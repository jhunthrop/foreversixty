package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// neutralSource is a lootIndex fixture whose one entry is always
// obtainable by both factions (Kind world_drop, no Side/Standing/
// Opens) - isFactionNeutralCandidate's own "every source obtainable by
// both factions" half, satisfied trivially so these tests can focus on
// the reconcile rule itself rather than on sourceFor's own gating.
func neutralSource(id int, label string) lootIndex {
	return lootIndex{id: {{Kind: "world_drop", Label: label}}}
}

func mergeLootIndex(idxs ...lootIndex) lootIndex {
	out := lootIndex{}
	for _, idx := range idxs {
		for id, srcs := range idx {
			out[id] = srcs
		}
	}
	return out
}

// measuredTrinket is trinket() (trinkets_test.go) plus the measured
// fields rankTrinketSlot's own tournament sets on its winner -
// reconcileFactionTrinkets' own real input shape (a slot's pick always
// carries these once it has been through a real tournament).
func measuredTrinket(id int, name string, gainDPS, gainStdErr, absoluteDPS float64) scored {
	s := trinket(id, name, 50)
	s.GainMeasured = true
	s.MeasuredGainDPS = gainDPS
	s.MeasuredGainStdErr = gainStdErr
	s.MeasuredDPS = absoluteDPS
	return s
}

// gearHasItem reports whether req's own gear equips itemID in slot -
// every fakeEngine.DPSFunc below reads this to tell "the candidate
// run" from "the no-trinket baseline run" apart, since swapSlot drops
// itemID 0 from the gear list entirely (verify.go's own doc) rather
// than equipping a literal zero id.
func gearHasItem(req api.SimRequest, slot string, itemID int) bool {
	for _, g := range req.Character.Gear {
		if g.Slot == slot && g.ItemID == itemID {
			return true
		}
	}
	return false
}

// (a) a faction-neutral trinket picked on Alliance, Horde's own slot
// empty -> Horde adopts it, with Horde's own measured numbers.
func TestReconcileFactionTrinketsFillsEmptyOtherFactionSlot(t *testing.T) {
	idx := neutralSource(201, "World drop")
	allianceItem := measuredTrinket(201, "Frozen Heart of the Mountain", 3.0, 0.5, 150.0)
	alliancePicks := map[string]slotPick{"trinket1": {Item: &allianceItem}}
	hordePicks := map[string]slotPick{"trinket1": {}}

	engine := &fakeEngine{
		DPSFunc: func(req api.SimRequest) (float64, error) {
			if gearHasItem(req, "trinket1", 201) {
				return 140.0, nil
			}
			return 136.0, nil
		},
	}
	alliance := factionTrinketInputs{Faction: "alliance", Race: "human"}
	horde := factionTrinketInputs{Faction: "horde", Race: "orc"}

	_, newHorde, notes := reconcileFactionTrinkets(engine, specInfo{ClassSlug: "hunter"}, "hunter", 50, "", idx, alliance, alliancePicks, horde, hordePicks)

	got := newHorde["trinket1"]
	if got.Item == nil || got.Item.ID != 201 {
		t.Fatalf("Horde trinket1 = %+v, want id 201 adopted from Alliance", got)
	}
	if !got.Item.GainMeasured || got.Item.MeasuredGainDPS <= 0 {
		t.Fatalf("Horde's adopted pick carries no real measured gain of its own: %+v", got.Item)
	}
	if got.Item.MeasuredDPS != 140.0 {
		t.Errorf("Horde's adopted pick MeasuredDPS = %v, want 140.0 (Horde's OWN measurement, never Alliance's)", got.Item.MeasuredDPS)
	}
	if len(notes) == 0 {
		t.Fatalf("expected a note describing the crossover")
	}
}

