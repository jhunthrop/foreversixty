"""The server's item push read from wago.tools' hotfix page (issue #27).

Between the 1.60.1.70790 and 70938 data runs wago's hotfixed CSV export stopped
applying push 16778086, and Fiery War Axe (870) went back to a 2.9 s swing. The
page still lists the push: its enUS copies carry no values, its deDE, frFR,
koKR and zhTW copies carry ItemDelay 3600. The fixtures are the page's real
70938 entries for 870 and 2801 and the committed play-session rows under them.
"""

from __future__ import annotations

import csv
import json
import shutil
from pathlib import Path

import pytest

from pipeline.hotfix_merge import SEEDED_MARKER, merge_hotfix_table
from pipeline.hotfix_push import STATUS_REMOVED, STATUS_VALID, PushEntry, newest_per_record
from pipeline.hotfix_push_store import read_store
from pipeline.item_push import ITEM_PUSH_TABLES, fetch_item_push, item_store, layer_item_push
from tests.test_hotfix_push_store import Wago

FIXTURES = Path(__file__).parent / "fixtures" / "item_push"
BUILD = "1.60.1.70938"


def page_entries() -> list[dict]:
    return json.loads((FIXTURES / "push-entries-70938.json").read_text(encoding="utf-8"))


@pytest.fixture
def raw(tmp_path: Path) -> Path:
    """A 70938 raw/ with one shipped row, the seeded play-session rows for 870
    and 2801 (870 at 2.9 s, 2801 at 2.1 s with five stats) and no CSV export."""
    raw = tmp_path / BUILD / "raw"
    (raw / "hotfixes").mkdir(parents=True)
    shutil.copyfile(FIXTURES / "ItemSparse.csv", raw / "ItemSparse.csv")
    shutil.copyfile(FIXTURES / "hotfixes-ItemSparse.csv", raw / "hotfixes" / "ItemSparse.csv")
    (raw / "hotfixes" / SEEDED_MARKER).write_text("1.60.1.70790\n", encoding="utf-8")
    return raw


def crawl(raw: Path) -> dict[str, int]:
    return fetch_item_push(BUILD, root=raw.parent.parent, client=Wago(page_entries()).client())


def by_id(rows: list[dict[str, str]]) -> dict[str, dict[str, str]]:
    return {row["ID"]: row for row in rows}


def test_a_valid_entry_without_values_never_outranks_an_older_one_with_values():
    with_values = PushEntry(1, 100, 870, STATUS_VALID, [870, "3600"], "deDE")
    empty = PushEntry(2, 100, 870, STATUS_VALID, None)
    assert newest_per_record([with_values, empty]) == {870: with_values}


def test_a_removal_or_a_newer_push_with_values_still_wins():
    old = PushEntry(1, 100, 870, STATUS_VALID, [870, "2900"])
    removed = PushEntry(2, 101, 870, STATUS_REMOVED, None)
    newer = PushEntry(3, 102, 870, STATUS_VALID, [870, "3600"], "frFR")
    assert newest_per_record([old, removed]) == {870: removed}
    assert newest_per_record([old, removed, newer]) == {870: newer}


def test_with_no_entry_carrying_values_the_newest_still_stands():
    first = PushEntry(1, 100, 870, STATUS_VALID, None)
    second = PushEntry(2, 100, 870, STATUS_VALID, None)
    assert newest_per_record([first, second]) == {870: second}


def test_the_crawl_keeps_every_locale_and_saves_one_deciding_entry_per_record(raw):
    assert crawl(raw) == {"ItemSparse": 2}
    saved = read_store(item_store(raw))
    assert saved is not None and saved["tables"] == list(ITEM_PUSH_TABLES)
    entries = {e.record_id: e for e in saved["entries"]["ItemSparse"]}
    assert set(entries) == {870, 2801}
    assert all(e.data and e.locale != "enUS" for e in entries.values())


def test_fiery_war_axe_swings_at_the_pushed_3_6_s_with_its_english_name(raw):
    crawl(raw)
    row = by_id(merge_hotfix_table(raw, "ItemSparse"))["870"]
    assert row["ItemDelay"] == "3600"
    assert row["Display_lang"] == "Fiery War Axe"


def test_blade_of_hanna_takes_the_pushed_2_0_s_and_drops_its_five_stats(raw):
    before = by_id(merge_hotfix_table(raw, "ItemSparse"))["2801"]
    assert before["ItemDelay"] == "2100" and before["StatModifier_bonusStat_0"] == "3"
    crawl(raw)
    row = by_id(merge_hotfix_table(raw, "ItemSparse"))["2801"]
    assert row["ItemDelay"] == "2000"
    assert all(row[f"StatModifier_bonusStat_{n}"] == "-1" for n in range(10))
    assert row["Display_lang"] == "Blade of Hanna"


def test_the_push_goes_over_the_csv_exports_row(raw):
    crawl(raw)
    export = raw / "wago-hotfixes" / "ItemSparse.csv"
    export.parent.mkdir()
    rows = list(csv.DictReader((raw / "hotfixes" / "ItemSparse.csv").open(encoding="utf-8")))
    fields = [k for k in rows[0] if not k.startswith("_")] + ["_HotfixBuild"]
    with export.open("w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=fields, extrasaction="ignore")
        writer.writeheader()
        writer.writerows({**r, "ItemDelay": "2900", "_HotfixBuild": BUILD} for r in rows)
    assert by_id(merge_hotfix_table(raw, "ItemSparse"))["870"]["ItemDelay"] == "3600"


def test_a_build_without_a_saved_item_push_passes_through(raw):
    rows = merge_hotfix_table(raw, "ItemSparse")
    assert layer_item_push(rows, list(rows[0]), raw, "ItemSparse") == rows
    assert by_id(rows)["870"]["ItemDelay"] == "2900"


def test_another_builds_saved_item_push_is_refused(raw, tmp_path):
    crawl(raw)
    other = tmp_path / "1.60.1.70999" / "raw"
    other.mkdir(parents=True)
    shutil.copytree(raw / "wago-push", other / "wago-push")
    with pytest.raises(SystemExit, match="is the push of 1.60.1.70938"):
        layer_item_push([], [], other, "ItemSparse")
