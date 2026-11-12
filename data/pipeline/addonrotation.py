"""Per-level rotation lines for the addon, from the curated APL + spellranks.

Wave B1 (docs/superpowers/specs/2026-09-28-addon-character-aware-design.md
section 2 item 3): the addon's rotation card and its level-up rotation
toast both want the same thing -- "the abilities you actually cast, in
priority order, at your level" -- read off data/curated/apl/<spec>.json
rather than duplicated by hand. This module turns one spec's priority list
into a per-level-band table of {spellId, name, condition, icon}.

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
import re
from collections.abc import Callable, Mapping
from dataclasses import dataclass
from pathlib import Path
from types import MappingProxyType
from typing import Any

from pipeline.csvio import read_csv
from pipeline.icons import PLACEHOLDER_ICON, icon_names, resolve_icon
from pipeline.models import AddonRotationBand, AddonRotationLine, PanelWhen, check_panel_when
from pipeline.spelltext import SpellText, load_spell_text
from pipeline.talent_levels import (
    first_point_levels,
    floor_pairs,
    load_spec_rank_levels,
    load_talent_owned_auras,
    tier_level,
)

#: sim/request/ladder.go's own `ladderLevels`. See the module docstring.
LEVEL_BANDS = [10, 20, 30, 38, 40, 50, 60]

#: The level at which the site runs a curated list exactly as written.
MAX_LEVEL = 60


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


def _family_ranks(
    spellranks: dict,
    class_slug: str,
    trainable_ids: frozenset[int] = frozenset(),
    talent_spells: Mapping[str, list[tuple[int, int]]] = MappingProxyType({}),
) -> dict[str, list[tuple[int, int]]]:
    """Spell family name -> every (level, id) pair it ranks up through,
    sorted ascending by level. Two ranks may share a level (spellranks.json
    sometimes lists an alternate id for the same rank); sorting is stable
    on that tie, which only matters for which of the two `_resolve` picks,
    and either is a real, castable id for that rank.

    An unranked family keeps only the ids the class's trainables list
    (`trainable_ids`, from SkillLineAbility), when it has any: the client
    carries copies with no trainer row at lower levels (Recklessness 13847
    at 22 beside the trained 1719 at 50, Evocation 456397 at 1, Barkskin
    428713 at 1), and the addon showed them bands before the player could
    learn the spell. A family none of whose ids are trainable (a talent or
    rune spell) keeps every row."""
    families: dict[str, list[tuple[int, int]]] = {}
    for name, ranks in spellranks.get("classes", {}).get(class_slug, {}).items():
        if all(rank.get("rank", 0) == 0 for rank in ranks):
            trained = [rank for rank in ranks if rank["id"] in trainable_ids]
            ranks = trained or ranks
        # An Ahn'Qiraj book rank is not a trainer's: the sim's rank rewrite
        # never casts one, so neither does the addon.
        ranks = [rank for rank in ranks if not rank.get("book")]
        # By level only: two ids of one tier keep spellranks.json's own
        # order, which is the order the sim's tier lists them in.
        pairs = sorted(((rank["level"], rank["id"]) for rank in ranks), key=lambda pair: pair[0])
        # A level-0 row is an alternate id the client carries with no learn
        # level of its own (Rapid Fire's 28755 beside the real 3045 at 26,
        # Aimed Shot's 1236188 beside rank 6 at 60). Left in, it made the
        # family "learned" from level 1, so the addon showed a level-20
        # Beast Mastery hunter a Rapid Fire it cannot train until 26.
        # Keep them only when the family states no level at all.
        levelled = [pair for pair in pairs if pair[0] > 0]
        families[name] = levelled or pairs
    # A family with no id in the class's trainables and no talent spell is a Season of
    # Discovery rune (SkillLineAbility AcquireMethod 3): a Forever player never learns it, so
    # no rotation line may name it (Water Shield 408510 sat in Restoration's band 20 on
    # 1.60.1.70790; design/specs/2026-11-08-addon-ux.md section 4.3, the trainables gate).
    # With no trainables file for the class there is nothing to judge by, and every family stays.
    if trainable_ids:
        talent_ids = {spell_id for pairs in talent_spells.values() for _, spell_id in pairs}
        families = {
            name: pairs
            for name, pairs in families.items()
            if any(spell_id in trainable_ids or spell_id in talent_ids for _, spell_id in pairs)
        }
    # A talent's own spell (Stormstrike, Adrenaline Rush, Cold Blood) is in no
    # trainer's table, so spellranks.json does not track it; the sim casts it
    # once the build takes the talent, so the rotation names it from the
    # first level its tier opens (load_talent_spells).
    tracked = {spell_id for pairs in families.values() for _, spell_id in pairs}
    for name, pairs in talent_spells.items():
        if name not in families and not any(spell_id in tracked for _, spell_id in pairs):
            families[name] = list(pairs)
    return families


def _id_to_family(families: dict[str, list[tuple[int, int]]]) -> dict[int, str]:
    lookup: dict[int, str] = {}
    for name, ranks in families.items():
        for _, spell_id in ranks:
            lookup[spell_id] = name
    return lookup


def _collect_cast_lines(action: dict, notes: str, lines: list[dict], line: int) -> None:
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
            lines.append({"spell_id": spell_id, "notes": notes, "line": line})
        return
    for key in ("strictSequence", "sequence"):
        sub = action.get(key)
        if isinstance(sub, dict):
            for step in sub.get("actions", []):
                _collect_cast_lines(step, notes, lines, line)


def _cast_lines(
    rotation: dict, panel_when: Mapping[int, tuple[PanelWhen, ...]] = MappingProxyType({})
) -> list[dict]:
    """Every priority-list cast, in the curated file's own order, with the
    0-based index of the entry it sits under (a rank choice's "lines"), that
    entry's condition (`_level_allows`) and its panel lines (`panel_when`, the
    curated file's top-level map; None where the entry has none). A sequence's
    steps share their entry's line, as they share its notes."""
    lines: list[dict] = []
    for index, entry in enumerate(rotation.get("priorityList", [])):
        action = entry.get("action", {})
        start = len(lines)
        _collect_cast_lines(action, entry.get("notes", ""), lines, index)
        for line in lines[start:]:
            line["condition"] = action.get("condition")
            line["when"] = panel_when.get(index)
    return lines


def _panel_when(curated: dict, apl_path: Path) -> dict[int, tuple[PanelWhen, ...]]:
    """The curated file's `panel_when`, checked (`check_panel_when`), by step index."""
    raw = curated.get("panel_when", {})
    if not isinstance(raw, dict):
        raise AddonRotationError(f"{apl_path}: panel_when must be an object")
    try:
        return check_panel_when(raw)
    except ValueError as error:
        raise AddonRotationError(f"{apl_path}: {error}") from error


def _check_panel_when_steps(
    panel_when: Mapping[int, tuple[PanelWhen, ...]], casts: list[dict], apl_path: Path
) -> None:
    """Every `panel_when` key names a priority-list step that casts a spell: a line on any
    other step (out of range, or an autocast/seal helper) would never print."""
    cast_steps = {cast["line"] for cast in casts}
    stray = sorted(index for index in panel_when if index not in cast_steps)
    if stray:
        raise AddonRotationError(
            f"{apl_path}: panel_when names step(s) {stray}, which cast no spell"
        )


def condition_at(cast: dict, level: int) -> str | None:
    """A line's panel condition at `level`: its curated `panel_when` line verbatim (the
    variant with the highest `from_level` at or below `level`), else the first sentence of
    its step's note (`_one_line`); None when every variant starts above `level`."""
    when = cast.get("when")
    if when is None:
        return _one_line(cast["notes"])
    lines = [variant.text for variant in when if variant.from_level <= level]
    return lines[-1] if lines else None


def _condition_of(cast: dict, level: int, spec: str) -> str:
    """`condition_at`, refusing a step whose first variant starts above a band it prints
    in: that band would have no line."""
    condition = condition_at(cast, level)
    if condition is None:
        raise AddonRotationError(
            f"{spec}: panel_when[{cast['line']}] starts at level {cast['when'][0].from_level}, "
            f"after the step's band at {level}"
        )
    return condition


#: The engine's APLValueCompare ops a currentLevel clause can use.
_LEVEL_OPS: dict[str, Callable[[float, float], bool]] = {
    "OpEq": lambda a, b: a == b,
    "OpNe": lambda a, b: a != b,
    "OpLt": lambda a, b: a < b,
    "OpLe": lambda a, b: a <= b,
    "OpGt": lambda a, b: a > b,
    "OpGe": lambda a, b: a >= b,
}


def _level_operand(node: object, level: int) -> float | None:
    """A comparison side at `level`: the character's level, a constant, else None."""
    if not isinstance(node, dict):
        return None
    if "currentLevel" in node:
        return float(level)
    const = node.get("const")
    if isinstance(const, dict):
        try:
            return float(const.get("val"))
        except (TypeError, ValueError):
            return None
    return None


def _level_truth(condition: object, level: int) -> bool | None:
    """The condition at `level` as far as its currentLevel comparisons decide
    it: True, False, or None (open). Mirrors web/src/lib/bis/apl-level-gate.ts."""
    if condition is None:
        return True
    if not isinstance(condition, dict):
        return None
    cmp = condition.get("cmp")
    if isinstance(cmp, dict):
        sides = (cmp.get("lhs"), cmp.get("rhs"))
        compare = _LEVEL_OPS.get(cmp.get("op", ""))
        if compare is None or not any(isinstance(s, dict) and "currentLevel" in s for s in sides):
            return None
        lhs, rhs = (_level_operand(side, level) for side in sides)
        return None if lhs is None or rhs is None else compare(lhs, rhs)
    inner = condition.get("not")
    if isinstance(inner, dict):
        value = _level_truth(inner.get("val"), level)
        return None if value is None else not value
    for key in ("and", "or"):
        group = condition.get(key)
        if isinstance(group, dict):
            values = [_level_truth(v, level) for v in group.get("vals", [])]
            if key == "and":
                if False in values:
                    return False
                return True if all(v is True for v in values) else None
            if True in values:
                return True
            return False if all(v is False for v in values) else None
    return None


def _level_allows(condition: object, level: int) -> bool:
    """False only when a currentLevel comparison rules the line out at `level`
    (round 19: a line scoped to a level range)."""
    return _level_truth(condition, level) is not False


#: The APL values that ask whether a spell or aura is known, and the key naming its id.
_KNOWN_VALUES = {"auraIsKnown": "auraId", "spellIsKnown": "spellId"}


def _known_gate(term: object) -> tuple[int, int] | None:
    """(spell id, rank) an auraIsKnown or spellIsKnown term asks for: its rank tag (a talent
    aura registered per rank, Maelstrom Weapon's tag 5) or rank 1 untagged."""
    for value, id_key in _KNOWN_VALUES.items():
        node = term.get(value) if isinstance(term, dict) else None
        action = node.get(id_key, {}) if isinstance(node, dict) else None
        spell_id = action.get("spellId") if isinstance(action, dict) else None
        if isinstance(spell_id, int):
            tag = action.get("tag")
            return spell_id, tag if isinstance(tag, int) and tag > 0 else 1
    return None


def _condition_terms(condition: object) -> list[object]:
    """The whole condition, or its top-level and's terms."""
    if not isinstance(condition, dict):
        return []
    group = condition.get("and")
    return group.get("vals", []) if isinstance(group, dict) else [condition]


def _required_known(condition: object) -> list[tuple[int, int]]:
    """The (spell id, rank) gates `condition` needs known: an auraIsKnown or spellIsKnown that
    is the whole condition or one of its top-level and's terms (an or-ed one gates nothing)."""
    return [gate for term in _condition_terms(condition) if (gate := _known_gate(term))]


def _excluded_ranks(condition: object) -> list[tuple[int, int]]:
    """The rank-tagged gates `condition` needs NOT known (a negated top-level term with a rank
    tag): the line stops firing once the build takes that rank. An untagged negation is left
    alone (round 23's rule): it may name a trainer spell."""
    gates = []
    for term in _condition_terms(condition):
        inner = term.get("not", {}).get("val") if isinstance(term, dict) else None
        action = next((v for v in (inner or {}).values() if isinstance(v, dict)), {})
        tagged = any(isinstance(a, dict) and "tag" in a for a in action.values())
        if tagged and (gate := _known_gate(inner)):
            gates.append(gate)
    return gates


def _rank_floor(
    gate: tuple[int, int], level: int, rank_levels: Mapping[tuple[int, int], int]
) -> int:
    """The level a gate's rank comes on the published build: `level` (no floor) for a spell no
    talent teaches, past 60 for a rank of a taken talent the build never reaches."""
    spell_id, rank = gate
    if gate in rank_levels:
        return rank_levels[gate]
    return MAX_RANK_FLOOR if (spell_id, 1) in rank_levels else level


#: Past the level cap: a rank the published build never takes.
MAX_RANK_FLOOR = 61


def _talents_allow(
    condition: object, level: int, rank_levels: Mapping[tuple[int, int], int]
) -> bool:
    """False while `condition` needs a talent rank known that the published build has not
    taken by `level` (round 23: Restoration's Healing Wave after Nature's Swiftness printed
    from 10, the talent comes at 44; round 24: a rank tag waits for that rank, Enhancement's
    five-stack Lightning Bolt for Maelstrom Weapon 5/5), or once it has taken a rank the
    condition needs unknown (round 24: Arcane's Missiles until Missile Barrage rank 2)."""
    starts = all(level >= _rank_floor(g, level, rank_levels) for g in _required_known(condition))
    ends = any(g in rank_levels and level >= rank_levels[g] for g in _excluded_ranks(condition))
    return starts and not ends


def _trained_allow(
    condition: object,
    level: int,
    families: Mapping[str, list[tuple[int, int]]],
    id_to_family: Mapping[int, str],
) -> bool:
    """False while `condition` needs a trained spell known (a top-level spellIsKnown or
    auraIsKnown on a spell a rank family tracks) that the character learns after `level`:
    the engine's spellIsKnown is false until the first rank is learned."""
    for spell_id, _ in _required_known(condition):
        family = id_to_family.get(spell_id)
        if family is not None and families[family] and families[family][0][0] > level:
            return False
    return True


def apl_cast_spell_ids(rotation: object) -> set[int]:
    """Every spell id a curated rotation casts, anywhere in it: the pre-pull
    actions (the guide's opener: a warlock's Summon Imp, a druid's Prowl) as
    well as the priority list, whether or not a rank family resolves it."""
    ids: set[int] = set()
    if isinstance(rotation, dict):
        cast = rotation.get("castSpell")
        if isinstance(cast, dict) and isinstance(cast.get("spellId"), dict):
            spell_id = cast["spellId"].get("spellId")
            if isinstance(spell_id, int) and not isinstance(spell_id, bool):
                ids.add(spell_id)
        for value in rotation.values():
            ids |= apl_cast_spell_ids(value)
    elif isinstance(rotation, list):
        for value in rotation:
            ids |= apl_cast_spell_ids(value)
    return ids


#: Forever's own spell ids start here; a rank with a classic id and a
#: Forever copy resolves within the band the list names
#: (sim/internal/spellranks foreverIDFloor).
FOREVER_ID_FLOOR = 100000


def _rank_choice_id(
    spellranks: dict, class_slug: str, family: str, spell_id: int, rank: int, level: int
) -> int | None:
    """The id of `rank` of `family` learned by `level`, in the list's id band
    (sim/internal/spellranks SpellIDOfRank), or None."""
    ids = [
        row["id"]
        for row in spellranks.get("classes", {}).get(class_slug, {}).get(family, [])
        if row.get("rank") == rank
        and not row.get("book")
        and not row.get("inferior")
        and 0 < row["level"] <= level
    ]
    same_band = [i for i in ids if (i >= FOREVER_ID_FLOOR) == (spell_id >= FOREVER_ID_FLOOR)]
    return (same_band or ids or [None])[0]


#: The presets a band is published under (sim/request PublishedPreset): the bare
#: preset below RAID_PRESET_LEVEL, the raid preset from it.
BARE_PRESET = "bare"
RAID_PRESET = "raid"
RAID_PRESET_LEVEL = 60


def published_preset(level: int) -> str:
    """The preset `level` is published under (sim/request PublishedPreset)."""
    return RAID_PRESET if level >= RAID_PRESET_LEVEL else BARE_PRESET


def _band_ranks(choices: dict, spell_id: int, level: int, line: int) -> dict[str | None, int]:
    """The ranks data/curated/rank-choices.json picks for the list's
    `spell_id` at `level` (one spec's rows), keyed by each race slug a row
    names (None for a row of every race); empty outside every row. A row for
    one preset applies only at the levels published under it (`published_preset`:
    the raid preset at 60, round 22's Holy paladin Flash of Light rank 5), a row
    naming no preset at every level, and a row scoped to lines only on those
    priority-list entries. sim/request refuses two rows that meet one race."""
    preset = published_preset(level)
    ranks: dict[str | None, int] = {}
    for row in choices.get(spell_id, []):
        if row.get("preset", preset) != preset:
            continue
        if row["lines"] is not None and line not in row["lines"]:
            continue
        if row["from_level"] <= level <= row["to_level"]:
            for race in row.get("races") or [None]:
                ranks[race] = row["rank"]
    return ranks


@dataclass(frozen=True)
class _RankPick:
    """One cast line's family at one band, for picking a chosen rank's id."""

    spellranks: dict
    class_slug: str
    family: str
    spell_id: int
    level: int

    def id_of(self, rank: int | None, fallback: int) -> int:
        """`rank`'s id (`_rank_choice_id`), or `fallback` with no rank or no such id."""
        if rank is None:
            return fallback
        chosen = _rank_choice_id(
            self.spellranks, self.class_slug, self.family, self.spell_id, rank, self.level
        )
        return chosen if chosen is not None else fallback


def _race_ids(pick: _RankPick, ranks: dict[str | None, int], resolved_id: int) -> dict[str, int]:
    """Each keyed race's id where its row casts other than `resolved_id`
    (AddonRotationLine.race_spell_ids); a race whose row resolves to
    `resolved_id` is left out."""
    ids = {race: pick.id_of(rank, resolved_id) for race, rank in ranks.items() if race is not None}
    return {race: i for race, i in sorted(ids.items()) if i != resolved_id}


def load_rank_choices(curated_dir: Path, spec: str) -> dict[int, list[dict]]:
    """One spec's per-band rank choices, by the list's authored id, each row
    carrying the priority-list lines it applies to (`_scoped_rows`). The file
    is optional here (a test's curated dir may have none)."""
    path = curated_dir / "rank-choices.json"
    if not path.exists():
        return {}
    spells = json.loads(path.read_text(encoding="utf-8")).get("specs", {}).get(spec, [])
    return {spell["authored"]: _scoped_rows(spell) for spell in spells}


def _scoped_rows(spell: dict) -> list[dict]:
    """`spell`'s rows, each with the priority-list "lines" it applies to: a band
    naming its own "lines" or "prepull" (round 17) keeps its own scope, a band
    naming neither takes its spell's (sim/request scopeBands)."""
    return [
        {**row, "lines": _priority_lines(row if _names_scope(row) else spell)}
        for row in spell["bands"]
    ]


def _names_scope(entry: dict) -> bool:
    """Whether a spell or band names a "lines" or "prepull" scope."""
    return "lines" in entry or "prepull" in entry


def _priority_lines(entry: dict) -> list[int] | None:
    """The priority-list lines a spell's or band's scope reaches: its "lines",
    none when it scopes only pre-pull actions, None (every line) when it
    scopes nothing."""
    if entry.get("lines"):
        return entry["lines"]
    if entry.get("prepull"):
        return []
    return None


def _rank_of(spellranks: dict, class_slug: str) -> dict[int, int]:
    """spell id -> its rank number (0 for an unranked spell), for one class."""
    return {
        rank["id"]: rank.get("rank", 0)
        for ranks in spellranks.get("classes", {}).get(class_slug, {}).values()
        for rank in ranks
    }


def _resolve(
    families: dict[str, list[tuple[int, int]]],
    id_to_family: dict[int, str],
    rank_of: dict[int, int],
    spell_id: int,
    level: int,
) -> tuple[str, int] | None:
    """The (family name, id) this ability casts at `level`, or None when
    the curated id is not one spellranks.json tracks for this class (an
    engine-internal or item-proc id the ladder itself would call
    unresolved) or the family's first rank is not learned yet at `level`.

    Below 60 this is the sim's default rule (sim/internal/spellranks
    LearnedSpellID): the highest rank learned that is not above the list's
    own, so a deliberate lower rank is never lifted (owner ruling 3)."""
    family = id_to_family.get(spell_id)
    if family is None:
        return None
    named = [rank[0] for rank in families[family] if rank[1] == spell_id]
    cap = min(level, named[0]) if named else level
    learned = [rank for rank in families[family] if rank[0] <= cap]
    if not learned:
        return None
    # The site runs a band-60 list as written, deliberate lower ranks and
    # all (sim/request's rank rewrite applies below 60 only).
    if level >= MAX_LEVEL and any(rank[1] == spell_id for rank in learned):
        return family, spell_id
    # Two ids can share the top learned level (Water Shield's buff and its
    # proc; a Burning Crusade copy of Lesser Healing Wave's rank 6; a
    # Forever copy of Exorcism rank 1): keep the one the list names, else
    # the first in the list's own id band, as the sim's rewrite does
    # (sim/internal/spellranks sameBandOrFirst).
    # Two ranks can share a level too (Counterattack ranks 1 and 2 at 30):
    # the tier is the highest rank learned there.
    top_level = [rank[1] for rank in learned if rank[0] == learned[-1][0]]
    top_rank = max(rank_of.get(i, 0) for i in top_level)
    top = [i for i in top_level if rank_of.get(i, 0) == top_rank]
    if spell_id in top:
        return family, spell_id
    same_band = [i for i in top if (i >= FOREVER_ID_FLOOR) == (spell_id >= FOREVER_ID_FLOOR)]
    return family, (same_band or top)[0]


def _resolve_icon(spell_id: int, icon_of: Callable[[int], str] | None) -> str:
    """The resolved rank's icon, via `icon_of` -- a lookup a caller supplies
    (see `build_rotations` below, which looks one up in the committed
    `spellicons.json`, and `compute_spell_icons`, which resolves one fresh
    off the client's own tables).

    `icon_of` is `None` when `build_rotation` is called with no way to
    resolve an icon at all (every existing caller before this field
    existed, and a build whose committed `spellicons.json` does not exist
    yet -- see `_load_spell_icons`): the placeholder is what every other
    icon-bearing field in this pipeline emits for "nothing to resolve"
    (`pipeline.icons.resolve_icon`'s own doc), never an empty string,
    which would put `icons/.webp` in the output.
    """
    if icon_of is None:
        return PLACEHOLDER_ICON
    return icon_of(spell_id)


def build_rotation(
    spec: str,
    curated_dir: Path,
    spellranks: dict,
    icon_of: Callable[[int], str] | None = None,
    trainable_ids: frozenset[int] = frozenset(),
    talent_spells: Mapping[str, list[tuple[int, int]]] = MappingProxyType({}),
    rank_levels: Mapping[tuple[int, int], int] = MappingProxyType({}),
) -> list[AddonRotationBand]:
    """One spec's rotation table: a band at every LEVEL_BANDS rung and at
    every level between where a line changes (`_band_levels`).
    `trainable_ids` is the class's trainer-learned ids (`load_trainable_ids`),
    `talent_spells` its talents' own spells (`load_talent_spells`), and
    `rank_levels` the level the spec's published build takes each rank of each talent's
    spell (`load_spec_rank_levels`): a talent's row starts at its first point, and a row whose
    condition needs a rank known at that rank (`_talents_allow`)."""
    apl_path = curated_dir / "apl" / f"{spec}.json"
    if not apl_path.exists():
        raise AddonRotationError(f"missing {apl_path}")
    curated = json.loads(apl_path.read_text(encoding="utf-8"))
    panel_when = _panel_when(curated, apl_path)
    casts = _cast_lines(curated.get("rotation", {}), panel_when)
    _check_panel_when_steps(panel_when, casts, apl_path)

    families = {
        name: floor_pairs(pairs, first_point_levels(rank_levels))
        for name, pairs in _family_ranks(
            spellranks, _class_slug(spec), trainable_ids, talent_spells
        ).items()
    }
    id_to_family = _id_to_family(families)
    rank_of = _rank_of(spellranks, _class_slug(spec))
    choices = load_rank_choices(curated_dir, spec)

    context = _SpecContext(spec, spellranks, families, id_to_family, rank_of, choices, icon_of)
    bands: list[AddonRotationBand] = []
    for level in _band_levels():
        lines = [
            line
            for cast in casts
            if _level_allows(cast.get("condition"), level)
            and _talents_allow(cast.get("condition"), level, rank_levels)
            and _trained_allow(cast.get("condition"), level, families, id_to_family)
            and (line := _line_at(context, cast, level)) is not None
        ]
        if level in LEVEL_BANDS or not bands or lines != bands[-1].lines:
            bands.append(AddonRotationBand(level=level, lines=lines))
    return bands


def _band_levels() -> range:
    """Every level from the first rung to the last. A band is published at
    each rung and at every level between where a line changes (a rank row
    or a rank learned mid-band); the addon reads the band whose level is the
    highest at or below the player's (Rotation.lua bandFor)."""
    return range(LEVEL_BANDS[0], LEVEL_BANDS[-1] + 1)


@dataclass(frozen=True)
class _SpecContext:
    """One spec's rank tables, read once for every level's lines."""

    spec: str
    spellranks: dict
    families: dict[str, list[tuple[int, int]]]
    id_to_family: dict[int, str]
    rank_of: dict[int, int]
    choices: dict[int, list[dict]]
    icon_of: Callable[[int], str] | None


def _line_at(context: _SpecContext, cast: dict, level: int) -> AddonRotationLine | None:
    """`cast` as the band at `level` publishes it, or None when not learned yet."""
    resolved = _resolve(
        context.families, context.id_to_family, context.rank_of, cast["spell_id"], level
    )
    if resolved is None:
        return None
    name, resolved_id = resolved
    # A per-band rank choice (sim/request rankResolver) wins at every level,
    # 60 included under the raid preset (`_band_ranks`); a race's own row wins
    # for that race.
    ranks = _band_ranks(context.choices, cast["spell_id"], level, cast["line"])
    class_slug = _class_slug(context.spec)
    pick = _RankPick(context.spellranks, class_slug, name, cast["spell_id"], level)
    resolved_id = pick.id_of(ranks.get(None), resolved_id)
    return AddonRotationLine(
        spell_id=resolved_id,
        name=name,
        condition=_condition_of(cast, level, context.spec),
        icon=_resolve_icon(resolved_id, context.icon_of),
        race_spell_ids=_race_ids(pick, ranks, resolved_id),
    )


#: The client tables a rotation line's icon needs, straight off `raw/` --
#: the same five `pipeline.normalize.__init__` already reads to build a
#: talent's own icon (`load_spell_text` for `SpellMisc.SpellIconFileDataID`,
#: `icon_names` for `ManifestInterfaceData`). Read only by
#: `compute_spell_icons` below -- the one function in this module that
#: still touches `raw/` at all.
_ICON_RAW_TABLES = (
    "Spell.csv",
    "SpellMisc.csv",
    "SpellEffect.csv",
    "SpellDuration.csv",
    "ManifestInterfaceData.csv",
)


def _load_client_spell_tables(raw: Path) -> tuple[SpellText | None, dict[int, str] | None]:
    """`(spell_text, icon_names)` off `build/raw/`, or `(None, None)` when
    this build has no `raw/` at all, or is missing one of the five tables.
    `compute_spell_icons` is the only caller: it needs these to resolve a
    rotation line's icon from the client's own tables in the first place,
    at the one point in the pipeline (`python -m pipeline icons`) raw/ is
    guaranteed to actually be there."""
    if not raw.exists() or not all((raw / name).exists() for name in _ICON_RAW_TABLES):
        return None, None
    spell_text = load_spell_text(
        read_csv(raw / "Spell.csv"),
        read_csv(raw / "SpellMisc.csv"),
        read_csv(raw / "SpellEffect.csv"),
        read_csv(raw / "SpellDuration.csv"),
    )
    icons = icon_names(read_csv(raw / "ManifestInterfaceData.csv"))
    return spell_text, icons


def compute_spell_icons(
    root: Path, build: str, curated_dir: Path = Path("curated")
) -> dict[int, str]:
    """spell id -> icon name for every id a rotation line resolves to,
    across every curated spec and every `LEVEL_BANDS` entry, read straight
    off this build's `raw/` client tables.

    This is the one function in this module that still needs `raw/` at
    all: `pipeline.icons.icons_for_build` (the step that already requires
    `raw/`, to download icon art) calls this and persists the result at
    `builds/<build>/spellicons.json` (`write_spell_icons`); `build_rotations`
    below then reads only that committed file, never `raw/` itself, which
    is what lets `addon-data --check` (data.yml's `test` job) and bis.yml's
    nightly regen -- both run against a plain checkout with no `raw/`
    CSVs -- see the same real icons the fetch job saw instead of
    re-deriving nothing and publishing placeholders over them.

    A build whose `raw/` is missing or incomplete resolves every id to
    the placeholder (`_resolve_icon`'s own fallback) rather than refusing;
    `write_spell_icons` would then commit an all-placeholder file, which is
    a visible regression, not a silent one, and no worse than what this
    build already had before `raw/` was fetched at all.

    Sorted by id on return: small, deterministic, and a stable diff when a
    curated APL or spellranks.json changes which ids are referenced.

    Player-review sweep 15 (2026-09-30) flagged this join as wrong for
    Heroic Strike (all 9 ranks: 78, 284, 285, 1608, 11564-11567, 25286) and
    Rend (all 7 ranks: 772, 6546-6548, 11572-11574) -- both resolve to a
    Rogue ability's icon file (Ambush's 132282/ability_rogue_ambush, Gouge's
    132155/ability_gouge). Sweep 16 re-fetched SpellMisc.csv and
    ManifestInterfaceData.csv fresh from wago.tools for both this build and
    Classic Era 1.15.9.70003, and cross-checked Wowhead Classic's own
    rendered icondb id for the same spell ids: all four sources agree, on
    both ability's every rank. This is a real, longstanding Classic client
    quirk (these two Vanilla abilities' base SpellMisc rows have carried a
    borrowed Rogue icon file since at least 1.12), not a join bug -- see
    test_known_spells_pin_to_their_verified_client_icon, which pins it
    (tenet 8) so a future sweep does not re-file the same false positive.
    """
    spellranks_path = root / build / "spellranks.json"
    if not spellranks_path.exists():
        raise AddonRotationError(f"missing {spellranks_path}")
    spellranks = json.loads(spellranks_path.read_text(encoding="utf-8"))
    spell_text, icon_lookup = _load_client_spell_tables(root / build / "raw")
    icon_of = (
        (
            lambda spell_id: resolve_icon(
                spell_text.icon_file_id(spell_id), icon_lookup, f"rotation spell {spell_id}"
            )
        )
        if spell_text is not None and icon_lookup is not None
        else None
    )

    apl_dir = curated_dir / "apl"
    specs = sorted(path.stem for path in apl_dir.glob("*.json")) if apl_dir.exists() else []
    icons: dict[int, str] = {}
    for spec in specs:
        class_slug = _class_slug(spec)
        for band in build_rotation(
            spec,
            curated_dir,
            spellranks,
            icon_of,
            load_trainable_ids(root / build, class_slug),
            load_talent_spells(root / build, class_slug),
        ):
            for line in band.lines:
                icons.setdefault(line.spell_id, line.icon)
        # Every cast the list names as written, pre-pull included: the guide's opener and its
        # priority rows show the client's icon for each, even a spell no rank family resolves
        # (a pet summon, Prowl, a talent spell such as Divine Favor).
        apl_text = (apl_dir / f"{spec}.json").read_text(encoding="utf-8")
        rotation = json.loads(apl_text).get("rotation", {})
        for spell_id in sorted(apl_cast_spell_ids(rotation)):
            icons.setdefault(spell_id, _resolve_icon(spell_id, icon_of))
    # Spells a guide names that no rotation line does (the cooldown list: a trainer spell such as
    # Evocation): `curated/extra-spell-icons.json` lists their ids so they get an icon too.
    extra_path = curated_dir / "extra-spell-icons.json"
    if extra_path.exists():
        extra = json.loads(extra_path.read_text(encoding="utf-8"))
        for spell_id in extra["spell_ids"]:
            icons.setdefault(spell_id, _resolve_icon(spell_id, icon_of))
        # An engine id the client's tables do not carry takes its own spell's other rank.
        for spell_id, client_row in extra.get("icon_rows", {}).items():
            icons[int(spell_id)] = _resolve_icon(client_row, icon_of)
    # The raid cover grid's spells (curated/raid-cover.json): each row draws its buff's or
    # debuff's icon, and the seal a judgement comes from.
    for spell_id in sorted(raid_cover_spell_ids(curated_dir)):
        icons.setdefault(spell_id, _resolve_icon(spell_id, icon_of))
    # Spells the log coach names (its rule files' core casts, auras, downranks, defensives):
    # the review's titles and chips show each one's icon, every rank.
    coach_ids = coach_spell_ids(curated_dir) | coach_persona_spell_ids(curated_dir, spellranks)
    for spell_id in sorted(coach_ids):
        icons.setdefault(spell_id, _resolve_icon(spell_id, icon_of))
    # The form spells that draw a spec the client gives no tree of its own (Cat Form for Feral,
    # Dire Bear Form for Feral Bear): pipeline.specs reads their icons from this file.
    from pipeline.specs import load_specs

    specs = load_specs(curated_dir) if (curated_dir / "specs.json").exists() else []
    for form_spell in sorted(s.form_spell for s in specs if s.form_spell is not None):
        icons.setdefault(form_spell, _resolve_icon(form_spell, icon_of))
    # Every rank of a family any of the above names: a page that downranks to the rank a
    # character's level has (Lesser Heal rank 1, spell 2050, on a level 1 Discipline card)
    # draws that rank's own icon, never the placeholder.
    for spell_id in sorted(rank_siblings(spellranks, set(icons))):
        icons.setdefault(spell_id, _resolve_icon(spell_id, icon_of))
    return dict(sorted(icons.items()))


def rank_siblings(spellranks: dict[str, Any], ids: set[int]) -> set[int]:
    """Every rank id of each `spellranks.json` family that has a rank in `ids`, `ids` excluded."""
    siblings: set[int] = set()
    for families in spellranks.get("classes", {}).values():
        for ranks in families.values():
            family = {row["id"] for row in ranks}
            if family & ids:
                siblings |= family - ids
    return siblings


def _collect_spell_ids(value: object, into: set[int]) -> None:
    """Every `spell_id` and `spell_ids` entry anywhere under `value`."""
    if isinstance(value, dict):
        for key, inner in value.items():
            if key == "spell_id" and isinstance(inner, int) and not isinstance(inner, bool):
                into.add(inner)
            elif key == "spell_ids" and isinstance(inner, list):
                into.update(i for i in inner if isinstance(i, int) and not isinstance(i, bool))
            else:
                _collect_spell_ids(inner, into)
    elif isinstance(value, list):
        for inner in value:
            _collect_spell_ids(inner, into)


def raid_cover_spell_ids(curated_dir: Path = Path("curated")) -> set[int]:
    """Every spell id `curated/raid-cover.json` names: each row's `spell_id` and each
    provider's `spell` (a judgement's seal, Demoralizing Roar). Empty without the file."""
    path = curated_dir / "raid-cover.json"
    if not path.exists():
        return set()
    ids: set[int] = set()
    for entry in json.loads(path.read_text(encoding="utf-8"))["entries"]:
        ids.add(entry["spell_id"])
        ids.update(p["spell"] for p in entry["providers"] if "spell" in p)
    return ids


def coach_spell_ids(curated_dir: Path = Path("curated")) -> set[int]:
    """Every spell id a log coach rule file (`curated/coach/*.json`) carries, every rank of
    every family, wherever the file holds one (`spell_id` or `spell_ids`). The coach's
    review names these spells with their icons, and no rotation line need name them
    (Greater Heal, Bloodrage, Inner Fire), so `compute_spell_icons` adds them."""
    ids: set[int] = set()
    for path in sorted((curated_dir / "coach").glob("*.json")):
        _collect_spell_ids(json.loads(path.read_text(encoding="utf-8")), ids)
    return ids


def coach_persona_spell_ids(curated_dir: Path, spellranks: dict) -> set[int]:
    """Every rank of every class spell a log coach persona (`curated/coach/<spec>.md`) names
    by its family name: the heals the F5 rank finding reads (Greater Heal) appear in no rule
    file, only in the persona and the player's log. A family name matches as whole words,
    case and all; a stray match only adds an icon."""
    ids: set[int] = set()
    for path in sorted((curated_dir / "coach").glob("*.md")):
        text = path.read_text(encoding="utf-8")
        for name, ranks in _family_ranks(spellranks, _class_slug(path.stem)).items():
            if re.search(rf"(?<![\w']){re.escape(name)}(?![\w'])", text):
                ids.update(spell_id for _, spell_id in ranks)
    return ids


def write_spell_icons(
    build: str, root: Path = Path("builds"), curated_dir: Path = Path("curated")
) -> Path:
    """Write `builds/<build>/spellicons.json`: `compute_spell_icons`'
    id -> icon map, committed like every other per-build JSON file
    (`spellranks.json` beside it) so a later run needs no `raw/` at all to
    read it back. Called by `python -m pipeline icons`, the step that
    already has `raw/` -- see `compute_spell_icons`'s own doc."""
    icons = compute_spell_icons(root, build, curated_dir)
    path = root / build / "spellicons.json"
    payload = {str(spell_id): icon for spell_id, icon in icons.items()}
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    return path


def load_talent_spells(build_dir: Path, class_slug: str) -> dict[str, list[tuple[int, int]]]:
    """Talent name -> [(the level its tier opens, its spell id)] for every
    talent of `class_slug` on this build (`builds/<build>/talents/<class>.json`),
    or none when the build has no talents file for the class."""
    path = build_dir / "talents" / f"{class_slug}.json"
    if not path.exists():
        return {}
    trees = json.loads(path.read_text(encoding="utf-8"))["trees"]
    return {
        talent["name"]: [(tier_level(talent["tier"]), talent["spell_id"])]
        for tree in trees
        for talent in tree["talents"]
        if isinstance(talent.get("spell_id"), int)
    }


def load_trainable_ids(build_dir: Path, class_slug: str) -> frozenset[int]:
    """Every spell id `class_slug` learns from a trainer on this build
    (`builds/<build>/trainables/<class>.json`, pipeline.trainables), or none
    when the build has no trainables file for the class."""
    path = build_dir / "trainables" / f"{class_slug}.json"
    if not path.exists():
        return frozenset()
    entries = json.loads(path.read_text(encoding="utf-8"))["trainables"]
    return frozenset(rank["id"] for entry in entries for rank in entry["ranks"])


def _load_spell_icons(build_dir: Path) -> dict[int, str] | None:
    """spell id -> icon name off the committed `build_dir/spellicons.json`
    (`write_spell_icons`), or `None` when this build predates that file.
    `build_rotations` is the only caller, and never falls back to `raw/`
    itself -- see `compute_spell_icons`'s own doc for why."""
    path = build_dir / "spellicons.json"
    if not path.exists():
        return None
    payload = json.loads(path.read_text(encoding="utf-8"))
    return {int(spell_id): icon for spell_id, icon in payload.items()}


def build_rotations(
    root: Path, build: str, curated_dir: Path = Path("curated")
) -> dict[str, list[AddonRotationBand]]:
    """Every curated spec's rotation table, keyed by spec slug."""
    spellranks_path = root / build / "spellranks.json"
    if not spellranks_path.exists():
        raise AddonRotationError(f"missing {spellranks_path}")
    spellranks = json.loads(spellranks_path.read_text(encoding="utf-8"))
    spell_icons = _load_spell_icons(root / build)
    icon_of = (
        (lambda spell_id: spell_icons.get(spell_id, PLACEHOLDER_ICON))
        if spell_icons is not None
        else None
    )
    apl_dir = curated_dir / "apl"
    specs = sorted(path.stem for path in apl_dir.glob("*.json")) if apl_dir.exists() else []
    return {
        spec: build_rotation(
            spec,
            curated_dir,
            spellranks,
            icon_of,
            load_trainable_ids(root / build, _class_slug(spec)),
            load_talent_spells(root / build, _class_slug(spec)),
            load_spec_rank_levels(
                root / build,
                spec,
                _class_slug(spec),
                load_talent_owned_auras(curated_dir, _class_slug(spec)),
            ),
        )
        for spec in specs
    }


def build_ability_levels(
    root: Path, build: str, curated_dir: Path = Path("curated")
) -> dict[int, int]:
    """Every rotation line's spell id -> the level its ability's FIRST rank
    is learned (Data.lua's `abilityLevels`, contract 1): the rotation
    region's "at level N". Keyed by each id a band actually carries, so the
    addon looks up the very id the rotation line hands it; every rank of
    one ability maps to the same level (Execute's 5308 through 20662 all
    say 24). The first rank is `_family_ranks`' own first entry, level-0
    alternate ids already set aside there, so a family's level is the one
    `_resolve` uses to decide when the ability joins a band."""
    spellranks_path = root / build / "spellranks.json"
    if not spellranks_path.exists():
        raise AddonRotationError(f"missing {spellranks_path}")
    spellranks = json.loads(spellranks_path.read_text(encoding="utf-8"))
    apl_dir = curated_dir / "apl"
    specs = sorted(path.stem for path in apl_dir.glob("*.json")) if apl_dir.exists() else []
    levels: dict[int, int] = {}
    for spec in specs:
        trainable_ids = load_trainable_ids(root / build, _class_slug(spec))
        talent_spells = load_talent_spells(root / build, _class_slug(spec))
        families = _family_ranks(spellranks, _class_slug(spec), trainable_ids, talent_spells)
        id_to_family = _id_to_family(families)
        for band in build_rotation(
            spec, curated_dir, spellranks, trainable_ids=trainable_ids, talent_spells=talent_spells
        ):
            for line in band.lines:
                level = families[id_to_family[line.spell_id]][0][0]
                if levels.setdefault(line.spell_id, level) != level:
                    raise AddonRotationError(
                        f"spell {line.spell_id} ({line.name}) is learned at "
                        f"{levels[line.spell_id]} in one spec and {level} in {spec}"
                    )
    return dict(sorted(levels.items()))
