package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

func fixtureGuideExpectation(t *testing.T) guideExpectation {
	t.Helper()
	root := writeRaidFixture(t, "build: 'FS1:testbuild:hunter:dwarf:2/50/0:'\nraidBuild: 'FS1:testbuild:hunter:dwarf:4/05/0:'\n")
	targets, err := loadPassTargets(root, "testbuild", raidFixtureSpec)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := leveling.LoadTalentTrees(root, "testbuild", "hunter")
	if err != nil {
		t.Fatal(err)
	}
	spec := raidFixtureSpec
	spec.Spec = "hunter-marksmanship"
	return guideExpectation{
		spec:    spec,
		races:   map[string]string{"alliance": "night-elf", "horde": "troll"},
		trees:   trees,
		targets: targets,
	}
}

func TestGuideDivergencesAcceptBandsRankedFromTheGuide(t *testing.T) {
	g := fixtureGuideExpectation(t)
	bands := []publishedBand{
		{Band: 60, Preset: presetBare, Faction: "alliance", Race: "night-elf", Talents: g.talentsFor(60, presetBare)},
		{Band: 60, Preset: presetRaid, Faction: "horde", Race: "troll", Talents: g.talentsFor(60, presetRaid)},
	}
	if got := guideDivergences(g, bands); len(got) != 0 {
		t.Errorf("bands ranked from the guide diverge: %v", got)
	}
	if g.talentsFor(60, presetRaid) == g.talentsFor(60, presetBare) {
		t.Fatal("fixture raid and leveling builds must differ for this test to mean anything")
	}
}

func TestGuideDivergencesNameAStaleRaceAndAStaleRaidBuild(t *testing.T) {
	g := fixtureGuideExpectation(t)
	stale := []publishedBand{
		{Band: 60, Preset: presetRaid, Faction: "alliance", Race: "dwarf", Talents: g.talentsFor(60, presetBare)},
	}
	got := guideDivergences(g, stale)
	if len(got) != 2 {
		t.Fatalf("want a race and a talents divergence, got %v", got)
	}
	if !strings.Contains(got[0], `race "dwarf"`) || !strings.Contains(got[0], `"night-elf"`) {
		t.Errorf("race line = %q", got[0])
	}
	if !strings.Contains(got[1], "raidBuild") {
		t.Errorf("talents line = %q, want it to name raidBuild", got[1])
	}
}

// knownStaleBISFiles are specs whose committed files predate their guide,
// each with the reason. It exists only while a split ranking lands one
// shard at a time; the run that lands the last shard empties it, and it
// never carries a spec past that landing.
var knownStaleBISFiles = map[string]string{
	"warrior-protection":   "sim-accuracy round 30 moved band 30's leveling talents (threat talents before Anticipation on the dungeon-table boss); the batch re-rank lands it",
	"warrior-arms":         "sim-accuracy 2026-11-12 (melee-hunters) moved the level 60 build (Spearing Strike for a point of Deflection) and with it the leveling talents from 17; the batch re-rank lands it",
	"hunter-beast-mastery": "sim-accuracy 2026-11-12 (melee-hunters) moved the level 60 build (Efficiency 5 and Improved Tracking 4 for Improved Stings, Rapid Killing and Mortal Shots) and the leveling talents from 46; the batch re-rank lands it",
	"mage-arcane":          "sim-accuracy 2026-11-12 (arcane) moved the level 60 leveling build (Arcane Wrath for the fourth Incineration point) and the leveling talents from 44; the batch re-rank lands it",
	"mage-frost":           "sim-accuracy 2026-11-12 (casters) moved the level 60 leveling build (Arcane Meditation 3 and Improved Channeling 5 for Incineration) and the leveling talents from 56; the batch re-rank lands it",
}

// TestPublishedBISMatchesGuides: the active build's committed bis files
// wear the races and talents their guides name. Not gated like
// TestPublishedBISFilesPassSanityChecks: a guide that moves its race or
// build must come with a re-rank (bis.yml, force), or this fails.
func TestPublishedBISMatchesGuides(t *testing.T) {
	activeBuild, err := leveling.ReadActiveBuild(publishedRepoRoot)
	if err != nil {
		t.Fatal(err)
	}
	buildDir := filepath.Join(publishedRepoRoot, "data", "builds", activeBuild)
	files, err := filepath.Glob(filepath.Join(buildDir, "bis", "*.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		if _, statErr := os.Stat(buildDir); statErr != nil {
			t.Skipf("data/builds/%s is not in this checkout", activeBuild)
		}
		t.Skipf("no bis files published for %s yet", activeBuild)
	}
	for _, path := range files {
		f, err := readPublishedSpecFile(path)
		if err != nil {
			t.Fatal(err)
		}
		t.Run(f.Spec, func(t *testing.T) {
			g, err := loadGuideExpectation(publishedRepoRoot, buildDir, activeBuild, f.Spec)
			if err != nil {
				t.Fatal(err)
			}
			if reason, stale := knownStaleBISFiles[f.Spec]; stale {
				t.Skipf("known stale: %s", reason)
			}
			for _, line := range guideDivergences(g, f.Bands) {
				t.Error(line)
			}
		})
	}
}
