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

from pipeline.classic_sources import ClassicDbSourceRecord, load_classic_sources
from pipeline.forkdb import ForkDatabase, load_fork_database
from pipeline.loot import types_from_committed_loot
from pipeline.loot.classicdb import (
    _CREATURE_KINDS,
    _item_instance_zone_consensus,
    classic_db_npc_names,
    classicdb_additions,
    fork_instance_npc_zones,
    instance_zone_by_map,
)
from pipeline.loot.sources import item_factions
from pipeline.normalize.gear import SLOT_BY_INVENTORY_TYPE
from pipeline.quest_levels import load_quest_levels

BUILD = "1.60.1.70291"
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


# --- `preferred` -- the Westfall/Deadmines duplicate-AreaTable-row finding
# (day3 data-followups-11 lane, 2026-09-30): The Deadmines' own map id 36
# carries a second `types`-marked zone row, 206, whose `Name_lang` is
# verbatim "Westfall" (the OPEN-WORLD zone's own name, a different map
# entirely) -- see `instance_zone_by_map`'s own doc for the measured
# regression this closes (a second `dungeon:westfall` LootSource with
# Captain Greenskin's own boss list duplicated under a zone-shaped name).

_DEADMINES_DUPLICATE_ROWS = [
    {"id": 206, "name": "Westfall", "map_id": 36},
    {"id": 1581, "name": "The Deadmines", "map_id": 36},
]
_DEADMINES_DUPLICATE_TYPES = {206: 1, 1581: 1}


def test_a_duplicate_zone_named_row_wins_without_a_preferred_id():
    """Pins the OLD, wrong default: with no `preferred` id, whichever row
    sorts first in `zone_rows` wins -- 206 ("Westfall") before 1581 ("The
    Deadmines") here, the same order this build's own `zones.json` lists
    them in."""
    by_map = instance_zone_by_map(_DEADMINES_DUPLICATE_ROWS, _DEADMINES_DUPLICATE_TYPES)
    assert by_map[36] == 206


def test_a_preferred_id_wins_over_a_zone_named_duplicate():
    by_map = instance_zone_by_map(
        _DEADMINES_DUPLICATE_ROWS, _DEADMINES_DUPLICATE_TYPES, frozenset({1581})
    )
    assert by_map[36] == 1581


def test_a_preferred_id_wins_regardless_of_which_row_comes_first():
    reversed_rows = list(reversed(_DEADMINES_DUPLICATE_ROWS))
    by_map = instance_zone_by_map(
        reversed_rows, _DEADMINES_DUPLICATE_TYPES, frozenset({1581})
    )
    assert by_map[36] == 1581


def test_a_preferred_id_that_names_no_row_for_this_map_changes_nothing():
    by_map = instance_zone_by_map(
        _DEADMINES_DUPLICATE_ROWS, _DEADMINES_DUPLICATE_TYPES, frozenset({99999})
    )
    assert by_map[36] == 206


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


# --- `_item_instance_zone_consensus` and the Molten Core Firelord/Lava
# Annihilator finding it fixes (day3 data-followups-5 lane, 2026-09-30) ---

_MC_ZONE_BY_MAP = {409: 2717}


def _creature(npc_id, name, map_id, chance):
    return ClassicDbSourceRecord(
        kind="creature_drop", npc_id=npc_id, name=name, map_id=map_id, chance=chance
    )


def _object(object_id, name, map_id, chance=None, is_reward_chest=True):
    return ClassicDbSourceRecord(
        kind="object_drop", object_id=object_id, name=name, map_id=map_id, chance=chance,
        is_reward_chest=is_reward_chest,
    )  # fmt: skip


def test_item_instance_zone_consensus_resolves_an_unmapped_npc_from_its_siblings():
    """Firelord's own shape for Fiery Core (17010): its own record states
    no map, but the item's other four creature records all resolve to
    Molten Core (2717) alone."""
    records = [
        _creature(11666, "Firewalker", 409, 42.2),
        _creature(11667, "Flameguard", 409, 34.7),
        _creature(11668, "Firelord", None, 23.4),
        _creature(12056, "Baron Geddon", 409, 14.0),
    ]
    assert _item_instance_zone_consensus(records, _MC_ZONE_BY_MAP, {}) == 2717


def test_item_instance_zone_consensus_is_none_without_a_single_agreeing_zone():
    """No fallback when the item's OTHER creature records don't themselves
    agree on one zone (a real cross-instance/world item) or resolve none
    at all -- Firelord's own OTHER 68 items are exactly this shape."""
    unresolved_only = [_creature(11668, "Firelord", None, 1.0), _creature(1, "Rat", 0, 1.0)]
    assert _item_instance_zone_consensus(unresolved_only, _MC_ZONE_BY_MAP, {}) is None

    two_zones = [
        _creature(11668, "Firelord", None, 1.0),
        _creature(2, "Deadmines Rat", 36, 1.0),
        _creature(3, "Blackrock Ogre", 230, 1.0),
    ]
    zone_by_map = {**_MC_ZONE_BY_MAP, 36: 1581, 230: 1584}
    assert _item_instance_zone_consensus(two_zones, zone_by_map, {}) is None


