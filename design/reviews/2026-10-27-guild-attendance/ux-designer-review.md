# Guild attendance and the raid-night report (Lane G2): ux-designer review

Reviewer: ux-designer, 2026-10-27. Spec: `design/specs/2026-10-27-guild-attendance.md` §1–§9 (rows 1–11 and 13; row 12 is not captured). G1 grammar: `design/specs/2026-10-26-guild-raids.md`. Evidence: the PNGs and JSON in this folder only.

Fixture names and figures are not defects. The `Sign in` header in the signed-out stubs is not a defect. The side-by-sides compare layout, not figures.

## Verdict: SHIP WITH FIXES

The report, attendance, picker, clear confirm, read-only gate, roster chip, guild line and readiness toggle are built to the spec. The page reads as a raid-night report. Every must-fix below blocks merge. Each must-fix lands with re-captures of the rows named.

## Rulings on the implementer's flags

- **(a) G1's officer row on a `done` raid: remove it.** On `done` the band draws no officer row at any width. The mock draws none (13-mock-officer-1440). `Edit`, `Lock sign-ups` and `Cancel raid` do nothing on a logged night. `Copy for Discord` replaces `Copy the sign-up list`. The read-only gate sentence then shows once, under the Report header (§3.4), not twice. The §4 tab order becomes: back link, match line, tabs. Finding 1.
- **(b) Notes and the Sign-ups roster open below Attendance: R-A4 stands; they render closed.** Finding 2.
- **(c) Consumables without `Before` / `In the log` heads: must-fix.** Both strings are in §6.1's verbatim copy. Without them the pill column reads as an unlabelled verdict. Finding 3.
- **(d) Empty loot icon squares: must-fix.** The tenet is real game icons everywhere. `Head of Onyxia` is a quest item and is in no class's item file, so the lookup in `RaidLootItem` can never find it. The Loot tab already carries `icon` per item from the build's `items.json` (`api/internal/guilds/loot.go`). Finding 4.
- **(e) `Sign in` in the header: ignored.**
- **(f) CLS 0.0001 warm, but the loading reserves are the spec's numbers: the CLS figure is accepted. The reserves are not.** G1 ruling (b) applies here too: a reserve equals the measured ready render, never a guessed constant. The §4 numbers (band + 22 + 40, the Report minimums) are floors, and the measured ready heights replace them. Finding 8.

## Findings

