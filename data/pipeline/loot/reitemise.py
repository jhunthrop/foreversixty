"""Re-itemisation inheritance (src-classicdb lane item 3, 2026-09-29;
gated by item level and subclass, src-classicdb-fixes lane, 2026-09-29):
a Forever-NEW item (id >= 200000) that shares its exact name and
inventory slot with an existing Classic-era item OFTEN drops, sells or
rewards from the SAME place the Classic item does -- Forever re-itemises
old content rather than inventing new drop tables for it. But name+slot
alone is not enough: Forever also reuses old item NAMES for unrelated
new content (Swamp Ring 270052, ilvl 35, no required level, is a
low-level leveling ring that merely reuses the name "Swamp Ring" from
the ilvl-57 Scholomance quest reward 12015 -- inheriting 12015's sources
put a Darkmaster Gandling drop atop the level-20 hunter list). So a
name+slot candidate is only trusted when the two items are plausibly the
SAME piece of content re-tuned: item levels within
`ITEM_LEVEL_TOLERANCE` of each other, and the same armor/weapon
subclass (`class_id` AND `subclass_id`, since subclass ids are only
unique within a class -- weapon subclass 0 is One-Handed Axe, armor
subclass 0 is Miscellaneous, the very subclass both Swamp Rings share).
A candidate that fails the gate is logged and left unsourced rather than
guessed at.

450 (name, inventory_type) pairs exist on build 1.60.1.70009; of those,
`apply_reitemisation` only ever ACTS on a new id that `build_loot`'s own
fork/classic-db/wowhead passes left wholly unsourced (most of the 450
already have their own real source and need nothing copied) and whose
name+slot key names EXACTLY ONE classic candidate (a PvP rank insignia's
per-rank ids -- eight different classic ids all named "Insignia of the
Alliance" -- is left alone rather than guessed at: never the fork's or
any scrape's own "invent nothing" policy this pipeline holds to
elsewhere).

Direction is one-way, always: a classic item's own sources are copied
onto the new item, never the reverse (`reitemised_pairs`' own
`newer_id -> classic_id` shape makes this true by construction). Every
bucket the copy lands in keeps its OWN existing `source_origin` --
copying does not create a new "reitemised" kind of place, it says "this
item drops from the exact place classic_id already does", so the
breadcrumb is per item, not per source: `LootSource.reitemised_from`/
`LootBoss.reitemised_from` (item id string -> the classic item id it was
copied from), the same per-item dict shape `item_chances` already uses.
"""

from __future__ import annotations

import logging
from collections import defaultdict

from pipeline.models import LootBoss, LootFile, LootSource

logger = logging.getLogger(__name__)

#: Forever's own item id split: 200000+ is a newly minted item id, never
#: a real Blizzard/Classic one (`pipeline.normalize.items`' own
#: convention, unchanged since the itemisation design landed).
NEW_ITEM_ID_FLOOR = 200000

#: How far a new item's `item_level` may drift from the classic item it
#: would inherit sources from before the two are treated as unrelated
#: content that merely share a name (see the module docstring's Swamp
#: Ring example, 22 ilvl apart and correctly rejected by this gate).
ITEM_LEVEL_TOLERANCE = 10

#: `items.json`'s own `faction_restriction` column spellings -> the
#: `QuestSource.faction` vocabulary -- same mapping
#: `test_loot_build.py`'s `by_restriction` checks against, restated here
#: because `apply_reitemisation` needs it at write time, not just at
#: test time.
_RESTRICTION_TO_FACTION = {"": "both", "alliance_only": "alliance", "horde_only": "horde"}


def _plausibly_the_same_item(new_row: dict, classic_row: dict) -> bool:
    """True when `new_row` and `classic_row` are close enough in power
    and category to be the same piece of content, re-tuned -- item
    levels within `ITEM_LEVEL_TOLERANCE`, and the same armor/weapon
    subclass (`class_id` and `subclass_id` both)."""
    ilvl_gap = abs(int(new_row["item_level"]) - int(classic_row["item_level"]))
    same_subclass = int(new_row["class_id"]) == int(classic_row["class_id"]) and int(
        new_row["subclass_id"]
    ) == int(classic_row["subclass_id"])
    return ilvl_gap <= ITEM_LEVEL_TOLERANCE and same_subclass


