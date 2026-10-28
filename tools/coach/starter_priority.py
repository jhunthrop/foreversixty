"""The priority list of a starter: every button the rotation data or the guide's rotation names.

Pure functions over data the caller loads. Each entry is (ability name, line). A line is a reason in a
sentence when one can be found in the note or the guide, else the bare name.
"""

from __future__ import annotations

import re
from typing import Any

from starter_text import digit_tokens, plain_sentence, sentences, usable_sentence

NAME_NEAR_FRONT_CHARS = 40
MIN_ABILITY_CHARS = 5

# Claims the rotation data makes that the guides do not, so a starter does not repeat them.
CLAIMS_THE_GUIDE_DOES_NOT_MAKE = re.compile(
    r"most healing per mana|lifts Rupture damage|more damage from your Rupture|\bMaul\b.*\brage\b",
    re.IGNORECASE,
)
# A sentence that rules an ability out or ranks it lower is not a mention of a core ability.
NEGATION = re.compile(
    r"\b(not|never|instead|rather|than|worse|lower|loss|cannot|can't|without|swap\w*|skip\w*|isn't|aren't)\b",
    re.IGNORECASE,
)

Entry = tuple[str, str]


def leads_with(name: str, sentence: str) -> bool:
    """The sentence is about the ability: it starts as a sentence and names it near the front."""
    return sentence[:1].isupper() and 0 <= sentence.lower().find(name.lower()) <= NAME_NEAR_FRONT_CHARS


def mentions(name: str, sentence: str) -> bool:
    return name.lower() in sentence.lower()


def corroborated_note(note: str, guide_digits: set[str]) -> str | None:
    """The note's lead sentence in words, if it reads alone and the guide backs every number and claim."""
    lead = sentences(note)[0] if note.strip() else ""
    if not lead or CLAIMS_THE_GUIDE_DOES_NOT_MAKE.search(lead) or not digit_tokens(lead) <= guide_digits:
        return None
    return usable_sentence(note)


def cast_spell_id(item: dict[str, Any]) -> int | None:
    cast = item["action"].get("castSpell")
    return cast.get("spellId", {}).get("spellId") if cast else None


def guide_abilities(prose: str, castable: set[str]) -> list[str]:
    """Castable abilities the guide's rotation names outside a sentence that rules them out, in its order."""
    names = [n for n in castable if len(n) >= MIN_ABILITY_CHARS]
    if not names:
        return []
    pattern = re.compile(r"\b(" + "|".join(re.escape(n) for n in sorted(names, key=len, reverse=True)) + r")\b")
    found: list[str] = []
    for sentence in sentences(prose):
        if NEGATION.search(sentence):
            continue
        found.extend(n for n in pattern.findall(sentence) if n not in found)
    return found


def positive_plain_sentences(prose: str) -> list[str]:
    plain = (plain_sentence(s) for s in sentences(prose) if not NEGATION.search(s))
    return [s for s in plain if s]


def apl_entries(
    items: list[dict[str, Any]], names: dict[int, str], guide_digits: set[str]
) -> list[tuple[str, str | None]]:
    """(name, note sentence that leads with it or None) per cast line; lines with no known spell keep only a sentence."""
    entries: list[tuple[str, str | None]] = []
    for item in items:
        sentence = corroborated_note(item.get("notes", ""), guide_digits)
        spell_id = cast_spell_id(item)
        name = names.get(spell_id) if spell_id is not None else None
        if name is None:
            if sentence:
                entries.append((sentence, sentence))
        else:
            entries.append((name, sentence if sentence and leads_with(name, sentence) else None))
    return entries


def guide_sentence_for(name: str, sentences_: list[str], used: set[str]) -> str | None:
    unused = [s for s in sentences_ if s not in used]
    return next((s for s in unused if leads_with(name, s)), None) or next(
        (s for s in unused if mentions(name, s)), None
    )


def priority_entries(
    items: list[dict[str, Any]],
    names: dict[int, str],
    castable: set[str],
    prose: str,
    guide_digits: set[str],
) -> list[Entry]:
    """APL buttons in their order, then the guide's other core abilities in the guide's order."""
    ordered: list[tuple[str, str | None]] = []
    seen: set[str] = set()
    for name, sentence in apl_entries(items, names, guide_digits):
        if name.lower() not in seen:
            seen.add(name.lower())
            ordered.append((name, sentence))
    ordered.extend((n, None) for n in guide_abilities(prose, castable) if n.lower() not in seen)
    guide = positive_plain_sentences(prose)
    used = {s for _, s in ordered if s}
    result: list[Entry] = []
    for name, sentence in ordered:
        line = sentence or guide_sentence_for(name, guide, used) or f"{name}."
        used.add(line)
        result.append((name, line if leads_with(name, line) or is_bare((name, line)) else f"{name}: {line}"))
    return result


def is_bare(entry: Entry) -> bool:
    return entry[1] == f"{entry[0]}."
