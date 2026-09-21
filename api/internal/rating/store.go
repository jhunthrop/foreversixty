// api/internal/rating/store.go
package rating

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// RatedFight is one verified, ranked fight handed to the rating pipeline - the already-
// decoded summary the ingest holds in memory (spec §4.5: "no extra summary read is
// needed"), plus the region/ruleset a roster row's name is keyed under (the same pair
// api/internal/reports.ReportRealm already resolves for i.rank/i.score).
//
// This is deliberately this package's own type, not an alias for
// api/internal/reports.RatedFight: that lets this file's tests run without depending on
// the reports package landing its own copy first. api/internal/rating/rater.go is the one
// place that converts between the two.
type RatedFight struct {
	ReportID        string
	FightIndex      int
	Region, Ruleset string
	FoughtAt        time.Time
	EncounterID     int64
	Summary         summary.Summary
}

// Store is every rating read and write.
type Store struct {
	Pool *pgxpool.Pool
	Log  *slog.Logger
}

func (s *Store) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

// RateFight scores every roster player in f and writes their rows, folding each scored
// component's value into its percentile digest exactly once - on this fight's first
// genuine write, never on a rewrite (a re-verified fight arriving again) or a backfill
// recompute. Mirrors rankings.Store.WriteFight's own lock-then-detect-rewrite shape
// closely, on purpose: the two stores solve the same idempotency problem the same way.
func (s *Store) RateFight(ctx context.Context, f RatedFight) error {
	if len(f.Summary.Roster) == 0 {
		return nil
	}
	if err := db.EnsureRatingsPartition(ctx, s.Pool, f.FoughtAt); err != nil {
		return err
	}
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("rating: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	// Serialises this fight's writers exactly as rankings.Store.WriteFight's own comment
	// explains: without the lock two writers of the same fight could both see zero rows
	// withdrawn, both fold every value, and double-count the fight in every digest. A
	// distinct string prefix ("rating:") keeps this lock namespace separate from
	// rankings' own advisory lock on the same (report_id, fight_index) pair - harmless if
	// they collided, but a stray cross-package wait is easier to reason about avoided.
	if _, err := tx.Exec(ctx, `select pg_advisory_xact_lock(hashtext('rating:' || $1), $2)`,
		f.ReportID, f.FightIndex); err != nil {
		return fmt.Errorf("rating: lock %s/%d: %w", f.ReportID, f.FightIndex, err)
	}

	withdrawn, err := tx.Exec(ctx, `delete from rating_scores where report_id = $1 and fight_index = $2`,
		f.ReportID, f.FightIndex)
	if err != nil {
		return fmt.Errorf("rating: withdraw %s/%d: %w", f.ReportID, f.FightIndex, err)
	}
	fold := withdrawn.RowsAffected() == 0

	if fold && f.Summary.Kill {
		if err := foldKillDuration(ctx, tx, f.Summary.EncounterID, f.Summary.Difficulty, f.Summary.DurationMS); err != nil {
			return err
		}
	}

	meta := fightMeta{
		ReportID: f.ReportID, FightIndex: f.FightIndex, EncounterID: f.Summary.EncounterID,
		Difficulty: f.Summary.Difficulty, DurationMS: f.Summary.DurationMS, Kill: f.Summary.Kill,
		FoughtAt: f.FoughtAt,
	}
	adapter := &percentileAdapter{Tx: tx, Fold: fold, Log: s.logger()}
	if band, ok := adapter.KillTimeBand(f.Summary.EncounterID, f.Summary.Difficulty, f.Summary.DurationMS); ok && f.Summary.Kill {
		meta.KillTimeBand = band
	}
	modelInfo := ratingengine.DefaultModelInfo()

	for _, row := range f.Summary.Roster {
		tables := curatedTablesFor(row, f.Summary.EncounterID)
		card := ratingengine.Score(f.Summary, row.GUID, tables, adapter, nil, modelInfo)
		cr := newCardRow(meta, row, card)
		cr.PlayerKey = character.KeyFromUnit(f.Region, f.Ruleset, row.Name)
		if err := insertCardRow(ctx, tx, cr); err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("rating: commit: %w", err)
	}
	return nil
}

func insertCardRow(ctx context.Context, tx pgx.Tx, cr CardRow) error {
	_, err := tx.Exec(ctx,
		`insert into rating_scores (report_id, fight_index, player_key, player_name, class, spec, role,
		   encounter_id, difficulty, size, duration_ms, kill, kill_time_band, overall, overall_uncapped,
		   overall_capped, components, model_version, fought_at)
		 values ($1, $2, $3, $4, nullif($5, ''), nullif($6, ''), nullif($7, ''), $8, $9, $10, $11, $12, $13,
		   $14, $15, $16, $17, $18, $19)
		 on conflict (report_id, fight_index, player_key, fought_at) do update set
		   player_name = excluded.player_name, class = excluded.class, spec = excluded.spec,
		   role = excluded.role, overall = excluded.overall, overall_uncapped = excluded.overall_uncapped,
		   overall_capped = excluded.overall_capped, components = excluded.components,
		   model_version = excluded.model_version, computed_at = now()`,
		cr.ReportID, cr.FightIndex, cr.PlayerKey, cr.PlayerName, cr.Class, cr.Spec, cr.Role,
		cr.EncounterID, cr.Difficulty, cr.Size, cr.DurationMS, cr.Kill, cr.KillTimeBand,
		cr.Overall, cr.OverallUncapped, cr.OverallCapped, cr.Components, cr.ModelVersion, cr.FoughtAt)
	if err != nil {
		return fmt.Errorf("rating: write row %s/%d/%s: %w", cr.ReportID, cr.FightIndex, cr.PlayerKey, err)
	}
	return nil
}
