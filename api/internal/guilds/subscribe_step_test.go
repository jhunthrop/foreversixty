// api/internal/guilds/subscribe_step_test.go
package guilds

import (
	"testing"
	"time"
)

var stepRateLimitedUntil = time.Date(2026, 11, 18, 0, 0, 0, 0, time.UTC)

// unclaimedFacts is a viewer in an unclaimed guild with nothing else set; each case below
// changes only the facts its step reads.
func unclaimedFacts(rank string) StepFacts {
	return StepFacts{ClaimState: ClaimStateUnclaimed, ViewerRank: rank}
}

func TestSubscribeStep(t *testing.T) {
	until := stepRateLimitedUntil
	cases := []struct {
		name  string
		facts StepFacts
		want  Step
	}{
		{"on_plan", StepFacts{OnLivePlan: true, ClaimState: ClaimStateClaimed, ViewerRank: "member"}, StepOnPlan},
		{"checkout_open", StepFacts{CheckoutOpen: true, ClaimState: ClaimStateClaimed, ViewerRank: "member"}, StepCheckoutOpen},
		{"frozen", StepFacts{ClaimState: ClaimStateContested, ViewerRank: "leader", ViewerVerified: true}, StepFrozen},
		{"ready", StepFacts{ClaimState: ClaimStateClaimed, ViewerHoldsClaim: true, ViewerRank: "leader", ViewerVerified: true}, StepReady},
		{"claimed_ask", StepFacts{ClaimState: ClaimStateClaimed, ViewerRank: "officer", ViewerVerified: true}, StepClaimedAsk},
		{"claim_pending_yours", StepFacts{ClaimState: ClaimStatePending, ViewerPendingClaim: true, ViewerRank: "officer"}, StepClaimPendingYours},
		{"confirm_claim", StepFacts{ClaimState: ClaimStatePending, ViewerRank: "officer"}, StepConfirmClaim},
		{"claim_pending_ask", StepFacts{ClaimState: ClaimStatePending, ViewerRank: "member"}, StepClaimPendingAsk},
		{"claim", StepFacts{ClaimState: ClaimStateUnclaimed, ViewerRank: "leader", HasBattleNet: true}, StepClaim},
		{"claim_link_bnet", unclaimedFacts("leader"), StepClaimLinkBnet},
		{"start_claim", unclaimedFacts("officer"), StepStartClaim},
		{"claim_blocked", StepFacts{ClaimState: ClaimStateUnclaimed, ViewerRank: "officer", RateLimitedUntil: &until}, StepClaimBlocked},
		{"ask_leader", unclaimedFacts("member"), StepAskLeader},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SubscribeStep(tc.facts)
			if got.Step != tc.want {
				t.Fatalf("step = %q, want %q", got.Step, tc.want)
			}
			if tc.want != StepReady && got.TrialUsed != nil {
				t.Errorf("TrialUsed = %v on %q, want nil", *got.TrialUsed, got.Step)
			}
			if tc.want != StepClaimBlocked && (got.AvailableAt != nil || got.Reason != "" || got.OtherGuildName != "") {
				t.Errorf("blocked fields set on %q: %+v", got.Step, got)
			}
		})
	}
}

func TestSubscribeStep_ReadyCarriesTrialUsed(t *testing.T) {
	for _, used := range []bool{false, true} {
		got := SubscribeStep(StepFacts{ClaimState: ClaimStateClaimed, ViewerHoldsClaim: true,
			ViewerRank: "officer", ViewerVerified: true, TrialUsed: used})
		if got.Step != StepReady || got.TrialUsed == nil || *got.TrialUsed != used {
			t.Fatalf("TrialUsed=%v: got %+v", used, got)
		}
	}
}

