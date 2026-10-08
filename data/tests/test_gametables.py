import json
from pathlib import Path

import httpx
import pytest

from pipeline.casc import CascError
from pipeline.gametables import (
    ABSENT_FROM_THE_CLASSIC_LINEAGE,
    EXPECTED_STAT_COLUMNS,
    EXPECTED_STAT_FILE,
    EXPECTED_STAT_TABLE,
    GAME_TABLE_IDS,
    GameTableError,
    parse_game_table,
    write_game_tables,
)
from pipeline.manifest import verify, write_manifest

BUILD = "9.9.9.9"

#: A header row then one row per level, tab separated, CRLF, exactly as the
#: client ships it. Shaped like the real basemp.txt, trimmed to four classes.
BASE_MP = "Level\tRogue\tMage\tPaladin\tWarrior\r\n1\t0\t100\t60\t0\r\n60\t0\t1213\t1512\t0\r\n"

#: The DB2 table beside them, as wago.tools serves it: a header naming the
#: columns the engine reads, then one row per class per level. Two classes,
#: two levels.
EXPECTED_STAT = (
    "ID,Level,ClassID,ContentSetID,BaseMana,HPPerStamina,CritPerAgility,SpellCritPerIntellect\n"
    "1,1,1,0,0,10,0.00249999994,0\n"
    "2,60,1,0,0,10,0.000500000024,0\n"
    "3,1,8,0,100,10,0.000899999985,0.00192000007\n"
    "4,60,8,0,1213,10,0.000514000026,0.000167999996\n"
)

EXPECTED_STAT_PATH = f"/db2/{EXPECTED_STAT_TABLE}/csv"


def transport(
    calls: list[str],
    bodies: dict[int, bytes] | None = None,
    expected_stat: str = EXPECTED_STAT,
) -> httpx.MockTransport:
    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(str(request.url))
        if request.url.path == EXPECTED_STAT_PATH:
            return httpx.Response(200, text=expected_stat)
        file_id = int(request.url.path.rsplit("/", 1)[-1])
        if bodies is not None and file_id in bodies:
            return httpx.Response(200, content=bodies[file_id])
        return httpx.Response(200, content=BASE_MP.encode("utf-8"))

    return httpx.MockTransport(handler)


def build_dir_with_manifest(tmp_path: Path) -> Path:
    build = tmp_path / BUILD
    build.mkdir(parents=True)
    (build / "items.json").write_text("[]\n")
    write_manifest(
        build, build=BUILD, product="wow_classic_beta", fetched_at="2026-01-01T00:00:00Z"
    )
    return build


def run(
    tmp_path: Path,
    calls: list[str],
    bodies: dict[int, bytes] | None = None,
    expected_stat: str = EXPECTED_STAT,
) -> Path:
    client = httpx.Client(
        transport=transport(calls, bodies, expected_stat), base_url="https://wago.tools"
    )
    return write_game_tables(BUILD, root=tmp_path, cache_dir=tmp_path / "cache", client=client)


def casc_calls(calls: list[str]) -> list[str]:
    return [call for call in calls if EXPECTED_STAT_PATH not in call]


def test_only_the_tables_this_lineage_serves_are_fetched():
    """Of the seven files tools/base_stats_parser.py names, this client ships
    one: combatratings.txt. The other six answer 200 with an empty body on
    both 1.60.1.69893 and 1.15.9.69722, because their file data ids were
    added to retail after Classic forked."""
    assert set(GAME_TABLE_IDS) == {"combatratings.txt", "basemp.txt", "hppersta.txt"}
    assert GAME_TABLE_IDS == {
        "basemp.txt": 1391664,
        "combatratings.txt": 1391669,
        "hppersta.txt": 1391642,
    }
    assert set(ABSENT_FROM_THE_CLASSIC_LINEAGE) == {
        "octbasempbyclass.txt",
        "chancetomeleecrit.txt",
        "chancetomeleecritbase.txt",
        "chancetospellcrit.txt",
        "chancetospellcritbase.txt",
        "octclasscombatratingscalar.txt",
    }
    assert not set(GAME_TABLE_IDS) & set(ABSENT_FROM_THE_CLASSIC_LINEAGE)


def test_every_id_is_in_the_block_the_classic_lineage_serves():
    """Every file data id this lineage answers for is in the original ~1.39M
    block; every id a later retail expansion added is empty. That is the
    pattern, and it is why the absent list is what it is."""
    assert all(1_000_000 < file_id < 1_400_000 for file_id in GAME_TABLE_IDS.values())
    assert all(file_id > 2_000_000 for file_id in ABSENT_FROM_THE_CLASSIC_LINEAGE.values())


