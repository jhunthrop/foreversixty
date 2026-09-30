"""src-classicdb lane, 2026-09-29: folding `pipeline.classic_sources`'
cmangos/classic-db dump parse into `loot.json`, for an item the engine
fork's own database names no source for (or an incomplete one) -- the
SECOND scrape/dump `pipeline.loot.sources.build_loot` folds in, applied
AFTER the fork's own sources and BEFORE `pipeline.loot.wowhead`'s scrape
(priority order: fork primary, classic-db fills gaps, wowhead lowest --
this lane's own brief). Mirrors `pipeline.loot.wowhead`'s own
bucket-then-emit/union-merge shape closely enough that this module
reuses its `merge_wowhead_sources` directly (fully generic over
`LootSource` lists despite the name -- it does not special-case
"wowhead" anywhere in its own logic, only in what it is named after).
"""

from __future__ import annotations

from collections import defaultdict

from pipeline.classic_sources import (
    CREATURE_DROP_KINDS,
    ClassicDbSourceRecord,
    direct_row_world_drop_items,
)
from pipeline.forkdb import ForkDatabase
from pipeline.loot.constants import INSTANCE_KIND, is_confirmed_boss_drop, world_drop_id
from pipeline.loot.wowhead import merge_wowhead_sources as merge_classicdb_sources  # noqa: F401
from pipeline.models import LootBoss, LootSource, QuestSource
from pipeline.normalize.classes import slugify
from pipeline.quest_levels import QuestLevelEntry

#: `ClassicDbSourceRecord.kind` values that name a creature (a spawn
#: table's own npc entry) as the drop's origin -- skinning/pickpocketing
#: loot comes off the SAME creature a normal drop would, so it is bucketed
#: identically (this lane's own brief: "classify as world" is what
#: distinguishes them from a `creature_drop`, not a separate bucket kind).
#: `pipeline.classic_sources.CREATURE_DROP_KINDS` (drop-sources-2 lane,
#: 2026-09-29) is the same set, shared so `direct_row_world_drop_items`
#: and this module never drift apart on what counts as a creature drop.
_CREATURE_KINDS = CREATURE_DROP_KINDS


def instance_zone_by_map(zone_rows: list[dict], types: dict[int, int]) -> dict[int, int]:
    """map id -> the zone id `types` marks as dungeon/raid for that map.

    classic-db's own spawn tables (`creature`/`gameobject`) give only a
    cmangos `map` id, never the finer AreaTable zone id the fork's own
    AtlasLoot-derived drops carry -- for an instance this is unambiguous
    (one zone per instance map), which is what this builds. Preferring a
    `zones.json` row whose id `types` already recognizes as an instance
    filters out a stale duplicate (Blackrock Depths carries both a
    `map_id: 0` row and its real `map_id: 230` one in `zones.json`; only
    the latter is in `types`).

    A zone whose OWN `map_id` is a bare continent id (0 Eastern Kingdoms,
    1 Kalimdor) is skipped outright, even when `types` marks it an
    instance: `zones.json`'s `map_id` is `AreaTable.ContinentID` for a
    zone `instance_types()` had to resolve through its own fallback join
    (its own doc: "wrong for Onyxia's Lair, whose AreaTable row says
    Kalimdor") -- exactly Onyxia's Lair's only `zones.json` row
    (`map_id: 1`, no better one exists for this build). Without this
    guard, EVERY open-world creature classic-db spawns anywhere on that
    whole continent resolves to this one zone id -- measured while
    regenerating loot.json for 1.60.1.70009: 1,381 "bosses" and 48,047
    items landed on `raid:onyxias-lair` alone, Teldrassil deer and boars
    included, before this guard existed.
    """
    by_map: dict[int, int] = {}
    for row in zone_rows:
        zone_id = int(row["id"])
        map_id = int(row["map_id"])
        if zone_id not in types or map_id in (0, 1):
            continue
        by_map.setdefault(map_id, zone_id)
    return by_map


