import json
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.forkdb import ForkDatabase, load_fork_database
from pipeline.loot.sources import (
    KIND_ORDER,
    SourceIdCollision,
    build_loot,
    instance_types,
    pvp_ranks,
    quest_ids_for_build,
)
from pipeline.quest_levels import QuestLevelEntry

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"


def zone_rows():
    return json.loads((ENGINE / "zones.json").read_text(encoding="utf-8"))


def built(quest_levels=None):
    fork = load_fork_database(ENGINE)
    rows = zone_rows()
    build_items = {
        row["id"] for row in json.loads((ENGINE / "items.json").read_text(encoding="utf-8"))
    }
    return build_loot(
        fork,
        {row["id"]: row["name"] for row in rows},
        instance_types(read_csv(ENGINE / "Map.csv"), rows),
        pvp_ranks(read_csv(ENGINE / "ItemSparse.csv")),
        build_items,
        {},
        quest_levels,
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
        "zone:16",
        "crafted:blacksmithing",
        "rep:argent-dawn:exalted",
        "pvp:rank-11",
        "quest",
    ]
    # The fixture has no vendor rows, so every kind but `vendor` is emitted
    # once, in the contract's order (Azuregos in Azshara is also zone 16's
    # one open-world drop).
    assert [s.kind for s in document.sources] == [k for k in KIND_ORDER if k != "vendor"]


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
    others = [s for s in document.sources if s.id.startswith("world:") and s.id != "world:azuregos"]
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


def test_quest_ids_for_build_is_every_quest_a_build_item_names():
    fork = load_fork_database(ENGINE)
    build_items = {
        row["id"] for row in json.loads((ENGINE / "items.json").read_text(encoding="utf-8"))
    }
    assert quest_ids_for_build(fork, build_items) == {42, 43}
    # An item the build does not have (113) names no quest here, but this
    # fixture's only quest items (108, 112) are both in the build, so the
    # id set does not change -- proven by narrowing build_items instead.
    assert quest_ids_for_build(fork, {108}) == {42}


def test_quests_map_falls_back_to_the_item_level_proxy_with_no_quest_levels_entry():
    """Item 108 (quest 42, item_level 40) and item 112 (quest 43, item_level
    70) both have no entry in this test's quest_levels (empty/None) --
    neither classic-db nor wowhead covers them here -- so both fall back
    to item_level_proxy: min(60, item_level - 5)."""
    document, _ = built()
    entry_108 = document.quests["108"][0]
    assert (entry_108.quest_id, entry_108.min_level, entry_108.level) == (42, 35, 35)
    assert entry_108.level_source == "item_level_proxy"
    entry_112 = document.quests["112"][0]
    # item_level 70 - 5 = 65, floored at MAX_QUEST_LEVEL (60).
    assert (entry_112.quest_id, entry_112.min_level, entry_112.level) == (43, 60, 60)
    assert entry_112.level_source == "item_level_proxy"


def test_quests_map_prefers_a_quest_levels_entry_over_the_item_level_proxy():
    """The entry's own `.source` (classic-db, or wowhead for an id
    classic-db lacks) passes straight through to level_source -- not a
    hardcoded label, since pipeline.quest_levels.load_quest_levels's
    result can carry either."""
    document, _ = built(
        quest_levels={
            42: QuestLevelEntry(min_level=12, level=14, source="classic-db", fetched_at="2026-09-28T00:00:00+00:00")  # noqa: E501
        }
    )
    entry_108 = document.quests["108"][0]
    assert (entry_108.min_level, entry_108.level, entry_108.level_source) == (12, 14, "classic-db")
    # Quest 43 still has no entry, so it is unaffected.
    entry_112 = document.quests["112"][0]
    assert entry_112.level_source == "item_level_proxy"


def test_quests_map_passes_through_a_wowhead_sourced_entry_unchanged():
    document, _ = built(
        quest_levels={
            42: QuestLevelEntry(min_level=9, level=11, source="wowhead", fetched_at="2026-09-28T00:00:00+00:00")  # noqa: E501
        }
    )
    entry_108 = document.quests["108"][0]
    assert (entry_108.min_level, entry_108.level, entry_108.level_source) == (9, 11, "wowhead")


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
    # Every fixture entry now has a kind: the unnamed open-world mob's drop
    # files under zone:16 and the vendor sale under its npc.
    assert stats.dropped_entries == 0
    # 105 (zone 16), 109 (the vendor) and 113 (the raid boss): the build's
    # item table has none of them, so they are counted before any kind
    # could take them.
    assert stats.absent_items == 3


