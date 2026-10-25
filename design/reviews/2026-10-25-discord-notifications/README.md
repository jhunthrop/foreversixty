# Discord notifications: acceptance captures

Head sha captured: 6613bb1c (branch worktree-lane-discord-3; the tree carried the follow-up commits c5fc10f8, f46eb52c and 9b51612f when the final run was made), re-taken for the connect lane; every PNG below is refreshed (copy and layout changed). The earlier lanes were captured at c0e41c7d plus 1660ae0a and a71bbb90. Fixture build (`FOREVER_DATA=fixture`), `/v1/me` and the notification routes stubbed, Chromium at device scale 1, reduced motion on (so fades and hovers show their end state).
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

## Follow-up lane (2026-10-25)

Spec section 9 captures 14 to 23. File names are `discord-<name>-<width>.png` (full page) with `-panel-` crops where the spec asks for the panel. Run: `CAPTURE_PORT=4451 npx playwright test -c tests/captures/discord-notifications/playwright.capture.config.ts` (137 passed).

| Capture | Files |
|---|---|
| 14 paste form with the three steps | `discord-paste-steps-<390|1440>` (full and panel); `paste-empty`, `paste-invalid`, `replace` refreshed |
| 15 destination | `discord-destination-named`, `-unnamed`, `-name-long` (cut at 40 with `...`), `-name-markup` (`@everyone <b>x</b>` as text) at 390 and 1440, plus `discord-saved-<value>` full page and panel |
| 16 event rows | `discord-events-private-on`, `-private-off`, `-logs-off-private-disabled`, `-lapsed-all-disabled` at 390 and 1440 |
| 17 warning | `discord-warning-one`, `-two`, `-four-long` at 390 and 1440 |
| 18, 19 switched off | `discord-switched-off`, `discord-switched-off-long` (360, 390, 1440), `discord-switched-off-remove-confirm` (no `Use a different webhook`) |
| 20 Premium row | `discord-row-<value>-<390|1024|1440>` for off, off-set-up, off-no-events, off-needs-attention, paused-plan-lapsed, on-1-event, on-2-events; `discord-feature-off` |
| 21 loading reserve | `discord-loading-warm-<none|free|form|saved|off|lapsed>-<390|1440>` with `-ready` beside each; `discord-loading-cold-<390|1440>` with `-ready` |
| 22 gold checkboxes | `discord-checkbox-checked-disabled-1440`, `discord-checkbox-focus-1440`, `rankings-checkbox-checked-1440`, `rankings-checkbox-focus-1440`, `checkbox-sim-1440` and `checkbox-sim-page-1440` (the sim ones were taken by a throwaway spec, not the committed one) |
| 13 | `discord-side-by-side-1440.png` regenerated after `design/mocks/gen_discord_notifications.py` |

Not captured: the other seven of the ten checkbox surfaces (22), and item 23 (needs a real channel).

Panel heights printed by the spec: switched off 553 px at 360 and 390, 561 px at 1440; switched off with the 140-character error 611.5 px at 360 and 390, 600 px at 1440. Crop heights at 1440 (panel height plus 48 px padding, so panel = crop minus 48): `discord-paste-steps-panel-1440` 695 (panel 647), `discord-switched-off-panel-1440` 609 (561), `discord-switched-off-long-panel-1440` 648 (600), `discord-warning-long-panel-1440` 621 (573), `discord-saved-panel-1440` 505 (457), `discord-lapsed-panel-1440` 526 (478). Against the section 7 caps at 1440 (form 640, switched off 560): the form is 7 px over and the one-line switched off 1 px over.

### CLS and footer travel (Task 4 report, desktop project; Pixel 7 gives the same pass)

Warm, CLS / footer travel in px:

| state | 390 | 1440 |
|---|---|---|
| none | 0.0000 / 0 | 0.0001 / 0.5 |
| free | 0.0000 / 0 | 0.0001 / 0.5 |
| form | 0.0000 / 0 | 0.0000 / 0.5 |
| saved | 0.0000 / 0.5 | 0.0000 / 0 |
| off | 0.0000 / 0 | 0.0001 / 0.5 |
| lapsed | 0.0000 / 0 | 0.0001 / 0.5 |

