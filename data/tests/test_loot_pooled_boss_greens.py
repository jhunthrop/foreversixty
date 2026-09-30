# data/tests/test_loot_pooled_boss_greens.py
"""pooled-boss-greens lane, 2026-09-29: a dungeon/raid `creature_drop`
row still keeps its own boss attribution UNLESS the item is ALSO a known
world-drop pool member by another route (`ClassicDbSourceRecord(kind=
"world_drop", ...)` for the same item id, produced by a differently
shaped -- marked/multi-map/fan-out -- `reference_loot_template` pool
elsewhere in the dump, `pipeline.classic_sources._world_drop_records`)
AND this row's own chance is unknown (cmangos' own `0`/`None` sentinel)
or below `WORLD_DROP_MAX_CHANCE_PERCENT`.

cmangos gives many dungeon/raid bosses a "(Boss Loot)" reference group
used by exactly ONE referencing creature (that boss alone) holding the
boss's real drops AND a large equal-weight random-BoE-green tail at
`ChanceOrQuestChance` 0 -- e.g. `reference_loot_template` 35009,
Maraudon's Princess Theradras, 265 rows, one user. `pipeline.
classic_sources._world_drop_pools`' three signals (a marked comment, more
than `_SHARED_REFERENCE_MAX_USERS` referencing entries, `_MULTI_MAP_MIN_
MAPS` or more distinct maps) all key off the REFERENCING creature count,
so a one-user reference like this is invisible to every one of them even
though each item inside it is, separately, a well-known world drop with
its own `world_drop` record from elsewhere in the dump. Before this
lane's fix, `is_confirmed_boss_drop` read "does this row resolve to a
dungeon/raid zone, whatever its chance" -- true for every one of these
rows -- so all ~260 pooled greens on Theradras's own `creature_loot_
template` row stayed attributed to her personally, on top of their own
honest `world_drop` source, measured on build 1.60.1.70009 as audit
category C minors 1,817 -> 27,927.

Reuses `tests/test_loot_sources_classicdb.py`'s own fixture (`tests/
fixtures/loot`, `UNSOURCED_ITEM` 110 "Suffixed Sword", npc 657 "Defias
Pirate" on map 36 -> `dungeon:the-deadmines`) rather than a new one --
same engine checkout, same dungeon, just a `world_drop` record added
alongside the creature row this lane's fix changes the handling of.
"""

import json
from pathlib import Path

from pipeline.classic_sources import ClassicDbSourceRecord
from pipeline.csvio import read_csv
from pipeline.forkdb import load_fork_database
from pipeline.loot.sources import build_loot, instance_types, pvp_ranks, source_item_ids

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"

UNSOURCED_ITEM = 110  # "Suffixed Sword" -- the fixture's one fork-unsourced item
DEFIAS_PIRATE = 657  # the-deadmines, map 36 -- tests/test_loot_sources_classicdb.py


def built(classic_sources=None):
    fork = load_fork_database(ENGINE)
    zone_rows = json.loads((ENGINE / "zones.json").read_text(encoding="utf-8"))
    item_rows = json.loads((ENGINE / "items.json").read_text(encoding="utf-8"))
    build_items = {row["id"] for row in item_rows}
    item_inventory_types = {row["id"]: row["inventory_type"] for row in item_rows}
    types = instance_types(read_csv(ENGINE / "Map.csv"), zone_rows)
    return build_loot(
        fork,
        {row["id"]: row["name"] for row in zone_rows},
        types,
        pvp_ranks(read_csv(ENGINE / "ItemSparse.csv")),
        build_items,
        item_inventory_types,
        None,
        None,
        classic_sources,
        zone_rows,
    )


def source(document, source_id: str):
    for candidate in document.sources:
        if candidate.id == source_id:
            return candidate
    raise AssertionError(f"no source {source_id}; have {[s.id for s in document.sources]}")