1. **must-fix**: officer row on a `done` raid (01-done-officer-1440, 01-done-officer-360, 08-read-only-panel-1440/390, 11-ready-officer-1440). Shipped: `Edit`, `Lock sign-ups`, `Copy the sign-up list` and `Cancel raid`, some dimmed. At 360 they add two rows of buttons (~110 px) above the tabs. The read-only gate sentence shows twice. Spec: §5, the night line is the answer and nothing competes with it; the mock draws no row. Fix: the raid page does not render the G1 officer row when the status is `done`. The read-only gate sentence renders only under the Report header. Amend the §4 tab order to back link, match line, tabs.
2. **must-fix**: Notes and Sign-ups open below Attendance (01-done-officer-1440 y 2350–2830, 01-done-officer-360 y 4140–5900, 02-member-present-1440, 13-side-by-side-member-390). Spec: R-A4, "Notes and the G1 roster stay below Attendance, closed". Shipped open, which adds ~480 px on desktop and ~1,700 px at 360. Fix: on `done`, wrap each in a native `<details>`, closed by default. Follow the pattern in `report/RaidCooldowns.svelte`. Each `<summary>` is a 44 px row holding the h2 (`Notes`, or `Sign-ups` with its mono `14/40`) and a 16 px chevron that rotates when the section opens, with a 2 px gold focus-visible ring. The officer selects inside take G1's disabled recipe (`opacity-50`, `cursor-not-allowed`, `disabled`), because sign-ups on a logged night are history.
3. **must-fix**: Consumables column heads (01-done-officer-1440 y 1240–1520, 01-done-officer-report-1440). Spec: §3.4 **Before** / **In the log**, §6.1 copy. Shipped: no heads. Fix: from 768, add a 24 px head row on the block's own grid (`minmax(180px,1fr) 120px minmax(0,1.5fr)`). It holds `Before` over column 2 and `In the log` over column 3, in `label` 11 px `--muted`, with a `--line-soft` bottom rule (as 13-mock-officer-1440). Below 768, put each stacked line behind an inline `label` 11 px lead-in, `Before` and then `In the log`. When `log_has_buffs` is false, only `Before` shows.
4. **must-fix**: loot icons are empty 28 / 20 px boxes (01-done-officer-report-1440/360, 02-member-present-1440, 13-built-*). Spec: §3.4 Loot, `ItemHover` with a 28 px icon. Tenet: real game icons, never placeholders. Fix: the raid view's `report.loot[].awards[]` gains `icon`, resolved server-side from the build's `items.json` exactly as `loot.go` does for the Loot tab. `RaidLootItem` renders that `icon` and drops the class-file `loadItems` lookup. The tooltip must still open for an item outside every class file: pass `ItemHover` a `model` or a class-free lookup, never a null panel. Re-capture with Head of Onyxia's icon and its open tooltip.
5. **must-fix**: the no-show `Your night` contradicts itself (02-member-no-show-panel-1440/360/390). It reads `You are not in this log.` and then `DPS 196 on Onyxia, up 16 since Thu 3 Dec`. Spec: §3.3, facts are deaths, the parse move and `Excused`, all of which come from the log. Fix: when the row's outcome is `no_show`, the facts line renders only `Excused` (when set) and otherwise nothing. Also correct the no-show fixture so its `moves.rows` drops the character.
6. **must-fix**: the first screen misses §1's targets. At 1440 the first kill-order row ends at ~y 680 against "above y 620" at 1280 x 800 (01-done-officer-1440). At 360 the member's first `Your night` row ends at ~y 760 against 640 (02-member-present-360). Fix, desktop: remove the 28 px empty strip between the header and the band (y 80–108, drawn on every guild page). It is shared shell, so run the full web vitest, Playwright and the lhci shard. Then set the Report section's top padding to 32 px, the h2-to-sub-head gap to 16 px and the sub-head-to-rows gap to 12 px (§3.4). Fix, phone: put the night line's 120 px bar on the count's line, as the mock draws (13-side-by-side-member-390). Set the trailing clause proportional (finding 7) so the night line is two lines, ≤ 44 px (§7). Set the when line in two lines: `Thursday 10 December`, then `8:00 pm to 11:00 pm · ended 1 h ago`. This amends G1's three lines on `done` and `ended` only. Amend §1's phone member target to: "the first `Your night` outcome sentence starts above y 620 at 360". Re-capture both widths with a 640 / 620 guide line drawn.
7. **must-fix**: night-line and facts typography (01-done-officer-panel-1440, 02-member-present-panel-1440). The ` · ` separators have no spaces: `came·1 kill·3 wipes`. The whole trailing clause is mono, including the words. The `Your night` facts line is mono end to end. Spec: §3.1, the clause is 13 px muted with figures in mono; §3.3 likewise; §6.1, the ` · ` separator is verbatim. Fix: join with `SEP` (` · `, spaces included). Set the clause in the body face at 13 px `--muted`, and wrap only the numerals (and the figures in `2 h 41 min`) in `font-mono`. Treat the facts line the same way: `up 16` stays `--kill`, and its words go in the body face.
8. **must-fix**: loading reserves do not equal the ready render (11-loading-* beside 11-ready-*). Measured band heights: officer 1440, loading 368 against ready 405; member 390, loading 386 against ready 438; member 1440, loading 368 against ready 358. `Your night` at 390: loading panel ~400 against ready ~228. Fix: after findings 1, 6 and 7, re-measure the ready band per role and layout. Reserve exactly that: the night line row plus the match line's two lines (one row per stored own character for `Your night`, at the measured row height plus the panel padding and sub-head). There is no officer row on `done`. Attach loading beside ready matching to ±1 px, and re-print `cls-390.json` / `cls-1440.json`.
9. **must-fix**: Deaths rows at phone widths (01-done-officer-360 y 1240–1790, 07-bare-report-panel-390). Names truncate to `Pyrewi…`, `Mend…`, `Shado…`. The split `3 on wipes · 1 on kills` wraps to three lines, so rows grow from 40 to ~55 px and stop aligning. Spec: §3.4 Deaths, 40 px rows with a class-coloured name. Fix: below 768 the row takes two lines, min 52 px. Line 1 holds crest 20, the name untruncated, and the mono count right-aligned. Line 2 holds the 8 px bar (`flex-1`, max 160) and then the 12 px muted split, `whitespace-nowrap`.
10. **should-fix**: Deaths split wraps at 1024 (13-side-by-side-officer-1024: `3 on wipes · 1 on / kills`). Fix: from 1024 to 1279 the bar's max width is 96 px, and the split is `whitespace-nowrap`.
11. **should-fix**: Deaths sub-head count (01-done-officer-1440: `DEATHS 11` beside the night line's `17 deaths`). The mock's count is the death total (`Deaths 14` = the night line's 14). Fix: the sub-head count is the sum of deaths. `Show all 11 raiders` keeps the raider count.
12. **should-fix**: Loot and Biggest changes rows are ragged (01-done-officer-1440, 01-done-officer-360). The arrow, awardee and `by Kraggor` drift with the item name's length. At 360, `by Kraggor` drops to its own line, and `+41` / `+16` wrap under the row. Mock: aligned columns (13-mock-officer-1440). Fix, Loot from 768: grid `minmax(0,1fr) 16px minmax(0,1fr) auto`, gap 12, with `by {officer}` right-aligned. Below 768, line 1 holds the item; line 2 holds the arrow, the awardee and `by {officer}`, which is muted and right-aligned. Fix, Changes from 768: grid `minmax(0,1fr) auto auto 56px`, with `on {Boss}`, the metric and the right-aligned delta. Below 768, line 1 holds the crest, the name and the right-aligned delta; line 2 holds `on Onyxia · DPS 238 → 291`.
13. **should-fix**: consumables row for a raider not in the log (01-done-officer-1440: Hexbramble, `No flask No food No potions`). That reports usage for someone who was not there, and R-A3 forbids it. Fix: when the character is not in the log, the In the log cell reads `Not in the log` in 13 px `--muted`. The row sorts after every in-log Short row.
14. **should-fix**: bare report (07-bare-report-panel-1440). The changes empty line `First Onyxia's Lair raid logged. Changes show from the next one.` stands without a sub-head. With buff data dropped, the `Short` pill sits at the far right edge (x ~1150). Fix: the block keeps its sub-head, `Biggest changes` (no `since {date}` when there is no previous raid). The consumables grid keeps its three columns with column 3 empty, so `Before` stays at the same x as the full report.
15. **should-fix**: picker rows at phone widths take four lines (04-picker-closest-panel-360, 04-picker-unknown-zone-panel-390). The duration wraps to its own line. Spec: §3.2 says "two lines, min 64", which cannot hold zone, uploader, time, pulls and duration at 360, so the spec is amended. Fix: at most three lines, at most 88 px. Line 1: zone, `Closest`, and the duration mono right-aligned. Line 2: the upload time in 13 px mono. Line 3: `by Kraggor · 4 pulls · 1 kill`, 12 px muted. The matched-elsewhere note is a fourth line on disabled rows only.
16. **should-fix**: readiness `No export yet` row shows export data (10-readiness-on-panel-1440, Shadowquill: `4 upgrades · +28 DPS`, `BOOTS`, `59 (−2)`). Fix: when `no_export`, cells Gear gap through Item level render `—` in `--muted`. Nudge stays. Correct the fixture to match.
17. **should-fix**: the 1024 mock board is rendered at 1440 wide (13-mock-officer-1024 is 1440 px; 13-side-by-side-officer-1024 is not same-scale). Fix: render the officer board in `gen_guild_attendance.py` at a 1024 viewport and rebuild the side-by-side.
18. **should-fix**: the clear confirm shows no focus ring on `Keep it` (05-clear-confirm-panel-1440/390). Spec: row 5, focus on `Keep it`. Fix: open the confirm from the keyboard (Tab to `Not this log`, Enter) so `:focus-visible` draws the 2 px gold ring, then re-capture. If the ring still does not draw, focus is not moving, and that becomes a must-fix.
19. **nit**: `Your night` omits `0 deaths` (02-member-present). The mock shows it, and it answers "was I counted". Fix: always render `{d} death(s)` for a character in the log, including 0.
20. **nit**: `Bench and out` compact count `0/4` is drawn in `--text` (06-attendance-panel-1440, Frostnettle). Spec: §3.5, only a Present row that missed pulls uses `--text`. Fix: `--muted` in `Bench and out`.
21. **nit**: `LOOT 0` (07-bare-report-panel-1440/390). Fix: drop the count at 0, as G1 ruled for `PLANNED 0`.
22. **nit**: empty picker repeats itself (04-picker-empty-panel-390). The helper and the empty line both name the zone and 14 days. Fix: in the empty state, hide the helper; `No {Zone} logs…` and `Cancel` stand alone (§3.2).
23. **nit**: the guild line breaks inside its figure (09-guild-line-panel-360: `came 7` / `of 8 raids`). Fix: `whitespace-nowrap` on the appended ` · came 7 of 8 raids` clause.
24. **nit**: Kill order row left-runs in a 608 px column (01-done-officer-1440). Fix: right-align the kill time and `after 3 wipes · 2 deaths`, as 13-mock-officer-1440 draws. The index and boss stay left.
25. **nit**: the `Before` pill is content-width (01-done-officer-1440). Fix: the pill fills the 120 px column, as the mock draws.
26. **nit, no change**: the night line's bar is `--muted`, while the mock draws it gold. The shipped colour is correct under §5's "gold once per region" (`Open the night`). The mock is amended to match.
27. **nit, out of lane**: Brightwick's roster crest is an empty ring (09-roster-raids-1440, `Combat in game`). This breaks the one crest language. File it against the roster: a known spec draws its class's ringed crest.

## Checklist (§9)
- [x] Rows 1–11 and 13 at their widths; row 12 skipped as instructed.
- [x] 2000 px: the 1344 column is centred and the band is full-bleed (01-2000, 06-2000, 13-2000).
- [x] Members see no Deaths, Consumables, Changes or Attendance rows (02-*).
- [x] No text node reads "no-show"; `Not in the log` and the death bars are `--muted` (R-A3).
- [x] Picker: `Closest` preselected, matched-elsewhere row disabled with its line, unknown-zone helper verbatim.
- [x] Officer-read-only: `Post to Discord`, `Not this log` and `Excused` disabled; `Copy for Discord` live (08).
- [x] Roster chip in both modes with the footnote; `No raids yet`; guild line appended (09).
- [x] Readiness toggle on by default, stale-export warn pill, `No export yet` last (10).
- [ ] First screen meets §1 at 360 and 1280 (finding 6).
- [ ] Copy verbatim: ` · ` spacing, `Before` / `In the log` (findings 3, 7).
- [ ] Real icons on every item (finding 4).
- [ ] R-A4: Notes and Sign-ups closed (finding 2).
- [ ] Loading beside ready matches ±1 px (finding 8).
- [ ] Every count and name survives 360 (findings 9, 12).

## Re-review

Reviewer: ux-designer, after the persona fix round (`b6f2b9da`, `7ed80843`). Evidence: the re-taken PNGs and JSON in this folder, plus the guild views' source for the shared padding change (no captures exist for the views without a band).

### Ruling: Head of Onyxia's icon

It shows its real game icon, `inv_misc_head_dragon_01`. The red question mark is not acceptable. `inv_misc_questionmark` is the client's missing-texture icon, and `/bis` already refuses it (`GearRow.astro`, `lib/gear/qualities.ts`: "never a red ?"). The premise is also wrong: the icon is in the build's data. `data/builds/1.60.1.70532/icons/inv_misc_head_dragon_01.webp` exists. Only `items.json` carries no `icon` for the quest item, and the loot table in `api/internal/guilds/loot.go` leaves it `""`.
- Fix: fill `icon` for 18422 and 18423 with `inv_misc_head_dragon_01` in `loot.go`. Fill every other `""` in that table (15410, 17966, 18705, 21108) from the client's item data, not by guessing. Make sure the published icon set for the build (`web/public/data/<build>/icons`) includes each file.
- Fallback, when an item really has no icon: use the `/bis` neutral tile. That is the empty-slot square in `--line` with a `bg-card-top` ground, at the same 28 / 20 px. Then show the name's letters. Remove the `QUESTION_MARK_ICON` step from `RaidLootItem` and `lib/guild/loot-icon.ts`.

### Findings

1. Fixed. There is no officer row on `done`. The gate sentence shows once (01-done-officer-1440, 01-done-officer-360, 08-read-only-panel-390). Ended keeps `Edit` (03-ended-officer-window-panel-1440).
2. Fixed. `Notes` and `Sign-ups 14/40` are closed, each with its chevron (01-done-officer-1440, 02-member-present-1440).
3. Fixed. The `Before` / `In the log` heads show from 768, and inline lead-ins show at 360 (01-done-officer-1440, 01-done-officer-360).
4. Not fixed for Head of Onyxia. It still shows the red `?`. The other three awards show real icons (01-done-officer-1440, 01b-first-screen-officer-1440, 02-member-present-390). Must fix, per the ruling above.
5. Fixed. The no-show row shows only the sentence and `No loot this raid.` (02-member-no-show-panel-360).
6. Fixed. The kill row ends at y 609 at 1440 × 800. The member's sentence starts at y 602 at 360 × 640. The phone bar sits on the count's line, and the when line takes two lines (01b-first-screen-*, first-screen-*.json).
7. Fixed. Separators are spaced, words are in the body face and figures are in mono, on both the night line and the facts line (01-done-officer-panel-1440, 02-member-present-panel-1440).
8. Fixed. Loading beside ready differs by 0 px for the tabs, the band, `Your night` and the Report top (heights-reserve-*.json, 11-loading-member-390 beside 11-ready-member-390). CLS is 0.0035 at most.
9. Fixed. Deaths take two phone lines, with whole names and nothing wrapping (01-done-officer-360).
10. Fixed. At 1024 the bar is capped and the split is on one line (13-side-by-side-officer-1024).
11. Fixed. `Deaths 17` matches `17 deaths`, and the link reads `Show all 11 raiders` (01-done-officer-1440).
12. Fixed. Loot and Changes sit on aligned grids from 768 and take two lines at 360 (01-done-officer-1440, 01-done-officer-360).
13. Fixed. Hexbramble is gone from Consumables, and every row is a raider in the log (01-done-officer-1440).
14. Fixed. `Biggest changes` is kept. The `Before` column sits at the same x as in the full report (07-bare-report-panel-1440).
15. Partly fixed. The rows are three lines, but at 360 the zone name breaks (`Onyxia's` / `Lair`) beside `Closest` and `2 h 41 min` (04-picker-closest-panel-360). Should fix: line 1 holds the zone and `Closest`, `whitespace-nowrap`. The duration moves to line 2, mono and right-aligned on the time's line.
16. Fixed. Shadowquill reads `—` from Gear gap to Item level (10-readiness-on-panel-1440).
17. Fixed. The 1024 board is drawn at 1024, and the side-by-side is at the same scale (13-side-by-side-officer-1024).
18. Fixed. `Keep it` draws the 2 px gold ring (05-clear-confirm-panel-1440).
19. Fixed. `0 deaths` shows (02-member-present-panel-1440).
20. Fixed. The row reads `Out` in `--muted` (06-attendance-panel-1440).
21. Fixed. The sub-head reads `LOOT` with no count (07-bare-report-panel-1440).
22. Fixed. The empty state shows no helper (04-picker-empty-panel-390).
23. Fixed. The figure no longer breaks (09-guild-line-panel-360). Polish: the wrapped line starts with a stray `·`. Below 768, drop the separator when the clause starts its own line.
24. Fixed. The kill time and note are right-aligned (01-done-officer-1440, 01-done-officer-360).
25. Fixed. The pill fills the 120 px column (01-done-officer-1440).
26. No change, as ruled. The mock amendment is deferred in the README.
27. Open, out of lane. It is deferred to the roster lane in the README. It stays tracked.

### New

- N1, must fix: a regression from the shared padding change. `md:pt-7` became `md:pt-0` on `main` in `guild.astro` and `guild/[...path].astro`. That is correct for the views that open with a full-bleed band: the raid page and the guild home (01, 09-roster-raids-1440). It is not correct for the views with no band. At 768 and up, their first line now sits flush against the header's bottom rule with 0 px between. Those views are `GuildLanding` (`/guild`), `GuildSettings`, `GuildClaim`, `GuildJoin`, the claim and settings chunk-pending wrappers in `GuildShell`, and `Guild.svelte`'s `missing`, `loading` and `failed` states. None of them is captured. Fix: keep `main` at `md:pt-0`, and give each band-less view's root `md:pt-7`. Attach 1440 and 390 captures of `/guild`, a guild's settings, claim, join with a bad token, and a guild address that does not parse. Shared layout changed, so run the full web vitest, the full Playwright suite and the lhci shard.
- N2, should fix: in the Attendance phone rows, the `BOSS PULLS` lead-in wraps to two lines on Mendwhisper's row (01-done-officer-360, Late group). Fix: `whitespace-nowrap` on the lead-in.
- N3, polish: at 360, `Biggest changes since Thu 3 Dec` leaves `DEC` alone on line 2 (01-done-officer-360). Fix: `since Thu 3 Dec` is `whitespace-nowrap`, so it breaks as one unit.
- N4, polish: in the when line, the words in `ended 1 h ago` are set in mono (01-done-officer-panel-1440). This is G1's grammar, but finding 7's rule applies: only the figures go in mono.

### Checklist (re-run)
- [x] The first screen meets §1 at 360 and 1280 (finding 6).
- [x] Copy is verbatim: ` · ` spacing, `Before` / `In the log` (findings 3, 7).
- [ ] Real icons on every item: Head of Onyxia still shows the red `?` (finding 4).
- [x] R-A4: Notes and Sign-ups are closed (finding 2).
- [x] Loading beside ready matches within ±1 px (finding 8).
- [ ] Every count and name survives 360: the picker zone and the `BOSS PULLS` lead-in break (15, N2).
- [ ] The other guild views keep their top spacing after the shared padding change (N1).

### Verdict: SHIP WITH FIXES

Two items block the merge. Finding 4: Head of Onyxia shows `inv_misc_head_dragon_01`, and no `?` fallback remains. N1: the band-less guild views get their 28 px top padding back, shown in the new captures. Finding 15 and N2 land in the same pass. N3, N4 and the finding 23 polish can follow.
