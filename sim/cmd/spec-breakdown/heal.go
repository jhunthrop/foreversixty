package main

// The -heal report: a healer is judged by what it heals and whether its
// mana lasts, so this prints the two sides of the mana bar (income by
// source, spend by spell) and the healing each spell did, from the same
// request and the same curated profile the ranker runs the healer under.

import (
	"errors"
	"fmt"
	"math"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/score"
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
	{stats.Crit, "crit (every table)"},
	{stats.SpellCrit, "spell crit (spell table only)"},
}

// manaRow is one mana flow: a source that paid in or a spell that paid out.
type manaRow struct {
	Name   string
	Casts  float64
	Mana   float64
	Healed float64
}

func runHealReport(o options) (string, error) {
	req, err := legalRequest(o)
	if err != nil {
		return "", err
	}
	if specs.ByKey[req.Spec].Role != healerRole {
		return "", fmt.Errorf("-heal needs a healer spec, got %q", req.Spec)
	}
	fight, seconds, err := healFightFor(o)
	if err != nil {
		return "", err
	}
	req.Encounter.DurationSec = seconds
	if o.duration > 0 {
		req.Encounter.DurationSec = o.duration
	}
	req.Encounter.Variation = 0
	inproc.Register()
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: true})
	if err != nil {
		return "", fmt.Errorf("building the request: %w", err)
	}
	fight.Attach(engineReq)
	engineReq.Encounter.Duration = float64(req.Encounter.DurationSec)
	if err := simdb.Attach(engineReq); err != nil {
		return "", fmt.Errorf("attaching the item database: %w", err)
	}
	forced, err := parseForcedRanks(o.rank)
	if err != nil {
		return "", err
	}
	rotation, err := rotationFor(o, forced, specs.ByKey[o.spec].ClassSlug, req.Character.Race, req.Character.Level)
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

// healFightFor is the fight -heal runs: the BiS profile, or with -raid-fight the raid-shaped
// one, and its length.
func healFightFor(o options) (request.HealFight, int, error) {
	curated := filepath.Join(o.repoRoot, "data", "curated")
	if o.raidFight {
		fight, err := request.LoadRaidHealProfile(filepath.Join(curated, request.RaidHealProfileFile))
		return fight, fight.DurationSec, err
	}
	profile, err := request.LoadHealProfile(filepath.Join(curated, healFile))
	return profile, profile.DurationSec, err
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
	fmt.Fprintf(&b, "\n| Metric | Value |\n|---|---|\n| effective HPS | %.2f +/- %.2f |\n| raw HPS | %.1f |\n| mana lasts (s) | %.1f |\n",
		p.GetEffectiveHps().GetAvg(), meanError(p.GetEffectiveHps(), iters), p.GetHps().GetAvg(), p.GetTto().GetAvg())
	if final := firstFinalStats(computed); final != nil {
		result := score.HealingResultOf(p, res.IterationsDone, secs).WithManaPool(final.GetStats()[stats.Mana])
		fmt.Fprintf(&b, "| overheal | %.1f%% |\n| mana left at the end | %.0f of %.0f (%.1f%%) |\n",
			100*result.OverhealShare(), result.ManaLeft, result.ManaPool, 100*result.ManaLeftShare())
	}

	income, spend := splitManaFlows(p, names, iters)
	writeManaTable(&b, "Mana income", "per fight", "per second", income, secs, false)
	writeManaTable(&b, "Mana spend", "per fight", "per second", spend, secs, true)
	writeFreeCasts(&b, p, names, iters, spend)
	writeHealSources(&b, p, names, iters)
	writeHealCrits(&b, p, names, iters)
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

// writeHealSources lists every action that healed or shielded with what it
// landed per fight: healing and shields as cast (shields at the amount
// applied), and the effective part (healing that raised a health bar,
// shields that soaked damage). A shield's gap between the two is what
// expired or was overwritten unused.
func writeHealSources(b *strings.Builder, p *proto.UnitMetrics, names map[int32]string, iters float64) {
	b.WriteString("\n| Healing source | healed or shielded per fight | of it shielding | effective per fight | effective share |\n|---|---|---|---|---|\n")
	var total float64
	for _, a := range p.GetActions() {
		for _, t := range a.GetTargets() {
			total += t.GetEffectiveHealing()
		}
	}
	for _, a := range p.GetActions() {
		var raw, shielding, effective float64
		for _, t := range a.GetTargets() {
			raw += t.GetHealing() + t.GetShielding()
			shielding += t.GetShielding()
			effective += t.GetEffectiveHealing()
		}
		if raw == 0 {
			continue
		}
		fmt.Fprintf(b, "| %s | %.0f | %.0f | %.0f | %s |\n", actionName(a.GetId(), names), raw/iters, shielding/iters, effective/iters, percentOf(effective, total))
	}
}

// writeHealCrits lists every action that healed with its direct landings
// and ticks per fight and the share of each that crit, so a reader can see
// which table a heal rolls and whether its ticks crit. The engine counts a
// landing as a hit or a crit, never both, and a tick likewise.
func writeHealCrits(b *strings.Builder, p *proto.UnitMetrics, names map[int32]string, iters float64) {
	b.WriteString("\n| Heal crits | landed per fight | crit % | ticks per fight | tick crit % |\n|---|---|---|---|---|\n")
	for _, a := range p.GetActions() {
		var healing, hits, crits, ticks, critTicks float64
		for _, t := range a.GetTargets() {
			healing += t.GetHealing()
			hits += float64(t.GetHits())
			crits += float64(t.GetCrits())
			ticks += float64(t.GetTicks())
			critTicks += float64(t.GetCritTicks())
		}
		if healing == 0 {
			continue
		}
		landed, ticked := hits+crits, ticks+critTicks
		fmt.Fprintf(b, "| %s | %.1f | %s | %.1f | %s |\n", actionName(a.GetId(), names), landed/iters, percentOf(crits, landed), ticked/iters, percentOf(critTicks, ticked))
	}
}

// percentOf is part as a percentage of whole to one decimal, or "-" when
// there is no whole.
func percentOf(part, whole float64) string {
	if whole == 0 {
		return "-"
	}
	return fmt.Sprintf("%.1f", 100*part/whole)
}

// meanError is the standard error of a per-iteration metric's mean: its stdev over the
// square root of the iterations, the same error the DPS line reports.
func meanError(d *proto.DistributionMetrics, iterations float64) float64 {
	if iterations <= 0 {
		return 0
	}
	return d.GetStdev() / math.Sqrt(iterations)
}
