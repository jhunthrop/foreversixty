"""Which weapons and armour each class can equip, and from what level, from the client.

Written to `builds/<build>/equip-skills.json` for the guild Loot tab, which lists a
character as a candidate for a drop only when the character can equip it
(`api/internal/guilds/equip_rules_table.go` is generated from this file by
`go generate ./internal/guilds/` in `api/`).

**Sources**, all three from wago.tools for the build (`fetch` writes them to `raw/`):

- `SkillLine`: the Weapon Skills lines (`CategoryID` 6: Axes, Swords, Dual Wield, ...)
  and the Armor Proficiencies lines (`CategoryID` 8: Cloth, Leather, Mail, Plate Mail,
  Shield).
- `SkillLineAbility`: the spell that teaches each line and the classes that can learn it
  (`ClassMask`). A row whose `AcquireMethod` is 3 is a Season of Discovery rune, which
  Forever does not teach (the shaman's Dual Wield row is one), and a row with no class
  bit names nobody; neither grants the skill.
- `SkillRaceClassInfo`: the level a class learns the line at (`MinLevel`): mail at 40
  for hunters and shamans, plate at 40 for warriors and paladins, Dual Wield at 10 for
  rogues and 20 for warriors and hunters. A class with no row learns it at level 1.

One client bit is dropped: druids in polearms. The client's mask carries it, but no
trainer teaches it (sim audit round 8, K1, from the Classic weapon masters' trainer
data). The file lists it under `dropped` so the next build's data can be checked
against the same evidence.

Rogues keep one-hand axes here (guild round 1, item 8), unlike
`pipeline.proficiency.VESTIGIAL_CLIENT_BITS`, which still drops them for the gear
pickers and sims: the client's Axes mask grants rogues the skill, and Forever's own
rogue talent Hack and Slash (spell 13960, `talents/rogue.json`) reads "Axe/Sword".
The two tables disagree on purpose until the BiS loop settles proficiency.py.
"""

from __future__ import annotations

import json
import logging
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.proficiency import (
    CLASS_MASK_BIT,
    CLASS_SLUG_BY_ID,
    WEAPON_SKILL_LINE,
    WEAPON_SKILL_LINE_NAME,
)

logger = logging.getLogger(__name__)

EQUIP_SKILLS = "equip-skills.json"
TABLES = ("SkillLine", "SkillLineAbility", "SkillRaceClassInfo")

WEAPON_SKILL_CATEGORY = 6
ARMOR_SKILL_CATEGORY = 8

#: Item.SubclassID (ARMOR) -> the Armor Proficiencies line that gates it. Miscellaneous
#: (0: rings, necks, trinkets, held in off-hand) needs no skill.
ARMOR_SKILL_LINE: dict[int, int] = {1: 415, 2: 414, 3: 413, 4: 293, 6: 433}
ARMOR_SKILL_LINE_NAME: dict[int, str] = {
    415: "Cloth",
    414: "Leather",
    413: "Mail",
    293: "Plate Mail",
    433: "Shield",
}

DUAL_WIELD_SKILL_LINE = 118
DUAL_WIELD_NAME = "Dual Wield"

#: (ChrClasses.ID, weapon Item.SubclassID) bits the client's masks carry that no trainer
#: teaches: druid polearms. Rogue one-hand axes stay (see the module docstring).
UNTAUGHT_CLIENT_BITS: frozenset[tuple[int, int]] = frozenset({(11, 6)})

#: SkillLineAbility.AcquireMethod of a Season of Discovery rune: not learnable in Forever.
RUNE_ACQUIRE_METHOD = 3


class EquipSkillsSchemaError(Exception):
    """A fetched table does not carry the skill lines this module expects."""


def _check_skill_lines(skill_lines: dict[int, dict[str, str]]) -> None:
    expected = {
        **{sid: (name, WEAPON_SKILL_CATEGORY) for sid, name in WEAPON_SKILL_LINE_NAME.items()},
        DUAL_WIELD_SKILL_LINE: (DUAL_WIELD_NAME, WEAPON_SKILL_CATEGORY),
        **{sid: (name, ARMOR_SKILL_CATEGORY) for sid, name in ARMOR_SKILL_LINE_NAME.items()},
    }
    for skill_id, (name, category) in expected.items():
        row = skill_lines.get(skill_id)
        if row is None or row["DisplayName_lang"] != name or int(row["CategoryID"]) != category:
            raise EquipSkillsSchemaError(
                f"SkillLine {skill_id} is not {name!r} in category {category} for this build"
            )


