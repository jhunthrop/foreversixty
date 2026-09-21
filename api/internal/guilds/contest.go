// api/internal/guilds/contest.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

var (
	ErrNoActiveClaim                 = errors.New("guilds: no claim to contest")
	ErrAlreadyContested              = errors.New("guilds: claim already contested")
	ErrInvalidOutcome                = errors.New("guilds: invalid resolution outcome")
	ErrContestRateLimited            = errors.New("guilds: only one contest attempt per account every 30 days")
	ErrAlreadyContestingAnotherGuild = errors.New("guilds: this account already has an open contest on another guild")
	ErrContestAlreadyUpheld          = errors.New("guilds: this account's contest of this guild's claim has already been upheld")
	ErrGuildRecentlyUpheld           = errors.New("guilds: this guild's claim was upheld within the last 30 days")
)

// claimUpheldCooldown is how long after an uphold a guild's claim
// cannot be re-contested by anyone (fifth security review response,
// MEDIUM): without it, a fresh Battle.net account could re-contest the
// instant a moderator upholds, and roughly one fresh account every
// contest-rate-limit window could keep a legitimately claimed guild
// frozen indefinitely - about 30 accounts to hold it frozen for a
// month.
const claimUpheldCooldown = 30 * 24 * time.Hour

// ClaimStateView is the claim state object GET /v1/guilds/{id}/settings
// and GET /v1/guilds/{id}/home both expose (2026-09-21 security review
// response, spec §2.4's amendment). Frozen (added by the second review
// response, simplified by the fourth) is exactly State == "contested":
// a contest always freezes the disputed claimant's officer tools for
// that guild until a moderator resolves it - two earlier attempts at a
// narrower rule (young-claim-only, then an independence-checked
// corroboration test) were each found gameable by a squatter, so the
// fourth response deletes the escape rather than patch it a third
// time. The field stays in the response shape so the web needs no
// change; see spec §2.4's fourth amendment for the reasoning and the
// accepted residual risk.
type ClaimStateView struct {
	State  string     `json:"state"`
	Since  *time.Time `json:"since,omitempty"`
	Frozen bool       `json:"frozen"`
}

// claimState derives the four-phase claim state a guild is in as of
// now, entirely pure (no DB access - the fourth security review
// response removed the only rule that ever needed one). Since is filled
// in for every state that has a meaningful "as of" instant: claimed_at
// for "claimed", claim_requested_at for "pending", claim_contested_at
// for "contested".
func claimState(g Guild, now time.Time) ClaimStateView {
	switch {
	case g.ClaimContestedAt != nil:
		return ClaimStateView{State: "contested", Since: g.ClaimContestedAt, Frozen: true}
	case g.ClaimedBy != nil:
		return ClaimStateView{State: "claimed", Since: g.ClaimedAt}
	case g.pendingActive(now):
		return ClaimStateView{State: "pending", Since: g.ClaimRequestedAt}
	default:
		return ClaimStateView{State: "unclaimed"}
	}
}

// checkContestRateLimit enforces A2: at most one open contest per
// account across every guild, and at most one contest attempt per
// account per rolling 30 days - the same shape checkClaimRateLimit
// already applies to claims, now sharing guild_claim_attempts via its
// kind column.
// checkContestRateLimit takes tx, not s.Pool, for the same reason
// checkClaimRateLimit does: it must run after
// lockAccountForClaimActivity, inside the same transaction that later
// inserts the attempt row and updates guilds.claim_contested_by, or a
// burst of concurrent contests from the same account could each read
// the same stale counts (HIGH, third security review response).
func (s *Store) checkContestRateLimit(ctx context.Context, tx pgx.Tx, userID int64) error {
	var openElsewhere int
	if err := tx.QueryRow(ctx,
		`select count(*) from guilds where claim_contested_by = $1`, userID).Scan(&openElsewhere); err != nil {
		return fmt.Errorf("guilds: contest rate limit: %w", err)
	}
	if openElsewhere > 0 {
		return ErrAlreadyContestingAnotherGuild
	}
	var recent int
	if err := tx.QueryRow(ctx,
		`select count(*) from guild_claim_attempts
		 where user_id = $1 and kind = 'contest' and attempted_at >= now() - interval '30 days'`,
		userID).Scan(&recent); err != nil {
		return fmt.Errorf("guilds: contest rate limit: %w", err)
	}
	if recent > 0 {
		return ErrContestRateLimited
	}
	return nil
}

