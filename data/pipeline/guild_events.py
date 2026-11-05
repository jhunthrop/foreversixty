"""The guild calendar's dungeon runs and guild events, read off the client.

design/specs/2026-11-05-guild-events.md section 3.1. An officer can post a
dungeon run or a guild event on the Raids tab as well as a raid. Each
dungeon's size and level range are the live client's own facts, from the
tables the in-game Group Finder draws (wago.tools DB2 exports for the
active build):

- `GroupFinderActivity`: one row per run the Group Finder lists, wings
  apart (Scarlet Monastery's four, Stratholme's two gates, Lower and Upper
  Blackrock Spire, Dire Maul's three). `MaxPlayers` is the size and
  `MinLevel` to `MaxLevelSuggestion` the level range. Only rows of the
  `Dungeons` category are read.
- `Map`: `InstanceType` must be 1 (party) or 2 (raid); anything else is
  refused.
- `MapDifficulty`: the map's normal-difficulty `MaxPlayers` must hold the
  activity's size, a cross-check that the two tables agree.
- `AreaTable`: the map's root area, which is both the zone name the game
  shows and the id `loot.json` names a dungeon source by.
- `LFGDungeons`: names the two Forever dungeons `loot.json` has no source
  for yet, so their inclusion is the client's word, not ours.
- `Lock`: a key item is shipped only when the client has a lock opened by
  that item (and the door it opens is named in the source notes).

Which dungeons: every `loot.json` dungeon source open at launch (no
`opens`; its event variants carry one), plus `NEW_DUNGEONS`. Dire Maul is
in the client but gated (`content-phases.json`, `loot.json` opens
`later`), so it is left out until its gate opens.

The run writes `curated/raids/events.json` and the API's Go twin
`api/internal/guilds/raids_zones_catalogue.go`; `raids_zones_test.go`
pins the two together.
"""

from __future__ import annotations

import csv
import json
import re
from dataclasses import dataclass
from pathlib import Path

import httpx

from pipeline.wago import BASE_URL, USER_AGENT, download_table

#: The client tables this generator reads, fetched into `--raw`.
TABLES: tuple[str, ...] = (
    "GroupFinderActivity",
    "GroupFinderCategory",
    "Map",
    "MapDifficulty",
    "AreaTable",
    "LFGDungeons",
    "Lock",
)

DUNGEONS_CATEGORY = "Dungeons"
#: Map.InstanceType: 1 party, 2 raid. 0 is a world map, 3 a battleground, 4 an arena.
DUNGEON_INSTANCE_TYPES = frozenset({1, 2})
#: MapDifficulty's normal five-player difficulty, which carries a dungeon's own cap.
NORMAL_DIFFICULTY = 1
#: Lock.Type for "an item opens this".
LOCK_TYPE_ITEM = 1

#: Forever's new dungeons `loot.json` has no source for yet, by LFGDungeons
#: row id: the client lists each as a Looking for Group dungeon and the
#: Group Finder lists it under Dungeons (Blizzard's What's Next panel
#: named them). Ruins of Lordaeron and Excavation Site: Wetlands come in
#: through `loot.json` already.
NEW_DUNGEONS: dict[int, str] = {3271: "City of Dalaran", 3273: "Hall of Thanes"}

#: The name of each run on a map the Group Finder splits into wings, by
#: GroupFinderActivity id. A split map whose row is missing here fails the
#: run. Stratholme's Main Gate is the living side and the Service Gate the
#: undead side.
WING_NAMES: dict[int, str] = {
    805: "Scarlet Monastery: Graveyard",
    829: "Scarlet Monastery: Library",
    827: "Scarlet Monastery: Armory",
    828: "Scarlet Monastery: Cathedral",
    813: "Dire Maul: East",
    814: "Dire Maul: West",
    815: "Dire Maul: North",
    816: "Stratholme: Live",
    1603: "Stratholme: Undead",
    812: "Lower Blackrock Spire",
    837: "Upper Blackrock Spire",
}


