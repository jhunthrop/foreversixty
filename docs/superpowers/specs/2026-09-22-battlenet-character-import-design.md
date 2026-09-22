# Battle.net character import, guild discovery, and the signed-in paste

**Date:** 2026-09-22
**Status:** APPROVED by the owner's direction ("sign in with Battle.net => import characters
=> learn guilds"; "build the product for what we will have on launch"). Two lanes execute
it: `bnet-api` (sections 2 to 6) and `bnet-web` (section 7). Section 1 binds both.

Every **RULING** is a call this spec makes where the owner's direction or the code left a
gap, with its reason and the cost of being wrong.

## 0. What exists today (read, not built here)

- Login: `GET /v1/auth/battlenet/start` and `/callback` (`api/internal/auth/handler.go`,
  `bnetStart`, `bnetCallback`). `NewBattleNet` (`api/internal/auth/bnet.go`) already requests
  scopes `openid` and `wow.profile`; `Identify` exchanges the code, reads userinfo, returns
  `BnetUser{Sub, Battletag}` and **discards the access token**. `?next=` is honoured through
  `safeNext`. `Store.UpsertBnetUser` writes `users.bnet_sub`, `users.battletag`.
- `characters` (`0005_logs.up.sql`): `key` (`region/ruleset/slug`, see `character.Key`),
  `region`, `ruleset`, `name`, `class`, `user_id`, `refreshed_at`. **Nothing in production
  writes it**: `auth.Store.LinkCharacter` has only test callers. It is read by `GET /v1/me`
  (`auth.Store.Characters`), by the rankings character page (`rankings/query.go:274`), by the
  rating anonymize check (`rating/store.go:249`) and by the data addon's anonymize exclusion
  (`dataaddon/store.go:49`). So today an account never lists its characters and anonymize
  never applies. Section 3 fixes this at the source.
- `addon_exports` is written only by `POST /v1/addon/exports` (device token,
  `addon.Store.PutExports` → `putOneExport` → `syncGuild`). `syncGuild` parses the export's
  `guild=` section, `resolveGuild` finds-or-creates the `guilds` row case-insensitively on
  `(region, ruleset, lower(name))`, `deriveRank(rankIndex, officerMax)` maps rank index to
  `leader|officer|member`, the `guild_characters` upsert keeps `verified_at`, then
  `AutoConfirmClaimIfPending` for a leader and `afterGuildChange` → `guilds.RecomputeMembership`.
  `guild_characters.source` is checked against `('export','invite','claim')` and
  `verified_by` against `('claim','officer','invite','logs')` (`0018_guild_membership.up.sql`).
- The web paste box (`web/src/components/AddonPasteBox.svelte`, on `/addon`) decodes the
  string client-side (`decodeFS1`), stores the current character in localStorage
  (`writeCurrent`) and links to the tools. It never calls the API.
- `web/src/components/Account.svelte` renders `/v1/me`: a Characters section over
  `me.characters` (`MeCharacter {key, region, ruleset, name, class?}` in
  `web/src/lib/account/api.ts`) and a separate "My guilds" section over `me.guilds`.
- Blizzard's API has **no beta or PTR namespace** (every `*-beta-*`, `*-ptr-*` probe answers
  403). Classic Era is `profile-classic1x-{region}` / `dynamic-classic1x-{region}`; Classic
  progression is `profile-classic-{region}`. Forever's namespace is unknown until Blizzard
  publishes it. Verified on Era on 2026-09-22 with a client-credentials token:
  - `GET https://{region}.api.blizzard.com/data/wow/realm/index?namespace=dynamic-classic1x-{region}`
    → `realms[] {id, name, slug}`; `GET /data/wow/realm/{slug}` → `type {type: NORMAL|PVP|RP|RP_PVP}`,
    `category`, `is_tournament`.
  - `GET /profile/wow/character/{realmSlug}/{lowercaseName}?namespace=profile-classic1x-{region}`
    → `name, level, faction{type}, race{name,id}, character_class{name,id}, realm{slug,name,id},
    guild{name,id}` (guild absent when unguilded), `last_login_timestamp`.
  - `GET /data/wow/guild/{realmSlug}/{guildSlug}/roster?namespace=profile-classic1x-{region}`
    → `guild{name,id,realm{slug}}, members[] {character{name,id,realm{slug},level,playable_class{id}}, rank}`
    (`rank` 0 is the guild master).
  - The account's own characters need the **user's** token: `GET /profile/user/wow?namespace=profile-classic1x-{region}`
    (Bearer = the OAuth access token from the login exchange) → `wow_accounts[] {id,
    characters[] {name, id, realm{slug,name,id}, playable_class{id,name}, playable_race{id,name},
    gender, faction{type}, level}}`. A user token lives about 24 hours and this app requests
    no refresh token (`AccessTypeOnline`).

