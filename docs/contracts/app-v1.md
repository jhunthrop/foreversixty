# Native apps v1: API and static-data contract (binding for the iOS and Android lanes)

Plan: `docs/superpowers/plans/2026-10-30-mobile-apps-lanes.md`. This file is lane C. It lists every
endpoint and every static file the v1 apps read or write, with the exact shapes the live handlers
answer today. Each response has a JSON fixture under `docs/contracts/fixtures/app/`, and a Go test
in the `api` module serves the real handler with test data and checks its answer against that
fixture. Either app can be built against the fixtures with no network.

## Conventions

**Hosts.** The API is `https://api.foreversixty.gg`. Static data is on the site origin,
`https://foreversixty.gg` (the Cloudflare CDN).

**Envelope.** Every API answer, success or failure, is
`{ "ok": bool, "data": ..., "error": ..., "request_id": "..." }`. A success has `data` and a null
`error`; a failure has a null `data` and `error = { code, message, fields? }`. `fields` is a map
from a field name to a short sentence, present on validation errors. `request_id` is also the
`X-Request-Id` header; send your own `X-Request-Id` to correlate a report. Exceptions that carry no
envelope: `204 No Content` answers, `304 Not Modified`, and the redirects of sign-in.

**Fixtures.** A fixture is named for its endpoint and state, `<endpoint>[.<variant>].json`;
failures carry `error-<code>` in the name. API fixtures are whole envelopes; `cdn.*` fixtures
are the static files as served and `push.*` fixtures are push payloads, neither enveloped. Both
apps read one route table from method and path to these names, generated from the spec's
`x-fixtures` (`ContractRoutes`, used by iOS `FixtureCatalog` and Android `FixtureRoutes`). A fixture is a floor, not a ceiling. Every field it
names is in the live answer with the same JSON kind; a field the live answer adds is allowed, and a
client ignores what it does not know. **Never pin a count, an id or a value from a fixture**: a
fixture shows one state, and a list is longer or shorter in production. A field marked *omitted
when empty* below is absent rather than null; treat absent and null alike. Every other field is
always present (null when unknown, `[]` when an empty list).

**Auth.** Three kinds:
- **public**: no credential.
- **session**: the browser session cookie. A signed-in app sends its **device bearer token**
  (`Authorization: Bearer fsd_...`) instead, and every session route accepts it (lane A1: a token
  issued by app sign-in is an app device, which `RequireSession` admits; a companion's token is
  still refused on session routes). A bearer request needs no CSRF header.
- **device**: a bearer token only. A cookie is refused.

