package main

import (
	"encoding/json"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

// testRotation is a small synthetic rotation covering the shapes the
// mutators care about: an autocast pseudo-action, a dot-maintenance
// line (not(dotIsActive)), a damage-gate line (bare dotIsActive, which
// is NOT a maintenance line and must not be touched by the refresh
// mutator), a mana-gated filler, and a plain spam line.
func testRotation() rotation {
	return rotation{
		Type: "TypeAPL",
		PriorityList: []entry{
			{Action: action{"autocastOtherCooldowns": map[string]any{}}},
			{
				Notes:  "keep the dot up",
				Action: action{"castSpell": map[string]any{"spellId": buildActionID(100, 1)}, "condition": notCondition(dotIsActiveCondition(100, 1))},
			},
			{
				Notes:  "damage bonus while the dot is up - not a maintenance line",
				Action: action{"castSpell": map[string]any{"spellId": buildActionID(150, 2)}, "condition": dotIsActiveCondition(100, 1)},
			},
			{
				Action: action{"castSpell": map[string]any{"spellId": buildActionID(200, 0)}, "condition": cmpCondition("OpGe", action{"currentManaPercent": map[string]any{}}, "30%")},
			},
			{
				Action: action{"castSpell": map[string]any{"spellId": buildActionID(300, 5)}},
			},
		},
	}
}

func testCandidates() []learnedCandidate {
	dotCond, dotKind := action(notCondition(dotIsActiveCondition(400, 1))), condDOT
	return []learnedCandidate{
		{Name: "DotSpell", ID: 400, Rank: 1, Condition: dotCond, ConditionLabel: string(dotKind)},
		{Name: "PlainSpell", ID: 500, Rank: 0, Condition: nil, ConditionLabel: string(condPlain)},
	}
}

// assertLegalAPL round-trips r through the same JSON shape
// request.rotation() parses (sim/request/request.go) and then through
// protojson into the engine's own proto.APLRotation, failing the test
// if either step errors - "every mutation yields a legal APL".
func assertLegalAPL(t *testing.T, label string, r rotation) {
	t.Helper()
	b, err := r.marshalIndent()
	if err != nil {
		t.Fatalf("%s: marshaling: %v", label, err)
	}
	var round rotation
	if err := json.Unmarshal(b, &round); err != nil {
		t.Fatalf("%s: round-tripping through this package's own shape: %v", label, err)
	}
	apl := &proto.APLRotation{}
	if err := protojson.Unmarshal(b, apl); err != nil {
		t.Fatalf("%s: does not parse as an engine APLRotation: %v\n%s", label, err, b)
	}
}

// entryKey is an entry's content, for duplicate detection - "no
// duplicate actions".
func entryKey(e entry) string {
	b, _ := json.Marshal(e)
	return string(b)
}

func assertNoDuplicateActions(t *testing.T, label string, list []entry) {
	t.Helper()
	seen := map[string]bool{}
	for _, e := range list {
		k := entryKey(e)
		if seen[k] {
			t.Errorf("%s: duplicate action in the resulting priority list: %s", label, k)
		}
		seen[k] = true
	}
}

func TestSwapAdjacentMutationsAreLegalAndSameLength(t *testing.T) {
	base := testRotation()
	muts := swapAdjacentMutations(base, nil)
	if len(muts) != len(base.PriorityList)-1 {
		t.Fatalf("got %d mutations, want %d", len(muts), len(base.PriorityList)-1)
	}
	for _, m := range muts {
		assertLegalAPL(t, m.Label, m.Rotation)
		assertNoDuplicateActions(t, m.Label, m.Rotation.PriorityList)
		if len(m.Rotation.PriorityList) != len(base.PriorityList) {
			t.Errorf("%s: priority list length changed: got %d, want %d", m.Label, len(m.Rotation.PriorityList), len(base.PriorityList))
		}
	}
}

func TestRemoveActionMutationsAreLegalAndShorter(t *testing.T) {
	base := testRotation()
	muts := removeActionMutations(base, nil)
	if len(muts) != len(base.PriorityList) {
		t.Fatalf("got %d mutations, want %d", len(muts), len(base.PriorityList))
	}
	for _, m := range muts {
		assertLegalAPL(t, m.Label, m.Rotation)
		if len(m.Rotation.PriorityList) != len(base.PriorityList)-1 {
			t.Errorf("%s: priority list length = %d, want %d", m.Label, len(m.Rotation.PriorityList), len(base.PriorityList)-1)
		}
	}
}

// TestInsertCandidateMutationsGiveADotCandidateADotCondition is the
// task's own "a dot gets a dot condition" check: every insertion of
// the DOT-classified candidate must carry a dotIsActive-shaped
// condition on the inserted line, whatever position it lands at.
func TestInsertCandidateMutationsGiveADotCandidateADotCondition(t *testing.T) {
	base := testRotation()
	candidates := testCandidates()
	muts := insertCandidateMutations(base, candidates)
	want := len(candidates) * (len(base.PriorityList) + 1)
	if len(muts) != want {
		t.Fatalf("got %d mutations, want %d", len(muts), want)
	}
	dotInsertions := 0
	for _, m := range muts {
		assertLegalAPL(t, m.Label, m.Rotation)
		assertNoDuplicateActions(t, m.Label, m.Rotation.PriorityList)
		inserted, ok := findCastEntry(m.Rotation.PriorityList, 400)
		if !ok {
			continue
		}
		dotInsertions++
		if !hasDotIsActive(inserted.Action["condition"]) {
			t.Errorf("%s: inserted DotSpell's condition = %v, want a dotIsActive node", m.Label, inserted.Action["condition"])
		}
	}
	if dotInsertions != len(base.PriorityList)+1 {
		t.Fatalf("found the dot candidate inserted %d times, want %d (one per position)", dotInsertions, len(base.PriorityList)+1)
	}
}

func TestInsertCandidateMutationsGivePlainCandidateNoCondition(t *testing.T) {
	base := testRotation()
	muts := insertCandidateMutations(base, testCandidates())
	found := false
	for _, m := range muts {
		inserted, ok := findCastEntry(m.Rotation.PriorityList, 500)
		if !ok {
			continue
		}
		found = true
		if _, hasCond := inserted.Action["condition"]; hasCond {
			t.Errorf("%s: a plain candidate's inserted condition = %v, want none", m.Label, inserted.Action["condition"])
		}
	}
	if !found {
		t.Fatal("never found PlainSpell inserted")
	}
}

func findCastEntry(list []entry, id int) (entry, bool) {
	for _, e := range list {
		if aid, ok := castSpellID(e.Action); ok && aid.SpellID == id {
			return e, true
		}
	}
	return entry{}, false
}

func hasDotIsActive(node any) bool {
	m, ok := node.(map[string]any)
	if !ok {
		return false
	}
	if _, ok := m["dotIsActive"]; ok {
		return true
	}
	for _, v := range m {
		if mv, ok := v.(map[string]any); ok && hasDotIsActive(mv) {
			return true
		}
		if lst, ok := v.([]any); ok {
			for _, c := range lst {
				if hasDotIsActive(c) {
					return true
				}
			}
		}
	}
	return false
}

func TestRefreshConditionMutationsOnlyTouchMaintenanceLines(t *testing.T) {
	base := testRotation()
	ticks := map[int]float64{100: 1.5}
	muts := refreshConditionMutations(base, ticks, nil)
	// The dot-maintenance line (id 100, not(dotIsActive)) has two
	// variants to try (below-3s, below-one-tick); the damage-gate line
	// (id 150, bare dotIsActive, no "not") must never be touched.
	if len(muts) != 2 {
		t.Fatalf("got %d mutations, want 2 (below-3s and below-one-tick for the one maintenance line)", len(muts))
	}
	for _, m := range muts {
		assertLegalAPL(t, m.Label, m.Rotation)
		changed, ok := findCastEntry(m.Rotation.PriorityList, 100)
		if !ok {
			t.Fatalf("%s: lost the maintenance line entirely", m.Label)
		}
		if hasDotIsActive(changed.Action["condition"]) {
			t.Errorf("%s: still has a dotIsActive condition, want it rewritten to dotRemainingTime", m.Label)
		}
		unchanged, ok := findCastEntry(m.Rotation.PriorityList, 150)
		if !ok || !hasDotIsActive(unchanged.Action["condition"]) {
			t.Errorf("%s: the damage-gate line (bare dotIsActive) was altered", m.Label)
		}
	}
}

func TestResourceGateMutationsVaryTheThreshold(t *testing.T) {
	base := testRotation()
	muts := resourceGateMutations(base, nil)
	// 5 thresholds minus the one already in use (30%) = 4.
	if len(muts) != len(gateThresholds)-1 {
		t.Fatalf("got %d mutations, want %d", len(muts), len(gateThresholds)-1)
	}
	seen := map[string]bool{}
	for _, m := range muts {
		assertLegalAPL(t, m.Label, m.Rotation)
		e, ok := findCastEntry(m.Rotation.PriorityList, 200)
		if !ok {
			t.Fatalf("%s: lost the mana-gated line", m.Label)
		}
		g, found := findResourceGate(e.Action["condition"])
		if !found {
			t.Fatalf("%s: lost the resource gate entirely", m.Label)
		}
		if g.RHS == "30%" {
			t.Errorf("%s: threshold unchanged (still 30%%)", m.Label)
		}
		seen[g.RHS] = true
	}
	if len(seen) != len(gateThresholds)-1 {
		t.Errorf("got %d distinct thresholds, want %d", len(seen), len(gateThresholds)-1)
	}
}
