"""Per-spec corrections to a generated priority list, each a short restatement of the spec's guide.

The generator reads the guide's rotation prose sentence by sentence. Where the guide states an order
in words ("goes down first", "sits ahead of", "comes last"), a condition the line must carry ("only
with a Windfury Totem in the group"), or a sentence too long to read at a glance, the reviewer's
correction is recorded here (wow-player read 4). Every line restates its guide: a test checks that
each ability is named in the guide's rotation section and that no line carries a digit.
"""

from __future__ import annotations

import re

from starter_text import short_sentences

Entry = tuple[str, str]

# ability -> line. A listed ability is added when the generator left it out.
GUIDE_LINES: dict[str, dict[str, str]] = {
    "druid-feral": {
        # feral.md: "Press it every time it comes off cooldown, but only when your energy is 60 or lower".
        "Shifting Power": (
            "Shifting Power every time it comes off cooldown, but only at sixty energy or lower. "
            "Then the whole forty energy fits under the cap."
        ),
    },
    "hunter-survival": {
        # survival.md: "it can only be performed after you dodge, and Expose Prey opens the same window".
        "Mongoose Bite": "Mongoose Bite only after you dodge, or while Expose Prey is up: it is a counterattack.",
        # survival.md: "only in a raid: with a Windfury Totem in the group ... a loss without the totem".
        "Wing Clip": (
            "Wing Clip only with a Windfury Totem in the group, on a global with nothing else ready. "
            "Without the totem it loses damage."
        ),
    },
    "warrior-arms": {
        # arms.md: "topping up Rage with Bloodrage whenever it drops below 80".
        "Bloodrage": "Bloodrage whenever Rage drops below eighty, since Rage limits everything else here.",
        # arms.md: "With a Windfury Totem in the group, spend any global that is still free on Hamstring".
        "Hamstring": (
            "Hamstring only with a Windfury Totem in the group, on a free global for ten rage. "
            "Without the totem it loses damage."
        ),
    },
    "mage-fire": {
        # fire.md: "Scorch is the filler for every global that isn't spent on a three-stack Pyroblast ...".
        "Scorch": (
            "Scorch on every global not spent on a three-stack Pyroblast, Fire Blast or Combustion. "
            "Every Scorch also renews the Improved Scorch debuff."
        ),
    },
    "paladin-holy": {
        # holy.md: "Holy Shock ... only fire while your mana is keeping up with the share of the fight".
        "Holy Shock": "Holy Shock only while your mana is keeping pace with the fight.",
    },
    "paladin-protection": {
        # protection.md list items 4 and 6.
        "Hammer of the Righteous": (
            "Hammer of the Righteous and Holy Strike on their cooldowns. "
            "Hammer of the Righteous is the largest single threat gain in this list."
        ),
        "Consecration": "Consecration only with spare mana, where it pays for itself in threat.",
    },
    "paladin-retribution": {
        # retribution.md: "reapply Seal of Righteousness just before it expires so uptime never lapses".
        "Seal of Righteousness": (
            "Seal of Righteousness about a second and a half before the pull. "
            "Reapply it just before it expires so it never lapses."
        ),
        "Hammer of Wrath": "Hammer of Wrath below twenty percent target health, at its highest rank.",
        "Holy Strike": (
            "Holy Strike on its own ten-second cooldown, whenever it is up. "
            "It spends the globals Judgement and seal upkeep leave idle."
        ),
        # retribution.md: "a loss on a bare character ... only cast while mana is above 65%".
        "Consecration": (
            "Consecration last, and only while mana is above sixty-five percent. "
            "Against one target it is a loss without the raid's mana support."
        ),
    },
    "priest-shadow": {
        "Shadow Word: Pain": "Shadow Word: Pain kept up, unless the fight is close to ending.",
        "Inner Focus": (
            "Inner Focus into Devouring Plague outside the execute window: "
            "a free, empowered cast at its full top rank."
        ),
        "Mind Flay": "Mind Flay on every free global.",
        "Mind Blast": "Mind Blast only when Mind Flay cannot be cast.",
        "Shadow Word: Death": (
            "Shadow Word: Death first below twenty percent target health. "
            "While it is on its fifteen-second cooldown, Mind Blast, then Mind Flay."
        ),
        # shadow.md: "Wand shots fill any global where nothing else is ready."
        "Shoot": "Shoot with your wand on any global where nothing else is ready.",
    },
    "shaman-elemental": {
        # elemental.md: "Mana Spring Totem goes down first ... starting it one global before the cooldowns".
        "Mana Spring Totem": (
            "Mana Spring Totem goes down first, one global before the cooldowns. "
            "The loop runs short of mana long before the fight ends."
        ),
        "Lava Burst": "Lava Burst on every cooldown when you have the talent, ahead of the totems and the cooldowns.",
        "Searing Totem": "Searing Totem against a single target, for the higher fire damage.",
        "Lightning Bolt": "Lightning Bolt on every free global.",
        "Chain Lightning": (
            "Chain Lightning only against two or more targets. "
            "It hits every target for close to single-target Lightning Bolt damage."
        ),
        "Magma Totem": "Magma Totem in place of Searing Totem against two or more targets.",
    },
    "shaman-enhancement": {
        # enhancement.md: imbues "applied before the pull and reapplied whenever it falls off".
        "Windfury Weapon": (
            "Windfury Weapon on your one weapon before the pull, and again whenever it falls off. "
            "Before level thirty, Rockbiter Weapon goes on it instead."
        ),
        "Earth Shock": "Earth Shock last, and only while mana is above half.",
    },
}