A missing or dead credential answers `401 unauthorized` (fixture `error-unauthorized`; a lapsed app token's one-time answer is `error-session-expired`). A dead bearer token is `401` with
`"that device token is not valid"`. An app token that lapsed (30 days unused or 180 days from
sign-in) answers `401 session_expired` once, then `401 unauthorized`. **On either 401 the app
clears its token and signs in again.** Billing writes (checkout, portal, guild upgrade, transfer)
answer an app token `403 browser_session_required`: send the person to foreversixty.gg in the
browser, never retry. An app token carries no moderator powers.

**Caching.** Cache-Control values are set by the API, and the app's HTTP cache honours them.
Classes: *live public* `public, max-age=30, stale-while-revalidate=300`; *public object*
`public, max-age=60, stale-while-revalidate=600`; *private* `private, no-cache` with
`Vary: Cookie, Authorization` (always revalidate; never a shared cache); *no-store*. Every 2xx
`GET` carries a weak `ETag` over `data`; send it back as `If-None-Match` and a match answers
`304` with no body. Use this on every private read: it makes a refresh free.

**Limits.** 120 requests a minute per IP across the router, and 120 a minute per device token.
`429 rate_limited` is the answer; back off, do not retry in a loop. Writes carry tighter budgets,
stated per endpoint.

**Common errors.** Any route can answer `401 unauthorized`, `429 rate_limited` and
`500 internal` (`"could not ... just now"`; retry later). `404 not_found` is also what a route
answers when the thing exists but is private to you.

| code | status | meaning |
|---|---|---|
| `unauthorized` | 401 | no, wrong or dead credential (fixture `error-unauthorized`) |
| `session_expired` | 401 | an app token lapsed; sign in again (answered once, then `unauthorized`) |
| `browser_session_required` | 403 | a billing write from an app token; open the site |
| `forbidden` | 403 | signed in but not allowed (not a member, not an officer) |
| `not_found` | 404 | no such thing, or private to you |
| `invalid` | 400 or 422 | bad input; `fields` names each (400 body/query, 422 field rules) |
| `conflict` | 409 | the state does not allow this (a raid's status) |
| `premium_required` | 402 | a Premium feature on a free account |
| `rate_limited` | 429 | too many requests |
| `update_required` | 426 | this app build is below the minimum; `fields.store_url` is the store page (see Forced update) |
| `internal` | 500 | server fault |

**Timestamps** are RFC 3339 with a zone, normally UTC (`Z`). Durations are `..._ms` or `..._sec`.
Unknown numbers are `null`, never `0`.

**Client identity (lane P2).** Both apps send two headers on every API request:

| header | value | example |
|---|---|---|
| `X-App` | `<platform>/<version> (<build>)`; platform `ios` or `android`, version and build from `apps/VERSION` | `ios/1.2.0 (42)` |
| `X-App-Install` | a random UUID made on first launch, tied to no user, gone on uninstall (not backed up) | `0b6f6b8e-…` |

The site never sends them. The API logs them and counts requests and errors per platform and
version (`APP_CLIENT_HEADERS`, off by default; docs/runbooks/mobile-apps.md). A missing or
malformed `X-App` is not an error.

**Forced update.** With a minimum set for a platform (`APP_MIN_VERSION_IOS`,
`APP_MIN_VERSION_ANDROID`), a request whose `X-App` names an older version (semver compare on
`<version>`; `<build>` is ignored) is answered before any route runs:

```json
{"ok":false,"data":null,"error":{"code":"update_required","message":"This version of Forever Sixty is too old. Update it from the store.","fields":{"store_url":"https://play.google.com/store/apps/details?id=gg.foreversixty.app","min_version":"1.2.0"}},"request_id":"…"}
```

with status `426`. Both apps then replace every screen with a blocking "Update Forever Sixty"
screen whose one button opens `store_url` (https only; otherwise the app's own store page).

**Character keys.** A character is `region/ruleset/slug`, for example `us/pvp/obnoxious-yell`.
`region` is `us|eu|kr|tw|cn`, `ruleset` is `normal|pvp|rp|hardcore`, and the slug is the name
lowercased with spaces turned into hyphens. The same three segments are the path of every
character route.

## Account

### GET /v1/me
Auth session. Caching private. Fixtures `me`, `me.premium`, `me.empty`.

The home answer: who you are, your characters, your guilds and what you may use. Errors: 401.

`data`:
- `user`: `{ id, battletag (nullable), email (nullable), role, anonymize }`.
- `characters[]`: `{ key, region, ruleset, name, source }` always; omitted when unknown: `class`
  (lowercase slug), `realm`, `level`, `faction` (`alliance|horde`), `race`, `gender`, `item_level`,
  `avatar_url`, `render_url` (Blizzard's render CDN), `spec` (primary tree name), `build`, `guild`.
  `build` is `{ source ("addon"|"bnet"), captured_at, median_sync_gap_sec (nullable), sync_error
  (nullable) }` plus, each omitted when not decoded: `gear` (slot to item id), `talents`
  (`{ trees[], points[] }`, one entry per tree), `level`, `data_build` (the client build the export
  was made on). `guild` is `{ id, name, rank, verified }` plus omitted `rank_index` and
  `attendance` (`{ present, nights }`, only once the guild counts raids).
- `guilds[]`: `{ id, region, ruleset, name, verified, verified_count, can[] }` plus omitted `rank`,
  `consent`, `plan`, `ruleset_label`, `slug`; and always-present nullable `crest_url`, `faction`,
  `roster_count`, `claim`, `verify`, `covered`, `subscribe`. `can[]` lists what you may do there
  (`approve`, `remove`, `settings`, `award`, `billing`, `release`); draw a control exactly when its
  word is present, never from the rank.
- `entitlements`: booleans `server_sims`, `retention`, `multi_compare`, `history`, `notifications`,
  `officer_views`, `roster_check`, `log_coach`, `supporter_mark`, `personal_portal`, and `billing`
  (null without a plan, else `{ plan, source ("stripe"|"grant"), status, current_period_end
  (nullable), cancel_at_period_end, amount_cents, interval, founder }`). The apps carry the web's
  exact rules: a feature is available exactly when its flag is true.
- `finished_sims[]`: your unseen finished server runs, newest first, at most three, only with
  `server_sims`: `{ sim_id, spec, title, dps, finished_at }`.
- `discord_link_available` (bool), `discord_link` (`{ display_name }` or null),
  `discord_link_suggested` (string or null).
- Omitted when empty: `main_character_key`, `bnet_imported_at`, `server_sims`
  (`{ runs_this_month }`), `notifications` (`{ connect, discord }`), `seat`, `seat_offer`,
  `ended_seat`.

A new account with no characters answers `characters: []`, `guilds: []` (fixture `me.empty`).

### DELETE /v1/me
Auth session or app device token. In-app account deletion (App Store guideline 5.1.1(v), lane F1).
Dark: no handler on main yet (the API lanes merge after Nov 6); the shape is final. Body
`{ "confirm": "delete my account" }`, sent only after the person has read what is deleted and
confirmed. Answer **202** `{ requested_at, complete_by }` (`me.delete`): every device token and
browser session of the account is revoked at once, so the app's next call answers 401, and the
data goes by `complete_by`, 30 days on, the privacy page's promise. Errors: 400 `invalid` when
`confirm` is not the phrase (`me.delete.error-invalid`), 401 (`me.delete.error-unauthorized`).
What is removed is the privacy page's list, the same as today's emailed request: see the
operation's description in `api/openapi.yaml`. The app shows `account.delete.done` with
`complete_by`, clears its token and every cache, and returns to signed out.

### GET /v1/devices and DELETE /v1/devices/{id}
Auth session or app device token. `GET` answers the account's devices, the array itself
(`devices.list`): `{ id, name, platform, kind, created_at, last_seen_at? }`. The app marks its own
row by the `device.id` the sign-in exchange returned, which it keeps beside the token. `DELETE`
revokes one device and answers 204; that device's token answers 401 at once; 404 `not_found`
for a device that is not yours (`devices.revoke.error-not-found`).

### PATCH /v1/me
Auth session. Request JSON, at least one of `anonymize` (bool) and `main_character_key` (one of your
own character keys). Answer: the same body as `GET /v1/me` (fixtures `me.patch`). Errors: 400
`invalid` for a broken body or nothing to change (`me.patch.error-invalid`), or a main key that is
not yours (`me.patch.error-not-your-character`).

## Characters

### GET /v1/characters/{region}/{ruleset}/{name}
The character card. Auth public. Caching live public. Fixtures `character.card`,
`character.card.empty`. Errors: 404 `not_found` for an unknown region, ruleset or character, and
for a character whose owner made their account anonymous (`character.card.error-not-found`).

`data`: `character` (`{ key, region, ruleset, name, guild, level, race, realm }`, the last four
nullable, plus omitted `class`, `avatar_url`, `render_url`; `guild` is `{ name, rank }`), `best[]`
and `history[]` (fights; `percentile`, `spec` omitted when unknown; `history[]` adds `kill`,
`duration_ms`, `talent_split`, `execution_score` nullable), `builds_seen[]`, `score_history[]`
(`{ date, score, tier }`, oldest first), `gear[]` (the gear panel, below), `gear_gap_total`,
`gear_spec`, `gear_spec_source` (`fight|talents|default`), `supporter`, and the Forever Score:
`score`, `tier`, `as_of`, `components` (`{ gear, experience, parses, attendance, discipline }`),
`experience` (`{ dungeons, raid_kills, alts }`). **Every score field is null for a character
with no rated fight; render "no logs yet", never 0** (fixture `character.card.empty`).

A gear row is `{ slot, worn, pick, gap }`: `worn` is `{ id, name, item_level }` or null for an
empty slot (`name` and `item_level` are null when the data build does not know the item); `pick`
is the band's recommendation `{ id, name, source, chance }` or null (`chance`, the drop fraction,
is null for a source with no rate); `gap` is the score points the pick adds, or null. `gear` is
`[]` for a character with no addon export. The rules (picks, gaps, scores) live in the API: the
apps never re-derive them.

### GET /v1/characters/{region}/{ruleset}/{name}/home
The home answer (lane A3, dark behind `APP_API`; full rules in lane A3's
`docs/contracts/app-v1-home-config.md`). The character page's visibility and caching. Computed in
Go by `api/internal/upgrades`, the package the Discord `/char` panel reads; show the rows in the
server's order and never re-rank them. `data` is `{ character, state (ok|no_gear|under_20|no_spec|
no_list), spec, band, band_label, preset, unit, unit_word, data_build, upgrade_count, total_gain,
not_sim_checked_count, not_modelled_count, synced_at, synced_from, upgrades[], already_best[],
talents { state, points_differ, band_points } }`. Any state but `ok` has `upgrades` and
`already_best` empty and `total_gain` null. Fixtures `character.home`, `character.home.empty`,
`character.home.no-gear`. Errors: 404 `not_found`. A pick's `source_line` is the site's formatted line (see the BiS
file below) and `drop_chance` the percent when known; lane A3 fills both. The apps list
upgrades measured gains largest first, then gains not modelled yet, then empty slots (coordinator
ruling, 2026-10-31); A3 should return that order too. The band is the site's `bandForLevel` over the character's level (20 to 60 in steps of 10, rounded
down; 60 reads the `raid` preset, every other band `bare`): the apps label Home's band by that rule,
the rule BiS opens the character's list with, so A3's `band` and `preset` must follow it too (a level-58
character is band 50, `bare`).

### GET /v1/characters/search
Typeahead. Auth public. Caching `public, max-age=60, stale-while-revalidate=300`. Query `q` (2 to
24 characters; shorter answers no rows), optional `region`, `ruleset`, `limit` (1 to 10, default
8). Answer `{ rows[] }`, each `{ key, region, ruleset, name, class (nullable), level (nullable),
spec (always null) }`. Fixtures `character.search`, `character.search.empty`. Errors: 400
`bad_request` for an unknown region, ruleset or limit (`character.search.error-bad-request`).

### GET /v1/characters/{region}/{ruleset}/{name}/sim-input
The character model a sim starts from. Auth public; caching `private,
no-cache` for a signed-in caller, `public, max-age=60, stale-while-revalidate=600` otherwise.
Answer `{ spec, spec_source, gear, talents, buffs[], captured_at, source }`. `source` is `addon`
or `blizzard` (an export) or `fight`. `gear` is **a JSON string** (the FS1 export) for an export
and an object `{ trinkets: [item ids] }` for `fight`; branch on `source`, not on the shape.
`talents` is the last fight's split (`"31/20/0"`) or `""`. `spec` and `spec_source`
(`fight|talents|default`) are `""` when unknown. Fixtures `character.sim-input.addon`,
`.fight`. Errors: 404 `not_found` (`character.sim-input.error-not-found`).

## Sims

A sim row's `state` is `queued|running|done|error`. A server run is Premium (`entitlements.server_sims`).

### GET /v1/sims?mine=1
Your sims. Auth session. Caching private. `mine=1` is required; optional `page` (from 1), `kind`
(`run|gear|talents|drops|weights`), `lane` (`server|browser`). Answer `{ rows[], total, page,
per_page }` (100 a page, newest first). A row is `{ sim_id, spec, kind, dps, headline, state,
engine_version, created_at, title, lane, seen }` plus omitted `finished_at`, `failure_reason`
(a plain sentence on `error`), `estimate_seconds` (queued and running only). `headline` is `""`
until done. Fixtures `sims.list`, `sims.list.empty`. Errors: 400 `invalid` (`sims.list.error-invalid`).

### GET /v1/sims/{id}
One saved result. Auth public (the id is the secret). Caching public object. The body is the
engine's `SimResult` (`sim/api` `SimResult`: `engine_version`, `request`, `lane`, `dps`
`{ mean, stddev, error, min, max }`, `iterations_run`, `duration_ms`, `summary`, optional `error`,
`aborted`, `combos`, `equipped`, `stages`, `weights`, `sample`, `healing`, `tank`) plus `sim_id`
and an omitted `title`. The fixture is the real engine fixture
(`web/src/fixtures/sim/result.json`); an app reads `dps`, `lane`, `request.spec`, `title`, and for
bulk runs `combos[]`. Fixtures `sims.result`. Errors: 404 (`sims.result.error-not-found`).

### GET /v1/sims/{id}/progress
Follow a server run, poll every 2 seconds. Auth public. Caching **no-store**. Answer
`{ state, iterations_done, dps (omitted until running or done), stage, combos_done, combos_total }`;
the last three are 0 for a plain run. Stop polling at `done` or `error`, then read the sim (or its
list row for `failure_reason`). Fixtures `sims.progress`, `sims.progress.queued`. Errors: 404.

### POST /v1/sims/{id}/rerun
Start a failed server run again as a new run. Auth session, Premium. No body. Budget 6 starts a
minute per account (shared with `POST /v1/sims/run`). Answer **202** `{ sim_id }` (the new run);
the failed row is left as it is. Fixtures `sims.rerun`. Errors: 402 `premium_required`
(`sims.rerun.error-premium-required`), 404 `not_found` for a row that is not yours or not failed
(`sims.rerun.error-not-found`), 400 `too_large` / `cap_exceeded` when the stored request no longer
fits the server lane, 502 `upstream` when the run was recorded but could not be started (list it
and rerun again). The route is mounted only where the server lane is available; elsewhere it answers
404 `not_found` "no such route".

### POST /v1/sims/{id}/seen
Mark a finished run read (clears it from `finished_sims`). Auth session. Answer **204**, no body.
404 when the run is not yours or not done.

## Guilds

A guild id is a number (`guilds[].id` of `/v1/me`). A raid id is a number. A `raids[]` row's `id` is a report id (a string), not a raid id; its `raid_id` is the raid. "Member" means a
verified member; officers are verified `officer` or `leader` rank.

### GET /v1/guilds/{id}/home
Auth session, verified members only. Caching private. Fixture `guild.home`. Errors: 403 `forbidden`
for a non-member or an unverified one (`guild.home.error-forbidden`), 404.

`data`: `guild` (`{ id, region, ruleset, name, faction, crest_url }`, the last two nullable),
`viewer` (`{ role, verified, officer_tools, raid_guild_tools, verify: { has_approver } }` plus
omitted `character_key`), `claim` (`{ state, frozen }` plus omitted `since`, `claimed_by_name`),
`summary`, `standing` (omitted when it cannot be derived), `reports[]` (the last 7 days;
`next_cursor` omitted), `roster[]` and `pending[]` (roster rows; `pending` is `[]` for a
non-officer), `members[]` and `counts` (the roster as a person counts it, with each member's
`state`: `approved|waiting|no_account|not_seen`), `addon_snapshot` (nullable), `next_raid`
(omitted when none; the band's raid line `{ id, zone, zone_name, title, starts_at, status, size,
guild_timezone, counts, own }`), `attendance_source` (`raids|reports`).

A roster row is `{ character_key, region, ruleset, name, rank, verified, consent, logged_at,
logged_recently, attendance { present, nights }, account_key, supporter, last_report_at, may_remove,
may_approve, addon }` plus omitted `class`, `spec`, `role`, `item_level` (withheld without gear
consent), `attendance_30`, `best_parse`, `rating`, `professions`. Draw Remove and Approve exactly
when `may_remove` and `may_approve`.

**Wave 3 W3-C16 (lane F16, the app's roster).** Each roster row also carries omitted `spec_name`
(the full spec name, catalogue name then class: `Fury Warrior`, never a bare `Fury`) and
`spec_icon` (the talent-tab icon stem), both withheld with `spec`. Each `members[]` entry carries
`role` (`tank|healer|dps`), `spec_name` and `spec_icon`, always present and null when unknown:
read from the entry's own `spec` (the addon snapshot's specs.json key), else from its roster row's
logged spec, and withheld with them. The role chips read `role`; the app never derives a role or a
name from a spec word. `members[].character_key` is absent for a `no_account` name. Fixture
`guild.home.roster` is a 38-person guild for the roster screen (4 in game with no account, 2
waiting).

### GET /v1/guilds/{id}/raids
Auth none required (the actor is read if present). Caching private. Query `limit` (default 20,
max 50), `cursor` (the `next_cursor` of a previous page). Answer `{ rows[], next_cursor (omitted at
the end) }`: `rows[]` are logged raid nights (`{ id (a report id), title, zone, created_at,
duration_ms, fight_count, kill_count, wipe_count, raiders, deaths, fights[], present[], raid_id
(nullable) }` plus omitted `top_parse`). **A verified member also gets `planned[]` (the open raids,
raid views), `zones[]`, `guild_timezone` and `discord`; anyone else gets none of those four keys**
(fixtures `guild.raids`, `guild.raids.public`, `guild.raids.empty`). An officer whose guild `can`
lists `approve` also gets `waiting_for_approval`, the count of characters waiting (lane F4,
handler pending). Errors: 400 `invalid` for a bad cursor, 404.

### GET /v1/guilds/{id}/raids/{raid}
One raid. Auth session, verified members only: anyone else gets 404, never 403. Caching private.
Fixtures `guild.raid`, `guild.raid.done`, and the app's states `guild.raid.officer`,
`guild.raid.locked`, `guild.raid.cancelled`, `guild.raid.no-invite`, `guild.raid.ask-spec` and
`guild.raid.two-characters`. Errors: 404 (`guild.raid.error-not-found`).

The raid view: `{ id, zone, zone_name, title, starts_at, duration_minutes, status, size,
guild_timezone, notes, notes_html, counts, signups[], not_signed[], viewer, discord, discord_post,
pre_done_status, match, night, attendance, report }`.
- `status`: `planned|locked|done|cancelled`. A **locked** raid takes no sign-up from a member;
  an officer may still write.
- `counts`: `{ in, late, bench, out, tanks, healers, dps }`; the role counts count `in` only.
- `signups[]`: `{ character_key, name, class, spec, role, state, note, own }`, `state` is
  `in|late|bench|out`, `role` is `tank|healer|dps`.
- `not_signed[]`: verified roster characters with no sign-up yet.
- `viewer.characters[]`: your own characters in the guild: `{ character_key, name, class, spec,
  verified, state (null until signed), role, raid_spec (null until picked) }`. An unverified
  character is listed but cannot sign up. When `raid_spec` is null, ask "Which spec tonight?" and
  send the answer as `spec`.
- `notes_html` is sanitised HTML of `notes`; render the HTML.
- **Lane F4 additions** (the app's raid page, spec 2026-10-30-mobile-apps C2; the handler adds
  them after launch, and the apps read each as absent until then):
  - `invite_at`: when invites go out, before `starts_at` (the pull); null when the guild sets none.
  - `composition`: the raid leader's read, computed by the API and rendered as is: `classes[]`
    (`{ class, count }`, all nine in the game's order, counting In and Late), `roles` (`{ tanks,
    healers, dps }`, In only), `targets` (the same shape, or null when the raid sets none) and
    `buffs[]` (`{ key, name, label, icon, present, count }`: fortitude, mark_of_the_wild,
    arcane_intellect, blessings, shadow_protection, soulstone, battle_rez and the shaman totems
    windfury_totem, strength_of_earth and mana_spring, from the classes signed In or Late; the
    list is `data/curated/raid-buffs.json`, and the apps draw `label` inside a tile whose icon is
    missing).
  - `signups[]` and `not_signed[]` gain `spec_name` and `spec_icon` (the talent-tab icon stem);
    `not_signed[]` gains `role`, the spec's role, for an officer's Add.
  - `viewer.characters[]` gains `spec_name`, `spec_icon` and `spec_options[]` (`{ spec, name,
    role, icon }`, the class's specs in catalogue order, for "Which spec tonight?").
  - `viewer.can`: `lock` (lock and unlock) and `move` (an officer's PUT for anyone). Draw officer
    tools exactly when the word is there. `viewer.tools_ended`: an officer of a guild whose plan
    ended; the tools draw disabled. `viewer.raid_night`: a Your raid night summary exists.
- After the night: `match`, `night`, `attendance[]` (each row has an `outcome`: `present|late|no_show|...`) and `report` are filled (null before a log is
  matched); `attendance` carries every row for an officer and only your own characters for a
  member (fixture `guild.raid.done`).

### PUT /v1/guilds/{id}/raids/{raid}/signups/{region}/{ruleset}/{name}
Answer for one character: the app's In / Late / Bench / Out buttons. Auth session (a device token
too, once the apps' sign-up is switched on, `APP_NOTIFICATIONS_ENABLED`, lane A2). A member
answers for their own verified characters; an officer of a guild with the plan answers for any
verified roster character, and may write to a locked raid. Budget 30 writes a minute per account.

Request `{ "state": "in|late|bench|out", "role": "" | "tank|healer|dps", "note": "", "spec": "" }`.
`role` empty takes the spec's default. `note` is at most 140 characters. `spec` is optional: a
`specs.json` key such as `warrior-fury`, saved as the character's raid spec before the sign-up; the
sign-up then takes that spec's role. Answer **200** `{ raid, moved }`: the raid view as it now
stands, and `moved`, the name of **your other character that the server moved to bench** (a player
brings one character; signing a second in benches the first) or null. Tell the player when `moved`
is set. A sign-up refreshes the guild's Discord raid post.

Fixtures `guild.signup`, `guild.signup.moved`, and the refusals `...error-raid-cancelled`,
`...error-rate-limited` and `...error-plan-required`. Errors: 422 `invalid` with `fields`
(`state`, `role`, `note`, `spec`; `guild.signup.error-invalid`); 403 `forbidden` for someone else's
character (`...error-forbidden`, the message is the sentence to show), an unverified one, or a
character not verified in the guild; 403 `guild_plan_required` for an officer signing others
without the plan; 404 `not_found` for a raid you cannot see or a character not in the guild
(`...error-character-not-found`); 409 `raid_locked` (`...error-raid-locked`), `raid_cancelled`,
`claim_contested`; 429 `rate_limited`.

### Raid history (Wave 3 F16)
The app's Raid history reads `GET /v1/guilds/{id}/raids` above, 20 rows a page, `next_cursor` for
Load more. Fixtures `guild.raids.history` (a first page from Wed 9 Dec 2026, when raids open, to
31 Dec, with Onyxia at 38 to 40 raiders, Upper Blackrock Spire at 10 and five-player Blackrock
Depths and Lower Blackrock Spire runs, every night a different length) and `guild.raids.history-2`
(the last page, no `next_cursor`). A row's boss count is the distinct `fights[].name`; its size is
`raiders`, which is never below a fight's `players`.

### Calendar link (Wave 3 W3-C17)
- `GET /v1/me/calendar-feed`: auth session or app token. Caching `private, no-store`. Answer
  `{ url, created_at }` (schema `CalendarFeed`). The first call makes the link and answers its
  `url`; every later call answers `url: null` with the same `created_at`, because only the token's
  SHA-256 is stored. Keep the url you were shown in the secure store (Keychain, Android's encrypted
  store), clear it at sign-out, and offer Reset link when you have none. Fixtures `calendar.feed`
  and `calendar.feed.shown`.
- `POST /v1/me/calendar-feed/reset`: revokes the live link and answers a new `{ url, created_at }`
  in one transaction; the old link answers 404 at once. 10 an hour per account (429
  `rate_limited`). Fixture `calendar.feed`.
- `GET /v1/calendar/{token}.ics`: public by token, opened by calendar apps (tag `app-browser`).
  `text/calendar; charset=utf-8`, `Cache-Control: private, max-age=300`, `Referrer-Policy:
  no-referrer`. RFC 5545, one `VEVENT` per planned, locked or done raid of every guild where the
  account is a verified member, from 30 days back; cancelled raids are left out so a subscribed
  calendar drops them; membership is read on every fetch. `UID raid-{id}@foreversixty.gg`,
  `SUMMARY {zone} · {In|Late|Bench|Out|Not signed up}`, `DTSTART` the start, `DTEND` the end,
  `LOCATION {guild} · {zone}`, `DESCRIPTION Pull {time}. {n} of {size} signed up.`, `URL` the raid
  page, no `VALARM`. Golden body: `docs/contracts/fixtures/app/calendar.ics`. 404 `not_found`,
  one body, for a malformed, unknown or revoked token (`calendar.ics.error-not-found`); 60 reads an
  hour per live token, and 30 misses an hour per address (429): only a link that does not exist
  spends the address budget, so calendar services fetching many real links from one address are
  never refused.
- **Security.** The token is 256 random bits (base64url, 43 characters), shown once, stored only
  as its SHA-256 (`calendar_feeds`, migration 0064, one live row per account), revocable, and never
  logged: the API's request and panic logs write the route as `/v1/calendar/:redacted`
  (`httpx.LogPath`). The feed carries the guild's raid times and your own answers, nothing else:
  no raid notes, no other player's name.

### GET /v1/me/raid-nights/{raid}
Your raid night (lane F4 contract; dark, handler pending): the after-raid summary behind the app's
Your raid night screen and the `raid.summary` push. Auth session or device token. Answer
`{ raid_id, guild_id, guild_name, zone_name, starts_at, report_id, character { character_key, name,
class, spec, spec_name, role }, kills, wipes, new_bests, loot[] { item_id, name, quality, icon,
slot, boss, awarded_by, replaces ({ item_id, name, quality, item_level } or null) }, bests[]
{ boss, dps, percentile, kill_ms }, coach_notes }`; `coach_notes` (at most three) is null without
`log_coach`. Fixtures `me.raid-night`, `me.raid-night.free`. Errors: 404 `not_found` when the raid
has no matched log, the viewer has no character in it, or it is not theirs to see
(`me.raid-night.error-not-found`).

### PATCH /v1/guilds/{id}/raids/{raid}
Officers edit a raid. Auth session, verified officer or leader of a guild with the plan, claim not
frozen. Budget 20 officer raid writes a minute per account (shared with lock, unlock, cancel and
create). Request is the whole form: `{ zone ("raid:onyxias-lair", from the raids list's `zones[]`),
title (may be ""), starts_at (RFC 3339, in the future, within a year), duration_minutes (60 to 240
in steps of 30), notes (markdown, at most 2000), timezone (an IANA name), post_to_discord (bool or
omitted), remind_day_before (bool or omitted) }`. `repeat_weeks` is create-only. Answer 200 the
raid view. Fixtures `guild.raid.patch`. Errors: 422 `invalid` with `fields`
(`guild.raid.patch.error-invalid`); 403 `forbidden` (`guild.raid.patch.error-forbidden`, also
`guild_plan_required`, `raid_guild_required`); 404; 409 `conflict` for a cancelled or done raid,
`claim_contested`.

### POST /v1/guilds/{id}/raids/{raid}/lock | unlock | cancel
Officer status moves, same auth and budget as PATCH. `cancel` takes an optional body
`{ post_to_discord }`. Answer 200 the raid view (`guild.raid.patch`). Errors as PATCH, plus 409
`conflict` when the status does not allow the move (`guild.raid.error-conflict`): lock from
`planned`, unlock from `locked`, cancel from `planned` or `locked`.

### GET /v1/guilds/{id}/loot
Auth session, verified members only. Caching private. Fixture `guild.loot`. Answer `{ encounters[]
(the boss picker: `{ encounter_id, name, zone, killed }`), selected, items[] }`. An item is
`{ item_id, name, icon, quality (0 to 5), slot, copies, awarded_to, candidates[] }`; `awarded_to` is
`{ award_id, character_key, name, at, by_name }` or null; a candidate is `{ character_key, name,
class, spec, gain_dps (null when not sim-checked), not_sim_checked, ilvl_delta (nullable),
attendance, already_equivalent, replaces, items_won }`, ranked best first. `icon` is an icon stem
(see Icons). Errors: 403 `forbidden` (`guild.home.error-forbidden`), 404.

The roster is the `roster[]` of the home read; there is no separate roster endpoint in v1.

## Reports

### GET /v1/reports?mine=1
Your reports, newest first. Auth session. Caching private. `mine=1` required, optional `page`.
Answer `{ rows[], total, page, per_page }` (100 a page). A row is `{ id, title, zone, status,
visibility, created_at, fight_count, kill_count, boss_count, logging_character, last_upload_at,
heartbeat_at }`; `status` is `live|processing|complete|failed`, `visibility`
`public|unlisted|guild|private`, `logging_character` is `{ key, name }` or null, `last_upload_at`
and `heartbeat_at` are nullable (a live report with a recent heartbeat is "Live"). `title` may be
`""`; show the zone. Fixtures `reports.list`, `reports.list.empty`. Errors: 400 `invalid`
(`reports.list.error-invalid`).

### GET /v1/reports/{id}
One report summary. Auth public for a `public` or `unlisted` report (caching public object); a
`guild`, `private` or flagged report needs the owner or a verified guild member and is cached
private; **anything you may not read is 404, never 403**. Answer `{ id, title, visibility, owner
({ id, battletag }, or null), guild (omitted when none), zone, status, engine_version, fights[],
players[], created_at, last_upload_at, heartbeat_at, data_base_url, logging_character }` plus
omitted `flagged`. A fight is `{ index, kind ("encounter"|"trash"), name, kill, in_progress, start,
end, duration_ms, players[], deaths, verified }` plus omitted `encounter_id`,
`difficulty`, `size`, `zone`, `boss_health_pct`, `mob_deaths`, `group_kills`, `group_deaths` and
`group[]`. The last four come from engine 0.6.0: `mob_deaths` is every hostile NPC that died in
range, once each; `group_kills` how many of those the logger's group killed; `group_deaths` the
group's own deaths; `group` the players flagged MINE, PARTY or RAID during that fight. A fight
indexed earlier has none of them, and the app shows no kill count for it (the retired
`npc_kills` counted a group kill twice). The per-fight tables are files under
`data_base_url`, which v1 does not read: the app shows the summary and opens the full report on the
web. Fixtures `reports.get`. Errors: 404 (`reports.get.error-not-found`).

`?fold=trash` (Wave 3, W3-C2 lane) adds `trash_groups[]`, each `{ key, zone, fights[], duration_ms }`:
each zone's trash as one entry, zone by zone in the order the night reached them, read from the
report's `report.json` (the database keeps no zone per fight) by the Go port of the web's
`fights.ts fightRows`. Trash with no zone is never grouped. Empty when `report.json` cannot be read.
The apps draw the Fights card from it and never group fights themselves.

### GET /v1/reports/{id}/me (W3-C2)
Your results in one call. Auth session or app. The report's visibility rules exactly: 404, never
403; a flagged report is 404 `flagged`; no report bucket is 503. Answer `{ report_id, characters[]
}`, a character `{ key, name, class, spec (slug, nullable), spec_name, role, fights[] }`, a fight
`{ index, encounter_id, encounter, kill (always true), duration_ms, value, unit ("dps"|"hps"),
percentile (int|null), ranked (int|null), new_best, deaths }`. Parses count boss kills only.
`value` is the role figure (hps for a healer, dps for everyone else, the web Summary tab's rule)
read off the fight's stored summary and rounded to hundredths: it is exactly the `dps` (or `hps`)
of that character's row in the meters below. `percentile` is the bracket placement the web's report
page asks for (the fight's phase, the row's spec, the rounded value), rounded; null when the
bracket has no ranked kill, and the app then prints no parse number. `new_best` is higher than
every earlier ranked kill of that boss and difficulty by that character; a first kill is not one.
The spec's `rank` is left out: a rank among kills would contradict the character-board rank F11
prints. Fixtures `reports.me`, `reports.me.empty`; errors `reports.me.error-not-found`.

### GET /v1/reports/{id}/fights/{n}/meters (W3-C3)
One fight's players for the Details!-style sheet. `?metric=damage|healing|taken` (default damage),
`?scope=trash` folds a zone's trash (n is its `trash_groups[].key`). Auth as the report. Answer
`{ fight_index, scope, metric, kill, duration_ms, players[] }`, a player `{ key (nullable), name,
class, spec (nullable), spec_name, role, active_ms, damage_done, healing_done, damage_taken, dps,
hps, dtps, deaths, mine }`, sorted high to low by the metric's per-second figure; a player with none
of it is left off. Per-second figures are the total over the player's time in the folded pulls,
rounded to hundredths (the web's `night.ts`). `key` is set only on a public or guild report and
only for a character with a page. A fight in progress has no meters (404). Fixtures
`reports.meters`; errors `reports.meters.error-not-found`.

### Coach (W3-C4): GET /v1/reports/{id}/coach/night, GET /v1/reports/{id}/fights/{n}/coach, GET /v1/me/coach/recent
Existing routes, now app operations; no behaviour change. The app draws the coach card only with
`entitlements.log_coach`: it asks the night line for the character (`?player=key`; data null means
no card), then that line's `fight_index` review. Free accounts never call them (the night route
would answer 403 `no_plan`, `coach.night.error-premium-required`). The app prints the review text
as given. Fixtures `coach.night`, `coach.fight`, `coach.recent`.

## Builds

### GET /v1/builds?mine=1
Your saved planner builds, newest first. Auth session. Caching private. `mine=1` required, optional
`page`. Answer `{ rows[], total, page, per_page }`; a row is a build (below). Fixtures
`builds.list`, `builds.list.empty`. Errors: 400 `invalid` (`builds.list.error-invalid`).

### GET /v1/builds/{id}
One build. Auth public. Caching `public, max-age=86400`. A build is `{ id, class_id, race_id,
tree_version, point_order[], gear, created_at, views }` plus omitted `title`. `point_order` is the
talent ids in the order the points were spent; `gear` maps a slot (`head`, `finger1`, `trinket2`,
`main_hand`, ...) to an item id (an open map); `tree_version` names the client build whose static
data it was made against (`/data/<tree_version>/talents/<class>.json`): **read the trees from that
build, not the active one**. Fixtures `builds.get`. Errors: 404 (`builds.get.error-not-found`).

### POST /v1/builds
Save a build. Auth public (a signed-in caller owns it). Budget 20 saves an hour per IP. Request
`{ class_id, race_id, tree_version, point_order[], gear, title }` (at most 8 KB; title at most 60
characters). Answer **201** (new) or **200** (the same build again) `{ id, url }`; the id is a hash of
the content. Fixtures `builds.save`. Errors: 400 `invalid` with `fields` keyed by `point_order[n]`
and the like (`builds.save.error-invalid`).

### DELETE /v1/builds/{id} (W3-C9, Wave 3 lane F13)
Take a build off the caller's own list (Your builds). Auth session or device. Answer **204** with no
body. The record and its `/b/{id}` link stay and keep working for anyone who has them; nothing is
deleted (`builds.user_id` is already nullable, migration 0015, so this lane adds no migration).
Saving the same build again (`POST /v1/builds`, the same content) puts it back on the list, which
is the apps' Undo. Errors: **404** `not_found` when the build is not on the caller's list (another
player's, nobody's, or no such id; never 403) (`builds.delete.error-not-found`).

### The planner's rules (Wave 3 lane F13)
The apps' planner refuses exactly what the site planner and `POST /v1/builds` refuse. The rules
live in `web/src/lib/planner` (`rules.ts`, `fs1.ts`, `derive.ts`); `web/scripts/gen-planner-rules.mjs`
replays them over the active build's talent files into `apps/shared/fixtures/planner-rules.json`
(add, remove, validate, split, points per level, FS1 code to order), and both apps' tests replay
every vector. The refusal sentences are the planner's own, copied into the app catalog as
`plan.refuse.*`; the app adds one, the level budget (`level - 9` points, at most 51): `No points
left. Take one off first.` `src/lib/mobile/planner-rules-codegen.test.ts` fails CI on drift.

## Static files

Served from `https://foreversixty.gg`. A build directory `/data/<build>/` is immutable per build in
practice (it can be regenerated for a corrected pipeline). The cache rule is
`public, max-age=3600, stale-while-revalidate=86400`; icons and tree art are
`public, max-age=31536000, immutable`. No envelope. Fixtures `cdn.*`; the tests hold them to the
data generator's real output for the active build.

**Finding the build.** Two ways: the active build is named by `/data/active-build.json`, and a
character's own export names the build it was made on (`/v1/me` `characters[].build.data_build`).
For the lists and the BiS, read the active build. Build ids look like `1.60.1.70658`.

> **Publication.** `web/scripts/sync-data.mjs` copies `web/src/data/active-build.json` and
> `dates.json` to `public/data/` on every build (lane F4, 2026-10-31), so they are served at
> `/data/active-build.json` and `/data/dates.json` from the next site deploy. Before that the CDN
> answered 404 and an app had no build to hang icon, list and BiS URLs off.

### /data/active-build.json
`{ "build": "1.60.1.70658" }`. Fixture `cdn.active-build`. Fetch on launch; revalidate (an hour).

### /data/{build}/bis/{spec}.json
The BiS lists for a spec (`warrior-fury`; one file per spec). Fixtures `cdn.bis.warrior-fury`, `cdn.bis.hunter-beast-mastery` (its level-60 bands).
**Pending (lane F2 fix round).** Each slot and alternative gains `source_line`, the site's
formatted source (`describeSourceCell` in `web/src/lib/bis/source-cell.ts`: `Quest: {name} ·
requires level {n}`, `{n}% from {boss}`, `Crafted: {profession}`, `{instance} · {boss}`), never
repeating the item's name, and `drop_chance` (percent, or null). The apps render `source_line`
and fall back to `source`. The hunter fixture carries both, from the site's formatter over
`loot.json`.
`{ spec, build, engine_version, generated_at, bands[], presets }`. A **band** is one level band
(`20|30|40|50|60`) for one `faction` (`alliance|horde`) under one `preset` (`bare`, and `raid` at
60): `{ spec, band, preset, faction, race, talents, talent_points, weights[], slots[], set_dps,
score_unit, reference_dps_per_point, ... }`. Pick a band by (band, faction, preset). Read:
- `slots[]` (the list): `{ slot, verified }` always; when the slot has a pick: `item_id`,
  `item_name`, `source` (the game's words), `source_kind` (`vendor|rep|quest|crafted|dungeon|raid|...`),
  `alternatives[]` (runners-up: `{ item_id, item_name, source, source_kind, dps_delta }`, `dps_delta`
  is negative against the pick), `score`, `sim_dps`, `chance` (drop fraction), `effect_unmodelled`
  (true when a proc the sim leaves out makes the item good: say so), `swap_note`, `ties`
  (`[{ item_id, item_name }]`, items that measure the same), `set_bonus` (`{ set, pieces, bonus }`),
  `faction_note`, `low_value`.
- `role` (`dps|healer|tank`): a healer band's `dps_delta` and `set_dps` are healing per second, a
  tank band's its tank score. `new_at_band[]` is `"<slot>: <item name>"` strings.
- `empty_reason` on a slot with no pick (`two_hand_equipped` for an off hand under a two-hander).
- `weights[]`: `{ stat, weight, error, scale_factor, dps_per_point, weight_per_percent, insignificant }`.
  **Show the scale factor** (normalised so the top stat is 1.00); haste is per 1%.
- `presets.raid.label` and its buff lists describe the raid preset.
A band's `slots[].item_id`s resolve in `items/<class>.json`. Do not quote numbers in prose: render
them from this file. Item ids are positive; `item_id <= 0` or absent means no recommendation.

### /data/{build}/items/{class}.json
`{ build, class_slug, items[] }`, one file per class (`warrior`, `hunter`, ...). An item is `{ id,
name, icon, slot, quality, required_level, item_level, armor, stats, damage_min, damage_max, speed,
dps, two_hand, effect_text, stats_source, set_id, unique, client_unconfirmed, percent_stats,
icon_source, weapon_type, superseded_by, required_level_source }`. `quality` is 0 to 5 (poor to
legendary; draw the quality colours). `stats` and `percent_stats` are open maps of a stat to a
number; `set_id` is nullable, with sets in `sets.json`. `slot` is the item slot
(`head ... ranged`, `main_hand`, `off_hand`, `finger`, `trinket`). `client_unconfirmed` marks stats
the client has not confirmed. The tooltip is built from these fields. Fixtures `cdn.items.warrior`,
`cdn.items.hunter`.

**Spec C6 item facts.** `bonding` (`pickup|equip|use|quest`, from ItemSparse `Bonding`; for an
item the client's ItemSparse lacks, classic-db's `item_template.bonding`, the same enum) and
`armor_type` (`cloth|leather|mail|plate|shield`, from Item `SubclassID` when `ClassID` is 4; a
cloak is `cloth`, as the client prints `Back ... Cloth`) are published by the generator from the
next data run of each build. Still pending: `min_standing` (`hated` ... `exalted`, from ItemSparse `MinReputation` when `MinFactionID` is
set). Each is nullable and absent until published; draw nothing for a missing one. The
`cdn.items.hunter` fixture carries them, read from the client's own Item and ItemSparse rows.
`stats_source` is null on most real rows.

**Pending (lane F2 fix round).** `subclass` (a weapon's type word: `polearm`, `sword`, `gun` ...,
from Item `SubclassID` when `ClassID` is 2) and `weapon_slot` (`one_hand`, `two_hand`,
`main_hand`, `off_hand`, `ranged`, `thrown`, from Item `InventoryType`): the client's
`Two-Hand ... Polearm` line. Nullable, absent until published.

### /data/{build}/sets.json
`[{ id, name, item_ids[], bonuses[] }]`, a bonus `{ pieces, description }`. An item's `set_id`
names its set; the tooltip draws the set's name and bonuses. Fixture `cdn.sets`.

### /data/{build}/talents/{class}.json
`{ build, class_id, class_slug, trees[] }`; a tree is `{ id, name, position, background, icon,
talents[] }`, a talent `{ id, name, icon, max_rank, tier, column, prereq_talent_id (nullable),
prereq_rank (nullable), ranks[], spell_id }`, a rank `{ spell_id, description }`. `tier` and
`column` place the talent on the grid; `background` names `/data/<build>/trees/<background>.webp`.
`id` is the id a build's `point_order` and `GET /v1/builds` use (the client's trait node id);
`spell_id` is the spell. Fixture `cdn.talents.warrior`.

### /data/{build}/classes.json
`[{ id, name, slug, color, forever_changes[] }]`, one per class, `color` a hex. `forever_changes[]`
is `{ text, sources[] }`. Use the id to read `class_id` of a build. Fixture `cdn.classes`.

### /data/{build}/races.json and /data/{build}/combos.json (W3-C10)
`races.json` is `[{ id, name, slug, faction (alliance|horde), placeholder, forever_changes[] }]`;
`placeholder` marks a race the client names but nobody can play yet. `combos.json` is
`[{ race_id, class_id, new_in_forever }]`, one per legal pair: a pair it does not list cannot be
made, and `POST /v1/builds` refuses it. The planner's Race sheet lists the non-placeholder races
paired with the class. Both are copied by `web/scripts/sync-data.mjs`. Fixtures `cdn.races`,
`cdn.combos` (held to the generator's files by `cdn_test.go`).

### /data/{build}/icons/{icon}.webp
The game icon for any item, talent or spell: `icon` is the stem a JSON file carries
(`inv_helmet_09` is `/data/<build>/icons/inv_helmet_09.webp`). A faithful HD, scalable reproduction;
never reinterpret or re-badge. A build that does not carry an icon answers 404: fall back to the
neutral placeholder. Tree art is `/data/<build>/trees/<background>.webp`.

### /data/dates.json
The key dates (see the open item above). `[{ key, value, ... }]`: `key` (a display date, "Nov 4"),
`value` (the label), omitted `note`, `iso` (YYYY-MM-DD), `endIso`, `time`, `running` (the sentence
while the date is in progress), `role`, `highlight`. Sort and compare by `iso`/`endIso`, never by
`key`. Fixture `cdn.dates`.

### Guides, tier lists and the changelog (W3-C1, Wave 3 lane F9)
Written at site build time by `web/src/lib/app-data` (Astro endpoints under `web/src/pages/data`)
from the guide Markdown and the BiS files, through the same `lib/guides` and `lib/tiers` functions
the web's guide and tier pages call. Public, no envelope. The site build fails a file over its
gzipped budget: a guide 60 KB, `tiers.json` 40 KB, the index 10 KB.

- `/data/{build}/guides/index.json`: `{ build, classes: [{ class, class_name, specs: [{ spec,
  spec_slug, spec_name, role, icon, updated, path }] }] }`. Classes in `classes.json` order, specs in
  tree order; `path` is the guide's file under `/data/{build}/`. Fixture `cdn.guides.index`.
- `/data/{build}/guides/{class}-{spec}.json`: one guide. `updated` is the web header's stamp (the
  later of the guide's date and its BiS file's). `stat_priority` is the web's pill order.
  `overview { strengths, weaknesses, raid }` are plain-text lines; `role_line` is the raid block's
  Role line. `talents { key[], leveling_build, raid_build }`: a key talent is `{ name, icon, rank,
  max_rank, text, why }` (`text` the top rank's game text, `why` null when it would repeat it); a
  build is `{ code, totals[], tree_names[] }` with `raid_build` null when it is the leveling build.
  `rotation { opener, priority, cooldowns }` (null with no rotation) lists `{ spell_id, name, rank,
  icon, clause }`, the rank the top one learned by 60. `gear { spec, band, preset_label }` names the
  BiS list the Gear row opens. `sections` are the sections the guide has, in reading order. Fixtures
  `cdn.guides.shaman-elemental`, `cdn.guides.mage-fire`, `cdn.guides.warrior-protection`.
- `/data/{build}/tiers.json`: `{ build, updated, preset_label, factions: { alliance|horde: { dps|tank|healer:
  { notes[], rows[] } } } }`. A row carries the web's rank, gap, bar fraction, tie, tier letter (null
  on tank and healer lists), `metric_text`, `gap_text` and the callout's `title` and `detail`
  sentences, so the apps never sort, gap or letter. The apps show it only behind the `guide_tiers`
  switch, off until the accuracy reviewer verifies every DPS band. Fixture `cdn.tiers`.
- `/data/changelog.json`: `{ entries: [{ id, date, title, kind, blocks: [{ type: p|h|li, text }],
  sources[] }] }`, newest first. `kind` `site` is a change to this site; the rest are game news.
  Fixture `cdn.changelog`.

## App sign-in and notifications (lanes A1 and A2)

These routes ship dark: unmounted until `APP_SIGN_IN=true` (A1) and `APP_NOTIFICATIONS_ENABLED=1`
(A2) are set, so they answer 404 until then. The shapes below are the lanes' final contracts
(reports `docs/superpowers/reports/2026-10-30-mobile-auth.md` and `...-mobile-push.md`). There is
no handler on main yet, so these fixtures are held only to the envelope rule; add an
`app_contract_test.go` beside each handler when its lane merges.

### Sign in (A1)
The app opens the system browser (ASWebAuthenticationSession, Custom Tabs) and never sees the
Battle.net credentials.

1. The app makes a random `code_verifier` (43 to 128 characters of `A-Za-z0-9._~-`) and a random
   `state` (16 to 128 of the same). It opens
   `GET https://api.foreversixty.gg/v1/auth/app/start?platform=ios|android&code_challenge=<base64url
   SHA-256 of the verifier, no padding, 43 characters>&code_challenge_method=S256&state=<state>`
   (public; 60 an hour per IP, shared with continue). Errors: 400 `invalid` with `fields` naming
   each bad parameter (`app.signin.start.error-invalid`), 429.
2. Start answers **200 with a confirm page** (`text/html`, no script, framing refused): "Sign in to
   the Forever Sixty app on your phone". The person taps "Continue to Battle.net", which is
   `GET /v1/auth/app/continue?state=...`; that answers 302 to Battle.net, or 400 `invalid` from a
   browser that did not start the flow. The flow may sit at Battle.net for 10 minutes.
3. After Battle.net the API redirects to **exactly** `https://foreversixty.gg/app/auth?code=<one-time
   code>&state=<your state, verbatim>`. There is no `/callback` suffix; register exactly that
   universal link / App Link. **Any `redirect_uri` the client sends is ignored**, so send none.
   **Check the returned `state` equals the one you sent.** The code is 43 base64url characters,
   single use, and expires 2 minutes after the callback. A forged, expired or replayed flow answers
   400 `invalid` ("start again in the app") and mints nothing.
4. `POST /v1/auth/app/exchange` (public, no cookies; 60 an hour per IP), body `{ code,
   code_verifier, device_name }` (`device_name` optional, trimmed to 60, otherwise "iOS" or
   "Android"). Answer **201** `{ token, device: { id, name, platform }, me }`. `token` is the device
   bearer token, shown once: keep it in the Keychain / encrypted storage. `me` is the body of
   `GET /v1/me`, so the first screen needs no second call. Fixtures `app.signin.exchange`. Errors:
   400 `invalid` for a malformed body, no code or a bad verifier (`...error-invalid`), 400
   `invalid_grant` for an unknown, spent, expired or mismatched code (`...error-invalid-grant`;
   start again; any attempt spends the code), 429.
5. From then on send `Authorization: Bearer <token>` on every call. The token reads every session
   route. On `401 unauthorized` or `401 session_expired`, clear it and sign in again; billing
   writes answer `403 browser_session_required` (see Auth above).

### POST /v1/auth/app/logout
Auth device. Revokes this device's token. Answer 200 `{ signed_out: true }` (`app.logout`); 401
with no token or a dead one. The apps call it on sign-out, then clear the token whatever the
answer. The account page lists app devices (`GET /v1/devices` items are `{ id, name, platform,
kind ("companion"|"app"), created_at, last_seen_at? }`) with Revoke; a revoke, `DELETE
/v1/sessions` or `POST /v1/auth/logout` makes the token answer 401.

### GET /v1/me/notifications and PATCH /v1/me/notifications
Auth session or device. Caching **no-store**. The player's choices per event and channel. Answer
`{ time_zone, quiet_hours { enabled, start, end }, events[] }`; an event is `{ kind, who, officer,
quiet_exempt, time_sensitive, channels { push, discord } }`, every catalogue event in catalogue
order. `kind` is the catalogue's name (`raid.posted`, `raid.signup_nudge`, `raid.changed`,
`raid.roster_set`, `raid.roster_moved`, `raid.starting`, `raid.summary`, `sim.finished`,
`report.ready`, ...; the list grows), `who` the audience key (`player`, `guild_members`,
`signed_up`, `in_and_late`, `raiders`, `officers`, `payer`, `everyone`), `officer` true for an
event only officers receive (show it only to officers), `quiet_exempt` true for an event sent
inside quiet hours, `time_sensitive` true for one that breaks through Focus (`raid.roster_moved`,
`raid.starting`). Show copy from your own strings keyed by `kind`; hide an unknown `kind`. Times
are `HH:MM` in `time_zone` (an IANA name).

PATCH body, every key optional but at least one: `{ time_zone, quiet_hours { enabled, start, end },
events: [{ kind, channel ("push"|"discord"), enabled }] }`. Answer 200 the GET body. Fixture
`app.notifications`. Errors: 400 `invalid` for a body that is not JSON; 422 `invalid` with `fields`
(`time_zone`, `quiet_hours`, `events`, `events[i]`; `app.notifications.patch.error-invalid`), and
422 "nothing to change".

### POST /v1/me/push-tokens and DELETE /v1/me/push-tokens/{token}
Auth device only (a cookie is refused). Register after the permission prompt and again whenever
the OS rotates the token. Body `{ platform ("ios"|"android"), token, app_version }`: an iOS token is
64 to 200 hex characters, an Android token the FCM registration token (`[A-Za-z0-9_:-]`, 16 to
4096 characters); `app_version` is optional, at most 32 characters. Answer **201** `{ platform,
app_version }`. A device holds one token at a time; registering an existing token moves it to this
device. `DELETE` answers **204** whether or not the token was there, and touches only the calling
device's token; send it on sign-out and when the player turns notifications off. Revoking the
device drops its tokens. Fixtures `app.push-tokens.register`. Errors: 400 `invalid` for a body
that is not JSON; 422 `invalid` with `fields` (`platform`, `token`, `app_version`;
`app.push-tokens.register.error-invalid`).

### Push payloads
What APNs and FCM hand the app (lane A2, `api/internal/push`). Fixtures `push.raid-posted.apns`,
`push.raid-starting.apns`, `push.raid-posted.fcm`.

- **APNs**: `aps.alert { title, body }`, `aps.sound` (`default`), `aps.thread-id`, `aps.category`
  (`RAID_SIGNUP` when the payload carries actions), `aps.interruption-level` (`time-sensitive` for
  `raid.roster_moved` and `raid.starting`, otherwise absent), and top-level `kind`, `link`,
  `actions` (omitted when none) and `data` (string values). The `apns-collapse-id` header carries
  the collapse id.
- **FCM**: `message.notification { title, body }`; `message.data` holds strings only: `kind`,
  `link`, `thread`, `category`, `actions` (the actions array as a JSON string) and the event's
  data; `message.android.collapse_key` and `message.android.notification.tag` are the collapse id,
  `message.android.notification.click_action` the category. `android.priority` is `HIGH` only for
  a time-sensitive event, otherwise `NORMAL`.
- **Ids.** The thread id is `raid-<raid id>`, so one raid's notifications group together. The
  collapse id is `<kind>-<raid id>` (at most 64 bytes), so a later notification of the same kind
  for the same raid replaces the earlier one.
- **`data`** on a raid event is `{ guild_id, raid_id }` (decimal strings). A raid event with
  actions also carries (lane F4 contract, A2 pending) `character_key`, the character the actions
  answer for (the one signed for this raid, else the main when it is a verified member, else the
  only verified one); `ask_spec` (`"true"` when that character has no raid spec and its class has
  two or more); and `answers`, a JSON string of `{ in, late, out, moved, locked, cancelled, failed }`
  each `{ title, body }` in the player's locale, which the app posts in place of the notification
  after an action (`moved.body` holds `{name}` and is appended to the In or Late body). Without
  `character_key` the app picks by the same rule from the raid view, and opens the raid page when
  it cannot. `link` is a
  `https://foreversixty.gg/...` page (guild, raid, character, sim, report, account): open it as a
  deep link, or in the browser when the app has no screen for it.
- **The RAID_SIGNUP category** has three buttons, In, Late and Out, on `raid.posted` and
  `raid.signup_nudge`. Register the category with these action ids. Each action is `{ id
  ("signup_in"|"signup_late"|"signup_out"), title ("In"|"Late"|"Out"), method ("PUT"), path,
  body ({ state }) }`, and `path` is `/v1/guilds/{g}/raids/{r}/signups/{character}`. **The app
  replaces `{character}`** with the `{region}/{ruleset}/{name}` of the character it answers for
  (the server sends one event to the whole guild and cannot pick), then sends the PUT with its
  device token. The PUT body writes `note` too, so an action without one clears the stored note.

### Officer raid actions from the app
Lock, unlock and cancel are `POST /v1/guilds/{id}/raids/{raid}/lock`, `/unlock` and `/cancel`
(see Guilds above for the bodies, budget and errors); moving a player between In and Bench is the
sign-up PUT with `{ "state": "in" | "bench" }` for another account's character. All of them take the
app's device token once A1 admits app tokens on session routes; the sign-up PUT takes one only
while `APP_NOTIFICATIONS_ENABLED` is on. There is no delete route: Out removes a player.

### GET /v1/app/config
Auth public. Caching `private, max-age=300`, `Vary: X-App` (never the CDN, which does not key on `X-App`). Call on launch. Answer
`{ web_checkout_regions[], min_app_version { ios, android }, active_build, features { <key>: bool } }`. `web_checkout_regions` are the
storefront countries (ISO 3166-1 alpha-2, upper case) where the app may show a "See plans" button
that opens `https://foreversixty.gg/premium` in the browser; read the storefront country from
StoreKit or Play Billing and show the button exactly when it is in the list, and show no purchase
prompt or call to action anywhere else. The list is per platform, read from the `X-App` header
(decision L: `ios` gets `["US"]`, `android` gets `[]` until the Play programme enrolment), so the
response carries `Vary: X-App`. `min_app_version` is the lowest version the API supports
for each platform (dotted, compared numerically): below it, block with an "Update the app" screen.
`active_build` is the client build the static data serves. `features` holds the feature
switches by key (`apps/shared/features.json` lists every key, its lane and its bundled default;
every one ships `false`). A feature lane ships dark in both apps and the API turns it on in both
at once. The apps read the config once at launch and cache it (iOS `AppConfigStore`, Android
`AppConfigRepository`): a key the answer leaves out takes the bundled default, a key the build
does not know is ignored, and offline the cached answer, else the defaults, stands. The map is
always present (`{}` when nothing is switched). Fixture `app.config`. The handler is lane A3's
(`api/internal/app`, behind `APP_API`). It fills `features` from `APP_FEATURES` (a comma list
of keys, each sent `true`; unset is `{}`), and the checkout lists from
`APP_WEB_CHECKOUT_REGIONS_IOS` (default `US`) and `APP_WEB_CHECKOUT_REGIONS_ANDROID` (default
`none`); a missing, malformed or unknown `X-App` platform gets `[]`. All of it is server config:
a change needs no code and no app release.

## Live surfaces (Wave 3 lane F14)

Spec `design/specs/2026-10-31-mobile-wave3.md` §5.6, contracts W3-C12 to W3-C14. Dark: mounted
only with `APP_WIDGETS=true` (404 until then), and the apps draw nothing until the `widgets`
feature switch is on. Handler `api/internal/livesurface`.

### GET /v1/me/widget
Auth session or device. Caching private; send the ETag back (an unchanged answer is a free 304).
Query `character` (optional): the widget's setting, a key of one of your characters; empty or
unknown reads your main, else your first. **The one snapshot every surface draws from**: the
home-screen and lock-screen widgets, the Live Activity and the Android ongoing notification and
quick-settings tile all read this answer from the app's shared container (iOS App Group, Android
DataStore) and compute their figures from it with the same clock, so no two surfaces disagree.
`data` is `{ generated_at, character, guild, next_raid, last_raid, upgrades, live }`:
- `character` `{ key, name, class, spec_name }` or null (no characters).
- `guild` `{ id, name }` or null: `No guild yet`.
- `next_raid` or null (`No raid planned`): the first raid of the raids tab's calendar that has not
  ended, cancelled ones included. `{ raid_id, guild_id, guild_name, zone, zone_name, starts_at,
  invite_at (null until raids carry an invite time), pull_at (= starts_at), ends_at, size,
  signed_up (In plus Late, the raid page's band count), status ("open"|"locked"|"cancelled"),
  my { state ("in"|"late"|"bench"|"out"|null), spec (null: `Choose a spec`), role }, can_sign_up }`.
  A press on In, Late or Out is the sign-up `PUT` above.
