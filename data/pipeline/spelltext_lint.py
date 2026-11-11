"""Find the rendered spell text that does not read like the game client's.

Every description the pipeline ships (talent ranks, spell and consumable
tooltips, item use/equip/proc text, set bonuses) goes through
`pipeline.spelltext`. This module reads those shipped strings back out of a
build directory and sorts what is wrong with them into defect classes, so a
renderer gap shows up as a count instead of as one tooltip somebody happened
to read (2026-11-10: Booming Voice's "Rage cost by of your Shouts").

`scan_build` is the one reader; `pipeline spelltext-check` and
`tests/test_spelltext_contract.py` both call it, so the gate and the report
agree on what a defect is.

A defect the client's own template carries word for word (Booming Voice's
"by of" is in the 1.60 client's Spell.Description_lang and on Wowhead's
Forever tooltip, both 2026-11-10) is a client typo, not a renderer gap: the
site prints what the game prints. Those strings are listed, each with its
source, in `CLIENT_VERBATIM` and never counted.
"""

from __future__ import annotations

import json
import re
from collections.abc import Iterator
from dataclasses import dataclass
from pathlib import Path

from pipeline.spelltext_grammar import MID_SENTENCE_SPACES

#: A `$` the renderer left: `$s1`, `${..}`, `$<mult>`, `$?s123[..]`, `$@spelldesc123`.
_DOLLAR = re.compile(r"\$")
#: A brace with no `$` before it: the remains of a half-worked `${..}`.
_BRACE = re.compile(r"[{}]")
#: An `@` token whose `$` is gone, or still there.
_AT = re.compile(r"@(?:spell|null)")
#: The words a value sits between, with the value gone: "cost by of your", "by %",
#: "for  sec", "of .", "( )".
_EMPTY_SUBSTITUTION = re.compile(
    r"\b(?:by|for|of|to)"
    r" +(?:of|by|for|to|and|sec|min|yd|%|\.|,)(?=\W|$)"
    r"|\(\s*\)"
    r"|(?<=\s)[%.,](?=\s|$)"
)
#: "-0", "-0.0".
_NEGATIVE_ZERO = re.compile(r"(?<![\w.])-0(?:\.0+)?(?!\.?\d)")
#: A duration or cooldown of nothing: "for 0 sec", "every 0 sec".
_ZERO_DURATION = re.compile(r"(?<![\w.])0 (?:sec|min|yd|yards)\b")
#: A chance or amount of nothing: "a 0% chance", "by 0%", "restore 0.0%". A talent whose
#: per-rank value lives on a curve renders 0 when it is read without that curve.
_ZERO_AMOUNT = re.compile(r"(?<![\w.])0(?:\.0+)?%")

DEFECT_PATTERNS: dict[str, re.Pattern[str]] = {
    "unrendered $ token": _DOLLAR,
    "stray brace": _BRACE,
    "unrendered @ token": _AT,
    "client colour code": re.compile(r"\|[cC][0-9A-Fa-f]{8}|\|[rR](?![a-z])"),
    "double space mid-sentence": MID_SENTENCE_SPACES,
    "empty substitution": _EMPTY_SUBSTITUTION,
    "negative zero": _NEGATIVE_ZERO,
    "zero duration": _ZERO_DURATION,
    "zero amount": _ZERO_AMOUNT,
}

#: Strings the client itself ships with the defect. Each entry names where the
#: verbatim text was confirmed; the test suite keeps the list honest.
CLIENT_VERBATIM: dict[str, str] = {
    # Booming Voice (12321): Spell.Description_lang in 1.60.1.70790's hotfix push reads
    # "...reduces the Rage cost by of your Shouts and Howls by $s2%."; Wowhead's Forever
    # tooltip (nether.wowhead.com/tooltip/spell/12321?dataEnv=16, 2026-11-10) prints the
    # same "by of".
    "reduces the Rage cost by of your Shouts and Howls": "client typo, Booming Voice 12321",
}


