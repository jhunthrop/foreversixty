package discordbot

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
	CharLine:       "Level {n} {Spec} {Class} · {Guild}",
	CharFacts:      "-# Item level {x.x} · Readiness: {word} · Synced <t:{unix}:R>",
	NoGear:         "No gear shared yet. Upgrades show once the character is synced.",
	CharacterPage:  "Character page",
	BisOpen:        "Open on Forever Sixty",
	BisAllSlots:    "All 17 slots",
	ItemLineOne:    "{Slot} · {Quality} · item level {n}",
	ItemLineTwo:    "{delta} · {source}",
	Share:          "Share in channel",
	Shared:         "Shared",
	NoLongerPublic: "That is no longer public, so it was not shared.",
}

// panelCopy is the words of the gear panel and the shared sheet's message (spec 2026-10-29,
// "Copy, verbatim"); the tests read each string back out of that table.
var panelCopy = struct {
	BisHeadTitle, BisHeadLine, BisHeadPage string
	// BisPages are the two pages' ranges: Head to Legs, Feet to Ranged.
	BisPages                                      [2]string
	BisSlotTitle, BisSlotNext                     string
	BisOffHandTwoHand, BisSlotNone                string
	NotModelled                                   string
	Replaces, ReplacesNoLevel                     string
	UpgradesTitle, UpgradesByDPS, UpgradesByScore string
	UpgradesMore, UpgradesMoreOne, UpgradesNone   string
	ShareCharGold, ShareCharGoldTank              string
	TileAlt, TileAltNoQuality, CrestAlt           string
	// SheetRange is the /bis sheet's range under its level: the Phase 1 raid, level 60 bare,
	// and the levels of a lower band.
	SheetRangeRaid, SheetRangeBare, SheetRangeLevels string
	SheetUpdated, SheetCharGold, SheetCharSynced     string
}{
	BisHeadTitle:      "### {Spec} {Class}",
	BisHeadLine:       "Best in slot · {band label} · {Faction}",
	BisHeadPage:       "-# {Page range}, {n} of 17 slots. Each gain is over the next best item.",
	BisPages:          [2]string{"Head to Legs", "Feet to Ranged"},
	BisSlotTitle:      "### {Slot} · {Spec} {Class}",
	BisSlotNext:       "-# The next best",
	BisOffHandTwoHand: "**Off hand** · none: {main hand name} is two-handed.",
	BisSlotNone:       "**{Slot}** · no pick for this range.",
	NotModelled:       "Gain not modelled yet",
	Replaces:          "-# Replaces {worn name}, item level {n}",
	ReplacesNoLevel:   "-# Replaces {worn name}",
	UpgradesTitle:     "**Top upgrades** · {band label}",
	UpgradesByDPS:     "-# By DPS gained, against the guide build.",
	UpgradesByScore:   "-# By Forever Score gained, against the guide build.",
	UpgradesMore:      "-# {n} more upgrades on the character page.",
	UpgradesMoreOne:   "-# 1 more upgrade on the character page.",
	UpgradesNone:      "Every slot already wears the guide's pick for {band label}.",
	ShareCharGold:     "-# Gold figures: the {band label} pick's gain over what is worn.",
	ShareCharGoldTank: "-# Gold figures: the {band label} pick's Forever Score gain over what is worn.",
	TileAlt:           "{Item name}, {Quality} {slot} item",
	TileAltNoQuality:  "{Item name}, {slot} item",
	CrestAlt:          "{Class} crest",
	SheetRangeRaid:    "Phase 1 raid",
	SheetRangeBare:    "No buffs or consumables",
	SheetRangeLevels:  "Levels {n} to {n+9}",
	SheetUpdated:      "Updated {d Mon yyyy}",
	SheetCharGold:     "Gold: gain over what is worn, {band label}",
	SheetCharSynced:   "Synced {d Mon yyyy}",
}

// qualityWords are the item qualities in the game's words, by quality number (spec §6.2).
var qualityWords = []string{"Poor", "Common", "Uncommon", "Rare", "Epic", "Legendary"}

// bandLabels are /bis's band labels by phase choice value (spec §6.2).
var bandLabels = map[string]string{
	"20": "level 20", "30": "level 30", "40": "level 40", "50": "level 50",
	"60": "level 60", "p1": "level 60 Phase 1 raid",
}

// pointsWord is a gain in Forever Score points, the character page's own "+18.4 points".
const pointsWord = "points"
