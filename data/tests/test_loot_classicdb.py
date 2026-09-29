# data/tests/test_loot_classicdb.py
"""`pipeline.loot.classicdb.instance_zone_by_map` -- the map-id -> zone-id
resolver `classicdb_additions` uses to place a classic-db creature/object
drop inside a dungeon/raid. Regression coverage for the bug this lane's
own report names: a zone resolved through `instance_types()`'s
ContinentID fallback (Onyxia's Lair, whose only `zones.json` row states
`map_id: 1`, the Kalimdor continent) must never be treated as a
resolvable instance map -- every open-world creature on that continent
would otherwise land on it.

Also `fork_instance_npc_zones` -- the npc_id -> zone_id fallback for a
dungeon/raid boss classic-db's own spawn table has no usable map for at
all (src-classicdb-fixes lane, 2026-09-29; Darkmaster Gandling, npc 1853,
is the real case: a scripted Scholomance encounter with no static
`creature` spawn row)."""

import json
import os
from pathlib import Path

import pytest

from pipeline.classic_sources import load_classic_sources
from pipeline.forkdb import ForkDatabase, load_fork_database
from pipeline.loot import types_from_committed_loot
from pipeline.loot.classicdb import (
    _CREATURE_KINDS,
    classicdb_additions,
    fork_instance_npc_zones,
    instance_zone_by_map,
)
from pipeline.loot.sources import item_factions
from pipeline.normalize.gear import SLOT_BY_INVENTORY_TYPE
from pipeline.quest_levels import load_quest_levels

BUILD = "1.60.1.70009"
BUILD_DIR = Path("builds") / BUILD


def _fork(items: list[dict]) -> ForkDatabase:
    return ForkDatabase(
        items=tuple(items), enchants=(), random_suffixes=(), zones={}, npcs={}, factions={},
        item_icons={}, spell_icons={}, spell_icon_rows=(), item_icon_rows=(),
    )

TYPES = {
    1581: 1,  # The Deadmines -- dungeon
    2159: 2,  # Onyxia's Lair -- raid, but map_id below is a bare continent id
}

ZONE_ROWS = [
    {"id": 1581, "name": "The Deadmines", "map_id": 36},
    {"id": 2159, "name": "Onyxia's Lair", "map_id": 1},  # AreaTable ContinentID fallback
]


def test_a_real_instance_map_resolves_to_its_zone_id():
    by_map = instance_zone_by_map(ZONE_ROWS, TYPES)
    assert by_map[36] == 1581


def test_a_bare_continent_map_id_is_never_resolved_even_when_types_marks_it_an_instance():
    by_map = instance_zone_by_map(ZONE_ROWS, TYPES)
    assert 1 not in by_map  # Kalimdor -- would otherwise catch every open-world creature
    assert 0 not in by_map  # Eastern Kingdoms, same reasoning


def test_a_zone_types_does_not_mark_as_an_instance_is_ignored():
    by_map = instance_zone_by_map(
        [*ZONE_ROWS, {"id": 999, "name": "Some Open Zone", "map_id": 99}], TYPES
    )
    assert 99 not in by_map


def test_fork_instance_npc_zones_names_every_npc_a_dungeon_or_raid_drop_places():
    fork = _fork(
        [
            {"id": 103, "sources": [{"drop": {"npcId": 902, "zoneId": 1581}}]},
            {"id": 200, "sources": [{"drop": {"npcId": 1853, "zoneId": 2159}}]},
        ]
    )
    zones = fork_instance_npc_zones(fork, TYPES)
    assert zones == {902: 1581, 1853: 2159}


def test_fork_instance_npc_zones_skips_an_open_world_drop():
    """A drop whose zone `types` does not mark a dungeon or raid (an
    open-world mob) contributes nothing -- only an instance boss is a
    fallback candidate."""
    fork = _fork([{"id": 300, "sources": [{"drop": {"npcId": 9000, "zoneId": 16}}]}])
    assert fork_instance_npc_zones(fork, TYPES) == {}


