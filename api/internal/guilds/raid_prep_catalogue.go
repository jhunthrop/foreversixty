// api/internal/guilds/raid_prep_catalogue.go
//
// What the raid-prep checks (owner decision D14) count: a copy of data/curated/raid-supplies.json's
// supplies, which the addon's Data.lua also carries (a test pins the two). Every id and name below is
// in data/builds/1.60.1.70790/simconsumes.json (checked 2026-11-06): the addon exports every
// Consumable-class stack in the bags (supplies=) and this table decides which of them a raid
// leader counts, so a new Forever consumable is a line here and in the curated file.
//
// Flasks: the four Classic raid flasks and Forever's four Natural flasks (Flask of
// Petrification is a utility flask and is left out). Elixirs: the raid-tier battle and
// guardian elixirs, the Blasted Lands and Zul'Gurub buffs and Forever's Fortitude and
// Mageblood elixirs. Food: every item whose spell is Forever's "Nutritious Food" (the Well
// Fed rework) at required level 30 or more, the level-1 copies of Tender Wolf Steak and
// Grilled Squid that share those spells, Dirge's Kickin' Chimaerok Chops and Blessed Sunfruit
// and its juice.
package guilds

// The kinds of raid supply a check counts.
const (
	supplyFlask  = "flask"
	supplyElixir = "elixir"
	supplyFood   = "food"
)

// raidSupply is one counted consumable.
type raidSupply struct {
	Name string
	Kind string
}

// raidSupplies is item id -> what it counts as.
var raidSupplies = map[int]raidSupply{
	13510:  {"Flask of the Titans", supplyFlask},
	13511:  {"Flask of Distilled Wisdom", supplyFlask},
	13512:  {"Flask of Supreme Power", supplyFlask},
	13513:  {"Flask of Chromatic Resistance", supplyFlask},
	274273: {"Flask of Natural Accuracy", supplyFlask},
	274274: {"Flask of Natural Aggression", supplyFlask},
	274275: {"Flask of Natural Precision", supplyFlask},
	274276: {"Flask of Natural Swiftness", supplyFlask},
	8410:   {"R.O.I.D.S.", supplyElixir},
	8411:   {"Lung Juice Cocktail", supplyElixir},
	8412:   {"Ground Scorpok Assay", supplyElixir},
	8423:   {"Cerebral Cortex Compound", supplyElixir},
	9088:   {"Gift of Arthas", supplyElixir},
	9155:   {"Arcane Elixir", supplyElixir},
	9179:   {"Elixir of Greater Intellect", supplyElixir},
	9187:   {"Elixir of Greater Agility", supplyElixir},
	9264:   {"Elixir of Shadow Power", supplyElixir},
	12820:  {"Winterfall Firewater", supplyElixir},
	13445:  {"Elixir of Greater Defense", supplyElixir},
	13447:  {"Elixir of the Sages", supplyElixir},
	13452:  {"Elixir of the Mongoose", supplyElixir},
	13453:  {"Elixir of Brute Force", supplyElixir},
	13454:  {"Greater Arcane Elixir", supplyElixir},
	17708:  {"Elixir of Frost Power", supplyElixir},
	20004:  {"Major Troll's Blood Elixir", supplyElixir},
	20007:  {"Mageblood Elixir", supplyElixir},
	20079:  {"Spirit of Zanza", supplyElixir},
	20080:  {"Sheen of Zanza", supplyElixir},
	20081:  {"Swiftness of Zanza", supplyElixir},
	23715:  {"Permanent Lung Juice Cocktail", supplyElixir},
	23718:  {"Permanent Ground Scorpok Assay", supplyElixir},
	23719:  {"Permanent Cerebral Cortex Compound", supplyElixir},
	23722:  {"Permanent R.O.I.D.S.", supplyElixir},
	23794:  {"Permanent Sheen of Zanza", supplyElixir},
	23795:  {"Permanent Spirit of Zanza", supplyElixir},
	23796:  {"Permanent Swiftness of Zanza", supplyElixir},
	250334: {"Elixir of Fortitude", supplyElixir},
	250335: {"Elixir of Greater Fortitude", supplyElixir},
	250341: {"Greater Mageblood Elixir", supplyElixir},
	6887:   {"Spotted Yellowtail", supplyFood},
	8957:   {"Spinefin Halibut", supplyFood},
	12212:  {"Jungle Stew", supplyFood},
	13810:  {"Blessed Sunfruit", supplyFood},
	13813:  {"Blessed Sunfruit Juice", supplyFood},
	13927:  {"Cooked Glossy Mightfish", supplyFood},
	13928:  {"Grilled Squid", supplyFood},
	13929:  {"Hot Smoked Bass", supplyFood},
	13930:  {"Filet of Redgill", supplyFood},
	13931:  {"Nightfin Soup", supplyFood},
	13932:  {"Poached Sunscale Salmon", supplyFood},
	13933:  {"Lobster Stew", supplyFood},
	13934:  {"Mightfish Steak", supplyFood},
	13935:  {"Baked Salmon", supplyFood},
	17222:  {"Spider Sausage", supplyFood},
	18045:  {"Tender Wolf Steak", supplyFood},
	18254:  {"Runn Tum Tuber Surprise", supplyFood},
	20452:  {"Smoked Desert Dumplings", supplyFood},
	21023:  {"Dirge's Kickin' Chimaerok Chops", supplyFood},
	21217:  {"Sagefish Delight", supplyFood},
	250065: {"Savory Stag Sliders", supplyFood},
	250066: {"Soaring Pamplona", supplyFood},
	250067: {"Bat Hachee", supplyFood},
	250068: {"Savory Turtle Stew", supplyFood},
	250069: {"Flank au Poivre", supplyFood},
	250070: {"Bear Bruscitti", supplyFood},
	250071: {"Steaming Stag Steak", supplyFood},
	250072: {"Swiftstrike Steak", supplyFood},
	250075: {"Prehistoric Pulled Raptor", supplyFood},
	250076: {"Raptor Rouladen", supplyFood},
	267341: {"Sweetpaw Jam", supplyFood},
	286152: {"Plated Armorfish", supplyFood},
	287973: {"Tender Wolf Steak", supplyFood},
	287974: {"Grilled Squid", supplyFood},
}

