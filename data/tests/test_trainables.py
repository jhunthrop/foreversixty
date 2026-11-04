import csv
import json
from pathlib import Path

import pytest

from pipeline.manifest import newest_build
from pipeline.trainables import (
    SPELLS_FILE,
    TRAINABLES,
    PreviousBuild,
    build_trainables,
    is_excluded_name,
    load_previous_build,
)

#: A floor, never an exact count: the 1.60.1.70009 class with the fewest active
#: trainables (warrior) has 51.
MIN_ACTIVE_TRAINABLES_PER_CLASS = 30
PLAYER_CLASSES = {
    "druid", "hunter", "mage", "paladin", "priest", "rogue", "shaman", "warlock", "warrior",
}  # fmt: skip

FROSTBOLT_1, FROSTBOLT_2, FIREBALL_NPC, FROSTFIRE_1, FROSTFIRE_2 = 116, 205, 9001, 401502, 1237312
PASSIVE, DNT_SPELL, TEST_SPELL, UA, PET_GROWL = 9100, 9101, 9102, 427717, 9200
COLD_SNAP, IGNITE, FROSTBOLT_VARIANT = 9300, 9301, 9302
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
        COLD_SNAP: ("Cold Snap", "", 0, 0, 0, 480_000),  # a talent: no learn row, no rank
        IGNITE: ("Ignite", "", 0, 0, 0, 0),  # a passive talent
        # a class-family Frostbolt with a rank and a level but no learn row (an NPC or proc copy)
        FROSTBOLT_VARIANT: ("Frostbolt", "Rank 2", 8, 3500, 35, 0),
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
        PET_GROWL: HUNTER_FAMILY, COLD_SNAP: MAGE_FAMILY, IGNITE: MAGE_FAMILY,
        FROSTBOLT_VARIANT: MAGE_FAMILY,
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
        ("Rogue Passive (DND)", True),
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


def _write_live_tree(talents_dir: Path, talents: list[dict], tree: str = "Fire") -> Path:
    """A live tree of `talents`, each {"name"} plus optional "spell_id" (its every rank)."""
    talents_dir.mkdir(parents=True, exist_ok=True)
    rows = [
        {
            "name": t["name"],
            "spell_id": t.get("spell_id", 0),
            "ranks": [{"spell_id": t.get("spell_id", 0)}],
        }
        for t in talents
    ]
    body = {"trees": [{"name": tree, "position": 0, "talents": rows}]}
    (talents_dir / "mage.json").write_text(json.dumps(body), encoding="utf-8")
    return talents_dir


def mage_with_tree(raw: Path, talents_dir: Path) -> dict:
    records = {r.class_slug: r for r in build_trainables("9.9.9.9", raw, talents_dir)}
    return {t.name: t for t in records["mage"].trainables}


def test_a_talent_moved_off_the_live_tree_keeps_its_learn_row_as_a_baseline_spell(raw, tmp_path):
    """The 70707 push repointed Voice of Truth's and Binding Heal's trait nodes (Oath of Light,
    Holy Binding) and Wowhead's live spellbook lists the old spells at their trainer levels: a
    spell with a learn row that no live talent teaches is a trainer spell, whatever an older
    trait row taught."""
    _write(raw / "TraitDefinition.csv", ["ID", "SpellID"], [[1, FROSTFIRE_1]])
    mage = mage_with_tree(raw, _write_live_tree(tmp_path / "talents", [{"name": "Ignite"}]))
    assert mage["Frostfire Bolt"].source == "skill_line_ability"
    assert "Frostbolt" in mage


def test_an_ability_a_live_talent_teaches_is_a_talent_not_a_trainer_spell(raw, tmp_path):
    mage = mage_with_tree(
        raw,
        _write_live_tree(
            tmp_path / "talents", [{"name": "Frostfire Bolt", "spell_id": FROSTFIRE_1}]
        ),
    )
    assert mage["Frostfire Bolt"].source == "talent"
    # its trainer-taught higher rank stays with the ability, under the one tag
    assert [r.id for r in mage["Frostfire Bolt"].ranks] == [FROSTFIRE_1, FROSTFIRE_2]
    assert mage["Frostbolt"].source == "skill_line_ability"


def test_a_talent_is_matched_by_its_spell_when_the_names_differ(raw, tmp_path):
    talents = _write_live_tree(
        tmp_path / "talents", [{"name": "Frozen Core", "spell_id": FROSTFIRE_2}]
    )
    assert mage_with_tree(raw, talents)["Frostfire Bolt"].source == "talent"


