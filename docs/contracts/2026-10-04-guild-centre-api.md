# Guild control centre: API contract (binding for the api and web lanes)

Spec: `design/specs/2026-10-04-guild-page.md` (v2). Every endpoint lives under
`/v1/guilds/{id}` and uses the standard envelope `{ ok, data, error, request_id }`.
Viewer roles: `public` (no session or not a member), `member` (verified member),
`officer` (verified officer or leader), `moderator`. Fields marked *officer* are omitted for
other roles, never sent empty. All timestamps are RFC 3339 UTC. All numbers that are
unknown are `null`, never `0`.

## GET /v1/guilds/{id}/home  (existing, extended)

```json
{
  "guild": { "id": 2, "name": "OLYMPUS XXVII", "region": "us", "ruleset": "pvp" },
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
  "pending": [ ...RosterRow ]          // officer: unverified members; others: []
}
```

### RosterRow

```json
{
  "character_key": "us/pvp/windtalon", "name": "Windtalon", "region": "us", "ruleset": "pvp",
  "class": "hunter", "spec": "Marksmanship", "role": "dps",
  "rank": "member", "verified": true, "consent": "gear",
  "item_level": 70,                     // null without gear consent
  "logged_at": "2026-10-04T21:10:00Z",  // addon_exports.updated_at; replaces logged_recently (keep logged_recently too)
  "logged_recently": true,
  "attendance": { "present": 6, "nights": 8 },   // last 8 guild reports
  "best_parse": { "metric": "dps", "value": 376.4, "percentile": null, "encounter": "Onyxia", "report_id": "...", "fight_index": 3 },
  "rating": { "overall": 65, "output": 70, "survival": 60, "mechanics": 62, "utility": 58, "preparation": 71, "activity": 66 },  // null when no rating_scores
  "professions": ["Leatherworking", "Skinning"],
  "account_key": "u:8812",             // same value for a main and its alts
  "last_report_at": "...",            // newest report this character's account uploaded, null if none
  "may_remove": false, "may_approve": false
}
```

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
Visibility follows the existing report rules for the viewer.

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

## GET /v1/guilds/{id}/readiness   (member and officer; member sees every row, officer gets `nudge_text`)

```json
{ "median_item_level": 63, "generated_at": "...",
  "rows": [ {
    "character_key": "...", "name": "...", "class": "...", "spec": "...", "consent": "gear",
    "gear_gap": { "upgrades": 5, "gain_dps": 48.0, "not_sim_checked": 1 },   // null without gear consent
    "enchants": { "missing_slots": ["chest"], "checked": true },             // checked=false without gear consent
    "consumables": { "state": "stocked" | "short" | "unknown" },             // unknown without gear_bags consent
    "talent_points_unspent": 1,
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
                      "gain_dps": 40.0, "not_sim_checked": false, "attendance": { "present": 8, "nights": 8 },
                      "already_equivalent": false } ]
  } ] }
```
Candidates: every roster character whose spec BiS band (level 60 band, their faction) names the
item as the pick or an alternative for that slot, ranked by `gain_dps` desc, then attendance.

## POST /v1/guilds/{id}/loot/awards  (officer)   body `{ "encounter_id", "item_id", "character_key" }`
## DELETE /v1/guilds/{id}/loot/awards/{award_id}  (officer)
Table `loot_awards (id bigserial, guild_id, encounter_id, item_id, character_key, awarded_by, awarded_at, report_id null)`.

## Existing officer endpoints stay as they are
Approve, remove, invite rotate, claim, settings, visibility. Add `POST /v1/guilds/{id}/roster/approve-all` (officer) that approves every pending row and returns the approved keys.

## Not in this build (spec §9, recorded)
Death causes (needs event data), recommended-enchant curation beyond what BiS files carry.
