from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.forkdb import FACTION_RESTRICTIONS
from pipeline.normalize.item_curves import load_item_curves
from pipeline.simdb.equip import SpellBonus
from pipeline.simdb.items import (
    FACTION_RESTRICTION_BY_SLUG,
    HAND_TYPE_BY_INVENTORY_TYPE,
    build_sim_items,
    simdb_item_rows,
)
from pipeline.simdb.weapons import WeaponCurves, load_weapon_curves
from pipeline.simproto import pb

HERE = Path(__file__).parent
FIXTURES = HERE / "fixtures"
SIM = FIXTURES / "sim"

#: tests/fixtures/ItemSparse_1_60.csv's row for MaxCount == 1 -- Novice's
#: Cloth Robe's columns (30001) reused under an id (14152) already present
#: in tests/fixtures/Item.csv, since simdb_item_rows keeps only ids in both.
UNIQUE_FIXTURE_ID = 14152

#: The committed build's own level-60 combatratings.txt row (see
#: tests/fixtures/sim/combatratings.txt), so a fixture test exercises the
#: same conversion the real build does.
RATING_FACTORS = {
    "hit": 10.0,
    "crit": 14.0,
    "dodge": 12.0,
    "parry": 15.0,
    "block": 5.0,
    "defense": 1.0,
}


def curves():
    return load_item_curves(
        read_csv(FIXTURES / "ItemArmorTotal.csv"),
        read_csv(FIXTURES / "ItemArmorQuality.csv"),
        read_csv(FIXTURES / "ItemArmorShield.csv"),
        read_csv(FIXTURES / "ArmorLocation.csv"),
        read_csv(FIXTURES / "RandPropPoints.csv"),
    )


def weapon_curves():
    return load_weapon_curves(
        read_csv(SIM / "ItemDamageOneHand.csv"),
        read_csv(SIM / "ItemDamageTwoHand.csv"),
        read_csv(SIM / "ItemDamageRanged.csv"),
        read_csv(SIM / "ItemDamageWand.csv"),
        read_csv(SIM / "ItemDamageThrown.csv"),
    )


def pairs():
    return simdb_item_rows(
        read_csv(FIXTURES / "ItemSparse_1_60.csv"),
        read_csv(FIXTURES / "Item.csv"),
    )


def built(set_names=None, equip=None, fork_columns=None):
    items = build_sim_items(
        pairs(),
        set_names or {},
        equip or {},
        curves(),
        weapon_curves(),
        RATING_FACTORS,
        fork_columns or {},
    )
    return {item.id: item for item in items}


def built_item(item_id, **kwargs):
    return built(**kwargs)[item_id]


def test_rows_are_sorted_by_item_id():
    ids = [int(sparse["ID"]) for sparse, _ in pairs()]
    assert ids == sorted(ids)


def test_the_filter_keeps_every_equippable_row_the_planner_would_drop():
    """The planner drops an item with no armour and no stats. The sim keeps it:
    Annihilator's whole value is its damage, and so is a plain white weapon's."""
    assert {int(sparse["ID"]) for sparse, _ in pairs()} == {
        12798,
        14152,
        16866,
        30001,
        30002,
        30003,
        30005,
        30006,
    }


def test_an_armour_piece_carries_its_curve_resolved_armour_and_stats():
    """Helm of Might (16866): plate, item level 66, epic, head. The planner
    emits 608 armour and 35 stamina / 15 strength for the same row."""
    helm = built()[16866]
    assert helm.type == pb.ItemType.Value("ItemTypeHead")
    assert helm.armor_type == pb.ArmorType.Value("ArmorTypePlate")
    assert helm.stats[pb.Stat.Value("StatArmor")] == 608.0
    assert helm.stats[pb.Stat.Value("StatStamina")] == 35.0
    assert helm.stats[pb.Stat.Value("StatStrength")] == 15.0


def test_a_column_stat_that_is_a_combat_rating_is_converted_to_a_percentage():
    """Loop of Minor Fortitude (30006) carries 5 crit through the curve, the
    same rating amount tests/test_normalize_gear.py checks the planner still
    shows raw (`items[30006].stats == {"strength": 8, "crit": 5}`) -- the
    planner keeps the rating number the client tooltip shows. simdb's own
    stats are what the engine reads as a flat percentage, so 5 rating points
    at this fixture's level-60 crit factor (14, see RATING_FACTORS above)
    convert to 5 / 14 %."""
    loop = built()[30006]
    assert loop.stats[pb.Stat.Value("StatCrit")] == pytest.approx(5.0 / 14.0)


def test_a_weapon_carries_its_damage_and_speed():
    """Annihilator (12798): item level 63, rare, 2.4 second one-hand axe,
    variance 0.6 -> 69 to 129."""
    axe = built()[12798]
    assert axe.type == pb.ItemType.Value("ItemTypeWeapon")
    assert axe.weapon_type == pb.WeaponType.Value("WeaponTypeAxe")
    assert axe.hand_type == pb.HandType.Value("HandTypeMainHand")
    assert (axe.weapon_damage_min, axe.weapon_damage_max, axe.weapon_speed) == (69.0, 129.0, 2.4)


def test_an_armour_piece_has_no_weapon_damage():
    helm = built()[16866]
    assert (helm.weapon_damage_min, helm.weapon_damage_max, helm.weapon_speed) == (0.0, 0.0, 0.0)


