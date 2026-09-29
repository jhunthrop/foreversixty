"""cmangos/classic-db (the same pinned dump `pipeline.classic_quest_levels`
already uses for quest levels) as a third `loot.json` source, between the
engine fork's own database (primary) and wowhead's item-page scrape
(lowest priority, night-item-sources lane). Src-classicdb lane, 2026-09-29
-- see the lane brief for the coverage measurement this closes: union of
creature+reference+object+vendor+quest drops from the dump covers 2,607 of
this build's ~2,900 Classic-id items the fork+wowhead pair leaves
unsourced.

Same split as `pipeline.item_sources`/`pipeline.quest_levels`: this module
is fetch-and-PARSE only. `fetch_classic_db_sources` downloads the pinned
dump (reusing `classic_quest_levels`' own pinned commit and URL) and
returns every table this lane needs, parsed into `ClassicDbSourceRecord`
rows keyed by item id. `write_classic_sources`/`load_classic_sources` are
the committed-cache reader/writer
(`data/builds/<build>/raw/classicdb/sources.json`) that
`pipeline.loot.classicdb.classicdb_additions` reads OFFLINE -- `loot`
itself never touches the network, the same "nightly fetches, CI reads a
committed file" contract every other raw/ cache in this pipeline keeps.

Tables read, and what this module does with each (see the lane brief for
the schema citations):

* `creature_loot_template` / `reference_loot_template`: a creature's own
  drops. A NEGATIVE `mincountOrRef` means the row is itself a reference
  into `reference_loot_template` (bounded, cycle-safe recursion --
  `_MAX_REFERENCE_DEPTH`), not a real item id.
* `gameobject_loot_template` / `gameobject_template`: chest drops, named
  via the object's own template row.
* `npc_vendor` + `npc_vendor_template` (joined through
  `creature_template.VendorTemplateId`, cmangos' own indirection for a
  vendor list shared by more than one creature entry) + `conditions`
  (type 5, `CONDITION_REPUTATION_RANK`, for a rep-gated vendor slot).
* `skinning_loot_template` / `pickpocketing_loot_template`: classified
  `world`, same as a plain creature drop, since the item comes off the
  same creature (no separate name/zone facts to draw on).
* `fishing_loot_template`: no creature to name (the entry is a fishing
  loot GROUP id, its own comments naming a list of AreaTable zone ids as
  free text) -- classified `world` under a flat "Fishing" bucket.
* `quest_template`: `RewChoiceItemId1-6`/`RewItemId1-4` for the reward
  items, `MinLevel`/`QuestLevel` for the quest's own level (the same two
  fields `classic_quest_levels.parse_quest_template` reads, reused here
  as `pipeline.quest_levels.QuestLevelEntry`-shaped data on each record
  rather than re-fetched), and `RequiredRaces` for the faction the
  reward is restricted to (see `_faction_from_required_races`).

A creature/object's own MAP id (cmangos' `creature.map`/`gameobject.map`,
the spawn table, aggregated per template entry by `Counter.most_common`)
is kept on the record as-is; resolving it to one of this build's own
dungeon/raid zone ids is `pipeline.loot.classicdb`'s job (it has `zones.
json`+`types`, this module does not). An open-world creature's map is
almost always continent id 0 or 1, which is not a specific zone at all --
this module makes no attempt to guess one; a source whose map does not
resolve to a known instance becomes a flat `world:<name>` bucket, per the
lane brief's own instruction.
"""

from __future__ import annotations

import gzip
import json
import logging
import re
from collections import Counter, defaultdict
from datetime import UTC, datetime
from pathlib import Path
from typing import Literal

import httpx
from pydantic import BaseModel

from pipeline.classic_quest_levels import SOURCE_URL
from pipeline.forkdb import REP_LEVELS, decode
from pipeline.sqldump import iter_table_records, unquote
from pipeline.wago import USER_AGENT

logger = logging.getLogger(__name__)

FILE_NAME = "sources.json"

#: cmangos' own `CONDITION_REPUTATION_RANK` type id (see the dump's
#: `conditions` table: row `(71,5,749,6,...)` reads "Has Minimum Rank
#: Revered With Faction ID: 749" -- value1 is the faction id, value2 the
#: rank, 0-indexed Hated..Exalted, one below `pipeline.forkdb.REP_LEVELS`'
#: own 1-indexed Hated..Exalted, hence the `+ 1` everywhere this is read).
_CONDITION_REPUTATION_RANK = 5