def test_fork_instance_npc_zones_skips_a_drop_with_no_npc_or_zone():
    fork = _fork(
        [
            {"id": 301, "sources": [{"drop": {"zoneId": 1581}}]},  # trash, no npc
            {"id": 302, "sources": [{}]},  # no drop info at all
        ]
    )
    assert fork_instance_npc_zones(fork, TYPES) == {}


def test_every_fork_known_instance_boss_lands_in_its_own_instance_from_classic_db():
    """Cross-check over the REAL committed build: for every npc_id the
    fork's own database already places in a dungeon or raid
    (`fork_instance_npc_zones`), every classic-db creature/skinning/
    pickpocketing record for that same npc must resolve to that SAME
    zone -- never a flat `world:<name>` bucket -- once
    `fork_instance_npc_zones`'s fallback is wired in (src-classicdb-fixes
    lane, 2026-09-29; Darkmaster Gandling, Scholomance's npc 1853, is the
    real case this closes: `map_id=None` in the dump, formerly stranded
    on `world:darkmaster-gandling`).

    Needs FOREVER_ENGINE_PATH (a wowsims-forever checkout) for the fork's
    own npc-placement facts; skips without it, same as
    test_simdb_fork_parity.py.
    """
    engine = os.environ.get("FOREVER_ENGINE_PATH")
    if not engine or not (Path(engine) / "assets" / "database" / "db.json").exists():
        pytest.skip("FOREVER_ENGINE_PATH is unset or has no assets/database/db.json")

    fork = load_fork_database(Path(engine))
    zone_rows = json.loads((BUILD_DIR / "zones.json").read_text(encoding="utf-8"))
    item_rows = json.loads((BUILD_DIR / "items.json").read_text(encoding="utf-8"))
    build_items = {int(row["id"]) for row in item_rows}
    item_inventory_types = {int(row["id"]): int(row["inventory_type"]) for row in item_rows}
    equippable = {
        item_id
        for item_id, inventory_type in item_inventory_types.items()
        if inventory_type in SLOT_BY_INVENTORY_TYPE
    }
    zone_names = {int(row["id"]): row["name"] for row in zone_rows}
    types = types_from_committed_loot(BUILD_DIR)
    classic_sources = load_classic_sources(BUILD_DIR)
    quest_levels = load_quest_levels(BUILD_DIR)
    factions = item_factions(fork, build_items)
    fork_instance_npcs = fork_instance_npc_zones(fork, types)

    # Every npc classic-db itself names via a creature/skinning/
    # pickpocketing record, that the fork already places in an instance --
    # the population this fallback must rescue from `world`. Checked at
    # the NPC level, not the item level: an item can drop from both an
    # instance boss and an unrelated open-world mob, so an item id alone
    # would still show up in a `world:` bucket for a completely different
    # reason and give a false positive.
    target_npc_ids = {
        record.npc_id
        for records in classic_sources.values()
        for record in records
        if record.kind in _CREATURE_KINDS and record.npc_id in fork_instance_npcs
    }
    assert target_npc_ids, "the pinned dump names no npc the fork places in an instance"

    def instance_boss_npc_ids(fork_npcs: dict[int, int]) -> set[int]:
        sources, _, _ = classicdb_additions(
            classic_sources, build_items, equippable, zone_names, types, zone_rows,
            quest_levels, factions, fork_npcs,
        )
        return {
            boss.npc_id
            for source in sources
            if source.kind in ("dungeon", "raid")
            for boss in source.bosses or []
        }

    before = instance_boss_npc_ids({})
    after = instance_boss_npc_ids(fork_instance_npcs)
    rescued = target_npc_ids - before
    assert rescued, "expected at least one fork-known instance npc to move off world: sources"
    still_world = target_npc_ids - after
    assert not still_world, (
        f"npc ids the fork places in an instance still fall through to world:, formerly "
        f"resolved via classic-db's own spawn map alone -- {sorted(still_world)}"
    )
