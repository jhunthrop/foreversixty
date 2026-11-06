import csv
import json
from pathlib import Path

import pytest

from pipeline.raid_bosses import (
    IN_CLIENT,
    NOT_IN_CLIENT,
    RaidBossesError,
    build_raid_bosses,
    fetch_missing_tables,
    write_raid_bosses,
)

BUILD = "1.60.1.1"

#: The first tier as curated/raids/first-tier.json lists it.
ZONES = [
    {"key": "raid:barrow-deeps", "name": "Barrow Deeps", "kind": "raid"},
    {"key": "raid:hyjal-summit", "name": "Hyjal Summit", "kind": "raid"},
    {"key": "raid:onyxias-lair", "name": "Onyxia's Lair", "kind": "raid"},
]

#: 1.60.1.70790's own rows: Onyxia's Lair by name, and two maps whose names only resemble
#: the announced raids.
MAPS = [
    {"ID": 249, "MapName_lang": "Onyxia's Lair"},
    {"ID": 2817, "MapName_lang": "Starfall Barrow Den"},
    {"ID": 2995, "MapName_lang": "Hyjal Crater"},
]
DIFFICULTIES = [
    {"ID": 4564, "MapID": 249, "MaxPlayers": 40},
    {"ID": 5621, "MapID": 2817, "MaxPlayers": 1},
]
ENCOUNTERS = [
    {"ID": 1084, "Name_lang": "Onyxia", "MapID": 249, "OrderIndex": 0},
    {"ID": 9001, "Name_lang": "Barrow Thing", "MapID": 2817, "OrderIndex": 0},
]
ITEMS = [
    {
        "id": 17064,
        "name": "Shard of the Scale",
        "quality": 4,
        "item_level": 71,
        "required_level": 60,
        "class_id": 4,
        "subclass_id": 0,
        "inventory_type": 12,
    },
    {
        "id": 18205,
        "name": "Eskhandar's Collar",
        "quality": 4,
        "item_level": 71,
        "required_level": 60,
        "class_id": 4,
        "subclass_id": 0,
        "inventory_type": 2,
    },
    {
        "id": 50001,
        "name": "Barrow Blade",
        "quality": 4,
        "item_level": 70,
        "required_level": 60,
        "class_id": 2,
        "subclass_id": 7,
        "inventory_type": 13,
    },
]
LOOT = {
    "sources": [
        {
            "id": "raid:onyxias-lair",
            "kind": "raid",
            "name": "Onyxia's Lair",
            "source_origin": "classic-db",
            "bosses": [
                {
                    "id": "raid:onyxias-lair:10184",
                    "name": "Onyxia",
                    "npc_id": 10184,
                    "items": [18205, 17064, 16900],
                },
            ],
        }
    ]
}


