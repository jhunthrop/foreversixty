"""Whether the client's own tables confirm a classic-db-only item.

`pipeline.classicdb_items` fills the ids neither the client's `ItemSparse`
(shipped, DBCache dump or wago's hotfix layers) nor wowhead's Forever
gear-planner scrape carry. On 1.60.1.70790 that is 1,354 items, and 1,349 of
them are not missing from the client at all: the client's `Item` table has
their row (class, subclass, inventory type, icon) and its
`ItemXItemEffect`/`ItemEffect` tables carry their proc and use spells. The
server sends `ItemSparse` (name, stats, required level) for such an id on
demand, when a client first meets the item, so no export and no wowhead
scrape names it until somebody loots one (client_unconfirmed lane,
2026-11-05; design/reviews/2026-11-05-client-unconfirmed.md).

So the client does confirm that the item exists in Forever, and in what
shape. `unconfirmed_reason` is the rule that decides whether that shape is
classic-db's 1.12 item:

* the `Item` row must exist and state classic-db's class, subclass and
  inventory type. Forever turned Orb of Deception into a toy (inventory
  type 0) and nine daggers and maces from one-hand into main-hand;
* every spell the client's item effects name must be one classic-db's row
  names. Forever reworked about 80 dungeon items by giving them new spells
  (Briarwood Reed's swamp-doubled spell damage, Blackhand's Breadth's mark,
  the Royal Seals of Eldre'Thalas), so their 1.12 stats are wrong;
* every 1.12 proc or use spell must still be on the client's item: a proc
  the client dropped means the item changed.

An item that passes is confirmed: it keeps classic-db's numbers
(`stats_source: "classic-db"` still says where they came from) and the
ranker may pick it. Its numbers are the residual risk: this rule does not read
them. Of the 624 comparable items that pass it and whose Forever numbers we
do know, 113 differ from 1.12. The review file lists the macro that settles
each likely pick in game.
"""

from __future__ import annotations

import logging
from collections import Counter
from collections.abc import Iterable, Mapping
from dataclasses import dataclass

from pipeline.classicdb_items import (
    AURA_PROC_TRIGGER_SPELL,
    EFFECT_APPLY_AURA,
    TRIGGER_CHANCE_ON_HIT,
    TRIGGER_ON_USE,
    ClassicDbItem,
    ClassicDbSpell,
)
from pipeline.simdb.equip import item_effect_spells

logger = logging.getLogger(__name__)

#: `unconfirmed_reason`'s answers, one per way the client can disagree.
NOT_IN_ITEM_TABLE = "not in the client's Item table"
ITEM_SHAPE_DIFFERS = "the client's Item row states another class, subclass or inventory type"
CLIENT_EFFECTS_DIFFER = "the client's item effects name spells classic-db does not (Forever rework)"
CLIENT_LACKS_EFFECT = "a 1.12 proc or use spell is not on the client's item"


#: The `Item` columns a client row's shape is read from.
_SHAPE_COLUMNS = ("ClassID", "SubclassID", "InventoryType")


@dataclass(frozen=True)
class ClientItemRow:
    """What the client's `Item` and item-effect tables say about one id."""

    class_id: int
    subclass_id: int
    inventory_type: int
    effect_spell_ids: frozenset[int]


def client_item_rows(
    item_rows: Iterable[Mapping[str, str]],
    item_effect_rows: list[dict[str, str]],
    item_x_item_effect_rows: list[dict[str, str]],
) -> dict[int, ClientItemRow]:
    """Item id -> its `ClientItemRow`, for every row of the client's `Item`
    table (hotfix-merged, as `normalize_build` reads it). A row that states no
    shape (a removed-id hotfix stub, a test table without the columns)
    confirms nothing and is left out."""
    shapes = {
        int(row["ID"]): (int(row["ClassID"]), int(row["SubclassID"]), int(row["InventoryType"]))
        for row in item_rows
        if all(row.get(column) for column in _SHAPE_COLUMNS)
    }
    spells = item_effect_spells(
        item_effect_rows, item_x_item_effect_rows, shapes.keys(), trigger_types=None
    )
    return {
        item_id: ClientItemRow(
            class_id=class_id,
            subclass_id=subclass_id,
            inventory_type=inventory_type,
            effect_spell_ids=frozenset(s for s in spells.get(item_id, ()) if s > 0),
        )
        for item_id, (class_id, subclass_id, inventory_type) in shapes.items()
    }


def _must_be_on_client(item: ClassicDbItem, spells: Mapping[int, ClassicDbSpell]) -> set[int]:
    """The 1.12 spells a client item must still carry: every use and
    chance-on-hit spell, and every on-equip proc (Hand of Justice). A flat
    on-equip stat spell is not one: the 1.60 client states those as
    `ItemSparse` stats, not as item effects."""
    required: set[int] = set()
    for slot in item.spells:
        if not slot.spell_id:
            continue
        if slot.trigger in (TRIGGER_ON_USE, TRIGGER_CHANCE_ON_HIT):
            required.add(slot.spell_id)
            continue
        spell = spells.get(slot.spell_id)
        if spell is not None and any(
            e.effect == EFFECT_APPLY_AURA and e.aura == AURA_PROC_TRIGGER_SPELL
            for e in spell.effects
        ):
            required.add(slot.spell_id)
    return required


def unconfirmed_reason(
    item: ClassicDbItem, client: ClientItemRow | None, spells: Mapping[int, ClassicDbSpell]
) -> str | None:
    """Why the client does not confirm `item` as classic-db states it, or
    `None` when it does (module doc has the rule)."""
    if client is None:
        return NOT_IN_ITEM_TABLE
    if (client.class_id, client.subclass_id, client.inventory_type) != (
        item.class_id,
        item.subclass_id,
        item.inventory_type,
    ):
        return ITEM_SHAPE_DIFFERS
    classic_ids = {slot.spell_id for slot in item.spells if slot.spell_id}
    if not client.effect_spell_ids <= classic_ids:
        return CLIENT_EFFECTS_DIFFER
    if not _must_be_on_client(item, spells) <= client.effect_spell_ids:
        return CLIENT_LACKS_EFFECT
    return None


def confirmed_ids(
    items: Iterable[ClassicDbItem],
    client_rows: Mapping[int, ClientItemRow],
    spells: Mapping[int, ClassicDbSpell],
) -> frozenset[int]:
    """The ids among `items` the client confirms; logs the count per reason
    for the rest."""
    reasons: Counter[str] = Counter()
    confirmed: set[int] = set()
    for item in items:
        reason = unconfirmed_reason(item, client_rows.get(item.id), spells)
        if reason is None:
            confirmed.add(item.id)
        else:
            reasons[reason] += 1
    logger.info("classic-db items: the client confirms %d", len(confirmed))
    for reason, count in reasons.most_common():
        logger.info("classic-db items: %d stay client_unconfirmed: %s", count, reason)
    return frozenset(confirmed)
