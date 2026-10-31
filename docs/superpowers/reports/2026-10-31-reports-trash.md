# Reports and trash lane: "Your reports", trash as one entry, themed scrollbars

Branch `wt/reports-trash`, from origin/main. Three owner-reported problems with logs and reports.

## 1. "Your reports" now opens a real page

The Logs menu linked "Your reports" to /reports. That URL was the report page's shell, so the
visitor saw "No report with that id".

- The shell moved from `web/src/pages/reports.astro` to `web/src/pages/report-shell.astro`
  (`dist/report-shell.html`). `src/worker.ts` SHELL_ROUTES now maps `/reports/` to
  `/report-shell.html`, so every `/reports/<id>` is served exactly as before. Bare `/reports` is not
  under `run_worker_first` (`/reports/*` does not match it), so it is a plain static asset.
- The new `web/src/pages/reports.astro` is "Your reports". It uses the /account look: a static
  PageHead (LOGS / Your reports / "Every fight you logged, newest first.") over one Card ("Reports").
  The card is the `YourReports.svelte` island (`web/src/components/logs/`), which reads the session
  and mounts the existing `MyReports.svelte`. That is the same list /logs and /account use, fed by
  `GET /v1/reports?mine=1` (session), with the same rows, guild tabs, paging and error state.
  - Signed in: the reports, newest first, each linking to `/reports/<id>`.
  - Empty: "No logs yet.", "Upload a log" pointing to /logs, and "Open the sample report".
    `MyReports` gained two optional props for this (`emptyMessage`, `emptySecondary`). /logs and
    /account are unchanged.
  - Signed out: the shared SignInPrompt (Battle.net, returning to /reports, plus the email link) and
    the sample report link. This is the server-rendered state, so a visitor with no session hint
    paints it at once. A visitor whose hint says signed in instead gets the list's 360 px loading
    reserve until the session answers.
  - Errors: a failed list shows MyReports' "Your reports did not load." with Try again. A failed
    session shows "Your account did not load." with Try again.
- Sitemap: /report-shell and /reports are both left out.
- Worker tests: the shell asset is now `report-shell.html`. A new test checks that bare /reports goes
  to the static assets and makes no API call. `_shells.test.ts` renders `report-shell.astro`.
- Playwright `tests/e2e/reports-page.spec.ts` covers: the menu target is a page, not the shell;
  signed out; signed in with a newest-first list; empty; list error with retry; session error with
  retry; and `/reports/fixture2abcd` still opening the report.

## 2. Trash is one entry per zone

- `lib/report/fights.ts` `fightRows` now works zone by zone, in the order the night reached each
  zone. Every boss pull keeps its own row, then that zone's trash becomes one `trash-group`.
  Previously only consecutive runs were folded. A zone with a single trash pull still gets the
  entry, and the entry opens that pull itself. Zone-less trash keeps its own rows, as it does today.
- The FightSelector entry reads "Trash · <zone>" with its pull count, total duration and deaths.
  - Selecting it (`fight-trash-<key>`) opens the combined view.
  - The pulls sit behind a separate chevron button (`fight-group-<key>`, `aria-expanded`). That
    button is held open (disabled) while one of its pulls is on screen.
  - The "Show N trash fights" toggle is gone.
  - The card head reads `bossPullsLabel`: "6 boss pulls", "1 boss pull" or "No boss pulls". It
    never counts trash.
- The combined trash view is built the same way as "All boss pulls":
  - `url.ts` adds `trashScope(key)` / `trashScopeKey`. The URL spells it `fight=trash-<first pull>`,
    and the state holds it below MISSING_FIGHT, so `fight` stays one number. `resolveFightIndex`
    accepts it, and a stale `trash-N` link says so and falls back.
  - `combinedFights(fights, selected)` returns the pulls to fold: all boss pulls for ALL_FIGHTS, or
    the zone's trash for a trash scope. ReportView's `nightMode` is now "a combined view is open",
    so every existing night rule applies to trash too: tables only, no per-second chart, the
    "one pull's" modes greyed out, and Ratings say to pick a pull. The fold effect loads and folds
    `combined` instead of the encounters.
  - `aggregateNight` / `nightSummary` fold the pulls the caller hands them. The boss-only filter now
    sits with the caller, and only boss pulls make a Bosses row or a mechanics boss. Trash pulls are
    numbered "Trash · pull 2 of 3".
  - NightView with `trash` shows "N trash pulls · time in combat · deaths", has no Bosses table, is
    headed "Players, across every trash pull", and shows no Kill or Wipe per pull. In the trash
    view the tables' Parse cells use the trash fallback rather than the night's. The boss filter has
    no bosses there.
- Checked in the browser: the trash view's Summary, Damage Done and the other tables fill from the
  merged trash summaries (captures `trash-combined-*`). As with the night view, no time chart is
  drawn over a combined view.
- The sample (79 trash pulls, no boss) shows "No boss pulls" and one "Trash · Zephras Isle" entry.
  It still opens on its longest pull, so the entry is expanded. On a phone the list stays folded to
  its one row, so the phone fold promise holds (`onboarding-second-screen` passes), and the
  sample's CLS specs pass.
