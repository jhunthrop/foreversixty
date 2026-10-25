# Discord notifications: acceptance captures

Head sha captured: c0e41c7d (design fixes) plus 1660ae0a, branch worktree-lane-discord-1; re-taken after the design fix round (first take was 8822bf07). Fixture build (`FOREVER_DATA=fixture`), `/v1/me` and the three notification routes stubbed, Chromium at device scale 1, reduced motion on (so fades and hovers show their end state).
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

- `saving` and `sending` now hold the PUT and the POST until after the frame: `Saving…` and `Sending…` show with the checkboxes disabled (grey) and the buttons in their busy labels. The `Sending…` button draws a light focus-style border, and in `saving` there is a blank gap of about one line between the URL input and the In Discord helper text.
- Free sentence: two lines at 1440 and two at 390 (`...or a log you uploaded is done. See plans`), the same 92 px crop height at both widths; no jump between widths.
- Saved panel: the gap above `Use a different webhook` is now about 50 px (was about 70), still larger than the mock's, which is about 40 px.
- Warning (2 failures) reads as two lines: `The last 2 posts failed. At 5 in a row it switches off.` then `Last error from Discord: Unknown Webhook.`, in the same muted small text.
- The long Discord error now cuts at a word and ends `someone with…` with no stray full stop (switched-off-long at 390: 606 px tall with the Post when rows below `Use a different webhook`; warning-long at 360 and 390 also cut at a word).
- Wording: the shipped event reads `A combat log you upload finishes parsing`; the mock render still says `A log file you upload finishes parsing`, so the side-by-side differs there.
- Checkboxes are still the browser's bright blue native control (grey when disabled); the mock draws a paler blue tick. The channel id is still in the body font, bolder; the mock uses a monospace face.
- Fixture artifacts, not the lane: Your reports shows `did not load` (unstubbed route), and the Plan panel reads `Not on Premium` beside a live Discord panel and `On · 2 events`.
- At 2000 px the page is centred (content 376 to 1624) with the rail as the right column.
- The loading skeleton has no Discord reserve (deferred); load CLS numbers are above.

## Verdicts

- ux-designer: SHIP WITH FIXES on the first pass (the free sentence broke into two lines with a 44 px link box; the `Saving…`/`Sending…` captures were stale; a 74 px gap above `Use a different webhook`; `Mana….`; the 1-to-4 warning read as small print). Fixed in c0e41c7d and re-captured; re-check at 6f493595: **SHIP**, no must-fix items. Rulings recorded in the spec: the free sentence may wrap to two lines where the rail is narrower than the sentence (§3.2); the 520 px switched-off cap holds for a one-line error, up to 560/590 px with the longest error (§7). Report: `ux-designer-review.md`.
- wow-player: SHIP WITH FIXES (11 items). Shipped this lane: the cut error now ends at a word with no stray full stop; "log file" became "combat log" (the site's own term) in the checkbox row, the test message, the Go copy table and the spec. The rest are below. Report: `wow-player-review.md`.

## Deferred, not this lane

- The loading skeleton reserves no height for the Discord panel (load-time CLS 0.04–0.06 higher at 1440; the page already shifts 0.35–0.45 at 390 with or without the block): a page-wide account-skeleton task with a Base.astro pre-paint hint.
- An inflight pre-save `/v1/me` request can overwrite the in-place patch for a few hundred ms (shared `query.ts`).
- wow-player: the switched-off helper wording and moving `Use a different webhook` into the notice; the channel NAME instead of the id (Discord's webhook GET carries no channel name — needs a bot or a member-typed label); numbered webhook how-to steps; the private-logs note wording or a separate private-logs checkbox; gold `accent-color` on checkboxes site-wide (ux polish too); Premium row wording (`Off · needs attention`); `Took` → `Run time`, the character name in the sim title, `Boss fights`.
- ux polish: non-breaking spaces in the Premium row value at 1024; hide `Use a different webhook` while the remove question is open; 8 px less between the input and the `In Discord:` helper.
- Spec §9 item 9 (the three messages in a real channel) needs a Discord webhook from the owner; the embeds are pinned by golden JSON tests meanwhile.
