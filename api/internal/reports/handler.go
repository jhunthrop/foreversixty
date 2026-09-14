package reports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
)

const (
	// maxJSONBody is the ceiling on the small JSON bodies here.
	maxJSONBody = 8 << 10
	// visibilityMaxAge is how long the site Worker may cache a report's
	// visibility, as the contract sets it.
	visibilityMaxAge = 60
	// MinePerPage is the page size of the caller's own report list,
	// the same hundred the rankings page uses.
	MinePerPage = 100
	// officer ranks that may edit a guild's reports.
	rankOfficer = "officer"
	rankLeader  = "leader"
)

// Signer hands out signed URLs for a report's files. *r2.Client
// satisfies it; the tests use a stub.
type Signer interface {
	PresignGet(ctx context.Context, key string, ttl time.Duration) (string, error)
}

// Accounts is the part of auth.Store the reports handlers need.
type Accounts interface {
	User(ctx context.Context, id int64) (auth.User, error)
	GuildRank(ctx context.Context, guildID, userID int64) (string, bool, error)
}

// Service serves the report routes.
type Service struct {
	Store    *Store
	Accounts Accounts
	Signer   Signer
	// Rank withdraws a report's ranking rows when a patch takes the
	// report out of the visibilities that may rank. Nil leaves the
	// rows in place, which is only ever right in a test that is not
	// looking at rankings.
	Rank Ranker
	// PublicBaseURL is the site, which serves public report files from
	// the bucket at /logs-data/.
	PublicBaseURL string
	// APIBaseURL is this service, which serves a private report's files
	// through signed redirects.
	APIBaseURL string
	Log        *slog.Logger
}

// Mount registers the report routes that are not ingest.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("POST /v1/reports", auth.Require(s.create))
	mux.HandleFunc("GET /v1/reports", auth.RequireSession(s.mine))
	mux.HandleFunc("GET /v1/reports/{id}", s.get)
	mux.HandleFunc("PATCH /v1/reports/{id}", auth.Require(s.patch))
	mux.HandleFunc("GET /v1/reports/{id}/visibility", s.visibility)
	mux.HandleFunc("GET /v1/reports/{id}/access", auth.RequireSession(s.access))
	mux.HandleFunc("GET /v1/reports/{id}/files/{path...}", s.file)
	mux.HandleFunc("GET /reports/{id}/card.png", s.card)
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("reports", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

// CreateInput is the body of POST /v1/reports.
type CreateInput struct {
	Title            string `json:"title"`
	Visibility       string `json:"visibility"`
	Zone             string `json:"zone"`
	LoggingCharacter *struct {
		Region  string `json:"region"`
		Ruleset string `json:"ruleset"`
		Name    string `json:"name"`
	} `json:"logging_character"`
}

func (s *Service) create(w http.ResponseWriter, r *http.Request) {
	var in CreateInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxJSONBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	if !ValidVisibility(in.Visibility) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a visibility",
			map[string]string{"visibility": "one of public, unlisted, private, guild"})
		return
	}
	a := auth.ActorFrom(r.Context())
	owner := a.UserID
	rep := Report{
		ID: auth.NewReportID(), OwnerID: &owner, Title: strings.TrimSpace(in.Title),
		Visibility: in.Visibility, Zone: strings.TrimSpace(in.Zone), Status: StatusLive,
	}
	if c := in.LoggingCharacter; c != nil && c.Name != "" {
		if !character.ValidRegion(c.Region) || !character.ValidRuleset(c.Ruleset) {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a character",
				map[string]string{"logging_character": "region must be one of us, eu, kr, tw, cn and ruleset one of normal, pvp, rp, hardcore"})
			return
		}
		key := character.Key(c.Region, c.Ruleset, c.Name)
		rep.LoggingCharacter = &key
	}
	stored, err := s.Store.Create(r.Context(), rep)
	if err != nil {
		s.fail(w, r, "create", err, "could not start that report just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusCreated, map[string]any{
		"id": stored.ID, "created_at": stored.CreatedAt,
	})
}

