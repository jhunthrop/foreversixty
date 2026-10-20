package rating

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Task L6.7: the character page's gear panel, read through the real endpoint (rankings.Mount
// with rating.Service as both Scorer and GearReader) over seeded rows and the repository's
// real data/builds (build 1.60.1.70291, warrior-arms). Every row these tests write carries
// the "gp67" prefix (character slugs "gp67...", report ids "gp67-...", the email
// "gp67-..."), and the cleanup deletes only rows with that prefix.

const gp67Build = "1.60.1.70291"

// gp67Picks are warrior-arms' band-60 alliance bare picks through L1.4's walk (the ranker now
// ranks confirmed items only, so the walk promotes nothing here); the band has no off-hand pick.
var gp67Picks = map[string]int{
	"head": 12640, "neck": 22150, "shoulder": 272105, "back": 272414, "chest": 252484,
	"wrist": 19578, "hands": 272095, "waist": 227814, "legs": 22385, "feet": 21995,
	"finger1": 234034, "finger2": 19325, "trinket1": 19289, "trinket2": 19287,
	"main_hand": 19323, "ranged": 22656,
}

type gearPanelFixture struct {
	pool   *pgxpool.Pool
	mux    *http.ServeMux
	tracer *countingTracer
}

// countingTracer counts the statements a pool runs, by their SQL text, so a test can hold how
// often a page reads something (fix round 1, item 6).
type countingTracer struct {
	mu  sync.Mutex
	sql []string
}

func (c *countingTracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	c.mu.Lock()
	c.sql = append(c.sql, data.SQL)
	c.mu.Unlock()
	return ctx
}

func (c *countingTracer) TraceQueryEnd(context.Context, *pgx.Conn, pgx.TraceQueryEndData) {}

func (c *countingTracer) reset() {
	c.mu.Lock()
	c.sql = nil
	c.mu.Unlock()
}

// count is how many statements since the last reset contain fragment.
func (c *countingTracer) count(fragment string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	n := 0
	for _, q := range c.sql {
		if strings.Contains(q, fragment) {
			n++
		}
	}
	return n
}

