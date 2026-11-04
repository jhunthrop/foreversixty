package main

import (
	"maps"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// relicItem is a candidate in the armor class's relic subclass: a libram
// carrying the given effect text.
func relicItem(id int, name, effectText string) scored {
	return scored{candidate: candidate{
		ID: id, Name: name, EffectText: effectText, Slots: []string{"ranged"},
		ClassID: armorClassID, SubclassID: armorSublibramID,
	}}
}

// Ids from build 1.60.1.70009, read off the client's own item-effect rows.
const (
	kindlingStave        = 11750  // chance-on-hit-style proc (aura 42), blank spell text
	lordGeneralsSword    = 11817  // chance on hit, blank spell text, engine-implemented
	theUnstoppableForce  = 19323  // chance on hit: a 1 sec stun
	blackbladeOfShahram  = 12592  // chance on hit: summons Shahram, engine-implemented
	cracklingStaff       = 19102  // equip spell power, folded into the item's stats
	ravencrestsLegacy    = 21520  // stats only: no client effect row at all
	shadowsongsSorrow    = 21522  // stats only: no client effect row at all
	riphook              = 12653  // equip attack power, folded into the item's stats
	libramOfLaw          = 272435 // equip modifier, modelled by the fork's relic work
	libramOfInvocation   = 249442 // equip modifier, modelled by the fork's relic work
	libramOfJudgementRne = 205420 // an engraving rune: no equip effect the engine models
)

func TestCarriesEffectReadsTheClientTablesWhenTheTextIsBlank(t *testing.T) {
	cases := []struct {
		name string
		c    candidate
		want bool
	}{
		{"effect text alone", candidate{ID: 1, EffectText: "Does a thing."}, true},
		{"a proc whose spell text is blank", candidate{ID: kindlingStave}, true},
		{"a blank-text chance-on-hit weapon the engine models", candidate{ID: lordGeneralsSword}, true},
		{"nothing at all", candidate{ID: 1}, false},
	}
	for _, tc := range cases {
		if got := carriesEffect(tc.c); got != tc.want {
			t.Errorf("%s: carriesEffect = %v, want %v", tc.name, got, tc.want)
		}
	}
}

// The brief named six weapon picks as unflagged procs. The client rows say
// which of them have one: Blackblade (modelled), Teebu's (modelled) and
// The Unstoppable Force carry a chance-on-hit spell; Shadowsong's Sorrow,
// Ravencrest's Legacy, Riphook and Crackling Staff do not (their only
// effects are stats the item data already holds), so there is nothing to
// flag on them.
func TestClientEffectTableMatchesWhatTheNamedWeaponsActuallyCarry(t *testing.T) {
	for id, want := range map[int]bool{
		blackbladeOfShahram: true, 1728: true, theUnstoppableForce: true, lordGeneralsSword: true, kindlingStave: true,
		shadowsongsSorrow: false, ravencrestsLegacy: false, riphook: false, cracklingStaff: false,
	} {
		if got := clientEffectItemIDs[id]; got != want {
			t.Errorf("clientEffectItemIDs[%d] = %v, want %v", id, got, want)
		}
	}
}

func TestAProcWithBlankTextIsRankedWhenModelledAndFlaggedWhenNot(t *testing.T) {
	modelled := candidate{ID: lordGeneralsSword}
	if !hasImplementedEffect(modelled) {
		t.Error("Lord General's Sword is engine-implemented with a blank spell text; hasImplementedEffect = false, want true so a sim ranks it")
	}
	// Kindling Stave was the unmodelled example until the engine registered
	// its Ignition proc; the engine's own answer now ranks it.
	if !hasImplementedEffect(candidate{ID: kindlingStave}) {
		t.Error("Kindling Stave's Ignition is registered in the engine; hasImplementedEffect = false, want true")
	}
	unmodelledID := 0
	for _, id := range slices.Sorted(maps.Keys(clientEffectItemIDs)) {
		if !effectImplemented(id) {
			unmodelledID = id
			break
		}
	}
	if unmodelledID == 0 {
		t.Fatal("every client-effect item is engine-modelled; the fixture needs another unmodelled example")
	}
	unmodelled := candidate{ID: unmodelledID}
	if hasImplementedEffect(unmodelled) {
		t.Errorf("item %d's effect is not in the engine; hasImplementedEffect = true, want false", unmodelledID)
	}
	if !trinketEffectUnmodelled(unmodelled) {
		t.Errorf("item %d carries an effect the engine does not simulate; the unmodelled flag = false, want true", unmodelledID)
	}
	if trinketEffectUnmodelled(candidate{ID: shadowsongsSorrow}) {
		t.Error("Shadowsong's Sorrow has no proc; the unmodelled flag = true, want false")
	}
}

func TestRelicShortlistKeepsOnlyRelicsASimCanVerify(t *testing.T) {
	pool := []scored{
		relicItem(libramOfLaw, "Libram of Law", "Increases the damage of your Judgement ability by 4%."),
		relicItem(libramOfJudgementRne, "Libram of Judgement", "Engrave your gloves with the Crusader Strike rune."),
		effectItem(3854, "Frost Tiger Blade", "text", "ranged"), // modelled, but not a relic
	}
	got := relicShortlist(pool)
	if len(got) != 1 || got[0].ID != libramOfLaw {
		t.Fatalf("relicShortlist = %+v, want only Libram of Law", got)
	}
}

// A relic scores nothing, so pick() chose an arbitrary one - here an
// unmodelled rune. The slot must still go to the modelled relic that
// simulates best, whatever the margin: rankSlotWithEffects' 1% rule would
// have kept the incumbent for any relic gaining less than that.
func TestRankSlotWithEffectsRanksRelicsByMeasuredGainNotByIncumbency(t *testing.T) {
	incumbent := relicItem(libramOfJudgementRne, "Libram of Judgement", "Engrave your gloves with the Crusader Strike rune.")
	picks := map[string]slotPick{"ranged": {Item: &incumbent}}
	bySlot := map[string][]scored{"ranged": {
		incumbent,
		relicItem(libramOfInvocation, "Libram of Invocation", "Reduces the mana cost of your Seal spells by 5%."),
		relicItem(libramOfLaw, "Libram of Law", "Increases the damage of your Judgement ability by 4%."),
	}}
	fake := &fakeEngine{
		DefaultDPS: 100,
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "ranged", ItemID: libramOfInvocation}}): 100.2,
			gearKey([]api.GearSlot{{Slot: "ranged", ItemID: libramOfLaw}}):        100.6,
		},
	}
	out, notes := rankSlotWithEffects(fake, specInfo{}, "human", "paladin", 60, "", picks, bySlot, "ranged")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	got := out["ranged"].Item
	if got == nil || got.ID != libramOfLaw {
		t.Fatalf("ranged pick = %+v, want Libram of Law (highest simulated DPS, 0.6 percent over an empty slot)", got)
	}
	if !got.GainMeasured || got.MeasuredGainDPS < 0.59 || got.MeasuredGainDPS > 0.61 {
		t.Errorf("gain = %v (measured %v), want 0.6 over the empty relic slot", got.MeasuredGainDPS, got.GainMeasured)
	}
	if out["ranged"].RunnerUp == nil || out["ranged"].RunnerUp.ID != libramOfInvocation {
		t.Errorf("runner-up = %+v, want Libram of Invocation", out["ranged"].RunnerUp)
	}
	for _, call := range fake.Calls {
		if call == gearKey([]api.GearSlot{{Slot: "ranged", ItemID: libramOfJudgementRne}}) {
			t.Errorf("an unmodelled relic was simmed: %v", fake.Calls)
		}
	}
}

