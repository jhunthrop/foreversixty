// api/internal/guilds/readiness_core_test.go
package guilds

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
)

func testBand() bis.Band {
	return bis.Band{Slots: []bis.Slot{
		{Slot: "chest", ItemID: 100, Alternatives: []bis.Alternative{{ItemID: 200, DpsDelta: -5}}},
		{Slot: "feet", ItemID: 300},
	}}
}

func lvl(n int) *int { return &n }

// at60 is a level-60 input with consent, everything else empty; tests fill what they check.
func at60(consent string) readinessInput {
	return readinessInput{Consent: consent, Level: lvl(60)}
}

func TestComputeReadinessWithoutConsentChecksNothingGear(t *testing.T) {
	in := at60("roster")
	in.Band, in.HasBand, in.Gear, in.HasBagsSection = testBand(), true, map[string]int{"chest": 999}, true
	cr := computeReadiness(in)
	if cr.GearChecked || cr.EnchantsChecked {
		t.Fatalf("checked = gear:%v enchants:%v, want both false without gear consent", cr.GearChecked, cr.EnchantsChecked)
	}
	if cr.ConsumablesState != "unknown" {
		t.Fatalf("consumables = %q, want unknown", cr.ConsumablesState)
	}
}

func TestComputeReadinessGearUpgradeAndEnchantGap(t *testing.T) {
	in := at60("gear")
	in.Band, in.HasBand = testBand(), true
	in.Gear = map[string]int{"chest": 200, "feet": 300} // chest is the listed alt (+5), feet is the pick (no gap)
	in.Enchants = map[string]int{"feet": 42}            // chest equipped, no enchant; feet enchanted
	cr := computeReadiness(in)
	if !cr.GearChecked || cr.GearGap.Upgrades != 1 || cr.GearGap.GainDps != 5 || cr.GearUnit != "dps" {
		t.Fatalf("gearGap = %+v unit %q, want 1 upgrade worth +5 DPS", cr.GearGap, cr.GearUnit)
	}
	if !cr.EnchantsChecked || !reflect.DeepEqual(cr.MissingEnchantSlots, []string{"chest"}) {
		t.Fatalf("missingEnchants = %v, want [chest]", cr.MissingEnchantSlots)
	}
}

// An export with no gear list is gear unknown: neither the gear gap nor the enchants are checked,
// never "all enchanted" (guild round 3).
func TestComputeReadinessWithoutGearChecksNoGearOrEnchants(t *testing.T) {
	in := at60("gear")
	in.Band, in.HasBand = testBand(), true
	cr := computeReadiness(in)
	if cr.GearChecked || cr.EnchantsChecked || len(cr.MissingEnchantSlots) != 0 {
		t.Fatalf("readiness = %+v, want gear and enchants unchecked with no gear on file", cr)
	}
	if got := readinessWord(cr, "gear"); got != readinessNotChecked {
		t.Fatalf("word = %q, want %q", got, readinessNotChecked)
	}
}

// A healer's or tank's band measures no DPS: the gap has no unit and never prints one.
func TestComputeReadinessHealerGapHasNoUnit(t *testing.T) {
	in := at60("gear")
	band := testBand()
	band.Role = "healer"
	in.Band, in.HasBand, in.Gear = band, true, map[string]int{"chest": 200}
	cr := computeReadiness(in)
	if cr.GearUnit != "" {
		t.Fatalf("unit = %q, want none for a healer band", cr.GearUnit)
	}
	if got := cr.failText()[0]; strings.Contains(got, "DPS") {
		t.Fatalf("healer fail text = %q, want no DPS figure", got)
	}
}

func TestComputeReadinessConsumablesGatedOnGearBagsConsent(t *testing.T) {
	short := at60("gear_bags")
	short.Bags, short.HasBagsSection = []int{}, true
	if got := computeReadiness(short).ConsumablesState; got != "short" {
		t.Fatalf("consumables = %q, want short for an empty bags section", got)
	}
	stocked := at60("gear_bags")
	stocked.Bags, stocked.HasBagsSection = []int{13510}, true
	if got := computeReadiness(stocked).ConsumablesState; got != "stocked" {
		t.Fatalf("consumables = %q, want stocked", got)
	}
	unknown := at60("gear")
	unknown.Bags, unknown.HasBagsSection = []int{13510}, true
	if got := computeReadiness(unknown).ConsumablesState; got != "unknown" {
		t.Fatalf("consumables = %q, want unknown without gear_bags consent even with bag data present", got)
	}
}

