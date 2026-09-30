import csv

import pytest

from pipeline import proficiency
from pipeline.proficiency import ARMOR, WEAPON, can_equip


@pytest.fixture(autouse=True)
def _reset_weapon_proficiency():
    """Every test in this module runs against the hardcoded fallback table
    unless it explicitly calls `proficiency.configure`; undo that afterwards
    so one test's client-derived table can never leak into the next."""
    yield
    proficiency.configure(None)


def test_plate_is_warrior_and_paladin_only():
    assert can_equip(1, ARMOR, 4)
    assert can_equip(2, ARMOR, 4)
    assert not can_equip(8, ARMOR, 4)


def test_every_class_can_wear_cloth_and_miscellaneous_armour():
    for class_id in (1, 2, 3, 4, 5, 7, 8, 9, 11):
        assert can_equip(class_id, ARMOR, 0)
        assert can_equip(class_id, ARMOR, 1)


def test_shields_exclude_the_classes_that_cannot_use_them():
    assert can_equip(1, ARMOR, 6)
    assert can_equip(7, ARMOR, 6)
    assert not can_equip(4, ARMOR, 6)


def test_wands_are_for_the_three_wand_classes():
    assert can_equip(5, WEAPON, 19)
    assert can_equip(8, WEAPON, 19)
    assert can_equip(9, WEAPON, 19)
    assert not can_equip(1, WEAPON, 19)


def test_two_handed_swords_exclude_rogues():
    assert can_equip(1, WEAPON, 8)
    assert not can_equip(4, WEAPON, 8)


def test_relics_are_one_class_each_and_nobody_else():
    # libram (7): paladin only
    assert can_equip(2, ARMOR, 7)
    for class_id in (1, 3, 4, 5, 7, 8, 9, 11):
        assert not can_equip(class_id, ARMOR, 7)
    # totem (9): shaman only
    assert can_equip(7, ARMOR, 9)
    for class_id in (1, 2, 3, 4, 5, 8, 9, 11):
        assert not can_equip(class_id, ARMOR, 9)
    # idol (8): druid only
    assert can_equip(11, ARMOR, 8)
    for class_id in (1, 2, 3, 4, 5, 7, 8, 9):
        assert not can_equip(class_id, ARMOR, 8)


def test_monster_and_fishing_subclasses_belong_to_nobody():
    for class_id in (1, 4, 8):
        assert not can_equip(class_id, WEAPON, 14)
        assert not can_equip(class_id, WEAPON, 17)
        assert not can_equip(class_id, WEAPON, 20)


def test_non_equipment_item_classes_are_never_equippable():
    assert not can_equip(1, 0, 0)
    assert not can_equip(1, 7, 5)


# --- Contract tests from the 2026-09-30 wow-player sweep (day3 data-followups-5) ---


def test_paladin_has_no_axe():
    assert not can_equip(2, WEAPON, 0)  # one-hand axe
    assert not can_equip(2, WEAPON, 1)  # two-hand axe


def test_druid_has_no_polearm():
    assert not can_equip(11, WEAPON, 6)


def test_shaman_has_two_hand_axe_and_two_hand_mace():
    assert can_equip(7, WEAPON, 1)
    assert can_equip(7, WEAPON, 5)


def test_rogue_has_one_hand_mace():
    assert can_equip(4, WEAPON, 4)


def test_warrior_has_every_melee_subclass():
    melee = {0, 1, 4, 5, 6, 7, 8, 10, 13, 15}
    for subclass_id in melee:
        assert can_equip(1, WEAPON, subclass_id), subclass_id


def test_wands_are_priest_mage_warlock_only():
    for class_id in (5, 8, 9):
        assert can_equip(class_id, WEAPON, 19)
    for class_id in (1, 2, 3, 4, 7, 11):
        assert not can_equip(class_id, WEAPON, 19)


# --- configure(): the client-derived table and its fallback ---


