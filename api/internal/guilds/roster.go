// api/internal/guilds/roster.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// ErrInvalidConsent is returned for a consent value outside roster,
// gear, gear_bags.
var ErrInvalidConsent = errors.New("guilds: invalid consent")

var validConsent = map[string]bool{"roster": true, "gear": true, "gear_bags": true}

// ApproveCharacter sets verified_at on a guild_characters row that does
// not already have it - an officer's social-trust decision, available
// immediately with no wait for a second raid night.
func (s *Store) ApproveCharacter(ctx context.Context, guildID int64, characterKey string) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: approve character: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	var userID int64
	err = tx.QueryRow(ctx,
		`update guild_characters set verified_at = coalesce(verified_at, now())
		 where guild_id = $1 and character_key = $2 returning user_id`,
		guildID, characterKey).Scan(&userID)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("guilds: approve character: %w", err)
	}
	if err := RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// CharacterOwner reads which account a guild_characters row belongs to,
// so a handler can decide "the character's own account" before acting.
func (s *Store) CharacterOwner(ctx context.Context, guildID int64, characterKey string) (int64, error) {
	var userID int64
	err := s.Pool.QueryRow(ctx,
		`select user_id from guild_characters where guild_id = $1 and character_key = $2`,
		guildID, characterKey).Scan(&userID)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, ErrNotFound
	}
	if err != nil {
		return 0, fmt.Errorf("guilds: character owner: %w", err)
	}
	return userID, nil
}

// RemoveCharacter deletes a guild_characters row and recomputes the
// account it belonged to. A wrongly-verified character an officer
// removes stays removed: nothing re-verifies it on its next export
// unless it earns verification again through one of the corroboration
// paths.
func (s *Store) RemoveCharacter(ctx context.Context, guildID int64, characterKey string) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: remove character: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	var userID int64
	err = tx.QueryRow(ctx,
		`delete from guild_characters where guild_id = $1 and character_key = $2 returning user_id`,
		guildID, characterKey).Scan(&userID)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("guilds: remove character: %w", err)
	}
	if err := RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
		return err
	}
	if err := ReleaseClaimIfLost(ctx, tx, guildID, userID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// SetConsent updates the caller's own per-guild consent level.
func (s *Store) SetConsent(ctx context.Context, guildID, userID int64, consent string) error {
	if !validConsent[consent] {
		return ErrInvalidConsent
	}
	tag, err := s.Pool.Exec(ctx,
		`update guild_members set consent = $3 where guild_id = $1 and user_id = $2`, guildID, userID, consent)
	if err != nil {
		return fmt.Errorf("guilds: set consent: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// Leave removes every one of the caller's own guild_characters rows in
// guildID, then recomputes and releases a lost claim.
func (s *Store) Leave(ctx context.Context, guildID, userID int64) error {
	ok, err := s.IsMember(ctx, guildID, userID)
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotFound
	}
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: leave: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx, `delete from guild_characters where guild_id = $1 and user_id = $2`, guildID, userID); err != nil {
		return fmt.Errorf("guilds: leave: %w", err)
	}
	if err := RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
		return err
	}
	if err := ReleaseClaimIfLost(ctx, tx, guildID, userID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// characterKeyFrom builds a character key from the three path segments
// {region}/{ruleset}/{name} every roster route takes.
func characterKeyFrom(r *http.Request) (string, bool) {
	region := strings.ToLower(r.PathValue("region"))
	ruleset := strings.ToLower(r.PathValue("ruleset"))
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
		return "", false
	}
	return character.Key(region, ruleset, r.PathValue("name")), true
}

func (s *Service) approveCharacter(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	key, ok := characterKeyFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "approve", err, "could not approve that character just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to approve a character", nil)
		return
	}
	err = s.Store.ApproveCharacter(r.Context(), guildID, key)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character on this guild's roster", nil)
	case err != nil:
		s.fail(w, r, "approve", err, "could not approve that character just now")
	default:
		s.logger().Info("guilds", "op", "character_approve", "guild_id", guildID, "character_key", key,
			"user_id", auth.ActorFrom(r.Context()).UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"character_key": key, "status": "approved"})
	}
}

func (s *Service) removeCharacter(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	key, ok := characterKeyFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	owner, err := s.Store.CharacterOwner(r.Context(), guildID, key)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character on this guild's roster", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "remove", err, "could not remove that character just now")
		return
	}
	if owner != actor.UserID {
		allowed, err := s.verifiedOfficerOrLeader(r, guildID)
		if err != nil {
			s.fail(w, r, "remove", err, "could not remove that character just now")
			return
		}
		if !allowed {
			httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
				"you must be the character's own account or a verified officer to remove it", nil)
			return
		}
	}
	if err := s.Store.RemoveCharacter(r.Context(), guildID, key); err != nil {
		if errors.Is(err, ErrNotFound) {
			httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character on this guild's roster", nil)
			return
		}
		s.fail(w, r, "remove", err, "could not remove that character just now")
		return
	}
	s.logger().Info("guilds", "op", "character_remove", "guild_id", guildID, "character_key", key, "user_id", actor.UserID)
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "removed"})
}

type consentInput struct {
	Consent string `json:"consent"`
}

func (s *Service) patchConsent(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	var in consentInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	actor := auth.ActorFrom(r.Context())
	err := s.Store.SetConsent(r.Context(), guildID, actor.UserID, in.Consent)
	switch {
	case errors.Is(err, ErrInvalidConsent):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a consent level",
			map[string]string{"consent": "one of roster, gear, gear_bags"})
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "you are not a member of that guild", nil)
	case err != nil:
		s.fail(w, r, "consent", err, "could not change that setting just now")
	default:
		s.logger().Info("guilds", "op", "consent_update", "guild_id", guildID, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"consent": in.Consent})
	}
}

func (s *Service) leaveGuild(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	err := s.Store.Leave(r.Context(), guildID, actor.UserID)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "you are not a member of that guild", nil)
	case err != nil:
		s.fail(w, r, "leave", err, "could not leave that guild just now")
	default:
		s.logger().Info("guilds", "op", "leave", "guild_id", guildID, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "left"})
	}
}
