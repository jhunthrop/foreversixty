"""The profession skill each crafted item needs, onto loot.json's crafted sources.

Spec 2026-11-05-web-layout 5.10 item 3: a crafted row reads `Crafted ·
Leatherworking (skill 40)`. The number is `ClassicDbCraftedRecipe.
required_skill` (`pipeline.classicdb_crafted`): the trainer's
`reqskillvalue` or the recipe item's `RequiredSkillRank` in the pinned
cmangos/classic-db dump. The client cannot supply it: its
SkillLineAbility rows carry MinSkillLineRank 1 for every craft, and the
trivial ranks there are the yellow and grey thresholds.

A Forever-new id that re-itemisation tied to a classic twin
(`reitemised_from` on any source, `pipeline.loot.reitemise`) is crafted
from the twin's recipe, so it takes the twin's skill. An item neither id
resolves is left out of the map: never guessed.
"""

from __future__ import annotations

import logging

from pipeline.classicdb_crafted import ClassicDbCraftedRecipe
from pipeline.models import LootFile

logger = logging.getLogger(__name__)


def _classic_twins(document: LootFile) -> dict[int, int]:
    """Forever-new item id -> the classic id re-itemisation copied it from,
    from every source and boss that records one."""
    twins: dict[int, int] = {}
    for source in document.sources:
        for mapping in [source.reitemised_from, *(b.reitemised_from for b in source.bosses or [])]:
            twins |= {int(item_id): classic for item_id, classic in (mapping or {}).items()}
    return twins


def apply_crafted_skill(
    document: LootFile, classic_crafted: dict[int, ClassicDbCraftedRecipe]
) -> LootFile:
    """Every crafted source with `required_skill` (item id -> skill) for the
    items whose recipe skill is known."""
    twins = _classic_twins(document)

    def skill(item_id: int) -> int | None:
        recipe = classic_crafted.get(item_id) or classic_crafted.get(twins.get(item_id, -1))
        return None if recipe is None else recipe.required_skill

    sources = []
    unknown = 0
    for source in document.sources:
        if source.kind != "crafted":
            sources.append(source)
            continue
        skills = {str(i): s for i in source.items or [] if (s := skill(i)) is not None}
        unknown += len(source.items or []) - len(skills)
        sources.append(source.model_copy(update={"required_skill": skills or None}))
    logger.info("crafted-skill: %d crafted items with no recipe skill in classic-db", unknown)
    return document.model_copy(update={"sources": sources})
