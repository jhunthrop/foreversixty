import csv
import json
from pathlib import Path

import pytest

from pipeline.manifest import newest_build
from pipeline.trainables import TRAINABLES, build_trainables, is_excluded_name

#: A floor, never an exact count: the 1.60.1.70009 class with the fewest active
#: trainables (warrior) has 51.
MIN_ACTIVE_TRAINABLES_PER_CLASS = 30
PLAYER_CLASSES = {
    "druid", "hunter", "mage", "paladin", "priest", "rogue", "shaman", "warlock", "warrior",
}  # fmt: skip

FROSTBOLT_1, FROSTBOLT_2, FIREBALL_NPC, FROSTFIRE_1, FROSTFIRE_2 = 116, 205, 9001, 401502, 1237312
PASSIVE, DNT_SPELL, TEST_SPELL, UA, PET_GROWL = 9100, 9101, 9102, 427717, 9200
MAGE_FAMILY, WARLOCK_FAMILY, HUNTER_FAMILY = 3, 5, 9
FIRE_LINE, FROST_LINE, PET_LINE, WEAPON_LINE = 8, 6, 7, 5


def _write(path: Path, header: list[str], rows: list[list[object]]) -> None:
    with path.open("w", newline="", encoding="utf-8") as handle:
        writer = csv.writer(handle)
        writer.writerow(header)
        writer.writerows(rows)


@pytest.fixture
def raw(tmp_path: Path) -> Path:
    spells = {
        # id: (name, subtext, level, cast_ms, cost, cooldown_ms)
        FROSTBOLT_1: ("Frostbolt", "Rank 1", 4, 2500, 25, 0),
        FROSTBOLT_2: ("Frostbolt", "Rank 2", 8, 3000, 35, 0),
        FROSTFIRE_1: ("Frostfire Bolt", "Rank 1", 40, 3000, 100, 0),
        FROSTFIRE_2: ("Frostfire Bolt", "Rank 2", 50, 3000, 120, 0),
        PASSIVE: ("Staves", "", 1, 0, 0, 0),
        DNT_SPELL: ("(DNT) Frostbolt", "Rank 1", 9, 3000, 5, 0),
        TEST_SPELL: ("Test Frostbolt", "Rank 1", 9, 3000, 5, 0),
        UA: ("Unstable Affliction", "Rank 1", 40, 1500, 200, 0),
        PET_GROWL: ("Growl", "Rank 1", 1, 0, 15, 5000),
        FIREBALL_NPC: ("Fireball", "", 1, 0, 0, 0),
    }
    _write(tmp_path / "SpellName.csv", ["ID", "Name_lang"], [[i, s[0]] for i, s in spells.items()])
    _write(
        tmp_path / "Spell.csv",
        ["ID", "NameSubtext_lang"],
        [[i, s[1]] for i, s in spells.items()],
    )
    _write(
        tmp_path / "SpellMisc.csv",
        ["ID", "DifficultyID", "SpellID", "CastingTimeIndex"],
        [[i, 0, i, i] for i in spells],
    )
    _write(tmp_path / "SpellCastTimes.csv", ["ID", "Base"], [[i, s[3]] for i, s in spells.items()])
    _write(
        tmp_path / "SpellLevels.csv",
        ["ID", "DifficultyID", "SpellID", "SpellLevel"],
        [[i, 0, i, s[2]] for i, s in spells.items()],
    )
    _write(
        tmp_path / "SpellPower.csv",
        ["ID", "SpellID", "ManaCost", "PowerType", "PowerCostPct"],
        [[i, i, s[4], 0, 0] for i, s in spells.items()],
    )
    _write(
        tmp_path / "SpellCooldowns.csv",
        ["ID", "DifficultyID", "SpellID", "RecoveryTime", "CategoryRecoveryTime"],
        [[i, 0, i, s[5], 0] for i, s in spells.items()],
    )
    families = {
        FROSTBOLT_1: MAGE_FAMILY, FROSTBOLT_2: MAGE_FAMILY, FROSTFIRE_1: MAGE_FAMILY,
        FROSTFIRE_2: MAGE_FAMILY, DNT_SPELL: MAGE_FAMILY, UA: WARLOCK_FAMILY,
        PET_GROWL: HUNTER_FAMILY,
    }  # fmt: skip
    _write(
        tmp_path / "SpellClassOptions.csv",
        ["ID", "SpellID", "SpellClassSet"],
        [[n, i, f] for n, (i, f) in enumerate(families.items())],
    )
    _write(
        tmp_path / "SkillLine.csv",
        ["ID", "DisplayName_lang", "CategoryID"],
        [
            [FIRE_LINE, "Fire", 7], [FROST_LINE, "Frost", 7],
            [PET_LINE, "Pet - Cat", 7], [WEAPON_LINE, "Staves", 6],
        ],
    )  # fmt: skip
    mage_mask = 128
    _write(
        tmp_path / "SkillLineAbility.csv",
        ["ID", "SkillLine", "Spell", "ClassMask", "SupercedesSpell", "AcquireMethod"],
        [
            [1, FROST_LINE, FROSTBOLT_1, mage_mask, 0, 0],
            [2, FROST_LINE, FROSTBOLT_2, mage_mask, FROSTBOLT_1, 0],
            [3, FIRE_LINE, FROSTFIRE_1, 0, 0, 0],  # no mask: the spell's own family names the class
            [4, FIRE_LINE, FROSTFIRE_2, mage_mask, 0, 0],  # unlinked rank: grouped by name
            [5, FROST_LINE, PASSIVE, mage_mask, 0, 0],
            [6, FROST_LINE, DNT_SPELL, mage_mask, 0, 0],
            [7, FROST_LINE, TEST_SPELL, mage_mask, 0, 0],
            [8, FROST_LINE, FIREBALL_NPC, 0, 0, 0],  # no mask, no family: not a player ability
            [9, PET_LINE, PET_GROWL, 0, 0, 0],
            [10, WEAPON_LINE, PASSIVE, mage_mask, 0, 0],  # a weapon skill line is not a class line
        ],
    )
    return tmp_path


