# Guild raids and sign-ups: acceptance captures

Captures for the experience spec `design/specs/2026-10-26-guild-raids.md` §9, rows 1–16 and 18.
Row 17 (the live Discord clients) belongs to PR G1b. The captures come from
`web/tests/captures/guild-raids/capture.spec.ts`, run against a fixture build served on port 4441:

```
cd web
FOREVER_DATA=fixture npm run build && cp -R dist /tmp/g1-capture-dist
npx astro preview --outDir /tmp/g1-capture-dist --port 4441   # in the background
python3 ../design/mocks/gen_guild_raids.py                     # mock boards for row 18
npx playwright test -c tests/captures/guild-raids/playwright.capture.config.ts
```

The stubs are shared with the e2e (`web/tests/e2e/support/guild-raids.ts`). The guild is the
seeded mock guild, Olympus XXVII. The clock is pinned to Tue 8 Dec 2026, 12:00 Chicago. The
browser runs on Chicago time, except for the `For you` form and row 18, which use New York
(the mock boards are drawn from Eastern time). The member is Zulmara, whose alt Duskstrider gives
R3 a second character. The officer is Kraggor.

## Capture list

Each state has a full-page PNG, `guild-raids-<nn>-<state>-<width>.png`, and most have a crop,
`guild-raids-<nn>-<state>-panel-<width>.png`, of the tab panel, the raid page or the band. A long
panel's crop is its top 640 px. The form crops show the whole form.

| # | State | File stem(s) | Widths |
|---|---|---|---|
| 1 | Raids tab, officer: `Plan a raid`, 3 planned (one cancelled), past nights | `01-tab-officer` | 360, 390, 1440, 2000 |
| 2 | Raids tab, member: `In · DPS` and `Not signed up` chips | `02-tab-member` | 390, 1440 |
| 3 | Raids tab, empty Planned, member and officer | `03-tab-empty-member`, `03-tab-empty-officer` | 390, 1440 |
| 4 | Raids tab, officer-read-only | `04-tab-read-only` | 360, 1440 |
| 5 | Raid page, member In on Zulmara, Duskstrider on Bench; plus a `-strip-` crop | `05-page-member-in` | 360, 390, 1440, 2000 |
| 6 | Raid page, member unsigned | `06-page-member-unsigned` | 360, 390, 1440 |
| 7 | Raid page locked; the lock race line, focused | `07-page-locked`, `07-page-lock-race` | 360, 390, 1440 |
| 8 | Raid page cancelled, ended, done | `08-page-cancelled`, `08-page-ended`, `08-page-done` | 360, 390, 1440 |
| 9 | Raid page, officer: inline selects, officer row; plus `-roster-` and `-confirm-` (cancel confirm) crops | `09-page-officer` | 390, 1440, 2000 |
| 10 | Raid page, officer-read-only | `10-page-read-only` | 390, 1440 |
| 11 | Raid page, signed-out visitor and unverified viewer | `11-page-visitor`, `11-page-unverified` | 360, 390, 1440 |
| 12 | Raid page and Raids tab: the loading reserve (warm), then ready | `12-page-loading`, `12-page-ready`, `12-tab-loading`, `12-tab-ready` | 390, 1440 |
| 13 | Form: create default, errors, `For you` (viewer in New York), repeat at the cap with no webhook | `13-form-default`, `13-form-errors`, `13-form-for-you`, `13-form-cap-no-webhook` | 360, 390, 1440 |
| 14 | Band line: In, unsigned, none planned (member and officer), with the needs line | `14-band-in`, `14-band-unsigned`, `14-band-none-member`, `14-band-none-officer` | 360, 390, 1440, 2000 |
| 15 | Focus-visible: planned row, segment, role, note, officer button, form zone/date/time zone/submit; hover on a planned row and an officer button | `15-focus-*`, `15-hover-*` | 390, 1440 (hover 1440) |
| 16 | Copy failed, the textarea shown | `16-copy-failed` | 390, 1440 |
| 18 | Same-scale side by side with the mock boards (`guild-raids-tab`, `guild-raids-page`), plus each half alone | `18-side-by-side-{tab,page}`, `18-mock-*`, `18-built-*` | 1024, 1440, 2000 |

## Heights against §7

Measured in the browser. The raw numbers are in `heights-<surface>-<width>.json`.