func TestRankSlotWithEffectsLeavesARelicSlotWithNoModelledRelicAlone(t *testing.T) {
	incumbent := relicItem(libramOfJudgementRne, "Libram of Judgement", "Engrave your gloves with the Crusader Strike rune.")
	picks := map[string]slotPick{"ranged": {Item: &incumbent}}
	bySlot := map[string][]scored{"ranged": {incumbent}}
	out, _ := rankRelicSlot(&fakeEngine{DefaultDPS: 100}, specInfo{}, "human", "paladin", 60, "", picks, bySlot, "ranged")
	if out["ranged"].Item == nil || out["ranged"].Item.ID != libramOfJudgementRne {
		t.Fatalf("ranged pick = %+v, want the incumbent unchanged (nothing to measure)", out["ranged"].Item)
	}
}

// A relic that measures exactly the empty slot's DPS on the same seed
// touches nothing the rotation casts (Enhancement's Totem of Urgency, a
// Lesser Healing Wave relic): it is no pick.
func TestBuildReportEmptiesAModelledRelicWhoseGainIsNoise(t *testing.T) {
	picks := map[string]slotPick{
		"ranged": {Item: &scored{candidate: relicItem(libramOfLaw, "Libram of Law", "Increases the damage of your Judgement ability by 4%.").candidate, MeasuredDPS: 250, MeasuredGainDPS: 0, MeasuredGainStdErr: 0.4, GainMeasured: true}},
	}
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 250, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "ranged" {
			row = s
		}
	}
	if row.ItemID != 0 || row.EmptyReason != noDPSValueReason {
		t.Fatalf("ranged row = %+v, want empty with %q (the simulated gain is inside its own error)", row, noDPSValueReason)
	}
}

