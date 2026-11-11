package request

import (
	"encoding/json"
	"strconv"
	"testing"
)

// Sim-accuracy round 30 (open items): the bear's Maul gate, retuned on round 29's leveling
// order and the dungeon-table boss. A bear queues Maul only once it has this much rage, so the
// rest goes into Lacerate and Primal Bite: 15 below 25, 20 from 25, 45 from 30 to 35, 25 from 36
// to 39 (where a higher gate costs over 5 points of chance of death once Frenzied Regeneration
// needs the rage), 45 from 40 to 59 and 25 at 60. Threat +0.4 to +6.3% at 25 to 59 and +0.5 to
// +1.8% at 60, both factions, seeds 1 and 50001 (design/reviews/2026-11-11-sim-accuracy/round30-open.md).
func TestBearMaulGateFollowsTheLevel(t *testing.T) {
	b, err := aplFS.ReadFile("apl/druid-feral-bear.apl.json")
	if err != nil {
		t.Fatal(err)
	}
	var apl struct {
		PriorityList []struct {
			Action map[string]json.RawMessage `json:"action"`
		} `json:"priorityList"`
	}
	if err := json.Unmarshal(b, &apl); err != nil {
		t.Fatal(err)
	}
	var cond json.RawMessage
	for _, line := range apl.PriorityList {
		var cast struct {
			SpellID struct {
				SpellID int `json:"spellId"`
			} `json:"spellId"`
		}
		if raw, ok := line.Action["castSpell"]; ok && json.Unmarshal(raw, &cast) == nil && cast.SpellID.SpellID == 9881 {
			cond = line.Action["condition"]
		}
	}
	if cond == nil {
		t.Fatal("no Maul line")
	}
	want := map[int]float64{10: 15, 24: 15, 25: 20, 29: 20, 30: 45, 35: 45, 36: 25, 39: 25, 40: 45, 53: 45, 54: 45, 59: 45, 60: 25}
	for level, rage := range want {
		if got := maulGate(t, cond, level); got != rage {
			t.Errorf("level %d: Maul queued from %v rage, want %v", level, got, rage)
		}
	}
}

// maulGate is the least rage at which cond holds for a character of level (the gate is
// monotonic in rage, so a scan finds it).
func maulGate(t *testing.T, cond json.RawMessage, level int) float64 {
	t.Helper()
	for rage := 0.0; rage <= 100; rage++ {
		if evalBearCond(t, cond, level, rage) {
			return rage
		}
	}
	return -1
}

func evalBearCond(t *testing.T, raw json.RawMessage, level int, rage float64) bool {
	t.Helper()
	var node map[string]json.RawMessage
	if err := json.Unmarshal(raw, &node); err != nil {
		t.Fatal(err)
	}
	var group struct {
		Vals []json.RawMessage `json:"vals"`
	}
	switch {
	case node["and"] != nil:
		_ = json.Unmarshal(node["and"], &group)
		for _, v := range group.Vals {
			if !evalBearCond(t, v, level, rage) {
				return false
			}
		}
		return true
	case node["or"] != nil:
		_ = json.Unmarshal(node["or"], &group)
		for _, v := range group.Vals {
			if evalBearCond(t, v, level, rage) {
				return true
			}
		}
		return false
	case node["cmp"] != nil:
		var c struct {
			Op  string                     `json:"op"`
			Lhs map[string]json.RawMessage `json:"lhs"`
			Rhs struct {
				Const struct {
					Val string `json:"val"`
				} `json:"const"`
			} `json:"rhs"`
		}
		if err := json.Unmarshal(node["cmp"], &c); err != nil {
			t.Fatal(err)
		}
		rhs, err := strconv.ParseFloat(c.Rhs.Const.Val, 64)
		if err != nil {
			t.Fatal(err)
		}
		var lhs float64
		switch {
		case c.Lhs["currentRage"] != nil:
			lhs = rage
		case c.Lhs["currentLevel"] != nil:
			lhs = float64(level)
		default:
			t.Fatalf("unexpected operand %v", c.Lhs)
		}
		switch c.Op {
		case "OpGe":
			return lhs >= rhs
		case "OpGt":
			return lhs > rhs
		case "OpLe":
			return lhs <= rhs
		case "OpLt":
			return lhs < rhs
		}
		t.Fatalf("unexpected op %s", c.Op)
	}
	t.Fatalf("unexpected node %s", raw)
	return false
}
