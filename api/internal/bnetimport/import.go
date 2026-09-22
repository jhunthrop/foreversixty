// api/internal/bnetimport/import.go

// Package bnetimport is the Battle.net character/guild import (spec
// docs/superpowers/specs/2026-09-22-battlenet-character-import-design.md
// §4): the at-login import (ImportAccount, satisfying auth.Importer) and
// the nightly refresh job (RefreshJobCommand). It writes directly to
// characters, guild_characters and guilds — never through auth.Store or
// addon.Store — and it must never import package addon (the two write
// the same tables through different entry points; addon already imports
// guilds, and this package would create a cycle the other way if it
// reached back into addon).
package bnetimport

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// RefreshJobCommand is this lane's Cloud Run job name, dispatched the
// same way dataaddon.JobCommand already is (api/cmd/api/main.go's
// os.Args[1] switch).
const RefreshJobCommand = "bnet-refresh"

// ErrBudget is returned (with the partial ImportSummary already applied)
// when ImportAccount's context deadline is reached mid-run. What was
// written stays written; the nightly refresh picks up the rest (spec
// §4.1).
var ErrBudget = errors.New("bnetimport: import budget exceeded")

// minImportLevel is spec §4.2 step 2's RULING: characters below this
// level are not imported — bank alts and throwaways clutter the list,
// and 10 is the level Forever's own tools start to matter. Cost if
// wrong: a player wonders where their level 8 alt is; they can still
// paste it by hand.
const minImportLevel = 10

// Service runs the Battle.net import and the nightly refresh.
// ImportAccount satisfies auth.Importer (it returns auth.ImportSummary,
// defined in that package rather than here so the dependency runs one
// way only: bnetimport already needs auth transitively through guilds,
// which reads auth.User, so auth itself never imports bnetimport).
type Service struct {
	Pool    *pgxpool.Pool
	Client  *bnetapi.Client
	Regions []string
	Log     *slog.Logger
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

// ImportAccount reads every character on the account behind userToken,
// across every configured region, and writes what it finds — the
// characters row, and (for a guilded character) its guild membership —
// stopping early only when ctx's deadline is reached. It always stamps
// users.bnet_imported_at, even on a partial run: a login that started an
// import always counts as "we tried," and GET /v1/me's "Imported from
// Battle.net" line and the nightly refresh both key off that timestamp.
func (s *Service) ImportAccount(ctx context.Context, userID int64, userToken string) (auth.ImportSummary, error) {
	summary := auth.ImportSummary{}
	rosterCache := map[string]bnetapi.Roster{}
	realmCache := map[string][]bnetapi.Realm{}

	var runErr error
	for _, region := range s.Regions {
		if ctx.Err() != nil {
			runErr = ErrBudget
			break
		}
		if err := s.importRegion(ctx, userID, userToken, region, &summary, rosterCache, realmCache); err != nil {
			runErr = err
			break
		}
	}

	if _, err := s.Pool.Exec(context.Background(),
		`update users set bnet_imported_at = now() where id = $1`, userID); err != nil {
		s.logger().Warn("bnetimport", "op", "stamp_imported_at", "user_id", userID, "err", err)
	}
	return summary, runErr
}

// importRegion imports every character the account has in one region. A
// soft failure (no account in this region, a single character's profile
// unreadable) is logged and does not stop the run; only ErrBudget
// (the context deadline) propagates.
func (s *Service) importRegion(ctx context.Context, userID int64, userToken, region string,
	summary *auth.ImportSummary, rosterCache map[string]bnetapi.Roster, realmCache map[string][]bnetapi.Realm) error {
	chars, err := s.Client.AccountCharacters(ctx, region, userToken)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return ErrBudget
		}
		if errors.Is(err, bnetapi.ErrForbidden) || errors.Is(err, bnetapi.ErrNotFound) {
			return nil // no WoW account under this namespace in this region
		}
		s.logger().Warn("bnetimport", "op", "account_characters", "region", region, "user_id", userID, "err", err)
		return nil
	}
	summary.Regions = append(summary.Regions, region)

	realms, ok := realmCache[region]
	if !ok {
		if realms, err = s.Client.Realms(ctx, region); err != nil {
			s.logger().Warn("bnetimport", "op", "realms", "region", region, "user_id", userID, "err", err)
			realms = nil
		}
		realmCache[region] = realms
	}

	for _, ch := range chars {
		if ctx.Err() != nil {
			return ErrBudget
		}
		if ch.Level < minImportLevel {
			summary.Skipped++
			continue
		}
		ruleset := rulesetForRealm(realms, ch.RealmSlug)
		key := character.Key(region, ruleset, ch.Name)

		wrote, guildWritten, err := s.importOneCharacter(ctx, userID, region, ruleset, key, ch, rosterCache)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				return ErrBudget
			}
			s.logger().Warn("bnetimport", "op", "import_character", "key", key, "user_id", userID, "err", err)
			continue
		}
		if !wrote {
			summary.Skipped++
			s.logger().Warn("bnetimport", "op", "collision", "key", key, "user_id", userID)
			continue
		}
		summary.Characters++
		if guildWritten {
			summary.Guilds++
		}
	}
	return nil
}
