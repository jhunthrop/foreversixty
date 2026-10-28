package coach

import (
	"context"
	"fmt"

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

// Resolver answers an Input's Refs: the top logs first, then the sim for a validated DPS
// spec whose pool is too thin. An unvalidated spec is never simmed and never "preparing".
type Resolver struct {
	Pool          *Pool
	Sims          *SimRefs
	Validated     SpecValidator
	EngineVersion string
	Build         string // the newest client build whose manifest fetched_at <= the fight (Ruling F11, approximate)
}

// Refs resolves the references for playerKey, who played k.Spec at level for durationMS.
func (r *Resolver) Refs(ctx context.Context, k PoolKey, playerKey, role string, level int, durationMS int64) (Refs, error) {
	count, pool, err := r.Pool.Reference(ctx, k, playerKey)
	if err != nil {
		return Refs{}, fmt.Errorf("coach: pool reference: %w", err)
	}
	refs := Refs{Pool: pool, PoolCount: count}
	if pool != nil || role != RoleDPS || LevelBand(level) == 0 {
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
	key := SimRefKey{EngineVersion: r.EngineVersion, Build: r.Build, Spec: slug,
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
