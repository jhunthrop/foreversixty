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
from pipeline.classic_sources import (
    ClassicDbSourceRecord,
    quest_classes_from_classic_sources,
    quest_factions_from_classic_sources,
    quest_profession_from_classic_sources,
    quest_turn_in_items_from_classic_sources,
)
from pipeline.classicdb_crafted import ClassicDbCraftedRecipe
from pipeline.csvio import populated
from pipeline.forkdb import FACTION_RESTRICTIONS, PROFESSIONS, REP_LEVELS, ForkDatabase, decode
from pipeline.item_sources import ItemSourceEntry
from pipeline.loot.classicdb import (
    classic_db_npc_names,
    classicdb_additions,
    fork_instance_npc_zones,
    merge_classicdb_sources,
)
from pipeline.loot.constants import INSTANCE_KIND
from pipeline.loot.wowhead import merge_wowhead_sources, wowhead_additions
from pipeline.models import LootBoss, LootFile, LootSource, QuestSource
from pipeline.normalize.classes import slugify
from pipeline.normalize.gear import SLOT_BY_INVENTORY_TYPE
from pipeline.quest_levels import QuestLevelEntry

logger = logging.getLogger(__name__)

#: The order sources are emitted in, which is the order the picker shows
#: them: instances first, then the things you can farm by zone, then the
#: things you buy or make. `world_drop` (src-classicdb world-drop-pool
#: lane, 2026-09-29) sits right after `world`: both are "go kill
#: something out in the world", and `world_drop`'s own generic,
#: auction-housable pool is a step LESS specific than a named `world`
#: mob's own drop.
KIND_ORDER = (
    "raid", "dungeon", "world", "world_drop", "zone", "vendor", "crafted", "rep", "pvp", "quest",
)  # fmt: skip

#: An item's own `factionRestriction` (0 included, unlike
#: `pipeline.forkdb.FACTION_RESTRICTIONS`), standing in for the side of the
#: quest that hands it out -- the fork database states no faction on a
#: quest directly. 0 means the quest is open to both, per the design.
QUEST_FACTION_BY_RESTRICTION = {0: "both", 1: "alliance", 2: "horde"}


def resolve_quest_faction(
    quest_id: int,
    item_id: int,
    quest_factions: dict[int, str],
    quest_levels: dict[int, QuestLevelEntry],
    item_sources: dict[int, ItemSourceEntry],
    item_faction_restrictions: dict[int, str],
) -> tuple[str, str]:
    """`(faction, faction_source)` for one (quest id, reward item id)
    pair, in primary-source order (data-followups-3 lane, 2026-09-30,
    item 2 -- the "Friend of the Library" defect, quest 78150: neither
    reward item, Erudite's Amulet nor Scholarly Pendant, carries a
    `factionRestriction`, so the OLD item-derived guess this function
    replaces was always "both" regardless of which faction can actually
    receive each one):

    1. `quest_factions` (`quest_factions_from_classic_sources`'s
       result): the quest's own classic-db `RequiredRaces` -- verified
       against a primary source, wins outright when present, for every
       reward item the quest hands out alike (a fact about the QUEST).
    2. `item_faction_restrictions` (`item_factions`'s own result): the
       reward item's own client-stated `factionRestriction`, when it is
       genuinely non-zero -- `QuestSource`'s own long-standing contract
       (`test_loot_sources_wowhead.
       test_a_wowhead_quest_reward_faction_follows_the_items_own_
       restriction_not_the_scrapes`: item 100's real `alliance_only`
       wins over a deliberately-mismatched wowhead scrape). Left
       UNTOUCHED by this lane -- only a REAL restriction short-circuits
       here; an unrestricted item (absent from this dict, its
       `factionRestriction` 0) falls through to step 3, which is what
       Friend of the Library's own two items do (neither carries one).
    3. `item_sources[item_id].quest_rewards` (`pipeline.
       wowhead_item_sources.QuestRewardSource.faction`, already
       committed in `raw/items/item-sources.json` for every item the
       fork database itself names no source for -- no new fetch needed):
       the matching `quest_id` row's own wowhead `side`, when it states
       "alliance" or "horde" -- a real, per-ITEM signal (measured while
       building this lane's report: 106 of 145 item/quest pairs among
       this build's `faction_source: "item"` quests resolve this way,
       including the "Guardian Talisman"-shaped case wowhead.py's own
       `wowhead_additions` doc already names -- the SAME item rewarded
       by two faction-mirrored quest ids, each correctly one-sided here).
       Skipped when wowhead states "both": indistinguishable, in
       practice, from "wowhead's scrape has not resolved this one either"
       -- Friend of the Library's own two items are exactly this case,
       both wowhead-labelled "both" despite the observed one-reward-per-
       faction split in play, so this tier does not report it as fact.
    4. `quest_levels[quest_id].faction` (`pipeline.quest_levels.
       QuestLevelEntry.faction`): wowhead's own `side` off the quest
       PAGE'S `g_quests[<id>]` payload (a per-QUEST, not per-item,
       fallback for an item the fork itself sources, so step 3's
       `item_sources` cache was never fetched for it) -- the client's
       own DB2 quest tables (`QuestV2`/`QuestInfo`'s allowable-races
       mask, whichever names it) would be the primary source ahead of
       even classic-db, but this pipeline does not fetch them at all (no
       build under `data/builds/<build>/raw/quests/` carries one) -- see
       this lane's own report for which table a future lane would add.
    5. `"unknown"`, published honestly rather than invented from a
       DEFAULTED (not real) item restriction -- `QuestSource`'s own doc,
       item 270018 Hammerbone/quest 914 for why a quest's own faction and
       an unrestricted item's can legitimately disagree.
       `sim/cmd/leveling-bis/data.go`'s own `questFactionSide` map has no
       `"unknown"` key, so a Go map lookup's own zero value reads this
       exactly like `"both"` until a real source covers it -- not a new
       runtime behaviour, just an honest label in place of a wrong guess.

    Every case is tagged `faction_source`: `"classic-db"`, `"item"`
    (either a real restriction, step 2, or unverified, step 5 -- kept as
    one tag for continuity with the pre-existing contract, which never
    distinguished the two), or `"wowhead"` (steps 3/4).
    """
    if quest_id in quest_factions:
        return quest_factions[quest_id], "classic-db"
    if item_id in item_faction_restrictions:
        return item_faction_restrictions[item_id], "item"
    item_entry = item_sources.get(item_id)
    if item_entry is not None:
        for reward in item_entry.quest_rewards:
            if reward.quest_id == quest_id and reward.faction != "both":
                return reward.faction, "wowhead"
    quest_entry = quest_levels.get(quest_id)
    if quest_entry is not None and quest_entry.faction is not None:
        return quest_entry.faction, "wowhead"
    return "unknown", "item"

