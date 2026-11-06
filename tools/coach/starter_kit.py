"""What a starter says about preparation, resources and findings, per spec.

Consumables come from the spec's published raid preset (starter_food), checked against the
catalogue the coach reads, never a role-wide or resource-wide list. Drills are shared by role;
a spec whose guide plays a drill differently gives its own text in starter_overrides.DRILL_LINES.
"""

from __future__ import annotations

from dataclasses import dataclass
from typing import Any

from starter_food import band_labels, first_level, preset_food, preset_labels
from starter_overrides import DRILL_LINES
from starter_text import number_words

ROLE_LABEL = {"dps": "damage dealer", "healer": "healer", "tank": "tank"}
HEALER_CODES = ("core_casts", "uptime", "resource", "ranks", "preparation", "death")
ALL_CODES = ("core_casts", "uptime", "idle", "resource", "ranks", "preparation", "death", "adherence")
ROLE_CODES = {"dps": ALL_CODES, "tank": ALL_CODES, "healer": HEALER_CODES}

# Each spec's main resource; the resource drill names its bar.
RESOURCE: dict[str, str] = {
    "druid-balance": "mana",
    "druid-feral": "energy",
    "druid-feral-bear": "rage",
    "druid-restoration": "mana",
    "hunter-beast-mastery": "mana",
    "hunter-marksmanship": "mana",
    "hunter-survival": "mana",
    "mage-arcane": "mana",
    "mage-fire": "mana",
    "mage-frost": "mana",
    "paladin-holy": "mana",
    "paladin-protection": "mana",
    "paladin-retribution": "mana",
    "priest-discipline": "mana",
    "priest-shadow": "mana",
    "rogue-assassination": "energy",
    "rogue-combat": "energy",
    "rogue-subtlety": "energy",
    "shaman-elemental": "mana",
    "shaman-enhancement": "mana",
    "shaman-restoration": "mana",
    "warlock-affliction": "mana",
    "warlock-demonology": "mana",
    "warlock-destruction": "mana",
    "warrior-arms": "rage",
}
# The preset's consumables a starter names, in pull-checklist order, with the word for each.
KIT_KINDS = (("flask", "a flask"), ("weapon_enchant", "a weapon enchant"), ("combat_potion", "a combat potion"))
# The coach reads dungeon bosses from the low teens (Ragefire Chasm); a consumable that first arrives
# in a later band is named with its level (round 26), as the Fury persona does by hand.
EARLIEST_DUNGEON_LEVEL = 13


@dataclass(frozen=True)
class KitItem:
    """One thing on the pull checklist: the top band's item(s), the level those items arrive at,
    and the level the first item of the kind arrives at (a Dense stone before the Elemental)."""

    thing: str
    name: str
    name_level: int
    thing_level: int


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
    "uptime": "a buff, debuff or damage over time left to fall off when its line says to keep it up.",
    "idle": "time with no button pressed while the boss was alive and in reach.",
    "resource": "the main resource bar left to cap or to run dry.",
    "ranks": "a spell at an older rank than your level calls for, because the trainer visit was skipped.",
    "death": "dying, which ends your contribution for the rest of the pull.",
    "adherence": "buttons outside the priority list taking globals the list would have used.",
}
HEALER_MISTAKES = {
    "core_casts": "a mana or efficiency cooldown left unused through a long fight when its line says to press it.",
    "resource": "the mana bar empty well before the end of the fight.",
    "ranks": "a big rank landing mostly as overheal where the rank your level calls for would have covered the hole.",
    "death": "dying, because a dead healer heals nothing for the rest of the pull.",
}

RESOURCE_DRILL = {
    "rage": "Watch the rage bar on a dummy. Spend it on the list's buttons before it can cap, and never let a button wait for want of it.",
    "energy": "Watch the energy bar on a dummy. Spend it on the list's buttons before it can cap, and never let a button wait for want of it.",
    "mana": "Watch the mana bar against the fight. Know where your mana comes back from, and use each source before the bar is empty, not after.",
}
# Ranks: the guides keep a cheaper mana rank on purpose at some levels (round 25), so the drill never
# asks for the top rank on every button. Energy and rage ranks all cost the same (round 26), so
# those specs get no cheaper-rank sentence.
RANKS_DRILL = (
    "Visit your trainer every time you level, then check each button on the priority list against the rank "
    "your level calls for."
)
CHEAPER_RANK = " Sometimes the right one is a lower, cheaper rank."
FLAT_COST_RESOURCES = frozenset({"energy", "rage"})
HEALER_RESOURCE_DRILL = (
    "Before the pull, note the mana you should have at the halfway mark. If you are behind it, "
    "heal only the real holes and let the small ones wait."
)
HEALER_CORE_CASTS_DRILL = (
    "When the mana bar drops, check your mana cooldowns before the next cast. "
    "Press each one as soon as the bar has room for its full return."
)

