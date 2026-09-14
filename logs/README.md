# logs — the combat log engine

One engine, three hosts: the desktop companion, the server-side parse job,
and the `forever-logs` command line. All three call `engine/session`.

Design: [`docs/superpowers/specs/2026-09-13-logs-engine-design.md`](../docs/superpowers/specs/2026-09-13-logs-engine-design.md).

```
lexer  →  layout  →  event  →  units  →  fight  →  summary  →  parquet  →  store
                                   └──────────── session ────────────┘
```

- **`engine/lexer`** splits a byte stream into timestamped parameter slices, carrying a partial line across chunk boundaries and ignoring bytes it has already seen.
- **`engine/layout`** is the format table: one row per client dialect.
- **`engine/event`** decodes parameters into one flat typed struct.
- **`engine/units`** parses GUIDs and tracks who is who.
- **`engine/fight`** splits the log into encounters and trash.
- **`engine/summary`** computes every table the report page opens.
- **`engine/parquet`** writes the per-fight events file the browser queries.
- **`engine/session`** is the public API: `New`, `Feed`, `Snapshot`, `State`, `Restore`, `Close`.
- **`engine/store`** writes the R2 key layout through a one-method interface.

## Running it

```bash
cd logs
go test ./... -race -cover
go run ./cmd/forever-logs fights path/to/WoWCombatLog.txt
go run ./cmd/forever-logs parse -out /tmp/report -report abc123 path/to/WoWCombatLog.txt
go run ./cmd/forever-logs tail path/to/WoWCombatLog.txt
go run ./cmd/forever-logs conformance path/to/logs/
```

`query` is deliberately absent: deep queries run in the browser over the
fight's Parquet file, per the design's read path.

## The format strategy

Blizzard has never documented the combat log, and the format differs between
retail and Classic and between patches. Rather than branch on a version
number in the decoder, the engine keeps a **table of layout rows** in
`engine/layout`. A row says how many parameters each suffix carries, whether
the event carries the seventeen-field advanced block, which widths a special
event may have, where `COMBATANT_INFO`'s parts sit, and whether the timestamp
carries a year and a zone. The decoder reads the row; it never sniffs.

Three rows ship today:

| Row | Version | Verified | Where it comes from |
|---|---|---|---|
| `retail-v16` | 16, project 1 | yes | every count checked against a real 77 MB retail log: 272,367 events, 105 `COMBATANT_INFO` lines, eighteen encounters |
| `classic-wiki` | — | no | the parameter lists warcraft.wiki.gg documents for `COMBAT_LOG_EVENT`, without the retail Shadowlands additions |
| `inferred` | from the log | no | built at parse time by counting fields, for a dialect no row matches |

**When Forever's first beta log arrives (Sept 17):**

1. Run `go run ./cmd/forever-logs conformance <dir>` over it. The report names
   the layout that was selected, whether it was inferred, and every unknown
   event and parse error with counts.
2. Add `engine/layout/forever.go` with a `Forever()` function returning the
   row, modelled on `retail.go`. Set `Version` and `ProjectID` from the
   header the log actually writes.
3. Add it to the `rows` slice in `engine/layout/layout.go`.
4. If its `COMBATANT_INFO` differs, set the indexes in the row's `Combatant`
   field; if its shape is not a reordering of retail's, add a decoder branch
   in `engine/event/special.go`. Nothing else changes.
5. Commit a redacted excerpt as `engine/event/testdata/forever.log` with the
   invented cast substituted, and add a width test beside the retail one.

Set `Verified: true` on a row only once its counts have been checked against
a real log of that dialect. An unverified row makes the decoder report a
width mismatch instead of mis-reading a field, which is the behaviour we
want: a `parse_error` in the conformance report is information, a silently
wrong number in a ranking is not.

A fourth row, retail v22 (a nineteen-field advanced block and a forty-two
field `SPELL_DAMAGE` with an `ST`/`AOE` hint, per wowcoach.gg), is documented
but not written: nothing here reads a v22 log yet.

## The fixture policy

**Every line under `engine/*/testdata/` is hand-written.** No line is copied
from a downloaded log. The characters are invented and used consistently:

| Name | GUID | Role |
|---|---|---|
| Baelgrim-Nightslayer | `Player-4184-000000A1` | protection warrior, tanking |
| Sunwick-Nightslayer | `Player-4184-000000A2` | priest, healing |
| Morrowlyn-Nightslayer | `Player-4184-000000A3` | mage |
| Thalgrit-Nightslayer | `Player-4184-000000A4` | hunter |
| Ashfang | `Pet-0-2085-2284-7855-165189-01000000B1` | Thalgrit's pet |
| Hollow Sentinel | `Creature-0-2085-2284-7855-169753-0000AA0001` | trash |
| Warden Kelthas | encounter `9001` | a synthetic encounter id |

Field *counts* and field *order* in the fixture are real — they match the
verified retail v16 table exactly — so the fixture exercises the decoder the
way a real log does. The values are not.

**The real retail sample is never committed.** The benchmark and the
conformance checks fetch it at test time into `logs/testdata/cache/`, which
is git-ignored:

