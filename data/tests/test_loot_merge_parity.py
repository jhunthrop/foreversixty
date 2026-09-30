"""loot-parity lane, 2026-09-30: `write_loot_files` (`loot`) and
`merge_loot_files` (`loot-merge`) already share one builder
(`pipeline.loot.sources.build_loot`) -- the gap this lane's own report
measured on build 1.60.1.70009 was never a second, unshared build
function. It was `merge_loot_files`'s OWN `types_from_committed_loot`
recovering an incomplete `types` (zone id -> dungeon/raid code) next to
`write_loot_files`'s raw-`Map.csv` `instance_types()`: a published
`LootSource.zone_id` is only ever the fork's OWN AtlasLoot-era zone id
for an instance, never the second, unpublished `zones.json` row for the
SAME instance that carries its real DB2 `Map.ID`-joined zone id (and the
real `map_id` `pipeline.loot.classicdb.instance_zone_by_map` needs to
place a classic-db-only trash mob inside it). Measured regression this
closed: Blackrock Depths 292 -> 143 items, Blackrock Spire 296 -> 106,
on that build -- see `pipeline/loot/__init__.py::types_from_committed_
loot`'s own doc for the fix.

This module is the regression guard: `test_types_from_committed_loot.py`
-- style unit tests for the fix's own name-based backfill (and the
false-positive it must not reintroduce, Westfall's own real open-world
zone id sharing its literal name with The Deadmines' third, mislabeled
`zones.json` row), plus one end-to-end test that runs BOTH `loot` and
`loot-merge` over the same small fixture build, with a synthetic
"Test Depths" dungeon shaped exactly like the BRD/BRS bug (a published
zone id with `map_id: 0` next to an unpublished one with the real
`map_id`, and a classic-db-only trash npc the fork's own database never
names at all), and asserts the two commands' `sources` agree.
"""

from __future__ import annotations

import json
import shutil
from pathlib import Path

from pipeline.classic_sources import ClassicDbSourceRecord, write_classic_sources
from pipeline.loot import LOOT, merge_loot_files, types_from_committed_loot, write_loot_files
from pipeline.manifest import write_manifest

HERE = Path(__file__).parent
FIXTURES = HERE / "fixtures/loot"
OVERLAY_FIXTURE = FIXTURES / "curated"
BUILD = "1.60.1.69893"

RAW_FILES = ("Map.csv", "ItemSparse.csv", "SpellMisc.csv", "Item.csv",
             "ManifestInterfaceData.csv")
JSON_FILES = ("zones.json", "items.json", "spells.json")


# ---------------------------------------------------------------------------
# types_from_committed_loot: the fix itself, isolated from any build at all.
# ---------------------------------------------------------------------------

_COMMITTED = {
    "sources": [
        {"id": "dungeon:blackrock-depths", "kind": "dungeon", "name": "Blackrock Depths",
         "zone_id": 1584, "bosses": [{"id": "x:1", "name": "A Boss", "npc_id": 1,
                                       "items": [1]}]},
        {"id": "raid:molten-core", "kind": "raid", "name": "Molten Core", "zone_id": 2717,
         "bosses": [{"id": "y:1", "name": "Lucifron", "npc_id": 2, "items": [2]}]},
    ]
}

#: The exact shape build 1.60.1.70009 carries: the published, fork-used
#: zone id (`map_id: 0`, a bare continent -- `instance_zone_by_map`'s own
#: exclusion), the real DB2 instance zone id sharing its name (`map_id:
#: 230`, what a classic-db creature record's own `map_id` column needs to
#: join against), and Westfall's own true open-world zone -- a real,
#: UNRELATED zone that happens to share its literal name with The
#: Deadmines' own third, mislabeled `zones.json` row (id 206 on the real
#: build; modelled here as sharing "Molten Core" instead, any name would
#: do) and must never be swept in by a name-only match.
_ZONE_ROWS = [
    {"id": 1584, "name": "Blackrock Depths", "map_id": 0},
    {"id": 17803, "name": "Blackrock Depths", "map_id": 230},
    {"id": 2717, "name": "Molten Core", "map_id": 409},
    {"id": 40, "name": "Molten Core", "map_id": 0},  # decoy: real open-world zone, same name
]


def test_the_published_zone_id_alone_is_recovered_without_zone_rows(tmp_path: Path):
    build_dir = tmp_path / "build"
    build_dir.mkdir()
    (build_dir / LOOT).write_text(json.dumps(_COMMITTED), encoding="utf-8")
    assert types_from_committed_loot(build_dir) == {1584: 1, 2717: 2}