def _learnable_class_masks(rows: list[dict[str, str]]) -> dict[int, int]:
    """SkillLine id -> the classes that can learn it (OR of every learnable row's mask)."""
    out: dict[int, int] = {}
    for row in rows:
        if int(row["AcquireMethod"]) == RUNE_ACQUIRE_METHOD:
            continue
        skill_id = int(row["SkillLine"])
        out[skill_id] = out.get(skill_id, 0) | int(row["ClassMask"])
    return out


def _learn_levels(rows: list[dict[str, str]]) -> dict[tuple[int, int], int]:
    """(SkillLine id, ChrClasses id) -> the lowest MinLevel of a row naming that class."""
    out: dict[tuple[int, int], int] = {}
    for row in rows:
        skill_id, mask, level = int(row["SkillID"]), int(row["ClassMask"]), int(row["MinLevel"])
        for class_id, bit in CLASS_MASK_BIT.items():
            if mask & bit:
                key = (skill_id, class_id)
                out[key] = min(out.get(key, level), level)
    return out


def _skill_entries(
    gates: dict[int, int],
    names: dict[int, str],
    class_id: int,
    masks: dict[int, int],
    levels: dict[tuple[int, int], int],
    dropped: set[tuple[int, int]],
) -> list[dict[str, object]]:
    entries = []
    for subclass_id, skill_id in sorted(gates.items()):
        if not masks.get(skill_id, 0) & CLASS_MASK_BIT[class_id]:
            continue
        if (class_id, subclass_id) in dropped:
            continue
        entries.append(
            {
                "subclass": subclass_id,
                "skill": names[skill_id],
                "level": max(levels.get((skill_id, class_id), 1), 1),
            }
        )
    return entries


def build_equip_skills(build: str, raw: Path) -> dict[str, object]:
    """The equip-skills document for `build` from `raw`'s three client tables."""
    skill_lines = {int(r["ID"]): r for r in read_csv(raw / "SkillLine.csv")}
    _check_skill_lines(skill_lines)
    masks = _learnable_class_masks(read_csv(raw / "SkillLineAbility.csv"))
    levels = _learn_levels(read_csv(raw / "SkillRaceClassInfo.csv"))

    classes: dict[str, object] = {}
    for class_id, slug in sorted(CLASS_SLUG_BY_ID.items(), key=lambda kv: kv[1]):
        dual = masks.get(DUAL_WIELD_SKILL_LINE, 0) & CLASS_MASK_BIT[class_id]
        classes[slug] = {
            "armor": _skill_entries(
                ARMOR_SKILL_LINE, ARMOR_SKILL_LINE_NAME, class_id, masks, levels, set()
            ),
            "weapons": _skill_entries(
                WEAPON_SKILL_LINE,
                WEAPON_SKILL_LINE_NAME,
                class_id,
                masks,
                levels,
                set(UNTAUGHT_CLIENT_BITS),
            ),
            "dual_wield_level": (
                max(levels.get((DUAL_WIELD_SKILL_LINE, class_id), 1), 1) if dual else None
            ),
        }
    dropped = [
        {
            "class": CLASS_SLUG_BY_ID[class_id],
            "subclass": subclass_id,
            "skill": WEAPON_SKILL_LINE_NAME[WEAPON_SKILL_LINE[subclass_id]],
            "why": "The client mask carries it, but no trainer teaches it (sim audit round 8, K1).",
        }
        for class_id, subclass_id in sorted(UNTAUGHT_CLIENT_BITS)
        if masks.get(WEAPON_SKILL_LINE[subclass_id], 0) & CLASS_MASK_BIT[class_id]
    ]
    return {
        "build": build,
        "source": [f"https://wago.tools/db2/{table}/csv?build={build}" for table in TABLES],
        "classes": classes,
        "dropped": dropped,
    }


def write_equip_skills(build: str, root: Path = Path("builds"), raw: Path | None = None) -> Path:
    """Write `builds/<build>/equip-skills.json` from the build's raw client tables."""
    build_dir = root / build
    raw = raw if raw is not None else build_dir / "raw"
    for table in TABLES:
        if not (raw / f"{table}.csv").exists():
            raise SystemExit(f"no {raw / f'{table}.csv'}; run `python -m pipeline fetch` first")
    out = build_dir / EQUIP_SKILLS
    doc = build_equip_skills(build, raw)
    out.write_text(json.dumps(doc, indent=2) + "\n", encoding="utf-8")
    logger.info("wrote %s", out)
    return out
