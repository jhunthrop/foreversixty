"""A shield's own block value, which the 1.60 client does not state.

ItemSparse has no block column and `gametables/shieldblockregular.txt` is empty
on 1.60.1.70707, so a client-row shield reached simdb with only the block value
of its on-equip spells (aura 274), usually none. Wowhead Forever renders the
server's figure ("44 Block" on Evergreen Shield 279262), and its feed's
`blockamount` is the shield's total, equip bonuses included (The Immovable
Object 19321: 66, of which 25 is its on-equip spell). Every shield in the feed
carries one (314 of 314 on 1.60.1.70707). Non-shields' `blockamount` is only
their on-equip bonus, which simdb already reads from the client, so they are
left alone.

A shield only classic-db carries (no client row, no Wowhead record) takes its
`item_template.block`, the 1.12 figure: 34 such shields reached simdb blocking
nothing on 1.60.1.70707 (Skullflame Shield 1168, Crest of Retribution 13375).
Wowhead's total wins wherever both exist.
"""

from __future__ import annotations

from collections.abc import Iterable, Mapping

from pipeline import classicdb_items as cdb
from pipeline import wowhead_items as wh
from pipeline.simdb.statmap import stat_index
from pipeline.simproto import pb

ITEM_CLASS_ARMOR = 4
ITEM_SUBCLASS_SHIELD = 6
WOWHEAD_BLOCK_KEY = "blockamount"
BLOCK_VALUE_KEY = "block_value"


def shield_block_values(items: Iterable[wh.WowheadItem]) -> dict[int, float]:
    """Shield id -> its total block value, from wowhead's feed."""
    return {
        item.id: item.stats[WOWHEAD_BLOCK_KEY]
        for item in items
        if item.class_id == ITEM_CLASS_ARMOR
        and item.subclass_id == ITEM_SUBCLASS_SHIELD
        and item.stats.get(WOWHEAD_BLOCK_KEY)
    }


def classicdb_shield_block_values(items: Iterable[cdb.ClassicDbItem]) -> dict[int, float]:
    """Shield id -> its classic-db `item_template.block`."""
    return {
        item.id: float(item.block)
        for item in items
        if item.class_id == ITEM_CLASS_ARMOR
        and item.subclass_id == ITEM_SUBCLASS_SHIELD
        and item.block
    }


def _with_block(item: pb.SimItem, block: float) -> pb.SimItem:
    index = stat_index(BLOCK_VALUE_KEY)
    stats = list(item.stats) + [0.0] * max(0, index + 1 - len(item.stats))
    stats[index] = max(stats[index], block)
    out = pb.SimItem()
    out.CopyFrom(item)
    del out.stats[:]
    out.stats.extend(stats)
    return out


def with_shield_block(items: Iterable[pb.SimItem], values: Mapping[int, float]) -> list[pb.SimItem]:
    """`items`, each shield in `values` carrying at least its total block value."""
    return [_with_block(item, values[item.id]) if item.id in values else item for item in items]
