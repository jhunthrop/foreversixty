import pytest

from pipeline.normalize.gear import TWO_HAND_INVENTORY_TYPES, weapon_fields


def row(**overrides):
    base = {
        "ID": "12640",
        "InventoryType": "13",
        "ItemDelay": "2600",
        "ItemDamageMin_0": "100",
        "ItemDamageMax_0": "180",
    }
    base.update({key: str(value) for key, value in overrides.items()})
    return base


def test_a_one_hander_reports_damage_speed_and_dps():
    fields = weapon_fields(row())
    assert fields.damage_min == 100
    assert fields.damage_max == 180
    assert fields.speed == 2.6
    assert fields.dps == pytest.approx(53.85, abs=0.01)
    assert fields.two_hand is False


def test_a_two_hander_is_flagged():
    assert weapon_fields(row(InventoryType=17)).two_hand is True
    assert 17 in TWO_HAND_INVENTORY_TYPES


def test_a_ranged_weapon_is_not_flagged():
    """R10: rule 6 asks whether an off-hand item may sit beside this one.

    A bow (InventoryType 15) takes the ranged slot, not the main hand, so it
    never refuses one. 26 is the trap the old set fell into -- in Classic it is
    Ranged Right, which is wands as much as guns and crossbows, so flagging it
    called 37 wands two-handed.
    """
    assert weapon_fields(row(InventoryType=15)).two_hand is False
    assert weapon_fields(row(InventoryType=26)).two_hand is False
    assert TWO_HAND_INVENTORY_TYPES.isdisjoint({15, 25, 26})


def test_a_non_weapon_reports_zeroes_and_no_flag():
    fields = weapon_fields(row(InventoryType=1, ItemDelay=0, ItemDamageMin_0=0, ItemDamageMax_0=0))
    assert (fields.damage_min, fields.damage_max, fields.speed, fields.dps) == (0, 0, 0.0, 0.0)
    assert fields.two_hand is False


def test_a_weapon_with_no_delay_reports_no_dps_rather_than_dividing_by_zero():
    fields = weapon_fields(row(ItemDelay=0))
    assert fields.speed == 0.0
    assert fields.dps == 0.0


def test_a_build_without_the_damage_columns_reports_zeroes():
    """The 1.60 client computes weapon damage from curves; absent is not malformed."""
    sparse = {"ID": "12640", "InventoryType": "13", "ItemDelay": "2600"}
    fields = weapon_fields(sparse)
    assert (fields.damage_min, fields.damage_max) == (0, 0)
    assert fields.speed == 2.6
