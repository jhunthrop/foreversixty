// api/internal/guilds/loot_item_extras.go
//
// The four facts of a drop that the generated boss table (raid_bosses_table.go, from the build's
// raid-bosses.json) does not carry, keyed by item id. Everything else about a drop (name, quality,
// item level, class, subclass, inventory type, its boss and raid) is the generated table's, so a
// drop has one source per field. These four stay here because the build's tracked files cannot
// give them for every drop: six of Onyxia's drops have no inventory slot, so neither the per-class
// item pools (data/builds/<build>/items/<class>.json, icon, slot and unique) nor classic-db's
// narrowed item_template (allowable_class, equippable items only) name them, and their icons are
// the client's Item.IconFileDataID through ManifestInterfaceData (wago.tools, build 1.60.1.70532).
// loot_table_test.go keeps this map and the table in step: every table drop has a row here and
// every row here is a table drop, so a client build that adds a boss fails until its drops do.
package guilds

// lootItemExtra is a drop's Icon (the client's file name), Slot (the BiS slot of the per-class
// item pools, "" for a drop with no slot), AllowableClass (classic-db item_template's
// allowable_class, -1 for no restriction) and Unique (item_template's unique flag: a character
// holds one).
type lootItemExtra struct {
	Icon           string
	Slot           string
	AllowableClass int
	Unique         bool
}

// anyClass is AllowableClass for a drop every class may use.
const anyClass = -1

// lootItemExtras is every table drop's extra facts.
var lootItemExtras = map[int]lootItemExtra{
	15410: {Icon: "inv_misc_monsterscales_07", AllowableClass: anyClass},                                // Scale of Onyxia
	16900: {Icon: "inv_helmet_09", Slot: "head", AllowableClass: 1024},                                  // Stormrage Cover
	16908: {Icon: "inv_helmet_41", Slot: "head", AllowableClass: 8},                                     // Bloodfang Hood
	16914: {Icon: "inv_helmet_70", Slot: "head", AllowableClass: 128},                                   // Netherwind Crown
	16921: {Icon: "inv_helmet_24", Slot: "head", AllowableClass: 16},                                    // Halo of Transcendence
	16929: {Icon: "inv_helmet_08", Slot: "head", AllowableClass: 256},                                   // Nemesis Skullcap
	16939: {Icon: "inv_helmet_05", Slot: "head", AllowableClass: 4},                                     // Dragonstalker's Helm
	16947: {Icon: "inv_helmet_69", Slot: "head", AllowableClass: 64},                                    // Helmet of Ten Storms
	16955: {Icon: "inv_helmet_74", Slot: "head", AllowableClass: 2},                                     // Judgement Crown
	16963: {Icon: "inv_helmet_71", Slot: "head", AllowableClass: 1},                                     // Helm of Wrath
	17064: {Icon: "inv_misc_monsterscales_15", Slot: "trinket", AllowableClass: anyClass, Unique: true}, // Shard of the Scale
	17067: {Icon: "inv_misc_book_07", Slot: "off_hand", AllowableClass: anyClass},                       // Ancient Cornerstone Grimoire
	17068: {Icon: "inv_axe_09", Slot: "main_hand", AllowableClass: anyClass},                            // Deathbringer
	17075: {Icon: "inv_sword_18", Slot: "main_hand", AllowableClass: anyClass, Unique: true},            // Vis'kag the Bloodletter
	17078: {Icon: "inv_misc_cape_16", Slot: "back", AllowableClass: anyClass},                           // Sapphiron Drape
	17966: {Icon: "inv_misc_bag_22", AllowableClass: anyClass},                                          // Onyxia Hide Backpack
	18205: {Icon: "inv_belt_12", Slot: "neck", AllowableClass: anyClass, Unique: true},                  // Eskhandar's Collar
	18422: {Icon: "inv_misc_head_dragon_01", AllowableClass: anyClass},                                  // Head of Onyxia
	18423: {Icon: "inv_misc_head_dragon_01", AllowableClass: anyClass},                                  // Head of Onyxia
	// classic-db's full item_template gives the Sinew AllowableClass 4 (hunter); the build's
	// narrowed raw/classicdb/item_template.json keeps equippable items only.
	18705: {Icon: "spell_holy_dispelmagic", AllowableClass: 4},                                   // Mature Black Dragon Sinew
	18813: {Icon: "inv_jewelry_ring_13", Slot: "finger", AllowableClass: anyClass, Unique: true}, // Ring of Binding
	21108: {Icon: "inv_misc_stonetablet_03", AllowableClass: anyClass},                           // Draconic for Dummies
}
