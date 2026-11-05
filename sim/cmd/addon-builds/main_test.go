package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// TestCommittedAddonBuildsFollowTheGuides: the addon's default talent
// builds (Data.lua's defaultBuilds) come from the committed
// addon-builds.json, which only the client activation used to write. A
// guide build changed since (Fire's raid build took Arcane Meditation in
// round 4 and Improved Fire Blast 2/2 in round 5) left the in-game addon
// on the old one. Regenerate with `go run ./sim/cmd/addon-builds -repo-root .`.
func TestCommittedAddonBuildsFollowTheGuides(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	active, want, err := generate(root)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(filepath.Join(root, "data", "builds", active, "addon-builds.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Error("data/builds/" + active + "/addon-builds.json is stale against the guides' builds; regenerate it with go run ./sim/cmd/addon-builds -repo-root .")
	}
}

// TestRaidOrderEndsOnTheGuideRaidBuild: round 22's regen wrote the
// Assassination raid order as a copy of the leveling order, because the
// raid build ended in an empty third tree and the parser fell back to the
// leveling line. Every committed raid order must spend its points exactly
// as the guide's raidBuild says, and may equal the leveling order only when
// the guide gives both lines the same talents.
func TestRaidOrderEndsOnTheGuideRaidBuild(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	active, err := leveling.ReadActiveBuild(root)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join(root, "data", "builds", active, "addon-builds.json"))
	if err != nil {
		t.Fatal(err)
	}
	var committed map[string]map[string]build
	if err := json.Unmarshal(raw, &committed); err != nil {
		t.Fatal(err)
	}
	for _, spec := range specs.All {
		_, levelDigits, err := leveling.GuideBuildTalents(root, spec.ClassSlug, spec.SpecSlug)
		if err != nil {
			t.Fatalf("%s leveling build: %v", spec.Spec, err)
		}
		_, raidDigits, err := leveling.GuideRaidBuildTalents(root, spec.ClassSlug, spec.SpecSlug)
		if err != nil {
			t.Fatalf("%s raid build: %v", spec.Spec, err)
		}
		trees, err := leveling.LoadTalentTrees(root, active, spec.ClassSlug)
		if err != nil {
			t.Fatal(err)
		}
		got := map[[3]int]int{}
		for _, p := range committed[spec.Spec]["raid"].Order {
			got[p]++
		}
		want := map[[3]int]int{}
		targets := leveling.GuideTalentTargets(trees, raidDigits)
		for ti, tree := range trees {
			for _, node := range tree.Talents {
				if r := targets[node.ID]; r > 0 {
					want[[3]int{ti + 1, node.Tier + 1, node.Column + 1}] = r
				}
			}
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: the raid order does not end on the guide's raidBuild %v", spec.Spec, raidDigits)
		}
		if levelDigits != raidDigits && reflect.DeepEqual(committed[spec.Spec]["raid"].Order, committed[spec.Spec]["leveling"].Order) {
			t.Errorf("%s: the raid order is a copy of the leveling order, but the guide's two builds differ", spec.Spec)
		}
	}
}
