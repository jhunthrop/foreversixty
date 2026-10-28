"""Turning rotation notes and guide prose into plain sentences a starter persona can carry.

Rules: only a sentence's own words are used (a later sentence depends on the first, so it is
never promoted); numbers become words; simulator talk, dev notes and dangling references drop
the sentence.
"""

from __future__ import annotations

import re

MIN_SENTENCE_CHARS = 20

SMALL = (
    "zero one two three four five six seven eight nine ten eleven twelve thirteen fourteen "
    "fifteen sixteen seventeen eighteen nineteen"
).split()
TENS = "_ _ twenty thirty forty fifty sixty seventy eighty ninety".split()

# Sentences about the simulator, the rule set or the author's own working, not about playing the spec.
NOT_ABOUT_PLAYING = re.compile(
    r"\b(engine|ladder|rung|golden|harness|inert|unresolved|gated|sim|simulator|simulated|simulation|"
    r"the search|conditions?|preset|cited|tried|measured|the client states?|fight clock|falls? through|"
    r"bare cast|era fallback|idle by design|guard|the rule|therefore|same as|this site|this build|"
    r"measures|our|raid-ready|search run|share of the fight)\b",
    re.IGNORECASE,
)
# Changelog and rotation-data wording (P2): a sentence about what changed or how a line is tracked.
CHANGELOG = re.compile(
    r"\b(this week|now the only|back in the rotation|tracked by slot|second [\w ]{0,20}line|fires itself|"
    r"the new|new baseline|new to the tree|fraction of a percent|whole of its gain|no longer)\b",
    re.IGNORECASE,
)
# References to other lines of the rotation data, which a lone sentence cannot carry.
DANGLING_REFERENCE = re.compile(
    r"\b(above|below it|same pace|paced?|paced like|the pass|the line)\b", re.IGNORECASE
)
PARENTHETICAL = re.compile(r"\s*\([^)]*\)")
MARKDOWN_EMPHASIS = re.compile(r"\*+")
SECONDS_SUFFIX = re.compile(r"(?<![\w.])(\d{1,3})\s?s\b(?!\w)")
NUMBER = re.compile(r"(?<![\w.])(\d{1,3})(%?)(?!\w|\.\d|\d)")
# Words about the site's own tooling that a player would not say, and their plain replacements.
TOOL_SPEAK = {"The curated rotation": "The rotation", ", hence the reserve floor": ", so keep a mana reserve"}
PERCENT_WORD = " percent"
DIGIT = re.compile(r"\d")
HALF = re.compile(r"(?<![\w.])(\d{1,2})\.5(?![\w.]|\d)")
# A figure: a number with its unit (seconds, percent, a multiplier).
FIGURE = re.compile(r"(\d+(?:\.\d+)?)\s?-?(%|percent\b|s\b|sec\b|seconds?\b|x\b)", re.IGNORECASE)
UNIT_KIND = {"%": "percent", "percent": "percent", "x": "times"}
SENTENCE_END = re.compile(r"(?<=[.!?])\s+")
# A colon (not the one in "Shadow Word: Pain"), a semicolon or a dash that can end a sentence instead.
SENTENCE_JOIN = re.compile(r"(?<!Word): |; | -- | — ")
MAX_SENTENCE_WORDS = 20
# Words that open a clause able to stand as its own sentence.
STANDALONE_OPENERS = frozenset({"it", "its", "the", "once", "then", "recast", "health", "no", "they", "this"})


def number_words(n: int) -> str:
    if n < 20:
        return SMALL[n]
    if n < 100:
        return TENS[n // 10] + (f"-{SMALL[n % 10]}" if n % 10 else "")
    rest = f" {number_words(n % 100)}" if n % 100 else ""
    return f"{SMALL[n // 100]} hundred{rest}"


def spell_out_numbers(text: str) -> str:
    """Replace whole integers up to 999 and halves (a percent sign, a bare 's' for seconds) with words."""
    text = SECONDS_SUFFIX.sub(lambda m: f"{m[1]} seconds", text)
    text = HALF.sub(lambda m: f"{number_words(int(m[1]))} and a half", text)
    words = NUMBER.sub(lambda m: number_words(int(m[1])) + (PERCENT_WORD if m[2] else ""), text)
    return words.replace("seconds cooldown", "second cooldown")


def sentences(text: str) -> list[str]:
    return SENTENCE_END.split(MARKDOWN_EMPHASIS.sub("", text).strip())


def plain_words(text: str) -> str:
    for tool, plain in TOOL_SPEAK.items():
        text = text.replace(tool, plain)
    return text


def plain_sentence(sentence: str) -> str | None:
    """The sentence in words, or None when it cannot stand alone."""
    words = spell_out_numbers(plain_words(PARENTHETICAL.sub("", sentence).replace("!", ".").strip()))
    if DIGIT.search(words) or any(p.search(words) for p in (NOT_ABOUT_PLAYING, CHANGELOG, DANGLING_REFERENCE)):
        return None
    if len(words) < MIN_SENTENCE_CHARS:
        return None
    return words if words.endswith((".", ":")) else words + "."


def usable_sentence(note: str) -> str | None:
    """A note's first sentence in words, or None. Later sentences depend on it, so are never used."""
    return plain_sentence(sentences(note)[0])


def digit_tokens(text: str) -> set[str]:
    return set(re.findall(r"\d+", text))


def figures(text: str) -> set[tuple[str, str]]:
    """(number, unit kind) for every figure in the text, so "15 s" and "15-second" match."""
    return {(m[1], UNIT_KIND.get(m[2].lower(), "seconds")) for m in FIGURE.finditer(text)}


def short_sentences(line: str, name: str) -> str:
    """Sentences longer than the limit split where the next part reads alone (read 4, ruling 3)."""
    prefix = f"{name}: " if line.startswith(f"{name}: ") else ""
    return prefix + " ".join(split_long(part) for part in SENTENCE_END.split(line.removeprefix(prefix)))


def split_long(sentence: str) -> str:
    if len(sentence.split()) <= MAX_SENTENCE_WORDS:
        return sentence
    pieces = SENTENCE_JOIN.split(sentence.rstrip("."))
    joined = pieces[0]
    for piece, separator in zip(pieces[1:], SENTENCE_JOIN.findall(sentence), strict=True):
        reads_alone = piece.split()[0].lower() in STANDALONE_OPENERS
        joined += f". {piece[:1].upper()}{piece[1:]}" if reads_alone else f"{separator}{piece}"
    return joined + "."