// Summary is one row of the signed-in person's own report list.
type Summary struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Zone       string    `json:"zone"`
	Status     string    `json:"status"`
	Visibility string    `json:"visibility"`
	CreatedAt  time.Time `json:"created_at"`
	FightCount int       `json:"fight_count"`
	// KillCount is how many of those fights were boss kills, which is
	// what the list shows beside the count: "8 fights, 3 kills".
	KillCount int `json:"kill_count"`
}

// MinePage is the body of GET /v1/reports?mine=1.
type MinePage struct {
	Rows    []Summary `json:"rows"`
	Total   int       `json:"total"`
	Page    int       `json:"page"`
	PerPage int       `json:"per_page"`
}

// mine lists the caller's own reports, newest first. It is the only
// listing this service offers: there is no browsing other people's.
func (s *Service) mine(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("mine") != "1" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"this route lists your own reports", map[string]string{"mine": "pass mine=1"})
		return
	}
	page := 1
	if v := r.URL.Query().Get("page"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "page must be 1 or more",
				map[string]string{"page": "a page number from 1"})
			return
		}
		page = n
	}
	rows, total, err := s.Store.OwnedBy(r.Context(), auth.ActorFrom(r.Context()).UserID, page, MinePerPage)
	if err != nil {
		s.fail(w, r, "mine", err, "could not list your reports just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, MinePage{
		Rows: rows, Total: total, Page: page, PerPage: MinePerPage,
	})
}