#: A loot-template reference chain longer than this is almost certainly a
#: cycle in the data (or a mistake in this parser) rather than a real
#: nested loot group -- cut it off rather than recurse forever. No chain
#: in the pinned dump is more than 2 deep in practice.
_MAX_REFERENCE_DEPTH = 5

#: cmangos' own tag for a GENERIC, level-gated "world drop" pool -- any
#: creature in the right level range can drop from the Grey/Green/Blue/
#: Purple/Epic tier this reference group represents, so it says nothing
#: specific about the one creature whose row happens to point at it.
#: Confirmed present on every comment for the outlier groups below
#: (2026-09-29 measurement on the pinned dump).
_WORLD_DROP_MARKER = "World Drop"

#: A reference group used by more than this many DISTINCT creature/
#: object/skinning/pickpocketing/fishing entries is, empirically, always
#: one of two things: a `_WORLD_DROP_MARKER` pool (already excluded by
#: name), or another game-wide generic pool that happens not to say
#: "World Drop" in its own comment (id 60446, "16 Slot Bag - NPC Levels:
#: 48+", shared by 777 creatures, is the one such case on the pinned
#: dump). A real PER-ZONE shared trash table -- Shadowfang Keep's own
#: "Zone Drop" reference, 27 users -- stays comfortably under this.
#: Without this cut, expanding every reference blindly turned ONE
#: Onyxia's Lair trash npc's own boss listing into 300+ items that have
#: nothing to do with the raid (measured while regenerating loot.json
#: for 1.60.1.70009, src-classicdb lane).
_SHARED_REFERENCE_MAX_USERS = 50

#: cmangos' own tag for a real PER-ZONE/PER-INSTANCE shared trash table
#: (Shadowfang Keep's, Gnomeregan's, Scarlet Monastery's, Blackfathom
#: Deeps' and Razorfen Downs' own "Zone Drop" references on the pinned
#: dump, 22-38 users each) -- a DIFFERENT, hand-authored cmangos label
#: from `_WORLD_DROP_MARKER`'s "World Drop", naming a table that is
#: still specific to the one instance it drops in (not bind-on-equip,
#: not auction-housable), even when the dump's own stale/duplicate
#: `map` column (the same Onyxia's-Lair-continent-id quirk
#: `instance_zone_by_map`'s own doc measures) makes `_world_drop_pools`'
#: distinct-map count for it look like more than one map. World-drop-
#: pool lane, 2026-09-29: this marker exempts a "Zone Drop" pool from
#: `_MULTI_MAP_MIN_MAPS`'s own signal (world-drop-pool lane's own
#: measurement: without this exemption, Gnomeregan's, Scarlet
#: Monastery's and Blackfathom Deeps' own "Zone Drop" pools -- real,
#: per-instance trash tables, not auction-housable world drops --
#: false-positive on 2 maps each).
_ZONE_DROP_MARKER = "Zone Drop"

#: A reference group whose creature/object/skinning/pickpocketing users
#: (NOT fishing -- see `_world_drop_pools`' own doc) spawn on this many
#: or more DISTINCT maps is a generic pool no single instance or open-
#: world zone owns, whatever its own comment says -- world-drop-pool
#: lane, 2026-09-29's own measurement: 15 reference ids on the pinned
#: dump (mostly "NPC LOOT ... Classic World bosses and dragons" and a
#: handful of profession-recipe tables shared across two Ahn'Qiraj wings)
#: cross this without EITHER carrying `_WORLD_DROP_MARKER`'s own text or
#: exceeding `_SHARED_REFERENCE_MAX_USERS`' own fan-out count -- exactly
#: the gap this lane's brief calls out (Lambent Scale Cloak, item 4706:
#: 11 narrow-banded reference ids, each well under 50 users and on a
#: single map on its own, together spanning Gnomeregan, Scarlet
#: Monastery, Shadowfang Keep and The Stockade once a stale dump-map
#: quirk is set aside -- see `_ZONE_DROP_MARKER`'s own doc for why that
#: quirk is excluded here rather than trusted).
_MULTI_MAP_MIN_MAPS = 2

#: `NPC Levels: <lo>[-<hi>]` or `Item Levels: <lo>[-<hi>]` inside a
#: `_WORLD_DROP_MARKER`-tagged row's own `comments` text (the pinned
#: dump's own convention, e.g. "NPC LOOT (Green World Drop) - (Item
#: Levels: 20-25) - (NPC Levels: 21-22)") -- `_level_range_from_comment`
#: prefers the NPC figure (closer to "the pool's own creature level
#: range", this lane's brief's own second option, than the item-level
#: figure is) and falls back to the item figure only when a row states
#: no NPC range at all.
_NPC_LEVEL_PATTERN = re.compile(r"NPC\s+Levels?:?\s*(\d+)(?:\s*-\s*(\d+))?")
_ITEM_LEVEL_PATTERN = re.compile(r"Item\s+Levels?:?\s*(\d+)(?:\s*-\s*(\d+))?")


