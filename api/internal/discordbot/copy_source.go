package discordbot

import "github.com/jhunthrop/foreversixty/api/internal/copygen"

// sourceCopy is the site BiS page's source-line labels (web/src/lib/bis/copy.ts, the
// questSourceLabel to worldDropSourceLabel block, and items/source-labels.ts), so an item
// reads the same in Discord as on the page. A quest names the level a player must be to take
// it ("requires level 58", spec 2026-10-29 §A.3).
var sourceCopy = struct {
	Quest, QuestLevel, Vendor, Rep, Crafted, CraftedNamed   string
	Chance, UnderOne, Trash, Unnamed                        string
	PvP, PvPUntitled, PvPReward, WorldDrop, WorldDropLevels string
	AllianceQuest, HordeQuest                               string
}{
	Quest:           copygen.BisSourceQuest,
	QuestLevel:      copygen.DiscordSourceQuestLevel,
	Vendor:          copygen.DiscordSourceVendor,
	Rep:             copygen.TooltipRepSource,
	Crafted:         copygen.BisSourceCrafted,
	CraftedNamed:    copygen.DiscordSourceCraftedNamed,
	Chance:          copygen.DiscordSourceChance,
	UnderOne:        copygen.DiscordSourceUnderOne,
	Trash:           copygen.DiscordSourceTrash,
	Unnamed:         copygen.DiscordSourceUnnamed,
	PvP:             copygen.DiscordSourcePvP,
	PvPUntitled:     copygen.DiscordSourcePvPUntitled,
	PvPReward:       copygen.DiscordSourcePvPReward,
	WorldDrop:       copygen.DiscordSourceWorldDrop,
	WorldDropLevels: copygen.DiscordSourceWorldDropLevels,
	AllianceQuest:   copygen.DiscordSourceAllianceQuest,
	HordeQuest:      copygen.DiscordSourceHordeQuest,
}

// sourceKindWords are the page's badge words by kind (web/src/lib/sim/loot.ts
// SOURCE_KIND_LABELS), the line a kind loot.json has nothing for falls back to.
var sourceKindWords = map[string]string{
	"raid": copygen.NotifSectionRaids, "dungeon": copygen.DiscordSourceKindWordsDungeon, "world": copygen.DiscordSourceKindWordsWorld, "world_drop": copygen.DiscordSourceKindWordsWorldDrop,
	"zone": copygen.DiscordSourceKindWordsZone, "vendor": copygen.DiscordSourceKindWordsVendor, "crafted": copygen.ItemPageKindCrafted, "rep": copygen.ItemPageKindRep, "pvp": copygen.AccountRulesetPvp,
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