def fork_instance_npc_zones(fork: ForkDatabase, types: dict[int, int]) -> dict[int, int]:
    """npc_id -> the dungeon/raid zone id the fork's OWN drops already
    place that npc in (its own `drop.npcId`/`drop.zoneId`, the same pair
    `pipeline.loot.sources._drop_sources` reads, filtered to a zone
    `types` marks as an instance).

    A dungeon/raid boss classic-db's own dump has no usable spawn map
    for at all -- a scripted/summoned encounter like Darkmaster Gandling
    (npc 1853, Scholomance) has no static `creature` spawn row, so
    `pipeline.classic_sources`' own `_spawn_map_by_entry` never states a
    `map_id` for it -- still resolves through this fallback instead of
    stranding the drop in a flat `world:<name>` bucket
    (`classicdb_additions`'s own doc has the measured regression this
    closes: 181 items on `world:darkmaster-gandling` alone, src-classicdb
    lane, before this fallback existed).
    """
    zones: dict[int, int] = {}
    for item in fork.items:
        for source in item.get("sources") or []:
            drop = source.get("drop")
            if drop is None:
                continue
            zone_id, npc_id = int(drop.get("zoneId", 0)), int(drop.get("npcId", 0))
            if not npc_id or not zone_id or INSTANCE_KIND.get(types.get(zone_id, 0)) is None:
                continue
            zones.setdefault(npc_id, zone_id)
    return zones


def classic_db_npc_names(classic_sources: dict[int, list[ClassicDbSourceRecord]]) -> dict[int, str]:
    """npc_id -> the name cmangos/classic-db's own `creature_template.
    Name` states for it, re-derived from the committed per-item records
    (`ClassicDbSourceRecord.name`, itself read off that same column while
    parsing the dump -- `pipeline.classic_sources.parse_classic_db_
    sources`'s own `creature_names`) rather than the raw SQL text, which
    this pipeline does not keep past that one parse.

    wowhead-world-drops lane, 2026-09-29's own addendum: `build_loot`'s
    fallback for a boss `LootBoss.name`'s own doc already admits the fork
    database does not name (35/67 raid bosses, 39/230 dungeon bosses on
    the pinned fork) -- classic-db drops loot for the SAME npc far more
    often than the fork's own AtlasLoot-derived table names it, so this
    covers 55 of the 58 empty names measured on build 1.60.1.70009.

    day3 data-followups-6 lane, 2026-09-30's own addendum: the remaining
    3 (npc 179703, 181366, 175245) are not really npcs at all -- the fork
    assigns a GAMEOBJECT's own entry to `drop.npcId` for a "boss" that is
    really a reward CHEST (Cache of the Firelord, Four Horsemen Chest,
    Father Flame; `pipeline.classic_sources._parse_object_drops`'s own
    doc), so `_resolve_or_drop_unnamed_bosses` looks up exactly this same
    numeric id -- it is simply never in the npc-only dict above. A second
    pass over every `object_drop` record's own `object_id`/`name` (never
    `npc_id`, which the first pass already covers) fills that gap: every
    one of the three DOES have a real classic-db name, just under
    `gameobject_template.name`, not `creature_template.Name`. Run strictly
    AFTER the npc pass (`setdefault`, so an npc id already named there
    always wins) -- an accidental numeric collision between a real npc id
    and an unrelated gameobject entry has never been observed on the
    pinned dump, but the npc fact is the more direct one either way.

    Keyed by whichever record names each id FIRST across every item's own
    list (`setdefault`) -- classic-db's own dump never disagrees with
    itself about one npc's or object's name, so which record wins is
    never ambiguous in practice.
    """
    names: dict[int, str] = {}
    for records in classic_sources.values():
        for record in records:
            if record.npc_id and record.name:
                names.setdefault(record.npc_id, record.name)
    for records in classic_sources.values():
        for record in records:
            if record.object_id and record.name:
                names.setdefault(record.object_id, record.name)
    return names


