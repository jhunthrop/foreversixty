"""Talent + TalentTab + spell text -> one talents/<class-slug>.json per class."""

from __future__ import annotations

import logging

from pipeline.icons import resolve_icon
from pipeline.models import ClassTalents, TalentEntry, TalentRank, TalentTree
from pipeline.normalize.classes import slugify
from pipeline.spelltext import SpellText

logger = logging.getLogger(__name__)

RANK_COLUMNS = [f"SpellRank_{i}" for i in range(9)]

#: Shown when the client has no `SpellName` row for a talent's first rank. The
#: planner never invents a name, and "" would render a nameless cell, so the
#: talent id goes on screen instead: obviously not a real spell name, and it
#: says exactly which row to go and look at.
UNKNOWN_TALENT_NAME = "Unknown talent"


def talent_name(talent_id: int, spell_id: int, spell_names: dict[int, str]) -> str:
    """The first rank's spell name, or a placeholder that names the missing row.

    Shared with `pipeline/normalize/trait_trees.py`, which names a talent the
    same way off the client's trait tables.
    """
    name = spell_names.get(spell_id)
    if name:
        return name
    placeholder = f"{UNKNOWN_TALENT_NAME} {talent_id}"
    logger.warning(
        "talent %s has no name: spell %s is not in SpellName; using %r",
        talent_id,
        spell_id,
        placeholder,
    )
    return placeholder


def _rank_spell_ids(row: dict[str, str]) -> list[int]:
    return [
        int(row[column])
        for column in RANK_COLUMNS
        if row.get(column, "0").strip() not in ("", "0")
    ]


def build_talent_trees(
    talent_rows: list[dict[str, str]],
    tab_rows: list[dict[str, str]],
    class_rows: list[dict[str, str]],
    spell_names: dict[int, str],
    spell_text: SpellText,
    icons: dict[int, str],
    build: str,
) -> list[ClassTalents]:
    tabs = [
        {
            "id": int(row["ID"]),
            "name": row["Name_lang"],
            "position": int(row["OrderIndex"]),
            "class_mask": int(row["ClassMask"]),
            "background": row["BackgroundFile"].lower(),
        }
        for row in tab_rows
    ]
    by_tab: dict[int, list[TalentEntry]] = {}
    for row in talent_rows:
        ranks = _rank_spell_ids(row)
        if not ranks:
            continue
        prereq_talent_id = int(row.get("PrereqTalent_0", "0") or 0) or None
        # PrereqRank_0 is a 0-based rank index; the contract wants a point count.
        prereq_rank = int(row.get("PrereqRank_0", "0") or 0) + 1 if prereq_talent_id else None
        talent_id = int(row["ID"])
        entry = TalentEntry(
            id=talent_id,
            name=talent_name(talent_id, ranks[0], spell_names),
            icon=resolve_icon(
                spell_text.icon_file_id(ranks[0]),
                icons,
                f"talent {talent_id} (spell {ranks[0]})",
            ),
            max_rank=len(ranks),
            tier=int(row["TierID"]),
            column=int(row["ColumnIndex"]),
            prereq_talent_id=prereq_talent_id,
            prereq_rank=prereq_rank,
            ranks=[
                TalentRank(spell_id=spell_id, description=spell_text.describe(spell_id))
                for spell_id in ranks
            ],
            # The legacy table's talent is a chain of per-rank spells, so the
            # talent's own spell is the first rank's -- the same id the client
            # writes when the talent is learned.
            spell_id=ranks[0],
        )
        by_tab.setdefault(int(row["TabID"]), []).append(entry)
    records: list[ClassTalents] = []
    for row in class_rows:
        class_id = int(row["ID"])
        class_mask = 1 << (class_id - 1)
        trees = [
            TalentTree(
                id=tab["id"],
                name=tab["name"],
                position=tab["position"],
                talents=sorted(by_tab.get(tab["id"], []), key=lambda t: (t.tier, t.column, t.id)),
                background=tab["background"],
            )
            for tab in sorted(tabs, key=lambda t: (t["position"], t["id"]))
            if tab["class_mask"] & class_mask
        ]
        if not trees:
            continue
        records.append(
            ClassTalents(
                build=build,
                class_id=class_id,
                class_slug=slugify(row["Name_lang"]),
                trees=trees,
            )
        )
    return sorted(records, key=lambda r: r.class_id)
