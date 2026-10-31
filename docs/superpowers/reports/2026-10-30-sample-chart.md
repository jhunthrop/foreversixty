# Sample chart lane — 2026-10-30

Branch `wt/sample-chart`, based on `wt/dropdown`. The owner asked whether the sample report's
charts were "still too small" and then said "do the chart changes". Three changes, each built
test-first and committed on its own. Nothing is pushed.

## 1. The full chart on desktop (`6dab5f14c`)

- From md (768 px) up, the sample report draws the damage chart the way a real report does. That
  means the head with its title and units, the legend, the Players and Window presets, the log
  switch, the brush, and a plot of 200 px (220 px from 1440).
- The 74 px `compact` strip now shows only on phones. `ReportView` reads it from a `MediaQuery`
  for `min-width: 768px`. The island is client-mounted, so the first paint already has the right
  value.
- The phone promise is unchanged and green: `onboarding-second-screen.spec.ts` still checks that
  "the damage rows start on the first screen". On a phone the chart opens collapsed by default
  (`chartOpen`). Once opened, the sample still draws the strip.
- **Desktop pins I checked:**
  - `launch-polish.spec.ts` › "the sample report says why its numbers are small…" assumed the
    strip. It pinned the chart at 100 px or less, at least two rows above a 720 px fold, and the
    Window field hidden. The owner overrode that earlier ruling. The test now pins the full chart:
    not compact, the Window field visible, and the table starting below the chart.
  - Measured at 1280 × 720:

    | Element | Top | Bottom |
    | --- | --- | --- |
    | Chart | 390 px | 885 px |
    | Plot | 596 px (222 px tall) | |
    | First row | 986 px | |

    At 1440 × 900 the chart runs from 390 px to 905 px and the first row starts at 1006 px.
  - `coach-report.spec.ts` › "at 1280 x 800 Fix first ends above the fold" is about the Coach tab
    of a real fixture report, which has no chart. It is unaffected and green.
  - `sample.astro`'s reserves (`md:min-h-[1556px] lg:min-h-[1325px]`) were measured on a real
    report that draws the full chart. The sample now matches that layout, so they stand.
- **New tests in `sample-report.spec.ts`:**
  - Desktop, at 1280, 1440 and 2000: the full chart, a plot of at least 222 px including its axis,
    and zero layout shift from anything inside `#report`.
  - Phone: the compact strip once the chart is opened.
- **CLS** inside the report is 0 at 1280, 1440 and 2000, and 0 overall at 390. The only shift on
  the desktop page is in the site header, which is outside the report (see Concerns).

## 2. Finer time buckets on short fights (`8979a853d`)

**Where the buckets come from.**
- The logs engine writes one value per second: `Options.Bucket` defaults to `time.Second` in
  `logs/engine/summary/summary.go`.
- The "5 s" was the client's own trailing average (`trailingAverage` in `lib/chart/drawn.ts`),
  drawn at one point per second.
- So the whole change is client-side, and **the engine needs nothing** for this. Only points finer
  than 1 s would need an engine change, which would be a smaller `Options.Bucket` and a larger
  summary. Nothing asks for that.

**The rule** (`lib/chart/buckets.ts`, `chartBucket`):
1. Take the widest of 1, 2 and 5 s that still gives at least 30 points. If none does, use 1 s.
2. Widen through 10, 15, 30 and 60 s until the fight fits in 300 points (`MAX_POINTS`).
3. The axis label floor is `max(2 s, bucket)`, so labels never step finer than the points.

