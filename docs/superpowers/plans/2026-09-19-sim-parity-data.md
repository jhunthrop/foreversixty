# Simulator parity: the data lane — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Give the simulator's Droptimizer and Top Gear pages the data they need — a per-build `loot.json`, `enchants.json` and `suffixes.json`, a curated loot overlay, `reference_stat` on every spec, and a build-validation guard that fails the day an item ships with a gem socket.

**Architecture:** Nothing in a DB2 export says which boss drops an item, so the engine fork's own `assets/database/db.json` (AtlasLoot + Wowhead sources for 7,553 Era items, 173 enchants, 1,168 random suffixes, 274 named NPCs, 18 factions) is the only source for loot, enchants and suffixes. A new offline-ish `loot` command reads that database from an engine checkout at the pinned sha, joins it to the build's own `zones.json`, `items.json` and two raw DB2 tables (`Map.csv` for instance types, `ItemSparse.csv` for PvP ranks), applies a curated overlay from `data/curated/loot/*.json`, and writes three files plus a `suffixes` column on `items.json`. The generated files are committed per build, exactly like `simdb.bin` and `gametables/`, and are gated in CI by a conformance test with golden counts rather than by regeneration — CI has no `raw/`.

**Tech Stack:** Python 3.12, uv, pydantic v2, pytest + pytest-cov (80% floor), ruff (line-length 100, `E,F,I,B,UP`). No new dependencies.

**Spec:**
- Design: `docs/superpowers/specs/2026-09-19-simulator-parity-design.md` (sections 4.4, 6.1, 6.2, 12, 13)
- Contract: `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` (section 6 is this lane's binding interface; section 1.4 and 8 for `reference_stat`)

## Global Constraints

- Python `>=3.12` (`data/pyproject.toml`), `.python-version` is `3.12`; everything runs under `uv` from the `data/` directory: `uv run python -m pipeline …`, `uv run pytest`, `uv run ruff check .`.
- Ruff: `line-length = 100`, `target-version = "py312"`, `select = ["E", "F", "I", "B", "UP"]`. Every new file must pass `uv run ruff check .`.
- pytest runs with `--cov=pipeline --cov-report=term-missing --cov-fail-under=80`. New modules need tests in the same run or the suite fails.
- Generated files are committed per build under `data/builds/<build>/`. The active build is `1.60.1.69893` (`web/src/data/active-build.json`).
- Every curated file carries `sources` (a list of `{label, url, kind}` with `kind` in `blizzard | datamined | community | site`) and `notes`. An unsourced claim stops the pipeline.
- Nothing fabricated. An item with no source in either database is simply absent from `loot.json`. A boss the fork database does not name is emitted with an empty `name`, never an invented one.
- `opens` values are phase names from `api/internal/phase/phase.go`: `pre-beta`, `beta`, `launch`, `raids-1`. There is no other raid phase.
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
| `data/pipeline/loot/gear.py` | new — `enchants.json`, `suffixes.json`, and `items.json`'s `suffixes` column |
| `data/pipeline/loot/overlay.py` | new — load and apply `data/curated/loot/*.json` |
| `data/pipeline/models.py` | modified — `SpecRecord.reference_stat`, `Item.suffixes`, and the loot/enchant/suffix records |
| `data/pipeline/simdb/statmap.py` | modified — `stat_keys`, the inverse of `stat_array` |
| `data/pipeline/specs.py` | modified — validate and render `reference_stat` |
| `data/pipeline/curated.py` | modified — `_sources` becomes the public `parse_sources` |
| `data/pipeline/normalize/__init__.py` | modified — call the socket guard; add `write_document` |
| `data/pipeline/__main__.py` | modified — the `loot` subcommand |
| `data/curated/specs.json` | modified — `reference_stat` per spec |
| `data/curated/loot/forever-first-raids.json` | new — the first overlay |
| `Makefile` | modified — `loot` and `loot-check` |
| `.github/workflows/data.yml` | modified — the fork clone and the `loot` step in the fetch job |

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

The table, exactly: attack power for every melee spec and for all three hunter specs, spell power for every caster spec.

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
- Modify: `data/pipeline/specs.py` (`REFERENCE_STATS`, `load_specs`, `render_go`, `render_ts`)
- Modify: `data/curated/specs.json` (all 27 rows)
- Modify: `sim/specs/specs.go`, `web/src/lib/sim/specs.ts` (regenerated, never hand-edited)
- Test: `data/tests/test_specs.py`

**Interfaces:**
- Consumes: `pipeline.simdb.statmap.PROTO_STAT_ALIASES: dict[str, tuple[str, ...]]`
- Produces: `SpecRecord.reference_stat: str`; `pipeline.specs.REFERENCE_STATS: frozenset[str]`; the Go field `ReferenceStat string \`json:"reference_stat"\`` on `specs.Spec`; the TypeScript field `reference_stat: string` on `Spec`.

- [ ] **Step 1: Write the failing tests**

Append to `data/tests/test_specs.py`:

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
Expected: FAIL — `pydantic_core.ValidationError: SpecRecord … reference_stat Field required` (or `AttributeError`)

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

In `data/pipeline/specs.py`, add the import after `from pipeline.models import SpecRecord`:

```python
from pipeline.simdb.statmap import PROTO_STAT_ALIASES
```

Add beside `ROLES`:

```python
#: The stat keys a spec may normalise its weights to: the planner's own stat
#: vocabulary, which is `pipeline/simdb/statmap.py`'s, minus its test probe.
#: Naming a stat the engine has no `Stat` enum entry for would produce a
#: `StatWeights` request the engine refuses at run time, so it is refused here.
REFERENCE_STATS = frozenset(key for key in PROTO_STAT_ALIASES if not key.startswith("__"))
```

In `load_specs`, after the `record.role not in ROLES` check, insert:

```python
        if record.reference_stat not in REFERENCE_STATS:
            raise SpecError(
                f"spec {record.spec} has reference_stat {record.reference_stat!r}; "
                f"use one of {sorted(REFERENCE_STATS)}"
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
      ] }
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
    { "id": 903, "name": "Azuregos" }
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
    assert len(fork.items) == 12
    assert len(fork.enchants) == 4
    assert len(fork.random_suffixes) == 2
    assert fork.zones == {2717: "Molten Core", 1581: "The Deadmines"}
    assert fork.npcs == {900: "Big Boss", 902: "Dungeon Boss", 903: "Azuregos"}
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

### Task 4: `suffixes.json`, and `items.json`'s `suffixes` column

Contract 6.3: `data/builds/<build>/suffixes.json` is `[ { "id", "name", "stats": {...} } ]`, and `items.json` rows gain `suffixes: [id, ...]` where the item rolls one.

The contract sources this from `ItemRandomSuffix`. **That table 404s on build 1.60.1.69893** (it exists for Classic Era 1.15.9.69722 and not for the beta client), which is exactly why `pipeline/simdb/__init__.py` already emits `SimDatabase.random_suffixes` empty. The fork database has the table the client does not: `randomSuffixes` (1,168 rows) and `randomSuffixOptions` on 1,688 items. This task reads those.

Measured on the pinned fork against build 1.60.1.69893: **1,168 suffix records**, and **69** of the build's 19,171 items carry suffix options (the fork's item table is Era's, and only 2,811 of its 7,553 ids exist in this client's `ItemSparse`).

**Files:**
- Modify: `data/pipeline/models.py` (`Item.suffixes`, new `SuffixRecord`)
- Create: `data/pipeline/loot/__init__.py` (package marker only in this task), `data/pipeline/loot/gear.py`
- Modify: `data/tests/golden/items.json` (13 rows gain `"suffixes": []`)
- Create: `data/tests/fixtures/loot/items.json`
- Test: `data/tests/test_loot_gear.py`

**Interfaces:**
- Consumes: `pipeline.forkdb.ForkDatabase`, `pipeline.forkdb.load_fork_database`, `pipeline.simdb.statmap.stat_keys`
- Produces:
  - `pipeline.models.SuffixRecord` with `id: int`, `name: str`, `stats: dict[str, float]`
  - `pipeline.models.Item.suffixes: list[int]` (default `[]`, last key)
  - `pipeline.loot.gear.build_suffixes(fork: ForkDatabase) -> list[SuffixRecord]`
  - `pipeline.loot.gear.suffix_options(fork: ForkDatabase) -> dict[int, list[int]]`
  - `pipeline.loot.gear.apply_suffixes(build_dir: Path, options: dict[int, list[int]]) -> int` (returns how many rows got a non-empty list)

- [ ] **Step 1: Write the fixture item table**

Create `data/tests/fixtures/loot/items.json` — the build's own item table, in `items.json` shape, deliberately missing three ids the fork database names (102, 104, 107) so the coverage gap is exercised:

```json
[
  { "id": 100, "name": "Raid Boss Drop", "quality": 4, "item_level": 76, "required_level": 60, "class_id": 4, "subclass_id": 4, "inventory_type": 5 },
  { "id": 101, "name": "Raid Trash Drop", "quality": 3, "item_level": 66, "required_level": 58, "class_id": 4, "subclass_id": 3, "inventory_type": 9 },
  { "id": 103, "name": "Dungeon Boss Drop", "quality": 3, "item_level": 50, "required_level": 45, "class_id": 4, "subclass_id": 2, "inventory_type": 6 },
  { "id": 106, "name": "Crafted Plate", "quality": 3, "item_level": 60, "required_level": 55, "class_id": 4, "subclass_id": 4, "inventory_type": 5 },
  { "id": 108, "name": "Quest Ring", "quality": 2, "item_level": 40, "required_level": 35, "class_id": 4, "subclass_id": 0, "inventory_type": 11 },
  { "id": 110, "name": "Suffixed Sword", "quality": 2, "item_level": 45, "required_level": 40, "class_id": 2, "subclass_id": 7, "inventory_type": 13 },
  { "id": 111, "name": "Rank Eleven Blade", "quality": 3, "item_level": 65, "required_level": 60, "class_id": 2, "subclass_id": 7, "inventory_type": 13 },
  { "id": 112, "name": "Two Sources", "quality": 4, "item_level": 70, "required_level": 60, "class_id": 4, "subclass_id": 4, "inventory_type": 1 }
]
```

- [ ] **Step 2: Write the failing test**

Create `data/tests/test_loot_gear.py`:

```python
import json
import shutil
from pathlib import Path

from pipeline.forkdb import load_fork_database
from pipeline.loot.gear import apply_suffixes, build_suffixes, suffix_options

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


def test_apply_suffixes_fills_the_column_and_leaves_the_rest_empty(tmp_path):
    build_dir = tmp_path / "1.60.1.69893"
    build_dir.mkdir()
    shutil.copy(ENGINE / "items.json", build_dir / "items.json")
    filled = apply_suffixes(build_dir, suffix_options(fork()))
    assert filled == 1
    rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    by_id = {row["id"]: row for row in rows}
    assert by_id[110]["suffixes"] == [5, 6]
    assert by_id[100]["suffixes"] == []
    assert len(rows) == 8


def test_apply_suffixes_keeps_the_key_order_and_puts_suffixes_last(tmp_path):
    build_dir = tmp_path / "1.60.1.69893"
    build_dir.mkdir()
    shutil.copy(ENGINE / "items.json", build_dir / "items.json")
    apply_suffixes(build_dir, suffix_options(fork()))
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
    ]


