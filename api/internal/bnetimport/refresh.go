// api/internal/bnetimport/refresh.go
package bnetimport

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// refreshWindow is spec §4.4's staleness threshold: a 'bnet'-sourced
// character not refreshed within this window is due for another look.
const refreshWindow = 20 * time.Hour

// refreshRateLimit bounds Blizzard calls during the nightly refresh
// (spec §4.4: "at most 4 Blizzard calls per second").
const refreshRateLimit = 4

// RefreshResult is what one refresh run produced.
type RefreshResult struct {
	Considered  int
	Refreshed   int
	Skipped     int
	RateLimited bool
}

// staleCharacter identifies one characters row the refresh must
// re-check, by Blizzard's own character id and the realm it was last
// seen on — not by its key, which can go stale between this batch's
// snapshot and the moment it is actually processed (spec A3: a
// concurrent login import can rekey a row mid-run). refreshOneCharacter
// re-resolves the current key from these two fields right before use.
type staleCharacter struct {
	BnetCharacterID int64
	RealmSlug       string
	UserID          int64
}

// staleBnetCharacters lists every 'bnet'-sourced character not refreshed
// within refreshWindow.
func (s *Service) staleBnetCharacters(ctx context.Context) ([]staleCharacter, error) {
	rows, err := s.Pool.Query(ctx,
		`select coalesce(bnet_character_id, 0), coalesce(realm_slug, ''), user_id from characters
		 where source = 'bnet' and refreshed_at < now() - $1::interval
		 order by refreshed_at`,
		fmt.Sprintf("%d seconds", int(refreshWindow.Seconds())))
	if err != nil {
		return nil, fmt.Errorf("bnetimport: stale characters: %w", err)
	}
	defer rows.Close()
	var out []staleCharacter
	for rows.Next() {
		var sc staleCharacter
		if err := rows.Scan(&sc.BnetCharacterID, &sc.RealmSlug, &sc.UserID); err != nil {
			return nil, fmt.Errorf("bnetimport: stale characters: %w", err)
		}
		out = append(out, sc)
	}
	return out, rows.Err()
}

// RunRefresh re-syncs every stale 'bnet' character's guild membership
// with the app token (no user token needed — a character profile and a
// guild roster are public data), at most refreshRateLimit Blizzard calls
// per second, stopping on the first rate limit (spec §4.4). It then runs
// the namespace probe (spec §5), logging one line per game and a WARN
// for any game other than the configured one that answers 200.
func (s *Service) RunRefresh(ctx context.Context, probeGames []string) (RefreshResult, error) {
	stale, err := s.staleBnetCharacters(ctx)
	if err != nil {
		return RefreshResult{}, err
	}
	result := RefreshResult{Considered: len(stale)}

	rosterCache := map[string]bnetapi.Roster{}
	realmCache := map[string][]bnetapi.Realm{}
	ticker := time.NewTicker(time.Second / refreshRateLimit)
	defer ticker.Stop()
	for _, sc := range stale {
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		case <-ticker.C:
		}
		if err := s.refreshOneCharacter(ctx, sc, rosterCache, realmCache); err != nil {
			if errors.Is(err, bnetapi.ErrRateLimited) {
				result.RateLimited = true
				s.logger().Warn("bnetimport", "op", "refresh", "err", "rate limited by Blizzard, stopping this run")
				break
			}
			s.logger().Warn("bnetimport", "op", "refresh_character", "bnet_character_id", sc.BnetCharacterID, "err", err)
			result.Skipped++
			continue
		}
		result.Refreshed++
	}

	s.runProbe(ctx, probeGames)
	return result, nil
}

