"""What a starter says about preparation, resources and findings, per spec.

Consumables come from the catalogue, filtered by the spec's own traits, never a role-wide list.
The guides' consumables sections are empty, so the traits table is the only per-spec input.
"""

from __future__ import annotations

from dataclasses import dataclass
from typing import Any

ROLE_LABEL = {"dps": "damage dealer", "healer": "healer", "tank": "tank"}
HEALER_CODES = ("core_casts", "uptime", "resource", "ranks", "preparation", "death")
ALL_CODES = ("core_casts", "uptime", "idle", "resource", "ranks", "preparation", "death", "adherence")
ROLE_CODES = {"dps": ALL_CODES, "tank": ALL_CODES, "healer": HEALER_CODES}

FLASK_BY_STYLE = {
    "caster": "Flask of Supreme Power",
    "healer": "Flask of Distilled Wisdom",
    "melee": "Flask of the Titans",
    "ranged": "Flask of the Titans",
    "tank": "Flask of the Titans",
}
OIL = "Brilliant Wizard Oil"
STONE = "Elemental Sharpening Stone"
RAGE_POTION = "Mighty Rage Potion"
MANA_POTION = "Major Mana Potion"


@dataclass(frozen=True)
class Traits:
    resource: str  # rage, energy or mana
    style: str  # caster, melee, ranged, healer or tank
    stone: bool = False  # an edged melee weapon that takes a sharpening stone


# Facets of each starter spec. specs.json carries the role and stats; resource and weapon use are here.
TRAITS: dict[str, Traits] = {
    "druid-balance": Traits("mana", "caster"),
    "druid-feral": Traits("energy", "melee"),
    "druid-feral-bear": Traits("rage", "tank"),
    "druid-restoration": Traits("mana", "healer"),
    "hunter-beast-mastery": Traits("mana", "ranged"),
    "hunter-marksmanship": Traits("mana", "ranged"),
    "hunter-survival": Traits("mana", "melee"),
    "mage-arcane": Traits("mana", "caster"),
    "mage-fire": Traits("mana", "caster"),
    "mage-frost": Traits("mana", "caster"),
    "paladin-holy": Traits("mana", "healer"),
    "paladin-protection": Traits("mana", "tank"),
    "paladin-retribution": Traits("mana", "melee"),
    "priest-discipline": Traits("mana", "healer"),
    "priest-shadow": Traits("mana", "caster"),
    "rogue-assassination": Traits("energy", "melee"),
    "rogue-combat": Traits("energy", "melee"),
    "rogue-subtlety": Traits("energy", "melee"),
    "shaman-elemental": Traits("mana", "caster"),
    "shaman-enhancement": Traits("mana", "melee"),
    "shaman-restoration": Traits("mana", "healer"),
    "warlock-affliction": Traits("mana", "caster"),
    "warlock-demonology": Traits("mana", "caster"),
    "warlock-destruction": Traits("mana", "caster"),
    "warrior-arms": Traits("rage", "melee", stone=True),
}

# The names a player knows each finding by (design spec section 6), not the raw ids.
FINDING_TITLES = {
    "core_casts": "Core abilities under-used",
    "uptime": "Key buff and debuff uptime",
    "idle": "Dead time",
    "resource": "Resource waste",
    "ranks": "Ranks",
    "preparation": "Pull preparation",
    "death": "Death review",
    "adherence": "Rotation adherence",
}
HEALER_FINDING_TITLES = {"resource": "Mana"}

MISTAKES = {
    "core_casts": "a button the priority list wants sitting ready while something lower went first.",
    "uptime": "a buff, debuff or damage over time left to fall off the target.",
    "idle": "time with no button pressed while the boss was alive and in reach.",
    "resource": "the main resource bar left to cap or to run dry.",
    "ranks": "a spell left at an old rank because the trainer visit was skipped.",
    "death": "dying, which ends your contribution for the rest of the pull.",
    "adherence": "buttons outside the priority list taking globals the list would have used.",
}
HEALER_MISTAKES = {
    "core_casts": "a mana or efficiency cooldown sitting ready for most of the fight.",
    "resource": "the mana bar empty well before the end of the fight.",
    "ranks": "a big rank landing mostly as overheal when a lower rank of the same heal would have covered the hole.",
    "death": "dying, because a dead healer heals nothing for the rest of the pull.",
}

RESOURCE_DRILL = {
    "rage": "Watch the rage bar on a dummy. Spend it on the list's buttons before it can cap, and never let a button wait for want of it.",
    "energy": "Watch the energy bar on a dummy. Spend it on the list's buttons before it can cap, and never let a button wait for want of it.",
    "mana": "Watch the mana bar against the fight. Know where your mana comes back from, and use each source before the bar is empty, not after.",
}
HEALER_RESOURCE_DRILL = (
    "Before the pull, note the mana you should have at the halfway mark. If you are behind it, "
    "lean on your cheapest heal and save the fast, dear heal for the real emergency."
)
HEALER_CORE_CASTS_DRILL = (
    "When the mana bar drops below half, check your mana cooldowns before the next cast. "
    "Press each one as soon as the bar has room for its full return."
)