def test_apply_suffixes_is_idempotent(tmp_path):
    build_dir = tmp_path / "1.60.1.69893"
    build_dir.mkdir()
    shutil.copy(ENGINE / "items.json", build_dir / "items.json")
    options = suffix_options(fork())
    apply_suffixes(build_dir, options)
    once = (build_dir / "items.json").read_bytes()
    apply_suffixes(build_dir, options)
    assert (build_dir / "items.json").read_bytes() == once
```

- [ ] **Step 3: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_loot_gear.py -q --no-cov`
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.loot'`

- [ ] **Step 4: Add the models**

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
```

and add, beside `ItemSetRecord`:

```python
class SuffixRecord(BaseModel):
    id: int
    name: str
    stats: dict[str, float]
```

- [ ] **Step 5: Create the package and write the builders**

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

from pipeline.forkdb import ForkDatabase
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


def apply_suffixes(build_dir: Path, options: dict[int, list[int]]) -> int:
    """Fill `items.json`'s `suffixes` column. Returns how many rows got one.

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
        Item(**row).model_copy(update={"suffixes": options.get(int(row["id"]), [])})
        for row in rows
    ]
    write_json(items, path)
    return sum(1 for item in items if item.suffixes)
```

- [ ] **Step 6: Update the normalize golden**

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
Expected: 13 rows each gain one `"suffixes": []` line.

- [ ] **Step 7: Run the tests**

Run: `cd data && uv run pytest tests/test_loot_gear.py tests/test_normalize_items.py tests/test_normalize_build.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 8: Commit**

```bash
git add data/pipeline/models.py data/pipeline/loot/ data/tests/test_loot_gear.py data/tests/fixtures/loot/items.json data/tests/golden/items.json
git commit -m "$(cat <<'EOF'
feat(data): random suffixes, and the items.json column for them

Contract 6.3. The contract names ItemRandomSuffix, which 404s on build
1.60.1.69893 -- the same absence that already leaves SimDatabase's
random_suffixes empty -- so the fork database's randomSuffixes and
per-item randomSuffixOptions are the source instead.

items.json's rows gain `suffixes`, appended last and defaulting to
empty, so a build that has not had `pipeline loot` run for it still
emits the key.

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

Contract 6.1. Seven kinds of source, built from the fork database's `sources`, joined to the build's own zones.

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
  - `pipeline.loot.sources.build_loot(fork, zone_names: dict[int, str], types: dict[int, int], ranks: dict[int, int]) -> tuple[LootFile, LootStats]`
  - `pipeline.loot.sources.LootStats(items: int, dropped_entries: int)`

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
    return build_loot(
        fork,
        {row["id"]: row["name"] for row in rows},
        instance_types(read_csv(ENGINE / "Map.csv"), rows),
        pvp_ranks(read_csv(ENGINE / "ItemSparse.csv")),
    )


def source(source_id: str):
    document, _ = built()
    for candidate in document.sources:
        if candidate.id == source_id:
            return candidate
    raise AssertionError(f"no source {source_id}; have {[s.id for s in document.sources]}")


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


