import json

import pytest

from pipeline.normalize.forever_talents import (
    ForeverTalentError,
    normalize_forever_talents,
)

CLASSES = [{"id": 1, "name": "Warrior", "slug": "warrior"}]
TREE_CLASS = {161: 1, 164: 1}
TREE_NAMES = {161: "Arms", 164: "Fury"}
TREE_BACKGROUNDS = {161: "warriorarms", 164: "warriorfury"}


def payload(**overrides):
    base = {
        "trees": {
            "161": {"id": 161, "description": "WarriorArms"},
            "164": {"id": 164, "description": "WarriorFury"},
        },
        "talents": {
            "161": {
                "105954": {
                    "id": 105954,
                    "row": 0,
                    "col": 1,
                    "icon": "ability_rogue_ambush",
                    "name": "Deflection",
                    "ranks": [None, None],
                    "requires": [],
                    "descriptions": {"1": "Parry 1%.", "2": "Parry 2%."},
                },
                "105956": {
                    "id": 105956,
                    "row": 2,
                    "col": 1,
                    "icon": "ability_warrior_savageblow",
                    "name": "Anger Management",
                    "ranks": [None],
                    "requires": [{"id": 105954, "qty": 2}],
                    "descriptions": {"1": "Rage over time."},
                },
            },
            "164": {
                "105960": {
                    "id": 105960,
                    "row": 0,
                    "col": 0,
                    "icon": "ability_warrior_rampage",
                    "name": "Booming Voice",
                    "ranks": [None],
                    "requires": [],
                    "descriptions": {"1": "Shout lasts longer."},
                }
            },
        },
    }
    base.update(overrides)
    return base


def normalize(p=None):
    return normalize_forever_talents(
        p or payload(),
        build="forever-prebeta",
        classes=CLASSES,
        tree_class=TREE_CLASS,
        tree_names=TREE_NAMES,
        tree_backgrounds=TREE_BACKGROUNDS,
    )


def test_one_file_per_class_with_its_trees_in_id_order():
    out = normalize()
    assert len(out) == 1
    warrior = out[0]
    assert (warrior.class_id, warrior.class_slug) == (1, "warrior")
    assert warrior.build == "forever-prebeta"
    assert [(t.id, t.name, t.position) for t in warrior.trees] == [
        (161, "Arms", 0),
        (164, "Fury", 1),
    ]


def test_the_tree_name_comes_from_our_tables_not_from_wowheads_glued_string():
    # Wowhead says "WarriorArms" and "HunterBeastMastery"; un-gluing those would have to
    # guess where the words break, and the tree ids match exactly, so we look them up.
    out = normalize()
    assert [t.name for t in out[0].trees] == ["Arms", "Fury"]


def test_every_rank_keeps_its_own_tooltip_text():
    deflection = normalize()[0].trees[0].talents[0]
    assert deflection.max_rank == 2
    assert [r.description for r in deflection.ranks] == ["Parry 1%.", "Parry 2%."]
    assert {r.spell_id for r in deflection.ranks} == {105954}


def test_a_prerequisite_becomes_a_talent_id_and_a_rank():
    anger = normalize()[0].trees[0].talents[1]
    assert (anger.prereq_talent_id, anger.prereq_rank) == (105954, 2)
    deflection = normalize()[0].trees[0].talents[0]
    assert (deflection.prereq_talent_id, deflection.prereq_rank) == (None, None)


def test_talents_come_out_in_grid_order():
    assert [(t.tier, t.column) for t in normalize()[0].trees[0].talents] == [(0, 1), (2, 1)]


def test_a_rank_count_that_disagrees_with_the_tooltips_is_refused():
    p = payload()
    p["talents"]["161"]["105954"]["ranks"] = [None, None, None]
    with pytest.raises(ForeverTalentError, match="3 ranks but 2 descriptions"):
        normalize(p)


def test_a_prerequisite_outside_the_tree_is_refused():
    p = payload()
    p["talents"]["161"]["105956"]["requires"] = [{"id": 999999, "qty": 1}]
    with pytest.raises(ForeverTalentError, match="not in the same tree"):
        normalize(p)


def test_a_second_prerequisite_is_refused_rather_than_dropped():
    p = payload()
    p["talents"]["161"]["105956"]["requires"].append({"id": 105954, "qty": 1})
    with pytest.raises(ForeverTalentError, match="2 prerequisites"):
        normalize(p)


def test_a_tree_we_do_not_know_is_refused():
    p = payload()
    p["trees"]["999"] = {"id": 999, "description": "MonkBrewmaster"}
    p["talents"]["999"] = {}
    with pytest.raises(ForeverTalentError, match="not one of ours"):
        normalize(p)


def test_a_tree_with_no_name_in_our_tables_is_refused():
    with pytest.raises(ForeverTalentError, match="no name in our tables"):
        normalize_forever_talents(
            payload(),
            build="forever-prebeta",
            classes=CLASSES,
            tree_class=TREE_CLASS,
            tree_names={161: "Arms"},
            tree_backgrounds=TREE_BACKGROUNDS,
        )


def test_a_tree_with_no_background_in_our_tables_is_refused():
    with pytest.raises(ForeverTalentError, match="no background in our tables"):
        normalize_forever_talents(
            payload(),
            build="forever-prebeta",
            classes=CLASSES,
            tree_class=TREE_CLASS,
            tree_names=TREE_NAMES,
            tree_backgrounds={161: "warriorarms"},
        )


def test_a_talent_with_no_ranks_is_refused():
    p = payload()
    p["talents"]["161"]["105954"]["ranks"] = []
    p["talents"]["161"]["105954"]["descriptions"] = {}
    with pytest.raises(ForeverTalentError, match="no ranks"):
        normalize(p)


def test_an_icon_no_client_holds_is_repointed_at_the_placeholder(tmp_path):
    from pipeline.forever import _repoint_to_placeholder
    from pipeline.icons import PLACEHOLDER_ICON

    talents = tmp_path / "talents"
    talents.mkdir()
    (talents / "druid.json").write_text(
        json.dumps(
            {
                "trees": [
                    {
                        "talents": [
                            {"icon": "classic_ability_druid_demoralizingroar"},
                            {"icon": "spell_nature_regeneration"},
                        ]
                    }
                ]
            }
        )
    )
    changed = _repoint_to_placeholder(tmp_path, {"classic_ability_druid_demoralizingroar"})
    assert changed == 1
    icons = [
        t["icon"]
        for tree in json.loads((talents / "druid.json").read_text())["trees"]
        for t in tree["talents"]
    ]
    assert icons == [PLACEHOLDER_ICON, "spell_nature_regeneration"]
