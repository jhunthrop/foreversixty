"""ItemSparse + Item -> items/<class-slug>.json, ItemSet -> sets.json."""

from __future__ import annotations

from pipeline.models import ClassItems, GearItem, ItemSetBonus, ItemSetRecord
from pipeline.normalize.classes import slugify
from pipeline.proficiency import can_equip
from pipeline.spelltext import SpellText

MAX_PLAYER_LEVEL = 60
STAT_COLUMNS = range(10)
SET_ITEM_COLUMNS = range(17)

#: InventoryType -> the planner's slot name. Anything absent is not gear the
#: planner tracks (shirts, tabards, bags, ammo, quivers, relics) and is dropped.
SLOT_BY_INVENTORY_TYPE: dict[int, str] = {
    1: "head",
    2: "neck",
    3: "shoulder",
    5: "chest",
    6: "waist",
    7: "legs",
    8: "feet",
    9: "wrist",
    10: "hands",
    11: "finger",
    12: "trinket",
    13: "main_hand",
    14: "off_hand",
    15: "ranged",
    16: "back",
    17: "main_hand",
    20: "chest",
    21: "main_hand",
    22: "off_hand",
    23: "off_hand",
    25: "ranged",
    26: "ranged",
}

#: StatModifier_bonusStat_<n> -> the planner's stat key, or None for a stat the
#: planner does not track. An id that is not a key here raises ItemDataError
#: rather than being dropped silently or guessed at. Classic Era only uses
#: 0, 1, 3, 4, 5, 6, 7, 38 and 51; the rest are here so a later build that
#: starts using them needs no code change.
STAT_BY_MODIFIER_ID: dict[int, str | None] = {
    0: None,  # mana
    1: None,  # health
    3: "agility",
    4: "strength",
    5: "intellect",
    6: "spirit",
    7: "stamina",
    12: "defense",
    13: "dodge",
    14: "parry",
    15: "block",
    31: "hit",
    32: "crit",
    38: "attack_power",
    39: None,  # ranged attack power
    41: "healing",
    42: "spell_power",
    43: "mp5",
    45: "spell_power",
    48: "block",
    51: "fire_res",
    52: "frost_res",
    53: None,  # holy resistance
    54: "shadow_res",
    55: "nature_res",
    56: "arcane_res",
}

#: Resistances_<n> -> stat key. Index 0 is armour and index 1 is holy
#: resistance, which the planner does not track.
RESISTANCE_KEYS: dict[int, str] = {
    2: "fire_res",
    3: "nature_res",
    4: "frost_res",
    5: "shadow_res",
    6: "arcane_res",
}


class ItemDataError(ValueError):
    """The item tables hold something this normalizer will not guess at."""


def _stats(row: dict[str, str]) -> dict[str, int]:
    stats: dict[str, int] = {}
    for index in STAT_COLUMNS:
        # The live table carries all ten stat column pairs, but they are
        # contiguous: a build that exports fewer simply has fewer to read.
        raw_stat_id = row.get(f"StatModifier_bonusStat_{index}")
        if raw_stat_id is None:
            break
        stat_id = int(raw_stat_id)
        if stat_id < 0:
            continue
        if stat_id not in STAT_BY_MODIFIER_ID:
            raise ItemDataError(
                f"item {row['ID']} uses unknown stat modifier id {stat_id}; "
                f"add it to STAT_BY_MODIFIER_ID in pipeline/normalize/gear.py"
            )
        key = STAT_BY_MODIFIER_ID[stat_id]
        amount = int(row[f"StatModifier_bonusAmount_{index}"])
        if key is None or amount == 0:
            continue
        stats[key] = stats.get(key, 0) + amount
    for index, key in RESISTANCE_KEYS.items():
        amount = int(row[f"Resistances_{index}"])
        if amount:
            stats[key] = stats.get(key, 0) + amount
    return stats


def build_class_items(
    sparse_rows: list[dict[str, str]],
    item_rows: list[dict[str, str]],
    class_rows: list[dict[str, str]],
    icons: dict[int, str],
    build: str,
) -> list[ClassItems]:
    """One equippable item list per class. Raises ItemDataError if a row is unreadable."""
    by_id = {int(row["ID"]): row for row in item_rows}
    candidates: list[tuple[GearItem, int, int, int]] = []
    for row in sparse_rows:
        inventory_type = int(row["InventoryType"])
        slot = SLOT_BY_INVENTORY_TYPE.get(inventory_type)
        if slot is None:
            continue
        if int(row["RequiredLevel"]) > MAX_PLAYER_LEVEL:
            continue
        item_id = int(row["ID"])
        item_row = by_id.get(item_id)
        if item_row is None:
            continue
        set_id = int(row["ItemSet"]) or None
        item = GearItem(
            id=item_id,
            name=row["Display_lang"],
            icon=icons.get(int(item_row["IconFileDataID"]), ""),
            slot=slot,
            quality=int(row["OverallQualityID"]),
            required_level=int(row["RequiredLevel"]),
            item_level=int(row["ItemLevel"]),
            armor=int(row["Resistances_0"]),
            stats=_stats(row),
            set_id=set_id,
            unique=int(row["MaxCount"]) == 1,
        )
        candidates.append(
            (
                item,
                int(row["AllowableClass"]),
                int(item_row["ClassID"]),
                int(item_row["SubclassID"]),
            )
        )
    records: list[ClassItems] = []
    for class_row in class_rows:
        class_id = int(class_row["ID"])
        class_mask = 1 << (class_id - 1)
        items = [
            item
            for item, allowable, item_class_id, subclass_id in candidates
            if (allowable < 0 or allowable & class_mask)
            and can_equip(class_id, item_class_id, subclass_id)
        ]
        records.append(
            ClassItems(
                build=build,
                class_slug=slugify(class_row["Name_lang"]),
                items=sorted(items, key=lambda i: (i.required_level, i.name, i.id)),
            )
        )
    return records


def build_item_sets(
    set_rows: list[dict[str, str]],
    set_spell_rows: list[dict[str, str]],
    spell_text: SpellText,
) -> list[ItemSetRecord]:
    bonuses: dict[int, list[ItemSetBonus]] = {}
    for row in set_spell_rows:
        bonuses.setdefault(int(row["ItemSetID"]), []).append(
            ItemSetBonus(
                pieces=int(row["Threshold"]),
                description=spell_text.describe(int(row["SpellID"])),
            )
        )
    records: list[ItemSetRecord] = []
    for row in set_rows:
        set_id = int(row["ID"])
        item_ids = [
            int(row[f"ItemID_{index}"])
            for index in SET_ITEM_COLUMNS
            if int(row[f"ItemID_{index}"]) != 0
        ]
        records.append(
            ItemSetRecord(
                id=set_id,
                name=row["Name_lang"],
                item_ids=sorted(item_ids),
                bonuses=sorted(bonuses.get(set_id, []), key=lambda b: b.pieces),
            )
        )
    return sorted(records, key=lambda r: r.id)
