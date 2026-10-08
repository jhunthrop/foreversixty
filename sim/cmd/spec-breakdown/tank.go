package main

// The -tank report: a tank is judged by what the boss does to it, not by
// what it does to the boss, so this prints the character's final defensive
// stats, the published tank figures and the boss's swing outcomes against
// it, from the same request the ranker builds (request.BuildWith applies the
// curated tank fight to a tank spec).

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

// tankStatNames are the final stats the report prints, in order.
var tankStatNames = []struct {
	stat stats.Stat
	name string
}{
	{stats.Health, "health"},
	{stats.Stamina, "stamina"},
	{stats.Strength, "strength"},
	{stats.Agility, "agility"},
	{stats.Armor, "armor"},
	{stats.BonusArmor, "bonus armor"},
	{stats.Defense, "defense"},
	{stats.Dodge, "dodge %"},
	{stats.Parry, "parry %"},
	{stats.Block, "block %"},
	{stats.BlockValue, "block value (items)"},
	{stats.Hit, "hit"},
	{stats.Crit, "crit"},
	{stats.AttackPower, "attack power"},
	{stats.Mana, "mana"},
}

// bossSwingRow is the boss's outcome tally against the tank, per swing.
type bossSwingRow struct {
	Swings                                      float64
	Miss, Dodge, Parry, Block, Crit, Crush, Hit float64
	DamagePerSwing, BlockedSwingsDamageShare    float64
}

func runTankReport(o options) (string, error) {
	req, err := buildRequest(o)
	if err != nil {
		return "", err
	}
	if !request.IsTankSpec(req.Spec) {
		return "", fmt.Errorf("-tank needs a tank spec, got %q", req.Spec)
	}
	inproc.Register()
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: true})
	if err != nil {
		return "", fmt.Errorf("building the request: %w", err)
	}
	if err := simdb.Attach(engineReq); err != nil {
		return "", fmt.Errorf("attaching the item database: %w", err)
	}
	rotation, err := loadRotation(o.rotation)
	if err != nil {
		return "", err
	}
	if rotation != nil {
		for _, party := range engineReq.GetRaid().GetParties() {
			for _, player := range party.GetPlayers() {
				player.Rotation = rotation
			}
		}
	}
	computed := core.ComputeStats(&proto.ComputeStatsRequest{Raid: engineReq.GetRaid(), Encounter: engineReq.GetEncounter()})
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(engineReq, reporter, "spec-breakdown-tank")
	res := simdrain.ToResult(reporter, nil)
	if res == nil {
		return "", errors.New("the engine produced no result")
	}
	if err := adapter.ResultError(res); err != nil {
		return "", err
	}
	player, err := adapter.PlayerMetrics(res)
	if err != nil {
		return "", err
	}
	names, err := loadSpellNames(o.repoRoot)
	if err != nil {
		return "", err
	}
	return tankMarkdown(req, computed, res, player, names), nil
}

func firstFinalStats(computed *proto.ComputeStatsResult) *proto.UnitStats {
	for _, party := range computed.GetRaidStats().GetParties() {
		for _, p := range party.GetPlayers() {
			return p.GetFinalStats()
		}
	}
	return nil
}

func bossSwings(res *proto.RaidSimResult, iterations float64) bossSwingRow {
	var row bossSwingRow
	var attempts, damage, blockedDamage float64
	for _, target := range res.GetEncounterMetrics().GetTargets() {
		for _, a := range target.GetActions() {
			for _, t := range a.GetTargets() {
				swings := float64(t.GetHits() + t.GetCrits() + t.GetMisses() + t.GetDodges() + t.GetParries() + t.GetBlocks() + t.GetCrushes())
				attempts += swings
				row.Miss += float64(t.GetMisses())
				row.Dodge += float64(t.GetDodges())
				row.Parry += float64(t.GetParries())
				row.Block += float64(t.GetBlocks())
				row.Crit += float64(t.GetCrits())
				row.Crush += float64(t.GetCrushes())
				row.Hit += float64(t.GetHits())
				damage += t.GetDamage()
				blockedDamage += t.GetBlockDamage()
			}
		}
	}
	if attempts == 0 {
		return row
	}
	row.Swings = attempts / iterations
	row.Miss, row.Dodge, row.Parry, row.Block = 100*row.Miss/attempts, 100*row.Dodge/attempts, 100*row.Parry/attempts, 100*row.Block/attempts
	row.Crit, row.Crush, row.Hit = 100*row.Crit/attempts, 100*row.Crush/attempts, 100*row.Hit/attempts
	row.DamagePerSwing = damage / attempts
	if damage+blockedDamage > 0 {
		row.BlockedSwingsDamageShare = 100 * blockedDamage / (damage + blockedDamage)
	}
	return row
}