DRILLS = {
    "dps": {
        "core_casts": "On a dummy, press the first button of the priority list the moment it is ready, before any filler. Read the list top to bottom and check it again after every cooldown.",
        "uptime": "Put the timers of your buffs and debuffs where you can see them. Recast each with 3 seconds left, and again the moment you are back up after a death or a phase change.",
        "idle": "On every move, keep the boss in range and press a button the moment you arrive. Practise on trash: the first press after a reposition comes before any walk to the side.",
        "ranks": "Visit your trainer every time you level and read your spellbook after: every button on the priority list should show its top rank.",
        "death": "Before the pull, find the spot where the boss does not cleave or breathe. Keep a healing potion on a key you can reach without looking.",
        "adherence": "On a dummy, play only the buttons on the priority list for 3 minutes. Anything else on your bars waits.",
    },
    "tank": {
        "core_casts": "On a dummy with a friend hitting you, press the first button of the priority list the moment it is ready. Fillers wait their turn.",
        "uptime": "Watch the buff on you and the debuffs on the boss, not the cooldowns. The moment one drops, press it again.",
        "idle": "When you move the boss, keep it in front of you and keep pressing buttons. Your first press after the move is a priority button, not a walk to the next spot.",
        "ranks": "Visit your trainer every time you level and read your spellbook after: every button on the priority list should show its top rank.",
        "death": "Put your emergency cooldown on a key you can hit without looking, and press it on a dummy pull the moment your health falls past half.",
        "adherence": "On a dummy, play only the buttons on the priority list for 3 minutes. Anything else on your bars waits.",
    },
    "healer": {
        "uptime": "Recast your self buff as part of the ready check, and again the moment you are back up after a death.",
        "ranks": "Bind a mid rank of your filler heal and use it for every small hole on the tank. Save the top rank for a hole as deep as it heals.",
        "death": "Before the pull, pick a spot in range of the tank and out of anything the boss breathes or cleaves. When damage lands on you, move first and finish the cast after.",
    },
}


def finding_title(code: str, role: str) -> str:
    if role == "healer" and code in HEALER_FINDING_TITLES:
        return HEALER_FINDING_TITLES[code]
    return FINDING_TITLES[code]


def traits_of(slug: str) -> Traits:
    try:
        return TRAITS[slug]
    except KeyError:
        raise SystemExit(f"no traits for {slug!r}: add it to starter_kit.TRAITS") from None


def join_names(names: list[str]) -> str:
    if len(names) <= 1:
        return "".join(names)
    return ", ".join(names[:-1]) + " and " + names[-1]


def catalogue_names(catalogue: dict[str, Any], kind: str) -> set[str]:
    names: set[str] = set()
    for role in catalogue["roles"].values():
        entry = role[kind]
        entries = entry["entries"] if isinstance(entry, dict) else entry
        names.update(e["name"] for e in entries)
    return names


def checked(catalogue: dict[str, Any], kind: str, name: str) -> str:
    """The name, verified against the catalogue so a rename fails loudly."""
    if name not in catalogue_names(catalogue, kind):
        raise SystemExit(f"the catalogue has no {kind} named {name!r}")
    return name


def kit(slug: str, role: str, catalogue: dict[str, Any]) -> list[tuple[str, str]]:
    """(thing, catalogue name) for what the catalogue looks for on this spec, in pull-checklist order."""
    traits = traits_of(slug)
    found = [("a flask", checked(catalogue, "flask", FLASK_BY_STYLE[traits.style]))]
    if traits.style != "caster":
        found.append(("food", catalogue["roles"][role]["food"][0]["name"]))
    if traits.style == "caster":
        found.append(("a weapon enchant", checked(catalogue, "weapon_enchant", OIL)))
    elif traits.stone:
        found.append(("a weapon enchant", checked(catalogue, "weapon_enchant", STONE)))
    potion = {"rage": RAGE_POTION, "mana": MANA_POTION}.get(traits.resource)
    if potion:
        found.append(("a combat potion", checked(catalogue, "combat_potion", potion)))
    return found


def consumable_paragraph(slug: str, role: str, catalogue: dict[str, Any]) -> str:
    listed = "; ".join(f"{thing} ({name})" for thing, name in kit(slug, role, catalogue))
    buffs = join_names([e["name"] for e in catalogue["roles"][role]["world_buffs"]])
    potion_limit = ", the potion up to its use limit in a fight" if "potion" in listed else ""
    return (
        f"For this spec the catalogue looks for {listed}{potion_limit}. "
        f"The world buffs it lists are {buffs}. They are listed as facts, never held against anyone."
    )


def preparation_drill(slug: str, role: str, catalogue: dict[str, Any]) -> str:
    things = [thing for thing, _ in kit(slug, role, catalogue)]
    if role == "tank":
        things.append("a healing potion")
    return (
        f"Before every boss, check {join_names(things)} on your bars and in your buffs. "
        "Do it on the walk to the pull, not after it."
    )


def preparation_mistake(slug: str, role: str, catalogue: dict[str, Any]) -> str:
    things = [thing.removeprefix("a ") for thing, _ in kit(slug, role, catalogue)]
    return f"a missing {join_names(things)}."


def resource_drill(slug: str, role: str) -> str:
    return HEALER_RESOURCE_DRILL if role == "healer" else RESOURCE_DRILL[traits_of(slug).resource]


def drills(slug: str, role: str, catalogue: dict[str, Any]) -> list[tuple[str, str]]:
    table = dict(DRILLS[role])
    table["resource"] = resource_drill(slug, role)
    table["preparation"] = preparation_drill(slug, role, catalogue)
    if role == "healer":
        table["core_casts"] = HEALER_CORE_CASTS_DRILL
    return [(code, table[code]) for code in ROLE_CODES[role]]


def mistakes(slug: str, role: str, catalogue: dict[str, Any]) -> list[str]:
    table = {**MISTAKES, **(HEALER_MISTAKES if role == "healer" else {})}
    table["preparation"] = preparation_mistake(slug, role, catalogue)
    return [f"- {finding_title(code, role)}: {table[code]}" for code in ROLE_CODES[role]]
