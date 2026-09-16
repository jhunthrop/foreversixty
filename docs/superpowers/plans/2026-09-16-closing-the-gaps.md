# Closing the Gaps Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Close the four places the report page sits behind Warcraft Logs — threat, Compare, resources at cap, phases — plus the two smaller asks every persona round repeated (pet casts, aura refreshes) and per-ability graphs, all from data the engine already parses.

**Architecture:** Six features, six task groups, each merged to `main` on its own. Four of them add fields to the engine's summary (`ThreatPair.series`, `ResourceTrack.max`/`at_max_ms`/`wasted`, `CastRow.owner_guid`, `Summary.phases`), all under one version bump to `0.4.0`; every new field is optional on the web so a report parsed before 0.4.0 renders as it does today with the new panels absent. Two features (Compare at ability level, per-ability graphs) are web-only: the ability diff is arithmetic over two summaries, and the ability line is one more measure on the shared DuckDB query layer. Phase data is curated into the existing per-encounter mechanics tables (`logs/engine/mechanics/tables/<id>.json`), which already ship inside the engine via `go:embed`.

**Tech Stack:** Go 1.22+ (`logs/engine`, `go:embed`), Astro 7 + Svelte 5 runes + Tailwind v4 (`web/`), vitest, Playwright, DuckDB-wasm.

**Spec:** `docs/superpowers/specs/2026-09-16-closing-the-gaps-design.md`

## Global Constraints

Copied verbatim from the spec's "Global constraints":

- The engine version goes to 0.4.0 with the first summary-shape change and stays there for
  the rest; the fixture is regenerated and prettier-formatted on every engine change, and
  the fixture test reads the version from the engine source.
- Every summary field added is optional on the web (`?:` in `types.ts`): a report parsed
  before 0.4.0 renders as it does today, with the new panels absent rather than broken.
- Every figure that is prorated rather than measured carries the tilde and the existing
  approximate title; a figure that is measured never does.
- The night (`?fight=all`) folds what folds by name and hides what cannot fold (per-second
  series across pulls); nothing on the night prints a figure the fold cannot stand behind.
- Phone first: every new panel is checked at 390px with no horizontal scroll, a 44px tap
  target, and its words beside its figures where the desktop has column headers.
- No third-party requests; the DuckDB measures use the shared query layer and the
  versioned fight-file urls.
- Each feature is a task group in the plan, reviewed and merged to main on its own, so a
  later feature never holds an earlier one back.

## Working rules

- Branch per group: `threat-over-time`, `compare-abilities`, `resources-at-cap`,
  `pet-casts-and-refreshes`, `ability-graphs`, `phases`. Never commit to `main` directly;
  merge each group when its own tests pass.
- Engine tasks bump nothing after the first: `logs/engine/session/session.go`'s
  `const Version` goes `"0.3.5"` → `"0.4.0"` in Task 1 and stays there. Tasks 8, 11 and 18
  change the summary shape without touching `Version`.
- Every engine task ends with `FOREVER_UPDATE_GOLDEN=1 go test ./...` then `go test ./...`
  in `logs/`, `go test ./...` in `api/`, and `npm run make:report-fixture` +
  `npx prettier --write src/fixtures/report` in `web/`. `web/src/fixtures/report/fixture.test.ts`
  reads the version out of `logs/engine/session/session.go`, so an engine bump with no
  fixture regeneration leaves vitest red.
- Node must be on the path as `~/.nvm/versions/node/v22.12.0/bin`; prefix web commands with
  `export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH"` if `node -v` is not v22.12.0.
- Run only the tests for the files you touch; CI runs the rest.
- Type check with `NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"` and read the
  count — colour codes defeat a grep for the word "error". It must report `- 0 errors`.
- Lint the files you touched: `npx eslint <files>`; format with `npx prettier --write <files>`.
- Before any e2e run: `FOREVER_DATA=fixture npm run build && ls dist/report-island.js`. A
  `{@const}` that is not an immediate child of `{#if}`/`{#each}` compiles fine in dev and
  fails only at island build time, silently, unless that file is checked for.
- Phone checks run as `npx playwright test tests/e2e/report-phone.spec.ts --project=mobile`.
  That suite measures at 360px, not the spec's 390px: 360 is the narrowest width
  `design/DESIGN-SYSTEM.md` covers and the width every other phone spec in the repo uses,
  so it is the stricter of the two and passing it passes 390.

---

## Group 1 — Threat over time, and threat measured inside a window

Tasks 1 to 4. Review gate at the end of Task 4: the whole group's e2e green, the fixture
regenerated, `astro check` at 0 errors, and the phone sweep passing.

---

### Task 1: The engine gives every threat pair a per-second series

**Files:**
- Modify: `logs/engine/summary/threat.go` (`ThreatPair`, `creditThreat`, `spreadThreat`, `threatPairs`)
- Modify: `logs/engine/summary/summary.go` (`Accumulator.threatBy` value type, `New`)
- Modify: `logs/engine/summary/damage.go` (the `creditThreat` call in the Damage case)
- Modify: `logs/engine/session/session.go` (`const Version`)
- Test: `logs/engine/summary/threat_test.go`
- Regenerate: `logs/engine/summary/testdata/v16.summary.json.golden`, `web/src/fixtures/report/`

**Interfaces:**
- Consumes: `Accumulator.bucket(time.Time) int`, `Accumulator.ms`, `Options.Threat` (`ThreatModel`), `units.Enemy`.
- Produces, on `summary.ThreatPair`:

```go
// Series is the threat this player built on this enemy in each whole second of
// the fight, on the same one-second buckets the damage series use, rounded to
// an integer. Its sum is Threat to the unit.
Series []int64 `json:"series"`
```

- Produces, unexported in `summary.go`:

```go
// threatPair is one player's running threat on one enemy: the total, and the
// per-second buckets it was built in.
type threatPair struct {
	total  float64
	series []float64
}
```

- Produces: `const Version = "0.4.0"` in `logs/engine/session/session.go`.

**Decision taken here:** the base threat model itself does not change. `BaseThreat.Version`
stays `"base-1"` and `Complete` stays false until the per-class modifiers land, as the spec
says; this task changes only where the threat that model produces is booked. The series is
accumulated as float64 per bucket and rounded once, at Snapshot, so rounding does not
compound over a fight's worth of spread healing threat.

- [ ] **Step 1: Write the failing test**

Append to `logs/engine/summary/threat_test.go` (it already imports `event`, `fight` and
`units`; add `"math"` and `"slices"` to the import block):

```go
// The chart on the Threat tab is drawn from the pair's own series, and a window's
// standing and built figures are summed from it, so a series that did not reconcile
// with the pair's total would be a second answer to the same question.
func TestThreatPairsCarryAPerSecondSeriesThatSumsToTheTotal(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		dmg(0.5, tank, boss, 1464, "Slam", 300, -1),
		dmg(2.5, tank, boss, 1464, "Slam", 700, -1),
		heal(2.9, healer, tank, 2050, "Holy Light", 400, 0),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(5),
		Players: []string{tank, healer}}, "test")

	find := func(guid, target string) ThreatPair {
		t.Helper()
		for _, p := range s.ThreatByTarget {
			if p.GUID == guid && p.TargetGUID == target {
				return p
			}
		}
		t.Fatalf("no pair for %s on %s in %+v", guid, target, s.ThreatByTarget)
		return ThreatPair{}
	}

	// 300 in the first second, nothing in the second, 700 in the third.
	tankPair := find(tank, boss)
	if want := []int64{300, 0, 700}; !slices.Equal(tankPair.Series, want) {
		t.Errorf("tank's series on the boss = %v, want %v", tankPair.Series, want)
	}
	var sum int64
	for _, v := range tankPair.Series {
		sum += v
	}
	if sum != int64(math.Round(tankPair.Threat)) {
		t.Errorf("series sums to %d, pair total is %v: the two must agree to the unit", sum, tankPair.Threat)
	}

	// 400 effective healing at 0.5 is 200 threat, spread over the one enemy engaged
	// within the window, and it lands in the second the heal landed in.
	healerPair := find(healer, boss)
	if want := []int64{0, 0, 200}; !slices.Equal(healerPair.Series, want) {
		t.Errorf("healer's series on the boss = %v, want %v", healerPair.Series, want)
	}
}

// A pair with threat only in the first second still carries a one-element series, and
// never a nil one: the web reads `series` as "this report has the per-second split",
// and a null there is not the same sentence as an empty array.
func TestThreatPairSeriesIsNeverNull(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	e := dmg(0.2, mage, boss, 116, "Frostbolt", 50, -1)
	reg.Observe(e)
	a.Add(e)
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(1),
		Players: []string{mage}}, "test")
	if len(s.ThreatByTarget) != 1 {
		t.Fatalf("pairs = %+v", s.ThreatByTarget)
	}
	if s.ThreatByTarget[0].Series == nil {
		t.Fatal("Series must be an empty or filled slice, never nil")
	}
	if want := []int64{50}; !slices.Equal(s.ThreatByTarget[0].Series, want) {
		t.Errorf("series = %v, want %v", s.ThreatByTarget[0].Series, want)
	}
}
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/logs && go test ./engine/summary/ -run 'TestThreatPair'`
Expected: FAIL to compile — `p.Series undefined (type ThreatPair has no field or method Series)`.

- [ ] **Step 3: Implement**

In `logs/engine/summary/threat.go`, add `"math"` to the imports, give `ThreatPair` the
series field, and replace `creditThreat`, `spreadThreat` and `threatPairs`:

```go
// ThreatPair is one player's threat on one enemy for the fight.
type ThreatPair struct {
	GUID       string  `json:"guid"`
	Name       string  `json:"name"`
	TargetGUID string  `json:"target_guid"`
	TargetName string  `json:"target_name"`
	Threat     float64 `json:"threat"`
	// Series is the threat this player built on this enemy in each whole
	// second of the fight, on the same one-second buckets the damage series
	// use, rounded to an integer. Its sum is Threat to the unit, so the chart
	// drawn from it and the table drawn from the total cannot disagree.
	Series []int64 `json:"series"`
}

// creditThreat books threat from one player against one enemy, in the second it
// was built. The caller is responsible for knowing enemy is on the other side;
// creditThreat itself trusts it, the same way the damage and heal cases already
// trust the event's own flags rather than a registry lookup (see engage below).
// The damage case credits the pair wherever it credits the total, so the two can
// never disagree about what an enemy is.
func (a *Accumulator) creditThreat(player, enemy string, threat float64, bucket int) {
	if threat == 0 || enemy == "" {
		return
	}
	by := a.threatBy[player]
	if by == nil {
		by = map[string]*threatPair{}
		a.threatBy[player] = by
	}
	pair := by[enemy]
	if pair == nil {
		pair = &threatPair{}
		by[enemy] = pair
	}
	pair.total += threat
	for len(pair.series) <= bucket {
		pair.series = append(pair.series, 0)
	}
	pair.series[bucket] += threat
}

// spreadThreat books healing threat over every enemy engaged within the window,
// in the second the heal landed.
func (a *Accumulator) spreadThreat(player string, threat float64, at time.Time) {
	if threat == 0 {
		return
	}
	var live []string
	for guid, last := range a.engaged {
		if at.Sub(last) <= a.opt.EngagedWindow {
			live = append(live, guid)
		}
	}
	if len(live) == 0 {
		return
	}
	each := threat / float64(len(live))
	bucket := a.bucket(at)
	for _, guid := range live {
		a.creditThreat(player, guid, each, bucket)
	}
}

// threatPairs renders the per-target table, largest first. The series is rounded
// per bucket rather than scaled to the total: a bucket is a second of the fight
// and reads as one, and the rounding error over a fight is under half a point a
// second against totals in the millions.
func (a *Accumulator) threatPairs() []ThreatPair {
	out := []ThreatPair{}
	for player, by := range a.threatBy {
		for enemy, pair := range by {
			series := make([]int64, len(pair.series))
			for i, v := range pair.series {
				series[i] = int64(math.Round(v))
			}
			out = append(out, ThreatPair{
				GUID: player, Name: a.name(player),
				TargetGUID: enemy, TargetName: a.name(enemy),
				Threat: pair.total, Series: series,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Threat != out[j].Threat {
			return out[i].Threat > out[j].Threat
		}
		return out[i].GUID+out[i].TargetGUID < out[j].GUID+out[j].TargetGUID
	})
	return out
}
```

In `logs/engine/summary/summary.go`, add the `threatPair` type next to the `Accumulator`
declaration and change the field's value type:

```go
// threatPair is one player's running threat on one enemy: the total, and the
// per-second buckets it was built in. Kept as float64 so the rounding happens
// once, at Snapshot, rather than compounding over a fight's worth of heals.
type threatPair struct {
	total  float64
	series []float64
}
```

```go
	// threatBy is per-target threat: player guid -> enemy guid -> the pair.
	threatBy map[string]map[string]*threatPair
```

and in `New`:

```go
		threatBy:        map[string]map[string]*threatPair{},
```

In `logs/engine/summary/damage.go`, the Damage case's credit call becomes:

```go
			if units.Enemy(e.Dest.Flags) {
				a.creditThreat(src, e.Dest.GUID, th, a.bucket(e.Time))
			}
```

In `logs/engine/session/session.go`:

```go
const Version = "0.4.0"
```

- [ ] **Step 4: Run the tests**

Run: `cd /Users/jh/code/forever/logs && go test ./engine/summary/ -run 'TestThreat'`
Expected: PASS, all threat tests including the two existing ones.

- [ ] **Step 5: Regenerate the goldens and the web fixture**

Run:
```bash
cd /Users/jh/code/forever/logs && FOREVER_UPDATE_GOLDEN=1 go test ./... && go test ./... && (cd ../api && go test ./...)
export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH"
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && npx vitest run src/fixtures/report/fixture.test.ts
```
Expected: the golden gains a `series` array inside every `threat_by_target` entry; the
fixture's `report.json` reads `"engine_version": "0.4.0"`; the fixture test passes because
it reads the version out of `session.go`.

- [ ] **Step 6: Commit**

```bash
git add logs/engine/summary logs/engine/session web/src/fixtures/report
git commit -m "feat(logs): every threat pair carries its per-second series"
```

---

### Task 2: The web reads standing and built threat out of the series

**Files:**
- Modify: `web/src/lib/report/types.ts` (`ThreatPair` gains `series`, `standing`, `built`, `measured`)
- Modify: `web/src/lib/report/window.ts` (`scopeThreatPairs` replaces `scaleThreatPairs` when the series is there)
- Modify: `web/src/lib/report/night.ts` (the fold drops the series and the window figures)
- Test: `web/src/lib/report/window.test.ts`, `web/src/lib/report/night.test.ts`

**Interfaces:**
- Consumes: `summary.ThreatPair.series` from Task 1; `TimeWindow`, `sliceSeries`, `sumSeries`, `isFullWindow` from `window.ts`.
- Produces, on `ThreatPair` in `types.ts`:

```ts
/** Threat built on this enemy in each whole second of the fight. Absent before engine 0.4.0. */
series?: number[];
/** Set by the window scope: the cumulative threat at the window's end — the number that decides aggro. */
standing?: number;
/** Set by the window scope: the threat built inside the window, the series summed over it. */
built?: number;
/** Set by the window scope: true when standing and built were measured from the series, not scaled. */
measured?: boolean;
```

- Produces, exported from `window.ts`:

```ts
export function scopeThreatPairs(pairs: ThreatPair[], window: TimeWindow): ThreatPair[]
```

- [ ] **Step 1: Write the failing tests**

In `web/src/lib/report/window.test.ts`, add `scopeThreatPairs` to the import list from
`./window`, `type ThreatPair` to the import from `./types`, and append:

```ts
describe('threat inside a window', () => {
  const pair = (guid: string, threat: number, series: number[]): ThreatPair => ({
    guid,
    name: guid,
    target_guid: 'Creature-1',
    target_name: 'Boss',
    threat,
    series,
  });

  it('measures standing at the window’s end and built inside it, from the pair’s own series', () => {
    const [scoped] = scopeThreatPairs([pair('P1', 100, [10, 20, 30, 40])], {
      startMs: 1000,
      endMs: 3000,
    });
    // Standing is everything up to the window's end: 10 + 20 + 30.
    expect(scoped.standing).toBe(60);
    // Built is the window's own buckets: 20 + 30.
    expect(scoped.built).toBe(50);
    expect(scoped.measured).toBe(true);
    // The table sorts and draws on standing, so `threat` is standing.
    expect(scoped.threat).toBe(60);
    expect(scoped.series).toEqual([20, 30]);
  });

  it('leaves a whole-fight window alone: standing at the end is the total', () => {
    const [scoped] = scopeThreatPairs([pair('P1', 100, [10, 20, 30, 40])], {
      startMs: 0,
      endMs: 4000,
    });
    expect(scoped.standing).toBe(100);
    expect(scoped.built).toBe(100);
    expect(scoped.threat).toBe(100);
  });

  it('says nothing was measured when the pair was written before the engine kept a series', () => {
    const old: ThreatPair = {
      guid: 'P1',
      name: 'P1',
      target_guid: 'Creature-1',
      target_name: 'Boss',
      threat: 100,
    };
    const [scoped] = scopeThreatPairs([old], { startMs: 1000, endMs: 3000 });
    expect(scoped.measured).toBeUndefined();
    expect(scoped.standing).toBeUndefined();
    expect(scoped.threat).toBe(100);
  });

  it('scopeSummary measures the pairs when they carry a series and scales them when they do not', () => {
    const base = summary;
    const withSeries: Summary = {
      ...base,
      threat: [{ guid: 'P1', name: 'P1', threat: 100, model_version: 'base-1', complete: false }],
      threat_by_target: [pair('P1', 100, [10, 20, 30, 40])],
      duration_ms: 4000,
      damage_done: [],
      healing: [],
    };
    const scoped = scopeSummary(withSeries, { startMs: 1000, endMs: 3000 });
    expect(scoped.threat_by_target?.[0].measured).toBe(true);
    expect(scoped.threat_by_target?.[0].standing).toBe(60);

    const withoutSeries: Summary = {
      ...withSeries,
      threat_by_target: [{ ...pair('P1', 100, []), series: undefined }],
    };
    const legacy = scopeSummary(withoutSeries, { startMs: 1000, endMs: 3000 });
    // No series, no damage and no healing in the window: the old ratio scales it to zero.
    expect(legacy.threat_by_target?.[0].measured).toBeUndefined();
    expect(legacy.threat_by_target?.[0].threat).toBe(0);
  });
});
```

In `web/src/lib/report/night.test.ts`, inside the existing describe that holds
`'folds per-target threat by enemy name and shifts taunts onto the night’s clock'`, add:

```ts
  it('drops the per-second series from the night’s threat pairs: a night has no one clock', () => {
    const fights = [fight(3, 'Kaal', false, 10_000), fight(4, 'Kaal', true, 10_000)];
    const summaries = new Map([
      [3, summary(3, 10_000, [roster('A', 10)])],
      [4, summary(4, 10_000, [roster('A', 10)])],
    ]);
    const pair = (threat: number, series: number[]) => [
      {
        guid: 'Player-1',
        name: 'Tank',
        target_guid: 'Creature-a',
        target_name: 'Kaal',
        threat,
        series,
      },
    ];
    summaries.set(3, { ...summaries.get(3)!, threat_by_target: pair(100, [60, 40]) });
    summaries.set(4, { ...summaries.get(4)!, threat_by_target: pair(50, [50]) });
    const night = nightSummary(fights, summaries);
    expect(night.threat_by_target?.[0].threat).toBe(150);
    expect(night.threat_by_target?.[0].series).toBeUndefined();
  });
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/window.test.ts src/lib/report/night.test.ts`
Expected: FAIL — `scopeThreatPairs is not a function` in window.test.ts, and the night case
fails on `series` still being `[60, 40]`.

- [ ] **Step 3: Implement**

In `web/src/lib/report/types.ts`, extend `ThreatPair`:

```ts
/** summary.ThreatPair — one player's threat on one enemy, for the threat-by-target view. */
export interface ThreatPair {
  guid: string;
  name: string;
  target_guid: string;
  target_name: string;
  threat: number;
  /** Threat built on this enemy in each whole second of the fight. Absent before engine 0.4.0. */
  series?: number[];
  /** Set by the window scope: the cumulative threat at the window's end — the number that decides aggro. */
  standing?: number;
  /** Set by the window scope: the threat built inside the window, the series summed over it. */
  built?: number;
  /** Set by the window scope: true when standing and built were measured from the series, not scaled. */
  measured?: boolean;
}
```

In `web/src/lib/report/window.ts`, add the exported function beside `scaleThreatPairs`
(keep `scaleThreatPairs`: it is still what a pre-0.4.0 report gets) and change
`scopeSummary` to pick between them:

```ts
/**
 * A pair's two honest window figures, both measured from its own per-second series:
 * `standing`, the cumulative threat at the window's end, which is the number that decides
 * who the enemy is looking at, and `built`, the threat made inside the window. `threat`
 * becomes standing, because that is what the table sorts by, draws a bar for and takes a
 * share of. A pair written before engine 0.4.0 has no series and comes back untouched, for
 * `scaleThreatPairs` to prorate the old way.
 */
export function scopeThreatPairs(pairs: ThreatPair[], window: TimeWindow): ThreatPair[] {
  return pairs.map((pair) => {
    if (pair.series === undefined) return pair;
    const standing = sumSeries(pair.series, { startMs: 0, endMs: window.endMs });
    const built = sumSeries(pair.series, window);
    return {
      ...pair,
      threat: standing,
      standing,
      built,
      measured: true,
      series: sliceSeries(pair.series, window),
    };
  });
}
```

In `scopeSummary`, replace the `scoped.threat_by_target = …` assignment:

```ts
  // Undefined survives scoping, the same way `taunts` does below: a summary written before
  // the engine kept the per-target split has no key at all, and the table reads that
  // differently from a split that came out empty. A split that carries per-second series
  // (engine 0.4.0) is measured rather than scaled, so the window's figures are the
  // window's own events; one that does not is prorated as before, and marked.
  scoped.threat_by_target =
    summary.threat_by_target === undefined
      ? undefined
      : summary.threat_by_target.every((pair) => pair.series !== undefined)
        ? scopeThreatPairs(summary.threat_by_target, window)
        : scaleThreatPairs(summary.threat_by_target, summary.threat, scoped.threat);
```

In `web/src/lib/report/night.ts`, the pair fold drops the series and the window figures —
a night is several clocks end to end and a cumulative line across them would claim a
standing nobody ever held:

```ts
    // Keyed by the target's name, not its GUID, for the same reason as auras and
    // exchanges: an add is a new GUID every pull. The night's target_guid is the name,
    // since the real per-pull GUID is not a stable identity across the fold. The
    // per-second series is dropped: the night has no one clock to draw threat on, and a
    // cumulative line spliced across pulls would read as a standing nobody ever held.
    for (const pair of summary.threat_by_target ?? []) {
      const key = `${pair.guid}|${pair.target_name}`;
      const found = threatPairs.get(key);
      threatPairs.set(
        key,
        found === undefined
          ? {
              ...pair,
              target_guid: pair.target_name,
              series: undefined,
              standing: undefined,
              built: undefined,
              measured: undefined,
            }
          : { ...found, threat: found.threat + pair.threat },
      );
    }
```

- [ ] **Step 4: Run the tests to verify they pass**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/window.test.ts src/lib/report/night.test.ts && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"`
Expected: PASS, and `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/lib/report/types.ts src/lib/report/window.ts src/lib/report/night.ts src/lib/report/window.test.ts src/lib/report/night.test.ts && npx prettier --write src/lib/report/types.ts src/lib/report/window.ts src/lib/report/night.ts src/lib/report/window.test.ts src/lib/report/night.test.ts
cd /Users/jh/code/forever && git add web/src/lib/report
git commit -m "feat(web): threat standing and threat built, measured from the pair's series"
```

---

### Task 3: The chart draws cumulative threat per player, with taunt marks

**Files:**
- Modify: `web/src/components/report/TimeChart.svelte` (two new optional props: `marks`, `perSecond`)
- Modify: `web/src/components/report/ThreatTable.svelte` (mount a TimeChart above the table)
- Modify: `web/src/components/report/ReportView.svelte` (pass the window and the setter through)
- Test: `web/tests/e2e/report-threat.spec.ts`

**Interfaces:**
- Consumes: `ThreatPair.series` (Task 2), `Taunt[]`, `classColorVar`, `formatDuration`, `TimeWindow`, `tauntKey`.
- Produces, on `TimeChart`'s props:

```ts
/** Ticks on the time axis, named on hover: the taunts under a threat chart. */
marks?: { atMs: number; label: string }[];
/** False when the lines are a running total rather than a rate: the caption and readout drop "per second". */
perSecond?: boolean;
```

- Produces, on `ThreatTable`'s props:

```ts
/** The page's window, so the chart's brush is the report's window. */
window?: TimeWindow;
/** True on the night, where there is no one clock and the chart is not drawn. */
nightMode?: boolean;
```

**Decisions taken here** (the spec leaves them open):
- The chart is `TimeChart` with an empty main `series` and one `extra` line per player, so
  the brush, the sliders, the hover readout and the reset button are the page's one
  implementation rather than a second canvas. `marks` and `perSecond` are the only new props.
- With "Every enemy" picked the chart shows the enemy carrying the most total threat and a
  caption saying which, exactly as the spec asks; the table above stays the totals table.
- Lines are cumulative: each player's series is its own running sum, so the line's value at
  a second is that player's standing then.

- [ ] **Step 1: Write the failing e2e tests**

Append to `web/tests/e2e/report-threat.spec.ts`:

