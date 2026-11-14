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
	"github.com/jhunthrop/foreversixty/api/internal/synclog"
)

// The refresh windows (spec §4.4, widened 2026-11-07): a 'bnet'-sourced character is due
// for another look once it has gone this long unrefreshed. An account active within
// activeAccountWithin (users.bnet_synced_at, which every sign-in and every visit refresh
// stamps) gets the short window; every other account keeps the nightly one.
const (
	activeRefreshWindow = 3 * time.Hour
	idleRefreshWindow   = 20 * time.Hour
	activeAccountWithin = 14 * 24 * time.Hour
)

// refreshRateLimit bounds Blizzard calls during a refresh batch
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

// refreshWindowFor is how long a character of an account whose last sync started at
// lastSync may go unrefreshed: activeRefreshWindow inside activeAccountWithin, else
// idleRefreshWindow. staleBnetCharacters' SQL is this rule; the two are tested together.
func refreshWindowFor(lastSync *time.Time, now time.Time) time.Duration {
	if lastSync != nil && now.Sub(*lastSync) < activeAccountWithin {
		return activeRefreshWindow
	}
	return idleRefreshWindow
}

// intervalArg renders d as a Postgres interval argument.
func intervalArg(d time.Duration) string {
	return fmt.Sprintf("%d seconds", int(d.Seconds()))
}

// staleBnetCharacters lists every 'bnet'-sourced character due under refreshWindowFor,
// oldest refresh first.
func (s *Service) staleBnetCharacters(ctx context.Context) ([]staleCharacter, error) {
	return s.queryCharacters(ctx,
		`select coalesce(c.bnet_character_id, 0), coalesce(c.realm_slug, ''), c.user_id
		 from characters c left join users u on u.id = c.user_id
		 where c.source = 'bnet' and c.user_id is not null and c.refreshed_at < now() -
		   (case when u.bnet_synced_at > now() - $3::interval then $1::interval else $2::interval end)
		 order by c.refreshed_at`,
		intervalArg(activeRefreshWindow), intervalArg(idleRefreshWindow), intervalArg(activeAccountWithin))
}

// accountBnetCharacters lists one account's 'bnet'-sourced characters, highest level
// first: the visit refresh's public-profile fallback.
func (s *Service) accountBnetCharacters(ctx context.Context, userID int64) ([]staleCharacter, error) {
	return s.queryCharacters(ctx,
		`select coalesce(bnet_character_id, 0), coalesce(realm_slug, ''), user_id from characters
		 where source = 'bnet' and user_id = $1
		 order by level desc nulls last, key`, userID)
}

func (s *Service) queryCharacters(ctx context.Context, sql string, args ...any) ([]staleCharacter, error) {
	rows, err := s.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("bnetimport: list characters: %w", err)
	}
	defer rows.Close()
	var out []staleCharacter
	for rows.Next() {
		var sc staleCharacter
		if err := rows.Scan(&sc.BnetCharacterID, &sc.RealmSlug, &sc.UserID); err != nil {
			return nil, fmt.Errorf("bnetimport: list characters: %w", err)
		}
		out = append(out, sc)
	}
	return out, rows.Err()
}

// RunRefresh is the bnet-refresh job: every due 'bnet' character (staleBnetCharacters) is
// re-synced through refreshBatch, then the namespace probe runs (spec §5), logging one line
// per game and a WARN for any game other than the configured one that answers 200.
func (s *Service) RunRefresh(ctx context.Context, probeGames []string) (RefreshResult, error) {
	result, err := s.RefreshDue(ctx)
	if err != nil {
		return result, err
	}
	s.runProbe(ctx, probeGames)
	return result, nil
}

// RefreshDue re-syncs every due 'bnet' character: the job's work and each pass of the
// API's in-process RefreshLoop.
func (s *Service) RefreshDue(ctx context.Context) (RefreshResult, error) {
	due, err := s.staleBnetCharacters(ctx)
	if err != nil {
		return RefreshResult{}, err
	}
	return s.refreshBatch(ctx, due)
}

