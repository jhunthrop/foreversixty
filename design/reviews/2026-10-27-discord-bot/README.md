# Discord bot acceptance captures, 2026-10-27

Spec: design/specs/2026-10-27-discord-bot.md section 9 site captures rows 1-11, the amendment row 12 (2b), and the bot-only raid form. Fixture build at the lane HEAD, API stubbed. Harness: web/tests/captures/discord-bot/. CLS for rows 10 is in cls-390.json and cls-1440.json (cap 0.02). Rows at 2000 exist where the table asks (1, 4, 8).

File names: `<prefix><state>-<width>.png`; widths are the viewport widths. `-crop` is a padded crop of the subject region from the same capture.

| File | Widths | State |
|---|---|---|
| `discord-bot-01-state-5-<w>.png` | 360, 390, 1440, 2000 | Settings #discord state 5: nothing connected, add link and webhook link |
| `discord-bot-01-state-5-crop-<w>.png` | 360, 390, 1440, 2000 | Settings #discord state 5: nothing connected, add link and webhook link (crop of the subject region) |
| `discord-bot-02-state-4-offer-<w>.png` | 360, 390, 1440 | State 4: webhook connected, the bot offer as an outline button |
| `discord-bot-02-state-4-offer-crop-<w>.png` | 360, 390, 1440 | State 4: webhook connected, the bot offer as an outline button (crop of the subject region) |
| `discord-bot-03-empty-<w>.png` | 360, 390, 1440 | Picker, no channels in the server |
| `discord-bot-03-empty-crop-<w>.png` | 360, 390, 1440 | Picker, no channels in the server (crop of the subject region) |
| `discord-bot-03-error-<w>.png` | 360, 390, 1440 | Picker, channel list failed to load (retry) |
| `discord-bot-03-error-crop-<w>.png` | 360, 390, 1440 | Picker, channel list failed to load (retry) (crop of the subject region) |
| `discord-bot-03-install-return-<w>.png` | 360, 390, 1440 | Picker after the install return (installed flag), focus on the select |
| `discord-bot-03-install-return-crop-<w>.png` | 360, 390, 1440 | Picker after the install return (installed flag), focus on the select (crop of the subject region) |
| `discord-bot-03-loading-<w>.png` | 360, 390, 1440 | Picker while the channel list loads (disabled select) |
| `discord-bot-03-loading-crop-<w>.png` | 360, 390, 1440 | Picker while the channel list loads (disabled select) (crop of the subject region) |
| `discord-bot-03-open-list-<w>.png` | 360, 390, 1440 | Picker open, list ready, raid channel preselected |
| `discord-bot-03-open-list-crop-<w>.png` | 360, 390, 1440 | Picker open, list ready, raid channel preselected (crop of the subject region) |
| `discord-bot-03-race-<w>.png` | 360, 390, 1440 | Picker after a 422 channel_unusable save: race line, list refetched |
| `discord-bot-03-race-crop-<w>.png` | 360, 390, 1440 | Picker after a 422 channel_unusable save: race line, list refetched (crop of the subject region) |
| `discord-bot-04-just-saved-<w>.png` | 360, 390, 1440, 2000 | Installed, just saved: Saved line, focus on Send test message |
| `discord-bot-04-just-saved-crop-<w>.png` | 360, 390, 1440, 2000 | Installed, just saved: Saved line, focus on Send test message (crop of the subject region) |
| `discord-bot-04-remove-question-<w>.png` | 360, 390, 1440, 2000 | Installed with a webhook, the remove question |
| `discord-bot-04-remove-question-crop-<w>.png` | 360, 390, 1440, 2000 | Installed with a webhook, the remove question (crop of the subject region) |
| `discord-bot-04-warning-3-failures-<w>.png` | 360, 390, 1440, 2000 | Installed with the 3-failure warning |
| `discord-bot-04-warning-3-failures-crop-<w>.png` | 360, 390, 1440, 2000 | Installed with the 3-failure warning (crop of the subject region) |
| `discord-bot-05-denied-<w>.png` | 360, 390, 1440 | Install return denied |
| `discord-bot-05-denied-crop-<w>.png` | 360, 390, 1440 | Install return denied (crop of the subject region) |
| `discord-bot-05-failed-<w>.png` | 360, 390, 1440 | Install return failed |
| `discord-bot-05-failed-crop-<w>.png` | 360, 390, 1440 | Install return failed (crop of the subject region) |
| `discord-bot-05-taken-<w>.png` | 360, 390, 1440 | Install return taken, naming the other guild |
| `discord-bot-05-taken-crop-<w>.png` | 360, 390, 1440 | Install return taken, naming the other guild (crop of the subject region) |
| `discord-bot-05-taken-expired-<w>.png` | 360, 390, 1440 | Install return taken, expired variant (no guild name) |
| `discord-bot-05-taken-expired-crop-<w>.png` | 360, 390, 1440 | Install return taken, expired variant (no guild name) (crop of the subject region) |
| `discord-bot-06-with-webhook-<w>.png` | 390, 1440 | Removed in Discord, a webhook is connected |
| `discord-bot-06-with-webhook-crop-<w>.png` | 390, 1440 | Removed in Discord, a webhook is connected (crop of the subject region) |
| `discord-bot-06-without-webhook-<w>.png` | 390, 1440 | Removed in Discord, no webhook |
| `discord-bot-06-without-webhook-crop-<w>.png` | 390, 1440 | Removed in Discord, no webhook (crop of the subject region) |
| `discord-bot-07-frozen-installed-<w>.png` | 360, 1440 | Frozen guild, installed: Remove disabled too |
| `discord-bot-07-frozen-installed-crop-<w>.png` | 360, 1440 | Frozen guild, installed: Remove disabled too (crop of the subject region) |
| `discord-bot-07-read-only-installed-<w>.png` | 360, 1440 | Officer-read-only installed: disabled test/change, read-only line |
| `discord-bot-07-read-only-installed-crop-<w>.png` | 360, 1440 | Officer-read-only installed: disabled test/change, read-only line (crop of the subject region) |
| `discord-bot-07-read-only-state-5-<w>.png` | 360, 1440 | Officer-read-only state 5: aria-disabled links |
| `discord-bot-07-read-only-state-5-crop-<w>.png` | 360, 1440 | Officer-read-only state 5: aria-disabled links (crop of the subject region) |
| `discord-bot-08-conflict-<w>.png` | 360, 390, 1440, 2000 | Account You panel after the conflict return with its help line |
| `discord-bot-08-conflict-crop-<w>.png` | 360, 390, 1440, 2000 | Account You panel after the conflict return with its help line (crop of the subject region) |
| `discord-bot-08-linked-<w>.png` | 360, 390, 1440, 2000 | Account You panel, linked |
| `discord-bot-08-linked-crop-<w>.png` | 360, 390, 1440, 2000 | Account You panel, linked (crop of the subject region) |
| `discord-bot-08-return-denied-<w>.png` | 360, 390, 1440, 2000 | Account You panel after the denied return |
| `discord-bot-08-return-denied-crop-<w>.png` | 360, 390, 1440, 2000 | Account You panel after the denied return (crop of the subject region) |
| `discord-bot-08-return-failed-<w>.png` | 360, 390, 1440, 2000 | Account You panel after the failed return |
| `discord-bot-08-return-failed-crop-<w>.png` | 360, 390, 1440, 2000 | Account You panel after the failed return (crop of the subject region) |
| `discord-bot-08-return-linked-<w>.png` | 360, 390, 1440, 2000 | Account You panel after the linked return (status, focus on Unlink) |
| `discord-bot-08-return-linked-crop-<w>.png` | 360, 390, 1440, 2000 | Account You panel after the linked return (status, focus on Unlink) (crop of the subject region) |
| `discord-bot-08-signed-out-<w>.png` | 360, 390, 1440, 2000 | Account signed out (the Discord row does not render) |
| `discord-bot-08-suggested-<w>.png` | 360, 390, 1440, 2000 | Account You panel, unlinked with the guild suggestion (the fill) |
| `discord-bot-08-suggested-crop-<w>.png` | 360, 390, 1440, 2000 | Account You panel, unlinked with the guild suggestion (the fill) (crop of the subject region) |
| `discord-bot-08-unlink-question-<w>.png` | 360, 390, 1440, 2000 | Account You panel, the unlink question |
| `discord-bot-08-unlink-question-crop-<w>.png` | 360, 390, 1440, 2000 | Account You panel, the unlink question (crop of the subject region) |
| `discord-bot-08-unlinked-<w>.png` | 360, 390, 1440, 2000 | Account You panel, unlinked (default lead) |
| `discord-bot-08-unlinked-crop-<w>.png` | 360, 390, 1440, 2000 | Account You panel, unlinked (default lead) (crop of the subject region) |
| `discord-bot-09-account-focus-discord-link-start-crop-<w>.png` | 390, 1440 | Focus-visible on the account Link Discord button |
| `discord-bot-09-account-hover-discord-link-start-crop-<w>.png` | 1440 | Hover on the account Link Discord button |
| `discord-bot-09-installed-focus-bot-change-crop-<w>.png` | 390, 1440 | Focus-visible on Change channel |
| `discord-bot-09-installed-focus-bot-open-crop-<w>.png` | 390, 1440 | Focus-visible on Open in Discord |
| `discord-bot-09-installed-focus-bot-remove-crop-<w>.png` | 390, 1440 | Focus-visible on Remove |
| `discord-bot-09-installed-focus-bot-test-crop-<w>.png` | 390, 1440 | Focus-visible on Send test message (installed) |
| `discord-bot-09-installed-hover-bot-test-crop-<w>.png` | 1440 | Hover on Send test message |
| `discord-bot-09-offer-focus-bot-add-crop-<w>.png` | 390, 1440 | Focus-visible on Add the bot (state 4 offer) |
| `discord-bot-09-offer-focus-discord-test-crop-<w>.png` | 390, 1440 | Focus-visible on the webhook Send test message |
| `discord-bot-09-offer-hover-bot-add-crop-<w>.png` | 1440 | Hover on Add the bot |
| `discord-bot-10-account-loading-<w>.png` | 390, 1440 | Account while /v1/me is held (warm reserve) |
| `discord-bot-10-account-ready-<w>.png` | 390, 1440 | Account after /v1/me answers; CLS in cls-<width>.json |
| `discord-bot-10-settings-loading-<w>.png` | 390, 1440 | Guild settings while the settings read is held (loading reserve) |
| `discord-bot-10-settings-ready-<w>.png` | 390, 1440 | Guild settings after the read answers; CLS in cls-<width>.json |
| `discord-bot-11-built-account-<w>.png` | 1440 | Built account page, suggested state |
| `discord-bot-11-built-settings-<w>.png` | 1440 | Built settings page, installed with a webhook |
| `discord-bot-11-mock-account-<w>.png` | 1440 | Mock board discord-bot-account |
| `discord-bot-11-mock-settings-<w>.png` | 1440 | Mock board discord-bot-settings (full board) |
| `discord-bot-11-side-by-side-account-<w>.png` | 1440 | Same-scale side by side, account mock left, built right |
| `discord-bot-11-side-by-side-settings-<w>.png` | 1440 | Same-scale side by side, settings mock left, built right |
| `discord-bot-12-deleted-no-webhook-notice-<w>.png` | 360, 390, 1440 | 2b channel deleted, no webhook (notice) |
| `discord-bot-12-deleted-no-webhook-notice-crop-<w>.png` | 360, 390, 1440 | 2b channel deleted, no webhook (notice) (crop of the subject region) |
| `discord-bot-12-deleted-webhook-notice-<w>.png` | 360, 390, 1440 | 2b channel deleted, webhook connected (notice) |
| `discord-bot-12-deleted-webhook-notice-crop-<w>.png` | 360, 390, 1440 | 2b channel deleted, webhook connected (notice) (crop of the subject region) |
| `discord-bot-12-deleted-webhook-picker-open-<w>.png` | 360, 390, 1440 | 2b channel deleted, webhook connected, Choose a channel picker open |
| `discord-bot-12-deleted-webhook-picker-open-crop-<w>.png` | 360, 390, 1440 | 2b channel deleted, webhook connected, Choose a channel picker open (crop of the subject region) |
| `discord-bot-12-permissions-no-webhook-notice-<w>.png` | 360, 390, 1440 | 2b missing permissions, no webhook (notice, Open in Discord) |
| `discord-bot-12-permissions-no-webhook-notice-crop-<w>.png` | 360, 390, 1440 | 2b missing permissions, no webhook (notice, Open in Discord) (crop of the subject region) |
| `discord-bot-12-permissions-no-webhook-picker-open-<w>.png` | 360, 390, 1440 | 2b missing permissions, no webhook, picker open |
| `discord-bot-12-permissions-no-webhook-picker-open-crop-<w>.png` | 360, 390, 1440 | 2b missing permissions, no webhook, picker open (crop of the subject region) |
| `discord-bot-12-permissions-webhook-notice-<w>.png` | 360, 390, 1440 | 2b missing permissions, webhook connected (notice) |
| `discord-bot-12-permissions-webhook-notice-crop-<w>.png` | 360, 390, 1440 | 2b missing permissions, webhook connected (notice) (crop of the subject region) |
| `discord-bot-raid-form-post-to-discord-bot-only-<w>.png` | 360, 390, 1440 | Raid form's Post to Discord box ticked for a bot-only guild |
| `discord-bot-raid-form-post-to-discord-bot-only-crop-<w>.png` | 360, 390, 1440 | Raid form's Post to Discord box ticked for a bot-only guild (crop of the subject region) |