@dataclass(frozen=True)
class KeyItem:
    item_id: int
    name: str
    #: Where the lock is stated; recorded in the file's notes.
    evidence: str
    #: The one door the key opens, when the run has another way in that needs
    #: no key: the site then says `{key} opens {door}.`, never `Needs a key`.
    door: str | None = None


#: The key a run needs, by GroupFinderActivity id. Each is checked against
#: the client's Lock table at generation; the door is verified on
#: cmangos/classic-db (commit ec4f5961, ClassicDB_1_12_1_z2815).
KEY_ITEMS: dict[int, KeyItem] = {
    797: KeyItem(
        13704,
        "Skeleton Key",
        "classic-db gameobject 174626 'Scholomance Door' (lock 1159)",
    ),
    1603: KeyItem(
        12382,
        "Key to the City",
        "classic-db gameobject 175368 'Service Entrance Gate', map 329 (lock 879)",
        # The key opens only the back gate; groups enter the Undead side through
        # the main gate (no lock), so the run does not need it.
        door="the service entrance",
    ),
    837: KeyItem(
        12344,
        "Seal of Ascension",
        "client Lock 820; classic-db quest 5584: the seal is the first step "
        "in gaining access to the upper citadel",
    ),
}

#: The guild event's six keys (spec 3.1): one name, an officer-chosen size.
EVENT_NAME = "Guild event"
EVENT_SIZES: tuple[int | None, ...] = (None, 5, 10, 15, 20, 40)

OUT_JSON = Path("curated/raids/events.json")
OUT_GO = Path("../api/internal/guilds/raids_zones_catalogue.go")


class GuildEventsError(SystemExit):
    """The client tables say something this generator will not publish."""


def fetch(build: str, raw: Path, client: httpx.Client | None = None) -> None:
    """Download TABLES for build into raw (a directory of its own)."""
    own = client is None
    client = client or httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    try:
        for table in TABLES:
            download_table(table, build, raw, client)
    finally:
        if own:
            client.close()


def _rows(raw: Path, table: str) -> list[dict[str, str]]:
    with (raw / f"{table}.csv").open(encoding="utf-8", newline="") as fh:
        return list(csv.DictReader(fh))


def slug(name: str) -> str:
    """`Scarlet Monastery: Library` -> `scarlet-monastery-library`."""
    return re.sub(r"[^a-z0-9]+", "-", name.lower().replace("'", "")).strip("-")


def _open_dungeon_areas(loot: dict) -> set[int]:
    """AreaTable ids of the loot.json dungeon sources open at launch."""
    return {
        int(s["zone_id"])
        for s in loot["sources"]
        if s.get("kind") == "dungeon" and "opens" not in s and s.get("zone_id")
    }


def _area_map(area: dict[str, str], maps: dict[int, dict[str, str]]) -> tuple[int, bool]:
    """The instance map an area names, and whether the area lies inside it.

    An area inside an instance has that map as its ContinentID (The
    Deadmines, 1581, on map 36); an entrance area outside it is the map's
    own AreaTableID (Blackrock Spire, 1583, on Eastern Kingdoms, for map 229).
    """
    continent = int(area["ContinentID"])
    if int(maps.get(continent, {}).get("InstanceType", 0)) in DUNGEON_INSTANCE_TYPES:
        return continent, True
    hits = [m for m, row in maps.items() if row["AreaTableID"] == area["ID"]]
    if len(hits) != 1:
        raise GuildEventsError(
            f"area {area['ID']} ({area['AreaName_lang']}) names {len(hits)} maps"
        )
    return hits[0], False