func TestComputeReadinessTalentPointsUnspent(t *testing.T) {
	in := at60("roster")
	in.TalentPointsSpent, in.HasTalents = 45, true
	cr := computeReadiness(in)
	if cr.TalentPointsUnspent != 6 || cr.TalentLevelAssumed {
		t.Fatalf("unspent = %d assumed %v, want 6 (51-45) on a real level", cr.TalentPointsUnspent, cr.TalentLevelAssumed)
	}
	in.TalentPointsSpent = 51
	if full := computeReadiness(in); full.TalentPointsUnspent != 0 {
		t.Fatalf("unspent = %d, want 0 at full spend", full.TalentPointsUnspent)
	}

	noLevel := readinessInput{Consent: "roster", TalentPointsSpent: 11, HasTalents: true}
	cr = computeReadiness(noLevel)
	if cr.TalentPointsUnspent != 40 || !cr.TalentLevelAssumed {
		t.Fatalf("unspent = %d assumed %v, want 40 (51-11) assuming level 60 with no level on file", cr.TalentPointsUnspent, cr.TalentLevelAssumed)
	}
}

// Talent points follow the level: 1 at level 10, one more each level, 51 at 60, none below 10.
func TestMaxTalentPointsFollowTheLevel(t *testing.T) {
	cases := []struct{ level, want int }{{5, 0}, {9, 0}, {10, 1}, {23, 14}, {52, 43}, {60, 51}}
	for _, c := range cases {
		if got, assumed := maxTalentPoints(lvl(c.level)); got != c.want || assumed {
			t.Errorf("maxTalentPoints(%d) = (%d, %v), want (%d, false)", c.level, got, assumed, c.want)
		}
	}
	if got, assumed := maxTalentPoints(nil); got != 51 || !assumed {
		t.Fatalf("maxTalentPoints(nil) = (%d, %v), want (51, true)", got, assumed)
	}
}

// A character under 60 is leveling: nothing is graded, whatever its gear, bags or talents say.
func TestComputeReadinessLevelingGradesNothing(t *testing.T) {
	in := readinessInput{
		Consent: "gear_bags", Band: testBand(), HasBand: true, Gear: map[string]int{"chest": 999},
		Bags: []int{}, HasBagsSection: true, HasTalents: true, Level: lvl(52), ItemLevel: lvl(40), MedianItemLevel: lvl(60),
	}
	cr := computeReadiness(in)
	if !cr.Leveling || cr.GearChecked || cr.EnchantsChecked || cr.ConsumablesState != "unknown" ||
		cr.TalentPointsUnspent != 0 || cr.ItemLevelDelta != nil || cr.FailingCount() != 0 {
		t.Fatalf("leveling readiness = %+v, want nothing graded", cr)
	}
	if cr.ItemLevel == nil || *cr.ItemLevel != 40 || cr.Level == nil || *cr.Level != 52 {
		t.Fatalf("level %v item level %v, want 52 and 40 kept", cr.Level, cr.ItemLevel)
	}
}

func TestComputeReadinessItemLevelDelta(t *testing.T) {
	in := at60("roster")
	in.ItemLevel, in.MedianItemLevel = lvl(58), lvl(63)
	cr := computeReadiness(in)
	if cr.ItemLevel == nil || *cr.ItemLevel != 58 {
		t.Fatalf("itemLevel = %v, want 58", cr.ItemLevel)
	}
	if cr.ItemLevelDelta == nil || *cr.ItemLevelDelta != -5 {
		t.Fatalf("itemLevelDelta = %v, want -5", cr.ItemLevelDelta)
	}
}

func TestFailTextOrderAndNeedsCap(t *testing.T) {
	cr := CharacterReadiness{
		GearChecked: true, GearGap: bis.GearGap{Upgrades: 3, GainDps: 23.8}, GearUnit: "dps",
		EnchantsChecked: true, MissingEnchantSlots: []string{"chest", "feet"},
		ConsumablesState:    "short",
		TalentPointsUnspent: 1,
	}
	fails := cr.failText()
	want := []string{
		"3 gear upgrades waiting (+23.8 DPS)",
		"no enchant: Chest, Feet",
		"consumables short",
		"1 unspent talent point",
	}
	if !reflect.DeepEqual(fails, want) {
		t.Fatalf("failText = %v, want %v", fails, want)
	}
	if cr.FailingCount() != 4 {
		t.Fatalf("failingCount = %d, want 4", cr.FailingCount())
	}
	if got := cr.Needs(3); len(got) != 3 {
		t.Fatalf("Needs(3) = %v, want 3 entries", got)
	}
	if got := cr.Needs(0); len(got) != 4 {
		t.Fatalf("Needs(0) = %v, want uncapped (4)", got)
	}
}

// An upgrade nothing sim-checked has no gain to print: never "(+0.0 DPS)".
func TestFailTextOmitsAnUnsimmedGain(t *testing.T) {
	cr := CharacterReadiness{GearChecked: true, GearGap: bis.GearGap{Upgrades: 2, NotSimChecked: 2}, GearUnit: "dps"}
	if got := cr.failText(); !reflect.DeepEqual(got, []string{"2 gear upgrades waiting"}) {
		t.Fatalf("failText = %v, want the count alone", got)
	}
}

