"""What a spec guide's rotation section says about each ability: which it names positively, in what order.

Rules (round 4 rulings P1, P4, P6):
- an ability is listed only when the rotation section names it positively: not in a clause about a
  loss or an older way to play, not as the far side of a comparison, not as another spell's effect;
- an ability the guide rules out (is no longer a, not part of the loop, left out, there is no, with
  no, replaces, its own guide) or conditions (only with a totem, without a talent) is never listed;
- order is the first clause the ability leads, else its first positive mention;
- a term the guide says the build no longer takes is never cited.
"""

from __future__ import annotations

import re

from starter_text import sentences

LIST_MARKER = re.compile(r"^\s*\d+\.\s+", re.MULTILINE)
BOLD_HEADING = re.compile(r"^\*\*([^*]+)\*\*:\s*(\w)", re.MULTILINE)
# A clause ends at a colon (not the one inside "Shadow Word: Pain"), a semicolon or a dash.
CLAUSE_BREAK = re.compile(r"(?<!Word): |; | — | -- ")
NAME_NEAR_FRONT_CHARS = 40

# Simulator talk: an ability counts here only when the clause is about it (it leads the clause).
SIM_CLAUSE = re.compile(r"\b(simulator|simulated|the search|search run|engine)\b", re.IGNORECASE)
# A loss, a measurement or an older way to play: no ability in the clause counts.
LOSS_CLAUSE = re.compile(
    r"\b(measured|measures|worse|lower|a loss|older community|earlier version)\b", re.IGNORECASE
)
# Words just before a name that make the mention a comparison or another spell's effect.
NOT_POSITIVE_BEFORE = re.compile(
    r"(\b(than|instead of|compete with|competes with|unlike|after an?|over|beats?|against|edged out)"
    r"\s+(?:\w+\s+)?|'s\s+(?:own\s+)?)$",
    re.IGNORECASE,
)
# Rule-outs whose subject is every ability named before the phrase in its clause.
RULED_OUT_ALL_BEFORE = re.compile(
    r"\b(not part of|not in the loop|do nothing|does nothing|is not cast|are not cast)\b", re.IGNORECASE
)
# Rule-outs whose subject is the first ability of the clause.
RULED_OUT_SUBJECT = re.compile(
    r"\b(is no longer an?|no upgrade|is not an?|has its own guide|removed|left out of|never be cast)\b",
    re.IGNORECASE,
)
# Rule-outs that name the ability right after the phrase.
RULED_OUT_NEXT = re.compile(r"\b(?:there is no|with no|leave)\s+(.*)$", re.IGNORECASE)
# Rule-outs that cover every ability after the phrase in the sentence.
RULED_OUT_ALL_AFTER = re.compile(r"\bleft out on purpose:?\s+(.*)$", re.IGNORECASE)
REPLACES = re.compile(r"\breplac(?:es|ing)\s+(\w+)", re.IGNORECASE)
# Group or talent conditions: an ability named outside the condition itself is conditional.
CONDITION = re.compile(
    r"\b(only in a raid|only with|with an? [\w' ]+? totem in the group|without the [\w' ]+? talent)\b",
    re.IGNORECASE,
)
# A talent or effect the guide says the build does not take (P6).
NEGATED_TERM = re.compile(r"\b(?:no longer takes|there is no)\s+((?:[A-Z][\w-]*(?:'s)?\s?)+)")


def prepared(prose: str, names: set[str]) -> str:
    """Rotation prose with list markers removed and bold headings that are not abilities dropped."""

    def heading(match: re.Match[str]) -> str:
        return match[0] if match[1] in names else match[2].upper()

    return BOLD_HEADING.sub(heading, LIST_MARKER.sub("", prose))


def clauses(sentence: str) -> list[str]:
    return CLAUSE_BREAK.split(sentence)


