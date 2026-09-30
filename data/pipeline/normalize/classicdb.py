"""Merges the classic-db item supplement into `normalize_build`'s outputs.

Runs strictly after `pipeline.normalize.wowhead`'s own merge: the client's
`ItemSparse`/`Item` are the source of truth for every id they carry, wowhead's
Forever gear-planner scrape is next (it is corroborated against this same
client's curve tables on every id both share -- `pipeline.wowhead_items`'s own
doc), and classic-db's 1.12 `item_template` only ever fills the ids BOTH of
those lack (`pipeline.classicdb_items.supplement`'s own doc). See
`pipeline.normalize.wowhead`'s module doc for why each `merge_*` function
returns a new list/records rather than mutating its input, and why a soft gap
(an untracked stat, a set id with no client set) is logged, never raised.
"""

from __future__ import annotations

import logging
from collections import Counter
from pathlib import Path

from pipeline.classicdb_items import (
    ClassicDbItem,
    ClassicDbSpell,
    class_allowed,
    load_extract,
    supplement,
    to_gear_item,
    to_item,
)
from pipeline.models import ClassItems, Item, ItemSetRecord
from pipeline.normalize.classes import slugify
from pipeline.spelltext import SpellText

logger = logging.getLogger(__name__)


def load_classicdb_supplement(
    build_dir: Path, known_ids: set[int]
) -> tuple[list[ClassicDbItem], dict[int, ClassicDbSpell]] | None:
    """The classic-db items neither the client nor wowhead's own supplement
    already cover, plus every spell those items' `spellid_<n>` reference
    (classicdb-fidelity lane, 2026-09-30: `merge_class_items` needs the
    structured spell data, not only the item rows) -- or `None` when the
    build carries no committed extract, the caller then leaves every output
    unchanged."""
    extract = load_extract(build_dir)
    if extract is None:
        return None
    records, spells = extract
    return supplement(records, known_ids), spells


def merge_items(items: list[Item], picked: list[ClassicDbItem]) -> list[Item]:
    """`items` plus a flat `Item` for every picked supplement item -- REPLACING,
    not duplicating, any existing row with the same id.

    `known_ids` (`load_classicdb_supplement`'s own caller in
    `pipeline.normalize`) only counts the client as already covering an id
    when its own row is real equippable data (`InventoryType != 0`), so a
    client row that names an id but states nothing usable (Orb of Deception,
    1973, on build 1.60.1.70009 -- see that caller's own doc) can still be
    picked here. Its flat `items.json` row already exists (with the client's
    broken `inventory_type: 0`); appending a second one would leave two rows
    for the same id, so the stale one is dropped in favour of classic-db's
    real data.
    """
    picked_ids = {item.id for item in picked}
    kept = [item for item in items if item.id not in picked_ids]
    logger.info("classic-db items: merged %d supplement items into items.json", len(picked))
    return [*kept, *(to_item(item) for item in picked)]


def merge_class_items(
    records: list[ClassItems],
    picked: list[ClassicDbItem],
    class_rows: list[dict[str, str]],
    spell_text: SpellText,
    fork_icons: dict[int, str],
    wowhead_icons: dict[int, str],
    spells: dict[int, ClassicDbSpell] | None = None,
    client_spell_names: dict[int, str] | None = None,
) -> list[ClassItems]:
    """Each class's `items/<class-slug>.json` plus the classic-db supplement
    gear that class may equip (classic-db's own `AllowableClass` mask and the
    client's own proficiency table).

    `spells`/`client_spell_names` feed `to_gear_item`'s structured
    `stats`/`effect_text` resolution (classicdb-fidelity lane, 2026-09-30);
    pass `None` (the default) for a caller with neither and every supplement
    row falls all the way back to its bare classic-db spell name, exactly as
    before that lane existed.
    """
    spells = spells or {}
    class_id_by_slug = {slugify(row["Name_lang"]): int(row["ID"]) for row in class_rows}
    placed = 0
    icon_origins: Counter[str] = Counter()
    merged: list[ClassItems] = []
    for record in records:
        class_id = class_id_by_slug.get(record.class_slug)
        allowed = (
            [item for item in picked if class_allowed(item, class_id)]
            if class_id is not None
            else []
        )
        placed += len(allowed)
        new_items = [
            to_gear_item(item, spell_text, spells, fork_icons, wowhead_icons, client_spell_names)
            for item in allowed
        ]
        for gear_item in new_items:
            if gear_item.icon_source:
                icon_origins[gear_item.icon_source] += 1
        items = [*record.items, *new_items]
        merged.append(
            record.model_copy(
                update={"items": sorted(items, key=lambda i: (i.required_level, i.name, i.id))}
            )
        )
    logger.info("classic-db items: merged %d supplement item placements into items/*.json", placed)
    if icon_origins:
        logger.info(
            "classic-db items icon origins: %d from the fork db, %d from wowhead's "
            "gear-planner payload, %d still on the placeholder",
            icon_origins["fork"],
            icon_origins["wowhead"],
            placed - icon_origins["fork"] - icon_origins["wowhead"],
        )
    return merged


def merge_sets(sets: list[ItemSetRecord], picked: list[ClassicDbItem]) -> list[ItemSetRecord]:
    """Each supplement item whose set id names an existing set joins its
    `item_ids` (sorted, deduped) -- same contract as
    `pipeline.normalize.wowhead.merge_sets`: an id with no matching
    `ItemSetRecord` is counted and skipped, not invented."""
    additions: dict[int, set[int]] = {record.id: set() for record in sets}
    skipped = 0
    for item in picked:
        if item.set_id is None:
            continue
        if item.set_id not in additions:
            skipped += 1
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
        merged.append(record.model_copy(update={"item_ids": sorted({*record.item_ids, *new_ids})}))
    logger.info("classic-db items: merged %d supplement items into existing sets", joined)
    if skipped:
        logger.info(
            "classic-db items: skipped %d supplement items whose set id has no client set",
            skipped,
        )
    return merged
