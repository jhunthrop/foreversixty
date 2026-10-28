// api/internal/guilds/claim_facts.go
//
// The claim and verify facts GET /v1/me adds to each guilds[] entry (spec
// design/specs/2026-10-27-subscribe-and-join.md §8), read for many guilds at once from the
// same columns Claim, ConfirmClaim and the contest code read.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// claimHolderState is a guild's claim view and the account a holder name would come from.
type claimHolderState struct {
	view   auth.GuildClaimView
	holder *int64 // the claim holder, else the active pending claimant
}

// ClaimAndVerifyViews answers, for each of guildIDs that exists, the claim view the viewer
// sees and whether anyone in the guild can approve a character. It is the one call
// GuildStanding makes for these facts.
func (s *Store) ClaimAndVerifyViews(ctx context.Context, viewerID int64, guildIDs []int64) (
	map[int64]auth.GuildClaimView, map[int64]auth.GuildVerifyView, error) {
	claims, err := s.claimViews(ctx, viewerID, guildIDs)
	if err != nil {
		return nil, nil, err
	}
	approvers, err := s.hasApprover(ctx, guildIDs)
	if err != nil {
		return nil, nil, err
	}
	verify := make(map[int64]auth.GuildVerifyView, len(guildIDs))
	for _, id := range guildIDs {
		verify[id] = auth.GuildVerifyView{HasApprover: approvers[id]}
	}
	return claims, verify, nil
}

// claimViews is each guild's claim as viewerID sees it. holder_name is the holder's (or the
// pending claimant's) character in the guild, and only a verified officer or leader of that
// guild is shown it.
func (s *Store) claimViews(ctx context.Context, viewerID int64, guildIDs []int64) (map[int64]auth.GuildClaimView, error) {
	states, err := s.claimStates(ctx, viewerID, guildIDs)
	if err != nil {
		return nil, err
	}
	approves, err := s.viewerApproves(ctx, viewerID, guildIDs)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]auth.GuildClaimView, len(states))
	for id, st := range states {
		if approves[id] && st.holder != nil {
			name, err := s.guildCharacterName(ctx, id, *st.holder)
			if err != nil {
				return nil, err
			}
			st.view.HolderName = name
		}
		out[id] = st.view
	}
	return out, nil
}

// claimStates derives each guild's claim phase with the rule the claim endpoints use.
func (s *Store) claimStates(ctx context.Context, viewerID int64, guildIDs []int64) (map[int64]claimHolderState, error) {
	rows, err := s.Pool.Query(ctx, `
		select id, claimed_by, claimed_at, claim_pending_by, claim_requested_at, claim_contested_at
		from guilds where id = any($1)`, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: claim facts: %w", err)
	}
	defer rows.Close()
	now := time.Now()
	out := make(map[int64]claimHolderState, len(guildIDs))
	for rows.Next() {
		var g Guild
		if err := rows.Scan(&g.ID, &g.ClaimedBy, &g.ClaimedAt, &g.ClaimPendingBy, &g.ClaimRequestedAt,
			&g.ClaimContestedAt); err != nil {
			return nil, fmt.Errorf("guilds: claim facts: %w", err)
		}
		out[g.ID] = claimHolderStateOf(g, viewerID, now)
	}
	return out, rows.Err()
}

// claimHolderStateOf is g's claim phase and who holds it, from viewerID's side.
func claimHolderStateOf(g Guild, viewerID int64, now time.Time) claimHolderState {
	view := auth.GuildClaimView{State: claimState(g, now).State}
	var holder *int64
	switch {
	case g.ClaimedBy != nil:
		holder = g.ClaimedBy
		view.Yours = *g.ClaimedBy == viewerID
	case g.pendingActive(now):
		holder = g.ClaimPendingBy
		view.PendingYours = *g.ClaimPendingBy == viewerID
	}
	return claimHolderState{view: view, holder: holder}
}

// viewerApproves is the guilds among guildIDs in which viewerID is a verified officer or
// leader.
func (s *Store) viewerApproves(ctx context.Context, viewerID int64, guildIDs []int64) (map[int64]bool, error) {
	rows, err := s.Pool.Query(ctx, `
		select guild_id from guild_members
		where user_id = $1 and guild_id = any($2) and verified_at is not null and rank in ('officer', 'leader')`,
		viewerID, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: viewer approves: %w", err)
	}
	return collectGuildIDs(rows)
}

// hasApprover is, per guild, whether at least one verified officer or leader exists: the
// predicate Capabilities and verifiedOfficerOrLeader read.
func (s *Store) hasApprover(ctx context.Context, guildIDs []int64) (map[int64]bool, error) {
	rows, err := s.Pool.Query(ctx, `
		select distinct guild_id from guild_members
		where guild_id = any($1) and verified_at is not null and rank in ('officer', 'leader')`, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: has approver: %w", err)
	}
	return collectGuildIDs(rows)
}

// collectGuildIDs reads a one-column guild id result into a set.
func collectGuildIDs(rows pgx.Rows) (map[int64]bool, error) {
	defer rows.Close()
	out := map[int64]bool{}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("guilds: guild ids: %w", err)
		}
		out[id] = true
	}
	return out, rows.Err()
}

// guildCharacterName is the name of userID's highest-ranked character in guildID (nil when
// the account has none with a known name).
func (s *Store) guildCharacterName(ctx context.Context, guildID, userID int64) (*string, error) {
	var name string
	err := s.Pool.QueryRow(ctx, `
		select coalesce(ae.name, c.name)
		from guild_characters gc
		left join addon_exports ae on ae.character_key = gc.character_key
		left join characters c on c.key = gc.character_key
		where gc.guild_id = $1 and gc.user_id = $2 and gc.character_key not like 'account:%'
		  and coalesce(ae.name, c.name) is not null
		order by (gc.rank = 'leader') desc, (gc.rank = 'officer') desc, gc.verified_at is not null desc, 1
		limit 1`, guildID, userID).Scan(&name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: claim holder name: %w", err)
	}
	return &name, nil
}
