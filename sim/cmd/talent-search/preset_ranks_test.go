package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/request"
)

// A search's request names the preset it layers, so the rotation casts that
// preset's rank choices (data/curated/rank-choices.json). Without it a level-60
// request falls back to the raid preset and the bare-60 rows (Shadow's Devouring
// Plague rank 1, Elemental's Lightning Bolt rank 4, ...) never reach a bare-60
// search, whose baseline then differs from the published bare-60 number.
func TestSearchRequestNamesItsPresetForTheRankChoices(t *testing.T) {
	bare := simSetup{spec: specInfo{Spec: "priest-shadow", ClassSlug: "priest"}, band: bisBand{Race: "undead"}, level: 60, seed: 1}
	if got := bare.request("0", 10).Preset; got != request.BarePreset {
		t.Errorf("bare search preset = %q, want %q", got, request.BarePreset)
	}
	raid := bare
	raid.preset = &request.ResolvedPreset{}
	if got := raid.request("0", 10).Preset; got != request.RaidPreset {
		t.Errorf("raid search preset = %q, want %q", got, request.RaidPreset)
	}
}
