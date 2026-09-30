"""Per-level rotation lines for the addon, from the curated APL + spellranks.

Wave B1 (docs/superpowers/specs/2026-09-28-addon-character-aware-design.md
section 2 item 3): the addon's rotation card and its level-up rotation
toast both want the same thing -- "the abilities you actually cast, in
priority order, at your level" -- read off data/curated/apl/<spec>.json
rather than duplicated by hand. This module turns one spec's priority list
into a per-level-band table of {spellId, name, condition}.

Two decisions are load-bearing:

* Level bands are the rotation-accuracy program's own rungs
  (sim/request/ladder.go's `ladderLevels`), copied here as LEVEL_BANDS
  rather than shared, because Go and Python do not import each other.
  test_addonrotation.py's `test_level_bands_match_the_ladders_own_rungs`
  reads ladder.go's literal so the two cannot drift silently.
* A curated line's `castSpell.spellId.spellId` is the spec's chosen RANK
  (usually the highest one learnable), not a stable per-ability id. A
  level-10 character does not know that rank, so each line is re-resolved
  per band to the highest-level rank of the same ability
  (spellranks.json's own family grouping) that is <= the band's level --
  the same rule the engine's own rank rewrite applies at request time
  (sim/internal/spellranks.HighestLearnedSpellID), and the same rule the
  rotation ladder's own header names. A line whose ability is not learned
  yet at a band (no rank at or below that level) is left out of that
  band entirely, which is how "Sunder Armor opens your rotation now" can
  be a level-20 event and not a level-10 one.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.models import AddonRotationBand, AddonRotationLine

#: sim/request/ladder.go's own `ladderLevels`. See the module docstring.
LEVEL_BANDS = [10, 20, 30, 38, 40, 50, 60]


class AddonRotationError(SystemExit):
    """A curated APL or spellranks file says something the generator will not publish."""


#: A condition longer than this is cut with an ellipsis rather than left
#: whole. Data.lua ships to every player on every login; a curated note
#: written as engine-debugging narration (the RAW note some curated files
#: carry runs past 2,000 characters, even 5,700 on druid-feral's own Claw
#: line) is not "one line" by any reading of the design, and the addon has
#: no use for the part a player would never read anyway.
#:
#: data-followups-10 lane, 2026-09-30, item 8: this cap used to be 160,
#: well under the real first-SENTENCE length `_one_line`'s own boundary
#: search already narrows every curated note down to (measured across
#: every curated/apl/*.json file on this build: 39 lines, 162-672 real
#: characters once correctly cut at the sentence boundary) -- the 160
#: cap was firing on an ordinary, single, complete sentence, not the
#: pathological run-on paragraph it exists to guard against, and cutting
#: mid-word (hunter-marksmanship/hunter-beast-mastery's own Arcane Shot
#: line published "...are the sh…", the "…" landing inside "shots") is
#: what the player-sweep that opened this lane's brief caught. Raised to
#: comfortably clear the real maximum (672) with headroom for a future
#: curated note's own first sentence running a little longer, while
#: still refusing to publish a multi-thousand-character paragraph
#: verbatim if one is ever written with no early sentence boundary at
#: all.
CONDITION_MAX_CHARS = 700


def _one_line(notes: str) -> str:
    """The rotation card's condition: the curated note's first sentence,
    not the paragraph of reasoning some curated files carry (a few run to
    engine-debugging narration hundreds of characters long -- see
    druid-feral's Claw line). Cut at the first sentence boundary (". " or
    a trailing "."); a note with neither is short enough to keep whole, up
    to CONDITION_MAX_CHARS as a last-resort cap for the rare run-on
    sentence. That last-resort cut itself never splits a WORD in half
    (data-followups-10 lane, 2026-09-30, item 8 -- the defect a player
    caught, "...the sh…" mid-word inside "shots"): it backs off to the
    last whitespace before the cap, keeping whatever whole words fit,
    and only falls back to a hard character cut when the cap's own
    window has no whitespace at all (a single unbroken token longer than
    the cap, which no real curated note today is)."""
    notes = notes.strip()
    if not notes:
        return notes
    boundary = notes.find(". ")
    if boundary != -1:
        notes = notes[: boundary + 1]
    elif notes.endswith("."):
        pass
    if len(notes) > CONDITION_MAX_CHARS:
        cut = notes[: CONDITION_MAX_CHARS - 1]
        word_boundary = cut.rfind(" ")
        if word_boundary > 0:
            cut = cut[:word_boundary]
        notes = cut.rstrip() + "…"
    return notes


def _class_slug(spec: str) -> str:
    return spec.split("-", 1)[0]


def _family_ranks(spellranks: dict, class_slug: str) -> dict[str, list[tuple[int, int]]]:
    """Spell family name -> every (level, id) pair it ranks up through,
    sorted ascending by level. Two ranks may share a level (spellranks.json
    sometimes lists an alternate id for the same rank); sorting is stable
    on that tie, which only matters for which of the two `_resolve` picks,
    and either is a real, castable id for that rank."""
    families: dict[str, list[tuple[int, int]]] = {}
    for name, ranks in spellranks.get("classes", {}).get(class_slug, {}).items():
        pairs = sorted((rank["level"], rank["id"]) for rank in ranks)
        # A level-0 row is an alternate id the client carries with no learn
        # level of its own (Rapid Fire's 28755 beside the real 3045 at 26,
        # Aimed Shot's 1236188 beside rank 6 at 60). Left in, it made the
        # family "learned" from level 1, so the addon showed a level-20
        # Beast Mastery hunter a Rapid Fire it cannot train until 26.
        # Keep them only when the family states no level at all.
        levelled = [pair for pair in pairs if pair[0] > 0]
        families[name] = levelled or pairs
    return families


def _id_to_family(families: dict[str, list[tuple[int, int]]]) -> dict[int, str]:
    lookup: dict[int, str] = {}
    for name, ranks in families.items():
        for _, spell_id in ranks:
            lookup[spell_id] = name
    return lookup


def _collect_cast_lines(action: dict, notes: str, lines: list[dict]) -> None:
    """Depth-first over one priority-list action: a direct castSpell is one
    line; a strictSequence/sequence unfolds into one line per step (each
    step is itself an action-shaped dict, not wrapped in another "action"
    key), carrying the sequence's own notes since the curated format has
    no per-step note. Anything else -- autocastOtherCooldowns, the
    paladin primary-seal helper, a bare condition with neither -- names no
    spell id and contributes no line."""
    cast = action.get("castSpell")
    if cast is not None:
        spell_id = cast.get("spellId", {}).get("spellId")
        if spell_id is not None:
            lines.append({"spell_id": spell_id, "notes": notes})
        return
    for key in ("strictSequence", "sequence"):
        sub = action.get(key)
        if isinstance(sub, dict):
            for step in sub.get("actions", []):
                _collect_cast_lines(step, notes, lines)


def _cast_lines(rotation: dict) -> list[dict]:
    """Every priority-list cast, in the curated file's own order."""
    lines: list[dict] = []
    for entry in rotation.get("priorityList", []):
        _collect_cast_lines(entry.get("action", {}), entry.get("notes", ""), lines)
    return lines


