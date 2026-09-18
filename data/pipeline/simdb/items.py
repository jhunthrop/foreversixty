"""ItemSparse + Item -> the engine's SimItem rows.

The filter is the planner's -- mapped inventory type, uncommon and better,
required level at most 60, not a gamemaster/test/monster row, present in both
tables -- minus the planner's "has gear value" clause. The sim needs every item
a player might own, and a weapon whose whole value is its damage has gear value
the planner cannot see.

On build 1.60.1.69893 that is 4,986 items: 759 with weapon damage, 828 in a
set, 285 with an on-equip spell (18 of which grant a stat).

Armour and stats are not read from the columns here. The 1.60 client states
neither: both come off the curve tables, and `normalize/gear.py`'s
`resolve_item_values` already decides literal-versus-curve per row for the
planner. Calling it means the planner and the sim cannot disagree about what an
item is worth. Weapon damage comes from `simdb/weapons.py` the same way.
"""

from __future__ import annotations

from collections.abc import Mapping

from pipeline.normalize.gear import (
    MAX_PLAYER_LEVEL,
    PLANNER_QUALITIES,
    column_value,
    int_column,
    is_junk_name,
    resolve_item_values,
)
from pipeline.normalize.item_curves import ItemCurves
from pipeline.simdb.equip import SpellBonus
from pipeline.simdb.statmap import stat_array, weapon_skill_array
from pipeline.simdb.weapons import WeaponCurves, weapon_damage
from pipeline.simproto import pb

ITEM_CLASS_WEAPON = 2
ITEM_CLASS_ARMOR = 4
ARMOR_SUBCLASS_SHIELD = 6

#: AllowableClass when every class may equip the item. Every other value is a
#: bitmask over ChrClasses.ID - 1, including the negative ones the 1.60 client
#: uses (-1136 is priest, mage and warlock, not "everyone").
ALLOWABLE_ALL_CLASSES = -1

#: InventoryType -> the engine's ItemType. An inventory type that is absent is
#: not a slot the engine equips (shirts, tabards, bags, ammo, quivers, relics),
#: and is the same set `normalize/gear.py`'s SLOT_BY_INVENTORY_TYPE keeps.
ITEM_TYPE_BY_INVENTORY_TYPE: dict[int, str] = {
    1: "ItemTypeHead",
    2: "ItemTypeNeck",
    3: "ItemTypeShoulder",
    5: "ItemTypeChest",
    6: "ItemTypeWaist",
    7: "ItemTypeLegs",
    8: "ItemTypeFeet",
    9: "ItemTypeWrist",
    10: "ItemTypeHands",
    11: "ItemTypeFinger",
    12: "ItemTypeTrinket",
    13: "ItemTypeWeapon",
    14: "ItemTypeWeapon",
    15: "ItemTypeRanged",
    16: "ItemTypeBack",
    17: "ItemTypeWeapon",
    20: "ItemTypeChest",
    21: "ItemTypeWeapon",
    22: "ItemTypeWeapon",
    23: "ItemTypeWeapon",
    25: "ItemTypeRanged",
    26: "ItemTypeRanged",
}

#: InventoryType -> HandType, for weapons only.
HAND_TYPE_BY_INVENTORY_TYPE: dict[int, str] = {
    13: "HandTypeMainHand",
    14: "HandTypeOffHand",
    17: "HandTypeTwoHand",
    21: "HandTypeMainHand",
    22: "HandTypeOffHand",
    23: "HandTypeOffHand",
}

#: Item.SubclassID -> ArmorType, for Item.ClassID 4. See pipeline/proficiency.py
#: for the whole subclass vocabulary.
ARMOR_TYPE_BY_SUBCLASS: dict[int, str] = {
    1: "ArmorTypeCloth",
    2: "ArmorTypeLeather",
    3: "ArmorTypeMail",
    4: "ArmorTypePlate",
}

#: Item.SubclassID -> WeaponType, for Item.ClassID 2.
WEAPON_TYPE_BY_SUBCLASS: dict[int, str] = {
    0: "WeaponTypeAxe",
    1: "WeaponTypeAxe",
    4: "WeaponTypeMace",
    5: "WeaponTypeMace",
    6: "WeaponTypePolearm",
    7: "WeaponTypeSword",
    8: "WeaponTypeSword",
    10: "WeaponTypeStaff",
    13: "WeaponTypeFist",
    15: "WeaponTypeDagger",
}

#: Item.SubclassID -> RangedWeaponType, for Item.ClassID 2.
RANGED_TYPE_BY_SUBCLASS: dict[int, str] = {
    2: "RangedWeaponTypeBow",
    3: "RangedWeaponTypeGun",
    16: "RangedWeaponTypeThrown",
    18: "RangedWeaponTypeCrossbow",
    19: "RangedWeaponTypeWand",
}