def instanced_zones(zone_ids: set[int], types: dict[int, int]) -> set[int]:
    """The subset of `zone_ids` whose Map.csv instance type names a dungeon or
    raid kind; a battleground (type 3) or arena zone gets no loot source."""
    return {zone_id for zone_id in zone_ids if INSTANCE_KIND.get(types.get(zone_id, 0)) is not None}


def _item_instance_zone_consensus(
    records: list[ClassicDbSourceRecord],
    zone_by_map: dict[int, int],
    fork_instance_npcs: dict[int, int],
) -> int | None:
    """The one instance zone id every OTHER creature-kind record for this
    same item resolves to (via its own `map_id` or `fork_instance_npcs`),
    when they all agree -- the third fallback tier `classicdb_additions`
    uses for a creature record whose own spawn data names no map at all,
    after `zone_by_map` and `fork_instance_npcs` have both already come up
    empty for it (that record's own `zone_id` is `None` at the call site).

    day3 data-followups-5 lane, 2026-09-30: Molten Core's Firelord (npc
    11668) and Lava Annihilator (npc 11665) have NO static spawn row
    anywhere in the pinned classic-db dump -- every one of their 69/70
    appearances across the whole dump states `map_id: null` -- and the
    fork's own database names neither npc at all, so `zone_by_map` and
    `fork_instance_npc_zones` are both legitimately empty for them (not a
    bug in either -- there is truly no map/zone fact recorded for these
    two npcs anywhere this pipeline reads). Both are summoned Molten Core
    trash: Fiery Core (17010) and Lava Core (17011), the two items this
    lane's brief names, list them as a source ALONGSIDE four/six other
    creatures that DO resolve to Molten Core's own zone (2717) and no
    other -- exactly the pattern this function looks for. Without this,
    Firelord's contribution reads as `world:firelord` (an open-world
    source) and Lava Annihilator's folds into a `world_drop:<range>`
    pool (`classicdb_additions`'s own `is_direct_world_drop` branch,
    since Fiery/Lava Core's 5-7 distinct creature-kind rows for one item
    clear `is_world_drop_pattern`'s `WORLD_DROP_MIN_CREATURES` floor on
    creature COUNT alone) -- both readings leave the item's real,
    already-present `raid:molten-core` source looking optional, which is
    exactly what let `apply_crafted_opens_gate`'s reagent check see an
    open source and leave every craft using either reagent ungated.

    Deliberately item-scoped, never npc-wide: Firelord alone drops 69
    different items, most of them ordinary multi-zone/continent trash
    pools (bind-on-equip greens, cloth, and the like) where its OTHER
    co-droppers span many different maps or none at all -- `zones` below
    would have zero or more than one member for those, so this returns
    `None` and they are left exactly as before. Only an item whose
    resolvable creature-kind siblings agree on ONE zone counts.

    day3 data-followups-6 lane, 2026-09-30's own addendum: an
    `object_drop` sibling counts toward consensus too, resolved the exact
    same way (`zone_by_map`, never `fork_instance_npcs`, which is an
    npc-only fallback) -- Cache of the Firelord's own two Flamewaker
    creature rows (npc 11663/11664, both `map_id: null`, the SAME "no
    static spawn row" shape Firelord/Lava Annihilator's own case above
    already documents) have no OTHER creature-kind sibling to agree with
    at all, but every Cache item also carries the chest's own `object_drop`
    record, which (once `pipeline.classic_sources._parse_object_drops`
    resolves `data1` to the chest's real spawn map, this lane's brief item
    a) DOES resolve to Molten Core on its own -- exactly the single
    resolvable sibling this fallback needs.
    """
    zones: set[int] = set()
    for record in records:
        if record.kind in _CREATURE_KINDS:
            zone_id = zone_by_map.get(record.map_id) if record.map_id else None
            if zone_id is None and record.npc_id:
                zone_id = fork_instance_npcs.get(record.npc_id)
        elif record.kind == "object_drop":
            zone_id = zone_by_map.get(record.map_id) if record.map_id else None
        else:
            continue
        if zone_id is not None:
            zones.add(zone_id)
    return next(iter(zones)) if len(zones) == 1 else None


