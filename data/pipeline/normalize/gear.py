"""ItemSparse + Item -> items/<class-slug>.json, ItemSet -> sets.json."""

from __future__ import annotations

import logging
import re
from typing import TYPE_CHECKING, NamedTuple

from pipeline.icons import resolve_icon
from pipeline.models import ClassItems, GearItem, ItemSetBonus, ItemSetRecord
from pipeline.normalize.classes import slugify
from pipeline.normalize.item_curves import ItemCurves, resolve_armor, stat_budget
from pipeline.proficiency import ARMOR, WEAPON, can_equip
from pipeline.spelltext import SpellText

if TYPE_CHECKING:
    # Importing pipeline.normalize.effects at module level would execute
    # pipeline/simdb/__init__.py (it imports pipeline.simdb.equip), whose own
    # line 45 imports names from this module -- a module-level import here
    # would re-enter gear.py while it is still initialising. This module
    # already has `from __future__ import annotations`, so the annotation
    # below stays a string and needs no runtime import.
    from pipeline.normalize.effects import EffectIndex

logger = logging.getLogger(__name__)

MAX_PLAYER_LEVEL = 60
STAT_COLUMNS = range(10)
SET_ITEM_COLUMNS = range(17)

#: Item level at which the sanity cap below applies -- the top of Classic-style
#: itemization, where real gear's own ceiling is well understood (the highest
#: real armour and stat values in build 1.60.1.69893 are 1,833 and 46; see
#: data/README.md). Distinct from MAX_PLAYER_LEVEL even though both are 60 in
#: Classic-style numbering: one is a player level, the other an item level.
SANITY_CHECK_ITEM_LEVEL = 60
#: A single stat or armour value past these on a SANITY_CHECK_ITEM_LEVEL item is
#: not real gear -- every QA/test row found in build 1.60.1.69893 (JUNK_NAME_PATTERN
#: is the primary defence) cleared both by 4x or more. This is a backstop for the
#: next one JUNK_NAME_PATTERN does not yet know to catch, not a tuned gameplay
#: number. Raised from 2000 to 2100 when re-emitting build 1.15.9.69722 (Task 6,
#: 2026-09-17): Era's own ItemSparse carries item 13375, "Crest of Retribution",
#: a real rare shield (RequiredLevel 55, quality 3) with 2057 armour -- absent
#: from 1.60.1.69893's own item set, so the original tuning pass never saw it.
MAX_LEVEL_60_STAT = 200
MAX_LEVEL_60_ARMOR = 2100

#: OverallQualityID values the planner keeps: uncommon, rare, epic, legendary.
#: Poor (0) and common (1) are vendor trash the planner never recommends, and
#: artifact (6) and heirloom (7) are not obtainable player gear in Era.
PLANNER_QUALITIES = frozenset({2, 3, 4, 5})

