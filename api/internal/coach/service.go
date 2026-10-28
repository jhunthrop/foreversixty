package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// ReviewJobCommand is the job that writes one review: coach-review <report> <fight>
// <player_key> <account_id> <lease holder>.
const ReviewJobCommand = "coach-review"

// Service serves the coach routes and runs the coach-review job.
type Service struct {
	Store        *Store
	Reports      *reports.Store
	Accounts     reports.GuildRanker
	Owners       Owners
	Entitlements Entitlementer
	Summaries    Getter
	Tables       TableSource
	Refs         *Resolver
	Personas     *Personas
	// Generator is the template the job copies: one Generator per job, whose Validator
	// carries Names and that fight's roster.
	Generator *Generator
	// Names is the process-wide name dictionary (LoadNames), built once at construction.
	Names *NameDict
	Caps  *Caps
	// Jobs is bound to the coach-review job only (one CloudRun per job, F23); nil: nothing
	// is queued and the rule-only rendering is served.
	Jobs          jobs.Runner
	Enabled       bool
	PublicBaseURL string
	Log           *slog.Logger
	Now           func() time.Time
	// ReviewDeadline bounds one coach-review job from its start; ReviewDeadline (the
	// constant) when zero. Tests shorten it.
	ReviewDeadline time.Duration
}

// Owners answers whether an account owns a character. auth.Store satisfies it.
type Owners interface {
	OwnsCharacter(ctx context.Context, userID int64, key string) (bool, error)
}

// Entitlementer answers feature questions. entitlements.Store satisfies it.
type Entitlementer interface {
	Can(ctx context.Context, userID int64, f entitlements.Feature) (bool, entitlements.Reason, error)
	CanForGuild(ctx context.Context, guildID, userID int64, f entitlements.Feature) (bool, entitlements.Reason, error)
}

// Getter reads one object out of the bucket. *r2.Client satisfies it.
type Getter interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
}

// Errors the fight route answers 404 for.
var (
	ErrNoFight  = errors.New("coach: no such fight")
	ErrNoPlayer = errors.New("coach: no such player in the fight")
)

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

// loaded is a pack with the tables its version was taken from (the DTO reads the aura
// kinds from the rule file).
type loaded struct {
	pack   Pack
	tables Tables
}

// PackFor loads the cached pack for the current pack version, or builds it from the bucket
// summary and caches it. A preparing pack is not cached.
func (s *Service) PackFor(ctx context.Context, rep reports.Report, fightIndex int, playerKey string) (Pack, error) {
	l, err := s.loadPack(ctx, rep, fightIndex, playerKey)
	return l.pack, err
}

func (s *Service) loadPack(ctx context.Context, rep reports.Report, fightIndex int, playerKey string) (loaded, error) {
	entry, err := s.fight(ctx, rep.ID, fightIndex)
	if err != nil {
		return loaded{}, err
	}
	if l, ok, err := s.cachedPack(ctx, rep.ID, entry, playerKey); err != nil || ok {
		return l, err
	}
	sum, err := s.summary(ctx, rep.ID, fightIndex)
	if err != nil {
		return loaded{}, err
	}
	return s.buildPack(ctx, rep, entry, sum, playerKey)
}

// cachedPack is the newest cached pack when its version is still the current one.
func (s *Service) cachedPack(ctx context.Context, reportID string, entry reports.FightEntry, playerKey string) (loaded, bool, error) {
	p, ok, err := s.Store.latestPack(ctx, reportID, entry.Index, playerKey)
	if err != nil || !ok {
		return loaded{}, false, err
	}
	row := summary.RosterRow{Class: p.Context.Class, Spec: p.Context.SpecLabel, Role: p.Context.Role, Level: int64(p.Context.Level)}
	tables, version, err := s.versionedTables(ctx, row, entry.EncounterID, entry.Start)
	if err != nil || (p.Version != version && p.Version != version+poolStateSep+poolOK) {
		return loaded{}, false, err
	}
	return loaded{pack: p, tables: tables}, true, nil
}

// buildPack reviews the player's row of sum and caches the pack unless it is preparing or
// its comparison pool is below the floor ("not enough logs yet"): that pack is served but
// rebuilt on the next view, so it picks up the top-logs basis once the pool fills.
func (s *Service) buildPack(ctx context.Context, rep reports.Report, entry reports.FightEntry, sum summary.Summary, playerKey string) (loaded, error) {
	row, ok := rosterRow(rep, sum, playerKey)
	if !ok {
		return loaded{}, ErrNoPlayer
	}
	tables, version, err := s.versionedTables(ctx, row, sum.EncounterID, entry.Start)
	if err != nil {
		return loaded{}, err
	}
	in := Input{Fight: sum, Row: row, Size: int(entry.Size), BossName: entry.Name, Tables: tables}
	if notReviewable(in) == "" {
		if in.Refs, err = s.refs(ctx, in, playerKey, entry.Start); err != nil {
			return loaded{}, err
		}
	}
	p := Review(in)
	p.Version = withPoolState(version, in)
	if p.Status != StatusPreparing && !thinPool(in) {
		if err := s.Store.SavePack(ctx, rep.ID, entry.Index, playerKey, p); err != nil {
			return loaded{}, err
		}
	}
	return loaded{pack: p, tables: tables}, nil
}

// Pool-state tokens on a reviewable pack's version: the version (and so the review key)
// moves exactly once, when the comparison pool reaches PoolMinPlayers, so a review written
// against "not enough logs yet" is regenerated once with the top-logs basis.
const (
	poolStateSep = "."
	poolThin     = "pool:thin"
	poolOK       = "pool:ok"
)

