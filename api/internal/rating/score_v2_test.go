package rating

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/digest"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

func TestTierNamesEveryBandAtItsThresholds(t *testing.T) {
	for _, c := range []struct {
		score int
		want  string
	}{
		{0, TierCommon}, {499, TierCommon},
		{500, TierUncommon}, {999, TierUncommon},
		{1000, TierRare}, {1499, TierRare},
		{1500, TierEpic}, {1830, "epic"}, {1999, TierEpic},
		{2000, TierLegendary}, {2600, TierLegendary},
	} {
		if got := Tier(c.score); got != c.want {
			t.Errorf("Tier(%d) = %q, want %q", c.score, got, c.want)
		}
	}
}

// The component ceilings are the v1 definition; their sum is the most a score can be.
func TestComponentCeilingsAreTheV1Definition(t *testing.T) {
	if GearMax != 500 || ExperienceMax != 700 || ParsesMax != 600 || AttendanceMax != 200 || DisciplineMax != 100 {
		t.Fatalf("ceilings = %d/%d/%d/%d/%d, want 500/700/600/200/100",
			GearMax, ExperienceMax, ParsesMax, AttendanceMax, DisciplineMax)
	}
	if ExperiencePerDungeon != 25 || ExperiencePerRaidBoss != 40 || AttendanceWindow != 90*24*time.Hour {
		t.Fatal("experience is 25 a dungeon and 40 a raid boss; attendance looks back 90 days")
	}
}

// --- per-fight discipline ---------------------------------------------------------------

// Slot indexes in COMBATANT_INFO's gear list (the client's inventory slot, less one).
const (
	slotChest    = 4
	slotFeet     = 7
	slotBack     = 14
	slotMainHand = 15
)

// preparedCombatant is a dps combatant with a flask, food, a weapon stone and every
// enchantable slot enchanted.
func preparedCombatant(guid string) summary.CombatantRow {
	cat, _ := consumables.Load().For("dps")
	gear := make([]event.Item, 19)
	for i := range gear {
		gear[i] = event.Item{ID: int64(10000 + i)}
	}
	for _, s := range []int{slotChest, slotFeet, slotBack, slotMainHand} {
		gear[s].Enchants = []int64{1900}
	}
	return summary.CombatantRow{
		GUID: guid, Gear: gear,
		Consumables: []summary.AuraRef{
			{SpellID: cat.Flask[0].SpellID}, {SpellID: cat.Food[0].SpellID}, {SpellID: cat.WeaponEnchant[0].SpellID},
		},
	}
}

func dropConsumable(c summary.CombatantRow, spellID int64) summary.CombatantRow {
	kept := []summary.AuraRef{}
	for _, a := range c.Consumables {
		if a.SpellID != spellID {
			kept = append(kept, a)
		}
	}
	c.Consumables = kept
	return c
}

func TestFightDisciplineNeedsFlaskFoodAndEnchants(t *testing.T) {
	dps, _ := consumables.Load().For("dps")
	tank, _ := consumables.Load().For("tank")
	row := summary.RosterRow{GUID: "g", Name: "Prepared", Class: "Warrior", Spec: "Arms", Role: "dps"}
	fightWith := func(c summary.CombatantRow) summary.Summary {
		return summary.Summary{Combatants: []summary.CombatantRow{c}, Roster: []summary.RosterRow{row}}
	}
	check := func(name string, c summary.CombatantRow, want bool) {
		t.Helper()
		if got := fightDiscipline(fightWith(c), row); got == nil || *got != want {
			t.Errorf("%s: %v, want %v", name, got, want)
		}
	}

	check("a fully prepared player", preparedCombatant("g"), true)
	check("no flask", dropConsumable(preparedCombatant("g"), dps.Flask[0].SpellID), false)
	check("no food", dropConsumable(preparedCombatant("g"), dps.Food[0].SpellID), false)
	unenchanted := preparedCombatant("g")
	unenchanted.Gear[slotBack].Enchants = nil
	check("an unenchanted cloak", unenchanted, false)
	empty := preparedCombatant("g")
	empty.Gear[slotFeet] = event.Item{}
	check("an empty slot is not a missing enchant", empty, true)
	// Weapon stones and oils are not required until a beta log shows how they appear.
	check("no weapon stone", dropConsumable(preparedCombatant("g"), dps.WeaponEnchant[0].SpellID), true)
	// A category is met by any role's consumable of it: a dps warrior on a tank's food.
	tankFood := dropConsumable(preparedCombatant("g"), dps.Food[0].SpellID)
	tankFood.Consumables = append(tankFood.Consumables, summary.AuraRef{SpellID: tank.Food[0].SpellID})
	if tank.Food[0].SpellID == dps.Food[0].SpellID {
		t.Fatal("the tank and dps foods must differ for this case to mean anything")
	}
	check("a dps warrior on a tank's food", tankFood, true)
	if got := fightDiscipline(summary.Summary{Roster: []summary.RosterRow{row}}, row); got != nil {
		t.Errorf("no combatant info: %v, want unknown (nil)", *got)
	}
}

