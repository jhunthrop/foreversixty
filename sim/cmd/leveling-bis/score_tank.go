package main

// Tank mode.
//
// A spec whose role in data/curated/specs.json is "tank" is ranked on a
// tank score rather than on damage. Every comparison the ranker makes
// (the stat weights, a trinket's tournament, a runner-up against a pick,
// a set bonus) asks the same question of a sim: which of these two sets
// is the better tank against the curated boss (data/curated/tank-
// encounter.json)? The tank runner below answers it with one number.
//
// The score, its three terms and their reasons are sim/score's: the
// ranker, inproc and the browser's wasm all call that one copy.
//
// Stat weights are the derivative of this score with respect to each
// stat, taken from the engine's own sweep (which reports the derivative
// of damage taken per second, threat per second and TMI) plus the
// maximum-health change each stat makes, which is deterministic and read
// from the engine's stat computation rather than simmed.

import (
	"fmt"
	"math"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/statid"
	"github.com/jhunthrop/foreversixty/sim/request"
	simscore "github.com/jhunthrop/foreversixty/sim/score"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	googleproto "google.golang.org/protobuf/proto"
)

const (
	// scoreUnitTankScore is bandReport.ScoreUnit for a tank band: every
	// sim_dps, dps_delta and per-point figure on such a band is in tank
	// score, never in damage.
	scoreUnitTankScore = "tank_score"

	// tankMetricsIterations and tankMetricsSeed are the final run that
	// publishes a band's figures. More iterations than a verify run
	// because the figures are quoted, not compared.
	tankMetricsIterations = 3000
	tankMetricsSeed       = 5

	// healthSweepStep is the stat nudge the health derivative is read
	// from. Maximum health is linear in every stat, so any step reads the
	// same slope; a large one keeps the engine's rounding out of it.
	healthSweepStep = 10.0
)

// tankMetricsReport is bandReport.Metrics: the figures the site shows in
// place of a DPS number on a tank band.
type tankMetricsReport struct {
	DTPS            float64 `json:"dtps"`
	TMI             float64 `json:"tmi"`
	ChanceOfDeath   float64 `json:"chance_of_death"`
	TPS             float64 `json:"tps"`
	EffectiveHealth float64 `json:"effective_health"`
	DPS             float64 `json:"dps"`
}

func tankReportOf(f simscore.TankFigures) *tankMetricsReport {
	return &tankMetricsReport{
		DTPS:            f.DTPS,
		TMI:             f.TMI,
		ChanceOfDeath:   f.ChanceOfDeath,
		TPS:             f.TPS,
		EffectiveHealth: f.EffectiveHealth,
		DPS:             f.DPS,
	}
}

// tankRunner is the engine as a tank sees it. It is a separate interface
// from engineRunner so that the tank-only steps (the published figures)
// can ask for what only a tank run measures.
type tankRunner interface {
	RunTankFigures(req api.SimRequest) (simscore.TankFigures, error)
}

// tankEngine ranks on tank score: it answers engineRunner's plain-DPS and
// weights questions with the score and its derivatives, and delegates the
// rest (hit profile) to the engine it wraps.
type tankEngine struct {
	engineRunner
}

// roleRunner is the one place the ranker's role is read: a tank spec is
// run by tankEngine, every other spec by the runner it was given.
func roleRunner(runner engineRunner, spec specInfo) engineRunner {
	if spec.Role != roleTank {
		return runner
	}
	return tankEngine{engineRunner: runner}
}

// roleTank is data/curated/specs.json's role string for a tank.
const roleTank = request.RoleTank

func (tankEngine) RunTankFigures(req api.SimRequest) (simscore.TankFigures, error) {
	rawDPS, err := simscore.BossRawDPS(req.Character.Level)
	if err != nil {
		return simscore.TankFigures{}, err
	}
	run, err := inproc.TankRun(req)
	if err != nil {
		return simscore.TankFigures{}, err
	}
	return simscore.NewTankFigures(run, rawDPS)
}

