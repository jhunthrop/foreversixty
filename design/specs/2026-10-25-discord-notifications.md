# Discord notifications: experience spec

Author: ux-designer. 2026-10-25. Lane 5 of the premium experience. Product spec: `docs/superpowers/specs/2026-10-24-premium-experience-design.md` section 4.6 (the promise, the rulings), section 5 (the account page is the Premium console; the upsell sentence), section 2 (anti-goals), section 8. References: `docs/tenets.md` 1, 4, 9, 10, 11, 13, 14; `design/specs/2026-10-24-server-runs.md` (the finish event this rides on; the account row grammar), `design/specs/2026-10-24-character-history.md` (the free owner's sentence, ruling 3.9). Mock: `design/mocks/gen_discord_notifications.py`, render `design/mocks/renders/discord-notifications-states.png` (1440 and 390: the rail with the block, every state, the controls, the Premium row values, and the embeds as Discord draws them).

Surfaces changed: a new `web/src/components/account/AccountDiscordBlock.svelte` (built only from existing parts, section 3), `web/src/components/Account.svelte` (one mount point, 3.1), `web/src/components/account/AccountPremiumBlock.svelte` (the `notifications` row goes live, 3.8), `web/src/lib/billing/copy.ts` (the row strings, the /premium line; `shipsMonth.notifications` is deleted), a new `web/src/lib/notifications/copy.ts` (every string below) and a new pure `web/src/lib/notifications/discord-block.ts` (the state function and the URL shape check). The embeds are built by the API; section 6.2 fixes their text.

## 0. The experience in one paragraph

A payer opens `/account` and, directly under the Plan panel, finds a `Discord` panel. With nothing saved it is a paste form: one field for a webhook URL, two checkboxes for what to post, one gold `Save webhook`. Once saved the panel says `Posts to channel 1098765432101234567`, shows the two checkboxes, and offers `Send test message` and `Remove`. When Discord rejects five posts in a row the panel says so in one sentence, names Discord's last error, and offers `Test and switch on`. When the plan ends the panel says posts are paused, keeps the webhook, and says why. A free member sees one sentence and `See plans` in the panel's place. The webhook token never appears on the page, not even masked.

## 1. The one question and the answer that must be visible

**"Will I hear about it where my guild already talks?"**

The answer is a state word and a destination: `Posts to channel 1098765432101234567`, or the sentence that says why it does not.

| | 360 / 390 px phone | 1280 / 1440 px desktop |
|---|---|---|
| Payer, saved | The `Posts to channel ...` line is the first line under the panel label. The whole panel (label, destination, two checkboxes, note, buttons, status line, `Use a different webhook`) is at most 420 px tall at 360 and 390, so it fits one 640 px screen once scrolled to (the Premium row's `On · 2 events` link lands on it, `#discord`). | The panel sits directly under Plan in the right rail (395 px wide at 1440). On a payer's account page Plan ends near y 895 at 1440, so the block begins at or below the fold; the Premium row above it answers the question in one line (`On · 2 events`), and the block is one tap or one scroll away. |
| Payer, nothing saved | The intro (two lines) and the URL field start inside the panel's first 160 px. | Same. |
| Payer, switched off or paused | The notice sentence is the first thing under the destination line, with a 2 px bar (`--wipe` switched off, `--gold` paused). | Same. |
| Free | One sentence, 44 px high, under Plan. | Same. |

The Premium row in the Plan panel is the answer for a reader who never scrolls to the block: `On · 2 events`, `Off · failed five times`, `Paused · plan lapsed`, `Off · set up`.

## 2. The reference, and exactly what we borrow

**Discord's own Integrations, Webhooks page** (Channel settings). The job players already know: one webhook is one row with a destination (the channel); it can be copied and deleted; nothing about the token is shown after creation except through an explicit Copy. Borrowed: the destination named first; `Copy Webhook URL` as the instruction we point the member to (the helper line names the exact menu path); Delete as a plain, confirmed action. Not borrowed: the webhook's avatar and name editor (the message carries our own username and avatar, 6.2).

**GitHub's repository Webhooks settings.** The best-known "paste a URL, choose events, test it" form. Borrowed: events as checkboxes under the URL, not a list of toggles; a failing delivery is shown with its error where the webhook is listed; a failed webhook gets one line saying why and one action to retry (their `Redeliver`; ours is `Test and switch on`); the secret is write-only. Not borrowed: the payload preview, the delivery history list, the content-type picker, the SSL toggle. The member has two events and one channel; a deliveries table would be a wall of rows for nothing (tenet 11).

**Warcraft Logs' Discord integration.** The incumbent for this audience: a guild links a channel and WCL posts a report link. Not borrowed: guild-level linking and channel pickers (a Discord bot with OAuth; the product spec's non-goals rule out a guild webhook). Ours is a member's own webhook.

**The existing account page** (`web/src/components/Account.svelte`): `StatePanel` label row (gold dot, 11 px label); the email form's `label.label` + `h-11` input + `SECONDARY_BUTTON_FIXED` button; the `anonymize` checkbox row (native checkbox in a 44 px label); the device list's in-place `Revoke` confirm (a question and two buttons replace the row); `min-h-[21px]` reserved status lines; the plan panel's underlined `See plans`. Every control below is one of those.

## 3. Layout, region by region

Tokens (all existing, `web/src/styles/tokens.css`): `--color-bg #07090d`, `--color-raised #0d111a` (panel ground), `--color-line #262e40` (panel border), `--color-line-warm #3a3326` (input and button borders), `--color-gold #e5b955`, `--color-gold-hover #f5d27a`, `--color-strong #f2eee4`, `--color-text #e9e4d8`, `--color-muted #9a9484`, `--color-wipe #f0736b` (failure text and the switched-off bar; AA on raised). No new colour. Type: Barlow for text; JetBrains Mono tabular for the channel id; `label` (11 px 700, .12em, uppercase) for the panel label, `Webhook URL` and `Post when`; body 14 px; helper and status 13 px; button text 12 px 700 .06em uppercase (`SECONDARY_BUTTON_FIXED`, `PRIMARY_BUTTON_FIXED`). Radius `rounded-panel` 6 px for the panel, `rounded-control` 4 px for input and buttons. Spacing: panel padding 16 px below 768 and 24 x 20 px from 768 (`StatePanel`'s own `px-4 py-4 md:px-6 md:py-5`), 16 px between regions (`StatePanel`'s `gap-4`), 8 px between a label and its field.

Shared components used: `StatePanel` (label `Discord`, `testid="account-discord"`), `PRIMARY_BUTTON_FIXED` and `SECONDARY_BUTTON_FIXED` from `lib/planner/styles.ts`, the native checkbox row of `anonymize`, the label-plus-`h-11`-input of the email form, the underlined `See plans` link (`billingBlockCopy.seePlans`, `text-text underline underline-offset-2`). `EmptyState`, `LoadError` and `Skeleton` are **not used**: the block makes no read of its own (it renders from `/v1/me`, which the page already holds), so it has no loading or load-error state, and its empty state is the paste form, which is the thing that fills it. Failed actions use the inline status line, the pattern of the email form's `account-error`.

### 3.1 Where it mounts

In `Account.svelte`, the right rail (`lg:col-span-4`), as the **last child of the `#plan` wrapper** (`<div id="plan" class="flex flex-col gap-4">`), after the `StatePanel` that holds the plan line and `AccountPremiumBlock`. The wrapper's `gap-4` puts the Discord panel (or the free sentence) 16 px under Plan, where the rail's other panels sit 32 px apart: Plan and Discord read as one console. Rail order, top to bottom: ratings, Devices, You, Plan, **Discord**. Below `lg` the rail follows the left column in the same order, so on a phone Discord is the last panel on the page; the Premium row's link (`#discord`) gets a payer there in one tap. The panel has `id="discord"` and `scroll-margin-top: 80px` (clear of the sticky header). Mount condition: signed in and `me` loaded. It never mounts for a visitor. At 1024 the rail is 288 px wide (inner 240): the destination line, the buttons and `Off · failed five times` wrap, nothing overflows horizontally, and the input stays full width (its placeholder truncates with an ellipsis).

### 3.2 The panel, by state

`AccountDiscordBlock` chooses one of five states from `discordBlockState(me)` (4.1) and renders it inside one `StatePanel`, except `free`, which renders a sentence and no panel.

**Paste form** (payer, nothing saved; and `Use a different webhook`):

| Order | Region | Content | Style |
|---|---|---|---|
| 1 | Panel label row | gold dot, `Discord` | `StatePanel`, as Devices |
| 2 | Intro | `Get a message in a Discord channel when a server sim finishes or a log finishes parsing.` | 13 px `--muted`, up to two lines. Omitted in the replace form. |
| 3 | Field group | `label.label` `Webhook URL`; the input; a 21 px reserved error line; the `HOWTO` helper; the `TRUST` helper | Input: `h-11` (44 px), `border-line-warm bg-raised rounded-control text-text px-3 text-[15px]`, full width. Helpers 13 px `--muted`. |
| 4 | Events | `Post when` label, then the two checkbox rows (3.4) and the note | |
| 5 | Actions | `Save webhook` (`PRIMARY_BUTTON_FIXED`, `px-5`) and, in the replace form only, `Cancel` (text link, 44 px) | The only filled gold element in the panel. |

**Saved** (payer, working):

| Order | Region | Content | Style |
|---|---|---|---|
| 1 | Label row | `Discord` | |
| 2 | Destination (primary) | `Posts to channel 1098765432101234567` | 14 px; `Posts to` `--muted`, `channel` `--text`, the id JetBrains Mono 13 px `--strong`, tabular. Wraps at word boundaries; the id never breaks. |
| 3 | Failure warning (only with 1 to 4 failures) | `The last 3 posts failed. At 5 in a row it switches off. Last error from Discord: Missing Permissions.` | 13 px `--muted` |
| 4 | Events | `Post when`, checkboxes, note | |
| 5 | Actions | `Send test message` and `Remove`, both `SECONDARY_BUTTON_FIXED`, `px-4`, 12 px gap, wrapping | No filled button: the panel is quiet when it works. |
| 6 | Status line | reserved 21 px, 13 px | `--text` for success, `--wipe` for failure; `role="status"`, `role="alert"` on failure |
| 7 | Replace link | `Use a different webhook` | 13 px `--text`, underline, 44 px tall, directly under the status line |

**Switched off** (`disabled_reason` is `failed five times`, plan active): label row; destination; a notice (2 px `--wipe` left bar, 14 px `--strong` sentence `Switched off after 5 failed posts in a row.`, under it 13 px `--muted` `Last error from Discord: Unknown Webhook`, omitted when `last_error` is null); the 13 px helper `Sends a test message. If Discord accepts it, posts resume.`; actions `Test and switch on` (`PRIMARY_BUTTON_FIXED`, the panel's only filled gold element) and `Remove`; status line; `Use a different webhook`; then the events (still editable).

**Plan lapsed** (`entitlements.notifications` is not true and a webhook is saved): label row; destination; a notice (2 px `--gold` left bar, `Posts are paused because your plan ended.`, under it `We kept your webhook. Posts start again when you are back on Premium.`); the `See plans` link; the two checkboxes disabled (read-only, showing what was chosen, no note); `Remove`. No `Send test message`, no `Use a different webhook`: the API refuses both.

**Free** (`entitlements.notifications` is not true and nothing is saved): the sentence `Discord messages tell you when a server sim or a log you uploaded is done. See plans` (13 px `--muted`, min height 44 px, `See plans` underlined, `href="/premium"`, `inline-flex min-h-11 items-center` on phone). No panel, no heading, no form, no lock icon. This is the same decision as the character page's sentence (history spec 3.9).

### 3.3 The input

`<input type="text" name="discord-webhook" id="discord-webhook" inputmode="url" autocomplete="off" autocapitalize="off" autocorrect="off" spellcheck="false" enterkeyhint="done" aria-describedby="discord-webhook-help discord-webhook-error">` inside a `<form>`, so Enter submits. It is a plain text field, not `type="password"` (decision 10): a password type invites the browser's password manager to save the URL, and the member needs to see that a paste landed. The typed value lives only in component state until the call returns. It is never written to `localStorage`, `sessionStorage`, the URL, a log, a test id, an error string or an analytics event, and it is cleared the moment the save succeeds. A failed save keeps the typed text so the member can fix it.

### 3.4 The checkbox rows

The native checkbox in a `label.flex.min-h-11.items-center.gap-3.text-[14px]`, exactly the `anonymize` row. Order is fixed: `A server sim finishes`, then `A log file you upload finishes parsing`. Defaults on a new webhook: both on. Under the second row, the note `Posts every log you upload, including private ones; each message says who can open it.` (13 px `--muted`), so a member who uploads a private report knows it will name that report in the channel. In the saved state a change **saves at once** (`PUT` with `events` only): both boxes are disabled with `aria-busy` while the call is in flight, the status line says `Saved.` for 4 s and then clears. A failed call puts the box back and the status line says `Could not save that. Try again.`

### 3.5 What the panel never shows

The webhook id path segment and the token, in any form: not masked, not truncated, not in `title` or `aria-label`, not in the DOM at all (the API never returns them, so they cannot be). The channel id is not a credential and is shown whole. No `Copy` button (there is nothing to copy). No delivery history, counters, "last posted" time or per-event counts: they are pipeline numbers (tenet 7), and the member learns it works from the channel itself. `failures` is shown only as the 1-to-4 warning sentence.

### 3.6 A second channel later (a design rule; nothing is drawn)

The panel is one **channel section**: destination line, `Post when` checkboxes, actions, status line. Events are saved per channel. When a second channel ships (the native apps' phone notifications, owner's direction 2026-10-24), the panel label changes from `Discord` to `Notifications`; each channel becomes a section with a 15 px 600 `--strong` sub-heading (`Discord`, `Phone`); sections are separated by a 1 px `--line-soft` rule with 20 px padding either side; the Premium row label changes from `Discord notifications` to `Notifications` and counts channels on. The `#discord` anchor stays on the Discord section. `AccountDiscordBlock` takes its channel name, field copy and event list from `lib/notifications/copy.ts` keyed by channel, so the second channel is a second entry and a second component for its own destination (a phone has no URL field), not a redesign. This lane builds only the Discord entry and writes no `Notifications` heading.

### 3.7 The Premium row (`AccountPremiumBlock.svelte`)

The `notifications` row stops being `Ships December` and reads the block's state. It links to `#discord` for every payer state and is muted plain text for `free`. `AccountPremiumBlock` gains one prop, `notifications` (the state key and the count of events on), and the row is live when `entitlements.notifications === true` or a webhook is saved.

| State | Value | Link |
|---|---|---|
| Free, nothing saved | `Off` | no, `--muted` |
| Payer, nothing saved | `Off · set up` | yes |
| Saved, two events on | `On · 2 events` | yes |
| Saved, one event on | `On · 1 event` | yes |
| Saved, no event on | `Off · no events chosen` | yes |
| Switched off | `Off · failed five times` | yes |
| Plan lapsed, webhook kept | `Paused · plan lapsed` | yes |

The value column is about 190 px at 390: `Off · failed five times` (23 characters, 14 px) wraps to two lines at 360 and the underline wraps with it, as the History row's does.

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

An older `/v1/me` with no `notifications` key and `entitlements.notifications` true shows the paste form. A saved webhook with both events off is `saved`; the Premium row says `Off · no events chosen` and the block shows the unchecked boxes.

### 4.2 Interactive elements

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Webhook URL input | `border-line-warm`, placeholder `https://discord.com/api/webhooks/…` in `--muted` | border `--line-warm-strong` | 2 px `--gold` outline, 2 px offset | n/a | while the save is in flight | n/a | placeholder shown | border `--wipe`, message in the 21 px line under it (4.3), `aria-invalid="true"`; typed text stays |
| `Save webhook` (primary) | gold fill, `--bg` text | fill `--gold-hover` | 2 px gold outline, 2 px offset | fill darkens to `#cfa64a` | never disabled for an empty field (a disabled button hides why; an empty press shows the error) | label `Saving…`, disabled, `aria-busy` | n/a | status line under the buttons (4.3); the button returns to rest |
| Checkbox row (each) | native box, 14 px label | none | 2 px gold outline on the box | native | disabled and `--muted` while a save is in flight, and always in `lapsed` | `aria-busy` on the group | n/a | box restored, status `Could not save that. Try again.` |
| `Send test message` | outline `--line-warm`, `--text` | border `--gold`, text `--gold-hover` | 2 px gold outline, 2 px offset | ground `--card-top` | while any call is in flight | label `Sending…`, disabled, `aria-busy` | n/a | status line, `--wipe` (4.3) |
| `Test and switch on` (primary, switched-off only) | gold fill | `--gold-hover` | gold outline | darker | in flight | label `Sending…` | n/a | status line; the notice stays |
| `Remove` | outline | border `--gold` | gold outline | ground | in flight | n/a | n/a | status `Could not remove it. Try again.` |
| Remove confirm | the question `Remove this webhook? Posts stop and we forget the URL.` replaces the actions row, with `Remove` and `Cancel` (both `SECONDARY_BUTTON_FIXED`) | as `Remove` | as above; focus moves to `Cancel` | | `Remove` while in flight | `Removing…` | n/a | status `Could not remove it. Try again.`; the question stays |
| `Use a different webhook` | underlined `--text` link-button, 44 px | `--gold-hover` | gold outline, 2 px offset, radius 2 | `--gold-hover` | in flight | n/a | n/a | n/a |
| `Cancel` (replace form) | as the link | as above | as above | as above | while saving | n/a | n/a | n/a |
| `See plans` | `--text` underlined | `--gold-hover` | gold outline, 2 px offset | `--gold-hover` | n/a | n/a | n/a | n/a |

Touch and keyboard parity: every control is 44 px high on phone (buttons `h-11`, link-buttons `min-h-11`, checkbox rows `min-h-11`, input `h-11`). Tab order in the form: input, `Save webhook`, `Cancel` (replace form only); checkboxes come before the actions in DOM order, so the order is input, the two checkboxes, `Save webhook`. Saved: the two checkboxes, `Send test message`, `Remove`, `Use a different webhook`. Enter in the input submits; Space toggles a focused checkbox; Escape in the replace form acts as `Cancel`; Escape in the remove question acts as `Cancel`. No hover-only behaviour exists, so touch needs nothing extra.

### 4.3 The flows, each state's words

**Save.** Client check first (`discordBlockState` module, `webhookShapeError(text)`): the trimmed text must parse as `https://`, with host exactly `discord.com` or `discordapp.com`, and a path of `/api/webhooks/<number>/<token>` (an optional `v6` to `v10` version segment is accepted); anything else shows the one message below and makes no call. The API repeats the check and is the authority; a 400 or 422 from it shows the same message. On 200 the typed value is cleared, the state becomes `saved`, focus moves to `Send test message`, and the status line says `Saved. Send a test message to see it in your channel.` (`role="status"`). The save does not post anything to Discord.

| Failure | Text (verbatim, in the 21 px line under the input) |
|---|---|
| Empty | `Paste a Discord webhook URL.` |
| Not a Discord webhook URL (wrong host, wrong path, not a URL, or the API said 400 or 422) | `That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/.` |
| 402 | `Discord messages are part of Premium.` followed by the `See plans` link |
| Network or 5xx | `Could not save that. Try again.` |

None of these messages repeats, quotes or hints at what was typed.

**Test.** `Send test message` posts `POST /v1/notifications/discord/test`. It reads `Sending…` and is disabled until the call answers (the API times out at 5 s, so a call ends within about 6 s). 204: status `Sent. Check the channel.` in `--text`. A Discord error: status `Discord said: Unknown Webhook.` in `--wipe`, using the error text the API returns (shortened to 140 characters by the API; the web also truncates at 140 and never renders HTML). 429 from us: `Wait a minute before sending another test.` Network or 5xx: `Could not send the test. Try again.` A manual test does not count toward the five failures. A test that succeeds on a switched-off webhook switches it on, and the block re-renders as `saved` with the status `Sent. Posts are on again.`

**Switched off after five failures.** The server sets `disabled_reason: "failed five times"` and keeps `last_error`. On the next `/v1/me` read the panel is the Switched off state and the Premium row says `Off · failed five times`. There is no live push to an open page; a member with the account page open learns on its next load. The way back is `Test and switch on` (a new webhook URL through `Use a different webhook` also resets the count).

**Failure warning.** With 1 to 4 failures in a row, the saved panel adds the 13 px sentence `The last N posts failed. At 5 in a row it switches off. Last error from Discord: …` (`The last post failed.` for one). It disappears when a post succeeds.

**Remove.** `Remove` replaces the actions row with the question. Confirm calls `DELETE /v1/notifications/discord`; success returns to the paste form with focus on the input and the status `Removed.` shown once under the panel's label row for 4 s (`role="status"`); in `lapsed`, success returns to `free`, which is the sentence, and the focus lands on the Plan panel's Discord row link.

**Replace.** `Use a different webhook` swaps the saved view for the paste form without the intro, with `Cancel`. The saved webhook keeps posting until a new one is saved. `Save webhook` here sends the new URL and the current events; success returns to `saved` with the new channel id.

**Free.** No state: the sentence only.

## 5. Hierarchy

- **One primary element per state.** Paste form: `Save webhook`, the only filled gold element. Saved: the destination line (`Posts to channel ...`); no filled button, because nothing is wrong. Switched off: the notice, and `Test and switch on` as the only filled gold element. Plan lapsed: the notice; `See plans` is a text link, not a button (the upsell grammar). Free: the sentence.
- **One level of emphasis.** The destination (14 px, the id in mono `--strong`) is the answer; everything else is 13 px `--muted` or a control. `--wipe` appears only on the switched-off bar, failed status text and the invalid-field border. `--gold` appears on the panel dot and label (as every `StatePanel`), the primary button, the lapsed bar and focus outlines.
- **Secondary is visibly secondary.** `Use a different webhook` and `Cancel` are text links under the actions; `Remove` is an outline button, never red, never first (the account page already keeps its one red for the guild-leave confirm).
- **Whitespace.** 16 px between regions, 8 px label to field, 44 px rows; the 21 px status line is reserved, so a result never moves the buttons. No dividers inside the panel.
- **Not a pitch.** The free state is one sentence with no panel, heading, lock or colour; a free member never sees a form they cannot use.

## 6. Copy, verbatim

New strings live in `web/src/lib/notifications/copy.ts`; the Premium row and /premium strings in `web/src/lib/billing/copy.ts`. Functions take the data and return the string; counts and the channel id come from `/v1/me`.

### 6.1 On the page

| Where | String |
|---|---|
| Panel label | `Discord` |
| Intro (form) | `Get a message in a Discord channel when a server sim finishes or a log finishes parsing.` |
| Field label | `Webhook URL` |
| Placeholder | `https://discord.com/api/webhooks/…` |
| Helper (how) | `In Discord: Channel settings, Integrations, Webhooks, New Webhook, Copy Webhook URL.` |
| Helper (trust) | `We store the URL encrypted, never show it again, and delete it with your account.` |
| Events label | `Post when` |
| Event 1 | `A server sim finishes` |
| Event 2 | `A log file you upload finishes parsing` |
| Events note | `Posts every log you upload, including private ones; each message says who can open it.` |
| Save | `Save webhook`; pending `Saving…` |
| Cancel | `Cancel` |
| Destination | `Posts to` `channel` `{channel id}` |
| Test | `Send test message`; pending `Sending…` |
| Remove | `Remove`; pending `Removing…` |
| Remove question | `Remove this webhook? Posts stop and we forget the URL.` (buttons `Remove`, `Cancel`) |
| Replace | `Use a different webhook` |
| Saved status | `Saved. Send a test message to see it in your channel.` after a first save; `Saved.` after an event change |
| Test sent | `Sent. Check the channel.`; after switching on `Sent. Posts are on again.` |
| Test failed | `Discord said: {error}.` (the error text, one trailing full stop) |
| Test rate limited | `Wait a minute before sending another test.` |
| Test, our error | `Could not send the test. Try again.` |
| Removed | `Removed.` |
| Remove failed | `Could not remove it. Try again.` |
| Save failed (event change or webhook) | `Could not save that. Try again.` |
| Empty field | `Paste a Discord webhook URL.` |
| Invalid URL | `That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/.` |
| Not entitled on save | `Discord messages are part of Premium.` + `See plans` |
| Warning, one | `The last post failed. At 5 in a row it switches off. Last error from Discord: {last_error}.` |
| Warning, several | `The last {n} posts failed. At 5 in a row it switches off. Last error from Discord: {last_error}.` |
| Switched off | `Switched off after 5 failed posts in a row.` / `Last error from Discord: {last_error}` / `Sends a test message. If Discord accepts it, posts resume.` / button `Test and switch on` |
| Lapsed | `Posts are paused because your plan ended.` / `We kept your webhook. Posts start again when you are back on Premium.` / `See plans` |
| Free | `Discord messages tell you when a server sim or a log you uploaded is done.` + `See plans` |
| Premium row label | `Discord notifications` |
| Premium row values | `Off`, `Off · set up`, `On · 2 events`, `On · 1 event`, `Off · no events chosen`, `Off · failed five times`, `Paused · plan lapsed` |
| /premium line | `Discord notifications when a server sim finishes or a log finishes parsing.` (replaces `Discord notifications (ships December).`; the line is edited when the lane ships, never ahead) |
| Changelog entry | `Added Discord notifications. Premium members can paste a Discord webhook on the account page and get a message when a server sim finishes or a log finishes parsing.` (dated the day it ships, past tense) |

`{error}` and `{last_error}` are Discord's own words for what went wrong (for example `Unknown Webhook`, `Missing Permissions`, `Invalid Webhook Token`). They are shown as plain text, never as HTML, shortened to 140 characters, and the API removes any URL or token from them before they reach the page. No string uses `soon`, an exclamation mark, or `unlock`.

### 6.2 In Discord: the messages

Every message is one POST with one embed, no `content`, `username` `Forever Sixty`, `avatar_url` the Forever Sixty seal, and `allowed_mentions` `{"parse": []}` so a title that contains `@everyone` pings no one. The embed colour is `15055189` (`#e5b955`). The footer text is `Forever Sixty`; Discord adds the time from the embed's `timestamp` (the moment the event finished). The title is the link. Text that members wrote (report and sim titles) has Discord's markdown characters escaped. Counts and durations come from the event's own row.

**Server sim finished** (the `sims` job's `Finish`; only a run in state `done` posts)

| Part | Text | Source |
|---|---|---|
| Title (link to `/sim/{id}`) | `Your Frost Mage sim finished` | `specLabel(spec)` |
| Field `Result` (inline) | `1,204 DPS`, or the API headline for gear and talent runs (`+41 DPS from Vis'kag`) | `sims.headline`, falling back to `{dps} DPS` |
| Field `Run` (inline, omitted when the member gave no title) | `Phase 1 BiS check` | `sims.title`, at most 100 characters |
| Field `Took` (inline) | `3 min 12 s`; `48 s`; `1 h 5 min` | `finished_at` minus `created_at` |
| Thumbnail | the spec's class crest (`/icons/hd/crests/mage.webp`, 80 px) | one crest language; omitted if the crest URL is not absolute |
| Footer | `Forever Sixty` | |

**Log parsed** (the report ingest's `complete`, once per whole-file upload; live logging does not post)

| Part | Text | Source |
|---|---|---|
| Title (link to `/reports/{id}`) | `Tuesday Molten Core finished parsing`; with no title, `Your log finished parsing` | `reports.title`, at most 200 characters |
| Field `Zone` (inline, omitted when empty) | `Molten Core` | `reports.zone` |
| Field `Fights` (inline) | `9 fights · 7 kills`; `1 fight · 1 kill`; `3 fights · no kills` | the report summary's fight and kill counts |
| Field `Visible to` (inline) | `Everyone` (public), `Anyone with the link` (unlisted), `Your guild` (guild), `Only you` (private) | `reports.visibility` |
| Footer | `Forever Sixty` | |

**Test message** (the `Send test message` call)

| Part | Text |
|---|---|
| Title (link to `/account#discord`) | `Test message from Forever Sixty` |
| Description | `Your webhook works. You will get a message here when a server sim finishes or a log file you upload finishes parsing.` Built from the events currently on; with one on, only that clause; with none, `Your webhook works. No events are on yet.` |
| Fields | none |

Limits respected: titles at most 256 characters (ours are at most 200), field values at most 1024 (ours at most 100), three fields at most (Discord allows 25), one embed. One event is one message; the site never batches, repeats or follows up. A member who uploads ten files in a row gets ten messages; that is the member's own action and is stated in decision 16.

## 7. Phone and desktop

**Never collapses, at any width:** the destination line and the channel id; the notice sentence in switched-off and lapsed; the two checkbox labels (they wrap, never truncate); the status line; `Remove`; the input at full width; the free sentence and `See plans`.

**Collapses or wraps:** buttons wrap onto a second row when the panel is narrower than their sum (at 360 the panel's inner width is 292 px: `Send test message` plus `Remove` fit on one row at 280 px; `Test and switch on` plus `Remove` wrap, `Remove` on its own 44 px row). Helpers wrap. The channel id sits on its own line when the line is full. The panel padding is 16 px below 768.

- **360 and 390:** the panel is the rail's width (the page's 18 px gutters, 324 or 354 px). Panel heights: saved at most 420 px, paste form at most 560 px, switched off at most 520 px. The Premium row's value wraps to two lines at 360 for the longest values.
- **768 to 1023:** the right rail drops under the left column (single column below `lg`), the panel is the full content width with 24 x 20 px padding, and nothing stretches the input beyond the panel.
- **1024 to 1279:** the rail is 288 px wide; see 3.1.
- **1440:** 395 px wide, as the mock.
- **2000:** the account page's container is the existing centred 1344 column, so the panel is the same 395 px and the page is centred. The block adds no `max-width` of its own and never widens. Capture at 2000 to confirm it is centred with the rail.
- **Keyboard on phone:** the URL input uses `inputmode="url"`, so the keyboard that opens has a paste key; `enterkeyhint="done"`.

## 8. Rulings this spec makes

1. **The block is one panel, and a free member gets a sentence, not a panel.** Same ruling as the history spec's 3.9: a panel with a heading and one sentence is a heading over one line.
2. **The block sits inside `#plan`, under the Plan panel,** 16 px away. It is a benefit of the plan and the Premium row links to it; the rail's other panels (Devices, You) are the account's own settings, 32 px apart.
3. **Events are two checkboxes saved on change, not a form with a Save.** The URL is saved once and the events are two booleans, so a second Save button would be a button with nothing to say. Alternative in decision 8.
4. **Re-enable is a test.** A switched-off webhook comes back only through a post that Discord accepts, so `Test and switch on` is both the proof and the switch. No separate `Re-enable` route and no way to switch on a webhook that does not work.
5. **The saved view shows the channel id and nothing derived from the token.** The page cannot offer a `Copy`, a `Reveal` or a masked tail, because it never holds the token.
6. **The messages are quiet.** One embed, three fields at most, no `content`, no pings (`allowed_mentions` empty), no follow-ups, no digest. The site tells the member once that a thing is done.
7. **A failed sim does not post** (decision 4). The event is "a server sim finished".
8. **Private uploads post, and say who can open them** (the `Visible to` field and the note under the toggles). The webhook is the member's own channel; hiding the report's existence from their own channel would be the site deciding for them.
9. **A lapsed plan keeps the webhook and says why, in the member's own words,** not the API's reason code. `Remove` stays available so a member can delete a credential at any time.
10. **No history, counters or delivery log** in the block (3.5).

## 9. Performance, polish, and acceptance

**Limits.** The block reads `/v1/me`, which the account page already holds: no request of its own on mount, no skeleton, no `LoadError`. It renders in the same paint as the Plan panel it sits under, so nothing appears after hydration beyond what the page itself reveals. No new font, no icon, no image, no new dependency; the Discord name is text. `Account.svelte` gains one import; `AccountDiscordBlock` is small enough to stay in the account island (no lazy chunk). Lighthouse (`web/lighthouserc.json`): `/account.html` is audited signed out, where the page renders its sign-in state and none of this, so the budgets hold by construction (CLS at most 0.05, TBT at most 100 ms, performance and accessibility at least 0.95, 0.9 where the file says so). Add no signed-in `/account` run to the config.

**No layout shift.** The block is the last panel in the rail, and nothing below it but the footer moves. After `me` resolves, every state's height is a function of `me` alone, known in the first render. The 21 px status line is always reserved, so a test result, `Saved.` or an error never moves the buttons. Changes the member causes (Save turning the form into the saved view, Remove turning it back) change the panel's height once and move the footer; nothing else on the page is below it. A member-caused shift is the only one allowed; the Playwright check below measures it at 390 and 1440.

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
12. The CLS numbers printed by the Playwright check at 390 and 1440.
13. A same-scale side-by-side of the shipped rail with the mock's matching panel at 1440.

## 10. API contract this spec assumes

The API lane builds exactly this; if it differs, change the contract here first.

- **`GET /v1/me`** carries `notifications: {discord: {channel_id, events: {sim_finished, upload_parsed}, disabled_reason, last_error, failures} | null}`. `channel_id` is a string (Discord's channel snowflake). `events` has exactly those two keys. `disabled_reason` is `null`, `"failed five times"` or `"plan lapsed"`. `last_error` is `null` or Discord's error text with any URL or token removed. `failures` is the consecutive failure count, 0 to 5. The block is returned for a lapsed member (it is how the page knows the webhook was kept). `entitlements.notifications` is the gate; the web never decides it. Channels are keyed (`discord` now), so a later channel adds a key and changes no field.
- **`PUT /v1/notifications/discord`** with `{webhook_url?, events}`. `webhook_url` is **required when nothing is saved or when replacing** and **omitted** for an events-only change (the page never has the URL, so it cannot resend it). 200 answers the same block as `/v1/me`'s, never the URL. The API validates the host (exactly `discord.com` or `discordapp.com`, https only) and the path (`/api/webhooks/<number>/<token>`, with an optional `v6` to `v10` segment), strips any query or fragment, and rejects userinfo and ports. A new URL resets `failures` to 0 and clears `disabled_reason` and `last_error`. 400 or 422 for a bad URL (the body names no part of the URL); 402 when `Can(FeatureNotifications)` is false. The web shows its own words and never the response body.
- **`channel_id` provenance (decision 1).** A Discord webhook URL contains the **webhook id** and the token, not the channel id. The channel id comes from Discord's `GET /webhooks/{id}/{token}` response (`channel_id`). The API calls it once on save, with a 5 s timeout. If Discord does not answer or the webhook does not exist, the save fails with 422 and the invalid-URL message, which also proves the webhook exists before it is stored. The webhook id itself is never returned.
- **`DELETE /v1/notifications/discord`** removes the URL and the settings. 204. Allowed for a lapsed member. The URL is also deleted with the account.
- **`POST /v1/notifications/discord/test`** posts the test embed (6.2) and answers 204, or a non-2xx with `{error: "<Discord's error text>"}` (URL and token removed, 140 characters at most). Limited to 5 per minute per member; 429 beyond. A successful test sets `failures` to 0 and clears `disabled_reason` when it was `failed five times`; a failed test changes no counter. 402 for a lapsed member.
- **Counting.** A failed event post (non-2xx other than 429, a timeout, a DNS or TLS failure) increments `failures` and stores `last_error`; a success resets `failures` to 0. A Discord 429 waits for the next event and does not count. At the fifth consecutive failure the API sets `disabled_reason: "failed five times"` and stops posting. A lapsed plan reads `disabled_reason: "plan lapsed"` while the URL stays; no post is attempted.
- **Posts** come from the job that produced the event, never a request handler, with a 5 s timeout and no retry beyond the next event. The payload is 6.2. The sim event is the `sims` job's `Finish` (state `done` only); the upload event is the report ingest's `complete`. Both carry the title, visibility, counts and duration the embeds need, so the poster makes no further read.
- **Gating** of the account block: `me.entitlements.notifications`.

## 11. Decisions the product spec left open (for the controller to rule)

1. **`channel_id` is not in the URL.** The product spec says the page shows "the channel id segment" of the URL. The path holds the webhook id and the token; the channel id needs one call to Discord (section 10). Recommended: the API fetches it on save, which also proves the webhook exists. Alternative: show `Webhook {id}` using the first path segment, which is not secret but names nothing a member recognises.
2. **Placement:** the right rail inside `#plan`, under the Plan panel (3.1). Alternative: the left column under Characters, where there is room (the rail ends lower than the left column at 1440). Recommended: under Plan, so the Premium row and the block read as one console.
3. **Free member: a bare sentence, no panel** (3.2, ruling 1). Alternative: a panel holding the sentence. Recommended: the sentence.
4. **A failed sim does not post.** The promise is "a server sim finished". A failed run is already on `/sim` and in the band's list. Alternative: post a failure embed under the same toggle (`Your Frost Mage sim failed`, field `Reason`), which tells a member who closed the tab that the run will never finish. Recommended: no, until a payer asks.
5. **Private uploads post** with `Visible to: Only you` and a note under the toggles (6.2, ruling 8). Alternative: only public, unlisted and guild reports post.
6. **No mentions, ever** (`allowed_mentions: {parse: []}`) and escaped member text, so a title with `@everyone` does nothing. Not optional; listed so the API lane builds it.
7. **The panel is `Discord` now and becomes `Notifications` when the second channel ships** (3.6). Alternative: label it `Notifications` now with a `Discord` sub-heading, which is a heading over one section today.
8. **Events save on change** (3.4). Alternative: one `Save` for URL and events together, which needs a visible dirty state.
9. **Re-enable is `Test and switch on`** (ruling 4), with no extra route. Alternative: a `Switch on` button that clears the flag without a test, which lets a dead webhook burn five more posts.
10. **The URL field is plain text, not masked** (3.3). Alternative: `type="password"`, which hides it from a shoulder surfer or a screen share but invites the password manager to save it and makes a bad paste harder to see.
11. **Accepted hosts:** exactly `discord.com` and `discordapp.com`, https, per the product spec. `ptb.discord.com` and `canary.discord.com` are rejected. Alternative: accept those two subdomains, which Discord's own Copy Webhook URL button may give members on those clients; the helper and the invalid-URL message stay unchanged either way.
12. **Failure counting:** a Discord 429 and a manual test do not count; every other failed event post does (section 10). The product spec's "fails five times in a row" does not say.
13. **Webhook name on the card.** The saved view could show the webhook's own name (Discord's webhook object has one) above the channel id, which a member recognises more easily than a snowflake. It needs the name stored from the same call as the channel id and rendered as escaped text. Recommended for a later pass; this spec draws the channel id the product spec asked for.
14. **Crest thumbnail and our own avatar on the messages.** The sim message carries the spec's class crest as its 80 px thumbnail (one crest language), and every message carries the Forever Sixty seal as the sender's avatar and `Forever Sixty` as its name, overriding the webhook's own. Needs absolute https URLs for both images; if the crest cannot be reached Discord omits it and the message is still whole.
15. **The `Took` field** (finish minus start) is an addition to the title, result and link the product spec named; it is the one number a member did not already know. Drop it if the controller wants two fields.
16. **No coalescing.** A member who uploads ten log files in a row gets ten messages. Alternative: after three messages in ten minutes, post one `and 7 more finished parsing`. Recommended: not in this lane; revisit if upload habits make it noisy.
17. **The trust sentence** (`We store the URL encrypted, never show it again, and delete it with your account.`) is true only if the API stores the URL encrypted (it must be recoverable to post). The API lane confirms, or the sentence loses `encrypted`.
18. **The Premium row's values** (3.7) are new wording for the state; `Off` and `On · n events` follow the server-sims row's `On · n runs this month`.
19. **No live update.** A switch-off that happens while the account page is open shows on the next load. Alternative: re-read `/v1/me` on tab focus, which the page does not do for anything else.
20. **Removing `shipsMonth.notifications`** from `billing/copy.ts` and the tests that reference `Ships December`; the /premium line and the changelog entry change on the day the lane ships, not before.
