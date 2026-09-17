# Retail v22 Layout Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a second verified combat-log dialect, `retail-v22`, so that a `COMBAT_LOG_VERSION,22` log parses with zero parse errors and no unknown events instead of falling back to the unverified inferred row.

**Architecture:** One new `Layout` row beside `RetailV16()` in `logs/engine/layout/retail.go`, selected by header version exactly as v16 is. The v22 advanced block is 19 fields instead of 17, the damage and miss suffixes carry one trailing field, COMBATANT_INFO's stat indices shift by one, and eleven event types v16 never saw need rows. Three small, shared mechanisms carry that: `Layout.Widths` (the set of field counts a shape may have, replacing the ad-hoc chain in the decoder), a longest-prefix-with-a-known-suffix `Split`, and a layout-driven COMBATANT_INFO stat index map. The v16 row sets none of the new flags, so v16 decoding is byte-identical.

**Tech Stack:** Go 1.x, standard library only. Test framework: `go test`. Golden files regenerated with `FOREVER_UPDATE_GOLDEN=1`.

**Spec:** `docs/superpowers/specs/2026-09-16-retail-v22-layout-design.md`

---

## Global Constraints

Copy these verbatim into your working notes; every task's requirements implicitly include them.

- The engine version in `logs/engine/session/session.go` does not change in this plan; another lane owns it and the web fixture. No file under `web/` is touched.
- The v16 layout, the v16 excerpt, and `v16.summary.json.golden` are byte-identical before and after; a task that regenerates goldens asserts `git diff --exit-code` on the v16 golden.
- `retail-v22` is `Verified: true` and selected by header version 22; the inferred fallback is untouched.
- Acceptance is `forever-logs conformance` over the 89-file corpus with `layout=retail-v22 verified=true parse_errors=0` on every file and no unknown events, or unknowns named with counts in a ledger ruling.
- Commits are made with `git commit -F <file>` in a command of its own (a hook blocks heredocs combined with git commit and the text "no-verify"); the plan's commit steps say so.

---

## Context: what was measured, and how

Everything below was measured before this plan was written. The implementer does not
need to re-derive it, but every number is reproducible from the commands named.

### Corpus

- 89 files, `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/logs/v22-corpus/`, 7.3 GB, **26,030,980 lines**.
- Headers seen, all with `ADVANCED_LOG_ENABLED,1` and `PROJECT_ID,1`:
  `BUILD_VERSION,12.0.5` (264 header lines), `12.0.7` (168), `12.1.0` (1). 433 headers in 89 files — logs are appended across sessions, so a file carries many headers.
- Timestamps are `7/1/2026 09:07:12.137-5`: a year in the date and an hour offset after the
  clock. `ParseStamp` already handles both; the v22 row sets `StampYear: true, StampZone: true`.
- Two probe files used for field-identity work:
  `.../scratchpad/logs/v22/retail-1207-small.txt` (9 MB, build 12.0.7) and
  `.../scratchpad/logs/v22/retail-1210-pvp.txt` (34 MB, build 12.1.0).
  The v16 sample used as the control is `.../scratchpad/logs/wowp.txt`.
- Baseline sweep with the engine at main: 6,725,218 parse errors (26%), 2,958 fights,
  every file `layout=inferred verified=false`, 41 distinct unknown event names.
  Per-file JSON: `.../scratchpad/conformance-v22-before.jsonl`.

Field counts below were produced with a faithful Python port of `lexer.SplitParams`
(commas inside quotes, `[...]` and `(...)` do not split). That matters: a naive CSV
split reports COMBATANT_INFO as 383–455 fields; the real lexer reports **34**, the same
as v16.

### The 19-field advanced block

v16 has 17. The two extra fields sit at block offsets 7 and 8, pushing `absorb` and
everything after it down by two.

