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

TOOLS_DIR = Path(__file__).resolve().parent
REPO = TOOLS_DIR.parent.parent
COACH_DIR = REPO / "data" / "curated" / "coach"
SPECS_PATH = "data/curated/specs.json"
CATALOGUE_PATH = "logs/engine/mechanics/consumables/catalogue.json"
FULL_SPECS = frozenset({"warrior-fury", "priest-holy", "warrior-protection"})
STARTER_VERSION = "starter edition"
MIN_SENTENCE_CHARS = 20

ROLE_LABEL = {"dps": "damage dealer", "healer": "healer", "tank": "tank"}
HEALER_CODES = ("core_casts", "uptime", "resource", "ranks", "preparation", "death")
ALL_CODES = ("core_casts", "uptime", "idle", "resource", "ranks", "preparation", "death", "adherence")
ROLE_CODES = {"dps": ALL_CODES, "tank": ALL_CODES, "healer": HEALER_CODES}

# What each finding code means, in the words of the rule.
MISTAKES = {
    "core_casts": "a button the priority list wants sitting ready while something lower went first.",
    "uptime": "a buff, debuff or damage over time left to fall off the target.",
    "idle": "time with no button pressed while the boss was alive and in reach.",
    "resource": "the main resource bar left to cap or to run dry.",
    "ranks": "a spell left at an old rank because the trainer visit was skipped.",
    "preparation": "a missing flask, food, weapon enchant or combat potion.",
    "death": "dying, which ends your contribution for the rest of the pull.",
    "adherence": "buttons outside the priority list taking globals the list would have used.",
}
HEALER_MISTAKES = {
    "core_casts": "a mana or efficiency cooldown sitting ready for most of the fight.",
    "resource": "the mana bar empty well before the end of the fight.",
    "ranks": "a big rank landing mostly as overheal when a lower rank of the same heal would have covered the hole.",
    "death": "dying, because a dead healer heals nothing for the rest of the pull.",
}

# The default drill per code, worded by role. Healer drills never say "heal more".
DRILLS = {
    "dps": {
        "core_casts": "On a dummy, press the first button of the priority list the moment it is ready, before any filler. Read the list top to bottom and check it again after every cooldown.",
        "uptime": "Put the timers of your buffs and debuffs where you can see them. Recast each with 3 seconds left, and again the moment you are back up after a death or a phase change.",
        "idle": "On every move, keep the boss in range and press a button the moment you arrive. Practise on trash: the first press after a reposition comes before any walk to the side.",
        "resource": "Watch the main resource bar on a dummy. Spend it on the list's buttons before it can cap, and never let a button wait for want of it.",
        "ranks": "Visit your trainer every time you level and read your spellbook after: every button on the priority list should show its top rank.",
        "preparation": "Before every boss, check flask, food buff, weapon enchant and a combat potion on your bars. Do it on the walk to the pull, not after it.",
        "death": "Before the pull, find the spot where the boss does not cleave or breathe. Keep a healing potion on a key you can reach without looking.",
        "adherence": "On a dummy, play only the buttons on the priority list for 3 minutes. Anything else on your bars waits.",
    },
    "tank": {
        "core_casts": "On a dummy with a friend hitting you, press the first button of the priority list the moment it is ready. Fillers wait their turn.",
        "uptime": "Watch the buff on you and the debuffs on the boss, not the cooldowns. The moment one drops, press it again.",
        "idle": "When you move the boss, keep it in front of you and keep pressing buttons. Your first press after the move is a priority button, not a walk to the next spot.",
        "resource": "Watch the main resource bar. Spend it on the list's buttons before it can cap, and never let a button wait for want of it.",
        "ranks": "Visit your trainer every time you level and read your spellbook after: every button on the priority list should show its top rank.",
        "preparation": "Before every boss, check flask, food buff, weapon enchant, a combat potion and a healing potion on your bars. Do it on the walk to the pull, not after it.",
        "death": "Put your emergency cooldown on a key you can hit without looking, and press it on a dummy pull the moment your health falls past half.",
        "adherence": "On a dummy, play only the buttons on the priority list for 3 minutes. Anything else on your bars waits.",
    },
    "healer": {
        "core_casts": "Put your efficiency cooldown next to your most expensive heal and press them as a pair. When the mana bar drops below half, check your mana cooldowns before the next cast.",
        "uptime": "Recast your self buff as part of the ready check, and again the moment you are back up after a death.",
        "resource": "Before the pull, note the mana you should have at the halfway mark. If you are behind it, let the cheap heals over time carry the tank and save the fast, dear heal for the real emergency.",
        "ranks": "Bind a mid rank of your filler heal and use it for every small hole on the tank. Save the top rank for a hole as deep as it heals.",
        "preparation": "Before every boss, check flask, mana food and a mana potion on your bars. Do it on the walk to the pull, not after it.",
        "death": "Before the pull, pick a spot in range of the tank and out of anything the boss breathes or cleaves. When damage lands on you, move first and finish the cast after.",
    },
}

