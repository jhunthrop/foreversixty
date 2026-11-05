# data/tests/test_guild_events.py
"""pipeline.guild_events: the guild calendar's dungeons read off the
client's Group Finder tables (spec 2026-11-05 section 3.1). Small synthetic
CSVs in the real tables' column names stand in for the wago exports."""

import csv
import json
from pathlib import Path

import pytest

from pipeline import guild_events as ge

BUILD = "1.60.1.70790"


def _write(raw: Path, table: str, rows: list[dict]) -> None:
    with (raw / f"{table}.csv").open("w", encoding="utf-8", newline="") as fh:
        w = csv.DictWriter(fh, fieldnames=list(rows[0]))
        w.writeheader()
        w.writerows(rows)


def _act(id_, name, map_id, size, lo, hi, cat="2"):
    return {
        "FullName_lang": name,
        "ID": id_,
        "GroupFinderCategoryID": cat,
        "MapID": map_id,
        "MaxPlayers": size,
        "MinLevel": lo,
        "MaxLevelSuggestion": hi,
    }


def _area(id_, name, continent):
    return {"ID": id_, "AreaName_lang": name, "ContinentID": continent, "ParentAreaID": "0"}


def _lock(id_, item):
    row = {"ID": id_}
    for i in range(8):
        row[f"_Index_{i}"] = item if i == 0 else "0"
        row[f"Type_{i}"] = "1" if i == 0 else "0"
    return row


@pytest.fixture
def raw(tmp_path: Path) -> Path:
    _write(
        tmp_path,
        "GroupFinderCategory",
        [{"ID": "2", "Name_lang": "Dungeons"}, {"ID": "114", "Name_lang": "Raids"}],
    )
    _write(
        tmp_path,
        "GroupFinderActivity",
        [
            _act("799", "Deadmines", "36", "5", "17", "26"),
            _act("812", "Lower Blackrock Spire", "229", "5", "55", "60"),
            _act("837", "Upper Blackrock Spire", "229", "10", "59", "60"),
            _act("1983", "City of Dalaran", "2959", "5", "28", "35"),
            _act("813", "Dire Maul (East)", "429", "5", "54", "60"),
            _act("838", "Onyxia's Lair", "249", "40", "60", "60", cat="114"),
        ],
    )
    _write(
        tmp_path,
        "Map",
        [
            {"ID": "0", "InstanceType": "0", "AreaTableID": "0"},
            {"ID": "36", "InstanceType": "1", "AreaTableID": "0"},
            {"ID": "229", "InstanceType": "1", "AreaTableID": "1583"},
            {"ID": "249", "InstanceType": "2", "AreaTableID": "2159"},
            {"ID": "429", "InstanceType": "1", "AreaTableID": "2557"},
            {"ID": "2959", "InstanceType": "1", "AreaTableID": "16544"},
        ],
    )
    _write(
        tmp_path,
        "MapDifficulty",
        [
            {"ID": "1", "DifficultyID": "1", "MaxPlayers": "5", "MapID": "36"},
            {"ID": "2", "DifficultyID": "1", "MaxPlayers": "10", "MapID": "229"},
            {"ID": "3", "DifficultyID": "1", "MaxPlayers": "5", "MapID": "429"},
            {"ID": "4", "DifficultyID": "1", "MaxPlayers": "5", "MapID": "2959"},
        ],
    )
    _write(
        tmp_path,
        "AreaTable",
        [
            _area("206", "Westfall", "36"),
            _area("1581", "The Deadmines", "36"),
            _area("1583", "Blackrock Spire", "0"),
            _area("2557", "Dire Maul", "429"),
            _area("16544", "City of Dalaran", "2959"),
        ],
    )
    _write(
        tmp_path,
        "LFGDungeons",
        [
            {"ID": "3271", "Name_lang": "City of Dalaran"},
            {"ID": "3273", "Name_lang": "Hall of Thanes"},
        ],
    )
    _write(tmp_path, "Lock", [_lock("820", "12344")])
    return tmp_path


ITEMS = [{"id": 12344, "quality": 3}]

LOOT = {
    "sources": [
        {"id": "dungeon:the-deadmines", "kind": "dungeon", "zone_id": 1581},
        {"id": "dungeon:blackrock-spire", "kind": "dungeon", "zone_id": 1583},
        {
            "id": "dungeon:blackrock-spire:tier-0-5",
            "kind": "dungeon",
            "zone_id": 1583,
            "opens": "later",
        },
        {"id": "dungeon:dire-maul", "kind": "dungeon", "zone_id": 2557, "opens": "later"},
        {"id": "raid:onyxias-lair", "kind": "raid", "zone_id": 2159},
    ]
}


@pytest.fixture
def no_thanes(monkeypatch):
    monkeypatch.setattr(ge, "NEW_DUNGEONS", {3271: "City of Dalaran"})
    monkeypatch.setattr(ge, "KEY_ITEMS", {837: ge.KEY_ITEMS[837]})


