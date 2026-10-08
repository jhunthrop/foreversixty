package main

// Healer ranking.
//
// A healer is ranked on effective healing per second under one stated
// incoming-damage profile (data/curated/heal-profile.json), never on
// damage. The ranker's tournaments, verification runs and stat-weight
// sweeps all go through an engineRunner that speaks damage; healEngine is
// the engineRunner that speaks healing, so none of that code changes:
//
//   - RunPlainDPS and RunPlainDPSWithError return the set's GUARDED
//     effective healing per second (healingScore).
//   - RunWeights returns weights per point of effective healing per
//     second, normalised to the spec's reference stat (healing power).
//
// The guard is mana longevity. A healer that runs out of mana before the
// fight ends has stopped healing for the rest of it, and the average over
// the fight already pays for that, but a set that front-loads its mana
// into a burst and then idles should not out-rank a set that lasts, so a
// set that empties early is scored by the share of the fight it lasted,
// squared (healingScore). A set whose mana lasts the fight is scored by
// its effective healing per second alone.
//
// The profile is retuned by one rule, so the ranking keeps its signal.
// Once healing meets the incoming damage, more healing power changes
// nothing and gear ties. So the tank hit size and the raid-wide pulse
// damage are scaled TOGETHER (their ratio and cadence kept) by the
// smallest multiplier at which, with the five healers at band 60: the
// strongest raid-ready set covers about 80 percent of the incoming
// damage and the weakest about 60; every bare set (no raid consumables)
// covers less than half; and mana stays a live constraint (mana_lasts_sec
// under the fight length for every bare set and for at least one raid
// set). The multipliers are swept, the table is in
// design/reviews/2026-10-08-heal-profile-retune.md, and the next retune
// (Phase 2 gear) follows the same rule. The profile is a stand-in sized
// by that rule, never a named boss.
//
// Every published healer number carries the profile's name: the band
// entry names it (healerFields.Profile) and the file carries the profile
// itself, reasons included (specReport.HealProfile).

import (
	"fmt"
	"math"
	"path/filepath"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/statid"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core"
)

const (
	// healerRole is data/curated/specs.json's role for a healer.
	healerRole = "healer"
	// healProfileFile is the curated profile under data/curated.
	healProfileFile = "heal-profile.json"
	// healMetricsIterations is how many iterations the published metrics
	// of a band's final set run: more than a verification run, because
	// these are the numbers a player reads.
	healMetricsIterations = 1000
	// healMetricsSeed is the seed of that run, so a report is reproducible.
	healMetricsSeed = 11
	// maxManaLastsSec caps the published "mana lasts" figure; the engine
	// projects a time past the fight for a set that never ran out, and an
	// hour is where its own convention stops.
	maxManaLastsSec = 3600.0
)

// An item states its healing as "healing" and a spec's weights state it as
// "healing_power" (the engine's stat id), so a candidate's healing is
// weighed by the healing_power row.
const (
	itemHealingStat   = "healing"
	healingWeightStat = "healing_power"
)

// isHealer reports whether spec is ranked as a healer.
func isHealer(spec specInfo) bool { return spec.Role == healerRole }

// healingMetrics is the contract's metrics block: what the healer's final
// set did against the profile.
type healingMetrics struct {
	// HPS is effective healing per second: what landed.
	HPS float64 `json:"hps"`
	// RawHPS includes the overheal.
	RawHPS float64 `json:"raw_hps"`
	// OverhealPct is the share of raw healing that overhealed, 0 to 1.
	OverhealPct float64 `json:"overheal_pct"`
	// ManaLastsSec is seconds until the first out-of-mana, capped at an hour.
	ManaLastsSec float64 `json:"mana_lasts_sec"`
	// HPM is effective healing per point of mana spent.
	HPM float64 `json:"hpm"`
}

// healerFields are the band-entry keys a healer adds to the DPS shape. A
// non-healer's entry carries none of them.
type healerFields struct {
	Role    string          `json:"role,omitempty"`
	Profile string          `json:"profile,omitempty"`
	Metrics *healingMetrics `json:"metrics,omitempty"`
}

// healingBackend is where a healing run and a healing weights sweep come
// from: the engine in production, a fake in a test.
type healingBackend interface {
	Run(req api.SimRequest, profile request.HealProfile) (inproc.HealingResult, error)
	Weights(req api.SimRequest, profile request.HealProfile) (map[string]api.StatWeight, float64, error)
}

// healEngine is the engineRunner of a healer ranked under one profile.
type healEngine struct {
	backend healingBackend
	profile request.HealProfile
}

// withHealerEngine swaps in the healing engine for a healer spec and loads
// its profile; any other spec keeps the runner it came with.
func withHealerEngine(runner engineRunner, repoRoot string, spec specInfo) (engineRunner, *request.HealProfile, error) {
	if !isHealer(spec) {
		return runner, nil, nil
	}
	profile, err := request.LoadHealProfile(filepath.Join(repoRoot, "data", "curated", healProfileFile))
	if err != nil {
		return nil, nil, fmt.Errorf("loading the heal profile for %s: %w", spec.Spec, err)
	}
	return healEngine{backend: realHealingBackend{}, profile: profile}, &profile, nil
}

// forProfile fixes the fight to the profile's: a healer is ranked over one
// stated fight, whatever length the request carried.
func (h healEngine) forProfile(req api.SimRequest) api.SimRequest {
	req.Encounter.DurationSec = h.profile.DurationSec
	req.Encounter.Variation = 0
	return req
}

