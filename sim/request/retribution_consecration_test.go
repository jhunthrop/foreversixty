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
// on the fight's pace: never a flat mana percentage. Sim-accuracy round 23:
// the pace became a mana budget (the pool plus what the paladin regains over
// the fight's remaining time against the Judgements still to come), which
// reads raid mana and bare mana alike: raid 60 ties, bare 60 and every level
// from 20 gain.
func TestRetributionConsecrationSpendsThePoolOnTheBudget(t *testing.T) {
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
		var flat, budget, reserve bool
		walkManaComparisons(cond, func(cmp map[string]any) {
			lhs, _ := cmp["lhs"].(map[string]any)
			rhs, _ := cmp["rhs"].(map[string]any)
			if _, ok := lhs["currentManaPercent"]; ok {
				_, isConst := rhs["const"]
				flat = flat || isConst
			}
			budget = budget || (containsKey(lhs, "manaIncomePerSecond") && containsKey(lhs, "remainingTime") && containsKey(rhs, "remainingTime"))
			reserve = reserve || (containsKey(lhs, "currentMana") && containsKey(rhs, "spellCurrentCost"))
		})
		if flat {
			t.Error("Consecration is rationed by a flat mana percentage")
		}
		if !budget {
			t.Error("Consecration does not weigh the pool and its regen against the fight's remaining time")
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

// Re-measured on the talents the leveling order gives at each level
// (sim-accuracy round 24, second pass, and again on 2026-11-12 with
// Judgement a ranged attack; 10,000 iterations, seeds 1 and 50001, Human,
// Dwarf and Undead). Rank 1 wins at 30 to 52 for every race. Holy Conduit's
// second point at 53 (Consecration 40% cheaper) makes rank 3 the best from
// 53 to 57 by 0.2 to 11% for every race; at 58 and 59 rank 4 beats it by 1.0
// to 1.8% for a Human and a Dwarf while an Undead keeps rank 3. At 60 the
// bare fight casts rank 4, the raid fight keeps rank 5.
func TestRetributionCastsTheReMeasuredConsecrationRanks(t *testing.T) {
	for _, c := range []struct {
		preset string
		level  int
		want   map[string]int32
	}{
		{BarePreset, 30, map[string]int32{"human": 26573, "dwarf": 26573, "undead": 26573}},
		{BarePreset, 43, map[string]int32{"human": 26573, "dwarf": 26573, "undead": 26573}},
		{BarePreset, 47, map[string]int32{"human": 26573, "dwarf": 26573, "undead": 26573}},
		{BarePreset, 49, map[string]int32{"human": 26573, "dwarf": 26573, "undead": 26573}},
		{BarePreset, 51, map[string]int32{"human": 26573, "dwarf": 26573, "undead": 26573}},
		{BarePreset, 52, map[string]int32{"human": 26573, "dwarf": 26573, "undead": 26573}},
		{BarePreset, 53, map[string]int32{"human": 20922, "dwarf": 20922, "undead": 20922}},
		{BarePreset, 56, map[string]int32{"human": 20922, "dwarf": 20922, "undead": 20922}},
		{BarePreset, 57, map[string]int32{"human": 20922, "dwarf": 20922, "undead": 20922}},
		{BarePreset, 58, map[string]int32{"human": 20923, "dwarf": 20923, "undead": 20922}},
		{BarePreset, 59, map[string]int32{"human": 20923, "dwarf": 20923, "undead": 20922}},
		{BarePreset, 60, map[string]int32{"human": 20923, "dwarf": 20923, "undead": 20923}},
		{RaidPreset, 60, map[string]int32{"human": retributionConsecrationID, "dwarf": retributionConsecrationID, "undead": retributionConsecrationID}},
	} {
		for race, want := range c.want {
			ids := rotationIDs(t, "paladin-retribution", "paladin", c.level, specRankChoices("paladin-retribution", c.preset, race))
			if !ids[want] {
				t.Errorf("%s %s %d: Consecration does not cast %d", race, c.preset, c.level, want)
			}
		}
	}
}
