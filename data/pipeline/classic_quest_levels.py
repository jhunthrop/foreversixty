"""The open-source Classic `quest_template` table as one source for a
quest's own min_level/level -- the fields the client's per-item
`required_level` (almost always 0 for a quest reward) does not carry.

2026-09-28 quest-levels lane, primary source (see `pipeline.quest_levels`
for the merged file this feeds and `pipeline.wowhead_quests` for the
secondary source it does not cover). Two facts drove the plan:

* The Forever client itself states no quest level anywhere `wago.tools`
  exports: `QuestV2` carries only `ID`/`UniqueBitFlag`/
  `UiQuestDetailsThemeID`, and neither `QuestInfo` nor `QuestXP` carries
  one either. A quest's level is server-side data on every WoW version,
  Forever included.
* Wowhead throttles this pipeline's address (`pipeline.wowhead_quests`'s
  own doc: a 403 storm at id 873, again at id 9248 even after backing
  off to 4s between requests -- a shared-IP limit, not something a
  slower crawl alone outlasts).

cmangos/classic-db (github.com/cmangos/classic-db, GPL-3.0) ships a
`quest_template` table with `MinLevel`/`QuestLevel` columns for 1.12
Classic's own ~4,245 quests. Forever's quest ids match 1.12's for every
quest that existed then (both are Blizzard's own ids, unchanged since
vanilla) -- confirmed for this build: of the 720 quest ids
`data/builds/1.60.1.70009/loot.json` needs, 718 (99.7%) are in this
table. The owner's own count (Forever adds 1,000+ NEW quests overall)
is about the FULL 5,198-quest Forever catalog, not this build's actual
quest-reward set, which is old 1.12 content almost entirely -- a
Forever-new quest (not in 1.12, 2 of this build's 720) has no row here
and needs `pipeline.wowhead_quests` (or, failing that too, the
item_level_proxy fallback) instead.

Cross-checked against 414 wowhead pages this lane fetched: level agrees
on 406/414 (98%), min_level on 347/414 (84%, Forever's own rebalancing
of a handful of quests -- Zul'Gurub's entry-level quests moved from 58
to 60, for one -- accounts for most of the gap). See
`data/tests/test_classic_quest_levels.py`'s own sample fixtures for the
exact mismatches.

This module is fetch-and-parse ONLY: `fetch_classic_db_quest_levels`
downloads the pinned dump and returns the parsed table in memory. It
writes nothing -- `pipeline.quest_levels.merge_classic_db` (the
`fetch-classic-quest-levels` CLI step) is what merges the result into
the build's committed, provenance-headered `quest-levels.json`.
"""

from __future__ import annotations

import gzip
import re

import httpx
from pydantic import BaseModel

from pipeline.wago import USER_AGENT

#: Pinned per the controller's own instruction: a raw.githubusercontent.com
#: URL naming an exact commit, never a moving branch ref, so a rerun of
#: `fetch-classic-quest-levels` is reproducible until deliberately
#: repinned (bumping SOURCE_COMMIT here is the only way this module's
#: output ever changes).
SOURCE_REPO = "cmangos/classic-db"
SOURCE_COMMIT = "ec4f596146be6467ea93c57397858e329e2db852"
SOURCE_FILE = "Full_DB/ClassicDB_1_12_1_z2815.sql.gz"
SOURCE_LICENSE = "GPL-3.0"
SOURCE_URL = f"https://raw.githubusercontent.com/{SOURCE_REPO}/{SOURCE_COMMIT}/{SOURCE_FILE}"

#: The fallback level for a quest (or crafted item) neither source
#: states a real level for: `min(60, item_level - 5)`, floored at 0.
#: This lane's brief's own formula -- a quest reward's item level tends
#: to run a few points above the level a character obtains it at, and
#: 60 is as high as a leveling character's gate is ever set.
MAX_QUEST_LEVEL = 60


def item_level_proxy(item_level: int) -> int:
    return max(0, min(MAX_QUEST_LEVEL, item_level - 5))


class ClassicQuestLevel(BaseModel):
    """One quest's level fields, from cmangos/classic-db's
    `quest_template` table (`entry`, `MinLevel`, `QuestLevel`)."""

    entry: int
    min_level: int
    level: int


