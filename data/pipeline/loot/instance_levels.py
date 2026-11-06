"""Each dungeon's and raid's level range, read off the client, onto loot.json.

Spec 2026-11-05-web-layout 5.10 item 2: the BiS page writes a drop as
`Boss · Instance (levels a-b)`. The range is the live client's own, from
the table the in-game Group Finder draws (wago.tools `GroupFinderActivity`:
`MinLevel` to `MaxLevelSuggestion`), the same rows `pipeline.guild_events`
reads for the guild calendar.

Which source wins, and why: the client. `curated/raids/events.json` carries
the same client numbers but only for the dungeons open at launch, split by
wing and with no raids; classic-db's `instance_template` holds 1.12's
server levels, which Forever retuned (Wailing Caverns: 17-24 in the pinned
classic-db dump, 15-24 in the client).
So this module reads the client rows itself, for every loot.json dungeon
and raid, gated or not.

Two stages, the `guild-events` shape:

* `python -m pipeline instance-levels --build B --raw DIR [--fetch]` reads
  the client CSVs and writes the committed `curated/instance-levels.json`,
  keyed by the AreaTable id loot.json names an instance by (`zone_id`).
* every `loot`/`loot-merge` run reads that committed file
  (`load_instance_levels`) and stamps `level_min`/`level_max` on each
  dungeon and raid source (`apply_instance_levels`). A dungeon or raid
  with no range fails the run, listing every one, unless `NO_RANGE`
  names it with a reason. Nothing is guessed. A source with no `zone_id`
  is a curated overlay's announcement of an instance the client has no
  map for yet, so there is no client range to read; it ships without one
  and the run logs it.

Only the classic difficulties count. The client also lists Season of
Discovery's raid versions of three dungeons and the later raid tunings on
the same maps (Gnomeregan's 10-player raid at 40-49, a 25-player
Ahn'Qiraj); a range is read from `CLASSIC_DIFFICULTIES` rows alone. An
instance the Group Finder splits into wings (Scarlet Monastery, Blackrock
Spire, Dire Maul, Stratholme) is one loot source, so its range spans
every wing: lowest `MinLevel` to highest `MaxLevelSuggestion`.
"""

from __future__ import annotations

import json
import logging
from pathlib import Path

from pipeline.guild_events import area_map, read_rows
from pipeline.models import LootFile

#: The kinds that carry a level range.
INSTANCE_KINDS = frozenset({"dungeon", "raid"})
#: GroupFinderCategory names whose rows are instances.
INSTANCE_CATEGORIES = frozenset({"Dungeons", "Raids"})
#: Difficulty ids of the original instances: 1 normal (five and ten
#: players), 9 forty players, 148 twenty players (Zul'Gurub, Ruins of
#: Ahn'Qiraj). Season of Discovery and later tunings use others (185, 186,
#: 197, 215, 226).
CLASSIC_DIFFICULTIES = frozenset({1, 9, 148})
#: AreaTable id -> why that dungeon or raid ships without a level range.
#: Empty: every instance on build 1.60.1.70790 has a classic Group Finder row.
NO_RANGE: dict[int, str] = {}

OUT_JSON = Path("curated/instance-levels.json")

logger = logging.getLogger(__name__)


class InstanceLevelsError(SystemExit):
    """A dungeon or raid has no level range, or the client tables disagree."""


def _instance_zone_ids(loot: dict) -> list[int]:
    return sorted(
        {
            int(source["zone_id"])
            for source in loot["sources"]
            if source.get("kind") in INSTANCE_KINDS and source.get("zone_id")
        }
    )


def _classic_activities(raw: Path) -> dict[int, list[dict[str, str]]]:
    """Map id -> its classic-difficulty dungeon and raid Group Finder rows."""
    categories = {
        row["ID"]
        for row in read_rows(raw, "GroupFinderCategory")
        if row["Name_lang"] in INSTANCE_CATEGORIES
    }
    by_map: dict[int, list[dict[str, str]]] = {}
    for row in read_rows(raw, "GroupFinderActivity"):
        if row["GroupFinderCategoryID"] in categories and (
            int(row["DifficultyID"]) in CLASSIC_DIFFICULTIES
        ):
            by_map.setdefault(int(row["MapID"]), []).append(row)
    return by_map


