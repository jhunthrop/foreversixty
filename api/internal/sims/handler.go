package sims

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

// filterKinds is the five kinds a history filter accepts, in the
// contract's own order (contract 1.1). Composed from the envelope's
// constants rather than spelled out, so the vocabulary has one source.
// Replace this with simapi.Kinds if the sim module ever publishes one.
var filterKinds = []string{
	simapi.KindRun, simapi.KindGear, simapi.KindTalents,
	simapi.KindDrops, simapi.KindWeights,
}

const (
	// maxResultBytes bounds a posted browser result. A 3,000-iteration
	// summary with per-ability rows, aura tracks and a cast timeline is
	// a few hundred kilobytes; two megabytes is room for the precision
	// run and still a bound.
	maxResultBytes = 2 << 20
	// maxRequestBytes bounds a POST /v1/sims/run body. The envelope is
	// JSON with a gear list in it, so 256 KB is generous.
	maxRequestBytes = 256 << 10
	// fetchMaxAge is how long a saved sim may be cached, in seconds:
	// it never changes once written.
	fetchMaxAge = 86400
	// SimRunJobCommand is the argument the image dispatches on for one
	// premium run, the way ParseJobCommand does for an upload.
	SimRunJobCommand = "sim-run"
	// maxTitle bounds the name a member may give a saved sim, in
	// runes.
	maxTitle = 120
)

// Service serves the simulator routes.
type Service struct {
	Store *Store
	// Accounts reads the premium flag. Nil means the premium lane is
	// not offered.
	Accounts Premiumer
	// Jobs starts the premium lane's Cloud Run job. Nil means the
	// deployment cannot reach it, and the premium lane is not offered.
	Jobs jobs.Runner
	// Planner counts a bulk or weights request before it is queued. Nil
	// means the premium lane cannot size one, and the run route is not
	// mounted: accepting a request we cannot bound would be worse than
	// not offering the route.
	Planner runner.Planner
	// Summaries reads stored fight summaries, for the buffs a
	// character's last fight recorded. Nil means no bucket, and
	// sim-input answers without them.
	Summaries Getter
	// EngineVersion is the pinned engine build this deployment runs.
	EngineVersion string
	Log           *slog.Logger
}

// Mount registers every simulator route. POST /v1/sims/run is mounted
// only when this deployment can actually dispatch it; without the job
// runner, the accounts store or the planner there is nothing behind
// it, and a 404 is a truer answer than a 500.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("POST /v1/sims", s.save)
	mux.HandleFunc("GET /v1/sims", auth.RequireSession(s.mine))
	mux.HandleFunc("GET /v1/sims/{id}", s.get)
	mux.HandleFunc("GET /v1/sims/{id}/progress", s.progress)
	mux.HandleFunc("GET /v1/specs", s.specs)
	mux.HandleFunc("GET /v1/characters/{region}/{ruleset}/{name}/sim-input", s.simInput)
	if s.Jobs != nil && s.Accounts != nil && s.Planner != nil {
		mux.HandleFunc("POST /v1/sims/run", auth.RequireSession(s.run))
	}
}

// loggerOr returns l, or the package's shared default when l is nil.
// Every deps struct in this package that carries an optional
// *slog.Logger falls back the same way; this is the one place that
// says so.
func loggerOr(l *slog.Logger) *slog.Logger {
	if l != nil {
		return l
	}
	return slog.Default()
}

func (s *Service) logger() *slog.Logger {
	return loggerOr(s.Log)
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("sims", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

// SaveInput is the body of POST /v1/sims: a browser-run result, with
// an optional name for the member's own history.
type SaveInput struct {
	simapi.SimResult
	Title string `json:"title"`
}

func (s *Service) save(w http.ResponseWriter, r *http.Request) {
	var in SaveInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxResultBytes)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"the body must be a sim result of at most 2 MB", nil)
		return
	}
	if in.Lane != simapi.LaneBrowser {
		// The server lane writes its own rows, from the job. A client
		// claiming a server result here would be claiming compute it
		// never paid for.
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"only a browser-run result is saved this way",
			map[string]string{"lane": simapi.LaneBrowser})
		return
	}
	in.Request.Encounter = withEncounterDefaults(in.Request.Encounter)
	// A save is not a run: the engine that produced this result need
	// not be this deployment's pin (ValidateSaved, not Validate) - a
	// member's cached browser bundle can be a build behind without
	// losing the sim they just ran.
	if err := in.ValidateSaved(); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(), nil)
		return
	}
	if in.EngineVersion == "" {
		in.EngineVersion = in.Request.EngineVersion
	}
	var owner *int64
	if a := auth.ActorFrom(r.Context()); a.Signed() {
		id := a.UserID
		owner = &id
	}
	id := auth.Base32ID(auth.ReportIDChars)
	if err := s.Store.Save(r.Context(), id, owner, trimTitle(in.Title), in.SimResult); err != nil {
		s.fail(w, r, "save", err, "could not save that sim just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusCreated, map[string]string{"sim_id": id})
}

func (s *Service) get(w http.ResponseWriter, r *http.Request) {
	res, err := s.Store.Get(r.Context(), r.PathValue("id"))
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such sim", nil)
	case err != nil:
		s.fail(w, r, "get", err, "could not read that sim just now")
	default:
		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(fetchMaxAge))
		httpx.WriteOK(w, r, http.StatusOK, res)
	}
}

func (s *Service) mine(w http.ResponseWriter, r *http.Request) {
	if !httpx.RequireMine(w, r) {
		return
	}
	page, ok := httpx.ParsePage(w, r)
	if !ok {
		return
	}
	// An unknown kind is refused rather than answered with an empty
	// list: "you have no topgear sims" is a true sentence about a word
	// that does not exist, and a typo in a link would read as a real
	// (empty) answer forever.
	kind := r.URL.Query().Get("kind")
	if kind != "" && !slices.Contains(filterKinds, kind) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"kind must be one of "+strings.Join(filterKinds, ", "),
			map[string]string{"kind": strings.Join(filterKinds, "|")})
		return
	}
	out, err := s.Store.Mine(r.Context(), auth.ActorFrom(r.Context()).UserID, page, kind)
	if err != nil {
		s.fail(w, r, "mine", err, "could not read your sims just now")
		return
	}
	httpx.SetPrivateListCache(w)
	httpx.WriteOK(w, r, http.StatusOK, out)
}

func (s *Service) progress(w http.ResponseWriter, r *http.Request) {
	p, err := s.Store.Progress(r.Context(), r.PathValue("id"))
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such sim", nil)
	case err != nil:
		s.fail(w, r, "progress", err, "could not read that run just now")
	default:
		// A run in flight changes every second; nothing here may be
		// cached at the edge.
		w.Header().Set("Cache-Control", "no-store")
		httpx.WriteOK(w, r, http.StatusOK, p)
	}
}

// trimTitle bounds the name a member gave a sim. A long title is cut
// rather than refused: it is a label, not data.
func trimTitle(s string) string { return trimRunes(s, maxTitle) }

// trimRunes cuts s to at most max runes. The cut is by rune, so a string
// in any script keeps its last character whole and the result is always
// valid UTF-8. Both the member's title and the composed headline are
// bounded this way.
func trimRunes(s string, max int) string {
	r := []rune(s)
	if len(r) > max {
		return string(r[:max])
	}
	return s
}