// withPoolState is version with the pool-state token of a reviewable pull; a pull that is
// not reviewable has no pool and keeps version as it is.
func withPoolState(version string, in Input) string {
	switch {
	case notReviewable(in) != "":
		return version
	case thinPool(in):
		return version + poolStateSep + poolThin
	}
	return version + poolStateSep + poolOK
}

// thinPool reports a reviewable pull whose comparison pool is below PoolMinPlayers.
func thinPool(in Input) bool {
	return notReviewable(in) == "" && in.Refs.Pool == nil
}

func (s *Service) refs(ctx context.Context, in Input, playerKey string, fought time.Time) (Refs, error) {
	k := PoolKey{EncounterID: in.Fight.EncounterID, Difficulty: in.Fight.Difficulty, Size: in.Size,
		Spec: in.Row.Spec, Phase: phase.At(fought)}
	return s.Refs.Refs(ctx, k, playerKey, in.Row.Role, int(in.Row.Level), in.Fight.DurationMS, fought)
}

// versionedTables loads the row's tables and the pack version they give. A row the tables
// do not cover (no spec, no rule data) has the zero tables, and a version of them.
func (s *Service) versionedTables(ctx context.Context, row summary.RosterRow, encounterID int64, fought time.Time) (Tables, string, error) {
	tables, err := s.Tables.TablesFor(ctx, row, encounterID, fought)
	if err != nil && !errors.Is(err, ErrNoTables) {
		return Tables{}, "", fmt.Errorf("coach: tables: %w", err)
	}
	if err != nil {
		s.warnNoTables(row, encounterID, err)
		tables = Tables{}
	}
	rules, err := json.Marshal(tables.Rules)
	if err != nil {
		return Tables{}, "", fmt.Errorf("coach: encode rules: %w", err)
	}
	cat, err := catalogueBytes()
	if err != nil {
		return Tables{}, "", err
	}
	return tables, PackVersion(rules, cat), nil
}

// warnNoTables tells the operator that a row with a class and spec has no data tables (a
// build file missing): its rule-file checks go to Couldn't check. A row with no class or
// spec is not a data defect and stays quiet. The reason never enters the pack.
func (s *Service) warnNoTables(row summary.RosterRow, encounterID int64, err error) {
	if row.Class == "" || row.Spec == "" {
		return
	}
	s.logger().Warn("coach: no tables; the rule file's checks go to Couldn't check",
		slog.String("class", row.Class), slog.String("spec", row.Spec),
		slog.Int64("encounter", encounterID), slog.String("error", err.Error()))
}

// catalogueBytes is the consumables catalogue as the pack version hashes it.
var catalogueBytes = sync.OnceValues(func() ([]byte, error) {
	b, err := json.Marshal(consumables.Load())
	if err != nil {
		return nil, fmt.Errorf("coach: encode the catalogue: %w", err)
	}
	return b, nil
})

// fight is the report's fight row; ErrNoFight when there is none.
func (s *Service) fight(ctx context.Context, reportID string, index int) (reports.FightEntry, error) {
	fights, err := s.Reports.Fights(ctx, reportID)
	if err != nil {
		return reports.FightEntry{}, fmt.Errorf("coach: fights: %w", err)
	}
	for _, f := range fights {
		if f.Index == index {
			return f, nil
		}
	}
	return reports.FightEntry{}, ErrNoFight
}

// summary reads one fight's stored summary out of the bucket.
func (s *Service) summary(ctx context.Context, reportID string, index int) (summary.Summary, error) {
	if s.Summaries == nil {
		return summary.Summary{}, errors.New("coach: no bucket to read the summary from")
	}
	body, err := s.Summaries.Get(ctx, store.Keys{ReportID: reportID}.FightSummary(index))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("coach: read summary %s/%d: %w", reportID, index, err)
	}
	defer body.Close()
	raw, err := io.ReadAll(io.LimitReader(body, sims.MaxSummaryBytes))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("coach: read summary %s/%d: %w", reportID, index, err)
	}
	var sum summary.Summary
	if err := json.Unmarshal(raw, &sum); err != nil {
		return summary.Summary{}, fmt.Errorf("coach: decode summary %s/%d: %w", reportID, index, err)
	}
	return sum, nil
}

// rosterRow is the row whose character key, under the report's realm, is playerKey.
func rosterRow(rep reports.Report, sum summary.Summary, playerKey string) (summary.RosterRow, bool) {
	region, ruleset := reports.ReportRealm(rep)
	for _, row := range sum.Roster {
		if character.KeyFromUnit(region, ruleset, row.Name) == playerKey {
			return row, true
		}
	}
	return summary.RosterRow{}, false
}

// fightRoster is the fight's player names as a review would write them, and their realms:
// what neither the prompt pack nor a model review may carry (§19).
type fightRoster struct {
	Names, Realms []string
}

// fightRosterOf is the fight's roster names and distinct realms.
func fightRosterOf(sum summary.Summary) fightRoster {
	out := fightRoster{Names: make([]string, 0, len(sum.Roster))}
	for _, row := range sum.Roster {
		name, realm := character.SplitUnit(row.Name)
		out.Names = append(out.Names, name)
		if realm != "" && !slices.Contains(out.Realms, realm) {
			out.Realms = append(out.Realms, realm)
		}
	}
	return out
}

// personaFor is the persona of the pack's spec; false when the spec has none.
func (s *Service) personaFor(p Pack) (Persona, bool) {
	if s.Personas == nil || p.Context.SpecLabel == "" {
		return Persona{}, false
	}
	persona, err := s.Personas.For(rating.DataSpecSlug(p.Context.Class, p.Context.SpecLabel))
	if err != nil {
		s.logger().Warn("coach: no persona; serving the rule-only review",
			slog.String("spec", p.Context.SpecLabel), slog.String("error", err.Error()))
		return Persona{}, false
	}
	return persona, true
}