// previouslyUpheld reports whether contesterID's contest of guildID's
// claim has already been decided in the claimant's favor once before -
// A3's "an uphold is final for that pair": the same account may not
// keep re-contesting a claim a moderator has already looked at and
// upheld.
func (s *Store) previouslyUpheld(ctx context.Context, guildID, contesterID int64) (bool, error) {
	var yes bool
	if err := s.Pool.QueryRow(ctx,
		`select exists(
		   select 1 from guild_claim_resolutions
		   where guild_id = $1 and contester_id = $2 and outcome = 'uphold'
		 )`, guildID, contesterID).Scan(&yes); err != nil {
		return false, fmt.Errorf("guilds: previously upheld: %w", err)
	}
	return yes, nil
}

// recentlyUpheld reports whether guildID's most recent uphold
// resolution (from any contester) fell within claimUpheldCooldown,
// checked from guild_claim_resolutions by guild id alone - who
// contests next does not matter, only that this guild was just
// settled. A moderator's claim_reopened_at (POST .../claim/reopen)
// clears the cooldown for exactly the uphold(s) resolved at or before
// it; a later, fresh uphold is not retroactively cleared by a stale
// reopen (its resolved_at is after g.ClaimReopenedAt).
func (s *Store) recentlyUpheld(ctx context.Context, g Guild) (bool, error) {
	var resolvedAt time.Time
	err := s.Pool.QueryRow(ctx, `
		select resolved_at from guild_claim_resolutions
		where guild_id = $1 and outcome = 'uphold'
		order by resolved_at desc limit 1
	`, g.ID).Scan(&resolvedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("guilds: recently upheld: %w", err)
	}
	if time.Since(resolvedAt) >= claimUpheldCooldown {
		return false, nil
	}
	if g.ClaimReopenedAt != nil && !g.ClaimReopenedAt.Before(resolvedAt) {
		return false, nil
	}
	return true, nil
}