def _level_range_from_comment(comment: str) -> tuple[int, int] | None:
    match = _NPC_LEVEL_PATTERN.search(comment) or _ITEM_LEVEL_PATTERN.search(comment)
    if match is None:
        return None
    low = int(match.group(1))
    high = int(match.group(2)) if match.group(2) else low
    return low, high


class WorldDropPool(BaseModel):
    """One `reference_loot_template` id `_world_drop_pools` classifies as
    a generic world-drop pool, with the level range its own referencing
    rows' `comments` state (`_level_range_from_comment`), widest first --
    `None` on either side when no referencing row's comment states one at
    all (kept as a real, honestly-level-less world drop rather than
    invented, per this file's own no-fabrication rule elsewhere)."""

    level_min: int | None = None
    level_max: int | None = None

#: Vanilla's own race bitmask (`ChrRaces.dbc` ids, 1-indexed bit
#: positions) split by faction, for `quest_template.RequiredRaces`. A
#: quest that admits races from both sides (or states none, 0, meaning
#: "any race") is `"both"`; a quest whose bits are one side only names
#: that side.
_ALLIANCE_RACE_MASK = 1 | 4 | 8 | 64  # Human, Dwarf, Night Elf, Gnome
_HORDE_RACE_MASK = 2 | 16 | 32 | 128  # Orc, Undead, Tauren, Troll


def _faction_from_required_races(races: int) -> Literal["alliance", "horde", "both"]:
    alliance = bool(races & _ALLIANCE_RACE_MASK)
    horde = bool(races & _HORDE_RACE_MASK)
    if alliance and not horde:
        return "alliance"
    if horde and not alliance:
        return "horde"
    return "both"


ClassicDbSourceKind = Literal[
    "creature_drop", "object_drop", "vendor", "quest_reward", "skinning", "pickpocketing", "fishing",
    "world_drop",
]  # fmt: skip


class ClassicDbCondition(BaseModel):
    """A rep-gated vendor slot's own condition row (`conditions` type 5)."""

    faction_id: int
    standing: str


class ClassicDbQuestInfo(BaseModel):
    quest_id: int
    min_level: int
    level: int
    faction: Literal["alliance", "horde", "both"]


class ClassicDbSourceRecord(BaseModel):
    """One item id's one way to get it, from the pinned classic-db dump.
    `pipeline.loot.classicdb.classicdb_additions` is what turns a list of
    these (per item id) into `LootSource`/`QuestSource` rows."""

    kind: ClassicDbSourceKind
    npc_id: int | None = None
    object_id: int | None = None
    #: The creature/object/quest name -- never the generic "Fishing"
    #: placeholder's npc, since fishing has none.
    name: str
    #: The spawn table's own map id (cmangos `creature.map`/
    #: `gameobject.map`), for `creature_drop`/`object_drop`/`skinning`/
    #: `pickpocketing` only.
    map_id: int | None = None
    #: Percent chance (0-100), `abs(ChanceOrQuestChance)` -- cmangos uses
    #: a negative value for "grouped, no other normal loot" bookkeeping
    #: this pipeline does not model; the magnitude is still the real
    #: percentage a player sees.
    chance: float | None = None
    condition: ClassicDbCondition | None = None
    quest: ClassicDbQuestInfo | None = None
    #: `world_drop` only -- the pool's own level range (`WorldDropPool`),
    #: merged across every reference id that names this exact item (a
    #: green-quality world drop is typically split across several
    #: narrow-banded reference ids on the pinned dump -- 11 of them for
    #: Lambent Scale Cloak alone -- that all describe the one real item;
    #: `_world_drop_records` merges them into the one honest range a
    #: player actually sees rather than one near-duplicate source per
    #: narrow band). `None` on either side when no contributing pool
    #: states a level at all.
    level_min: int | None = None
    level_max: int | None = None


def _rows_by_entry(records: list[dict[str, str]]) -> dict[int, list[dict[str, str]]]:
    out: dict[int, list[dict[str, str]]] = defaultdict(list)
    for row in records:
        out[int(row["entry"])].append(row)
    return out


