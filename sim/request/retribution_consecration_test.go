package request

import (
	"encoding/json"
	"testing"
)

// retributionConsecrationID is the rank Retribution's rotation authors
// Consecration at (rank 5); rank-choices.json moves it below 60.
const retributionConsecrationID = 20924

// Owner ruling 2, round 10: Retribution's Consecration was rationed behind
// a flat "mana above 65%" line. The line left a fifth to a half of the pool
// unspent at the end of every fight, and it is what made rank 1 look like a
// rank choice. With no guard at all, though, Consecration drains the pool
// and the seal lapses (bare 20 to 50 spend 75 to 115 s of the fight out of
// mana). So the line keeps the core buttons affordable and spends the pool
// on the fight's pace: never a flat mana percentage.
func TestRetributionConsecrationSpendsThePoolOnThePace(t *testing.T) {
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
	found := false
	for _, e := range apl.PriorityList {
		cast := e.Action.CastSpell
		if cast == nil || cast.SpellID.SpellID != retributionConsecrationID {
			continue
		}
		found = true
		var cond any
		if err := json.Unmarshal(e.Action.Condition, &cond); err != nil {
			t.Fatalf("Consecration's condition: %v", err)
		}
		var flat, pace, reserve bool
		walkManaComparisons(cond, func(cmp map[string]any) {
			lhs, _ := cmp["lhs"].(map[string]any)
			rhs, _ := cmp["rhs"].(map[string]any)
			if _, ok := lhs["currentManaPercent"]; ok {
				_, isConst := rhs["const"]
				flat = flat || isConst
				pace = pace || containsKey(rhs, "remainingTimePercent")
			}
			reserve = reserve || (containsKey(lhs, "currentMana") && containsKey(rhs, "spellCurrentCost"))
		})
		if flat {
			t.Error("Consecration is rationed by a flat mana percentage")
		}
		if !pace {
			t.Error("Consecration does not compare the pool with the fight's remaining time")
		}
		if !reserve {
			t.Error("Consecration does not keep the seal, Judgement and Holy Strike affordable")
		}
	}
	if !found {
		t.Fatalf("the rotation casts no Consecration %d", retributionConsecrationID)
	}
}

func walkManaComparisons(node any, visit func(map[string]any)) {
	switch v := node.(type) {
	case map[string]any:
		if cmp, ok := v["cmp"].(map[string]any); ok {
			visit(cmp)
		}
		for _, child := range v {
			walkManaComparisons(child, visit)
		}
	case []any:
		for _, child := range v {
			walkManaComparisons(child, visit)
		}
	}
}

func containsKey(node any, key string) bool {
	switch v := node.(type) {
	case map[string]any:
		if _, ok := v[key]; ok {
			return true
		}
		for _, child := range v {
			if containsKey(child, key) {
				return true
			}
		}
	case []any:
		for _, child := range v {
			if containsKey(child, key) {
				return true
			}
		}
	}
	return false
}

// Re-measured on the paced line (round 10): rank 1 still wins from 30 to 59,
// with the gate taken away as well, where Retribution then runs dry; at 60
// the bare fight runs dry and casts rank 4, the raid fight keeps rank 5.
func TestRetributionCastsTheReMeasuredConsecrationRanks(t *testing.T) {
	for _, c := range []struct {
		preset string
		level  int
		want   int32
	}{
		{BarePreset, 30, 26573},
		{BarePreset, 45, 26573},
		{BarePreset, 59, 26573},
		{BarePreset, 60, 20923},
		{RaidPreset, 60, retributionConsecrationID},
	} {
		ids := rotationIDs(t, "paladin-retribution", "paladin", c.level, specRankChoices("paladin-retribution", c.preset, "human"))
		if !ids[c.want] {
			t.Errorf("%s %d: Consecration does not cast %d", c.preset, c.level, c.want)
		}
	}
}
