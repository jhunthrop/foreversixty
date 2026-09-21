package rankings

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
)

// cacheSeconds is how long a rankings answer may be cached at the edge.
// The design's caching note sets it at thirty seconds.
const cacheSeconds = 30

// Service serves the rankings, character, and guild routes.
type Service struct {
	Store *Store
	Log   *slog.Logger
	// Now is the clock, so a test can fix "today".
	Now func() time.Time
}

// Mount registers the read routes.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET /v1/rankings", s.rankings)
	mux.HandleFunc("GET /v1/rankings/percentile", s.percentile)
	mux.HandleFunc("GET /v1/rankings/guilds", s.guildRankings)
	mux.HandleFunc("GET /v1/encounters", s.encounters)
	mux.HandleFunc("GET /v1/characters/{region}/{ruleset}/{name}", s.character)
	mux.HandleFunc("GET /v1/guilds/{region}/{ruleset}/{name}", s.guild)
}

func (s *Service) now() time.Time {
	if s.Now != nil {
		return s.Now()
	}
	return time.Now().UTC()
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("rankings", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

// cache marks a read cacheable for the edge.
func cache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(cacheSeconds))
}

// encounterOf reads the `encounter` parameter, which may be the
// numeric id or the slug the site's own URL carries. It writes the
// response itself when there is nothing to rank.
func (s *Service) encounterOf(w http.ResponseWriter, r *http.Request, v string) (int64, bool) {
	if v == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "encounter is required",
			map[string]string{"encounter": "the encounter id or its slug"})
		return 0, false
	}
	id, err := s.Store.ResolveEncounter(r.Context(), v)
	if errors.Is(err, ErrNoEncounter) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such encounter", nil)
		return 0, false
	}
	if err != nil {
		s.fail(w, r, "encounter", err, "could not read the rankings just now")
		return 0, false
	}
	return id, true
}

func (s *Service) rankings(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	encounter, ok := s.encounterOf(w, r, q.Get("encounter"))
	if !ok {
		return
	}
	query := Query{
		EncounterID: encounter, Metric: MetricDPS, Spec: q.Get("spec"), Class: q.Get("class"),
		Phase: q.Get("phase"), Region: strings.ToLower(q.Get("region")),
		Ruleset: strings.ToLower(q.Get("ruleset")), Faction: strings.ToLower(q.Get("faction")),
		Since: q.Get("since"),
	}
	if v := q.Get("metric"); v != "" {
		if !ValidMetric(v) {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a metric",
				map[string]string{"metric": "one of " + strings.Join(Metrics, ", ")})
			return
		}
		query.Metric = v
	}
	if v := q.Get("difficulty"); v != "" {
		d, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "difficulty must be a number",
				map[string]string{"difficulty": "the encounter's difficulty id"})
			return
		}
		query.Difficulty = &d
	}
	if query.Phase != "" && !phase.Valid(query.Phase) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a phase",
			map[string]string{"phase": strings.Join(phase.Names(), ", ")})
		return
	}
	if query.Ruleset != "" && !character.ValidRuleset(query.Ruleset) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a ruleset",
			map[string]string{"ruleset": strings.Join(character.Rulesets, ", ")})
		return
	}
	if query.Region != "" && !character.ValidRegion(query.Region) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a region",
			map[string]string{"region": strings.Join(character.Regions, ", ")})
		return
	}
	if v := q.Get("page"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil || p < 1 {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "page must be 1 or more",
				map[string]string{"page": "a page number from 1"})
			return
		}
		query.Page = p
	}
	page, err := s.Store.Rankings(r.Context(), query, s.now())
	if err != nil {
		if strings.HasPrefix(err.Error(), "since must be") {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(),
				map[string]string{"since": "today, or a number of days like 7d"})
			return
		}
		s.fail(w, r, "rankings", err, "could not read the rankings just now")
		return
	}
	cache(w)
	httpx.WriteOK(w, r, http.StatusOK, page)
}