# --- Parsing a mysqldump's quest_template INSERT statement(s) ---------
#
# The dump is one `CREATE TABLE` plus one or more "INSERT INTO
# `quest_template` VALUES (...),(...),...;" statements. A regex over
# the raw text cannot safely find where one row ends and the next
# begins: every row's tail is free-form quest text (Title, Details,
# Objectives, ...) that can itself contain literal commas, parentheses
# and semicolons inside single-quoted strings. This lane's first attempt
# -- a regex for "six leading integers in parens" run over the whole
# file -- found 180,935 matches for what turned out to be 4,245 real
# quests: the same shape recurs by chance inside OTHER tables' data
# throughout a 73MB dump. The fix is the same one a real SQL parser
# uses: track single-quote state (with backslash-escaping) and paren
# depth char by char, so a row boundary is only ever recognized OUTSIDE
# a quoted string.
#
# It only needs to track enough to find each row's boundaries, though
# -- once a row is isolated, `quest_template`'s own column order
# (`entry`, `Method`, `ZoneOrSort`, `MinLevel`, `MaxLevel`, `QuestLevel`,
# ...) puts every field this module wants in the first SIX columns,
# all plain integers, before the first text field (`Title`, column 31)
# -- so a plain regex on the row's own (now-safe) leading substring is
# enough to read them.
_ROW_PREFIX = re.compile(r"^(\d+),(\d+),(-?\d+),(\d+),(\d+),(-?\d+),")
_INSERT_MARKER = "INSERT INTO `quest_template` VALUES"


def _statement_span(text: str, marker_start: int) -> str:
    """The full "(...),(...),..." value list of one INSERT statement
    starting at `marker_start` (the index `_INSERT_MARKER` was found
    at), stopping at the statement's own top-level `;` -- one outside
    any quoted string and at paren depth 0."""
    i = marker_start + len(_INSERT_MARKER)
    while text[i] in " \n\t":
        i += 1
    depth = 0
    in_string = False
    j = i
    n = len(text)
    while j < n:
        c = text[j]
        if in_string:
            if c == "\\":
                j += 2
                continue
            if c == "'":
                in_string = False
        else:
            if c == "'":
                in_string = True
            elif c == "(":
                depth += 1
            elif c == ")":
                depth -= 1
            elif c == ";" and depth == 0:
                return text[i:j]
        j += 1
    raise ValueError("quest_template INSERT statement has no top-level terminator")


def _split_rows(statement: str) -> list[str]:
    """Every "(...)" row in one INSERT statement's value list, quote-aware
    (see this module's own parsing doc)."""
    rows: list[str] = []
    depth = 0
    in_string = False
    row_start = 0
    n = len(statement)
    j = 0
    while j < n:
        c = statement[j]
        if in_string:
            if c == "\\":
                j += 2
                continue
            if c == "'":
                in_string = False
        else:
            if c == "'":
                in_string = True
            elif c == "(":
                if depth == 0:
                    row_start = j + 1
                depth += 1
            elif c == ")":
                depth -= 1
                if depth == 0:
                    rows.append(statement[row_start:j])
        j += 1
    return rows


def parse_quest_template(sql_text: str) -> dict[int, ClassicQuestLevel]:
    """Every `quest_template` row's entry/MinLevel/QuestLevel, from one or
    more "INSERT INTO `quest_template` VALUES ...;" statements anywhere
    in a mysqldump's text (cmangos/classic-db's own full-DB dump splits
    large tables across several INSERT statements)."""
    out: dict[int, ClassicQuestLevel] = {}
    start = sql_text.find(_INSERT_MARKER)
    while start != -1:
        statement = _statement_span(sql_text, start)
        for row in _split_rows(statement):
            match = _ROW_PREFIX.match(row)
            if not match:
                raise ValueError(f"quest_template row does not start with 6 integers: {row[:60]!r}")
            entry, _method, _zone, min_level, _max_level, level = match.groups()
            out[int(entry)] = ClassicQuestLevel(
                entry=int(entry), min_level=int(min_level), level=int(level)
            )
        start = sql_text.find(_INSERT_MARKER, start + len(_INSERT_MARKER))
    return out


def fetch_classic_db_quest_levels(
    client: httpx.Client | None = None, url: str = SOURCE_URL
) -> dict[int, ClassicQuestLevel]:
    """Download the pinned cmangos/classic-db dump and extract
    `quest_template`'s entry/MinLevel/QuestLevel. Gzip-decodes the
    response itself (httpx does not auto-decode a `.gz` file's content,
    unlike a `Content-Encoding: gzip` response) so the caller never
    needs a second library for it. Writes nothing -- see
    `pipeline.quest_levels.merge_classic_db` for the writer.
    """
    own = client is None
    client = client or httpx.Client(headers={"User-Agent": USER_AGENT})
    try:
        response = client.get(url, timeout=120, follow_redirects=True)
        response.raise_for_status()
        sql_text = gzip.decompress(response.content).decode("utf-8", errors="replace")
        return parse_quest_template(sql_text)
    finally:
        if own:
            client.close()