def test_reads_size_levels_and_wings_off_the_group_finder(raw, no_thanes):
    zones = ge.build_events(raw, BUILD, LOOT, ITEMS)["zones"]
    dungeons = [z for z in zones if z["kind"] == "dungeon"]
    assert [(z["key"], z["name"], z["size"], z["min_level"], z["max_level"]) for z in dungeons] == [
        ("dungeon:the-deadmines", "The Deadmines", 5, 17, 26),
        ("dungeon:city-of-dalaran", "City of Dalaran", 5, 28, 35),
        ("dungeon:lower-blackrock-spire", "Lower Blackrock Spire", 5, 55, 60),
        ("dungeon:upper-blackrock-spire", "Upper Blackrock Spire", 10, 59, 60),
    ]
    assert all(z["build"] == BUILD for z in dungeons)


def test_gated_dungeons_and_raids_stay_out(raw, no_thanes):
    keys = {z["key"] for z in ge.build_events(raw, BUILD, LOOT, ITEMS)["zones"]}
    assert not any("dire-maul" in k or "onyxia" in k for k in keys)


def test_a_key_ships_only_with_a_client_lock(raw, no_thanes):
    zones = {z["key"]: z for z in ge.build_events(raw, BUILD, LOOT, ITEMS)["zones"]}
    assert zones["dungeon:upper-blackrock-spire"]["key_item"] == {
        "id": 12344,
        "name": "Seal of Ascension",
        "quality": 3,
    }
    assert zones["dungeon:the-deadmines"]["key_item"] is None
    _write(raw, "Lock", [_lock("1", "1")])
    with pytest.raises(ge.GuildEventsError, match="no client Lock"):
        ge.build_events(raw, BUILD, LOOT, ITEMS)


def test_a_map_that_is_not_an_instance_is_refused(raw, no_thanes):
    maps = [
        {"ID": "0", "InstanceType": "0", "AreaTableID": "0"},
        {"ID": "36", "InstanceType": "4", "AreaTableID": "1581"},
        {"ID": "229", "InstanceType": "1", "AreaTableID": "1583"},
        {"ID": "2959", "InstanceType": "1", "AreaTableID": "16544"},
    ]
    _write(raw, "Map", maps)
    with pytest.raises(ge.GuildEventsError):
        ge.build_events(raw, BUILD, LOOT, ITEMS)


def test_a_split_map_needs_every_wing_named(raw, no_thanes, monkeypatch):
    monkeypatch.setattr(ge, "WING_NAMES", {812: "Lower Blackrock Spire"})
    with pytest.raises(ge.GuildEventsError, match="needs a wing name"):
        ge.build_events(raw, BUILD, LOOT, ITEMS)


def test_a_new_dungeon_the_client_does_not_name_fails(raw):
    with pytest.raises(ge.GuildEventsError, match="Hall of Thanes"):
        ge.build_events(raw, BUILD, LOOT, ITEMS)


def test_six_guild_event_keys_after_the_dungeons(raw, no_thanes):
    events = [z for z in ge.build_events(raw, BUILD, LOOT, ITEMS)["zones"] if z["kind"] == "event"]
    assert [(z["key"], z["name"], z["size"]) for z in events] == [
        ("event:guild", "Guild event", None),
        ("event:guild-5", "Guild event", 5),
        ("event:guild-10", "Guild event", 10),
        ("event:guild-15", "Guild event", 15),
        ("event:guild-20", "Guild event", 20),
        ("event:guild-40", "Guild event", 40),
    ]


def test_writes_json_and_go_and_reports_drift(raw, no_thanes, tmp_path):
    doc = ge.build_events(raw, BUILD, LOOT, ITEMS)
    out_json, out_go = tmp_path / "events.json", tmp_path / "catalogue.go"
    ge.write(doc, out_json, out_go)
    assert json.loads(out_json.read_text())["zones"] == doc["zones"]
    go = out_go.read_text()
    assert "DO NOT EDIT" in go
    assert 'Key: "dungeon:upper-blackrock-spire"' in go
    assert '&KeyItem{ID: 12344, Name: "Seal of Ascension", Quality: 3}' in go
    assert ge.stale(doc, out_json, out_go) == []
    out_go.write_text("package guilds\n")
    assert ge.stale(doc, out_json, out_go) == [out_go]


def test_the_committed_file_is_the_active_builds_client_data():
    doc = json.loads(Path("curated/raids/events.json").read_text(encoding="utf-8"))
    dungeons = [z for z in doc["zones"] if z["kind"] == "dungeon"]
    assert dungeons, "no dungeons in events.json"
    assert all(z["size"] and z["min_level"] <= z["max_level"] for z in dungeons)
    order = [(z["min_level"], z["max_level"], z["name"]) for z in dungeons]
    assert order == sorted(order)


def test_slug():
    assert ge.slug("Scarlet Monastery: Library") == "scarlet-monastery-library"
    assert ge.slug("Zul'Farrak") == "zulfarrak"
    assert ge.slug("Excavation Site: Wetlands") == "excavation-site-wetlands"
