// logs/engine/rating/rating.go
// Package rating scores one player's performance for one closed fight
// against the six components docs/superpowers/specs/2026-09-21-performance-
// rating-design.md §1 defines: pure functions over an already-finished
// summary.Summary plus curated tables, with no network or database access,
// so a unit test can hand it a hand-built fixture and a fake percentile
// source. Sibling to logs/engine/summary and logs/engine/mechanics, never
// imported by either (§4.1's own ruling: a rating is a post-fight
// computation, not something summary.Accumulator.Snapshot needs to finish
// mid-fight).
package rating

import (
	"math"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/utility"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/weights"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The six component names, exactly as they appear in Component.Name and
// the site's public API (spec §4.1, §5.2).
const (
	ComponentNameOutput      = "output"
	ComponentNameSurvival    = "survival"
	ComponentNameMechanics   = "mechanics"
	ComponentNameUtility     = "utility"
	ComponentNamePreparation = "preparation"
	ComponentNameActivity    = "activity"
)

// A component's basis (spec §1.2): which of the two rules produced its
// score, or "" when the component is excluded.
const (
	BasisPercentile = "percentile"
	BasisAbsolute   = "absolute"
)

// The three roles summary.RosterRow.Role already produces (roster.go's own
// role() function), reused here rather than a second vocabulary.
const (
	RoleDPS    = "dps"
	RoleHealer = "healer"
	RoleTank   = "tank"
)

// The three kill-time bands, spec §1.2.
const (
	BandFast    = "fast"
	BandTypical = "typical"
	BandSlow    = "slow"
)

// MinSample is spec §1.2's RULING: a percentile computed from fewer than
// this many kills is a percentile of noise dressed as a number, so the
// absolute standard is used instead until a bracket crosses it.
const MinSample = 20

// DefaultCapThreshold is spec §1.5's site-wide default: the score an
// avoidable-death-triggered cap holds Overall at. A guild-adjustable
// setting (out of scope here, officer tooling); ModelInfo carries whatever
// value a caller wants — this is only the default a caller not overriding
// it should pass.
const DefaultCapThreshold = 40.0

// DefaultModelVersion is this spec's own model version stamp (spec §4.4:
// "a short string... bumped whenever a formula, a weight, or any curated
// table changes in a way that could change a stored score").
const DefaultModelVersion = "rating-2026-09-21"

// The fixed machine-readable exclusion reasons Component.Reason carries
// (spec §5.2's own JSON example shows "no_mechanics_table" in exactly this
// shape; the web lane maps these to spec §6.4's reader-facing copy —
// RULING R10 in docs/superpowers/plans/2026-09-21-rating-engine.md: the
// spec's own "§7.4" cross-references for this copy are a drafting slip for
// "§6.4", the section that actually lists the fixed strings).
const (
	ReasonNoMechanicsTable      = "no_mechanics_table"
	ReasonNoUtilityTable        = "no_utility_table"
	ReasonNoConsumableCatalogue = "no_consumable_catalogue"
	ReasonWipe                  = "wipe"
	ReasonNotEnoughSamples      = "not_enough_samples"
)

// The Bracket.Component values for Mechanics' and Survival's independently-
// percentiled sub-parts, distinct from the six top-level component names
// since a single top-level component can combine more than one percentile
// query (RULING R9).
const (
	ComponentSurvivalAvoidableHit  = "survival_avoidable_hit"
	ComponentMechanicsInterrupt    = "mechanics_interrupt"
	ComponentMechanicsDispel       = "mechanics_dispel"
	ComponentMechanicsDebuffUptime = "mechanics_debuff_uptime"
)

// Card is one player's rating for one fight (spec §4.1, verbatim).
type Card struct {
	// Overall is the site-default figure: OverallUncapped, with §1.5's cap
	// applied if OverallCapped fired. OverallUncapped is always stored too,
	// so a guild-adjustable setting that turns the cap off — or a future
	// recompute of the cap's own threshold — reads it straight back with
	// no backfill (§1.5, §4.2).
	Overall         float64
	OverallUncapped float64
	OverallCapped   bool // §1.5's Survival-catastrophe cap condition fired this fight
	Components      [6]Component
	Basis           string // "percentile" | "absolute" | "mixed" (per-component; see Component.Basis)
	ModelVersion    string
	KillTimeBand    string // "fast" | "typical" | "slow" | "" (excluded/wipe)
}

// Component is one of the six parts of a Card (spec §4.1, verbatim).
type Component struct {
	Name       string   // "output" | "survival" | "mechanics" | "utility" | "preparation" | "activity"
	Score      float64  // 0-100; meaningless if Excluded
	Weight     float64  // the *renormalised* weight actually applied
	Basis      string   // "percentile" | "absolute" | ""
	Percentile *float64
	BracketN   int64
	Excluded   bool
	Reason     string   // set iff Excluded; one of the Reason* constants
	Moments    []Moment // the "opens into the specific moments" data; see §7.1
}

// Moment is one specific instant a component's score is built from — a
// death, an avoidable hit, an interrupt made, a dispel made, a dropped
// debuff — the "opens into the moments in the log" data spec §6.1/§7.1
// describe. Anchor, when set, is a DOM anchor id an existing tab already
// renders (spec §6.1's "Jumping to a moment"): "death-<guid>-<at_ms>" for a
// death, matching DeathsTab.svelte's own convention.
type Moment struct {
	Kind      string
	AtMS      int64
	SpellID   int64
	SpellName string
	Avoidable bool
	Anchor    string
}

// Assignment is an officer-marked job for one player on one fight, supplied
// later by officer tooling that does not exist yet (§9). Score's
// assignments parameter is always an empty slice today — no caller
// populates it — but the signature accepts it now so officer tooling can
// be built against a stable contract without a model change once it
// exists. (Spec §2, verbatim.)
type Assignment struct {
	PlayerKey string
	// Job is the officer's own free-text label for what this window covers
	// — "decurse", "interrupt rotation", "add duty", "kiting" — shown back
	// on the card wherever the exemption applies, never interpreted by the
	// engine.
	Job    string
	FromMS int64
	ToMS   int64 // exclusive
}

// ModelInfo is every guild-adjustable setting Score needs, bundled so
// Score's signature does not grow every time officer tooling adds one
// (spec §4.1).
type ModelInfo struct {
	ModelVersion string
	Weights      weights.Roles
	// CapEnabled and CapThreshold are §1.5's cap: a guild may turn it off
	// entirely, or (in principle) recompute its threshold. A caller wanting
	// the site default passes CapEnabled: true, CapThreshold:
	// DefaultCapThreshold.
	CapEnabled   bool
	CapThreshold float64
}

// DefaultModelInfo is the site-wide default a caller not implementing a
// guild override should pass.
func DefaultModelInfo() ModelInfo {
	return ModelInfo{
		ModelVersion: DefaultModelVersion,
		Weights:      weights.Default(),
		CapEnabled:   true,
		CapThreshold: DefaultCapThreshold,
	}
}

// CuratedTables bundles the curated data a fight's rating reads, already
// selected for this player's encounter/spec/role by the caller: Mechanics
// is the fight's own encounter table (nil if none curated yet), Utility is
// this player's spec's owned-utility table (nil if none curated for this
// spec — never expected in practice, all 27 specs have a file, some
// sparse), Consumables is this player's role's consumable catalogue row
// (nil if the caller has not loaded one).
type CuratedTables struct {
	Mechanics   *mechanics.Table
	Utility     *utility.Table
	Consumables *consumables.RoleCatalogue
}

// Bracket is the percentile-digest key spec §1.2 defines:
// (encounter_id, difficulty, spec, role, kill_time_band, component), plus
// Kill (spec §2's wipe rule: a wipe's digests are kept apart from kills').
// Component is one of the six top-level component names or one of the four
// sub-part names above (RULING R9).
type Bracket struct {
	EncounterID  int64
	Difficulty   int64
	Spec         string
	Role         string
	KillTimeBand string
	Kill         bool
	Component    string
}

// PercentileSource is the small interface Score reads distributions
// through (spec §4.1), so logs/engine never depends on api/internal: the
// API layer wires a concrete implementation backed by
// api/internal/digest.Unmarshal/Placement over rows read from
// rating_percentile_digests.
type PercentileSource interface {
	// Placement is spec §1.2's percentile-within-bracket lookup: the share
	// of the bracket's other values this value beats, in 0..1, and how many
	// kills the bracket has. ok is false when the source has nothing for
	// this bracket at all (a brand-new encounter/spec/role combination).
	Placement(bracket Bracket, value float64) (pct float64, n int64, ok bool)
	// KillTimeBand classifies durationMS against the (encounterID,
	// difficulty) bracket's own trailing kill-duration median (spec §1.2).
	// ok is false when the source has no kill-duration digest for this
	// bracket yet, in which case the caller uses BandTypical — spec §1.2's
	// own note that an unclassifiable band "degenerates to a no-op, not a
	// wrong answer" (RULING R3: added to the interface — spec §4.1 names
	// only Placement, which alone cannot answer a median-duration query).
	KillTimeBand(encounterID, difficulty int64, durationMS int64) (band string, ok bool)
}

// rosterRow finds player's row in the fight's roster.
func rosterRow(fight summary.Summary, player string) (summary.RosterRow, bool) {
	for _, r := range fight.Roster {
		if r.GUID == player {
			return r, true
		}
	}
	return summary.RosterRow{}, false
}

// round2 rounds to two decimal places: spec §1.5's storage precision ("the
// stored value keeps two decimal places... so a recompute is stable to the
// same input").
func round2(x float64) float64 {
	return roundTo(x, 2)
}

// roundTo rounds x to n decimal places, half away from zero — spec §1.5
// does not specify a tie-breaking rule, so math.Round's ordinary
// round-half-away-from-zero is used rather than a hand-rolled truncation
// (which mis-rounds negative values, since Go's int64 conversion truncates
// toward zero, not toward negative infinity).
func roundTo(x float64, n int) float64 {
	m := math.Pow(10, float64(n))
	return math.Round(x*m) / m
}

// specSlug turns a roster row's display class/spec (e.g. "Warrior",
// "Protection") into the slug data/curated/specs.json and
// logs/engine/mechanics/utility's file names use (e.g.
// "warrior-protection"). Every one of the 27 specs in specs.json follows
// exactly this "<lower class>-<lower spec, spaces to hyphens>" pattern
// (checked against specs.json's own "spec" field for all 27 rows), so this
// is a direct transform, not a lookup table this package would otherwise
// need to embed a copy of specs.json just to build.
func specSlug(class, spec string) string {
	return toSlug(class) + "-" + toSlug(spec)
}

func toSlug(s string) string {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == ' ':
			b = append(b, '-')
		case c >= 'A' && c <= 'Z':
			b = append(b, c+('a'-'A'))
		default:
			b = append(b, c)
		}
	}
	return string(b)
}
