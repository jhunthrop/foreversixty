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
    others = [
        s for s in document.sources if s.id.startswith("world:") and s.id != "world:azuregos"
    ]
    assert not others


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
