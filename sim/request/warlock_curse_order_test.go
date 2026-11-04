package request

import (
	"encoding/json"
	"testing"
)

// Sim-accuracy round 15 (warlock lane, guide-accuracy round 1 S1): the
// warlock who holds Curse of the Elements puts it up before Corruption. It
// raises every magic school the raid deals from its first global, which a
// one-player sim cannot see (the personal order measures within ±0.1% at bare
// 60 for all three specs). Its own line still skips when the raid's copy is
// already on the target.
func TestWarlockRotationsCurseOfTheElementsBeforeCorruption(t *testing.T) {
	const curseOfTheElements, corruption = 1311680, 11672
	for _, spec := range []string{"warlock-affliction", "warlock-demonology", "warlock-destruction"} {
		raw, err := aplFS.ReadFile("apl/" + spec + ".apl.json")
		if err != nil {
			t.Fatal(err)
		}
		var rot struct {
			PriorityList []struct {
				Action struct {
					CastSpell *struct {
						SpellID struct {
							SpellID int `json:"spellId"`
						} `json:"spellId"`
					} `json:"castSpell"`
				} `json:"action"`
			} `json:"priorityList"`
		}
		if err := json.Unmarshal(raw, &rot); err != nil {
			t.Fatal(err)
		}
		at := map[int]int{}
		for i, entry := range rot.PriorityList {
			if entry.Action.CastSpell != nil {
				if _, seen := at[entry.Action.CastSpell.SpellID.SpellID]; !seen {
					at[entry.Action.CastSpell.SpellID.SpellID] = i
				}
			}
		}
		coe, okCoE := at[curseOfTheElements]
		corr, okCorr := at[corruption]
		if !okCoE || !okCorr {
			t.Fatalf("%s: Curse of the Elements or Corruption missing from the priority list", spec)
		}
		if coe > corr {
			t.Errorf("%s: Curse of the Elements at priority %d, after Corruption at %d", spec, coe, corr)
		}
	}
}
