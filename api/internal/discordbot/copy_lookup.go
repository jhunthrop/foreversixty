package discordbot

import (
	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
)

// lookupCopy is the words of /char, /bis, the item embed and Share in channel (experience spec
// §6.2); the tests read each string back out of the spec.
var lookupCopy = struct {
	// CharLine is the card's level line; the guild clause follows only when there is one.
	CharLine string
	// CharFacts is the /char header's subtext: item level, readiness and synced, each clause
	// only when known (spec 2026-10-29 copy table, char.facts).
	CharFacts string
	// NoGear stands under Top upgrades when the character has no gear shared.
	NoGear string
	// CharacterPage is the card's link button.
	CharacterPage string
	// BisOpen and BisAllSlots are the link buttons with and without a slot.
	BisOpen, BisAllSlots string
	// ItemLineOne and ItemLineTwo are an item embed's description (/report keeps the embed).
	ItemLineOne, ItemLineTwo string
	// Share is the button; Shared its label after a share; NoLongerPublic a share with nothing left.
	Share, Shared, NoLongerPublic string
}{
	CharLine:       copygen.DiscordLookupCharLine,
	CharFacts:      copygen.DiscordLookupCharFacts,
	NoGear:         copygen.DiscordLookupNoGear,
	CharacterPage:  copygen.DiscordLookupCharacterPage,
	BisOpen:        copygen.DiscordLookupBisOpen,
	BisAllSlots:    copygen.HomeUpgradesAllSlots,
	ItemLineOne:    copygen.DiscordLookupItemLineOne,
	ItemLineTwo:    copygen.DiscordLookupItemLineTwo,
	Share:          copygen.DiscordLookupShare,
	Shared:         copygen.DiscordLookupShared,
	NoLongerPublic: copygen.DiscordLookupNoLongerPublic,
}

// panelCopy is the words of the gear panel and the shared sheet's message (spec 2026-10-29,
// "Copy, verbatim"); the tests read each string back out of that table.
var panelCopy = struct {
	BisHeadTitle, BisHeadLine, BisHeadPage string
	// BisPages are the two pages' ranges: Head to Legs, Feet to Ranged.
	BisPages                  [2]string
	BisSlotTitle, BisSlotNext string
	// BisAltDelta is an alternative\'s line-2 figure: its own measured gap to the pick.
	BisAltDelta                    string
	BisOffHandTwoHand, BisSlotNone string
	NotModelled                    string
	// BisTie and BisOnly fill a pick's gain when there is no figure to state: a tie with the
	// next best ({unit} is DPS or points), and a slot with nothing else ranked.
	BisTie, BisOnly                               string
	Replaces, ReplacesNoLevel                     string
	UpgradesTitle, UpgradesByDPS, UpgradesByScore string
	// UpgradesByLevel heads a healer's or tank's list, never ranked by DPS; LevelGain,
	// LevelGainOne, LevelEmpty and LevelOnList are its rows' gain words.
	UpgradesByLevel, LevelGain, LevelGainOne, LevelEmpty, LevelOnList string
	// AllUpgrades labels the character page link when the panel shows only some of them.
	AllUpgrades, UpgradesNone           string
	TileAlt, TileAltNoQuality, CrestAlt string
	// SheetRange is the /bis sheet's range under its level: the Phase 1 raid, level 60 bare,
	// and the levels of a lower band.
	SheetRangeRaid, SheetRangeBare, SheetRangeLevels string
	SheetUpdated, SheetCharFooter                    string
}{
	BisHeadTitle:      copygen.DiscordPanelBisHeadTitle,
	BisHeadLine:       copygen.DiscordPanelBisHeadLine,
	BisHeadPage:       copygen.DiscordPanelBisHeadPage,
	BisPages:          [2]string{copygen.DiscordPanelBisPages1, copygen.DiscordPanelBisPages2},
	BisSlotTitle:      copygen.DiscordPanelBisSlotTitle,
	BisAltDelta:       copygen.DiscordPanelBisAltDelta,
	BisSlotNext:       copygen.DiscordPanelBisSlotNext,
	BisOffHandTwoHand: copygen.DiscordPanelBisOffHandTwoHand,
	BisSlotNone:       copygen.DiscordPanelBisSlotNone,
	NotModelled:       upgrades.NotModelledWords,
	BisTie:            copygen.DiscordPanelBisTie,
	BisOnly:           copygen.TooltipNothingElse,
	Replaces:          copygen.DiscordPanelReplaces,
	ReplacesNoLevel:   copygen.DiscordPanelReplacesNoLevel,
	UpgradesTitle:     copygen.DiscordPanelUpgradesTitle,
	UpgradesByDPS:     "",
	UpgradesByScore:   copygen.DiscordPanelUpgradesByScore,
	UpgradesByLevel:   copygen.DiscordPanelUpgradesByLevel,
	LevelGain:         copygen.DiscordPanelLevelGain,
	LevelGainOne:      copygen.DiscordPanelLevelGainOne,
	LevelEmpty:        copygen.DiscordPanelLevelEmpty,
	LevelOnList:       copygen.DiscordPanelLevelOnList,
	AllUpgrades:       copygen.DiscordPanelAllUpgrades,
	UpgradesNone:      copygen.DiscordPanelUpgradesNone,
	TileAlt:           copygen.DiscordPanelTileAlt,
	TileAltNoQuality:  copygen.DiscordPanelTileAltNoQuality,
	CrestAlt:          copygen.DiscordPanelCrestAlt,
	SheetRangeRaid:    copygen.ItemPagePresetRaid,
	SheetRangeBare:    copygen.DiscordPanelSheetRangeBare,
	SheetRangeLevels:  copygen.DiscordPanelSheetRangeLevels,
	SheetUpdated:      copygen.DiscordPanelSheetUpdated,
	SheetCharFooter:   copygen.DiscordPanelSheetCharFooter,
}

// bandLabels are /bis's band labels by phase choice value (spec §6.2).
var bandLabels = map[string]string{
	"20": copygen.DiscordBandLabelsK20, "30": copygen.DiscordBandLabelsK30, "40": copygen.DiscordBandLabelsK40, "50": copygen.DiscordBandLabelsK50,
	"60": copygen.DiscordBandLabelsK60, "p1": copygen.DiscordBandLabelsP1,
}

// pointsWord is a gain in Forever Score points, the character page's own "+18.4 points".
const pointsWord = copygen.DiscordLookupPointsWord
