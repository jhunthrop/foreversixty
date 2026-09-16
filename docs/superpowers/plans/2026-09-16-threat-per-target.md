# Threat Per Target Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** The Threat tab shows each player's threat on each enemy, not only totals, and the timeline and the Threat tab show every taunt, so a tank can see who pulled a mob off them and when.

**Architecture:** The engine already turns damage and healing into threat per player through `ThreatModel`. This plan credits each event's threat to the enemy it concerns as well: damage threat to the unit hit, healing threat spread over the hostile units the raid was engaged with at that moment. Taunts are read off cast-success lines for a small list of taunt spell ids. Both land in `summary.json` (`threat_by_target`, `taunts`); the web adds a target picker to the Threat tab, a taunt list under it, and taunt marks on the timeline; the night fold sums pairs by enemy name.

**Tech Stack:** Go (`logs/engine`), Svelte 5 + Astro 7 (`web/`), vitest, Playwright.

**Spec:** this plan is its own design; the tank persona's reviews (`scratchpad/persona/tank-r*/review.md`) are the requirement: "each player's threat on each mob, plus a taunt row on the timeline with off times".

## Global Constraints

- Work on the `mechanics-mode` worktree branch after the mechanics plan, or on its own branch `threat-per-target`; never on `main` while the persona review loop runs its harness from `main`.
- Threat stays the model's: this plan changes attribution, not coefficients. `ThreatModel.Damage`/`Healing` are called exactly as today; the per-player totals must not change (the golden's `threat` rows stay identical).
- "Engaged" hostile units: a hostile unit that dealt or took damage within the last 10 seconds (`EngagedWindow`, an `Options` field, default `10 * time.Second`).
- Taunt spell ids (`taunts.go`): 355 Taunt, 694 Mocking Blow, 1161 Challenging Shout, 5209 Challenging Roar, 6795 Growl, 17735 Suffering, 31789 Righteous Defense, 56222 Dark Command, 62124 Hand of Reckoning, 115546 Provoke, 185245 Torment, 116189 Provoke (statue). Keep the list in one place with a comment that it is data awaiting Forever's own spell ids.
- Every summary change bumps `logs/engine/session/session.go` `Version` (`0.3.0` → `0.3.1` if after the mechanics plan; otherwise `0.2.4` → `0.3.0`), regenerates goldens (`FOREVER_UPDATE_GOLDEN=1 go test ./...` in `logs/`) and the web fixture (`node scripts/make-report-fixture.mjs` in `web/`).
- The persona review harness serves the `main` build on port 4321. In the worktree, never run `scratchpad/local-report.sh`, and run Playwright against your own preview on port 4322: `E2E_PORT=4322 npx astro preview --port 4322` with `E2E_PORT=4322 npx playwright test ...`. The first task of whichever plan runs first makes `web/playwright.config.ts` read `process.env.E2E_PORT ?? '4321'` for both `use.baseURL` and `webServer.port`/`command` (one small commit: `chore(web): e2e port from E2E_PORT`).
- Run only the tests for files you touch; CI runs the rest. After `npm run build`, check `ls dist/report-island.js`. Type check with `NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"`.

---

### Task 1: Threat per target and taunts in the engine

**Files:**
- Modify: `logs/engine/summary/threat.go` (new row types)
- Create: `logs/engine/summary/taunts.go`
- Modify: `logs/engine/summary/summary.go` (Options.EngagedWindow, Accumulator fields, Snapshot)
- Modify: `logs/engine/summary/damage.go:150-176` (credit per target; track engagement)
- Modify: `logs/engine/summary/summary.go` `Add` (record taunts from cast-success lines)
- Modify: `logs/engine/session/session.go` `Version`
- Test: `logs/engine/summary/threat_test.go`

**Interfaces:**
- Produces on `Summary`:

```go
// ThreatPair is one player's threat on one enemy for the fight.
type ThreatPair struct {
	GUID       string  `json:"guid"`
	Name       string  `json:"name"`
	TargetGUID string  `json:"target_guid"`
	TargetName string  `json:"target_name"`
	Threat     float64 `json:"threat"`
}
// Taunt is one taunt cast: who taunted what, when.
type Taunt struct {
	AtMS       int64  `json:"at_ms"`
	SourceGUID string `json:"source_guid"`
	SourceName string `json:"source_name"`
	TargetGUID string `json:"target_guid"`
	TargetName string `json:"target_name"`
	SpellID    int64  `json:"spell_id"`
	SpellName  string `json:"spell_name"`
}
// on Summary:
ThreatByTarget []ThreatPair `json:"threat_by_target"`
Taunts         []Taunt      `json:"taunts"`
```

