package main

import (
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
	unmodelled := candidate{ID: kindlingStave}
	if hasImplementedEffect(unmodelled) {
		t.Error("Kindling Stave's proc is not in the engine; hasImplementedEffect = true, want false")
	}
	if !trinketEffectUnmodelled(unmodelled) {
		t.Error("Kindling Stave carries a proc the engine does not simulate; the unmodelled flag = false, want true")
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

// A relic is gated on its simulated gain exactly as a trinket is: one whose
// gain over an empty slot is noise is no pick.
func TestBuildReportEmptiesAModelledRelicWhoseGainIsNoise(t *testing.T) {
	picks := map[string]slotPick{
		"ranged": {Item: &scored{candidate: relicItem(libramOfLaw, "Libram of Law", "Increases the damage of your Judgement ability by 4%.").candidate, MeasuredDPS: 250, MeasuredGainDPS: 0.1, MeasuredGainStdErr: 0.4, GainMeasured: true}},
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