## 1. Global constraints (both lanes)

1. **Never persist a Battle.net user access token** anywhere (database, logs, files). It is
   used inside the callback request and dropped. "Refresh from Battle.net" is a fresh OAuth
   round trip, not a stored token. Client-credentials tokens are held in memory only.
2. **Never log a token, a `code`, or a `state`.** Log Blizzard status codes, namespaces,
   realm slugs, character counts.
3. **Namespace and region are configuration**, never literals in a handler:
   `BNET_PROFILE_GAME` (default `classic1x`; the value between `profile-` and `-{region}`),
   `BNET_REGIONS` (default `us,eu`; which regional hosts to try for an account). Forever's
   real namespace is one env change on launch day.
4. **A Blizzard roster is authoritative over an export.** Membership rows it produces carry
   `source = 'bnet'` and `verified_by = 'bnet'` with `verified_at = now()`; an export never
   downgrades them (section 4.3).
5. **The `characters` table becomes the single list of an account's characters**, written
   by every path that learns one: the Battle.net import, the companion's export upload, and
   the signed-in paste. Readers already exist and are not changed.
6. Everything else in `.superpowers/journeys/lane-common.md` (worktrees, model budget,
   commit rules, scoped checks, copy modules, function and file size) binds as before.
7. **Lane ownership.** `bnet-api` owns `api/**` and `docs/superpowers/plans/2026-09-22-bnet-api.md`.
   `bnet-web` owns `web/**` and `docs/superpowers/plans/2026-09-22-bnet-web.md`. The
   interface between them is section 6's JSON, frozen here; the web lane stubs it in tests
   and never waits on the API lane.

## 2. Migration `0024_bnet_characters`

```sql
alter table characters
  add column if not exists realm_slug text,
  add column if not exists realm_name text,
  add column if not exists level int,
  add column if not exists faction text check (faction in ('alliance', 'horde')),
  add column if not exists bnet_character_id bigint,
  add column if not exists source text not null default 'export'
    check (source in ('export', 'bnet')),
  add column if not exists imported_at timestamptz;
create index if not exists characters_bnet_idx on characters (bnet_character_id) where bnet_character_id is not null;

alter table guild_characters drop constraint if exists guild_characters_source_check;
alter table guild_characters add constraint guild_characters_source_check
  check (source in ('export', 'invite', 'claim', 'bnet'));
alter table guild_characters drop constraint if exists guild_characters_verified_by_check;
alter table guild_characters add constraint guild_characters_verified_by_check
  check (verified_by in ('claim', 'officer', 'invite', 'logs', 'bnet'));

alter table guilds
  add column if not exists bnet_guild_id bigint,
  add column if not exists realm_slug text,
  add column if not exists roster_refreshed_at timestamptz;

alter table users add column if not exists bnet_imported_at timestamptz;
```

The constraint names must match what `0018` created; the implementer reads `0018` and uses
the real names. A down migration reverses every statement. The migration-numbering test
(`api/internal/db`) demands `0024` because `0023` is main's highest.

**RULING (key format):** the character key stays `region/ruleset/slug` with no realm. On
Forever one realm per ruleset per region is the announced shape, and every reader keys that
way. On Era, where several realms share a ruleset, two same-named characters on different
PVP realms collide; the import keeps the first owner (`user_id` already set by another
account → skip with a logged `collision`) and records `realm_slug` on the row it does write.
Cost if wrong: Forever launches with several realms per ruleset and the key needs a realm
segment, a migration plus a re-key of `addon_exports`, `guild_characters`, `rating_scores`.

