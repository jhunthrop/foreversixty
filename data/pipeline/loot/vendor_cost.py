"""vendor-11036 lane, 2026-09-30: a wowhead-scraped `vendor:<npc_id>`
`LootSource` whose item is not actually buyable at launch, split out of
`pipeline.loot.sources` (already pushing 1300 lines) the same way
`pipeline.loot.wowhead`/`pipeline.loot.supersede` were before it.

Twelfth wow-player sweep finding: vendor 11036 ("Leonid Barthalomew the
Revered", Light's Hope Chapel) and its twin 240248 ("Bryon Steelblade")
sell 192 items each -- full quality-4 Forever sets 18-38 item levels
above rank-18 PvP epics -- with no gold price and no reputation
requirement wowhead's own `dropped-by`/`sold-by` listview states
directly. Primary-source check (wowhead's live item pages,
2026-09-30, item 240023 "Inquisition Vambraces" among them): every one
of the 208 rows wowhead's own `sells` listview names for 11036 DOES
carry an `ItemExtendedCost`-shaped `cost` -- never gold, always a
single item id as the price -- but every one of those 40 distinct price
items is itself either (a) not a real item at all (239759 and 23 others:
absent from this build's own `items.json` AND a 404 on wowhead's own
`item=<id>` page -- a placeholder the server-side data table names but
nothing ever implemented), or (b) a REAL catalogued item (16 of them,
the "Inquisition" plate set among them) whose own ONLY source, anywhere,
is these same two vendors -- a closed loop: buying anything requires
already owning an item that can only be bought by buying something
first. Neither is a functioning launch-day economy; this is beta/test
vendor plumbing Blizzard wired with placeholder costs before whatever
real acquisition method (most likely a seasonal token or a raid-tied
currency, given the set's own item level) ships.

`apply_vendor_cost_gate` encodes this as a rule, not a hand list of
these two npc ids, so it also catches a future vendor of the same
shape:

1. An item id named as an `ItemExtendedCost` price that is not in this
   build's own catalogue at all is unresolvable currency outright.
2. A GREATEST fixed point of "obtainable" (mirrors `pipeline.loot.
   sources.apply_quest_opens_gate`'s own item-reachability loop, run
   the other way round): an item starts obtainable only when it has an
   independent, non-vendor source, or a `sold-by` row needing no
   `ItemExtendedCost` at all; the closure then grows that set to cover
   an item payable using only ALREADY-obtainable ingredients. Anything
   left ungrounded when the closure settles is unresolvable -- which is
   what catches the Inquisition-set circular loop (A prices in B, B
   prices in A, neither has any other source) that a one-directional
   walk from a known-bad seed would miss entirely, without naming
   either vendor or either item.
3. A `vendor:<npc_id>` item every one of whose rows resolves this way
   is moved to a `vendor:<npc_id>:later` sibling (`opens="later"`),
   same split shape `apply_crafted_opens_gate` already uses for a
   `crafted:<profession>:<phase>` sibling -- the ranker needs no Go
   change to read it.
4. A `sold-by` row with NO `ItemExtendedCost` at all (an ordinary gold
   buy wowhead states no override for) and no reputation requirement is
   left alone UNLESS the item is quality 4+ and above this build's own
   launch-phase item level ceiling (the highest item level any
   currently-ungated source in `document` already names -- dungeons,
   quests, PvP ranks, ordinary vendors: whatever a player can actually
   reach before the first raid tier opens) -- the same "reads like beta
   inventory" heuristic the lane's own brief describes, applied without
   a hand list of ids either.

Must run after `pipeline.loot.sources.apply_overlays`/
`apply_quest_opens_gate`/`apply_crafted_opens_gate`: the launch-phase
ceiling in point 4 is read off THEIR finished `opens` gates, the same
"needs the document's other gates already resolved" ordering rule every
gate in this pipeline already follows.
"""

from __future__ import annotations

from collections import defaultdict

from pipeline.item_sources import ItemSourceEntry
from pipeline.loot.sources import KIND_ORDER, source_item_ids
from pipeline.models import LootFile, LootSource

#: Bound on the fixed-point loop in `_obtainable_items` -- same role and
#: same headroom as `apply_quest_opens_gate`'s own
#: `_MAX_QUEST_GATE_PASSES`.
_MAX_PASSES = 5

#: `LootSource.source_note` text for each of this module's two gates --
#: pipeline bookkeeping, never shown in player-facing copy (tenet 7).
_PHANTOM_COST_NOTE = (
    "wowhead prices this row in an ItemExtendedCost item id that is not itself a real, "
    "separately obtainable item at launch (missing from this build's own catalogue, or "
    "obtainable only from this same vendor) -- beta/test data, not a real purchase."
)
_OVERLEVELED_NOTE = (
    "wowhead states no gold price and no reputation requirement for this row at all, and "
    "the item's own quality/level sit above anything else this build sources as open at "
    "launch -- beta/test inventory, not a real purchase."
)


def _row_payable(row, obtainable: set[int]) -> bool:
    """A `sold-by` row is payable when it states no `ItemExtendedCost`
    at all (`cost_item_ids` empty -- an ordinary gold buy, whatever
    `cost_money` is) OR every one of its cost items is already known
    obtainable -- `ItemExtendedCost` requires handing over every listed
    item, never a choice of one, so a single bad ingredient spoils the
    whole row."""
    return not row.cost_item_ids or all(c in obtainable for c in row.cost_item_ids)


