# /logs build report — logs landing rebuild (2026-10-04)

Branch: `web-logs-rebuild`, worktree `.worktrees/web-logs-rebuild`, based on `main@ab279af8`.
Spec: `design/specs/2026-10-04-logs-landing.md`, amended by `design/reviews/2026-10-04-logs-mock-wow-player.md`
and `design/reviews/2026-10-04-logs-mock-ux-designer.md` (both folded in below; see main's
own `design/specs/2026-10-04-logs-landing.md` §12 for the amendment text, which my worktree
predates and does not carry a copy of).

## What was built, by spec section

**§4.A Header band (new).** `LogsHeroBand.svelte` + `LogsHero.svelte`, mounted `client:idle`
as the first child of `<main>` in `logs.astro`. Full-bleed `100vw` breakout reusing
`PlannerHeaderBand.svelte`'s own technique (no second way of doing the same thing). Hero
selection is exact per spec: `listMyReports(1).rows[0]` when signed in with a report, else
the canonical sample (`data/sample-report.json`, id `5lop7n5kwysf`) via the public
`GET /v1/reports/{id}`. Eyebrow, h1 (title/zone fallback), `pill-sample` only on the sample
(kept even though the sample's own title already says "sample", per the designer's ruling),
facts line (`ReportRow.svelte`'s own fragment, now shared via `lib/reports/hero.ts`'s
`reportDay`), description sentence with its two `text-nav underline` anchors, `Open this
report` as `SECONDARY_BUTTON_FIXED` with a gold border/text (never the filled gold button —
designer ruling 2).

**§4.A.1 Per-player hook (new).** Signed-in hero only, a deferred third fetch
(`fetchReportMeta` → last-kill fight → `fetchSummary` → roster match by character name,
else top dps → `plannerLinkFor`/`simLinkFor` reused unchanged). Line format per the
designer's amendment: `Last kill, {fight}: {Name} · {dps} DPS · Gear in the planner | Sim`,
dps one decimal under 1,000 / thousands-separated whole number at or above. Sample hero
shows the static sentence instead and never runs this fetch. Failure is silent (no error
banner) and never blocks the hero's own primary content.

**§4.B Spine.** `CurrentCharacterBar spine` unchanged, directly under the band.

**§4.C Your reports / Recent public reports.**
- Guild tab strip (`MyReports.svelte`): shows when `me.guilds.length > 0`, "Mine" + one pill
  per guild, backed by the new `listGuildReports(guildId, page)` client
  (`GET /v1/guilds/{id}/reports`, confirmed merged on `main` mid-build — same `MyReportPage`
  shape and page-number pagination as `listMyReports`, not the cursor the task brief first
  guessed). Empty copy `No {guild.name} reports yet.`
- `RecentReports.svelte` filters the sample id out of its own rows before rendering, so its
  existing, already-correct empty string is reachable.
- **§4.C.3 data defect, filed, not fixed:** the canonical sample report
  (`5lop7n5kwysf`) is confirmed (by the coordinator) to be a real retail dungeon whose
  roster carries a Mistweaver — not a Classic Forever class
  (`forever-class-rules`: nine vanilla classes + six new race/class pairs, no Monk). Per
  the coordinator's ruling the build keeps the real title and the Sample pill and does not
  alter the data. **Blocking fix needed before this is a reference example**: regenerate
  that report's roster with valid Classic class/spec rows, or point
  `data/sample-report.json` at a different already-public report with none.

**§4.D Companion / Upload.**
- `CompanionStatus.svelte`: one line per device, `{name} · {platform} · last seen
  {relative}` / `· paired, not seen yet`, plain text, never "Connected", never a coloured
  dot (designer ruling 5). Mounted above the existing pointer line.
