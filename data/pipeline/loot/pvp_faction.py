"""PvP honor-rank source faction split (third wow-player sweep defect,
2026-09-29): `pipeline.loot.sources.build_loot`'s own `pvp:rank-N`
buckets mix Alliance and Horde titled rewards under one numeric id --
Blizzard's own `RequiredPVPRank` column names a RANK, not a side, so
"Knight-Lieutenant's Pauldrons" (Alliance) and "Blood Guard's
Pauldrons" (Horde) land in the exact same `pvp:rank-11` bucket. The
site then has no fact to gate a hybrid spec's ranker on (`sim/cmd/
leveling-bis/data.go`'s own `Side`), and the bare bucket name "Rank 9"
reads like neither faction's own Quartermaster to a player.

`split_pvp_sources_by_faction` is `build_loot`'s own post-step, run the
same way `pipeline.loot.reitemise.apply_reitemisation` is -- AFTER
`build_loot`, BEFORE the reitemisation pass, so a Forever-new item
copied onto a classic twin's source inherits whichever faction split
that source already landed in. Every `kind == "pvp"` source becomes up
to two, `pvp:rank-N:alliance` / `pvp:rank-N:horde`, name "Rank N
(Alliance)" / "Rank N (Horde)", `rank` unchanged.

Two-tier resolution, primary source first:

1. `vendor_npc_factions` -- the classic-db `npc_vendor` row selling the
   CLASSIC item id (`pipeline.classic_sources.ClassicDbSourceRecord`,
   kind `"vendor"`, `map_id` now carried for this kind too -- see that
   module's own doc). Ranks 10 through 18 are sold out of one of
   exactly two capital-city PvP halls classic-db's own `creature` spawn
   table places the vendor in -- this build's own `zones.json` names
   them plainly ("Alliance PVP Barracks"/"Horde PVP Barracks", zone
   2918 Champions' Hall and 2917 Hall of Legends) -- which settles the
   vendor's side outright. Ranks 5 through 9's own two vendors (Officer
   Areyn, Sergeant Ba'sha) stand in their OWN faction's capital on the
   open continent map (Eastern Kingdoms/Kalimdor), which both sides
   hold territory on, so the map alone settles nothing there;
   `vendor_npc_factions` instead reads whichever side the SAME vendor's
   own OTHER items already prove via `title_faction` below (unanimous
   only -- a vendor whose titled items disagree resolves nothing,
   logged rather than guessed at) -- still a classic-db `npc_vendor`
   fact, not a title guess, so this still counts as
   `faction_source: "classic-db"`.

   (`creature_template.Faction`, the pinned dump's only per-npc
   team-adjacent column, is a `FactionTemplate.dbc` id this dump
   carries no table to resolve -- measured values 35, 12 and 29
   corroborate nothing (two Alliance-flavoured and two Horde-flavoured
   rank-10+ quartermasters all share Faction 35), and
   `creature_template_armor.FactionAlliance`/`FactionHorde`, the dump's
   only table with those exact column names, is zero for every
   quartermaster checked -- a legacy compatibility table this pinned
   dump never populates. Neither is a usable primary source here; the
   capital-city PvP hall map is.)

2. `title_faction` -- the item's OWN display name against Vanilla's own
   14-rank title ladder per side (owner brief, 2026-09-29). Settles
   every item whose vendor did not (`faction_source: "title"`), and
   agreed with every vendor-resolved item that also carried an
   unambiguous title when this lane measured build 1.60.1.70009's own
   659 pvp item/rank pairs (0 disagreements, 0 unresolved).

An item neither tier resolves is dropped from both split buckets and
counted in `PvpFactionStats.unresolved` -- tenet 8's "unverifiable is
labelled or left out, never shown as fact" -- rather than published
under a guessed side.
"""

from __future__ import annotations

import logging
from collections import defaultdict
from dataclasses import dataclass
from typing import Literal

from pipeline.classic_sources import ClassicDbSourceRecord
from pipeline.models import LootFile, LootSource

logger = logging.getLogger(__name__)

Faction = Literal["alliance", "horde"]

