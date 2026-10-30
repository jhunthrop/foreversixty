package app

import "time"

// The home answer's states (app spec §5.2 F2): ok, or why there is no comparison to show.
const (
	StateOK      = "ok"
	StateNoGear  = "no_gear"
	StateUnder20 = "under_20"
	StateNoSpec  = "no_spec"
	StateNoList  = "no_list"
)

// The talents status (web/src/lib/home/talent-delta.ts's band compare).
const (
	TalentsOptimized = "optimized"
	TalentsDiffer    = "differs"
	TalentsNoExport  = "no_export"
	TalentsNoBuild   = "no_build"
)

// WornNotInData is a worn item's label when the build's item data lacks it.
const WornNotInData = "Not in our item data yet"

// Home is GET /v1/characters/{region}/{ruleset}/{name}/home: the character's upgrades for its
// level range, what is already best in slot, its talents against the band's build and when
// it last synced. Every array is present, never null.
type Home struct {
	// Character is the character key.
	Character string `json:"character"`
	State     string `json:"state"`
	// Spec is the spec the band is for (the site's key, "hunter-beast-mastery"); null when
	// the character has no spec the data knows.
	Spec *string `json:"spec"`
	// Band is the level band, BandLabel its words ("20 to 29", "60") and Preset what it was
	// ranked under ("raid", "bare"); all null when no band applies.
	Band      *int    `json:"band"`
	BandLabel *string `json:"band_label"`
	Preset    *string `json:"preset"`
	// Unit is what every gain counts ("dps", "hps", "tank_score") and UnitWord how a figure
	// reads it ("DPS", "HPS", "score"); null without a band.
	Unit     *string `json:"unit"`
	UnitWord *string `json:"unit_word"`
	// DataBuild is the build the item ids and icons belong to.
	DataBuild string `json:"data_build"`
	// UpgradeCount is len(Upgrades). TotalGain is the sum of every gain with a figure, one
	// decimal; null when it is not above 0 (the site shows no gain figure then).
	UpgradeCount int      `json:"upgrade_count"`
	TotalGain    *float64 `json:"total_gain"`
	// NotSimCheckedCount counts weapons with no listed alternative; NotModelledCount every
	// upgrade whose words are not a figure.
	NotSimCheckedCount int `json:"not_sim_checked_count"`
	NotModelledCount   int `json:"not_modelled_count"`
	// SyncedAt is when the stored export was true in the game and SyncedFrom where it came
	// from ("addon", "blizzard"); both null with no export.
	SyncedAt    *time.Time `json:"synced_at"`
	SyncedFrom  *string    `json:"synced_from"`
	Upgrades    []Upgrade  `json:"upgrades"`
	AlreadyBest []Best     `json:"already_best"`
	Talents     Talents    `json:"talents"`
}

// Item is an item as the build's table knows it; Name, Quality, Icon and ItemLevel are null
// when the table lacks the item. Icon is the icon id: the image is
// /data/{data_build}/icons/{icon}.webp on the site.
type Item struct {
	ID        int     `json:"id"`
	Name      *string `json:"name"`
	Quality   *int    `json:"quality"`
	Icon      *string `json:"icon"`
	ItemLevel *int    `json:"item_level"`
}

// Pick is the band's pick with its short source line.
type Pick struct {
	Item
	SourceLine string `json:"source_line"`
}

// Worn is the worn item. Known is false when the item data lacks it; Label is its name, or
// WornNotInData.
type Worn struct {
	Item
	Known bool   `json:"known"`
	Label string `json:"label"`
}

// Upgrade is one slot whose pick beats what is worn, in the order the apps show them: empty
// slots, then measured gains largest first, then the rest in frame order.
type Upgrade struct {
	Slot      string `json:"slot"`
	SlotLabel string `json:"slot_label"`
	Pick      Pick   `json:"pick"`
	// Worn is null for an empty slot.
	Worn *Worn `json:"worn"`
	// Gain is the pick's gain in Unit, one decimal, null when there is no figure. GainWords
	// is what the row reads: the signed figure, "Gain not modelled yet" (no figure, or a pick
	// whose effect the sim leaves out) or "not sim-checked". Modelled is true when GainWords
	// is the figure.
	Gain      *float64 `json:"gain"`
	GainWords string   `json:"gain_words"`
	Modelled  bool     `json:"modelled"`
}

// Best is a slot already wearing the pick, or a measured tie with it.
type Best struct {
	Slot      string `json:"slot"`
	SlotLabel string `json:"slot_label"`
	Item      Item   `json:"item"`
}

// Talents is the worn talents against the band's build: PointsDiffer of the band's
// BandPoints points are unmatched (Σ max(band − worn, 0)). Both null unless State is
// optimized or differs.
type Talents struct {
	State        string `json:"state"`
	PointsDiffer *int   `json:"points_differ"`
	BandPoints   *int   `json:"band_points"`
}
