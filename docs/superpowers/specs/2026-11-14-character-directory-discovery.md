# Character directory: discovery

Date: 2026-11-13 (probes), for the 2026-11-14 design lane.
Scope: what Forever Sixty can learn about a Forever character and guild before the player signs in,
what we already hold, how the Discord raid sign-up gates on a site account today, the Blizzard API
terms that bound a directory, and a proposed design. No code changes in this lane.

Method: real calls with the app's Battle.net client credentials (client credentials grant, app
token only, read-only), and read-only queries against production through the Cloud SQL proxy.
Every count below is from 2026-11-13. Player and guild names are left out on purpose; examples
use made-up names.

## 1. Blizzard APIs without sign-in

### 1.1 Character profile (namespace `profile-forever-{us,eu}`)

The key is the realmless slug of the full name, `{name}-{surname}` lowercased
(`/profile/wow/character/vela-thornbrook`). Every realmed form answers 404 for Forever:
`/{realm-slug}/{name}`, `/{realm-id}/{name}`, `/connected-realm-{id}/{name}`, a bare character id,
and the bare first name. `protected-character/{realmId}-{characterId}` answers 404 to an app token
(it needs the owner's user token).

Probed on 3 public characters (US and EU) with an app token:

| Endpoint (after `/profile/wow/character/{slug}`) | Result | What comes back |
|---|---|---|
| (summary) | 200 | id, name, surname, gender, faction, race, character_class, level, experience, guild (id, name, slug, faction; no realm), ruleset (id, name), last_login_timestamp, average_item_level, equipped_item_level, achievement_points, is_remix, name_search, links to every sub-resource. **No `realm` object.** |
| `/equipment` | 200 | equipped_items: slot, item id, name, quality, level, stats, binding, durability, enchant/transmog where set, media link (13 to 17 items) |
| `/specializations` | 200 | specialization_groups: is_active, per tree spent_points, specialization_name, talents (ids and ranks) |
| `/statistics` | 200 | health, power, primary stats (base and effective), armor, resistances, crit/hit/haste per school, weapon damage and speed, attack power, defense, dodge, parry, block |
| `/achievements` | 200 | total_quantity, total_points, achievements, category_progress, recent_events (small: 3 to 12 per character) |
| `/achievements/statistics` | 200 | categories |
| `/appearance` | 200 | playable race/class, gender, faction, customizations, item appearances |
| `/professions` | 200 | primaries (2) and secondaries (1 to 3), each with tiers and skill points |
| `/pvp-summary` | 200 | honorable_kills, pvp_rank |
| `/character-media` | 200 | assets: avatar, avatar-sd, main-raw, main-sd-raw (portrait and full render URLs) |
| `/encounters`, `/encounters/raids`, `/encounters/dungeons` | 200 | empty for Forever today |
| `/collections`, `/collections/mounts`, `/collections/pets` | 200 | mounts, pets, toys, heirlooms, transmogs |
| `/titles` | 200 | empty |
| `/reputations` | 200 | 14 to 15 factions with standing |
| `/quests`, `/quests/completed` | 200 | completed quest ids (375 to 566 per character), in_progress |
| `/hunter-pets` | 200 for a hunter, 404 otherwise | pets |
| `/mythic-keystone-profile` | 200 | current_period (empty) |
| `/status` | 200 | id, is_valid (the 30-day validation call the terms ask for) |
| `/soulbinds` | 404 | not in this game |

Coverage of characters who never signed up: a sweep of the rosters of the 23 guilds our
signed-in members belong to returned 3,869 members, 3,836 of them unknown to us. Of 60 random
unknown members, 54 (90%) answered 200 on the summary and on equipment; 6 answered 404. All 60
were below level 60 (the game is nine days old). A fresh character that has not yet been made
public answers 404 until it is (see the note in `api/internal/bnetapi/forever_realms.go`).

### 1.2 Guild (namespace `profile-forever-{us,eu}`)

The key is the realmless guild slug, `/data/wow/guild/{guild-slug}`. The realmed form
`/data/wow/guild/{realm-slug}/{guild-slug}/roster` answers 404 (23 of 23).

| Endpoint | Result | What comes back |
|---|---|---|
| `/data/wow/guild/{slug}` | 200 (23 of 23 found from profiles; 27 of the 31 guilds in our table by name) | id, name, slug, faction, ruleset (id, name), member_count, achievement_points, created_timestamp, links |
| `/roster` | 200 (23 of 23) | members: rank (0 = guild master) and character id, name, surname, slug, level, playable_class, playable_race, faction. No realm. Includes non-public and low-level characters. |
| `/achievements` | 200 | empty for Forever |
| `/activity` | 200 | empty for Forever |

The 4 guilds that answered 404 by name are either renamed or disbanded, or our stored name does
not slug the way Blizzard's does; the import already stores `bnet_guild_id` for 28 of 31.

The roster is the single richest source: one call names every member of a guild, with level,
class, race and Blizzard character id, whether or not the member is public.

### 1.3 Game data (namespaces `dynamic-forever-*`, `static-forever-*`)

| Endpoint | US | EU |
|---|---|---|
| `/data/wow/realm/index` | 404 | 404 |
| `/data/wow/connected-realm/index` | 404 | 404 |
| `/data/wow/realm/{id}`, `/data/wow/connected-realm/{id}` (6794, 6929, 6942) | 404 | 404 |
| `/data/wow/search/realm`, `/data/wow/search/connected-realm` | 200, 0 results | 200, 0 results |
| `/data/wow/ruleset/index` | 200: ids 1, 2, 3, 15 | same |
| `/data/wow/ruleset/{id}` | 200: 1 PvP, 2 Normal, 3 Roleplay, 15 Roleplay PvP; status Up; an `auctions` link | same |
| `/data/wow/ruleset/{id}/auctions/` | 200 (an auction listing per ruleset; not explored here) | same |
| `/data/wow/search/character`, `/data/wow/search/guild` | 404 | 404 |
| `/data/wow/playable-class/index`, `/data/wow/item/{id}` (static) | 200 | 200 |
| `/data/wow/pvp-season/index`, `/data/wow/token/index` | 404 | 404 |

There is no character or guild search, so a directory cannot be enumerated from Blizzard alone:
it has to start from names or guilds we already know and grow through rosters.

### 1.4 Rate limits

The terms cap an app at 36,000 calls per hour (section 2(r)); the documented per-second limit is
100. No response carried a rate-limit header in about 550 calls. A sequential client measured about
0.6 s per call, so one worker does about 6,000 calls an hour; the cap is reached with six.

## 2. Realms: where the ids come from

The owner asked whether some other source lists Forever realms with ids, names and rulesets.
Result: **no source does, because Forever realms have no names.** The ids are internal shards.

- Blizzard's own server-status page for `forever-us` and `forever-eu` (its GraphQL `Realms`
  query, read once in a headless browser) lists exactly four rows per region, each flagged
  `isRuleset: true`: Normal, PvP, Roleplay, Roleplay PvP. No realm ids, no shard names.
