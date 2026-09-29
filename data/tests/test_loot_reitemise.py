# data/tests/test_loot_reitemise.py
"""src-classicdb lane item 3, 2026-09-29: a Forever-new item (id >=
200000) that shares its name and slot with an existing Classic item
inherits the classic item's own sources when it has none of its own."""

from pipeline.loot.reitemise import apply_reitemisation, reitemised_pairs
from pipeline.models import LootBoss, LootFile, LootSource, QuestSource

CLASSIC_ID = 22016
NEW_ID = 226884
AMBIGUOUS_CLASSIC_A = 18854
AMBIGUOUS_CLASSIC_B = 18856
AMBIGUOUS_NEW = 209611

ITEM_ROWS = [
    {"id": CLASSIC_ID, "name": "Beastmaster's Mantle", "inventory_type": 3},
    {"id": NEW_ID, "name": "Beastmaster's Mantle", "inventory_type": 3},
    {"id": AMBIGUOUS_CLASSIC_A, "name": "Insignia of the Alliance", "inventory_type": 12},
    {"id": AMBIGUOUS_CLASSIC_B, "name": "Insignia of the Alliance", "inventory_type": 12},
    {"id": AMBIGUOUS_NEW, "name": "Insignia of the Alliance", "inventory_type": 12},
    {"id": 999999, "name": "Unrelated New Item", "inventory_type": 3},
]


def test_reitemised_pairs_matches_an_unambiguous_name_and_slot():
    pairs = reitemised_pairs(ITEM_ROWS)
    assert pairs[NEW_ID] == CLASSIC_ID


def test_reitemised_pairs_skips_an_ambiguous_name_and_slot():
    pairs = reitemised_pairs(ITEM_ROWS)
    assert AMBIGUOUS_NEW not in pairs


def test_reitemised_pairs_skips_a_new_item_with_no_classic_match():
    pairs = reitemised_pairs(ITEM_ROWS)
    assert 999999 not in pairs


def test_apply_reitemisation_copies_a_boss_drop_onto_the_new_item():
    document = LootFile(
        sources=[
            LootSource(
                id="raid:x", kind="raid", name="X", zone_id=1,
                bosses=[LootBoss(id="raid:x:1", name="Boss", npc_id=1, items=[CLASSIC_ID])],
            )
        ]
    )
    updated, filled = apply_reitemisation(document, ITEM_ROWS)
    assert filled == 1
    boss = updated.sources[0].bosses[0]
    assert set(boss.items) == {CLASSIC_ID, NEW_ID}
    assert boss.reitemised_from == {str(NEW_ID): CLASSIC_ID}


def test_apply_reitemisation_copies_a_flat_items_list_and_a_vendor_condition():
    document = LootFile(
        sources=[
            LootSource(
                id="vendor:5", kind="vendor", name="V", npc_id=5, items=[CLASSIC_ID],
                faction_id=76, standing="exalted",
            )
        ]
    )
    updated, filled = apply_reitemisation(document, ITEM_ROWS)
    assert filled == 1
    vendor = updated.sources[0]
    assert set(vendor.items) == {CLASSIC_ID, NEW_ID}
    assert vendor.reitemised_from == {str(NEW_ID): CLASSIC_ID}
    assert vendor.faction_id == 76  # untouched


def test_apply_reitemisation_copies_the_quest_detail_map_too():
    document = LootFile(
        sources=[LootSource(id="quest", kind="quest", name="Quests", items=[CLASSIC_ID])],
        quests={
            str(CLASSIC_ID): [
                QuestSource(
                    quest_id=53, name="Sweet Amber", faction="both",
                    min_level=40, level=44, level_source="classic-db",
                )
            ]
        },
    )
    updated, filled = apply_reitemisation(document, ITEM_ROWS)
    assert filled == 1
    assert NEW_ID in updated.sources[0].items
    assert updated.quests[str(NEW_ID)] == updated.quests[str(CLASSIC_ID)]


def test_apply_reitemisation_re_derives_faction_from_the_new_items_own_restriction():
    """The new item can carry a DIFFERENT `faction_restriction` than the
    classic one it inherits sources from -- a copied `QuestSource` gets
    its own `faction` re-derived from the NEW item, never the classic
    item's value copied verbatim (models.QuestSource's own doc; the same
    fix pipeline.loot.classicdb/wowhead both apply for quest rewards
    they name directly)."""
    restricted_rows = [
        {"id": CLASSIC_ID, "name": "Beastmaster's Mantle", "inventory_type": 3,
         "faction_restriction": ""},
        {"id": NEW_ID, "name": "Beastmaster's Mantle", "inventory_type": 3,
         "faction_restriction": "horde_only"},
    ]
    document = LootFile(
        sources=[LootSource(id="quest", kind="quest", name="Quests", items=[CLASSIC_ID])],
        quests={
            str(CLASSIC_ID): [
                QuestSource(
                    quest_id=53, name="Sweet Amber", faction="alliance",
                    min_level=40, level=44, level_source="classic-db",
                )
            ]
        },
    )
    updated, filled = apply_reitemisation(document, restricted_rows)
    assert filled == 1
    assert updated.quests[str(NEW_ID)][0].faction == "horde"
    assert updated.quests[str(CLASSIC_ID)][0].faction == "alliance"  # untouched


def test_apply_reitemisation_never_touches_an_item_that_already_has_its_own_source():
    document = LootFile(
        sources=[
            LootSource(id="zone:1", kind="zone", name="Z", zone_id=1, items=[CLASSIC_ID]),
            LootSource(id="zone:2", kind="zone", name="Z2", zone_id=2, items=[NEW_ID]),
        ]
    )
    updated, filled = apply_reitemisation(document, ITEM_ROWS)
    assert filled == 0
    assert updated == document


def test_apply_reitemisation_does_nothing_when_the_classic_item_itself_has_no_source():
    document = LootFile(sources=[])
    updated, filled = apply_reitemisation(document, ITEM_ROWS)
    assert filled == 0
    assert updated.sources == []
