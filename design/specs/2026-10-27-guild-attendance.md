# Guild attendance and the raid-night report (Lane G2): experience spec

Author: ux-designer, 2026-10-27. Product design (binding): `docs/superpowers/specs/2026-10-26-guild-suite-design.md` Lane G2, §3, §5. Page it extends (binding grammar, chips, band, `done` state): `design/specs/2026-10-26-guild-raids.md` (G1). Base page: `design/specs/2026-10-04-guild-page.md` §4.B, §4.C, §4.E. Read-only pattern: `design/specs/2026-10-25-officer-tools-gate.md` §3.2–3.4. Embed grammar: G1 §6.2. Mock: `design/mocks/gen_guild_attendance.py` → `design/mocks/renders/guild-attendance-*.html`; it imports G1's generator helpers, so the two boards cannot drift. No new token, no new control (tenet 10). Every token, type step and chip below is from G1 §3 and §3.6.

## 1. The one question, and the first screen

| Viewer | Question | Answer visible with no scroll, 360 x 640 and 1280 x 800 |
|---|---|---|
| Officer, a `done` raid page | "Who came, and how did the night go?" | The band's **night line** (§3.1) `22 of 23 came · 1 kill · 3 wipes · 14 deaths · 2 h 41 min`, the match line under it, then the Report's kill order. 360: the night line ends above y 420. 1280: the first kill-order row ends above y 620. |
| Member, a `done` raid page | "Was I counted, and what did I get?" | The night line, then `Your night` (§3.3): the outcome sentence, pulls, deaths, loot, parse move. 360: the first `Your night` outcome sentence starts above y 620 (amended 2026-10-27, persona round). |
| Officer, an ended raid with no log | "Why is there no report?" | `Ended · no log yet` and `Match a log`, inside the band at every width. |
| Member or officer, Roster tab | "Who shows up?" | Each row's `Came 7 of 8 · 26 of 30` chip (§3.6), sortable by the existing `Attendance` sort chip. |
| Raid leader, Readiness tab | "Is the raid I am about to pull ready?" | The toggle `Next raid · Thu 10 Dec`, on by default, and the board cut to that raid's In and Late roster (§3.7). |

## 2. Reference, and what we borrow

- **Warcraft Logs' report summary** (the Fights list, the Deaths pane). Borrowed: kill order with each boss's kill time, wipes before it, deaths counted per raider. Not borrowed: per-ability tables (our report pages and the Raids tab expandable own pull detail: one job per page).
- **Details!'s Deaths segment.** Borrowed: one row per raider, class-coloured name, a bar to the count, most first.
- **The attendance sheets guilds keep from Raid-Helper and RCLootCouncil.** Borrowed: one row per raider with signed state against outcome; `Late` and `Excused` as their own marks; a figure over a window. Not borrowed: a red no-show column (R-A3).
- **Our G1 raid page**: band, chips, panel grammar, roster group heads (`label`, mono count), clipboard pattern, 10 s poll.
- **Our readiness board** (§4.E): its pills (`Stocked`, `Short`, `Not shared`), worst-first sort, and the roster's toggle filter buttons (`FILTER_BTN` / `FILTER_BTN_ON` in `GuildRosterTable`).

## 3. Layout, region by region

Reused: everything G1 §3 lists, plus `ItemHover` (icon, quality-coloured name, real tooltip) for loot, `FILTER_BTN`/`FILTER_BTN_ON`, the readiness pill recipe (`consumablesKind`), `nightsLabel`, `CharacterGuildLine`. New files, no new look: `guild/RaidMatchLine.svelte`, `guild/RaidMatchPicker.svelte`, `guild/RaidYourNight.svelte`, `guild/RaidReport.svelte` (children `RaidReportKills`, `RaidReportDeaths`, `RaidReportConsumables`, `RaidReportLoot`, `RaidReportMoves`), `guild/RaidAttendance.svelte`; copy in `web/src/lib/guild/attendance-copy.ts`.

### 3.1 Band on a `done` or ended raid (amends G1 §3.2 line 5 and the `done` line)

