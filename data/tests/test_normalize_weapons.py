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


def test_a_ranged_two_hander_is_flagged():
    """A bow occupies the ranged slot and both hands; rule 6 has to know."""
    assert weapon_fields(row(InventoryType=15)).two_hand is True


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
