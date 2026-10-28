"""Generate the starter coach personas (design spec section 8.2).

A starter says what the spec's guide and rotation data say and nothing more:
the best-player habits are left for a later full write. Output is deterministic.

    python tools/coach/persona_starter.py --spec mage-frost --out data/curated/coach/mage-frost.md
    python tools/coach/persona_starter.py --all
"""

from __future__ import annotations

import argparse
import json
import re
from pathlib import Path
from typing import Any

from starter_kit import ALL_CODES, ROLE_CODES, consumable_paragraph, drills, mistakes
from starter_text import (
    DIGIT,
    digit_tokens,
    number_words,
    plain_sentence,
    sentences,
    spell_out_numbers,
    usable_sentence,
)

__all__ = ["ALL_CODES", "ROLE_CODES", "number_words", "usable_sentence"]

TOOLS_DIR = Path(__file__).resolve().parent
REPO = TOOLS_DIR.parent.parent
COACH_DIR = REPO / "data" / "curated" / "coach"
SPECS_PATH = "data/curated/specs.json"
CATALOGUE_PATH = "logs/engine/mechanics/consumables/catalogue.json"
ACTIVE_BUILD_PATH = "web/src/data/active-build.json"
FULL_SPECS = frozenset({"warrior-fury", "priest-holy", "warrior-protection"})
STARTER_VERSION = "starter edition"

# Claims the rotation data makes that the guides do not, so a starter does not repeat them.
CLAIMS_THE_GUIDE_DOES_NOT_MAKE = re.compile(
    r"most healing per mana|lifts Rupture damage|more damage from your Rupture|\bMaul\b.*\brage\b",
    re.IGNORECASE,
)
ROTATION_HEADING = "## Rotation and priority"
NAME_NEAR_FRONT_CHARS = 40
# Spells the class spellconst files do not carry.
EXTRA_SPELL_NAMES = {5019: "Shoot"}


def load_json(relative: str) -> Any:
    return json.loads((REPO / relative).read_text(encoding="utf-8"))


def read_text(relative: str) -> str:
    return (REPO / relative).read_text(encoding="utf-8")


def spec_row(slug: str) -> dict[str, Any]:
    for row in load_json(SPECS_PATH):
        if row["spec"] == slug:
            return row
    raise SystemExit(f"unknown spec {slug!r}: not in {SPECS_PATH}")


def spec_slugs() -> list[str]:
    return [row["spec"] for row in load_json(SPECS_PATH)]


def guide_path(row: dict[str, Any]) -> str:
    return f"web/src/content/guides/{row['class_slug']}/{row['spec_slug']}.md"


def apl_path(slug: str) -> str:
    return f"data/curated/apl/{slug}.json"


def spellconst_path(row: dict[str, Any]) -> str:
    build = load_json(ACTIVE_BUILD_PATH)["build"]
    return f"data/builds/{build}/spellconst/{row['class_slug']}.json"


def section_after(text: str, heading: str) -> str:
    start = text.index(heading) + len(heading)
    rest = text[start:]
    following = re.search(r"^## ", rest, re.MULTILINE)
    return rest[: following.start()] if following else rest


def overview(row: dict[str, Any]) -> str:
    match = re.search(r"^## Overview\n\n(.+?)\n\n", read_text(guide_path(row)), re.MULTILINE | re.DOTALL)
    if not match:
        raise SystemExit(f"{guide_path(row)}: no Overview paragraph")
    paragraph = spell_out_numbers(" ".join(match[1].split()))
    if DIGIT.search(paragraph):
        raise SystemExit(f"{guide_path(row)}: Overview paragraph carries a number")
    return paragraph


def spell_names(row: dict[str, Any]) -> dict[int, str]:
    spells = load_json(spellconst_path(row))["spells"]
    return {**{int(spell_id): spell["name"] for spell_id, spell in spells.items()}, **EXTRA_SPELL_NAMES}


def guide_rotation_sentences(row: dict[str, Any]) -> list[str]:
    prose = section_after(read_text(guide_path(row)), ROTATION_HEADING)
    return [s for s in (plain_sentence(raw) for raw in sentences(prose)) if s]


def leads_with(name: str, sentence: str) -> bool:
    """The sentence is about the ability: it starts as a sentence and names it near the front."""
    return sentence[:1].isupper() and 0 <= sentence.lower().find(name.lower()) <= NAME_NEAR_FRONT_CHARS


def sentence_naming(name: str, candidates: list[str]) -> str | None:
    return next((s for s in candidates if leads_with(name, s)), None)


