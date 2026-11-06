"""A card note never names a talent or spell at a band before the player has it.

Card notes are the `notes` strings of curated/apl/<spec>.json's priority steps (and their
`level_notes`, from their `from_level`). For every sentence of every note, each whole
talent name (from the class's talent trees) and trained spell name (from its trainables)
it mentions is checked at every band the row renders on in the active build's
addon-data.json: the player has a talent from its first point in the spec's published
leveling order, or in its raid order (`pipeline.talent_levels`: point i is spent at level
10 + i), and a trained spell from its lowest trainer rank.

A name is scoped, and passes, when its clause (the sentence split at , ; : and dashes)
says "from N", "at N", "by N", "until N" or "before N" (N a level), "once you have", or
"with the raid build" / "in the raid build"; when a conditional ("with", "without", "once",
"until", "before", "after") comes before it in its clause; or when a clause of the sentence
is nothing but a level ("From 25, ..."). Anything else that is right as written goes in
fixtures/card-note-levels-allowlist.txt, one line per exception with its reason:

    <spec> <name> <sentence hash>  <reason>

where the hash is the first 12 hex of the sentence's sha256, as the failure prints it.
"""

from __future__ import annotations

import hashlib
import json
import re
from collections.abc import Iterator
from dataclasses import dataclass
from functools import cache
from pathlib import Path

from conftest import ACTIVE_BUILD
from pipeline.addonrotation import _one_line
from pipeline.talent_levels import FIRST_TALENT_LEVEL

DATA = Path(__file__).resolve().parents[1]
BUILD_DIR = DATA / "builds" / ACTIVE_BUILD
APL_DIR = DATA / "curated" / "apl"
ALLOWLIST = Path(__file__).parent / "fixtures" / "card-note-levels-allowlist.txt"
#: Trainable sources a class trainer teaches; a "talent" source is read from the tree.
TRAINER_SOURCES = frozenset({"skill_line_ability", "class_spell"})
NEVER = 10_000
HASH_CHARS = 12
#: After the spec: "<name> <hash>  <reason>".
ALLOWLIST_LINE = re.compile(rf"(.+) ([0-9a-f]{{{HASH_CHARS}}})  \S")

SENTENCE_SPLIT = re.compile(r"(?<=[.!?])\s+")
CLAUSE_SPLIT = re.compile(r"[,;:]|\s[-–—]\s")
#: A level qualifier: a two-digit number that is not a duration, stack count or amount.
LEVEL_QUALIFIER = re.compile(
    r"\b(?:from|at|by|until|before)(?: level)? [1-6]\d\b"
    r"(?!\s*(?:%|seconds?|stacks?|yards?|rage|energy|mana|combo|points?|percent|health))",
    re.IGNORECASE,
)
PHRASE_QUALIFIER = re.compile(r"\bonce you have\b|\b(?:with|in) the raid build\b", re.IGNORECASE)
#: A clause that says nothing but a level: "From 25", "from 35".
LEVEL_ONLY = re.compile(r"(?:from|at|by|until|before)(?: level)? [1-6]\d", re.IGNORECASE)
#: A conditional the name follows in its clause: "once you have X", "with X", "until X".
CONDITION_BEFORE = re.compile(r"\b(?:once|with|without|until|before|after)\b", re.IGNORECASE)


@dataclass(frozen=True)
class Hit:
    spec: str
    band: int
    name: str
    has_at: int
    sentence: str

    @property
    def key(self) -> tuple[str, str, str]:
        digest = hashlib.sha256(self.sentence.encode("utf-8")).hexdigest()[:HASH_CHARS]
        return (self.spec, self.name, digest)

    def describe(self) -> str:
        have = "never" if self.has_at >= NEVER else f"at {self.has_at}"
        return (
            f"{' '.join(self.key)}  -- band {self.band}, the player has {self.name} {have}: "
            f"{self.sentence}"
        )


def _read(path: Path) -> dict:
    return json.loads(path.read_text(encoding="utf-8"))


@cache
def _addon_data() -> dict:
    return _read(BUILD_DIR / "addon-data.json")


@cache
def _builds() -> dict:
    return _read(BUILD_DIR / "addon-builds.json")


def talent_levels(spec: str, class_slug: str) -> dict[str, int]:
    """Talent name -> the first level the spec's leveling or raid order spends a point in it."""
    trees = _read(BUILD_DIR / "talents" / f"{class_slug}.json")["trees"]
    by_position = {
        (tree_index + 1, talent["tier"] + 1, talent["column"] + 1): talent["name"]
        for tree_index, tree in enumerate(trees)
        for talent in tree["talents"]
    }
    levels = {talent["name"]: NEVER for tree in trees for talent in tree["talents"]}
    for kind in ("leveling", "raid"):
        for index, point in enumerate(_builds().get(spec, {}).get(kind, {}).get("order", [])):
            name = by_position.get(tuple(point))
            if name is not None:
                levels[name] = min(levels[name], FIRST_TALENT_LEVEL + index)
    return levels


