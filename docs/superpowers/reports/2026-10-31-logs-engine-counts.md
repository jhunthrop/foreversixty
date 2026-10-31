# Logs engine counts lane — 2026-10-31

Branch `wt/logs-engine-counts`, based on `origin/main` (`4761fe711`). This lane follows up the
finding in `wt/sample-chart`'s report (`docs/superpowers/reports/2026-10-30-sample-chart.md` §5).
Each change was built test-first. Nothing is pushed.

## Commits

| Commit | What |
| --- | --- |
| `07fb1ae9d` | fix(logs): engine 0.6.0 counts each mob death once and records the fight's group |
| `ad54158ea` | feat(api): fights carry mob deaths, group kills, group deaths and the group |
| `afd0ece6f` | feat(api): samplecut -recount refreshes the sample's kill counts from its events |
| `bd6112b18` | feat(web): trash shows its true kill counts again, and none for an old report |

## 1. Kills counted once (engine)

`fight.Counter` (`logs/engine/fight/counts.go`) replaces `npc_kills`. The segmenter and the API's
`engine.Rebuild` both count through it, so a parsed fight and an uploaded fight agree. Before
this, they ran two copies of the same double-counting switch.

| Field | Meaning |
| --- | --- |
| `mob_deaths` | Each hostile non-player unit that died in the fight, once per GUID. UNIT_DIED and PARTY_KILL for the same unit are one death. UNIT_DESTROYED and UNIT_DISSIPATES (totems, summons) don't count, and neither do neutral critters. |
| `group_kills` | How many of those deaths the group caused: a PARTY_KILL names the unit, or the last damage the unit took came from a unit flagged MINE, PARTY or RAID. The flag test covers the group's pets. This is always at most `mob_deaths`. |
| `group` | Player GUIDs flagged MINE, PARTY or RAID on any line of this fight. Pets are left out. |
| `group_deaths` | Player deaths among the group. `deaths` still counts every player death in range. |

`npc_kills` is gone from `fight.Fight` and from `store.FightEntry`. On the report.json line, the
four new fields are pointers with `omitempty`, so "absent" (an old line) can never be read as
zero. `session.Version` goes from 0.5.3 to 0.6.0, so fight-file URLs re-key.

**Where the group lives.** It is on the fight entry: in report.json's `fights[]`, in the API meta
`fights[]`, and in the sample `meta.json`. summary.json is unchanged. Adding `group` to
`summary.Summary` would break the sim adapter's rule that summary lists are never null, and the
client already has the fight entry before it loads any summary. So the fight entry is the one
place the group is stored.

**Tests**
- `fight/counts_test.go`, 12 cases:
  - a group kill logged twice
  - an outsider's kill
  - the group's killing blow with no PARTY_KILL
  - an outsider's killing blow
  - critters and totems
  - a duplicate UNIT_DIED
  - group deaths
  - membership changing between two pulls (the mage leaves the party)
  - pets left out of the group
  - a pet's killing blow
  - the Open snapshot
  - Counter on a hand-built fight
- `session/counts_test.go`: fixture tests over the full pipeline. The oracle is read straight from
  the log text: for each fight, the distinct hostile units that a UNIT_DIED or PARTY_KILL line
  names. No count is pinned. The v16 Hollow Sentinel (UNIT_DIED plus PARTY_KILL) is one death and
  one group kill. In every fight, the group is a subset of the players and group deaths are at
  most deaths.
- The `v16` and `v22` summary goldens are unchanged.

## 2–3. Per-fight group membership and group deaths (API)

- Migration `0059_fight_counts` adds four nullable columns to `fights`: `mob_deaths`,
  `group_kills`, `group_deaths` and `group_players text[]`.
- The parse job, the ingest and `Store.Fights` all write and read them.
- For a fight row from before 0.6.0, all four are null, and `GET /v1/reports/{id}` leaves the
  fields out.
- **Migration number:** 0059 stays. The coordinator renumbers the held sign-in lane's own 0059
  when that lane merges after launch.
- `npc_kills` is no longer written or read. The column stays, default 0, so that a revision still
  running the old code during the rollout can still insert. **Follow-up:** a later migration drops
  it.
- Updated to match: `openapi.yaml`, `docs/contracts/app-v1.md` with the `reports.get` fixture, and
  the generated iOS and Android models. The apps never read `npcKills`.

