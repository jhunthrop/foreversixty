# Discord bot acceptance captures, 2026-10-27

Spec: design/specs/2026-10-27-discord-bot.md section 9 site captures rows 1-11, the amendment row 12 (2b), and the bot-only raid form. Fixture build at the lane HEAD, API stubbed. Harness: web/tests/captures/discord-bot/. CLS for rows 10 is in cls-390.json and cls-1440.json (cap 0.02). Rows at 2000 exist where the table asks (1, 4, 8).

The account captures (rows 8, 9 and 10) set the session cookie the header reads, so the header shows the signed-in selector, as the settings captures do. Captured 2026-10-27 after the persona fixes, from a fixture build served from a dist copy.

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
| `discord-bot-03-open-list-<w>.png` | 360, 390, 1440 | Picker opened from the install, list ready, raid channel preselected. The native list itself cannot be screenshotted open: its order (uncategorised channels first, then one optgroup per Discord category in Discord's order) is asserted by web/tests/e2e/guild-discord-bot.spec.ts, `the picker › ready`, on the `optgroup` label and its options |
| `discord-bot-03-open-list-crop-<w>.png` | 360, 390, 1440 | Picker open, list ready, raid channel preselected (crop of the subject region) |
| `discord-bot-03-race-<w>.png` | 360, 390, 1440 | Picker after a 422 channel_unusable save: race line, list refetched |
| `discord-bot-03-race-crop-<w>.png` | 360, 390, 1440 | Picker after a 422 channel_unusable save: race line, list refetched (crop of the subject region) |
| `discord-bot-04-just-saved-<w>.png` | 360, 390, 1440, 2000 | Installed, just saved: Saved line, focus on Send test message |
| `discord-bot-04-just-saved-crop-<w>.png` | 360, 390, 1440, 2000 | Installed, just saved: Saved line, focus on Send test message (crop of the subject region) |
| `discord-bot-04-remove-question-<w>.png` | 360, 390, 1440, 2000 | Installed with a webhook, the remove question |
| `discord-bot-04-remove-question-crop-<w>.png` | 360, 390, 1440, 2000 | Installed with a webhook, the remove question (crop of the subject region) |
| `discord-bot-04-warning-3-failures-<w>.png` | 360, 390, 1440, 2000 | Installed with the 3-failure warning: the count and the channel, then the cause from the stored error (`discord status 403 code 50001`) as the missing permissions in the channel; no switch-off clause (the bot never switches itself off) |
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
| `discord-bot-08-unlink-question-<w>.png` | 360, 390, 1440, 2000 | Account You panel, the unlink question under the kept linked line, Cancel an outline button with its focus ring |
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

Discord posts and ephemerals have no web surface. `web/tests/captures/discord-bot/payloads.json` is what the bot's own code sends: `payload-dump.go.txt` (with `payload-dump-guilds.go.txt`, which only exposes the readiness answer and the attendance chip the guilds reads build) runs as a throwaway test through `go test -overlay` from /private/tmp, nothing added to api/. It calls the real `Service` handlers (`press`, `pick`, every slash command, `share`) and the real drainer (`DrainOnce`, whose `CreateMessage` is captured on the `FakeREST`) against the real guilds service on a seeded test database (`forever_test_cap`) and the repo's real build data (`data/builds`, active build 1.60.1.70630: the BiS files and the per-class item tables, so every item is the game's own with its icon, quality, item level and source). Where a read needs a matched log or an addon export the test database cannot hold, the guilds answer is scripted in the guilds package's own types: the report (Onyxia's real drops from the build's loot table), attendance (the chip from the roster's own `attendanceChip`), readiness (the rows from the board's own `readinessAnswerOf`) and the character card (Zulmara wears the bare level-60 Alliance Beast Mastery band's picks except three slots where she wears the band's first alternative, so each `+x.x DPS` is the BiS file's measured delta for that pair). A done raid's `Came` and `Kills` are added to the real post facts the same way.

One example throughout: Olympus XXVII (server Olympus, `#raid-signups`, raid zone America/Chicago), Onyxia's Lair on Thursday 10 December 2026 at 8:00 pm CST, 40-man. Before Zulmara: 30 In (3 tanks, 8 healers, 19 DPS), 2 Late, 3 Bench, 4 Out; her In makes `33 of 40 · 3 tanks · 8 healers · 20 DPS · 2 late` everywhere. Zulmara (Beast Mastery Hunter, her main) and Duskstrider (Combat Rogue) are one account; Mendwhisper is the officer; Thalwyn is pending. The done raid: 30 of 33 came, Onyxia killed at 4:12 after 2 wipes. Each frame has its own render clock (in payloads.json, `clock`): Thursday 4:00 pm CST before the raid, 7:30 pm for locked, 11:30 pm (after the 11:00 pm end) for ended, Friday 9:00 am for done and the report, Wednesday for the moved post. Relative stamps (`in 4 hours`, `4 hours ago`) and the `Today at` lines are computed from that clock in the raid zone.

`discord-render.ts` draws the JSON as a Discord-like card: dark theme, colour bar, fields, the real thumbnail (item icons from `data/builds/<build>/icons`, crests from `web/public/icons/hd/crests`, routed in place of foreversixty.gg), buttons and the string select (closed, and opened in the `-open` frames). A shared follow-up is drawn under Discord's `Zulmara used /bis` line. It is NOT the real client: fonts, spacing and the APP tag are stand-ins. Widths: desktop 640, phone 390. Real captures from the private test server (spec D1-D12) are still owed.

| Row | File | Message | Clock (UTC) |
|---|---|---|---|
| D1 | `discord-msg-post-cancelled-{desktop,phone}.png` | Raid post: cancelled | 2026-12-10T22:00:00Z |
| D1 | `discord-msg-post-done-{desktop,phone}.png` | Raid post: done | 2026-12-11T15:00:00Z |
| D1 | `discord-msg-post-ended-{desktop,phone}.png` | Raid post: ended, no log yet | 2026-12-11T05:30:00Z |
| D1 | `discord-msg-post-locked-{desktop,phone}.png` | Raid post: locked | 2026-12-11T01:30:00Z |
| D1 | `discord-msg-post-moved-{desktop,phone}.png` | Raid post: moved to Friday | 2026-12-09T22:00:00Z |
| D1 | `discord-msg-post-open-{desktop,phone}.png` | Raid post: open | 2026-12-10T22:00:00Z |
| D1 | `discord-msg-post-webhook-{desktop,phone}.png` | Raid post through a webhook (no buttons, never edited) | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-already-{desktop,phone}.png` | Press: already In | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-bench-{desktop,phone}.png` | Press: Bench | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-cancelled-{desktop,phone}.png` | Press: cancelled | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-ended-{desktop,phone}.png` | Press: ended | 2026-12-11T05:30:00Z |
| D3 | `discord-msg-press-failure-{desktop,phone}.png` | Press: the save failed | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-in-{desktop,phone}.png` | Press: In (from the select) | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-late-{desktop,phone}.png` | Press: Late, with Add a note | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-locked-{desktop,phone}.png` | Press: locked | 2026-12-11T01:30:00Z |
| D3 | `discord-msg-press-moved-{desktop,phone}.png` | Press: In on the other character, the R3 move | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-not-member-{desktop,phone}.png` | Press: not a member | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-out-{desktop,phone}.png` | Press: Out | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-pending-{desktop,phone}.png` | Press: waiting for approval | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-rate-limit-{desktop,phone}.png` | Press: rate limit | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-select-open-{desktop,phone}.png` | Press: two characters, the select, the select opened | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-select-{desktop,phone}.png` | Press: two characters, the select | 2026-12-10T22:00:00Z |
| D3 | `discord-msg-press-unlinked-{desktop,phone}.png` | Press: not linked | 2026-12-10T22:00:00Z |
| D4 | `discord-msg-raid-unlinked-{desktop,phone}.png` | /raid, not linked | 2026-12-10T22:00:00Z |
| D4 | `discord-msg-raid-{desktop,phone}.png` | /raid | 2026-12-10T22:00:00Z |
| D4 | `discord-msg-signup-{desktop,phone}.png` | /signup status:Late character:Zulmara | 2026-12-10T22:00:00Z |
| D5 | `discord-msg-char-no-gear-{desktop,phone}.png` | /char with no gear shared | 2026-12-10T22:00:00Z |
| D5 | `discord-msg-char-unknown-{desktop,phone}.png` | /char unknown | 2026-12-10T22:00:00Z |
| D5 | `discord-msg-char-upgrades-{desktop,phone}.png` | /char with upgrades (a member's view) | 2026-12-10T22:00:00Z |
| D6 | `discord-msg-bis-after-share-{desktop,phone}.png` | /bis ephemeral after sharing | 2026-12-10T22:00:00Z |
| D6 | `discord-msg-bis-shared-{desktop,phone}.png` | /bis shared in the channel | 2026-12-10T22:00:00Z |
| D6 | `discord-msg-bis-slot-{desktop,phone}.png` | /bis slot:Ring 1 (defaults from the main) | 2026-12-10T22:00:00Z |
| D6 | `discord-msg-bis-top-{desktop,phone}.png` | /bis without a slot | 2026-12-10T22:00:00Z |
| D7 | `discord-msg-report-after-share-{desktop,phone}.png` | /report ephemeral after sharing | 2026-12-11T15:00:00Z |
| D7 | `discord-msg-report-member-{desktop,phone}.png` | /report, a member | 2026-12-11T15:00:00Z |
| D7 | `discord-msg-report-officer-{desktop,phone}.png` | /report, an officer | 2026-12-11T15:00:00Z |
| D7 | `discord-msg-report-shared-{desktop,phone}.png` | /report shared in the channel | 2026-12-11T15:00:00Z |
| D8 | `discord-msg-attendance-officer-{desktop,phone}.png` | /attendance name:Zulmara, an officer | 2026-12-10T22:00:00Z |
| D8 | `discord-msg-attendance-own-{desktop,phone}.png` | /attendance, own (two characters) | 2026-12-10T22:00:00Z |
| D8 | `discord-msg-attendance-refusal-{desktop,phone}.png` | /attendance name:Kraggor, a member | 2026-12-10T22:00:00Z |
| D9 | `discord-msg-readiness-guild-tier-{desktop,phone}.png` | /readiness, Guild tier | 2026-12-10T22:00:00Z |
| D9 | `discord-msg-readiness-lapsed-{desktop,phone}.png` | /readiness, plan lapsed | 2026-12-10T22:00:00Z |
| D9 | `discord-msg-readiness-member-{desktop,phone}.png` | /readiness, a member | 2026-12-10T22:00:00Z |
| D9 | `discord-msg-readiness-nobody-{desktop,phone}.png` | /readiness, nobody signed In or Late | 2026-12-10T22:00:00Z |
| D9 | `discord-msg-readiness-officer-{desktop,phone}.png` | /readiness, an officer | 2026-12-10T22:00:00Z |
| D9 | `discord-msg-readiness-unlinked-{desktop,phone}.png` | /readiness, not linked | 2026-12-10T22:00:00Z |
| D10 | `discord-msg-help-{desktop,phone}.png` | /help | 2026-12-10T22:00:00Z |
| D10 | `discord-msg-link-linked-{desktop,phone}.png` | /link, linked | 2026-12-10T22:00:00Z |
| D10 | `discord-msg-link-unlinked-{desktop,phone}.png` | /link, not linked | 2026-12-10T22:00:00Z |

Not rendered: D2 (a screen recording of the live edit), the `Forever Sixty is thinking…` deferral and Discord's command picker (Discord's own drawing, like D11's install and OAuth dialogs), and D12 (G1's rows through the real client).
