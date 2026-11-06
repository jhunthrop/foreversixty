"""The first raid tier's bosses and their loot, tied to the client's own rows.

Written to `builds/<build>/raid-bosses.json` for the guild loot picker and the raid
assignment templates (`api/internal/guilds/raid_bosses_table.go` is generated from it by
`go generate ./internal/guilds/` in `api/`). The raids are the zones of
`curated/raids/first-tier.json`; this step adds what the client says about each one.

**Ties are by the client's own names and ids, never by guess.** A raid is in the client
when a `Map` row (`MapName_lang`) or a `JournalInstance` row (`Name_lang`) carries the
announced name exactly; its bosses are the `DungeonEncounter` rows on those maps (and the
`JournalEncounter` rows of that instance). A map whose name only resembles the announced
one (2817 Starfall Barrow Den, 2995 Hyjal Crater) is not a tie. A raid the client does not
name is written with no bosses and the status `not_in_client`.

**Loot** is the boss's `JournalEncounterItem` rows where the client ships the Dungeon
Journal. 1.60.1.70790 does not (wago answers 400 for all three Journal tables), so a boss
with no journal loot takes the drops of the build's own `loot.json` raid source of the
same name (classic-db's creature loot, already filtered to this client) when that source
names a boss with the encounter's exact name; `loot_source` says which. Every item must be
in the build's `items.json`; the ones that are not are counted under `dropped_items`.

**Sources** (all from wago.tools for the build; `fetch` writes them to `raw/`): `Map`,
`MapDifficulty` (the raid's player count), `DungeonEncounter`, `JournalInstance`,
`JournalEncounter`, `JournalEncounterItem`.
"""

from __future__ import annotations

import json
import logging
from pathlib import Path

from pipeline.csvio import read_csv

logger = logging.getLogger(__name__)

RAID_BOSSES = "raid-bosses.json"
FIRST_TIER = Path("raids/first-tier.json")
REQUIRED_TABLES = ("Map", "MapDifficulty", "DungeonEncounter")
JOURNAL_TABLES = ("JournalInstance", "JournalEncounter", "JournalEncounterItem")
TABLES = REQUIRED_TABLES + JOURNAL_TABLES

IN_CLIENT = "in_client"
NOT_IN_CLIENT = "not_in_client"

#: The items.json fields each loot row carries, for the API's loot table.
ITEM_FIELDS = (
    "name",
    "quality",
    "item_level",
    "required_level",
    "class_id",
    "subclass_id",
    "inventory_type",
)


class RaidBossesError(SystemExit):
    """A raid-bosses input is missing or malformed."""


def _int(row: dict[str, str], column: str) -> int:
    value = row.get(column, "")
    try:
        return int(value)
    except ValueError as exc:
        raise RaidBossesError(f"column {column} is not an integer: {value!r} in {row}") from exc


def _rows(raw: Path, table: str) -> list[dict[str, str]]:
    path = raw / f"{table}.csv"
    if not path.exists():
        if table in JOURNAL_TABLES:
            return []
        raise RaidBossesError(f"no {path}; run `python -m pipeline fetch` first")
    return read_csv(path)


def _row_ref(table: str, row: dict[str, str], name_column: str | None) -> dict:
    ref: dict = {"table": table, "id": _int(row, "ID")}
    if name_column:
        ref["name"] = row.get(name_column, "")
    return ref


class ClientTables:
    """The raw client tables this step reads, indexed for the ties it makes."""

    def __init__(self, raw: Path) -> None:
        self.maps = _rows(raw, "Map")
        self.difficulties = _rows(raw, "MapDifficulty")
        self.encounters = _rows(raw, "DungeonEncounter")
        self.instances = _rows(raw, "JournalInstance")
        self.journal_encounters = _rows(raw, "JournalEncounter")
        self.journal_items = _rows(raw, "JournalEncounterItem")
        self.has_journal = bool(self.journal_encounters)

    def maps_named(self, name: str) -> list[dict[str, str]]:
        return [r for r in self.maps if r.get("MapName_lang") == name]

    def instances_for(self, name: str, map_ids: set[int]) -> list[dict[str, str]]:
        return [
            r
            for r in self.instances
            if r.get("Name_lang") == name or (r.get("MapID") and _int(r, "MapID") in map_ids)
        ]

    def difficulties_on(self, map_ids: set[int]) -> list[dict[str, str]]:
        return [r for r in self.difficulties if _int(r, "MapID") in map_ids]

    def encounters_on(self, map_ids: set[int]) -> list[dict[str, str]]:
        rows = [r for r in self.encounters if _int(r, "MapID") in map_ids]
        return sorted(rows, key=lambda r: (_int(r, "OrderIndex"), _int(r, "ID")))

    def journal_encounter_for(
        self, encounter_id: int, instance_ids: set[int]
    ) -> dict[str, str] | None:
        for r in self.journal_encounters:
            if r.get("DungeonEncounterID") and _int(r, "DungeonEncounterID") == encounter_id:
                if not instance_ids or _int(r, "JournalInstanceID") in instance_ids:
                    return r
        return None

    def journal_item_ids(self, journal_encounter_id: int) -> list[int]:
        return sorted(
            {
                _int(r, "ItemID")
                for r in self.journal_items
                if _int(r, "JournalEncounterID") == journal_encounter_id
            }
        )


def _loot_boss_items(loot: dict, raid_name: str, boss_name: str) -> tuple[list[int], str] | None:
    """The drops loot.json's raid source `raid_name` names for a boss `boss_name`."""
    for source in loot.get("sources", []):
        if source.get("kind") != "raid" or source.get("name") != raid_name:
            continue
        for boss in source.get("bosses", []):
            if boss.get("name") == boss_name:
                origin = source.get("source_origin", "unknown")
                label = f"loot.json {boss['id']} (npc {boss.get('npc_id')}, {origin})"
                return sorted(set(boss.get("items", []))), label
    return None


