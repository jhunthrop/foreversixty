# Discord notifications: acceptance captures

Head sha captured: 8822bf07 (branch worktree-lane-discord-1). Fixture build (`FOREVER_DATA=fixture`), `/v1/me` and the three notification routes stubbed, Chromium at device scale 1, reduced motion on (so fades and hovers show their end state).
Spec: `design/specs/2026-10-25-discord-notifications.md` section 9. Item 9 (real Discord messages) is the controller's.

Regenerate: build and `astro preview --port 4411`, then from `web/`
`npx playwright test -c tests/captures/discord-notifications/playwright.capture.config.ts`
(the spec and config live under `web/tests/captures/discord-notifications/` so `@playwright/test` resolves), then `python3 design/reviews/2026-10-25-discord-notifications/compose_side_by_side.py` after `python3 design/mocks/gen_discord_notifications.py`.

## Files

`discord-<state>-<width>.png` is the full page; `discord-<state>-panel-<width>.png` is the Discord panel with 24 px padding (the free sentence for `free`).

| Item | Files (widths) |
|---|---|
| 1 saved, Premium row `On · 2 events` | `saved` (360, 390, 768, 1024, 1440, 2000), with panel crops |
| 2 paste form | `paste-empty` (input focused), `paste-invalid` (typed value in the field, absent from the message), `paste-not-entitled` (402 refusal with See plans), `saving` (`Saving…`) |
| 3 saved states | `saved-just-now`, `sending`, `test-sent`, `test-failed` (Discord error in wipe), `test-rate-limited`, `events-saved` (`Saved.`), `warning` (2 failures), `warning-long` (4 failures, 140-char Discord error; also 360) |
| 4 switched off | `switched-off`, `switched-off-long` (also 360), `switched-on-again` (after Test and switch on) |
| 5 lapsed | `lapsed` (also 360) |
| 6 remove / replace | `remove-confirm`, `removed`, `replace` |
| 7 free | `free` (360, 390, 768, 1024, 1440, 2000; full page and sentence crop), `feature-off` (deployment without the `notifications` key) |
| 8 focus-visible, 1440 | `discord-focus-<control>-1440.png`: input, save, checkbox-sim-paste, checkbox-sim, checkbox-upload, test, remove, replace, see-plans (panel or free-sentence crops) |
| 8 hover, 1440 | `discord-hover-save|test|remove-1440.png` (panel crops) |
| 10 Premium row | `discord-row-<value>-<390|1440>.png` for off, off-set-up, off-no-events, off-failed-five-times, paused-plan-lapsed, on-1-event, on-2-events |
| 11 | 1024, 768 and 360 are in the `saved` set; 360 also for warning-long, switched-off(-long), lapsed |
| 13 | `discord-side-by-side-1440.png` (mock left, shipped right, 1:1) |
| skeleton | `discord-skeleton-390.png`, `discord-skeleton-1440.png`: `/v1/me` held, session cookie hint set, so the page draws its loading skeletons |

139 PNGs. 360, 390, 768 and 1024 widths other than the ones in the table were not asked for.

## CLS (item 12, Task 6 report, desktop project)

| Width | load | save | test | remove |
|---|---|---|---|---|
| 390 | 0.352 | 0 | 0 | 0 |
| 1440 | 0.077 | 0 | 0 | 0 |

Mobile project: all 0. Load-time footer shift is not zero: at 390 it is 0.4476 for free, empty and saved alike (account skeletons giving way to content); at 1440 free 0.040, saved 0.083, empty 0.098, so the Discord block adds about 0.04 to 0.06 because the skeletons do not reserve its height. Lighthouse audits signed out only.

## Noticed while capturing

- Fixture artifacts, not the lane: the stub leaves `/v1/reports` unanswered, so Your reports shows "did not load"; `mockEntitlements({ notifications: true })` alone leaves the Plan panel reading "Not on Premium. See plans" beside a live Discord panel and an `On · 2 events` row.
- The long Discord error is cut mid-word at 140 characters and ends `Mana….` (ellipsis followed by a full stop) in both the switched-off notice and the 4-failure warning.
- The 1-to-4-failures warning is plain muted text, while switched-off has a coloured bar and a heading line; in the 4-failure state the error sentence reads as small print.
- Checkboxes are the browser's bright blue native control (and grey when disabled in lapsed); the mock draws a paler blue tick.
- The channel id is set in the body font with a bolder weight; the mock draws it in a monospace face.
- Saved: about 70 px of empty space between the buttons and `Use a different webhook` (the reserved status line plus gaps); the mock shows the same reserve at a smaller gap. In switched-off the gap above `Use a different webhook` is of the same size.
- In switched-off the Post when rows sit below `Use a different webhook`, as the mock draws them; the panel is the tallest state (see `switched-off-long-panel-360.png`).
- At 2000 the page is centred (content 376 to 1624) and the rail is the right column; nothing stretches.
- The mock's Plan panel shows Premium renew data (example); the shipped one cannot match it in the fixture.
- The loading skeleton (`discord-skeleton-*`) has no Discord reserve; the lane deferred this, see CLS above.

## Verdicts

## Deferred, not this lane