def classicdb_additions(
    classic_sources: dict[int, list[ClassicDbSourceRecord]],
    build_items: set[int],
    equippable: set[int],
    zone_names: dict[int, str],
    types: dict[int, int],
    zone_rows: list[dict],
    quest_levels: dict[int, QuestLevelEntry],
    item_factions: dict[int, str],
    fork_instance_npcs: dict[int, int] | None = None,
) -> tuple[list[LootSource], list[int], dict[int, list[QuestSource]]]:
    """Extra sources `pipeline.classic_sources` names for an item, in the
    same `(LootSource list, quest item ids, quest detail)` shape
    `pipeline.loot.wowhead.wowhead_additions` returns, so `build_loot`
    folds both in the same way. Every `LootSource` here carries
    `source_origin="classic-db"`.

    An open-world creature/object's own map (cmangos' `map` column) is
    almost always a bare continent id (0 Eastern Kingdoms, 1 Kalimdor),
    never a specific zone -- `instance_zone_by_map` only ever resolves an
    INSTANCE map, so anything else falls to the flat `world:<name>`
    bucket, per this lane's own brief.

    `fork_instance_npcs` (`fork_instance_npc_zones`' own result, npc_id ->
    zone_id) is the fallback for a creature/skinning/pickpocketing record
    whose own spawn map does not resolve one -- either because
    classic-db's dump states no `map_id` at all (a scripted/summoned
    dungeon or raid boss with no static spawn row: Darkmaster Gandling,
    Scholomance's npc 1853, is the measured case) or because `zone_by_map`
    itself has no entry for that map id. Every npc the fork already places
    in a dungeon or raid must land in that SAME instance here too, never
    `world` -- src-classicdb-fixes lane, 2026-09-29 (181 items measured on
    `world:darkmaster-gandling` alone before this fallback existed).

    `_item_instance_zone_consensus` (day3 data-followups-5 lane,
    2026-09-30) is a further, item-scoped fallback for a creature BOTH of
    those still miss -- one with no static spawn row AND no fork mention
    at all (Molten Core's Firelord and Lava Annihilator, its own doc's
    measured case). See its own doc for the full reasoning.

    `item_factions` (`pipeline.loot.sources.item_factions`' own result)
    is what a quest reward's `QuestSource.faction` is built from --
    `QuestSource`'s own doc: "the item it awards carries whichever
    factionRestriction the quest's side amounts to in practice", the
    SAME rule the fork's own `_keyed_sources` quest handling uses, so a
    classic-db-sourced quest reward's faction always agrees with the
    item's own client-stated restriction. `record.quest.faction`
    (classic-db's own `RequiredRaces` reading) is deliberately NOT used
    here even though it is real data -- it can legitimately disagree
    with the item's own restriction (a handful of quests measured while
    regenerating loot.json for 1.60.1.70009), and this file's own
    invariant (`test_the_quests_map_faction_matches_the_items_own_
    restriction`) is that every origin agrees on the one fact the ITEM
    states.

    drop-sources-2 lane, 2026-09-29: an item in `direct_world_drop_items`
    (`pipeline.classic_sources.direct_row_world_drop_items`' own result,
    intersected with `build_items`) reads as a generic world-drop pattern
    from its OWN direct `creature_loot_template` rows alone -- items
    7909/7910/4306, ~900-1,500 per-creature rows each, measured on build
    1.60.1.70009's audit. Every one of that item's creature-kind records
    that `is_confirmed_boss_drop` does not exempt folds into ONE
    synthetic `world_drop` source for the item instead of its own
    boss/world bucket entry, its level range the min/max of every folded
    row's own creature level (`ClassicDbSourceRecord.level_min`/
    `level_max`, from `creature_template.MinLevel`/`MaxLevel`) and its
    chance the highest of theirs -- the same "never invented, never
    averaged" merge `pipeline.classic_sources._world_drop_records`
    already uses for a reference-pool world drop.

    An open-world creature (`is_confirmed_boss_drop` false because the
    row resolves to no dungeon/raid zone at all) always folds here,
    whatever its own chance -- the world-pool rule applies to open-world
    creatures only. Tier armour in cmangos routinely drops from several
    bosses, or many trash creatures, of ONE instance, sometimes at a
    chance under 1% or none stated -- gating the dungeon/raid exemption
    on a chance floor (the retired `WORLD_DROP_BOSS_MIN_CHANCE_PERCENT`)
    folded exactly those rows into `world_drop` instead, the regression
    a prior lane's report measured (raid distinct items 767 -> 350 on
    build 1.60.1.70009).

    pooled-boss-greens lane, 2026-09-29's own addendum: a dungeon/raid
    row ALSO folds -- without ever needing `is_direct_world_drop` at all
    -- when the item already carries a separate `world_drop` record from
    a differently-shaped pool AND this row's own chance is unknown or
    below `WORLD_DROP_MAX_CHANCE_PERCENT` (`is_confirmed_boss_drop`'s own
    doc: cmangos' "(Boss Loot)" reference groups, used by exactly one
    boss, that `pipeline.classic_sources._world_drop_pools`' fan-out/
    multi-map signals cannot see). That row needs no new pool -- the
    `elif record.kind == "world_drop":` branch below, over this SAME
    item's own separate `world_drop` record, already lists it.
    """
    zone_by_map = instance_zone_by_map(zone_rows, types)
    fork_instance_npcs = fork_instance_npcs or {}
    direct_world_drop_items = direct_row_world_drop_items(classic_sources) & build_items

    bosses: dict[tuple[int, int], set[int]] = defaultdict(set)
    boss_names: dict[int, str] = {}
    boss_chances: dict[tuple[int, int], dict[int, float]] = defaultdict(dict)
    trash: dict[int, set[int]] = defaultdict(set)
    # World bucket keys are (namespace, id) so a creature and a
    # gameobject that happen to share a raw id never collide.
    world: dict[tuple[str, int], set[int]] = defaultdict(set)
    world_names: dict[tuple[str, int], str] = {}
    world_chances: dict[tuple[str, int], dict[int, float]] = defaultdict(dict)
    vendor_items: dict[int, set[int]] = defaultdict(set)
    vendor_names: dict[int, str] = {}
    vendor_condition: dict[int, tuple[int, str]] = {}
    quest: set[int] = set()
    quest_detail: dict[int, list[QuestSource]] = defaultdict(list)
    # World-drop bucket keys are the pool's own (level_min, level_max) --
    # `pipeline.classic_sources._world_drop_records`' own doc: a green-
    # quality world drop is typically split across several narrow-banded
    # reference ids on the pinned dump that all describe the one real
    # item, already merged to one range per item there; grouping here by
    # that same range (rather than by item id 1:1, which would need no
    # grouping at all) is what turns every item sharing a range into ONE
    # `world_drop` source instead of one per item, the same "single
    # synthetic source" shape a raid/dungeon zone's own bucket has.
    world_drop: dict[tuple[int | None, int | None], set[int]] = defaultdict(set)
    world_drop_chances: dict[tuple[int | None, int | None], dict[int, float]] = defaultdict(dict)

    for item_id, records in classic_sources.items():
        if item_id not in build_items:
            continue
        # `is_direct_world_drop`'s own item: every creature-kind record
        # `is_confirmed_boss_drop` does not exempt folds into this one
        # `(level_min, level_max)` pool instead of its own boss/world
        # entry below -- `direct_pool_level`'s own `is not None` (not its
        # contents, which may honestly be `(None, None)`) is the "did
        # anything fold at all" flag; an item every one of whose own rows
        # IS exempted gets no synthetic source at all (`classicdb_
        # additions`'s own doc, above).
        is_direct_world_drop = item_id in direct_world_drop_items
        # pooled-boss-greens lane, 2026-09-29: whether THIS item already
        # carries its own `world_drop` `ClassicDbSourceRecord` (a
        # differently-shaped, marked/multi-map/fan-out pool elsewhere in
        # the dump names this same item) -- `is_confirmed_boss_drop`'s
        # own doc for why a dungeon/raid row still folds when this is
        # true and the row's own chance is unknown or below
        # `WORLD_DROP_MAX_CHANCE_PERCENT`, rather than "whatever its
        # chance" keeping it on the boss unconditionally.
        has_world_drop_record = any(record.kind == "world_drop" for record in records)
        # `_item_instance_zone_consensus`'s own doc: the third fallback tier
        # for a creature record whose own map is None AND `fork_instance_npcs`
        # names nothing for its npc -- computed once per item, not per
        # record, since it reads every OTHER creature record for this same
        # item_id.
        item_zone_consensus = _item_instance_zone_consensus(
            records, zone_by_map, fork_instance_npcs
        )
        direct_pool_level: tuple[int | None, int | None] | None = None
        direct_pool_chance: float | None = None
        for record in records:
            if record.kind in _CREATURE_KINDS:
                zone_id = zone_by_map.get(record.map_id) if record.map_id else None
                npc_id = record.npc_id or 0
                if zone_id is None and npc_id:
                    zone_id = fork_instance_npcs.get(npc_id)
                # The consensus fallback is narrower than the two above: only
                # for a record whose own `map_id` is a true `None` (classic-db
                # records no spawn map for this npc at all -- Firelord, Lava
                # Annihilator), never one recorded as a bare continent id (0
                # Eastern Kingdoms, 1 Kalimdor -- Trash One..Six's own shape,
                # `_item_instance_zone_consensus`'s own doc), which is real,
                # if coarse, open-world data that a same-item instance sibling
                # must never override.
                if zone_id is None and npc_id and record.map_id is None:
                    zone_id = item_zone_consensus
                confirmed = is_confirmed_boss_drop(
                    zone_id is not None, has_world_drop_record, record.chance
                )
                if is_direct_world_drop and not confirmed:
                    lo, hi = direct_pool_level or (None, None)
                    if record.level_min is not None:
                        lo = record.level_min if lo is None else min(lo, record.level_min)
                    if record.level_max is not None:
                        hi = record.level_max if hi is None else max(hi, record.level_max)
                    direct_pool_level = (lo, hi)
                    if record.chance:
                        direct_pool_chance = (
                            record.chance
                            if direct_pool_chance is None
                            else max(direct_pool_chance, record.chance)
                        )
                    continue
                if zone_id is not None and npc_id and not confirmed:
                    # A dungeon/raid row `is_confirmed_boss_drop` no
                    # longer confirms: this item's own separate
                    # `world_drop` source already lists it (the `elif
                    # record.kind == "world_drop":` branch below, over a
                    # DIFFERENT record for this same item_id) -- no new
                    # pool needed, just no attribution to this boss.
                    continue
                if zone_id is not None and npc_id:
                    bosses[(zone_id, npc_id)].add(item_id)
                    boss_names[npc_id] = record.name
                    # wowhead-world-drops lane, 2026-09-29's own addendum:
                    # cmangos/classic-db's own `ChanceOrQuestChance` uses 0
                    # as ITS OWN "no chance recorded" sentinel (not a real
                    # 0% -- a drop with a truly 0% chance would simply not
                    # be a row in the dump at all), so `record.chance` (a
                    # truthy check, not `is not None`) excludes it the same
                    # way an absent chance already is -- an omitted key
                    # (unknown) is honest; a published `0` reads as "this
                    # never drops", which is not a fact this dump states.
                    if record.chance:
                        boss_chances[(zone_id, npc_id)][item_id] = record.chance
                elif npc_id and record.name:
                    # An unnamed npc (creature_template names none) has
                    # nowhere honest to go: `world:<slug>` is keyed by
                    # NAME, so two unrelated unnamed creatures would
                    # otherwise both slugify to the same blank
                    # "world:" id and get merged into one fake shared
                    # bucket (measured while regenerating loot.json for
                    # 1.60.1.70009 -- fixed here, matching the fork's
                    # own `if npc_id in fork.npcs` gate on its identical
                    # world bucket, which already excludes an unnamed
                    # npc the same way).
                    key = ("npc", npc_id)
                    world[key].add(item_id)
                    world_names[key] = record.name
                    # See the creature_drop branch's own comment above: 0
                    # is classic-db's "unknown" sentinel, never a real 0%.
                    if record.chance:
                        world_chances[key][item_id] = record.chance
            elif record.kind == "object_drop":
                zone_id = zone_by_map.get(record.map_id) if record.map_id else None
                object_id = record.object_id or 0
                if zone_id is not None and object_id and record.is_reward_chest and record.name:
                    # A named, `data1`-curated reward chest (Cache of the
                    # Firelord, Father Flame, Four Horsemen Chest --
                    # `ClassicDbSourceRecord.is_reward_chest`'s own doc)
                    # spawning in an instance IS the boss: the fork's own
                    # `_drop_sources` already keys the identical chest as a
                    # "boss" by this SAME object entry in its own `drop.
                    # npcId` slot (day3 data-followups-6 lane, 2026-09-30's
                    # own measurement), so this merges into that one entry
                    # rather than publishing an anonymous trash line
                    # alongside it.
                    bosses[(zone_id, object_id)].add(item_id)
                    boss_names[object_id] = record.name
                    if record.chance:
                        boss_chances[(zone_id, object_id)][item_id] = record.chance
                elif zone_id is not None:
                    # An ORDINARY instance object -- an ore vein, or a
                    # generic chest whose own entry IS its own loot key
                    # because `gameobject_template.data1` is unset --
                    # stays anonymous trash exactly as before, whatever its
                    # own name: `is_reward_chest` is what tells the two
                    # apart, not `record.name`'s own truthiness (130 type-3
                    # rows on the pinned dump are named AND `data1`-less,
                    # `ClassicDbSourceRecord.is_reward_chest`'s own doc).
                    trash[zone_id].add(item_id)
                elif object_id and record.name:
                    key = ("object", object_id)
                    world[key].add(item_id)
                    world_names[key] = record.name
                    if record.chance:
                        world_chances[key][item_id] = record.chance
            elif record.kind == "fishing":
                key = ("fishing", 0)
                world[key].add(item_id)
                world_names[key] = record.name
                if record.chance:
                    world_chances[key][item_id] = record.chance
            elif record.kind == "world_drop":
                level_key = (record.level_min, record.level_max)
                world_drop[level_key].add(item_id)
                if record.chance:
                    world_drop_chances[level_key][item_id] = record.chance
            elif record.kind == "vendor":
                # `vendor:<npc_id>` is keyed by id, not name, so an
                # unnamed vendor is harmless to keep (no collision risk
                # the way the name-keyed `world:` bucket has) -- still
                # requires a name to show something better than a blank
                # line on the site, same policy as everywhere else here.
                if not record.npc_id or not record.name or item_id not in equippable:
                    continue
                vendor_items[record.npc_id].add(item_id)
                vendor_names.setdefault(record.npc_id, record.name)
                if record.condition is not None:
                    vendor_condition.setdefault(
                        record.npc_id, (record.condition.faction_id, record.condition.standing)
                    )
            elif record.kind == "quest_reward":
                assert record.quest is not None
                quest.add(item_id)
                known = quest_levels.get(record.quest.quest_id)
                if known is not None:
                    min_level, level = known.min_level, known.level
                else:
                    min_level, level = record.quest.min_level, record.quest.level
                quest_detail[item_id].append(
                    QuestSource(
                        quest_id=record.quest.quest_id,
                        name=record.name,
                        faction=item_factions.get(item_id, "both"),
                        min_level=min_level,
                        level=level,
                        level_source="classic-db",
                        required_rep_faction=record.quest.required_rep_faction,
                        required_rep_standing=record.quest.required_rep_standing,
                    )
                )
        if direct_pool_level is not None:
            world_drop[direct_pool_level].add(item_id)
            if direct_pool_chance:
                world_drop_chances[direct_pool_level][item_id] = direct_pool_chance

    out: list[LootSource] = []
    # Only zones whose Map.csv instance type is a dungeon or raid get a source
    # here; a fresh export also lists battlegrounds (type 3) and arenas, which
    # data.yml's 2026-09-30 rebuild hit as KeyError 3 through the trash set.
    instanced = instanced_zones({zone_id for zone_id, _ in bosses} | set(trash), types)
    for zone_id in sorted(
        instanced,
        key=lambda z: (INSTANCE_KIND[types[z]], slugify(zone_names.get(z, str(z)))),
    ):
        kind = INSTANCE_KIND[types[zone_id]]
        slug = f"{kind}:{slugify(zone_names.get(zone_id) or str(zone_id))}"
        in_zone = sorted(npc for zone, npc in bosses if zone == zone_id)
        out.append(
            LootSource(
                id=slug,
                kind=kind,
                name=zone_names.get(zone_id, ""),
                zone_id=zone_id,
                bosses=[
                    LootBoss(
                        id=f"{slug}:{npc_id}",
                        name=boss_names.get(npc_id, ""),
                        npc_id=npc_id,
                        items=sorted(bosses[(zone_id, npc_id)]),
                        item_chances={
                            str(item_id): chance
                            for item_id, chance in sorted(boss_chances[(zone_id, npc_id)].items())
                        }
                        or None,
                    )
                    for npc_id in in_zone
                ]
                or None,
                trash=sorted(trash[zone_id]) or None,
                source_origin="classic-db",
            )
        )
    out.extend(
        LootSource(
            id=f"world:{slugify(world_names[key])}",
            kind="world",
            name=world_names[key],
            items=sorted(items),
            item_chances={
                str(item_id): chance for item_id, chance in sorted(world_chances[key].items())
            }
            or None,
            source_origin="classic-db",
        )
        for key, items in sorted(world.items(), key=lambda pair: slugify(world_names[pair[0]]))
    )
    out.extend(
        LootSource(
            id=world_drop_id(level_min, level_max),
            kind="world_drop",
            name="World drop",
            items=sorted(items),
            item_chances={
                str(item_id): chance
                for item_id, chance in sorted(world_drop_chances[(level_min, level_max)].items())
            }
            or None,
            level_min=level_min,
            level_max=level_max,
            source_origin="classic-db",
        )
        # Sorted by level_min (unknown-level buckets, `None`, last -- a
        # bucket this pinned dump's own comments state no level for is
        # rarer and less specific than one that does) so the picker's own
        # per-kind list reads low-to-high, matching every other level-
        # banded kind's own convention.
        for (level_min, level_max), items in sorted(
            world_drop.items(),
            key=lambda pair: (
                pair[0][0] is None, pair[0][0] or 0,
                pair[0][1] is None, pair[0][1] or 0,
            ),
        )
    )
    out.extend(
        LootSource(
            id=f"vendor:{npc_id}",
            kind="vendor",
            name=vendor_names[npc_id],
            npc_id=npc_id,
            items=sorted(items),
            faction_id=vendor_condition.get(npc_id, (None, None))[0],
            standing=vendor_condition.get(npc_id, (None, None))[1],
            source_origin="classic-db",
        )
        for npc_id, items in sorted(vendor_items.items())
    )
    return out, sorted(quest), dict(quest_detail)
