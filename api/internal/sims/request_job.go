package sims

// A sim the player asked for in game. The addon writes the request into its
// outbox (docs/contracts/addon-v2.md, contract 4) and the companion posts it
// with the character's export; this turns it into one premium server run,
// the same queued sims row and the same `api sim-run` job POST /v1/sims/run
// starts (Run reads it back off the row), built from the character's
// sim-input with the request's own gear code.

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// AddonSimRequest is contract 4's sim_request.
type AddonSimRequest struct {
	// Character is Name-Realm, the addon's own name for the character.
	Character string `json:"character"`
	Spec      string `json:"spec"`
	// RequestedAt is the player's local wall time, RequestedAtLayout. It
	// has no zone, so it is kept as the text the addon wrote: it names
	// the request, it is not a moment the server can compare.
	RequestedAt string `json:"requested_at"`
	// GearCode is the character's FS1 export as of the request.
	GearCode string `json:"gear_code"`
}

// RequestedAtLayout is contract 4's YYYY-MM-DD HH:MM.
const RequestedAtLayout = "2006-01-02 15:04"

// SimRequestWindow is how often one character's requests may start a run:
// a second request inside it, or the same request at any time, is a
// duplicate. The companion re-posts the outbox on every file change, so
// the identical request arrives many times.
const SimRequestWindow = time.Hour

// What RequestFromAddon did with a request.
const (
	RequestQueued           = "queued"
	RequestDuplicate        = "duplicate"
	RequestPremiumRequired  = "premium_required"
	RequestUnavailable      = "unavailable"
	RequestUnknownCharacter = "unknown_character"
	RequestInvalid          = "invalid"
	RequestNotStarted       = "not_started"
)

// baseLevel is the planner's BASE_LEVEL: a character with no talent points
// is level 9, and every point is one level on (web/src/lib/planner/types.ts).
const baseLevel = 9

// Validate checks the request's shape: a spec the site models, a request
// time in the contract's layout, and a gear code that is an FS1 export.
func (r AddonSimRequest) Validate() error {
	if strings.TrimSpace(r.Character) == "" {
		return errors.New("sim_request.character is required")
	}
	if _, ok := specs.ByKey[r.Spec]; !ok {
		return fmt.Errorf("sim_request.spec %q is not a spec", r.Spec)
	}
	if _, err := time.Parse(RequestedAtLayout, r.RequestedAt); err != nil {
		return errors.New("sim_request.requested_at must be YYYY-MM-DD HH:MM")
	}
	if _, ok := fs1.Decode(r.GearCode); !ok {
		return errors.New("sim_request.gear_code must be an FS1 export")
	}
	return nil
}

// RequestFromAddon queues one server run for characterKey (a character of
// userID's own; the caller resolves it) and starts the job. It answers what
// it did rather than failing for anything about the request itself: the
// export that carried it is already stored, and a request that cannot run
// must not fail the sync. Only a database failure is an error.
func (s *Service) RequestFromAddon(ctx context.Context, userID int64, characterKey string, in AddonSimRequest) (string, error) {
	if s.Jobs == nil || s.Accounts == nil {
		// POST /v1/sims/run is not mounted without both either.
		return RequestUnavailable, nil
	}
	allowed, _, err := s.Accounts.Can(ctx, userID, entitlements.FeatureServerSims)
	if err != nil {
		return "", fmt.Errorf("sims: check server sims for %d: %w", userID, err)
	}
	if !allowed {
		return RequestPremiumRequired, nil
	}
	_, ref, ok, err := s.Store.SimInput(ctx, characterKey)
	if err != nil {
		return "", err
	}
	if !ok {
		return RequestUnknownCharacter, nil
	}
	name, _ := character.SplitUnit(in.Character)
	char, err := characterFromCode(name, in.GearCode)
	if err != nil {
		s.logger().Warn("sims", "op", "addon request", "character", characterKey, "err", err)
		return RequestInvalid, nil
	}
	char.Buffs = s.recordedBuffs(ctx, ref)
	req := simapi.SimRequest{
		EngineVersion: s.EngineVersion, Spec: in.Spec, Iterations: defaultIterations,
		Source: simapi.CharacterSource{
			Kind: simapi.SourceAddon, Ref: characterKey, CapturedAt: in.RequestedAt,
		},
		Character: char,
		Encounter: withEncounterDefaults(simapi.EncounterSpec{}),
	}
	if err := req.ValidateLane(simapi.LaneServer); err != nil {
		s.logger().Warn("sims", "op", "addon request", "character", characterKey, "err", err)
		return RequestInvalid, nil
	}

	id := auth.Base32ID(auth.ReportIDChars)
	queued, err := s.Store.QueueAddonRequest(ctx, id, userID, characterKey, in.RequestedAt, req)
	if err != nil {
		return "", err
	}
	if !queued {
		return RequestDuplicate, nil
	}
	if err := s.Jobs.Run(ctx, SimRunJobCommand, id); err != nil {
		// The same compensation POST /v1/sims/run makes: the row says
		// the run failed rather than sitting on "queued" forever. The
		// request stays recorded, so the outbox's next re-post does not
		// start it again.
		s.logger().Error("sims", "id", httpx.RequestIDFrom(ctx), "op", "addon request job",
			"sim", id, "err", err)
		failCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), failCompensationTimeout)
		defer cancel()
		if err := s.Store.Fail(failCtx, id, "the job could not be started"); err != nil {
			return "", err
		}
		return RequestNotStarted, nil
	}
	return RequestQueued, nil
}

