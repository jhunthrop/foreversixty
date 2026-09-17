import pytest
from traits_fixture import TRAITS, trait_rows

from pipeline.csvio import read_csv
from pipeline.curves import load_rank_points
from pipeline.icons import PLACEHOLDER_ICON, icon_names
from pipeline.normalize.trait_trees import TraitRankError, build_trait_talent_trees
from pipeline.spelltext import load_spell_text


def fixture_records(**replace):
    return build_trait_talent_trees(
        trait_rows(**replace),
        read_csv(TRAITS / "ChrClasses.csv"),
        {int(r["ID"]): r["Name_lang"] for r in read_csv(TRAITS / "SpellName.csv")},
        load_spell_text(
            read_csv(TRAITS / "Spell.csv"),
            read_csv(TRAITS / "SpellMisc.csv"),
            read_csv(TRAITS / "SpellEffect.csv"),
            read_csv(TRAITS / "SpellDuration.csv"),
        ),
        load_rank_points(
            read_csv(TRAITS / "TraitDefinitionEffectPoints.csv"),
            read_csv(TRAITS / "CurvePoint.csv"),
        ),
        icon_names(read_csv(TRAITS / "ManifestInterfaceData.csv")),
        "1.60.1.69893",
    )


def test_only_the_class_with_a_trait_tree_is_emitted():
    records = fixture_records()
    assert [r.class_slug for r in records] == ["warrior"]
    assert records[0].build == "1.60.1.69893"
    assert records[0].class_id == 1


def test_the_three_trees_carry_the_games_order_and_their_background():
    (warrior,) = fixture_records()
    assert [(t.id, t.name, t.position, t.background) for t in warrior.trees] == [
        (161, "Arms", 0, "warriorarms"),
        (164, "Fury", 1, "warriorfury"),
        (163, "Protection", 2, "warriorprotection"),
    ]


def test_a_talent_is_its_node_with_the_clients_spell_name_and_icon():
    (warrior,) = fixture_records()
    arms = {t.id: t for t in warrior.trees[0].talents}
    heroic = arms[900001]
    assert heroic.name == "Improved Heroic Strike"
    assert heroic.icon == "ability_rogue_ambush"
    assert heroic.spell_id == 12282
    assert (heroic.tier, heroic.column, heroic.max_rank) == (0, 0, 3)


def test_every_rank_gets_its_own_text_from_the_curve():
    (warrior,) = fixture_records()
    arms = {t.id: t for t in warrior.trees[0].talents}
    assert [r.description for r in arms[900001].ranks] == [
        "Increases your Strength by 1%.",
        "Increases your Strength by 2%.",
        "Increases your Strength by 3%.",
    ]
    # Every rank carries the same spell id: the client has no per-rank spells.
    assert [r.spell_id for r in arms[900001].ranks] == [12282, 12282, 12282]


def test_a_talent_with_no_curve_repeats_the_clients_one_sentence():
    (warrior,) = fixture_records()
    fury = {t.id: t for t in warrior.trees[1].talents}
    assert fury[900013].max_rank == 1
    assert fury[900013].ranks[0].description == "Dazes the target for 6 sec."


def test_a_curve_with_more_points_than_the_talent_has_ranks_is_clamped():
    # 700004 is a 3-rank talent on a 5-point curve, the shape three of the
    # beta build's single-rank talents have.
    (warrior,) = fixture_records()
    arms = {t.id: t for t in warrior.trees[0].talents}
    assert len(arms[900004].ranks) == 3
    assert arms[900004].ranks[2].description == "Reduces the target's attack speed by 9%."


def test_a_multi_rank_talent_with_no_value_for_a_rank_is_refused():
    entries = read_csv(TRAITS / "TraitNodeEntry.csv")
    for entry in entries:
        if entry["ID"] == "800013":  # 1 rank and no curve; ask for 4
            entry["MaxRanks"] = "4"
    with pytest.raises(TraitRankError, match="700013"):
        fixture_records(node_entry=entries)


def test_prerequisites_are_carried_through_as_talent_ids():
    (warrior,) = fixture_records()
    arms = {t.id: t for t in warrior.trees[0].talents}
    assert (arms[900003].prereq_talent_id, arms[900003].prereq_rank) == (900002, 5)
    assert (arms[900001].prereq_talent_id, arms[900001].prereq_rank) == (None, None)


def test_a_talent_whose_spell_has_no_icon_falls_back_to_the_placeholder():
    (warrior,) = fixture_records()
    protection = {t.id: t for t in warrior.trees[2].talents}
    assert protection[900022].icon == PLACEHOLDER_ICON


def test_talents_are_sorted_by_cell():
    (warrior,) = fixture_records()
    for tree in warrior.trees:
        cells = [(t.tier, t.column) for t in tree.talents]
        assert cells == sorted(cells)
