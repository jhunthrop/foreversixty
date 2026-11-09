"""Level-scoped panel lines (`panel_when` variants) and the rule that a panel line never
names a class spell the character cannot have at its band (pipeline.panelnames)."""

import json
from pathlib import Path

import pytest

from pipeline.addonrotation import AddonRotationError, build_rotation
from pipeline.models import PanelWhen, check_panel_when
from pipeline.panelnames import (
    ClassSpellNames,
    known_names,
    load_class_spell_names,
    named_spells,
    panel_name_violations,
    unavailable_names,
)
from pipeline.specs import active_build_dir

SPELLRANKS = {
    "classes": {
        "warrior": {
            "Bloodrage": [{"id": 2687, "level": 1, "rank": 0}],
            "Whirlwind": [{"id": 1680, "level": 36, "rank": 0}],
        }
    }
}

BLOODRAGE = {"notes": "N.", "action": {"castSpell": {"spellId": {"spellId": 2687}}}}


def _write_apl(curated_dir: Path, panel_when: dict) -> None:
    apl_dir = curated_dir / "apl"
    apl_dir.mkdir(parents=True, exist_ok=True)
    (apl_dir / "warrior-fury.json").write_text(
        json.dumps(
            {
                "spec": "warrior-fury",
                "rotation": {"priorityList": [BLOODRAGE]},
                "panel_when": panel_when,
            }
        ),
        encoding="utf-8",
    )


def _condition_at(curated_dir: Path, level: int) -> str:
    """The panel line a character of `level` reads (the band at or below it, Rotation.lua)."""
    bands = build_rotation("warrior-fury", curated_dir, SPELLRANKS)
    return [band for band in bands if band.level <= level][-1].lines[0].condition


def test_a_plain_line_is_one_variant_from_level_one():
    assert check_panel_when({"0": "on cooldown"}) == {0: (PanelWhen(1, "on cooldown"),)}


def test_each_band_reads_the_variant_for_its_level(tmp_path):
    _write_apl(
        tmp_path,
        {
            "0": [
                {"from_level": 1, "text": "on cooldown"},
                {"from_level": 36, "text": "on cooldown, before Whirlwind"},
            ]
        },
    )
    assert _condition_at(tmp_path, 10) == _condition_at(tmp_path, 35) == "on cooldown"
    assert _condition_at(tmp_path, 36) == "on cooldown, before Whirlwind"
    assert _condition_at(tmp_path, 60) == "on cooldown, before Whirlwind"


@pytest.mark.parametrize(
    ("value", "match"),
    [
        ([], "empty"),
        ([{"from_level": 1}], "exactly from_level and text"),
        ([{"from_level": 1, "text": "x", "to_level": 9}], "exactly from_level and text"),
        ([{"from_level": 0, "text": "x"}], "not a level 1 to 60"),
        ([{"from_level": 61, "text": "x"}], "not a level 1 to 60"),
        ([{"from_level": "1", "text": "x"}], "not a level 1 to 60"),
        (
            [{"from_level": 20, "text": "a"}, {"from_level": 10, "text": "b"}],
            "not strictly ascending",
        ),
        (
            [{"from_level": 10, "text": "a"}, {"from_level": 10, "text": "b"}],
            "not strictly ascending",
        ),
        ([{"from_level": 1, "text": "x" * 61}], "61 characters"),
        ([{"from_level": 1, "text": " padded"}], "leading or trailing space"),
        ([{"from_level": 1, "text": "a — b"}], "em dash"),
        ([{"from_level": 1, "text": ""}], "empty"),
    ],
)
def test_a_bad_variant_is_refused(tmp_path, value, match):
    _write_apl(tmp_path, {"0": value})
    with pytest.raises(AddonRotationError, match=match):
        build_rotation("warrior-fury", tmp_path, SPELLRANKS)


def test_a_first_variant_after_the_steps_first_band_is_refused(tmp_path):
    _write_apl(tmp_path, {"0": [{"from_level": 12, "text": "on cooldown"}]})
    with pytest.raises(
        AddonRotationError, match=r"starts at level 12, after the step's band at 10"
    ):
        build_rotation("warrior-fury", tmp_path, SPELLRANKS)


NAMES = ClassSpellNames(
    trained={"Rend": 4, "Multi-Shot": 18, "Arcane Shot": 6},
    talents={"Improved Rend": frozenset({12286}), "Bane": frozenset({17788})},
    unlearnable=frozenset({"Water Shield"}),
)


