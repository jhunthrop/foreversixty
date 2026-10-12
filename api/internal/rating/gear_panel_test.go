package rating

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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

// gp67Picks are warrior-arms' band-60 alliance bare picks through L1.4's walk: Hand of
// Justice and Blackhand's Breadth are client-unconfirmed, so Frozen Heart of the Mountain
// holds trinket2; the band has no off-hand pick.
var gp67Picks = map[string]int{
	"head": 12640, "neck": 22150, "shoulder": 20057, "back": 272414, "chest": 252483,
	"wrist": 19578, "hands": 15063, "waist": 20041, "legs": 15062, "feet": 20048,
	"finger1": 234034, "finger2": 19325, "trinket1": 19289, "trinket2": 249469,
	"main_hand": 19323, "ranged": 16004,
}

type gearPanelFixture struct {
	pool *pgxpool.Pool
	mux  *http.ServeMux
}

func newGearPanelFixture(t *testing.T) *gearPanelFixture {
	t.Helper()
	pool := testPool(t)
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
	rankings.Mount(mux, &rankings.Service{Store: &rankings.Store{Pool: pool}, Scores: svc, Gear: svc})
	return &gearPanelFixture{pool: pool, mux: mux}
}

// seed writes a warrior named name (key us/normal/<lower name>) with one rated Arms fight in a
// public report, and, when gear is not nil, an addon export of a level-60 human wearing it.
func (fx *gearPanelFixture) seed(t *testing.T, name string, gear map[string]int) string {
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
	if gear != nil {
		var entries []string
		for slot, id := range gear {
			entries = append(entries, fmt.Sprintf("%s=%d", slot, id))
		}
		export := "FS1:" + gp67Build + ":warrior:human:0/0/0:" + strings.Join(entries, ",") + "|level=60"
		mustExec(t, fx.pool, `insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at)
			values ($1, $2, 'us', 'normal', $3, $4, now())`, key, userID, name, export)
	}
	reportID := "gp67-" + strings.ToLower(name)
	at := time.Now().UTC().Add(-48 * time.Hour)
	mustExec(t, fx.pool, `insert into reports (id, owner_id, title, visibility, zone, status, created_at)
		values ($1, $2, 'Gear panel', 'public', 'Molten Core', 'complete', $3)`, reportID, userID, at)
	mustExec(t, fx.pool, `insert into fights (report_id, fight_index, encounter_id, name, difficulty, kill, start_ms)
		values ($1, 1, 990671, 'Gear panel boss', 3, true, $2)`, reportID, at.UnixMilli())
	row := summary.RosterRow{GUID: "gp67-" + name, Name: name, Class: "Warrior", Spec: "Arms", Role: "dps"}
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
		"chest":    252484, // Timbermaw Tunic; the pick is Dawn Armor (252483)
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

	// Dawn Armor scores 506.2 and Timbermaw Tunic 300.8 under the arms weights; the band-60
	// mark is 4766.6, so the pick adds 500 * 205.4 / 4766.6 = 21.5457... points.
	chest := gp67Slot(t, body, "chest")
	if chest.Worn == nil || chest.Worn.ID != 252484 || chest.Pick == nil || chest.Pick.ID != 252483 ||
		chest.Pick.Name != "Dawn Armor" || chest.Pick.Source != "Leatherworking" || chest.Gap == nil || *chest.Gap != 21.5 {
		t.Errorf("chest = worn %+v pick %+v gap %s; want Timbermaw Tunic under Dawn Armor, gap 21.5", chest.Worn, chest.Pick, fmtGap(chest.Gap))
	}

	// Nothing worn at the neck: the pick (Beads of Ogre Might, 284) adds 500 * 284 / 4766.6.
	neck := gp67Slot(t, body, "neck")
	if neck.Worn != nil || neck.Pick == nil || neck.Pick.ID != 22150 || neck.Gap == nil || *neck.Gap != 29.8 {
		t.Errorf("neck = worn %+v pick %+v gap %s; want nothing worn, Beads of Ogre Might, gap 29.8", neck.Worn, neck.Pick, fmtGap(neck.Gap))
	}

	// trinket2's own pick and first runner-up are client-unconfirmed: the promoted pick has no chance.
	tr := gp67Slot(t, body, "trinket2")
	if tr.Pick == nil || tr.Pick.ID != 249469 || tr.Pick.Chance != nil {
		t.Errorf("trinket2 pick = %+v, want Frozen Heart of the Mountain (249469) with no chance", tr.Pick)
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
// mark is the band's best set across presets and factions (warrior-arms band 60: horde raid,
// 4766.6), and the alliance bare picks through the walk score 4514.4, so Gear is
// round(500 * 4514.4 / 4766.6) = 474, 26 below its top.
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
	if body.Components == nil || body.Components.Gear != 474 {
		t.Errorf("components.gear = %+v, want 474 (26 below its top of %d)", body.Components, GearMax)
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