def test_a_same_named_zone_row_with_a_real_map_id_is_backfilled(tmp_path: Path):
    """The BRD/BRS fix: `17803` never became a `LootSource.zone_id`
    itself, but shares Blackrock Depths' name and carries a real
    (non-continent) `map_id` -- exactly what `instance_zone_by_map`
    needs and could not otherwise ever learn from committed data alone."""
    build_dir = tmp_path / "build"
    build_dir.mkdir()
    (build_dir / LOOT).write_text(json.dumps(_COMMITTED), encoding="utf-8")
    types = types_from_committed_loot(build_dir, _ZONE_ROWS)
    assert types[17803] == 1  # dungeon, same code as the published 1584
    assert types[1584] == 1
    assert types[2717] == 2


def test_a_same_named_zone_row_with_a_bare_continent_map_id_is_never_backfilled(tmp_path: Path):
    """The false positive the fix must not reintroduce: zone 40 shares
    Molten Core's name but is a real, unrelated open-world zone (`map_id:
    0`) -- promoting it to "raid" would fold every one of the fork's own
    open-world drops there into the raid bucket (measured on the real
    build as `dungeon:westfall` inflating 125 -> 191 items before this
    exclusion existed)."""
    build_dir = tmp_path / "build"
    build_dir.mkdir()
    (build_dir / LOOT).write_text(json.dumps(_COMMITTED), encoding="utf-8")
    types = types_from_committed_loot(build_dir, _ZONE_ROWS)
    assert 40 not in types


def test_a_zone_with_currently_zero_sourced_items_stays_unrecovered(tmp_path: Path):
    """Documented degradation (the function's own doc): a dungeon/raid
    zone absent from the committed sources entirely -- not even under a
    same-named duplicate -- is invisible to this recovery, same as
    before the fix."""
    build_dir = tmp_path / "build"
    build_dir.mkdir()
    (build_dir / LOOT).write_text(json.dumps(_COMMITTED), encoding="utf-8")
    types = types_from_committed_loot(
        build_dir, [*_ZONE_ROWS, {"id": 9999, "name": "Uldaman", "map_id": 70}]
    )
    assert 9999 not in types


# ---------------------------------------------------------------------------
# End-to-end: `loot` and `loot-merge` over the same small fixture build,
# with a synthetic BRD/BRS-shaped dungeon.
# ---------------------------------------------------------------------------

#: cmangos map id for the synthetic "Test Depths" instance -- what the
#: classic-db-only trash npc's own record states, and what `zones.json`'s
#: real (unpublished) duplicate row must carry for `instance_zone_by_map`
#: to resolve it.
_TEST_DEPTHS_MAP_ID = 500


def _augmented_engine(tmp_path: Path) -> Path:
    """A copy of the fixture engine checkout with one extra fork-known
    boss: npc 910 in the synthetic "Test Depths" dungeon (published zone
    id 5001), dropping item 5011."""
    engine_dir = tmp_path / "engine"
    shutil.copytree(FIXTURES, engine_dir)
    db_path = engine_dir / "assets" / "database" / "db.json"
    db = json.loads(db_path.read_text(encoding="utf-8"))
    db["items"].append(
        {
            "id": 5011,
            "name": "Test Depths Boss Drop",
            "icon": "inv_f",
            "type": 4,
            "ilvl": 60,
            "phase": 1,
            "quality": 3,
            "sources": [{"drop": {"difficulty": 1, "npcId": 910, "zoneId": 5001}}],
        }
    )
    db["npcs"].append({"id": 910, "name": "Test Depths Boss", "zoneId": 5001})
    db["zones"].append({"id": 5001, "name": "Test Depths", "expansion": 1})
    db_path.write_text(json.dumps(db), encoding="utf-8")
    return engine_dir


