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
