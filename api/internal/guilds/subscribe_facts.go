// api/internal/guilds/subscribe_facts.go
//
// The facts SubscribeStep reads, gathered for every guild of one viewer, and the subscribe
// view GET /v1/me sends (spec design/specs/2026-10-27-subscribe-and-join.md §8). Each fact
// comes from the read the rule it mirrors already makes: checkGuildCheckout's live plan and
// pending checkout, the trial record, Claim's raw rank and rate limit, Capabilities' verified
// rank. Who may bill or claim is decided in those places, never here.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// StepBillingFacts is the checkout reads the subscribe step needs, for many guilds in one
// query (billing.Store).
type StepBillingFacts interface {
	GuildCheckoutFacts(ctx context.Context, guildIDs []int64) (map[int64]billing.CheckoutFacts, error)
}

// GuildPlans reads many guilds' plan rows in one query (entitlements.Store). GuildStanding
// reads them once and hands them to the subscribe and covered views.
type GuildPlans interface {
	GuildBillings(ctx context.Context, guildIDs []int64) (map[int64]*entitlements.GuildBilling, error)
}

// rankMember is the rank StepFacts names for a viewer with no officer or leader character.
const rankMember = "member"

// heldClaim is a guild whose claim the viewer's account holds.
type heldClaim struct {
	id   int64
	name string
}

// viewerStepFacts is the account-wide half of StepFacts plus the per-guild ranks, read once
// for all of a viewer's guilds.
type viewerStepFacts struct {
	ranks            map[int64]string // highest raw officer/leader rank per guild
	verifiedOfficer  map[int64]bool   // Capabilities' verified officer or leader read
	hasBattleNet     bool
	held             []heldClaim
	rateLimitedUntil *time.Time
}

// subscribeViews answers each guild's subscribe view for viewerID, or nil when the billing
// reads are not wired (the older /v1/me shape). claims and approvers are what GuildStanding
// already read; plans is its one plan-row read. No read here is per guild.
func (s *Store) subscribeViews(ctx context.Context, viewerID int64, guildIDs []int64,
	claims map[int64]auth.GuildClaimView, approvers map[int64]approverFact,
	plans map[int64]*entitlements.GuildBilling) (map[int64]*auth.GuildSubscribeView, error) {
	if s.Billing == nil || s.Plans == nil {
		return nil, nil
	}
	checkouts, err := s.Billing.GuildCheckoutFacts(ctx, guildIDs)
	if err != nil {
		return nil, err
	}
	viewer, err := s.readViewerStepFacts(ctx, viewerID, guildIDs)
	if err != nil {
		return nil, err
	}
	viewer.verifiedOfficer = viewerApproves(approvers)
	out := make(map[int64]*auth.GuildSubscribeView, len(claims))
	for id, claim := range claims {
		out[id] = subscribeView(SubscribeStep(stepFacts(id, claim, viewer, plans[id], checkouts[id])))
	}
	return out, nil
}

// stepFacts is guildID's StepFacts: the viewer's facts plus the guild's plan and checkout facts.
func stepFacts(guildID int64, claim auth.GuildClaimView, v viewerStepFacts, plan *entitlements.GuildBilling,
	checkout billing.CheckoutFacts) StepFacts {
	other := v.otherHeldClaim(guildID)
	return StepFacts{
		OnLivePlan:            entitlements.LiveWholePlan(plan),
		CheckoutOpen:          checkout.Open,
		ClaimState:            claim.State,
		ViewerHoldsClaim:      claim.Yours,
		ViewerPendingClaim:    claim.PendingYours,
		ViewerRank:            v.rank(guildID),
		ViewerVerifiedOfficer: v.verifiedOfficer[guildID],
		HasBattleNet:          v.hasBattleNet,
		HoldsOtherClaim:       other != nil,
		OtherGuildName:        nameOf(other),
		RateLimitedUntil:      v.rateLimitedUntil,
		TrialUsed:             checkout.TrialUsed,
	}
}

func (v viewerStepFacts) rank(guildID int64) string {
	if r, ok := v.ranks[guildID]; ok {
		return r
	}
	return rankMember
}

