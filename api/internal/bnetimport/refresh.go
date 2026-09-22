// api/internal/bnetimport/refresh.go
package bnetimport

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
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

// staleCharacter is one characters row the refresh must re-check.
type staleCharacter struct {
	Key, Region, Ruleset, RealmSlug, Name string
	UserID                                int64
}

// staleBnetCharacters lists every 'bnet'-sourced character not refreshed
// within refreshWindow.
func (s *Service) staleBnetCharacters(ctx context.Context) ([]staleCharacter, error) {
	rows, err := s.Pool.Query(ctx,
		`select key, region, ruleset, coalesce(realm_slug, ''), name, user_id from characters
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
		if err := rows.Scan(&sc.Key, &sc.Region, &sc.Ruleset, &sc.RealmSlug, &sc.Name, &sc.UserID); err != nil {
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
	ticker := time.NewTicker(time.Second / refreshRateLimit)
	defer ticker.Stop()
	for _, sc := range stale {
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		case <-ticker.C:
		}
		if err := s.refreshOneCharacter(ctx, sc, rosterCache); err != nil {
			if errors.Is(err, bnetapi.ErrRateLimited) {
				result.RateLimited = true
				s.logger().Warn("bnetimport", "op", "refresh", "err", "rate limited by Blizzard, stopping this run")
				break
			}
			s.logger().Warn("bnetimport", "op", "refresh_character", "key", sc.Key, "err", err)
			result.Skipped++
			continue
		}
		result.Refreshed++
	}

	s.runProbe(ctx, probeGames)
	return result, nil
}

// refreshOneCharacter bumps a stale row's refreshed_at and re-runs the
// guild sync steps of the import (spec §4.2 steps 4-5) — it never
// re-runs step 3 (the characters row's own class/level/faction/realm),
// which needs the account's own user token the nightly job does not
// have.
func (s *Service) refreshOneCharacter(ctx context.Context, sc staleCharacter, rosterCache map[string]bnetapi.Roster) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("bnetimport: refresh begin %s: %w", sc.Key, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if _, err := tx.Exec(ctx, `update characters set refreshed_at = now() where key = $1`, sc.Key); err != nil {
		return fmt.Errorf("bnetimport: refresh stamp %s: %w", sc.Key, err)
	}
	if _, err := s.syncCharacterGuild(ctx, tx, sc.UserID, sc.Region, sc.Ruleset, sc.Key, sc.RealmSlug, sc.Name, rosterCache); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("bnetimport: refresh commit %s: %w", sc.Key, err)
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