## Discord messages (faithful render of the payload, not the real client)

Discord posts and ephemerals have no web surface. `payloads.json` is the JSON the Go builders produce (PostMessage, RaidPostedEmbed, pressReply with pressCopy, BrandEmbed for /help, attendanceEmbed, readinessEmbed, ItemEmbed), dumped by `web/tests/captures/discord-bot/payload-dump.go.txt` run as a throwaway test through `go test -overlay` (nothing was added to api/). `discord-render.ts` draws that JSON as a Discord-like card: dark theme, colour bar, fields, buttons and select. It is NOT the real client: fonts, spacing, the APP tag and Discord relative-time text ("in 5 hours") are stand-ins, and the thumbnail is a placeholder box. Real captures from the private test server (spec D1-D12) are still owed. Widths: desktop 640, phone 390.

| File | Message |
|---|---|
| `discord-msg-attendance-{desktop,phone}.png` | D8 /attendance own |
| `discord-msg-bis-slot-{desktop,phone}.png` | D6 /bis slot (item embed grammar) |
| `discord-msg-help-{desktop,phone}.png` | D10 /help |
| `discord-msg-post-cancelled-{desktop,phone}.png` | D1 raid post, cancelled |
| `discord-msg-post-done-{desktop,phone}.png` | D1 raid post, done |
| `discord-msg-post-ended-{desktop,phone}.png` | D1 raid post, ended |
| `discord-msg-post-locked-{desktop,phone}.png` | D1 raid post, locked |
| `discord-msg-post-moved-{desktop,phone}.png` | D1 raid post, moved |
| `discord-msg-post-open-{desktop,phone}.png` | D1 raid post, open |
| `discord-msg-post-webhook-{desktop,phone}.png` | D1 webhook post (no buttons) |
| `discord-msg-press-already-{desktop,phone}.png` | D3 press reply, already |
| `discord-msg-press-bench-{desktop,phone}.png` | D3 press reply, Bench |
| `discord-msg-press-cancelled-{desktop,phone}.png` | D3 press reply, cancelled |
| `discord-msg-press-ended-{desktop,phone}.png` | D3 press reply, ended |
| `discord-msg-press-failure-{desktop,phone}.png` | Press reply, save failed |
| `discord-msg-press-in-{desktop,phone}.png` | D3 press reply, In |
| `discord-msg-press-late-{desktop,phone}.png` | D3 press reply, Late with Add a note |
| `discord-msg-press-locked-{desktop,phone}.png` | D3 press reply, locked |
| `discord-msg-press-moved-{desktop,phone}.png` | D3 press reply, In with the Bench move |
| `discord-msg-press-not-member-{desktop,phone}.png` | D3 press reply, not a member |
| `discord-msg-press-out-{desktop,phone}.png` | Press reply, Out |
| `discord-msg-press-pending-{desktop,phone}.png` | D3 press reply, pending approval |
| `discord-msg-press-rate-limit-{desktop,phone}.png` | D3 press reply, rate limit |
| `discord-msg-press-select-{desktop,phone}.png` | D3 press reply, character select open |
| `discord-msg-press-unlinked-{desktop,phone}.png` | D3 press reply, not linked |
| `discord-msg-readiness-{desktop,phone}.png` | D9 /readiness officer |

Not rendered (needs a database-backed service, no pure builder): /raid, /signup, /char, /report, /link, /readiness member and lapsed variants, the unlinked-select states, thinking and the Discord command picker, and D11 install dialogs (Discord UI).
