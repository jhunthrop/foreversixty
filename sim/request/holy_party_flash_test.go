package request

import (
	"encoding/json"
	"testing"
)

// holyPartyFlashOfLightLines are the Holy rotation's party Flash of Light
// lines (everyone else below 80%): priorityList 12 to 16, authored at
// rank 4 (19941).
var holyPartyFlashOfLightLines = []int{12, 13, 14, 15, 16}

const (
	flashOfLightRank4 = 19941
	flashOfLightRank5 = 19942
)

// holyLineCasts is the spell each rewritten priorityList line casts.
func holyLineCasts(t *testing.T, preset string, level int) []int32 {
	t.Helper()
	raw, err := aplFS.ReadFile("apl/paladin-holy.apl.json")
	if err != nil {
		t.Fatal(err)
	}
	got, err := rewriteRotationRanks(raw, "paladin", level, specRankChoices("paladin-holy", preset, "human"))
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
			} `json:"action"`
		} `json:"priorityList"`
	}
	if err := json.Unmarshal(got, &apl); err != nil {
		t.Fatal(err)
	}
	casts := make([]int32, len(apl.PriorityList))
	for i, e := range apl.PriorityList {
		if e.Action.CastSpell != nil {
			casts[i] = e.Action.CastSpell.SpellID.SpellID
		}
	}
	return casts
}

// Sim-accuracy round 22: at raid 60 the party Flash of Light heals more at
// rank 5 than at the authored rank 4, beyond error at seeds 1 and 50001 on
// both factions; without raid buffs the party lines never fire, so bare
// keeps the authored rank.
func TestHolyPartyFlashOfLightIsRankFiveInARaid(t *testing.T) {
	for _, c := range []struct {
		preset string
		want   int32
	}{
		{RaidPreset, flashOfLightRank5},
		{BarePreset, flashOfLightRank4},
	} {
		casts := holyLineCasts(t, c.preset, 60)
		for _, line := range holyPartyFlashOfLightLines {
			if casts[line] != c.want {
				t.Errorf("%s 60: party line %d casts %d, want %d", c.preset, line, casts[line], c.want)
			}
		}
	}
}