def test_a_shield_is_a_weapon_type_not_an_armour_type():
    shield = built()[30005]
    assert shield.weapon_type == pb.WeaponType.Value("WeaponTypeShield")
    assert shield.armor_type == pb.ArmorType.Value("ArmorTypeUnknown")


def test_a_set_piece_carries_its_set_id_and_name():
    helm = built(set_names={209: "Battlegear of Might"})[16866]
    assert helm.set_id == 209
    assert helm.set_name == "Battlegear of Might"


def test_an_item_in_no_set_carries_neither():
    axe = built()[12798]
    assert axe.set_id == 0
    assert axe.set_name == ""


def test_an_on_equip_bonus_is_folded_into_the_stat_array():
    bonus = SpellBonus(
        stats={"attack_power": 62.0},
        weapon_skills={"WeaponSkillAxes": 3.0},
        bonus_physical_damage=2.0,
    )
    axe = built(equip={12798: bonus})[12798]
    assert axe.stats[pb.Stat.Value("StatAttackPower")] == 62.0
    assert axe.weapon_skills[pb.WeaponSkill.Value("WeaponSkillAxes")] == 3.0
    assert axe.bonus_physical_damage == 2.0


def test_an_unrestricted_item_has_an_empty_class_allowlist():
    assert list(built()[12798].class_allowlist) == []


def test_a_class_restricted_item_lists_the_proto_classes():
    helm = built()[16866]  # AllowableClass 1: warrior only
    assert list(helm.class_allowlist) == [pb.Class.Value("ClassWarrior")]


def test_a_negative_mask_that_is_not_minus_one_excludes_rather_than_permits():
    """Build 1.60.1.69893 uses masks like -1136, which is 'these classes' with
    the sign bit set, not 'everyone'. Reading only `> 0` as restricted would
    hand a priest-and-mage-and-warlock item to a warrior."""
    pair = pairs()[0]
    sparse = dict(pair[0])
    sparse["AllowableClass"] = "-1136"
    item = build_sim_items(
        [(sparse, pair[1])], {}, {}, curves(), weapon_curves(), RATING_FACTORS, {}
    )[0]
    assert list(item.class_allowlist) == [
        pb.Class.Value("ClassMage"),
        pb.Class.Value("ClassPriest"),
        pb.Class.Value("ClassWarlock"),
    ]


def test_inventory_type_13_is_one_hand_not_main_hand():
    """InventoryType 13 covers weapons a player may equip in either hand
    (e.g. daggers, one-hand swords/maces/axes). HandTypeMainHand (21) is a
    disjoint InventoryType reserved for main-hand-only weapons. Collapsing 13
    into HandTypeMainHand left the engine with zero HandTypeOneHand rows,
    which made every InventoryType-13 weapon un-offhandable."""
    assert HAND_TYPE_BY_INVENTORY_TYPE[13] == "HandTypeOneHand"


def test_a_build_with_no_weapon_curves_still_emits_its_items():
    items = build_sim_items(pairs(), {}, {}, curves(), WeaponCurves(), RATING_FACTORS, {})
    axe = {item.id: item for item in items}[12798]
    assert axe.weapon_speed == 0.0
    assert axe.type == pb.ItemType.Value("ItemTypeWeapon")


def test_unique_and_required_level_come_from_item_sparse():
    item = built_item(16866)  # Helm of Might: MaxCount 0, RequiredLevel 60
    assert item.unique is False
    assert item.required_level == 60


def test_max_count_one_is_unique():
    assert built_item(UNIQUE_FIXTURE_ID).unique is True


def test_the_fork_columns_are_carried_through_from_items_json():
    item = built_item(12798, fork_columns={12798: ([5, 6], "horde_only")})
    assert list(item.random_suffix_options) == [5, 6]
    assert item.faction_restriction == pb.SimItem.FactionRestriction.Value(
        "FACTION_RESTRICTION_HORDE_ONLY"
    )


def test_an_item_with_no_fork_column_is_left_unrestricted():
    item = built_item(12798, fork_columns={})
    assert list(item.random_suffix_options) == []
    assert item.faction_restriction == pb.SimItem.FactionRestriction.Value(
        "FACTION_RESTRICTION_UNSPECIFIED"
    )


def test_an_unknown_faction_restriction_slug_raises():
    """Every other operator-facing failure on the simdb path -- a missing raw
    directory, a build that skipped `loot` -- raises SystemExit directly
    rather than being caught and converted somewhere upstream, since
    `python -m pipeline simdb` does not catch anything from this path. An
    unknown slug is the same kind of failure, so it raises the same way."""
    with pytest.raises(SystemExit, match="bogus_slug"):
        built_item(12798, fork_columns={12798: ([], "bogus_slug")})


def test_the_faction_vocabulary_is_the_same_on_both_ends():
    """`pipeline/forkdb.py`'s FACTION_RESTRICTIONS is the write side (the
    fork's faction enum -> the slug `loot` writes into items.json);
    FACTION_RESTRICTION_BY_SLUG here is the read side. A lane adding a third
    restriction to one without the other should get a red test here, not a
    SystemExit the first time someone builds a simdb with that item."""
    assert set(FACTION_RESTRICTIONS.values()) | {""} == set(FACTION_RESTRICTION_BY_SLUG)