| Off | v16 | **v22** | How it was pinned |
|---|---|---|---|
| 0 | infoGUID | infoGUID | unchanged |
| 1 | ownerGUID | ownerGUID | unchanged; `DAMAGE_SPLIT` shows a pet's owner here |
| 2 | currentHP | currentHP | unchanged |
| 3 | maxHP | maxHP | unchanged |
| 4 | attackPower | attackPower | unchanged |
| 5 | spellPower | spellPower | equals COMBATANT_INFO `intellect` for 6/6 players in both probe files |
| 6 | armor | armor | equals COMBATANT_INFO field 24 for 6/6 players in both probe files (v16 control: field 23) |
| 7 | absorb | **versatility** (new) | its only correlated transitions are versatility buffs (Rune of the Versatile Warrior, Mark of the Wild, Lycara's Teachings); across 12 players it is 1.8513 × the COMBATANT_INFO versatility rating, spread 0.18 — the tightest of all 22 COMBATANT_INFO fields tried, the runner-up being 0.79. See the caveat below. |
| 8 | powerType | **unknown8** (new) | 0 for almost every unit; a small per-unit constant (301, 402, 453, 597, 633, 1088) for a few. Matches no COMBATANT_INFO field. **Not pinned** — see the ledger note in Task 1. |
| 9 | currentPower | **absorb** | decisive: walking each unit's advanced blocks in order, the drop in offset 9 equals the `absorbed` field of the damage that hit it in 112 transitions across the two probe files; offset 7 matched once and offset 8 never. |
| 10 | maxPower | **powerType** | equals the `powerType` in the SPELL_ENERGIZE suffix on 5,360 of 6,521 energize lines. v16 control puts the same match at offset 8. |
| 11 | powerCost | **currentPower** | follows from 10 and 12 |
| 12 | positionX | **maxPower** | equals the `maxPower` in the SPELL_ENERGIZE suffix on the same 5,360 lines. v16 control: offset 10. |
| 13 | positionY | **powerCost** | follows |
| 14 | uiMapID | **positionX** | the first of the two decimal-fraction fields |
| 15 | facing | **positionY** | second decimal-fraction field |
| 16 | level | **uiMapID** | integer zone id, 0 in arenas |
| 17 | — | **facing** | radians, 0–2π |
| 18 | — | **itemLevel** | 281–339 on level-80 characters |

Two notes carried into the code as comments:

- **Offset 7 caveat.** Both probe files are arena logs, and arena applies a PvP scaling
  factor to secondary stats. 1/1.8513 = 0.5402, which is consistent with COMBATANT_INFO
  reporting the PvP-adjusted rating and the advanced block reporting the raw one. The
  field is named `Versatility` because every piece of evidence points there, and the
  measured ratio is recorded in the struct comment so a later reader can check it against
  a non-arena log.
- **`Advanced.Level` is item level in v16 too.** The v16 sample shows 189–191 on a
  level-60 character. Renaming the existing field is out of scope for this plan; the v22
  field is documented as item level and a ledger note records the v16 misnomer.

A third-party reference in the scratchpad
(`.../scratchpad/logs/wc_docs_combat-log_advanced-logging`) also describes a 19-field
block, but places `absorb` at offset 7 and calls 8 and 9 "always seen as 0 in modern
logs". The measurements above contradict it on all three, so the measurements win. This
is recorded so nobody "fixes" the layout back to the document.

### The trailing field on damage and miss lines

v22 writes **one extra field at the very end** of some damage and miss lines. It is
either the two-valued tag `ST` / `AOE`, or, on the `_SUPPORT` variants, the supporting
player's GUID.

| Event | trailing field |
|---|---|
| SPELL_DAMAGE, SPELL_PERIODIC_DAMAGE, RANGE_DAMAGE, DAMAGE_SPLIT | `ST` / `AOE` |
| SPELL_MISSED, SPELL_PERIODIC_MISSED, DAMAGE_SHIELD_MISSED | `ST` / `AOE` |
| SPELL_DAMAGE_SUPPORT, SPELL_PERIODIC_DAMAGE_SUPPORT, RANGE_DAMAGE_SUPPORT, SWING_DAMAGE_LANDED_SUPPORT | supporter GUID |
| SPELL_HEAL_SUPPORT, SPELL_PERIODIC_HEAL_SUPPORT, SPELL_ABSORBED_SUPPORT | supporter GUID |
| **SWING_DAMAGE, SWING_DAMAGE_LANDED, SWING_MISSED, RANGE_MISSED, ENVIRONMENTAL_DAMAGE** | **none** |

`RANGE_MISSED` is the awkward one: `RANGE_DAMAGE` carries the tag and `RANGE_MISSED`
does not, at 4,103 lines, so it is not a sampling accident. No prefix rule fits. The
design therefore reads the tag **by value** — `ST` and `AOE` are the only two strings it
ever holds, and no neighbouring field (`critical`, `isOffHand`, `crushing`, all `nil` /
`0` / `1`) can collide with them. The reasoning is written into the code.

`_MISSED` also gained a conditional field independent of the tag: when `missType` is
`BLOCK` or `RESIST` the line carries one extra `amountMissed` (21 lines corpus-wide,
e.g. `...,BLOCK,nil,0,ST`). `ABSORB` still carries three extras as in v16.

### COMBATANT_INFO

Still **34 fields**, but reshuffled. One extra defensive rating is inserted among
dodge / parry / block (v16 fields 7–9 become v22 fields 7–10; all four are 0 in every
sample, so which one is new cannot be told apart and does not matter), shifting
everything up to and including armor by +1. v16's borrowed-power field is gone, which is
why the total is unchanged.

| | v16 | v22 |
|---|---|---|
| strength / agility / stamina / intellect | 3 / 4 / 5 / 6 | 3 / 4 / 5 / 6 |
| dodge, parry, block (+1 unnamed in v22) | 7, 8, 9 | 7, 8, 9, 10 |
| crit (melee/ranged/spell) | 10, 11, 12 | 11, 12, 13 |
| speed | 13 | 14 |
| lifesteal | 14 | 15 |
| haste (melee/ranged/spell) | 15, 16, 17 | 16, 17, 18 |
| avoidance | 18 | 19 |
| mastery | 19 | 20 |
| versatility (done/heal/taken) | 20, 21, 22 | 21, 22, 23 |
| **armor** | **23** | **24** |
| **specID** | **24** | **25** |
| **talents** | **25** | **26** |
| **pvpTalents** | **26** | **27** |
| **borrowed power** | **27** | **absent** |
| **gear** | **28** | **28** |
| **auras** | **29** | **29** |

The decoder currently hard-codes the stat indices 3–23 inside `readCombatant`. This plan
moves them into the layout row as a `StatIndex` map so each dialect declares its own.

### Full width table, 89 files, 26,030,980 lines

Every event name the corpus contains, with total field counts under the real lexer. The
"v16" column is the v16 row's declared width where it has one. This is the table the
width tests assert against.

| Event | count | v22 widths | v16 | reading |
|---|---|---|---|---|
| ARENA_MATCH_END | 587 | 5 | — | new |
| ARENA_MATCH_START | 1,262 | 5 | — | new |
| COMBATANT_INFO | 7,216 | 34 | 34 | same width, shifted indices |
| COMBAT_LOG_VERSION | 433 | 8 | 8 | unchanged |
| DAMAGE_SHIELD_MISSED | 4 | 15 | — | new; 9+3+2+tag |
| DAMAGE_SPLIT | 138,590 | 42 | — | new; SPELL_DAMAGE's shape |
| EMOTE | 431 | 6 | 6 | unchanged |
| ENCHANT_APPLIED | 132 | 12 | 12 | unchanged |
| ENCHANT_REMOVED | 83 | 12 | 12 | unchanged |
| ENVIRONMENTAL_DAMAGE | 881 | 39 | 37 | advanced +2, no tag |
| MAP_CHANGE | 1,141 | 7 | 7 | unchanged |
| PARTY_KILL | 1,601 | 10 | 10 | unchanged |
| RANGE_DAMAGE | 20,853 | 42 | 39 | advanced +2, tag +1 |
| RANGE_DAMAGE_SUPPORT | 14 | 42 | — | new |
| RANGE_MISSED | 4,103 | 14, 17 | 14, 17 | unchanged, no tag |
| SPELL_ABSORBED | 1,845,765 | 19, 22 | 19, 22 | unchanged |
| SPELL_ABSORBED_SUPPORT | 213 | 20, 23 | — | new; +supporter GUID |
| SPELL_AURA_APPLIED | 2,853,535 | 13, 14, 15 | 13, 14 | new 15-wide form |
| SPELL_AURA_APPLIED_DOSE | 1,165,212 | 14 | 14 | unchanged |
| SPELL_AURA_BROKEN | 34 | 13 | 13 | unchanged |
| SPELL_AURA_BROKEN_SPELL | 15,047 | 16 | 16 | unchanged |
| SPELL_AURA_REFRESH | 2,125,275 | 13, 14 | 13, 14 | unchanged |
| SPELL_AURA_REMOVED | 2,081,340 | 13, 14, 15 | 13, 14 | new 15-wide form |
| SPELL_AURA_REMOVED_DOSE | 458,446 | 14 | 14 | unchanged |
| SPELL_CAST_FAILED | 53,656 | 13 | 13 | unchanged |
| SPELL_CAST_START | 319,509 | 12 | 12 | unchanged |
| SPELL_CAST_SUCCESS | 1,393,121 | 31 | 29 | advanced +2 |
| SPELL_CREATE | 4,460 | 12 | 12 | unchanged |
| SPELL_DAMAGE | 4,019,238 | 42 | 39 | advanced +2, tag +1 |
| SPELL_DAMAGE_SUPPORT | 67,386 | 42 | — | new |
| SPELL_DISPEL | 27,262 | 16 | 16 | unchanged |
| SPELL_DRAIN | 1,466 | 35 | 33 | advanced +2 |
| SPELL_EMPOWER_END | 4,439 | 13 | — | new |
| SPELL_EMPOWER_INTERRUPT | 294 | 13 | — | new |
| SPELL_EMPOWER_START | 4,617 | 12 | — | new |
| SPELL_ENERGIZE | 960,018 | 35 | 33 | advanced +2 |
| SPELL_EXTRA_ATTACKS | 5,156 | 13 | 13 | unchanged |
| SPELL_HEAL | 3,360,710 | 36 | 34 | advanced +2 |
| SPELL_HEAL_ABSORBED | 210,597 | 21 | 21 | unchanged |
| SPELL_HEAL_SUPPORT | 1,036 | 37 | — | new |
| SPELL_INSTAKILL | 114 | 13 | 13 | unchanged |
| SPELL_INTERRUPT | 4,523 | 15 | 15 | unchanged |
| SPELL_MISSED | 513,406 | 15, 16, 18 | 14, 17 | tag +1; 16 = BLOCK/RESIST amount |
| SPELL_PERIODIC_DAMAGE | 1,700,800 | 42 | 39 | advanced +2, tag +1 |
| SPELL_PERIODIC_DAMAGE_SUPPORT | 14,900 | 42 | — | new |
| SPELL_PERIODIC_ENERGIZE | 99,719 | 35 | 33 | advanced +2 |
| SPELL_PERIODIC_HEAL | 932,694 | 36 | 34 | advanced +2 |
| SPELL_PERIODIC_HEAL_SUPPORT | 320 | 37 | — | new |
| SPELL_PERIODIC_MISSED | 283,411 | 15, 18 | 14, 17 | tag +1 |
| SPELL_RESURRECT | 2 | 12 | 12 | unchanged |
| SPELL_STOLEN | 768 | 16 | 16 | unchanged |
| SPELL_SUMMON | 118,117 | 12 | 12 | unchanged |
| STAGGER_CLEAR | 172 | 3 | — | new; `guid, amount` |
| STAGGER_PREVENTED | 5 | 4 | — | new; `guid, spellID, amount` |
| SWING_DAMAGE | 509,671 | 38 | 36 | advanced +2, no tag |
| SWING_DAMAGE_LANDED | 550,396 | 38 | 36 | advanced +2, no tag |
| SWING_DAMAGE_LANDED_SUPPORT | 4,206 | 42 | — | new; carries a spell triple |
| SWING_MISSED | 112,089 | 11, 14 | 11, 14 | unchanged |
| UNIT_DIED | 28,556 | 10 | 10 | unchanged |
| WORLD_MARKER_PLACED | 6 | 5 | — | new; `mapID, index, x, y` |
| WORLD_MARKER_REMOVED | 6 | 2 | — | new; `index` |
| ZONE_CHANGE | 1,936 | 4 | 4 | unchanged |

Events in the v16 row that the corpus never shows: `UNIT_DESTROYED`, `UNIT_DISSIPATES`,
`ENCOUNTER_START`, `ENCOUNTER_END`, `CHALLENGE_MODE_START`, `CHALLENGE_MODE_END`,
`SPELL_BUILDING_*`, `SPELL_LEECH`, `SPELL_DISPEL_FAILED`, `SPELL_DURABILITY_DAMAGE`.
The v22 row keeps them at their v16 widths (the corpus is arena- and dummy-heavy, not
proof of absence) except where a width provably changed.

`SWING_DAMAGE_LANDED_SUPPORT` is the one event whose name lies about its shape: despite
the `SWING` name it carries a spell triple naming the *supporting* spell
(`395152,"Ebon Might",0xc`), so it is 9 + 3 + 19 + 10 + 1 = 42, not 38 + 1.

---

## File Structure

| File | Change | Responsibility |
|---|---|---|
| `logs/engine/event/testdata/v22.log` | create (Task 1) | contiguous 4,000-line excerpt of a real v22 arena log; drives the v22 summary golden |
| `logs/engine/event/testdata/v22-shapes.log` | create (Task 1) | one line per event shape the excerpt does not contain; drives shape coverage only |
| `logs/engine/layout/measured_v22_test.go` | create (Task 1) | asserts the committed testdata really contains the measured widths, independent of any layout |
| `logs/engine/layout/layout.go` | modify (Task 2) | `Suffix` flags, `Combatant.StatIndex`, `Layout.Widths`, `Split` |
| `logs/engine/layout/retail.go` | modify (Task 2, 3) | the `RetailV22()` row |
| `logs/engine/layout/layout_test.go` | modify (Task 2, 3) | v22 twins of the v16 width and selection tests |
| `logs/engine/event/event.go` | modify (Task 3) | `Advanced.Versatility`, `Advanced.Unknown8`, `Event.Scope`, `Event.Supporter` |
| `logs/engine/event/decode.go` | modify (Task 3) | 19-field advanced block, trailing tag, `_SUPPORT`, `_SPLIT`, `_EMPOWER_*`, width set check |
| `logs/engine/event/special.go` | modify (Task 3) | layout-driven COMBATANT_INFO stats, the new specials |
| `logs/engine/event/decode_v22_test.go` | create (Task 3) | every v22 shape decodes without error |
| `logs/engine/summary/testdata/v22.summary.json.golden` | create (Task 4) | the v22 summary golden |
| `logs/engine/summary/golden_test.go` | modify (Task 4) | v22 golden twin |
| `docs/ledger/2026-09-16-retail-v22.md` | create (Task 1), append (Task 3, 5) | unpinned fields and documented exceptions |

---

## Task 1: Commit the measured v22 testdata and lock the measurement

The layout does not exist yet, so this task's test asserts facts about the **log files**,
not about any `Layout`. That is what makes the measurement a regression test rather than
a comment: if someone later regenerates the excerpt and drops a shape, this fails.

**Files:**
- Create: `logs/engine/event/testdata/v22.log`
- Create: `logs/engine/event/testdata/v22-shapes.log`
- Create: `logs/engine/layout/measured_v22_test.go`
- Create: `docs/ledger/2026-09-16-retail-v22.md`

**Interfaces:**
- Consumes: nothing.
- Produces: the two testdata files, at the paths above, read by Tasks 2, 3 and 4. `v22.log`
  is a contiguous excerpt (a real arena match, usable for fight segmentation and the
  summary golden). `v22-shapes.log` is a synthetic concatenation of real lines, in
  timestamp order, with its own `COMBAT_LOG_VERSION` line first; it is **not** used for
  the summary golden.

- [ ] **Step 1: Build the two excerpts**

`SRC` and `CORPUS` are session-local scratchpad paths; after this step the testdata is
committed and self-contained.

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
SCRATCH=/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad
SRC="$SCRATCH/logs/v22/retail-1207-small.txt"
CORPUS="$SCRATCH/logs/v22-corpus"

# The contiguous excerpt: header, two ZONE_CHANGEs, ARENA_MATCH_START, six
# COMBATANT_INFO lines and the opening of a Rated Solo Shuffle round.
head -4000 "$SRC" > logs/engine/event/testdata/v22.log
wc -l logs/engine/event/testdata/v22.log
```

Expected: `4000 logs/engine/event/testdata/v22.log`

- [ ] **Step 2: Build the shapes file**

The excerpt covers 32 of the corpus's 65 event names. This collects one real line for
every remaining `(event, width)` pair. The widths are the ones in the context table.

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
SCRATCH=/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad
CORPUS="$SCRATCH/logs/v22-corpus"
OUT=logs/engine/event/testdata/v22-shapes.log

# grep -m1 over the corpus for each shape; awk -F, filters on the naive comma
# count, which equals the lexer's field count for every event below except
# COMBATANT_INFO, which the contiguous excerpt already covers.
{
  printf '7/1/2026 09:00:00.000-5  COMBAT_LOG_VERSION,22,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,12.0.7,PROJECT_ID,1\n'
  for spec in \
    "ARENA_MATCH_END 5" "DAMAGE_SHIELD_MISSED 15" "DAMAGE_SPLIT 42" "EMOTE 6" \
    "ENCHANT_APPLIED 12" "ENCHANT_REMOVED 12" "ENVIRONMENTAL_DAMAGE 39" "MAP_CHANGE 7" \
    "PARTY_KILL 10" "RANGE_DAMAGE_SUPPORT 42" "RANGE_MISSED 14" \
    "SPELL_ABSORBED_SUPPORT 20" "SPELL_ABSORBED_SUPPORT 23" "SPELL_AURA_BROKEN 13" \
    "SPELL_AURA_APPLIED 15" "SPELL_AURA_REMOVED 15" "SPELL_DAMAGE_SUPPORT 42" \
    "SPELL_DRAIN 35" "SPELL_EMPOWER_START 12" "SPELL_EMPOWER_END 13" \
    "SPELL_EMPOWER_INTERRUPT 13" "SPELL_HEAL_ABSORBED 21" "SPELL_HEAL_SUPPORT 37" \
    "SPELL_INSTAKILL 13" "SPELL_INTERRUPT 15" "SPELL_MISSED 16" \
    "SPELL_PERIODIC_DAMAGE_SUPPORT 42" "SPELL_PERIODIC_HEAL_SUPPORT 37" \
    "SPELL_PERIODIC_MISSED 15" "SPELL_RESURRECT 12" "SPELL_STOLEN 16" \
    "STAGGER_CLEAR 3" "STAGGER_PREVENTED 4" "SWING_DAMAGE_LANDED_SUPPORT 42" \
    "WORLD_MARKER_PLACED 5" "WORLD_MARKER_REMOVED 2"
  do
    set -- $spec
    grep -h "  $1," "$CORPUS"/*.txt | awk -F, -v n="$2" 'NF==n {print; exit}'
  done
} > "$OUT"
wc -l "$OUT"
```

Expected: `37 logs/engine/event/testdata/v22-shapes.log` (one header plus 36 shape lines).
If a line is missing the count is lower; re-run the single `grep` for that shape and add
it by hand rather than accepting a short file.

- [ ] **Step 3: Write the failing measurement test**

Create `logs/engine/layout/measured_v22_test.go`:

```go
// logs/engine/layout/measured_v22_test.go
package layout

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// v22Testdata is the committed v22 excerpt and the shapes file beside it.
// They live in the event package's testdata because the decoder tests read
// them too; this package reads them to pin the measurement the retail-v22
// row is built from.
var v22Testdata = []string{
	filepath.Join("..", "event", "testdata", "v22.log"),
	filepath.Join("..", "event", "testdata", "v22-shapes.log"),
}

// v22MeasuredWidths is the field count of every event shape the committed
// testdata must contain. It is the subset of the 89-file corpus measurement
// (26,030,980 lines) that the excerpts were cut to cover, and it is what the
// retail-v22 row's widths are checked against in TestRetailV22Widths.
//
// A shape dropped from the testdata is a shape nothing tests, so this list
// failing is a real failure, not a chore.
var v22MeasuredWidths = map[string][]int{
	"ARENA_MATCH_END":               {5},
	"ARENA_MATCH_START":             {5},
	"COMBATANT_INFO":                {34},
	"COMBAT_LOG_VERSION":            {8},
	"DAMAGE_SHIELD_MISSED":          {15},
	"DAMAGE_SPLIT":                  {42},
	"EMOTE":                         {6},
	"ENCHANT_APPLIED":               {12},
	"ENCHANT_REMOVED":               {12},
	"ENVIRONMENTAL_DAMAGE":          {39},
	"MAP_CHANGE":                    {7},
	"PARTY_KILL":                    {10},
	"RANGE_DAMAGE":                  {42},
	"RANGE_DAMAGE_SUPPORT":          {42},
	"RANGE_MISSED":                  {14, 17},
	"SPELL_ABSORBED":                {19, 22},
	"SPELL_ABSORBED_SUPPORT":        {20, 23},
	"SPELL_AURA_APPLIED":            {13, 14, 15},
	"SPELL_AURA_APPLIED_DOSE":       {14},
	"SPELL_AURA_BROKEN":             {13},
	"SPELL_AURA_BROKEN_SPELL":       {16},
	"SPELL_AURA_REFRESH":            {13, 14},
	"SPELL_AURA_REMOVED":            {13, 14, 15},
	"SPELL_AURA_REMOVED_DOSE":       {14},
	"SPELL_CAST_FAILED":             {13},
	"SPELL_CAST_START":              {12},
	"SPELL_CAST_SUCCESS":            {31},
	"SPELL_CREATE":                  {12},
	"SPELL_DAMAGE":                  {42},
	"SPELL_DAMAGE_SUPPORT":          {42},
	"SPELL_DISPEL":                  {16},
	"SPELL_DRAIN":                   {35},
	"SPELL_EMPOWER_END":             {13},
	"SPELL_EMPOWER_INTERRUPT":       {13},
	"SPELL_EMPOWER_START":           {12},
	"SPELL_ENERGIZE":                {35},
	"SPELL_EXTRA_ATTACKS":           {13},
	"SPELL_HEAL":                    {36},
	"SPELL_HEAL_ABSORBED":           {21},
	"SPELL_HEAL_SUPPORT":            {37},
	"SPELL_INSTAKILL":               {13},
	"SPELL_INTERRUPT":               {15},
	"SPELL_MISSED":                  {15, 16, 18},
	"SPELL_PERIODIC_DAMAGE":         {42},
	"SPELL_PERIODIC_DAMAGE_SUPPORT": {42},
	"SPELL_PERIODIC_ENERGIZE":       {35},
	"SPELL_PERIODIC_HEAL":           {36},
	"SPELL_PERIODIC_HEAL_SUPPORT":   {37},
	"SPELL_PERIODIC_MISSED":         {15, 18},
	"SPELL_RESURRECT":               {12},
	"SPELL_STOLEN":                  {16},
	"SPELL_SUMMON":                  {12},
	"STAGGER_CLEAR":                 {3},
	"STAGGER_PREVENTED":             {4},
	"SWING_DAMAGE":                  {38},
	"SWING_DAMAGE_LANDED":           {38},
	"SWING_DAMAGE_LANDED_SUPPORT":   {42},
	"SWING_MISSED":                  {11, 14},
	"UNIT_DIED":                     {10},
	"WORLD_MARKER_PLACED":           {5},
	"WORLD_MARKER_REMOVED":          {2},
	"ZONE_CHANGE":                   {4},
}

// v22TestdataWidths lexes the committed excerpts and returns every
// (event, width) pair they contain.
func v22TestdataWidths(t *testing.T) map[string][]int {
	t.Helper()
	seen := map[string]map[int]bool{}
	for _, path := range v22Testdata {
		text, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		for _, ln := range lines(t, string(text)) {
			if len(ln.Params) == 0 || ln.Params[0] == "" {
				continue
			}
			ev := ln.Params[0]
			if seen[ev] == nil {
				seen[ev] = map[int]bool{}
			}
			seen[ev][len(ln.Params)] = true
		}
	}
	out := map[string][]int{}
	for ev, ws := range seen {
		for w := range ws {
			out[ev] = append(out[ev], w)
		}
		sort.Ints(out[ev])
	}
	return out
}

func TestTheV22TestdataCoversEveryMeasuredShape(t *testing.T) {
	got := v22TestdataWidths(t)
	for ev, want := range v22MeasuredWidths {
		have := got[ev]
		for _, w := range want {
			found := false
			for _, g := range have {
				if g == w {
					found = true
				}
			}
			if !found {
				t.Errorf("the v22 testdata has no %s line of width %d (it has %v)", ev, w, have)
			}
		}
	}
}

func TestTheV22TestdataHasNoShapeTheMeasurementDoesNotName(t *testing.T) {
	for ev, have := range v22TestdataWidths(t) {
		want, known := v22MeasuredWidths[ev]
		if !known {
			t.Errorf("the v22 testdata contains %s, which the measurement does not name", ev)
			continue
		}
		for _, g := range have {
			found := false
			for _, w := range want {
				if g == w {
					found = true
				}
			}
			if !found {
				t.Errorf("the v22 testdata has a %s line of width %d; measured widths are %v", ev, g, want)
			}
		}
	}
}

func TestTheV22ExcerptHeaderIsVersion22(t *testing.T) {
	text, err := os.ReadFile(v22Testdata[0])
	if err != nil {
		t.Fatal(err)
	}
	ls := lines(t, string(text))
	h, ok := ParseHeader(ls[0])
	if !ok {
		t.Fatal("the first line of the v22 excerpt is not a COMBAT_LOG_VERSION line")
	}
	if h.Version != 22 || !h.Advanced || h.ProjectID != 1 || h.Fields != 8 {
		t.Fatalf("header = %+v, want version 22, advanced, project 1, 8 fields", h)
	}
	if h.Build != "12.0.7" {
		t.Errorf("build = %q, want 12.0.7", h.Build)
	}
}

func TestTheV22ExcerptTimestampsCarryAYearAndAZone(t *testing.T) {
	text, err := os.ReadFile(v22Testdata[0])
	if err != nil {
		t.Fatal(err)
	}
	year, zone := inferStamp(lines(t, string(text))[0].Stamp)
	if !year || !zone {
		t.Fatalf("year=%v zone=%v, want both true", year, zone)
	}
}
```

- [ ] **Step 4: Run the test to verify it passes on the committed data**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go test ./engine/layout/ -run 'TestTheV22' -v 2>&1 | tail -20
```

Expected: `PASS`, with `TestTheV22TestdataCoversEveryMeasuredShape`,
`TestTheV22TestdataHasNoShapeTheMeasurementDoesNotName`,
`TestTheV22ExcerptHeaderIsVersion22` and
`TestTheV22ExcerptTimestampsCarryAYearAndAZone` all `--- PASS`.

If `TestTheV22TestdataCoversEveryMeasuredShape` fails, Step 2's grep missed a shape: go
back and add the missing line. If `...HasNoShapeTheMeasurementDoesNotName` fails, the
excerpt contains an event the context table does not list — stop and add it to the table
and to `v22MeasuredWidths`, because the corpus measurement was then incomplete.

- [ ] **Step 5: Write the ledger note for the fields that are not pinned**

Create `docs/ledger/2026-09-16-retail-v22.md`:

```markdown
# Retail v22 layout: what is measured and what is not

Measured over 89 real v22 logs (26,030,980 lines, builds 12.0.5, 12.0.7 and 12.1.0).
The plan at `docs/superpowers/plans/2026-09-16-retail-v22-layout.md` carries the full
field table; this file records only what could not be pinned, so that a later reader
knows which names in the code are evidence and which are placeholders.

## Advanced block offset 8: `Unknown8`

Zero for almost every unit in every log. Non-zero as a small per-unit constant for a
handful of players: 301, 402, 453, 597, 633, 1088. It matches no COMBATANT_INFO stat
field, does not move when the unit takes or absorbs damage, and does not track any
aura. It is counted so that the 19-field block's widths verify, and it is read into
`Advanced.Unknown8` so nothing silently shifts if it is later identified.

To identify it: find a character whose value is non-zero, and diff their
COMBATANT_INFO and equipped items against a character whose value is zero.

## Advanced block offset 7: `Versatility`, with a scale caveat

Offset 7 tracks versatility buffs and nothing else, and sits at a constant
1.8513x the COMBATANT_INFO versatility rating across 12 players in two independent
logs (spread 0.18; the runner-up field's spread is 0.79). Both logs are arena logs and
arena scales secondary stats, so the most likely reading is that COMBATANT_INFO reports
the PvP-adjusted rating and the advanced block reports the raw one: 1/1.8513 = 0.5402.

To confirm: run the same ratio over a raid or dungeon log, where the factor should be
1.0. Until then the name stands on the correlation, not on the scale.

## COMBATANT_INFO's fourth defensive rating

v22 writes four fields where v16 writes dodge, parry and block. All four are 0 in every
sample in the corpus, so which of the four is new cannot be told apart. The decoder
reads dodge, parry and block from the first three and ignores the fourth. Harmless
while they are all zero; revisit if a tank log ever shows non-zero values.

## `Advanced.Level` is item level, in v16 as well as v22

The v16 sample shows 189-191 in this field on a level-60 character, and v22 shows
281-339 on level-80 characters: both are item level, not character level. The v22
documentation names it correctly. Renaming the Go field would touch v16 decoding and is
out of scope for the v22 lane.
```

- [ ] **Step 6: Commit**

The commit message goes in a file first; `git commit -F` is used in a command of its
own, because a hook blocks a heredoc combined with `git commit` and the text
`no-verify`.

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
cat > /tmp/v22-msg-1.txt <<'MSG'
test(logs): commit the v22 excerpts and pin the measured shapes

Two committed excerpts of real COMBAT_LOG_VERSION 22 logs: a contiguous
4,000-line arena excerpt, and one line per event shape the excerpt does
not contain. The new test asserts the excerpts really hold every shape
the 89-file corpus measurement names, so a shape cannot quietly leave
the testdata.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
git add logs/engine/event/testdata/v22.log logs/engine/event/testdata/v22-shapes.log logs/engine/layout/measured_v22_test.go docs/ledger/2026-09-16-retail-v22.md
```

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22 && git commit -F /tmp/v22-msg-1.txt
```

---

## Task 2: The `retail-v22` row and its selection

**Files:**
- Modify: `logs/engine/layout/layout.go`
- Modify: `logs/engine/layout/retail.go`
- Modify: `logs/engine/layout/layout_test.go`

**Interfaces:**
- Consumes: `v22MeasuredWidths` from Task 1.
- Produces:
  - `layout.RetailV22() Layout` — the v22 row, `Name: "retail-v22"`, `Version: 22`, `Verified: true`.
  - `Suffix.Tag bool`, `Suffix.MissAmount bool`, `Suffix.AuraExtra bool` — new optional-field flags, all false on v16 and classic.
  - `Combatant.StatIndex map[string]int` — COMBATANT_INFO stat name to field index.
  - `(Layout).Widths(prefix, suffix string) []int` — every total field count the row accepts for that shape, ascending, deduplicated.
  - `(Layout).Split` now returns the longest prefix **whose remainder is a known suffix**.

- [ ] **Step 1: Write the failing selection and width tests**

Append to `logs/engine/layout/layout_test.go`:

```go
func TestLookupPicksRetailV22(t *testing.T) {
	on, ok := Lookup(Header{Version: 22, ProjectID: 1, Advanced: true})
	if !ok || on.Name != "retail-v22" || on.Advanced != 19 {
		t.Fatalf("advanced on: name=%q advanced=%d ok=%v", on.Name, on.Advanced, ok)
	}
	if !on.Verified {
		t.Error("retail-v22 must be a verified row")
	}
	if !on.StampYear || !on.StampZone {
		t.Errorf("stampYear=%v stampZone=%v, want both true", on.StampYear, on.StampZone)
	}
	off, ok := Lookup(Header{Version: 22, ProjectID: 1, Advanced: false})
	if !ok || off.Advanced != 0 {
		t.Fatalf("advanced off: advanced=%d ok=%v", off.Advanced, ok)
	}
	// v16 must still win for version 16.
	v16, ok := Lookup(Header{Version: 16, ProjectID: 1, Advanced: true})
	if !ok || v16.Name != "retail-v16" || v16.Advanced != 17 {
		t.Fatalf("v16 selection regressed: name=%q advanced=%d", v16.Name, v16.Advanced)
	}
}

// TestRetailV22WidthsMatchTheVerifiedCounts is the v22 twin of
// TestRetailV16WidthsMatchTheVerifiedCounts. wantWidths is the complete set
// the row accepts; the measured widths from the corpus must all be in it.
func TestRetailV22WidthsMatchTheVerifiedCounts(t *testing.T) {
	l := RetailV22()
	for _, tc := range []struct {
		event      string
		wantWidths []int
		wantAdvAt  int
	}{
		{"SPELL_DAMAGE", []int{41, 42}, 12},
		{"SPELL_PERIODIC_DAMAGE", []int{41, 42}, 12},
		{"RANGE_DAMAGE", []int{41, 42}, 12},
		{"DAMAGE_SPLIT", []int{41, 42}, 12},
		{"SWING_DAMAGE", []int{38, 39}, 9},
		{"SWING_DAMAGE_LANDED", []int{38, 39}, 9},
		{"SPELL_DAMAGE_SUPPORT", []int{42}, 12},
		{"SPELL_PERIODIC_DAMAGE_SUPPORT", []int{42}, 12},
		{"RANGE_DAMAGE_SUPPORT", []int{42}, 12},
		{"SWING_DAMAGE_LANDED_SUPPORT", []int{42}, 12},
		{"SPELL_HEAL", []int{36}, 12},
		{"SPELL_PERIODIC_HEAL", []int{36}, 12},
		{"SPELL_HEAL_SUPPORT", []int{37}, 12},
		{"SPELL_PERIODIC_HEAL_SUPPORT", []int{37}, 12},
		{"SPELL_ENERGIZE", []int{35}, 12},
		{"SPELL_PERIODIC_ENERGIZE", []int{35}, 12},
		{"SPELL_DRAIN", []int{35}, 12},
		{"SPELL_CAST_SUCCESS", []int{31}, 12},
		{"SPELL_CAST_START", []int{12}, -1},
		{"SPELL_CAST_FAILED", []int{13}, -1},
		{"SPELL_EMPOWER_START", []int{12}, -1},
		{"SPELL_EMPOWER_END", []int{13}, -1},
		{"SPELL_EMPOWER_INTERRUPT", []int{13}, -1},
		{"SPELL_AURA_APPLIED", []int{13, 14, 15}, -1},
		{"SPELL_AURA_REMOVED", []int{13, 14, 15}, -1},
		{"SPELL_AURA_REFRESH", []int{13, 14, 15}, -1},
		{"SPELL_AURA_APPLIED_DOSE", []int{14}, -1},
		{"SPELL_AURA_BROKEN_SPELL", []int{16}, -1},
		{"SPELL_INTERRUPT", []int{15}, -1},
		{"SPELL_DISPEL", []int{16}, -1},
		{"SPELL_STOLEN", []int{16}, -1},
		{"SPELL_SUMMON", []int{12}, -1},
		{"SPELL_CREATE", []int{12}, -1},
		{"SPELL_RESURRECT", []int{12}, -1},
		{"SPELL_INSTAKILL", []int{13}, -1},
		{"SPELL_EXTRA_ATTACKS", []int{13}, -1},
		{"SWING_MISSED", []int{11, 12, 13, 14, 15}, -1},
		{"RANGE_MISSED", []int{14, 15, 16, 17, 18}, -1},
		{"SPELL_MISSED", []int{14, 15, 16, 17, 18}, -1},
		{"SPELL_PERIODIC_MISSED", []int{14, 15, 16, 17, 18}, -1},
		{"DAMAGE_SHIELD_MISSED", []int{14, 15, 16, 17, 18}, -1},
	} {
		t.Run(tc.event, func(t *testing.T) {
			prefix, suffix, ok := l.Split(tc.event)
			if !ok {
				t.Fatalf("Split(%q) = %q %q, not known", tc.event, prefix, suffix)
			}
			_, advAt := l.Width(prefix, suffix)
			if advAt != tc.wantAdvAt {
				t.Errorf("advAt = %d, want %d", advAt, tc.wantAdvAt)
			}
			got := l.Widths(prefix, suffix)
			if len(got) != len(tc.wantWidths) {
				t.Fatalf("widths = %v, want %v", got, tc.wantWidths)
			}
			for i := range got {
				if got[i] != tc.wantWidths[i] {
					t.Fatalf("widths = %v, want %v", got, tc.wantWidths)
				}
			}
		})
	}
}

// TestRetailV22AcceptsEveryMeasuredWidth ties the row back to the corpus
// measurement: every (event, width) pair seen in 26,030,980 real lines must
// be a width the row accepts, whether the event is a prefix/suffix shape or
// a special.
func TestRetailV22AcceptsEveryMeasuredWidth(t *testing.T) {
	l := RetailV22()
	for event, widths := range v22MeasuredWidths {
		if event == "COMBAT_LOG_VERSION" {
			continue // the header is parsed before the row is consulted
		}
		for _, w := range widths {
			if s, ok := l.Specials[event]; ok {
				if !s.Accepts(w) {
					t.Errorf("special %s does not accept measured width %d (accepts %v)", event, w, s.Widths)
				}
				continue
			}
			prefix, suffix, known := l.Split(event)
			if !known {
				t.Errorf("%s is neither a special nor a known prefix/suffix on retail-v22", event)
				continue
			}
			found := false
			for _, got := range l.Widths(prefix, suffix) {
				if got == w {
					found = true
				}
			}
			if !found {
				t.Errorf("%s width %d is not accepted (row accepts %v)", event, w, l.Widths(prefix, suffix))
			}
		}
	}
}

// TestRetailV16WidthsAreUnchangedByTheWidthsHelper guards the refactor: the
// set of widths the v16 row accepts must be exactly what the decoder
// accepted before Widths existed.
func TestRetailV16WidthsAreUnchangedByTheWidthsHelper(t *testing.T) {
	l := RetailV16()
	for _, tc := range []struct {
		event string
		want  []int
	}{
		{"SPELL_DAMAGE", []int{39}},
		{"SWING_DAMAGE", []int{36}},
		{"SPELL_HEAL", []int{34}},
		{"SPELL_ENERGIZE", []int{33}},
		{"SPELL_CAST_SUCCESS", []int{29}},
		{"SPELL_MISSED", []int{14, 17}},
		{"SWING_MISSED", []int{11, 14}},
		{"SPELL_AURA_APPLIED", []int{13, 14}},
		{"SPELL_AURA_APPLIED_DOSE", []int{14}},
		{"SPELL_AURA_BROKEN_SPELL", []int{16}},
	} {
		t.Run(tc.event, func(t *testing.T) {
			prefix, suffix, ok := l.Split(tc.event)
			if !ok {
				t.Fatalf("Split(%q) failed", tc.event)
			}
			got := l.Widths(prefix, suffix)
			if len(got) != len(tc.want) {
				t.Fatalf("widths = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("widths = %v, want %v", got, tc.want)
				}
			}
		})
	}
}

// TestSplitPrefersThePrefixThatLeavesAKnownSuffix is the rule that lets the
// v22 row register both SWING and SWING_DAMAGE_LANDED without the longer
// one swallowing the shorter one's events.
func TestSplitPrefersThePrefixThatLeavesAKnownSuffix(t *testing.T) {
	l := RetailV22()
	for _, tc := range []struct{ event, prefix, suffix string }{
		{"SWING_DAMAGE_LANDED", "SWING", "_DAMAGE_LANDED"},
		{"SWING_DAMAGE_LANDED_SUPPORT", "SWING_DAMAGE_LANDED", "_SUPPORT"},
		{"SPELL_DAMAGE", "SPELL", "_DAMAGE"},
		{"SPELL_DAMAGE_SUPPORT", "SPELL_DAMAGE", "_SUPPORT"},
		{"SPELL_PERIODIC_DAMAGE", "SPELL_PERIODIC", "_DAMAGE"},
		{"SPELL_PERIODIC_DAMAGE_SUPPORT", "SPELL_PERIODIC_DAMAGE", "_SUPPORT"},
		{"RANGE_DAMAGE", "RANGE", "_DAMAGE"},
		{"RANGE_DAMAGE_SUPPORT", "RANGE_DAMAGE", "_SUPPORT"},
		{"SPELL_HEAL", "SPELL", "_HEAL"},
		{"SPELL_HEAL_SUPPORT", "SPELL", "_HEAL_SUPPORT"},
		{"DAMAGE_SPLIT", "DAMAGE", "_SPLIT"},
		{"DAMAGE_SHIELD_MISSED", "DAMAGE_SHIELD", "_MISSED"},
	} {
		t.Run(tc.event, func(t *testing.T) {
			prefix, suffix, ok := l.Split(tc.event)
			if !ok || prefix != tc.prefix || suffix != tc.suffix {
				t.Fatalf("Split(%q) = %q %q ok=%v, want %q %q true",
					tc.event, prefix, suffix, ok, tc.prefix, tc.suffix)
			}
		})
	}
}

// TestSplitIsUnchangedForRetailV16 is the other half: the new rule must not
// move a single v16 event.
func TestSplitIsUnchangedForRetailV16(t *testing.T) {
	l := RetailV16()
	for _, tc := range []struct {
		event, prefix, suffix string
		ok                    bool
	}{
		{"SPELL_DAMAGE", "SPELL", "_DAMAGE", true},
		{"SPELL_PERIODIC_DAMAGE", "SPELL_PERIODIC", "_DAMAGE", true},
		{"SPELL_BUILDING_DAMAGE", "SPELL_BUILDING", "_DAMAGE", true},
		{"SWING_DAMAGE_LANDED", "SWING", "_DAMAGE_LANDED", true},
		{"SPELL_AURA_APPLIED_DOSE", "SPELL", "_AURA_APPLIED_DOSE", true},
		{"SPELL_MADE_UP", "SPELL", "_MADE_UP", false},
		{"NOT_AN_EVENT", "", "", false},
	} {
		t.Run(tc.event, func(t *testing.T) {
			prefix, suffix, ok := l.Split(tc.event)
			if ok != tc.ok || prefix != tc.prefix || suffix != tc.suffix {
				t.Fatalf("Split(%q) = %q %q ok=%v, want %q %q %v",
					tc.event, prefix, suffix, ok, tc.prefix, tc.suffix, tc.ok)
			}
		})
	}
}
```

Also update the existing `TestLookupPicksRetailV16AndHonoursTheAdvancedFlag`, whose last
assertion says v22 has no row. Replace:

```go
	if _, ok := Lookup(Header{Version: 22, ProjectID: 1, Advanced: true}); ok {
		t.Fatal("Lookup matched a version with no row; v22 is not implemented")
	}
```

with:

```go
	if _, ok := Lookup(Header{Version: 23, ProjectID: 1, Advanced: true}); ok {
		t.Fatal("Lookup matched a version with no row; v23 is not implemented")
	}
```

- [ ] **Step 2: Run the tests to verify they fail**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go test ./engine/layout/ -run 'TestLookupPicksRetailV22|TestRetailV22|TestSplitPrefers|TestRetailV16WidthsAreUnchanged' 2>&1 | head -20
```

Expected: compile failure, `undefined: RetailV22` and `l.Widths undefined (type Layout has no field or method Widths)`.

- [ ] **Step 3: Add the new `Suffix` flags, `Combatant.StatIndex` and `Widths` to `layout.go`**

In `logs/engine/layout/layout.go`, extend `Suffix`:

```go
// Suffix describes the parameters that follow the prefix.
type Suffix struct {
	Params      int  // parameters after the prefix and after the advanced block
	Advanced    bool // the event carries the advanced block
	AbsorbExtra int  // extra parameters present only when missType == "ABSORB"
	OffHand     bool // a trailing isOffHand that may be absent
	BaseAmount  bool // the damage suffix carries the unmodified base amount at index 1
	HealedToHP  bool // the heal suffix carries healedToHP at index 0
	// Tag marks a suffix whose line may end in the single-target / area
	// tag, the two-valued field combat-log version 22 writes as "ST" or
	// "AOE". It is optional because the same suffix is written with the
	// tag under a SPELL prefix and without it under SWING, and because
	// RANGE writes it on _DAMAGE but not on _MISSED. No prefix rule fits
	// all four cases, so the decoder recognises the tag by its value; see
	// isScopeTag in the event package.
	Tag bool
	// MissAmount marks a _MISSED suffix that carries one extra field, the
	// amount missed, when the miss type is BLOCK or RESIST. ABSORB's three
	// extras are counted by AbsorbExtra and are unaffected.
	MissAmount bool
	// AuraExtra marks an aura suffix that may carry a second trailing
	// number after the absorb size. Version 22 writes two where version 16
	// writes one; the second one's meaning is not pinned, so it is
	// counted and not read.
	AuraExtra bool
}
```

Extend `Combatant`:

```go
// Combatant locates the parts of a COMBATANT_INFO line.
type Combatant struct {
	Present        bool
	Params         int
	SpecIndex      int
	TalentIndex    int
	PvPTalentIndex int
	// BorrowIndex is the borrowed-power field, or 0 when the dialect
	// writes none. Zero is unambiguous: field 0 is always the event name.
	BorrowIndex int
	GearIndex   int
	AuraIndex   int
	// StatIndex maps a stat name to the field that holds it. The names are
	// the keys the decoder publishes in Combatant.Stats. A dialect that
	// does not write a stat leaves it out of the map rather than pointing
	// it at a field that means something else.
	StatIndex map[string]int
}
```

Add `Widths` next to `Width`:

```go
// Widths returns every total field count this row accepts for a shape, in
// ascending order. Width gives the base count the row's arithmetic
// produces; the optional trailing fields a dialect may or may not write
// (an absorb's extras, a Classic isOffHand, an aura's absorb size, version
// 22's single-target tag) turn that one number into a small set. The
// decoder checks membership in this set; the layout tests check the set
// against the counts measured on real logs.
func (l Layout) Widths(prefix, suffix string) []int {
	base, _ := l.Width(prefix, suffix)
	s := l.Suffixes[suffix]
	widths := []int{base}
	if s.AbsorbExtra > 0 {
		widths = append(widths, base+s.AbsorbExtra)
	}
	if s.MissAmount {
		widths = append(widths, base+1)
	}
	if s.OffHand {
		widths = append(widths, base+1)
	}
	if AuraCarriesAmount(suffix) {
		widths = append(widths, base+1)
		if s.AuraExtra {
			widths = append(widths, base+2)
		}
	}
	if s.Tag {
		for _, w := range append([]int(nil), widths...) {
			widths = append(widths, w+1)
		}
	}
	sort.Ints(widths)
	out := widths[:0]
	for i, w := range widths {
		if i == 0 || w != widths[i-1] {
			out = append(out, w)
		}
	}
	return out
}

// AuraCarriesAmount reports whether an aura suffix may be followed by the
// size of the absorb the aura provides. The dose suffixes carry a stack
// count in a field of their own and _AURA_BROKEN_SPELL carries the
// breaking spell, so neither takes the optional amount.
func AuraCarriesAmount(suffix string) bool {
	return strings.HasPrefix(suffix, "_AURA_") &&
		!strings.HasSuffix(suffix, "_DOSE") &&
		suffix != "_AURA_BROKEN_SPELL"
}
```

Replace `Split` with the longest-prefix-that-leaves-a-known-suffix rule:

```go
// Split separates an event name into its prefix and suffix. The winner is
// the longest registered prefix whose remainder is a suffix this row knows,
// which is what lets a row register both "SWING" and "SWING_DAMAGE_LANDED"
// without the longer one swallowing the shorter one's events. When no
// prefix leaves a known suffix the longest prefix match is reported with
// ok false, so the caller's error names the closest thing the row knows.
// Events handled by Specials must be checked first.
func (l Layout) Split(event string) (prefix, suffix string, ok bool) {
	best, bestSuffix := "", ""
	for p := range l.Prefixes {
		if len(p) <= len(best) || !strings.HasPrefix(event, p) {
			continue
		}
		if _, known := l.Suffixes[event[len(p):]]; known {
			best, bestSuffix = p, event[len(p):]
		}
	}
	if best != "" {
		return best, bestSuffix, true
	}
	for p := range l.Prefixes {
		if strings.HasPrefix(event, p) && len(p) > len(best) {
			best = p
		}
	}
	if best == "" {
		return "", "", false
	}
	return best, event[len(best):], false
}
```

Register the new row, most specific first:

```go
// rows is the table, most specific first.
var rows = []Layout{RetailV22(), RetailV16(), ClassicWiki()}
```

- [ ] **Step 4: Give the v16 and classic rows an explicit `StatIndex`**

In `logs/engine/layout/retail.go`, inside `RetailV16()`'s `Combatant` literal, add the
map that the decoder currently hard-codes:

```go
		Combatant: Combatant{
			Present:        true,
			Params:         34,
			SpecIndex:      24,
			TalentIndex:    25,
			PvPTalentIndex: 26,
			BorrowIndex:    27,
			GearIndex:      28,
			AuraIndex:      29,
			StatIndex: map[string]int{
				"strength": 3, "agility": 4, "stamina": 5, "intellect": 6,
				"dodge": 7, "parry": 8, "block": 9, "crit": 10, "speed": 13,
				"lifesteal": 14, "haste": 15, "avoidance": 18, "mastery": 19,
				"versatility": 20, "armor": 23,
			},
		},
```

`ClassicWiki()` sets `Combatant{}` with `Present: false`, so it needs no map.

- [ ] **Step 5: Add `RetailV22()`**

Append to `logs/engine/layout/retail.go`:

```go
// RetailV22 is the modern retail dialect, COMBAT_LOG_VERSION 22,
// PROJECT_ID 1. Every count below was measured over 89 real logs,
// 26,030,980 lines, builds 12.0.5, 12.0.7 and 12.1.0, all with advanced
// logging on. The field-by-field derivation is in
// docs/superpowers/plans/2026-09-16-retail-v22-layout.md; the fields that
// could not be identified are in docs/ledger/2026-09-16-retail-v22.md.
//
// Three things changed from v16 and drive everything here: the advanced
// block grew from 17 fields to 19, spell-prefixed damage and miss lines
// gained a trailing single-target / area tag, and COMBATANT_INFO's stat
// fields all moved one to the right.
func RetailV22() Layout {
	return Layout{
		Name:      "retail-v22",
		Version:   22,
		ProjectID: 1,
		Advanced:  19,
		StampYear: true,
		StampZone: true,
		Verified:  true,
		Prefixes: map[string]int{
			"SWING":          0,
			"RANGE":          3,
			"SPELL_PERIODIC": 3,
			"SPELL_BUILDING": 3,
			"SPELL":          3,
			// DAMAGE_SPLIT is a spell-prefixed damage line whose name
			// happens not to start with SPELL.
			"DAMAGE": 3,
			// DAMAGE_SHIELD_MISSED likewise. The longer prefix wins over
			// "DAMAGE" because Split prefers the one that leaves a known
			// suffix, and "_SHIELD_MISSED" is not a suffix.
			"DAMAGE_SHIELD": 3,
			// The four prefixes below exist only so that the "_SUPPORT"
			// suffix attaches to the right spell-triple offset. A support
			// line names the supporting spell, not the supported one, so
			// even SWING_DAMAGE_LANDED_SUPPORT carries a spell triple and
			// is 42 fields wide where SWING_DAMAGE_LANDED is 38.
			"SPELL_DAMAGE":          3,
			"SPELL_PERIODIC_DAMAGE": 3,
			"RANGE_DAMAGE":          3,
			"SWING_DAMAGE_LANDED":   3,
		},
		Suffixes: map[string]Suffix{
			// amount, baseAmount, overkill, school, resisted, blocked,
			// absorbed, critical, glancing, crushing, then the optional
			// "ST" / "AOE" tag.
			"_DAMAGE":        {Params: 10, Advanced: true, BaseAmount: true, Tag: true},
			"_DAMAGE_LANDED": {Params: 10, Advanced: true, BaseAmount: true, Tag: true},
			"_SPLIT":         {Params: 10, Advanced: true, BaseAmount: true, Tag: true},
			// The same ten damage fields, then the supporting player's
			// GUID in place of the tag.
			"_SUPPORT": {Params: 11, Advanced: true, BaseAmount: true},
			// healedToHP, amount, overheal, absorbed, critical.
			"_HEAL": {Params: 5, Advanced: true, HealedToHP: true},
			// The same five, then the supporting player's GUID.
			"_HEAL_SUPPORT": {Params: 6, Advanced: true, HealedToHP: true},
			// missType, isOffHand; three more on ABSORB, one more on
			// BLOCK and RESIST, and the tag after whichever of those
			// applies.
			"_MISSED": {Params: 2, AbsorbExtra: 3, MissAmount: true, Tag: true},
			// amount, overEnergize, powerType, maxPower.
			"_ENERGIZE": {Params: 4, Advanced: true},
			"_DRAIN":    {Params: 4, Advanced: true},
			"_LEECH":    {Params: 4, Advanced: true},
			// auraType, then the absorb size, then one more number whose
			// meaning is not pinned.
			"_AURA_APPLIED":      {Params: 1, AuraExtra: true},
			"_AURA_REMOVED":      {Params: 1, AuraExtra: true},
			"_AURA_REFRESH":      {Params: 1, AuraExtra: true},
			"_AURA_BROKEN":       {Params: 1, AuraExtra: true},
			"_AURA_APPLIED_DOSE": {Params: 2},
			"_AURA_REMOVED_DOSE": {Params: 2},
			"_AURA_BROKEN_SPELL": {Params: 4},
			"_INTERRUPT":         {Params: 3},
			"_DISPEL_FAILED":     {Params: 3},
			"_DISPEL":            {Params: 4},
			"_STOLEN":            {Params: 4},
			"_CAST_START":        {Params: 0},
			"_CAST_SUCCESS":      {Params: 0, Advanced: true},
			"_CAST_FAILED":       {Params: 1},
			"_SUMMON":            {Params: 0},
			"_CREATE":            {Params: 0},
			"_RESURRECT":         {Params: 0},
			"_INSTAKILL":         {Params: 1},
			"_EXTRA_ATTACKS":     {Params: 1},
			"_DURABILITY_DAMAGE": {Params: 0},
			// Evoker empowered casts: start carries nothing extra, end and
			// interrupt carry the empowerment stage reached.
			"_EMPOWER_START":     {Params: 0},
			"_EMPOWER_END":       {Params: 1},
			"_EMPOWER_INTERRUPT": {Params: 1},
		},
		Specials: map[string]Special{
			"COMBAT_LOG_VERSION":  {Widths: []int{8}},
			"UNIT_DIED":           {Widths: []int{10}},
			"UNIT_DESTROYED":      {Widths: []int{10}},
			"UNIT_DISSIPATES":     {Widths: []int{10}},
			"PARTY_KILL":          {Widths: []int{10}},
			"SPELL_ABSORBED":      {Widths: []int{19, 22}},
			"SPELL_HEAL_ABSORBED": {Widths: []int{21}},
			// The absorb shapes plus the supporting player's GUID.
			"SPELL_ABSORBED_SUPPORT": {Widths: []int{20, 23}},
			"ENCOUNTER_START":        {Widths: []int{6}},
			"ENCOUNTER_END":          {Widths: []int{6}},
			"ZONE_CHANGE":            {Widths: []int{4}},
			"MAP_CHANGE":             {Widths: []int{7}},
			"CHALLENGE_MODE_START":   {Widths: []int{6}},
			"CHALLENGE_MODE_END":     {Widths: []int{5}},
			"ENCHANT_APPLIED":        {Widths: []int{12}},
			"ENCHANT_REMOVED":        {Widths: []int{12}},
			"EMOTE":                  {Widths: []int{6}},
			"COMBATANT_INFO":         {Widths: []int{34}},
			// 9 header + 19 advanced + environmental type + 10 damage.
			"ENVIRONMENTAL_DAMAGE": {Widths: []int{39}},
			// instanceID, bracket, matchType, isRated.
			"ARENA_MATCH_START": {Widths: []int{5}},
			// winningTeam, duration, newRatingTeam1, newRatingTeam2.
			"ARENA_MATCH_END": {Widths: []int{5}},
			// playerGUID, amount cleared.
			"STAGGER_CLEAR": {Widths: []int{3}},
			// playerGUID, spellID, amount prevented.
			"STAGGER_PREVENTED": {Widths: []int{4}},
			// uiMapID, markerIndex, x, y.
			"WORLD_MARKER_PLACED": {Widths: []int{5}},
			// markerIndex.
			"WORLD_MARKER_REMOVED": {Widths: []int{2}},
		},
		Combatant: Combatant{
			Present:        true,
			Params:         34,
			SpecIndex:      25,
			TalentIndex:    26,
			PvPTalentIndex: 27,
			// Version 22 writes no borrowed-power field; the line is the
			// same width as v16's because the stat block gained one field
			// and this one went away.
			BorrowIndex: 0,
			GearIndex:   28,
			AuraIndex:   29,
			// Every stat sits one field later than in v16: v22 writes four
			// defensive ratings where v16 writes dodge, parry and block.
			// All four are zero in every sample, so the fourth is counted
			// and not named; see the ledger.
			StatIndex: map[string]int{
				"strength": 3, "agility": 4, "stamina": 5, "intellect": 6,
				"dodge": 7, "parry": 8, "block": 9, "crit": 11, "speed": 14,
				"lifesteal": 15, "haste": 16, "avoidance": 19, "mastery": 20,
				"versatility": 21, "armor": 24,
			},
		},
	}
}
```

- [ ] **Step 6: Run the layout tests**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go test ./engine/layout/ 2>&1 | tail -20
```

Expected: `ok  github.com/jhunthrop/foreversixty/logs/engine/layout`.

If `TestRetailV22AcceptsEveryMeasuredWidth` reports `SPELL_AURA_REFRESH width 15 is not
accepted`, note that the corpus shows only 13 and 14 for `_AURA_REFRESH` — the row
accepts 15 because the flag is per-suffix. That direction (row accepts more than
measured) is fine; the test only fails when a measured width is rejected.

- [ ] **Step 7: Run the whole suite to prove nothing v16 moved**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go test ./... 2>&1 | tail -20
git -C /Users/jh/code/forever/.worktrees/retail-v22 diff --exit-code -- logs/engine/summary/testdata/v16.summary.json.golden && echo "v16 golden unchanged"
```

Expected: every package `ok`, then `v16 golden unchanged`.

- [ ] **Step 8: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
cat > /tmp/v22-msg-2.txt <<'MSG'
feat(logs): a verified retail-v22 layout row, selected by header version

The advanced block is 19 fields, the timestamp carries a year and a
zone, and COMBATANT_INFO's stats all sit one field later. Three shared
mechanisms carry the differences: Layout.Widths, which turns a shape's
base width into the set of widths a dialect may write; a Split that
prefers the longest prefix leaving a known suffix, so SWING and
SWING_DAMAGE_LANDED can both be prefixes; and a per-dialect
COMBATANT_INFO stat index map in place of hard-coded offsets.

The v16 row sets none of the new flags and its accepted widths are
asserted unchanged.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
git add logs/engine/layout/layout.go logs/engine/layout/retail.go logs/engine/layout/layout_test.go
```

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22 && git commit -F /tmp/v22-msg-2.txt
```

---

## Task 3: Decode every v22 shape

**Files:**
- Modify: `logs/engine/event/event.go`
- Modify: `logs/engine/event/decode.go`
- Modify: `logs/engine/event/special.go`
- Create: `logs/engine/event/decode_v22_test.go`
- Append: `docs/ledger/2026-09-16-retail-v22.md`

**Interfaces:**
- Consumes: `layout.RetailV22()`, `(Layout).Widths`, `Combatant.StatIndex`, `Suffix.Tag`, `Suffix.MissAmount`, `Suffix.AuraExtra` from Task 2; the two testdata files from Task 1.
- Produces:
  - `event.Advanced` gains `Versatility int64` and `Unknown8 int64`.
  - `event.Event` gains `Scope string` (`"ST"`, `"AOE"` or empty) and `Supporter string` (the supporting player's GUID, empty when the event is not a `_SUPPORT` variant).
  - `event.Kind` gains `DamageSplit`, `EmpowerStart`, `EmpowerEnd`, `ArenaMatchStart`, `ArenaMatchEnd`, `StaggerClear`, `StaggerPrevented`, `WorldMarker`.
  - No change to the parquet schema and no change to any summary field.

**How the summary treats support damage.** A `_SUPPORT` line is not extra damage: it is
the game attributing a slice of damage *already reported on another line* to the player
whose buff caused it (Augmentation Evoker's Ebon Might, Shifting Sands). Counting it
would double-count every augmented hit. **This plan therefore has the summary ignore
`_SUPPORT` events entirely**: they decode without error, carry `Supporter`, and the
accumulator skips them. Attributing support damage to the supporter is a report-design
question, not a parsing question, and belongs to whichever lane owns the Augmentation
view. This decision is written into the ledger in Step 8.

- [ ] **Step 1: Write the failing decode test**

Create `logs/engine/event/decode_v22_test.go`:

```go
// logs/engine/event/decode_v22_test.go
package event

import (
	"os"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// v22Base is unused by v22 itself: its timestamps carry a year. It is
// passed anyway because NewDecoder wants a base, and a wrong one must not
// be able to leak into a dated dialect.
var v22Base = time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)

func decodeV22File(t *testing.T, path string) (map[string][]Event, []Event) {
	t.Helper()
	text, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	d := NewDecoder(layout.RetailV22(), v22Base)
	byName := map[string][]Event{}
	var all []Event
	l := lexer.New()
	emit := func(ln lexer.Line) error {
		e := d.Decode(ln)
		byName[e.Name] = append(byName[e.Name], e)
		all = append(all, e)
		return nil
	}
	if err := l.Feed(text, 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := l.Flush(emit); err != nil {
		t.Fatal(err)
	}
	return byName, all
}

func TestEveryV22LineDecodesWithoutError(t *testing.T) {
	for _, path := range []string{"testdata/v22.log", "testdata/v22-shapes.log"} {
		t.Run(path, func(t *testing.T) {
			_, all := decodeV22File(t, path)
			if len(all) == 0 {
				t.Fatal("no lines decoded")
			}
			for _, e := range all {
				if e.Kind == ParseError {
					t.Errorf("line %d (%s) failed: %s\n%s", e.Line, e.Name, e.Error, e.Raw)
				}
				if e.Kind == Unknown {
					t.Errorf("line %d (%s) was not recognised", e.Line, e.Name)
				}
			}
		})
	}
}

func TestV22SpellDamageReadsTheNineteenFieldAdvancedBlockAndTheScopeTag(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22.log")
	e := one(t, by, "SPELL_DAMAGE", 0)
	if e.Kind != Damage {
		t.Fatalf("kind = %v, want Damage", e.Kind)
	}
	if !e.Adv.OK {
		t.Fatal("the advanced block was not read")
	}
	if e.Adv.PositionX == 0 || e.Adv.PositionY == 0 {
		t.Errorf("position = %g, %g; both zero means the block was read at the v16 offsets",
			e.Adv.PositionX, e.Adv.PositionY)
	}
	if e.Adv.Level < 200 || e.Adv.Level > 500 {
		t.Errorf("item level = %d, want a plausible level-80 item level", e.Adv.Level)
	}
	if e.Scope != "ST" && e.Scope != "AOE" {
		t.Errorf("scope = %q, want ST or AOE", e.Scope)
	}
	if e.OffHand.OK {
		t.Error("the scope tag was read as isOffHand")
	}
	if !e.Amount.OK || !e.BaseAmount.OK {
		t.Errorf("amount=%+v baseAmount=%+v, both must be read", e.Amount, e.BaseAmount)
	}
}

func TestV22SwingDamageHasNoScopeTag(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22.log")
	e := one(t, by, "SWING_DAMAGE", 0)
	if e.Kind != Damage {
		t.Fatalf("kind = %v, want Damage", e.Kind)
	}
	if e.Scope != "" {
		t.Errorf("scope = %q, want empty: SWING lines carry no tag", e.Scope)
	}
	if !e.Adv.OK {
		t.Error("the advanced block was not read")
	}
}

func TestV22PowerFieldsComeFromTheShiftedOffsets(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22.log")
	e := one(t, by, "SPELL_ENERGIZE", 0)
	if e.Kind != Energize {
		t.Fatalf("kind = %v, want Energize", e.Kind)
	}
	// The energize suffix names the same power the advanced block
	// describes, which is what pinned the block's offsets in the first
	// place.
	if e.Adv.PowerType != e.PowerType.V {
		t.Errorf("advanced powerType = %d, suffix powerType = %d; the block is misaligned",
			e.Adv.PowerType, e.PowerType.V)
	}
	if e.Adv.MaxPower != e.MaxPower.V {
		t.Errorf("advanced maxPower = %d, suffix maxPower = %d; the block is misaligned",
			e.Adv.MaxPower, e.MaxPower.V)
	}
}

func TestV22MissedReadsTheAbsorbExtrasPastTheTag(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22.log")
	var absorbed *Event
	for i := range by["SPELL_MISSED"] {
		if by["SPELL_MISSED"][i].MissType == "ABSORB" {
			absorbed = &by["SPELL_MISSED"][i]
			break
		}
	}
	if absorbed == nil {
		t.Fatal("the excerpt has no absorbed SPELL_MISSED")
	}
	if !absorbed.Amount.OK || !absorbed.BaseAmount.OK {
		t.Errorf("amount=%+v baseAmount=%+v; the absorb extras were not read past the tag",
			absorbed.Amount, absorbed.BaseAmount)
	}
	if absorbed.Scope != "ST" && absorbed.Scope != "AOE" {
		t.Errorf("scope = %q, want ST or AOE", absorbed.Scope)
	}
}

func TestV22BlockedMissCarriesAnAmount(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	e := one(t, by, "SPELL_MISSED", 0)
	if e.MissType != "BLOCK" && e.MissType != "RESIST" {
		t.Fatalf("the shapes file's SPELL_MISSED is %q, want the 16-wide BLOCK or RESIST form", e.MissType)
	}
	if !e.Amount.OK {
		t.Error("the amount missed was not read")
	}
	if e.Scope != "ST" && e.Scope != "AOE" {
		t.Errorf("scope = %q, want ST or AOE", e.Scope)
	}
}

func TestV22RangeMissedHasNoScopeTag(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	e := one(t, by, "RANGE_MISSED", 0)
	if e.Kind != Missed {
		t.Fatalf("kind = %v, want Missed", e.Kind)
	}
	if e.Scope != "" {
		t.Errorf("scope = %q, want empty: RANGE_MISSED carries no tag", e.Scope)
	}
}

func TestV22SupportEventsNameTheSupporter(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	for _, name := range []string{
		"SPELL_DAMAGE_SUPPORT", "SPELL_PERIODIC_DAMAGE_SUPPORT",
		"RANGE_DAMAGE_SUPPORT", "SWING_DAMAGE_LANDED_SUPPORT",
		"SPELL_HEAL_SUPPORT", "SPELL_PERIODIC_HEAL_SUPPORT",
		"SPELL_ABSORBED_SUPPORT",
	} {
		t.Run(name, func(t *testing.T) {
			e := one(t, by, name, 0)
			if e.Kind == ParseError || e.Kind == Unknown {
				t.Fatalf("kind = %v: %s", e.Kind, e.Error)
			}
			if e.Supporter == "" {
				t.Error("the supporting player's GUID was not read")
			}
			if e.Scope != "" {
				t.Errorf("scope = %q; a support line ends in a GUID, not a tag", e.Scope)
			}
		})
	}
}

func TestV22DamageSplitIsDamage(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	e := one(t, by, "DAMAGE_SPLIT", 0)
	if e.Kind != DamageSplit {
		t.Fatalf("kind = %v, want DamageSplit", e.Kind)
	}
	if e.Spell.ID == 0 || e.Spell.Name == "" {
		t.Errorf("spell = %+v, want the splitting spell", e.Spell)
	}
	if !e.Adv.OK {
		t.Error("the advanced block was not read")
	}
	if !e.Amount.OK {
		t.Error("the amount was not read")
	}
}

func TestV22EmpowerAndArenaAndStaggerAndMarkers(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	for name, want := range map[string]Kind{
		"SPELL_EMPOWER_START":     EmpowerStart,
		"SPELL_EMPOWER_END":       EmpowerEnd,
		"SPELL_EMPOWER_INTERRUPT": EmpowerEnd,
		"ARENA_MATCH_START":       ArenaMatchStart,
		"ARENA_MATCH_END":         ArenaMatchEnd,
		"STAGGER_CLEAR":           StaggerClear,
		"STAGGER_PREVENTED":       StaggerPrevented,
		"WORLD_MARKER_PLACED":     WorldMarker,
		"WORLD_MARKER_REMOVED":    WorldMarker,
	} {
		t.Run(name, func(t *testing.T) {
			e := one(t, by, name, 0)
			if e.Kind != want {
				t.Fatalf("kind = %v, want %v (error: %s)", e.Kind, want, e.Error)
			}
		})
	}
}

func TestV22CombatantInfoReadsTheShiftedStats(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22.log")
	e := one(t, by, "COMBATANT_INFO", 0)
	if e.Kind != CombatantInfo {
		t.Fatalf("kind = %v: %s", e.Kind, e.Error)
	}
	c := e.Combatant
	if c == nil {
		t.Fatal("no combatant")
	}
	if c.SpecID <= 0 || c.SpecID > 2000 {
		t.Errorf("specID = %d, want a plausible spec id; the index may still be v16's 24", c.SpecID)
	}
	if len(c.Gear) == 0 {
		t.Error("no gear; the gear index is wrong")
	}
	if c.ItemLevel < 200 || c.ItemLevel > 500 {
		t.Errorf("item level = %d, want a plausible level-80 item level", c.ItemLevel)
	}
	if len(c.Talents) == 0 {
		t.Error("no talents; the talent index is wrong")
	}
	if c.Borrowed != "" {
		t.Errorf("borrowed = %q, want empty: v22 writes no borrowed-power field", c.Borrowed)
	}
	// The advanced block's armor is the same number COMBATANT_INFO
	// reports, which is what pinned the stat shift.
	if c.Stats["armor"] <= 0 {
		t.Errorf("armor = %d, want a positive value", c.Stats["armor"])
	}
}

func TestV22AuraWithTwoTrailingNumbers(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	e := one(t, by, "SPELL_AURA_APPLIED", 0)
	if e.Kind != AuraApplied {
		t.Fatalf("kind = %v: %s", e.Kind, e.Error)
	}
	if e.AuraType != "BUFF" && e.AuraType != "DEBUFF" {
		t.Errorf("auraType = %q", e.AuraType)
	}
}

func TestTheV16FixtureStillDecodesTheSameWayAfterTheV22Work(t *testing.T) {
	_, all := decodeAll(t)
	if len(all) != 33 {
		t.Fatalf("decoded %d lines, want 33", len(all))
	}
	for _, e := range all {
		if e.Kind == ParseError || e.Kind == Unknown {
			t.Errorf("line %d (%s): kind=%v %s", e.Line, e.Name, e.Kind, e.Error)
		}
		if e.Scope != "" {
			t.Errorf("line %d (%s) got a scope tag on a v16 line", e.Line, e.Name)
		}
		if e.Supporter != "" {
			t.Errorf("line %d (%s) got a supporter on a v16 line", e.Line, e.Name)
		}
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go test ./engine/event/ -run 'V22' 2>&1 | head -20
```

Expected: compile failure, `e.Scope undefined`, `e.Supporter undefined`, `undefined: DamageSplit`.

- [ ] **Step 3: Add the new Event and Advanced fields and kinds**

In `logs/engine/event/event.go`, extend `Advanced` with the two v22-only fields:

```go
type Advanced struct {
	OK           bool
	InfoGUID     string
	OwnerGUID    string
	CurrentHP    int64
	MaxHP        int64
	AttackPower  int64
	SpellPower   int64
	Armor        int64
	// Versatility is written only by combat-log version 22, at block
	// offset 7. It tracks versatility buffs and nothing else, and across
	// twelve players in two logs it is 1.8513x the versatility rating
	// COMBATANT_INFO reports for the same character. Both logs are arena
	// logs, where secondary stats are scaled, so the likeliest reading is
	// that this is the unscaled rating; see
	// docs/ledger/2026-09-16-retail-v22.md.
	Versatility int64
	// Unknown8 is block offset 8 in combat-log version 22. Zero for almost
	// every unit; a small per-unit constant for a few. It matches no
	// COMBATANT_INFO field and moves with nothing. It is read so that the
	// nineteen fields account for themselves; see the ledger.
	Unknown8     int64
	Absorb       int64
	PowerType    int64
	CurrentPower int64
	MaxPower     int64
	PowerCost    int64
	PositionX    float64
	PositionY    float64
	UIMapID      int64
	Facing       float64
	// Level is the unit's equipped item level, not its character level:
	// the version 16 sample shows 189-191 on a level-60 character. The
	// name predates the measurement and is kept because renaming it would
	// touch v16 decoding; see the ledger.
	Level int64
}
```

Add to the `Event` struct, beside `MissType` / `AuraType` / `FailedType`:

```go
	// Scope is the single-target / area tag combat-log version 22 writes
	// at the end of spell-prefixed damage and miss lines: "ST", "AOE", or
	// empty on a dialect or an event that does not write one.
	Scope string
	// Supporter is the GUID of the player whose buff caused a _SUPPORT
	// line's damage or healing. Empty on every other event. A _SUPPORT
	// line restates damage already reported elsewhere, so the summary
	// ignores these events rather than double-counting them.
	Supporter string
```

Add the new kinds to the `Kind` enum, after the existing ones so no existing
value changes, and add their names to the kind-name map:

```go
	DamageSplit
	EmpowerStart
	EmpowerEnd
	ArenaMatchStart
	ArenaMatchEnd
	StaggerClear
	StaggerPrevented
	WorldMarker
```

```go
	DamageSplit: "damage_split", EmpowerStart: "empower_start",
	EmpowerEnd: "empower_end", ArenaMatchStart: "arena_match_start",
	ArenaMatchEnd: "arena_match_end", StaggerClear: "stagger_clear",
	StaggerPrevented: "stagger_prevented", WorldMarker: "world_marker",
```

- [ ] **Step 4: Teach `readAdvanced` the 19-field block**

In `logs/engine/event/decode.go`, replace `readAdvanced`:

```go
// readAdvanced reads the advanced block. The two layouts differ by two
// fields inserted after armor, which pushes absorb and every power field
// down; reading a 19-field block with the 17-field offsets silently yields
// a position of (0, 0) and a power type taken from the absorb slot, so the
// length picks the mapping rather than an index guard.
func readAdvanced(f []string) Advanced {
	if len(f) == 19 {
		return Advanced{
			OK:           true,
			InfoGUID:     f[0],
			OwnerGUID:    f[1],
			CurrentHP:    intOf(f[2]),
			MaxHP:        intOf(f[3]),
			AttackPower:  intOf(f[4]),
			SpellPower:   intOf(f[5]),
			Armor:        intOf(f[6]),
			Versatility:  intOf(f[7]),
			Unknown8:     intOf(f[8]),
			Absorb:       intOf(f[9]),
			PowerType:    intOf(f[10]),
			CurrentPower: intOf(f[11]),
			MaxPower:     intOf(f[12]),
			PowerCost:    intOf(f[13]),
			PositionX:    floatOf(f[14]),
			PositionY:    floatOf(f[15]),
			UIMapID:      intOf(f[16]),
			Facing:       floatOf(f[17]),
			Level:        intOf(f[18]),
		}
	}
	if len(f) < 17 {
		return Advanced{}
	}
	return Advanced{
		OK:           true,
		InfoGUID:     f[0],
		OwnerGUID:    f[1],
		CurrentHP:    intOf(f[2]),
		MaxHP:        intOf(f[3]),
		AttackPower:  intOf(f[4]),
		SpellPower:   intOf(f[5]),
		Armor:        intOf(f[6]),
		Absorb:       intOf(f[7]),
		PowerType:    intOf(f[8]),
		CurrentPower: intOf(f[9]),
		MaxPower:     intOf(f[10]),
		PowerCost:    intOf(f[11]),
		PositionX:    floatOf(f[12]),
		PositionY:    floatOf(f[13]),
		UIMapID:      intOf(f[14]),
		Facing:       floatOf(f[15]),
		Level:        intOf(f[16]),
	}
}
```

- [ ] **Step 5: Replace the ad-hoc width chain with the layout's width set, and strip the trailing field**

In `decodeStandard`, replace the block from `spec := d.lay.Suffixes[suffix]` down to the
`if len(p) != want` check with:

```go
	p := ln.Params
	spec := d.lay.Suffixes[suffix]
	want, advAt := d.lay.Width(prefix, suffix)

	// The row says which widths this shape may have: the base count plus
	// whichever optional trailing fields the dialect writes. The exact
	// width in hand then fixes where the suffix's own fields stop.
	allowed := d.lay.Widths(prefix, suffix)
	ok := false
	for _, w := range allowed {
		if len(p) == w {
			ok = true
		}
	}
	if !ok {
		return fail(e, ln, fmt.Sprintf("%s has %d fields, layout %q allows %v",
			e.Name, len(p), d.lay.Name, allowed))
	}
	want = len(p)
```

Keep the rest of the function as it is up to `rest := p[i:]`, then insert the trailing-field
strip immediately after it:

```go
	rest := p[i:]

	// Version 22 ends a spell-prefixed damage or miss line with a
	// single-target / area tag. It cannot be located by counting, because
	// RANGE_DAMAGE writes it and RANGE_MISSED does not, so it is
	// recognised by value: "ST" and "AOE" are the only two strings it ever
	// holds, and every field it could be confused with (critical,
	// isOffHand, crushing) holds "nil", "0" or "1".
	if spec.Tag && len(rest) > 0 && isScopeTag(rest[len(rest)-1]) {
		e.Scope = rest[len(rest)-1]
		rest = rest[:len(rest)-1]
	}
```

Add the helper next to `boolOf`:

```go
// isScopeTag reports whether a field is the single-target / area tag.
func isScopeTag(s string) bool { return s == "ST" || s == "AOE" }
```

- [ ] **Step 6: Extend the suffix switch**

In the `switch suffix` of `decodeStandard`, change the `_MISSED` case to read the
BLOCK / RESIST amount, and add the new suffixes:

```go
	case "_MISSED":
		e.Kind = Missed
		e.MissType = rest[0]
		e.OffHand = boolOf(rest[1])
		switch {
		case len(rest) >= 5:
			// ABSORB: amountMissed, baseAmount, critical.
			e.Amount, e.BaseAmount, e.Critical = optInt(rest[2]), optInt(rest[3]), boolOf(rest[4])
		case len(rest) == 4:
			e.Amount, e.Critical = optInt(rest[2]), boolOf(rest[3])
		case len(rest) == 3:
			// BLOCK and RESIST: the amount missed, and nothing else.
			e.Amount = optInt(rest[2])
		}
	case "_SPLIT":
		e.Kind = DamageSplit
		readDamage(&e, rest, spec)
	case "_SUPPORT":
		// The supporting player's GUID is the last field; the ten before
		// it are an ordinary damage suffix.
		e.Kind = Damage
		e.Supporter = rest[len(rest)-1]
		readDamage(&e, rest[:len(rest)-1], spec)
	case "_HEAL_SUPPORT":
		e.Kind = Heal
		e.Supporter = rest[len(rest)-1]
		readHeal(&e, rest[:len(rest)-1], spec)
	case "_EMPOWER_START":
		e.Kind = EmpowerStart
	case "_EMPOWER_END", "_EMPOWER_INTERRUPT":
		e.Kind = EmpowerEnd
		e.Stacks = optInt(rest[0]) // the empowerment stage reached
```

Extend the aura case to tolerate the second trailing number:

```go
	case "_AURA_APPLIED", "_AURA_REMOVED", "_AURA_REFRESH", "_AURA_BROKEN":
		e.Kind = map[string]Kind{
			"_AURA_APPLIED": AuraApplied, "_AURA_REMOVED": AuraRemoved,
			"_AURA_REFRESH": AuraRefresh, "_AURA_BROKEN": AuraBroken,
		}[suffix]
		e.AuraType = rest[0]
		if len(rest) > 1 {
			e.Absorbed = optInt(rest[1])
		}
		// A version 22 aura line may carry one more number after the
		// absorb size. Its meaning is not pinned, so it is counted by the
		// layout and left unread; see the ledger.
```

Extend `suffixNeeds` so the new suffixes declare what the switch indexes:

```go
	case "_SPLIT":
		return 10
	case "_SUPPORT":
		return 11
	case "_HEAL_SUPPORT":
		return 6
	case "_EMPOWER_END", "_EMPOWER_INTERRUPT":
		return 1
```

- [ ] **Step 7: Make `readCombatant` layout-driven and add the new specials**

In `logs/engine/event/special.go`, replace the hard-coded `Stats` map and the
`BorrowIndex` handling in `readCombatant`:

```go
	for _, at := range []int{c.SpecIndex, c.TalentIndex, c.PvPTalentIndex, c.GearIndex, c.AuraIndex} {
		if at <= 0 || at >= len(p) {
			return fail(e, ln, fmt.Sprintf("COMBATANT_INFO has %d fields, layout %q indexes field %d",
				len(p), d.lay.Name, at))
		}
	}
	// BorrowIndex is zero on a dialect that writes no borrowed-power
	// field, and field zero is always the event name, so zero is an
	// unambiguous "absent" rather than a missing bounds check.
	if c.BorrowIndex < 0 || c.BorrowIndex >= len(p) {
		return fail(e, ln, fmt.Sprintf("COMBATANT_INFO has %d fields, layout %q indexes field %d",
			len(p), d.lay.Name, c.BorrowIndex))
	}
	stats := make(map[string]int64, len(c.StatIndex))
	for name, at := range c.StatIndex {
		if at <= 0 || at >= len(p) {
			return fail(e, ln, fmt.Sprintf("COMBATANT_INFO has %d fields, layout %q reads %s from field %d",
				len(p), d.lay.Name, name, at))
		}
		stats[name] = intOf(p[at])
	}
	e.Kind = CombatantInfo
	info := &Combatant{
		GUID:       p[1],
		Faction:    intOf(p[2]),
		SpecID:     intOf(p[c.SpecIndex]),
		Stats:      stats,
		Talents:    intList(p[c.TalentIndex]),
		PvPTalents: intList(p[c.PvPTalentIndex]),
		Gear:       gearList(p[c.GearIndex]),
		Auras:      auraList(p[c.AuraIndex]),
	}
	if c.BorrowIndex > 0 {
		info.Borrowed = p[c.BorrowIndex]
	}
```

Add the new specials to `specialNeeds`:

```go
	"SPELL_ABSORBED_SUPPORT": layout.BaseParams,
	"ARENA_MATCH_START":      5,
	"ARENA_MATCH_END":        5,
	"STAGGER_CLEAR":          3,
	"STAGGER_PREVENTED":      4,
	"WORLD_MARKER_PLACED":    5,
	"WORLD_MARKER_REMOVED":   2,
```

Add their branches to `decodeSpecial`'s switch:

```go
	case "SPELL_ABSORBED_SUPPORT":
		// The same shape as SPELL_ABSORBED with the supporting player's
		// GUID appended; readAbsorbed walks from the front and stops at
		// the fields it knows, so the GUID is taken off first.
		short := ln
		short.Params = p[:len(p)-1]
		out := d.readAbsorbed(e, short)
		out.Supporter = p[len(p)-1]
		out.Raw = ln.Raw
		return out
	case "ARENA_MATCH_START":
		e.Kind = ArenaMatchStart
		e.Zone = &Zone{ID: intOf(p[1])}
		e.Amount = optInt(p[2])  // bracket
		e.ItemName = nilless(p[3]) // match type, e.g. "Rated Solo Shuffle"
		e.Critical = OptBool{V: intOf(p[4]) == 1, OK: true}
	case "ARENA_MATCH_END":
		e.Kind = ArenaMatchEnd
		e.Amount = optInt(p[1]) // winning team
		e.Total = optInt(p[2])  // duration in seconds
	case "STAGGER_CLEAR":
		e.Kind = StaggerClear
		e.Source = Unit{GUID: p[1]}
		e.Amount = optInt(p[2])
	case "STAGGER_PREVENTED":
		e.Kind = StaggerPrevented
		e.Source = Unit{GUID: p[1]}
		e.Spell = Spell{ID: intOf(p[2])}
		e.Amount = optInt(p[3])
	case "WORLD_MARKER_PLACED":
		e.Kind = WorldMarker
		e.Zone = &Zone{ID: intOf(p[1])}
		e.Amount = optInt(p[2])
		e.Critical = OptBool{V: true, OK: true} // placed
	case "WORLD_MARKER_REMOVED":
		e.Kind = WorldMarker
		e.Amount = optInt(p[1])
		e.Critical = OptBool{V: false, OK: true} // removed
```

`readAbsorbed` is a method on `*Decoder` taking `(Event, lexer.Line)`; check its exact
signature before writing the `SPELL_ABSORBED_SUPPORT` branch and match it.

- [ ] **Step 8: Have the summary ignore `_SUPPORT` events**

In `logs/engine/summary/summary.go`, at the top of `(*Accumulator).Add`, before any
bucket is touched:

```go
	// A _SUPPORT line restates damage or healing already reported on
	// another line, attributing a slice of it to the player whose buff
	// caused it (an Augmentation Evoker's Ebon Might, for instance).
	// Counting it would double every augmented hit, so the accumulator
	// skips it. Attributing support damage to the supporter is a report
	// question and belongs to whichever lane owns that view.
	if e.Supporter != "" {
		return
	}
```

Append to `docs/ledger/2026-09-16-retail-v22.md`:

```markdown
## `_SUPPORT` events are parsed and not counted

Version 22 emits SPELL_DAMAGE_SUPPORT, SPELL_PERIODIC_DAMAGE_SUPPORT,
RANGE_DAMAGE_SUPPORT, SWING_DAMAGE_LANDED_SUPPORT, SPELL_HEAL_SUPPORT,
SPELL_PERIODIC_HEAL_SUPPORT and SPELL_ABSORBED_SUPPORT: 88,075 lines across the 89-file
corpus. Each restates a slice of damage or healing already reported on an ordinary line,
attributing it to the player whose buff caused it.

The layout decodes them and records the supporting player in `Event.Supporter`. The
summary accumulator skips any event with a non-empty `Supporter`, because adding them to
the existing damage would double-count every augmented hit. Showing a supporter's
contribution is a report-design decision and is deliberately not made here.
```

- [ ] **Step 9: Run the event and summary tests**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go test ./engine/event/ ./engine/summary/ 2>&1 | tail -20
```

Expected: `ok` for both packages.

- [ ] **Step 10: Run the whole suite and prove the v16 golden is untouched**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go vet ./... && go test ./... 2>&1 | tail -20
git -C /Users/jh/code/forever/.worktrees/retail-v22 diff --exit-code -- logs/engine/summary/testdata/v16.summary.json.golden logs/engine/event/testdata/v16.log && echo "v16 fixture and golden unchanged"
```

Expected: every package `ok`, then `v16 fixture and golden unchanged`.

- [ ] **Step 11: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
cat > /tmp/v22-msg-3.txt <<'MSG'
feat(logs): decode every combat-log v22 shape

The nineteen-field advanced block is read by its own offsets, so
position, power and item level land in the right fields. The trailing
single-target / area tag is recognised by value, which is the only rule
that fits RANGE_DAMAGE writing it and RANGE_MISSED not. _SUPPORT lines
name their supporting player and the summary skips them, because they
restate damage already counted elsewhere. COMBATANT_INFO's stat offsets
now come from the layout row instead of being hard-coded, which is what
lets v22 shift them all by one.

New events decoded: DAMAGE_SPLIT, DAMAGE_SHIELD_MISSED, the seven
_SUPPORT variants, SPELL_EMPOWER_START/END/INTERRUPT, ARENA_MATCH_START
and _END, STAGGER_CLEAR and _PREVENTED, WORLD_MARKER_PLACED and
_REMOVED.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
git add logs/engine/event/event.go logs/engine/event/decode.go logs/engine/event/special.go logs/engine/event/decode_v22_test.go logs/engine/summary/summary.go docs/ledger/2026-09-16-retail-v22.md
```

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22 && git commit -F /tmp/v22-msg-3.txt
```

---

## Task 4: The v22 summary golden and the probe-file fight check

**Files:**
- Modify: `logs/engine/summary/golden_test.go`
- Create: `logs/engine/summary/testdata/v22.summary.json.golden`
- Modify: `logs/cmd/forever-logs/main_test.go`

**Interfaces:**
- Consumes: `layout.RetailV22()`, the v22 excerpt, `Event.Supporter`.
- Produces: `logs/engine/summary/testdata/v22.summary.json.golden`, regenerated by
  `FOREVER_UPDATE_GOLDEN=1 go test ./engine/summary/`.

- [ ] **Step 1: Generalise the fixture helper and add the v22 golden test**

In `logs/engine/summary/golden_test.go`, replace the two constants and `fixtureSummary`'s
hard-coded path and layout with a parameterised helper, keeping the v16 test's behaviour
exactly:

```go
// goldenFile is the committed summary for the hand-written v16 fixture.
const goldenFile = "testdata/v16.summary.json.golden"

// goldenFileV22 is the committed summary for the v22 excerpt. A second
// dialect earns a second golden: the v16 one cannot catch a v22 field
// landing in the wrong column.
const goldenFileV22 = "testdata/v22.summary.json.golden"

// fixtureSummary runs the hand-written v16 fixture through the real
// pipeline and returns the summary of the first fight that closes.
func fixtureSummary(t *testing.T) (fight.Fight, Summary) {
	t.Helper()
	return excerptSummary(t, filepath.Join("..", "event", "testdata", "v16.log"),
		layout.RetailV16(), time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC))
}

// fixtureSummaryV22 does the same for the v22 excerpt. Version 22
// timestamps carry a year, so the base is never consulted; a deliberately
// wrong one is passed so that a regression that starts consulting it shows
// up as a wildly wrong date rather than as a plausible one.
func fixtureSummaryV22(t *testing.T) (fight.Fight, Summary) {
	t.Helper()
	return excerptSummary(t, filepath.Join("..", "event", "testdata", "v22.log"),
		layout.RetailV22(), time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC))
}
```

Rename the body of the current `fixtureSummary` to `excerptSummary` and give it the three
parameters, changing only these lines inside it:

```go
func excerptSummary(t *testing.T, path string, lay layout.Layout, base time.Time) (fight.Fight, Summary) {
	t.Helper()
	text, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	dec := event.NewDecoder(lay, base)
```

Everything from `reg := units.NewRegistry(...)` onwards is unchanged.

Add the v22 golden test:

```go
func TestTheV22ExcerptSummaryMatchesTheCommittedGolden(t *testing.T) {
	_, s := fixtureSummaryV22(t)
	got := append([]byte(jsonIndentOf(t, s)), '\n')

	if os.Getenv(regenEnv) != "" {
		if err := os.MkdirAll(filepath.Dir(goldenFileV22), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(goldenFileV22, got, 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("regenerated %s (%d bytes)", goldenFileV22, len(got))
		return
	}

	want, err := os.ReadFile(goldenFileV22)
	if err != nil {
		t.Fatalf("%v\nregenerate it with: %s=1 go test ./engine/summary/ -run %s", err, regenEnv, t.Name())
	}
	if string(got) != string(want) {
		t.Fatalf("the v22 excerpt's summary no longer matches %s.\n"+
			"If the change is intended, regenerate with:\n"+
			"    %s=1 go test ./engine/summary/ -run %s\n"+
			"and review the diff in the commit.\ngot  %d bytes\nwant %d bytes",
			goldenFileV22, regenEnv, t.Name(), len(got), len(want))
	}
}

// TestTheV22SummaryIsDatedFromTheLineNotTheBase proves the excerpt's own
// timestamps are used: the base handed to the decoder is in 2001 and the
// log is from 2026.
func TestTheV22SummaryIsDatedFromTheLineNotTheBase(t *testing.T) {
	f, _ := fixtureSummaryV22(t)
	if f.Start.Year() != 2026 {
		t.Fatalf("fight starts in %d, want 2026: the year came from the base, not the line", f.Start.Year())
	}
	if _, off := f.Start.Zone(); off != -5*3600 {
		t.Errorf("zone offset = %d seconds, want -18000", off)
	}
}

func TestTheV22GoldenSummaryIsStableAcrossRuns(t *testing.T) {
	_, first := fixtureSummaryV22(t)
	for i := 0; i < 10; i++ {
		_, again := fixtureSummaryV22(t)
		if jsonIndentOf(t, again) != jsonIndentOf(t, first) {
			t.Fatalf("run %d of the v22 excerpt produced a different summary", i)
		}
	}
}
```

- [ ] **Step 2: Run to verify the new test fails for the right reason**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go test ./engine/summary/ -run TestTheV22ExcerptSummaryMatchesTheCommittedGolden 2>&1 | head -10
```

Expected: FAIL with `open testdata/v22.summary.json.golden: no such file or directory`
followed by the regenerate instruction.

- [ ] **Step 3: Generate the golden and confirm the v16 one did not move**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
FOREVER_UPDATE_GOLDEN=1 go test ./engine/summary/ 2>&1 | tail -5
git -C /Users/jh/code/forever/.worktrees/retail-v22 diff --exit-code -- logs/engine/summary/testdata/v16.summary.json.golden && echo "v16 golden unchanged"
git -C /Users/jh/code/forever/.worktrees/retail-v22 status --short -- logs/engine/summary/testdata/
```

Expected: `ok`, then `v16 golden unchanged`, then exactly one untracked file,
`?? logs/engine/summary/testdata/v22.summary.json.golden`.

- [ ] **Step 4: Sanity-read the generated golden**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
python3 -c "
import json
s=json.load(open('engine/summary/testdata/v22.summary.json.golden'))
print('engine_version', s.get('engine_version'))
print('keys', sorted(s.keys())[:12])
print('players', len(s.get('players') or []))
"
```

Expected: `engine_version golden`, and a non-zero player count. A player count of 0 means
the excerpt's fight closed with no damage attributed, which would point at the
`Supporter` skip being too broad — check that `e.Supporter` is empty on ordinary
`SPELL_DAMAGE` events before accepting the golden.

- [ ] **Step 5: Add the probe-file fight and player check to the CLI test**

Append to `logs/cmd/forever-logs/main_test.go`:

```go
// TestFightsOnAV22ExcerptReportsPlayers runs the committed v22 excerpt
// through the fights command the way a user would, so the acceptance
// numbers in the plan are checked by something that runs in CI rather than
// only by the corpus sweep, which needs logs that are not in the repo.
func TestFightsOnAV22ExcerptReportsPlayers(t *testing.T) {
	var out strings.Builder
	if err := run([]string{"fights", filepath.Join("..", "..", "engine", "event", "testdata", "v22.log")}, &out, io.Discard); err != nil {
		t.Fatal(err)
	}
	got := out.String()
	if !strings.Contains(got, "retail-v22") {
		t.Errorf("the report does not name the retail-v22 layout:\n%s", got)
	}
	if strings.Contains(got, "inferred") {
		t.Errorf("the excerpt fell back to the inferred layout:\n%s", got)
	}
}

// TestConformanceOnTheV22ExcerptIsClean is the in-repo half of the
// acceptance: the committed excerpts must show the verified row, no parse
// errors and no unknown events.
func TestConformanceOnTheV22ExcerptIsClean(t *testing.T) {
	var out strings.Builder
	dir := filepath.Join("..", "..", "engine", "event", "testdata")
	if err := run([]string{"conformance", "-json", dir}, &out, io.Discard); err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(strings.TrimSpace(out.String()), "\n") {
		var row ConformanceRow
		if err := json.Unmarshal([]byte(line), &row); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(row.Path, "v22") {
			continue // the v16 fixture is covered by its own tests
		}
		if row.Layout != "retail-v22" || !row.Verified || row.Inferred {
			t.Errorf("%s: layout=%s verified=%v inferred=%v", row.Path, row.Layout, row.Verified, row.Inferred)
		}
		if row.ParseErrors != 0 {
			t.Errorf("%s: %d parse errors", row.Path, row.ParseErrors)
		}
		if len(row.UnknownEvents) != 0 {
			t.Errorf("%s: unknown events %v", row.Path, row.UnknownEvents)
		}
	}
}
```

Add `encoding/json`, `io`, `path/filepath` and `strings` to the file's imports if they
are not already there.

- [ ] **Step 6: Run the CLI and summary tests**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go test ./cmd/forever-logs/ ./engine/summary/ -run 'V22|Conformance' -v 2>&1 | grep -E '^(=== RUN|--- (PASS|FAIL)|ok|FAIL)' | head -20
```

Expected: `--- PASS` for `TestFightsOnAV22ExcerptReportsPlayers`,
`TestConformanceOnTheV22ExcerptIsClean`,
`TestTheV22ExcerptSummaryMatchesTheCommittedGolden`,
`TestTheV22SummaryIsDatedFromTheLineNotTheBase` and
`TestTheV22GoldenSummaryIsStableAcrossRuns`.

- [ ] **Step 7: Run the whole suite**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go test ./... 2>&1 | tail -20
```

Expected: every package `ok`.

- [ ] **Step 8: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
cat > /tmp/v22-msg-4.txt <<'MSG'
test(logs): a v22 summary golden and an in-repo conformance check

The golden test helper is now dialect-parameterised, so v16 and v22 each
pin their own decoded output; the v16 golden is byte-identical. The
decoder for the v22 excerpt is seeded with a 2001 base it must never
consult, so a regression that starts taking the year from the base
shows up as a date twenty-five years out.

The CLI test runs conformance over the committed testdata and requires
layout=retail-v22, verified, no parse errors and no unknown events, so
the acceptance shape is checked in CI and not only by the corpus sweep.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
git add logs/engine/summary/golden_test.go logs/engine/summary/testdata/v22.summary.json.golden logs/cmd/forever-logs/main_test.go
```

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22 && git commit -F /tmp/v22-msg-4.txt
```

---

## Task 5: The corpus conformance sweep and the acceptance ruling

**Files:**
- Create: `logs/scripts/v22-conformance.sh`
- Modify: `logs/engine/layout/retail.go` (only if the sweep finds a width outlier)
- Modify: `logs/engine/layout/measured_v22_test.go` (only if the sweep finds a width outlier)
- Append: `docs/ledger/2026-09-16-retail-v22.md`

**Interfaces:**
- Consumes: everything from Tasks 1–4.
- Produces: `logs/scripts/v22-conformance.sh`, a committed, re-runnable sweep; and the
  acceptance ruling in the ledger.

**Baseline to beat** (measured with the engine at main, `.../scratchpad/conformance-v22-before.jsonl`):
89 files, 26,030,980 lines, **6,725,218 parse errors**, 2,958 fights, every file
`layout=inferred verified=false`, 41 distinct unknown event names.

- [ ] **Step 1: Write the sweep script**

Create `logs/scripts/v22-conformance.sh`:

```bash
#!/usr/bin/env bash
# Sweep a directory of combat logs and fail unless every file parsed
# cleanly under a verified layout.
#
# Usage: logs/scripts/v22-conformance.sh <log-dir> [out.jsonl]
#
# The v22 corpus this was written against is 89 files and 7.3 GB and is
# not in the repository; it lives in the session scratchpad. The script is
# committed so the sweep is reproducible against any directory of logs.
set -euo pipefail

dir="${1:?usage: v22-conformance.sh <log-dir> [out.jsonl]}"
out="${2:-/tmp/v22-conformance.jsonl}"

cd "$(dirname "$0")/.."
go build -o /tmp/forever-logs ./cmd/forever-logs
/tmp/forever-logs conformance -json "$dir" > "$out"

python3 - "$out" <<'PY'
import json, sys, collections
rows = [json.loads(line) for line in open(sys.argv[1])]
unknown = collections.Counter()
errors = lines = fights = 0
bad_layout = []
for r in rows:
    for k, v in (r.get("unknown_events") or {}).items():
        unknown[k] += v
    errors += r["parse_errors"]
    lines += r["lines"]
    fights += r["fights"]
    if r["layout"] != "retail-v22" or not r["layout_verified"] or r["layout_inferred"]:
        bad_layout.append((r["path"], r["layout"], r["layout_verified"]))

print(f"files        {len(rows)}")
print(f"lines        {lines}")
print(f"fights       {fights}")
print(f"parse errors {errors}")
print(f"unknown      {len(unknown)} distinct")
for name, n in sorted(unknown.items()):
    print(f"  {name:<34} {n}")

fail = False
if bad_layout:
    fail = True
    print(f"\nFAIL: {len(bad_layout)} files did not select a verified retail-v22 row")
    for path, lay, ver in bad_layout[:10]:
        print(f"  {path} layout={lay} verified={ver}")
if errors:
    fail = True
    print(f"\nFAIL: {errors} parse errors")
if unknown:
    fail = True
    print(f"\nFAIL: {len(unknown)} distinct unknown events")
if fail:
    sys.exit(1)
print("\nPASS: every file parsed clean under retail-v22")
PY
```

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
chmod +x logs/scripts/v22-conformance.sh
```

- [ ] **Step 2: Run the sweep over the 89-file corpus**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
SCRATCH=/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad
logs/scripts/v22-conformance.sh "$SCRATCH/logs/v22-corpus" "$SCRATCH/conformance-v22-after.jsonl" 2>&1 | tail -40
```

Expected, and this is the acceptance:

```
files        89
lines        26030980
fights       <n>
parse errors 0
unknown      0 distinct

PASS: every file parsed clean under retail-v22
```

The sweep took a few minutes at main; expect the same order.

- [ ] **Step 3: If the sweep reports parse errors, fold the outliers back into the row**

The measurement in this plan's context section covers every `(event, width)` pair in the
corpus, so a non-zero count means one of three things. Diagnose before changing anything:

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
SCRATCH=/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad
# Which files, and how many each?
python3 -c "
import json
rows=[json.loads(l) for l in open('$SCRATCH/conformance-v22-after.jsonl')]
bad=[r for r in rows if r['parse_errors']]
print(len(bad),'files with errors')
for r in sorted(bad,key=lambda r:-r['parse_errors'])[:5]:
    print(r['parse_errors'], r['path'])
"
```

Then take the worst file and read the actual messages:

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
go run ./cmd/forever-logs parse --layout retail-v22 <worst-file> 2>&1 | grep -i 'fields, layout' | sort | uniq -c | sort -rn | head -20
```

Check `forever-logs parse -h` for the flag that prints per-line errors if `--layout` is
not it; the command's job here is only to surface the `has N fields, layout "retail-v22"
allows [...]` messages.

- If the message names a width one larger than an allowed one on a damage or miss
  suffix, it is an optional trailing field on a shape the corpus scan did not reach.
  Add the width to `v22MeasuredWidths` in `measured_v22_test.go` **and** make the row
  produce it, then add a line of that shape to `testdata/v22-shapes.log` so the decode
  test covers it.
- If the message names an event not in the context table at all, add a `Special` with
  the measured width and a `decodeSpecial` branch, exactly as Task 3 did for
  `STAGGER_CLEAR`.
- If a handful of lines in one file are truncated mid-write (a log the client was still
  appending to), that is not a layout bug. Count them, name the file, and record them in
  the ledger as a documented exception rather than widening the row to accept corruption.

Re-run Step 2 after each change until it prints `PASS`.

- [ ] **Step 4: If the sweep reports unknown events, rule on each one**

An unknown event is an event name the row has no entry for. For each name the sweep
prints, either add it to the row (a `Special` with its measured width, or a suffix) and
re-run, or rule it out in the ledger with its count and the reason. Do not leave an
unknown event unnamed in both places.

- [ ] **Step 5: Record the acceptance numbers in the ledger**

Append to `docs/ledger/2026-09-16-retail-v22.md`, filling in the numbers the sweep
actually printed:

```markdown
## Acceptance: the 89-file corpus sweep

Re-runnable with `logs/scripts/v22-conformance.sh <log-dir>`. The corpus is 89 files and
7.3 GB and is not in the repository.

| | before (engine at main) | after |
|---|---|---|
| files | 89 | 89 |
| lines | 26,030,980 | 26,030,980 |
| layout | `inferred`, unverified, on every file | `retail-v22`, verified, on every file |
| parse errors | 6,725,218 (26%) | <fill in: expected 0> |
| fights | 2,958 | <fill in> |
| distinct unknown events | 41 | <fill in: expected 0> |

<If anything is non-zero, one subsection per exception: the event or the file, the exact
count, why it is not a layout bug, and what would have to be true to fix it.>
```

- [ ] **Step 6: Run the whole suite one last time and check the v16 golden**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22/logs
gofmt -l . && go vet ./... && go test ./... 2>&1 | tail -20
git -C /Users/jh/code/forever/.worktrees/retail-v22 diff --exit-code main -- logs/engine/summary/testdata/v16.summary.json.golden logs/engine/event/testdata/v16.log && echo "v16 fixture and golden identical to main"
```

Expected: `gofmt -l` prints nothing, `go vet` is silent, every package `ok`, then
`v16 fixture and golden identical to main`.

Also confirm the constraint that this lane does not touch the engine version or the web:

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
git diff --stat main -- logs/engine/session/session.go web/ | tail -3
```

Expected: no output — neither path was modified.

- [ ] **Step 7: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22
cat > /tmp/v22-msg-5.txt <<'MSG'
test(logs): the v22 corpus conformance sweep and its acceptance ruling

A committed script builds the engine, runs conformance over a directory
of logs and fails unless every file selected a verified row with no
parse errors and no unknown events. Run over the 89-file, 26-million-line
v22 corpus it replaces a baseline of 6,725,218 parse errors on an
unverified inferred layout.

The ledger records the before and after numbers and any exception with
its count.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
git add logs/scripts/v22-conformance.sh docs/ledger/2026-09-16-retail-v22.md
git add -u logs/engine/layout/ logs/engine/event/testdata/ 2>/dev/null || true
```

```bash
cd /Users/jh/code/forever/.worktrees/retail-v22 && git commit -F /tmp/v22-msg-5.txt
```

---

## Self-Review

**1. Spec coverage.**

| Spec requirement | Task |
|---|---|
| A second verified layout `retail-v22` beside `retail-v16`, `Version: 22`, `StampYear`, `StampZone`, `Advanced: 19` | Task 2, Step 5 |
| `_DAMAGE` params for SPELL prefixes vs SWING ("or one rule if the extra field proves to be a prefix field") | Task 2: the measurement showed it is a suffix field, and the one rule is `Suffix.Tag` read by value |
| `_MISSED` params 3 | Task 2: refined by measurement to `Params: 2` plus `Tag` plus `MissAmount`, which is what the 15/16/18 widths require; the spec's "3" was the probe-file reading of the 15-wide form |
| The same specials as v16 | Task 2, Step 5 |
| `DAMAGE_SPLIT` and the `_SUPPORT` variants read as damage or heal lines | Task 2 (`_SPLIT`, `_SUPPORT`, `_HEAL_SUPPORT`), Task 3 (decode) |
| The summary decides whether to fold support damage | Task 3, Step 8: it does not fold them, with the reason in the ledger |
| `SPELL_EMPOWER_*` as cast-like lines with a stage field | Task 2 (`_EMPOWER_*` suffixes), Task 3 (`e.Stacks`) |
| `ARENA_MATCH_START/END` as specials, not used as fight boundaries | Task 2 (specials), Task 3 (decode); no fight-splitter change |
| Selection by header version; inferred fallback untouched | Task 2, Steps 3 and 5; `infer.go` is not in any task's file list |
| Engine version unchanged, v16 golden and fixture unchanged | Global Constraints; asserted in Task 2 Step 7, Task 3 Step 10, Task 4 Step 3, Task 5 Step 6 |
| Committed `v22.log` excerpt covering every event type | Task 1, Steps 1–2 (split into two files, because no contiguous excerpt contains all 65 event names) |
| `v22.summary.json.golden` | Task 4 |
| `internal/sample` stays on v16 | not in any task's file list |
| Acceptance: corpus sweep, `layout=retail-v22 verified=true`, 0 parse errors, unknowns named | Task 5 |
| Layout width test gains a v22 twin | Task 2, Step 1 |
| `go test ./...` green | every task's last test step |
| Probe files parse to the same fights with zero errors, player counts match COMBATANT_INFO | Task 4 Step 5 and Task 5 Step 2 — the corpus sweep supersedes the two-probe check, and its fight total is recorded in the ledger |

Two spec numbers were superseded by the larger measurement and are called out above:
`_MISSED` params, and `ENVIRONMENTAL_DAMAGE` (the spec said "width to verify against
v16's 37"; it is 39). The spec's "keep the v16 combatant parser; verify on the corpus"
resolved the other way: verification showed the indices must shift, which is Task 2
Step 5 and Task 3 Step 7.

**2. Placeholder scan.** The only intentional fill-ins are in Task 5 Step 5's ledger
table, which records numbers that do not exist until the sweep runs, and the
`<worst-file>` argument in Task 5 Step 3, which is output from the preceding command.
Both are marked. No step says "add error handling", "similar to Task N", or "write tests
for the above".

**3. Type consistency.** Checked across tasks: `RetailV22()` (Task 2) is called by Tasks
3, 4 and 5. `Layout.Widths(prefix, suffix) []int` is defined in Task 2 Step 3 and used in
Task 2 Step 1's test and Task 3 Step 5's decoder. `Suffix.Tag` / `.MissAmount` /
`.AuraExtra` are declared in Task 2 Step 3 and consumed in Task 2 Step 5's row and Task 3
Steps 5–6. `Combatant.StatIndex` is declared in Task 2 Step 3, populated for v16 in Step
4 and for v22 in Step 5, and read in Task 3 Step 7. `Event.Scope` and `Event.Supporter`
are declared in Task 3 Step 3 and used in Task 3 Steps 5–6 and 8 and Task 4. `Advanced.Versatility`
and `Advanced.Unknown8` are declared in Task 3 Step 3 and written in Step 4.
`AuraCarriesAmount` is exported from the layout package in Task 2 Step 3 because the
decoder's aura branch and `Widths` both need the same predicate. `v22MeasuredWidths` is
defined in Task 1 and read in Task 2 Step 1 — both live in package `layout`, so the
test-only identifier is visible. `excerptSummary` (Task 4 Step 1) is the renamed body of
the existing `fixtureSummary`, and both `fixtureSummary` and `fixtureSummaryV22` call it,
so the existing v16 tests keep compiling.

One thing an implementer must check rather than assume: `readAbsorbed`'s exact signature,
flagged in Task 3 Step 7, and the flag `forever-logs parse` uses to print per-line
errors, flagged in Task 5 Step 3.
