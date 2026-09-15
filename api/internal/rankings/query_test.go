package rankings

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/metrics"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// seedReportVisibility makes a report owned by the guild with the given
// visibility, so a test can check what a non-public report does and
// does not leak into a guild's public reads.
func (h *harness) seedReportVisibility(id, visibility string) {
	h.t.Helper()
	character := "us/hardcore/baelgrim"
	if _, err := h.reports.Create(context.Background(), reports.Report{
		ID: id, OwnerID: &h.owner, GuildID: &h.guildID, Title: "Tuesday",
		Visibility: visibility, Status: reports.StatusComplete, LoggingCharacter: &character,
	}); err != nil {
		h.t.Fatal(err)
	}
}

// seedFightRecord writes only the fights-table row for one fight,
// mirroring what the ingest writes before the ranking gate
// (reports.Ranked) runs. It is used to test the guild progression
// query's own visibility filter independently of whether the fight
// ever produced ranking rows.
func (h *harness) seedFightRecord(reportID string, index int, at time.Time, kill bool) {
	h.t.Helper()
	if _, err := h.reports.UpsertFight(context.Background(), reports.FightRecord{
		ReportID: reportID, Index: index, Name: "Warden Kelthas", Kill: kill,
		EncounterID: ptr(int64(9001)), Difficulty: ptr(int64(8)), Size: ptr(int64(5)),
		DurationMS: 30000, StartMS: at.UnixMilli(), Verified: true,
	}); err != nil {
		h.t.Fatal(err)
	}
}

// seedTiedRankingRows writes n fight_metrics rows directly, every one
// tied on both the ranked metric and fought_at - the exact condition
// under which an ORDER BY with no unique final column can repeat or
// drop rows across a page boundary. Writing straight to the table
// rather than through WriteFight keeps a 100+ row fixture cheap.
func (h *harness) seedTiedRankingRows(n int, value float64, at time.Time) {
	h.t.Helper()
	for i := 0; i < n; i++ {
		reportID := fmt.Sprintf("tie-report-%04d", i)
		key := fmt.Sprintf("us/hardcore/tie-%04d", i)
		if _, err := h.pool.Exec(context.Background(),
			`insert into fight_metrics (report_id, fight_index, player_key, player_name, role,
			   metric_dps, encounter_id, difficulty, kill, phase, fought_at, state)
			 values ($1, 1, $2, $2, 'dps', $3, 9001, 8, true, 'raids-1', $4, 'ok')`,
			reportID, key, value, at); err != nil {
			h.t.Fatal(err)
		}
	}
}

// serve mounts the read routes over a harness, with a fixed clock so
// "today" means the fixture's day.
func serve(t *testing.T, h *harness) {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, &Service{Store: h.store, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
		Now: func() time.Time { return engine.FixtureBase }})
	h.server = httptest.NewServer(mux)
	t.Cleanup(h.server.Close)
}

func (h *harness) get(path string) *http.Response {
	h.t.Helper()
	res, err := http.Get(h.server.URL + path)
	if err != nil {
		h.t.Fatal(err)
	}
	return res
}

