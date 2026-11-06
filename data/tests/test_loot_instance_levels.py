"""pipeline.loot.instance_levels: each dungeon's and raid's level range, from
the client's Group Finder rows onto loot.json. Synthetic CSVs in the real
tables' column names stand in for the wago exports; the committed file is
checked against the active build's loot.json."""

import csv
import json
from pathlib import Path

import pytest

from conftest import ACTIVE_BUILD
from pipeline.loot import instance_levels as il
from pipeline.models import LootFile, LootSource


def _write(raw: Path, table: str, rows: list[dict]) -> None:
    with (raw / f"{table}.csv").open("w", encoding="utf-8", newline="") as fh:
        writer = csv.DictWriter(fh, fieldnames=list(rows[0]))
        writer.writeheader()
        writer.writerows(rows)


def _act(id_, map_id, lo, hi, cat="2", difficulty="1"):
    return {
        "ID": id_,
        "GroupFinderCategoryID": cat,
        "MapID": map_id,
        "DifficultyID": difficulty,
        "MinLevel": lo,
        "MaxLevelSuggestion": hi,
    }


@pytest.fixture
def raw(tmp_path: Path) -> Path:
    _write(
        tmp_path,
        "GroupFinderCategory",
        [
            {"ID": "2", "Name_lang": "Dungeons"},
            {"ID": "114", "Name_lang": "Raids"},
            {"ID": "116", "Name_lang": "Quests & Zones"},
        ],
    )
    _write(
        tmp_path,
        "GroupFinderActivity",
        [
            _act("796", "43", "15", "24"),
            _act("803", "90", "29", "38"),
            # Season of Discovery's Gnomeregan raid on the same map: not classic.
            _act("1605", "90", "40", "49", cat="114", difficulty="197"),
            _act("812", "229", "55", "60"),
            _act("837", "229", "59", "60"),
            _act("838", "249", "60", "60", cat="114", difficulty="9"),
            _act("853", "0", "1", "14", cat="116", difficulty="0"),
        ],
    )
    _write(
        tmp_path,
        "Map",
        [
            {"ID": "0", "InstanceType": "0", "AreaTableID": "0"},
            {"ID": "43", "InstanceType": "1", "AreaTableID": "718"},
            {"ID": "90", "InstanceType": "1", "AreaTableID": "721"},
            {"ID": "229", "InstanceType": "1", "AreaTableID": "1583"},
            {"ID": "249", "InstanceType": "2", "AreaTableID": "2159"},
            {"ID": "70", "InstanceType": "1", "AreaTableID": "1337"},
        ],
    )
    _write(
        tmp_path,
        "AreaTable",
        [
            {"ID": "718", "AreaName_lang": "Wailing Caverns", "ContinentID": "43"},
            {"ID": "721", "AreaName_lang": "Gnomeregan", "ContinentID": "90"},
            {"ID": "1583", "AreaName_lang": "Blackrock Spire", "ContinentID": "0"},
            {"ID": "2159", "AreaName_lang": "Onyxia's Lair", "ContinentID": "0"},
            {"ID": "1337", "AreaName_lang": "Uldaman", "ContinentID": "70"},
        ],
    )
    return tmp_path


def _loot(*zones: tuple[str, str, int]) -> dict:
    return {"sources": [{"id": i, "kind": k, "zone_id": z} for i, k, z in zones]}


def test_reads_the_group_finder_range_by_area(raw):
    doc = il.build_instance_levels(raw, "b", _loot(("dungeon:wailing-caverns", "dungeon", 718)))
    assert doc["zones"]["718"] == {
        "name": "Wailing Caverns",
        "map_id": 43,
        "activity_ids": [796],
        "level_min": 15,
        "level_max": 24,
    }


def test_a_split_instance_spans_every_wing(raw):
    doc = il.build_instance_levels(raw, "b", _loot(("dungeon:blackrock-spire", "dungeon", 1583)))
    zone = doc["zones"]["1583"]
    assert (zone["level_min"], zone["level_max"], zone["activity_ids"]) == (55, 60, [812, 837])


