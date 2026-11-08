// api/internal/guilds/ready_check_handler.go
//
// GET /v1/characters/{region}/{ruleset}/{name}/ready?raid={id}: the signed-in owner's ready check
// for one of their characters (ready_check.go). It reads the bags, so only the owner reads it:
// anyone else, a character the site does not have, or a raid of a guild the character is not
// verified in reads 404.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const readyFailMessage = "could not load your ready check just now"

func (s *Service) readyCheck(w http.ResponseWriter, r *http.Request) {
	key, ok := characterKeyFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	owned, err := s.Store.isOwnCharacter(r.Context(), auth.ActorFrom(r.Context()).UserID, key)
	if err != nil {
		s.fail(w, r, "ready_check", err, readyFailMessage)
		return
	}
	if !owned {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character on your account", nil)
		return
	}
	zone, ok := s.readyZone(w, r, key)
	if !ok {
		return
	}
	check, err := s.Store.ReadyCheckOf(r.Context(), key, zone)
	if err != nil {
		s.fail(w, r, "ready_check", err, readyFailMessage)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, check)
}

// readyZone is the ?raid= raid's zone (nil without one); ok is false once it has answered 404
// for a raid that is not one of the character's guilds'.
func (s *Service) readyZone(w http.ResponseWriter, r *http.Request, key string) (*Zone, bool) {
	q := r.URL.Query().Get("raid")
	if q == "" {
		return nil, true
	}
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil || id <= 0 {
		writeRaidNotFound(w, r)
		return nil, false
	}
	zoneKey, found, err := s.Store.characterRaidZone(r.Context(), key, id)
	if err != nil {
		s.fail(w, r, "ready_check", err, readyFailMessage)
		return nil, false
	}
	if !found {
		writeRaidNotFound(w, r)
		return nil, false
	}
	zone := zoneOf(zoneKey)
	return &zone, true
}

// isOwnCharacter reports whether key is userID's character.
func (s *Store) isOwnCharacter(ctx context.Context, userID int64, key string) (bool, error) {
	var owned bool
	if err := s.Pool.QueryRow(ctx, `select exists (select 1 from characters where key = $1 and user_id = $2)`,
		key, userID).Scan(&owned); err != nil {
		return false, fmt.Errorf("guilds: owns character: %w", err)
	}
	return owned, nil
}

// characterRaidZone is the zone of raidID when its guild has key verified on its roster.
func (s *Store) characterRaidZone(ctx context.Context, key string, raidID int64) (string, bool, error) {
	var zone string
	err := s.Pool.QueryRow(ctx, `select r.zone from guild_raids r
		join guild_characters gc on gc.guild_id = r.guild_id and gc.character_key = $1 and gc.verified_at is not null
		where r.id = $2`, key, raidID).Scan(&zone)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("guilds: character raid zone: %w", err)
	}
	return zone, true, nil
}
