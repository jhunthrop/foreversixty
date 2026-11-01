package rankings

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
)

// cacheSeconds is how long a rankings answer may be cached at the edge.
// The design's caching note sets it at thirty seconds.
const cacheSeconds = 30

// liveBoardStale is the Live public boards class's
// stale-while-revalidate window (spec §2.2).
const liveBoardStale = 300 * time.Second

// Service serves the rankings, character, and guild routes.
type Service struct {
	Store *Store
	Log   *slog.Logger
	// Now is the clock, so a test can fix "today".
	Now func() time.Time
	// Scores adds the Forever Score to the character page; nil, or a failed read, leaves it
	// null.
	Scores Scorer
	// ScoreGear adds the Forever Score and the gear panel to the character page in one read,
	// in place of Scores when it is set; nil leaves the panel [] and its total null, and a
	// failed read leaves the score null.
	ScoreGear ScoreGearReader
	// Supporters adds the supporter mark to the character page; nil means nobody shows it.
	Supporters Supporters
}

// Supporters answers whether an account shows the supporter mark; nil means
// nobody does (the fixture build, and tests that do not care).
type Supporters interface {
	IsSupporter(ctx context.Context, userID int64) (bool, error)
}

// supporterOf reads the mark for a character's owning account. A cosmetic flag never takes
// a page down: no service, no owner or a failed read is false (the failure logged at warn).
func (s *Service) supporterOf(ctx context.Context, ownerID *int64) bool {
	if s.Supporters == nil || ownerID == nil {
		return false
	}
	ok, err := s.Supporters.IsSupporter(ctx, *ownerID)
	if err != nil {
		s.logger().Warn("rankings", "id", httpx.RequestIDFrom(ctx), "op", "supporter", "err", err)
		return false
	}
	return ok
}

// Scorer reads a character's Forever Score by character key. rating.Service is the one
// implementation: the formula lives there, and this package only carries the result. A
// character with no score is the zero ForeverScore, not an error.
type Scorer interface {
	ForeverScore(ctx context.Context, key string) (ForeverScore, error)
}

// Mount registers the read routes.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET /v1/rankings", s.rankings)
	mux.HandleFunc("GET /v1/rankings/percentile", s.percentile)
	mux.HandleFunc("GET /v1/rankings/guilds", s.guildRankings)
	mux.HandleFunc("GET /v1/rankings/encounters", s.rankedEncounters)
	mux.HandleFunc("GET /v1/me/watching", auth.RequireSession(s.watching))
	mux.HandleFunc("PUT /v1/me/watching/{region}/{ruleset}/{name}", auth.RequireSession(s.watch))
	mux.HandleFunc("DELETE /v1/me/watching/{region}/{ruleset}/{name}", auth.RequireSession(s.unwatch))
	mux.HandleFunc("GET /v1/encounters", s.encounters)
	mux.HandleFunc("GET /v1/characters/search", s.searchCharacters)
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

// CacheLive marks a read cacheable for the edge: rankings, ratings, the
// reports feed, and a guild or character's public page are all the
// spec's "Live public boards" class (§2.2). The apps' character home
// answer is cached the same as the character page it is read from.
func CacheLive(w http.ResponseWriter) {
	httpx.CachePublic(w, cacheSeconds*time.Second, liveBoardStale)
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
	if !s.readCharacterBoard(w, r, &query) {
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
	CacheLive(w)
	httpx.WriteOK(w, r, http.StatusOK, page)
}

// readCharacterBoard reads `distinct` and `character`, writing the 400
// itself when they are wrong. A character can only be placed on a board
// that ranks characters.
func (s *Service) readCharacterBoard(w http.ResponseWriter, r *http.Request, query *Query) bool {
	q := r.URL.Query()
	if v := q.Get("distinct"); v != "" {
		if v != DistinctCharacter {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a way to dedupe",
				map[string]string{"distinct": DistinctCharacter})
			return false
		}
		query.Distinct = v
	}
	if v := strings.ToLower(strings.TrimSpace(q.Get("character"))); v != "" {
		if query.Distinct != DistinctCharacter || len(strings.SplitN(v, "/", 3)) != 3 {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
				"character needs distinct=character and a region/ruleset/name key",
				map[string]string{"character": "region/ruleset/name, with distinct=character"})
			return false
		}
		query.Character = v
	}
	return true
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
	CacheLive(w)
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
	CacheLive(w)
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
	CacheLive(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": rows})
}

func (s *Service) character(w http.ResponseWriter, r *http.Request) {
	c, ok, err := s.CharacterCard(r.Context(), r.PathValue("region"), r.PathValue("ruleset"), r.PathValue("name"))
	if err != nil {
		s.fail(w, r, "character", err, "could not read that character just now")
		return
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	CacheLive(w)
	httpx.WriteOK(w, r, http.StatusOK, c)
}

// CharacterCard is the public character page (GET /v1/characters/{region}/{ruleset}/{name})
// without HTTP: found is false for an unknown region or ruleset, an unknown character and an
// anonymised one. A score or gear panel that cannot be read is logged and left empty.
func (s *Service) CharacterCard(ctx context.Context, region, ruleset, name string) (Character, bool, error) {
	region, ruleset = strings.ToLower(region), strings.ToLower(ruleset)
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
		return Character{}, false, nil
	}
	c, ok, err := s.Store.Character(ctx, region, ruleset, character.Slug(name))
	if err != nil || !ok {
		return Character{}, ok, err
	}
	s.addScoreAndGear(ctx, &c)
	c.Supporter = s.supporterOf(ctx, c.ownerID)
	return c, true, nil
}

// addScoreAndGear fills the Forever Score and the gear panel; a failed read leaves them null.
func (s *Service) addScoreAndGear(ctx context.Context, c *Character) {
	switch {
	case s.ScoreGear != nil:
		score, gear, err := s.ScoreGear.ScoreAndGear(ctx, c.Character.Key)
		if err != nil {
			s.logger().Error("rankings", "id", httpx.RequestIDFrom(ctx), "op", "forever_score",
				"character", c.Character.Key, "err", err)
		} else {
			c.ForeverScore = score
		}
		if gear.Slots != nil {
			c.Gear, c.GearGapTotal = gear.Slots, gear.GapTotal
			if gear.Spec != "" {
				spec, source := gear.Spec, gear.SpecSource
				c.GearSpec, c.GearSpecSource = &spec, &source
			}
		}
	case s.Scores != nil:
		score, err := s.Scores.ForeverScore(ctx, c.Character.Key)
		if err != nil {
			s.logger().Error("rankings", "id", httpx.RequestIDFrom(ctx), "op", "forever_score",
				"character", c.Character.Key, "err", err)
		} else {
			c.ForeverScore = score
		}
	}
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
	CacheLive(w)
	httpx.WriteOK(w, r, http.StatusOK, g)
}
