# data/tests/test_loot_vendor_cost.py
"""vendor-11036 lane, 2026-09-30: `apply_vendor_cost_gate`'s own contract
tests, on synthetic sources (never the network) -- see
`pipeline.loot.vendor_cost`'s own module doc for the rule and the
primary-source measurement that motivated it.
"""

from pipeline.item_sources import ItemSourceEntry
from pipeline.loot.vendor_cost import apply_vendor_cost_gate
from pipeline.models import LootFile, LootSource
from pipeline.wowhead_item_sources import NpcSource


def _vendor(npc_id: int, items: list[int], **extra) -> LootSource:
    return LootSource(
        id=f"vendor:{npc_id}",
        kind="vendor",
        name=f"Vendor {npc_id}",
        npc_id=npc_id,
        items=items,
        source_origin="wowhead",
        **extra,
    )


def test_phantom_cost_item_gates_the_row():
    """Item 100's `ItemExtendedCost` prices it in item 999, which
    `build_items` does not contain at all (a 404 on wowhead, 239759's
    own real-world twin) -- gated to a `vendor:1:later` sibling with a
    note, the original id keeping nothing."""
    document = LootFile(sources=[_vendor(1, [100])])
    item_sources = {
        100: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=1, name="Vendor 1", zone_ids=[], cost_item_ids=[999])],
            source="wowhead",
            fetched_at="x",
        )
    }
    new_document, moved = apply_vendor_cost_gate(document, item_sources, {100}, {100: 60}, {100: 4})
    assert moved == 1
    assert [s.id for s in new_document.sources if s.items] == ["vendor:1:later"]
    gated = next(s for s in new_document.sources if s.id == "vendor:1:later")
    assert gated.items == [100]
    assert gated.opens == "later"
    assert gated.source_note


def test_circular_cost_item_gates_both_sides():
    """Item 100 costs item 200, and item 200's OWN only source is the
    SAME vendor's row costing item 100 back -- neither is in the
    phantom seed (both ARE in `build_items`), but the fixed-point pass
    still finds neither payable: each one's only route to being owned
    is buying the other, which is itself ungated nowhere else."""
    document = LootFile(sources=[_vendor(1, [100, 200])])
    item_sources = {
        100: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=1, name="Vendor 1", zone_ids=[], cost_item_ids=[200])],
            source="wowhead",
            fetched_at="x",
        ),
        200: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=1, name="Vendor 1", zone_ids=[], cost_item_ids=[100])],
            source="wowhead",
            fetched_at="x",
        ),
    }
    new_document, moved = apply_vendor_cost_gate(
        document, item_sources, {100, 200}, {100: 60, 200: 60}, {100: 4, 200: 4}
    )
    assert moved == 2
    gated = next(s for s in new_document.sources if s.id == "vendor:1:later")
    assert gated.items == [100, 200]


def test_real_gold_price_is_never_gated():
    """An ordinary gold buy (`cost_item_ids` empty, `cost_money` set) is
    left on the original, ungated source -- the overwhelming majority of
    every vendor this build has."""
    document = LootFile(sources=[_vendor(1, [100])])
    item_sources = {
        100: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=1, name="Vendor 1", zone_ids=[], cost_money=500)],
            source="wowhead",
            fetched_at="x",
        )
    }
    new_document, moved = apply_vendor_cost_gate(document, item_sources, {100}, {100: 60}, {100: 4})
    assert moved == 0
    assert new_document == document


def test_an_alternative_unresolvable_ingredient_does_not_spoil_an_otherwise_real_item():
    """`build_items` does not need to cover every referenced cost id for
    an ORDINARY item to stay ungated -- only a row that names an
    unresolvable id as ITS OWN cost is gated, never merely because some
    other unrelated item elsewhere in the same scrape does."""
    document = LootFile(sources=[_vendor(1, [100]), _vendor(2, [200])])
    item_sources = {
        100: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=1, name="Vendor 1", zone_ids=[], cost_money=500)],
            source="wowhead",
            fetched_at="x",
        ),
        200: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=2, name="Vendor 2", zone_ids=[], cost_item_ids=[999])],
            source="wowhead",
            fetched_at="x",
        ),
    }
    new_document, moved = apply_vendor_cost_gate(
        document, item_sources, {100, 200}, {100: 60, 200: 60}, {100: 4, 200: 4}
    )
    assert moved == 1
    assert source_ids_with_items(new_document) == {"vendor:1", "vendor:2:later"}


def test_no_cost_no_rep_but_overleveled_quality_gates_as_beta_inventory():
    """Rule 4: a row with NO cost at all and no reputation requirement,
    but whose item is quality 4+ and above this build's own launch-phase
    ceiling (here, 60 -- the only ungated item level `document` names),
    reads as beta/test inventory."""
    document = LootFile(
        sources=[
            _vendor(1, [100]),
            LootSource(id="world:trainer", kind="world", name="A Trainer", items=[50]),
        ]
    )
    item_sources = {
        100: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=1, name="Vendor 1", zone_ids=[])],
            source="wowhead",
            fetched_at="x",
        )
    }
    new_document, moved = apply_vendor_cost_gate(
        document, item_sources, {50, 100}, {50: 60, 100: 98}, {50: 2, 100: 4}
    )
    assert moved == 1
    gated = next(s for s in new_document.sources if s.id == "vendor:1:later")
    assert gated.items == [100]


def test_no_cost_no_rep_and_within_the_launch_cap_stays_ungated():
    """The same shape as above, but the item's own level does not
    exceed the ceiling -- an ordinary, ungated vendor sale."""
    document = LootFile(
        sources=[
            _vendor(1, [100]),
            LootSource(id="world:trainer", kind="world", name="A Trainer", items=[50]),
        ]
    )
    item_sources = {
        100: ItemSourceEntry(
            sold_by=[NpcSource(npc_id=1, name="Vendor 1", zone_ids=[])],
            source="wowhead",
            fetched_at="x",
        )
    }
    new_document, moved = apply_vendor_cost_gate(
        document, item_sources, {50, 100}, {50: 60, 100: 60}, {50: 2, 100: 4}
    )
    assert moved == 0


def source_ids_with_items(document: LootFile) -> set[str]:
    return {s.id for s in document.sources if s.items}