DRILLS = {
    "dps": {
        "core_casts": "On a dummy, press the first button of the priority list the moment it is ready, before any filler. Read the list top to bottom and check it again after every cooldown.",
        "uptime": "Put the timers of your buffs and debuffs where you can see them. Recast each when its line says, never earlier, and again the moment you are back up after a death or a phase change.",
        "idle": "On every move, keep the boss in range and press a button the moment you arrive. Practise on trash: the first press after a reposition comes before any walk to the side.",
        "death": "Before the pull, find the spot where the boss does not cleave or breathe. Keep a healing potion on a key you can reach without looking.",
        "adherence": "On a dummy, play only the buttons on the priority list for three minutes. Anything else on your bars waits.",
    },
    "tank": {
        "core_casts": "On a dummy with a friend hitting you, press the first button of the priority list the moment it is ready. Fillers wait their turn.",
        "uptime": "Watch the debuffs on the boss and the buff on you. Refresh each where its line says, before it drops: a stacking debuff that falls off takes its stacks with it.",
        "idle": "When you move the boss, keep it in front of you and keep pressing buttons. Your first press after the move is a priority button, not a walk to the next spot.",
        "death": "Put your emergency cooldown on a key you can hit without looking, and press it on a dummy pull the moment your health falls past half.",
        "adherence": "On a dummy, play only the buttons on the priority list for three minutes. Anything else on your bars waits.",
    },
    "healer": {
        "uptime": "Recast your self buff as part of the ready check, and again the moment you are back up after a death.",
        "ranks": "Bind the ranks your level calls for and check them after every trainer visit. Save the top rank for a hole as deep as it heals.",
        "death": "Before the pull, pick a spot in range of the tank and out of anything the boss breathes or cleaves. When damage lands on you, move first and finish the cast after.",
    },
}


def finding_title(code: str, role: str) -> str:
    if role == "healer" and code in HEALER_FINDING_TITLES:
        return HEALER_FINDING_TITLES[code]
    return FINDING_TITLES[code]


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


def kit(slug: str, role: str, catalogue: dict[str, Any]) -> list[KitItem]:
    """What the spec brings in its raid preset, in pull-checklist order, each with its level.
    Only what the catalogue lists for the role is named: those are the things the coach can see."""
    food = checked(catalogue, "food", preset_food(slug))
    food_known = band_labels(slug, "food") & catalogue_names(catalogue, "food")
    found = [KitItem("food", food, first_level(slug, {food}), first_level(slug, food_known))]
    for kind, thing in KIT_KINDS:
        known = catalogue_names_for(catalogue, role, kind)
        names = [name for name in preset_labels(slug, kind) if name in known]
        if names:
            levels = (first_level(slug, set(names)), first_level(slug, band_labels(slug, kind) & known))
            found.append(KitItem(thing, join_names(names), *levels))
    flask = [item for item in found if item.thing == "a flask"]
    return flask + [item for item in found if item.thing != "a flask"]


def from_level(level: int) -> str:
    """' from <level>' when the item arrives after the first dungeons, else nothing."""
    return f" from {number_words(level)}" if level > EARLIEST_DUNGEON_LEVEL else ""


def catalogue_names_for(catalogue: dict[str, Any], role: str, kind: str) -> set[str]:
    entry = catalogue["roles"][role][kind]
    entries = entry["entries"] if isinstance(entry, dict) else entry
    return {e["name"] for e in entries}


def consumable_paragraph(slug: str, role: str, catalogue: dict[str, Any]) -> str:
    listed = "; ".join(f"{i.thing} ({i.name}){from_level(i.name_level)}" for i in kit(slug, role, catalogue))
    potion_limit = ", and the potion as often as its timer allows" if "potion" in listed else ""
    return (
        f"Before a boss you want {listed}{potion_limit}. "
        "World buffs are never expected: Forever raids have none, so a missing one is never held against anyone."
    )


def preparation_drill(slug: str, role: str, catalogue: dict[str, Any]) -> str:
    things = [item.thing + from_level(item.thing_level) for item in kit(slug, role, catalogue)]
    if role == "tank":
        things.append("a healing potion")
    return (
        f"Before every boss, check {join_names(things)} on your bars and in your buffs. "
        "Do it on the walk to the pull, not after it."
    )


def preparation_mistake(slug: str, role: str, catalogue: dict[str, Any]) -> str:
    things = [item.thing.removeprefix("a ") for item in kit(slug, role, catalogue)]
    return f"a missing {join_names(things)}."


def resource_drill(slug: str, role: str) -> str:
    if role == "healer":
        return HEALER_RESOURCE_DRILL
    try:
        return RESOURCE_DRILL[RESOURCE[slug]]
    except KeyError:
        raise SystemExit(f"no resource for {slug!r}: add it to starter_kit.RESOURCE") from None


def ranks_drill(slug: str) -> str:
    if RESOURCE.get(slug) in FLAT_COST_RESOURCES:
        return RANKS_DRILL
    return RANKS_DRILL + CHEAPER_RANK


def drills(slug: str, role: str, catalogue: dict[str, Any]) -> list[tuple[str, str]]:
    table = dict(DRILLS[role])
    if role != "healer":
        table["ranks"] = ranks_drill(slug)
    table["resource"] = resource_drill(slug, role)
    table["preparation"] = preparation_drill(slug, role, catalogue)
    if role == "healer":
        table["core_casts"] = HEALER_CORE_CASTS_DRILL
    table.update(DRILL_LINES.get(slug, {}))
    return [(code, table[code]) for code in ROLE_CODES[role]]


def mistakes(slug: str, role: str, catalogue: dict[str, Any]) -> list[str]:
    table = {**MISTAKES, **(HEALER_MISTAKES if role == "healer" else {})}
    table["preparation"] = preparation_mistake(slug, role, catalogue)
    return [f"- {finding_title(code, role)}: {table[code]}" for code in ROLE_CODES[role]]
