package rankings

import (
	"context"
	"net/http/httptest"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/metrics"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/spec"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

type harness struct {
	t       *testing.T
	store   *Store
	reports *reports.Store
	pool    *pgxpool.Pool
	server  *httptest.Server
	guildID int64
	owner   int64
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	ctx := context.Background()
	if _, err := pool.Exec(ctx,
		`truncate users, reports, guilds, fight_metrics, percentile_digests, moderation cascade`); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureMetricsPartitions(ctx, pool, engine.FixtureBase, 1); err != nil {
		t.Fatal(err)
	}
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	build, _ := data.Latest()

	h := &harness{t: t, pool: pool, reports: &reports.Store{Pool: pool},
		store: &Store{Pool: pool, Specs: spec.New(build)}}
	owner, err := (&auth.Store{Pool: pool}).UpsertEmailUser(ctx, "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	h.owner = owner.ID
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Forever Sixty') returning id`).
		Scan(&h.guildID); err != nil {
		t.Fatal(err)
	}
	return h
}

// seedReport makes a report owned by the guild.
func (h *harness) seedReport(id string) {
	h.t.Helper()
	character := "us/hardcore/baelgrim"
	if _, err := h.reports.Create(context.Background(), reports.Report{
		ID: id, OwnerID: &h.owner, GuildID: &h.guildID, Title: "Tuesday",
		Visibility: reports.Public, Status: reports.StatusComplete, LoggingCharacter: &character,
	}); err != nil {
		h.t.Fatal(err)
	}
}

// seedFight writes one fixture fight's ranking rows, bending them first.
func (h *harness) seedFight(reportID string, index int, at time.Time, bend func([]metrics.Row)) []metrics.Row {
	h.t.Helper()
	fx, err := engine.NewFixture(reportID)
	if err != nil {
		h.t.Fatal(err)
	}
	rows := metrics.Derive(fx.Fight, fx.Summary)
	if bend != nil {
		bend(rows)
	}
	if _, err := h.reports.UpsertFight(context.Background(), reports.FightRecord{
		ReportID: reportID, Index: index, Name: "Warden Kelthas", Kill: true,
		EncounterID: ptr(int64(9001)), Difficulty: ptr(int64(8)), Size: ptr(int64(5)),
		DurationMS: 30000, StartMS: at.UnixMilli(), Verified: true,
	}); err != nil {
		h.t.Fatal(err)
	}
	if err := h.store.WriteFight(context.Background(), reports.RankedFight{
		ReportID: reportID, FightIndex: index, FoughtAt: at, Region: "us", Ruleset: "hardcore",
		GuildID: &h.guildID, Rows: rows, Combatants: fx.Summary.Combatants,
		Factions: map[string]string{rows[0].PlayerGUID: "alliance"},
	}); err != nil {
		h.t.Fatal(err)
	}
	return rows
}

func ptr[T any](v T) *T { return &v }

func TestWriteFightStoresRowsAndADigest(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	rows := h.seedFight("report-one", 1, engine.FixtureBase, nil)

	var n int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from fight_metrics where report_id = 'report-one'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != len(rows) {
		t.Fatalf("rows = %d, want %d", n, len(rows))
	}
	var phaseName, key string
	if err := h.pool.QueryRow(t.Context(),
		`select phase, player_key from fight_metrics where report_id = 'report-one' order by player_key limit 1`).
		Scan(&phaseName, &key); err != nil {
		t.Fatal(err)
	}
	if phaseName != "raids-1" {
		t.Fatalf("phase = %q, want the December raids phase", phaseName)
	}
	if key != "us/hardcore/baelgrim" {
		t.Fatalf("player key = %q", key)
	}
	var digests int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from percentile_digests`).Scan(&digests); err != nil {
		t.Fatal(err)
	}
	if digests == 0 {
		t.Fatal("no digest was written")
	}
}

func TestWritingTheSameFightTwiceReplacesItsRows(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)
	h.seedFight("report-one", 1, engine.FixtureBase, nil)
	var n int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from fight_metrics where report_id = 'report-one'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("rows = %d, want the three the fight has", n)
	}
}

// A fight that arrives twice - a second verifying bundle for the same
// index, with a different raw range - must not land in the distribution
// twice, because a t-digest cannot have a value taken back out.
func TestWritingTheSameFightTwiceDoesNotCountItTwice(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	rows := h.seedFight("report-one", 1, engine.FixtureBase, nil)
	first := h.digested()
	if first != int64(len(rows)) {
		t.Fatalf("digests hold %d values after one write, want %d", first, len(rows))
	}
	h.seedFight("report-one", 1, engine.FixtureBase, nil)
	if got := h.digested(); got != first {
		t.Fatalf("digests hold %d values after a re-write, want them unchanged at %d", got, first)
	}
}

// digested is how many values the stored digests hold between them. The
// fixture's three players are a tank, a healer and a damage dealer, so
// they land in one bracket each.
func (h *harness) digested() int64 {
	h.t.Helper()
	var n int64
	if err := h.pool.QueryRow(context.Background(),
		`select coalesce(sum(n), 0) from percentile_digests`).Scan(&n); err != nil {
		h.t.Fatal(err)
	}
	return n
}

// A fight that was not a kill is still stored, but its numbers are not
// comparable with a kill's, so they stay out of the distribution.
func TestAWipeIsStoredButDoesNotFeedTheDigest(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, func(rows []metrics.Row) {
		for i := range rows {
			rows[i].Kill = false
		}
	})
	var n, digests int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from fight_metrics where report_id = 'report-one'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("rows = %d, want the three the fight has", n)
	}
	if err := h.pool.QueryRow(t.Context(), `select count(*) from percentile_digests`).Scan(&digests); err != nil {
		t.Fatal(err)
	}
	if digests != 0 {
		t.Fatalf("digests = %d, want none: a wipe does not rank", digests)
	}
}

func TestWriteFightWithNoRowsDoesNothing(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	if err := h.store.WriteFight(t.Context(), reports.RankedFight{
		ReportID: "report-one", FightIndex: 1, FoughtAt: engine.FixtureBase,
		Region: "us", Ruleset: "hardcore",
	}); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from fight_metrics`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("rows = %d, want none", n)
	}
}

// A deployment with no talent data still ranks: the rows keep the spec
// the engine named, and nothing panics on the nil inferrer.
func TestAStoreWithNoTalentDataStillWritesRows(t *testing.T) {
	h := newHarness(t)
	h.store.Specs = nil
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, func(rows []metrics.Row) {
		rows[0].Spec = "fury"
	})
	var got string
	if err := h.pool.QueryRow(t.Context(),
		`select coalesce(spec, '') from fight_metrics where report_id = 'report-one' order by player_key limit 1`).
		Scan(&got); err != nil {
		t.Fatal(err)
	}
	if got != "fury" {
		t.Fatalf("spec = %q, want the engine's own", got)
	}
}

