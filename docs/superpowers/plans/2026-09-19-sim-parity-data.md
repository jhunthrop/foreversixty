# Simulator parity: the data lane — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Give the simulator's Droptimizer and Top Gear pages the data they need — a per-build `loot.json`, `enchants.json` and `suffixes.json`, a curated loot overlay, `reference_stat` on every spec, and a build-validation guard that fails the day an item ships with a gem socket.

**Architecture:** Nothing in a DB2 export says which boss drops an item, so the engine fork's own `assets/database/db.json` (AtlasLoot + Wowhead sources for 7,553 Era items, 173 enchants, 1,168 random suffixes, 274 named NPCs, 18 factions) is the only source for loot, enchants and suffixes. A new offline-ish `loot` command reads that database from an engine checkout at the pinned sha, joins it to the build's own `zones.json`, `items.json` and two raw DB2 tables (`Map.csv` for instance types, `ItemSparse.csv` for PvP ranks), applies a curated overlay from `data/curated/loot/*.json`, and writes three files plus a `suffixes` column on `items.json`. The generated files are committed per build, exactly like `simdb.bin` and `gametables/`, and are gated in CI by a conformance test with golden counts rather than by regeneration — CI has no `raw/`.

**Tech Stack:** Python 3.12, uv, pydantic v2, pytest + pytest-cov (80% floor), ruff (line-length 100, `E,F,I,B,UP`). No new dependencies.

**Spec:**
- Design: `docs/superpowers/specs/2026-09-19-simulator-parity-design.md` (sections 4.4, 6.1, 6.2, 12, 13)
- Contract: `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` — section 6 is this lane's binding interface, **as corrected by section 10.4**; section 10.3's last three bullets and section 10.1's A7 are also this lane's. Where 10 contradicts 6, 10 wins.

## Global Constraints

- Python `>=3.12` (`data/pyproject.toml`), `.python-version` is `3.12`; everything runs under `uv` from the `data/` directory: `uv run python -m pipeline …`, `uv run pytest`, `uv run ruff check .`.
- Ruff: `line-length = 100`, `target-version = "py312"`, `select = ["E", "F", "I", "B", "UP"]`. Every new file must pass `uv run ruff check .`.
- pytest runs with `--cov=pipeline --cov-report=term-missing --cov-fail-under=80`. New modules need tests in the same run or the suite fails.
- Generated files are committed per build under `data/builds/<build>/`. The active build is `1.60.1.69893` (`web/src/data/active-build.json`).
- Every curated file carries `sources` (a list of `{label, url, kind}` with `kind` in `blizzard | datamined | community | site`) and `notes`. An unsourced claim stops the pipeline.
- Nothing fabricated. An item with no source in either database is simply absent from `loot.json`. A boss the fork database does not name is emitted with an empty `name`, never an invented one.
- `opens` values are phase names from `api/internal/phase/phase.go` — `pre-beta`, `beta`, `launch`, `raids-1` — plus the sentinel `later` (contract 10.4) for a source whose date is unknown, which the page shows as unreleased without a date. `later` is not a boundary and never appears in `phases.json`.
- `loot.json` lists only items the build has (contract 10.4). An id the fork database names that this client's `ItemSparse` does not have is dropped and counted; the first overlay records the resulting gap per raid.
- The stat vocabulary for `reference_stat` is the fork's `proto.Stat` enum names in lower snake case (contract 10.1 A7): `spell_haste` and `melee_haste`, never a bare `haste`.
- JSON serialization is `json.dumps(payload, indent=2, ensure_ascii=False) + "\n"` via `pipeline.normalize._write` and its public wrappers. Never hand-roll it.
- One commit per task: `feat(data): …` for code and data, `test(data): …` for a test-only task. Every commit message ends with the trailer:
  ```
  Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
  ```
- Never `git stash` — this repository shares a stash stack across worktrees. Use a temporary WIP commit instead.
- All paths below are relative to the repository root unless they start with `data/`, in which case they are still repository-relative; commands are run from `data/`.

---

## File map

| File | Responsibility |
| --- | --- |
| `data/pipeline/normalize/sockets.py` | new — the gem-socket guard (contract 6.4) |
| `data/pipeline/forkdb.py` | new — read the engine fork's `assets/database/db.json`, and the enum tables that decode it |
| `data/pipeline/loot/__init__.py` | new — the `loot` command's orchestration: read, build, overlay, write, refresh the manifest |
| `data/pipeline/loot/sources.py` | new — `loot.json` from the fork database's item sources |
| `data/pipeline/loot/gear.py` | new — `enchants.json`, `suffixes.json`, and `items.json`'s two new columns |
| `data/pipeline/loot/overlay.py` | new — load and apply `data/curated/loot/*.json` |
| `data/pipeline/loot/buffs.py` | new — `simbuffs.json`: a name and icon per IDS.md id (contract 10.4) |
| `data/pipeline/phases.py` | new — `curated/phases.json` → `web/src/data/phases.json` (contract 10.4) |
| `data/pipeline/models.py` | modified — `SpecRecord.reference_stat`, `Item.suffixes` and `Item.faction_restriction`, and the loot/enchant/suffix/buff/phase records |
| `data/pipeline/simdb/statmap.py` | modified — `stat_keys`, the inverse of `stat_array`; `STAT_IDS`, the A7 vocabulary |
| `data/pipeline/simdb/items.py` | modified — `SimItem`'s `unique`, `required_level`, `faction_restriction`, `random_suffix_options` (contract 10.3) |
| `data/pipeline/specs.py` | modified — validate and render `reference_stat` |
| `data/pipeline/curated.py` | modified — `_sources` becomes the public `parse_sources` |
| `data/pipeline/normalize/__init__.py` | modified — call the socket guard; add `write_document` |
| `data/pipeline/__main__.py` | modified — the `loot` and `phases` subcommands |
| `data/proto/*`, `data/pipeline/simproto/*_pb2.py` | re-vendored and regenerated at the new engine pin (contract 10.3) |
| `data/curated/specs.json` | modified — `reference_stat` per spec |
| `data/curated/loot/forever-raid-phases.json` | new — the first overlay |
| `data/curated/simbuffs.json` | new — the fourteen IDS.md ids no name join resolves |
| `data/curated/phases.json` | new — the phase boundaries, the source of truth for Go and the web |
| `web/src/data/phases.json` | new (generated, committed) |
| `Makefile` | modified — `loot` and `loot-check` |
| `.github/workflows/data.yml` | modified — the fork clone, the `loot` step, and `phases --check` |

---

### Task 1: The gem-socket guard

Design section 4.4 says gems, sockets and socket bonuses are deliberately absent because the 1.60 client's 19,171 items set no socket column. Contract 6.4 turns that into a build check: the pipeline fails when any `ItemSparse` row has a non-zero `SocketType_*`.

**Files:**
- Create: `data/pipeline/normalize/sockets.py`
- Modify: `data/pipeline/normalize/__init__.py` (inside `normalize_build`, before the flat-entity writes)
- Test: `data/tests/test_normalize_sockets.py`

**Interfaces:**
- Consumes: `pipeline.csvio.populated(row, key) -> str | None`
- Produces: `pipeline.normalize.sockets.SOCKET_COLUMNS: tuple[str, ...]`, `socketed_item_ids(sparse_rows: list[dict[str, str]]) -> list[int]`, `check_no_sockets(sparse_rows: list[dict[str, str]], build: str) -> None`, `class SocketError(SystemExit)`

- [ ] **Step 1: Write the failing test**

Create `data/tests/test_normalize_sockets.py`:

```python
import pytest

from pipeline.normalize.sockets import SocketError, check_no_sockets, socketed_item_ids


def row(item_id: int, *sockets: int) -> dict[str, str]:
    padded = list(sockets) + [0] * (3 - len(sockets))
    return {"ID": str(item_id)} | {
        f"SocketType_{index}": str(value) for index, value in enumerate(padded)
    }


def test_a_client_with_no_socketed_item_passes():
    rows = [row(1), row(2), row(3)]
    assert socketed_item_ids(rows) == []
    check_no_sockets(rows, "1.60.1.69893")


def test_a_row_missing_the_socket_columns_entirely_is_not_socketed():
    """Classic Era's ItemSparse export carries the columns; a trimmed
    fixture or a future client that drops them must read as 'no sockets',
    not raise inside int()."""
    assert socketed_item_ids([{"ID": "7"}]) == []


def test_an_empty_socket_column_is_not_socketed():
    assert socketed_item_ids([{"ID": "7", "SocketType_0": ""}]) == []


def test_socketed_items_are_reported_in_id_order():
    assert socketed_item_ids([row(9, 0, 2), row(4), row(7, 1)]) == [7, 9]


def test_a_socketed_item_stops_the_build_and_names_the_ids():
    with pytest.raises(SocketError) as error:
        check_no_sockets([row(4), row(7, 1), row(9, 0, 0, 3)], "1.61.0.1")
    message = str(error.value)
    assert "1.61.0.1" in message
    assert "7, 9" in message
    assert "2 item(s)" in message


def test_a_long_list_is_truncated_but_counted():
    rows = [row(item_id, 1) for item_id in range(100, 130)]
    with pytest.raises(SocketError) as error:
        check_no_sockets(rows, "1.61.0.1")
    message = str(error.value)
    assert "30 item(s)" in message
    assert "and 10 more" in message
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd data && uv run pytest tests/test_normalize_sockets.py -q --no-cov`
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.normalize.sockets'`

- [ ] **Step 3: Write the implementation**

Create `data/pipeline/normalize/sockets.py`:

```python
"""The gem signal.

The parity design's section 4.4 states outright that gems, sockets, socket
bonuses and the whole gem economy do not exist here: the 1.60 client's item
table carries vanilla's socket columns and **none of its 19,171 rows sets
one**. Every consumer downstream -- the planner's item cards, `simdb.bin`,
the Top Gear expander -- is built on that measurement.

The day a build arrives with a socketed item, that is a design decision to
make, not a column to ignore. So the pipeline stops here and names the
items, rather than emitting a build whose sockets nothing reads.
"""

from __future__ import annotations

from pipeline.csvio import populated

#: The ItemSparse columns that would hold a gem socket's colour. A client
#: that does not export them at all reads as "no sockets", which is what
#: `populated` returning None means.
SOCKET_COLUMNS = ("SocketType_0", "SocketType_1", "SocketType_2")

#: How many ids the error message spells out before it just counts.
SHOWN = 20


class SocketError(SystemExit):
    """An item has a gem socket, which nothing in this repository models."""


def socketed_item_ids(sparse_rows: list[dict[str, str]]) -> list[int]:
    """Every ItemSparse id with a non-zero socket colour, in id order."""
    found: list[int] = []
    for row in sparse_rows:
        for column in SOCKET_COLUMNS:
            value = populated(row, column)
            if value is not None and int(value) != 0:
                found.append(int(row["ID"]))
                break
    return sorted(found)


def check_no_sockets(sparse_rows: list[dict[str, str]], build: str) -> None:
    found = socketed_item_ids(sparse_rows)
    if not found:
        return
    shown = ", ".join(str(item_id) for item_id in found[:SHOWN])
    more = f" and {len(found) - SHOWN} more" if len(found) > SHOWN else ""
    raise SocketError(
        f"build {build} has {len(found)} item(s) with a gem socket: {shown}{more}. "
        "Nothing here models gems -- see the parity design's section 4.4, which "
        "records that no item on 1.60.1.69893 had one. Design gems before "
        "emitting this build."
    )
```

- [ ] **Step 4: Wire it into the build**

In `data/pipeline/normalize/__init__.py`, add the import to the block of local imports inside `normalize_build` (alphabetical among the `pipeline.normalize.*` imports, after `from pipeline.normalize.items import normalize_items`):

```python
    from pipeline.normalize.sockets import check_no_sockets
```

Then, immediately after the line `class_rows = t("ChrClasses")` and before the `write_json(normalize_zones(...))` line, insert:

```python
    # Contract 6.4: a socketed item means the gem decision in the parity
    # design's section 4.4 no longer holds. Checked before anything is
    # written, so a build that trips it leaves no half-emitted directory.
    check_no_sockets(t("ItemSparse"), build)
```

- [ ] **Step 5: Run the tests**

Run: `cd data && uv run pytest tests/test_normalize_sockets.py tests/test_normalize_build.py -q --no-cov`
Expected: PASS (the fixtures set no socket column)

- [ ] **Step 6: Lint**

Run: `cd data && uv run ruff check .`
Expected: `All checks passed!`

- [ ] **Step 7: Commit**

```bash
git add data/pipeline/normalize/sockets.py data/pipeline/normalize/__init__.py data/tests/test_normalize_sockets.py
git commit -m "$(cat <<'EOF'
feat(data): fail the build when an item has a gem socket

Contract 6.4. The parity design's section 4.4 records that none of the
1.60 client's 19,171 items sets a socket column, and every consumer is
built on that. A build that breaks it now stops the pipeline and names
the item ids instead of shipping sockets nothing reads.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 2: `reference_stat` on every spec

Contract 1.4: `WeightsSpec.Reference` is "the stat normalised to 1.0; the spec's default from `data/curated/specs.json`". Contract 8: `GET /v1/specs` rows carry `reference_stat`. `curated/specs.json` is the single source and both `sim/specs/specs.go` and `web/src/lib/sim/specs.ts` are generated from it and committed, so all three change together.

**Contract 10.1 A7 fixes the vocabulary**: the stat ids are the fork's `proto.Stat` enum names in lower snake case — `StatSpellPower` → `spell_power`, `StatMeleeHaste` → `melee_haste`, `StatSpellHaste` → `spell_haste`, `StatMP5` → `mp5` — and never a bare `haste`. The sim module publishes them as a generated **Stats** section in IDS.md; this lane derives the same 41 names from the vendored `pb.Stat` enum so the two cannot drift. `pipeline/simdb/statmap.py`'s `PROTO_STAT_ALIASES` is the *planner's* vocabulary and is **not** what `reference_stat` is validated against.

The table, exactly: attack power for every melee spec and for all three hunter specs, spell power for every caster spec. Both names are already `proto.Stat` names in snake case, so no value in the table changes under A7 — only what validates it.

| spec | reference_stat | | spec | reference_stat |
| --- | --- | --- | --- | --- |
| `druid-balance` | `spell_power` | | `rogue-assassination` | `attack_power` |
| `druid-feral` | `attack_power` | | `rogue-combat` | `attack_power` |
| `druid-restoration` | `spell_power` | | `rogue-subtlety` | `attack_power` |
| `hunter-beast-mastery` | `attack_power` | | `shaman-elemental` | `spell_power` |
| `hunter-marksmanship` | `attack_power` | | `shaman-enhancement` | `attack_power` |
| `hunter-survival` | `attack_power` | | `shaman-restoration` | `spell_power` |
| `mage-arcane` | `spell_power` | | `warlock-affliction` | `spell_power` |
| `mage-fire` | `spell_power` | | `warlock-demonology` | `spell_power` |
| `mage-frost` | `spell_power` | | `warlock-destruction` | `spell_power` |
| `paladin-holy` | `spell_power` | | `warrior-arms` | `attack_power` |
| `paladin-protection` | `attack_power` | | `warrior-fury` | `attack_power` |
| `paladin-retribution` | `attack_power` | | `warrior-protection` | `attack_power` |
| `priest-discipline` | `spell_power` | | | |
| `priest-holy` | `spell_power` | | | |
| `priest-shadow` | `spell_power` | | | |

Fifteen `spell_power`, twelve `attack_power`.

**Files:**
- Modify: `data/pipeline/models.py` (`SpecRecord`)
- Modify: `data/pipeline/simdb/statmap.py` (`stat_id`, `STAT_IDS`)
- Modify: `data/pipeline/specs.py` (`load_specs`, `render_go`, `render_ts`)
- Modify: `data/curated/specs.json` (all 27 rows)
- Modify: `sim/specs/specs.go`, `web/src/lib/sim/specs.ts` (regenerated, never hand-edited)
- Test: `data/tests/test_specs.py`

**Interfaces:**
- Consumes: `pipeline.simproto.pb.Stat`
- Produces: `SpecRecord.reference_stat: str`; `pipeline.simdb.statmap.stat_id(enum_name: str) -> str` and `STAT_IDS: frozenset[str]` (the 41 snake-case `proto.Stat` names); the Go field `ReferenceStat string \`json:"reference_stat"\`` on `specs.Spec`; the TypeScript field `reference_stat: string` on `Spec`.

- [ ] **Step 1: Write the failing tests**

Append to `data/tests/test_specs.py`, and add `from pipeline.simdb.statmap import STAT_IDS` to its imports:

```python
#: The parity contract's section 1.4 default per spec: attack power for
#: melee and hunters, spell power for casters. Transcribed here from the
#: plan's table rather than imported, so the curated file is checked
#: against the decision and not against itself.
REFERENCE_BY_SPEC = {
    "druid-balance": "spell_power",
    "druid-feral": "attack_power",
    "druid-restoration": "spell_power",
    "hunter-beast-mastery": "attack_power",
    "hunter-marksmanship": "attack_power",
    "hunter-survival": "attack_power",
    "mage-arcane": "spell_power",
    "mage-fire": "spell_power",
    "mage-frost": "spell_power",
    "paladin-holy": "spell_power",
    "paladin-protection": "attack_power",
    "paladin-retribution": "attack_power",
    "priest-discipline": "spell_power",
    "priest-holy": "spell_power",
    "priest-shadow": "spell_power",
    "rogue-assassination": "attack_power",
    "rogue-combat": "attack_power",
    "rogue-subtlety": "attack_power",
    "shaman-elemental": "spell_power",
    "shaman-enhancement": "attack_power",
    "shaman-restoration": "spell_power",
    "warlock-affliction": "spell_power",
    "warlock-demonology": "spell_power",
    "warlock-destruction": "spell_power",
    "warrior-arms": "attack_power",
    "warrior-fury": "attack_power",
    "warrior-protection": "attack_power",
}


def test_every_spec_names_the_reference_stat_the_contract_gives_it():
    assert {record.spec: record.reference_stat for record in specs()} == REFERENCE_BY_SPEC


def test_the_stat_vocabulary_is_the_proto_enum_in_snake_case():
    """Contract 10.1 A7. Spot values, plus the shape of the whole set, so a
    renamed or added engine stat shows up here and not in a sim that
    silently normalises to nothing."""
    assert "spell_power" in STAT_IDS
    assert "attack_power" in STAT_IDS
    assert "spell_haste" in STAT_IDS and "melee_haste" in STAT_IDS
    assert "haste" not in STAT_IDS
    assert "mp5" in STAT_IDS
    assert len(STAT_IDS) == 41
    assert all(name == name.lower() for name in STAT_IDS)


def test_a_reference_stat_the_engine_has_no_stat_for_is_refused(tmp_path):
    rows = [
        {
            "spec": "warrior-arms",
            "class_slug": "warrior",
            "spec_slug": "arms",
            "name": "Arms",
            "role": "dps",
            "tree_index": 0,
            "reference_stat": "swagger",
        }
    ]
    (tmp_path / "specs.json").write_text(json.dumps(rows), encoding="utf-8")
    with pytest.raises(SpecError, match="reference_stat"):
        load_specs(tmp_path)


def test_both_generated_files_carry_the_reference_stat():
    records = specs()
    go, ts = render_go(records), render_ts(records)
    assert 'ReferenceStat string `json:"reference_stat"`' in go
    assert 'ReferenceStat: "attack_power"' in go
    assert "reference_stat: string;" in ts
    assert "reference_stat: 'spell_power'," in ts
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd data && uv run pytest tests/test_specs.py -q --no-cov`
Expected: FAIL — `ImportError: cannot import name 'STAT_IDS'`, and after that is fixed, `pydantic_core.ValidationError: SpecRecord … reference_stat Field required`

- [ ] **Step 3: Add the model field**

In `data/pipeline/models.py`, replace the `SpecRecord` class with:

```python
class SpecRecord(BaseModel):
    spec: str
    class_slug: str
    spec_slug: str
    name: str
    role: str
    tree_index: int
    #: The stat `StatWeights` normalises to 1.0 for this spec (parity
    #: contract 1.4). Appended last, like `TalentEntry.spell_id`: the
    #: emitted key order is the generated files' only compatibility
    #: surface, so new fields go on the end and existing ones never move.
    reference_stat: str
```

- [ ] **Step 4: Validate and render it**

First, the vocabulary itself. In `data/pipeline/simdb/statmap.py`, add after `PROTO_STAT_ALIASES`:

```python
def stat_id(enum_name: str) -> str:
    """A `proto.Stat` enum name as the request vocabulary's id.

    Contract 10.1 A7: the stat ids `WeightsSpec.Reference` and
    `specs.json`'s `reference_stat` use are the fork's own enum names in
    lower snake case -- `StatSpellPower` is `spell_power`, `StatMeleeHaste`
    is `melee_haste`, and there is deliberately no bare `haste`. Derived
    from the enum rather than listed, so an engine that renames a stat
    renames the id with it.

    This is NOT `PROTO_STAT_ALIASES`. That is the planner's vocabulary,
    which merges several enum values onto one key (`hit` covers both
    `StatHit` and the lineage's `StatMeleeHit`) and exists to read item
    columns; A7's is one id per enum value.
    """
    core = enum_name.removeprefix("Stat")
    return re.sub(r"(?<=[a-z0-9])(?=[A-Z])", "_", core).lower()


#: Every stat id A7's vocabulary has, from the vendored enum. 41 on the
#: pinned engine, all distinct.
STAT_IDS = frozenset(stat_id(name) for name in pb.Stat.keys())
```

and add `import re` to the module's imports.

Then, in `data/pipeline/specs.py`, add the import after `from pipeline.models import SpecRecord`:

```python
from pipeline.simdb.statmap import STAT_IDS
```

In `load_specs`, after the `record.role not in ROLES` check, insert:

```python
        if record.reference_stat not in STAT_IDS:
            raise SpecError(
                f"spec {record.spec} has reference_stat {record.reference_stat!r}; "
                f"contract 10.1 A7's vocabulary is the engine's Stat enum in snake "
                f"case -- use one of {sorted(STAT_IDS)}"
            )
```

In `render_go`, replace the `rows = …` expression with:

```python
    rows = "\n".join(
        f'\t{{Spec: "{s.spec}", ClassSlug: "{s.class_slug}", SpecSlug: "{s.spec_slug}", '
        f'Name: "{s.name}", Role: "{s.role}", TreeIndex: {s.tree_index}, '
        f'ReferenceStat: "{s.reference_stat}"}},'
        for s in specs
    )
```

and add the struct field after `TreeIndex int    \`json:"tree_index"\``:

```
\tReferenceStat string `json:"reference_stat"`
```

(in the f-string the literal line is `\tReferenceStat string \`json:"reference_stat"\``; the whole template is inside a `"""…"""` so the backticks need no escaping, and `{{`/`}}` doubling already applies to the braces around it.)

In `render_ts`, add `reference_stat: s.reference_stat` to the row template and the field to the interface:

```python
        f"    tree_index: {s.tree_index},\n"
        f"    reference_stat: '{s.reference_stat}',\n"
```

```
  tree_index: number;
  reference_stat: string;
```

- [ ] **Step 5: Add the field to the curated list**

Add `"reference_stat": "<value>"` to each of the 27 objects in `data/curated/specs.json`, after `"tree_index"`, using the table at the top of this task. For example the first and last rows become:

```json
  { "spec": "druid-balance", "class_slug": "druid", "spec_slug": "balance", "name": "Balance", "role": "dps", "tree_index": 0, "reference_stat": "spell_power" },
```
```json
  { "spec": "warrior-protection", "class_slug": "warrior", "spec_slug": "protection", "name": "Protection", "role": "tank", "tree_index": 2, "reference_stat": "attack_power" }
```

- [ ] **Step 6: Regenerate the two generated files**

Run: `cd data && uv run python -m pipeline specs`
Expected: prints `../sim/specs/specs.go` and `../web/src/lib/sim/specs.ts`

Then: `cd data && uv run python -m pipeline specs --check`
Expected: exit 0, no output

- [ ] **Step 7: Run the tests**

Run: `cd data && uv run pytest tests/test_specs.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 8: Check the Go side still compiles**

Run: `cd sim && gofmt -l specs/ && go build ./specs/`
Expected: no output from either

- [ ] **Step 9: Commit**

```bash
git add data/pipeline/models.py data/pipeline/specs.py data/curated/specs.json data/tests/test_specs.py sim/specs/specs.go web/src/lib/sim/specs.ts
git commit -m "$(cat <<'EOF'
feat(data): reference_stat per spec

Contract 1.4 and 8: the stat a spec's stat weights normalise to 1.0.
Attack power for the twelve melee and hunter specs, spell power for the
fifteen casters. curated/specs.json stays the single source; the Go and
TypeScript lists are regenerated from it, and a reference stat the
engine's Stat enum has no entry for is refused at load.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 3: The engine fork's item database, and the inverse of `stat_array`

Everything in contract section 6 comes out of one file that is **not in this repository**: `assets/database/db.json` in the wowsims-forever checkout. Measured on the pinned fork on 2026-09-19: `items` 7,553 (4,481 of them with `sources`), `randomSuffixes` 1,168, `enchants` 173 (150 distinct effect ids), `zones` 27, `npcs` 274, `factions` 18, `itemIcons` 103, `spellIcons` 4,038, `encounters` 2. This task is the reader and the enum tables that decode it, plus the one thing `statmap` is missing: reading a `repeated double stats` array back into planner stat keys.

**Files:**
- Create: `data/pipeline/forkdb.py`
- Modify: `data/pipeline/simdb/statmap.py` (add `stat_keys`)
- Create: `data/tests/fixtures/loot/db.json` (shared by tasks 4, 5 and 6)
- Test: `data/tests/test_forkdb.py`, `data/tests/test_simdb_statmap.py` (extend)

**Interfaces:**
- Consumes: nothing from earlier tasks.
- Produces:
  - `pipeline.forkdb.ForkDbError(SystemExit)`
  - `pipeline.forkdb.DB_RELATIVE: Path` = `Path("assets/database/db.json")`
  - `@dataclass(frozen=True) pipeline.forkdb.ForkDatabase` with fields `items: tuple[dict, ...]`, `enchants: tuple[dict, ...]`, `random_suffixes: tuple[dict, ...]`, `zones: dict[int, str]`, `npcs: dict[int, str]`, `factions: dict[int, str]`, `item_icons: dict[int, str]`, `spell_icons: dict[int, str]`
  - `pipeline.forkdb.load_fork_database(engine_dir: Path) -> ForkDatabase`
  - `pipeline.forkdb.PROFESSIONS`, `REP_LEVELS`, `ITEM_TYPE_SLOTS`, `ENCHANT_TYPES`, `CLASS_SLUGS`
  - `pipeline.forkdb.decode(table: dict[int, T], value: int, what: str) -> T`
  - `pipeline.simdb.statmap.stat_keys(array: Sequence[float]) -> dict[str, float]`

