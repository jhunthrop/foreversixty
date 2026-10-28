package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// TableSource loads the data Tables for a roster row (spec rules, spellconst, trainables,
// rotation, catalogue, mechanics, WorstActions) as of the fight's build.
type TableSource interface {
	TablesFor(ctx context.Context, row summary.RosterRow, encounterID int64, foughtAt time.Time) (Tables, error)
}

// ErrNoTables is what a TableSource returns when it has no data tables for the row (a spec
// without rule files, a build it does not know). The writer skips that row; any other
// TablesFor error fails the hook.
var ErrNoTables = errors.New("coach: no data tables for this player")

// ProfileWriter writes coach_profiles for a rated fight: kills only, on reports that are
// ranked (reports.Ranked: not private) and evidence (presence.EvidenceReport), read in the
// same transaction. It implements rating.FightHook.
type ProfileWriter struct {
	Tables TableSource
	Log    *slog.Logger
}

var _ rating.FightHook = (*ProfileWriter)(nil)

// evidenceReportSQL subsumes reports.Ranked: both refuse a private report.
var evidenceReportSQL = `select exists (select 1 from reports r where r.id = $1 and ` +
	presence.EvidenceReport("r") + `)`

const (
	deleteProfilesSQL = `delete from coach_profiles where report_id = $1 and fight_index = $2`

	insertProfileSQL = `insert into coach_profiles
	  (report_id, fight_index, player_key, class, spec, role, encounter_id, difficulty, size, kill,
	   duration_ms, phase, metric, fought_at, model_version, profile)
	  values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`
)

func (w *ProfileWriter) logger() *slog.Logger {
	if w.Log != nil {
		return w.Log
	}
	return slog.Default()
}

// InTx replaces the fight's profiles with the ones its cards now justify. A wipe, a fight
// shorter than MinFightMS, a report that is not evidence and a row with no spec write
// nothing; any earlier rows of the fight are still withdrawn, as RateFight withdraws its own.
func (w *ProfileWriter) InTx(ctx context.Context, tx pgx.Tx, f rating.RatedFight, cards []rating.CardRow) error {
	if _, err := tx.Exec(ctx, deleteProfilesSQL, f.ReportID, f.FightIndex); err != nil {
		return fmt.Errorf("coach: withdraw profiles %s/%d: %w", f.ReportID, f.FightIndex, err)
	}
	if !f.Summary.Kill || f.Summary.DurationMS < MinFightMS {
		return nil
	}
	var evidence bool
	if err := tx.QueryRow(ctx, evidenceReportSQL, f.ReportID).Scan(&evidence); err != nil {
		return fmt.Errorf("coach: read report %s: %w", f.ReportID, err)
	}
	if !evidence {
		return nil
	}
	cardOf := make(map[string]rating.CardRow, len(cards))
	for _, c := range cards {
		cardOf[c.PlayerName] = c
	}
	for _, row := range f.Summary.Roster {
		card, ok := cardOf[row.Name]
		if !ok || row.Spec == "" || unmeasured(row, card) {
			continue
		}
		if err := w.writeRow(ctx, tx, f, row, card); err != nil {
			return err
		}
	}
	return nil
}

func (w *ProfileWriter) writeRow(ctx context.Context, tx pgx.Tx, f rating.RatedFight,
	row summary.RosterRow, card rating.CardRow) error {
	tables, err := w.Tables.TablesFor(ctx, row, f.EncounterID, f.FoughtAt)
	if errors.Is(err, ErrNoTables) {
		// The spec has no data tables: that row has no profile, the rest of the fight does.
		w.logger().Warn("coach", "op", "profile_tables", "report", f.ReportID, "fight", f.FightIndex,
			"player", card.PlayerKey, "err", err)
		return nil
	}
	if err != nil {
		// Any other failure rolls the rating back, so the fight is retried and the old
		// profile survives.
		return fmt.Errorf("coach: tables for %s: %w", card.PlayerKey, err)
	}
	profile := BuildProfile(Input{Fight: f.Summary, Row: row, Size: card.Size, Tables: tables})
	raw, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("coach: marshal profile of %s: %w", card.PlayerKey, err)
	}
	if _, err := tx.Exec(ctx, insertProfileSQL,
		f.ReportID, f.FightIndex, card.PlayerKey, row.Class, row.Spec, row.Role, f.EncounterID,
		f.Summary.Difficulty, card.Size, f.Summary.Kill, f.Summary.DurationMS, phase.At(f.FoughtAt),
		profileMetric(row, card), f.FoughtAt, RulesVersion, raw); err != nil {
		return fmt.Errorf("coach: write profile of %s: %w", card.PlayerKey, err)
	}
	return nil
}

// unmeasured reports a healer or tank whose card is insufficient: its Overall is 0 by
// contract, and a 0 must never enter the pool as a metric.
func unmeasured(row summary.RosterRow, card rating.CardRow) bool {
	return row.Role != RoleDPS && card.Insufficient
}

// profileMetric is what "top" means: DPS for damage dealers, the rating's overall score
// for healers and tanks (§11.1).
func profileMetric(row summary.RosterRow, card rating.CardRow) float64 {
	if row.Role == RoleDPS {
		return row.DPS
	}
	return card.Overall
}
