import struct
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.db2 import Column
from pipeline.hotfix_cache import (
    HotfixCacheError,
    HotfixStatus,
    decode_table_hotfixes,
    read_cache,
)
from pipeline.hotfix_merge import (
    HOTFIX_STATUS_COLUMN,
    SEEDED_MARKER,
    STATUS_REMOVED,
    STATUS_VALID,
    hotfix_csv_rows,
    merge_hotfix_rows,
    merge_hotfix_table,
    write_build_hotfix_tables,
    write_hotfix_csv,
)

TINY_SCHEMA: list[Column] = [
    Column("Name_lang", "str"),
    Column("Level", "i", bits=8),
]


# --- merge_hotfix_rows: override / add / remove -----------------------------


def test_merge_hotfix_rows_overrides_a_shipped_id():
    shipped = [{"ID": "1", "Name_lang": "Old", "Level": "0"}]
    hotfix = [{"ID": "1", "Name_lang": "New", "Level": "9", HOTFIX_STATUS_COLUMN: STATUS_VALID}]

    merged = merge_hotfix_rows(shipped, hotfix)

    assert merged == [{"ID": "1", "Name_lang": "New", "Level": "9"}]


def test_merge_hotfix_rows_adds_an_id_the_shipped_table_never_had():
    shipped = [{"ID": "1", "Name_lang": "Alpha", "Level": "1"}]
    hotfix = [{"ID": "2", "Name_lang": "Beta", "Level": "2", HOTFIX_STATUS_COLUMN: STATUS_VALID}]

    merged = merge_hotfix_rows(shipped, hotfix)

    assert merged == [
        {"ID": "1", "Name_lang": "Alpha", "Level": "1"},
        {"ID": "2", "Name_lang": "Beta", "Level": "2"},
    ]


def test_merge_hotfix_rows_removed_status_drops_a_shipped_row():
    shipped = [
        {"ID": "1", "Name_lang": "Alpha", "Level": "1"},
        {"ID": "2", "Name_lang": "Beta", "Level": "2"},
    ]
    hotfix = [{"ID": "1", HOTFIX_STATUS_COLUMN: STATUS_REMOVED}]

    merged = merge_hotfix_rows(shipped, hotfix)

    assert merged == [{"ID": "2", "Name_lang": "Beta", "Level": "2"}]


def test_merge_hotfix_rows_removed_status_on_an_id_shipped_never_had_is_a_no_op():
    shipped = [{"ID": "1", "Name_lang": "Alpha", "Level": "1"}]
    hotfix = [{"ID": "99", HOTFIX_STATUS_COLUMN: STATUS_REMOVED}]

    merged = merge_hotfix_rows(shipped, hotfix)

    assert merged == shipped


def test_merge_hotfix_rows_result_never_carries_the_status_column():
    shipped: list[dict[str, str]] = []
    hotfix = [{"ID": "1", "Name_lang": "New", HOTFIX_STATUS_COLUMN: STATUS_VALID}]

    merged = merge_hotfix_rows(shipped, hotfix)

    assert HOTFIX_STATUS_COLUMN not in merged[0]


def test_merge_hotfix_rows_orders_the_result_by_id_numerically():
    shipped = [{"ID": "10", "Name_lang": "Ten"}]
    hotfix = [{"ID": "2", "Name_lang": "Two", HOTFIX_STATUS_COLUMN: STATUS_VALID}]

    merged = merge_hotfix_rows(shipped, hotfix)

    assert [row["ID"] for row in merged] == ["2", "10"]  # not ["10", "2"] (string sort)


# --- merge_hotfix_table: passthrough / csv round-trip -----------------------