#: Vanilla's own Alliance PvP rank ladder, ranks 1-14 in order
#: (Blizzard's own `RequiredPVPRank` item column is these ranks + 4) --
#: owner brief, 2026-09-29.
ALLIANCE_TITLES: tuple[str, ...] = (
    "Private", "Corporal", "Sergeant", "Master Sergeant", "Sergeant Major",
    "Knight", "Knight-Lieutenant", "Knight-Captain", "Knight-Champion",
    "Lieutenant Commander", "Commander", "Marshal", "Field Marshal", "Grand Marshal",
)  # fmt: skip

#: Vanilla's own Horde PvP rank ladder, ranks 1-14, same source and
#: reasoning as `ALLIANCE_TITLES`.
HORDE_TITLES: tuple[str, ...] = (
    "Scout", "Grunt", "Sergeant", "Senior Sergeant", "First Sergeant",
    "Stone Guard", "Blood Guard", "Legionnaire", "Centurion", "Champion",
    "Lieutenant General", "General", "Warlord", "High Warlord",
)  # fmt: skip

#: title -> faction, EXCLUDING a title both ladders name ("Sergeant",
#: rank 3 on both sides -- the one rank Blizzard never gave the two
#: factions distinct wording for): `title_faction` must never guess
#: between the two sides for a title this ambiguous, only settle one
#: that is not.
TITLE_FACTION: dict[str, Faction] = {
    **{title: "alliance" for title in ALLIANCE_TITLES},
    **{title: "horde" for title in HORDE_TITLES},
}
for _shared in set(ALLIANCE_TITLES) & set(HORDE_TITLES):
    del TITLE_FACTION[_shared]

#: Longest title first, so a compound title ("Master Sergeant's
#: Insignia") is checked before any shorter title that happens to
#: prefix it could be -- the ordering is the contract here, not a
#: coincidence of today's ladder.
_TITLES_LONGEST_FIRST: tuple[str, ...] = tuple(sorted(TITLE_FACTION, key=len, reverse=True))


def title_faction(item_name: str) -> Faction | None:
    """`item_name`'s own rank title -- Vanilla's PvP reward naming is
    always "<Title>'s <slot>" (e.g. "Knight-Lieutenant's Pauldrons") --
    or literal faction word ("Insignia of the Alliance"/"Insignia of
    the Horde" name no rank title at all), matched against
    `TITLE_FACTION`. `None` when neither pattern matches, or the
    matched title is the one both sides share ("Sergeant") and no
    vendor already settled it.
    """
    for title in _TITLES_LONGEST_FIRST:
        if item_name.startswith(f"{title}'s "):
            return TITLE_FACTION[title]
    if item_name.endswith("of the Alliance"):
        return "alliance"
    if item_name.endswith("of the Horde"):
        return "horde"
    return None


def _barracks_map_factions(zone_rows: list[dict]) -> dict[int, Faction]:
    """map id -> faction, for every map this build's own `zones.json`
    names as one side's PvP hall ("Alliance PVP Barracks"/"Horde PVP
    Barracks") -- matched by substring on `zones.json`'s own `map_name`
    rather than a hardcoded map id, since that id is this build's own
    DB2 export, not a Blizzard constant worth pinning by number.
    """
    by_map: dict[int, Faction] = {}
    for row in zone_rows:
        map_id = row.get("map_id")
        if map_id is None:
            continue
        name = (row.get("map_name") or "").lower()
        if "alliance" in name:
            by_map.setdefault(int(map_id), "alliance")
        elif "horde" in name:
            by_map.setdefault(int(map_id), "horde")
    return by_map


