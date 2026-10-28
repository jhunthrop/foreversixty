// api/internal/rating/backfill.go
package rating

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"io"
	"time"

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

// RerateFlaggedJobCommand is the Cloud Run job that brings the stored Forever Score back
// into line after reports were flagged (RerateFlagged), dispatched beside BackfillJobCommand.
const RerateFlaggedJobCommand = "rating-rerate-flagged"

// rerateGroup is why a character's score moved with a flagged report, the order
// RerateFlagged takes them in: a character is counted in the first group it is in.
type rerateGroup int

const (
	// rerateInFlagged: the character has a rated row in a flagged report (its fights,
	// as_of, discipline and parses counted it).
	rerateInFlagged rerateGroup = iota
	// rerateGuildmate: the character is a verified member of a guild with a flagged report, or
	// has a rated fight in the attendance window in one of its reports (the night counted
	// toward its attendance's nights).
	rerateGuildmate
	// rerateLinked: the character is linked to the same account as a character in a flagged
	// report (the report's kills counted toward its experience).
	rerateLinked
)

// RerateGroup is what one RerateFlagged run did for one group of characters.
type RerateGroup struct {
	// Characters is how many characters are in the group.
	Characters int
	// Fights is how many distinct fights the group's characters picked that no earlier
	// group had picked: each character's newest rated fight in a public, unflagged report (a
	// raid's whole roster usually shares one). A character with no such fight left has no
	// score and nothing to roll it onto.
	Fights int
	// Rerated is how many of those fights were re-rated and rolled up; Failed how many could
	// not be (each logged).
	Rerated, Failed int
}

// RerateCounts is what one RerateFlagged run did, per group, for its log line.
type RerateCounts struct {
	// Flagged is the characters with a rated row in a flagged report.
	Flagged RerateGroup
	// Guildmates is the other verified members of a guild that has a flagged report (the
	// membership guildNights counts nights for), and the other characters with a rated fight in
	// the last AttendanceWindow in one of its reports: the night was one of their guild's nights,
	// so their attendance's denominator moved, whether they were there or not.
	Guildmates RerateGroup
	// Linked is the other characters linked to the account of a character in Flagged: the
	// flagged report's kills counted toward their experience.
	Linked RerateGroup
}

// Failed is how many fights of every group could not be re-rated.
func (c RerateCounts) Failed() int { return c.Flagged.Failed + c.Guildmates.Failed + c.Linked.Failed }

// Fights is how many distinct fights every group picked.
func (c RerateCounts) Fights() int { return c.Flagged.Fights + c.Guildmates.Fights + c.Linked.Fights }

func (c *RerateCounts) group(g rerateGroup) *RerateGroup {
	switch g {
	case rerateGuildmate:
		return &c.Guildmates
	case rerateLinked:
		return &c.Linked
	default:
		return &c.Flagged
	}
}

// RerateFlagged re-rates, through Backfill's own path (the fight's stored summary read back
// out of the bucket, then RateFight), the newest rated fight in a public, unflagged report
// of every character whose score a flagged report moved, and rolls the fight's players'
// Forever Scores up onto it (RollUpScores, the rater's own second step). The score's inputs
// skip a flagged report (presence.EvidenceReport), so the score rolled up is the one without
// it, and the snapshot the character page's score history ends on drops back. The model
// version is not bumped: the formula is the same, only its inputs changed.
//
// The affected set is three groups (RerateCounts): the characters in a flagged report; the
// guild-mates of a flagged night (the verified members of its guild, and every character with a
// rated fight in the attendance window in that guild's reports, whose attendance counted the
// night among the guild's nights, present or absent); and the characters linked to the account
// of a character in a flagged report (whose experience counted its kills). A character is in
// the first group it fits.
//
// It is idempotent: a second run picks the same fights and writes the same figures (the
// score as it stands now). A flagged report's own rows are left in place, never read; older
// snapshots, which recorded the score of their own day, are left as they were. The roll-up
// writes the score as it stands now for every player of a picked fight, which is the
// rater's own rule. A run with no bucket configured is a no-op, logged, as Backfill's is.
func RerateFlagged(ctx context.Context, d BackfillDeps) (RerateCounts, error) {
	var out RerateCounts
	if d.Store == nil || d.Summaries == nil {
		d.logf().Warn("rating", "op", "rerate_flagged", "err", "no-op: Store or Summaries is nil (this deployment has no bucket configured)")
		return out, nil
	}
	picks, err := d.Store.flaggedRerates(ctx, time.Now().Add(-AttendanceWindow))
	if err != nil {
		return out, fmt.Errorf("rating: rerate flagged: %w", err)
	}
	type fightKey struct {
		report string
		index  int
	}
	seen := map[fightKey]bool{}
	for _, p := range picks {
		g := out.group(p.group)
		g.Characters++
		if p.ReportID == "" {
			continue
		}
		k := fightKey{p.ReportID, p.FightIndex}
		if seen[k] {
			continue
		}
		seen[k] = true
		g.Fights++
		if err := rerateFight(ctx, d, p.staleFight); err != nil {
			d.logf().Error("rating", "op", "rerate_flagged", "report", p.ReportID, "fight", p.FightIndex, "err", err)
			g.Failed++
			continue
		}
		g.Rerated++
	}
	return out, nil
}

// loadRatedFight reads a picked fight's stored summary into the RatedFight RateFight takes,
// defaulting an unusable region or ruleset to the one ingest uses for an unattributed report.
func loadRatedFight(ctx context.Context, d BackfillDeps, sf staleFight) (RatedFight, error) {
	region, ruleset := sf.Region, sf.Ruleset
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
		region, ruleset = "us", character.RulesetNormal
	}
	sum, err := readSummary(ctx, d.Summaries, sf.ReportID, sf.FightIndex)
	if err != nil {
		return RatedFight{}, err
	}
	return RatedFight{
		ReportID: sf.ReportID, FightIndex: sf.FightIndex, Region: region, Ruleset: ruleset,
		FoughtAt: sf.FoughtAt, EncounterID: sum.EncounterID, Summary: sum,
	}, nil
}

// rerateFight is one picked fight through Backfill's path, then the roll-up.
func rerateFight(ctx context.Context, d BackfillDeps, sf staleFight) error {
	f, err := loadRatedFight(ctx, d, sf)
	if err != nil {
		return err
	}
	if err := d.Store.RateFight(ctx, f); err != nil {
		return err
	}
	return d.Store.RollUpScores(ctx, f)
}