func TestRemoveReportWithdrawsItsRowsAndRecordsWhy(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)
	if err := h.store.RemoveReport(t.Context(), "report-one", "the stored events do not match the raw log"); err != nil {
		t.Fatal(err)
	}
	var rows, moderations int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from fight_metrics where report_id = 'report-one'`).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	if rows != 0 {
		t.Fatalf("rows = %d, want none", rows)
	}
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from moderation where target_id = 'report-one' and state = 'removed'`).
		Scan(&moderations); err != nil {
		t.Fatal(err)
	}
	if moderations != 1 {
		t.Fatalf("moderation rows = %d, want one", moderations)
	}
	var reason string
	if err := h.pool.QueryRow(t.Context(),
		`select reason from moderation where target_id = 'report-one'`).Scan(&reason); err != nil {
		t.Fatal(err)
	}
	if reason != "the stored events do not match the raw log" {
		t.Fatalf("reason = %q, want the caller's own words", reason)
	}
	page, err := h.store.Rankings(t.Context(), Query{EncounterID: 9001}, engine.FixtureBase)
	if err != nil {
		t.Fatal(err)
	}
	if page.Total != 0 || len(page.Rows) != 0 {
		t.Fatalf("rankings = %+v, want a withdrawn report to be gone from the leaderboard", page)
	}
}

