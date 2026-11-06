// api/internal/guilds/invite_shown.go
//
// Invite links an officer can show again (spec 2026-11-05 guild gating and invites §4.1,
// migration 0070). The hash stays the only lookup key for redeeming; the same token is also kept
// sealed (TokenSealer, the Discord webhook envelope) so a verified officer can copy the current
// link on a later visit, and the account that made it is kept so settings and the join page can
// name it. The token is never logged and never leaves this package except to a verified officer.
package guilds

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// TokenSealer seals the invite token at rest; *notify.Sealer (AES-256-GCM under
// NOTIFY_ENCRYPTION_KEY) satisfies it.
type TokenSealer interface {
	Seal(plain string) ([]byte, error)
	Open(box []byte) (string, error)
}

// errInviteUnreadable is a sealed token that no longer opens: the key changed, or the box was
// altered. The link is then treated as one that cannot be shown again.
var errInviteUnreadable = errors.New("guilds: sealed invite token does not open")

// InviteView is the guild's current invite link as settings reads it. URL is set only for a
// verified officer or leader (Service.withInviteURL), and only while the kept token is the live
// one; MadeBy is the maker's character on this guild's roster, nil when it has none.
type InviteView struct {
	RotatedAt *time.Time      `json:"rotated_at"`
	URL       *string         `json:"url"`
	MadeBy    *ClaimHolderRef `json:"made_by"`
}

// invitePath is the site path a token opens, the one shape rotate and settings both answer.
func invitePath(token string) string {
	return "/guild/invite/" + token
}

// sealInvite is token sealed for storage, or nil when no sealer is wired (the link is then
// shown once, right after it is made).
func (s *Store) sealInvite(token string) ([]byte, error) {
	if s.InviteSealer == nil {
		return nil, nil
	}
	box, err := s.InviteSealer.Seal(token)
	if err != nil {
		return nil, fmt.Errorf("guilds: seal invite: %w", err)
	}
	return box, nil
}

// inviteView is g's invite as any settings reader sees it: when it was made and by whom.
func (s *Store) inviteView(ctx context.Context, g Guild) (InviteView, error) {
	view := InviteView{RotatedAt: g.InviteTokenRotatedAt}
	if g.InviteTokenRotatedAt == nil || g.InviteRotatedBy == nil {
		return view, nil
	}
	by, err := s.claimHolderRef(ctx, g.ID, *g.InviteRotatedBy)
	if err != nil {
		return InviteView{}, err
	}
	view.MadeBy = by
	return view, nil
}

// InviteURL is the path of guildID's current invite link, read back from its sealed token. Nil
// when there is no link, no sealer, no kept token, or a kept token that is no longer the live
// one (a build without 0070's columns rotated it: the hash moved, the box did not). A box that
// does not open answers errInviteUnreadable. Callers must have checked the reader is a verified
// officer or leader of the guild.
func (s *Store) InviteURL(ctx context.Context, guildID int64) (*string, error) {
	var box, hash []byte
	err := s.Pool.QueryRow(ctx,
		`select invite_token_enc, invite_token_hash from guilds where id = $1`, guildID).Scan(&box, &hash)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: read invite for %d: %w", guildID, err)
	}
	if s.InviteSealer == nil || box == nil || hash == nil {
		return nil, nil
	}
	token, err := s.InviteSealer.Open(box)
	if err != nil {
		return nil, errInviteUnreadable
	}
	if !bytes.Equal(auth.TokenHash(token), hash) {
		return nil, nil
	}
	path := invitePath(token)
	return &path, nil
}

// withInviteURL is view with the current link's URL when the request's account is a verified
// officer or leader of guildID. A moderator reads settings without it: the link is a membership
// capability, and only the guild's own officers hand it out. A box that no longer opens is logged
// (never the token) and read as a link that cannot be shown again.
func (s *Service) withInviteURL(r *http.Request, guildID int64, view SettingsView) (SettingsView, error) {
	officer, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil || !officer {
		return view, err
	}
	url, err := s.Store.InviteURL(r.Context(), guildID)
	if errors.Is(err, errInviteUnreadable) {
		s.logger().Warn("guilds", "op", "invite_unseal", "guild_id", guildID, "err", err)
		return view, nil
	}
	if err != nil {
		return view, err
	}
	shown := view
	shown.Invite.URL = url
	return shown, nil
}

// rotatedInviteView is the settings shape of a link just made by userID: the same InviteView a
// later settings read answers, its URL set only when the token was kept (an InviteSealer is
// wired), so the page knows whether the link can be shown again.
func (s *Service) rotatedInviteView(ctx context.Context, guildID, userID int64, token string, rotatedAt time.Time) (InviteView, error) {
	by, err := s.Store.claimHolderRef(ctx, guildID, userID)
	if err != nil {
		return InviteView{}, err
	}
	view := InviteView{RotatedAt: &rotatedAt, MadeBy: by}
	if s.Store.InviteSealer != nil {
		path := invitePath(token)
		view.URL = &path
	}
	return view, nil
}