// (b) a faction-neutral trinket held on both sides, Horde's own pick
// measures lower - Horde swaps to Alliance's pick and its own former
// pick becomes the runner-up (so it survives as an alternative).
func TestReconcileFactionTrinketsSwapsALowerGainNeutralPick(t *testing.T) {
	idx := mergeLootIndex(neutralSource(201, "World drop A"), neutralSource(202, "World drop B"))
	allianceItem := measuredTrinket(201, "The Stronger Neutral Pick", 5.0, 0.5, 150.0)
	hordeItem := measuredTrinket(202, "Horde's Own Weaker Neutral Pick", 1.0, 0.2, 130.0)
	alliancePicks := map[string]slotPick{"trinket1": {Item: &allianceItem}}
	hordePicks := map[string]slotPick{"trinket1": {Item: &hordeItem}}

	engine := &fakeEngine{
		DPSFunc: func(req api.SimRequest) (float64, error) {
			if gearHasItem(req, "trinket1", 201) {
				return 145.0, nil // Horde's own measurement of item 201: gain 10.0, above Horde's own 1.0
			}
			return 135.0, nil
		},
	}
	alliance := factionTrinketInputs{Faction: "alliance", Race: "human"}
	horde := factionTrinketInputs{Faction: "horde", Race: "orc"}

	_, newHorde, notes := reconcileFactionTrinkets(engine, specInfo{ClassSlug: "hunter"}, "hunter", 50, "", idx, alliance, alliancePicks, horde, hordePicks)

	got := newHorde["trinket1"]
	if got.Item == nil || got.Item.ID != 201 {
		t.Fatalf("Horde trinket1 = %+v, want id 201 (swapped in from Alliance)", got)
	}
	if got.RunnerUp == nil || got.RunnerUp.ID != 202 {
		t.Fatalf("Horde trinket1 RunnerUp = %+v, want the old pick (id 202) demoted to it", got.RunnerUp)
	}
	if len(notes) == 0 {
		t.Fatalf("expected a note describing the swap")
	}
}

// (c) Horde measures Alliance's own neutral pick as a real, negative-
// beyond-error loss - a racial difference, not a defect - so Horde
// keeps its own verdict (here, empty) and the row carries a
// faction_note explaining why, rather than silently adopting a worse
// item or silently diverging with no record of the comparison.
func TestReconcileFactionTrinketsKeepsTargetWhenGainMeasuresNegativeBeyondError(t *testing.T) {
	idx := neutralSource(401, "World drop")
	allianceItem := measuredTrinket(401, "Fire Ruby", 2.0, 0.3, 150.0)
	alliancePicks := map[string]slotPick{"trinket1": {Item: &allianceItem}}
	hordePicks := map[string]slotPick{"trinket1": {}}

	engine := &fakeEngine{
		DPSFunc: func(req api.SimRequest) (float64, error) {
			if gearHasItem(req, "trinket1", 401) {
				return 120.0, nil // with the item
			}
			return 125.0, nil // no-trinket baseline is HIGHER: a real loss
		},
	}
	alliance := factionTrinketInputs{Faction: "alliance", Race: "human"}
	horde := factionTrinketInputs{Faction: "horde", Race: "orc"}

	_, newHorde, notes := reconcileFactionTrinkets(engine, specInfo{ClassSlug: "hunter"}, "hunter", 50, "", idx, alliance, alliancePicks, horde, hordePicks)

	got := newHorde["trinket1"]
	if got.Item != nil {
		t.Fatalf("Horde trinket1 = %+v, want it to stay empty (kept its own verdict)", got)
	}
	if !strings.Contains(got.FactionNote, "Fire Ruby") || !strings.Contains(got.FactionNote, "racial") {
		t.Fatalf("FactionNote = %q, want it to name Fire Ruby and call out a racial difference", got.FactionNote)
	}
	if len(notes) == 0 {
		t.Fatalf("expected a note describing the kept verdict")
	}
	// bis-ranker-integrity-16, item 3: this note names only the
	// REJECTED crossing candidate (Fire Ruby), never Item (which is
	// nil here) - that claim stays true no matter what later happens
	// to this slot's own pick, so report.go must never drop it.
	if got.FactionNoteNeedsPick {
		t.Fatalf("FactionNoteNeedsPick = true, want false: this note names the rejected candidate, not Item")
	}
}

