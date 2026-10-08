"""What is committed under builds/1.60.1.69893/gametables/ must be this
build's own game tables, not another build's.

One of the handful of tests the Global Constraints allow a build string in:
the column names and spot values below are only right for this build.
"""

import csv
from functools import cache
from pathlib import Path

import pytest

from pipeline.gametables import (
    ABSENT_FROM_THE_CLASSIC_LINEAGE,
    EXPECTED_STAT_COLUMNS,
    EXPECTED_STAT_FILE,
    GAME_TABLE_IDS,
    GameTable,
    parse_game_table,
)

BUILD = "1.60.1.69893"
GAMETABLES_DIR = Path("builds") / BUILD / "gametables"

#: The class columns this client's basemp.txt carries, in its own order --
#: which is NOT the order the engine's tools/base_stats_parser.py assumes for
#: octbasempbyclass.txt, so the engine must index these by name.
BASE_MP_CLASSES = (
    "Rogue",
    "Druid",
    "Hunter",
    "Mage",
    "Paladin",
    "Priest",
    "Shaman",
    "Warlock",
    "Warrior",
    "Death Knight",
    "Monk",
    "Demon Hunter",
    "Evoker",
    "Adventurer",
    "Traveler",
)


@cache
def table(name: str) -> GameTable:
    return parse_game_table(name, (GAMETABLES_DIR / name).read_text(encoding="utf-8"))


def value(name: str, key: str, column: str) -> str:
    row = table(name)
    return row.rows[key][row.columns.index(column) - 1]


def test_exactly_the_tables_this_lineage_serves_are_committed():
    assert {path.name for path in GAMETABLES_DIR.glob("*.txt")} == set(GAME_TABLE_IDS)


def test_none_of_the_absent_six_was_written():
    """They answer 200 with an empty body on this build; an unpinned fetch
    would have written retail's copies instead."""
    for name in ABSENT_FROM_THE_CLASSIC_LINEAGE:
        assert not (GAMETABLES_DIR / name).exists(), name


def test_base_mana_names_its_classes_and_states_this_build_s_numbers():
    assert table("basemp.txt").columns == ("Level", *BASE_MP_CLASSES)
    assert value("basemp.txt", "60", "Mage") == "1213"
    assert value("basemp.txt", "60", "Paladin") == "1512"
    assert value("basemp.txt", "60", "Warrior") == "0"


def test_health_per_stamina_is_the_two_column_table_it_should_be():
    assert table("hppersta.txt").columns == ("Level", "Health")
    assert value("hppersta.txt", "60", "Health") == "10"


def test_combat_ratings_is_keyed_by_level_and_names_its_ratings():
    ratings = table("combatratings.txt")
    assert ratings.columns[0] == "Level"
    for name in ("Dodge", "Parry", "Hit - Melee", "Crit - Melee", "Crit - Spell"):
        assert name in ratings.columns, name
    assert value("combatratings.txt", "60", "Dodge") == "12"
    assert value("combatratings.txt", "60", "Parry") == "15"
    assert value("combatratings.txt", "60", "Crit - Melee") == "14"


def test_every_committed_table_covers_the_same_levels():
    """All three are keyed by level 1-123 on this client. A table that stopped
    short would leave the engine reading a missing level as zero."""
    for name in GAME_TABLE_IDS:
        assert {int(key) for key in table(name).rows} == set(range(1, 124)), name


def test_every_committed_table_parses():
    """A table that does not parse here is one the engine's own
    csv.reader(delimiter='\\t') would misread."""
    for name in GAME_TABLE_IDS:
        assert table(name).rows, name


@cache
def expected_stat() -> dict[tuple[int, int], dict[str, str]]:
    with (GAMETABLES_DIR / EXPECTED_STAT_FILE).open(encoding="utf-8", newline="") as handle:
        rows = list(csv.DictReader(handle))
    return {(int(row["ClassID"]), int(row["Level"])): row for row in rows}


def test_expected_stat_names_the_columns_the_engine_reads():
    with (GAMETABLES_DIR / EXPECTED_STAT_FILE).open(encoding="utf-8", newline="") as handle:
        header = next(csv.reader(handle))
    assert set(EXPECTED_STAT_COLUMNS) <= set(header)


def test_expected_stat_covers_every_class_at_every_level():
    """Nine playable classes, levels 1 to 123, like the three game tables."""
    keys = set(expected_stat())
    assert {class_id for class_id, _ in keys} == {1, 2, 3, 4, 5, 7, 8, 9, 11}
    for class_id in {1, 2, 3, 4, 5, 7, 8, 9, 11}:
        assert {level for c, level in keys if c == class_id} == set(range(1, 124)), class_id


def test_expected_stat_states_this_build_s_crit_curves():
    """Level 60 is vanilla's rate (a warrior's 20 Agility per 1%); below it
    the client asks less Agility per point, which is the curve the engine
    pins at 60 unless it reads this file. Intellect is the wowhead planner's
    curve, from this table. Warriors and rogues have no spell crit row."""
    rows = expected_stat()

    def rate(class_id: int, level: int, column: str) -> float:
        return float(rows[(class_id, level)][column])

    close = 5e-7
    assert rate(1, 60, "CritPerAgility") == pytest.approx(0.0005, abs=close)
    assert rate(1, 30, "CritPerAgility") == pytest.approx(0.000962, abs=close)
    assert rate(1, 1, "CritPerAgility") == pytest.approx(0.0025, abs=close)
    assert rate(4, 60, "CritPerAgility") == pytest.approx(0.000345, abs=close)
    assert rate(3, 60, "CritPerAgility") == pytest.approx(0.000189, abs=close)
    assert rate(11, 60, "SpellCritPerIntellect") == pytest.approx(0.000167, abs=close)
    assert rate(11, 30, "SpellCritPerIntellect") == pytest.approx(0.000352, abs=close)
    assert rows[(1, 60)]["SpellCritPerIntellect"] == "0"
    assert rows[(4, 60)]["SpellCritPerIntellect"] == "0"


def test_expected_stat_base_mana_is_the_base_mana_table():
    """The same numbers as basemp.txt, so a reader of either agrees with the
    other."""
    rows = expected_stat()
    assert rows[(8, 60)]["BaseMana"] == value("basemp.txt", "60", "Mage")
    assert rows[(2, 60)]["BaseMana"] == value("basemp.txt", "60", "Paladin")
    assert rows[(11, 30)]["BaseMana"] == value("basemp.txt", "30", "Druid")