// refreshBatch re-syncs each character with the app token (no user token needed: a
// character profile and a guild roster are public data), at most refreshRateLimit
// characters per second. A 429 stops the batch and backs every later batch in this process
// off (rateBackoff); a batch that starts inside a backoff refreshes nothing and reports
// RateLimited.
func (s *Service) refreshBatch(ctx context.Context, chars []staleCharacter) (RefreshResult, error) {
	result := RefreshResult{Considered: len(chars)}
	if len(chars) == 0 {
		return result, nil
	}
	if s.backoff.blocked(time.Now()) {
		result.RateLimited = true
		return result, nil
	}
	rosterCache := map[string]bnetapi.Roster{}
	realmCache := map[string][]bnetapi.Realm{}
	ticker := time.NewTicker(time.Second / refreshRateLimit)
	defer ticker.Stop()
	for _, sc := range chars {
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		case <-ticker.C:
		}
		if err := s.refreshOneCharacter(ctx, sc, rosterCache, realmCache); err != nil {
			if errors.Is(err, bnetapi.ErrRateLimited) {
				result.RateLimited = true
				wait := s.backoff.hit(time.Now())
				s.logger().Warn("bnetimport", "op", "refresh", "err", "rate limited by Blizzard, stopping this run",
					"backoff", wait.String())
				return result, nil
			}
			s.logger().Warn("bnetimport", "op", "refresh_character", "bnet_character_id", sc.BnetCharacterID, "err", err)
			s.recordRefreshFailure(ctx, sc)
			result.Skipped++
			continue
		}
		result.Refreshed++
	}
	s.backoff.reset()
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
	var realmID int64
	err = tx.QueryRow(ctx,
		`select key, region, ruleset, name, coalesce(bnet_realm_id, 0) from characters
		 where bnet_character_id = $1 and coalesce(realm_slug, '') = $2 and source = 'bnet'`,
		sc.BnetCharacterID, sc.RealmSlug).Scan(&key, &region, &ruleset, &name, &realmID)
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
	key, ruleset, err = s.resolveUnknownInPlace(ctx, tx, key, region, ruleset, name, realmID, sc.BnetCharacterID)
	if err != nil {
		return err
	}

	ref := characterRef{
		UserID: sc.UserID, Region: region, Ruleset: ruleset, Key: key,
		RealmSlug: sc.RealmSlug, Name: name, BnetID: sc.BnetCharacterID, RealmID: realmID,
	}
	var guildWritten, unavailable bool
	if err := s.enrichInTx(ctx, tx, &ref, rosterCache, &guildWritten, &unavailable); err != nil {
		return err
	}
	// The account's "Updated from Battle.net" stamp (GET /v1/me's bnet_imported_at).
	if _, err := tx.Exec(ctx, `update users set bnet_imported_at = now() where id = $1`, sc.UserID); err != nil {
		return fmt.Errorf("bnetimport: refresh stamp account %s: %w", key, err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("bnetimport: refresh commit %s: %w", key, err)
	}
	s.simImport(ctx, ref, false)
	return nil
}

// recordRefreshFailure notes in the sync history that sc's refresh failed, for GET /v1/me's
// build.sync_error. It runs outside the refresh's transaction (which the failure rolled back)
// and looks the key up afresh, since the rolled-back transaction may have been the one
// re-keying it. A history that cannot be written is logged, never a second failure.
func (s *Service) recordRefreshFailure(ctx context.Context, sc staleCharacter) {
	var key string
	err := s.Pool.QueryRow(ctx,
		`select key from characters where bnet_character_id = $1 and coalesce(realm_slug, '') = $2 and source = 'bnet'`,
		sc.BnetCharacterID, sc.RealmSlug).Scan(&key)
	if errors.Is(err, pgx.ErrNoRows) {
		return
	}
	if err == nil {
		err = synclog.Record(ctx, s.Pool, key, synclog.SourceBlizzard, synclog.OutcomeRefreshFailed)
	}
	if err != nil {
		s.logger().Warn("bnetimport", "op", "record_refresh_failure", "bnet_character_id", sc.BnetCharacterID, "err", err)
	}
}

// runProbe requests the namespace probe for every configured region and
// logs the result; it never fails the refresh run.
// Probe runs the namespace probe alone: the nightly job's path while the character import
// is off (config.BnetImportCharacters), so the namespace_appeared warning still fires the
// day Blizzard serves Forever, with no character re-synced from a namespace that is not it.
func (s *Service) Probe(ctx context.Context, games []string) {
	s.runProbe(ctx, games)
}

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
	// A realmless (Forever) row's ruleset is the profile's own, which the
	// enrich pass adopts (adoptProfileRuleset); the realm step leaves it.
	if realms == nil || realmSlug == "" {
		return key, ruleset, nil
	}
	newRuleset, ok := rulesetForRealm(realms, realmSlug)
	newKey := character.Key(region, newRuleset, name)
	if !ok || newKey == key {
		return key, ruleset, nil
	}
	_, moved, err := s.moveKey(ctx, tx, key, newKey, newRuleset)
	if err != nil || !moved {
		return key, ruleset, err
	}
	return newKey, newRuleset, nil
}
