package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestToggleExecuteGateShapes(t *testing.T) {
	bare := map[string]any(isExecutePhaseCondition())
	other := map[string]any(dotIsActiveCondition(5, 0))

	if c, label, ok := toggleExecuteGate(nil); !ok || label != "add" || !isBareExecutePhase(c) {
		t.Errorf("nil condition: %v %q %v", c, label, ok)
	}
	if c, label, ok := toggleExecuteGate(bare); !ok || label != "remove" || c != nil {
		t.Errorf("bare gate: %v %q %v", c, label, ok)
	}
	for _, vals := range [][]any{{bare, other}, {other, bare}} {
		and := map[string]any{"and": map[string]any{"vals": vals}}
		c, label, ok := toggleExecuteGate(and)
		if !ok || label != "remove" || !reflect.DeepEqual(map[string]any(c), other) {
			t.Errorf("and with a gate: %v %q %v", c, label, ok)
		}
	}
	plain := map[string]any{"and": map[string]any{"vals": []any{other, other}}}
	c, label, ok := toggleExecuteGate(plain)
	if !ok || label != "add" || !hasExecutePhase(c) {
		t.Errorf("and without a gate must gain one: %v %q %v", c, label, ok)
	}
	for name, bad := range map[string]any{
		"non-map":     "x",
		"not an and":  other,
		"three vals":  map[string]any{"and": map[string]any{"vals": []any{other, other, other}}},
		"vals absent": map[string]any{"and": map[string]any{}},
	} {
		if _, _, ok := toggleExecuteGate(bad); ok {
			t.Errorf("%s: toggled although the gate is ambiguous", name)
		}
	}
}

func TestIsBareExecutePhaseRequiresTheOnlyKey(t *testing.T) {
	if isBareExecutePhase(map[string]any{"isExecutePhase": map[string]any{}, "x": 1}) || isBareExecutePhase("s") || isBareExecutePhase(map[string]any{}) {
		t.Fatal("non-bare accepted")
	}
}

func TestToggleExecuteGateMutationsOnlyTouchGatedCasts(t *testing.T) {
	rot := rotation{Type: "TypeAPL", PriorityList: []entry{
		{Action: action{"autocastOtherCooldowns": map[string]any{}, "condition": isExecutePhaseCondition()}}, // not a cast
		buildCastEntry("ungated", 1, 0, nil),
		buildCastEntry("gated", 2, 0, isExecutePhaseCondition()),
	}}
	muts := toggleExecuteGateMutations(rot, map[int]string{2: "Execute"})
	if len(muts) != 1 {
		t.Fatalf("got %d mutations", len(muts))
	}
	if !strings.Contains(muts[0].Label, "#3 (Execute): execute gate -> remove") {
		t.Fatal(muts[0].Label)
	}
	if _, has := muts[0].Rotation.PriorityList[2].Action["condition"]; has {
		t.Fatal("the gate was not removed")
	}
	assertLegalAPL(t, "toggle", muts[0].Rotation)
	if _, has := rot.PriorityList[2].Action["condition"]; !has {
		t.Fatal("the base rotation was edited")
	}
}

func TestHasExecutePhaseSearchesNestedSlices(t *testing.T) {
	deep := map[string]any{"and": map[string]any{"vals": []any{map[string]any{"not": map[string]any{"val": isExecutePhaseCondition()}}}}}
	if !hasExecutePhase(deep) || hasExecutePhase(map[string]any(dotIsActiveCondition(1, 0))) {
		t.Fatal("hasExecutePhase misjudged")
	}
}

func TestFindResourceGateAndConstVal(t *testing.T) {
	cond := map[string]any{"and": map[string]any{"vals": []any{
		map[string]any(dotIsActiveCondition(1, 0)),
		map[string]any(cmpCondition("OpGe", action{"currentRagePercent": map[string]any{}}, "40%")),
	}}}
	g, ok := findResourceGate(cond)
	if !ok || g != (resourceGate{LHSKey: "currentRagePercent", Op: "OpGe", RHS: "40%"}) {
		t.Fatalf("got %v %v", g, ok)
	}
	if _, ok := findResourceGate(map[string]any(dotIsActiveCondition(1, 0))); ok {
		t.Fatal("found a gate in a dot condition")
	}
	if constVal("x") != "" || constVal(map[string]any{"const": 3}) != "" || constVal(map[string]any{"const": map[string]any{"val": "7%"}}) != "7%" {
		t.Fatal("constVal misread")
	}
	newCond, replaced := replaceResourceGateValue(cond, g, "20%")
	if !replaced {
		t.Fatal("not replaced")
	}
	if g2, _ := findResourceGate(newCond); g2.RHS != "20%" {
		t.Fatalf("got %v", g2)
	}
	if g3, _ := findResourceGate(cond); g3.RHS != "40%" {
		t.Fatal("original edited")
	}
	if _, replaced := replaceResourceGateValue(cond, resourceGate{LHSKey: "currentRagePercent", Op: "OpLt", RHS: "40%"}, "0%"); replaced {
		t.Fatal("replaced a gate with a different operator")
	}
}

