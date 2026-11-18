import shutil
from dataclasses import dataclass
from pathlib import Path

import pytest

from pipeline.blizzard_cdn import BuildConfigs
from pipeline.cdn_tables import write_table
from pipeline.client_tables import FILE_DATA_IDS, OPTIONAL_TABLES, TABLES
from pipeline.dbd import DBD_COMMIT
from tests.cdn_fakes import FakeSource

FIXTURES = Path(__file__).parent / "fixtures" / "cdn"


@dataclass
class _Client:
    cache: Path


class FakeBuild(FakeSource):
    def __init__(self, cache: Path, files: dict[int, bytes]) -> None:
        super().__init__(files)
        self.client = _Client(cache)
        self.configs = BuildConfigs("1.60.1.70938", "bc", "cc")


def build_with(tmp_path: Path, tables: list[str]) -> FakeBuild:
    dbd = tmp_path / "cache" / "dbd" / DBD_COMMIT
    dbd.mkdir(parents=True)
    for table in tables:
        shutil.copy(FIXTURES / f"{table}.dbd", dbd / f"{table}.dbd")
    files = {FILE_DATA_IDS[t]: (FIXTURES / f"{t}.db2").read_bytes() for t in tables}
    return FakeBuild(tmp_path / "cache", files)


def test_every_fetched_table_has_a_file_data_id():
    assert set(TABLES) <= set(FILE_DATA_IDS)
    assert OPTIONAL_TABLES <= set(TABLES)


def test_a_table_is_written_in_the_export_shape(tmp_path: Path):
    build = build_with(tmp_path, ["ArmorLocation"])
    path = write_table("ArmorLocation", build, tmp_path / "raw", http=None)  # type: ignore[arg-type]
    assert path.read_text() == (FIXTURES / "ArmorLocation.wago-70532.csv").read_text()


def test_a_missing_optional_table_is_empty_and_a_missing_required_one_fails(tmp_path: Path):
    build = build_with(tmp_path, [])
    path = write_table("JournalInstance", build, tmp_path / "raw", http=None)  # type: ignore[arg-type]
    assert path.read_text() == "ID\n"
    with pytest.raises(SystemExit, match="SpellName"):
        write_table("SpellName", build, tmp_path / "raw", http=None)  # type: ignore[arg-type]