def _loot_rows(item_ids: list[int], items: dict[int, dict]) -> tuple[list[dict], list[int]]:
    kept = [
        {"item_id": i, **{f: items[i].get(f) for f in ITEM_FIELDS}} for i in item_ids if i in items
    ]
    dropped = [i for i in item_ids if i not in items]
    return kept, dropped


def _boss(
    tables: ClientTables,
    encounter: dict[str, str],
    order: int,
    raid_name: str,
    instance_ids: set[int],
    items: dict[int, dict],
    loot: dict,
) -> dict:
    encounter_id = _int(encounter, "ID")
    name = encounter.get("Name_lang", "")
    rows = [_row_ref("DungeonEncounter", encounter, "Name_lang")]
    journal = tables.journal_encounter_for(encounter_id, instance_ids)
    item_ids: list[int] = []
    loot_source = None
    if journal is not None:
        rows.append(_row_ref("JournalEncounter", journal, "Name_lang"))
        item_ids = tables.journal_item_ids(_int(journal, "ID"))
        if item_ids:
            loot_source = f"JournalEncounterItem (JournalEncounterID {_int(journal, 'ID')})"
    if not item_ids:
        found = _loot_boss_items(loot, raid_name, name)
        if found:
            item_ids, loot_source = found
    kept, dropped = _loot_rows(item_ids, items)
    return {
        "encounter_id": encounter_id,
        "journal_encounter_id": _int(journal, "ID") if journal is not None else None,
        "name": name,
        "order": order,
        "rows": rows,
        "loot_source": loot_source,
        "loot": kept,
        "dropped_items": dropped,
    }


def _raid(zone: dict, tables: ClientTables, build: str, items: dict, loot: dict) -> dict:
    name = zone["name"]
    maps = tables.maps_named(name)
    map_ids = {_int(r, "ID") for r in maps}
    instances = tables.instances_for(name, map_ids)
    map_ids |= {_int(r, "MapID") for r in instances if r.get("MapID")}
    instance_ids = {_int(r, "ID") for r in instances}
    rows = [_row_ref("Map", r, "MapName_lang") for r in maps]
    rows += [_row_ref("JournalInstance", r, "Name_lang") for r in instances]
    difficulties = tables.difficulties_on(map_ids)
    rows += [
        {**_row_ref("MapDifficulty", r, None), "max_players": _int(r, "MaxPlayers")}
        for r in difficulties
    ]
    encounters = tables.encounters_on(map_ids)
    bosses = [
        _boss(tables, e, i + 1, name, instance_ids, items, loot) for i, e in enumerate(encounters)
    ]
    in_client = bool(maps or instances)
    note = None
    if not in_client:
        note = f"not in client {build}: no Map or JournalInstance row is named {name!r}"
    elif not tables.has_journal:
        note = f"client {build} ships no Dungeon Journal; loot is the build's loot.json source"
    return {
        "key": zone["key"],
        "name": name,
        "status": IN_CLIENT if in_client else NOT_IN_CLIENT,
        "note": note,
        "max_players": max((_int(r, "MaxPlayers") for r in difficulties), default=None),
        "rows": rows,
        "bosses": bosses,
    }


def _first_tier_zones(curated_dir: Path) -> list[dict]:
    path = curated_dir / FIRST_TIER
    zones = json.loads(path.read_text(encoding="utf-8")).get("zones", [])
    raids = [z for z in zones if z.get("kind") == "raid"]
    if not raids:
        raise RaidBossesError(f"{path} lists no raid zones")
    return raids


def build_raid_bosses(build: str, build_dir: Path, raw: Path, curated_dir: Path) -> dict:
    """The raid-bosses.json document for `build`."""
    items_path = build_dir / "items.json"
    if not items_path.exists():
        raise RaidBossesError(f"no {items_path}; run `python -m pipeline normalize` first")
    items = {i["id"]: i for i in json.loads(items_path.read_text(encoding="utf-8"))}
    loot_path = build_dir / "loot.json"
    loot = json.loads(loot_path.read_text(encoding="utf-8")) if loot_path.exists() else {}
    tables = ClientTables(raw)
    return {
        "build": build,
        "source": [f"https://wago.tools/db2/{t}/csv?build={build}" for t in TABLES],
        "journal_in_client": tables.has_journal,
        "raids": [_raid(z, tables, build, items, loot) for z in _first_tier_zones(curated_dir)],
    }


def render(doc: dict) -> str:
    return json.dumps(doc, indent=2) + "\n"


def write_raid_bosses(
    build: str,
    root: Path = Path("builds"),
    curated_dir: Path = Path("curated"),
    raw: Path | None = None,
    check: bool = False,
) -> Path:
    """Write (or, with `check`, verify) `builds/<build>/raid-bosses.json`."""
    build_dir = root / build
    doc = build_raid_bosses(build, build_dir, raw or build_dir / "raw", curated_dir)
    out = build_dir / RAID_BOSSES
    text = render(doc)
    if check:
        if not out.exists() or out.read_text(encoding="utf-8") != text:
            raise RaidBossesError(
                f"{out} has drifted; run `python -m pipeline raid-bosses --build {build}`"
            )
        return out
    out.write_text(text, encoding="utf-8")
    for raid in doc["raids"]:
        logger.info("%s: %s, %d bosses", raid["name"], raid["status"], len(raid["bosses"]))
    return out
