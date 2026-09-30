"""Which armour and weapon subclasses each class can equip.

Item.ClassID 4 is armour with subclasses
    0 miscellaneous (rings, necks, trinkets)  1 cloth (includes cloaks)
    2 leather  3 mail  4 plate  6 shield  7 libram  8 idol  9 totem
Item.ClassID 2 is weapons with subclasses
    0 one-hand axe   1 two-hand axe   2 bow          3 gun
    4 one-hand mace  5 two-hand mace  6 polearm      7 one-hand sword
    8 two-hand sword 10 staff         13 fist        15 dagger
    16 thrown        18 crossbow      19 wand
Subclasses 14 (monster items), 17 (deprecated spear) and 20 (fishing pole)
are deliberately absent: they are not player gear.

**Armour** proficiency has no client export the pipeline reads today: a
character learns it from trainer spells, so `ARMOR_SUBCLASSES` below stays a
hand-written table of the World of Warcraft Classic 1.x proficiencies, keyed
by ChrClasses.ID. Revisit it when the Forever client ships its own class
data; a wrong entry shows up as an item missing from, or wrongly offered in,
one class's gear picker.

**Weapon** proficiency does have a client export: `SkillLineAbility`'s
`ClassMask` column, joined on `SkillLine.CategoryID == 6` ("Weapon Skills"),
names exactly which classes may train (and therefore equip) each weapon
subclass. `configure()` derives `WEAPON_SUBCLASSES`'s runtime table from
`raw/SkillLineAbility.csv` + `raw/SkillLine.csv` when both are present (a
normal `fetch` for build 1.60.1.70009 downloads them) and falls back to the
hand-written `WEAPON_SUBCLASSES` constant otherwise -- the same 1.x table
as before, corrected (2026-09-30, day3 data-followups-5 lane) for two
classes a wow-player sweep caught publishing items nobody can wear:
paladins had one- and two-hand axes (they don't -- paladins train maces,
swords and polearms) and druids had polearms (they don't -- druids train
daggers, fist weapons, maces and staves). Forever's own two documented
deviations from vanilla (shamans train one- and two-hand axes and maces;
rogues train maces; `research/01-official-facts.md`) are unaffected.

Cross-checking every other weapon subclass's `ClassMask` in build
1.60.1.70009 against that same 1.x-plus-Forever's-two-changes table found
them an exact match, with three exceptions the client's raw `ClassMask`
carries but that were never reachable in play: Paladin in Axes (44) and
Two-Handed Axes (172), Druid in Polearms (229), and Rogue in Axes (44). The
first two are exactly the bug this lane fixes; the third is not a contract
in this lane's brief but agrees with today's already-shipped (never
challenged) hand-written table, so it is excluded too rather than silently
granting rogues an axe nothing asked for. `_VESTIGIAL_CLIENT_BITS` documents
all three with this same reasoning so a future build's data can be checked
against it rather than re-derived from scratch. `SkillRaceClassInfo` (the
brief's other named candidate) was tried first and rejected: its
`Availability` column does not cleanly separate real from vestigial rows
either way -- filtering to `Availability == 1` drops Paladin, Mage and
Warlock from Swords entirely, which is wrong for all three.
"""

from __future__ import annotations

import csv
import logging
from pathlib import Path

logger = logging.getLogger(__name__)

WEAPON = 2
ARMOR = 4

_MISC = 0
_CLOTH = 1
_LEATHER = 2
_MAIL = 3
_PLATE = 4
_SHIELD = 6
_LIBRAM = 7
_IDOL = 8
_TOTEM = 9

#: Relic subclasses (Item.ClassID 4/ARMOR, SubclassID 7/8/9) -- the single
#: source of truth `pipeline.normalize.gear._has_gear_value` reads from so a
#: relic's exemption there and its class filter here never drift apart. A
#: relic carries its value in an on-equip spell effect, not armour or a
#: literal stat, so it needs the same "not judged on armour/stats" exemption
#: the WEAPON class id already gets in _has_gear_value.
RELIC_SUBCLASSES: frozenset[int] = frozenset({_LIBRAM, _IDOL, _TOTEM})

