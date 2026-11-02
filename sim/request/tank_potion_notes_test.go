package request

import (
	"encoding/json"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

// Round 11 (regression lane): round 10 moved the bear's raid potion from
// Greater Stoneshield Potion to Mighty Rage Potion (S fd109430b), but the
// rotation's potion line, the preset's notes and the dropped Stoneshield row
// still told players the bear drinks Stoneshield. Greater Stoneshield
// Potion's client RequiredLevel is 46, the same as Mighty Rage Potion's, so
// it is never the bear's potion at any level.
func TestTheBearsPotionCopyNamesTheRaidPotion(t *testing.T) {
	raw, err := os.ReadFile(raidConsumesPath)
	if err != nil {
		t.Fatal(err)
	}
	var file raidConsumesFile
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	rows := file.Specs["druid-feral-bear"]
	if len(rows) == 0 || !slices.Contains(rows[len(rows)-1].Consumes, "mighty_rage_potion") {
		t.Fatalf("the bear's level-60 raid row no longer drinks Mighty Rage Potion: %v", rows)
	}

	curated, err := loadLadderCurated(repoRoot, "druid-feral-bear")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		if item.GetAction().GetAutocastOtherCooldowns() == nil {
			continue
		}
		if notes := item.GetNotes(); strings.Contains(notes, "Stoneshield") || !strings.Contains(notes, "Mighty Rage Potion") {
			t.Errorf("the bear's potion line says %q; the raid preset drinks Mighty Rage Potion", notes)
		}
	}

	presets, err := os.ReadFile(repoRoot + "/data/curated/presets.json")
	if err != nil {
		t.Fatal(err)
	}
	for _, stale := range []string{"the bear's Stoneshield", "the bear's potion below level 46"} {
		if strings.Contains(string(presets), stale) {
			t.Errorf("data/curated/presets.json still says %q", stale)
		}
	}
}