func (t tankEngine) RunPlainDPS(req api.SimRequest) (float64, error) {
	score, _, err := t.RunPlainDPSWithError(req)
	return score, err
}

func (t tankEngine) RunPlainDPSWithError(req api.SimRequest) (mean, stdErr float64, err error) {
	figures, err := t.RunTankFigures(req)
	if err != nil {
		return 0, 0, err
	}
	return figures.Score, figures.ScoreError, nil
}

// statSweep is the engine's sweep of one stat: the slope of each tank
// measure per point, and each slope's standard error.
type statSweep struct {
	dtps, tps, tmi          float64
	dtpsErr, tpsErr, tmiErr float64
	health                  float64
}

// tankWeightsFromSweeps turns per-stat sweeps and the baseline run into
// the published weights: the derivative of the score with respect to each
// stat, normalised so reference weighs exactly 1. The second return is
// the reference stat's own derivative, tank score per point.
func tankWeightsFromSweeps(ids []string, reference string, base simscore.TankFigures, baseHealth float64, sweeps map[string]statSweep) ([]api.StatWeight, float64, error) {
	raw := make(map[string]float64, len(ids))
	rawError := make(map[string]float64, len(ids))
	for _, id := range ids {
		s := sweeps[id]
		relative := s.health/baseHealth - s.dtps/base.DTPS +
			simscore.TankThreatExponent*s.tps/math.Max(base.TPS, 1) -
			simscore.TankRiskPerTMI*s.tmi
		relativeError := s.dtpsErr/base.DTPS +
			simscore.TankThreatExponent*s.tpsErr/math.Max(base.TPS, 1) +
			simscore.TankRiskPerTMI*s.tmiErr
		raw[id] = base.Score * relative
		rawError[id] = base.Score * relativeError
	}
	scale := raw[reference]
	if scale <= 0 {
		return nil, 0, fmt.Errorf("the reference stat %q did not raise the tank score (%.4f per point), so nothing can be normalised against it", reference, scale)
	}
	out := make([]api.StatWeight, 0, len(ids))
	for _, id := range ids {
		weight := raw[id] / scale
		err := rawError[id] / scale
		out = append(out, api.StatWeight{Stat: id, Weight: weight, Error: err, Insignificant: err >= math.Abs(weight)})
	}
	return out, scale, nil
}

// RunWeights sweeps the stats on the tank fight and publishes the score's
// derivatives. The engine's sweep supplies the slope of damage taken,
// threat and TMI; the health slope is read from the engine's own stat
// computation; the baseline is one plain run of the same request.
func (t tankEngine) RunWeights(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
	inproc.Register()
	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return nil, 0, fmt.Errorf("building the weights request: %w", err)
	}
	if err := simdb.AttachWeights(engineReq); err != nil {
		return nil, 0, fmt.Errorf("attaching the item database: %w", err)
	}
	baseline := req
	baseline.Weights = nil
	baseline.Iterations = req.Iterations * api.WeightsIterationsFactor
	baseRun, err := inproc.TankRun(baseline)
	if err != nil {
		return nil, 0, fmt.Errorf("the baseline run: %w", err)
	}
	rawDPS, err := simscore.BossRawDPS(req.Character.Level)
	if err != nil {
		return nil, 0, err
	}
	base, err := simscore.NewTankFigures(baseRun, rawDPS)
	if err != nil {
		return nil, 0, fmt.Errorf("the baseline run: %w", err)
	}

	res := core.StatWeights(engineReq)
	if res.GetError().GetMessage() != "" {
		return nil, 0, fmt.Errorf("the engine's weights sweep failed: %s", res.GetError().GetMessage())
	}
	healthSlopes, err := healthSlopesOf(engineReq, req.Weights.Stats)
	if err != nil {
		return nil, 0, err
	}
	sweeps, err := sweepsOf(res, req, healthSlopes)
	if err != nil {
		return nil, 0, err
	}
	rows, perPoint, err := tankWeightsFromSweeps(req.Weights.Stats, req.Weights.Reference, base, baseRun.Health, sweeps)
	if err != nil {
		return nil, 0, err
	}
	out := make(map[string]api.StatWeight, len(rows))
	for _, row := range rows {
		out[row.Stat] = row
	}
	return out, perPoint, nil
}

