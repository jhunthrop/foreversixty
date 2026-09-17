from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.curves import CurveDataError, load_rank_points

HERE = Path(__file__).parent
TRAITS = HERE / "fixtures/traits"


def fixture_points():
    return load_rank_points(
        read_csv(TRAITS / "TraitDefinitionEffectPoints.csv"),
        read_csv(TRAITS / "CurvePoint.csv"),
    )


def test_each_rank_gets_the_curve_value_for_every_effect_of_the_definition():
    points = fixture_points()
    # 700001 has one effect; 700002 has two, whose curves rise at different rates.
    assert points.for_rank(700001, 1) == {0: 1}
    assert points.for_rank(700001, 3) == {0: 3}
    assert points.for_rank(700002, 1) == {0: 2, 1: 20}
    assert points.for_rank(700002, 5) == {0: 10, 1: 100}


def test_highest_rank_is_the_last_point_on_the_curve():
    points = fixture_points()
    assert points.highest_rank(700001) == 3
    assert points.highest_rank(700002) == 5


def test_a_definition_with_no_curve_has_no_points_and_says_so_quietly():
    points = fixture_points()
    assert points.for_rank(700013, 1) == {}
    assert points.highest_rank(700013) == 0


def test_an_operation_other_than_setting_the_value_is_refused():
    rows = [
        {
            "ID": "1",
            "TraitDefinitionID": "700001",
            "EffectIndex": "0",
            "OperationType": "1",
            "CurveID": "90001",
        }
    ]
    with pytest.raises(CurveDataError, match="OperationType 1"):
        load_rank_points(rows, read_csv(TRAITS / "CurvePoint.csv"))


def test_a_curve_with_no_points_is_refused():
    rows = [
        {
            "ID": "1",
            "TraitDefinitionID": "700001",
            "EffectIndex": "0",
            "OperationType": "0",
            "CurveID": "99999",
        }
    ]
    with pytest.raises(CurveDataError, match="99999"):
        load_rank_points(rows, read_csv(TRAITS / "CurvePoint.csv"))