// fireResistanceZones are the raids whose checklist shows the fire resistance total: only a raid
// that asks for it. Of Forever's first tier (Barrow Deeps, Hyjal Summit, Onyxia's Lair) that is
// Onyxia's breath and Flame Breath; a later raid joins by its key.
var fireResistanceZones = map[string]bool{
	"raid:onyxias-lair": true,
}

// Attunement is a zone's attunement by quest: done when the client flags any of Quests
// complete (one per faction), or when Item, what the chain hands out (nil for a chain that
// hands out none), is held.
type Attunement struct {
	Quests []int
	Item   *KeyItem
}

// The attunement chains' last quests. Each id is a row of the 1.60.1.70790 client's QuestV2
// table (wago.tools db2/QuestV2?build=1.60.1.70790, fetched 2026-11-06; the table carries no
// titles), and its title and chain order are cmangos classic-db quest_template's
// (PrevQuestId), the source data/builds/1.60.1.70790/raw/quests/quest-levels.json names.
// RaidPrep.ATTUNEMENT_QUESTS asks the client about the same two.
const (
	// questDrakefireAmulet ends the Alliance chain: Stormwind Rendezvous 6402, The Great
	// Masquerade 6403, The Dragon's Eye 6501, Drakefire Amulet 6502.
	questDrakefireAmulet = 6502
	// questBlackDragonChampion ends the Horde chain: Ascension... 6601, Blood of the Black
	// Dragon Champion 6602.
	questBlackDragonChampion = 6602
)

// drakefireAmulet is the item both Onyxia chains reward (items.json: 16309, a neck item of
// quality 3); the export says whether it is held (owned=).
var drakefireAmulet = KeyItem{ID: 16309, Name: "Drakefire Amulet", Quality: 3}

// zoneAttunements is zone key -> its attunement by quest. Kept beside the zone catalogue
// rather than on Zone: first-tier.json, which the raid zones mirror, carries no quests. Of the
// first tier only Onyxia's Lair has an attunement.
var zoneAttunements = map[string]Attunement{
	"raid:onyxias-lair": {Quests: []int{questDrakefireAmulet, questBlackDragonChampion}, Item: &drakefireAmulet},
}

// attunementFor is the zone's attunement by quest, ok false when it has none.
func attunementFor(zone Zone) (Attunement, bool) {
	a, ok := zoneAttunements[zone.Key]
	return a, ok
}