// ReopenClaimCooldown clears guildID's per-guild contest cooldown
// (recentlyUpheld) once - moderator only (fifth security review
// response). Idempotent: harmless, and simply a no-op resolution-wise,
// when no cooldown is currently active.
func (s *Store) ReopenClaimCooldown(ctx context.Context, guildID int64) error {
	tag, err := s.Pool.Exec(ctx, `update guilds set claim_reopened_at = now() where id = $1`, guildID)
	if err != nil {
		return fmt.Errorf("guilds: reopen claim cooldown for guild %d: %w", guildID, err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// ContestClaim flags guildID's current claim (pending or claimed) as
// disputed by contesterID, always freezing the claimant's officer
// powers for this guild until a moderator resolves it (A4, simplified
// by the fourth security review response). The contester must: hold a
// linked Battle.net
// identity (A1, exactly as a leader claim does - a free email-only
// account with one forged officer export must never be able to freeze
// a real guild's officer tools); hold a raw officer/leader rank
// character in this guild (eligibleClaimRank - the same raw,
// unverified-tolerant check the claim flow itself uses, since
// disputing a possibly-forged claim cannot itself require verification
// the dispute exists to question) on a different account from the
// current claimant/pending claimant; not have already had a contest of
// this same guild upheld (A3); and stay within the same per-account
// rate limit claiming uses (A2).
func (s *Store) ContestClaim(ctx context.Context, guildID, contesterID int64, hasBattleNetIdentity bool) error {
	if !hasBattleNetIdentity {
		return ErrNoBattleNetIdentity
	}
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	claimant, ok := g.activeClaimant(time.Now())
	if !ok {
		return ErrNoActiveClaim
	}
	if g.ClaimContestedAt != nil {
		return ErrAlreadyContested
	}
	if upheld, err := s.recentlyUpheld(ctx, g); err != nil {
		return err
	} else if upheld {
		return ErrGuildRecentlyUpheld
	}
	if claimant == contesterID {
		return ErrSameAccount
	}
	if _, ok, err := s.eligibleClaimRank(ctx, guildID, contesterID); err != nil {
		return err
	} else if !ok {
		return ErrNotEligible
	}
	if upheld, err := s.previouslyUpheld(ctx, guildID, contesterID); err != nil {
		return err
	} else if upheld {
		return ErrContestAlreadyUpheld
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: contest claim: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if err := lockAccountForClaimActivity(ctx, tx, contesterID); err != nil {
		return err
	}
	if err := s.checkContestRateLimit(ctx, tx, contesterID); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx,
		`insert into guild_claim_attempts (user_id, kind) values ($1, 'contest')`, contesterID); err != nil {
		return fmt.Errorf("guilds: contest claim: record attempt: %w", err)
	}
	// Guarded by "claim_contested_at is null" so two concurrent
	// contesters of the SAME guild cannot both silently overwrite each
	// other's claim_contested_by; guilds_claim_contested_by_idx (a
	// second, independent guard) then catches the cross-guild race -
	// the SAME account trying to hold an open contest on two different
	// guilds at once - as a commit-time unique violation, since the
	// pre-transaction checkContestRateLimit read above can still be
	// stale against a sibling transaction racing on a different guild
	// row entirely, which the per-account advisory lock above does not
	// by itself prevent (HIGH, third security review response).
	tag, err := tx.Exec(ctx,
		`update guilds set claim_contested_at = now(), claim_contested_by = $2
		 where id = $1 and claim_contested_at is null`,
		guildID, contesterID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrAlreadyContestingAnotherGuild
		}
		return fmt.Errorf("guilds: contest claim: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrAlreadyContested
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("guilds: contest claim: commit: %w", err)
	}
	return nil
}

// unverifyClaimSourcedRows clears verified_at/verified_by on exactly
// userID's guild_characters rows in guildID whose verification came
// from the claim being resolved - scoped to that one account, not
// every claim-sourced row in the guild, so an unrelated account's
// long-settled verification (e.g. from an earlier, already-released
// claim) is never touched by resolving a different claim.
func unverifyClaimSourcedRows(ctx context.Context, tx pgx.Tx, guildID, userID int64) error {
	if _, err := tx.Exec(ctx,
		`update guild_characters set verified_at = null, verified_by = null
		 where guild_id = $1 and user_id = $2 and verified_by = 'claim'`, guildID, userID); err != nil {
		return fmt.Errorf("guilds: unverify claim-sourced rows for guild %d: %w", guildID, err)
	}
	return nil
}

// recordResolution writes the durable (guild, contester, outcome,
// moderator) row A3 needs to refuse a repeat contest after an uphold,
// and the audit trail every claim/resolve/settings change in this
// package already keeps via a log line (recorded here as a table row
// instead, since a repeat-contest check needs to query it back later).
func recordResolution(ctx context.Context, tx pgx.Tx, guildID, contesterID, moderatorID int64, outcome string) error {
	if _, err := tx.Exec(ctx,
		`insert into guild_claim_resolutions (guild_id, contester_id, outcome, moderator_id) values ($1, $2, $3, $4)`,
		guildID, contesterID, outcome, moderatorID); err != nil {
		return fmt.Errorf("guilds: record resolution for guild %d: %w", guildID, err)
	}
	return nil
}

// ResolveClaim is a moderator's decision on a contested claim: uphold
// (dismiss the contest, claim stands, and the contester's own
// UNVERIFIED characters in this guild are removed - a verified row
// stays, since that person is a real member who lost a dispute, not an
// impostor), release (clear the claim and un-verify the resolved
// claimant's characters whose only verification source was that
// claim), or transfer (move the claim, and the same claim-sourced
// verification, to the contesting account). Every outcome is recorded
// in guild_claim_resolutions with moderatorID (A3).
func (s *Store) ResolveClaim(ctx context.Context, guildID, moderatorID int64, outcome string) error {
	if outcome != "uphold" && outcome != "release" && outcome != "transfer" {
		return ErrInvalidOutcome
	}
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	if g.ClaimContestedAt == nil {
		return ErrNoActiveClaim
	}
	if g.ClaimContestedBy == nil {
		return fmt.Errorf("guilds: resolve claim: no contesting account recorded")
	}
	contester := *g.ClaimContestedBy
	// The account whose claim is being resolved - claimed takes
	// priority over merely-pending, mirroring ContestClaim's own
	// resolution logic. nil only when a contested claim was somehow
	// never actually claimed or pending, which contest state implies
	// cannot happen, but is guarded rather than assumed below.
	claimant := g.ClaimedBy
	if claimant == nil {
		claimant = g.ClaimPendingBy
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: resolve claim: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	switch outcome {
	case "uphold":
		if _, err := tx.Exec(ctx,
			`update guilds set claim_contested_at = null, claim_contested_by = null where id = $1`, guildID); err != nil {
			return fmt.Errorf("guilds: resolve claim (uphold): %w", err)
		}
		if _, err := tx.Exec(ctx,
			`delete from guild_characters where guild_id = $1 and user_id = $2 and verified_at is null`,
			guildID, contester); err != nil {
			return fmt.Errorf("guilds: resolve claim (uphold): remove unverified contester rows: %w", err)
		}
		if err := RecomputeMembership(ctx, tx, guildID, &contester); err != nil {
			return err
		}
	case "release":
		if _, err := tx.Exec(ctx,
			`update guilds set claimed_by = null, claimed_at = null, claim_pending_by = null, claim_requested_at = null,
			   claim_contested_at = null, claim_contested_by = null where id = $1`, guildID); err != nil {
			return fmt.Errorf("guilds: resolve claim (release): %w", err)
		}
		if claimant != nil {
			if err := unverifyClaimSourcedRows(ctx, tx, guildID, *claimant); err != nil {
				return err
			}
		}
		if err := RecomputeMembership(ctx, tx, guildID, nil); err != nil {
			return err
		}
	case "transfer":
		if _, err := tx.Exec(ctx,
			`update guilds set claimed_by = $2, claimed_at = now(), claim_pending_by = null, claim_requested_at = null,
			   claim_contested_at = null, claim_contested_by = null where id = $1`, guildID, contester); err != nil {
			return fmt.Errorf("guilds: resolve claim (transfer): %w", err)
		}
		if claimant != nil {
			if err := unverifyClaimSourcedRows(ctx, tx, guildID, *claimant); err != nil {
				return err
			}
		}
		if err := setVerifiedForAccount(ctx, tx, guildID, contester, "claim"); err != nil {
			return err
		}
		if err := RecomputeMembership(ctx, tx, guildID, nil); err != nil {
			return err
		}
	}
	if err := recordResolution(ctx, tx, guildID, contester, moderatorID, outcome); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("guilds: resolve claim: commit: %w", err)
	}
	return nil
}

// FrozenClaimant reports whether guildID's claim is currently contested
// (always frozen when it is, per the fourth security review response)
// with userID as the disputed claimant - the hook reports/handler.go's
// mayEdit consults (D, 2026-09-21 second security review response):
// while frozen, the disputed claimant's officer-derived edit right
// over their guild's reports is suspended right alongside every other
// officer power; their own reports (owned by them regardless of guild
// role) and every other verified officer or moderator are unaffected,
// since mayEdit only ever calls this after its own
// owner/moderator/officer-rank checks have already passed.
func (s *Store) FrozenClaimant(ctx context.Context, guildID, userID int64) (bool, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return false, nil
		}
		return false, err
	}
	if g.ClaimContestedAt == nil {
		return false, nil
	}
	claimant, ok := g.activeClaimant(time.Now())
	return ok && claimant == userID, nil
}

func (s *Service) contestClaim(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	u, err := s.Accounts.User(r.Context(), actor.UserID)
	if err != nil {
		s.fail(w, r, "contest_claim", err, "could not contest that claim just now")
		return
	}
	err = s.Store.ContestClaim(r.Context(), guildID, actor.UserID, u.BnetSub != "")
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case errors.Is(err, ErrNoActiveClaim):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "there is no claim on this guild to contest", nil)
	case errors.Is(err, ErrAlreadyContested):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "this claim is already contested", nil)
	case errors.Is(err, ErrGuildRecentlyUpheld):
		httpx.WriteError(w, r, http.StatusConflict, "conflict",
			"this guild's claim was upheld within the last 30 days", nil)
	case errors.Is(err, ErrSameAccount):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you cannot contest your own claim", nil)
	case errors.Is(err, ErrNotEligible):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you need an officer or leader character in this guild to contest its claim", nil)
	case errors.Is(err, ErrNoBattleNetIdentity):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"contesting a claim requires a linked Battle.net account", nil)
	case errors.Is(err, ErrContestAlreadyUpheld):
		httpx.WriteError(w, r, http.StatusConflict, "conflict",
			"this account's contest of this guild's claim has already been upheld", nil)
	case errors.Is(err, ErrContestRateLimited):
		httpx.WriteError(w, r, http.StatusTooManyRequests, "rate_limited",
			"you may only attempt one guild contest every 30 days", nil)
	case errors.Is(err, ErrAlreadyContestingAnotherGuild):
		httpx.WriteError(w, r, http.StatusConflict, "conflict",
			"you already have an open contest on another guild", nil)
	case err != nil:
		s.fail(w, r, "contest_claim", err, "could not contest that claim just now")
	default:
		// Warn, not Info: every contest is a moderation-queue event a
		// human should be alerted on (item 4, fourth security review
		// response), not routine activity.
		s.logger().Warn("guilds", "op", "claim_contest", "guild_id", guildID, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "contested"})
	}
}

