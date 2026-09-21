// api/internal/rating/store.go
package rating

import (
	"context"
	"errors"
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

// cursorPos is the trend list's keyset position: the (fought_at, report_id, fight_index)
// of the last row a page returned, matching reports/recent.go's own cursor shape.
type cursorPos struct {
	FoughtAt   time.Time
	ReportID   string
	FightIndex int
}

// ReadFightRatings reads every stored rating row for one fight. ok is false when the
// fight has no rows at all - never verified, still queued in the async Rater, or (per §4.3)
// a fight this lane never rates at all (trash, no ranker, or a report that was never
// Ranked). The caller (the handler) does not distinguish those three at the HTTP layer:
// all read as 404, matching how a report's own file routes already treat "nothing here"
// without describing why.
func (s *Store) ReadFightRatings(ctx context.Context, reportID string, fightIndex int) ([]CardRow, bool, error) {
	rows, err := s.Pool.Query(ctx,
		`select player_key, player_name, coalesce(class, ''), coalesce(spec, ''), coalesce(role, ''),
		    encounter_id, difficulty, coalesce(size, 0), coalesce(duration_ms, 0), kill, kill_time_band,
		    overall, overall_uncapped, overall_capped, components, model_version, fought_at
		 from rating_scores where report_id = $1 and fight_index = $2 order by player_name`,
		reportID, fightIndex)
	if err != nil {
		return nil, false, fmt.Errorf("rating: read %s/%d: %w", reportID, fightIndex, err)
	}
	defer rows.Close()
	var out []CardRow
	for rows.Next() {
		var cr CardRow
		if err := rows.Scan(&cr.PlayerKey, &cr.PlayerName, &cr.Class, &cr.Spec, &cr.Role,
			&cr.EncounterID, &cr.Difficulty, &cr.Size, &cr.DurationMS, &cr.Kill, &cr.KillTimeBand,
			&cr.Overall, &cr.OverallUncapped, &cr.OverallCapped, &cr.Components, &cr.ModelVersion, &cr.FoughtAt); err != nil {
			return nil, false, fmt.Errorf("rating: scan %s/%d: %w", reportID, fightIndex, err)
		}
		cr.ReportID, cr.FightIndex = reportID, fightIndex
		out = append(out, cr)
	}
	if err := rows.Err(); err != nil {
		return nil, false, err
	}
	return out, len(out) > 0, nil
}

// anonymized reports whether playerKey's owning account has asked to be anonymous
// (users.anonymize), joined the same way auth.User.PublicName's own doc names as the join
// a character-scoped anonymize check would need. A player_key with no characters row, or
// no linked account, is never anonymized - there is nothing to hide.
func (s *Store) anonymized(ctx context.Context, playerKey string) (bool, error) {
	var anon bool
	err := s.Pool.QueryRow(ctx,
		`select u.anonymize from characters c join users u on u.id = c.user_id where c.key = $1`,
		playerKey).Scan(&anon)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("rating: anonymize check %s: %w", playerKey, err)
	}
	return anon, nil
}

// characterTrendPageSize bounds one page of the character endpoint's trend[]; a career's
// worth of rated fights is unbounded, so this list is paginated like every other list read
// in this API (reports/recent.go's own convention).
const characterTrendPageSize = 20

// ReadCharacterRating reads up to limit rows of playerKey's rating history, newest first,
// from public reports only (spec §5.1's ruling: "public at time of query", enforced here
// by joining the live reports.visibility column rather than trusting anything cached on
// the row at write time), starting after before when it is non-nil. hasMore is true when
// a further page exists.
func (s *Store) ReadCharacterRating(ctx context.Context, playerKey string, limit int, before *cursorPos) ([]CardRow, bool, error) {
	if limit <= 0 || limit > characterTrendPageSize {
		limit = characterTrendPageSize
	}
	query := `select rs.player_key, rs.player_name, coalesce(rs.class, ''), coalesce(rs.spec, ''),
	    coalesce(rs.role, ''), rs.encounter_id, rs.difficulty, coalesce(rs.size, 0),
	    coalesce(rs.duration_ms, 0), rs.kill, rs.kill_time_band, rs.overall, rs.overall_uncapped,
	    rs.overall_capped, rs.components, rs.model_version, rs.report_id, rs.fight_index, rs.fought_at
	 from rating_scores rs
	 join reports r on r.id = rs.report_id
	 where rs.player_key = $1 and r.visibility = 'public'`
	args := []any{playerKey}
	if before != nil {
		query += ` and (rs.fought_at, rs.report_id, rs.fight_index) < ($2, $3, $4)`
		args = append(args, before.FoughtAt, before.ReportID, before.FightIndex)
	}
	query += ` order by rs.fought_at desc, rs.report_id desc, rs.fight_index desc limit $` + fmt.Sprint(len(args)+1)
	args = append(args, limit+1) // one extra row, to answer hasMore without a second query

	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, false, fmt.Errorf("rating: read character %s: %w", playerKey, err)
	}
	defer rows.Close()
	var out []CardRow
	for rows.Next() {
		var cr CardRow
		if err := rows.Scan(&cr.PlayerKey, &cr.PlayerName, &cr.Class, &cr.Spec, &cr.Role,
			&cr.EncounterID, &cr.Difficulty, &cr.Size, &cr.DurationMS, &cr.Kill, &cr.KillTimeBand,
			&cr.Overall, &cr.OverallUncapped, &cr.OverallCapped, &cr.Components, &cr.ModelVersion,
			&cr.ReportID, &cr.FightIndex, &cr.FoughtAt); err != nil {
			return nil, false, fmt.Errorf("rating: scan character %s: %w", playerKey, err)
		}
		out = append(out, cr)
	}
	if err := rows.Err(); err != nil {
		return nil, false, err
	}
	hasMore := len(out) > limit
	if hasMore {
		out = out[:limit]
	}
	return out, hasMore, nil
}
