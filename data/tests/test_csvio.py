import logging
from pathlib import Path

import pytest

from pipeline.csvio import check_item_sparse_completeness, read_csv

FIXTURES = Path(__file__).parent / "fixtures"


def test_read_csv_returns_rows_keyed_by_header():
    rows = read_csv(FIXTURES / "Tiny.csv")
    assert rows == [
        {"ID": "1", "Name_lang": "Alpha", "Flags": "0"},
        {"ID": "2", "Name_lang": "Beta, with comma", "Flags": "4"},
    ]


def _write_item_tables(raw: Path, *, item_rows: int, sparse_rows: int) -> None:
    raw.mkdir(parents=True, exist_ok=True)
    (raw / "Item.csv").write_text(
        "ID,ClassID\n" + "".join(f"{i},4\n" for i in range(item_rows)), encoding="utf-8"
    )
    (raw / "ItemSparse.csv").write_text(
        "ID,Display_lang\n" + "".join(f"{i},Item {i}\n" for i in range(sparse_rows)),
        encoding="utf-8",
    )


def test_check_item_sparse_completeness_passes_when_the_ratio_is_met(tmp_path: Path):
    _write_item_tables(tmp_path, item_rows=100, sparse_rows=90)
    check_item_sparse_completeness(tmp_path)  # 90/100 == the 0.9 floor; no raise


def test_check_item_sparse_completeness_raises_on_a_truncated_export(tmp_path: Path):
    # The real incident: wago.tools served 19,226 of Item.csv's 31,819 rows (60%).
    _write_item_tables(tmp_path, item_rows=100, sparse_rows=60)
    with pytest.raises(SystemExit, match="ItemSparse.csv has only 60 rows"):
        check_item_sparse_completeness(tmp_path)


def test_check_item_sparse_completeness_allow_shrink_skips_and_logs_loudly(
    tmp_path: Path, caplog
):
    caplog.set_level(logging.WARNING)
    _write_item_tables(tmp_path, item_rows=100, sparse_rows=60)
    check_item_sparse_completeness(tmp_path, allow_shrink=True)  # does not raise
    assert "ItemSparse completeness gate skipped (--allow-shrink)" in caplog.text


def test_check_item_sparse_completeness_does_not_divide_by_zero(tmp_path: Path):
    """An empty Item.csv is some other check's problem, not this one's."""
    _write_item_tables(tmp_path, item_rows=0, sparse_rows=0)
    check_item_sparse_completeness(tmp_path)