def test_an_active_live_talent_without_a_learn_row_is_listed_as_a_talent(raw, tmp_path):
    """Cold Snap, Adrenaline Rush, Shadowform: no SkillLineAbility row and no rank subtext, so
    neither source finds them; the live tree teaches them, so they are listed, as talents."""
    talents = _write_live_tree(
        tmp_path / "talents",
        [{"name": "Cold Snap", "spell_id": COLD_SNAP}, {"name": "Ignite", "spell_id": IGNITE}],
        tree="Frost",
    )
    mage = mage_with_tree(raw, talents)
    cold_snap = mage["Cold Snap"]
    assert (cold_snap.source, cold_snap.skill_line, cold_snap.active) == ("talent", "Frost", True)
    assert cold_snap.cooldown_ms == 480_000
    assert [(r.id, r.rank, r.level) for r in cold_snap.ranks] == [(COLD_SNAP, 0, 0)]
    assert "Ignite" not in mage  # a passive talent stays in the talent file alone


def test_a_spell_on_two_learn_rows_is_one_rank(raw):
    rows = list(csv.DictReader((raw / "SkillLineAbility.csv").open()))
    rows.append({**rows[1], "ID": "99"})  # Frostbolt rank 2 on a second row of the same line
    with (raw / "SkillLineAbility.csv").open("w", newline="") as handle:
        writer = csv.DictWriter(handle, fieldnames=list(rows[0]))
        writer.writeheader()
        writer.writerows(rows)
    mage = {t.name: t for t in by_slug(raw)["mage"].trainables}
    assert [r.id for r in mage["Frostbolt"].ranks] == [FROSTBOLT_1, FROSTBOLT_2]


def test_without_live_trees_nothing_is_tagged_a_talent(raw, tmp_path):
    for record in by_slug(raw).values():
        assert all(t.source != "talent" for t in record.trainables)
    mage = mage_with_tree(raw, tmp_path / "no-talents")
    assert mage["Frostfire Bolt"].source == "skill_line_ability"


@pytest.mark.parametrize(
    ("slug", "name"),
    [
        # Era talents Forever's trees never teach; trainer spells at their client levels
        # (Blizzard recap, Wowhead's live spellbook, beta casts; round 12 trainables report)
        ("hunter", "Aimed Shot"), ("paladin", "Blessing of Kings"), ("paladin", "Consecration"),
        ("priest", "Divine Spirit"), ("druid", "Omen of Clarity"), ("priest", "Lightwell"),
        # a talent whose node the 70707 push repointed; a trainer spell since (Binding Heal
        # went the same way until 70755's push removed its learn rows for good: see
        # test_abilities_the_client_removed_are_not_trainable)
        ("paladin", "Voice of Truth"),
    ],
)  # fmt: skip
def test_the_baseline_spells_round_12_settled_are_trainer_spells(slug, name):
    sources = {t["source"] for t in committed()[slug]["trainables"] if t["name"] == name}
    assert sources == {"skill_line_ability"}


def test_a_class_spell_named_like_a_learn_row_ability_is_not_a_rank_of_it(raw):
    """Aimed Shot 27632 (rank 6, 3 s) beside the trainer's 20904 (2 s), Primal Bite 1238073
    (rank 4 at 60, its learn row removed by push 113283): the learn rows are the ability's
    ranks, and a no-learn-row copy is neither a rank nor the source of its figures."""
    frostbolt = {t.name: t for t in by_slug(raw)["mage"].trainables}["Frostbolt"]
    assert [r.id for r in frostbolt.ranks] == [FROSTBOLT_1, FROSTBOLT_2]
    assert frostbolt.cast_time_ms == 3000


def test_a_learn_row_a_hotfix_push_removed_is_not_resurrected_as_a_class_spell(raw):
    """A push (status 2) drops the row from raw/ but not from raw/shipped/: the spell had a
    learn row, so the no-learn-row fallback must not list it."""
    shipped = raw / "shipped"
    shipped.mkdir()
    _write(
        shipped / "SkillLineAbility.csv",
        ["ID", "SkillLine", "Spell", "ClassMask", "SupercedesSpell", "AcquireMethod"],
        [[11, FIRE_LINE, UA, 256, 0, 0]],
    )
    assert "Unstable Affliction" not in names(by_slug(raw)["warlock"])


def _add_spell(
    raw: Path, spell_id: int, name: str, subtext: str, level: int, cost: int, family: int
) -> None:
    """Append one spell (no cast time, no cooldown) to every table the fixture writes."""
    rows = {
        "SpellName.csv": [spell_id, name],
        "Spell.csv": [spell_id, subtext],
        "SpellMisc.csv": [spell_id, 0, spell_id, spell_id],
        "SpellCastTimes.csv": [spell_id, 0],
        "SpellLevels.csv": [spell_id, 0, spell_id, level],
        "SpellPower.csv": [spell_id, spell_id, cost, 0, 0],
        "SpellCooldowns.csv": [spell_id, 0, spell_id, 0, 0],
        "SpellClassOptions.csv": [10_000 + spell_id, spell_id, family],
    }
    for table, row in rows.items():
        with (raw / table).open("a", newline="", encoding="utf-8") as handle:
            csv.writer(handle).writerow(row)


