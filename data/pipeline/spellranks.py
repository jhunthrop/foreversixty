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

import json
import logging
from collections import defaultdict
from pathlib import Path

from pipeline.manifest import refresh_manifest
from pipeline.models import ClassSpellConstants, SpellConstant, SpellRank, SpellRanksFile
from pipeline.normalize import _write
from pipeline.simconst import SPELLCONST

logger = logging.getLogger(__name__)

SPELLRANKS_FILE = "spellranks.json"

#: The curated statement of which spell ids are Ahn'Qiraj book ranks; a rank whose
#: id is listed is written `book: true`, and nothing reads it as learnable unless
#: the engine's `core.IncludeAQ` is on.
BOOK_RANKS_FILE = "book-ranks.json"

#: The curated statement of ranks whose effect is weaker than the rank before them in
#: the client; a rank whose id is listed is written `inferior: true`, and the rewrite
#: never resolves a cast to it.
INFERIOR_RANKS_FILE = "inferior-ranks.json"

#: data/curated, found from this file rather than the working directory.
CURATED_DIR = Path(__file__).resolve().parent.parent / "curated"

#: One spell entry as (id, rank, level, castable), the minimal shape this module
#: works in before it becomes a `SpellRank`. `castable` is whether the client
#: gives the id a mana cost, a cast time or a global cooldown -- what a spell the
#: player presses has, and what the copies beside it lack.
_Entry = tuple[int, int, int, bool]

#: The client names an NPC's or a scripted event's duplicate of a real player
#: spell "Copy of <name>" rather than reusing the player's own name (Copy of
#: Deadly Poison IV, Copy of Frostbolt, Copy of Mortal Strike). It is never a
#: spell the player learns a rank of, and unlike the rank-0 duplicates the
#: module docstring's "reference caveat" describes, it does not even share the
#: real spell's name to be filtered out of that spell's own chain by rank --
#: it forms a one-name chain of its own instead, which is how these were
#: found leaking into `sim/request`'s ladder as "learned but unused" abilities
#: (e.g. "Copy of Mortal Strike" beside the real Mortal Strike). Dropped
#: before grouping, so it never reaches `_entries_by_name` at all.
_COPY_NAME_PREFIX = "Copy of "


def _castable(spell: SpellConstant) -> bool:
    return spell.cost > 0 or spell.cast_time_ms > 0 or spell.gcd_ms > 0


def is_copy_name(name: str) -> bool:
    return name.startswith(_COPY_NAME_PREFIX)


def load_class_spell_constants(spellconst_dir: Path) -> list[ClassSpellConstants]:
    """Every class's spell constants in a `spellconst/` directory, in file order."""
    return [
        ClassSpellConstants.model_validate_json(path.read_text(encoding="utf-8"))
        for path in sorted(spellconst_dir.glob("*.json"))
    ]


def _entries_by_name(spells: dict[str, SpellConstant]) -> dict[str, list[_Entry]]:
    groups: dict[str, list[_Entry]] = defaultdict(list)
    for spell_id, spell in spells.items():
        if is_copy_name(spell.name):
            continue
        groups[spell.name].append((int(spell_id), spell.rank, spell.spell_level, _castable(spell)))
    return groups


def count_copy_names(records: list[ClassSpellConstants]) -> int:
    """How many "Copy of ..." ids `_entries_by_name` dropped, across every
    class -- for `write_spell_ranks`' log line, the same way `loot`'s and
    `normalize`'s own stats count what they filtered or derived."""
    return sum(
        1 for record in records for spell in record.spells.values() if is_copy_name(spell.name)
    )


def _rank_chain(entries: list[_Entry]) -> list[_Entry]:
    """The player chain for one spell name: see the module docstring's caveat."""
    ranked = [entry for entry in entries if entry[1] >= 1]
    if ranked:
        # The 1.60 client keeps cost-less, cast-less, GCD-less copies of a ranked spell
        # beside the player's own ids at the same ranks (Lightning Bolt 403 beside
        # 408439, Magma Totem's totem 10585 beside its pulse 10579); a rewrite that
        # landed on one of those would name a spell the engine never registers, or
        # registers without the periodic effect the rotation asks about. Keep the
        # castable ids when the chain has any.
        castable = [entry for entry in ranked if entry[3]]
        chain = castable or ranked
        return sorted(chain, key=lambda entry: (entry[1], entry[0]))
    return sorted(entries, key=lambda entry: entry[0])


def load_book_ids(curated: Path) -> frozenset[int]:
    """The ids `curated/book-ranks.json` names as Ahn'Qiraj book ranks."""
    path = curated / BOOK_RANKS_FILE
    if not path.exists():
        raise SystemExit(f"no {path}; the spellranks emitter marks book ranks from it")
    books = json.loads(path.read_text(encoding="utf-8"))["books"]
    return frozenset(int(book["id"]) for book in books)


def load_inferior_ids(curated: Path) -> frozenset[int]:
    """The ids `curated/inferior-ranks.json` names as ranks weaker than their predecessor."""
    path = curated / INFERIOR_RANKS_FILE
    if not path.exists():
        raise SystemExit(f"no {path}; the spellranks emitter marks inferior ranks from it")
    inferior = json.loads(path.read_text(encoding="utf-8"))["inferior"]
    return frozenset(int(entry["id"]) for entry in inferior)


def _spell_ranks_for_class(
    spells: dict[str, SpellConstant],
    book_ids: frozenset[int] = frozenset(),
    inferior_ids: frozenset[int] = frozenset(),
) -> dict[str, list[SpellRank]]:
    names: dict[str, list[SpellRank]] = {}
    for name, entries in sorted(_entries_by_name(spells).items()):
        chain = _rank_chain(entries)
        if len(chain) <= 1 and all(level == 0 for _id, _rank, level, _castable in chain):
            continue
        names[name] = [
            SpellRank(id=i, rank=r, level=lvl, book=i in book_ids, inferior=i in inferior_ids)
            for i, r, lvl, _castable in chain
        ]
    return names


def build_spell_ranks(
    build: str,
    records: list[ClassSpellConstants],
    book_ids: frozenset[int] = frozenset(),
    inferior_ids: frozenset[int] = frozenset(),
) -> SpellRanksFile:
    """The pure transform: every class's spell constants, no file I/O."""
    classes = {
        record.class_slug: _spell_ranks_for_class(record.spells, book_ids, inferior_ids)
        for record in sorted(records, key=lambda r: r.class_slug)
    }
    return SpellRanksFile(build=build, classes=classes)


def write_spell_ranks(
    build: str, root: Path = Path("builds"), curated: Path = CURATED_DIR
) -> Path:
    build_dir = root / build
    spellconst_dir = build_dir / SPELLCONST
    if not spellconst_dir.exists():
        raise SystemExit(f"no {spellconst_dir}; run `python -m pipeline simconst` first")
    records = load_class_spell_constants(spellconst_dir)
    result = build_spell_ranks(
        build, records, load_book_ids(curated), load_inferior_ids(curated)
    )
    path = build_dir / SPELLRANKS_FILE
    _write(result.model_dump(exclude_defaults=True), path, sort_keys=True)
    refresh_manifest(build_dir)
    total = sum(len(names) for names in result.classes.values())
    logger.info(
        'wrote %s: %d ranked spell names over %d classes; %d "Copy of" ids dropped',
        path,
        total,
        len(result.classes),
        count_copy_names(records),
    )
    return path