- `Account.svelte` pairing mode: a 5s poll (`lib/reports/pairing-poll.ts`'s
  `createPairingPoll`, fed by the new `refreshDevices` — `listDevices` alone would have kept
  handing the poll the same cached array forever under `query()`'s "already loaded this
  page" rule, caught and fixed during this build) ends on a new device id, swapping the code
  block for `{device.name} paired. It starts uploading as soon as you are logging.` in
  place; the existing expiry path is unchanged. Upload panel kept verbatim.

**§7 Phone order / designer-ruling fixes folded in:**
- The dungeon-framing sentence sits directly under the spine on every viewport, before
  "Your reports" (ruling 4).
- `ReportRow.svelte` stacks its facts line under the title below 640px instead of letting a
  title and a pinned date risk sharing one line (player-review finding 3).
- No "Newest first…" caption anywhere (ruling 7 — was never added).
- Sign-in actions inside `Your reports` and elsewhere use the existing outline
  `SignInPrompt.svelte` only; no filled-gold button added to this page (ruling 2).

**§8 Performance/CLS — one real regression found and fixed during this pass.** The first
skeleton height I shipped (a flat 132px) did not match either the real ready hero (measured
292px at 360px/mobile, 228–232px at ≥1024px) or the `LoadError` branch (44px) — a direct
Lighthouse audit of `/logs.html` caught a CLS of **0.0518**, over the 0.05 budget, from the
skeleton→content (and skeleton→error) swap. Fixed by moving the reservation onto the hero's
own outer wrapper (`min-h-[292px] lg:min-h-[232px]`) so the loading, failed and ready
branches all share one floor — re-measured equal (292/292 and 232/232 px) at both
breakpoints, and the direct `/logs.html` Lighthouse re-run came back **CLS 0** on all three
runs. A new e2e test (`logs-hero.spec.ts`, "no layout shift on hydration") pins this: the
hero's bounding-box height must be identical in the failed and ready states, at 360 and
1280px.

## Gates

**Unit suite (`npx vitest run`):** `290 files passed, 3033 tests passed, 1 skipped` — exit
0. 28 of those tests are new, across `hero.test.ts`, `hero-hook.test.ts`,
`device-status.test.ts`, `pairing-poll.test.ts`, `copy.test.ts`: hero selection, hook row
selection (character match → top dps), dps formatting, the sample-id filter (covered
indirectly through `RecentReports.svelte`'s own existing test plus the e2e sample-feed
test), device status copy (both states), the pairing poll (new id ends it; a short expiry
stops it with no success; `stop()` is immediate), and the guild-empty/pairing-success copy
strings.

**Full Playwright suite, both projects, no file filter (`playwright test --project=desktop
--project=mobile`):**
- First full run (before the CLS fix): 1193 passed, exit 0.
- After the CLS fix: 2 new mobile-only failures surfaced
  (`logs-recent-reports.spec.ts`'s spine-bar-mounts test,
  `logs-upload.spec.ts`'s signed-out-sign-in test) — root cause confirmed by direct
  measurement: the header band legitimately grew taller on a short mobile viewport (the CLS
  fix above), which meant a single `scrollIntoViewIfNeeded()` on a sibling heading no longer
  reliably carried a `client:visible` island's own root into view. Fixed by scrolling to the
  **next** panel down instead (verified empirically with a throwaway diagnostic spec, then
  deleted) — both tests pass on mobile and desktop afterward.
- Re-run after that fix: 1195 passed, exit 0 (see `/tmp/e2e-full2.log`; not kept past this
  session).
- Final clean run after all four commits: **1197 passed, 67 skipped, 0 failed, exit 0**
  (`/tmp/e2e-full4.log`). The run immediately prior to it (same commit state, same command)
  came back 1196 passed / 1 failed / exit 1, and the one failure
  (`sim-tabs.spec.ts`'s "active tab scrolled into view" test, a sub-pixel `390.34375 >
  390` rounding check) is unrelated to `/logs`, pre-existing, and non-reproducible: repeated
  3/3 pass in isolation immediately after, and absent entirely from this final run.
- New specs added this pass: `logs-hero.spec.ts` (6 tests: 4 hero-selection/hook/sample-
  filter cases + 2 CLS-regression-guard cases), `logs-guild-tabs.spec.ts` (2),
  `logs-companion.spec.ts` (4, including the full code→poll→success sequence over real 5s
  ticks), `logs-widths.spec.ts` (10: 5 widths × signed-out/signed-in, the
  `scrollWidth === clientWidth` assertion at 1024/1100/1280/1440/1920).
- One pre-existing test-fixture bug fixed as a side effect: `logs-upload.spec.ts`'s
  `/v1/devices` mock returned `{ devices: [] }` instead of the real API's bare `Device[]`;
  harmless until the pairing poll started reading `devices` itself in `onPair`, at which
  point `devices.map` on a non-array silently swallowed every "Show pairing code" click.

**Lighthouse.** Official gate per the task's hard rule — `FOREVER_DATA=fixture
LHCI_PUSH_SHARD=2 npm run lhci:push` (shard 2 is the one carrying the
`(?:planner|logs|setup)\.html$` assertMatrix pattern /logs.html itself matches, per
`lighthouserc.push.cjs`'s own documented shard-URL mapping): **0 assertion failures, exit
0**, both before and after the CLS fix (shard 2's own representative URL is `planner.html`,
unaffected by `LogsHero.svelte`). Median `planner.html`: performance 0.97, accessibility 1,
SEO 1, best-practices 0.96, TBT 68ms, CLS 0.0313 — all well inside budget.

Supplementary, beyond the required gate: a direct Lighthouse audit of `/logs.html` alone
(not part of the official sharded command, run for my own verification of the CLS fix).
Its assertMatrix entry (`.*/(?:planner|logs|setup)\.html$`) carries no LCP assertion at all
— only performance ≥0.90, accessibility/SEO ≥0.95, best-practices ≥0.95 (warn), TBT
≤100ms, CLS ≤0.05 — so the spec text's own "LCP ≤2200ms" is aspirational for this page,
not something `lighthouserc.json` currently enforces. Before the fix: CLS 0.0518 (fails).
After: **performance 0.93–0.95, accessibility 0.96, SEO 1.0, best-practices 0.96, TBT 0ms,
CLS 0** across 3 runs — passes with margin on every assertion this pattern actually makes.
LCP read 2.7–3.0s in this sandbox specifically because `/v1/me` et al. resolve against the
real `https://api.foreversixty.gg`, which this sandboxed environment cannot reach
(`statusCode: -1` in the Lighthouse trace) — every hero fetch fails and the hero settles
into its `LoadError` branch before paint. That is an artifact of this environment lacking
outbound network access, not a measurement of the hero's real-network LCP; I could not
verify the real-network number from here. The official CI runner (GitHub-hosted, real
internet access) should see the true success-path LCP; worth a follow-up check there.

## Captures

`design/mocks/renders/logs-build/` (git-ignored, 33 files), driven against the real
`FOREVER_DATA=fixture` build via `npm run preview` (never a fixture-build screenshot trick),
with `/v1/me` routed to a hand-built subset of `src/fixtures/me-addon.ts`'s Zulmara fixture
for signed-in states and the `fs_csrf` cookie set so the spine's pre-paint rule shows it
(caught during this pass — my first capture pass omitted the cookie and silently rendered
every signed-in board with no spine bar at all; re-run once I noticed it on inspection):

- `signed-out-{360,390,1024,1280,1440,1920}.png` — full page, sample hero, sign-in prompt,
  honest empty `Recent public reports`
- `signed-in-reports-{360,390,1024,1280,1440,1920}.png` — full page, own newest report as
  hero with the resolved per-player hook line, real `planner`/`sim` hrefs
- `signed-in-zero-{360,390,1024,1280,1440,1920}.png` — full page, sample hero (identical to
  signed-out's), "Your reports"' own empty state with `Upload a log`
- `guild-tab-active-{1440,390}.png` — the guild pill active, that guild's own rows
- `companion-status-seen.png` / `companion-status-never-seen.png` — both device-status
  copies, cropped to the status block
- `pairing-code-shown.png` / `pairing-polling.png` / `pairing-success.png` — the full
  code→poll→success sequence, cropped to the pairing block
- `upload-idle.png` / `upload-chosen.png` / `upload-progress.png` / `upload-error.png` —
  cropped to the upload panel (`upload-error.png` is from the pass before the `fs_csrf` fix;
  its own content is unaffected by that cookie, so it is still valid)
- `hero-skeleton-{1280,360}.png` — the loading state, reserved height confirmed equal to
  ready/failed (see the CLS fix above and its e2e regression guard)
- `hover-open-this-report.png` / `focus-visible-guild-pill.png` — the two states tenet 9
  asks for

**Side-by-side (required deliverable):**
`design/mocks/renders/logs-build/side-by-side-1440.png` — my `signed-in-reports-1440.png`
beside the regenerated mock `design/mocks/renders/logs-signed-in.png` (that mock file lives
on `main`, past my worktree's branch point; read directly off disk rather than merged in),
both at 1440px, same scale, padded to a common canvas height. Structural match is strong:
band layout, facts line, description/anchors, hook line format (`Last kill, {fight}: {Name}
· {dps} DPS · Gear in the planner | Sim`), spine position, framing line position, the two
report-list panels, guild tabs, companion panel order (status → pointer → pairing), upload
panel. The only differences are expected ones: the mock's illustrative guild name/report
rows (`Olympus XXVII`, three rows) versus my fixture's real ones (`Sample Guild`, one row);
the mock's companion panel defaulting to a shown pairing code versus the real app's
interactive idle state (captured separately above); and the spine's own door labels/count
(`PLAN SIM LOGS RANKINGS FOR HUNTER`, no Forget shown at this width) — `CurrentCharacterBar`
is the existing, unchanged component per §4.B, not something this pass touched, and the
mock's own spine row is an illustration, not literal output of that component.

I reviewed this side-by-side myself before writing this report, per the hard rule.

## Concerns / follow-ups

1. **Blocking data defect, §4.C.3**: the sample report's Monk roster row — filed above, not
   fixed here per the explicit task instruction not to fake it away.
2. **LCP for `/logs.html` is not actually gated** by `lighthouserc.json`'s own assertMatrix
   (confirmed by reading the config, not assumed) — the spec's §8/§11 "LCP ≤2200ms" is
   real-world guidance the build should still meet, but nothing in CI currently fails if it
   doesn't. Worth raising with whoever owns `lighthouserc.json` if that's unintentional.
3. **LCP could not be verified against a reachable API** from this sandbox; the real-network
   number should be spot-checked once this merges and CI (or a developer with real network
   access) runs the nightly Lighthouse audit, which does collect `/logs.html` directly.
4. `GET /v1/guilds/{id}/reports` was still being built in a parallel worktree
   (`api-guild-reports`) when I started; I initially shaped `listGuildReports` around a
   cursor per the task brief, corrected to page-number pagination once the coordinator
   confirmed the merged endpoint's real shape, before writing any client code against the
   wrong one.
5. `web/src/components/report/ReportView.svelte`'s own `resolveTreeSizes`/`loadTalents`
   pattern was reused unchanged for the hook's talent-tree-size lookup (one class per hero,
   not the whole roster) — no new network contract invented there.

## Files touched

Lib/data (commit `e1ae4d19`): `web/src/data/sample-report.json`,
`web/src/lib/account/api.ts`, `web/src/lib/reports/{copy,hero,hero-hook,device-status,
pairing-poll}.ts` + their `.test.ts` files.

Components (commit `3af6a43d`): `web/src/components/{LogsHero,LogsHeroBand,
CompanionStatus}.svelte` (new), `web/src/components/{Account,MyReports,RecentReports,
ReportRow}.svelte` (changed).

Page wiring (commit `623038de`): `web/src/pages/logs.astro`, `web/src/pages/_logs.test.ts`.

e2e (commit `9cbc1a8e`): `web/tests/e2e/logs-{hero,guild-tabs,companion,widths}.spec.ts`
(new), `web/tests/e2e/support/logs-hero-fixture.ts` (new),
`web/tests/e2e/logs-{upload,recent-reports}.spec.ts` (changed).

Captures: `design/mocks/renders/logs-build/*.png` (git-ignored, not committed).
