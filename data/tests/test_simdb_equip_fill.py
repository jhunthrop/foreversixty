"""On-equip stats the client rows lose (sim-accuracy round 5).

A server-hotfix item's ItemSparse row reaches the client, but its
ItemXItemEffect link does not: neither the shipped table nor wago's hotfixed
one carries it. 153 client-path items in simdb therefore lost their on-equip
spell damage, healing or attack power (Staff of Jordan 873, Enhancement's
band-40 staff; once wago's newer rows land, Glowing Brightwood Staff 812, the
Restoration shaman's band-60 main hand, "+50 spell damage, +150 healing" on
Wowhead Forever). Wowhead Forever's feed states those amounts.
"""

from __future__ import annotations

from pipeline import wowhead_items as wh
from pipeline.simdb.equip_fill import with_wowhead_equip_stats
from pipeline.simdb.statmap import stat_array, stat_index
from pipeline.simproto import pb

FACTORS = {"hit": 10.0, "crit": 14.0, "dodge": 12.0, "parry": 15.0, "block": 5.0, "defense": 1.0}


def _wowhead(item_id: int, stats: dict[str, float], armor: int = 0) -> wh.WowheadItem:
    return wh.WowheadItem(
        id=item_id,
        name="Glowing Brightwood Staff",
        quality=3,
        item_level=50,
        required_level=49,
        class_id=2,
        subclass_id=10,
        inventory_type=17,
        icon="x",
        class_mask=None,
        stats=stats,
        set_id=None,
        unique=False,
        armor=armor,
        damage_min=94,
        damage_max=142,
        speed=3.1,
        dps=38.06,
    )


STAFF = _wowhead(812, {"int": 29, "spi": 9, "sta": 15, "spldmg": 50, "splheal": 150}, armor=99)


def _client() -> pb.SimItem:
    return pb.SimItem(
        id=812,
        stats=stat_array([("intellect", 29.0), ("spirit", 9.0), ("stamina", 15.0), ("armor", 7.0)]),
    )


def test_an_unlinked_client_item_takes_wowheads_missing_equip_stats():
    (out,) = with_wowhead_equip_stats([_client()], [STAFF], FACTORS)
    assert out.stats[stat_index("spell_power")] == 50.0
    assert out.stats[stat_index("healing")] == 150.0
    assert out.stats[stat_index("intellect")] == 29.0


def test_the_client_armor_and_stats_it_states_are_kept():
    (out,) = with_wowhead_equip_stats([_client()], [STAFF], FACTORS)
    assert out.stats[stat_index("armor")] == 7.0


def test_a_stat_the_client_states_is_never_overridden():
    client = pb.SimItem(id=812, stats=stat_array([("intellect", 31.0), ("spell_power", 12.0)]))
    (out,) = with_wowhead_equip_stats([client], [STAFF], FACTORS)
    assert out.stats[stat_index("intellect")] == 31.0
    assert out.stats[stat_index("spell_power")] == 12.0
    assert out.stats[stat_index("healing")] == 150.0


def test_an_item_wowhead_does_not_list_is_unchanged():
    client = _client()
    (out,) = with_wowhead_equip_stats([client], [], FACTORS)
    assert out == client
