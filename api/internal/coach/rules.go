package coach

// RulesVersion names the threshold set below; a change to any value moves it.
const RulesVersion = "coach-rules-v1"

// The §6.1 and §6.2 v1 values (docs/superpowers/specs/2026-10-27-log-coach-design.md).
const (
	// Shared rules (§6.1).
	MinFightMS            = 30_000 // a reviewable fight is at least this long
	WipeTailMS            = 15_000 // on a wipe, deaths this close to the end are part of the wipe
	MaxFindings           = 8      // one finding per type
	MoreAbilitiesInDetail = 2      // F1 and F2 name up to this many more in the detail

	// F1 core casts.
	CoreCooldownMinMS     = 4_000
	CoreCooldownMaxMS     = 600_000
	CoreGapGraceMS        = 1_500
	ResourceWholeSecondMS = 1_000
	CoreFiresLost         = 2
	CoreFiresPct          = 10.0
	CoreMajorPct          = 20.0

	// F2 uptime.
	UptimeFiresPoints   = 10.0
	UptimeFiresUnderPct = 90.0
	UptimeMajorPoints   = 25.0

	// F3 dead time.
	IdleFiresMS      = 10_000
	IdleFiresPct     = 5.0
	IdleMajorPct     = 10.0
	IdleWindowsShown = 3

	// F4 resource waste.
	CapFiresMS         = 10_000
	CapFiresPct        = 5.0
	CapMajorPct        = 15.0
	ManaLowPct         = 5.0
	ManaFiresMS        = 10_000
	ManaFiresBeforePct = 90.0
	ManaMajorBeforePct = 75.0

	// F5 ranks.
	RanksFiresCasts     = 3
	RanksFiresPct       = 10.0
	HealRankMinCasts    = 10
	HealRankFiresPoints = 15.0
	HealRankMajorPoints = 30.0

	// F7 death review.
	DeathFiresBeforePct = 90.0
	DeathMajorBeforePct = 75.0
	DeathPotionWindowMS = 10_000

	// F8 rotation adherence.
	AdherenceFiresPoints   = 10.0
	AdherenceFiresUnderPct = 85.0
	AdherenceMajorPoints   = 25.0
	OffRotationShown       = 3

	// The pool (§11.1).
	PoolMinPlayers = 20
	PoolTopN       = 25
)

// Roles as summary.RosterRow.Role spells them.
const (
	RoleDPS    = "dps"
	RoleHealer = "healer"
	RoleTank   = "tank"
)

// Precedence is the order within a severity tier (§6.3). A role's list holds only the
// findings that role gets: a healer has no idle and no adherence.
var Precedence = map[string][]Code{
	RoleDPS:    {CodeDeath, CodeCoreCasts, CodeIdle, CodePreparation, CodeUptime, CodeResource, CodeAdherence, CodeRanks},
	RoleHealer: {CodeDeath, CodeResource, CodeRanks, CodeCoreCasts, CodePreparation, CodeUptime},
	RoleTank:   {CodeDeath, CodeUptime, CodeCoreCasts, CodeIdle, CodePreparation, CodeResource, CodeAdherence, CodeRanks},
}