def trainer_levels(class_slug: str) -> dict[str, int]:
    """Trained spell name -> its lowest trainer rank's level."""
    levels: dict[str, int] = {}
    for spell in _read(BUILD_DIR / "trainables" / f"{class_slug}.json")["trainables"]:
        ranks = [rank["level"] for rank in spell["ranks"]]
        if spell["source"] in TRAINER_SOURCES and ranks:
            levels[spell["name"]] = min(levels.get(spell["name"], NEVER), *ranks)
    return levels


def learn_levels(spec: str) -> dict[str, int]:
    class_slug = spec.split("-", 1)[0]
    return {**trainer_levels(class_slug), **talent_levels(spec, class_slug)}


def note_bands(spec: str) -> Iterator[tuple[str, list[int]]]:
    """Each card note of the spec and the bands it renders on."""
    rotation = _addon_data()["rotations"][spec]
    curated = _read(APL_DIR / f"{spec}.json")
    steps = curated.get("rotation", {}).get("priorityList", [])

    def bands_of(note: str) -> list[int]:
        line = _one_line(note)
        return [
            band["level"]
            for band in rotation
            if any(row["condition"] == line for row in band["lines"])
        ]

    for step in steps:
        if step.get("notes"):
            yield step["notes"], bands_of(step["notes"])
    for level_note in curated.get("level_notes", []):
        base = steps[level_note["step"]].get("notes", "")
        from_level = level_note["from_level"]
        bands = [band for band in bands_of(base) if band >= from_level]
        yield level_note["notes"], bands or [from_level]


def _clauses(sentence: str) -> list[tuple[int, int]]:
    edges = [0, *(m.end() for m in CLAUSE_SPLIT.finditer(sentence)), len(sentence)]
    return list(zip(edges, edges[1:], strict=False))


def is_scoped(sentence: str, start: int) -> bool:
    """A level, "once you have" or raid-build qualifier in the name's clause, a conditional
    ("with", "until", ...) before the name in it, or a clause that is only a level qualifier
    ("From 25, ...", "Boot Knife, from 35, gives ...") anywhere in the sentence."""
    clauses = _clauses(sentence)
    if any(LEVEL_ONLY.fullmatch(sentence[a:b].strip(" ,;:")) for a, b in clauses):
        return True
    begin, end = next((a, b) for a, b in clauses if a <= start < b)
    clause = sentence[begin:end]
    return bool(
        LEVEL_QUALIFIER.search(clause)
        or PHRASE_QUALIFIER.search(clause)
        or CONDITION_BEFORE.search(sentence[begin:start])
    )


def spec_hits(spec: str) -> Iterator[Hit]:
    levels = learn_levels(spec)
    names = sorted(levels, key=len, reverse=True)
    pattern = re.compile(r"(?<![\w-])(?:" + "|".join(map(re.escape, names)) + r")(?![\w-])")
    for note, bands in note_bands(spec):
        for sentence in SENTENCE_SPLIT.split(note.strip()):
            for match in pattern.finditer(sentence):
                has_at = levels[match.group(0)]
                early = [band for band in bands if band < has_at]
                if early and not is_scoped(sentence, match.start()):
                    yield Hit(spec, min(early), match.group(0), has_at, sentence)


def allowlisted() -> set[tuple[str, str, str]]:
    keys = set()
    for line in ALLOWLIST.read_text(encoding="utf-8").splitlines():
        if line.strip() and not line.startswith("#"):
            spec, rest = line.split(" ", 1)
            match = ALLOWLIST_LINE.match(rest)
            assert match, f"allowlist line needs '<spec> <name> <hash>  <reason>': {line}"
            keys.add((spec, match.group(1), match.group(2)))
    return keys


def all_hits() -> list[Hit]:
    return [hit for spec in sorted(_addon_data()["rotations"]) for hit in spec_hits(spec)]


def test_no_card_note_names_a_talent_or_spell_before_the_player_has_it():
    accepted = allowlisted()
    unscoped = sorted(
        {hit.key: hit for hit in all_hits() if hit.key not in accepted}.values(),
        key=lambda h: h.key,
    )
    assert not unscoped, (
        "card notes name a talent or spell at a band before the player has it. Scope the name "
        '("from 30", "once you have X", "with the raid build"), fix the note in both apl '
        f"copies, or add the line to {ALLOWLIST.name} with a reason:\n"
        + "\n".join(hit.describe() for hit in unscoped)
    )


def test_every_allowlist_line_still_matches_a_note():
    stale = allowlisted() - {hit.key for hit in all_hits()}
    assert not stale, f"{ALLOWLIST.name} lines that match no card note any more: {sorted(stale)}"


def test_scope_reads_the_clause_around_the_name():
    sentence = "Fireball is the filler at 20, and Arcane Blast hits harder."
    assert is_scoped(sentence, sentence.index("Fireball"))
    assert not is_scoped(sentence, sentence.index("Arcane Blast"))
    assert is_scoped("Once you have Heating Up, Fire Blast is a stack.", 14)
    assert not is_scoped("At 6 seconds Pyroblast is free.", 13)