| Fight | Bucket | Points | Label floor |
| --- | --- | --- | --- |
| 5 s | 1 s | 5 | 2 s |
| 10 s | 1 s | 10 | 2 s |
| 60 s | 2 s | 30 | 2 s |
| 4 min | 5 s | 48 | 5 s (today's average) |
| 15 min | 5 s | 180 | 5 s |
| 30 min | 10 s | 180 | 10 s |
| 3 h | 60 s | 180 | 60 s |

**Implementation.**
- Lines are bucket means, each value being that bucket's mean per second. The trailing average is
  deleted along with its tests.
- Each `ChartSeries` carries `stepMs`. The plot's x values, its empty check, the dead player's end
  dot, the tooltip and the legend's scrub all read through it (`valueAt`). Threat and Resources set
  no step, so they keep one point per second.
- The shared axis (`time-axis.ts`) gains a 2 s step behind a floor parameter. The floor defaults to
  5 s, so every other caller (RaidCooldowns, TimelinesView, Threat) labels exactly as before.
- The tooltip head names the bucket: "DPS, each second" or "DPS, 5 s average".

**Tests.** Unit tests cover the bucket choice at 5 s, 10 s, 60 s, 4 min and 15 min, plus 30 min,
1 h and 3 h for the 300-point bound. They also cover the bucket means, `valueAt`, the label floor,
`stopAtDeath` on 5 s points, the rate model at 340 s and at 10 s, the tooltip on 5 s points, and
the copy.

## 3. The sample opens on a boss fight (`fb8bbdb52`)

- The default is set in `ReportView` (`firstFight`). It used `biggestFightIndex` for the sample.
- That is now `sampleFightIndex`: the first boss kill, else the longest boss attempt, else the
  longest pull.
- A real report keeps `defaultFightIndex` (the first boss pull), and a `?fight=` in the url still
  wins.
- **Finding.** The bundled sample (`web/public/sample-report/sample/meta.json`) has 79 fights and
  every one is trash. It already opened on fight 6 (12:43, its longest pull), not on a 10 s pull,
  and it still does. The 10 s pull the owner saw must have been picked by hand or reached through a
  link. A boss only shows first once the sample is replaced with a log that has one.

## Gates

All exit-code checked, with no pipes.

| Gate | Result |
| --- | --- |
| Full vitest | 686 files, 7741 tests, pass |
| `npm run check` (fixture data) | 0 errors |
| `npm run lint` | pass |
| `npm run format:check` | pass |
| `check-island-size.mjs` | pass: `report-island.js` 394,950 B raw / 126,635 B gzip; chart chunk 79,297 B raw / 31,915 B gzip; no budget moved |
| Targeted Playwright, desktop and mobile (sample-report, launch-polish, onboarding-second-screen, coach-report) | 108 passed, 3 skipped (project-specific), 0 failed |
| Full Playwright suite, both projects | 3383 passed, 410 skipped, 1 failed. The failure was `[mobile] character-rating.spec.ts:25` (the rating panel not found within 5 s), a character page this lane does not touch. Rerun alone with `--repeat-each=3`: 18 passed. Counted as load flake |

## Captures

In `design/reviews/2026-10-30-controls-charts/`:
- `sample-chart-1440.png`
- `sample-chart-2000.png`
- `sample-chart-390.png` (the phone, with the chart opened to the strip)
- `sample-chart-1440-10s-pull.png` (fight 2, the 9.9 s pull: one point a second, a label every
  2 s)

## Concerns

- **One-second points are spiky.** Low-level trash at one point a second draws jagged lines; that
  is the point of the change on a 10 s pull. From 30 s to about 2.5 min (2 s points) the lines are
  noisier than the old 5 s average was.
- **Long fights changed shape slightly.** A 5 s bucket now draws one point per 5 s instead of a
  5 s trailing average at every second. A 15-minute fight drops from 900 points to 180.
- **The desktop fold.** At 1280 × 720 and 1440 × 900, no damage row is on the first screen of the
  sample any more. That is the owner's call, and the phone still keeps its promise.
- **Header CLS, not changed in this lane.** The site header logs a font-swap layout shift of about
  0.0004 (a 3 px text move in the nav at about 40 ms). It is outside the report.

## Fix round (ux-designer and wow-player HOLDs)

### 4. The brush's end handle 

- **Cause.** The handles are range inputs with `step=1000` and `max=domainMs`. When a fight does
  not end on a whole second, the browser snaps the value to the last whole step. On the 9.879 s
  pull the browser reported the end input as 9000, so the handle drew at about 89% while Window
  read "Whole fight". The cause was the input's step. Neither `brushShape` nor the buckets were
  involved.
- **Fix.**
  - The inputs take `step="any"`.
  - A dragged edge snaps to the nearest whole second, except that the fight's end lands on its
    exact millisecond (`lib/chart/brush.ts`).
  - The end and domain labels read the fight's length as the fight list does: `0:09.9` under a
    minute, the rounded second past it (`endClockLabel`).
- **Tests.**
  - Unit tests cover `brushSnap` and `endClockLabel`.
  - A server-render test of OverviewBrush covers fights on 1 s, 2 s, 5 s and 10 s points:
    9.879 s, 80.261 s, 763.14 s and 30:00.437.
  - An e2e test runs on sample fights 2, 9 and 6, on both desktop and mobile, at 1024 px so the
    mobile project also draws the brush. It checks that the browser's own input value equals
    `max`. It was red before the fix: 9000 for 9879, and 80000 for 80261.

### 5. Counts that say what they count

**The engine and group membership.**
- The engine reads the affiliation bits (MINE `0x1`, PARTY `0x2`, RAID `0x4`, OUTSIDER `0x8`;
  `logs/engine/units/units.go`).
- `report.json`'s units keep each unit's **last-seen** flags (`Registry.touch`). The per-fight
  summary keeps no flags at all.
- In the sample, the units are 1 MINE, 4 PARTY and 955 OUTSIDER players. The logger was in a
  party of five.
- So the group can be found client-side, from report-wide, last-seen flags.

**What the client now does** (`web/src/lib/report/group.ts`):
- **The default "friendlies" scope is the logger's group** when outsiders outnumber the group (an
  open-world log) and the pull has a group member in it. The Source list then offers "Your group"
  (or "Your raid") and "Everyone in range" (`?source=everyone`).
