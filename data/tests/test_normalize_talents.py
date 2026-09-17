from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.models import ClassTalents, TalentEntry, TalentRank, TalentTree
from pipeline.normalize import write_json
from pipeline.normalize.talents import TalentDataError, flat_talents, normalize_talents

HERE = Path(__file__).parent


def test_talents_match_golden(tmp_path: Path):
    talent_rows = read_csv(HERE / "fixtures/Talent.csv")
    tab_rows = read_csv(HERE / "fixtures/TalentTab.csv")
    nodes = normalize_talents(talent_rows, tab_rows)
    out = tmp_path / "talents.json"
    write_json(nodes, out)
    assert out.read_text() == (HERE / "golden/talents.json").read_text()


def test_a_zero_class_id_falls_back_to_the_tab_class_mask():
    """The 1.60 client (Forever beta) leaves Talent.ClassID at 0 for every row; class
    ownership moved to TalentTab.ClassMask. Era's Talent.ClassID is still the real
    value (see test_talents_match_golden's fixture, ClassID 1 throughout) and is used
    as-is; a build that states 0 falls back to the owning tab's single-bit mask."""
    talent_rows = read_csv(HERE / "fixtures/Talent_1_60.csv")
    tab_rows = read_csv(HERE / "fixtures/TalentTab.csv")  # tab 161 has ClassMask 1: warrior
    nodes = normalize_talents(talent_rows, tab_rows)
    assert [n.class_id for n in nodes] == [1]


def test_a_tab_whose_mask_names_no_single_class_is_an_error():
    talent_rows = read_csv(HERE / "fixtures/Talent_1_60.csv")
    tab_rows = read_csv(HERE / "fixtures/TalentTab.csv")
    for row in tab_rows:
        if row["ID"] == "161":
            row["ClassMask"] = "3"  # two classes at once: not a single bit
    with pytest.raises(TalentDataError, match="161"):
        normalize_talents(talent_rows, tab_rows)


def _entry(talent_id, tier, column, prereq=None, prereq_rank=None):
    return TalentEntry(
        id=talent_id,
        name=f"Talent {talent_id}",
        icon="icon",
        max_rank=2,
        tier=tier,
        column=column,
        prereq_talent_id=prereq,
        prereq_rank=prereq_rank,
        ranks=[
            TalentRank(spell_id=100 + talent_id, description="a"),
            TalentRank(spell_id=200 + talent_id, description="b"),
        ],
        spell_id=100 + talent_id,
    )


def test_flat_talents_folds_already_built_records_into_talent_nodes():
    """flat_talents is the trait path's (and forever.py's) equivalent of
    normalize_talents: both must produce the same TalentNode shape from
    different sources, since talents.json does not distinguish where a
    build's trees came from."""
    records = [
        ClassTalents(
            build="b",
            class_id=1,
            class_slug="warrior",
            trees=[
                TalentTree(
                    id=161,
                    name="Arms",
                    position=0,
                    background="warriorarms",
                    talents=[
                        _entry(900001, 0, 0),
                        _entry(900002, 1, 0, prereq=900001, prereq_rank=2),
                    ],
                ),
                TalentTree(id=164, name="Fury", position=1, background="warriorfury", talents=[]),
            ],
        )
    ]
    nodes = flat_talents(records)
    assert [n.id for n in nodes] == [900001, 900002]
    first, second = nodes
    assert (first.tab_id, first.tab_name, first.class_id) == (161, "Arms", 1)
    assert (first.tier, first.column) == (0, 0)
    assert first.spell_ids == [900101, 900201]
    assert first.prereq_talent_id is None
    assert second.prereq_talent_id == 900001


def test_flat_talents_agrees_with_normalize_talents_on_the_era_fixture():
    """Both readers must emit the same TalentNode for the same data: build a
    ClassTalents record with the same shape normalize_talents would produce
    from the Era fixture's one talent, and compare the flattened rows."""
    talent_rows = read_csv(HERE / "fixtures/Talent.csv")
    tab_rows = read_csv(HERE / "fixtures/TalentTab.csv")
    node = normalize_talents(talent_rows, tab_rows)[0]
    record = ClassTalents(
        build="b",
        class_id=node.class_id,
        class_slug="warrior",
        trees=[
            TalentTree(
                id=node.tab_id,
                name=node.tab_name,
                position=0,
                background="warriorarms",
                talents=[
                    TalentEntry(
                        id=node.id,
                        name="whatever",
                        icon="icon",
                        max_rank=len(node.spell_ids),
                        tier=node.tier,
                        column=node.column,
                        prereq_talent_id=node.prereq_talent_id,
                        prereq_rank=1 if node.prereq_talent_id else None,
                        ranks=[TalentRank(spell_id=s, description="") for s in node.spell_ids],
                        spell_id=node.spell_ids[0],
                    )
                ],
            )
        ],
    )
    (folded,) = flat_talents([record])
    assert folded == node