SMALL = (
    "zero one two three four five six seven eight nine ten eleven twelve thirteen fourteen "
    "fifteen sixteen seventeen eighteen nineteen"
).split()
TENS = "_ _ twenty thirty forty fifty sixty seventy eighty ninety".split()
# Words that mark a sentence as about the simulator rather than about playing the spec.
SIMULATOR_JARGON = re.compile(
    r"\b(engine|ladder|rung|golden|harness|inert|unresolved|gated|sim|simulator|simulated|"
    r"simulation|the search|conditions?)\b",
    re.IGNORECASE,
)
# References to other lines of the rotation data, which a lone sentence cannot carry.
DANGLING_REFERENCE = re.compile(r"\b(above|below it|same pace|paced like|the pass)\b", re.IGNORECASE)
PARENTHETICAL = re.compile(r"\s*\([^)]*\)")
NUMBER = re.compile(r"(?<![\w.])(\d{1,3})(%?)(?![\w.]|\d)")
PERCENT_WORD = " percent"
DIGIT = re.compile(r"\d")


def number_words(n: int) -> str:
    if n < 20:
        return SMALL[n]
    if n < 100:
        return TENS[n // 10] + (f"-{SMALL[n % 10]}" if n % 10 else "")
    rest = f" {number_words(n % 100)}" if n % 100 else ""
    return f"{SMALL[n // 100]} hundred{rest}"


def spell_out_numbers(text: str) -> str:
    """Replace whole integers up to 999 (and a trailing percent sign) with words."""
    return NUMBER.sub(lambda m: number_words(int(m[1])) + (PERCENT_WORD if m[2] else ""), text)


def usable_sentence(note: str) -> str | None:
    """A note's first sentence in words, or None. Later sentences depend on it, so are never used."""
    lead = re.split(r"(?<=[.!?])\s+", note.strip())[0]
    words = spell_out_numbers(PARENTHETICAL.sub("", lead).replace("!", "."))
    if DIGIT.search(words) or SIMULATOR_JARGON.search(words) or DANGLING_REFERENCE.search(words):
        return None
    if len(words) < MIN_SENTENCE_CHARS:
        return None
    return words if words.endswith((".", ":")) else words + "."


def load_json(relative: str) -> Any:
    return json.loads((REPO / relative).read_text(encoding="utf-8"))


def spec_row(slug: str) -> dict:
    for row in load_json(SPECS_PATH):
        if row["spec"] == slug:
            return row
    raise SystemExit(f"unknown spec {slug!r}: not in {SPECS_PATH}")


def spec_slugs() -> list[str]:
    return [row["spec"] for row in load_json(SPECS_PATH)]


def guide_path(row: dict) -> str:
    return f"web/src/content/guides/{row['class_slug']}/{row['spec_slug']}.md"


def apl_path(slug: str) -> str:
    return f"data/curated/apl/{slug}.json"


def overview(row: dict) -> str:
    text = (REPO / guide_path(row)).read_text(encoding="utf-8")
    match = re.search(r"^## Overview\n\n(.+?)\n\n", text, re.MULTILINE | re.DOTALL)
    if not match:
        raise SystemExit(f"{guide_path(row)}: no Overview paragraph")
    paragraph = spell_out_numbers(" ".join(match[1].split()))
    if DIGIT.search(paragraph):
        raise SystemExit(f"{guide_path(row)}: Overview paragraph carries a number")
    return paragraph


def priority_lines(slug: str) -> tuple[int, list[str]]:
    """The count of lines in the rotation data and the reasons that read on their own, in order."""
    items = load_json(apl_path(slug))["rotation"]["priorityList"]
    lines = (usable_sentence(item.get("notes", "")) for item in items)
    return len(items), [line for line in lines if line]


def join_names(names: list[str]) -> str:
    if len(names) <= 1:
        return "".join(names)
    return ", ".join(names[:-1]) + " or " + names[-1]


def names_of(entries: list[dict]) -> list[str]:
    return [entry["name"] for entry in entries]


def consumable_paragraph(role: str) -> str:
    entry = load_json(CATALOGUE_PATH)["roles"][role]
    wanted = [f"a flask ({join_names(names_of(entry['flask']))})", f"food ({join_names(names_of(entry['food']))})"]
    if entry["weapon_enchant"]:
        wanted.append(f"a weapon enchant ({join_names(names_of(entry['weapon_enchant']))})")
    potion = join_names(names_of(entry["combat_potion"]["entries"]))
    wanted.append(f"a combat potion ({potion}), up to its use limit in a fight")
    return (
        f"For a {ROLE_LABEL[role]} the catalogue looks for {'; '.join(wanted)}. "
        f"The world buffs it lists are {join_names(names_of(entry['world_buffs']))}. "
        "They are listed as facts, never held against anyone."
    )


def mistakes(role: str) -> list[str]:
    table = {**MISTAKES, **(HEALER_MISTAKES if role == "healer" else {})}
    return [f"- {code}: {table[code]}" for code in ROLE_CODES[role]]


def who_section(row: dict) -> str:
    return (
        f"This is the {row['name']} {row['class_slug'].capitalize()} as the site's guide describes it. "
        f"{overview(row)} "
        "The best players' habits for this spec are not written down here yet. "
        "This file says what the guide says and leaves the rest for a later full write."
    )


def know_section(row: dict) -> str:
    total, lines = priority_lines(row["spec"])
    priority = "\n".join(f"- {line}" for line in lines)
    return "\n\n".join([
        f"**The priority, in words.** The site's rotation data for this spec holds {number_words(total)} lines. "
        "These are the reasons written on them that read on their own, in the order the lines run. "
        "The guide and the rotation data hold the rest.",
        priority,
        "**The mistakes the coach looks for, by finding.**\n" + "\n".join(mistakes(row["role"])),
        "**Consumables and world buffs.** " + consumable_paragraph(row["role"]),
    ])


def drills_section(role: str) -> str:
    return "\n".join(f"- drill.{code}: {DRILLS[role][code]}" for code in ROLE_CODES[role])


def front_matter(row: dict) -> str:
    sources = [guide_path(row), apl_path(row["spec"]), CATALOGUE_PATH, SPECS_PATH]
    listed = "\n".join(f"  - {s}" for s in sources)
    return (
        f"---\nspec: {row['spec']}\nrole: {row['role']}\nstatus: starter\n"
        f"version: {STARTER_VERSION}\nsources:\n{listed}\nreviewed: pending\n---\n"
    )


def render(slug: str) -> str:
    row = spec_row(slug)
    voice = (TOOLS_DIR / "voice.md").read_text(encoding="utf-8").strip()
    return (
        f"{front_matter(row)}\n## Who they are\n\n{who_section(row)}\n\n"
        f"## What they know\n\n{know_section(row)}\n\n"
        f"## Drills\n\n{drills_section(row['role'])}\n\n"
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
