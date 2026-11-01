package request

import (
	"bytes"
	"os"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

const publishedPresetsPath = "published-presets.json"

// The /sim page's default preset and its bare set are this file (owner ruling, round 6). A
// stale copy is a failing test here, not a page number that quietly differs from the
// published one.
func TestPublishedPresetsJSONIsCommitted(t *testing.T) {
	want, err := PublishedPresetsJSON()
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(publishedPresetsPath)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bytes.TrimSpace(got), bytes.TrimSpace(want)) {
		t.Error("sim/request/published-presets.json is stale; run `go run ./internal/genraidconsumes` from sim/")
	}
}

// Every band below 60 publishes bare and 60 publishes raid, the ranker's own passes.
func TestPublishedPresetFollowsTheRaidBand(t *testing.T) {
	for _, tc := range []struct {
		level int
		want  string
	}{{1, BarePreset}, {20, BarePreset}, {40, BarePreset}, {59, BarePreset}, {60, RaidPreset}} {
		if got := PublishedPreset(tc.level); got != tc.want {
			t.Errorf("PublishedPreset(%d) = %q, want %q", tc.level, got, tc.want)
		}
	}
	want := []PresetFrom{{FromLevel: 1, Preset: BarePreset}, {FromLevel: RaidPresetBand, Preset: RaidPreset}}
	if got := PublishedPresets(); !slices.Equal(got, want) {
		t.Errorf("PublishedPresets() = %v, want %v", got, want)
	}
}

// The bare set at every level is the class kit, what the ranker's bare pass sends.
func TestBareSetsAreTheKitAtEveryLevel(t *testing.T) {
	sets := BareSets()
	if len(sets) != len(specs.All) {
		t.Fatalf("%d specs, want %d", len(sets), len(specs.All))
	}
	for _, spec := range specs.All {
		rows := sets[spec.Spec]
		if len(rows) == 0 || rows[0].FromLevel != 1 {
			t.Fatalf("%s: rows %v do not start at level 1", spec.Spec, rows)
		}
		for level := 1; level <= api.MaxLevel; level++ {
			at := SetAtLevel(rows, level)
			if !slices.Equal(at.Buffs, nonNil(leveling.KitBuffs(spec.Spec, level))) ||
				!slices.Equal(at.Consumes, nonNil(leveling.KitConsumes(spec.Spec, level))) {
				t.Errorf("%s level %d: bare set %+v is not the kit", spec.Spec, level, at)
			}
		}
	}
}
