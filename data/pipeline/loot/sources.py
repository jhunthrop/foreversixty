"""`loot.json`: every place the two databases say an item comes from.

Parity contract 6.1. Droptimizer's source picker and Top Gear's "pin this
drop" both read this file, so the ids here are stable keys: a candidate's
`origin` is `drop:<source id>`.

Nothing is invented. An item the fork database gives no source and the
client gives no PvP rank is simply absent; a boss the fork does not name
is emitted with an empty name; a drop source whose kind has no home in the
contract's vocabulary -- a vendor with no rank, an unnamed open-world mob --
is dropped and counted, never guessed into a kind.
"""

from __future__ import annotations

import logging
from collections import defaultdict
from dataclasses import dataclass

from pipeline.csvio import populated
from pipeline.forkdb import PROFESSIONS, REP_LEVELS, ForkDatabase, decode
from pipeline.models import LootBoss, LootFile, LootSource
from pipeline.normalize.classes import slugify

logger = logging.getLogger(__name__)

#: The order sources are emitted in, which is the order the picker shows
#: them: instances first, then the things you buy or make.
KIND_ORDER = ("raid", "dungeon", "world", "crafted", "rep", "pvp", "quest")

#: `Map.InstanceType`. 3 (battleground) and 4 (arena) are instances whose
#: loot the contract has no kind for -- a battleground's rewards are
#: reputation and rank, which are their own kinds -- so only these two
#: become drop sources.
INSTANCE_KIND = {1: "dungeon", 2: "raid"}


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
) -> tuple[list[LootSource], int]:
    """The raid, dungeon and world sources, and how many drops had no home.

    `absent` collects, in place, every item id this build's table does not
    have, so the caller can count the re-itemisation gap once across all
    the builders rather than three times.
    """
    bosses: dict[tuple[int, int], set[int]] = defaultdict(set)
    trash: dict[int, set[int]] = defaultdict(set)
    world: dict[int, set[int]] = defaultdict(set)
    dropped = 0
    for item in fork.items:
        item_id = int(item["id"])
        for source in item.get("sources") or []:
            drop = source.get("drop")
            if drop is None:
                continue
            zone_id, npc_id = int(drop.get("zoneId", 0)), int(drop.get("npcId", 0))
            kind = INSTANCE_KIND.get(types.get(zone_id, 0))
            if kind is None and npc_id not in fork.npcs:
                # An open-world mob the fork does not name, or a drop with no
                # zone at all. The contract has no kind for either.
                dropped += 1
                continue
            if item_id not in build_items:
                # Contract 10.4: the build has no such item, so nothing could
                # render or sim it.
                absent.add(item_id)
                continue
            if kind is not None:
                (bosses[(zone_id, npc_id)] if npc_id else trash[zone_id]).add(item_id)
            else:
                world[npc_id].add(item_id)

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
    return out, dropped


def _keyed_sources(
    fork: ForkDatabase, build_items: set[int], absent: set[int]
) -> tuple[list[LootSource], list[int], int]:
    """The crafted, rep and quest sources, plus how many entries had no home."""
    crafted: dict[str, set[int]] = defaultdict(set)
    rep: dict[tuple[int, str], set[int]] = defaultdict(set)
    quest: set[int] = set()
    dropped = 0
    for item in fork.items:
        item_id = int(item["id"])
        for source in item.get("sources") or []:
            if not ({"crafted", "rep", "quest"} & set(source)):
                if "soldBy" in source:
                    # A vendor. Rank vendors are covered by the pvp kind,
                    # read off the client; anything else has no kind in 6.1.
                    dropped += 1
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
    return out, sorted(quest), dropped


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
) -> tuple[LootFile, LootStats]:
    absent: set[int] = set()
    drops, dropped_drops = _drop_sources(fork, zone_names, types, build_items, absent)
    keyed, quest, dropped_keyed = _keyed_sources(fork, build_items, absent)
    sources = [*drops, *keyed, *_pvp_sources(ranks, build_items)]
    if quest:
        sources.append(LootSource(id="quest", kind="quest", name="Quests", items=quest))
    # A source the filter emptied is not a source. `_drop_sources` already
    # drops a boss with no items left (its `bosses` set is simply never
    # created), so this is the last sweep: a zone whose every drop was
    # re-itemised away, like Onyxia's Lair on build 1.60.1.69893.
    sources = [source for source in sources if source_item_ids(source)]
    sources.sort(key=lambda source: (KIND_ORDER.index(source.kind), source.id))
    named = {item_id for source in sources for item_id in source_item_ids(source)}
    return LootFile(sources=sources), LootStats(
        items=len(named),
        dropped_entries=dropped_drops + dropped_keyed,
        absent_items=len(absent),
    )
