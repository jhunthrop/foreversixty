package request

import (
	"encoding/json"
	"strconv"
	"testing"
)

// combatSliceAndDiceID is the rank Combat's rotation authors Slice and Dice at.
const combatSliceAndDiceID = 6774

// Sim-accuracy round 26: Combat cast Slice and Dice at one or two combo
// points every time it dropped, so below 19 it never reached a five-point
// Eviscerate. On the bare fights (10,000 iterations, seeds 1 and 50001, both
// factions) skipping Slice and Dice and spending every five points on
// Eviscerate does 0.5 to 6% more at every level from 10 to 60; only the raid
// fight at 60 keeps it (+2.4%). The line now waits for level 60.
func TestCombatSliceAndDiceWaitsForSixty(t *testing.T) {
	raw, err := aplFS.ReadFile("apl/rogue-combat.apl.json")
	if err != nil {
		t.Fatal(err)
	}
	var apl struct {
		PriorityList []struct {
			Action struct {
				CastSpell *struct {
					SpellID struct {
						SpellID int32 `json:"spellId"`
					} `json:"spellId"`
				} `json:"castSpell"`
				Condition json.RawMessage `json:"condition"`
			} `json:"action"`
		} `json:"priorityList"`
	}
	if err := json.Unmarshal(raw, &apl); err != nil {
		t.Fatal(err)
	}
	found := false
	for _, e := range apl.PriorityList {
		cast := e.Action.CastSpell
		if cast == nil || cast.SpellID.SpellID != combatSliceAndDiceID {
			continue
		}
		found = true
		var cond any
		if err := json.Unmarshal(e.Action.Condition, &cond); err != nil {
			t.Fatalf("Slice and Dice's condition: %v", err)
		}
		if !hasLevelFloor(cond, 60) {
			t.Error("Slice and Dice is cast below level 60")
		}
	}
	if !found {
		t.Fatalf("the rotation casts no Slice and Dice %d", combatSliceAndDiceID)
	}
}

// hasLevelFloor reports whether cond holds a top-level (or and-joined)
// "currentLevel >= floor" clause.
func hasLevelFloor(cond any, floor int) bool {
	m, ok := cond.(map[string]any)
	if !ok {
		return false
	}
	if and, ok := m["and"].(map[string]any); ok {
		vals, _ := and["vals"].([]any)
		for _, v := range vals {
			if hasLevelFloor(v, floor) {
				return true
			}
		}
		return false
	}
	cmp, ok := m["cmp"].(map[string]any)
	if !ok || cmp["op"] != "OpGe" {
		return false
	}
	lhs, _ := cmp["lhs"].(map[string]any)
	rhs, _ := cmp["rhs"].(map[string]any)
	if _, ok := lhs["currentLevel"]; !ok {
		return false
	}
	c, _ := rhs["const"].(map[string]any)
	s, _ := c["val"].(string)
	n, err := strconv.Atoi(s)
	return err == nil && n == floor
}
