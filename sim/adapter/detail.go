package adapter

// The saved result page's detail (spec 2026-11-04-sim-results 11.2): what a
// player reads about one run beyond the summary's rows. Summarize keeps the
// logs engine's shape, and its goldens pin it; everything here is laid over
// its output by WithDetail, which both lanes call on a run the page will
// show. A weights request never reaches it, and sim/bulk takes it away again.

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

const (
	auraTypeDebuff = "DEBUFF"
	auraOnPlayer   = "player"
	auraOnTarget   = "target"

	// extraAttackBookkeepingKey is the engine's store of pending extra
	// attacks, a permanent counter aura it files under Thrash's id
	// (sim/core/attack.go StoreExtraMHAttack, sim/core/buffs.go
	// CreateExtraAttackAuraCommon). It is up all fight on every character a
	// Windfury Totem reaches and no player has such a buff, so it never
	// reaches a page.
	extraAttackBookkeepingKey = "spell:21919"

	// maxHistogramBins is the widest spread drawn bin by bin. The engine bins
	// at 10 DPS, which is 39 bins on a level 60 raid run; a spread this wide is
	// a far longer axis than a chart can draw, so the bins widen instead.
	maxHistogramBins = 200

	// amountDecimals is the precision of a per-fight average on the page.
	amountDecimals = 10
)

// WithDetail lays the page's detail over a finished run's result. The summary
// gains the target's debuff tracks and loses what the engine registered but no
// fight ever had. It returns a new result and leaves its argument alone.
func WithDetail(out api.SimResult, res *proto.RaidSimResult) (api.SimResult, error) {
	player, err := PlayerMetrics(res)
	if err != nil {
		return api.SimResult{}, err
	}
	iters := float64(res.IterationsDone)
	sample := res.GetSampleIteration()
	sampled := sampleAuras(sample)

	out.DPSHistogram = dpsHistogram(res.GetRaidMetrics().GetDps())
	out.ResourceSources = resourceSources(player, iters)
	out.ExtraAttackSources = extraAttackSources(player, iters)
	out.SampleDurationMS = int64(math.Round(sample.GetDurationSeconds() * 1000))
	out.SampleAuras = sampled
	out.FinalStats, out.TargetStats = sheetStats(sample.GetSheet(), out.Request.Spec)
	out.Summary.Auras = detailAuras(out.Summary.Auras, res.GetEncounterMetrics().GetTargets(), sampled)
	out.Summary.DamageDone = withoutBookkeepingRows(out.Summary.DamageDone)
	return out, nil
}

// withoutBookkeepingRows takes the engine's resource pseudo-actions (the rage
// gain ticks, which travel through the action metrics) out of the damage table.
// They dealt nothing and no player pressed them. isPlayerCast already keeps them
// off the casts and the sample; Summarize keeps them for its goldens' sake.
func withoutBookkeepingRows(actors []summary.Actor) []summary.Actor {
	out := make([]summary.Actor, len(actors))
	for i, actor := range actors {
		actor.Abilities = slices.DeleteFunc(slices.Clone(actor.Abilities), func(ab summary.Ability) bool {
			return ab.Total == 0 && strings.HasPrefix(ab.Name, otherActionKeyPrefix) && !isSwingKey(ab.Name)
		})
		out[i] = actor
	}
	return out
}

const otherActionKeyPrefix = "other:"

// isSwingKey says a key is a white swing or a shot, which a row keeps even when
// it dealt nothing: a fight's misses are still a fight's.
func isSwingKey(key string) bool {
	return strings.HasPrefix(key, "other:attack") || strings.HasPrefix(key, "other:shoot")
}

// sampleOutcome is a sample row's outcome word. A landed result with no damage
// is the engine resolving an application rather than a blow (Deep Wounds books
// a hit each time a crit refreshes it), and its damage comes as ticks, so the
// row reads as a cast with no outcome of its own.
func sampleOutcome(c *proto.SampleCast) string {
	if c.GetOutcome() == proto.SampleOutcome_SampleOutcomeHit && c.GetAmount() == 0 {
		return ""
	}
	return sampleOutcomeWords[c.GetOutcome()]
}

