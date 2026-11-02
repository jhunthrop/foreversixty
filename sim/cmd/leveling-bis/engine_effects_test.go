package main

import "testing"

// Round 9 (items lane X1): the static list of engine-modelled items dated
// from fork 912a458f7 (2026-10-07), so seventeen picked items whose effects
// the engine runs were ranked as if they had none and published "not
// simulated". The ranker now asks the engine it links.
func TestTheEngineDecidesWhichItemEffectsAreModelled(t *testing.T) {
	for id, name := range map[int]string{
		1713:  "Ankh of Life",
		11750: "Kindling Stave",
		14152: "Robe of the Archmage",
		9453:  "Toxic Revenger",
		6831:  "Black Menace",
		11815: "Hand of Justice",
		3854:  "Frost Tiger Blade",
	} {
		if !effectImplemented(id) {
			t.Errorf("%s (%d): the engine registers its effect, effectImplemented = false", name, id)
		}
	}
	if effectImplemented(424242) {
		t.Error("424242 is no item; effectImplemented = true")
	}
}
