package enginetalents

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/wowsims/classic/sim/core"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// repoRoot is the site repository root, from this package's own
// directory (sim/internal/enginetalents) - the same depth
// sim/cmd/talent-search's own data_test.go resolves from.
const repoRoot = "../../.."

// engineFieldRank reads one named field off a filled engine talents
// message - the int32 rank, or 1/0 for a bool (one-rank) field - the
// one shape Reposition's callers need to check regardless of which of
// those two kinds a given talent happens to compile to.
func engineFieldRank(m protoreflect.Message, field string) int {
	fd := m.Descriptor().Fields().ByName(protoreflect.Name(field))
	if fd == nil {
		return -1
	}
	v := m.Get(fd)
	if fd.Kind() == protoreflect.BoolKind {
		if v.Bool() {
			return 1
		}
		return 0
	}
	return int(v.Int())
}

// TestRepositionPutsRankedTalentsOnTheEnginesOwnIndex is the table
// test this lane's brief asks for: a known Ret Paladin and Elemental
// Shaman build, read from the real guides and the real active build's
// trees, must put every ranked talent on the COMPILED engine's own
// field - proto/<class>.proto's own "// node <id>" comment, read by
// ForClass - never on whatever index that talent happens to occupy in
// the active build's own (tier, column) order.
//
// Both cases are real, not synthetic: the engine's proto (built from
// client 1.60.1.69893) still carries Improved Holy Strike and Crusade
// in paladin's Holy and Retribution trees - both gone from the active
// build (1.60.1.70009) - which shifts every talent after them by one
// field; and the engine's shaman proto has Elemental Fury and
// Elemental Alacrity in each other's slot outright. A positional
// string bypassing Reposition would misread exactly the talents this
// test pins.
func TestRepositionPutsRankedTalentsOnTheEnginesOwnIndex(t *testing.T) {
	build, err := leveling.ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	dir := engineDir(t)

	cases := []struct {
		name, class, specSlug string
		treeIndex             int
		// wantField maps a stable talent id this build ranks to the
		// real engine field (proto/<class>.proto's own field name)
		// Reposition must place its rank on.
		wantField map[int]string
	}{
		{
			// data/curated/specs.json: paladin-retribution's own
			// tree_index.
			name:      "retribution paladin",
			class:     "paladin",
			specSlug:  "retribution",
			treeIndex: 2,
			wantField: map[int]string{
				105703: "conviction",        // unaffected row (idx4 both layouts) - still must resolve by id, not coast on position.
				110880: "instrument_of_law", // active idx15, engine idx16 - Crusade's removal shifts this field by one.
				105692: "twist_of_light",    // active idx16, engine idx17 - same shift.
				105693: "vengeance",         // active idx12 unranked (0) in this build - Encode must still omit it, not leak a stray digit onto it.
			},
		},
		{
			// data/curated/specs.json: shaman-elemental's own
			// tree_index.
			name:      "elemental shaman",
			class:     "shaman",
			specSlug:  "elemental",
			treeIndex: 0,
			wantField: map[int]string{
				104766: "elemental_fury",     // active idx14, engine idx7 - the direct swap this lane's brief names.
				104765: "elemental_alacrity", // active idx7, engine idx14 - the other half of the same swap.
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			trees, err := leveling.LoadTalentTrees(repoRoot, build, tc.class)
			if err != nil {
				t.Fatal(err)
			}
			layout, err := ForClass(dir, tc.class)
			if err != nil {
				t.Fatal(err)
			}
			guideBuild, digits, err := leveling.GuideBuildTalents(repoRoot, tc.class, tc.specSlug)
			if err != nil {
				t.Fatal(err)
			}
			if err := leveling.RequireGuideBuildMatchesActive(guideBuild, build); err != nil {
				t.Fatalf("fixture drifted - %v", err)
			}
			targets := leveling.GuideTalentTargets(trees, digits)
			site := leveling.LadderTalentString(trees, targets, tc.treeIndex, 60)

			engineString, err := layout.Reposition(trees, site)
			if err != nil {
				t.Fatalf("Reposition: %v", err)
			}

			cl := classLayouts()[tc.class]
			core.FillTalentsProto(cl.msg.ProtoReflect(), engineString, cl.sizes)

			for id, field := range tc.wantField {
				want := targets[id]
				if got := engineFieldRank(cl.msg.ProtoReflect(), field); got != want {
					t.Errorf("talent %d: engine field %q = %d, want %d (the guide's own target)", id, field, got, want)
				}
			}
		})
	}
}

// TestRepositionBypassGuardWouldHaveCaughtTheRealDrift is the guard
// this lane's brief asks for: it fails if a class whose site layout
// and compiled-engine layout genuinely differ - paladin and shaman,
// verified below by direct comparison - were ever read by a code path
// that skips Reposition and hands the engine the positional string
// unconverted. It is not a test of Reposition itself (the table test
// above already pins that); it is a test that the DIVERGENCE those two
// classes carry today is real, so a future regression that quietly
// drops a Reposition call cannot pass by accident just because the
// drift it was protecting against happened to close.
func TestRepositionBypassGuardWouldHaveCaughtTheRealDrift(t *testing.T) {
	build, err := leveling.ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	dir := engineDir(t)

	for _, tc := range []struct {
		class, specSlug string
		treeIndex       int
	}{
		{"paladin", "retribution", 2},
		{"shaman", "elemental", 0},
	} {
		t.Run(tc.class, func(t *testing.T) {
			trees, err := leveling.LoadTalentTrees(repoRoot, build, tc.class)
			if err != nil {
				t.Fatal(err)
			}
			layout, err := ForClass(dir, tc.class)
			if err != nil {
				t.Fatal(err)
			}
			guideBuild, digits, err := leveling.GuideBuildTalents(repoRoot, tc.class, tc.specSlug)
			if err != nil {
				t.Fatal(err)
			}
			if err := leveling.RequireGuideBuildMatchesActive(guideBuild, build); err != nil {
				t.Fatalf("fixture drifted - %v", err)
			}
			targets := leveling.GuideTalentTargets(trees, digits)
			site := leveling.LadderTalentString(trees, targets, tc.treeIndex, 60)

			engineString, err := layout.Reposition(trees, site)
			if err != nil {
				t.Fatal(err)
			}
			if engineString == site {
				t.Fatalf("%s: the site and engine talent strings are identical (%q) - this class no longer "+
					"exercises the drift this guard exists to catch; replace it with a class/spec that still does",
					tc.class, site)
			}
		})
	}
}
