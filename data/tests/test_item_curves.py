from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.normalize.item_curves import (
    CHEST_INVENTORY_TYPE,
    CLOTH,
    PLATE,
    ROBE_INVENTORY_TYPE,
    SHIELD,
    load_item_curves,
    resolve_armor,
    stat_budget,
)

HERE = Path(__file__).parent

#: Item.SubclassID for ClassID ARMOR's "miscellaneous" category (rings, necks,
#: trinkets) -- see pipeline/proficiency.py's docstring for the full vocabulary.
MISC = 0


def fixture_curves():
    return load_item_curves(
        read_csv(HERE / "fixtures/ItemArmorTotal.csv"),
        read_csv(HERE / "fixtures/ItemArmorQuality.csv"),
        read_csv(HERE / "fixtures/ItemArmorShield.csv"),
        read_csv(HERE / "fixtures/ArmorLocation.csv"),
        read_csv(HERE / "fixtures/RandPropPoints.csv"),
    )


def test_available_requires_every_table():
    assert fixture_curves().available is True
    assert load_item_curves([], [], [], [], []).available is False
    partial = load_item_curves(read_csv(HERE / "fixtures/ItemArmorTotal.csv"), [], [], [], [])
    assert partial.available is False


def test_resolve_armor_matches_the_literal_era_value_for_the_same_item():
    """Plate, item level 66, epic (quality 4), head (InventoryType 1) -- the same
    item level/quality/slot as Helm of Might (16866), whose Era build states 608
    armour literally. See data/README.md for how this was cross-checked broadly."""
    curves = fixture_curves()
    assert resolve_armor(curves, 66, 4, 1, PLATE) == 608


def test_resolve_armor_treats_a_robe_as_a_chest():
    """InventoryType 20 (robe) has no row of its own in ArmorLocation; the client
    scores it the same as InventoryType 5 (chest)."""
    curves = fixture_curves()
    assert resolve_armor(curves, 40, 3, ROBE_INVENTORY_TYPE, CLOTH) == resolve_armor(
        curves, 40, 3, CHEST_INVENTORY_TYPE, CLOTH
    )


def test_resolve_armor_uses_item_armor_shield_not_the_location_table():
    curves = fixture_curves()
    # InventoryType 14 (a shield's own slot) carries a 0 ArmorLocation modifier for
    # every armour type; the shield's own curve table is used instead of that path.
    assert resolve_armor(curves, 40, 3, 14, SHIELD) == 1078


def test_resolve_armor_is_zero_for_a_subclass_with_no_armour_curve():
    """Rings, necks, trinkets, librams, idols and totems (Item.SubclassID 0, 7, 8,
    9 under ClassID ARMOR) carry no armour curve."""
    curves = fixture_curves()
    assert resolve_armor(curves, 40, 3, 11, MISC) == 0


def test_resolve_armor_is_zero_when_a_curve_table_has_no_data_at_all():
    empty = load_item_curves([], [], [], [], [])
    assert resolve_armor(empty, 66, 4, 1, PLATE) == 0


def test_stat_budget_is_none_for_a_quality_the_planner_never_reaches():
    """RandPropPoints has no column for poor/common (quality 0/1); PLANNER_QUALITIES
    already excludes those rows before gear.py ever calls stat_budget, but the
    function itself refuses to guess rather than relying on that caller behaviour."""
    curves = fixture_curves()
    assert stat_budget(curves, 40, 0, 1) is None
    assert stat_budget(curves, 40, 1, 1) is None


def test_stat_budget_is_none_for_an_inventory_type_with_no_slot_group():
    curves = fixture_curves()
    assert stat_budget(curves, 40, 3, 999) is None


def test_stat_budget_clamps_to_the_curve_table_s_own_range():
    """A fixture only carries item levels 40 and 66; an item level outside that
    range clamps to the nearest edge rather than raising -- the same behaviour the
    real, gapless 1-to-N table gives for an item level within its own range."""
    curves = fixture_curves()
    assert stat_budget(curves, 1, 4, 1) == stat_budget(curves, 40, 4, 1)
    assert stat_budget(curves, 9999, 4, 1) == stat_budget(curves, 66, 4, 1)
