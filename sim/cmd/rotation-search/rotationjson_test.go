package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParseRotationRoundTripsAndRejectsBadJSON(t *testing.T) {
	raw := json.RawMessage(`{"type":"TypeAPL","priorityList":[{"notes":"n","action":{"castSpell":{"spellId":{"spellId":5}}}}]}`)
	r, err := parseRotation(raw)
	if err != nil {
		t.Fatal(err)
	}
	if r.Type != "TypeAPL" || len(r.PriorityList) != 1 || r.PriorityList[0].Notes != "n" {
		t.Fatalf("got %+v", r)
	}
	if id, ok := castSpellID(r.PriorityList[0].Action); !ok || id.SpellID != 5 {
		t.Fatalf("castSpellID = %v %v", id, ok)
	}
	if _, err := parseRotation(json.RawMessage(`{"type":`)); err == nil {
		t.Fatal("malformed JSON parsed")
	}
}

func TestCloneIsDeepAndWithPriorityListKeepsPrepull(t *testing.T) {
	orig := rotation{
		Type:           "TypeAPL",
		PrepullActions: []entry{buildCastEntry("pre", 1, 0, nil)},
		PriorityList:   []entry{buildCastEntry("a", 2, 1, notCondition(dotIsActiveCondition(2, 1)))},
	}
	c := orig.clone()
	if !reflect.DeepEqual(orig, c) {
		t.Fatal("clone differs from the original")
	}
	c.PriorityList[0].Action["castSpell"].(map[string]any)["spellId"].(map[string]any)["spellId"] = float64(99)
	c.PrepullActions[0].Action["extra"] = true
	if id, _ := castSpellID(orig.PriorityList[0].Action); id.SpellID != 2 {
		t.Fatal("mutating the clone's priority list changed the original")
	}
	if _, ok := orig.PrepullActions[0].Action["extra"]; ok {
		t.Fatal("mutating the clone's prepull changed the original")
	}

	w := orig.withPriorityList(nil)
	if len(w.PriorityList) != 0 || len(w.PrepullActions) != 1 || w.Type != "TypeAPL" {
		t.Fatalf("withPriorityList = %+v", w)
	}
	w.PrepullActions[0].Action["extra"] = true
	if _, ok := orig.PrepullActions[0].Action["extra"]; ok {
		t.Fatal("withPriorityList aliased the prepull actions")
	}
	if (rotation{}).clone().PriorityList != nil {
		t.Fatal("cloning a nil list must stay nil")
	}
}

func TestCloneAnyCopiesNestedSlices(t *testing.T) {
	src := map[string]any{"vals": []any{map[string]any{"k": "v"}, 1.5, true}}
	dup := cloneAny(src).(map[string]any)
	dup["vals"].([]any)[0].(map[string]any)["k"] = "changed"
	if src["vals"].([]any)[0].(map[string]any)["k"] != "v" {
		t.Fatal("cloneAny shared a nested map")
	}
	if cloneAction(nil) != nil {
		t.Fatal("nil action must clone to nil")
	}
}

func TestReadActionIDRejectsMalformedNodes(t *testing.T) {
	if _, ok := readActionID("x"); ok {
		t.Error("string accepted")
	}
	if _, ok := readActionID(map[string]any{"rank": 1.0}); ok {
		t.Error("node without spellId accepted")
	}
	if id, ok := readActionID(buildActionID(9, 4)); !ok || id != (actionID{SpellID: 9, Rank: 4}) {
		t.Errorf("got %v %v", id, ok)
	}
	if _, ok := castSpellID(action{"autocastOtherCooldowns": map[string]any{}}); ok {
		t.Error("non-cast action reported a spell id")
	}
}

func TestFindBareActiveGatesSkipsNotWrappedMaintenance(t *testing.T) {
	cond := map[string]any{"and": map[string]any{"vals": []any{
		map[string]any(dotIsActiveCondition(10, 0)),
		map[string]any(auraIsActiveCondition(20, 1)),
		map[string]any(notCondition(dotIsActiveCondition(30, 0))),
	}}}
	var got []actionID
	findBareActiveGates(cond, &got)
	have := map[int]bool{}
	for _, g := range got {
		have[g.SpellID] = true
	}
	if !have[10] || !have[20] || have[30] || len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}