def _wanted_maps(raw: Path, loot: dict, maps: dict[int, dict[str, str]]) -> dict[int, str]:
    """Each wanted instance map and the name its single run takes.

    The name is the loot source's area, an area inside the map winning over
    an entrance area (Sunken Temple over The Temple of Atal'Hakkar).
    """
    areas = {int(a["ID"]): a for a in _rows(raw, "AreaTable")}
    named: dict[int, tuple[bool, str]] = {}
    for area_id in sorted(_open_dungeon_areas(loot)):
        map_id, inside = _area_map(areas[area_id], maps)
        if map_id not in named or (inside and not named[map_id][0]):
            named[map_id] = (inside, areas[area_id]["AreaName_lang"])
    lfg = {int(r["ID"]): r["Name_lang"] for r in _rows(raw, "LFGDungeons")}
    for lfg_id, name in NEW_DUNGEONS.items():
        if lfg.get(lfg_id) != name:
            raise GuildEventsError(f"LFGDungeons {lfg_id} is {lfg.get(lfg_id)!r}, not {name!r}")
        hits = {
            int(a["ContinentID"])
            for a in areas.values()
            if a["AreaName_lang"].removeprefix("The ") == name
            and int(maps.get(int(a["ContinentID"]), {}).get("InstanceType", 0))
            in DUNGEON_INSTANCE_TYPES
        }
        if len(hits) != 1:
            raise GuildEventsError(f"{name}: {len(hits)} instance maps carry that area name")
        named[hits.pop()] = (True, name)
    return {m: name for m, (_, name) in named.items()}


def _check_key(raw: Path, activity: int, key: KeyItem) -> None:
    for lock in _rows(raw, "Lock"):
        for i in range(8):
            if lock[f"_Index_{i}"] == str(key.item_id) and lock[f"Type_{i}"] == str(LOCK_TYPE_ITEM):
                return
    raise GuildEventsError(f"activity {activity}: no client Lock opens with item {key.item_id}")


def _dungeon_rows(raw: Path, build: str, loot: dict, qualities: dict[int, int]) -> list[dict]:
    maps = {int(r["ID"]): r for r in _rows(raw, "Map")}
    normal_cap = {
        int(r["MapID"]): int(r["MaxPlayers"])
        for r in _rows(raw, "MapDifficulty")
        if r["DifficultyID"] == str(NORMAL_DIFFICULTY)
    }
    category = next(
        r["ID"] for r in _rows(raw, "GroupFinderCategory") if r["Name_lang"] == DUNGEONS_CATEGORY
    )
    wanted = _wanted_maps(raw, loot, maps)
    acts = [
        r
        for r in _rows(raw, "GroupFinderActivity")
        if r["GroupFinderCategoryID"] == category and int(r["MapID"]) in wanted
    ]
    per_map: dict[int, int] = {}
    for a in acts:
        per_map[int(a["MapID"])] = per_map.get(int(a["MapID"]), 0) + 1
    if missing := set(wanted) - set(per_map):
        raise GuildEventsError(f"maps with no Group Finder dungeon row: {sorted(missing)}")
    out = [_dungeon_row(raw, a, maps, wanted, normal_cap, per_map, build, qualities) for a in acts]
    return sorted(out, key=lambda z: (z["min_level"], z["max_level"], z["name"]))


def _dungeon_row(raw, act, maps, names, normal_cap, per_map, build, qualities) -> dict:
    act_id, map_id = int(act["ID"]), int(act["MapID"])
    instance = int(maps[map_id]["InstanceType"])
    if instance not in DUNGEON_INSTANCE_TYPES:
        raise GuildEventsError(f"map {map_id} has InstanceType {instance}, not a dungeon")
    size = int(act["MaxPlayers"])
    if size <= 0 or size > normal_cap.get(map_id, 0):
        raise GuildEventsError(f"activity {act_id}: size {size} over map {map_id}'s cap")
    if per_map[map_id] > 1:
        if act_id not in WING_NAMES:
            raise GuildEventsError(f"activity {act_id} ({act['FullName_lang']}) needs a wing name")
        name = WING_NAMES[act_id]
    else:
        name = names[map_id]
    key = KEY_ITEMS.get(act_id)
    if key is not None:
        _check_key(raw, act_id, key)
    return {
        "key": f"dungeon:{slug(name)}",
        "name": name,
        "kind": "dungeon",
        "size": size,
        "min_level": int(act["MinLevel"]),
        "max_level": int(act["MaxLevelSuggestion"]),
        "key_item": None if key is None else _key_row(key, qualities),
        "log_zone": None,
        "map_id": map_id,
        "activity_id": act_id,
        "build": build,
    }