- `Options.EngagedWindow time.Duration` (default 10s).

- [ ] **Step 1: Write the failing tests**

```go
// logs/engine/summary/threat_test.go
package summary

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

// Two enemies, one healer. Damage threat lands on the unit hit; healing threat is
// spread over the enemies engaged in the last ten seconds; a taunt is recorded.
func TestThreatIsCreditedPerTargetAndTauntsAreKept(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	const bossB = "Creature-0-2085-2284-7855-169754-0000AA0002"
	events := []event.Event{
		dmg(1, tank, boss, 1, "Melee", 1000, -1),
		dmg(2, mage, bossB, 116, "Frostbolt", 500, -1),
		heal(3, healer, tank, 2050, "Holy Light", 400, 0),
		cast(4, tank, boss, 355, "Taunt"), // add a cast() helper if summary_test.go lacks one
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(5), Players: []string{tank, mage, healer}}, "test")

	pair := func(guid, target string) float64 {
		for _, p := range s.ThreatByTarget {
			if p.GUID == guid && p.TargetGUID == target {
				return p.Threat
			}
		}
		return -1
	}
	if got := pair(tank, boss); got != 1000 {
		t.Errorf("tank on boss = %v, want 1000 (a point of damage is a point of threat)", got)
	}
	if got := pair(mage, bossB); got != 500 {
		t.Errorf("mage on bossB = %v, want 500", got)
	}
	// 400 effective healing at 0.5 = 200 threat, over the two engaged enemies: 100 each.
	if got := pair(healer, boss); got != 100 {
		t.Errorf("healer on boss = %v, want 100", got)
	}
	if got := pair(healer, bossB); got != 100 {
		t.Errorf("healer on bossB = %v, want 100", got)
	}
	// The per-player totals are untouched by attribution.
	for _, row := range s.Threat {
		if row.GUID == healer && row.Threat != 200 {
			t.Errorf("healer total = %v, want 200", row.Threat)
		}
	}
	if len(s.Taunts) != 1 || s.Taunts[0].SourceGUID != tank || s.Taunts[0].TargetGUID != boss || s.Taunts[0].AtMS != 4000 {
		t.Fatalf("taunts = %+v", s.Taunts)
	}
}
```

Read `summary_test.go`'s helpers (`opts`, `dmg`, `heal`, the guids `tank`, `mage`, `healer`, `boss`, the flags they use) and match them. If there is no `cast()` helper, add one that builds an `event.Event{Kind: event.CastSuccess, Source, Dest, Spell}` with the same flags `dmg` uses (players `0x511`, enemies `0xa48`).

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd logs && go test ./engine/summary/ -run TestThreatIsCredited`
Expected: FAIL, `ThreatByTarget` undefined.

- [ ] **Step 3: Implement**

`summary.go` `Options`: add

```go
	// EngagedWindow is how long after dealing or taking damage a hostile unit
	// still counts as engaged, for spreading healing threat. Default ten seconds.
	EngagedWindow time.Duration
```

default it in `New` (`if o.EngagedWindow == 0 { o.EngagedWindow = 10 * time.Second }`). `Accumulator` gains `threatBy map[string]map[string]float64` (player → enemy → threat), `engaged map[string]time.Time` (hostile guid → last damage instant), `taunts []Taunt`; initialise the maps in `New`.

`damage.go`, `case event.Damage:` after the existing `a.threat[src] += ...`:

```go
		th := a.opt.Threat.Damage(e)
		a.creditThreat(src, e.Dest.GUID, th)
		a.engage(e.Source.GUID, e.Time)
		a.engage(e.Dest.GUID, e.Time)
```

and in `case event.Heal:` after `a.threat[src] += a.opt.Threat.Healing(e)`:

```go
		a.spreadThreat(src, a.opt.Threat.Healing(e), e.Time)
```

`threat.go` additions:

```go
// creditThreat books threat from one player against one enemy.
func (a *Accumulator) creditThreat(player, enemy string, threat float64) {
	if threat == 0 || enemy == "" {
		return
	}
	if u, ok := a.opt.Registry.Get(enemy); !ok || !units.Hostile(u.Flags) {
		return
	}
	by := a.threatBy[player]
	if by == nil {
		by = map[string]float64{}
		a.threatBy[player] = by
	}
	by[enemy] += threat
}

