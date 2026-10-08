"""ItemSparse + Item -> items/<class-slug>.json, ItemSet -> sets.json."""

from __future__ import annotations

import logging
import re
from collections import Counter
from typing import TYPE_CHECKING, NamedTuple

from pipeline.icons import resolve_icon, resolve_icon_name
from pipeline.models import ClassItems, GearItem, ItemSetBonus, ItemSetRecord
from pipeline.normalize.classes import slugify
from pipeline.normalize.item_curves import ItemCurves, resolve_armor, stat_budget
from pipeline.normalize.weapon_curves import (
    WeaponCurves,
    WeaponDamage,
    resolve_weapon_damage,
    row_has_literal_weapon_damage,
)
from pipeline.proficiency import ARMOR, RELIC_SUBCLASSES, WEAPON, can_equip
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
#: classic-db follow-up, 2026-09-30: "Fishing Pole (JEFFTEST)" slipped past
#: every alternative above -- \btest\b needs a word boundary on both sides,
#: and "JEFFTEST" has none before "test". The two alternatives below catch a
#: token ENDING in "test" that is either parenthesised ("(JEFFTEST)") or a
#: whole ALL-CAPS word on its own ("JEFFTEST" with no parens), scoped
#: case-sensitive with (?-i:...) so they only fire on shouting-case tokens --
#: not on real, ordinarily-cased names that happen to end in "test", such as
#: "Contest Winner's Tabard" or "Rexxar's Testament" (see
#: test_the_junk_name_matcher_keeps_real_items_that_merely_contain_the_letters).
JUNK_NAME_PATTERN = re.compile(
    r"\bgamemaster\b|\bgm\b|\btest\b|\btesting\b|\bqatest\b"
    r"|\bdeprecated\b|\bmonster\b|\bunused\b|\bplaceholder\b|\bdnd\b"
    r"|\(old\)|zzold|\[ph\]|\[dnt\]"
    # catalogue-universe follow-up, 2026-09-30: classic-db's own QA rows
    # 5031-5039 are named "ZZZZZ", "ZZZZZZZZ", "ZZZZZ sword N".
    r"|\bz{4,}"
    # catalogue-universe follow-up, 2026-09-30: an ALL-CAPS token ending in
    # "test", parenthesised or standalone (e.g. "(JEFFTEST)", "JEFFTEST").
    r"|(?-i:\(\w*TEST\)|\b[A-Z]+TEST\b)",
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
    # INVTYPE_RELIC: librams (paladin), idols (druid) and totems (shaman)
    # occupy the ranged slot in Classic Era, same as a bow or wand. This
    # was missing entirely until the 2026-09-28 night-bis-sources lane,
    # which silently dropped every relic (65 on build 1.60.1.70009,
    # confirmed present with real names in the unfiltered items.json) out
    # of every class's items/<class>.json -- emptying paladin's ranged
    # slot at every leveling band, not a true absence of librams in the
    # client.
    28: "ranged",
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
    48: "block_value",
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
    # 101, 125, 126, 129, 134, 139 below all surfaced only once the hotfix
    # cache is merged (normalize-levels lane, 2026-09-29): each is a real,
    # defined ItemModType a hotfix-only item's StatModifier slot names --
    # verified via wowhead's Forever tooltip corroboration for the item that
    # carries it, never guessed at -- but every one is either a
    # creature-type-conditional damage bonus ("+N Attack Power against
    # <type>", rendered through wowhead's own `rtg<id>` template tag: 125
    # Humanoids on item 7683 "Bloody Brass Knuckles", 126 Elementals on item
    # 5444 "Miner's Cape", 129 Dragonkin on item 282778 "Mark of the Red
    # Flight") or a weapon-skill-style bonus (101, "Increased Polearms +1" on
    # item 13056 "Frenzied Striker"; 134, no rendered line at all on item
    # 2169 "Buzzer Blade" -- see that item's separate, unrelated "Equip:
    # spell power" line, which comes from an on-equip spell, not this
    # StatModifier slot; 139, "Increases damage done to Beasts..." on item
    # 4771 "Harvest Cloak") -- none of which this Classic-style planner's
    # stat vocabulary has a key for.
    101: None,
    103: None,
    112: None,
    113: None,
    114: None,
    115: None,
    117: None,
    119: None,
    121: None,
    124: None,
    125: None,
    126: None,
    127: None,
    128: None,
    129: None,
    131: None,
    132: None,
    134: None,
    135: None,
    136: None,
    139: None,
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

#: InventoryType values whose item takes the MAIN HAND and, in doing so,
#: leaves no off-hand free. That is two-handed melee (17) and nothing else.
#:
#: The ranged inventory types are deliberately absent. Rule 6 asks whether an
#: off-hand item may sit beside this one, and a bow (15), a thrown weapon (25)
#: or a Ranged Right item (26) occupies the ranged slot, not the main hand, so
#: it never refuses one -- a hunter carries a bow, a sword and a shield at
#: once. 26 is the trap: in Classic it is wands as much as guns and crossbows,
#: so including it flagged 169 items, 37 of them wands, as occupying both
#: hands.
#:
#: `pipeline.normalize.weapon_curves.TWO_HAND_INVENTORY_TYPE` is 17 for its own
#: reason -- it picks which melee damage curve a weapon scores on, and a
#: ranged weapon is resolved by SubclassID before that branch is reached. The
#: two constants
#: answer different questions and now happen to agree on the answer.
TWO_HAND_INVENTORY_TYPES = frozenset({17})

#: Item.ClassID 2 (WEAPON) rows whose InventoryType is an actual weapon slot
#: -- verified by joining build 1.60.1.70009's own Item.csv and
#: ItemSparse.csv on id: one-hand (13), ranged/bow (15), two-hand (17),
#: main-hand (21), off-hand weapon (22), thrown (25), ranged-right / wand,
#: gun, crossbow (26). Of the 1,775 class-2 rows in that build, all but
#: three sit at one of these seven values; the three (a junk-named "OLD..."
#: item, a "...Test" item, and one real weapon misfiled at 23 HOLDABLE) are
#: not worth widening the set for.
#:
#: 23 (HOLDABLE) is deliberately absent even though it is also an off-hand
#: slot: it is where a non-weapon, item-class-4 held item lives (a tome,
#: idol, relic or -- the defect this set fixes -- Father Flame, item 13371).
#: InventoryType alone cannot tell a weapon from a non-weapon here: the same
#: build has non-weapon (class != 2) rows at InventoryType 15, 17, 21, 22
#: and 26 too, so `is_weapon_row` below checks Item.ClassID first and
#: InventoryType second, never InventoryType alone.
WEAPON_INVENTORY_TYPES = frozenset({13, 15, 17, 21, 22, 25, 26})


#: Item.ClassID 2 (WEAPON) SubclassID -> the planner's `weapon_type` vocabulary
#: (classicdb-fidelity lane, 2026-09-30). Melee subclasses collapse their
#: one-hand/two-hand pair into one type -- `two_hand` (TWO_HAND_INVENTORY_TYPES)
#: already carries that distinction, so a caller does not need it twice.
#: Ranged subclasses stay one-to-one: `sim/cmd/leveling-bis`'s
#: `weapon_requirements.go` (`bis-ranker-integrity-6` lane) gates a caster's
#: wand slot and a hunter's/warrior's/rogue's bow-or-gun-or-crossbow-or-thrown
#: slot by this exact vocabulary. Values and ids are `pipeline.proficiency`'s
#: own WEAPON_SUBCLASSES table (its module doc has the citation for each);
#: 14 (Miscellaneous, fishing poles among other unshipped rows) and 20
#: (Fishing Pole) are deliberately absent, same as that module.
WEAPON_TYPE_BY_SUBCLASS: dict[int, str] = {
    0: "axe",
    1: "axe",
    4: "mace",
    5: "mace",
    6: "polearm",
    7: "sword",
    8: "sword",
    10: "staff",
    13: "fist",
    15: "dagger",
    2: "bow",
    3: "gun",
    16: "thrown",
    18: "crossbow",
    19: "wand",
}

#: The five ranged `weapon_type` values -- the contract test's own vocabulary
#: ("every row with slot: ranged carries one of the five ranged types").
RANGED_WEAPON_TYPES = frozenset({"wand", "bow", "gun", "crossbow", "thrown"})


def weapon_type_for(item_class_id: int, subclass_id: int, inventory_type: int) -> str | None:
    """The row's `weapon_type`, or None for a non-weapon row or a weapon
    whose SubclassID `WEAPON_TYPE_BY_SUBCLASS` does not map (a fishing pole
    or a misfiled row `is_weapon_row`'s own doc already excludes from dps
    scoring) -- never guessed at, same policy as `STAT_BY_MODIFIER_ID`."""
    if not is_weapon_row(item_class_id, inventory_type):
        return None
    return WEAPON_TYPE_BY_SUBCLASS.get(subclass_id)


def is_weapon_row(item_class_id: int, inventory_type: int) -> bool:
    """True when a row is a real weapon: Item.ClassID 2 in a weapon-only
    InventoryType slot. `weapon_fields` is only meaningful for such a row --
    a non-weapon that happens to state a nonzero ItemDelay (Father Flame,
    an off-hand HOLDABLE item, states 2000) is not a weapon and must not
    get a speed or damage from it, or a picker that ranks by dps equips it
    as one and hands the engine a "weapon" with a speed the engine's own
    simdb (gated the same way, see pipeline/simdb/items.py's
    `class_id == ITEM_CLASS_WEAPON` check) says is zero -- a zero-speed
    swing loops forever.
    """
    return item_class_id == WEAPON and inventory_type in WEAPON_INVENTORY_TYPES


#: Stat keys that are a combat-rating point count when ItemSparse states
#: them (`STAT_BY_MODIFIER_ID`'s 12/13/14/15/31/32/48) but a flat literal
#: percentage when an on-equip spell states them instead
#: (`pipeline.simdb.equip.STAT_AURAS`'s 47/49/51/52/54, plus the
#: MOD_SKILL/defense branch) -- see data/README.md, "Hit, crit, dodge, parry
#: and block as percentages". The two units cannot be summed into one
#: `stats` entry; `_merge_effect_stats` raises rather than do it.
RATING_FAMILY_STAT_KEYS = frozenset({"hit", "crit", "dodge", "parry", "block", "defense"})


def percent_stats_of(equip_stats: dict[str, int]) -> dict[str, int]:
    """The rating-family entries of an on-equip spell's stats: literal
    percentages, recorded as `GearItem.percent_stats` so a consumer never
    divides them by a rating factor."""
    return {key: amount for key, amount in equip_stats.items() if key in RATING_FAMILY_STAT_KEYS}


#: `resolve_required_level`'s item-level-proxy offset: `sim/leveling.
#: ItemLevelProxyRequiredLevel`'s own formula, `item_level - 5`, kept as a
#: named constant here rather than inlined so the Go file's comment ("the
#: two cannot share code across languages, only the number") has one Python
#: number to point at.
REQUIRED_LEVEL_ITEM_LEVEL_PROXY_OFFSET = 5


def resolve_required_level(
    client_level: int, item_level: int, wowhead_level: int | None
) -> tuple[int, str]:
    """The level gate one gear row really has, and which of four sources
    produced it -- `models.GearItem.required_level_source`'s own doc has the
    full rationale; this is the precedence in code:

    1. `client_level`, when non-zero -- the shipped or hotfix-merged row
       states its own gate. Source `"client"`.
    2. `wowhead_level`, when `client_level` is 0 (or the row has no client
       side at all -- a wowhead-supplement item, whose own `required_level`
       IS `wowhead_level`) and wowhead names a non-zero level for this id.
       Source `"wowhead"`.
    3. The item-level proxy, for real gear (`item_level > 1`) neither of the
       above resolved: `item_level - REQUIRED_LEVEL_ITEM_LEVEL_PROXY_OFFSET`,
       floored at 0 and capped at `MAX_PLAYER_LEVEL` -- the same formula
       `sim/leveling.ItemLevelProxyRequiredLevel` computes independently in
       Go. Source `"item_level_proxy"`.
    4. 0, for an `item_level`-1 row (or anything else with no proxy to
       compute) -- required_level 0 IS the right answer here, not a gap.
       Source `"none"`.
    """
    if client_level:
        return client_level, "client"
    if wowhead_level:
        return wowhead_level, "wowhead"
    if item_level > 1:
        proxy = item_level - REQUIRED_LEVEL_ITEM_LEVEL_PROXY_OFFSET
        return max(0, min(MAX_PLAYER_LEVEL, proxy)), "item_level_proxy"
    return 0, "none"


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


#: What every non-weapon row gets instead of `weapon_fields`' curve/literal
#: resolution -- see `is_weapon_row`. Real zeroes, not "unknown": a
#: non-weapon has no damage and no swing speed, full stop.
NOT_A_WEAPON = WeaponFields(damage_min=0, damage_max=0, speed=0.0, dps=0.0, two_hand=False)


def weapon_fields(
    row: dict[str, str],
    subclass_id: int,
    curves: WeaponCurves | None = None,
) -> WeaponFields:
    """Weapon damage, speed and the two-handed flag for one ItemSparse row.

    Assumes the row is already known to be a real weapon -- call this only
    when `is_weapon_row(item_class_id, inventory_type)` is True; every other
    row gets `NOT_A_WEAPON` instead. Nothing here checks Item.ClassID or
    restricts InventoryType to a weapon slot, so calling it on a non-weapon
    row (an off-hand HOLDABLE item that happens to state a nonzero
    ItemDelay, say) reads that column as a swing speed anyway and hands out
    a real dps for an item that has none.

    Classic Era's ItemSparse states damage outright (`MinDamage_0`/
    `MaxDamage_0`); the 1.60 client (Forever beta) states neither and leaves
    it to the `ItemDamage<kind>` curve tables, exactly as armour and stats do
    in `item_curves.py`. `curves` resolves that case; pass None (the
    default) or a `WeaponCurves` whose own tables are incomplete and such a
    row gets damage_min = damage_max = dps = 0, only speed and two_hand
    real -- the same shape a row with neither column reports.

    `pipeline/simdb/weapons.py` resolves the simulator's own weapon damage
    off the same curve tables, through `pipeline.normalize.weapon_curves`'s
    `resolve_weapon_damage` -- this function's curve branch below is the same
    call, not a second copy of the formula. `simdb/weapons.py` already
    imports `column_value`/`int_column` from this module, so importing it
    back here would be a circular import; `weapon_curves.py` is a module
    neither of the two imports, which is why the shared pieces live there.
    """
    delay = _optional_int(row, "ItemDelay") or 0
    speed = round(delay / 1000, 2)
    two_hand = int_column(row, "InventoryType") in TWO_HAND_INVENTORY_TYPES
    if row_has_literal_weapon_damage(row):
        damage_min = _optional_int(row, "MinDamage_0") or 0
        damage_max = _optional_int(row, "MaxDamage_0") or 0
    else:
        damage = _curve_weapon_damage(row, subclass_id, speed, curves)
        damage_min = int(damage.minimum) if damage is not None else 0
        damage_max = int(damage.maximum) if damage is not None else 0
    # Never divide by a zero speed: an item with damage and no delay is a
    # thrown weapon or a malformed row, and either way it has no dps.
    dps = round((damage_min + damage_max) / 2 / speed, 2) if speed > 0 else 0.0
    return WeaponFields(damage_min, damage_max, speed, dps, two_hand)


def _curve_weapon_damage(
    row: dict[str, str],
    subclass_id: int,
    speed: float,
    curves: WeaponCurves | None,
) -> WeaponDamage | None:
    """`weapon_fields`' curve-resolved damage, or None when there is nothing
    to resolve (no curves passed, an incomplete curve set, or a build whose
    schema carries no `DmgVariance` at all -- the same "absent column is a
    schema, not a malformed row" rule `_optional_int` documents)."""
    if curves is None or "DmgVariance" not in row:
        return None
    return resolve_weapon_damage(
        curves,
        subclass_id,
        inventory_type=int_column(row, "InventoryType"),
        item_level=int_column(row, "ItemLevel"),
        quality=int_column(row, "OverallQualityID"),
        speed=speed,
        variance=float(column_value(row, "DmgVariance")),
    )


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


def _has_gear_value(
    armor: int, stats: dict[str, int], item_class_id: int, subclass_id: int, effect_text: str
) -> bool:
    """True when the item carries something the planner can compare.

    An item with no armour, no non-zero stat and no effect text gives the
    planner nothing to reason about, so it is dropped -- for armour and for
    every other item class.

    Weapons (`Item.ClassID` 2) are exempt. A weapon's value is its damage,
    which lives in `weapon_fields`/`is_weapon_row` output the caller checks
    separately, not in `armor`/`stats` here -- judging a weapon on armour
    and stats alone would drop real gear: Annihilator, Arcanite Champion and
    the Hakkari warblades all carry nothing but their damage. A weapon
    therefore survives on the quality and junk-name clauses alone. A real
    client weapon the `ItemDamage*` curve tables state no dps for (or a
    curve-unavailable build) keeps damage 0 and is still not dropped here --
    see `build_class_items`.

    Relics (`Item.ClassID` 4/ARMOR, `SubclassID` in `RELIC_SUBCLASSES` --
    libram, idol, totem) get the same exemption for the same reason: a
    relic's value is an on-equip spell effect (cast-speed/proc/dummy auras),
    not armour or a flat stat. Judging one on armour and stats alone dropped
    every relic on this build -- 65 rows on build 1.60.1.70009, emptying
    paladin/druid/shaman's `ranged` slot at every leveling band (see
    `SLOT_BY_INVENTORY_TYPE`'s InventoryType 28 comment). A relic survives on
    the quality and junk-name clauses alone, same as a weapon; whether its
    effect is one the simulator can model at all is judged downstream by
    `sim/cmd/leveling-bis`'s ranker, not here.

    An `effect_text` (`EffectIndex.text`, a use or proc description this
    build's build_class_items resolves before this call) also exempts any
    item class: a trinket or off-hand rod whose whole value is a use/proc
    effect and no plain stat -- Carrot on a Stick (11122, "Use: dismounts
    ... increases speed"), Six Demon Bag (7734), Tidal Charm (1404),
    Antipodean Rod/Gossamer Rod/Abjurer's Crystal (off-hand rods whose
    effect is a ranged spell attack) -- carries nothing in `armor`/`stats`
    at all, only its effect. Missing this exemption dropped 85 such items
    from every items/<class>.json on the full rebuild at commit a4194d05
    (catalogue-completeness lane, 2026-09-29): a client-shaped row's
    `effect_text` was only ever populated after this gate ran, so every one
    of these real, effect-bearing items looked indistinguishable from a
    genuinely valueless row and was dropped. A relic or weapon still does
    not need this clause -- their own exemptions above already cover the
    case where they also happen to carry no effect text -- but a trinket,
    ring, off-hand rod or any other slot with an effect and nothing else
    now survives here too.
    """
    if item_class_id == WEAPON:
        return True
    if item_class_id == ARMOR and subclass_id in RELIC_SUBCLASSES:
        return True
    if effect_text:
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


def _icon_name(
    item_row: dict[str, str],
    icons: dict[int, str],
    display_name: str,
    fork_icons: dict[int, str],
    wowhead_icons: dict[int, str],
    icon_origins: Counter[str],
) -> tuple[str, str]:
    """The item's icon name and where it came from, falling back past the
    client's placeholder art.

    Some client rows carry `IconFileDataID` 0, meaning the client itself ships
    no art for the item; see `pipeline.icons.resolve_icon` for what happens
    then, and `pipeline.icons.resolve_icon_name` for the fork-db/wowhead
    fallback chain applied on top of it here. `icon_origins` is mutated with
    a count per source ("client", "fork", "wowhead"), for `build_class_items`
    to log a per-build summary; the same per-row origin is now also returned
    for `GearItem.icon_source` (classicdb-fidelity lane, 2026-09-30).
    """
    base = resolve_icon(
        int_column(item_row, "IconFileDataID"),
        icons,
        f"item {item_row.get('ID', '?')} ({display_name})",
    )
    icon, origin = resolve_icon_name(base, int_column(item_row, "ID"), fork_icons, wowhead_icons)
    icon_origins[origin] += 1
    return icon, origin


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
    weapon_curves: WeaponCurves | None = None,
    fork_icons: dict[int, str] | None = None,
    wowhead_icons: dict[int, str] | None = None,
    wowhead_required_levels: dict[int, int] | None = None,
) -> list[ClassItems]:
    """One equippable item list per class. Raises ItemDataError if a row is unreadable.

    `fork_icons` and `wowhead_icons` are the same fallback chain
    `pipeline.icons_fix.fix_placeholder_icons` applies to an already-committed
    build, applied here instead at normalize time so a build fetched with a
    complete raw export never needs the standalone pass at all. Pass None
    (the default, same as an empty dict) for a caller with neither -- every
    item whose `IconFileDataID` the client states as 0 (or names no row in
    `ManifestInterfaceData`) simply keeps PLACEHOLDER_ICON, exactly as before
    either fallback existed. A one-line summary of how many items resolved
    from each source is logged once the class lists are built.

    `curves` resolves armour and stats for a build whose ItemSparse carries no
    literal amounts (the 1.60 client / Forever beta); pass None (the default)
    or an `ItemCurves` whose own tables are incomplete and such a row simply
    gets no armour and no stats, exactly as before curve support existed.

    `weapon_curves` is the same idea for a weapon's damage (`weapon_fields`);
    pass None (the default) or a `WeaponCurves` whose own tables are
    incomplete and such a weapon keeps damage_min = damage_max = dps = 0 --
    a real client weapon the `ItemDamage*` tables state no dps for stays at
    zero and is not dropped (see `_has_gear_value`'s weapon exemption); a
    caller counts those to report the before/after zero-damage rate.

    `effects` folds an item's on-equip spell stats (`pipeline.normalize.
    effects.EffectIndex`) into the same `stats` dict ItemSparse's own columns
    populate, and sets `effect_text` from its use/proc spells. Pass None (the
    default) for a caller that has not built one and every item gets no
    extra stats and an empty `effect_text`, exactly as before this existed.

    `wowhead_required_levels` (item id -> wowhead's own `required_level`, every
    id the payload names, not only the ones the client lacks -- see
    `pipeline.icons_fix.load_wowhead_icons`'s own doc for why "every id" and
    not just the supplement) resolves `required_level`/`required_level_source`
    per `resolve_required_level`. Pass None (the default, same as an empty
    dict) for a caller with no wowhead payload -- every row's `required_level`
    then falls back to the item-level proxy or 0, never wowhead's number.
    """
    by_id = {int_column(row, "ID"): row for row in item_rows}
    fork_icons = fork_icons or {}
    wowhead_icons = wowhead_icons or {}
    wowhead_required_levels = wowhead_required_levels or {}
    icon_origins: Counter[str] = Counter()
    candidates: list[tuple[GearItem, int, int, int]] = []
    for row in sparse_rows:
        inventory_type = int_column(row, "InventoryType")
        slot = SLOT_BY_INVENTORY_TYPE.get(inventory_type)
        if slot is None:
            continue
        client_required_level = int_column(row, "RequiredLevel")
        if client_required_level > MAX_PLAYER_LEVEL:
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
        percent_stats: dict[str, int] = {}
        if effects is not None:
            _merge_effect_stats(stats, item_id, display_name, effects)
            percent_stats = percent_stats_of(effects.stats(item_id))
        _check_level_60_sanity(item_id, display_name, item_level, armor, stats)
        effect_text = "" if effects is None else effects.text(item_id)
        if not _has_gear_value(armor, stats, item_class_id, subclass_id, effect_text):
            continue
        weapon = (
            weapon_fields(row, subclass_id, weapon_curves)
            if is_weapon_row(item_class_id, inventory_type)
            else NOT_A_WEAPON
        )
        required_level, required_level_source = resolve_required_level(
            client_required_level, item_level, wowhead_required_levels.get(item_id)
        )
        icon, icon_source = _icon_name(
            item_row, icons, display_name, fork_icons, wowhead_icons, icon_origins
        )
        item = GearItem(
            id=item_id,
            name=display_name,
            icon=icon,
            slot=slot,
            quality=quality,
            required_level=required_level,
            required_level_source=required_level_source,
            item_level=item_level,
            armor=armor,
            stats=stats,
            percent_stats=percent_stats,
            damage_min=weapon.damage_min,
            damage_max=weapon.damage_max,
            speed=weapon.speed,
            dps=weapon.dps,
            two_hand=weapon.two_hand,
            effect_text=effect_text,
            set_id=int_column(row, "ItemSet") or None,
            unique=int_column(row, "MaxCount") == 1,
            icon_source=icon_source,
            weapon_type=weapon_type_for(item_class_id, subclass_id, inventory_type),
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
    if icon_origins["fork"] or icon_origins["wowhead"]:
        logger.info(
            "items %s icon origins: %d from the client, %d from the fork db, "
            "%d from wowhead's gear-planner payload",
            build,
            icon_origins["client"],
            icon_origins["fork"],
            icon_origins["wowhead"],
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
