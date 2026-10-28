package coach

import ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"

// Code names one of the eight findings (§6.2). The JSON names are the spec's.
type Code string

const (
	CodeCoreCasts   Code = "core_casts"
	CodeUptime      Code = "uptime"
	CodeIdle        Code = "idle"
	CodeResource    Code = "resource"
	CodeRanks       Code = "ranks"
	CodePreparation Code = "preparation"
	CodeDeath       Code = "death"
	CodeAdherence   Code = "adherence"
)

// Severity is a finding's tier: major ranks before minor (§6.1).
type Severity string

const (
	SeverityMajor Severity = "major"
	SeverityMinor Severity = "minor"
)

// BasisKind is where a measurement's yardstick comes from (§11).
type BasisKind string

const (
	BasisTopLogs   BasisKind = "top_logs"
	BasisSim       BasisKind = "sim"
	BasisSpellData BasisKind = "spell_data"
	BasisOwnLog    BasisKind = "own_log"
)

// Basis is what a finding is measured against, in plain words, with the pool size when
// the basis is the top logs.
type Basis struct {
	Label string    `json:"label"`
	Kind  BasisKind `json:"kind"`
	N     int       `json:"n"`
}

// NumberFact is one number a finding states: its value, the text it appears as, its
// unit and what it measures.
type NumberFact struct {
	V    float64 `json:"v"`
	Text string  `json:"text"`
	Unit string  `json:"unit"`
	Of   string  `json:"of"`
}

// SpellFact is one spell a finding names; the spell id is the rank.
type SpellFact struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

// Facts is everything a finding or strength may be quoted on (§7).
type Facts struct {
	Numbers []NumberFact `json:"numbers"`
	Spells  []SpellFact  `json:"spells"`
	Items   []string     `json:"items"`
	Times   []string     `json:"times"`
}

// Finding is one ranked pack finding (§5.2, §7).
type Finding struct {
	ID         string   `json:"id"`
	Code       Code     `json:"code"`
	Severity   Severity `json:"severity"`
	Rank       int      `json:"rank"`
	Title      string   `json:"title"`
	Detail     string   `json:"detail"`
	Basis      Basis    `json:"basis"`
	Comparison *Basis   `json:"comparison,omitempty"`
	Anchors    []string `json:"anchors"`
	Facts      Facts    `json:"facts"`
}

// Strength is one measurement at or above its comparison; the review's "Doing well"
// may use only these.
type Strength struct {
	ID       string   `json:"id"`
	Sentence string   `json:"sentence"`
	Facts    Facts    `json:"facts"`
	Anchors  []string `json:"anchors"`
}

// LinkKind is one of the deterministic connections between two findings (§7).
type LinkKind string

const (
	LinkGapOverlapsCap     LinkKind = "gap_overlaps_cap"
	LinkOutOfManaAtDeath   LinkKind = "oom_at_death"
	LinkDeathInGap         LinkKind = "death_in_gap"
	LinkOverhealOnMostCast LinkKind = "overheal_on_most_cast"
	LinkConsumableCovers   LinkKind = "consumable_covers"
)

// Link is a connection the review may state, only when its windows or ids meet.
type Link struct {
	ID       string    `json:"id"`
	Findings [2]string `json:"findings"`
	Sentence string    `json:"sentence"`
	Kind     LinkKind  `json:"kind"`
}

// AnchorKind is the report-page target an evidence anchor opens.
type AnchorKind string

const (
	AnchorCast     AnchorKind = "cast"
	AnchorAura     AnchorKind = "aura"
	AnchorDeath    AnchorKind = "death"
	AnchorResource AnchorKind = "resource"
	AnchorTime     AnchorKind = "time"
	AnchorBuffs    AnchorKind = "buffs"
	AnchorHealing  AnchorKind = "healing"
)

// Anchor maps an id to a report-page target: a Casts row, an aura row, a death, the
// resource graph or a time.
type Anchor struct {
	ID      string     `json:"id"`
	Kind    AnchorKind `json:"kind"`
	SpellID int64      `json:"spell_id,omitempty"`
	AtMS    *int64     `json:"at_ms,omitempty"`
	ToMS    *int64     `json:"to_ms,omitempty"`
	Death   string     `json:"death,omitempty"`
}

// Context is the fight as the model may see it: no character name, account, guild or
// realm (§7).
type Context struct {
	SpecLabel string `json:"spec"`
	Class     string `json:"class"`
	Role      string `json:"role"`
	Boss      string `json:"boss"`
	Size      int    `json:"size"`
	Kill      bool   `json:"kill"`
	Duration  string `json:"duration"`
	Level     int    `json:"level,omitempty"`
}

// Allowed is the validator's precomputed set: every number as it may appear, every name
// and every id. Each list is sorted and unique.
type Allowed struct {
	Numbers []string `json:"numbers"`
	Names   []string `json:"names"`
	IDs     []string `json:"ids"`
}

// Status is the pack's readiness.
type Status string

const (
	StatusReady         Status = "ready"
	StatusPreparing     Status = "preparing"
	StatusNotReviewable Status = "not_reviewable"
)

// Pack is the evidence pack (§7).
type Pack struct {
	Version       string     `json:"pack_version"`
	Status        Status     `json:"status"`
	NotReviewable string     `json:"not_reviewable,omitempty"`
	Context       Context    `json:"context"`
	Findings      []Finding  `json:"findings"`
	Strengths     []Strength `json:"strengths"`
	Links         []Link     `json:"links"`
	CouldNotCheck []string   `json:"couldNotCheck"`
	Anchors       []Anchor   `json:"anchors"`
	Allowed       Allowed    `json:"allowed"`
	BasisLabels   []string   `json:"basis_labels"`
}

// outcome is one check's answer for one player and one fight. Review (Task 14) numbers
// the ids: local anchors are referred to as "#0", "#1"… inside Finding.Anchors and
// Strength.Anchors.
type outcome struct {
	Finding       *Finding // nil: nothing fires
	Strengths     []Strength
	CouldNotCheck []string
	Anchors       []Anchor              // local, ID empty
	Windows       []ratingengine.Window // the finding's windows, for links (gaps, caps, out of mana, the death instant)
}

// check is one finding's rule (F1–F8), one function per file.
type check func(in Input, c Counted) outcome