func newGearPanelFixture(t *testing.T) *gearPanelFixture {
	t.Helper()
	testPool(t) // migrates
	cfg, err := pgxpool.ParseConfig(testURL(t))
	if err != nil {
		t.Fatal(err)
	}
	tracer := &countingTracer{}
	cfg.ConnConfig.Tracer = tracer
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	useRepoGearData(t)
	cleanup := func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'gp67-%'`)
		pool.Exec(bg, `delete from fight_metrics where report_id like 'gp67-%'`)
		pool.Exec(bg, `delete from fights where report_id like 'gp67-%'`)
		pool.Exec(bg, `delete from reports where id like 'gp67-%'`)
		pool.Exec(bg, `delete from addon_exports where character_key like 'us/normal/gp67%'`)
		pool.Exec(bg, `delete from characters where key like 'us/normal/gp67%'`)
		pool.Exec(bg, `delete from users where email like 'gp67-%'`)
	}
	cleanup()
	t.Cleanup(cleanup)
	svc := &Service{Store: &Store{Pool: pool}}
	mux := http.NewServeMux()
	rankings.Mount(mux, &rankings.Service{Store: &rankings.Store{Pool: pool}, Scores: svc, ScoreGear: svc})
	return &gearPanelFixture{pool: pool, mux: mux, tracer: tracer}
}

// seedOpts vary a seeded character: its rated spec (Arms when empty), its export's race slug
// (human when empty) and its characters.faction (none when empty).
type seedOpts struct {
	spec, race, faction string
}

// seed writes a warrior named name (key us/normal/<lower name>) with one rated Arms fight in a
// public report, and, when gear is not nil, an addon export of a level-60 human wearing it.
func (fx *gearPanelFixture) seed(t *testing.T, name string, gear map[string]int) string {
	t.Helper()
	return fx.seedWith(t, name, gear, seedOpts{})
}

// seedWith is seed with opts.
func (fx *gearPanelFixture) seedWith(t *testing.T, name string, gear map[string]int, opts seedOpts) string {
	t.Helper()
	if opts.spec == "" {
		opts.spec = "Arms"
	}
	if opts.race == "" {
		opts.race = "human"
	}
	ctx := context.Background()
	key := "us/normal/" + strings.ToLower(name)
	var userID int64
	if err := fx.pool.QueryRow(ctx, `insert into users (email) values ($1) returning id`,
		"gp67-"+strings.ToLower(name)+"@example.com").Scan(&userID); err != nil {
		t.Fatal(err)
	}
	var faction *string
	if opts.faction != "" {
		faction = &opts.faction
	}
	mustExec(t, fx.pool, `insert into characters (key, region, ruleset, name, class, user_id, faction)
		values ($1, 'us', 'normal', $2, 'Warrior', $3, $4)`, key, name, userID, faction)
	if gear != nil {
		var entries []string
		for slot, id := range gear {
			entries = append(entries, fmt.Sprintf("%s=%d", slot, id))
		}
		export := "FS1:" + gp67Build + ":warrior:" + opts.race + ":0/0/0:" + strings.Join(entries, ",") + "|level=60"
		mustExec(t, fx.pool, `insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at)
			values ($1, $2, 'us', 'normal', $3, $4, now())`, key, userID, name, export)
	}
	reportID := "gp67-" + strings.ToLower(name)
	at := time.Now().UTC().Add(-48 * time.Hour)
	mustExec(t, fx.pool, `insert into reports (id, owner_id, title, visibility, zone, status, created_at)
		values ($1, $2, 'Gear panel', 'public', 'Molten Core', 'complete', $3)`, reportID, userID, at)
	mustExec(t, fx.pool, `insert into fights (report_id, fight_index, encounter_id, name, difficulty, kill, start_ms)
		values ($1, 1, 990671, 'Gear panel boss', 3, true, $2)`, reportID, at.UnixMilli())
	row := summary.RosterRow{GUID: "gp67-" + name, Name: name, Class: "Warrior", Spec: opts.spec, Role: "dps"}
	if err := (&Store{Pool: fx.pool}).RateFight(ctx, RatedFight{
		ReportID: reportID, FightIndex: 1, Region: "us", Ruleset: "normal", FoughtAt: at, EncounterID: 990671,
		Summary: summary.Summary{EncounterID: 990671, Difficulty: 3, DurationMS: 180000, Kill: true,
			Roster: []summary.RosterRow{row}, Combatants: []summary.CombatantRow{preparedCombatant(row.GUID)}},
	}); err != nil {
		t.Fatal(err)
	}
	return key
}

type gearPanelBody struct {
	Gear         []rankings.GearSlot `json:"gear"`
	GearGapTotal *float64            `json:"gear_gap_total"`
	Components   *struct {
		Gear int `json:"gear"`
	} `json:"components"`
	raw map[string]json.RawMessage
}

func (fx *gearPanelFixture) read(t *testing.T, name string) gearPanelBody {
	t.Helper()
	rec := httptest.NewRecorder()
	fx.mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v1/characters/us/normal/"+name, nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("%s: status %d: %s", name, rec.Code, rec.Body)
	}
	var env struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	var out gearPanelBody
	if err := json.Unmarshal(env.Data, &out); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(env.Data, &out.raw); err != nil {
		t.Fatal(err)
	}
	return out
}

func gp67Slot(t *testing.T, body gearPanelBody, slot string) rankings.GearSlot {
	t.Helper()
	for _, s := range body.Gear {
		if s.Slot == slot {
			return s
		}
	}
	t.Fatalf("no %s row in %+v", slot, body.Gear)
	return rankings.GearSlot{}
}

func fmtGap(g *float64) string {
	if g == nil {
		return "null"
	}
	return fmt.Sprint(*g)
}

// Ruling 1 and 5: a seeded level-60 alliance arms warrior with an export gets the 17 slots in
// the export's order, and three slots read as expected: the off-hand has no pick (gap null),
// the head is worn as the pick (gap 0), and the chest is an upgrade. An empty slot is worn
// null and its gap is the pick's whole worth.
func TestTheGearPanelListsTheSeventeenSlotsWithPicksAndGaps(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seed(t, "Gp67mixed", map[string]int{
		"head":     12640,  // Lionheart Helm, the band's pick
		"chest":    6627,   // a strength-10 chest piece (20); the pick is Timbermaw Tunic (252484)
		"off_hand": 271798, // Desperate Barrier; the band has no off-hand pick
	})
	body := fx.read(t, "gp67mixed")

	var slots []string
	for _, s := range body.Gear {
		slots = append(slots, s.Slot)
	}
	want := "head,neck,shoulder,back,chest,wrist,hands,waist,legs,feet,finger1,finger2,trinket1,trinket2,main_hand,off_hand,ranged"
	if got := strings.Join(slots, ","); got != want {
		t.Fatalf("slots = %s, want %s", got, want)
	}

	off := gp67Slot(t, body, "off_hand")
	if off.Worn == nil || off.Worn.ID != 271798 || off.Worn.Name == nil || *off.Worn.Name != "Desperate Barrier" ||
		off.Worn.ItemLevel == nil || *off.Worn.ItemLevel != 43 || off.Pick != nil || off.Gap != nil {
		t.Errorf("off_hand = worn %+v pick %+v gap %s; want Desperate Barrier (43), no pick, gap null", off.Worn, off.Pick, fmtGap(off.Gap))
	}

	head := gp67Slot(t, body, "head")
	if head.Worn == nil || head.Worn.ID != 12640 || head.Pick == nil || head.Pick.ID != 12640 ||
		head.Pick.Name != "Lionheart Helm" || head.Pick.Source != "Blacksmithing" || head.Pick.Chance != nil ||
		head.Gap == nil || *head.Gap != 0 {
		t.Errorf("head = worn %+v pick %+v gap %s; want Lionheart Helm worn and picked, gap 0", head.Worn, head.Pick, fmtGap(head.Gap))
	}

	// Timbermaw Tunic scores 300.8 and the worn 6627 (strength 10) 20 under the arms weights;
	// the band-60 mark is 5344.4 (the horde bare set), so the pick adds
	// 500 * 280.8 / 5344.4 = 26.271... points.
	chest := gp67Slot(t, body, "chest")
	if chest.Worn == nil || chest.Worn.ID != 6627 || chest.Pick == nil || chest.Pick.ID != 252484 ||
		chest.Pick.Name != "Timbermaw Tunic" || chest.Pick.Source != "Leatherworking" || chest.Gap == nil || *chest.Gap != 26.3 {
		t.Errorf("chest = worn %+v pick %+v gap %s; want 6627 under Timbermaw Tunic, gap 26.3", chest.Worn, chest.Pick, fmtGap(chest.Gap))
	}

	// Nothing worn at the neck: the pick (Beads of Ogre Might, 284) adds 500 * 284 / 5344.4.
	neck := gp67Slot(t, body, "neck")
	if neck.Worn != nil || neck.Pick == nil || neck.Pick.ID != 22150 || neck.Gap == nil || *neck.Gap != 26.6 {
		t.Errorf("neck = worn %+v pick %+v gap %s; want nothing worn, Beads of Ogre Might, gap 26.6", neck.Worn, neck.Pick, fmtGap(neck.Gap))
	}

	// trinket2's pick is a quest reward: no drop chance.
	tr := gp67Slot(t, body, "trinket2")
	if tr.Pick == nil || tr.Pick.ID != 19287 || tr.Pick.Chance != nil {
		t.Errorf("trinket2 pick = %+v, want Darkmoon Card: Heroism (19287) with no chance", tr.Pick)
	}

	// Null rules are serialised as null, never omitted.
	var rows []map[string]json.RawMessage
	if err := json.Unmarshal(body.raw["gear"], &rows); err != nil {
		t.Fatal(err)
	}
	for _, r := range rows {
		for _, k := range []string{"slot", "worn", "pick", "gap"} {
			if _, ok := r[k]; !ok {
				t.Errorf("row %s has no %q key", r["slot"], k)
			}
		}
	}
	if string(rows[15]["pick"]) != "null" || string(rows[15]["gap"]) != "null" || string(rows[1]["worn"]) != "null" {
		t.Errorf("off_hand pick %s gap %s, neck worn %s; want null", rows[15]["pick"], rows[15]["gap"], rows[1]["worn"])
	}
}

// Ruling 2 (as re-ruled): gear_gap_total is the plain sum of the non-null gaps, one function.
func TestGearGapTotalIsTheSumOfTheGaps(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seed(t, "Gp67sum", map[string]int{"head": 12640, "chest": 252484, "off_hand": 271798})
	body := fx.read(t, "gp67sum")
	sum, n := 0.0, 0
	for _, s := range body.Gear {
		if s.Gap != nil {
			sum += *s.Gap
			n++
		}
	}
	if n != 16 {
		t.Errorf("%d non-null gaps, want 16 (every slot but the off-hand)", n)
	}
	if body.GearGapTotal == nil || math.Abs(*body.GearGapTotal-sum) > 1e-9 {
		t.Fatalf("gear_gap_total = %s, want the sum of the gaps %v", fmtGap(body.GearGapTotal), sum)
	}
	if got := rankings.GearGapTotal(body.Gear); got == nil || *got != *body.GearGapTotal {
		t.Errorf("rankings.GearGapTotal = %s, the page says %s", fmtGap(got), fmtGap(body.GearGapTotal))
	}
	if rankings.GearGapTotal([]rankings.GearSlot{{Slot: "off_hand"}}) != nil {
		t.Error("no gap at all: want a null total")
	}
}

// Ruling 2 (as re-ruled): a character wearing every pick of its own faction's bare set has
// every gap 0 and a gear_gap_total of 0, while its Gear component still sits below 500: the
// mark is the band's best set across presets and factions (warrior-arms band 60: horde bare,
// 5344.4), and the alliance bare picks through the walk score 5120.4, so Gear is
// round(500 * 5120.4 / 5344.4) = 479, 21 below its top.
func TestWearingEveryPickZeroesTheGapsButNotTheGearComponentsDistance(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seed(t, "Gp67picks", gp67Picks)
	body := fx.read(t, "gp67picks")
	for _, s := range body.Gear {
		if s.Slot == "off_hand" {
			if s.Pick != nil || s.Gap != nil {
				t.Errorf("off_hand pick %+v gap %s, want no pick and a null gap", s.Pick, fmtGap(s.Gap))
			}
			continue
		}
		if s.Pick == nil || s.Worn == nil || s.Worn.ID != s.Pick.ID || s.Gap == nil || *s.Gap != 0 {
			t.Errorf("%s: worn %+v pick %+v gap %s, want the pick worn and gap 0", s.Slot, s.Worn, s.Pick, fmtGap(s.Gap))
		}
	}
	if body.GearGapTotal == nil || *body.GearGapTotal != 0 {
		t.Errorf("gear_gap_total = %s, want 0", fmtGap(body.GearGapTotal))
	}
	if body.Components == nil || body.Components.Gear != 479 {
		t.Errorf("components.gear = %+v, want 479 (21 below its top of %d)", body.Components, GearMax)
	}
}

// Ruling 1 and 5: no export, no panel: gear is [] and its total null.
func TestTheGearPanelIsEmptyWithNoExport(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seed(t, "Gp67none", nil)
	body := fx.read(t, "gp67none")
	if string(body.raw["gear"]) != "[]" || string(body.raw["gear_gap_total"]) != "null" {
		t.Errorf("gear = %s, gear_gap_total = %s; want [] and null", body.raw["gear"], body.raw["gear_gap_total"])
	}
}

// Fix round 1, item 2: rings and trinkets are matched to the picks the way that gives the
// smaller total, so a pick worn in the other slot of its pair has a gap of 0.
func TestASwappedPairOfRingsHasNoGap(t *testing.T) {
	fx := newGearPanelFixture(t)
	// finger1's pick is the Signet Ring (234034, 308) and finger2's Don Julio's Band
	// (19325, 500); this character wears them the other way round.
	fx.seed(t, "Gp67swap", map[string]int{"finger1": 19325, "finger2": 234034})
	body := fx.read(t, "gp67swap")
	for _, slot := range []string{"finger1", "finger2"} {
		row := gp67Slot(t, body, slot)
		if row.Gap == nil || *row.Gap != 0 {
			t.Errorf("%s: worn %+v pick %+v gap %s, want 0 (the pick is worn in the other ring slot)",
				slot, row.Worn, row.Pick, fmtGap(row.Gap))
		}
	}
	if w := gp67Slot(t, body, "finger1").Worn; w == nil || w.ID != 19325 {
		t.Errorf("finger1 worn = %+v, want what the export says (19325)", w)
	}
}

// The same matching for the trinkets: trinket1's pick is 19289 and trinket2's Darkmoon Card:
// Heroism (19287); this character wears them the other way round.
func TestASwappedPairOfTrinketsHasNoGap(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seed(t, "Gp67trswap", map[string]int{"trinket1": 19287, "trinket2": 19289})
	body := fx.read(t, "gp67trswap")
	for _, slot := range []string{"trinket1", "trinket2"} {
		row := gp67Slot(t, body, slot)
		if row.Gap == nil || *row.Gap != 0 {
			t.Errorf("%s: worn %+v pick %+v gap %s, want 0 (the pick is worn in the other trinket slot)",
				slot, row.Worn, row.Pick, fmtGap(row.Gap))
		}
	}
	if w := gp67Slot(t, body, "trinket1").Worn; w == nil || w.ID != 19287 {
		t.Errorf("trinket1 worn = %+v, want what the export says (19287)", w)
	}
}

// Fix round 1, item 3, a picked two-hander: the arms band-60 pick is The Unstoppable Force
// (19323, two-handed, 486) with no off-hand pick. Worn Persuader (484) and a Grand Marshal's
// Left Hand Blade (252) score 736 together, more than the two-hander, so the main hand's gap is
// 0 (on its own slot it would read 500 x 2 / 5344.4 = 0.2) and the off-hand's is null.
func TestAPickedTwoHanderIsWeighedAgainstBothWornHands(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seed(t, "Gp67twohandpick", map[string]int{"main_hand": 22384, "off_hand": 18847})
	body := fx.read(t, "gp67twohandpick")
	mh, oh := gp67Slot(t, body, "main_hand"), gp67Slot(t, body, "off_hand")
	if mh.Pick == nil || mh.Pick.ID != 19323 || mh.Gap == nil || *mh.Gap != 0 {
		t.Errorf("main_hand pick %+v gap %s, want The Unstoppable Force and 0", mh.Pick, fmtGap(mh.Gap))
	}
	if oh.Pick != nil || oh.Gap != nil {
		t.Errorf("off_hand pick %+v gap %s, want no pick and null", oh.Pick, fmtGap(oh.Gap))
	}
}

// Fix round 1, item 3, a worn two-hander: the fury band-60 picks are Ravencrest's Legacy
// (21520, 31.0) in the main hand and Shadowsong's Sorrow (21522, 26.2) in the off-hand, both
// one-handed; the character wears a Centaur Spear (271794, two-handed, 20.9) and nothing in
// the off-hand. The main hand carries the pair's whole gap, 500 x (57.2 - 20.9) / 6047.1 =
// 3.0016 -> 3, and the off-hand 0.
func TestAWornTwoHanderIsWeighedAgainstBothPickedHands(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seedWith(t, "Gp67twohandworn", map[string]int{"main_hand": 271794}, seedOpts{spec: "Fury"})
	body := fx.read(t, "gp67twohandworn")
	mh, oh := gp67Slot(t, body, "main_hand"), gp67Slot(t, body, "off_hand")
	if mh.Pick == nil || mh.Pick.ID != 21520 || mh.Gap == nil || *mh.Gap != 3 {
		t.Errorf("main_hand pick %+v gap %s, want Ravencrest's Legacy and 3", mh.Pick, fmtGap(mh.Gap))
	}
	if oh.Pick == nil || oh.Pick.ID != 21522 || oh.Gap == nil || *oh.Gap != 0 {
		t.Errorf("off_hand pick %+v gap %s, want Shadowsong's Sorrow and 0", oh.Pick, fmtGap(oh.Gap))
	}
}

// Fix round 1, item 5: when the export's race is not one the build knows, the faction comes
// from characters.faction.
func TestTheFactionFallsBackToTheCharactersRow(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seedWith(t, "Gp67faction", map[string]int{"head": 12640}, seedOpts{race: "murloc", faction: "alliance"})
	body := fx.read(t, "gp67faction")
	if head := gp67Slot(t, body, "head"); head.Pick == nil || head.Pick.ID != 12640 || head.Gap == nil || *head.Gap != 0 {
		t.Errorf("head pick %+v gap %s, want the alliance pick Lionheart Helm and 0", head.Pick, fmtGap(head.Gap))
	}
	fx.seedWith(t, "Gp67nofaction", map[string]int{"head": 12640}, seedOpts{race: "murloc"})
	body = fx.read(t, "gp67nofaction")
	if head := gp67Slot(t, body, "head"); head.Pick != nil {
		t.Errorf("no faction anywhere: head pick %+v, want null", head.Pick)
	}
}

// Fix round 1, item 6: the Forever Score and the gear panel share one read of the export and
// one read of the rated fights per page.
func TestTheScoreAndTheGearPanelReadTheExportAndTheFightsOnce(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seed(t, "Gp67once", map[string]int{"head": 12640})
	fx.tracer.reset()
	body := fx.read(t, "gp67once")
	if body.Components == nil || len(body.Gear) != 17 {
		t.Fatalf("components %+v, %d gear rows; want a score and a panel", body.Components, len(body.Gear))
	}
	if n := fx.tracer.count("from addon_exports where character_key"); n != 1 {
		t.Errorf("%d reads of the export, want 1", n)
	}
	if n := fx.tracer.count("count(rs.discipline)"); n != 1 {
		t.Errorf("%d reads of the rated fights, want 1", n)
	}
}

// seedExportOnly writes a level-21 human warrior with an addon export whose talents are
// talents ("0/05052/0": twelve Fury points) and no logged fight at all.
func (fx *gearPanelFixture) seedExportOnly(t *testing.T, name, talents string) {
	t.Helper()
	ctx := context.Background()
	key := "us/normal/" + strings.ToLower(name)
	var userID int64
	if err := fx.pool.QueryRow(ctx, `insert into users (email) values ($1) returning id`,
		"gp67-"+strings.ToLower(name)+"@example.com").Scan(&userID); err != nil {
		t.Fatal(err)
	}
	mustExec(t, fx.pool, `insert into characters (key, region, ruleset, name, class, user_id)
		values ($1, 'us', 'normal', $2, 'Warrior', $3)`, key, name, userID)
	export := "FS1:" + gp67Build + ":warrior:human:" + talents + ":head=12640|level=21"
	mustExec(t, fx.pool, `insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at)
		values ($1, $2, 'us', 'normal', $3, $4, now())`, key, userID, name, export)
}

// A character with an export and no logged fight gets picks for the spec its talents name,
// and the panel says so.
func TestTheGearPanelReadsTheSpecFromTheExportsTalentsWhenNoFightNamedOne(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seedExportOnly(t, "Gp67nologs", "0/05052/0")
	body := fx.read(t, "gp67nologs")
	var spec, source string
	if err := json.Unmarshal(body.raw["gear_spec"], &spec); err != nil || spec != "warrior-fury" {
		t.Errorf("gear_spec = %s, want warrior-fury", body.raw["gear_spec"])
	}
	if err := json.Unmarshal(body.raw["gear_spec_source"], &source); err != nil || source != "talents" {
		t.Errorf("gear_spec_source = %s, want talents", body.raw["gear_spec_source"])
	}
	picks := 0
	for _, s := range body.Gear {
		if s.Pick != nil {
			picks++
		}
	}
	if picks == 0 {
		t.Errorf("no slot has a pick for a level-21 Fury warrior: %+v", body.Gear)
	}
}

// With no points spent the class's first spec answers, as the default.
func TestTheGearPanelFallsBackToTheClassDefaultWithNoPointsSpent(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seedExportOnly(t, "Gp67nopoints", "0/0/0")
	body := fx.read(t, "gp67nopoints")
	if got := string(body.raw["gear_spec"]); got != `"warrior-arms"` {
		t.Errorf("gear_spec = %s, want warrior-arms", got)
	}
	if got := string(body.raw["gear_spec_source"]); got != `"default"` {
		t.Errorf("gear_spec_source = %s, want default", got)
	}
}

// A rated fight's spec wins over the talents.
func TestTheGearPanelPrefersTheFightsSpecToTheTalents(t *testing.T) {
	fx := newGearPanelFixture(t)
	fx.seed(t, "Gp67fight", map[string]int{"head": 12640})
	body := fx.read(t, "gp67fight")
	if got := string(body.raw["gear_spec"]); got != `"warrior-arms"` {
		t.Errorf("gear_spec = %s, want warrior-arms", got)
	}
	if got := string(body.raw["gear_spec_source"]); got != `"fight"` {
		t.Errorf("gear_spec_source = %s, want fight", got)
	}
}
