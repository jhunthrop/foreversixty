package discordbot

// sourceCopy is the site BiS page's source-line labels (web/src/lib/bis/copy.ts, the
// questSourceLabel to worldDropSourceLabel block, and items/source-labels.ts), so an item
// reads the same in Discord as on the page. A quest names the level a player must be to take
// it ("requires level 58", spec 2026-10-29 §A.3).
var sourceCopy = struct {
	Quest, QuestLevel, Vendor, Rep, Crafted, CraftedNamed string
	Chance, UnderOne, Trash, Unnamed                      string
	PvP, PvPUntitled, WorldDrop, WorldDropLevels          string
	AllianceQuest, HordeQuest                             string
}{
	Quest:           "Quest: {name}",
	QuestLevel:      "requires level {min_level}",
	Vendor:          "Vendor: {npc}",
	Rep:             "{faction} ({standing})",
	Crafted:         "Crafted: {name}",
	CraftedNamed:    "Crafted: {Profession} · {name}",
	Chance:          "{chance}% from {from}",
	UnderOne:        "<1",
	Trash:           "{instance} trash",
	Unnamed:         "Unnamed source in {zone}",
	PvP:             "PvP rank {rank} · {title} · {Faction}",
	PvPUntitled:     "PvP rank {rank} · {Faction}",
	WorldDrop:       "World drop (BoE)",
	WorldDropLevels: " · levels {min}-{max}",
	AllianceQuest:   "Alliance quest",
	HordeQuest:      "Horde quest",
}

// sourceKindWords are the page's badge words by kind (web/src/lib/sim/loot.ts
// SOURCE_KIND_LABELS), the line a kind loot.json has nothing for falls back to.
var sourceKindWords = map[string]string{
	"raid": "Raids", "dungeon": "Dungeons", "world": "World bosses", "world_drop": "World drops",
	"zone": "Zone drops", "vendor": "Vendors", "crafted": "Crafted", "rep": "Reputation", "pvp": "PvP",
}

// pvpRankOffset turns loot.json's rank (the client's RequiredPVPRank, 5 to 18) into a title
// index; pvpTitles are the ranks 1 to 14 of each side (items/source-labels.ts).
const pvpRankOffset = 5

var pvpTitles = map[string][]string{
	"alliance": {"Private", "Corporal", "Sergeant", "Master Sergeant", "Sergeant Major", "Knight", "Knight-Lieutenant",
		"Knight-Captain", "Knight-Champion", "Lieutenant Commander", "Commander", "Marshal", "Field Marshal", "Grand Marshal"},
	"horde": {"Scout", "Grunt", "Sergeant", "Senior Sergeant", "First Sergeant", "Stone Guard", "Blood Guard", "Legionnaire",
		"Centurion", "Champion", "Lieutenant General", "General", "Warlord", "High Warlord"},
}
