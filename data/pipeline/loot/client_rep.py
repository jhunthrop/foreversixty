"""Reputation sources from an item's own client requirement.

The fork's database names the reputation rewards of classic factions. Forever's
new ones live only in the client: ItemSparse.MinFactionID and MinReputation,
which normalize carries into items/<class>.json as `reputation`, together with
the faction's name and side (`pipeline.normalize.factions`). Darkspear Islands
is the case at launch. Creeg Bothunk sells the Darkspear Raiders set, and
Wowhead names him. Nothing named a seller of the Theramore Expeditionary Force
set, so the Alliance half had no source at all.

`apply_client_reputations` gives every such item that no source names a
`rep:<faction>:<standing>` source, the same shape `pipeline.loot.sources`
builds from the fork. It also records the faction's side in the document's
`factions`, so each set stays on its own side in loot and in the ranker. An
item a fork `rep` source already names keeps that source. Neutral gates
nothing, so a neutral requirement adds no source.

An item another source or a quest already names (a vendor's stock, Creeg
Bothunk's Darkspear Raiders set) gets its side and no rep source (round 10).
That source is where the item comes from, and the ranker holds every source
to the item's own client standing (`sim/cmd/leveling-bis` reputation.go). A
rep source beside it was a second way in that no gate on the vendor closes:
the gated Undermine traders' Booty Bay, Ratchet, Gadgetzan and Everlook
stock opened through it.
"""

from __future__ import annotations

import json
from collections import defaultdict
from collections.abc import Collection, Mapping
from pathlib import Path

from pipeline.loot.sources import KIND_ORDER
from pipeline.models import ItemReputation, LootFile, LootSource
from pipeline.normalize.classes import slugify

REP_KIND = "rep"
UNGATED_STANDING = "neutral"


def load_item_reputations(build_dir: Path) -> dict[int, ItemReputation]:
    """Item id -> its client reputation requirement, from every items/<class>.json."""
    out: dict[int, ItemReputation] = {}
    for path in sorted((build_dir / "items").glob("*.json")):
        for row in json.loads(path.read_text(encoding="utf-8"))["items"]:
            if row.get("reputation"):
                out[int(row["id"])] = ItemReputation.model_validate(row["reputation"])
    return out


def _rep_source_id(reputation: ItemReputation) -> str:
    return f"{REP_KIND}:{slugify(reputation.faction_name or '')}:{reputation.standing}"


def _named_by_another_way(document: LootFile) -> set[int]:
    """Every item a non-rep source (items, trash or a boss) or a quest names."""
    named: set[int] = set()
    for source in document.sources:
        if source.kind == REP_KIND:
            continue
        named.update(source.items or [])
        named.update(source.trash or [])
        for boss in source.bosses or []:
            named.update(boss.items)
    named.update(int(item_id) for item_id in document.quests)
    return named


def apply_client_reputations(
    document: LootFile,
    reputations: Mapping[int, ItemReputation],
    build_items: Collection[int],
) -> tuple[LootFile, int]:
    """`document` with a `rep` source and a side for every client reputation
    reward no `rep` source names, and how many items it sourced."""
    rep_named = {
        item_id
        for source in document.sources
        if source.kind == REP_KIND
        for item_id in source.items or []
    }
    sourced_elsewhere = _named_by_another_way(document)
    added: dict[str, set[int]] = defaultdict(set)
    by_source: dict[str, ItemReputation] = {}
    factions = dict(document.factions)
    for item_id, reputation in sorted(reputations.items()):
        if (
            item_id not in build_items
            or item_id in rep_named
            or reputation.faction_name is None
            or reputation.standing == UNGATED_STANDING
        ):
            continue
        if reputation.side is not None:
            factions.setdefault(str(item_id), reputation.side)
        if item_id in sourced_elsewhere:
            continue
        source_id = _rep_source_id(reputation)
        added[source_id].add(item_id)
        by_source[source_id] = reputation
    sourced = sum(len(items) for items in added.values())
    sources = []
    for source in document.sources:
        extra = added.pop(source.id, set())
        sources.append(
            source.model_copy(update={"items": sorted({*(source.items or []), *extra})})
            if extra
            else source
        )
    sources += [
        LootSource(
            id=source_id,
            kind=REP_KIND,
            name=by_source[source_id].faction_name or "",
            faction_id=by_source[source_id].faction_id,
            standing=by_source[source_id].standing,
            items=sorted(items),
        )
        for source_id, items in added.items()
    ]
    sources.sort(key=lambda source: (KIND_ORDER.index(source.kind), source.id))
    return document.model_copy(update={"sources": sources, "factions": factions}), sourced