- `dynamic-forever-*` realm and connected-realm indexes answer 404 and the searches answer
  empty (section 1.3). The sibling namespaces answer, but none lists a Forever id:
  `dynamic-classic1x-*` (66 US / 58 EU realms), `dynamic-classic-*` (50 / 34),
  `dynamic-classicann-*` (5 / 5) hold zero of the 29 ids below; `dynamic-classic-hc-*` and
  `dynamic-classicsod-*` answer 403.
- The profile responses name no realm at all: the character summary, the guild and the guild
  roster have no `realm` object, and the account profile (`/profile/user/wow`, stored for 84
  characters) carries only the `protected-character/{realmId}-{characterId}` link.
- Forever Logs: its site calls an internal `/api/realms`, which refuses outside callers
  (`403 API access not allowed`). Its documented public API (`/api/public/v1/openapi.json`)
  has no realm endpoint, only a free-text `realm` filter; its realm list comes from logs.
- Warcraft Logs: `forever.warcraftlogs.com` exists but sits behind a bot check for scripts; the
  v2 GraphQL API (which has `worldData.region.servers`) needs an OAuth client we do not have
  (we hold only a v1 key, and v1 has no server list). WCL server ids are WCL's own, not
  Blizzard's.
- Icy Veins' realm page and a third-party server-status page refuse scripted requests (403);
  the web search found no Forever armory.

Where realm ids do appear:

