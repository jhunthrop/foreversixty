"""ItemSparse + Item -> items/<class-slug>.json, ItemSet -> sets.json."""

from __future__ import annotations

import logging
import re

from pipeline.icons import resolve_icon
from pipeline.models import ClassItems, GearItem, ItemSetBonus, ItemSetRecord
from pipeline.normalize.classes import slugify
from pipeline.proficiency import WEAPON, can_equip
from pipeline.spelltext import SpellText

logger = logging.getLogger(__name__)

MAX_PLAYER_LEVEL = 60
STAT_COLUMNS = range(10)
SET_ITEM_COLUMNS = range(17)

#: OverallQualityID values the planner keeps: uncommon, rare, epic, legendary.
#: Poor (0) and common (1) are vendor trash the planner never recommends, and
#: artifact (6) and heirloom (7) are not obtainable player gear in Era.
PLANNER_QUALITIES = frozenset({2, 3, 4, 5})

#: Display names the client uses for rows that are not shipping player gear:
#: Gamemaster/GM items, QA and test rows ("AHNQIRAJ TEST ITEM"), the
#: "Deprecated ..." leftovers, the "Monster - ..." display weapons NPCs hold,
#: and the "(OLD)" duplicates kept beside their replacements. Matched
#: case-insensitively; every alternative but "(old)" is whole-word, so real
#: items whose names merely contain the letters ("Testament of Hope" contains
#: "test", "Magma Forged Band" contains "gm") are kept.
JUNK_NAME_PATTERN = re.compile(
    r"\bgamemaster\b|\bgm\b|\btest\b|\bdeprecated\b|\bmonster\b|\(old\)",
    re.IGNORECASE,
)

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
    # The 1.60 client (Forever beta) is a modern client build: its ItemSparse rows
    # reference the full retail ITEM_MOD vocabulary even though Forever's Classic-style
    # planner has no use for combat ratings retail grew after Classic (haste, expertise,
    # armor penetration, mastery, versatility, avoidance, leech, speed, indestructible,
    # corruption resistance, socket bonuses and the rest). None of Classic Era's items
    # use them, so they are mapped to None here rather than guessed at with a stat key
    # the planner does not have.
    36: None,  # haste rating
    37: None,  # expertise rating
    44: None,  # armor penetration rating
    46: None,  # health regen
    47: None,  # spell penetration
    50: None,  # mastery rating
    83: None,
    84: None,
    85: None,
    86: None,
    87: None,
    88: None,
    89: None,
    90: None,
    91: None,
    92: None,
    96: None,
    98: None,
    103: None,
    112: None,
    113: None,
    114: None,
    115: None,
    117: None,
    119: None,
    121: None,
    124: None,
    127: None,
    128: None,
    131: None,
    132: None,
    135: None,
    136: None,
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


def _column(row: dict[str, str], column: str) -> str:
    """One column's value, or ItemDataError if the row does not supply one.

    csv.DictReader pads a short row with None rather than dropping the key, so
    a present key with no value means a truncated row, and a missing key means a
    header that does not carry the column at all (for instance a bonusStat
    column with no paired bonusAmount). Both are unreadable, and this module
    raises rather than guess at them. Every ItemSparse and Item column read on
    the build_class_items path goes through here or through _int, so a
    malformed row is always the ItemDataError the orchestrator catches, never a
    KeyError or TypeError that takes the whole run down with it.
    """
    value = row.get(column)
    if value is None:
        raise ItemDataError(
            f"item {row.get('ID', '?')} has no readable {column}; "
            f"the item row or its header is malformed"
        )
    return value


def _int(row: dict[str, str], column: str) -> int:
    """One column's value as an int, or ItemDataError if it is not readable as one."""
    value = _column(row, column)
    try:
        return int(value)
    except ValueError as error:
        raise ItemDataError(
            f"item {row.get('ID', '?')} has a non-numeric {column} {value!r}"
        ) from error


def _optional_int(row: dict[str, str], column: str) -> int | None:
    """One column's value as an int, or None if this build's schema has no such column.

    The 1.60 client (Forever beta) dropped ItemSparse's flat `Resistances_*` and
    `StatModifier_bonusAmount_*` columns: armour and stat amounts are now computed from
    curve tables (`RandPropPoints`, `ItemArmorTotal`, `ItemArmorQuality`) this pipeline
    does not resolve, so the client states nothing in a column for them. A column
    entirely absent from the row is that -- an older-schema build (Classic Era) still
    carries every one of these columns, so this only ever fires on a build that truly
    lacks the column. A present key with an empty value is still a truncated row, and
    _int still raises for that.
    """
    if column not in row:
        return None
    return _int(row, column)