def _excluded_reference_ids(sql_text: str) -> set[int]:
    """Every `reference_loot_template` id a NEGATIVE `mincountOrRef` row
    anywhere in the dump points at, that this lane treats as a generic
    pool rather than that row's own creature/object's specific loot --
    see `_WORLD_DROP_MARKER`/`_SHARED_REFERENCE_MAX_USERS`' own docs for
    the two signals and why each is needed (neither alone covers every
    case on the pinned dump)."""
    users: dict[int, set[tuple[str, int]]] = defaultdict(set)
    world_drop: set[int] = set()
    for table in (
        "creature_loot_template", "gameobject_loot_template",
        "skinning_loot_template", "pickpocketing_loot_template", "fishing_loot_template",
    ):
        for row in iter_table_records(sql_text, table):
            min_ref = int(row["mincountOrRef"])
            if min_ref >= 0:
                continue
            ref_id = -min_ref
            users[ref_id].add((table, int(row["entry"])))
            if _WORLD_DROP_MARKER in row["comments"]:
                world_drop.add(ref_id)
    shared = {ref_id for ref_id, rows in users.items() if len(rows) > _SHARED_REFERENCE_MAX_USERS}
    return world_drop | shared


def _world_drop_pools(
    sql_text: str, npc_map: dict[int, int], object_map: dict[int, int]
) -> dict[int, WorldDropPool]:
    """Every `reference_loot_template` id this lane classifies as a
    generic world-drop pool -- one this lane's brief asks be recorded
    under a single synthetic `world_drop` source (`_world_drop_records`),
    rather than attributed to each creature/object that happens to point
    at it (`_excluded_reference_ids`'s own, coarser rule: drop it
    entirely). World-drop-pool lane, 2026-09-29's own measurement on the
    pinned dump, three signals (a pool matching any one is classified):

    * `_WORLD_DROP_MARKER` in a referencing row's own `comments` -- 309
      reference ids, the same set `_excluded_reference_ids` already
      caught by name.
    * More than `_SHARED_REFERENCE_MAX_USERS` distinct referencing
      creature/object/skinning/pickpocketing entries -- kept at 50, not
      lowered to this lane's brief's own proposed 8: Shadowfang Keep's
      real, single-map "Zone Drop" pool has exactly 27 (this file's own
      `_SHARED_REFERENCE_MAX_USERS` doc), which 8 would misclassify as a
      world drop even though it is a real per-instance trash table, not
      an auction-housable one.
    * Referencing entries spawning on `_MULTI_MAP_MIN_MAPS` or more
      DISTINCT maps, UNLESS every referencing comment carries
      `_ZONE_DROP_MARKER` (`_ZONE_DROP_MARKER`'s own doc: a stale/
      duplicate map id, the same quirk `instance_zone_by_map` already
      guards against, would otherwise misclassify Gnomeregan's, Scarlet
      Monastery's and Blackfathom Deeps' own real "Zone Drop" pools too).
      This is this lane's brief's own second proposed signal, and the one
      that actually explains the reported defect: Lambent Scale Cloak's
      11 reference ids (60125-60135) are each under 50 users and marked
      `_WORLD_DROP_MARKER` already (so already excluded, just not
      bucketed) -- a pool with NEITHER signal but a real multi-instance
      fan-out (Gnomeregan + Scarlet Monastery + Shadowfang Keep + The
      Stockade, in the reported screenshot's own case) is exactly what
      this third signal is for.

    Deliberately does NOT implement this lane's brief's own third
    proposed signal ("rows all `ChanceOrQuestChance` <= 0.5 with no boss
    in it"): "boss" is a dungeon/raid zone-type fact
    (`pipeline.loot.classicdb`'s own `zones.json`/`types`), not something
    this module -- pure SQL-dump parsing, no zone data -- can answer.
    Left for a follow-up that thread that data through, rather than
    guessed at here.

    `fishing_loot_template` is deliberately excluded from every signal
    here (unlike `_excluded_reference_ids`, which still covers it for
    that function's own, unrelated caller): a fishing pool spanning many
    zones is fishing working as designed, not a symptom of anything, and
    `pipeline.loot.classicdb._parse_fishing`'s own flat "Fishing" bucket
    already names it honestly regardless of which reference id
    contributed the catch.
    """
    users: dict[int, set[tuple[str, int]]] = defaultdict(set)
    comments: dict[int, list[str]] = defaultdict(list)
    for table in (
        "creature_loot_template", "gameobject_loot_template",
        "skinning_loot_template", "pickpocketing_loot_template",
    ):
        for row in iter_table_records(sql_text, table):
            min_ref = int(row["mincountOrRef"])
            if min_ref >= 0:
                continue
            ref_id = -min_ref
            users[ref_id].add((table, int(row["entry"])))
            comments[ref_id].append(row["comments"])

    pools: dict[int, WorldDropPool] = {}
    for ref_id, entries in users.items():
        comment_texts = comments[ref_id]
        marker = any(_WORLD_DROP_MARKER in text for text in comment_texts)
        zone_drop = any(_ZONE_DROP_MARKER in text for text in comment_texts)
        maps = {
            (object_map if table == "gameobject_loot_template" else npc_map).get(entry)
            for table, entry in entries
        } - {None}
        multi_map = len(maps) >= _MULTI_MAP_MIN_MAPS and not zone_drop
        big_fan_out = len(entries) > _SHARED_REFERENCE_MAX_USERS
        if not (marker or multi_map or big_fan_out):
            continue
        ranges = [
            level_range
            for text in comment_texts
            if (level_range := _level_range_from_comment(text)) is not None
        ]
        pools[ref_id] = WorldDropPool(
            level_min=min((lo for lo, _ in ranges), default=None),
            level_max=max((hi for _, hi in ranges), default=None),
        )
    return pools