#: `pipeline.forkdb.FACTION_RESTRICTIONS`' 1/2 (alliance_only/horde_only),
#: minus the `_only` suffix, for comparing an item's own restriction
#: against the side a rep faction id belongs to.
_RESTRICTION_SIDE = {1: "alliance", 2: "horde"}

#: Which side a battleground reputation faction id is on. Sourced from the
#: fork's own `factions` table names, not invented: Silverwing Sentinels
#: and Stormpike Guard are Alliance-run; Warsong Outriders and Frostwolf
#: Clan are Horde-run. Used only to detect and correct
#: `_ATLASLOOT_WSG_REP_FACTION_SWAP` below, never to reject an id outright.
_BG_REP_FACTION_SIDE = {889: "alliance", 890: "horde", 729: "horde", 730: "alliance"}

#: Warsong Gulch's two reputation factions come out of the fork's
#: `assets/database/db.json` swapped: every item mined under 889
#: (Silverwing Sentinels, Alliance) actually carries `factionRestriction:
#: 2` (horde_only), and every item under 890 (Warsong Outriders, Horde)
#: carries `factionRestriction: 1` (alliance_only) -- confirmed on the
#: 2026-09-28 hunter/paladin audits for Scout's/Sentinel's Medallion
#: (20442/20444/19541) and Lorekeeper's Staff (19573), and for all 124 WSG
#: rep-sourced items on this build by cross-checking both the item's own
#: `factionRestriction` and the same source's `playerFaction` field, which
#: agree with each other and disagree with `repFactionId` every time.
#: Arathi Basin (509/510) and Alterac Valley (729/730) are not affected --
#: only WSG's ids are reversed.
#:
#: The defect is upstream of this pipeline, in the wowsims-forever engine
#: fork's `tools/database/atlasloot.go` (~line 406), whose hardcoded
#: `RepFactionId` map has the ALLIANCE and HORDE keys pointing at each
#: other's WSG faction id (its own comments name the right faction for the
#: wrong key). That needs a one-line swap and a re-pin on the engine side;
#: this pipeline cannot regenerate `db.json` and so corrects the id here
#: instead, so `loot.json` itself -- not just its Go consumer's separate
#: safety net -- is right.
#:
#: `_corrected_rep_faction_id` only ever applies this when doing so turns a
#: mismatch with the item's own `factionRestriction` into a match, so once
#: the engine fix lands and `db.json` stops swapping the ids, this becomes
#: a no-op rather than re-breaking correct data.
_ATLASLOOT_WSG_REP_FACTION_SWAP = {889: 890, 890: 889}