def _resolve(families: dict[str, list[tuple[int, int]]], id_to_family: dict[int, str],
             spell_id: int, level: int) -> tuple[str, int] | None:
    """The (family name, id) this ability casts at `level`, or None when
    the curated id is not one spellranks.json tracks for this class (an
    engine-internal or item-proc id the ladder itself would call
    unresolved) or the family's first rank is not learned yet at `level`."""
    family = id_to_family.get(spell_id)
    if family is None:
        return None
    learned = [rank for rank in families[family] if rank[0] <= level]
    if not learned:
        return None
    return family, learned[-1][1]


def build_rotation(spec: str, curated_dir: Path, spellranks: dict) -> list[AddonRotationBand]:
    """One spec's rotation table, one band per LEVEL_BANDS entry."""
    apl_path = curated_dir / "apl" / f"{spec}.json"
    if not apl_path.exists():
        raise AddonRotationError(f"missing {apl_path}")
    curated = json.loads(apl_path.read_text(encoding="utf-8"))
    casts = _cast_lines(curated.get("rotation", {}))

    families = _family_ranks(spellranks, _class_slug(spec))
    id_to_family = _id_to_family(families)

    bands = []
    for level in LEVEL_BANDS:
        lines = []
        for cast in casts:
            resolved = _resolve(families, id_to_family, cast["spell_id"], level)
            if resolved is None:
                continue
            name, resolved_id = resolved
            condition = _one_line(cast["notes"])
            lines.append(AddonRotationLine(spell_id=resolved_id, name=name, condition=condition))
        bands.append(AddonRotationBand(level=level, lines=lines))
    return bands


def build_rotations(
    root: Path, build: str, curated_dir: Path = Path("curated")
) -> dict[str, list[AddonRotationBand]]:
    """Every curated spec's rotation table, keyed by spec slug."""
    spellranks_path = root / build / "spellranks.json"
    if not spellranks_path.exists():
        raise AddonRotationError(f"missing {spellranks_path}")
    spellranks = json.loads(spellranks_path.read_text(encoding="utf-8"))
    apl_dir = curated_dir / "apl"
    specs = sorted(path.stem for path in apl_dir.glob("*.json")) if apl_dir.exists() else []
    return {spec: build_rotation(spec, curated_dir, spellranks) for spec in specs}