def by_slug(raw: Path) -> dict:
    return {record.class_slug: record for record in build_trainables("9.9.9.9", raw)}


def names(record) -> set[str]:
    return {t.name for t in record.trainables}


def test_there_is_one_record_per_player_class_even_when_it_has_none(raw):
    records = by_slug(raw)
    assert set(records) == PLAYER_CLASSES
    assert records["paladin"].trainables == []
    assert records["mage"].build == "9.9.9.9"


def test_ranks_group_by_supersede_link_and_by_name(raw):
    mage = {t.name: t for t in by_slug(raw)["mage"].trainables}
    assert [(r.id, r.rank, r.level) for r in mage["Frostbolt"].ranks] == [
        (FROSTBOLT_1, 1, 4),
        (FROSTBOLT_2, 2, 8),
    ]
    assert [r.id for r in mage["Frostfire Bolt"].ranks] == [FROSTFIRE_1, FROSTFIRE_2]
    assert mage["Frostfire Bolt"].skill_line == "Fire"
    assert mage["Frostfire Bolt"].source == "skill_line_ability"


def test_active_means_a_cost_a_cast_time_or_a_cooldown(raw):
    mage = {t.name: t for t in by_slug(raw)["mage"].trainables}
    assert mage["Frostfire Bolt"].active is True
    assert mage["Frostfire Bolt"].cost == 120  # of the highest-level rank
    assert mage["Frostfire Bolt"].cast_time_ms == 3000
    assert mage["Staves"].active is False


def test_a_spell_without_a_skill_line_row_is_found_through_its_class_family(raw):
    warlock = {t.name: t for t in by_slug(raw)["warlock"].trainables}
    assert warlock["Unstable Affliction"].source == "class_spell"
    assert warlock["Unstable Affliction"].skill_line == ""
    assert warlock["Unstable Affliction"].active is True


def test_rune_rows_are_not_trainable(raw):
    """AcquireMethod 3 and learn level 0 rows are Season of Discovery runes."""
    rows = list(csv.DictReader((raw / "SkillLineAbility.csv").open()))
    for row in rows:
        row["AcquireMethod"] = "3" if row["Spell"] == str(FROSTFIRE_1) else "0"
    with (raw / "SkillLineAbility.csv").open("w", newline="") as handle:
        writer = csv.DictWriter(handle, fieldnames=list(rows[0]))
        writer.writeheader()
        writer.writerows(rows)
    mage = {t.name: t for t in by_slug(raw)["mage"].trainables}
    assert [r.id for r in mage["Frostfire Bolt"].ranks] == [FROSTFIRE_2]


