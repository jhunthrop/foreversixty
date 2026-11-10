"""Race icons, from the client: race slug -> icon name.

The guides' race list, the character header and the roster show each race's own portrait
icon, the square `Achievement_Character_<Race>_Male.blp` the client ships in
`Interface\\ICONS\\` (it matches the item and spell icon frames). The race's part of the
name is the client's own `ChrRaces.Name_lang` with the spaces dropped (Night Elf ->
nightelf, Undead -> undead). A race the client carries no portrait for (Forever's two
Skyborne races) takes the portrait of the race the client itself draws it with,
`ChrRaces.MaleModelFallbackRaceID` (Blood Elf for both), so no race is left without art.

`builds/<build>/raceicons.json` is written by `python -m pipeline icons` beside
`racialicons.json`, keyed by `races.json`'s own slug, and the art is downloaded with the
rest, so a later run with no `raw/` reads the committed file back.
"""

from __future__ import annotations

import json
import logging
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.icons import icon_names

logger = logging.getLogger(__name__)

RACE_ICONS = "raceicons.json"

#: The raw client tables `race_icons` reads; `icons_for_build` skips the file without them.
RACE_TABLES = ("ChrRaces.csv", "ManifestInterfaceData.csv")

#: One portrait per race, the male one, everywhere the site shows a race.
RACE_ICON_PATTERN = "achievement_character_{race}_male"

#: The ChrRaces column naming the race whose model (and so portrait) the client falls back to.
FALLBACK_RACE_COLUMN = "MaleModelFallbackRaceID"


def _portrait(race_id: int, rows: dict[int, dict[str, str]], available: set[str]) -> str | None:
    """The race's own portrait, else its fallback race's, else None."""
    row = rows.get(race_id)
    if row is None:
        return None
    name = RACE_ICON_PATTERN.format(race=row["Name_lang"].replace(" ", "").lower())
    if name in available:
        return name
    fallback = int(row.get(FALLBACK_RACE_COLUMN) or 0)
    if fallback in (0, race_id):
        return None
    return _portrait(fallback, rows, available)


def race_icons(raw: Path, races: list[dict[str, object]]) -> dict[str, str]:
    """Race slug -> icon name for every race in `races` (races.json rows), sorted by slug.

    A race with no portrait anywhere along its fallback chain is logged and left out, so
    the site draws no icon for it rather than a placeholder."""
    rows = {int(row["ID"]): row for row in read_csv(raw / "ChrRaces.csv")}
    available = set(icon_names(read_csv(raw / "ManifestInterfaceData.csv")).values())
    icons: dict[str, str] = {}
    for race in races:
        slug, race_id = str(race["slug"]), int(str(race["id"]))
        icon = _portrait(race_id, rows, available)
        if icon is None:
            logger.warning("race %s (%s) has no portrait icon in the client", race_id, slug)
            continue
        icons[slug] = icon
    return dict(sorted(icons.items()))


def write_race_icons(build: str, root: Path = Path("builds")) -> Path:
    """Write `builds/<build>/raceicons.json` from the build's `raw/` tables and races.json."""
    build_dir = root / build
    races = json.loads((build_dir / "races.json").read_text(encoding="utf-8"))
    path = build_dir / RACE_ICONS
    payload = race_icons(build_dir / "raw", races)
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    return path