1. **Combat-log GUIDs.** `fights.players` holds 4,591 distinct player GUIDs of the form
   `Player-{realmId}-{hex}` across 20 realm ids. The hex part is the Blizzard character id
   (3 of 3 GUIDs whose hex matched a stored `bnet_character_id` were the same character).
2. **The account profile's protected-character link** (signed-in players only).

Realm ids seen in logs: 4613, 4619 (beta-era ids), 6798, 6800, 6801, 6841, 6855, 6922, 6924,
6926, 6929, 6930, 6932, 6934, 6937, 6939, 6941, 6945, 6947, 6949. Of these, 6924, 6926, 6937 and
6941 are not in `forever_realms.go`.

Joining log GUIDs to the 23 guild rosters by character id gives each realm id a ruleset with no
sign-in: US 6924 PvP (2 members) and US 6926 PvP (4 members) are new; every match on an id
already in the table agreed with it. The join must key on region, since character ids repeat
across regions (two EU guild members matched US log GUIDs by id alone).

**Recommendation:** stop needing a realm table. A character's ruleset comes, in order, from its
public profile's `ruleset`, then from its guild's `ruleset` (the roster lists non-public members,
so this covers the fresh characters that today fall into `needs_ruleset`), then from a learned
realm-id table in the database (not committed code) that any character of known ruleset fills,
from logs or sign-in. A realm id is then unknown only until the first character on it is seen
in a guild, a log or a sign-in.

## 3. What we already hold

| Source | Count | Names characters who never signed up? |
|---|---|---|
| `users` | 47 | n/a |
| `characters` | 85, all owned by an account | No: today only sign-ins and addon exports create rows |
| `guilds` | 31 (28 with a Blizzard guild id) | Guild names only |
| `guild_characters` | 34 (31 from Battle.net, 3 from addon exports) | No: `user_id` is not null, so only owned characters |
| Blizzard guild rosters | fetched at import, used only to find the signer's rank, then dropped | **Yes, 3,836 characters across 23 guilds, discarded today** |
| `guild_roster_snapshots` (addon guild tab) | 0 rows | Would, once members run the addon's outbox |
| `fights.players` (log GUIDs) | 4,591 distinct GUIDs in 8 reports | Yes, by GUID; names stay in the raw log and are not kept beside the GUID |
| `fight_metrics` | 10 distinct players | Yes, with name, class, spec |
| `rating_scores` / ratings pack | 0 rows | Would, once rated logs arrive |
| `addon_exports` | 45 | No: the exporter's own characters; the stored FS1 string has no realm or group |
| `discord_links` | 1 | Discord users with a site account only |

The Discord bot cannot list a server's members (it has no member-list intent), so Discord
membership is known only for people who press a button or run a command.

## 4. Discord raid sign-up today

Flow (`api/internal/discordbot/press.go`, `candidates.go`; `guilds/raids_signup_service.go`):

1. `resolveAsker` reads `discord_links` by the presser's Discord id. No row: the ephemeral
   "not linked" answer with a link button. The link (`link.go`) starts at the site and needs a
   signed-in session, which needs Battle.net sign-in. **This is the barrier.**
2. The server must be bound to a guild (`guild_discord_installs`).
3. `SignupCandidates` lists the account's characters in the guild from `guild_characters`
   (`user_id = account`); only `verified` rows can sign up. None: "not a member" with a link to
   `/setup`; unverified: "waiting for approval".
4. The sign-up row (`guild_raid_signups`) needs `signed_by` and `acted_by`, both
   `not null references users(id)`.

Why the account is needed: every sign-up and every roster row hangs off a `users.id`, and the only
proof that a Discord user plays a character is a Battle.net sign-in or an officer's approval.

Minimal change:

1. **Discord-only account.** On a press from an unlinked Discord user, create a `users` row with
   no Battle.net identity (`bnet_sub` is already nullable) plus its `discord_links` row. No email,
   no BattleTag.
2. **Pick from the roster.** Answer with the guild's directory characters (section 6) that no
   verified account holds, through `/signup` autocomplete (a select caps at 25 options, a guild
   roster runs to about 170). The pick writes `guild_characters` with the Discord-only user,
   a new `source = 'discord'` and `verified_by = 'discord'`, and signs up as today.
3. **Who may sign up.** Treat a `discord` claim as signable but marked "claimed on Discord" for
   officers, who can approve, reassign or remove it from the existing roster tools. One Discord
   user per character; a second claimer is told the character is taken and pointed to an officer.
