"""`loot.json`: every place the two databases say an item comes from.

Parity contract 6.1. Droptimizer's source picker and Top Gear's "pin this
drop" both read this file, so the ids here are stable keys: a candidate's
`origin` is `drop:<source id>`.

Nothing is invented. An item the fork database gives no source and the
client gives no PvP rank is simply absent; a boss the fork does not name
is emitted with an empty name; a drop source whose kind has no home in the
contract's vocabulary -- an open-world drop with neither a named npc nor a
zone, a vendor sale with no npc id -- is dropped and counted, never guessed
into a kind.

Two kinds grew a second face here beyond the flat item lists every other
kind carries. A vendor npc (`soldBy`) becomes a `vendor` source, one per
npc, of the equippable items it sells -- a vendor selling only reagents or
consumables names no source at all. A `drop` whose npc is not a dungeon or
raid boss becomes a `zone` source, one per zone, alongside (not instead of)
the existing `world` per-npc bucket; an item several creatures in the same
zone drop appears once in that zone's list. Quests keep their single flat
`quest` bucket for compatibility, and additionally grow a `quests` map
(item id -> quest id, name and faction) on `LootFile` itself, and every
faction-restricted item this build has -- quest or not -- is named in
`LootFile.factions`.
"""

from __future__ import annotations

import logging
from collections import defaultdict
from dataclasses import dataclass

from pipeline.classic_quest_levels import item_level_proxy
from pipeline.csvio import populated
from pipeline.forkdb import FACTION_RESTRICTIONS, PROFESSIONS, REP_LEVELS, ForkDatabase, decode
from pipeline.models import LootBoss, LootFile, LootSource, QuestSource
from pipeline.normalize.classes import slugify
from pipeline.normalize.gear import SLOT_BY_INVENTORY_TYPE
from pipeline.quest_levels import QuestLevelEntry

logger = logging.getLogger(__name__)

#: The order sources are emitted in, which is the order the picker shows
#: them: instances first, then the things you can farm by zone, then the
#: things you buy or make.
KIND_ORDER = ("raid", "dungeon", "world", "zone", "vendor", "crafted", "rep", "pvp", "quest")

#: `Map.InstanceType`. 3 (battleground) and 4 (arena) are instances whose
#: loot the contract has no kind for -- a battleground's rewards are
#: reputation and rank, which are their own kinds -- so only these two
#: become drop sources.
INSTANCE_KIND = {1: "dungeon", 2: "raid"}

#: An item's own `factionRestriction` (0 included, unlike
#: `pipeline.forkdb.FACTION_RESTRICTIONS`), standing in for the side of the
#: quest that hands it out -- the fork database states no faction on a
#: quest directly. 0 means the quest is open to both, per the design.
QUEST_FACTION_BY_RESTRICTION = {0: "both", 1: "alliance", 2: "horde"}


@dataclass(frozen=True)
class LootStats:
    """What one build's loot.json covers, for the log and the test."""

    #: Distinct item ids the emitted file names.
    items: int
    #: Fork source entries whose kind has no home in contract 6.1.
    dropped_entries: int
    #: Distinct item ids the fork sources name that this build's item table
    #: does not have. Contract 10.4 leaves them out; the count is the size
    #: of the re-itemisation gap and is logged and pinned by a test.
    absent_items: int
    #: `zone` sources whose zone id `zones[]` itself does not name (emitted
    #: with `name: ""` rather than invented, same policy as an unnamed
    #: boss).
    unnamed_zones: int


