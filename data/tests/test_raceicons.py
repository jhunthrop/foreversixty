import json
from pathlib import Path

import pytest

from pipeline.raceicons import RACE_ICONS, race_icons

#: Every race a Forever player can pick: the eight classic races and the two Skyborne.
FOREVER_RACES = (
    "human",
    "orc",
    "dwarf",
    "night-elf",
    "undead",
    "tauren",
    "gnome",
    "troll",
    "high-order-skyborne",
    "windshaper-skyborne",
)

CHR_RACES = [
    "1,Human,0",
    "2,Orc,0",
    "3,Dwarf,0",
    "4,Night Elf,0",
    "5,Undead,0",
    "6,Tauren,0",
    "7,Gnome,0",
    "8,Troll,0",
    "10,Blood Elf,0",
    "95,High Order Skyborne,10",
    "96,Windshaper Skyborne,10",
    "97,Nowhere Folk,0",
]

PORTRAITS = ["human", "orc", "dwarf", "nightelf", "undead", "tauren", "gnome", "troll", "bloodelf"]


def _csv(path: Path, header: str, rows: list[str]) -> None:
    path.write_text("\n".join([header, *rows]) + "\n", encoding="utf-8")


def _raw(tmp_path: Path) -> Path:
    _csv(tmp_path / "ChrRaces.csv", "ID,Name_lang,MaleModelFallbackRaceID", CHR_RACES)
    manifest = [
        f"{100 + index},Interface\\ICONS\\,Achievement_Character_{name.title()}_Male.blp"
        for index, name in enumerate(PORTRAITS)
    ]
    female = "200,Interface\\ICONS\\,Achievement_Character_Human_Female.blp"
    _csv(tmp_path / "ManifestInterfaceData.csv", "ID,FilePath,FileName", [*manifest, female])
    return tmp_path


RACE_ROWS: list[dict[str, object]] = [
    {"id": race_id, "slug": slug}
    for race_id, slug in zip((1, 2, 3, 4, 5, 6, 7, 8, 95, 96), FOREVER_RACES, strict=True)
]


def test_every_forever_race_gets_its_male_portrait(tmp_path):
    icons = race_icons(_raw(tmp_path), RACE_ROWS)
    assert set(icons) == set(FOREVER_RACES)
    assert icons["night-elf"] == "achievement_character_nightelf_male"
    assert icons["undead"] == "achievement_character_undead_male"
    assert icons["human"] == "achievement_character_human_male"


def test_skyborne_take_the_portrait_of_the_race_the_client_draws_them_with(tmp_path):
    icons = race_icons(_raw(tmp_path), RACE_ROWS)
    assert icons["high-order-skyborne"] == "achievement_character_bloodelf_male"
    assert icons["windshaper-skyborne"] == "achievement_character_bloodelf_male"


def test_a_race_with_no_portrait_anywhere_is_left_out(tmp_path):
    icons = race_icons(_raw(tmp_path), [{"id": 97, "slug": "nowhere-folk"}])
    assert icons == {}


def test_the_committed_race_icons_cover_every_forever_race():
    active_build = Path("../web/src/data/active-build.json").read_text(encoding="utf-8")
    active = json.loads(active_build)["build"]
    path = Path(f"builds/{active}/{RACE_ICONS}")
    if not path.exists():
        pytest.skip(f"{path} is written by the next data run")
    icons = json.loads(path.read_text(encoding="utf-8"))
    races = json.loads(Path(f"builds/{active}/races.json").read_text(encoding="utf-8"))
    assert set(FOREVER_RACES) <= set(icons)
    assert {race["slug"] for race in races} <= set(icons)
    for icon in icons.values():
        assert Path(f"builds/{active}/icons/{icon}.webp").exists()
