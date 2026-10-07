# Hit to cap on the BiS page (2026-10-07)

Lane `hit-cap-web`, branch `hit-cap-web` from `main`. Web only; no data files touched.

## What changed

- `BisBand.hit_to_cap?: { baseline, specials, white? }` is declared in `types.ts` and read through
  `normaliseBisFile`, which defaults an absent key to `null`.
- `lib/bis/hit-cap.ts` holds the one formatter (`hitCapLine`), shared by the BiS weight rail and the
  guides' stat table. Whole percents print bare, fractions to one place. It returns `undefined` for an
  absent key, and callers then render nothing.
- BiS page: one line under the weights table, "Hit to cap: N% for specials[, M% for white swings]".
  The tooltip says hit is worth its full weight until the cap and nothing past it, and that the weights
  are per rating point (10 hit rating and 14 crit rating are 1%).
- Guides: `GuideStatTable` shows the same line under the table, through `band60Weights` (new `hitToCap`).
- Preset caption: every BiS band's caption now carries one static sentence saying the sim fights a
  level-63 target of no creature type over 180 seconds, with slaying bonuses counting for nothing
  (coordinator addition). Static copy in `bisCopy.simTargetNote`.
- Fixture `warrior-protection.json`: alliance raid carries `{3, 6, 25}`, alliance bare `{3, 6}` (no
  white), horde carries none, so both shapes and the absent case are covered.

## Tests

Unit: `hit-cap.test.ts`, `panel-view.test.ts` (both shapes, absent, null), `load.test.ts`,
`class-dps.test.ts`, `GuideStatTable.test.ts` (line present exactly when real data has the key).
Playwright `bis-preset.spec.ts`: the line per preset with its tooltip, the absent case on horde, the
simulation-target sentence on both presets.

## Screenshots

`design/reviews/hit-cap-web/bis-preset-{raid,bare}-{1440,2000}.png`. The line is one 12px muted
paragraph inside the existing panel flex column; the table and panel heights above it do not move.

## Notes

- `npm run sync:rotations` (the `pretest`/`prebuild` step) fails on main: hunter-survival step 4's
  player-facing notes carry the date 2026-10-07 (`rotation-apl.mjs` guard). Not in this lane's files.
  `astro build` and vitest were run directly.
- `bulk-store.test.ts` (5 tests) and `rotation-apl.test.ts` fail in this worktree for the same
  reason or stale symlinked generated data; neither touches this change.
- `bis-preset.spec.ts` "below level 60 there is no control" counts preset options on the whole
  hunter/marksmanship page, which now carries a level-60 control from real data; not changed here.