def boss_items(document, source_id: str, npc_id: int) -> list[int]:
    dungeon = source(document, source_id)
    boss = next((b for b in (dungeon.bosses or []) if b.npc_id == npc_id), None)
    return boss.items if boss is not None else []


def test_a_pooled_boss_green_with_a_zero_chance_folds_into_world_drop_not_the_boss():
    """The item carries its own `world_drop:20-25` record (a marked
    pool elsewhere names it) AND a `creature_drop` row on Defias Pirate
    (an instance boss) whose own chance is 0 -- cmangos' "unknown"
    sentinel, exactly the "(Boss Loot)" equal-share-group shape Princess
    Theradras's own reference 35009 has. The row folds: Defias Pirate
    gets no entry for it, and the item stays exactly once, under its own
    `world_drop` source."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(kind="world_drop", name="World drop", level_min=20, level_max=25),
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=DEFIAS_PIRATE, name="Defias Pirate", map_id=36,
                chance=0.0,
            ),
        ]
    }
    document, _ = built(classic_sources)
    assert UNSOURCED_ITEM not in boss_items(document, "dungeon:the-deadmines", DEFIAS_PIRATE)
    world_drop = source(document, "world_drop:20-25")
    assert world_drop.items == [UNSOURCED_ITEM]
    assert not any(
        candidate.id.startswith("world:") and UNSOURCED_ITEM in source_item_ids(candidate)
        for candidate in document.sources
    )


def test_a_tier_item_with_no_world_drop_record_keeps_its_boss_attribution_at_any_chance():
    """The raid-tier invariant `raid-loot-regression` established: an
    item with NO `world_drop` record anywhere keeps every instance
    attribution whatever its own row's chance -- Defias Pirate's own row
    here states 0 (unknown) and still keeps the item, since nothing else
    names this item a world-drop pool member."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=DEFIAS_PIRATE, name="Defias Pirate", map_id=36,
                chance=0.0,
            ),
        ]
    }
    document, _ = built(classic_sources)
    assert boss_items(document, "dungeon:the-deadmines", DEFIAS_PIRATE) == [UNSOURCED_ITEM]
    assert not any(candidate.id.startswith("world_drop") for candidate in document.sources)


def test_a_pooled_boss_green_with_a_real_chance_keeps_its_boss_attribution():
    """Same shape as the zero-chance test above, but Defias Pirate's own
    row states a real 20% chance -- clearing `WORLD_DROP_MAX_CHANCE_
    PERCENT` means this IS a specific, confirmed boss kill regardless of
    what else the item is also known for, so it keeps its attribution
    (and still legitimately also appears under its own `world_drop`
    pool, from the OTHER, unrelated pool that names it)."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(kind="world_drop", name="World drop", level_min=20, level_max=25),
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=DEFIAS_PIRATE, name="Defias Pirate", map_id=36,
                chance=20.0,
            ),
        ]
    }
    document, _ = built(classic_sources)
    assert boss_items(document, "dungeon:the-deadmines", DEFIAS_PIRATE) == [UNSOURCED_ITEM]
    world_drop = source(document, "world_drop:20-25")
    assert world_drop.items == [UNSOURCED_ITEM]


def test_an_open_world_creature_row_is_unaffected_by_the_world_drop_record_check():
    """An open-world creature (map 1, a bare continent id, never a
    specific instance) is already exempt from boss attribution -- the
    new world-drop-record check only ever applies to a row that resolves
    to a dungeon/raid zone in the first place, so a `world_drop` record
    on the same item changes nothing about the open-world row's own flat
    `world:<name>` bucket."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(kind="world_drop", name="World drop", level_min=20, level_max=25),
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=9999, name="Some Kalimdor Mob", map_id=1, chance=0.0,
            ),
        ]
    }
    document, _ = built(classic_sources)
    world = source(document, "world:some-kalimdor-mob")
    assert world.items == [UNSOURCED_ITEM]
    world_drop = source(document, "world_drop:20-25")
    assert world_drop.items == [UNSOURCED_ITEM]