def test_parse_reads_the_header_and_keys_rows_by_their_first_column():
    table = parse_game_table("basemp.txt", BASE_MP)
    assert table.columns == ("Level", "Rogue", "Mage", "Paladin", "Warrior")
    assert set(table.rows) == {"1", "60"}
    assert table.rows["60"][table.columns.index("Mage") - 1] == "1213"


def test_parse_tolerates_crlf_and_a_trailing_newline():
    assert parse_game_table("x.txt", BASE_MP) == parse_game_table(
        "x.txt", BASE_MP.replace("\r\n", "\n")
    )


def test_a_row_with_the_wrong_width_is_an_error_not_a_short_row():
    with pytest.raises(GameTableError, match="60"):
        parse_game_table("basemp.txt", BASE_MP.replace("60\t0\t1213\t1512\t0", "60\t0\t1213"))


def test_a_header_with_no_rows_is_an_error():
    with pytest.raises(GameTableError, match="no rows"):
        parse_game_table("basemp.txt", "Level\tMage\r\n")


def test_write_game_tables_pins_the_build_and_covers_the_files(tmp_path: Path):
    build = build_dir_with_manifest(tmp_path)
    calls: list[str] = []
    out = run(tmp_path, calls)
    assert out == build / "gametables"
    assert {path.name for path in out.glob("*.txt")} == set(GAME_TABLE_IDS)
    assert len(casc_calls(calls)) == len(GAME_TABLE_IDS)
    assert all(call.endswith(f"?version={BUILD}") for call in casc_calls(calls)), calls
    assert [call for call in calls if EXPECTED_STAT_PATH in call] == [
        f"https://wago.tools{EXPECTED_STAT_PATH}?build={BUILD}"
    ]
    files = json.loads((build / "manifest.json").read_text())["files"]
    assert "gametables/basemp.txt" in files
    assert f"gametables/{EXPECTED_STAT_FILE}" in files
    assert verify(build) == []


def test_the_expected_stat_table_is_written_as_served_under_its_own_name(tmp_path: Path):
    """The CSV is the client's per-level crit curves; the engine's generator
    reads CritPerAgility and SpellCritPerIntellect from it by name."""
    build_dir_with_manifest(tmp_path)
    out = run(tmp_path, [])
    assert (out / EXPECTED_STAT_FILE).read_text(encoding="utf-8") == EXPECTED_STAT
    assert {path.name for path in out.iterdir()} == set(GAME_TABLE_IDS) | {EXPECTED_STAT_FILE}
    header = EXPECTED_STAT.split("\n", 1)[0].split(",")
    assert set(EXPECTED_STAT_COLUMNS) <= set(header)


def test_an_expected_stat_table_missing_a_crit_column_is_an_error(tmp_path: Path):
    """A build that renames or drops a column the engine reads is a schema
    change to look at, not a file to ship."""
    build_dir_with_manifest(tmp_path)
    renamed = EXPECTED_STAT.replace("CritPerAgility", "Field_1_60_1_99999_007")
    with pytest.raises(GameTableError, match="CritPerAgility"):
        run(tmp_path, [], expected_stat=renamed)
    assert {path.name for path in (tmp_path / BUILD / "gametables").iterdir()} == set(
        GAME_TABLE_IDS
    )


def test_the_bytes_are_written_exactly_as_the_client_ships_them(tmp_path: Path):
    """The engine reads these with csv.reader(delimiter='\\t'); a rewritten
    line ending or a re-serialised float would be our numbers, not the
    client's."""
    build_dir_with_manifest(tmp_path)
    out = run(tmp_path, [])
    assert (out / "basemp.txt").read_bytes() == BASE_MP.encode("utf-8")


def test_a_second_run_serves_from_the_cache_and_re_emits(tmp_path: Path):
    build = build_dir_with_manifest(tmp_path)
    first: list[str] = []
    second: list[str] = []
    run(tmp_path, first)
    run(tmp_path, second)
    assert len(casc_calls(first)) == len(GAME_TABLE_IDS)
    assert casc_calls(second) == []
    assert (build / "gametables" / "combatratings.txt").exists()


def test_an_empty_download_fails_the_run_rather_than_writing_nothing(tmp_path: Path):
    """This is the failure the six absent ids produce, and the reason the
    absent list is a constant rather than something to rediscover."""
    build_dir_with_manifest(tmp_path)
    with pytest.raises(CascError, match=str(GAME_TABLE_IDS["combatratings.txt"])):
        run(tmp_path, [], {GAME_TABLE_IDS["combatratings.txt"]: b""})


def test_a_build_that_was_never_normalized_is_a_clear_error(tmp_path: Path):
    (tmp_path / "1.0.0.0").mkdir()
    client = httpx.Client(transport=transport([]), base_url="https://wago.tools")
    with pytest.raises(SystemExit, match="normalize"):
        write_game_tables("1.0.0.0", root=tmp_path, cache_dir=tmp_path / "cache", client=client)