- Source: `https://raw.githubusercontent.com/rp4rk/WoWP/main/WoWCombatLog.txt`
- 77 MB, `COMBAT_LOG_VERSION 16`, build 9.0.2, advanced logging on.
- **Pinned**: exactly 76,979,002 bytes, SHA-256
  `72b3ee25ac51b0e08c2b250e71171ec4c22ab6069df961e946031105c1cfa2bf`. The
  digest is verified after a download and on every cache hit, because the
  source is a third-party repository that can edit or replace the file:
  every measurement in this module — the verified retail row, the 77 MB
  throughput figure, the "0 parse errors over 272,367 lines" conformance
  evidence — describes those exact bytes and nothing else. A mismatch is
  an `ErrUnavailable` naming both digests, not a quiet re-measurement.
- Its repository is AGPL-3.0, which is why it stays out of this one.
- Offline, the tests that need it call `t.Skip` with a message saying so.
- Point `FOREVER_LOGS_SAMPLE` at a local copy to use one instead of
  downloading. A file supplied that way is deliberately **not**
  digest-checked: it is your log, not ours.

## Determinism

`events.parquet` and `summary.json` must be byte-identical across runs and
machines. That means no map iteration in an output path — every slice is
sorted before it is emitted — no wall clock, no randomness, and a pinned
`CreatedBy` in the Parquet metadata. `TestWriteIsByteIdenticalAcrossRuns`
and `TestSnapshotIsDeterministic` are the in-process guards, and
`engine/summary/testdata/v16.summary.json.golden` is the committed one: it
catches the drift two runs in one process cannot see, such as a
dependency bump that changes a rounding, a schema reorder, or a renamed
`Kind`. Regenerate it deliberately, never to make a test pass:

```bash
FOREVER_UPDATE_GOLDEN=1 go test ./engine/summary/ -run TestTheFixtureSummaryMatchesTheCommittedGolden
```

There is no committed Parquet golden. A binary asserting byte-identity
across architectures is a claim this branch cannot verify before CI runs;
it is recorded as a recommendation rather than written on a guess.

Untrusted input is fuzzed rather than only sampled: `FuzzDecode` in
`engine/event` and `FuzzSplitParams` in `engine/lexer` run for twenty
seconds each in CI, seeded from the v16 fixture and from the malformed
shapes that once panicked. A crasher lands in
`engine/*/testdata/fuzz/` and is committed with its fix.

## Unresolved on purpose

These are marked in code rather than guessed, and the tests assert that they
stay marked:

- **The threat model's coefficients.** `summary.BaseThreat` applies one point
  of threat per point of damage and half per point of healing, and carries an
  empty per-spell `Modifiers` table. `Complete()` returns false while it is
  empty, and every `ThreatRow` carries that flag so the report can label the
  table provisional. Fill the table from data when Forever's numbers settle.
- **Classic's `COMBATANT_INFO`.** No source documents it, so the Classic row
  sets `Combatant.Present: false` and such a line is kept raw.
- **Which `COMBAT_LOG_VERSION` Classic clients write.** The Classic row has
  `Version: 0`, so `Lookup` never selects it automatically; pass
  `-layout classic-wiki`, or let the inferred row handle it.
- **`EventNames()` is a cross-product, not a list of real events.**
  `layout.Layout.EventNames()` returns every prefix joined to every suffix,
  so it lists combinations the game never emits — `SWING_HEAL` among them.
  Nothing consumes it today. It must be narrowed to the combinations a
  dialect actually writes before the conformance report adopts it, or that
  report will overstate what the layout can decode.
- **Which patch added the year and zone to the timestamp.** Detected from the
  log by `layout.Infer`, never assumed.
- **`_DRAIN` and `_LEECH` carry the advanced block on the retail row,
  unverified.** The verified list of v16 events carrying the 17-field
  advanced block, established by parsing the real sample, does not include
  `SPELL_DRAIN` or `SPELL_LEECH` — the sample contains none of either. The
  retail row groups them with `_ENERGIZE`, whose advanced block IS verified
  and whose suffix shape they share. If the guess is wrong, such a line
  becomes a width mismatch, which the decoder turns into a `parse_error`
  with the raw text kept and the conformance command reports. Flip one
  boolean in `retail.go` when a log settles it.
- **A fight's summary duration is the span of its events, not the
  segmenter's window.** `summary.Summary.DurationMS` comes from the
  accumulator's own first and last event, while `fight.Fight.Duration()`
  includes the two-second trailing window that keeps damage-over-time ticks
  inside a finished encounter. For an encounter whose trailing window closes
  with nothing landing in it the two disagree by up to that window, so
  per-second figures are computed over the event span rather than over dead
  padding. Trash fights closed by a gap show no discrepancy.
- **The benchmark's memory assertion is a coarse gate.** It asserts
  `runtime.MemStats.Sys` under 1 GiB, and `Sys` is a whole-process,
  run-lifetime high-water mark — so it catches a catastrophic regression but
  would not by itself catch a change that buffered the whole file. The
  measured figure is the real evidence: 37 MB of process footprint for a
  77 MB log is less than half the file, which cannot happen if the file is
  being buffered. Tightening the bound, by asserting the peak stays under
  the file size or by sampling `HeapInuse` during the parse, is an open
  improvement.