func (s *Service) get(w http.ResponseWriter, r *http.Request) {
	rep, err := s.Store.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "get", err, "could not load that report just now")
		return
	}
	if !s.mayView(r, rep) {
		// A private report answers 404 rather than 403: whether a
		// report exists is itself private.
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	view, err := s.view(r.Context(), rep)
	if err != nil {
		s.fail(w, r, "get", err, "could not load that report just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// view assembles the report body the page reads.
func (s *Service) view(ctx context.Context, rep Report) (View, error) {
	fights, err := s.Store.Fights(ctx, rep.ID)
	if err != nil {
		return View{}, err
	}
	players, err := s.Store.Players(ctx, rep.ID)
	if err != nil {
		return View{}, err
	}
	v := View{
		ID: rep.ID, Title: rep.Title, Visibility: rep.Visibility, Zone: rep.Zone,
		Status: rep.Status, EngineVersion: rep.EngineVersion, Fights: fights, Players: players,
		CreatedAt: rep.CreatedAt, DataBaseURL: s.dataBaseURL(rep),
	}
	if rep.Flagged != nil {
		v.Flagged = *rep.Flagged
	}
	if rep.OwnerID != nil {
		owner := Owner{ID: *rep.OwnerID}
		if u, err := s.Accounts.User(ctx, *rep.OwnerID); err == nil {
			// PublicName, never Name: this body is served to anyone
			// holding the link of a public or unlisted report, and an
			// account signed in by email has no battletag to show.
			owner.Battletag = u.PublicName()
		}
		v.Owner = &owner
	}
	if rep.GuildID != nil {
		g, err := s.guild(ctx, *rep.GuildID)
		if err != nil {
			return View{}, err
		}
		v.Guild = g
	}
	return v, nil
}

func (s *Service) guild(ctx context.Context, id int64) (*GuildRef, error) {
	g := GuildRef{ID: id}
	err := s.Store.Pool.QueryRow(ctx,
		`select name, region, ruleset from guilds where id = $1`, id).Scan(&g.Name, &g.Region, &g.Ruleset)
	if err != nil {
		return nil, fmt.Errorf("reports: read guild %d: %w", id, err)
	}
	return &g, nil
}

// dataBaseURL is where the island reads the report's files. Public and
// unlisted reports are served straight off the bucket by the site
// Worker; everything else goes through this service's signed redirects,
// which /v1/reports/{id}/access hands out.
func (s *Service) dataBaseURL(rep Report) string {
	if rep.Visibility == Public || rep.Visibility == Unlisted {
		return s.PublicBaseURL + "/logs-data/reports/" + rep.ID
	}
	return s.APIBaseURL + "/v1/reports/" + rep.ID + "/files"
}

func (s *Service) patch(w http.ResponseWriter, r *http.Request) {
	rep, err := s.Store.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "patch", err, "could not change that report just now")
		return
	}
	if !s.mayView(r, rep) {
		// Same rule as every read path: whether a private report exists
		// is itself private, so a caller who could not even see it gets
		// the same 404 as a caller patching an id that does not exist -
		// never the 403 that would confirm it is there.
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	may, err := s.mayEdit(r.Context(), auth.ActorFrom(r.Context()), rep)
	if err != nil {
		s.fail(w, r, "patch", err, "could not change that report just now")
		return
	}
	if !may {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "that report is not yours to change", nil)
		return
	}
	var p Patch
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxJSONBody)).Decode(&p); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	if p.Visibility != nil && !ValidVisibility(*p.Visibility) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a visibility",
			map[string]string{"visibility": "one of public, unlisted, private, guild"})
		return
	}
	if p.GuildID != nil {
		// Setting guild_id hands that guild's officers edit rights on
		// this report (via mayEdit) and its members read rights once
		// visibility is guild, so the caller needs standing in the
		// guild they are naming, not just in the report they own. A
		// moderator keeps the same wider reach mayEdit already gives
		// them elsewhere in this handler.
		a := auth.ActorFrom(r.Context())
		if !a.IsModerator() {
			rank, ok, err := s.Accounts.GuildRank(r.Context(), *p.GuildID, a.UserID)
			if err != nil {
				s.fail(w, r, "patch", err, "could not change that report just now")
				return
			}
			if !ok || (rank != rankOfficer && rank != rankLeader) {
				httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
					"you must be an officer of that guild to attach a report to it", nil)
				return
			}
		}
	}
	if p.Title != nil {
		trimmed := strings.TrimSpace(*p.Title)
		p.Title = &trimmed
	}
	// A report that stops being rankable takes its ranking rows with
	// it: the body 404s from here on, and leaving the rows would keep
	// publishing the player keys, names, guild and numbers the owner
	// just withdrew. Only the crossing is acted on, so patching a
	// report that is already private is not a second withdrawal, and
	// a patch that does not touch visibility withdraws nothing. The
	// way back is deliberately one-way: re-ranking a report returned
	// to public would have to re-derive every fight, and nothing asks
	// for it.
	//
	// This runs before the update, not after. A withdrawal that fails
	// then leaves a report that is still public and still rankable, so
	// the retry crosses again and withdraws again; the other order
	// would leave a private report's rows on the leaderboards with
	// nothing left to notice them.
	if s.Rank != nil && p.Visibility != nil && Ranked(rep.Visibility) && !Ranked(*p.Visibility) {
		if err := s.Rank.RemoveReport(r.Context(), rep.ID, ReasonNotRankable); err != nil {
			s.fail(w, r, "patch", err, "could not change that report just now")
			return
		}
	}
	updated, err := s.Store.Update(r.Context(), rep.ID, p)
	if err != nil {
		s.fail(w, r, "patch", err, "could not change that report just now")
		return
	}
	view, err := s.view(r.Context(), updated)
	if err != nil {
		s.fail(w, r, "patch", err, "could not change that report just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// visibility is what the site Worker asks before serving a report's
// files off the bucket.
func (s *Service) visibility(w http.ResponseWriter, r *http.Request) {
	rep, err := s.Store.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "visibility", err, "could not read that report just now")
		return
	}
	w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(visibilityMaxAge))
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"visibility": rep.Visibility})
}

// access hands a signed-in reader a base URL for a report whose files
// the Worker refuses to serve.
//
// The contract calls for "signed R2 URLs valid 10 minutes" behind one
// data_base_url. One signature cannot cover a whole prefix, and the
// island reads report.json and then a summary per fight on demand, so
// the base URL points at this service's own file route, which checks
// the session again and redirects each request to a freshly signed URL
// for that one object. The island's code is then the same for a public
// report and a private one.
func (s *Service) access(w http.ResponseWriter, r *http.Request) {
	rep, err := s.Store.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "access", err, "could not open that report just now")
		return
	}
	if !s.mayView(r, rep) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{
		"data_base_url": s.APIBaseURL + "/v1/reports/" + rep.ID + "/files",
		"expires_in":    int(r2.AccessTTL.Seconds()),
	})
}

