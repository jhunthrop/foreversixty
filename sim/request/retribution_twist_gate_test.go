package request

import (
	"encoding/json"
	"strings"
	"testing"
)

// Sim-accuracy round 21 (MythicSim comparison lead): the Unsealed Wrath twist
// lines (the Seal of Righteousness release, the switch back to Seal of Command
// and Seal of Command's upkeep) were gated on 1200 attack power and 30% mana.
// The attack power gate stood in for "raid buffs" and failed under other raid
// buff sets: without Heart of the Lion and Trueshot Aura the lines stopped
// firing and the raid build lost 14 to 15% (seeds 1 and 50001, both
// factions). Pacing the twist on the fight instead (mana percent at least four
// fifths of the fight left, the rule Consecration already uses) wins or ties
// in every buff state measured: raid 60 +0.3%, bare 60 +0.5 to +1.8%, no
// Blessing of Might or Battle Shout a tie, no Heart of the Lion or Trueshot
// Aura +15%.
func TestRetributionTwistIsPacedNotGatedOnAttackPower(t *testing.T) {
	raw, err := aplFS.ReadFile("apl/paladin-retribution.apl.json")
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
	const unsealedWrath = "1310735"
	twistLines := 0
	for index, line := range apl.PriorityList {
		cast := line.Action.CastSpell
		if cast == nil || (cast.SpellID.SpellID != 20920 && cast.SpellID.SpellID != 20293) {
			continue
		}
		var cond any
		if len(line.Action.Condition) > 0 {
			if err := json.Unmarshal(line.Action.Condition, &cond); err != nil {
				t.Fatalf("line %d: %v", index, err)
			}
		}
		if containsKey(cond, "currentAttackPower") {
			t.Errorf("line %d (%d) is gated on attack power", index, cast.SpellID.SpellID)
		}
		if !containsKey(cond, "auraNumStacks") && !containsKey(cond, "auraRemainingTime") && !(cast.SpellID.SpellID == 20920 && containsKey(cond, "currentSealRemainingTime")) {
			continue
		}
		if !conditionNamesSpell(cond, unsealedWrath) {
			continue
		}
		twistLines++
		var paced, flat bool
		walkManaComparisons(cond, func(cmp map[string]any) {
			lhs, _ := cmp["lhs"].(map[string]any)
			rhs, _ := cmp["rhs"].(map[string]any)
			if _, ok := lhs["currentManaPercent"]; ok {
				_, isConst := rhs["const"]
				flat = flat || isConst
				paced = paced || containsKey(rhs, "remainingTimePercent")
			}
		})
		if flat || !paced {
			t.Errorf("line %d (%d): the twist is not paced on the fight's remaining time (flat %v, paced %v)", index, cast.SpellID.SpellID, flat, paced)
		}
	}
	if twistLines != 3 {
		t.Errorf("found %d Unsealed Wrath twist lines, want 3", twistLines)
	}
}

func conditionNamesSpell(node any, spellID string) bool {
	raw, err := json.Marshal(node)
	return err == nil && strings.Contains(string(raw), `"spellId":`+spellID)
}
