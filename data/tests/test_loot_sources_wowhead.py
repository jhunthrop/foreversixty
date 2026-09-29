# data/tests/test_loot_sources_wowhead.py
"""night-item-sources lane, 2026-09-28: `pipeline.loot.sources.build_loot`
folding in `pipeline.item_sources`' wowhead scrape for an item the fork
database itself names no source for at all (item 110, "Suffixed Sword",
in tests/fixtures/loot/items.json -- its fork db.json row states
`sources: null`).
"""

import json
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.forkdb import load_fork_database
from pipeline.item_sources import ItemSourceEntry
from pipeline.loot.sources import build_loot, instance_types, pvp_ranks, source_item_ids
from pipeline.loot.wowhead import (
    FOREVER_NEW_ID_THRESHOLD,
    is_placeholder_item,
    load_class_item_rows,
    named_items_in_committed_loot,
    unsourced_real_item_ids,
)
from pipeline.wowhead_item_sources import CraftedSource, NpcSource, QuestRewardSource

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"

UNSOURCED_ITEM = 110  # "Suffixed Sword" -- the fixture's one fork-unsourced item


def built(item_sources=None):
    fork = load_fork_database(ENGINE)
    rows = json.loads((ENGINE / "zones.json").read_text(encoding="utf-8"))
    item_rows = json.loads((ENGINE / "items.json").read_text(encoding="utf-8"))
    build_items = {row["id"] for row in item_rows}
    item_inventory_types = {row["id"]: row["inventory_type"] for row in item_rows}
    return build_loot(
        fork,
        {row["id"]: row["name"] for row in rows},
        instance_types(read_csv(ENGINE / "Map.csv"), rows),
        pvp_ranks(read_csv(ENGINE / "ItemSparse.csv")),
        build_items,
        item_inventory_types,
        None,
        item_sources,
    )


def source(document, source_id: str):
    for candidate in document.sources:
        if candidate.id == source_id:
            return candidate
    raise AssertionError(f"no source {source_id}; have {[s.id for s in document.sources]}")


def test_with_no_item_sources_the_item_stays_unsourced():
    document, stats = built()
    named = {item_id for s in document.sources for item_id in source_item_ids(s)}
    assert UNSOURCED_ITEM not in named
    assert stats.wowhead_items == 0


def test_a_wholly_new_wowhead_vendor_source_is_appended_and_tagged():
    item_sources = {
        UNSOURCED_ITEM: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=950, name="Wowhead Vendor", zone_ids=[16])],
            source="wowhead",
            fetched_at="x",
        )
    }
    document, stats = built(item_sources)
    vendor = source(document, "vendor:950")
    assert vendor.kind == "vendor"
    assert vendor.name == "Wowhead Vendor"
    assert vendor.items == [UNSOURCED_ITEM]
    assert vendor.source_origin == "wowhead"
    assert stats.wowhead_items == 1


def test_a_wowhead_drop_in_an_existing_zone_unions_into_the_fork_source_and_stays_unset():
    """Item 104 ("World Boss Drop") already sources Azuregos's own
    `zone:16` bucket (npc 903 names no `zoneId` on ITS OWN drop row, so
    it is `world`-only -- but the OPEN-WORLD zone bucket itself, `zone:
    16`, still exists because of it). A wowhead drop naming the SAME
    zone id for item 110 has to land in the SAME `zone:16` source,
    unioned, not a second one -- and the union keeps `source_origin`
    unset, since the fork found this zone bucket first.
    """
    item_sources = {
        UNSOURCED_ITEM: ItemSourceEntry(
            dropped_by=[NpcSource(npc_id=0, name="", zone_ids=[16])],
            source="wowhead",
            fetched_at="x",
        )
    }
    document, stats = built(item_sources)
    zone = source(document, "zone:16")
    assert set(zone.items) == {104, UNSOURCED_ITEM}
    assert zone.source_origin is None
    assert stats.wowhead_items == 1


def test_a_wowhead_crafted_addition_unions_into_the_existing_profession_bucket():
    item_sources = {
        UNSOURCED_ITEM: ItemSourceEntry(
            crafted_by=[CraftedSource(profession="blacksmithing")],
            source="wowhead",
            fetched_at="x",
        )
    }
    document, _ = built(item_sources)
    crafted = source(document, "crafted:blacksmithing")
    assert set(crafted.items) == {106, UNSOURCED_ITEM}
    assert crafted.source_origin is None  # the fork already had this bucket


def test_a_wowhead_quest_reward_joins_the_flat_quest_bucket_with_wowhead_level_source():
    item_sources = {
        UNSOURCED_ITEM: ItemSourceEntry(
            quest_rewards=[
                QuestRewardSource(
                    quest_id=99, name="A Wowhead Errand", min_level=38, level=40, faction="horde"
                )
            ],
            source="wowhead",
            fetched_at="x",
        )
    }
    document, _ = built(item_sources)
    quest_source = source(document, "quest")
    assert UNSOURCED_ITEM in quest_source.items
    detail = document.quests[str(UNSOURCED_ITEM)]
    assert len(detail) == 1
    assert detail[0].quest_id == 99
    assert detail[0].faction == "horde"
    assert detail[0].level_source == "wowhead"