// file redirects one of a report's objects to a signed R2 URL. Only the
// files a report page reads are reachable: raw chunks are private to
// their owner and are not served here.
func (s *Service) file(w http.ResponseWriter, r *http.Request) {
	rep, err := s.Store.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "file", err, "could not open that file just now")
		return
	}
	if !s.mayView(r, rep) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	path := r.PathValue("path")
	if !readablePath(path) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such file", nil)
		return
	}
	if s.Signer == nil {
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "unavailable",
			"report files are not configured on this deployment", nil)
		return
	}
	// The key is built from store.Keys out of the report's own id, never
	// a raw concatenation of the request's path: a request can only ever
	// name one of the objects readablePath has already accepted, and
	// never reach outside that report's own prefix.
	key, ok := objectKey(rep.ID, path)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such file", nil)
		return
	}
	url, err := s.Signer.PresignGet(r.Context(), key, r2.AccessTTL)
	if err != nil {
		s.fail(w, r, "file", err, "could not open that file just now")
		return
	}
	w.Header().Set("Cache-Control", "private, max-age=60")
	http.Redirect(w, r, url, http.StatusFound)
}

// readablePath allows exactly the objects the report page reads:
// report.json and a fight's summary, events, or live snapshot. Anything
// else - raw chunks above all - is refused, and no path may climb.
func readablePath(path string) bool {
	if path == "" || strings.Contains(path, "..") {
		return false
	}
	if path == "report.json" {
		return true
	}
	_, file, ok := fightPath(path)
	if !ok {
		return false
	}
	switch file {
	case "summary.json", "events.parquet", "live.json":
		return true
	}
	return false
}

// fightPath parses the fights/<n>/<file> shape of a path, which is the
// only shape readablePath accepts besides the bare report.json.
func fightPath(path string) (index int, file string, ok bool) {
	rest, ok := strings.CutPrefix(path, "fights/")
	if !ok {
		return 0, "", false
	}
	indexStr, file, ok := strings.Cut(rest, "/")
	if !ok {
		return 0, "", false
	}
	n, err := strconv.Atoi(indexStr)
	if err != nil {
		return 0, "", false
	}
	return n, file, true
}

// objectKey turns a path readablePath has already accepted into the R2
// object key it names, built from store.Keys out of reportID rather
// than the path itself - the r2 client does not confine keys on its
// own, so nothing here may hand it a key assembled from request input.
func objectKey(reportID, path string) (string, bool) {
	keys := Keys(reportID)
	if path == "report.json" {
		return keys.Report(), true
	}
	n, file, ok := fightPath(path)
	if !ok {
		return "", false
	}
	switch file {
	case "summary.json":
		return keys.FightSummary(n), true
	case "events.parquet":
		return keys.FightEvents(n), true
	case "live.json":
		return keys.FightLive(n), true
	}
	return "", false
}

// mayView applies the visibility rules. Public and unlisted reports are
// readable by anyone with the link; a private one by its owner and
// moderators; a guild one by the guild as well.
func (s *Service) mayView(r *http.Request, rep Report) bool {
	if rep.Visibility == Public || rep.Visibility == Unlisted {
		return true
	}
	a := auth.ActorFrom(r.Context())
	if !a.Signed() {
		return false
	}
	if a.IsModerator() || (rep.OwnerID != nil && *rep.OwnerID == a.UserID) {
		return true
	}
	if rep.Visibility == GuildTo && rep.GuildID != nil {
		if _, ok, err := s.Accounts.GuildRank(r.Context(), *rep.GuildID, a.UserID); err == nil && ok {
			return true
		}
	}
	return false
}

// mayEdit is the owner, a moderator, or an officer of the report's guild.
func (s *Service) mayEdit(ctx context.Context, a auth.Actor, rep Report) (bool, error) {
	if !a.Signed() {
		return false, nil
	}
	if a.IsModerator() || (rep.OwnerID != nil && *rep.OwnerID == a.UserID) {
		return true, nil
	}
	if rep.GuildID == nil {
		return false, nil
	}
	rank, ok, err := s.Accounts.GuildRank(ctx, *rep.GuildID, a.UserID)
	if err != nil || !ok {
		return false, err
	}
	return rank == rankOfficer || rank == rankLeader, nil
}
