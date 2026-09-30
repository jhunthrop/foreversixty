package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// TestEffectiveWeightsZeroesANegativeWeightRegardlessOfSignificance is
// night-bis-sanity's own bug: a negative weight -- in this engine, no
// stat ever lowers a damage spec's own DPS, so a negative weight is
// measurement noise around a true value at or near zero -- must not
// survive into score()'s ranking at its own negative value, whether or
// not isWeightSignificant would call it significant (this one would
// not: error 0.02 is 40% of |weight| 0.05, over the 25% bar).
func TestEffectiveWeightsZeroesANegativeWeightRegardlessOfSignificance(t *testing.T) {
	raw := map[string]api.StatWeight{
		"spirit": {Stat: "spirit", Weight: -0.05, Error: 0.02},
	}
	got := effectiveWeights(raw)
	if got["spirit"] != 0 {
		t.Errorf(`effectiveWeights(%+v)["spirit"] = %v, want 0 (a negative weight is noise, never a real "this stat hurts" signal)`, raw, got["spirit"])
	}

	c := candidate{Stats: map[string]float64{"spirit": 40}}
	if s := score(c, "wrist", got, 0, false); s != 0 {
		t.Errorf("score() with the item's only stat carrying a negative weight = %v, want 0, not negative", s)
	}
}

// TestEffectiveWeightsKeepsAPositiveWeightEvenWhenInsignificant is the
// controller's own correction (2026-09-29): the first version of this
// helper zeroed every weight isWeightSignificant called insignificant,
// not just negative ones, and a before/after run of mage-frost -bands
// 20 showed that DROPS real signal -- verified set_dps fell for both
// factions (alliance 29.8->29.2, horde 28.1->27.0) versus the
// un-zeroed baseline once insignificant-but-positive weights stopped
// counting. This command's 100-iteration sweep flags a real, useful
// positive weight "insignificant" more often than a ranking decision
// should discard it over; report.go's Insignificant flag is calibrated
// for "should a player trust this printed number", not "should this
// stat move a ranking". So a positive weight is kept at its own value
// here even when the sweep cannot pin its error under the 25% bar.
func TestEffectiveWeightsKeepsAPositiveWeightEvenWhenInsignificant(t *testing.T) {
	raw := map[string]api.StatWeight{
		// error (6.0) is over 25% of |weight| (14.87): insignificant by
		// isWeightSignificant, but positive, so effectiveWeights keeps it.
		"melee_haste": {Stat: "melee_haste", Weight: 14.87, Error: 6.0},
	}
	if isWeightSignificant(raw["melee_haste"]) {
		t.Fatal("test fixture's own premise is wrong: melee_haste should be insignificant by isWeightSignificant")
	}
	got := effectiveWeights(raw)
	if got["melee_haste"] != 14.87 {
		t.Errorf(`effectiveWeights(%+v)["melee_haste"] = %v, want 14.87 unchanged (positive, kept regardless of significance)`, raw, got["melee_haste"])
	}
}

// TestEffectiveWeightsKeepsASignificantWeightUnchanged is the
// uncontroversial case both this rule and the rejected one agree on: a
// weight isWeightSignificant trusts reaches score() at its own real
// value, unrounded and unscaled.
func TestEffectiveWeightsKeepsASignificantWeightUnchanged(t *testing.T) {
	raw := map[string]api.StatWeight{
		// error (0.1) is well under 25% of |weight| (2.05): significant.
		"agility": {Stat: "agility", Weight: 2.05, Error: 0.1},
	}
	got := effectiveWeights(raw)
	if got["agility"] != 2.05 {
		t.Errorf(`effectiveWeights(%+v)["agility"] = %v, want 2.05 unchanged`, raw, got["agility"])
	}

	c := candidate{Stats: map[string]float64{"agility": 10}}
	want := 10 * 2.05
	if s := score(c, "wrist", got, 0, false); s != want {
		t.Errorf("score() with a significant weight = %v, want %v", s, want)
	}
}

// TestEffectiveWeightsZeroesAZeroWeightToo pins the boundary: a
// hard-capped or unmoved stat (weight=0, error=0) is neither positive
// nor negative, and effectiveWeights' own "keep only Weight > 0" rule
// contributes nothing for it, same as a negative weight.
func TestEffectiveWeightsZeroesAZeroWeightToo(t *testing.T) {
	raw := map[string]api.StatWeight{
		"stamina": {Stat: "stamina", Weight: 0, Error: 0},
	}
	got := effectiveWeights(raw)
	if got["stamina"] != 0 {
		t.Errorf(`effectiveWeights(%+v)["stamina"] = %v, want 0`, raw, got["stamina"])
	}
}

// TestReferenceMeasurementReasonCatchesANegativeReference pins this
// lane's brief, item 2: warlock-demonology band 60 published
// reference_dps_per_point -0.189 (spell_power's own measured DPS
// delta went negative) with every OTHER weight sign-flipped
// (intellect 4.55, crit -1.77, hit -5.57), none marked insignificant.
// wresult[reference].Error is sim/adapter.Weights' own errAmt/scale;
// this test's error (-0.62) times its referenceDPSPerPoint (-0.189)
// recovers a positive raw standard error (0.117), and -0.189 does not
// clear it, so this band must not be trusted.
func TestReferenceMeasurementReasonCatchesANegativeReference(t *testing.T) {
	wresult := map[string]api.StatWeight{
		"spell_power": {Stat: "spell_power", Weight: 1, Error: -0.6199033936015775},
	}
	reason := referenceMeasurementReason("spell_power", wresult, -0.1893283096184771)
	if reason == "" {
		t.Fatal("referenceMeasurementReason(...) = \"\", want a non-empty reason for a negative reference delta")
	}
}