def test_a_wowhead_source_for_an_item_the_build_does_not_have_is_ignored():
    item_sources = {
        999999: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=950, name="Wowhead Vendor", zone_ids=[16])],
            source="wowhead",
            fetched_at="x",
        )
    }
    document, stats = built(item_sources)
    assert not any(s.id == "vendor:950" for s in document.sources)
    assert stats.wowhead_items == 0


def test_is_placeholder_item_matches_the_client_naming_conventions():
    assert is_placeholder_item({"name": "90 Epic Rogue Dagger"})
    assert is_placeholder_item({"name": "Bland Dagger"})
    assert is_placeholder_item({"name": "Copy of Truestrike Blade"})
    assert not is_placeholder_item(
        {"name": "Truestrike Blade", "armor": 0, "damage_max": 10, "stats": {}}
    )


def test_is_placeholder_item_also_catches_a_zero_stat_row_by_shape():
    assert is_placeholder_item(
        {
            "name": "Some Real-Looking Name",
            "armor": 0,
            "damage_max": 0,
            "stats": {},
            "effect_text": "",
            "set_id": None,
        }
    )


def test_unsourced_real_item_ids_excludes_named_and_placeholder_items():
    rows = [
        {"id": 1, "name": "Real Sword", "armor": 0, "damage_max": 5, "stats": {}},
        {"id": 2, "name": "90 Epic Rogue Dagger", "armor": 0, "damage_max": 0, "stats": {}},
        {"id": 3, "name": "Another Real Item", "armor": 5, "damage_max": 0, "stats": {}},
    ]
    assert unsourced_real_item_ids(rows, named={3}) == [1]


def test_unsourced_real_item_ids_orders_forever_new_before_classic_by_required_level():
    """src-crawl-order lane, 2026-09-29: every Forever-new id (>=
    FOREVER_NEW_ID_THRESHOLD) sorts before every Classic id, ascending
    required_level within each bucket, unset (0, or the key entirely
    missing) required_level sorting last within its own bucket -- NOT
    the ascending-item-id order these rows are given in."""
    assert FOREVER_NEW_ID_THRESHOLD == 200_000
    rows = [
        {"id": 100, "name": "Classic High", "armor": 1, "required_level": 40},
        {"id": 50, "name": "Classic Low", "armor": 1, "required_level": 5},
        {"id": 60, "name": "Classic Unset", "armor": 1, "required_level": 0},
        {"id": 70, "name": "Classic No Level Key", "armor": 1},
        {"id": 200500, "name": "Forever-new High", "armor": 1, "required_level": 30},
        {"id": 200050, "name": "Forever-new Low", "armor": 1, "required_level": 10},
        {"id": 200100, "name": "Forever-new Unset", "armor": 1, "required_level": 0},
    ]
    assert unsourced_real_item_ids(rows, named=set()) == [
        200050,  # Forever-new bucket, ascending required_level
        200500,
        200100,  # Forever-new bucket, unset required_level -- last
        50,  # Classic bucket, ascending required_level
        100,
        60,  # Classic bucket, unset required_level -- last (0 and missing tie on id)
        70,
    ]


def test_named_items_in_committed_loot_reads_every_bucket_shape(tmp_path: Path):
    build_dir = tmp_path / "1.60.1.70009"
    build_dir.mkdir()
    (build_dir / "loot.json").write_text(
        json.dumps(
            {
                "sources": [
                    {"id": "vendor:1", "kind": "vendor", "name": "V", "items": [1, 2]},
                    {"id": "zone:1", "kind": "zone", "name": "Z", "items": [3], "trash": [4]},
                    {
                        "id": "raid:x",
                        "kind": "raid",
                        "name": "X",
                        "bosses": [{"id": "raid:x:1", "name": "B", "npc_id": 1, "items": [5]}],
                    },
                ]
            }
        ),
        encoding="utf-8",
    )
    assert named_items_in_committed_loot(build_dir) == {1, 2, 3, 4, 5}


def test_named_items_in_committed_loot_with_no_file_returns_empty(tmp_path: Path):
    assert named_items_in_committed_loot(tmp_path / "1.60.1.70009") == set()


def test_load_class_item_rows_dedupes_an_item_shared_by_two_classes(tmp_path: Path):
    build_dir = tmp_path / "1.60.1.70009"
    items_dir = build_dir / "items"
    items_dir.mkdir(parents=True)
    row = {"id": 1, "name": "Shared Item", "armor": 0, "damage_max": 5, "stats": {}}
    (items_dir / "warrior.json").write_text(
        json.dumps({"items": [row, {"id": 2, "name": "Warrior Only"}]}), encoding="utf-8"
    )
    (items_dir / "rogue.json").write_text(json.dumps({"items": [row]}), encoding="utf-8")
    rows = load_class_item_rows(build_dir)
    assert {r["id"] for r in rows} == {1, 2}
