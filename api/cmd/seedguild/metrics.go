// api/cmd/seedguild/metrics.go
//
// Per-fight, per-raider fight_metrics rows, in the exact shape
// api/internal/rankings.Store.WriteFight writes after a real ingest (same columns, same
// state/phase conventions) — see that function's own insert for the shape this mirrors.
// Values are synthesized with a fixed random seed, not hand-picked, so every character
// gets Classic-plausible but not identical numbers across fights.
package main

import "math/rand"

// fightMetricRow is one fight_metrics row this tool writes.
type fightMetricRow struct {
	PlayerKey   string
	PlayerName  string
	Class       string
	Spec        string
	Role        string
	Ilvl        int
	MetricDPS   float64
	MetricHPS   *float64
	DamageTaken int64
	ActiveMS    int
	Deaths      int
}

// seededRand is a deterministic generator keyed on the report tag and fight index, so a
// rerun of this tool (or a dry-run followed by an apply) produces byte-identical numbers
// for the same fight — nothing here depends on wall-clock entropy.
func seededRand(reportTag string, fightIndex int) *rand.Rand {
	var seed int64
	for _, r := range reportTag {
		seed = seed*131 + int64(r)
	}
	seed = seed*131 + int64(fightIndex)
	return rand.New(rand.NewSource(seed))
}

// deathPronePoolWeight is how many extra entries a DeathProne raider (roster.go) gets in
// rollDeaths' own victim pool, concentrating most deaths on a handful of raiders rather
// than spreading them evenly - the way most real guilds have a few reliable fire-standers.
const deathPronePoolWeight = 5

// deathVictimPool is the roster's indices, repeated deathPronePoolWeight times for a
// DeathProne character and once for everyone else, for rollDeaths to sample from.
func deathVictimPool(roster []mockCharacter) []int {
	pool := make([]int, 0, len(roster)*2)
	for i, c := range roster {
		weight := 1
		if c.DeathProne {
			weight = deathPronePoolWeight
		}
		for j := 0; j < weight; j++ {
			pool = append(pool, i)
		}
	}
	return pool
}

// rollDeaths picks which roster indices die on a fight and how many times: zero on
// trash, zero to three on a kill (weighted toward zero), four to eight on a wipe -
// concentrated on this roster's DeathProne raiders via deathVictimPool, the same way a
// handful of raiders account for most of a real guild's deaths.
func rollDeaths(rng *rand.Rand, roster []mockCharacter, kill, isTrash bool) map[int]int {
	out := map[int]int{}
	if isTrash || len(roster) == 0 {
		return out
	}
	var count int
	switch {
	case !kill:
		count = 4 + rng.Intn(5) // 4-8 on a wipe
	default:
		switch r := rng.Float64(); {
		case r < 0.45:
			count = 0
		case r < 0.75:
			count = 1
		case r < 0.92:
			count = 2
		default:
			count = 3
		}
	}
	if count > len(roster) {
		count = len(roster)
	}
	pool := deathVictimPool(roster)
	for attempts := 0; len(out) < count && attempts < count*30; attempts++ {
		idx := pool[rng.Intn(len(pool))]
		if _, already := out[idx]; already {
			continue
		}
		hits := 1
		if !kill && rng.Intn(3) == 0 {
			hits = 2 // a wipe's unlucky raider who ate a second hit before release
		}
		out[idx] = hits
	}
	return out
}

// buildFightMetrics generates one row per roster character for a single fight.
func buildFightMetrics(roster []mockCharacter, keys []string, plan fightPlan, rng *rand.Rand) []fightMetricRow {
	deaths := rollDeaths(rng, roster, plan.Kill, plan.Trash)
	rows := make([]fightMetricRow, len(roster))
	for i, c := range roster {
		activeFrac := 0.85 + rng.Float64()*0.13
		row := fightMetricRow{
			PlayerKey: keys[i], PlayerName: c.Name, Class: c.Class, Spec: c.Spec, Role: c.Role,
			Ilvl: c.ItemLevel, Deaths: deaths[i],
			ActiveMS: int(float64(plan.DurationMS) * activeFrac),
		}
		switch c.Role {
		case roleTank:
			row.MetricDPS = 110 + rng.Float64()*90
			row.DamageTaken = int64(45_000+rng.Intn(45_000)) * int64(plan.DurationMS) / 240_000
		case roleHealer:
			hps := 260 + rng.Float64()*280
			row.MetricHPS = &hps
			row.MetricDPS = 15 + rng.Float64()*45
			row.DamageTaken = int64(8_000+rng.Intn(12_000)) * int64(plan.DurationMS) / 240_000
		default:
			row.MetricDPS = 170 + rng.Float64()*260
			row.DamageTaken = int64(10_000+rng.Intn(15_000)) * int64(plan.DurationMS) / 240_000
		}
		rows[i] = row
	}
	return rows
}

// talentSplitFor is the single-string talent summary fight_metrics.talent_split carries
// (rankings.Store.WriteFight's own "split" value) — this tool's own flavour string, not a
// read of a real talent tree, matching talentString's own leniency.
func talentSplitFor(role string) string {
	switch role {
	case roleTank:
		return "5/46/0"
	case roleHealer:
		return "0/46/5"
	default:
		return "0/5/46"
	}
}