def _key_row(key: KeyItem, qualities: dict[int, int]) -> dict:
    if key.item_id not in qualities:
        raise GuildEventsError(f"{key.name} ({key.item_id}) is not in the build's items.json")
    row = {"id": key.item_id, "name": key.name, "quality": qualities[key.item_id]}
    return row if key.door is None else {**row, "door": key.door}


def _event_rows() -> list[dict]:
    return [
        {
            "key": "event:guild" if size is None else f"event:guild-{size}",
            "name": EVENT_NAME,
            "kind": "event",
            "size": size,
            "min_level": None,
            "max_level": None,
            "key_item": None,
            "log_zone": None,
            "map_id": None,
            "activity_id": None,
            "build": None,
        }
        for size in EVENT_SIZES
    ]


def build_events(raw: Path, build: str, loot: dict, items: list[dict]) -> dict:
    """items is the build's items.json: a key's quality colour is the client's."""
    qualities = {int(i["id"]): int(i["quality"]) for i in items}
    keys_note = "; ".join(f"{k.name}: {k.evidence}" for k in KEY_ITEMS.values())
    return {
        "notes": (
            "Generated by `python -m pipeline guild-events` (data/pipeline/guild_events.py); "
            "do not edit by hand. Dungeon size, level range and instance type are the client's "
            f"own Group Finder rows for build {build} (wago.tools GroupFinderActivity, Map, "
            "MapDifficulty, AreaTable); Dire Maul stays out while content-phases.json gates it. "
            f"Keys: {keys_note}. Guild events are curated (spec 2026-11-05 section 3.1)."
        ),
        "build": build,
        "zones": _dungeon_rows(raw, build, loot, qualities) + _event_rows(),
    }


def _go_int(v: int | None) -> str:
    return "nil" if v is None else f"intPtr({v})"


def _go_row(z: dict) -> str:
    key = z["key_item"]
    key_go = "nil"
    if key is not None:
        name = json.dumps(key["name"])
        door = "" if key.get("door") is None else f", Door: {json.dumps(key['door'])}"
        key_go = f"&KeyItem{{ID: {key['id']}, Name: {name}, Quality: {key['quality']}{door}}}"
    kind = {"dungeon": "ZoneDungeon", "event": "ZoneEvent"}[z["kind"]]
    return (
        f"\t{{Key: {json.dumps(z['key'])}, Name: {json.dumps(z['name'])}, Kind: {kind}, "
        f"Size: {_go_int(z['size'])}, MinLevel: {_go_int(z['min_level'])}, "
        f"MaxLevel: {_go_int(z['max_level'])}, KeyItem: {key_go}}},"
    )


def render_go(doc: dict) -> str:
    rows = "\n".join(_go_row(z) for z in doc["zones"])
    return (
        "// Code generated by `python -m pipeline guild-events` (data/pipeline/guild_events.py)."
        " DO NOT EDIT.\n\n"
        "package guilds\n\n"
        f"// eventZones are the dungeons (client build {doc['build']}) and guild events of\n"
        "// data/curated/raids/events.json, in its order.\n"
        f"var eventZones = []Zone{{\n{rows}\n}}\n"
    )


def render_json(doc: dict) -> str:
    return json.dumps(doc, indent=2, ensure_ascii=False) + "\n"


def write(doc: dict, out_json: Path = OUT_JSON, out_go: Path = OUT_GO) -> list[Path]:
    out_json.write_text(render_json(doc), encoding="utf-8")
    out_go.write_text(render_go(doc), encoding="utf-8")
    return [out_json, out_go]


def stale(doc: dict, out_json: Path = OUT_JSON, out_go: Path = OUT_GO) -> list[Path]:
    """The outputs whose committed text differs from doc's."""
    want = {out_json: render_json(doc), out_go: render_go(doc)}
    return [
        p for p, text in want.items() if not p.exists() or p.read_text(encoding="utf-8") != text
    ]
