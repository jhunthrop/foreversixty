// api/internal/rating/rebuild_preparation.go
package rating

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// RebuildPreparationJobCommand is the Cloud Run job that re-rates every stored fight under
// the current model version and refolds the preparation digests from the stored summaries,
// after the consumables-at-pull fix moved DefaultModelVersion (Backfill alone would leave
// the digests holding the values the old, wrong reading placed).
const RebuildPreparationJobCommand = "rating-rebuild-preparation"

// RebuildBatchSize bounds one transaction of digest folds and one page of stale fights.
const RebuildBatchSize = 200

// RebuildCounts is what one RebuildPreparation run did.
type RebuildCounts struct {
	// Folded is the rated kills whose preparation values were folded back into the digests.
	Folded int
	// Rerated is the fights re-rated to the current model version; RolledUp those whose
	// players' Forever Scores were then rolled up.
	Rerated, RolledUp int
	// Failed is the fights that could not be read, folded, re-rated or rolled up (each logged).
	Failed int
}

// RebuildOptions tunes one run.
type RebuildOptions struct {
	// BatchSize is RebuildBatchSize when zero.
	BatchSize int
	// DryRun counts what a run would fold and re-rate and writes nothing.
	DryRun bool
	// After resumes the fold phase past a fight (the cursor Progress last reported). A run
	// with a cursor does not clear the digests again, since the earlier run already did.
	After *FightCursor
	// Progress, when set, is called after every batch with the running counts and the cursor
	// to resume from.
	Progress func(phase string, counts RebuildCounts, cursor FightCursor)
}

// FightCursor is the fold phase's keyset position.
type FightCursor struct {
	ReportID   string
	FightIndex int
}

// RebuildPreparation runs RebuildPreparationWith with the defaults: clears every
// component='preparation' digest row, folds each rated kill's raw PreparationValue back in
// (one read of each stored summary), then re-rates until no fight is stale and rolls each
// re-rated fight's scores up. Idempotent: a second run refolds from an empty digest again
// and re-rates nothing.
func RebuildPreparation(ctx context.Context, d BackfillDeps) (RebuildCounts, error) {
	return RebuildPreparationWith(ctx, d, RebuildOptions{})
}

// RebuildPreparationWith is RebuildPreparation with options. A run with no bucket is a
// logged no-op, as Backfill's is.
func RebuildPreparationWith(ctx context.Context, d BackfillDeps, o RebuildOptions) (RebuildCounts, error) {
	var out RebuildCounts
	if d.Store == nil || d.Summaries == nil {
		d.logf().Warn("rating", "op", "rebuild_preparation", "err", "no-op: Store or Summaries is nil (this deployment has no bucket configured)")
		return out, nil
	}
	if o.BatchSize <= 0 {
		o.BatchSize = RebuildBatchSize
	}
	if o.After == nil && !o.DryRun {
		if _, err := d.Store.Pool.Exec(ctx, `delete from rating_percentile_digests where component = $1`,
			ratingengine.ComponentNamePreparation); err != nil {
			return out, fmt.Errorf("rating: rebuild: clear preparation digests: %w", err)
		}
	}
	if err := d.Store.refoldPreparation(ctx, d, o, &out); err != nil {
		return out, err
	}
	if err := d.Store.rerateStale(ctx, d, o, &out); err != nil {
		return out, err
	}
	return out, nil
}

// refoldPreparation pages through every rated fight, folding the kills' preparation values.
func (s *Store) refoldPreparation(ctx context.Context, d BackfillDeps, o RebuildOptions, out *RebuildCounts) error {
	cursor := FightCursor{FightIndex: -1}
	if o.After != nil {
		cursor = *o.After
	}
	for {
		page, err := s.ratedFightPage(ctx, cursor, o.BatchSize)
		if err != nil {
			return err
		}
		if len(page) == 0 {
			return nil
		}
		if err := s.foldPage(ctx, d, o, page, out); err != nil {
			return err
		}
		cursor = page[len(page)-1]
		if o.Progress != nil {
			o.Progress("fold", *out, cursor)
		}
	}
}