4. **Claim on sign-in.** When a Battle.net import finds a character its account owns that a
   Discord-only user holds, the Battle.net account wins: `guild_characters.user_id`,
   `guild_raid_signups.signed_by/acted_by` and `characters.user_id` move to it. If the same person
   then links that Discord account, the Discord-only user's link and rows merge into the real
   account and the empty user is deleted. If they differ, the Discord user is told the character
   was claimed by its owner.

Copy (`copygen` keys) and the persona gates apply to the new answers when this is built.

## 5. Privacy and terms

Blizzard Developer API Terms of Use (sections as published):

- **2(s)** Refresh at least every 30 days and keep data no longer than 30 days; for WoW characters,
  validate the character id through the API and delete the data when it no longer exists. The
  `/status` endpoint (id, is_valid) is that check.
- **11(c)** When a person asks us to stop using their data, directly, through Blizzard, or by
  de-linking, stop and delete every copy. This is the opt-out.
- **2(p)** A privacy policy covering the data, consistent with Blizzard's; personal information
  only as the law allows and with notice.
- **2(b), 2(h), 2(i)** No for-pay features built on the API data, no selling or transferring it.
  The directory and anything shown from it stay free.
- **2(m)** Name Blizzard as the source, without implying endorsement.
- **2(r), 2(q)** 36,000 calls an hour; no scraping outside the API. The one-off headless read of
  Blizzard's status page in section 2 is not something the import should depend on (and it gives
  no ids anyway).

What the directory may hold for a player who never signed up: public profile data only (name,
level, class, race, faction, guild, ruleset, gear, talents, professions, portrait), refreshed
or dropped within 30 days. Never a BattleTag (the app token cannot read one anyway), never the
account profile, never anything behind a user token.

The privacy page (`web/src/content/pages/privacy.md`) already promises: Blizzard API data kept at
most 30 days and never sold or paywalled; public profile data "no longer linked to you" survives
account deletion and still ages out in 30 days; "builds no profiles". Before the directory ships
the page needs: a line under "What is collected" for public character data of players without an
account and for the Discord id and display name of Discord-only sign-ups; the opt-out path; and a
reading of "builds no profiles" that the directory does not break (it holds public game data per
character, not profiles of people, and is never joined to identity until the player claims it).

## 6. Proposed design

**Directory.** A `character_directory` table keyed by region and Blizzard character id, holding
name, ruleset, realm id when known, class, race, level, faction, guild id, last login, public flag,
sources (roster, log, addon, ranking, sign-in) and `refreshed_at`. A learned `realm_rulesets`
table (region, realm id, ruleset, evidence count) replaces the committed realm map. Seeds:

- every guild we know (the `guilds` table, guilds named in profiles, guilds of logged players):
  one roster call each, daily; each roster adds its members and stamps their ruleset from the
  guild;
- log GUIDs, with the parser keeping the name beside each GUID (realm id and character id from
  the GUID, name from the log); a logged character's guild adds that guild to the sweep;
- addon guild-tab snapshots and ratings rows as they arrive.

**Refresh within the limits.** Budget 20,000 calls an hour, leaving the rest of the 36,000 for
sign-ins and the nightly jobs. Rosters daily; the summary weekly for every directory character
(it doubles as the `/status` validation); equipment, specializations and media every few days
only for characters in guilds with the bot installed or signed up for a raid; everything dropped
when a 30-day validation fails or the row has not refreshed in 30 days. At today's 3,869 known
members this is under 2,000 calls a day; 100,000 characters still fits in under a day a week.

**Discord sign-up by character pick.** As in section 4: a Discord-only account, the guild's
directory roster offered through autocomplete, a `discord` claim marked for officers.

**Claim on sign-in.** The Battle.net import matches its account characters by Blizzard character
id against the directory and moves any Discord-only claim, sign-ups and roster rows to the real
account; the Battle.net owner always wins.

**Opt-out.** A form on the site and a bot command take a character name and region and add a
`directory_optouts` row (region plus Blizzard character id only). The directory deletes the
character, never stores it again, and every surface hides it, as section 11(c) requires. A signed-in
owner's own character settings already cover visibility; the opt-out is for people who never sign in.