def _corrected_rep_faction_id(faction_id: int, faction_restriction: int) -> int:
    """`faction_id`, or its swap partner when the swap partner is the one
    that actually matches the item's own `factionRestriction`.

    See `_ATLASLOOT_WSG_REP_FACTION_SWAP`'s docstring for the defect this
    corrects and why it self-disables once the upstream data is fixed.
    """
    swapped = _ATLASLOOT_WSG_REP_FACTION_SWAP.get(faction_id)
    if swapped is None or not faction_restriction:
        return faction_id
    expected = _RESTRICTION_SIDE.get(faction_restriction)
    if expected is None:
        return faction_id
    if _BG_REP_FACTION_SIDE.get(faction_id) != expected and (
        _BG_REP_FACTION_SIDE.get(swapped) == expected
    ):
        return swapped
    return faction_id


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
    #: Distinct item ids sourced ONLY because `item_sources` (a wowhead
    #: scrape, night-item-sources lane) named one for an item the fork
    #: database itself named none for at all -- disjoint from `items`'
    #: fork-derived count above, for the report this lane's brief asks
    #: for.
    wowhead_items: int = 0
    #: Distinct item ids sourced ONLY because `classic_sources` (the
    #: cmangos/classic-db dump, src-classicdb lane) named one for an item
    #: the fork database itself named none for -- disjoint from `items`
    #: (fork) and `wowhead_items` (both counted before this origin is
    #: applied, so a quest item the fork already covers never double
    #: counts here).
    classicdb_items: int = 0
    #: wowhead-world-drops lane, 2026-09-29's own addendum: bosses whose
    #: name resolved to neither the fork database's own npcs table nor
    #: classic-db's `creature_template.Name` (`classic_db_npc_names`) --
    #: dropped from their source rather than published with an empty
    #: name (`_resolve_or_drop_unnamed_bosses`'s own doc).
    dropped_unnamed_bosses: int = 0


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
                faction_id = _corrected_rep_faction_id(
                    int(source["rep"]["repFactionId"]), int(item.get("factionRestriction") or 0)
                )
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


#: Bound on `apply_quest_opens_gate`'s own fixed-point loop -- a quest
#: gated only through ANOTHER quest's own gate (item-level recursion,
#: not the `PrevQuestId` chain `pipeline.classic_sources` already
#: resolves) needs at most a couple of passes in practice; this is
#: headroom, the same role every other bounded-recursion constant in
#: this pipeline plays.
_MAX_QUEST_GATE_PASSES = 5

#: `quest_template.RequiredMinRepFaction` (or the same item's own
#: `rep`-kind `LootSource`, `_rep_gate_by_item`'s own doc) -> the content
#: phase that faction's own reputation track requires -- mirrored from
#: `sim/cmd/leveling-bis/data.go`'s own `repFactionRaidPhaseOpens` (name
#: it once in Python; not imported from there, and not placed in
#: `pipeline.loot.overlay` beside `RAID_DEFAULT_OPENS`, because that
#: module already imports `KIND_ORDER` FROM this one -- importing this
#: constant back the other way would be circular). Rep-gate lane,
#: 2026-09-30, this lane's brief item 1: Cenarion Circle's own
#: reputation is Gates of Ahn'Qiraj (Patch 1.9) content, the same patch
#: this build's own `raid:ahnqiraj` source is already gated `"later"`
#: for (`curated/loot/forever-raid-phases.json`), so a quest that needs
#: ANY standing with it cannot be turned in any earlier than the raid
#: itself opens. Writing this fact into `QuestSource.opens` here (rather
#: than only in the Go ranker, where `firstNonEmpty(src.Opens,
#: repFactionRaidPhaseOpens[factionID])` already covers `rep`-kind
#: sources) is what fixes a quest reward: `loot.json` states the gate
#: directly, and the ranker's OWN table becomes a fallback for whatever
#: this one is silent on, exactly the same relationship `firstNonEmpty`
#: already has with a curated raid/dungeon `opens`.
REP_FACTION_RAID_PHASE_OPENS: dict[int, str] = {
    609: "later",  # Cenarion Circle - Gates of Ahn'Qiraj (AQ War Effort)
}


def _rep_gate_by_item(sources: list[LootSource]) -> dict[int, tuple[int, str]]:
    """item id -> (faction_id, standing), for every `rep`-kind
    `LootSource` -- `apply_quest_opens_gate`'s own fallback signal for a
    quest reward classic-db's `quest_template` states no reputation
    requirement for directly (Earthstrike, item 21180, quest 8573
    "Champion's Battlegear": the fork database already names that exact
    item from a `rep:cenarion-circle:exalted` source, so the quest
    reward inherits that source's own gate rather than bypassing it)."""
    by_item: dict[int, tuple[int, str]] = {}
    for source in sources:
        if source.kind != "rep" or source.faction_id is None or source.standing is None:
            continue
        for item_id in source_item_ids(source):
            by_item.setdefault(item_id, (source.faction_id, source.standing))
    return by_item


