# Logs, reports, rankings and ratings in the flat web system: experience spec

Author: ux-designer, 2026-11-07. Mock first (tenet 15): the build matches these mocks region by region; nothing is built from this round.
Surfaces: `web/src/pages/logs.astro`, `reports.astro`, `reports/[id].astro`, `reports/sample.astro`, `report-shell.astro`, `rankings.astro`, `rankings/[slug].astro`, `ratings.astro`; `components/logs/*`, `components/report/*` (the report island, `ReportView.svelte` and its 40 children), `components/rankings/*`, `Rankings.svelte`, `MyReports.svelte`, `RecentReports.svelte`, `ReportRow.svelte`, `Upload.svelte`, `RatingCard.svelte`, `ForeverScoreCard.svelte`.
System: `design/specs/2026-11-05-web-layout.md` (section 2 principles, 3 tokens and type, 3.4 control bar, 3A primitives and the guard, 3.7 layouts, 5.13 build decisions). Latest pages in the system, whose rulings stand here: `2026-11-06-tiers.md` (RankList, the signed-in answer slot, T6), `2026-11-06-sim.md` (StatTable, Facts, the one ± rule, S1 to S8, Button case).
Content rules that stand unchanged: `2026-10-11-logs.md`, `2026-10-11-rankings.md`, `2026-10-28-log-coach.md`, `2026-10-30-controls-charts.md` (every figure's definition, the group scope, the threat estimate, the coach's sentences). This spec changes frames, order and hierarchy. It changes no figure's definition.
Memory rules applied: dense-clean-paperdoll, one-job-per-page, wide-screen-check, best-addon-standard, one-crest-language, provenance-quiet, prose-never-quotes-nightly-numbers, design-work-lives-in-repo, blizzard-profile-import (the header selector owns the character), and the parity gate (`design/parity.md`).
Type: Marcellus once (the H1, `--color-strong`: none of these pages has one class), Barlow everywhere else, no condensed face, no mono except literal code (`WoWCombatLog.txt`, `/combatlog`, the Queries SQL). One Plate per page. Quality and class colours from props. The HD circular crest wherever a character appears; a character whose class the log does not carry draws the ringed empty disc (`EmptyCrest`), never a dot.

## 0. Files

| What | Path |
|---|---|
| Generators (compose only from `ws_prims`; `build_page` asserts no page draws a surface) | `design/mocks/ws_logs.py` (/logs, /reports, /rankings, /ratings), `design/mocks/ws_report.py` with `ws_report_lib.py` (every report tab and state, /reports/sample) |
| Facts reader (copy modules and e2e stubs, bundled with esbuild, never retyped) | `design/mocks/logs_facts.mjs` -> `design/mocks/data/logs-facts.json` |
| Stub data copied out of e2e specs (gaps G7, G8: move them into `web/src/fixtures`) | `design/mocks/data/logs-rating-fight3.json`, `design/mocks/data/logs-reports-stub.json` |
| Primitives added (appended block "added by the logs round"; nothing existing changed) | `design/mocks/ws_prims.py`: `TabStripMore`, `LineChart`, `Spark`, `Menu`, `DropZone`, `Disclosure`, `EmptyCrest`, the `.st-sort` reset moved into the primitive. `ws_lib.header` accepts an account with no character (the empty disc) |
| Mocks (45) | `design/mocks/web/web-logs.html` (signed out), `web-logs-empty.html` (signed in, no logs), `web-logs-you.html` (signed in, with logs); `web-reports.html` (signed out), `web-reports-empty.html`, `web-reports-list.html`, `web-reports-failed.html`; `web-report-<tab>.html` for `summary`, `damage-done`, `healing`, `damage-taken`, `deaths`, `mechanics`, `timeline`, `threat`, `casts`, `buffs`, `debuffs`, `rating`, `coach`, `interrupts`, `dispels`, `resources`, `events`, `queries`, `compare`; `web-report-player.html` (Source = one player: abilities and gear at the pull), `web-report-fights.html` (the fight menu open), `web-report-bar.html` (phone controls open); states `web-report-loading.html`, `-error.html`, `-missing.html` (404), `-empty.html` (a pull with no data), `-processing.html`, `-waiting.html` (live, no fight yet), `-failed.html` (parse failed); `web-report-sample.html`; `web-rankings.html` (index), `web-rankings-empty.html`, `web-rankings-board.html` (signed out, every class), `web-rankings-you.html` (your row in place), `web-rankings-off.html` (your best kill off page, pinned), `web-rankings-unranked.html` (signed in, not on the board), `web-rankings-bar.html` (phone controls open); `web-ratings.html` |
| Renders (gitignored) | `design/mocks/renders/web/<mock>-{390,1440,1920,2000}.png` and `-fold.png` (all 45 at all four widths) |
| Before/after at the same scale, and the tracked copies (256-colour) | `design/reviews/2026-11-07-logs/after/compare-logs-<page>-<w>.png` (19 pairs), the 1440 and 390 first screens of every mock and the 2000 ones of the key pages; `design/reviews/2026-11-07-logs/before/` (live and fixture captures). Generator `design/mocks/ws_logs_compare.py` |
| Regenerate | `node design/mocks/logs_facts.mjs && python3 design/mocks/ws_logs.py && python3 design/mocks/ws_report.py`, then `node design/mocks/ws_shoot.mjs <mock> 390,1440,1920,2000` for each name in `python3 design/mocks/ws_logs.py --list` and `ws_report.py --list`, then `python3 design/mocks/ws_logs_compare.py <before dir>` |