def instance_types(
    map_rows: list[dict[str, str]], zone_rows: list[dict]
) -> dict[int, int]:
    """AreaTable zone id -> `Map.InstanceType`, for the zones inside one.

    Two joins, in this order, because neither covers every instance on
    build 1.60.1.69893:

    * `Map.AreaTableID` names the area a map's entrance is in, and resolves
      19 of the fork's 27 instance zones -- Molten Core, Blackwing Lair,
      Onyxia's Lair, Zul'Gurub among them.
    * `AreaTable.ContinentID`, which `zones.json` already carries as
      `map_id`, resolves the other 8 (Deadmines, Stratholme, Scholomance,
      Scarlet Monastery, Zul'Farrak, both Razorfens, Shadowfang Keep). It
      is the fallback and not the first join because it is wrong for
      Onyxia's Lair, whose AreaTable row says Kalimdor.

    The second join is also what finds Ahn'Qiraj, its Ruins and Naxxramas,
    whose drops the fork database references but whose zones its own
    `zones` list omits.
    """
    by_map = {int(row["ID"]): int(row["InstanceType"] or 0) for row in map_rows}
    by_area: dict[int, int] = {}
    for row in map_rows:
        area = int(populated(row, "AreaTableID") or 0)
        kind = int(row["InstanceType"] or 0)
        if area and kind:
            by_area.setdefault(area, kind)
    types: dict[int, int] = {}
    for zone in zone_rows:
        zone_id = int(zone["id"])
        kind = by_area.get(zone_id) or by_map.get(int(zone["map_id"]), 0)
        if kind:
            types[zone_id] = kind
    return types


def pvp_ranks(sparse_rows: list[dict[str, str]]) -> dict[int, int]:
    """Item id -> the PvP rank it requires, for the items that require one.

    The fork database has no rank data at all; the client states it, and
    266 of the fork's 273 vendor-only items are rank sets, so this is what
    makes a `pvp` source possible without inventing anything.
    """
    ranks: dict[int, int] = {}
    for row in sparse_rows:
        rank = int(populated(row, "RequiredPVPRank") or 0)
        if rank:
            ranks[int(row["ID"])] = rank
    return ranks


def _drop_sources(
    fork: ForkDatabase,
    zone_names: dict[int, str],
    types: dict[int, int],
    build_items: set[int],
    absent: set[int],
) -> tuple[list[LootSource], int, int]:
    """The raid, dungeon, world and zone sources, how many drops had no
    home, and how many zone sources named a zone `zones[]` does not.

    `absent` collects, in place, every item id this build's table does not
    have, so the caller can count the re-itemisation gap once across all
    the builders rather than three times.

    A drop whose zone is itself a dungeon or raid instance becomes a boss
    or trash entry, same as always. Everything else -- a named open-world
    mob, an unnamed one, a drop with no npc at all -- gets a `zone` entry
    when it has a zone id, *in addition to* the existing per-npc `world`
    entry when the npc is one the fork names; the two are complementary
    views of the same drop, not alternatives. Only a drop with neither a
    named npc nor any zone id has nowhere at all to go.
    """
    bosses: dict[tuple[int, int], set[int]] = defaultdict(set)
    trash: dict[int, set[int]] = defaultdict(set)
    world: dict[int, set[int]] = defaultdict(set)
    zone: dict[int, set[int]] = defaultdict(set)
    dropped = 0
    for item in fork.items:
        item_id = int(item["id"])
        for source in item.get("sources") or []:
            drop = source.get("drop")
            if drop is None:
                continue
            zone_id, npc_id = int(drop.get("zoneId", 0)), int(drop.get("npcId", 0))
            kind = INSTANCE_KIND.get(types.get(zone_id, 0))
            if kind is None and npc_id not in fork.npcs and not zone_id:
                # No dungeon/raid boss, no named open-world mob, no zone to
                # file a zone-drop under either. The contract has no kind
                # for it.
                dropped += 1
                continue
            if item_id not in build_items:
                # Contract 10.4: the build has no such item, so nothing could
                # render or sim it.
                absent.add(item_id)
                continue
            if kind is not None:
                (bosses[(zone_id, npc_id)] if npc_id else trash[zone_id]).add(item_id)
                continue
            if npc_id in fork.npcs:
                world[npc_id].add(item_id)
            if zone_id:
                zone[zone_id].add(item_id)

    out: list[LootSource] = []
    for zone_id in sorted({zone for zone, _ in bosses} | set(trash), key=lambda z: (
        INSTANCE_KIND[types[z]], slugify(zone_names[z])
    )):
        kind = INSTANCE_KIND[types[zone_id]]
        slug = f"{kind}:{slugify(zone_names[zone_id])}"
        in_zone = sorted(npc for zone, npc in bosses if zone == zone_id)
        out.append(
            LootSource(
                id=slug,
                kind=kind,
                name=zone_names[zone_id],
                zone_id=zone_id,
                bosses=[
                    LootBoss(
                        id=f"{slug}:{npc_id}",
                        name=fork.npcs.get(npc_id, ""),
                        npc_id=npc_id,
                        items=sorted(bosses[(zone_id, npc_id)]),
                    )
                    for npc_id in in_zone
                ]
                or None,
                trash=sorted(trash[zone_id]) or None,
            )
        )
    out.extend(
        LootSource(
            id=f"world:{slugify(fork.npcs[npc_id])}",
            kind="world",
            name=fork.npcs[npc_id],
            items=sorted(items),
        )
        for npc_id, items in sorted(world.items(), key=lambda pair: slugify(fork.npcs[pair[0]]))
    )
    unnamed_zones = sum(1 for zone_id in zone if zone_id not in zone_names)
    out.extend(
        LootSource(
            id=f"zone:{zone_id}",
            kind="zone",
            # A zone id the fork's drops name that `zones[]` itself omits
            # (the same gap `instance_types`' docstring measures for
            # instances) gets an empty name rather than an invented one.
            name=zone_names.get(zone_id, ""),
            zone_id=zone_id,
            items=sorted(items),
        )
        for zone_id, items in sorted(zone.items())
    )
    return out, dropped, unnamed_zones


