// api/internal/guilds/subscribe_step.go
package guilds

import "time"

// Step is the one next step the checkout guild step offers a viewer for a guild (spec
// 2026-10-27-subscribe-and-join §8). The web draws the card for the step it is given and
// never derives one from a rank.
type Step string

const (
	StepOnPlan            Step = "on_plan"
	StepCheckoutOpen      Step = "checkout_open"
	StepFrozen            Step = "frozen"
	StepReady             Step = "ready"
	StepClaimedAsk        Step = "claimed_ask"
	StepClaimPendingYours Step = "claim_pending_yours"
	StepConfirmClaim      Step = "confirm_claim"
	StepClaimPendingAsk   Step = "claim_pending_ask"
	StepClaim             Step = "claim"
	StepClaimLinkBnet     Step = "claim_link_bnet"
	StepStartClaim        Step = "start_claim"
	StepClaimBlocked      Step = "claim_blocked"
	StepAskLeader         Step = "ask_leader"
)

// The reasons a claim_blocked step carries, in the order Claim checks them.
const (
	BlockedHoldsOther  = "holds_other"
	BlockedRateLimited = "rate_limited"
)

// The claim states /v1/me reports for a guild.
const (
	ClaimStateUnclaimed = "unclaimed"
	ClaimStatePending   = "pending"
	ClaimStateClaimed   = "claimed"
	ClaimStateContested = "contested"
)

// The raw ranks a viewer's characters carry in a guild.
const (
	rankLeader  = "leader"
	rankOfficer = "officer"
)

// StepFacts is everything SubscribeStep reads, gathered by the caller from billing and the
// guild's claim. It holds facts, never decisions.
type StepFacts struct {
	OnLivePlan            bool       // a whole plan in trial, active, ending or past_due
	CheckoutOpen          bool       // billing.Store.PendingGuildCheckout
	ClaimState            string     // unclaimed | pending | claimed | contested
	ViewerHoldsClaim      bool       // the viewer's account holds the claim
	ViewerPendingClaim    bool       // the viewer started the pending claim
	ViewerRank            string     // highest raw rank in the guild: "leader" | "officer" | "member"
	ViewerVerifiedOfficer bool       // guild_members' verified rank is officer or leader (Store.Capabilities' read)
	HasBattleNet          bool       // a linked Battle.net identity
	HoldsOtherClaim       bool       // the account already holds another guild's claim
	OtherGuildName        string     // that other guild's name
	RateLimitedUntil      *time.Time // set when the 30-day claim limit applies
	TrialUsed             bool       // billing.Store.GuildTrialUsed
}

// StepResult is the step and the few values its card needs.
type StepResult struct {
	Step           Step
	TrialUsed      *bool      // set only on ready
	AvailableAt    *time.Time // set only on claim_blocked with rate_limited
	Reason         string     // "" | BlockedHoldsOther | BlockedRateLimited
	OtherGuildName string     // set only with holds_other
}

// stepRule is one step and the predicate that selects it.
type stepRule struct {
	step    Step
	matches func(StepFacts) bool
}

// stepRules is §8's first-match order. ask_leader is the fallback when none matches. It is
// read-only: nothing writes it after initialisation.
var stepRules = []stepRule{
	{StepOnPlan, func(f StepFacts) bool { return f.OnLivePlan }},
	{StepCheckoutOpen, func(f StepFacts) bool { return f.CheckoutOpen }},
	{StepFrozen, func(f StepFacts) bool { return f.ClaimState == ClaimStateContested }},
	{StepReady, isReady},
	{StepClaimedAsk, func(f StepFacts) bool { return f.ClaimState == ClaimStateClaimed }},
	{StepClaimPendingYours, func(f StepFacts) bool { return f.ClaimState == ClaimStatePending && f.ViewerPendingClaim }},
	{StepConfirmClaim, func(f StepFacts) bool { return f.ClaimState == ClaimStatePending && f.mayClaim() }},
	{StepClaimPendingAsk, func(f StepFacts) bool { return f.ClaimState == ClaimStatePending }},
	{StepClaim, func(f StepFacts) bool { return f.unclaimedLeader() && f.HasBattleNet && !f.claimBlocked() }},
	{StepClaimLinkBnet, func(f StepFacts) bool { return f.unclaimedLeader() && !f.HasBattleNet }},
	{StepStartClaim, func(f StepFacts) bool { return f.unclaimedClaimer() && !f.claimBlocked() }},
	{StepClaimBlocked, func(f StepFacts) bool { return f.unclaimedClaimer() && f.claimBlocked() }},
}

// SubscribeStep resolves the one step for a viewer and a guild, first match in §8 order.
func SubscribeStep(f StepFacts) StepResult {
	for _, rule := range stepRules {
		if rule.matches(f) {
			return resultFor(rule.step, f)
		}
	}
	return StepResult{Step: StepAskLeader}
}

// isReady is a claimed guild the viewer may bill, by the rule Capabilities' billing word uses.
// A claim holder who is not a verified officer falls through to claimed_ask (spec §9.1, the
// claim.yours row).
func isReady(f StepFacts) bool {
	return f.ClaimState == ClaimStateClaimed && mayBill(f.ViewerHoldsClaim, f.ViewerVerifiedOfficer)
}

// mayClaim is eligibleClaimRank's rule: an officer or leader character, verified or not.
func (f StepFacts) mayClaim() bool {
	return f.ViewerRank == rankOfficer || f.ViewerRank == rankLeader
}

func (f StepFacts) unclaimedLeader() bool {
	return f.ClaimState == ClaimStateUnclaimed && f.ViewerRank == rankLeader
}

func (f StepFacts) unclaimedClaimer() bool {
	return f.ClaimState == ClaimStateUnclaimed && f.mayClaim()
}

// claimBlocked is checkClaimRateLimit's two refusals as facts.
func (f StepFacts) claimBlocked() bool {
	return f.HoldsOtherClaim || f.RateLimitedUntil != nil
}

// resultFor fills the values the step's card reads, and only those.
func resultFor(step Step, f StepFacts) StepResult {
	switch step {
	case StepReady:
		used := f.TrialUsed
		return StepResult{Step: step, TrialUsed: &used}
	case StepClaimBlocked:
		return blockedResult(f)
	default:
		return StepResult{Step: step}
	}
}

// blockedResult names holds_other before rate_limited, the order checkClaimRateLimit uses:
// a viewer who still holds a guild hears about that, not a generic rate limit.
func blockedResult(f StepFacts) StepResult {
	if f.HoldsOtherClaim {
		return StepResult{Step: StepClaimBlocked, Reason: BlockedHoldsOther, OtherGuildName: f.OtherGuildName}
	}
	until := *f.RateLimitedUntil
	return StepResult{Step: StepClaimBlocked, Reason: BlockedRateLimited, AvailableAt: &until}
}
