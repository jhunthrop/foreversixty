// api/internal/guilds/contest.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

var (
	ErrNoActiveClaim    = errors.New("guilds: no claim to contest")
	ErrAlreadyContested = errors.New("guilds: claim already contested")
	ErrInvalidOutcome   = errors.New("guilds: invalid resolution outcome")
)

// ClaimStateView is the claim state object GET /v1/guilds/{id}/settings
// and GET /v1/guilds/{id}/home both expose (2026-09-21 security review
// response, spec §2.4's amendment).
type ClaimStateView struct {
	State string     `json:"state"`
	Since *time.Time `json:"since,omitempty"`
}

// claimState derives the four-phase claim state a guild is in as of now.
func claimState(g Guild, now time.Time) ClaimStateView {
	switch {
	case g.ClaimContestedAt != nil:
		return ClaimStateView{State: "contested", Since: g.ClaimContestedAt}
	case g.ClaimedBy != nil:
		return ClaimStateView{State: "claimed"}
	case g.pendingActive(now):
		return ClaimStateView{State: "pending", Since: g.ClaimRequestedAt}
	default:
		return ClaimStateView{State: "unclaimed"}
	}
}

// ContestClaim flags guildID's current claim (pending or claimed) as
// disputed by contesterID, freezing the claimant's officer powers until
// a moderator resolves it. The contester must hold a raw officer/leader
// rank character in this guild (eligibleClaimRank - the same raw,
// unverified-tolerant check the claim flow itself uses, since disputing
// a possibly-forged claim cannot itself require verification the dispute
// exists to question) on a different account from the current
// claimant/pending claimant.
func (s *Store) ContestClaim(ctx context.Context, guildID, contesterID int64) error {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	current := g.ClaimedBy
	if current == nil && g.pendingActive(time.Now()) {
		current = g.ClaimPendingBy
	}
	if current == nil {
		return ErrNoActiveClaim
	}
	if g.ClaimContestedAt != nil {
		return ErrAlreadyContested
	}
	if *current == contesterID {
		return ErrSameAccount
	}
	if _, ok, err := s.eligibleClaimRank(ctx, guildID, contesterID); err != nil {
		return err
	} else if !ok {
		return ErrNotEligible
	}
	if _, err := s.Pool.Exec(ctx,
		`update guilds set claim_contested_at = now(), claim_contested_by = $2 where id = $1`,
		guildID, contesterID); err != nil {
		return fmt.Errorf("guilds: contest claim: %w", err)
	}
	return nil
}

// ResolveClaim is a moderator's decision on a contested claim: uphold
// (dismiss the contest, claim stands), release (clear the claim and
// un-verify every character whose only verification source was the
// claim), or transfer (move the claim, and the same claim-sourced
// verification, to the contesting account).
func (s *Store) ResolveClaim(ctx context.Context, guildID int64, outcome string) error {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	if g.ClaimContestedAt == nil {
		return ErrNoActiveClaim
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
	case "release":
		if _, err := tx.Exec(ctx,
			`update guilds set claimed_by = null, claim_pending_by = null, claim_requested_at = null,
			   claim_contested_at = null, claim_contested_by = null where id = $1`, guildID); err != nil {
			return fmt.Errorf("guilds: resolve claim (release): %w", err)
		}
		if _, err := tx.Exec(ctx,
			`update guild_characters set verified_at = null, verified_by = null
			 where guild_id = $1 and verified_by = 'claim'`, guildID); err != nil {
			return fmt.Errorf("guilds: resolve claim (release): un-verify: %w", err)
		}
		if err := RecomputeMembership(ctx, tx, guildID, nil); err != nil {
			return err
		}
	case "transfer":
		if g.ClaimContestedBy == nil {
			return fmt.Errorf("guilds: resolve claim (transfer): no contesting account recorded")
		}
		newClaimant := *g.ClaimContestedBy
		if _, err := tx.Exec(ctx,
			`update guilds set claimed_by = $2, claim_pending_by = null, claim_requested_at = null,
			   claim_contested_at = null, claim_contested_by = null where id = $1`, guildID, newClaimant); err != nil {
			return fmt.Errorf("guilds: resolve claim (transfer): %w", err)
		}
		if _, err := tx.Exec(ctx,
			`update guild_characters set verified_at = null, verified_by = null
			 where guild_id = $1 and verified_by = 'claim'`, guildID); err != nil {
			return fmt.Errorf("guilds: resolve claim (transfer): un-verify: %w", err)
		}
		if err := setVerifiedForAccount(ctx, tx, guildID, newClaimant, "claim"); err != nil {
			return err
		}
		if err := RecomputeMembership(ctx, tx, guildID, nil); err != nil {
			return err
		}
	default:
		return ErrInvalidOutcome
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("guilds: resolve claim: commit: %w", err)
	}
	return nil
}

func (s *Service) contestClaim(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	err := s.Store.ContestClaim(r.Context(), guildID, actor.UserID)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case errors.Is(err, ErrNoActiveClaim):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "there is no claim on this guild to contest", nil)
	case errors.Is(err, ErrAlreadyContested):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "this claim is already contested", nil)
	case errors.Is(err, ErrSameAccount):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you cannot contest your own claim", nil)
	case errors.Is(err, ErrNotEligible):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you need an officer or leader character in this guild to contest its claim", nil)
	case err != nil:
		s.fail(w, r, "contest_claim", err, "could not contest that claim just now")
	default:
		s.logger().Info("guilds", "op", "claim_contest", "guild_id", guildID, "user_id", actor.UserID)
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
	err := s.Store.ResolveClaim(r.Context(), guildID, in.Outcome)
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