// Every Classic class/spec pair the log engine can name (units.RetailSpecName) must
// resolve to a spec the build data can weigh and measure against a full set.
func TestEveryLoggedClassicSpecResolvesToADataSpec(t *testing.T) {
	useRepoGearData(t)
	classic := map[string]bool{"Druid": true, "Hunter": true, "Mage": true, "Paladin": true, "Priest": true,
		"Rogue": true, "Shaman": true, "Warlock": true, "Warrior": true}
	seen := 0
	for id, spec := range units.RetailSpecName {
		class := units.RetailSpecClass[id]
		if !classic[class] {
			continue
		}
		seen++
		slug := dataSpecSlug(class, spec)
		if _, err := sims.FullGearScore(context.Background(), "1.60.1.70291", slug); err != nil {
			t.Errorf("%s %s -> %q: %v", class, spec, slug, err)
		}
	}
	if seen != 28 {
		t.Errorf("%d Classic spec ids, want 28 (nine classes, three specs each, four for the druid)", seen)
	}
	if got := dataSpecSlug("Rogue", "Outlaw"); got != "rogue-combat" {
		t.Errorf("Outlaw -> %q, want rogue-combat", got)
	}
	if got := dataSpecSlug("Druid", "Guardian"); got != "druid-feral-bear" {
		t.Errorf("Guardian -> %q, want druid-feral-bear", got)
	}
}

// --- ScoreV2 against the database -------------------------------------------------------

var (
	repoGearOnce sync.Once
	repoGearErr  error
)

// useRepoGearData wires sims.GearScore to the repository's own data/builds, once per test
// binary (loading every build's item tables takes a few seconds).
func useRepoGearData(t *testing.T) {
	t.Helper()
	const dir = "../../../data/builds"
	repoGearOnce.Do(func() {
		var data *trees.Data
		if data, repoGearErr = trees.Load(dir); repoGearErr == nil {
			sims.SetGearData(data, dir)
		}
	})
	if repoGearErr != nil {
		t.Fatal(repoGearErr)
	}
}

// fixtureGear is docs/contracts/fixtures/gearscore.json's code and its pinned score.
func fixtureGear(t *testing.T) (code string, score float64) {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "docs", "contracts", "fixtures", "gearscore.json"))
	if err != nil {
		t.Fatal(err)
	}
	var f struct {
		GearCode string  `json:"gear_code"`
		Expected float64 `json:"expected_score"`
	}
	if err := json.Unmarshal(b, &f); err != nil {
		t.Fatal(err)
	}
	return f.GearCode, f.Expected
}

type scoreFixture struct {
	pool               *pgxpool.Pool
	store              *Store
	userID, guildID    int64
	main, altA, altB   string
	newest             time.Time
	parsePercentile    float64
	gearScore          float64
	rated, missedFlask int
}

const (
	sv2Encounter = 990601 // 990601..990606: this file's own encounter ids
	sv2Report    = "sv2-report-"
)

func mustExec(t *testing.T, pool *pgxpool.Pool, sql string, args ...any) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), sql, args...); err != nil {
		t.Fatalf("%s: %v", sql, err)
	}
}