def test_the_stats_count_what_was_emitted_and_what_was_dropped():
    _, stats = built()
    # 100, 101, 102, 103, 104, 106, 107, 108, 111, 112
    assert stats.items == 10
    # the vendor-only item and the unnamed open-world mob's drop
    assert stats.dropped_entries == 2
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

    items: int
    dropped_entries: int


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
    fork: ForkDatabase, zone_names: dict[int, str], types: dict[int, int]
) -> tuple[list[LootSource], int]:
    """The raid, dungeon and world sources, and how many drops had no home."""
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
            if kind is not None:
                (bosses[(zone_id, npc_id)] if npc_id else trash[zone_id]).add(item_id)
            elif npc_id in fork.npcs:
                world[npc_id].add(item_id)
            else:
                # An open-world mob the fork does not name, or a drop with no
                # zone at all. The contract has no kind for either.
                dropped += 1

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


def _keyed_sources(fork: ForkDatabase) -> tuple[list[LootSource], list[int], int]:
    """The crafted, rep and quest sources, plus how many entries had no home."""
    crafted: dict[str, set[int]] = defaultdict(set)
    rep: dict[tuple[int, str], set[int]] = defaultdict(set)
    quest: set[int] = set()
    dropped = 0
    for item in fork.items:
        item_id = int(item["id"])
        for source in item.get("sources") or []:
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
            elif "soldBy" in source:
                # A vendor. Rank vendors are covered by the pvp kind, read
                # off the client; anything else has no kind in contract 6.1.
                dropped += 1
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


def _pvp_sources(ranks: dict[int, int]) -> list[LootSource]:
    by_rank: dict[int, set[int]] = defaultdict(set)
    for item_id, rank in ranks.items():
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


def build_loot(
    fork: ForkDatabase,
    zone_names: dict[int, str],
    types: dict[int, int],
    ranks: dict[int, int],
) -> tuple[LootFile, LootStats]:
    drops, dropped_drops = _drop_sources(fork, zone_names, types)
    keyed, quest, dropped_keyed = _keyed_sources(fork)
    sources = [*drops, *keyed, *_pvp_sources(ranks)]
    if quest:
        sources.append(LootSource(id="quest", kind="quest", name="Quests", items=quest))
    sources.sort(key=lambda source: (KIND_ORDER.index(source.kind), source.id))
    named = {
        item_id
        for source in sources
        for item_id in [
            *(source.items or []),
            *(source.trash or []),
            *(item for boss in (source.bosses or []) for item in boss.items),
        ]
    }
    return LootFile(sources=sources), LootStats(
        items=len(named), dropped_entries=dropped_drops + dropped_keyed
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

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 7: The curated loot overlay

Contract 6.1: the generated file is "then overlaid by `data/curated/loot/*.json`, which use the same shape plus `sources` and `notes` like every curated file and may add, replace or remove a source or an item".

The contract's own top-level key for loot sources is `sources`, and `sources` is also every curated file's provenance key. The two cannot both be the top level, so an overlay document keeps `sources`/`notes` in their curated meaning and carries loot-shaped documents underneath:

```json
{
  "sources": [{ "label": "...", "url": "https://...", "kind": "blizzard" }],
  "notes": "why this file exists",
  "add":     { "sources": [ <a whole LootSource> ] },
  "replace": { "sources": [ { "id": "raid:molten-core", "opens": "raids-1" } ] },
  "remove":  { "sources": ["dungeon:x"], "items": [{ "source": "raid:y", "item_id": 123 }] }
}
```

- `add.sources` — whole new sources. An id that already exists is an error; use `replace`.
- `replace.sources` — `id` plus the keys to overwrite, and only those. The id must exist. `kind` is deliberately not overwritable: a source's kind decides the shape of its id and of everything keyed on it.
- `remove.sources` — ids to drop whole. Each must exist.
- `remove.items` — one item id out of one source, wherever it appears in it (`items`, `trash`, or any boss). It must appear, or the overlay is stale and says so.

Files apply in filename order; within a file, `add`, then `replace`, then `remove`. Every mismatch is a hard error, the same policy `pipeline/curated.py` has for an unsourced claim.

**Files:**
- Modify: `data/pipeline/curated.py` (`_sources` becomes the public `parse_sources`; three call sites)
- Modify: `data/pipeline/models.py` (`LootPatch`, `LootPatchFile`, `LootItemRemoval`, `LootRemoval`, `LootOverlay`)
- Create: `data/pipeline/loot/overlay.py`
- Create: `data/tests/fixtures/loot/curated/10-add.json`, `data/tests/fixtures/loot/curated/20-replace-and-remove.json`
- Test: `data/tests/test_loot_overlay.py`

**Interfaces:**
- Consumes: `pipeline.curated.SOURCE_KINDS`, `pipeline.curated.parse_sources`; `pipeline.models.LootFile`, `LootSource`
- Produces:
  - `pipeline.curated.parse_sources(raw: list[dict], where: str) -> list[Source]` (renamed from `_sources`)
  - `pipeline.models.LootOverlay` with `sources: list[Source]`, `notes: str`, `add: LootFile | None`, `replace: LootPatchFile | None`, `remove: LootRemoval | None`
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
  "add": {
    "sources": [
      {
        "id": "raid:barrow-deeps",
        "kind": "raid",
        "name": "Barrow Deeps",
        "opens": "raids-1",
        "bosses": []
      }
    ]
  }
}
```

Create `data/tests/fixtures/loot/curated/20-replace-and-remove.json`:

```json
{
  "sources": [
    { "label": "A test source", "url": "https://example.invalid/two", "kind": "site" }
  ],
  "notes": "Gates the raid, drops a dungeon and one item.",
  "replace": { "sources": [{ "id": "raid:molten-core", "opens": "raids-1" }] },
  "remove": {
    "sources": ["dungeon:the-deadmines"],
    "items": [{ "source": "raid:molten-core", "item_id": 101 }]
  }
}
```

- [ ] **Step 2: Write the failing test**

Create `data/tests/test_loot_overlay.py`:

```python
import json
from pathlib import Path

import pytest

from pipeline.curated import CuratedError
from pipeline.models import LootBoss, LootFile, LootSource
from pipeline.loot.overlay import OverlayError, apply_overlays, load_overlays

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
    assert by_id["raid:molten-core"].opens == "raids-1"
    assert by_id["raid:molten-core"].trash is None or 101 not in by_id["raid:molten-core"].trash
    assert by_id["raid:barrow-deeps"].bosses == []


def test_a_replace_only_touches_the_keys_it_names():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert molten.name == "Molten Core"
    assert molten.zone_id == 2717
    assert [b.items for b in molten.bosses] == [[100]]


def test_the_result_stays_sorted_by_kind_then_id():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    assert [s.id for s in result.sources] == ["raid:barrow-deeps", "raid:molten-core"]


def test_an_overlay_with_no_source_is_refused(tmp_path):
    write(tmp_path, "a.json", {"notes": "x", "add": {"sources": []}})
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
        "add": {"sources": [{"id": "raid:molten-core", "kind": "raid", "name": "Again"}]},
    })
    with pytest.raises(OverlayError, match="already"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_replacing_patching_or_removing_an_unknown_source_is_refused(tmp_path):
    for payload in (
        {"replace": {"sources": [{"id": "raid:nope", "opens": "raids-1"}]}},
        {"remove": {"sources": ["raid:nope"]}},
        {"remove": {"items": [{"source": "raid:nope", "item_id": 1}]}},
    ):
        write(tmp_path, "a.json", {
            "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
            "notes": "x",
        } | payload)
        with pytest.raises(OverlayError, match="raid:nope"):
            apply_overlays(base(), load_overlays(tmp_path))


def test_removing_an_item_the_source_does_not_have_is_refused(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "remove": {"items": [{"source": "raid:molten-core", "item_id": 999}]},
    })
    with pytest.raises(OverlayError, match="999"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_removing_an_item_clears_it_from_a_boss_too(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "remove": {"items": [{"source": "raid:molten-core", "item_id": 100}]},
    })
    result = apply_overlays(base(), load_overlays(tmp_path))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert [b.items for b in molten.bosses] == [[]]


