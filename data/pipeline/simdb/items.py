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

from collections import Counter
from collections.abc import Iterable, Mapping

from pipeline import classicdb_items as cdb
from pipeline import wowhead_items as wh
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
from pipeline.simdb.ratings import convert_rating_stats
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
#: not a slot the engine equips (shirts, tabards, bags, ammo, quivers), and is
#: the same set `normalize/gear.py`'s SLOT_BY_INVENTORY_TYPE keeps.
#:
#: simdb-supplement lane, 2026-09-30: 28 (INVTYPE_RELIC -- librams, idols and
#: totems) used to be absent here even though `SLOT_BY_INVENTORY_TYPE` grew it
#: on 2026-09-28 (the night-bis-sources lane, which fixed the PLANNER's own
#: relic gap): four real client relics (Idol of the Moon 23197, Totem of the
#: Storm 23199, Totem of Thunder 228176, Howling Idol 272427) were published
#: leveling-bis picks with no `simitems.json` row at all, so `simdb.Attach`'s
#: `UnequipUnknown` stripped every one of them the same way a classic-db-only
#: id used to. `SLOT_BY_INVENTORY_TYPE` files a relic under the ranged slot
#: (same as a bow or wand, `normalize/gear.py`'s own comment), so this does
#: too.
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
    28: "ItemTypeRanged",
}

