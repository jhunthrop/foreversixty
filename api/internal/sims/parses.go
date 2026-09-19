package sims

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// ParseReader reads the top parses for a spec out of the ranking rows
// and the fight summaries they point at. The summary read itself is
// summaries.go's, shared with GET /v1/characters/…/sim-input: one
// bucket read, written once.
type ParseReader struct {
	Pool *pgxpool.Pool
	Get  Getter
	Log  *slog.Logger
}

func (p *ParseReader) logger() *slog.Logger {
	if p.Log != nil {
		return p.Log
	}
	return slog.Default()
}

// TopParses answers the n best DPS parses for one spec in one phase.
// A parse whose stored summary cannot be read, or which the summary
// has no combatant for, is skipped: the validation figure is a
// measurement, and measuring a character we cannot reconstruct would
// make it worse rather than more complete.
func (p *ParseReader) TopParses(ctx context.Context, spec, phase string, n int) ([]Parse, error) {
	rows, err := p.Pool.Query(ctx,
		`select m.report_id, m.fight_index, m.player_key, m.player_name,
		        coalesce(m.class, ''), m.metric_dps, coalesce(m.duration_ms, 0)
		 from fight_metrics m
		 where m.spec = $1 and m.phase = $2 and m.kill and m.state <> 'removed'
		   and m.role = $3 and m.metric_dps is not null and m.duration_ms > 0
		 order by m.metric_dps desc limit $4`, spec, phase, roleDPS, n)
	if err != nil {
		return nil, fmt.Errorf("sims: top parses for %s: %w", spec, err)
	}
	defer rows.Close()

	// The rows are read out in full before anything touches the
	// bucket: a query's rows must be drained before the connection
	// is used again, and the summary reads below are slow.
	type ref struct {
		Parse
		playerName string
	}
	refs := []ref{}
	for rows.Next() {
		var (
			r          ref
			durationMS int
		)
		if err := rows.Scan(&r.ReportID, &r.FightIndex, &r.PlayerKey, &r.playerName,
			&r.Class, &r.ActualDPS, &durationMS); err != nil {
			return nil, fmt.Errorf("sims: scan parse: %w", err)
		}
		r.Spec, r.DurationSec = spec, durationMS/1000
		refs = append(refs, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	out := make([]Parse, 0, len(refs))
	for _, r := range refs {
		s, err := fightSummary(ctx, p.Get, r.ReportID, r.FightIndex)
		if err != nil {
			p.logger().Warn("sims", "op", "validate", "report", r.ReportID,
				"fight", r.FightIndex, "err", err)
			continue
		}
		c, ok := combatantNamed(s, r.playerName)
		if !ok {
			continue
		}
		r.Combatant, r.ActualCasts = c, castsBy(s, c.GUID)
		out = append(out, r.Parse)
	}
	return out, nil
}

// castsBy is how many times one actor landed each ability, which is
// the actual half of the cast-count comparison. It is keyed by the
// summary's row identity, the same key the sim's rows carry, and keeps
// the log's name for the card.
func castsBy(s summary.Summary, guid string) map[int64]CastCount {
	out := map[int64]CastCount{}
	for _, c := range s.Casts {
		if c.GUID != guid {
			continue
		}
		row := out[c.SpellID]
		row.Casts += c.Succeeded
		if row.Name == "" {
			row.Name = c.SpellName
		}
		out[c.SpellID] = row
	}
	return out
}