def _prepare_build(tmp_path: Path) -> Path:
    """`tests/test_loot.py::prepare_build`'s own shape, plus the
    synthetic dungeon's zone/map/item rows and a classic-db-only trash
    npc the fork's database above never names."""
    root = tmp_path / "builds"
    build_dir = root / BUILD
    raw = build_dir / "raw"
    raw.mkdir(parents=True)
    for name in RAW_FILES:
        shutil.copy(FIXTURES / name, raw / name)
    for name in JSON_FILES:
        shutil.copy(FIXTURES / name, build_dir / name)
    write_manifest(
        build_dir, build=BUILD, product="wow_classic_beta", fetched_at="2026-01-01T00:00:00Z"
    )

    zones = json.loads((build_dir / "zones.json").read_text(encoding="utf-8"))
    zones += [
        # Published: what the fork's own drop.zoneId names -- a bare
        # continent map id, exactly Blackrock Depths' real shape.
        {"id": 5001, "name": "Test Depths", "map_id": 0, "map_name": "Eastern Kingdoms",
         "parent_id": None, "level_min": None, "level_max": None},
        # Unpublished: the real DB2 instance zone, never a
        # LootSource.zone_id anywhere -- only `instance_types()` (raw
        # Map.csv) or this lane's committed-data backfill ever learns it.
        {"id": 5002, "name": "Test Depths", "map_id": _TEST_DEPTHS_MAP_ID,
         "map_name": "Test Depths", "parent_id": None, "level_min": None, "level_max": None},
    ]
    (build_dir / "zones.json").write_text(json.dumps(zones), encoding="utf-8")

    items = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    items += [
        {"id": 5010, "name": "Test Depths Relic", "quality": 3, "item_level": 60,
         "required_level": 55, "class_id": 4, "subclass_id": 4, "inventory_type": 5},
        {"id": 5011, "name": "Test Depths Boss Drop", "quality": 3, "item_level": 60,
         "required_level": 55, "class_id": 4, "subclass_id": 4, "inventory_type": 5},
    ]
    (build_dir / "items.json").write_text(json.dumps(items), encoding="utf-8")

    map_csv = (build_dir / "raw" / "Map.csv").read_text(encoding="utf-8")
    map_csv += f"{_TEST_DEPTHS_MAP_ID},TestDepths,Test Depths,1,5001\n"
    (build_dir / "raw" / "Map.csv").write_text(map_csv, encoding="utf-8")

    # The classic-db-only trash mob: the fork's own database (`_augmented_
    # engine`) never names npc 911 at all, so only `instance_zone_by_map`
    # -- keyed by this record's own `map_id` -- can place item 5010
    # inside the dungeon instead of a flat `world:test-depths-trash`.
    write_classic_sources(
        build_dir,
        {
            5010: [
                ClassicDbSourceRecord(
                    kind="creature_drop", npc_id=911, name="Test Depths Trash",
                    map_id=_TEST_DEPTHS_MAP_ID, chance=10.0,
                )
            ]
        },
    )
    return root


def _dungeon_items(document: dict, source_id: str) -> set[int]:
    for source in document["sources"]:
        if source["id"] == source_id:
            items = set(source.get("items") or [])
            items.update(source.get("trash") or [])
            for boss in source.get("bosses") or []:
                items.update(boss.get("items") or [])
            return items
    raise AssertionError(f"no source {source_id!r} in {[s['id'] for s in document['sources']]}")


def test_loot_places_the_classicdb_only_trash_mob_inside_the_dungeon(tmp_path: Path):
    """Baseline: `write_loot_files`, with `raw/Map.csv` available, already
    gets this right -- both the fork's own boss (5011) and classic-db's
    rescued trash mob (5010) land in `dungeon:test-depths`."""
    root = _prepare_build(tmp_path)
    engine_dir = _augmented_engine(tmp_path)
    curated = tmp_path / "curated"
    curated.mkdir()
    shutil.copy(FIXTURES / "curated-simbuffs.json", curated / "simbuffs.json")
    write_loot_files(
        BUILD, engine_dir=engine_dir, root=root, overlay_dir=tmp_path / "no-overlay",
        curated_dir=curated, ids_md=FIXTURES / "IDS.md",
    )
    document = json.loads((root / BUILD / LOOT).read_text(encoding="utf-8"))
    assert _dungeon_items(document, "dungeon:test-depths") == {5010, 5011}


def test_loot_merge_agrees_with_loot_on_the_same_committed_caches(tmp_path: Path):
    """The actual parity contract: run `loot` once (so `raw/classicdb/
    sources.json`, `zones.json`, `items.json` and `loot.json` itself are
    all committed the way a real nightly leaves them), then run
    `loot-merge` alone over that SAME build dir -- no raw/Map.csv or
    ItemSparse.csv read a second time -- and assert its own
    `dungeon:test-depths` names the exact same items. Before this lane's
    fix, `loot-merge` dropped item 5010 to `world:test-depths-trash`
    instead (`types_from_committed_loot` never recovering zone id 5002's
    real `map_id`)."""
    root = _prepare_build(tmp_path)
    engine_dir = _augmented_engine(tmp_path)
    curated = tmp_path / "curated"
    curated.mkdir()
    shutil.copy(FIXTURES / "curated-simbuffs.json", curated / "simbuffs.json")
    write_loot_files(
        BUILD, engine_dir=engine_dir, root=root, overlay_dir=tmp_path / "no-overlay",
        curated_dir=curated, ids_md=FIXTURES / "IDS.md",
    )
    full_document = json.loads((root / BUILD / LOOT).read_text(encoding="utf-8"))

    merge_loot_files(BUILD, engine_dir=engine_dir, root=root, overlay_dir=tmp_path / "no-overlay")
    merged_document = json.loads((root / BUILD / LOOT).read_text(encoding="utf-8"))

    assert _dungeon_items(merged_document, "dungeon:test-depths") == {5010, 5011}

    def by_id(document: dict) -> dict[str, set[int]]:
        return {
            source["id"]: _dungeon_items(document, source["id"])
            for source in document["sources"]
        }

    assert by_id(merged_document) == by_id(full_document)
