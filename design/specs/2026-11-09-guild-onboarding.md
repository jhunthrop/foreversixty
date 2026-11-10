# Guild onboarding and the free Discord bot (2026-11-09)

Owner, verbatim: "we need a guild onboarding setup flow. When someone claims a guild we should walk them through setting everything up for their guild, like adding the discord bot, and setting an invites channel up. etc.. everything they need to do to get the full functionality of the app. We should also make the discord bot free, because they have a competitor that's free already."

Mocks: `design/mocks/ws_guild_setup.py` (`web-guild-setup`, `web-guild-setup-mixed`, `web-guild-claim-done`, `web-guild-settings-setup-row`), renders in `design/mocks/renders/web/`.

## 1. The free Discord bot

### 1.1 Ruling (recorded so the owner can override it)

Free for any claimed guild's verified officer or leader:

- installing the bot and picking its channels;
- the webhook channel (the bot's fallback);
- raid posts with the In, Late, Bench and Out buttons;
- sign-ups and attendance taken from the sign-ups;
- the recruiting channel's application posts;
- member commands: `/raid`, `/signup`, `/bis`, `/char`, `/report`, `/attendance`, `/link` and `/help`.

The raid posts need raids, so the **raid calendar** is free too: create, edit, lock, unlock, cancel and Post now, and an officer answering for a member. The checklist cannot be finished without it.

Stays paid, unchanged:

- **Raid guild:** the readiness board (web and `/readiness`) and the loot council.
- **Guild plan (the officer views):** comp targets, Discord pings for non-signers, boss assignments, log matching, attendance overrides, the raid-night report post, and the coach on guild reports.

No other gate was widened or narrowed.

`/sim` is in the owner's list, but the bot has no `/sim` command. That is a follow-up, not a gate.

### 1.2 Inventory: every bot command and post, before and after

| Surface | Route or code | Before | After |
| --- | --- | --- | --- |
| Install the bot | `GET /v1/guilds/{id}/discord/bot/install` (`discordbot.installStart`) | officer + guild plan (`needPlan` true) | officer, no plan |
| Install return | `installOutcome` | officer + `standing.Plan` | officer |
| List the server's channels | `GET …/discord/bot/channels` | officer + plan | officer |
| Save the raid channel | `PUT …/discord/bot/channel` | officer + plan | officer |
| Save the recruiting channel | `PUT …/discord/bot/recruit-channel` (new) | n/a | officer |
| Test message | `POST …/discord/bot/test` | officer + plan | officer |
| Remove the bot | `DELETE …/discord/bot` | officer | officer |
| Webhook connect / test | `GET …/notifications/discord/connect`, `POST …/test` | officer + plan | officer |
| Webhook remove | `DELETE …/notifications/discord` | officer | officer |
| Raid post (bot or webhook) | `RaidRouter.RaidPosted` from a raid save | the save needed the plan | free with the calendar |
| Raid edit, move, cancel and lock posts | `RaidRouter.RaidChanged` | the write needed the plan | free |
| Day-before reminder | `RunRaidSchedule` | scheduled from a plan-gated save | free |
| Post now / Repost | `POST …/raids/{raid}/discord/post` (new) | n/a | free (calendar gate) |
| Sign-up buttons (In, Late, Bench, Out) | `press.go` → `guilds.SignUp` | own character free; for a member, officer + plan | free |
| `/raid`, `/signup` | `cmd_raid.go`, `cmd_signup.go` | linked member | unchanged |
| `/bis`, `/char`, `/report`, `/attendance`, `/link`, `/help` | `cmd_*.go` | free | unchanged |
| `/readiness` | `cmd_readiness.go` → `guilds.readinessGate` | officer + plan, then Raid guild | officer, then Raid guild. A free guild gets the Raid guild sentence and `Move to Raid guild` (checkout `plan=raid_guild`) |
| Application post, with Accept and Decline | `discordbot.PostApplication`, `applicationPress` (new) | n/a | free; Accept/Decline need a linked verified officer and an unfrozen claim, and run the site's `Decide` |
| Accepted applicant's role and DM | `GrantGuildRole`, `DirectMessage` | officer | unchanged |
| Ping on Discord | `POST …/raids/{raid}/ping` | officer + plan | unchanged |
| Assignments post | `POST …/raids/{raid}/assignments/post` | officer + plan | unchanged |
| Raid-night report post | `POST …/raids/{raid}/report/post` | officer + plan | unchanged |

Code: `notify.GuildOfficerGate` loses `needPlan` and `GuildStanding` loses `Plan`. `guilds.officerRaidGate` takes `needPlan`: the calendar routes go through `calendarRaid` (false) and the raid tools through `officerRaid` (true). `officerGateFor` checks only the rank and the freeze.

### 1.3 Copy changed

- `refusal.RaidPlanRequired` and the apps' `officer.gate`: `Comp targets, Discord pings, assignments and log matching are part of the guild plan.` iOS and Android strings were regenerated.
- The gate's part 1 (`guildGateCopy.lead.raids`): `Comp targets, pings, assignments and log matching are part of the Guild and Raid guild plans.` The Discord sentence is gone, because no Discord gate is left.
- `/premium`:
  - `guildRegion.lede`, `guild.features[0]` and `guild.features[2]`, and `raidGuild.features[4]`, recorded in launch-pricing §6.1.
  - The Guild card no longer sells the bot or raid posts.
  - The guild leaders' lede says what is free.
- Bot `/readiness`: the lapsed-plan sentence and `Subscribe your guild` are removed, recorded in discord-bot spec §6.2.
- Guild settings:
  - the Discord section's read-only state and gate sentence are removed;
  - the bot's `Members can still sign up…` read-only line and the `planEnded` lines are removed.
- Claim page: `Events and the Discord bot are part of the Guild and Raid guild plans.` is removed.
- Raid form: a save posts at once whatever the start date (owner ruling 2026-11-09, hotfix cd60429f9). `form.postLater` is removed, and `form.postRepeat` reads `All four raids post as soon as you save.`

## 2. The setup checklist

### 2.1 Route

- **The checklist:** `/guild/<region>/<ruleset>/<slug>/setup`, a chunk of `GuildShell` beside `/settings` and `/claim`.
- **Claim page:** once you hold the claim, `Set up <Guild>` lists the steps with their states and leads to it (`Start setup`). A claim link made for a plan keeps its subscribe button as the one fill.
- **Settings:** until every step is done or skipped, a Notice heads the page: `Setup · 5 of 8 done` with `Continue setup`.

### 2.2 Steps, in order of value, and how each state is derived

A step is **done** when the guild's data says so. Otherwise it is **skipped** or **done** by an officer's mark (`guild_setup_marks`), or else **to do**. A derived done always wins over a mark.

| # | Step | Done when (derived) | Mark | Controls in place |
| --- | --- | --- | --- | --- |
| 1 | Officer ranks | `officer_max_rank_index` ≠ 1 (the column default) | `done` = keep the current ranks | Each rank's count and names from the roster's `rank_index`; `Officers are ranks 0 to [n]` with Save; `Keep 0 to n` |
| 2 | Discord bot | the bot is installed with a raid channel, or (no bot in the deployment) the webhook is connected | skip only | `Add the Discord bot` (install with `?return=setup`, back to the checklist); a channel picker once installed |
| 3 | Recruiting | recruiting open, a pitch, wanted lines, and a recruiting channel | skip only | Parts: Pitch, Nights, Wanted, Recruit role, Recruit channel; the recruiting channel picker; `Edit the pitch` |
| 4 | Raid schedule | `raid_timezone` set and a planned or locked raid still ahead | skip only | Parts: Time zone, First raid, Comp targets (tagged Guild); `Plan a raid`; with a raid ahead, `Posted to #channel` or `Post now` (any state) |
| 5 | Guild invite link | a link exists and at least one account joined by one | `done` = I posted it | The invite link field; `Paste in guild chat` with the exact line `Join <Guild> on Forever Sixty and sign up for raids: <url>`; `I posted it` |
| 6 | Addon and companion | at least half of the verified accounts have an addon export in the last 14 days | skip only | `<n> of <m> members synced in the last 14 days` (short of half: `· <half> needed`) with a bar, green when done; the paste line `Get the Forever Sixty addon and companion so your gear and the roster sync: https://foreversixty.gg/setup` |
| 7 | Report visibility | `default_visibility` ≠ `public` | `done` = keep Public | Public, Unlisted, Guild only; `Keep Public` |
| 8 | Discord linked role (optional) | never derivable: the role lives in the officer's server | `done` | The Server Settings path to the Forever Sixty Premium requirement; `Done` |

The plan is not a step. It is the rail's Plan panel:

- the current plan (`Free`, `Guild` or `Raid guild`) as its meta;
- three lines: `Free:`, `Guild:` and `Raid guild:`;
- `Compare plans` (`/premium#guild-leaders`).

There is no subscribe button.

Every step except the plan has `Skip`. A marked step has `Undo`. The step's mark carries its state: a check for done, the number in a ring for to do, a dash for skipped. Only `Skipped` is also written. The first step to do carries the page's one primary, and every other step's controls are quiet (ghost). On step 1, `Keep 0 to n` is the action until the number changes; then `Save` is. The checklist reads again after each write and when the window regains focus, so a bot install, a raid planned on the guild page or a member's sync shows without a reload.

### 2.3 The recruiting channel

The officer picks it from the bot's server channels (`PUT …/discord/bot/recruit-channel`; `""` clears it). The picker guesses only a channel whose name reads as recruiting (`recruit`, `apply`, `application`, `join`). It never guesses the raid channel. Each new application then posts there as an embed:

- `Application · <Name>`;
- the spec and class as the author line;
- the message, cut at 600 characters;
- Level, Item level and Available fields;
- `Accept` and `Decline` buttons: a linked verified officer decides in Discord through the site's own `Decide`, so the recruit role, the DM and the votes all follow, and the answer is ephemeral;
- a `Review application` link to the applications page.

The post goes out after the apply has answered, so a failure never fails the application.

## 3. API

| Route | Who | Answer |
| --- | --- | --- |
| `GET /v1/guilds/{id}/setup` | verified officer or leader | the checklist: steps, done, total, finished, plan, and each step's facts |
| `PUT /v1/guilds/{id}/setup/{step}` `{"state":"done"\|"skipped"}` | officer, claim not frozen | the checklist; 400 for an unknown step, or `done` on a step whose done is derived only |
| `DELETE /v1/guilds/{id}/setup/{step}` | officer, claim not frozen | the checklist |
| `PUT /v1/guilds/{id}/discord/bot/recruit-channel` | officer | the settings `discord` object |
| `GET …/discord/bot/install?return=setup` | officer | 302 to Discord; the return lands on `/setup?discord_bot=<outcome>` |
| `POST /v1/guilds/{id}/raids/{raid}/discord/post` | officer (calendar gate) | the raid view; 409 `no_channel` or `conflict`, 502 `post_failed` |

Migration `0094_guild_setup`, for the owner to review before the deploy runs it. Nothing was run against production.

- `guild_setup_marks (guild_id, step, state, set_by, set_at)`: only the marks. Every state is derived.
- `guild_discord_installs.recruit_channel_id` and `recruit_channel_name`.
- The connect-state kind `bot_install_setup`.

## 4. Coordination

Shared changes, all additive:

- `ui/StepMark.svelte` is extracted from `Steps.svelte` and gains the `skip` mark.
- `VISIBILITY_OPTIONS` moves to `lib/guild/settings-view.ts`.
- `GuildShell` gains the `/setup` route.
- `characters.ts` gains `parseGuildSetupPath` and `guildSetupHref`.

The f-guild lane's tabs were edited only where a plan gate came off the free calendar:

- `Guild.svelte` drops `officerTools` from the Raids tab, band line and week-one card mounts;
- `RaidPlanner` passes `CALENDAR_FREE` to the officer row, roster and bench;
- `PlannerNotSigned`: the in-row answer is free and the ping stays paid.
