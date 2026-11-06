"""`write_loot_files` (`loot`) and `merge_loot_files` (`loot-merge`) must
write the same loot.json from the same inputs.

Both build it through one function (`pipeline.loot._build_and_write_loot`)
from one classification of the client's Map.csv and ItemSparse.csv, which
`loot` writes to `raw/client/loot-classification.json` and the merge reads.
Before round 8 the merge recovered instance types and PvP ranks from the
committed loot.json instead: first it lost Blackrock Depths' real instance
zone (loot-parity lane, 2026-09-30, the "Test Depths" fixture below), then it
disagreed about Alterac Valley and Arathi Basin and a re-itemised PvP tabard,
so loot.json flipped on every pair of runs (round 8, owner ruling 3).
"""

from __future__ import annotations

import json
import os
import shutil
from pathlib import Path

import pytest
from test_loot import prepare_build, prepare_curated

from pipeline.classic_sources import ClassicDbSourceRecord, write_classic_sources
from pipeline.loot import LOOT, merge_loot_files, merge_simbuffs_file, write_loot_files
from pipeline.loot.buffs import SIMBUFFS
from pipeline.manifest import read_manifest

HERE = Path(__file__).parent
FIXTURES = HERE / "fixtures/loot"
OVERLAY_FIXTURE = FIXTURES / "curated"
BUILD = "1.60.1.69893"


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
    """`tests/test_loot.py::prepare_build`, plus the
    synthetic dungeon's zone/map/item rows and a classic-db-only trash
    npc the fork's database above never names."""
    root = prepare_build(tmp_path)
    build_dir = root / BUILD

    zones = json.loads((build_dir / "zones.json").read_text(encoding="utf-8"))
    zones += [
        # Published: what the fork's own drop.zoneId names -- a bare
        # continent map id, exactly Blackrock Depths' real shape.
        {
            "id": 5001,
            "name": "Test Depths",
            "map_id": 0,
            "map_name": "Eastern Kingdoms",
            "parent_id": None,
            "level_min": None,
            "level_max": None,
        },
        # Unpublished: the real DB2 instance zone, never a
        # LootSource.zone_id anywhere -- only `instance_types()` (raw
        # Map.csv) or this lane's committed-data backfill ever learns it.
        {
            "id": 5002,
            "name": "Test Depths",
            "map_id": _TEST_DEPTHS_MAP_ID,
            "map_name": "Test Depths",
            "parent_id": None,
            "level_min": None,
            "level_max": None,
        },
    ]
    (build_dir / "zones.json").write_text(json.dumps(zones), encoding="utf-8")

    items = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    items += [
        {
            "id": 5010,
            "name": "Test Depths Relic",
            "quality": 3,
            "item_level": 60,
            "required_level": 55,
            "class_id": 4,
            "subclass_id": 4,
            "inventory_type": 5,
        },
        {
            "id": 5011,
            "name": "Test Depths Boss Drop",
            "quality": 3,
            "item_level": 60,
            "required_level": 55,
            "class_id": 4,
            "subclass_id": 4,
            "inventory_type": 5,
        },
    ]
    (build_dir / "items.json").write_text(json.dumps(items), encoding="utf-8")

    map_csv = (build_dir / "raw" / "Map.csv").read_text(encoding="utf-8")
    map_csv += f"{_TEST_DEPTHS_MAP_ID},TestDepths,Test Depths,1,5001\n"
    (build_dir / "raw" / "Map.csv").write_text(map_csv, encoding="utf-8")

    # The classic-db-only trash mob: the fork's own database (`_augmented_
    # engine`) never names npc 911 at all, so only `instance_zone_by_map`
    # -- keyed by this record's own `map_id` -- can place item 5010
    # inside the dungeon instead of a flat `world:test-depths-trash`.
    #
    # loot-parity-2 lane, 2026-09-30: item 5011 ALSO gets a classic-db
    # `creature_drop` record for the SAME npc (910) the fork's own boss
    # already names it under -- the Onyxia's Lair shape
    # (`pipeline.loot.sources._classic_db_corroborated_raid_dungeon_
    # origin`'s own doc) in miniature: classic-db independently
    # corroborates the fork's ENTIRE contribution to this dungeon, so
    # `dungeon:test-depths` should be promoted to `source_origin:
    # "classic-db"` in both `loot` and `loot-merge`'s output --
    # `test_loot_merge_agrees_on_source_shape_and_origin_labels` below.
    write_classic_sources(
        build_dir,
        {
            5010: [
                ClassicDbSourceRecord(
                    kind="creature_drop",
                    npc_id=911,
                    name="Test Depths Trash",
                    map_id=_TEST_DEPTHS_MAP_ID,
                    chance=10.0,
                )
            ],
            5011: [
                ClassicDbSourceRecord(
                    kind="creature_drop",
                    npc_id=910,
                    name="Test Depths Boss",
                    map_id=_TEST_DEPTHS_MAP_ID,
                    chance=50.0,
                )
            ],
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


def _source_shape(source: dict) -> tuple[frozenset, str | None, tuple]:
    """loot-parity-2 lane, 2026-09-30: the per-source shape `test_loot_
    merge_agrees_on_source_shape_and_origin_labels` diffs between the two
    generators, over and above `_dungeon_items`'s own item-id-only
    comparison -- the KEY SET a source carries (whether `source_origin`
    is present at all, same gap the lane's own report measured on
    `raid:onyxias-lair`) and each boss's own `source_origin`/`item_
    source_origin` attribution."""
    bosses = tuple(
        sorted(
            (boss["npc_id"], tuple(sorted((boss.get("item_source_origin") or {}).items())))
            for boss in source.get("bosses") or []
        )
    )
    return frozenset(source.keys()), source.get("source_origin"), bosses


def test_loot_places_the_classicdb_only_trash_mob_inside_the_dungeon(tmp_path: Path):
    """Baseline: `write_loot_files`, with `raw/Map.csv` available, already
    gets this right -- both the fork's own boss (5011) and classic-db's
    rescued trash mob (5010) land in `dungeon:test-depths`."""
    root = _prepare_build(tmp_path)
    engine_dir = _augmented_engine(tmp_path)
    curated = prepare_curated(tmp_path)
    write_loot_files(
        BUILD,
        engine_dir=engine_dir,
        root=root,
        overlay_dir=tmp_path / "no-overlay",
        curated_dir=curated,
        ids_md=FIXTURES / "IDS.md",
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
    instead (its recovery from the committed loot.json never finding zone id 5002's
    real `map_id`)."""
    root = _prepare_build(tmp_path)
    engine_dir = _augmented_engine(tmp_path)
    curated = prepare_curated(tmp_path)
    write_loot_files(
        BUILD,
        engine_dir=engine_dir,
        root=root,
        overlay_dir=tmp_path / "no-overlay",
        curated_dir=curated,
        ids_md=FIXTURES / "IDS.md",
    )
    full_document = json.loads((root / BUILD / LOOT).read_text(encoding="utf-8"))

    merge_loot_files(
        BUILD,
        engine_dir=engine_dir,
        root=root,
        overlay_dir=tmp_path / "no-overlay",
        curated_dir=curated,
    )
    merged_document = json.loads((root / BUILD / LOOT).read_text(encoding="utf-8"))

    assert _dungeon_items(merged_document, "dungeon:test-depths") == {5010, 5011}

    def by_id(document: dict) -> dict[str, set[int]]:
        return {
            source["id"]: _dungeon_items(document, source["id"]) for source in document["sources"]
        }

    assert by_id(merged_document) == by_id(full_document)


def test_loot_merge_agrees_on_source_shape_and_origin_labels(tmp_path: Path):
    """loot-parity-2 lane, 2026-09-30: the parity contract above only ever
    diffed item ID sets -- it would have missed `raid:onyxias-lair`'s own
    real-build gap (this lane's report: the `loot` path's raid sources
    carried no `source_origin` key at all where `loot-merge` did) because
    both sides named the exact same items either way. This diffs the
    FULL per-source shape instead -- every key a source carries, plus
    each boss's own `source_origin`/`item_source_origin` attribution
    (`_source_shape`) -- between the same two runs
    `test_loot_merge_agrees_with_loot_on_the_same_committed_caches` makes,
    and additionally asserts `dungeon:test-depths` (the fixture's own
    Onyxia's-Lair-shaped case -- classic-db's dump independently
    corroborates the fork's entire boss item 5011, `_prepare_build`'s own
    doc) is promoted to `source_origin: "classic-db"` in BOTH."""
    root = _prepare_build(tmp_path)
    engine_dir = _augmented_engine(tmp_path)
    curated = prepare_curated(tmp_path)
    write_loot_files(
        BUILD,
        engine_dir=engine_dir,
        root=root,
        overlay_dir=tmp_path / "no-overlay",
        curated_dir=curated,
        ids_md=FIXTURES / "IDS.md",
    )
    full_document = json.loads((root / BUILD / LOOT).read_text(encoding="utf-8"))

    merge_loot_files(
        BUILD,
        engine_dir=engine_dir,
        root=root,
        overlay_dir=tmp_path / "no-overlay",
        curated_dir=curated,
    )
    merged_document = json.loads((root / BUILD / LOOT).read_text(encoding="utf-8"))

    def shapes(document: dict) -> dict[str, tuple]:
        return {source["id"]: _source_shape(source) for source in document["sources"]}

    full_shapes = shapes(full_document)
    merged_shapes = shapes(merged_document)
    assert set(full_shapes) == set(merged_shapes)
    assert full_shapes == merged_shapes

    depths_keys, depths_origin, _ = full_shapes["dungeon:test-depths"]
    assert "source_origin" in depths_keys
    assert depths_origin == "classic-db"


# ---------------------------------------------------------------------------
# Round 8, owner ruling 3: one source of truth for the battleground and PvP
# rank classification. `loot` read Map.csv (Alterac Valley is a type-3
# battleground, so its spawns got no source) while `loot-merge` recovered
# instance types from the committed loot.json (no battlegrounds, so the same
# spawns became world sources), and loot.json flipped on every pair of runs.
# ---------------------------------------------------------------------------

#: Alterac Valley's client map id and AreaTable zone id.
_AV_MAP_ID = 30
_AV_ZONE_ID = 2597
_AV_CREATURE_ITEM = 5020
_AV_OBJECT_ITEM = 5021
#: The tabard shape: a classic id the client gives a RequiredPVPRank, and
#: the Forever re-itemised id with the same name and slot that supersedes it.
_CLASSIC_TABARD = 15196
_FOREVER_TABARD = 272694


def _prepare_battleground_build(tmp_path: Path) -> Path:
    root = _prepare_build(tmp_path)
    build_dir = root / BUILD
    zones = json.loads((build_dir / "zones.json").read_text(encoding="utf-8"))
    zones.append(
        {
            "id": _AV_ZONE_ID,
            "name": "Alterac Valley",
            "map_id": _AV_MAP_ID,
            "map_name": "Alterac Valley",
            "parent_id": None,
            "level_min": None,
            "level_max": None,
        }
    )
    (build_dir / "zones.json").write_text(json.dumps(zones), encoding="utf-8")

    tabard = {
        "quality": 1,
        "item_level": 20,
        "required_level": 0,
        "class_id": 4,
        "subclass_id": 0,
        "inventory_type": 19,
    }
    items = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    items += [
        {
            "id": _AV_CREATURE_ITEM,
            "name": "Alterac Ram Hide",
            "quality": 1,
            "item_level": 55,
            "required_level": 0,
            "class_id": 12,
            "subclass_id": 0,
            "inventory_type": 0,
        },
        {
            "id": _AV_OBJECT_ITEM,
            "name": "Coldtooth Supplies",
            "quality": 1,
            "item_level": 55,
            "required_level": 0,
            "class_id": 12,
            "subclass_id": 0,
            "inventory_type": 0,
        },
        {"id": _CLASSIC_TABARD, "name": "Private's Tabard", **tabard},
        {"id": _FOREVER_TABARD, "name": "Private's Tabard", **tabard},
    ]
    (build_dir / "items.json").write_text(json.dumps(items), encoding="utf-8")

    raw = build_dir / "raw"
    with (raw / "Map.csv").open("a", encoding="utf-8") as handle:
        handle.write(f"{_AV_MAP_ID},PVPZone01,Alterac Valley,3,{_AV_ZONE_ID}\n")
    with (raw / "ItemSparse.csv").open("a", encoding="utf-8") as handle:
        handle.write(f"{_CLASSIC_TABARD},Private's Tabard,5,0,0,0\n")

    from pipeline.classic_sources import load_classic_sources

    records = load_classic_sources(build_dir)
    records[_AV_CREATURE_ITEM] = [
        ClassicDbSourceRecord(
            kind="creature_drop",
            npc_id=10990,
            name="Alterac Ram",
            map_id=_AV_MAP_ID,
            chance=50.0,
        )
    ]
    records[_AV_OBJECT_ITEM] = [
        ClassicDbSourceRecord(
            kind="object_drop",
            object_id=178785,
            name="Coldtooth Supplies",
            map_id=_AV_MAP_ID,
            chance=100.0,
        )
    ]
    write_classic_sources(build_dir, records)
    return root


def _items_anywhere(document: dict) -> set[int]:
    return {
        item for source in document["sources"] for item in _dungeon_items(document, source["id"])
    }


@pytest.mark.parametrize("spawns_are_sources", [True, False])
def test_loot_and_loot_merge_write_byte_identical_loot_json(
    tmp_path: Path, monkeypatch: pytest.MonkeyPatch, spawns_are_sources: bool
):
    """Full run, merge, full run, merge over one fixture with Alterac
    Valley spawns and a client-ranked PvP tabard: all four loot.json files
    are byte-identical, whichever way the battleground switch is set, and
    the switch alone decides whether the spawns are sources."""
    monkeypatch.setattr(
        "pipeline.loot.constants.BATTLEGROUND_SPAWNS_ARE_SOURCES",
        spawns_are_sources,
        raising=False,
    )
    root = _prepare_battleground_build(tmp_path)
    engine_dir = _augmented_engine(tmp_path)
    curated = prepare_curated(tmp_path)
    loot_path = root / BUILD / LOOT
    outputs = []
    for _ in range(2):
        write_loot_files(
            BUILD,
            engine_dir=engine_dir,
            root=root,
            overlay_dir=tmp_path / "no-overlay",
            curated_dir=curated,
            ids_md=FIXTURES / "IDS.md",
        )
        outputs.append(loot_path.read_bytes())
        merge_loot_files(
            BUILD,
            engine_dir=engine_dir,
            root=root,
            overlay_dir=tmp_path / "no-overlay",
            curated_dir=curated,
        )
        outputs.append(loot_path.read_bytes())

    assert outputs[1] == outputs[0], "loot-merge disagrees with loot"
    assert len(set(outputs)) == 1, "loot.json flips between runs"

    document = json.loads(outputs[0])
    sourced = _items_anywhere(document)
    by_id = {source["id"]: source for source in document["sources"]}
    if spawns_are_sources:
        assert _AV_CREATURE_ITEM in by_id["world:alterac-ram"]["items"]
        assert _AV_OBJECT_ITEM in by_id["world:coldtooth-supplies"]["items"]
    else:
        assert {_AV_CREATURE_ITEM, _AV_OBJECT_ITEM}.isdisjoint(sourced)
    assert _FOREVER_TABARD in sourced


def test_loot_merge_fails_fast_without_the_full_runs_classification(tmp_path: Path):
    """No silent fallback to guessing instance types from an old loot.json:
    a build `loot` never ran on is refused with the command to run."""
    root = _prepare_battleground_build(tmp_path)
    engine_dir = _augmented_engine(tmp_path)
    with pytest.raises(SystemExit, match="python -m pipeline loot"):
        merge_loot_files(
            BUILD,
            engine_dir=engine_dir,
            root=root,
            overlay_dir=tmp_path / "no-overlay",
            curated_dir=prepare_curated(tmp_path),
        )


# ---------------------------------------------------------------------------
# The committed build: the merge is a fixed point on its own inputs, and the
# full run (where the raw client tables are on disk) writes the same bytes.
# Needs FOREVER_ENGINE_PATH, like test_loot_classicdb.py's committed-build
# checks.
# ---------------------------------------------------------------------------

ACTIVE_BUILD = json.loads(
    (HERE.parent.parent / "web/src/data/active-build.json").read_text(encoding="utf-8")
)["build"]


def _engine_or_skip() -> Path:
    engine = os.environ.get("FOREVER_ENGINE_PATH")
    if not engine or not (Path(engine) / "assets" / "database" / "db.json").exists():
        pytest.skip("FOREVER_ENGINE_PATH is unset or has no assets/database/db.json")
    return Path(engine)


def _copy_committed_build(tmp_path: Path, with_raw_tables: bool) -> Path:
    """The active build's committed files (raw/ client CSVs only when asked)."""
    source = Path("builds") / ACTIVE_BUILD
    root = tmp_path / "builds"
    shutil.copytree(
        source,
        root / ACTIVE_BUILD,
        ignore=None if with_raw_tables else shutil.ignore_patterns("*.csv", "*.txt"),
    )
    if not with_raw_tables:
        hotfixes = source / "raw" / "hotfixes"
        shutil.copytree(hotfixes, root / ACTIVE_BUILD / "raw" / "hotfixes", dirs_exist_ok=True)
    return root


def test_the_committed_loot_json_is_a_fixed_point_of_the_merge(tmp_path: Path):
    engine = _engine_or_skip()
    root = _copy_committed_build(tmp_path, with_raw_tables=False)
    loot_path = root / ACTIVE_BUILD / LOOT
    merge_loot_files(ACTIVE_BUILD, engine_dir=engine, root=root)
    first = loot_path.read_bytes()
    merge_loot_files(ACTIVE_BUILD, engine_dir=engine, root=root)
    assert loot_path.read_bytes() == first


def test_the_full_run_writes_the_merges_bytes_on_the_committed_build(tmp_path: Path):
    engine = _engine_or_skip()
    raw = Path("builds") / ACTIVE_BUILD / "raw"
    if not all((raw / name).exists() for name in ("Map.csv", "ItemSparse.csv")):
        pytest.skip(f"no raw client tables under {raw}; run `python -m pipeline fetch`")
    merged_root = _copy_committed_build(tmp_path / "merge", with_raw_tables=False)
    merge_loot_files(ACTIVE_BUILD, engine_dir=engine, root=merged_root)
    full_root = _copy_committed_build(tmp_path / "full", with_raw_tables=True)
    write_loot_files(ACTIVE_BUILD, engine_dir=engine, root=full_root)
    assert (full_root / ACTIVE_BUILD / LOOT).read_bytes() == (
        merged_root / ACTIVE_BUILD / LOOT
    ).read_bytes()


def test_battleground_spawns_are_loot_sources():
    """Owner ruling 3 (round 8): battlegrounds are in Forever at launch, so
    Alterac Valley's and Arathi Basin's creatures and objects stay sources."""
    from pipeline.loot import constants

    assert constants.BATTLEGROUND_SPAWNS_ARE_SOURCES is True


def test_buffs_only_names_a_new_id_and_leaves_loot_json_alone(tmp_path: Path):
    """`loot-merge --buffs-only`: an id IDS.md gained since the last full
    `loot` run is named again, and the nightly-owned loot.json is not
    rewritten; the manifest re-hashes simbuffs.json."""
    root = _prepare_build(tmp_path)
    engine_dir = _augmented_engine(tmp_path)
    curated = prepare_curated(tmp_path)
    ids_md = FIXTURES / "IDS.md"
    write_loot_files(
        BUILD,
        engine_dir=engine_dir,
        root=root,
        overlay_dir=tmp_path / "no-overlay",
        curated_dir=curated,
        ids_md=ids_md,
    )
    build_dir = root / BUILD
    full = json.loads((build_dir / SIMBUFFS).read_text(encoding="utf-8"))
    dropped = sorted(full["entries"])[0]
    trimmed = {**full, "entries": {k: v for k, v in full["entries"].items() if k != dropped}}
    (build_dir / SIMBUFFS).write_text(json.dumps(trimmed), encoding="utf-8")
    sentinel = b'{"sentinel": true}'
    (build_dir / LOOT).write_bytes(sentinel)

    written = merge_simbuffs_file(
        BUILD, engine_dir=engine_dir, root=root, curated_dir=curated, ids_md=ids_md
    )

    assert written == [build_dir / SIMBUFFS]
    merged = json.loads((build_dir / SIMBUFFS).read_text(encoding="utf-8"))
    assert merged["entries"][dropped] == full["entries"][dropped]
    assert (build_dir / LOOT).read_bytes() == sentinel
    manifest_files = json.dumps(read_manifest(build_dir))
    assert SIMBUFFS in manifest_files


def test_buffs_only_refuses_a_build_with_no_simbuffs(tmp_path: Path):
    root = _prepare_build(tmp_path)
    with pytest.raises(SystemExit, match="run the full `loot` command first"):
        merge_simbuffs_file(BUILD, engine_dir=_augmented_engine(tmp_path), root=root)