def _world_drop_records(
    sql_text: str, pools: dict[int, WorldDropPool]
) -> dict[int, ClassicDbSourceRecord]:
    """One `world_drop` `ClassicDbSourceRecord` per item id any pool in
    `pools` (`_world_drop_pools`' own result) directly names, its own
    `level_min`/`level_max` the widest range across every pool that names
    it (this function's own doc reason: the pinned dump splits one real
    item's world-drop pool across several narrow-banded reference ids far
    more often than not) and its own `chance` the highest of theirs (a
    generic pool's per-item chance is itself approximate -- cmangos'
    own reference rows record it as the item's weight within ONE narrow
    band, not the real overall drop chance across every band that
    carries it -- so the highest band's figure is kept as the more
    honest of several approximations, never averaged or invented).

    Only a pool's own DIRECT rows are read (a POSITIVE `mincountOrRef`)
    -- no pool in the pinned dump references another one, so the
    recursive expansion `_expand_loot_template` needs for a creature's
    own drop list is not needed here.
    """
    reference_rows = _rows_by_entry(list(iter_table_records(sql_text, "reference_loot_template")))
    levels: dict[int, tuple[int | None, int | None]] = {}
    chances: dict[int, float] = {}
    for ref_id, pool in pools.items():
        for row in reference_rows.get(ref_id, []):
            if int(row["mincountOrRef"]) < 0:
                continue
            item_id = int(row["item"])
            chance = abs(float(row["ChanceOrQuestChance"]))
            chances[item_id] = max(chances.get(item_id, 0.0), chance)
            lo, hi = levels.get(item_id, (None, None))
            if pool.level_min is not None:
                lo = pool.level_min if lo is None else min(lo, pool.level_min)
            if pool.level_max is not None:
                hi = pool.level_max if hi is None else max(hi, pool.level_max)
            levels[item_id] = (lo, hi)
    return {
        item_id: ClassicDbSourceRecord(
            kind="world_drop", name="World drop", chance=chances.get(item_id),
            level_min=lo, level_max=hi,
        )
        for item_id, (lo, hi) in levels.items()
    }  # fmt: skip


def _expand_loot_template(
    entry: int,
    rows_by_entry: dict[int, list[dict[str, str]]],
    reference_rows_by_entry: dict[int, list[dict[str, str]]],
    excluded_refs: frozenset[int],
    *,
    depth: int = 0,
    seen: frozenset[int] = frozenset(),
) -> list[tuple[int, float]]:
    """`entry`'s own (item id, percent chance) pairs from a
    `*_loot_template` table, expanding a NEGATIVE `mincountOrRef` row
    into `reference_loot_template[-mincountOrRef]`'s own rows
    (recursively -- a reference can itself reference another), bounded by
    `_MAX_REFERENCE_DEPTH` and a `seen` set so a cycle in the data
    degrades to "stop expanding" rather than an infinite loop. A row
    pointing at an id in `excluded_refs` (a generic pool, not this
    entry's own specific loot -- `_excluded_reference_ids`' own doc)
    contributes nothing at all, at any depth.
    """
    if depth > _MAX_REFERENCE_DEPTH or entry in seen:
        return []
    seen = seen | {entry}
    out: list[tuple[int, float]] = []
    for row in rows_by_entry.get(entry, []):
        min_ref = int(row["mincountOrRef"])
        chance = abs(float(row["ChanceOrQuestChance"]))
        if min_ref < 0:
            ref_id = -min_ref
            if ref_id in excluded_refs:
                continue
            out.extend(
                _expand_loot_template(
                    ref_id, reference_rows_by_entry, reference_rows_by_entry, excluded_refs,
                    depth=depth + 1, seen=seen,
                )
            )
        else:
            out.append((int(row["item"]), chance))
    return out


