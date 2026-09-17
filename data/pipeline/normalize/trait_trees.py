"""The client's trait tables -> one talents/<class-slug>.json per class.

The trait-table twin of `talent_trees.py`: same output shape, same field
meanings, different source. Names, icons and descriptions still come from the
spell tables, so a talent looks the same to the site whichever reader emitted
it; what differs is that a trait talent is one spell whose per-rank numbers
come off a curve (`pipeline/curves.py`) rather than a chain of per-rank
spells.
"""

from __future__ import annotations

import logging

from pipeline.curves import RankPoints
from pipeline.icons import resolve_icon
from pipeline.models import ClassTalents, TalentEntry, TalentRank, TalentTree
from pipeline.normalize.classes import slugify
from pipeline.normalize.talent_trees import talent_name
from pipeline.normalize.traits import TraitDataError, TraitRows, TraitTalent, read_trait_trees
from pipeline.spelltext import SpellText

logger = logging.getLogger(__name__)


class TraitRankError(ValueError):
    """A talent has ranks the client gives no values for."""


def build_trait_talent_trees(
    rows: TraitRows,
    class_rows: list[dict[str, str]],
    spell_names: dict[int, str],
    spell_text: SpellText,
    rank_points: RankPoints,
    icons: dict[int, str],
    build: str,
) -> list[ClassTalents]:
    slugs = {int(row["ID"]): slugify(row["Name_lang"]) for row in class_rows}
    records: list[ClassTalents] = []
    for tree in read_trait_trees(rows):
        slug = slugs.get(tree.class_id)
        if slug is None:
            raise TraitDataError(
                f"tree {tree.tree_id} names class {tree.class_id}, which has no ChrClasses row"
            )
        records.append(
            ClassTalents(
                build=build,
                class_id=tree.class_id,
                class_slug=slug,
                trees=[
                    TalentTree(
                        id=tab.tab_id,
                        name=tab.name,
                        position=tab.position,
                        background=tab.background.lower(),
                        talents=sorted(
                            (
                                _entry(talent, spell_names, spell_text, rank_points, icons)
                                for talent in tab.talents
                            ),
                            key=lambda t: (t.tier, t.column, t.id),
                        ),
                    )
                    for tab in tree.tabs
                ],
            )
        )
    return sorted(records, key=lambda r: r.class_id)


def _entry(
    talent: TraitTalent,
    spell_names: dict[int, str],
    spell_text: SpellText,
    rank_points: RankPoints,
    icons: dict[int, str],
) -> TalentEntry:
    highest = rank_points.highest_rank(talent.definition_id)
    if talent.max_rank > 1 and highest < talent.max_rank:
        raise TraitRankError(
            f"definition {talent.definition_id} (node {talent.node_id}, spell "
            f"{talent.spell_id}) has {talent.max_rank} ranks but its curves "
            f"reach rank {highest}"
        )
    if highest > talent.max_rank:
        # Three single-rank talents in build 1.60.1.69893 (Ice Block,
        # Thousand Cuts, Last Stand) sit on curves with points the node can
        # never reach. Clamping to MaxRanks is what the client draws.
        logger.debug(
            "definition %s has curve points up to rank %s but only %s ranks; clamping",
            talent.definition_id,
            highest,
            talent.max_rank,
        )
    return TalentEntry(
        id=talent.node_id,
        name=talent_name(talent.node_id, talent.spell_id, spell_names),
        icon=resolve_icon(
            spell_text.icon_file_id(talent.spell_id),
            icons,
            f"talent {talent.node_id} (spell {talent.spell_id})",
        ),
        max_rank=talent.max_rank,
        tier=talent.row,
        column=talent.column,
        prereq_talent_id=talent.prereq_node_id,
        prereq_rank=talent.prereq_rank,
        ranks=[
            TalentRank(
                spell_id=talent.spell_id,
                description=spell_text.describe(
                    talent.spell_id, rank_points.for_rank(talent.definition_id, rank)
                ),
            )
            for rank in range(1, talent.max_rank + 1)
        ],
        spell_id=talent.spell_id,
    )
