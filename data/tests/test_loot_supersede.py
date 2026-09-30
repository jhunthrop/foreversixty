# data/tests/test_loot_supersede.py
"""data-followups-4 lane, 2026-09-30: ninth wow-player sweep finding 3
(`day3/player-review-24/casters.md`) -- "Grand Marshal's Stave" 18873
(ilvl 78, client vanilla id) and 234571 (ilvl 80, Forever's re-itemised
copy, +28 crit rating) are BOTH real client rows, both listed by
`pvp:rank-18:alliance`. `apply_reitemisation` never touches this pair
(both already independently sourced -- it only fills a wholly unsourced
new id), so a second mechanism is needed: `superseded_pairs` finds a pair
BOTH of whose ids are already co-listed in the same source, and
`apply_supersession` drops the legacy id from wherever that is true.
"""

import json

from pipeline.loot.reitemise import ITEM_LEVEL_TOLERANCE
from pipeline.loot.supersede import (
    apply_supersession,
    mark_superseded_items,
    superseded_pairs,
)
from pipeline.models import ClassItems, GearItem, Item, LootBoss, LootFile, LootSource

# Real ids/names/levels off build 1.60.1.70009's own items.json (item 1's
# own verification: `reitemised_pairs` already pairs these two on that
# build -- see the lane report).
LEGACY_STAVE = 18873
NEW_STAVE = 234571
LEGACY_WARSTAFF = 18874
NEW_WARSTAFF = 234549

ITEM_ROWS = [
    {"id": LEGACY_STAVE, "name": "Grand Marshal's Stave", "inventory_type": 17,
     "item_level": 78, "class_id": 2, "subclass_id": 10},
    {"id": NEW_STAVE, "name": "Grand Marshal's Stave", "inventory_type": 17,
     "item_level": 80, "class_id": 2, "subclass_id": 10},
    {"id": LEGACY_WARSTAFF, "name": "High Warlord's War Staff", "inventory_type": 17,
     "item_level": 78, "class_id": 2, "subclass_id": 10},
    {"id": NEW_WARSTAFF, "name": "High Warlord's War Staff", "inventory_type": 17,
     "item_level": 80, "class_id": 2, "subclass_id": 10},
]

# A pair `reitemised_pairs` itself would reject (item-level gap past
# `ITEM_LEVEL_TOLERANCE`) -- `superseded_pairs` must inherit that gate,
# not just co-listing, or it would supersede two unrelated items that
# merely share a name and slot.
UNRELATED_CLASSIC = 40001
UNRELATED_NEW = 240001
UNRELATED_ROWS = [
    {"id": UNRELATED_CLASSIC, "name": "Swamp Ring", "inventory_type": 11,
     "item_level": 57, "class_id": 4, "subclass_id": 0},
    {"id": UNRELATED_NEW, "name": "Swamp Ring", "inventory_type": 11,
     "item_level": 57 - ITEM_LEVEL_TOLERANCE - 1, "class_id": 4, "subclass_id": 0},
]


def _pvp_source(item_ids: list[int]) -> LootSource:
    return LootSource(
        id="pvp:rank-18:alliance", kind="pvp", name="Rank 18 (Alliance)",
        rank=18, faction="alliance", items=item_ids,
    )


def test_superseded_pairs_finds_a_pair_co_listed_in_the_same_source():
    document = LootFile(sources=[_pvp_source([LEGACY_STAVE, NEW_STAVE])])
    assert superseded_pairs(document, ITEM_ROWS) == {LEGACY_STAVE: NEW_STAVE}


def test_superseded_pairs_finds_both_staff_pairs_independently():
    document = LootFile(
        sources=[_pvp_source([LEGACY_STAVE, NEW_STAVE, LEGACY_WARSTAFF, NEW_WARSTAFF])]
    )
    assert superseded_pairs(document, ITEM_ROWS) == {
        LEGACY_STAVE: NEW_STAVE,
        LEGACY_WARSTAFF: NEW_WARSTAFF,
    }


def test_superseded_pairs_ignores_ids_never_co_listed_anywhere():
    """Both ids exist in the catalogue but neither appears in any source
    (unsourced entirely) -- nothing to supersede, since nothing proves
    they are duplicate listings for the same reward."""
    document = LootFile(sources=[])
    assert superseded_pairs(document, ITEM_ROWS) == {}


def test_superseded_pairs_ignores_a_pair_in_two_different_sources():
    """The legacy id and its copy each real, but in DIFFERENT sources --
    not proof of a duplicate listing the way co-occurrence in the SAME
    bucket is, so left alone."""
    document = LootFile(
        sources=[
            LootSource(id="vendor:1", kind="vendor", name="V1", npc_id=1, items=[LEGACY_STAVE]),
            LootSource(id="vendor:2", kind="vendor", name="V2", npc_id=2, items=[NEW_STAVE]),
        ]
    )
    assert superseded_pairs(document, ITEM_ROWS) == {}


