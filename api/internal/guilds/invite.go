// api/internal/guilds/invite.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// InviteTokenPrefix marks an invite token wherever it appears, the same
// way auth.DeviceTokenPrefix does for a device token.
const InviteTokenPrefix = "fsg_"

// inviteTokenChars is the random half of an invite token.
const inviteTokenChars = 32

// syntheticKey is the character_key an invite-joined account gets before
// it ever syncs a real character — a namespace that never collides with a
// real character.Key output (always region/ruleset/name-slug, never
// prefixed "account:").
func syntheticKey(userID int64) string {
	return fmt.Sprintf("account:%d", userID)
}

// RotateInvite issues a fresh invite token for guildID, storing only its
// SHA-256 hash. The token is returned once; nothing here or later can
// recover it.
func (s *Store) RotateInvite(ctx context.Context, guildID int64) (string, time.Time, error) {
	token := InviteTokenPrefix + auth.Base32ID(inviteTokenChars)
	hash := auth.TokenHash(token)
	var rotatedAt time.Time
	err := s.Pool.QueryRow(ctx,
		`update guilds set invite_token_hash = $2, invite_token_rotated_at = now()
		 where id = $1 returning invite_token_rotated_at`, guildID, hash).Scan(&rotatedAt)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("guilds: rotate invite for %d: %w", guildID, err)
	}
	return token, rotatedAt, nil
}

// InviteAccept is what AcceptInvite answers with.
type InviteAccept struct {
	GuildID               int64
	Region, Ruleset, Name string
	Rank                  string
}

// AcceptInvite redeems token for userID: a verified, source="invite" row
// under the synthetic account: key, since an invite-joined account has no
// real character to anchor a normal row by. An officer sharing the link
// is itself the corroboration, so verified_at is set immediately. If the
// account already holds a synthetic row for a different guild, that row
// is transferred (deleted, then recreated here) rather than left behind —
// the unique index on character_key allows only one guild at a time,
// synthetic rows included.
func (s *Store) AcceptInvite(ctx context.Context, token string, userID int64) (InviteAccept, error) {
	hash := auth.TokenHash(token)
	var g Guild
	err := s.Pool.QueryRow(ctx,
		`select id, region, ruleset, name from guilds where invite_token_hash = $1`, hash).
		Scan(&g.ID, &g.Region, &g.Ruleset, &g.Name)
	if err != nil {
		// Never distinguish "no such token" from "rotated away" - both a
		// bad token and a hash that no longer matches any row answer the
		// same ErrNotFound here.
		return InviteAccept{}, ErrNotFound
	}

	key := syntheticKey(userID)
	var prevGuildID *int64
	err = s.Pool.QueryRow(ctx,
		`select guild_id from guild_characters where character_key = $1`, key).Scan(&prevGuildID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return InviteAccept{}, fmt.Errorf("guilds: accept invite: read previous: %w", err)
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return InviteAccept{}, fmt.Errorf("guilds: accept invite: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if prevGuildID != nil && *prevGuildID != g.ID {
		if _, err := tx.Exec(ctx,
			`delete from guild_characters where character_key = $1 and guild_id = $2`, key, *prevGuildID); err != nil {
			return InviteAccept{}, fmt.Errorf("guilds: accept invite: clear previous: %w", err)
		}
		if err := RecomputeMembership(ctx, tx, *prevGuildID, &userID); err != nil {
			return InviteAccept{}, err
		}
		if err := ReleaseClaimIfLost(ctx, tx, *prevGuildID, userID); err != nil {
			return InviteAccept{}, err
		}
	}
	if _, err := tx.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank, source, verified_at, verified_by, refreshed_at)
		 values ($1, $2, $3, 'member', 'invite', now(), 'invite', now())
		 on conflict (guild_id, character_key) do update set
		   user_id = excluded.user_id, refreshed_at = now(),
		   verified_at = coalesce(guild_characters.verified_at, now()),
		   verified_by = coalesce(guild_characters.verified_by, 'invite')`,
		g.ID, key, userID); err != nil {
		return InviteAccept{}, fmt.Errorf("guilds: accept invite: %w", err)
	}
	if err := RecomputeMembership(ctx, tx, g.ID, &userID); err != nil {
		return InviteAccept{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return InviteAccept{}, fmt.Errorf("guilds: accept invite: commit: %w", err)
	}
	return InviteAccept{GuildID: g.ID, Region: g.Region, Ruleset: g.Ruleset, Name: g.Name, Rank: "member"}, nil
}

func (s *Service) rotateInvite(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "rotate_invite", err, "could not rotate that invite just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to rotate its invite link", nil)
		return
	}
	if contested, err := s.Store.contested(r.Context(), guildID); err != nil {
		s.fail(w, r, "rotate_invite", err, "could not rotate that invite just now")
		return
	} else if contested {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return
	}
	token, rotatedAt, err := s.Store.RotateInvite(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, "rotate_invite", err, "could not rotate that invite just now")
		return
	}
	s.logger().Info("guilds", "op", "invite_rotate", "guild_id", guildID, "user_id", actor.UserID)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{
		"token": token, "url": "/guild/invite/" + token, "rotated_at": rotatedAt,
	})
}

func (s *Service) acceptInvite(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")
	actor := auth.ActorFrom(r.Context())
	result, err := s.Store.AcceptInvite(r.Context(), token, actor.UserID)
	switch {
	case errors.Is(err, ErrNotFound):
		// Never distinguish "no such token" from "rotated away" — both
		// fall out of the same zero-row lookup in AcceptInvite.
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "that invite link is not valid", nil)
	case err != nil:
		s.fail(w, r, "accept_invite", err, "could not accept that invite just now")
	default:
		s.logger().Info("guilds", "op", "invite_accept", "guild_id", result.GuildID, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]any{
			"guild": map[string]any{
				"id": result.GuildID, "region": result.Region, "ruleset": result.Ruleset, "name": result.Name,
			},
			"rank": result.Rank,
		})
	}
}
