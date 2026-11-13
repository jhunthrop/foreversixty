// api/internal/guilds/raid_bosses.go
//
// The first raid tier's bosses and their loot, from the client's own rows
// (data/builds/<build>/raid-bosses.json, data/pipeline/raid_bosses.py; Go twin
// raid_bosses_table.go). The loot picker and the raid assignment templates read it. A raid the
// client does not name yet (Barrow Deeps and Hyjal Summit on 1.60.1.70790) has no entry.
package guilds

//go:generate go run ../../cmd/genraidbosses -build 1.60.1.70938

// tierItem is one boss drop: id, name and quality with items.json's item_level, required_level,
// class_id, subclass_id and inventory_type.
type tierItem struct {
	ItemID        int
	Name          string
	Quality       int
	ItemLevel     int
	RequiredLevel int
	ItemClass     int
	Subclass      int
	InventoryType int
}

// tierBoss is one DungeonEncounter row (EncounterID is the combat log's encounter id).
// JournalEncounterID is 0 when the client ships no Dungeon Journal row for it; LootSource names
// the rows the loot came from.
type tierBoss struct {
	EncounterID        int64
	JournalEncounterID int64
	Name               string
	Order              int
	LootSource         string
	Loot               []tierItem
}

// tierRaid is one first-tier raid the client names. Key is data/curated/raids/first-tier.json's
// zone key; MaxPlayers is the client's MapDifficulty.MaxPlayers.
type tierRaid struct {
	Key        string
	Name       string
	MaxPlayers int
	Bosses     []tierBoss
}
