package coach

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"slices"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	// SamplesJobCommand is the Cloud Run job that samples a day's reviews for the owner.
	SamplesJobCommand = "coach-samples"
	// SampleSize is how many reviews a day's sample holds.
	SampleSize = 20

	daySeedStream = 0x7e57_c0ac
	hoursPerDay   = 24
	sampleDayFmt  = "2006-01-02"
)

// sampleCandidate is one review generated on the sampled day, with what the owner reads.
type sampleCandidate struct {
	reportID, playerKey, name, spec, source string
	fightIndex                              int
	validation, pack, review                json.RawMessage
}

// DaySeed is the random source for a day: the same day always picks the same reviews.
func DaySeed(day time.Time) *rand.Rand {
	return rand.New(rand.NewPCG(uint64(day.Unix()), daySeedStream))
}

// ParseSampleDay reads a YYYY-MM-DD day as midnight UTC.
func ParseSampleDay(s string) (time.Time, error) {
	d, err := time.ParseInLocation(sampleDayFmt, s, time.UTC)
	if err != nil {
		return time.Time{}, fmt.Errorf("coach: day %q is not YYYY-MM-DD: %w", s, err)
	}
	return d, nil
}

// SampleDay picks SampleSize reviews generated on day (UTC) at random, spread across specs
// (round-robin over specs in random order), and writes each one's pack, output, source and
// validation result to coach_review_samples, with one counts-only log entry per sample. It
// returns how many it wrote: zero when the day was already sampled.
func SampleDay(ctx context.Context, pool *pgxpool.Pool, day time.Time, rnd *rand.Rand, log *slog.Logger) (int, error) {
	day = day.UTC().Truncate(hoursPerDay * time.Hour)
	done, err := daySampled(ctx, pool, day)
	if err != nil || done {
		return 0, err
	}
	cands, err := dayCandidates(ctx, pool, day)
	if err != nil {
		return 0, err
	}
	picked := pickSpread(cands, SampleSize, rnd)
	if err := writeSamples(ctx, pool, day, picked); err != nil {
		return 0, err
	}
	for _, c := range picked {
		log.Info("coach sample", "day", day.Format(sampleDayFmt), "spec", c.spec, "source", c.source, "problems", problemCount(c.validation))
	}
	return len(picked), nil
}

func daySampled(ctx context.Context, pool *pgxpool.Pool, day time.Time) (bool, error) {
	var found bool
	err := pool.QueryRow(ctx, `select exists (select 1 from coach_review_samples where day = $1)`, day).Scan(&found)
	if err != nil {
		return false, fmt.Errorf("coach: check sampled day: %w", err)
	}
	return found, nil
}

func dayCandidates(ctx context.Context, pool *pgxpool.Pool, day time.Time) ([]sampleCandidate, error) {
	rows, err := pool.Query(ctx, `
		select cr.report_id, cr.fight_index, cr.player_key,
		       coalesce((select rs.player_name from rating_scores rs
		                  where rs.report_id = cr.report_id and rs.fight_index = cr.fight_index
		                    and rs.player_key = cr.player_key limit 1), ''),
		       coalesce(cp.pack #>> '{context,spec}', ''), cr.source, cr.validation, cp.pack, cr.review
		from coach_reviews cr
		join coach_packs cp using (report_id, fight_index, player_key, pack_version)
		where cr.generated_at >= $1 and cr.generated_at < $2
		order by cr.report_id, cr.fight_index, cr.player_key, cr.model_version, cr.persona_version`,
		day, day.Add(hoursPerDay*time.Hour))
	if err != nil {
		return nil, fmt.Errorf("coach: read the day's reviews: %w", err)
	}
	defer rows.Close()
	var out []sampleCandidate
	for rows.Next() {
		var c sampleCandidate
		if err := rows.Scan(&c.reportID, &c.fightIndex, &c.playerKey, &c.name, &c.spec, &c.source, &c.validation, &c.pack, &c.review); err != nil {
			return nil, fmt.Errorf("coach: scan a review: %w", err)
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// pickSpread takes up to n candidates, one spec at a time in a random spec order, a random
// review from each, round after round; a spec that runs out drops from the rotation.
func pickSpread(cands []sampleCandidate, n int, rnd *rand.Rand) []sampleCandidate {
	bySpec := map[string][]sampleCandidate{}
	for _, c := range cands {
		bySpec[c.spec] = append(bySpec[c.spec], c)
	}
	specs := make([]string, 0, len(bySpec))
	for s := range bySpec {
		specs = append(specs, s)
	}
	slices.Sort(specs)
	rnd.Shuffle(len(specs), func(i, j int) { specs[i], specs[j] = specs[j], specs[i] })
	for _, s := range specs {
		g := bySpec[s]
		rnd.Shuffle(len(g), func(i, j int) { g[i], g[j] = g[j], g[i] })
	}
	var out []sampleCandidate
	for took := true; took && len(out) < n; {
		took = false
		for _, s := range specs {
			if g := bySpec[s]; len(g) > 0 && len(out) < n {
				out, bySpec[s], took = append(out, g[0]), g[1:], true
			}
		}
	}
	return out
}

func writeSamples(ctx context.Context, pool *pgxpool.Pool, day time.Time, picked []sampleCandidate) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("coach: begin samples: %w", err)
	}
	defer tx.Rollback(ctx) //nolint:errcheck // rollback after commit is a no-op
	for _, c := range picked {
		if err := insertSample(ctx, tx, day, c); err != nil {
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("coach: commit samples: %w", err)
	}
	return nil
}

func insertSample(ctx context.Context, tx pgx.Tx, day time.Time, c sampleCandidate) error {
	_, err := tx.Exec(ctx, `
		insert into coach_review_samples
		  (day, report_id, fight_index, player_key, character_name, spec, source, validation, pack, review)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		on conflict do nothing`,
		day, c.reportID, c.fightIndex, c.playerKey, c.name, c.spec, c.source, c.validation, c.pack, c.review)
	if err != nil {
		return fmt.Errorf("coach: write a sample: %w", err)
	}
	return nil
}

// problemCount is how many problems the last model attempt left: a number, never their text.
func problemCount(validation json.RawMessage) int {
	var problems []json.RawMessage
	if json.Unmarshal(validation, &problems) != nil {
		return 0
	}
	return len(problems)
}