- **Night line** replaces the count line once `night` is present. The 8 px bar (160 px desktop, 120 phone) is filled `came / signed`. Then mono 15 px `--strong` `22 of 23 came`. Then 13 px muted ` · 2 joined · 1 kill · 3 wipes · 14 deaths · 2 h 41 min` in the body face, only the figures in mono; ` · {j} joined` shows when any joined, and no clause breaks inside (`2 h 41 min` stays whole). The bar sits on the count's line on phone, and the line is two lines at most (amended 2026-10-27, persona round). `came` counts present plus late among characters signed In or Late. It never names anyone. Members and officers both see it.
- **Match line** (`RaidMatchLine`), 13 px, under the night line. Line 1 is G1's `Logged as a raid night.` + `Open the night`; the link now targets the night expanded, `{guildPath}#raids-{report_id}`. Line 2, muted: `Log uploaded Thu 10 Dec, 11:42 pm by Kraggor · matched by time`. A hand match reads `· matched by an officer`. ` (10:42 pm guild time)` follows the time when offsets differ. Officers: then `Not this log`, a 44 px underlined `--text` text button.
- **Ended, no log** (status `planned` or `locked`, past its end, `match` null). Chip `Ended` (G1). G1's `This raid has ended.` becomes `Ended · no log yet`, 14 px `--text`. Within end + 6 h, ` We match the log when it is uploaded.` follows in muted. After that, ` Upload the log, then match it here.` follows, `Upload the log` linking to `/upload`. Officers get `Match a log` (`SECONDARY_BUTTON_FIXED px-4`): on the same row from `md`, its own full-width row on phone. The count line stays G1's muted started count.
- **No officer row on `done`** (amended 2026-10-27, persona round): the G1 officer row (`Edit`, `Lock sign-ups`, `Copy the sign-up list`, `Cancel raid`) is not drawn on a `done` raid at any width; `Not this log` in the match line is the band's only officer control, and the officer-read-only gate sentence shows once, under the Report header. An ended raid with no log keeps `Edit` alone in the officer row, beside the band's `Match a log`. On `done` and ended raids the phone when line takes two lines: `Thursday 10 December`, then `8:00 pm to 11:00 pm · ended 1 h ago`.
- **Clear confirm** replaces line 2's button: `Clear the match? This raid's attendance is removed until another log is matched.`, then `Clear the match` (G1's `--wipe` outline cancel recipe) and `Keep it`. Focus goes to `Keep it`; Escape acts as `Keep it`. On clear: status returns to `pre_done_status`, the band shows `Ended · no log yet`, and the auto-matcher never picks that report for this raid again.

### 3.2 Match picker (`RaidMatchPicker`, officers)

Opens in place under the band's match row, in the strip's panel (`bg-raised`, `border-line`, `rounded-panel`, the top highlight, padding 16, 20 from `md`). Focus moves to the selected row. Sub-head `Match a log`, 15 px `section-title`. Helper 13 px muted: `{Zone} logs this guild uploaded in the last 14 days.` When the zone's `log_zone` is not yet known: `Logs this guild uploaded in the last 14 days. {Zone}'s log name is not known yet, so every zone shows.`

