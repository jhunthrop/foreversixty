"""What is committed under builds/1.60.1.69893/gametables/ must be this
build's own game tables, not another build's.

One of the handful of tests the Global Constraints allow a build string in:
the column names and spot values below are only right for this build.
"""

from functools import cache
from pathlib import Path

from pipeline.gametables import (
    ABSENT_FROM_THE_CLASSIC_LINEAGE,
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