def _obtainable_items(
    item_sources: dict[int, ItemSourceEntry], build_items: set[int]
) -> set[int]:
    """The GREATEST fixed point of "obtainable" over `build_items`:
    grounded facts first (point 1's mirror -- a real catalogued item
    with an independent, non-vendor way to get it, OR one this scrape
    never even names, an ordinary fork/classic-db item outside its
    scope), then point 2's closure, growing the set to cover an item
    payable using only ALREADY-obtainable ingredients. An item id NOT in
    `build_items` at all can never enter this set (the phantom-currency
    case: not itself a real item, so never grounded and never reachable
    through it), and neither can one whose every `sold-by` row needs an
    ingredient that never becomes obtainable either -- a circular pair
    (A costs B, B costs A, neither has any other source) grounds
    NEITHER side, unlike a one-directional seed-and-grow walk from a
    known-bad set would miss.
    """
    obtainable: set[int] = set()
    for item_id in build_items:
        entry = item_sources.get(item_id)
        if entry is None or entry.dropped_by or entry.crafted_by or entry.quest_rewards:
            obtainable.add(item_id)
    for _ in range(_MAX_PASSES):
        changed = False
        for item_id, entry in item_sources.items():
            if item_id in obtainable or item_id not in build_items:
                continue
            if any(_row_payable(row, obtainable) for row in entry.sold_by):
                obtainable.add(item_id)
                changed = True
        if not changed:
            break
    return obtainable


def _launch_phase_item_level_cap(document: LootFile, item_levels: dict[int, int]) -> int:
    """The highest item level any item in a currently-UNGATED,
    NON-VENDOR source (`opens` unset -- overlays/quest-gate/crafted-gate
    have all already run by the time this is called) carries: the
    practical ceiling on what a player can reach before the first raid
    tier opens, read off the document itself rather than a hand-picked
    number. `vendor`-kind sources are excluded from the measurement
    itself -- this cap is what a vendor row's OWN item level is compared
    against, so including vendor sources would let an ungated beta
    vendor's own overleveled stock inflate the very ceiling meant to
    catch it."""
    levels = [
        item_levels[item_id]
        for source in document.sources
        if not source.opens and source.kind != "vendor"
        for item_id in source_item_ids(source)
        if item_id in item_levels
    ]
    return max(levels, default=0)


def apply_vendor_cost_gate(
    document: LootFile,
    item_sources: dict[int, ItemSourceEntry],
    build_items: set[int],
    item_levels: dict[int, int],
    item_qualities: dict[int, int],
) -> tuple[LootFile, int]:
    """Splits every wowhead-scraped `vendor:<npc_id>` source's items into
    their ungated remainder (kept on the original id) and a
    `vendor:<npc_id>:later` sibling for the items this module's doc
    gates -- see that doc for the full rule. Returns the new document and
    how many (npc, item) rows moved, the lane's own before/after count.
    """
    obtainable = _obtainable_items(item_sources, build_items)
    cap = _launch_phase_item_level_cap(document, item_levels)
    scraped_npc_ids = {row.npc_id for entry in item_sources.values() for row in entry.sold_by}
    new_sources: list[LootSource] = []
    gated_items: dict[int, set[int]] = defaultdict(set)
    gated_notes: dict[int, set[str]] = defaultdict(set)
    moved = 0
    for source in document.sources:
        npc_id = source.npc_id
        if source.kind != "vendor" or not source.items or npc_id not in scraped_npc_ids:
            new_sources.append(source)
            continue
        ungated_items: list[int] = []
        for item_id in source.items:
            entry = item_sources.get(item_id)
            row = next((r for r in (entry.sold_by if entry else []) if r.npc_id == npc_id), None)
            note = None
            if row is not None and not _row_payable(row, obtainable):
                note = _PHANTOM_COST_NOTE
            elif (
                row is not None
                and not row.cost_item_ids
                and not row.cost_money
                and (entry is None or entry.required_faction_id is None)
                and item_qualities.get(item_id, 0) >= 4
                and item_levels.get(item_id, 0) > cap
            ):
                note = _OVERLEVELED_NOTE
            if note:
                gated_items[npc_id].add(item_id)
                gated_notes[npc_id].add(note)
                moved += 1
            else:
                ungated_items.append(item_id)
        if ungated_items:
            new_sources.append(source.model_copy(update={"items": sorted(ungated_items)}))
    if not moved:
        return document, 0
    for npc_id, item_ids in sorted(gated_items.items()):
        base = next(s for s in document.sources if s.kind == "vendor" and s.npc_id == npc_id)
        new_sources.append(
            base.model_copy(
                update={
                    "id": f"{base.id}:later",
                    "items": sorted(item_ids),
                    "opens": "later",
                    "source_note": "; ".join(sorted(gated_notes[npc_id])),
                    "faction_id": None,
                    "standing": None,
                }
            )
        )
    new_sources.sort(key=lambda source: (KIND_ORDER.index(source.kind), source.id))
    return document.model_copy(update={"sources": new_sources}), moved