**Tests**
- `TestStoreReadsBackFightsPlayersAndRanges` round-trips all four fields.
- `TestStoreLeavesOutTheCountsOfAFightFromBeforeThem` nulls the columns and checks that the JSON
  carries none of `mob_deaths`, `group_kills`, `group_deaths`, `group` or `npc_kills`.
- `TestRebuildCountsEachMobDeathOnceAndTheGroupsKills` covers the ingest path.
- The app contract test passes on the new fixture.

## 4. Consumers

### On main (done here)

- **`lib/report/kill-counts.ts`**:
  - `killCounts(fight)` returns null unless `mob_deaths` and `group_kills` are both finite numbers.
  - `sumKillCounts(run)` returns null if any pull in the run has no figure. It never gives a
    partial total.
  - `killCountsLabel` gives "177 mobs died · your group killed 26", or "No mobs died" when the
    count is zero.
- **`outcomeLabel`** gives that label for trash, or `''` for an old trash line.
- **`FightSelector`**:
  - On a trash row, the count is on the second line (`fight-N-kills`). It is too long for the
    outcome slot beside the title in a 300 px column.
  - The outcome slot now shows only for bosses and live pulls.
  - A folded zone row shows the summed figure (`fight-group-K-kills`).
  - Both meta lines wrap rather than truncate a true figure.
- **The phone fight summary and Compare's fight list** join their parts through `joinFacts`, so an
  empty outcome leaves no stray " · ".
- **`FightEntry` type**: `npc_kills` removed; `mob_deaths?`, `group_kills?`, `group_deaths?` and
  `group?` added.
- **The web report fixture** is regenerated with `make:report-fixture`. It was also stale from
  earlier engine versions, so it picks up item levels, seeded auras and the fight header fields.
  `meta.json` mirrors the counts.
- **e2e** (`report-players.spec.ts`):
  - true counts on a zone-less trash row and on a folded zone row;
  - a new test for an old-format report: no `fight-1-kills`, and no "230", "kills" or "mobs died"
    anywhere in the fight list.

### What `wt/sample-chart` should switch to after merge

Its `lib/report/group.ts` builds the group from report.json's last-seen unit flags, and from the
`group` that samplecut `groupOf` writes into meta. It should:

1. **Scope.** For the pull on show, use `fight.group` (`FightEntry.group`, present in both meta
   `fights[]` and report.json `fights[]`) when the fight carries the new counts. The test is
   `killCounts(fight) !== null`, or `fight.mob_deaths !== undefined`. Use it for a group of any
   size, including instance logs: per-fight flags have no "left before the log ended" problem, so
   the rule that outsiders must outnumber the group first is only needed in the fallback. Fall back
   to today's report-wide derivation only when the fight has no counts (a pre-0.6.0 report) or
   `group` is absent (nobody was flagged).
2. **Meta `group`.** samplecut's top-level `group` becomes redundant for any report with per-fight
   groups. Keep it only as part of the fallback.
3. **Deaths under the group scope.** Use `fight.group_deaths` for the subtitle and the fight-list
   death marks when present. Otherwise count deaths from the summary under the derived group, as
   today.
4. **Kill counts.** That lane removed them. Bring them back through `killCounts` and
   `killCountsLabel` (`fight-N-kills`, `fight-group-K-kills`) as done here, and drop its
   "no kill counts" e2e assertions.