def test_merge_hotfix_rows_add_only_adds_missing_ids_and_never_overrides_or_removes():
    """A dump seeded from another build (SEEDED_MARKER) may only add the ids
    the shipped table lacks: a shipped row keeps the new client's values and a
    removed-status id stays."""
    shipped = [{"ID": "1", "Display_lang": "Shipped One"}, {"ID": "3", "Display_lang": "Three"}]
    hotfix = [
        {"ID": "1", "Display_lang": "Old Hotfix One", HOTFIX_STATUS_COLUMN: STATUS_VALID},
        {"ID": "2", "Display_lang": "Hotfix Only Two", HOTFIX_STATUS_COLUMN: STATUS_VALID},
        {"ID": "3", HOTFIX_STATUS_COLUMN: STATUS_REMOVED},
    ]
    merged = merge_hotfix_rows(shipped, hotfix, add_only=True)
    assert merged == [
        {"ID": "1", "Display_lang": "Shipped One"},
        {"ID": "2", "Display_lang": "Hotfix Only Two"},
        {"ID": "3", "Display_lang": "Three"},
    ]


def test_merge_hotfix_table_merges_add_only_when_the_dump_is_seeded(tmp_path: Path):
    raw = tmp_path / "raw"
    (raw / "hotfixes").mkdir(parents=True)
    (raw / "T.csv").write_text("ID,Display_lang\n1,Shipped One\n", encoding="utf-8")
    (raw / "hotfixes" / "T.csv").write_text(
        f"ID,Display_lang,{HOTFIX_STATUS_COLUMN}\n1,Old One,{STATUS_VALID}\n2,Two,{STATUS_VALID}\n",
        encoding="utf-8",
    )
    assert merge_hotfix_table(raw, "T") == [
        {"ID": "1", "Display_lang": "Old One"},
        {"ID": "2", "Display_lang": "Two"},
    ]
    (raw / "hotfixes" / SEEDED_MARKER).write_text("9.9.9.1\n", encoding="utf-8")
    assert merge_hotfix_table(raw, "T") == [
        {"ID": "1", "Display_lang": "Shipped One"},
        {"ID": "2", "Display_lang": "Two"},
    ]


def test_merge_hotfix_table_is_a_passthrough_with_no_hotfix_file(tmp_path: Path):
    raw = tmp_path / "raw"
    raw.mkdir()
    (raw / "ItemSparse.csv").write_text("ID,Name_lang\n1,Alpha\n", encoding="utf-8")

    rows = merge_hotfix_table(raw, "ItemSparse")

    assert rows == [{"ID": "1", "Name_lang": "Alpha"}]


def test_merge_hotfix_table_merges_a_real_hotfix_csv_on_disk(tmp_path: Path):
    raw = tmp_path / "raw"
    raw.mkdir()
    (raw / "ItemSparse.csv").write_text("ID,Name_lang,Level\n1,Old,0\n2,Keep,5\n", encoding="utf-8")
    hotfix_dir = raw / "hotfixes"
    hotfix_dir.mkdir()
    (hotfix_dir / "ItemSparse.csv").write_text(
        f"ID,Name_lang,Level,{HOTFIX_STATUS_COLUMN}\n"
        f"1,New,9,{STATUS_VALID}\n"
        f"2,,,{STATUS_REMOVED}\n"
        f"3,Added,3,{STATUS_VALID}\n",
        encoding="utf-8",
    )

    rows = merge_hotfix_table(raw, "ItemSparse")

    assert rows == [
        {"ID": "1", "Name_lang": "New", "Level": "9"},
        {"ID": "3", "Name_lang": "Added", "Level": "3"},
    ]


# --- hotfix_csv_rows / write_hotfix_csv: TableHotfixes -> csv -> read_csv ---


