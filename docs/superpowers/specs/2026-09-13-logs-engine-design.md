# Forever Sixty — combat log engine design

Date: 2026-09-13. Status: approved in conversation (the user set the bar: match or beat Warcraft
Logs' report experience on day one; storage decision taken by the assistant as recommended).
This is the engine only: parser, model, fight detection, analyses, storage layout, and a CLI.
Upload, accounts, report pages, rankings, and the companion are Phase 3 and 4 work built on it.

## Goal

Turn `WoWCombatLog.txt` from the Forever client into a complete, queryable record of every fight
in it, with the analyses a raid leader opens first already computed: damage and healing done and
taken by source, target, and ability; deaths with their last events; buffs and debuffs with
uptimes; casts, interrupts, and dispels; resources; per-second timelines. Fidelity is the whole
point: every event type the format defines is parsed into typed fields, nothing is dropped, and a
report can be rebuilt from its stored events alone.

## Decisions

| Question | Decision |
|---|---|
| Scope | Full fidelity from the start, not a damage meter. |
| Storage | Parquet per report in object storage as the source of truth, precomputed fight summaries beside it; DuckDB embedded in the Go service for per-report queries; self-hosted ClickHouse on one small VM in Phase 4 for rankings and history, reading the same files. |
| Location | New top-level `logs/` Go module: library, CLI, tests, later a Cloud Run job. Shares nothing with `api/` today. |
| Fixtures | Public sample logs from open-source parser repositories plus hand-written lines from the wiki format reference; a real Forever dungeon log from the beta when available. |
| Streaming | Live logging is a first-class mode: the engine is an incremental session fed chunks as the game writes them, and batch parsing is the same session fed one whole file. |
| Client format | Classic-family advanced combat logging (the same line format Classic Era and SoD write); `COMBAT_LOG_VERSION` and `ADVANCED_LOG_ENABLED` headers are parsed and the version is stored with the report. |

## Architecture

```
WoWCombatLog.txt ──▶ reader (streaming, tolerant) ──▶ line lexer ──▶ event decoder (typed)
        │                                                         │
        ▼                                                         ▼
  report metadata                                   unit registry (GUID → name, flags, class, owner)
                                                                  │
                                            fight detector (encounters, pulls, trash segments)
                                                                  │
                        ┌─────────────────────────────────────────┴───────────────────┐
                        ▼                                                             ▼
        events.parquet per fight (typed columns)                       summaries.json per fight
        (source of truth, replayable)                                  (tables the report page opens first)
                        │
                        ▼
        DuckDB on demand: any query over a report's fight files (Phase 3 API)
        ClickHouse (Phase 4): ingest the same Parquet for rankings and history
```

- **Streaming.** Files are hundreds of megabytes; the reader never loads a file into memory, and
  memory use is bounded by the current fight's events plus registries.
- **Tolerant.** A malformed line is recorded as a `parse_error` event with the raw text and the
  reason; parsing continues. A report's health (lines, errors, unknown events) is part of its
  metadata.
- **Deterministic.** The same file yields byte-identical Parquet and JSON, tested.

## Streaming

Warcraft Logs' live logging is table stakes for raid nights: the uploader tails
`WoWCombatLog.txt` while the game writes it, and the report updates as fights end. The engine is
therefore built as an incremental session, and batch parsing is a special case of it.

- **Session.** `logs.NewSession(reportID, opts)` holds all parser state: the partial trailing
  line, the unit registry, the open fight, health counters, the year and last timestamp for
  rollover handling. `session.Feed(chunk []byte)` decodes every complete line in the chunk and
  keeps the remainder; it returns the events decoded and any fights that closed. `session.Close()`
  closes an open fight as unfinished and flushes. Feeding a whole file in one call is batch mode;
  the CLI does exactly that.
- **Fight lifecycle.** A fight opens on its first event, accumulates events in memory, and on
  close writes `events.parquet` and `summary.json`. While open, `session.Snapshot()` returns the
  in-progress summary (damage, healing, deaths so far, elapsed time) cheaply, computed from
  running accumulators rather than a rescan, so a live report page can refresh every few seconds.
- **Chunk boundaries** never matter: lines split across chunks, chunks that end mid-quote, and
  chunks containing a year rollover all parse identically to the whole file, tested by feeding
  every fixture in random chunk sizes and comparing outputs byte for byte with the batch result.
- **Ordering and idempotency.** Each fed byte range is addressed by file offset; a re-fed range
  (the uploader retried) is detected by offset and ignored, so at-least-once delivery from the
  companion is safe. Offsets are recorded in `report.json` so a session can resume after a
  process restart by re-reading the stored tail state (`session.State()` serializes it).
- **Runtime shape (Phase 3).** The companion posts chunks to an ingest endpoint with the report
  id and offset; one session per live report runs in a long-lived Cloud Run instance with
  session affinity, or a small always-on VM, writing fight files to the bucket as fights close.
  The engine exposes only the session API; the transport is the Phase 3 design's job.

## Format coverage

Every line is `MM/DD HH:MM:SS.mmm  EVENT,params...` with quoted strings, nested brackets for
advanced fields, and a date without a year (the year comes from file metadata or upload time and
the reader handles a year rollover mid-file). The decoder covers:

- **Header lines**: `COMBAT_LOG_VERSION`, `ADVANCED_LOG_ENABLED`, `BUILD_VERSION`, `PROJECT_ID`.
- **Base fields** on every combat event: source and dest GUID, name, flags, raid flags.
- **Prefixes**: `SWING`, `RANGE`, `SPELL`, `SPELL_PERIODIC`, `SPELL_BUILDING`, `ENVIRONMENTAL`,
  with their spell id, name, and school fields.
- **Suffixes**: `_DAMAGE`, `_MISSED`, `_HEAL`, `_HEAL_ABSORBED`, `_ABSORBED`, `_ENERGIZE`,
  `_DRAIN`, `_LEECH`, `_INTERRUPT`, `_DISPEL`, `_DISPEL_FAILED`, `_STOLEN`, `_EXTRA_ATTACKS`,
  `_AURA_APPLIED`, `_AURA_REMOVED`, `_AURA_APPLIED_DOSE`, `_AURA_REMOVED_DOSE`,
  `_AURA_REFRESH`, `_AURA_BROKEN`, `_AURA_BROKEN_SPELL`, `_CAST_START`, `_CAST_SUCCESS`,
  `_CAST_FAILED`, `_INSTAKILL`, `_DURABILITY_DAMAGE`, `_DURABILITY_DAMAGE_ALL`, `_CREATE`,
  `_SUMMON`, `_RESURRECT`, `_EMPOWER_*`.
- **Advanced fields** (present when advanced logging is on): info GUID, owner GUID, current and
  max HP, attack and spell power, armor, absorb, power type, current and max power, power cost,
  position x and y, map id, facing, item level.
- **Special events**: `ENCOUNTER_START`, `ENCOUNTER_END`, `COMBATANT_INFO` (gear, talents,
  auras at pull), `UNIT_DIED`, `UNIT_DESTROYED`, `UNIT_DISSIPATES`, `PARTY_KILL`,
  `ZONE_CHANGE`, `MAP_CHANGE`, `WORLD_MARKER_PLACED`, `WORLD_MARKER_REMOVED`, `EMOTE`,
  `SPELL_ABSORBED` in both its shapes.
- **Unknown events** are kept with their raw params as a string array, counted, and reported.

The wiki's `COMBAT_LOG_EVENT` reference and the community field references are the source; the
plan pins the field positions with a fixture per event shape.

## Data model

**Units.** GUID parsed into kind (Player, Creature, Pet, Vehicle, GameObject, Item), server id,
npc id or player id, spawn id. The registry records first-seen name, flags (affiliation, reaction,
control, type), class when derivable (COMBATANT_INFO, spell school heuristics as fallback, labeled
as inferred), and pet owner from `_SUMMON` and the advanced owner GUID. Players are keyed by
GUID; names are display data.

**Fights.** An encounter is `ENCOUNTER_START` to `ENCOUNTER_END` with its id, name, difficulty,
group size, and success. Pulls without encounter events (dungeons, trash) are segmented by
combat activity gaps: a fight starts at the first hostile damage after idle and ends after a gap
with no hostile events longer than a threshold, or at a `UNIT_DIED` of every hostile. Each fight
gets an index, start and end offsets, duration, a kill flag, and the list of players present
(any player event inside it).

**Events table** (one Parquet file per fight, sorted by timestamp then line number):

| column | type |
|---|---|
| ts_ms | int64, milliseconds from report start |
| line | int64 |
| event | dictionary string |
| prefix, suffix | dictionary strings |
| source_guid, source_name, source_flags, source_raid_flags | string, string, uint32, uint32 |
| dest_guid, dest_name, dest_flags, dest_raid_flags | same |
| spell_id, spell_name, spell_school | int32, string, uint8 |
| amount, overkill, over_heal, absorbed, blocked, resisted, critical, glancing, crushing, is_offhand | numeric and bool as appropriate |
| miss_type, aura_type, aura_stacks, power_type, extra_spell_id, extra_spell_name, extra_school | as appropriate |
| adv_* | the advanced fields, nullable |
| raw | string, null unless the event is unknown or errored |

Nullable columns are null where a suffix has no such field; there is no overloading of columns
across event kinds beyond what the format itself overloads.

**Summaries** (`summary.json` per fight, and `report.json` at the report level):
- damage done and taken: per source, per ability, per target, with hits, crits, misses by type,
  amount, absorbed, overkill, and per-second series at one-second resolution
- healing done and taken: per source, per ability, per target, with amount, overheal, absorbs
- deaths: per player, timestamp, killing blow, the last 10 damage and heal events before it and
  the last aura applications, with HP from the advanced fields when present
- buffs and debuffs: per unit, per aura, uptime segments and percentage, stacks over time
- casts: per source, per spell, count, and cast time when start and success pair
- interrupts and dispels: per source, what was interrupted or dispelled, on whom
- resources: per player, power type series from the advanced fields
- combatant info: gear with item ids and enchants, talents, at pull, per player
- players: GUID, name, class (with an `inferred` flag), spec guess from talents when present

`report.json`: file metadata, log version, build, health counters, zone and map changes, fight
list with the fields above, and the units registry.

## Storage layout

```
reports/<report-id>/report.json
reports/<report-id>/fights/<n>/events.parquet
reports/<report-id>/fights/<n>/summary.json
reports/<report-id>/raw.txt.zst            (the upload, compressed, kept for reparse)
```
`<report-id>` is assigned by the uploader in Phase 3; the engine takes it as an argument. The
CLI writes the same layout to a local directory. Object storage is Google Cloud Storage, in the
existing project, via the standard client; the writer is an interface with a local implementation
used by tests and the CLI.

## CLI

`forever-logs tail <file> --out <dir>` follows a growing file and writes fights as they close, printing each fight's summary line, which is the live mode exercised locally against the game client. `forever-logs parse <file> --out <dir> [--report-id <id>] [--year 2026]` runs the whole
pipeline and prints the health summary. `forever-logs fights <file>` lists fights. `forever-logs
query <dir> "<sql>"` runs DuckDB over a report's Parquet for inspection. The CLI is how the engine
is validated on real logs before any upload path exists.

## Performance targets

A 500 MB raid-night log parses in under 60 seconds on one Cloud Run CPU with under 1 GB of
memory; per-fight Parquet plus summaries total under 20% of the raw size; a summary for any fight
is served from JSON without touching Parquet.

## Error handling

Malformed lines become `parse_error` events and health counters, never a stopped parse. Unknown
events are preserved raw. A file that is not a combat log at all (no header, no recognizable
event in the first 1,000 lines) is rejected with a message. Year rollover, duplicate lines from
log concatenation, and timestamps going backwards (client clock changes) are handled and counted.

## Testing

- Lexer tests for quoting, escaping, nested brackets, and every timestamp shape.
- A fixture line per event shape from the format reference, with the expected typed event; a
  golden test over a concatenated fixture file producing the Parquet and JSON outputs.
- Public sample logs from open-source parser repositories checked in under `logs/testdata/` with
  their licenses, used for fight detection, summaries, and the performance benchmark.
- Property tests: parse then serialize then parse yields the same events; summaries recomputed
  from Parquet match the ones written at parse time.
- A benchmark over the largest fixture with the memory ceiling asserted.
- Streaming equivalence: every fixture fed in chunk sizes of 1, 7, 64, 4096 bytes and random sizes yields byte-identical outputs to batch; duplicate and overlapping chunk ranges are ignored; a session serialized mid-fight and restored continues to the same result.
- Snapshot tests: the in-progress summary after N events equals the summary of a fight truncated at N.
- Coverage floor 80%, `go vet`, `gofmt`, race detector in CI.

## Out of scope here

Upload endpoint and the companion; report pages; accounts and ownership; rankings and
ClickHouse; replay and positional maps beyond storing positions; Retail log variants.