def test_a_longer_name_wins_over_one_inside_it():
    assert named_spells("with Improved Rend, then Rend", NAMES.all_names) == [
        "Improved Rend",
        "Rend",
    ]


def test_a_name_matches_whole_words_with_its_case():
    names = NAMES.all_names
    assert named_spells("rend it, Rending, Rend's bleed", names) == ["Rend"]


def test_known_names_hold_trainer_spells_by_level_and_talents_once_taken():
    rank_levels = {(12286, 1): 21, (12286, 2): 22}
    assert known_names(NAMES, 17, rank_levels) == {"Rend", "Arcane Shot"}
    assert known_names(NAMES, 21, rank_levels) == {
        "Rend",
        "Arcane Shot",
        "Multi-Shot",
        "Improved Rend",
    }


def test_unavailable_names_skip_the_rows_own_spell_and_known_ones():
    line = "when Multi-Shot is 1.5+ s from ready, before Arcane Shot"
    known = known_names(NAMES, 10, {})
    assert unavailable_names(line, "Arcane Shot", known, NAMES.all_names) == ["Multi-Shot"]
    assert unavailable_names(line, "Arcane Shot", known_names(NAMES, 18, {}), NAMES.all_names) == []
    assert unavailable_names("never Water Shield", "Rend", known, NAMES.all_names) == [
        "Water Shield"
    ]


def test_class_spell_names_read_trainables_talents_and_talent_auras(tmp_path):
    (tmp_path / "trainables").mkdir()
    (tmp_path / "talents").mkdir()
    trainables = [
        {
            "name": "Arcane Shot",
            "source": "skill_line_ability",
            "ranks": [{"id": 3044, "level": 6}],
        },
        {"name": "Counterattack", "source": "talent", "ranks": [{"id": 19306, "level": 25}]},
    ]
    talents = {
        "trees": [{"talents": [{"name": "Elemental Focus", "spell_id": 16164, "ranks": []}]}]
    }
    (tmp_path / "trainables" / "shaman.json").write_text(json.dumps({"trainables": trainables}))
    (tmp_path / "talents" / "shaman.json").write_text(json.dumps(talents))
    spellranks = {
        "classes": {
            "shaman": {
                "Clearcasting": [{"id": 16246, "level": 10, "rank": 0}],
                "Water Shield": [{"id": 408510, "level": 20, "rank": 0}],
            }
        }
    }
    names = load_class_spell_names(tmp_path, "shaman", spellranks, {16246: 16164})
    assert names.trained == {"Arcane Shot": 6}
    assert names.talents == {
        "Elemental Focus": frozenset({16164}),
        "Counterattack": frozenset({19306}),
        "Clearcasting": frozenset({16246}),
    }
    assert names.unlearnable == {"Water Shield"}


def test_no_panel_line_names_a_spell_the_character_cannot_have_at_its_band():
    """Coordinator ruling 2026-11-08: over every spec and every published band of the active
    build, a condition line names only class spells the character has at that band (trainer
    spells learned by then, talents the published build has taken by then) or the spell the
    row casts. Beast Mastery's Arcane Shot read "when Multi-Shot is 1.5+ s from ready" at 10."""
    build_dir = active_build_dir()
    violations = panel_name_violations(build_dir.parent, build_dir.name)
    assert [str(v) for v in violations] == []


def test_a_step_that_needs_a_trained_spell_known_waits_for_it(tmp_path):
    """The engine's spellIsKnown is false until the first rank is trained: Fire's "Fireball
    as your filler once you have Arcane Blast" printed from 10, Arcane Blast comes at 54."""
    gated = {
        "notes": "N.",
        "action": {
            "condition": {"spellIsKnown": {"spellId": {"spellId": 1680}}},
            "castSpell": {"spellId": {"spellId": 2687}},
        },
    }
    apl_dir = tmp_path / "apl"
    apl_dir.mkdir()
    (apl_dir / "warrior-fury.json").write_text(
        json.dumps({"spec": "warrior-fury", "rotation": {"priorityList": [gated]}}),
        encoding="utf-8",
    )
    bands = build_rotation("warrior-fury", tmp_path, SPELLRANKS)
    assert [band.level for band in bands if band.lines][0] == 36
