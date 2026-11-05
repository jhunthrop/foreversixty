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

// Sim-accuracy round 22 (follow-up): Forever has no sub-20 cut on heals, so
// Holy Light's ranks learned below 20 were re-judged level by level. Rank 1
// (635) beats the top rank learned by 11 to 16% at 18 and 19, both seeds and
// factions; at 17 rank 3 (647) still wins, so 17 keeps the default.
func TestHolyLightIsRankOneAtEighteenAndNineteen(t *testing.T) {
	const holyLightLine = 3
	for _, c := range []struct {
		level int
		want  int32
	}{
		{17, 647},
		{18, 635},
		{19, 635},
		{20, 635},
	} {
		casts := holyLineCasts(t, BarePreset, c.level)
		if casts[holyLightLine] != c.want {
			t.Errorf("level %d: Holy Light line casts %d, want %d", c.level, casts[holyLightLine], c.want)
		}
	}
}