Rows form one `role="radiogroup"`, `aria-label` `Logs to match`. Desktop row 56 px, grid `24px minmax(0,1fr) 200px 140px 120px`, gap 12. Phone: at most three lines and 88 px (amended 2026-10-27, persona round): zone and `Closest`, never broken; `Uploaded` (muted) and the upload time, mono 13 px, with the duration right-aligned on the same line (amended 2026-10-27, final round); then `by Kraggor · 4 pulls · 1 kill`, 12 px muted (a disabled row's matched-elsewhere note is a fourth line). Cells: a 16 px radio in the site checkbox's gold; zone 14 px 600 `--strong` over `by Kraggor` 12 px muted; uploaded time after a muted `Uploaded`, mono 13 px, with guild time muted beside it when offsets differ (the candidates carry no pull time yet) (amended 2026-10-27, final round); `9 pulls · 1 kill` 13 px; duration mono 13 px `2 h 41 min`. Order: nearest to the raid's window first. The nearest carries `Closest` (`pill-site`) and is preselected. A report matched to another raid stays listed, `aria-disabled`, opacity 50, its second line `Matched to {Title} · {Thu 3 Dec}. Clear it there first.` with the title linked. Hover `--card-top`; focus-visible 2 px gold at −2 px; arrows move, Space selects.

Actions: `Match this log` (`PRIMARY_BUTTON_FIXED px-4`, the page's one fill while open) and `Cancel` (44 px underlined text button). A 20 px status line. Empty: `No {Zone} logs uploaded in the last 14 days. Upload the log, then match it here.` and `Cancel` only; the helper is hidden (amended 2026-10-27, persona round). At most 20 rows.

### 3.3 Your night (`RaidYourNight`, in the strip's slot on a `done` raid)

On `done` it replaces `RaidSignupStrip` for a viewer with at least one own character signed up or in the log. Officers do not get it: their rows sit in the full sections with the own-row wash `rgba(229,185,85,.06)`. Panel: the strip's. Sub-head `Your night`. One row per own character, the header character first, divided on `--line-soft`. Desktop min 68 px, grid `minmax(180px,1fr) minmax(0,2fr) minmax(160px,auto)`, gap 16. Phone: stacked, at most 172 px.

1. Who: `ClassCrestRing` 32, the name in Cinzel 14 class colour, the spec 12 px muted.
2. The outcome sentence, 14 px `--text` (§6.1), for example `Signed In as DPS. Present from the first pull, 11 of 11 pulls.` Under it, 13 px muted facts joined by ` · `, words in the body face and figures in mono: `2 deaths` (`0 deaths` too, for a character in the log) (amended 2026-10-27, persona round); `DPS 291 on Onyxia, up 53 since Thu 3 Dec` (mono figures, `up 53` in `--kill`, `down 12` stays muted); `Excused` when set. A character not in the log has no deaths and no move fact; only `Excused` shows, when set (amended 2026-10-27, persona round).
3. Loot: each award an `ItemHover` (20 px icon, quality-coloured name, 13 px). None: `No loot this raid.` muted.

A character waiting for approval keeps G1's muted row. A viewer with no own row gets no panel; the Report follows the band.

### 3.4 Report (`RaidReport`, `<section id="report">`)

Header row `flex items-center justify-between gap-3`: h2 `Report` 18 px. Right, officers only: `Copy for Discord` and `Post to Discord`, both `SECONDARY_BUTTON_FIXED px-4`; the report is read, so no fill. `Post to Discord` exists only when `discord.connected`. Once posted it is replaced by 13 px muted `Posted to Discord at 11:58 pm.` and never posts twice (R-A6). With no webhook, G1's line `Connect a Discord channel in Settings to post raids there.` stands in its place. Officer-read-only: `Post to Discord` is `disabled`, the G1 gate sentence sits under the header, `Copy for Discord` stays live.

Body from 1024: grid `minmax(0,1fr) minmax(0,1fr)`, gap 32. Left column: Kill order, then Deaths. Right column: Loot, then Biggest changes. Consumables spans both underneath. Below 1024: one column, in the order Kill order, Loot, Deaths, Biggest changes, Consumables. Each block: a 15 px `section-title` sub-head with its mono muted count, 12 px gap, rows divided on `--line-soft`.

- **Kill order** (members and officers). One 44 px row per named boss in kill order: mono 13 px muted index `1`; boss 14 px 600 `--strong`; kill time (the kill's duration) mono 14 px `--text` `1:42`; 13 px muted `after 3 wipes`, or `first pull` in `--kill`; then ` · 2 deaths`. Each row links to its fight on the report page, hover `--card-top`. A named boss never killed reads `Not killed · 4 wipes`, its name in `--muted`. Pulls with no published encounter (Barrow Deeps, Hyjal Summit today) take one line: `7 pulls with no named encounter published: 5 kills, 2 wipes.` (guild spec §12.1).
- **Deaths** (officers; a member's own count is in `Your night`). The sub-head's count is the deaths, as the night line counts them; `Show all` counts raiders (amended 2026-10-27, persona round). Below 768 a row is two lines, min 52 px: crest, the whole name and the count; then the bar (`flex-1`, max 160) and the split, never wrapped. From 1024 to 1279 the bar is at most 96 px. One 40 px row per raider who died, most first, then the game's sort: crest 20; name 13 px 600 class colour; an 8 px bar on the `--line-soft` track with `--muted` fill, never `--wipe` (R-A3), max 160 px; mono 13 px `--strong` `4`; 12 px muted `3 on wipes · 1 on kills`. The first 10 show, then G1's underlined 44 px text button `Show all 14 raiders` expands in place. None: `Nobody died.`
- **Consumables** (officers). Rows only for characters in the log (amended 2026-10-27, persona round). From 768 a 24 px head row, `label` 11 px muted on a `--line-soft` rule, puts `Before` over column 2 and `In the log` over column 3; below 768 each stacked value takes an inline 11 px lead-in. The grid keeps its three columns without buff data, so `Before` holds its x; the pill fills its 120 px column. Desktop grid `minmax(180px,1fr) 120px minmax(0,1.5fr)`, rows 44 px: who; **Before**, the readiness pill as the board read it at the raid's cutoff (`Stocked` ok, `Short` warn, `Not shared` muted); **In the log**, `.pill pill-sample` chips for what was used (`Flask`, `Food`, `Potions 3`) and muted text for what was not (`No flask`). Worst first: Short with nothing used, then Short, then the rest. First 10, then `Show all 23 raiders`. When the log has no buff events (`log_has_buffs` false) the In the log column is dropped and one line sits under the sub-head: `This log carries no buff data, so only the readiness check shows.`
- **Loot** (members and officers; the Loot tab already shows awards to members). Grouped by boss under a `label` head with mono count (`ONYXIA 6`); a one-boss raid drops the boss head, and a count of 0 is left out (amended 2026-10-27, persona round). Each award carries its icon from the loot table (the Loot tab's), the client's own art (`Item.IconFileDataID`), quest items included; an item that truly has none shows the `/bis` neutral tile with the name's letters, never the red question mark (amended 2026-10-27, final round), and its tooltip opens with the award's own facts. From 768 the row is a grid `minmax(0,1fr) 16px minmax(0,1fr) auto`, `by {officer}` right-aligned; below, the item on line 1, then the arrow, awardee and `by` on line 2. One 44 px row per award: `ItemHover` (28 px icon, quality name 14 px), a muted arrow, crest 20 and the class-coloured name, 12 px muted `by Bow Jackzon`. Sorted by quality, then name. None: officer `No loot awarded for this raid. Award it on the Loot tab.` (`Loot tab` links to the Loot tab); member `No loot recorded for this raid.`
- **Biggest changes** (officers; a member's own is in `Your night`). Exactly the five largest absolute moves in a raider's best `metric_dps` or `metric_hps` on the same boss, against the previous `done` raid of the same zone. Row 44 px: crest 20; class-coloured name; 13 px muted `on Onyxia`; mono 14 px `DPS 238 → 291`; mono 13 px `+53` in `--kill`, or `−12` in `--muted`. Sub-head `Biggest changes since Thu 3 Dec`. No previous raid of the zone: the sub-head `Biggest changes` stays, over `First {Zone} raid logged. Changes show from the next one.` (amended 2026-10-27, persona round) From 768 the row is a grid `minmax(0,1fr) auto auto 56px`; below, the name and the delta on line 1, `on Onyxia · DPS 238 → 291` on line 2.

### 3.5 Attendance (`RaidAttendance`, officers, `<section id="attendance">`, after the Report)

h2 `Attendance`. Under it, the summary in 14 px: `22 of 23 came · 1 late · 1 not in the log · 2 more joined`; a zero clause is left out; ` · 1 excused` is added when any row is excused.

Groups, in this order, each with a G1 group head: `Signed up, not in the log`, `Late`, `Joined from the bench or without a sign-up`, `Present`, `Bench and out`. Within a group: role, then class in the game's order, then name. An empty group is left out. `Late` holds only characters signed In (amended 2026-10-27, persona round): a character signed Late who is first seen after the first pull plus 15 minutes did what they said and is `Present`; their `Your night` sentence still reads `Signed Late as {Role}. First seen {m} min after the first pull, {n} of {t} pulls.` A Bench or unsigned character first seen after the first pull plus 15 minutes is still `Joined`, never `Late`.

**Full rows** (the first three groups). Desktop grid `minmax(200px,1.4fr) 132px minmax(180px,1fr) 120px 112px`, gap 16, min 52 px. Cells: who (crest 24, name 14 px 600 class colour, spec 12 px muted); the signed chip, G1 §3.6's viewer chips (`In · DPS`, `Late · Healer`, `Bench`, `Out`, `Not signed up`); outcome 13 px (`Not in the log` in `--muted`, `Late, first seen 22 min after the first pull`, `Present`, `Came as Grimtotem` under R-A2); pulls `9 of 11`, figures mono 13 px, with a 48 px bar, under one `Boss pulls` column head from 768 (a phone lead-in below) (amended 2026-10-27, persona round); **Excused**, the site's gold checkbox inside a 44 px label `Excused`, on `no_show` and `late` rows only. Phone: line 1 who with the chip right; line 2 outcome and pulls; line 3 the checkbox; at most 148 px.

**Compact rows** (`Present`, `Bench and out`). From 1024 a `repeat(3, minmax(0,1fr))` grid of 36 px rows: crest 24, name, and mono muted `11/11` at the right edge. 768–1023 two columns; phone one column of 44 px rows. A Present row that missed pulls shows its `8/11` in `--text`. On phone the `Present` group head carries `Boss pulls` at its right (amended 2026-10-27, final round). A `Bench and out` row shows `Bench` or `Out`, muted, in place of a pull count (amended 2026-10-27, persona round). A `Came as` row in Present carries `as {Alt}`, 12 px muted, after the name.

### 3.6 The member record (roster row, loot line, character guild line)

- **Roster chip** (`rosterStatChips`). In raid mode (`attendance_source = "raids"`) it reads `Came 7 of 8 · 26 of 30 raids`. The windows are the guild's last 8 and last 30 logged raids. The numerator is raids attended. The denominator is raids signed In or Late, or attended; excused raids are left out. Fewer than 8 raids: `Came 3 of 3 raids` alone. Under the roster, once, 12 px muted: `Came: raids attended of raids signed up for or attended, over this guild's last 8 and 30 logged raids.` Report mode keeps today's `5 of 8 nights` and no footnote. The `Attendance` sort uses the 8-raid ratio.
- **Loot line** (`attendanceLabel`): raid mode `7 of 8 raids`; report mode unchanged.
- **Character guild line** (`CharacterGuildLine`, own characters, from `/me`): raid mode appends ` · came 7 of 8 raids`, 13 px muted, figures mono. Otherwise nothing is appended.

### 3.7 Readiness tab, the next-raid filter (amends guild spec §4.E)

When a `planned` or `locked` raid exists, the filter row gains, first, the toggle `Next raid · Thu 10 Dec` (`FILTER_BTN`, 44 px phone, 32 px desktop; on: `FILTER_BTN_ON`; `aria-pressed`). It is on by default. On: rows are that raid's In and Late sign-ups. A 13 px muted line under the filter row says `Showing the 23 signed In or Late for Onyxia's Lair, Thu 10 Dec, 8:00 pm.`, with ` (9:00 pm guild time)` when offsets differ and the title linked to the raid page. The Last synced cell judges freshness against the raid: an export synced before `starts_at` minus 7 days reads `Synced 9 days before the raid` in the warn pill and counts as one failing check (one constant, `READINESS_FRESH_DAYS = 7`). A signed character with no export: `No export yet`, muted, sorted last, and `—` from Gear gap to Item level. The note under the board names the order: `Sorted by the most checks to fix, then the biggest gear gap.` (amended 2026-10-27, final round) Off: today's board. No raid planned: no toggle.

## 4. States

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| `Match a log`, `Not this log`, `Copy for Discord`, `Post to Discord` | G1 outline or text | G1 | 2 px gold, 2 px offset | `--card-top` | officer-read-only (`Copy` stays live) | `Matching…`, `Clearing…`, `Posting…` | n/a | 13 px `--wipe` line under the row |
| Picker row | §3.2 | `--card-top` | 2 px gold, −2 px | radio filled | matched elsewhere | 3 `Skeleton` rows of 56 px | §3.2 empty | `These logs did not load.` + `Try again` |
| `Excused` | unticked | border `--line-warm-strong` | gold ring | ticked | officer-read-only | row `aria-busy` | n/a | reverts; the row's status line |
| Kill, loot, death rows | §3.4 | `--card-top` on links | 2 px gold, −2 px | `#171d2b` | n/a | `Skeleton` at ready height | §3.4 lines | section `LoadError` |
| `Show all {n} raiders` | underlined `--text` | `--gold-hover` | gold ring | n/a | n/a | n/a | absent at 10 or fewer | n/a |
| Readiness toggle | off look | border `--line-warm-strong` | gold ring | on look | n/a | board skeleton | `Nobody signed up for this raid yet.` | board `LoadError` |

Touch and keyboard parity: every control is 44 px tall on phone; nothing is hover-only; `ItemHover` opens on tap and on focus. Tab order on a `done` page: back link (amended 2026-10-27, persona round), match line (`Open the night`, `Not this log`), tabs, Your night, Report header buttons, kill rows, loot items, `Show all`, attendance checkboxes.

**Page states.** Loading a `done` raid: G1's band reserve plus the night line (22 px) and the match line (40 px); `Your night` reserves one row per stored own character (`fs.guildRaid.rows`); the Report `min-h-[520px]` desktop and `min-h-[900px]` phone; Attendance `min-h-[420px]` for a cached officer. **Poll:** G1's 10 s poll keeps running on an ended raid until end + 6 h, so a log matched by the parse turns the page `done` in place. That swap is the one allowed height change; an sr-only polite live region says `The raid's log was matched.` **Re-match**: attendance and report recompute; excused flags carry over by character.

**Races.** 409 `report_matched` on `Match this log`: the row turns matched-elsewhere and the focused status line says `That log was just matched to another raid.` 403 `guild_plan_required`: the gate spec §4 handling and `Not saved. This guild's officer tools have ended.`

**Copy for Discord** writes plain text in guild time, at most 12 lines, in this order: `Onyxia's Lair · Thu 10 Dec · Olympus XXVII`; `1 kill, 3 wipes, 2 h 41 min. 22 of 23 came.`; `Kills: Onyxia 1:42 (after 3 wipes)`; `Most deaths: Pyrewisp 4, Windtalon 3, Vexlash 2`; one `Loot: {Boss}: {Item} to {Name}; ...` line per boss; `Biggest changes: Pyrewisp DPS +53, Lightbrand HPS +41` (rises only); `Report:` and the raid page's absolute URL with the report anchor. No absentee names (R-A3). `Copied` for 4 s. Failure: `Could not copy. Select the report and copy it.` and an 8-row read-only textarea, selected.

## 5. Hierarchy

- `done` raid page: the night line is the answer. `Your night` (member) or the Report (officer) is the primary region. No filled button unless the picker is open.
- Gold once per region: `Open the night` in the band, `Closest` in the picker. `--kill` only on `first pull`, rises and `Done`. `--wipe` only on errors and the clear confirm.
- Secondary is visibly secondary: uploader, matched-by, `on wipes / on kills` and `by {officer}` are 12–13 px muted.

## 6. Copy, verbatim (`attendance-copy.ts`; functions take the data; no number typed into prose)

### 6.1 On the page

| Where | String |
|---|---|
| Night line | `{came} of {signed} came` + ` · {j} joined · {k} kill(s) · {w} wipe(s) · {d} death(s) · {duration}`; a zero clause is left out (amended 2026-10-27, persona round) |
| Match line | `Logged as a raid night.` + `Open the night`; `Log uploaded {when} by {name} · matched by time` / `· matched by an officer`; `Not this log` (`Clearing…`); `Open the raid` (a night row's link to its raid) |
| Ended, no log | `Ended · no log yet`; ` We match the log when it is uploaded.`; ` Upload the log, then match it here.`; `Match a log` |
| Clear confirm | `Clear the match? This raid's attendance is removed until another log is matched.`, `Clear the match`, `Keep it` |
| Picker | `Match a log`; `{Zone} logs this guild uploaded in the last 14 days.`; `Logs this guild uploaded in the last 14 days. {Zone}'s log name is not known yet, so every zone shows.`; `by {name}`; `{p} pulls · {k} kill(s)`; `Closest`; `Matched to {Title} · {date}. Clear it there first.`; `Match this log` (`Matching…`), `Cancel`; `No {Zone} logs uploaded in the last 14 days. Upload the log, then match it here.`; `These logs did not load.`; `That log was just matched to another raid.` |
| Your night sentences | present `Signed {State} as {Role}. Present from the first pull, {n} of {t} pulls.`; late `Signed {State} as {Role}. First seen {m} min after the first pull, {n} of {t} pulls.`; came as `Signed In on {Name}; came as {Alt}, {n} of {t} pulls.`; no-show `Signed In as {Role}. You are not in this log. If you were there under another name, tell an officer.`; unsigned `Not signed up. In the log for {n} of {t} pulls.`; bench joined `Signed Bench. Joined for {n} of {t} pulls.` |
| Your night facts | `Your night`; `{d} death(s)`; `{Metric} {after} on {Boss}, up {x} since {date}` / `down {x}`; `Excused`; `No loot this raid.` |
| Report | `Report`; `Copy for Discord` (`Copied`); `Post to Discord` (`Posting…`); `Posted to Discord at {time}.`; `Kill order`, `Deaths`, `Consumables`, `Loot`, `Biggest changes since {date}` |
| Kill order | `after {w} wipe(s)`; `first pull`; ` · {d} death(s)`; `Not killed · {w} wipe(s)`; `{p} pulls with no named encounter published: {k} kills, {w} wipes.` |
| Deaths | `{w} on wipes · {k} on kills`; `Show all {n} raiders`; `Nobody died.` |
| Consumables | `Before`, `In the log`; `Stocked`, `Short`, `Not shared`; `Flask`, `Food`, `Potions {n}`, `No flask`, `No food`, `No potions`; `This log carries no buff data, so only the readiness check shows.` |
| Loot | `by {officer}`; `No loot awarded for this raid. Award it on the Loot tab.`; `No loot recorded for this raid.` |
| Changes | `on {Boss}`; `{Metric} {before} → {after}`; `+{x}`, `−{x}`; `First {Zone} raid logged. Changes show from the next one.` |
| Attendance | `Attendance`; `{came} of {signed} came · {l} late · {n} not in the log · {j} more joined · {e} excused`; `Signed up, not in the log`, `Late`, `Joined from the bench or without a sign-up`, `Present`, `Bench and out`; `Present`, `Not in the log`, `Late, first seen {m} min after the first pull`, `Came as {Alt}`; `Excused` |
| Roster | `Came {p} of {n} · {p30} of {n30} raids` (one unit for both windows, amended 2026-10-27, final round); `Came: raids attended of raids signed up for or attended, over this guild's last 8 and 30 logged raids.`; loot line `{p} of {n} raids`; guild line ` · came {p} of {n} raids`; `No raids yet` (a chip with no raids signed up for or attended) |
| Readiness | `Next raid · {Thu 10 Dec}`; `Showing the {n} signed In or Late for {Title}, {Thu 10 Dec, 8:00 pm}.`; `Synced {d} days before the raid`; `No export yet`; `Nobody signed up for this raid yet.` |
| Errors | `That did not save; try again.` (buttons), `That did not save. Try again.` (rows), `Could not post. Try again.`, `This report did not load.`, `Not saved. This guild's officer tools have ended.`, `Could not copy. Select the report and copy it.` |
| Live region | `The raid's log was matched.` |
| Changelog | `Added attendance and raid-night reports. A raid's log is matched when it is uploaded, and the raid page shows who came, kill times, deaths, loot and the biggest changes since last time.` |

### 6.2 In Discord: `report_ready` (G1 §6.2 grammar)

Title `Report: {Title} · Thursday 10 December`, linking to the raid page's report anchor. Colour `8377482` (`#7fd48a`, the Done mark). Inline fields: `Kills` (one line per kill, `Onyxia 1:42`, at most 6, then `+{n} more`); `Wipes` `3`; `Came` `22 of 23`; `Loot` `6 items`. Description: `Kill times, loot and the biggest changes on Forever Sixty.` No raider names (R-A3). It posts only from an officer's `Post to Discord` (R-A6) and is never edited.

## 7. Phone and desktop

**Never collapses**: the night line and its figures; every outcome word; every kill time; every pull count; every death count; every item name; the excused state; the gate sentence; the roster footnote. **Collapses**: below 1024 the Report goes to one column; below 768 attendance and consumables rows stack, picker rows take two lines, and the Report header buttons wrap under the h2. At 2000 everything stays in the centred 1344 column. Caps: picker 20 rows; deaths and consumables 10 before `Show all`; changes exactly 5; kill order every kill. Heights: night line 22 desktop, at most 44 phone; match line at most 66 desktop, 110 phone; Your night row 68 / 172; full attendance row 52 / 148; compact row 36 / 44.

## 8. Rulings

- **R-A1 Members see the guild's night, not each other's rows.** Kill order, loot and the night line's counts are member-visible (the Raids and Loot tabs already are). Deaths, consumables, changes and attendance rows are officer views. A member sees their own in `Your night`. The roster's aggregate chip keeps today's roster visibility. A read-only officer still reads everything; only writes are disabled.
- **R-A2 Came on another character counts.** A signed character absent from the log, whose account has another character of this guild in it, reads `present` or `late` with `came_as`. This prevents the commonest false no-show. It is a product amendment: the attendance row gains `came_as`.
- **R-A3 No shaming.** The word "no-show" never renders. `Not in the log`, death bars and falls are `--muted`. Discord text carries counts, never absentees.
- **R-A4 The report replaces the strip, not the roster.** On `done` the disabled strip goes; Notes and the G1 roster stay below Attendance, closed.
- **R-A5 First seen is relative.** Combat logs carry no absolute clock, so first seen reads in minutes after the first pull and needs no time zone.
- **R-A6 `report_ready` posts on an officer's press, once per raid.** Every Discord post is a choice (product §2).
- **R-A7 One log, one raid.** A report matched elsewhere shows, disabled, in other pickers.
- **R-A8 Readiness defaults to the next raid** when one is planned, because the board is a pre-pull tool.

## 9. Performance, contract and acceptance

Limits: the report rides the raid view; the picker fetches on open; `ItemHover` loads its panel lazily. The Lighthouse `/guild/` row is unchanged (the audit is signed out and draws nothing new). Playwright CLS on a `done` raid with a warm reserve, at 390 and 1440: at most 0.02; the poll's announced `done` swap is measured on its own.

**Contract this spec assumes.** The raid view gains:
- `pre_done_status: "planned"|"locked"|null`; `match: null | {report_id, zone_name, uploaded_at, uploaded_by_name, pulls, kills, duration_ms, how: "auto"|"hand"}`.
- `night: null | {signed, came, late, not_in_log, joined, excused, kills, wipes, deaths, duration_ms}` (everyone).
- `report: null | {kill_order [{encounter_id, name, fight_index, killed, kill_ms|null, wipes_before, deaths}], unnamed {pulls, kills, wipes}|null, deaths [{character_key, name, class, on_wipes, on_kills}], consumables {log_has_buffs, rows [{character_key, name, class, before: "stocked"|"short"|"unknown", flask: bool|null, food: bool|null, potions: int|null}]}, loot [{encounter_id, boss, awards [{item_id, item_name, quality, character_key, name, class, by_name, at}]}], moves {previous {id, title, starts_at}|null, rows [{character_key, name, class, metric: "dps"|"hps", encounter, before, after}]}, posted_at|null}`. Officers (rank officer or leader, with or without the plan — R-A1) get every row; a member gets `kill_order`, `unnamed`, `loot`, and `deaths`, `consumables.rows` and `moves.rows` filtered to their own characters. A never-killed boss counts `deaths` over every one of its pulls.
- `attendance: null | [{character_key, name, class, spec, signed: "in"|"late"|"bench"|"out"|null, role|null, outcome: "present"|"late"|"absent"|"no_show", came_as|null, first_seen_offset_s|null, pulls_present, pulls_total, excused, own}]` (officers all rows; a member their own).
- Match candidates, officers: `GET` on the raid's `match-candidates` route answers `{zone_known, rows [{report_id, zone_name, uploaded_at, uploaded_by_name, pulls, kills, duration_ms, closest, matched_raid {id, title, starts_at}|null}]}`.
- Three officer actions behind `CanForGuild(officer_views)`, each answering the raid view: set the match (`PUT` on the raid's `match` route, body `{report_id}`); clear it (`DELETE` on the same route); excuse (`PUT` on the raid's `attendance` route plus the character key's three segments, body `{excused}`). Plus posting (`POST` on the raid's `report/post` route) answering `{posted_at}`, 409 `already_posted` on a second press, 502 `post_failed` when the post reached no channel (the claim is released, so the officer can press again). Errors: 409 `report_matched`, 409 `conflict` (raid not ended), 422 `invalid`, 403 `guild_plan_required`, 403 `forbidden`.
- Readiness: the board read with `raid={id}` (a planned or locked raid of this guild; anything else answers 404 `not_found`) adds `raid {id, title, starts_at, cutoff_at}` and, per row, `synced_before_cutoff` and `no_export`; the unfiltered read adds `next_raid {id, title, zone_name, starts_at}|null`.
- Roster: top-level `attendance_source: "raids"|"reports"`; rows gain `attendance_30 {present, nights}` in raid mode. `Attendance{present, nights}` keeps its name and shape, but in raid mode `nights` is per character, so `loot.go` (~182–188, which takes the last row's nights as the guild's) must read each candidate's own. `/me`'s `guild` gains `attendance {present, nights}` in raid mode.

**E2E** (`guild-attendance.spec.ts`, API stubbed): an ended raid shows `Ended · no log yet`; the poll turns it `done` and the live region speaks; an officer opens the picker, `Closest` is preselected, a matched-elsewhere row is disabled, `Match this log` matches; `Not this log`, then the confirm, clears; `Excused` ticks and the summary updates; a member sees `Your night` and no Deaths, Consumables, Changes or Attendance rows but their own; the no-show sentence renders and no text node contains "no-show"; `log_has_buffs` false shows its line; no previous raid shows its line; `Copy for Discord` text equals the fixture; `Post to Discord` posts once and the stub receives one `report_ready`; officer-read-only: match, clear, excused and post disabled, copy live; the roster chip in both modes; the readiness toggle on by default with rows equal to the sign-ups; axe clean in every state; every control at least 44 px tall at 360.

**Screenshots** (fixture build; x = required):

| # | State | 360 | 390 | 1440 | 2000 |
|---|---|---|---|---|---|
| 1 | Done raid, officer: night line, match line, Report, Attendance | x | x | x | x |
| 2 | Done raid, member: `Your night` present, and the no-show variant | x | x | x | |
| 3 | Ended, no log: officer and member, inside and after the 6 h window | x | x | x | |
| 4 | Picker: `Closest`, a matched-elsewhere row, empty, unknown log zone | x | x | x | |
| 5 | Clear confirm, focus on `Keep it` | | x | x | |
| 6 | Attendance: every group, excused ticked, a `Came as` row | x | x | x | x |
| 7 | Report with no buff data, no previous raid, no loot | | x | x | |
| 8 | Officer-read-only done raid | | x | x | |
| 9 | Roster chip in raid and report mode with the footnote; character guild line | x | x | x | |
| 10 | Readiness next-raid on (a stale export, no export) and off | x | x | x | |
| 11 | Loading beside ready, CLS printed | | x | x | |
| 12 | `report_ready` embed and the copied text in a staging channel, desktop and phone clients | | | | |
| 13 | Same-scale side-by-side with the mock boards at 1440, plus 1024 and 2000 | | | x | x |
