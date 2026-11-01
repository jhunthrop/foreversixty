"""A shield's own block value (sim-accuracy round 5, physical lane).

The 1.60 client states no shield block anywhere: ItemSparse has no column for
it and `gametables/shieldblockregular.txt` comes back empty. Wowhead Forever's
server tooltip does ("44 Block" on Evergreen Shield 279262, the Protection
warrior's band-60 shield, read 2026-11-01), and its feed's `blockamount` is that
total, equip bonuses included (The Immovable Object 19321: 66 against its 25
on-equip). simdb gave every shield 0, so Shield Slam ("increased by your Block
Value") and every block absorbed only Strength / 20.
"""

from __future__ import annotations

from pipeline import wowhead_items as wh
from pipeline.classicdb_items import ClassicDbItem
from pipeline.simdb.shield_block import (
    classicdb_shield_block_values,
    shield_block_values,
    with_shield_block,
)
from pipeline.simdb.statmap import stat_array, stat_index
from pipeline.simproto import pb

BLOCK = stat_index("block_value")


def _wowhead(item_id: int, class_id: int, subclass_id: int, block: float) -> wh.WowheadItem:
    return wh.WowheadItem(
        id=item_id,
        name="x",
        quality=3,
        item_level=60,
        required_level=60,
        class_id=class_id,
        subclass_id=subclass_id,
        inventory_type=14,
        icon="x",
        class_mask=None,
        stats={"blockamount": block},
        set_id=None,
        unique=False,
        armor=0,
        damage_min=0,
        damage_max=0,
        speed=0.0,
        dps=0.0,
    )


def test_only_shields_carry_a_block_value():
    values = shield_block_values([_wowhead(279262, 4, 6, 44), _wowhead(16732, 4, 4, 17)])
    assert values == {279262: 44.0}


def test_a_shield_takes_wowheads_total_block_over_its_equip_part():
    shield = pb.SimItem(id=19321, stats=stat_array([("block_value", 25.0)]))
    other = pb.SimItem(id=16732, stats=stat_array([("block_value", 17.0)]))
    out = {item.id: item for item in with_shield_block([shield, other], {19321: 66.0})}
    assert out[19321].stats[BLOCK] == 66.0
    assert out[16732].stats[BLOCK] == 17.0


def test_a_shield_with_no_stats_gains_its_block():
    shield = pb.SimItem(id=279262, stats=stat_array([("stamina", 16.0)]))
    (out,) = with_shield_block([shield], {279262: 44.0})
    assert out.stats[BLOCK] == 44.0
    assert out.stats[stat_index("stamina")] == 16.0


def _classicdb(item_id: int, subclass_id: int, block: int) -> ClassicDbItem:
    return ClassicDbItem(
        id=item_id,
        name="x",
        quality=4,
        item_level=59,
        required_level=54,
        class_id=4,
        subclass_id=subclass_id,
        inventory_type=14,
        allowable_class=-1,
        allowable_race=-1,
        armor=2256,
        raw_stats={},
        resistances={},
        damage_min=0,
        damage_max=0,
        delay=0,
        set_id=None,
        unique=False,
        spells=[],
        block=block,
    )


def test_classic_db_shields_carry_their_item_template_block():
    """A shield Wowhead's feed does not carry still has its classic-db block
    (Skullflame Shield 1168: 40); a non-shield's column is ignored."""
    values = classicdb_shield_block_values([_classicdb(1168, 6, 40), _classicdb(2, 4, 9)])
    assert values == {1168: 40.0}
