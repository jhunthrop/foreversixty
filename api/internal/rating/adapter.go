// api/internal/rating/adapter.go
package rating

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/digest"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

// killDurationBandFast/Typical/Slow thresholds, spec §1.2's exact formula.
const (
	killDurationFastRatio = 0.85
	killDurationSlowRatio = 1.15
)

// foldEligible reports whether bracket b's value may be folded into its digest for this
// fight, spec §2's wipe rule: Survival/Mechanics' four sub-brackets fold on a wipe as well
// as a kill; every other component - Output, Utility, Preparation, Activity, addressed
// directly by their top-level component names since only Survival and Mechanics have
// sub-brackets (RULING R9 in logs/engine/rating) - folds only on a kill.
func foldEligible(b ratingengine.Bracket) bool {
	if b.Kill {
		return true
	}
	switch b.Component {
	case ratingengine.ComponentSurvivalAvoidableHit,
		ratingengine.ComponentMechanicsInterrupt,
		ratingengine.ComponentMechanicsDispel,
		ratingengine.ComponentMechanicsDebuffUptime:
		return true
	default:
		return false
	}
}

// percentileAdapter is the ratingengine.PercentileSource logs/engine/rating.Score reads
// through, backed by rating_percentile_digests and kill_duration_digests. Fold controls
// whether a Placement call may also add this fight's value into its digest: true on a
// fight's first, genuine write; false on a rewrite (a re-verified fight) or a backfill
// recompute, so a value already folded once is never folded twice (spec's idempotency
// requirement). KillTimeBand never folds regardless of Fold - see foldKillDuration, called
// once per fight rather than once per roster player. Log is optional, following this
// codebase's Service.Log/logger() convention (api/internal/rankings/handler.go,
// api/internal/reports/ingest.go): nil is fine and falls back to slog.Default() via
// logger() below, so every construction of this type before Task 5 wires a real logger
// through keeps working unchanged.
type percentileAdapter struct {
	Tx   pgx.Tx
	Fold bool
	Log  *slog.Logger
}

var _ ratingengine.PercentileSource = (*percentileAdapter)(nil)

// logger returns a.Log, or slog.Default() when a.Log is unset.
func (a *percentileAdapter) logger() *slog.Logger {
	if a.Log != nil {
		return a.Log
	}
	return slog.Default()
}

// Placement reads bracket b's digest, answers this value's percentile within it (before
// this value is added - the same "read, then decide, then maybe fold" order
// rankings.Store.updateDigest already uses), and, when a.Fold and foldEligible(b), adds
// value and persists the row afterward, all under the caller's transaction. The returned
// pct/n are captured from the digest's state BEFORE the fold mutates it (Go passes them
// back by value, so foldValue's later mutation of the same *digest.Digest cannot change
// what was already returned) - a fight is never compared against its own not-yet-folded
// value.
func (a *percentileAdapter) Placement(b ratingengine.Bracket, value float64) (pct float64, n int64, ok bool) {
	var raw []byte
	err := a.Tx.QueryRow(context.Background(),
		`insert into rating_percentile_digests
		   (encounter_id, difficulty, spec, role, kill_time_band, kill, component, digest, n)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, 0)
		 on conflict (encounter_id, difficulty, spec, role, kill_time_band, kill, component)
		 do update set digest = rating_percentile_digests.digest
		 returning digest`,
		b.EncounterID, b.Difficulty, b.Spec, b.Role, b.KillTimeBand, b.Kill, b.Component, []byte{},
	).Scan(&raw)
	if err != nil {
		return 0, 0, false
	}
	d, derr := digest.Unmarshal(raw)
	if derr != nil {
		// A corrupted digest row falls back to a fresh, empty digest so this bracket can
		// keep accumulating - but that fallback must never be silent: if this call is also
		// fold-eligible, the fresh digest is about to be persisted over the corrupted row,
		// permanently discarding whatever history it held. Logged with the full bracket key
		// so the affected row can be found and, if the corruption's cause turns out to be
		// fixable, investigated before more folds compound the loss.
		a.logger().Error("rating", "op", "placement", "encounter_id", b.EncounterID,
			"difficulty", b.Difficulty, "spec", b.Spec, "role", b.Role,
			"kill_time_band", b.KillTimeBand, "kill", b.Kill, "component", b.Component, "err", derr)
		d = digest.New()
	}
	// digest.Digest.Placement already returns the share of the bracket's other values
	// this value beats in 0..1 (api/internal/digest/digest.go:203-212's own doc comment),
	// exactly the range ratingengine.PercentileSource.Placement documents - no further
	// scaling belongs here.
	if d.Count() > 0 {
		pct, n, ok = d.Placement(value), d.Count(), true
	}
	if a.Fold && foldEligible(b) {
		a.foldValue(b, d, value)
	}
	return pct, n, ok
}

