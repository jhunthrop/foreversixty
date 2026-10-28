package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/jobs"
)

const (
	// SimRefJobCommand is the job that fills one coach_sim_refs row.
	SimRefJobCommand = "coach-simref"
	// DurationBucketSec is the fight length the sim reference is rounded to.
	DurationBucketSec = 30
	// SimRefRequeueAfter is how long a queued key is not queued again.
	SimRefRequeueAfter = time.Hour

	simRefKeyArgs = 5 // engine version, build, spec, band, duration bucket
)

// SimRefKey is one coach_sim_refs row's key. DurationBucket is seconds, rounded to 30.
type SimRefKey struct {
	EngineVersion, Build, Spec string
	Band, DurationBucket       int
}

// DurationBucket is the fight length in seconds rounded to the nearest 30, at least 30.
func DurationBucket(durationMS int64) int {
	step := int64(DurationBucketSec) * 1000
	return int(max(1, (durationMS+step/2)/step)) * DurationBucketSec
}

// Args is the key as the job's arguments after SimRefJobCommand.
func (k SimRefKey) Args() []string {
	return []string{k.EngineVersion, k.Build, k.Spec, strconv.Itoa(k.Band), strconv.Itoa(k.DurationBucket)}
}

// ParseSimRefKey is Args' inverse, for the job entry point.
func ParseSimRefKey(args []string) (SimRefKey, error) {
	if len(args) != simRefKeyArgs {
		return SimRefKey{}, fmt.Errorf("coach: sim ref key needs %d arguments, got %d", simRefKeyArgs, len(args))
	}
	band, err := strconv.Atoi(args[3])
	if err != nil {
		return SimRefKey{}, fmt.Errorf("coach: sim ref band %q: %w", args[3], err)
	}
	bucket, err := strconv.Atoi(args[4])
	if err != nil {
		return SimRefKey{}, fmt.Errorf("coach: sim ref duration bucket %q: %w", args[4], err)
	}
	return SimRefKey{EngineVersion: args[0], Build: args[1], Spec: args[2], Band: band, DurationBucket: bucket}, nil
}

// simRefSummary is coach_sim_refs.summary: the Profile's units (uptime is a percent 0-100).
type simRefSummary struct {
	CastsPerMin map[string]float64 `json:"cpm"`
	Uptime      map[string]float64 `json:"uptime"`
}

func (s simRefSummary) ref(band int) (*SimRef, error) {
	ref := &SimRef{Band: band, CastsPerMin: s.CastsPerMin, Uptime: map[int64]float64{}}
	for key, v := range s.Uptime {
		id, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("coach: sim ref uptime key %q: %w", key, err)
		}
		ref.Uptime[id] = v
	}
	return ref, nil
}

// SimRefs reads coach_sim_refs. A miss queues one coach-simref job, once per key per hour
// (in memory), and answers preparing.
type SimRefs struct {
	DB   *pgxpool.Pool
	Jobs jobs.Runner
	Now  func() time.Time // time.Now when nil
	Log  *slog.Logger     // slog.Default when nil

	mu     sync.Mutex
	queued map[SimRefKey]time.Time
}

func (s *SimRefs) now() time.Time {
	if s.Now != nil {
		return s.Now()
	}
	return time.Now()
}

// Get is the stored reference, or (nil, true) on a miss. A miss queues the job; a job that
// cannot be queued (no runner, a launch error) is logged and still answers preparing, which
// is never cached, so the next view retries rather than failing the pack.
func (s *SimRefs) Get(ctx context.Context, k SimRefKey) (*SimRef, bool, error) {
	var raw []byte
	err := s.DB.QueryRow(ctx, `select summary from coach_sim_refs
		where engine_version = $1 and build = $2 and spec = $3 and band = $4 and duration_bucket = $5`,
		k.EngineVersion, k.Build, k.Spec, k.Band, k.DurationBucket).Scan(&raw)
	if err == nil {
		ref, derr := decodeSimRef(raw, k.Band)
		return ref, false, derr
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, false, fmt.Errorf("coach: read sim ref: %w", err)
	}
	if err := s.queue(ctx, k); err != nil {
		s.logger().Warn("coach: sim ref not queued; the pack is preparing",
			slog.String("spec", k.Spec), slog.Int("band", k.Band), slog.String("error", err.Error()))
	}
	return nil, true, nil
}

func (s *SimRefs) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func decodeSimRef(raw []byte, band int) (*SimRef, error) {
	var sum simRefSummary
	if err := json.Unmarshal(raw, &sum); err != nil {
		return nil, fmt.Errorf("coach: decode sim ref: %w", err)
	}
	return sum.ref(band)
}

// queue runs the job unless this key was queued within SimRefRequeueAfter. The key is
// marked under the lock and the launch happens outside it, so one slow launch never stalls
// another key; a failed launch clears the mark.
func (s *SimRefs) queue(ctx context.Context, k SimRefKey) error {
	if s.Jobs == nil {
		return errors.New("coach: sim refs have no job runner")
	}
	if !s.claim(k) {
		return nil
	}
	if err := s.Jobs.Run(ctx, append([]string{SimRefJobCommand}, k.Args()...)...); err != nil {
		s.mu.Lock()
		delete(s.queued, k)
		s.mu.Unlock()
		return fmt.Errorf("coach: queue sim ref: %w", err)
	}
	return nil
}

// claim marks k queued now and reports whether the caller should launch it.
func (s *SimRefs) claim(k SimRefKey) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if at, ok := s.queued[k]; ok && s.now().Sub(at) < SimRefRequeueAfter {
		return false
	}
	if s.queued == nil {
		s.queued = map[SimRefKey]time.Time{}
	}
	s.queued[k] = s.now()
	return true
}