def apply_quest_opens_gate(
    document: LootFile, classic_sources: dict[int, list[ClassicDbSourceRecord]]
) -> LootFile:
    """Sets `QuestSource.opens` for every quest in `document.quests`
    gated one of two ways.

    First, a REPUTATION gate (rep-gate lane, 2026-09-30, this lane's
    brief item 1): `entry.required_rep_faction`, whether classic-db's own
    `quest_template` row states it directly or (Earthstrike's own case)
    it is filled in here as a fallback from the same item's `rep`-kind
    `LootSource` (`_rep_gate_by_item`) -- `REP_FACTION_RAID_PHASE_OPENS`
    turns that faction id into the phase its own reputation track needs.

    Second (unchanged since quest-gates lane, 2026-09-29), whether the
    quest's classic-db turn-in item(s) (`quest_turn_in_items_from_
    classic_sources`, already resolved through the quest's own
    `PrevQuestId` chain) are THEMSELVES only obtainable from a source
    `opens` is already set on: a raid boss drop directly (Onyxia's
    Lair's "Head of Onyxia" gates "For All To See"/"Celebrating Good
    Times"; Ruins of Ahn'Qiraj's "Head of Ossirian" gates "The Fall of
    Ossirian"), or another such quest, recursively (the fixed-point loop
    below -- an item that is itself a reward of an already-gated quest).

    Must run AFTER `apply_overlays`: a raid's own `LootSource.opens` is
    itself an overlay fact (`curated/loot/forever-raid-phases.json`) --
    nothing upstream of the overlay ever sets it, so calling this any
    earlier would find every raid source ungated and gate nothing.

    A turn-in item this build's `document` names NO source for at all
    is left alone -- never invented into either a gate or a guarantee of
    none (tenet 8). A turn-in item with even ONE un-gated way to get it
    (a vendor, a farmed drop, an alternate quest with no raid gate of
    its own) leaves the quest open, same as today: `SrcItemId`/
    `ReqItemId` on the overwhelming majority of quests names an ordinary
    farmed/vendored item, and this function must never gate one of
    those.
    """
    quest_turn_ins = quest_turn_in_items_from_classic_sources(classic_sources)
    # Every NON-"quest"-kind source's own opens, per item id it names --
    # the flat "quest" LootSource carries no per-quest opens of its own
    # (LootSource.opens is never set on it), so it is excluded here the
    # same way sim/cmd/leveling-bis/data.go's loadLootIndex already
    # skips that kind in favour of the per-quest detail below.
    non_quest_ways: dict[int, list[str]] = defaultdict(list)
    for source in document.sources:
        if source.kind == "quest":
            continue
        opens = source.opens or ""
        for item_id in source_item_ids(source):
            non_quest_ways[item_id].append(opens)

    quests: dict[str, list[QuestSource]] = {
        item_id: list(entries) for item_id, entries in document.quests.items()
    }

    rep_gate_by_item = _rep_gate_by_item(document.sources)
    for item_id_str, entries in quests.items():
        item_id = int(item_id_str)
        for i, entry in enumerate(entries):
            faction_id, standing = entry.required_rep_faction, entry.required_rep_standing
            if faction_id is None:
                fallback = rep_gate_by_item.get(item_id)
                if fallback is None:
                    continue
                faction_id, standing = fallback
                entries[i] = entry = entry.model_copy(
                    update={
                        "required_rep_faction": faction_id,
                        "required_rep_standing": standing,
                    }
                )
            if entry.opens:
                continue
            gate = REP_FACTION_RAID_PHASE_OPENS.get(faction_id)
            if gate:
                entries[i] = entry.model_copy(update={"opens": gate})

    def item_gate(item_id: int) -> str | None:
        ways = [*non_quest_ways.get(item_id, [])]
        ways += [entry.opens or "" for entry in quests.get(str(item_id), [])]
        if not ways or any(not way for way in ways):
            # No known source at all, or at least one un-gated way --
            # either way, this item never gates a quest that needs it.
            return None
        return sorted(set(ways))[0]

    for _ in range(_MAX_QUEST_GATE_PASSES):
        changed = False
        for entries in quests.values():
            for i, entry in enumerate(entries):
                if entry.opens:
                    continue
                required = quest_turn_ins.get(entry.quest_id) or []
                gate = next((g for g in (item_gate(item_id) for item_id in required) if g), None)
                if gate:
                    entries[i] = entry.model_copy(update={"opens": gate})
                    changed = True
        if not changed:
            break
    return document.model_copy(update={"quests": quests})


def item_effective_gate(document: LootFile, item_id: int) -> str | None:
    """The SAME "does this item have even one un-gated way to get it"
    check `apply_quest_opens_gate`'s own `item_gate` closure runs
    mid-fixed-point-loop, but over the FINAL, fully-resolved `document`
    (every raid/dungeon `LootSource.opens` an overlay set, every
    `QuestSource.opens` `apply_quest_opens_gate` itself already
    computed) -- `apply_crafted_opens_gate`'s own use, for a recipe or
    reagent item, needs no fixed-point loop of its own: by the time it
    runs (after `apply_quest_opens_gate`), nothing it reads changes
    again.

    `None` for an item with no known source at all (never invented into
    a gate, tenet 8) OR at least one un-gated way; otherwise the sole
    (or, when more than one way is gated to a different phase,
    deterministically first-sorted) gate value -- identical semantics to
    `apply_quest_opens_gate`'s own `item_gate`, just not tied to its
    local, in-progress `quests`/`non_quest_ways` dicts.
    """
    ways = [
        source.opens or ""
        for source in document.sources
        if source.kind != "quest" and item_id in source_item_ids(source)
    ]
    ways += [entry.opens or "" for entry in document.quests.get(str(item_id), [])]
    if not ways or any(not way for way in ways):
        return None
    return sorted(set(ways))[0]


#: `apply_crafted_opens_gate`'s own two known phase values, LEAST
#: restrictive first -- `curated/loot/forever-raid-phases.json`'s own
#: notes: "raids-1" (Onyxia's Lair, the one announced-date raid, opens 9
#: December) opens strictly BEFORE "later" (every other raid, no
#: announced date at all). An alternative recipe source or reagent this
#: tuple does not name (a future third phase this file's own curators
#: have not added here yet) sorts LAST -- treated as the MOST
#: restrictive rather than guessed into either known slot, so a new
#: phase value never silently under-gates an item.
_CRAFTED_OPENS_ORDER = ("raids-1", "later")


