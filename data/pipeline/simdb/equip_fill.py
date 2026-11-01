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
overridden.
"""

from __future__ import annotations

from collections.abc import Iterable, Mapping

from pipeline import wowhead_items as wh
from pipeline.simdb.items import build_wowhead_sim_items
from pipeline.simdb.statmap import stat_index
from pipeline.simproto import pb

_KEEP_CLIENT = frozenset({stat_index("armor")})


def _filled(client: pb.SimItem, wowhead: pb.SimItem) -> pb.SimItem:
    size = max(len(client.stats), len(wowhead.stats))
    ours = list(client.stats) + [0.0] * (size - len(client.stats))
    theirs = list(wowhead.stats) + [0.0] * (size - len(wowhead.stats))
    stats = [
        mine if mine or index in _KEEP_CLIENT else max(0.0, other)
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
) -> list[pb.SimItem]:
    """`items`, each carrying Wowhead's stat wherever its own is zero."""
    by_id = {item.id: item for item in wowhead_items}
    out: list[pb.SimItem] = []
    for item in items:
        record = by_id.get(item.id)
        if record is None:
            out.append(item)
            continue
        (wowhead,) = build_wowhead_sim_items([record], rating_factors)
        out.append(_filled(item, wowhead))
    return out