def _write_skill_line(path, rows):
    fieldnames = ["ID", "DisplayName_lang", "CategoryID"]
    with path.open("w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(rows)


def _write_skill_line_ability(path, rows):
    fieldnames = ["SkillLine", "ClassMask"]
    with path.open("w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(rows)


def _minimal_client_tables(tmp_path, *, class_mask_overrides=None):
    """A synthetic `raw/` with just enough of `SkillLine.csv` +
    `SkillLineAbility.csv` for `_derive_weapon_subclasses_from_client` to
    build a full table -- one row per real weapon skill line/name pair, and
    a ClassMask that reproduces the true 1.x table plus Forever's two
    documented changes (shaman axes/maces, rogue maces), so the default
    fixture matches `WEAPON_SUBCLASSES` exactly unless a test overrides one
    skill line's mask."""
    masks = {
        44: 1 + 4 + 8 + 64,  # Axes: warrior, hunter, rogue, shaman
        172: 1 + 4 + 64,  # Two-Handed Axes: warrior, hunter, shaman
        45: 1 + 4 + 8,  # Bows
        46: 1 + 4 + 8,  # Guns
        54: 1 + 2 + 8 + 16 + 64 + 1024,  # Maces: + paladin, rogue, priest, shaman, druid
        160: 1 + 2 + 64 + 1024,  # Two-Handed Maces
        229: 1 + 2 + 4,  # Polearms: warrior, paladin, hunter
        43: 1 + 2 + 4 + 8 + 128 + 256,  # Swords
        55: 1 + 2 + 4,  # Two-Handed Swords
        136: 1 + 4 + 16 + 64 + 128 + 256 + 1024,  # Staves (warrior + hunter get it too)
        473: 1 + 4 + 8 + 64 + 1024,  # Fist Weapons
        173: 1 + 4 + 8 + 16 + 64 + 128 + 256 + 1024,  # Daggers
        176: 1 + 4 + 8,  # Thrown
        226: 1 + 4 + 8,  # Crossbows
        228: 16 + 128 + 256,  # Wands
    }
    if class_mask_overrides:
        masks.update(class_mask_overrides)
    raw = tmp_path / "raw"
    raw.mkdir()
    _write_skill_line(
        raw / "SkillLine.csv",
        [
            {"ID": skill_id, "DisplayName_lang": name, "CategoryID": 6}
            for skill_id, name in proficiency._WEAPON_SKILL_LINE_NAME.items()
        ],
    )
    _write_skill_line_ability(
        raw / "SkillLineAbility.csv",
        [{"SkillLine": skill_id, "ClassMask": mask} for skill_id, mask in masks.items()],
    )
    return raw


def test_configure_derives_from_client_tables(tmp_path):
    raw = _minimal_client_tables(tmp_path)
    source = proficiency.configure(raw)
    assert source.startswith("client:")
    for class_id, expected in proficiency.WEAPON_SUBCLASSES.items():
        for subclass_id in range(20):
            assert can_equip(class_id, WEAPON, subclass_id) == (subclass_id in expected)


def test_configure_still_excludes_the_vestigial_bits_even_if_the_client_carries_them(tmp_path):
    raw = _minimal_client_tables(
        tmp_path,
        class_mask_overrides={
            44: 1 + 2 + 4 + 8 + 64,  # Axes: client carries paladin (2) and rogue-axe again
            229: 1 + 2 + 4 + 1024,  # Polearms: client carries druid (1024) again
        },
    )
    proficiency.configure(raw)
    assert not can_equip(2, WEAPON, 0)  # paladin, one-hand axe
    assert not can_equip(4, WEAPON, 0)  # rogue, one-hand axe
    assert not can_equip(11, WEAPON, 6)  # druid, polearm


def test_configure_falls_back_when_the_client_tables_are_absent(tmp_path):
    raw = tmp_path / "raw"
    raw.mkdir()
    source = proficiency.configure(raw)
    assert source == "fallback:hardcoded WEAPON_SUBCLASSES"
    assert not can_equip(2, WEAPON, 0)


def test_configure_falls_back_when_a_weapon_skill_line_is_renumbered(tmp_path):
    raw = _minimal_client_tables(tmp_path)
    rows = list(csv.DictReader((raw / "SkillLine.csv").open(encoding="utf-8")))
    for row in rows:
        if row["ID"] == "44":
            row["DisplayName_lang"] = "Something Else"
    _write_skill_line(raw / "SkillLine.csv", rows)
    source = proficiency.configure(raw)
    assert source == "fallback:hardcoded WEAPON_SUBCLASSES"