| Surface | 360 | 390 | 1440 | 2000 | Cap (phone / desktop) |
|---|---|---|---|---|---|
| Planned row | 97 | 97 | **77** | **77** | 96–120 / 64–72 |
| Strip row | 183 (whole row with status line) | 183 | 68 (grid row); 100 with the 12 px top pad and 20 px status line | 68 | ≤ 196 / 68 |
| Band line | 98 | 98 | 44 | 44 | ≤ 112 / 44 |
| Roster row, no note | 44 | 44 | **38** | **38** | 44 / 36 |
| Form | 989 | 989 | 726 | — | ≤ 1180 / ≤ 760 |

Two surfaces are over their desktop caps: the planned row is 77 px against 72 and the roster row
is 38 px against 36. Everything else is within its cap.

**CLS with a warm reserve** (`cls-<width>.json`, cap 0.02). The value is summed from the first
paint, so it includes the band and tabs settling as well as the reserve.

| Width | Raid page | Raids tab |
|---|---|---|
| 390 | **0.226** | **0.359** |
| 1440 | **0.024** | **0.126** |

Every value is over 0.02. In `12-page-loading-390` beside `12-page-ready-390`, the band reserve
(300 px) and the strip reserve (2 rows × 196 + 56 px) do not match the ready heights. The band
ends near 420 px and the strip is shorter, so the notes and roster jump.

## Noticed while capturing

- **The when and count lines lose their spaces from `md` up.** The 1440 band reads
  `Thursday 10 December·8:00 pm to 11:00 pm·in 2 days 7 h` and `14 of 40 signed up· 1 tank`.
  The ` · ` separators sit in `hidden md:inline` spans, and their leading and trailing spaces are
  trimmed (05, 09, row 18).
- **The notes list has no bullets** (05 and row 18): `prose-forever` drops `list-style`, so
  `Onyxia Scale Cloak on everyone` reads as an indented line. The mock has bullets.
- **Late, Bench and Out stack in one column at 1440.** The mock sets them side by side in three
  columns under the role columns (row 18, page). The built page stacks each group full width,
  so the page runs taller. §3.7 does not say which.
- **The no-webhook form line's inline `Settings` link is 16 px tall at 360,** with no phone hit
  area (13). The e2e keeps this as a `test.fail` in `guild-raids.spec.ts`, which turns red once
  the link is fixed.
- **The form checkboxes are the browser's native box,** not the site-wide gold checkbox §2 names
  (13).
- **The officer band has no standing line** in the built page (row 18, tab), where the mock shows
  `Kraggor · 2 of 3 Protection Warriors…`. That comes from the fixture: the mock home builds a
  standing only for the member viewer's roster row.
- **The chip order on a cancelled row is reversed:** the built page shows `IN · TANK` then
  `CANCELLED`, and the mock shows `CANCELLED` then the viewer chip (01, row 18).
- **The site header reads `Sign in` in the captures.** The header selector reads its own
  endpoints, which these stubs do not answer. This is fixture-only.
- The ux-designer and wow-player reviews of these PNGs (brief Step 4) were not run in this task.
  The controller dispatches them.

## Pre-review fix round

Re-taken: rows 1, 2, 5, 9, 12 and 13 at their widths, from a fresh fixture build on port 4441.
`guild-raids-13-form-checked-panel-1440.png` is new: the form with `Post to Discord` ticked.

- **Separators** (05, 09, 18): the when line, count line and title suffix now read
  `Thursday 10 December · 8:00 pm to 11:00 pm · in 2 days 7 h` and `14 of 40 signed up · 1 tank …`
  from `md` up. Svelte trims the edge spaces of literal text inside a span, so each separator is
  now an expression (`{SEP}`). The planned row's phone line had the same pattern and is fixed too.
- **Notes lists** (05): the notes container adds `list-disc` and `list-decimal` to the
  `prose-forever` style, with muted markers.
- **Checkboxes** (13): no change was needed. The site-wide rule already applies: the computed
  `accent-color` is `rgb(229, 185, 85)` and `color-scheme` is `dark`. The earlier captures only
  showed unchecked boxes, and an unchecked box is the native dark box under the §9.2 rule. A
  checked box is gold (`13-form-checked-panel-1440`).
- **Settings link** (13): the no-webhook line's link takes the gate link's 44 px `::after` at
  phone widths (`INLINE_LINK` in `lib/guild/field-styles.ts`). The e2e `test.fail` is removed
  and the 360 control check passes.
- **Cancelled row** (01): the status chip now comes before the viewer chip (`CANCELLED`, then
  `IN · TANK`).

