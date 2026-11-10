// api/internal/guilds/rank_override.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// Rank override ranks: the only ranks guild_rank_overrides accepts.
const (
	RankOfficer = "officer"
	RankLeader  = "leader"
)

// RankOverride is an operator's time-boxed rank for one account in one
// guild (api grant-rank). RecomputeMembership applies it while unexpired.
type RankOverride struct {
	GuildID   int64
	UserID    int64
	Rank      string
	ExpiresAt time.Time
	Note      string
	CreatedBy string
}

// RankOverrideResult is the account's guild_members state after a grant or
// revoke: Rank is empty when the account has no character in the guild
// (an override never creates membership on its own).
type RankOverrideResult struct {
	Rank     string
	Verified bool
}

// ValidRankOverride reports whether rank is one an override may grant.
func ValidRankOverride(rank string) bool {
	return rank == RankOfficer || rank == RankLeader
}

// GrantRankOverride writes (or replaces) the override and recomputes the
// account's membership in the same transaction.
func (s *Store) GrantRankOverride(ctx context.Context, o RankOverride) (RankOverrideResult, error) {
	if !ValidRankOverride(o.Rank) {
		return RankOverrideResult{}, fmt.Errorf("guilds: rank override: rank %q must be %s or %s", o.Rank, RankOfficer, RankLeader)
	}
	return s.withMembershipRecompute(ctx, o.GuildID, o.UserID, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, `
			insert into guild_rank_overrides (guild_id, user_id, rank, expires_at, note, created_by)
			values ($1, $2, $3, $4, $5, $6)
			on conflict (guild_id, user_id) do update set
			  rank = excluded.rank, expires_at = excluded.expires_at, note = excluded.note,
			  created_by = excluded.created_by, created_at = now()`,
			o.GuildID, o.UserID, o.Rank, o.ExpiresAt, o.Note, o.CreatedBy)
		return err
	})
}

// RevokeRankOverride deletes the override and recomputes the account's
// membership, releasing a claim it no longer qualifies for. It reports
// false when there was no override to revoke.
func (s *Store) RevokeRankOverride(ctx context.Context, guildID, userID int64) (RankOverrideResult, bool, error) {
	var removed bool
	res, err := s.withMembershipRecompute(ctx, guildID, userID, func(tx pgx.Tx) error {
		tag, err := tx.Exec(ctx,
			`delete from guild_rank_overrides where guild_id = $1 and user_id = $2`, guildID, userID)
		removed = tag.RowsAffected() == 1
		return err
	})
	return res, removed, err
}

// withMembershipRecompute runs write, then RecomputeMembership and
// ReleaseClaimIfLost for the account, in one transaction, and reads back
// the account's resulting guild_members row.
func (s *Store) withMembershipRecompute(ctx context.Context, guildID, userID int64, write func(pgx.Tx) error) (RankOverrideResult, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return RankOverrideResult{}, fmt.Errorf("guilds: rank override: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := write(tx); err != nil {
		return RankOverrideResult{}, fmt.Errorf("guilds: rank override for guild %d user %d: %w", guildID, userID, err)
	}
	if err := RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
		return RankOverrideResult{}, err
	}
	if err := ReleaseClaimIfLost(ctx, tx, guildID, userID); err != nil {
		return RankOverrideResult{}, err
	}
	res, err := readMembership(ctx, tx, guildID, userID)
	if err != nil {
		return RankOverrideResult{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return RankOverrideResult{}, fmt.Errorf("guilds: rank override: commit: %w", err)
	}
	return res, nil
}

func readMembership(ctx context.Context, tx pgx.Tx, guildID, userID int64) (RankOverrideResult, error) {
	var res RankOverrideResult
	err := tx.QueryRow(ctx,
		`select rank, verified_at is not null from guild_members where guild_id = $1 and user_id = $2`,
		guildID, userID).Scan(&res.Rank, &res.Verified)
	if errors.Is(err, pgx.ErrNoRows) {
		return RankOverrideResult{}, nil
	}
	if err != nil {
		return RankOverrideResult{}, fmt.Errorf("guilds: rank override: read membership: %w", err)
	}
	return res, nil
}

// ExpireRankOverrides deletes every override expired as of at and
// recomputes each affected account, so an expired override stops
// granting rank on the next sweep tick even if the account never syncs.
// It returns the (guild, user) pairs it expired.
func (s *Store) ExpireRankOverrides(ctx context.Context, at time.Time) ([][2]int64, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("guilds: expire rank overrides: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	rows, err := tx.Query(ctx,
		`delete from guild_rank_overrides where expires_at <= $1 returning guild_id, user_id`, at)
	if err != nil {
		return nil, fmt.Errorf("guilds: expire rank overrides: delete: %w", err)
	}
	affected, err := scanGuildUserPairs(rows)
	if err != nil {
		return nil, fmt.Errorf("guilds: expire rank overrides: %w", err)
	}
	expired := make([][2]int64, 0, len(affected))
	for p := range affected {
		uid := p.userID
		if err := RecomputeMembership(ctx, tx, p.guildID, &uid); err != nil {
			return nil, err
		}
		if err := ReleaseClaimIfLost(ctx, tx, p.guildID, uid); err != nil {
			return nil, err
		}
		expired = append(expired, [2]int64{p.guildID, uid})
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("guilds: expire rank overrides: commit: %w", err)
	}
	return expired, nil
}