def corroborated_note(note: str, guide_digits: set[str]) -> str | None:
    """The note's lead sentence in words, if it reads alone and the guide backs every number and claim."""
    lead = sentences(note)[0] if note.strip() else ""
    if not lead or CLAIMS_THE_GUIDE_DOES_NOT_MAKE.search(lead) or not digit_tokens(lead) <= guide_digits:
        return None
    return usable_sentence(note)


def cast_spell_id(item: dict[str, Any]) -> int | None:
    cast = item["action"].get("castSpell")
    return cast.get("spellId", {}).get("spellId") if cast else None


def priority_line(item: dict[str, Any], names: dict[int, str], guide: list[str], guide_digits: set[str]) -> tuple[str, str] | None:
    """(ability name, line): the note's sentence if it names the ability, else the guide's, else the name alone."""
    sentence = corroborated_note(item.get("notes", ""), guide_digits)
    spell_id = cast_spell_id(item)
    name = names.get(spell_id) if spell_id is not None else None
    if name is None:
        return (sentence, sentence) if sentence else None
    if sentence and leads_with(name, sentence):
        return name, sentence
    return name, sentence_naming(name, guide) or f"{name}."


def priority_lines(row: dict[str, Any]) -> list[str]:
    names, guide = spell_names(row), guide_rotation_sentences(row)
    guide_digits = digit_tokens(read_text(guide_path(row)))
    seen: set[str] = set()
    lines: list[str] = []
    for item in load_json(apl_path(row["spec"]))["rotation"]["priorityList"]:
        built = priority_line(item, names, guide, guide_digits)
        if built and built[0].lower() not in seen and built[1] not in lines:
            seen.add(built[0].lower())
            lines.append(built[1])
    return lines


def who_section(row: dict[str, Any]) -> str:
    return (
        f"This is the {row['name']} {row['class_slug'].capitalize()} as the site's guide describes it. "
        f"{overview(row)} "
        "The best players' habits for this spec are not written down here yet. "
        "This file says what the guide says and leaves the rest for a later full write."
    )


def know_section(row: dict[str, Any], catalogue: dict[str, Any]) -> str:
    slug, role = row["spec"], row["role"]
    priority = "\n".join(f"- {line}" for line in priority_lines(row))
    return "\n\n".join([
        "**The priority, in words.** The site's rotation data for this spec runs these buttons in this order, "
        "from the top of the list down. Where the guide or the rotation data gives a reason, it follows the name.",
        priority,
        "**The mistakes the coach looks for, by finding.**\n" + "\n".join(mistakes(slug, role, catalogue)),
        "**Consumables and world buffs.** " + consumable_paragraph(slug, role, catalogue),
    ])


def drills_section(row: dict[str, Any], catalogue: dict[str, Any]) -> str:
    return "\n".join(f"- drill.{code}: {text}" for code, text in drills(row["spec"], row["role"], catalogue))


def front_matter(row: dict[str, Any]) -> str:
    sources = [guide_path(row), apl_path(row["spec"]), spellconst_path(row), CATALOGUE_PATH, SPECS_PATH]
    listed = "\n".join(f"  - {s}" for s in sources)
    return (
        f"---\nspec: {row['spec']}\nrole: {row['role']}\nstatus: starter\n"
        f"version: {STARTER_VERSION}\nsources:\n{listed}\nreviewed: pending\n---\n"
    )


def render(slug: str) -> str:
    row = spec_row(slug)
    catalogue = load_json(CATALOGUE_PATH)
    voice = (TOOLS_DIR / "voice.md").read_text(encoding="utf-8").strip()
    return (
        f"{front_matter(row)}\n## Who they are\n\n{who_section(row)}\n\n"
        f"## What they know\n\n{know_section(row, catalogue)}\n\n"
        f"## Drills\n\n{drills_section(row, catalogue)}\n\n"
        f"## Voice\n\n{voice}\n"
    )


def write_starter(slug: str, out: Path) -> None:
    if slug in FULL_SPECS:
        raise SystemExit(f"{slug} is a full persona; not generated")
    out.parent.mkdir(parents=True, exist_ok=True)
    out.write_text(render(slug), encoding="utf-8")


def write_all(out_dir: Path) -> list[str]:
    slugs = [s for s in spec_slugs() if s not in FULL_SPECS]
    for slug in slugs:
        write_starter(slug, out_dir / f"{slug}.md")
    return slugs


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument("--spec", help="spec slug, e.g. mage-frost")
    group.add_argument("--all", action="store_true", help="every spec that is not a full persona")
    parser.add_argument("--out", type=Path, help="output file (--spec) or directory (--all)")
    args = parser.parse_args()
    if args.all:
        write_all(args.out or COACH_DIR)
    else:
        write_starter(args.spec, args.out or COACH_DIR / f"{args.spec}.md")


if __name__ == "__main__":
    main()
