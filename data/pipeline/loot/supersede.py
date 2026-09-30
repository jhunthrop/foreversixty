"""A Forever re-itemisation whose LEGACY id the client still ships as a
real, independently-sourced row -- not just the sourceless gap
`pipeline.loot.reitemise.apply_reitemisation` fills.

Ninth wow-player sweep finding (`day3/player-review-24/casters.md`,
finding 3, 2026-09-30): every caster spec's band-60 `main_hand` lists TWO
client rows named "Grand Marshal's Stave" -- 18873 (ilvl 78, 71 spell
power, the vanilla id) and 234571 (ilvl 80, 81 spell power, +28 crit
rating, Forever's re-itemised copy) -- and likewise "High Warlord's War
Staff" 18874/234549. `pvp:rank-18:alliance` lists both ids for the first
pair. Both rows are real client rows (`stats_source` client,
`reitemised_from` unset on both), so `apply_reitemisation` never touched
them -- that step only ever fills a new id `build_loot`'s own passes left
WHOLLY unsourced (`apply_reitemisation`'s own doc), and here the new id
already has its own real source.

Verified against two primary sources (tenet 8): `pipeline.loot.reitemise.
reitemised_pairs` itself (the pinned direction is always
new-id-owns-a-Classic-twin's name+slot, gated the same way
`apply_reitemisation` gates it -- ilvl gap and armor/weapon subclass),
and wowhead's live Forever `npc=12782`/`npc=14581` "sells" listviews
(Captain O'Neal, Sergeant Thunderhorn), fetched 2026-09-30: the LEGACY id
carries `envChange.status: "unchanged"` (or a change unrelated to the
re-itemisation, High Warlord's own damage-curve note) while the NEW id
carries `"effects_changed"`/`"stats_changed"` naming the exact stat this
build's own `items/*.json` already shows it gained ("Critical Strike
Rating 28 added") -- the new id is Forever's own hotfixed row, the
legacy id is the vanilla leftover the client has not stopped shipping.

`superseded_pairs` is therefore narrower than `reitemised_pairs`: only a
pair BOTH of whose ids `build_loot` already, independently, named in the
SAME source bucket -- proof the two are literally duplicate listings for
the one reward, not merely two catalogue rows close enough in item level
and subclass to be plausibly-the-same-content (`reitemised_pairs`'s own,
looser bar). `apply_supersession` drops the legacy id from every such
bucket; `mark_superseded_items` writes `superseded_by` (legacy id ->
new id) onto the flat and per-class rows so a legacy row that is no
longer offered anywhere still shows up as itself in a tooltip for
whatever already-owned character equips one -- items.json's own
no-shrink gate (`pipeline.csvio.check_item_sparse_completeness`) counts
it exactly like any other row; nothing here removes it from the file.

The ranker (`sim/cmd/leveling-bis`, Go) is the intended reader of
`superseded_by`: "a row with `superseded_by` set is never a pick
candidate" is a one-line change on that side, out of scope for this
lane -- see the lane report for the exact JSON key.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.loot.reitemise import reitemised_pairs
from pipeline.loot.sources import source_item_ids
from pipeline.models import ClassItems, Item, LootBoss, LootFile, LootSource
from pipeline.normalize import write_json, write_model


def superseded_pairs(document: LootFile, item_rows: list[dict]) -> dict[int, int]:
    """Legacy (classic) item id -> the Forever-new id that supersedes it,
    for every `reitemised_pairs()` entry where BOTH ids are already named
    in the exact same source bucket -- `document`'s own `sources`, via
    `source_item_ids` (flat `items`/`trash`, or a boss's own `items`)."""
    pairs = reitemised_pairs(item_rows)
    if not pairs:
        return {}
    superseded: dict[int, int] = {}
    for new_id, classic_id in pairs.items():
        for source in document.sources:
            ids = source_item_ids(source)
            if new_id in ids and classic_id in ids:
                superseded[classic_id] = new_id
                break
    return superseded


def _drop_legacy(ids: list[int], superseded: dict[int, int]) -> tuple[list[int], int]:
    """`ids` with every legacy id dropped wherever ITS OWN replacement is
    also present in `ids` -- never dropped alone, since `superseded_pairs`
    only proposes a pair both ids already co-occur in, but a single bucket
    is checked again here rather than trusted blindly (a curated overlay
    or a later pass could, in principle, narrow one side after this ran)."""
    present = set(ids)
    kept: list[int] = []
    removed = 0
    for item_id in ids:
        new_id = superseded.get(item_id)
        if new_id is not None and new_id in present:
            removed += 1
            continue
        kept.append(item_id)
    return kept, removed


def _dedup_boss(boss: LootBoss, superseded: dict[int, int]) -> tuple[LootBoss, int]:
    kept, removed = _drop_legacy(boss.items, superseded)
    if removed == 0:
        return boss, 0
    return boss.model_copy(update={"items": kept}), removed


def _dedup_source(source: LootSource, superseded: dict[int, int]) -> tuple[LootSource, int]:
    update: dict = {}
    removed = 0
    if source.items:
        kept, n = _drop_legacy(source.items, superseded)
        if n:
            update["items"] = kept
            removed += n
    if source.trash:
        kept, n = _drop_legacy(source.trash, superseded)
        if n:
            update["trash"] = kept
            removed += n
    if source.bosses:
        new_bosses = []
        changed = False
        for boss in source.bosses:
            new_boss, n = _dedup_boss(boss, superseded)
            new_bosses.append(new_boss)
            if n:
                changed = True
                removed += n
        if changed:
            update["bosses"] = new_bosses
    if not update:
        return source, 0
    return source.model_copy(update=update), removed


def apply_supersession(
    document: LootFile, superseded: dict[int, int]
) -> tuple[LootFile, int]:
    """`document` with every legacy id in `superseded` (classic id -> new
    id) dropped from any source bucket that lists both it and its
    replacement. Returns the new document and how many (bucket, id)
    removals it made, for the lane's own before/after report."""
    if not superseded:
        return document, 0
    sources = []
    removed_total = 0
    for source in document.sources:
        new_source, removed = _dedup_source(source, superseded)
        sources.append(new_source)
        removed_total += removed
    return document.model_copy(update={"sources": sources}), removed_total


def mark_superseded_items(build_dir: Path, superseded: dict[int, int]) -> int:
    """Read-modify-write `items.json` and every `items/<class>.json`,
    setting `superseded_by` on each legacy id's row -- same discipline
    `pipeline.loot.gear.apply_fork_columns` uses for its own two columns,
    so a build with nothing superseded rewrites identical bytes. Returns
    how many flat-catalogue (`items.json`) rows carry the field, which is
    the lane-report count (a per-class row is the same id repeated per
    class that can use it, not a second fact). A no-op (no files touched)
    when `superseded` is empty."""
    if not superseded:
        return 0
    items_path = build_dir / "items.json"
    rows = json.loads(items_path.read_text(encoding="utf-8"))
    items = [
        Item(**row).model_copy(update={"superseded_by": superseded.get(int(row["id"]))})
        for row in rows
    ]
    write_json(items, items_path)
    marked = sum(1 for item in items if item.superseded_by is not None)

    items_dir = build_dir / "items"
    for path in sorted(items_dir.glob("*.json")):
        record = ClassItems(**json.loads(path.read_text(encoding="utf-8")))
        updated = [
            gear_item.model_copy(update={"superseded_by": superseded.get(gear_item.id)})
            for gear_item in record.items
        ]
        write_model(record.model_copy(update={"items": updated}), path)
    return marked
