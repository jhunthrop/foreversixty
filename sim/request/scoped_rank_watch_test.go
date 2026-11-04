package request

import (
	"encoding/json"
	"fmt"
	"slices"
	"testing"
)

// Round 18 priest X4, the production side: a rank row scoped to some lines
// rewrites the authored id on those lines only. A condition on any OTHER line
// that reads the same authored id (a dotIsActive, a spellTimeToReady) would
// then watch a rank the scoped line no longer casts: a DoT gate that never
// sees its DoT, and the line recasting on every global. No such row may exist.
func TestScopedRankRowsAreNotWatchedOutsideTheirLines(t *testing.T) {
	all, err := embeddedRankChoices()
	if err != nil {
		t.Fatal(err)
	}
	for spec, choices := range all {
		raw, err := aplFS.ReadFile("apl/" + spec + ".apl.json")
		if err != nil {
			t.Fatalf("%s: %v", spec, err)
		}
		watchers, err := outOfScopeWatchers(raw, choices)
		if err != nil {
			t.Fatalf("%s: %v", spec, err)
		}
		for _, w := range watchers {
			t.Errorf("%s: %s", spec, w)
		}
	}
}

func TestOutOfScopeWatchersFindsAConditionOnAnotherLine(t *testing.T) {
	raw := []byte(`{"prepullActions":[],"priorityList":[
		{"action":{"castSpell":{"spellId":{"spellId":7}}}},
		{"action":{"condition":{"dotIsActive":{"spellId":{"spellId":7}}},"castSpell":{"spellId":{"spellId":8}}}},
		{"action":{"castSpell":{"spellId":{"spellId":7}}}}
	]}`)
	choices := rankChoices{7: {{FromLevel: 40, ToLevel: 49, Rank: 3, Lines: []int{0}}}}
	got, err := outOfScopeWatchers(raw, choices)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("want the condition on line 1 alone (line 2 casts the spell, which is allowed), got %v", got)
	}
}

// outOfScopeWatchers lists, for every rank row of choices scoped to lines or
// pre-pull steps, each line outside that scope that reads the row's authored
// id anywhere but as the thing it casts.
func outOfScopeWatchers(raw []byte, choices rankChoices) ([]string, error) {
	var root struct {
		PrepullActions []any `json:"prepullActions"`
		PriorityList   []any `json:"priorityList"`
	}
	if err := json.Unmarshal(raw, &root); err != nil {
		return nil, err
	}
	lists := map[bool][]any{true: root.PrepullActions, false: root.PriorityList}
	var out []string
	for id := range choices {
		scope := choices.scopedLines(id)
		if len(scope) == 0 {
			continue
		}
		for prepull, entries := range lists {
			key := map[bool]string{true: "prepullActions", false: "priorityList"}[prepull]
			for index, entry := range entries {
				if slices.Contains(scope, castLine{prepull: prepull, index: index}) {
					continue
				}
				if readsOutsideCast(entry, id, false) {
					out = append(out, fmt.Sprintf("%s line %d reads spell %d in a condition outside the rank row's scope %v", key, index, id, scope))
				}
			}
		}
	}
	slices.Sort(out)
	return out, nil
}

// readsOutsideCast reports whether node names id in an ActionID that is not
// under one of castKeys.
func readsOutsideCast(node any, id int32, underCast bool) bool {
	switch v := node.(type) {
	case map[string]any:
		if num, ok := v["spellId"].(float64); ok {
			return int32(num) == id && !underCast
		}
		for key, val := range v {
			if readsOutsideCast(val, id, underCast || castKeys[key]) {
				return true
			}
		}
	case []any:
		for _, item := range v {
			if readsOutsideCast(item, id, underCast) {
				return true
			}
		}
	}
	return false
}
