"""Per-level base stats and race offsets, from wowhead's Forever gear planner.

Level 60 has been the sim's only level: `sim/core/character.go`'s base
strength/agility/intellect/spirit/stamina exist only as the level-60 table,
and there is no game source that carries a lower level's values -- DB2 has no
`CharBaseStats` on this client (see `pipeline/gametables.py`'s docstring) and
neither does `GameTables`. Wowhead's own gear planner already resolves the
per-level numbers it needs to project a lower-level character's stats onto
gear, at `baseStats.stats[classId][statId][level]` (statId 0 mana, 1 health,
3 agility, 4 strength, 5 intellect, 6 spirit, 7 stamina), `baseStats.raceOffsets`
and, separately, spell crit per point of intellect at `critSpell[classId][level]`
(absent for a class with no spell-crit curve at all). Level-60 values match the
engine's own tables to the point -- warrior stamina 110, health 1689 -- which is
this module's own regression check.

`load_levels` returns `None`, logging why, for a build with no wowhead payload
at all (Classic Era, or a 1.60 build `fetch-wowhead` has not run for yet): see
`docs/superpowers/specs/2026-09-27-level-aware-sim-design.md`. This mirrors
`pipeline.normalize.wowhead.load_supplement`'s soft-gap contract, but a build
with no wowhead payload writes no `levels.json` at all rather than an
unchanged one -- there is no non-wowhead source for this file to fall back to.
"""

from __future__ import annotations

import json
import logging
from pathlib import Path

from pipeline.manifest import refresh_manifest
from pipeline.models import ClassLevels, LevelsFile, LevelsSource, LevelStats, RaceStatOffsets
from pipeline.normalize import _write
from pipeline.wowhead_items import META_FILE, parse_page_data, raw_path

logger = logging.getLogger(__name__)

LEVELS_FILE = "levels.json"
BASESTATS_KEY = "wow.gearPlanner.classicplus.baseStats"
CRITSPELL_KEY = "wow.gearPlanner.classicplus.critSpell"
MAX_LEVEL = 60

#: `baseStats.stats[classId]` stat id -> the field it fills on `LevelStats`.
STAT_FIELD_BY_ID: dict[int, str] = {
    0: "mana",
    1: "health",
    3: "agility",
    4: "strength",
    5: "intellect",
    6: "spirit",
    7: "stamina",
}

#: `baseStats.raceOffsets[raceId]` stat id -> the field it fills on
#: `RaceStatOffsets`. Mana and health have no race offset in the payload --
#: only the five stats a gear planner projects onto a character sheet do.
RACE_OFFSET_FIELD_BY_ID: dict[int, str] = {
    3: "agility",
    4: "strength",
    5: "intellect",
    6: "spirit",
    7: "stamina",
}


class UnmappableIdError(ValueError):
    """A ChrClasses or ChrRaces id in the wowhead payload has no row in the
    build's own `classes.json`/`races.json`. Raised rather than guessed at --
    a build that adds a class or race the planner already knows about needs a
    new row in those files before its numbers can be trusted."""


def _slug_by_id(rows: list[dict]) -> dict[int, str]:
    return {int(row["id"]): row["slug"] for row in rows}


def _class_levels(
    stat_arrays: dict[str, list[float]], crit_by_level: list[float] | None
) -> ClassLevels:
    levels = [
        LevelStats(
            level=level,
            **{
                field: int(stat_arrays[str(stat_id)][level])
                for stat_id, field in STAT_FIELD_BY_ID.items()
            },
            spell_crit_per_int=float(crit_by_level[level]) if crit_by_level is not None else None,
        )
        for level in range(1, MAX_LEVEL + 1)
    ]
    return ClassLevels(levels=levels)


def _race_offsets(offsets: dict[str, float]) -> RaceStatOffsets:
    return RaceStatOffsets(
        **{
            field: int(offsets.get(str(stat_id), 0))
            for stat_id, field in RACE_OFFSET_FIELD_BY_ID.items()
        }
    )


def build_levels_file(
    build: str,
    base_stats: dict,
    crit_spell: dict,
    source: LevelsSource,
    class_rows: list[dict],
    race_rows: list[dict],
) -> LevelsFile:
    """The pure transform: wowhead's two page-data objects plus the build's
    own class/race slugs, with no file I/O."""
    class_slug_by_id = _slug_by_id(class_rows)
    race_slug_by_id = _slug_by_id(race_rows)

    classes: dict[str, ClassLevels] = {}
    for class_id_str, stat_arrays in base_stats.get("stats", {}).items():
        class_id = int(class_id_str)
        slug = class_slug_by_id.get(class_id)
        if slug is None:
            raise UnmappableIdError(
                f"ChrClasses id {class_id} (baseStats.stats) has no row in classes.json"
            )
        classes[slug] = _class_levels(stat_arrays, crit_spell.get(class_id_str))

    race_offsets: dict[str, RaceStatOffsets] = {}
    for race_id_str, offsets in base_stats.get("raceOffsets", {}).items():
        race_id = int(race_id_str)
        slug = race_slug_by_id.get(race_id)
        if slug is None:
            raise UnmappableIdError(
                f"ChrRaces id {race_id} (baseStats.raceOffsets) has no row in races.json"
            )
        race_offsets[slug] = _race_offsets(offsets)

    return LevelsFile(
        build=build,
        source=source,
        classes=dict(sorted(classes.items())),
        race_offsets=dict(sorted(race_offsets.items())),
    )


def load_levels(build: str, root: Path = Path("builds")) -> LevelsFile | None:
    """`build_levels_file` fed from a build directory, or None -- logged --
    when the build has no wowhead payload at all."""
    build_dir = root / build
    path = raw_path(build_dir)
    if not path.exists():
        logger.info("levels: no wowhead payload at %s; levels.json not written", path)
        return None
    text = path.read_text(encoding="utf-8")
    base_stats = parse_page_data(text, BASESTATS_KEY)
    crit_spell = parse_page_data(text, CRITSPELL_KEY)

    meta_path = path.parent / META_FILE
    if not meta_path.exists():
        raise SystemExit(f"no {meta_path}; run `python -m pipeline fetch-wowhead` first")
    meta = json.loads(meta_path.read_text(encoding="utf-8"))
    source = LevelsSource(url=meta["url"], fetched_at=meta["fetched_at"])

    class_rows = json.loads((build_dir / "classes.json").read_text(encoding="utf-8"))
    race_rows = json.loads((build_dir / "races.json").read_text(encoding="utf-8"))
    return build_levels_file(build, base_stats, crit_spell, source, class_rows, race_rows)


def write_levels(build: str, root: Path = Path("builds")) -> Path | None:
    result = load_levels(build, root)
    if result is None:
        return None
    path = root / build / LEVELS_FILE
    _write(result.model_dump(), path)
    refresh_manifest(root / build)
    logger.info(
        "wrote %s: %d classes, %d races", path, len(result.classes), len(result.race_offsets)
    )
    return path
