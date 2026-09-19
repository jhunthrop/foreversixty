package sims

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// MaxSummaryBytes bounds a stored fight summary read back out of the
// bucket. The ingest never accepted a bundle larger than this, so
// nothing bigger can be there, and no reader here may be at the mercy
// of the bucket for how much it holds in memory.
const MaxSummaryBytes = 64 << 20

// Getter reads one object back. *r2.Client satisfies it; a nil Getter
// means this deployment has no bucket, and every caller degrades to
// what the database alone can answer.
type Getter interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
}

// fightSummary reads one fight's stored summary out of the bucket.
func fightSummary(ctx context.Context, get Getter, reportID string, index int) (summary.Summary, error) {
	body, err := get.Get(ctx, store.Keys{ReportID: reportID}.FightSummary(index))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("sims: read summary %s/%d: %w", reportID, index, err)
	}
	defer body.Close()
	b, err := io.ReadAll(io.LimitReader(body, MaxSummaryBytes))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("sims: read summary %s/%d: %w", reportID, index, err)
	}
	var s summary.Summary
	if err := json.Unmarshal(b, &s); err != nil {
		return summary.Summary{}, fmt.Errorf("sims: decode summary %s/%d: %w", reportID, index, err)
	}
	return s, nil
}

// combatantNamed finds one player's recorded gear, talents and auras in
// a fight's summary.
func combatantNamed(s summary.Summary, name string) (summary.CombatantRow, bool) {
	for _, c := range s.Combatants {
		if c.Name == name {
			return c, true
		}
	}
	return summary.CombatantRow{}, false
}
