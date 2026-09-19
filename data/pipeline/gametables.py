"""The client's game tables, which the engine's base stats are parsed from.

Vanilla's per-class base mana and its combat-rating conversions are not in
DB2. Probed against build 1.60.1.69893: GameTables, gtCombatRatings,
CombatRatings, GtChanceToMeleeCrit, GtOCTRegenHP, BaseMP, CharBaseStats and
CombatRatingsMultByILvl all 404, and CharBaseInfo is only the race/class
validity matrix. They live in the client's `GameTables/*.txt`, which is what
the engine's own `tools/base_stats_parser.py` reads, and they come down the
same road the talent-tree art does: by CASC file data id, through
`pipeline.casc.fetch_casc_file`, pinned to the build.

**Six of the seven files that parser names are not in this client.** Pinned
to either Classic-lineage build, `octbasempbyclass.txt` (4049853),
`chancetomeleecrit.txt` (3999262), `chancetomeleecritbase.txt` (3999263),
`chancetospellcrit.txt` (3999265), `chancetospellcritbase.txt` (3999264) and
`octclasscombatratingscalar.txt` (4526467) all answer 200 with an empty body,
on 1.60.1.69893 and on 1.15.9.69722 alike, while `combatratings.txt`
(1391669) returns 12,144 and 31,614 bytes respectively. The dividing line is
the id range: the Classic lineage serves the original ~1.39M block and
nothing a later retail expansion added. An unpinned request returned those
six from retail, which is how an earlier draft of this module came to treat
retail's stat curves as Forever's.

So this emits the three the lineage does serve -- `combatratings.txt`,
`basemp.txt` (the lineage's base-mana table, and the only stand-in for the
absent `octbasempbyclass.txt`) and `hppersta.txt` -- and keeps the six absent
ones as `ABSENT_FROM_THE_CLASSIC_LINEAGE`, so the engine lane reads the gap
off a constant with a test behind it instead of inheriting Era's numbers
without knowing.

`spellscaling.txt` (1391660) and `combatratingsmultbyilvl.txt` (1391670) are
served too and deliberately left out: nothing needs them, and vanilla spell
coefficients are a convention the engine owns rather than a table
(research/07-simulator.md 5.3).

Per-race base Strength, Agility, Stamina, Intellect and Spirit are in neither
DB2 nor GameTables -- the listfile's whole gametables/ has no
base-primary-stat file -- and stay with the engine lane.
"""

from __future__ import annotations

import logging
import shutil
from dataclasses import dataclass
from pathlib import Path

import httpx

from pipeline.casc import fetch_casc_file
from pipeline.icons import CACHE_DIR, _atomic_write
from pipeline.manifest import refresh_manifest
from pipeline.wago import BASE_URL, USER_AGENT

logger = logging.getLogger(__name__)

GAMETABLES = "gametables"

#: File name -> CASC file data id, for the game tables build 1.60.1.69893
#: actually serves. Ids from wowdev/wow-listfile's community-listfile.csv,
#: each verified 200 with content for this build.
GAME_TABLE_IDS: dict[str, int] = {
    "basemp.txt": 1391664,
    "combatratings.txt": 1391669,
    "hppersta.txt": 1391642,
}

#: The six files `tools/base_stats_parser.py` names that this client does not
#: ship: each answers 200 with an empty body, on 1.60.1.69893 and on Classic
#: Era alike. Kept as data, not as a comment, because the engine lane needs to
#: know which of its inputs no Forever client can supply -- and because a
#: later build that starts shipping one should show up as a test failure here
#: rather than as a number nobody re-derived.
ABSENT_FROM_THE_CLASSIC_LINEAGE: dict[str, int] = {
    "chancetomeleecrit.txt": 3999262,
    "chancetomeleecritbase.txt": 3999263,
    "chancetospellcrit.txt": 3999265,
    "chancetospellcritbase.txt": 3999264,
    "octbasempbyclass.txt": 4049853,
    "octclasscombatratingscalar.txt": 4526467,
}

_SEPARATOR = "\t"


class GameTableError(ValueError):
    """A game table is not the shape the engine's parser reads."""


@dataclass(frozen=True)
class GameTable:
    """One `GameTables/*.txt`: a header row, then one row per key."""

    name: str
    columns: tuple[str, ...]
    rows: dict[str, tuple[str, ...]]


def parse_game_table(name: str, text: str) -> GameTable:
    """Parse a tab-separated game table, keyed by each row's first column.

    Used only to check what was downloaded -- the bytes are written through
    unchanged -- but a table that does not parse here is one the engine's own
    `csv.reader(delimiter="\\t")` would misread, which is worth catching at
    fetch time rather than in a Go build.
    """
    lines = [line for line in text.replace("\r\n", "\n").split("\n") if line]
    if not lines:
        raise GameTableError(f"{name} is empty")
    header = tuple(lines[0].split(_SEPARATOR))
    if len(header) < 2:
        raise GameTableError(f"{name} has a {len(header)}-column header, not a game table")
    rows: dict[str, tuple[str, ...]] = {}
    for line in lines[1:]:
        cells = tuple(line.split(_SEPARATOR))
        if len(cells) != len(header):
            raise GameTableError(
                f"{name} row {cells[0]!r} has {len(cells)} cells against {len(header)} columns"
            )
        rows[cells[0]] = cells[1:]
    if not rows:
        raise GameTableError(f"{name} has a header and no rows")
    return GameTable(name=name, columns=header, rows=rows)


def write_game_tables(
    build: str,
    root: Path = Path("builds"),
    cache_dir: Path = CACHE_DIR,
    client: httpx.Client | None = None,
) -> Path:
    """Fetch this build's game tables and write them under its build directory.

    The bytes are written exactly as the client ships them, CRLF and all: the
    engine parses these files directly, so re-serialising a float or
    normalising a line ending would make them our numbers rather than the
    client's.
    """
    build_dir = root / build
    if not (build_dir / "manifest.json").exists():
        raise SystemExit(
            f"no manifest in {build_dir}; run `python -m pipeline normalize` for it first"
        )
    out = build_dir / GAMETABLES
    # Rebuilt from scratch so a table that leaves GAME_TABLE_IDS does not
    # leave a stale file behind, exactly as simconst does for spellconst/.
    shutil.rmtree(out, ignore_errors=True)
    out.mkdir(parents=True, exist_ok=True)
    own = client is None
    if client is None:
        client = httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    try:
        for name, file_id in sorted(GAME_TABLE_IDS.items()):
            data = fetch_casc_file(
                file_id, build, cache_dir=cache_dir, client=client, suffix=".txt"
            )
            table = parse_game_table(name, data.decode("utf-8"))
            logger.info("%s: %d columns, %d rows", name, len(table.columns), len(table.rows))
            _atomic_write(out / name, data)
    finally:
        if own:
            client.close()
    refresh_manifest(build_dir)
    logger.info("wrote %s: %d game tables", out, len(GAME_TABLE_IDS))
    return out
