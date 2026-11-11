"""`curated/spelltext-gaps.json`: description tokens no client table carries a value for.

A token such as Berserking's `$26635d` names a spell the server knows and the client
export does not, so the renderer has nothing to put there. Each entry gives the value
Wowhead's Forever tooltip prints, with its source; `pipeline.spelltext.load_spell_text`
writes it into that spell's template before rendering and fails when the template no
longer holds the token, so a stale entry is removed rather than silently ignored.
"""

from __future__ import annotations

import json
from pathlib import Path

from pydantic import BaseModel

from pipeline.curated import CuratedError, parse_sources

GAPS_FILE = "spelltext-gaps.json"


class SpellTextGap(BaseModel):
    spell_id: int
    token: str
    value: str
    sources: list[dict]


def load_spelltext_gaps(curated: Path) -> dict[int, dict[str, str]]:
    """Spell id -> token -> the words to print, every entry sourced. A curated directory
    without the file (a test fixture's) has no gaps."""
    path = curated / GAPS_FILE
    if not path.exists():
        return {}
    out: dict[int, dict[str, str]] = {}
    for index, entry in enumerate(json.loads(path.read_text(encoding="utf-8"))["gaps"]):
        where = f"{GAPS_FILE} gap {index}"
        gap = SpellTextGap(**entry)
        parse_sources(gap.sources, where)
        if not gap.token.startswith("$") or not gap.value.strip():
            raise CuratedError(f"{where} needs a `$` token and a value")
        out.setdefault(gap.spell_id, {})[gap.token] = gap.value
    return out