def test_a_spell_with_a_rune_row_is_not_resurrected_by_the_family_fallback(raw):
    rows = list(csv.DictReader((raw / "SkillLineAbility.csv").open()))
    for row in rows:
        row["AcquireMethod"] = "3"
    with (raw / "SkillLineAbility.csv").open("w", newline="") as handle:
        writer = csv.DictWriter(handle, fieldnames=list(rows[0]))
        writer.writeheader()
        writer.writerows(rows)
    assert "Frostfire Bolt" not in names(by_slug(raw)["mage"])


def test_test_dnt_pet_and_unattributed_spells_are_dropped(raw):
    records = by_slug(raw)
    everything = set().union(*(names(r) for r in records.values()))
    assert "(DNT) Frostbolt" not in everything
    assert "Test Frostbolt" not in everything
    assert "Fireball" not in everything
    assert "Growl" not in names(records["hunter"])


@pytest.mark.parametrize(
    ("name", "excluded"),
    [
        ("Frostbolt", False),
        ("Testimony", False),
        ("(DNT) Frostbolt", True),
        ("[DNT] Hydrate Seed Pod", True),
        ("Copy of Frostbolt", True),
        ("Runecarving Test - Crippling Poison", True),
        ("", True),
    ],
)
def test_is_excluded_name(name, excluded):
    assert is_excluded_name(name) is excluded


# --- contracts on the committed build -------------------------------------------------

COMMITTED = Path("builds") / newest_build() / TRAINABLES


def committed() -> dict[str, dict]:
    if not COMMITTED.exists():
        pytest.skip("trainables have not been generated for the newest build")
    return {p.stem: json.loads(p.read_text(encoding="utf-8")) for p in COMMITTED.glob("*.json")}


def active_names(record: dict) -> set[str]:
    return {t["name"] for t in record["trainables"] if t["active"]}


def test_every_class_has_a_floor_of_active_trainables():
    records = committed()
    assert set(records) == PLAYER_CLASSES
    for slug, record in records.items():
        assert len(active_names(record)) >= MIN_ACTIVE_TRAINABLES_PER_CLASS, slug


def learnable_ids(record: dict, name: str) -> set[int]:
    return {
        r["id"]
        for t in record["trainables"]
        if t["name"] == name and t["source"] == "skill_line_ability"
        for r in t["ranks"]
    }


@pytest.mark.parametrize(
    ("slug", "name", "ids"),
    [
        ("mage", "Frostfire Bolt", {401502, 1237312, 1237313}),
        ("hunter", "Lacerate", {1299332}),
        ("priest", "Dark Sacrifice", {1277324}),
        ("warlock", "Curse of the Elements", {440892, 1311676, 1311677, 1311680}),
    ],
)
def test_learnable_ranks_the_owner_named_are_present(slug, name, ids):
    assert ids <= learnable_ids(committed()[slug], name)


@pytest.mark.parametrize(
    ("slug", "name"),
    [("shaman", "Lava Lash"), ("priest", "Dispersion"), ("warrior", "Raging Blow"),
     ("paladin", "Divine Storm"), ("rogue", "Saber Slash")],
)  # fmt: skip
def test_runes_are_not_listed(slug, name):
    assert name not in {t["name"] for t in committed()[slug]["trainables"]}


@pytest.mark.parametrize(
    ("slug", "name"),
    [
        ("mage", "Frostfire Bolt"),
        ("warlock", "Unstable Affliction"),
        ("warlock", "Curse of the Elements"),
        ("hunter", "Hydra Shot"),
    ],
)
def test_the_named_new_forever_abilities_are_trainable(slug, name):
    assert name in active_names(committed()[slug])


def test_no_development_spell_is_trainable():
    for slug, record in committed().items():
        for trainable in record["trainables"]:
            assert not is_excluded_name(trainable["name"]), (slug, trainable["name"])