// sweepsOf reads the engine's weights result into one statSweep per
// requested stat. The engine's standard deviation is the population
// spread of the per-iteration deltas; dividing by the square root of the
// sample count behind it makes it a standard error, the way adapter.Weights
// does for the damage weights.
func sweepsOf(res *proto.StatWeightsResult, req api.SimRequest, healthSlopes map[string]float64) (map[string]statSweep, error) {
	sampleCount := float64(req.Iterations * api.WeightsIterationsFactor)
	if sampleCount <= 0 {
		return nil, fmt.Errorf("the request has no iterations to build standard errors from")
	}
	at := func(values []float64, s proto.Stat) float64 {
		if int(s) >= len(values) {
			return 0
		}
		return values[s]
	}
	standardError := func(stdev float64) float64 { return stdev / math.Sqrt(sampleCount) }
	sweeps := make(map[string]statSweep, len(req.Weights.Stats))
	for _, id := range req.Weights.Stats {
		s, ok := statid.Parse(id)
		if !ok {
			return nil, fmt.Errorf("%q is not a known stat id", id)
		}
		sweeps[id] = statSweep{
			dtps:    at(res.GetDtps().GetWeights().GetStats(), s),
			tps:     at(res.GetTps().GetWeights().GetStats(), s),
			tmi:     at(res.GetTmi().GetWeights().GetStats(), s),
			dtpsErr: standardError(at(res.GetDtps().GetWeightsStdev().GetStats(), s)),
			tpsErr:  standardError(at(res.GetTps().GetWeightsStdev().GetStats(), s)),
			tmiErr:  standardError(at(res.GetTmi().GetWeightsStdev().GetStats(), s)),
			health:  healthSlopes[id],
		}
	}
	return sweeps, nil
}

// healthSlopesOf is how many maximum health points one point of each stat
// buys, read from the engine's own stat computation for the sweep's
// character with the stat nudged. It is exact (maximum health is linear in
// every stat), so no standard error rides on it.
func healthSlopesOf(req *proto.StatWeightsRequest, ids []string) (map[string]float64, error) {
	healthOf := func(bonus *proto.UnitStats) (float64, error) {
		player := cloneWithBonus(req.Player, bonus)
		raid := core.SinglePlayerRaidProto(player, req.PartyBuffs, req.RaidBuffs, req.Debuffs)
		raid.Tanks = req.Tanks
		return simscore.MaxHealth(raid, req.Encounter)
	}
	base, err := healthOf(nil)
	if err != nil {
		return nil, err
	}
	slopes := make(map[string]float64, len(ids))
	for _, id := range ids {
		s, ok := statid.Parse(id)
		if !ok {
			return nil, fmt.Errorf("%q is not a known stat id", id)
		}
		nudge := &proto.UnitStats{Stats: make([]float64, stats.Len), PseudoStats: make([]float64, stats.PseudoStatsLen)}
		stats.UnitStatFromStat(stats.Stat(s)).AddToStatsProto(nudge, healthSweepStep)
		nudged, err := healthOf(nudge)
		if err != nil {
			return nil, err
		}
		slopes[id] = (nudged - base) / healthSweepStep
	}
	return slopes, nil
}

// cloneWithBonus is player with bonus added to its bonus stats; the
// original is left alone. A nil bonus is a plain copy.
func cloneWithBonus(player *proto.Player, bonus *proto.UnitStats) *proto.Player {
	clone := googleproto.Clone(player).(*proto.Player)
	if bonus == nil {
		return clone
	}
	if clone.BonusStats == nil {
		clone.BonusStats = &proto.UnitStats{}
	}
	if clone.BonusStats.Stats == nil {
		clone.BonusStats.Stats = make([]float64, stats.Len)
	}
	for i, v := range bonus.Stats {
		clone.BonusStats.Stats[i] += v
	}
	return clone
}