func (h healEngine) measure(req api.SimRequest) (inproc.HealingResult, error) {
	return h.backend.Run(h.forProfile(req), h.profile)
}

// RunPlainDPS is the set's guarded effective healing per second.
func (h healEngine) RunPlainDPS(req api.SimRequest) (float64, error) {
	score, _, err := h.RunPlainDPSWithError(req)
	return score, err
}

// RunPlainDPSWithError is RunPlainDPS with the score's standard error,
// scaled by the same guard so a comparison keeps its signal-to-noise.
func (h healEngine) RunPlainDPSWithError(req api.SimRequest) (float64, float64, error) {
	result, err := h.measure(req)
	if err != nil {
		return 0, 0, err
	}
	factor := lastingFactor(result.ManaLastsSec, h.profile.Duration())
	return result.Effective.Mean * factor, result.Effective.Error * factor, nil
}

// RunWeights sweeps the spec's weight stats over the profile's fake raid.
func (h healEngine) RunWeights(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
	return h.backend.Weights(h.forProfile(req), h.profile)
}

// HitProfileFor is empty: a healer has no miss table to cap against.
func (healEngine) HitProfileFor(api.SimRequest) (core.HitProfile, error) {
	return core.HitProfile{}, nil
}

// lastingFactor is the share of the fight a set's mana lasted, squared,
// capped at one: 1 for a set that never runs dry, 0.64 for one that runs
// out at four fifths of the fight.
func lastingFactor(manaLastsSec float64, fight time.Duration) float64 {
	if fight <= 0 || manaLastsSec >= fight.Seconds() {
		return 1
	}
	share := math.Max(manaLastsSec, 0) / fight.Seconds()
	return share * share
}

// healingScore is the number a healing set ranks on: effective healing per
// second, scaled by lastingFactor.
func healingScore(result inproc.HealingResult, fight time.Duration) float64 {
	return result.Effective.Mean * lastingFactor(result.ManaLastsSec, fight)
}

// metricsFor is the contract's metrics block for one run.
func metricsFor(result inproc.HealingResult) healingMetrics {
	return healingMetrics{
		HPS:          result.Effective.Mean,
		RawHPS:       result.Raw.Mean,
		OverhealPct:  result.OverhealShare(),
		ManaLastsSec: math.Min(result.ManaLastsSec, maxManaLastsSec),
		HPM:          result.HealingPerMana(),
	}
}

// healerFieldsFor measures the band's final set under the profile and
// returns the band entry's healer keys. set_dps is the effective healing
// per second, so nothing that reads set_dps breaks.
func (h healEngine) healerFieldsFor(spec specInfo, race, classSlug string, band int, talents string, picks map[string]slotPick) (healerFields, float64, error) {
	req := plainRequest(spec, bandCharacter("heal-metrics", race, classSlug, spec.Spec, band, talents, buildGear(picks)), healMetricsIterations, healMetricsSeed)
	result, err := h.measure(req)
	if err != nil {
		return healerFields{}, 0, err
	}
	metrics := metricsFor(result)
	return healerFields{Role: healerRole, Profile: h.profile.ID, Metrics: &metrics}, metrics.HPS, nil
}

// healFieldsAttacher is what a band report needs of a healer's engine; a
// damage engine does not have it.
type healFieldsAttacher interface {
	healerFieldsFor(spec specInfo, race, classSlug string, band int, talents string, picks map[string]slotPick) (healerFields, float64, error)
}

// attachHealerFields puts a healer's keys on its band report; a damage
// spec's runner is not a healFieldsAttacher and the report is left alone.
func attachHealerFields(report *bandReport, runner engineRunner, spec specInfo, race, classSlug string, band int, talents string, picks map[string]slotPick) error {
	attacher, ok := runner.(healFieldsAttacher)
	if !ok {
		return nil
	}
	fields, hps, err := attacher.healerFieldsFor(spec, race, classSlug, band, talents, picks)
	if err != nil {
		return fmt.Errorf("measuring the healer's set: %w", err)
	}
	report.Role = fields.Role
	report.Profile = fields.Profile
	report.Metrics = fields.Metrics
	report.SetDPS = hps
	return nil
}

// realHealingBackend runs the engine in this process.
type realHealingBackend struct{}

func (realHealingBackend) Run(req api.SimRequest, profile request.HealProfile) (inproc.HealingResult, error) {
	return inproc.HealingRun(req, profile)
}

// Weights is runWeights for a healer: the same request, the profile laid
// over it, and the engine's healing table read instead of its damage one.
func (realHealingBackend) Weights(req api.SimRequest, profile request.HealProfile) (map[string]api.StatWeight, float64, error) {
	registerEngine()
	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return nil, 0, fmt.Errorf("building the weights request: %w", err)
	}
	profile.AttachWeights(engineReq)
	if err := simdb.AttachWeights(engineReq); err != nil {
		return nil, 0, fmt.Errorf("attaching the item database: %w", err)
	}
	res := core.StatWeights(engineReq)
	weights, err := adapter.HealingWeights(res, req)
	if err != nil {
		return nil, 0, fmt.Errorf("reading the engine's healing weights: %w", err)
	}
	out := make(map[string]api.StatWeight, len(weights))
	for _, w := range weights {
		out[w.Stat] = w
	}
	reference, ok := statid.Parse(req.Weights.Reference)
	if !ok {
		return nil, 0, fmt.Errorf("reference stat %q is not a known stat id", req.Weights.Reference)
	}
	raw := res.GetHps().GetWeights().GetStats()
	if int(reference) >= len(raw) {
		return out, 0, nil
	}
	return out, raw[reference], nil
}
