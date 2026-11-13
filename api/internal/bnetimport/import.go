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

// ErrTokenRejected is ImportAccount's answer when Battle.net refused the user token (401):
// it expired early or the player revoked it. Nothing was written for that region.
var ErrTokenRejected = errors.New("bnetimport: battle.net refused the user token")

// enrichFreshness is how recently a row must have been enriched for a run
// to skip it: long enough to cover the sign-in run and its background
// continuation, far shorter than activeRefreshWindow.
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
	// ImportSims queues the auto-sim of a character whose gear the enrichment changed; nil
	// queues none.
	ImportSims ImportSimmer

	// staticSeen remembers which static namespaces noteStaticNamespace has
	// already logged.
	staticSeen sync.Map
	// unlistedRealms remembers which region/realm ids noteUnlistedRealm has
	// already logged.
	unlistedRealms sync.Map

	enrichSlotsOnce sync.Once
	enrichSlots     chan struct{}

	// Sealer seals the sign-in's access token at rest (KeepToken); nil keeps none, and every
	// visit refresh then reads public profiles.
	Sealer TokenSealer
	// backoff pauses automatic refreshes after a 429 (backoff.go).
	backoff rateBackoff
	// visiting holds the user ids whose visit refresh is running in this process.
	visiting sync.Map
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

	if summary.RateLimited {
		s.backoff.hit(time.Now())
	}
	s.stampAccountSync(userID)
	return summary, runErr
}

// stampAccountSync records an account import's end: bnet_imported_at and bnet_synced_at,
// and, when the account has no main yet, its highest-level character as the main, so the
// site opens on it after the first import. A failed write is logged, never the import's.
func (s *Service) stampAccountSync(userID int64) {
	ctx := context.Background()
	if _, err := s.Pool.Exec(ctx,
		`update users set bnet_imported_at = now(), bnet_synced_at = now() where id = $1`, userID); err != nil {
		s.logger().Warn("bnetimport", "op", "stamp_imported_at", "user_id", userID, "err", err)
	}
	if _, err := s.Pool.Exec(ctx,
		`update users u set main_character_key = c.key
		 from (select key from characters where user_id = $1 order by level desc nulls last, key limit 1) c
		 where u.id = $1 and u.main_character_key is null`, userID); err != nil {
		s.logger().Warn("bnetimport", "op", "pick_main", "user_id", userID, "err", err)
	}
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
// context deadline) and ErrTokenRejected (a 401) propagate, with summary.Pending counting the rows the
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
		if errors.Is(err, bnetapi.ErrUnauthorized) {
			return ErrTokenRejected
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

// opNeedsRuleset is the log op a character no source names a ruleset for is
// reported under; the bnet_needs_ruleset log-based metric counts it. Such a
// character is imported under character.RulesetUnknown and its player asked.
const opNeedsRuleset = "needs_ruleset"

// writeAccountRows is the first pass: one characters row per character,
// keyed by the ruleset rulesetFor resolves, or by character.RulesetUnknown
// when nothing names one. A collision (the key already belongs to another
// account) is counted in summary.Skipped and logged.
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
// writes its row. A character no source names a ruleset for is keyed under
// character.RulesetUnknown, never a guess, and logged with its realm id so
// the Forever realm table can learn it.
func (s *Service) writeAccountRow(ctx context.Context, userID int64, region string, realms []bnetapi.Realm,
	ch bnetapi.AccountCharacter) (characterRef, bool, error) {
	resolved, ok, err := s.rulesetFor(ctx, userID, region, realms, ch)
	if err != nil {
		return characterRef{Key: ch.Name}, false, err
	}
	if !ok {
		// Counted by the bnet_needs_ruleset log-based metric (docs/runbooks/launch-day.md).
		_, realmKnown := s.Client.KnownRealm(region, ch.RealmID)
		s.logger().Warn("bnetimport", "op", opNeedsRuleset, "user_id", userID, "region", region,
			"bnet_character_id", ch.ID, "realm_id", ch.RealmID, "realm_known", realmKnown,
			"fields_checked", rulesetSources, "fallback", "ask")
		resolved = resolvedRuleset{Ruleset: character.RulesetUnknown}
	}
	ref := characterRef{
		UserID: userID, Region: region, Ruleset: resolved.Ruleset, Key: character.Key(region, resolved.Ruleset, ch.Name),
		RealmSlug: ch.RealmSlug, Name: ch.Name, BnetID: ch.ID, RealmID: ch.RealmID, Answered: resolved.Answered,
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
