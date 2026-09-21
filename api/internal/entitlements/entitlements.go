// api/internal/entitlements/entitlements.go
//
// Package entitlements answers "may this account use this feature right
// now" from the entitlements table (migration 0020), and nothing else: it
// has no Stripe dependency, no HTTP handlers, and no opinion about how a
// row got there (a webhook, a grant, or the users.premium backfill all
// look identical from here). See
// docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md §1.
package entitlements

// Feature is one gated capability. The web/API lanes both spell these as
// the exact snake_case strings below (spec §1.3) — JSON on EntitlementsView
// (auth package) uses the same spelling.
type Feature string

const (
	FeatureServerSims    Feature = "server_sims"
	FeatureRetention     Feature = "retention"
	FeatureMultiCompare  Feature = "multi_compare"
	FeatureHistory       Feature = "history"
	FeatureNotifications Feature = "notifications"
	// FeatureOfficerViews is guild-plan-and-verified-officer/leader only,
	// never unlockable by an individual's own premium subscription
	// (spec RULING 3).
	FeatureOfficerViews Feature = "officer_views"
	FeatureRosterCheck  Feature = "roster_check"
)

// Reason is why Can answered the way it did — informational for the
// caller's own error body/copy, never a branch the caller inspects
// beyond the boolean itself (spec §1.3).
type Reason string

const (
	ReasonEntitled       Reason = "entitled"
	ReasonSignInRequired Reason = "sign_in_required"
	ReasonNoPlan         Reason = "no_plan"
	ReasonNotOfficer     Reason = "not_officer"
)

// PlanPremium and PlanGuild are the two values entitlements.plan holds.
const (
	PlanPremium = "premium"
	PlanGuild   = "guild"
)

// verifiedGuildRanks are the two guild_members.rank values spec §0
// calls a guild's VerifiedGuildRank.
var verifiedGuildRanks = []string{"officer", "leader"}
