package request

import (
	"encoding/json"
	"testing"
)

// protectionConsecrationID is the rank Protection's rotation authors
// Consecration at (rank 5); rank-choices.json moves it below 60.
const protectionConsecrationID = 20924

// Sim-accuracy round 24: Protection's Consecration was held behind "mana above
// 50%", which spent the bar at the band floors where Holy Strike needed it and
// sat on it where nothing else could use it. The line now casts while the pool
// plus the net rate (what comes in less what the rotation spends, over the last
// minute) times the fight's remaining time still covers Consecration, and also
// when you drop below 35% health, so mana is not held through a fight you may
// not finish; Holy Shield's reserve stays.
func TestProtectionConsecrationSpendsWhatTheNetRateCovers(t *testing.T) {
	raw, err := aplFS.ReadFile("apl/paladin-protection.apl.json")
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
		if cast == nil || cast.SpellID.SpellID != protectionConsecrationID {
			continue
		}
		found = true
		var cond any
		if err := json.Unmarshal(e.Action.Condition, &cond); err != nil {
			t.Fatalf("Consecration's condition: %v", err)
		}
		var flat, netRate, reserve bool
		walkManaComparisons(cond, func(cmp map[string]any) {
			lhs, _ := cmp["lhs"].(map[string]any)
			rhs, _ := cmp["rhs"].(map[string]any)
			if _, ok := lhs["currentManaPercent"]; ok {
				_, isConst := rhs["const"]
				flat = flat || isConst
			}
			netRate = netRate || (containsKey(lhs, "manaIncomePerSecond") && containsKey(lhs, "manaSpentPerSecond") && containsKey(lhs, "remainingTime"))
			reserve = reserve || (containsKey(lhs, "currentMana") && containsKey(rhs, "spellCurrentCost") && !containsKey(lhs, "remainingTime"))
		})
		if flat {
			t.Error("Consecration is rationed by a flat mana percentage")
		}
		if !netRate {
			t.Error("Consecration does not weigh the pool against income less spend over the fight's remaining time")
		}
		if !reserve {
			t.Error("Consecration does not keep Holy Shield's cost in reserve")
		}
	}
	if !found {
		t.Fatalf("the rotation casts no Consecration %d", protectionConsecrationID)
	}
}

// Re-judged on the net-rate line (sim-accuracy round 24) on the old 180 s
// fight: rank 1 at 30 to 33, 40 to 48 and 50 to 57. Re-judged on the 48 s
// dungeon pull (sim-accuracy 2026-11-12, paladins): rank 1 ties rank 2 at 30
// to 33 and beats it by 0.2 to 0.8% at 34 to 38, loses 0.7 to 1.4% to rank 3
// at 40 to 49 and 1.2 to 2.1% to rank 4 at 50 to 59, so the row is 30 to 38
// and the trained top rank follows from 39; rank 1 keeps bare 60 (+0.6%).
func TestProtectionCastsTheReMeasuredConsecrationRanks(t *testing.T) {
	for _, c := range []struct {
		preset string
		level  int
		want   int32
	}{
		{BarePreset, 30, 26573},
		{BarePreset, 33, 26573},
		{BarePreset, 34, 26573},
		{BarePreset, 38, 26573},
		{BarePreset, 39, 20116},
		{BarePreset, 40, 20922},
		{BarePreset, 48, 20922},
		{BarePreset, 49, 20922},
		{BarePreset, 50, 20923},
		{BarePreset, 57, 20923},
		{BarePreset, 59, 20923},
		{BarePreset, 60, 26573},
		{RaidPreset, 60, protectionConsecrationID},
	} {
		for _, race := range []string{"dwarf", "undead"} {
			ids := rotationIDs(t, "paladin-protection", "paladin", c.level, specRankChoices("paladin-protection", c.preset, race))
			if !ids[c.want] {
				t.Errorf("%s %s %d: Consecration does not cast %d", race, c.preset, c.level, c.want)
			}
		}
	}
}
