import json
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.icons import icon_names
from pipeline.normalize import write_model
from pipeline.normalize.talent_trees import build_talent_trees
from pipeline.spelltext import load_spell_text

HERE = Path(__file__).parent


def build_warrior():
    spell_text = load_spell_text(
        read_csv(HERE / "fixtures/Spell.csv"),
        read_csv(HERE / "fixtures/SpellMisc.csv"),
        read_csv(HERE / "fixtures/SpellEffect.csv"),
        read_csv(HERE / "fixtures/SpellDuration.csv"),
    )
    records = build_talent_trees(
        read_csv(HERE / "fixtures/Talent.csv"),
        read_csv(HERE / "fixtures/TalentTab.csv"),
        read_csv(HERE / "fixtures/ChrClasses.csv"),
        {int(r["ID"]): r["Name_lang"] for r in read_csv(HERE / "fixtures/SpellName.csv")},
        spell_text,
        icon_names(read_csv(HERE / "fixtures/ManifestInterfaceData.csv")),
        "1.0.0.1",
    )
    return records


def test_only_classes_with_tabs_are_emitted():
    assert [r.class_slug for r in build_warrior()] == ["warrior"]


def test_warrior_trees_match_golden(tmp_path: Path):
    out = tmp_path / "warrior.json"
    write_model(build_warrior()[0], out)
    assert out.read_text() == (HERE / "golden/talents_warrior.json").read_text()


def test_trees_are_ordered_by_position():
    trees = build_warrior()[0].trees
    assert [(t.id, t.name, t.position) for t in trees] == [(161, "Arms", 0), (164, "Fury", 1)]


def test_ranks_carry_one_entry_per_spell_rank_with_resolved_text():
    talents = {t.id: t for t in build_warrior()[0].trees[0].talents}
    improved = talents[124]
    assert improved.max_rank == 3
    assert [r.spell_id for r in improved.ranks] == [12282, 12663, 12664]
    assert [r.description for r in improved.ranks] == [
        "Increases your Strength by 2%.",
        "Increases your Strength by 4%.",
        "Increases your Strength by 6%.",
    ]


def test_prereq_rank_is_a_point_count_not_a_rank_index():
    talents = {t.id: t for t in build_warrior()[0].trees[0].talents}
    assert (talents[130].prereq_talent_id, talents[130].prereq_rank) == (124, 4)
    assert (talents[124].prereq_talent_id, talents[124].prereq_rank) == (None, None)


def test_icon_is_the_first_rank_spell_icon():
    talents = {t.id: t for t in build_warrior()[0].trees[0].talents}
    assert talents[124].icon == "ability_golemthunderclap"


def test_output_is_deterministic(tmp_path: Path):
    first, second = tmp_path / "a.json", tmp_path / "b.json"
    write_model(build_warrior()[0], first)
    write_model(build_warrior()[0], second)
    assert first.read_bytes() == second.read_bytes()
    assert json.loads(first.read_text())["build"] == "1.0.0.1"