// seedLoggedRaider builds the brief's character: twelve rated fights over three guild
// nights (one without a flask), a 71st-percentile best parse, full attendance, an addon
// export to score gear from, and two linked alts with kills of their own.
func seedLoggedRaider(t *testing.T) scoreFixture {
	t.Helper()
	pool := testPool(t)
	ctx := context.Background()
	useRepoGearData(t)
	code, gearScore := fixtureGear(t)
	fx := scoreFixture{
		pool: pool, store: &Store{Pool: pool},
		main: "us/normal/sv2main", altA: "us/normal/sv2alta", altB: "us/normal/sv2altb",
		gearScore: gearScore,
	}
	t.Cleanup(func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'sv2-%'`)
		pool.Exec(bg, `delete from fight_metrics where report_id like 'sv2-%'`)
		pool.Exec(bg, `delete from reports where id like 'sv2-%'`)
		pool.Exec(bg, `delete from percentile_digests where encounter_id between 990601 and 990606`)
		pool.Exec(bg, `delete from rating_percentile_digests where encounter_id between 990601 and 990606`)
		pool.Exec(bg, `delete from kill_duration_digests where encounter_id between 990601 and 990606`)
		pool.Exec(bg, `delete from addon_exports where character_key like 'us/normal/sv2%'`)
		pool.Exec(bg, `delete from guild_characters where character_key like 'us/normal/sv2%'`)
		pool.Exec(bg, `delete from guilds where name = 'SV2 Guild'`)
		pool.Exec(bg, `delete from characters where key like 'us/normal/sv2%'`)
		pool.Exec(bg, `delete from users where email = 'scorev2-test@example.com'`)
	})

	if err := pool.QueryRow(ctx, `insert into users (email) values ('scorev2-test@example.com')
		on conflict (email) do update set email = excluded.email returning id`).Scan(&fx.userID); err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ key, name string }{
		{fx.main, "Sv2main"}, {fx.altA, "Sv2alta"}, {fx.altB, "Sv2altb"},
	} {
		mustExec(t, pool, `insert into characters (key, region, ruleset, name, class, user_id)
			values ($1, 'us', 'normal', $2, 'Warrior', $3)
			on conflict (key) do update set user_id = excluded.user_id`, c.key, c.name, fx.userID)
	}
	mustExec(t, pool, `update characters set equipped_item_level = 58 where key = $1`, fx.main)
	mustExec(t, pool, `update users set main_character_key = $1 where id = $2`, fx.main, fx.userID)
	mustExec(t, pool, `insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at)
		values ($1, $2, 'us', 'normal', 'Sv2main', $3, now())
		on conflict (character_key) do update set export = excluded.export`, fx.main, fx.userID, code)
	if err := pool.QueryRow(ctx, `insert into guilds (region, ruleset, name) values ('us', 'normal', 'SV2 Guild')
		on conflict (region, ruleset, name) do update set name = excluded.name returning id`).Scan(&fx.guildID); err != nil {
		t.Fatal(err)
	}
	mustExec(t, pool, `insert into guild_characters (guild_id, character_key, user_id, verified_at)
		values ($1, $2, $3, now()) on conflict (guild_id, character_key) do update set verified_at = now()`,
		fx.guildID, fx.main, fx.userID)

	// Three guild nights, ten days apart, four boss kills each.
	today := time.Now().UTC().Truncate(24 * time.Hour)
	row := summary.RosterRow{GUID: "sv2-g", Name: "Sv2main", Class: "Warrior", Spec: "Arms", Role: "dps",
		ExecutionScore: execScorePtr(0.8)}
	dps, _ := consumables.Load().For("dps")
	for night := 0; night < 3; night++ {
		reportID := sv2Report + string(rune('a'+night))
		at := today.Add(time.Duration(night-3)*10*24*time.Hour + 20*time.Hour)
		mustExec(t, pool, `insert into reports (id, owner_id, guild_id, title, visibility, zone, status, created_at)
			values ($1, $2, $3, 'Score v2', 'public', 'Molten Core', 'complete', $4)`, reportID, fx.userID, fx.guildID, at)
		for i := 1; i <= 4; i++ {
			combatant := preparedCombatant(row.GUID)
			if night == 0 && i == 1 {
				combatant = dropConsumable(combatant, dps.Flask[0].SpellID)
				fx.missedFlask++
			}
			foughtAt := at.Add(time.Duration(i) * 15 * time.Minute)
			mustExec(t, pool, `insert into fights (report_id, fight_index, encounter_id, name, difficulty, kill, start_ms)
				values ($1, $2, $3, 'Score v2 boss', 3, true, $4)`, reportID, i, sv2Encounter+i-1, foughtAt.UnixMilli())
			if err := fx.store.RateFight(ctx, RatedFight{
				ReportID: reportID, FightIndex: i, Region: "us", Ruleset: "normal",
				FoughtAt: foughtAt, EncounterID: int64(sv2Encounter + i - 1),
				Summary: summary.Summary{
					EncounterID: int64(sv2Encounter + i - 1), Difficulty: 3, DurationMS: 180000, Kill: true,
					Roster: []summary.RosterRow{row}, Combatants: []summary.CombatantRow{combatant},
				},
			}); err != nil {
				t.Fatal(err)
			}
			fx.rated++
			fx.newest = foughtAt
		}
	}

	// The best parse: one ranked kill placed against a curve of 1..100.
	d := digest.New()
	for v := 1; v <= 100; v++ {
		d.Add(float64(v))
	}
	raw, err := d.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	fx.parsePercentile = d.Placement(71) * 100
	mustExec(t, pool, `insert into percentile_digests (encounter_id, difficulty, spec, phase, metric, digest, n)
		values ($1, 3, 'Arms', '', 'dps', $2, 100)`, sv2Encounter, raw)
	metric := func(reportID string, fightIndex int, key, name string, encounter int64, dps float64) {
		at := today.Add(-25 * 24 * time.Hour)
		if err := db.EnsureMetricsPartition(ctx, pool, at); err != nil {
			t.Fatal(err)
		}
		mustExec(t, pool, `insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec,
			role, metric_dps, encounter_id, difficulty, kill, phase, fought_at)
			values ($1, $2, $3, $4, 'Warrior', 'Arms', 'dps', $5, $6, 3, true, '', $7)`,
			reportID, fightIndex, key, name, dps, encounter, at)
	}
	metric(sv2Report+"a", 1, fx.main, "Sv2main", sv2Encounter, 71)
	// The alts: one boss the main never killed each, and one the main did.
	metric(sv2Report+"a", 5, fx.altA, "Sv2alta", sv2Encounter+4, 50)
	metric(sv2Report+"a", 6, fx.altB, "Sv2altb", sv2Encounter+5, 50)
	metric(sv2Report+"a", 1, fx.altB, "Sv2altb", sv2Encounter, 50)
	return fx
}

func TestScoreV2ScoresALoggedRaidersFiveComponents(t *testing.T) {
	fx := seedLoggedRaider(t)
	got, err := fx.store.ScoreV2(context.Background(), fx.main)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ScoreV2(%s) = %+v", fx.main, got)

	full, err := sims.FullGearScore(context.Background(), "1.60.1.70291", "warrior-arms")
	if err != nil {
		t.Fatal(err)
	}
	wantGear := int(math.Round(GearMax * fx.gearScore / full))
	wantParses := int(math.Round(fx.parsePercentile / 100 * ParsesMax))
	wantDiscipline := int(math.Round(DisciplineMax * float64(fx.rated-fx.missedFlask) / float64(fx.rated)))
	want := ScoreComponents{
		Gear:       wantGear,
		Experience: 6 * ExperiencePerRaidBoss, // 990601..990606, the alts' kills included
		Parses:     wantParses,
		Attendance: AttendanceMax, // three guild nights, present at all three
		Discipline: wantDiscipline,
	}
	if got.Components != want {
		t.Errorf("components = %+v, want %+v", got.Components, want)
	}
	for name, c := range map[string]struct{ v, max int }{
		"gear": {got.Components.Gear, GearMax}, "experience": {got.Components.Experience, ExperienceMax},
		"parses": {got.Components.Parses, ParsesMax}, "attendance": {got.Components.Attendance, AttendanceMax},
		"discipline": {got.Components.Discipline, DisciplineMax},
	} {
		if c.v <= 0 || c.v > c.max {
			t.Errorf("%s = %d, outside (0, %d]", name, c.v, c.max)
		}
	}
	if wantParses < 420 || wantParses > 432 {
		t.Errorf("a 71st percentile best parse scales to %d, want about 426 of 600", wantParses)
	}
	if wantDiscipline != 92 {
		t.Errorf("eleven of twelve prepared fights is %d, want 92", wantDiscipline)
	}
	sum := want.Gear + want.Experience + want.Parses + want.Attendance + want.Discipline
	if got.Score != sum {
		t.Errorf("score = %d, want the components' sum %d", got.Score, sum)
	}
	if got.Tier != Tier(sum) {
		t.Errorf("tier = %q, want %q", got.Tier, Tier(sum))
	}
	if !got.AsOf.UTC().Equal(fx.newest.UTC()) {
		t.Errorf("as_of = %v, want the newest fight %v", got.AsOf, fx.newest)
	}
	if got.Experience != (ScoreExperience{Dungeons: 0, RaidKills: 6, Alts: 2}) {
		t.Errorf("experience = %+v, want 0 dungeons, 6 raid kills, 2 alts", got.Experience)
	}
	// Contract 5: gear_score is the gear component in points; attendance is the fraction.
	if got.GearScore != got.Components.Gear || got.GearScoreRaw != int(math.Round(fx.gearScore)) || got.ItemLevel != 58 {
		t.Errorf("gear score %d (raw %d) at item level %d, want %d (raw %d) at 58",
			got.GearScore, got.GearScoreRaw, got.ItemLevel, got.Components.Gear, int(math.Round(fx.gearScore)))
	}
	if got.Attendance != 1 {
		t.Errorf("attendance = %v, want 1 (every night)", got.Attendance)
	}
	if got.Spec != "warrior-arms" || got.Fights != fx.rated {
		t.Errorf("spec %q over %d fights, want warrior-arms over %d", got.Spec, got.Fights, fx.rated)
	}
}

// Two of three guild nights is two thirds of the attendance component.
func TestScoreV2AttendanceCountsThePartOfTheNightsAttended(t *testing.T) {
	fx := seedLoggedRaider(t)
	// The main was not at the first night: none of its rows there remain.
	mustExec(t, fx.pool, `delete from rating_scores where report_id = $1 and player_key = $2`, sv2Report+"a", fx.main)
	mustExec(t, fx.pool, `delete from fight_metrics where report_id = $1 and player_key = $2`, sv2Report+"a", fx.main)
	got, err := fx.store.ScoreV2(context.Background(), fx.main)
	if err != nil {
		t.Fatal(err)
	}
	if got.Components.Attendance != 133 || math.Abs(got.Attendance-2.0/3) > 1e-9 {
		t.Errorf("attendance %d (%v), want 133 (2/3)", got.Components.Attendance, got.Attendance)
	}
}

// "No logs" is never a zero score: a guild member with no rated fight has none at all.
func TestScoreV2WithNoRatedFightIsErrNoLogs(t *testing.T) {
	fx := seedLoggedRaider(t)
	got, err := fx.store.ScoreV2(context.Background(), fx.altA) // ranked, never rated
	if !errors.Is(err, ErrNoLogs) {
		t.Fatalf("err = %v, want ErrNoLogs", err)
	}
	if got != (ScoreV2{}) {
		t.Errorf("ScoreV2 = %+v alongside ErrNoLogs, want the zero value", got)
	}
	mustExec(t, fx.pool, `insert into guild_characters (guild_id, character_key, user_id, verified_at)
		values ($1, $2, $3, now())`, fx.guildID, fx.altB, fx.userID)
	if _, err := fx.store.ScoreV2(context.Background(), fx.altB); !errors.Is(err, ErrNoLogs) {
		t.Fatalf("a verified guild member with no rated fight: err = %v, want ErrNoLogs", err)
	}
	if _, err := fx.store.ScoreV2(context.Background(), "us/normal/sv2nobody"); !errors.Is(err, ErrNoLogs) {
		t.Fatalf("a character nobody has seen: err = %v, want ErrNoLogs", err)
	}
}

// A fight rated only in a private report gives no score: the score is public.
func TestScoreV2CountsOnlyPublicReports(t *testing.T) {
	fx := seedLoggedRaider(t)
	mustExec(t, fx.pool, `update reports set visibility = 'private' where id like 'sv2-%'`)
	if _, err := fx.store.ScoreV2(context.Background(), fx.main); !errors.Is(err, ErrNoLogs) {
		t.Fatalf("only private fights: err = %v, want ErrNoLogs", err)
	}
}

// The rater rolls the character's score up onto the fight it just rated.
func TestRaterRollsTheScoreUpOntoTheRatedFight(t *testing.T) {
	fx := seedLoggedRaider(t)
	f := RatedFight{ReportID: sv2Report + "c", FightIndex: 4, Region: "us", Ruleset: "normal",
		Summary: summary.Summary{Roster: []summary.RosterRow{{Name: "Sv2main"}, {Name: "Sv2nobody"}}}}
	if err := fx.store.RollUpScores(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	want, err := fx.store.ScoreV2(context.Background(), fx.main)
	if err != nil {
		t.Fatal(err)
	}
	var score, gear, experience int
	var attendance float64
	var discipline bool
	var tier string
	if err := fx.pool.QueryRow(context.Background(),
		`select score, gear_score, experience, attendance::float8, discipline, tier from rating_scores
		 where report_id = $1 and fight_index = 4 and player_key = $2`, f.ReportID, fx.main).
		Scan(&score, &gear, &experience, &attendance, &discipline, &tier); err != nil {
		t.Fatal(err)
	}
	if score != want.Score || gear != want.Components.Gear || experience != want.Components.Experience ||
		attendance != want.Attendance || !discipline || tier != want.Tier {
		t.Errorf("rolled up %d/%d/%d/%v/%v/%q, want %d/%d/%d/%v/true/%q", score, gear, experience, attendance,
			discipline, tier, want.Score, want.Components.Gear, want.Components.Experience, want.Attendance, want.Tier)
	}
}

// --- the character endpoint -------------------------------------------------------------

func TestTheCharacterEndpointCarriesTheForeverScore(t *testing.T) {
	fx := seedLoggedRaider(t)
	svc := &Service{Store: fx.store}
	mux := http.NewServeMux()
	rankings.Mount(mux, &rankings.Service{Store: &rankings.Store{Pool: fx.pool}, Scores: svc})

	read := func(name string) map[string]json.RawMessage {
		t.Helper()
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v1/characters/us/normal/"+name, nil))
		if rec.Code != http.StatusOK {
			t.Fatalf("%s: status %d: %s", name, rec.Code, rec.Body)
		}
		var env struct {
			Data map[string]json.RawMessage `json:"data"`
		}
		if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
			t.Fatal(err)
		}
		return env.Data
	}

	body := read("sv2main")
	for _, field := range []string{"character", "best", "history", "builds_seen"} {
		if _, ok := body[field]; !ok {
			t.Errorf("the existing field %q is gone", field)
		}
	}
	want, err := fx.store.ScoreV2(context.Background(), fx.main)
	if err != nil {
		t.Fatal(err)
	}
	var score int
	var tier, asOf string
	var components struct{ Gear, Experience, Parses, Attendance, Discipline int }
	var experience struct {
		Dungeons  int `json:"dungeons"`
		RaidKills int `json:"raid_kills"`
		Alts      int `json:"alts"`
	}
	for field, into := range map[string]any{
		"score": &score, "tier": &tier, "as_of": &asOf, "components": &components, "experience": &experience,
	} {
		if err := json.Unmarshal(body[field], into); err != nil {
			t.Fatalf("%s = %s: %v", field, body[field], err)
		}
	}
	if wantAsOf := want.AsOf.UTC().Format(time.DateOnly); score != want.Score || tier != want.Tier || asOf != wantAsOf {
		t.Errorf("score %d %q as of %s, want %d %q as of %s", score, tier, asOf, want.Score, want.Tier, wantAsOf)
	}
	if components != (struct{ Gear, Experience, Parses, Attendance, Discipline int }(want.Components)) {
		t.Errorf("components = %+v, want %+v", components, want.Components)
	}
	if experience.RaidKills != 6 || experience.Alts != 2 || experience.Dungeons != 0 {
		t.Errorf("experience = %+v", experience)
	}

	// An alt with ranked fights but no rated one: the page is there, the score is null.
	alt := read("sv2alta")
	for _, field := range []string{"score", "tier", "components", "as_of", "experience"} {
		if string(alt[field]) != "null" {
			t.Errorf("no logs: %s = %s, want null", field, alt[field])
		}
	}
}

type failingScorer struct{}

func (failingScorer) ForeverScore(context.Context, string) (rankings.ForeverScore, error) {
	return rankings.ForeverScore{}, errors.New("the score store is down")
}

// A score that cannot be read is logged and left out; the character page still renders.
func TestTheCharacterPageSurvivesAScoreFailure(t *testing.T) {
	fx := seedLoggedRaider(t)
	mux := http.NewServeMux()
	rankings.Mount(mux, &rankings.Service{Store: &rankings.Store{Pool: fx.pool}, Scores: failingScorer{}})
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v1/characters/us/normal/sv2main", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d: %s", rec.Code, rec.Body)
	}
	var env struct {
		Data map[string]json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if _, ok := env.Data["history"]; !ok {
		t.Error("the page lost its history")
	}
	for _, field := range []string{"score", "tier", "components", "as_of", "experience"} {
		if string(env.Data[field]) != "null" {
			t.Errorf("%s = %s, want null", field, env.Data[field])
		}
	}
}