## 3. The Blizzard client: package `api/internal/bnetapi`

A small client for the game-data and profile APIs, separate from `auth.BattleNet` (which is
the OAuth login and stays as is).

```go
type Client struct {
  HTTP     *http.Client   // 10 s timeout, injected in tests
  TokenURL string         // https://oauth.battle.net/token
  APIHost  func(region string) string // "https://us.api.blizzard.com"
  ClientID, ClientSecret string
  Game     string         // "classic1x"
  Regions  []string       // {"us","eu"}
  now      func() time.Time
  mu       sync.Mutex; token string; tokenExpiry time.Time
}
func New(cfg Config) *Client
func (c *Client) ProfileNamespace(region string) string // "profile-classic1x-us"
func (c *Client) DynamicNamespace(region string) string // "dynamic-classic1x-us"
func (c *Client) AppToken(ctx) (string, error)          // client credentials, cached until 60 s before expiry
func (c *Client) Realms(ctx, region string) ([]Realm, error)   // index + per-realm type, cached 24 h in memory
func (c *Client) AccountCharacters(ctx, region, userToken string) ([]AccountCharacter, error)
func (c *Client) Character(ctx, region, realmSlug, name string) (CharacterProfile, error)
func (c *Client) GuildRoster(ctx, region, realmSlug, guildName string) (Roster, error)
```

- `Realm{ID int64; Slug, Name, Type, Category string}`; `RulesetOf(realm Realm) string`
  maps `PVP → pvp`, `RP, RP_PVP → rp`, `NORMAL → normal`, and any realm whose `Category`
  or `Name` contains "Hardcore" (case-insensitive) → `hardcore`. Unknown types → `normal`
  with a WARN log naming the type. `RulesetOf` is a pure function with a table test.
