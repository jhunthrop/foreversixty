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
    r"bare cast|era fallback|idle by design|guard|the rule|therefore|same as|this site|this build)\b",
    re.IGNORECASE,
)
# References to other lines of the rotation data, which a lone sentence cannot carry.
DANGLING_REFERENCE = re.compile(
    r"\b(above|below it|same pace|paced?|paced like|the pass|the line)\b", re.IGNORECASE
)
PARENTHETICAL = re.compile(r"\s*\([^)]*\)")
MARKDOWN_EMPHASIS = re.compile(r"\*+")
SECONDS_SUFFIX = re.compile(r"(?<![\w.])(\d{1,3})\s?s\b(?!\w)")
NUMBER = re.compile(r"(?<![\w.])(\d{1,3})(%?)(?![\w.]|\d)")
PERCENT_WORD = " percent"
DIGIT = re.compile(r"\d")
SENTENCE_END = re.compile(r"(?<=[.!?])\s+")


def number_words(n: int) -> str:
    if n < 20:
        return SMALL[n]
    if n < 100:
        return TENS[n // 10] + (f"-{SMALL[n % 10]}" if n % 10 else "")
    rest = f" {number_words(n % 100)}" if n % 100 else ""
    return f"{SMALL[n // 100]} hundred{rest}"


def spell_out_numbers(text: str) -> str:
    """Replace whole integers up to 999 (a trailing percent sign, a bare 's' for seconds) with words."""
    text = SECONDS_SUFFIX.sub(lambda m: f"{m[1]} seconds", text)
    words = NUMBER.sub(lambda m: number_words(int(m[1])) + (PERCENT_WORD if m[2] else ""), text)
    return words.replace("seconds cooldown", "second cooldown")


def sentences(text: str) -> list[str]:
    return SENTENCE_END.split(MARKDOWN_EMPHASIS.sub("", text).strip())


def plain_sentence(sentence: str) -> str | None:
    """The sentence in words, or None when it cannot stand alone."""
    words = spell_out_numbers(PARENTHETICAL.sub("", sentence).replace("!", ".").strip())
    if DIGIT.search(words) or NOT_ABOUT_PLAYING.search(words) or DANGLING_REFERENCE.search(words):
        return None
    if len(words) < MIN_SENTENCE_CHARS:
        return None
    return words if words.endswith((".", ":")) else words + "."


def usable_sentence(note: str) -> str | None:
    """A note's first sentence in words, or None. Later sentences depend on it, so are never used."""
    return plain_sentence(sentences(note)[0])


def digit_tokens(text: str) -> set[str]:
    return set(re.findall(r"\d+", text))
