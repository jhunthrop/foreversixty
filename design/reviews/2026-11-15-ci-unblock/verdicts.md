# CI unblock, 2026-11-15: persona verdicts

Captures from the e2e build on the test box (fixture data, signed out for /setup and
/remove-my-character, the unknown-realm account from tests/e2e/setup-unknown-realm.spec.ts).

| Surface | ux-designer | wow-player | ai-slop-reviewer |
| --- | --- | --- | --- |
| /setup Commands (addon 2.0.55 list, capped command column) | SHIP | SHIP | PASS |
| /remove-my-character failure line ("That did not go through. Try again.") | SHIP | SHIP | PASS |
| Unknown-realm question (hairlines from ui/Ruled) | SHIP | SHIP after recapture | PASS |

## Round 1 findings

- The first unknown-realm captures came from a build before the Commands column cap, so their
  Commands list overflowed at 390. All three reviewers traced it to the stale build. The
  unknown-realm/ captures here were retaken on the capped build: question-390.png is 390 wide with a
  clean Commands list.
- The empty boxes under "What it does in the game" in the full-page captures are lazy images that
  had not loaded when the shot was taken. That is how the capture works, the page is fine.
  handoffs-setup.spec.ts checks every image loads once scrolled to, and it passes.

## Notes, not blockers

- At 390 the description column wraps short lines onto two rows, because the command column takes
  up to 45% of the rail.
- "the Set up steps on Me again" is the addon's own Locale.lua text, and the site mirrors it.
