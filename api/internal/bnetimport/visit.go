// api/internal/bnetimport/visit.go
package bnetimport

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// visitStaleAfter is how old an account's last sync (users.bnet_synced_at) may be before a
// signed-in visit refreshes it. The visit claims the column with a compare-and-set, so it
// is also the per-account cooldown: one visit refresh per account per visitStaleAfter,
// across every instance, whether or not the refresh succeeded.
const visitStaleAfter = time.Hour

// visitBudget bounds one visit refresh. It runs after the request answered, on its own
// context.
const visitBudget = 3 * time.Minute

// visitDue reports whether an account whose last sync started at lastSync is due a visit
// refresh at now. A never-synced account is due.
func visitDue(lastSync *time.Time, now time.Time) bool {
	return lastSync == nil || now.Sub(*lastSync) >= visitStaleAfter
}

// Visited is the authenticator's hook on every signed-in request (auth.VisitHook). For a
// Battle.net account whose last sync is older than visitStaleAfter it starts a background
// refresh and returns at once: the request never waits. A refresh already running for the
// account in this process, or a 429 backoff, starts nothing.
func (s *Service) Visited(u auth.User) {
	now := time.Now()
	if u.BnetSub == "" || !visitDue(u.BnetSyncedAt, now) || s.backoff.blocked(now) {
		return
	}
	if _, running := s.visiting.LoadOrStore(u.ID, struct{}{}); running {
		return
	}
	go func() {
		defer s.visiting.Delete(u.ID)
		ctx, cancel := context.WithTimeout(context.Background(), visitBudget)
		defer cancel()
		if err := s.RefreshAccount(ctx, u.ID); err != nil {
			s.logger().Warn("bnetimport", "op", "visit_refresh", "user_id", u.ID, "err", err)
		}
	}()
}

// RefreshAccount claims the account's sync and refreshes it: through the account's stored
// access token while it is usable (the sign-in import's own path, which also finds new
// characters and levels), else character by character from public profiles, which need no
// user token. A token Battle.net refuses is forgotten and the public path runs instead.
// An account another visit claimed within visitStaleAfter is left alone.
func (s *Service) RefreshAccount(ctx context.Context, userID int64) error {
	claimed, err := s.claimSync(ctx, userID)
	if err != nil || !claimed {
		return err
	}
	token, err := s.liveToken(ctx, userID)
	if err != nil {
		return err
	}
	if token != "" {
		summary, err := s.ImportAccount(ctx, userID, token)
		if !errors.Is(err, ErrTokenRejected) {
			s.logger().Info("bnetimport", "op", "visit_refresh", "user_id", userID, "path", "account",
				"characters", summary.Characters, "pending", summary.Pending, "rate_limited", summary.RateLimited)
			return err
		}
		if err := s.forgetToken(ctx, userID); err != nil {
			return err
		}
	}
	return s.refreshPublic(ctx, userID)
}

// refreshPublic re-syncs every one of the account's 'bnet' characters from public profiles.
func (s *Service) refreshPublic(ctx context.Context, userID int64) error {
	chars, err := s.accountBnetCharacters(ctx, userID)
	if err != nil {
		return err
	}
	result, err := s.refreshBatch(ctx, chars)
	s.logger().Info("bnetimport", "op", "visit_refresh", "user_id", userID, "path", "public",
		"considered", result.Considered, "refreshed", result.Refreshed, "rate_limited", result.RateLimited)
	return err
}

// claimSync stamps the account's bnet_synced_at when it is due (visitDue's rule, in SQL)
// and reports whether this call did. Two visits racing on two instances: one wins.
func (s *Service) claimSync(ctx context.Context, userID int64) (bool, error) {
	tag, err := s.Pool.Exec(ctx,
		`update users set bnet_synced_at = now()
		 where id = $1 and bnet_sub is not null
		   and (bnet_synced_at is null or bnet_synced_at <= now() - $2::interval)`,
		userID, intervalArg(visitStaleAfter))
	if err != nil {
		return false, fmt.Errorf("bnetimport: claim sync: %w", err)
	}
	return tag.RowsAffected() == 1, nil
}
