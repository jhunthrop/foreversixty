import json
from pathlib import Path

import httpx
import pytest

from pipeline.wago import TABLES, USER_AGENT, download_table, fetch_build, latest_build


def fake_transport(
    calls: list[str], missing_tables: frozenset[str] = frozenset()
) -> httpx.MockTransport:
    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(str(request.url))
        if request.url.path == "/api/builds":
            return httpx.Response(
                200,
                json={
                    "wow_classic_era": [
                        {"version": "1.15.7.61582"},
                        {"version": "1.15.6.60000"},
                    ]
                },
            )
        if request.url.path.startswith("/db2/"):
            table = request.url.path.split("/")[2]
            if table in missing_tables:
                return httpx.Response(404, json={"errors": "Table not found."})
            return httpx.Response(200, text=f"ID,Name_lang\n1,{table}\n")
        return httpx.Response(404)

    return httpx.MockTransport(handler)


def test_latest_build_picks_first_entry():
    client = httpx.Client(transport=fake_transport([]), base_url="https://wago.tools")
    assert latest_build("wow_classic_era", client) == "1.15.7.61582"


def test_fetch_build_downloads_every_table(tmp_path: Path):
    calls: list[str] = []
    client = httpx.Client(transport=fake_transport(calls), base_url="https://wago.tools")
    out = fetch_build("wow_classic_era", None, root=tmp_path, client=client)
    assert out == tmp_path / "1.15.7.61582"
    for table in TABLES:
        assert (out / "raw" / f"{table}.csv").read_text() == f"ID,Name_lang\n1,{table}\n"
    meta = json.loads((out / "raw" / "_meta.json").read_text())
    assert meta["product"] == "wow_classic_era" and meta["build"] == "1.15.7.61582"
    assert any("build=1.15.7.61582" in c for c in calls)


def test_download_table_writes_empty_csv_when_table_missing_for_product(tmp_path: Path):
    client = httpx.Client(
        transport=fake_transport([], missing_tables=frozenset({"JournalInstance"})),
        base_url="https://wago.tools",
    )
    path = download_table("JournalInstance", "1.15.7.61582", tmp_path, client)
    assert path.read_text() == "ID\n"


def test_download_table_404_on_required_table_fails(tmp_path: Path):
    client = httpx.Client(
        transport=fake_transport([], missing_tables=frozenset({"ItemSparse"})),
        base_url="https://wago.tools",
    )
    with pytest.raises(SystemExit):
        download_table("ItemSparse", "1.15.7.61582", tmp_path, client)


def test_download_table_accepts_id_column_anywhere_in_header(tmp_path: Path):
    def handler(request: httpx.Request) -> httpx.Response:
        return httpx.Response(200, text="Name_lang,ID\nWarrior,1\n")

    client = httpx.Client(transport=httpx.MockTransport(handler), base_url="https://wago.tools")
    path = download_table("ChrClasses", "1.15.7.61582", tmp_path, client)
    assert path.read_text() == "Name_lang,ID\nWarrior,1\n"


def test_tables_cover_every_phase_1_input():
    assert TABLES == [
        "Map",
        "AreaTable",
        "JournalInstance",
        "ItemSparse",
        "Item",
        "SpellName",
        "Spell",
        "SpellEffect",
        "SpellDuration",
        "SpellMisc",
        "ManifestInterfaceData",
        "ItemSet",
        "ItemSetSpell",
        "ChrClasses",
        "ChrRaces",
        "Talent",
        "TalentTab",
        "ItemArmorTotal",
        "ItemArmorQuality",
        "ItemArmorShield",
        "ArmorLocation",
        "RandPropPoints",
    ]
    assert "foreversixty-pipeline" in USER_AGENT
