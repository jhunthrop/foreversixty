package linkedrole

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// syncWithin bounds one account's sync (a refresh and a push).
const syncWithin = 15 * time.Second

// Store is what the syncer needs from discord_role_links; *PGStore satisfies it.
type Store interface {
	Get(ctx context.Context, userID int64) (Link, bool, error)
	Save(ctx context.Context, l Link) error
	Rotate(ctx context.Context, userID int64, oldBox, newBox []byte) (bool, error)
	DeleteIf(ctx context.Context, userID int64, box []byte) (bool, error)
	MarkPushed(ctx context.Context, userID int64, premium bool) error
	Profile(ctx context.Context, userID int64) (Profile, error)
	LinkedUsers(ctx context.Context, guildID int64) ([]int64, error)
}

// Sealer seals refresh tokens at rest; *notify.Sealer satisfies it.
type Sealer interface {
	Seal(plain string) ([]byte, error)
	Open(box []byte) (string, error)
}

// Syncer keeps each linked account's role connection equal to its Premium.
type Syncer struct {
	Store   Store
	Sealer  Sealer
	Discord Discord
	Log     *slog.Logger
}

func (s *Syncer) log() *slog.Logger {
	if s.Log == nil {
		return slog.Default()
	}
	return s.Log
}

// Link saves a fresh authorization and pushes the account's Premium with its access token. An
// earlier link of the account to another Discord account is told premium false first (best
// effort), so swapping Discord accounts never leaves a stale Premium behind.
func (s *Syncer) Link(ctx context.Context, userID int64, discordUserID string, t Tokens) error {
	if prev, ok, err := s.Store.Get(ctx, userID); err == nil && ok && prev.DiscordUserID != discordUserID {
		if err := s.withAccess(ctx, prev, func(access string) error {
			return s.Discord.Push(ctx, access, Connection{Premium: false})
		}); err != nil {
			s.log().Warn("linkedrole", "op", "clear previous", "user_id", userID, "err", err)
		}
	}
	box, err := s.Sealer.Seal(t.Refresh)
	if err != nil {
		return fmt.Errorf("linkedrole: seal: %w", err)
	}
	if err := s.Store.Save(ctx, Link{UserID: userID, DiscordUserID: discordUserID, RefreshBox: box}); err != nil {
		return err
	}
	return s.push(ctx, userID, t.Access)
}

// SyncUser re-pushes the account's Premium; an account with no link is a no-op, and a link
// Discord has revoked is removed.
func (s *Syncer) SyncUser(ctx context.Context, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, syncWithin)
	defer cancel()
	link, ok, err := s.Store.Get(ctx, userID)
	if err != nil || !ok {
		return err
	}
	return s.withAccess(ctx, link, func(access string) error { return s.push(ctx, userID, access) })
}

// withAccess refreshes the link's token, stores the rotated one and runs fn with the access
// token. A revoked token removes the link (unless a racing sync rotated it) and runs nothing.
func (s *Syncer) withAccess(ctx context.Context, link Link, fn func(access string) error) error {
	refresh, err := s.Sealer.Open(link.RefreshBox)
	if err != nil {
		return fmt.Errorf("linkedrole: open %d: %w", link.UserID, err)
	}
	t, err := s.Discord.Refresh(ctx, refresh)
	if errors.Is(err, ErrRevoked) {
		removed, delErr := s.Store.DeleteIf(ctx, link.UserID, link.RefreshBox)
		s.log().Info("linkedrole", "op", "revoked", "user_id", link.UserID, "removed", removed)
		return delErr
	}
	if err != nil {
		return fmt.Errorf("linkedrole: refresh %d: %w", link.UserID, err)
	}
	box, err := s.Sealer.Seal(t.Refresh)
	if err != nil {
		return fmt.Errorf("linkedrole: seal: %w", err)
	}
	if _, err := s.Store.Rotate(ctx, link.UserID, link.RefreshBox, box); err != nil {
		return err
	}
	return fn(t.Access)
}

// push reads the account's Premium and battletag and writes the role connection. An account
// that asked to be anonymous shows no username.
func (s *Syncer) push(ctx context.Context, userID int64, access string) error {
	p, err := s.Store.Profile(ctx, userID)
	if err != nil {
		return err
	}
	c := Connection{Premium: p.Premium}
	if !p.Anonymize {
		c.Username = p.Battletag
	}
	if err := s.Discord.Push(ctx, access, c); err != nil {
		return fmt.Errorf("linkedrole: push %d: %w", userID, err)
	}
	return s.Store.MarkPushed(ctx, userID, p.Premium)
}

// OnEntitlementChange is entitlements.Store's OnChange: it re-pushes the accounts whose Premium
// the change can move (the user, or a guild's linked members). Failures are logged, never
// returned: the entitlement write already stands, and the daily reconcile retries.
func (s *Syncer) OnEntitlementChange(ctx context.Context, subj entitlements.Subject) {
	var ids []int64
	switch {
	case subj.UserID != nil:
		ids = []int64{*subj.UserID}
	case subj.GuildID != nil:
		var err error
		if ids, err = s.Store.LinkedUsers(ctx, *subj.GuildID); err != nil {
			s.log().Error("linkedrole", "op", "entitlement change", "err", err)
			return
		}
	}
	for _, id := range ids {
		if err := s.SyncUser(ctx, id); err != nil {
			s.log().Warn("linkedrole", "op", "entitlement change", "user_id", id, "err", err)
		}
	}
}

// ReconcileResult counts a reconcile.
type ReconcileResult struct{ Synced, Failed int }

// Reconcile re-pushes every linked account (billing-daily). It fails only when every account
// failed.
func (s *Syncer) Reconcile(ctx context.Context) (ReconcileResult, error) {
	ids, err := s.Store.LinkedUsers(ctx, 0)
	if err != nil {
		return ReconcileResult{}, err
	}
	var r ReconcileResult
	var last error
	for _, id := range ids {
		if err := s.SyncUser(ctx, id); err != nil {
			r.Failed++
			last = err
			s.log().Warn("linkedrole", "op", "reconcile", "user_id", id, "err", err)
			continue
		}
		r.Synced++
	}
	if r.Failed > 0 && r.Synced == 0 {
		return r, fmt.Errorf("linkedrole: reconcile: every account failed: %w", last)
	}
	return r, nil
}
