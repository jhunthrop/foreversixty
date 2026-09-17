import pytest

from pipeline.models import ClassTalents, TalentEntry, TalentRank, TalentTree
from pipeline.wowhead_diff import SnapshotShapeError, diff_snapshot


def entry(talent_id, name, tier, column, max_rank=1, prereq=None, prereq_rank=None):
    return TalentEntry(
        id=talent_id,
        name=name,
        icon="icon",
        max_rank=max_rank,
        tier=tier,
        column=column,
        prereq_talent_id=prereq,
        prereq_rank=prereq_rank,
        ranks=[TalentRank(spell_id=1, description="x") for _ in range(max_rank)],
        spell_id=1,
    )


def records(talents):
    return [
        ClassTalents(
            build="b",
            class_id=1,
            class_slug="warrior",
            trees=[
                TalentTree(
                    id=161, name="Arms", position=0, talents=talents, background="warriorarms"
                )
            ],
        )
    ]


def snapshot(talents):
    return {
        "trees": {"161": {"id": 161, "description": "WarriorArms"}},
        "talents": {"161": {str(t["id"]): t for t in talents}},
    }


def snap_talent(talent_id, name, row, col, ranks=1, requires=()):
    return {
        "id": talent_id,
        "name": name,
        "row": row,
        "col": col,
        "ranks": [None] * ranks,
        "descriptions": {str(n): "x" for n in range(1, ranks + 1)},
        "requires": [{"id": rid, "qty": qty} for rid, qty in requires],
    }


def test_identical_trees_report_no_change():
    result = diff_snapshot(
        snapshot([snap_talent(1, "Deflection", 0, 1)]),
        records([entry(900002, "Deflection", 0, 1)]),
    )
    assert result["totals"] == {"snapshot": 1, "build": 1}
    assert result["trees"] == []


def test_a_new_name_in_the_same_cell_is_a_rename():
    result = diff_snapshot(
        snapshot([snap_talent(1, "Drain Hope", 6, 1)]),
        records([entry(905916, "Wrack", 6, 1)]),
    )
    (tree,) = result["trees"]
    assert tree["renamed"] == [{"cell": [6, 1], "snapshot": "Drain Hope", "build": "Wrack"}]
    assert tree["added"] == [] and tree["removed"] == []


def test_a_name_that_appears_in_another_cell_is_a_move():
    result = diff_snapshot(
        snapshot([snap_talent(1, "Tidal Mastery", 0, 2), snap_talent(2, "Totemic Focus", 3, 0)]),
        records([entry(11, "Totemic Focus", 0, 2), entry(12, "Tidal Mastery", 3, 0)]),
    )
    (tree,) = result["trees"]
    assert tree["moved"] == [
        {"name": "Tidal Mastery", "snapshot": [0, 2], "build": [3, 0]},
        {"name": "Totemic Focus", "snapshot": [3, 0], "build": [0, 2]},
    ]
    assert tree["renamed"] == []


def test_an_empty_cell_on_one_side_is_an_addition_or_a_removal():
    result = diff_snapshot(
        snapshot([snap_talent(1, "Balance of Nature", 2, 3)]),
        records([entry(11, "Improved Serpent Sting", 3, 3, max_rank=5)]),
    )
    (tree,) = result["trees"]
    assert tree["removed"] == [{"cell": [2, 3], "name": "Balance of Nature", "ranks": 1}]
    assert tree["added"] == [{"cell": [3, 3], "name": "Improved Serpent Sting", "ranks": 5}]


def test_rank_and_prerequisite_changes_are_listed_by_name():
    result = diff_snapshot(
        snapshot(
            [
                snap_talent(1, "Nature's Majesty", 1, 2, ranks=2, requires=((2, 1),)),
                snap_talent(2, "Nature's Splendor", 2, 2, ranks=3),
            ]
        ),
        records(
            [
                entry(11, "Nature's Majesty", 1, 2, max_rank=2),
                entry(12, "Nature's Splendor", 2, 2, max_rank=1, prereq=11, prereq_rank=2),
            ]
        ),
    )
    (tree,) = result["trees"]
    assert tree["rank_changes"] == [{"name": "Nature's Splendor", "snapshot": 3, "build": 1}]
    assert tree["prerequisite_changes"] == [
        {"name": "Nature's Majesty", "snapshot": ["Nature's Splendor", 1], "build": None},
        {"name": "Nature's Splendor", "snapshot": None, "build": ["Nature's Majesty", 2]},
    ]


def test_a_tree_the_snapshot_names_that_the_build_lacks_is_reported():
    """The symmetric case to missing_from_snapshot: a tree id only the
    snapshot has is otherwise invisible except as a mismatch between the two
    totals, which is easy to miss since the totals disagree for other
    reasons on every real run so far."""
    snap = snapshot([snap_talent(1, "Deflection", 0, 1)])
    snap["trees"]["999"] = {"id": 999, "description": "GoneTree"}
    result = diff_snapshot(snap, records([entry(900002, "Deflection", 0, 1)]))
    assert {"tree_id": 999, "tree": "GoneTree", "missing_from_build": True} in result["trees"]


def test_two_snapshot_talents_sharing_a_cell_is_an_error_not_a_silent_drop():
    """The build side's cells are unique by construction (see
    tests/test_beta_build.py); the snapshot side is a hand-saved payload with
    no such guarantee, and the dict comprehensions in _tree_diff would
    otherwise silently keep only one of two talents sharing a cell."""
    snap = snapshot([snap_talent(1, "Alpha", 0, 0), snap_talent(2, "Beta", 0, 0)])
    with pytest.raises(SnapshotShapeError, match="161"):
        diff_snapshot(snap, records([entry(900002, "Alpha", 0, 0)]))