type resolveClaimInput struct {
	Outcome string `json:"outcome"`
}

func (s *Service) resolveClaim(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	if !actor.IsModerator() {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "only a moderator may resolve a contested claim", nil)
		return
	}
	var in resolveClaimInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	if in.Outcome != "uphold" && in.Outcome != "release" && in.Outcome != "transfer" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a resolution outcome",
			map[string]string{"outcome": "one of uphold, release, transfer"})
		return
	}
	err := s.Store.ResolveClaim(r.Context(), guildID, actor.UserID, in.Outcome)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case errors.Is(err, ErrNoActiveClaim):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "there is no contested claim on this guild", nil)
	case err != nil:
		s.fail(w, r, "resolve_claim", err, "could not resolve that claim just now")
	default:
		s.logger().Info("guilds", "op", "claim_resolve", "guild_id", guildID,
			"moderator_id", actor.UserID, "outcome", in.Outcome)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "resolved", "outcome": in.Outcome})
	}
}

// reopenClaimCooldown answers 404, not 403, for a non-moderator - the
// same hidden-standing pattern moderationClaims uses (fifth security
// review response), checked before the guild id is even parsed so a
// non-moderator learns nothing about whether the guild exists either.
func (s *Service) reopenClaimCooldown(w http.ResponseWriter, r *http.Request) {
	if !auth.ActorFrom(r.Context()).IsModerator() {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such route", nil)
		return
	}
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	if err := s.Store.ReopenClaimCooldown(r.Context(), guildID); err != nil {
		if errors.Is(err, ErrNotFound) {
			httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
			return
		}
		s.fail(w, r, "reopen_claim_cooldown", err, "could not reopen that guild's claim cooldown just now")
		return
	}
	s.logger().Info("guilds", "op", "claim_reopen", "guild_id", guildID, "moderator_id", actor.UserID)
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "reopened"})
}