Cold (nothing stored, hint `none`): CLS 0 for all six at 390; at 1440 0.0009 none, 0.0005 free, 0.0060 form, 0.0043 saved, 0.0052 off, 0.0045 lapsed. Cold footer travel is 27 to 656 px at 390 and 27.5 to 663.5 px at 1440. This capture run printed warm CLS 0.0000 at both widths for all six states (footer within 0.5 px) and cold CLS 0.0000 with the footer 1424 to 2080 at 390 and 987 to 1650.5 at 1440 (form state held).

### Noticed while capturing

- Paste form: three numbered steps sit above the field; the input bottom to the `We store the URL encrypted` trust line is 21 px at both 390 and 1440 (measured).
- Panel heights at 1440: paste form 647 px (cap 640, 7 over), one-line switched off 561 px (cap 560, 1 over); the 140-character error reaches 600 px at 1440 and 611.5 px at 360 and 390 (spec allows 620 and 650 there).
- Switched off carries the four-line notice with `Use a different webhook` inside it; with the remove question open that link is absent.
- Premium row values (`Off · needs attention`, `Off · Premium`, `On · 2 events`) are captured at 390, 1024 and 1440; the 1024 captures are the evidence on whether `2 events` splits.
- Loading skeleton, warm, 390 and 1440: the footer's top is the same (within 0.5 px) before and after the data arrives for all six `data-discord` states; cold, the page grows by 656 px (390) and 663 px (1440) when the Discord block lands.
- Gold checkboxes show on the Discord row and the rankings filter. On `/sim` both settings boxes (`Execute phase`, `Target dummy`) are disabled in the fixture build because the sim worker chunk 404s under `astro preview` and the run never finishes; `checkbox-sim-1440.png` was taken with the box force-enabled and checked (computed accent `rgb(229, 185, 85)`).
- The account `anonymize` box has no capture case and is not in this set.
- Fixture artifacts as before: Your reports reads `did not load` where unstubbed, and the Plan panel reads `Not on Premium` beside a live Discord panel.
- While this ran, `git status` showed uncommitted edits by someone else to `api/internal/notify/store_test.go` and seven files under `web/src` (Connect, AccountDiscordBlock, account api, layout, layout test and notifications, the changelog); the fixture build was made during that period, so the captures may or may not include them.

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

## Follow-up lane verdicts (2026-10-25)

- ux-designer: **SHIP** on the follow-up captures (no must-fix). Rulings recorded in the spec: §7's 1440 caps are 650 (form) and 565 (one-line switched off), plus a 590 px cap for the saved panel with the 1–4 warning and a 140-character error; the loading reserve is accepted (warm skeletons start within 1 px of the ready panels at both widths; cold is recorded); the gold checkbox rule is accepted; the changelog sentence approved. Report: `ux-designer-review.md` (`## Follow-up lane re-check`).
- wow-player: **SHIP WITH FIXES** (5). Shipped: the changelog names the webhook and the channel id and gives private logs their own paragraph; the short Premium row values no longer split (`On ·` / `2 events` at 1024 was real; spec §3.7 amended). Deferred below: the disabled-but-ticked private box beside its note, the error cut at a word rather than a clause. Report: `wow-player-review.md` (`## Follow-up lane re-check`).

## Deferred after the follow-up lane

- Cold `/account` load with no stored heights still grows by the Discord panel's height (~660 px) below the fold; §9.1 records cold and gates warm.
- When the log row is off, `Include private logs` stays ticked (value kept) but disabled while its note still reads "Off, private logs never post." — a note variant for that state is a later pass (ux-designer passed it; wow-player would change it).
- The long Discord error is cut at a word boundary, not at the last clause.
- A one-time account-page line for members who saved a webhook before the private-logs opt-in existed (they keep posting private logs; only the changelog says so).
- Spec §9 item 9 / capture 23: the three messages in a real Discord channel need a webhook from the owner.
- A checked `anonymize` capture and a disabled sim checkbox capture for the §9.2 record; the Premium value column width at `lg` (three-line values at 1024); 4 px above the trust line when the invalid-URL message wraps.
- `Account.svelte` is ~960 lines (pre-existing); extracting the rail panels needs the full shared-layout suite.

