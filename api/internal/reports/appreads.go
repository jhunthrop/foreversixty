package reports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	// maxObjectBytes bounds a report.json or fight summary read back out of the bucket: the
	// ingest never accepted a bundle larger than this.
	maxObjectBytes = 64 << 20
	// summaryReads is how many fight summaries one call reads at once.
	summaryReads = 4
	// The meters a fight sheet can sort by.
	meterDamage = "damage"
	meterHeal   = "healing"
	meterTaken  = "taken"
	// scopeTrash names a zone's trash group in a meters call: {n} is the group's key.
	scopeTrash = "trash"
	// foldTrash asks GET /v1/reports/{id} for the trash groups the apps draw.
	foldTrash = "trash"
)

// Objects reads one object out of the report bucket. *r2.Client satisfies it; nil means this
// deployment has no bucket, and the routes that need a fight summary answer 503.
type Objects interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
}

// Placer places a parse among the ranked kills of its bracket and reads a character's best
// earlier kill. *rankings.Store satisfies it; nil leaves every percentile null and no new best.
type Placer interface {
	Percentile(ctx context.Context, encounterID, difficulty int64, spec, at, metric string, value float64) (float64, int64, bool, error)
	BestBefore(ctx context.Context, key string, encounterID, difficulty int64, metric string, before time.Time, exceptReport string) (float64, bool, error)
}

// MeFight is one boss kill of one of the caller's characters (W3-C2).
type MeFight struct {
	Index       int     `json:"index"`
	EncounterID int64   `json:"encounter_id"`
	Encounter   string  `json:"encounter"`
	Kill        bool    `json:"kill"`
	DurationMS  int64   `json:"duration_ms"`
	Value       float64 `json:"value"`
	Unit        string  `json:"unit"`
	// Percentile is the parse, rounded as the web prints it; null when the bracket has no
	// ranked kill (the app prints no parse number then). Ranked is how many kills it holds.
	Percentile *int   `json:"percentile"`
	Ranked     *int64 `json:"ranked"`
	NewBest    bool   `json:"new_best"`
	Deaths     int    `json:"deaths"`
}

// MeCharacter is one of the caller's characters in the report, with its kills.
type MeCharacter struct {
	Key      string    `json:"key"`
	Name     string    `json:"name"`
	Class    string    `json:"class"`
	Spec     *string   `json:"spec"`
	SpecName string    `json:"spec_name"`
	Role     string    `json:"role"`
	Fights   []MeFight `json:"fights"`
}

// MeView is the body of GET /v1/reports/{id}/me.
type MeView struct {
	ReportID   string        `json:"report_id"`
	Characters []MeCharacter `json:"characters"`
}

// MeterPlayer is one row of a fight's meters (W3-C3).
type MeterPlayer struct {
	// Key is the player's character page, null when the report is not public or guild-visible
	// or the player has none.
	Key         *string `json:"key"`
	Name        string  `json:"name"`
	Class       string  `json:"class"`
	Spec        *string `json:"spec"`
	SpecName    string  `json:"spec_name"`
	Role        string  `json:"role"`
	ActiveMS    int64   `json:"active_ms"`
	DamageDone  int64   `json:"damage_done"`
	HealingDone int64   `json:"healing_done"`
	DamageTaken int64   `json:"damage_taken"`
	DPS         float64 `json:"dps"`
	HPS         float64 `json:"hps"`
	DTPS        float64 `json:"dtps"`
	Deaths      int     `json:"deaths"`
	Mine        bool    `json:"mine"`
}

// MetersView is the body of GET /v1/reports/{id}/fights/{n}/meters.
type MetersView struct {
	FightIndex int           `json:"fight_index"`
	Scope      string        `json:"scope"`
	Metric     string        `json:"metric"`
	Kill       bool          `json:"kill"`
	DurationMS int64         `json:"duration_ms"`
	Players    []MeterPlayer `json:"players"`
}

// MountApp registers the app's report reads (W3-C2, W3-C3).
func MountApp(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET /v1/reports/{id}/me", auth.RequireSession(s.me))
	mux.HandleFunc("GET /v1/reports/{id}/fights/{n}/meters", s.meters)
}

// readable is the path's report when the caller may see it. Anything else has already been
// answered: 404 for a report that does not exist or is not theirs to see (never 403), 404
// flagged for a flagged one they may see, 503 with no bucket.
func (s *Service) readable(w http.ResponseWriter, r *http.Request, op string) (Report, bool) {
	rep, err := s.Store.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, ErrNotFound) || (err == nil && !s.mayView(r, rep)) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return Report{}, false
	}
	if err != nil {
		s.fail(w, r, op, err, "could not load that report just now")
		return Report{}, false
	}
	if rep.Flagged != nil {
		httpx.WriteError(w, r, http.StatusNotFound, "flagged", "this report cannot be shown", nil)
		return Report{}, false
	}
	if s.Objects == nil {
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "unavailable",
			"report files are not configured on this deployment", nil)
		return Report{}, false
	}
	return rep, true
}