func TestBuildReportPublishesAModelledRelicWithARealGain(t *testing.T) {
	picks := map[string]slotPick{
		"ranged": {Item: &scored{candidate: relicItem(libramOfLaw, "Libram of Law", "Increases the damage of your Judgement ability by 4%.").candidate, MeasuredDPS: 250, MeasuredGainDPS: 1.5, MeasuredGainStdErr: 0.2, GainMeasured: true}},
	}
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 250, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "ranged" {
			row = s
		}
	}
	if row.ItemID != libramOfLaw || row.EmptyReason != "" || row.EffectUnmodelled {
		t.Fatalf("ranged row = %+v, want Libram of Law published, modelled", row)
	}
}

// relicRow is buildReport's ranged row for a lone relic pick.
func relicRow(t *testing.T, gain, stdErr float64) slotRow {
	t.Helper()
	picks := map[string]slotPick{
		"ranged": {Item: &scored{candidate: relicItem(libramOfLaw, "Libram of Law", "Increases the damage of your Judgement ability by 4%.").candidate, MeasuredDPS: 250, MeasuredGainDPS: gain, MeasuredGainStdErr: stdErr, GainMeasured: true}},
	}
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 250, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	for _, s := range r.Slots {
		if s.Slot == "ranged" {
			return s
		}
	}
	t.Fatal("no ranged row")
	return slotRow{}
}

// Round 19 (round 18 shaman S3): Enhancement's totems gain under a DPS
// (Totem of the Storm about +0.8, Totem of Thunder +0.24), which the 2x
// unpaired error of two 100-iteration runs (about 1.4 DPS) can never
// resolve, so the slot published empty although every Enhancement shaman
// wears a totem and a relic has no stats to lose. A relic whose measured
// gain moved off the empty slot's exact DPS (same seed: its effect reached a
// spell the rotation casts) and is not significantly negative now publishes
// as a near tie instead.
func TestBuildReportPublishesATouchingRelicInsideItsErrorAsANearTie(t *testing.T) {
	row := relicRow(t, 0.8, 0.7)
	if row.ItemID != libramOfLaw || row.EmptyReason != "" || !row.NearTie {
		t.Fatalf("ranged row = %+v, want Libram of Law published as a near tie", row)
	}
	if row.LowValue {
		t.Errorf("a relic near tie is not the weapon fallback: low_value set")
	}
	if worse := relicRow(t, -1.6, 0.7); worse.ItemID != 0 || worse.EmptyReason != noDPSValueReason {
		t.Errorf("a relic measurably worse than an empty slot = %+v, want empty", worse)
	}
	if real := relicRow(t, 1.5, 0.2); real.NearTie {
		t.Errorf("a relic with a real gain is no near tie: %+v", real)
	}
}

// A relic that touches nothing can top a noisy ordering by measuring exactly
// the empty slot while the one that touches a cast spell lands a hair below
// it (Totem of the Storm read -0.04 at one seed). The slot goes to the one
// that touches the rotation, as a near tie, not to the inert one that would
// then publish empty.
func TestRankRelicSlotPrefersATouchingRelicOverAnInertOneOnANearTie(t *testing.T) {
	inert := relicItem(libramOfInvocation, "Libram of Invocation", "Reduces the mana cost of your Seal spells by 5%.")
	touching := relicItem(libramOfLaw, "Libram of Law", "Increases the damage of your Judgement ability by 4%.")
	picks := map[string]slotPick{"ranged": {Item: &inert}}
	bySlot := map[string][]scored{"ranged": {inert, touching}}
	fake := &fakeEngine{
		DefaultDPS:    100,
		DefaultStdErr: 0.5,
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "ranged", ItemID: libramOfLaw}}): 99.96,
		},
	}
	out, _ := rankRelicSlot(fake, specInfo{}, "human", "paladin", 60, "", picks, bySlot, "ranged")
	got := out["ranged"].Item
	if got == nil || got.ID != libramOfLaw || !relicNearTie(got) {
		t.Fatalf("ranged pick = %+v, want Libram of Law as a near tie", got)
	}
	if out["ranged"].RunnerUp == nil || out["ranged"].RunnerUp.ID != libramOfInvocation {
		t.Errorf("runner-up = %+v, want the inert relic", out["ranged"].RunnerUp)
	}
}
