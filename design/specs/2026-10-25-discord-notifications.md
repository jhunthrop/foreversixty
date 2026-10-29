# Discord notifications: experience spec

Author: ux-designer. 2026-10-25. Lane 5 of the premium experience. Product spec: `docs/superpowers/specs/2026-10-24-premium-experience-design.md` section 4.6 (the promise, the rulings), section 5 (the account page is the Premium console; the upsell sentence), section 2 (anti-goals), section 8. References: `docs/tenets.md` 1, 4, 9, 10, 11, 13, 14; `design/specs/2026-10-24-server-runs.md` (the finish event this rides on; the account row grammar), `design/specs/2026-10-24-character-history.md` (the free owner's sentence, ruling 3.9). Mock: `design/mocks/gen_discord_notifications.py`, render `design/mocks/renders/discord-notifications-states.png` (1440 and 390: the rail with the block, every state, the controls, the Premium row values, and the embeds as Discord draws them).

Surfaces changed: a new `web/src/components/account/AccountDiscordBlock.svelte` (built only from existing parts, section 3), `web/src/components/Account.svelte` (one mount point, 3.1), `web/src/components/account/AccountPremiumBlock.svelte` (the `notifications` row goes live, 3.8), `web/src/lib/billing/copy.ts` (the row strings, the /premium line; `shipsMonth.notifications` is deleted), a new `web/src/lib/notifications/copy.ts` (every string below) and a new pure `web/src/lib/notifications/discord-block.ts` (the state function and the URL shape check). The embeds are built by the API; section 6.2 fixes their text.

## Amendment log (amended 2026-10-25, follow-up lane)

Ten items the first lane deferred are ruled here. Each changed passage below carries the marker `(amended 2026-10-25, follow-up lane)`. The rulings, in one line each:

1. Switched off: the cure is named in the notice, and `Use a different webhook` moves into the notice (3.2, 6.1).
2. Destination: Discord's own webhook name plus a shortened channel id; no member label field (3.2, 6.1, 10).
3. How-to: three numbered steps above the field; the field-to-helper dead band shrinks (3.2, 6.1).
4. Private logs get their own checkbox, off by default, and the API gets a third event flag (3.4, 6.1, 10).
5. Premium row: `Off · needs attention`, `Off · Premium`; the row disappears when the deployment has notifications off; non-breaking spaces keep a count with its noun (3.7).
6. The 1 to 4 failures warning adds `Send a test message to check.` (3.2, 4.3, 6.1).
7. Messages: `Run time`, `Boss fights`, the character name in the sim title when the row has one (6.2).
8. `Use a different webhook` is absent while the remove question is open (3.2, 4.2, 4.3).
9. The account page's loading state draws the real two-column grid and reserves heights from the cached `/v1/me` before first paint (9.1).
10. Checkboxes are gold site-wide, one rule in `global.css` (9.2).

## Amendment log (amended 2026-10-25, connect lane)

The owner asked for one click instead of copying a webhook URL. Discord's OAuth2 `webhook.incoming` flow gives exactly that: Discord shows its own server-and-channel picker and asks the member to allow `Forever Sixty`; our API receives a new incoming webhook and stores it as it stores a pasted one. Every changed passage below carries the marker `(amended 2026-10-25, connect lane)`. The rulings, in one line each:

1. When the API says `connect: true`, the first state is the Connect state: one gold `Connect a Discord channel` and a one-sentence explanation; the paste form is a swap view behind `Paste a webhook URL instead` (3.2, 6.1). When `connect` is false the paste-first state is unchanged.
2. The round trip is a full redirect (nothing is drawn while away). The return lands on `/account?discord=connected|denied|failed#discord`; the page turns the flag into one status line, moves focus, and removes the flag from the URL (4.3).
3. A connected webhook's destination is `Connected to a Discord channel` over `Channel 1198…4568`. Discord's own webhook name (`Forever Sixty`) is never shown (3.2).
4. A connected webhook's replace link is `Connect a different channel`, one click to Discord; the API deletes the old webhook from Discord, best effort, and the link's note says so (3.2, 4.3).
5. Remove asks a different question by source: connected, `we delete its webhook from Discord`; pasted, `It stays in Discord until you delete it there` (4.3, 6.1).
6. The Premium row's seven values are unchanged (3.7).
7. No Discord logo, wordmark glyph or blurple: a plain text button in our own gold (3.2, 8).
8. Contract: `connect` and `source`, two routes, the delete rule (10).

## 0. The experience in one paragraph

A payer opens `/account` and, directly under the Plan panel, finds a `Discord` panel. With nothing saved it is a paste form: three numbered steps, one field for a webhook URL, three checkboxes for what to post, one gold `Save webhook` (amended 2026-10-25, follow-up lane). Once saved the panel says `Posts through “Raid pings”` with the shortened channel id under it, shows the three checkboxes, and offers `Send test message` and `Remove` (amended 2026-10-25, follow-up lane). When Discord rejects five posts in a row the panel says so in one sentence, names Discord's last error, says how to fix it, and offers `Test and switch on` (amended 2026-10-25, follow-up lane). When the plan ends the panel says posts are paused, keeps the webhook, and says why. A free member sees one sentence and `See plans` in the panel's place. (amended 2026-10-25, connect lane) Where the deployment has Discord connect on, the nothing-saved panel is not the paste form but one gold button, `Connect a Discord channel`: one click sends the member to Discord to pick a channel, and they come back to `Connected. Send a test message to see it in your channel.` The paste form stays one link away. The webhook token never appears on the page, not even masked.

## 1. The one question and the answer that must be visible

**"Will I hear about it where my guild already talks?"**

The answer is a state word and a destination: `Posts through “Raid pings”`, or the sentence that says why it does not (amended 2026-10-25, follow-up lane).

| | 360 / 390 px phone | 1280 / 1440 px desktop |
|---|---|---|
| Payer, saved | The `Posts through ...` line is the first line under the panel label. The whole panel (label, destination, three checkboxes, note, buttons, status line, `Use a different webhook`) is at most 480 px tall at 360 and 390, so it fits one 640 px screen once scrolled to (amended 2026-10-25, follow-up lane) (the Premium row's `On · 2 events` link lands on it, `#discord`). | The panel sits directly under Plan in the right rail (395 px wide at 1440). On a payer's account page Plan ends near y 895 at 1440, so the block begins at or below the fold; the Premium row above it answers the question in one line (`On · 2 events`), and the block is one tap or one scroll away. |
| Payer, nothing saved | The intro (two lines) and the three numbered steps fill the panel's first 220 px; the URL field starts by 260 px, inside one 640 px screen (amended 2026-10-25, follow-up lane). | Same. |
| Payer, nothing saved, connect on (amended 2026-10-25, connect lane) | The label, the intro and the gold `Connect a Discord channel` button end by 190 px, the explanation (four lines at 360 and 390) by 270 px; the whole panel is at most 340 px tall at 360 and 390 (measured 324.5, amended 2026-10-25, connect lane re-check); with a `failed` return line (three lines) it grows to 370 at 390 and 360 at 1440 (measured 362 and 350). | The button's top is at 120 px inside the panel; the panel is at most 320 px tall at 1440 (measured 313). |
| Payer, switched off or paused | The notice sentence is the first thing under the destination line, with a 2 px bar (`--wipe` switched off, `--gold` paused). | Same. |
| Free | One sentence, 44 px high, under Plan. | Same. |

The Premium row in the Plan panel is the answer for a reader who never scrolls to the block: `On · 2 events`, `Off · needs attention`, `Paused · plan lapsed`, `Off · set up` (amended 2026-10-25, follow-up lane).

## 2. The reference, and exactly what we borrow

**Discord's own Integrations, Webhooks page** (Channel settings). The job players already know: one webhook is one row with a destination (the channel); it can be copied and deleted; nothing about the token is shown after creation except through an explicit Copy. Borrowed: the destination named first; `Copy Webhook URL` as the instruction we point the member to (the helper line names the exact menu path); Delete as a plain, confirmed action. Not borrowed: the webhook's avatar and name editor (the message carries our own username and avatar, 6.2).

**GitHub's repository Webhooks settings.** The best-known "paste a URL, choose events, test it" form. Borrowed: events as checkboxes under the URL, not a list of toggles; a failing delivery is shown with its error where the webhook is listed; a failed webhook gets one line saying why and one action to retry (their `Redeliver`; ours is `Test and switch on`); the secret is write-only. Not borrowed: the payload preview, the delivery history list, the content-type picker, the SSL toggle. The member has two events and one channel; a deliveries table would be a wall of rows for nothing (tenet 11).

**Warcraft Logs' Discord integration.** The incumbent for this audience: a guild links a channel and WCL posts a report link. Not borrowed: guild-level linking and channel pickers (a Discord bot with OAuth; the product spec's non-goals rule out a guild webhook). Ours is a member's own webhook.

**The existing account page** (`web/src/components/Account.svelte`): `StatePanel` label row (gold dot, 11 px label); the email form's `label.label` + `h-11` input + `SECONDARY_BUTTON_FIXED` button; the `anonymize` checkbox row (native checkbox in a 44 px label); the device list's in-place `Revoke` confirm (a question and two buttons replace the row); `min-h-[21px]` reserved status lines; the plan panel's underlined `See plans`. Every control below is one of those.

## 3. Layout, region by region

Tokens (all existing, `web/src/styles/tokens.css`): `--color-bg #07090d`, `--color-raised #0d111a` (panel ground), `--color-line #262e40` (panel border), `--color-line-warm #3a3326` (input and button borders), `--color-gold #e5b955`, `--color-gold-hover #f5d27a`, `--color-strong #f2eee4`, `--color-text #e9e4d8`, `--color-muted #9a9484`, `--color-wipe #f0736b` (failure text and the switched-off bar; AA on raised). No new colour. Type: Barlow for text; JetBrains Mono tabular for the channel id; `label` (11 px 700, .12em, uppercase) for the panel label, `Webhook URL` and `Post when`; body 14 px; helper and status 13 px; button text 12 px 700 .06em uppercase (`SECONDARY_BUTTON_FIXED`, `PRIMARY_BUTTON_FIXED`). Radius `rounded-panel` 6 px for the panel, `rounded-control` 4 px for input and buttons. Spacing: panel padding 16 px below 768 and 24 x 20 px from 768 (`StatePanel`'s own `px-4 py-4 md:px-6 md:py-5`), 16 px between regions (`StatePanel`'s `gap-4`), 8 px between a label and its field.

Shared components used: `StatePanel` (label `Discord`, `testid="account-discord"`), `PRIMARY_BUTTON_FIXED` and `SECONDARY_BUTTON_FIXED` from `lib/planner/styles.ts`, the native checkbox row of `anonymize`, the label-plus-`h-11`-input of the email form, the underlined `See plans` link (`billingBlockCopy.seePlans`, `text-text underline underline-offset-2`). `EmptyState`, `LoadError` and `Skeleton` are **not used**: the block makes no read of its own (it renders from `/v1/me`, which the page already holds), so it has no loading or load-error state, and its empty state is the paste form, which is the thing that fills it. Failed actions use the inline status line, the pattern of the email form's `account-error`.

### 3.1 Where it mounts

In `Account.svelte`, the right rail (`lg:col-span-4`), as the **last child of the `#plan` wrapper** (`<div id="plan" class="flex flex-col gap-4">`), after the `StatePanel` that holds the plan line and `AccountPremiumBlock`. The wrapper's `gap-4` puts the Discord panel (or the free sentence) 16 px under Plan, where the rail's other panels sit 32 px apart: Plan and Discord read as one console. Rail order, top to bottom: ratings, Devices, You, Plan, **Discord**. Below `lg` the rail follows the left column in the same order, so on a phone Discord is the last panel on the page; the Premium row's link (`#discord`) gets a payer there in one tap. The panel has `id="discord"` and `scroll-margin-top: 80px` (clear of the sticky header). Mount condition: signed in and `me` loaded. It never mounts for a visitor. At 1024 the rail is 288 px wide (inner 240): the destination line, the buttons and `Off · needs attention` wrap (amended 2026-10-25, follow-up lane), nothing overflows horizontally, and the input stays full width (its placeholder truncates with an ellipsis).

### 3.2 The panel, by state

`AccountDiscordBlock` chooses one of five states from `discordBlockState(me)` (4.1) and renders it inside one `StatePanel`, except `free`, which renders a sentence and no panel.

**Paste form** (payer, nothing saved; and `Use a different webhook`):

| Order | Region | Content | Style |
|---|---|---|---|
| 1 | Panel label row | gold dot, `Discord` | `StatePanel`, as Devices |
| 2 | Intro | `Get a message in a Discord channel when a server sim finishes or a log finishes parsing.` | 13 px `--muted`, up to two lines. Omitted in the replace form. |
| 3 | Steps (amended 2026-10-25, follow-up lane) | An ordered list, three items, copy in 6.1 (`Step 1` to `Step 3`). Shown in the paste form and in the replace form. | 13 px `--muted`, `ol` with `list-decimal`, `padding-left: 20px`, markers `--muted` and tabular, 4 px between items, no margin under the last item; the panel's 16 px region gap separates it from the field group. Steps replace the one-line `HOWTO` helper, which is deleted. |
| 4 | Field group | `label.label` `Webhook URL`; the input; the 21 px reserved error line; the `TRUST` helper (amended 2026-10-25, follow-up lane) | Input: `h-11` (44 px), `border-line-warm bg-raised rounded-control text-text px-3 text-[15px]`, full width. The error line sits directly under the input (no gap between them; the line has `pt-1` inside its 21 px, so a message starts 4 px under the border) and the `TRUST` helper follows the error line with no extra gap. The dead band between the input's bottom edge and the helper's first line is 21 px (was about 40). `TRUST` is 13 px `--muted`. |
| 5 | Events | `Post when` label, then the three checkbox rows and the note (3.4) | |
| 6 | Actions | `Save webhook` (`PRIMARY_BUTTON_FIXED`, `px-5`) and, in the replace form only, `Cancel` (text link, 44 px) | The only filled gold element in the panel. |

**Connect state** (payer, nothing saved, `me.notifications.connect` is true) (amended 2026-10-25, connect lane). This replaces the paste form as the first view. The paste form is the other view of the same panel (below).

| Order | Region | Content | Style |
|---|---|---|---|
| 1 | Panel label row | gold dot, `Discord` | `StatePanel` |
| 2 | Intro | `Get a message in a Discord channel when a server sim finishes or a log finishes parsing.` | 13 px `--muted`. Omitted in the replace view. |
| 3 | Connect button (primary) | `Connect a Discord channel` | An `<a href="<API origin>/v1/notifications/discord/connect">` styled as `PRIMARY_BUTTON_FIXED` (`h-11`, 44 px, `px-5`, gold fill, `--bg` text). A link, not a `<button>`: it is a navigation, middle-click and long-press work. Same tab. Not full width: it hugs its label (about 232 px). The only filled gold element. |
| 4 | Explanation (amended 2026-10-25, connect lane re-check) | `Discord asks for a server and a text channel, not a voice channel, then lets Forever Sixty post there. You need Manage Webhooks permission in that server. We keep only that webhook, encrypted.` | 13 px `--muted`, 8 px under the button, up to four lines. It says what Discord will show and what we keep, and nothing else. (amended 2026-10-25, first-use fixes) The reworded text is 172 characters against 168 and must still wrap to four lines at 360 and 390 and stay inside the caps (340/340/320). The owner picked a voice channel because Discord's picker lists voice channels; `text channel, not a voice channel` is the only guard we have. If the capture shows five lines, the budget is 360 at 360 and 390 px and the cap at 1440 stays 320; the words do not change. |
| 5 | Status line (amended 2026-10-25, connect lane re-check) | reserved 21 px (one line), empty until a return outcome (4.3); a two- or three-line outcome (`failed`, the stuck line) grows the line and pushes `Paste a webhook URL instead` down | 13 px, `--text` for `denied`, `--wipe` for `failed`; `role="status"` for denied, `role="alert"` for failed. |
| 6 | Paste disclosure | `Paste a webhook URL instead` | Link-button, 13 px `--text`, underline, `min-h-11`, `aria-controls` the paste view. Activating it swaps the panel's body for the paste view with focus on the URL input. |

No steps, no field and no checkboxes here: a connected webhook starts with the defaults (sim on, log on, private off) and the member edits them in the saved state, where changes save at once. Region gaps are the panel's 16 px; the button, explanation, status line and link sit in one group with 8 px between them.

**Paste view when connect is on** (amended 2026-10-25, connect lane): the paste form above, unchanged in every row, plus one link-button under `Save webhook`, `Connect a Discord channel instead` (13 px `--text`, underline, 44 px), which swaps back with focus on the connect button. The view swap is not an accordion: the other view is removed, so the panel never holds both. The `Webhook URL` field keeps what was typed while the member stays on the paste view; switching away clears it.

**Replace view when connect is on** (a pasted webhook, `Use a different webhook`): the connect view without the intro and with `Cancel` (text link, 44 px) as its last row; the saved webhook keeps posting until a new one is saved. A connected webhook has no replace view: its link goes straight to Discord (Saved, below).

**Connect off** (`connect` is false or absent): the paste-first form above, exactly as shipped. The connect button, the disclosure and the back link do not exist.

**Saved** (payer, working):

| Order | Region | Content | Style |
|---|---|---|---|
| 1 | Label row | `Discord` | |
| 2 | Destination (primary) (amended 2026-10-25, follow-up lane) (amended 2026-10-25, connect lane) | For a pasted webhook: line one `Posts through “Raid pings”` (the name is Discord's own name for the webhook, 6.1). Line two: `Channel 1198…4568` (the first four and last four digits of the channel id, joined by an ellipsis). Without a stored name, line one reads `Posts to a Discord channel`. For a connected webhook (`source` is `connected`): line one `Connected to a Discord channel`, line two the same shortened channel id followed by ` · Open in Discord` (decision 22, amended: in this lane, inline; shown for any webhook whose `guild_id` is known). Ruling: a connected webhook's Discord-side name is the application's name (`Forever Sixty`), so it says nothing and is never shown (the API returns `name: null` for a connected webhook; the web would ignore it regardless). The API cannot return a channel or server name (that needs a bot), so the member recognises the channel by the one they chose in Discord's dialog and by the test message. `Connected to` is `--muted`, `a Discord channel` is `--strong` 600; line two as for pasted. | Line one 14 px: `Posts through` `--muted`, the quoted name `--strong` 600; the name is escaped text, at most 40 characters shown then `…`, wraps at word boundaries, `overflow-wrap: anywhere`. Line two 13 px `--muted`: `Channel` in Barlow, the shortened id JetBrains Mono 13 px tabular `--text`; the id never breaks. 2 px between the lines. |
| 3 | Failure warning (only with 1 to 4 failures) (amended 2026-10-25, follow-up lane) | Paragraph one: `The last 3 posts failed. At 5 in a row it switches off. Send a test message to check.` Paragraph two: `Last error from Discord: Missing Permissions.` | Paragraph one 13 px `--text`; paragraph two 13 px `--muted`; 2 px apart; no bar, no new colour. |
| 4 | Events | `Post when`, the three checkbox rows (3.4), the note (amended 2026-10-25, follow-up lane) | |
| 5 | Actions | `Send test message` and `Remove`, both `SECONDARY_BUTTON_FIXED`, `px-4`, 12 px gap, wrapping | No filled button: the panel is quiet when it works. |
| 6 | Status line | reserved 21 px, 13 px | `--text` for success, `--wipe` for failure; `role="status"`, `role="alert"` on failure |
| 7 | Replace link | Pasted: `Use a different webhook`. Connected: `Connect a different channel` with the note `Connecting again replaces this channel and removes its webhook from Discord.` under it (amended 2026-10-25, connect lane) | 13 px `--text`, underline, 44 px tall, directly under the status line. Not rendered while the remove question is open (amended 2026-10-25, follow-up lane). The connected link is an `<a href="<API origin>/v1/notifications/discord/connect">` (the API's origin, `api.foreversixty.gg` in production — the page host has no `/v1`; amended 2026-10-25 after the whole-branch review) (one click to Discord, no intermediate view); its note is 13 px `--muted`, no gap above, up to two lines, always shown. |

**Switched off** (`disabled_reason` is `failed five times`, plan active) (amended 2026-10-25, follow-up lane): label row; destination; a notice; actions `Test and switch on` (`PRIMARY_BUTTON_FIXED`, the panel's only filled gold element) and `Remove`; status line; then the events (still editable). The notice is one block with a 2 px `--wipe` left bar and four lines, top to bottom: (1) 14 px `--strong` `Switched off after 5 failed posts in a row.`; (2) 13 px `--muted` `Last error from Discord: Unknown Webhook.`, omitted when `last_error` is null, cut at a word at 140 characters and ending in one `…` with no full stop after it; (3) 13 px `--muted` `Fixed it in Discord? Test and switch on. If the webhook is gone, make a new one.`; (4) the link-button `Use a different webhook` (13 px `--text`, underline, 44 px tall, 0 gap above it). The link moved out of its place under the status line: the cure for a deleted webhook is now read in the same block as the cause, and `Test and switch on` sits directly below the notice. The old helper `Sends a test message. If Discord accepts it, posts resume.` is deleted. For a CONNECTED webhook (amended 2026-10-25, connect lane re-check) the fourth line is `Connect a different channel` (straight to Discord) with its note `Connecting again replaces this channel and removes its webhook from Discord.` (13 px `--muted`, two lines at 360 and 390) under it; the panel is at most 620 px at 360 and 390 and 605 px at 1440 (measured 592 and 600).

**Plan lapsed** (`entitlements.notifications` is not true and a webhook is saved): label row; destination; a notice (2 px `--gold` left bar, `Posts are paused because your plan ended.`, under it `We kept your webhook. Posts start again when you are back on Premium.`); the `See plans` link; the three checkboxes disabled (read-only, showing what was chosen, no note) (amended 2026-10-25, follow-up lane); `Remove`. No `Send test message`, no `Use a different webhook`: the API refuses both.

A lapsed connected webhook shows `Connected to a Discord channel` as its destination line (amended 2026-10-25, connect lane).

**Free** (`entitlements.notifications` is not true and nothing is saved): the sentence `Discord messages tell you when a server sim or a log you uploaded is done. See plans` (13 px `--muted`, min height 44 px on the paragraph — `min-h-11 flex items-center`; the sentence wraps to two lines where the rail is narrower than the sentence, which includes 1440 at a 395 px rail, with `See plans` ending the second line, never alone on it; `See plans` underlined, `href="/premium"`, an inline link whose focus ring hugs the text). No panel, no heading, no form, no lock icon. This is the same decision as the character page's sentence (history spec 3.9).

### 3.3 The input

`<input type="text" name="discord-webhook" id="discord-webhook" inputmode="url" autocomplete="off" autocapitalize="off" autocorrect="off" spellcheck="false" enterkeyhint="done" aria-describedby="discord-webhook-help discord-webhook-error">` inside a `<form>`, so Enter submits. It is a plain text field, not `type="password"` (decision 10): a password type invites the browser's password manager to save the URL, and the member needs to see that a paste landed. The typed value lives only in component state until the call returns. It is never written to `localStorage`, `sessionStorage`, the URL, a log, a test id, an error string or an analytics event, and it is cleared the moment the save succeeds. A failed save keeps the typed text so the member can fix it.

### 3.4 The checkbox rows (amended 2026-10-25, follow-up lane)

The native checkbox in a `label.flex.min-h-11.items-center.gap-3.text-[14px]`, exactly the `anonymize` row, gold through the site-wide rule in 9.2. Order is fixed: `A server sim finishes`, then `A combat log you upload finishes parsing`, then `Include private logs`. The third row is a modifier of the second, not an event: it is indented 28 px (the box's 16 px plus the 12 px gap) so it reads as a child, and it is disabled (greyed, value kept) while the second row is off. Defaults on a new webhook: sim on, log on, private off. Under the third row, the note `Off, private logs never post. On, everyone in the channel sees the log's name and “Only you”.` (13 px `--muted`, always shown, so toggling the box never changes the panel's height). A private log posts only when the third box is on; the `Visible to` field of a posted private log reads `Only you` (6.2). The third box is not counted by the Premium row (`On · 2 events` counts the first two rows only). In the saved state a change **saves at once** (`PUT` with `events` only): both boxes are disabled with `aria-busy` while the call is in flight, the status line says `Saved.` for 4 s and then clears. A failed call puts the box back and the status line says `Could not save that. Try again.`

### 3.5 What the panel never shows

The webhook id path segment and the token, in any form: not masked, not truncated, not in `title` or `aria-label`, not in the DOM at all (the API never returns them, so they cannot be). The channel id is not a credential; it is shown shortened (`1198…4568`) only because 19 digits are noise, and the webhook's name is shown as escaped text (amended 2026-10-25, follow-up lane). No `Copy` button (there is nothing to copy). No delivery history, counters, "last posted" time or per-event counts: they are pipeline numbers (tenet 7), and the member learns it works from the channel itself. `failures` is shown only as the 1-to-4 warning sentence.

### 3.6 A second channel later (a design rule; nothing is drawn)

The panel is one **channel section**: destination line, `Post when` checkboxes, actions, status line. Events are saved per channel. When a second channel ships (the native apps' phone notifications, owner's direction 2026-10-24), the panel label changes from `Discord` to `Notifications`; each channel becomes a section with a 15 px 600 `--strong` sub-heading (`Discord`, `Phone`); sections are separated by a 1 px `--line-soft` rule with 20 px padding either side; the Premium row label changes from `Discord notifications` to `Notifications` and counts channels on. The `#discord` anchor stays on the Discord section. `AccountDiscordBlock` takes its channel name, field copy and event list from `lib/notifications/copy.ts` keyed by channel, so the second channel is a second entry and a second component for its own destination (a phone has no URL field), not a redesign. This lane builds only the Discord entry and writes no `Notifications` heading.

### 3.7 The Premium row (`AccountPremiumBlock.svelte`)

The `notifications` row stops being `Ships December` and reads the block's state. It links to `#discord` for every payer state and is muted plain text for `free`. When `/v1/me` has no `notifications` key (the deployment has notifications off), the row is not rendered at all (amended 2026-10-25, follow-up lane). `AccountPremiumBlock` gains one prop, `notifications` (the state key and the count of events on), and the row is live when `entitlements.notifications === true` or a webhook is saved.

| State | Value | Link |
|---|---|---|
| Free, nothing saved (amended 2026-10-25, follow-up lane) | `Off · Premium` | no, `--muted` |
| Payer, nothing saved | `Off · set up` | yes |
| Saved, two events on | `On · 2 events` | yes |
| Saved, one event on | `On · 1 event` | yes |
| Saved, no event on | `Off · no events chosen` | yes |
| Switched off (amended 2026-10-25, follow-up lane) | `Off · needs attention` | yes |
| Plan lapsed, webhook kept | `Paused · plan lapsed` | yes |

The value column is about 190 px at 390: `Off · needs attention` (21 characters, 14 px) fits one line at 390 and wraps to two at 360 and at 1024, and the underline wraps with it, as the History row's does. Non-breaking spaces (amended 2026-10-25, follow-up lane): the value is written with U+00A0 between the state word and the dot (`On ·`) and between a count and its noun (`2 events`, `1 event`), and with ordinary spaces everywhere else. The short values — `On · 2 events`, `On · 1 event`, `Off · Premium`, `Off · set up` — are written with U+00A0 throughout and never wrap (amended 2026-10-25 after the follow-up captures: `On ·` / `2 events` split at 1024); the longer values wrap only between ordinary words (`Off · needs` / `attention`, `Off · no events` / `chosen`, `Paused · plan` / `lapsed`), never after the dot and never between a count and its noun. The mid-dot is U+00B7.

## 4. States (every interactive element)

### 4.1 Which state shows

`discordBlockState(me)` returns exactly one key, in this precedence:

| Key | When | Shows |
|---|---|---|
| `lapsed` | `me.notifications.discord` is not null and `me.entitlements.notifications` is not true | Plan lapsed (3.2). Wins over a saved `disabled_reason` of any kind: nothing can be done until the plan is back. |
| `switched-off` | saved, entitled, `disabled_reason` is not null | Switched off. The sentence for `failed five times` is fixed; any other non-null reason reads `Switched off.` and shows `last_error` if present. |
| `saved` | saved, entitled, `disabled_reason` is null | Saved |
| `form` | entitled and nothing saved (`notifications` or `notifications.discord` is null or missing) | Paste form |
| `free` | not entitled and nothing saved | The sentence |

Which body the `form` state draws (amended 2026-10-25, connect lane): the Connect state when `me.notifications.connect` is true, else the paste form; a per-panel flag (`pasteView`) holds the member's swap to the paste view. `discordBlockState` still returns `form`; the view is a second pure function of `me` and the flag.

A `/v1/me` with no `notifications` key means the deployment has notifications off (§10: the key is omitted when the API has no encryption key): the block renders nothing and the Premium row is not rendered, whatever the entitlement (amended 2026-10-25, follow-up lane). A saved webhook with both events off is `saved`; the Premium row says `Off · no events chosen` and the block shows the unchecked boxes.

### 4.2 Interactive elements

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Webhook URL input | `border-line-warm`, placeholder `https://discord.com/api/webhooks/…` in `--muted` | border `--line-warm-strong` | 2 px `--gold` outline, 2 px offset | n/a | while the save is in flight | n/a | placeholder shown | border `--wipe`, message in the 21 px line under it (4.3), `aria-invalid="true"`; typed text stays |
| `Save webhook` (primary) | gold fill, `--bg` text | fill `--gold-hover` | 2 px gold outline, 2 px offset | fill darkens to `#cfa64a` | never disabled for an empty field (a disabled button hides why; an empty press shows the error) | label `Saving…`, disabled, `aria-busy` | n/a | status line under the buttons (4.3); the button returns to rest |
| Checkbox row (each; three rows (amended 2026-10-25, follow-up lane)) | native box, gold (9.2), 14 px label | none | 2 px gold outline on the box | native | disabled and `--muted` while a save is in flight, and always in `lapsed` | `aria-busy` on the group | n/a | box restored, status `Could not save that. Try again.` |
| `Send test message` | outline `--line-warm`, `--text` | border `--gold`, text `--gold-hover` | 2 px gold outline, 2 px offset | ground `--card-top` | while any call is in flight | label `Sending…`, disabled, `aria-busy` | n/a | status line, `--wipe` (4.3) |
| `Test and switch on` (primary, switched-off only) | gold fill | `--gold-hover` | gold outline | darker | in flight | label `Sending…` | n/a | status line; the notice stays |
| `Remove` | outline | border `--gold` | gold outline | ground | in flight | n/a | n/a | status `Could not remove it. Try again.` |
| Remove confirm | the question `Remove this webhook? Posts stop and we forget the URL.` replaces the actions row, with `Remove` and `Cancel` (both `SECONDARY_BUTTON_FIXED`) | as `Remove` | as above; focus moves to `Cancel` | | `Remove` while in flight | `Removing…` | n/a | status `Could not remove it. Try again.`; the question stays |
| `Use a different webhook` (in the saved panel, and inside the switched-off notice (amended 2026-10-25, follow-up lane)) | underlined `--text` link-button, 44 px; not rendered while the remove question is open | `--gold-hover` | gold outline, 2 px offset, radius 2 | `--gold-hover` | in flight | n/a | n/a | n/a |
| `Cancel` (replace form) | as the link | as above | as above | as above | while saving | n/a | n/a | n/a |
| `See plans` | `--text` underlined | `--gold-hover` | gold outline, 2 px offset | `--gold-hover` | n/a | n/a | n/a | n/a |

**Connect elements** (amended 2026-10-25, connect lane):

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| `Connect a Discord channel` (link styled as primary) | gold fill, `--bg` text | fill `--gold-hover` | 2 px gold outline, 2 px offset | fill `#cfa64a` | n/a (an anchor); `aria-disabled` after the first click | label `Opening Discord…`, `aria-busy`, pointer events off, from the click until the page unloads; restored on `pageshow` with `persisted` (the back button) | n/a | none here: failure arrives as `?discord=failed` (4.3) |
| `Paste a webhook URL instead`, `Connect a Discord channel instead`, `Connect a different channel` | underlined `--text` link-button, 44 px | `--gold-hover` | gold outline, 2 px offset, radius 2 | `--gold-hover` | in flight | the connected link shows `Opening Discord…` as the button does | n/a | n/a |

Tab order (amended 2026-10-25, connect lane): connect state, `Connect a Discord channel`, `Paste a webhook URL instead`, then (replace view) `Cancel`; paste view, as before and then `Connect a Discord channel instead`; saved connected, the checkboxes, `Send test message`, `Remove`, `Connect a different channel`. Enter activates the anchor and Space the link-buttons as the platform does; Escape in the replace view acts as `Cancel`.

Touch and keyboard parity: every control is 44 px high on phone (buttons `h-11`, link-buttons `min-h-11`, checkbox rows `min-h-11`, input `h-11`). Tab order in the form: input, `Save webhook`, `Cancel` (replace form only); checkboxes come before the actions in DOM order, so the order is input, the three checkboxes, `Save webhook` (amended 2026-10-25, follow-up lane). Saved: the three checkboxes (the third skipped while the second is off, since a disabled box takes no focus), `Send test message`, `Remove`, `Use a different webhook`. Switched off: `Use a different webhook` (in the notice), `Test and switch on`, `Remove`, then the checkboxes (amended 2026-10-25, follow-up lane). Enter in the input submits; Space toggles a focused checkbox; Escape in the replace form acts as `Cancel`; Escape in the remove question acts as `Cancel`. No hover-only behaviour exists, so touch needs nothing extra.

### 4.3 The flows, each state's words

**Save.** Client check first (`discordBlockState` module, `webhookShapeError(text)`): the trimmed text must parse as `https://`, with host exactly `discord.com` or `discordapp.com`, and a path of `/api/webhooks/<number>/<token>` (an optional `v6` to `v10` version segment is accepted); anything else shows the one message below and makes no call. The API repeats the check and is the authority; a 400 or 422 from it shows the same message. On 200 the typed value is cleared, the state becomes `saved`, focus moves to `Send test message`, and the status line says `Saved. Send a test message to see it in your channel.` (`role="status"`). The save does not post anything to Discord.

| Failure | Text (verbatim, in the 21 px line under the input) |
|---|---|
| Empty | `Paste a Discord webhook URL.` |
| Not a Discord webhook URL (wrong host, wrong path, not a URL, or the API said 400 or 422) | `That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/.` |
| 402 | `Discord messages are part of Premium.` followed by the `See plans` link |
| Network or 5xx | `Could not save that. Try again.` |

None of these messages repeats, quotes or hints at what was typed.

**Test.** `Send test message` posts `POST /v1/notifications/discord/test`. It reads `Sending…` and is disabled until the call answers (the API times out at 5 s, so a call ends within about 6 s). 204: status `Sent. Check the channel.` in `--text`. A Discord error: status `Discord said: Unknown Webhook.` in `--wipe`, using the error text the API returns (shortened to 140 characters by the API; the web also truncates at 140 and never renders HTML). 429 from us: `Wait a minute before sending another test.` Network or 5xx: `Could not send the test. Try again.` Status lines that last (amended 2026-10-25, first-use fixes): the test's success lines `Sent. Check the channel.` and `Sent. Posts are on again.` and the connect line `Connected. Send a test message to see it in your channel.` have no timer. They stay until the next action: a press on any button or link in the panel, a checkbox change, or a leave. The next action replaces the line at the moment of the press (`Sending…` is on the button, and the line is empty while the call is in flight, then holds the new result). A reload shows none of them (4, step 4). Every error line already stays the same way. Only `Saved.` (the event toggles) and `Removed.` keep their 4 s, because they confirm something the member can see change on the page. The reason is the first real use: the member pressed Send test message, switched to Discord, came back after 4 s and saw nothing. The line is the reserved 21 px line, so staying costs no layout. A manual test does not count toward the five failures. A test that succeeds on a switched-off webhook switches it on, and the block re-renders as `saved` with the status `Sent. Posts are on again.`

**Switched off after five failures.** The server sets `disabled_reason: "failed five times"` and keeps `last_error`. On the next `/v1/me` read the panel is the Switched off state and the Premium row says `Off · needs attention` (amended 2026-10-25, follow-up lane). There is no live push to an open page; a member with the account page open learns on its next load. The way back is `Test and switch on` when the problem was fixed in Discord, or `Use a different webhook` when the webhook is gone; both are in reach of the notice that says why (amended 2026-10-25, follow-up lane).

**Failure warning.** With 1 to 4 failures in a row, the saved panel adds two paragraphs (amended 2026-10-25, follow-up lane): `The last N posts failed. At 5 in a row it switches off. Send a test message to check.` (`The last post failed.` for one), then `Last error from Discord: …`. They disappear when a post succeeds.

**Connect, the round trip** (amended 2026-10-25, connect lane).

1. The click is a full-page navigation to `GET /v1/notifications/discord/connect`, which answers 302 to `https://discord.com/oauth2/authorize` (10). The page draws nothing while the member is away: no overlay, no spinner, no polling. Discord's own dialog asks which server and channel and asks them to authorise `Forever Sixty`; the member needs permission to manage webhooks in the channel and Discord's dialog says so. The only thing we draw before leaving is the button's `Opening Discord…`.
2. Discord sends the member to `GET /v1/notifications/discord/callback`, which answers 302 to `/account?discord=<outcome>#discord`. The outcomes:

| Outcome | Meaning | Panel after return | Status line (verbatim) | Focus |
|---|---|---|---|---|
| `connected` | the webhook was created and stored | `saved`, connected; the page treats its cached `/v1/me` as stale and renders the block only from the fresh read | `Connected. Send a test message to see it in your channel.` (`--text`, `role="status"`, stays until the next action) | `Send test message` |
| `denied` | the member cancelled in Discord (no code) | unchanged: Connect state, or the saved panel when reconnecting; nothing is stored | `Cancelled in Discord. Nothing was connected.` (`--text`, `role="status"`) | the connect control that was used (`Connect a Discord channel` or `Connect a different channel`) |
| `failed` | anything else (bad state, exchange failed, no webhook in the answer, not entitled, a store error) | unchanged | `Could not connect. You need to be able to manage webhooks in that channel. Ask a server admin, or paste a webhook URL instead.` (`--wipe`, `role="alert"`) | the same control |

3. The status line of a connect outcome is the panel's own (the reserved 21 px line of the connect state, or the saved panel's status line), so it never moves a button. A `denied` or `failed` line stays until the member acts or leaves.
4. URL clean-up: on mount the page reads `discord`, keeps the value in component state, and calls `history.replaceState` to the same path and query without `discord`, keeping the hash (`/account#discord`), the pattern of the existing `?refreshed=1` toast. Any value other than the three is ignored and removed. Reloading the page therefore shows no status line. The hash scrolls the panel to the top (`scroll-margin-top: 80px`).
5. Back button: from Discord the member may press back; the page restores from the bfcache or reloads; the button's `Opening Discord…` is reset on `pageshow` with `persisted`. Stuck click (amended 2026-10-25, connect lane re-check): if the page is still here 10 s after the click (no navigation, so no `pageshow`), the button label is restored and the status line (`role="status"`) shows `Discord did not open. Try again, or paste a webhook URL instead.`; the status line grows like the `failed` outcome (two to three lines at 390).
6. Reconnect on a saved webhook: `denied` and `failed` leave the old webhook untouched and posting. `connected` replaces it: events, `failures`, `disabled_reason` and `last_error` follow 10; the old webhook is deleted from Discord when it was a connected one.
7. A connect while switched off returns `connected` as a fresh saved webhook (failures 0, switched on); the status line is the connected line, not `Sent. Posts are on again.`

**Remove.** `Remove` replaces the actions row with the question, and `Use a different webhook` is removed from the panel until the question is answered (amended 2026-10-25, follow-up lane). The question depends on `source` (amended 2026-10-25, connect lane): connected, `Remove this connection? Posts stop and we delete its webhook from Discord.`; pasted, `Remove this webhook? Posts stop and we forget the URL. It stays in Discord until you delete it there.` The delete from Discord is best effort and silent: if Discord refuses it the member still sees `Removed.`, because the webhook is already unusable by us. Confirm calls `DELETE /v1/notifications/discord`; success returns to the nothing-saved state (the Connect state when connect is on, else the paste form) with focus on the connect button or the input and the status `Removed.` shown once under the panel's label row for 4 s (`role="status"`); in `lapsed`, success returns to `free`, which is the sentence, and the focus lands on the Plan panel's Discord row link.

**Replace.** `Use a different webhook` swaps the saved view for the paste form without the intro, with `Cancel`. The saved webhook keeps posting until a new one is saved. `Save webhook` here sends the new URL and the current events; success returns to `saved` with the new channel id.

**Free.** No state: the sentence only.

## 5. Hierarchy

- **One primary element per state.** Connect state (amended 2026-10-25, connect lane): `Connect a Discord channel`, the only filled gold element; `Paste a webhook URL instead` is a text link. Paste form: `Save webhook`, the only filled gold element. Saved: the destination line (`Posts through ...`) (amended 2026-10-25, follow-up lane); no filled button, because nothing is wrong. Switched off: the notice, and `Test and switch on` as the only filled gold element. Plan lapsed: the notice; `See plans` is a text link, not a button (the upsell grammar). Free: the sentence.
- **One level of emphasis.** The destination (14 px, the id in mono `--strong`) is the answer; everything else is 13 px `--muted` or a control. `--wipe` appears only on the switched-off bar, failed status text and the invalid-field border. `--gold` appears on the panel dot and label (as every `StatePanel`), the primary button, the lapsed bar and focus outlines.
- **Secondary is visibly secondary.** `Use a different webhook` and `Cancel` are text links under the actions; `Remove` is an outline button, never red, never first (the account page already keeps its one red for the guild-leave confirm).
- **Whitespace.** 16 px between regions, 8 px label to field, 44 px rows; the 21 px status line is reserved, so a result never moves the buttons. No dividers inside the panel.
- **Not a pitch.** The free state is one sentence with no panel, heading, lock or colour; a free member never sees a form they cannot use.

## 6. Copy, verbatim

New strings live in `web/src/lib/notifications/copy.ts`; the Premium row and /premium strings in `web/src/lib/billing/copy.ts`. Functions take the data and return the string; counts and the channel id come from `/v1/me`.

### 6.1 On the page

| Where | String |
|---|---|
| Connect button (amended 2026-10-25, connect lane) | `Connect a Discord channel`; pending `Opening Discord…` |
| Connect explanation (amended 2026-10-25, first-use fixes) | `Discord asks for a server and a text channel, not a voice channel, then lets Forever Sixty post there. You need Manage Webhooks permission in that server. We keep only that webhook, encrypted.` |
| Paste disclosure (amended 2026-10-25, connect lane) | `Paste a webhook URL instead` |
| Back to connect (amended 2026-10-25, connect lane) | `Connect a Discord channel instead` |
| Destination, connected (amended 2026-10-28, polish task 4, player ruling) | line one `Posts to a channel in your Discord server`; no id and no line two. Never the Discord-side webhook name |
| Reconnect link and note (amended 2026-10-25, connect lane) | `Connect a different channel` / `Connecting again replaces this channel and removes its webhook from Discord.` |
| Return, connected (amended 2026-10-25, connect lane) | `Connected. Send a test message to see it in your channel.` Stays until the next action (amended 2026-10-25, first-use fixes; 4.3). |
| Return, denied (amended 2026-10-25, connect lane) | `Cancelled in Discord. Nothing was connected.` |
| Return, failed (amended 2026-10-25, connect lane re-check) | `Could not connect. You need to be able to manage webhooks in that channel. Ask a server admin, or paste a webhook URL instead.` |
| Connect did not open (amended 2026-10-25, connect lane re-check) | `Discord did not open. Try again, or paste a webhook URL instead.` |
| Remove question, connected (amended 2026-10-25, connect lane) | `Remove this connection? Posts stop and we delete its webhook from Discord.` (buttons `Remove`, `Cancel`) |
| Remove question, pasted (amended 2026-10-25, connect lane) | `Remove this webhook? Posts stop and we forget the URL. It stays in Discord until you delete it there.` (buttons `Remove`, `Cancel`) |
| Panel label | `Discord` |
| Intro (form) | `Get a message in a Discord channel when a server sim finishes or a log finishes parsing.` |
| Field label | `Webhook URL` |
| Placeholder | `https://discord.com/api/webhooks/…` |
| Step 1 (amended 2026-10-25, follow-up lane) | `In Discord, open the channel's settings, then Integrations, then Webhooks.` |
| Step 2 (amended 2026-10-25, follow-up lane) | `Choose New Webhook, name it something you will recognize, then Copy Webhook URL.` |
| Step 3 (amended 2026-10-25, follow-up lane) | `Paste it below.` |
| Helper (trust) | `We store the URL encrypted, never show it again, and delete it with your account.` |
| Events label | `Post when` |
| Event 1 | `A server sim finishes` |
| Event 2 | `A combat log you upload finishes parsing` |
| Event 3 (amended 2026-10-25, follow-up lane) | `Include private logs` |
| Events note (amended 2026-10-28, polish task 4, player ruling) | `Off: private logs are never posted. On: the channel sees the log's name and “Only you”.` |
| Save | `Save webhook`; pending `Saving…` |
| Cancel | `Cancel` |
| Destination, with a name (amended 2026-10-25, follow-up lane) | line one `Posts through` `“{name}”`; line two `Channel` `{first 4 digits}…{last 4 digits}` (for example `Posts through “Raid pings”` / `Channel 1198…4568`) |
| Destination, no name (amended 2026-10-28, polish task 4, player ruling) | line one `Posts to a channel in your Discord server`; no line two |
| Test | `Send test message`; pending `Sending…` |
| Remove | `Unlink Discord`; pending `Unlinking…` (amended 2026-10-28, polish task 4, player ruling) |
| Replace (pasted) (amended 2026-10-25, connect lane) | `Use a different webhook` |
| Saved status | `Saved. Send a test message to see it in your channel.` after a first save (stays until the next action, like the connected line: it is a prompt to act; controller ruling 2026-10-25, first-use fixes); `Saved.` after an event change (4 s) |
| Test sent | `Sent. Check the channel.`; after switching on `Sent. Posts are on again.` Both stay until the next action (amended 2026-10-25, first-use fixes; 4.3). |
| Test failed | `Discord said: {error}.` (the error text, one trailing full stop) |
| Test rate limited | `Wait a minute before sending another test.` |
| Test, our error | `Could not send the test. Try again.` |
| Removed | `Removed.` |
| Remove failed | `Could not remove it. Try again.` |
| Save failed (event change or webhook) | `Could not save that. Try again.` |
| Empty field | `Paste a Discord webhook URL.` |
| Invalid URL | `That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/.` |
| Not entitled on save | `Discord messages are part of Premium.` + `See plans` |
| Warning, one (amended 2026-10-25, follow-up lane) | `The last post failed. At 5 in a row it switches off. Send a test message to check.` then `Last error from Discord: {last_error}.` |
| Warning, several (amended 2026-10-25, follow-up lane) | `The last {n} posts failed. At 5 in a row it switches off. Send a test message to check.` then `Last error from Discord: {last_error}.` |
| Switched off (amended 2026-10-25, follow-up lane) | `Switched off after 5 failed posts in a row.` / `Last error from Discord: {last_error}.` / `Fixed it in Discord? Test and switch on. If the webhook is gone, make a new one.` / link `Use a different webhook` / button `Test and switch on`. A `{last_error}` cut at 140 characters ends `…` and takes no full stop. |
| Lapsed | `Posts are paused because your plan ended.` / `We kept your webhook. Posts start again when you are back on Premium.` / `See plans` |
| Free | `Discord messages tell you when a server sim or a log you uploaded is done.` + `See plans` |
| Premium row label | `Discord notifications` |
| Premium row values (amended 2026-10-25, follow-up lane; the free value amended 2026-10-27 by the launch pricing persona round, player 13) | `Premium only`, `Off · set up`, `On · 2 events`, `On · 1 event`, `Off · no events chosen`, `Off · needs attention`, `Paused · plan lapsed` (U+00A0 as in 3.7: between the state word and the dot, and between a count and its noun) |
| /premium line | `Discord notifications when a server sim finishes or a log finishes parsing.` (replaces `Discord notifications (ships December).`; the line is edited when the lane ships, never ahead) |
| Changelog entry | `Added Discord notifications. Premium members can paste a Discord webhook on the account page and get a message when a server sim finishes or a log finishes parsing.` (dated the day it ships, past tense) |

`{error}` and `{last_error}` are Discord's own words for what went wrong (for example `Unknown Webhook`, `Missing Permissions`, `Invalid Webhook Token`). They are shown as plain text, never as HTML, shortened to 140 characters, and the API removes any URL or token from them before they reach the page. No string uses `soon`, an exclamation mark, or `unlock`.

### 6.2 In Discord: the messages

Every message is one POST with one embed, no `content`, `username` `Forever Sixty`, `avatar_url` the Forever Sixty seal, and `allowed_mentions` `{"parse": []}` so a title that contains `@everyone` pings no one. The embed colour is `15055189` (`#e5b955`). The footer text is `Forever Sixty`; Discord adds the time from the embed's `timestamp` (the moment the event finished). The title is the link. Text that members wrote (report and sim titles) has Discord's markdown characters escaped. Counts and durations come from the event's own row. `Result` keeps its bare unit (`1,204 DPS`): the title already says the figure is a sim, so no `(sim)` hint is added (amended 2026-10-25, follow-up lane).

**Server sim finished** (the `sims` job's `Finish`; only a run in state `done` posts)

| Part | Text | Source |
|---|---|---|
| Title (link to `/sim/{id}`) (amended 2026-10-25, follow-up lane) | `Frostbolt's Frost Mage sim finished` when the run's character name is known; `Your Frost Mage sim finished` when it is not (`{name}'s`, always `'s`, the name escaped) | the sims row's character name when it stores one (else the poster adds nothing and the title reads `Your ...`); `specLabel(spec)` |
| Field `Result` (inline) | `1,204 DPS`, or the API headline for gear and talent runs (`+41 DPS from Vis'kag`) | `sims.headline`, falling back to `{dps} DPS` |
| Field `Run` (inline, omitted when the member gave no title) | `Phase 1 BiS check` | `sims.title`, at most 100 characters |
| Field `Run time` (inline) (amended 2026-10-25, follow-up lane) | `3 min 12 s`; `48 s`; `1 h 5 min` | `finished_at` minus `created_at` |
| Thumbnail | the spec's class crest (`/icons/hd/crests/mage.webp`, 80 px) | one crest language; omitted if the crest URL is not absolute |
| Footer | `Forever Sixty` | |

**Log parsed** (the report ingest's `complete`, once per whole-file upload; live logging does not post)

| Part | Text | Source |
|---|---|---|
| Title (link to `/reports/{id}`) | `Tuesday Molten Core finished parsing`; with no title, `Your log finished parsing` | `reports.title`, at most 200 characters |
| Field `Zone` (inline, omitted when empty) | `Molten Core` | `reports.zone` |
| Field `Boss fights` (inline, omitted when the count is 0) (amended 2026-10-25, follow-up lane) | `9 · 7 kills`; `1 · 1 kill`; `3 · no kills` | the report summary's fight and kill counts (the summary's fight count is boss pulls only; trash is already excluded) |
| Field `Visible to` (inline) | `Everyone` (public), `Anyone with the link` (unlisted), `Your guild` (guild), `Only you` (private; posted only when `Include private logs` is on, 3.4) (amended 2026-10-25, follow-up lane) | `reports.visibility` |
| Footer | `Forever Sixty` | |

**Test message** (the `Send test message` call)

| Part | Text |
|---|---|
| Title (link to `/account#discord`) | `Test message from Forever Sixty` |
| Description | `Your webhook works. You will get a message here when a server sim finishes or a combat log you upload finishes parsing.` Built from the events currently on; with one on, only that clause; with none, `Your webhook works. No events are on yet.` With `Include private logs` on and the log clause present, the clause ends `finishes parsing, private logs included.` (amended 2026-10-25, follow-up lane) |
| Fields | none |

Limits respected: titles at most 256 characters (ours are at most 200), field values at most 1024 (ours at most 100), three fields at most (Discord allows 25), one embed. One event is one message; the site never batches, repeats or follows up. A member who uploads ten files in a row gets ten messages; that is the member's own action and is stated in decision 16.

## 7. Phone and desktop

**Never collapses, at any width:** the destination line and the channel id; the notice sentence in switched-off and lapsed; the three checkbox labels (they wrap, never truncate) (amended 2026-10-25, follow-up lane); the destination's two lines; the status line; `Remove`; the input at full width; the free sentence and `See plans`.

Also never, (amended 2026-10-25, connect lane): `Connect a Discord channel` and its explanation, and the connect status line.

**Collapses or wraps:** buttons wrap onto a second row when the panel is narrower than their sum (at 360 the panel's inner width is 292 px: `Send test message` plus `Remove` fit on one row at 280 px; `Test and switch on` plus `Remove` wrap, `Remove` on its own 44 px row). Helpers wrap. The channel id sits on its own line when the line is full. The panel padding is 16 px below 768.

- **360 and 390:** the panel is the rail's width (the page's 18 px gutters, 324 or 354 px). Panel heights (amended 2026-10-25, follow-up lane): saved at most 480 px (540 with the 1-to-4-failures warning), paste form at most 640 px, switched off at most 580 px with a one-line Discord error, plan lapsed at most 480 px; with the longest error the API keeps (140 characters, three lines) the switched-off panel may reach 620 px at 390 and 650 px at 360 — the error is never cut shorter, because it is the one thing the member can act on. At 1440 the same states are at most 470, 650, 565 and 490 px (amended after the follow-up captures: the form measured 647 and the one-line switched-off 561, the 8 px being `StatePanel`'s `md:py-5` over `py-4`); saved with the 1-to-4-failures warning and a 140-character error may reach 590 px at 360 and 390 (measured 584 at 390). The heights grew from the first lane's caps (420, 560, 520) by the third checkbox row (44 px), the destination's second line (20 px) and the steps (about 70 px, form only). The Premium row's value wraps to two lines at 360 for the longest values. Connect lane re-check (amended 2026-10-25): the paste view with the back link under `Save webhook` (connect on) is at most 700 px at 360 and 390 and 710 px at 1440 (measured 683, 683 and 691); the connect-off form keeps 640 and 650. A switched-off CONNECTED webhook (the notice also holds `Connect a different channel` and its two-line note) is at most 620 px at 360 and 390 and 605 px at 1440 (measured 592, 592 and 600, with a one-line error).
- Connect heights (amended 2026-10-25, connect lane): the connect state is at most 340 px at 360 and 390 and 320 px at 1440; saved and connected is at most 520 px at 360 and 390 and 500 px at 1440 (the pasted caps plus the reconnect note, 40 px). The button hugs its label and never wraps its text; at 360 it is 232 px of the 292 px inner width.
- **768 to 1023:** the right rail drops under the left column (single column below `lg`), the panel is the full content width with 24 x 20 px padding, and nothing stretches the input beyond the panel.
- **1024 to 1279:** the rail is 288 px wide; see 3.1.
- **1440:** 395 px wide, as the mock.
- **2000:** the account page's container is the existing centred 1344 column, so the panel is the same 395 px and the page is centred. The block adds no `max-width` of its own and never widens. Capture at 2000 to confirm it is centred with the rail.
- **Keyboard on phone:** the URL input uses `inputmode="url"`, so the keyboard that opens has a paste key; `enterkeyhint="done"`.

## 8. Rulings this spec makes

1. **The block is one panel, and a free member gets a sentence, not a panel.** Same ruling as the history spec's 3.9: a panel with a heading and one sentence is a heading over one line.
2. **The block sits inside `#plan`, under the Plan panel,** 16 px away. It is a benefit of the plan and the Premium row links to it; the rail's other panels (Devices, You) are the account's own settings, 32 px apart.
3. **Events are checkboxes saved on change, not a form with a Save** (two events and the private-logs modifier, (amended 2026-10-25, follow-up lane)). The URL is saved once and the events are three booleans, so a second Save button would be a button with nothing to say. Alternative in decision 8.
4. **Re-enable is a test.** A switched-off webhook comes back only through a post that Discord accepts, so `Test and switch on` is both the proof and the switch. No separate `Re-enable` route and no way to switch on a webhook that does not work.
5. **The saved view shows the channel id and nothing derived from the token.** The page cannot offer a `Copy`, a `Reveal` or a masked tail, because it never holds the token.
6. **The messages are quiet.** One embed, three fields at most, no `content`, no pings (`allowed_mentions` empty), no follow-ups, no digest. The site tells the member once that a thing is done.
7. **A failed sim does not post** (decision 4). The event is "a server sim finished".
8. **Private uploads post only when the member turns on `Include private logs`, and say who can open them** (the `Visible to` field and the note under the toggles) (amended 2026-10-25, follow-up lane). The first lane posted every upload and warned in a note; a veteran player read that as a leak into a shared guild channel. A separate off-by-default box removes the warning and keeps the choice the member's.
9. **A lapsed plan keeps the webhook and says why, in the member's own words,** not the API's reason code. `Remove` stays available so a member can delete a credential at any time.
10. **No history, counters or delivery log** in the block (3.5).
11. **The destination is Discord's own webhook name, not a member label** (amended 2026-10-25, follow-up lane). A webhook's name is set by the member in the same Discord dialog that makes the URL; step 2 asks them to name it something they will recognize. A second field in our form would be a second thing to type for the same word. The shortened channel id stays beneath it to tell two webhooks of the same name apart.
12. **The account page's loading state is the real grid** (amended 2026-10-25, follow-up lane). A Discord-only reserve inside an absent rail cannot hold; the reserve is the page's (9.1).
13. **Connect is the first path, paste stays one link away** (amended 2026-10-25, connect lane). The member should not have to copy a URL; the product spec's non-goal (no bot, no guild linking) holds, because the OAuth scope `webhook.incoming` gives only a webhook. The paste path stays for members who do not want to authorise an app and for deployments without the Discord application.
14. **A connected webhook says `Connected to a Discord channel`, never the Discord-side name** (amended 2026-10-25, connect lane). The name defaults to `Forever Sixty`, which tells the member nothing; a pasted name was typed by the member and does.
15. **We delete what we made, never what the member made** (amended 2026-10-25, connect lane). A connected webhook is created by our authorisation, so replacing or removing it deletes it from Discord (best effort); a pasted webhook is the member's and stays.
16. **No Discord branding** (amended 2026-10-25, connect lane). No logo, wordmark glyph or blurple; the button is our own gold text button, and the word `Discord` appears as text only. The Developer Portal application is named exactly `Forever Sixty` with our seal as its icon, because Discord's dialog shows both.

## 9. Performance, polish, and acceptance

**Limits.** The block reads `/v1/me`, which the account page already holds: no request of its own on mount, no skeleton of its own inside the ready view, no `LoadError`. It renders in the same paint as the Plan panel it sits under, so nothing appears after hydration beyond what the page itself reveals. The loading state before `/v1/me` answers is specified in 9.1 (amended 2026-10-25, follow-up lane). No new font, no icon, no image, no new dependency; the Discord name is text. `Account.svelte` gains one import; `AccountDiscordBlock` is small enough to stay in the account island (no lazy chunk). Lighthouse (`web/lighthouserc.json`): `/account.html` is audited signed out, where the page renders its sign-in state and none of this, so the budgets hold by construction (CLS at most 0.05, TBT at most 100 ms, performance and accessibility at least 0.95, 0.9 where the file says so). Add no signed-in `/account` run to the config.

**No layout shift.** The block is the last panel in the rail, and nothing below it but the footer moves. The load-time shift is governed by 9.1 (amended 2026-10-25, follow-up lane). After `me` resolves, every state's height is a function of `me` alone, known in the first render. The 21 px status line is always reserved, so a test result, `Saved.` or an error never moves the buttons. Changes the member causes (Save turning the form into the saved view, Remove turning it back) change the panel's height once and move the footer; nothing else on the page is below it. A member-caused shift is the only one allowed; the Playwright check below measures it at 390 and 1440.

### 9.1 The account page's loading reserve (amended 2026-10-25, follow-up lane)

**Measured today.** At 390 the loading page (three stacked `Skeleton` blocks: `account-characters-skeleton`, `account-identity-skeleton`, `account-more-skeleton`, 1396 px with gaps) gives way to a ready page whose rail arrives all at once: load-time CLS 0.35 to 0.45 with or without the Discord block. At 1440 it is 0.04 to 0.10, and the block adds about 0.04 to 0.06 of that. The loading page is one full-width column of bars; the ready page is a header row, then a 12-column grid with the left column at 8 and the rail at 4. They share nothing, so no panel-sized fix can hold. Reference: the in-game character panel, which draws its frames before it draws what is in them; the first paint is the final layout with the contents missing.

**Scope.** The whole `/account` loading state, for a visitor with the session hint (`data-session="1"`), in one stroke: the header row, the left column, and the whole rail (Devices, You, Plan, Discord). Not other pages. A visitor without the hint keeps the collapsed state the `.session-reserve` rule already gives (no skeletons, the sign-in prompt only). Lighthouse still audits signed out, so its budgets are untouched.

**What the loading state draws.** The ready page's skeleton at its real positions, with the existing `Skeleton` component (the bars the page already uses; no new look; `prefers-reduced-motion` as it already handles):

| Region | Skeleton | Height | Position |
|---|---|---|---|
| Header row | the title stays real text; a reserved row beside or under it | the ready header row: 44 px at 1024 and up; below 1024 the ready row (title, 12 px, identity line with `Sign out`), measured and written in `layout.ts`, expected 84 px | at 1024 and up the page's `lg:flex-row` row; below, stacked |
| Left column (`account-left-skeleton`, `lg:col-span-8`) | one `Skeleton`, `lines={3}` | `max(stored left height, left floor)`; the left floor is the measured height of an empty Characters panel plus an empty Your reports panel plus the 32 px gap (about 300 px at 1440) | left 8 columns at 1024 and up, first in the stack below |
| Rail (`account-rail-skeleton`, `lg:col-span-4`) | a column whose `min-height` is `var(--account-rail-reserve)` when stored, else the sum of its floors; inside it the blocks below | | right 4 columns at 1024 and up, after the left column below, 32 px gap |
| Rail: Devices (`account-devices-skeleton`) | `Skeleton` | 232 px at 1440 (empty devices, measured on the fixture) | first |
| Rail: You (`account-you-skeleton`) | `Skeleton` | 192 px at 1440 | 32 px under Devices |
| Rail: Plan (`account-plan-skeleton`) | `Skeleton` | 219 px at 1440 | 32 px under You, inside a wrapper with `gap-4`, as `#plan` |
| Rail: Discord (`account-discord-skeleton`) | `Skeleton`, no panel chrome | by `data-discord` (table below) | 16 px under Plan, in the same wrapper |

A guild panel above Plan, extra devices, or any panel taller than its floor is absorbed by the stored rail height: the container's `min-height` is the stored total, so the blocks inside may sit at slightly different positions from the ready panels (they vanish when the page is ready, and nothing persistent sits among them). Below 1024 the floors measured at 390 are written beside the 1440 ones in `web/src/lib/account/layout.ts`; the three old constants (`CHARACTERS_SKELETON_MIN_H`, `IDENTITY_SKELETON_MIN_H`, `MORE_SKELETON_MIN_H`) are replaced by the new ones, and `SIGNED_OUT_MIN_H` stays the loading page's total by the same arithmetic.

**What the page knows before first paint.** A third block in `Base.astro`'s inline script, written like the server-runs and history blocks (lines 131 to 185), runs only when `location.pathname` is `/account` or `/account/` and `<html data-session="1">`. It reads two localStorage values:

1. The cached `/v1/me` (`data-me-storage-key`, as the other blocks). From it, one attribute on `<html>`: `data-discord` with exactly one of `none`, `free`, `connect`, `form`, `saved`, `off`, `lapsed` ((amended 2026-10-25, connect lane): `connect` is `form` with `notifications.connect` true), by the same precedence as `discordBlockState` (4.1): no cache, or no `notifications` key: `none`; otherwise `lapsed`, `off`, `saved`, `form`, `free` by 4.1's rows. A unit test runs the script's expression against every fixture `discordBlockState` is tested with and asserts equal keys.
2. The last rendered heights (`fs.accountReserve`, named in a `data-account-reserve-key` attribute): `{"v":1,"lg":{"left":N,"rail":N},"stack":{"left":N,"rail":N}}`. The bucket is `lg` when `matchMedia('(min-width: 1024px)').matches`, else `stack`. It sets `--account-left-reserve` and `--account-rail-reserve` in px on `<html>` when each value is above 0 and at most 4000. A malformed value, a wrong `v` or unavailable storage sets nothing and the floors apply (the other blocks' pattern: `try`/`catch`, silent).

**What writes the heights.** `Account.svelte`, once two animation frames after the ready view mounts, and again on `pagehide` (so the Your reports list that arrives later is counted), writes the rounded `getBoundingClientRect().height` of `account-left-col` and `account-rail-col` (two new test ids on the existing column wrappers) into the current bucket, keeping the other bucket. It skips a write when either height is 0 or above 4000. Sign-out removes the key.

**Discord skeleton height by state.** Written in `layout.ts` as one table. These are targets (the built panel's measured height at 1440 and at 390, rounded up to 5), replaced by the measured values when the lane lands, and used only when no rail height is stored:

| `data-discord` | Draws | 1440 | 390 |
|---|---|---|---|
| `none` | nothing (no block; no row) | 0 | 0 |
| `free` | one 44 px line, no panel chrome | 44 | 44 |
| `connect` (amended 2026-10-25, connect lane) | one block | 320 | 340 |
| `form` | one block | 610 | 630 |
| `saved` | one block | 460 | 475 |
| `off` | one block | 540 | 565 |
| `lapsed` | one block | 475 | 475 |

The `none` and `free` states draw no panel, because the ready page draws none.

**Acceptance numbers** (Playwright, `/v1/me` stubbed with a 400 ms delay so the loading state is on screen; the Chromium layout-shift observer over the whole load, full page):

| Case | 390 | 1440 |
|---|---|---|
| Warm (localStorage holds a cached `/v1/me` and `fs.accountReserve` from a previous ready page), each of the seven `data-discord` states (amended 2026-10-25, connect lane) | CLS at most 0.05; footer top moves at most 8 px | CLS at most 0.03; footer top moves at most 8 px |
| Cold (session hint set, nothing stored) | recorded, not gated; at most 0.15 | recorded, not gated; at most 0.08 |

Cold is a new device; its numbers go in the PR so a regression is visible. The hint-less signed-out visitor is untouched and keeps its existing CLS 0.

**Hydration order.** The attribute and the variables are set before first paint by the inline script, so the first paint already has the right heights; the island does not set them. No new request, no font, no image.

### 9.2 Gold checkboxes, site-wide (amended 2026-10-25, follow-up lane)

One rule set in `web/src/styles/global.css`, in the base layer, not a Discord rule:

```
input[type='checkbox'] { accent-color: var(--color-gold); color-scheme: dark; }
input[type='checkbox']:disabled { accent-color: var(--color-muted); }
input[type='checkbox']:focus-visible { outline: 2px solid var(--color-gold); outline-offset: 2px; }
```

The checked box is gold with the browser's dark tick; unchecked it is the dark-scheme outline. A disabled box is `--muted`, and its label is `--muted` (the Discord lapsed rows already do this; every disabled-label pair gets the same colour from this rule, not from local classes). Focus-visible is the 2 px gold ring, 2 px offset, on the box, as the Discord rows ship. The rule covers all ten surfaces: Account `anonymize`, planner `SharePanel`, rankings `RankingsFilters`, guide `GuideActionRail`, report `EventsView` and `FilterBar`, sim `BuffPanel`, `SettingsSheet` and `ReportOptions`, and the Discord rows. The implementer greps for `type="checkbox"` and for any local `accent-*`, `appearance-none` or hand-drawn checkbox; a local override that draws a non-gold box is removed, one that draws its own gold box is left only if it passes the focus-visible and disabled looks above. Hit area is unchanged: every row's label stays 44 px tall on phone. Contrast: gold `#e5b955` on `--raised` is above 3:1 for the box edge; the tick is the browser's own contrast colour.

**Tests the lane adds.**

- Unit, web: `webhookShapeError` accepts `discord.com` and `discordapp.com` with an optional `v6` to `v10` segment, and rejects `http:`, other hosts, `discord.com.evil.example`, userinfo, a missing token, a trailing path, an empty string and text that is not a URL; `discordBlockState` covers every row of 4.1, including lapsed beating switched-off.
- Unit and job, API (the product spec's list): URL validation; the embed builder (field counts, limits, escaping, `allowed_mentions`); a finished run posts once; a webhook failing five times in a row is switched off; no response body and no log line contains the URL or the token.
- E2E (Playwright, API stubbed to section 10): an invalid URL shows the fixed message and the typed text appears nowhere in the DOM outside the input; save, and the channel id shows while a grep of the page HTML for the fixture token finds nothing; test sent and test failed; an event turned off; the five-failure state; the lapsed state; remove with the confirm; the free sentence.
- An axe pass on every state. A CLS check of `/account` at 390 and 1440 through save and remove, asserting no shift above 0.02 other than the member-caused one.

**Acceptance screenshots the implementer attaches** (fixture build, API stubbed; 390 and 1440 unless stated; the mock's boards are the reference for each):

1. `/account`, payer with a saved webhook: the rail with Plan and the Discord panel under it, at 1440, 2000 and 390, the Premium row reading `On · 2 events`.
2. The paste form: empty, input focused, and the invalid-URL message (one capture with the typed value visible in the field and absent from the message).
3. Saved: default, `Saving…`, `Sending…`, `Sent. Check the channel.`, a Discord error in `--wipe`, and the 1-to-4-failures warning.
4. Switched off after five failures with the last error, and after `Test and switch on` succeeds.
5. Plan lapsed, webhook kept, boxes read-only.
6. Remove confirm, and the replace form.
7. Free member: the sentence under Plan at 390, 1440 and 2000, and the Premium row `Off`.
8. Every control's focus-visible: input, each checkbox, `Save webhook`, `Send test message`, `Remove`, `Use a different webhook`, `See plans`; and hover on the three buttons at 1440.
9. The three messages as they arrive in a real channel (sim, log, test) from a staging webhook, in the desktop and the phone Discord client, beside the mock's embed boards.
10. The Premium row's seven values at 390 and 1440.
11. 1024 and 768 captures of the panel (rail and stacked), and 360 for the longest state.
12. The CLS numbers printed by the Playwright check at 390 and 1440 (member-caused shifts, as before).
13. A same-scale side-by-side of the shipped rail with the mock's matching panel at 1440.

New or changed in the follow-up lane (amended 2026-10-25, follow-up lane) (390 and 1440 unless stated):

14. The paste form and the replace form with the three steps, the dead band shortened (measure and print the input-bottom to `TRUST` first line distance: 21 px), and the invalid-URL state.
15. The destination with a name (`Posts through “Raid pings”`, `Channel 1198…4568`) and without (`Posts to a Discord channel`), plus a 40-character-plus name truncated with `…`, and a name containing `@everyone` and `<b>x</b>` shown as plain text.
16. Saved with the three checkboxes: third on, third off, second off (third disabled), and lapsed (all three disabled).
17. The warning with `Send a test message to check.` for one failure, two failures, and four failures with a 140-character error.
18. Switched off with the four-line notice (one-line error, and the 140-character error at 390 and 360), with `Use a different webhook` inside the notice and the panel height printed against the caps in 7.
19. Remove confirm: `Use a different webhook` absent.
20. The Premium row's seven values at 390, 1024 and 1440, with the 1024 capture proving `2 events` never splits; and the feature-off page with no Discord row.
21. The loading state at 390 and 1440 with the cached `/v1/me` for each of the seven `data-discord` states (amended 2026-10-25, connect lane), then the ready page beside it (same scale), the printed warm and cold CLS numbers, and the footer's top before and after.
22. A capture of each of the ten checkbox surfaces at 1440 showing a gold checked box; one disabled and one focus-visible capture of the Discord row and of one non-Discord surface (the rankings filters).
23. The sim embed with and without the character name, the log embed with `Boss fights`, the test embed with private logs on, all three in the mock's boards and in a real channel when the owner supplies a webhook.

New or changed in the connect lane (amended 2026-10-25, connect lane) (390 and 1440 unless stated; the mock's `CONNECT ON` boards are the reference):

24. The Connect state with `connect: true`, nothing saved: default, button hover, button focus-visible, `Opening Discord…`, and the 360 capture; the panel height printed against 340 and 320.
25. The paste view with connect on (with the back link), and the paste-first state with `connect: false` (the connect button, disclosure and back link absent). Re-check additions: the 360 capture of the paste view; the switched-off connected webhook at 360, 390 and 1440.
26. The three return outcomes from a real redirect in staging: `connected` (saved connected, focus on `Send test message`), `denied`, `failed`; each with the URL bar capture showing `/account#discord` with no `discord=`; a reload showing no status line.
27. Saved connected: the destination `Connected to a Discord channel` over `Channel 1198…4568` with no Discord-side name anywhere in the DOM text; the reconnect link and its note at 360 and 390; a saved pasted webhook beside it.
28. Remove confirm for a connected and for a pasted webhook (the two questions), and replace view for a pasted webhook with connect on.
29. Discord's own dialog (staging) showing the application named `Forever Sixty` with the seal icon, and the new webhook listed under the channel's Integrations; after a reconnect and after a remove, the old webhook gone from Discord.
30. The loading state for the seventh `data-discord` value (`connect`), warm CLS numbers as in 9.1.
31. A grep of the built page and the bundle for the Discord logo asset and for `5865F2` finds nothing.

## 10. API contract this spec assumes

The API lane builds exactly this; if it differs, change the contract here first.

- **`GET /v1/me`** carries `notifications: {discord: {channel_id, name, events: {sim_finished, upload_parsed, upload_private}, disabled_reason, last_error, failures} | null}` (amended 2026-10-25, follow-up lane). `channel_id` is a string (Discord's channel snowflake). `name` is `null` or a string (see the next item). `events` has exactly those three keys. `disabled_reason` is `null`, `"failed five times"` or `"plan lapsed"`. `last_error` is `null` or Discord's error text with any URL or token removed. `failures` is the consecutive failure count, 0 to 5. The block is returned for a lapsed member (it is how the page knows the webhook was kept). `entitlements.notifications` is the gate; the web never decides it. Channels are keyed (`discord` now), so a later channel adds a key and changes no field.
- **Connect additions** (amended 2026-10-25, connect lane).
  - `GET /v1/me` carries `notifications.connect: boolean`: true only when the Discord application is configured (client id, client secret, registered redirect) and the deployment's encryption key is present. It is absent whenever `notifications` is.
  - A saved `discord` block carries `source: "connected" | "pasted"`; rows saved before this change are `pasted`. For a connected webhook `name` is `null`.
  - `GET /v1/notifications/discord/connect` needs the session cookie and `Can(FeatureNotifications)`. It answers 302 to `https://discord.com/oauth2/authorize?response_type=code&client_id=…&scope=webhook.incoming&redirect_uri=<callback>&state=<value>`. `state` is a random value held in an httpOnly, `SameSite=Lax` cookie scoped to the callback path (the Battle.net sign-in pattern), compared in constant time on return, cleared on first use, and expiring in 10 minutes; the callback also requires the member's session, so the saved webhook can only bind to the account that started the flow (amended 2026-10-25, connect lane, controller ruling). Without a session or entitlement it answers 302 to `/account?discord=failed#discord`. When `connect` is false it answers 404.
  - `GET /v1/notifications/discord/callback`: with no `code`, or with `error=access_denied`, 302 to `/account?discord=denied#discord`. With a bad or expired `state`, a failed exchange, an answer with no `webhook`, or a store error, 302 to `/account?discord=failed#discord`. Otherwise it exchanges the code (`POST https://discord.com/api/oauth2/token`, form-encoded, the same `redirect_uri`), takes `webhook.{id, token, url, channel_id, guild_id, name}`, stores the sealed URL, the channel id and `source = "connected"`, and answers 302 to `/account?discord=connected#discord`. The access and refresh tokens are discarded unread, never stored, never logged. No response, redirect or log line contains the code, the token or the URL.
  - A connect over an existing webhook keeps its events; with nothing saved the events are the defaults (sim on, log on, private off). `failures`, `disabled_reason` and `last_error` are reset.
  - **The delete-from-Discord rule.** When a connected webhook is replaced (by a new connect, or by a `PUT` with a URL), removed (`DELETE`), or its account is deleted, the API calls Discord's webhook delete (`DELETE /webhooks/{id}/{token}`) with the webhook's own token: 5 s timeout, no retry, after the database change has landed. A failure is logged without the URL or token and never shown. A pasted webhook is never deleted from Discord. A lapsed plan deletes nothing.
  - A `PUT` with a `webhook_url` stores `source = "pasted"`.
- **`PUT /v1/notifications/discord`** with `{webhook_url?, events}`, `events` carrying all three keys (amended 2026-10-25, follow-up lane). `webhook_url` is **required when nothing is saved or when replacing** and **omitted** for an events-only change (the page never has the URL, so it cannot resend it). 200 answers the same block as `/v1/me`'s, never the URL. The API validates the host (exactly `discord.com` or `discordapp.com`, https only) and the path (`/api/webhooks/<number>/<token>`, with an optional `v6` to `v10` segment), refuses a query or fragment (a pasted webhook URL has neither; the web pre-check mirrors this), and rejects userinfo and ports. A new URL resets `failures` to 0 and clears `disabled_reason` and `last_error`. 400 or 422 for a bad URL (the body names no part of the URL); 402 when `Can(FeatureNotifications)` is false. The web shows its own words and never the response body.
- **The webhook `name` and the third event flag** (amended 2026-10-25, follow-up lane). This is an API change from the first lane. (a) One new nullable text column on the stored webhook, `webhook_name`: the `name` field of the same `GET /webhooks/{id}/{token}` response that supplies `channel_id`, trimmed, control characters removed, cut at 80 characters, stored on every save that sends a `webhook_url` (a replace refreshes it; an events-only `PUT` leaves it). `null` when Discord returns none. The web escapes it and shows at most 40 characters. Discord does not return the channel's name (that needs a bot), so the page cannot say `#raid-pings` and does not try. (b) One new boolean event flag, `upload_private`, beside the other two: default `false` on a new webhook; webhooks saved before this change keep posting private logs (their flag is set `true` once, so no member's behaviour changes silently). The report ingest posts when `upload_parsed` is on and either the report is not private or `upload_private` is on. The test message reads the flag (6.2). An events-only `PUT` validates all three keys. The Premium row's count is `sim_finished` plus `upload_parsed` only.
- **`channel_id` provenance (decision 1).** A Discord webhook URL contains the **webhook id** and the token, not the channel id. The channel id comes from Discord's `GET /webhooks/{id}/{token}` response (`channel_id`). The API calls it once on save, with a 5 s timeout. If Discord does not answer or the webhook does not exist, the save fails with 422 and the invalid-URL message, which also proves the webhook exists before it is stored. The webhook id itself is never returned.
- **`DELETE /v1/notifications/discord`** removes the URL and the settings. 204. Allowed for a lapsed member. The URL is also deleted with the account.
- **`POST /v1/notifications/discord/test`** posts the test embed (6.2) and answers 204, or a non-2xx with `{error: "<Discord's error text>"}` (URL and token removed, 140 characters at most). Limited to 5 per minute per member; 429 beyond. A successful test sets `failures` to 0 and clears `disabled_reason` when it was `failed five times`; a failed test changes no counter. 402 for a lapsed member.
- **Counting.** A failed event post (non-2xx other than 429, a timeout, a DNS or TLS failure) increments `failures` and stores `last_error`; a success resets `failures` to 0. A Discord 429 waits for the next event and does not count. At the fifth consecutive failure the API sets `disabled_reason: "failed five times"` and stops posting. A lapsed plan reads `disabled_reason: "plan lapsed"` while the URL stays; no post is attempted.
- **Posts** come from the job that produced the event, never a request handler, with a 5 s timeout and no retry beyond the next event. The payload is 6.2. The sim event is the `sims` job's `Finish` (state `done` only); the upload event is the report ingest's `complete`. Both carry the title, visibility, counts and duration the embeds need, so the poster makes no further read. The sim event also carries the character's name when the `sims` row stores it (amended 2026-10-25, follow-up lane); if it does not, the title reads `Your ...` (6.2) and no column is added for this.
- **Gating** of the account block: `me.entitlements.notifications`.

## 11. Decisions the product spec left open (for the controller to rule)

1. **`channel_id` is not in the URL.** The product spec says the page shows "the channel id segment" of the URL. The path holds the webhook id and the token; the channel id needs one call to Discord (section 10). Recommended: the API fetches it on save, which also proves the webhook exists. Alternative: show `Webhook {id}` using the first path segment, which is not secret but names nothing a member recognises.
2. **Placement:** the right rail inside `#plan`, under the Plan panel (3.1). Alternative: the left column under Characters, where there is room (the rail ends lower than the left column at 1440). Recommended: under Plan, so the Premium row and the block read as one console.
3. **Free member: a bare sentence, no panel** (3.2, ruling 1). Alternative: a panel holding the sentence. Recommended: the sentence.
4. **A failed sim does not post.** The promise is "a server sim finished". A failed run is already on `/sim` and in the band's list. Alternative: post a failure embed under the same toggle (`Your Frost Mage sim failed`, field `Reason`), which tells a member who closed the tab that the run will never finish. Recommended: no, until a payer asks.
5. **Private uploads post only on request** (amended 2026-10-25, follow-up lane): the `Include private logs` box, off by default (3.4, ruling 8), with `Visible to: Only you` on the message. Alternative: never post private logs; or post them by default, as the first lane did.
6. **No mentions, ever** (`allowed_mentions: {parse: []}`) and escaped member text, so a title with `@everyone` does nothing. Not optional; listed so the API lane builds it.
7. **The panel is `Discord` now and becomes `Notifications` when the second channel ships** (3.6). Alternative: label it `Notifications` now with a `Discord` sub-heading, which is a heading over one section today.
8. **Events save on change** (3.4). Alternative: one `Save` for URL and events together, which needs a visible dirty state.
9. **Re-enable is `Test and switch on`** (ruling 4), with no extra route. Alternative: a `Switch on` button that clears the flag without a test, which lets a dead webhook burn five more posts.
10. **The URL field is plain text, not masked** (3.3). Alternative: `type="password"`, which hides it from a shoulder surfer or a screen share but invites the password manager to save it and makes a bad paste harder to see.
11. **Accepted hosts:** exactly `discord.com` and `discordapp.com`, https, per the product spec. `ptb.discord.com` and `canary.discord.com` are rejected. Alternative: accept those two subdomains, which Discord's own Copy Webhook URL button may give members on those clients; the helper and the invalid-URL message stay unchanged either way.
12. **Failure counting:** a Discord 429 and a manual test do not count; every other failed event post does (section 10). The product spec's "fails five times in a row" does not say.
13. **Webhook name on the card** (amended 2026-10-25, follow-up lane). Ruled in: the saved view shows Discord's webhook name (`Posts through “Raid pings”`) with the shortened channel id under it (3.2, ruling 11). Alternative: an optional label field in the form (`Call it`, `#raid-pings`), which would read as a channel name but costs a second field to type; or the id alone.
14. **Crest thumbnail and our own avatar on the messages.** The sim message carries the spec's class crest as its 80 px thumbnail (one crest language), and every message carries the Forever Sixty seal as the sender's avatar and `Forever Sixty` as its name, overriding the webhook's own. Needs absolute https URLs for both images; if the crest cannot be reached Discord omits it and the message is still whole.
15. **The `Run time` field** (was `Took`; finish minus start) (amended 2026-10-25, follow-up lane) is an addition to the title, result and link the product spec named; it is the one number a member did not already know. Drop it if the controller wants two fields.
16. **No coalescing.** A member who uploads ten log files in a row gets ten messages. Alternative: after three messages in ten minutes, post one `and 7 more finished parsing`. Recommended: not in this lane; revisit if upload habits make it noisy.
17. **The trust sentence** (`We store the URL encrypted, never show it again, and delete it with your account.`) is true only if the API stores the URL encrypted (it must be recoverable to post). The API lane confirms, or the sentence loses `encrypted`.
18. **The Premium row's values** (3.7) are new wording for the state; `On · n events` follows the server-sims row's `On · n runs this month`; `Off · needs attention` and `Off · Premium` (amended 2026-10-25, follow-up lane) replace `Off · failed five times` and `Off`.
19. **No live update.** A switch-off that happens while the account page is open shows on the next load. Alternative: re-read `/v1/me` on tab focus, which the page does not do for anything else.
20. **Removing `shipsMonth.notifications`** from `billing/copy.ts` and the tests that reference `Ships December`; the /premium line and the changelog entry change on the day the lane ships, not before.
21. **Connect is first when configured** (amended 2026-10-25, connect lane). Alternative: keep paste first with a secondary connect button. Recommended: connect first, per the owner's direction.
22. **A connected webhook shows no name** (amended 2026-10-25, connect lane). Alternative: a link `Open in Discord` built from `guild_id` and `channel_id` (`https://discord.com/channels/<guild>/<channel>`), which would let the member see the channel's name in Discord. Controller ruling 2026-10-25: IN this lane — the API stores and returns `guild_id` for a connected webhook, and the saved view carries the text link `Open in Discord` INLINE on the channel line — `Channel 1198…4568 · Open in Discord` (13 px; the link part underlined `--text`, with a 44 px tall hit area given by padding, not by a taller line; `href="https://discord.com/channels/<guild_id>/<channel_id>"`, `target="_blank" rel="noopener"`), so the member can see the channel's name in Discord itself without the panel growing past §7's caps (ruling revised 2026-10-25 after the build measured the third-line version at 543 px against the 500 px cap). A pasted webhook has no guild id (Discord's lookup gives `guild_id` too — store it for pasted webhooks as well, so both sources get the link when it is known).
23. **A connected webhook offers no paste replace** (amended 2026-10-25, connect lane). A member who wants a pasted URL removes the connection first. Alternative: a second link under the reconnect link; rejected as two secondary links stacked for a rare path.
24. **`denied` and `failed` are the only error words** (amended 2026-10-25, connect lane): the API knows no more. Alternative: a reason code (`?discord=failed&why=scope`); Discord's answers do not give one.
25. **Events are not chosen before connecting** (amended 2026-10-25, connect lane). The redirect would have to carry them; the defaults suit most members and the saved state edits them in one click.
