# data/tests/test_loot_reitemise.py
"""src-classicdb lane item 3, 2026-09-29: a Forever-new item (id >=
200000) that shares its name and slot with an existing Classic item
inherits the classic item's own sources when it has none of its own --
but only when it is plausibly the same item re-tuned (src-classicdb-fixes
lane, 2026-09-29): item levels within `ITEM_LEVEL_TOLERANCE` and the same
armor/weapon subclass, or nothing is inherited."""

from pipeline.loot.reitemise import (
    ITEM_LEVEL_TOLERANCE,
    apply_reitemisation,
    reitemised_pairs,
)
from pipeline.models import LootBoss, LootFile, LootSource, QuestSource

CLASSIC_ID = 22016
NEW_ID = 226884
AMBIGUOUS_CLASSIC_A = 18854
AMBIGUOUS_CLASSIC_B = 18856
AMBIGUOUS_NEW = 209611

# Same item level and armor subclass as the classic item they pair with,
# so the item-level/subclass gate never rejects these unless a test says
# otherwise -- the gate is exercised on its own by the Swamp Ring rows
# below.
ITEM_ROWS = [
    {"id": CLASSIC_ID, "name": "Beastmaster's Mantle", "inventory_type": 3,
     "item_level": 61, "class_id": 4, "subclass_id": 1},
    {"id": NEW_ID, "name": "Beastmaster's Mantle", "inventory_type": 3,
     "item_level": 61, "class_id": 4, "subclass_id": 1},
    {"id": AMBIGUOUS_CLASSIC_A, "name": "Insignia of the Alliance", "inventory_type": 12,
     "item_level": 26, "class_id": 4, "subclass_id": 0},
    {"id": AMBIGUOUS_CLASSIC_B, "name": "Insignia of the Alliance", "inventory_type": 12,
     "item_level": 30, "class_id": 4, "subclass_id": 0},
    {"id": AMBIGUOUS_NEW, "name": "Insignia of the Alliance", "inventory_type": 12,
     "item_level": 30, "class_id": 4, "subclass_id": 0},
    {"id": 999999, "name": "Unrelated New Item", "inventory_type": 3,
     "item_level": 10, "class_id": 4, "subclass_id": 1},
]

# The lane brief's own defect: Swamp Ring 270052 (Forever-new, ilvl 35, a
# leveling ring) merely reuses the name+slot of Swamp Ring 12015 (Classic,
# ilvl 57, required level 52, a Scholomance reward) -- the two are
# unrelated content and must NOT inherit each other's sources.
SWAMP_RING_CLASSIC_ID = 12015
SWAMP_RING_NEW_ID = 270052
SWAMP_RING_ROWS = [
    {"id": SWAMP_RING_CLASSIC_ID, "name": "Swamp Ring", "inventory_type": 11,
     "item_level": 57, "class_id": 4, "subclass_id": 0},
    {"id": SWAMP_RING_NEW_ID, "name": "Swamp Ring", "inventory_type": 11,
     "item_level": 35, "class_id": 4, "subclass_id": 0},
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
         "item_level": 61, "class_id": 4, "subclass_id": 1, "faction_restriction": ""},
        {"id": NEW_ID, "name": "Beastmaster's Mantle", "inventory_type": 3,
         "item_level": 61, "class_id": 4, "subclass_id": 1, "faction_restriction": "horde_only"},
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


def test_reitemised_pairs_rejects_the_swamp_ring_pair_the_item_level_gap_is_too_wide():
    """Swamp Ring 270052 (ilvl 35) must not inherit Swamp Ring 12015's
    (ilvl 57) sources: a 22-ilvl gap exceeds `ITEM_LEVEL_TOLERANCE`."""
    pairs = reitemised_pairs(SWAMP_RING_ROWS)
    assert SWAMP_RING_NEW_ID not in pairs


def test_apply_reitemisation_never_copies_the_swamp_rings_scholomance_drop():
    document = LootFile(
        sources=[
            LootSource(
                id="dungeon:scholomance", kind="dungeon", name="Scholomance", zone_id=289,
                bosses=[
                    LootBoss(
                        id="dungeon:scholomance:darkmaster-gandling",
                        name="Darkmaster Gandling", npc_id=1853,
                        items=[SWAMP_RING_CLASSIC_ID],
                    )
                ],
            )
        ]
    )
    updated, filled = apply_reitemisation(document, SWAMP_RING_ROWS)
    assert filled == 0
    assert updated == document


def test_reitemised_pairs_accepts_a_pair_exactly_at_the_item_level_tolerance():
    rows = [
        {"id": CLASSIC_ID, "name": "Beastmaster's Mantle", "inventory_type": 3,
         "item_level": 50, "class_id": 4, "subclass_id": 1},
        {"id": NEW_ID, "name": "Beastmaster's Mantle", "inventory_type": 3,
         "item_level": 50 + ITEM_LEVEL_TOLERANCE, "class_id": 4, "subclass_id": 1},
    ]
    pairs = reitemised_pairs(rows)
    assert pairs[NEW_ID] == CLASSIC_ID


def test_reitemised_pairs_rejects_a_matching_item_level_but_different_subclass():
    """Same name, slot and item level, but a plate item reusing a cloth
    item's name is not the same piece of content."""
    rows = [
        {"id": CLASSIC_ID, "name": "Beastmaster's Mantle", "inventory_type": 3,
         "item_level": 61, "class_id": 4, "subclass_id": 1},
        {"id": NEW_ID, "name": "Beastmaster's Mantle", "inventory_type": 3,
         "item_level": 61, "class_id": 4, "subclass_id": 4},
    ]
    pairs = reitemised_pairs(rows)
    assert NEW_ID not in pairs
