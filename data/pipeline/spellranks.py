"""spellranks.json: per class, each spell name's player-castable rank chain --
`sim/request`'s own answer to "what id does this rotation slot resolve to at
this character's level" (docs/superpowers/specs/2026-09-27-level-aware-sim-
design.md, design point 4). It reads `builds/<build>/spellconst/<class>.json`
(`pipeline.simconst`), the constants file `normalize` neither writes nor
reads, so this has to be its own command rather than a step inside either.

**The reference caveat.** `spellconst` groups nothing by name -- it is keyed
by spell id -- so grouping by name here meets the client's own habit of
reusing one ability's name across ids that never share a rank. Sinister
Strike carries five rank-0 ids at level 20 beside its real ranks 1-8: NPC or
internal copies, not ranks. The player chain for a name is its entries with
rank >= 1, sorted by rank then id; a rank-0 duplicate never joins that chain.
A name whose entries are *all* rank 0 has no chain at all -- it is a spell
without ranks, most often carrying a single id, occasionally more when the
client duplicates an unranked ability under NPC or spec-variant ids -- and
keeps every one of its own entries, sorted by id, rather than being collapsed
or dropped: `sim/request` still needs the id(s) and the learn level to gate
the action, even though there is no rank to rewrite to.

A name is included at all only when its resulting list carries more than one
id or a nonzero level -- the vast majority of spells, unranked and learned at
level 0, need neither a rewrite nor a level gate and are left out entirely.
"""

from __future__ import annotations

import logging
from collections import defaultdict
from pathlib import Path

from pipeline.manifest import refresh_manifest
from pipeline.models import ClassSpellConstants, SpellConstant, SpellRank, SpellRanksFile
from pipeline.normalize import _write
from pipeline.simconst import SPELLCONST

logger = logging.getLogger(__name__)

SPELLRANKS_FILE = "spellranks.json"

#: One spell entry as (id, rank, level), the minimal shape this module works in
#: before it becomes a `SpellRank`.
_Entry = tuple[int, int, int]


def load_class_spell_constants(spellconst_dir: Path) -> list[ClassSpellConstants]:
    """Every class's spell constants in a `spellconst/` directory, in file order."""
    return [
        ClassSpellConstants.model_validate_json(path.read_text(encoding="utf-8"))
        for path in sorted(spellconst_dir.glob("*.json"))
    ]


def _entries_by_name(spells: dict[str, SpellConstant]) -> dict[str, list[_Entry]]:
    groups: dict[str, list[_Entry]] = defaultdict(list)
    for spell_id, spell in spells.items():
        groups[spell.name].append((int(spell_id), spell.rank, spell.spell_level))
    return groups


def _rank_chain(entries: list[_Entry]) -> list[_Entry]:
    """The player chain for one spell name: see the module docstring's caveat."""
    ranked = [entry for entry in entries if entry[1] >= 1]
    if ranked:
        return sorted(ranked, key=lambda entry: (entry[1], entry[0]))
    return sorted(entries, key=lambda entry: entry[0])


def _spell_ranks_for_class(spells: dict[str, SpellConstant]) -> dict[str, list[SpellRank]]:
    names: dict[str, list[SpellRank]] = {}
    for name, entries in sorted(_entries_by_name(spells).items()):
        chain = _rank_chain(entries)
        if len(chain) <= 1 and all(level == 0 for _id, _rank, level in chain):
            continue
        names[name] = [SpellRank(id=i, rank=r, level=lvl) for i, r, lvl in chain]
    return names


def build_spell_ranks(build: str, records: list[ClassSpellConstants]) -> SpellRanksFile:
    """The pure transform: every class's spell constants, no file I/O."""
    classes = {
        record.class_slug: _spell_ranks_for_class(record.spells)
        for record in sorted(records, key=lambda r: r.class_slug)
    }
    return SpellRanksFile(build=build, classes=classes)


def write_spell_ranks(build: str, root: Path = Path("builds")) -> Path:
    build_dir = root / build
    spellconst_dir = build_dir / SPELLCONST
    if not spellconst_dir.exists():
        raise SystemExit(f"no {spellconst_dir}; run `python -m pipeline simconst` first")
    result = build_spell_ranks(build, load_class_spell_constants(spellconst_dir))
    path = build_dir / SPELLRANKS_FILE
    _write(result.model_dump(), path, sort_keys=True)
    refresh_manifest(build_dir)
    total = sum(len(names) for names in result.classes.values())
    logger.info(
        "wrote %s: %d ranked spell names over %d classes", path, total, len(result.classes)
    )
    return path
