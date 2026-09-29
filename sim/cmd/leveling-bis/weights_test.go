package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// TestEffectiveWeightsZeroesAnInsignificantNegativeWeight is the exact
// bug night-bis-sanity found: a weight this command's own significance
// bar (isWeightSignificant) cannot tell apart from noise, and which
// happens to be small and negative, must not survive into score()'s
// ranking at all -- not even at its own (negative) value.
func TestEffectiveWeightsZeroesAnInsignificantNegativeWeight(t *testing.T) {
	raw := map[string]api.StatWeight{
		// error (0.02) is 40% of |weight| (0.05), over the 25% bar:
		// insignificant.
		"spirit": {Stat: "spirit", Weight: -0.05, Error: 0.02},
	}
	got := effectiveWeights(raw)
	if got["spirit"] != 0 {
		t.Errorf(`effectiveWeights(%+v)["spirit"] = %v, want 0 (insignificant weight must not move a ranking)`, raw, got["spirit"])
	}

	c := candidate{Stats: map[string]float64{"spirit": 40}}
	if s := score(c, "wrist", got); s != 0 {
		t.Errorf("score() with the item's only stat carrying an insignificant negative weight = %v, want 0, not negative", s)
	}
}

// TestEffectiveWeightsKeepsASignificantWeightUnchanged is the other
// half of the same rule: a weight isWeightSignificant trusts must
// reach score() at its own real value, unrounded and unscaled.
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
	if s := score(c, "wrist", got); s != want {
		t.Errorf("score() with a significant weight = %v, want %v", s, want)
	}
}

// TestEffectiveWeightsZeroesAZeroWeightToo pins isWeightSignificant's
// own "a weight of exactly zero is never significant" rule as seen
// through effectiveWeights: a hard-capped or unmoved stat (weight=0,
// error=0) contributes nothing, same as an insignificant one.
func TestEffectiveWeightsZeroesAZeroWeightToo(t *testing.T) {
	raw := map[string]api.StatWeight{
		"stamina": {Stat: "stamina", Weight: 0, Error: 0},
	}
	got := effectiveWeights(raw)
	if got["stamina"] != 0 {
		t.Errorf(`effectiveWeights(%+v)["stamina"] = %v, want 0`, raw, got["stamina"])
	}
}

// TestBandPoolNeverPublishesANegativeScoreFromAnInsignificantWeight is
// an end-to-end pin of the same bug at buildBandPool's own level (the
// actual call site score() is reached from): an item whose only stat
// carries an insignificant negative weight must score 0, not
// negative, once buildBandPool has turned the raw weights result into
// what buildBandPool itself expects. This mirrors what main.go now
// does (effectiveWeights(wresult) before buildBandPool), rather than
// buildBandPool's own map[string]float64 parameter, since that
// parameter's contract (already-effective weights) does not change --
// only what main.go passes it does.
func TestBandPoolNeverPublishesANegativeScoreFromAnInsignificantWeight(t *testing.T) {
	raw := map[string]api.StatWeight{
		"spirit": {Stat: "spirit", Weight: -0.05, Error: 0.02},
	}
	weights := effectiveWeights(raw)

	items := []candidate{
		{ID: 1, Name: "Evergreen Gloves", RequiredLevel: 5, EffectiveRequiredLevel: 5, Stats: map[string]float64{"spirit": 40}, Slots: []string{"hands"}},
	}
	idx := lootIndex{1: {{Kind: "quest", Label: "A Quest"}}}
	pool := buildBandPool(items, idx, "mage", 10, "alliance", weights)
	if len(pool.Scored) != 1 {
		t.Fatalf("len(pool.Scored) = %d, want 1", len(pool.Scored))
	}
	if pool.Scored[0].Score != 0 {
		t.Errorf("Score = %v, want 0 (spirit's weight is insignificant), not negative", pool.Scored[0].Score)
	}
}
