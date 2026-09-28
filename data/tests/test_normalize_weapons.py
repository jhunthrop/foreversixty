from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.normalize.gear import TWO_HAND_INVENTORY_TYPES, weapon_fields
from pipeline.normalize.weapon_curves import WeaponCurves, load_weapon_curves

FIXTURES = Path(__file__).parent / "fixtures" / "sim"

AXE_TWO_HAND = 1
MACE_ONE_HAND = 4
GUN = 3
WAND = 19


def curves() -> WeaponCurves:
    return load_weapon_curves(
        read_csv(FIXTURES / "ItemDamageOneHand.csv"),
        read_csv(FIXTURES / "ItemDamageTwoHand.csv"),
        read_csv(FIXTURES / "ItemDamageRanged.csv"),
        read_csv(FIXTURES / "ItemDamageWand.csv"),
        read_csv(FIXTURES / "ItemDamageThrown.csv"),
    )


def row(**overrides):
    base = {
        "ID": "12640",
        "InventoryType": "13",
        "ItemDelay": "2600",
        "MinDamage_0": "100",
        "MaxDamage_0": "180",
    }
    base.update({key: str(value) for key, value in overrides.items()})
    return base


def test_a_one_hander_reports_damage_speed_and_dps():
    fields = weapon_fields(row(), MACE_ONE_HAND)
    assert fields.damage_min == 100
    assert fields.damage_max == 180
    assert fields.speed == 2.6
    assert fields.dps == pytest.approx(53.85, abs=0.01)
    assert fields.two_hand is False


def test_a_two_hander_is_flagged():
    assert weapon_fields(row(InventoryType=17), AXE_TWO_HAND).two_hand is True
    assert 17 in TWO_HAND_INVENTORY_TYPES


def test_a_ranged_weapon_is_not_flagged():
    """R10: rule 6 asks whether an off-hand item may sit beside this one.

    A bow (InventoryType 15) takes the ranged slot, not the main hand, so it
    never refuses one. 26 is the trap the old set fell into -- in Classic it is
    Ranged Right, which is wands as much as guns and crossbows, so flagging it
    called 37 wands two-handed.
    """
    assert weapon_fields(row(InventoryType=15), GUN).two_hand is False
    assert weapon_fields(row(InventoryType=26), GUN).two_hand is False
    assert TWO_HAND_INVENTORY_TYPES.isdisjoint({15, 25, 26})


def test_a_non_weapon_reports_zeroes_and_no_flag():
    fields = weapon_fields(
        row(InventoryType=1, ItemDelay=0, MinDamage_0=0, MaxDamage_0=0), MACE_ONE_HAND
    )
    assert (fields.damage_min, fields.damage_max, fields.speed, fields.dps) == (0, 0, 0.0, 0.0)
    assert fields.two_hand is False


def test_a_weapon_with_no_delay_reports_no_dps_rather_than_dividing_by_zero():
    fields = weapon_fields(row(ItemDelay=0), MACE_ONE_HAND)
    assert fields.speed == 0.0
    assert fields.dps == 0.0


# --- Curve-derived damage (the 1.60 client / Forever beta) --------------------


def curve_row(**overrides: object) -> dict[str, str]:
    """A row with no literal MinDamage_0/MaxDamage_0 -- the 1.60 client's shape."""
    base = {
        "ID": "12784",
        "ItemLevel": "63",
        "OverallQualityID": "3",
        "InventoryType": "17",
        "ItemDelay": "3800",
        "DmgVariance": "0.5",
    }
    base.update({key: str(value) for key, value in overrides.items()})
    return base


def test_a_build_without_the_damage_columns_reports_zeroes_when_no_curves_are_passed():
    """The 1.60 client computes weapon damage from curves; absent is not malformed."""
    sparse = {"ID": "12640", "InventoryType": "13", "ItemDelay": "2600"}
    fields = weapon_fields(sparse, MACE_ONE_HAND)
    assert (fields.damage_min, fields.damage_max) == (0, 0)
    assert fields.speed == 2.6


def test_a_curve_only_row_resolves_damage_off_the_curve_tables():
    """Arcanite Reaper (12784) on build 1.60.1.69893: item level 63, rare,
    3.8 second two-hander, variance 0.5 -> 153 to 256 (see test_simdb_weapons.py,
    which proves the same numbers off the same curve tables)."""
    fields = weapon_fields(curve_row(), AXE_TWO_HAND, curves())
    assert (fields.damage_min, fields.damage_max) == (153, 256)
    assert fields.speed == 3.8
    assert fields.dps == pytest.approx((153 + 256) / 2 / 3.8, abs=0.01)


def test_an_unavailable_curve_set_leaves_a_curve_only_row_at_zero():
    """A build whose fetch could not supply every ItemDamage* table resolves
    nothing rather than guessing -- the row stays a documented zero, exactly
    like a row with no damage columns at all, and is not dropped (see
    `_has_gear_value`'s weapon exemption in gear.py)."""
    fields = weapon_fields(curve_row(), AXE_TWO_HAND, WeaponCurves())
    assert (fields.damage_min, fields.damage_max, fields.dps) == (0, 0, 0.0)
    assert fields.speed == 3.8


def test_no_curves_argument_at_all_leaves_a_curve_only_row_at_zero():
    fields = weapon_fields(curve_row(), AXE_TWO_HAND)
    assert (fields.damage_min, fields.damage_max, fields.dps) == (0, 0, 0.0)


def test_literal_columns_win_over_the_curve_when_both_are_available():
    """A row that states MinDamage_0/MaxDamage_0 outright is never sent
    through the curve, even when a full curve set is passed alongside it."""
    literal = curve_row(MinDamage_0="138", MaxDamage_0="256")
    fields = weapon_fields(literal, AXE_TWO_HAND, curves())
    assert (fields.damage_min, fields.damage_max) == (138, 256)


def test_a_schema_with_no_variance_column_resolves_nothing_from_the_curve():
    """An absent DmgVariance is a schema without curve support, not a
    truncated row -- the same "absent column is a schema" rule
    `_optional_int` documents for armour and stats."""
    without_variance = {key: value for key, value in curve_row().items() if key != "DmgVariance"}
    fields = weapon_fields(without_variance, AXE_TWO_HAND, curves())
    assert (fields.damage_min, fields.damage_max) == (0, 0)
