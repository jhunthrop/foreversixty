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
	"sort"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/bnetbuild"
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

// enrichFreshness is how recently a row must have been enriched for a run
// to skip it: long enough to cover the sign-in run and its background
// continuation, far shorter than the nightly refreshWindow.
const enrichFreshness = 10 * time.Minute

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
	// Tables is the active client build's talent/enchant/suffix/race data (spec §2.2,
	// §2.3), loaded once per process by cmd/api's wiring. A zero Tables (Trees nil) means
	// no build data is available and buildAndWriteExport (build.go) is a no-op — never a
	// crash.
	Tables bnetbuild.Tables
	// MaxConcurrentEnrichments caps how many characters this process enriches at
	// once. Each enrichment holds one pool connection inside a transaction while it
	// waits on several Battle.net calls, so a sign-in burst without a cap would pin
	// every connection and stall every other request on the instance. Zero means no
	// cap (the nightly refresh, which enriches one character at a time). The API sets
	// it with EnrichmentSlotsFor.
	MaxConcurrentEnrichments int

	// staticSeen remembers which static namespaces noteStaticNamespace has
	// already logged.
	staticSeen sync.Map

	enrichSlotsOnce sync.Once
	enrichSlots     chan struct{}
}

// EnrichmentSlotsFor is the sign-in enrichment cap for a process whose pool is
// pool: half its connections, at least one, so the request path always keeps the
// other half.
func EnrichmentSlotsFor(pool *pgxpool.Pool) int {
	return max(1, int(pool.Config().MaxConns)/2)
}

// acquireEnrichSlot waits for one of MaxConcurrentEnrichments slots, or for ctx to
// end. The returned release gives the slot back.
func (s *Service) acquireEnrichSlot(ctx context.Context) (release func(), err error) {
	if s.MaxConcurrentEnrichments <= 0 {
		return func() {}, nil
	}
	s.enrichSlotsOnce.Do(func() { s.enrichSlots = make(chan struct{}, s.MaxConcurrentEnrichments) })
	select {
	case s.enrichSlots <- struct{}{}:
		return func() { <-s.enrichSlots }, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
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

// importRegion imports every character the account has in one region, in
// two passes so a first sign-in lists every character at once:
//
//  1. every eligible character's row is written straight from the account
//     profile, which is one call for the whole account and cheap DB work;
//  2. each row is then enriched (public profile, guild, gear, talents,
//     media, the Blizzard-sourced build), highest level first, so the
//     character the site will most likely pick as the main is ready first.
//
// A soft failure (no account in this region, a single character's profile
// unreadable) is logged and does not stop the run; only ErrBudget (the
// context deadline) propagates, with summary.Pending counting the rows the
// second pass did not reach.
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

	written, err := s.writeAccountRows(ctx, userID, region, realms, byLevel(chars), summary)
	summary.Characters += len(written)
	if err != nil {
		summary.Pending += len(written)
		return err
	}
	return s.enrichAll(ctx, written, summary, rosterCache)
}

// byLevel is every account character, highest level first (stable, so
// Blizzard's own order breaks ties). Every Forever character is imported,
// whatever its level: the old level-10 floor (spec §4.2 step 2) skipped every
// character on launch night, when all of them were brand new (2026-11-04,
// user 78: five characters, all skipped).
func byLevel(chars []bnetapi.AccountCharacter) []bnetapi.AccountCharacter {
	out := append([]bnetapi.AccountCharacter(nil), chars...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Level > out[j].Level })
	return out
}

// writeAccountRows is the first pass: one characters row per character,
// keyed by the ruleset its realm resolves to. A collision (the key already
// belongs to another account) is counted in summary.Skipped and logged.
func (s *Service) writeAccountRows(ctx context.Context, userID int64, region string, realms []bnetapi.Realm,
	chars []bnetapi.AccountCharacter, summary *auth.ImportSummary) ([]characterRef, error) {
	written := make([]characterRef, 0, len(chars))
	for _, ch := range chars {
		if ctx.Err() != nil {
			return written, ErrBudget
		}
		ref, wrote, err := s.writeAccountRow(ctx, userID, region, realms, ch)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) || ctx.Err() != nil {
				return written, ErrBudget
			}
			s.logger().Warn("bnetimport", "op", "write_character", "key", ref.Key, "user_id", userID, "err", err)
			continue
		}
		if !wrote {
			summary.Skipped++
			s.logger().Warn("bnetimport", "op", "collision", "key", ref.Key, "user_id", userID)
			continue
		}
		written = append(written, ref)
	}
	return written, nil
}

// writeAccountRow resolves one account character's ruleset and key and
// writes its row.
func (s *Service) writeAccountRow(ctx context.Context, userID int64, region string, realms []bnetapi.Realm,
	ch bnetapi.AccountCharacter) (characterRef, bool, error) {
	ruleset, err := s.rulesetFor(ctx, userID, region, realms, ch)
	if err != nil {
		return characterRef{Key: ch.Name}, false, err
	}
	ref := characterRef{
		UserID: userID, Region: region, Ruleset: ruleset, Key: character.Key(region, ruleset, ch.Name),
		RealmSlug: ch.RealmSlug, Name: ch.Name, BnetID: ch.ID,
	}
	wrote, err := s.writeCharacterRow(ctx, ref, ch)
	return ref, wrote, err
}

// enrichAll is the second pass over the rows writeAccountRows wrote. A row
// enriched within enrichFreshness is skipped, so a run that continues an
// earlier one (auth's background continuation after the sign-in budget)
// spends its calls only on what the first run did not reach.
func (s *Service) enrichAll(ctx context.Context, refs []characterRef, summary *auth.ImportSummary,
	rosterCache map[string]bnetapi.Roster) error {
	for i, ref := range refs {
		if ctx.Err() != nil {
			summary.Pending += len(refs) - i
			return ErrBudget
		}
		fresh, err := s.recentlyEnriched(ctx, ref.Key)
		if err != nil {
			s.logger().Warn("bnetimport", "op", "enrich_freshness", "key", ref.Key, "err", err)
		}
		if fresh {
			continue
		}
		release, err := s.acquireEnrichSlot(ctx)
		if err != nil {
			summary.Pending += len(refs) - i
			return ErrBudget
		}
		guildWritten, unavailable, err := s.enrichCharacter(ctx, ref, rosterCache)
		release()
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) || ctx.Err() != nil {
				summary.Pending += len(refs) - i
				return ErrBudget
			}
			// A 429 stops the batch (bnetapi.ErrRateLimited): every further call
			// would spend the shared client quota for another refusal. The rows stay
			// listed and the nightly refresh enriches them.
			if errors.Is(err, bnetapi.ErrRateLimited) {
				summary.Pending += len(refs) - i
				summary.RateLimited = true
				return err
			}
			s.logger().Warn("bnetimport", "op", "enrich_character", "key", ref.Key, "user_id", ref.UserID, "err", err)
			continue
		}
		if guildWritten {
			summary.Guilds++
		}
		if unavailable {
			summary.Unavailable++
		}
	}
	return nil
}