- Also fixed: the selected fight row's bottom border took `currentColor` (a white rule under the
  gold edge). It now uses the soft line colour.

### Merging with wt/logs-engine-counts

That lane edits the same trash-group block (it adds `kills` and `fightDeathMark`, and drops
`deaths`). I kept the `TrashGroup` field names (`key`, `zone`, `fights`, `duration_ms`, `deaths`)
and the `fight-group-<key>` test id. Whichever lane merges second resolves it this way:

- In `trashGroup()` in fights.ts, swap `deaths` for `kills: sumKillCounts(fights)`.
- In the new entry's meta line in FightSelector, swap the deaths mark for their kill-count span and
  `groupMark`.
- Their e2e `trash from before engine 0.6` clicks `toggle-trash`. That line has to go, since the
  toggle no longer exists. Their `fight-group-1-kills` test id can stay on the entry's meta line.

Their fights.test "folds consecutive trash" case conflicts with my rewritten version. Keep mine and
add their `kills` expectation to it.

## 3. Themed scrollbars

`web/src/styles/global.css` @layer base now has `* { scrollbar-width: thin; scrollbar-color:
var(--color-line-warm-strong) transparent; }`, plus `::-webkit-scrollbar` rules for Safari: 8 px, a
transparent track, a 4 px-radius thumb in the same token, and `--color-muted` on hover. Strips that
hide their bar (`scrollbar-width: none`, unlayered) still win. The capture launches Chromium with
macOS's per-process `-AppleShowScrollBars Always`, so the bar shows the way it does with a mouse
plugged in.

## Gates (exit codes checked, no pipes)

- vitest, full web suite: 689 files, 7772 tests, exit 0
- `npm run check`: 0 errors. `npm run lint`: exit 0. `npm run format:check`: exit 0
- worker tests (`src/worker.test.ts`): 47 passed (included in the full vitest)
- island size: `check-island-size` passes in postbuild (report-island 127.9 KB gzipped)
- Playwright, full suite, desktop and mobile, fixture build on port 4707:
  - Run 1, default build: 3412 passed, 410 skipped, 0 failed.
  - Run 2, rebuilt and served with `PUBLIC_API_BASE_URL=http://127.0.0.1:9` (the coordinator's
    hermetic rule): 3382 passed, 30 failed, 410 skipped. All 30 failures are in specs this lane does
    not touch (account-discord, account-discord-link, guild-discord-bot, guild-settings-discord,
    header-account, header-finished-sim, home-panel, home-feeds-settle, report-queries). They
    hard-code `https://api.foreversixty.gg` in href assertions or route globs, or rely on an
    unstubbed production call. Under run 1 they all pass. The hermetic lane owns them. Every spec
    this lane added or edited passes under both builds and stubs its own API calls.

## Captures

`design/reviews/2026-10-31-reports-trash/`, from
`web/tests/captures/reports-trash/capture.spec.ts`:

- reports-{signed-out,empty,list,error}-{1440,390}
- fights-trash-collapsed-1440 and fights-trash-expanded-1440
- trash-combined-{1440,390} and trash-combined-damage-1440
- raid-night-{1440,390}
- scrollbar-fight-list-1440 and scrollbar-page-1440

## Concerns

- A combined view (trash or All boss pulls) keeps the report shell's desktop min-height
  reservation, so a short combined table leaves space above the footer. This predates this lane:
  the night view has always done it.
- The trash entry sits after its zone's boss pulls. On a night that clears trash first, the entry
  comes after the bosses, not in time order. That is deliberate (the Warcraft Logs layout).

## Finish round (2026-10-31)

- **Reviews.** ux-designer SHIP and wow-player SHIP, both on round 2; verdicts in
  `design/reviews/2026-10-31-reports-trash/verdicts.md`.
  - ux round 1 blocked on a stray scrollbar thumb on the mode tab row. Fixed with a `scrollbar-none`
    utility on sideways tab rows.
  - wow-player's owner follow-ups are logged in the verdicts: report-level fight totals counting
    trash, the "mobs died" wording, the Active % denominator, and the list row affordance.
- **Merge.** Merged origin/main, including logs-engine-counts. The trash entry now carries its
  summed kill counts and the group death mark. The `toggle-trash` e2e steps are gone.
- **Gates.**
  - check, lint, format:check and the full vitest pass (7783 tests).
  - Full Playwright, built and served with `PUBLIC_API_BASE_URL=http://127.0.0.1:9`: 3387 passed,
    27 failed.
    - The account-discord(-link), guild-discord-bot, guild-settings-discord, report-queries and
      guild-raid-spec-ask failures also fail on a clean origin/main built the same way.
    - home-panel's late-shift check failed under full-suite load. It passes in isolation, run
      twice.
  - lhci: the reports URL passes. bis/marksmanship LCP, sim and planner miss locally on both this
    branch and clean main (machine noise; this lane's only change there is the scrollbar CSS).