| Surface | 360 | 390 | 1440 | 2000 | Cap (phone / desktop) |
|---|---|---|---|---|---|
| Planned row | 97 | 97 | 71.5, 71, 71 | 71.5, 71, 71 | 96–120 / 64–72 |
| Roster row, no note | 44 | 44 | 36 | 36 | 44 / 36 |
| Strip row | 183, 182 | 183, 182 | 68 (101/100 whole) | 68 | ≤ 196 / 68 |
| Form | 989 | 989 | 725.5 | — | ≤ 1180 / ≤ 760 |

The desktop planned row's 5 extra px came from the 56 px date block plus 10 px of vertical
padding on each side (77). The vertical padding is now 7 px, which holds every cell. A bare
roster row now centres its crest, name and select on the 36 px floor with no padding past them.
A row with a note keeps its top padding.

**CLS with a warm reserve** (`cls-<width>.json`):

| Width | Raid page | Raids tab |
|---|---|---|
| 390 | **0** (was 0.226) | 0.206 (was 0.359) |
| 1440 | **0.0001** (was 0.024) | 0.123 (was 0.126) |

Most of the shift came from the footer, not from the reserves. Every /guild/* path serves the
same prerendered shell. Its first paint was about 630 px of lede and skeleton, so the footer
started inside the viewport at y 703 and jumped when the hydrated view grew. `GuildShell` now
holds every guild view at least one screen tall from the first paint
(`min-h-[calc(100svh-72px)] md:min-h-[calc(100svh-109px)]`), so the footer starts below the
fold. Content inside the reserves is replaced with new nodes, not moved, so it adds no shift.

The reserves now equal the ready heights measured on the fixture (`lib/guild/raid-layout.ts`):
- raid band: 347 / 311 px for the member;
- strip: 62 + rows × 182 px on phone, 70 + rows × 100 px on desktop (426 and 270 for two rows,
  the same as ready);
- Raids tab planned block: 62 + rows × 97 px on phone, 62 + rows × 71 px on desktop (the tab's
  header row, `PLANNED n`, and the rows: 353 and 275 for three rows against ready 353 and 275).

**Floor on the Raids tab:** 0.12–0.21. Layout-shift sources show the guild band itself growing
when `/home` answers. The tabs move from y 273 to y 417 at 1440, and the needs line and raid
line appear, so everything under the band moves. This is the guild band's own loading state in
Guild.svelte's header, not a raid reserve; this round was limited to the reserve block. The
value varies with how close `/v1/me` and `/home` land to first paint (a diagnostic run read
0.003 at 390). The fix is a member-band reserve while `/home` is in flight, held to the band's
ready height. That is a guild-band change for the band's owner.

## Review round (2026-10-27)

The ux-designer and wow-player fixes (both SHIP WITH FIXES), every item accepted, plus the Raids
tab CLS floor. Spec amended in place, each change marked `(amended 2026-10-27, review round)`.
Re-taken from a fresh fixture build on port 4441: rows 1, 2, 5, 7, 8, 9 (with the cancel
confirm), 12, 13 and 14 at their widths. New files:
`guild-raids-05-page-member-in-new-york-{390,1440}` (+ `-panel-`): the viewer in
America/New_York, the guild on America/Chicago, so the band's when line reads
`9:00 pm to 12:00 am · 8:00 pm CST guild time`; and `guild-raids-09-page-officer-roster-1024`
beside the 1440 roster.

**Raid page**
- Inert states (07, 08): locked, ended and done draw the segment group, the role select and the
  note at `opacity-50` with `cursor-not-allowed`. The group dims as one, its border included.
  `FilterSelect` dims only when its `select` is disabled; a disabled `—` option does not count.
- Gold only for a raid still ahead (08, 01): a cancelled raid, or one that has started, draws its
  countdown and bar fill `--muted`. The cancelled page's count is muted too, and the page shows
  only the `Cancelled` chip.
- Counts that add up: `14 of 40 signed up · 1 tank · 3 healers · 9 DPS · 1 late`. The late clause
  shows only when late > 0. The page band, the planned row and the copied list all use it. The
  copied list reads `Onyxia's Lair · Thu 10 Dec, 8:00 pm CST · <t:1796954400:F>`, then
  `14 of 40: 1 tank, 3 healers, 9 DPS, 1 late`.
- Chosen segment (05 strip): no class-order bug. The checked segment computes
  `rgb(19, 24, 36)` (`--card-top`) against `rgb(13, 17, 26)` (`--raised`) at rest, with
  `--strong` text and the 2 px line, exactly as §3.3 says. The fill was already drawn; the two
  grounds are close, so it reads faint. The designer rules whether §3.3 needs a stronger ground.
  The wow-player asked for `--kill` at about 15 % for In and Late.
- Roster (09 at 1440 and 1024): the name sits over the spec at every width, one line each with
  an ellipsis and the full text in `title`. Bare rows are 36 / 44 px and line up across columns.
  Late, Bench and Out form one grid: three columns from 1024, two at 768–1023, stacked below.
  Empty groups are left out. `Not signed up` runs full width 24 px under them. Every group head
  is the 11 px `label`, muted, with its count in mono.
- Cancel confirm (09 confirm): `Cancel the raid` is drawn in the `--wipe` outline. The Discord
  box starts from `discord.post_default`. The localStorage memory of the last answer
  (`fs.guildRaid.postCancel`) is deleted.
- Loading (12): the back link, the eyebrow and `GuildTabs` are drawn for real while the raid
  loads. Only lines 3–5 are skeleton, at their ready 163 / 105 px. A cached officer adds 108 px
  on phone and 22 px on desktop for the officer row.

**Raids tab**
- Cancelled row (01, 02): shows `Cancelled` alone, with a muted countdown and bar. Phone line 1
  holds one chip: the status chip when there is one, otherwise the viewer's.
- Planned row: the desktop bar now sits on the count's line. The row stays 71 px, because the
  56 px date block sets the height, not the count cell. On phone the breakdown has its own
  fourth line, and the title wraps to two lines (`line-clamp-2`).
- `Plan a raid` is hidden while the form is open. On close it returns with focus (e2e).
  `PLANNED` drops its count at 0.

**Form (13)**: `Raid`, `Duration` and `Guild time zone` take `SELECT_CHROME` and the shared
`SelectChevron`, which is extracted from `FilterSelect`, so no native arrow is left. The repeat
box at the cap was already `disabled` in the DOM (unit test). It looked live because a native
disabled box looks the same, so its row now takes the disabled look.

**Band line (14)**: when `own` is another of the account's characters at In or Late, the
segments act on that character, with its role. Their `aria-label` names it (`Duskstrider's
sign-up`). That is the designer's narrower rule of fix 11. When the other character is only on
Bench or Out, the segments stay on the current character, because a press there benches no one.
The unsigned sentence reads `Thursday 8 pm, you have not signed up yet. 13 of 40 are in.`. The
needs line is `--text`, not `--wipe`; a passing line stays green.

**Heights**

| Surface | 360 | 390 | 1440 | 2000 | Cap (phone / desktop) |
|---|---|---|---|---|---|
| Planned row | 114, **136.5**, 114 | 114, 114, 114 | 71, 71, 71 | 71, 71, 71 | 96–120 / 64–72 |
| Roster row, no note | 44 | 44 | 36 | 36 | 44 / 36 |
| Strip row | 183, 182 | 183, 182 | 68 (101/100 whole) | 68 | ≤ 196 / 68 |
| Band line | 98 | 98 | 44 | 44 | ≤ 112 / 44 |
| Form | 989 | 989 | 725.5 | — | ≤ 1180 / ≤ 760 |

At 360, `Barrow Deeps: first clear` wraps to two lines under item 9, and with the breakdown on
its own line under item 11 that row is 136.5 px, over the 120 cap. Both items together cannot fit
120 px with a two-line title. One-line rows are 114 px. The designer rules: raise the phone cap
to 140 for a two-line title, or keep one of the two items.

**CLS with a warm reserve** (`cls-<width>.json`). Capture 12 now makes a real warm visit first:
the page itself stores its strip rows, planned rows and band height. It then drops the read
cache (`fs.q.*`) and measures a reload with the raid read held.

| Width | Raid page | Raids tab |
|---|---|---|
| 390 | 0.0001 | **0.0017** (was 0.206) |
| 1440 | 0.0001 | **0.0006** (was 0.123) |

**The Raids tab floor took both measures: the band hold, and the one-screen floor in
`GuildShell`.** The band hold (`lib/guild/band-hold.ts`) stores the band header's last height in
`fs.guildBand.{region}.{ruleset}.{slug}`, per guild and per layout (phone / desktop), for a
verified viewer, behind try/catch. While `/home` is in flight it holds that height as the
header's `min-height`. It is forgotten once the viewer is no longer verified there. It holds the
measured header, not the designer's flat 44 + 12 / 98 + 12. The member block that `/home` adds
is 101 px on desktop and 215 px on phone (standing line, raid line, needs line, settings link),
and `/v1/me` adds the `Claimed by` line, which alone moved the tabs 31 px in a run that held only
the raid line. The planned reserve now also shows from the first paint for a viewer last seen
verified. Before, it waited for `/home` to name the role, and the skeleton under it jumped
307 px. Without `GuildShell`'s `min-h-[calc(100svh-…)]`, the same runs read 0.232 / 0.173 at 390
and 0.040 / 0.030 at 1440 (page / tab), because the footer starts inside the viewport and moves.
The floor stays.

**Reserves** (`lib/guild/raid-layout.ts`): raid band lines 163 / 105, officer row 108 / 22.
Planned block: 62 + rows × 121.5 on phone (the fixture member's mean of one-line 114 and
two-line 136.5 rows), 62 + rows × 71 on desktop. Strip unchanged.

Size: `GuildShell` 16292 bytes gzipped (budget 16384).

## Discord section (2026-10-27)

Guild settings · Discord (spec §3.8, lane G1b web half), captured from a fresh fixture build on
port 4441 by `web/tests/captures/guild-raids/capture-discord.spec.ts` (the capture config now
matches `capture*.spec.ts`). Files are `guild-raids-<row>-settings-discord-<state>-<width>.png`,
each with a `-panel-` crop of the section (row 19's crop starts at Billing so the order shows),
and `heights-discord-<state>-<width>.json`.

- **19** not connected, beside Billing: 360, 390, 1440, 2000.
- **20** connected at rest, with the 3-failure warning (`three-failures`), and the remove question
  (focus on `Cancel`): 360, 390, 1440.
- **21** switched off: 360, 390, 1440. `Test and switch on` is the one fill.
- **22** returns: `return-connected` (focus ring on `Send test message`, the connected line),
  `return-denied` and `return-failed` (focus on `Connect a Discord channel`): 360, 390, 1440.
- **23** officer-read-only, not connected and connected (`Send test message` dimmed and
  disabled, `Remove` live, the gate sentence's `Subscribe the guild` to `#billing`), and frozen:
  360, 1440.
- **24** (the test embed in the Discord clients) needs a staging channel; not captured here.

**Heights against §3.8**

| State | 360 | 390 | 1440 | 2000 | Cap (360 / 1440) |
|---|---|---|---|---|---|
| Not connected | 198 | 198 | 139.5 | 139.5 | 210 / 140 |
| Connected, at rest | 249.5 | 249.5 | 230 | — | 300 / 230 |
| Connected, remove question | 208.5 | 208.5 | 187.5 | — | 300 / 230 |
| Connected, 3 failures | **322** | **322** | **283** | — | 300 / 230 |
| Switched off | 349 | 349 | 310 | — | 440 / 330 |
| Read-only, not connected | 250 | — | 171.5 | — | no cap (adds the gate sentence) |
| Read-only, connected | 301.5 | — | 262 | — | no cap (adds the gate sentence) |

- The 140 at 1440 holds only with the explanation on one line (27 + 12 + 44 + 8 + 20 + 8 + 21).
  Held to the 720 px prose measure, it wraps to two lines and the state is 159 px. So the
  not-connected column is not capped at 720: the explanation runs 1053 px at 1440. The connected
  and switched-off columns keep 720.
- The connected cap is the state at rest (230 exactly). The 1 to 4 failure warning adds its two
  lines on top (52 px desktop, 72 px phone), so the warned state is over both caps. The designer
  rules: a cap for the warned state, or a shorter warning.
- The reserves in `lib/guild/discord-state.ts` are these measured rest heights. They are the
  placeholder while the section's chunk loads, and the connected hold after a `connected` return.

**Segment ruling (row 5, re-taken at 390 and 1440).** The checked segment's ground is now
`--kill` at 15 % for In and Late, and `--line-warm` for Bench and Out (§3.3 amended). See
`guild-raids-05-page-member-in-strip-390.png`: Zulmara is In, Duskstrider on Bench.

**Planned row (§7 amended).** A planned row whose title wraps to two lines at 360 may run to
140 px. The 136.5 px `Barrow Deeps: first clear` row in the heights table above is inside it. No
code change.

Size: `GuildShell` 16321 bytes gzipped (budget 16384; 16292 at 9158489e). The Discord section is
its own chunk, `GuildDiscordSection` 3853 gzipped. It shares the account's Discord components
chunk (5912) and loads only when `settings.discord` is present.