```ts
// The chart is the answer to "did I pull it off the tank at 3:54": one cumulative line
// per player on the picked enemy, the taunts marked on the axis, and the same brush the
// rest of the page uses.
test('the threat chart draws one line per player and marks the taunts', async ({ page }) => {
  await page.goto(FIGHT);
  const chart = page.getByTestId('threat-chart');
  await expect(chart).toBeVisible();
  // Four players built threat on Warden Kelthas, so four legend entries.
  await expect(chart.getByTestId('chart-line-label')).toHaveCount(4);
  await expect(chart).toContainText('Baelgrim');
  // With no enemy picked the chart says which one it is drawing.
  await expect(page.getByTestId('threat-chart-note')).toContainText('Warden Kelthas');
  await expect(chart.getByTestId('chart-mark')).toHaveCount(1);
  await expect(chart.getByTestId('chart-mark').first()).toHaveAttribute('title', /Taunt/);
});

test('the threat chart’s own brush sets the report’s window', async ({ page }) => {
  await page.goto(FIGHT);
  await page.getByTestId('threat-chart').getByTestId('window-start').fill('6000');
  await expect(page).toHaveURL(/start=6000/);
});

test('the night has no threat chart, and says why', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&tab=threat');
  await expect(page.getByTestId('threat-chart')).toHaveCount(0);
  await expect(page.getByTestId('threat-chart-note')).toContainText('no clock');
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run:
```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-threat.spec.ts --project=desktop
```
Expected: FAIL — `threat-chart` not found on all three new cases.

- [ ] **Step 3: Implement the chart props**

In `web/src/components/report/TimeChart.svelte`, extend the props block:

```svelte
  let {
    series,
    extra = [],
    marks = [],
    perSecond = true,
    durationMs,
    window: current,
    deaths,
    label,
    onWindow,
  }: {
    series: number[];
    /** More lines drawn behind the main one, each in its token's colour. */
    extra?: { label: string; series: number[]; token: string }[];
    /** Ticks on the time axis, named on hover: the taunts under a threat chart. */
    marks?: { atMs: number; label: string }[];
    /** False when the lines are a running total rather than a rate: the caption and readout drop "per second". */
    perSecond?: boolean;
    durationMs: number;
    window: TimeWindow;
    deaths: { at_ms: number; name: string }[];
    label: string;
    onWindow: (window: TimeWindow | null) => void;
  } = $props();
```

Draw the marks in `draw()`, after the deaths block and before the hover line (gold ticks
rising from the axis, so they read as moments rather than as another series):

```js
    if (marks.length > 0) {
      context.strokeStyle = gold;
      context.fillStyle = gold;
      context.lineWidth = 2;
      for (const mark of marks) {
        const x = Math.round(xOf(mark.atMs)) + 0.5;
        context.beginPath();
        context.moveTo(x, HEIGHT);
        context.lineTo(x, HEIGHT - 14);
        context.stroke();
        context.beginPath();
        context.arc(x, HEIGHT - 16, 3, 0, Math.PI * 2);
        context.fill();
      }
    }
```

and add `marks` to the redraw effect's dependency list:

```js
    void [series, extra, marks, current, deaths, width, peak, hoverMs];
```

The marks also need a name on hover for a pointer and for a test, which a canvas cannot
give: put them over the canvas as absolutely positioned buttons, inside the same
`relative pl-12` wrapper that holds `chart-scale`, after the `{#if peak > 0}` block:

```svelte
    {#if marks.length > 0}
      <div class="pointer-events-none absolute inset-y-0 right-0 left-12" aria-hidden="true">
        {#each marks as mark, position (`${mark.atMs}-${position}`)}
          <span
            class="bg-gold pointer-events-auto absolute bottom-0 block h-[18px] w-[3px]"
            style={`left: ${durationMs === 0 ? 0 : (mark.atMs / durationMs) * 100}%`}
            title={`${mark.label} · ${formatDuration(mark.atMs)}`}
            data-testid="chart-mark"
          ></span>
        {/each}
      </div>
    {/if}
```

The caption's main swatch and the readout's main figure are the fight's damage line; with
an empty `series` there is none to name. Replace the `figcaption`'s first span and the
readout so both are conditional, and so the unit follows `perSecond`:

```svelte
    <span class="label text-muted"
      >{#if series.length > 0}<span class="whitespace-nowrap"
          ><span class="bg-gold mr-1 inline-block h-[2px] w-[14px] align-middle" aria-hidden="true"
          ></span>{label}{perSecond ? ' per second' : ''}</span
        >{:else}<span class="whitespace-nowrap">{label}</span>{/if}{#each extra as line (line.label)}
        <span class="ml-3 tracking-normal whitespace-nowrap normal-case" data-testid="chart-line-label"
          ><span
            class="mr-1 inline-block h-[2px] w-[14px] align-middle"
            style={`background: ${line.token}`}
            aria-hidden="true"
          ></span>{line.label}</span
        >{/each}{#if deaths.length > 0}
        <span class="ml-3 tracking-normal whitespace-nowrap normal-case"
          ><span class="bg-death mr-1 inline-block h-[10px] w-[2px] align-middle" aria-hidden="true"
          ></span>{deaths.length === 1 ? 'a death' : `${deaths.length} deaths`}</span
        >{/if}</span
    >
```

and in the readout span:

```svelte
      {#if hoverMs !== null && hoverValue !== null}
        <span class="text-text mr-3" data-testid="chart-readout"
          >{formatDuration(hoverMs)}{#if series.length > 0} · {label.toLowerCase()}
            {formatAmount(hoverValue)}{perSecond ? '/s' : ''}{/if}{#each hoverExtra as line (line.label)}
            · {line.label.toLowerCase()} {formatAmount(line.value)}{perSecond ? '/s' : ''}{/each}</span
        >
      {/if}
```

`hoverIndex` must not go negative on an empty main series, so the longest line decides it:

```js
  const seriesLength = $derived(
    extra.reduce((longest, line) => Math.max(longest, line.series.length), series.length),
  );
  const hoverIndex = $derived(
    hoverMs === null ? null : Math.min(seriesLength - 1, Math.floor(hoverMs / BUCKET_MS)),
  );
  const hoverValue = $derived(hoverIndex === null ? null : (series[hoverIndex] ?? 0));
```

- [ ] **Step 4: Implement the threat chart**

In `web/src/components/report/ThreatTable.svelte`, add the imports and props:

```ts
  import TimeChart from './TimeChart.svelte';
```

The file already imports `aroundWindow` and `type TimeWindow` from `../../lib/report/window`,
so nothing else is needed there. Add to the props block, after `durationMs` so the default
can read it:

```ts
    window = { startMs: 0, endMs: durationMs },
    nightMode = false,
```

```ts
    /** The page's window, so the chart's brush is the report's window. */
    window?: TimeWindow;
    /** True on the night, where there is no one clock and the chart is not drawn. */
    nightMode?: boolean;
```

Note the prop is declared after `durationMs` in the destructuring so the default can read
it. Then derive the chart's lines:

```ts
  /**
   * The enemy the chart draws: the picked one, or — with "Every enemy" picked — the one
   * carrying the most threat, named under the chart so nobody reads it as the raid's total.
   */
  const charted = $derived(picked ?? everyoneGroups[0] ?? groups[0]);
  /** A running total, so the line's value at a second is that player's standing then. */
  function cumulative(series: number[]): number[] {
    let running = 0;
    return series.map((value) => (running += value));
  }
  /**
   * One line per player on the charted enemy, each in their class colour, longest first so
   * the legend reads top-down like the table. Every pair of a player on this enemy is
   * summed: six adds of one name are one enemy, and a player's line on it is the sum.
   */
  const chartLines = $derived.by(() => {
    if (charted === undefined) return [];
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byPlayer = new Map<string, number[]>();
    for (const pair of pairs) {
      if (pair.series === undefined || !charted.guids.includes(pair.target_guid)) continue;
      const found = byPlayer.get(pair.guid) ?? [];
      const merged = [...found];
      pair.series.forEach((value, index) => {
        merged[index] = (merged[index] ?? 0) + value;
      });
      byPlayer.set(pair.guid, merged);
    }
    return [...byPlayer.entries()]
      .map(([guid, series]) => ({
        label: splitUnitName(names.get(guid) ?? guid).name,
        series: cumulative(series.map((value) => value ?? 0)),
        token: classColorVar(classOf.get(guid)),
      }))
      .sort((a, b) => (b.series[b.series.length - 1] ?? 0) - (a.series[a.series.length - 1] ?? 0));
  });
  /** The taunts on the charted enemy, as marks on the chart's time axis. */
  const chartMarks = $derived(
    orderedTaunts.map((taunt) => ({
      atMs: taunt.at_ms,
      label: `${nameOf(taunt.source_guid, taunt.source_name)} · ${taunt.spell_name} on ${nameOf(taunt.target_guid, taunt.target_name)}`,
    })),
  );
  const showChart = $derived(!nightMode && chartLines.length > 0);
```

`nameOf` and `splitUnitName` are already in the file; `charted`, `chartLines` and
`chartMarks` must be declared after `picked`, `everyoneGroups`, `groups` and
`orderedTaunts` so the `$derived` graph resolves. The `onWindow` prop's type widens to
`(window: TimeWindow | null) => void`, because the chart's own `null` -- a tap rather than
a drag -- clears the window; the taunt list's existing `onWindow(aroundWindow(...))` call
still passes a window and is untouched. Mount the chart directly above the enemy picker,
inside the root `div`, wrapped so the test id sits on one element rather than a sibling
that has to be kept in step with it:

```svelte
  {#if showChart}
    <div data-testid="threat-chart">
      <TimeChart
        series={[]}
        extra={chartLines}
        marks={chartMarks}
        perSecond={false}
        {durationMs}
        {window}
        deaths={[]}
        label={`Threat on ${charted?.name ?? ''}`}
        {onWindow}
      />
    </div>
  {/if}
  <p class="text-muted text-[12px]" data-testid="threat-chart-note">
    {#if nightMode}
      A night has no clock to draw threat on, so there is no chart here: the totals below are every
      pull's threat added up.
    {:else if showChart}
      One cumulative line per player, on {charted?.name}{picked === undefined
        ? ', the enemy carrying the most threat; pick another above'
        : ''}. A taunt in the game puts the taunter on top, and the base threat model does not: the
      lines are what damage and healing built, so a taunt mark is the moment the order stopped
      matching them.
    {:else}
      This report was parsed before threat was kept second by second, so there is no chart; the
      totals below are the whole fight's.
    {/if}
  </p>
```

In `web/src/components/report/ReportView.svelte`, pass the two new props to the existing
`ThreatTable` mount:

```svelte
            window={cutWindow}
            {nightMode}
```

`ReportView`'s `onWindow` is already `setWindow`, which takes `TimeWindow | null`, so the
widened prop type is exactly what it already passes:

```ts
    onWindow: (window: TimeWindow | null) => void;
```

- [ ] **Step 5: Run the tests to verify they pass**

Run:
```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-threat.spec.ts --project=desktop && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: PASS on every case in the file (the pre-existing ones included), `- 0 errors`.

- [ ] **Step 6: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/components/report/TimeChart.svelte src/components/report/ThreatTable.svelte src/components/report/ReportView.svelte tests/e2e/report-threat.spec.ts && npx prettier --write src/components/report/TimeChart.svelte src/components/report/ThreatTable.svelte src/components/report/ReportView.svelte tests/e2e/report-threat.spec.ts
cd /Users/jh/code/forever && git add web/src/components/report web/tests/e2e/report-threat.spec.ts
git commit -m "feat(web): threat over time, one line per player, with the taunts marked"
```

---

### Task 4: The Threat table stops prorating under a window

**Files:**
- Modify: `web/src/components/report/ThreatTable.svelte` (standing and built columns, the note, the CSV)
- Modify: `web/src/components/report/Glossary.svelte` (two terms)
- Test: `web/tests/e2e/report-threat.spec.ts`, `web/tests/e2e/report-phone.spec.ts`

**Interfaces:**
- Consumes: `ThreatPair.standing`, `ThreatPair.built`, `ThreatPair.measured` (Task 2).
- Produces: no new exports. `ThreatLine` inside `ThreatTable.svelte` gains two fields:

```ts
  interface ThreatLine {
    guid: string;
    name: string;
    threat: number;
    /** The threat built inside the window; undefined when the whole fight is on screen. */
    built?: number;
    /** True when threat and built came from the pair's own series rather than a ratio. */
    measured?: boolean;
  }
```

**Decision taken here:** the measured path applies to both tables. On a picked enemy every
line is a pair and carries standing and built. On "Every enemy" a player's standing is the
sum of their standings across every enemy, and their built the sum of their built; the
enemies' own rows have no pair at all, keep the scaled total, and stay outside the share
exactly as they do today.

- [ ] **Step 1: Write the failing e2e tests**

In `web/tests/e2e/report-threat.spec.ts`, replace the existing case
`'a brushed window shows greyed totals with no ranking, and the taunt keeps its time'`
with the two cases below (the old behaviour is the thing this task removes), and append
the third:

```ts
// Inside a brush the table is measured from the pairs' own per-second series: standing at
// the window's end is what decides aggro, built inside it is what the window did. Both
// are measured, so there is no tilde, no "no order to read" note, and the ranking stands.
test('a brushed window reads standing and built from the series, measured', async ({ page }) => {
  await page.goto(`${FIGHT}&start=3000&end=14000&target=${KELTHAS}`);
  await expect(page.getByTestId('threat-approximate-note')).toHaveCount(0);
  await expect(page.getByTestId('threat-window-note')).toContainText('Standing');
  const rows = page.getByTestId('threat-on-target').locator('li');
  await expect(rows.first()).toContainText('Baelgrim');
  await expect(rows.first().getByTestId('threat-standing')).not.toContainText('~');
  await expect(rows.first().getByTestId('threat-built')).toBeVisible();
  await expect(page.getByTestId('threat-share')).not.toHaveCount(0);
  await expect(page.getByTestId('threat-table').locator('ul[data-ranked="true"]')).toHaveCount(1);
  // The taunt keeps the pull's clock, so the same taunt reads the same time from either side.
  await expect(page.getByTestId('threat-taunts')).toContainText('8.5s');
});

test('the whole fight is unchanged: standing at the end is the total', async ({ page }) => {
  await page.goto(`${FIGHT}&target=${KELTHAS}`);
  const rows = page.getByTestId('threat-on-target').locator('li');
  await expect(rows.first().getByTestId('threat-share')).toContainText('49.2%');
  // No second figure when there is no window: the whole fight has only one number.
  await expect(page.getByTestId('threat-built')).toHaveCount(0);
});

// "Around it" on a taunt sets the window; the table it lands on is measured, not prorated.
test('a taunt’s window lands on a measured table', async ({ page }) => {
  await page.goto(`${FIGHT}&target=${KELTHAS}`);
  await page.getByTestId('threat-taunts').getByTestId('taunt-window').first().click();
  await expect(page).toHaveURL(/start=3000&end=14000/);
  await expect(page.getByTestId('threat-window-note')).toContainText('Standing');
  await expect(page.getByTestId('threat-approximate-note')).toHaveCount(0);
});

test('the glossary explains standing and built', async ({ page }) => {
  await page.goto(FIGHT);
  await page.getByTestId('glossary').locator('summary').click();
  await expect(page.getByTestId('glossary')).toContainText('Standing threat');
  await expect(page.getByTestId('glossary')).toContainText('Threat built in a window');
});
```

In `web/tests/e2e/report-phone.spec.ts`, no new state is needed (`tab=threat` is already in
`STATES`); add one case at the end of the file so the two new figures say what they are
without a column header:

```ts
test('a windowed threat row says which figure is standing and which is built', async ({ page }) => {
  await page.goto(`${REPORT}&tab=threat&start=3000&end=14000&target=Creature-0-2085-2284-7855-169754-0000AA0002`);
  const row = page.getByTestId('threat-on-target').locator('li').first();
  await expect(row.getByTestId('threat-standing')).toContainText('standing');
  await expect(row.getByTestId('threat-built')).toContainText('built');
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run:
```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-threat.spec.ts --project=desktop
```
Expected: FAIL — `threat-window-note` not found, `threat-approximate-note` still present.

- [ ] **Step 3: Implement**

In `web/src/components/report/ThreatTable.svelte`:

Extend `ThreatLine` as in Interfaces above. In `groupPairs`, carry the two figures through:

```ts
      const have = found.players.get(pair.guid);
      found.players.set(pair.guid, {
        guid: pair.guid,
        name: splitUnitName(pair.name).name,
        threat: (have?.threat ?? 0) + pair.threat,
        built: pair.built === undefined ? have?.built : (have?.built ?? 0) + pair.built,
        measured: pair.measured === true || have?.measured === true,
      });
```

Add a fold of the same figures per player across every enemy, for the totals table, and
use it where a player has one:

```ts
  /**
   * Per player, standing and built across every enemy: the totals table's own measured
   * figures. An enemy's own threat row has no pair and keeps the summary's scaled total,
   * which is what it has always been and what the share already leaves out.
   */
  const measuredTotals = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const out = new Map<string, { threat: number; built: number }>();
    for (const pair of pairs) {
      if (pair.measured !== true) continue;
      const found = out.get(pair.guid) ?? { threat: 0, built: 0 };
      out.set(pair.guid, {
        threat: found.threat + (pair.standing ?? pair.threat),
        built: found.built + (pair.built ?? 0),
      });
    }
    return out;
  });
  /** The totals table's lines, with a player's measured standing standing in for the scaled total. */
  const totalLines = $derived<ThreatLine[]>(
    ordered.map((row) => {
      const found = measuredTotals.get(row.guid);
      return found === undefined
        ? { guid: row.guid, name: row.name, threat: row.threat }
        : { guid: row.guid, name: row.name, threat: found.threat, built: found.built, measured: true };
    }),
  );
```

`lines` now picks between the two and no longer reorders under a window, because a measured
standing is a ranking:

```ts
  /** Every line on screen is measured: the window's figures came from the series, not a ratio. */
  const allMeasured = $derived(lines.length > 0 && lines.every((line) => line.measured === true));
  /**
   * Under a brush with nothing measured the figures are scaled totals, not a standing, so
   * the rows keep the roster's order (by name) rather than an order that reads as a
   * ranking. Measured lines are a ranking and keep it.
   */
  const lines = $derived<ThreatLine[]>(
    approximate && !(picked?.lines ?? totalLines).every((line) => line.measured === true)
      ? [...(picked?.lines ?? totalLines)].sort((a, b) => a.name.localeCompare(b.name))
      : [...(picked?.lines ?? totalLines)].sort((a, b) => b.threat - a.threat),
  );
```

Svelte 5 resolves `$derived` lazily, so `allMeasured` may be declared before `lines`; keep
them adjacent for a reader. `ranked` becomes true whenever nothing is prorated:

```ts
  /**
   * Bars and shares draw a ranking. A measured window is a ranking — standing is exactly
   * the number that decides aggro — so they are drawn. A window with nothing measured
   * (a report parsed before engine 0.4.0) still holds them back.
   */
  const ranked = $derived(!approximate || allMeasured);
  /** The tilde is for a scaled figure only; a measured one never carries it. */
  const scaled = $derived(approximate && !allMeasured);
  const mark = $derived(approximateMark(scaled));
  const title = $derived(approximateTitle(scaled));
```

(the two existing `mark`/`title` declarations are replaced by these; every other use of
`approximate` in the markup for marking purposes becomes `scaled`.)

The grid gains a column when the built figure is on screen:

```ts
  /** The second figure exists only under a window: the whole fight has one number. */
  const showBuilt = $derived(allMeasured && approximate);
  const rowGrid = $derived(
    showBuilt
      ? 'md:grid-cols-[minmax(120px,1.2fr)_minmax(0,3fr)_96px_96px_64px]'
      : ranked
        ? 'md:grid-cols-[minmax(120px,1.2fr)_minmax(0,3fr)_96px_64px]'
        : 'md:grid-cols-[minmax(120px,1.2fr)_96px]',
  );