// foldValue adds value to d and persists the encoded digest under bracket b. Called only
// after Placement has already captured its return values from d's pre-fold state.
func (a *percentileAdapter) foldValue(b ratingengine.Bracket, d *digest.Digest, value float64) {
	d.Add(value)
	encoded, err := d.MarshalBinary()
	if err != nil {
		// MarshalBinary never errors in the current implementation (digest.go:237-249
		// builds a byte slice with no fallible step); a future change that makes it
		// fallible must not silently drop a fold, so this panics rather than swallowing
		// the error.
		panic("rating: digest must always marshal: " + err.Error())
	}
	_, _ = a.Tx.Exec(context.Background(),
		`update rating_percentile_digests set digest = $8, n = $9, updated_at = now()
		 where encounter_id = $1 and difficulty = $2 and spec = $3 and role = $4
		   and kill_time_band = $5 and kill = $6 and component = $7`,
		b.EncounterID, b.Difficulty, b.Spec, b.Role, b.KillTimeBand, b.Kill, b.Component,
		encoded, d.Count())
}

// KillTimeBand reads (never folds) the kill-duration digest for (encounterID,
// difficulty) and classifies durationMS against its median, spec §1.2's exact bands.
func (a *percentileAdapter) KillTimeBand(encounterID, difficulty int64, durationMS int64) (string, bool) {
	var raw []byte
	err := a.Tx.QueryRow(context.Background(),
		`select digest from kill_duration_digests where encounter_id = $1 and difficulty = $2`,
		encounterID, difficulty).Scan(&raw)
	if err != nil {
		return "", false
	}
	d, err := digest.Unmarshal(raw)
	if err != nil || d.Count() == 0 {
		return "", false
	}
	median := d.Quantile(0.5)
	return classifyBand(float64(durationMS), median), true
}

func classifyBand(durationMS, medianMS float64) string {
	switch {
	case durationMS < killDurationFastRatio*medianMS:
		return ratingengine.BandFast
	case durationMS > killDurationSlowRatio*medianMS:
		return ratingengine.BandSlow
	default:
		return ratingengine.BandTypical
	}
}

// foldKillDuration folds one kill's duration into its (encounterID, difficulty) bracket
// exactly once. Called once per fight, before the per-player Score loop - never from
// inside percentileAdapter.KillTimeBand, which would otherwise fold the same fight's
// duration once per roster player.
func foldKillDuration(ctx context.Context, tx pgx.Tx, encounterID, difficulty int64, durationMS int64) error {
	var raw []byte
	if err := tx.QueryRow(ctx,
		`insert into kill_duration_digests (encounter_id, difficulty, digest, n)
		 values ($1, $2, $3, 0)
		 on conflict (encounter_id, difficulty) do update set digest = kill_duration_digests.digest
		 returning digest`,
		encounterID, difficulty, []byte{}).Scan(&raw); err != nil {
		return fmt.Errorf("rating: lock kill duration digest: %w", err)
	}
	d, err := digest.Unmarshal(raw)
	if err != nil {
		return fmt.Errorf("rating: decode kill duration digest: %w", err)
	}
	d.Add(float64(durationMS))
	encoded, err := d.MarshalBinary()
	if err != nil {
		return fmt.Errorf("rating: encode kill duration digest: %w", err)
	}
	if _, err := tx.Exec(ctx,
		`update kill_duration_digests set digest = $3, n = $4, updated_at = now()
		 where encounter_id = $1 and difficulty = $2`,
		encounterID, difficulty, encoded, d.Count()); err != nil {
		return fmt.Errorf("rating: write kill duration digest: %w", err)
	}
	return nil
}