def _add_learn_rows(raw: Path, rows: list[list[object]]) -> None:
    """Append SkillLineAbility rows: [ID, SkillLine, Spell, ClassMask, SupercedesSpell,
    AcquireMethod]."""
    with (raw / "SkillLineAbility.csv").open("a", newline="", encoding="utf-8") as handle:
        csv.writer(handle).writerows(rows)


FROSTBOLT_3, FROST_ARMOR, DEMON_CHARGE, FROSTFIRE_3 = 9400, 9401, 9402, 9403
ICE_SHIELD, ICE_SHIELD_PROC = 9404, 9405


def test_an_auto_learned_row_is_a_rank(raw):
    """AcquireMethod 2 rows are learned with the skill line, not bought: every class's
    starting rank 1s (Heroic Strike 78, Fireball 133, Smite 585), Battle Stance 2457,
    Auto Shot 75, and Sniper Shot rank 3 1310786 at 58. Wowhead's live spellbook lists
    each at its client level."""
    _add_spell(raw, FROSTBOLT_3, "Frostbolt", "Rank 3", 14, 50, MAGE_FAMILY)
    _add_spell(raw, FROST_ARMOR, "Frost Armor", "Rank 1", 1, 60, MAGE_FAMILY)
    _add_learn_rows(
        raw,
        [
            [20, FROST_LINE, FROSTBOLT_3, 128, FROSTBOLT_2, 2],
            [21, FROST_LINE, FROST_ARMOR, 128, 0, 2],
        ],
    )
    mage = {t.name: t for t in by_slug(raw)["mage"].trainables}
    assert [r.id for r in mage["Frostbolt"].ranks] == [FROSTBOLT_1, FROSTBOLT_2, FROSTBOLT_3]
    assert mage["Frostbolt"].source == "skill_line_ability"
    assert [(r.id, r.level) for r in mage["Frost Armor"].ranks] == [(FROST_ARMOR, 1)]


def test_an_unmasked_auto_learned_row_counts_only_as_a_rank_of_a_learn_row_ability(raw):
    """Sniper Shot rank 3 (mask 0, the family names the class) joins its trainer ranks;
    an unmasked auto-learned spell of no learnable ability (Metamorphosis : Demon Charge
    412783, a Season of Discovery rune effect) is not a trainable."""
    _add_spell(raw, FROSTFIRE_3, "Frostfire Bolt", "Rank 3", 60, 140, MAGE_FAMILY)
    _add_spell(raw, DEMON_CHARGE, "Demon Charge", "", 1, 10, WARLOCK_FAMILY)
    _add_learn_rows(
        raw,
        [
            [22, FIRE_LINE, FROSTFIRE_3, 0, 0, 2],
            [23, FIRE_LINE, DEMON_CHARGE, 0, 0, 2],
        ],
    )
    records = by_slug(raw)
    mage = {t.name: t for t in records["mage"].trainables}
    assert [r.id for r in mage["Frostfire Bolt"].ranks] == [FROSTFIRE_1, FROSTFIRE_2, FROSTFIRE_3]
    assert "Demon Charge" not in names(records["warlock"])


def test_an_inactive_learn_row_spell_on_an_active_rank_is_its_effect_not_a_rank(raw):
    """Lightning Shield 324 and its damage spell 26364 both carry a rank-1 learn row at 8
    (so do Shadowguard 18137/28377, Explosive Trap 13813/13812, Tame Beast 1515/13481).
    The effect is not a second rank 1, and the figures are the castable spell's."""
    _add_spell(raw, ICE_SHIELD, "Ice Shield", "Rank 1", 8, 45, MAGE_FAMILY)
    _add_spell(raw, ICE_SHIELD_PROC, "Ice Shield", "Rank 1", 8, 0, MAGE_FAMILY)
    _add_learn_rows(
        raw,
        [
            [24, FROST_LINE, ICE_SHIELD, 128, 0, 0],
            [25, FROST_LINE, ICE_SHIELD_PROC, 0, 0, 0],
        ],
    )
    shield = {t.name: t for t in by_slug(raw)["mage"].trainables}["Ice Shield"]
    assert [r.id for r in shield.ranks] == [ICE_SHIELD]
    assert shield.cost == 45


# --- removed learn rows carried from the previous build ---------------------------------

LAVA_BURST, BINDING_HEAL = 408490, 401937


def test_a_no_learn_row_spell_the_previous_build_knew_and_did_not_list_is_a_removed_learn_row(
    raw,
):
    """70790 shipped without the Lava Burst learn rows 70755's push had removed: the previous
    build knew the spell (spells.json) and did not list it without a learn row, so the
    fallback must not resurrect it."""
    previous = PreviousBuild("1.60.1.70755", known=frozenset({UA}), no_learn_row=frozenset())
    records = {r.class_slug: r for r in build_trainables("9.9.9.9", raw, previous=previous)}
    assert "Unstable Affliction" not in names(records["warlock"])
    assert "Frostfire Bolt" in names(records["mage"])  # learn rows are never filtered