def name_pattern(names: set[str]) -> re.Pattern[str]:
    """Whole ability names, longest first; a plural counts, a hyphenated compound does not."""
    ordered = sorted(names, key=len, reverse=True)
    return re.compile(r"\b(" + "|".join(re.escape(n) for n in ordered) + r")s?\b(?!-)")


def first_ability(text: str, pattern: re.Pattern[str]) -> str | None:
    match = pattern.search(text)
    return match[1] if match else None


def leads(match: re.Match[str]) -> bool:
    return match.start() <= NAME_NEAR_FRONT_CHARS


def ruled_out_in(clause: str, pattern: re.Pattern[str], names: set[str]) -> set[str]:
    out: set[str] = set()
    all_before = RULED_OUT_ALL_BEFORE.search(clause)
    if all_before:
        out.update(pattern.findall(clause[: all_before.start()]))
    subject = RULED_OUT_SUBJECT.search(clause)
    if subject:
        out.update(filter(None, [first_ability(clause[: subject.start()], pattern)]))
    following = RULED_OUT_NEXT.search(clause)
    named_next = pattern.match(following[1]) if following else None
    if named_next:
        out.add(named_next[1])
    for word in REPLACES.findall(clause):
        out.update(n for n in names if n.split()[0] == word)
    return out


def conditioned_in(sentence: str, pattern: re.Pattern[str]) -> set[str]:
    conditions = list(CONDITION.finditer(sentence))
    outside = sentence
    for match in reversed(conditions):
        outside = outside[: match.start()] + " " + outside[match.end():]
    return set(pattern.findall(outside)) if conditions else set()


def excluded_in(sentence: str, pattern: re.Pattern[str], names: set[str]) -> set[str]:
    out = conditioned_in(sentence, pattern)
    all_after = RULED_OUT_ALL_AFTER.search(sentence)
    if all_after:
        out |= set(pattern.findall(all_after[1]))
    for clause in clauses(sentence):
        out |= ruled_out_in(clause, pattern, names)
    return out


def positive_mentions(clause: str, pattern: re.Pattern[str]) -> list[re.Match[str]]:
    if LOSS_CLAUSE.search(clause):
        return []
    found = [m for m in pattern.finditer(clause) if not NOT_POSITIVE_BEFORE.search(clause[: m.start()])]
    return [m for m in found if leads(m)] if SIM_CLAUSE.search(clause) else found


def guide_sentences(prose: str, names: set[str]) -> list[str]:
    return [s for s in sentences(prepared(prose, names)) if s.strip()]


def excluded_abilities(prose: str, names: set[str]) -> set[str]:
    """Abilities the rotation section rules out or makes conditional anywhere (P1, P6)."""
    if not names:
        return set()
    pattern = name_pattern(names)
    return set().union(*(excluded_in(s, pattern, names) for s in guide_sentences(prose, names)))


def guide_order(prose: str, names: set[str]) -> list[str]:
    """Abilities the rotation section names positively and never rules out, in the guide's order (P1, P4)."""
    if not names:
        return []
    pattern = name_pattern(names)
    all_clauses = [c for s in guide_sentences(prose, names) for c in clauses(s)]
    excluded = excluded_abilities(prose, names)
    led: dict[str, tuple[int, int]] = {}
    mentioned: dict[str, tuple[int, int]] = {}
    for index, clause in enumerate(all_clauses):
        first = pattern.search(clause)
        for match in positive_mentions(clause, pattern):
            mentioned.setdefault(match[1], (index, match.start()))
            if first is not None and match.start() == first.start() and leads(match):
                led.setdefault(match[1], (index, match.start()))
    keys = {name: led.get(name, at) for name, at in mentioned.items() if name not in excluded}
    return sorted(keys, key=lambda name: keys[name])


def negated_terms(guide: str) -> set[str]:
    """Talents and buttons the guide says the build no longer takes or no longer has (P6)."""
    return {m[1].strip().removesuffix("'s") for m in NEGATED_TERM.finditer(guide.replace("**", ""))}
