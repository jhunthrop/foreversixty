"""wago.tools' hotfix rows layered over the play-session dump (sim-accuracy round 5).

The committed `raw/hotfixes/` dump came from one DBCache read on 1.60.1.70009
and was seeded forward add-only through every build since. Blizzard re-pushed
hundreds of item rows after it: Fiery War Axe 870 went from 2.9 s to 3.6 s,
Blade of Hanna 2801 lost its five stats and went to 2.0 s, and The Stitcher
279874 dropped from rare to uncommon. wago.tools records the pushes it sees for
each build, and serves them applied with `useHotfixes=1`.
"""

from __future__ import annotations

import json
from pathlib import Path

import httpx

from pipeline.hotfix_merge import (
    HOTFIX_STATUS_COLUMN,
    SEEDED_MARKER,
    STATUS_VALID,
    merge_hotfix_table,
)
from pipeline.wago_hotfixes import (
    HOTFIX_BUILD_COLUMN,
    changed_rows,
    fetch_wago_hotfixes,
    hotfix_builds,
    layer_newest,
)


def test_changed_rows_keeps_new_ids_and_edited_rows_only():
    plain = [{"ID": "1", "ItemDelay": "2900"}, {"ID": "2", "ItemDelay": "2000"}]
    hotfixed = [
        {"ID": "1", "ItemDelay": "3600"},
        {"ID": "2", "ItemDelay": "2000"},
        {"ID": "3", "ItemDelay": "1800"},
    ]
    assert changed_rows(plain, hotfixed) == {
        "1": {"ID": "1", "ItemDelay": "3600"},
        "3": {"ID": "3", "ItemDelay": "1800"},
    }


def test_changed_rows_ignores_float_print_width():
    plain = [{"ID": "1", "DmgVariance": "0.4000000059604645"}]
    hotfixed = [{"ID": "1", "DmgVariance": "0.40000000596046"}]
    assert changed_rows(plain, hotfixed) == {}


def test_layer_newest_lets_the_later_build_win_and_tags_the_build():
    layered = layer_newest(
        [
            (
                "1.60.1.70532",
                {"1": {"ID": "1", "ItemDelay": "3600"}, "2": {"ID": "2", "ItemDelay": "2000"}},
            ),
            ("1.60.1.70707", {"1": {"ID": "1", "ItemDelay": "3700"}}),
        ]
    )
    assert layered == [
        {"ID": "1", "ItemDelay": "3700", HOTFIX_BUILD_COLUMN: "1.60.1.70707"},
        {"ID": "2", "ItemDelay": "2000", HOTFIX_BUILD_COLUMN: "1.60.1.70532"},
    ]


def test_hotfix_builds_are_the_products_patch_line_up_to_the_build_oldest_first():
    api = {
        "wowf": [
            {"version": "1.60.1.70720"},
            {"version": "1.60.1.70707"},
            {"version": "1.60.1.70532"},
            {"version": "1.60.0.69000"},
        ],
        "wow_classic_era": [{"version": "1.15.8.1"}],
    }
    assert hotfix_builds(api, "wowf", "1.60.1.70707") == ["1.60.1.70532", "1.60.1.70707"]


def _raw(tmp_path: Path, build: str) -> Path:
    raw = tmp_path / build / "raw"
    (raw / "hotfixes").mkdir(parents=True)
    return raw


def _wago(raw: Path, rows: str) -> None:
    path = raw / "wago-hotfixes" / "T.csv"
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(f"ID,ItemDelay,{HOTFIX_BUILD_COLUMN}\n{rows}", encoding="utf-8")


def test_a_wago_row_replaces_a_seeded_dump_row_the_shipped_table_lacks(tmp_path: Path):
    raw = _raw(tmp_path, "1.60.1.70707")
    (raw / "T.csv").write_text("ID,ItemDelay\n5,1500\n", encoding="utf-8")
    (raw / "hotfixes" / "T.csv").write_text(
        f"ID,ItemDelay,{HOTFIX_STATUS_COLUMN}\n870,2900,{STATUS_VALID}\n", encoding="utf-8"
    )
    (raw / "hotfixes" / SEEDED_MARKER).write_text("1.60.1.70685\n", encoding="utf-8")
    _wago(raw, "870,3600,1.60.1.70532\n")
    assert merge_hotfix_table(raw, "T") == [
        {"ID": "5", "ItemDelay": "1500"},
        {"ID": "870", "ItemDelay": "3600"},
    ]


def test_an_older_builds_wago_row_never_replaces_a_shipped_row(tmp_path: Path):
    raw = _raw(tmp_path, "1.60.1.70707")
    (raw / "T.csv").write_text("ID,ItemDelay\n5,1500\n6,1600\n", encoding="utf-8")
    _wago(raw, "5,9999,1.60.1.70532\n6,1700,1.60.1.70707\n")
    assert merge_hotfix_table(raw, "T") == [
        {"ID": "5", "ItemDelay": "1500"},
        {"ID": "6", "ItemDelay": "1700"},
    ]


def test_this_builds_own_session_dump_beats_wago(tmp_path: Path):
    raw = _raw(tmp_path, "1.60.1.70707")
    (raw / "T.csv").write_text("ID,ItemDelay\n", encoding="utf-8")
    (raw / "hotfixes" / "T.csv").write_text(
        f"ID,ItemDelay,{HOTFIX_STATUS_COLUMN}\n870,3800,{STATUS_VALID}\n", encoding="utf-8"
    )
    _wago(raw, "870,3600,1.60.1.70707\n")
    assert merge_hotfix_table(raw, "T") == [{"ID": "870", "ItemDelay": "3800"}]


def test_fetch_wago_hotfixes_writes_the_layered_rows(tmp_path: Path):
    tables = {
        ("1.60.1.70532", False): "ID,ItemDelay\n1,2900\n",
        ("1.60.1.70532", True): "ID,ItemDelay\n1,3600\n2,2000\n",
        ("1.60.1.70707", False): "ID,ItemDelay\n1,2900\n",
        ("1.60.1.70707", True): "ID,ItemDelay\n1,2900\n",
    }
    builds = {"wowf": [{"version": "1.60.1.70707"}, {"version": "1.60.1.70532"}]}

    def handler(request: httpx.Request) -> httpx.Response:
        if request.url.path == "/api/builds":
            return httpx.Response(200, text=json.dumps(builds))
        build = request.url.params["build"]
        hot = request.url.params.get("useHotfixes") == "1"
        return httpx.Response(200, text=tables[(build, hot)])

    client = httpx.Client(transport=httpx.MockTransport(handler), base_url="https://wago.tools")
    written = fetch_wago_hotfixes("1.60.1.70707", "wowf", tmp_path, client=client, tables=("T",))
    assert written == {"T": 2}
    out = (tmp_path / "1.60.1.70707" / "raw" / "wago-hotfixes" / "T.csv").read_text(
        encoding="utf-8"
    )
    assert out.splitlines() == [
        f"ID,ItemDelay,{HOTFIX_BUILD_COLUMN}",
        "1,3600,1.60.1.70532",
        "2,2000,1.60.1.70532",
    ]
