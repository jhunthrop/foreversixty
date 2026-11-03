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
    "druid-restoration": {
        # restoration.md: "Swiftmend goes on the tank whenever he is under three quarters health
        # and your mana is keeping pace" (the rotation's own gate is under 75%).
        "Swiftmend": "Swiftmend on the tank whenever he is under three quarters health and your mana is keeping pace.",
    },
    "druid-feral": {
        # feral.md: "Shred fills every other global, building combo points the rest of the time."
        "Shred": "Shred fills every other global, building combo points.",
        # feral.md: "Press it every time it comes off cooldown, but only when your energy is 60 or lower".
        "Shifting Power": (
            "Shifting Power every time it comes off cooldown, but only at sixty energy or lower. "
            "Then the whole forty energy fits under the cap."
        ),
    },
    "druid-feral-bear": {
        # feral-bear.md priority items 2 and 3.
        "Lacerate": "Lacerate until it has five stacks, then again when it has under three seconds left.",
        "Faerie Fire": "Faerie Fire whenever the boss does not have it.",
    },
    "hunter-beast-mastery": {
        # beast-mastery.md: "Aimed Shot is only the opener: before the pull, Aspect of the Hawk goes up
        # and the first Aimed Shot begins casting"; "Multi-Shot is then cast every time the shared
        # cooldown comes up".
        "Aspect of the Hawk": "Aspect of the Hawk goes up before the pull.",
        "Aimed Shot": "Aimed Shot only as the opener, cast before the pull so it lands as the fight opens.",
        "Multi-Shot": "Multi-Shot every time the cooldown it shares with Aimed Shot comes up.",
        # beast-mastery.md: "Multi-Shot every six seconds, Serpent Sting when it is about to drop".
        "Serpent Sting": "Serpent Sting when it is about to drop.",
    },
    "hunter-marksmanship": {
        # marksmanship.md: Multi-Shot and Aimed Shot share a cooldown; "Aimed Shot is the opener only:
        # before the pull, Aspect of the Hawk goes up and the opening Aimed Shot begins casting".
        "Aspect of the Hawk": "Aspect of the Hawk goes up before the pull.",
        "Aimed Shot": "Aimed Shot only as the opener, cast before the pull so it lands as the fight opens.",
        "Multi-Shot": "Multi-Shot every time the cooldown it shares with Aimed Shot comes up.",
    },
    "hunter-survival": {
        # survival.md: "it can only be performed after you dodge, and Expose Prey opens the same window".
        # survival.md: Counterattack is "a free, unavoidable strike after a parry (cannot be blocked,
        # dodged, or parried)"; Lacerate "is recast whenever its own bleed has fallen off" (read 5).
        "Counterattack": "Counterattack only after you parry: it is free and cannot be blocked, dodged or parried.",
        "Lacerate": "Lacerate whenever its bleed has fallen off.",
        # survival.md: "Strider Kick comes next on its own 8-second cooldown, ahead of Immolation Trap".
        "Strider Kick": "Strider Kick on its own short cooldown, ahead of Immolation Trap.",
        "Mongoose Bite": "Mongoose Bite only after you dodge, or while Expose Prey is up: it is a counterattack.",
        # survival.md: "only in a raid: with a Windfury Totem in the group ... a loss without the totem".
        "Wing Clip": (
            "Wing Clip only with a Windfury Totem in the group, on a global with nothing else ready. "
            "Without the totem it loses damage."
        ),
    },
    "warlock-affliction": {
        # affliction.md: "Immolate is the fifth DoT ... recasting it when under 3 seconds remain";
        # "Life Tap keeps you casting once mana falls below about 10%".
        "Immolate": "Immolate is the fifth DoT, recast when under three seconds remain.",
        "Life Tap": "Life Tap keeps you casting once mana runs low.",
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
    "mage-arcane": {
        # arcane.md: "Otherwise cast Arcane Blast until its buff carries three stacks. Past three, keep
        # blasting while your mana is ahead of the fight".
        "Arcane Blast": (
            "Arcane Blast until its buff carries three stacks. "
            "Past three, keep blasting while your mana is ahead of the fight."
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
        # protection.md: "Templar's Bulwark first ... Divine Shield and Divine Protection come after it,
        # once Forbearance is gone."
        "Divine Shield": "Divine Shield and Divine Protection come after Templar's Bulwark, once Forbearance is gone.",
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
        # retribution.md (round 10, ranks lane): "cast it only while the mana left after it still
        # pays for your Seal, Judgement and Holy Strike, and while your mana is no more than about a
        # fifth behind the share of the fight still to go".
        "Consecration": (
            "Consecration last, while the mana left after it still pays for your seal, Judgement and Holy Strike. "
            "Do not let your mana fall far behind the fight."
        ),
    },
    "priest-discipline": {
        # discipline.md: "Major Mana Potion and Demonic Rune once you are down enough, Dark Sacrifice if you are Undead."
        "Dark Sacrifice": "Dark Sacrifice if you are Undead, once your mana is down enough.",
    },
    "priest-shadow": {
        "Shadow Word: Pain": "Shadow Word: Pain kept up, unless the fight is close to ending.",
        # shadow.md: "Inner Focus on Mind Blast every time it is up"; "Mind Blast on cooldown.
        # Shadow Word: Death on cooldown. Mind Flay with everything else"; "Nothing changes below 20%".
        "Inner Focus": "Inner Focus on Mind Blast every time it is up: the Mind Blast is free and more likely to crit.",
        "Mind Blast": "Mind Blast on cooldown, ahead of Mind Flay.",
        "Shadow Word: Death": (
            "Shadow Word: Death on cooldown, ahead of Mind Flay. "
            "Keep Vampiric Embrace up so your own damage heals back its backlash."
        ),
        "Mind Flay": "Mind Flay with everything else, and never clip it: the third tick is the biggest.",
        # shadow.md: "Wand shots fill any global where nothing else is ready."
        "Shoot": "Shoot with your wand on any global where nothing else is ready.",
        # shadow.md: "On long fights use a Major Mana Potion and a Demonic Rune once you are down that
        # much, and an Undead uses Dark Sacrifice."
        "Dark Sacrifice": "Dark Sacrifice if you are Undead, on a long fight once your mana is well down.",
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
        # elemental.md: "Spend it on Fire Nova only if your Fire totem is out and the boss is within
        # 10 yards of it, which from 20 yards it usually is not. Otherwise it goes to your next Lightning Bolt."
        "Fire Nova": (
            "Fire Nova on a Clearcasting proc only when your Fire totem is out and the boss stands close to it. "
            "Otherwise the proc goes to your next Lightning Bolt."
        ),
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
    # feral.md: Shifting Power "sits ahead of every finisher and builder in the order"; Rip at five
    # combo points, and Shred "fills every other global".
    "druid-feral": ("Shifting Power", "Rip", "Shred"),
    # feral-bear.md: Dire Bear Form and Enrage before the pull, then the numbered priority list.
    "druid-feral-bear": (
        "Dire Bear Form", "Enrage", "Primal Bite", "Lacerate", "Faerie Fire", "Maul", "Swipe",
    ),
    # beast-mastery.md: the opener, then Bestial Wrath with Rapid Fire, then Multi-Shot, Serpent
    # Sting and Arcane Shot last.
    "hunter-beast-mastery": (
        "Aspect of the Hawk", "Aimed Shot", "Bestial Wrath", "Rapid Fire", "Multi-Shot", "Serpent Sting",
        "Arcane Shot",
    ),
    "hunter-marksmanship": ("Aspect of the Hawk", "Aimed Shot"),
    # protection.md list items 1-6.
    "paladin-protection": (
        "Righteous Fury", "Seal of Fury", "Holy Shield", "Hammer of the Righteous", "Holy Strike",
        "Judgement", "Consecration",
    ),
    # shadow.md: "Mind Blast on cooldown. Shadow Word: Death on cooldown. Mind Flay with everything else".
    "priest-shadow": (
        "Shadow Word: Pain", "Inner Focus", "Mind Blast", "Shadow Word: Death", "Mind Flay", "Shoot",
    ),
    # enhancement.md: the imbue is "applied before the pull", ahead of every rotation line.
    "shaman-enhancement": ("Windfury Weapon",),
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
