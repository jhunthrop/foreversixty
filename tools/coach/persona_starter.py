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
from starter_overrides import GUIDE_LINES, GUIDE_ORDER, corrected
from starter_priority import corroborated_note, is_bare, priority_entries
from starter_text import DIGIT, number_words, spell_out_numbers, usable_sentence

__all__ = [
    "ALL_CODES",
    "GUIDE_LINES",
    "GUIDE_ORDER",
    "ROLE_CODES",
    "corroborated_note",
    "number_words",
    "usable_sentence",
]

TOOLS_DIR = Path(__file__).resolve().parent
REPO = TOOLS_DIR.parent.parent
COACH_DIR = REPO / "data" / "curated" / "coach"
SPECS_PATH = "data/curated/specs.json"
CATALOGUE_PATH = "logs/engine/mechanics/consumables/catalogue.json"
ACTIVE_BUILD_PATH = "web/src/data/active-build.json"
FULL_SPECS = frozenset({"warrior-fury", "priest-holy", "warrior-protection"})
STARTER_VERSION = "starter edition"
# The recorded wow-player read of every starter (design/reviews/2026-10-27-log-coach-personas), pinned
# so regeneration is deterministic. Change it only when a new read is recorded.
REVIEWED = "wow-player, 2026-10-28"

ROTATION_HEADING = "## Rotation and priority"
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


def castable_names(row: dict[str, Any]) -> set[str]:
    """Names of spells a player presses: those with a cost, a global cooldown, a cooldown or a cast time."""
    spells = load_json(spellconst_path(row))["spells"].values()
    pressed = ("cost", "gcd_ms", "cooldown_ms", "cast_time_ms")
    return {spell["name"] for spell in spells if any(spell[key] for key in pressed)}


def priority_entries_of(row: dict[str, Any]) -> list[tuple[str, str]]:
    prose = section_after(read_text(guide_path(row)), ROTATION_HEADING)
    items = load_json(apl_path(row["spec"]))["rotation"]["priorityList"]
    entries = priority_entries(items, spell_names(row), castable_names(row), prose, read_text(guide_path(row)))
    return corrected(row["spec"], entries)


def priority_lines(row: dict[str, Any]) -> list[str]:
    return [line for _, line in priority_entries_of(row)]


def bare_lines(row: dict[str, Any]) -> list[str]:
    """Abilities that carry no reason in their line: nothing in the notes or the guide led with them."""
    return [name for name, line in priority_entries_of(row) if is_bare((name, line))]


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
        "**The priority, in words.** These are the buttons the guide's rotation names, in the order it gives them. "
        "Where the guide or the rotation data gives a reason, it follows the name.",
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
        f"version: {STARTER_VERSION}\nsources:\n{listed}\nreviewed: {REVIEWED}\n---\n"
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
