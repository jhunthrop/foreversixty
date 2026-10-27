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