func TestFailTextCleanRowIsEmpty(t *testing.T) {
	cr := CharacterReadiness{GearChecked: true, EnchantsChecked: true, ConsumablesState: "stocked"}
	if got := cr.failText(); len(got) != 0 {
		t.Fatalf("failText = %v, want empty for a clean row", got)
	}
}

// TestComputeReadinessTalentsRead: a zero unspent is only a real zero when the export carried
// talents at all; the board says "not read yet" for the other zero.
func TestComputeReadinessTalentsRead(t *testing.T) {
	in := at60("roster")
	in.TalentPointsSpent, in.HasTalents = 51, true
	if read := computeReadiness(in); !read.TalentsRead || read.TalentPointsUnspent != 0 {
		t.Fatalf("read = %v unspent = %d, want a real zero", read.TalentsRead, read.TalentPointsUnspent)
	}
	in.TalentPointsSpent, in.HasTalents = 0, false
	if unread := computeReadiness(in); unread.TalentsRead || unread.TalentPointsUnspent != 0 {
		t.Fatalf("read = %v unspent = %d, want an unread zero", unread.TalentsRead, unread.TalentPointsUnspent)
	}
}

// The nudge asks for the fix by the planned raid's day, and only lists it with none planned.
func TestNudgeTextNamesTheRaidDay(t *testing.T) {
	cr := CharacterReadiness{ConsumablesState: "short", Level: lvl(60)}
	if got := nudgeText("Rootbound", cr, "Tue 8 Dec"); got != "Rootbound: consumables short. Check before Tue 8 Dec." {
		t.Fatalf("nudge = %q", got)
	}
	if got := nudgeText("Rootbound", cr, ""); got != "Rootbound: consumables short." {
		t.Fatalf("nudge with no raid = %q", got)
	}
}

func TestRaidDayReadsTheGuildZone(t *testing.T) {
	// 02:00 UTC on Friday 11 Dec 2026 is Thursday 10 Dec in the evening in Chicago.
	start := time.Date(2026, 12, 11, 2, 0, 0, 0, time.UTC)
	if got := raidDay(start, "America/Chicago"); got != "Thu 10 Dec" {
		t.Fatalf("Chicago day = %q, want Thu 10 Dec", got)
	}
	if got := raidDay(start, ""); got != "Fri 11 Dec" {
		t.Fatalf("UTC day = %q, want Fri 11 Dec", got)
	}
}

// Graded rows lead worst-first; leveling rows follow, highest level first, then by name.
func TestSortReadinessRowsPutsLevelingLast(t *testing.T) {
	rows := []ReadinessRow{
		{Name: "Low", Level: lvl(20)},
		{Name: "Clean", Level: lvl(60)},
		{Name: "Bravo", Level: lvl(52)},
		{Name: "Worst", Level: lvl(60), readinessScore: 300},
		{Name: "Alpha", Level: lvl(52)},
	}
	sortReadinessRows(rows)
	var got []string
	for _, r := range rows {
		got = append(got, r.Name)
	}
	if want := []string{"Worst", "Clean", "Alpha", "Bravo", "Low"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("order = %v, want %v", got, want)
	}
}

// The bot's gear words: a healer's or tank's gap is the count alone, and a figure needs a gain.
func TestGearGapWordsPrintAFigureOnlyWithAUnitAndAGain(t *testing.T) {
	dps := gearUnitDPS
	cases := []struct {
		gap  GearGapView
		want string
	}{
		{GearGapView{Upgrades: 4, GainDps: 31}, "4 upgrades"},
		{GearGapView{Upgrades: 2, NotSimChecked: 2, Unit: &dps}, "2 upgrades · not simmed yet"},
		{GearGapView{Upgrades: 1, Unit: &dps}, "1 upgrade"},
		{GearGapView{Upgrades: 3, GainDps: 24, Unit: &dps}, "3 upgrades · +24 DPS"},
	}
	for _, c := range cases {
		if got := gearGapWords(c.gap); got != c.want {
			t.Errorf("gearGapWords(%+v) = %q, want %q", c.gap, got, c.want)
		}
	}
}

// The median is the level-60 raiders': a leveling character's item level is left out.
func TestMedianItemLevelSkipsLevelingRows(t *testing.T) {
	roster := []RosterRow{
		{ItemLevel: lvl(60), Level: lvl(60)},
		{ItemLevel: lvl(62), Level: lvl(60)},
		{ItemLevel: lvl(30), Level: lvl(41)},
	}
	if got := medianItemLevel(roster); got == nil || *got != 61 {
		t.Fatalf("median = %v, want 61 from the two 60s", got)
	}
}