// Withdrawing a report that never ranked - every fight failed
// verification, say - is not an error: there is simply nothing to take
// back, and the withdrawal is still recorded.
func TestRemoveReportWithNoRowsIsHarmless(t *testing.T) {
	h := newHarness(t)
	if err := h.store.RemoveReport(t.Context(), "never-ranked", "the stored events do not match the raw log"); err != nil {
		t.Fatalf("removing a report with no rows: %v", err)
	}
	var moderations int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from moderation where target_id = 'never-ranked' and state = 'removed'`).
		Scan(&moderations); err != nil {
		t.Fatal(err)
	}
	if moderations != 1 {
		t.Fatalf("moderation rows = %d, want the withdrawal recorded once", moderations)
	}
}

func TestPercentilePlacesAValueAmongTheKills(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	var specName, phaseName string
	if err := h.pool.QueryRow(t.Context(),
		`select spec, phase from percentile_digests where metric = $1`, MetricDPS).
		Scan(&specName, &phaseName); err != nil {
		t.Fatal(err)
	}
	got, ok, err := h.store.Percentile(t.Context(), 9001, 8, specName, phaseName, MetricDPS, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("the bracket has a digest, so it must place a value")
	}
	if got != 0 {
		t.Fatalf("percentile = %v, want the bottom: zero beats no parse", got)
	}
	if _, ok, err := h.store.Percentile(t.Context(), 4242, 8, specName, phaseName, MetricDPS, 1); err != nil || ok {
		t.Fatalf("an encounter with no digest = %v, %v", ok, err)
	}
}

func TestMetricHelpers(t *testing.T) {
	for _, m := range Metrics {
		if !ValidMetric(m) {
			t.Errorf("%q should be a metric", m)
		}
	}
	if ValidMetric("threat") {
		t.Error("unknown names must be refused")
	}
	if metricOf("healer") != MetricHPS || metricOf("tank") != MetricDamageTaken || metricOf("dps") != MetricDPS {
		t.Error("each role ranks on its own metric")
	}
	if got := valueOf(1, 2, 3, MetricHPS); got != 2 {
		t.Errorf("valueOf = %v", got)
	}
	if got := valueOf(1, 2, 3, MetricDamageTaken); got != 3 {
		t.Errorf("valueOf = %v", got)
	}
	if got := valueOf(1, 2, 3, MetricDPS); got != 1 {
		t.Errorf("valueOf = %v", got)
	}
}

func TestSinceReadsTheAllTimeSwitch(t *testing.T) {
	now := time.Date(2026, 12, 9, 12, 0, 0, 0, time.UTC)
	if got, err := since("", now); err != nil || !got.IsZero() {
		t.Fatalf("empty = %v, %v", got, err)
	}
	got, err := since("today", now)
	if err != nil || got.Hour() != 0 {
		t.Fatalf("today = %v, %v", got, err)
	}
	if got, err := since("7d", now); err != nil || !got.Equal(now.AddDate(0, 0, -7)) {
		t.Fatalf("7d = %v, %v", got, err)
	}
	if _, err := since("forever", now); err == nil {
		t.Fatal("nonsense must be refused")
	}
	if _, err := since("0d", now); err == nil {
		t.Fatal("a zero-day window must be refused")
	}
}

func TestFactionsFromEventsReadsCombatantInfo(t *testing.T) {
	fx, err := engine.NewFixture("r")
	if err != nil {
		t.Fatal(err)
	}
	// The fixture has no COMBATANT_INFO, so nothing is read from it.
	if got := FactionsFromEvents(fx.Events); len(got) != 0 {
		t.Fatalf("factions = %v", got)
	}
	got := FactionsFromEvents([]event.Event{
		{Kind: event.CombatantInfo, Combatant: &event.Combatant{GUID: "player-1", Faction: 1}},
		{Kind: event.CombatantInfo, Combatant: &event.Combatant{GUID: "player-2", Faction: 0}},
		{Kind: event.CombatantInfo, Combatant: &event.Combatant{GUID: "player-3", Faction: 7}},
		{Kind: event.CombatantInfo},
	})
	if len(got) != 2 || got["player-1"] != "alliance" || got["player-2"] != "horde" {
		t.Fatalf("factions = %v", got)
	}
	if got := factionName(0); got != "horde" {
		t.Errorf("faction 0 = %q", got)
	}
	if got := factionName(1); got != "alliance" {
		t.Errorf("faction 1 = %q", got)
	}
	if got := factionName(7); got != "" {
		t.Errorf("an unknown faction = %q", got)
	}
}

func TestTrinketsAndBuffCountReadTheCombatantRow(t *testing.T) {
	var c summary.CombatantRow
	if got := trinketsOf(c); len(got) != 0 {
		t.Fatalf("trinkets = %v", got)
	}
	if got := buffCountOf(c); got != 0 {
		t.Fatalf("buff count = %d", got)
	}
	c.Gear = make([]event.Item, 14)
	c.Gear[12] = event.Item{ID: 19406}
	c.Gear[13] = event.Item{ID: 18820}
	got := trinketsOf(c)
	if len(got) != 2 || got[0] != 19406 || got[1] != 18820 {
		t.Fatalf("trinkets = %v, want both trinket slots", got)
	}
	c.RaidBuffs = []summary.AuraRef{{}, {}}
	c.Consumables = []summary.AuraRef{{}}
	if got := buffCountOf(c); got != 3 {
		t.Fatalf("buff count = %d, want three", got)
	}
}

// A fight that carried COMBATANT_INFO stores what only that line
// knows: the two trinkets, the buff count, and the faction.
func TestARowCarriesItsTrinketsBuffsAndFaction(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	fx, err := engine.NewFixture("report-one")
	if err != nil {
		t.Fatal(err)
	}
	rows := metrics.Derive(fx.Fight, fx.Summary)
	gear := make([]event.Item, 14)
	gear[12] = event.Item{ID: 19406}
	gear[13] = event.Item{ID: 18820}
	if err := h.store.WriteFight(t.Context(), reports.RankedFight{
		ReportID: "report-one", FightIndex: 1, FoughtAt: engine.FixtureBase,
		Region: "us", Ruleset: "hardcore", GuildID: &h.guildID, Rows: rows,
		Combatants: []summary.CombatantRow{{
			GUID: rows[0].PlayerGUID, Gear: gear,
			RaidBuffs: []summary.AuraRef{{}, {}}, Consumables: []summary.AuraRef{{}},
		}},
		Factions: map[string]string{rows[0].PlayerGUID: "alliance"},
	}); err != nil {
		t.Fatal(err)
	}
	var trinkets []int64
	var buffs int
	var faction string
	if err := h.pool.QueryRow(t.Context(),
		`select trinkets, buff_count, coalesce(faction, '') from fight_metrics
		 where report_id = 'report-one' and player_key = $1`,
		character.KeyFromUnit("us", "hardcore", rows[0].Name)).Scan(&trinkets, &buffs, &faction); err != nil {
		t.Fatal(err)
	}
	if len(trinkets) != 2 || trinkets[0] != 19406 || trinkets[1] != 18820 {
		t.Errorf("trinkets = %v", trinkets)
	}
	if buffs != 3 {
		t.Errorf("buff count = %d, want three", buffs)
	}
	if faction != "alliance" {
		t.Errorf("faction = %q", faction)
	}
}

func TestEveryCallFailsWhenTheDatabaseIsGone(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	fx, err := engine.NewFixture("report-one")
	if err != nil {
		t.Fatal(err)
	}
	rows := metrics.Derive(fx.Fight, fx.Summary)
	h.pool.Close()

	if err := h.store.WriteFight(t.Context(), reports.RankedFight{
		ReportID: "report-one", FightIndex: 1, FoughtAt: engine.FixtureBase,
		Region: "us", Ruleset: "hardcore", Rows: rows,
	}); err == nil {
		t.Error("writing a fight must fail when the database is gone")
	}
	if err := h.store.RemoveReport(t.Context(), "report-one", "the stored events do not match the raw log"); err == nil {
		t.Error("withdrawing a report must fail when the database is gone")
	}
	if _, _, err := h.store.Percentile(t.Context(), 9001, 8, "", "raids-1", MetricDPS, 1); err == nil {
		t.Error("reading a percentile must fail when the database is gone")
	}
}

// The second fight in a bracket merges into the digest that is already
// there rather than replacing it.
func TestASecondFightFoldsIntoTheSameBracket(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	rows := h.seedFight("report-one", 1, engine.FixtureBase, nil)
	h.seedReport("report-two")
	h.seedFight("report-two", 1, engine.FixtureBase, nil)

	if got, want := h.digested(), int64(2*len(rows)); got != want {
		t.Fatalf("digests hold %d values, want %d: the second fight must merge in", got, want)
	}
	var brackets int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from percentile_digests`).Scan(&brackets); err != nil {
		t.Fatal(err)
	}
	if brackets != 3 {
		t.Fatalf("brackets = %d, want one per role", brackets)
	}
}

