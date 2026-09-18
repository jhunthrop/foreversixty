from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.simdb.weapons import WeaponCurves, load_weapon_curves, weapon_damage

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


def row(**overrides: object) -> dict[str, str]:
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


def test_a_two_hander_resolves_to_the_numbers_the_client_shows():
    """Arcanite Reaper (12784) on build 1.60.1.69893: item level 63, rare,
    3.8 second two-hander, variance 0.5 -> 153 to 256."""
    damage = weapon_damage(row(), AXE_TWO_HAND, curves())
    assert damage is not None
    assert (damage.minimum, damage.maximum, damage.speed) == (153.0, 256.0, 3.8)


def test_a_one_hander_uses_the_one_hand_curve():
    """Mace of Unending Life (21407): item level 70, epic, 2.6 second one-hander,
    variance 0.6 -> 93 to 175."""
    damage = weapon_damage(
        row(ID=21407, ItemLevel=70, OverallQualityID=4, InventoryType=21, ItemDelay=2600,
            DmgVariance="0.6"),
        MACE_ONE_HAND,
        curves(),
    )
    assert damage is not None
    assert (damage.minimum, damage.maximum, damage.speed) == (93.0, 175.0, 2.6)


def test_a_gun_uses_the_ranged_curve_not_the_one_hand_one():
    damage = weapon_damage(
        row(ID=217314, ItemLevel=27, OverallQualityID=2, InventoryType=26, ItemDelay=1700,
            DmgVariance="0.6"),
        GUN,
        curves(),
    )
    assert damage is not None
    assert (damage.minimum, damage.maximum, damage.speed) == (13.0, 24.0, 1.7)


def test_a_wand_uses_the_wand_curve():
    damage = weapon_damage(
        row(ID=217287, ItemLevel=35, OverallQualityID=2, InventoryType=26, ItemDelay=2000,
            DmgVariance="1.0"),
        WAND,
        curves(),
    )
    assert damage is not None
    assert (damage.minimum, damage.maximum, damage.speed) == (29.0, 88.0, 2.0)


def test_an_item_level_past_the_table_is_clamped_to_its_last_row():
    """The real curves stop at item level 100; a higher one reads the last row
    rather than falling off the end, the same clamp item_curves.py uses."""
    high = weapon_damage(row(ItemLevel=140), AXE_TWO_HAND, curves())
    last = weapon_damage(row(ItemLevel=100), AXE_TWO_HAND, curves())
    assert high == last


def test_a_row_with_no_delay_is_not_a_weapon():
    assert weapon_damage(row(ItemDelay=0), AXE_TWO_HAND, curves()) is None


def test_a_schema_with_no_delay_column_at_all_resolves_nothing():
    """An absent column is a schema, not a broken row: it resolves to None
    rather than the ItemDataError a present-but-empty column raises."""
    without = {key: value for key, value in row().items() if key != "ItemDelay"}
    assert weapon_damage(without, AXE_TWO_HAND, curves()) is None


def test_an_older_schema_row_uses_its_literal_columns():
    """Classic Era states MinDamage_0/MaxDamage_0 outright. The curve is not
    consulted, and it does not need to be available."""
    literal = row(MinDamage_0="138", MaxDamage_0="256")
    damage = weapon_damage(literal, AXE_TWO_HAND, WeaponCurves())
    assert damage is not None
    assert (damage.minimum, damage.maximum, damage.speed) == (138.0, 256.0, 3.8)


def test_a_curve_set_that_could_not_be_fetched_resolves_nothing():
    assert WeaponCurves().available is False
    assert weapon_damage(row(), AXE_TWO_HAND, WeaponCurves()) is None


def test_an_unmapped_subclass_falls_back_to_the_hand_curves():
    """Fishing poles (subclass 20) and any subclass the client adds are melee as
    far as the damage curve goes; only bows, guns, crossbows, wands and thrown
    weapons have their own table."""
    pole = weapon_damage(row(InventoryType=17), 20, curves())
    axe = weapon_damage(row(InventoryType=17), AXE_TWO_HAND, curves())
    assert pole == axe
