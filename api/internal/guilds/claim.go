// api/internal/guilds/claim.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

var (
	ErrAlreadyClaimed            = errors.New("guilds: already claimed")
	ErrClaimPending              = errors.New("guilds: a claim is already pending")
	ErrNotEligible               = errors.New("guilds: no officer or leader character in this guild")
	ErrSameAccount               = errors.New("guilds: cannot confirm your own pending claim")
	ErrNoPendingClaim            = errors.New("guilds: no pending claim")
	ErrNotClaimant               = errors.New("guilds: not the claimant")
	ErrNoBattleNetIdentity       = errors.New("guilds: claiming as guild master requires a linked Battle.net account")
	ErrClaimRateLimited          = errors.New("guilds: only one claim attempt per account every 30 days")
	ErrAlreadyClaimsAnotherGuild = errors.New("guilds: this account already holds another guild's claim")
)

// claimRateLimitWindow and claimRateLimitReason: an account may attempt
// at most one claim (successful or pending) per rolling window, and may
// hold at most one claimed guild at a time - both per the 2026-09-21
// security review response (spec §2.4's amendment).
const claimRateLimitWindow = 30 * 24 * time.Hour

// ClaimResult is what Claim answers with.
type ClaimResult struct {
	Status    string     `json:"status"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// eligibleClaimRank reads the caller's highest raw rank in guildID
// straight from guild_characters, deliberately bypassing the
// verified-only guild_members/GuildRank: claiming is itself the
// corroboration mechanism, so it does not require verified_at
// beforehand.
func (s *Store) eligibleClaimRank(ctx context.Context, guildID, userID int64) (string, bool, error) {
	var rank string
	err := s.Pool.QueryRow(ctx,
		`select rank from guild_characters where guild_id = $1 and user_id = $2 and rank in ('officer', 'leader')
		 order by (rank = 'leader') desc limit 1`, guildID, userID).Scan(&rank)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("guilds: eligible claim rank: %w", err)
	}
	return rank, true, nil
}

// checkClaimRateLimit enforces the 2026-09-21 hardening: at most one
// currently-claimed guild per account, and at most one claim attempt
// (successful or pending) per account per rolling 30 days. Checked
// before either branch of Claim acts. The "already claims another
// guild" check runs first: holding a claim always also means a recent
// attempt exists, so checking attempts first would make the
// already-claims case unreachable in practice - a caller who still
// holds a guild always hears about that specifically, not a generic
// rate limit, even though both are true.
func (s *Store) checkClaimRateLimit(ctx context.Context, userID int64) error {
	var claimedElsewhere int
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from guilds where claimed_by = $1`, userID).Scan(&claimedElsewhere); err != nil {
		return fmt.Errorf("guilds: claim rate limit: %w", err)
	}
	if claimedElsewhere > 0 {
		return ErrAlreadyClaimsAnotherGuild
	}
	var recent int
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from guild_claim_attempts where user_id = $1 and attempted_at >= now() - interval '30 days'`,
		userID).Scan(&recent); err != nil {
		return fmt.Errorf("guilds: claim rate limit: %w", err)
	}
	if recent > 0 {
		return ErrClaimRateLimited
	}
	return nil
}

// Claim starts or completes a claim on guildID for userID. A
// guild-master-rank character claims immediately - the bootstrap for an
// unclaimed guild, since it has no verified member who could vouch for
// a second signal (spec §2.4's amendment records why this stays
// single-signal rather than being closed outright) - hardened by
// requiring a linked Battle.net identity and the account-wide rate
// limit above. An officer's claim goes pending until a second officer
// confirms it, or the guild master's own export does
// (AutoConfirmClaimIfPending).
func (s *Store) Claim(ctx context.Context, guildID, userID int64, hasBattleNetIdentity bool) (ClaimResult, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return ClaimResult{}, err
	}
	if g.ClaimedBy != nil {
		return ClaimResult{}, ErrAlreadyClaimed
	}
	if g.pendingActive(time.Now()) {
		return ClaimResult{}, ErrClaimPending
	}
	rank, ok, err := s.eligibleClaimRank(ctx, guildID, userID)
	if err != nil {
		return ClaimResult{}, err
	}
	if !ok {
		return ClaimResult{}, ErrNotEligible
	}
	if rank == "leader" && !hasBattleNetIdentity {
		return ClaimResult{}, ErrNoBattleNetIdentity
	}
	if err := s.checkClaimRateLimit(ctx, userID); err != nil {
		return ClaimResult{}, err
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return ClaimResult{}, fmt.Errorf("guilds: claim: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if _, err := tx.Exec(ctx, `insert into guild_claim_attempts (user_id) values ($1)`, userID); err != nil {
		return ClaimResult{}, fmt.Errorf("guilds: claim: record attempt: %w", err)
	}

	if rank == "leader" {
		if _, err := tx.Exec(ctx, `update guilds set claimed_by = $2 where id = $1`, guildID, userID); err != nil {
			return ClaimResult{}, fmt.Errorf("guilds: claim: %w", err)
		}
		if err := setVerifiedForAccount(ctx, tx, guildID, userID, "claim"); err != nil {
			return ClaimResult{}, err
		}
		if err := RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
			return ClaimResult{}, err
		}
		if err := tx.Commit(ctx); err != nil {
			return ClaimResult{}, fmt.Errorf("guilds: claim: commit: %w", err)
		}
		return ClaimResult{Status: "confirmed"}, nil
	}

	if _, err := tx.Exec(ctx,
		`update guilds set claim_pending_by = $2, claim_requested_at = now() where id = $1`,
		guildID, userID); err != nil {
		return ClaimResult{}, fmt.Errorf("guilds: claim: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return ClaimResult{}, fmt.Errorf("guilds: claim: commit: %w", err)
	}
	expires := time.Now().Add(ClaimPendingTTL)
	return ClaimResult{Status: "pending", ExpiresAt: &expires}, nil
}

// ConfirmClaim completes a pending officer claim: a second, distinct
// officer/leader account vouches for the pending claimant. Returns the
// claimant's user id so the handler can look up their battletag.
func (s *Store) ConfirmClaim(ctx context.Context, guildID, confirmerID int64) (int64, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return 0, err
	}
	if !g.pendingActive(time.Now()) {
		return 0, ErrNoPendingClaim
	}
	claimant := *g.ClaimPendingBy
	if claimant == confirmerID {
		return 0, ErrSameAccount
	}
	if _, ok, err := s.eligibleClaimRank(ctx, guildID, confirmerID); err != nil {
		return 0, err
	} else if !ok {
		return 0, ErrNotEligible
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("guilds: confirm claim: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx,
		`update guilds set claimed_by = $2, claim_pending_by = null, claim_requested_at = null where id = $1`,
		guildID, claimant); err != nil {
		return 0, fmt.Errorf("guilds: confirm claim: %w", err)
	}
	if err := setVerifiedForAccount(ctx, tx, guildID, claimant, "claim"); err != nil {
		return 0, err
	}
	if err := RecomputeMembership(ctx, tx, guildID, &claimant); err != nil {
		return 0, err
	}
	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("guilds: confirm claim: commit: %w", err)
	}
	return claimant, nil
}

// AutoConfirmClaimIfPending confirms a pending officer claim the moment
// any account's export shows guild-master rank for guildID, per the
// design's "failing that, by the guild master's export" - the guild
// master need never call the confirm endpoint themselves. Requires
// triggeringUserID (the account whose export produced the rank-0
// signal) to differ from claim_pending_by, mirroring ConfirmClaim's own
// ErrSameAccount check - a single account must not be able to
// self-confirm its own pending claim with a second forged character
// (2026-09-21 security review finding). A pending claim older than
// ClaimPendingTTL has already expired and is left untouched; a fresh
// Claim call is what clears it. Called from addon.Store.syncGuild in
// the same transaction as the character write that made this account's
// rank "leader".
func AutoConfirmClaimIfPending(ctx context.Context, tx pgx.Tx, guildID, triggeringUserID int64) error {
	var claimant int64
	err := tx.QueryRow(ctx,
		`update guilds set claimed_by = claim_pending_by, claim_pending_by = null, claim_requested_at = null
		 where id = $1 and claim_pending_by is not null and claim_pending_by != $2
		   and claim_requested_at >= now() - interval '14 days'
		 returning claimed_by`, guildID, triggeringUserID).Scan(&claimant)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("guilds: auto-confirm claim for guild %d: %w", guildID, err)
	}
	if err := setVerifiedForAccount(ctx, tx, guildID, claimant, "claim"); err != nil {
		return err
	}
	return RecomputeMembership(ctx, tx, guildID, &claimant)
}

// ReleaseClaim clears guildID's claim. Only the current claimant or a
// moderator may do it.
func (s *Store) ReleaseClaim(ctx context.Context, guildID, userID int64, moderator bool) error {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	if !moderator && (g.ClaimedBy == nil || *g.ClaimedBy != userID) {
		return ErrNotClaimant
	}
	if _, err := s.Pool.Exec(ctx, `update guilds set claimed_by = null where id = $1`, guildID); err != nil {
		return fmt.Errorf("guilds: release claim: %w", err)
	}
	return nil
}