def _wsg_fork(repFactionId: int, factionRestriction: int) -> ForkDatabase:
    return ForkDatabase(
        items=(
            {
                "id": 1,
                "factionRestriction": factionRestriction,
                "sources": [{"rep": {"repFactionId": repFactionId, "repLevel": 6}}],
            },
        ),
        enchants=(),
        random_suffixes=(),
        zones={},
        npcs={},
        factions={889: "Silverwing Sentinels", 890: "Warsong Outriders"},
        item_icons={},
        spell_icons={},
        spell_icon_rows=(),
        item_icon_rows=(),
    )


def test_wsg_rep_sources_are_corrected_to_match_the_items_own_faction_restriction():
    """The fork's db.json mines every WSG rep-sourced item under the wrong
    faction id (`_ATLASLOOT_WSG_REP_FACTION_SWAP`'s docstring has the
    upstream defect and the evidence). An alliance_only item (restriction
    1) mined under 890 (Warsong Outriders, Horde) is corrected to 889
    (Silverwing Sentinels, Alliance); a horde_only item mined under 889 is
    corrected to 890."""
    document, _ = build_loot(
        _wsg_fork(repFactionId=890, factionRestriction=1), {}, {}, {}, {1}, {}
    )
    rep = document.sources[0]
    assert (rep.faction_id, rep.name) == (889, "Silverwing Sentinels")

    document, _ = build_loot(
        _wsg_fork(repFactionId=889, factionRestriction=2), {}, {}, {}, {1}, {}
    )
    rep = document.sources[0]
    assert (rep.faction_id, rep.name) == (890, "Warsong Outriders")


def test_wsg_rep_source_correction_is_a_no_op_once_already_correct():
    """Proves the fix does not blindly flip 889/890 -- an item already
    mined under the faction id that matches its own restriction is left
    alone, so a future upstream fix to the swap does not get re-broken by
    this correction."""
    document, _ = build_loot(
        _wsg_fork(repFactionId=889, factionRestriction=1), {}, {}, {}, {1}, {}
    )
    assert document.sources[0].faction_id == 889

    document, _ = build_loot(
        _wsg_fork(repFactionId=890, factionRestriction=2), {}, {}, {}, {1}, {}
    )
    assert document.sources[0].faction_id == 890


def test_non_wsg_rep_sources_are_never_touched_by_the_correction():
    """Arathi Basin/Alterac Valley ids (or anything else) are outside
    `_ATLASLOOT_WSG_REP_FACTION_SWAP` and pass through untouched regardless
    of factionRestriction."""
    fork = ForkDatabase(
        items=(
            {
                "id": 1,
                "factionRestriction": 2,
                "sources": [{"rep": {"repFactionId": 509, "repLevel": 6}}],
            },
        ),
        enchants=(),
        random_suffixes=(),
        zones={},
        npcs={},
        factions={509: "The League of Arathor"},
        item_icons={},
        spell_icons={},
        spell_icon_rows=(),
        item_icon_rows=(),
    )
    document, _ = build_loot(fork, {}, {}, {}, {1}, {})
    assert document.sources[0].faction_id == 509


def test_two_world_bosses_with_the_same_name_raise_instead_of_silently_colliding():
    """`apply_overlays` keys its sources by id (`by_id = {source.id:
    source for ...}`); two world npcs the fork names identically would
    both slugify to the same `world:<slug>` id and one would silently
    vanish. `build_loot` must refuse instead, the same policy this module
    already applies to a dropped kind or an absent item."""
    fork = ForkDatabase(
        items=(
            {"id": 1, "sources": [{"drop": {"npcId": 10, "zoneId": 0}}]},
            {"id": 2, "sources": [{"drop": {"npcId": 20, "zoneId": 0}}]},
        ),
        enchants=(),
        random_suffixes=(),
        zones={},
        npcs={10: "Doomsayer", 20: "Doomsayer"},
        factions={},
        item_icons={},
        spell_icons={},
        spell_icon_rows=(),
        item_icon_rows=(),
    )
    with pytest.raises(SourceIdCollision, match="world:doomsayer"):
        build_loot(fork, {}, {}, {}, {1, 2}, {})
