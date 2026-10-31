# Discord bot: sign-ups and slash commands in Discord (Lane G5): experience spec

Author: ux-designer, 2026-10-27. Product design (binding): `docs/superpowers/specs/2026-10-27-discord-bot-design.md` §1–8, §8 "Decisions from the builder check" above all. Surfaces it extends: `design/specs/2026-10-26-guild-raids.md` (G1: §3.8 guild settings · Discord, §6.2 embed grammar, the sign-up rules R1 and R3), `design/specs/2026-10-25-discord-notifications.md` (the account block: the connect round trip §4.3, status lines, the 10 s stuck watchdog, ruling 16 no Discord branding on the site), `design/specs/2026-10-27-guild-attendance.md` (G2: the night line, `report_ready`, R-A1, R-A3). Tenets: `docs/tenets.md` 1, 2, 7, 9, 10, 11, 14.

Mock: `design/mocks/gen_discord_bot.py`. It imports G1's generator, so tokens, SIGNUPS and the Discord chrome cannot drift. It writes seven boards into the ignored `design/mocks/renders/`: `discord-bot-settings.html` and `discord-bot-settings-phone.html`, `discord-bot-account.html` and `discord-bot-account-phone.html`, `discord-bot-post.html`, `discord-bot-replies.html`, `discord-bot-commands.html`. Screenshot them with `render.py` run from `web/`, as its docstring says. The `/bis` boards read the active build's BiS file and item table (icon, quality, item level), so names, sources and DPS deltas are the nightly's; everything else is labelled mock.

No new token, no new site control (tenet 10). The Discord side uses only Discord's own components (embeds, buttons, string selects); its look is Discord's, its words are ours.

## 0. What this spec amends

- G1 §3.8: the Discord section gains the bot as its first path; the webhook connect becomes the alternative one link away (§3.1).
- G1 §6.2: the raid embed's `Signed up` field becomes a full-width two-line counts field; a bot post is edited for life, so "the message is never edited" holds only for webhook posts (§3.3.1). The `raid_changed` time and cancel embeds become states of the same message when a bot post exists (product §8.1).
- Account page: the `You` panel gains the Discord link row (§3.2). The Premium `Discord` notifications block is untouched; the two never share a panel.

## 1. The one question, and the answer on the first screen