```

In the header row, after the existing Total heading, add:

```svelte
      {#if showBuilt}
        <span class="text-right" title="Threat this player built inside the window">Built</span>
      {/if}
```

and rename the Total heading's own title under a window:

```svelte
      <span
        class="text-right"
        title={showBuilt
          ? 'Cumulative threat at the window’s end: the number that decides who the enemy attacks'
          : picked === undefined
            ? 'Threat accumulated from damage and healing over this window'
            : `Threat this player built on ${picked.name} over this window`}
        >{showBuilt ? 'Standing' : 'Total'}</span
      >
```

In the row, give the amount cell its test id and phone word, and add the built cell:

```svelte
          <span
            class={`tabular text-right font-mono ${ranked ? '' : 'text-muted'}`}
            {title}
            data-testid="threat-standing"
            aria-label={approximateAriaLabel(scaled, formatAmount(Math.round(row.threat)))}
          >
            {mark}{formatAmount(Math.round(row.threat))}<span class="label font-body ml-1.5 md:hidden"
              >{showBuilt ? 'standing' : 'threat'}</span
            >
          </span>
          {#if showBuilt}
            <span
              class="text-muted tabular text-right font-mono"
              title="Threat this player built inside the window"
              data-testid="threat-built"
            >
              {formatAmount(Math.round(row.built ?? 0))}<span class="label font-body ml-1.5 md:hidden"
                >built</span
              >
            </span>
          {/if}
```

Replace the `{#if approximate}` note block with one that keeps the old words only when
nothing was measured, and says which figure is which when something was:

```svelte
    {#if showBuilt}
      <p class="text-muted text-[12px]" data-testid="threat-window-note">
        Standing is the threat this player had built by the window’s end, which is the number that
        decides who the enemy attacks; Built is what they made inside the window. Both are measured
        from the fight’s own seconds, so neither is marked. The table sorts by Standing.
      </p>
    {:else if scaled}
      <p class="text-muted text-[12px]" data-testid="threat-approximate-note">
        This report was parsed before threat was kept second by second, so threat inside a window is
        not measured. Each figure is marked {mark} because it is the whole fight’s threat scaled to
        this window’s share of that player’s total, not the window’s own events, so it is not a
        standing and is not drawn as one: no bars, no shares, no order to read. Parse the report
        again for the measured figures. The whole pull is exact.
      </p>
    {/if}
```

The CSV gains the column:

```ts
  function csvLines(): string[][] {
    const share = (line: ThreatLine): string[] =>
      ranked ? [(total === 0 ? 0 : (line.threat / total) * 100).toFixed(1)] : [];
    const amount = showBuilt ? 'Standing' : 'Threat';
    return [
      [
        ...(picked === undefined ? ['Unit', amount] : ['Player', `${amount} on ${picked.name}`]),
        ...(showBuilt ? ['Built in the window'] : []),
        ...(ranked ? ['Share %'] : []),
      ],
      ...lines.map((line) => [
        splitUnitName(line.name).name,
        String(Math.round(line.threat)),
        ...(showBuilt ? [String(Math.round(line.built ?? 0))] : []),
        ...(shares(line) ? share(line) : ranked ? [''] : []),
      ]),
    ];
  }
```

In `web/src/components/report/Glossary.svelte`, add two entries to `TERMS`, after the
`'Threat on a target'` entry:

```ts
    {
      term: 'Standing threat',
      meaning:
        'On the Threat tab under a window: the threat a player had built on that enemy by the window’s end, counted from the pull’s start. It is the number that decides who the enemy attacks, so the table sorts by it. Measured from the fight’s own seconds, never scaled.',
    },
    {
      term: 'Threat built in a window',
      meaning:
        'The threat a player made inside the window itself, the seconds of the window added up. A player can be top of the standing and bottom of the built: they came in with a lead.',
    },
```

- [ ] **Step 4: Run the tests to verify they pass**

Run:
```bash
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && npx vitest run src/lib/report src/fixtures/report/fixture.test.ts && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-threat.spec.ts --project=desktop && npx playwright test tests/e2e/report-phone.spec.ts --project=mobile && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: the fixture regeneration is a no-op (Task 1 already did it, and nothing has
changed the engine since — `git status` shows no change under `web/src/fixtures/report`);
every threat case passes; the phone sweep passes with no sideways scroll; `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/components/report/ThreatTable.svelte src/components/report/Glossary.svelte tests/e2e/report-threat.spec.ts tests/e2e/report-phone.spec.ts && npx prettier --write src/components/report/ThreatTable.svelte src/components/report/Glossary.svelte tests/e2e/report-threat.spec.ts tests/e2e/report-phone.spec.ts
cd /Users/jh/code/forever && git add web/src/components/report web/tests/e2e web/src/fixtures/report
git commit -m "feat(web): threat inside a window is standing and built, measured"
```

- [ ] **Step 6: Group 1 review gate**

Run the whole group's checks and merge:
```bash
cd /Users/jh/code/forever/logs && go test ./... && (cd ../api && go test ./...)
cd /Users/jh/code/forever/web && npx vitest run && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?" && npx playwright test tests/e2e/report-threat.spec.ts tests/e2e/report-tabs.spec.ts tests/e2e/report-brush.spec.ts --project=desktop
```
Then merge `threat-over-time` into `main`.

---

## Group 2 — Compare at ability level, player against player, with a CSV

Tasks 5 to 7. No engine change. Review gate at the end of Task 7.

---

### Task 5: The ability diff builders

**Files:**
- Create: `web/src/lib/report/compare.ts`
- Create: `web/src/lib/report/compare.test.ts`

**Interfaces:**
- Consumes: `Summary`, `Ability`, `abilityKey` from `types.ts`.
- Produces:

```ts
export type CompareMetric =
  | 'damage_done' | 'dps' | 'healing_done' | 'hps'
  | 'damage_taken' | 'dtps' | 'threat' | 'tps';
export const COMPARE_METRICS: readonly CompareMetric[];
export const METRIC_LABELS: Record<CompareMetric, string>;
export const PER_SECOND_METRICS: ReadonlySet<CompareMetric>;
/** Which actor table a metric's ability split comes from; null for threat, which has none. */
export function metricTable(metric: CompareMetric): 'damage_done' | 'healing' | 'damage_taken' | null;
export interface AbilityDiff {
  /** types.ts's abilityKey: the spell and the pet it came via. */
  key: string;
  name: string;
  via?: string;
  /** Null when that side never used the ability, which the table prints as a dash. */
  a: number | null;
  b: number | null;
  /** (a ?? 0) - (b ?? 0): positive means this side did more. */
  delta: number;
}
/** One player's abilities in two fights, largest difference first. */
export function abilityDiff(
  left: Summary | null, right: Summary | null, guid: string, metric: CompareMetric,
): AbilityDiff[];
/** Two players inside one fight, ability by ability, largest difference first. */
export function playerAbilityDiff(
  summary: Summary | null, leftGuid: string, rightGuid: string, metric: CompareMetric,
): AbilityDiff[];
```

`METRIC_LABELS` and `PER_SECOND_METRICS` move out of `CompareMode.svelte` into this module
so the component and the diff builders read one list rather than two that can drift.

- [ ] **Step 1: Write the failing tests**

```ts
// web/src/lib/report/compare.test.ts
import { describe, expect, it } from 'vitest';
import { abilityDiff, metricTable, playerAbilityDiff } from './compare';
import type { Ability, Actor, Summary } from './types';

function ability(spell_id: number, name: string, effective: number, via?: string): Ability {
  return { spell_id, name, via, total: effective, effective, hits: 1, crits: 0, ticks: 0, min: 0, max: 0 };
}

function actor(guid: string, abilities: Ability[]): Actor {
  return {
    guid,
    name: guid,
    total: abilities.reduce((sum, a) => sum + a.total, 0),
    effective: abilities.reduce((sum, a) => sum + a.effective, 0),
    active_ms: 1000,
    abilities,
    targets: [],
    series: [],
  };
}

function summary(damage: Actor[], healing: Actor[] = []): Summary {
  return {
    engine_version: 'test',
    fight_index: 1,
    duration_ms: 10_000,
    damage_done: damage,
    damage_taken: [],
    healing,
    healing_taken: [],
    deaths: [],
    auras: [],
    casts: [],
    interrupts: [],
    dispels: [],
    resources: [],
    threat: [],
    combatants: [],
    roster: [],
  };
}

describe('metricTable', () => {
  it('sends each metric to the table its ability split lives in, and threat to none', () => {
    expect(metricTable('damage_done')).toBe('damage_done');
    expect(metricTable('dps')).toBe('damage_done');
    expect(metricTable('healing_done')).toBe('healing');
    expect(metricTable('hps')).toBe('healing');
    expect(metricTable('damage_taken')).toBe('damage_taken');
    expect(metricTable('dtps')).toBe('damage_taken');
    expect(metricTable('threat')).toBeNull();
    expect(metricTable('tps')).toBeNull();
  });
});

describe('abilityDiff', () => {
  const left = summary([actor('P1', [ability(1, 'Slam', 400), ability(2, 'Cleave', 100)])]);
  const right = summary([actor('P1', [ability(1, 'Slam', 250), ability(3, 'Execute', 300)])]);

  it('lists both pulls’ abilities with a signed difference, biggest first', () => {
    const rows = abilityDiff(left, right, 'P1', 'dps');
    expect(rows.map((row) => [row.name, row.a, row.b, row.delta])).toEqual([
      ['Execute', null, 300, -300],
      ['Slam', 400, 250, 150],
      ['Cleave', 100, null, 100],
    ]);
  });

  it('keeps a pet’s ability apart from its owner’s own of the same spell', () => {
    const owner = summary([actor('P1', [ability(1, 'Melee', 100), ability(1, 'Melee', 60, 'Ashfang')])]);
    const rows = abilityDiff(owner, summary([actor('P1', [])]), 'P1', 'dps');
    expect(rows.map((row) => [row.name, row.via, row.a])).toEqual([
      ['Melee', undefined, 100],
      ['Melee', 'Ashfang', 60],
    ]);
  });

  it('has nothing to split for threat, which has no per-ability table', () => {
    expect(abilityDiff(left, right, 'P1', 'threat')).toEqual([]);
  });

  it('is empty rather than throwing when a side is missing or the player is not in it', () => {
    expect(abilityDiff(left, null, 'P1', 'dps')).toEqual([
      { key: '1|', name: 'Slam', via: undefined, a: 400, b: null, delta: 400 },
      { key: '2|', name: 'Cleave', via: undefined, a: 100, b: null, delta: 100 },
    ]);
    expect(abilityDiff(left, right, 'nobody', 'dps')).toEqual([]);
  });

  it('reads healing off the healing table, not the damage one', () => {
    const healers = summary([], [actor('P2', [ability(9, 'Heal', 900)])]);
    expect(abilityDiff(healers, summary([], []), 'P2', 'hps').map((row) => row.name)).toEqual(['Heal']);
    expect(abilityDiff(healers, summary([], []), 'P2', 'dps')).toEqual([]);
  });
});

describe('playerAbilityDiff', () => {
  it('puts two players of one fight side by side', () => {
    const one = summary([
      actor('P1', [ability(1, 'Slam', 400)]),
      actor('P2', [ability(5, 'Frostbolt', 310)]),
    ]);
    expect(playerAbilityDiff(one, 'P1', 'P2', 'dps').map((row) => [row.name, row.a, row.b, row.delta])).toEqual(
      [
        ['Slam', 400, null, 400],
        ['Frostbolt', null, 310, -310],
      ],
    );
  });
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/compare.test.ts`
Expected: FAIL — `Failed to resolve import "./compare"`.

- [ ] **Step 3: Implement**

```ts
// web/src/lib/report/compare.ts
// Compare mode's arithmetic, kept out of the component so it can be tested without a DOM.
//
// Two pulls, or two players inside one pull, ability by ability. There is no engine call
// and no query here: every figure is already in the summaries the page has loaded, and the
// window has already been applied to both sides by window.ts's scopeSummary before these
// are asked. The one thing the summary cannot split is threat, which the engine keeps per
// player and never per ability, so a threat comparison expands to nothing and says so.
import { abilityKey, type Ability, type Summary } from './types';

export type CompareMetric =
  | 'damage_done'
  | 'dps'
  | 'healing_done'
  | 'hps'
  | 'damage_taken'
  | 'dtps'
  | 'threat'
  | 'tps';

/** The metric picker's order, which is the order the options are listed in. */
export const COMPARE_METRICS: readonly CompareMetric[] = [
  'dps',
  'damage_done',
  'hps',
  'healing_done',
  'dtps',
  'threat',
  'tps',
  'damage_taken',
];

/** The caption's words for each metric: a key like dtps is not a sentence. */
export const METRIC_LABELS: Record<CompareMetric, string> = {
  damage_done: 'damage done',
  dps: 'DPS',
  healing_done: 'healing done',
  hps: 'HPS',
  damage_taken: 'damage taken',
  dtps: 'damage taken per second',
  threat: 'threat',
  tps: 'threat per second',
};

export const PER_SECOND_METRICS: ReadonlySet<CompareMetric> = new Set<CompareMetric>([
  'dps',
  'hps',
  'dtps',
  'tps',
]);

/** Which actor table a metric's ability split comes from; null for threat, which has none. */
export function metricTable(metric: CompareMetric): 'damage_done' | 'healing' | 'damage_taken' | null {
  switch (metric) {
    case 'damage_done':
    case 'dps':
      return 'damage_done';
    case 'healing_done':
    case 'hps':
      return 'healing';
    case 'damage_taken':
    case 'dtps':
      return 'damage_taken';
    default:
      return null;
  }
}

/** One ability on both sides of a comparison. */
export interface AbilityDiff {
  /** types.ts's abilityKey: the spell and the pet it came via. */
  key: string;
  name: string;
  via?: string;
  /** Null when that side never used the ability, which the table prints as a dash. */
  a: number | null;
  b: number | null;
  /** (a ?? 0) - (b ?? 0): positive means this side did more. */
  delta: number;
}

/** One actor's abilities out of the table a metric reads, keyed the way the tables key them. */
function abilityAmounts(
  summary: Summary | null,
  guid: string,
  metric: CompareMetric,
): Map<string, Ability> {
  // A plain Map: built once and returned to diffOf, never read reactively by key.
  // eslint-disable-next-line svelte/prefer-svelte-reactivity
  const out = new Map<string, Ability>();
  const table = metricTable(metric);
  if (summary === null || table === null) return out;
  const actor = summary[table].find((row) => row.guid === guid);
  for (const ability of actor?.abilities ?? []) out.set(abilityKey(ability), ability);
  return out;
}

/**
 * The union of two ability tables as one signed list, largest difference first, so the
 * thing that changed most is the first line a reader sees. Ties break on the name, so two
 * runs over the same pull produce the same order.
 */
function diffOf(a: Map<string, Ability>, b: Map<string, Ability>): AbilityDiff[] {
  const out: AbilityDiff[] = [];
  for (const key of new Set([...a.keys(), ...b.keys()])) {
    const left = a.get(key);
    const right = b.get(key);
    const named = left ?? right;
    if (named === undefined) continue;
    const x = left === undefined ? null : left.effective;
    const y = right === undefined ? null : right.effective;
    out.push({ key, name: named.name, via: named.via, a: x, b: y, delta: (x ?? 0) - (y ?? 0) });
  }
  return out.sort((p, q) => Math.abs(q.delta) - Math.abs(p.delta) || p.name.localeCompare(q.name));
}

/** One player's abilities in two fights, largest difference first. */
export function abilityDiff(
  left: Summary | null,
  right: Summary | null,
  guid: string,
  metric: CompareMetric,
): AbilityDiff[] {
  return diffOf(abilityAmounts(left, guid, metric), abilityAmounts(right, guid, metric));
}

/** Two players inside one fight, ability by ability, largest difference first. */
export function playerAbilityDiff(
  summary: Summary | null,
  leftGuid: string,
  rightGuid: string,
  metric: CompareMetric,
): AbilityDiff[] {
  return diffOf(abilityAmounts(summary, leftGuid, metric), abilityAmounts(summary, rightGuid, metric));
}
```

- [ ] **Step 4: Run the tests to verify they pass**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/compare.test.ts && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"`
Expected: PASS, `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/lib/report/compare.ts src/lib/report/compare.test.ts && npx prettier --write src/lib/report/compare.ts src/lib/report/compare.test.ts
cd /Users/jh/code/forever && git add web/src/lib/report/compare.ts web/src/lib/report/compare.test.ts
git commit -m "feat(web): the ability diff behind Compare, as its own module"
```

---

### Task 6: The url carries the compared player as `vs`

**Files:**
- Modify: `web/src/lib/report/url.ts` (`ReportState.compareVs`, parse, serialise)
- Test: `web/src/lib/report/url.test.ts`

**Interfaces:**
- Produces, on `ReportState`:

```ts
  /** Compare mode's second player, as a GUID; '' compares two fights instead. */
  compareVs: string;
```

serialised as `vs=<guid>`, parsed with the same pattern `source` uses (a GUID is the
engine's own format: a kind, then hyphen-separated numeric fields).

- [ ] **Step 1: Write the failing tests**

Append to `web/src/lib/report/url.test.ts`:

```ts
describe('the compared player', () => {
  it('reads vs as a GUID and writes it back beside with and cmetric', () => {
    const state = parseReportState('?fight=3&mode=compare&with=4&cmetric=hps&vs=Player-4184-000000A3', 1);
    expect(state.compareVs).toBe('Player-4184-000000A3');
    expect(reportSearch(state, 1)).toContain('vs=Player-4184-000000A3');
  });

  it('defaults to no player, and ignores a vs that is not a GUID', () => {
    expect(parseReportState('?fight=3', 1).compareVs).toBe('');
    expect(parseReportState('?fight=3&vs=Morrowlyn the Great', 1).compareVs).toBe('');
    expect(reportSearch(parseReportState('?fight=3', 1), 1)).not.toContain('vs=');
  });
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/url.test.ts`
Expected: FAIL — `expected undefined to be 'Player-4184-000000A3'`.

- [ ] **Step 3: Implement**

In `web/src/lib/report/url.ts`, in `ReportState` beside `compareMetric`:

```ts
  /** Compare mode's second player, as a GUID; '' compares two fights instead. */
  compareVs: string;
```

in `defaultState`:

```ts
    compareVs: '',
```

in `parseReportState`, after the `cmetric` block:

```ts
  // The same shape `source` accepts: a GUID is the engine's own format, a kind then
  // hyphen-separated numeric fields. A name typed by hand is not resolved here the way
  // `source` resolves one, because Compare's picker is a list of this pull's players and
  // a link that names a player who was not in it has nothing to fall back to.
  const compareVs = params.get('vs');
  if (compareVs !== null && /^[A-Za-z0-9-]{1,64}$/.test(compareVs)) state.compareVs = compareVs;
```

in `reportSearch`, after the `cmetric` line:

```ts
  if (state.compareVs !== '') params.set('vs', state.compareVs);
```

- [ ] **Step 4: Run the tests to verify they pass**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/url.test.ts && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"`
Expected: PASS, `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/lib/report/url.ts src/lib/report/url.test.ts && npx prettier --write src/lib/report/url.ts src/lib/report/url.test.ts
cd /Users/jh/code/forever && git add web/src/lib/report/url.ts web/src/lib/report/url.test.ts
git commit -m "feat(web): the url carries Compare's second player as vs"
```

---

### Task 7: Compare expands to abilities, compares two players, and copies a CSV

**Files:**
- Modify: `web/src/components/report/CompareMode.svelte`
- Modify: `web/src/components/report/ReportView.svelte` (pass `vs` through)
- Modify: `web/src/components/report/Glossary.svelte` (one term)
- Create: `web/tests/e2e/report-compare.spec.ts`
- Test: `web/tests/e2e/report-phone.spec.ts`

**Interfaces:**
- Consumes: `abilityDiff`, `playerAbilityDiff`, `metricTable`, `METRIC_LABELS`,
  `PER_SECOND_METRICS`, `COMPARE_METRICS`, `type CompareMetric` from `compare.ts`;
  `state.compareVs` from Task 6; `toCsv` through `CopyCsv`.
- Produces, on `CompareMode`'s props:

```ts
    /** The compared player's GUID from the url; '' compares two fights. */
    vs?: string;
    onPatch: (patch: {
      compareWith?: number | null;
      compareMetric?: string;
      compareVs?: string;
    }) => void;
```

- [ ] **Step 1: Write the failing e2e tests**

```ts
// web/tests/e2e/report-compare.spec.ts
// Compare's two questions. "What did the player above me do differently" is one player's
// abilities in two pulls, which is a row expanding. "What did I do differently from them"
// is two players in this pull, which is the second picker.
import { expect, test } from '@playwright/test';

const COMPARE = '/reports/fixture2abcd?fight=3&mode=compare';

test('a player row expands to that player’s abilities in both pulls', async ({ page }) => {
  await page.goto(`${COMPARE}&with=4`);
  const row = page.getByTestId('compare-Player-4184-000000A1');
  await expect(row).toBeVisible();
  await row.getByRole('button', { name: /abilities/i }).click();
  const diff = page.getByTestId('compare-abilities-Player-4184-000000A1');
  await expect(diff).toBeVisible();
  // Baelgrim's Slam is 4,400 in fight 3 and 482,100 in fight 4.
  await expect(diff).toContainText('Slam');
  await expect(diff.getByTestId('ability-delta').first()).toContainText('−');
});

test('an ability only one pull has shows a dash on the other side', async ({ page }) => {
  await page.goto(`${COMPARE}&with=1`);
  const row = page.getByTestId('compare-Player-4184-000000A1');
  await row.getByRole('button', { name: /abilities/i }).click();
  await expect(page.getByTestId('compare-abilities-Player-4184-000000A1')).toContainText('—');
});

test('threat has no ability split, and the row says so rather than opening empty', async ({ page }) => {
  await page.goto(`${COMPARE}&with=4&cmetric=threat`);
  const row = page.getByTestId('compare-Player-4184-000000A1');
  await row.getByRole('button', { name: /abilities/i }).click();
  await expect(page.getByTestId('compare-abilities-Player-4184-000000A1')).toContainText(
    'no ability split',
  );
});

test('picking a player compares the two inside this pull, and rides in the url', async ({ page }) => {
  await page.goto(COMPARE);
  await page.getByTestId('compare-vs').selectOption('Player-4184-000000A3');
  await expect(page).toHaveURL(/vs=Player-4184-000000A3/);
  // The player table collapses to the two of them.
  const rows = page.getByTestId('compare-cards').locator('li');
  await expect(rows).toHaveCount(2);
  // And the ability diff below is Slam against Frostbolt.
  const diff = page.getByTestId('compare-players');
  await expect(diff).toContainText('Slam');
  await expect(diff).toContainText('Frostbolt');
});

test('the link alone opens on two players', async ({ page }) => {
  await page.goto(`${COMPARE}&vs=Player-4184-000000A3`);
  await expect(page.getByTestId('compare-vs')).toHaveValue('Player-4184-000000A3');
  await expect(page.getByTestId('compare-players')).toContainText('Frostbolt');
});

test('whatever is on screen copies as a CSV', async ({ page }) => {
  await page.goto(`${COMPARE}&with=4`);
  await expect(page.getByTestId('compare-mode').getByTestId('copy-csv').first()).toBeVisible();
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run:
```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-compare.spec.ts --project=desktop
```
Expected: FAIL — no `abilities` button, no `compare-vs` picker.

- [ ] **Step 3: Implement**

In `web/src/components/report/CompareMode.svelte`:

Replace the local `CompareMetric` type, `PER_SECOND` set, `METRIC_LABELS` and `METRIC_IDS`
with imports, and add the diff builders and `CopyCsv`:

```ts
  import {
    COMPARE_METRICS,
    METRIC_LABELS,
    PER_SECOND_METRICS,
    abilityDiff,
    metricTable,
    playerAbilityDiff,
    type CompareMetric,
  } from '../../lib/report/compare';
  import CopyCsv from './CopyCsv.svelte';
```

```ts
  const metric = $derived<CompareMetric>(
    (COMPARE_METRICS as readonly string[]).includes(metricParam)
      ? (metricParam as CompareMetric)
      : 'dps',
  );
```

and every `PER_SECOND.has(metric)` becomes `PER_SECOND_METRICS.has(metric)`.

Add the props:

```ts
    vs = '',
```

```ts
    /** The compared player's GUID from the url; '' compares two fights. */
    vs?: string;
    onPatch: (patch: {
      compareWith?: number | null;
      compareMetric?: string;
      compareVs?: string;
    }) => void;
```

Add the player mode's state and lines. `lines` already exists for the two-fight table;
narrow it when a player is picked, and build the two diffs:

```ts
  /** The players this pull's roster offers as a second player, the picked one excluded. */
  const vsOptions = $derived(
    left.roster
      .filter((row) => row.guid !== topPlayer)
      .map((row) => ({ guid: row.guid, name: splitUnitName(row.name).name }))
      .sort((a, b) => a.name.localeCompare(b.name)),
  );
  /**
   * The player the picked one is compared against. With a player picked the question is
   * "what did I do differently from them", so the first side is the pull's own top row of
   * the metric — the player a reader is looking at — unless the url named it.
   */
  const topPlayer = $derived(lines[0]?.guid ?? '');
  /** True while the mode is two players inside this pull rather than two pulls. */
  const versusPlayer = $derived(vs !== '' && left.roster.some((row) => row.guid === vs));
  /** The player table, cut to the two players when one is picked. */
  const shownLines = $derived(
    versusPlayer ? lines.filter((line) => line.guid === topPlayer || line.guid === vs) : lines,
  );
  /** The ability rows the reader opened, by GUID: a row is opened, not a tab. */
  let openRows = $state<string[]>([]);
  function toggleRow(guid: string): void {
    openRows = openRows.includes(guid) ? openRows.filter((entry) => entry !== guid) : [...openRows, guid];
  }
  const splitTable = $derived(metricTable(metric));
  /** Two players inside this pull, ability by ability. */
  const versusRows = $derived(versusPlayer ? playerAbilityDiff(left, topPlayer, vs, metric) : []);
  function rowsFor(guid: string): AbilityDiff[] {
    return abilityDiff(left, right, guid, metric);
  }
  /**
   * An ability split inside a window is prorated: window.ts scales each ability by the
   * window's share of its actor's total, because the summary keeps no per-ability series.
   * So the ability table's own figures carry the tilde whenever a window is set, whatever
   * the metric -- which is a different mark from `threatScaled`, the whole-table one that
   * only applies to the threat metrics.
   */
  const splitScaled = $derived(leftWindow !== null || rightWindow !== null);
  const splitMark = $derived(splitScaled ? '~' : '');
  /** A signed difference in the ability table, marked when the split it came from was scaled. */
  const splitSigned = (value: number): string =>
    `${value >= 0 ? '+' : '−'}${splitMark}${formatAmount(Math.abs(value))}`;
```

`leftWindow` and `rightWindow` do not exist until Task 21 adds the phase picker; until
then they are the one `window` prop, so declare them now as

```ts
  const leftWindow = $derived(window);
  const rightWindow = $derived(window);
```

and Task 21 replaces both bodies. That keeps `splitScaled` written once.

`lines` must be computed before `topPlayer` reads it, and `vsOptions` reads `topPlayer`;
Svelte's `$derived` is lazy, so declaration order does not matter, but keep them together.
Import `type AbilityDiff` alongside the functions.

When a player is picked the two-fight comparison is not what is on screen, so the second
fight is not needed for it; `right` stays whatever it was and the per-fight columns are
replaced by the two players' panel. The render's three branches become four:

```svelte
  {#if error !== ''}
    <p class="text-[14px]" role="alert">{error}</p>
  {:else if versusPlayer}
    <!-- Branch A below: the two players' panel. -->
  {:else if right === null}
    <p class="text-muted text-[14px]">Pick a second fight, or a second player, to see the difference.</p>
  {:else}
    <!-- Branch B below: the two fights' cards and table, each row now expandable. -->
  {/if}
```

Add the player picker beside "Compare with":

```svelte
    <label class="label text-muted flex w-full flex-wrap items-center gap-2 md:w-auto" for="compare-vs">
      Or a player
      <select
        id="compare-vs"
        class="border-line-warm bg-raised rounded-control text-text h-11 w-full max-w-full min-w-0 px-2 text-[13px] md:h-9 md:w-auto"
        data-testid="compare-vs"
        value={versusPlayer ? vs : ''}
        onchange={(event) =>
          onPatch({ compareVs: (event.currentTarget as HTMLSelectElement).value })}
      >
        <option value="">Compare fights instead</option>
        {#each vsOptions as option (option.guid)}
          <option value={option.guid}>{option.name}</option>
        {/each}
      </select>
    </label>
```

One snippet renders an ability diff, so the row expansion and the two-player panel share
it rather than repeating a table:

```svelte
{#snippet abilityTable(rows: AbilityDiff[], aHead: string, bHead: string, testid: string)}
  <div class="flex flex-col gap-1" data-testid={testid}>
    {#if splitTable === null}
      <p class="text-muted text-[13px]">
        Threat has no ability split: the engine keeps it per player, not per spell. Pick a damage or
        healing metric to read the abilities.
      </p>
    {:else if rows.length === 0}
      <p class="text-muted text-[13px]">Neither side used an ability of this kind here.</p>
    {:else}
      <!-- A grid, not a table: three numbers and a name at 360px read better stacked than
           they do as columns that must each keep a header. -->
      <div class="text-muted label hidden grid-cols-[minmax(0,1fr)_96px_96px_96px] gap-x-3 px-2 pb-1 md:grid">
        <span>Ability</span>
        <span class="text-right">{aHead}</span>
        <span class="text-right">{bHead}</span>
        <span class="text-right">Difference</span>
      </div>
      <ul class="flex flex-col">
        {#each rows as row (row.key)}
          <li
            class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-0.5 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(0,1fr)_96px_96px_96px]"
          >
            <span class="truncate"
              >{row.name}{#if row.via}<span class="text-muted ml-1 text-[11px]">· {row.via}</span>{/if}</span
            >
            <span class="tabular text-right font-mono md:col-start-2"
              >{row.a === null ? '—' : splitMark + formatAmount(row.a)}<span
                class="label font-body text-muted ml-1.5 md:hidden">{aHead}</span
              ></span
            >
            <span class="text-muted tabular col-start-1 text-left font-mono md:col-start-3 md:text-right"
              >{row.b === null ? '—' : splitMark + formatAmount(row.b)}<span
                class="label font-body ml-1.5 md:hidden">{bHead}</span
              ></span
            >
            <span
              class="tabular text-right font-mono"
              class:text-gold={row.delta >= 0}
              data-testid="ability-delta">{splitSigned(row.delta)}</span
            >
          </li>
        {/each}
      </ul>
      <CopyCsv lines={() => abilityCsv(rows, aHead, bHead)} />
      {#if splitScaled}
        <p class="text-muted text-[12px]" data-testid="compare-split-note">
          Marked ~: inside a window the split by ability is each ability's whole-fight amount
          scaled to the window's share of that player's total, not the window's own events. The
          player totals above are exact.
        </p>
      {/if}
    {/if}
  </div>
{/snippet}
```

with the CSV builder beside `signed`:

```ts
  /** An ability diff as lines, in the same shape as every other table's CSV. */
  function abilityCsv(rows: AbilityDiff[], aHead: string, bHead: string): string[][] {
    return [
      ['Ability', 'Via', aHead, bHead, 'Difference'],
      ...rows.map((row) => [
        row.name,
        row.via ?? '',
        row.a === null ? '' : String(row.a),
        row.b === null ? '' : String(row.b),
        String(row.delta),
      ]),
    ];
  }
  /** The player table as lines, for the CSV of what is on screen. */
  function playerCsv(): string[][] {
    return [
      ['Player', 'This fight', 'Compared with', 'Difference'],
      ...shownLines.map((line) => [
        splitUnitName(line.name).name,
        String(line.a),
        String(line.b),
        String(line.a - line.b),
      ]),
    ];
  }
```

One snippet renders a phone card, so both branches show the same card rather than two
copies of the same markup that can drift:

```svelte
{#snippet playerCard(line: Line)}
  <li
    class="border-line-soft grid grid-cols-[minmax(0,1fr)_auto] gap-x-3 gap-y-0.5 border-b py-2 text-[14px]"
    data-testid={`compare-card-${line.guid}`}
  >
    <span class="truncate font-semibold" style={`color: ${classColorVar(line.class)}`}
      >{splitUnitName(line.name).name}</span
    >
    <span
      class="tabular text-right font-mono"
      class:text-gold={line.a >= line.b}
      title="This fight less the compared fight"
      data-testid="compare-card-delta">{signed(line.a - line.b)}</span
    >
    <span class="text-muted col-span-2 text-[12px]"
      ><span class="tabular font-mono">{mark}{formatAmount(line.a)}</span> this fight ·
      <span class="tabular font-mono">{mark}{formatAmount(line.b)}</span> compared with</span
    >
    <span class="col-span-2">
      <button
        type="button"
        class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase"
        aria-expanded={openRows.includes(line.guid)}
        onclick={() => toggleRow(line.guid)}
        >{openRows.includes(line.guid) ? 'Hide abilities' : 'Abilities'}</button
      >
    </span>
    {#if openRows.includes(line.guid)}
      <span class="col-span-2">
        {@render abilityTable(
          rowsFor(line.guid),
          'This fight',
          'Compared with',
          `compare-abilities-${line.guid}`,
        )}
      </span>
    {/if}
  </li>
{/snippet}
```

The existing phone list becomes `{#each shownLines as line (line.guid)}{@render
playerCard(line)}{/each}`, and the desktop `<tbody>` gains the same expander as a fifth
cell and the snippet as a second row:

```svelte
          {#each shownLines as line (line.guid)}
            <tr class="border-line-soft min-h-11 border-b" data-testid={`compare-${line.guid}`}>
              <!-- The four existing cells (name, this fight, compared with, difference)
                   are unchanged. -->
              <td class="px-2 py-2 text-right">
                <button
                  type="button"
                  class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-0"
                  aria-expanded={openRows.includes(line.guid)}
                  onclick={() => toggleRow(line.guid)}
                  >{openRows.includes(line.guid) ? 'Hide abilities' : 'Abilities'}</button
                >
              </td>
            </tr>
            {#if openRows.includes(line.guid)}
              <tr class="bg-card-top">
                <td colspan="5" class="px-2 py-3">
                  {@render abilityTable(
                    rowsFor(line.guid),
                    'This fight',
                    'Compared with',
                    `compare-abilities-${line.guid}`,
                  )}
                </td>
              </tr>
            {/if}
          {/each}
```

The header row gains a fifth `<th scope="col"><span class="sr-only">Abilities</span></th>`.
Both lists iterate `shownLines`, not `lines`.

Branch A, when a player is picked, is the whole content:

```svelte
  {:else if versusPlayer}
    <p class="text-muted text-[12px]" data-testid="compare-players-scope">
      {splitUnitName(left.roster.find((row) => row.guid === topPlayer)?.name ?? '').name} against
      {splitUnitName(left.roster.find((row) => row.guid === vs)?.name ?? '').name}, in this pull{window ===
      null
        ? ''
        : `'s ${formatDuration(window.startMs)} to ${formatDuration(window.endMs)}`}.
    </p>
    <ul class="flex flex-col" data-testid="compare-cards">
      {#each shownLines as line (line.guid)}{@render playerCard(line)}{/each}
    </ul>
    {@render abilityTable(
      versusRows,
      splitUnitName(left.roster.find((row) => row.guid === topPlayer)?.name ?? '').name,
      splitUnitName(left.roster.find((row) => row.guid === vs)?.name ?? '').name,
      'compare-players',
    )}
    <CopyCsv lines={playerCsv} label="Copy the two players as CSV" />
```

Branch B keeps its `<ul data-testid="compare-cards">` (now rendering `playerCard`) and its
table, and adds `<CopyCsv lines={playerCsv} />` under them.

In `web/src/components/report/ReportView.svelte`, pass the url value through:

```svelte
          vs={state.compareVs}
```

In `web/src/components/report/Glossary.svelte`, add:

```ts
    {
      term: 'Difference (Compare)',
      meaning:
        'This fight’s figure less the compared one’s: a plus means this fight did more. Expanding a row gives the same three columns ability by ability, and an ability only one side used shows a dash on the other. Threat has no ability split, because the engine keeps it per player.',
    },
```

- [ ] **Step 4: Run the tests to verify they pass**

Run:
```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-compare.spec.ts --project=desktop && npx playwright test tests/e2e/report-phone.spec.ts --project=mobile && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: PASS on every case, no sideways scroll at 360px on `mode=compare`, `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/components/report/CompareMode.svelte src/components/report/ReportView.svelte src/components/report/Glossary.svelte tests/e2e/report-compare.spec.ts && npx prettier --write src/components/report/CompareMode.svelte src/components/report/ReportView.svelte src/components/report/Glossary.svelte tests/e2e/report-compare.spec.ts
cd /Users/jh/code/forever && git add web/src/components/report web/tests/e2e/report-compare.spec.ts
git commit -m "feat(web): Compare at ability level, player against player, with a CSV"
```

- [ ] **Step 6: Group 2 review gate**

```bash
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && git status --short web/src/fixtures/report
npx vitest run && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?" && npx playwright test tests/e2e/report-compare.spec.ts tests/e2e/report-tabs.spec.ts --project=desktop
```
The fixture regeneration must produce no diff: this group changed no engine code. Merge
`compare-abilities` into `main`.

---

## Group 3 — Resources: cap, time at maximum, wasted at cap, CSV

Tasks 8 to 10. Review gate at the end of Task 10.

---

### Task 8: The engine records the cap, the time at it, and what was wasted past it

**Files:**
- Modify: `logs/engine/summary/deaths.go` (`ResourceTrack`, `addResources`, `resourceRows`)
- Create: `logs/engine/summary/resources_test.go`
- Modify: `web/src/fixtures/report/fixture.log` (one overcapping energize, so the fixture has a track at its cap)
- Regenerate: `logs/engine/summary/testdata/v16.summary.json.golden`, `web/src/fixtures/report/`

**Interfaces:**
- Consumes: `event.Event.PowerType`, `.MaxPower`, `.OverEnergize` (energize lines) and
  `event.Advanced.PowerType`, `.CurrentPower`, `.MaxPower` (every advanced block).
- Produces, on `summary.ResourceTrack`:

```go
	// Max is the largest maximum the log reported for this power: the cap the
	// graph draws a line at. Zero when no line ever carried one.
	Max int64 `json:"max"`
	// AtMaxMS is the whole seconds the reading sat at Max, times 1000, on the
	// same buckets Series uses: the time a rage bar or an energy bar was full
	// and everything poured into it was poured away.
	AtMaxMS int64 `json:"at_max_ms"`
	// Wasted is the power the client says was gained past the cap, summed over
	// this track's energize lines.
	Wasted int64 `json:"wasted"`
```

**Decision taken here:** `AtMaxMS` is read off the finished series at Snapshot rather than
counted as events arrive, so it means exactly what the drawn line means — a second with no
reading carries the last one forward, and a bar that was full through a quiet stretch was
full. The engine version stays `0.4.0`, set in Task 1.

- [ ] **Step 1: Write the failing test**

```go
// logs/engine/summary/resources_test.go
package summary

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
)

// energize is one power gain: what was gained, what was gained past the cap, and the
// reading the advanced block leaves the bar at.
func energize(sec float64, guid string, powerType, amount, over, current, max int64) event.Event {
	return event.Event{
		Time: at(sec), Kind: event.Energize, Name: "SPELL_ENERGIZE",
		Source: event.Unit{GUID: guid, Flags: 0x512}, Dest: event.Unit{GUID: guid, Flags: 0x512},
		Spell:        event.Spell{ID: 34428, Name: "Victory Rush"},
		Amount:       event.OptInt{V: amount, OK: true},
		OverEnergize: event.OptInt{V: over, OK: true},
		PowerType:    event.OptInt{V: powerType, OK: true},
		MaxPower:     event.OptInt{V: max, OK: true},
		Adv: event.Advanced{OK: true, InfoGUID: guid, PowerType: powerType,
			CurrentPower: current, MaxPower: max},
	}
}

// A rage bar that fills and stays full is the thing this table exists to show: every
// point poured into it after that is a point thrown away, and the seconds it spent full
// are seconds the player had a resource they were not spending.
func TestResourcesRecordTheCapTheTimeAtItAndWhatWasWastedPastIt(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		energize(0, tank, 1, 60, 0, 60, 100),
		energize(1, tank, 1, 40, 25, 100, 100),
		energize(2, tank, 1, 10, 10, 100, 100),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(3),
		Players: []string{tank}}, "test")

	if len(s.Resources) != 1 {
		t.Fatalf("resources = %+v", s.Resources)
	}
	tr := s.Resources[0]
	if tr.Max != 100 {
		t.Errorf("max = %d, want 100 (the largest maximum any line reported)", tr.Max)
	}
	// Two of the three seconds read 100 of 100.
	if tr.AtMaxMS != 2000 {
		t.Errorf("at_max_ms = %d, want 2000", tr.AtMaxMS)
	}
	if tr.Wasted != 35 {
		t.Errorf("wasted = %d, want 35 (25 and 10 past the cap)", tr.Wasted)
	}
	if tr.Gained != 110 {
		t.Errorf("gained = %d, want 110", tr.Gained)
	}
}

// A track the log never gave a maximum for has no cap to draw and no time at one: zero,
// not the peak standing in for a cap nobody reported.
func TestResourcesWithNoReportedMaximumHaveNoCapAndNoTimeAtIt(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	e := event.Event{
		Time: at(0), Kind: event.Damage, Name: "SPELL_DAMAGE",
		Source: event.Unit{GUID: mage, Flags: 0x512}, Dest: event.Unit{GUID: boss, Flags: 0xa48},
		Spell:    event.Spell{ID: 116, Name: "Frostbolt", School: 0x10},
		Amount:   event.OptInt{V: 100, OK: true},
		Overkill: event.OptInt{V: -1, OK: true},
		Adv:      event.Advanced{OK: true, InfoGUID: mage, PowerType: 0, CurrentPower: 500, MaxPower: 0},
	}
	reg.Observe(e)
	a.Add(e)
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(1),
		Players: []string{mage}}, "test")
	for _, tr := range s.Resources {
		if tr.Max != 0 || tr.AtMaxMS != 0 {
			t.Errorf("track %+v: with no reported maximum both max and at_max_ms must be zero", tr)
		}
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/forever/logs && go test ./engine/summary/ -run TestResources`
Expected: FAIL to compile — `tr.Max undefined (type ResourceTrack has no field or method Max)`.

- [ ] **Step 3: Implement**

In `logs/engine/summary/deaths.go`, extend `ResourceTrack`:

```go
// ResourceTrack is one actor's power over the fight.
type ResourceTrack struct {
	GUID      string  `json:"guid"`
	Name      string  `json:"name"`
	PowerType int64   `json:"power_type"`
	Series    []int64 `json:"series"`
	Gained    int64   `json:"gained"`
	Spent     int64   `json:"spent"`
	ZeroMS    int64   `json:"zero_ms"`
	// Max is the largest maximum the log reported for this power: the cap the
	// graph draws a line at. Zero when no line ever carried one.
	Max int64 `json:"max"`
	// AtMaxMS is the whole seconds the reading sat at Max, times 1000, on the
	// same buckets Series uses: the time a rage bar or an energy bar was full
	// and everything poured into it was poured away.
	AtMaxMS int64 `json:"at_max_ms"`
	// Wasted is the power the client says was gained past the cap, summed over
	// this track's energize lines.
	Wasted int64 `json:"wasted"`
}
```

In `addResources`, the energize branch takes the cap and the waste, and the advanced
branch takes the cap wherever a line carries one:

```go
	if e.Kind == event.Energize && e.Dest.GUID != "" && e.Dest.GUID != units.NoGUID {
		k := resourceKey{guid: e.Dest.GUID, powerType: e.PowerType.V}
		tr := a.resource(k, e.Time)
		tr.Gained += e.Amount.V
		// What the client says was gained past the cap. Negative would be a
		// malformed line, and a negative waste is not a thing to report.
		tr.Wasted += max(e.OverEnergize.V, 0)
		if e.MaxPower.V > tr.Max {
			tr.Max = e.MaxPower.V
		}
	}
	if !e.Adv.OK || e.Adv.InfoGUID == "" || e.Adv.InfoGUID == units.NoGUID {
		return
	}
	k := resourceKey{guid: e.Adv.InfoGUID, powerType: e.Adv.PowerType}
	tr := a.resource(k, e.Time)
	if e.Adv.MaxPower > tr.Max {
		tr.Max = e.Adv.MaxPower
	}
```

(the rest of `addResources` is unchanged; the old `a.resource(k, e.Time).Gained += …`
one-liner is what the first block replaces.)

In `resourceRows`, fill the derived figure after the series is copied:

```go
func (a *Accumulator) resourceRows() []ResourceTrack {
	out := make([]ResourceTrack, 0, len(a.resources))
	for _, tr := range a.resources {
		row := tr.ResourceTrack
		row.Series = copySlice(tr.Series)
		if row.Series == nil {
			row.Series = []int64{}
		}
		row.AtMaxMS = atMaxMS(row.Series, row.Max, a.opt.Bucket)
		out = append(out, row)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].GUID != out[j].GUID {
			return out[i].GUID < out[j].GUID
		}
		return out[i].PowerType < out[j].PowerType
	})
	return out
}

// atMaxMS is the time the series sat at the cap: one bucket per second whose
// reading is the maximum. Read off the finished series rather than counted as
// the events arrive, so it means what the drawn line means -- a second with no
// reading carries the last one forward, and a bar that was full through a quiet
// stretch was full. A track the log never gave a maximum for has no cap to be at.
func atMaxMS(series []int64, maximum int64, bucket time.Duration) int64 {
	if maximum <= 0 {
		return 0
	}
	var total int64
	for _, value := range series {
		if value >= maximum {
			total += bucket.Milliseconds()
		}
	}
	return total
}
```

- [ ] **Step 4: Give the web fixture a bar that reaches its cap**

Add one line to `web/src/fixtures/report/fixture.log`, between the
`20:12:16` Wrack Soul line and the `20:12:20` Frostbolt line, so Baelgrim's rage fills
and overcaps 18 seconds into fight 3:

```
9/26 20:12:18.000  SPELL_ENERGIZE,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,34428,"Victory Rush",0x1,Player-4184-000000A1,0000000000000000,9800,11000,388,0,4120,0,1,1000,1000,0,-1489.90,6410.05,1675,4.1002,183,60.0000,25.0000,1,1000
```

It adds no damage, no heal, no cast and no aura, so every figure any existing test pins is
untouched; it adds one resource track (Baelgrim, power type 1) reading 1000 of 1000.

- [ ] **Step 5: Run the tests, regenerate the goldens and the fixture**

```bash
cd /Users/jh/code/forever/logs && go test ./engine/summary/ -run TestResources && FOREVER_UPDATE_GOLDEN=1 go test ./... && go test ./... && (cd ../api && go test ./...)
export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH"
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && npx vitest run src/fixtures/report/fixture.test.ts
```
Expected: PASS everywhere; `web/src/fixtures/report/fights/3/summary.json` now holds a
`power_type: 1` track for `Player-4184-000000A1` with `"max": 1000`, `"at_max_ms": 1000`
and `"wasted": 25`. Read those three values out of the file and note them: Task 10's e2e
asserts on them.

- [ ] **Step 6: Commit**

```bash
git add logs/engine/summary web/src/fixtures/report
git commit -m "feat(logs): resources carry their cap, the time at it and what was wasted past it"
```

---

### Task 9: The web scopes the cap figures to the window and folds them over the night

**Files:**
- Modify: `web/src/lib/report/types.ts` (`ResourceTrack` gains `max`, `at_max_ms`, `wasted`)
- Modify: `web/src/lib/report/window.ts` (`scopeResource` recomputes the time at cap)
- Modify: `web/src/lib/report/night.ts` (the resource fold)
- Test: `web/src/lib/report/window.test.ts`, `web/src/lib/report/night.test.ts`

**Interfaces:**
- Consumes: the JSON shape from Task 8.
- Produces, on `ResourceTrack` in `types.ts`:

```ts
  /** The largest maximum the log reported for this power: the cap the graph draws. Absent before engine 0.4.0. */
  max?: number;
  /** Whole seconds the reading was at the cap, times 1000. Recomputed by the window scope from the sliced series. */
  at_max_ms?: number;
  /** Power gained past the cap. The whole fight's figure under any window: the summary cannot cut it down. */
  wasted?: number;
```

- Produces, exported from `window.ts`: nothing new — `scopeResource` keeps its signature.

**Decision taken here:** the series is exact under a window, so `at_max_ms` is recomputed
from the sliced series and is a measured figure with no mark. `wasted` cannot be cut down
by the summary, so it keeps the whole-fight dagger the way `zero_ms`, `gained` and `spent`
already do.

- [ ] **Step 1: Write the failing tests**

In `web/src/lib/report/window.test.ts`, append:

```ts
describe('a resource track in a window', () => {
  const track = {
    guid: 'P1',
    name: 'P1',
    power_type: 1,
    series: [0, 100, 100, 40, 100],
    gained: 300,
    spent: 60,
    zero_ms: 1000,
    max: 100,
    at_max_ms: 3000,
    wasted: 25,
  };

  it('recomputes the time at the cap from the window’s own buckets', () => {
    const scoped = scopeResource(track, { startMs: 1000, endMs: 4000 });
    expect(scoped.series).toEqual([100, 100, 40]);
    expect(scoped.at_max_ms).toBe(2000);
  });

  it('leaves the cap and the waste alone: one is a property of the bar, the other whole-fight', () => {
    const scoped = scopeResource(track, { startMs: 1000, endMs: 4000 });
    expect(scoped.max).toBe(100);
    expect(scoped.wasted).toBe(25);
    expect(scoped.zero_ms).toBe(1000);
  });

  it('leaves a track written before the engine kept a cap untouched', () => {
    const { max: _max, at_max_ms: _atMax, wasted: _wasted, ...old } = track;
    const scoped = scopeResource(old, { startMs: 1000, endMs: 4000 });
    expect(scoped.max).toBeUndefined();
    expect(scoped.at_max_ms).toBeUndefined();
  });
});
```

In `web/src/lib/report/night.test.ts`, append inside the nightSummary describe:

```ts
  it('folds the cap figures: wasted and at-cap sum, the cap is the largest of them', () => {
    const fights = [fight(3, 'Kaal', false, 4000), fight(4, 'Kaal', true, 4000)];
    const track = (max: number, atMax: number, wasted: number, series: number[]) => [
      {
        guid: 'Player-1',
        name: 'Tank',
        power_type: 1,
        series,
        gained: 10,
        spent: 5,
        zero_ms: 0,
        max,
        at_max_ms: atMax,
        wasted,
      },
    ];
    const summaries = new Map([
      [3, { ...summary(3, 4000, [roster('A', 1)]), resources: track(100, 2000, 25, [100, 100]) }],
      [4, { ...summary(4, 4000, [roster('A', 1)]), resources: track(120, 1000, 5, [120]) }],
    ]);
    const night = nightSummary(fights, summaries as never);
    const folded = night.resources.find((row) => row.guid === 'Player-1');
    expect(folded?.max).toBe(120);
    expect(folded?.at_max_ms).toBe(3000);
    expect(folded?.wasted).toBe(30);
  });
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/window.test.ts src/lib/report/night.test.ts`
Expected: FAIL — `expected undefined to be 2000` on the window case, `expected undefined to be 120` on the night case.

- [ ] **Step 3: Implement**

In `web/src/lib/report/types.ts`:

```ts
/** summary.ResourceTrack. `power_type` is the game's power index (0 mana, 1 rage, 3 energy). */
export interface ResourceTrack {
  guid: string;
  name: string;
  power_type: number;
  series: number[];
  gained: number;
  spent: number;
  zero_ms: number;
  /** The largest maximum the log reported for this power: the cap the graph draws. Absent before engine 0.4.0. */
  max?: number;
  /** Whole seconds the reading was at the cap, times 1000. Recomputed by the window scope from the sliced series. */
  at_max_ms?: number;
  /** Power gained past the cap. The whole fight's figure under any window: the summary cannot cut it down. */
  wasted?: number;
}
```

In `web/src/lib/report/window.ts`:

```ts
/**
 * The series is exact under a window -- this only slices it -- so the time at the cap is
 * recomputed from the window's own buckets and stays a measured figure. `wasted`,
 * `gained`, `spent` and `zero_ms` are not sliceable from the summary and come through
 * untouched, which is the whole fight's figure and is marked as one where it is printed.
 */
export function scopeResource(track: ResourceTrack, window: TimeWindow): ResourceTrack {
  const series = sliceSeries(track.series, window);
  if (track.max === undefined) return { ...track, series };
  return { ...track, series, at_max_ms: atMaxMs(series, track.max) };
}

/** The buckets whose reading is at the cap, as milliseconds: the engine's own rule. */
function atMaxMs(series: number[], max: number): number {
  if (max <= 0) return 0;
  return series.filter((value) => value >= max).length * BUCKET_MS;
}
```

In `web/src/lib/report/night.ts`, extend the resource merge's `else` branch (the
`resources.set(rkey, {...})` that sums `gained`, `spent` and `zero_ms`) and the first
branch so a single-pull night keeps its figures:

```ts
      if (found === undefined) {
        const lead = new Array<number>(Math.max(0, Math.round(padTo))).fill(0);
        resources.set(rkey, { ...track, series: [...lead, ...track.series] });
      } else {
        // `last`, `gap` and `joined` above are unchanged.
        resources.set(rkey, {
          ...found,
          series: [...found.series, ...new Array<number>(gap).fill(last), ...joined],
          gained: found.gained + track.gained,
          spent: found.spent + track.spent,
          zero_ms: found.zero_ms + track.zero_ms,
          // The cap is a property of the bar, not of the night: the largest any pull
          // reported. The time at it and the waste are counts, and counts add up.
          max: maxOptional(found.max, track.max),
          at_max_ms: sumOptional(found.at_max_ms, track.at_max_ms),
          wasted: sumOptional(found.wasted, track.wasted),
        });
      }
```

and add the helper beside `sumOptional`:

```ts
function maxOptional(a: number | undefined, b: number | undefined): number | undefined {
  if (a === undefined && b === undefined) return undefined;
  return Math.max(a ?? 0, b ?? 0);
}
```

- [ ] **Step 4: Run the tests to verify they pass**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/window.test.ts src/lib/report/night.test.ts && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"`
Expected: PASS, `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/lib/report/types.ts src/lib/report/window.ts src/lib/report/night.ts src/lib/report/window.test.ts src/lib/report/night.test.ts && npx prettier --write src/lib/report/types.ts src/lib/report/window.ts src/lib/report/night.ts src/lib/report/window.test.ts src/lib/report/night.test.ts
cd /Users/jh/code/forever && git add web/src/lib/report
git commit -m "feat(web): the resource cap, the time at it and the waste, scoped and folded"
```

---

### Task 10: The resource graph draws the cap, shades the time at it, and copies a CSV

**Files:**
- Modify: `web/src/components/report/ResourceGraphs.svelte`
- Modify: `web/src/components/report/Glossary.svelte` (two terms)
- Test: `web/tests/e2e/report-tables.spec.ts`, `web/tests/e2e/report-phone.spec.ts`

**Interfaces:**
- Consumes: `ResourceTrack.max`, `.at_max_ms`, `.wasted` (Task 9); `wholeFightMark`,
  `wholeFightTitle`, `wholeFightAriaLabel`, `formatAmount`, `formatDuration`, `formatPercent`.
- Produces: `ResourceGraphs` uses `CopyCsv` with a per-track `lines` builder:
  `second, reading, at cap`.

- [ ] **Step 1: Write the failing e2e tests**

In `web/tests/e2e/report-tables.spec.ts`, append:

```ts
test('a resource graph draws the cap, shades the time at it and says what was wasted', async ({
  page,
}) => {
  await page.goto(`${FIGHT}&tab=resources`);
  // Baelgrim's rage fills to 1000 of 1000 eighteen seconds in and stays read at the cap
  // for that one second of the fixture's sixty.
  const rage = page.getByTestId('resource-Player-4184-000000A1-1');
  await expect(rage).toBeVisible();
  await expect(rage.getByTestId('resource-cap-line')).toBeVisible();
  await expect(rage.getByTestId('resource-at-max')).toHaveCount(1);
  const figures = rage.getByTestId('resource-cap-figures');
  await expect(figures).toContainText('peak 1,000');
  await expect(figures).toContainText('at cap');
  await expect(figures).toContainText('wasted 25');
});

test('a resource with no reported cap draws no cap line and no cap figures', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=resources&source=Player-4184-000000A3`);
  const mana = page.getByTestId('resource-Player-4184-000000A3-0');
  await expect(mana).toBeVisible();
  // The mage's mana does have a reported maximum, so this asserts the other half: it
  // never reaches it, so nothing is shaded.
  await expect(mana.getByTestId('resource-at-max')).toHaveCount(0);
  await expect(mana.getByTestId('resource-cap-figures')).toContainText('at cap 0.0%');
});

test('a resource series copies as a CSV of second, reading and at cap', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=resources`);
  await expect(
    page.getByTestId('resource-Player-4184-000000A1-1').getByTestId('copy-csv'),
  ).toBeVisible();
});
```

In `web/tests/e2e/report-phone.spec.ts`, append:

```ts
test('a resource row’s cap figures sit beside the line at 360px', async ({ page }) => {
  await page.goto(`${REPORT}&tab=resources`);
  const figures = page.getByTestId('resource-cap-figures').first();
  await expect(figures).toBeVisible();
  const box = await figures.boundingBox();
  expect(box?.width ?? 0).toBeLessThanOrEqual(PHONE_WIDTH);
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run:
```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-tables.spec.ts --project=desktop
```
Expected: FAIL — `resource-cap-line` not found.

- [ ] **Step 3: Implement**

In `web/src/components/report/ResourceGraphs.svelte`, add `formatPercent` to the imports
from `format.ts` and `CopyCsv` from `./CopyCsv.svelte`, and add the derivations:

```ts
  /** Where the cap sits on the 26-unit viewBox the sparkline is drawn in. */
  function capY(series: number[], max: number): number {
    const peak = peakOf(series);
    const top = Math.max(peak, max);
    return top === 0 ? 24 : 24 - (max / top) * 22;
  }
  /** The seconds whose reading is at the cap, as [from, to] percentages of the line's width. */
  function atMaxSpans(series: number[], max: number): { from: number; to: number }[] {
    if (max <= 0 || series.length === 0) return [];
    const width = 100 / series.length;
    const out: { from: number; to: number }[] = [];
    series.forEach((value, index) => {
      if (value < max) return;
      const from = index * width;
      const last = out[out.length - 1];
      if (last !== undefined && Math.abs(last.to - from) < 0.0001) last.to = from + width;
      else out.push({ from, to: from + width });
    });
    return out;
  }
  /** The share of the window the bar spent full. */
  function atCapPct(atMaxMs: number): number {
    return durationMs === 0 ? 0 : (atMaxMs / durationMs) * 100;
  }
  /** One track's series as lines: the second, the reading, and whether it was at the cap. */
  function seriesCsv(track: ResourceTrack): string[][] {
    return [
      ['Second', 'Reading', 'At cap'],
      ...track.series.map((value, index) => [
        String(index),
        String(value),
        track.max !== undefined && track.max > 0 && value >= track.max ? 'yes' : 'no',
      ]),
    ];
  }
```

`peakOf` already exists above; keep the new functions beside it. The sparkline drawing
must be scaled against the cap as well as the peak, or a line that reaches its cap would
touch the top of the box and the cap line would sit on it: change `points` to take the
cap too.

```ts
  function points(series: number[], max = 0): string {
    const top = Math.max(peakOf(series), max);
    if (top === 0 || series.length < 2) return '';
    return series
      .map((value, index) => `${(index / (series.length - 1)) * 100},${24 - (value / top) * 22}`)
      .join(' ');
  }
```

In the markup, inside the `<svg>` and before the `<polyline>`, add the shading and the cap
line; and after the low/peak figure line, add the cap figures and the CSV button:

```svelte
              {#if track.max !== undefined && track.max > 0}
                {#each atMaxSpans(track.series, track.max) as span, i (`${span.from}-${i}`)}
                  <rect
                    x={span.from}
                    y="0"
                    width={span.to - span.from}
                    height="26"
                    fill="var(--color-gold)"
                    opacity="0.18"
                    data-testid="resource-at-max"
                  />
                {/each}
                <line
                  x1="0"
                  y1={capY(track.series, track.max)}
                  x2="100"
                  y2={capY(track.series, track.max)}
                  stroke="var(--color-gold)"
                  stroke-width="1"
                  stroke-dasharray="3 2"
                  vector-effect="non-scaling-stroke"
                  data-testid="resource-cap-line"
                />
              {/if}
              <polyline
                points={points(track.series, track.max ?? 0)}
                fill="none"
                stroke="var(--color-gold)"
                stroke-width="1.5"
                vector-effect="non-scaling-stroke"
              />
```

`<rect>` and `<line>` are inside an `aria-hidden` svg, so a `data-testid` on them is for
the test alone, which is what the e2e counts. Then, under the existing peak/low line:

```svelte
          {#if track.max !== undefined && track.max > 0}
            <span
              class="text-muted tabular flex flex-wrap gap-x-3 font-mono text-[11px]"
              data-testid="resource-cap-figures"
            >
              <span title="The cap the log reported for this power">cap {formatAmount(track.max)}</span>
              <span title="The share of this window the bar spent full, measured from the window's own seconds"
                >at cap {formatPercent(atCapPct(track.at_max_ms ?? 0))} of the fight</span
              >
              <!-- One title per element: the whole-fight one from wholeFightTitle(true),
                   already bound to this file's `title` const, and the words that explain
                   what the figure is in the aria-label, which is where a title on an
                   element with visible text would not reliably be read out anyway. -->
              <span
                {title}
                aria-label={wholeFightAriaLabel(
                  true,
                  `${formatAmount(track.wasted ?? 0)} of power gained past the cap and thrown away`,
                )}>{mark}wasted {formatAmount(track.wasted ?? 0)}</span
              >
            </span>
          {/if}
          <CopyCsv lines={() => seriesCsv(track)} label="Copy this line as CSV" />
```

The figure line already says "peak N"; the spec's line reads
`peak N · at cap M% of the fight · wasted W`, which is the existing peak/low span followed
by this one. Extend the note at the bottom so the dagger is explained once:

```svelte
  <p class="text-muted text-[12px]" data-testid="resource-wholefight-note">
    Time empty and power wasted are marked {mark} because the summary tracks them for the whole fight
    only; the line, the cap and the time at the cap are this window's own.
  </p>
```

In `web/src/components/report/Glossary.svelte`, add:

```ts
    {
      term: 'At cap',
      meaning:
        'On the Resources tab: the share of the window a bar sat at its maximum. A rage or energy bar at its cap is generating nothing, so time at cap is time a resource was going to waste. Measured from the window’s own seconds.',
    },
    {
      term: 'Wasted',
      meaning:
        'Power gained past the cap, as the client reports it on each gain: the rage a hit would have given had the bar had room. A dagger marks it because the summary keeps it for the whole fight, not for a window.',
    },
```

- [ ] **Step 4: Run the tests to verify they pass**

```bash
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && git status --short src/fixtures/report && npx vitest run src/lib/report src/fixtures/report/fixture.test.ts && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-tables.spec.ts --project=desktop && npx playwright test tests/e2e/report-phone.spec.ts --project=mobile && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: the fixture regeneration produces no diff (Task 8 already did it); every case
passes; `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/components/report/ResourceGraphs.svelte src/components/report/Glossary.svelte tests/e2e/report-tables.spec.ts tests/e2e/report-phone.spec.ts && npx prettier --write src/components/report/ResourceGraphs.svelte src/components/report/Glossary.svelte tests/e2e/report-tables.spec.ts tests/e2e/report-phone.spec.ts
cd /Users/jh/code/forever && git add web/src/components/report web/tests/e2e
git commit -m "feat(web): resource graphs draw the cap, the time at it and the waste"
```

- [ ] **Step 6: Group 3 review gate**

```bash
cd /Users/jh/code/forever/logs && go test ./... && (cd ../api && go test ./...)
cd /Users/jh/code/forever/web && npx vitest run && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?" && npx playwright test tests/e2e/report-tables.spec.ts tests/e2e/report-tabs.spec.ts --project=desktop
```
Merge `resources-at-cap` into `main`.

---

## Group 4 — Pet rows on Casts, and aura refresh lines in Events

Tasks 11 to 14. Review gate at the end of Task 14.

---

### Task 11: Every cast row names its caster's owner

**Files:**
- Modify: `logs/engine/summary/deaths.go` (`CastRow`, `cast`)
- Test: `logs/engine/summary/summary_test.go`
- Modify: `web/src/fixtures/report/fixture.log` (one pet cast, so the fixture has a pet's row to fold)
- Regenerate: `logs/engine/summary/testdata/v16.summary.json.golden`, `web/src/fixtures/report/`

**Interfaces:**
- Consumes: `Accumulator.owner(guid) string`, which reads `units.Registry.Owner`.
- Produces, on `summary.CastRow`:

```go
	// OwnerGUID is the caster's owner when the caster is a pet or a guardian,
	// and the caster's own GUID otherwise, so the Casts tab can keep a pet's
	// rows under the player who owns it the way the damage tables already keep
	// a pet's damage.
	OwnerGUID string `json:"owner_guid"`
```

- [ ] **Step 1: Write the failing test**

Append to `logs/engine/summary/summary_test.go`:

```go
// A statue's casts and a hunter's pet's casts belong on their owner's page: the damage
// tables already fold them there, and a Casts tab that does not is the one place a pet's
// work disappears when a reader picks the player who owns it.
func TestCastRowsNameTheCastersOwner(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		{Time: at(0), Kind: event.Summon, Name: "SPELL_SUMMON",
			Source: event.Unit{GUID: hunter, Name: "Thalgrit-Nightslayer", Flags: 0x512},
			Dest:   event.Unit{GUID: pet, Name: "Ashfang", Flags: 0x1114}},
		cast(1, hunter, boss, 34026, "Kill Command"),
		cast(2, pet, boss, 17253, "Bite"),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(3),
		Players: []string{hunter}}, "test")

	row := func(guid string, spellID int64) CastRow {
		t.Helper()
		for _, r := range s.Casts {
			if r.GUID == guid && r.SpellID == spellID {
				return r
			}
		}
		t.Fatalf("no cast row for %s / %d in %+v", guid, spellID, s.Casts)
		return CastRow{}
	}
	if got := row(pet, 17253).OwnerGUID; got != hunter {
		t.Errorf("the pet's row owner = %q, want the hunter %q", got, hunter)
	}
	if got := row(hunter, 34026).OwnerGUID; got != hunter {
		t.Errorf("the hunter's own row owner = %q, want their own guid %q", got, hunter)
	}
	// The pet keeps its own row and its own name: "via Ashfang" is the web's job, and it
	// needs the two apart.
	if got := row(pet, 17253).Name; got != "Ashfang" {
		t.Errorf("the pet's row name = %q, want Ashfang", got)
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/forever/logs && go test ./engine/summary/ -run TestCastRowsNameTheCastersOwner`
Expected: FAIL to compile — `r.OwnerGUID undefined`.

- [ ] **Step 3: Implement**

In `logs/engine/summary/deaths.go`, add the field to `CastRow` after `Name`:

```go
// CastRow is one caster's use of one spell.
type CastRow struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
	// OwnerGUID is the caster's owner when the caster is a pet or a guardian,
	// and the caster's own GUID otherwise, so the Casts tab can keep a pet's
	// rows under the player who owns it the way the damage tables already keep
	// a pet's damage. The row itself stays the pet's, name and all: the web
	// prints "via Ashfang" and needs the two apart.
	OwnerGUID   string           `json:"owner_guid"`
	SpellID     int64            `json:"spell_id"`
	SpellName   string           `json:"spell_name"`
	Started     int64            `json:"started"`
	Succeeded   int64            `json:"succeeded"`
	Failed      int64            `json:"failed"`
	FailReasons map[string]int64 `json:"fail_reasons,omitempty"`
	CastTimeMS  int64            `json:"cast_time_ms"`
	Sequence    []int64          `json:"sequence"`
}
```

and set it where the row is opened:

```go
func (a *Accumulator) cast(e event.Event) *castRow {
	k := castKey{guid: e.Source.GUID, spellID: e.Spell.ID}
	row, ok := a.casts[k]
	if !ok {
		row = &castRow{}
		row.GUID, row.Name = e.Source.GUID, a.name(e.Source.GUID)
		row.OwnerGUID = a.owner(e.Source.GUID)
		row.SpellID, row.SpellName = e.Spell.ID, e.Spell.Name
		a.casts[k] = row
	}
	return row
}
```

- [ ] **Step 4: Give the web fixture a pet cast**

Add one line to `web/src/fixtures/report/fixture.log`, between the `20:12:01.000`
COMBATANT_INFO for Player-4184-000000A2 and the `20:12:02.000` SPELL_INSTAKILL that
dismisses the pet, so Ashfang casts once inside fight 3 while it is still out:

```
9/26 20:12:01.500  SPELL_CAST_SUCCESS,Pet-0-2085-2284-7855-165189-01000000B1,"Ashfang",0x1111,0x0,Creature-0-2085-2284-7855-169754-0000AA0002,"Warden Kelthas",0xa48,0x0,17253,"Bite",0x1
```

It adds no damage and no aura. It does mark the pet's owner active for one gap (1.5 s),
so `Player-4184-000000A4`'s `active_ms` in fight 3 goes from 0 to 1500 — check no test
pins that before moving on.

- [ ] **Step 5: Run the tests, regenerate the goldens and the fixture**

```bash
cd /Users/jh/code/forever/logs && go test ./engine/summary/ && FOREVER_UPDATE_GOLDEN=1 go test ./... && go test ./... && (cd ../api && go test ./...)
export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH"
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && npx vitest run && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-tabs.spec.ts tests/e2e/report-tables.spec.ts tests/e2e/report-casts.spec.ts tests/e2e/report-threat.spec.ts --project=desktop
```
Expected: PASS. Every `casts` entry in the golden and the fixture now carries
`owner_guid`; `fights/3/summary.json` has a row with
`"guid": "Pet-0-2085-2284-7855-165189-01000000B1"` and
`"owner_guid": "Player-4184-000000A4"`. If a spec pins a cast-row count under the
friendlies scope it will still pass here — the pet's row is filtered out until Task 12.

- [ ] **Step 6: Commit**

```bash
git add logs/engine/summary web/src/fixtures/report
git commit -m "feat(logs): every cast row names its caster's owner"
```

---

### Task 12: The source scope keeps a pet's casts under its owner

**Files:**
- Modify: `web/src/lib/report/types.ts` (`CastRow` gains `owner_guid`)
- Modify: `web/src/lib/report/source.ts` (`scopeSource` scopes casts by owner)
- Test: `web/src/lib/report/source.test.ts`

**Interfaces:**
- Consumes: `CastRow.owner_guid` from Task 11.
- Produces, on `CastRow` in `types.ts`:

```ts
  /** The caster's owner when the caster is a pet, their own GUID otherwise. Absent before engine 0.4.0. */
  owner_guid?: string;
```

- [ ] **Step 1: Write the failing test**

In `web/src/lib/report/source.test.ts`, add casts to the shared `summary()` fixture and
append the cases. The `casts: []` line in the fixture becomes:

```ts
    casts: [
      {
        guid: 'Player-1',
        name: 'One',
        owner_guid: 'Player-1',
        spell_id: 1,
        spell_name: 'Slam',
        started: 1,
        succeeded: 1,
        failed: 0,
        cast_time_ms: 0,
        sequence: [1000],
      },
      {
        guid: 'Pet-7',
        name: 'Ashfang',
        owner_guid: 'Player-2',
        spell_id: 2,
        spell_name: 'Bite',
        started: 1,
        succeeded: 1,
        failed: 0,
        cast_time_ms: 0,
        sequence: [2000],
      },
      {
        guid: 'Creature-9',
        name: 'Boss',
        owner_guid: 'Creature-9',
        spell_id: 3,
        spell_name: 'Cleave',
        started: 1,
        succeeded: 1,
        failed: 0,
        cast_time_ms: 0,
        sequence: [3000],
      },
    ],
```

and the new cases:

```ts
describe('casts under the source scope', () => {
  it('keeps a pet’s casts under the player who owns it', () => {
    const scoped = scopeSource(summary(), 'Player-2', players);
    expect(scoped.casts.map((row) => row.spell_name)).toEqual(['Bite']);
  });

  it('shows a pet’s casts under all friendlies, where the pet itself is not a player', () => {
    const scoped = scopeSource(summary(), 'friendlies', players);
    expect(scoped.casts.map((row) => row.spell_name).sort()).toEqual(['Bite', 'Slam']);
  });

  it('leaves a pet out of the enemies, because its owner is a friendly', () => {
    const scoped = scopeSource(summary(), 'enemies', players, new Set([...players, 'Pet-7']));
    expect(scoped.casts.map((row) => row.spell_name)).toEqual(['Cleave']);
  });

  it('falls back to the caster’s own guid on a report parsed before owners were kept', () => {
    const old = summary();
    old.casts = old.casts.map((row) => ({ ...row, owner_guid: undefined }));
    expect(scopeSource(old, 'Player-2', players).casts).toEqual([]);
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/source.test.ts`
Expected: FAIL — the pet's cast is dropped under `Player-2` and under `friendlies`, so the
first two cases come back `[]` and `['Slam']`.

- [ ] **Step 3: Implement**

In `web/src/lib/report/types.ts`:

```ts
/** summary.CastRow. `sequence` holds millisecond offsets from the fight's start. */
export interface CastRow {
  guid: string;
  name: string;
  /** The caster's owner when the caster is a pet, their own GUID otherwise. Absent before engine 0.4.0. */
  owner_guid?: string;
  spell_id: number;
  // spell_name, started, succeeded, failed, fail_reasons, cast_time_ms and sequence
  // are unchanged.
}
```

In `web/src/lib/report/source.ts`, scope the cast rows by their owner:

```ts
    // A pet's casts are its owner's work, the way a pet's damage is its owner's damage:
    // picking the healer shows the statue's casts too. A row written before the engine
    // kept the owner falls back to the caster's own guid, which is what the scope did
    // before, so an older report narrows exactly as it used to.
    casts: summary.casts.filter((row) => keep(row.owner_guid ?? row.guid)),
```

- [ ] **Step 4: Run the tests to verify they pass**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/source.test.ts && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"`
Expected: PASS, `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/lib/report/types.ts src/lib/report/source.ts src/lib/report/source.test.ts && npx prettier --write src/lib/report/types.ts src/lib/report/source.ts src/lib/report/source.test.ts
cd /Users/jh/code/forever && git add web/src/lib/report
git commit -m "feat(web): a pet's casts scope under the player who owns it"
```

---

### Task 13: The Casts tab says which pet cast it

**Files:**
- Modify: `web/src/components/report/CastTable.svelte`
- Modify: `web/src/components/report/Glossary.svelte` (one term)
- Test: `web/tests/e2e/report-casts.spec.ts`

**Interfaces:**
- Consumes: `CastRow.owner_guid` (Task 12), `classOf` (already a prop), `splitUnitName`.
- Produces: no new props. The Caster cell reads the owner's name and class; the Spell cell
  carries `· via Ashfang` when the row is a pet's, the way `ActorRow` prints an ability's
  `via`.

- [ ] **Step 1: Write the failing e2e tests**

Append to `web/tests/e2e/report-casts.spec.ts`:

```ts
// Ashfang bites once at 1.5s of fight 3. Under all friendlies the row is there, filed
// under its owner and saying which pet it was.
test('a pet’s cast row sits under its owner and says which pet cast it', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=casts');
  const bite = page.getByTestId('cast-Pet-0-2085-2284-7855-165189-01000000B1-17253');
  await expect(bite).toBeVisible();
  await expect(bite).toContainText('Thalgrit');
  await expect(bite).toContainText('via Ashfang');
});

test('picking the pet’s owner keeps the pet’s casts on screen', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=casts&source=Player-4184-000000A4');
  await expect(page.getByTestId('cast-Pet-0-2085-2284-7855-165189-01000000B1-17253')).toBeVisible();
});

test('picking another player leaves the pet’s casts out', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=casts&source=Player-4184-000000A3');
  await expect(page.getByTestId('cast-Pet-0-2085-2284-7855-165189-01000000B1-17253')).toHaveCount(0);
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run:
```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-casts.spec.ts --project=desktop
```
Expected: FAIL — the first case finds the row but it reads "Ashfang" in the Caster column
and has no "via Ashfang" beside the spell.

- [ ] **Step 3: Implement**

In `web/src/components/report/CastTable.svelte`, add a prop for the names (the component
already receives `classOf`; it needs GUID-to-name for the owner, which `ReportView`
already holds as `unitNames`):

```ts
    names = new Map<string, string>(),
```

```ts
    /** GUID to unit name, so a pet's row can be filed under its owner's name. */
    names?: ReadonlyMap<string, string>;
```

and the two derivations:

```ts
  /** The player a row belongs to: a pet's owner, or the caster themselves. */
  const ownerOf = (row: CastRow): string => row.owner_guid ?? row.guid;
  /** The pet's own name when the row is a pet's; '' for a caster's own row. */
  const viaOf = (row: CastRow): string =>
    ownerOf(row) === row.guid ? '' : splitUnitName(row.name).name;
  /** The name the Caster column shows: the owner's, since that is whose page this is. */
  const casterName = (row: CastRow): string =>
    splitUnitName(names.get(ownerOf(row)) ?? (ownerOf(row) === row.guid ? row.name : ownerOf(row))).name;
```

In the row markup, the Caster cell and the Spell cell become:

```svelte
          <span
            class="truncate font-semibold"
            style={`color: ${classColorVar(classOf.get(ownerOf(row)))}`}
            title={casterName(row)}
          >
            {casterName(row)}
          </span>
          <span class="truncate" title={row.spell_name}
            >{row.spell_name}{#if viaOf(row) !== ''}
              <span
                class="text-muted ml-1 text-[11px]"
                title="Cast by this pet or guardian, counted on its owner's row"
                data-testid="cast-via">· via {viaOf(row)}</span
              >{/if}{#if sameName.has(`${row.guid}|${row.spell_name}`)}
              <span
                class="text-muted ml-1 font-mono text-[11px]"
                title="Two spells share this name; this is spell id {row.spell_id}">#{row.spell_id}</span
              >{/if}</span
          >
```

The CSV gains the pet so a spreadsheet can tell the two apart:

```ts
  function csvLines(): string[][] {
    return [
      ['Player', 'Via', 'Spell', 'Spell id', 'Casts', 'Started', 'Cancelled', 'Failed', 'Casting s'],
      ...ordered.map((row) => [
        casterName(row),
        viaOf(row),
        row.spell_name,
        String(row.spell_id),
        String(row.succeeded),
        String(row.started),
        scaled ? '' : String(cancelled(row)),
        failedKnown(row) ? String(refused(row)) : '',
        (row.cast_time_ms / 1000).toFixed(1),
      ]),
    ];
  }
```

The "rhythm" line counts one caster's casts; a player and their pet are one caster's
rhythm, so it keys on the owner too:

```ts
    const casters = new Set(rows.map(ownerOf));
```

In `web/src/components/report/ReportView.svelte`, pass the names to the CastTable mount:

```svelte
            names={unitNames}
```

In `web/src/components/report/Glossary.svelte`, add:

```ts
    {
      term: 'via a pet',
      meaning:
        'A cast or an ability a player’s pet, totem or guardian did, counted on the player’s own row and named after the spell. Picking a player on the Source control shows their pets’ work with their own.',
    },
```

- [ ] **Step 4: Run the tests to verify they pass**

```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-casts.spec.ts tests/e2e/report-tabs.spec.ts --project=desktop && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: PASS, `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/components/report/CastTable.svelte src/components/report/ReportView.svelte src/components/report/Glossary.svelte tests/e2e/report-casts.spec.ts && npx prettier --write src/components/report/CastTable.svelte src/components/report/ReportView.svelte src/components/report/Glossary.svelte tests/e2e/report-casts.spec.ts
cd /Users/jh/code/forever && git add web/src/components/report web/tests/e2e/report-casts.spec.ts
git commit -m "feat(web): pet rows on Casts, under the player who owns them"
```

---

### Task 14: Aura refreshes in the full event stream

**Files:**
- Modify: `web/src/lib/report/exact.ts` (`eventStreamSql`, `StreamLine`, `loadEventStream`)
- Modify: `web/src/lib/report/events.ts` (`streamEvents` words a refresh)
- Modify: `web/src/fixtures/report/fixture.log` (one `SPELL_AURA_REFRESH` inside fight 3)
- Test: `web/src/lib/report/events.test.ts`, `web/src/lib/report/exact.test.ts`, `web/tests/e2e/report-tabs.spec.ts`

**Interfaces:**
- Consumes: the parquet's `kind = 'aura_refresh'` rows (`logs/engine/event/event.go` line 64)
  and its `spell_name`, `source_name`, `dest_name` columns.
- Produces, on `StreamLine` in `exact.ts`:

```ts
  kind: 'damage' | 'heal' | 'missed' | 'aura_refresh';
```

  and in `events.ts`, a refresh becomes a `SummaryEvent` of kind `'aura-applied'`, so the
  existing "Auras applied" toggle governs it.

**Decision taken here:** the refresh line is web-only, as the spec says — the summary's
aura tracks fold a refresh into the running segment and the engine keeps no refresh count,
so this reads the parquet the full stream already reads rather than adding a summary field.

- [ ] **Step 1: Write the failing tests**

In `web/src/lib/report/exact.test.ts`, append to the `eventStreamSql` describe (or create
one if the file has none):

```ts
describe('eventStreamSql', () => {
  it('reads aura refreshes alongside the hits, heals and misses', () => {
    const sql = eventStreamSql({ startMs: 0, endMs: 10_000 });
    expect(sql).toContain("kind IN ('heal', 'missed', 'aura_refresh')");
  });
});
```

In `web/src/lib/report/events.test.ts`, append:

```ts
describe('aura refreshes in the full stream', () => {
  it('words a refresh from the caster’s side and files it under auras applied', () => {
    const [line] = streamEvents([
      {
        atMs: 20_000,
        kind: 'aura_refresh',
        sourceGuid: 'Player-2',
        sourceName: 'Sunwick-Nightslayer',
        destGuid: 'Player-1',
        destName: 'Baelgrim-Nightslayer',
        spellName: 'Power Word: Fortitude',
        amount: 0,
        overheal: 0,
        absorbed: 0,
        blocked: 0,
        missType: '',
      },
    ]);
    expect(line.kind).toBe('aura-applied');
    expect(line.text).toBe('Sunwick refreshed Power Word: Fortitude on Baelgrim');
    expect(line.guid).toBe('Player-1');
    expect(line.guids).toEqual(['Player-2', 'Player-1']);
    expect(line.amount).toBeUndefined();
  });

  it('says the aura refreshed itself when the log names no caster', () => {
    const [line] = streamEvents([
      {
        atMs: 1000,
        kind: 'aura_refresh',
        sourceGuid: '',
        sourceName: '',
        destGuid: 'Player-1',
        destName: 'Baelgrim-Nightslayer',
        spellName: 'Rend',
        amount: 0,
        overheal: 0,
        absorbed: 0,
        blocked: 0,
        missType: '',
      },
    ]);
    expect(line.text).toBe('Rend refreshed on Baelgrim');
  });
});
```

In `web/tests/e2e/report-tabs.spec.ts`, append:

```ts
test('the full event stream lists aura refreshes under Auras applied', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto('/reports/fixture2abcd?fight=3&view=events');
  await page.getByTestId('events-stream').click();
  await expect(page.getByTestId('event-list')).toContainText('refreshed Power Word: Fortitude', {
    timeout: 60_000,
  });
  // The "Auras applied" toggle governs it: off, the refresh goes with the applications.
  await page.getByLabel('Auras applied').uncheck();
  await expect(page.getByTestId('event-list')).not.toContainText('refreshed Power Word: Fortitude');
});
```

`report-tabs.spec.ts` must import `serveDuckdbRuntime` from `./support/duckdb-runtime` if
it does not already.

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/events.test.ts src/lib/report/exact.test.ts`
Expected: FAIL — the SQL has no `aura_refresh`, and `streamEvents` types a refresh as damage.

- [ ] **Step 3: Give the fixture an aura refresh**

Add one line to `web/src/fixtures/report/fixture.log`, between the `20:12:18.000` energize
added in Task 8 and the `20:12:20.000` Frostbolt, refreshing the Fortitude the priest put
on the tank at `20:12:04`:

```
9/26 20:12:19.000  SPELL_AURA_REFRESH,Player-4184-000000A2,"Sunwick-Nightslayer",0x512,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,1243,"Power Word: Fortitude",0x2,BUFF
```

The aura is already open, so `addAuras` neither counts an application nor closes a
segment: `fights/3/summary.json` comes out byte-identical and only `events.parquet` gains
a row. Confirm that after the regeneration in Step 5 — `git status` must show a change to
`events.parquet` and none to `summary.json`.

- [ ] **Step 4: Implement**

In `web/src/lib/report/exact.ts`, widen the stream:

```ts
export interface StreamLine {
  atMs: number;
  kind: 'damage' | 'heal' | 'missed' | 'aura_refresh';
  // sourceGuid, sourceName, destGuid, destName, spellName, amount, overheal, absorbed,
  // blocked and missType are unchanged.
}
```

```ts
/**
 * A refresh rides with the hits and heals rather than with the summary's aura segments:
 * the summary folds a refresh into the running segment and keeps no count of them, so
 * "when did this actually get re-applied" is only answerable from the fight's own lines.
 */
export function eventStreamSql(window: TimeWindow): string {
  return `SELECT ${FIGHT_MS} AS fight_ms, kind, source_guid, source_name, dest_guid, dest_name, spell_name,
  coalesce(amount, 0) AS amount, coalesce(overheal, 0) AS overheal, coalesce(absorbed, 0) AS absorbed,
  coalesce(blocked, 0) AS blocked, coalesce(miss_type, '') AS miss_type
FROM (SELECT * FROM ${DAMAGE_LINES} UNION ALL SELECT * FROM ${EVENTS_TABLE} WHERE kind IN ('heal', 'missed', 'aura_refresh'))
WHERE ${windowClause(window)}
ORDER BY time_unix_nano, line`;
}
```

and the mapper:

```ts
    kind:
      row.kind === 'heal'
        ? 'heal'
        : row.kind === 'missed'
          ? 'missed'
          : row.kind === 'aura_refresh'
            ? 'aura_refresh'
            : 'damage',
```

In `web/src/lib/report/events.ts`, handle the refresh at the top of `streamEvents`'s map,
before the miss branch:

```ts
    if (line.kind === 'aura_refresh') {
      // Filed under "Auras applied": a refresh is the aura going back up, and a reader
      // who switched applications off does not want refreshes either. The summary's own
      // aura lines fold a refresh into the running segment and never say it happened,
      // which is exactly the gap this fills.
      return {
        atMs: line.atMs,
        kind: 'aura-applied',
        guid: line.destGuid,
        guids: [line.sourceGuid, line.destGuid],
        text: who === 'Something' ? `${spell} refreshed on ${whom}` : `${who} refreshed ${spell} on ${whom}`,
        tags: ['refresh', 'refreshed'],
      };
    }
```

`who`, `whom` and `spell` are already computed above the miss branch in the same callback;
the guard reads `who === 'Something'` because that is the fallback the function already
uses for a line with no source name.

- [ ] **Step 5: Regenerate the fixture and run the tests**

```bash
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && git status --short src/fixtures/report && npx vitest run && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-tabs.spec.ts --project=desktop && npx playwright test tests/e2e/report-phone.spec.ts --project=mobile && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: `git status` shows `fights/3/events.parquet` changed and `fights/3/summary.json`
unchanged; every test passes; `- 0 errors`.

- [ ] **Step 6: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/lib/report/exact.ts src/lib/report/events.ts src/lib/report/events.test.ts src/lib/report/exact.test.ts tests/e2e/report-tabs.spec.ts && npx prettier --write src/lib/report/exact.ts src/lib/report/events.ts src/lib/report/events.test.ts src/lib/report/exact.test.ts tests/e2e/report-tabs.spec.ts
cd /Users/jh/code/forever && git add web/src/lib/report web/tests/e2e/report-tabs.spec.ts web/src/fixtures/report
git commit -m "feat(web): aura refreshes in the full event stream"
```

- [ ] **Step 7: Group 4 review gate**

```bash
cd /Users/jh/code/forever/logs && go test ./... && (cd ../api && go test ./...)
cd /Users/jh/code/forever/web && npx vitest run && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?" && npx playwright test tests/e2e/report-casts.spec.ts tests/e2e/report-tabs.spec.ts tests/e2e/report-tables.spec.ts --project=desktop
```
Merge `pet-casts-and-refreshes` into `main`.

---

## Group 5 — Per-ability graphs on the main chart

Tasks 15 and 16. No engine change. Review gate at the end of Task 16.

---

### Task 15: One ability's per-second line, measured from the fight's events

**Files:**
- Modify: `web/src/lib/report/exact.ts` (`abilitySeriesSql`, `measureAbilitySeries`)
- Test: `web/src/lib/report/exact.test.ts`

**Interfaces:**
- Consumes: `rowsSql`, `MeasureOptions`, `ALL_ROWS`, `rowsOf`, `num`, `quote` (all already
  in `exact.ts`); `BUCKET_MS` from `window.ts`; `QueryLayer` from `query.ts`.
- Produces:

```ts
/** One actor's one ability, per whole second of the fight, for the main chart. */
export function abilitySeriesSql(
  kind: ActorKind, guid: string, spellId: number, via: string, window: TimeWindow,
): string;
/**
 * Measures one ability's effective amount per second over the whole fight, so the line
 * can ride behind the main series and be brushed with it.
 */
export async function measureAbilitySeries(
  layer: QueryLayer, eventsUrl: string, kind: ActorKind,
  guid: string, spellId: number, via: string, durationMs: number,
): Promise<number[]>;
```

**Decision taken here:** the measure is over the whole fight, not the window, so brushing
does not re-measure — the chart already slices what it draws, and re-reading the parquet on
every drag would put a DuckDB query behind a pointer move.

- [ ] **Step 1: Write the failing tests**

Append to `web/src/lib/report/exact.test.ts`:

```ts
describe('abilitySeriesSql', () => {
  it('buckets one actor’s one ability by whole second', () => {
    const sql = abilitySeriesSql('damage-done', 'Player-1', 1464, '', { startMs: 0, endMs: 60_000 });
    expect(sql).toContain('AND spell_id = 1464');
    expect(sql).toContain("WHERE actor = 'Player-1' AND via = ''");
    expect(sql).toContain('floor(fight_ms / 1000) AS second');
    expect(sql).toContain('sum(effective) AS amount');
    expect(sql).toContain('GROUP BY 1');
  });

  it('keeps a pet’s ability apart from its owner’s own of the same spell', () => {
    expect(abilitySeriesSql('healing', 'Player-2', 115175, 'Jade Serpent Statue', {
      startMs: 0,
      endMs: 1000,
    })).toContain("AND via = 'Jade Serpent Statue'");
  });

  it('reads damage taken from the victim’s side', () => {
    expect(abilitySeriesSql('damage-taken', 'Player-4', 334660, '', { startMs: 0, endMs: 1000 })).toContain(
      'SELECT dest_guid AS actor',
    );
  });
});

describe('measureAbilitySeries', () => {
  it('fills a bucket per second of the fight, zero where the ability did nothing', async () => {
    const layer = {
      run: async () => ({
        columns: ['second', 'amount'],
        rows: [
          [0n, 2100n],
          [3n, 2300n],
        ] as unknown[][],
      }),
    } as unknown as QueryLayer;
    const series = await measureAbilitySeries(layer, 'x.parquet', 'damage-done', 'P1', 1464, '', 5000);
    expect(series).toEqual([2100, 0, 0, 2300, 0]);
  });

  it('never returns an empty line for a fight with a length', async () => {
    const layer = { run: async () => ({ columns: [], rows: [] }) } as unknown as QueryLayer;
    expect(await measureAbilitySeries(layer, 'x.parquet', 'healing', 'P1', 1, '', 2000)).toEqual([0, 0]);
  });
});
```

Add `abilitySeriesSql` and `measureAbilitySeries` to the file's import from `./exact`.

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/exact.test.ts`
Expected: FAIL — `abilitySeriesSql is not a function`.

- [ ] **Step 3: Implement**

In `web/src/lib/report/exact.ts`, change the `window.ts` import to bring in the bucket
width and add the two functions at the end of the file:

```ts
import { BUCKET_MS, type TimeWindow } from './window';
```

```ts
/**
 * One actor's one ability, bucketed by whole second: the line the main chart draws behind
 * its own series when a reader puts an ability "on the chart". It counts what the tables
 * count -- rowsSql is the same projection the totals and the splits are read from -- so
 * the line's peak and the row's Max are the same number.
 */
export function abilitySeriesSql(
  kind: ActorKind,
  guid: string,
  spellId: number,
  via: string,
  window: TimeWindow,
  options: MeasureOptions = {},
): string {
  return `WITH rows AS (${rowsSql(kind, window, { ...options, ability: spellId })})
SELECT floor(fight_ms / ${BUCKET_MS}) AS second, sum(effective) AS amount
FROM rows
WHERE actor = ${quote(guid)} AND via = ${quote(via)}
GROUP BY 1
ORDER BY 1`;
}

/**
 * Measures one ability's effective amount per second over the whole fight. The whole
 * fight, not the window: the chart slices what it draws, so a brush costs nothing, where
 * re-measuring would put a parquet read behind every pointer move.
 */
export async function measureAbilitySeries(
  layer: QueryLayer,
  eventsUrl: string,
  kind: ActorKind,
  guid: string,
  spellId: number,
  via: string,
  durationMs: number,
  options: MeasureOptions = {},
): Promise<number[]> {
  const endMs = Math.max(durationMs, BUCKET_MS);
  const result = await layer.run(
    eventsUrl,
    abilitySeriesSql(kind, guid, spellId, via, { startMs: 0, endMs }, options),
    ALL_ROWS,
  );
  const series = new Array<number>(Math.ceil(endMs / BUCKET_MS)).fill(0);
  for (const row of rowsOf(result)) {
    const second = num(row.second);
    if (second >= 0 && second < series.length) series[second] = num(row.amount);
  }
  return series;
}
```

`options` is threaded rather than dropped because the call site (Task 16) passes the
page's `measureOptions`, whose `pets` map is what folds a pet's lines onto its owner: the
actor asked for is the owner's GUID as the summary files it, and `via` -- exactly what
`Ability.via` carries -- tells the pet's lines from the owner's own of the same spell.
`ability` is overwritten rather than merged, because this measure is always one spell's.
The tests above call both without options, which is the default.

- [ ] **Step 4: Run the tests to verify they pass**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/exact.test.ts && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"`
Expected: PASS, `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/lib/report/exact.ts src/lib/report/exact.test.ts && npx prettier --write src/lib/report/exact.ts src/lib/report/exact.test.ts
cd /Users/jh/code/forever && git add web/src/lib/report
git commit -m "feat(web): measure one ability's per-second line from the fight's events"
```

---

### Task 16: "On the chart" puts an ability behind the main series

**Files:**
- Modify: `web/src/components/report/ActorRow.svelte` (the control in the abilities table)
- Modify: `web/src/components/report/ActorTable.svelte` (thread the two props)
- Modify: `web/src/components/report/ReportView.svelte` (hold the picked line, measure it, draw it)
- Modify: `web/src/components/report/Glossary.svelte` (one term)
- Test: `web/tests/e2e/report-tables.spec.ts`, `web/tests/e2e/report-phone.spec.ts`

**Interfaces:**
- Consumes: `measureAbilitySeries` (Task 15), `sharedQueryLayer`, `eventsUrl`,
  `schoolToken`, `abilityKey`, `splitUnitName`.
- Produces, on `ActorRow` and `ActorTable`:

```ts
  /** Puts this ability on the main chart, or takes it off when it is already there. */
  onChart?: (actor: Actor, ability: Ability) => void;
  /** `${guid}|${abilityKey(ability)}` of the ability currently on the chart; '' for none. */
  charted?: string;
```

- Produces, in `ReportView.svelte`:

```ts
  /** The one ability drawn behind the main chart's series; null when none is. */
  let chartedAbility = $state<{
    key: string;
    label: string;
    token: string;
    series: number[];
  } | null>(null);
```

**Decisions taken here:** one ability at a time, so picking another replaces it (the spec
says so). The line is not in the url — it is a look, not a view. It is offered only where
the chart is: one pull, Analyze, the tables view, and one of the three actor tabs; the
night has no chart of this kind and the control is absent there. The colour is the
ability's school colour, through `schoolToken`, as the spec asks.

- [ ] **Step 1: Write the failing e2e tests**

Append to `web/tests/e2e/report-tables.spec.ts`:

```ts
test('an ability goes on the main chart in its school colour and comes off again', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto(`${FIGHT}&tab=damage-done`);
  await page.getByTestId('actor-Player-4184-000000A1').getByRole('button').first().click();
  const control = page
    .getByTestId('row-abilities')
    .getByRole('row', { name: /Slam/ })
    .getByTestId('ability-chart');
  await expect(control).toHaveText('On the chart');
  await control.click();
  // The legend names the actor and the ability once the measure lands.
  await expect(page.getByTestId('time-chart')).toContainText('Baelgrim · Slam', { timeout: 60_000 });
  await expect(control).toHaveText('Off the chart');
  await control.click();
  await expect(page.getByTestId('time-chart')).not.toContainText('Baelgrim · Slam');
});

test('picking a second ability replaces the first: one line at a time', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto(`${FIGHT}&tab=healing`);
  await page.getByTestId('actor-Player-4184-000000A2').getByRole('button').first().click();
  const controls = page.getByTestId('row-abilities').getByTestId('ability-chart');
  await controls.first().click();
  await expect(page.getByTestId('time-chart')).toContainText('Sunwick · ', { timeout: 60_000 });
  await controls.nth(1).click();
  await expect(controls.first()).toHaveText('On the chart');
  await expect(controls.nth(1)).toHaveText('Off the chart');
});

test('the night offers no ability line, because it has no chart of this kind', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&tab=damage-done');
  await page.getByTestId('actor-Player-4184-000000A1').getByRole('button').first().click();
  await expect(page.getByTestId('ability-chart')).toHaveCount(0);
});
```

Append to `web/tests/e2e/report-phone.spec.ts`:

```ts
test('the on-the-chart control is a 44px target on a phone', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  await page.getByTestId('actor-Player-4184-000000A1').getByRole('button').first().click();
  const control = page.getByTestId('ability-chart').first();
  await expect(control).toBeVisible();
  const box = await control.boundingBox();
  expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run:
```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-tables.spec.ts --project=desktop
```
Expected: FAIL — `ability-chart` not found.

- [ ] **Step 3: Implement the control**

In `web/src/components/report/ActorRow.svelte`, add the props:

```ts
    onChart = undefined,
    charted = '',
```

```ts
    /** Puts this ability on the main chart, or takes it off when it is already there. */
    onChart?: (actor: Actor, ability: Ability) => void;
    /** `${guid}|${abilityKey(ability)}` of the ability currently on the chart; '' for none. */
    charted?: string;
```

and a helper beside `landed`:

```ts
  /** The key the page identifies a charted line by: this row's actor and this ability. */
  const chartKey = (ability: Ability): string => `${actor.guid}|${abilityKey(ability)}`;
```

In the abilities table's `<thead>`, after the "Not landed" header:

```svelte
              {#if onChart !== undefined}
                <th scope="col" class="py-1 pl-3 text-right font-normal"
                  ><span class="sr-only">On the chart</span></th
                >
              {/if}
```

and in each `<tr>`, after the notes cell:

```svelte
                {#if onChart !== undefined}
                  <td class="py-1.5 pl-3 text-right whitespace-nowrap">
                    <button
                      type="button"
                      class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-9"
                      title="Draw this ability's amount per second behind the chart above"
                      aria-pressed={charted === chartKey(ability)}
                      data-testid="ability-chart"
                      onclick={() => onChart?.(actor, ability)}
                      >{charted === chartKey(ability) ? 'Off the chart' : 'On the chart'}</button
                    >
                  </td>
                {/if}
```

In `web/src/components/report/ActorTable.svelte`, add the same two props to the props
block and its type, and pass them down in the `<ActorRow …>` mount:

```svelte
          {onChart}
          {charted}
```

- [ ] **Step 4: Implement the line on the chart**

In `web/src/components/report/ReportView.svelte`, add the import:

```ts
    measureAbilitySeries,
```
to the existing `from '../../lib/report/exact'` import list, and `schoolToken` to the
`format` import, and `abilityKey, type Ability` to the `types` import.

Add the state and the toggle, beside the other measures:

```ts
  /**
   * The one ability drawn behind the main chart's series: a look, not a view, so it is
   * not in the url. Measured over the whole fight once and sliced by the brush like every
   * other line, and dropped whenever the fight, the tab or the source scope changes,
   * because it answered a question about the table that was on screen then.
   */
  let chartedAbility = $state<{ key: string; label: string; token: string; series: number[] } | null>(
    null,
  );
  let chartedError = $state('');
  /** The measure that is wanted now; an answer for an older pick is dropped. */
  let chartedToken = 0;
  $effect(() => {
    void [state.fight, state.tab, state.source, nightMode];
    chartedAbility = null;
    chartedError = '';
    chartedToken += 1;
  });
  async function toggleAbilityOnChart(actor: Actor, ability: Ability): Promise<void> {
    const key = `${actor.guid}|${abilityKey(ability)}`;
    if (chartedAbility?.key === key) {
      chartedAbility = null;
      return;
    }
    const token = ++chartedToken;
    chartedError = '';
    try {
      const series = await measureAbilitySeries(
        sharedQueryLayer(),
        eventsUrl(dataBase, state.fight, engineVersion),
        tableKind,
        actor.guid,
        ability.spell_id,
        ability.via ?? '',
        base?.duration_ms ?? 0,
        measureOptions,
      );
      if (token !== chartedToken) return;
      chartedAbility = {
        key,
        label: `${splitUnitName(actor.name).name} · ${ability.name}`,
        token: schoolToken(ability.school),
        series,
      };
    } catch (thrown) {
      if (token !== chartedToken) return;
      chartedError = `That ability's line did not load${thrown instanceof Error ? ` (${thrown.message})` : ''}.`;
    }
  }
  /** True where there is a chart to put an ability on: one pull, Analyze, an actor tab. */
  const abilityChartAvailable = $derived(
    !nightMode &&
      state.mode === 'analyze' &&
      state.view === 'tables' &&
      (state.tab === 'damage-done' || state.tab === 'damage-taken' || state.tab === 'healing'),
  );
```

Append the line to `chartExtra`:

```ts
  const chartExtra = $derived([
    ...(scoped === null || state.tab !== 'summary'
      ? []
      : [
          {
            label: 'Damage taken',
            series: combinedSeries(
              scoped.damage_taken.filter((actor) =>
                inSource(actor.guid, state.source, playerSet, friendlySet),
              ),
            ),
            token: 'var(--color-death)',
          },
          {
            label: 'Healing',
            series: combinedSeries(
              scoped.healing.filter((actor) => inSource(actor.guid, state.source, playerSet, friendlySet)),
            ),
            token: 'var(--color-kill)',
          },
        ]),
    // The line respects the brush the way the main series does: TimeChart is handed the
    // whole fight's buckets and draws the window over them.
    ...(chartedAbility === null
      ? []
      : [{ label: chartedAbility.label, series: chartedAbility.series, token: chartedAbility.token }]),
  ]);
```

Pass the two props to the `ActorTable` mount:

```svelte
            onChart={abilityChartAvailable ? toggleAbilityOnChart : undefined}
            charted={chartedAbility?.key ?? ''}
```

and put the failure where the other measure failures go, just under the approximate note:

```svelte
          {#if chartedError !== ''}
            <p class="text-wipe text-[12px]" role="alert" data-testid="ability-chart-error">
              {chartedError}
            </p>
          {/if}
```

In `web/src/components/report/Glossary.svelte`, add:

```ts
    {
      term: 'On the chart',
      meaning:
        'On an opened row’s abilities table: draws that one ability’s amount per second behind the chart above, in its spell school’s colour, measured from the fight’s own events. One ability at a time; picking another replaces it. It is a look rather than a view, so a copied link does not carry it.',
    },
```

- [ ] **Step 5: Run the tests to verify they pass**

```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-tables.spec.ts --project=desktop && npx playwright test tests/e2e/report-phone.spec.ts --project=mobile && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: PASS, `- 0 errors`.

- [ ] **Step 6: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/components/report/ActorRow.svelte src/components/report/ActorTable.svelte src/components/report/ReportView.svelte src/components/report/Glossary.svelte tests/e2e/report-tables.spec.ts tests/e2e/report-phone.spec.ts && npx prettier --write src/components/report/ActorRow.svelte src/components/report/ActorTable.svelte src/components/report/ReportView.svelte src/components/report/Glossary.svelte tests/e2e/report-tables.spec.ts tests/e2e/report-phone.spec.ts
cd /Users/jh/code/forever && git add web/src/components/report web/tests/e2e
git commit -m "feat(web): put one ability's measured line on the main chart"
```

- [ ] **Step 7: Group 5 review gate**

```bash
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && git status --short src/fixtures/report
npx vitest run && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?" && npx playwright test tests/e2e/report-tables.spec.ts tests/e2e/report-brush.spec.ts --project=desktop
```
The fixture regeneration must produce no diff: this group changed no engine code. Merge
`ability-graphs` into `main`.

---

## Group 6 — Phases, from curated per-boss tables

Tasks 17 to 22. Review gate at the end of Task 22.

---

### Task 17: The mechanics tables carry phases

**Files:**
- Modify: `logs/engine/mechanics/mechanics.go` (`Phase`, `PhaseStart`, the trigger kinds, `Parse`)
- Modify: `logs/engine/mechanics/tables/9001.json` (one phase, for the fixture)
- Test: `logs/engine/mechanics/mechanics_test.go`

**Interfaces:**
- Produces, in `package mechanics`:

```go
const (
	OnCastStart   = "cast_start"
	OnCastSuccess = "cast_success"
	OnAuraApplied = "aura_applied"
	OnAuraRemoved = "aura_removed"
)

// PhaseStart is what begins a phase: an enemy's cast or aura change, or the
// boss's own health passing a percentage. Exactly one of the two forms is set.
type PhaseStart struct {
	SpellID   int64   `json:"spell_id,omitempty"`
	On        string  `json:"on,omitempty"`
	HealthPct float64 `json:"health_pct,omitempty"`
}

// Phase is one named stretch of an encounter, and what begins it.
type Phase struct {
	Name   string     `json:"name"`
	Starts PhaseStart `json:"starts"`
}
```

  and `Phases []Phase \`json:"phases,omitempty"\`` on `Table`.

**Decision taken here:** the spec lists three trigger kinds; the engine accepts a fourth,
`cast_start`, because a boss's phase-opening channel begins the phase when the cast begins,
not when it lands — and a channel that is interrupted never lands at all. The fixture's
Anima Surge is exactly that case: `SPELL_CAST_START` at 14.0 s of fight 3 and no success
line, ever. Four kinds, one map, validated the same way.

- [ ] **Step 1: Write the failing tests**

Append to `logs/engine/mechanics/mechanics_test.go`:

```go
func TestParseAcceptsPhasesAndRefusesAMalformedTrigger(t *testing.T) {
	good := []byte(`{"encounter_id": 9001, "name": "Warden Kelthas", "mechanics": [],
		"phases": [
			{"name": "Phase 2", "starts": {"spell_id": 334653, "on": "cast_start"}},
			{"name": "Phase 3", "starts": {"health_pct": 30}}]}`)
	table, err := Parse(good)
	if err != nil {
		t.Fatal(err)
	}
	if len(table.Phases) != 2 {
		t.Fatalf("phases = %+v", table.Phases)
	}
	if table.Phases[0].Starts.On != OnCastStart || table.Phases[1].Starts.HealthPct != 30 {
		t.Fatalf("phases = %+v", table.Phases)
	}

	refused := map[string][]byte{
		"an unnamed phase":      []byte(`{"encounter_id":1,"name":"X","phases":[{"starts":{"health_pct":50}}]}`),
		"an unknown on":         []byte(`{"encounter_id":1,"name":"X","phases":[{"name":"P2","starts":{"spell_id":5,"on":"cast_finished"}}]}`),
		"a spell with no on":    []byte(`{"encounter_id":1,"name":"X","phases":[{"name":"P2","starts":{"spell_id":5}}]}`),
		"both forms at once":    []byte(`{"encounter_id":1,"name":"X","phases":[{"name":"P2","starts":{"spell_id":5,"on":"cast_start","health_pct":50}}]}`),
		"an on with no spell":   []byte(`{"encounter_id":1,"name":"X","phases":[{"name":"P2","starts":{"on":"cast_start"}}]}`),
		"a health out of range": []byte(`{"encounter_id":1,"name":"X","phases":[{"name":"P2","starts":{"health_pct":140}}]}`),
		"no trigger at all":     []byte(`{"encounter_id":1,"name":"X","phases":[{"name":"P2","starts":{}}]}`),
	}
	for why, data := range refused {
		if _, err := Parse(data); err == nil {
			t.Errorf("%s must be refused", why)
		}
	}
}

func TestTheFixtureEncounterNamesItsSecondPhase(t *testing.T) {
	table, ok := Load(9001)
	if !ok {
		t.Fatal("the fixture's encounter must have a table")
	}
	if len(table.Phases) != 1 {
		t.Fatalf("phases = %+v", table.Phases)
	}
	if table.Phases[0].Name != "Phase 2" || table.Phases[0].Starts.SpellID != 334653 {
		t.Fatalf("phase = %+v", table.Phases[0])
	}
}
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/logs && go test ./engine/mechanics/`
Expected: FAIL to compile — `table.Phases undefined`.

- [ ] **Step 3: Implement**

In `logs/engine/mechanics/mechanics.go`, after the `Mechanic` type:

```go
// The trigger kinds a phase can start on. cast_start is here alongside the
// spec's three because a boss's phase-opening channel begins the phase when
// the cast begins -- and a channel a raid interrupts never lands at all, so
// keying its phase on the success would be keying it on the raid failing.
const (
	OnCastStart   = "cast_start"
	OnCastSuccess = "cast_success"
	OnAuraApplied = "aura_applied"
	OnAuraRemoved = "aura_removed"
)

var phaseOn = map[string]bool{
	OnCastStart: true, OnCastSuccess: true, OnAuraApplied: true, OnAuraRemoved: true,
}

// PhaseStart is what begins a phase: an enemy's cast or aura change, or the
// boss's own health passing a percentage. Exactly one of the two forms is set.
type PhaseStart struct {
	// SpellID with On: the spell whose cast or aura change starts the phase,
	// on any enemy.
	SpellID int64  `json:"spell_id,omitempty"`
	On      string `json:"on,omitempty"`
	// HealthPct starts the phase the first time the boss's own health is at or
	// below this percentage, read from the advanced block on its lines.
	HealthPct float64 `json:"health_pct,omitempty"`
}

// Phase is one named stretch of an encounter, and what begins it. Phase 1
// starts at the pull and needs no entry.
type Phase struct {
	Name   string     `json:"name"`
	Starts PhaseStart `json:"starts"`
}
```

Add the field to `Table`:

```go
// Table is one encounter's mechanics, and the phases it is fought in.
type Table struct {
	EncounterID int64      `json:"encounter_id"`
	Name        string     `json:"name"`
	Mechanics   []Mechanic `json:"mechanics"`
	// Phases are the stretches the encounter is fought in, in the order they
	// happen. Empty for an encounter nobody has curated phases for.
	Phases []Phase `json:"phases,omitempty"`
}
```

and validate them at the end of `Parse`, before the `return t, nil`:

```go
	for i, p := range t.Phases {
		if p.Name == "" {
			return Table{}, fmt.Errorf("phases[%d]: name is required", i)
		}
		switch {
		case p.Starts.SpellID > 0:
			if !phaseOn[p.Starts.On] {
				return Table{}, fmt.Errorf(
					"phases[%d]: on %q is not cast_start, cast_success, aura_applied or aura_removed", i, p.Starts.On)
			}
			if p.Starts.HealthPct != 0 {
				return Table{}, fmt.Errorf(
					"phases[%d]: a phase starts on a spell or on a health percentage, not both", i)
			}
		case p.Starts.HealthPct > 0 && p.Starts.HealthPct <= 100:
			if p.Starts.On != "" {
				return Table{}, fmt.Errorf("phases[%d]: on belongs to a spell trigger; a health trigger takes none", i)
			}
		default:
			return Table{}, fmt.Errorf(
				"phases[%d]: starts must name a spell_id with on, or a health_pct above 0 and at most 100", i)
		}
	}
```

In `logs/engine/mechanics/tables/9001.json`, add the phases after the mechanics array:

```json
  ],
  "phases": [
    {
      "name": "Phase 2",
      "starts": { "spell_id": 334653, "on": "cast_start" }
    }
  ]
}
```

(the file's `mechanics` array is unchanged; only the closing brace gains a sibling key.)

- [ ] **Step 4: Run the tests to verify they pass**

Run: `cd /Users/jh/code/forever/logs && go test ./engine/mechanics/`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add logs/engine/mechanics
git commit -m "feat(logs): the mechanics tables carry an encounter's phases"
```

---

### Task 18: The summary says which phase each stretch of a fight was

**Files:**
- Create: `logs/engine/summary/phases.go`
- Create: `logs/engine/summary/phases_test.go`
- Modify: `logs/engine/summary/summary.go` (`Summary.Phases`, `Accumulator.phaseAt`, `New`, `addNow`, `Snapshot`)
- Regenerate: `logs/engine/summary/testdata/v16.summary.json.golden`, `web/src/fixtures/report/`

**Interfaces:**
- Consumes: `Options.Mechanics *mechanics.Table` (already set per fight in
  `logs/engine/session/session.go`'s `startFight`), `mechanics.Phase`, `mechanics.PhaseStart`,
  the trigger constants, `Accumulator.isPlayer`, `Accumulator.ms`, `Accumulator.name`.
- Produces:

```go
// Phase is one named stretch of the fight, from its trigger to the next
// trigger or the fight's end.
type Phase struct {
	Name    string `json:"name"`
	StartMS int64  `json:"start_ms"`
	EndMS   int64  `json:"end_ms"`
}
```

  and `Phases []Phase \`json:"phases"\`` on `Summary`, always an array and never null.

**Decisions taken here:**
- A trigger fires once: a boss that casts its opener twice is still in the phase its first
  cast began.
- A fight whose table has phases but whose triggers never fired emits an empty list, not a
  single "Phase 1" spanning the pull — a wipe in the opening is a fight with no phases to
  read, and "Phase 1" against the whole pull is a band with no information in it.
- The stretch before the first trigger is called "Phase 1" for every encounter: the table
  names the phases its triggers open, and the pull's own opening has no trigger.
- A health trigger is read off the advanced block of any line describing the unit the
  table's `name` names, which is the same identity `fight.bossHealthPct` uses.

- [ ] **Step 1: Write the failing tests**

```go
// logs/engine/summary/phases_test.go
package summary

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

// phaseFight is the fight the phase tests snapshot: sixty seconds of Warden Kelthas.
func phaseFight() fight.Fight {
	return fight.Fight{Index: 1, Kind: fight.Encounter, EncounterID: 9001, Name: "Warden Kelthas",
		Start: at(0), End: at(60), Players: []string{tank, mage}}
}

func withPhases(t *testing.T, phases []mechanics.Phase) (Options, *units.Registry) {
	t.Helper()
	o, reg := opts(t)
	o.Mechanics = &mechanics.Table{EncounterID: 9001, Name: "Warden Kelthas", Phases: phases}
	return o, reg
}

// The cast the table names opens the phase, and the phase before it runs up to that
// instant: the bands on the chart and the presets in the strip are these spans.
func TestPhasesRunFromEachTriggerToTheNext(t *testing.T) {
	o, reg := withPhases(t, []mechanics.Phase{
		{Name: "Phase 2", Starts: mechanics.PhaseStart{SpellID: 334653, On: mechanics.OnCastStart}},
	})
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		dmg(1, mage, boss, 116, "Frostbolt", 100, -1),
		{Time: at(14), Kind: event.CastStart, Name: "SPELL_CAST_START",
			Source: event.Unit{GUID: boss, Flags: 0xa48}, Spell: event.Spell{ID: 334653, Name: "Anima Surge"}},
		// A second cast of the same spell is not a second phase.
		{Time: at(30), Kind: event.CastStart, Name: "SPELL_CAST_START",
			Source: event.Unit{GUID: boss, Flags: 0xa48}, Spell: event.Spell{ID: 334653, Name: "Anima Surge"}},
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(phaseFight(), "test")
	want := []Phase{{Name: "Phase 1", StartMS: 0, EndMS: 14000}, {Name: "Phase 2", StartMS: 14000, EndMS: 60000}}
	if len(s.Phases) != len(want) {
		t.Fatalf("phases = %+v, want %+v", s.Phases, want)
	}
	for i, p := range want {
		if s.Phases[i] != p {
			t.Errorf("phase %d = %+v, want %+v", i, s.Phases[i], p)
		}
	}
}

// Each trigger kind, one fight each, so a curator can trust all four.
func TestEveryTriggerKindOpensItsPhase(t *testing.T) {
	cases := []struct {
		name  string
		start mechanics.PhaseStart
		event event.Event
	}{
		{"cast success", mechanics.PhaseStart{SpellID: 334653, On: mechanics.OnCastSuccess},
			cast(20, boss, tank, 334653, "Anima Surge")},
		{"aura applied", mechanics.PhaseStart{SpellID: 321038, On: mechanics.OnAuraApplied},
			event.Event{Time: at(20), Kind: event.AuraApplied, Name: "SPELL_AURA_APPLIED",
				Source: event.Unit{GUID: boss, Flags: 0xa48}, Dest: event.Unit{GUID: mage, Flags: 0x512},
				Spell: event.Spell{ID: 321038, Name: "Wrack Soul"}, AuraType: "DEBUFF"}},
		{"aura removed", mechanics.PhaseStart{SpellID: 321038, On: mechanics.OnAuraRemoved},
			event.Event{Time: at(20), Kind: event.AuraRemoved, Name: "SPELL_AURA_REMOVED",
				Source: event.Unit{GUID: boss, Flags: 0xa48}, Dest: event.Unit{GUID: mage, Flags: 0x512},
				Spell: event.Spell{ID: 321038, Name: "Wrack Soul"}, AuraType: "DEBUFF"}},
		{"boss health", mechanics.PhaseStart{HealthPct: 50},
			event.Event{Time: at(20), Kind: event.Damage, Name: "SPELL_DAMAGE",
				Source: event.Unit{GUID: mage, Flags: 0x512},
				Dest:   event.Unit{GUID: boss, Name: "Warden Kelthas", Flags: 0xa48},
				Spell:  event.Spell{ID: 116, Name: "Frostbolt", School: 0x10},
				Amount: event.OptInt{V: 10, OK: true}, Overkill: event.OptInt{V: -1, OK: true},
				Adv: event.Advanced{OK: true, InfoGUID: boss, CurrentHP: 400000, MaxHP: 1000000}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			o, reg := withPhases(t, []mechanics.Phase{{Name: "Phase 2", Starts: c.start}})
			a := New(o)
			a.Start(at(0))
			// The boss must be named before the health trigger can recognise it.
			naming := dmg(1, mage, boss, 116, "Frostbolt", 1, -1)
			naming.Dest.Name = "Warden Kelthas"
			for _, e := range []event.Event{naming, c.event} {
				reg.Observe(e)
				a.Add(e)
			}
			s := a.Snapshot(phaseFight(), "test")
			if len(s.Phases) != 2 || s.Phases[1].Name != "Phase 2" || s.Phases[1].StartMS != 20000 {
				t.Fatalf("phases = %+v", s.Phases)
			}
		})
	}
}

// A pull that never reached the trigger has no phases to read, and says so with an empty
// list rather than one band called Phase 1 across the whole pull.
func TestAFightThatNeverLeftTheOpeningHasNoPhases(t *testing.T) {
	o, reg := withPhases(t, []mechanics.Phase{
		{Name: "Phase 2", Starts: mechanics.PhaseStart{SpellID: 334653, On: mechanics.OnCastStart}},
	})
	a := New(o)
	a.Start(at(0))
	e := dmg(1, mage, boss, 116, "Frostbolt", 100, -1)
	reg.Observe(e)
	a.Add(e)
	s := a.Snapshot(phaseFight(), "test")
	if len(s.Phases) != 0 {
		t.Fatalf("phases = %+v, want none", s.Phases)
	}
}

// An encounter with no table at all, which is most of them, emits an empty list.
func TestAnEncounterWithNoTableHasAnEmptyPhaseList(t *testing.T) {
	_, _, s := build(t)
	if s.Phases == nil {
		t.Fatal("phases must be an empty array, never null")
	}
	if len(s.Phases) != 0 {
		t.Fatalf("phases = %+v, want none", s.Phases)
	}
}
```

The file needs `"github.com/jhunthrop/foreversixty/logs/engine/units"` for the registry in
`withPhases`'s return type.

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/logs && go test ./engine/summary/ -run 'Phase'`
Expected: FAIL to compile — `s.Phases undefined`.

- [ ] **Step 3: Implement**

```go
// logs/engine/summary/phases.go
package summary

import (
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

// Phase is one named stretch of the fight, from its trigger to the next
// trigger or the fight's end.
type Phase struct {
	Name    string `json:"name"`
	StartMS int64  `json:"start_ms"`
	EndMS   int64  `json:"end_ms"`
}

// firstPhaseName is what the stretch before the first trigger is called. The
// table names the phases its triggers open; the pull's own opening has no
// trigger and no entry, and every encounter calls it the same thing.
const firstPhaseName = "Phase 1"

// notePhase records the first instant each of the encounter's phase triggers
// fired. A trigger fires once: a boss that casts its opener twice is still in
// the phase its first cast began, and a health threshold crossed again after a
// heal did not start the phase over.
func (a *Accumulator) notePhase(e event.Event) {
	if a.opt.Mechanics == nil || len(a.opt.Mechanics.Phases) == 0 {
		return
	}
	for i, p := range a.opt.Mechanics.Phases {
		if _, fired := a.phaseAt[i]; fired {
			continue
		}
		if a.phaseFires(p.Starts, e) {
			a.phaseAt[i] = e.Time
		}
	}
}

// phaseFires reports whether this event is the trigger. A spell trigger is the
// enemies' own line -- a player casting the same id is not the boss changing
// phase -- and a health trigger is the boss's own health, read off the advanced
// block of any line that describes the unit the table names, which is the same
// identity the fight list reads a wipe percentage from.
func (a *Accumulator) phaseFires(start mechanics.PhaseStart, e event.Event) bool {
	if start.SpellID > 0 {
		if e.Spell.ID != start.SpellID || a.isPlayer(e.Source.GUID) {
			return false
		}
		switch start.On {
		case mechanics.OnCastStart:
			return e.Kind == event.CastStart
		case mechanics.OnCastSuccess:
			return e.Kind == event.CastSuccess
		case mechanics.OnAuraApplied:
			return e.Kind == event.AuraApplied
		case mechanics.OnAuraRemoved:
			return e.Kind == event.AuraRemoved
		}
		return false
	}
	if !e.Adv.OK || e.Adv.MaxHP <= 0 || e.Adv.InfoGUID == "" {
		return false
	}
	if a.name(e.Adv.InfoGUID) != a.opt.Mechanics.Name {
		return false
	}
	return float64(e.Adv.CurrentHP)/float64(e.Adv.MaxHP)*100 <= start.HealthPct
}

// phaseRows renders the fight's phases: each from its trigger to the next
// trigger or the fight's end, with the pull's own opening in front of them.
// A trigger that never fired is not a phase the fight reached and is left out.
// A fight that reached none at all has no phases: a single band called "Phase 1"
// across the whole pull says nothing a reader can act on, and a preset for it
// would be the whole-fight preset under another name.
func (a *Accumulator) phaseRows(durationMS int64) []Phase {
	out := []Phase{}
	if a.opt.Mechanics == nil || len(a.phaseAt) == 0 {
		return out
	}
	type fired struct {
		name string
		at   int64
	}
	rows := []fired{{name: firstPhaseName, at: 0}}
	for i, p := range a.opt.Mechanics.Phases {
		if t, ok := a.phaseAt[i]; ok {
			rows = append(rows, fired{name: p.Name, at: a.ms(t)})
		}
	}
	// Stable, and by instant rather than by table order: a curated list can name a
	// health threshold that a fight crossed before a cast the table lists first.
	sort.SliceStable(rows, func(i, j int) bool { return rows[i].at < rows[j].at })
	for i, r := range rows {
		end := durationMS
		if i+1 < len(rows) {
			end = rows[i+1].at
		}
		if end < r.at {
			end = r.at
		}
		out = append(out, Phase{Name: r.name, StartMS: r.at, EndMS: end})
	}
	return out
}
```

The file imports only `sort`, `event` and `mechanics`: nothing in it names `time.Time` in
a signature, so no `time` import.

In `logs/engine/summary/summary.go`:

- `Summary` gains, after `Mechanics`:

```go
	// Phases are the stretches this fight was fought in, from the encounter's
	// curated table. Empty for an encounter with no table, no phases, or a pull
	// that never reached one.
	Phases []Phase `json:"phases"`
```

- `Accumulator` gains:

```go
	// phaseAt is the first instant each of the encounter's phase triggers fired,
	// by its index in the table's Phases.
	phaseAt map[int]time.Time
```

- `New` initialises it: `phaseAt: map[int]time.Time{},`
- `addNow` calls it beside the other note functions:

```go
	a.noteMechanicCast(e)
	a.notePhase(e)
	a.noteTaunt(e)
```

- `Snapshot` sets it after `s.Mechanics`:

```go
	s.Mechanics = a.mechanicsBlock(s.Deaths)
	s.Phases = a.phaseRows(s.DurationMS)
	s.Roster = a.rosterRows(f, s)
```

- [ ] **Step 4: Run the tests, regenerate the goldens and the fixture**

```bash
cd /Users/jh/code/forever/logs && go test ./engine/summary/ && FOREVER_UPDATE_GOLDEN=1 go test ./... && go test ./... && (cd ../api && go test ./...)
export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH"
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && npx vitest run
```
Expected: PASS. `web/src/fixtures/report/fights/3/summary.json` now holds
`"phases": [{"name":"Phase 1","start_ms":0,"end_ms":14000},{"name":"Phase 2","start_ms":14000,"end_ms":60000}]`,
and `fights/4/summary.json` holds `"phases": []`. Read the fixture back and confirm those
two spans before moving on: Task 20's e2e asserts on them.

- [ ] **Step 5: Add the fixture assertion**

In `web/src/fixtures/report/fixture.test.ts`, append inside the main describe:

```ts
  it('splits the encounter into the phases its table names', () => {
    // 9001's table opens Phase 2 on the Warden's Anima Surge cast, which starts at
    // 20:12:14 -- fourteen seconds into the pull -- and is never completed. The opening
    // stretch is Phase 1 and needs no entry in the table.
    expect(three.phases).toEqual([
      { name: 'Phase 1', start_ms: 0, end_ms: 14_000 },
      { name: 'Phase 2', start_ms: 14_000, end_ms: 60_000 },
    ]);
    // Skolex has no table, so no phases and an empty list rather than a missing key.
    expect(four.phases).toEqual([]);
  });
```

and add `summary.phases` to `everyArray`'s list so the never-null invariant covers it.

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/web && npx vitest run src/fixtures/report/fixture.test.ts
cd /Users/jh/code/forever && git add logs/engine/summary web/src/fixtures/report
git commit -m "feat(logs): the summary carries each fight's phases"
```

---

### Task 19: The web reads phases into presets, the outcome and the night

**Files:**
- Modify: `web/src/lib/report/types.ts` (`Phase`, `Summary.phases`)
- Modify: `web/src/lib/report/window.ts` (`windowPresets` gains the phase chips)
- Modify: `web/src/lib/report/format.ts` (`phaseReached`, `outcomeLabel`'s second argument)
- Modify: `web/src/lib/report/night.ts` (`Night.phaseReached`)
- Test: `web/src/lib/report/window.test.ts`, `web/src/lib/report/format.test.ts`, `web/src/lib/report/night.test.ts`

**Interfaces:**
- Produces, in `types.ts`:

```ts
/** summary.Phase — one named stretch of a fight, from the encounter's curated table. */
export interface Phase {
  name: string;
  start_ms: number;
  end_ms: number;
}
```
  and `phases?: Phase[];` on `Summary` (absent from summaries written before engine 0.4.0).

- Produces, in `window.ts`:

```ts
export function windowPresets(summary: Summary, wholeDurationMs = summary.duration_ms): WindowPreset[]
```

- Produces, in `format.ts`:

```ts
/** The phase a fight ended in: the last one that started. '' when it has none. */
export function phaseReached(phases: readonly { name: string }[] | undefined): string;
export function outcomeLabel(fight: {…}, phaseReached?: string): string;
```

- Produces, in `night.ts`, on `Night`:

```ts
  /** Per fight index, the phase that pull reached; only pulls whose summary loaded. */
  phaseReached: Map<number, string>;
```

- [ ] **Step 1: Write the failing tests**

In `web/src/lib/report/window.test.ts`, extend the `windowPresets` describe:

```ts
  it('offers each phase as a preset, in order, after the fixed ones', () => {
    const phased: Summary = {
      ...summary,
      phases: [
        { name: 'Phase 1', start_ms: 0, end_ms: 14_000 },
        { name: 'Phase 2', start_ms: 14_000, end_ms: 60_000 },
      ],
    };
    const presets = windowPresets(phased);
    const phases = presets.filter((preset) => preset.label.startsWith('Phase'));
    expect(phases.map((preset) => preset.label)).toEqual([
      'Phase 1 · 0.0s to 14.0s',
      'Phase 2 · 14.0s to 1:00',
    ]);
    expect(phases[1].window).toEqual({ startMs: 14_000, endMs: 60_000 });
  });

  it('clamps a phase against the whole fight, not against a window already set', () => {
    const scopedToAWindow: Summary = {
      ...summary,
      duration_ms: 5000,
      phases: [{ name: 'Phase 2', start_ms: 14_000, end_ms: 60_000 }],
    };
    const [phase] = windowPresets(scopedToAWindow, 60_000).filter((preset) =>
      preset.label.startsWith('Phase'),
    );
    expect(phase.window).toEqual({ startMs: 14_000, endMs: 60_000 });
  });

  it('offers no phase presets for a fight that has none', () => {
    expect(windowPresets(summary).some((preset) => preset.label.startsWith('Phase'))).toBe(false);
  });
```

In `web/src/lib/report/format.test.ts`:

```ts
describe('the phase a fight reached', () => {
  it('is the last phase that started, and nothing when there are none', () => {
    expect(
      phaseReached([
        { name: 'Phase 1', start_ms: 0, end_ms: 14_000 },
        { name: 'Phase 2', start_ms: 14_000, end_ms: 60_000 },
      ]),
    ).toBe('Phase 2');
    expect(phaseReached([])).toBe('');
    expect(phaseReached(undefined)).toBe('');
  });

  it('a wipe’s outcome says which phase it got to; a kill’s does not need one', () => {
    const base = { kind: 'encounter', kill: false, in_progress: false, npc_kills: 0 };
    expect(outcomeLabel({ ...base, boss_health_pct: 59 }, 'Phase 3')).toBe('Wipe 59% · in Phase 3');
    expect(outcomeLabel({ ...base, boss_health_pct: 59 })).toBe('Wipe 59%');
    expect(outcomeLabel({ ...base, kill: true }, 'Phase 3')).toBe('Kill');
  });
});
```

(add `phaseReached` to the file's import from `./format`.)

In `web/src/lib/report/night.test.ts`:

```ts
  it('folds the phase each pull reached into the night, by fight index', () => {
    const fights = [fight(3, 'Kaal', false, 10_000), fight(4, 'Kaal', true, 10_000)];
    const summaries = new Map([
      [
        3,
        {
          ...summary(3, 10_000, [roster('A', 1)]),
          phases: [
            { name: 'Phase 1', start_ms: 0, end_ms: 4000 },
            { name: 'Phase 2', start_ms: 4000, end_ms: 10_000 },
          ],
        },
      ],
      [4, { ...summary(4, 10_000, [roster('A', 1)]), phases: [] }],
    ]);
    const night = aggregateNight(fights, summaries as never);
    expect(night.phaseReached.get(3)).toBe('Phase 2');
    expect(night.phaseReached.has(4)).toBe(false);
  });
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/window.test.ts src/lib/report/format.test.ts src/lib/report/night.test.ts`
Expected: FAIL — no phase presets, `phaseReached is not a function`, `night.phaseReached` undefined.

- [ ] **Step 3: Implement**

In `web/src/lib/report/types.ts`, before `Summary`:

```ts
/** summary.Phase — one named stretch of a fight, from the encounter's curated table. */
export interface Phase {
  name: string;
  start_ms: number;
  end_ms: number;
}
```

and on `Summary`:

```ts
  /** The stretches this fight was fought in. Absent from summaries written before engine 0.4.0. */
  phases?: Phase[];
```

In `web/src/lib/report/window.ts`, replace `windowPresets`:

```ts
/**
 * The window chips under the chart: the whole fight, the two fixed slices, each phase the
 * encounter's curated table named, and one per death.
 *
 * `wholeDurationMs` is the fight's own length, which is not `summary.duration_ms` when a
 * window is already set (scopeSummary narrows that to the window). The phases are the
 * fight's own spans and must be clamped against the fight, or a chip set inside one
 * window could not reach outside it.
 */
export function windowPresets(summary: Summary, wholeDurationMs = summary.duration_ms): WindowPreset[] {
  const duration = summary.duration_ms;
  const presets: WindowPreset[] = [
    { label: 'Whole fight', window: null },
    { label: 'First 30s', window: clampWindow({ startMs: 0, endMs: 30_000 }, duration) },
    { label: 'Last 30s', window: clampWindow({ startMs: duration - 30_000, endMs: duration }, duration) },
  ];
  for (const phase of summary.phases ?? []) {
    presets.push({
      label: `${phase.name} · ${formatDuration(phase.start_ms)} to ${formatDuration(phase.end_ms)}`,
      window: clampWindow({ startMs: phase.start_ms, endMs: phase.end_ms }, wholeDurationMs),
    });
  }
  for (const death of summary.deaths) {
    presets.push({
      // The time is part of the label: one player can die twice in a fight (a battle
      // rez), and two chips reading "Before Thalgrit died" leave the reader guessing.
      label: `20s before ${splitUnitName(death.name).name} died · ${formatDuration(death.at_ms)}`,
      window: deathWindow(death.at_ms, duration),
    });
  }
  return presets;
}
```

In `web/src/lib/report/format.ts`, add the helper and the argument:

```ts
/**
 * The phase a fight ended in: the last one that started, since the engine writes them in
 * order and runs the last one to the fight's end. '' when the encounter has no phases.
 */
export function phaseReached(phases: readonly { name: string }[] | undefined): string {
  return phases === undefined || phases.length === 0 ? '' : phases[phases.length - 1].name;
}

/** A fight's outcome word for a list: Kill, Wipe, Live, or how much trash died. */
export function outcomeLabel(
  fight: {
    kind: string;
    kill: boolean;
    in_progress: boolean;
    npc_kills: number;
    boss_health_pct?: number;
  },
  phase = '',
): string {
  if (fight.in_progress) return 'Live';
  if (fight.kind !== 'encounter') return `${fight.npc_kills} killed`;
  if (fight.kill) return 'Kill';
  // "Wipe 23%": how far the pull got, which is what separates one wipe from the next.
  // With phases curated for the boss, how far is also which phase it died in.
  const health = fight.boss_health_pct;
  const wipe = health !== undefined && health >= 0 ? `Wipe ${Math.round(health)}%` : 'Wipe';
  return phase === '' ? wipe : `${wipe} · in ${phase}`;
}
```

In `web/src/lib/report/night.ts`, add the import (the file has none from `format.ts` yet):

```ts
import { phaseReached } from './format';
```

the field to `Night`:

```ts
  /** Per fight index, the phase that pull reached; only pulls whose summary loaded and had one. */
  phaseReached: Map<number, string>;
```

fill it inside `aggregateNight`'s loop, just after `loaded += 1`:

```ts
    const reached = phaseReached(summary.phases);
    if (reached !== '') phases.set(fight.index, reached);
```

with `const phases = new Map<number, string>();` declared beside `players` and `bosses`
(a plain Map: it is built once and handed out whole, the same reasoning the rest of the
fold gives), and return it:

```ts
  return {
    fights: loaded,
    expected: encounters.length,
    time_ms: time,
    deaths,
    players: rows,
    bosses: [...bosses.values()],
    phaseReached: phases,
  };
```

- [ ] **Step 4: Run the tests to verify they pass**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"`
Expected: PASS, `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/lib/report/types.ts src/lib/report/window.ts src/lib/report/format.ts src/lib/report/night.ts src/lib/report/window.test.ts src/lib/report/format.test.ts src/lib/report/night.test.ts && npx prettier --write src/lib/report/types.ts src/lib/report/window.ts src/lib/report/format.ts src/lib/report/night.ts src/lib/report/window.test.ts src/lib/report/format.test.ts src/lib/report/night.test.ts
cd /Users/jh/code/forever && git add web/src/lib/report
git commit -m "feat(web): phases as window presets, in the outcome and folded over the night"
```

---

### Task 20: Phase bands on the chart, and the phase a wipe reached

**Files:**
- Modify: `web/src/components/report/TimeChart.svelte` (a `phases` prop: bands and names)
- Modify: `web/src/components/report/FightSelector.svelte` (the phase on a wipe's row)
- Modify: `web/src/components/report/ReportView.svelte` (wire the three of them)
- Modify: `web/src/components/report/Glossary.svelte` (one term)
- Test: `web/tests/e2e/report-tabs.spec.ts`, `web/tests/e2e/report-phone.spec.ts`

**Interfaces:**
- Consumes: `Summary.phases` (Task 19), `phaseReached`, `windowPresets`'s second argument,
  `Night.phaseReached`.
- Produces, on `TimeChart`:

```ts
/** The fight's phases, drawn as bands with their names at their left edge. */
phases?: { name: string; start_ms: number; end_ms: number }[];
```

- Produces, on `FightSelector`:

```ts
/** Per fight index, the phase that pull reached; a wipe's row says which. */
phaseOf?: ReadonlyMap<number, string>;
```

**Decision taken here:** the band's divider is drawn on the canvas (one more vertical rule,
in the line token) and the name is a positioned span over it, like `chart-scale` — a canvas
`fillText` would not use the page's font or its tokens and could not be read by a test.

- [ ] **Step 1: Write the failing e2e tests**

Append to `web/tests/e2e/report-tabs.spec.ts`:

```ts
test('the chart bands the fight’s phases and names them at their left edge', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  const bands = page.getByTestId('time-chart').getByTestId('phase-band');
  await expect(bands).toHaveCount(2);
  await expect(bands.first()).toHaveText('Phase 1');
  await expect(bands.nth(1)).toHaveText('Phase 2');
});

test('a phase is a window preset, so every table reads per phase in one click', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  await page
    .getByTestId('window-presets')
    .getByRole('button', { name: /Phase 2 · 14.0s to 1:00/ })
    .click();
  await expect(page).toHaveURL(/start=14000&end=60000/);
});

test('a fight with no phases shows no bands and no phase presets', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=4');
  await expect(page.getByTestId('phase-band')).toHaveCount(0);
  await expect(page.getByTestId('window-presets')).not.toContainText('Phase');
});
```

The fixture's only encounter with phases is a kill, so the wipe wording is covered by the
unit test in Task 19 rather than here; the fight list's own row is still assertable:

```ts
test('the fight list says nothing about a phase on a kill', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Kill');
});
```

Append to `web/tests/e2e/report-phone.spec.ts`:

```ts
test('a phase preset is a 44px target on a phone', async ({ page }) => {
  await page.goto(`${REPORT}`);
  const chip = page.getByTestId('window-presets').getByRole('button', { name: /Phase 2/ });
  await expect(chip).toBeVisible();
  const box = await chip.boundingBox();
  expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run:
```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-tabs.spec.ts --project=desktop
```
Expected: FAIL — `phase-band` not found, and no Phase 2 chip.

- [ ] **Step 3: Implement**

In `web/src/components/report/TimeChart.svelte`, add the prop:

```ts
    phases = [],
```

```ts
    /** The fight's phases, drawn as bands with their names at their left edge. */
    phases?: { name: string; start_ms: number; end_ms: number }[];
```

In `draw()`, after the ten-second grid and before the extra lines, draw one rule at each
phase's start (the first phase starts at zero and needs none) and shade alternate bands so
the boundary reads as a band rather than as one more grid line:

```js
    phases.forEach((phase, index) => {
      if (index % 2 === 1) {
        context.fillStyle = `${gold}0d`;
        context.fillRect(xOf(phase.start_ms), 0, xOf(phase.end_ms) - xOf(phase.start_ms), HEIGHT);
      }
      if (phase.start_ms <= 0) return;
      const x = Math.round(xOf(phase.start_ms)) + 0.5;
      context.strokeStyle = gold;
      context.lineWidth = 1;
      context.setLineDash([2, 2]);
      context.beginPath();
      context.moveTo(x, 0);
      context.lineTo(x, HEIGHT);
      context.stroke();
      context.setLineDash([]);
    });
```

`gold` is already read from the computed style at the top of `draw`. Add `phases` to the
redraw effect's dependency list.

The names go over the canvas, in the same `relative pl-12` wrapper as `chart-scale` and
the marks:

```svelte
    {#if phases.length > 0}
      <div class="pointer-events-none absolute inset-y-0 right-0 left-12" aria-hidden="true">
        {#each phases as phase (phase.name)}
          <span
            class="text-muted absolute top-0 font-mono text-[10px] leading-none"
            style={`left: calc(${durationMs === 0 ? 0 : (phase.start_ms / durationMs) * 100}% + 2px)`}
            data-testid="phase-band">{phase.name}</span
          >
        {/each}
      </div>
    {/if}
```

In `web/src/components/report/FightSelector.svelte`, add the prop and print it beside the
outcome on a wipe:

```ts
  let {
    fights,
    selected,
    onSelect,
    phaseOf = new Map<number, string>(),
  }: {
    fights: FightEntry[];
    selected: number;
    onSelect: (index: number) => void;
    /** Per fight index, the phase that pull reached; a wipe's row says which. */
    phaseOf?: ReadonlyMap<number, string>;
  } = $props();
```

and the outcome span becomes:

```svelte
              {outcome(fight, fight.kill ? '' : (phaseOf.get(fight.index) ?? ''))}
```

`const outcome = outcomeLabel;` already aliases the function, and `outcomeLabel`'s second
argument is optional, so every other caller is untouched.

In `web/src/components/report/ReportView.svelte`:

- import `phaseReached` from `format.ts`;
- pass the phases to the chart and the whole length to the presets:

```svelte
        <TimeChart
          series={chartSeries}
          extra={chartExtra}
          phases={summary.phases ?? []}
          durationMs={summary.duration_ms}
          window={timeWindow}
          deaths={summary.deaths
            .filter((death) => !playerSet.has(state.source) || death.guid === state.source)
            .map((death) => ({ at_ms: death.at_ms, name: death.name }))}
          label={chartLabel}
          onWindow={setWindow}
        />
```

`phases` is the one added line; every other prop keeps the value it already has.

```ts
  const presets = $derived(
    scoped === null ? [] : windowPresets(scoped, base?.duration_ms ?? scoped.duration_ms),
  );
```

- the header's outcome names the phase on a wipe:

```svelte
            data-testid="report-fight-outcome"
            >{outcomeLabel(fight, fight.kill ? '' : phaseReached(summary?.phases))}</span
          >
```

- the fight list gets whatever phases the page knows about: the night's fold for every
  loaded pull, and the selected fight's own on a single pull.

```ts
  /**
   * Per fight index, the phase that pull reached. Over the night every pull's summary is
   * folded, so the whole list can say; on a single pull only the one on screen is loaded,
   * and the others' rows say nothing rather than guessing.
   */
  const phaseOf = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const out = new Map<number, string>(night?.phaseReached ?? []);
    const own = phaseReached(summary?.phases);
    if (own !== '' && fight !== null) out.set(fight.index, own);
    return out;
  });
```

```svelte
    <FightSelector {fights} selected={state.fight} onSelect={(index) => patch({ fight: index })} {phaseOf} />
```

In `web/src/components/report/Glossary.svelte`, add:

```ts
    {
      term: 'Phase',
      meaning:
        'A named stretch of a boss fight, from the curated table for that encounter: the cast, debuff or health percentage that opens it is a person’s judgement, the same way the mechanics table is. Phase 1 is the pull itself. Each phase is a band on the chart and a window preset, so every table can be read per phase, and a wipe’s outcome says which phase it got to. A boss nobody has curated phases for has none, and the page reads exactly as it did before.',
    },
```

- [ ] **Step 4: Run the tests to verify they pass**

```bash
cd /Users/jh/code/forever/web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-tabs.spec.ts tests/e2e/report-brush.spec.ts tests/e2e/report-threat.spec.ts --project=desktop && npx playwright test tests/e2e/report-phone.spec.ts --project=mobile && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: PASS, `- 0 errors`. `report-brush.spec.ts` is in the list because the presets
strip gained chips and anything that counts them must still hold.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/components/report/TimeChart.svelte src/components/report/FightSelector.svelte src/components/report/ReportView.svelte src/components/report/Glossary.svelte tests/e2e/report-tabs.spec.ts tests/e2e/report-phone.spec.ts && npx prettier --write src/components/report/TimeChart.svelte src/components/report/FightSelector.svelte src/components/report/ReportView.svelte src/components/report/Glossary.svelte tests/e2e/report-tabs.spec.ts tests/e2e/report-phone.spec.ts
cd /Users/jh/code/forever && git add web/src/components/report web/tests/e2e
git commit -m "feat(web): phase bands on the chart, phase presets, and the phase a wipe reached"
```

---

### Task 21: Compare aligns two pulls by phase

**Files:**
- Modify: `web/src/lib/report/compare.ts` (`sharedPhases`, `phaseWindow`)
- Modify: `web/src/lib/report/compare.test.ts`
- Modify: `web/src/components/report/CompareMode.svelte` (the phase picker)
- Test: `web/tests/e2e/report-compare.spec.ts`

**Interfaces:**
- Consumes: `Summary.phases` (Task 19), `TimeWindow`, `clampWindow`, `scopeSummary`.
- Produces, in `compare.ts`:

```ts
/** The phase names both pulls have, in the left pull's order; empty when they share none. */
export function sharedPhases(left: Summary | null, right: Summary | null): string[];
/** One pull's own span for a named phase, or null when it has none of that name. */
export function phaseWindow(summary: Summary | null, name: string): TimeWindow | null;
```

**Known fixture limitation:** the fixture has one encounter with phases (9001) and one
without (9002), so no two of its fights share a phase name. The positive path — two pulls
of one boss aligned by phase — is covered by the unit tests here; the e2e covers the
picker's absence and its wording. The harness smoke against the sample log (two General
Kaal pulls) is where the picker is exercised end to end, once 2363's table has phases.

- [ ] **Step 1: Write the failing tests**

Append to `web/src/lib/report/compare.test.ts`:

```ts
describe('aligning two pulls by phase', () => {
  const phased = (phases: { name: string; start_ms: number; end_ms: number }[]): Summary => ({
    ...summary([]),
    phases,
  });
  const left = phased([
    { name: 'Phase 1', start_ms: 0, end_ms: 20_000 },
    { name: 'Phase 2', start_ms: 20_000, end_ms: 90_000 },
  ]);
  const right = phased([
    { name: 'Phase 1', start_ms: 0, end_ms: 35_000 },
    { name: 'Phase 2', start_ms: 35_000, end_ms: 60_000 },
    { name: 'Phase 3', start_ms: 60_000, end_ms: 120_000 },
  ]);

  it('offers only the phases both pulls reached, in the first pull’s order', () => {
    expect(sharedPhases(left, right)).toEqual(['Phase 1', 'Phase 2']);
  });

  it('gives each side its own span for the named phase', () => {
    expect(phaseWindow(left, 'Phase 2')).toEqual({ startMs: 20_000, endMs: 90_000 });
    expect(phaseWindow(right, 'Phase 2')).toEqual({ startMs: 35_000, endMs: 60_000 });
  });

  it('has nothing to offer when a side has no phases at all', () => {
    expect(sharedPhases(left, phased([]))).toEqual([]);
    expect(sharedPhases(left, null)).toEqual([]);
    expect(phaseWindow(phased([]), 'Phase 2')).toBeNull();
    expect(phaseWindow(null, 'Phase 2')).toBeNull();
  });
});
```

Append to `web/tests/e2e/report-compare.spec.ts`:

```ts
// The fixture's two encounters are a boss with phases and a boss with none, so they share
// no phase: the picker says so rather than offering an alignment it cannot make.
test('the phase picker says when the two pulls share no phase', async ({ page }) => {
  await page.goto(`${COMPARE}&with=4`);
  await expect(page.getByTestId('compare-phase')).toHaveCount(0);
  await expect(page.getByTestId('compare-phase-note')).toContainText('no phase in common');
});

test('there is no phase picker before a second fight is picked', async ({ page }) => {
  await page.goto(COMPARE);
  await expect(page.getByTestId('compare-phase')).toHaveCount(0);
  await expect(page.getByTestId('compare-phase-note')).toHaveCount(0);
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd /Users/jh/code/forever/web && npx vitest run src/lib/report/compare.test.ts`
Expected: FAIL — `sharedPhases is not a function`.

- [ ] **Step 3: Implement**

In `web/src/lib/report/compare.ts`, add the import and the two functions:

```ts
import type { TimeWindow } from './window';
```

```ts
/**
 * The phase names both pulls reached, in the first pull's order. Aligning by phase only
 * makes sense for a phase both sides got to: a pull that wiped in Phase 2 has no Phase 3
 * to compare, and offering one would compare a real span against nothing.
 */
export function sharedPhases(left: Summary | null, right: Summary | null): string[] {
  const theirs = new Set((right?.phases ?? []).map((phase) => phase.name));
  return (left?.phases ?? []).map((phase) => phase.name).filter((name) => theirs.has(name));
}

/** One pull's own span for a named phase, or null when it has none of that name. */
export function phaseWindow(summary: Summary | null, name: string): TimeWindow | null {
  const found = (summary?.phases ?? []).find((phase) => phase.name === name);
  return found === undefined ? null : { startMs: found.start_ms, endMs: found.end_ms };
}
```

In `web/src/components/report/CompareMode.svelte`, add the imports and the state:

```ts
  // Added to the existing import from compare.ts (Task 7 already imports the metric
  // helpers and the diff builders from it).
  import { phaseWindow, sharedPhases } from '../../lib/report/compare';
```

```ts
  /**
   * The phase both sides are read over, when one is picked: each side is scoped to its
   * own span for it, so a pull whose Phase 2 ran twice as long is still compared phase
   * against phase. Not in the url: it is a way of reading the two fights on screen, and
   * a link already carries which two they are.
   */
  let phase = $state('');
  const phases = $derived(sharedPhases(leftWhole, rightWhole));
  // A second fight picked while a phase is chosen may not have that phase.
  $effect(() => {
    if (phase !== '' && !phases.includes(phase)) phase = '';
  });
```

and give each side its own window. Task 7 declared `leftWindow` and `rightWindow` as the
page's one `window`; replace both bodies:

```ts
  /** Each side in its own phase's span when one is picked, else in the page's window. */
  const leftWindow = $derived(phase === '' ? window : phaseWindow(leftWhole, phase));
  const rightWindow = $derived(phase === '' ? window : phaseWindow(rightWhole, phase));
  const left = $derived(
    leftWindow === null ? leftWhole : scopeSummary(leftWhole, clampWindow(leftWindow, leftWhole.duration_ms)),
  );
  const right = $derived(
    rightWhole === null || rightWindow === null
      ? rightWhole
      : scopeSummary(rightWhole, clampWindow(rightWindow, rightWhole.duration_ms)),
  );
```

`fightLabel` reads `window` for the stretch it prints; give it the side's own:

```ts
  function fightLabel(fight: FightEntry | null, own: TimeWindow | null): string {
    if (fight === null) return '';
    const stretch =
      own === null
        ? formatDuration(fight.duration_ms)
        : `${formatDuration(own.startMs)} to ${formatDuration(Math.min(own.endMs, fight.duration_ms))} of ${formatDuration(fight.duration_ms)}`;
    return `${fight.name} · ${stretch} · ${outcomeLabel(fight).toLowerCase()}`;
  }
```

and its two call sites become `fightLabel(currentFight, leftWindow)` and
`fightLabel(rightFight, rightWindow)`.

Add the picker beside the metric, and the note when there is nothing to pick:

```svelte
    {#if rightWhole !== null}
      {#if phases.length > 0}
        <label class="label text-muted flex items-center gap-2" for="compare-phase">
          Phase
          <select
            id="compare-phase"
            class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
            data-testid="compare-phase"
            value={phase}
            onchange={(event) => (phase = (event.currentTarget as HTMLSelectElement).value)}
          >
            <option value="">The whole pull</option>
            {#each phases as name (name)}
              <option value={name}>{name}</option>
            {/each}
          </select>
        </label>
      {:else}
        <p class="text-muted text-[12px]" data-testid="compare-phase-note">
          These two pulls have no phase in common, so there is nothing to align by: either their
          bosses have no curated phases, or one pull did not reach the other's.
        </p>
      {/if}
    {/if}
```

and the scope line says which reading is on screen:

```svelte
  <p class="text-muted text-[12px]" data-testid="compare-scope">
    {#if phase !== ''}
      Each side shows its own {phase}: {formatDuration(leftWindow?.startMs ?? 0)} to {formatDuration(
        leftWindow?.endMs ?? 0,
      )} of this pull against {formatDuration(rightWindow?.startMs ?? 0)} to {formatDuration(
        rightWindow?.endMs ?? 0,
      )} of the other, so a long phase and a short one are read phase against phase.
    {:else if window === null}
      Both sides show the whole fight. Set a window in Analyze to compare the same stretch of each pull.
    {:else}
      Both sides show {formatDuration(window.startMs)} to {formatDuration(window.endMs)} of each fight, so a long
      wipe and a short one are read over the same stretch. Clear the window in Analyze to compare whole fights.
    {/if}
  </p>
```

- [ ] **Step 4: Run the tests to verify they pass**

```bash
cd /Users/jh/code/forever/web && npx vitest run src/lib/report/compare.test.ts && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-compare.spec.ts --project=desktop && npx playwright test tests/e2e/report-phone.spec.ts --project=mobile && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: PASS, `- 0 errors`.

- [ ] **Step 5: Lint, format, commit**

```bash
cd /Users/jh/code/forever/web && npx eslint src/lib/report/compare.ts src/lib/report/compare.test.ts src/components/report/CompareMode.svelte tests/e2e/report-compare.spec.ts && npx prettier --write src/lib/report/compare.ts src/lib/report/compare.test.ts src/components/report/CompareMode.svelte tests/e2e/report-compare.spec.ts
cd /Users/jh/code/forever && git add web/src/lib/report web/src/components/report web/tests/e2e/report-compare.spec.ts
git commit -m "feat(web): Compare aligns two pulls by phase"
```

---

### Task 22: The drafting tool proposes an encounter's phase candidates

**Files:**
- Modify: `logs/cmd/forever-logs/mechanics_draft.go` (the draft gains a `phases` list)
- Test: `logs/cmd/forever-logs/mechanics_draft_test.go`

**Interfaces:**
- Consumes: the parse path `runMechanicsDraft` already uses to turn a log into fights and
  summaries, plus each fight's `summary.Casts` and `summary.Auras`, and `fight.Fight`'s
  own `Name` for the boss's health readings.
- Produces: the emitted `mechanics.Table` carries a `phases` array of candidates, each
  with a `name` of `"Phase N"` in the order they happen and a `starts` in the spec's
  format, and each carrying its evidence in the table's own JSON so a curator can judge
  it. Because `mechanics.Phase` has no note field, the evidence goes to stderr, one line
  per candidate, keyed by the phase name.

**Decision taken here:** a candidate is an enemy cast (`cast_start`, which is what a
channel logs) or an enemy aura application that happened **exactly once in every pull of
that encounter** in the log, at a boss health within ten percentage points across those
pulls. That is the spec's rule — "the enemy casts and aura changes that happen exactly
once per pull at a consistent boss health" — written as a test.

- [ ] **Step 1: Write the failing test**

Append to `logs/cmd/forever-logs/mechanics_draft_test.go`:

```go
// The fixture log has one pull of Warden Kelthas. Its Anima Surge cast starts once and
// only once, so it is exactly the shape a phase trigger has; Frostbolt is a player's and
// is not a candidate at all.
func TestMechanicsDraftProposesPhaseCandidates(t *testing.T) {
	var out, errOut bytes.Buffer
	err := run([]string{"mechanics-draft", "-encounter", "9001",
		"../../../web/src/fixtures/report/fixture.log"}, &out, &errOut)
	if err != nil {
		t.Fatal(err, errOut.String())
	}
	table, err := mechanics.Parse(out.Bytes())
	if err != nil {
		t.Fatalf("the draft must be a valid table: %v\n%s", err, out.String())
	}
	var surge *mechanics.Phase
	for i := range table.Phases {
		if table.Phases[i].Starts.SpellID == 334653 {
			surge = &table.Phases[i]
		}
	}
	if surge == nil {
		t.Fatalf("Anima Surge must be a phase candidate: %+v", table.Phases)
	}
	if surge.Starts.On != mechanics.OnCastStart {
		t.Errorf("a channel's candidate must key on its start, not its success: %+v", surge.Starts)
	}
	if surge.Name != "Phase 2" {
		t.Errorf("the first candidate after the pull is Phase 2, got %q", surge.Name)
	}
	for _, p := range table.Phases {
		if p.Starts.SpellID == 116 {
			t.Errorf("a player's spell is not a phase candidate: %+v", p)
		}
	}
	if !strings.Contains(errOut.String(), "Phase 2") {
		t.Errorf("the evidence for each candidate must reach stderr:\n%s", errOut.String())
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/forever/logs && go test ./cmd/forever-logs/ -run TestMechanicsDraftProposesPhaseCandidates`
Expected: FAIL — `table.Phases` is empty, so `surge` is nil.

- [ ] **Step 3: Implement**

In `logs/cmd/forever-logs/mechanics_draft.go`, after the mechanics are drafted for an
encounter, gather the candidates across that encounter's pulls and attach them:

```go
// phaseCandidate is one spell that could open a phase: how it was seen, when, and on
// how many of the encounter's pulls.
type phaseCandidate struct {
	spellID   int64
	name      string
	on        string
	pulls     int
	firstMS   int64
	healthPct []float64
}

// draftPhases proposes phase triggers for one encounter. The rule is the spec's: an
// enemy cast or aura application that happened exactly once in every pull, at a boss
// health that did not move much between them. Anything that happened twice in a pull is
// a rotation, not a phase; anything that happened in one pull of five is a fluke.
func draftPhases(pulls []draftPull) ([]mechanics.Phase, []string) {
	seen := map[int64]*phaseCandidate{}
	for _, pull := range pulls {
		once := map[int64]*phaseCandidate{}
		for _, row := range pull.enemyCasts {
			count(once, row.SpellID, row.SpellName, mechanics.OnCastStart, row.firstMS, row.count, pull.bossPctAt(row.firstMS))
		}
		for _, row := range pull.enemyAuras {
			count(once, row.SpellID, row.SpellName, mechanics.OnAuraApplied, row.firstMS, row.count, pull.bossPctAt(row.firstMS))
		}
		for id, c := range once {
			if c.pulls == 0 {
				continue
			}
			found := seen[id]
			if found == nil {
				seen[id] = c
				continue
			}
			found.pulls++
			found.healthPct = append(found.healthPct, c.healthPct...)
			if c.firstMS < found.firstMS {
				found.firstMS = c.firstMS
			}
		}
	}
	kept := []*phaseCandidate{}
	for _, c := range seen {
		if c.pulls == len(pulls) && spread(c.healthPct) <= 10 {
			kept = append(kept, c)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].firstMS < kept[j].firstMS })
	phases := make([]mechanics.Phase, 0, len(kept))
	evidence := make([]string, 0, len(kept))
	for i, c := range kept {
		name := fmt.Sprintf("Phase %d", i+2)
		phases = append(phases, mechanics.Phase{
			Name:   name,
			Starts: mechanics.PhaseStart{SpellID: c.spellID, On: c.on},
		})
		evidence = append(evidence, fmt.Sprintf(
			"%s: %s (%d) %s once on each of %d pulls, first at %s, boss at %.0f%%-%.0f%%",
			name, c.name, c.spellID, c.on, c.pulls, formatMS(c.firstMS), minOf(c.healthPct), maxOf(c.healthPct)))
	}
	return phases, evidence
}
```

`draftPull`, `count`, `spread`, `minOf`, `maxOf` and `formatMS` are small helpers this
file writes alongside `draftPhases`: `draftPull` holds one pull's enemy cast rows and
enemy aura tracks with their first offsets and counts (read off the pull's
`summary.Casts` and `summary.Auras`, keeping only rows whose GUID the registry does not
know as a player) and its boss health readings; `count` records a spell in `once` only
when its count in that pull is exactly one, setting `pulls` to 1; `spread` is
`max - min` of a float slice and 0 for fewer than two entries; `formatMS` prints
`m:ss`. Write them to match the file's existing style — `runMechanicsDraft` already walks
the parsed fights and their summaries, so reuse that walk rather than parsing twice.

Set the result on the emitted table and print the evidence:

```go
	table.Phases, phaseEvidence = draftPhases(pulls)
	for _, line := range phaseEvidence {
		fmt.Fprintln(errOut, line)
	}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd /Users/jh/code/forever/logs && go test ./cmd/forever-logs/`
Expected: PASS.

- [ ] **Step 5: Draft the sample encounters' phases and curate them**

```bash
cd /Users/jh/code/forever/logs && go run ./cmd/forever-logs mechanics-draft /Users/jh/code/forever/scratchpad/logs/wowp.txt > /private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/phase-drafts.json
```
Read the drafted candidates for 2363 (General Kaal) and 2360 (Kryxis) and the evidence on
stderr, keep the ones that are really phase changes, name them what the community calls
them, and write them into `logs/engine/mechanics/tables/2363.json` and `2360.json`. A
candidate whose evidence does not convince you is left out: the tables are curated, and a
wrong phase band is worse than none.

- [ ] **Step 6: Run the tests and commit**

```bash
cd /Users/jh/code/forever/logs && go test ./... && (cd ../api && go test ./...)
git add logs/cmd/forever-logs logs/engine/mechanics/tables
git commit -m "feat(logs): forever-logs mechanics-draft proposes an encounter's phases"
```

- [ ] **Step 7: Group 6 review gate**

```bash
cd /Users/jh/code/forever/logs && go test ./... && (cd ../api && go test ./...)
cd /Users/jh/code/forever/web && npm run make:report-fixture && npx prettier --write src/fixtures/report && git status --short src/fixtures/report
npx vitest run && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?" && npx playwright test --project=desktop && npx playwright test tests/e2e/report-phone.spec.ts --project=mobile
```
The fixture regeneration must produce no diff unless Step 5 changed 9001's table, which it
should not. Merge `phases` into `main`.

---

## After all six have merged

- [ ] **Harness smoke, one per feature.** Re-parse `scratchpad/logs/wowp.txt`, stage it the
  way the persona harness does, and reconcile each new figure against the parquet in
  DuckDB, the way the swing and absorb bugs were found:
  1. **Threat series.** For a handful of pairs in `fights/20/summary.json`, check
     `sum(series) == round(threat)` and that the last bucket's cumulative equals the total.
  2. **Compare.** For one player across two General Kaal pulls, check the ability diff's
     two columns against `SELECT spell_name, sum(amount - greatest(overkill,0)) FROM
     read_parquet('events.parquet') WHERE source_guid = … GROUP BY 1` on each fight.
  3. **Resources.** Check `wasted` against `SELECT sum(over_energize) FROM … WHERE
     kind = 'energize' AND dest_guid = … AND power_type = …`, and `at_max_ms` against the
     seconds the advanced block reads `adv_current_power = adv_max_power`.
  4. **Pet casts and refreshes.** Check that every cast row's `owner_guid` matches the
     report's `units` owner, and that the events view's refresh count matches
     `SELECT count(*) FROM … WHERE kind = 'aura_refresh'`.
  5. **Ability lines.** Put one ability on the chart and check its peak second against
     `SELECT floor(fight_ms/1000), sum(...) … GROUP BY 1 ORDER BY 2 DESC LIMIT 1`.
  6. **Phases.** With 2363's and 2360's phases curated in Task 22, check each phase's
     `start_ms` in `fights/*/summary.json` against the line that triggers it in the log,
     and check that a wipe's phase in the fight list is the phase its last event fell in.
- [ ] **Persona review.** Run the five personas once, after all six have merged, not after
  each — as the spec's Testing section asks.

## Out of scope, and why (from the spec)

Tank stance and taunt threat multipliers wait for Forever's ability data; rankings depth is
a population problem; replay needs positions the engine does not record; spell and item
tooltips wait for Forever's data files. Nothing in this plan adds any of them, and the
threat chart's note says out loud that the base model does not move a taunter to the top.
