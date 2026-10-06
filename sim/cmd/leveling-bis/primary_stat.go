package main

// primaryStatBySpec is THE single source of truth for which raw stat
// is this spec's primary stat for SimulationCraft-convention scale-
// factor anchoring (normalizeScaleFactors, weights.go) - owner ruling
// 2026-10-05 (bis-weights-simc lane, ranker-weights-anchor): Strength
// for warrior and paladin melee/tank specs and shaman enhancement;
// Agility for rogue, hunter (every spec, including survival, which
// fights in melee but is still a hunter) and feral druid; Intellect
// for every caster and healer spec (mage, warlock, priest, elemental
// shaman, holy paladin, restoration shaman/druid, balance druid).
//
// This table names the CANONICAL stat, not necessarily the exact
// weight_stats row id a given spec anchors on in practice -
// primaryAnchorStat (below) is the one function that turns this
// answer into an actual row id, applying the Intellect fallback this
// lane's brief calls for when a spec's own weight_stats carries no
// literal "intellect" row.
//
// Every spec in data/curated/specs.json must have an entry here -
// TestPrimaryStatCoversEverySpec (primary_stat_test.go) walks the
// real, published catalogue and fails if one is missing or is not
// one of the three canonical stats below.
var primaryStatBySpec = map[string]string{
	"druid-balance":        "intellect",
	"druid-feral":          "agility",
	"druid-restoration":    "intellect",
	"hunter-beast-mastery": "agility",
	"hunter-marksmanship":  "agility",
	"hunter-survival":      "agility",
	"mage-arcane":          "intellect",
	"mage-fire":            "intellect",
	"mage-frost":           "intellect",
	"paladin-holy":         "intellect",
	"paladin-protection":   "strength",
	"paladin-retribution":  "strength",
	"priest-discipline":    "intellect",
	"priest-holy":          "intellect",
	"priest-shadow":        "intellect",
	"rogue-assassination":  "agility",
	"rogue-combat":         "agility",
	"rogue-subtlety":       "agility",
	"shaman-elemental":     "intellect",
	"shaman-enhancement":   "strength",
	"shaman-restoration":   "intellect",
	"warlock-affliction":   "intellect",
	"warlock-demonology":   "intellect",
	"warlock-destruction":  "intellect",
	"warrior-arms":         "strength",
	"warrior-fury":         "strength",
	"warrior-protection":   "strength",
}

// intellectFallbackRows is this lane's brief, item 1's parenthetical:
// "if a healer or caster spec's own weight set has no intellect row,
// use its spell or healing power row and document it". Checked in
// order - healing_power first (a healer's own weight_stats array
// always leads with it, data/curated/specs.json), then spell_power (a
// caster's own primary damage-scaling stat). Not exercised by any
// spec in today's committed data - every Intellect-primary spec's own
// weight_stats already carries a literal "intellect" row (see
// TestPrimaryAnchorStatUsesIntellectDirectlyWhenPresent) - but kept so
// a future weight_stats edit that drops intellect degrades to a named
// row instead of normalizeScaleFactors' own "primary row absent"
// fallback (today's largest-significant-weight rule, which this
// lane's brief restricts to exactly that case).
var intellectFallbackRows = []string{"healing_power", "spell_power"}

// primaryAnchorStat turns primaryStatBySpec's canonical answer
// ("strength"/"agility"/"intellect") into the actual weight_stats row
// id normalizeScaleFactors should anchor the published scale-factor
// table to for spec - this lane's brief, item 1 and item 2.
//
// ok is false only when spec.Spec has no primaryStatBySpec entry at
// all (unreachable for a real spec given TestPrimaryStatCoversEvery
// Spec); the caller treats that exactly like normalizeScaleFactors'
// own "primary row absent" case, by passing it an empty string, which
// falls back to the pre-existing largest-significant-weight rule.
func primaryAnchorStat(spec specInfo) (stat string, ok bool) {
	primary, ok := primaryStatBySpec[spec.Spec]
	if !ok {
		return "", false
	}
	if containsStat(spec.WeightStats, primary) {
		return primary, true
	}
	if primary == "intellect" {
		for _, fallback := range intellectFallbackRows {
			if containsStat(spec.WeightStats, fallback) {
				return fallback, true
			}
		}
	}
	// Nothing in this spec's own weight_stats names the primary stat
	// or (for Intellect) either of its fallback rows - returning the
	// canonical id anyway (rather than "", false) lets
	// normalizeScaleFactors' own row search fail to find a match on
	// its own and fall back to the largest-significant-weight rule,
	// the same "primary row absent" path a stat that is not even in
	// weight_stats always takes.
	return primary, true
}

// containsStat is a plain linear membership check over a spec's own
// weight_stats list (data/curated/specs.json) - these lists are a
// handful of entries each, so a map is not worth the allocation.
func containsStat(stats []string, want string) bool {
	for _, s := range stats {
		if s == want {
			return true
		}
	}
	return false
}
