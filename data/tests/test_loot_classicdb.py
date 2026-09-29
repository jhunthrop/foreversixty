# data/tests/test_loot_classicdb.py
"""`pipeline.loot.classicdb.instance_zone_by_map` -- the map-id -> zone-id
resolver `classicdb_additions` uses to place a classic-db creature/object
drop inside a dungeon/raid. Regression coverage for the bug this lane's
own report names: a zone resolved through `instance_types()`'s
ContinentID fallback (Onyxia's Lair, whose only `zones.json` row states
`map_id: 1`, the Kalimdor continent) must never be treated as a
resolvable instance map -- every open-world creature on that continent
would otherwise land on it.
"""

from pipeline.loot.classicdb import instance_zone_by_map

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
