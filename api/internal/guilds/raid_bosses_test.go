package guilds

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/raidbosses"
)

// TestRaidBossesTableMatchesItsGenerator pins raid_bosses_table.go to what cmd/genraidbosses
// writes from the build's raid-bosses.json.
func TestRaidBossesTableMatchesItsGenerator(t *testing.T) {
	body, err := os.ReadFile(filepath.Join(repoDataDirForTrees(t), raidBossesBuild, "raid-bosses.json"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := raidbosses.Parse(body)
	if err != nil {
		t.Fatal(err)
	}
	want, err := raidbosses.GoSource(doc, raidbosses.SourcePath(raidBossesBuild))
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile("raid_bosses_table.go")
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) {
		t.Fatal("raid_bosses_table.go is stale: run go generate ./internal/guilds/ from api/")
	}
}

// TestRaidBossesAreFirstTierZones keeps every table raid a zone of first-tier.json (raidZones).
func TestRaidBossesAreFirstTierZones(t *testing.T) {
	zones := map[string]bool{}
	for _, z := range raidZones {
		zones[z.Key] = true
	}
	for _, r := range firstTierRaids {
		if !zones[r.Key] {
			t.Errorf("%s is in the boss table but not a first-tier zone", r.Key)
		}
	}
}

// TestRaidBossesCarryOnyxiaAsTheLootCouncilKnowsHer is the swap's safety net: the table's Onyxia
// is progression.go's encounter and carries every drop onyxiaLoot lists, with the same item rows.
func TestRaidBossesCarryOnyxiaAsTheLootCouncilKnowsHer(t *testing.T) {
	var onyxia *tierBoss
	for _, r := range firstTierRaids {
		for i := range r.Bosses {
			if r.Bosses[i].EncounterID == encounterOnyxia {
				if r.Key != onyxiaZoneKey {
					t.Errorf("encounter %d is under %s, want %s", encounterOnyxia, r.Key, onyxiaZoneKey)
				}
				onyxia = &r.Bosses[i]
			}
		}
	}
	if onyxia == nil {
		t.Fatalf("the table has no encounter %d", encounterOnyxia)
	}
	drops := map[int]tierItem{}
	for _, it := range onyxia.Loot {
		drops[it.ItemID] = it
	}
	for _, want := range onyxiaLoot {
		got, ok := drops[want.ItemID]
		if !ok {
			t.Errorf("Onyxia's table lacks %d %s", want.ItemID, want.Name)
			continue
		}
		if got.Name != want.Name || got.Quality != want.Quality || got.ItemLevel != want.ItemLevel ||
			got.RequiredLevel != want.RequiredLevel || got.ItemClass != want.ItemClass ||
			got.Subclass != want.Subclass || got.InventoryType != want.InventoryType {
			t.Errorf("item %d: table %+v, onyxiaLoot %+v", want.ItemID, got, want)
		}
	}
}
