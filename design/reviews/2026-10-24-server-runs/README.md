# Server runs review: "Your server runs" and the done-while-away line

Date: 2026-10-24. Plan: `docs/superpowers/plans/2026-10-24-server-runs.md`.
Experience spec: `design/specs/2026-10-24-server-runs.md` (ux-designer). Product spec:
`docs/superpowers/specs/2026-10-24-premium-experience-design.md` §4.3, §5.
Mock: `design/mocks/gen_server_runs.py` → `design/mocks/renders/server-runs-states.png` (git-ignored).

Captures taken from the fixture build with the API stubbed, at head b6690e3d:

- `sim-payer-states-{390,1024,1440,2000}`: running with an estimate, running without, done unseen,
  done seen, failed with Run again. `sim-titled-rows-1440`: the same with member titles.
- `sim-empty-*`, `sim-loading-*`, `sim-error-*`: the list's empty, skeleton and failed reads.
- `sim-overflow-1440` (six rows → the footer line), `sim-overdue-1440`, `sim-starting-1440`,
  `sim-reduced-motion-1440`.
- `sim-failed-runagain-{rest,hover,focus,pending,started,couldnotstart,lapsed}-*`: every Run again
  state, the last one the 402 `See plans` variant. `sim-done-row-{hover,focus}-*`.
- `home-line-{one,three,dismissed}-{390,1024,1440,2000}`, `home-line-focus-{link,dismiss}-390`,
  `bis-line-1440`, `rankings-line-1440`, `logs-line-1440`: the band line on every page that mounts
  the shared character header.
- `account-runs-{4,1,none,nocount}-*`, `account-off-*`: the account page's Server sims row.

Measured CLS, hinted second visit: `/sim` list 0.0006 (390) / 0.0004 (1440); home line
0.0059 (390) / 0.0002 (1440); the spec's limit is 0.02.

## Verdicts

- ux-designer: SHIP WITH FIXES on the first pass (the done-row `›` and the dismiss `×` rendered at
  6–8 px; the skeleton lacked crests, separators and the state bar; the error state collapsed the
  reserve; five captures missing). Fixed in commit b6690e3d and re-captured; re-check: **SHIP**.
  Accepted: phone rows put state and age on one line; the list sits above the run block; running
  and failed rows are not links; the line renders in the band layout only; the reserve stores the
  rendered px height (spec §3.7 amended). Report: `ux-designer-review.md`.
- wow-player: SHIP WITH FIXES. Of the eleven items, four were fixture artefacts (no character
  loaded, untitled runs, two different fixtures' DPS, a guild-plan account's "Not on Premium" line),
  one a misread of low-resolution text (the copy is `Could not start it again. Try again.`), four
  are by spec (no estimate under three finished runs, `Open result` dropped on the phone, the
  unread dot plus tint, the fixed error copy), and the glyph size was fixed with the ux-designer's
  item. Report: `wow-player-review.md`.

## Deferred, not this lane

- Name the character on the band line when the finished run's spec differs from the band's
  character (wow-player 6); rows could carry the character or gear set once the API sends one
  (wow-player 2).
- A visibly distinct unread mark beyond the dot and tint (wow-player 3), if a later pass finds it
  missed in use.
- Fight context on the DPS headline (fight length, target count) is the API headline's to grow.
- "Supporter through <guild>" on the account plan line for guild-plan members (from lane 1).
- `SimView.svelte` (883 lines) and `CharacterHeader.svelte` (933) were over the 800-line cap before
  this lane; the additions are small, the split is a maintenance pass.
