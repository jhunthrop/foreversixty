package guilds

import "testing"

// testOnyxiaEncounter is the encounter the tests' logs and fixtures fight: Onyxia, the one boss
// the 1.60.1.70790 table carries (DungeonEncounter 1084).
const testOnyxiaEncounter = int64(1084)

// testOnyxiaLoot is the table's Onyxia drops, as the loot council reads them.
var testOnyxiaLoot = func() []lootItem {
	b, _ := lootBossByEncounter(testOnyxiaEncounter)
	return b.Items
}()

// TestLootTableKeepsItsExtrasInStep holds lootItemExtras to the generated table: every drop has
// its row and every row is a drop, so a build that adds a boss fails until its drops do.
func TestLootTableKeepsItsExtrasInStep(t *testing.T) {
	drops := map[int]bool{}
	for _, r := range firstTierRaids {
		for _, b := range r.Bosses {
			for _, it := range b.Loot {
				drops[it.ItemID] = true
				if _, ok := lootItemExtras[it.ItemID]; !ok {
					t.Errorf("%s drops %d %s, which lootItemExtras has no row for", b.Name, it.ItemID, it.Name)
				}
			}
		}
	}
	for id := range lootItemExtras {
		if !drops[id] {
			t.Errorf("lootItemExtras has %d, which no table boss drops", id)
		}
	}
}

// TestLootTableIsTheGeneratedTable reads every boss and drop from firstTierRaids, in its order,
// with the boss's raid and the drop's encounter.
func TestLootTableIsTheGeneratedTable(t *testing.T) {
	i := 0
	for _, r := range firstTierRaids {
		for _, b := range r.Bosses {
			got := lootBosses[i]
			if got.RaidKey != r.Key || got.RaidName != r.Name || got.EncounterID != b.EncounterID || got.Name != b.Name {
				t.Errorf("boss %d = %+v, want %s %d of %s", i, got, b.Name, b.EncounterID, r.Key)
			}
			for j, it := range b.Loot {
				item := got.Items[j]
				if item.ItemID != it.ItemID || item.Name != it.Name || item.Quality != it.Quality || item.ItemLevel != it.ItemLevel ||
					item.InventoryType != it.InventoryType || item.EncounterID != b.EncounterID {
					t.Errorf("%s drop %d = %+v, want %+v", b.Name, j, item, it)
				}
			}
			i++
		}
	}
	if i != len(lootBosses) || len(testOnyxiaLoot) == 0 {
		t.Fatalf("lootBosses = %d, table = %d; Onyxia drops %d", len(lootBosses), i, len(testOnyxiaLoot))
	}
}

// TestLootRaidGroupsNameEveryFirstTierRaid lists every first-tier raid for the picker, the ones
// the client does not carry yet with no encounters.
func TestLootRaidGroupsNameEveryFirstTierRaid(t *testing.T) {
	groups := lootRaidGroups()
	if len(groups) != len(raidZones) {
		t.Fatalf("groups = %d, want %d", len(groups), len(raidZones))
	}
	for i, g := range groups {
		if g.Key != raidZones[i].Key || g.Name != raidZones[i].Name {
			t.Errorf("group %d = %s, want %s", i, g.Key, raidZones[i].Key)
		}
		if g.InClient != (len(g.Encounters) > 0) || g.InClient != raidHasBosses(g.Key) {
			t.Errorf("%s: in_client %v with %d encounters", g.Key, g.InClient, len(g.Encounters))
		}
	}
	byKey := map[string]LootRaidGroup{}
	for _, g := range groups {
		byKey[g.Key] = g
	}
	if byKey["raid:onyxias-lair"].InClient != true || byKey["raid:barrow-deeps"].InClient || byKey["raid:hyjal-summit"].InClient {
		t.Errorf("groups = %+v, want Onyxia's Lair alone in the client", groups)
	}
}

// TestZoneEncounterPairsAreParallel keeps bossZones and bossEncounters one pair per boss.
func TestZoneEncounterPairsAreParallel(t *testing.T) {
	zones, encs := bossZones(), bossEncounters()
	if len(zones) != len(encs) || len(zones) != len(lootBosses) {
		t.Fatalf("zones %d, encounters %d, bosses %d", len(zones), len(encs), len(lootBosses))
	}
	for i, b := range lootBosses {
		if zones[i] != b.RaidKey || encs[i] != b.EncounterID {
			t.Errorf("pair %d = %s/%d, want %s/%d", i, zones[i], encs[i], b.RaidKey, b.EncounterID)
		}
	}
	if dropEncounter(testOnyxiaLoot[0].ItemID) != testOnyxiaEncounter || dropEncounter(1) != 0 {
		t.Error("dropEncounter does not read the table")
	}
}
