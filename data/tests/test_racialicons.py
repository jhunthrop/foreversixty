import json
from pathlib import Path

from pipeline.racialicons import racial_icons


def _csv(path: Path, header: str, rows: list[str]) -> None:
    path.write_text("\n".join([header, *rows]) + "\n", encoding="utf-8")


def test_racial_icons_reads_only_the_racial_skill_lines(tmp_path):
    _csv(
        tmp_path / "SkillLine.csv",
        "DisplayName_lang,ID,CategoryID",
        ["Racial - Gnome,753,9", "Cooking,185,9", "Fire,8,7"],
    )
    _csv(
        tmp_path / "SkillLineAbility.csv",
        "ID,SkillLine,Spell",
        ["1,753,1259813", "2,753,1259812", "3,185,2550", "4,8,133"],
    )
    _csv(
        tmp_path / "SpellName.csv",
        "ID,Name_lang",
        ["1259813,Eureka!", "1259812,Eureka!", "2550,Cooking", "133,Fireball"],
    )
    _csv(
        tmp_path / "SpellMisc.csv",
        "ID,SpellIconFileDataID,SpellID",
        ["1,100,1259812", "2,101,1259813", "3,102,2550", "4,103,133"],
    )
    _csv(
        tmp_path / "ManifestInterfaceData.csv",
        "ID,FilePath,FileName",
        [
            "100,Interface/Icons/,inv_gnometoy.blp",
            "101,Interface/Icons/,other.blp",
            "102,Interface/Icons/,cook.blp",
            "103,Interface/Icons/,fire.blp",
        ],
    )
    # One icon per racial name, from its lowest spell id; Cooking and Fireball are not racials.
    assert racial_icons(tmp_path) == {"Eureka!": "inv_gnometoy"}


def test_the_committed_racial_icons_name_the_offensive_racials():
    active_build = Path("../web/src/data/active-build.json").read_text(encoding="utf-8")
    active = json.loads(active_build)["build"]
    icons = json.loads(Path(f"builds/{active}/racialicons.json").read_text(encoding="utf-8"))
    for name in ("Blood Fury", "Berserking", "Eureka!", "Will of the Forsaken"):
        assert name in icons