def build_instance_levels(raw: Path, build: str, loot: dict) -> dict:
    """The committed document: every loot.json dungeon and raid area that
    has a classic Group Finder row, with its map, activities and range."""
    maps = {int(row["ID"]): row for row in read_rows(raw, "Map")}
    areas = {int(row["ID"]): row for row in read_rows(raw, "AreaTable")}
    activities = _classic_activities(raw)
    zones: dict[str, dict] = {}
    for zone_id in _instance_zone_ids(loot):
        if zone_id not in areas:
            raise InstanceLevelsError(f"loot.json names area {zone_id}; AreaTable has no such row")
        map_id, _ = area_map(areas[zone_id], maps)
        rows = activities.get(map_id)
        if not rows:
            continue
        zones[str(zone_id)] = {
            "name": areas[zone_id]["AreaName_lang"],
            "map_id": map_id,
            "activity_ids": sorted(int(row["ID"]) for row in rows),
            "level_min": min(int(row["MinLevel"]) for row in rows),
            "level_max": max(int(row["MaxLevelSuggestion"]) for row in rows),
        }
    return {
        "notes": (
            "Generated by `python -m pipeline instance-levels` "
            "(data/pipeline/loot/instance_levels.py); do not edit by hand. Each dungeon's "
            "and raid's level range is the client's own Group Finder row for build "
            f"{build} (wago.tools GroupFinderActivity MinLevel to MaxLevelSuggestion, "
            "classic difficulties only, every wing of a split instance), keyed by the "
            "AreaTable id loot.json names the instance by."
        ),
        "build": build,
        "zones": zones,
    }


def render_json(doc: dict) -> str:
    return json.dumps(doc, indent=2, ensure_ascii=False) + "\n"


def write(doc: dict, out_json: Path = OUT_JSON) -> Path:
    out_json.write_text(render_json(doc), encoding="utf-8")
    return out_json


def stale(doc: dict, out_json: Path = OUT_JSON) -> bool:
    return not out_json.exists() or out_json.read_text(encoding="utf-8") != render_json(doc)


def load_instance_levels(curated_dir: Path) -> dict[int, tuple[int, int]]:
    """AreaTable id -> (level_min, level_max), from the committed file."""
    path = curated_dir / OUT_JSON.name
    if not path.exists():
        raise InstanceLevelsError(f"no {path}; run `python -m pipeline instance-levels` first")
    zones = json.loads(path.read_text(encoding="utf-8"))["zones"]
    return {int(key): (row["level_min"], row["level_max"]) for key, row in zones.items()}


def apply_instance_levels(document: LootFile, levels: dict[int, tuple[int, int]]) -> LootFile:
    """Every dungeon and raid source with its level range. Fails, naming
    each one, when a source has no range and `NO_RANGE` does not excuse it."""
    instances = [source for source in document.sources if source.kind in INSTANCE_KINDS]
    announced = [source.id for source in instances if source.zone_id is None]
    if announced:
        logger.info("instance-levels: no client area, so no range, for %s", ", ".join(announced))
    missing = sorted(
        f"{source.id} (area {source.zone_id})"
        for source in instances
        if source.zone_id is not None
        and source.zone_id not in levels
        and source.zone_id not in NO_RANGE
    )
    if missing:
        raise InstanceLevelsError(
            "dungeon/raid sources with no level range in curated/instance-levels.json "
            f"(rerun `python -m pipeline instance-levels`, or add to NO_RANGE with a reason): "
            f"{', '.join(missing)}"
        )
    sources = [
        source.model_copy(
            update=dict(zip(("level_min", "level_max"), levels[source.zone_id], strict=True))
        )
        if source.kind in INSTANCE_KINDS and source.zone_id in levels
        else source
        for source in document.sources
    ]
    return document.model_copy(update={"sources": sources})