def test_item_instance_zone_consensus_falls_back_to_fork_instance_npcs_too():
    records = [_creature(11668, "Firelord", None, 1.0), _creature(1853, "Gandling", None, 1.0)]
    assert _item_instance_zone_consensus(records, {}, {1853: 2717}) == 2717


def test_item_instance_zone_consensus_counts_a_resolvable_object_drop_sibling():
    """day3 data-followups-6 lane, 2026-09-30, this lane's brief item (c):
    Cache of the Firelord's own shape -- two Flamewaker creature records
    with no static spawn map at all, and no OTHER creature-kind sibling
    to agree with -- but the item's own chest `object_drop` record DOES
    resolve to Molten Core once `map_id` is correct (this lane's brief
    item a), and that alone must be enough for the two creature rows to
    reach consensus."""
    records = [
        _creature(11663, "Flamewaker Healer", None, 0.44),
        _creature(11664, "Flamewaker Elite", None, 0.64),
        _object(179703, "Cache of the Firelord", 409, 50.0),
    ]
    assert _item_instance_zone_consensus(records, _MC_ZONE_BY_MAP, {}) == 2717


def test_item_instance_zone_consensus_ignores_an_object_drop_with_no_map():
    """An object_drop record contributes nothing to consensus when ITS
    OWN map is unresolved either -- the fallback only ever reaches a real
    fact, never invents one from an equally-unmapped sibling."""
    records = [
        _creature(11663, "Flamewaker Healer", None, 0.44),
        _object(179703, "Cache of the Firelord", None, 50.0),
    ]
    assert _item_instance_zone_consensus(records, _MC_ZONE_BY_MAP, {}) is None


_MC_TYPES = {2717: 2}  # raid
_MC_ZONE_ROWS = [{"id": 2717, "name": "Molten Core", "map_id": 409}]


def _classicdb_additions(classic_sources):
    return classicdb_additions(
        classic_sources,
        build_items=set(classic_sources),
        equippable=set(),
        zone_names={2717: "Molten Core"},
        types=_MC_TYPES,
        zone_rows=_MC_ZONE_ROWS,
        quest_levels={},
        item_factions={},
    )


def _raid_boss_items(sources, npc_id):
    for source in sources:
        if source.kind == "raid":
            for boss in source.bosses or []:
                if boss.npc_id == npc_id:
                    return set(boss.items)
    return set()


def _world_items(sources):
    return {item_id for source in sources if source.kind == "world" for item_id in source.items}


def _world_drop_items(sources):
    return {
        item_id for source in sources if source.kind == "world_drop" for item_id in source.items
    }


def test_fiery_core_s_firelord_lands_in_raid_trash_not_a_world_bucket():
    """Regression for the exact bug the 9th wow-player sweep found
    (day3/player-review-24): Fiery Core (item id below) named `world:
    firelord` beside `raid:molten-core`, which let a reagent-gate check
    see an open source and leave Nightfall/Ebon Hand/Blackfury ungated."""
    fiery_core = 90010
    classic_sources = {
        fiery_core: [
            _creature(11666, "Firewalker", 409, 42.2),
            _creature(11667, "Flameguard", 409, 34.7),
            _creature(11668, "Firelord", None, 23.4),
            _creature(11659, "Molten Destroyer", 409, 14.9),
            _creature(12056, "Baron Geddon", 409, 14.0),
        ]
    }
    sources, _, _ = _classicdb_additions(classic_sources)
    assert fiery_core in _raid_boss_items(sources, 11668)
    assert fiery_core not in _world_items(sources)


def test_lava_core_s_lava_annihilator_lands_in_raid_trash_not_a_world_drop_pool():
    """Regression for Lava Core: 7 creature-kind records for one item
    clears `is_world_drop_pattern`'s creature-count floor on its own, so
    Lava Annihilator's unmapped record used to fold into a synthetic
    `world_drop:<range>` pool instead of Molten Core's own trash."""
    lava_core = 90011
    classic_sources = {
        lava_core: [
            _creature(11665, "Lava Annihilator", None, 12.5),
            _creature(12076, "Lava Elemental", 409, 23.8),
            _creature(12101, "Lava Surger", 409, 3.5),
            _creature(12100, "Lava Reaver", 409, 22.1),
            _creature(11659, "Molten Destroyer", 409, 15.4),
            _creature(11988, "Golemagg the Incinerator", 409, 14.0),
            _creature(12057, "Garr", 409, 14.0),
        ]
    }
    sources, _, _ = _classicdb_additions(classic_sources)
    assert lava_core in _raid_boss_items(sources, 11665)
    assert lava_core not in _world_drop_items(sources)