def test_superseded_pairs_respects_reitemised_pairs_own_gate():
    """A name+slot match `reitemised_pairs` itself rejects (item-level gap
    past `ITEM_LEVEL_TOLERANCE`, the Swamp Ring shape) is never
    superseded even when both ids are co-listed."""
    document = LootFile(
        sources=[_pvp_source([UNRELATED_CLASSIC, UNRELATED_NEW])]
    )
    assert superseded_pairs(document, UNRELATED_ROWS) == {}


def test_apply_supersession_drops_the_legacy_id_from_a_shared_flat_items_list():
    document = LootFile(sources=[_pvp_source([12584, LEGACY_STAVE, NEW_STAVE, 18876])])
    updated, removed = apply_supersession(document, {LEGACY_STAVE: NEW_STAVE})
    assert removed == 1
    assert updated.sources[0].items == [12584, NEW_STAVE, 18876]


def test_apply_supersession_drops_the_legacy_id_from_a_boss():
    document = LootFile(
        sources=[
            LootSource(
                id="raid:x", kind="raid", name="X", zone_id=1,
                bosses=[
                    LootBoss(
                        id="raid:x:1", name="Boss", npc_id=1,
                        items=[LEGACY_STAVE, NEW_STAVE],
                    )
                ],
            )
        ]
    )
    updated, removed = apply_supersession(document, {LEGACY_STAVE: NEW_STAVE})
    assert removed == 1
    assert updated.sources[0].bosses[0].items == [NEW_STAVE]


def test_apply_supersession_leaves_a_bucket_with_only_the_legacy_id_alone():
    """A pair superseded ELSEWHERE (globally) is still left alone in a
    bucket that names only the legacy id -- there is no copy present to
    de-duplicate against here, so removing it would be a real data loss,
    not a de-duplication."""
    document = LootFile(
        sources=[LootSource(id="vendor:3", kind="vendor", name="V3", npc_id=3,
                             items=[LEGACY_STAVE])]
    )
    updated, removed = apply_supersession(document, {LEGACY_STAVE: NEW_STAVE})
    assert removed == 0
    assert updated.sources[0].items == [LEGACY_STAVE]


def test_apply_supersession_is_a_noop_with_no_pairs():
    document = LootFile(sources=[_pvp_source([LEGACY_STAVE, NEW_STAVE])])
    updated, removed = apply_supersession(document, {})
    assert removed == 0
    assert updated is document


def test_mark_superseded_items_writes_the_flat_and_per_class_rows(tmp_path):
    build_dir = tmp_path
    (build_dir / "items.json").write_text(
        json.dumps(
            [
                Item(
                    id=LEGACY_STAVE, name="Grand Marshal's Stave", quality=4, item_level=78,
                    required_level=60, class_id=2, subclass_id=10, inventory_type=17,
                ).model_dump(),
                Item(
                    id=NEW_STAVE, name="Grand Marshal's Stave", quality=4, item_level=80,
                    required_level=60, class_id=2, subclass_id=10, inventory_type=17,
                ).model_dump(),
            ]
        ),
        encoding="utf-8",
    )
    items_dir = build_dir / "items"
    items_dir.mkdir()
    (items_dir / "mage.json").write_text(
        json.dumps(
            ClassItems(
                build="test", class_slug="mage",
                items=[
                    GearItem(
                        id=LEGACY_STAVE, name="Grand Marshal's Stave", icon="inv_staff_14",
                        slot="main_hand", quality=4, required_level=60,
                        required_level_source="client", item_level=78, armor=0,
                        stats={"spell_power": 71}, set_id=None, unique=False,
                    ).model_dump(),
                    GearItem(
                        id=NEW_STAVE, name="Grand Marshal's Stave", icon="inv_staff_14",
                        slot="main_hand", quality=4, required_level=60,
                        required_level_source="client", item_level=80, armor=0,
                        stats={"spell_power": 81, "crit": 28}, set_id=None, unique=False,
                    ).model_dump(),
                ],
            ).model_dump()
        ),
        encoding="utf-8",
    )

    marked = mark_superseded_items(build_dir, {LEGACY_STAVE: NEW_STAVE})
    assert marked == 1

    flat = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    by_id = {row["id"]: row for row in flat}
    assert by_id[LEGACY_STAVE]["superseded_by"] == NEW_STAVE
    assert by_id[NEW_STAVE]["superseded_by"] is None

    per_class = json.loads((items_dir / "mage.json").read_text(encoding="utf-8"))
    by_class_id = {row["id"]: row for row in per_class["items"]}
    assert by_class_id[LEGACY_STAVE]["superseded_by"] == NEW_STAVE
    assert by_class_id[NEW_STAVE]["superseded_by"] is None


def test_mark_superseded_items_is_a_noop_with_no_pairs(tmp_path):
    build_dir = tmp_path
    items_path = build_dir / "items.json"
    items_path.write_text("[]", encoding="utf-8")
    items_dir = build_dir / "items"
    items_dir.mkdir()
    before = items_path.read_bytes()
    marked = mark_superseded_items(build_dir, {})
    assert marked == 0
    assert items_path.read_bytes() == before