#: Relic subclasses (7 libram, 8 idol, 9 totem) were absent from every
#: class's set below entirely until the 2026-09-28 night-bis-sources
#: lane, which (together with `pipeline.normalize.gear.
#: SLOT_BY_INVENTORY_TYPE` lacking InventoryType 28) silently dropped
#: every relic from every class's items/<class>.json, emptying paladin's
#: ranged slot at every leveling band. Only the class trainer spells
#: actually give a relic to gets it: a libram is paladin-only, an idol
#: druid-only, a totem shaman-only.
ARMOR_SUBCLASSES: dict[int, frozenset[int]] = {
    1: frozenset({_MISC, _CLOTH, _LEATHER, _MAIL, _PLATE, _SHIELD}),  # warrior
    2: frozenset({_MISC, _CLOTH, _LEATHER, _MAIL, _PLATE, _SHIELD, _LIBRAM}),  # paladin
    3: frozenset({_MISC, _CLOTH, _LEATHER, _MAIL}),  # hunter
    4: frozenset({_MISC, _CLOTH, _LEATHER}),  # rogue
    5: frozenset({_MISC, _CLOTH}),  # priest
    7: frozenset({_MISC, _CLOTH, _LEATHER, _MAIL, _SHIELD, _TOTEM}),  # shaman
    8: frozenset({_MISC, _CLOTH}),  # mage
    9: frozenset({_MISC, _CLOTH}),  # warlock
    11: frozenset({_MISC, _CLOTH, _LEATHER, _IDOL}),  # druid
}

#: Fallback only -- see `configure()`. Paladin carries no axe (0, 1) and
#: druid no polearm (6): the 2026-09-30 fix (see module docstring).
WEAPON_SUBCLASSES: dict[int, frozenset[int]] = {
    1: frozenset({0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 13, 15, 16, 18}),  # warrior
    2: frozenset({4, 5, 6, 7, 8}),  # paladin
    3: frozenset({0, 1, 2, 3, 6, 7, 8, 10, 13, 15, 16, 18}),  # hunter
    4: frozenset({2, 3, 4, 7, 13, 15, 16, 18}),  # rogue
    5: frozenset({4, 10, 15, 19}),  # priest
    7: frozenset({0, 1, 4, 5, 10, 13, 15}),  # shaman
    8: frozenset({7, 10, 15, 19}),  # mage
    9: frozenset({7, 10, 15, 19}),  # warlock
    11: frozenset({4, 5, 10, 13, 15}),  # druid
}

#: Item.SubclassID (WEAPON) -> the SkillLine.ID that gates it, restricted to
#: `SkillLine.CategoryID == 6` ("Weapon Skills"). Verified against build
#: 1.60.1.70009's own `SkillLine.csv` -- see `_WEAPON_SKILL_LINE_NAME`, which
#: `configure()` cross-checks at run time so a build that renumbers these
#: ids fails loudly rather than deriving a silently wrong table.
_WEAPON_SKILL_LINE: dict[int, int] = {
    0: 44,
    1: 172,
    2: 45,
    3: 46,
    4: 54,
    5: 160,
    6: 229,
    7: 43,
    8: 55,
    10: 136,
    13: 473,
    15: 173,
    16: 176,
    18: 226,
    19: 228,
}

_WEAPON_SKILL_CATEGORY = 6  # SkillLine.CategoryID for "Weapon Skills"

_WEAPON_SKILL_LINE_NAME: dict[int, str] = {
    44: "Axes",
    172: "Two-Handed Axes",
    45: "Bows",
    46: "Guns",
    54: "Maces",
    160: "Two-Handed Maces",
    229: "Polearms",
    43: "Swords",
    55: "Two-Handed Swords",
    136: "Staves",
    473: "Fist Weapons",
    173: "Daggers",
    176: "Thrown",
    226: "Crossbows",
    228: "Wands",
}

#: ChrClasses.ID -> its bit in SkillLineAbility/SkillRaceClassInfo's
#: ClassMask (`1 << (ChrClasses.ID - 1)`, the same convention
#: `pipeline.normalize.gear.build_class_items` uses for `AllowableClass`).
_CLASS_MASK_BIT: dict[int, int] = {
    1: 1,  # warrior
    2: 2,  # paladin
    3: 4,  # hunter
    4: 8,  # rogue
    5: 16,  # priest
    7: 64,  # shaman
    8: 128,  # mage
    9: 256,  # warlock
    11: 1024,  # druid
}

#: (ChrClasses.ID, Item.SubclassID) pairs `SkillLineAbility`'s `ClassMask`
#: carries for build 1.60.1.70009 but that were never reachable in play --
#: see the module docstring for the evidence trail. `configure()` drops
#: these out of the client-derived table the same way the hand-written
#: `WEAPON_SUBCLASSES` fallback already omits them.
_VESTIGIAL_CLIENT_BITS: frozenset[tuple[int, int]] = frozenset(
    {
        (2, 0),  # paladin, one-hand axe
        (2, 1),  # paladin, two-hand axe
        (11, 6),  # druid, polearm
        (4, 0),  # rogue, one-hand axe
    }
)