def test_a_replace_may_not_change_a_sources_kind(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "replace": {"sources": [{"id": "raid:molten-core", "kind": "dungeon"}]},
    })
    with pytest.raises(Exception):
        load_overlays(tmp_path)


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
class LootPatch(BaseModel):
    """A partial `LootSource`: the id, and only the keys to overwrite.

    `kind` is absent on purpose. A source's kind decides the shape of its
    id, which candidates carry as `drop:<source id>`, so changing it would
    orphan every reference rather than edit one.
    """

    id: str
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


class LootPatchFile(BaseModel):
    sources: list[LootPatch]


class LootItemRemoval(BaseModel):
    source: str
    item_id: int


class LootRemoval(BaseModel):
    sources: list[str] = []
    items: list[LootItemRemoval] = []


class LootOverlay(BaseModel):
    """One `data/curated/loot/*.json`.

    `sources` and `notes` are the curated provenance every hand-maintained
    fact here carries; the loot sources live inside `add` and `replace`,
    each of which is a `loot.json`-shaped document.
    """

    sources: list[Source] = []
    notes: str = ""
    add: LootFile | None = None
    replace: LootPatchFile | None = None
    remove: LootRemoval | None = None

    model_config = {"extra": "forbid"}
```

- [ ] **Step 6: Write the overlay loader and applier**

Create `data/pipeline/loot/overlay.py`:

```python
"""Forever's own loot facts, laid over the two databases' Era ones.

Parity contract 6.1 and the design's risk list: the fork's item database
covers Classic Era, and Forever re-itemises. 2,811 of the fork's 7,553
item ids exist in the 1.60 client's `ItemSparse` at all, so a great deal
of what a Forever player will actually loot has no source in either
database. This is where a sourced statement about that goes -- and where
the phase a raid opens in goes, since neither database states a date.

Every file is a `LootOverlay`: curated `sources` and `notes` like every
other hand-maintained fact here, plus `add`, `replace` and `remove`.
Files apply in filename order; within a file, add, then replace, then
remove. A stale instruction -- adding a source that exists, patching one
that does not, removing an item that is not there -- is an error, because
the alternative is an overlay that quietly stops doing anything.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.curated import parse_sources
from pipeline.models import LootFile, LootOverlay, LootSource
from pipeline.loot.sources import KIND_ORDER


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


def _without(item_id: int, items: list[int] | None) -> tuple[list[int] | None, bool]:
    if not items or item_id not in items:
        return items, False
    remaining = [candidate for candidate in items if candidate != item_id]
    return remaining or None, True


def _drop_item(source: LootSource, item_id: int) -> tuple[LootSource, bool]:
    items, in_items = _without(item_id, source.items)
    trash, in_trash = _without(item_id, source.trash)
    bosses = source.bosses
    in_boss = False
    if bosses is not None:
        rebuilt = []
        for boss in bosses:
            if item_id in boss.items:
                in_boss = True
                boss = boss.model_copy(
                    update={"items": [c for c in boss.items if c != item_id]}
                )
            rebuilt.append(boss)
        bosses = rebuilt
    updated = source.model_copy(update={"items": items, "trash": trash, "bosses": bosses})
    return updated, in_items or in_trash or in_boss


def apply_overlays(
    document: LootFile, overlays: list[tuple[Path, LootOverlay]]
) -> LootFile:
    by_id = {source.id: source for source in document.sources}
    for path, overlay in overlays:
        for source in (overlay.add.sources if overlay.add else []):
            if source.id in by_id:
                raise OverlayError(
                    f"{path} adds source {source.id!r}, which the generated file "
                    f"already has; use `replace` to change it"
                )
            by_id[source.id] = source
        for patch in (overlay.replace.sources if overlay.replace else []):
            if patch.id not in by_id:
                raise OverlayError(
                    f"{path} replaces keys on source {patch.id!r}, which no source has; "
                    f"use `add`, or delete the stale instruction"
                )
            changes = patch.model_dump(exclude_unset=True)
            changes.pop("id")
            by_id[patch.id] = by_id[patch.id].model_copy(update=changes)
        removal = overlay.remove
        for source_id in (removal.sources if removal else []):
            if source_id not in by_id:
                raise OverlayError(f"{path} removes source {source_id!r}, which is not there")
            del by_id[source_id]
        for entry in (removal.items if removal else []):
            if entry.source not in by_id:
                raise OverlayError(
                    f"{path} removes item {entry.item_id} from source "
                    f"{entry.source!r}, which is not there"
                )
            updated, found = _drop_item(by_id[entry.source], entry.item_id)
            if not found:
                raise OverlayError(
                    f"{path} removes item {entry.item_id} from {entry.source!r}, "
                    f"which does not list it"
                )
            by_id[entry.source] = updated
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

Contract 6.1's overlay, resolving the one collision in it: `sources` is
both the contract's key for loot sources and every curated file's
provenance key, so provenance keeps the top-level name and the
loot-shaped documents live under `add` and `replace`.

Three operations: add a whole source, replace named keys on an existing
one (which is how a raid gets its `opens`), remove a source or one item
from one. A stale instruction is an error, because the alternative is an
overlay that quietly stops doing anything.

curated.py's `_sources` becomes the public `parse_sources` so both
validators state the provenance rule once.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 8: The first overlay — Forever's raid open phases

Two facts the databases cannot state, both required by the design's section 6.1 ("the phase each raid opens") and its risk list ("Forever-only items with no source").

**The phases.** `api/internal/phase/phase.go` has four names and exactly one raid phase, `raids-1` (9 December 2026). So every raid source is gated behind `raids-1` — none of them is open at launch. That is conservative in the right direction: a raid that opens later than `raids-1` still shows as gated, whereas a source with no `opens` shows as open from launch, which would be wrong for all seven.

**The Forever-only items.** Measured on the pinned fork against build 1.60.1.69893: 2,811 of the fork's 7,553 item ids exist in this client's `ItemSparse` at all, and 1,809 of the 4,316 item ids `loot.json` names are not in the build's item table. Forever has re-itemised the raid tier — Bonereaver's Edge (17076), Band of Accuria (17063), every tier-1 and tier-2 set piece are gone from the client — and **no source exists for their replacements in either database**. Per the global constraint, that means this overlay adds no items. It says so, with the measurement, so the next person does not go looking again.

Also recorded, and deliberately *not* acted on: the site's own `web/src/data/dates.json` names the 9 December raids as "Barrow Deeps · Hyjal Summit · Onyxia". The 1.60 client has candidate maps for the first two (2817 "Starfall Barrow Den", 2832 "Nightmare Grove", 2995 "Hyjal Crater"), but nothing sourced maps the announced names onto those ids, and neither database gives them a single item. Naming them would be a guess, so the overlay states the open question instead of answering it.

**Files:**
- Create: `data/curated/loot/forever-raid-phases.json`
- Test: `data/tests/test_loot_overlay.py` (extend)

**Interfaces:**
- Consumes: `pipeline.loot.overlay.load_overlays`
- Produces: the committed overlay, whose seven `replace.sources` ids must match the raid source ids task 6 generates.

- [ ] **Step 1: Write the failing test**

Append to `data/tests/test_loot_overlay.py`:

```python
#: Every raid source the fork database produces for build 1.60.1.69893,
#: from tests/test_loot_sources.py's rules. All seven are gated: the
#: phase calendar has exactly one raid phase.
RAID_SOURCE_IDS = [
    "raid:ahnqiraj",
    "raid:blackwing-lair",
    "raid:molten-core",
    "raid:naxxramas",
    "raid:onyxias-lair",
    "raid:ruins-of-ahnqiraj",
    "raid:zulgurub",
]


def test_the_raid_phase_overlay_gates_every_raid_behind_raids_1():
    loaded = dict(
        (path.name, document) for path, document in load_overlays(CURATED)
    )
    document = loaded["forever-raid-phases.json"]
    patches = document.replace.sources
    assert sorted(patch.id for patch in patches) == RAID_SOURCE_IDS
    assert {patch.opens for patch in patches} == {"raids-1"}


def test_the_raid_phase_overlay_adds_and_removes_nothing():
    """Per the plan's measurement: no Forever-only item has a source in
    either database, so this overlay states that rather than inventing
    one. If it ever gains an `add`, the note has to change with it."""
    loaded = dict((path.name, d) for path, d in load_overlays(CURATED))
    document = loaded["forever-raid-phases.json"]
    assert document.add is None
    assert document.remove is None


def test_every_phase_an_overlay_names_is_a_real_phase():
    """The phase names are api/internal/phase/phase.go's, transcribed:
    an `opens` the API does not know is a filter that matches nothing."""
    phases = {"pre-beta", "beta", "launch", "raids-1"}
    for path, document in load_overlays(CURATED):
        for patch in (document.replace.sources if document.replace else []):
            assert patch.opens is None or patch.opens in phases, (path, patch.id)
        for source in (document.add.sources if document.add else []):
            assert source.opens is None or source.opens in phases, (path, source.id)
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd data && uv run pytest tests/test_loot_overlay.py -q --no-cov`
Expected: FAIL — `KeyError: 'forever-raid-phases.json'`

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
  "notes": "Two things neither database states. First, when a raid opens: no raid is available at launch on 4 November, the first raid tier opens on 9 December, and api/internal/phase/phase.go has exactly one raid phase for it, so all seven raid sources are gated behind raids-1. A source with no `opens` reads as open from launch, which would be wrong for every one of them; a raid that turns out to open after the first tier is still shown as gated, which is the safe direction to be wrong in. Second, Forever's own loot: no item is added here, because no Forever-only item has a source in either database. Measured on the pinned fork against build 1.60.1.69893 -- 2,811 of the fork's 7,553 item ids exist in this client's ItemSparse, and 1,809 of the 4,316 ids loot.json names are absent from the build's item table. The client has re-itemised the raid tier outright (Bonereaver's Edge 17076, Band of Accuria 17063 and every tier-1 and tier-2 set piece are gone) and states no replacement's source anywhere. Also open, and deliberately unanswered: the site's dates.json names the 9 December raids as Barrow Deeps, Hyjal Summit and Onyxia, and the client carries candidate maps for the first two (2817 Starfall Barrow Den, 2832 Nightmare Grove, 2995 Hyjal Crater), but nothing sourced maps the announced names onto those ids and neither database gives them an item, so no source is added for them.",
  "replace": {
    "sources": [
      { "id": "raid:ahnqiraj", "opens": "raids-1" },
      { "id": "raid:blackwing-lair", "opens": "raids-1" },
      { "id": "raid:molten-core", "opens": "raids-1" },
      { "id": "raid:naxxramas", "opens": "raids-1" },
      { "id": "raid:onyxias-lair", "opens": "raids-1" },
      { "id": "raid:ruins-of-ahnqiraj", "opens": "raids-1" },
      { "id": "raid:zulgurub", "opens": "raids-1" }
    ]
  }
}
```

- [ ] **Step 4: Run the tests**

Run: `cd data && uv run pytest tests/test_loot_overlay.py -q --no-cov && uv run ruff check .`
Expected: PASS, `All checks passed!`

- [ ] **Step 5: Commit**

```bash
git add data/curated/loot/forever-raid-phases.json data/tests/test_loot_overlay.py
git commit -m "$(cat <<'EOF'
feat(data): Forever's raid open phases, and the loot gap

The first curated loot overlay. All seven raid sources are gated behind
raids-1: nothing raids at launch, the first tier opens 9 December, and
the phase calendar has one raid phase to say it with. A source with no
`opens` would read as open from launch, which is wrong for every one.

It adds no items, on purpose, and the note carries the measurement: 2,811
of the fork's 7,553 item ids exist in the 1.60 client at all, 1,809 of
the 4,316 ids loot.json names are absent from the build's item table, and
Forever's re-itemised raid loot has no source in either database. The
Barrow Deeps / Hyjal Summit mapping is recorded as an open question
rather than guessed.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 9: The `loot` command, and the build's three files

One command writes all four outputs and refreshes the manifest, the way `simdb` does. It needs the build's `raw/` CSVs — `Map.csv` for the instance types and `ItemSparse.csv` for the PvP ranks — exactly as `simdb` does, and an engine checkout for the fork database, exactly as `simproto` does.

**Files:**
- Modify: `data/pipeline/loot/__init__.py`
- Modify: `data/pipeline/__main__.py`
- Modify: `data/README.md` (Commands, Layout, New build checklist, Known gaps)
- Create (generated, committed): `data/builds/1.60.1.69893/loot.json`, `enchants.json`, `suffixes.json`; modified `items.json` and `manifest.json`

**Interfaces:**
- Consumes: `pipeline.forkdb.load_fork_database`; `pipeline.loot.sources.build_loot`, `instance_types`, `pvp_ranks`; `pipeline.loot.gear.build_enchants`, `build_suffixes`, `suffix_options`, `apply_suffixes`; `pipeline.loot.overlay.load_overlays`, `apply_overlays`; `pipeline.normalize.write_document`, `write_records`; `pipeline.manifest.refresh_manifest`
- Produces: `pipeline.loot.LOOT`, `ENCHANTS`, `SUFFIXES` (the three filenames); `pipeline.loot.write_loot_files(build: str, engine_dir: Path, root: Path = Path("builds"), overlay_dir: Path = Path("curated/loot")) -> list[Path]`; the CLI `python -m pipeline loot --build <build> --engine <path>`

- [ ] **Step 1: Write the orchestration**

Replace the contents of `data/pipeline/loot/__init__.py`:

```python
"""The Droptimizer and Top Gear data for one build.

