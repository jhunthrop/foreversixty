# Guild attendance and the raid-night report: acceptance captures

Captures for the experience spec `design/specs/2026-10-27-guild-attendance.md` §9, rows 1–11 and 13.
Row 12 (the `report_ready` embed and the copied text in a staging channel's Discord clients) needs a
live channel and is not captured here. The captures come from
`web/tests/captures/guild-attendance/capture.spec.ts`, run against a fixture build served on port 4451:

```
cd web
FOREVER_DATA=fixture npm run build && cp -R dist <dir>
npx astro preview --outDir <dir> --port 4451        # backgrounds itself
python3 ../design/mocks/gen_guild_attendance.py      # mock boards for row 13
npx playwright test -c tests/captures/guild-attendance/playwright.capture.config.ts
```

The stubs are shared with the e2e (`web/tests/e2e/support/guild-attendance.ts`, on top of
`guild-raids.ts`). The raid is G1's Onyxia's Lair of Thu 10 Dec 2026, 8:00 pm Chicago, logged as
3 wipes then a kill (`guild-attendance-data.ts`): 13 of 14 came, Mendwhisper and Arcanemoor late,
Hexbramble not in the log, Felsnap came as Mirestalk, Soulgrave joined from the bench, Thornhide
joined without a sign-up, Frostnettle out. The clock is 1 h after the raid's end (inside the
matcher's 6 h window), or 7 h after it for the `-after` states. The member is Zulmara, whose alt
Duskstrider is the member's character in the log; the officer is Kraggor. Browser on Chicago time,
except row 13, which uses New York (the mock boards are drawn from Eastern time).

## Capture list

Each state has a full-page PNG, `guild-attendance-<nn>-<state>-<width>.png`, and most have a
`-panel-` crop (the band, the report, a section or the tab; at most 900 px tall).

| # | State | File stem(s) | Widths |
|---|---|---|---|
| 1 | Done raid, officer: night line, match line, Report, Attendance; plus a `-report-` crop | `01-done-officer` | 360, 390, 1440, 2000 |
| 2 | Done raid, member: `Your night` present, and the no-show variant | `02-member-present`, `02-member-no-show` | 360, 390, 1440 |
| 3 | Ended, no log: officer and member, inside (`-window`) and after (`-after`) the 6 h window | `03-ended-{officer,member}-{window,after}` | 360, 390, 1440 |
| 4 | Picker: `Closest` with a matched-elsewhere row, empty, unknown log zone | `04-picker-closest`, `04-picker-empty`, `04-picker-unknown-zone` | 360, 390, 1440 |
| 5 | Clear confirm, focus on `Keep it` | `05-clear-confirm` | 390, 1440 |
| 6 | Attendance: every group, Mendwhisper excused, Felsnap `as Mirestalk` | `06-attendance` | 360, 390, 1440, 2000 |
| 7 | Report with no buff data, no previous raid, no loot | `07-bare-report` | 390, 1440 |
| 8 | Officer-read-only done raid; plus an `-attendance-` crop | `08-read-only` | 390, 1440 |
| 9 | Roster chip in raid mode (with the footnote and a `No raids yet` row) and report mode; the account page's character guild line | `09-roster-raids`, `09-roster-reports`, `09-guild-line` | 360, 390, 1440 |
| 10 | Readiness, next raid on (Windtalon's stale export, Shadowquill's missing one) and off | `10-readiness-on`, `10-readiness-off` | 360, 390, 1440 |
| 11 | Loading beside ready, officer and member, warm reserve; CLS in `cls-<width>.json` | `11-loading-{officer,member}`, `11-ready-{officer,member}` | 390, 1440 |
| 13 | Same-scale side by side with the mock boards: `guild-attendance-page` (officer) at 1024, 1440, 2000; `guild-attendance-page-phone` (member) at 390; plus each half alone | `13-side-by-side-{officer,member}`, `13-mock-*`, `13-built-*` | 390, 1024, 1440, 2000 |

## CLS (row 11, cap 0.02)

Summed from the first paint of a reload whose raid read is held, after a warm visit (the page
stores `fs.guildRaid.rows` and the new `fs.guildRaidDone.{id}`; the read cache is dropped).

| Width | Officer, done raid | Member, done raid | Poll's done swap, on its own |
|---|---|---|---|
| 390 | 0.0001 | 0.0001 | 0.2732 |
| 1440 | 0.0001 | 0.0001 | 0.202 |

The poll swap is the spec's one allowed height change (§4 "Poll"), announced by the live region;
it is measured from a settled ended page to the done one and reported, not capped.

## Reviews

The ux-designer and wow-player reviews of these PNGs are dispatched by the controller.

## Persona fix round (2026-10-27)

The ux-designer (UX 1–27) and wow-player (WP 1–21) reviews, both SHIP WITH FIXES, with the
controller's rulings. Every capture row was re-taken from a fresh fixture build on port 4451. API
commit `b6f2b9da`; the web fixes are in the `fix(web): persona round …` commit after it.

| Finding | Fix | Capture |
|---|---|---|
| UX 1, WP 5, WP 18 | No G1 officer row on a `done` raid; an ended raid keeps `Edit` only beside `Match a log`; the read-only gate sentence shows once, under the Report header | 01, 03, 08 |
| UX 2, WP 6 | Notes and Sign-ups are closed `<details>` (`RaidDisclosure`, 44 px summary, chevron) below Attendance; the selects inside stay disabled | 01, 02 |
| UX 3, UX 13, WP 3 | Consumables: `Before` / `In the log` heads from 768, inline lead-ins below; rows only for raiders in the log (the API already filtered; the fixture now matches); three columns kept without buff data; the pill fills its column (UX 25) | 01, 07 |
| UX 4, WP 1 | Awards carry `icon` from the API's loot table; none (Head of Onyxia) or a failed file shows the client's question mark, then the name's letters; the tooltip opens with the award's own facts for an item outside every class file | 01, 02 |
| UX 5, WP 2 | A character not in the log shows no deaths and no move; only `Excused` when set; the no-show fixture drops the alt's rows | 02 |
| UX 6 | The 28 px strip above the band is gone (guild pages' `md:pt-7`); `Not this log` is 44 px on phone only; the phone bar sits on the count's line; done and ended when lines are two phone lines. Kill-order row 1 ends at y 609 at 1440 × 800; the member's first sentence starts at y 602 at 360 × 640 (`first-screen-*.json`, guide line drawn in `01b-first-screen-*`) | 01b |
| UX 7, WP 13, WP 19 | Night line and Your night facts: words in the body face, figures in mono (`Figures`), each clause whole | 01, 02 |
| UX 8 | Reserves equal the ready render: loading beside ready deltas are 0 px for the tabs, the band, Your night and the Report top (`heights-reserve-*.json`) | 11 |
| UX 9, UX 10 | Deaths: two phone lines (whole name and count; bar and split); the bar caps at 96 px from 1024 to 1279; the split never wraps | 01 |
| UX 11 | Deaths sub-head counts deaths (17); `Show all 11 raiders` counts raiders | 01 |
| UX 12 | Loot and Changes on aligned grids from 768, two lines below | 01 |
| UX 14, WP 12 | `Biggest changes` heading kept over the no-previous line | 07 |
| UX 15, UX 22 | Picker phone rows: three lines (zone, Closest, duration; time; `by · figures`); the empty list hides the helper | 04 |
| UX 16, WP 14 | `No export yet` rows read `—` from Gear gap to Item level; the fixture sends no gear for them | 10 |
| UX 17 | The officer mock board is drawn at 1024 for the 1024 side by side (`guild-attendance-page-1024`) | 13 |
| UX 18 | Row 5 opens the confirm from the keyboard, so `Keep it` shows its ring | 05 |
| UX 19 | `0 deaths` shows for a character in the log | 02 |
| UX 20, WP 17 | `Bench and out` rows say `Bench` or `Out`, muted, in place of `0/4` | 06 |
| UX 21 | A count of 0 is left out of a sub-head (`LOOT`) | 07 |
| UX 23 | The guild line's `came 7 of 8 raids` never breaks | 09 |
| UX 24 | Kill time and note right-aligned | 01 |
| WP 4 | Cause: the fixture, not the code. The API reads both from one `raidAttendance`; the stub now builds the roster chip and the account's guild line from one map (`rosterAttendance`), and Kraggor reads his raids | 09 |
| WP 8 | A `Boss pulls` column head; full rows read `4 of 4` | 06 |
| WP 9 | A raider signed Late who came late is present (API, test); their sentence still says when they came; spec §3.5 amended | 06, 02 |
| WP 10 | The night line adds ` · 2 joined` | 01 |
| WP 20 | `Sorted worst first. Nudge …` keeps its space | 10 |
| WP 21 | A one-boss raid drops the boss head over its loot | 01 |
| Fixture realism | Onyxia 4:12 after 3 wipes; consumables = every raider in the log; no `+0 DPS` on the stale readiness row | 01, 10 |

**CLS** (cap 0.02): officer 0.0001 / 0.0035 (1440 / 390), member 0.0001 / 0.0001, the poll's done
swap 0.0039 / 0.0008.

**Reserves** (`lib/guild/raid-layout.ts`, measured): done band lines, member 223 / 157 px, officer
267 / 157 px (phone / desktop); Your night 56 + rows × 172 / 68 + rows × 68 px, rows from
`fs.guildRaidDone.{id}` (the own attendance rows last drawn).

### Follow-ups (deferred, not built)

- WP 11: the picker leads with the first pull's time; the candidates carry no pull time yet.
- WP 15: `Felsnap as Mirestalk` in Deaths, Consumables and Changes.
- WP 16: boss health at each wipe, killing blows, named fire-protection potions.
- UX 26: the mock's night bar is amended to `--muted` (no build change).
- UX 27: Brightwick's empty roster crest (roster lane).

## Final round (2026-10-27, re-reviews)

Both re-reviews SHIP WITH FIXES. API `0d2a4aa0`; web in the `fix(web): final round …` commit.

| Finding | Fix | Capture |
|---|---|---|
| UX 4 (must) | Head of Onyxia shows `inv_misc_head_dragon_01`. The loot table's six slotless drops take the client's own icon, `Item.IconFileDataID` through `ManifestInterfaceData` (wago.tools, build 1.60.1.70532), never guessed: Scale of Onyxia `inv_misc_monsterscales_07`, Onyxia Hide Backpack `inv_misc_bag_22`, Mature Black Dragon Sinew `spell_holy_dispelmagic`, Draconic for Dummies `inv_misc_stonetablet_03`. The two files missing from the build's icon set (published through `web/public/data/<build>/icons`) were added from CASC with the pipeline's `download_icons`. No question-mark step remains: an item with no icon draws the `/bis` neutral tile with the name's letters | 01, 02 |
| UX N1 (must) | `main` stays `md:pt-0`; every band-less view's root takes `md:pt-7` (`GUILD_BANDLESS_TOP`): GuildLanding, GuildSettings, GuildClaim, GuildJoin, GuildShell's claim, settings and join chunk placeholders, Guild.svelte's missing, loading and failed states | `N1-{landing,claim,settings,join-bad-token,not-an-address}-{390,1440}` |
| WP A (must) | The roster chip names its unit once for both windows: `Came 7 of 8 · 26 of 30 raids` | 09 |
| WP 6 (must) | Notes and Sign-ups opened on a done raid: the selects are disabled | `08-sections-open`, `08-sections-open-roster` |
| UX 15 | Picker phone line 1 is the zone and `Closest`, never broken; the duration moves to line 2 beside the time | 04 |
| UX N2 | The `Boss pulls` lead-in never wraps | 01 (360), 06 |
| WP B | The readiness note names the order: `Sorted by the most checks to fix, then the biggest gear gap.` | 10 |
| WP C | On phone the Present group head carries `Boss pulls` | 06 (360, 390) |
| WP 11 (interim) | A muted `Uploaded` before each candidate's time | 04 |
| UX N3 | `since Thu 3 Dec` breaks as one unit | 01 (360) |
| UX N4 | `ended 1 h ago`: figures in mono, words in the body face | 01 |
| UX 23, WP E | The guild line's separator ends the line before (a no-break space ties it to `Verified`); the clause never starts with `·` | 09 |
| WP D | The officer captures read `Kraggor` in the site header (the stubs answer `/v1/me` as the officer); 01 and 08 re-taken | 01, 08 |