func tankMarkdown(req api.SimRequest, computed *proto.ComputeStatsResult, res *proto.RaidSimResult, p *proto.UnitMetrics, names map[int32]string) string {
	var b strings.Builder
	iters := float64(res.IterationsDone)
	secs := float64(req.Encounter.DurationSec)
	fmt.Fprintf(&b, "%s tank fight, %d iterations\n\n| Stat before the fight's auras | Value |\n|---|---|\n", req.Spec, res.IterationsDone)
	if final := firstFinalStats(computed); final != nil {
		for _, s := range tankStatNames {
			fmt.Fprintf(&b, "| %s | %.2f |\n", s.name, final.GetStats()[s.stat])
		}
		pseudo := func(ps proto.PseudoStat) float64 { return final.GetPseudoStats()[ps] }
		fmt.Fprintf(&b, "| block value multiplier | %.2f |\n| block value per strength | %.3f |\n",
			pseudo(proto.PseudoStat_PseudoStatBlockValueMultiplier), pseudo(proto.PseudoStat_PseudoStatBlockValuePerStrength))
	}
	fmt.Fprintf(&b, "\n| Metric | Value |\n|---|---|\n| damage taken /s | %.1f |\n| TMI | %.1f |\n| chance of death | %.1f%% |\n| threat /s | %.1f |\n| own damage /s | %.1f |\n",
		p.GetDtps().GetAvg(), p.GetTmi().GetAvg(), 100*p.GetChanceOfDeath(), p.GetThreat().GetAvg(), p.GetDps().GetAvg())

	swing := bossSwings(res, iters)
	fmt.Fprintf(&b, "\n| Boss swings per fight | miss | dodge | parry | block | crit | crush | hit | damage per landed-or-not swing | share of damage taken on blocked swings |\n|---|---|---|---|---|---|---|---|---|---|\n| %.1f | %.1f%% | %.1f%% | %.1f%% | %.1f%% | %.1f%% | %.1f%% | %.1f%% | %.0f | %.1f%% |\n",
		swing.Swings, swing.Miss, swing.Dodge, swing.Parry, swing.Block, swing.Crit, swing.Crush, swing.Hit, swing.DamagePerSwing, swing.BlockedSwingsDamageShare)

	b.WriteString("\n| Tank action | casts/fight | shielding/s | threat/s |\n|---|---|---|---|\n")
	for _, a := range p.GetActions() {
		var casts, shield, threat float64
		for _, t := range a.GetTargets() {
			casts += float64(t.GetCasts())
			shield += t.GetShielding()
			threat += t.GetThreat()
		}
		if casts == 0 && shield == 0 {
			continue
		}
		fmt.Fprintf(&b, "| %s | %.2f | %.1f | %.1f |\n", actionName(a.GetId(), names), casts/iters, shield/iters/secs, threat/iters/secs)
	}
	b.WriteString("\n| Aura | Uptime % | Procs/iter |\n|---|---|---|\n")
	for _, a := range p.GetAuras() {
		if a.GetUptimeSecondsAvg() == 0 {
			continue
		}
		fmt.Fprintf(&b, "| %s | %.1f | %.2f |\n", actionName(a.GetId(), names), 100*a.GetUptimeSecondsAvg()/secs, a.GetProcsAvg())
	}
	return b.String()
}