var sampleOutcomeWords = map[proto.SampleOutcome]string{
	proto.SampleOutcome_SampleOutcomeHit:      "hit",
	proto.SampleOutcome_SampleOutcomeCrit:     "crit",
	proto.SampleOutcome_SampleOutcomeMiss:     "miss",
	proto.SampleOutcome_SampleOutcomeDodge:    "dodge",
	proto.SampleOutcome_SampleOutcomeParry:    "parry",
	proto.SampleOutcome_SampleOutcomeGlance:   "glance",
	proto.SampleOutcome_SampleOutcomeBlock:    "block",
	proto.SampleOutcome_SampleOutcomeTick:     "tick",
	proto.SampleOutcome_SampleOutcomeTickCrit: "tick_crit",
	// The page's vocabulary has no crush: a crushing blow is a mob's, and a
	// row here is the player's. Should a player's ever crush, it landed.
	proto.SampleOutcome_SampleOutcomeCrush: "hit",
}

// dpsHistogram turns the engine's sparse map of rounded DPS into contiguous
// bins. The engine rounds to tens on both entry points
// (DistributionMetrics.doneIteration); the keys' own spacing is what says so.
func dpsHistogram(d *proto.DistributionMetrics) *api.DPSHistogram {
	hist := d.GetHist()
	if len(hist) == 0 {
		return nil
	}
	keys := make([]int, 0, len(hist))
	for k := range hist {
		keys = append(keys, int(k))
	}
	slices.Sort(keys)
	start, last := keys[0], keys[len(keys)-1]

	bin := histogramBaseBin
	for (last-start)/bin+1 > maxHistogramBins {
		bin *= 2
	}
	counts := make([]int, (last-start)/bin+1)
	for k, n := range hist {
		counts[(int(k)-start)/bin] += int(n)
	}
	return &api.DPSHistogram{Start: start, BinDPS: bin, Counts: counts}
}

// histogramBaseBin is the engine's own rounding: math.Round(dps/10)*10.
const histogramBaseBin = 10

// resourceNames is the resources a player reads about, by the engine's enum.
// Combo points are a count, and health here is a cost, not a bar anyone spends.
var resourceNames = map[proto.ResourceType]string{
	proto.ResourceType_ResourceTypeMana:   "mana",
	proto.ResourceType_ResourceTypeEnergy: "energy",
	proto.ResourceType_ResourceTypeRage:   "rage",
	proto.ResourceType_ResourceTypeFocus:  "focus",
}

// resourceSources is the engine's per-action ResourceMetrics, split the way
// resources() splits them: a row's sign is its direction. So Bloodrage, which
// pays in health and is booked in rage only as a gain, is a gain.
func resourceSources(u *proto.UnitMetrics, iters float64) []api.ResourceSources {
	byType := map[proto.ResourceType]*api.ResourceSources{}
	order := []proto.ResourceType{}
	for _, rm := range u.GetResources() {
		name, ok := resourceNames[rm.GetType()]
		if !ok {
			continue
		}
		bucket, seen := byType[rm.GetType()]
		if !seen {
			bucket = &api.ResourceSources{Resource: name, Gained: []api.ResourceSourceRow{}, Spent: []api.ResourceSourceRow{}}
			byType[rm.GetType()] = bucket
			order = append(order, rm.GetType())
		}
		_, action := ActionName(rm.GetId())
		if rm.GetActualGain() >= 0 {
			bucket.Gained = append(bucket.Gained, api.ResourceSourceRow{
				Action: action,
				Amount: OneDecimal(rm.GetActualGain() / iters),
				Wasted: OneDecimal(max(rm.GetGain()-rm.GetActualGain(), 0) / iters),
			})
		} else {
			bucket.Spent = append(bucket.Spent, api.ResourceSourceRow{
				Action: action,
				Amount: OneDecimal(-rm.GetActualGain() / iters),
			})
		}
	}
	out := make([]api.ResourceSources, 0, len(order))
	for _, t := range order {
		bucket := *byType[t]
		sortResourceRows(bucket.Gained)
		sortResourceRows(bucket.Spent)
		out = append(out, bucket)
	}
	return out
}

func sortResourceRows(rows []api.ResourceSourceRow) {
	slices.SortStableFunc(rows, func(a, b api.ResourceSourceRow) int {
		return cmp.Or(cmp.Compare(b.Amount, a.Amount), cmp.Compare(a.Action, b.Action))
	})
}

// OneDecimal rounds a per-fight average to the precision the page prints.
func OneDecimal(v float64) float64 {
	return math.Round(v*amountDecimals) / amountDecimals
}