@dataclass(frozen=True)
class Finding:
    defect: str
    file: str
    where: str
    text: str


@dataclass(frozen=True)
class ShippedText:
    file: str
    where: str
    text: str


def _talent_texts(path: Path, rel: str) -> Iterator[ShippedText]:
    data = json.loads(path.read_text(encoding="utf-8"))
    for tree in data.get("trees", []):
        for talent in tree.get("talents", []):
            for number, rank in enumerate(talent.get("ranks", []), start=1):
                yield ShippedText(rel, f"{talent['name']} rank {number}", rank["description"])


def _spelltip_texts(path: Path, rel: str) -> Iterator[ShippedText]:
    data = json.loads(path.read_text(encoding="utf-8"))
    for group, entries in data.items():
        if not isinstance(entries, dict):
            continue
        for key, entry in entries.items():
            if not isinstance(entry, dict):
                continue
            for field in ("description", "use"):
                if isinstance(entry.get(field), str):
                    yield ShippedText(rel, f"{group}/{key}.{field}", entry[field])


def _item_texts(path: Path, rel: str) -> Iterator[ShippedText]:
    data = json.loads(path.read_text(encoding="utf-8"))
    for item in data.get("items", []):
        if item.get("effect_text"):
            yield ShippedText(rel, f"item {item['id']} {item['name']}", item["effect_text"])


def _set_texts(path: Path, rel: str) -> Iterator[ShippedText]:
    for item_set in json.loads(path.read_text(encoding="utf-8")):
        for bonus in item_set.get("bonuses", []):
            where = f"set {item_set.get('id')} {item_set.get('name')} ({bonus.get('pieces')})"
            yield ShippedText(rel, where, bonus["description"])


def shipped_texts(build_dir: Path) -> Iterator[ShippedText]:
    """Every rendered description a build directory ships, one per occurrence."""
    readers = (
        ("talents/*.json", _talent_texts),
        ("spelltips/*.json", _spelltip_texts),
        ("items/*.json", _item_texts),
        ("sets.json", _set_texts),
    )
    for pattern, reader in readers:
        for path in sorted(build_dir.glob(pattern)):
            yield from reader(path, path.relative_to(build_dir).as_posix())


def defects_in(text: str) -> list[str]:
    """The defect classes one rendered string falls into, client typos excepted."""
    checked = text
    for verbatim in CLIENT_VERBATIM:
        checked = checked.replace(verbatim, "verbatim")
    return [name for name, pattern in DEFECT_PATTERNS.items() if pattern.search(checked)]


def scan_build(build_dir: Path) -> tuple[int, list[Finding]]:
    """How many shipped descriptions the build has, and every defect in them."""
    total = 0
    findings: list[Finding] = []
    for shipped in shipped_texts(build_dir):
        total += 1
        findings.extend(
            Finding(defect, shipped.file, shipped.where, shipped.text)
            for defect in defects_in(shipped.text)
        )
    return total, findings


def summarize(findings: list[Finding]) -> dict[str, dict[str, int]]:
    """Defect class -> occurrences and unique strings, in `DEFECT_PATTERNS` order."""
    out: dict[str, dict[str, int]] = {}
    for defect in DEFECT_PATTERNS:
        rows = [finding for finding in findings if finding.defect == defect]
        if rows:
            out[defect] = {"occurrences": len(rows), "unique": len({r.text for r in rows})}
    return out


#: How many findings `check_build` prints before it stops listing them.
_SHOWN = 40


def check_build(build: str, root: Path = Path("builds")) -> int:
    """`pipeline spelltext-check`: 0 when every shipped description renders clean, else 1."""
    total, findings = scan_build(root / build)
    print(f"{build}: {total} shipped descriptions, {len(findings)} defects")
    for defect, counts in summarize(findings).items():
        print(f"  {defect}: {counts['occurrences']} ({counts['unique']} unique)")
    for finding in findings[:_SHOWN]:
        print(f"  [{finding.defect}] {finding.file} {finding.where}: {finding.text[:160]!r}")
    return 1 if findings else 0
