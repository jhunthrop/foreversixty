package coach

import (
	"context"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

// simBands are the level bands the BiS ladder ranks (leveling-bis defaultBandsFlag).
var simBands = []int{20, 30, 40, 50, 60}

// LevelBand is the highest ladder band at or below level; 0 when none (level unknown or
// under the lowest band).
func LevelBand(level int) int {
	band := 0
	for _, b := range simBands {
		if level >= b {
			band = b
		}
	}
	return band
}

// SpecValidator says whether a spec's sim is validated. *sims.Store satisfies it.
type SpecValidator interface {
	Validated(ctx context.Context, spec string) (bool, error)
}

// Resolver answers an Input's Refs: the top logs, and beside them the sim for a validated
// DPS spec, so that comparison can take the first basis that exists for each measurement
// (§11). An unvalidated spec is never simmed and never "preparing".
type Resolver struct {
	Pool          *Pool
	Sims          *SimRefs
	Validated     SpecValidator
	EngineVersion string
	// BuildAt is the build in force at a fight: the newest client build whose manifest
	// fetched_at is at or before the fight (Ruling F11, an approximation).
	BuildAt func(fought time.Time) string
}

// Refs resolves the references for playerKey, who played k.Spec at level for durationMS in the fight at fought.
func (r *Resolver) Refs(ctx context.Context, k PoolKey, playerKey, role string, level int, durationMS int64, fought time.Time) (Refs, error) {
	count, pool, err := r.Pool.Reference(ctx, k, playerKey)
	if err != nil {
		return Refs{}, fmt.Errorf("coach: pool reference: %w", err)
	}
	refs := Refs{Pool: pool, PoolCount: count}
	if role != RoleDPS {
		return refs, nil
	}
	slug, ok := dpsSpecSlug(k.Spec)
	if !ok {
		return refs, nil
	}
	validated, err := r.Validated.Validated(ctx, slug)
	if err != nil {
		return Refs{}, fmt.Errorf("coach: spec validated: %w", err)
	}
	if !validated {
		return refs, nil
	}
	if LevelBand(level) == 0 {
		refs.SimNoLevel = true
		return refs, nil
	}
	key := SimRefKey{EngineVersion: r.EngineVersion, Build: r.BuildAt(fought), Spec: slug,
		Band: LevelBand(level), DurationBucket: DurationBucket(durationMS)}
	refs.Sim, refs.SimPreparing, err = r.Sims.Get(ctx, key)
	return refs, err
}

// dpsSpecSlug is the site's spec key ("warrior-fury") for a display name; DPS spec names
// are unique across classes.
func dpsSpecSlug(name string) (string, bool) {
	for _, s := range specs.All {
		if s.Role == "dps" && s.Name == name {
			return s.Spec, true
		}
	}
	return "", false
}
