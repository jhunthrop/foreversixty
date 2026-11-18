import csv
import struct
from pathlib import Path

import httpx
import pytest

from pipeline.dbd import Field
from pipeline.hotfix_push import STATUS_REMOVED, STATUS_VALID, PushEntry
from pipeline.hotfix_push_store import write_store
from pipeline.raidbots_hotfixes import (
    STORE,
    STORE_DIR,
    HotfixDecodeError,
    decode_record,
    download,
    layer,
    save_push,
)
from pipeline.wago_http import USER_AGENT

BUILD = "1.60.1.70959"


def field(name, kind="int", bits=32, signed=True, count=1, is_id=False, inline=True, rel=False):
    return Field(name, kind, bits, signed, count, is_id, inline, rel)


FIELDS = [
    field("ID", is_id=True, inline=False),
    field("Name_lang", kind="locstring"),
    field("Speed", bits=16, signed=False),
    field("Ratio", kind="float"),
    field("Pair", count=2),
    field("ParentID", inline=False, rel=True),
]


def test_a_pushed_record_decodes_in_column_order():
    data = b"Fiery War Axe\0" + struct.pack("<Hf2ii", 3600, 0.4, -1, 7, 870)
    assert decode_record(data, 870, FIELDS) == [
        "870",
        "Fiery War Axe",
        "3600",
        "0.40000000596",
        "-1",
        "7",
        "870",
    ]


def test_a_record_that_does_not_fit_the_definition_is_an_error():
    with pytest.raises(HotfixDecodeError, match="870"):
        decode_record(b"Axe\0" + bytes(30), 870, FIELDS)


def test_the_cache_is_downloaded_only_when_raidbots_changed_it(tmp_path: Path, monkeypatch):
    monkeypatch.delenv("FOREVER_DBCACHE", raising=False)
    seen: list[dict[str, str]] = []

    def handler(request: httpx.Request) -> httpx.Response:
        seen.append(dict(request.headers))
        if request.headers.get("if-none-match") == '"v1"':
            return httpx.Response(304)
        return httpx.Response(
            200, content=b"XFTH", headers={"etag": '"v1"', "last-modified": "Wed, 18 Nov 2026"}
        )

    http = httpx.Client(transport=httpx.MockTransport(handler))
    first = download(tmp_path, http)
    second = download(tmp_path, http)
    assert first == second and first.read_bytes() == b"XFTH"
    assert [h.get("user-agent") for h in seen] == [USER_AGENT, USER_AGENT]
    assert "if-none-match" not in seen[0] and seen[1]["if-none-match"] == '"v1"'


def test_a_cache_for_another_build_saves_nothing(tmp_path: Path):
    cache = tmp_path / "DBCache.bin"
    cache.write_bytes(b"XFTH" + struct.pack("<II", 9, 70959) + bytes(32))
    assert save_push("1.60.1.70938", tmp_path, cache_file=cache) is None
    assert not (tmp_path / "1.60.1.70938").exists()


def write_table(path: Path, header: list[str], rows: list[list[str]]) -> None:
    with path.open("w", newline="", encoding="utf-8") as f:
        writer = csv.writer(f, lineterminator="\n")
        writer.writerow(header)
        writer.writerows(rows)


def test_the_saved_push_is_layered_over_the_spell_side_and_written_for_the_items(tmp_path: Path):
    raw = tmp_path / BUILD / "raw"
    raw.mkdir(parents=True)
    write_table(raw / "SpellName.csv", ["ID", "Name_lang"], [["1", "Old"], ["2", "Gone"]])
    write_table(raw / "ItemSparse.csv", ["ID", "Display_lang"], [["870", "Fiery War Axe"]])
    write_table(raw / "Item.csv", ["ID", "ClassID"], [["870", "2"]])
    write_store(
        raw / STORE_DIR / STORE,
        BUILD,
        ["SpellName", "ItemSparse", "Item"],
        {},
        {
            "SpellName": [
                PushEntry(1, 113000, 1, STATUS_VALID, ["1", "New"]),
                PushEntry(2, 113000, 2, STATUS_REMOVED, None),
            ],
            "ItemSparse": [PushEntry(3, 16778086, 870, STATUS_VALID, ["870", "Fiery War Axe"])],
            "Item": [],
        },
    )
    summary = layer(BUILD, tmp_path)
    assert summary is not None
    assert summary["tables"]["SpellName"]["changed"] == 1
    assert summary["tables"]["SpellName"]["removed"] == 1
    assert (raw / "SpellName.csv").read_text() == "ID,Name_lang\n1,New\n"
    items = (raw / "wago-hotfixes" / "ItemSparse.csv").read_text()
    assert items == f"ID,Display_lang,_HotfixBuild\n870,Fiery War Axe,{BUILD}\n"
    assert summary["item_hotfixes"] == {"ItemSparse": 1, "Item": 0}


def test_a_build_with_no_saved_push_layers_nothing(tmp_path: Path):
    (tmp_path / BUILD / "raw").mkdir(parents=True)
    assert layer(BUILD, tmp_path) is None