def test_season_of_discovery_raid_rows_are_not_read(raw):
    doc = il.build_instance_levels(raw, "b", _loot(("dungeon:gnomeregan", "dungeon", 721)))
    assert (doc["zones"]["721"]["level_min"], doc["zones"]["721"]["level_max"]) == (29, 38)


def test_raids_are_read_and_other_kinds_ignored(raw):
    loot = _loot(("raid:onyxias-lair", "raid", 2159), ("zone:718", "zone", 718))
    assert set(il.build_instance_levels(raw, "b", loot)["zones"]) == {"2159"}


def test_an_instance_with_no_group_finder_row_is_left_for_the_loot_run_to_refuse(raw):
    doc = il.build_instance_levels(raw, "b", _loot(("dungeon:uldaman", "dungeon", 1337)))
    assert doc["zones"] == {}


def _document(*sources: LootSource) -> LootFile:
    return LootFile(sources=list(sources))


def _source(id_: str, kind: str, zone_id: int | None) -> LootSource:
    return LootSource(id=id_, kind=kind, name=id_, zone_id=zone_id, items=[1])


def test_apply_stamps_dungeons_and_raids_only():
    doc = il.apply_instance_levels(
        _document(
            _source("dungeon:wailing-caverns", "dungeon", 718),
            _source("raid:onyxias-lair", "raid", 2159),
            _source("zone:718", "zone", 718),
        ),
        {718: (15, 24), 2159: (60, 60)},
    )
    ranges = {s.id: (s.level_min, s.level_max) for s in doc.sources}
    assert ranges == {
        "dungeon:wailing-caverns": (15, 24),
        "raid:onyxias-lair": (60, 60),
        "zone:718": (None, None),
    }


def test_apply_fails_naming_every_instance_with_no_range():
    with pytest.raises(il.InstanceLevelsError) as err:
        il.apply_instance_levels(
            _document(
                _source("dungeon:uldaman", "dungeon", 1337),
                _source("raid:onyxias-lair", "raid", 2159),
            ),
            {},
        )
    assert "dungeon:uldaman (area 1337)" in str(err.value)
    assert "raid:onyxias-lair (area 2159)" in str(err.value)


def test_no_range_excuses_a_named_area(monkeypatch):
    monkeypatch.setattr(il, "NO_RANGE", {1337: "test: no Group Finder row"})
    doc = il.apply_instance_levels(_document(_source("dungeon:uldaman", "dungeon", 1337)), {})
    assert doc.sources[0].level_min is None


def test_an_announced_raid_with_no_area_ships_without_a_range():
    doc = il.apply_instance_levels(_document(_source("raid:announced", "raid", None)), {})
    assert doc.sources[0].level_min is None


def test_stale_and_write_round_trip(raw, tmp_path):
    doc = il.build_instance_levels(raw, "b", _loot(("dungeon:wailing-caverns", "dungeon", 718)))
    out = tmp_path / "instance-levels.json"
    assert il.stale(doc, out)
    il.write(doc, out)
    assert not il.stale(doc, out)
    assert il.load_instance_levels(tmp_path) == {718: (15, 24)}


def test_load_fails_without_the_committed_file(tmp_path):
    with pytest.raises(il.InstanceLevelsError):
        il.load_instance_levels(tmp_path)


def test_the_committed_file_is_the_clients_and_covers_every_instance_in_loot_json():
    """Wailing Caverns is 15-24 in build 1.60.1.70790's own Group Finder row
    (GroupFinderActivity 796); classic-db's 1.12 row says 17-24."""
    levels = il.load_instance_levels(Path("curated"))
    assert levels[718] == (15, 24)
    assert levels[2159] == (60, 60)
    assert levels[796] == (30, 46)
    loot = json.loads((Path("builds") / ACTIVE_BUILD / "loot.json").read_text(encoding="utf-8"))
    instance_areas = {
        s["zone_id"] for s in loot["sources"] if s["kind"] in il.INSTANCE_KINDS and s.get("zone_id")
    }
    assert instance_areas - set(levels) - set(il.NO_RANGE) == set()