Parity contract section 6. Four outputs, one command:

    builds/<build>/loot.json        every source the two databases state
    builds/<build>/enchants.json    the enchants Top Gear's picker offers
    builds/<build>/suffixes.json    every random suffix and what it is worth
    builds/<build>/items.json       gains its `suffixes` column

Inputs, and why each is where it is:

* The engine fork's `assets/database/db.json`, at a path, because it moves
  with the engine pin and is not ours to vendor (see `pipeline/forkdb.py`).
* The build's own `zones.json` and `items.json`, already committed, which
  is what "joined to the build's zones" means.
* The build's `raw/Map.csv` and `raw/ItemSparse.csv`. Like `simdb`, this
  command needs the downloaded DB2 exports and not just the normalized
  JSON: only `Map` says which zone is a raid and only `ItemSparse` says
  which item needs a PvP rank. Run `python -m pipeline fetch` first.
* `curated/loot/*.json`, the overlay.

Run it after `normalize`, for the same reason `simdb` is: it adds files to
a normalized build directory and refreshes the manifest to cover them.
"""

from __future__ import annotations

import json
import logging
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.forkdb import load_fork_database
from pipeline.loot.gear import apply_suffixes, build_enchants, build_suffixes, suffix_options
from pipeline.loot.overlay import apply_overlays, load_overlays
from pipeline.loot.sources import build_loot, instance_types, pvp_ranks
from pipeline.manifest import refresh_manifest
from pipeline.normalize import write_document, write_records
from pipeline.normalize.sockets import check_no_sockets

logger = logging.getLogger(__name__)

LOOT = "loot.json"
ENCHANTS = "enchants.json"
SUFFIXES = "suffixes.json"


def write_loot_files(
    build: str,
    engine_dir: Path,
    root: Path = Path("builds"),
    overlay_dir: Path = Path("curated/loot"),
) -> list[Path]:
    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    zones_path = build_dir / "zones.json"
    if not zones_path.exists():
        raise SystemExit(
            f"no {zones_path}; run `python -m pipeline normalize` for this build first"
        )

    fork = load_fork_database(engine_dir)
    sparse_rows = read_csv(raw / "ItemSparse.csv")
    # The same guard normalize runs. A build reaching this command with a
    # socketed item would otherwise get a loot table for gear the rest of
    # the repository cannot model.
    check_no_sockets(sparse_rows, build)
    zone_rows = json.loads(zones_path.read_text(encoding="utf-8"))

    document, stats = build_loot(
        fork,
        {int(row["id"]): row["name"] for row in zone_rows},
        instance_types(read_csv(raw / "Map.csv"), zone_rows),
        pvp_ranks(sparse_rows),
    )
    document = apply_overlays(document, load_overlays(overlay_dir))

    enchants = build_enchants(fork)
    suffixes = build_suffixes(fork)
    write_document(document, build_dir / LOOT)
    write_records(enchants, build_dir / ENCHANTS)
    write_records(suffixes, build_dir / SUFFIXES)
    with_suffixes = apply_suffixes(build_dir, suffix_options(fork))

    build_items = {int(row["id"]) for row in json.loads(
        (build_dir / "items.json").read_text(encoding="utf-8")
    )}
    named = {
        item_id
        for source in document.sources
        for item_id in [
            *(source.items or []),
            *(source.trash or []),
            *(item for boss in (source.bosses or []) for item in boss.items),
        ]
    }
    logger.info(
        "loot: %d sources naming %d items, %d of them not in this build's item table; "
        "%d fork source entries had no kind and were dropped; "
        "%d enchants, %d suffixes, %d items carrying suffix options",
        len(document.sources),
        len(named),
        len(named - build_items),
        stats.dropped_entries,
        len(enchants),
        len(suffixes),
        with_suffixes,
    )
    refresh_manifest(build_dir)
    return [build_dir / LOOT, build_dir / ENCHANTS, build_dir / SUFFIXES,
            build_dir / "items.json"]
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

Expected, on the log line and then the four paths:
```
INFO pipeline.loot: loot: 78 sources naming 4316 items, 1809 of them not in this build's item table; 864 fork source entries had no kind and were dropped; 173 enchants, 1168 suffixes, 69 items carrying suffix options
builds/1.60.1.69893/loot.json
builds/1.60.1.69893/enchants.json
builds/1.60.1.69893/suffixes.json
builds/1.60.1.69893/items.json
```

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
print("molten core:", mc["opens"], len(mc["bosses"]), "bosses,", len(mc["trash"]), "trash")
print("keys on a crafted source:", sorted(
    next(s for s in d["sources"] if s["kind"] == "crafted")))
EOF
```
Expected: `raid 7`, `dungeon 19`, `world 2`, `crafted 5`, `rep 31`, `pvp 13`, `quest 1`; `molten core: raids-1 11 bosses, 45 trash`; the crafted source's keys are exactly `['id', 'items', 'kind', 'name', 'profession']` — no nulls.

- [ ] **Step 7: Verify the manifest covers the new files**

Run:
```bash
cd data && uv run python -c "
from pathlib import Path
from pipeline.manifest import read_manifest, verify
m = read_manifest(Path('builds/1.60.1.69893'))
assert verify(Path('builds/1.60.1.69893')) == [], verify(Path('builds/1.60.1.69893'))
print(sorted(n for n in m['files'] if n in ('loot.json','enchants.json','suffixes.json','items.json')))
"
```
Expected: `['enchants.json', 'items.json', 'loot.json', 'suffixes.json']`

- [ ] **Step 8: Update the README**

In `data/README.md`:

Add to the **Commands** block, after the `simdb` line:
```
uv run python -m pipeline loot --build <build> --engine "$FOREVER_ENGINE_PATH"   # needs raw/ and an engine checkout
```

and amend the paragraph under it to read:

> `fetch`, `icons`, `tree-art` and `gametables` are the only commands that use the network.
> `normalize`, `diff`, `simdb`, `simconst`, `loot` and `specs` are offline and fully unit-tested
> against the fixtures in `tests/fixtures/`. `simproto` and `loot` read a local engine checkout
> and are the only commands that need one.

Add to the **Layout** block, after the `simconsumes.json` line:
```
builds/<build>/loot.json        every source the fork database and the client state, by kind
builds/<build>/enchants.json    the enchants Top Gear's picker offers, with slots and classes
builds/<build>/suffixes.json    every random suffix and the stats it grants
curated/loot/*.json             overlays: Forever's own loot facts, with sources
```

Add a **New build checklist** step after the current step 5:

> 6. Run `loot` after `simdb`, with an engine checkout: `uv run python -m pipeline loot --build <build> --engine "$FOREVER_ENGINE_PATH"`. It needs the build's `raw/` the way `simdb` does, plus the fork's `assets/database/db.json` at the pinned sha. It writes `loot.json`, `enchants.json`, `suffixes.json` and `items.json`'s `suffixes` column, and its log line states the coverage; compare it against `tests/test_loot_build.py`'s constants before committing. Running `normalize` again afterwards clears the `suffixes` column, so re-run `loot` if you do.

(renumber the following steps)

Add to **Known gaps**:

> - `ItemRandomSuffix` 404s on build 1.60.1.69893 and `JournalInstance` 404s on both Classic-lineage products, so neither random suffixes nor the Dungeon Journal comes from the client: `suffixes.json` and `dungeons.json` have different answers to that. Suffixes come from the engine fork's own database, which has them; `dungeons.json` stays empty, and `loot.json` gets its dungeon and raid list from `Map.InstanceType` instead.
> - The fork's item database is Classic Era's. Only 2,811 of its 7,553 item ids exist in build 1.60.1.69893's `ItemSparse`, and 1,809 of the 4,316 ids `loot.json` names are absent from the build's item table — Forever has re-itemised the raid tier and neither database states a source for the replacements. The ids are emitted as written; `curated/loot/` is where a sourced Forever fact goes when one exists.
> - 864 of the fork database's source entries have no kind in contract 6.1 — 501 plain vendors (the 266 rank sets among them are covered by the `pvp` kind, read off the client's `RequiredPVPRank`) and 363 open-world drops from NPCs the fork does not name — and are dropped rather than guessed into a kind.

- [ ] **Step 9: Run the whole suite**

Run: `cd data && uv run ruff check . && uv run pytest`
Expected: PASS, coverage at or above 80%

- [ ] **Step 10: Commit**

```bash
git add data/pipeline/loot/__init__.py data/pipeline/__main__.py data/README.md \
        data/builds/1.60.1.69893/loot.json data/builds/1.60.1.69893/enchants.json \
        data/builds/1.60.1.69893/suffixes.json data/builds/1.60.1.69893/items.json \
        data/builds/1.60.1.69893/manifest.json
git commit -m "$(cat <<'EOF'
feat(data): the loot command, and build 1.60.1.69893's three files

78 sources naming 4,316 items, 173 enchants, 1,168 suffixes, and the
suffixes column on the 69 of the build's items that roll one.

The command needs the build's raw/ CSVs the way simdb does -- only Map
says which zone is a raid and only ItemSparse says which item needs a
PvP rank -- and an engine checkout the way simproto does. It refreshes
the manifest so the build directory stays verifiable.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 10: The committed build's conformance test

`simdb.bin` and `gametables/` are checked in CI not by regenerating them — CI has no `raw/` and no engine checkout in the test job — but by a test that reads what is committed and holds it to counts and spot values. `loot.json`, `enchants.json` and `suffixes.json` get the same treatment, in the same style as `tests/test_simdb_build.py` and `tests/test_gametables_build.py`.

Every number below was measured on build 1.60.1.69893 against the pinned fork. A regeneration that moves one is something to look at, which is the whole point of pinning it.

**Files:**
- Create: `data/tests/test_loot_build.py`

**Interfaces:**
- Consumes: the committed `data/builds/1.60.1.69893/{loot,enchants,suffixes,items}.json`; `pipeline.forkdb.CLASS_SLUGS`, `ENCHANT_TYPES`, `PROFESSIONS`, `REP_LEVELS`; `pipeline.loot.sources.KIND_ORDER`
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
from pipeline.loot.sources import KIND_ORDER

BUILD = "1.60.1.69893"
BUILD_DIR = Path("builds") / BUILD

#: Sources per kind. Seven raids (Molten Core, Blackwing Lair, Onyxia's
#: Lair, Zul'Gurub, both Ahn'Qiraj zones and Naxxramas -- the last three
#: found only by the zones.json fallback join, since the fork's own zone
#: list omits them), nineteen dungeons, two world bosses (Azuregos and
#: Lord Kazzak, the only open-world drop NPCs the fork names), five
#: professions, thirty-one faction-and-standing pairs, thirteen PvP ranks,
#: one quest list.
SOURCES_PER_KIND = {
    "raid": 7,
    "dungeon": 19,
    "world": 2,
    "crafted": 5,
    "rep": 31,
    "pvp": 13,
    "quest": 1,
}

RAID_SOURCE_IDS = [
    "raid:ahnqiraj",
    "raid:blackwing-lair",
    "raid:molten-core",
    "raid:naxxramas",
    "raid:onyxias-lair",
    "raid:ruins-of-ahnqiraj",
    "raid:zulgurub",
]
#: raid source id -> (bosses, trash items, distinct items).
RAID_SHAPE = {
    "raid:ahnqiraj": (12, 14, 111),
    "raid:blackwing-lair": (9, 21, 132),
    "raid:molten-core": (11, 45, 139),
    "raid:naxxramas": (15, 12, 111),
    "raid:onyxias-lair": (1, 0, 16),
    "raid:ruins-of-ahnqiraj": (6, 10, 61),
    "raid:zulgurub": (13, 18, 98),
}
RAID_BOSSES = 67
RAID_ITEMS = 668
#: Bosses the fork database names no NPC for. An invented name would be
#: worse than a blank one, so this is measured rather than forbidden.
UNNAMED_RAID_BOSSES = 35
UNNAMED_DUNGEON_BOSSES = 39

DUNGEON_BOSSES = 230
DUNGEONS_WITH_TRASH = 17
WORLD_SOURCE_IDS = ["world:azuregos", "world:lord-kazzak"]
CRAFTED_ITEMS = {
    "crafted:blacksmithing": 195,
    "crafted:enchanting": 5,
    "crafted:engineering": 48,
    "crafted:leatherworking": 191,
    "crafted:tailoring": 157,
}
QUEST_ITEMS = 1228
PVP_ITEMS_PER_RANK = {5: 2, 6: 16, 7: 6, 8: 6, 9: 23, 10: 2, 11: 66, 12: 93,
                      14: 63, 15: 2, 16: 80, 17: 48, 18: 42}

#: Every item id the file names, and how many of them this client's own
#: item table does not have. The gap is Forever re-itemising on top of the
#: fork's Era database; curated/loot/forever-raid-phases.json carries the
#: measurement in prose.
NAMED_ITEMS = 4316
NAMED_ITEMS_NOT_IN_BUILD = 1809

ENCHANT_ROWS = 173
ENCHANT_EFFECT_IDS = 150
SUFFIX_ROWS = 1168
ITEMS_WITH_SUFFIXES = 69

PHASES = {"pre-beta", "beta", "launch", "raids-1"}


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
    assert sum(SOURCES_PER_KIND.values()) == len(loot()["sources"]) == 78


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
        assert len(source["bosses"]) == bosses, source_id
        assert len(source.get("trash", [])) == trash, source_id
        assert len(source_items(source)) == distinct, source_id
    assert sum(shape[0] for shape in RAID_SHAPE.values()) == RAID_BOSSES
    assert len({
        item for s in loot()["sources"] if s["kind"] == "raid" for item in source_items(s)
    }) == RAID_ITEMS


def test_a_boss_without_a_name_is_blank_and_counted_not_invented():
    for kind, expected in (("raid", UNNAMED_RAID_BOSSES), ("dungeon", UNNAMED_DUNGEON_BOSSES)):
        blank = [
            boss
            for source in loot()["sources"]
            if source["kind"] == kind
            for boss in source["bosses"]
            if boss["name"] == ""
        ]
        assert len(blank) == expected, kind


def test_every_boss_id_is_its_source_id_plus_its_npc_id():
    for source in loot()["sources"]:
        for boss in source.get("bosses", []):
            assert boss["id"] == f"{source['id']}:{boss['npc_id']}"
            assert boss["npc_id"] > 0


def test_the_dungeon_sources_are_the_nineteen_measured():
    dungeons = [s for s in loot()["sources"] if s["kind"] == "dungeon"]
    assert len(dungeons) == SOURCES_PER_KIND["dungeon"]
    assert sum(len(s["bosses"]) for s in dungeons) == DUNGEON_BOSSES
    assert sum(1 for s in dungeons if s.get("trash")) == DUNGEONS_WITH_TRASH


def test_the_only_world_sources_are_the_two_named_world_bosses():
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


def test_every_raid_is_gated_behind_a_real_phase_and_nothing_else_is():
    for source in loot()["sources"]:
        opens = source.get("opens")
        if source["kind"] == "raid":
            assert opens == "raids-1", source["id"]
        else:
            assert opens is None, source["id"]
        assert opens is None or opens in PHASES


def test_the_coverage_against_this_builds_item_table_is_what_was_measured():
    named = {item for source in loot()["sources"] for item in source_items(source)}
    assert len(named) == NAMED_ITEMS
    assert len(named - {row["id"] for row in items()}) == NAMED_ITEMS_NOT_IN_BUILD


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


def test_items_json_carries_the_suffix_column_on_every_row():
    assert all("suffixes" in row for row in items())
    assert sum(1 for row in items() if row["suffixes"]) == ITEMS_WITH_SUFFIXES
```

- [ ] **Step 2: Run the test**

Run: `cd data && uv run pytest tests/test_loot_build.py -q --no-cov`
Expected: PASS, 18 tests

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
than regenerated. 78 sources, 7 raids with 67 bosses over 668 items,
19 dungeons, 2 world bosses, 173 enchants, 1,168 suffixes, and the
4,316/1,809 coverage gap against this client's item table.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 11: The Makefile target and the CI job

How the existing build files are handled, and what this matches:

- **Regeneration** happens in `data.yml`'s `fetch` job, which only runs on `workflow_dispatch`, downloads `raw/`, runs each per-build emitter, and commits `builds/`. `loot` joins that list.
- **Checking** happens in `data.yml`'s `test` job, which runs `ruff`, `specs --check` and `pytest` — and pytest is where `test_simdb_build.py`, `test_gametables_build.py` and now `test_loot_build.py` hold the committed files to their measurements. Nothing is regenerated there, because the test job has neither `raw/` nor an engine checkout.
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
	@(cd data && uv run pytest tests/test_loot_build.py tests/test_loot_overlay.py -q --no-cov)
```

- [ ] **Step 2: Verify the targets**

Run: `make loot-check`
Expected: PASS

Run: `make loot ENGINE_DIR=/nonexistent`
Expected: exits non-zero with `no item database at /nonexistent/assets/database/db.json; set ENGINE_DIR`

- [ ] **Step 3: Add the regeneration step to the workflow**

In `.github/workflows/data.yml`, in the `fetch` job, insert after the `gametables` step and before the `diff` step:

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
      # After simdb, because it rewrites items.json and refreshes the same
      # manifest; before diff, so the diff sees the finished directory.
      - run: uv run python -m pipeline loot --build "$BUILD" --engine ../.engine
```

Also extend the workflow's `paths` filters (both `push` and `pull_request`) with `'sim/enginever/version.go'`, so a pin bump runs the data tests:

```yaml
    paths: ['data/**', 'sim/specs/**', 'sim/enginever/version.go', 'web/src/lib/sim/specs.ts', '.github/workflows/data.yml']
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

`make loot` regenerates the active build's loot, enchant and suffix
files from an engine checkout; `make loot-check` is the offline gate on
what it wrote.

CI matches what simdb.bin and gametables/ already do: the fetch job
regenerates (cloning the fork at the pinned sha the way sim.yml's
rotations job does) and the test job checks the committed files through
pytest, because it has neither raw/ nor an engine checkout. A pin bump
now runs the data tests.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

## Self-review

**Spec coverage.** Contract 6.1 `loot.json` — tasks 6, 7, 8, 9. Contract 6.2 `enchants.json` — task 5. Contract 6.3 `suffixes.json` and `items.json`'s `suffixes` — task 4. Contract 6.4 the socket guard — task 1. Contract 1.4 and 8's `reference_stat` — task 2. Design 6.2's "joined to the build's zones" — task 6's two-step join. Design 12's ordering (dungeon and crafted sources before launch, raid sources with the phase gate before 9 December) — every kind lands together in task 9, ahead of both dates. Design 13's "Forever-only items with no source" — task 8's overlay states the measured gap. The reading layer the contract assumes but does not name (the fork database, `statmap`'s inverse) — task 3. Tests, Makefile and CI — tasks 10 and 11.

Not in this lane, deliberately: `sim/bulk`, the wasm exports, migration 0014, the web routes, the addon export. `dungeons.json` stays empty — `JournalInstance` 404s on **both** Classic-lineage products (checked live against wago.tools on 2026-09-19 for 1.60.1.69893 and for Era 1.15.9.69722), so the Dungeon Journal offers this lane nothing; `loot.json` takes its dungeon and raid list from `Map.InstanceType` instead, which is why no task tries to revive it.

**Type consistency.** `LootSource` / `LootBoss` / `LootFile` are defined in task 6 and used unchanged in 7, 8, 9 and 10. `ForkDatabase`'s field names (`random_suffixes`, `item_icons`, `spell_icons`) are fixed in task 3 and read in 4, 5, 6 and 9. `build_loot` returns `(LootFile, LootStats)` in task 6 and is unpacked that way in task 9. `KIND_ORDER` is defined in task 6 and imported by task 7's `apply_overlays` and task 10's ordering test. `parse_sources` is renamed in task 7 and called only from `curated.py` and `overlay.py`. The three filename constants `LOOT` / `ENCHANTS` / `SUFFIXES` live in `pipeline/loot/__init__.py` (task 9); task 10's test spells the filenames out, on purpose, because a conformance test that imported them would only prove the pipeline agrees with itself.

## Contract ambiguities and errors found

Each is resolved as described in the task that hits it; they are collected here so the contract can be amended in its own commit.

1. **`zone_id: 409` for Molten Core is a map id, not a zone id.** The fork database and the build's `zones.json` both key Molten Core as AreaTable 2717; 409 is its `Map.ID`, and AreaTable 409 is "Island of Doctor Lapidis". The plan emits 2717, so the join to `zones.json` works.
2. **`ItemRandomSuffix`, contract 6.3's stated source, 404s on build 1.60.1.69893** — the same absence that already leaves `SimDatabase.random_suffixes` empty — so suffixes come from the fork database's `randomSuffixes` and per-item `randomSuffixOptions`.
3. **`UIEnchant.effect_id` is not unique** (173 rows over 150 ids), so `id` alone cannot key `enchants.json`; the emitted rows also carry `spell_id` and `item_id`.
4. **`slots` and `item_types` in contract 6.2 would be the same field** — both can only derive from `UIEnchant.type` plus `extra_types` — so `item_types` is read as the `EnchantType` shape restriction (normal, two_hand, shield, kit, staff) instead.
5. **Contract 6.1's id examples (`raid:mc`, `raid:mc:lucifron`) are not derivable**: nothing in either database gives a short code, and 74 of the 297 bosses have no name at all, so ids are `raid:<zone-slug>` and `raid:<zone-slug>:<npc-id>`.
6. **A boss's `name` is not always available** — 35 raid and 39 dungeon bosses are unnamed in the fork database — so it may be empty rather than invented.
7. **The `pvp` kind has no source in either database**: the fork carries no rank at all, so it is built from the client's `ItemSparse.RequiredPVPRank` (449 items over 13 ranks).
8. **The contract has no kind for a plain vendor or an open-world trash drop**, which together are 864 of the fork's source entries; they are dropped and counted rather than forced into a kind.
9. **`opens` cannot distinguish one raid tier from another**: `api/internal/phase/phase.go` has exactly one raid phase, so Ahn'Qiraj and Naxxramas share `raids-1` with Molten Core.
10. **The overlay's "same shape plus `sources` and `notes`" collides with itself** — `sources` is both contract 6.1's key for loot sources and every curated file's provenance key — so the overlay keeps `sources`/`notes` as provenance and nests the loot-shaped documents under `add` and `replace`.
11. **Contract 1.4 cites "IDS.md stat ids", but `sim/request/IDS.md` has no stat section** and nothing anywhere defines a plain `haste`; the stat vocabulary that exists is `pipeline/simdb/statmap.py`'s `PROTO_STAT_ALIASES`, which is what `reference_stat` is validated against.