| Viewer, surface | Question | Answer visible with no scroll, 360 x 640 and 1280 x 800 |
|---|---|---|
| Officer, guild settings `#discord` | "How do I get sign-up buttons into our Discord?" | `Add the Forever Sixty bot`, the section's one gold fill, and one sentence under it saying what it does. The `#discord` anchor lands the h2 at the top; the button ends above y 130 inside the section at both widths. |
| Officer, settings after install | "Where does it post, and is it working?" | `Posts raids to #raid-signups in Olympus Discord` with `With sign-up buttons` under it: the first line under the h2. |
| Member, account page `#you` | "Is my Discord linked?" | `Linked to Discord as zulmara`, or the `Link Discord` button with one sentence. Inside the `You` panel; the `#you` anchor lands it at the top. |
| Raider, a raid post in Discord | "When, how full, and am I in?" | The embed: title (zone and the guild's date), `When` in the reader's own time, `Signed up 23 of 40 · 3 tanks · 6 healers · 12 DPS · 2 late`, and the four buttons. On the phone client the whole post, buttons included, fits one 640 px screen (the mock's open post is about 490 px tall at 354). The raider's own state is never in a public post; one press answers it ephemerally. |
| Raider, after a press | "Did it take?" | One ephemeral sentence: `Zulmara is In as DPS for Onyxia's Lair, Thu 10 Dec. 23 of 40 signed up.` |
| Anyone, a slash command | The command's question | The answer embed is the first thing in the reply; item answers lead with the item's icon and quality bar (tenet 2). |

Tenet 14: every surface leads with the player's character or raid. The bot has no welcome message, no feature pitch and no unprompted post (product anti-goal).

## 2. References, and what we borrow

- Raid-Helper's Discord sign-up post, the tool guilds use today. Borrowed: one public post per raid that updates live; sign-up from buttons on the post; counts by role in the post; ephemeral confirmations. Not borrowed: the emoji role wall, per-class columns of names in the public post (the roster lives on the raid page), reactions as input. Ours beats it on one axis Raid-Helper cannot match: a press is the site's sign-up row, with the per-account rule, the guild's verified characters and the attendance that follows (G2).
- The in-game Calendar invite (G1 §2): one press per state, Accept the one emphasised choice. Borrowed: In is the only coloured button (Discord's Success green); Late, Bench and Out are Discord's grey.
- Wowhead's item tooltip and BiS table: an item is its icon, its quality colour and its source line. Discord draws no tooltip and no coloured text, so each item is its own embed: thumbnail the item's own icon (80 px), the embed's left bar its quality colour, the quality word in the text (`Epic`), then slot, item level and the source line in the game's words (tenet 7). The title links to the item on our site, where the shared tooltip opens.
- Discord's own Connections screen (User Settings): an external account is one row, its name, and a remove action. Borrowed for the account page: one row, the Discord display name, `Unlink`.
- Our webhook connect (G1 §3.8, account §4.3): the round trip, `Opening Discord…`, the 10 s stuck line, the `pageshow` reset, return flags dropped with `history.replaceState`, status lines that stay until the next action. Every one is reused unchanged.

## 3. Layout, region by region

Site tokens (all in `web/src/styles/tokens.css`): `--color-raised #0d111a`, `--color-card-top #131824`, `--color-line #262e40`, `--color-line-soft #1c2230`, `--color-line-warm #3a3326`, `--color-line-warm-strong #4a4030`, `--color-gold #e5b955`, `--color-gold-hover #f5d27a`, `--color-strong #f2eee4`, `--color-text #e9e4d8`, `--color-muted #9a9484`, `--color-kill #7fd48a`, `--color-wipe #f0736b`, the `--color-rarity-*` and `--color-class-*` sets. Type: `section-title` 18 px h2; Barlow 13–14 px; JetBrains Mono for ids; `label` 11 px uppercase. Radii `rounded-control` 4 px, `rounded-panel` 6 px.

Shared site parts reused: `PRIMARY_BUTTON_FIXED` and `SECONDARY_BUTTON_FIXED` (`lib/planner/styles.ts`); `FilterSelect` (`components/rankings/FilterSelect.svelte`) for the channel select, 44 px, warm border, its own chevron; `GuildPlanGate`; `StatePanel` (the `You` panel as it is); `account/discord/DiscordConnectLink` (it already takes `href`), `DiscordStatusLine`, `DiscordActions` (with G1's `gateId` prop), `DiscordNotice`; the in-place confirm pattern (a question and two buttons replace the actions row, focus on `Cancel`, Escape acts as `Cancel`); the 21 px reserved status line. New files, no new look: `guild/GuildDiscordBot.svelte` (installed and removed states), `guild/GuildDiscordChannelPick.svelte` (the picker), `account/AccountDiscordLink.svelte`. Copy lives in `web/src/lib/guild/discord-bot-copy.ts` and `web/src/lib/account/discord-link-copy.ts`; the Discord strings live in the API's `notify` copy structs, as `raidCopy` does today.

### 3.1 Guild settings · Discord (amends G1 §3.8)

Position, markup, the flat-section grammar (no `StatePanel`), the officer-only presence (`settings.discord` present), ruling G1 (every verified officer may act) and ruling G2 (the gate is `CanForGuild(officer_views)`) are G1 §3.8's, unchanged. Mock: `discord-bot-settings.html` and `discord-bot-settings-phone.html`.

The section reads `settings.discord` as today plus `settings.discord.bot` (§9 contract). One state at a time, chosen in this precedence:

| # | State | When |
|---|---|---|
| 1 | Picker | `bot` present and `bot.channel_id` null (installed, no channel yet), or `Change channel` pressed |
| 2 | Installed | `bot` present with a channel, `bot.removed_in_discord` false |
| 3 | Removed in Discord | `bot.removed_in_discord` true |
| 4 | Webhook only | `bot` null, a webhook connected: G1's connected or switched-off state, plus the bot offer |
| 5 | Nothing connected | `bot` null, no webhook |

When the deployment has the bot off (`settings.discord.bot_available` false), states 1 to 3 never occur, state 4 has no bot offer, and state 5 is G1's not-connected state exactly as shipped.

State 5, nothing connected. Top to bottom:
1. h2 `Discord`.
2. `Add the Forever Sixty bot`: `PRIMARY_BUTTON_FIXED px-5`, an `<a>` to `GET /v1/guilds/{id}/discord/bot/install` on the API origin, same tab, as wide as its label (about 248 px). The section's one fill.
3. 8 px gap, the explanation, 13 px `--muted`, the full section width (G1's amended rule).
4. The reserved 21 px status line.
5. 12 px gap, `Connect a channel without the bot`: a 13 px `--text` underlined link-button, `min-h-11`, an `<a>` to the existing webhook connect route (`guildDiscordConnectUrl`). It is the webhook path one link away, as `Paste a webhook URL instead` is on the account block. Under it, no gap, its note, 13 px `--muted`, one line at 1440, two at 360.

Height cap: 250 at 1440, 340 at 360; with a three-line refusal in the status line, 300 and 410.

State 4, webhook only. G1's connected (or switched-off) state, unchanged, with one block inserted directly after the destination lines and 12 px above the actions row: the sentence `Add the bot and raid posts get In, Late, Bench and Out buttons. This channel stays as the fallback.` (14 px `--text`), 8 px, then `Add the Forever Sixty bot` as `SECONDARY_BUTTON_FIXED px-4`. It is not a fill: the connected state is quiet (G1 §3.8, one fill per state), and switched off keeps `Test and switch on` as its one fill. Why here: tenet 14, the next step that improves the guild sits under the answer and above the maintenance actions. Cap: G1's caps plus 100 px.

State 1, picker. Reached on the install return (`?discord_bot=installed`) and from `Change channel`.
1. h2.
2. A 14 px line: `Bot added to` in `--muted` on the first pick after the install (outside it, from `Change channel` or the can't-be-used notice, `Raid channel in`; controller's ruling 2026-10-27), the server name in `--strong` 600 (escaped, at most 60 characters, then `…`).
3. 12 px gap, the field group, 8 px label to field: `label` `Raid channel`; a `FilterSelect` (44 px, full width up to 360 px). Its options are the server's text channels in which the bot holds View Channel, Send Messages, Embed Links and Read Message History, grouped by Discord category as `<optgroup>` in Discord's order, uncategorised channels first, each option `#channel-name`. Preselected: the current channel on `Change channel`; else the first channel whose name contains `raid` (case-insensitive); else none, with the placeholder `Choose a channel`. Under the select, 13 px `--muted`: `Text channels the bot can post in.`
4. 12 px gap, the actions: `Save channel` (`PRIMARY_BUTTON_FIXED px-4`, the one fill), then `Remove the bot` on the install return, or `Cancel` on `Change channel` (both 13 px underlined link-buttons, 44 px tall).
5. The status line.

On the install return the status line reads `Bot added. Choose the channel it posts raids in.` (`role="status"`, stays until the next action) and focus moves to the select. The channel list is fetched when the picker opens (`GET …/discord/bot/channels`). Its states: loading, the select drawn `disabled` with the one option `Loading channels…`, `aria-busy`, and `Save channel` disabled, at the ready height so nothing shifts; empty, the select and helper replaced by the 14 px empty sentence and `Reload the list` (`SECONDARY_BUTTON_FIXED px-4`); error, `Could not load the channels.` (13 px `--wipe`, `role="alert"`) under the label and `Try again` (`SECONDARY_BUTTON_FIXED px-4`) in place of `Save channel`. A successful save shows the installed state with `Saved. Raids post here with sign-up buttons.` (stays) and focus on `Send test message`. Cap: 330 at 1440, 400 at 360; with the open list it overlays, it never pushes.

State 2, installed.
1. h2.
2. The destination, 14 px: `Posts raids to` in `--muted`, `#raid-signups` in `--strong` 600, `in` in `--muted`, the server name in `--strong` 600. Line two, 13 px `--muted`: `With sign-up buttons · Open in Discord`. `Open in Discord` is underlined `--text`, has a 44 px hit area by a transparent pseudo-element (so its focus outline, radius 2 at 2 px offset, rings the text alone), opens the channel URL on discord.com in a new tab with `rel="noopener"`. The bot can read channel and server names, so unlike the webhook it names them.
3. The failure warning, from `bot.failures` 1 on, 13 px: line one `--text`, the count and the channel; line two `--muted`, the cause in words a GM can act on (a permission refusal names the three permissions and the channel; no answer from Discord says the post is retried), never Discord's raw error. The bot never switches itself off, so the webhook's `At 5 in a row it switches off` clause does not apply (2026-10-27 persona fix).
4. The linked-members line, 13 px `--muted`: `9 of 31 members have linked Discord. The others get a link when they press a button.` Both figures count verified members of this guild (ruling §8.6).
5. 12 px gap, the actions: `Send test message`, `Change channel`, `Remove`, each `SECONDARY_BUTTON_FIXED px-4`, 12 px apart, wrapping. No fill: the section is quiet when it works.
6. The status line.
7. When a webhook is also connected, 13 px `--muted`: `Your webhook channel stays as the fallback if the bot is removed.` The webhook's own controls are not drawn while the bot is installed: it is retired for raid posts (product §4.1).

Cap: 330 at 1440, 450 at 360; with the warning, 380 and 520.

State 3, removed in Discord (someone removed the bot in Discord; the API learns it from Discord's guild-delete event or a 403 on the next post). A notice with a 2 px `--gold` left bar, the account block's paused grammar: line one 14 px `--strong` `The bot is no longer in Olympus Discord.`; line two 13 px `--muted` `It was removed in Discord.`, then ` Raids post through your webhook channel again.` or ` Raids stop posting to Discord until the bot is added again.` Then `Add the Forever Sixty bot` (the one fill) and the status line. Cap: 230 at 1440, 300 at 360.

Remove question (from state 2, or `Remove the bot` in the picker). It replaces the actions row: `Remove the bot from Olympus Discord? Raid posts lose their buttons and the commands stop.`, then ` Raids post through your webhook channel again.` or ` Raids stop posting to Discord.`; buttons `Remove` (`Removing…`) and `Cancel`, both `SECONDARY_BUTTON_FIXED`; focus to `Cancel`; Escape acts as `Cancel`. On success the API makes the bot leave the server; the section becomes state 4 or 5 with `Removed.` (4 s) and focus on the add control. Failure: `Could not remove it. Try again.` and the question stays.

Officer-read-only (plan lapsed, `officerTools === false`), G1 §3.8's rules extended. The gate sentence `p#officer-tools-gate-discord` sits under the h2. `Add the Forever Sixty bot` and `Connect a channel without the bot` render as `<a>` with no `href`, `aria-disabled="true"`, `opacity-50`, out of the tab order, described by the sentence. `Send test message`, `Change channel` and `Save channel` are native `disabled`. `Remove` stays live. Installed adds, above the actions, 13 px `--muted`: `Members can still sign up with the buttons on posted raids and use the commands.` (product §8.6). Frozen claim: every control disabled, `Remove` included. A 403 `guild_plan_required` mid-action follows G1's handling: `Not saved. This guild's officer tools have ended.`, focused.

Return from Discord. The install callback answers 302 to the settings page with `?discord_bot=installed`, `denied`, `failed` or `taken` and the `#discord` hash. The flag has its own name, so it never collides with the webhook's `discord=`. It is read once on mount, dropped with `history.replaceState` keeping `#discord`, and the section scrolls to the top: G1's return code path with a second outcome table.

| Flag | Section after return | Status line (verbatim, §6.1) | Role | Focus |
|---|---|---|---|---|
| `installed` | the picker; the settings read is stale, so the section holds the picker height (330 / 400) until it answers and state 5 never flashes | `Bot added. Choose the channel it posts raids in.` | status | the channel select |
| `denied` | unchanged | `Cancelled in Discord. Nothing was added.` | status | the add control used |
| `failed` | unchanged | `Could not add the bot. You need Manage Server in that Discord server and to be an officer of this guild.` | alert | the add control used |
| `taken` | unchanged | `That Discord server is linked to Mograine Raiders. One server works with one guild: ask their officers to remove the bot from their guild settings, or add it to another server.` | alert | the add control used |

The guild named in the `taken` line comes from the settings read (`bot_refusal.guild_name`, §9), never from the URL. With no refusal in the read (it expired), the line reads `That Discord server is linked to another guild. One server works with one guild: ask its officers to remove the bot from their guild settings, or add it to another server.`

Picker races. `Save channel` answering 422 `channel_unusable` (the channel was deleted, or the bot lost a permission since the list loaded): the list re-fetches, the select clears to `Choose a channel`, and the status line says `That channel is gone, or the bot cannot post there. Choose another; the list is fresh.` (`--wipe`, alert, focused). Two officers installing at once: the second return finds the bot installed and reads as `installed` with the saved channel preselected. A second install from a stale tab while a server is bound: the API answers `taken` naming this same guild.

Test message. `POST …/discord/bot/test` posts G1's guild test embed through the bot to the raid channel. Status lines are G1's verbatim (`Sent. Check the channel.` and the rest). A test that succeeds clears `failures`.

### 3.2 Account page · the Discord link (`account/AccountDiscordLink.svelte`)

Mount: inside the `You` `StatePanel`, after the pseudonym row and its note, as the panel's second region (the panel's `gap-4`). The panel gains `id="you"` and `scroll-margin-top: 80px`. Drawn when `me.discord_link_available` is true; absent otherwise. Why here: a Discord link is an identity setting like the pseudonym and free for every account, so it does not belong under Plan (the Premium console) beside the notifications block. Mock: `discord-bot-account.html` and `discord-bot-account-phone.html`.

Unlinked, top to bottom, `flex flex-col items-start gap-2`:
1. The lead. Default, 13 px `--muted`: `Link your Discord account so your guild's sign-up buttons and Forever Sixty commands know your characters.` When a guild of the account has the bot installed (`me.discord_link_suggested` names it), 14 px `--text`: `Olympus XXVII signs up in Discord. Link your Discord account to use the buttons.`
2. `Link Discord`: an `<a>` to `GET /v1/discord/link/start` on the API origin, same tab. `SECONDARY_BUTTON_FIXED px-4` by default; `PRIMARY_BUTTON_FIXED px-5` when suggested, where it is the panel's one fill (the `You` panel has no other).
3. The 21 px status line; under it, with the conflict line only, a 13 px `--muted` help line.

Linked: a 14 px line, `Linked to Discord as` in `--muted`, the display name in `--strong` 600 (escaped, at most 32 characters, `overflow-wrap: anywhere`); then `Unlink` (`SECONDARY_BUTTON_FIXED px-4`); then the status line. `Unlink` asks in place, under the linked line (which stays): `Unlink Discord? Sign-up buttons and commands will ask you to link again.`, with `Unlink` (`Unlinking…`) and `Cancel`, both `SECONDARY_BUTTON_FIXED px-4`; focus to `Cancel`, focus-visible; Escape acts as `Cancel`. Success: the unlinked state, `Unlinked.` for 4 s, focus on `Link Discord`. Failure: `Could not unlink. Try again.` and the question stays.

Return: the callback answers 302 to `/account?discord_link=linked#you` (or `denied`, `failed`, `conflict`), read, dropped and scrolled as the webhook return is (account §4.3 step 4).

| Flag | After | Status line | Role | Focus |
|---|---|---|---|---|
| `linked` | linked, held at the linked height until the fresh `/v1/me` | `Linked. In Discord, sign-up buttons and commands now know your characters.` (stays until the next action) | status | `Unlink` |
| `denied` | unchanged | `Cancelled in Discord. Nothing was linked.` | status | `Link Discord` |
| `failed` | unchanged | `Could not link. Try again.` | alert | `Link Discord` |
| `conflict` | unchanged | `That Discord account is linked to another Forever Sixty account.` in `--wipe`, then the help line `Sign in to that account and unlink it on its account page, then try again.` | alert | `Link Discord` |

`Opening Discord…`, `aria-busy` and pointer events off run from the click until unload; the `pageshow` reset and the 10 s stuck line `Discord did not open. Try again.` are the webhook connect's. A signed-out visitor following a link from Discord: the start route without a session answers 302 to sign-in with the return set to the start route, then on to Discord (product §8.4); they land on `/account#you` with the flag. It is not a login method: nothing on the sign-in page mentions Discord.

Loading reserve (account spec §9.1): the `You` floor grows by the link row. The account block in `Base.astro` adds `data-discord-link` with one of `none`, `unlinked`, `suggested`, `linked` from the cached `/v1/me`, and `layout.ts` adds the measured row height per value. Targets at 1440 and 390: unlinked 115 and 135, suggested 120 and 140, linked 115 and 120 (the line, `Unlink` and the status line stacked; 2026-10-27 persona fix). The warm-CLS gates of §9.1 apply unchanged.

### 3.3 In Discord

Reference for every frame: the boards `discord-bot-post.html`, `discord-bot-replies.html` and `discord-bot-commands.html`, each message drawn in the desktop client (600 px message, embed at most 520) beside the phone client (354 px, embed 270).

#### 3.3.0 Grammar and Discord's limits

G1 §6.2's grammar holds for every message the bot sends: `username` and avatar are the application's own (`Forever Sixty`, the seal; the Developer Portal application is named exactly that, notifications ruling 16); `allowed_mentions` `{"parse": []}`; every member-written string (raid titles, character names, notes) has Discord's markdown escaped; titles link to the site with absolute URLs only; the footer of the last embed in a message is `{Guild} · Forever Sixty` on guild answers and `Forever Sixty` on `/char`, `/bis` and `/help`, with the embed `timestamp` beside it. Only the last embed of a multi-embed message carries a footer.

Colours, as embed colour integers: gold `15055189` (`#e5b955`) for an open raid and neutral answers; muted `10130564` (`#9a9484`) for a locked or ended raid; wipe `15758187` (`#f0736b`) for cancelled; Done green `8377482` (`#7fd48a`) for a done raid and reports; the class colour (`--color-class-*`) for a character card and an attendance record; the quality colour (`--color-rarity-*`, the game's own values, not the lightened text variants) for an item.

Images: an item embed's thumbnail is the item's own icon at the URL the site serves it (`dataUrl(build, "icons/{icon}.webp")`, absolute), never a reinterpretation (owner rule: faithful HD icons); a character card's thumbnail is the ringed HD crest (`/icons/hd/crests/{class}.webp`, absolute; one crest language); a raid post's thumbnail is the guild crest when `crest_url` is absolute, else none. If Discord cannot fetch an image the embed is still whole.

Limits this spec stays inside, and the builder's tests assert: at most 10 embeds per message (the largest answer, `/report`, is 9); at most 6000 characters across a message's embeds; title 256; description 4096; at most 25 fields, name 256, value 1024 (the counts field is under 120); footer 2048; at most 5 action rows of at most 5 components (the raid post uses one row of 5); button label 80 (ours at most 21); string select at most 25 options, option label and description 100 each, placeholder 150; message content 2000; slash command description 100 and option description 100 (§6.2 strings are all under 90); custom ids 100.

Buttons use Discord's styles, never a colour of ours: Success (3) for In, Secondary (2) for Late, Bench, Out and `Share in channel`, Link (5) for every link to the site. No Primary (blurple) and no Danger (red) anywhere: Out is a normal answer, not an alarm (G2 R-A3's spirit). No emoji on any button (G1 §2: the emoji role wall is not borrowed).

#### 3.3.1 The raid post (one message per raid, edited for life)

One embed and one action row. Mock: `discord-bot-post.html`, seven states.

| Part | Open (planned) | Locked | Moved | Cancelled | Ended, no log | Done |
|---|---|---|---|---|---|---|
| Title (links to the raid page) | `{Title} · Thursday 10 December` (the guild's date) | same | `Moved: {Title} · Friday 11 December` until the raid starts, then plain | `Cancelled: {Title} · Thursday 10 December` | plain | `Done: {Title} · Thursday 10 December` |
| Colour | gold | muted | gold | wipe | muted | Done green |
| Description | `Sign up with the buttons below.` | `Sign-ups are locked. An officer can still change yours.` | `Sign-ups carry over. Check you can still make it.` | `This raid was cancelled.` | `This raid has ended.` | `Logged as a raid night.` |
| Fields | `When` (inline), `Guild time` (inline, only when the guild has a zone), `Signed up` (full width) | same | `When` (new), `Was` (old stamp), `Guild time`, `Signed up` | `When` struck through, `Signed up` | `When` with a relative stamp, `Guild time`, `Signed up` | `When`, `Came` `22 of 23`, `Kills` (one line per kill, `Onyxia 1:42`, at most 6, then `+{n} more`) |
| Row | `In` `Late` `Bench` `Out` live, `Raid page` link | the four disabled, the link live | live | the link only | the four disabled, the link live | `Open the night` link only, to the night |

`When` is the Discord timestamp markup `<t:{unix}:F> (<t:{unix}:R>)`, so each reader sees their own time. `Guild time` is `8:00 pm CST`. The `Signed up` field, full width, two lines:
- line one: `23 of 40 · 3 tanks · 6 healers · 12 DPS · 2 late`, the G1 count grammar verbatim (signed up counts In plus Late; the roles count In rows; the late clause only when late is above 0; `1 tank`, `1 healer`; unknown size reads `23 signed up` with no `of`); the first clause in bold;
- line two: `Bench 3 · Out 3`, each clause only when above 0; the line is absent when both are 0;
- with nobody signed up: `None yet`.

No names in the public post: who is coming is on the raid page for members (R1 keeps the schedule off the open web; the channel is the officer's choice, product §8.7, but the roster stays one tap away rather than in a channel any guest may read).

Edits: every change that the raid page would show (a sign-up on either surface, lock, unlock, edit, cancel, the end time passing, a log matched by G2) marks the post dirty; the per-raid worker edits at most once per 2 s and the last state wins (product §8.9). The end-time edit is scheduled at `starts_at + duration`. A webhook post is the same builder with no row, the description `Sign up on Forever Sixty.` and the counts at posting (product §8.11); it is never edited.

#### 3.3.2 A press on the post (and `/signup`)

Every answer is ephemeral (flag 64), one sentence, at most one link button. Mock: `discord-bot-replies.html`.

Flow, in this order (product §8.5): resolve the Discord user to an account; resolve the server to the guild; the raid's state; the account's verified characters in the guild; then the site's sign-up service with the per-account limit.
1. Not linked: `Link your Discord account to Forever Sixty to sign up from here, then press the button again.` with the link button `Link Discord` (the start route; it carries no Discord token and no raid id).
2. Raid cancelled, ended or locked: its G1 sentence (`This raid was cancelled.`, `This raid has ended.`, `Sign-ups are locked. An officer can still change yours.`). A lock that lands between the post's last edit and the press reads the same, and the post's next edit disables the buttons.
3. No character of the account in the guild: `Sign-ups are for members of Olympus XXVII. Save your character on Forever Sixty and an officer can approve it.` with `Get set up` (link to `/setup`).
4. Only pending characters: `Duskstrider is waiting for approval. You can sign it up once an officer approves it.` (G1's waiting sentence with the name; with two or more pending, `Your characters are waiting for approval. You can sign them up once an officer approves them.`).
5. One verified character: sign it up at once, with the role of its current spec (`SPEC_ROLE`).
6. Two or more verified characters: the select. Text `Sign up as In for Onyxia's Lair, <t:{unix}:f>, with which character?` (the pressed state in bold); a string select, placeholder `Choose a character`, one option per verified character, the account's main character first (`main_character_key`) then the game's sort; option label the name, description `Marksmanship Hunter · DPS`, with ` · your main` on the main. Over 25 characters: the first 25 in that order and the text gains `Showing 25 of {n}. Sign up the others on the raid page.` A choice replaces this same ephemeral message (interaction type 7) with the outcome.
7. The outcome, from the service:

| Outcome | Sentence |
|---|---|
| In or Bench or Out saved | `Zulmara is In as DPS for Onyxia's Lair, Thu 10 Dec. 23 of 40 signed up.` (the raid's day, `Thu 10 Dec`, in the guild's raid zone, so a guild with two raids of one title a week can tell which the press landed on; controller's ruling 2026-10-27) (`is on the Bench`, `is Out` for those states; the count clause only for In and Late) |
| Late saved | `Zulmara is Late as DPS for Onyxia's Lair, Thu 10 Dec. Add a note for the officers on the raid page.` with the link button `Add a note` (the raid page, `#signup`) |
| R3 move | `Zulmara is In as DPS for Onyxia's Lair, Thu 10 Dec. Your other character Duskstrider moved to Bench.` (names it as the presser's own character, so it never reads as another raider bumped; controller's ruling 2026-10-27) |
| Same state pressed again | `Zulmara is already In as DPS for Onyxia's Lair, Thu 10 Dec.` (no toggle, product §8.5) |
| Rate limit (30 a minute per account) | `Too many sign-up changes. Wait a minute, then try again.` |
| 409 `raid_locked` or `raid_cancelled` from the service | the step 2 sentence |
| Any other failure | `That did not save. Try again.` |

Timing: a press is acknowledged within 3 s. If the outcome is ready by 2.5 s the reply is the outcome; otherwise the reply is deferred ephemerally (type 5, flag 64), Discord draws `Forever Sixty is thinking…` for the presser only, and the outcome edits it in. The post itself is never edited by the press's reply; it follows through the coalesced edit.

`/signup status:<In|Late|Bench|Out> [raid] [character]` runs the same flow: `raid` defaults to the next planned raid (autocomplete lists the guild's planned raids as `Onyxia's Lair · Thu 10 Dec`, soonest first, at most 25), `character` skips the select (autocomplete lists the account's verified characters in the guild). With no planned raid: `No raid planned yet.`

#### 3.3.3 Slash commands

Registered globally at deploy (product §8.10), usable in servers only (contexts: guild; integration type: guild install). Every answer is ephemeral; `Share in channel` appears only where §3.3.4 allows. Commands resolve the guild from the server; in a server with no bound guild every guild command answers `This server is not linked to a Forever Sixty guild. An officer can add the bot from the guild's settings.` A viewer who is not linked sees what a signed-out visitor sees (product §4.4). Each command acknowledges within 3 s; any read not answered by 2 s defers (Discord draws `Forever Sixty is thinking…` ephemerally) and the answer follows within the interaction's 15 minutes, past which it is dropped and logged (product §8.9). Mock: `discord-bot-commands.html`.

Item embed (the one grammar for every item in `/bis`, `/char` and `/report`): title the item name, linking to the item on the site (the BiS page's band anchor for `/bis`, which has no slot anchors (persona round 2, 2026-10-27), the character page's slot for `/char`, the report's loot row for `/report`); colour its quality colour; thumbnail its icon; description line one `Head · Epic · item level 61` (slot only where the answer is about a slot), line two the delta first when there is one, then the source exactly as the site's BiS page renders it (its source-cell rule over the build's loot.json: `−15.0 DPS · PvP rank 14 · Lieutenant Commander · Alliance`; `Vendor: Pix Xizzix`; `Stormpike Guard (exalted)`; `Quest: Divino-matic Rod · Level 47`; `Crafted: Blacksmithing`; the page's kind word when loot.json has nothing for the item; persona round 2, 2026-10-27), or for a report award `to Vexlash · Onyxia`. The quality word is there because a colour bar alone fails colour-blind readers and the Discord phone client narrows the bar to 4 px. DPS figures follow the SimC convention, one decimal, the minus sign U+2212.

| Command | Options (name, type, required) | Answer |
|---|---|---|
| `/raid` | none | The raid post's embed for the guild's next planned raid, with the description replaced by the viewer's state: `Zulmara is In as DPS.` / `Zulmara is Late as DPS.` / `Zulmara is on the Bench.` / `Zulmara is Out.` / `You have not signed up yet.` (persona round 2, 2026-10-27); then the post's row (the buttons act as on the post). No Share. Unlinked: `Link your Discord account to see your guild's next raid.` with `Link Discord`. None planned: `No raid planned yet.` |
| `/signup` | `status` (choice: In, Late, Bench, Out; required), `raid` (autocomplete), `character` (autocomplete) | §3.3.2. No Share. |
| `/char` | `name` (string, required; autocomplete from the guild's public characters) | One card embed, then up to three item embeds. Card: title the name (link to the character page), class colour, the ringed crest as thumbnail, description `Level 60 Marksmanship Hunter · Olympus XXVII` (the guild only when the character has one), inline fields `Item level` `61.4` (the mean of the worn items' item levels), then, for a verified member of the guild only and never in a share, `Readiness` (the board's word: `Stocked`; `Short: ` and the first two failing checks in the board's own phrases, `Short: 51 unspent talent points`; `Not shared`) and `Synced` (`2 days ago`), then the full-width field `Top 3 upgrades · level 60 Phase 1 raid` (the band the gains are against) with `By DPS gained, against the guide build.` when the spec is not a tank and the band `/bis` defaults to for this character (its level band, the Phase 1 raid at 60; the character's faction) measured a gain of its pick over the worn item for at least one slot: the three largest of those gains, so an item shows the same gain here as in `/bis` (persona round 2, 2026-10-27); otherwise `By Forever Score gained, against the guide build.`, the gear panel's three largest gaps in points against its band (`Top 3 upgrades · level 60`). The three item embeds follow, largest first, each with its gain in that one unit (`+9.2 DPS` or `+7.0 points`) and source. No gear shared: the card alone with `Top 3 upgrades` reading `No gear shared yet. Upgrades show once the character is synced.` Row: `Share in channel`, `Character page` link. Resolves the name within the guild's region and ruleset. Private or unknown: `No public character by that name.` |
| `/bis` | `class` (choice of the nine), `spec` (autocomplete filtered by class), `slot` (choice of 17: Head, Neck, Shoulder, Back, Chest, Wrist, Hands, Waist, Legs, Feet, Finger 1, Finger 2, Trinket 1, Trinket 2, Main hand, Off hand, Ranged), `phase` (choice: Level 20, Level 30, Level 40, Level 50, Level 60, Phase 1 raid); all optional | With a slot: the text line `Head for a Fury Warrior, level 60 Phase 1 raid:`, the pick's item embed with the author line `Best in slot`, then its first two alternatives as item embeds with their DPS deltas (tenet 7: a pick shows its alternatives). Without a slot: the text line `Fury Warrior, level 60 Phase 1 raid. The five slots where the best item gains the most over the next best:` and five item embeds (ruling §8.4). Row: `Share in channel`, then `Open on Forever Sixty` (with a slot) or `All 17 slots` (without), both links to the BiS page. Class and spec omitted: the linked viewer's main character's class and spec, and its level's band; unlinked: `Name a class and spec, or link Discord to use your character.` with `Link Discord`. A spec given without a class resolves from the spec. Phase omitted: the main character's band when the class matches it, else `Phase 1 raid`. Faction: the guild's. |
| `/report` | `raid` (autocomplete of the guild's last 10 done raids) | The last done raid. Main embed: title `Report: Onyxia's Lair · Thursday 10 December` (link to the raid page `#report`), Done green, description the G2 night line (`22 of 23 came · 1 kill · 3 wipes · 14 deaths · 2 h 41 min`), full-width fields `Kill order` (`1. Onyxia 1:42 after 3 wipes`, one line per named boss, G2 §3.4 words), and for officers `Most deaths` (the top three, `Pyrewisp 4 · Windtalon 3 · Vexlash 2`) and `Consumables` (`3 short before the raid · 2 with no flask in the log`), then `Loot` (`6 items:` when every item follows, else `10 items. The first 8:`; persona round 2, 2026-10-27). Then one item embed per award, most valuable quality first, at most 8 (`to {Name} · {Boss}`). Row: `Share in channel` (only when §3.3.4 allows), `Open the report` link. None: `No raid logged yet. A raid shows here once its log is matched.` Unlinked: `Link your Discord account to see your guild's raids.` with `Link Discord`. |
| `/attendance` | `name` (string; autocomplete from the guild's roster, officers only) | One embed: title `Zulmara · attendance` (link to the roster), class colour, description `Came 7 of 8 · 26 of 30 raids` (G2's roster chip verbatim; report-mode guilds `5 of 8 nights`, the web chip verbatim; fewer than 8 raids `Came 3 of 3 raids`), full-width field `Last 5 raids`, one line each, `Thu 10 Dec Present`, `Thu 3 Dec Late, excused`, `Sat 28 Nov Not in the log` (G2 outcome words; never "no-show"). With two or more own characters, one embed per character, main first, at most 5. Row: `Roster` link. No Share. A member giving `name`: `Only officers can look up another member's attendance.` Unknown name: `No member of Olympus XXVII by that name.` Unlinked: `Link your Discord account to see your attendance.` with `Link Discord`. No raids: `No raids yet.` |
| `/readiness` | none | Officers: one embed, title `Readiness · Onyxia's Lair · Thursday 10 December` (link to the Readiness tab with the next-raid filter), gold, description `23 signed In or Late. 5 have something to fix.` (or `Everyone signed In or Late is ready.`), full-width field `To fix, most checks first`, one line per failing row, at most 10, `Pyrewisp · Short · Synced 9 days before the raid` (the board's own words), then `+{n} more on the board` when over 10. Row: `Readiness board` link. No Share. A member: `Readiness is for officers of Olympus XXVII.` Plan lapsed: `Planning and running raids are part of the guild plan. An officer can subscribe the guild.` with `Subscribe the guild` (to the settings `#billing`). Frozen: `This guild's claim is contested. Officer actions are frozen until a moderator resolves it.` No raid planned: `No raid planned yet.` Unlinked: `Readiness is for officers of Olympus XXVII. Link your Discord account if you are one.` with `Link Discord`. Rank, plan and freeze are checked on every call (product §8.5). On purpose, `/readiness` is stricter than the site's board, which still shows a lapsed or Guild-tier officer a read-only board: an officer-only command answers only with the plan that carries the board, Raid guild (launch pricing §3.2.8), else the plan sentences in §6.2 (Task 14 note, 2026-10-27). |
| `/link` | none | Unlinked: `Link this Discord account to Forever Sixty so sign-up buttons and commands know your characters.` with `Link Discord`. Linked: `This Discord account is linked to Forever Sixty. Unlink it on your account page.` with `Account page` (to `/account#you`). |
| `/help` | none | One gold embed, title `Forever Sixty commands`, description one line per command, the command in bold then its line (§6.2), footer `Forever Sixty`. No Share, no button. |

Any command's own failure: `Could not load that. Try again.`

#### 3.3.4 Share in channel

`Share in channel` (Secondary) posts the answer again as a public follow-up, recomputed as a signed-out visitor would see it (product §8.7). Discord shows it under `{user} used /{command}`, so the share needs no "shared by" text of ours. After a share the ephemeral's button turns disabled with the label `Shared`; one share per answer. Present on: `/char`, `/bis`, and `/report` only when the matched log's visibility is `Everyone` (a visitor can open that report page; the shared answer is the kill order and the night line, with no loot names, deaths or consumables, and the link goes to the public report, not the raid page). Absent on: `/raid` and `/signup` (a visitor sees no planned raid, R1), `/attendance` (another viewer's record is not public), `/readiness` (officer-only), `/link`, `/help`. A share whose recomputed answer is empty (the character turned private since) answers ephemerally `That is no longer public, so it was not shared.`

## 4. States

### 4.1 Site controls

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| `Add the Forever Sixty bot` (fill in states 3 and 5; outline in state 4) | gold fill, `--bg` text / outline `--line-warm` | `--gold-hover` fill / border `--gold`, text `--gold-hover` | 2 px gold outline, 2 px offset | `#cfa64a` / `--card-top` | read-only and frozen: `aria-disabled`, no `href`, `opacity-50`, out of tab order | `Opening Discord…`, `aria-busy`, pointer events off until unload; `pageshow` reset; 10 s stuck line | n/a | arrives as a return flag (§3.1) |
| `Connect a channel without the bot` | 13 px `--text` underlined, 44 px | `--gold-hover` | gold outline, 2 px offset, radius 2 | `--gold-hover` | as above | `Opening Discord…` | n/a | the webhook's own return lines (G1) |
| Channel select (`FilterSelect`) | warm border, value or placeholder `Choose a channel` | border `--line-warm-strong` | gold ring | native | while saving; read-only | `Loading channels…`, `disabled`, `aria-busy` | the empty sentence and `Reload the list` replace it | `Could not load the channels.` and `Try again` |
| `Save channel` | gold fill | `--gold-hover` | gold outline | `#cfa64a` | never disabled for an empty choice (a disabled button hides why); a press with none shows `Choose a channel.` under the select (`--wipe`, `aria-invalid`). Disabled only while loading, saving and read-only | `Saving…` | n/a | `That did not save. Try again.`; the race line (§3.1) |
| `Send test message`, `Change channel`, `Remove`, `Reload the list`, `Try again` | outline `--line-warm`, `--text` | border `--gold`, text `--gold-hover` | gold outline | `--card-top` | read-only (all but `Remove`); frozen (all) | `Sending…`, `Removing…` | n/a | the status line, `--wipe` |
| Remove question | question plus `Remove`, `Cancel`; focus on `Cancel`; Escape = `Cancel` | as buttons | as buttons | as buttons | `Remove` while in flight | `Removing…` | n/a | `Could not remove it. Try again.`; the question stays |
| `Open in Discord` | underlined `--text` | `--gold-hover` | gold outline | `--gold-hover` | n/a | n/a | n/a | n/a |
| `Link Discord` | outline (fill when suggested) | as the button kind | gold outline, 2 px offset | as the button kind | never (the row is absent when unavailable) | `Opening Discord…`; `pageshow` reset; stuck line | n/a | return flags (§3.2) |
| `Unlink` and its question | outline | border `--gold` | gold outline | `--card-top` | while in flight | `Unlinking…` | n/a | `Could not unlink. Try again.` |

Touch and keyboard parity: every site control is 44 px tall at every width; nothing is hover-only. Tab order in the section: state 5, add, the webhook link; picker, select, `Save channel`, `Remove the bot` or `Cancel`; installed, `Open in Discord`, `Send test message`, `Change channel`, `Remove`. Account: `Link Discord`, or `Unlink` then the question's `Unlink`, `Cancel`.

### 4.2 Discord components

Discord draws hover, focus and press itself, on desktop and phone. We control: label, style, `disabled`, and what each press answers.

| Component | Live | Disabled | Pressed, before the answer | Answer |
|---|---|---|---|---|
| `In`, `Late`, `Bench`, `Out` on the post | locked: no; open and moved: yes | locked and ended: Discord's disabled look (50 %), presses not delivered | Discord's own spinner on the button for under 3 s, or the ephemeral `Forever Sixty is thinking…` when deferred | §3.3.2 |
| `In`… on a `/raid` answer | same as the post | same | same | same |
| Character select | until the ephemeral is dismissed; a choice after the raid locked answers the locked line | n/a | Discord's spinner | the outcome replaces the message |
| `Share in channel` | once | after a share, label `Shared`, disabled | spinner | a public follow-up |
| Link buttons | always | never | n/a (opens the browser) | the site |
| An old bot post after `Remove` | buttons stay drawn, Discord stops delivering presses (product §8.6) | n/a | Discord shows its own "This interaction failed" | none: a removed bot cannot answer |

## 5. Hierarchy

- Settings, state 5: `Add the Forever Sixty bot` is the one fill; the webhook path is a text link, visibly secondary. Picker: `Save channel` is the one fill. Installed: no fill; the destination line is the answer. Removed in Discord: the notice, then the add button as the one fill.
- Account: no fill unless a guild of the account uses the bot; then `Link Discord` is the panel's one fill.
- Raid post: the title and the `Signed up` line are the answer; `In` is the one coloured button. Bench and Out are grey, never red.
- Command answers: the first embed is the answer; buttons are below it, the link last. Item answers: the icon and the quality bar lead, the name is the link, the source is one line.
- `--wipe` on the site only for refusals, failures and races. In Discord only the cancelled post's bar is red.

## 6. Copy, verbatim

Site strings live in `web/src/lib/guild/discord-bot-copy.ts` and `web/src/lib/account/discord-link-copy.ts`; Discord strings in the API's `notify` copy structs (`botCopy`, beside `raidCopy`). Functions take the data; no number is typed into prose. `{Guild}`, `{Server}`, `{Name}` and the rest are escaped data. No string uses `soon`, an exclamation mark or `unlock`.

### 6.1 On the site

| Where | String |
|---|---|
| Add button | `Add the Forever Sixty bot`; pending `Opening Discord…` |
| Add explanation | `Raid posts get In, Late, Bench and Out buttons, and members can ask /raid, /bis and /char in the server. Discord asks which server; you need Manage Server in it.` |
| Webhook path | `Connect a channel without the bot` / `Raids post there, and members sign up on Forever Sixty.` |
| Bot offer (state 4) | `Add the bot and raid posts get In, Late, Bench and Out buttons. This channel stays as the fallback.` |
| Picker | `Bot added to` `{Server}` (the first pick after the install), else `Raid channel in` `{Server}`; label `Raid channel`; placeholder `Choose a channel`; helper `Text channels the bot can post in.`; `Save channel` (`Saving…`); `Remove the bot`; `Cancel`; `Choose a channel.` (empty press) |
| Picker list states | `Loading channels…`; `The bot cannot post in any text channel of {Server}. In Discord, give the Forever Sixty role View Channel, Send Messages and Embed Links in your raid channel, then reload the list.` + `Reload the list`; `Could not load the channels.` + `Try again` |
| Picker race | `That channel is gone, or the bot cannot post there. Choose another; the list is fresh.` |
| Installed | `Posts raids to` `#{channel}` `in` `{Server}` / `With sign-up buttons` · `Open in Discord`; `Send test message` (`Sending…`), `Change channel`, `Remove` (`Removing…`); `Saved. Raids post here with sign-up buttons.` |
| Linked members | `{linked} of {members} members have linked Discord. The others get a link when they press a button.` (one member: `1 of 31 members has linked Discord.`; none: `No member has linked Discord yet. Members get a link when they press a button.`) |
| Fallback note | `Your webhook channel stays as the fallback if the bot is removed.` |
| Failure warning | `The last post to #{channel} failed. Send a test message to check.` (`The last {n} posts to #{channel} failed.`); cause `Discord refused it for a missing permission. In Discord, give the Forever Sixty role View Channel, Send Messages and Embed Links in #{channel}.` (a 403, or Discord code 50001 or 50013) or `Discord did not answer, so the post is tried again on its own.` (a 5xx or no answer); any other refusal has no cause line |
| Removed in Discord | `The bot is no longer in {Server}.` / `It was removed in Discord.` + ` Raids post through your webhook channel again.` or ` Raids stop posting to Discord until the bot is added again.` |
| Remove question | `Remove the bot from {Server}? Raid posts lose their buttons and the commands stop.` + ` Raids post through your webhook channel again.` or ` Raids stop posting to Discord.` (`Remove`, `Cancel`); `Removed.`; `Could not remove it. Try again.` |
| Read-only | G1 gate sentence verbatim (`Planning and running raids are part of the guild plan. Subscribe the guild.`, link to `#billing`); `Members can still sign up with the buttons on posted raids and use the commands.`; `Not saved. This guild's officer tools have ended.` |
| Install returns | `Bot added. Choose the channel it posts raids in.`; `Cancelled in Discord. Nothing was added.`; `Could not add the bot. You need Manage Server in that Discord server and to be an officer of this guild.`; `That Discord server is linked to {Guild}. One server works with one guild: ask their officers to remove the bot from their guild settings, or add it to another server.`; expired refusal `That Discord server is linked to another guild. One server works with one guild: ask its officers to remove the bot from their guild settings, or add it to another server.`; stuck `Discord did not open. Try again.` |
| Test | G1 §6.1 `Settings · Discord test` row, verbatim |
| Account lead | `Link your Discord account so your guild's sign-up buttons and Forever Sixty commands know your characters.`; suggested `{Guild} signs up in Discord. Link your Discord account to use the buttons.` |
| Account controls | `Link Discord` (`Opening Discord…`); `Linked to Discord as` `{display name}`; `Unlink` (`Unlinking…`); `Unlink Discord? Sign-up buttons and commands will ask you to link again.` (`Unlink`, `Cancel`); `Unlinked.`; `Could not unlink. Try again.` |
| Account returns | `Linked. In Discord, sign-up buttons and commands now know your characters.`; `Cancelled in Discord. Nothing was linked.`; `Could not link. Try again.`; `That Discord account is linked to another Forever Sixty account.` + `Sign in to that account and unlink it on its account page, then try again.`; stuck `Discord did not open. Try again.` |
| Changelog | `Added the Forever Sixty Discord bot. Raid posts carry In, Late, Bench and Out buttons, and /raid, /bis, /char and more answer in your guild's Discord. Officers add it from guild settings; members link Discord on the account page.` (past tense, dated the day it ships) |

### 6.2 In Discord

| Where | String |
|---|---|
| Post descriptions | open `Sign up with the buttons below.`; locked `Sign-ups are locked. An officer can still change yours.`; moved `Sign-ups carry over. Check you can still make it.`; cancelled `This raid was cancelled.`; ended `This raid has ended.`; done `Logged as a raid night.`; webhook `Sign up on Forever Sixty.` |
| Post titles | `{Title} · {Weekday} {d} {Month}`; `Moved: …`; `Cancelled: …`; `Done: …` |
| Post fields | `When`, `Was`, `Guild time`, `Signed up`, `Came`, `Kills`; values `{signed} of {size}` plus the G1 role clauses plus ` · {l} late`, line two `Bench {b}` and `Out {o}` joined by ` · `; `None yet`; `{n} signed up`; `+{n} more` |
| Post buttons | `In`, `Late`, `Bench`, `Out`, `Raid page`, `Open the night` |
| Press: not linked | `Link your Discord account to Forever Sixty to sign up from here, then press the button again.` + `Link Discord` |
| Press: refusals | `Sign-ups are locked. An officer can still change yours.`; `This raid was cancelled.`; `This raid has ended.`; `Sign-ups are for members of {Guild}. Save your character on Forever Sixty and an officer can approve it.` + `Get set up`; `{Name} is waiting for approval. You can sign it up once an officer approves it.`; `Your characters are waiting for approval. You can sign them up once an officer approves them.`; `Too many sign-up changes. Wait a minute, then try again.`; `That did not save. Try again.` |
| Press: select | `Sign up as {State} for {Title}, <t:{unix}:f>, with which character?`; placeholder `Choose a character`; option description `{Spec} {Class} · {Role}`, on the main ` · your main`; `Showing 25 of {n}. Sign up the others on the raid page.` |
| Press: outcomes | `{Name} is In as {Role} for {Title}, {Date}. {signed} of {size} signed up.`; `{Name} is Late as {Role} for {Title}, {Date}. Add a note for the officers on the raid page.` + `Add a note`; `{Name} is on the Bench for {Title}, {Date}.`; `{Name} is Out for {Title}, {Date}.`; ` Your other character {Other} moved to Bench.`; `{Name} is already {State} as {Role} for {Title}, {Date}.` (`is already on the Bench`, `is already Out`); `{Date}` is `Thu 10 Dec` in the guild's raid zone (controller's ruling 2026-10-27) |
| `/raid` | `{Name} is In as {Role}.` / `{Name} is Late as {Role}.` / `{Name} is on the Bench.` / `{Name} is Out.` / `You have not signed up yet.`; `No raid planned yet.`; `Link your Discord account to see your guild's next raid.` When the signed-up character cannot be named (controller's addition, 2026-10-27): `You are In as {Role}.` / `You are Late as {Role}.` / `You are on the Bench.` / `You are Out.` |
| `/char` | `Level {n} {Spec} {Class} · {Guild}`; `Item level`, `Readiness`, `Synced`, `Top 3 upgrades` plus ` · {band label}`; `By DPS gained, against the guide build.` / `By Forever Score gained, against the guide build.`; `No gear shared yet. Upgrades show once the character is synced.`; `Character page` |
| `/bis` | `{Slot} for a {Spec} {Class}, {band label}:`; `{Spec} {Class}, {band label}. The five slots where the best item gains the most over the next best:`; author `Best in slot`; `Open on Forever Sixty`; `All 17 slots`; band labels `level 20`, `level 30`, `level 40`, `level 50`, `level 60`, `level 60 Phase 1 raid` |
| Item embed | `{Slot} · {Quality} · item level {n}` over `{delta} · {source}` ({source} the BiS page's source line, §3.3.3); quality words `Poor`, `Common`, `Uncommon`, `Rare`, `Epic`, `Legendary` |
| `/report` | `Report: {Title} · {Weekday} {d} {Month}`; the G2 night line; `Kill order`, `Most deaths`, `Consumables`, `Loot`; `{n} items:` (every item follows), else `{n} items.` plus ` The first {k}:`; `to {Name} · {Boss}`; `Open the report`; `No raid logged yet. A raid shows here once its log is matched.`; `Link your Discord account to see your guild's raids.` |
| `/attendance` | `{Name} · attendance`; the G2 roster chip verbatim; `Last 5 raids`; `Roster`; `Only officers can look up another member's attendance.`; `No member of {Guild} by that name.`; `No raids yet.`; `Link your Discord account to see your attendance.` A linked viewer with no verified character in the guild (Task 14's addition, 2026-10-27): `Attendance is for members of {Guild}.` |
| `/readiness` | `Readiness · {Title} · {Weekday} {d} {Month}`; `{n} signed In or Late. {k} have something to fix.` (`1 has`; controller's ruling 2026-10-27 dropped the leading "The"); `Everyone signed In or Late is ready.`; `To fix, most checks first`; `+{n} more on the board`; `Readiness board`; `Readiness is for officers of {Guild}.`; `Readiness is for officers of {Guild}. Link your Discord account if you are one.`; `Planning and running raids are part of the guild plan. An officer can subscribe the guild.` plus `Subscribe the guild`; `This guild's claim is contested. Officer actions are frozen until a moderator resolves it.` A Guild-tier guild, whose plan does not carry the Raid guild tools (Task 14's addition, 2026-10-27, after launch pricing §3.2.8): `The readiness board is part of the Raid guild plan. An officer can move the guild to the Raid guild plan.` plus `Move to Raid guild`. Nobody signed In or Late for the next raid (controller's addition, 2026-10-27): `No one has signed In or Late yet.` |
| `/link` | `Link this Discord account to Forever Sixty so sign-up buttons and commands know your characters.` plus `Link Discord`; `This Discord account is linked to Forever Sixty. Unlink it on your account page.` plus `Account page` |
| `/help` title | `Forever Sixty commands` |
| Command descriptions, also the `/help` lines | `/raid` `Your guild's next raid. Sign up from the answer.`; `/signup` `Sign up for a raid without the post.`; `/char` `A character: item level, readiness and the top 3 upgrades.`; `/bis` `Best in slot for a class and spec.`; `/report` `The last logged raid: kills and loot.`; `/attendance` `Your attendance. Officers can name anyone.`; `/readiness` `Officers: who is not ready for the next raid.`; `/link` `Link this Discord account to Forever Sixty.`; `/help` `This list.` |
| Option descriptions | `status` `In, Late, Bench or Out. Leave out to pick on a card.`; `raid` `Which raid. The next one if left out.`; `character` `Which of your characters.`; `name` (`/char`) `A character name. Yours if left out.`; `class` `The class. Leave out to pick it.`; `spec` `The spec. Leave out to pick it.`; `slot` `One slot. The top five if left out.`; `phase` `Level band or Phase 1 raid.`; `name` (`/attendance`) `Officers: a member of the guild.`; `raid` (`/report`) `Which raid. The last one if left out.` (controller's addition: the `raid` string above says "the next one", which is wrong for `/report`) |
| Share | `Share in channel`; after a share `Shared`; `That is no longer public, so it was not shared.` |
| Unbound server | `This server is not linked to a Forever Sixty guild. An officer can add the bot from the guild's settings.` |
| Any failure | `Could not load that. Try again.` |
| Footers | `{Guild} · Forever Sixty`; `Forever Sixty`; on `/bis`, `{Spec} {Class} · {band label} · Forever Sixty` |

## 7. Phone and desktop

Site. Never collapses: the add button and its explanation, the destination line and the server name, the linked-members line, the gate sentence, every status and race line, the remove question, the account row's lead, the display name and the status line. Wraps: the actions row (at 360 the three installed buttons take two rows, `Send test message` and `Change channel` on the first, `Remove` on the second); the long refusal (three lines at 360); the server name (at word boundaries). The select is full width up to 360 px and stays 360 px from 768 up. At 1024 the account rail is 288 px wide: the lead wraps to four lines and the button never wraps its label. At 2000 everything stays in the centred 1344 column; the section keeps its 720 px prose measure and the rail its 395 px.

Discord desktop client (the mock's 600 px message): inline fields run three to a row, so `When` and `Guild time` share a row and `Signed up` spans it; the five post buttons fit one row. Phone client (the mock's 354 px): every field stacks; the four sign-up buttons fill one row (Discord's phone buttons are at least 40 px tall) and `Raid page` drops to a second row; item thumbnails stay 80 px and the text column narrows to about 160 px, so a source line may take two lines and is never cut. Neither client is ours to lay out; the captures prove each state reads in both (§9).

## 8. Rulings this spec makes

1. Bot first, the webhook one link away. One primary per state; the bot is the better tool (buttons, commands); the webhook stays for guilds that will not add a bot (product goal 5).
2. No names in the public post. Counts only; the roster is on the raid page. A guild channel can hold guests, and R1 keeps who is coming for members.
3. One coloured button. In is Discord's Success green, the others grey; no red, no blurple. The palette is Discord's own, so no site branding enters Discord beyond the name and the seal.
4. `/bis` without a slot lists the five slots where the pick beats its next best by the most DPS (the BiS file's first alternative's `dps_delta`), not the five highest-scoring items. The score is in reference-stat points and is missing for weapons and trinkets in the current files, so "top by score" would be wrong for every spec. The lead line says what the five are.
5. R3 is a move, not a refusal. The product lists "one character per account" among the refusals; on the site, In or Late on a second character moves the first to Bench (G1 R3), and the service the click calls does the same (product §8.5). The reply says `Your other character Duskstrider moved to Bench.`
6. The linked-members line in the installed state. Raid leaders lose sign-ups to members who never linked; the count tells an officer whom to nudge. It is a guild fact, not a pipeline counter (tenet 7), and costs one count in the settings read.
7. The account link sits in the `You` panel, not under Plan: it is free and an identity setting.
8. The character select lists the main character first. The product says "header character first"; the header selector's current character is browser state the API cannot see, and its fallback is already the main (`main_character_key`). The option says ` · your main`.
9. `/bis` options are all optional and default to the linked viewer's main character, its level band and the guild's faction (tenet 14: lead with the player's character). The product listed `class` and `spec` without saying they are required.
10. Lock, end and done change the post's buttons (disabled, or replaced by one link) rather than leaving live buttons that refuse every press. Done adds `Came` and `Kills` from G2 so the post closes the loop it opened.
11. Bench and Out are not counted in `Signed up` (G1), but the post shows them on a second line because a raid leader filling a 40-man needs the bench number at a glance (Raid-Helper shows it).

## 9. Performance, contract and acceptance

Limits. The site adds no font, no image and no request on page load: the section and the account row render from the settings read and the `/v1/me` the pages already make; the channel list is fetched only when the picker opens. Lighthouse budgets (`web/lighthouserc.json`, the `/guild/` and `/account` rows) are unchanged: both audits run signed out, where neither surface renders. No layout shift: every status line is reserved; the `installed` and `linked` returns hold their target heights until the fresh read answers; the account reserve grows by the link row (§3.2). Playwright CLS on guild settings and `/account`, warm, at 390 and 1440: at most 0.02 beyond member-caused changes.

Contract this spec assumes (the API lane builds this; change it here first if it differs):
- Settings read: `settings.discord` gains `bot_available` (boolean) and `bot`, which is null or an object with `server_id`, `server_name`, `channel_id` (null until picked), `channel_name`, `linked_members`, `members`, `failures`, `last_error`, `removed_in_discord`. It also gains `bot_refusal`, null or an object with `guild_name`, present for 10 minutes for the officer whose install was refused and cleared on read.
- `GET /v1/guilds/{id}/discord/bot/install` (officer session; rank, plan, freeze) answers 302 to Discord's OAuth2 authorize with scopes `bot applications.commands`, the four permissions, and server-side state carrying the guild and the officer (product §8.3). The callback answers 302 to the settings page with `discord_bot` set to `installed`, `denied`, `failed` or `taken`, and `#discord`.
- `GET /v1/guilds/{id}/discord/bot/channels` answers a `channels` array; each entry has `id`, `name` and `category` (null or the category's name), in Discord's order. `PUT …/discord/bot/channel` with `channel_id` answers the settings `discord` object, or 422 `channel_unusable`. `POST …/discord/bot/test` behaves as the webhook test. `DELETE …/discord/bot` answers 204 and makes the bot leave the server.
- `/v1/me` gains `discord_link_available` (boolean), `discord_link` (null or an object with `display_name`) and `discord_link_suggested` (null or a guild name). `GET /v1/discord/link/start` starts the flow; its callback answers 302 to `/account` with `discord_link` set to `linked`, `denied`, `failed` or `conflict`, and `#you`. `DELETE /v1/discord/link` answers 204.
- Interactions: the endpoint verifies Ed25519 and answers PING (product §3). Custom ids: `raid:` plus the raid id and the state; `pick:` plus the raid id and the state; `share:` plus an opaque token. Nothing in a custom id, link or message identifies another account or carries a credential.

E2E (`guild-discord-bot.spec.ts` and `account-discord-link.spec.ts`, API stubbed):
1. State 5: the add link's `href` is the install route on the API origin; `Connect a channel without the bot` links to the webhook route; with the bot off, G1's not-connected state draws exactly.
2. Each install return flag shows its line, moves focus as §3.1 says and leaves `#discord` with no query; `installed` never draws state 5; `taken` names the guild from the stubbed read, and the expired variant reads without it.
3. Picker: loading, empty, error and ready; the preselect picks the channel containing `raid`; an empty press shows `Choose a channel.`; a save shows the installed state with focus on `Send test message`; a 422 shows the race line, focused, and refetches.
4. Installed: the destination names channel and server; the linked-members line in its three forms; the warning at 3 failures; `Change channel` reopens the picker with the current channel and `Cancel`; the remove question, Escape, and success to state 4 with a webhook and to state 5 without.
5. Webhook only: the bot offer is present as an outline button, and the connected state's controls still work.
6. Removed in Discord: the notice in both variants and the add button.
7. Officer-read-only: add and the webhook link are `aria-disabled` with no `href`; test, change and save are `disabled`; `Remove` works. Frozen disables `Remove` too.
8. Account: unlinked, default and suggested (the fill); each `discord_link` flag; unlink with the question; the conflict and help lines; `Opening Discord…` and the 10 s stuck line under fake timers; the row is absent when unavailable; `#you` scrolls the panel to the top.
9. Axe clean in every state of both surfaces; every control at least 44 px tall at 360; every height inside its §3.1 cap.
10. API unit tests, listed so the lane builds them: each post state's embed and row against §3.3.1 (colour, fields, disabled flags); the counts field's G1 grammar; every limit in §3.3.0; each press outcome and refusal sentence; the select's order and its 25 cap; Share present and absent per §3.3.4, and its recomputation as a visitor; no message or custom id carries a token.

Site captures (fixture build; x = required):

| # | State | 360 | 390 | 1440 | 2000 |
|---|---|---|---|---|---|
| 1 | Settings state 5 under Billing (section order visible) | x | x | x | x |
| 2 | Settings state 4, the bot offer under the webhook destination | x | x | x | |
| 3 | Picker: install return (focus on the select), open list, loading, empty, error, race | x | x | x | |
| 4 | Installed: just saved, the 3-failure warning, the remove question | x | x | x | x |
| 5 | Returns: `taken` (and the expired variant), `denied`, `failed` | x | x | x | |
| 6 | Removed in Discord, with and without a webhook | | x | x | |
| 7 | Officer-read-only state 5 and installed; frozen | x | | x | |
| 8 | Account `You` panel: unlinked, suggested, linked, unlink question, conflict, each return | x | x | x | x |
| 9 | Focus-visible on every site control; hover on the buttons at 1440 | | x | x | |
| 10 | Loading beside ready for settings and `/account`, CLS printed | | x | x | |
| 11 | Same-scale side by side with the mock boards `discord-bot-settings` and `discord-bot-account` at 1440 | | | x | |

Discord captures. The mock boards are the reference until real captures exist. Real captures come from the private test server against the production application (product §8.10), each in the desktop client and the iOS or Android client, beside the matching mock frame:

| # | Message | Desktop client | Phone client |
|---|---|---|---|
| D1 | Raid post: open, locked, moved, cancelled, ended, done; and a webhook post | x | x |
| D2 | The post updating within 2 s of a press on the site and of a press in Discord (a short screen recording) | x | |
| D3 | Press replies: unlinked, select closed and open, In, Late, the Bench move, already, locked, cancelled, ended, pending, not a member, rate limit, thinking | x | x |
| D4 | `/raid` and `/signup` | x | x |
| D5 | `/char` with upgrades, with no gear shared, and unknown | x | x |
| D6 | `/bis` with a slot and without; a shared `/bis` in the channel; the ephemeral after sharing (`Shared`) | x | x |
| D7 | `/report` member and officer; a shared public report | x | x |
| D8 | `/attendance` own (two characters) and officer by name; the member refusal | x | x |
| D9 | `/readiness` officer, member, plan lapsed | x | x |
| D10 | `/link` both states; `/help`; Discord's command picker showing the nine descriptions | x | x |
| D11 | Discord's install dialog showing `Forever Sixty`, the seal and the four permissions; the OAuth `identify` dialog for the link | x | |
| D12 | G1's still-uncaptured rows 17 (posted, moved, cancelled) and 24 (test embed), through the bot | x | x |

A Discord capture that differs from its mock frame in words, colour, button style or order is a defect against this spec; a difference in Discord's own chrome (font rendering, spacing, the APP tag) is not.

## Amendment 2026-10-27: channel can't be used

Source: the controller's ruling of 2026-10-27. The bot is still in the server, but it cannot post in the chosen raid channel: the channel was deleted, or the bot lacks View Channel, Send Messages or Embed Links there. Raid posts fall back to the guild's webhook when one is set; otherwise nothing posts. Saving a channel, the same one or another, clears the state. Reference: this is state 3's notice, one step milder, because the bot is still here and the fix is one pick.

### A.1 Contract

`settings.discord.bot` gains `channel_unusable`: null, or an object `{ reason, missing }`. `reason` is `deleted` or `permissions`. `missing` is the list of lacking permissions, a subset of `View Channel`, `Send Messages`, `Embed Links`, in that order; empty when `deleted`. `channel_name` is null when the API never knew the name. `PUT …/discord/bot/channel` clears `channel_unusable` when it answers 200. The API sets it when a post answers Discord 404 (unknown channel) or 403 on the channel (not on the server, which is state 3).

### A.2 Where it sits

A new row in the §3.1 precedence table, between rows 2 and 3 in weight, evaluated after `removed_in_discord`:

| # | State | When |
|---|---|---|
| 1 | Picker | unchanged |
| 3 | Removed in Discord | unchanged, wins over everything below |
| 2b | Channel can't be used | `bot.channel_unusable` non-null, `removed_in_discord` false, the picker not open |
| 2 | Installed | unchanged |

State 2b, channel can't be used. Replaces state 2's body; the destination line, the linked-members line and the three action buttons are not drawn (a destination that does not work is not a destination). Top to bottom:
1. h2 `Discord`.
2. A notice with a 2 px `--gold` left bar, state 3's grammar exactly (same padding, same bar, no `--wipe`: nothing is lost, the officer has one step to take). Line one, 14 px `--strong`. Line two, 13 px `--muted`. In the permissions case `#raid-signups` is `--strong` 600 within line one.
3. 12 px gap, the actions: `Choose a channel`, `PRIMARY_BUTTON_FIXED px-4`, the section's one fill. In the permissions case, 12 px to its right, wrapping, `Open in Discord`: 13 px `--text` underlined link-button, 44 px hit area, opens the channel on discord.com in a new tab with `rel="noopener"`. In the deleted case there is no secondary action: the channel does not exist.
4. The status line (reserved 21 px).
5. When a webhook is connected, nothing more is drawn; its controls stay hidden as in state 2.

Primary action. `Choose a channel` opens the picker (state 1, `Change channel` variant): the channel list is fetched, the select takes focus, nothing is preselected when `deleted`, and the current channel is preselected when `permissions` and it is still in the list (it is not, unless the officer fixed permissions meanwhile). The cancel link reads `Cancel` and returns to state 2b with focus on `Choose a channel`. Escape acts as `Cancel`. A successful save shows state 2 with `Saved. Raids post here with sign-up buttons.` and focus on `Send test message`, as state 1 does.

Read-only and frozen follow §3.1: `Choose a channel` is native `disabled` when read-only or frozen; `Open in Discord` stays live; the gate sentence sits under the h2 as in the other states.

Caps: 240 at 1440, 320 at 360 (state 3's 230 and 300 plus the extra line of the permissions copy and the secondary link). The picker keeps its own caps (330 and 400).

States of the new controls: `Choose a channel` has the add button's row in §4.1 (hover `--gold-hover` fill, gold focus-visible outline with 2 px offset, active `#cfa64a`, disabled `opacity-50`; loading is the picker's `Loading channels…`). `Open in Discord` has the existing `Open in Discord` row. Touch and keyboard parity: 44 px tall at every width; tab order `Choose a channel`, `Open in Discord`.

### A.3 Copy, verbatim

Where the channel name is known, `{channel}` is `#raid-signups` (escaped, at most 60 characters, then `…`). `{Server}` is the server name. `{list}` joins the missing permissions with commas and `and` (`Send Messages`; `Send Messages and Embed Links`; `View Channel, Send Messages and Embed Links`).

| Case | Line one (14 px `--strong`) | Line two (13 px `--muted`) |
|---|---|---|
| Deleted, name known | `The raid channel {channel} was deleted in {Server}.` | `The bot cannot post there. Choose another channel.` + fallback |
| Deleted, name unknown | `The raid channel in {Server} was deleted.` | `The bot cannot post there. Choose another channel.` + fallback |
| Permissions | `The bot cannot post in {channel}.` | `In Discord, give the Forever Sixty role {list} in that channel, or choose another channel.` + fallback |
| Permissions, name unknown (controller's ruling, 2026-10-27) | `The bot cannot post in the raid channel.` (no `--strong` run) | As the Permissions row |

Fallback, appended to line two with a leading space: ` Raids post through your webhook channel until then.` when a webhook is connected; ` Raids stop posting to Discord until then.` when not.

Controls: `Choose a channel`; `Open in Discord`. The picker, its placeholder `Choose a channel`, `Cancel` and the saved line are §6.1 verbatim. A 422 `channel_unusable` on save (the same channel still unusable, or another that turned bad) reads the §3.1 picker-race line verbatim: `That channel is gone, or the bot cannot post there. Choose another; the list is fresh.`

Add to §6.1: one row, `Channel can't be used`, holding the table above.

### A.4 What the raid list and the raid page show meanwhile

Nothing new. The removed state adds nothing to the raid list or the raid page, and neither does this one. A raid still posts through the webhook when one is set, and the raid page reads as for a webhook-only guild. When nothing posts, the raid page and list read as for a guild with no Discord destination: no error line, no badge. The notice lives in guild settings only, where the officer fixes it. Members are never shown the state. Existing bot posts keep their buttons and keep trying to edit; an edit that fails is not shown.

### A.5 Acceptance

E2E, added to `guild-discord-bot.spec.ts` (API stubbed), numbered after §9's 1 to 10:
11. Channel can't be used, deleted: the notice with the name known and with `channel_name` null; line two carries the webhook fallback with a webhook and the stop sentence without; `Choose a channel` is the only fill; there is no `Open in Discord`; the destination line and the three action buttons are absent.
12. Channel can't be used, permissions: line one names `#raid-signups`; line two lists exactly the stubbed `missing` permissions in the specified order and joining (one, two and three items); `Open in Discord` has the channel URL, `target="_blank"` and `rel="noopener"`.
13. Precedence: `removed_in_discord` true with `channel_unusable` set draws state 3; `channel_unusable` set with the picker open draws the picker; null draws state 2.
14. Fix by picking: `Choose a channel` opens the picker with focus on the select and nothing preselected (deleted); `Cancel` and Escape return to the notice with focus on `Choose a channel`; a save shows state 2 with `Saved. Raids post here with sign-up buttons.` and focus on `Send test message`; the stubbed read after save has `channel_unusable` null.
15. Fix by saving the same channel: with the channel still unusable the 422 shows the picker-race line, focused, and refetches; with it fixed the 200 clears the notice.
16. Officer-read-only and frozen: `Choose a channel` is `disabled`; `Open in Discord` works.
17. Axe clean and 44 px controls at 360 in both causes; heights inside the A.2 caps at 360 and 1440; zero layout shift between the notice and state 2 beyond the cap change.
18. API unit tests: a Discord 404 on post sets `channel_unusable` `deleted`; a 403 on the channel sets `permissions` with the computed `missing`; a 403 on the server sets `removed_in_discord`, not this; a successful `PUT …/channel` clears it; with a webhook the raid posts through it, without one nothing posts and no error reaches the raid page.

Site captures, added to §9: 12. Channel can't be used, deleted with webhook and permissions without webhook, each with the picker open, at 360, 390 and 1440.