def test_hotfix_csv_rows_and_write_hotfix_csv_round_trip_through_read_csv(tmp_path: Path):
    _HEADER_FMT = "<4sII32s"
    _RECORD_FMT = "<4siiIIIIB3x"

    def encode(name: str, level: int) -> bytes:
        return name.encode("utf-8") + b"\x00" + int(level).to_bytes(1, "little", signed=True)

    def record(table_hash, record_id, status, data=b"") -> bytes:
        header = struct.pack(
            _RECORD_FMT, b"XFTH", 0, 1, 1, table_hash, record_id, len(data), int(status)
        )
        return header + data

    records_bytes = record(0xAAAA, 1, HotfixStatus.VALID, encode("Winner", 5)) + record(
        0xAAAA, 2, HotfixStatus.RECORD_REMOVED
    )
    cache_path = tmp_path / "DBCache.bin"
    cache_bytes = struct.pack(_HEADER_FMT, b"XFTH", 9, 70009, b"\x00" * 32) + records_bytes
    cache_path.write_bytes(cache_bytes)
    header, records = read_cache(cache_path)
    hotfixes = decode_table_hotfixes(header, records, 0xAAAA, TINY_SCHEMA)

    rows = hotfix_csv_rows(hotfixes)
    csv_path = tmp_path / "ItemSparse.csv"
    write_hotfix_csv(rows, csv_path)
    read_back = read_csv(csv_path)

    assert read_back == [
        {"ID": "1", "Name_lang": "Winner", "Level": "5", HOTFIX_STATUS_COLUMN: STATUS_VALID},
        {"ID": "2", "Name_lang": "", "Level": "", HOTFIX_STATUS_COLUMN: STATUS_REMOVED},
    ]


# --- write_build_hotfix_tables: the CLI's own entry point -------------------


def _write_real_cache(tmp_path: Path, *, build: int = 70009) -> Path:
    from pipeline.hotfix_cache import ITEM_SPARSE_TABLE_HASH, ITEM_TABLE_HASH

    _HEADER_FMT = "<4sII32s"
    _RECORD_FMT = "<4siiIIIIB3x"

    def encode(name: str, level: int) -> bytes:
        return name.encode("utf-8") + b"\x00" + int(level).to_bytes(1, "little", signed=True)

    # These use TINY_SCHEMA's shape, not the real ItemSparse/Item schemas --
    # write_build_hotfix_tables always decodes with the real schemas, so this
    # helper is only used by the build-mismatch test below, which never gets
    # far enough to decode a row.
    def record(table_hash, record_id, status, data=b"") -> bytes:
        header = struct.pack(
            _RECORD_FMT, b"XFTH", 0, 1, 1, table_hash, record_id, len(data), int(status)
        )
        return header + data

    records_bytes = record(ITEM_SPARSE_TABLE_HASH, 1, HotfixStatus.RECORD_REMOVED) + record(
        ITEM_TABLE_HASH, 1, HotfixStatus.RECORD_REMOVED
    )
    path = tmp_path / "DBCache.bin"
    path.write_bytes(struct.pack(_HEADER_FMT, b"XFTH", 9, build, b"\x00" * 32) + records_bytes)
    return path


def test_write_build_hotfix_tables_rejects_a_build_mismatch(tmp_path: Path):
    cache_path = _write_real_cache(tmp_path, build=70009)

    with pytest.raises(HotfixCacheError, match="build 70009"):
        write_build_hotfix_tables("1.60.1.99999", cache_path, root=tmp_path)


def test_write_build_hotfix_tables_writes_both_tables_and_reports_removed(tmp_path: Path):
    cache_path = _write_real_cache(tmp_path, build=70009)
    build_dir = tmp_path / "1.60.1.70009"
    (build_dir / "raw").mkdir(parents=True)

    results = write_build_hotfix_tables("1.60.1.70009", cache_path, root=tmp_path)

    by_table = {r.table: r for r in results}
    assert set(by_table) == {"ItemSparse", "Item"}
    assert by_table["ItemSparse"].removed == 1
    assert by_table["ItemSparse"].valid == 0
    assert by_table["ItemSparse"].path == build_dir / "raw" / "hotfixes" / "ItemSparse.csv"
    assert by_table["ItemSparse"].path.exists()
