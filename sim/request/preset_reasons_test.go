package request

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// raidBuffReasons reads the raid preset's buff and debuff reasons straight
// from the curated file: the reasons are not part of the resolved preset,
// but the BiS page prints them.
func raidBuffReasons(t *testing.T) map[string]string {
	t.Helper()
	raw, err := os.ReadFile(realPresetsPath)
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Presets map[string]presetDef `json:"presets"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	reasons := map[string]string{}
	raid := file.Presets[RaidPreset]
	for _, entry := range append(raid.Buffs, raid.Debuffs...) {
		reasons[entry.ID] = entry.Reason
	}
	return reasons
}

// The raid preset's reasons state what the 1.60.1.70753 client says (round
// 12 auras A1, A5, A7, B2; a later merge had put the older text back):
// Divine Spirit is a trainer spell, Trueshot Aura is ranged attack power
// only, Heart of the Lion is 280 attack power at 60 (40 + 4 a level), and
// Windfury Totem's bonus ends with the totem.
func TestRaidPresetReasonsMatchTheClient(t *testing.T) {
	reasons := raidBuffReasons(t)
	cases := []struct {
		id, want, forbidden string
	}{
		{"divine_spirit", "trainer spell", "Discipline talent"},
		{"trueshot_aura", "Ranged attack power only", ""},
		{"heart_of_the_lion", "280 at 60", "276"},
		{"windfury_totem", "ends with the totem", ""},
	}
	for _, c := range cases {
		reason, ok := reasons[c.id]
		if !ok {
			t.Errorf("the raid preset has no %s entry", c.id)
			continue
		}
		if !strings.Contains(reason, c.want) {
			t.Errorf("%s's reason does not say %q:\n%s", c.id, c.want, reason)
		}
		if c.forbidden != "" && strings.Contains(reason, c.forbidden) {
			t.Errorf("%s's reason still says %q:\n%s", c.id, c.forbidden, reason)
		}
	}
}