def quest_ids_for_build(fork: ForkDatabase, build_items: set[int]) -> set[int]:
    """Every quest id `_keyed_sources` will emit a `QuestSource` for --
    the same `"quest" in source` walk, filtered to items this build
    actually has. Used by `verify-wowhead-quests` to scope its spot-check
    to ids this build's loot.json actually needs, rather than every id
    this pipeline has ever seen.
    """
    ids: set[int] = set()
    for item in fork.items:
        item_id = int(item["id"])
        if item_id not in build_items:
            continue
        for source in item.get("sources") or []:
            if "quest" in source:
                ids.add(int(source["quest"]["id"]))
    return ids


def _keyed_sources(
    fork: ForkDatabase,
    build_items: set[int],
    equippable: set[int],
    absent: set[int],
    quest_levels: dict[int, QuestLevelEntry],
) -> tuple[list[LootSource], list[int], dict[int, list[QuestSource]], int]:
    """The crafted, rep, vendor and quest sources, the per-item quest detail
    behind the flat `quest` bucket, and how many entries had no home.

    `equippable` gates the vendor kind only: a vendor who sells nothing but
    reagents or consumables sells nothing this file has a slot for, so it
    names no source at all (the item still counts toward `absent` if the
    build itself lacks it -- that check runs before the equippable one, the
    same order every other kind here uses).

    `quest_levels` is `pipeline.quest_levels.load_quest_levels`'s
    result: quest id -> the min_level/level either upstream source
    (cmangos/classic-db, or wowhead for the ids classic-db lacks)
    states for it. A quest id absent from it (neither source covers it)
    falls back to
    `item_level_proxy(item's own item_level)` for BOTH fields,
    `level_source` recording which happened.
    """
    crafted: dict[str, set[int]] = defaultdict(set)
    rep: dict[tuple[int, str], set[int]] = defaultdict(set)
    vendor_items: dict[int, set[int]] = defaultdict(set)
    vendor_names: dict[int, str] = {}
    quest: set[int] = set()
    quest_detail: dict[int, list[QuestSource]] = defaultdict(list)
    dropped = 0
    for item in fork.items:
        item_id = int(item["id"])
        for source in item.get("sources") or []:
            if not ({"crafted", "rep", "quest", "soldBy"} & set(source)):
                continue
            if item_id not in build_items:
                absent.add(item_id)
                continue
            if "crafted" in source:
                profession = decode(
                    PROFESSIONS, int(source["crafted"]["profession"]), "profession"
                )
                crafted[profession].add(item_id)
            elif "rep" in source:
                standing = decode(REP_LEVELS, int(source["rep"]["repLevel"]), "rep level")
                faction_id = int(source["rep"]["repFactionId"])
                if faction_id not in fork.factions:
                    # A faction the fork's own table does not name: there is
                    # nothing to call the source, so it is not emitted.
                    dropped += 1
                    continue
                rep[(faction_id, standing)].add(item_id)
            elif "quest" in source:
                quest.add(item_id)
                # The fork states no faction on the quest itself; the item
                # it hands out carries the quest's side as its own
                # factionRestriction, so that is what `quests` reports.
                faction = QUEST_FACTION_BY_RESTRICTION[int(item.get("factionRestriction", 0))]
                quest_id = int(source["quest"]["id"])
                known = quest_levels.get(quest_id)
                if known is not None:
                    # known.source is "classic-db" or "wowhead" --
                    # whichever pipeline.quest_levels.load_quest_levels'
                    # merged file actually resolved this quest id from.
                    min_level, level, level_source = known.min_level, known.level, known.source
                else:
                    # Neither upstream source covers this quest id: the
                    # item's own item_level stands in for both fields
                    # (this lane's brief's own fallback), and
                    # level_source says so.
                    proxy = item_level_proxy(int(item.get("ilvl") or 0))
                    min_level, level, level_source = proxy, proxy, "item_level_proxy"
                quest_detail[item_id].append(
                    QuestSource(
                        quest_id=quest_id,
                        name=source["quest"]["name"],
                        faction=faction,
                        min_level=min_level,
                        level=level,
                        level_source=level_source,
                    )
                )
            elif "soldBy" in source:
                npc_id = int(source["soldBy"].get("npcId", 0))
                if not npc_id:
                    # No npc to key a vendor source on.
                    dropped += 1
                    continue
                if item_id not in equippable:
                    continue
                vendor_items[npc_id].add(item_id)
                vendor_names.setdefault(npc_id, source["soldBy"].get("npcName", ""))
    out = [
        LootSource(
            id=f"crafted:{profession}",
            kind="crafted",
            name=profession.replace("-", " ").title(),
            profession=profession,
            items=sorted(items),
        )
        for profession, items in sorted(crafted.items())
    ]
    out += [
        LootSource(
            id=f"rep:{slugify(fork.factions[faction_id])}:{standing}",
            kind="rep",
            name=fork.factions[faction_id],
            faction_id=faction_id,
            standing=standing,
            items=sorted(items),
        )
        for (faction_id, standing), items in sorted(
            rep.items(), key=lambda pair: (slugify(fork.factions[pair[0][0]]), pair[0][1])
        )
    ]
    out += [
        LootSource(
            id=f"vendor:{npc_id}",
            kind="vendor",
            name=vendor_names[npc_id],
            npc_id=npc_id,
            items=sorted(items),
        )
        for npc_id, items in sorted(vendor_items.items())
    ]
    return out, sorted(quest), dict(quest_detail), dropped