func TestSubscribeStep_BlockedReasons(t *testing.T) {
	until := stepRateLimitedUntil
	limited := SubscribeStep(StepFacts{ClaimState: ClaimStateUnclaimed, ViewerRank: "leader", HasBattleNet: true,
		RateLimitedUntil: &until})
	if limited.Step != StepClaimBlocked || limited.Reason != BlockedRateLimited {
		t.Fatalf("rate limited: got %+v", limited)
	}
	if limited.AvailableAt == nil || !limited.AvailableAt.Equal(until) {
		t.Errorf("AvailableAt = %v, want %v", limited.AvailableAt, until)
	}
	if limited.OtherGuildName != "" {
		t.Errorf("OtherGuildName = %q on rate_limited, want empty", limited.OtherGuildName)
	}

	other := SubscribeStep(StepFacts{ClaimState: ClaimStateUnclaimed, ViewerRank: "officer",
		HoldsOtherClaim: true, OtherGuildName: "Other Guild", RateLimitedUntil: &until})
	if other.Step != StepClaimBlocked || other.Reason != BlockedHoldsOther || other.OtherGuildName != "Other Guild" {
		t.Fatalf("holds other: got %+v", other)
	}
	if other.AvailableAt != nil {
		t.Errorf("AvailableAt = %v on holds_other, want nil", other.AvailableAt)
	}
}

// The owner's real facts on 2026-10-27: a member, not verified, in two guilds nobody has
// claimed or verified in, with no plan.
func TestSubscribeStep_OwnerFacts(t *testing.T) {
	for _, guild := range []string{"OLYMPUS XXVII", "second guild"} {
		got := SubscribeStep(StepFacts{ClaimState: ClaimStateUnclaimed, ViewerRank: "member"})
		if got.Step != StepAskLeader {
			t.Errorf("%s: step = %q, want ask_leader", guild, got.Step)
		}
	}
}

// Capture 1b: the owner after an addon save that makes them guild master.
func TestSubscribeStep_OwnerAsGuildMaster(t *testing.T) {
	got := SubscribeStep(StepFacts{ClaimState: ClaimStateUnclaimed, ViewerRank: "leader", HasBattleNet: true})
	if got.Step != StepClaim {
		t.Fatalf("step = %q, want claim", got.Step)
	}
}

func TestSubscribeStep_FirstMatchOverlaps(t *testing.T) {
	until := stepRateLimitedUntil
	holder := StepFacts{ClaimState: ClaimStateClaimed, ViewerHoldsClaim: true, ViewerRank: "leader", ViewerVerified: true}
	withCheckout := holder
	withCheckout.CheckoutOpen = true
	cases := []struct {
		name   string
		facts  StepFacts
		want   Step
		reason string
	}{
		{"on a plan and contested", StepFacts{OnLivePlan: true, ClaimState: ClaimStateContested, ViewerRank: "member"}, StepOnPlan, ""},
		{"holder with a checkout open", withCheckout, StepCheckoutOpen, ""},
		{"leader without Battle.net, rate-limited", StepFacts{ClaimState: ClaimStateUnclaimed, ViewerRank: "leader",
			RateLimitedUntil: &until}, StepClaimLinkBnet, ""},
		{"officer holding another claim", StepFacts{ClaimState: ClaimStateUnclaimed, ViewerRank: "officer",
			HoldsOtherClaim: true, OtherGuildName: "Elsewhere"}, StepClaimBlocked, BlockedHoldsOther},
		{"lapsed plan, viewer holds the claim", holder, StepReady, ""},
		{"lapsed plan, viewer not the holder", StepFacts{ClaimState: ClaimStateClaimed, ViewerRank: "officer",
			ViewerVerified: true}, StepClaimedAsk, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SubscribeStep(tc.facts)
			if got.Step != tc.want || got.Reason != tc.reason {
				t.Fatalf("got %q/%q, want %q/%q", got.Step, got.Reason, tc.want, tc.reason)
			}
		})
	}
}

// The holder whose verified rank is member does not bill (Capabilities' CapBilling needs a
// verified officer), so the step is not ready and falls through to claimed_ask.
func TestSubscribeStep_HolderWithoutVerifiedOfficerIsNotReady(t *testing.T) {
	got := SubscribeStep(StepFacts{ClaimState: ClaimStateClaimed, ViewerHoldsClaim: true, ViewerRank: "leader"})
	if got.Step != StepClaimedAsk {
		t.Fatalf("step = %q, want claimed_ask", got.Step)
	}
}