// extraAttackSources says where the extra swings on the summary's own row came
// from, most first. The engine counts what each source granted, and a grant is
// not always a swing of its own: one that arrives on a normal swing's heels
// only speeds up the next swing (WeaponAttack.castExtraAttacksTriggered). So
// the sources share out the swings the summary reports under the extra attack
// tag, in proportion to their grants, and the page's two figures agree.
//
// Windfury Totem is named as the buff the request asks for; every other source
// is a proc, named by its action key.
func extraAttackSources(u *proto.UnitMetrics, iters float64) []api.ExtraAttackSource {
	rows := u.GetExtraAttacks()
	if len(rows) == 0 {
		return nil
	}
	var granted int64
	for _, row := range rows {
		granted += int64(row.GetAttacks())
	}
	share := 1 / iters
	if swings := extraAttackSwings(u); swings > 0 && granted > 0 {
		share = float64(swings) / float64(granted) / iters
	}
	out := make([]api.ExtraAttackSource, 0, len(rows))
	for _, row := range rows {
		out = append(out, api.ExtraAttackSource{
			Source: extraAttackSourceKey(row.GetId()),
			Count:  OneDecimal(float64(row.GetAttacks()) * share),
		})
	}
	slices.SortStableFunc(out, func(a, b api.ExtraAttackSource) int {
		return cmp.Or(cmp.Compare(b.Count, a.Count), cmp.Compare(a.Source, b.Source))
	})
	return out
}

// extraAttackTag is the engine's tag on an extra swing's metrics
// (sim/core/attack.go tagExtraAttack).
const extraAttackTag = 3

// extraAttackSwings is the casts booked under the extra attack tag on the
// engine's swing actions, summed over every iteration.
func extraAttackSwings(u *proto.UnitMetrics) int64 {
	var total int64
	for _, am := range u.GetActions() {
		other := am.GetId().GetOtherId()
		if am.GetId().GetTag() != extraAttackTag ||
			(other != proto.OtherAction_OtherActionAttack && other != proto.OtherAction_OtherActionShoot) {
			continue
		}
		for _, t := range am.GetTargets() {
			total += int64(t.GetCasts())
		}
	}
	return total
}

func extraAttackSourceKey(id *proto.ActionID) string {
	_, key := ActionName(id)
	return key
}

// sampleAuras folds the engine's gain, stack and fade events into per-aura
// windows on the sample run's own clock. Pre-pull time clamps to the pull, the
// way the engine's own uptime counts it. An aura that was never up in the fight
// has no row, and neither does the engine's extra attack counter.
func sampleAuras(sample *proto.SampleIteration) []api.SampleAura {
	if sample == nil {
		return nil
	}
	type key struct {
		action string
		target string
	}
	type track struct {
		aura    api.SampleAura
		openAt  int64
		open    bool
		stacked bool
	}
	tracks := map[key]*track{}
	order := []key{}
	for _, e := range sample.GetAuras() {
		_, action := ActionName(e.GetAuraId())
		if action == extraAttackBookkeepingKey {
			continue
		}
		k := key{action, e.GetTarget()}
		t, seen := tracks[k]
		if !seen {
			on := auraOnPlayer
			if e.GetTarget() != "" {
				on = auraOnTarget
			}
			t = &track{aura: api.SampleAura{Action: action, On: on, Segments: [][2]int64{}}}
			tracks[k] = t
			order = append(order, k)
		}
		at := max(e.GetAtMs(), 0)
		switch {
		case e.GetActive() && !t.open:
			t.open, t.openAt = true, at
		case !e.GetActive() && t.open:
			t.open = false
			if at > t.openAt {
				t.aura.Segments = append(t.aura.Segments, [2]int64{t.openAt, at})
			}
		}
		if e.GetStacks() > 0 {
			t.aura.Stacks = appendStack(t.aura.Stacks, at, int64(e.GetStacks()))
		}
	}
	out := make([]api.SampleAura, 0, len(order))
	for _, k := range order {
		if a := tracks[k].aura; len(a.Segments) > 0 {
			out = append(out, a)
		}
	}
	return out
}

// appendStack adds one change point: an aura that rises through several stacks
// inside one millisecond keeps the last, and a repeat of the current count is
// not a change.
func appendStack(points [][2]int64, at, stacks int64) [][2]int64 {
	if n := len(points); n > 0 {
		if points[n-1][1] == stacks {
			return points
		}
		if points[n-1][0] == at {
			return append(points[:n-1:n-1], [2]int64{at, stacks})
		}
	}
	return append(points, [2]int64{at, stacks})
}