5. **Expected merge conflicts:**
   - `format.ts` (`outcomeLabel`)
   - `FightSelector.svelte`
   - `fights.ts` (trash-group `kills`)
   - `format.test.ts` and `fights.test.ts`
   - `report-players.spec.ts`
   - `api/cmd/samplecut/main.go` (flag block versus `groupOf`)
   - `web/public/sample-report/sample/meta.json`

   For meta.json, take that lane's file and then rerun `go run ./cmd/samplecut -recount
   ../web/public/sample-report/sample`. Recount keeps every other meta key, `group` included.

## 5. Old reports: reprocessed or not?

**They are not reprocessed.**
- There is no reprocess or backfill job. `store.Decompress`'s comment calls a "Phase 3 reprocess
  job" planned, and it was never built.
- The whole-file parse runs once per upload (`reports/uploads.go` → `parse-report`). It can be
  rerun by hand for one report with `api parse-report <report_id>`, which upserts the fight rows
  and rewrites report.json, as long as the upload object still exists.
- A live companion report can't be re-parsed whole. Each fight is fixed when it is ingested.

So existing reports keep their old fight rows and old report.json. The UI reads both formats
safely:
- An old line has no `mob_deaths` or `group_kills`, so it shows no kill count. The API answers
  such rows without the fields, never as zero.
- A folded zone row with any old pull shows no count.
- Group scoping falls back to the report-wide flags, which is what the sample-chart lane does now.

## 6. The committed sample report

The original log isn't kept (it isn't in the WoW Logs folder either), so `samplecut -log` can't
re-cut it. Each fight's committed `events.parquet` holds every event the segmenter assigned to
that fight, so `fight.Counter` over it gives the same figures the engine writes.

`samplecut -recount <dir>` (`recount.go`, with a Go test):
- rewrites only the fight lines' counts in `report.json` and `meta.json`, and drops `npc_kills`;
- keeps every other meta key;
- leaves the summaries and `engine_version` (0.5.3) as they were cut.

Results on the committed sample:
- Fight 6 reads 177 mob deaths and 26 group kills. The old figure was 203, and this matches the
  sample-chart lane's hand count exactly.
- Fight 6's group is **five** players for that pull. The last-seen flags gave four.
- Across the night: 675 mob deaths, 78 group kills.

To re-cut from a log later: `go run ./cmd/samplecut -log <log> -out ../web/public/sample-report/sample`.
That writes 0.6.0 output directly.

## Gates

All checks were exit-code checked, with no pipes into pushes.

| Gate | Result |
| --- | --- |
| `go vet ./...` (logs, api, companion build) | clean |
| `go test -race ./...` (logs) | all pass |
| `go test -race -p 1` api: reports, parse, engine, coach, db, cmd/api, cmd/coacheval, cmd/samplecut, cmd/seedguild (private db `forever_gc_logs`, dropped afterwards) | all pass |
| `go test ./adapter/` (sim) | pass. Other sim packages need `make simdb` data and are unrelated |
| vitest: lib/report, lib/coach, fixtures, lib/mobile, components/report | 78 files, 785 tests, pass |
| `astro check` | 0 errors |
| `npm run lint`, `prettier --check .` | pass |
| Fixture build (`npm run build`, includes the island-size check) | pass |
| Playwright on 4705, both projects: every `report-*.spec.ts`, sample-report, coach-report, logs-recent-reports | 376 passed, 54 skipped (project-specific), 0 failed |
| Playwright: launch-polish, onboarding-second-screen, first-impression, handoffs-report | 82 passed |
| After the "No mobs died" copy: rebuild, then report-players, sample-report, report-compare, report-phone | 54 passed, 10 skipped |

## Concerns

- **`node_modules` link.** The brief's `charts` worktree `node_modules` predates main's
  `@floating-ui/dom`, so `astro check` failed in `popover.ts`. I linked
  `.claude/worktrees/dropdown/web/node_modules` instead: its lockfile is byte-identical to this
  branch's. The link was removed before each commit.
- **`npc_kills` column.** It is still in the database and needs a drop migration once 0.6.0 is
  deployed everywhere.
- **The sample is mixed-version.** Its fight lines carry 0.6.0 counts while its summaries and
  `engine_version` are 0.5.3. That is honest, but a fresh cut from the log would be cleaner.
- **Kill attribution by last hit.** If an outsider lands the last hit after the group did the work,
  the death counts as not the group's, unless the game also logged a PARTY_KILL. That is the
  intended rule. PARTY_KILL is the game's own verdict, and it wins whenever it is present.
- **Visual review.** The trash row is now a two-line row with the count on the meta line. On the
  folded zone row, "79 pulls ‹" wraps to its own line. I checked this on screenshots at 1440 and
  390, but the ux-designer and wow-player reviewers have not reviewed it.

## Round 2: merged with main (sample-chart lane), per-fight scope, captures

### Merge

- `origin/main` was merged into this branch (a merge, not a rebase). The commit is `e7fcd7154`.
- Conflicts were in `format.ts`, `FightSelector.svelte`, `fights.ts`, `CompareMode.svelte`, their
  tests and `report-players.spec.ts`. Each was resolved on this branch's side:
  - the true kill line stays;
  - `joinFacts` stays the one way to join a row's facts;
  - the folded row keeps its summed counts.
- From main's side I kept the "still being logged" live title and an extra `Live` test.
- The sample's `meta.json` and `report.json` are main's files, recounted with `samplecut -recount`.
  Main's top-level `group` key is kept.

### The client reads each pull's group

**`lib/report/group.ts` `fightGroupScope(fights, raidFlagged)`** builds the scope from the
fight's own `group` and `players`. For the night view it uses the union over the boss pulls. It
returns null when any of those pulls predates engine 0.6.0, and the old report-wide `groupScope`
then runs exactly as before.

| Pull | Scope |
| --- | --- |
| Anyone outside the group in it | Scoped to the group, in an instance pull too. The "outsiders outnumber" rule is only needed by the fallback. |
| Only the group in it | Every player, named Group or Raid. |
| No group in it | Everyone in range. |

- **Noun:** Raid when the group is larger than 5, or when any member's report units carry the
  RAID flag.
- **Source:** `ReportView` reads report.json's own fight lines once they arrive, and the meta's
  before that.
- **`fightDeathMark`:**
  - A fight row's dagger (and a folded row's) is the group's deaths. Its title reads "1 death in
    your group · 3 among everyone in range". That makes the dagger match the row's "your group
    killed" line.
  - A pull the group was not in marks everyone in range's deaths.
  - An old fight keeps its in-range count, worded as before.
- **Subtitle:** reads `group_deaths` under the group scope.
- Main's honest wording is unchanged: "players seen", "Your group / Everyone in range", and
  "Raid / Group / Total" on the chart.

**What this changes on the sample's fight 6:**
- The default group is **five** players: the fight's own flags put a party member in it that the
  last-seen flags called an outsider. The roster reads "5 players · 0 deaths".
- The row reads "177 mobs died · your group killed 26" with no dagger. The 2 deaths were
  outsiders', and they show under Everyone in range.
- Main's e2e test asserted the in-range dagger on fight 6. It now asserts the kill line and no
  dagger.

**Test changes:**
- `group.test.ts` adds 14 tests for `fightGroupScope` and `fightDeathMark`.
- The `raid40.ts` mock now carries the pull's own group, all forty players. They were all
  MINE-flagged already.
- FightRow's `deaths` field was only feeding the old dagger, so it is gone.

**For reviewers:** trash rows with outsider-only deaths no longer show a dagger. Whether that is
right is a design call.

### Captures

In `design/reviews/2026-10-31-logs-counts/`:

| File | Shows |
| --- | --- |
| `sample-fights-1440.png`, `sample-fights-390.png` | The sample's fight list, trash rows reading "N mobs died · your group killed M" |
| `sample-page-1440.png`, `sample-page-390.png` | The whole sample page on fight 6, under the group scope |
| `folded-zone-row-1440.png` | The fixture report's folded "Trash · Sanguine Depths" row, with its summed count |
| `no-mobs-died-1440.png` | The same group opened: "No mobs died" on the 6.4 s pull |
| `old-format-no-count-1440.png` | The sample served with its counts stripped (an engine 0.5 report): no kill line on any row, the in-range daggers as before |

### Gates (round 2)

All checks were exit-code checked.

| Gate | Result |
| --- | --- |
| `go vet ./...` on logs and api | clean |
| `go test -race ./...` on logs | pass |
| `go test -race -p 1` on api reports, parse, engine, coach, db and cmd/... (private db `forever_gc_logs`, dropped afterwards) | pass |
| vitest: lib/report, components/report, lib/coach, fixtures, lib/mobile | 79 files, 809 tests, pass |
| `astro check` | 0 errors |
| lint, format | pass |
| Playwright on 4705, desktop and mobile: every `report-*.spec.ts`, plus sample-report, coach-report, onboarding-second-screen and launch-polish | 442 passed, 56 skipped (project-specific), 0 failed |

## Round 3: the code review's FIX FIRST

The review is `docs/superpowers/reports/2026-10-31-logs-engine-counts-review.md`, committed with
both design verdicts.

### H1. Neutral mobs (fixed)

**What counts as a mob** (`fight.Counter`): a non-player unit that either
- was hostile on any line of the fight, or
- is neutral and the group damaged it or PARTY_KILLed it.

**How the counts follow:**
- A PARTY_KILL on a neutral NPC counts toward both `mob_deaths` and `group_kills`.
- Group damage on neutral NPCs is tracked (`groupHit`), and so is the last hit
  (`lastHitByGroup`). A neutral unit the group never touched (someone else's pull, or a critter
  in someone else's AoE) is not counted.

**Critters.** This repository's data has no creature types: `data/builds` carries no creature
table, and the classic-db `creature_template` exists only in the audit pipeline's uncommitted
cache. So the line is drawn by health:
- A neutral unit whose known maximum health is at most `fight.Options.CritterMaxHP` is a critter
  and never a mob. The default is 10, in `DefaultOptions`.
- The weakest level-one neutral beasts in the Forever sample have 42 HP.
- When no advanced block ever showed the unit's health, it counts.

**Segmentation fix.** The neutral-beast test showed that a pull of only neutral beasts **never
opened a fight at all**: `hostileCombat` required one side to be hostile. It now uses
`units.Enemy`, which the summary already used for damage and threat. Two neutral units, or two
friendly ones, are still no fight. The `v16`, `v22` and Forever summary goldens are unchanged, and
so is the regenerated web report fixture.

**Tests:**
- The old critter case is now separate cases: a neutral beast the group killed (counts, with and
  without a PARTY_KILL), a neutral unit the group never hit (does not count), a critter the group
  PARTY_KILLed (does not count), and the configurable cutoff.
- A segmenter test checks that a neutral-beast pull opens a fight.
- The session oracle now runs over the v16, v22 and Forever fixtures. It reads mob deaths from the
  text by the same rule, using a CSV reader so names with commas parse.

### M1. Feign Death (handled; not confirmed on a real log)

- A player UNIT_DIED counts, but stays pending.
- If the same GUID then casts (SPELL_CAST_START or SPELL_CAST_SUCCESS) or swings (SWING_*)
  before any SPELL_RESURRECT names them, the death is taken back from `deaths` and, when it
  counted there, from `group_deaths`.
- Ticks from a DoT and hits from a pet carry on after their owner dies, so they are not signs of
  life.
- The synthetic test covers a feign undone by the hunter's swing, and a real death that a DoT
  tick does not undo, followed by a resurrect and a cast, which keeps it.
- **No real Forever hunter log was available to confirm how the client logs Feign Death.**
- **Known gap:** a shaman's Reincarnation, if the client logs it as a cast with no
  SPELL_RESURRECT, would undo a real death.

### M2. Charmed mobs (fixed)

An NPC seen hostile anywhere in the fight counts when it dies, even when it dies carrying a group
member's charm flags. Covered by a test.

### LOWs

| Finding | Action |
| --- | --- |
| L1, L2 (rollout window) | No change. Both need an old revision to serve traffic during the deploy, and the web shows no count rather than a wrong one. |
| L3 (no backfill for old fights) | Follow-up: a production recount over each fight's stored `events.parquet`, the way `samplecut -recount` works. Not built here. |
| L4 (no migration test) | Added `api/internal/db/fight_counts_migration_test.go`. Up adds the four columns as nullable, and down removes them. |
| L5 (subtitle and death list disagree) | Under the group scope the subtitle reads the summary's death list again, the same list as the Deaths tab, so the two always agree. The fight list's mark still reads `group_deaths`, which also takes back a feign. Making the summary's death list feign-aware is a follow-up. |

### Sample

`samplecut -recount` was rerun:
- **Fight 6 is unchanged: 177 mobs died, your group killed 26, 0 group deaths, a group of five.**
- Fight 20 is now 5 mob deaths and 2 group kills (was 3 and 0).
- Fight 21 is now 2 mob deaths and 1 group kill (was 1 and 0). The new kills in fights 20 and 21
  are the Galestriders.
- Night totals: 678 mob deaths, 81 group kills.
- **Caveat:** recount keeps the sample's 0.5 fight boundaries. A re-cut from the log could add
  pulls of only neutral beasts that the old segmenter never made into fights.
- The captures were re-taken. `sample-neutral-kills-1440.png` is new and shows fights 20 and 21.

### Gates (round 3)

All checks were exit-code checked.

| Gate | Result |
| --- | --- |
| `go vet ./...` on logs and api | clean |
| `go test -race ./...` on logs | pass |
| `go test -race -p 1` on api reports, parse, engine, coach, db (with the new 0059 test) and cmd/... (private db, dropped afterwards) | pass |
| sim adapter tests; companion build | pass |
| vitest: lib/report, components/report, lib/coach, fixtures, lib/mobile | 79 files, 809 tests, pass |
| `astro check` | 0 errors |
| lint, format | pass |
| Playwright on 4705, desktop and mobile: every `report-*.spec.ts`, plus sample-report, coach-report, onboarding-second-screen and launch-polish | 442 passed, 56 skipped, 0 failed |

The engine version stays 0.6.0, because 0.6.0 has not shipped.

## Round 4: the re-review's R1 and R2

### R1. The Feign Death undo now needs a Feign Death (fixed)

A player death is taken back only when **that same player cast Feign Death within
`Options.FeignDeathWindow` (default 1 s) before the UNIT_DIED**, and then casts or swings again.
Every other death stands, whatever follows it: a corpse run, Reincarnation and a Soulstone all
log the dead player acting again with no SPELL_RESURRECT.

- **Spell ids.** `Options.FeignDeathSpells` defaults to `[5384]`. The active build's spell data
  (`data/builds/1.60.1.70630/spells.json`) has 18 names containing "Feign Death". Only 5384 is
  the hunter's spell. The rest are Improved Feign Death talents (19286, 19287, 24432) and
  creature "Permanent Feign Death" auras.
- **Removed:** the `Resurrect` case, which only the old undo needed.
- **Tests:**
  - The synthetic feign (a cast of 5384, the death 0.3 s later, a swing) is still undone.
  - Kaal sequence: four players die, nobody is resurrected, each casts and swings two minutes
    later. Every death stays, and so does every group death.
  - Reincarnation sequence: the death, the dead shaman's cast of 21169 3.5 s later, then a
    Lightning Bolt. The death stays.
  - A Feign Death cast 3 s before a death doesn't make it a feign.
- **General Kaal wipe on the cached retail log** (`logs/testdata/cache/wowp-retail-v16.txt`,
  gitignored; `forever-logs fights -json`):

  | Commit | Kaal wipe deaths / group deaths | All fights' deaths |
  | --- | --- | --- |
  | Before round 3 (`4e221dcc6`) | 9 / 9 | 65 |
  | Round 3 (`1c26af802`) | 0 / 0 | 44 |
  | Now | 9 / 9 | 65 |

- No real Forever hunter log was available to confirm whether the client logs Feign Death as
  UNIT_DIED at all.

### R2. Boundary fixture (added)

`logs/engine/event/testdata/v16-levelling.log` is hand-made from the v16 fixture's own line
shapes. It contains:
- a pull of two neutral Plainstriders, which the group kills (one with a PARTY_KILL, one by last
  hit);
- an ambient fight between a neutral Plainstrider and a hostile Gluttonous Tick, far from the
  group (the retail log's ambient case);
- a General Kaal wipe: the hunter feigns, two players die, nobody is resurrected, and both run
  back and cast again before ENCOUNTER_END.

`TestALevellingNightSplitsAndCountsByTheRules` pins structure only, never counts:
- there are no parse errors;
- the neutral pull is a fight of its own, with group kills of at least 1 and no more than its mob
  deaths;
- no fight opens on the ambient NPCs;
- the wipe is one encounter and not a kill;
- its deaths and group deaths equal a text oracle: every player UNIT_DIED except one within 1 s
  of that player's own Feign Death cast.

The fixture is also in the mob-death oracle test. Under the old hostile-only rule, the neutral
pull makes no fight and the ambient tick fight does, so both checks fail. Under round 3's undo,
the wipe reads 0 deaths.

R2's other point (ambient friendly-against-neutral NPCs keeping a fight open) is not a blocker. It
is noted as a follow-up: require a group-affiliated unit on one side of `hostileCombat`.

### Sample

`samplecut -recount` was rerun and produced no change. Fight 6 is still 177 mob deaths, 26 group
kills and 0 group deaths. The sample has no Feign Death cast.

### Gates (round 4)

All checks were exit-code checked.

| Gate | Result |
| --- | --- |
| `go vet ./...` on logs and api | clean |
| `go test -race ./...` on logs | pass |
| `go test -race -p 1` on api reports, parse, engine, coach, db and cmd/... (private db, dropped afterwards) | pass |
| sim adapter (`-race`); companion build | pass |
| vitest: lib/report, components/report, lib/coach, fixtures, lib/mobile | 79 files, 809 tests, pass |
| `astro check` | 0 errors |
| lint, format | pass |
| Playwright on 4705, desktop and mobile: every `report-*.spec.ts`, plus sample-report, coach-report, onboarding-second-screen and launch-polish | 442 passed, 56 skipped, 0 failed |