def _opens_rank(value: str) -> int:
    try:
        return _CRAFTED_OPENS_ORDER.index(value)
    except ValueError:
        return len(_CRAFTED_OPENS_ORDER)


def _least_restrictive_gate(gates: list[str]) -> str | None:
    """For an OR relationship (alternative recipe items -- only one is
    ever needed): `None` the moment any one alternative is itself
    ungated (`item_effective_gate` already returned `None` for it), else
    whichever named gate opens soonest."""
    if not gates:
        return None
    return min(gates, key=_opens_rank)


def _most_restrictive_gate(gates: list[str]) -> str | None:
    """For an AND relationship (every reagent is needed, and the recipe
    AND its reagents together): the LATEST-opening gate among every
    non-`None` input, or `None` when none of them gates at all."""
    named = [gate for gate in gates if gate]
    if not named:
        return None
    return max(named, key=_opens_rank)


def _crafted_item_gate(document: LootFile, recipe: ClassicDbCraftedRecipe) -> str | None:
    """`opens` for one crafted item, per this lane's own brief: the most
    restrictive of (a) its recipe item(s) -- `None` outright when
    trainer-taught (`recipe_item_ids` empty, always open) -- and (b) its
    reagents, each checked as an ordinary item (a reagent with NO known
    source at all is silently skipped, never invented into a gate,
    exactly `item_effective_gate`'s own "no source" case; a reagent
    obtainable from even one un-gated place is likewise not a gate)."""
    if recipe.recipe_item_ids:
        recipe_ways = [item_effective_gate(document, item_id) for item_id in recipe.recipe_item_ids]
        recipe_gate = None if any(way is None for way in recipe_ways) else _least_restrictive_gate(
            [way for way in recipe_ways if way]
        )
    else:
        recipe_gate = None  # trainer-taught: no item stands in the way at all
    reagent_gates = [item_effective_gate(document, item_id) for item_id in recipe.reagent_item_ids]
    reagent_gate = _most_restrictive_gate(reagent_gates)
    return _most_restrictive_gate([recipe_gate, reagent_gate])


def apply_crafted_opens_gate(
    document: LootFile, classic_crafted: dict[int, ClassicDbCraftedRecipe]
) -> LootFile:
    """Splits each `crafted:<profession>` `LootSource` into its ungated
    items (kept on the original id, `opens` still unset) and, for every
    distinct non-empty gate `_crafted_item_gate` computes, a same-
    profession `crafted:<profession>:<phase>` sibling with `opens` set --
    the shape `sim/cmd/leveling-bis/data.go`'s own `loadLootIndex`
    already understands with NO Go change at all: it applies one
    `itemSource.Opens` per SOURCE (`lootSource.Opens`' own doc), never a
    per-item map for a `crafted`-kind source, so two differently-gated
    items sharing one profession bucket need two source ids, not one
    (this lane's brief, item 1, asked which shape the ranker already
    reads; this is it).

    Must run AFTER `apply_overlays` (a raid/dungeon source's own `opens`
    is itself an overlay fact) AND AFTER `apply_quest_opens_gate` (a
    recipe item taught by a quest needs that quest's own `opens` already
    resolved) -- same ordering rule both of those already state for
    themselves.

    A crafted item id `classic_crafted` does not cover (a Forever-new
    id, or a Classic id classic-db's own dump names no create-item spell
    for -- `pipeline.audit.check_crafted`'s own finding) is left in its
    original, ungated bucket -- never invented into a gate.
    """
    if not classic_crafted:
        return document
    new_sources: list[LootSource] = []
    gated: dict[tuple[str, str], set[int]] = defaultdict(set)
    changed = False
    for source in document.sources:
        if source.kind != "crafted" or not source.items:
            new_sources.append(source)
            continue
        ungated_items: list[int] = []
        for item_id in source.items:
            recipe = classic_crafted.get(item_id)
            gate = _crafted_item_gate(document, recipe) if recipe is not None else None
            if gate:
                gated[(source.profession or "", gate)].add(item_id)
                changed = True
            else:
                ungated_items.append(item_id)
        if ungated_items:
            new_sources.append(source.model_copy(update={"items": sorted(ungated_items)}))
    if not changed:
        return document
    for (profession, phase), item_ids in sorted(gated.items()):
        new_sources.append(
            LootSource(
                id=f"crafted:{profession}:{phase}",
                kind="crafted",
                name=profession.replace("-", " ").title(),
                profession=profession,
                items=sorted(item_ids),
                opens=phase,
            )
        )
    new_sources.sort(key=lambda source: (KIND_ORDER.index(source.kind), source.id))
    return document.model_copy(update={"sources": new_sources})