**Data in the mocks (the same files a parity spec serves).**
- Report tabs: `web/src/fixtures/report` (`fixture2abcd`, the e2e report): `meta.json`, `fights/<n>/summary.json`, and `fights/<n>/events.parquet`, which the Events, Timeline and Queries mocks read through DuckDB exactly as the shipped page does in the browser (the Queries result is the shipped template `damage-by-ability` run on the file). Fight 3 (Warden Kelthas, kill, 1:00, 5 players, 1 death) is the default; Interrupts and Dispels use fight 1 (the only pull with one of each); the empty state is fight 2 (one player, nothing logged); Compare sets fight 3 against fight 4.
- Coach: `web/src/fixtures/coach/fury-model.json` (the coach's own output), on fight 3 re-read as the Onyxia pull exactly as `tests/e2e/support/coach-stubs.ts` serves it (title, boss, 3:52, the two warriors Fury). The fixture's 60 s phases are not shown against 3:52 (G10).
- Rating: the `GET .../fights/3/ratings` answer `report-rating.spec.ts` serves (Elyra Duskvale, Holy Priest, 70), copied to `data/logs-rating-fight3.json` (G7).
- /reports/sample: `web/public/sample-report/sample`, pull 6 (Trash 12:43), Source `Your group` from the pull's own `group` (`lib/report/group.ts`), as the live page opens it. Checked against production on 2026-11-07: group DPS 3, five players, Soren Everhart 1.8 DPS, Melee 582 taken at 12.3% (the share is of every player), no deaths in the group. All match.
- Rankings: `tests/e2e/support/rankings-fixture.ts` `boardFor()` and `characterPage()` (signed out: the every-class board on Warden Kelthas; signed in: Obnoxious Yell, Fury Warrior, 14th of 212 on Onyxia; off page: 204th), through `logs_facts.mjs`.
- /logs and /reports: the `/v1/reports?mine=1` rows `reports-page.spec.ts` serves (Tuesday raid, Sunday dungeons; G8), the sample meta, an empty public feed (production's own state on 2026-11-07).
- /ratings: `web/src/pages/ratings.astro`'s own sample score (1,830, Epic), sample trend, weight table and prose, parsed from the source; `lib/score/copy.ts`, `lib/rating/copy.ts` through the facts file.
- Names and icons: build 1.60.1.70790 `spells.json` and `spellicons.json`. A spell the build has no icon for draws the dark socket (the fixture's retail ids, G2).
- **SAMPLE (typed, marked):** none beyond the stubs. The Coach viewer is `coach-stubs.ts`'s premium owner of Baelgrim.

**Before captures.** Production on 2026-11-07 at 390, 1440, 1920, 2000: `/logs`, `/reports`, `/reports/sample`, `/rankings`, `/ratings` (all signed out; the public feed and the rankings are empty in production). Signed-in and report states: the fixture build (`FOREVER_DATA=fixture`, origin/main of 2026-11-06) with the e2e stubs, at 1440 and 390: `/logs` signed in with and without logs, `/reports` empty and listed, `/reports/fixture2abcd` on every tab, mode and view, `/rankings/warden-kelthas` signed out, `/rankings` signed in. The Rankings mode capture failed on the fixture build (an unstubbed API read), so its before is the Summary's.

## 0A. Rulings

Pending. Every decision in section 16 is PROPOSED with a recommendation; the coordinator runs the wow-player persona on `design/reviews/2026-11-07-logs/after/` with the questions in section 17 and records the rulings here, then the mocks are revised to them.

## 1. Diagnosis of the live pages (measured 2026-11-07)

**/logs** (`before/logs-*`).
1. **Four gold-edged cards with icon chips for a page with one job.** Your reports, Recent public reports, Upload a log, The companion; caps titles in Barlow Condensed, a caps eyebrow "COMBAT LOGS" over a caps "LOGS". Nothing is primary.
2. **Signed out, the page cannot do its job and says so in three places.** "Sign in to see your reports." with a Sign in button in one card, "Uploading needs a signed-in account" as plain text in another, no drop zone. The one useful thing a signed-out visitor can open (the sample) is a row at the bottom of a card, y 615 at 1440.
3. **Width.** 1,248 px of content at 1440, 1920 and 2000 (the column is not widened); at 2000 the page uses 62% of the screen.
4. **Signed in, the newest report is not the answer.** The character header (shared band) leads with the latest kill; the newest report is a 13 px row inside a card; at 390 "Upload a log" starts at y 685 and the companion at y 1,040.

**/reports** (`before/reports-*`). One card ("REPORTS") inside a page head with an eyebrow. Signed in, two 13 px rows; signed out a sentence and a button. The card title restates the H1. Uses 1,248 px at every width.

**/reports/[id] and /reports/sample** (`before/report-*`, `sample-*`).
5. **Three rows of navigation before a number.** Modes (Analyze, Compare, Rankings, Mechanics), then views as boxed pills (Tables, Timelines, Events, Queries) with a Source select, then the tab row (Summary, Damage Done, Healing, Damage Taken, Deaths, Threat, More), in three visual languages (caps underline, outlined boxes, caps underline). Plus the Players / Window / Log scale controls inside the chart card. Eleven controls before the table.
6. **The table is below the fold.** On the sample at 1440 the Players card starts at y 865 under a 430 px chart; at 390 at y 770 under the actions, a fight strip, three nav rows and a collapsed chart line. On the fixture the Damage Done table starts at y 1,016.
7. **A 300 px fight column beside a 900 px report.** On the sample it lists 79 trash pulls in a box 680 px tall; on a boss night it is a short box beside a tall column (the shipped comment already notes the stretch problem).
8. **Every region is a card.** Damage Done (chart), Players, Damage taken by ability, Deaths, Performance rating, At pull, each with an icon chip and caps title; "What do these words mean?" is a seventh bordered strip.
9. **Names unresolved on the Buffs tab.** The fixture's Buffs tab reads "Spell #17" and "Spell #871" though the build's `spells.json` names them (Power Word: Shield, Shield Wall) (bug B2).
10. **Rating says "Ratings did not load." on the fixture** because no ratings route is stubbed in a plain fixture run; with the stub it shows "Compared with other Holy Priestss on this fight (40 logs)." (bug B1, a double plural in `ratingCopy.percentileBasis`).
11. **Nodes that are not players.** The fixture's Summary lists "Elyra Duskvale · Hardcore" and "Thalgrit" with a white dot for a class; the one crest language has the empty ringed disc for this.
12. **What is right and stays:** every figure and its definition, the group scope (sample: 265 players in range, the group is five), the threat estimate and its caveat, the coach's review, the deaths recap, the shared item and spell tooltips, Sim this fight and Copy link, the honest notices (missing fight, unknown source, parses unreachable).

**/rankings and /rankings/[slug]** (`before/rankings-*`).
13. **The board is a card with a filter wall.** Eight labelled selects in two rows plus a Today only switch, inside a card whose caps title restates the class filter; then the rows. Characters / Guilds is a pair of boxed buttons in the card head. At 1440 the first row is at y 616 on the every-class board.
14. **Signed out with nothing ranked, production shows an empty card with a disabled Encounter select** and the Guilds link.
15. **Signed in, the answer is a second band card** ("YOUR CHARACTER", "YOUR RANK 14th of 212 Fury Warriors", a gold SEE YOUR PARSES button) above the board card: two framed objects.
16. **Single-class boards draw no crest**, against the one crest language (decision K3).

**/ratings** (`before/ratings-*`). A 4,286 px reference page (6,874 at 390) of seven gold-edged cards; the sample score card and the sample rating card are two more framed objects; the caps headings "WHAT MOVES EACH PART", "ONE NUMBER FIRST, THEN SIX PARTS" and so on are condensed caps. Width 1,248 px everywhere.

## 2. The one question per page, and the answer in the first screen

| Page | Question | 390 x 844 (mock) | 1440 x 1000 (mock) | 1920 / 2000 |
|---|---|---|---|---|
| /logs signed out | "How do I get a log in, and what will I see?" | Title, the Upload plate with the disabled drop zone, Sign in and Open the sample report (plate y 235) | The Upload plate (y 248) with Sign in, the companion and Your reports in the rail | Same, wide column |
| /logs signed in, logs | "What is my latest report?" | The answer (Tuesday raid, Open it) under the title, then Your reports (plate y 364) | Answer at the right of the title; Your reports plate y 290; the companion beside it | Same |
| /logs signed in, none | "How do I get my first log in?" | The Upload plate with the live drop zone | Same, rail: the companion with pairing | Same |
| /reports | "Which report do I open?" | Rows from y 240 | Rows from y 194 (live: 13 px rows in a card at y 271) | Same |
| /reports/[id] | "How did this pull go, and who did what?" | Pull, outcome, group DPS under the title; the tab's table from y 624 to 720 (the phone head is two title lines, the answer and two tab lines) | The answer in the head; the tab's table from y 428 (Summary 471): up to nine 56 px rows in the first screen (the fixture has three to five). Live: y 865 to 1,016 | Same; the table's bar column caps at 280 px |
| /rankings/[slug] | "Where do I stand on this boss?" | Your rank under the title; the board from y 515 | Your rank in the head; board from y 389, ten rows in view | Same |
| /rankings | "Which boss board?" | Encounters from y 330 | Encounters from y 340 | Same |
| /ratings | "What is the score and what moves it?" | The sample score figure, then the five parts | The score in the head, the five parts plate at y 354 | Same |

Page height at 1440: sample report 2,149 -> 1,709; fixture Summary 2,312 -> 1,585; ratings 4,286 -> 3,528; signed-in rankings 1,935 -> 1,662. No mock scrolls sideways at 390 (`scrollWidth` 390 on all 45, measured), except the Timeline's plot, which scrolls inside its own frame as shipped.

## 3. Reference surfaces, and what is borrowed

- **Warcraft Logs report.** Borrowed: one fight picker at the top of the report, one row of tabs (Damage Done, Healing, Damage Taken, Deaths, Casts, Buffs, Debuffs, Interrupts, Dispels, Resources, Threat, Events, Queries), a source picker, the player row that opens that player's abilities, the death recap with last hits and health, the casts timeline, the events list. Not borrowed: three rows of nav, the fight list as a column.
- **Details! damage meter.** Borrowed: one row per actor, the class-colour bar scaled to the top, the figure right-aligned, the abilities one click in. The table leads; the graph is secondary (decision R4).
- **Wowhead / WCL rankings.** Borrowed: rank coloured by percentile on the parse ladder every log reader knows (grey, green, blue, purple, orange, pink), the character in its class colour, the report link per row.
- **Our tier list, sim and BiS as shipped in the new system.** Borrowed: the PageHeader with the answer at the right (label, Figure, caption, links), the ControlBar grammar (selects with a caret, segmented controls with gold selection, a Filters popover like BiS Sources), one Plate, open Sections, a 340 / 400 px rail, `StatTable`, `RankList`, `Facts`, `Notice`.
- **The in-game character panel** through the shared `Paperdoll` in `report` mode for the gear a player wore at the pull.

## 3A. Primitives (added by this round; a design-review item under tenet 10, decision P1)

All appended to `ws_prims.py`; built in `web/src/components/ui/`.
| Primitive | What | Replaces |
|---|---|---|
| `TabStripMore(tabs, active, md, xl)` | `TabStrip` with the shipped More overflow: the first `md` tabs below 1760, the first `xl` from 1760, the rest in a More menu. The More item reads `More: {tab}` and carries the gold underline when the active tab is inside it. Phone: wraps (the guild rule) and keeps the first `md` | `ModeBar`'s mode row, view pills and tab row; `MoreMenu`'s trigger |
| `LineChart(series, duration, height, phases, marks, step_ms)` | The report chart on S0: a legend list (name in class colour, figure), the plot, phase bands with caps names, death marks in `--wipe`, taunts in gold, a 5-tick time axis. `step_ms` places one value per second on a time axis (the summaries' series are 1 s buckets) | `TimeChart`'s card frame and header, `ThreatChart`'s card; the shipped canvas and brush stay inside |
| `Spark(values | marks)` | A 24 px line (a resource bar over the fight) or tick strip (a cast sequence) in a table cell | `ResourceGraphs` rows, `CastTable` sequence |
| `Menu(label, groups)` | Overlay of kind `menu` anchored under its control: grouped rows (title, sub, right figures), one `aria-checked` row with the gold selection | `FightSelector`'s column |
| `DropZone(lines, action, off)` | The upload target: a dashed socket, the one dashed surface (as the empty doll socket); `off` at 55% | `Upload.svelte`'s dashed box |
| `Disclosure(summary, body)` | The shipped disclosure, styled once | the phone's collapsed chart strip |
| `EmptyCrest(px)` | The ringed empty disc for a character with no class in the log | the white dot |

## 4. /logs

`Page layout="split"`. Region order: `head`, the Plate's Section, open Sections, rail.
- **PageHeader:** title `Logs`; meta one line, `logsBandCopy.jobLine`. The eyebrow retires. Signed in with a report, `answer`: label `Your latest report · {age}`, `Figure size="big"` holding the report's title (strong), Small `{zone} · {n} bosses · {n} fights`, link `Open it ›`. No answer signed out or with no report. The shared CharacterHeader band leaves /logs (decision L2).
- **The one Plate is the next step (decision L1).**
  - Signed out (`web-logs`): Section `Upload a log` (meta `For a night already logged`), Plate holding Upload's two paragraphs, the `DropZone` at 55% (`aria-disabled`), the line `Uploading needs a signed-in account; the report is filed under it.` and one row: primary `Sign in with Battle.net`, link `Open the sample report ›` (decision L3).
  - Signed in, no logs (`web-logs-empty`): the same Plate with the live `DropZone`; under it Section `Your reports` with `myReportsCopy.empty`.
  - Signed in with logs (`web-logs-you`): Section `Your reports` (meta `{n} reports, newest first`, action `All your reports ›`) holding the Plate of report `Row`s (see 5); then Section `Upload a log`, open (no Plate), the paragraphs and the `DropZone`.
  - After a file is chosen (shipped flow, not re-mocked): the Title `Field` and the `Who can open it` Segmented (`Public`, `Unlisted`, `Private`, `Guild`, each with its shipped note) appear under the zone; Upload is the primary Button.
- **Section `Recent public reports`** (action `All public reports ›`): the public feed as `Row`s, or `recentReportsCopy.empty`, then the sample `Row` (title, `{zone} · {n} fights`, Tag `Sample report`, date).
- **Rail:** Section `The companion` (meta `Live, while you raid`): `Log live while you raid.`, secondary `Get the companion`; signed in, Label `Pair it with your account`, `pairingCopy.companionFirst`, secondary sm `Get a pairing code` (the shipped pairing island). Signed out, Section `Your reports` with `Sign in to see your reports.` (no second button: one primary per view).

## 5. /reports (Your reports)

`Page layout="ledger"`. PageHeader title `Your reports`, meta `Every fight you logged, newest first.` and, listed, `{n} reports`; `actions` secondary `Upload a log` (signed in). The card title `Reports` and the eyebrow retire.
- **Plate of report `Row`s** (link rows to `/reports/<id>`): title (the report title, or its zone), sub `{zone} · {n} bosses · {n} fights · {credited name or Not credited}` (`accountReportSummary`), trail a muted status `Tag` while not complete (`Ready`, `Live`, `Processing`) and the age (`today`, `2 days ago`, then `17 Sept`). A guild member gets the shipped guild `TabStrip` (`Yours`, each guild) in the head's `tabs` slot; `logsCopy.guildEmpty` per guild.
- **Empty** (`web-reports-empty`): one Plate line `No logs yet. Upload a log, or let the companion send your next raid.`, primary `Upload a log`, link `Open the sample report ›`.
- **Signed out** (`web-reports`): `Sign in to see every report you logged.`, primary `Sign in with Battle.net`, link `Open the sample report ›`.
- **Failed** (`web-reports-failed`): `Your reports did not load.` with secondary sm `Try again`; session failed: `Your account did not load.` with the same.
- **Older reports:** secondary `Older reports` at the plate's foot when the API has more (shipped).

## 6. /reports/[id] and /reports/sample

### 6.1 Frame (every tab)
`Page layout="split"`: main 1fr and the rail (340 px from 1024, 400 from 1760).
- **PageHeader.** Title: the report's title (Marcellus, `--color-strong`). Meta line: `{zone} · {d MMM yyyy} · {n} fights`, then a muted `Tag` `Not credited` whose `title` carries the shipped credit sentence (decision R8), or `Sample report` on the sample followed by `reportCopy.sampleNote`. A credited report reads `Logged by {name}` in place of the Tag. Status `Live` (kill colour) or `Quiet` while live.
- **Answer (decision R1).** Label `{pull} · {Kill | Wipe at n%} · {m:ss}` (a trash pull: `Trash · 12:43`). Figure `big`: the group's DPS over the pull in `--color-strong`, unit `DPS`; caption `{Group | Your group} damage · {n} death(s)`. Signed in with a character in the pull: the figure is that character's DPS in gold, caption `{name} · {spec} {class} · {damage} damage` (the page leads with the player, tenet 14). Actions: secondary sm `Sim this fight`, `Copy link` (`Copied`, `Copy failed`). Phone: the answer stacks under the title; the actions stay words.
- **Tabs (decision R3).** `TabStripMore` in the head's `tabs` slot: `Summary`, `Damage done`, `Healing`, `Damage taken`, `Deaths`, `Mechanics`, `Timeline` (below 1760); from 1760 also `Threat`, `Casts`, `Buffs`, `Debuffs`, `Rating`, `Coach`; in More: `Interrupts`, `Dispels`, `Resources`, `Events`, `Queries`, `Compare`. `Coach` is listed only for the owner of a character in the pull or an officer (shipped); without it the from-1760 set is 12. Modes and views retire as separate rows: Mechanics, Compare, Timelines (now `Timeline`), Events and Queries are tabs; the Rankings mode retires to a link (`Rankings for {boss} ›` in the Summary's rail, decision R3b). The URL keeps `tab=`; old `mode=` and `view=` links map to the tab.
- **ControlBar** (`label="Report choices"`, group labels from 1920, sticky): `Fight` (select reading `{pull} · {outcome} · {m:ss}`; opens the fight `Menu`, decision R2), `Source` (select: `All friendlies`, `Your group`, `Your raid`, `Everyone in range`, `All enemies`, then players under `Tanks`, `Healers`, `Damage`), `Window` (select: `Whole fight`, each phase, each death's 20 s, the brush's own range). Damage done, Healing and Damage taken add `Target` (`Every target`), `Ability` (`Every ability`) and a `Filters` button opening a popover (`Boss damage only`, `Players only`, `Count overkill`, `Ignore events while dead`, `Clear filters`; the button reads `Filters ({n})` when any is on). Threat adds `Enemy` (`Every enemy`). Compare replaces Window with `Compare with` (the fight menu, or a player under `Or a player`) and `Metric`. Phone: the 48 px summary `{pull} {m:ss} · {source}` and `Change ▾` (`web-report-bar`).
- **Fight menu** (`web-report-fights`): groups `Boss pulls` (first row `All boss pulls · {n} pulls`, the night view) and `Trash` (each trash group as on the shipped selector); a row: `{pull} · {outcome} · {m:ss}` over `{time of day} · {n} deaths` (trash `{n} mobs died`), the current row selected. 480 px wide, 70vh tall, scrolls; Escape closes; arrows move, Enter picks. Phone: the same menu full width under the bar.
- **Notices** (`summary` region, one line each, `Notice tone="neutral"`): the shipped lines for an unknown source, unknown ability, missing fight, a fight that did not load, and `The rankings could not be reached, so the Parse column is unknown; every other figure is from the log.` with `Try again` (boss pulls only).
- **Rail on every tab:** Section `This pull` (meta `Started {hh:mm:ss}`, the viewer's clock): `Facts` Pull, Result, Length, Players (in the source scope), Deaths (in scope), Mobs died (trash), Phases (`Phase 1 0:00–0:14 · Phase 2 0:14–1:00`) (decision R7). Then the tab's own rail Sections below.
- **Popovers not mocked separately:** the `Filters` popover is the BiS `Sources` popover pattern (`web-bis-sources`), and the More menu is the `Menu` of `web-report-fights` with the tab names.
- **The one Plate holds the tab's table**, always first in main (decision R4); the chart is an open Section under it on the tabs that chart (R6).

### 6.2 Each tab (`web-report-<tab>`)
| Tab | Plate (a `StatTable` unless named) | Main, after the Plate | Rail, after This pull |
|---|---|---|---|
| Summary | Section `Players` (meta `{n} players · {n} deaths · {m:ss}`): crest, name in class colour (realm tag muted when not the report's home realm), sub `{spec} {class} · item level {n}`; columns `Parse`, `DPS`, `HPS`, `Taken / s`, `Active`, `Deaths` (red when above 0). Phone: name, sub, DPS under it | `Damage over the fight` chart; Section `Damage taken by ability` (open StatTable: spell icon, name, Meter, `Share`, `Amount`, `Per sec`; the caption `Share is of everything every player took in this window, whatever the source scope shows.`; action `Damage taken ›`) | `Deaths` (Rows: time, crest, name, `{ability} · {source} · {amount}`; action `Deaths ›`; empty `Nobody died.`), `Performance rating` (a Row per rated player with `Figure` `70 of 100`; action `Rating tab ›`; trash `No rating for trash. Ratings are made for boss fights.`; none `No ratings for this fight yet.`), `At pull` (Rows: crest, name, `{spec} · item level {n} · {n} items`, `Gear ›` opens the player's view) |
| Damage done, Healing, Damage taken | Section `{tab}` (meta `Click a player to see their abilities`): crest, name, sub, a class-colour `Meter` (share of the top), `Amount` (Healing: `Effective`, sub adds `{n} overheal ({p}%)`), `DPS` / `HPS` / `Taken / s`, `Share`, `Active`; a Total row `Total, {n} players`. Sort heads are buttons with `aria-sort` | the chart of the tab's metric | `Targets` / `Healed` / `Sources` (Rows with a Meter) |
| Player view (`web-report-player`, Source = one player, decision R5) | Section `{name}'s damage`: abilities (spell icon, name, sub `{hits} hits · {crits} crit · {min}–{max}`, Meter, `Amount`, `DPS`, `Share`, `Hits`, `Crit`) | the chart; Section `Gear at the pull` (meta `{spec} · item level {n}`, action `Open in the planner ›`): the shared `Paperdoll` in `report` mode, open | `Targets` |
| Deaths | One block per death: a Row (time, crest, name, `Killed by {source} · {ability} · {amount} ({overkill} overkill) · {max} health`, trail `The 20 s before it ›`), the last hits (`Amount`, `Health` after, a green health Meter, `dead` on the killing blow), `Holding: {auras}. No healing landed in the last 10 seconds.` | the chart with death marks | none |
| Mechanics | Rows per mechanic in the encounter's table: icon, name, a kind `Tag` (`Avoidable` red when someone failed it, else muted; `Interrupt`, `Dispel`, `Unavoidable`), sub `{who and how often} · {the table's note}`, `Failed`, `Damage`. Empty: `Nothing the table lists went wrong.`; no table: `No boss of the night has a table yet.` | none | `Clean` (Rows of players with no avoidable damage) |
| Timeline | `Lanes`: a 160 px label column (crest and name per player, then the boss), one lane per actor on one time axis: casts (gold ticks), hits (class colour), buffs (green bars), debuffs (red bars), a death (red line); legend above; the axis under | none | none |
| Threat | Section `Estimated threat` (meta `On {enemy} at the end of the fight`): crest, name, Meter, `Threat`, `Share`, `Of the top` | Section `Estimated threat on {enemy}`: the cumulative line chart with taunts (gold) and deaths (red) and the shipped caveat `An estimate: 1 threat per point of damage and 0.5 per point of healing, split across the enemies in combat. No stance, form, talent or taunt modifiers yet.` + `What it leaves out` | `Taunts` (Rows: time, icon, `{name} taunted {enemy}`, the spell) |
| Casts | spell icon, spell, sub the caster (class colour; a pet reads `Ashfang (Thalgrit's pet)`), a `Spark` tick strip of the sequence, `Casts`, `Per min`, `Failed` | none | none |
| Buffs, Debuffs | aura icon, name, sub `on {target} · from {applier} · up at the pull`, an uptime Meter (muted on Debuffs), `Up`, `Applied` | none | none |
| Rating | First row: crest, `{name}'s rating`, the basis sentence, `Figure big gold` `70 of 100`; then a row per part: name, the part's one-line definition, a score Meter (green from 75), `Of 100`, `Weight`, `Basis` (`Percentile of 40`, `Absolute`); an excluded part reads `ratingCopy.excludedReason` with dashes. Section action `How is this calculated? ›` | none | `Performance rating` (the rated players), `The cap` (`ratingCopy.cappedRule`) |
| Coach | Section `{persona.label}` (meta `{name} · {spec} {class} · reviewed {d MMM}`): the opening sentence (17/26, strong), `Fix first` (gold Label), its title (H3), why, the anchors (a spell with its icon; muted time Tags), `The drill` and its sentence; `Then fix` blocks; `Doing well` with green checks | none | `Measured against` (the basis labels), `What the coach could not check` (the review's list) |
| Interrupts, Dispels | spell icon, `{name} · {spell}` (class colour), sub `Interrupted {cast} on {target}` / `Removed {aura} on {target}`, `Count`. Empty `Nothing interrupted.` / `Nothing dispelled.` | none | none |
| Resources | crest, name, sub the power, a `Spark` of the bar over the fight in the class colour, `Gained`, `Spent`, `Wasted`, `Time empty`. Empty `No resource changes in this window.` | none | none |
| Events | Rows in order: the event's spell icon, a sentence with names in class colour (`Morrowlyn hits Warden Kelthas with Frostbolt`, `Thalgrit dies`), `Time`, `Amount`. Empty `No rows matched.` | none | `Show` (checkbox chips: Casts, Damage, Heals, Auras, Deaths, Power) |
| Queries | `Starting points` Segmented (the five shipped templates), Label `Query`, the SQL in a code block (mono: literal code), primary `Run`, `{n} rows`, then the result as a StatTable with the query's own column names | none | `Columns` (the parquet's column list) |
| Compare | crest, name, sub, `This pull`, `{other pull}`, `Difference` (green gain, red loss) | none | none; Notice `Players missing from the other pull show a dash.` |

**/reports/sample** (`web-report-sample`) is the same page: Summary on pull 6, Source `Your group`, Tag `Sample report` and `reportCopy.sampleNote` in the head, no credit line, no Coach.

### 6.3 States (`web-report-<state>`)
| State | Treatment |
|---|---|
| Loading (`loading`) | The head reserves its size with Skeleton blocks (title 520 x 40, meta 320 x 16, the answer); one disabled `Summary` tab; the bar; the Plate's heads and six 56 px skeleton rows; the rail's This pull as five skeleton lines. Same geometry as the ready page: no shift |
| Error (`error`) | Title `Report`; Plate line `Report data did not load.` with secondary sm `Try again` (role alert) |
| Not found (`missing`) | `No report with that id.` and `Open your reports ›` |
| A pull with no data (`empty`, fight 2) | Every tab's Plate shows its own empty line (`Nothing in this window.` on the tables), the chart Section `Nothing in this window.`, This pull still filled |
| Processing (`processing`) | Title `Your upload`, meta `Uploaded {age}`; one disabled tab; Plate line `Parsing your log… this takes a minute or two for a long night` with the spinner (role status) |
| Live, waiting (`waiting`) | Title the report; meta Tag `Live` and `Logging now`; `reportCopy.waiting`; after quiet minutes `reportCopy.quiet(n)` and the Tag `Quiet` |
| Parse failed (`failed`) | `This log could not be parsed.` + `Check your logging setup` (role alert) |
| Night view (All boss pulls) | Only Summary, Damage done, Healing, Damage taken, Mechanics and Rating mean something over a night; the other tabs read the shipped `Ratings are one pull’s. Pick a boss pull from the list to see one.` pattern per tab |
| A tab's lazy chunk failed | `This view did not load.` + `Try again` inside the Plate |

## 7. /rankings and /rankings/[slug]

`Page layout="ledger"`.
- **/rankings** (`web-rankings`): PageHeader `Rankings`, meta `How characters rank on each encounter, kill by kill.` and `Updated {d MMM}`; Notice `Comparing specs, not characters?` + `Tier list ›`; Section `Choose an encounter` (meta `{n} encounters ranked`) with the Plate of encounter `Row`s (`Characters and guilds, every class and spec`, trail `Open the board ›`) (decision K4); the legend line under it. Signed in, `/rankings` goes to the encounter the character last killed (shipped).
- **Empty** (`web-rankings-empty`): Plate `No encounters have been ranked yet.` / `The Guilds board ranks guild progress across every encounter and needs none.` / `Open the Guilds board ›`.
- **/rankings/[slug]** (`web-rankings-board`, `-you`, `-off`, `-unranked`):
  - PageHeader title the encounter (`Onyxia`, decision K1); meta `Rankings · {cohort, or Every class and spec} · {ruleset}` and `Updated {d MMM}`.
  - Signed-in `answer` (the tier list's slot): Label `Your rank`, crest 40 and the name in class colour, `Figure big gold` `{ordinal}` with unit `of {total} {Spec} {Class}s`, Small `On {encounter}, {ruleset} ruleset`, link `See your parses ›`. Unranked: Small `Not ranked yet on {encounter}. Upload a log of a kill to get on the board.` and `Upload a log ›`. Failed: `Your rank did not load.`
  - ControlBar (`label="Ranking choices"`, labels from 1920): `Encounter`, `Class`, `Spec`, `Ruleset` selects; `Metric` segmented `Damage | Healing`; `Board` segmented `Characters | Guilds`; `Filters` popover holding `Region`, `Phase`, `Faction`, `Today only` (decision K2). Phone summary `{encounter} · {spec or class}`.
  - Notice `Comparing specs, not characters?` + `Tier list ›`.
  - Section `{cohort}` (meta `{total} ranked, each character's best kill`) holding the Plate: `RankList` with heads `#`, `Character`, `Share of the top`, `DPS`, `Executed`, `Length`, `Date`; a `RankRow` per kill: rank in its percentile colour (`rankColor`), crest 32 (decision K3), name in class colour, sub `{guild} · {spec} {class} · {size}-man`, class Meter of value over the top, `{value} DPS`, `{executed}%`, `m:ss`, `{n} days ago` (then `13 Dec`), links `Report` (gold) and `Build`. Your row in place: selected (gold ground and rail) with Tag `You`. Off page: pinned first under `Your best kill`, the page's rows after it (decision K5). A moderated row: Tag `Removed by a moderator` or `Under review by a moderator` (shipped). The foot: `Page 1 of {n}`, `Previous` (muted on page 1), `Next`.
  - Footer note (`footer-note`): `Rank colour is the percentile:` `99+` `95+` `75+` `50+` `25+` `under 25`, each in its colour.
  - Guilds board: the same RankList with the faction emblem as the crest, guild name, `{n} bosses`, progress Meter.
  - Empty board: `No ranked fights yet for this filter. Rankings fill in as reports are uploaded.` + `Upload a log ›`. Failed: `Rankings did not load.` + `Try again`.

## 8. /ratings

`Page layout="split"`. A reference page, one level down (tenet 14): linked from every rating surface (`How is this calculated?`), not from the nav.
- PageHeader `Forever Score and ratings`; meta `Forever Score: your whole record, 0 to 2100.` and `Performance Rating: one fight, 0 to 100.`; `tabs` two anchors `Forever Score`, `Performance Rating` (the jump row); answer: Label `Sample · Forever Score`, the figure `1,830` in the tier's colour with `of 2,100 · Epic`, Small `Most to gain: Gear, 88 below its top.` (decision T1). Signed in with a scored character, the answer is that character's score with `Sample` dropped (data gap G9 until `/me` carries it).
- Main: Section `Forever Score` (meta the caption): the lead, the one Plate (`StatTable` of the five parts: name, its `moves` sentence, a Meter of value over its ceiling, `Points`, `Top` as `up to 500`), then `explainerWhat` and `explainerLaunch` (muted). Section `Performance Rating`: the bridge sentence, three Figures (`Latest fight 65 of 100` gold, `Last 6 fights` with a Spark and `Up 25 since the first`, `Weakest part` `Weakest part of a fight: Activity.`), the visibility paragraph (muted) and the two "One number first" paragraphs. Section `What each component measures` (open Rows: part, its paragraph). Section `Weights, by role` (open StatTable `Component`, `DPS`, `Healer`, `Tank`; the two muted notes). Section `Ranked against your spec, or absolute` (three paragraphs). Section `The cap` (`ratingCopy.cappedRule`).
- Rail: Section `Tiers` (a Row per tier, name in its quality colour, range; the sample's tier selected; `tiersNote`), Section `How to get your logs here` (`explainerLogsBody`, secondary `Get the companion`).
- Prose stays as shipped except that em dashes in the page's own paragraphs are a copy item for the content owner (not changed by this spec).

## 9. Every state of every interactive element

| Element | Default | Hover | Focus-visible | Active | Selected | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|---|
| TabStrip tab / More | 14 px 600 `--color-nav`, 44 px | `--color-strong` | 2 px gold ring, 2 px offset | `#1a2030` | `--color-strong`, 2 px gold underline, `aria-current="page"`; More reads `More: {tab}` | while processing: 50%, `aria-disabled`, title `reportCopy.parsing` | n/a | n/a | n/a |
| More menu | `Menu`, items as the tabs | S2 | ring on the item | pressed | gold ground and 3 px bar | n/a | n/a | n/a | n/a |
| ControlBar select / segmented / Filters | outlined 40 px (44 phone), 14 px 600 | S2, `--color-strong` | ring | `#1a2030` | segmented: gold text, `#e5b9551a`, 2 px underline; Filters `Filters ({n})` | Compare in night view: 50%, title `Compare is one pull’s.` | the bar never shifts; the Plate shows its skeleton | n/a | the bar is unaffected |
| Fight menu row | as Menu | S2 | inset ring | pressed | gold ground and bar, `aria-checked` | a pull still being logged: `The fight is still being logged` title, selectable | n/a | `No boss pulls` group line | n/a |
| StatTable row (actor) | plate ground, hairline | S2 | the name is the link; ring | pressed | the Source player's row when the bar's Source is them | n/a | 56 px skeleton rows | the tab's empty line | `This view did not load.` + Try again |
| Sort head | 12 px caps muted | `--color-strong` | ring | n/a | `--color-strong` + gold `▾`, `aria-sort` | n/a | n/a | n/a | n/a |
| Spell / item icon | 24 to 28 px, quality border for items | the shared tooltip after 150 ms | the tooltip on focus | n/a | n/a | n/a | the dark socket at size | the socket when the build has no icon | n/a |
| Report Row (/logs, /reports) | open row | S2 | ring | pressed | n/a | n/a | three skeleton rows | per section 5 | per section 5 |
| DropZone | dashed `--line`, 168 px | border `--color-nav` | ring on `choose a file` | n/a | dragging: gold border, gold 6% ground | signed out 55%, `aria-disabled` | uploading: the shipped progress line | n/a | `The upload did not finish. Your file is untouched; try again.` |
| RankRow | as the tier list | S2, rail slice in class colour | ring on each link | pressed | your row: gold ground, gold slice, Tag `You` | n/a | the shipped skeleton | the empty board line | `Rankings did not load.` |
| Pager | `Previous` muted on page 1, `Next` gold | gold-hover | ring | pressed | n/a | muted | n/a | hidden on one page | n/a |
| Chart (phone Disclosure) | collapsed `Show ▾` | n/a | ring | n/a | open `Hide ▾` | n/a | Skeleton at height | `Nothing in this window.` | n/a |
| Buttons | system Button (primary gold, secondary outline) | as system | ring | as system | n/a | 45% | `Copying…` n/a | n/a | `Copy failed` |

Keyboard: the bar's selects and segments in order, then the TabStrip, then the Plate (sort heads, then each row's link), then the main Sections, then the rail. Arrow keys inside a segmented control and inside the fight menu and the More menu; Escape closes either and returns focus. Touch: 44 px targets on phone; tooltips on long-press. Reduced motion: no transitions, the spinner is static.

## 10. Copy (verbatim)

Figures in braces render from data. Strings from the copy modules are quoted by key and stay as shipped unless this list changes them.
- /logs: H1 `Logs`; meta `Upload a combat log or log live, then read every pull.`; answer label `Your latest report · {age}`, Small `{zone} · {n} bosses · {n} fights`, link `Open it ›`; Sections `Upload a log` (meta `For a night already logged`), `Your reports` (meta `{n} reports, newest first`, action `All your reports ›`), `Recent public reports` (action `All public reports ›`), `The companion` (meta `Live, while you raid`); Upload paragraphs as `Upload.svelte`; zone `Drop WoWCombatLog.txt here, or` / `choose a file`; signed out `Uploading needs a signed-in account; the report is filed under it.`, `Sign in with Battle.net`, `Open the sample report ›`; `myReportsCopy.empty`, `myReportsCopy.signedOut`; companion `Log live while you raid.`, `Get the companion`, `Pair it with your account`, `pairingCopy.companionFirst`, `Get a pairing code`. Retired: `logsBandCopy.eyebrow`, the card titles' caps, the phone header action `Open your latest report →`.
- /reports: H1 `Your reports`; meta `Every fight you logged, newest first.`, `{n} reports`; action `Upload a log`; row sub `{zone} · {n} bosses · {n} fights · {credited or Not credited}`; ages `today`, `{n} days ago`; empty `No logs yet. Upload a log, or let the companion send your next raid.`; signed out `Sign in to see every report you logged.`; failed `Your reports did not load.`, `Try again`; `Older reports`. Retired: `yourReportsPageCopy.eyebrow`, `cardTitle`.
- Report head: answer label `{pull} · {Kill | Wipe at {n}%} · {m:ss}` (trash `Trash · {m:ss}`); captions `Group damage · {n} death(s)`, `Your group damage · …`, `{name} · {spec} {class} · {n} damage`; Tags `Not credited`, `Sample report`, `Live`, `Quiet`; `Logged by {name}`; actions `Sim this fight`, `Copy link`.
- Report bar: groups `Fight`, `Source`, `Window`, `Target`, `Ability`, `Filters`, `Enemy`, `Compare with`, `Metric`; values as `reportCopy` (`All friendlies`, `Your group`, `Your raid`, `Everyone in range`, `All enemies`, `Every target`, `Every ability`, `Whole fight`, `Pick a fight`, `Boss pulls`, `Trash`, `Or a player`); fight menu `All boss pulls`, `{n} pulls`, `{hh:mm:ss} · {n} deaths`, `{n} mobs died`; phone `{pull} {m:ss} · {source}`, `Change ▾`.
- Tabs: `Summary`, `Damage done`, `Healing`, `Damage taken`, `Deaths`, `Mechanics`, `Timeline`, `Threat`, `Casts`, `Buffs`, `Debuffs`, `Rating`, `Coach`, `Interrupts`, `Dispels`, `Resources`, `Events`, `Queries`, `Compare`, `More`, `More: {tab}`. (Sentence case; `Damage Done` in title case retires.)
- Report sections and heads: per 6.2 (`Players`, `Damage taken by ability`, `Damage over the fight`, `Healing over the fight`, `Damage taken over the fight`, `Group {n} {DPS | HPS | taken / s} over the whole fight`, `Click a player to see their abilities`, `{name}'s damage`, `Gear at the pull`, `Open in the planner ›`, `Targets`, `Healed`, `Sources`, `This pull`, `Started {hh:mm:ss}`, `Pull`, `Result`, `Length`, `Players`, `Deaths`, `Mobs died`, `Phases`, `At pull`, `Gear ›`, `Rankings for {boss} ›`, `Estimated threat`, `Estimated threat on {enemy}`, `On {enemy} at the end of the fight`, `Taunts`, `{name} taunted {enemy}`, `Casts`, `Every caster, {m:ss}`, `{n} on players · uptime over the whole fight`, `on {target} · from {applier} · up at the pull`, `Deaths`, `{n} in this fight`, `Killed by {source} · {ability} · {amount} ({n} overkill) · {n} health`, `The 20 s before it ›`, `Holding: {auras}. No healing landed in the last 10 seconds.`, `Mechanics`, `Problems first, then what went right`, `Clean`, `No avoidable damage`, `Timeline`, `Every player, {m:ss} · hover for the event`, legend `Cast`, `Hit, in the class colour`, `Buff up`, `Debuff up`, `Death`, `Events`, `{n} events in this window, in order`, `Show`, `Queries`, `SQL over this fight’s events, run in your browser`, `Starting points`, `Query`, `Run`, `{n} rows`, `Columns`, `Compare`, `This pull`, `Difference`, `Players missing from the other pull show a dash.`, `Measured against`, `What the coach could not check`, `Fix first`, `Then fix`, `The drill`, `Doing well`, `The cap`).
- Column heads: `Player`, `Parse`, `DPS`, `HPS`, `Taken / s`, `Active`, `Deaths`, `Amount`, `Effective`, `Share`, `Ability`, `Per sec`, `Hits`, `Crit`, `Threat`, `Of the top`, `Spell`, `When`, `Per min`, `Failed`, `Aura`, `Uptime`, `Up`, `Applied`, `Last hits`, `Health after`, `Health`, `Mechanic`, `Who, with what`, `Count`, `Over the fight`, `Gained`, `Spent`, `Wasted`, `Time empty`, `Event`, `Time`, `Part`, `Score`, `Of 100`, `Weight`, `Basis`.
- Report empties and errors: `Nobody died.`, `Nothing in this window.`, `Nobody took damage in this window.`, `Nothing interrupted.`, `Nothing dispelled.`, `No resource changes in this window.`, `No casts in this window.`, `No rows matched.`, `Nothing the table lists went wrong.`, `No boss of the night has a table yet.`, `No threat in this window.`, `Nobody taunted.`, `No {buffs | debuffs} in this window.`, `Report data did not load.`, `No report with that id.`, `Open your reports ›`, `This view did not load.`, `Try again`, and every `reportCopy` and `ratingCopy` state string as shipped.
- Rating: `{name}’s rating`; `Compared with other {spec} {class}s on this fight ({n} logs).` (B1 fixed); `Percentile of {n}`, `Absolute`; `How is this calculated? ›`; the part definitions as `ratingCopy.componentDefinition` without the `{Part}: ` prefix.
- Rankings: H1 `Rankings` (index) or `{encounter}`; meta `How characters rank on each encounter, kill by kill.`, `Rankings · {cohort} · {ruleset}`, `Updated {d MMM}`; `Comparing specs, not characters?`, `Tier list ›`; `Choose an encounter`, `{n} encounters ranked`, `Characters and guilds, every class and spec`, `Open the board ›`; answer `Your rank`, `{ordinal}`, `of {total} {Spec} {Class}s`, `On {encounter}, {ruleset} ruleset`, `See your parses ›` (the arrow `→` becomes `›`); heads `#`, `Character`, `Share of the top`, `DPS`, `Executed`, `Length`, `Date`; links `Report`, `Build`; `You`, `Your best kill`; `Page {n} of {n}`, `Previous`, `Next`; legend as `rankingsCopy`. Retired: `rankingsCopy.eyebrow`, the card title `Every class and spec` as a card head (it becomes the Section title).
- Ratings: as `scoreCopy` and `ratingCopy` and the page's own paragraphs; `Sample · Forever Score`, `of {max} · {Tier}`, `Points`, `Top`, `Latest fight`, `Last {n} fights`, `Up {n} since the first`, `Weakest part`, `Tiers`.

## 11. What each figure reads; gaps; bugs found

| Shown | Reads | Exists? |
|---|---|---|
| Group DPS, a player's DPS / HPS / taken, active, deaths | `summary.roster`, `damage_done[].effective` over `duration_ms`, in the source scope | yes |
| Source `Your group` | the pull's own `group` (engine 0.6.0), else `meta.group` (`lib/report/group.ts`) | yes |
| Share of damage taken | over every player, whatever the scope (the shipped caption) | yes |
| Charts | `series`, one value per second from the pull (1 s buckets); `phases`; `deaths[].at_ms`; `taunts` | yes |
| Parse | the rankings for the pull's players | **no fixture serves it** (G1) |
| Death recap | `deaths[].killing_blow`, `last`, `heals`, `auras_held` | yes |
| Mechanics | `mechanics.rows`, `interrupts`, `dispels` | yes |
| Timeline, Events, Queries | `events.parquet` | yes (DuckDB) |
| Gear at the pull | `combatants[].gear` (slot order), the items file | yes; fixture items are retail ids, unknown to our data (G2) |
| Rating | `GET .../fights/{n}/ratings` | only in an e2e spec (G7) |
| Coach | `GET .../coach` | yes (`fixtures/coach`) |
| Reports list | `GET /v1/reports?mine=1` | only in an e2e spec (G8) |
| Latest report on /logs | `newestOwnReport()` | yes |
| Rank, total, your rank, off-page rank | `GET /v1/rankings` `rows`, `total`, `character_rank` | yes (stub) |
| Rank colour | `rankColor(rank, total)` | yes |
| Forever Score, parts, tiers | the sample in `ratings.astro`, `COMPONENT_MAX`, `TIER_FROM` | sample only (G9) |

Gaps (named, not invented):
- **G1. No parses in any served fixture** for the report's players. The Parse column is `–` under the shipped notice. A parity stub of the rankings answer for fixture2abcd's players is needed before the column can be mocked filled.
- **G2. fixture2abcd is a retail-shaped synthetic log** (Sanguine Depths, Anima Lash, item level 183, spell ids 334660, 320462 with no icon in our build, item ids 175850 and 175885 unknown to our item data). The mocks show what it serves: dark sockets and `Item 175850 · Not in our item data yet`. Recommendation for the data lane: replace it with a real Forever pull in the sample's data model (decision D1), so the report's parity mocks show Forever spells and items.
- **G3. Two players have no class in the roster** (Thalgrit, Elyra Duskvale): they draw the empty crest. The rating answer names Elyra a Holy Priest, so the Rating tab can colour her; the roster cannot.
- **G4. The charts must place values by time.** The summaries' series are 1 s buckets of different lengths (Baelgrim 31, the boss 11); drawing them by index stretches every series to the fight's length. The mock places them by `step_ms`.
- **G5. The pull's clock.** The page shows the viewer's own time of day; the mocks render UTC (as the captures did).
- **G7. The rating stub lives in `report-rating.spec.ts`.** Move it to `web/src/fixtures/report/ratings-3.json` so the parity spec and the mock read one file (the mock reads `design/mocks/data/logs-rating-fight3.json` until then).
- **G8. The reports-list stub lives in `reports-page.spec.ts`.** Move it to `web/src/fixtures/reports/mine.json` (the mock reads `design/mocks/data/logs-reports-stub.json` until then).
- **G9. No score on `/me`.** The ratings page shows the sample; a signed-in answer needs the selected character's Forever Score from the character API.
- **G10. The coach's Onyxia re-read keeps the fixture's 60 s phases** against a 3:52 pull. The mock hides Phases on the Coach tab.
- **G11. `execution_score` is 0.9 on every stub row**, so Executed reads 90% down the board. Real boards vary.

Bugs found while measuring (for engineering):
- **B1.** `ratingCopy.percentileBasis` pluralises twice: `Compared with other Holy Priestss on this fight (40 logs).` (`roleNoun` adds `s`, then the template adds another). The mock shows the fixed sentence.
- **B2.** The Buffs tab shows `Spell #17` and `Spell #871` on the fixture, where `spells.json` names them; the aura names from the log should resolve through the build's names as the cast and damage rows do.
- **B3.** The fixture build's Rankings mode throws on an unreachable rankings API in the capture spec (a console error fails it). Engineering confirms; the mode retires to a link under R3b either way.

## 12. Phone and desktop

- **Never collapses:** the report's answer figure, the tab's table rows and their primary figure, the death recap, Run on Queries, every RankRow's rank, name, figure and Report link, your rank, the sign-in action on /logs.
- **Collapses:** the bar to its summary line (the stack in place, labelled, 44 px); the tabs wrap to two or three lines below 768 and keep the first seven; the answer stacks under the title below 900; StatTables to icon, name and sub, primary figure (the others into the sub line) below 1024; the chart into a `Disclosure` below 768; the Timeline's plot scrolls sideways inside its frame with the labels sticky; the RankList to the tier list's phone row; the rail follows main below 1024.
- **Breakpoints:** below 768 phone bar; 768 to 1023 one column; 1024 to 1759 main + 340 rail, `--page-max` 1344, seven tabs; 1760 and up main + 400 rail, `--page-max` to 1792, 12 or 13 tabs, the StatTable bar column 280 px, the RankList bar 560 px.
- **Alignment at 2000** (`web-report-summary-2000-fold.png`, `web-rankings-you-2000-fold.png`): column 1,792, margins 104; title, tabs, bar, plate and rail share x 168 and x 1,832.

## 13. Performance and polish

- No layout shift: the report reserves its head, bar, the Plate's heads and six rows, and the rail at their final sizes before the island mounts (the per-breakpoint `min-h` reserves in `[id].astro` are re-measured on the new default view: the Summary drops by about 700 px at 1440); the fight menu and More are overlays; the answer slot keeps one size from skeleton to figure.
- No missing-icon flash: icons carry width and height and draw the socket until loaded.
- Fewer nodes: seven cards, the fight column, three nav rows and the glossary strip go; one Plate per route (`plates.spec.ts` on `/logs`, `/reports`, `/reports/fixture2abcd`, `/reports/sample`, `/rankings`, `/rankings/warden-kelthas`, `/ratings`).
- Island size: `report-island.js` stays at or under its current gzip size; the Menu, LineChart and Spark replace FightSelector, the chart cards and ResourceGraphs' frames.
- Lighthouse (`web/lighthouserc.json` entries for these routes, unchanged budgets): performance at least 0.90 on the report, 0.95 elsewhere; accessibility and SEO at least 0.95; CLS at most 0.05 on `/reports/fixture2abcd.html` at 360 px, 4x CPU.
- The guard: every logs, report and rankings component leaves `surface-guard.allowlist.json` in its lane.
- Fonts: Marcellus loads on every route here (the title is in view).

## 14. Parity regions

Every mock carries `head`, `head-title`, `head-meta`, `head-answer` (where there is an answer), `controls` (where there is a bar), `summary` (the Notice lines: the report's notices, the rankings' Tier list line), `plate`, `slots` (the player view's doll), `footer-note` (the rankings legend), and `section-<slug>` on every Section (`section-players`, `section-damage-taken-by-ability`, `section-damage-over-the-fight`, `section-this-pull`, `section-deaths`, `section-performance-rating`, `section-at-pull`, `section-targets`, `section-estimated-threat`, `section-taunts`, `section-measured-against`, `section-what-the-coach-could-not-check`, `section-upload-a-log`, `section-your-reports`, `section-recent-public-reports`, `section-the-companion`, `section-choose-an-encounter`, `section-fury-warriors`, `section-every-class-and-spec`, `section-forever-score`, `section-performance-rating`, `section-what-each-component-measures`, `section-weights-by-role`, `section-ranked-against-your-spec-or-absolute`, `section-the-cap`, `section-tiers`, `section-how-to-get-your-logs-here`, and the rest by title). The tabs sit inside `head`; the fight menu is an overlay outside regions (it is closed in the parity states). Parity states to start with: `report-summary`, `report-damage-done`, `report-deaths`, `report-sample`, `logs`, `logs-you`, `reports-list`, `rankings-board`, `rankings-you`, `ratings`, at 1440 and 390.

## 15. Acceptance screenshots the implementer attaches

At 390, 1024, 1440, 1920 and 2000 unless width-specific, full page and first screen, each beside its mock at the same scale; the 2000 capture opened first:
1. /logs signed out, signed in with no logs, signed in with logs; a file chosen (title and visibility); an upload failing.
2. /reports signed out, empty, listed, failed; a guild member's tab strip.
3. /reports/fixture2abcd on every tab of 6.2, the player view, the fight menu open, the More menu open, the phone bar open, the Filters popover open, a sorted column.
4. Every state of 6.3, including the night view and a lazy chunk failing.
5. /reports/sample.
6. /rankings, empty, a board signed out, signed in with your row in place, off page, unranked, the Guilds board, the phone bar open, the Filters popover open, an empty board, a failed board.
7. /ratings signed out (and signed in once G9 lands).
8. Focus-visible on a tab, a bar select, a sort head, a table row link, a RankRow link, the drop zone's `choose a file` (1440).
9. The parity spec's region-by-region side-by-sides for the section 14 states.
10. Lighthouse for each route and the island size line.
11. The grep for `CharacterPortrait`, `classIconUrl`, `rounded-[3px]` and the white class dot on the built pages.

## 16. Decisions (PROPOSED, each with a recommendation; ruled after the persona review)

- **L1. /logs's Plate is the next step.** Signed out and with no logs: the upload; with logs: Your reports (upload open below). Recommended. Alternative: a fixed Your reports Plate in every state, which leaves a signed-out plate holding one sentence.
- **L2. /logs leads with the latest report** in the head's answer slot; the shared CharacterHeader band (latest kill) leaves /logs. Recommended (the latest kill is the character page's and the report's). Alternative: keep the band.
- **L3. A sign-in button on signed-out /logs.** The shipped rule was "the page carries no sign-in button". Recommended: one primary `Sign in with Battle.net` in the upload Plate, the drop zone shown disabled, so the visitor sees what they get. Alternative: the shipped sentence only.
- **R1. The report's answer:** the group's DPS for the pull; the viewer's own DPS in gold when their character is in it. Recommended. Alternative: no figure (the shipped head).
- **R2. Fights in the bar's `Fight` select with a menu**; the 300 px fight column retires. Recommended (WCL's fight picker; the column was 680 px of trash on the sample). Alternative: keep a left column from 1760 only.
- **R3. One tab row.** Modes and views fold into tabs; seven tabs below 1760, thirteen from 1760, the rest in More. Recommended. **R3b.** The Rankings mode retires to `Rankings for {boss} ›` (the rankings page owns ranks; one job per page). Alternative: keep Rankings as a tab.
- **R4. The table first, the chart under it.** Recommended (the answer is the table; up to nine rows in the first screen at 1440 against none today). Alternative: WCL's order, chart first at 180 px (the table then starts at y 741).
- **R5. A player row sets Source to that player** (their abilities, their gear at the pull); nested row expansion retires. Recommended (WCL). Alternative: keep in-row expansion.
- **R6. Charts only where the tab's metric is charted** (Summary, Damage done, Healing, Damage taken, Deaths; Threat its own). Buffs, Debuffs, Casts, Interrupts, Dispels, Resources, Mechanics, Rating and Coach draw no damage chart. Recommended.
- **R7. `This pull` Facts at the top of the rail on every tab.** Recommended (the fight's facts stay in view when the head's answer scrolls away).
- **R8. `Not credited` as a muted Tag** with the credit sentence in its title (provenance quiet). Recommended.
- **R9. The chart's Group median line retires** (on a five-player pull it reads as noise along zero). Recommended; ask the persona.
- **R10. The `What do these words mean?` glossary strip retires**; each term's definition moves to its column head's `title` and to /ratings. Recommended.
- **K1. A board's title is its encounter** (`Onyxia`); `Rankings` stays the index's title. Recommended.
- **K2. Region, Phase, Faction and Today only behind `Filters`.** Recommended (the bar stays one line at 1440).
- **K3. A crest on every ranking row**, single-class boards included (one crest language), replacing `showRowCrests`. Recommended.
- **K4. /rankings is an encounter list** (a Plate of encounters) rather than the every-class board opened on the first encounter. Recommended.
- **K5. Your row in place when it is on the page, pinned first under `Your best kill` when it is not** (shipped behaviour, kept). Recommended.
- **T1. /ratings: the sample score in the head's answer and the five parts as the one Plate; prose in open Sections; tiers in the rail.** Recommended.
- **P1. Approve the new primitives** `TabStripMore`, `LineChart`, `Spark`, `Menu`, `DropZone`, `Disclosure`, `EmptyCrest` (next consumers: the guild raid report, the character page's history, the sim's resource lane).
- **D1. Replace fixture2abcd with a real Forever pull** (data lane), so the report's parity mocks show Forever spells, items and icons (G2).

## 17. Questions for the wow-player persona

The persona could not be spawned from this agent. The coordinator runs it on `design/reviews/2026-11-07-logs/after/` (the `compare-logs-*` sheets and the first screens):
1. Report Summary at 1440 (`web-report-summary-1440-fold`): does the head's `136 DPS · Group damage · 1 death` plus the players table answer "how did the pull go" faster than the live page? Would you rather see your own DPS and parse there when you are in the pull (R1)?
2. Fights in a menu from the bar (`web-report-fights`) instead of a column: on a raid night with 12 boss pulls and 40 trash groups, is one click to the menu fast enough, or do you want the pulls always visible (R2)?
3. One tab row with More (R3): are Summary, Damage done, Healing, Damage taken, Deaths, Mechanics and Timeline the right seven to see without More at 1440? Is losing the Rankings mode (a link to the boss's board instead) a loss?
4. Table above the chart (R4): in Warcraft Logs you read the graph first. Here the graph is under the table. Right for you?
5. Clicking a player opens their abilities and the gear they wore (`web-report-player`). Is that what you expect, or do you want the row to expand in place (R5)?
6. Deaths (`web-report-deaths`): the last hits with health after each and what you were holding. Anything missing from a death recap you would use (heals received, cooldowns available)?
7. Coach (`web-report-coach`): does the review read as a coach, and is `Measured against` and `What the coach could not check` in the rail the right weight?
8. Rankings with your character (`web-rankings-you`): does `14th of 212 Fury Warriors` in the head plus your marked row tell you where you stand? Crests on every row of a single-class board: noise or right (K3)? Are Executed and Length the columns you look at after DPS?
9. /logs signed out (`web-logs`): would a sign-in button next to a disabled drop zone get you to upload, or does it read as a wall (L3)?
10. /ratings: is the page something you would read, and does the sample score at the top make the Forever Score clear?
11. Phone (`web-report-summary-390`): the table starts at y 720 under a two-line title, the answer and two lines of tabs. Is that too much head on a phone? What would you drop first?
12. Anything a Classic raider expects from a log report that is missing: per-phase damage, a comparison to your previous kill, consumables on the pull, a WCL export link?

Verdict asked for: SHIP the mocks to build, or the list of changes.