func (h *harness) data(res *http.Response, into any) {
	h.t.Helper()
	defer res.Body.Close()
	var env struct {
		OK   bool            `json:"ok"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		h.t.Fatal(err)
	}
	if !env.OK {
		h.t.Fatalf("envelope reports failure for %s", res.Request.URL)
	}
	if into != nil {
		if err := json.Unmarshal(env.Data, into); err != nil {
			h.t.Fatal(err)
		}
	}
}

func fmtFloat(v float64) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestRankingsAnswerAPageWithGuildsAndRanks(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	res := h.get("/v1/rankings?encounter=9001&metric=dps")
	if got := res.Header.Get("Cache-Control"); got != "public, max-age=30" {
		t.Fatalf("cache-control = %q", got)
	}
	var page Page
	h.data(res, &page)
	if page.PerPage != 100 || page.Page != 1 || page.Total != 3 {
		t.Fatalf("page = %+v", page)
	}
	if len(page.Rows) != 3 {
		t.Fatalf("rows = %d", len(page.Rows))
	}
	first := page.Rows[0]
	if first.Rank != 1 || first.Player.Key == "" || first.ReportID != "report-one" {
		t.Fatalf("row = %+v", first)
	}
	if first.Guild == nil || first.Guild.Name != "Forever Sixty" || first.Guild.Ruleset != "hardcore" {
		t.Fatalf("guild = %+v", first.Guild)
	}
	if first.State != StateOK || first.Trinkets == nil {
		t.Fatalf("row = %+v", first)
	}
	for i := 1; i < len(page.Rows); i++ {
		if page.Rows[i-1].Value < page.Rows[i].Value {
			t.Fatalf("rows are not ordered by value: %+v", page.Rows)
		}
	}
}

func TestRankingsFilterOnEveryAxis(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	for name, query := range map[string]string{
		"phase":   "&phase=raids-1",
		"region":  "&region=us",
		"ruleset": "&ruleset=hardcore",
		"since":   "&since=30d",
	} {
		var page Page
		h.data(h.get("/v1/rankings?encounter=9001"+query), &page)
		if page.Total == 0 {
			t.Errorf("filter %s matched nothing", name)
		}
	}
	for name, query := range map[string]string{
		"another phase":   "&phase=launch",
		"another region":  "&region=eu",
		"another ruleset": "&ruleset=pvp",
		"faction":         "&faction=horde",
		"difficulty":      "&difficulty=3",
	} {
		var page Page
		h.data(h.get("/v1/rankings?encounter=9001"+query), &page)
		if page.Total != 0 {
			t.Errorf("filter %s matched %d rows, want none", name, page.Total)
		}
	}
	var page Page
	h.data(h.get("/v1/rankings?encounter=9001&faction=alliance"), &page)
	if page.Total != 1 {
		t.Fatalf("the faction the fight recorded matched %d rows, want 1", page.Total)
	}
}

func TestRankingsRejectBadParameters(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	for name, path := range map[string]string{
		"no encounter": "/v1/rankings",
		"bad metric":   "/v1/rankings?encounter=1&metric=threat",
		"bad phase":    "/v1/rankings?encounter=1&phase=season-of-mastery",
		"bad ruleset":  "/v1/rankings?encounter=1&ruleset=nightslayer",
		"bad region":   "/v1/rankings?encounter=1&region=mars",
		"bad page":     "/v1/rankings?encounter=1&page=0",
		"bad since":    "/v1/rankings?encounter=1&since=forever",
		"bad kind":     "/v1/rankings/guilds?kind=wipes",
		"no value":     "/v1/rankings/percentile?encounter=1&difficulty=8",
	} {
		res := h.get(path)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d, want 400", name, res.StatusCode)
		}
	}
}

func TestPercentilePlacesAParse(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	rows := h.seedFight("report-one", 1, engine.FixtureBase, nil)
	var top float64
	for _, r := range rows {
		if r.MetricDPS > top {
			top = r.MetricDPS
		}
	}

	res := h.get("/v1/rankings/percentile?encounter=9001&difficulty=8&metric=dps&phase=raids-1&value=" +
		fmtFloat(top*2))
	var out struct {
		Percentile float64 `json:"percentile"`
	}
	h.data(res, &out)
	if out.Percentile < 99 {
		t.Fatalf("a parse above everything is in the %.0fth percentile", out.Percentile)
	}

	res = h.get("/v1/rankings/percentile?encounter=4242&difficulty=8&value=1")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("an empty bracket = %d, want 404", res.StatusCode)
	}
}

func TestTheCharacterPageShowsBestsHistoryAndBuilds(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, func(rows []metrics.Row) {
		for i := range rows {
			rows[i].Class = "Warrior"
		}
	})

	res := h.get("/v1/characters/us/hardcore/baelgrim")
	var c Character
	h.data(res, &c)
	if c.Character.Key != "us/hardcore/baelgrim" {
		t.Fatalf("character = %+v", c.Character)
	}
	if len(c.History) != 1 || c.History[0].ReportID != "report-one" {
		t.Fatalf("history = %+v", c.History)
	}
	if len(c.Best) != 1 || c.Best[0].EncounterID != 9001 {
		t.Fatalf("best = %+v", c.Best)
	}

	res = h.get("/v1/characters/us/hardcore/nobody")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("an unknown character = %d, want 404", res.StatusCode)
	}
	res = h.get("/v1/characters/mars/hardcore/baelgrim")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("an unknown region = %d, want 404", res.StatusCode)
	}
}

func TestTheGuildPageShowsProgressionRosterAndReports(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	res := h.get("/v1/guilds/us/hardcore/forever-sixty")
	var g Guild
	h.data(res, &g)
	if g.Guild.Name != "Forever Sixty" {
		t.Fatalf("guild = %+v", g.Guild)
	}
	if len(g.Progression) != 1 || g.Progression[0].Kills != 1 || g.Progression[0].FirstKillAt == nil {
		t.Fatalf("progression = %+v", g.Progression)
	}
	if len(g.RosterBest) != 3 {
		t.Fatalf("roster = %+v", g.RosterBest)
	}
	if len(g.Reports) != 1 || g.Reports[0].ID != "report-one" {
		t.Fatalf("reports = %+v", g.Reports)
	}

	res = h.get("/v1/guilds/us/hardcore/nobody")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("an unknown guild = %d, want 404", res.StatusCode)
	}
	res = h.get("/v1/guilds/us/nightslayer/forever-sixty")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("a realm that is not a ruleset = %d, want 404", res.StatusCode)
	}
}

func TestGuildRankingsRankSpeedExecutionAndProgress(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	for _, kind := range []string{KindSpeed, KindExecution, KindProgress} {
		var out struct {
			Rows []GuildRow `json:"rows"`
		}
		h.data(h.get("/v1/rankings/guilds?encounter=9001&kind="+kind), &out)
		if len(out.Rows) != 1 {
			t.Fatalf("%s rows = %+v", kind, out.Rows)
		}
		if out.Rows[0].Guild.Name != "Forever Sixty" || out.Rows[0].Rank != 1 {
			t.Fatalf("%s row = %+v", kind, out.Rows[0])
		}
	}
	var out struct {
		Rows []GuildRow `json:"rows"`
	}
	h.data(h.get("/v1/rankings/guilds?kind=progress&phase=raids-1"), &out)
	if len(out.Rows) != 1 || out.Rows[0].Kills != 1 {
		t.Fatalf("progress rows = %+v", out.Rows)
	}
	res := h.get("/v1/rankings/guilds?kind=speed")
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("speed with no encounter = %d, want 400", res.StatusCode)
	}
}

func TestEveryReadAnswers500WhenTheDatabaseIsGone(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.pool.Close()
	for name, path := range map[string]string{
		"rankings":   "/v1/rankings?encounter=9001",
		"percentile": "/v1/rankings/percentile?encounter=9001&difficulty=8&value=1",
		"guilds":     "/v1/rankings/guilds?encounter=9001&kind=speed",
		"character":  "/v1/characters/us/hardcore/baelgrim",
		"guild":      "/v1/guilds/us/hardcore/forever-sixty",
	} {
		res := h.get(path)
		res.Body.Close()
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("%s = %d, want 500", name, res.StatusCode)
		}
	}
}

func TestTheServiceClockDefaultsToNow(t *testing.T) {
	s := &Service{}
	if s.now().IsZero() {
		t.Fatal("the default clock should be the real one")
	}
}

func TestRankingsAcceptAnEncounterSlugAsWellAsAnID(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	var byID, bySlug Page
	h.data(h.get("/v1/rankings?encounter=9001"), &byID)
	h.data(h.get("/v1/rankings?encounter=warden-kelthas"), &bySlug)
	if byID.Total == 0 || bySlug.Total != byID.Total {
		t.Fatalf("slug gave %d rows, id gave %d", bySlug.Total, byID.Total)
	}

	res := h.get("/v1/rankings?encounter=no-such-boss")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("an unknown slug = %d, want 404", res.StatusCode)
	}

	var guilds struct {
		Rows []GuildRow `json:"rows"`
	}
	h.data(h.get("/v1/rankings/guilds?encounter=warden-kelthas&kind=speed"), &guilds)
	if len(guilds.Rows) != 1 {
		t.Fatalf("guild rows = %+v", guilds.Rows)
	}
	res = h.get("/v1/rankings/percentile?encounter=warden-kelthas&difficulty=8&value=1")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("percentile by slug = %d", res.StatusCode)
	}
}

func TestTheCharacterPageNamesEncountersAndPlacesParses(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	var c Character
	h.data(h.get("/v1/characters/us/hardcore/baelgrim"), &c)
	if len(c.Best) != 1 || c.Best[0].Encounter != "Warden Kelthas" {
		t.Fatalf("best = %+v, want the encounter named", c.Best)
	}
	if len(c.History) != 1 || c.History[0].Encounter != "Warden Kelthas" {
		t.Fatalf("history = %+v", c.History)
	}
	if c.History[0].Percentile == nil {
		t.Fatal("a fight in a bracket with a digest should carry a percentile")
	}
	if *c.History[0].Percentile < 0 || *c.History[0].Percentile > 100 {
		t.Fatalf("percentile = %v", *c.History[0].Percentile)
	}
	if c.Best[0].Percentile == nil {
		t.Fatal("a best parse should carry a percentile too")
	}
}

func TestTheGuildPageNamesEncounters(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	var g Guild
	h.data(h.get("/v1/guilds/us/hardcore/forever-sixty"), &g)
	if len(g.Progression) != 1 || g.Progression[0].Encounter != "Warden Kelthas" {
		t.Fatalf("progression = %+v", g.Progression)
	}
	if len(g.RosterBest) == 0 || g.RosterBest[0].Encounter != "Warden Kelthas" {
		t.Fatalf("roster = %+v", g.RosterBest)
	}
	if g.RosterBest[0].Player.Key == "" {
		t.Fatalf("roster rows keep the nested player: %+v", g.RosterBest[0])
	}
}

func TestEncountersAreListedWithTheirSlugs(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	all, err := h.store.Encounters(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 1 || all[0].ID != 9001 || all[0].Slug != "warden-kelthas" {
		t.Fatalf("encounters = %+v", all)
	}
	if _, err := h.store.ResolveEncounter(t.Context(), "0"); !errors.Is(err, ErrNoEncounter) {
		t.Fatalf("a zero id = %v, want ErrNoEncounter", err)
	}
	if _, err := h.store.ResolveEncounter(t.Context(), " "); !errors.Is(err, ErrNoEncounter) {
		t.Fatalf("an empty slug = %v, want ErrNoEncounter", err)
	}
	if id, err := h.store.ResolveEncounter(t.Context(), "4242"); err != nil || id != 4242 {
		t.Fatalf("a numeric id should pass straight through: %d, %v", id, err)
	}
	names, err := h.store.EncounterNames(t.Context(), nil)
	if err != nil || len(names) != 0 {
		t.Fatalf("no ids = %v, %v", names, err)
	}
}

// A guild's progression must not reveal what an officer marked private:
// fights are written to the fights table unconditionally at ingest,
// before the ranking gate, so the progression query needs its own
// visibility filter rather than relying on fight_metrics being empty.
func TestGuildProgressionIgnoresPrivateReports(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReportVisibility("report-private", reports.Private)
	h.seedFightRecord("report-private", 1, engine.FixtureBase, true)
	h.seedReport("report-public")
	h.seedFightRecord("report-public", 1, engine.FixtureBase, true)

	var g Guild
	h.data(h.get("/v1/guilds/us/hardcore/forever-sixty"), &g)
	if len(g.Progression) != 1 {
		t.Fatalf("progression = %+v, want the one bracket both fights share", g.Progression)
	}
	if g.Progression[0].Kills != 1 || g.Progression[0].PullCount != 1 {
		t.Fatalf("progression = %+v, want only the public report's kill and pull counted",
			g.Progression[0])
	}
}

// A leaderboard page must not repeat or drop rows across a page
// boundary that falls inside a tie on both the ranked metric and
// fought_at - the exact pattern flagged as binding: ORDER BY needs a
// unique final column.
func TestRankingsPageBoundaryIsStableAcrossATie(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedTiedRankingRows(PerPage+5, 1000, engine.FixtureBase)

	var page1, page1Again, page2 Page
	h.data(h.get("/v1/rankings?encounter=9001&page=1"), &page1)
	h.data(h.get("/v1/rankings?encounter=9001&page=1"), &page1Again)
	h.data(h.get("/v1/rankings?encounter=9001&page=2"), &page2)

	if len(page1.Rows) != PerPage || len(page2.Rows) != 5 {
		t.Fatalf("page1 = %d rows, page2 = %d rows, want %d and 5", len(page1.Rows), len(page2.Rows), PerPage)
	}
	for i := range page1.Rows {
		if page1.Rows[i].Player.Key != page1Again.Rows[i].Player.Key {
			t.Fatalf("page 1 is not stable across repeated calls at row %d: %q vs %q",
				i, page1.Rows[i].Player.Key, page1Again.Rows[i].Player.Key)
		}
	}
	seen := map[string]bool{}
	for _, r := range page1.Rows {
		seen[r.Player.Key] = true
	}
	for _, r := range page2.Rows {
		if seen[r.Player.Key] {
			t.Fatalf("player %s appears on both page 1 and page 2", r.Player.Key)
		}
	}
}

// The guild page's report list is report identities, not aggregates: a
// report marked "guild" is readable by the guild, and listing its id,
// title, zone and timestamp to an unauthenticated caller tells the
// world it exists and what it is called. reports.mayView then refuses
// the body, which is the giveaway that the list was wrong.
func TestTheGuildPageListsOnlyPublicAndUnlistedReports(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReportVisibility("report-guild", reports.GuildTo)
	h.seedReportVisibility("report-private", reports.Private)
	h.seedReportVisibility("report-unlisted", reports.Unlisted)
	h.seedReport("report-public")

	var g Guild
	h.data(h.get("/v1/guilds/us/hardcore/forever-sixty"), &g)
	listed := map[string]bool{}
	for _, r := range g.Reports {
		listed[r.ID] = true
	}
	for _, id := range []string{"report-guild", "report-private"} {
		if listed[id] {
			t.Fatalf("%s is listed to a stranger: %+v", id, g.Reports)
		}
	}
	for _, id := range []string{"report-public", "report-unlisted"} {
		if !listed[id] {
			t.Fatalf("%s should be listed: %+v", id, g.Reports)
		}
	}
}

// The guild route reconstructs a guild's name by mapping each hyphen
// in the path back to one space, so the slug a URL is built from must
// spend one hyphen per space. character.Slug does, including for a
// doubled space, and this pins that: collapsing a run of spaces to a
// single hyphen would make a guild whose name carries one unreachable.
func TestAGuildNameWithADoubleSpaceRoundTripsThroughTheRoute(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	if _, err := h.pool.Exec(t.Context(),
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', $1)`,
		"Lady  Sunwick's Own"); err != nil {
		t.Fatal(err)
	}

	slug := character.Slug("Lady  Sunwick's Own")
	if slug != "lady--sunwick's-own" {
		t.Fatalf("slug = %q, want one hyphen per space", slug)
	}
	var g Guild
	h.data(h.get("/v1/guilds/us/hardcore/"+url.PathEscape(slug)), &g)
	if g.Guild.Name != "Lady  Sunwick's Own" {
		t.Fatalf("guild = %+v, want the doubled-space name resolved", g.Guild)
	}
}