def reitemised_pairs(item_rows: list[dict]) -> dict[int, int]:
    """Forever-new item id -> the single Classic item id sharing its
    exact `(name, inventory_type)`, for every such pair that is both
    UNAMBIGUOUS (exactly one classic candidate) and plausibly the same
    item (`_plausibly_the_same_item`). `item_rows` is a build's
    `items.json` (or any list of dicts with `id`/`name`/
    `inventory_type`/`item_level`/`class_id`/`subclass_id`).
    """
    classic_by_key: dict[tuple[str, int], list[dict]] = defaultdict(list)
    for row in item_rows:
        item_id = int(row["id"])
        if item_id < NEW_ITEM_ID_FLOOR:
            classic_by_key[(row["name"], int(row["inventory_type"]))].append(row)

    pairs: dict[int, int] = {}
    for row in item_rows:
        item_id = int(row["id"])
        if item_id < NEW_ITEM_ID_FLOOR:
            continue
        candidates = classic_by_key.get((row["name"], int(row["inventory_type"])))
        if not candidates or len(candidates) != 1:
            continue
        classic_row = candidates[0]
        if not _plausibly_the_same_item(row, classic_row):
            logger.info(
                "reitemisation: %s (id %d, ilvl %d) not inherited from %s "
                "(id %d, ilvl %d): item level gap or armor/weapon subclass "
                "mismatch",
                row["name"],
                item_id,
                int(row["item_level"]),
                classic_row["name"],
                int(classic_row["id"]),
                int(classic_row["item_level"]),
            )
            continue
        pairs[item_id] = int(classic_row["id"])
    return pairs


def _copy_boss(boss: LootBoss, new_id: int, classic_id: int) -> LootBoss:
    return boss.model_copy(
        update={
            "items": sorted({*boss.items, new_id}),
            "reitemised_from": {**(boss.reitemised_from or {}), str(new_id): classic_id},
        }
    )


def _copy_into_source(source: LootSource, new_id: int, classic_id: int) -> LootSource:
    """`source` with `new_id` added wherever `classic_id` already
    appears in it (its flat `items`, `trash`, or a boss's own `items`),
    tagging each touched bucket's `reitemised_from`. A source that does
    not name `classic_id` at all is returned unchanged."""
    update: dict = {}
    if source.items and classic_id in source.items:
        update["items"] = sorted({*source.items, new_id})
        update["reitemised_from"] = {**(source.reitemised_from or {}), str(new_id): classic_id}
    if source.trash and classic_id in source.trash:
        update["trash"] = sorted({*source.trash, new_id})
        update["reitemised_from"] = {**(source.reitemised_from or {}), str(new_id): classic_id}
    if source.bosses:
        new_bosses = [
            _copy_boss(boss, new_id, classic_id) if classic_id in boss.items else boss
            for boss in source.bosses
        ]
        if new_bosses != source.bosses:
            update["bosses"] = new_bosses
    return source.model_copy(update=update) if update else source


def apply_reitemisation(document: LootFile, item_rows: list[dict]) -> tuple[LootFile, int]:
    """`document` with every unsourced Forever-new item that
    `reitemised_pairs` can resolve added alongside its classic
    counterpart, everywhere that counterpart already appears (including
    the flat `quest` bucket and `document.quests`' own per-item detail,
    copied verbatim -- the reward IS the same quest). Returns the new
    document and how many item ids were actually filled in (for the
    lane's own before/after report).
    """
    from pipeline.loot.sources import source_item_ids  # local: avoid a circular import

    already_named = {item_id for source in document.sources for item_id in source_item_ids(source)}
    pairs = {
        new_id: classic_id
        for new_id, classic_id in reitemised_pairs(item_rows).items()
        if new_id not in already_named and classic_id in already_named
    }
    if not pairs:
        return document, 0

    sources = list(document.sources)
    filled: set[int] = set()
    for new_id, classic_id in pairs.items():
        touched = False
        for index, source in enumerate(sources):
            updated = _copy_into_source(source, new_id, classic_id)
            if updated is not source:
                sources[index] = updated
                touched = True
        if touched:
            filled.add(new_id)

    # A quest's faction detail is the ITEM's own truth, not the quest's
    # (models.QuestSource's own doc, and the same fix
    # pipeline.loot.classicdb/wowhead both apply for the identical
    # reason) -- the new item can carry a DIFFERENT factionRestriction
    # than the classic one it inherits sources from (measured while
    # regenerating loot.json for 1.60.1.70009: item 226969 inherits
    # quest detail from an alliance-restricted classic item while being
    # unrestricted itself), so a copied `QuestSource` gets its own
    # `faction` re-derived from the NEW item, never copied verbatim.
    new_item_faction = {
        int(row["id"]): _RESTRICTION_TO_FACTION.get(row.get("faction_restriction") or "", "both")
        for row in item_rows
        if int(row["id"]) in filled
    }
    quests = dict(document.quests)
    for new_id in filled:
        classic_id = pairs[new_id]
        detail = quests.get(str(classic_id))
        if detail is not None:
            faction = new_item_faction.get(new_id, "both")
            quests[str(new_id)] = [
                entry.model_copy(update={"faction": faction}) for entry in detail
            ]

    return document.model_copy(update={"sources": sources, "quests": quests}), len(filled)
