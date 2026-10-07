"""Lay Wowhead's live Forever talent payload over the trait-table trees.

The client's trait tables, as wago.tools exports them, miss every talent change Blizzard
ships as a hotfix (the 1 October 2026 Fury rebuild is in neither build 1.60.1.70009 nor
1.60.1.70245). Wowhead's talent payload is read off the live game, so when
`raw/wowhead-talents.json` exists `normalize` keeps the trait reader's trees (their
identity, order, background art and tab icon) and replaces every tree's talents with the
payload's, mapped onto the same `TalentEntry` schema field for field:

* `id`, `row`, `col`, `name`, `icon`, `requires` become the node id, tier, column, name,
  icon and prerequisite;
* `ranks` is one client spell id per rank and `descriptions` one live tooltip per rank;
* a spell id that is in none of the build's client tables was added by a hotfix, so it
  is kept as Wowhead states it and the talent is flagged `hotfix_only`.
"""

from __future__ import annotations

import json
import logging
import re
from collections.abc import Collection
from pathlib import Path

from pipeline.models import ClassTalents, TalentEntry, TalentRank, TalentTree
from pipeline.normalize.forever_talents import ForeverTalentError, rank_count, single_prereq
from pipeline.wowhead_talents import raw_talents_path

logger = logging.getLogger(__name__)

#: The provenance a manifest records for talent files that came from the overlay.
OVERLAY_PROVENANCE = "wowhead-talents-overlay"
OVERLAY_SOURCE_FILE = "raw/wowhead-talents.json"


_SINGULAR = re.compile(
    r"(?P<number>\d+(?:\.\d+)?)(?P<gap>\s*)<!--singular:(?P<one>[^:>]*):(?P<many>[^>]*)-->"
    r".*?<!--singular-->"
)
_SINGULAR_ANYWHERE = re.compile(
    r"<!--singular:(?P<one>[^:>]*):(?P<many>[^>]*)-->.*?<!--singular-->"
)
_INNERMOST_TABLE = re.compile(r"<table\b[^>]*>(?:(?!<table\b).)*?</table>", re.DOTALL)
_COMMENT = re.compile(r"<!--.*?-->", re.DOTALL)
_LINE_BREAK = re.compile(r"<br\s*/?>")
_TAG = re.compile(r"</?[a-zA-Z][^>]*>")
_BLANK_RUN = re.compile(r"\n{3,}")


def _resolve_singular(match: re.Match[str]) -> str:
    word = match["one"] if float(match["number"]) == 1 else match["many"]
    return f"{match['number']}{match['gap']}{word}"


def clean_description(html: str) -> str:
    """Wowhead's tooltip markup as the plain text the planner shows.

    The payload is the HTML Wowhead's own page injects: line breaks, colour spans, marker
    comments around values the page resolves client-side, `singular:one:many` pairs whose
    shown word follows the number before them, and an embedded tooltip table for talents
    that grant a spell. The planner renders text, so the shown value is kept, the pair
    picks its word by that number, and the embedded table is dropped.
    """
    text = _SINGULAR.sub(_resolve_singular, html)
    text = _SINGULAR_ANYWHERE.sub(lambda match: match["many"], text)
    previous = None
    while previous != text:
        previous, text = text, _INNERMOST_TABLE.sub("", text)
    text = _LINE_BREAK.sub("\n", _COMMENT.sub("", text))
    text = _BLANK_RUN.sub("\n\n", _TAG.sub("", text))
    return "\n".join(line.rstrip() for line in text.split("\n")).strip()


def load_overlay_payload(build_dir: Path) -> dict | None:
    """The saved payload for a build, `_meta` removed, or None when none was fetched."""
    path = raw_talents_path(build_dir)
    if not path.is_file():
        return None
    payload = json.loads(path.read_text(encoding="utf-8"))
    payload.pop("_meta", None)
    return payload


def _entry(talent: dict, tree_ids: set[int], known_spells: Collection[int]) -> TalentEntry:
    count = rank_count(talent)
    prereq_id, prereq_rank = single_prereq(talent, tree_ids)
    spell_ids = [int(spell_id) for spell_id in talent["ranks"]]
    descriptions = talent["descriptions"]
    return TalentEntry(
        id=int(talent["id"]),
        name=talent["name"],
        icon=talent.get("icon", ""),
        max_rank=count,
        tier=int(talent["row"]),
        column=int(talent["col"]),
        prereq_talent_id=prereq_id,
        prereq_rank=prereq_rank,
        ranks=[
            TalentRank(spell_id=spell_id, description=clean_description(descriptions[str(rank)]))
            for rank, spell_id in enumerate(spell_ids, start=1)
        ],
        spell_id=spell_ids[0],
        hotfix_only=any(spell_id not in known_spells for spell_id in spell_ids),
    )


def _overlay_tree(tree: TalentTree, raw: dict, known_spells: Collection[int]) -> TalentTree:
    tree_ids = {int(talent["id"]) for talent in raw.values()}
    talents = sorted(
        (_entry(talent, tree_ids, known_spells) for talent in raw.values()),
        key=lambda t: (t.tier, t.column, t.id),
    )
    cells = [(t.tier, t.column) for t in talents]
    if len(set(cells)) != len(cells):
        raise ForeverTalentError(f"tree {tree.id} places two talents in one cell")
    return tree.model_copy(update={"talents": talents})


def overlay_wowhead_talents(
    records: list[ClassTalents], payload: dict, known_spells: Collection[int]
) -> list[ClassTalents]:
    """`records` with every tree's talents replaced by the payload's, as new records.

    Fails if the payload names a tree we do not have or lacks one we do: a partial
    overlay would mix two eras of the tree.
    """
    raw_by_tree: dict[str, dict] = payload.get("talents") or {}
    ours = {str(tree.id) for record in records for tree in record.trees}
    if set(raw_by_tree) != ours:
        raise ForeverTalentError(
            f"payload trees {sorted(set(raw_by_tree) - ours)} are not ours; "
            f"ours {sorted(ours - set(raw_by_tree))} are missing from the payload"
        )
    out = [
        record.model_copy(
            update={
                "trees": [
                    _overlay_tree(tree, raw_by_tree[str(tree.id)], known_spells)
                    for tree in record.trees
                ]
            }
        )
        for record in records
    ]
    flagged = sum(t.hotfix_only for r in out for tree in r.trees for t in tree.talents)
    logger.info("wowhead talents overlay: %d hotfix-only talents flagged", flagged)
    return out
