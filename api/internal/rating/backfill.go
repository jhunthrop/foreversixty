// api/internal/rating/backfill.go
package rating

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"io"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// BackfillJobCommand is this lane's Cloud Run job name, dispatched the same way
// sims.ValidateJobCommand/sims.SimRunJobCommand already are (api/cmd/api/main.go's
// os.Args[1] switch).
const BackfillJobCommand = "rating-backfill"

// BackfillBatchSize bounds one run: the spec's own cost note (§4.4/§4.5) treats
// over-recomputing as a compute cost, not a correctness one, so this can be generous
// without risk - 500 fights is comfortably inside one Cloud Run job's timeout at the
// per-fight cost §4.5 already budgets for fight-close itself.
const BackfillBatchSize = 500

// MaxBackfillSummaryBytes bounds a stored fight summary read back out of the bucket -
// mirrors sims.MaxSummaryBytes (api/internal/sims/summaries.go): nothing the ingest ever
// accepted is bigger than this.
const MaxBackfillSummaryBytes = 64 << 20

// Getter reads one stored object back. *r2.Client satisfies it; a nil Getter means this
// deployment has no bucket, and Backfill no-ops, logging why rather than failing.
type Getter interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
}

// BackfillDeps is what one backfill run needs.
type BackfillDeps struct {
	Store     *Store
	Summaries Getter
	Log       interface {
		Error(msg string, args ...any)
		Warn(msg string, args ...any)
	}
}

func (d BackfillDeps) logf() interface {
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
} {
	if d.Log != nil {
		return d.Log
	}
	return noopLogger{}
}

type noopLogger struct{}

func (noopLogger) Error(string, ...any) {}
func (noopLogger) Warn(string, ...any)  {}

// Backfill recomputes every rating_scores row not yet at logs/engine/rating's current
// DefaultModelVersion, up to batchSize fights, and upserts them via Store.RateFight -
// which, because every one of these rows already exists, always takes the rewrite path
// (spec §4.4: "digest tables are not rebuilt from scratch on a backfill"). It is
// resumable with no extra state: each run only ever selects rows still stale after the
// previous run's writes, so a crash mid-run simply leaves the next run a slightly larger
// batch to work through, and safe to run concurrently with live ingest, because
// RateFight's own per-fight advisory lock (store.go) serialises the two.
func Backfill(ctx context.Context, d BackfillDeps, batchSize int) (recomputed int, err error) {
	if d.Store == nil || d.Summaries == nil {
		d.logf().Warn("rating", "op", "backfill", "err", "no-op: Store or Summaries is nil (this deployment has no bucket configured)")
		return 0, nil
	}
	if batchSize <= 0 {
		batchSize = BackfillBatchSize
	}
	fights, err := d.Store.staleFights(ctx, batchSize)
	if err != nil {
		return 0, fmt.Errorf("rating: backfill: %w", err)
	}
	for _, sf := range fights {
		region, ruleset := sf.Region, sf.Ruleset
		if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
			// A report with no logging character: the same default ingest uses.
			region, ruleset = "us", character.RulesetNormal
		}
		sum, err := readSummary(ctx, d.Summaries, sf.ReportID, sf.FightIndex)
		if err != nil {
			d.logf().Error("rating", "op", "backfill", "report", sf.ReportID, "fight", sf.FightIndex, "err", err)
			continue
		}
		if err := d.Store.RateFight(ctx, RatedFight{
			ReportID: sf.ReportID, FightIndex: sf.FightIndex, Region: region, Ruleset: ruleset,
			FoughtAt: sf.FoughtAt, EncounterID: sum.EncounterID, Summary: sum,
		}); err != nil {
			d.logf().Error("rating", "op", "backfill", "report", sf.ReportID, "fight", sf.FightIndex, "err", err)
			continue
		}
		recomputed++
	}
	return recomputed, nil
}

// readSummary reads one fight's stored summary out of the bucket - the same shape
// api/internal/sims/summaries.go's own unexported fightSummary uses (MaxSummaryBytes
// bound, store.Keys for the object key), duplicated in miniature rather than imported:
// importing api/internal/sims here to reach one unexported helper is not possible, and
// the alternative (exporting it from sims for one cross-lane caller) is out of this
// lane's file ownership. The duplicated part is plain object-read boilerplate, not
// business logic.
func readSummary(ctx context.Context, get Getter, reportID string, index int) (summary.Summary, error) {
	body, err := get.Get(ctx, store.Keys{ReportID: reportID}.FightSummary(index))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("rating: read summary %s/%d: %w", reportID, index, err)
	}
	defer body.Close()
	b, err := io.ReadAll(io.LimitReader(body, MaxBackfillSummaryBytes))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("rating: read summary %s/%d: %w", reportID, index, err)
	}
	var s summary.Summary
	if err := json.Unmarshal(b, &s); err != nil {
		return summary.Summary{}, fmt.Errorf("rating: decode summary %s/%d: %w", reportID, index, err)
	}
	return s, nil
}