// me is GET /v1/reports/{id}/me: the caller's characters in the report and each one's boss
// kills, with the figure, the parse and whether it is a new best, so the app never joins
// history rows to fights itself. Every figure is read off the fight's stored summary, the
// file the web's report page reads, through the arithmetic the meters use.
func (s *Service) me(w http.ResponseWriter, r *http.Request) {
	rep, ok := s.readable(w, r, "me")
	if !ok {
		return
	}
	ctx := r.Context()
	mine, err := character.LinkedCharacters(ctx, s.Store.Pool, auth.ActorFrom(ctx).UserID)
	if err != nil {
		s.fail(w, r, "me", err, "could not read your results just now")
		return
	}
	fights, err := s.Store.Fights(ctx, rep.ID)
	if err != nil {
		s.fail(w, r, "me", err, "could not read your results just now")
		return
	}
	kills := make([]FightEntry, 0, len(fights))
	for _, f := range fights {
		if f.Kind == string(fight.Encounter) && f.Kill && !f.InProgress {
			kills = append(kills, f)
		}
	}
	view := MeView{ReportID: rep.ID, Characters: []MeCharacter{}}
	if len(mine) > 0 && len(kills) > 0 {
		sums, err := s.summaries(ctx, rep.ID, indexesOf(kills))
		if err != nil {
			s.fail(w, r, "me", err, "could not read your results just now")
			return
		}
		if view.Characters, err = s.myKills(ctx, rep, mine, kills, sums); err != nil {
			s.fail(w, r, "me", err, "could not read your results just now")
			return
		}
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// myKills builds the caller's characters from the kills' summaries, in the order the
// characters first appear, each character's kills in fight order.
func (s *Service) myKills(ctx context.Context, rep Report, mine []character.Linked,
	kills []FightEntry, sums []summary.Summary) ([]MeCharacter, error) {
	owned := make(map[string]character.Linked, len(mine))
	for _, c := range mine {
		owned[c.Key] = c
	}
	region, ruleset := ReportRealm(rep)
	var order []string
	byKey := map[string]*MeCharacter{}
	for i, kill := range kills {
		for _, row := range foldRosters(sums[i : i+1]) {
			key := character.KeyFromUnit(region, ruleset, row.Name)
			linked, ok := owned[key]
			if !ok {
				continue
			}
			c, seen := byKey[key]
			if !seen {
				c = newMeCharacter(linked, row)
				byKey[key] = c
				order = append(order, key)
			}
			f, err := s.myFight(ctx, rep, key, kill, row)
			if err != nil {
				return nil, err
			}
			c.Fights = append(c.Fights, f)
		}
	}
	out := make([]MeCharacter, 0, len(order))
	for _, k := range order {
		out = append(out, *byKey[k])
	}
	return out, nil
}

func newMeCharacter(linked character.Linked, row meterRow) *MeCharacter {
	class := linked.Class
	if class == "" {
		class = classSlug(row.Class)
	}
	return &MeCharacter{
		Key: linked.Key, Name: linked.Name, Class: class, Spec: nullable(specSlug(row.Class, row.Spec)),
		SpecName: row.Spec, Role: row.Role, Fights: []MeFight{},
	}
}

// myFight is one kill's line: the role figure, placed in its bracket the way the web's report
// page asks the rankings (the fight's phase, the row's spec, the rounded value).
func (s *Service) myFight(ctx context.Context, rep Report, key string, kill FightEntry, row meterRow) (MeFight, error) {
	metric, value := row.roleFigure()
	f := MeFight{
		Index: kill.Index, EncounterID: kill.EncounterID, Encounter: kill.Name, Kill: true,
		DurationMS: kill.DurationMS, Value: value, Unit: metric, Deaths: row.Deaths,
	}
	if s.Places == nil {
		return f, nil
	}
	if row.Spec != "" {
		pct, ranked, ok, err := s.Places.Percentile(ctx, kill.EncounterID, kill.Difficulty, row.Spec,
			phase.At(kill.Start), metric, value)
		if err != nil {
			return MeFight{}, fmt.Errorf("reports: place %s/%d: %w", rep.ID, kill.Index, err)
		}
		if ok {
			p := int(math.Round(pct))
			f.Percentile, f.Ranked = &p, &ranked
		}
	}
	best, found, err := s.Places.BestBefore(ctx, key, kill.EncounterID, kill.Difficulty, metric, kill.Start, rep.ID)
	if err != nil {
		return MeFight{}, fmt.Errorf("reports: best before %s/%d: %w", rep.ID, kill.Index, err)
	}
	f.NewBest = found && value > best
	return f, nil
}

// meters is GET /v1/reports/{id}/fights/{n}/meters[?metric=damage|healing|taken][&scope=trash]:
// one fight's players, or a zone's trash folded together (n is the group's key), sorted by the
// metric's per-second figure, high to low. A player who did none of the metric is left off, as
// Details! leaves them off. It follows the report's own visibility exactly.
func (s *Service) meters(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil || n < 1 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", msgFightIndex, map[string]string{"n": msgFightIndex})
		return
	}
	metric := r.URL.Query().Get("metric")
	if metric == "" {
		metric = meterDamage
	}
	if metric != meterDamage && metric != meterHeal && metric != meterTaken {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a meter",
			map[string]string{"metric": "one of damage, healing, taken"})
		return
	}
	scope := r.URL.Query().Get("scope")
	if scope != "" && scope != scopeTrash {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a scope",
			map[string]string{"scope": "trash, or leave it out"})
		return
	}
	rep, ok := s.readable(w, r, "meters")
	if !ok {
		return
	}
	view, found, err := s.meterView(r, rep, n, scope, metric)
	if err != nil {
		s.fail(w, r, "meters", err, "could not read that fight just now")
		return
	}
	if !found {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such fight", nil)
		return
	}
	if rep.Visibility == Public || rep.Visibility == Unlisted {
		httpx.CachePublic(w, reportsObjectMaxAge, reportsObjectStale)
	} else {
		httpx.CachePrivate(w)
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// meterView reads and folds the pulls a meters call names. found is false for a fight the
// report does not have, a fight still in progress, or a trash key that names no group.
func (s *Service) meterView(r *http.Request, rep Report, n int, scope, metric string) (MetersView, bool, error) {
	ctx := r.Context()
	fights, err := s.Store.Fights(ctx, rep.ID)
	if err != nil {
		return MetersView{}, false, err
	}
	view := MetersView{FightIndex: n, Scope: scope, Metric: metric}
	var pulls []int
	if scope == scopeTrash {
		entries, err := s.reportFights(ctx, rep.ID)
		if err != nil {
			return MetersView{}, false, err
		}
		group, ok := trashGroupKeyed(TrashGroups(entries), n)
		if !ok {
			return MetersView{}, false, nil
		}
		pulls, view.DurationMS = group.Fights, group.DurationMS
	} else {
		i := slices.IndexFunc(fights, func(f FightEntry) bool { return f.Index == n && !f.InProgress })
		if i < 0 {
			return MetersView{}, false, nil
		}
		pulls, view.Kill, view.DurationMS = []int{n}, fights[i].Kill, fights[i].DurationMS
	}
	sums, err := s.summaries(ctx, rep.ID, pulls)
	if err != nil {
		return MetersView{}, false, err
	}
	if view.Players, err = s.meterPlayers(ctx, rep, foldRosters(sums), metric); err != nil {
		return MetersView{}, false, err
	}
	return view, true, nil
}

// meterPlayers shapes, filters and sorts the folded rows.
func (s *Service) meterPlayers(ctx context.Context, rep Report, rows []meterRow, metric string) ([]MeterPlayer, error) {
	region, ruleset := ReportRealm(rep)
	keys := make([]string, len(rows))
	for i, row := range rows {
		keys[i] = character.KeyFromUnit(region, ruleset, row.Name)
	}
	mine, err := s.callerKeys(ctx)
	if err != nil {
		return nil, err
	}
	pages, err := s.characterPages(ctx, rep, keys)
	if err != nil {
		return nil, err
	}
	out := make([]MeterPlayer, 0, len(rows))
	for i, row := range rows {
		p := MeterPlayer{
			Name: unitName(row.Name), Class: classSlug(row.Class), Spec: nullable(specSlug(row.Class, row.Spec)),
			SpecName: row.Spec, Role: row.Role, ActiveMS: row.ActiveMS,
			DamageDone: row.Damage, HealingDone: row.Healing, DamageTaken: row.Taken,
			DPS: row.dps(), HPS: row.hps(), DTPS: row.dtps(), Deaths: row.Deaths, Mine: mine[keys[i]],
		}
		if pages[keys[i]] {
			p.Key = &keys[i]
		}
		if meterValue(p, metric) > 0 {
			out = append(out, p)
		}
	}
	slices.SortStableFunc(out, func(a, b MeterPlayer) int {
		if va, vb := meterValue(a, metric), meterValue(b, metric); va != vb {
			if va > vb {
				return -1
			}
			return 1
		}
		return strings.Compare(a.Name, b.Name)
	})
	return out, nil
}

func meterValue(p MeterPlayer, metric string) float64 {
	switch metric {
	case meterHeal:
		return p.HPS
	case meterTaken:
		return p.DTPS
	}
	return p.DPS
}

// callerKeys is the set of the signed-in caller's character keys; empty when signed out.
func (s *Service) callerKeys(ctx context.Context) (map[string]bool, error) {
	a := auth.ActorFrom(ctx)
	out := map[string]bool{}
	if !a.Signed() {
		return out, nil
	}
	linked, err := character.LinkedCharacters(ctx, s.Store.Pool, a.UserID)
	if err != nil {
		return nil, err
	}
	for _, c := range linked {
		out[c.Key] = true
	}
	return out, nil
}

// characterPages is which keys open a character page from this report's meters: only on a
// public or guild-visible report, and only for a character the site has a page for (a known
// character, or one with ranked fights).
func (s *Service) characterPages(ctx context.Context, rep Report, keys []string) (map[string]bool, error) {
	out := map[string]bool{}
	if (rep.Visibility != Public && rep.Visibility != GuildTo) || len(keys) == 0 {
		return out, nil
	}
	rows, err := s.Store.Pool.Query(ctx,
		`select key from characters where key = any($1)
		 union
		 select distinct player_key from fight_metrics where player_key = any($1) and state = 'ok'`, keys)
	if err != nil {
		return nil, fmt.Errorf("reports: character pages: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var k string
		if err := rows.Scan(&k); err != nil {
			return nil, fmt.Errorf("reports: character pages: %w", err)
		}
		out[k] = true
	}
	return out, rows.Err()
}

// unitName is a log unit's character name without its realm ("Baelgrim-Nightslayer" is
// "Baelgrim").
func unitName(unit string) string {
	name, _ := character.SplitUnit(unit)
	return name
}

func nullable(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func indexesOf(fights []FightEntry) []int {
	out := make([]int, len(fights))
	for i, f := range fights {
		out[i] = f.Index
	}
	return out
}

// summaries reads the stored summaries of the given fights, a few at a time, in the given
// order. Any failed read fails the call: a figure is never quietly left out.
func (s *Service) summaries(ctx context.Context, reportID string, indexes []int) ([]summary.Summary, error) {
	out := make([]summary.Summary, len(indexes))
	errs := make([]error, len(indexes))
	sem := make(chan struct{}, summaryReads)
	var wg sync.WaitGroup
	for i, n := range indexes {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			errs[i] = s.readJSON(ctx, Keys(reportID).FightSummary(n), &out[i])
		}()
	}
	wg.Wait()
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}
	return out, nil
}

// reportFights reads the fights of the report's report.json: the engine's own list, which
// carries each fight's zone (the database does not).
func (s *Service) reportFights(ctx context.Context, reportID string) ([]store.FightEntry, error) {
	var rep store.Report
	if err := s.readJSON(ctx, Keys(reportID).Report(), &rep); err != nil {
		return nil, err
	}
	return rep.Fights, nil
}

func (s *Service) readJSON(ctx context.Context, key string, into any) error {
	body, err := s.Objects.Get(ctx, key)
	if err != nil {
		return fmt.Errorf("reports: read %s: %w", key, err)
	}
	defer body.Close()
	b, err := io.ReadAll(io.LimitReader(body, maxObjectBytes))
	if err != nil {
		return fmt.Errorf("reports: read %s: %w", key, err)
	}
	if err := json.Unmarshal(b, into); err != nil {
		return fmt.Errorf("reports: decode %s: %w", key, err)
	}
	return nil
}

// trashGroupsFor is the trash groups of a report, for GET /v1/reports/{id}?fold=trash. With no
// bucket, or a report.json that cannot be read, there are none and the apps draw each trash
// pull on its own row, as they do for zone-less trash; the failure is logged, never hidden.
func (s *Service) trashGroupsFor(ctx context.Context, r *http.Request, reportID string) []TrashGroup {
	if s.Objects == nil {
		return []TrashGroup{}
	}
	entries, err := s.reportFights(ctx, reportID)
	if err != nil {
		s.logger().Warn("reports", "id", httpx.RequestIDFrom(r.Context()), "op", "trash groups",
			"report", reportID, "err", err)
		return []TrashGroup{}
	}
	return TrashGroups(entries)
}