- [ ] **Step 1: Write the fixture database**

Create `data/tests/fixtures/loot/db.json`. It is deliberately tiny and exercises every branch the three builders have: a named raid boss, an unnamed raid boss, raid trash, a dungeon boss reached through the *other* zone join, a named world boss, an unnamed open-world mob (dropped), crafted, reputation, quest, a vendor-only item (dropped), suffix options, and two enchants that share an effect id.

```json
{
  "items": [
    { "id": 100, "name": "Raid Boss Drop", "icon": "inv_a", "type": 5, "ilvl": 76, "phase": 1, "quality": 4,
      "factionRestriction": 1,
      "sources": [{ "drop": { "difficulty": 1, "npcId": 900, "zoneId": 2717 } }] },
    { "id": 101, "name": "Raid Trash Drop", "icon": "inv_b", "type": 6, "ilvl": 66, "phase": 1, "quality": 3,
      "sources": [{ "drop": { "difficulty": 1, "zoneId": 2717, "otherName": "Trash" } }] },
    { "id": 102, "name": "Nameless Boss Drop", "icon": "inv_c", "type": 7, "ilvl": 76, "phase": 1, "quality": 4,
      "sources": [{ "drop": { "difficulty": 1, "npcId": 901, "zoneId": 2717 } }] },
    { "id": 103, "name": "Dungeon Boss Drop", "icon": "inv_d", "type": 8, "ilvl": 50, "phase": 1, "quality": 3,
      "sources": [{ "drop": { "difficulty": 1, "npcId": 902, "zoneId": 1581 } }] },
    { "id": 104, "name": "World Boss Drop", "icon": "inv_e", "type": 2, "ilvl": 74, "phase": 1, "quality": 4,
      "sources": [{ "drop": { "difficulty": 1, "npcId": 903, "zoneId": 16 } }] },
    { "id": 105, "name": "Roadside Mob Drop", "icon": "inv_f", "type": 9, "ilvl": 30, "phase": 1, "quality": 2,
      "sources": [{ "drop": { "difficulty": 1, "npcId": 904, "zoneId": 16 } }] },
    { "id": 106, "name": "Crafted Plate", "icon": "inv_g", "type": 5, "ilvl": 60, "phase": 1, "quality": 3,
      "sources": [{ "crafted": { "profession": 2, "spellId": 9999 } }] },
    { "id": 107, "name": "Exalted Cloak", "icon": "inv_h", "type": 4, "ilvl": 68, "phase": 1, "quality": 4,
      "sources": [{ "rep": { "repFactionId": 529, "repLevel": 8, "playerFaction": 1 } }] },
    { "id": 108, "name": "Quest Ring", "icon": "inv_i", "type": 11, "ilvl": 40, "phase": 1, "quality": 2,
      "sources": [{ "quest": { "id": 42, "name": "A Small Errand" } }] },
    { "id": 109, "name": "Vendor Trash", "icon": "inv_j", "type": 12, "ilvl": 20, "phase": 1, "quality": 1,
      "sources": [{ "soldBy": { "npcId": 905, "npcName": "A Vendor", "zoneId": 16 } }] },
    { "id": 110, "name": "Suffixed Sword", "icon": "inv_k", "type": 13, "ilvl": 45, "phase": 1, "quality": 2,
      "randomSuffixOptions": [5, 6] },
    { "id": 112, "name": "Two Sources", "icon": "inv_m", "type": 1, "ilvl": 70, "phase": 1, "quality": 4,
      "sources": [
        { "drop": { "difficulty": 1, "npcId": 900, "zoneId": 2717 } },
        { "quest": { "id": 43, "name": "Another Errand" } }
      ] },
    { "id": 113, "name": "Re-itemised Away", "icon": "inv_n", "type": 5, "ilvl": 76, "phase": 1, "quality": 4,
      "sources": [{ "drop": { "difficulty": 1, "npcId": 906, "zoneId": 2717 } }] }
  ],
  "randomSuffixes": [
    { "id": 5, "name": "of Intellect", "stats": [0, 0, 0, 4] },
    { "id": 6, "name": "of Strength", "stats": [7] }
  ],
  "enchants": [
    { "effectId": 41, "spellId": 7420, "name": "Enchant Chest - Minor Health", "type": 5,
      "stats": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12],
      "quality": 1, "phase": 1 },
    { "effectId": 41, "spellId": 7418, "name": "Enchant Bracer - Minor Health", "type": 6,
      "stats": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12],
      "quality": 1 },
    { "effectId": 15, "itemId": 2304, "spellId": 2831, "name": "Light Armor Kit", "type": 5,
      "extraTypes": [9, 7, 10], "enchantType": 3,
      "stats": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8],
      "quality": 1, "classAllowlist": [9, 4] },
    { "effectId": 241, "spellId": 7745, "name": "Enchant 2H Weapon - Minor Impact", "type": 13,
      "enchantType": 1,
      "stats": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12],
      "quality": 1, "phase": 2 }
  ],
  "zones": [
    { "id": 2717, "name": "Molten Core", "expansion": 1 },
    { "id": 1581, "name": "The Deadmines", "expansion": 1 }
  ],
  "npcs": [
    { "id": 900, "name": "Big Boss", "zoneId": 2717 },
    { "id": 902, "name": "Dungeon Boss", "zoneId": 1581 },
    { "id": 903, "name": "Azuregos" },
    { "id": 906, "name": "Re-itemised Boss", "zoneId": 2717 }
  ],
  "factions": [{ "id": 529, "name": "Argent Dawn", "expansion": 1 }],
  "itemIcons": [{ "id": 2304, "name": "Light Armor Kit", "icon": "inv_misc_armorkit_17" }],
  "spellIcons": [
    { "id": 7420, "name": "Enchant Chest - Minor Health", "icon": "spell_holy_chest" },
    { "id": 7418, "name": "Enchant Bracer - Minor Health", "icon": "spell_holy_bracer" },
    { "id": 7745, "name": "Enchant 2H Weapon - Minor Impact", "icon": "spell_fire_impact" }
  ],
  "encounters": []
}
```

- [ ] **Step 2: Write the failing tests**

Create `data/tests/test_forkdb.py`:

```python
from pathlib import Path

import pytest

from pipeline.forkdb import (
    CLASS_SLUGS,
    ENCHANT_TYPES,
    ITEM_TYPE_SLOTS,
    PROFESSIONS,
    REP_LEVELS,
    ForkDbError,
    decode,
    load_fork_database,
)

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"


def test_a_missing_checkout_says_where_it_looked(tmp_path):
    with pytest.raises(ForkDbError, match="assets/database/db.json"):
        load_fork_database(tmp_path)


def test_the_fixture_database_loads_every_table():
    fork = load_fork_database(ENGINE)
    assert len(fork.items) == 13
    assert len(fork.enchants) == 4
    assert len(fork.random_suffixes) == 2
    assert fork.zones == {2717: "Molten Core", 1581: "The Deadmines"}
    assert fork.npcs == {
        900: "Big Boss",
        902: "Dungeon Boss",
        903: "Azuregos",
        906: "Re-itemised Boss",
    }
    assert fork.factions == {529: "Argent Dawn"}
    assert fork.item_icons[2304] == "inv_misc_armorkit_17"
    assert fork.spell_icons[7420] == "spell_holy_chest"


def test_the_tables_are_immutable():
    fork = load_fork_database(ENGINE)
    with pytest.raises(TypeError):
        fork.items.append({})  # type: ignore[attr-defined]


def test_decode_names_the_value_it_could_not_read():
    assert decode(PROFESSIONS, 2, "profession") == "blacksmithing"
    with pytest.raises(ForkDbError, match="profession 7"):
        decode(PROFESSIONS, 7, "profession")


def test_the_enum_tables_match_the_forks_protos():
    """Transcribed from wowsims-forever proto/common.proto and
    proto/ui.proto; a mismatch here is a fork change to pick up."""
    assert PROFESSIONS[11] == "tailoring"
    assert REP_LEVELS[8] == "exalted"
    assert ITEM_TYPE_SLOTS[13] == ("main_hand", "off_hand")
    assert ITEM_TYPE_SLOTS[4] == ("back",)
    assert ENCHANT_TYPES[3] == "kit"
    assert CLASS_SLUGS[9] == "warrior"
    assert sorted(CLASS_SLUGS) == list(range(1, 10))
```

Append to `data/tests/test_simdb_statmap.py`:

```python
def test_stat_keys_is_the_inverse_of_stat_array():
    array = stat_array({"strength": 10, "crit": 2.5, "attack_power": 40})
    assert stat_keys(array) == {"strength": 10.0, "crit": 2.5, "attack_power": 40.0}


def test_stat_keys_drops_zero_amounts():
    assert stat_keys([0, 0, 0, 0]) == {}


def test_stat_keys_reads_a_shared_index_back_as_the_first_name():
    """Forever merges spell and melee hit into one stat, so `hit` and
    `spell_hit` are the same index; the array can only be read back as one
    of them, and PROTO_STAT_ALIASES' order decides which."""
    assert stat_keys(stat_array({"spell_hit": 3})) == {"hit": 3.0}


def test_stat_keys_refuses_an_index_no_planner_key_covers():
    array = [0.0] * (len(pb.Stat.keys()))
    array[pb.Stat.Value("StatMana")] = 5
    with pytest.raises(StatMapError, match="no planner key"):
        stat_keys(array)
```

(the test module already imports `pytest`, `pb`, `StatMapError` and `stat_array`; add `stat_keys` to the `from pipeline.simdb.statmap import …` line. If `StatMana` has a planner key by the time this runs, pick any `pb.Stat` name absent from `PROTO_STAT_ALIASES`' values.)

- [ ] **Step 3: Run the tests to verify they fail**

Run: `cd data && uv run pytest tests/test_forkdb.py tests/test_simdb_statmap.py -q --no-cov`
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.forkdb'`

- [ ] **Step 4: Write `stat_keys`**

In `data/pipeline/simdb/statmap.py`, add after `stat_array`:

```python
def _key_by_index() -> dict[int, str]:
    """Engine stat index -> the planner key it is read back as.

    Not a bijection: Forever merges spell and melee hit into one `Stat`, so
    `hit` and `spell_hit` resolve to the same index (likewise crit). The
    first key `PROTO_STAT_ALIASES` declares for an index wins, which is why
    the merged names are declared before the vanilla-lineage ones.
    """
    by_index: dict[int, str] = {}
    for key in PROTO_STAT_ALIASES:
        if key.startswith("__"):
            continue
        by_index.setdefault(stat_index(key), key)
    return by_index


def stat_keys(array: Iterable[float]) -> dict[str, float]:
    """The planner stat keys a `repeated double stats` array carries.

    The inverse of `stat_array`, for reading the engine fork's own database
    (enchants and random suffixes state their stats as that array and
    nothing else). Zero amounts are dropped: an absent key means the row
    grants none of that stat, which is what every consumer already assumes
    of `GearItem.stats`.
    """
    by_index = _key_by_index()
    out: dict[str, float] = {}
    for index, amount in enumerate(array):
        if not amount:
            continue
        key = by_index.get(index)
        if key is None:
            raise StatMapError(
                f"stat index {index} has no planner key; add one to "
                f"PROTO_STAT_ALIASES in pipeline/simdb/statmap.py"
            )
        out[key] = float(amount)
    return out
```

- [ ] **Step 5: Write the fork database reader**

Create `data/pipeline/forkdb.py`:

```python
"""The engine fork's own item database, read for the facts DB2 cannot state.

`assets/database/db.json` in the wowsims-forever checkout is the fork's
`UIDatabase`: items with their AtlasLoot and Wowhead sources, the enchants
the fork's own UI offers, random suffixes, the instance zones, the NPCs it
names, and the reputation factions. Nothing in a DB2 export says which boss
drops an item, which profession crafts it, or what "of the Bear" is worth,
so for `loot.json`, `enchants.json` and `suffixes.json` this file is the
only source there is.

It is read from a checkout at a path and never vendored here: it moves with
the engine pin, the same way `pipeline/simproto`'s bindings do. `make loot`
passes the local checkout; the data workflow clones the fork at the sha in
`sim/enginever/version.go` and passes that.

Measured on the pinned fork, 2026-09-19: 7,553 items (4,481 with sources),
1,168 random suffixes, 173 enchant rows over 150 distinct effect ids, 27
instance zones, 274 named NPCs, 18 factions.

Every enum below is transcribed from the fork's `proto/common.proto` and
`proto/ui.proto`. A value that is not in one of them raises rather than
being guessed at or dropped -- the same policy as
`pipeline/normalize/gear.py`'s `STAT_BY_MODIFIER_ID`.
"""

from __future__ import annotations

import json
from dataclasses import dataclass
from pathlib import Path
from typing import TypeVar

#: Where the database sits inside an engine checkout.
DB_RELATIVE = Path("assets/database/db.json")

T = TypeVar("T")


class ForkDbError(SystemExit):
    """The fork's database is missing, or says something we cannot decode."""


#: proto/common.proto's `Profession`. 6 and 7 are unused in the fork's enum.
PROFESSIONS: dict[int, str] = {
    1: "alchemy",
    2: "blacksmithing",
    3: "enchanting",
    4: "engineering",
    5: "herbalism",
    8: "leatherworking",
    9: "mining",
    10: "skinning",
    11: "tailoring",
}

#: proto/ui.proto's `RepLevel`, minus `RepLevelUnknown`.
REP_LEVELS: dict[int, str] = {
    1: "hated",
    2: "hostile",
    3: "unfriendly",
    4: "neutral",
    5: "friendly",
    6: "honored",
    7: "revered",
    8: "exalted",
}

#: proto/common.proto's `ItemType` -> the planner's slot names, the same
#: vocabulary `pipeline/normalize/gear.py`'s SLOT_BY_INVENTORY_TYPE emits.
#: `ItemTypeWeapon` covers both hands because an enchant that can go on a
#: weapon can go on either.
ITEM_TYPE_SLOTS: dict[int, tuple[str, ...]] = {
    1: ("head",),
    2: ("neck",),
    3: ("shoulder",),
    4: ("back",),
    5: ("chest",),
    6: ("wrist",),
    7: ("hands",),
    8: ("waist",),
    9: ("legs",),
    10: ("feet",),
    11: ("finger",),
    12: ("trinket",),
    13: ("main_hand", "off_hand"),
    14: ("ranged",),
}

#: proto/common.proto's `EnchantType`: what shape of item the enchant needs,
#: which is a different question from which slot it goes in.
ENCHANT_TYPES: dict[int, str] = {
    0: "normal",
    1: "two_hand",
    2: "shield",
    3: "kit",
    4: "staff",
}

#: proto/common.proto's `Class` -> the site's class slugs (the values
#: `pipeline/normalize/classes.py`'s slugify produces from ChrClasses).
CLASS_SLUGS: dict[int, str] = {
    1: "druid",
    2: "hunter",
    3: "mage",
    4: "paladin",
    5: "priest",
    6: "rogue",
    7: "shaman",
    8: "warlock",
    9: "warrior",
}


def decode(table: dict[int, T], value: int, what: str) -> T:
    """`table[value]`, or a clear error naming the enum and the number.

    Widening a filter to make an unknown value pass is how a wrong fact
    ships quietly; this is the one place that refuses instead.
    """
    if value not in table:
        raise ForkDbError(
            f"the fork database has {what} {value}, which pipeline/forkdb.py does "
            f"not decode; add it from the fork's protos (known: {sorted(table)})"
        )
    return table[value]


@dataclass(frozen=True)
class ForkDatabase:
    """The five tables of db.json anything here reads, plus the two icon maps."""

    items: tuple[dict, ...]
    enchants: tuple[dict, ...]
    random_suffixes: tuple[dict, ...]
    zones: dict[int, str]
    npcs: dict[int, str]
    factions: dict[int, str]
    item_icons: dict[int, str]
    spell_icons: dict[int, str]


def _named(rows: list[dict]) -> dict[int, str]:
    return {int(row["id"]): row["name"] for row in rows}


def _icons(rows: list[dict]) -> dict[int, str]:
    return {int(row["id"]): row["icon"] for row in rows if row.get("icon")}


def load_fork_database(engine_dir: Path) -> ForkDatabase:
    path = engine_dir / DB_RELATIVE
    if not path.exists():
        raise ForkDbError(
            f"no fork item database at {path}; point --engine at a "
            f"wowsims-forever checkout"
        )
    raw = json.loads(path.read_text(encoding="utf-8"))
    missing = sorted({"items", "enchants", "randomSuffixes", "zones", "npcs"} - set(raw))
    if missing:
        raise ForkDbError(f"{path} has no {', '.join(missing)}; is this the fork's db.json?")
    return ForkDatabase(
        items=tuple(raw["items"]),
        enchants=tuple(raw["enchants"]),
        random_suffixes=tuple(raw["randomSuffixes"]),
        zones=_named(raw["zones"]),
        npcs=_named(raw["npcs"]),
        factions=_named(raw.get("factions", [])),
        item_icons=_icons(raw.get("itemIcons", [])),
        spell_icons=_icons(raw.get("spellIcons", [])),
    )
```

- [ ] **Step 6: Run the tests**

Run: `cd data && uv run pytest tests/test_forkdb.py tests/test_simdb_statmap.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 7: Commit**

```bash
git add data/pipeline/forkdb.py data/pipeline/simdb/statmap.py data/tests/test_forkdb.py data/tests/test_simdb_statmap.py data/tests/fixtures/loot/db.json
git commit -m "$(cat <<'EOF'
feat(data): read the engine fork's item database

Contract section 6's three files all come out of the fork's own
assets/database/db.json, which is the only source for which boss drops
an item, which profession crafts it, or what a random suffix is worth.
The reader takes an engine checkout so it moves with the pin, and every
enum it decodes is transcribed from the fork's protos with an unknown
value raising rather than being guessed.

statmap gains stat_keys, the inverse of stat_array, because the fork
states enchant and suffix stats as the engine's own float array.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 4: `suffixes.json`, and the two fork columns on `items.json`

Contract 6.3: `data/builds/<build>/suffixes.json` is `[ { "id", "name", "stats": {...} } ]`, and `items.json` rows gain `suffixes: [id, ...]` where the item rolls one.

The contract sources this from `ItemRandomSuffix`. **That table 404s on build 1.60.1.69893** (it exists for Classic Era 1.15.9.69722 and not for the beta client), which is exactly why `pipeline/simdb/__init__.py` already emits `SimDatabase.random_suffixes` empty. The fork database has the table the client does not: `randomSuffixes` (1,168 rows) and `randomSuffixOptions` on 1,688 items. This task reads those.

Measured on the pinned fork against build 1.60.1.69893: **1,168 suffix records**, and **69** of the build's 19,171 items carry suffix options (the fork's item table is Era's, and only 2,811 of its 7,553 ids exist in this client's `ItemSparse`).

`items.json` gains a **second** fork column in the same pass, `faction_restriction`, which contract 10.3 needs on `SimItem` (task 15) and `sim/bulk`'s expander needs for its faction rule. It is here and not in `simdb` because **the client does not state it**: 19,066 of the 1.60 client's 19,171 `ItemSparse` rows have `AllowableRace` `-1/-1`, including every one of the 819 items the fork database marks Alliance- or Horde-only. Keeping both fork-derived columns in one place means `pipeline/simdb` never needs an engine checkout. Values are `UIItem.FactionRestriction` in snake case: `alliance_only`, `horde_only`, or absent.

**Files:**
- Modify: `data/pipeline/models.py` (`Item.suffixes`, `Item.faction_restriction`, new `SuffixRecord`)
- Modify: `data/pipeline/forkdb.py` (`FACTION_RESTRICTIONS`)
- Create: `data/pipeline/loot/__init__.py` (package marker only in this task), `data/pipeline/loot/gear.py`
- Modify: `data/tests/golden/items.json` (13 rows gain `"suffixes": []` and `"faction_restriction": ""`)
- Create: `data/tests/fixtures/loot/items.json`
- Test: `data/tests/test_loot_gear.py`

**Interfaces:**
- Consumes: `pipeline.forkdb.ForkDatabase`, `pipeline.forkdb.load_fork_database`, `pipeline.forkdb.FACTION_RESTRICTIONS`, `decode`, `pipeline.simdb.statmap.stat_keys`
- Produces:
  - `pipeline.models.SuffixRecord` with `id: int`, `name: str`, `stats: dict[str, float]`
  - `pipeline.models.Item.suffixes: list[int]` (default `[]`) and `Item.faction_restriction: str` (default `""`), in that order, both last
  - `pipeline.loot.gear.build_suffixes(fork: ForkDatabase) -> list[SuffixRecord]`
  - `pipeline.loot.gear.suffix_options(fork: ForkDatabase) -> dict[int, list[int]]`
  - `pipeline.loot.gear.faction_restrictions(fork: ForkDatabase) -> dict[int, str]`
  - `pipeline.loot.gear.apply_fork_columns(build_dir: Path, options: dict[int, list[int]], restrictions: dict[int, str]) -> tuple[int, int]` (rows with a non-empty suffix list, rows with a restriction)

- [ ] **Step 1: Write the fixture item table**

Create `data/tests/fixtures/loot/items.json` — the build's own item table, in `items.json` shape. It deliberately omits item 113, which the fork database gives a raid drop, so the "re-itemised away" filter of contract 10.4 is exercised:

```json
[
  { "id": 100, "name": "Raid Boss Drop", "quality": 4, "item_level": 76, "required_level": 60, "class_id": 4, "subclass_id": 4, "inventory_type": 5 },
  { "id": 101, "name": "Raid Trash Drop", "quality": 3, "item_level": 66, "required_level": 58, "class_id": 4, "subclass_id": 3, "inventory_type": 9 },
  { "id": 102, "name": "Nameless Boss Drop", "quality": 4, "item_level": 76, "required_level": 60, "class_id": 4, "subclass_id": 4, "inventory_type": 7 },
  { "id": 103, "name": "Dungeon Boss Drop", "quality": 3, "item_level": 50, "required_level": 45, "class_id": 4, "subclass_id": 2, "inventory_type": 6 },
  { "id": 104, "name": "World Boss Drop", "quality": 4, "item_level": 74, "required_level": 60, "class_id": 4, "subclass_id": 0, "inventory_type": 2 },
  { "id": 106, "name": "Crafted Plate", "quality": 3, "item_level": 60, "required_level": 55, "class_id": 4, "subclass_id": 4, "inventory_type": 5 },
  { "id": 107, "name": "Exalted Cloak", "quality": 4, "item_level": 68, "required_level": 60, "class_id": 4, "subclass_id": 1, "inventory_type": 16 },
  { "id": 108, "name": "Quest Ring", "quality": 2, "item_level": 40, "required_level": 35, "class_id": 4, "subclass_id": 0, "inventory_type": 11 },
  { "id": 110, "name": "Suffixed Sword", "quality": 2, "item_level": 45, "required_level": 40, "class_id": 2, "subclass_id": 7, "inventory_type": 13 },
  { "id": 111, "name": "Rank Eleven Blade", "quality": 3, "item_level": 65, "required_level": 60, "class_id": 2, "subclass_id": 7, "inventory_type": 13 },
  { "id": 112, "name": "Two Sources", "quality": 4, "item_level": 70, "required_level": 60, "class_id": 4, "subclass_id": 4, "inventory_type": 1 }
]
```

- [ ] **Step 2: Add the faction enum to the fork reader**

In `data/pipeline/forkdb.py`, beside the other enum tables:

```python
#: proto/ui.proto's `UIItem.FactionRestriction`, minus its unspecified
#: zero. The client cannot supply this on build 1.60.1.69893 -- 19,066 of
#: its 19,171 ItemSparse rows have AllowableRace -1/-1, the 819 the fork
#: marks restricted among them -- so the fork is the only source.
FACTION_RESTRICTIONS: dict[int, str] = {1: "alliance_only", 2: "horde_only"}
```

- [ ] **Step 3: Write the failing test**

Create `data/tests/test_loot_gear.py`:

```python
import json
import shutil
from pathlib import Path

from pipeline.forkdb import load_fork_database
from pipeline.loot.gear import (
    apply_fork_columns,
    build_suffixes,
    faction_restrictions,
    suffix_options,
)

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"


def fork():
    return load_fork_database(ENGINE)


def test_every_suffix_is_emitted_with_its_stats_named():
    records = build_suffixes(fork())
    assert [(r.id, r.name) for r in records] == [(5, "of Intellect"), (6, "of Strength")]
    assert records[0].stats == {"intellect": 4.0}
    assert records[1].stats == {"strength": 7.0}


def test_suffix_options_index_only_the_items_that_roll_one():
    assert suffix_options(fork()) == {110: [5, 6]}


def test_faction_restrictions_index_only_the_restricted_items():
    assert faction_restrictions(fork()) == {100: "alliance_only"}


def prepared(tmp_path):
    build_dir = tmp_path / "1.60.1.69893"
    build_dir.mkdir()
    shutil.copy(ENGINE / "items.json", build_dir / "items.json")
    return build_dir


def test_apply_fork_columns_fills_both_and_leaves_the_rest_empty(tmp_path):
    build_dir = prepared(tmp_path)
    with_suffixes, restricted = apply_fork_columns(
        build_dir, suffix_options(fork()), faction_restrictions(fork())
    )
    assert (with_suffixes, restricted) == (1, 1)
    rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    by_id = {row["id"]: row for row in rows}
    assert by_id[110]["suffixes"] == [5, 6]
    assert by_id[100]["suffixes"] == []
    assert by_id[100]["faction_restriction"] == "alliance_only"
    assert by_id[110]["faction_restriction"] == ""
    assert len(rows) == 11


def test_apply_fork_columns_keeps_the_key_order_and_appends_the_new_ones(tmp_path):
    build_dir = prepared(tmp_path)
    apply_fork_columns(build_dir, suffix_options(fork()), faction_restrictions(fork()))
    rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    assert list(rows[0]) == [
        "id",
        "name",
        "quality",
        "item_level",
        "required_level",
        "class_id",
        "subclass_id",
        "inventory_type",
        "suffixes",
        "faction_restriction",
    ]


def test_apply_fork_columns_is_idempotent(tmp_path):
    build_dir = prepared(tmp_path)
    options, restrictions = suffix_options(fork()), faction_restrictions(fork())
    apply_fork_columns(build_dir, options, restrictions)
    once = (build_dir / "items.json").read_bytes()
    apply_fork_columns(build_dir, options, restrictions)
    assert (build_dir / "items.json").read_bytes() == once
```