def _resolve_or_drop_unnamed_bosses(
    sources: list[LootSource], npc_names: dict[int, str]
) -> tuple[list[LootSource], int]:
    """Every `LootBoss` with an empty `name` (`LootBoss.name`'s own doc:
    the fork database simply does not name every raid/dungeon npc --
    measured 58 on build 1.60.1.70009, every one a raid boss) gets
    `npc_names`' (`pipeline.loot.classicdb.classic_db_npc_names`'s own
    result) name for its npc_id when one exists; a boss STILL unnamed
    after that is DROPPED from its source entirely rather than published
    empty -- wowhead-world-drops lane, 2026-09-29's own addendum: the
    site renders an empty name as "0% from Unnamed source in <zone>",
    which is a worse fact than the boss simply not being listed at all
    (tenet 7's "one line a player recognises", tenet 8's "unverifiable
    is labelled or left out, never shown as fact"). Runs at the END of
    `build_loot`, after every origin (fork, classic-db, wowhead) has
    already merged, so it resolves a name regardless of which origin's
    union kept an empty one.

    Returns the resolved source list and how many bosses were dropped,
    for `LootStats.dropped_unnamed_bosses` and the build's own log line.
    A source a drop leaves with no bosses and no other items is pruned
    by `build_loot`'s own final `source_item_ids` sweep, same as any
    other source a filter emptied.
    """
    dropped = 0
    resolved: list[LootSource] = []
    for source in sources:
        if not source.bosses:
            resolved.append(source)
            continue
        kept_bosses: list[LootBoss] = []
        for boss in source.bosses:
            if boss.name:
                kept_bosses.append(boss)
                continue
            real_name = npc_names.get(boss.npc_id)
            if real_name:
                kept_bosses.append(boss.model_copy(update={"name": real_name}))
                continue
            dropped += 1
            logger.warning(
                "loot: dropping boss npc %d from %s -- neither the fork database nor "
                "classic-db's own creature_template names it; publishing an empty boss "
                "name would be worse than omitting it",
                boss.npc_id,
                source.id,
            )
        resolved.append(source.model_copy(update={"bosses": kept_bosses or None}))
    return resolved, dropped


def _classic_db_creature_items(
    classic_sources: dict[int, list[ClassicDbSourceRecord]],
) -> dict[int, set[int]]:
    """npc_id -> every item id classic-db's own `creature_loot_template`
    (direct or reference -- `pipeline.classic_sources.parse_classic_db_
    sources` already expanded a reference row into this SAME `kind=
    "creature_drop"` shape) names for it. `pipeline.audit.check_drops`'
    own `_boss_items` check reads the identical fact off `classic_
    sources_by_item`; this is that same corroboration test, inverted to
    npc-keyed so `_boss_item_source_origins` can ask it once per boss
    rather than once per (boss, item) pair.
    """
    by_npc: dict[int, set[int]] = defaultdict(set)
    for item_id, records in classic_sources.items():
        for record in records:
            if record.kind == "creature_drop" and record.npc_id:
                by_npc[record.npc_id].add(item_id)
    return dict(by_npc)


def _boss_item_source_origins(
    sources: list[LootSource],
    fork_boss_items: set[tuple[int, int]],
    classic_sources: dict[int, list[ClassicDbSourceRecord]],
) -> list[LootSource]:
    """`LootBoss.item_source_origin`'s own doc: tag every dungeon/raid
    boss item classic-db does NOT corroborate with whichever origin --
    `"fork"` or `"wowhead"` -- actually named it, drop-sources-2 lane
    2026-09-29. `fork_boss_items` (every (npc_id, item_id) pair the
    fork's OWN `_drop_sources` result names, captured in `build_loot`
    BEFORE classic-db/wowhead merge in) is what tells the two apart: a
    boss item classic-db does not corroborate and the fork's own
    pre-merge data does not name either can only have reached this boss
    through `wowhead_additions`' own union merge -- fork, classic-db and
    wowhead are the only three origins a boss item can have at all.

    Runs at the end of `build_loot`, after every merge, alongside
    `_resolve_or_drop_unnamed_bosses` -- same reasoning: it needs to see
    the FINAL boss item list regardless of which origin's union produced
    it.
    """
    classic_items_by_npc = _classic_db_creature_items(classic_sources)

    def tag(boss: LootBoss) -> LootBoss:
        origins = {
            str(item_id): "fork" if (boss.npc_id, item_id) in fork_boss_items else "wowhead"
            for item_id in boss.items
            if item_id not in classic_items_by_npc.get(boss.npc_id, set())
        }
        return boss if not origins else boss.model_copy(update={"item_source_origin": origins})

    return [
        source.model_copy(update={"bosses": [tag(boss) for boss in source.bosses]})
        if source.bosses
        else source
        for source in sources
    ]