- **Instance logs keep every player.** There, a few players read as outsiders only because they
  left the group before the log ended. Scoping those out would drop real raiders, so outsiders
  have to outnumber the group first.
- **A crowded pull the group was not in** (the 9.9 s pull) keeps every player, and the Source list
  says "Everyone in range". The page doesn't go empty.
- **The sample's meta carries `group`.** That lets its summary-first table scope before
  `report.json` arrives. `api/cmd/samplecut` now writes the field (`groupOf`, with a Go test), and
  the committed `meta.json` got the same derivation from its `report.json` units: five GUIDs added
  and nothing else changed.

**Copy fixes:**

| Where | Change |
| --- | --- |
| Chart head | Names the set it sums: "Raid", "Group" (a group of five or fewer with no RAID flag), or "Total" for everyone in range |
| Median line | "Raid median", "Group median" or "Median player" |
| Players field (everyone in range) | "Top 4, bottom 4 of 90 players seen", in a field widened to 260 px so it never truncates |
| Players card | "265 players seen" under everyone in range; "4 players · 0 deaths" under the group |
| Subtitle deaths | The group's deaths under the group scope. Fight 6's "2 deaths" are outsiders' and show only under everyone in range |
| Fight list death marks | Tooltip reads "2 deaths among everyone in range" when the default is the group |
| Chart markers | Under the group scope, the chart no longer marks outsiders' deaths and cooldowns |
| Trash kill counts | Removed everywhere: fight rows, trash group rows, the phone fight summary, Compare's fight list |

**Why the kill counts went.** The engine's `npc_kills` (`logs/engine/fight/fight.go`, `assign`)
adds one for every hostile UNIT_DIED **and** one for every PARTY_KILL. Each kill the party makes
is therefore counted twice. Every mob anyone in range killed counts as well. Measured on the
sample's own `events.parquet`:

