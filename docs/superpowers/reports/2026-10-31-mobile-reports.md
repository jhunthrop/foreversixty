# Mobile Wave 3 F10: Your reports (2026-10-31)

Branch `wt/mobile-reports`, lane `mobile-reports`. One lane, both platforms and the API. Spec:
`design/specs/2026-10-31-mobile-wave3.md` §5.2 and §10 (W3-C2 to W3-C4). Boards:
`design/mocks/2026-10-31-mobile-wave3/renders/f10-reports-{ios,and}`.

## API (W3-C2, W3-C3, W3-C4)

The web's report logic is ported into `api/internal/reports`, so every app figure comes from the
same file the web report page reads (the fight's stored `summary.json`) through the same arithmetic.

| Route | What it does |
|---|---|
| `GET /v1/reports/{id}/me` (session or app) | The caller's characters in the report and their boss kills. Each kill has the role figure: hps for a healer, dps for everyone else, as on the web's Summary tab. It also has the parse (placed as the web asks: the fight's phase, the row's spec, the value rounded to hundredths), `ranked` and `new_best`. |
| `GET /v1/reports/{id}/fights/{n}/meters?metric=&scope=trash` | Details!-style rows, sorted by the metric's per-second figure. Players with none of the metric are left off. `scope=trash` folds a zone's trash group. |
| `GET /v1/reports/{id}?fold=trash` | Adds `trash_groups` read from `report.json`. The database keeps no zone per fight. |
| Coach night, fight review, recent | The existing routes, now documented as app operations with fixtures and the 426 gate. |

- `fold.go` ports `fights.ts fightRows` (`TrashGroups`) and `night.ts aggregateNight` (`foldRosters`). It also defines `Figure`, the one rounding behind both Your parses and the meters. `TestMe_AgreesWithTheMeters` pins that the two print the same number.
- `rankings.Store.BestBefore` answers "earlier best kill", leaving out the report being read, for `new_best`.
- **Visibility.** All of these routes follow the report's own rules. A report you may not see is 404, never 403. A flagged report is 404 `flagged`. With no bucket the routes answer 503.
- **Dropped field.** The spec's `rank` on `/me` is left out. A rank among kills would contradict the character-board rank that F11 prints (the wow-player's "one character, one number").
- **Wiring.** `cmd/api` sets `Places: rankStore`, and `Objects: client` when R2 is configured. No migration.

## Fixtures: one story

`design/mocks/2026-10-31-mobile-wave3/f10_fixtures.py` writes every F10 fixture from one night:

- **The night.** Obnoxious Yell's guild Onyxia run, Monday 14 Dec in the Chicago evening. It has 9 fights: 6 trash pulls (7:40 in one group), a wipe at 61% (1:12), a phase-3 wipe at 22% (4:52) and the 4:28 kill (40 players, 3 deaths).
- **Your figure.** 190,923 damage over 268 s is 712.40 DPS. That is the 712.4 DPS and 91st percentile of `me.raid-night`; it is 6th of 40 on the meters.
- **Other fixtures.** `reports.list` (newest first, counts derived from the fights), `reports.get.live` (the live Blackrock Spire run on the list), `reports.meters.trash` and `coach.fight`. The raid night's coach notes now say what the fight coach says.

## Apps

| | iOS (`apps/ios/App/Reports`) | Android (`app/.../reports`) |
|---|---|---|
| Logic | `ReportsLogic`: list sections and row words, fight rows from `trash_groups`, pull numbers, wipe and kill meta, the parse ladder, meter bars, the 11-row cut and the largest-text 4 rows plus your row, coach sections, live polling | `logic/ReportsLogic.kt`, the same |
| Data | `ReportsStore`, `ReportStore` (report, parses, coach), `MetersStore`; disk cache cleared at sign-out | the same |
| Screens | list, report, meters sheet, coach card | the same, `ModalBottomSheet` with a drag handle |

- **Entry points.** Home's `Logs` row ("Your last raid: Onyxia's Lair, 1 kill", the newest complete report with a boss pull), the raid night's `Open the full report` row, and `/app/report/{id}` (new in `apps/shared/links.json`).
- **Coach.** The card is drawn only with `entitlements.log_coach`. The app reads the night line, then that line's pull review. A free player gets no card and no call.
- **Copy and switch.** New `reports.*`, `report.*`, `meters.*`, `coach.*` and `home.logs*` keys in `en.json`. The feature switch is `reports`.

## Reviews

Captures (default and largest text, both platforms, 64 files) are in
`design/reviews/2026-10-31-mobile-wave3-reports/build/`.

- **Round 1: both HOLD.**
  - ux-designer: the AX5 list gutter.
  - wow-player: Home's "last raid" named the live dungeon run; the live capture was drawn on the finished Onyxia report; your own row was lost at largest text.
  - The SHOULDs on largest-text sizing, flagged segments, pull numbers, fixture dates and deaths, and the coach notes were fixed in the same round.
- **Round 2: ux-designer SHIP (`ux-verdict-build.md`), wow-player SHIP (`player-verdict-build.md`).**

## Gates

Each gate was run on its own, and its exit code checked:

- `go test` on reports, rankings, coach, contract, server and cmd/api, against the private DB `forever_test_mobile_reports`.
- `npx vitest run src/lib/mobile`: the codegen drift tests.
- `make -C apps/ios test`.
- `./gradlew test :app:verifyRoborazziDebug lint`.

## Turning it on

`features.json` keeps `reports` default false: the codegen test holds every bundled default dark, and a feature turns on from the API. To turn it on, add `reports` to `APP_FEATURES` on the API deploy, the same step as F14. That is production config, so this lane did not do it.

## Gaps and follow-ups

1. **Meter names do not open a character page.** No character screen exists until F11. Rows already carry `key` for that.
2. **The trailing bar buttons are not in the iOS captures.** iOS snapshots do not draw iOS 26 glass trailing buttons (refresh, Open on foreversixty.gg). The same gap exists in the Raids captures.
3. **Signed-out list and deep-linked public report.** Both are built and captured, but Home gates signed-out players before the stack, so neither is reachable yet.
4. **Two sources for one parse.** `/me` and the meters read the stored summary, the web report page's source. F11's character history reads `fight_metrics`, which the ingest derives from the rebuilt events within 0.5%. A rounded figure almost always agrees. A shared source would need the ingest to publish the derived rows into the summary.
5. **Deferred SHOULDs:**
   - "was N DPS" beside New best, and "of 1,840" under the percentile.
   - Damage-taken bar scaling.
   - Home row chevrons (a shared pattern).
   - Meter active time and share of total.
   - Near-duplicate coach basis labels (only exact duplicates are removed).
6. **No push.** No push carries `report.ready` yet, so a report opens from Home, the raid night or the link.
