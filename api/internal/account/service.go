// api/internal/account/service.go
package account

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

// ErrConfirmMismatch is a typed confirmation that is neither the BattleTag nor the word.
var ErrConfirmMismatch = errors.New("account: the typed confirmation does not match")

// ErrBilling marks a Stripe cancel that did not happen (no gateway, or Stripe refused): the
// deletion stopped before anything was removed.
var ErrBilling = errors.New("account: live subscriptions could not be cancelled")

// ErrStaffAccount is a moderator or admin asking to delete itself.
var ErrStaffAccount = errors.New("account: staff accounts are closed by support")

// Canceller cancels every live Stripe subscription the account pays for, at once;
// *billing.Service satisfies it.
type Canceller interface {
	CancelAccountSubscriptions(ctx context.Context, userID int64) error
}

// RoleRevoker clears the Discord Premium role and revokes the stored authorization;
// *linkedrole.Syncer satisfies it.
type RoleRevoker interface {
	Revoke(ctx context.Context, userID int64) error
}

// ObjectDeleter removes one stored object; *r2.Client satisfies it.
type ObjectDeleter interface {
	Delete(ctx context.Context, key string) error
}

// Service deletes accounts. Billing is required (a deployment without Stripe still answers
// for an account with nothing to cancel); Discord and Objects are nil when the deployment has
// no linked-role syncer or object storage, and their step is skipped.
type Service struct {
	Store   *Store
	Billing Canceller
	Discord RoleRevoker
	Objects ObjectDeleter
	Log     *slog.Logger
	// Now is the clock; nil means time.Now.
	Now func() time.Time
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) now() time.Time {
	if s.Now != nil {
		return s.Now()
	}
	return time.Now()
}

// Delete deletes userID's account when typed confirms it. The refusals come first and
// change nothing: a staff account, a mismatched confirmation, a guild obligation. Then the
// live subscriptions are cancelled at Stripe (a failure stops here, with nothing deleted:
// deleting first could leave a deleted account still being charged), the Discord role is
// cleared (best effort: the transaction deletes the stored token either way), the
// transaction erases the account, and the deleted reports' storage objects are removed
// (best effort, logged). It answers when the account was deleted.
func (s *Service) Delete(ctx context.Context, userID int64, typed string) (time.Time, error) {
	acct, err := s.Store.Account(ctx, userID)
	if err != nil {
		return time.Time{}, err
	}
	if acct.Staff() {
		return time.Time{}, ErrStaffAccount
	}
	if !Confirmed(typed, acct.Battletag) {
		return time.Time{}, ErrConfirmMismatch
	}
	obligations, err := s.Store.Obligations(ctx, userID)
	if err != nil {
		return time.Time{}, err
	}
	if len(obligations) > 0 {
		return time.Time{}, &ObligationError{Guilds: obligations}
	}
	if err := s.Billing.CancelAccountSubscriptions(ctx, userID); err != nil {
		return time.Time{}, fmt.Errorf("%w: account %d: %w", ErrBilling, userID, err)
	}
	s.revokeDiscord(ctx, userID)
	erased, err := s.Store.Erase(ctx, userID)
	if err != nil {
		return time.Time{}, err
	}
	at := s.now().UTC()
	// The audit line: the account id and the time, nothing that names the person.
	s.logger().Info("account", "op", "deleted", "user_id", userID, "at", at.Format(time.RFC3339))
	s.deleteObjects(ctx, userID, erased.ObjectKeys)
	return at, nil
}

// revokeDiscord runs the Discord step and logs a failure; the stored token goes with the
// account row regardless.
func (s *Service) revokeDiscord(ctx context.Context, userID int64) {
	if s.Discord == nil {
		return
	}
	if err := s.Discord.Revoke(ctx, userID); err != nil {
		s.logger().Warn("account", "op", "discord_revoke", "user_id", userID, "err", err)
	}
}

// deleteObjects removes the deleted reports' and uploads' stored objects. The rows are gone,
// so nothing can serve them any more; a failure is logged for a hand clean-up.
func (s *Service) deleteObjects(ctx context.Context, userID int64, keys []string) {
	if s.Objects == nil || len(keys) == 0 {
		return
	}
	failed := 0
	for _, key := range keys {
		if err := s.Objects.Delete(ctx, key); err != nil {
			failed++
			s.logger().Warn("account", "op", "object_delete", "user_id", userID, "key", key, "err", err)
		}
	}
	if failed > 0 {
		s.logger().Error("account", "op", "object_delete", "user_id", userID, "failed", failed, "of", len(keys))
	}
}
