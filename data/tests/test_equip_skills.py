import csv
import json
from pathlib import Path

import pytest

from pipeline.equip_skills import (
    EquipSkillsSchemaError,
    build_equip_skills,
    write_equip_skills,
)

SKILL_LINES = [
    (43, "Swords", 6),
    (44, "Axes", 6),
    (45, "Bows", 6),
    (46, "Guns", 6),
    (54, "Maces", 6),
    (55, "Two-Handed Swords", 6),
    (118, "Dual Wield", 6),
    (136, "Staves", 6),
    (160, "Two-Handed Maces", 6),
    (172, "Two-Handed Axes", 6),
    (173, "Daggers", 6),
    (176, "Thrown", 6),
    (226, "Crossbows", 6),
    (228, "Wands", 6),
    (229, "Polearms", 6),
    (473, "Fist Weapons", 6),
    (293, "Plate Mail", 8),
    (413, "Mail", 8),
    (414, "Leather", 8),
    (415, "Cloth", 8),
    (433, "Shield", 8),
]

#: (SkillLine, ClassMask, AcquireMethod), the build's own masks for the lines the tests read.
ABILITIES = [
    (44, 79, 2),  # axes: warrior, paladin, hunter, rogue (vestigial), shaman
    (229, 1031, 2),  # polearms: warrior, paladin, hunter, druid (vestigial)
    (473, 1101, 2),  # fist weapons: warrior, hunter, rogue, shaman, druid
    (118, 13, 2),  # dual wield: warrior, hunter, rogue
    (118, 64, 3),  # dual wield for shamans: a Season of Discovery rune
    (293, 3, 2),  # plate: warrior, paladin
    (413, 71, 2),  # mail: warrior, paladin, hunter, shaman
    (415, 1503, 2),  # cloth: everyone
    (415, 0, 2),  # a row naming no class
]

#: (SkillID, ClassMask, MinLevel)
RACE_CLASS_INFO = [
    (293, 3, 40),
    (413, 68, 40),
    (413, 3, 0),
    (118, 8, 10),
    (118, 5, 20),
    (118, 64, 0),
    (229, 7, 20),
]


def _write(path: Path, header: list[str], rows: list[dict[str, object]]) -> None:
    with path.open("w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=header)
        writer.writeheader()
        writer.writerows(rows)


@pytest.fixture
def raw(tmp_path: Path) -> Path:
    _write(
        tmp_path / "SkillLine.csv",
        ["DisplayName_lang", "ID", "CategoryID"],
        [{"DisplayName_lang": n, "ID": i, "CategoryID": c} for i, n, c in SKILL_LINES],
    )
    _write(
        tmp_path / "SkillLineAbility.csv",
        ["ID", "SkillLine", "Spell", "ClassMask", "AcquireMethod"],
        [
            {"ID": n, "SkillLine": s, "Spell": 1000 + n, "ClassMask": m, "AcquireMethod": a}
            for n, (s, m, a) in enumerate(ABILITIES)
        ],
    )
    _write(
        tmp_path / "SkillRaceClassInfo.csv",
        ["ID", "SkillID", "ClassMask", "MinLevel"],
        [
            {"ID": n, "SkillID": s, "ClassMask": m, "MinLevel": lvl}
            for n, (s, m, lvl) in enumerate(RACE_CLASS_INFO)
        ],
    )
    return tmp_path


def _skills(doc: dict, slug: str, kind: str) -> dict[str, int]:
    return {e["skill"]: e["level"] for e in doc["classes"][slug][kind]}


def test_plate_is_warrior_and_paladin_from_forty(raw: Path) -> None:
    doc = build_equip_skills("1.60.1.1", raw)
    assert _skills(doc, "warrior", "armor")["Plate Mail"] == 40
    assert _skills(doc, "paladin", "armor")["Plate Mail"] == 40
    assert "Plate Mail" not in _skills(doc, "hunter", "armor")


def test_mail_is_hunter_and_shaman_from_forty_and_warrior_from_the_start(raw: Path) -> None:
    doc = build_equip_skills("1.60.1.1", raw)
    assert _skills(doc, "hunter", "armor")["Mail"] == 40
    assert _skills(doc, "shaman", "armor")["Mail"] == 40
    assert _skills(doc, "warrior", "armor")["Mail"] == 1
    assert "Mail" not in _skills(doc, "rogue", "armor")


def test_shamans_train_fists_and_never_dual_wield(raw: Path) -> None:
    doc = build_equip_skills("1.60.1.1", raw)
    assert "Fist Weapons" in _skills(doc, "shaman", "weapons")
    assert doc["classes"]["shaman"]["dual_wield_level"] is None
    assert doc["classes"]["rogue"]["dual_wield_level"] == 10
    assert doc["classes"]["warrior"]["dual_wield_level"] == 20


def test_the_untaught_client_bits_are_dropped_and_listed(raw: Path) -> None:
    doc = build_equip_skills("1.60.1.1", raw)
    assert "Axes" not in _skills(doc, "rogue", "weapons")
    assert "Polearms" not in _skills(doc, "druid", "weapons")
    assert _skills(doc, "hunter", "weapons")["Polearms"] == 20
    assert {(d["class"], d["skill"]) for d in doc["dropped"]} == {
        ("rogue", "Axes"),
        ("druid", "Polearms"),
    }


def test_a_renumbered_skill_line_fails_loudly(raw: Path) -> None:
    rows = [{"DisplayName_lang": n, "ID": i, "CategoryID": c} for i, n, c in SKILL_LINES]
    rows[0]["DisplayName_lang"] = "Not Swords"
    _write(raw / "SkillLine.csv", ["DisplayName_lang", "ID", "CategoryID"], rows)
    with pytest.raises(EquipSkillsSchemaError):
        build_equip_skills("1.60.1.1", raw)


def test_write_names_its_sources(raw: Path, tmp_path: Path) -> None:
    root = tmp_path / "builds"
    (root / "1.60.1.1").mkdir(parents=True)
    out = write_equip_skills("1.60.1.1", root, raw)
    doc = json.loads(out.read_text())
    assert doc["source"][0] == "https://wago.tools/db2/SkillLine/csv?build=1.60.1.1"


def test_write_needs_the_tables(tmp_path: Path) -> None:
    (tmp_path / "builds" / "1.60.1.1" / "raw").mkdir(parents=True)
    with pytest.raises(SystemExit):
        write_equip_skills("1.60.1.1", tmp_path / "builds")


def test_the_committed_live_file_holds_the_classic_facts() -> None:
    """The committed table against the Classic facts the loot spec lists (§3.2 rule 3)."""
    doc = json.loads(Path("builds/1.60.1.70790/equip-skills.json").read_text())

    def weapons(slug: str) -> set[str]:
        return set(_skills(doc, slug, "weapons"))

    assert "Axes" not in weapons("rogue")
    assert not {"Two-Handed Axes", "Two-Handed Maces", "Two-Handed Swords", "Polearms",
                "Staves"} & weapons("rogue")  # fmt: skip
    assert "Swords" not in weapons("priest")
    assert "Maces" not in weapons("hunter")
    assert not {"Daggers", "Fist Weapons", "Staves"} & weapons("paladin")
    assert "Fist Weapons" in weapons("shaman")
    assert "Maces" in weapons("rogue")