def vendor_npc_factions(
    classic_sources: dict[int, list[ClassicDbSourceRecord]],
    zone_rows: list[dict],
    item_names: dict[int, str],
) -> dict[int, Faction]:
    """npc_id -> faction, for every classic-db vendor npc this build's
    pvp rank items name -- see the module doc's two-tier explanation.
    """
    barracks = _barracks_map_factions(zone_rows)
    npc_items: dict[int, set[int]] = defaultdict(set)
    npc_map_id: dict[int, int | None] = {}
    for item_id, records in classic_sources.items():
        for record in records:
            if record.kind == "vendor" and record.npc_id:
                npc_items[record.npc_id].add(item_id)
                npc_map_id.setdefault(record.npc_id, record.map_id)

    factions: dict[int, Faction] = {}
    for npc_id, item_ids in npc_items.items():
        map_id = npc_map_id.get(npc_id)
        if map_id is not None and map_id in barracks:
            factions[npc_id] = barracks[map_id]
            continue
        votes: set[Faction] = set()
        for item_id in item_ids:
            name = item_names.get(item_id)
            faction = title_faction(name) if name else None
            if faction:
                votes.add(faction)
        if len(votes) == 1:
            factions[npc_id] = votes.pop()
        elif len(votes) > 1:
            logger.warning(
                "pvp-faction: classic-db npc %d sells items whose own titles disagree "
                "(%s) -- not resolved via vendor",
                npc_id,
                sorted(votes),
            )
    return factions


def _item_vendor_npcs(
    classic_sources: dict[int, list[ClassicDbSourceRecord]], item_id: int
) -> list[int]:
    return [
        record.npc_id
        for record in classic_sources.get(item_id, [])
        if record.kind == "vendor" and record.npc_id
    ]


@dataclass(frozen=True)
class PvpFactionStats:
    """Lane report counters for `split_pvp_sources_by_faction`: how many
    pvp item/rank pairs resolved via the classic-db vendor tier, how
    many via the title fallback, and which item ids resolved via
    neither (dropped from both split buckets)."""

    resolved_vendor: int = 0
    resolved_title: int = 0
    unresolved: tuple[int, ...] = ()


def split_pvp_sources_by_faction(
    document: LootFile,
    item_names: dict[int, str],
    classic_sources: dict[int, list[ClassicDbSourceRecord]] | None,
    zone_rows: list[dict] | None,
) -> tuple[LootFile, PvpFactionStats]:
    """`document` with every `kind == "pvp"` source replaced by its
    alliance/horde split (module doc). Runs BEFORE
    `pipeline.loot.reitemise.apply_reitemisation` so a re-itemised item
    copied onto a classic twin's source lands in the SAME
    already-split bucket its twin resolved to.
    """
    classic_sources = classic_sources or {}
    npc_factions = vendor_npc_factions(classic_sources, zone_rows or [], item_names)

    resolved_vendor = 0
    resolved_title = 0
    unresolved: list[int] = []
    sources: list[LootSource] = []
    for source in document.sources:
        if source.kind != "pvp":
            sources.append(source)
            continue
        by_faction: dict[Faction, list[int]] = {"alliance": [], "horde": []}
        origin_for: dict[Faction, Literal["classic-db", "title"]] = {}
        for item_id in source.items or []:
            faction: Faction | None = None
            origin: Literal["classic-db", "title"] | None = None
            for npc_id in _item_vendor_npcs(classic_sources, item_id):
                if npc_id in npc_factions:
                    faction, origin = npc_factions[npc_id], "classic-db"
                    break
            if faction is None:
                name = item_names.get(item_id)
                title = title_faction(name) if name else None
                if title is not None:
                    faction, origin = title, "title"
            if faction is None or origin is None:
                unresolved.append(item_id)
                continue
            by_faction[faction].append(item_id)
            origin_for.setdefault(faction, origin)
            if origin == "classic-db":
                resolved_vendor += 1
            else:
                resolved_title += 1
        for faction in ("alliance", "horde"):
            items = sorted(by_faction[faction])
            if not items:
                continue
            sources.append(
                source.model_copy(
                    update={
                        "id": f"{source.id}:{faction}",
                        "name": f"{source.name} ({faction.capitalize()})",
                        "items": items,
                        "faction": faction,
                        "faction_source": origin_for[faction],
                    }
                )
            )
    if unresolved:
        logger.warning(
            "pvp-faction: %d pvp item(s) resolved to neither faction, dropped: %s",
            len(unresolved),
            sorted(set(unresolved)),
        )
    stats = PvpFactionStats(
        resolved_vendor=resolved_vendor,
        resolved_title=resolved_title,
        unresolved=tuple(unresolved),
    )
    return document.model_copy(update={"sources": sources}), stats