- `last_raid` `{ zone_name, starts_at }` or null: `Last: {zone}, {d Mon}.`
- `upgrades` `{ state, sentence, count, total_gain, unit, rows[] }` or null: the Home answer's
  first three rows in its order with the whole list's count and total; any state but `ok` has no
  rows and a `sentence` the widget shows as is (an `ok` answer with no rows says every slot
  matches). Null while the Home answer (`APP_API`) is off.
- `live` `{ sim_runs_running }`.

`generated_at` is to the minute. `Updated {n} h ago` (past 12 hours) is measured from when the app
last fetched, a 304 included. Fixtures `widget.get`, `widget.get.no-guild`. Errors: 401 (the
widget shows its signed-out state).

### POST /v1/me/live-activities and DELETE /v1/me/live-activities/{id}
Auth session or device; 30 registrations a minute per account, at most 20 kept (the oldest goes).
Body `{ kind ("raid"|"sim"), platform ("ios"|"android"), push_token, push_to_start_token?,
object_id }`: `push_token` is the ActivityKit update token (hex, 64 to 200) or the FCM token,
`push_to_start_token` iOS only, `object_id` the raid id or the server sim's id. Answer **201**
`{ id }` (`live.register`). Registering the same token for the same object again refreshes it.
A registration ends with its object (a raid's end plus 15 minutes; a sim after 6 hours), when its
token fails, or on `DELETE` (**204** whether or not it was there; only your own). Tokens are never
logged or answered back. Errors: 400 `invalid` (not JSON), 422 `invalid` with `fields` (`kind`,
`platform`, `push_token`, `push_to_start_token`, `object_id`), 402 `premium_required` for a sim
without `server_sims` (`live.register.error-premium-required`), 404 for a raid you may not see,
one that has ended, or a sim that is not yours, 429.

### Live-surface pushes (W3-C14)
Built by `api/internal/livesurface/payloads.go` and pinned to their fixtures exactly. Events:
`widget.refresh` (silent, collapses per user), `live.raid.start`, `live.raid.update`,
`live.raid.end`, `live.sim.update`, `live.sim.end`.
- **APNs Live Activity** (`apns-push-type: liveactivity`): `aps { timestamp, event
  ("start"|"update"|"end"), content-state, stale-date?, dismissal-date? }`; a push-to-start adds
  `attributes-type` (`RaidActivityAttributes`|`SimActivityAttributes`), `attributes` and `alert`.
  The raid's attributes are `{ raid_id, guild_id, guild_name, zone_name, character_name }` and its
  content-state `{ starts_at, invite_at, status, my_state, role, signed_up, size }`; a sim's
  `{ sim_id, title }` and `{ state ("running"|"finished"|"failed"), done, total, eta_sec, dps }`.
  **Times are Unix seconds** (ActivityKit decodes a Date as seconds since 2001).
  Fixtures `surface.live-raid-start.apns`, `surface.live-sim-update.apns`.
- **APNs widget refresh** (`apns-push-type: background`): `{ aps { content-available: 1 }, kind }`
  (`surface.widget-refresh.apns`).
- **FCM**: a data message `{ message { token, data { kind, ... }, android { priority,
  collapse_key } } }`; `data` is the attributes and state above as strings, and `collapse_key` is
  `widget.refresh`, `live-raid-<id>` or `live-sim-<id>`. Android draws the ongoing notification
  itself. Fixtures `surface.*.fcm`.

**Pending.** No live-surface push is sent yet. Lane A2's senders (`api/internal/push`) send alert
pushes only; a Live Activity push needs `apns-push-type: liveactivity` and its own topic, and
Apple's APNs key does not exist yet. So iOS Live Activities and the Android ongoing notification
are started and updated by the app while it runs, and widgets tick on their timelines; the
registrations are kept and swept so the senders can use them the day they exist.

## Not in v1

No in-app purchase (decision P). Not read: rankings, tier lists, guides (static site pages), the
report per-fight files and coach, `POST /v1/sims/run` (the app reruns, it does not compose a sim),
loot awards, and the roster approve and remove routes (the officer roster actions stay on the web).
If a later lane needs one, it adds its section here and its fixtures in the same commit.

## How the tests pin this

**`api/openapi.yaml` is the source of truth.** Every route and static file above is an operation
there (tagged `app`, `app-static` or `app-browser`), and every fixture in
`docs/contracts/fixtures/app/` is claimed by an `x-fixtures` list on its response (or on a schema,
for the push payloads). `api/internal/contract/openapi_test.go` validates each fixture against its
response schema with libopenapi-validator, and fails on any fixture key the spec does not declare
and on any fixture no response claims. The apps' models, endpoints and fixture route tables are
generated from the spec by `web/scripts/gen-mobile-api.mjs` (CI's `mobile-codegen` workflow fails
on drift). To change a shape: edit the spec and the fixture together, run the generator, commit.
Every `app` operation lists the shared `X-App` and `X-App-Install` header parameters and the shared
`426` `UpdateRequired` response (fixture `error-update-required`); a test holds every operation to
that. The headers are sent by each app's transport, never by the generated endpoints.


`api/internal/contract` compares a response with a fixture (fields, kinds, nullability, never
values or counts; nulls match any kind; `Optional` marks an omitted-when-empty field; `Open` marks
a map whose keys are data). Each `app_contract_test.go` in `auth`, `builds`, `guilds`, `rankings`,
`reports` and `sims` serves the real handler with test data and asserts its answer, and its
`Cache-Control`, against the fixtures above. `cdn_test.go` holds the static fixtures to the real
generator output for the active build, and loads the BiS fixture through the API's own reader.
Run them with `TEST_DATABASE_URL` set and `go test -p 1`, since the packages share one database.