def _pvp_sources(ranks: dict[int, int], build_items: set[int]) -> list[LootSource]:
    """One source per PvP rank. `ranks` is read off the client's own
    `ItemSparse`, so its ids are the build's by construction; the filter
    is applied anyway so one rule governs every kind."""
    by_rank: dict[int, set[int]] = defaultdict(set)
    for item_id, rank in ranks.items():
        if item_id in build_items:
            by_rank[rank].add(item_id)
    return [
        LootSource(
            id=f"pvp:rank-{rank}",
            kind="pvp",
            name=f"Rank {rank}",
            rank=rank,
            items=sorted(items),
        )
        for rank, items in sorted(by_rank.items())
    ]


def item_factions(fork: ForkDatabase, build_items: set[int]) -> dict[int, str]:
    """Item id -> "alliance" or "horde", for every faction-restricted item
    this build has -- quest items and non-quest items alike, since a
    restricted vendor, drop or crafted item has no quest to carry the fact
    on. `LootFile.quests`' own per-item `faction` covers the quest items
    a second time, from the same `factionRestriction` column; this is the
    general map the design also asks for.
    """
    factions: dict[int, str] = {}
    for item in fork.items:
        item_id = int(item["id"])
        restriction = item.get("factionRestriction")
        if not restriction or item_id not in build_items:
            continue
        factions[item_id] = decode(
            FACTION_RESTRICTIONS, int(restriction), "faction restriction"
        ).removesuffix("_only")
    return factions