def _spawn_map_by_entry(spawn_records: list[dict[str, str]]) -> dict[int, int]:
    """template entry id -> the map id it spawns on most often
    (`Counter.most_common`) -- almost always unambiguous for an
    instance-bound creature/object; an open-world one may spawn on
    several continents' worth of copies, in which case the most common
    single map id stands in (this module never needs more precision than
    "is this an instance map or not", which `pipeline.loot.classicdb`
    checks afterwards).
    """
    by_entry: dict[int, Counter[int]] = defaultdict(Counter)
    for row in spawn_records:
        by_entry[int(row["id"])][int(row["map"])] += 1
    return {entry: counts.most_common(1)[0][0] for entry, counts in by_entry.items()}


class ParsedClassicDbSources(BaseModel):
    """Every table this lane reads, parsed -- `fetch_classic_db_sources`'
    return shape, before it is regrouped by item id (`sources_by_item`)."""

    items: dict[int, list[ClassicDbSourceRecord]]


def _parse_creature_drops(
    sql_text: str,
    creature_names: dict[int, str],
    npc_map: dict[int, int],
    into: dict[int, list[ClassicDbSourceRecord]],
    excluded_refs: frozenset[int],
    *,
    kind: ClassicDbSourceKind,
    table: str,
) -> None:
    rows_by_entry = _rows_by_entry(list(iter_table_records(sql_text, table)))
    reference_rows = _rows_by_entry(list(iter_table_records(sql_text, "reference_loot_template")))
    for npc_id in rows_by_entry:
        name = creature_names.get(npc_id, "")
        map_id = npc_map.get(npc_id)
        for item_id, chance in _expand_loot_template(
            npc_id, rows_by_entry, reference_rows, excluded_refs
        ):
            into[item_id].append(
                ClassicDbSourceRecord(
                    kind=kind, npc_id=npc_id, name=name, map_id=map_id, chance=chance
                )
            )


def _parse_object_drops(
    sql_text: str,
    object_names: dict[int, str],
    object_map: dict[int, int],
    into: dict[int, list[ClassicDbSourceRecord]],
    excluded_refs: frozenset[int],
) -> None:
    rows_by_entry = _rows_by_entry(list(iter_table_records(sql_text, "gameobject_loot_template")))
    reference_rows = _rows_by_entry(list(iter_table_records(sql_text, "reference_loot_template")))
    for object_id in rows_by_entry:
        name = object_names.get(object_id, "")
        map_id = object_map.get(object_id)
        for item_id, chance in _expand_loot_template(
            object_id, rows_by_entry, reference_rows, excluded_refs
        ):
            into[item_id].append(
                ClassicDbSourceRecord(
                    kind="object_drop", object_id=object_id, name=name, map_id=map_id,
                    chance=chance,
                )
            )


def _parse_conditions(sql_text: str) -> dict[int, ClassicDbCondition]:
    out: dict[int, ClassicDbCondition] = {}
    for row in iter_table_records(sql_text, "conditions"):
        if int(row["type"]) != _CONDITION_REPUTATION_RANK:
            continue
        entry = int(row["condition_entry"])
        faction_id = int(row["value1"])
        rank = int(row["value2"])
        out[entry] = ClassicDbCondition(
            faction_id=faction_id, standing=decode(REP_LEVELS, rank + 1, "rep level")
        )
    return out


def _parse_vendors(
    sql_text: str,
    creature_names: dict[int, str],
    vendor_template_id: dict[int, int],
    conditions: dict[int, ClassicDbCondition],
    into: dict[int, list[ClassicDbSourceRecord]],
) -> None:
    direct = _rows_by_entry(list(iter_table_records(sql_text, "npc_vendor")))
    templated = _rows_by_entry(list(iter_table_records(sql_text, "npc_vendor_template")))
    # Every creature entry's own effective vendor list: its direct
    # npc_vendor rows, plus its VendorTemplateId's npc_vendor_template
    # rows when it has one -- cmangos' own indirection for a vendor list
    # shared verbatim by more than one creature entry (a faction's
    # generic reagent vendor, for one).
    npc_ids = set(direct) | set(vendor_template_id)
    for npc_id in npc_ids:
        rows = [*direct.get(npc_id, [])]
        template_id = vendor_template_id.get(npc_id, 0)
        if template_id:
            rows += templated.get(template_id, [])
        if not rows:
            continue
        name = creature_names.get(npc_id, "")
        for row in rows:
            condition_id = int(row["condition_id"])
            into[int(row["item"])].append(
                ClassicDbSourceRecord(
                    kind="vendor", npc_id=npc_id, name=name,
                    condition=conditions.get(condition_id) if condition_id else None,
                )
            )


