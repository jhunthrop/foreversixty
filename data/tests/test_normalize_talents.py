from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.normalize import write_json
from pipeline.normalize.talents import TalentDataError, normalize_talents

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
