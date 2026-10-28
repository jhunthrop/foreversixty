// api/internal/guilds/covered_facts.go
//
// The covered view GET /v1/me puts on every guilds[] entry (spec
// design/specs/2026-10-27-subscribe-and-join.md §8): what plan covers the guild, in what
// state, who bills it, and how many of its seats are given. The plan row is the one
// entitlements.Store.GuildBilling reads for the billing page; the seat count follows the
// seats package's rule (slots within the tier's count, held or accepted).
package guilds

import (
	"context"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/seats"
)

// coveredLapsedShown is how long after its end a lapsed plan is still shown.
const coveredLapsedShown = 90 * 24 * time.Hour

// covered states, as GuildCoveredView.State words them.
const (
	coveredTrial   = "trial"
	coveredActive  = "active"
	coveredEnding  = "ending"
	coveredPastDue = "past_due"
	coveredLapsed  = "lapsed"
)

// Stripe statuses the covered state reads.
const (
	statusTrialing = "trialing"
	statusPastDue  = "past_due"
)

// coveredState is the state of a whole plan at now, and the moment it ends; ok is false when
// the plan is not shown (ended more than coveredLapsedShown ago, or ended with no end date).
func coveredState(gb *entitlements.GuildBilling, now time.Time) (state string, endsAt *time.Time, ok bool) {
	endsAt = gb.CurrentPeriodEnd
	if gb.Status == statusTrialing && gb.TrialEnd != nil {
		endsAt = gb.TrialEnd
	}
	if !gb.Active {
		if gb.CurrentPeriodEnd == nil || now.Sub(*gb.CurrentPeriodEnd) > coveredLapsedShown {
			return "", nil, false
		}
		return coveredLapsed, gb.CurrentPeriodEnd, true
	}
	switch {
	case gb.Status == statusPastDue:
		return coveredPastDue, endsAt, true
	case gb.CancelAtPeriodEnd:
		return coveredEnding, endsAt, true
	case gb.Status == statusTrialing:
		return coveredTrial, endsAt, true
	}
	return coveredActive, endsAt, true
}

// coveredViews answers each guild's covered view for viewerID from plans, GuildStanding's one
// plan-row read; a guild with no shown whole plan has none. Nil when the plan read is not
// wired (the older /v1/me shape). Its reads are fixed whatever the guild count: verified
// guilds, given seats, and the payers' names (skipped when none is to be shown).
func (s *Store) coveredViews(ctx context.Context, viewerID int64, guildIDs []int64,
	plans map[int64]*entitlements.GuildBilling) (map[int64]*auth.GuildCoveredView, error) {
	if s.Plans == nil {
		return nil, nil
	}
	verified, err := s.verifiedGuilds(ctx, viewerID, guildIDs)
	if err != nil {
		return nil, err
	}
	held, err := s.heldSeatSlots(ctx, guildIDs)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	out := make(map[int64]*auth.GuildCoveredView, len(guildIDs))
	payers := map[int64]int64{} // guild -> the paying account this viewer may see named
	for _, id := range guildIDs {
		gb := plans[id]
		view := coveredView(gb, viewerID, held[id], now)
		if view == nil {
			continue
		}
		out[id] = view
		if verified[id] && gb.BillingUserID != nil {
			payers[id] = *gb.BillingUserID
		}
	}
	return out, s.namePayers(ctx, out, payers)
}

// namePayers fills billed_by with each payer's character in the guild, named by the rule
// claim.holder_name uses (guildCharacterNames), so a card names one person one way. A payer
// with no named character there stays "".
func (s *Store) namePayers(ctx context.Context, views map[int64]*auth.GuildCoveredView, payers map[int64]int64) error {
	names, err := s.guildCharacterNames(ctx, payers)
	if err != nil {
		return err
	}
	for id, name := range names {
		views[id].BilledBy = name
	}
	return nil
}

// coveredView is one guild's covered view without billed_by, nil when it has no shown whole
// plan.
func coveredView(gb *entitlements.GuildBilling, viewerID int64, held []seatHolder, now time.Time) *auth.GuildCoveredView {
	if gb == nil || gb.Scoped() {
		return nil
	}
	state, endsAt, ok := coveredState(gb, now)
	if !ok {
		return nil
	}
	view := &auth.GuildCoveredView{
		Tier: gb.Tier, State: state, EndsAt: formatTime(endsAt),
		YouPay:     gb.BillingUserID != nil && *gb.BillingUserID == viewerID,
		SeatsTotal: catalog.SeatsForTier(gb.Tier),
	}
	view.SeatsGiven, view.YouHoldSeat = seatFacts(held, view.SeatsTotal, viewerID)
	return view
}

func formatTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.UTC().Format(time.RFC3339)
	return &s
}

// seatHolder is one seat row that is given: its slot and the account holding it.
type seatHolder struct {
	slot   int
	userID int64
}

// seatFacts counts the given seats within the tier's total and whether viewerID holds one.
func seatFacts(rows []seatHolder, total int, viewerID int64) (given int, youHold bool) {
	for _, r := range rows {
		if r.slot < 1 || r.slot > total {
			continue
		}
		given++
		youHold = youHold || r.userID == viewerID
	}
	return given, youHold
}

// heldSeatSlots reads every given (held or accepted) seat of guildIDs in one query.
func (s *Store) heldSeatSlots(ctx context.Context, guildIDs []int64) (map[int64][]seatHolder, error) {
	rows, err := s.Pool.Query(ctx, `
		select guild_id, slot, holder_user_id from premium_seats
		where guild_id = any($1) and state in ($2, $3) and holder_user_id is not null`,
		guildIDs, string(seats.StateHeld), string(seats.StateAccepted))
	if err != nil {
		return nil, fmt.Errorf("guilds: covered seats: %w", err)
	}
	defer rows.Close()
	out := map[int64][]seatHolder{}
	for rows.Next() {
		var guildID int64
		var h seatHolder
		if err := rows.Scan(&guildID, &h.slot, &h.userID); err != nil {
			return nil, fmt.Errorf("guilds: covered seats: %w", err)
		}
		out[guildID] = append(out[guildID], h)
	}
	return out, rows.Err()
}

// verifiedGuilds is the set of guildIDs where viewerID has a verified character.
func (s *Store) verifiedGuilds(ctx context.Context, viewerID int64, guildIDs []int64) (map[int64]bool, error) {
	rows, err := s.Pool.Query(ctx, `
		select distinct guild_id from guild_characters
		where user_id = $1 and guild_id = any($2) and verified_at is not null`, viewerID, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: covered verified: %w", err)
	}
	defer rows.Close()
	out := map[int64]bool{}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("guilds: covered verified: %w", err)
		}
		out[id] = true
	}
	return out, rows.Err()
}