- Guild slug for the roster URL: lowercase, spaces to hyphens, everything outside
  `[a-z0-9-]` dropped (Blizzard's own rule). Character name in URLs: lowercase.
- Errors: `ErrNotFound` for 404, `ErrForbidden` for 403 (the namespace does not exist or the
  character is private), `ErrRateLimited` for 429 (the caller stops the batch and logs it),
  anything else wrapped with the status code. Response bodies are decoded into narrow
  structs; unknown fields are ignored.
- Every call is retried once on a 5xx or a network error after 500 ms; never on 4xx.
- Tests use `httptest.Server` fixtures under `api/internal/bnetapi/testdata/` taken from the
  shapes in section 0 (invented names, no real players).

## 4. Import

### 4.1 At login

`bnetCallback` changes as little as possible: `Identify` returns the access token alongside
`BnetUser` (a new field `AccessToken string` on the returned struct, never logged), and after
`UpsertBnetUser` the handler calls `s.Importer.ImportAccount(ctx, userID, token)` **with a
5 second budget** (`context.WithTimeout`), then redirects exactly as before. `Importer` is an
interface on the auth `Service` (nil-safe: when nil, login behaves exactly as today, which is
what every existing auth test sees).

```go
type Importer interface {
  ImportAccount(ctx context.Context, userID int64, userToken string) (ImportSummary, error)
}
type ImportSummary struct {
  Regions    []string // regions that answered 200 for /profile/user/wow
  Characters int      // characters written
  Guilds     int      // guild memberships written
  Skipped    int      // collisions and level < 10 skipped
}
```

The summary is logged at INFO with the user id; an error is logged at WARN and **never fails
the login**. If the budget runs out mid-import, what was written stays written and the
remainder is picked up by the nightly refresh (4.4).

### 4.2 `ImportAccount` (package `api/internal/bnetimport`)

For each region in `BNET_REGIONS`:
1. `AccountCharacters(region, userToken)`. 403/404 → that region has no account: continue.
2. For each character: skip `level < 10` (**RULING:** bank alts and throwaways clutter the
   list; 10 is the level at which Forever's own tools start to matter; cost if wrong: a
   player wonders where their level 8 alt is; they can paste it). Resolve the realm through
   `Realms(region)` → `ruleset`. `key := character.Key(region, ruleset, name)`.
3. Upsert `characters`: `(key, region, ruleset, name, class, user_id, realm_slug, realm_name,
   level, faction, bnet_character_id, source='bnet', imported_at=now(), refreshed_at=now())`,
   `on conflict (key) do update` **only when** `characters.user_id is null or
   characters.user_id = excluded.user_id`; otherwise the row is a collision: count it in
   `Skipped`, log `collision` with the key, move on.
4. `Character(region, realmSlug, name)` → guild name or none. Unguilded: if a
   `guild_characters` row for the key exists with `source = 'bnet'`, delete it and run
   `afterGuildChange` (the player left); rows from other sources are left alone (the export
   is that path's evidence, not this one's).
5. Guilded: `GuildRoster(region, realmSlug, guildName)`; find this character's `rank`
   (0 = guild master). Guild row: `resolveGuild(region, ruleset, guildName)` (moved from
   `addon` into `guilds` as `guilds.ResolveGuild` so both packages share it; the `addon`
   package keeps calling it) and set `guilds.bnet_guild_id`, `realm_slug`,
   `roster_refreshed_at`. Membership: the same upsert as `syncGuild` but with
   `source = 'bnet'`, `verified_by = 'bnet'`, `verified_at = now()`, then the same
   `AutoConfirmClaimIfPending` for a leader and `afterGuildChange`. This upsert is extracted
   from `addon.syncGuild` into `guilds.UpsertCharacterMembership(ctx, tx, MembershipRow)` so
   there is one writer of `guild_characters` for both paths.
6. One transaction per character, so a failure on the 7th leaves the first six written.

Rosters are fetched once per guild per import (a map by `(region, realmSlug, guildName)`).
Each Blizzard call counts against the 5 second budget; a `context.DeadlineExceeded` returns
the partial summary with `err = ErrBudget`.

### 4.3 Precedence with the export path

`syncGuild` (export) keeps its behaviour with one change: when the existing
`guild_characters` row for the key has `source = 'bnet'` and names the **same** guild, the
export does not touch `source`, `verified_by` or `verified_at` (it may still update
`rank_index`/`rank`, since an export is the player's live client). When the export names a
**different** guild than a `bnet` row, the export wins (the player moved and the roster is
stale) exactly as a transfer does today, and the new row is `source = 'export'`, unverified.
**RULING:** the freshest evidence wins on guild identity; Blizzard wins on verification.
Cost if wrong: a forged export could unverify a real member until the next nightly refresh
(4.4) restores the Blizzard row; bounded to one day and visible in `refreshed_at`.

### 4.4 Nightly refresh job `bnet-refresh`

`main.go` gains `bnetimport.RefreshJobCommand = "bnet-refresh"`, created and scheduled by
the owner like `data-addon` (the README recipe is part of the lane). It walks every
`characters` row with `source = 'bnet'` and `refreshed_at < now() - interval '20 hours'`,
re-runs steps 4 and 5 of 4.2 with the app token (no user token needed: character profile
and roster are public data), at most 4 Blizzard calls per second, and stops the run on the
first `ErrRateLimited`. It also runs the namespace probe (section 5).

### 4.5 The companion's export and the signed-in paste also write `characters`

`putOneExport` upserts `characters (key, region, ruleset, name, class, user_id, source='export',
refreshed_at)` with the same ownership guard as 4.2 step 3 (never steal a key from another
account), reading `class` from the export's head the way the planner's decoder does
(`addon.ParseFS1Class`, new, beside `ParseFS1Guild`; if the head has no class the column
stays as it was). This is what makes `GET /v1/me` list a companion user's characters.

`POST /v1/me/exports` (new, session + CSRF, rate-limited like `POST /v1/guilds/{id}/claim`):
body `{"exports": [{"name","region","ruleset","export"}]}`, the same `Export` struct and the
same validation as `putExports` (`MaxExportLen`, non-empty name, valid region and ruleset),
then `PutExports(ctx, userID, exports)`. Response `{"ok": true, "data": {"characters": [<the
`/v1/me` character objects for the keys written>]}}`. A paste is the addon-less way to reach
the same rows the companion writes.

## 5. Namespace probe

`bnetapi.Probe(ctx, region) []ProbeResult` requests
`GET /data/wow/realm/index?namespace=dynamic-{game}-{region}` for each game in
`BNET_PROBE_GAMES` (default `classic1x,classic-forever,classicforever,forever,classic60,anniversary`)
and returns `{Game, Status}`. The refresh job logs one line per game; any game other than the
configured one that answers 200 is logged at WARN as `namespace_appeared`. Also
`GET https://us.version.battle.net/wow_classic_beta/versions` style CDN products are **not**
probed here (that is the data pipeline's job). Nothing acts automatically on a probe result;
the owner changes `BNET_PROFILE_GAME`.

## 6. `GET /v1/me` and the refresh entry point (the frozen interface)

`characters[]` gains fields; existing ones are unchanged:

```json
{
  "key": "us/pvp/thoradin",
  "region": "us", "ruleset": "pvp", "name": "Thoradin",
  "class": "warrior",
  "realm": "Whitemane",
  "level": 60,
  "faction": "alliance",
  "source": "bnet",
  "guild": { "id": 12, "name": "Iron Vanguard", "rank": "officer", "rank_index": 1, "verified": true }
}
```

`realm`, `level`, `faction` are omitted when unknown; `guild` is omitted when the character
has no `guild_characters` row; `verified` is `verified_at is not null`. The top-level
response also gains `"bnet_imported_at": "<RFC3339>"` (omitted when never imported).

Refresh: the web sends the user to `/v1/auth/battlenet/start?next=/account%3Frefreshed%3D1`
(the existing route, the existing `next` rule). A second login of the same account re-runs
4.1. No new endpoint.

## 7. Web (`bnet-web`)

1. `MeCharacter` gains `realm?, level?, faction?, source?, guild?` and `Me` gains
   `bnet_imported_at?` (`web/src/lib/account/api.ts`), with the normaliser treating every
   new field as optional.
2. The account page's Characters section (`Account.svelte`, extracted into a new
   `web/src/components/account/CharacterList.svelte` under 200 lines with its own copy
   module) shows per character: name, realm (when present), level, class, and a guild line
   `Iron Vanguard · Officer` with a "verified" pill when `guild.verified`, or nothing when
   unguilded. Empty state: "No characters yet." with two actions: **Refresh from Battle.net**
   (a link to the start route with `next=/account?refreshed=1`) and **Paste an export**
   (link to `/addon#paste`). When `bnet_imported_at` is set, a muted line "Imported from
   Battle.net <relative time>" and the same refresh link.
3. `?refreshed=1` on `/account` shows a one-line toast "Characters refreshed from
   Battle.net." and is removed from the URL with `history.replaceState` after render.
4. `AddonPasteBox.svelte`: when `fetchMeOnce` says the viewer is signed in, a successful
   decode also `POST`s `/v1/me/exports` (via a new `postMyExports(exports)` in
   `web/src/lib/account/api.ts`, CSRF header as the other POSTs) and shows "Saved to your
   account" or the API's error message; signed out, behaviour is unchanged plus a muted
   "Sign in to keep this character on your account." line. `#paste` scrolls the box into view.
5. The current-character bar (`CurrentCharacterBar.svelte`) shows the guild line when the
   `/v1/me` character matching the current key carries `guild` (read through the existing
   `fetchMeOnce`, no new fetch).
6. Tests: vitest SSR-render tests for `CharacterList` (guilded, unguilded, empty, imported
   line) and for the paste box's signed-in branch; the Playwright account spec gains one
   journey with `/v1/me` stubbed to the section 6 shape (fixture in `web/src/fixtures/`).

## 8. Out of scope

Storing refresh tokens; importing gear or talents from Blizzard (the export remains the
source of a build); guild rosters for guilds nobody on the site belongs to; retail
namespaces; officer tooling changes (a `bnet`-verified leader claims through the existing
flow).