func (s *Service) percentile(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	encounter, ok := s.encounterOf(w, r, q.Get("encounter"))
	if !ok {
		return
	}
	difficulty, err := strconv.ParseInt(q.Get("difficulty"), 10, 64)
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "difficulty is required",
			map[string]string{"difficulty": "the encounter's difficulty id"})
		return
	}
	value, err := strconv.ParseFloat(q.Get("value"), 64)
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "value is required",
			map[string]string{"value": "the metric value to place"})
		return
	}
	metric := q.Get("metric")
	if metric == "" {
		metric = MetricDPS
	}
	// Not ValidMetric: a percentile is a place on a folded curve, and
	// percentile_digests holds no curve for execution scores.
	if !ValidDigestMetric(metric) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a metric",
			map[string]string{"metric": "one of " + strings.Join(DigestedMetrics, ", ")})
		return
	}
	at := q.Get("phase")
	if at == "" {
		at = phase.At(s.now())
	}
	pct, ranked, ok, err := s.Store.Percentile(r.Context(), encounter, difficulty, q.Get("spec"), at, metric, value)
	if err != nil {
		s.fail(w, r, "percentile", err, "could not place that parse just now")
		return
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found",
			"nothing has been ranked in that bracket yet", nil)
		return
	}
	cache(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"percentile": pct, "ranked": ranked})
}

func (s *Service) guildRankings(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	kind := q.Get("kind")
	if kind == "" {
		kind = KindSpeed
	}
	if !ValidKind(kind) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a guild ranking",
			map[string]string{"kind": "one of speed, execution, progress"})
		return
	}
	var encounter int64
	if v := q.Get("encounter"); v != "" {
		id, ok := s.encounterOf(w, r, v)
		if !ok {
			return
		}
		encounter = id
	}
	if kind != KindProgress && encounter == 0 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"speed and execution rank one encounter",
			map[string]string{"encounter": "the encounter id or its slug"})
		return
	}
	at := q.Get("phase")
	if at != "" && !phase.Valid(at) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a phase",
			map[string]string{"phase": strings.Join(phase.Names(), ", ")})
		return
	}
	rows, err := s.Store.GuildRankings(r.Context(), encounter, kind, at)
	if err != nil {
		s.fail(w, r, "guild rankings", err, "could not read the guild rankings just now")
		return
	}
	cache(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": rows})
}

// encounters lists every encounter any report has ever ranked, for the
// /rankings picker: a visitor with no encounter to name in the URL
// cannot ask GET /v1/rankings for one, so the picker reads this route
// first to offer the ones that exist.
func (s *Service) encounters(w http.ResponseWriter, r *http.Request) {
	rows, err := s.Store.Encounters(r.Context())
	if err != nil {
		s.fail(w, r, "encounters", err, "could not read the encounters just now")
		return
	}
	cache(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": rows})
}

func (s *Service) character(w http.ResponseWriter, r *http.Request) {
	region, ruleset := strings.ToLower(r.PathValue("region")), strings.ToLower(r.PathValue("ruleset"))
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	c, ok, err := s.Store.Character(r.Context(), region, ruleset, character.Slug(r.PathValue("name")))
	if err != nil {
		s.fail(w, r, "character", err, "could not read that character just now")
		return
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	cache(w)
	httpx.WriteOK(w, r, http.StatusOK, c)
}

func (s *Service) guild(w http.ResponseWriter, r *http.Request) {
	region, ruleset := strings.ToLower(r.PathValue("region")), strings.ToLower(r.PathValue("ruleset"))
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	g, ok, err := s.Store.Guild(r.Context(), region, ruleset, strings.ReplaceAll(r.PathValue("name"), "-", " "))
	if err != nil {
		s.fail(w, r, "guild", err, "could not read that guild just now")
		return
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	cache(w)
	httpx.WriteOK(w, r, http.StatusOK, g)
}
