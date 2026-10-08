package main

// The -heal report: a healer is judged by what it heals and whether its
// mana lasts, so this prints the two sides of the mana bar (income by
// source, spend by spell) and the healing each spell did, from the same
// request and the same curated profile the ranker runs the healer under.

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

// healFile is the curated profile under data/curated.
// healerRole is data/curated/specs.json's role for a healer.
const healerRole = "healer"

const healFile = "heal-profile.json"

// healStatNames are the final stats the healer report prints, in order.
var healStatNames = []struct {
	stat stats.Stat
	name string
}{
	{stats.Mana, "mana pool"},
	{stats.Intellect, "intellect"},
	{stats.Spirit, "spirit"},
	{stats.MP5, "mp5"},
	{stats.HealingPower, "healing power"},
	{stats.SpellPower, "spell power"},
	{stats.Crit, "crit"},
}

// manaRow is one mana flow: a source that paid in or a spell that paid out.
type manaRow struct {
	Name   string
	Casts  float64
	Mana   float64
	Healed float64
}

func runHealReport(o options) (string, error) {
	req, err := buildRequest(o)
	if err != nil {
		return "", err
	}
	if specs.ByKey[req.Spec].Role != healerRole {
		return "", fmt.Errorf("-heal needs a healer spec, got %q", req.Spec)
	}
	profile, err := request.LoadHealProfile(filepath.Join(o.repoRoot, "data", "curated", healFile))
	if err != nil {
		return "", err
	}
	req.Encounter.DurationSec = profile.DurationSec
	if o.duration > 0 {
		req.Encounter.DurationSec = o.duration
	}
	req.Encounter.Variation = 0
	inproc.Register()
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: true})
	if err != nil {
		return "", fmt.Errorf("building the request: %w", err)
	}
	profile.Attach(engineReq)
	engineReq.Encounter.Duration = float64(req.Encounter.DurationSec)
	if err := simdb.Attach(engineReq); err != nil {
		return "", fmt.Errorf("attaching the item database: %w", err)
	}
	rotation, err := loadRotation(o.rotation)
	if err != nil {
		return "", err
	}
	if rotation != nil {
		engineReq.GetRaid().GetParties()[0].GetPlayers()[0].Rotation = rotation
	}
	computed := core.ComputeStats(&proto.ComputeStatsRequest{Raid: engineReq.GetRaid(), Encounter: engineReq.GetEncounter()})
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(engineReq, reporter, "spec-breakdown-heal")
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
	return healMarkdown(req, computed, res, player, names), nil
}

func healMarkdown(req api.SimRequest, computed *proto.ComputeStatsResult, res *proto.RaidSimResult, p *proto.UnitMetrics, names map[int32]string) string {
	var b strings.Builder
	iters := float64(res.IterationsDone)
	secs := float64(req.Encounter.DurationSec)
	fmt.Fprintf(&b, "%s healer fight, %d iterations, %.0f s\n\nconsumes: %v\nbuffs: %v\n\n| Stat before the fight's auras | Value |\n|---|---|\n", req.Spec, res.IterationsDone, secs, req.Character.Consumes, req.Character.Buffs)
	if final := firstFinalStats(computed); final != nil {
		for _, s := range healStatNames {
			fmt.Fprintf(&b, "| %s | %.2f |\n", s.name, final.GetStats()[s.stat])
		}
	}
	fmt.Fprintf(&b, "\n| Metric | Value |\n|---|---|\n| effective HPS | %.1f |\n| raw HPS | %.1f |\n| mana lasts (s) | %.1f |\n",
		p.GetEffectiveHps().GetAvg(), p.GetHps().GetAvg(), p.GetTto().GetAvg())

	income, spend := splitManaFlows(p, names, iters)
	writeManaTable(&b, "Mana income", "per fight", "per second", income, secs, false)
	writeManaTable(&b, "Mana spend", "per fight", "per second", spend, secs, true)
	writeFreeCasts(&b, p, names, iters, spend)
	return b.String()
}

// splitManaFlows folds the player's mana resources into income and spend
// rows, each per iteration, with the healing each spell landed.
func splitManaFlows(p *proto.UnitMetrics, names map[int32]string, iters float64) (income, spend []manaRow) {
	healed := map[string]float64{}
	casts := map[string]float64{}
	for _, a := range p.GetActions() {
		name := actionName(a.GetId(), names)
		for _, t := range a.GetTargets() {
			healed[name] += t.GetEffectiveHealing()
			casts[name] += float64(t.GetCasts())
		}
	}
	for _, r := range p.GetResources() {
		if r.GetType() != proto.ResourceType_ResourceTypeMana {
			continue
		}
		name := actionName(r.GetId(), names)
		if r.GetGain() < 0 {
			spend = append(spend, manaRow{Name: name, Casts: casts[name] / iters, Mana: -r.GetGain() / iters, Healed: healed[name] / iters})
			continue
		}
		income = append(income, manaRow{Name: name, Casts: float64(r.GetEvents()) / iters, Mana: r.GetActualGain() / iters})
	}
	byMana := func(rows []manaRow) {
		sort.Slice(rows, func(i, j int) bool { return rows[i].Mana > rows[j].Mana })
	}
	byMana(income)
	byMana(spend)
	return income, spend
}

func writeManaTable(b *strings.Builder, title, perFight, perSecond string, rows []manaRow, secs float64, withHealing bool) {
	head := fmt.Sprintf("\n| %s | events | mana %s | mana %s |", title, perFight, perSecond)
	rule := "\n|---|---|---|---|"
	if withHealing {
		head += " healed per fight | healed per mana |"
		rule += "---|---|"
	}
	b.WriteString(head + rule + "\n")
	var total float64
	for _, r := range rows {
		total += r.Mana
		fmt.Fprintf(b, "| %s | %.1f | %.0f | %.2f |", r.Name, r.Casts, r.Mana, r.Mana/secs)
		if withHealing {
			per := 0.0
			if r.Mana > 0 {
				per = r.Healed / r.Mana
			}
			fmt.Fprintf(b, " %.0f | %.2f |", r.Healed, per)
		}
		b.WriteString("\n")
	}
	fmt.Fprintf(b, "| total | | %.0f | %.2f |", total, total/secs)
	if withHealing {
		b.WriteString(" | |")
	}
	b.WriteString("\n")
}

// writeFreeCasts lists the actions that cast without paying mana (a
// cooldown, a free cast, a totem the engine charges nothing for), which
// the spend table cannot show.
func writeFreeCasts(b *strings.Builder, p *proto.UnitMetrics, names map[int32]string, iters float64, spend []manaRow) {
	paid := map[string]bool{}
	for _, r := range spend {
		paid[r.Name] = true
	}
	b.WriteString("\n| Cast without a mana row | casts per fight | healed per fight |\n|---|---|---|\n")
	for _, a := range p.GetActions() {
		name := actionName(a.GetId(), names)
		var casts, healed float64
		for _, t := range a.GetTargets() {
			casts += float64(t.GetCasts())
			healed += t.GetEffectiveHealing()
		}
		if casts == 0 || paid[name] {
			continue
		}
		fmt.Fprintf(b, "| %s | %.1f | %.0f |\n", name, casts/iters, healed/iters)
	}
}