func TestRefreshHelpersPickDotOrAuraShapes(t *testing.T) {
	id := actionID{SpellID: 7, Rank: 1}
	if _, ok := remainActiveCondition("auraIsActive", id)["auraIsActive"]; !ok {
		t.Error("aura kind lost")
	}
	if _, ok := remainActiveCondition("dotIsActive", id)["dotIsActive"]; !ok {
		t.Error("dot kind lost")
	}
	if _, ok := remainingTimeValue("auraIsActive", id)["auraRemainingTime"]; !ok {
		t.Error("aura remaining lost")
	}
	if _, ok := remainingTimeValue("dotIsActive", id)["dotRemainingTime"]; !ok {
		t.Error("dot remaining lost")
	}
	if got := refreshVariants("dotIsActive", id, 0); len(got) != 2 {
		t.Errorf("unknown tick length must skip the one-tick variant, got %d", len(got))
	}
	if got := refreshVariants("dotIsActive", id, 2); len(got) != 3 {
		t.Errorf("got %d variants", len(got))
	}
}

func TestReplaceNotActiveRewritesOnlyTheMatchingNode(t *testing.T) {
	cond := map[string]any{"and": map[string]any{"vals": []any{
		map[string]any(notCondition(auraIsActiveCondition(8, 0))),
		map[string]any(notCondition(auraIsActiveCondition(9, 0))),
	}}}
	repl := action{"marker": true}
	out, changed := replaceNotActive(cond, "auraIsActive", actionID{SpellID: 9}, repl)
	if !changed {
		t.Fatal("no change")
	}
	vals := out.(map[string]any)["and"].(map[string]any)["vals"].([]any)
	if _, ok := vals[1].(map[string]any)["marker"]; !ok {
		t.Errorf("target not replaced: %v", vals[1])
	}
	if _, ok := vals[0].(map[string]any)["not"]; !ok {
		t.Errorf("other node touched: %v", vals[0])
	}
	if _, changed := replaceNotActive(cond, "auraIsActive", actionID{SpellID: 1}, repl); changed {
		t.Error("replaced a node that is not there")
	}
	if _, changed := replaceNotActive("x", "auraIsActive", actionID{SpellID: 9}, repl); changed {
		t.Error("scalar changed")
	}
}

func TestAllMutationsPoolsEveryOperator(t *testing.T) {
	muts := allMutations(testRotation(), testCandidates(), map[int]float64{100: 3}, nil)
	var swaps, removes, inserts, refresh, gates int
	for _, m := range muts {
		switch {
		case strings.HasPrefix(m.Label, "swap"):
			swaps++
		case strings.HasPrefix(m.Label, "remove"):
			removes++
		case strings.HasPrefix(m.Label, "insert"):
			inserts++
		case strings.Contains(m.Label, "refresh condition"):
			refresh++
		case strings.Contains(m.Label, "gate"):
			gates++
		}
	}
	if swaps == 0 || removes == 0 || inserts == 0 || refresh == 0 || gates == 0 {
		t.Fatalf("pool misses an operator: swap %d remove %d insert %d refresh %d gate %d", swaps, removes, inserts, refresh, gates)
	}
}

func TestAuthoredCastIDsAndAnyAuthored(t *testing.T) {
	rot := testRotation()
	rot.PrepullActions = []entry{buildCastEntry("", 777, 0, nil)}
	got := authoredCastIDs(rot)
	for _, id := range []int{100, 150, 200, 300, 777} {
		if !got[id] {
			t.Errorf("missing %d", id)
		}
	}
	if got[999] || !anyAuthored([]int{1, 100}, got) || anyAuthored([]int{1, 2}, got) {
		t.Error("anyAuthored misjudged")
	}
}