def _stats(row: dict[str, str]) -> dict[str, int]:
    stats: dict[str, int] = {}
    for index in STAT_COLUMNS:
        stat_column = f"StatModifier_bonusStat_{index}"
        # The live table carries all ten stat column pairs, but they are
        # contiguous: a header that stops early simply has fewer to read.
        # Only an absent key means that; a key whose value is missing is a
        # truncated row, which _column turns into an error.
        if stat_column not in row:
            break
        stat_id = _int(row, stat_column)
        if stat_id < 0:
            continue
        if stat_id not in STAT_BY_MODIFIER_ID:
            raise ItemDataError(
                f"item {row['ID']} uses unknown stat modifier id {stat_id}; "
                f"add it to STAT_BY_MODIFIER_ID in pipeline/normalize/gear.py"
            )
        key = STAT_BY_MODIFIER_ID[stat_id]
        amount = _optional_int(row, f"StatModifier_bonusAmount_{index}") or 0
        if key is None or amount == 0:
            continue
        stats[key] = stats.get(key, 0) + amount
    for index, key in RESISTANCE_KEYS.items():
        amount = _optional_int(row, f"Resistances_{index}") or 0
        if amount:
            stats[key] = stats.get(key, 0) + amount
    return stats


def is_junk_name(name: str) -> bool:
    """True when the display name marks the row as non-shipping client data.

    See JUNK_NAME_PATTERN for what counts and why the alternatives are
    whole-word: a legitimate name that merely contains the letters must survive.
    """
    return JUNK_NAME_PATTERN.search(name) is not None


def _has_gear_value(armor: int, stats: dict[str, int], item_class_id: int) -> bool:
    """True when the item carries something the planner can compare.

    An item with no armour and no non-zero stat gives the planner nothing to
    reason about, so it is dropped -- for armour and for every other item class.

    Weapons (`Item.ClassID` 2) are exempt. A weapon's value is its damage, and
    this pipeline does not emit damage yet, so judging a weapon on armour and
    stats alone drops real gear: Annihilator, Arcanite Champion and the Hakkari
    warblades all carry nothing but their damage. A weapon therefore survives on
    the quality and junk-name clauses alone. Remove the exemption once weapon
    damage is emitted and a damage-less weapon really is valueless.
    """
    if item_class_id == WEAPON:
        return True
    return armor != 0 or any(stats.values())


def _icon_name(item_row: dict[str, str], icons: dict[int, str], display_name: str) -> str:
    """The item's icon name, falling back to the client's placeholder art.

    Some client rows carry `IconFileDataID` 0, meaning the client itself ships
    no art for the item; see `pipeline.icons.resolve_icon` for what happens then.
    """
    return resolve_icon(
        _int(item_row, "IconFileDataID"),
        icons,
        f"item {item_row.get('ID', '?')} ({display_name})",
    )


def build_class_items(
    sparse_rows: list[dict[str, str]],
    item_rows: list[dict[str, str]],
    class_rows: list[dict[str, str]],
    icons: dict[int, str],
    build: str,
) -> list[ClassItems]:
    """One equippable item list per class. Raises ItemDataError if a row is unreadable."""
    by_id = {_int(row, "ID"): row for row in item_rows}
    candidates: list[tuple[GearItem, int, int, int]] = []
    for row in sparse_rows:
        slot = SLOT_BY_INVENTORY_TYPE.get(_int(row, "InventoryType"))
        if slot is None:
            continue
        required_level = _int(row, "RequiredLevel")
        if required_level > MAX_PLAYER_LEVEL:
            continue
        if _int(row, "OverallQualityID") not in PLANNER_QUALITIES:
            continue
        display_name = _column(row, "Display_lang")
        if is_junk_name(display_name):
            continue
        item_id = _int(row, "ID")
        item_row = by_id.get(item_id)
        if item_row is None:
            logger.warning("item %s is in ItemSparse but not in Item; skipping it", item_id)
            continue
        item_class_id = _int(item_row, "ClassID")
        armor = _optional_int(row, "Resistances_0") or 0
        stats = _stats(row)
        if not _has_gear_value(armor, stats, item_class_id):
            continue
        item = GearItem(
            id=item_id,
            name=display_name,
            icon=_icon_name(item_row, icons, display_name),
            slot=slot,
            quality=_int(row, "OverallQualityID"),
            required_level=required_level,
            item_level=_int(row, "ItemLevel"),
            armor=armor,
            stats=stats,
            set_id=_int(row, "ItemSet") or None,
            unique=_int(row, "MaxCount") == 1,
        )
        candidates.append(
            (
                item,
                _int(row, "AllowableClass"),
                item_class_id,
                _int(item_row, "SubclassID"),
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
        set_id = int(row["ItemSetID"])
        spell_id = int(row["SpellID"])
        description = spell_text.describe(spell_id)
        if not description:
            logger.warning(
                "set %s has a %s-piece bonus whose spell %s has no description",
                set_id,
                row["Threshold"],
                spell_id,
            )
        bonuses.setdefault(set_id, []).append(
            ItemSetBonus(pieces=int(row["Threshold"]), description=description)
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
