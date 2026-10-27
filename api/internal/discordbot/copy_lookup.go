package discordbot

// lookupCopy is the words of /char, /bis, the item embed and Share in channel (experience spec
// §6.2); the tests read each string back out of the spec.
var lookupCopy = struct {
	// CharLine is the card's description; the guild clause follows only when there is one.
	CharLine string
	// The card's field names.
	ItemLevel, Readiness, Synced, TopUpgrades string
	// UpgradesBand follows TopUpgrades: the band the gains are against.
	UpgradesBand string
	// UpgradesByDPS and UpgradesByScore say what the upgrades are ranked by (tanks by score).
	UpgradesByDPS, UpgradesByScore string
	// NoGear stands in the upgrades field when the character has no gear shared.
	NoGear string
	// CharacterPage is the card's link button; NoCharacter answers an unknown or private name.
	CharacterPage, NoCharacter string
	// BisSlotLead and BisTopLead are /bis's text lines; BisAuthor heads the pick.
	BisSlotLead, BisTopLead, BisAuthor string
	// BisOpen and BisAllSlots are the link buttons with and without a slot.
	BisOpen, BisAllSlots string
	// BisNeedsSpec answers a /bis with nothing to resolve the spec from.
	BisNeedsSpec string
	// BisFooter closes a /bis answer.
	BisFooter string
	// ItemLineOne and ItemLineTwo are an item embed's description.
	ItemLineOne, ItemLineTwo string
	// Footer closes /char (spec §3.3.0).
	Footer string
	// Share is the button; Shared its label after a share; NoLongerPublic a share with nothing left.
	Share, Shared, NoLongerPublic string
}{
	CharLine:        "Level {n} {Spec} {Class} · {Guild}",
	ItemLevel:       "Item level",
	Readiness:       "Readiness",
	Synced:          "Synced",
	TopUpgrades:     "Top 3 upgrades",
	UpgradesBand:    " · {band label}",
	UpgradesByDPS:   "By DPS gained, against the guide build.",
	UpgradesByScore: "By Forever Score gained, against the guide build.",
	NoGear:          "No gear shared yet. Upgrades show once the character is synced.",
	CharacterPage:   "Character page",
	NoCharacter:     "No public character by that name.",
	BisSlotLead:     "{Slot} for a {Spec} {Class}, {band label}:",
	BisTopLead:      "{Spec} {Class}, {band label}. The five slots where the best item gains the most over the next best:",
	BisAuthor:       "Best in slot",
	BisOpen:         "Open on Forever Sixty",
	BisAllSlots:     "All 17 slots",
	BisNeedsSpec:    "Name a class and spec, or link Discord to use your character.",
	BisFooter:       "{Spec} {Class} · {band label} · Forever Sixty",
	ItemLineOne:     "{Slot} · {Quality} · item level {n}",
	ItemLineTwo:     "{delta} · {source}",
	Footer:          "Forever Sixty",
	Share:           "Share in channel",
	Shared:          "Shared",
	NoLongerPublic:  "That is no longer public, so it was not shared.",
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
