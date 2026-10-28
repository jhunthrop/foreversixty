# Task 28 fixes (P1 to P5)

Built and tested from an isolated archive of HEAD plus these edits (/tmp/t28f), `FOREVER_DATA=fixture`,
because another session was syncing `web/public/data` in the worktree (the build failed there with ENOTEMPTY).
Previews ran from dist copies on private ports and were killed by PID.

- P1 `premium-examples.ts`: `with { type: 'json' }`. Astro build accepts it; Playwright loads all 3422 tests
  with no Node hook (the harness hook is no longer needed).
- P2 already fixed by c70596df0 (site-column.test.ts passes, no change made).
- P3 checkout LCP: the band's static render now draws the Guild summary frame (h1, size line, trial line),
  hidden by CSS unless the head script marked a guild URL (`html[data-checkout-step]`); Raid guild's size
  line is patched on DOMContentLoaded by the head script; name and price fill in on hydration into reserved
  lines. The page URL cannot be read at build (static site), so "server-side from the URL" was done through
  the head script signal. That alone gave 2585 ms; the rest was the hydration race the guild page already
  solved, so the island is now `client:after-paint`. lhci checkout median 1503 ms (7 runs: 1502 to 1503),
  main 2105.
- P4 row 17: not CopyLinkField. The markup is byte-identical. The block sits at a half-pixel different page
  offset because this branch's billing panel above it is 56.5 px shorter (`guild-settings-billing` 91 vs
  147.5 px), which moves 1 px lines in a screenshot. The capture now snaps the block's section to a whole
  pixel in both runs and ignores a 1-level channel difference (corner rasteriser noise, 2 px). Row 17
  passes at 390 and 1440 against a regenerated main baseline.
- P5 premium first screen: region A buttons end at y 517 at 1280x800 (was 589). Premium and Free price row
  112 to 84 px (premium content needs 82), card top padding 20 to 12 and audience gap 12 to 8 from 768,
  action slot 60 to 52 for every card (keeps the spec B3.2 equal button heights), hero bottom padding
  24 to 12 and body top padding 24 to 12. No font sizes changed; CARD/CARD_SURFACE untouched.

## Results

- `npm run check`: 0 errors. eslint on touched dirs: clean.
- vitest full (archive): 7218 passed, 3 failed only because the archive lacks addon/ and the report
  fixture; those 3 pass in the worktree. `_premium.test.ts` gap assertion updated 12 to 8.
- Playwright full, desktop + mobile: 3181 passed / 13 failed / 228 skipped. Failures: coach-account (8) and
  coach-report (4) are the coach lane; home-character-search (1) is pre-existing (fails on main too).
  The first full run's premium.spec 1024 equal-height failure was mine and is fixed (52 px slot).
- lhci full (22 URLs x 3): one failing assertion, planner perf 0.88 (main 0.89, pre-existing). Checkout passes.
- Captures regenerated: launch-pricing premium-value 25/25, subscribe-and-join 120/120.