def _parse_quest_rewards(sql_text: str, into: dict[int, list[ClassicDbSourceRecord]]) -> None:
    for row in iter_table_records(sql_text, "quest_template"):
        entry = int(row["entry"])
        title = unquote(row["Title"]) or ""
        quest = ClassicDbQuestInfo(
            quest_id=entry,
            min_level=int(row["MinLevel"]),
            level=int(row["QuestLevel"]),
            faction=_faction_from_required_races(int(row["RequiredRaces"])),
        )
        reward_ids = {
            int(row[f"RewChoiceItemId{n}"]) for n in range(1, 7)
        } | {int(row[f"RewItemId{n}"]) for n in range(1, 5)}
        for item_id in reward_ids:
            if not item_id:
                continue
            into[item_id].append(
                ClassicDbSourceRecord(kind="quest_reward", name=title, quest=quest)
            )


def _parse_fishing(
    sql_text: str, into: dict[int, list[ClassicDbSourceRecord]], excluded_refs: frozenset[int]
) -> None:
    # Fishing's own rows reference `reference_loot_template` the same way
    # creature/object drops do (the sample row `(1,11000,100,1,-11000,...
    # 'Fishing Loot - Zone Area: ...')` points at reference entry 11000),
    # so this needs the REAL reference table, not itself.
    rows_by_entry = _rows_by_entry(list(iter_table_records(sql_text, "fishing_loot_template")))
    reference_rows = _rows_by_entry(list(iter_table_records(sql_text, "reference_loot_template")))
    for entry in rows_by_entry:
        for item_id, chance in _expand_loot_template(
            entry, rows_by_entry, reference_rows, excluded_refs
        ):
            into[item_id].append(
                ClassicDbSourceRecord(kind="fishing", name="Fishing", chance=chance)
            )


def parse_classic_db_sources(sql_text: str) -> dict[int, list[ClassicDbSourceRecord]]:
    """The pinned dump's text, parsed into every `ClassicDbSourceRecord`
    this lane names, keyed by item id. Pure parsing -- no network, no
    filesystem; `fetch_classic_db_sources` is the network wrapper and
    `write_classic_sources` is the writer.
    """
    creature_names: dict[int, str] = {}
    vendor_template_id: dict[int, int] = {}
    for row in iter_table_records(sql_text, "creature_template"):
        entry = int(row["Entry"])
        creature_names[entry] = unquote(row["Name"]) or ""
        template_id = int(row["VendorTemplateId"])
        if template_id:
            vendor_template_id[entry] = template_id

    object_names: dict[int, str] = {
        int(row["entry"]): unquote(row["name"]) or ""
        for row in iter_table_records(sql_text, "gameobject_template")
    }

    npc_map = _spawn_map_by_entry(list(iter_table_records(sql_text, "creature")))
    object_map = _spawn_map_by_entry(list(iter_table_records(sql_text, "gameobject")))
    # `world_drop_pools` (creature/object/skinning/pickpocketing's own
    # exclusion set, world-drop-pool lane 2026-09-29) supersedes the
    # older, coarser `_excluded_reference_ids` for those four tables --
    # a row pointing at a classified pool still contributes nothing to
    # ITS OWN creature/object's attribution (same mechanism as before),
    # but the pool's own items are no longer simply dropped: they are
    # recorded once each under a synthetic `world_drop` record below.
    # Fishing keeps `_excluded_reference_ids` unchanged (`_parse_fishing`'s
    # own doc, and `_world_drop_pools`' own doc, for why the two tables
    # are not treated the same way).
    world_drop_pools = _world_drop_pools(sql_text, npc_map, object_map)
    excluded_refs = frozenset(world_drop_pools)
    fishing_excluded_refs = frozenset(_excluded_reference_ids(sql_text))

    into: dict[int, list[ClassicDbSourceRecord]] = defaultdict(list)
    _parse_creature_drops(
        sql_text, creature_names, npc_map, into, excluded_refs,
        kind="creature_drop", table="creature_loot_template",
    )
    _parse_creature_drops(
        sql_text, creature_names, npc_map, into, excluded_refs,
        kind="skinning", table="skinning_loot_template",
    )
    _parse_creature_drops(
        sql_text, creature_names, npc_map, into, excluded_refs,
        kind="pickpocketing", table="pickpocketing_loot_template",
    )
    _parse_object_drops(sql_text, object_names, object_map, into, excluded_refs)
    conditions = _parse_conditions(sql_text)
    _parse_vendors(sql_text, creature_names, vendor_template_id, conditions, into)
    _parse_quest_rewards(sql_text, into)
    _parse_fishing(sql_text, into, fishing_excluded_refs)
    for item_id, record in _world_drop_records(sql_text, world_drop_pools).items():
        into[item_id].append(record)
    return dict(into)