// detailAuras is the summary's aura list as the page reads it: the player's
// buffs that were ever up, then the target's debuffs that were.
func detailAuras(rows []summary.AuraTrack, targets []*proto.UnitMetrics, sampled []api.SampleAura) []summary.AuraTrack {
	out := make([]summary.AuraTrack, 0, len(rows))
	for _, row := range rows {
		if row.Name == extraAttackBookkeepingKey || (row.Applications == 0 && row.UptimeMS == 0) {
			continue
		}
		out = append(out, row)
	}
	debuffs := debuffTracks(targets, sampled)
	return append(out, debuffs...)
}

// debuffTracks is the engine's target unit metrics as summary rows. Like the
// player's, these are per-iteration averages already. The engine reports no
// stack count, so a row's is the sample run's highest.
func debuffTracks(targets []*proto.UnitMetrics, sampled []api.SampleAura) []summary.AuraTrack {
	rows := []summary.AuraTrack{}
	for i, target := range targets {
		for _, am := range target.GetAuras() {
			spellID, name := ActionName(am.GetId())
			applications := int64(math.Round(am.GetProcsAvg()))
			uptime := int64(math.Round(am.GetUptimeSecondsAvg() * 1000))
			if applications == 0 && uptime == 0 {
				continue
			}
			rows = append(rows, summary.AuraTrack{
				TargetGUID:   fmt.Sprintf("target-%d", i+1),
				TargetName:   target.GetName(),
				SpellID:      spellID,
				Name:         name,
				Type:         auraTypeDebuff,
				Applications: applications,
				MaxStacks:    highestStack(sampled, name),
				UptimeMS:     uptime,
				Segments:     []summary.Segment{},
				Appliers:     []string{},
			})
		}
	}
	rows = foldAuras(rows)
	slices.SortStableFunc(rows, func(a, b summary.AuraTrack) int {
		return cmp.Or(cmp.Compare(b.UptimeMS, a.UptimeMS), strings.Compare(a.Name, b.Name))
	})
	return rows
}

func highestStack(sampled []api.SampleAura, action string) int64 {
	var most int64
	for _, a := range sampled {
		if a.Action != action || a.On != auraOnTarget {
			continue
		}
		for _, point := range a.Stacks {
			most = max(most, point[1])
		}
	}
	return most
}

// sheetStats lifts the engine's sheet into the envelope. A run with no sample
// has no sheet, and then no block: the page hides it rather than guessing.
func sheetStats(sheet *proto.SampleSheet, spec string) (*api.FinalStats, *api.TargetStats) {
	if sheet == nil {
		return nil, nil
	}
	final := &api.FinalStats{
		AttackPower:          int(math.Round(attackPowerFor(sheet, spec))),
		MeleeHitPct:          OneDecimal(sheet.GetMeleeHitPct()),
		MeleeCritPctVsTarget: OneDecimal(sheet.GetMeleeCritPctVsTarget()),
		MeleeHastePct:        OneDecimal(sheet.GetMeleeHastePct()),
		WeaponSkill: &api.WeaponSkill{
			MainHand: int(math.Round(sheet.GetMainHandSkill())),
			OffHand:  int(math.Round(sheet.GetOffHandSkill())),
			Ranged:   int(math.Round(sheet.GetRangedSkill())),
		},
	}
	if class, _ := splitSpecSlug(spec); sheet.GetHasManaBar() && class != hunterClass {
		power := int(math.Round(sheet.GetSpellPower()))
		hit := OneDecimal(sheet.GetSpellHitPct())
		crit := OneDecimal(sheet.GetSpellCritPct())
		final.SpellPower, final.SpellHitPct, final.SpellCritPct = &power, &hit, &crit
	}
	target := &api.TargetStats{
		Level:             int(sheet.GetTargetLevel()),
		Armor:             int(math.Round(sheet.GetTargetArmor())),
		ArmorAfterDebuffs: int(math.Round(sheet.GetTargetArmorAfterDebuffs())),
	}
	return final, target
}

const hunterClass = "hunter"

// meleeHunterSpec is the one hunter spec that fights in melee in Forever.
const meleeHunterSpec = "hunter-survival"

func attackPowerFor(sheet *proto.SampleSheet, spec string) float64 {
	if class, _ := splitSpecSlug(spec); class == hunterClass && spec != meleeHunterSpec {
		return sheet.GetRangedAttackPower()
	}
	return sheet.GetAttackPower()
}