// (c2) data-followups-10 lane, 2026-09-30: two DIFFERENT faction-
// neutral trinkets, each already its own faction's own tournament
// winner, whose gains on EACH OTHER's race are within their own
// combined standard error (a genuine cross-faction tie, the exact
// shape of the live Second Wind/Burst of Knowledge repro) must keep
// each faction's own verdict AND record a FactionNote explaining the
// comparison - never silently diverge with no note at all.
func TestReconcileFactionTrinketsNotesANearTieInsteadOfSilentlyDiverging(t *testing.T) {
	idx := mergeLootIndex(neutralSource(701, "Blackrock Depths: Ambassador Flamelash"), neutralSource(702, "Blackrock Depths: Golem Lord Argelmach"))
	allianceItem := measuredTrinket(701, "Burst of Knowledge", 5.7, 1.0, 150.0)
	hordeItem := measuredTrinket(702, "Second Wind", 5.5, 1.0, 148.0)
	alliancePicks := map[string]slotPick{"trinket1": {Item: &allianceItem}}
	hordePicks := map[string]slotPick{"trinket1": {Item: &hordeItem}}

	// Cross-measuring either item on the OTHER faction's race lands it
	// close to, but not exactly at, its own side's number - well inside
	// the two measurements' own combined error (sqrt(1.0^2+1.0^2) =
	// 1.41), the same way independent sim noise separated the live
	// repro's own two numbers by about 1.3 DPS.
	engine := &fakeEngine{
		DPSFunc: func(req api.SimRequest) (float64, error) {
			switch {
			case gearHasItem(req, "trinket1", 701):
				return 141.0, nil
			case gearHasItem(req, "trinket1", 702):
				return 140.8, nil
			default:
				return 135.0, nil
			}
		},
	}
	alliance := factionTrinketInputs{Faction: "alliance", Race: "human"}
	horde := factionTrinketInputs{Faction: "horde", Race: "orc"}

	newAlliance, newHorde, notes := reconcileFactionTrinkets(engine, specInfo{ClassSlug: "paladin"}, "paladin", 60, "", idx, alliance, alliancePicks, horde, hordePicks)

	if newAlliance["trinket1"].Item == nil || newAlliance["trinket1"].Item.ID != 701 {
		t.Fatalf("Alliance trinket1 = %+v, want it to keep its own pick (id 701)", newAlliance["trinket1"])
	}
	if newHorde["trinket1"].Item == nil || newHorde["trinket1"].Item.ID != 702 {
		t.Fatalf("Horde trinket1 = %+v, want it to keep its own pick (id 702)", newHorde["trinket1"])
	}
	if newAlliance["trinket1"].FactionNote == "" {
		t.Fatalf("Alliance trinket1 carries no FactionNote explaining the cross-faction tie")
	}
	if newHorde["trinket1"].FactionNote == "" {
		t.Fatalf("Horde trinket1 carries no FactionNote explaining the cross-faction tie")
	}
	if !strings.Contains(newAlliance["trinket1"].FactionNote, "indistinguishable") {
		t.Fatalf("Alliance FactionNote = %q, want it to call out the tie as statistically indistinguishable", newAlliance["trinket1"].FactionNote)
	}
	if !strings.Contains(newHorde["trinket1"].FactionNote, "indistinguishable") {
		t.Fatalf("Horde FactionNote = %q, want it to call out the tie as statistically indistinguishable", newHorde["trinket1"].FactionNote)
	}
	if len(notes) < 2 {
		t.Fatalf("notes = %v, want at least one per direction describing the kept tie", notes)
	}
	// bis-ranker-integrity-16, item 3: this note names Item ("this
	// faction's own pick") by name - a claim report.go's own
	// trinketLowGain gate can later falsify if Item's measured gain
	// does not clear significance. FactionNoteNeedsPick is how
	// report.go (publishableFactionNote) knows to drop the note rather
	// than publish it next to an empty row.
	if !newAlliance["trinket1"].FactionNoteNeedsPick {
		t.Fatalf("Alliance trinket1.FactionNoteNeedsPick = false, want true: the note names Item as its own pick")
	}
	if !newHorde["trinket1"].FactionNoteNeedsPick {
		t.Fatalf("Horde trinket1.FactionNoteNeedsPick = false, want true: the note names Item as its own pick")
	}
}