def fetch_classic_db_sources(
    client: httpx.Client | None = None, url: str = SOURCE_URL
) -> dict[int, list[ClassicDbSourceRecord]]:
    """Download the pinned cmangos/classic-db dump (same commit
    `classic_quest_levels` pins) and parse every table this lane needs.
    Writes nothing -- see `write_classic_sources`."""
    own = client is None
    client = client or httpx.Client(headers={"User-Agent": USER_AGENT})
    try:
        response = client.get(url, timeout=120, follow_redirects=True)
        response.raise_for_status()
        sql_text = gzip.decompress(response.content).decode("utf-8", errors="replace")
        return parse_classic_db_sources(sql_text)
    finally:
        if own:
            client.close()


def raw_path(build_dir: Path) -> Path:
    """`data/builds/<build>/raw/classicdb/sources.json` -- committed
    despite living inside the otherwise-gitignored `raw/` tree, the same
    exception `.gitignore` already carries for `raw/quests/
    quest-levels.json` and `raw/items/item-sources.json`:
    `pipeline.loot.classicdb` reads it in CI, which has no business (and,
    per this lane's own brief, no ALLOWANCE) to hit the network."""
    return build_dir / "raw" / "classicdb" / FILE_NAME


def write_classic_sources(build_dir: Path, items: dict[int, list[ClassicDbSourceRecord]]) -> Path:
    path = raw_path(build_dir)
    path.parent.mkdir(parents=True, exist_ok=True)
    document = {
        "generated_at": datetime.now(UTC).isoformat(),
        "source": {
            "classic-db": {
                "repo": "cmangos/classic-db",
                "commit": _source_commit(),
                "license": "GPL-3.0",
                "note": "creature/reference/gameobject loot templates, npc_vendor(+template), "
                "quest_template rewards and skinning/pickpocketing/fishing loot templates -- "
                "see pipeline.classic_sources' own doc.",
            },
        },
        "items": {
            str(item_id): [record.model_dump() for record in records]
            for item_id, records in sorted(items.items())
        },
    }
    path.write_text(json.dumps(document, indent=1) + "\n", encoding="utf-8")
    return path


def load_classic_sources(build_dir: Path) -> dict[int, list[ClassicDbSourceRecord]]:
    """The committed cache for one build, or `{}` (with a warning) when
    `fetch-classic-sources` has not run for it yet -- `loot`/`loot-merge`
    still run, every item this cache does not cover simply stays
    unsourced by classic-db (fork and wowhead coverage is unaffected)."""
    path = raw_path(build_dir)
    if not path.exists():
        logger.warning(
            "classic-sources: no %s; run `python -m pipeline fetch-classic-sources` first -- "
            "every item this cache does not cover gets no classic-db source until then",
            path,
        )
        return {}
    document = json.loads(path.read_text(encoding="utf-8"))
    return {
        int(item_id): [ClassicDbSourceRecord(**record) for record in records]
        for item_id, records in document["items"].items()
    }


def _source_commit() -> str:
    from pipeline.classic_quest_levels import SOURCE_COMMIT

    return SOURCE_COMMIT


def fetch_and_write_classic_sources(
    build: str, root: Path = Path("builds"), client: httpx.Client | None = None
) -> Path:
    """`python -m pipeline fetch-classic-sources`: the (one-time /
    occasional, same cadence as `fetch-classic-quest-levels`) step that
    downloads the pinned dump and writes the committed cache
    `pipeline.loot.classicdb` reads from then on."""
    items = fetch_classic_db_sources(client=client)
    build_dir = root / build
    path = write_classic_sources(build_dir, items)
    logger.info(
        "classic-sources: parsed %d item ids with at least one classic-db source; wrote %s",
        len(items),
        path,
    )
    return path