// refreshOneCharacter re-resolves sc's *current* characters row by
// Blizzard character id and realm (spec A3 — never trusting the key as
// it stood when the batch was snapshotted, which a concurrent login
// import may have rekeyed since), bumps its refreshed_at, and re-runs
// the guild sync, equipment capture, and media capture steps of the import (spec §4.2
// steps 4-5, §B). It never re-runs step 3 (the characters row's own
// class/level/faction/realm), which needs the account's own user token
// the nightly job does not have. A row that has since been rekeyed or
// removed (no match on bnet_character_id/realm_slug) is silently skipped
// — nothing to refresh.
func (s *Service) refreshOneCharacter(ctx context.Context, sc staleCharacter, rosterCache map[string]bnetapi.Roster,
	realmCache map[string][]bnetapi.Realm) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("bnetimport: refresh begin bnet_character_id=%d: %w", sc.BnetCharacterID, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	var key, region, ruleset, name string
	err = tx.QueryRow(ctx,
		`select key, region, ruleset, name from characters
		 where bnet_character_id = $1 and coalesce(realm_slug, '') = $2 and source = 'bnet'`,
		sc.BnetCharacterID, sc.RealmSlug).Scan(&key, &region, &ruleset, &name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("bnetimport: refresh resolve bnet_character_id=%d: %w", sc.BnetCharacterID, err)
	}
	key, ruleset, err = s.rekeyInPlace(ctx, tx, key, region, ruleset, name, sc.RealmSlug, realmCache)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, `update characters set refreshed_at = now() where key = $1`, key); err != nil {
		return fmt.Errorf("bnetimport: refresh stamp %s: %w", key, err)
	}
	profile, _, _, err := s.syncCharacterGuild(ctx, tx, sc.UserID, region, ruleset, key, sc.RealmSlug, name, rosterCache)
	if err != nil {
		return err
	}
	rawEquipment, err := s.captureEquipment(ctx, tx, key, region, sc.RealmSlug, name)
	if err != nil {
		return err
	}
	rawSpecializations, err := s.captureSpecializations(ctx, tx, key, region, sc.RealmSlug, name)
	if err != nil {
		return err
	}
	if err := s.captureMedia(ctx, tx, key, region, sc.RealmSlug, name); err != nil {
		return err
	}
	if err := s.buildAndWriteExport(ctx, tx, sc.UserID, key, region, ruleset, profile, rawEquipment, rawSpecializations); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("bnetimport: refresh commit %s: %w", key, err)
	}
	return nil
}

// runProbe requests the namespace probe for every configured region and
// logs the result; it never fails the refresh run.
func (s *Service) runProbe(ctx context.Context, games []string) {
	if len(games) == 0 {
		return
	}
	for _, region := range s.Regions {
		for _, pr := range s.Client.Probe(ctx, region, games) {
			s.logger().Info("bnetimport", "op", "probe", "region", region, "game", pr.Game, "status", pr.Status)
			if pr.Game != s.Client.Game && pr.Status == http.StatusOK {
				s.logger().Warn("bnetimport", "op", "namespace_appeared", "region", region, "game", pr.Game)
			}
		}
	}
}

// rekeyInPlace re-resolves the row's ruleset from its realm (the login
// import keyed the first production rows under a realm list whose types
// had not loaded, see the 2026-09-22 brief, A1) and, when the key it
// implies differs, moves the characters and guild_characters rows to the
// new key inside the refresh's transaction. A realm list that cannot be
// fetched leaves the key alone for this run; a new key already taken by
// another row (the same name on two realms sharing a ruleset) is logged
// and left alone too. Returns the key and ruleset to carry on with.
func (s *Service) rekeyInPlace(ctx context.Context, tx pgx.Tx, key, region, ruleset, name, realmSlug string,
	realmCache map[string][]bnetapi.Realm) (string, string, error) {
	realms, ok := realmCache[region]
	if !ok {
		var err error
		if realms, err = s.Client.Realms(ctx, region); err != nil {
			s.logger().Warn("bnetimport", "op", "realms", "region", region, "err", err)
			realms = nil
		}
		realmCache[region] = realms
	}
	if realms == nil {
		return key, ruleset, nil
	}
	newRuleset := rulesetForRealm(realms, realmSlug)
	newKey := character.Key(region, newRuleset, name)
	if newKey == key {
		return key, ruleset, nil
	}
	var taken bool
	if err := tx.QueryRow(ctx, `select exists (select 1 from characters where key = $1)`, newKey).Scan(&taken); err != nil {
		return "", "", fmt.Errorf("bnetimport: rekey check %s: %w", newKey, err)
	}
	if taken {
		s.logger().Warn("bnetimport", "op", "rekey_taken", "old_key", key, "new_key", newKey)
		return key, ruleset, nil
	}
	if _, err := tx.Exec(ctx, `update characters set key = $1, ruleset = $2 where key = $3`, newKey, newRuleset, key); err != nil {
		return "", "", fmt.Errorf("bnetimport: rekey %s: %w", key, err)
	}
	if _, err := tx.Exec(ctx, `update guild_characters set character_key = $1 where character_key = $2`, newKey, key); err != nil {
		return "", "", fmt.Errorf("bnetimport: rekey membership %s: %w", key, err)
	}
	// The account's chosen main follows its character.
	if _, err := tx.Exec(ctx, `update users set main_character_key = $1 where main_character_key = $2`, newKey, key); err != nil {
		return "", "", fmt.Errorf("bnetimport: rekey main %s: %w", key, err)
	}
	s.logger().Info("bnetimport", "op", "rekey", "old_key", key, "new_key", newKey)
	return newKey, newRuleset, nil
}