// TestReferenceMeasurementReasonCatchesAReferenceWithinItsOwnError is
// the boundary the brief's own wording draws: "positive beyond its own
// error" is stricter than merely positive - a small positive delta
// that does not clear its own standard error is exactly as
// untrustworthy as a negative one, since the "true" value could
// plausibly be zero or negative.
func TestReferenceMeasurementReasonCatchesAReferenceWithinItsOwnError(t *testing.T) {
	// error 1.5 * scale/delta 0.1 recovers raw stderr = 0.15, which the
	// reference delta itself (0.1) does not clear.
	wresult := map[string]api.StatWeight{
		"attack_power": {Stat: "attack_power", Weight: 1, Error: 1.5},
	}
	reason := referenceMeasurementReason("attack_power", wresult, 0.1)
	if reason == "" {
		t.Fatal("referenceMeasurementReason(...) = \"\", want a non-empty reason: 0.1 does not clear its own 0.15 error")
	}
}

// TestReferenceMeasurementReasonAllowsAConfidentlyPositiveReference is
// the ordinary case: a reference delta clearly above its own error
// (error 0.1/scale 0.07 recovers raw stderr 0.007, well under the
// 0.07 delta) must not be flagged - every band this lane measured but
// warlock-demonology 60 looks like this.
func TestReferenceMeasurementReasonAllowsAConfidentlyPositiveReference(t *testing.T) {
	wresult := map[string]api.StatWeight{
		"attack_power": {Stat: "attack_power", Weight: 1, Error: 0.1},
	}
	reason := referenceMeasurementReason("attack_power", wresult, 0.07)
	if reason != "" {
		t.Errorf("referenceMeasurementReason(...) = %q, want \"\" (0.07 clears its own 0.007 error)", reason)
	}
}

// TestBandPoolNeverPublishesANegativeScoreFromANegativeWeight is an
// end-to-end pin of the same bug at buildBandPool's own level (the
// actual call site score() is reached from): an item whose only stat
// carries a negative weight must score 0, not negative, once
// buildBandPool has turned the raw weights result into what
// buildBandPool itself expects. This mirrors what main.go now does
// (effectiveWeights(wresult) before buildBandPool), rather than
// buildBandPool's own map[string]float64 parameter, since that
// parameter's contract (already-effective weights) does not change --
// only what main.go passes it does.
func TestBandPoolNeverPublishesANegativeScoreFromANegativeWeight(t *testing.T) {
	raw := map[string]api.StatWeight{
		"spirit": {Stat: "spirit", Weight: -0.05, Error: 0.02},
	}
	weights := effectiveWeights(raw)

	items := []candidate{
		{ID: 1, Name: "Evergreen Gloves", RequiredLevel: 5, EffectiveRequiredLevel: 5, Stats: map[string]float64{"spirit": 40}, Slots: []string{"hands"}},
	}
	idx := lootIndex{1: {{Kind: "quest", Label: "A Quest"}}}
	pool := buildBandPool(items, idx, "mage", 10, "alliance", weights, 0, true)
	if len(pool.Scored) != 1 {
		t.Fatalf("len(pool.Scored) = %d, want 1", len(pool.Scored))
	}
	if pool.Scored[0].Score != 0 {
		t.Errorf("Score = %v, want 0 (spirit's weight is negative), not negative", pool.Scored[0].Score)
	}
}

// TestTrinketGainSignificantRequiresTwiceItsOwnError pins this lane's
// brief, item 2's own dogfooded numbers (weights.go's
// trinketGainSignificanceMultiplier doc): Fire Ruby on hunter-beast-
// mastery band 50 Alliance measured gain 0.773 against its own
// combined stdErr 0.7125 - clears a bare 1x bar but must not clear the
// 2x this gate actually uses, since it sits right next to Sanctified
// Orb's own (0.691/0.699) - two same-magnitude, effectively
// indistinguishable-from-noise readings for two trinkets with no real
// DPS relevance to a non-mana, non-mage spec.
func TestTrinketGainSignificantRequiresTwiceItsOwnError(t *testing.T) {
	if trinketGainSignificant(0.773, 0.7125) {
		t.Error("trinketGainSignificant(0.773, 0.7125) = true, want false (clears 1x but not 2x its own error - this lane's own Fire Ruby repro)")
	}
}

// TestTrinketGainSignificantAllowsARealSmallGain is the same
// dogfooded run's other half: Frozen Heart of the Mountain's own +9
// Hit measured gain 3.03 against stdErr 0.98 for rogue-assassination
// band 50 Horde - comfortably beyond 2x (1.97) - must stay a
// legitimate, verified pick under this rule (this lane's brief: "Frozen
// Heart of the Mountain at band 50... stays legitimate").
func TestTrinketGainSignificantAllowsARealSmallGain(t *testing.T) {
	if !trinketGainSignificant(3.03, 0.98) {
		t.Error("trinketGainSignificant(3.03, 0.98) = false, want true (comfortably beyond 2x its own error - Frozen Heart of the Mountain's own repro)")
	}
}

// TestTrinketGainSignificantRejectsAZeroNoiseZeroGain is the hard-
// capped boundary (matching isWeightSignificant's own >= convention,
// referenceMeasurementTrustworthy's own strict >): a gain of exactly 0
// with a stdErr of exactly 0 (an item simdb silently stripped, or one
// truly identical to the baseline every iteration) is not
// distinguishably positive and must not pass.
func TestTrinketGainSignificantRejectsAZeroNoiseZeroGain(t *testing.T) {
	if trinketGainSignificant(0, 0) {
		t.Error("trinketGainSignificant(0, 0) = true, want false")
	}
}