#: InventoryType -> HandType, for weapons only.
HAND_TYPE_BY_INVENTORY_TYPE: dict[int, str] = {
    13: "HandTypeOneHand",
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

#: `items.json`'s `faction_restriction` column -> the proto enum value.
#: The column's vocabulary is `pipeline/forkdb.py`'s FACTION_RESTRICTIONS;
#: this is the other end of it, kept here because only simdb needs the
#: enum and only forkdb needs the slug.
FACTION_RESTRICTION_BY_SLUG: dict[str, str] = {
    "": "FACTION_RESTRICTION_UNSPECIFIED",
    "alliance_only": "FACTION_RESTRICTION_ALLIANCE_ONLY",
    "horde_only": "FACTION_RESTRICTION_HORDE_ONLY",
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


def _apply_class_subclass(item: pb.SimItem, class_id: int, subclass: int, inventory: int) -> None:
    """Armour type, weapon type, hand type and ranged weapon type from
    Item.ClassID/SubclassID and InventoryType.

    A client row and a wowhead supplement row resolve this identically once
    each source's own columns are read into the same class/subclass/
    inventory-type vocabulary (`wowhead_items.WowheadItem` already does that
    translation for the wowhead side), so both `build_sim_items` and
    `build_wowhead_sim_items` call this rather than each keeping its own copy
    of the branch.
    """
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


def build_sim_items(
    pairs: list[tuple[dict[str, str], dict[str, str]]],
    set_names: Mapping[int, str],
    equip: Mapping[int, SpellBonus],
    curves: ItemCurves,
    weapon_curves: WeaponCurves,
    rating_factors: Mapping[str, float],
    fork_columns: Mapping[int, tuple[list[int], str]],
) -> list[pb.SimItem]:
    """`rating_factors` is `ratings.load_rating_factors`'s output: level-60
    rating points per 1% for every `ItemSparse`-column stat the client states
    as a combat rating (hit, crit, dodge, parry, block, defense). It converts
    only `resolve_item_values`'s output -- an on-equip spell's stat (`equip`)
    already states a flat percentage; see `pipeline/simdb/ratings.py`.

    `fork_columns` is item id -> (random suffix options, faction restriction
    slug), read from `items.json`'s two fork-derived columns
    (`pipeline/simdb/__init__.py`'s `_fork_columns`) rather than from the
    fork database directly, so this module still needs no engine checkout.
    An item with neither is absent from the mapping or maps to `([], "")`.
    """
    items: list[pb.SimItem] = []
    for sparse, item_row in pairs:
        item_id = int_column(sparse, "ID")
        inventory = int_column(sparse, "InventoryType")
        class_id = int_column(item_row, "ClassID")
        subclass = int_column(item_row, "SubclassID")

        armor, resolved = resolve_item_values(sparse, item_row, curves)
        stat_pairs: list[tuple[str, float]] = list(
            convert_rating_stats(resolved, rating_factors).items()
        )
        if armor:
            stat_pairs.append(("armor", armor))
        bonus = equip.get(item_id, SpellBonus())
        stat_pairs.extend(bonus.stats.items())

        item = pb.SimItem(
            id=item_id,
            name=column_value(sparse, "Display_lang"),
            type=pb.ItemType.Value(ITEM_TYPE_BY_INVENTORY_TYPE[inventory]),
            stats=stat_array(stat_pairs),
            weapon_skills=weapon_skill_array(bonus.weapon_skills),
            bonus_physical_damage=bonus.bonus_physical_damage,
            class_allowlist=_class_allowlist(int_column(sparse, "AllowableClass")),
            unique=int_column(sparse, "MaxCount") == 1,
            required_level=int_column(sparse, "RequiredLevel"),
        )
        _apply_class_subclass(item, class_id, subclass, inventory)
        if class_id == ITEM_CLASS_WEAPON:
            damage = weapon_damage(sparse, subclass, weapon_curves)
            if damage is not None:
                item.weapon_damage_min = damage.minimum
                item.weapon_damage_max = damage.maximum
                item.weapon_speed = damage.speed

        set_id = int_column(sparse, "ItemSet")
        if set_id:
            item.set_id = set_id
            item.set_name = set_names.get(set_id, "")

        suffix_options, restriction = fork_columns.get(item_id, ([], ""))
        if suffix_options:
            item.random_suffix_options.extend(suffix_options)
        if restriction not in FACTION_RESTRICTION_BY_SLUG:
            raise SystemExit(
                f"item {item_id} has faction_restriction {restriction!r} in "
                f"items.json; pipeline/simdb/items.py knows "
                f"{sorted(FACTION_RESTRICTION_BY_SLUG)}"
            )
        item.faction_restriction = pb.SimItem.FactionRestriction.Value(
            FACTION_RESTRICTION_BY_SLUG[restriction]
        )
        items.append(item)
    return items


def build_wowhead_sim_items(
    items: Iterable[wh.WowheadItem],
    rating_factors: Mapping[str, float],
    set_names: Mapping[int, str] | None = None,
    untracked: Counter[str] | None = None,
) -> list[pb.SimItem]:
    """`SimItem` rows for wowhead's supplement -- the ids the client's own
    `ItemSparse` lacks (`wowhead_items.supplement`'s output is what a caller
    passes as `items` here).

    Armour type, weapon type, hand type and ranged weapon type resolve
    through `_apply_class_subclass`, the same helper `build_sim_items` calls,
    since `WowheadItem.class_id`/`subclass_id`/`inventory_type` are already
    the client's own vocabulary (`wowhead_items._client_subclass` does that
    translation on load). Weapon damage and speed come straight from
    wowhead's own `damage_min`/`damage_max`/`speed` rather than a curve --
    wowhead states them outright; `simdb/weapons.py`'s curves are only for a
    client row that has none.

    Stats go through `wowhead_items.planner_stats` (which drops and counts a
    wowhead key the planner does not track, the same as the planner's own
    output) and then `ratings.convert_rating_stats`, exactly as a client
    row's `resolve_item_values` output does: wowhead's crit/hit/defense/
    dodge/parry/block amounts are combat-rating points, calibrated to match
    the client's own resolved values one for one (see
    `docs/superpowers/specs/2026-09-27-wowhead-item-supplement-design.md`),
    so they need the same rating-to-percentage conversion before the engine
    reads them as a flat percentage.

    No on-equip effects: wowhead states them as prose, not data, so
    `random_suffix_options` and `faction_restriction` stay at their proto
    defaults (a supplement item is never fork-restricted or fork-suffixed --
    `pipeline/simdb/__init__.py`'s caller logs the count and `untracked`
    rather than raising, since dropping a stat the planner cannot place is
    expected here, not an error).
    """
    set_names = set_names or {}
    out: list[pb.SimItem] = []
    for w in items:
        stat_pairs: list[tuple[str, float]] = list(
            convert_rating_stats(wh.planner_stats(w, untracked), rating_factors).items()
        )
        if w.armor:
            stat_pairs.append(("armor", float(w.armor)))

        item = pb.SimItem(
            id=w.id,
            name=w.name,
            type=pb.ItemType.Value(ITEM_TYPE_BY_INVENTORY_TYPE[w.inventory_type]),
            stats=stat_array(stat_pairs),
            class_allowlist=_class_allowlist(w.class_mask) if w.class_mask is not None else [],
            unique=w.unique,
            required_level=w.required_level,
        )
        _apply_class_subclass(item, w.class_id, w.subclass_id, w.inventory_type)
        if w.class_id == ITEM_CLASS_WEAPON and w.speed > 0:
            item.weapon_damage_min = float(w.damage_min)
            item.weapon_damage_max = float(w.damage_max)
            item.weapon_speed = w.speed

        if w.set_id:
            item.set_id = w.set_id
            item.set_name = set_names.get(w.set_id, "")
        out.append(item)
    return out


def _classicdb_class_allowlist(item: cdb.ClassicDbItem) -> list[int]:
    """The proto `Class` values classic-db's own `AllowableClass` mask AND the
    client's proficiency table both permit -- `pipeline.classicdb_items.
    class_allowed`, the SAME test `pipeline.normalize.classicdb.
    merge_class_items` uses to decide which class's `items/<class-slug>.json`
    a supplement item joins, so this allowlist always agrees with where the
    planner actually placed the item. A raw mask alone (`_class_allowlist`,
    the client-row path above) is not enough here: classic-db's own
    `AllowableClass` is sometimes `-1` (unrestricted) even where real Classic
    proficiency -- plate, a caster's lack of a shield, ... -- still narrows
    who can equip it (`is_gm_class_mask`'s own doc has the exact rows this
    was measured against)."""
    return sorted(
        pb.Class.Value(name)
        for chr_class_id, name in PROTO_CLASS_BY_CHR_CLASS_ID.items()
        if cdb.class_allowed(item, chr_class_id)
    )


def build_classicdb_sim_items(
    items: Iterable[cdb.ClassicDbItem],
    spells: Mapping[int, cdb.ClassicDbSpell],
    rating_factors: Mapping[str, float],
    set_names: Mapping[int, str] | None = None,
    fork_columns: Mapping[int, tuple[list[int], str]] | None = None,
    untracked: Counter[str] | None = None,
) -> list[pb.SimItem]:
    """`SimItem` rows for classic-db's supplement -- the ids neither the
    client's own `ItemSparse` nor wowhead's Forever gear-planner scrape carry
    at all (`pipeline.classicdb_items`'s own doc: Hand of Justice, Blackhand's
    Breadth, Devilsaur Eye and the rest of cmangos/classic-db's 1.12
    `item_template`). Same shape as `build_wowhead_sim_items` immediately
    above, extended rather than duplicated, with two differences forced by
    what classic-db's own extract states:

    * `class_allowlist` comes from `_classicdb_class_allowlist` (per-class
      membership, mask AND proficiency) rather than a raw mask -- see that
      helper's own doc.
    * A stat off `raw_stats`/resistances (`pipeline.classicdb_items.
      planner_stats`) is a combat rating and needs `ratings.
      convert_rating_stats`, exactly like a client row's `resolve_item_values`
      output; a stat off an on-equip spell (`pipeline.classicdb_items.
      equip_stats`) is already the flat percentage Classic states directly
      (verified against Blackhand's Breadth/Eye of the Beast -- see
      `classicdb_items.effect_text`'s own doc) and bypasses that conversion,
      the same split `build_sim_items` applies between `resolved` and
      `bonus.stats` above.

    `fork_columns` is the SAME `items.json` columns the client path reads:
    a classic-db row merges into `items.json` before `python -m pipeline
    loot` runs (`pipeline.normalize.classicdb.merge_items`), so `loot` places
    suffixes/faction restrictions on it exactly as it does any other row --
    `None` (the default) leaves every classic-db `SimItem` unrestricted and
    unsuffixed, the state a caller with no fork columns at all (no engine
    checkout) leaves every OTHER item in too.

    An item whose `InventoryType` this module does not equip at all
    (`ITEM_TYPE_BY_INVENTORY_TYPE` -- relics, InventoryType 28, are the one
    gap: `normalize/gear.py`'s `SLOT_BY_INVENTORY_TYPE` covers them but this
    dict does not) is counted in `untracked` and dropped, never guessed at.
    """
    set_names = set_names or {}
    fork_columns = fork_columns or {}
    out: list[pb.SimItem] = []
    for row in items:
        inventory = row.inventory_type
        if inventory not in ITEM_TYPE_BY_INVENTORY_TYPE:
            if untracked is not None:
                untracked[f"inventory_type_{inventory}"] += 1
            continue

        stat_pairs: list[tuple[str, float]] = list(
            convert_rating_stats(cdb.planner_stats(row), rating_factors).items()
        )
        if row.armor and row.class_id == ITEM_CLASS_ARMOR:
            stat_pairs.append(("armor", float(row.armor)))
        stat_pairs.extend(cdb.equip_stats(row, spells).items())

        item = pb.SimItem(
            id=row.id,
            name=row.name,
            type=pb.ItemType.Value(ITEM_TYPE_BY_INVENTORY_TYPE[inventory]),
            stats=stat_array(stat_pairs),
            class_allowlist=_classicdb_class_allowlist(row),
            unique=row.unique,
            required_level=row.required_level,
        )
        _apply_class_subclass(item, row.class_id, row.subclass_id, inventory)
        if row.class_id == ITEM_CLASS_WEAPON and row.delay > 0:
            item.weapon_damage_min = float(row.damage_min)
            item.weapon_damage_max = float(row.damage_max)
            item.weapon_speed = row.delay / 1000

        if row.set_id:
            item.set_id = row.set_id
            item.set_name = set_names.get(row.set_id, "")

        suffix_options, restriction = fork_columns.get(row.id, ([], ""))
        if suffix_options:
            item.random_suffix_options.extend(suffix_options)
        if restriction not in FACTION_RESTRICTION_BY_SLUG:
            raise SystemExit(
                f"item {row.id} has faction_restriction {restriction!r} in "
                f"items.json; pipeline/simdb/items.py knows "
                f"{sorted(FACTION_RESTRICTION_BY_SLUG)}"
            )
        item.faction_restriction = pb.SimItem.FactionRestriction.Value(
            FACTION_RESTRICTION_BY_SLUG[restriction]
        )
        out.append(item)
    return out
