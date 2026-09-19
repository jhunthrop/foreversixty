from pathlib import Path

import pytest

from pipeline.simdb.ratings import (
    RATING_STAT_COLUMNS,
    RatingFactorError,
    convert_rating_stats,
    load_rating_factors,
)

FIXTURE = Path(__file__).parent / "fixtures" / "sim" / "combatratings.txt"


def _write(tmp_path: Path, text: str) -> Path:
    build_dir = tmp_path / "build"
    (build_dir / "gametables").mkdir(parents=True)
    (build_dir / "gametables" / "combatratings.txt").write_text(text, encoding="utf-8")
    return build_dir


def test_factors_are_read_from_the_build_s_own_gametable(tmp_path: Path):
    """The committed build's real numbers: never typed as constants in the
    pipeline, always read off gametables/combatratings.txt."""
    build_dir = _write(tmp_path, FIXTURE.read_text(encoding="utf-8"))
    assert load_rating_factors(build_dir) == {
        "hit": 10.0,
        "crit": 14.0,
        "dodge": 12.0,
        "parry": 15.0,
        "block": 5.0,
        "defense": 1.0,
    }


def test_every_rating_stat_column_is_covered():
    """A key this module converts must name at least one real column."""
    for key, columns in RATING_STAT_COLUMNS.items():
        assert columns, key


def test_a_missing_gametables_file_is_a_clear_error(tmp_path: Path):
    build_dir = tmp_path / "build"
    build_dir.mkdir()
    with pytest.raises(RatingFactorError, match="gametables"):
        load_rating_factors(build_dir)


def test_a_missing_level_60_row_is_a_clear_error(tmp_path: Path):
    build_dir = _write(
        tmp_path,
        "Level\tHit - Melee\tHit - Ranged\tHit - Spell\tCrit - Melee\tCrit - Ranged\t"
        "Crit - Spell\tDodge\tParry\tBlock\tDefense Skill\n"
        "59\t10\t10\t10\t14\t14\t14\t12\t15\t5\t1\n",
    )
    with pytest.raises(RatingFactorError, match="level 60"):
        load_rating_factors(build_dir)


def test_schools_that_disagree_on_their_factor_are_an_error(tmp_path: Path):
    """The engine's Hit and Crit are each one unified Stat (statmap.py); a
    build whose melee and spell hit factors actually diverged at level 60
    would need a real design decision, not a silently picked column."""
    build_dir = _write(
        tmp_path,
        "Level\tHit - Melee\tHit - Ranged\tHit - Spell\tCrit - Melee\tCrit - Ranged\t"
        "Crit - Spell\tDodge\tParry\tBlock\tDefense Skill\n"
        "60\t10\t10\t11\t14\t14\t14\t12\t15\t5\t1\n",
    )
    with pytest.raises(RatingFactorError, match="disagree"):
        load_rating_factors(build_dir)


def test_convert_rating_stats_only_touches_known_keys():
    factors = {"hit": 10.0, "crit": 14.0}
    converted = convert_rating_stats(
        {"hit": 20.0, "crit": 28.0, "strength": 18.0}, factors
    )
    assert converted == {"hit": 2.0, "crit": 2.0, "strength": 18.0}


def test_convert_rating_stats_returns_a_new_mapping():
    """The caller's `resolve_item_values` output must survive untouched: the
    planner reads the same dict elsewhere and shows the client's raw rating
    number, not the engine's converted percentage."""
    original = {"hit": 20.0}
    convert_rating_stats(original, {"hit": 10.0})
    assert original == {"hit": 20.0}
