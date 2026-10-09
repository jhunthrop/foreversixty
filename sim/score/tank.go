package score

// The tank score, in effective-health-equivalent hit points:
//
//	score = effective health * (threat per second / 100) ^ 0.25 * exp(-0.5 * TMI / 100)
//
// where effective health is the hit points the boss's raw damage per
// second would take to kill the tank at the damage per second the tank
// actually takes: health * rawDPS / DTPS. The three terms are the order
// of the brief:
//
//   - Mitigation comes first and is a straight multiplier. Effective
//     health moves one for one with the health the set has and with the
//     inverse of the damage it lets through, so a point of stamina and a
//     point of armor are priced in the same coin, and a 1% change in
//     either is a 1% change in the score.
//   - TMI is the risk guard. The Theck-Meloree Index is the log-sum of
//     the damage the tank takes in sliding burst windows as a share of
//     its health, so it penalises what effective health cannot see: the
//     same damage per second taken in spikes is worth less than taken
//     steadily. Each TMI point costs half a percent of score.
//   - Threat is last and a quarter-power. A tank that cannot hold the
//     boss is not tanking, but past holding it more threat is worth far
//     less than staying alive: a 16% gain in threat per second is worth
//     a 4% gain in effective health.
//
// Chance of death is published, not scored: it is a rare event whose
// estimate at a few thousand iterations is noisier than the score's margins.

import (
	"errors"
	"fmt"
	"math"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	// TankThreatExponent is the power threat per second enters the score
	// at: a 16% gain in threat is worth a 4% gain in effective health.
	TankThreatExponent = 0.25
	// TankThreatReference is the threat per second that multiplies the
	// score by one. It only sets the score's magnitude.
	TankThreatReference = 100.0
	// TankRiskPerTMI is the share of score one TMI point costs.
	TankRiskPerTMI = 0.005
)

// TankRunResult is what a tank fight reports: what the tank took, how
// close it came to dying, the threat it made, its own damage, and the
// hit points that all of it is measured against.
type TankRunResult struct {
	DPS           api.Estimate
	DTPS          api.Estimate
	TPS           api.Estimate
	TMI           api.Estimate
	ChanceOfDeath float64
	// Health is the character's maximum health, from the engine's own
	// stat computation for the same request.
	Health float64
}

// TankRunOf reads a tank's metrics from a run of iterations.
func TankRunOf(player *proto.UnitMetrics, iterations int32, health float64) TankRunResult {
	return TankRunResult{
		DPS:           EstimateOf(player.GetDps(), iterations),
		DTPS:          EstimateOf(player.GetDtps(), iterations),
		TPS:           EstimateOf(player.GetThreat(), iterations),
		TMI:           EstimateOf(player.GetTmi(), iterations),
		ChanceOfDeath: player.GetChanceOfDeath(),
		Health:        health,
	}
}

// TankFigures is what one tank run measured and the score made of it.
type TankFigures struct {
	EffectiveHealth float64
	DTPS            float64
	TPS             float64
	TMI             float64
	DPS             float64
	ChanceOfDeath   float64
	// Score is the tank score and ScoreError its standard error, summed
	// linearly over the three measured terms (the upper bound: the three
	// are positively correlated within a run).
	Score      float64
	ScoreError float64
}

// TankScoreOf is the score of the three measured terms. threat is floored
// at one so a tank that made none scores low rather than scoring zero.
func TankScoreOf(effectiveHealth, tps, tmi float64) float64 {
	threat := math.Pow(math.Max(tps, 1)/TankThreatReference, TankThreatExponent)
	return effectiveHealth * threat * math.Exp(-TankRiskPerTMI*tmi)
}

// NewTankFigures reads one run against the boss's raw damage per second.
// A run in which the boss dealt nothing is not a tank fight and fails
// rather than scoring infinite effective health.
func NewTankFigures(run TankRunResult, bossRawDPS float64) (TankFigures, error) {
	if run.DTPS.Mean <= 0 {
		return TankFigures{}, errors.New("the boss dealt the tank no damage; the request is not a tank fight")
	}
	effectiveHealth := run.Health * bossRawDPS / run.DTPS.Mean
	score := TankScoreOf(effectiveHealth, run.TPS.Mean, run.TMI.Mean)
	relativeError := run.DTPS.Error/run.DTPS.Mean +
		TankThreatExponent*run.TPS.Error/math.Max(run.TPS.Mean, 1) +
		TankRiskPerTMI*run.TMI.Error
	return TankFigures{
		EffectiveHealth: effectiveHealth,
		DTPS:            run.DTPS.Mean,
		TPS:             run.TPS.Mean,
		TMI:             run.TMI.Mean,
		DPS:             run.DPS.Mean,
		ChanceOfDeath:   run.ChanceOfDeath,
		Score:           score,
		ScoreError:      score * relativeError,
	}, nil
}

// BossRawDPS is the curated boss's unmitigated damage per second for a
// character of level: the denominator of effective health.
func BossRawDPS(level int) (float64, error) {
	profile, err := request.TankProfileForLevel(level)
	if err != nil {
		return 0, err
	}
	return profile.Boss.RawDPS(), nil
}

// TankBlockOf measures a tank run for a character of level and returns the
// envelope's tank block.
func TankBlockOf(run TankRunResult, level int) (*api.TankResult, error) {
	rawDPS, err := BossRawDPS(level)
	if err != nil {
		return nil, err
	}
	figures, err := NewTankFigures(run, rawDPS)
	if err != nil {
		return nil, err
	}
	return &api.TankResult{
		DTPS:            run.DTPS,
		TPS:             run.TPS,
		TMI:             run.TMI,
		Health:          run.Health,
		EffectiveHealth: figures.EffectiveHealth,
		ChanceOfDeath:   run.ChanceOfDeath,
		Score:           api.Estimate{Mean: figures.Score, Error: figures.ScoreError},
	}, nil
}

// MaxHealth is the first player's maximum health as the engine computes
// it for raid: gear, talents, buffs and consumables included.
func MaxHealth(raid *proto.Raid, encounter *proto.Encounter) (float64, error) {
	result := core.ComputeStats(&proto.ComputeStatsRequest{Raid: raid, Encounter: encounter})
	for _, party := range result.GetRaidStats().GetParties() {
		for _, player := range party.GetPlayers() {
			final := player.GetFinalStats().GetStats()
			if int(stats.Health) < len(final) && final[stats.Health] > 0 {
				return final[stats.Health], nil
			}
		}
	}
	return 0, fmt.Errorf("the engine computed no maximum health for the character")
}
