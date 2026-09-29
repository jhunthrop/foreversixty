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

from pipeline.classic_sources import ClassicDbSourceRecord
from pipeline.forkdb import ForkDatabase
from pipeline.loot.constants import INSTANCE_KIND
from pipeline.loot.wowhead import merge_wowhead_sources as merge_classicdb_sources  # noqa: F401
from pipeline.models import LootBoss, LootSource, QuestSource
from pipeline.normalize.classes import slugify
from pipeline.quest_levels import QuestLevelEntry

#: `ClassicDbSourceRecord.kind` values that name a creature (a spawn
#: table's own npc entry) as the drop's origin -- skinning/pickpocketing
#: loot comes off the SAME creature a normal drop would, so it is bucketed
#: identically (this lane's own brief: "classify as world" is what
#: distinguishes them from a `creature_drop`, not a separate bucket kind).
_CREATURE_KINDS = frozenset({"creature_drop", "skinning", "pickpocketing"})


def _world_drop_id(level_min: int | None, level_max: int | None) -> str:
    """`world_drop:<level_min>-<level_max>`, or `world_drop:unknown` for a
    pool the pinned dump's own comments name no level range for at all
    (`pipeline.classic_sources._world_drop_pools`' own doc) -- never a
    raw reference-template id, which is an implementation detail of the
    dump, not a fact the site should ever show or key a URL on."""
    if level_min is None or level_max is None:
        return "world_drop:unknown"
    return f"world_drop:{level_min}-{level_max}"


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
    """
    zone_by_map = instance_zone_by_map(zone_rows, types)
    fork_instance_npcs = fork_instance_npcs or {}

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
        for record in records:
            if record.kind in _CREATURE_KINDS:
                zone_id = zone_by_map.get(record.map_id) if record.map_id else None
                npc_id = record.npc_id or 0
                if zone_id is None and npc_id:
                    zone_id = fork_instance_npcs.get(npc_id)
                if zone_id is not None and npc_id:
                    bosses[(zone_id, npc_id)].add(item_id)
                    boss_names[npc_id] = record.name
                    if record.chance is not None:
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
                    if record.chance is not None:
                        world_chances[key][item_id] = record.chance
            elif record.kind == "object_drop":
                zone_id = zone_by_map.get(record.map_id) if record.map_id else None
                object_id = record.object_id or 0
                if zone_id is not None:
                    trash[zone_id].add(item_id)
                elif object_id and record.name:
                    key = ("object", object_id)
                    world[key].add(item_id)
                    world_names[key] = record.name
                    if record.chance is not None:
                        world_chances[key][item_id] = record.chance
            elif record.kind == "fishing":
                key = ("fishing", 0)
                world[key].add(item_id)
                world_names[key] = record.name
                if record.chance is not None:
                    world_chances[key][item_id] = record.chance
            elif record.kind == "world_drop":
                level_key = (record.level_min, record.level_max)
                world_drop[level_key].add(item_id)
                if record.chance is not None:
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
                    )
                )

    out: list[LootSource] = []
    for zone_id in sorted(
        {zone_id for zone_id, _ in bosses} | set(trash),
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
            id=_world_drop_id(level_min, level_max),
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