class Db2SchemaError(Exception):
    """Raised when a fetched DB2 CSV doesn't match this module's expectations."""


def _derive_weapon_subclasses_from_client(raw: Path) -> dict[int, frozenset[int]]:
    """Build `WEAPON_SUBCLASSES`'s shape from `raw`'s `SkillLine.csv` +
    `SkillLineAbility.csv`. Raises `Db2SchemaError`/`KeyError`/`ValueError` on
    anything that doesn't match this build's known schema; `configure()` is
    the only caller and treats every one of those as "fall back"."""
    with (raw / "SkillLine.csv").open(newline="", encoding="utf-8") as f:
        skill_lines = {int(row["ID"]): row for row in csv.DictReader(f)}
    for skill_id, name in _WEAPON_SKILL_LINE_NAME.items():
        row = skill_lines.get(skill_id)
        if (
            row is None
            or row["DisplayName_lang"] != name
            or int(row["CategoryID"]) != _WEAPON_SKILL_CATEGORY
        ):
            raise Db2SchemaError(
                f"SkillLine {skill_id} is not {name!r} in category "
                f"{_WEAPON_SKILL_CATEGORY} for this build"
            )

    skill_class_mask: dict[int, int] = {}
    with (raw / "SkillLineAbility.csv").open(newline="", encoding="utf-8") as f:
        for row in csv.DictReader(f):
            skill_id = int(row["SkillLine"])
            if skill_id not in _WEAPON_SKILL_LINE_NAME:
                continue
            skill_class_mask[skill_id] = skill_class_mask.get(skill_id, 0) | int(
                row["ClassMask"]
            )

    derived: dict[int, set[int]] = {class_id: set() for class_id in _CLASS_MASK_BIT}
    for subclass_id, skill_id in _WEAPON_SKILL_LINE.items():
        mask = skill_class_mask.get(skill_id, 0)
        for class_id, bit in _CLASS_MASK_BIT.items():
            if mask & bit and (class_id, subclass_id) not in _VESTIGIAL_CLIENT_BITS:
                derived[class_id].add(subclass_id)
    return {class_id: frozenset(subclasses) for class_id, subclasses in derived.items()}


#: The table `can_equip` reads for WEAPON, and the source that produced it.
#: Module-level rather than threaded through every `can_equip` call site
#: deliberately: `can_equip` is called from `pipeline/normalize/gear.py`,
#: `pipeline/normalize/classicdb.py`, `pipeline/normalize/wowhead.py` and
#: `pipeline/simdb/items.py`, none of which otherwise carry a build's `raw`
#: path down to this check -- threading one through all four for a single
#: build-wide setting would touch `normalize/classicdb.py` far beyond this
#: lane's fix. `configure()` is the one intended mutation point, called once
#: per build by `normalize_build` before any of those run; every other
#: caller is unaffected and keeps reading whichever table is active.
_active_weapon_subclasses: dict[int, frozenset[int]] = WEAPON_SUBCLASSES
_active_weapon_source = "fallback:hardcoded WEAPON_SUBCLASSES"


def configure(raw: Path | None) -> str:
    """Point `can_equip` at `raw`'s client weapon-proficiency tables when
    both exist and parse as expected; otherwise (or when `raw` is None) fall
    back to the hardcoded `WEAPON_SUBCLASSES`. Returns the source string, so
    a caller (`normalize_build`) can log or report which one produced the
    table actually used."""
    global _active_weapon_subclasses, _active_weapon_source
    has_client_tables = (
        raw is not None
        and (raw / "SkillLine.csv").exists()
        and (raw / "SkillLineAbility.csv").exists()
    )
    if raw is not None and has_client_tables:
        try:
            _active_weapon_subclasses = _derive_weapon_subclasses_from_client(raw)
            _active_weapon_source = f"client:{raw / 'SkillLineAbility.csv'}"
            return _active_weapon_source
        except (Db2SchemaError, KeyError, ValueError) as exc:
            logger.warning(
                "weapon proficiency: %s not usable (%s); falling back to hardcoded table",
                raw,
                exc,
            )
    _active_weapon_subclasses = WEAPON_SUBCLASSES
    _active_weapon_source = "fallback:hardcoded WEAPON_SUBCLASSES"
    return _active_weapon_source


def can_equip(class_id: int, item_class_id: int, subclass_id: int) -> bool:
    if item_class_id == ARMOR:
        return subclass_id in ARMOR_SUBCLASSES.get(class_id, frozenset())
    if item_class_id == WEAPON:
        return subclass_id in _active_weapon_subclasses.get(class_id, frozenset())
    return False