## Connect lane (2026-10-25)

Spec section 9 captures 24 to 28, 30 and 31. Run: `CAPTURE_PORT=4501 npx playwright test -c tests/captures/discord-notifications/playwright.capture.config.ts` against the fixture build (170 tests). 168 passed; the two `connect opening Discord` tests (24) time out under the committed spec: the held cross-origin navigation never shows the `Opening Discord…` label to the locator. Those four PNGs were taken with a throwaway spec (deleted, not committed) that cancels the link's default navigation after the click, so the link's own `opening` state is on screen with no navigation. Item 29 (Discord's own dialog) is manual; item 31 (grep for the logo and `5865f2`) passed as a test.

| Capture | Files (`discord-` prefix; full page, `-panel-` crops where noted) |
|---|---|
| 24 connect state | `connect-<360|390|1440>` (+ `connect-panel-<360|390|1440>`), `connect-hover-panel-1440`, `connect-focus-panel-1440`, `connect-opening-<390|1440>` (+ panel) |
| 25 paste view | `connect-paste-view-<390|1440>` (+ panel), `connect-off-paste-first-<390|1440>` (+ panel) |
| 26 returns | `return-<connected|denied|failed>-<390|1440>` (+ panel) and `return-<flag>-reloaded-<390|1440>` (+ panel); the URL is asserted to be `/account#discord` by the test |
| 27 saved connected | `connected-saved-<360|390|1440>` (+ panel), `connected-destination-<360|390|1440>` (crop of the destination), `connected-pasted-beside-<390|1440>` (+ panel) |
| 28 remove, replace, reconnect | `connect-remove-confirm-<connected|pasted>-<390|1440>` (+ panel), `connect-replace-view-<390|1440>` (+ panel), `connect-reconnect-<390|1440>` (+ panel) |
| 30 loading reserve | `loading-warm-connect-<390|1440>` and `-ready` beside each |
| 13 | `discord-side-by-side-1440.png` regenerated after `design/mocks/gen_discord_notifications.py` (the mock render itself is under `design/mocks/renders/`, not part of this commit) |

Panel heights (crop minus 48): connect 305 px at 360, 285 at 390, 293 at 1440 against the caps 340, 340, 320; `Opening Discord…` 285 at 390; saved connected 488 at 360 and 390, 496 at 1440; connect paste view 683 at 390, 691 at 1440; connected remove question 447 at 390.

Design fix round 1 (2026-10-25, connect lane re-check): copy changed (explanation, `failed` line, new `Discord did not open` line after 10 s), and the connect state, return outcomes, saved connected crops and paste view were re-captured. Panel heights (bounding box, equal to crop minus 48): connect state 324.5 at 360, 324.5 at 390, 313 at 1440 (caps 340, 340, 320); `failed` return 362 at 390 and 350 at 1440 (the status line grows to three lines, pushing the paste link down); paste view with the back link 683 at 360, 683 at 390, 691 at 1440 (caps 700, 700, 710); switched-off connected webhook (new `discord-connected-switched-off-<360|390|1440>` and `-panel-`) 592, 592, 600 (caps 620, 620, 605); saved connected 488.5, 488.5, 496.5 (caps 520, 520, 500), with the separator now `Channel 1198…4568 · Open in Discord`.

Reserve numbers for the `connect` state (task-3 report): warm 390 CLS 0.0000, footer 0.5 px; warm 1440 CLS 0.0001, footer 0 px (gates CLS 0.05 / 0.03, footer 8 px). Cold 390 CLS 0.0000, footer 302.5 px; cold 1440 CLS 0.0028, footer 310 px (cold is the `none` hint, as for every state, recorded not gated).

Pixel diff, `connect: false` paste state against the previous lane: `discord-connect-off-paste-first-<390|1440>` (full page and panel) against the follow-up lane's `discord-paste-steps-<390|1440>` (the unfocused paste form): 0 changed pixels in all four pairs (390 full 927030 px, panel 267930; 1440 full 2604960, panel 307190). Against `discord-paste-empty` (input focused) 10084 / 4975 (390) and 11374 / 5035 (1440) pixels differ, which is the focus ring and caret.

### Noticed while capturing (connect lane)

- Connect state at 390: gold `CONNECT A DISCORD CHANNEL` button, then one muted explanation (`Discord asks which server and channel, then lets Forever Sixty post there. We keep only that webhook, encrypted.`, two lines), a gap of about 40 px, then `Paste a webhook URL instead` as an underlined white link. No status line is present until a return flag. At 1440 the same copy wraps in two lines inside a 394 px panel; at 360 the button and copy are the same, the panel is 20 px taller (305 against 285).
- `Opening Discord…` replaces the button label and the button shrinks to the label's width (about 160 px against 227 px); the panel height is unchanged (285 at 390), the explanation and the paste link stay put.
- Return outcomes: `denied` reads `Cancelled in Discord. Nothing was connected.` in the body colour, `failed` reads `Could not connect. Check that you can manage webhooks in that channel, then try again.` in red over two lines; both draw the gold button with a focus ring (focus is on `discord-connect`) and the status line sits between the explanation and `Paste a webhook URL instead`. `connected` shows the saved connected panel, focus ring on `Send test message`, and the line `Connected. Send a test message to see it in your channel.` (two lines at 390) under the buttons. After a reload none of the three shows a status line (asserted).
- Paste view with connect on, 390: the three steps, field, trust line, three event rows, `Save webhook`, then `Connect a Discord channel instead` as a link under Save, 24 px under the button. Panel is 683 px at 390, 691 at 1440, above the 640 / 650 form caps in section 7 (the caps were set before the back link).
- Paste first with connect off is pixel for pixel the previous lane's paste form (0 changed pixels), with no connect button, no disclosure and no back link.
- Saved connected, 390 and 360: `Connected to **a Discord channel**` over `Channel 1198…4568· Open in Discord` stays on one line at both 360 and 390 (the destination crop is 90 px tall at both, 368 and 338 px wide). `Open in Discord` is underlined and white like the other links, so it reads as a link; the separator dot sits directly against `4568` with a space only after it (`4568· Open in Discord`), which looks like a glitch in the crop. The channel id is bold and the label muted.
- Reconnect: `Connect a different channel` is an underlined link about 36 px under the buttons, with the muted note `Connecting again replaces this channel and removes its webhook from Discord.` under it (two lines at 360 and 390, 488 px panel). A saved pasted webhook with connect on still shows `Use a different webhook` and not the reconnect link or note.
- Remove questions: connected reads `Remove this connection? Posts stop and we delete its webhook from Discord.`, pasted reads `Remove this webhook? Posts stop and we forget the URL. It stays in Discord until you delete it there.`; both sit over `Remove` and `Cancel` buttons, and the reconnect link and note are hidden while the question is open. The Destination line above stays as it was (`Posts through "Raid pings"` for the pasted one).
- Replace view of a pasted webhook with connect on, 390: the gold connect button, its explanation, `Paste a webhook URL instead`, then `Cancel`; no `Webhook URL` field until the paste link is chosen.
- The committed `connect opening Discord` test fails against a held navigation (the label is never seen), see the top of this section; the PNGs exist from the throwaway spec.
- Capture 29 and the in-staging parts of 26 (real redirect, URL bar) are not captured here: the URL is asserted by the test, not shown.
- The worktree moved while this ran (commits c5fc10f8 connect link to the API origin, f46eb52c, 9b51612f, later 381e6d7f, and uncommitted edits to the spec and the capture spec by someone else). The first run, at a build made before those landed, linked a relative `/v1/notifications/discord/connect`; the final full run and every PNG here come from a fresh fixture build at 9b51612f, and its output is identical in the panel sizes quoted above.
