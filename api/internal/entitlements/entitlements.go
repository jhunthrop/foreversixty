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
	// (spec RULING 3). A guild grant may cover this feature alone
	// (FeatureScopeOfficerViews); Store.CanForGuild answers it per guild.
	FeatureOfficerViews Feature = "officer_views"
	FeatureRosterCheck  Feature = "roster_check"
	// FeatureRaidGuildTools is the Raid guild tier's per-guild tools: a
	// verified officer/leader of a guild whose active row covers officer
	// views and carries tier raid_guild (Store.CanForGuild; Can refuses it).
	// It is officer-only by construction: a member-level "this guild is on
	// raid_guild" check reads GuildBilling.Tier, not CanForGuild.
	FeatureRaidGuildTools Feature = "raid_guild_tools"
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

// TierGuild and TierRaidGuild are the two values entitlements.tier holds
// on a guild-plan row (null on premium rows).
const (
	TierGuild     = "guild"
	TierRaidGuild = "raid_guild"
)

// AnomalyKind is entitlement_anomalies.kind (migration 0020, security
// review 2026-09-21's double-billing fix).
type AnomalyKind string

const (
	// AnomalyDuplicateSubscription is UpsertStripe finding an already-
	// active row for a (subject, plan) pointing at a different
	// stripe_subscription_id: the row is left untouched and the newcomer
	// is canceled at Stripe instead of overwriting it.
	AnomalyDuplicateSubscription AnomalyKind = "duplicate_subscription"
	// AnomalyOrphanSubscription is stripe-reconcile finding a live Stripe
	// subscription for one of our own products with no matching
	// entitlements row at all (spec §2.8's two-directional check).
	AnomalyOrphanSubscription AnomalyKind = "orphan_subscription"
	// AnomalyUnexpectedTransfer is a transfer subscription with no offered
	// transfer behind it: the old subscription keeps billing.
	AnomalyUnexpectedTransfer AnomalyKind = "unexpected_transfer"
	// AnomalyUnknownPrice is a subscription whose price is not in the
	// catalog for its plan.
	AnomalyUnknownPrice AnomalyKind = "unknown_price"
)

// StatusGrantsAccess reports whether a Stripe subscription status counts
// as having the plan: active, trialing, and past_due (dunning keeps access
// while Stripe retries the card). Store's entitlement SQL grants on the
// same three, and billing's checkout refuses a second subscription while
// one of them stands.
func StatusGrantsAccess(status string) bool {
	switch status {
	case "active", "trialing", "past_due":
		return true
	}
	return false
}

// Standard reports whether f resolves from the account's own Premium (a
// personal plan or a held seat) rather than per guild.
func (f Feature) Standard() bool {
	return f != FeatureOfficerViews && f != FeatureRaidGuildTools
}