def test_a_no_learn_row_spell_the_previous_build_listed_that_way_stays(raw):
    previous = PreviousBuild("1.60.1.70755", known=frozenset({UA}), no_learn_row=frozenset({UA}))
    records = {r.class_slug: r for r in build_trainables("9.9.9.9", raw, previous=previous)}
    assert "Unstable Affliction" in names(records["warlock"])


def test_a_no_learn_row_spell_the_previous_build_never_knew_is_new_and_listed(raw):
    previous = PreviousBuild(
        "1.60.1.70755", known=frozenset({FROSTBOLT_1}), no_learn_row=frozenset()
    )
    records = {r.class_slug: r for r in build_trainables("9.9.9.9", raw, previous=previous)}
    assert "Unstable Affliction" in names(records["warlock"])


def _write_build(
    root: Path,
    build: str,
    known: list[int],
    class_spells: list[int],
    trainer_spells: list[int] = (),
    family_talents: list[int] = (),
) -> None:
    """A committed build: its spell list and one warlock trainables file with class spells
    (no learn row), trainer spells (a skill line) and talents the live tree re-tagged from
    class spells (talent source, the empty family skill line)."""
    build_dir = root / build
    (build_dir / TRAINABLES).mkdir(parents=True)
    (build_dir / SPELLS_FILE).write_text(
        json.dumps([{"id": i, "name": f"Spell {i}"} for i in known]), encoding="utf-8"
    )

    def entry(i: int, source: str, skill_line: str) -> dict:
        return {
            "name": f"Spell {i}",
            "skill_line": skill_line,
            "source": source,
            "active": True,
            "cost": 1,
            "cost_type": 0,
            "cast_time_ms": 0,
            "cooldown_ms": 0,
            "ranks": [{"id": i, "rank": 1, "level": 10}],
        }

    record = {
        "build": build,
        "class_slug": "warlock",
        "trainables": [
            *(entry(i, "class_spell", "") for i in class_spells),
            *(entry(i, "skill_line_ability", "Affliction") for i in trainer_spells),
            *(entry(i, "talent", "") for i in family_talents),
        ],
    }
    (build_dir / TRAINABLES / "warlock.json").write_text(json.dumps(record), encoding="utf-8")


def test_load_previous_build_reads_the_newest_older_committed_build(tmp_path):
    _write_build(tmp_path, "1.60.1.70700", known=[1, 2], class_spells=[1])
    _write_build(
        tmp_path, "1.60.1.70755", known=[2, 3, 5, 6], class_spells=[3],
        trainer_spells=[5], family_talents=[6],
    )  # fmt: skip
    _write_build(tmp_path, "1.60.1.70800", known=[4], class_spells=[4])  # newer: ignored
    (tmp_path / "notes").mkdir()  # not a build directory
    previous = load_previous_build(tmp_path, "1.60.1.70790")
    assert previous is not None
    assert previous.build == "1.60.1.70755"
    assert previous.known == {2, 3, 5, 6}
    assert previous.no_learn_row == {3, 6}
    assert previous.keeps_without_learn_row(3)  # listed without a learn row before
    assert previous.keeps_without_learn_row(6)  # a talent re-tagged from a class spell (Turn Evil)
    assert previous.keeps_without_learn_row(9)  # unknown to 70755: new
    assert not previous.keeps_without_learn_row(2)  # known, not listed: its learn rows are gone
    assert not previous.keeps_without_learn_row(5)  # had a trainer row: gone now means removed


def test_load_previous_build_is_none_for_the_first_build(tmp_path):
    _write_build(tmp_path, "1.60.1.70800", known=[4], class_spells=[4])
    assert load_previous_build(tmp_path, "1.60.1.70790") is None
    assert load_previous_build(tmp_path, "not-a-build") is None


@pytest.mark.parametrize(
    ("slug", "name", "spell_id"),
    [("shaman", "Lava Burst", LAVA_BURST), ("priest", "Binding Heal", BINDING_HEAL)],
)
def test_abilities_the_client_removed_are_not_trainable(slug, name, spell_id):
    """1.60.1.70755's push removed both abilities' learn rows and every later client ships
    without them (Lava Burst is gone from the Elemental talent texts too); neither may come
    back through the no-learn-row fallback."""
    record = committed()[slug]
    listed = [
        t
        for t in record["trainables"]
        if t["name"] == name or any(r["id"] == spell_id for r in t["ranks"])
    ]
    assert listed == [], [(t["name"], t["source"]) for t in listed]
