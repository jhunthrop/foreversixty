"""Merges wowhead's Forever gear-planner supplement into normalize_build's outputs.

The client's `ItemSparse`/`Item` stay the source of truth for every id they carry;
wowhead only ever adds the ids they lack, through the same planner gates
`pipeline.normalize.gear.build_class_items` applies to a client row (see
`pipeline.wowhead_items.supplement`). See
`docs/superpowers/specs/2026-09-27-wowhead-item-supplement-design.md` for the design.

Each `merge_*` function returns a new list/records rather than mutating its input, and
every soft gap -- an untracked stat key, a set id the client has no set for -- is
logged, never raised: a build with no wowhead payload (Classic Era, or a 1.60 build
`fetch-wowhead` has not run for yet) must normalize exactly as it did before this
module existed.
"""

from __future__ import annotations

import logging
from collections import Counter
from pathlib import Path

from pipeline.models import ClassItems, Item, ItemSetRecord
from pipeline.normalize.classes import slugify
from pipeline.wowhead_items import (
    WowheadItem,
    class_allowed,
    load_items,
    raw_path,
    supplement,
    to_gear_item,
    to_item,
)

logger = logging.getLogger(__name__)


def load_supplement(build_dir: Path, client_ids: set[int]) -> list[WowheadItem] | None:
    """The wowhead items the client's own tables lack, or None when the build carries
    no wowhead payload -- the caller then leaves every output unchanged."""
    path = raw_path(build_dir)
    if not path.exists():
        logger.info("wowhead: no payload at %s; items unchanged", path)
        return None
    return supplement(load_items(path), client_ids)


def merge_items(items: list[Item], picked: list[WowheadItem]) -> list[Item]:
    """items.json plus a flat `Item` for every picked supplement item."""
    logger.info("wowhead: merged %d supplement items into items.json", len(picked))
    return [*items, *(to_item(item) for item in picked)]


def merge_class_items(
    records: list[ClassItems],
    picked: list[WowheadItem],
    class_rows: list[dict[str, str]],
) -> list[ClassItems]:
    """Each class's items/<class-slug>.json plus the supplement gear that class may
    equip (wowhead's class mask and the client's own proficiency table)."""
    class_id_by_slug = {slugify(row["Name_lang"]): int(row["ID"]) for row in class_rows}
    untracked: Counter[str] = Counter()
    placed = 0
    merged: list[ClassItems] = []
    for record in records:
        class_id = class_id_by_slug.get(record.class_slug)
        allowed = (
            [item for item in picked if class_allowed(item, class_id)]
            if class_id is not None
            else []
        )
        placed += len(allowed)
        items = [*record.items, *(to_gear_item(item, untracked) for item in allowed)]
        merged.append(
            record.model_copy(
                update={"items": sorted(items, key=lambda i: (i.required_level, i.name, i.id))}
            )
        )
    logger.info("wowhead: merged %d supplement item placements into items/*.json", placed)
    if untracked:
        logger.info(
            "wowhead: %d supplement stat values have no planner mapping: %s",
            sum(untracked.values()),
            dict(untracked),
        )
    return merged


def merge_sets(sets: list[ItemSetRecord], picked: list[WowheadItem]) -> list[ItemSetRecord]:
    """Each supplement item whose set id names an existing set joins its `item_ids`
    (sorted, deduped). Wowhead's set payload carries no set names, so a set id the
    client has no `ItemSetRecord` for cannot be created here -- it is counted and
    skipped, a known gap (see the design spec's "Out of scope for this round")."""
    additions: dict[int, set[int]] = {record.id: set() for record in sets}
    skipped: Counter[int] = Counter()
    for item in picked:
        if item.set_id is None:
            continue
        if item.set_id not in additions:
            skipped[item.set_id] += 1
            continue
        additions[item.set_id].add(item.id)

    joined = 0
    merged: list[ItemSetRecord] = []
    for record in sets:
        new_ids = additions[record.id] - set(record.item_ids)
        if not new_ids:
            merged.append(record)
            continue
        joined += len(new_ids)
        merged.append(
            record.model_copy(update={"item_ids": sorted({*record.item_ids, *new_ids})})
        )
    logger.info("wowhead: merged %d supplement items into existing sets", joined)
    if skipped:
        logger.info(
            "wowhead: skipped %d supplement items whose set id has no client set: %s",
            sum(skipped.values()),
            dict(skipped),
        )
    return merged
