"""On-equip stats for client rows whose effect links the client lacks.

A server-hotfix item's ItemSparse row reaches the client (the DBCache dump,
wago's hotfix pushes), but its ItemXItemEffect link does not: neither the
shipped table nor wago's hotfixed one carries it. Such an item kept its
ItemSparse stats and lost every on-equip spell, which is where Classic puts
spell damage, healing and most attack power (153 client-path items on
1.60.1.70707: Staff of Jordan 873, Glowing Brightwood Staff 812). An item
whose link reaches only its use or proc spell loses the same way (Hand of
Edward the Odd 2243, "+90 spell damage" on Wowhead, 0 in simdb). Wowhead
Forever's feed states the server's amounts, so each stat a client row leaves
at zero takes Wowhead's,
converted like a Wowhead-supplement item (`build_wowhead_sim_items`). Armour
stays the client's curve value, and a stat the client row states is never
overridden, and a stat a hotfix removed from the row
(`pipeline.normalize.dropped_stats`) is not put back from Wowhead's older
record.
"""

from __future__ import annotations

from collections.abc import Iterable, Mapping

from pipeline import wowhead_items as wh
from pipeline.simdb.items import build_wowhead_sim_items
from pipeline.simdb.statmap import stat_index
from pipeline.simproto import pb

_KEEP_CLIENT = frozenset({stat_index("armor")})


def _filled(client: pb.SimItem, wowhead: pb.SimItem, keep: frozenset[int]) -> pb.SimItem:
    size = max(len(client.stats), len(wowhead.stats))
    ours = list(client.stats) + [0.0] * (size - len(client.stats))
    theirs = list(wowhead.stats) + [0.0] * (size - len(wowhead.stats))
    stats = [
        mine if mine or index in keep else max(0.0, other)
        for index, (mine, other) in enumerate(zip(ours, theirs, strict=True))
    ]
    out = pb.SimItem()
    out.CopyFrom(client)
    del out.stats[:]
    out.stats.extend(stats)
    return out


def with_wowhead_equip_stats(
    items: Iterable[pb.SimItem],
    wowhead_items: Iterable[wh.WowheadItem],
    rating_factors: Mapping[str, float],
    dropped: Mapping[int, frozenset[str]] | None = None,
) -> list[pb.SimItem]:
    """`items`, each carrying Wowhead's stat wherever its own is zero, except
    the stats `dropped` says a hotfix removed from that item's row."""
    by_id = {item.id: item for item in wowhead_items}
    dropped = dropped or {}
    out: list[pb.SimItem] = []
    for item in items:
        record = by_id.get(item.id)
        if record is None:
            out.append(item)
            continue
        (wowhead,) = build_wowhead_sim_items([record], rating_factors)
        removed = frozenset(stat_index(key) for key in dropped.get(item.id, frozenset()))
        out.append(_filled(item, wowhead, _KEEP_CLIENT | removed))
    return out