| Fight | Hostile deaths | PARTY_KILL | Also counted as UNIT_DIED | Engine shows |
| --- | --- | --- | --- | --- |
| 6 | 177 | 26 | all 26 | 203 |

Fight 6's true counts are 177 mobs died in range, and 26 were the party's. No client-side figure
can be made true, so the count is not shown. Critters are not counted unless the party kills one:
neutral units are not hostile, but PARTY_KILL counts any party kill.

**What the engine needs** (not changed in this lane):
1. **Kills.** In `logs/engine/fight/fight.go` `Segmenter.assign`, count each hostile NPC GUID's
   death once: de-duplicate UNIT_DIED and PARTY_KILL by dest GUID. Add a separate count for the
   group's kills (PARTY_KILL, or a death whose last hostile-damage source carried MINE, PARTY or
   RAID). Then the fight entry can say "26 kills by your group · 177 died in range".
2. **Per-fight group membership.** Today the flags are last-seen and report-wide. The engine
   should OR the affiliation bits each player carried during the fight into the summary's roster
   row, as e.g. `affiliation`, or a `group` list per `FightEntry`. Then someone who joins or
   leaves the party mid-night is scoped right per pull, and instance logs can scope too. The
   client would read that instead of the units' flags.
3. **Per-fight group deaths.** `FightEntry.deaths` counts every player death in range. A
   `group_deaths` field would let the fight list show the group's own deaths.

**Tests.**
- `group.test.ts`: 9 tests.
- Copy, rate and format/fights updates.
- samplecut's `groupOf` Go tests.
- New e2e in `sample-report.spec.ts`:
  - the group default: "Your group", "4 players · 0 deaths", "Group … DPS", "Group median", no
    "Raid", no death markers, no kill counts, and the in-range death tooltip;
  - Everyone in range: "265 players seen", "Total … DPS", "of N players seen" not truncated,
    "Median player", "2 deaths";
  - the 9.9 s pull reading "Everyone in range".
- `report-players.spec.ts` no longer expects "115 kills" or "230 kills".

### Gates (fix round)

| Gate | Result |
| --- | --- |
| Full vitest | 689 files, 7764 tests, pass |
| `check` (fixture data) | 0 errors |
| `lint`, `format:check` | pass |
| Island size | pass: `report-island.js` 396,864 B raw / 127,439 B gzip, +1.9 KB raw over the first round; no budget moved |
| `go test ./cmd/samplecut/` | pass |
| Full Playwright suite on the count-fix build | 3396 passed, 410 skipped, 0 failed |
| After the last two changes (marker scope, field width), desktop and mobile: sample-report, launch-polish, onboarding-second-screen, coach-report and every `report-*.spec.ts` | 440 passed, 56 skipped (project-specific), 0 failed |

### Captures (re-taken)

- `sample-chart-1440.png`, `-2000.png` and `-390.png` show the group default.
- `sample-chart-1440-10s-pull.png` has the end handle at the track's end, labelled `0:09.9`, and
  reads Everyone in range.
- `sample-chart-1440-everyone.png` is new: fight 6 under Everyone in range.

### Concerns (fix round)

- **The sample's group is in few pulls.** The logger was solo for most of the night. Only fight 6
  (4 of the group), fight 45 (2) and single-player pulls have the group in them. Fight 6 is still
  the default and shows four players.
- **Last-seen flags are report-wide.** Someone who left the party mid-night reads as an outsider
  in every pull. Engine item 2 fixes that.
- **Some copy still says "raid".** The Debuffs tab's note "the debuffs on the raid" keeps that
  word under a group scope. The Ratings tab lists every rated player under either friendlies
  scope.
- **Not done here:** the player's should-level notes (axis unit, marker key, step lines, a
  finding under the chart) and the ux shoulds.
