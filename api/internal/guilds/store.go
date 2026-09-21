// api/internal/guilds/store.go

// Package guilds is every guild mutation route: claim, invite, settings,
// roster approve/remove, consent, and the signed-in home. auth.Store keeps
// owning guild_members reads for report access (GuildRank); this package
// owns guild_characters, the row an export actually writes, and derives
// guild_members from it via RecomputeMembership.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNotFound is returned when a guild id names no row.
var ErrNotFound = errors.New("guilds: not found")

// ClaimPendingTTL is how long a pending officer claim stays open before a
// fresh Claim call may start a new one (kept from the design's RULING 5).
const ClaimPendingTTL = 14 * 24 * time.Hour

// Guild is one row of the guilds table, as this package's handlers need it.
type Guild struct {
	ID                    int64
	Region, Ruleset, Name string
	DefaultVisibility     string
	ClaimedBy             *int64
	ClaimPendingBy        *int64
	ClaimRequestedAt      *time.Time
	OfficerMaxRankIndex   int
	InviteTokenRotatedAt  *time.Time
}

// pendingActive reports whether g carries a claim pending within the
// ClaimPendingTTL window as of now.
func (g Guild) pendingActive(now time.Time) bool {
	return g.ClaimPendingBy != nil && g.ClaimRequestedAt != nil &&
		now.Sub(*g.ClaimRequestedAt) <= ClaimPendingTTL
}

// Store is every guild-mutation read and write.
type Store struct{ Pool *pgxpool.Pool }

func (s *Store) getGuild(ctx context.Context, id int64) (Guild, error) {
	var g Guild
	err := s.Pool.QueryRow(ctx,
		`select id, region, ruleset, name, default_visibility, claimed_by, claim_pending_by,
		        claim_requested_at, officer_max_rank_index, invite_token_rotated_at
		 from guilds where id = $1`, id).
		Scan(&g.ID, &g.Region, &g.Ruleset, &g.Name, &g.DefaultVisibility, &g.ClaimedBy,
			&g.ClaimPendingBy, &g.ClaimRequestedAt, &g.OfficerMaxRankIndex, &g.InviteTokenRotatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Guild{}, ErrNotFound
	}
	if err != nil {
		return Guild{}, fmt.Errorf("guilds: read %d: %w", id, err)
	}
	return g, nil
}

// IsMember reports whether this account has any guild_characters row in
// this guild, verified or not — the gate for GET /v1/guilds/{id}/home,
// never for report visibility or edit rights (those stay on
// auth.Store.GuildRank, which is verified-only).
func (s *Store) IsMember(ctx context.Context, guildID, userID int64) (bool, error) {
	var exists bool
	if err := s.Pool.QueryRow(ctx,
		`select exists(select 1 from guild_characters where guild_id = $1 and user_id = $2)`,
		guildID, userID).Scan(&exists); err != nil {
		return false, fmt.Errorf("guilds: is member: %w", err)
	}
	return exists, nil
}

// RecomputeMembership derives the account-level guild_members row for
// (guildID, userID) from that account's current guild_characters rows,
// within tx. userID nil recomputes every account in the guild at once
// (used by the officer-threshold settings PATCH and the ageing/
// corroboration sweeps). It is also called by addon.Store.PutExports in
// the same transaction as the character write it follows — this is why
// it takes a pgx.Tx rather than opening its own.
func RecomputeMembership(ctx context.Context, tx pgx.Tx, guildID int64, userID *int64) error {
	if _, err := tx.Exec(ctx, `
		insert into guild_members (guild_id, user_id, rank, verified_at, refreshed_at)
		select $1, gc.user_id,
		       case max(case when gc.verified_at is not null then
		                   case gc.rank when 'leader' then 3 when 'officer' then 2 else 1 end
		                 end)
		         when 3 then 'leader' when 2 then 'officer' else 'member' end,
		       min(gc.verified_at) filter (where gc.verified_at is not null),
		       now()
		from guild_characters gc
		where gc.guild_id = $1 and ($2::bigint is null or gc.user_id = $2)
		group by gc.user_id
		on conflict (guild_id, user_id) do update set
		  rank = excluded.rank, verified_at = excluded.verified_at, refreshed_at = now()
	`, guildID, userID); err != nil {
		return fmt.Errorf("guilds: recompute membership for guild %d: %w", guildID, err)
	}
	if _, err := tx.Exec(ctx, `
		delete from guild_members m
		where m.guild_id = $1 and ($2::bigint is null or m.user_id = $2)
		  and not exists (
		    select 1 from guild_characters gc where gc.guild_id = m.guild_id and gc.user_id = m.user_id
		  )
	`, guildID, userID); err != nil {
		return fmt.Errorf("guilds: prune membership for guild %d: %w", guildID, err)
	}
	return nil
}

// ReleaseClaimIfLost clears guilds.claimed_by when the account that holds
// it no longer has a verified guild_members row in that guild — leaving,
// or losing verification in, a guild you lead releases your claim. Call
// it after RecomputeMembership, in the same transaction.
func ReleaseClaimIfLost(ctx context.Context, tx pgx.Tx, guildID, userID int64) error {
	if _, err := tx.Exec(ctx,
		`update guilds set claimed_by = null
		 where id = $1 and claimed_by = $2
		   and not exists (
		     select 1 from guild_members
		     where guild_id = $1 and user_id = $2 and verified_at is not null
		   )`, guildID, userID); err != nil {
		return fmt.Errorf("guilds: release claim for guild %d: %w", guildID, err)
	}
	return nil
}