// QueueAddonRequest writes the queued sims row for an addon request and
// records the request against the character, unless the same request was
// already recorded or another was recorded inside SimRequestWindow; it
// reports whether it queued. The check and both writes are one
// transaction under a per-character lock, so two posts of the same outbox
// racing each other queue one run.
func (s *Store) QueueAddonRequest(ctx context.Context, id string, userID int64, characterKey, requestedAt string, req simapi.SimRequest) (bool, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("sims: queue addon request %s: begin: %w", characterKey, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx,
		`select pg_advisory_xact_lock(hashtext('addon_sim_requests:' || $1))`, characterKey); err != nil {
		return false, fmt.Errorf("sims: lock addon requests for %s: %w", characterKey, err)
	}
	var seen bool
	if err := tx.QueryRow(ctx,
		`select exists (select 1 from addon_sim_requests
		  where character_key = $1
		    and (requested_at = $2 or created_at > now() - make_interval(secs => $3)))`,
		characterKey, requestedAt, SimRequestWindow.Seconds()).Scan(&seen); err != nil {
		return false, fmt.Errorf("sims: read addon requests for %s: %w", characterKey, err)
	}
	if seen {
		return false, nil
	}
	if err := queueRow(ctx, tx, id, userID, req); err != nil {
		return false, err
	}
	if _, err := tx.Exec(ctx,
		`insert into addon_sim_requests (character_key, requested_at, user_id, sim_id)
		 values ($1, $2, $3, $4)`, characterKey, requestedAt, userID, id); err != nil {
		return false, fmt.Errorf("sims: record addon request for %s: %w", characterKey, err)
	}
	if err := tx.Commit(ctx); err != nil {
		return false, fmt.Errorf("sims: queue addon request %s: commit: %w", characterKey, err)
	}
	return true, nil
}

// characterFromCode turns an FS1 export into the engine's character: the
// web's characterFromFs1 and toCharacterSpec (web/src/lib/sim/character.ts)
// for the fields an export carries. Buffs and consumables are empty; the
// caller adds what the sim-input recorded.
func characterFromCode(name, code string) (simapi.CharacterSpec, error) {
	d, ok := fs1.Decode(code)
	if !ok {
		return simapi.CharacterSpec{}, errors.New("sims: the gear code is not an FS1 export")
	}
	trees := make([]string, len(d.Talents.Trees))
	points := 0
	for i, tree := range d.Talents.Trees {
		var b strings.Builder
		for _, digit := range tree {
			rank, err := strconv.ParseInt(string(digit), 36, 0)
			if err != nil {
				return simapi.CharacterSpec{}, fmt.Errorf("sims: talent digit %q: %w", digit, err)
			}
			// talentsString's own clamp: one decimal digit per talent.
			b.WriteByte(byte('0' + min(rank, 9)))
		}
		trees[i] = strings.TrimRight(b.String(), "0")
		points += d.Talents.Points[i]
	}
	level := d.Level
	if !d.HasLevel {
		level = min(simapi.MaxLevel, baseLevel+points)
	}
	slots := make([]string, 0, len(d.Gear))
	for slot, itemID := range d.Gear {
		if itemID > 0 {
			slots = append(slots, slot)
		}
	}
	sort.Strings(slots)
	gear := make([]simapi.GearSlot, len(slots))
	for i, slot := range slots {
		gear[i] = simapi.GearSlot{Slot: slot, ItemID: d.Gear[slot], Enchant: d.Enchants[slot]}
	}
	out := simapi.CharacterSpec{
		Name: name, Race: d.RaceSlug, Class: d.ClassSlug, Level: level,
		Talents: strings.Join(trees, "-"), Gear: gear,
		Buffs: []string{}, Consumes: []string{},
	}
	if len(d.Professions) > 0 {
		out.Profession = d.Professions
	}
	return out, nil
}