#: Display names the client uses for rows that are not shipping player gear:
#: Gamemaster/GM items, QA and test rows ("AHNQIRAJ TEST ITEM", "QATest +1000
#: Spell Dmg Ring", "Ring of Critical Testing"), the "Deprecated ..." and
#: "[DNT] ..." (do-not-translate/internal) leftovers, the "Monster - ..."
#: display weapons NPCs hold, "[PH] ..." placeholder rows (an entire unshipped
#: "Brilliant/Rising/Shining Dawn" tier set was found this way in build
#: 1.60.1.69893 -- see data/README.md), "UNUSED ..." rows, "(DND)" GM/event
#: props, and the "(OLD)"/"zzOLD..." duplicates kept beside their
#: replacements. Matched case-insensitively; every alternative but "(old)",
#: "[ph]", "[dnt]" and "zzold" is whole-word, so real items whose names merely
#: contain the letters ("Testament of Hope" contains "test", "Magma Forged
#: Band" contains "gm") are kept. "testing" and "qatest" are separate
#: alternatives from "test" because neither is a whole-word match for it
#: ("Testing"/"QATest" do not end where "test" does); "testing" does have a
#: known false positive on real quest items whose name uses "Testing" as
#: ordinary English ("Field Testing Kit") -- accepted rather than narrowed
#: further because every one of those is not equippable gear (InventoryType 0),
#: so build_class_items's own slot filter drops it before is_junk_name is ever
#: called; see test_the_testing_pattern_has_a_known_false_positive_on_real_
#: consumables.
JUNK_NAME_PATTERN = re.compile(
    r"\bgamemaster\b|\bgm\b|\btest\b|\btesting\b|\bqatest\b"
    r"|\bdeprecated\b|\bmonster\b|\bunused\b|\bplaceholder\b|\bdnd\b"
    r"|\(old\)|zzold|\[ph\]|\[dnt\]",
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

#: InventoryType values that occupy both hands: two-handed melee (17),
#: bows (15), guns (26) and crossbows (26 shares the ranged type), plus
#: fishing poles (23 is off-hand-only so it is not here). A ranged weapon
#: counts because rule 6's question is "may an off-hand item sit beside
#: this", and in Classic itemisation a bow and a shield do coexist -- but
#: the engine models a two-handed ranged weapon as occupying the ranged
#: slot alone, and the planner mirrors the engine.
#:
#: Not the same question as `pipeline.simdb.weapons.TWO_HAND_INVENTORY_TYPE`
#: (17 only): that constant picks which melee damage curve a weapon scores
#: on, and a ranged weapon is never on that path -- it is resolved by
#: SubclassID instead. This set answers validation rule 6 -- does this item
#: occupy both hands -- which a bow does answer yes to.
TWO_HAND_INVENTORY_TYPES = frozenset({15, 17, 25, 26})

#: Stat keys that are a combat-rating point count when ItemSparse states
#: them (`STAT_BY_MODIFIER_ID`'s 12/13/14/15/31/32/48) but a flat literal
#: percentage when an on-equip spell states them instead
#: (`pipeline.simdb.equip.STAT_AURAS`'s 47/49/51/52/54/55/138/552, plus the
#: MOD_SKILL/defense branch) -- see data/README.md, "Hit, crit, dodge, parry
#: and block as percentages". The two units cannot be summed into one
#: `stats` entry; `_merge_effect_stats` raises rather than do it.
RATING_FAMILY_STAT_KEYS = frozenset({"hit", "crit", "dodge", "parry", "block", "defense"})


class ItemDataError(ValueError):
    """The item tables hold something this normalizer will not guess at."""


def column_value(row: dict[str, str], column: str) -> str:
    """One column's value, or ItemDataError if the row does not supply one.

    csv.DictReader pads a short row with None rather than dropping the key, so
    a present key with no value means a truncated row, and a missing key means a
    header that does not carry the column at all (for instance a bonusStat
    column with no paired bonusAmount). Both are unreadable, and this module
    raises rather than guess at them. Every ItemSparse and Item column read on
    the build_class_items path goes through here or through int_column, so a
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


def int_column(row: dict[str, str], column: str) -> int:
    """One column's value as an int, or ItemDataError if it is not readable as one."""
    value = column_value(row, column)
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
    int_column still raises for that.
    """
    if column not in row:
        return None
    return int_column(row, column)


def _apply_stat(stats: dict[str, int], row: dict[str, str], index: int, amount: int) -> None:
    """Add one StatModifier_bonusStat_<index> pair's amount to `stats`, if any.

    Shared by the literal-amount path (_stats) and the curve-resolved path
    (_curve_stats): both already know the stat id column exists (their
    STAT_COLUMNS loop has not broken out yet) and differ only in how `amount`
    was computed -- read straight from a column, or from a curve formula.
    """
    stat_id = int_column(row, f"StatModifier_bonusStat_{index}")
    if stat_id < 0:
        return
    if stat_id not in STAT_BY_MODIFIER_ID:
        raise ItemDataError(
            f"item {row['ID']} uses unknown stat modifier id {stat_id}; "
            f"add it to STAT_BY_MODIFIER_ID in pipeline/normalize/gear.py"
        )
    key = STAT_BY_MODIFIER_ID[stat_id]
    if key is None or amount == 0:
        return
    stats[key] = stats.get(key, 0) + amount


def _stats(row: dict[str, str]) -> dict[str, int]:
    stats: dict[str, int] = {}
    for index in STAT_COLUMNS:
        stat_column = f"StatModifier_bonusStat_{index}"
        # The live table carries all ten stat column pairs, but they are
        # contiguous: a header that stops early simply has fewer to read.
        # Only an absent key means that; a key whose value is missing is a
        # truncated row, which column_value turns into an error.
        if stat_column not in row:
            break
        amount = _optional_int(row, f"StatModifier_bonusAmount_{index}") or 0
        _apply_stat(stats, row, index, amount)
    for index, key in RESISTANCE_KEYS.items():
        amount = _optional_int(row, f"Resistances_{index}") or 0
        if amount:
            stats[key] = stats.get(key, 0) + amount
    return stats


class WeaponFields(NamedTuple):
    """The gear tail's weapon numbers, computed once per row."""

    damage_min: int
    damage_max: int
    speed: float
    dps: float
    two_hand: bool


def weapon_fields(row: dict[str, str]) -> WeaponFields:
    """Weapon damage, speed and the two-handed flag for one ItemSparse row.

    Every number is optional: the 1.60 client computes weapon damage from
    curve tables this pipeline does not resolve, so a missing column is an
    honest zero rather than a malformed row. A present but non-numeric
    column is still an ItemDataError, through int_column.

    `pipeline/simdb/weapons.py` resolves the simulator's own weapon damage
    and speed off the client's ItemDamage* curve tables -- the same numbers
    this function leaves at zero when ItemSparse states no literal damage
    column. The two are deliberate siblings, not an oversight: wiring the
    curve resolver in here is not available, since `simdb/weapons.py`
    already imports `pipeline.normalize.gear` (this module), so importing it
    back would be a circular import, and its `WeaponCurves` are a simdb-stage
    input this normalize stage does not build. On build 1.60.1.69893 every
    weapon here therefore has damage_min = damage_max = dps = 0 and only
    speed and two_hand real -- that is the intended, documented behaviour,
    not a bug.
    """
    delay = _optional_int(row, "ItemDelay") or 0
    damage_min = _optional_int(row, "ItemDamageMin_0") or 0
    damage_max = _optional_int(row, "ItemDamageMax_0") or 0
    speed = round(delay / 1000, 2)
    # Never divide by a zero speed: an item with damage and no delay is a
    # thrown weapon or a malformed row, and either way it has no dps.
    dps = round((damage_min + damage_max) / 2 / speed, 2) if speed > 0 else 0.0
    two_hand = int_column(row, "InventoryType") in TWO_HAND_INVENTORY_TYPES
    return WeaponFields(damage_min, damage_max, speed, dps, two_hand)


def _curve_stats(
    row: dict[str, str], curves: ItemCurves, item_level: int, quality: int, inventory_type: int
) -> dict[str, int]:
    """The row's stats computed from the curve tables, for a client whose
    ItemSparse states a stat type per slot (StatModifier_bonusStat_<n>) but no
    amount: the amount is `budget * StatPercentEditor_<n> / 10000`, where
    `budget` is the item's RandPropPoints stat-point budget. See item_curves.py
    for the formula and its verification."""
    budget = stat_budget(curves, item_level, quality, inventory_type)
    if budget is None:
        return {}
    stats: dict[str, int] = {}
    for index in STAT_COLUMNS:
        stat_column = f"StatModifier_bonusStat_{index}"
        if stat_column not in row:
            break
        editor_column = f"StatPercentEditor_{index}"
        amount = (
            round(budget * int_column(row, editor_column) / 10000)
            if editor_column in row
            else 0
        )
        _apply_stat(stats, row, index, amount)
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


def _row_has_literal_amounts(row: dict[str, str]) -> bool:
    """True when this build's ItemSparse states flat Resistances_*/
    StatModifier_bonusAmount_* columns (Classic Era's shape), as opposed to the
    curve-only shape the 1.60 client (Forever beta) uses -- see
    pipeline/normalize/item_curves.py for the curve-resolved alternative. Both
    columns disappear together (see data/README.md), so checking one suffices.
    """
    return "Resistances_0" in row


def _check_level_60_sanity(
    item_id: int, display_name: str, item_level: int, armor: int, stats: dict[str, int]
) -> None:
    """Raise ItemDataError for an armour or stat value no real level-60 item has.

    A backstop behind JUNK_NAME_PATTERN: a QA/test row's absurd, hand-typed value
    (700 crit, 1000 spell power) is how three such rows were first noticed leaking
    through the curve resolver's stat path in build 1.60.1.69893 -- see
    data/README.md. This catches the next one by value rather than by name.
    """
    if item_level != SANITY_CHECK_ITEM_LEVEL:
        return
    if armor > MAX_LEVEL_60_ARMOR:
        raise ItemDataError(
            f"item {item_id} ({display_name}) has {armor} armour, over the "
            f"level-60 sanity cap of {MAX_LEVEL_60_ARMOR}; this is almost certainly "
            f"a QA/test row JUNK_NAME_PATTERN should catch, not real gear"
        )
    for key, amount in stats.items():
        if amount > MAX_LEVEL_60_STAT:
            raise ItemDataError(
                f"item {item_id} ({display_name}) has {amount} {key}, over the "
                f"level-60 sanity cap of {MAX_LEVEL_60_STAT}; this is almost "
                f"certainly a QA/test row JUNK_NAME_PATTERN should catch, not real gear"
            )


def _icon_name(item_row: dict[str, str], icons: dict[int, str], display_name: str) -> str:
    """The item's icon name, falling back to the client's placeholder art.

    Some client rows carry `IconFileDataID` 0, meaning the client itself ships
    no art for the item; see `pipeline.icons.resolve_icon` for what happens then.
    """
    return resolve_icon(
        int_column(item_row, "IconFileDataID"),
        icons,
        f"item {item_row.get('ID', '?')} ({display_name})",
    )


def resolve_item_values(
    sparse: dict[str, str],
    item_row: dict[str, str],
    curves: ItemCurves | None,
) -> tuple[int, dict[str, int]]:
    """One item's armour and stats, however this build states them.

    Classic Era states both in columns; the 1.60 client (Forever beta) states
    neither and computes them from the curve tables. That decision used to live
    inside build_class_items, which meant the simulator's SimItem rows would
    have had to make it a second time -- and a planner and a sim that disagree
    about what an item is worth is the one bug neither would show. One function
    now answers it for both. See item_curves.py for the formulas.
    """
    inventory_type = int_column(sparse, "InventoryType")
    quality = int_column(sparse, "OverallQualityID")
    item_level = int_column(sparse, "ItemLevel")
    item_class_id = int_column(item_row, "ClassID")
    subclass_id = int_column(item_row, "SubclassID")
    if _row_has_literal_amounts(sparse):
        return (_optional_int(sparse, "Resistances_0") or 0), _stats(sparse)
    if curves is not None and curves.available:
        armor = (
            resolve_armor(curves, item_level, quality, inventory_type, subclass_id)
            if item_class_id == ARMOR
            else 0
        )
        return armor, _curve_stats(sparse, curves, item_level, quality, inventory_type)
    return 0, {}


def _merge_effect_stats(
    stats: dict[str, int], item_id: int, display_name: str, effects: EffectIndex
) -> None:
    """Fold an item's on-equip spell stats into its ItemSparse-sourced ones.

    Almost always safe to just add: the two sources agree on units for
    every stat except the rating family (hit, crit, dodge, parry, block,
    defense). There, ItemSparse's own `StatModifier_bonusStat` columns state
    a combat-rating point count -- `pipeline/simdb/ratings.py` converts it
    for the simulator, but `items/<class-slug>.json` itself keeps the raw
    rating, matching what the client's own tooltip shows -- while an
    on-equip spell's flat stat (`pipeline.simdb.equip.STAT_AURAS`) already
    states a literal percentage, the older Classic itemisation convention
    (see data/README.md, "Hit, crit, dodge, parry and block as
    percentages"). Summing a rating into a percentage would silently
    produce a number that is neither. No item on build 1.60.1.69893 mixes
    the two -- see test_an_equip_percentage_never_meets_an_itemsparse_rating
    -- so this raises rather than guess which side is right the first time
    one does.
    """
    for key, amount in effects.stats(item_id).items():
        if key in RATING_FAMILY_STAT_KEYS and stats.get(key):
            raise ItemDataError(
                f"item {item_id} ({display_name}) has {stats[key]} {key} from ItemSparse's "
                f"own rating columns and {amount} more {key} from an on-equip spell; these "
                f"are different units (data/README.md, 'Hit, crit, dodge, parry and block "
                f"as percentages') and pipeline/normalize/gear.py will not sum them blindly"
            )
        stats[key] = stats.get(key, 0) + amount


def build_class_items(
    sparse_rows: list[dict[str, str]],
    item_rows: list[dict[str, str]],
    class_rows: list[dict[str, str]],
    icons: dict[int, str],
    build: str,
    curves: ItemCurves | None = None,
    effects: EffectIndex | None = None,
) -> list[ClassItems]:
    """One equippable item list per class. Raises ItemDataError if a row is unreadable.

    `curves` resolves armour and stats for a build whose ItemSparse carries no
    literal amounts (the 1.60 client / Forever beta); pass None (the default)
    or an `ItemCurves` whose own tables are incomplete and such a row simply
    gets no armour and no stats, exactly as before curve support existed.

    `effects` folds an item's on-equip spell stats (`pipeline.normalize.
    effects.EffectIndex`) into the same `stats` dict ItemSparse's own columns
    populate, and sets `effect_text` from its use/proc spells. Pass None (the
    default) for a caller that has not built one and every item gets no
    extra stats and an empty `effect_text`, exactly as before this existed.
    """
    by_id = {int_column(row, "ID"): row for row in item_rows}
    candidates: list[tuple[GearItem, int, int, int]] = []
    for row in sparse_rows:
        inventory_type = int_column(row, "InventoryType")
        slot = SLOT_BY_INVENTORY_TYPE.get(inventory_type)
        if slot is None:
            continue
        required_level = int_column(row, "RequiredLevel")
        if required_level > MAX_PLAYER_LEVEL:
            continue
        quality = int_column(row, "OverallQualityID")
        if quality not in PLANNER_QUALITIES:
            continue
        display_name = column_value(row, "Display_lang")
        if is_junk_name(display_name):
            continue
        item_id = int_column(row, "ID")
        item_row = by_id.get(item_id)
        if item_row is None:
            logger.warning("item %s is in ItemSparse but not in Item; skipping it", item_id)
            continue
        item_class_id = int_column(item_row, "ClassID")
        subclass_id = int_column(item_row, "SubclassID")
        item_level = int_column(row, "ItemLevel")
        armor, stats = resolve_item_values(row, item_row, curves)
        if effects is not None:
            _merge_effect_stats(stats, item_id, display_name, effects)
        _check_level_60_sanity(item_id, display_name, item_level, armor, stats)
        if not _has_gear_value(armor, stats, item_class_id):
            continue
        weapon = weapon_fields(row)
        item = GearItem(
            id=item_id,
            name=display_name,
            icon=_icon_name(item_row, icons, display_name),
            slot=slot,
            quality=quality,
            required_level=required_level,
            item_level=item_level,
            armor=armor,
            stats=stats,
            damage_min=weapon.damage_min,
            damage_max=weapon.damage_max,
            speed=weapon.speed,
            dps=weapon.dps,
            two_hand=weapon.two_hand,
            effect_text="" if effects is None else effects.text(item_id),
            set_id=int_column(row, "ItemSet") or None,
            unique=int_column(row, "MaxCount") == 1,
        )
        candidates.append(
            (
                item,
                int_column(row, "AllowableClass"),
                item_class_id,
                subclass_id,
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
