# A verified layout for combat log version 22

Status: approved in conversation 2026-09-16 (the user asked for this as a lane parallel to
the closing-the-gaps plan). Written from measurements on 89 real v22 logs.

## Why

The engine's only verified layout is `retail-v16`, written against the public WoWP sample.
The World of Warcraft 20th Anniversary Edition clients (Classic Fresh / TBC Anniversary on
the modern client) emit `COMBAT_LOG_VERSION,22`, and Forever's client is expected to do the
same. On a v22 log today the engine falls back to the inferred, unverified layout: on two
real files 16% and 26% of lines are parse errors, and a dozen event types the engine knows on
v16 (absorbs, deaths, party kills, dispels, combatant info, map changes) are reported unknown
because the inferred layout carries no specials. A v22 log therefore produces a report that is
wrong in every table.

## What is measured (89 files, builds 12.0.5 to 12.1.0, all `ADVANCED_LOG_ENABLED,1`)

Header: `COMBAT_LOG_VERSION,22,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,12.0.7,PROJECT_ID,1`
(8 fields, as v16).

Timestamp: `7/1/2026 09:07:12.137-5` — a year in the date and a UTC offset in hours after the
clock. The stamp parser already supports both (`StampYear`, `StampZone`); the v22 layout sets
both true.

Field counts per event, v16 → v22 (dominant widths):

| Event | v16 | v22 | Reading |
|---|---|---|---|
| SPELL_CAST_SUCCESS | 29 | 31 | advanced block 17 → 19 |
| SWING_DAMAGE, SWING_DAMAGE_LANDED | 36 | 38 | advanced +2, damage suffix stays 10 |
| SPELL_DAMAGE, SPELL_PERIODIC_DAMAGE | 39 | 42 | advanced +2, damage suffix 10 → 11 |
| SPELL_HEAL, SPELL_PERIODIC_HEAL | 34 | 36 | advanced +2, heal suffix stays 5 |
| SPELL_ENERGIZE, SPELL_PERIODIC_ENERGIZE | 33 | 35 | advanced +2, suffix stays 4 |
| SPELL_MISSED, SPELL_PERIODIC_MISSED | 14 / 17 | 15 / 18 | missed suffix 2 → 3, absorb extra stays 3 |
| SWING_MISSED | 11 / 14 | 11 / 14 | unchanged |
| SPELL_AURA_APPLIED / REMOVED / REFRESH | 13 (14 with amount) | 13 (14) | unchanged |
| SPELL_AURA_*_DOSE | 14 | 14 | unchanged |
| SPELL_CAST_START, SPELL_SUMMON | 12 | 12 | unchanged |
| SPELL_ABSORBED | 19 / 22 | 19 / 22 | unchanged |
| UNIT_DIED, PARTY_KILL | 10 | 10 | unchanged |
| SPELL_EXTRA_ATTACKS, SPELL_CAST_FAILED | 13 | 13 | unchanged |
| COMBATANT_INFO | variable (34 base) | variable | keep the v16 combatant parser; verify on the corpus |

The identity of the two new advanced fields and the one new spell-damage field is to be
determined from the lines themselves during the plan (compare a v16 and a v22 SPELL_DAMAGE
line field by field); the layout must name them, not skip them blindly.

New event types v22 emits that v16 never had, with counts over the two probe files:
`DAMAGE_SPLIT` (1,571), `SPELL_DAMAGE_SUPPORT` (100), `SPELL_PERIODIC_DAMAGE_SUPPORT` (17),
`SPELL_HEAL_SUPPORT` (3), `SPELL_EMPOWER_START` (15), `SPELL_EMPOWER_END` (14),
`SPELL_EMPOWER_INTERRUPT` (2), `ARENA_MATCH_START` (3), `ARENA_MATCH_END` (2),
`ENVIRONMENTAL_DAMAGE` (5, width to verify against v16's 37), `SPELL_AURA_BROKEN_SPELL` (10),
`SPELL_DRAIN` (9), `SPELL_DISPEL` (170), `MAP_CHANGE` (6).

## Design

- A second verified layout, `retail-v22`, beside `retail-v16` in `logs/engine/layout/retail.go`:
  `Version: 22`, `StampYear: true`, `StampZone: true`, `Advanced: 19`, `_DAMAGE` params 11 for
  the SPELL prefixes and 10 for SWING (or one rule if the extra field proves to be a prefix
  field rather than a suffix field), `_MISSED` params 3, the same specials as v16, and
  entries for the new events: `DAMAGE_SPLIT` and the `_SUPPORT` variants parsed as damage or
  heal lines whose amounts are attributed to the supporting player (the summary decides
  whether to fold them; the layout only has to read them), `SPELL_EMPOWER_*` as cast-like
  lines with a stage field, `ARENA_MATCH_START/END` as specials the fight splitter may use as
  boundaries later (not in this plan).
- Layout selection by header version picks v22 as it picks v16 today; the inferred fallback
  is untouched and still catches versions we have not verified.
- The engine version does not change: no summary field is added or renamed, and a v16 log
  parses byte-for-byte as before (the v16 golden and the fixture are unchanged). The
  closing-the-gaps lane owns `session.go` and the fixture; this lane never touches them.
- Test data: a committed excerpt `logs/engine/event/testdata/v22.log` of a few thousand
  lines cut from one of the user's own arena logs (real player names; the user owns the
  logs and has accepted that), covering every event type in the table above; a
  `v22.summary.json.golden` for the summary package built from that excerpt. The full
  89-file corpus stays in the session scratchpad
  (`/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/logs/v22-all/`)
  and is the acceptance sweep, never committed.
- `internal/sample` stays pointed at the public v16 sample; the v22 excerpt is plain testdata.

## Acceptance

- `forever-logs conformance` over the 89-file corpus: `layout=retail-v22 verified=true` on
  every file, `parse_errors=0`, and `unknown` empty or limited to events named in a ledger
  ruling with the count and the reason.
- The layout test that checks v16 widths against the verified counts gains a v22 twin.
- `go test ./...` green in `logs/`; the v16 golden is byte-identical before and after.
- The two probe files parse to 5 and 19 fights as today with zero errors, and `fights`
  reports player counts that match COMBATANT_INFO (6 and 6 in the arena files).

## Out of scope

Fight boundaries from ARENA_MATCH_*; a PvP report view; anonymising the excerpt; the
Anniversary client's own idiosyncrasies until a log from it exists (the user will record one).
