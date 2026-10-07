package main

import (
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// integrationRepoRoot is the real site repository root, from this
// package's own directory (sim/cmd/leveling-bis) - the same depth
// sim/cmd/talent-search's own repoRoot resolves from. The fixture-
// based tests elsewhere in this package (testdata/reporoot,
// identityTalentLayout) deliberately avoid the real repo; these two
// tests need it, to exercise a real class the site and engine layouts
// are known to differ on.
const integrationRepoRoot = "../../.."

// TestBandTalentStringsRoutesPaladinAndShamanThroughTheEngineLayout is
// the guard this lane's brief asks for: runSpec's own band loop calls
// bandTalentStrings, never leveling.LadderTalentString directly, to
// get the string a character is built from. For a real class whose
// site (active-build-positional) and compiled-engine layouts are
// known to differ - paladin and shaman; sim/internal/enginetalents'
// own doc - this proves bandTalentStrings still performs that
// conversion rather than something having quietly gone back to
// handing the engine the site string unconverted.
func TestBandTalentStringsRoutesPaladinAndShamanThroughTheEngineLayout(t *testing.T) {
	build, err := leveling.ReadActiveBuild(integrationRepoRoot)
	if err != nil {
		t.Fatal(err)
	}
	engineDir, err := enginetalents.SourceDir(filepath.Join(integrationRepoRoot, "sim"))
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		class, specSlug string
		treeIndex       int
	}{
		{"paladin", "retribution", 2},
		{"shaman", "elemental", 0},
	} {
		t.Run(tc.class, func(t *testing.T) {
			trees, err := leveling.LoadTalentTrees(integrationRepoRoot, build, tc.class)
			if err != nil {
				t.Fatal(err)
			}
			layout, err := enginetalents.ForClass(engineDir, tc.class)
			if err != nil {
				t.Fatal(err)
			}
			guideBuild, digits, err := leveling.GuideBuildTalents(integrationRepoRoot, tc.class, tc.specSlug)
			if err != nil {
				t.Fatal(err)
			}
			if err := leveling.RequireGuideBuildMatchesActive(guideBuild, build); err != nil {
				t.Fatalf("fixture drifted - %v", err)
			}
			targets := leveling.GuideTalentTargets(trees, digits)

			site, engine, err := bandTalentStrings(trees, targets, tc.treeIndex, 60, layout)
			if err != nil {
				t.Fatal(err)
			}
			// The engine proto was regenerated from the live trees on
			// 2026-10-07, so paladin and shaman no longer drift from the
			// site layout; the conversion itself is pinned by
			// TestBandTalentStringsUsesTheGivenLayoutNotTheSiteOrder.
			if site != engine {
				t.Fatalf("%s: the engine string %q differs from the site string %q - the engine proto lags "+
					"the site's active build", tc.class, engine, site)
			}
		})
	}
}

// TestBandTalentStringsUsesTheGivenLayoutNotTheSiteOrder is a direct
// unit check, independent of real engine/guide data: a layout whose
// Reposition swaps two talents must have that swap reflected in
// bandTalentStrings' own "engine" return, never silently dropped in
// favor of the site string it also computes.
func TestBandTalentStringsUsesTheGivenLayoutNotTheSiteOrder(t *testing.T) {
	trees := []leveling.TalentTree{{Talents: []leveling.TalentNode{{ID: 1, MaxRank: 5}, {ID: 2, MaxRank: 5}}}}
	targets := map[int]int{1: 3, 2: 1}

	site, engine, err := bandTalentStrings(trees, targets, 0, 60, swapFirstTwoLayout{})
	if err != nil {
		t.Fatal(err)
	}
	if site != "31" {
		t.Fatalf("site = %q, want 31", site)
	}
	if engine != "13" {
		t.Fatalf("engine = %q, want 13 (swapped by the given layout)", engine)
	}
}

// swapFirstTwoLayout is a talentLayout stand-in that reverses a
// two-digit positional string - enough to prove bandTalentStrings
// reads its "engine" value from the given layout's own Reposition,
// not from the site string it computed.
type swapFirstTwoLayout struct{}

func (swapFirstTwoLayout) Reposition(_ []leveling.TalentTree, s string) (string, error) {
	b := []byte(s)
	b[0], b[1] = b[1], b[0]
	return string(b), nil
}