def test_an_unmapped_npc_with_no_resolving_sibling_still_reads_as_open_world():
    """The consensus fallback must not fire when there is nothing to reach
    consensus WITH -- Firelord's own OTHER 68 items on the real dump are
    typically shared with hundreds of other creatures that themselves
    never resolve to one zone either; the simplest such shape is a lone
    unmapped npc with no creature-kind sibling at all, which must still
    read as `world:<name>`, exactly as before this fallback existed."""
    unrelated_item = 90012
    classic_sources = {unrelated_item: [_creature(11668, "Firelord", None, 1.0)]}
    sources, _, _ = _classicdb_additions(classic_sources)
    assert unrelated_item in _world_items(sources)
    assert unrelated_item not in _raid_boss_items(sources, 11668)


def _boss(sources, npc_or_object_id):
    for source in sources:
        for boss in source.bosses or []:
            if boss.npc_id == npc_or_object_id:
                return source, boss
    return None, None


def _trash_items(sources, zone_id):
    for source in sources:
        if source.zone_id == zone_id:
            return set(source.trash or [])
    return set()


def test_a_named_reward_chest_in_an_instance_becomes_a_boss_not_anonymous_trash():
    """day3 data-followups-6 lane, 2026-09-30, this lane's brief item (b):
    Cache of the Firelord's own shape once `pipeline.classic_sources.
    _parse_object_drops` resolves `data1` to the owning gameobject
    (this lane's brief item a) -- a named, `is_reward_chest` object_drop
    record in Molten Core becomes a boss entry keyed by the object's own
    entry, carrying the chest's real name, not `trash[zone_id]`."""
    cache_item = 18803
    classic_sources = {
        cache_item: [
            _creature(11663, "Flamewaker Healer", None, 0.14),
            _creature(11664, "Flamewaker Elite", None, 0.54),
            _object(179703, "Cache of the Firelord", 409, 0.0),
        ]
    }
    sources, _, _ = _classicdb_additions(classic_sources)
    source, boss = _boss(sources, 179703)
    assert source is not None and source.kind == "raid"
    assert boss.name == "Cache of the Firelord"
    assert cache_item in boss.items
    assert cache_item not in _trash_items(sources, 2717)
    assert cache_item not in _world_drop_items(sources)
    assert cache_item not in _world_items(sources)


def test_an_ordinary_instance_object_with_no_data1_redirect_stays_trash():
    """The negative case `is_reward_chest`'s own doc names: an ore vein or
    a generic chest whose own entry IS its own loot key (`data1` unset)
    stays anonymous trash even though it is named -- `record.name`'s own
    truthiness is never the signal, `is_reward_chest` is."""
    ore = 90013
    classic_sources = {
        ore: [_object(50001, "Rich Thorium Vein", 409, 100.0, is_reward_chest=False)]
    }
    sources, _, _ = _classicdb_additions(classic_sources)
    assert ore in _trash_items(sources, 2717)
    source, boss = _boss(sources, 50001)
    assert boss is None


def test_classic_db_npc_names_resolves_a_gameobject_id_in_the_npc_slot_too():
    """day3 data-followups-6 lane, 2026-09-30, this lane's brief item (b):
    the fork puts a reward chest's own GAMEOBJECT entry in `drop.npcId`
    (Cache of the Firelord, Four Horsemen Chest, Father Flame) --
    `_resolve_or_drop_unnamed_bosses` looks that same numeric id up in
    this dict, so an `object_drop` record's own `object_id`/`name` must
    resolve it too, not just a real `npc_id`."""
    classic_sources = {
        18803: [
            _creature(11663, "Flamewaker Healer", None, 0.14),
            _object(179703, "Cache of the Firelord", 409, 0.0),
        ]
    }
    names = classic_db_npc_names(classic_sources)
    assert names[11663] == "Flamewaker Healer"
    assert names[179703] == "Cache of the Firelord"


def test_classic_db_npc_names_prefers_a_real_npc_name_over_an_object_one():
    """`setdefault`, npc pass first: an id that happens to be both a real
    creature_template entry AND an unrelated gameobject entry keeps the
    npc's own name -- the more direct fact of the two."""
    classic_sources = {
        1: [_creature(500, "Real Creature", None, 1.0)],
        2: [_object(500, "Unrelated Object", None, 1.0)],
    }
    assert classic_db_npc_names(classic_sources)[500] == "Real Creature"
