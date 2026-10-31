# data/tests/test_build_floors.py
"""The cross-build floors (design/reviews/2026-10-30-70707-dps-drop.md).

The shrink gates compared a build only with its own last run, so a first run always
passed: 70685 lost 313 items and 681 of its 688 wowhead-sourced items with every
gate green. A first run is now compared with the newest other build, and the
wowhead-sourced count has a floor of its own. Floors, never exact counts.
"""

import json
from pathlib import Path

import pytest

from pipeline import build_floors
from pipeline.csvio import check_item_sparse_completeness, check_wowhead_items_floor
from pipeline.normalize import previous_class_item_counts


def _items(build_dir: Path, total: int, wowhead: int = 0) -> None:
    build_dir.mkdir(parents=True, exist_ok=True)
    rows = [
        {"id": i, "stats_source": "wowhead" if i < wowhead else None} for i in range(total)
    ]
    (build_dir / "items.json").write_text(json.dumps(rows))


def _class_items(build_dir: Path, **counts: int) -> None:
    (build_dir / "items").mkdir(parents=True, exist_ok=True)
    for slug, count in counts.items():
        payload = {"items": [{"id": i} for i in range(count)]}
        (build_dir / "items" / f"{slug}.json").write_text(json.dumps(payload))


def test_the_previous_build_is_the_newest_other_build_by_version(tmp_path: Path) -> None:
    for name in ("1.60.1.9", "1.60.1.10", "1.15.9.99999", "forever-prebeta", "1.60.1.11"):
        _items(tmp_path / name, 1)
    (tmp_path / "1.60.1.12").mkdir()  # no items.json: not a reference
    found = build_floors.previous_build_dir(tmp_path / "1.60.1.11", "items.json")
    assert found == tmp_path / "1.60.1.10"
    assert build_floors.previous_build_dir(tmp_path / "1.60.1.13", "items.json") == (
        tmp_path / "1.60.1.11"
    )
    assert build_floors.previous_build_dir(tmp_path / "x" / "1.60.1.1", "items.json") is None


def test_a_first_run_is_compared_with_the_previous_build(tmp_path: Path) -> None:
    _items(tmp_path / "1.60.1.1", 1000)
    new = tmp_path / "1.60.1.2"
    new.mkdir()
    check_item_sparse_completeness(new, 980)  # a 2% loss is a smaller client
    with pytest.raises(SystemExit, match="1.60.1.1"):
        check_item_sparse_completeness(new, 979)
    check_item_sparse_completeness(new, 979, allow_shrink=True)


def test_a_builds_own_last_run_still_may_not_lose_one_item(tmp_path: Path) -> None:
    _items(tmp_path / "1.60.1.1", 10)
    _items(tmp_path / "1.60.1.2", 1000)
    with pytest.raises(SystemExit, match="1000 to 999"):
        check_item_sparse_completeness(tmp_path / "1.60.1.2", 999)


def test_the_very_first_build_has_nothing_to_compare_with(tmp_path: Path) -> None:
    (tmp_path / "1.60.1.1").mkdir()
    check_item_sparse_completeness(tmp_path / "1.60.1.1", 1)
    check_wowhead_items_floor(tmp_path / "1.60.1.1", 0)


def test_the_wowhead_sourced_count_has_a_floor_against_the_previous_build(
    tmp_path: Path,
) -> None:
    _items(tmp_path / "1.60.1.70658", 25678, wowhead=688)
    new = tmp_path / "1.60.1.70685"
    new.mkdir()
    check_wowhead_items_floor(new, 620)  # a floor, not an exact count: 90% of 688
    with pytest.raises(SystemExit, match="7 .*688"):
        check_wowhead_items_floor(new, 7)  # 70685's own figure
    check_wowhead_items_floor(new, 7, allow_shrink=True)


def test_the_wowhead_floor_holds_against_the_builds_own_last_run(tmp_path: Path) -> None:
    _items(tmp_path / "1.60.1.70707", 25534, wowhead=688)
    with pytest.raises(SystemExit, match="396"):
        check_wowhead_items_floor(tmp_path / "1.60.1.70707", 396)


def test_class_items_fall_back_to_the_previous_builds_files(tmp_path: Path) -> None:
    _class_items(tmp_path / "1.60.1.1", mage=4000, rogue=3000)
    (tmp_path / "1.60.1.2").mkdir()
    assert previous_class_item_counts(tmp_path / "1.60.1.2") == {"mage": 4000, "rogue": 3000}
    _class_items(tmp_path / "1.60.1.2", mage=10)
    assert previous_class_item_counts(tmp_path / "1.60.1.2") == {"mage": 10}


def test_floor_problems_names_both_floors(tmp_path: Path) -> None:
    _items(tmp_path / "1.60.1.70658", 25678, wowhead=688)
    _items(tmp_path / "1.60.1.70707", 25534, wowhead=396)
    _items(tmp_path / "1.60.1.70800", 25700, wowhead=700)
    problems = build_floors.floor_problems(tmp_path / "1.60.1.70707", tmp_path / "1.60.1.70658")
    assert len(problems) == 1 and "396" in problems[0] and "688" in problems[0]
    assert build_floors.floor_problems(tmp_path / "1.60.1.70800", tmp_path / "1.60.1.70658") == []


def test_the_build_floors_command_refuses_a_build_below_the_active_ones(
    tmp_path: Path, monkeypatch: pytest.MonkeyPatch, caplog: pytest.LogCaptureFixture
) -> None:
    """tools/activate-build.sh runs it before a move: 70707's committed items took
    396 items' stats from wowhead against the active 70658's 688."""
    from pipeline.__main__ import main

    _items(tmp_path / "builds" / "1.60.1.70658", 25678, wowhead=688)
    _items(tmp_path / "builds" / "1.60.1.70707", 25534, wowhead=396)
    _items(tmp_path / "builds" / "1.60.1.70800", 25700, wowhead=700)
    monkeypatch.chdir(tmp_path)
    assert main(["build-floors", "--build", "1.60.1.70707", "--against", "1.60.1.70658"]) == 1
    assert "396" in caplog.text
    assert main(["build-floors", "--build", "1.60.1.70800", "--against", "1.60.1.70658"]) == 0