# The guide's own order statements, which win over the order of first mention. Listed abilities lead.
GUIDE_ORDER: dict[str, tuple[str, ...]] = {
    # feral.md: Shifting Power "sits ahead of Tiger's Fury and every builder in the order".
    "druid-feral": ("Shifting Power",),
    # protection.md list items 1-6.
    "paladin-protection": (
        "Righteous Fury", "Seal of Fury", "Holy Shield", "Hammer of the Righteous", "Holy Strike",
        "Judgement", "Consecration",
    ),
    # shadow.md: "After that, Mind Flay on every free global comes before Mind Blast".
    "priest-shadow": (
        "Shadow Word: Pain", "Inner Focus", "Mind Flay", "Mind Blast", "Shadow Word: Death", "Shoot",
    ),
    # elemental.md: Mana Spring Totem first, Lava Burst ahead of the totems, Lightning Bolt the rest.
    "shaman-elemental": (
        "Mana Spring Totem", "Lava Burst", "Searing Totem", "Lightning Bolt", "Chain Lightning",
        "Magma Totem",
    ),
}


def with_lines(slug: str, entries: list[Entry]) -> list[Entry]:
    lines = GUIDE_LINES.get(slug, {})
    present = {name for name, _ in entries}
    replaced = [(name, lines.get(name, line)) for name, line in entries]
    return replaced + [(name, line) for name, line in lines.items() if name not in present]


def in_guide_order(slug: str, entries: list[Entry]) -> list[Entry]:
    leading = GUIDE_ORDER.get(slug, ())
    rank = {name: index for index, name in enumerate(leading)}
    return sorted(entries, key=lambda entry: rank.get(entry[0], len(leading)))


def without_repeated_bare(entries: list[Entry]) -> list[Entry]:
    """A bare name another line already names is a repeat (read 4, pattern C)."""
    lines = [line for _, line in entries]
    return [
        (name, line)
        for name, line in entries
        if line != f"{name}." or not any(re.search(rf"\b{re.escape(name)}\b", o) for o in lines if o != line)
    ]


def corrected(slug: str, entries: list[Entry]) -> list[Entry]:
    ordered = without_repeated_bare(in_guide_order(slug, with_lines(slug, entries)))
    return [(name, short_sentences(line, name)) for name, line in ordered]