func (s *Store) ratedFightPage(ctx context.Context, after FightCursor, limit int) ([]FightCursor, error) {
	rows, err := s.Pool.Query(ctx,
		`select distinct report_id, fight_index from rating_scores
		  where (report_id, fight_index) > ($1, $2) order by report_id, fight_index limit $3`,
		after.ReportID, after.FightIndex, limit)
	if err != nil {
		return nil, fmt.Errorf("rating: rebuild: list rated fights: %w", err)
	}
	defer rows.Close()
	var out []FightCursor
	for rows.Next() {
		var c FightCursor
		if err := rows.Scan(&c.ReportID, &c.FightIndex); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// foldPage folds one page of fights in a single transaction (nothing in a dry run).
func (s *Store) foldPage(ctx context.Context, d BackfillDeps, o RebuildOptions, page []FightCursor, out *RebuildCounts) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("rating: rebuild: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	adapter := &percentileAdapter{Tx: tx, Fold: !o.DryRun, Log: s.logger()}
	var folded int
	for _, c := range page {
		sum, err := readSummary(ctx, d.Summaries, c.ReportID, c.FightIndex)
		if err != nil {
			d.logf().Error("rating", "op", "rebuild_preparation", "report", c.ReportID, "fight", c.FightIndex, "err", err)
			out.Failed++
			continue
		}
		if foldPreparation(sum, adapter) {
			folded++
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("rating: rebuild: commit: %w", err)
	}
	out.Folded += folded
	return nil
}

// foldPreparation places every roster player's raw preparation value of a kill through the
// adapter, which folds it; false for a wipe, which never enters the digest (foldEligible).
func foldPreparation(sum summary.Summary, adapter *percentileAdapter) bool {
	if !sum.Kill {
		return false
	}
	band, ok := adapter.KillTimeBand(sum.EncounterID, sum.Difficulty, sum.DurationMS)
	if !ok {
		band = ratingengine.BandTypical
	}
	for _, row := range sum.Roster {
		var cat *consumables.RoleCatalogue
		if rc, ok := consumables.Load().For(row.Role); ok {
			cat = &rc
		}
		value, ok := ratingengine.PreparationValue(sum, row.GUID, cat)
		if !ok {
			continue
		}
		adapter.Placement(preparationBracket(sum, row, band), value)
	}
	return true
}

func preparationBracket(sum summary.Summary, row summary.RosterRow, band string) ratingengine.Bracket {
	return ratingengine.Bracket{
		EncounterID: sum.EncounterID, Difficulty: sum.Difficulty, Spec: specSlug(row.Class, row.Spec),
		Role: row.Role, KillTimeBand: band, Kill: true, Component: ratingengine.ComponentNamePreparation,
	}
}

// rerateStale re-rates and rolls up stale fights batch by batch until none is left. A fight
// that fails stays stale, so later batches ask for enough extra rows to step past it.
func (s *Store) rerateStale(ctx context.Context, d BackfillDeps, o RebuildOptions, out *RebuildCounts) error {
	if o.DryRun {
		return s.countStale(ctx, out)
	}
	failed := map[FightCursor]bool{}
	for {
		stale, err := s.staleFights(ctx, o.BatchSize+len(failed))
		if err != nil {
			return err
		}
		progressed := false
		for _, sf := range stale {
			key := FightCursor{sf.ReportID, sf.FightIndex}
			if failed[key] {
				continue
			}
			progressed = true
			if !rerateAndRollUp(ctx, d, sf, out) {
				failed[key] = true
			}
		}
		if !progressed {
			return nil
		}
		if o.Progress != nil {
			o.Progress("rerate", *out, FightCursor{})
		}
	}
}

// rerateAndRollUp re-rates one stale fight, then rolls its scores up; false when either
// step failed (logged and counted).
func rerateAndRollUp(ctx context.Context, d BackfillDeps, sf staleFight, out *RebuildCounts) bool {
	f, err := loadRatedFight(ctx, d, sf)
	if err == nil {
		err = d.Store.RateFight(ctx, f)
	}
	if err != nil {
		d.logf().Error("rating", "op", "rebuild_preparation", "report", sf.ReportID, "fight", sf.FightIndex, "err", err)
		out.Failed++
		return false
	}
	out.Rerated++
	if err := d.Store.RollUpScores(ctx, f); err != nil {
		d.logf().Error("rating", "op", "rebuild_roll_up", "report", sf.ReportID, "fight", sf.FightIndex, "err", err)
		out.Failed++
		return true
	}
	out.RolledUp++
	return true
}

// countStale is the dry run's re-rate figure: the fights still on an older model version.
func (s *Store) countStale(ctx context.Context, out *RebuildCounts) error {
	var n int
	err := s.Pool.QueryRow(ctx,
		`select count(*) from (select distinct report_id, fight_index from rating_scores
		   where model_version <> $1) t`, ratingengine.DefaultModelVersion).Scan(&n)
	if err != nil {
		return fmt.Errorf("rating: rebuild: count stale: %w", err)
	}
	out.Rerated = n
	return nil
}
