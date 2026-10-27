# Discord bot: sign-ups and slash commands in Discord (Lane G5, promoted)

Status: draft for the owner's approval, 2026-10-27. Promotes Lane G5 of
`docs/superpowers/specs/2026-10-26-guild-suite-design.md` ahead of G3 and G4.

## 1. Intent

Meet players where they are. A guild lives in Discord: raid sign-ups, questions about gear,
"who's coming Thursday". Today the site posts a raid to the guild's channel through a webhook,
and every action needs a trip to the website. A raid leader reviewing the pricing said it
plainly: if raiders must open a website to sign up, "I lose half of them and keep Raid-Helper".

Two jobs:

1. **Sign up from the raid post.** The post carries `In`, `Late`, `Bench`, `Out` buttons and
   live counts by role. A click signs the clicker's character up, exactly as the site's
   one-tap sign-up does, and the post updates.
2. **Ask the site from Discord.** Slash commands answer the questions a guild asks in chat
   (my next raid, my character, BiS for a slot, last night's report) without leaving Discord.

Why now: with millions of Forever players expected, the tool that is already inside guild
Discords on 4 November becomes the default. Lookup commands are also the widest funnel we
have: every answer in a guild channel is the site seen by twenty people.

## 2. Goals, non-goals, anti-goals

Goals
- A guild officer installs the bot in one flow from guild settings and picks the raid channel.
- Raid posts carry sign-up buttons; counts update within seconds of a click on either surface.
- Slash commands for the guild's next raid, sign-up, a character, BiS, the last report,
  attendance and readiness.
- A Discord user links their site account once; after that every button and command knows
  who they are.
- The webhook stays for guilds that do not install the bot.

Non-goals (this lane)
- Raid assignments, soft-reserves and wishlists in Discord (G3, G4 add their commands later).
- Any game-data command that needs Blizzard data we do not have.
- A bot presence in voice, moderation, role management, reminders by DM.

Anti-goals
- A bot that spams: no unprompted messages beyond the raid posts the officer chose.
- Showing in Discord anything the site would not show that viewer (private reports,
  another member's attendance rows, unverified members' data).
- Sign-ups that drift between Discord and the site. There is one sign-up row; both surfaces
  read and write it.

## 3. Constraints

- Discord's HTTP interactions model, no gateway connection: Discord POSTs each button click
  and command to an Interactions Endpoint URL on the API, signed with Ed25519 (verify with the
  application's public key; reject unsigned or stale requests with 401). Every interaction is
  acknowledged within 3 seconds; slow answers defer (type 5/6) and follow up.
- The bot posts and edits messages with a bot token (`DISCORD_BOT_TOKEN`, Secret Manager,
  never logged). Commands are registered globally through the REST API at deploy.
- Discord rate limits: message edits per channel are limited, so post updates are coalesced
  (at most one edit per raid post per 2 seconds; the last state wins).
- Every rule on the site applies unchanged: R1 (visitors and unverified members see no
  planned raid), one In/Late per account (R3), sign-ups closed when locked, officer writes
  behind rank → plan → claim freeze, private reports invisible to others.
- Credentials: bot token, client secret, OAuth codes and Discord access tokens are never in
  logs, responses, the DOM or storage (the connect lane's rule).

## 4. What gets built

### 4.1 Install
- Guild settings' Discord section gains `Add the Forever Sixty bot` beside the existing
  webhook connect. It runs Discord's OAuth2 install with scopes `bot applications.commands`
  and the minimum permissions (View Channel, Send Messages, Embed Links, Read Message
  History), with server-side state like the webhook connect.
- After install the officer picks the raid channel from the server's text channels the bot
  can post in. The guild's existing webhook, if any, is retired for raid posts (kept as a
  fallback if the bot is removed).
- Removing the bot from the server (or `Remove` on the site) reverts to the webhook or to
  nothing.

### 4.2 Account link
- `Link Discord` on the account page (OAuth2 `identify`) stores the Discord user id against
  the site account. `/link` in Discord answers ephemerally with the same one-click link.
- An unlinked user who presses a sign-up button gets an ephemeral reply: what to do, one link,
  nothing else.

### 4.3 Raid posts with buttons
- The raid embed (G1's `raid_posted` and `raid_changed`) gains four buttons and a counts line
  by role (`In 22 · Late 3 · Bench 4` and tanks/healers/DPS).
- A click resolves the clicker's account, then their character for this guild: one verified
  character signs straight away; several bring an ephemeral select (header character first).
  A Late press asks nothing more; a note can be added on the site.
- Refusals answer ephemerally with the site's own sentences (locked, one character per
  account, not verified in this guild, raid cancelled).
- Site-side sign-ups and officer actions (lock, cancel, edit) update the post through the
  same coalesced edit.

### 4.4 Slash commands
Responses are ephemeral by default, with a `Share in channel` button that reposts the answer
publicly. All commands respect the site's visibility rules for the linked viewer; an unlinked
viewer sees what a signed-out visitor sees.

| Command | Answers |
|---|---|
| `/raid` | The guild's next raid: time in the viewer's zone, counts by role, your status, buttons |
| `/signup status:<In/Late/Bench/Out> [raid] [character]` | Signs up without the post |
| `/char name:<character>` | Character card: level, spec, item level, readiness check, top 3 upgrades, link |
| `/bis class:<> spec:<> [slot] [phase]` | BiS for the slot (or the top 5 slots) with sources and icons |
| `/report [raid]` | Last done raid's night line, kill order and loot; officers also deaths and consumables |
| `/attendance [name]` | Your record (`Came 7 of 8`); officers may name anyone in the guild |
| `/readiness` | Officers: the next raid's readiness board summary (failing rows first) |
| `/link` | Link this Discord account to Forever Sixty |
| `/help` | The list above, one line each |

Item and character answers use embeds with the item's icon and quality colour (tenet 2);
class colours and ringed crests where an image appears (tenet 7).

### 4.5 Gating
- Installing the bot and posting raids with buttons: guild plan (both tiers), like the
  webhook channel today.
- Sign-up clicks and member commands: free for every verified member (member writes are free
  so a lapsed guild still fills its raid — guild suite §5).
- Lookup commands (`/char`, `/bis`, `/report` member view): free for everyone. They are the
  funnel.
- A server that installed the bot for a guild whose plan lapses keeps lookups and sign-ups;
  new raid posts fall back to "post from the site" with the read-only sentence.

## 5. Owner steps (cannot be done by the builder)

1. Discord developer portal → the existing Forever Sixty application → **Bot**: create the
   bot user and reset its token; store it as `DISCORD_BOT_TOKEN` in Secret Manager.
2. **General Information**: copy the Public Key into config as `DISCORD_PUBLIC_KEY` (it is
   public, not a secret).
3. After the API deploys the endpoint: set **Interactions Endpoint URL** to
   `https://api.foreversixty.gg/v1/discord/interactions` (Discord verifies it with a PING).
4. **Installation**: add the bot scopes to the install link if the portal asks.

## 6. Order

1. API: interactions endpoint with signature check and PING; the account link; command
   registration at deploy; the bot install flow and channel pick.
2. API: buttons on raid posts, sign-up from a click, coalesced post edits.
3. API: the slash commands, one by one, each answering from the existing read paths.
4. Web: guild settings install and channel pick, the account page link, copy.
5. Captures of the Discord messages on desktop and phone clients in a staging server (the
   G1/G2 row 12 that is still uncaptured gets done here too).

Target: before the first raid night, 9 December; sign-up buttons first, lookups after.

## 7. Left to the builder

Table names (a `discord_links` user↔Discord id table, a `guild_discord_installs` table),
the exact embed layouts within the site's embed grammar, the select-menu wording within §6.1
copy rules, and the command option names beyond those above.

## 8. Decisions from the builder check (2026-10-27, the controller's calls unless marked)

1. **One post per raid, edited for life.** A `guild_raid_discord_posts` row holds the channel
   and message id. The bot posts on `raid_posted` and edits that message on every change
   (lock, cancel, edit, sign-ups); `raid_changed` no longer posts a new message when a bot
   post exists. Raids posted through the webhook before install keep their old posts; the
   next change to such a raid makes the bot post a fresh message with buttons and records it.
2. **One Discord server ↔ one Forever guild**, unique on the Discord server id. A guild has
   one server and one raid channel. Commands resolve the guild from the server. A second
   guild trying to install in a taken server is refused with a sentence naming the guild
   that holds it.
3. **Install binds through the officer's site session.** The install starts from guild
   settings (rank → plan → claim freeze, like the webhook connect), with server-side state
   carrying the guild id and the officer; Discord's install itself proves the installer has
   Manage Server on that server. No extra scope.
4. **Account link.** `Link Discord` uses OAuth2 `identify` on the existing redirect URI with
   its own state kind. One Discord id ↔ one site account, both unique; linking a Discord id
   held by another account is refused ("That Discord account is linked to another Forever
   Sixty account."). `Unlink` on the account page. Only the Discord user id and the display
   name at link time are stored; the access token is discarded. A signed-out user who follows
   the link signs in first and returns to the OAuth step. Not a login method.
5. **A click is the site's sign-up.** The interaction calls the same service function as
   `PUT …/signups/{region}/{ruleset}/{name}`, with the same per-account limit. Only the
   account's verified characters in that guild are offered; pending characters are not.
   The role is the character's current spec role; a different role is set on the site. A
   click on the current state is a no-op that answers ephemerally with the current state
   (no toggle, matching the site's segments). Officer-only commands evaluate rank, plan and
   freeze on every interaction.
6. **Lapsed plan.** Existing posts keep their buttons and keep updating (member writes are
   free; lock and cancel results render). New posts from the bot need the plan; without it the
   raid is posted from the site as today. Removing the bot reactivates the webhook if one
   was connected; old bot posts lose their buttons (Discord stops delivering clicks).
7. **Visibility.** The raid post shows what the webhook post shows today (the officer chose the
   channel). `Share in channel` recomputes the answer as a signed-out visitor would see it and
   is absent on officer-only answers (`/readiness`, `/attendance name:`, officer `/report`
   rows). A lookup never shows a private report or another member's attendance rows.
8. **Lookups read service functions, not HTTP handlers.** Where a read exists only as a
   handler, the lane extracts the service function first (one task per command). `/char`
   answers for any character with a public character page; a private or unknown one answers
   "No public character by that name."
9. **Durable edits.** Pending post edits live in Postgres (`guild_raid_discord_posts.dirty_at`),
   drained by a per-raid worker with a 2 s debounce and retries with backoff; any instance can
   drain. Deferred command answers use the interaction token within its 15 minutes; past that
   the answer is dropped and logged.
10. **Staging.** The controller's default while the owner has not chosen: test against the
    production application in a private Discord test server, with guild-scoped command
    registration to that server for unreleased commands. Config is per environment
    (`DISCORD_APPLICATION_ID`, `DISCORD_PUBLIC_KEY`, `DISCORD_BOT_TOKEN`), so a separate
    staging application can be added later with no code change. Commands are registered per
    application at deploy with a hash check, so an unchanged command set is not re-sent.
11. **Embeds.** One raid embed builder with an optional counts line; webhook posts omit the
    buttons and keep the counts line.
12. **Order.** Must-have before 9 December: install, link, buttons, post edits, `/raid`,
    `/signup`, `/link`, `/help`. Lookups (`/char`, `/bis`, `/report`, `/attendance`,
    `/readiness`) follow and may slip past 9 December.
