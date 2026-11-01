package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

// SimRefIterations is how many iterations a reference run takes (about 3 s of compute).
const SimRefIterations = 1200

// SimRefCharacters builds the character a reference is measured on: the spec's band preset
// gear, with the preset and race the BiS ladder uses.
type SimRefCharacters interface {
	Character(ctx context.Context, build, spec string, band int) (simapi.CharacterSpec, error)
}

// SimRefDeps is what the coach-simref job needs.
type SimRefDeps struct {
	DB         *pgxpool.Pool
	Engine     runner.Runner
	Characters SimRefCharacters
	Tables     TableSource
	Now        func() time.Time // time.Now when nil
}

func (d SimRefDeps) now() time.Time {
	if d.Now != nil {
		return d.Now()
	}
	return time.Now()
}

// RunSimRef runs the spec's written rotation at the band for the bucket's length and stores
// the casts per minute and uptimes its Summary measures, in the units of a log Profile.
func RunSimRef(ctx context.Context, d SimRefDeps, k SimRefKey) error {
	char, err := d.Characters.Character(ctx, k.Build, k.Spec, k.Band)
	if err != nil {
		return fmt.Errorf("coach: sim ref character %s band %d: %w", k.Spec, k.Band, err)
	}
	res, err := d.Engine.Run(ctx, simRefRequest(k, char, d.now()), nil)
	if err != nil {
		return fmt.Errorf("coach: sim ref run %s: %w", k.Spec, err)
	}
	profile, err := d.profileOf(ctx, res.Summary)
	if err != nil {
		return err
	}
	raw, err := json.Marshal(simRefSummary{CastsPerMin: profile.CastsPerMin, Uptime: profile.Uptime})
	if err != nil {
		return fmt.Errorf("coach: encode sim ref: %w", err)
	}
	_, err = d.DB.Exec(ctx, `insert into coach_sim_refs (engine_version, build, spec, band, duration_bucket, summary)
		values ($1, $2, $3, $4, $5, $6)
		on conflict (engine_version, build, spec, band, duration_bucket)
		do update set summary = excluded.summary, computed_at = now()`,
		k.EngineVersion, k.Build, k.Spec, k.Band, k.DurationBucket, raw)
	if err != nil {
		return fmt.Errorf("coach: store sim ref: %w", err)
	}
	return nil
}

func simRefRequest(k SimRefKey, char simapi.CharacterSpec, at time.Time) simapi.SimRequest {
	enc := simapi.DefaultEncounter()
	enc.DurationSec = k.DurationBucket
	// The published numbers' target (sim/leveling RankerEncounter): no creature type. The
	// empty type is the engine's Humanoid, which credits Murder and slaying gear.
	enc.TargetType = simapi.TargetTypeUnknown
	return simapi.SimRequest{
		EngineVersion: k.EngineVersion, Spec: k.Spec, Iterations: SimRefIterations,
		Source: simapi.CharacterSource{
			Kind: simapi.SourceBuild, Ref: fmt.Sprintf("%s:%s:%d", k.Build, k.Spec, k.Band),
			CapturedAt: at.UTC().Format(time.RFC3339),
		},
		Character: char, Encounter: enc,
	}
}

// profileOf measures the sim's player with the same builder the logs use.
func (d SimRefDeps) profileOf(ctx context.Context, s summary.Summary) (Profile, error) {
	row, ok := simPlayer(s)
	if !ok {
		return Profile{}, errors.New("coach: the sim summary has no damage-dealing player")
	}
	tables, err := d.Tables.TablesFor(ctx, row, 0, d.now())
	if err != nil {
		return Profile{}, fmt.Errorf("coach: sim ref tables: %w", err)
	}
	return BuildProfile(Input{Fight: s, Row: row, Tables: tables}), nil
}

func simPlayer(s summary.Summary) (summary.RosterRow, bool) {
	for _, r := range s.Roster {
		if r.Role == RoleDPS {
			return r, true
		}
	}
	return summary.RosterRow{}, false
}