// (d) a faction-restricted trinket never crosses: it is never even
// offered as a candidate, so the other faction's own verdict (here,
// empty) is left completely untouched and the engine is never called
// at all.
func TestReconcileFactionTrinketsNeverCrossesAFactionRestrictedPick(t *testing.T) {
	idx := lootIndex{}
	allianceItem := measuredTrinket(501, "Alliance-Only Trinket", 5.0, 0.2, 150.0)
	allianceItem.FactionRestriction = "alliance"
	alliancePicks := map[string]slotPick{"trinket1": {Item: &allianceItem}}
	hordePicks := map[string]slotPick{"trinket1": {}}

	engine := &fakeEngine{}
	alliance := factionTrinketInputs{Faction: "alliance", Race: "human"}
	horde := factionTrinketInputs{Faction: "horde", Race: "orc"}

	_, newHorde, notes := reconcileFactionTrinkets(engine, specInfo{ClassSlug: "hunter"}, "hunter", 50, "", idx, alliance, alliancePicks, horde, hordePicks)

	if newHorde["trinket1"].Item != nil {
		t.Fatalf("Horde trinket1 = %+v, want it untouched (a faction-restricted pick must never cross)", newHorde["trinket1"])
	}
	if len(engine.Calls) != 0 {
		t.Fatalf("engine was called %d time(s), want 0 - a faction-restricted pick is never even offered as a candidate", len(engine.Calls))
	}
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
}

// (e) idempotence: once a slot has converged, a second full call
// changes nothing further - reconcileFactionTrinkets' own doc names
// this as the "are the two factions' picks already the same physical
// item" short-circuit every direction checks first.
func TestReconcileFactionTrinketsSecondRoundIsANoOp(t *testing.T) {
	idx := neutralSource(601, "World drop")
	allianceItem := measuredTrinket(601, "Frozen Heart of the Mountain", 4.0, 0.1, 150.0)
	alliancePicks := map[string]slotPick{"trinket1": {Item: &allianceItem}}
	hordePicks := map[string]slotPick{"trinket1": {}}

	engine := &fakeEngine{
		DPSFunc: func(req api.SimRequest) (float64, error) {
			if gearHasItem(req, "trinket1", 601) {
				return 140.0, nil
			}
			return 136.0, nil
		},
	}
	alliance := factionTrinketInputs{Faction: "alliance", Race: "human"}
	horde := factionTrinketInputs{Faction: "horde", Race: "orc"}

	firstAlliance, firstHorde, firstNotes := reconcileFactionTrinkets(engine, specInfo{ClassSlug: "hunter"}, "hunter", 50, "", idx, alliance, alliancePicks, horde, hordePicks)
	if len(firstNotes) == 0 {
		t.Fatalf("expected the first round to record a crossover")
	}

	secondAlliance, secondHorde, secondNotes := reconcileFactionTrinkets(engine, specInfo{ClassSlug: "hunter"}, "hunter", 50, "", idx, alliance, firstAlliance, horde, firstHorde)

	if secondHorde["trinket1"].Item == nil || secondHorde["trinket1"].Item.ID != firstHorde["trinket1"].Item.ID {
		t.Fatalf("second round changed Horde's own pick: first %+v, second %+v", firstHorde["trinket1"], secondHorde["trinket1"])
	}
	if secondAlliance["trinket1"].Item.ID != firstAlliance["trinket1"].Item.ID {
		t.Fatalf("second round changed Alliance's own pick: first %+v, second %+v", firstAlliance["trinket1"], secondAlliance["trinket1"])
	}
	if len(secondNotes) != 0 {
		t.Fatalf("second round produced notes %v, want none - a converged slot must short-circuit with no work at all", secondNotes)
	}
}

// isFactionNeutralCandidate's own three cases directly, since
// reconcileFactionTrinkets' own tests above only ever exercise it
// through a full reconcile pass.
func TestIsFactionNeutralCandidate(t *testing.T) {
	idx := mergeLootIndex(
		neutralSource(1, "World drop"),
		lootIndex{2: {{Kind: "rep", Label: "Silverwing Sentinels (honored)", Side: "alliance"}}},
	)
	cases := []struct {
		name string
		c    candidate
		want bool
	}{
		{"a faction-restricted item", candidate{ID: 1, FactionRestriction: "alliance"}, false},
		{"an unrestricted item sourced for only one faction", candidate{ID: 2}, false},
		{"an unrestricted item with a source obtainable by both", candidate{ID: 1}, true},
		{"an unrestricted item with no known source at all", candidate{ID: 999}, false},
	}
	for _, tc := range cases {
		if got := isFactionNeutralCandidate(tc.c, 60, idx); got != tc.want {
			t.Errorf("%s: isFactionNeutralCandidate = %v, want %v", tc.name, got, tc.want)
		}
	}
}
