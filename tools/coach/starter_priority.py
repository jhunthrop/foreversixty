"""The priority list of a starter: the buttons the guide's rotation section names positively, in its order.

Pure functions over data the caller loads. Each entry is (ability name, line). The line is a reason
sentence that is about the ability (it names it first, or is prefixed with it), taken from the
rotation data's note or the guide; else the bare name. Each ability appears once, and a guide
sentence serves only the first ability it names, so no sentence repeats.
"""

from __future__ import annotations

import re
from collections.abc import Collection
from typing import Any

from starter_guide import (
    clauses,
    excluded_abilities,
    first_ability,
    guide_order,
    guide_sentences,
    name_pattern,
    negated_terms,
)
from starter_text import digit_tokens, figures, plain_sentence, sentences, usable_sentence

NAME_NEAR_FRONT_CHARS = 40
MIN_ABILITY_CHARS = 3
ARTICLES = frozenset({"The", "A", "An"})
SEMICOLON = "; "

# Claims the rotation data makes that the guides do not, so a starter does not repeat them.
CLAIMS_THE_GUIDE_DOES_NOT_MAKE = re.compile(
    r"most healing per mana|lifts Rupture damage|more damage from your Rupture|\bMaul\b.*\brage\b|"
    r"biggest threat hit|three of the four|one more Judgement a minute|the two curses",
    re.IGNORECASE,
)

Entry = tuple[str, str]


def leads_with(name: str, sentence: str) -> bool:
    """The sentence is about the ability: it starts as a sentence and names it near the front."""
    return sentence[:1].isupper() and 0 <= sentence.find(name) <= NAME_NEAR_FRONT_CHARS


def cites_negated(sentence: str, negated: Collection[str]) -> bool:
    return any(re.search(rf"\b{re.escape(term)}\b", sentence) for term in negated)


def corroborated_note(note: str, guide: str, negated: frozenset[str] = frozenset()) -> str | None:
    """The note's lead sentence in words, if it reads alone and the guide states every number it carries (P7)."""
    lead = sentences(note)[0] if note.strip() else ""
    if not lead or CLAIMS_THE_GUIDE_DOES_NOT_MAKE.search(lead) or cites_negated(lead, negated):
        return None
    if not digit_tokens(lead) <= digit_tokens(guide) or not figures(lead) <= figures(guide):
        return None
    return usable_sentence(note)


def cast_spell_id(item: dict[str, Any]) -> int | None:
    cast = item["action"].get("castSpell")
    return cast.get("spellId", {}).get("spellId") if cast else None


def note_reasons(
    items: list[dict[str, Any]], names: dict[int, str], guide: str, negated: frozenset[str]
) -> dict[str, str]:
    """The first rotation-data note per ability that leads with its name and passes every check."""
    reasons: dict[str, str] = {}
    for item in items:
        spell_id = cast_spell_id(item)
        name = names.get(spell_id) if spell_id is not None else None
        sentence = corroborated_note(item.get("notes", ""), guide, negated)
        usable = sentence and leads_with(name or "", sentence) and not is_orphan(name or "", sentence)
        if name and usable and name not in reasons:
            reasons[name] = sentence
    return reasons


def is_orphan(name: str, sentence: str) -> bool:
    """A noun phrase with no verb, such as "The instant Healing Wave." (P3)."""
    return sentence.split()[0] in ARTICLES and sentence.rstrip(".").endswith(name)


def sentence_parts(sentence: str, pattern: re.Pattern[str]) -> list[list[str]]:
    """Semicolon parts that each lead with an ability stand alone; else the sentence, then its first clause."""
    parts = [p for p in sentence.split(SEMICOLON) if (m := pattern.search(p)) and leads_with(m[1], p)]
    if len(parts) > 1:
        return [[p] for p in parts]
    return [[sentence, clauses(sentence)[0]]]


def guide_candidates(prose: str, names: set[str], negated: frozenset[str]) -> list[tuple[str, str]]:
    """(first ability named, plain sentence) for each usable guide sentence or part of one."""
    pattern = name_pattern(names)
    found: list[tuple[str, str]] = []
    for sentence in guide_sentences(prose, names):
        for texts in sentence_parts(sentence, pattern):
            usable = (p for p in map(plain_sentence, texts) if p and not cites_negated(p, negated))
            plain = next(usable, None)
            subject = first_ability(plain, pattern) if plain else None
            if plain and subject and not is_orphan(subject, plain):
                found.append((subject, plain))
    return found


def with_name(name: str, sentence: str) -> str:
    return sentence if leads_with(name, sentence) else f"{name}: {sentence}"


def priority_entries(
    items: list[dict[str, Any]],
    names: dict[int, str],
    castable: set[str],
    prose: str,
    guide: str,
) -> list[Entry]:
    """Guide-named buttons in the guide's order, each with a note or guide reason or bare (P1, P3, P4)."""
    cast_in_rotation = {names[i] for i in map(cast_spell_id, items) if i is not None and i in names}
    candidates_names = {n for n in castable | cast_in_rotation if len(n) >= MIN_ABILITY_CHARS}
    negated = frozenset(negated_terms(guide) | excluded_abilities(prose, candidates_names))
    notes = note_reasons(items, names, guide, negated)
    guide_reasons = guide_candidates(prose, candidates_names, negated)
    result: list[Entry] = []
    for name in guide_order(prose, candidates_names):
        line = notes.get(name) or next((s for subject, s in guide_reasons if subject == name), None)
        result.append((name, with_name(name, line) if line else f"{name}."))
    return result


def is_bare(entry: Entry) -> bool:
    return entry[1] == f"{entry[0]}."