- [ ] **Step 4: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_loot_gear.py -q --no-cov`
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.loot'`

- [ ] **Step 5: Add the models**

In `data/pipeline/models.py`, add `suffixes` to `Item` (last, so the emitted key order stays append-only):

```python
class Item(BaseModel):
    id: int
    name: str
    quality: int
    item_level: int
    required_level: int
    class_id: int
    subclass_id: int
    inventory_type: int
    #: The random suffixes this item rolls, from the engine fork's
    #: `randomSuffixOptions` (parity contract 6.3). Empty for an item that
    #: rolls none, and for every row until `python -m pipeline loot` has run
    #: for the build -- `normalize` has no fork database to read.
    suffixes: list[int] = []
    #: "alliance_only", "horde_only", or "" for no restriction, from the
    #: fork's `factionRestriction` (contract 10.3's `SimItem` field). The
    #: client states none of this on build 1.60.1.69893, so like `suffixes`
    #: it is empty until `loot` has run.
    faction_restriction: str = ""
```

and add, beside `ItemSetRecord`:

```python
class SuffixRecord(BaseModel):
    id: int
    name: str
    stats: dict[str, float]
```

- [ ] **Step 6: Create the package and write the builders**

Create `data/pipeline/loot/__init__.py` with just a docstring for now (task 9 fills it in):

```python
"""The Droptimizer and Top Gear data: loot.json, enchants.json, suffixes.json."""
```

Create `data/pipeline/loot/gear.py`:

```python
"""The two gear tables Top Gear needs, and the column they add to items.json.

Parity contract 6.2 and 6.3. Both come from the engine fork's database
rather than from DB2:

* **Suffixes.** `ItemRandomSuffix` 404s on build 1.60.1.69893 -- which is
  why `pipeline/simdb/__init__.py` already emits `SimDatabase.random_suffixes`
  empty -- so the fork's `randomSuffixes` and its per-item
  `randomSuffixOptions` are the only statement of what "of the Bear" is
  worth and which items roll it.
* **Enchants.** The client's `SpellItemEnchantment` has 2,216 rows and no
  idea which of them a player can actually apply; the fork's `UIEnchant` is
  the curated 173 its own UI offers, with the slots, classes and phase the
  picker needs.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.forkdb import FACTION_RESTRICTIONS, ForkDatabase, decode
from pipeline.models import Item, SuffixRecord
from pipeline.normalize import write_json
from pipeline.simdb.statmap import stat_keys


def build_suffixes(fork: ForkDatabase) -> list[SuffixRecord]:
    return sorted(
        (
            SuffixRecord(
                id=int(row["id"]),
                name=row["name"],
                stats=stat_keys(row.get("stats", [])),
            )
            for row in fork.random_suffixes
        ),
        key=lambda record: record.id,
    )


def suffix_options(fork: ForkDatabase) -> dict[int, list[int]]:
    """Item id -> the suffix ids it rolls, for the items that roll any."""
    return {
        int(row["id"]): sorted(int(suffix) for suffix in row["randomSuffixOptions"])
        for row in fork.items
        if row.get("randomSuffixOptions")
    }


def faction_restrictions(fork: ForkDatabase) -> dict[int, str]:
    """Item id -> "alliance_only" or "horde_only", for the items so marked.

    The client cannot answer this on build 1.60.1.69893: 19,066 of its
    19,171 `ItemSparse` rows carry `AllowableRace` -1/-1, and every one of
    the 819 items the fork marks restricted is among them.
    """
    return {
        int(row["id"]): decode(
            FACTION_RESTRICTIONS, int(row["factionRestriction"]), "faction restriction"
        )
        for row in fork.items
        if row.get("factionRestriction")
    }


def apply_fork_columns(
    build_dir: Path,
    options: dict[int, list[int]],
    restrictions: dict[int, str],
) -> tuple[int, int]:
    """Fill `items.json`'s two fork-derived columns.

    Returns (rows with a suffix list, rows with a faction restriction).

    Read-modify-write through the `Item` model rather than through the raw
    dicts, so the file comes back out in exactly the key order and
    serialization `normalize` would have written -- re-running this on an
    already-filled build rewrites the same bytes.
    """
    path = build_dir / "items.json"
    if not path.exists():
        raise SystemExit(f"no {path}; run `python -m pipeline normalize` for this build first")
    rows = json.loads(path.read_text(encoding="utf-8"))
    items = [
        Item(**row).model_copy(
            update={
                "suffixes": options.get(int(row["id"]), []),
                "faction_restriction": restrictions.get(int(row["id"]), ""),
            }
        )
        for row in rows
    ]
    write_json(items, path)
    return (
        sum(1 for item in items if item.suffixes),
        sum(1 for item in items if item.faction_restriction),
    )
```

- [ ] **Step 7: Update the normalize golden**

`tests/golden/items.json` is the golden for `normalize_items`, and the new model field makes every row gain `"suffixes": []`. Regenerate rather than hand-edit:

Run: `cd data && uv run pytest tests/test_normalize_items.py -q --no-cov`
Expected: FAIL on the golden comparison. Then regenerate with:

```bash
cd data && uv run python - <<'EOF'
from pathlib import Path
from pipeline.csvio import read_csv
from pipeline.normalize import write_json
from pipeline.normalize.items import normalize_items
here = Path("tests")
write_json(
    normalize_items(
        read_csv(here / "fixtures/ItemSparse.csv"), read_csv(here / "fixtures/Item.csv")
    ),
    here / "golden/items.json",
)
EOF
git diff --stat data/tests/golden/items.json
```
Expected: 13 rows each gain a `"suffixes": []` and a `"faction_restriction": ""` line.

- [ ] **Step 8: Run the tests**

Run: `cd data && uv run pytest tests/test_loot_gear.py tests/test_normalize_items.py tests/test_normalize_build.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 9: Commit**

```bash
git add data/pipeline/models.py data/pipeline/loot/ data/tests/test_loot_gear.py data/tests/fixtures/loot/items.json data/tests/golden/items.json
git commit -m "$(cat <<'EOF'
feat(data): random suffixes, and the two fork columns on items.json

Contract 6.3. The contract names ItemRandomSuffix, which 404s on build
1.60.1.69893 -- the same absence that already leaves SimDatabase's
random_suffixes empty -- so the fork database's randomSuffixes and
per-item randomSuffixOptions are the source instead.

items.json's rows gain `suffixes` and `faction_restriction`, appended
last and defaulting to empty, so a build that has not had `pipeline
loot` run for it still emits both keys. faction_restriction is here
rather than in simdb because the client states none of it -- 19,066 of
19,171 ItemSparse rows are AllowableRace -1/-1, the 819 restricted ones
among them -- and keeping both fork columns together means simdb never
needs an engine checkout.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 5: `enchants.json`

Contract 6.2: `[ { "id": <effect_id>, "name", "icon", "slots": ["head", ...], "item_types": [...], "classes": [...], "stats": {...}, "phase" } ]` from the fork database's `UIEnchant`.

Two facts the contract does not anticipate, both measured:

- **`effect_id` is not unique.** 173 rows over 150 distinct effect ids: 19 ids appear twice or three times (effect 247 is Minor Agility on a cloak, a bracer *and* a boot, each a different spell). The fork's own comment says uniqueness needs effect id plus slot, or effect id plus item/spell id. So the rows carry `spell_id` and `item_id` as well, appended after the contract's keys; `id` stays the effect id.
- **`slots` and `item_types` would otherwise be the same field.** Both can only come from `UIEnchant.type` + `extra_types`. `slots` is that, mapped to the planner's slot names; `item_types` is `enchant_type` — `normal`, `two_hand`, `shield`, `kit`, `staff` — which is the other restriction an enchant carries and the only reading that gives the two fields distinct content.

Measured on the pinned fork: **173 rows**, all 150 distinct effect ids present in the build's `simdb.bin`, every row's icon resolvable, 30 rows with a `phase`, 9 with a class allowlist.

**Files:**
- Modify: `data/pipeline/models.py` (`EnchantRecord`)
- Modify: `data/pipeline/loot/gear.py` (`build_enchants`)
- Test: `data/tests/test_loot_gear.py` (extend)

**Interfaces:**
- Consumes: `pipeline.forkdb.ForkDatabase`, `pipeline.forkdb.ITEM_TYPE_SLOTS`, `ENCHANT_TYPES`, `CLASS_SLUGS`, `decode`; `pipeline.simdb.statmap.stat_keys`; `pipeline.normalize.write_records`
- Produces: `pipeline.models.EnchantRecord` with `id: int`, `name: str`, `icon: str`, `slots: list[str]`, `item_types: list[str]`, `classes: list[str]`, `stats: dict[str, float]`, `phase: int`, `spell_id: int`, `item_id: int`; `pipeline.loot.gear.build_enchants(fork: ForkDatabase) -> list[EnchantRecord]`

- [ ] **Step 1: Write the failing test**

Append to `data/tests/test_loot_gear.py`:

```python
from pipeline.loot.gear import build_enchants  # add to the existing import line


def enchant(effect_id: int, spell_id: int):
    for record in build_enchants(fork()):
        if record.id == effect_id and record.spell_id == spell_id:
            return record
    raise AssertionError(f"no enchant {effect_id}/{spell_id}")


def test_every_fork_enchant_row_is_emitted_even_when_it_shares_an_effect_id():
    records = build_enchants(fork())
    assert len(records) == 4
    assert [r.id for r in records] == [15, 41, 41, 241]
    assert [r.spell_id for r in records if r.id == 41] == [7418, 7420]


def test_an_enchants_slots_come_from_its_item_type_and_extra_types():
    assert enchant(41, 7420).slots == ["chest"]
    assert enchant(41, 7418).slots == ["wrist"]
    assert enchant(15, 2831).slots == ["chest", "feet", "hands", "legs"]
    assert enchant(241, 7745).slots == ["main_hand", "off_hand"]


def test_item_types_is_the_shape_restriction_not_the_slot():
    assert enchant(41, 7420).item_types == ["normal"]
    assert enchant(15, 2831).item_types == ["kit"]
    assert enchant(241, 7745).item_types == ["two_hand"]


def test_classes_are_slugs_and_sorted_and_empty_means_anyone():
    assert enchant(15, 2831).classes == ["paladin", "warrior"]
    assert enchant(41, 7420).classes == []


def test_icons_come_from_the_item_when_there_is_one_and_the_spell_otherwise():
    assert enchant(15, 2831).icon == "inv_misc_armorkit_17"
    assert enchant(41, 7420).icon == "spell_holy_chest"


def test_stats_and_phase_are_read_off_the_row():
    assert enchant(15, 2831).stats == {"armor": 8.0}
    assert enchant(241, 7745).stats == {"attack_power": 12.0}
    assert enchant(241, 7745).phase == 2
    assert enchant(41, 7418).phase == 0
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_loot_gear.py -q --no-cov`
Expected: FAIL — `ImportError: cannot import name 'build_enchants'`

- [ ] **Step 3: Add the model**

In `data/pipeline/models.py`, beside `SuffixRecord`:

```python
class EnchantRecord(BaseModel):
    #: The effect id. NOT unique: the fork's own table repeats an effect
    #: across the slots it can go in (19 of its 150 ids), which is why
    #: `spell_id` and `item_id` are here too.
    id: int
    name: str
    icon: str
    #: The planner slot names this enchant can be applied in.
    slots: list[str]
    #: The shape of item it needs: normal, two_hand, shield, kit or staff.
    item_types: list[str]
    #: Class slugs allowed to use it. Empty means no restriction.
    classes: list[str]
    stats: dict[str, float]
    phase: int
    #: Appended after the contract's keys, for telling two rows with the
    #: same effect id apart. 0 where the fork states neither.
    spell_id: int
    item_id: int
```

- [ ] **Step 4: Write the builder**

Append to `data/pipeline/loot/gear.py` (and extend the imports at its top with `EnchantRecord` and `from pipeline.forkdb import CLASS_SLUGS, ENCHANT_TYPES, ITEM_TYPE_SLOTS, ForkDatabase, decode`):

```python
def _enchant_icon(fork: ForkDatabase, row: dict) -> str:
    """The enchant's art: its own item's icon where it has one, else its
    spell's. Every row on the pinned fork resolves through one or the
    other; a row that resolved through neither would reach the site as
    `icons/.webp`, so it is refused rather than emitted blank."""
    item_id, spell_id = int(row.get("itemId", 0)), int(row.get("spellId", 0))
    icon = fork.item_icons.get(item_id) or fork.spell_icons.get(spell_id)
    if not icon:
        raise SystemExit(
            f"enchant {row['effectId']} ({row.get('name', '?')}) has no icon in the "
            f"fork's itemIcons or spellIcons"
        )
    return icon


def build_enchants(fork: ForkDatabase) -> list[EnchantRecord]:
    records = [
        EnchantRecord(
            id=int(row["effectId"]),
            name=row["name"],
            icon=_enchant_icon(fork, row),
            slots=sorted(
                {
                    slot
                    for item_type in [row["type"], *row.get("extraTypes", [])]
                    for slot in decode(ITEM_TYPE_SLOTS, int(item_type), "enchant item type")
                }
            ),
            item_types=[decode(ENCHANT_TYPES, int(row.get("enchantType", 0)), "enchant type")],
            classes=sorted(
                decode(CLASS_SLUGS, int(class_id), "enchant class")
                for class_id in row.get("classAllowlist", [])
            ),
            stats=stat_keys(row.get("stats", [])),
            phase=int(row.get("phase", 0)),
            spell_id=int(row.get("spellId", 0)),
            item_id=int(row.get("itemId", 0)),
        )
        for row in fork.enchants
    ]
    # Not `write_json`'s sort: `id` alone is not unique here, so the tie is
    # broken by the spell and then the item, which together are.
    return sorted(records, key=lambda record: (record.id, record.spell_id, record.item_id))
```

- [ ] **Step 5: Run the tests**

Run: `cd data && uv run pytest tests/test_loot_gear.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 6: Commit**

```bash
git add data/pipeline/models.py data/pipeline/loot/gear.py data/tests/test_loot_gear.py
git commit -m "$(cat <<'EOF'
feat(data): the enchant table Top Gear's picker reads

Contract 6.2, from the fork's UIEnchant rather than the client's
SpellItemEnchantment: 2,216 client rows say nothing about which enchant
a player can apply, and the fork's 173 are the curated list its own UI
offers.

Two things the contract did not anticipate, both measured: effect ids
repeat across slots (19 of 150), so the rows also carry spell_id and
item_id; and `slots` and `item_types` can only be distinct if the first
is the item type mapped to planner slots and the second the
EnchantType shape restriction.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 6: `loot.json`

Contract 6.1 as corrected by 10.4. Seven kinds of source, built from the fork database's `sources`, joined to the build's own zones — **and listing only items the build has**.

**How each kind is decided** (every rule below is measured against the pinned fork and build 1.60.1.69893, not guessed):

| Kind | From | Source id |
| --- | --- | --- |
| `raid` | a `drop` whose zone is a `Map.InstanceType == 2` instance; one boss per `npcId`, the rest into `trash` | `raid:<zone-slug>`, boss `raid:<zone-slug>:<npc-id>` |
| `dungeon` | the same with `InstanceType == 1` | `dungeon:<zone-slug>`, boss `dungeon:<zone-slug>:<npc-id>` |
| `world` | a `drop` outside any instance **whose `npcId` the fork's `npcs` table names**. On the pinned fork exactly two do: Azuregos (6109) and Lord Kazzak (12397); the other 212 open-world drop NPCs are unnamed trash mobs and are dropped | `world:<npc-slug>` |
| `crafted` | a `crafted` source, grouped by `Profession` | `crafted:<profession>` |
| `rep` | a `rep` source, grouped by faction and standing | `rep:<faction-slug>:<standing>` |
| `pvp` | the **client's** `ItemSparse.RequiredPVPRank`, not the fork — the fork has no rank data, and 266 of its 273 vendor-only items are rank sets | `pvp:rank-<n>` |
| `quest` | every `quest` source, one flat list | `quest` |

A `soldBy` source with no PvP rank has no kind in the contract's vocabulary and is dropped. So is an open-world drop from an unnamed NPC. Together that is **864 dropped source entries** on the pinned fork, and that number is logged.

**The build filter (contract 10.4).** The fork's item table is Classic Era's and Forever has re-itemised: only 2,811 of its 7,553 ids exist in this client's `ItemSparse`, and **1,809** of the ids its sources name do not. An id the build does not have would render as a blank Droptimizer row and could never be simmed, so it is dropped and counted. A boss left with no items is dropped; a source left with no boss, trash or items is dropped. That is what turns 78 sources into **62**, and what makes raid Droptimizer honest and thin (Molten Core keeps 4 of its 11 bosses and 10 of its 139 items; Onyxia's Lair disappears entirely, all 16 of her items being gone from the client). The first overlay records the gap per raid.

**The zone join needs two steps, in this order**, because neither covers every instance:

1. `Map.AreaTableID` → the map, whose `InstanceType` is the answer. Resolves 19 of the fork's 27 instance zones, Molten Core, Blackwing Lair, Onyxia's Lair and Zul'Gurub among them.
2. The build's own `zones.json` `map_id` (which is `AreaTable.ContinentID`) → the map. Resolves the other 8 — Deadmines, Stratholme, Scholomance, Scarlet Monastery, Zul'Farrak, both Razorfens, Shadowfang Keep — and gets **Onyxia's Lair wrong**, calling it Kalimdor, which is why it is the fallback and not the first join.

That join also finds three raid zones the fork's own `zones` list omits but its drops reference: Ahn'Qiraj (3428), Ruins of Ahn'Qiraj (3429) and Naxxramas (3456).

**Files:**
- Modify: `data/pipeline/models.py` (`LootBoss`, `LootSource`, `LootFile`)
- Modify: `data/pipeline/normalize/__init__.py` (`write_document`)
- Create: `data/pipeline/loot/sources.py`
- Create: `data/tests/fixtures/loot/Map.csv`, `data/tests/fixtures/loot/ItemSparse.csv`, `data/tests/fixtures/loot/zones.json`
- Test: `data/tests/test_loot_sources.py`

**Interfaces:**
- Consumes: `pipeline.forkdb.ForkDatabase`, `PROFESSIONS`, `REP_LEVELS`, `decode`; `pipeline.normalize.classes.slugify`; `pipeline.csvio.read_csv`, `populated`
- Produces:
  - `pipeline.models.LootBoss(id: str, name: str, npc_id: int, items: list[int])`
  - `pipeline.models.LootSource(id, kind, name, zone_id=None, opens=None, profession=None, faction_id=None, standing=None, rank=None, bosses=None, trash=None, items=None)`
  - `pipeline.models.LootFile(sources: list[LootSource])`
  - `pipeline.normalize.write_document(record: BaseModel, path: Path) -> None`
  - `pipeline.loot.sources.KIND_ORDER: tuple[str, ...]`
  - `pipeline.loot.sources.instance_types(map_rows, zone_rows) -> dict[int, int]`
  - `pipeline.loot.sources.pvp_ranks(sparse_rows) -> dict[int, int]`
  - `pipeline.loot.sources.build_loot(fork, zone_names: dict[int, str], types: dict[int, int], ranks: dict[int, int], build_items: set[int]) -> tuple[LootFile, LootStats]`
  - `pipeline.loot.sources.LootStats(items: int, dropped_entries: int, absent_items: int)`

- [ ] **Step 1: Write the fixture client tables**

Create `data/tests/fixtures/loot/Map.csv` — Molten Core resolves through `AreaTableID`, Deadmines only through `zones.json`'s `map_id`, Kalimdor is not an instance at all:

```
ID,Directory,MapName_lang,InstanceType,AreaTableID
409,MoltenCore,Molten Core,2,2717
36,DeadminesInstance,Deadmines,1,0
1,Kalimdor,Kalimdor,0,0
```

Create `data/tests/fixtures/loot/zones.json` — the build's zone table, in `zones.json` shape:

```json
[
  { "id": 16, "name": "Azshara", "map_id": 1, "map_name": "Kalimdor", "parent_id": null, "level_min": 45, "level_max": null },
  { "id": 1581, "name": "The Deadmines", "map_id": 36, "map_name": "Deadmines", "parent_id": null, "level_min": null, "level_max": null },
  { "id": 2717, "name": "Molten Core", "map_id": 409, "map_name": "Molten Core", "parent_id": null, "level_min": null, "level_max": null }
]
```

Create `data/tests/fixtures/loot/ItemSparse.csv` — only the two columns this reads, plus the socket columns so the same fixture can serve the guard:

```
ID,Display_lang,RequiredPVPRank,SocketType_0,SocketType_1,SocketType_2
110,Suffixed Sword,0,0,0,0
111,Rank Eleven Blade,11,0,0,0
112,Two Sources,0,0,0,0
```

- [ ] **Step 2: Write the failing test**

Create `data/tests/test_loot_sources.py`:

```python
import json
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.forkdb import load_fork_database
from pipeline.loot.sources import KIND_ORDER, build_loot, instance_types, pvp_ranks

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"


def zone_rows():
    return json.loads((ENGINE / "zones.json").read_text(encoding="utf-8"))


def built():
    fork = load_fork_database(ENGINE)
    rows = zone_rows()
    build_items = {
        row["id"]
        for row in json.loads((ENGINE / "items.json").read_text(encoding="utf-8"))
    }
    return build_loot(
        fork,
        {row["id"]: row["name"] for row in rows},
        instance_types(read_csv(ENGINE / "Map.csv"), rows),
        pvp_ranks(read_csv(ENGINE / "ItemSparse.csv")),
        build_items,
    )


def source(source_id: str):
    document, _ = built()
    for candidate in document.sources:
        if candidate.id == source_id:
            return candidate
    raise AssertionError(f"no source {source_id}; have {[s.id for s in document.sources]}")


def source_item_ids(entry) -> set[int]:
    return set(
        (entry.items or [])
        + (entry.trash or [])
        + [item for boss in (entry.bosses or []) for item in boss.items]
    )


def test_the_map_join_prefers_area_table_id_and_falls_back_to_the_zones_map():
    types = instance_types(read_csv(ENGINE / "Map.csv"), zone_rows())
    assert types[2717] == 2  # through Map.AreaTableID
    assert types[1581] == 1  # through zones.json's map_id
    assert 16 not in types  # Azshara's map is not an instance


def test_pvp_ranks_index_only_the_items_that_require_one():
    assert pvp_ranks(read_csv(ENGINE / "ItemSparse.csv")) == {111: 11}


def test_every_kind_is_emitted_once_and_in_the_contracts_order():
    document, _ = built()
    assert [s.id for s in document.sources] == [
        "raid:molten-core",
        "dungeon:the-deadmines",
        "world:azuregos",
        "crafted:blacksmithing",
        "rep:argent-dawn:exalted",
        "pvp:rank-11",
        "quest",
    ]
    assert [s.kind for s in document.sources] == list(KIND_ORDER)


def test_a_raid_lists_a_boss_per_npc_and_everything_else_as_trash():
    raid = source("raid:molten-core")
    assert raid.kind == "raid"
    assert raid.name == "Molten Core"
    assert raid.zone_id == 2717
    assert raid.opens is None
    assert raid.trash == [101]
    assert [(b.id, b.name, b.npc_id, b.items) for b in raid.bosses] == [
        ("raid:molten-core:900", "Big Boss", 900, [100, 112]),
        ("raid:molten-core:901", "", 901, [102]),
    ]


def test_a_dungeon_reached_through_the_fallback_join_is_still_a_dungeon():
    dungeon = source("dungeon:the-deadmines")
    assert dungeon.kind == "dungeon"
    assert [b.items for b in dungeon.bosses] == [[103]]
    assert dungeon.trash is None


def test_only_a_named_npc_outside_an_instance_becomes_a_world_boss():
    world = source("world:azuregos")
    assert world.kind == "world"
    assert world.name == "Azuregos"
    assert world.items == [104]
    document, _ = built()
    assert not [s for s in document.sources if s.id.startswith("world:") and s.id != "world:azuregos"]


def test_crafted_rep_pvp_and_quest_carry_the_keys_their_kind_needs():
    crafted = source("crafted:blacksmithing")
    assert (crafted.profession, crafted.items) == ("blacksmithing", [106])
    rep = source("rep:argent-dawn:exalted")
    assert (rep.faction_id, rep.standing, rep.items) == (529, "exalted", [107])
    assert rep.name == "Argent Dawn"
    pvp = source("pvp:rank-11")
    assert (pvp.rank, pvp.items) == (11, [111])
    assert source("quest").items == [108, 112]


def test_an_item_with_two_sources_appears_under_both():
    assert 112 in source("raid:molten-core").bosses[0].items
    assert 112 in source("quest").items


def test_an_item_the_build_does_not_have_is_left_out_with_its_boss():
    """Contract 10.4: loot.json lists only items the build has. Item 113
    is a raid drop the fork knows and this client's item table does not,
    so both it and the boss it was the only drop of are gone."""
    raid = source("raid:molten-core")
    assert 113 not in source_item_ids(raid)
    assert [boss.npc_id for boss in raid.bosses] == [900, 901]


def test_the_stats_count_what_was_emitted_dropped_and_absent():
    _, stats = built()
    # 100, 101, 102, 103, 104, 106, 107, 108, 111, 112
    assert stats.items == 10
    # the vendor-only item and the unnamed open-world mob's drop
    assert stats.dropped_entries == 2
    # item 113, which the build's item table does not have
    assert stats.absent_items == 1
```

- [ ] **Step 3: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_loot_sources.py -q --no-cov`
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.loot.sources'`

- [ ] **Step 4: Add the models and the document writer**

In `data/pipeline/models.py`, after `SuffixRecord`:

```python
class LootBoss(BaseModel):
    id: str
    #: Empty where the fork database names no NPC for the id. 35 of the 67
    #: raid bosses and 39 of the 230 dungeon bosses on the pinned fork are
    #: in that state; an invented name would be worse than a blank one.
    name: str
    npc_id: int
    items: list[int]


class LootSource(BaseModel):
    """One place loot comes from (parity contract 6.1).

    Every key after `name` is optional and omitted when it does not apply
    to the kind, which is what `write_document`'s `exclude_none` is for: a
    crafted source carries `profession` and `items`, a raid carries
    `zone_id`, `bosses` and `trash`.
    """

    id: str
    kind: str
    name: str
    zone_id: int | None = None
    #: A phase name from api/internal/phase. None means open from launch.
    #: Never set by the generator -- only by a curated overlay, because the
    #: databases state no dates.
    opens: str | None = None
    profession: str | None = None
    faction_id: int | None = None
    standing: str | None = None
    rank: int | None = None
    bosses: list[LootBoss] | None = None
    trash: list[int] | None = None
    items: list[int] | None = None


class LootFile(BaseModel):
    sources: list[LootSource]
```

In `data/pipeline/normalize/__init__.py`, beside `write_model`:

```python
def write_document(record: BaseModel, path: Path) -> None:
    """Write a single record as a JSON object, dropping the keys it leaves unset.

    `write_model` emits every field, which is right for a record whose shape
    is fixed. `loot.json`'s sources are a union -- a crafted source has no
    `bosses` and a raid has no `profession` -- and a file of nulls is a file
    every consumer has to filter, so `None` means "not part of this kind"
    and is left out.
    """
    _write(record.model_dump(exclude_none=True), path)
```

- [ ] **Step 5: Write the builder**

Create `data/pipeline/loot/sources.py`:

```python
"""`loot.json`: every place the two databases say an item comes from.

Parity contract 6.1. Droptimizer's source picker and Top Gear's "pin this
drop" both read this file, so the ids here are stable keys: a candidate's
`origin` is `drop:<source id>`.

Nothing is invented. An item the fork database gives no source and the
client gives no PvP rank is simply absent; a boss the fork does not name
is emitted with an empty name; a drop source whose kind has no home in the
contract's vocabulary -- a vendor with no rank, an unnamed open-world mob --
is dropped and counted, never guessed into a kind.
"""

from __future__ import annotations

import logging
from collections import defaultdict
from dataclasses import dataclass

from pipeline.csvio import populated
from pipeline.forkdb import PROFESSIONS, REP_LEVELS, ForkDatabase, decode
from pipeline.models import LootBoss, LootFile, LootSource
from pipeline.normalize.classes import slugify

logger = logging.getLogger(__name__)

#: The order sources are emitted in, which is the order the picker shows
#: them: instances first, then the things you buy or make.
KIND_ORDER = ("raid", "dungeon", "world", "crafted", "rep", "pvp", "quest")

#: `Map.InstanceType`. 3 (battleground) and 4 (arena) are instances whose
#: loot the contract has no kind for -- a battleground's rewards are
#: reputation and rank, which are their own kinds -- so only these two
#: become drop sources.
INSTANCE_KIND = {1: "dungeon", 2: "raid"}


@dataclass(frozen=True)
class LootStats:
    """What one build's loot.json covers, for the log and the test."""

    #: Distinct item ids the emitted file names.
    items: int
    #: Fork source entries whose kind has no home in contract 6.1.
    dropped_entries: int
    #: Distinct item ids the fork sources name that this build's item table
    #: does not have. Contract 10.4 leaves them out; the count is the size
    #: of the re-itemisation gap and is logged and pinned by a test.
    absent_items: int


def instance_types(
    map_rows: list[dict[str, str]], zone_rows: list[dict]
) -> dict[int, int]:
    """AreaTable zone id -> `Map.InstanceType`, for the zones inside one.

    Two joins, in this order, because neither covers every instance on
    build 1.60.1.69893:

    * `Map.AreaTableID` names the area a map's entrance is in, and resolves
      19 of the fork's 27 instance zones -- Molten Core, Blackwing Lair,
      Onyxia's Lair, Zul'Gurub among them.
    * `AreaTable.ContinentID`, which `zones.json` already carries as
      `map_id`, resolves the other 8 (Deadmines, Stratholme, Scholomance,
      Scarlet Monastery, Zul'Farrak, both Razorfens, Shadowfang Keep). It
      is the fallback and not the first join because it is wrong for
      Onyxia's Lair, whose AreaTable row says Kalimdor.

    The second join is also what finds Ahn'Qiraj, its Ruins and Naxxramas,
    whose drops the fork database references but whose zones its own
    `zones` list omits.
    """
    by_map = {int(row["ID"]): int(row["InstanceType"] or 0) for row in map_rows}
    by_area: dict[int, int] = {}
    for row in map_rows:
        area = int(populated(row, "AreaTableID") or 0)
        kind = int(row["InstanceType"] or 0)
        if area and kind:
            by_area.setdefault(area, kind)
    types: dict[int, int] = {}
    for zone in zone_rows:
        zone_id = int(zone["id"])
        kind = by_area.get(zone_id) or by_map.get(int(zone["map_id"]), 0)
        if kind:
            types[zone_id] = kind
    return types


def pvp_ranks(sparse_rows: list[dict[str, str]]) -> dict[int, int]:
    """Item id -> the PvP rank it requires, for the items that require one.

    The fork database has no rank data at all; the client states it, and
    266 of the fork's 273 vendor-only items are rank sets, so this is what
    makes a `pvp` source possible without inventing anything.
    """
    ranks: dict[int, int] = {}
    for row in sparse_rows:
        rank = int(populated(row, "RequiredPVPRank") or 0)
        if rank:
            ranks[int(row["ID"])] = rank
    return ranks
```

- [ ] **Step 6: Write the five source builders and the assembly**

Append to `data/pipeline/loot/sources.py`:

```python
def _drop_sources(
    fork: ForkDatabase,
    zone_names: dict[int, str],
    types: dict[int, int],
    build_items: set[int],
    absent: set[int],
) -> tuple[list[LootSource], int]:
    """The raid, dungeon and world sources, and how many drops had no home.

    `absent` collects, in place, every item id this build's table does not
    have, so the caller can count the re-itemisation gap once across all
    the builders rather than three times.
    """
    bosses: dict[tuple[int, int], set[int]] = defaultdict(set)
    trash: dict[int, set[int]] = defaultdict(set)
    world: dict[int, set[int]] = defaultdict(set)
    dropped = 0
    for item in fork.items:
        item_id = int(item["id"])
        for source in item.get("sources") or []:
            drop = source.get("drop")
            if drop is None:
                continue
            zone_id, npc_id = int(drop.get("zoneId", 0)), int(drop.get("npcId", 0))
            kind = INSTANCE_KIND.get(types.get(zone_id, 0))
            if kind is None and npc_id not in fork.npcs:
                # An open-world mob the fork does not name, or a drop with no
                # zone at all. The contract has no kind for either.
                dropped += 1
                continue
            if item_id not in build_items:
                # Contract 10.4: the build has no such item, so nothing could
                # render or sim it.
                absent.add(item_id)
                continue
            if kind is not None:
                (bosses[(zone_id, npc_id)] if npc_id else trash[zone_id]).add(item_id)
            else:
                world[npc_id].add(item_id)

    out: list[LootSource] = []
    for zone_id in sorted({zone for zone, _ in bosses} | set(trash), key=lambda z: (
        INSTANCE_KIND[types[z]], slugify(zone_names[z])
    )):
        kind = INSTANCE_KIND[types[zone_id]]
        slug = f"{kind}:{slugify(zone_names[zone_id])}"
        in_zone = sorted(npc for zone, npc in bosses if zone == zone_id)
        out.append(
            LootSource(
                id=slug,
                kind=kind,
                name=zone_names[zone_id],
                zone_id=zone_id,
                bosses=[
                    LootBoss(
                        id=f"{slug}:{npc_id}",
                        name=fork.npcs.get(npc_id, ""),
                        npc_id=npc_id,
                        items=sorted(bosses[(zone_id, npc_id)]),
                    )
                    for npc_id in in_zone
                ]
                or None,
                trash=sorted(trash[zone_id]) or None,
            )
        )
    out.extend(
        LootSource(
            id=f"world:{slugify(fork.npcs[npc_id])}",
            kind="world",
            name=fork.npcs[npc_id],
            items=sorted(items),
        )
        for npc_id, items in sorted(world.items(), key=lambda pair: slugify(fork.npcs[pair[0]]))
    )
    return out, dropped


def _keyed_sources(
    fork: ForkDatabase, build_items: set[int], absent: set[int]
) -> tuple[list[LootSource], list[int], int]:
    """The crafted, rep and quest sources, plus how many entries had no home."""
    crafted: dict[str, set[int]] = defaultdict(set)
    rep: dict[tuple[int, str], set[int]] = defaultdict(set)
    quest: set[int] = set()
    dropped = 0
    for item in fork.items:
        item_id = int(item["id"])
        for source in item.get("sources") or []:
            if not ({"crafted", "rep", "quest"} & set(source)):
                if "soldBy" in source:
                    # A vendor. Rank vendors are covered by the pvp kind,
                    # read off the client; anything else has no kind in 6.1.
                    dropped += 1
                continue
            if item_id not in build_items:
                absent.add(item_id)
                continue
            if "crafted" in source:
                profession = decode(
                    PROFESSIONS, int(source["crafted"]["profession"]), "profession"
                )
                crafted[profession].add(item_id)
            elif "rep" in source:
                standing = decode(REP_LEVELS, int(source["rep"]["repLevel"]), "rep level")
                faction_id = int(source["rep"]["repFactionId"])
                if faction_id not in fork.factions:
                    # A faction the fork's own table does not name: there is
                    # nothing to call the source, so it is not emitted.
                    dropped += 1
                    continue
                rep[(faction_id, standing)].add(item_id)
            elif "quest" in source:
                quest.add(item_id)
    out = [
        LootSource(
            id=f"crafted:{profession}",
            kind="crafted",
            name=profession.replace("-", " ").title(),
            profession=profession,
            items=sorted(items),
        )
        for profession, items in sorted(crafted.items())
    ]
    out += [
        LootSource(
            id=f"rep:{slugify(fork.factions[faction_id])}:{standing}",
            kind="rep",
            name=fork.factions[faction_id],
            faction_id=faction_id,
            standing=standing,
            items=sorted(items),
        )
        for (faction_id, standing), items in sorted(
            rep.items(), key=lambda pair: (slugify(fork.factions[pair[0][0]]), pair[0][1])
        )
    ]
    return out, sorted(quest), dropped


def _pvp_sources(ranks: dict[int, int], build_items: set[int]) -> list[LootSource]:
    """One source per PvP rank. `ranks` is read off the client's own
    `ItemSparse`, so its ids are the build's by construction; the filter
    is applied anyway so one rule governs every kind."""
    by_rank: dict[int, set[int]] = defaultdict(set)
    for item_id, rank in ranks.items():
        if item_id in build_items:
            by_rank[rank].add(item_id)
    return [
        LootSource(
            id=f"pvp:rank-{rank}",
            kind="pvp",
            name=f"Rank {rank}",
            rank=rank,
            items=sorted(items),
        )
        for rank, items in sorted(by_rank.items())
    ]


def source_item_ids(source: LootSource) -> set[int]:
    """Every item id one source names, wherever it names it."""
    return set(
        (source.items or [])
        + (source.trash or [])
        + [item for boss in (source.bosses or []) for item in boss.items]
    )


def build_loot(
    fork: ForkDatabase,
    zone_names: dict[int, str],
    types: dict[int, int],
    ranks: dict[int, int],
    build_items: set[int],
) -> tuple[LootFile, LootStats]:
    absent: set[int] = set()
    drops, dropped_drops = _drop_sources(fork, zone_names, types, build_items, absent)
    keyed, quest, dropped_keyed = _keyed_sources(fork, build_items, absent)
    sources = [*drops, *keyed, *_pvp_sources(ranks, build_items)]
    if quest:
        sources.append(LootSource(id="quest", kind="quest", name="Quests", items=quest))
    # A source the filter emptied is not a source. `_drop_sources` already
    # drops a boss with no items left (its `bosses` set is simply never
    # created), so this is the last sweep: a zone whose every drop was
    # re-itemised away, like Onyxia's Lair on build 1.60.1.69893.
    sources = [source for source in sources if source_item_ids(source)]
    sources.sort(key=lambda source: (KIND_ORDER.index(source.kind), source.id))
    named = {item_id for source in sources for item_id in source_item_ids(source)}
    return LootFile(sources=sources), LootStats(
        items=len(named),
        dropped_entries=dropped_drops + dropped_keyed,
        absent_items=len(absent),
    )
```

- [ ] **Step 7: Run the tests**

Run: `cd data && uv run pytest tests/test_loot_sources.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 8: Commit**

```bash
git add data/pipeline/models.py data/pipeline/normalize/__init__.py data/pipeline/loot/sources.py data/tests/test_loot_sources.py data/tests/fixtures/loot/Map.csv data/tests/fixtures/loot/ItemSparse.csv data/tests/fixtures/loot/zones.json
git commit -m "$(cat <<'EOF'
feat(data): loot.json, every source the two databases state

Contract 6.1. Raids and dungeons per boss with their trash, world bosses,
crafted by profession, reputation by faction and standing, PvP by rank,
quests.

The zone join is two steps because neither covers every instance:
Map.AreaTableID first (19 of 27, and the only one that gets Onyxia's Lair
right), the build's own zones.json map_id as the fallback (the other 8,
plus the three Ahn'Qiraj and Naxxramas zones the fork's zone list omits).

PvP comes from the client's RequiredPVPRank because the fork has no rank
data. A vendor with no rank and an open-world mob the fork does not name
have no kind in the contract, so they are dropped and counted rather
than guessed into one.

Contract 10.4's build filter: an id the fork names that this client's
item table does not have could never render or sim, so it is left out
and counted. 1,809 of them, which is Forever's re-itemisation showing
through, and which is what takes 78 sources down to 62.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 7: The curated loot overlay

Contract 10.4 fixes the shape: `{ "sources": [...provenance...], "notes": "...", "add": [loot sources], "replace": [loot sources], "remove": [ids] }`. `sources` and `notes` keep their curated meaning — the provenance every hand-maintained fact here carries — and the three operations are flat arrays.

```json
{
  "sources": [{ "label": "...", "url": "https://...", "kind": "blizzard" }],
  "notes": "why this file exists",
  "add":     [ { "id": "raid:barrow-deeps", "kind": "raid", "name": "Barrow Deeps", "opens": "raids-1" } ],
  "replace": [ { "id": "raid:molten-core", "opens": "later" } ],
  "remove":  [ "dungeon:x" ]
}
```

- `add` — whole new sources, each a `LootSource`. An id that already exists is an error; use `replace`.
- `replace` — loot sources too, but with every key except `id` optional: **only the keys the file actually writes are overwritten**, so gating a raid is one line rather than a restatement of its eleven bosses. The id must exist. `kind` may be written but must match the kind the source already has — a source's kind decides the shape of its id, which candidates carry as `drop:<source id>`, so changing it would orphan every reference rather than edit one.
- `remove` — source ids to drop whole. Each must exist.

Contract 10.4 has no per-item removal, so there is none: a wrong item is corrected by `replace`-ing the source's `items`, `trash` or `bosses`.

Files apply in filename order; within a file, `add`, then `replace`, then `remove`. Every mismatch is a hard error, the same policy `pipeline/curated.py` has for an unsourced claim: an overlay that has quietly stopped doing anything is worse than one that fails.

**Files:**
- Modify: `data/pipeline/curated.py` (`_sources` becomes the public `parse_sources`; three call sites)
- Modify: `data/pipeline/models.py` (`LootSourcePatch`, `LootOverlay`)
- Create: `data/pipeline/loot/overlay.py`
- Create: `data/tests/fixtures/loot/curated/10-add.json`, `data/tests/fixtures/loot/curated/20-replace-and-remove.json`
- Test: `data/tests/test_loot_overlay.py`

**Interfaces:**
- Consumes: `pipeline.curated.SOURCE_KINDS`, `pipeline.curated.parse_sources`; `pipeline.models.LootFile`, `LootSource`
- Produces:
  - `pipeline.curated.parse_sources(raw: list[dict], where: str) -> list[Source]` (renamed from `_sources`)
  - `pipeline.models.LootSourcePatch` — a `LootSource` with every key but `id` optional
  - `pipeline.models.LootOverlay` with `sources: list[Source]`, `notes: str`, `add: list[LootSource]`, `replace: list[LootSourcePatch]`, `remove: list[str]`
  - `pipeline.loot.overlay.OverlayError(SystemExit)`
  - `pipeline.loot.overlay.load_overlays(overlay_dir: Path) -> list[tuple[Path, LootOverlay]]`
  - `pipeline.loot.overlay.apply_overlays(document: LootFile, overlays: list[tuple[Path, LootOverlay]]) -> LootFile`

- [ ] **Step 1: Write the fixture overlays**

Create `data/tests/fixtures/loot/curated/10-add.json`:

```json
{
  "sources": [
    { "label": "A test source", "url": "https://example.invalid/one", "kind": "site" }
  ],
  "notes": "Adds a raid the databases have never heard of.",
  "add": [
    {
      "id": "raid:barrow-deeps",
      "kind": "raid",
      "name": "Barrow Deeps",
      "opens": "raids-1",
      "items": [100]
    }
  ]
}
```

Create `data/tests/fixtures/loot/curated/20-replace-and-remove.json`:

```json
{
  "sources": [
    { "label": "A test source", "url": "https://example.invalid/two", "kind": "site" }
  ],
  "notes": "Gates the raid and drops a dungeon.",
  "replace": [{ "id": "raid:molten-core", "opens": "later" }],
  "remove": ["dungeon:the-deadmines"]
}
```

- [ ] **Step 2: Write the failing test**

Create `data/tests/test_loot_overlay.py`:

```python
import json
from pathlib import Path

import pytest

from pipeline.curated import CuratedError
from pipeline.loot.overlay import OverlayError, apply_overlays, load_overlays
from pipeline.models import LootBoss, LootFile, LootSource

HERE = Path(__file__).parent
FIXTURE = HERE / "fixtures/loot/curated"
CURATED = Path("curated/loot")


def base() -> LootFile:
    return LootFile(
        sources=[
            LootSource(
                id="raid:molten-core",
                kind="raid",
                name="Molten Core",
                zone_id=2717,
                bosses=[LootBoss(id="raid:molten-core:900", name="Big Boss", npc_id=900,
                                 items=[100])],
                trash=[101],
            ),
            LootSource(
                id="dungeon:the-deadmines",
                kind="dungeon",
                name="The Deadmines",
                zone_id=1581,
                bosses=[LootBoss(id="dungeon:the-deadmines:902", name="Dungeon Boss",
                                 npc_id=902, items=[103])],
            ),
        ]
    )


def write(tmp_path: Path, name: str, payload: dict) -> Path:
    (tmp_path / name).write_text(json.dumps(payload), encoding="utf-8")
    return tmp_path


def test_a_missing_overlay_directory_is_not_an_error(tmp_path):
    assert load_overlays(tmp_path / "nope") == []


def test_the_fixture_overlays_add_patch_and_remove():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    by_id = {source.id: source for source in result.sources}
    assert set(by_id) == {"raid:molten-core", "raid:barrow-deeps"}
    assert by_id["raid:molten-core"].opens == "later"
    assert by_id["raid:barrow-deeps"].items == [100]


def test_a_replace_only_touches_the_keys_it_names():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert molten.name == "Molten Core"
    assert molten.zone_id == 2717
    assert molten.trash == [101]
    assert [b.items for b in molten.bosses] == [[100]]


def test_the_result_stays_sorted_by_kind_then_id():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    assert [s.id for s in result.sources] == ["raid:barrow-deeps", "raid:molten-core"]


def test_an_overlay_with_no_source_is_refused(tmp_path):
    write(tmp_path, "a.json", {"notes": "x", "add": []})
    with pytest.raises(CuratedError):
        load_overlays(tmp_path)


def test_an_overlay_source_kind_outside_the_vocabulary_is_refused(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "u", "kind": "hearsay"}], "notes": "x"})
    with pytest.raises(CuratedError, match="hearsay"):
        load_overlays(tmp_path)


def test_adding_a_source_that_already_exists_is_refused(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "add": [{"id": "raid:molten-core", "kind": "raid", "name": "Again"}],
    })
    with pytest.raises(OverlayError, match="already"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_replacing_or_removing_an_unknown_source_is_refused(tmp_path):
    for payload in (
        {"replace": [{"id": "raid:nope", "opens": "later"}]},
        {"remove": ["raid:nope"]},
    ):
        write(tmp_path, "a.json", {
            "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
            "notes": "x",
        } | payload)
        with pytest.raises(OverlayError, match="raid:nope"):
            apply_overlays(base(), load_overlays(tmp_path))


def test_a_replace_may_restate_a_sources_item_lists(tmp_path):
    # Contract 10.4 has no per-item removal, so this is how a wrong item
    # is corrected: restate the list it is in.
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "replace": [{"id": "raid:molten-core", "trash": []}],
    })
    result = apply_overlays(base(), load_overlays(tmp_path))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert molten.trash == []
    assert [b.items for b in molten.bosses] == [[100]]


def test_a_replace_may_not_change_a_sources_kind(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "replace": [{"id": "raid:molten-core", "kind": "dungeon"}],
    })
    with pytest.raises(OverlayError, match="kind"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_the_real_curated_overlays_all_load_and_apply_cleanly():
    """The committed overlays must be loadable on their own terms; that
    they match the committed loot.json is tests/test_loot_build.py's job."""
    for path, document in load_overlays(CURATED):
        assert document.sources, path
        assert document.notes.strip(), path
```

- [ ] **Step 3: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_loot_overlay.py -q --no-cov`
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.loot.overlay'`

- [ ] **Step 4: Make the source parser public**

In `data/pipeline/curated.py`, rename `_sources` to `parse_sources`, give it the docstring below, and update its three call sites (`_changes`, `_merge_races` via `_changes`, and `_build_combos`) — a grep for `_sources(` finds them all:

```python
def parse_sources(raw: list[dict], where: str) -> list[Source]:
    """The provenance every curated fact carries, validated.

    Public because `pipeline/loot/overlay.py` states the same rule about
    the same vocabulary, and two copies of "at least one source, from this
    list of kinds, with a non-empty label and url" is two things to keep
    right.
    """
```

- [ ] **Step 5: Add the overlay models**

In `data/pipeline/models.py`, after `LootFile`:

```python
class LootSourcePatch(BaseModel):
    """Contract 10.4's `replace` entry: a loot source with every key but
    `id` optional, so gating a raid is one line instead of a restatement
    of its eleven bosses. Only the keys the file actually writes are
    applied -- `model_dump(exclude_unset=True)` is what tells "absent"
    from "explicitly null".

    `kind` is accepted but may not change: a source's kind decides the
    shape of its id, which candidates carry as `drop:<source id>`, so
    changing it would orphan every reference rather than edit one.
    """

    id: str
    kind: str | None = None
    name: str | None = None
    zone_id: int | None = None
    opens: str | None = None
    profession: str | None = None
    faction_id: int | None = None
    standing: str | None = None
    rank: int | None = None
    bosses: list[LootBoss] | None = None
    trash: list[int] | None = None
    items: list[int] | None = None

    model_config = {"extra": "forbid"}


class LootOverlay(BaseModel):
    """One `data/curated/loot/*.json`, in contract 10.4's shape.

    `sources` and `notes` are the curated provenance every hand-maintained
    fact here carries; `add`, `replace` and `remove` are the three
    operations, as flat arrays.
    """

    sources: list[Source] = []
    notes: str = ""
    add: list[LootSource] = []
    replace: list[LootSourcePatch] = []
    remove: list[str] = []

    model_config = {"extra": "forbid"}
```

- [ ] **Step 6: Write the overlay loader and applier**

Create `data/pipeline/loot/overlay.py`:

```python
"""Forever's own loot facts, laid over the two databases' Era ones.

Parity contract 6.1 as corrected by 10.4, and the design's risk list: the
fork's item database covers Classic Era, and Forever re-itemises. Only
2,811 of the fork's 7,553 item ids exist in the 1.60 client at all, so a
great deal of what a Forever player will actually loot has no source in
either database. This is where a sourced statement about that goes -- and
where the phase a raid opens in goes, since neither database states a date.

Every file is a `LootOverlay`: curated `sources` and `notes` like every
other hand-maintained fact here, plus `add`, `replace` and `remove`.
Files apply in filename order; within a file, add, then replace, then
remove. A stale instruction -- adding a source that exists, patching or
removing one that does not -- is an error, because the alternative is an
overlay that quietly stops doing anything.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.curated import parse_sources
from pipeline.loot.sources import KIND_ORDER
from pipeline.models import LootFile, LootOverlay


class OverlayError(SystemExit):
    """A curated loot overlay names something the generated file does not."""


def load_overlays(overlay_dir: Path) -> list[tuple[Path, LootOverlay]]:
    """Every overlay in filename order, with its provenance validated.

    A missing directory is not an error: a build with nothing to say about
    its loot is a normal state, and the day one exists the file appears.
    """
    if not overlay_dir.is_dir():
        return []
    loaded: list[tuple[Path, LootOverlay]] = []
    for path in sorted(overlay_dir.glob("*.json")):
        document = LootOverlay(**json.loads(path.read_text(encoding="utf-8")))
        parse_sources([source.model_dump() for source in document.sources], str(path))
        if not document.notes.strip():
            raise OverlayError(f"{path} has no notes; say why the overlay exists")
        loaded.append((path, document))
    return loaded


def apply_overlays(
    document: LootFile, overlays: list[tuple[Path, LootOverlay]]
) -> LootFile:
    by_id = {source.id: source for source in document.sources}
    for path, overlay in overlays:
        for source in overlay.add:
            if source.id in by_id:
                raise OverlayError(
                    f"{path} adds source {source.id!r}, which the generated file "
                    f"already has; use `replace` to change it"
                )
            by_id[source.id] = source
        for patch in overlay.replace:
            if patch.id not in by_id:
                raise OverlayError(
                    f"{path} replaces keys on source {patch.id!r}, which no source "
                    f"has; use `add`, or delete the stale instruction"
                )
            existing = by_id[patch.id]
            changes = patch.model_dump(exclude_unset=True)
            changes.pop("id")
            kind = changes.pop("kind", existing.kind)
            if kind != existing.kind:
                raise OverlayError(
                    f"{path} would change source {patch.id!r} from kind "
                    f"{existing.kind!r} to {kind!r}; a source's kind is part of "
                    f"its id and of every `drop:` origin that names it"
                )
            by_id[patch.id] = existing.model_copy(update=changes)
        for source_id in overlay.remove:
            if source_id not in by_id:
                raise OverlayError(
                    f"{path} removes source {source_id!r}, which is not there"
                )
            del by_id[source_id]
    return LootFile(
        sources=sorted(
            by_id.values(), key=lambda source: (KIND_ORDER.index(source.kind), source.id)
        )
    )
```

- [ ] **Step 7: Create the curated directory**

The directory has to exist for `load_overlays(Path("curated/loot"))` to find task 8's file, and git does not track empty directories — task 8 creates the first file, so nothing is needed here beyond leaving `CURATED` in the test pointing at it (that test iterates an empty list until task 8 lands).

- [ ] **Step 8: Run the tests**

Run: `cd data && uv run pytest tests/test_loot_overlay.py tests/test_curated.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 9: Commit**

```bash
git add data/pipeline/curated.py data/pipeline/models.py data/pipeline/loot/overlay.py data/tests/test_loot_overlay.py data/tests/fixtures/loot/curated/
git commit -m "$(cat <<'EOF'
feat(data): the curated loot overlay

Contract 10.4's overlay: provenance keeps the top-level `sources` and
`notes`, and the three operations are flat arrays beside them.

Three operations, in contract 10.4's shape: add a whole source, replace
named keys on an existing one (which is how a raid gets its `opens`),
remove a source. A stale instruction is an error, because the
alternative is an overlay that quietly stops doing anything.

curated.py's `_sources` becomes the public `parse_sources` so both
validators state the provenance rule once.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 8: The first overlay — Forever's raid phases and the loot gap

Two facts neither database states, both required by the design's section 6.1 ("the phase each raid opens"), its risk list ("Forever-only items with no source"), and contract 10.4's last bullet ("the first overlay records the gap per raid in its notes").

**The phases.** `api/internal/phase/phase.go` has four names and exactly one raid phase, `raids-1` (9 December 2026). The site's own `web/src/data/dates.json` names that day's raids: **Barrow Deeps, Hyjal Summit and Onyxia**. So exactly one of the generated raid sources has a known date — and it is the one the build filter removed, because all sixteen of Onyxia's fork-database drops are gone from the 1.60 client. The overlay therefore:

- **adds** `raid:onyxias-lair` with `opens: "raids-1"` and no items: an announced raid whose loot table neither database knows. The zone is real (AreaTable 2159) and the date is sourced; the item list is honestly empty until someone curates it.
- **replaces** `opens: "later"` on the six Era raids the generator does emit. Contract 10.4 gives `later` exactly this meaning — "unreleased, no date" — and it is the honest answer: Blizzard has announced no date for Molten Core, Blackwing Lair, Zul'Gurub, either Ahn'Qiraj or Naxxramas, and leaving `opens` unset would read as "open from launch", which is wrong for all six.

**The gap, per raid.** Measured on the pinned fork against build 1.60.1.69893, distinct items kept of items the fork's sources name:

| raid | kept | fork names | lost |
| --- | --- | --- | --- |
| Ahn'Qiraj | 68 | 111 | 43 |
| Blackwing Lair | 21 | 132 | 111 |
| Molten Core | 10 | 139 | 129 |
| Naxxramas | 45 | 111 | 66 |
| Onyxia's Lair | 0 | 16 | 16 |
| Ruins of Ahn'Qiraj | 5 | 61 | 56 |
| Zul'Gurub | 2 | 98 | 96 |

That is the whole of contract 10.4's "raid Droptimizer is honest and thin until sourced replacements are curated", and the notes carry it so the next person does not re-derive it.

Barrow Deeps and Hyjal Summit are **not** added. The client has candidate maps for both (2817 "Starfall Barrow Den", 2832 "Nightmare Grove", 2995 "Hyjal Crater"), but nothing sourced maps the announced names onto those ids, so the notes record the open question instead of answering it.

**Files:**
- Create: `data/curated/loot/forever-raid-phases.json`
- Test: `data/tests/test_loot_overlay.py` (extend)

**Interfaces:**
- Consumes: `pipeline.loot.overlay.load_overlays`
- Produces: the committed overlay, whose `replace` ids must match the six raid source ids task 6 generates for this build, and whose one `add` id must not.

- [ ] **Step 1: Write the failing test**

Append to `data/tests/test_loot_overlay.py`:

```python
#: The raid sources the generator emits for build 1.60.1.69893 after
#: contract 10.4's build filter. Onyxia's Lair is absent -- all sixteen
#: of her fork-database drops are gone from the client -- which is why
#: the overlay adds her rather than patching her.
GENERATED_RAID_IDS = [
    "raid:ahnqiraj",
    "raid:blackwing-lair",
    "raid:molten-core",
    "raid:naxxramas",
    "raid:ruins-of-ahnqiraj",
    "raid:zulgurub",
]
#: api/internal/phase/phase.go's names, plus contract 10.4's sentinel.
PHASES = {"pre-beta", "beta", "launch", "raids-1"}
OPENS_LATER = "later"


def raid_phase_overlay():
    for path, document in load_overlays(CURATED):
        if path.name == "forever-raid-phases.json":
            return document
    raise AssertionError("no forever-raid-phases.json under curated/loot")


def test_the_six_generated_raids_are_gated_as_unreleased_without_a_date():
    document = raid_phase_overlay()
    assert sorted(patch.id for patch in document.replace) == GENERATED_RAID_IDS
    assert {patch.opens for patch in document.replace} == {OPENS_LATER}


def test_onyxia_is_added_with_the_announced_date_and_no_items():
    """The one raid the phase calendar has a date for, and the one the
    build filter removed entirely. An empty item list is the honest
    statement: the raid is announced, its loot table is not known."""
    document = raid_phase_overlay()
    assert [source.id for source in document.add] == ["raid:onyxias-lair"]
    onyxia = document.add[0]
    assert (onyxia.kind, onyxia.zone_id, onyxia.opens) == ("raid", 2159, "raids-1")
    assert onyxia.items == []


def test_the_raid_phase_overlay_removes_nothing():
    assert raid_phase_overlay().remove == []


def test_the_notes_state_the_per_raid_gap():
    """Contract 10.4: 'the first overlay records the gap per raid in its
    notes'. Spot-check the two extremes rather than the whole table."""
    notes = raid_phase_overlay().notes
    for fragment in ("Molten Core", "10 of 139", "Zul'Gurub", "2 of 98"):
        assert fragment in notes


def test_every_phase_an_overlay_names_is_a_real_phase_or_the_sentinel():
    """An `opens` the API does not know is a filter that matches nothing."""
    allowed = PHASES | {OPENS_LATER}
    for path, document in load_overlays(CURATED):
        for patch in document.replace:
            assert patch.opens is None or patch.opens in allowed, (path, patch.id)
        for source in document.add:
            assert source.opens is None or source.opens in allowed, (path, source.id)
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_loot_overlay.py -q --no-cov`
Expected: FAIL — `AssertionError: no forever-raid-phases.json under curated/loot`

- [ ] **Step 3: Write the overlay**

Create `data/curated/loot/forever-raid-phases.json`:

```json
{
  "sources": [
    {
      "label": "Blizzard, Deep Dive panel recap",
      "url": "https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap",
      "kind": "blizzard"
    },
    {
      "label": "Blizzard, What's Next panel recap",
      "url": "https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap",
      "kind": "blizzard"
    }
  ],
  "notes": "When each raid opens, and what its loot table is missing. Nothing raids at launch on 4 November; the first tier opens on 9 December and is Barrow Deeps, Hyjal Summit and Onyxia. Of those three only Onyxia's Lair exists as a zone the fork database names, and the build filter removes it outright because all sixteen of its fork-database drops are gone from the 1.60 client, so it is added back here with its announced phase and an empty item list: the raid is real and dated, its Forever loot table is not known. Barrow Deeps and Hyjal Summit are not added at all -- the client has candidate maps (2817 Starfall Barrow Den, 2832 Nightmare Grove, 2995 Hyjal Crater) but nothing sourced maps the announced names onto those ids, and neither database gives them an item. The six Era raids the generator does emit have no announced date, so they carry opens 'later' rather than nothing, which would read as open from launch. The re-itemisation gap per raid, measured on the pinned fork against build 1.60.1.69893 as items kept of items the fork's sources name: Ahn'Qiraj 68 of 111, Blackwing Lair 21 of 132, Molten Core 10 of 139, Naxxramas 45 of 111, Onyxia's Lair 0 of 16, Ruins of Ahn'Qiraj 5 of 61, Zul'Gurub 2 of 98. Overall 1,809 of the 4,316 ids the fork's sources name are absent from this client's item table: Forever has re-itemised the raid tier -- Bonereaver's Edge 17076, Band of Accuria 17063 and every tier-1 and tier-2 set piece are gone -- and states no replacement's source anywhere, so raid Droptimizer stays honest and thin until replacements are curated here.",
  "add": [
    {
      "id": "raid:onyxias-lair",
      "kind": "raid",
      "name": "Onyxia's Lair",
      "zone_id": 2159,
      "opens": "raids-1",
      "items": []
    }
  ],
  "replace": [
    { "id": "raid:ahnqiraj", "opens": "later" },
    { "id": "raid:blackwing-lair", "opens": "later" },
    { "id": "raid:molten-core", "opens": "later" },
    { "id": "raid:naxxramas", "opens": "later" },
    { "id": "raid:ruins-of-ahnqiraj", "opens": "later" },
    { "id": "raid:zulgurub", "opens": "later" }
  ]
}
```

Note the consequence for task 6's pruning sweep: `raid:onyxias-lair` carries an empty `items` list and would be dropped if the sweep ran after the overlay. It does not — `build_loot` prunes, and `apply_overlays` runs on its result — which is exactly why a curated empty source survives and a generated one does not. Task 11's orchestration must keep that order, and task 12 asserts the source is present.

- [ ] **Step 4: Run the tests**

Run: `cd data && uv run pytest tests/test_loot_overlay.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 5: Commit**

```bash
git add data/curated/loot/forever-raid-phases.json data/tests/test_loot_overlay.py
git commit -m "$(cat <<'EOF'
feat(data): Forever's raid phases, and the loot gap per raid

The first curated loot overlay. Onyxia's Lair is added back with its
announced 9 December phase and an empty item list -- the build filter
removes it because all sixteen of its fork-database drops are gone from
the 1.60 client -- and the six Era raids the generator still emits carry
opens "later", contract 10.4's value for an unreleased source with no
date. Leaving opens unset would read as open from launch, which is wrong
for all six.

The notes carry the per-raid gap contract 10.4 asks for: Molten Core
keeps 10 of the 139 items the fork's sources name for it, Zul'Gurub 2 of
98, and 1,809 of 4,316 ids overall are absent from this client's table.
Barrow Deeps and Hyjal Summit are recorded as an open question rather
than guessed onto client zone ids.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 9: `phases.json`, the phase calendar both lanes read

Contract 10.4: a new `data/curated/phases.json` of `[ { "name", "start" } ]`, "the source of truth `api/internal/phase.Boundaries` is tested against, emitted to `web/src/data/phases.json` by the pipeline". Contract 10.6 adds `GET /v1/phases`, which serves it.

Today the boundaries are stated twice — as a Go table in `api/internal/phase/phase.go` and as prose in `web/src/data/dates.json` — with a comment asking whoever changes one to change the other. This makes one of them data.

The four boundaries, transcribed from `phase.go`: `pre-beta` from the zero time, `beta` from midnight UTC on 17 September 2026, `launch` from 15:00 PST on 4 November 2026 (23:00 UTC), `raids-1` from midnight UTC on 9 December 2026. `later` is **not** a boundary: it is contract 10.4's sentinel for a source with no date, it would break `phase.At`, and it never appears in this file.

This is not a per-build emitter, so it follows `specs`: a `phases` command that writes, a `--check` that writes nothing and exits non-zero on drift, and CI's test job running the check. The Go test that holds `phase.Boundaries` to this file is the API lane's; this task produces the files it reads.

**Files:**
- Create: `data/curated/phases.json`
- Create: `data/pipeline/phases.py`
- Create: `web/src/data/phases.json` (generated, committed)
- Modify: `data/pipeline/models.py` (`PhaseBoundary`)
- Modify: `data/pipeline/__main__.py` (the `phases` subcommand)
- Test: `data/tests/test_phases.py`

**Interfaces:**
- Consumes: `pipeline.normalize.write_records`
- Produces:
  - `pipeline.models.PhaseBoundary` with `name: str`, `start: str`
  - `pipeline.phases.WEB_PATH: Path` = `Path("../web/src/data/phases.json")`
  - `pipeline.phases.PhaseError(SystemExit)`
  - `pipeline.phases.load_phases(curated_dir: Path = Path("curated")) -> list[PhaseBoundary]`
  - `pipeline.phases.write_phases(curated_dir: Path = Path("curated"), web_path: Path = WEB_PATH) -> Path`
  - `pipeline.phases.check_phases(curated_dir: Path = Path("curated"), web_path: Path = WEB_PATH) -> bool` (True when the emitted file has drifted)
  - the CLI `python -m pipeline phases` and `python -m pipeline phases --check`

- [ ] **Step 1: Write the failing test**

Create `data/tests/test_phases.py`:

```python
import json
from pathlib import Path

import pytest

from pipeline.phases import PhaseError, check_phases, load_phases, write_phases

CURATED = Path("curated")
WEB = Path("../web/src/data/phases.json")

#: api/internal/phase/phase.go's table, transcribed. The Go side is
#: tested against the file this produces; this is the other direction, so
#: that an edit to the curated file that nobody meant shows up here too.
BOUNDARIES = [
    ("pre-beta", "0001-01-01T00:00:00Z"),
    ("beta", "2026-09-17T00:00:00Z"),
    ("launch", "2026-11-04T23:00:00Z"),
    ("raids-1", "2026-12-09T00:00:00Z"),
]


def test_the_curated_calendar_is_the_go_tables_boundaries_in_order():
    assert [(p.name, p.start) for p in load_phases(CURATED)] == BOUNDARIES


def test_the_sentinel_is_not_a_boundary():
    """`later` means 'unreleased, no date' on a loot source (contract
    10.4). It is not a phase: phase.At would bracket into it."""
    assert "later" not in {p.name for p in load_phases(CURATED)}


def test_the_emitted_web_copy_matches_the_curated_one():
    assert not check_phases(CURATED, WEB)
    emitted = json.loads(WEB.read_text(encoding="utf-8"))
    assert [(row["name"], row["start"]) for row in emitted] == BOUNDARIES


def test_a_start_that_is_not_rfc3339_utc_is_refused(tmp_path):
    (tmp_path / "phases.json").write_text(
        json.dumps([{"name": "beta", "start": "2026-09-17"}]), encoding="utf-8"
    )
    with pytest.raises(PhaseError, match="RFC 3339"):
        load_phases(tmp_path)


def test_boundaries_out_of_order_are_refused(tmp_path):
    (tmp_path / "phases.json").write_text(
        json.dumps(
            [
                {"name": "launch", "start": "2026-11-04T23:00:00Z"},
                {"name": "beta", "start": "2026-09-17T00:00:00Z"},
            ]
        ),
        encoding="utf-8",
    )
    with pytest.raises(PhaseError, match="order"):
        load_phases(tmp_path)


def test_a_duplicate_phase_name_is_refused(tmp_path):
    (tmp_path / "phases.json").write_text(
        json.dumps(
            [
                {"name": "beta", "start": "2026-09-17T00:00:00Z"},
                {"name": "beta", "start": "2026-11-04T23:00:00Z"},
            ]
        ),
        encoding="utf-8",
    )
    with pytest.raises(PhaseError, match="twice"):
        load_phases(tmp_path)


def test_write_phases_is_deterministic(tmp_path):
    out = tmp_path / "phases.json"
    write_phases(CURATED, out)
    once = out.read_bytes()
    write_phases(CURATED, out)
    assert out.read_bytes() == once
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_phases.py -q --no-cov`
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.phases'`

- [ ] **Step 3: Write the curated calendar**

Create `data/curated/phases.json`:

```json
[
  { "name": "pre-beta", "start": "0001-01-01T00:00:00Z" },
  { "name": "beta", "start": "2026-09-17T00:00:00Z" },
  { "name": "launch", "start": "2026-11-04T23:00:00Z" },
  { "name": "raids-1", "start": "2026-12-09T00:00:00Z" }
]
```

`pre-beta` starts at Go's zero time, which is what `phase.go` writes as `time.Time{}`; `launch` is 15:00 PST, which is 23:00 UTC. This file carries no `sources`/`notes` block: it is not a claim about the game, it is the same four instants `phase.go` and `dates.json` already state, in one place. The dates themselves are sourced in `data/curated/loot/forever-raid-phases.json`.

- [ ] **Step 4: Write the module**

Create `data/pipeline/phases.py`:

```python
"""The content phase calendar, stated once.

Contract 10.4. The four boundaries were written out twice -- as a Go
table in `api/internal/phase/phase.go` and as prose in
`web/src/data/dates.json` -- with a comment asking whoever moved one to
move the other. `curated/phases.json` is now the statement; the Go
package is tested against it (the API lane's test), the web reads the
copy this emits, and `GET /v1/phases` serves it.

`later` is deliberately not in here. It is contract 10.4's sentinel for a
loot source with no known date, and a boundary by that name would make
`phase.At` bracket a moment into it.
"""

from __future__ import annotations

import json
from datetime import datetime
from pathlib import Path

from pipeline.models import PhaseBoundary
from pipeline.normalize import write_records

WEB_PATH = Path("../web/src/data/phases.json")

#: The sentinel a loot source carries when its date is unknown. Named
#: here so the loot validator and this module cannot disagree about it.
OPENS_LATER = "later"


class PhaseError(SystemExit):
    """curated/phases.json says something the pipeline will not publish."""


def _instant(value: str, where: str) -> datetime:
    """Parse an RFC 3339 UTC instant, which is what Go's time.Time reads.

    `Z` and nothing else: an offset would still parse here and then
    compare unequal to the Go table's UTC instants for no visible reason.
    """
    if not value.endswith("Z"):
        raise PhaseError(f"{where} start {value!r} is not RFC 3339 UTC (it must end in Z)")
    try:
        return datetime.fromisoformat(value.replace("Z", "+00:00"))
    except ValueError as error:
        raise PhaseError(f"{where} start {value!r} is not RFC 3339: {error}") from error


def load_phases(curated_dir: Path = Path("curated")) -> list[PhaseBoundary]:
    path = curated_dir / "phases.json"
    if not path.exists():
        raise PhaseError(f"missing {path}")
    rows = [PhaseBoundary(**entry) for entry in json.loads(path.read_text(encoding="utf-8"))]
    if not rows:
        raise PhaseError(f"{path} lists no phase")
    seen: set[str] = set()
    previous: datetime | None = None
    for row in rows:
        if row.name in seen:
            raise PhaseError(f"{path} names phase {row.name!r} twice")
        seen.add(row.name)
        if row.name == OPENS_LATER:
            raise PhaseError(
                f"{path} names a phase {OPENS_LATER!r}; that is the sentinel a loot "
                f"source carries when its date is unknown, not a boundary"
            )
        start = _instant(row.start, f"{path} phase {row.name}")
        if previous is not None and start < previous:
            raise PhaseError(
                f"{path} is out of order: {row.name!r} starts before the phase above it"
            )
        previous = start
    return rows


def write_phases(
    curated_dir: Path = Path("curated"), web_path: Path = WEB_PATH
) -> Path:
    write_records(load_phases(curated_dir), web_path)
    return web_path


def check_phases(
    curated_dir: Path = Path("curated"), web_path: Path = WEB_PATH
) -> bool:
    """True when the emitted copy has drifted from the curated one.

    Compares contents, not timestamps, so it catches an edit made straight
    to the generated file as well as one to the curated list that was
    never re-emitted -- the same gate `pipeline/specs.py`'s
    `check_specs` is.
    """
    import tempfile

    rows = load_phases(curated_dir)
    with tempfile.TemporaryDirectory() as directory:
        expected = Path(directory) / "phases.json"
        write_records(rows, expected)
        wanted = expected.read_text(encoding="utf-8")
    return not web_path.exists() or web_path.read_text(encoding="utf-8") != wanted
```

- [ ] **Step 5: Add the model**

In `data/pipeline/models.py`, beside `SpecRecord`:

```python
class PhaseBoundary(BaseModel):
    """One content phase and the instant it opens (parity contract 10.4)."""

    name: str
    start: str
```

`write_records` writes in the order given rather than sorting by id, which is what this needs: the boundaries are ordered by time, and `PhaseBoundary` has no `id`.

- [ ] **Step 6: Wire up the CLI**

In `data/pipeline/__main__.py`, add the subparser after the `specs` block:

```python
    ph = sub.add_parser("phases", help="emit the phase calendar for the web")
    ph.add_argument("--web", default="../web/src/data/phases.json")
    ph.add_argument(
        "--check",
        action="store_true",
        help="write nothing; exit non-zero if the emitted file has drifted",
    )
```

and the dispatch branch:

```python
    elif args.command == "phases":
        from pathlib import Path

        from pipeline.phases import check_phases, write_phases

        if args.check:
            if check_phases(Path("curated"), Path(args.web)):
                logging.getLogger("pipeline").error(
                    "%s does not match curated/phases.json; run `python -m pipeline phases`",
                    args.web,
                )
                return 1
            return 0
        print(write_phases(Path("curated"), Path(args.web)))
```

- [ ] **Step 7: Emit and verify**

Run: `cd data && uv run python -m pipeline phases && uv run python -m pipeline phases --check`
Expected: prints `../web/src/data/phases.json`, then exits 0 silently

Run: `cd data && uv run pytest tests/test_phases.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 8: Point `phase.go` at the file in a comment**

`api/internal/phase/phase.go`'s package comment says the boundaries are "the site's own dates data (web/src/data/dates.json), copied here as a table" and asks for both to be changed together. The Go-side test is the API lane's; this lane's part is to make the comment name the new source of truth. Change the second paragraph of that package comment to:

```go
// The boundaries are data/curated/phases.json, copied here as a table
// rather than read at runtime: the API does not ship the site's source,
// the dates are four fixed instants, and a ranking bracket that could
// change under a running deployment would silently re-bucket stored
// rows. phase_test.go holds this table to that file, so the two cannot
// drift; `python -m pipeline phases` emits the web's copy from the same
// source.
```

- [ ] **Step 9: Commit**

```bash
git add data/curated/phases.json data/pipeline/phases.py data/pipeline/models.py data/pipeline/__main__.py data/tests/test_phases.py web/src/data/phases.json api/internal/phase/phase.go
git commit -m "$(cat <<'EOF'
feat(data): the phase calendar as data

Contract 10.4. The four boundaries were stated twice -- a Go table and
the site's dates prose -- with a comment asking whoever moved one to
move the other. curated/phases.json is now the statement, the pipeline
emits web/src/data/phases.json from it with a --check drift gate like
specs, and the Go package's comment names it (the test that holds
phase.Boundaries to it is the API lane's).

`later` is refused as a phase name: it is contract 10.4's sentinel for a
loot source with no known date, and a boundary by that name would make
phase.At bracket into it.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 10: `simbuffs.json`, a name and an icon per IDS.md id

Contract 10.4: a new `data/builds/<build>/simbuffs.json` of `{ "entries": { "<id>": { "name", "icon" } } }` for every IDS.md buff, debuff, world buff and consumable id. Design section 4.3 is what needs it: the full buff panel renders "the id's display name and icon from the build", and today the settings bar has neither.

`sim/request/IDS.md` is generated by the sim module from the engine's protobuf descriptors, so it is the id list and this lane parses it rather than restating it. Measured on the pinned engine: the **Buffs** table has 76 rows and the **Consumables** table 92, over **165 distinct ids** (`scroll_of_agility`, `scroll_of_protection` and `scroll_of_strength` appear in both). The **Professions** table is not included: those are slugs on the character, not buffs, and the page already has names for them.

**Resolving an id to a name and an icon.** An id is a protobuf field or enum-value name, not a spell id, so the join is by name, normalised the same way `slugify` normalises a zone: lowercase, apostrophes dropped, everything else to `_`. Two candidate spellings per id, both documented by IDS.md itself — the id as written, and the id with its leading enum-name token removed ("Where a value repeats its enum's name — `Food`'s `FoodGrilledSquid` — the name without that prefix (`grilled_squid`) answers to the id as well") — plus the singular of each, for the engine's plural fields (`innervates`, `power_infusions`, `mana_tide_totems`). Four tables are tried in order:

1. the fork database's `spellIcons` (4,038 rows: name and icon together),
2. the fork database's `itemIcons` (103 rows, same),
3. the build's `spells.json`, whose id resolves an icon through `raw/SpellMisc.csv`'s `SpellIconFileDataID`,
4. the build's `items.json`, whose id resolves an icon through `raw/Item.csv`'s `IconFileDataID`,

both of the last two through `pipeline.icons.icon_names(raw/ManifestInterfaceData.csv)` and `pipeline.icons.resolve_icon`, which is the same machinery every item and talent icon already goes through.

That resolves **151 of the 165**. The remaining **14** are ids the engine spells differently from the client — four Atiesh variants of one item, `curse_of_elements` against the client's "Curse of the Elements", `elixir_of_firepower` against "Elixir of Fire Power" — and they come from a curated file with sources, like every other hand-maintained fact here. **An id that resolves through neither path stops the pipeline**, naming the ids: a buff panel row with a blank name and `icons/.webp` is worse than a build that refuses to emit.

**Files:**
- Create: `data/pipeline/loot/buffs.py`
- Create: `data/curated/simbuffs.json`
- Create: `data/tests/fixtures/loot/IDS.md`, `data/tests/fixtures/loot/spells.json`, `data/tests/fixtures/loot/SpellMisc.csv`, `data/tests/fixtures/loot/Item.csv`, `data/tests/fixtures/loot/ManifestInterfaceData.csv`, `data/tests/fixtures/loot/curated-simbuffs.json`
- Modify: `data/pipeline/models.py` (`SimBuffEntry`, `SimBuffsFile`, `BuffOverride`)
- Modify: `data/pipeline/forkdb.py` (`ForkDatabase.spell_icon_rows`, `.item_icon_rows`)
- Modify: `data/tests/test_forkdb.py` (the two new fields)
- Test: `data/tests/test_loot_buffs.py`

**Interfaces:**
- Consumes: `pipeline.forkdb.ForkDatabase`; `pipeline.icons.icon_names`, `resolve_icon`, `PLACEHOLDER_ICON`; `pipeline.curated.parse_sources`; `pipeline.csvio.read_csv`, `int_column` (from `pipeline.normalize.gear`)
- Produces:
  - `pipeline.models.SimBuffEntry` with `name: str`, `icon: str`
  - `pipeline.models.SimBuffsFile` with `entries: dict[str, SimBuffEntry]`
  - `pipeline.models.BuffOverride` with `name: str`, `item_id: int = 0`, `spell_id: int = 0`
  - `pipeline.forkdb.ForkDatabase.spell_icon_rows: tuple[dict, ...]` and `.item_icon_rows: tuple[dict, ...]`
  - `pipeline.loot.buffs.SIMBUFFS: str` = `"simbuffs.json"`; `IDS_MD: Path` = `Path("../sim/request/IDS.md")`
  - `pipeline.loot.buffs.BuffError(SystemExit)`
  - `pipeline.loot.buffs.normalise(name: str) -> str`
  - `pipeline.loot.buffs.ids_md_ids(text: str) -> list[str]`
  - `pipeline.loot.buffs.NameTable` — a `dict[str, tuple[str, str]]` alias, normalised name -> (display name, icon)
  - `pipeline.loot.buffs.fork_tables(fork: ForkDatabase) -> list[NameTable]`
  - `pipeline.loot.buffs.client_tables(build_dir: Path, raw: Path) -> tuple[list[NameTable], IconFor]` where `IconFor` is `Callable[[int, int], str]` taking `(item_id, spell_id)`
  - `pipeline.loot.buffs.load_overrides(curated_dir: Path = Path("curated"), name: str = SIMBUFFS) -> dict[str, BuffOverride]`
  - `pipeline.loot.buffs.build_simbuffs(ids: list[str], tables: list[NameTable], overrides: dict[str, BuffOverride], icon_for: IconFor) -> SimBuffsFile`

- [ ] **Step 1: Write the fixtures**

Create `data/tests/fixtures/loot/IDS.md` — the two tables in the real file's shape, small:

```markdown
# The sim request's id vocabulary

## Buffs

| id | lands in |
| --- | --- |
| `arcane_brilliance` | RaidBuffs |
| `hunters_mark` | Debuffs |
| `innervates` | IndividualBuffs |
| `curse_of_elements` | Debuffs |

## Consumables

| id | sets |
| --- | --- |
| `food_grilled_squid` | Consumes.food |
| `elixir_of_fire_power` | Consumes.spell_power_buff |
| `arcane_brilliance` | Consumes.scroll |

## Professions

| slug | engine enum |
| --- | --- |
| `alchemy` | Profession.Alchemy |
```

Create `data/tests/fixtures/loot/spells.json`:

```json
[
  { "id": 1459, "name": "Arcane Brilliance" },
  { "id": 1130, "name": "Hunter's Mark" },
  { "id": 29166, "name": "Innervate" }
]
```

Create `data/tests/fixtures/loot/SpellMisc.csv`:

```
ID,SpellID,SpellIconFileDataID
1,1459,900001
2,1130,900002
3,29166,900003
```

Create `data/tests/fixtures/loot/Item.csv` — only the two columns read, for the items the fixture `items.json` already has:

```
ID,ClassID,SubclassID,IconFileDataID
110,2,7,900004
```

Create `data/tests/fixtures/loot/ManifestInterfaceData.csv`:

```
ID,FilePath,FileName
900001,Interface/Icons/,Spell_Holy_ArcaneIntellect.blp
900002,Interface/Icons/,Ability_Hunter_SniperShot.blp
900003,Interface/Icons/,Spell_Nature_Lightning.blp
900004,Interface/Icons/,INV_Sword_04.blp
```

Create `data/tests/fixtures/loot/curated-simbuffs.json` — the override file's shape, one entry:

```json
{
  "sources": [
    { "label": "A test source", "url": "https://example.invalid/buffs", "kind": "site" }
  ],
  "notes": "One id the name join cannot reach.",
  "entries": {
    "curse_of_elements": { "spell_id": 1459, "name": "Curse of the Elements" }
  }
}
```

Finally, add two rows to `data/tests/fixtures/loot/db.json`'s `spellIcons` array, so the fork tables cover the prefix-stripping and exact-match paths:

```json
    { "id": 24800, "name": "Grilled Squid", "icon": "inv_misc_fish_13" },
    { "id": 7845, "name": "Elixir of Fire Power", "icon": "inv_potion_60" }
```

- [ ] **Step 2: Write the failing test**

Create `data/tests/test_loot_buffs.py`:

```python
import json
from pathlib import Path

import pytest

from pipeline.forkdb import load_fork_database
from pipeline.loot.buffs import (
    BuffError,
    build_simbuffs,
    client_tables,
    fork_tables,
    ids_md_ids,
    load_overrides,
)

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"
CURATED = Path("curated")


def ids():
    return ids_md_ids((ENGINE / "IDS.md").read_text(encoding="utf-8"))


def tables_and_icons():
    client, icon_for = client_tables(ENGINE, ENGINE)
    return fork_tables(load_fork_database(ENGINE)) + client, icon_for


def built():
    tables, icon_for = tables_and_icons()
    return build_simbuffs(
        ids(), tables, load_overrides(ENGINE, name="curated-simbuffs.json"), icon_for
    ).entries


def test_ids_md_yields_the_buff_and_consumable_ids_once_each_not_professions():
    assert ids() == [
        "arcane_brilliance",
        "curse_of_elements",
        "elixir_of_fire_power",
        "food_grilled_squid",
        "hunters_mark",
        "innervates",
    ]
    assert "alchemy" not in ids()


def test_every_id_resolves_to_a_name_and_an_icon():
    entries = built()
    assert set(entries) == set(ids())
    assert all(entry.name and entry.icon for entry in entries.values())


def test_an_id_matching_a_client_spell_takes_the_clients_icon():
    entries = built()
    assert entries["arcane_brilliance"].name == "Arcane Brilliance"
    assert entries["arcane_brilliance"].icon == "spell_holy_arcaneintellect"


def test_an_apostrophe_and_a_plural_both_resolve():
    entries = built()
    assert entries["hunters_mark"].name == "Hunter's Mark"
    assert entries["innervates"].name == "Innervate"


def test_an_enum_name_prefix_is_stripped_the_way_ids_md_says():
    entries = built()
    assert entries["food_grilled_squid"].name == "Grilled Squid"


def test_a_curated_override_wins_over_every_join():
    entries = built()
    assert entries["curse_of_elements"].name == "Curse of the Elements"
    # spell 1459's icon in the fixture: the override named a client row,
    # so the art came from the build and not from the curated file.
    assert entries["curse_of_elements"].icon == "spell_holy_arcaneintellect"


def test_an_id_nothing_resolves_stops_the_build_and_names_it():
    tables, icon_for = tables_and_icons()
    with pytest.raises(BuffError, match="mystery_buff"):
        build_simbuffs(["mystery_buff"], tables, {}, icon_for)


def test_an_override_naming_both_or_neither_id_is_refused(tmp_path):
    for entry in ({"name": "X"}, {"name": "X", "item_id": 1, "spell_id": 2}):
        (tmp_path / "simbuffs.json").write_text(
            json.dumps(
                {
                    "sources": [
                        {"label": "l", "url": "https://x.invalid", "kind": "site"}
                    ],
                    "notes": "x",
                    "entries": {"curse_of_elements": entry},
                }
            ),
            encoding="utf-8",
        )
        with pytest.raises(BuffError, match="exactly one"):
            load_overrides(tmp_path)


def test_the_real_override_file_is_sourced_and_covers_only_real_ids():
    document = json.loads((CURATED / "simbuffs.json").read_text(encoding="utf-8"))
    assert document["sources"] and document["notes"].strip()
    real = set(ids_md_ids(Path("../sim/request/IDS.md").read_text(encoding="utf-8")))
    assert set(document["entries"]) <= real
    assert len(document["entries"]) == 14
```

- [ ] **Step 3: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_loot_buffs.py -q --no-cov`
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.loot.buffs'`

- [ ] **Step 4: Add the models**

In `data/pipeline/models.py`, beside `ConsumableRecord`:

```python
class SimBuffEntry(BaseModel):
    name: str
    icon: str


class SimBuffsFile(BaseModel):
    """`simbuffs.json` (parity contract 10.4): the display name and icon
    for every id `sim/request/IDS.md` lets a request name, so the settings
    bar's full buff panel can render one."""

    entries: dict[str, SimBuffEntry]


class BuffOverride(BaseModel):
    """A curated IDS.md entry: the display name, and the client row whose
    icon it takes. Exactly one of the two ids. An explicit icon is not
    accepted -- art the build does not ship would reach the site as a
    404, and the point of naming a row is that the build resolves it."""

    name: str
    item_id: int = 0
    spell_id: int = 0
```

- [ ] **Step 5: Write the module**

Create `data/pipeline/loot/buffs.py`:

```python
"""`simbuffs.json`: what to call every id a sim request may name.

Parity contract 10.4, for design section 4.3's full buff panel. An id in
`sim/request/IDS.md` is a protobuf field or enum-value name, not a spell
id, so nothing downstream can render it: this file is the display name
and icon per id, resolved once at build time.

IDS.md is generated by the sim module from the engine's own descriptors,
so it is parsed here rather than restated -- an id the engine adds shows
up as a build that needs a name, not as a panel row nobody notices is
missing.

The join is by normalised name, with the two spellings IDS.md itself
documents (the id, and the id without its leading enum-name token) and
the singular of each for the engine's plural fields. 151 of the 165 ids
resolve that way on the pinned engine and build; the other 14 are
`data/curated/simbuffs.json`. An id that resolves through neither is a
hard error: a row with a blank name and `icons/.webp` is worse than a
build that refuses to emit.
"""

from __future__ import annotations

import json
import re
from collections.abc import Callable
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.curated import parse_sources
from pipeline.forkdb import ForkDatabase
from pipeline.icons import icon_names, resolve_icon
from pipeline.models import BuffOverride, SimBuffEntry, SimBuffsFile

SIMBUFFS = "simbuffs.json"
IDS_MD = Path("../sim/request/IDS.md")

#: Normalised name -> (display name, icon).
type NameTable = dict[str, tuple[str, str]]

_ROW = re.compile(r"^\| `([a-z0-9_]+)` \|", re.M)
#: The IDS.md sections whose ids this file covers. "Professions" is
#: deliberately absent: those are character slugs, not buffs.
_SECTIONS = ("## Buffs", "## Consumables")


class BuffError(SystemExit):
    """An IDS.md id has no name the build can show."""


def normalise(name: str) -> str:
    """A display name as an id-shaped key: lowercase, apostrophes gone."""
    return re.sub(r"[^a-z0-9]+", "_", name.lower().replace("'", "")).strip("_")


def ids_md_ids(text: str) -> list[str]:
    """Every buff and consumable id IDS.md lists, sorted and deduplicated.

    Three ids (`scroll_of_agility`, `scroll_of_protection`,
    `scroll_of_strength`) are in both tables, which is why this is a set
    and not a concatenation. A missing section is an error rather than an
    empty answer: IDS.md is generated, and a shape change there must not
    silently empty the buff panel.
    """
    found: set[str] = set()
    for marker in _SECTIONS:
        if marker not in text:
            raise BuffError(f"{marker!r} is not in IDS.md; its shape changed")
        section = text.split(marker, 1)[1].split("\n## ", 1)[0]
        found |= set(_ROW.findall(section))
    return sorted(found)


def _candidates(buff_id: str) -> list[str]:
    """The spellings a display name may take for this id, best first."""
    out: list[str] = []
    for base in (buff_id, buff_id[:-1] if buff_id.endswith("s") else None):
        if base is None:
            continue
        out.append(base)
        if "_" in base:
            out.append(base.split("_", 1)[1])
    return out


def fork_tables(fork: ForkDatabase) -> list[NameTable]:
    """The fork's own icon tables: name and icon in one row, so no client
    lookup is needed for the 97 ids they cover."""
    return [_table_from(fork.spell_icon_rows), _table_from(fork.item_icon_rows)]


def _table_from(rows: tuple[dict, ...]) -> NameTable:
    table: NameTable = {}
    for row in rows:
        if row.get("name") and row.get("icon"):
            table.setdefault(normalise(row["name"]), (row["name"], row["icon"]))
    return table


#: (item_id, spell_id) -> the icon the build ships for whichever is set.
type IconFor = Callable[[int, int], str]


def client_tables(build_dir: Path, raw: Path) -> tuple[list[NameTable], IconFor]:
    """The build's own spell and item names, with icons resolved through
    `ManifestInterfaceData` exactly as every other icon here is.

    Returns the two tables and the resolver, because a curated override
    names a client row rather than an icon and needs the same lookup.
    """
    names = icon_names(read_csv(raw / "ManifestInterfaceData.csv"))
    spell_icon = {
        int(row["SpellID"]): int(row["SpellIconFileDataID"] or 0)
        for row in read_csv(raw / "SpellMisc.csv")
        if row.get("SpellID")
    }
    item_icon = {
        int(row["ID"]): int(row["IconFileDataID"] or 0)
        for row in read_csv(raw / "Item.csv")
    }

    def table(path: Path, icons: dict[int, int], what: str) -> NameTable:
        built: NameTable = {}
        for row in json.loads(path.read_text(encoding="utf-8")):
            key = normalise(row["name"])
            if key in built:
                continue
            icon = resolve_icon(
                icons.get(int(row["id"]), 0), names, f"{what} {row['id']} ({row['name']})"
            )
            built[key] = (row["name"], icon)
        return built

    def icon_for(item_id: int, spell_id: int) -> str:
        if item_id:
            return resolve_icon(item_icon.get(item_id, 0), names, f"item {item_id}")
        return resolve_icon(spell_icon.get(spell_id, 0), names, f"spell {spell_id}")

    return (
        [
            table(build_dir / "spells.json", spell_icon, "spell"),
            table(build_dir / "items.json", item_icon, "item"),
        ],
        icon_for,
    )


def load_overrides(
    curated_dir: Path = Path("curated"), name: str = SIMBUFFS
) -> dict[str, BuffOverride]:
    """The curated ids, validated. A missing file is no overrides, which
    is a legal state only while every id happens to join."""
    path = curated_dir / name
    if not path.exists():
        return {}
    document = json.loads(path.read_text(encoding="utf-8"))
    parse_sources(document.get("sources", []), str(path))
    if not document.get("notes", "").strip():
        raise BuffError(f"{path} has no notes; say why each id needs a hand")
    overrides: dict[str, BuffOverride] = {}
    for key, value in document["entries"].items():
        override = BuffOverride(**value)
        if bool(override.item_id) == bool(override.spell_id):
            raise BuffError(
                f"{path} entry {key!r} must name exactly one of item_id or spell_id, "
                f"so its icon comes from the build rather than from this file"
            )
        overrides[key] = override
    return overrides


def build_simbuffs(
    ids: list[str],
    tables: list[NameTable],
    overrides: dict[str, BuffOverride],
    icon_for: IconFor,
) -> SimBuffsFile:
    entries: dict[str, SimBuffEntry] = {}
    unresolved: list[str] = []
    for buff_id in ids:
        override = overrides.get(buff_id)
        if override is not None:
            entries[buff_id] = SimBuffEntry(
                name=override.name,
                icon=icon_for(override.item_id, override.spell_id),
            )
            continue
        found = None
        for candidate in _candidates(buff_id):
            for table in tables:
                if candidate in table:
                    found = table[candidate]
                    break
            if found is not None:
                break
        if found is None:
            unresolved.append(buff_id)
            continue
        entries[buff_id] = SimBuffEntry(name=found[0], icon=found[1])
    if unresolved:
        raise BuffError(
            f"{len(unresolved)} IDS.md id(s) have no name in the fork database or "
            f"this build: {', '.join(unresolved)}. Add each to "
            f"data/curated/simbuffs.json with a source, or fix the join."
        )
    return SimBuffsFile(entries=dict(sorted(entries.items())))
```

`fork_tables` needs the fork's raw icon rows, not just the id→icon maps, so add them to `ForkDatabase` in `data/pipeline/forkdb.py`:

```python
    #: The raw icon rows, kept alongside the id maps because `simbuffs`
    #: joins on the name in them and `enchants` joins on the id.
    spell_icon_rows: tuple[dict, ...]
    item_icon_rows: tuple[dict, ...]
```

filled in `load_fork_database` with `spell_icon_rows=tuple(raw.get("spellIcons", []))` and `item_icon_rows=tuple(raw.get("itemIcons", []))`. Task 3's `test_the_tables_are_immutable` covers the new fields for free; extend `test_the_fixture_database_loads_every_table` with `assert len(fork.spell_icon_rows) == 5` and `assert len(fork.item_icon_rows) == 1`.

- [ ] **Step 6: Write the curated overrides**

Create `data/curated/simbuffs.json` — the fourteen ids the name join cannot reach on the pinned engine and build, each with the client row it means:

```json
{
  "sources": [
    {
      "label": "Wowhead, Classic item and spell database",
      "url": "https://www.wowhead.com/classic",
      "kind": "community"
    },
    {
      "label": "This site, build 1.60.1.69893's own items.json and spells.json",
      "url": "https://foreversixty.gg/data/1.60.1.69893/items.json",
      "kind": "site"
    }
  ],
  "notes": "The fourteen sim request ids the engine spells differently from the client, so no name join reaches them. Each names the client row it means; the icon is that row's, read off the build. Four Atiesh ids share one item (22589) and differ only in which class's proc they name, which is why their display names are written out. `curse_of_elements` is the client's 'Curse of the Elements'; `elixir_of_firepower` is 'Elixir of Fire Power' (6373) and `elixir_of_ogres_strength` is 'Elixir of Ogre Strength' (3391), both without the engine's spelling; `sayges_fortune` is the Darkmoon Faire buff, whose client rows are one spell per stat, so it takes the generic name and the Strength row's icon. Delete an entry the day the join reaches it.",
  "entries": {
    "atiesh_druid": { "item_id": 22589, "name": "Atiesh (Druid)" },
    "atiesh_mage": { "item_id": 22589, "name": "Atiesh (Mage)" },
    "atiesh_priest": { "item_id": 22589, "name": "Atiesh (Priest)" },
    "atiesh_warlock": { "item_id": 22589, "name": "Atiesh (Warlock)" },
    "curse_of_elements": { "spell_id": 11723, "name": "Curse of the Elements" },
    "sayges_fortune": { "spell_id": 23735, "name": "Sayge's Dark Fortune" },
    "conjured_rogue_thistle_tea": { "item_id": 7676, "name": "Thistle Tea" },
    "dragon_breath_chili": { "item_id": 12217, "name": "Dragonbreath Chili" },
    "elixir_of_firepower": { "item_id": 6373, "name": "Elixir of Fire Power" },
    "elixir_of_ogres_strength": { "item_id": 3391, "name": "Elixir of Ogre Strength" },
    "food_bless_sunfruit": { "item_id": 13810, "name": "Blessed Sunfruit" },
    "food_dirges_kick_chimaerok_chops": { "item_id": 21023, "name": "Dirge's Kickin' Chimaerok Chops" },
    "food_smoked_desert_dumpling": { "item_id": 20452, "name": "Smoked Desert Dumplings" },
    "sapper_goblin_sapper": { "item_id": 10646, "name": "Goblin Sapper Charge" }
  }
}
```

Every entry names the client row it means, never an icon: the icon is resolved from the build by `icon_for`, so a curated override cannot point at art this client does not ship.

- [ ] **Step 7: Run the tests**

Run: `cd data && uv run pytest tests/test_loot_buffs.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 8: Commit**

```bash
git add data/pipeline/loot/buffs.py data/pipeline/models.py data/pipeline/forkdb.py data/curated/simbuffs.json data/tests/test_loot_buffs.py data/tests/fixtures/loot/
git commit -m "$(cat <<'EOF'
feat(data): simbuffs.json, a name and icon per IDS.md id

Contract 10.4, for the design's full buff panel: an IDS.md id is a
protobuf field name, not a spell id, so nothing downstream could render
one. IDS.md is parsed rather than restated, so an id the engine adds
shows up as a build that needs a name.

151 of the 165 ids resolve by normalised name against the fork's icon
tables and the build's own spells and items, icons through
ManifestInterfaceData like every other icon here. The other 14 are
curated with the client row each means -- the icon is still read off
the build, never typed in. An id nothing resolves stops the pipeline.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 11: The `loot` command, and the build's new files

One command writes all five outputs and refreshes the manifest, the way `simdb` does. It needs the build's `raw/` CSVs — `Map.csv` for the instance types, `ItemSparse.csv` for the PvP ranks, and `SpellMisc.csv`, `Item.csv` and `ManifestInterfaceData.csv` for `simbuffs.json`'s icons — exactly as `simdb` does, and an engine checkout for the fork database, exactly as `simproto` does.

**Order matters inside the command.** `build_loot` prunes a source the build filter emptied; `apply_overlays` runs on its result, which is what lets task 8's curated, deliberately-empty `raid:onyxias-lair` survive while a generated empty source does not. And **`loot` must run before `simdb`** from now on, because task 15 has `simdb` read the two fork columns `loot` writes onto `items.json`.

**Files:**
- Modify: `data/pipeline/loot/__init__.py`
- Modify: `data/pipeline/__main__.py`
- Modify: `data/README.md` (Commands, Layout, New build checklist, Known gaps)
- Create (generated, committed): `data/builds/1.60.1.69893/loot.json`, `enchants.json`, `suffixes.json`, `simbuffs.json`; modified `items.json` and `manifest.json`

**Interfaces:**
- Consumes: `pipeline.forkdb.load_fork_database`; `pipeline.loot.sources.build_loot`, `instance_types`, `pvp_ranks`; `pipeline.loot.gear.build_enchants`, `build_suffixes`, `suffix_options`, `faction_restrictions`, `apply_fork_columns`; `pipeline.loot.overlay.load_overlays`, `apply_overlays`; `pipeline.loot.buffs.ids_md_ids`, `fork_tables`, `client_tables`, `load_overrides`, `build_simbuffs`, `IDS_MD`, `SIMBUFFS`; `pipeline.normalize.write_document`, `write_records`; `pipeline.manifest.refresh_manifest`
- Produces: `pipeline.loot.LOOT`, `ENCHANTS`, `SUFFIXES` (the filenames; `SIMBUFFS` lives in `pipeline.loot.buffs`); `pipeline.loot.write_loot_files(build: str, engine_dir: Path, root: Path = Path("builds"), overlay_dir: Path = Path("curated/loot"), curated_dir: Path = Path("curated"), ids_md: Path = IDS_MD) -> list[Path]`; the CLI `python -m pipeline loot --build <build> --engine <path>`

- [ ] **Step 1: Write the orchestration**

Replace the contents of `data/pipeline/loot/__init__.py`:

```python
"""The Droptimizer, Top Gear and settings-panel data for one build.

Parity contract section 6 as corrected by 10.4. Five outputs, one command:

    builds/<build>/loot.json        every source the two databases state
    builds/<build>/enchants.json    the enchants Top Gear's picker offers
    builds/<build>/suffixes.json    every random suffix and what it is worth
    builds/<build>/simbuffs.json    a name and icon per IDS.md id
    builds/<build>/items.json       gains `suffixes` and `faction_restriction`

Inputs, and why each is where it is:

* The engine fork's `assets/database/db.json`, at a path, because it moves
  with the engine pin and is not ours to vendor (see `pipeline/forkdb.py`).
* The build's own `zones.json`, `items.json` and `spells.json`, already
  committed. `items.json` is both an input (contract 10.4's build filter,
  and the buff panel's item names) and an output.
* The build's `raw/` CSVs. Like `simdb`, this command needs the downloaded
  DB2 exports and not just the normalized JSON: only `Map` says which zone
  is a raid, only `ItemSparse` says which item needs a PvP rank, and
  `SpellMisc` + `Item` + `ManifestInterfaceData` are where an icon comes
  from. Run `python -m pipeline fetch` first.
* `curated/loot/*.json` and `curated/simbuffs.json`.
* `../sim/request/IDS.md`, generated by the sim module from the engine's
  descriptors, which is the buff id list.

Run it after `normalize` and **before** `simdb`: it adds files to a
normalized build directory, refreshes the manifest to cover them, and
writes the two fork columns `pipeline/simdb/items.py` reads.
"""

from __future__ import annotations

import json
import logging
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.forkdb import load_fork_database
from pipeline.loot.buffs import (
    IDS_MD,
    SIMBUFFS,
    build_simbuffs,
    client_tables,
    fork_tables,
    ids_md_ids,
    load_overrides,
)
from pipeline.loot.gear import (
    apply_fork_columns,
    build_enchants,
    build_suffixes,
    faction_restrictions,
    suffix_options,
)
from pipeline.loot.overlay import apply_overlays, load_overlays
from pipeline.loot.sources import build_loot, instance_types, pvp_ranks
from pipeline.manifest import refresh_manifest
from pipeline.normalize import write_document, write_records
from pipeline.normalize.sockets import check_no_sockets

logger = logging.getLogger(__name__)

LOOT = "loot.json"
ENCHANTS = "enchants.json"
SUFFIXES = "suffixes.json"


def _require(path: Path, command: str) -> None:
    if not path.exists():
        raise SystemExit(f"no {path}; run `python -m pipeline {command}` for this build first")


def write_loot_files(
    build: str,
    engine_dir: Path,
    root: Path = Path("builds"),
    overlay_dir: Path = Path("curated/loot"),
    curated_dir: Path = Path("curated"),
    ids_md: Path = IDS_MD,
) -> list[Path]:
    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    for name in ("zones.json", "items.json", "spells.json"):
        _require(build_dir / name, "normalize")

    fork = load_fork_database(engine_dir)
    sparse_rows = read_csv(raw / "ItemSparse.csv")
    # The same guard normalize runs. A build reaching this command with a
    # socketed item would otherwise get a loot table for gear the rest of
    # the repository cannot model.
    check_no_sockets(sparse_rows, build)
    zone_rows = json.loads((build_dir / "zones.json").read_text(encoding="utf-8"))
    build_items = {
        int(row["id"])
        for row in json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    }

    # Contract 10.4's build filter happens inside build_loot, and its
    # pruning sweep with it -- so this runs BEFORE the overlay, which is
    # what lets a curated source with a deliberately empty item list (the
    # announced-but-unlooted Onyxia's Lair) survive.
    document, stats = build_loot(
        fork,
        {int(row["id"]): row["name"] for row in zone_rows},
        instance_types(read_csv(raw / "Map.csv"), zone_rows),
        pvp_ranks(sparse_rows),
        build_items,
    )
    document = apply_overlays(document, load_overlays(overlay_dir))

    enchants = build_enchants(fork)
    suffixes = build_suffixes(fork)
    client, icon_for = client_tables(build_dir, raw)
    simbuffs = build_simbuffs(
        ids_md_ids(ids_md.read_text(encoding="utf-8")),
        fork_tables(fork) + client,
        load_overrides(curated_dir),
        icon_for,
    )
    write_document(document, build_dir / LOOT)
    write_records(enchants, build_dir / ENCHANTS)
    write_records(suffixes, build_dir / SUFFIXES)
    write_document(simbuffs, build_dir / SIMBUFFS)
    with_suffixes, restricted = apply_fork_columns(
        build_dir, suffix_options(fork), faction_restrictions(fork)
    )

    logger.info(
        "loot: %d sources naming %d items; %d fork ids left out because this build "
        "has no such item, %d fork source entries with no kind dropped; "
        "%d enchants, %d suffixes, %d buff ids; "
        "items.json: %d with suffix options, %d faction-restricted",
        len(document.sources),
        stats.items,
        stats.absent_items,
        stats.dropped_entries,
        len(enchants),
        len(suffixes),
        len(simbuffs.entries),
        with_suffixes,
        restricted,
    )
    refresh_manifest(build_dir)
    return [
        build_dir / LOOT,
        build_dir / ENCHANTS,
        build_dir / SUFFIXES,
        build_dir / SIMBUFFS,
        build_dir / "items.json",
    ]
```

- [ ] **Step 2: Wire up the CLI**

In `data/pipeline/__main__.py`, add the subparser after the `simdb` block:

```python
    lt = sub.add_parser("loot", help="build the Droptimizer and Top Gear data for a build")
    lt.add_argument("--build", required=True)
    lt.add_argument(
        "--engine",
        required=True,
        help="path to the wowsims-forever checkout, e.g. $FOREVER_ENGINE_PATH",
    )
```

and the dispatch branch after the `simdb` one:

```python
    elif args.command == "loot":
        from pathlib import Path

        from pipeline.loot import write_loot_files

        for path in write_loot_files(args.build, Path(args.engine)):
            print(path)
```

- [ ] **Step 3: Verify the command's wiring without touching a build**

Run: `cd data && uv run python -m pipeline loot --build nope --engine /nonexistent`
Expected: exits non-zero with `no raw data at builds/nope/raw; run \`python -m pipeline fetch\` first`

- [ ] **Step 4: Fetch the active build's raw tables**

`raw/` is gitignored, so it has to be downloaded. This takes a few minutes and ~40 MB.

Run: `cd data && uv run python -m pipeline fetch --product wow_classic_beta --build 1.60.1.69893`
Expected: prints `1.60.1.69893`; `data/builds/1.60.1.69893/raw/` now holds the CSVs and `_meta.json`

Sanity check that nothing else in the build directory changed:

Run: `git status --short data/builds/1.60.1.69893`
Expected: no output (`raw/` is ignored)

- [ ] **Step 5: Generate the files**

Run: `cd data && ENGINE=${FOREVER_ENGINE_PATH:-/Users/jh/code/wowsims-forever} uv run python -m pipeline loot --build 1.60.1.69893 --engine "$ENGINE"`

Expected, on the log line and then the five paths:
```
INFO pipeline.loot: loot: 63 sources naming 2507 items; 1809 fork ids left out because this build has no such item, 864 fork source entries with no kind dropped; 173 enchants, 1168 suffixes, 165 buff ids; items.json: 69 with suffix options, 819 faction-restricted
builds/1.60.1.69893/loot.json
builds/1.60.1.69893/enchants.json
builds/1.60.1.69893/suffixes.json
builds/1.60.1.69893/simbuffs.json
builds/1.60.1.69893/items.json
```

63, not 62: the generator emits 62 and task 8's overlay adds `raid:onyxias-lair` back.

If any number differs, the fork checkout is at a different sha than the one these were measured on — check `git -C "$ENGINE" rev-parse --short HEAD` against `sim/enginever/version.go` before changing anything in the pipeline.

- [ ] **Step 6: Check the shape by hand once**

Run:
```bash
cd data && uv run python - <<'EOF'
import json
d = json.load(open("builds/1.60.1.69893/loot.json"))
kinds = {}
for s in d["sources"]:
    kinds.setdefault(s["kind"], []).append(s["id"])
for kind, ids in kinds.items():
    print(kind, len(ids), ids[:3])
mc = next(s for s in d["sources"] if s["id"] == "raid:molten-core")
print("molten core:", mc["opens"], len(mc["bosses"]), "bosses,", len(mc.get("trash", [])), "trash")
print("keys on a crafted source:", sorted(
    next(s for s in d["sources"] if s["kind"] == "crafted")))
EOF
```
Expected: `raid 7`, `dungeon 5`, `world 1`, `crafted 5`, `rep 31`, `pvp 13`, `quest 1`; `molten core: later 4 bosses, 5 trash`; the crafted source's keys are exactly `['id', 'items', 'kind', 'name', 'profession']` — no nulls. (`raid 7` is the six the generator emits plus the overlay's Onyxia.)

- [ ] **Step 7: Verify the manifest covers the new files**

Run:
```bash
cd data && uv run python -c "
from pathlib import Path
from pipeline.manifest import read_manifest, verify
m = read_manifest(Path('builds/1.60.1.69893'))
assert verify(Path('builds/1.60.1.69893')) == [], verify(Path('builds/1.60.1.69893'))
print(sorted(n for n in m['files'] if n in ('loot.json','enchants.json','suffixes.json','simbuffs.json','items.json')))
"
```
Expected: `['enchants.json', 'items.json', 'loot.json', 'simbuffs.json', 'suffixes.json']`

- [ ] **Step 8: Update the README**

In `data/README.md`:

Add to the **Commands** block, before the `simdb` line (which is now the order they run in):
```
uv run python -m pipeline loot --build <build> --engine "$FOREVER_ENGINE_PATH"   # needs raw/ and an engine checkout
uv run python -m pipeline phases                                                 # emits web/src/data/phases.json
uv run python -m pipeline phases --check                                         # CI's drift gate; writes nothing
```

and amend the paragraph under it to read:

> `fetch`, `icons`, `tree-art` and `gametables` are the only commands that use the network.
> `normalize`, `diff`, `simdb`, `simconst`, `loot`, `phases` and `specs` are offline and fully
> unit-tested against the fixtures in `tests/fixtures/`. `simproto` and `loot` read a local
> engine checkout and are the only commands that need one.

Add to the **Layout** block, after the `simconsumes.json` line:
```
builds/<build>/loot.json        every source the fork database and the client state, by kind
builds/<build>/enchants.json    the enchants Top Gear's picker offers, with slots and classes
builds/<build>/suffixes.json    every random suffix and the stats it grants
builds/<build>/simbuffs.json    a display name and icon per sim/request/IDS.md id
curated/loot/*.json             overlays: Forever's own loot facts, with sources
curated/simbuffs.json           the IDS.md ids no name join reaches, with sources
curated/phases.json             the content phase calendar; emitted to web/src/data/phases.json
```

Add a **New build checklist** step after the current step 5:

> 6. Run `loot` **before** `simdb`, with an engine checkout: `uv run python -m pipeline loot --build <build> --engine "$FOREVER_ENGINE_PATH"`. It needs the build's `raw/` the way `simdb` does, plus the fork's `assets/database/db.json` at the pinned sha. It writes `loot.json`, `enchants.json`, `suffixes.json`, `simbuffs.json` and `items.json`'s `suffixes` and `faction_restriction` columns — which `simdb` then reads, so the order is not optional — and its log line states the coverage; compare it against `tests/test_loot_build.py`'s constants before committing. Running `normalize` again afterwards clears both columns, so re-run `loot` if you do. If it stops on an unresolved IDS.md id, add that id to `curated/simbuffs.json` with the client row it means and a source.

(renumber the following steps)

Add to **Known gaps**:

> - `ItemRandomSuffix` 404s on build 1.60.1.69893 and `JournalInstance` 404s on both Classic-lineage products, so neither random suffixes nor the Dungeon Journal comes from the client: `suffixes.json` and `dungeons.json` have different answers to that. Suffixes come from the engine fork's own database, which has them; `dungeons.json` stays empty, and `loot.json` gets its dungeon and raid list from `Map.InstanceType` instead.
> - The fork's item database is Classic Era's. Only 2,811 of its 7,553 item ids exist in build 1.60.1.69893's `ItemSparse`, and 1,809 of the ids its sources name are absent from the build's item table — Forever has re-itemised the raid tier and neither database states a source for the replacements. Contract 10.4 has `loot.json` list only items the build has, so those are left out and counted: Molten Core keeps 10 of 139, Zul'Gurub 2 of 98, and Onyxia's Lair 0 of 16. `curated/loot/` is where a sourced Forever fact goes when one exists, and `forever-raid-phases.json` carries the gap per raid.
> - 864 of the fork database's source entries have no kind in contract 6.1 — 501 plain vendors (the 266 rank sets among them are covered by the `pvp` kind, read off the client's `RequiredPVPRank`) and 363 open-world drops from NPCs the fork does not name — and are dropped rather than guessed into a kind.

- [ ] **Step 9: Run the whole suite**

Run: `cd data && uv run ruff check . && uv run pytest`
Expected: PASS, coverage at or above 80%

- [ ] **Step 10: Commit**

```bash
git add data/pipeline/loot/__init__.py data/pipeline/__main__.py data/README.md \
        data/builds/1.60.1.69893/loot.json data/builds/1.60.1.69893/enchants.json \
        data/builds/1.60.1.69893/suffixes.json data/builds/1.60.1.69893/simbuffs.json \
        data/builds/1.60.1.69893/items.json data/builds/1.60.1.69893/manifest.json
git commit -m "$(cat <<'EOF'
feat(data): the loot command, and build 1.60.1.69893's new files

63 sources naming 2,507 items, 173 enchants, 1,168 suffixes, 165 buff
ids, and the two fork columns on items.json: suffix options on 69 rows,
a faction restriction on 819.

The command needs the build's raw/ CSVs the way simdb does -- only Map
says which zone is a raid, only ItemSparse says which item needs a PvP
rank, and icons come from SpellMisc, Item and ManifestInterfaceData --
and an engine checkout the way simproto does. It runs before simdb now,
because simdb reads the columns it writes, and it refreshes the manifest
so the build directory stays verifiable.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 12: The committed build's conformance test

`simdb.bin` and `gametables/` are checked in CI not by regenerating them — CI has no `raw/` and no engine checkout in the test job — but by a test that reads what is committed and holds it to counts and spot values. `loot.json`, `enchants.json`, `suffixes.json`, `simbuffs.json` and `items.json`'s two new columns get the same treatment, in the same style as `tests/test_simdb_build.py` and `tests/test_gametables_build.py`.

Every number below was measured on build 1.60.1.69893 against the pinned fork. A regeneration that moves one is something to look at, which is the whole point of pinning it.

**Files:**
- Create: `data/tests/test_loot_build.py`

**Interfaces:**
- Consumes: the committed `data/builds/1.60.1.69893/{loot,enchants,suffixes,simbuffs,items}.json` and `simdb.bin`; `../sim/request/IDS.md`; `pipeline.forkdb.CLASS_SLUGS`, `ENCHANT_TYPES`, `PROFESSIONS`, `REP_LEVELS`; `pipeline.loot.sources.KIND_ORDER`; `pipeline.loot.buffs.SIMBUFFS`, `ids_md_ids`
- Produces: nothing; this is the gate.

- [ ] **Step 1: Write the test**

Create `data/tests/test_loot_build.py`:

```python
"""What is committed under builds/1.60.1.69893/ must satisfy contract 6.

Like tests/test_simdb_build.py and tests/test_gametables_build.py, this
reads the real output rather than running the pipeline over fixtures:
regenerating needs the build's raw/ CSVs and an engine checkout, and CI's
test job has neither. The counts are the measurement, so a regeneration
that moves one has to be looked at.

One of the handful of tests the Global Constraints allow a build string in.
"""

import json
from functools import cache
from pathlib import Path

from pipeline.forkdb import CLASS_SLUGS, ENCHANT_TYPES, PROFESSIONS, REP_LEVELS
from pipeline.loot.buffs import SIMBUFFS, ids_md_ids
from pipeline.loot.sources import KIND_ORDER

BUILD = "1.60.1.69893"
BUILD_DIR = Path("builds") / BUILD
IDS_MD = Path("../sim/request/IDS.md")

#: Sources per kind. Seven raids: the six the generator emits after
#: contract 10.4's build filter, plus Onyxia's Lair, which the filter
#: empties and curated/loot/forever-raid-phases.json adds back with its
#: announced phase. Five dungeons, one world boss (Lord Kazzak; Azuregos
#: keeps none of his ten drops on this client), five professions,
#: thirty-one faction-and-standing pairs, thirteen PvP ranks, one quest
#: list.
SOURCES_PER_KIND = {
    "raid": 7,
    "dungeon": 5,
    "world": 1,
    "crafted": 5,
    "rep": 31,
    "pvp": 13,
    "quest": 1,
}
TOTAL_SOURCES = 63

RAID_SOURCE_IDS = [
    "raid:ahnqiraj",
    "raid:blackwing-lair",
    "raid:molten-core",
    "raid:naxxramas",
    "raid:onyxias-lair",
    "raid:ruins-of-ahnqiraj",
    "raid:zulgurub",
]
#: raid source id -> (bosses, trash items, distinct items). The gap
#: against what the fork's sources name -- Molten Core 10 of 139,
#: Zul'Gurub 2 of 98 -- is Forever's re-itemisation, and the overlay's
#: notes carry it.
RAID_SHAPE = {
    "raid:ahnqiraj": (12, 7, 68),
    "raid:blackwing-lair": (7, 3, 21),
    "raid:molten-core": (4, 5, 10),
    "raid:naxxramas": (15, 1, 45),
    "raid:onyxias-lair": (0, 0, 0),
    "raid:ruins-of-ahnqiraj": (3, 1, 5),
    "raid:zulgurub": (1, 1, 2),
}
RAID_BOSSES = 42
RAID_ITEMS = 151
#: Bosses the fork database names no NPC for. An invented name would be
#: worse than a blank one, so this is measured rather than forbidden.
UNNAMED_RAID_BOSSES = 31
UNNAMED_DUNGEON_BOSSES = 1

DUNGEON_BOSSES = 6
DUNGEONS_WITH_TRASH = 1
WORLD_SOURCE_IDS = ["world:lord-kazzak"]
CRAFTED_ITEMS = {
    "crafted:blacksmithing": 190,
    "crafted:enchanting": 3,
    "crafted:engineering": 47,
    "crafted:leatherworking": 191,
    "crafted:tailoring": 157,
}
QUEST_ITEMS = 1058
PVP_ITEMS_PER_RANK = {5: 2, 6: 16, 7: 6, 8: 6, 9: 23, 10: 2, 11: 66, 12: 93,
                      14: 63, 15: 2, 16: 80, 17: 48, 18: 42}

#: Every distinct item id the file names. Contract 10.4: all of them are
#: the build's own, the 1,809 the fork names and this client does not
#: having been left out.
NAMED_ITEMS = 2507

ENCHANT_ROWS = 173
ENCHANT_EFFECT_IDS = 150
SUFFIX_ROWS = 1168
ITEMS_WITH_SUFFIXES = 69
ITEMS_FACTION_RESTRICTED = 819
SIMBUFF_ENTRIES = 165

PHASES = {"pre-beta", "beta", "launch", "raids-1"}
OPENS_LATER = "later"
#: The one raid the phase calendar has a date for, per the site's
#: dates.json and the overlay that records it.
DATED_RAID = "raid:onyxias-lair"


@cache
def loot() -> dict:
    return json.loads((BUILD_DIR / "loot.json").read_text(encoding="utf-8"))


@cache
def by_id() -> dict[str, dict]:
    return {source["id"]: source for source in loot()["sources"]}


@cache
def enchants() -> list[dict]:
    return json.loads((BUILD_DIR / "enchants.json").read_text(encoding="utf-8"))


@cache
def suffixes() -> list[dict]:
    return json.loads((BUILD_DIR / "suffixes.json").read_text(encoding="utf-8"))


@cache
def items() -> list[dict]:
    return json.loads((BUILD_DIR / "items.json").read_text(encoding="utf-8"))


@cache
def simbuffs() -> dict:
    return json.loads((BUILD_DIR / SIMBUFFS).read_text(encoding="utf-8"))


def source_items(source: dict) -> set[int]:
    return set(
        source.get("items", [])
        + source.get("trash", [])
        + [item for boss in source.get("bosses", []) for item in boss["items"]]
    )


def test_the_file_is_an_object_with_one_key():
    assert list(loot()) == ["sources"]


def test_every_kind_has_the_number_of_sources_measured():
    counted: dict[str, int] = {}
    for source in loot()["sources"]:
        counted[source["kind"]] = counted.get(source["kind"], 0) + 1
    assert counted == SOURCES_PER_KIND
    assert sum(SOURCES_PER_KIND.values()) == len(loot()["sources"]) == TOTAL_SOURCES


def test_sources_are_ordered_by_kind_then_id():
    order = [(KIND_ORDER.index(s["kind"]), s["id"]) for s in loot()["sources"]]
    assert order == sorted(order)


def test_source_ids_are_unique():
    ids = [source["id"] for source in loot()["sources"]]
    assert len(ids) == len(set(ids))


def test_the_raid_sources_are_the_seven_measured_with_their_shape():
    assert sorted(s["id"] for s in loot()["sources"] if s["kind"] == "raid") == RAID_SOURCE_IDS
    for source_id, (bosses, trash, distinct) in RAID_SHAPE.items():
        source = by_id()[source_id]
        assert len(source.get("bosses", [])) == bosses, source_id
        assert len(source.get("trash", [])) == trash, source_id
        assert len(source_items(source)) == distinct, source_id
    assert sum(shape[0] for shape in RAID_SHAPE.values()) == RAID_BOSSES
    assert len({
        item for s in loot()["sources"] if s["kind"] == "raid" for item in source_items(s)
    }) == RAID_ITEMS


def test_the_one_curated_raid_survives_the_generators_pruning():
    """`build_loot` drops a source the build filter emptied; the overlay
    runs after it, so an announced raid whose loot table nobody knows
    stays in the picker with an empty item list."""
    onyxia = by_id()[DATED_RAID]
    assert onyxia["items"] == []
    assert onyxia["zone_id"] == 2159


def test_a_boss_without_a_name_is_blank_and_counted_not_invented():
    for kind, expected in (("raid", UNNAMED_RAID_BOSSES), ("dungeon", UNNAMED_DUNGEON_BOSSES)):
        blank = [
            boss
            for source in loot()["sources"]
            if source["kind"] == kind
            for boss in source.get("bosses", [])
            if boss["name"] == ""
        ]
        assert len(blank) == expected, kind


def test_every_boss_id_is_its_source_id_plus_its_npc_id():
    for source in loot()["sources"]:
        for boss in source.get("bosses", []):
            assert boss["id"] == f"{source['id']}:{boss['npc_id']}"
            assert boss["npc_id"] > 0


def test_the_dungeon_sources_are_the_five_that_survived_the_filter():
    dungeons = [s for s in loot()["sources"] if s["kind"] == "dungeon"]
    assert len(dungeons) == SOURCES_PER_KIND["dungeon"]
    assert sum(len(s.get("bosses", [])) for s in dungeons) == DUNGEON_BOSSES
    assert sum(1 for s in dungeons if s.get("trash")) == DUNGEONS_WITH_TRASH


def test_the_only_world_source_is_the_one_world_boss_with_loot_left():
    assert [s["id"] for s in loot()["sources"] if s["kind"] == "world"] == WORLD_SOURCE_IDS


def test_crafted_rep_pvp_and_quest_carry_their_own_keys_and_counts():
    assert {
        s["id"]: len(s["items"]) for s in loot()["sources"] if s["kind"] == "crafted"
    } == CRAFTED_ITEMS
    for source in loot()["sources"]:
        if source["kind"] == "crafted":
            assert source["profession"] in PROFESSIONS.values()
        if source["kind"] == "rep":
            assert source["standing"] in REP_LEVELS.values()
            assert source["faction_id"] > 0
        if source["kind"] == "pvp":
            assert source["rank"] > 0
    assert {
        s["rank"]: len(s["items"]) for s in loot()["sources"] if s["kind"] == "pvp"
    } == PVP_ITEMS_PER_RANK
    assert len(by_id()["quest"]["items"]) == QUEST_ITEMS


def test_a_source_only_carries_the_keys_its_kind_needs():
    """`write_document` drops the unset ones, so a crafted source has no
    null `bosses` for the page to filter out."""
    assert sorted(by_id()["crafted:tailoring"]) == ["id", "items", "kind", "name", "profession"]
    assert sorted(by_id()["quest"]) == ["id", "items", "kind", "name"]
    assert "profession" not in by_id()["raid:molten-core"]


def test_every_item_list_is_sorted_and_free_of_duplicates():
    for source in loot()["sources"]:
        for items_list in [source.get("items"), source.get("trash")] + [
            boss["items"] for boss in source.get("bosses", [])
        ]:
            if items_list is None:
                continue
            assert items_list == sorted(set(items_list))


def test_every_raid_is_gated_and_nothing_else_is():
    for source in loot()["sources"]:
        opens = source.get("opens")
        if source["id"] == DATED_RAID:
            assert opens == "raids-1"
        elif source["kind"] == "raid":
            assert opens == OPENS_LATER, source["id"]
        else:
            assert opens is None, source["id"]
        assert opens is None or opens in PHASES | {OPENS_LATER}


def test_the_file_names_only_items_this_build_has():
    """Contract 10.4. The 1,809 ids the fork's sources name that this
    client does not carry are left out, which is what makes every row
    renderable and simmable."""
    named = {item for source in loot()["sources"] for item in source_items(source)}
    assert len(named) == NAMED_ITEMS
    assert named <= {row["id"] for row in items()}


def test_the_enchant_table_is_the_forks_with_its_repeated_effect_ids():
    rows = enchants()
    assert len(rows) == ENCHANT_ROWS
    assert len({row["id"] for row in rows}) == ENCHANT_EFFECT_IDS
    assert [(r["id"], r["spell_id"], r["item_id"]) for r in rows] == sorted(
        (r["id"], r["spell_id"], r["item_id"]) for r in rows
    )


def test_every_enchant_names_a_slot_an_icon_and_a_known_shape():
    for row in enchants():
        assert row["icon"], row["id"]
        assert row["slots"], row["id"]
        assert row["item_types"] and set(row["item_types"]) <= set(ENCHANT_TYPES.values())
        assert set(row["classes"]) <= set(CLASS_SLUGS.values())


def test_every_enchant_effect_id_is_one_the_engine_can_apply():
    """A picker row the engine's own database has no stats for would sim
    as nothing. All 150 are in simdb.bin today."""
    from pipeline.simproto import pb

    database = pb.SimDatabase()
    database.ParseFromString((BUILD_DIR / "simdb.bin").read_bytes())
    have = {enchant.effect_id for enchant in database.enchants}
    assert {row["id"] for row in enchants()} <= have


def test_the_suffix_table_is_complete_and_every_option_resolves():
    rows = suffixes()
    assert len(rows) == SUFFIX_ROWS
    assert [row["id"] for row in rows] == sorted(row["id"] for row in rows)
    known = {row["id"] for row in rows}
    referenced = {suffix for row in items() for suffix in row["suffixes"]}
    assert referenced <= known


def test_items_json_carries_both_fork_columns_on_every_row():
    assert all("suffixes" in row and "faction_restriction" in row for row in items())
    assert sum(1 for row in items() if row["suffixes"]) == ITEMS_WITH_SUFFIXES
    assert sum(1 for row in items() if row["faction_restriction"]) == ITEMS_FACTION_RESTRICTED
    assert {row["faction_restriction"] for row in items()} == {
        "",
        "alliance_only",
        "horde_only",
    }


def test_simbuffs_names_every_id_the_engine_lets_a_request_send():
    entries = simbuffs()["entries"]
    assert len(entries) == SIMBUFF_ENTRIES
    assert set(entries) == set(ids_md_ids(IDS_MD.read_text(encoding="utf-8")))
    for buff_id, entry in entries.items():
        assert entry["name"].strip(), buff_id
        assert entry["icon"].strip(), buff_id
```

- [ ] **Step 2: Run the test**

Run: `cd data && uv run pytest tests/test_loot_build.py -q --no-cov`
Expected: PASS, 21 tests

- [ ] **Step 3: Run the whole suite and lint**

Run: `cd data && uv run ruff check . && uv run pytest`
Expected: PASS, coverage at or above 80%

- [ ] **Step 4: Commit**

```bash
git add data/tests/test_loot_build.py
git commit -m "$(cat <<'EOF'
test(data): hold the committed build's loot data to its measurements

The same gate simdb.bin and gametables/ already have: CI's test job has
no raw/ and no engine checkout, so what is committed is checked rather
than regenerated. 63 sources naming 2,507 items, all of them the
build's own; 7 raids (6 generated plus the curated Onyxia) with 42
bosses over 151 items; 5 dungeons, 1 world boss, 173 enchants, 1,168
suffixes, 165 buff ids, and both fork columns on every items.json row.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 13: The Makefile target and the CI job

How the existing build files are handled, and what this matches:

- **Regeneration** happens in `data.yml`'s `fetch` job, which only runs on `workflow_dispatch`, downloads `raw/`, runs each per-build emitter, and commits `builds/`. `loot` joins that list.
- **Checking** happens in `data.yml`'s `test` job, which runs `ruff`, `specs --check` and `pytest` — and pytest is where `test_simdb_build.py`, `test_gametables_build.py` and now `test_loot_build.py` hold the committed files to their measurements. Nothing is regenerated there, because the test job has neither `raw/` nor an engine checkout. `phases --check` joins `specs --check` beside it: `phases` is not a per-build emitter, so the fetch job never runs it, and its drift gate belongs in the test job for exactly the reason `specs`' does.
- **The engine checkout** is cloned at the pinned sha, the way `sim.yml`'s `rotations` job already does it (`actions/checkout` treats a short sha as a branch name, so it is `git clone --filter=blob:none` plus `git checkout`).

**Files:**
- Modify: `Makefile`
- Modify: `.github/workflows/data.yml`
- Modify: `.gitignore`

**Interfaces:**
- Consumes: `pipeline.loot.write_loot_files` via the CLI; the Makefile's existing `ENGINE_DIR` and `ACTIVE_BUILD` variables
- Produces: `make loot`, `make loot-check`; the `loot` step in the data workflow's fetch job

- [ ] **Step 1: Add the Makefile targets**

Append to `Makefile`, after the `apl-check` target:

```make
.PHONY: loot
# loot regenerates the ACTIVE build's Droptimizer and Top Gear data:
# loot.json, enchants.json, suffixes.json, and items.json's `suffixes`
# column. Nothing in a DB2 export says which boss drops an item, so the
# engine fork's own assets/database/db.json is the only source for it,
# and this is the only target that reads it.
#
# Two prerequisites the target cannot supply itself:
#   * the build's raw/ CSVs (Map.csv for the instance types,
#     ItemSparse.csv for the PvP ranks), the same as `python -m pipeline
#     simdb` -- run `python -m pipeline fetch` for the build first;
#   * an engine checkout at ENGINE_DIR, as for engine-pin and apl-sync.
#
# ENGINE_DIR goes through `abspath` because the recipe runs from data/
# and a relative override would resolve against the wrong root.
loot:
	@test -n "$(ACTIVE_BUILD)" || { echo "$(ACTIVE_BUILD_JSON) names no build"; exit 1; }
	@test -f "$(ENGINE_DIR)/assets/database/db.json" || { \
	  echo "no item database at $(ENGINE_DIR)/assets/database/db.json; set ENGINE_DIR"; exit 1; }
	@test -d "data/builds/$(ACTIVE_BUILD)/raw" || { \
	  echo "no data/builds/$(ACTIVE_BUILD)/raw; run \`python -m pipeline fetch\` first"; exit 1; }
	@(cd data && uv run python -m pipeline loot \
	  --build "$(ACTIVE_BUILD)" --engine "$(abspath $(ENGINE_DIR))")

.PHONY: loot-check
# loot-check is the offline gate on what `loot` wrote: the committed
# build's counts, and the curated overlays' shape and provenance. It
# needs no engine checkout and no raw/, which is why CI runs it in the
# data workflow's test job (as part of pytest) rather than regenerating
# -- the same split simdb.bin and gametables/ already have.
loot-check:
	@(cd data && uv run python -m pipeline phases --check)
	@(cd data && uv run pytest tests/test_loot_build.py tests/test_loot_overlay.py \
	  tests/test_loot_buffs.py tests/test_phases.py -q --no-cov)
```

- [ ] **Step 2: Verify the targets**

Run: `make loot-check`
Expected: PASS

Run: `make loot ENGINE_DIR=/nonexistent`
Expected: exits non-zero with `no item database at /nonexistent/assets/database/db.json; set ENGINE_DIR`

- [ ] **Step 3: Add the regeneration step to the workflow**

In `.github/workflows/data.yml`, in the `test` job, add a step after `specs --check`:

```yaml
      # phases, like specs, is not a per-build emitter, so the fetch job
      # never runs it. This is the gate that web/src/data/phases.json was
      # re-emitted after curated/phases.json moved -- and that nobody
      # edited the generated copy by hand.
      - run: uv run python -m pipeline phases --check
```

In the `fetch` job, insert after the second `normalize` step and **before** the `simdb` step — `simdb` reads the two fork columns `loot` writes onto `items.json`:

```yaml
      # loot.json, enchants.json and suffixes.json come out of the engine
      # fork's own item database, which is not in this repository: nothing
      # in a DB2 export says which boss drops an item. The fork is cloned
      # at the PINNED sha, the way sim.yml's `rotations` job does it
      # (actions/checkout treats a short sha as a branch name and fails to
      # fetch it), so the committed files name one engine build.
      - name: the pinned engine sha
        id: pin
        working-directory: .
        run: |
          sha=$(sed -n 's/.*Version = "\(.*\)"/\1/p' sim/enginever/version.go)
          test -n "$sha" || { echo "sim/enginever/version.go has no Version"; exit 1; }
          echo "sha=$sha" >> "$GITHUB_OUTPUT"
      - name: the engine fork at the pin
        working-directory: .
        run: |
          git clone --quiet --filter=blob:none https://github.com/jhunthrop/wowsims-forever .engine
          git -C .engine checkout --quiet "${{ steps.pin.outputs.sha }}"
      # Before simdb: it writes items.json's `suffixes` and
      # `faction_restriction`, which pipeline/simdb/items.py reads into
      # SimItem. Both refresh the same manifest, and the diff step at the
      # end sees the finished directory either way.
      - run: uv run python -m pipeline loot --build "$BUILD" --engine ../.engine
```

Also extend the workflow's `paths` filters (both `push` and `pull_request`) so a pin bump, an IDS.md regeneration, or an edit to the emitted phase calendar runs the data tests:

```yaml
    paths: ['data/**', 'sim/specs/**', 'sim/enginever/version.go', 'sim/request/IDS.md', 'web/src/lib/sim/specs.ts', 'web/src/data/phases.json', '.github/workflows/data.yml']
```

- [ ] **Step 4: Keep the clone out of the tree**

The commit step runs `git add builds` and `git add diffs`, so the untracked `.engine/` clone is never staged by CI. Add it to `.gitignore` anyway, beside the other scratch entries, so a local reproduction of the same commands cannot stage it either:

```
# the engine fork, cloned at the pin by CI (sim.yml's rotations job and
# data.yml's fetch job) and by anyone reproducing them locally
.engine/
```

- [ ] **Step 5: Lint the workflow**

Run: `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/data.yml')); print('ok')"`
Expected: `ok`

- [ ] **Step 6: Run the full data suite one more time**

Run: `cd data && uv run ruff check . && uv run python -m pipeline specs --check && uv run pytest`
Expected: PASS, coverage at or above 80%

- [ ] **Step 7: Commit**

```bash
git add Makefile .github/workflows/data.yml .gitignore
git commit -m "$(cat <<'EOF'
ci(data): regenerate and check the loot data

`make loot` regenerates the active build's loot, enchant, suffix and
buff files from an engine checkout; `make loot-check` is the offline
gate on what it wrote, the phase calendar included.

CI matches what simdb.bin and gametables/ already do: the fetch job
regenerates (cloning the fork at the pinned sha the way sim.yml's
rotations job does, and running loot before simdb because simdb reads
the columns loot writes) and the test job checks the committed files
through pytest, because it has neither raw/ nor an engine checkout.
`phases --check` joins `specs --check` there. A pin bump, an IDS.md
regeneration or an edit to the emitted calendar now runs the data
tests.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 14: Re-vendor the engine protos at the new pin

Contract 10.3's last bullet: "Two pins move together: `sim/enginever/version.go` (sim module lane) and `data/proto/ENGINE_SHA` with the vendored protos (data lane, `python -m pipeline genproto`), **in that order, in the same round**."

**Do not start this task until the sim module lane has landed its pin bump.** The order is not cosmetic: `sim/enginever/version.go` names the sha the artifacts are built from, and this lane's vendored copies must be of that same sha. Re-vendoring first would mean the Python bindings describe a proto the shipped engine does not have, and every generated file this lane writes would be pinned to a sha nothing else uses.

The engine additions this picks up (contract 10.3) are `Encounter.movement`, `targets_over_time` and `target_dummy`, `SimOptions.sample_iteration`, `RaidSimResult.sample_iteration` with `SampleIteration`/`SampleCast`, and the four new `SimItem` fields task 15 fills.

**Files:**
- Modify: `data/proto/common.proto`, `data/proto/apl.proto`, `data/proto/shaman.proto`, `data/proto/ENGINE_SHA`
- Modify: `data/pipeline/simproto/common_pb2.py`, `apl_pb2.py`, `shaman_pb2.py`
- Test: `data/tests/test_genproto.py` (extend)

**Interfaces:**
- Consumes: `pipeline.genproto.refresh(engine_path: Path, proto_dir: Path, out_dir: Path) -> str` (unchanged)
- Produces: a vendored proto set at the new sha, and therefore `pb.SimItem` with `unique`, `required_level`, `faction_restriction` and `random_suffix_options`, which task 15 needs.

- [ ] **Step 1: Check the sim module lane has moved first**

Run:
```bash
cd /Users/jh/code/forever && \
  echo "sim pin:  $(sed -n 's/.*Version = "\(.*\)"/\1/p' sim/enginever/version.go)" && \
  echo "data pin: $(cat data/proto/ENGINE_SHA)"
```
Expected: the two differ, and the sim pin is the new sha. **If they are equal, this task has nothing to do — stop here.** If the sim pin is still the old one, the sim module lane has not landed; wait rather than re-vendoring ahead of it.

- [ ] **Step 2: Write the failing test**

Append to `data/tests/test_genproto.py`:

```python
def test_the_vendored_pin_matches_the_engine_the_artifacts_are_built_from():
    """Contract 10.3: the two pins move together, sim first. A data lane
    pinned to a different sha would generate against a proto the shipped
    engine does not have."""
    sim = re.search(
        r'Version = "(.*)"', Path("../sim/enginever/version.go").read_text(encoding="utf-8")
    )
    assert sim is not None
    assert Path("proto/ENGINE_SHA").read_text(encoding="utf-8").strip() == sim.group(1)


def test_sim_item_carries_the_fields_contract_10_3_adds():
    from pipeline.simproto import pb

    fields = {field.name for field in pb.SimItem.DESCRIPTOR.fields}
    assert {
        "unique",
        "required_level",
        "faction_restriction",
        "random_suffix_options",
    } <= fields
```

(add `import re` and `from pathlib import Path` to the module's imports if they are not already there.)

- [ ] **Step 3: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_genproto.py -q --no-cov`
Expected: FAIL — the pins differ, and `SimItem` has none of the four fields

- [ ] **Step 4: Re-vendor and regenerate**

Run:
```bash
cd data && uv run python -m pipeline simproto --engine "${FOREVER_ENGINE_PATH:-/Users/jh/code/wowsims-forever}"
```
Expected: prints the new short sha, which must equal `sim/enginever/version.go`'s

- [ ] **Step 5: Run the tests**

Run: `cd data && uv run pytest -q --no-cov && uv run ruff check .`
Expected: PASS. If `tests/test_simdb_*.py` now fails on a stat index, the engine renumbered `Stat`; `pipeline/simdb/statmap.py` resolves by name for exactly that reason, so investigate before touching a constant.

- [ ] **Step 6: Commit**

```bash
git add data/proto data/pipeline/simproto data/tests/test_genproto.py
git commit -m "$(cat <<'EOF'
chore(data): re-vendor the engine protos at the new pin

Contract 10.3's ordering: the sim module moves sim/enginever/version.go
first, this lane's data/proto/ENGINE_SHA and generated bindings follow
in the same round. Picks up Encounter.movement, targets_over_time,
target_dummy, the sample-iteration messages, and the four SimItem
fields the next task fills. A test now holds the two pins equal.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 15: `SimItem`'s four new fields

Contract 10.3: "`SimItem` … gains `bool unique = 20; int32 required_level = 21; UIItem.FactionRestriction faction_restriction = 22; repeated int32 random_suffix_options = 23;` and the data lane fills them in `pipeline/simdb/items.py`." `sim/bulk`'s expander needs all four: unique-equipped, required level, faction and the suffix options a "copy and modify" candidate can pick from.

Two come straight from `ItemSparse`, which `build_sim_items` already reads: `unique` is `MaxCount == 1` (the same test `pipeline/normalize/gear.py` already applies for `GearItem.unique`) and `required_level` is `RequiredLevel`. The other two are the fork columns task 4 put on `items.json`, read back here — the client states neither, and this keeps `pipeline/simdb` free of any engine checkout.

**Depends on task 14**: the four fields do not exist on `pb.SimItem` until the protos are re-vendored.

**Files:**
- Modify: `data/pipeline/simdb/items.py` (`build_sim_items` and its signature)
- Modify: `data/pipeline/simdb/__init__.py` (pass the new argument)
- Test: `data/tests/test_simdb_items.py`, `data/tests/test_simdb_build.py`

**Interfaces:**
- Consumes: `pb.SimItem`, `pb.UIItem.FactionRestriction`; `pipeline.normalize.gear.int_column`; the committed `items.json`'s `suffixes` and `faction_restriction` columns
- Produces: `pipeline.simdb.items.FACTION_RESTRICTION_BY_SLUG: dict[str, str]` (the column's value -> the proto enum value name) and `build_sim_items(..., fork_columns: dict[int, tuple[list[int], str]])` — item id -> (suffix options, faction restriction slug) — appended as the last parameter.

- [ ] **Step 1: Write the failing test**

Append to `data/tests/test_simdb_items.py` (it already builds items from the 1.60 fixtures; follow the file's existing helper for that):

```python
def test_unique_and_required_level_come_from_item_sparse():
    item = built_item(16866)  # Helm of Might: MaxCount 0, RequiredLevel 60
    assert item.unique is False
    assert item.required_level == 60


def test_max_count_one_is_unique():
    assert built_item(UNIQUE_FIXTURE_ID).unique is True


def test_the_fork_columns_are_carried_through_from_items_json():
    item = built_item(12798, fork_columns={12798: ([5, 6], "horde_only")})
    assert list(item.random_suffix_options) == [5, 6]
    assert item.faction_restriction == pb.UIItem.FactionRestriction.Value(
        "FACTION_RESTRICTION_HORDE_ONLY"
    )


def test_an_item_with_no_fork_column_is_left_unrestricted():
    item = built_item(12798, fork_columns={})
    assert list(item.random_suffix_options) == []
    assert item.faction_restriction == pb.UIItem.FactionRestriction.Value(
        "FACTION_RESTRICTION_UNSPECIFIED"
    )
```

Add a row to `data/tests/fixtures/ItemSparse_1_60.csv` with `MaxCount` 1 for `UNIQUE_FIXTURE_ID` (reuse an existing id's columns and change `ID` and `MaxCount`), and define `UNIQUE_FIXTURE_ID` at the top of the test module to that id.

Append to `data/tests/test_simdb_build.py`:

```python
#: The committed database's four contract-10.3 fields, measured against
#: the committed items.json that feeds two of them.
EXPECTED_UNIQUE = None  # fill from the run in step 5
EXPECTED_FACTION_RESTRICTED = 819
EXPECTED_WITH_SUFFIX_OPTIONS = 69


def test_sim_items_carry_the_four_fields_contract_10_3_adds():
    items = database().items
    assert sum(1 for item in items if item.faction_restriction) == (
        EXPECTED_FACTION_RESTRICTED
    )
    assert sum(1 for item in items if item.random_suffix_options) == (
        EXPECTED_WITH_SUFFIX_OPTIONS
    )
    assert all(item.required_level >= 0 for item in items)
    assert sum(1 for item in items if item.unique) == EXPECTED_UNIQUE
```

`EXPECTED_UNIQUE` is the one number this plan cannot state ahead of the run: it is `MaxCount == 1` over the items `simdb_item_rows` keeps, which no earlier task measured. Step 5 prints it; put that number here before committing, and treat a later change to it the way every other constant in that file is treated.

Note that `EXPECTED_FACTION_RESTRICTED` and `EXPECTED_WITH_SUFFIX_OPTIONS` are the whole-items.json figures from task 11's log line; `simdb` keeps only 4,986 of the build's 19,171 items, so if either assertion fails, compare against the kept set rather than assuming the column is wrong — and pin the smaller number with a comment saying so.

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd data && uv run pytest tests/test_simdb_items.py tests/test_simdb_build.py -q --no-cov`
Expected: FAIL — `TypeError: build_sim_items() got an unexpected keyword argument 'fork_columns'`

- [ ] **Step 3: Fill the fields**

In `data/pipeline/simdb/items.py`, add beside the other enum tables:

```python
#: `items.json`'s `faction_restriction` column -> the proto enum value.
#: The column's vocabulary is `pipeline/forkdb.py`'s FACTION_RESTRICTIONS;
#: this is the other end of it, kept here because only simdb needs the
#: enum and only forkdb needs the slug.
FACTION_RESTRICTION_BY_SLUG = {
    "": "FACTION_RESTRICTION_UNSPECIFIED",
    "alliance_only": "FACTION_RESTRICTION_ALLIANCE_ONLY",
    "horde_only": "FACTION_RESTRICTION_HORDE_ONLY",
}
```

Add `fork_columns: dict[int, tuple[list[int], str]]` as the last parameter of `build_sim_items`, and extend the `pb.SimItem(...)` construction with:

```python
            unique=int_column(sparse, "MaxCount") == 1,
            required_level=int_column(sparse, "RequiredLevel"),
```

then, after the construction and beside the other conditional assignments:

```python
        suffix_options, restriction = fork_columns.get(item_id, ([], ""))
        if suffix_options:
            item.random_suffix_options.extend(suffix_options)
        if restriction not in FACTION_RESTRICTION_BY_SLUG:
            raise ItemDataError(
                f"item {item_id} has faction_restriction {restriction!r} in "
                f"items.json; pipeline/simdb/items.py knows "
                f"{sorted(FACTION_RESTRICTION_BY_SLUG)}"
            )
        item.faction_restriction = pb.UIItem.FactionRestriction.Value(
            FACTION_RESTRICTION_BY_SLUG[restriction]
        )
```

- [ ] **Step 4: Read the columns in the caller**

In `data/pipeline/simdb/__init__.py`, add beside `_set_names`:

```python
def _fork_columns(build_dir: Path) -> dict[int, tuple[list[int], str]]:
    """`items.json`'s two fork-derived columns, for `SimItem`.

    They are on `items.json` and not read from the fork database here so
    that `simdb` never needs an engine checkout; `python -m pipeline loot`
    writes them, which is why it runs first. A build whose items.json
    predates that command has neither key, and every item comes out
    unrestricted with no suffix options -- which is what a build with no
    fork data honestly knows.
    """
    path = build_dir / "items.json"
    if not path.exists():
        raise SystemExit(f"no {path}; run `python -m pipeline normalize` for this build first")
    return {
        int(row["id"]): (row.get("suffixes", []), row.get("faction_restriction", ""))
        for row in json.loads(path.read_text(encoding="utf-8"))
    }
```

and pass `_fork_columns(build_dir)` as `build_sim_items`' last argument in `build_sim_database`.

- [ ] **Step 5: Regenerate the build's database and read the unique count**

Run:
```bash
cd data && ENGINE=${FOREVER_ENGINE_PATH:-/Users/jh/code/wowsims-forever} \
  uv run python -m pipeline loot --build 1.60.1.69893 --engine "$ENGINE" && \
  uv run python -m pipeline simdb --build 1.60.1.69893 && \
  uv run python -c "
from pathlib import Path
from pipeline.simproto import pb
db = pb.SimDatabase(); db.ParseFromString(Path('builds/1.60.1.69893/simdb.bin').read_bytes())
print('unique:', sum(1 for i in db.items if i.unique))
print('faction-restricted:', sum(1 for i in db.items if i.faction_restriction))
print('with suffix options:', sum(1 for i in db.items if i.random_suffix_options))
"
```
Put the three numbers into `tests/test_simdb_build.py`'s constants.

- [ ] **Step 6: Re-embed and check the Go side still builds**

Run: `cd /Users/jh/code/forever && make simdb && (cd sim && go build ./... && go test ./internal/simdb/... )`
Expected: the embedded copy is refreshed and both succeed

- [ ] **Step 7: Run the tests**

Run: `cd data && uv run ruff check . && uv run pytest`
Expected: PASS, coverage at or above 80%

- [ ] **Step 8: Commit**

```bash
git add data/pipeline/simdb/items.py data/pipeline/simdb/__init__.py \
        data/tests/test_simdb_items.py data/tests/test_simdb_build.py \
        data/tests/fixtures/ItemSparse_1_60.csv \
        data/builds/1.60.1.69893/simdb.bin data/builds/1.60.1.69893/manifest.json
git commit -m "$(cat <<'EOF'
feat(data): SimItem's unique, required level, faction and suffixes

Contract 10.3. sim/bulk's expander needs all four: unique-equipped,
required level, faction and the suffix options a copy-and-modify
candidate can pick from.

unique and required_level come straight from ItemSparse, the same
MaxCount == 1 test the planner's GearItem already applies. The other
two are the fork columns items.json carries, read back here rather than
from the fork database, so simdb still needs no engine checkout -- which
is why `pipeline loot` runs before it.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

## Self-review

**Spec coverage.** Contract 6.1 `loot.json`, as corrected by 10.4 — tasks 6, 7, 8, 11. Contract 6.2 `enchants.json` — task 5. Contract 6.3 `suffixes.json` and `items.json`'s `suffixes` — task 4. Contract 6.4 the socket guard — task 1. Contract 10.1 A7's `reference_stat` and its vocabulary — task 2. Contract 10.4's `simbuffs.json` — task 10; its `phases.json` — task 9; its `opens: "later"` — tasks 6, 8, 12; its build filter and the per-raid gap note — tasks 6, 8. Contract 10.3's `SimItem` fields — task 15, and the pin ordering that has to precede them — task 14. The reading layer the contract assumes but does not name (the fork database, `statmap`'s inverse) — task 3. Tests, Makefile and CI — tasks 12 and 13.

Not in this lane, deliberately: `sim/bulk`, the wasm exports, migration 0014, the web routes, the addon export, and the Go test that holds `phase.Boundaries` to `curated/phases.json` (contract 10.4 assigns the file to this lane and the test to the API lane; task 9 produces the file and updates the Go package comment to name it). `dungeons.json` stays empty — `JournalInstance` 404s on **both** Classic-lineage products (checked live against wago.tools on 2026-09-19 for 1.60.1.69893 and for Era 1.15.9.69722) — so `loot.json` takes its dungeon and raid list from `Map.InstanceType` instead, and no task tries to revive it.

**Ordering.** Tasks 1–13 are independent of the other lanes and can run in order today. Tasks 14 and 15 are gated on the sim module lane's pin bump and are last for that reason; task 14's first step refuses to run early. Within the lane, task 3 must precede 4, 5, 6 and 10 (the fork reader); 6 must precede 7 (`KIND_ORDER`); 7 must precede 8; 4, 5, 6, 7, 10 must precede 11; 11 must precede 12; 14 must precede 15.

**Type consistency.** `LootSource` / `LootBoss` / `LootFile` are defined in task 6 and used unchanged in 7, 8, 11 and 12. `LootSourcePatch` is defined in task 7 and used by task 8's file and task 12's phase test. `ForkDatabase`'s fields (`random_suffixes`, `item_icons`, `spell_icons`, and `spell_icon_rows` / `item_icon_rows` added in task 10) are read in 4, 5, 6, 10 and 11. `build_loot` returns `(LootFile, LootStats)` with `items`, `dropped_entries`, `absent_items` in task 6 and is unpacked that way in task 11. `apply_fork_columns` returns `(with_suffixes, restricted)` in task 4 and is unpacked that way in task 11. `client_tables` returns `(tables, icon_for)` in task 10 and is unpacked that way in task 11. `KIND_ORDER` is defined in task 6 and imported by task 7 and task 12. `SIMBUFFS` lives in `pipeline/loot/buffs.py` and is imported by task 11 and task 12. `parse_sources` is renamed in task 7 and called from `curated.py`, `overlay.py` and `buffs.py`. The `faction_restriction` vocabulary has exactly two ends: `forkdb.FACTION_RESTRICTIONS` (task 4) writes the slug, `simdb.items.FACTION_RESTRICTION_BY_SLUG` (task 15) reads it.

**One number this plan does not state.** `tests/test_simdb_build.py`'s `EXPECTED_UNIQUE` in task 15: it is `MaxCount == 1` over the 4,986 items `simdb` keeps, which nothing measured before that task runs. Step 5 prints it and the step says to write it down. Every other constant in the plan is measured.

## Rulings applied from contract section 10

Section 10 rules on all eleven findings the first draft of this plan reported. Ten are adopted as written and are simply what the tasks now do:

- **10.4** — AreaTable zone ids (Molten Core is 2717); `raid:<zone-slug>` and `raid:<zone-slug>:<npc-id>` source ids; an unnamed boss carries an empty name; `pvp:rank-<n>` from `ItemSparse.RequiredPVPRank`; plain vendors and unnamed open-world drops dropped and counted; suffixes from the fork database because the client has no `ItemRandomSuffix`; `enchants.json` keyed by effect id plus spell/item id, with `item_types` as the `EnchantType` shape restriction and `slots` from `type` + `extra_types`; the overlay shape `{sources, notes, add, replace, remove}`.
- **10.1 A7** — the stat vocabulary is the fork's `proto.Stat` enum names in snake case, not `PROTO_STAT_ALIASES`. Task 2 validates against the enum. No value in the `reference_stat` table changed: `spell_power` and `attack_power` are already enum names in snake case.

Three rulings changed the plan's shape rather than confirming it:

- **`opens: "later"`** replaced the first draft's "gate every raid behind `raids-1`". Task 8 now gates the six Era raids as `later` and reserves `raids-1` for Onyxia's Lair, the one raid the calendar dates.
- **`loot.json` lists only items the build has** reversed the first draft's "ids as written, no filter". Every count in tasks 6, 8, 11 and 12 was re-measured against the filtered output: 62 generated sources (63 after the overlay) naming 2,507 items, 42 raid bosses over 151 items, 1,809 ids left out.
- **Three new deliverables** — `simbuffs.json` (task 10), `phases.json` (task 9), and `SimItem`'s four fields with the pin move that has to precede them (tasks 14 and 15).

### Two rulings this lane could not apply as literally as written

1. **10.4's overlay `replace: [loot sources]`**, read strictly as whole `LootSource` objects, would make setting one raid's `opens` a restatement of its bosses and trash, and would silently wipe any key the file left out. Task 7 implements `replace` entries as loot sources with every key but `id` optional, applying only the keys the file actually writes (`model_dump(exclude_unset=True)`), and refusing a `kind` that differs from the source's own. Same shape on disk, usable semantics. 10.4 also drops the first draft's per-item removal, so a wrong item is now corrected by `replace`-ing the list it is in — noted in the module docstring.
2. **10.3's `SimItem.faction_restriction` cannot be filled from the client.** 19,066 of build 1.60.1.69893's 19,171 `ItemSparse` rows carry `AllowableRace` `-1/-1`, every one of the 819 items the fork marks Alliance- or Horde-only among them. Rather than give `pipeline/simdb` an engine checkout, task 4 adds a `faction_restriction` column to `items.json` beside the contract's `suffixes` — one fork-derived pass, one place — and task 15 reads both back. That is one more `items.json` key than 10.4 asks for, and it is the reason `loot` must now run before `simdb`.