// otherHeldClaim is a guild other than guildID whose claim the viewer holds, else nil.
func (v viewerStepFacts) otherHeldClaim(guildID int64) *heldClaim {
	for i := range v.held {
		if v.held[i].id != guildID {
			return &v.held[i]
		}
	}
	return nil
}

func nameOf(c *heldClaim) string {
	if c == nil {
		return ""
	}
	return c.name
}

// readViewerStepFacts reads the viewer's ranks, Battle.net link, held claims and claim rate
// limit (four queries). The verified officer guilds come from GuildStanding's approver read.
func (s *Store) readViewerStepFacts(ctx context.Context, viewerID int64, guildIDs []int64) (viewerStepFacts, error) {
	var v viewerStepFacts
	var err error
	if v.ranks, err = s.claimRanks(ctx, viewerID, guildIDs); err != nil {
		return v, err
	}
	if v.hasBattleNet, err = s.hasBattleNet(ctx, viewerID); err != nil {
		return v, err
	}
	if v.held, err = s.heldClaims(ctx, viewerID); err != nil {
		return v, err
	}
	v.rateLimitedUntil, err = s.claimRateLimitedUntil(ctx, viewerID)
	return v, err
}

// claimRanks is eligibleClaimRank for many guilds: the viewer's highest raw officer or leader
// rank per guild, verified or not.
func (s *Store) claimRanks(ctx context.Context, viewerID int64, guildIDs []int64) (map[int64]string, error) {
	rows, err := s.Pool.Query(ctx, `
		select guild_id, rank from guild_characters
		where user_id = $1 and guild_id = any($2) and rank in ('officer', 'leader')`, viewerID, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: step ranks: %w", err)
	}
	defer rows.Close()
	out := map[int64]string{}
	for rows.Next() {
		var id int64
		var rank string
		if err := rows.Scan(&id, &rank); err != nil {
			return nil, fmt.Errorf("guilds: step ranks: %w", err)
		}
		if out[id] != rankLeader {
			out[id] = rank
		}
	}
	return out, rows.Err()
}

// hasBattleNet is the claim handler's Battle.net check (a non-empty bnet_sub).
func (s *Store) hasBattleNet(ctx context.Context, userID int64) (bool, error) {
	var linked bool
	err := s.Pool.QueryRow(ctx, `select coalesce(bnet_sub, '') <> '' from users where id = $1`, userID).Scan(&linked)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("guilds: step battle.net: %w", err)
	}
	return linked, nil
}

// heldClaims is every guild whose claim userID holds (checkClaimRateLimit's first refusal).
func (s *Store) heldClaims(ctx context.Context, userID int64) ([]heldClaim, error) {
	rows, err := s.Pool.Query(ctx, `select id, name from guilds where claimed_by = $1 order by id`, userID)
	if err != nil {
		return nil, fmt.Errorf("guilds: step held claims: %w", err)
	}
	defer rows.Close()
	var out []heldClaim
	for rows.Next() {
		var c heldClaim
		if err := rows.Scan(&c.id, &c.name); err != nil {
			return nil, fmt.Errorf("guilds: step held claims: %w", err)
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// claimRateLimitedUntil is when checkClaimRateLimit's second refusal lifts (the latest claim
// attempt in the window plus the window), nil when no claim attempt is in it.
func (s *Store) claimRateLimitedUntil(ctx context.Context, userID int64) (*time.Time, error) {
	var until *time.Time
	if err := s.Pool.QueryRow(ctx,
		`select max(attempted_at) + `+claimAttemptWindowSQL+` `+recentClaimAttemptsSQL, userID).Scan(&until); err != nil {
		return nil, fmt.Errorf("guilds: step rate limit: %w", err)
	}
	return until, nil
}

// subscribeView is a StepResult as /v1/me sends it: empty values are null.
func subscribeView(r StepResult) *auth.GuildSubscribeView {
	v := &auth.GuildSubscribeView{Step: string(r.Step), TrialUsed: r.TrialUsed}
	if r.AvailableAt != nil {
		at := r.AvailableAt.UTC().Format(time.RFC3339)
		v.AvailableAt = &at
	}
	if r.Reason != "" {
		reason := r.Reason
		v.Reason = &reason
	}
	if r.OtherGuildName != "" {
		name := r.OtherGuildName
		v.OtherGuildName = &name
	}
	return v
}
