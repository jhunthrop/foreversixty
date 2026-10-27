# Guild control centre: API contract (binding for the api and web lanes)

Spec: `design/specs/2026-10-04-guild-page.md` (v2). Every endpoint lives under
`/v1/guilds/{id}` and uses the standard envelope `{ ok, data, error, request_id }`.
Viewer roles: `public` (no session or not a member), `member` (verified member),
`officer` (verified officer or leader), `moderator`. Fields marked *officer* are omitted for
other roles, never sent empty. All timestamps are RFC 3339 UTC. All numbers that are
unknown are `null`, never `0`.

## GET /v1/guilds/{id}/home  (existing, extended)

`guild.faction` is new (migration 0031, `guilds.RecomputeFaction` - api/internal/guilds/
faction.go): `"alliance"` or `"horde"`, the majority faction among the guild's exported
roster, or `null` when it has never resolved to either (no exports yet, a tie, or every
known character neutral/unknown-race). Recomputed on every membership change
(export sync, approve, remove, leave, invite accept) and backfilled for any
pre-migration guild within one `MembershipJob` sweep tick (5 minutes).

```json
{
  "guild": { "id": 2, "name": "OLYMPUS XXVII", "region": "us", "ruleset": "pvp", "faction": "alliance" },
  "viewer": { "role": "officer", "character_key": "us/pvp/obnoxious-yell", "verified": true },
  "claim": { "state": "claimed", "since": "...", "frozen": false, "claimed_by_name": "Obnoxious Yell" },
  "summary": {
    "raiders": 25, "waiting_for_approval": 3, "below_rating_floor": 2,
    "named_encounters_down": 1, "pulls_this_tier": 27,
    "updated_at": "..."
  },
  "standing": {                       // member and officer only; null when not derivable
    "spec": "Fury", "class": "warrior",
    "same_spec_count": 3, "rank_by_item_level": 2, "item_level": 66,
    "needs_before_next_raid": ["Enchant chest", "Spend 1 talent point"]   // from readiness, max 3
  },
  "reports": [ ...ReportSummary (last 7 days, existing shape plus "zone" and "duration_ms") ],
  "roster": [ ...RosterRow ],
  "pending": [ ...RosterRow ],         // officer: unverified members; others: []
  "addon_snapshot": {                  // the guild's newest addon roster snapshot; null when it has none
    "received_at": "2026-11-04T20:58:00Z", "by_character": "us/pvp/thorgrim", "members": 14
  }
}
```

`addon_snapshot` is the newest roster snapshot a member's addon sent for this guild
(`POST /v1/addon/exports`' outbox, `docs/contracts/addon-v2.md` §4, stored in
`guild_roster_snapshots`): `received_at` is when the site received it (RFC 3339 UTC),
`by_character` the character key whose export carried it, and `members` how many member lines
it listed, the ones that match no roster row included. It lets the page say "Roster from the
addon, 3 minutes ago, by Thorgrim". The newest snapshot wins even when it lists nobody (an
empty roster is a fact): `members` is then 0 and every row's `addon` is null. Like the rest of
the home, only a verified member reads it.

### RosterRow

```json
{
  "character_key": "us/pvp/windtalon", "name": "Windtalon", "region": "us", "ruleset": "pvp",
  "class": "hunter", "spec": "Marksmanship", "role": "dps",
  "rank": "member", "verified": true, "consent": "gear",
  "item_level": 70,                     // null without gear consent
  "logged_at": "2026-10-04T21:10:00Z",  // addon_exports.updated_at; replaces logged_recently (keep logged_recently too)
  "logged_recently": true,
  "attendance": { "present": 6, "nights": 8 },   // last 8 non-private, unflagged guild reports someone is present in
  "best_parse": { "metric": "dps", "value": 376.4, "percentile": null, "encounter": "Onyxia", "report_id": "...", "fight_index": 3 },
  "rating": { "overall": 65, "output": 70, "survival": 60, "mechanics": 62, "utility": 58, "preparation": 71, "activity": 66 },  // null when no rating_scores
  "professions": ["Leatherworking", "Skinning"],
  "account_key": "u:8812",             // same value for a main and its alts
  "last_report_at": "...",            // newest report this character's account uploaded, null if none
  "may_remove": false, "may_approve": false,
  "addon": {                           // this character's line in addon_snapshot; null when not in it
    "level": 58, "spec": "hunter-marksmanship", "gear_gap": 6, "build": "MM raid",   // all four null without gear consent; build also null when none announced
    "seen_at": "2026-11-04 14:01",     // the sender's local YYYY-MM-DD HH:MM, as stored
    "snapshot_at": "2026-11-04T20:58:00Z"   // addon_snapshot.received_at
  }
}
```

