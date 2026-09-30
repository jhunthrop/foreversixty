import json
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


def _commit_items(build_dir: Path, count: int) -> None:
    build_dir.mkdir(parents=True, exist_ok=True)
    (build_dir / "items.json").write_text(
        json.dumps([{"id": i} for i in range(count)]), encoding="utf-8"
    )


def test_check_item_sparse_completeness_passes_when_the_count_holds(tmp_path: Path):
    _commit_items(tmp_path, 100)
    check_item_sparse_completeness(tmp_path, 100)  # no shrink; no raise
    check_item_sparse_completeness(tmp_path, 150)  # a growth; no raise


def test_check_item_sparse_completeness_raises_on_a_shrink(tmp_path: Path):
    _commit_items(tmp_path, 100)
    with pytest.raises(SystemExit, match="would shrink from 100 to 60"):
        check_item_sparse_completeness(tmp_path, 60)


def test_check_item_sparse_completeness_allow_shrink_skips_and_logs_loudly(
    tmp_path: Path, caplog
):
    caplog.set_level(logging.WARNING)
    _commit_items(tmp_path, 100)
    check_item_sparse_completeness(tmp_path, 60, allow_shrink=True)  # does not raise
    assert "ItemSparse completeness gate skipped (--allow-shrink)" in caplog.text


def test_check_item_sparse_completeness_first_ever_run_has_nothing_to_shrink_against(
    tmp_path: Path,
):
    """No committed `items.json` yet (a new build, or the first build ever
    normalized) is never a shrink -- same rule as
    `normalize._check_class_items_not_shrunk`'s `previous_count == 0`."""
    check_item_sparse_completeness(tmp_path, 0)


def test_check_item_sparse_completeness_an_empty_committed_file_is_also_never_a_shrink(
    tmp_path: Path,
):
    _commit_items(tmp_path, 0)
    check_item_sparse_completeness(tmp_path, 0)