#: ChrClasses.ID -> the engine's Class enum name. The two numbering schemes
#: disagree (the client's rogue is 4, the engine's is 6), so this is explicit.
PROTO_CLASS_BY_CHR_CLASS_ID: dict[int, str] = {
    1: "ClassWarrior",
    2: "ClassPaladin",
    3: "ClassHunter",
    4: "ClassRogue",
    5: "ClassPriest",
    7: "ClassShaman",
    8: "ClassMage",
    9: "ClassWarlock",
    11: "ClassDruid",
}


def simdb_item_rows(
    sparse_rows: list[dict[str, str]],
    item_rows: list[dict[str, str]],
) -> list[tuple[dict[str, str], dict[str, str]]]:
    """The (ItemSparse, Item) pairs the sim database keeps, sorted by item id."""
    by_id = {int_column(row, "ID"): row for row in item_rows}
    kept: list[tuple[dict[str, str], dict[str, str]]] = []
    for sparse in sparse_rows:
        if int_column(sparse, "InventoryType") not in ITEM_TYPE_BY_INVENTORY_TYPE:
            continue
        if int_column(sparse, "OverallQualityID") not in PLANNER_QUALITIES:
            continue
        if int_column(sparse, "RequiredLevel") > MAX_PLAYER_LEVEL:
            continue
        if is_junk_name(column_value(sparse, "Display_lang")):
            continue
        item = by_id.get(int_column(sparse, "ID"))
        if item is None:
            continue
        kept.append((sparse, item))
    return sorted(kept, key=lambda pair: int_column(pair[0], "ID"))


def _class_allowlist(allowable: int) -> list[int]:
    """The proto Class values a mask permits, sorted, empty when unrestricted."""
    if allowable == ALLOWABLE_ALL_CLASSES:
        return []
    return sorted(
        pb.Class.Value(name)
        for chr_class_id, name in PROTO_CLASS_BY_CHR_CLASS_ID.items()
        if allowable & (1 << (chr_class_id - 1))
    )


def build_sim_items(
    pairs: list[tuple[dict[str, str], dict[str, str]]],
    set_names: Mapping[int, str],
    equip: Mapping[int, SpellBonus],
    curves: ItemCurves,
    weapon_curves: WeaponCurves,
) -> list[pb.SimItem]:
    items: list[pb.SimItem] = []
    for sparse, item_row in pairs:
        item_id = int_column(sparse, "ID")
        inventory = int_column(sparse, "InventoryType")
        class_id = int_column(item_row, "ClassID")
        subclass = int_column(item_row, "SubclassID")

        armor, resolved = resolve_item_values(sparse, item_row, curves)
        stats: dict[str, float] = {key: float(value) for key, value in resolved.items()}
        if armor:
            stats["armor"] = stats.get("armor", 0.0) + armor
        bonus = equip.get(item_id, SpellBonus())
        for key, amount in bonus.stats.items():
            stats[key] = stats.get(key, 0.0) + amount

        item = pb.SimItem(
            id=item_id,
            name=column_value(sparse, "Display_lang"),
            type=pb.ItemType.Value(ITEM_TYPE_BY_INVENTORY_TYPE[inventory]),
            stats=stat_array(stats),
            weapon_skills=weapon_skill_array(bonus.weapon_skills),
            bonus_physical_damage=bonus.bonus_physical_damage,
            class_allowlist=_class_allowlist(int_column(sparse, "AllowableClass")),
        )
        if class_id == ITEM_CLASS_ARMOR:
            if subclass in ARMOR_TYPE_BY_SUBCLASS:
                item.armor_type = pb.ArmorType.Value(ARMOR_TYPE_BY_SUBCLASS[subclass])
            elif subclass == ARMOR_SUBCLASS_SHIELD:
                item.weapon_type = pb.WeaponType.Value("WeaponTypeShield")
        if class_id == ITEM_CLASS_WEAPON:
            if subclass in WEAPON_TYPE_BY_SUBCLASS:
                item.weapon_type = pb.WeaponType.Value(WEAPON_TYPE_BY_SUBCLASS[subclass])
                if inventory in HAND_TYPE_BY_INVENTORY_TYPE:
                    item.hand_type = pb.HandType.Value(HAND_TYPE_BY_INVENTORY_TYPE[inventory])
            if subclass in RANGED_TYPE_BY_SUBCLASS:
                item.ranged_weapon_type = pb.RangedWeaponType.Value(
                    RANGED_TYPE_BY_SUBCLASS[subclass]
                )
            damage = weapon_damage(sparse, subclass, weapon_curves)
            if damage is not None:
                item.weapon_damage_min = damage.minimum
                item.weapon_damage_max = damage.maximum
                item.weapon_speed = damage.speed

        set_id = int_column(sparse, "ItemSet")
        if set_id:
            item.set_id = set_id
            item.set_name = set_names.get(set_id, "")
        items.append(item)
    return items