def _write_csv(path: Path, rows: list[dict]) -> None:
    columns = list(rows[0]) if rows else ["ID"]
    with path.open("w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=columns)
        writer.writeheader()
        writer.writerows(rows)


def _layout(tmp_path: Path, *, journal: dict[str, list[dict]] | None = None, maps=MAPS) -> Path:
    root = tmp_path / "builds"
    raw = root / BUILD / "raw"
    raw.mkdir(parents=True)
    _write_csv(raw / "Map.csv", maps)
    _write_csv(raw / "MapDifficulty.csv", DIFFICULTIES)
    _write_csv(raw / "DungeonEncounter.csv", ENCOUNTERS)
    journal = journal or {}
    for table in ("JournalInstance", "JournalEncounter", "JournalEncounterItem"):
        _write_csv(raw / f"{table}.csv", journal.get(table, []))
    (root / BUILD / "items.json").write_text(json.dumps(ITEMS), encoding="utf-8")
    (root / BUILD / "loot.json").write_text(json.dumps(LOOT), encoding="utf-8")
    curated = tmp_path / "curated" / "raids"
    curated.mkdir(parents=True)
    (curated / "first-tier.json").write_text(json.dumps({"zones": ZONES}), encoding="utf-8")
    return tmp_path


def _doc(tmp_path: Path) -> dict:
    root = tmp_path / "builds"
    return build_raid_bosses(BUILD, root / BUILD, root / BUILD / "raw", tmp_path / "curated")


def _raid(doc: dict, key: str) -> dict:
    return next(r for r in doc["raids"] if r["key"] == key)


def test_a_raid_the_client_does_not_name_is_empty_and_says_so(tmp_path):
    doc = _doc(_layout(tmp_path))
    for key, name in (("raid:barrow-deeps", "Barrow Deeps"), ("raid:hyjal-summit", "Hyjal Summit")):
        raid = _raid(doc, key)
        assert raid["status"] == NOT_IN_CLIENT
        assert raid["bosses"] == [] and raid["rows"] == []
        assert (
            raid["note"]
            == f"not in client {BUILD}: no Map or JournalInstance row is named {name!r}"
        )


def test_a_map_whose_name_only_resembles_the_raid_is_not_a_tie(tmp_path):
    """2817 Starfall Barrow Den carries an encounter, but it is not named Barrow Deeps."""
    raid = _raid(_doc(_layout(tmp_path)), "raid:barrow-deeps")
    assert all(b["encounter_id"] != 9001 for b in raid["bosses"])


def test_onyxia_ties_by_map_name_and_takes_the_loot_sources_drops(tmp_path):
    raid = _raid(_doc(_layout(tmp_path)), "raid:onyxias-lair")
    assert raid["status"] == IN_CLIENT
    assert raid["max_players"] == 40
    assert {"table": "Map", "id": 249, "name": "Onyxia's Lair"} in raid["rows"]
    [boss] = raid["bosses"]
    assert (boss["encounter_id"], boss["journal_encounter_id"], boss["name"], boss["order"]) == (
        1084,
        None,
        "Onyxia",
        1,
    )
    assert boss["rows"] == [{"table": "DungeonEncounter", "id": 1084, "name": "Onyxia"}]
    assert boss["loot_source"] == "loot.json raid:onyxias-lair:10184 (npc 10184, classic-db)"
    # 16900 is not in this build's items.json: counted, never shipped.
    assert [i["item_id"] for i in boss["loot"]] == [17064, 18205]
    assert boss["dropped_items"] == [16900]
    assert boss["loot"][0]["name"] == "Shard of the Scale"
    assert boss["loot"][0]["inventory_type"] == 12


def test_journal_loot_wins_when_the_client_ships_the_journal(tmp_path):
    maps = [*MAPS, {"ID": 3300, "MapName_lang": "Barrow Deeps"}]
    journal = {
        "JournalInstance": [{"ID": 700, "Name_lang": "Barrow Deeps", "MapID": 3300}],
        "JournalEncounter": [
            {
                "ID": 2500,
                "Name_lang": "Onyxia",
                "JournalInstanceID": 701,
                "DungeonEncounterID": 1084,
            },
            {
                "ID": 2600,
                "Name_lang": "Deep Warden",
                "JournalInstanceID": 700,
                "DungeonEncounterID": 9100,
            },
        ],
        "JournalEncounterItem": [
            {"ID": 1, "JournalEncounterID": 2500, "ItemID": 17064},
            {"ID": 2, "JournalEncounterID": 2600, "ItemID": 50001},
            {"ID": 3, "JournalEncounterID": 2600, "ItemID": 99999},
        ],
    }
    root = _layout(tmp_path, journal=journal, maps=maps)
    _write_csv(
        root / "builds" / BUILD / "raw" / "DungeonEncounter.csv",
        [*ENCOUNTERS, {"ID": 9100, "Name_lang": "Deep Warden", "MapID": 3300, "OrderIndex": 0}],
    )
    doc = _doc(root)
    assert doc["journal_in_client"] is True
    barrow = _raid(doc, "raid:barrow-deeps")
    assert barrow["status"] == IN_CLIENT and barrow["note"] is None
    assert {"table": "JournalInstance", "id": 700, "name": "Barrow Deeps"} in barrow["rows"]
    [warden] = barrow["bosses"]
    assert warden["journal_encounter_id"] == 2600
    assert warden["loot_source"] == "JournalEncounterItem (JournalEncounterID 2600)"
    assert [i["item_id"] for i in warden["loot"]] == [50001]
    assert warden["dropped_items"] == [99999]
    [onyxia] = _raid(doc, "raid:onyxias-lair")["bosses"]
    assert [i["item_id"] for i in onyxia["loot"]] == [17064]


def test_bosses_follow_the_clients_order_index(tmp_path):
    root = _layout(tmp_path)
    _write_csv(
        root / "builds" / BUILD / "raw" / "DungeonEncounter.csv",
        [
            {"ID": 1085, "Name_lang": "Second", "MapID": 249, "OrderIndex": 2000},
            {"ID": 1084, "Name_lang": "Onyxia", "MapID": 249, "OrderIndex": 0},
        ],
    )
    bosses = _raid(_doc(root), "raid:onyxias-lair")["bosses"]
    assert [(b["name"], b["order"]) for b in bosses] == [("Onyxia", 1), ("Second", 2)]
    assert bosses[1]["loot"] == [] and bosses[1]["loot_source"] is None


def test_missing_required_tables_fail_fast(tmp_path):
    root = _layout(tmp_path)
    (root / "builds" / BUILD / "raw" / "DungeonEncounter.csv").unlink()
    with pytest.raises(RaidBossesError, match="DungeonEncounter.csv"):
        _doc(root)


def test_write_then_check(tmp_path):
    root = _layout(tmp_path)
    kwargs = {"root": root / "builds", "curated_dir": root / "curated"}
    out = write_raid_bosses(BUILD, **kwargs)
    assert json.loads(out.read_text())["build"] == BUILD
    assert write_raid_bosses(BUILD, check=True, **kwargs) == out
    out.write_text("{}\n")
    with pytest.raises(RaidBossesError, match="drifted"):
        write_raid_bosses(BUILD, check=True, **kwargs)


def test_fetch_downloads_only_the_tables_raw_lacks(tmp_path):
    # activate-build.sh runs the step with --fetch when data.yml's fetch did not write the file.
    import httpx

    root = _layout(tmp_path)
    raw = root / "builds" / BUILD / "raw"
    (raw / "DungeonEncounter.csv").unlink()
    asked = []

    def handler(request):
        asked.append(request.url.path)
        if "Journal" in request.url.path:
            return httpx.Response(400, text="Table not found")
        return httpx.Response(200, text="ID,Name_lang,MapID,OrderIndex\n1084,Onyxia,249,0\n")

    client = httpx.Client(base_url="https://wago.tools", transport=httpx.MockTransport(handler))
    for journal in ("JournalInstance", "JournalEncounter", "JournalEncounterItem"):
        (raw / f"{journal}.csv").unlink(missing_ok=True)
    fetched = fetch_missing_tables(BUILD, raw, client)
    assert set(fetched) == {
        "DungeonEncounter",
        "JournalInstance",
        "JournalEncounter",
        "JournalEncounterItem",
    }
    assert asked and all(p.startswith("/db2/") for p in asked)
    assert (raw / "DungeonEncounter.csv").read_text().startswith("ID,")
    assert fetch_missing_tables(BUILD, raw, client) == []


def test_the_committed_build_carries_the_first_tier():
    """The committed file covers every first-tier raid and keeps Onyxia's encounter
    (progression.go's encounterOnyxia) and her loot as floors, never exact counts."""
    data = Path(__file__).resolve().parents[1]
    active = json.loads((data.parent / "web/src/data/active-build.json").read_text())["build"]
    path = data / "builds" / active / "raid-bosses.json"
    if not path.exists():
        pytest.skip(f"{path} not generated for the active build")
    doc = json.loads(path.read_text())
    zones = json.loads((data / "curated/raids/first-tier.json").read_text())["zones"]
    assert [r["key"] for r in doc["raids"]] == [z["key"] for z in zones]
    onyxia = _raid(doc, "raid:onyxias-lair")
    assert onyxia["status"] == IN_CLIENT
    [boss] = [b for b in onyxia["bosses"] if b["encounter_id"] == 1084]
    assert len(boss["loot"]) >= 20
    for raid in doc["raids"]:
        assert raid["status"] in {IN_CLIENT, NOT_IN_CLIENT}
        assert (raid["note"] is not None) or raid["status"] == IN_CLIENT