// engage marks a hostile unit as in the fight now.
func (a *Accumulator) engage(guid string, at time.Time) {
	if u, ok := a.opt.Registry.Get(guid); ok && units.Hostile(u.Flags) {
		a.engaged[guid] = at
	}
}

// spreadThreat books healing threat over every hostile unit engaged within the window.
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
	for _, guid := range live {
		a.creditThreat(player, guid, each)
	}
}

// threatPairs renders the per-target table, largest first.
func (a *Accumulator) threatPairs() []ThreatPair {
	out := []ThreatPair{}
	for player, by := range a.threatBy {
		for enemy, threat := range by {
			out = append(out, ThreatPair{GUID: player, Name: a.name(player), TargetGUID: enemy, TargetName: a.name(enemy), Threat: threat})
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

`taunts.go`:

```go
// logs/engine/summary/taunts.go
package summary

import "github.com/jhunthrop/foreversixty/logs/engine/event"

// tauntSpells are the taunts the deaths and threat views mark. Data awaiting
// Forever's own spell ids; the retail sample's Hand of Reckoning and Provoke are
// here so the feature can be reviewed against it.
var tauntSpells = map[int64]bool{
	355: true, 694: true, 1161: true, 5209: true, 6795: true, 17735: true,
	31789: true, 56222: true, 62124: true, 115546: true, 116189: true, 185245: true,
}

// noteTaunt records a taunt cast at an enemy. Called from Add for every event.
func (a *Accumulator) noteTaunt(e event.Event) {
	if e.Kind != event.CastSuccess || !tauntSpells[e.Spell.ID] || e.Dest.GUID == "" {
		return
	}
	a.taunts = append(a.taunts, Taunt{
		AtMS: a.ms(e.Time), SourceGUID: e.Source.GUID, SourceName: a.name(e.Source.GUID),
		TargetGUID: e.Dest.GUID, TargetName: a.name(e.Dest.GUID), SpellID: e.Spell.ID, SpellName: e.Spell.Name,
	})
}
```

In `Add`, call `a.noteTaunt(e)` beside `a.addCastsAndExchanges(e)`. In `Snapshot`, set `ThreatByTarget: a.threatPairs()` and `Taunts: copySlice(a.taunts)` (use `[]Taunt{}` when nil so the JSON is `[]`, matching the other slices). Confirm `event.CastSuccess` is the constant's name (`event.go` line 65: `CastSuccess: "cast_success"`). Bump `Version`.

- [ ] **Step 4: Run the tests and regenerate goldens**

Run: `cd logs && go test ./engine/summary/ && FOREVER_UPDATE_GOLDEN=1 go test ./... && go test ./...`
Expected: PASS; the golden gains `threat_by_target` and `taunts`; existing `threat` rows unchanged (diff the golden to confirm).

- [ ] **Step 5: Commit**

```bash
git add logs/engine/summary logs/engine/session
git commit -m "feat(logs): threat per target and taunts in the summary"
```

---

### Task 2: Web types, the window scale and the night fold

**Files:**
- Modify: `web/src/lib/report/types.ts` (Summary gains `threat_by_target`, `taunts`)
- Modify: `web/src/lib/report/window.ts:201-214` (scale pairs with their player's ratio, like `scopeThreat`)
- Modify: `web/src/lib/report/night.ts:298-306` (fold pairs by player guid and enemy name; taunts shifted by offset with the pull label)
- Test: `web/src/lib/report/night.test.ts`, `web/src/lib/report/window.test.ts`

**Interfaces:**
- Produces:

```ts
export interface ThreatPair { guid: string; name: string; target_guid: string; target_name: string; threat: number }
export interface Taunt { at_ms: number; source_guid: string; source_name: string; target_guid: string; target_name: string; spell_id: number; spell_name: string; label?: string }
// on Summary (absent before engine 0.3.x):
threat_by_target?: ThreatPair[];
taunts?: Taunt[];
```

- [ ] **Step 1: Write the failing tests**

In `window.test.ts` (read its existing `scopeSummary` fixtures):

```ts
it('scales a player’s per-target threat by the same ratio as their total', () => {
  const summary = { ...base, threat: [{ guid: 'P1', name: 'Tank', threat: 1000, model_version: 'base-1', complete: false }],
    threat_by_target: [{ guid: 'P1', name: 'Tank', target_guid: 'E1', target_name: 'Boss', threat: 600 },
                       { guid: 'P1', name: 'Tank', target_guid: 'E2', target_name: 'Add', threat: 400 }] };
  const scoped = scopeSummary(summary, halfWindow);
  const ratio = scoped.threat[0].threat / 1000;
  expect(scoped.threat_by_target?.[0].threat).toBeCloseTo(600 * ratio);
});
```

In `night.test.ts`:

```ts
it('folds per-target threat by enemy name and shifts taunts onto the night’s clock', () => {
  const withThreat = new Map(summaries);
  const pairs = (threat: number) => [{ guid: 'Player-1', name: 'Tank', target_guid: 'Creature-a', target_name: 'Kaal', threat }];
  withThreat.set(3, { ...summaries.get(3)!, threat_by_target: pairs(100), taunts: [{ at_ms: 1000, source_guid: 'Player-1', source_name: 'Tank', target_guid: 'Creature-a', target_name: 'Kaal', spell_id: 355, spell_name: 'Taunt' }] });
  withThreat.set(4, { ...summaries.get(4)!, threat_by_target: [{ ...pairs(50)[0], target_guid: 'Creature-b' }] });
  const night = nightSummary(fights, withThreat);
  expect(night.threat_by_target).toEqual([{ guid: 'Player-1', name: 'Tank', target_guid: 'Kaal', target_name: 'Kaal', threat: 150 }]);
  expect(night.taunts?.[0].label).toMatch(/pull 1/);
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd web && npx vitest run src/lib/report/window.test.ts src/lib/report/night.test.ts`
Expected: FAIL on the new cases.

- [ ] **Step 3: Implement** the type additions; in `scopeSummary`, after `scoped.threat = scopeThreat(...)`, add `scoped.threat_by_target = scaleThreatPairs(summary.threat_by_target ?? [], summary.threat, scoped.threat)` where the ratio per player is `scopedTotal / wholeTotal` (0 when the whole is 0); `scoped.taunts = (summary.taunts ?? []).filter((taunt) => taunt.at_ms >= window.startMs && taunt.at_ms <= window.endMs)`. In `nightSummary`, fold pairs into a `Map` keyed by `${guid}|${target_name}` (the enemy's GUID differs per pull, so the night key is its name, and `target_guid` becomes the name), and push each taunt with `at_ms + offset` and `label`.

- [ ] **Step 4: Run the tests to verify they pass** (same command).

- [ ] **Step 5: Commit**

```bash
git add web/src/lib/report
git commit -m "feat(web): per-target threat and taunts on the summary, scaled and folded"
```

---

### Task 3: The Threat tab: a target picker and the taunt list

**Files:**
- Modify: `web/src/components/report/ThreatTable.svelte`
- Modify: `web/src/components/report/ReportView.svelte` (pass `pairs={scoped.threat_by_target ?? []}` and `taunts={scoped.taunts ?? []}`; `names={unitNames}`)
- Modify: `web/src/components/report/Glossary.svelte` (term "Threat on a target", "Taunt")
- Test: `web/tests/e2e/report-threat.spec.ts`

**Interfaces:**
- Consumes: `ThreatPair[]`, `Taunt[]` from Task 2.
- Produces: `ThreatTable` props gain `pairs?: ThreatPair[]`, `taunts?: Taunt[]`, `names?: ReadonlyMap<string, string>`.

- [ ] **Step 1: Write the failing e2e test** (the fixture: Warden Kelthas fight 3; the tank taunts once after Task 4's fixture change; read `web/src/fixtures/report/fixture.log` for names)

```ts
// web/tests/e2e/report-threat.spec.ts
import { expect, test } from '@playwright/test';

const FIGHT = '/reports/fixture2abcd?fight=3&tab=threat';

test('threat on a target lists each player’s threat on the picked enemy, and the taunts', async ({ page }) => {
  await page.goto(FIGHT);
  const picker = page.getByTestId('threat-target');
  await expect(picker).toBeVisible();
  await picker.selectOption({ label: /Warden Kelthas/ });
  await expect(page.getByTestId('threat-on-target')).toContainText('Baelgrim');
  await expect(page.getByTestId('threat-on-target').getByTestId('threat-share').first()).toContainText('%');
  await expect(page.getByTestId('threat-taunts')).toContainText('Taunt');
});
```

- [ ] **Step 2: Run the test to verify it fails.** Run: `cd web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-threat.spec.ts`. Expected: FAIL, `threat-target` not found.

- [ ] **Step 3: Implement** in `ThreatTable.svelte`: a `<select data-testid="threat-target">` listing every distinct `target_name` in `pairs` (grouped: name once, with `×N` when several GUIDs share it, folding their threat), defaulting to the enemy with the most total threat; below it a list (`data-testid="threat-on-target"`) of players with their threat on that enemy, a bar against the highest, and `data-testid="threat-share"` per row as the share of all players' threat on that enemy; the existing totals list stays above under a heading "Overall". Then `data-testid="threat-taunts"`: one line per taunt, "1:04 · Hobolol taunted General Kaal (Hand of Reckoning)", with the `label` (pull) over the night. Keep the `~` mark under a window on both lists (pairs are scaled). Glossary: "Threat on a target: the threat one player has built on one enemy, from the damage they did to it and their share of the raid's healing while it was engaged. Whoever has the most is who it attacks." "Taunt: a cast that forces an enemy onto the caster; the Threat tab lists every one, and the timeline marks them."

- [ ] **Step 4: Run the test to verify it passes** (after Task 4's fixture regeneration; run Task 4 first if the fixture has no taunt yet).

- [ ] **Step 5: Lint, type check, commit**

```bash
git add web/src/components/report/ThreatTable.svelte web/src/components/report/ReportView.svelte web/src/components/report/Glossary.svelte web/tests/e2e/report-threat.spec.ts
git commit -m "feat(web): threat on a target, and the taunts"
```

---

### Task 4: Taunt marks on the timeline, and a fixture taunt

**Files:**
- Modify: `web/src/fixtures/report/fixture.log` (add one line after the boss's first melee: `9/26 20:12:05.000  SPELL_CAST_SUCCESS,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,<Warden Kelthas guid>,"Warden Kelthas",0xa48,0x0,355,"Taunt",0x1` — copy the exact guid and flags from neighbouring lines) and regenerate with `node scripts/make-report-fixture.mjs`
- Modify: `web/src/components/report/TimelinesView.svelte` (taunt marks on the player's lane and the boss lane; legend entry; readout names it)
- Modify: `web/src/components/report/ReportView.svelte` (`taunts={scoped.taunts ?? []}` to TimelinesView)
- Test: `web/tests/e2e/report-threat.spec.ts` (one more case)

- [ ] **Step 1: Write the failing e2e case**

```ts
test('the timeline marks the taunt on the tank’s lane and names it on hover', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&view=timelines');
  const mark = page.getByTestId('lane-Player-4184-000000A1').getByTestId('taunt-mark').first();
  await expect(mark).toBeVisible();
  await mark.hover();
  await expect(page.getByTestId('timeline-picked')).toContainText('Taunt');
});
```

- [ ] **Step 2: Run it to verify it fails.** Expected: FAIL, no `taunt-mark`.

- [ ] **Step 3: Implement**: `TimelinesView` gains `taunts: Taunt[] = []`; each lane gets `taunts` (by `source_guid`), drawn as a 3px-wide, full-height `bg-gold` mark with `data-testid="taunt-mark"` and `title="Taunt · 0:05 · on Warden Kelthas"`; `pickNearest` considers taunts within the cast tolerance and sets `picked = { at, name: \`${spell_name} on ${target}\` }`; the legend line gains "taunt". The boss lane draws the same marks (the enemy taunted) in gold beside its casts.

- [ ] **Step 4: Regenerate the fixture, run the threat spec and the tabs spec** (`npx playwright test tests/e2e/report-threat.spec.ts tests/e2e/report-tabs.spec.ts tests/e2e/report-tables.spec.ts`: the fixture changed, so anything pinning fixture numbers must still hold; a taunt adds no damage).

- [ ] **Step 5: Commit**

```bash
git add web/src/fixtures/report web/src/components/report/TimelinesView.svelte web/src/components/report/ReportView.svelte web/tests/e2e/report-threat.spec.ts
git commit -m "feat(web): taunts on the timeline"
```

---

### Task 5: Merge, re-parse, and hand to the tank

- [ ] **Step 1: Merge to main** after `go test ./...` in `logs/` and `api/` and the full web e2e pass.
- [ ] **Step 2: Wait for API CI, re-parse the production sample** (`gcloud run jobs execute parse-report --region us-east1 --args=parse-report,5lop7n5kwysf --wait`).
- [ ] **Step 3: Restage the review harness** (re-parse `scratchpad/logs/wowp.txt`, copy into `scratchpad/report-real`, run `scratchpad/local-report.sh`).
- [ ] **Step 4: Dispatch the tank persona** with the standing brief plus: "The Threat tab has a target picker and a taunt list, and the timeline marks taunts." Fold the findings in.