// Two writers of the same fight must not both call themselves the
// first write, so a fight is never folded into the digests twice. The
// lock that guarantees it is held here by another connection.
func TestWriteFightWaitsForAnotherWriterOfTheSameFight(t *testing.T) {
	h := newHarness(t)
	h.seedReport("report-one")
	fx, err := engine.NewFixture("report-one")
	if err != nil {
		t.Fatal(err)
	}
	fight := reports.RankedFight{
		ReportID: "report-one", FightIndex: 1, FoughtAt: engine.FixtureBase,
		Region: "us", Ruleset: "hardcore", Rows: metrics.Derive(fx.Fight, fx.Summary),
	}

	holder, err := h.pool.Acquire(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer holder.Release()
	if _, err := holder.Exec(context.Background(),
		`select pg_advisory_lock(hashtext($1), $2)`, fight.ReportID, fight.FightIndex); err != nil {
		t.Fatal(err)
	}

	waiting, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	if err := h.store.WriteFight(waiting, fight); err == nil {
		t.Fatal("a second writer of the same fight must wait for the first")
	}
	var n int
	if err := h.pool.QueryRow(context.Background(), `select count(*) from fight_metrics`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("rows = %d, want none: the blocked write must roll back", n)
	}

	if _, err := holder.Exec(context.Background(),
		`select pg_advisory_unlock(hashtext($1), $2)`, fight.ReportID, fight.FightIndex); err != nil {
		t.Fatal(err)
	}
	if err := h.store.WriteFight(context.Background(), fight); err != nil {
		t.Fatalf("once the lock is free the write goes through: %v", err)
	}
	if got := h.digested(); got != int64(len(fight.Rows)) {
		t.Fatalf("digests hold %d values, want %d", got, len(fight.Rows))
	}
}

func TestSortedKeysLockTheBracketsInAFixedOrder(t *testing.T) {
	pending := map[digestKey][]float64{
		{MetricHPS, "holy"}:         nil,
		{MetricDPS, "fury"}:         nil,
		{MetricDamageTaken, "prot"}: nil,
		{MetricDPS, "arms"}:         nil,
	}
	want := []digestKey{
		{MetricDamageTaken, "prot"}, {MetricDPS, "arms"}, {MetricDPS, "fury"}, {MetricHPS, "holy"},
	}
	for range 8 {
		if got := sortedKeys(pending); !slices.Equal(got, want) {
			t.Fatalf("sortedKeys = %v, want %v", got, want)
		}
	}
}
