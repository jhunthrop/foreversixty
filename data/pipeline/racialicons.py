"""Racial ability icons, from the client: racial name -> icon name.

The guides' Races table names each race's racials (Blood Fury, Eureka!, Read Ley Line) by
the name the race data uses. The client lists every racial on a racial skill line
(`SkillLine.CategoryID == 9`: "Orc Racial", "Racial - Gnome", "Skyborne Racial") in
`SkillLineAbility`, and the spell's own `SpellMisc.SpellIconFileDataID` names its icon,
resolved through `ManifestInterfaceData` like every other icon in this pipeline. A racial
with one spell per class (Eureka!, Expansive Mind) takes the icon of its lowest spell id.

`builds/<build>/racialicons.json` is written by `python -m pipeline icons` beside
`spellicons.json`, from the same `raw/` tables, and the icon art is downloaded with the
rest, so a later run with no `raw/` reads the committed file back.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.icons import PLACEHOLDER_ICON, icon_names, resolve_icon

#: SkillLine.CategoryID of the racial skill lines.
RACIAL_SKILL_CATEGORY = 9

#: Category 9 also holds the secondary skills (Cooking, Fishing, languages); the racial
#: lines are the ones the client names "... Racial" or "Racial - ...".
RACIAL_LINE_WORD = "Racial"

RACIAL_ICONS = "racialicons.json"

#: The raw client tables `racial_icons` reads; `icons_for_build` skips the file without them.
RACIAL_TABLES = (
    "SkillLine.csv",
    "SkillLineAbility.csv",
    "SpellName.csv",
    "SpellMisc.csv",
    "ManifestInterfaceData.csv",
)


def racial_icons(raw: Path) -> dict[str, str]:
    """Racial name -> icon name, for every spell on a racial skill line, sorted by name."""
    racial_lines = {
        int(row["ID"])
        for row in read_csv(raw / "SkillLine.csv")
        if int(row["CategoryID"] or 0) == RACIAL_SKILL_CATEGORY
        and RACIAL_LINE_WORD in row["DisplayName_lang"]
    }
    names = {int(row["ID"]): row["Name_lang"] for row in read_csv(raw / "SpellName.csv")}
    icon_ids = {
        int(row["SpellID"]): int(row["SpellIconFileDataID"] or 0)
        for row in read_csv(raw / "SpellMisc.csv")
    }
    manifest = icon_names(read_csv(raw / "ManifestInterfaceData.csv"))
    spells = sorted(
        int(row["Spell"])
        for row in read_csv(raw / "SkillLineAbility.csv")
        if int(row["SkillLine"] or 0) in racial_lines
    )
    icons: dict[str, str] = {}
    for spell_id in spells:
        name = names.get(spell_id, "")
        file_id = icon_ids.get(spell_id, 0)
        if name == "" or name in icons or file_id == 0:
            continue
        icon = resolve_icon(file_id, manifest, f"racial {spell_id} ({name})")
        if icon != PLACEHOLDER_ICON:
            icons[name] = icon
    return dict(sorted(icons.items()))


def write_racial_icons(build: str, root: Path = Path("builds")) -> Path:
    """Write `builds/<build>/racialicons.json` from the build's `raw/` tables."""
    path = root / build / RACIAL_ICONS
    payload = racial_icons(root / build / "raw")
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    return path
