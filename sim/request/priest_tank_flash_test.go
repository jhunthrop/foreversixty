package request

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// The tank's emergency Flash Heal waits on the mana pace on both priest
// lists (sim-accuracy 2026-11-13, discipline.md F1). Flash Heal heals the
// least for its mana of anything the two lists cast; unpaced, it fired
// whenever the tank fell under 35% all fight, after the bar had run dry,
// and spent mana a Greater Heal would have made more of: about a quarter
// of Discipline's spend at raid 60. The party's emergency Flash Heals were
// paced in round 6; the tank's was not.
func TestPriestTankFlashHealWaitsOnTheManaPace(t *testing.T) {
	const flashHeal, tank = 10917, 5
	for _, spec := range []string{"priest-discipline", "priest-holy"} {
		raw, err := os.ReadFile("../../data/curated/apl/" + spec + ".json")
		if err != nil {
			t.Fatal(err)
		}
		var file struct {
			Rotation struct {
				PriorityList []struct {
					Action struct {
						CastSpell *struct {
							SpellID struct {
								SpellID int `json:"spellId"`
							} `json:"spellId"`
							Target *struct {
								Index int `json:"index"`
							} `json:"target"`
						} `json:"castSpell"`
						Condition json.RawMessage `json:"condition"`
					} `json:"action"`
				} `json:"priorityList"`
			} `json:"rotation"`
		}
		if err := json.Unmarshal(raw, &file); err != nil {
			t.Fatal(err)
		}
		lines := 0
		for _, line := range file.Rotation.PriorityList {
			cast := line.Action.CastSpell
			if cast == nil || cast.SpellID.SpellID != flashHeal || cast.Target == nil || cast.Target.Index != tank {
				continue
			}
			lines++
			condition := string(line.Action.Condition)
			if !strings.Contains(condition, "currentManaPercent") || !strings.Contains(condition, "remainingTimePercent") {
				t.Errorf("%s: the tank's Flash Heal does not wait on the mana pace: %s", spec, condition)
			}
		}
		if lines != 1 {
			t.Errorf("%s: %d tank Flash Heal lines, want 1", spec, lines)
		}
	}
}
