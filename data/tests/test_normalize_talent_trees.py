import json
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.icons import PLACEHOLDER_ICON, icon_names
from pipeline.normalize import write_model
from pipeline.normalize.talent_trees import build_talent_trees
from pipeline.spelltext import load_spell_text

HERE = Path(__file__).parent
#: The Arms talent used across these tests, and the spell and icon file backing it.
TALENT_ID = 124
FIRST_RANK_SPELL_ID = 12282
ICON_FILE_ID = 132154


def fixture_spell_names() -> dict[int, str]:
    return {int(r["ID"]): r["Name_lang"] for r in read_csv(HERE / "fixtures/SpellName.csv")}


def fixture_icons() -> dict[int, str]:
    return icon_names(read_csv(HERE / "fixtures/ManifestInterfaceData.csv"))


def build_records(
    spell_names: dict[int, str] | None = None,
    icons: dict[int, str] | None = None,
):
    spell_text = load_spell_text(
        read_csv(HERE / "fixtures/Spell.csv"),
        read_csv(HERE / "fixtures/SpellMisc.csv"),
        read_csv(HERE / "fixtures/SpellEffect.csv"),
        read_csv(HERE / "fixtures/SpellDuration.csv"),
    )
    return build_talent_trees(
        read_csv(HERE / "fixtures/Talent.csv"),
        read_csv(HERE / "fixtures/TalentTab.csv"),
        read_csv(HERE / "fixtures/ChrClasses.csv"),
        fixture_spell_names() if spell_names is None else spell_names,
        spell_text,
        fixture_icons() if icons is None else icons,
        "1.0.0.1",
    )


def build_warrior():
    return build_records()


def arms_talents(records) -> dict:
    return {t.id: t for t in records[0].trees[0].talents}


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


def test_a_talent_whose_icon_the_client_cannot_resolve_falls_back_to_the_placeholder(caplog):
    """An empty icon would be emitted as icons/.webp and 404; the placeholder will not."""
    icons = {k: v for k, v in fixture_icons().items() if k != ICON_FILE_ID}
    with caplog.at_level("WARNING"):
        improved = arms_talents(build_records(icons=icons))[TALENT_ID]
    assert improved.icon == PLACEHOLDER_ICON
    assert improved.icon != ""
    messages = [record.getMessage() for record in caplog.records]
    assert any(str(TALENT_ID) in m and str(ICON_FILE_ID) in m for m in messages)


def test_a_talent_whose_first_rank_spell_has_no_name_gets_an_obvious_placeholder(caplog):
    """A nameless cell is not acceptable, and the pipeline never invents a spell name."""
    spell_names = {k: v for k, v in fixture_spell_names().items() if k != FIRST_RANK_SPELL_ID}
    with caplog.at_level("WARNING"):
        improved = arms_talents(build_records(spell_names=spell_names))[TALENT_ID]
    assert improved.name == f"Unknown talent {TALENT_ID}"
    messages = [record.getMessage() for record in caplog.records]
    assert any(str(TALENT_ID) in m and str(FIRST_RANK_SPELL_ID) in m for m in messages)


def test_every_emitted_talent_icon_and_name_is_usable():
    """No talent may carry an empty icon or name: the site builds icons/<icon>.webp from
    the one and renders the other as the talent's label."""
    talents = [t for r in build_warrior() for tree in r.trees for t in tree.talents]
    assert talents
    assert all(t.icon and t.name for t in talents)