class SourceIdCollision(SystemExit):
    """Two sources slugified to the same id.

    `apply_overlays` keys its sources by id (`by_id = {source.id: source
    for ...}`); two sources sharing one id would silently collapse into
    whichever the dict comprehension saw last, and a whole source's items
    would vanish from `loot.json` with nothing to say so. This is the one
    place that refuses instead, the same policy the rest of this module
    applies to a source with no kind, a boss with no name, or an item this
    build does not have.
    """


def _check_unique_ids(sources: list[LootSource]) -> None:
    seen: set[str] = set()
    for source in sources:
        if source.id in seen:
            raise SourceIdCollision(
                f"two loot sources both slugify to id {source.id!r}; "
                f"apply_overlays would silently keep only one of them. Rename "
                f"the zone, NPC or faction that produced the collision."
            )
        seen.add(source.id)


def source_item_ids(source: LootSource) -> set[int]:
    """Every item id one source names, wherever it names it."""
    return set(
        (source.items or [])
        + (source.trash or [])
        + [item for boss in (source.bosses or []) for item in boss.items]
    )


def build_loot(
    fork: ForkDatabase,
    zone_names: dict[int, str],
    types: dict[int, int],
    ranks: dict[int, int],
    build_items: set[int],
    item_inventory_types: dict[int, int],
    quest_levels: dict[int, QuestLevelEntry] | None = None,
) -> tuple[LootFile, LootStats]:
    absent: set[int] = set()
    # The vendor kind's own filter: a vendor selling only reagents or
    # consumables sells nothing this file tracks a source for.
    equippable = {
        item_id
        for item_id, inventory_type in item_inventory_types.items()
        if inventory_type in SLOT_BY_INVENTORY_TYPE
    }
    drops, dropped_drops, unnamed_zones = _drop_sources(
        fork, zone_names, types, build_items, absent
    )
    keyed, quest, quest_detail, dropped_keyed = _keyed_sources(
        fork, build_items, equippable, absent, quest_levels or {}
    )
    sources = [*drops, *keyed, *_pvp_sources(ranks, build_items)]
    if quest:
        sources.append(LootSource(id="quest", kind="quest", name="Quests", items=quest))
    # A source the filter emptied is not a source. `_drop_sources` already
    # drops a boss with no items left (its `bosses` set is simply never
    # created), so this is the last sweep: a zone whose every drop was
    # re-itemised away, like Onyxia's Lair on build 1.60.1.69893.
    sources = [source for source in sources if source_item_ids(source)]
    sources.sort(key=lambda source: (KIND_ORDER.index(source.kind), source.id))
    _check_unique_ids(sources)
    named = {item_id for source in sources for item_id in source_item_ids(source)}
    document = LootFile(
        sources=sources,
        quests={str(item_id): entries for item_id, entries in sorted(quest_detail.items())},
        factions={
            str(item_id): faction
            for item_id, faction in sorted(item_factions(fork, build_items).items())
        },
    )
    return document, LootStats(
        items=len(named),
        dropped_entries=dropped_drops + dropped_keyed,
        absent_items=len(absent),
        unnamed_zones=unnamed_zones,
    )
