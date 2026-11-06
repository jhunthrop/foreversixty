// api/internal/guilds/loot_table.go
//
// The loot council's table: every boss of the generated first-tier table (firstTierRaids,
// raid_bosses_table.go) with its drops, each joined to the facts the client table does not carry
// (lootItemExtras). Every reader of a drop, a boss or a loot zone goes through here, keyed by the
// raid's zone key and the boss's EncounterID; nothing names an encounter or a zone by hand.
package guilds

import "strconv"

// lootBoss is one table boss as the council reads it: its raid's zone key and name, the combat
// log's encounter id, and its drops in the table's order.
type lootBoss struct {
	RaidKey     string
	RaidName    string
	EncounterID int64
	Name        string
	Items       []lootItem
}

// lootBosses is every table boss, raid by raid in the table's order.
var lootBosses = buildLootBosses(firstTierRaids)

// lootTable is every drop of every table boss, in the table's order.
var lootTable = flattenLoot(lootBosses)

func buildLootBosses(raids []tierRaid) []lootBoss {
	var out []lootBoss
	for _, r := range raids {
		for _, b := range r.Bosses {
			items := make([]lootItem, 0, len(b.Loot))
			for _, t := range b.Loot {
				items = append(items, toLootItem(t, b.EncounterID))
			}
			out = append(out, lootBoss{RaidKey: r.Key, RaidName: r.Name, EncounterID: b.EncounterID, Name: b.Name, Items: items})
		}
	}
	return out
}

func flattenLoot(bosses []lootBoss) []lootItem {
	var out []lootItem
	for _, b := range bosses {
		out = append(out, b.Items...)
	}
	return out
}

// toLootItem is a table drop with its extra facts; a drop with no extras row (which
// loot_table_test.go refuses) has no icon or slot and no class restriction.
func toLootItem(t tierItem, encounterID int64) lootItem {
	extra, ok := lootItemExtras[t.ItemID]
	if !ok {
		extra = lootItemExtra{AllowableClass: anyClass}
	}
	return lootItem{
		ItemID: t.ItemID, Name: t.Name, Icon: extra.Icon, Quality: t.Quality, Slot: extra.Slot,
		ItemClass: t.ItemClass, Subclass: t.Subclass, InventoryType: t.InventoryType,
		RequiredLevel: t.RequiredLevel, ItemLevel: t.ItemLevel, AllowableClass: extra.AllowableClass,
		Unique: extra.Unique, EncounterID: encounterID,
	}
}

// lootItemByID is the table's row for itemID.
func lootItemByID(itemID int) (lootItem, bool) {
	for _, it := range lootTable {
		if it.ItemID == itemID {
			return it, true
		}
	}
	return lootItem{}, false
}

// validLootItem reports whether itemID is a drop of the table.
func validLootItem(itemID int) bool {
	_, ok := lootItemByID(itemID)
	return ok
}

// lootBossByEncounter is the table boss of encounterID.
func lootBossByEncounter(encounterID int64) (lootBoss, bool) {
	for _, b := range lootBosses {
		if b.EncounterID == encounterID {
			return b, true
		}
	}
	return lootBoss{}, false
}

// dropEncounter is the encounter itemID drops from; 0 for an item the table does not hold.
func dropEncounter(itemID int) int64 {
	it, _ := lootItemByID(itemID)
	return it.EncounterID
}

// bossName is the table's name for encounterID, "Encounter <id>" for one it does not hold.
func bossName(encounterID int64) string {
	if b, ok := lootBossByEncounter(encounterID); ok {
		return b.Name
	}
	return "Encounter " + strconv.FormatInt(encounterID, 10)
}

// lootScope is the zones and encounters a loot read or night lookup covers.
type lootScope struct {
	zones      []string
	encounters []int64
}

// allLoot is every table raid and boss.
var allLoot = scopeOf(lootBosses...)

// scopeOf is the scope of bosses: their raids' zones and their encounters.
func scopeOf(bosses ...lootBoss) lootScope {
	s := lootScope{zones: []string{}, encounters: []int64{}}
	seen := map[string]bool{}
	for _, b := range bosses {
		if !seen[b.RaidKey] {
			seen[b.RaidKey] = true
			s.zones = append(s.zones, b.RaidKey)
		}
		s.encounters = append(s.encounters, b.EncounterID)
	}
	return s
}

// bossZones is each table boss's zone key, parallel to bossEncounters: the pairs
// zoneEncounterSQL unnests, so a raid in one zone never takes another zone's loot.
func bossZones() []string {
	out := make([]string, 0, len(lootBosses))
	for _, b := range lootBosses {
		out = append(out, b.RaidKey)
	}
	return out
}

// bossEncounters is each table boss's encounter id, parallel to bossZones.
func bossEncounters() []int64 { return allLoot.encounters }

// zoneEncounterSQL is the filter "the row's encounter is a table boss of the zone", over the
// parallel arrays bound at zonesArg and encountersArg (bossZones, bossEncounters).
func zoneEncounterSQL(zoneColumn, encounterColumn, zonesArg, encountersArg string) string {
	return "exists (select 1 from unnest(" + zonesArg + "::text[], " + encountersArg + "::bigint[]) tb(zone, encounter_id)" +
		" where tb.zone = " + zoneColumn + " and tb.encounter_id = " + encounterColumn + ")"
}

// lootRaidName is the table's name for a zone key, else the curated zone's, else "".
func lootRaidName(zoneKey string) string {
	for _, b := range lootBosses {
		if b.RaidKey == zoneKey {
			return b.RaidName
		}
	}
	if z, ok := ZoneByKey(zoneKey); ok {
		return z.Name
	}
	return ""
}

// bossRaidName is the raid name of encounterID's table boss, "" for one the table does not hold.
func bossRaidName(encounterID int64) string {
	b, _ := lootBossByEncounter(encounterID)
	return b.RaidName
}
