"""A quest-only item's per-class catalogue row, narrowed to the classes
its own quest(s) can actually accept.

Eleventh wow-player sweep finding, day3 data-followups-7 lane,
2026-09-30: Fire Ruby (item 20036) is published as a trinket pick or
alternative for hunter-beast-mastery, shaman-elemental and paladin-
retribution, but its ONLY source is quest 8253 "Destroy Morphaz" --
`quest_template.RequiredClasses` names mages alone (`pipeline.
classic_sources._classes_from_required_classes`, `ClassicDbQuestInfo.
classes`, threaded onto `pipeline.models.QuestSource.classes` by
`pipeline.loot.sources.build_loot`), so no other class can ever accept
the quest at all, whatever the reward item's own `AllowableClass` says
(Fire Ruby carries none -- `AllowableClass` -1, "any class", same as
every other quest-only reward that names no class restriction of its
own; the QUEST is the only place this fact lives).

`pipeline.normalize.gear.build_class_items` has no notion of a quest at
all -- it builds `items/<class>.json` from the client's own
`AllowableClass`/weapon-and-armour-proficiency gates alone, before
`loot.json` exists for a build. This module runs AFTER `loot.json` is
built (same "needs the FINAL, post-merge document" placement as
`pipeline.loot.supersede.mark_superseded_items`) and narrows the
per-class catalogue `normalize` already wrote, the same read-modify-
write discipline that function already uses.

Only a QUEST-ONLY item is touched: one whose every `LootSource` bucket
across the whole build is the flat `quest` kind (`kind="quest"`), never
also a raid/dungeon/world/vendor/crafted/rep/pvp source -- an item with
another, independently-verified source stays in every class it can
equip, whatever one particular quest says. The flat `items.json` is
left alone entirely: `pipeline.csvio.check_item_sparse_completeness`'s
own no-shrink gate counts a quest-only row for every class regardless
of which one can actually complete the quest, so narrowing that file
too would trip the gate for the wrong reason.
"""

from __future__ import annotations

import json
import logging
from collections import defaultdict
from pathlib import Path

from pipeline.loot.sources import source_item_ids
from pipeline.models import ClassItems, LootFile
from pipeline.normalize import write_model

logger = logging.getLogger(__name__)


def quest_only_class_exclusions(
    document: LootFile, class_slugs: list[str]
) -> dict[str, set[int]]:
    """class slug -> quest-only item ids that class can never actually
    obtain, because every `QuestSource` entry `document.quests` names for
    that item excludes it (`QuestSource.classes` -- `None` means "any
    class", so a single entry with no restriction, or one that DOES admit
    the class, keeps the item in that class's catalogue).

    An item absent from `document.quests` entirely, or with at least one
    quest entry this build cannot verify a class restriction for, is
    never excluded -- unverifiable stays unlabelled, never guessed
    (tenet 8)."""
    kinds_by_item: dict[int, set[str]] = defaultdict(set)
    for source in document.sources:
        for item_id in source_item_ids(source):
            kinds_by_item[item_id].add(source.kind)
    quest_only_items = {
        item_id for item_id, kinds in kinds_by_item.items() if kinds == {"quest"}
    }
    exclusions: dict[str, set[int]] = defaultdict(set)
    for item_id in quest_only_items:
        entries = document.quests.get(str(item_id))
        if not entries:
            continue
        for class_slug in class_slugs:
            if all(
                entry.classes is not None and class_slug not in entry.classes
                for entry in entries
            ):
                exclusions[class_slug].add(item_id)
    return dict(exclusions)


def apply_quest_class_gate_to_class_items(build_dir: Path, document: LootFile) -> dict[str, int]:
    """Read-modify-write every `items/<class>.json` under `build_dir`,
    dropping a quest-only item `quest_only_class_exclusions` names for
    that class. Returns class slug -> how many rows it dropped; a class
    with nothing dropped is absent from the dict entirely (a no-op run
    rewrites nothing, the same discipline `pipeline.loot.supersede.
    mark_superseded_items` follows for its own two columns)."""
    items_dir = build_dir / "items"
    class_slugs = sorted(path.stem for path in items_dir.glob("*.json"))
    if not class_slugs:
        return {}
    exclusions = quest_only_class_exclusions(document, class_slugs)
    dropped: dict[str, int] = {}
    for class_slug in class_slugs:
        excluded = exclusions.get(class_slug)
        if not excluded:
            continue
        path = items_dir / f"{class_slug}.json"
        record = ClassItems(**json.loads(path.read_text(encoding="utf-8")))
        kept = [item for item in record.items if item.id not in excluded]
        removed = len(record.items) - len(kept)
        if removed == 0:
            continue
        write_model(record.model_copy(update={"items": kept}), path)
        dropped[class_slug] = removed
    if dropped:
        logger.info(
            "quest-class-gate: dropped %d quest-only item rows across %d class catalogues "
            "(class-excluded by their own quest's RequiredClasses): %s",
            sum(dropped.values()),
            len(dropped),
            ", ".join(f"{slug}={count}" for slug, count in sorted(dropped.items())),
        )
    return dropped