def _classic_db_corroborated_raid_dungeon_origin(
    sources: list[LootSource],
    fork_boss_items: set[tuple[int, int]],
    classic_sources: dict[int, list[ClassicDbSourceRecord]],
) -> list[LootSource]:
    """loot-parity-2 lane, 2026-09-30: promotes a raid/dungeon
    `LootSource.source_origin` from `None` (fork-primary) to
    `"classic-db"` when classic-db's own `creature_loot_template`
    independently corroborates EVERY item the fork's own `_drop_sources`
    pass put on that source (`fork_boss_items`, same pre-merge capture
    `_boss_item_source_origins` already uses) -- never when trash is
    present, since a zone's flat `trash` bucket has no per-npc key to
    check corroboration against.

    Onyxia's Lair is the measured case (this lane's own report): the
    fork's own AtlasLoot table names 16 items for npc 10184, and
    classic-db's own dump -- read fresh -- independently names all 16
    PLUS 5 more, so the fork contributes nothing classic-db does not
    already confirm; `_union_source`'s "base (fork) stays primary, its
    own `source_origin` is left alone" rule was correct for a source
    where fork and classic-db name DISJOINT items (`dungeon:the-
    deadmines`'s own two-boss test: classic-db's npc 657 corroborates
    nothing about the fork's own npc 902/item 103, so that source stays
    fork-primary), but wrong here, where classic-db has independently
    verified the fork's entire claim and is the more authoritative table
    (tenet 8: a primary source, not a legacy migration) -- `source_
    origin` should say so.

    A source with even one fork-only (unconfirmed by classic-db) boss
    item is left alone: the fork still contributes something classic-db
    cannot itself vouch for, so `None` (fork-primary, with that one item
    already carrying its own `item_source_origin: "fork"` tag from
    `_boss_item_source_origins` above) remains the honest label.
    """
    classic_items_by_npc = _classic_db_creature_items(classic_sources)

    def fully_corroborated(source: LootSource) -> bool:
        fork_items = [
            (boss.npc_id, item_id)
            for boss in (source.bosses or [])
            for item_id in boss.items
            if (boss.npc_id, item_id) in fork_boss_items
        ]
        if not fork_items or source.trash:
            return False
        return all(
            item_id in classic_items_by_npc.get(npc_id, set()) for npc_id, item_id in fork_items
        )

    return [
        source.model_copy(update={"source_origin": "classic-db"})
        if source.kind in ("raid", "dungeon")
        and source.source_origin is None
        and fully_corroborated(source)
        else source
        for source in sources
    ]