`addon` is present on every row (in `pending` too) and is `null` when the guild has no
snapshot or the newest one does not carry this character. A snapshot line is matched to a row
by character name, any case: the snapshot's names carry no realm (a `Name-Realm` line drops
its realm half), so the guild's realm is the row's. The first line of a name wins. A line that
matches no row is never added as one; it counts only in `addon_snapshot.members`. `level`,
`spec` (a spec slug), `gear_gap` and `build` are as that member's addon announced them
(contract 2's `hello`); the snapshot carries only the members its sender's guild tab had live
data for, so a member that addon had no data for reads `null`. Those four fields are also `null`
when the row's member did not give `gear` (or `gear_bags`) consent, the same rule that withholds
the row's class, spec and item level; `seen_at` and `snapshot_at` stay and `addon` stays non-null.

`supporter` is present on every `roster` and `pending` row: true when the row's owning account
shows the supporter mark, false otherwise. `members` entries, which are names without an
account, never carry it.

A character is present in a report when it has a ranked or rated row there (`fight_metrics`
not removed, or `rating_scores`), read by character key, never the log's GUIDs in
`fights.players`. `attendance` counts the guild's last 8 reports that are non-private,
unflagged (not tampered or unverified) and carry at least one such row, so a night nobody can
be present in (private, flagged, trash-only, or not ranked yet) is never everyone's absence
(task L6.8). Raid mode (lane G2): once a guild has one done raid with attendance rows, the home
carries `attendance_source: "raids"` (else `"reports"`), and `attendance` becomes raids
attended over the raids this character signed in or late for, or attended, among the guild's
last 8 logged raids; `attendance_30` (raid mode only) is the same over 30. `nights` is then per
character, so a consumer reads its own row's; an excused raid is neither attended nor a night.
`GET /v1/me`'s `characters[].guild` gains the same `attendance` in raid mode and omits it
otherwise. Verification by logs (`verified_by: "logs"`) and the moderator's
`nights_in_others_reports` read the same presence in non-private, unflagged reports only, so
they count unflagged, non-private boss fights: the same evidence the Forever Score trusts. The
guild's readers of rows a player sees (attendance, best parse, rating, class and spec, top
parses, best per role) and every score input pick their reports by one rule
(`presence.EvidenceReport`: not private, not flagged), so a flagged report counts toward no score
either (task L6.10). The guild's own counts (the summary's `named_encounters_down` and
`pulls_this_tier`, the Loot tab's `killed`, and progression's counts for a verified member)
include its private nights but never a flagged report (`presence.NotFlagged`, task L6.12). Home
and Loot open to a verified member only: an account whose only rows in the guild are unverified
gets 403, like any other non-member.

## GET /v1/guilds/{id}/raids?limit=20&cursor=

```json
{ "rows": [ {
  "id": "...", "title": "Onyxia · Wednesday", "zone": "Onyxia's Lair", "created_at": "...",
  "duration_ms": 5400000, "fight_count": 6, "kill_count": 5, "wipe_count": 1,
  "raiders": 20, "deaths": 7,
  "top_parse": { "name": "Pyrewisp", "class": "mage", "metric": "dps", "value": 358.2 },
  "fights": [ { "index": 0, "name": "Onyxia", "encounter_id": 1084, "kill": true, "duration_ms": 420000, "deaths": 1, "players": 20 } ],
  "present": [ { "character_key": "...", "name": "...", "class": "..." } ]
} ], "next_cursor": null }
```
Visibility follows the existing report rules for the viewer, and a flagged report (tampered or
unverified) is listed, row and fights, only to a verified member who owns it: everyone else,
signed out or not, sees the public reports that are not flagged (`presence.EvidenceReport`) and
their own unflagged ones, and a verified member the guild's public and guild-visibility reports
that are not flagged (`presence.NotFlagged`) and their own. `present` (and `raiders`, its
count) is read by character key from the same presence as `attendance`, so a private,
flagged or trash-only night lists nobody; `fights[].players` stays the count of units in the log.

