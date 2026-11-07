package sims

// The auto-sim on import (design/specs/2026-11-07-home-b.md, "Auto-sim on import"): when the
// Battle.net import or an addon sync brings gear that differs from the character's last
// import-simmed gear, one server run of that gear and talents, the bare preset (no buffs, no
// consumables), the default iterations and encounter. Every account gets it, free included,
// so it is bounded twice: one run per character per gear change, and ImportSimDailyCap runs
// a day per account. A plain run (no stat weights) is the cheapest server run there is:
// defaultIterations against the default encounter, one engine pass, well inside the
// server lane's caps.

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/activity"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// ImportSimDailyCap is how many import runs one account starts in a day.
const ImportSimDailyCap = 12

// What RequestFromImport did.
const (
	ImportQueued    = "queued"
	ImportUnchanged = "unchanged"
	ImportCapped    = "capped"
	ImportNoGear    = "no_gear"
)

// ImportPreset is the preset an import run is simmed under.
const ImportPreset = "bare"

// importPlan is one import run, ready to queue.
type importPlan struct {
	req      simapi.SimRequest
	spec     string
	band     int
	gearHash string
}

// RequestFromImport queues the import run for characterKey (a character of userID's) when
// its gear changed since the last one. Anything about the character that keeps it from
// running is an answer, not an error; only the database and the job can fail it.
func (s *Service) RequestFromImport(ctx context.Context, userID int64, characterKey string) (string, error) {
	if s.Jobs == nil {
		return RequestUnavailable, nil
	}
	plan, ok, err := s.importPlan(ctx, userID, characterKey)
	if err != nil || !ok {
		return ImportNoGear, err
	}
	if err := plan.req.ValidateLane(simapi.LaneServer); err != nil {
		s.logger().Warn("sims", "op", "import sim", "character", characterKey, "err", err)
		return RequestInvalid, nil
	}
	id := auth.Base32ID(auth.ReportIDChars)
	outcome, err := s.Store.queueImport(ctx, id, userID, characterKey, plan)
	if err != nil || outcome != ImportQueued {
		return outcome, err
	}
	if err := s.Jobs.Run(ctx, SimRunJobCommand, id); err != nil {
		s.logger().Error("sims", "op", "import sim job", "sim", id, "err", err)
		failCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), failCompensationTimeout)
		defer cancel()
		if err := s.Store.Fail(failCtx, id, "the job could not be started"); err != nil {
			return "", err
		}
		return RequestNotStarted, nil
	}
	return ImportQueued, nil
}

func (s *Service) importPlan(ctx context.Context, userID int64, key string) (importPlan, bool, error) {
	var code, name, source string
	var captured time.Time
	err := s.Store.Pool.QueryRow(ctx,
		`select export, name, source, captured_at from addon_exports where character_key = $1 and user_id = $2`,
		key, userID).Scan(&code, &name, &source, &captured)
	if errors.Is(err, pgx.ErrNoRows) {
		return importPlan{}, false, nil
	}
	if err != nil {
		return importPlan{}, false, fmt.Errorf("sims: read export %s: %w", key, err)
	}
	char, err := characterFromCode(name, code)
	if err != nil || len(char.Gear) == 0 {
		return importPlan{}, false, nil
	}
	input, _, ok, err := s.Store.SimInput(ctx, key)
	if err != nil || !ok || input.Spec == "" {
		return importPlan{}, false, err
	}
	kind := simapi.SourceAddon
	if source == simapi.SourceBlizzard {
		kind = simapi.SourceBlizzard
	}
	raw, err := json.Marshal(char)
	if err != nil {
		return importPlan{}, false, fmt.Errorf("sims: hash gear %s: %w", key, err)
	}
	sum := sha256.Sum256(raw)
	return importPlan{
		req: simapi.SimRequest{
			EngineVersion: s.EngineVersion, Spec: input.Spec, Iterations: defaultIterations,
			Source:    simapi.CharacterSource{Kind: kind, Ref: key, CapturedAt: captured.UTC().Format(time.RFC3339)},
			Character: char, Encounter: defaultEncounterOf(0),
		},
		spec: input.Spec, band: activity.BandOf(char.Level), gearHash: hex.EncodeToString(sum[:]),
	}, true, nil
}

// queueImport writes the queued sims row and its character_sims row unless the character's
// last import run had this gear, or the account is at its daily cap; one transaction under
// a per-character lock, so two syncs racing queue one run.
func (s *Store) queueImport(ctx context.Context, id string, userID int64, key string, p importPlan) (string, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("sims: queue import %s: %w", key, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx, `select pg_advisory_xact_lock(hashtext('import_sim:' || $1))`, key); err != nil {
		return "", fmt.Errorf("sims: lock import %s: %w", key, err)
	}
	var same bool
	var today int
	if err := tx.QueryRow(ctx,
		`select coalesce((select gear_hash = $2 from character_sims where character_key = $1 and kind = 'import'
		                  order by created_at desc limit 1), false),
		        (select count(*) from character_sims where user_id = $3 and kind = 'import'
		                  and created_at > now() - interval '1 day')`, key, p.gearHash, userID).Scan(&same, &today); err != nil {
		return "", fmt.Errorf("sims: read imports %s: %w", key, err)
	}
	if same {
		return ImportUnchanged, nil
	}
	if today >= ImportSimDailyCap {
		return ImportCapped, nil
	}
	if err := queueRow(ctx, tx, id, userID, p.req); err != nil {
		return "", err
	}
	if err := recordCharacterSim(ctx, tx, id, key, userID, "import", p.spec, p.band, ImportPreset, p.gearHash); err != nil {
		return "", err
	}
	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("sims: queue import %s: commit: %w", key, err)
	}
	return ImportQueued, nil
}

// recordCharacterSim links a sim to the character it simmed, with what the guild ladder
// compares it on.
func recordCharacterSim(ctx context.Context, db execer, simID, key string, userID int64, kind, spec string, band int, preset, gearHash string) error {
	role := "dps"
	if sp, ok := specs.ByKey[spec]; ok && sp.Role != "" {
		role = sp.Role
	}
	if _, err := db.Exec(ctx,
		`insert into character_sims (sim_id, character_key, user_id, kind, spec, role, band, preset, gear_hash)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9) on conflict (sim_id) do nothing`,
		simID, key, userID, kind, spec, role, band, preset, gearHash); err != nil {
		return fmt.Errorf("sims: record character sim %s: %w", simID, err)
	}
	return nil
}

// presetOf is the preset a run compares under on the guild ladder: bare with no buffs and no
// consumables, raid with any.
func presetOf(req simapi.SimRequest) string {
	if len(req.Character.Buffs) == 0 && len(req.Character.Consumes) == 0 {
		return ImportPreset
	}
	return "raid"
}