def build_loot(
    fork: ForkDatabase,
    zone_names: dict[int, str],
    types: dict[int, int],
    ranks: dict[int, int],
    build_items: set[int],
    item_inventory_types: dict[int, int],
    quest_levels: dict[int, QuestLevelEntry] | None = None,
    item_sources: dict[int, ItemSourceEntry] | None = None,
    classic_sources: dict[int, list[ClassicDbSourceRecord]] | None = None,
    zone_rows: list[dict] | None = None,
    required_levels: dict[int, int] | None = None,
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
    # `_boss_item_source_origins`' own doc: captured BEFORE classic-db/
    # wowhead merge in, so it is the fork's OWN, unmixed boss item list --
    # the one fact that lets that function tell "wowhead's own union
    # merge added this" apart from "the fork always named this".
    fork_boss_items = {
        (boss.npc_id, item_id)
        for source in drops
        for boss in (source.bosses or [])
        for item_id in boss.items
    }
    keyed, quest, quest_detail, dropped_keyed = _keyed_sources(
        fork, build_items, equippable, absent, quest_levels or {}
    )
    sources = [*drops, *keyed, *_pvp_sources(ranks, build_items)]
    # Fork-only coverage, INCLUDING the quest items `quest`/`quest_detail`
    # already name at this point (the flat "quest" LootSource itself is
    # appended below, after any wowhead quest items join `quest`) --
    # `LootStats.wowhead_items` below is what a wowhead scrape adds ON
    # TOP of this set, so quest items the fork itself already covers
    # must count as fork's, not wowhead's.
    fork_named = {
        item_id for source in sources for item_id in source_item_ids(source)
    } | set(quest)
    if classic_sources:
        # src-classicdb lane, 2026-09-29: applied BEFORE wowhead (priority
        # order fork > classic-db > wowhead) -- a bucket classic-db
        # creates first gets source_origin="classic-db"; if wowhead later
        # names the same bucket for an overlapping item, the union below
        # keeps THIS origin, not wowhead's, since merge unions extra into
        # base and only a brand-new id keeps extra's own origin.
        classicdb_sources, classicdb_quest, classicdb_quest_detail = classicdb_additions(
            classic_sources, build_items, equippable, zone_names, types, zone_rows or [],
            quest_levels or {}, item_factions(fork, build_items),
            fork_instance_npc_zones(fork, types),
        )
        sources = merge_classicdb_sources(sources, classicdb_sources)
        quest = sorted(set(quest) | set(classicdb_quest))
        for item_id, entries in classicdb_quest_detail.items():
            quest_detail[item_id] = [*quest_detail.get(item_id, []), *entries]
    # Fork + classic-db coverage, the baseline `LootStats.wowhead_items`
    # is measured against -- a quest item classic-db already covers must
    # count as classic-db's, not wowhead's, same reasoning as `fork_named`
    # above.
    fork_and_classicdb_named = {
        item_id for source in sources for item_id in source_item_ids(source)
    } | set(quest)
    if item_sources:
        # night-item-sources lane, 2026-09-28: fills the gap ABOVE, never
        # replaces a fork-found source -- `merge_wowhead_sources` unions
        # into an existing id and only appends a wholly new one.
        #
        # wowhead-world-drops lane, 2026-09-29: every item id ALREADY in
        # a `world_drop` source at this point got there from classic-db
        # (fork itself never emits the kind) -- `wowhead_additions`' own
        # precedence rule skips these entirely rather than resurrecting
        # their per-creature wowhead rows or appending a second,
        # differently-leveled `world_drop` source under a different id.
        classicdb_world_drop_items = {
            item_id
            for existing in sources
            if existing.kind == "world_drop"
            for item_id in source_item_ids(existing)
        }
        wowhead_sources, wowhead_quest, wowhead_quest_detail = wowhead_additions(
            item_sources, build_items, equippable, zone_names, types,
            item_factions(fork, build_items),
            required_levels,
            classicdb_world_drop_items,
            fork.factions,
        )
        sources = merge_wowhead_sources(sources, wowhead_sources)
        quest = sorted(set(quest) | set(wowhead_quest))
        for item_id, entries in wowhead_quest_detail.items():
            quest_detail[item_id] = [*quest_detail.get(item_id, []), *entries]
    # quest-faction lane, 2026-09-29, extended by the data-followups-3
    # lane, 2026-09-30 (item 2): the quest's own primary-source faction
    # wins over every `QuestSource.faction` built above (fork, classic-db
    # and wowhead alike all currently guess a quest's faction from the
    # reward ITEM's own `factionRestriction`, which can legitimately
    # disagree with the quest, or -- 78150 Friend of the Library -- carry
    # no restriction at all even though the quest itself is single-
    # faction). Applied once here, after every scrape's quest_detail has
    # been folded in, so it covers a quest id regardless of which one
    # produced the link -- see `resolve_quest_faction`'s own doc for the
    # full fallback order (classic-db, then the item's own REAL
    # restriction, then wowhead, then "unknown" -- never a defaulted
    # "both" republished as fact).
    quest_factions = quest_factions_from_classic_sources(classic_sources or {})
    quest_levels_by_id = quest_levels or {}
    item_sources_by_id = item_sources or {}
    item_faction_restrictions = item_factions(fork, build_items)
    # Day3 data-followups-7 lane, 2026-09-30: the quest's own `classes`/
    # `profession`+`skill` (`ClassicDbQuestInfo`'s own facts, both
    # `RequiredClasses`/`RequiredSkill` columns this pipeline did not read
    # before this lane) applied the SAME way `quest_factions` above
    # already is -- once here, after every scrape's quest_detail has been
    # folded in, so a class- or skill-gated quest is covered regardless
    # of which one (fork, classic-db or wowhead) produced the item link.
    # A quest id absent from either dict simply keeps whatever the
    # freshly-built `QuestSource` already carries (`None`/`None`/`None`:
    # neither the fork nor wowhead publishes either fact today).
    quest_classes = quest_classes_from_classic_sources(classic_sources or {})
    quest_professions = quest_profession_from_classic_sources(classic_sources or {})
    quest_detail = {
        item_id: [
            entry.model_copy(
                update=dict(
                    zip(
                        ("faction", "faction_source"),
                        resolve_quest_faction(
                            entry.quest_id,
                            item_id,
                            quest_factions,
                            quest_levels_by_id,
                            item_sources_by_id,
                            item_faction_restrictions,
                        ),
                        strict=True,
                    ),
                    classes=quest_classes.get(entry.quest_id),
                    profession=(quest_professions.get(entry.quest_id) or (None, None))[0],
                    skill=(quest_professions.get(entry.quest_id) or (None, None))[1],
                )
            )
            for entry in entries
        ]
        for item_id, entries in quest_detail.items()
    }
    if quest:
        sources.append(LootSource(id="quest", kind="quest", name="Quests", items=quest))
    # wowhead-world-drops lane, 2026-09-29's own addendum: resolves an
    # empty `LootBoss.name` from classic-db's own creature_template, or
    # drops the boss when neither origin names it -- runs after every
    # origin has merged (fork, classic-db, wowhead) so it catches an
    # empty name regardless of which origin's union kept one, and BEFORE
    # the pruning sweep below so a source a drop leaves empty is removed
    # the same way any other emptied source already is.
    sources, dropped_unnamed_bosses = _resolve_or_drop_unnamed_bosses(
        sources, classic_db_npc_names(classic_sources) if classic_sources else {}
    )
    # drop-sources-2 lane, 2026-09-29: labels every boss item classic-db
    # does not corroborate with whichever origin (fork/wowhead) actually
    # named it -- `_boss_item_source_origins`' own doc -- same "after
    # every merge" placement as the unnamed-boss resolution just above,
    # for the same reason (it needs the FINAL, post-merge boss lists).
    sources = _boss_item_source_origins(sources, fork_boss_items, classic_sources or {})
    # loot-parity-2 lane, 2026-09-30: promotes a raid/dungeon source's own
    # `source_origin` to "classic-db" when classic-db's dump fully
    # corroborates the fork's own contribution to it --
    # `_classic_db_corroborated_raid_dungeon_origin`'s own doc (the
    # Onyxia's Lair case this lane's report measured). Runs after
    # `_boss_item_source_origins`, same "needs the FINAL, post-merge boss
    # lists" reasoning.
    sources = _classic_db_corroborated_raid_dungeon_origin(
        sources, fork_boss_items, classic_sources or {}
    )
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
        wowhead_items=len(named - fork_and_classicdb_named),
        classicdb_items=len(fork_and_classicdb_named - fork_named),
        dropped_unnamed_bosses=dropped_unnamed_bosses,
    )