## GET /v1/guilds/{id}/progression

```json
{ "tier": { "name": "First tier", "raids": ["Barrow Deeps", "Hyjal Summit", "Onyxia's Lair"], "named_encounters": 1, "down": 1 },
  "encounters": [ {
    "encounter_id": 1084, "name": "Onyxia", "zone": "Onyxia's Lair",
    "first_kill_at": "...", "pulls": 12, "kills": 2, "best_kill_ms": 380000,
    "pulls_by_night": [ { "report_id": "...", "date": "2026-10-01", "pulls": 7, "killed": true } ],
    "deaths_per_pull": 1.4,
    "best_by_role": { "tank": {...ParseRef}, "healer": {...ParseRef}, "dps": {...ParseRef} }
  } ],
  "unnamed": [ { "zone": "Barrow Deeps", "pulls": 7, "nights": 1 } ] }
```
`ParseRef` = `{ "name", "class", "spec", "metric", "value", "report_id", "fight_index" }`.

The tab is public (no session needed), and the counts depend on the viewer (task L6.12). A
verified member of the guild reads its own numbers: `pulls`, `kills`, `first_kill_at`,
`best_kill_ms`, `deaths_per_pull`, `pulls_by_night` (with each night's `report_id`),
`unnamed` and `tier.down` count the guild's private nights too, never a flagged report
(`presence.NotFlagged`), and the answer is `Cache-Control: private, no-cache`. Everyone else
(signed out, another account, an account whose rows in the guild are unverified) reads the
public numbers, the ones the public guild page prints: reports that are not private and not
flagged (`presence.EvidenceReport`), in a shareable answer (`public, max-age=30,
stale-while-revalidate=60`). Both answers carry `Vary: Cookie, Authorization`. `best_by_role`
reads `presence.EvidenceReport` for every viewer.

## GET /v1/guilds/{id}/readiness   (member and officer; member sees every row, officer gets `nudge_text`)

```json
{ "median_item_level": 63, "generated_at": "...",
  "rows": [ {
    "character_key": "...", "name": "...", "class": "...", "spec": "...", "consent": "gear",
    "gear_gap": { "upgrades": 5, "gain_dps": 48.0, "not_sim_checked": 1 },   // null without gear consent
    "enchants": { "missing_slots": ["chest"], "checked": true },             // checked=false without gear consent
    "consumables": { "state": "stocked" | "short" | "unknown" },             // unknown without gear_bags consent
    "talent_points_unspent": 1,
    "talent_points_read": true,
    "item_level": 58, "item_level_delta": -5,
    "logged_at": "...",
    "failing": 4,                                                            // count of failing checks
    "nudge_text": "Thornhide: 5 upgrades (+48 DPS), chest enchant missing, 1 talent point unspent"  // officer
  } ] }
```
Gear gap uses the planner's rule (`web/src/lib/home/upgrades.ts`, ported to Go in the api lane):
pick == worn → none; worn in the band's alternatives → gain = −dps_delta; otherwise not sim-checked.
Band = the character's level band and faction from the spec's BiS file in `data/builds/<build>/bis/`.
Enchant "missing" = the slot carries no enchant in the export while the band's BiS lists an enchant
for it; when the BiS file carries no enchant data for the slot, the slot is not checked.

## GET /v1/guilds/{id}/loot?encounter_id=1084   (member read-only, officer full)

```json
{ "encounters": [ { "encounter_id": 1084, "name": "Onyxia", "zone": "Onyxia's Lair", "killed": true } ],
  "selected": 1084,
  "items": [ {
    "item_id": 16955, "name": "Helm of Wrath", "icon": "inv_helmet_71", "quality": 4, "slot": "head",
    "awarded_to": { "character_key": "...", "name": "Grimtotem", "at": "...", "by_name": "Kraggor" } | null,
    "candidates": [ { "character_key": "...", "name": "...", "class": "...", "spec": "...",
                      "gain_dps": 40.0, "not_sim_checked": false, "ilvl_delta": null,
                      "attendance": { "present": 8, "nights": 8 }, "already_equivalent": false } ]
  } ] }
```
Candidates, two tiers:
- **Tier 0 (BiS-matched)**: every roster character whose spec BiS band (level 60 band, their
  faction) names the item as the pick or an alternative for that slot — `gain_dps`/`not_sim_checked`
  from the normal gear-gap rule, `ilvl_delta` always `null`.
- **Tier 1 (fallback)**, added 2026-10-04 (loot candidates must not read empty on real data —
  no BiS file in this build names any raid-tier item, so tier 0 is empty for every item today):
  a character with no tier-0 match is still a candidate when (a) their class can use the item
  at all, read from the build's own per-class item table (`data/builds/<build>/items/<class>.json`
  — presence there is itself the class/armor-type restriction) and (b) the item's own item level
  is higher than what they wear in that slot (from the export), or the slot is empty. A tier-1
  candidate carries `gain_dps: null`, `not_sim_checked: true`, and `ilvl_delta` (item item level
  minus worn item level, or the item's own item level when the slot is empty).

Ranking: every tier-0 candidate ranks ahead of every tier-1 candidate, regardless of either's own
numbers. Within tier 0: `gain_dps` desc (`null` sorts last within the tier). Within tier 1:
`ilvl_delta` desc. Either tier, attendance (`present` over the candidate's own `nights`, then `present`) breaks the remaining tie.

## POST /v1/guilds/{id}/loot/awards  (officer)   body `{ "encounter_id", "item_id", "character_key", "raid_id"? }`
## DELETE /v1/guilds/{id}/loot/awards/{award_id}  (officer)
Table `loot_awards (id bigserial, guild_id, encounter_id, item_id, character_key, awarded_by, awarded_at, report_id null, raid_id null, attached_by_match)`.

`raid_id` is optional and names the done Onyxia raid of this guild the award was for. A
`raid_id` that is not one answers 400 `invalid` ("raid N is not a done Onyxia raid of this
guild"). Without it the award goes to the raid being played now: the guild's Onyxia raid (any
status but cancelled) whose match window, an hour before its start to six hours after its end,
holds the award's time. Otherwise it has no raid until a match attaches it; a previous done raid
never takes it. One award per drop per raid: re-awarding the same drop in the same raid
replaces it, and a later raid's award of the same item is a new row.

Attach on match (Lane G2): an award made live has no raid, since a raid is not done until its
log matches. When a raid's log matches (auto or by hand), the guild's raid-less awards whose
`awarded_at` falls between an hour before the raid's start and six hours after the log's anchor
join that raid, in the match's own transaction. If the raid already holds an award of the same
drop, the newer of the two stays on the raid and the older is left with no raid. Clearing the
match takes back only the awards the match attached (`attached_by_match`); an award filed under
the raid with `raid_id` is never detached.

## Existing officer endpoints stay as they are
Approve, remove, invite rotate, claim, settings, visibility. Add `POST /v1/guilds/{id}/roster/approve-all` (officer) that approves every pending row and returns the approved keys.

## Not in this build (spec §9, recorded)
Death causes (needs event data), recommended-enchant curation beyond what BiS files carry.

## Public guild page shape also gains `guild.faction`

Out of this contract's own scope (`GET /v1/guilds/{region}/{ruleset}/{name}` is
`rankings.Store.Guild`, api/internal/rankings/guilds.go, mounted outside this package) but
the same migration/column: its `guild` object gains the identical `faction` field
("alliance" | "horde" | null), read straight from `guilds.faction`.
