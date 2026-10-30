"""The mock's data, read from data.json (build 1.60.1.70532, Elemental Shaman, level 60 Phase 1 raid, Horde)
plus a hand-made guild roster. Every item, icon, quality, item level, source and DPS delta is the build's own;
the example character wears listed alternatives, so each gain is the band's measured delta."""
from kit import DATA, item

SLOT_ORDER = ["head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs", "feet",
              "finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged"]
SLOT_LABEL = {"head": "Head", "neck": "Neck", "shoulder": "Shoulder", "back": "Back", "chest": "Chest", "wrist": "Wrist",
              "hands": "Hands", "waist": "Waist", "legs": "Legs", "feet": "Feet", "finger1": "Ring 1", "finger2": "Ring 2",
              "trinket1": "Trinket 1", "trinket2": "Trinket 2", "main_hand": "Main hand", "off_hand": "Off hand", "ranged": "Ranged"}

# What Lightningbolltz wears where it is not the pick: each id is one of the slot's listed alternatives.
WORN = {"head": 231663, "shoulder": 231659, "back": 18510, "wrist": 254123, "waist": 18405,
        "finger1": 272407, "finger2": 13001, "trinket2": 272438}


def source_line(slot: dict) -> str:
    kind, src = slot.get("source_kind"), slot.get("source", "")
    if kind == "dungeon" and ":" in src:
        inst, boss = (p.strip() for p in src.split(":", 1))
        return f"{inst} · {boss}"
    if kind == "crafted":
        return f"Crafted: {src}"
    if kind == "vendor":
        return f"Vendor: {src}"
    if kind == "quest":
        return f"Quest: {src}"
    if kind == "world_drop":
        return "World drop (BoE)"
    return src


def band(preset: str = "raid") -> dict:
    return DATA["bands"][preset]


def slots(preset: str = "raid") -> dict:
    return {s["slot"]: s for s in band(preset)["slots"] if "item_id" in s}


def upgrades() -> tuple[list[dict], list[str], float]:
    """The bot's /char order: measured gains high to low, then the picks whose gain is not modelled."""
    measured, unmodelled, already = [], [], []
    for key in SLOT_ORDER:
        s = slots()[key]
        worn = WORN.get(key, s["item_id"])
        if worn == s["item_id"]:
            already.append(SLOT_LABEL[key])
            continue
        alt = next(a for a in s["alternatives"] if a["item_id"] == worn)
        row = {"slot": key, "pick": s["item_id"], "worn": worn, "source": source_line(s),
               "gain": None if s.get("effect_unmodelled") else round(-alt["dps_delta"], 1)}
        (unmodelled if row["gain"] is None else measured).append(row)
    measured.sort(key=lambda r: -r["gain"])
    total = round(sum(r["gain"] for r in measured), 1)
    return measured + unmodelled, already, total


def next_best(key: str) -> list[tuple[int, str, float]]:
    s = slots()[key]
    return [(a["item_id"], a["item_name"], a["dps_delta"]) for a in s.get("alternatives", [])[:2]]


def stat_lines(item_id: int) -> list[str]:
    labels = {"intellect": "Intellect", "spirit": "Spirit", "stamina": "Stamina", "spell_power": "Spell Power", "hit": "Hit",
              "crit": "Crit", "healing": "Healing"}
    it = item(item_id)
    order = ["stamina", "intellect", "spirit", "hit", "crit", "spell_power", "healing"]
    return [f"+{it['stats'][k]} {labels[k]}" for k in order if k in it["stats"]]


SPEC = {  # slug: (class, name, role, icon)
    "warrior-protection": ("warrior", "Protection", "Tank", "inv_shield_06"),
    "warrior-fury": ("warrior", "Fury", "DPS", "ability_warrior_innerrage"),
    "warrior-arms": ("warrior", "Arms", "DPS", "ability_rogue_eviscerate"),
    "druid-feral-bear": ("druid", "Feral Bear", "Tank", "ability_racial_bearform"),
    "druid-restoration": ("druid", "Restoration", "Healer", "spell_nature_healingtouch"),
    "druid-balance": ("druid", "Balance", "DPS", "spell_nature_starfall"),
    "priest-holy": ("priest", "Holy", "Healer", "spell_holy_guardianspirit"),
    "priest-discipline": ("priest", "Discipline", "Healer", "spell_holy_wordfortitude"),
    "priest-shadow": ("priest", "Shadow", "DPS", "spell_shadow_shadowwordpain"),
    "shaman-restoration": ("shaman", "Restoration", "Healer", "spell_nature_magicimmunity"),
    "shaman-elemental": ("shaman", "Elemental", "DPS", "spell_nature_lightning"),
    "shaman-enhancement": ("shaman", "Enhancement", "DPS", "spell_nature_lightningshield"),
    "mage-frost": ("mage", "Frost", "DPS", "spell_frost_frostbolt02"),
    "mage-fire": ("mage", "Fire", "DPS", "spell_fire_firebolt02"),
    "mage-arcane": ("mage", "Arcane", "DPS", "spell_holy_magicalsentry"),
    "rogue-combat": ("rogue", "Combat", "DPS", "ability_backstab"),
    "rogue-subtlety": ("rogue", "Subtlety", "DPS", "ability_stealth"),
    "hunter-marksmanship": ("hunter", "Marksmanship", "DPS", "ability_marksmanship"),
    "hunter-survival": ("hunter", "Survival", "DPS", "ability_hunter_swiftstrike"),
    "hunter-beast-mastery": ("hunter", "Beast Mastery", "DPS", "ability_hunter_beasttaming"),
    "warlock-destruction": ("warlock", "Destruction", "DPS", "spell_shadow_rainoffire"),
    "warlock-affliction": ("warlock", "Affliction", "DPS", "spell_shadow_deathcoil"),
    "warlock-demonology": ("warlock", "Demonology", "DPS", "spell_shadow_metamorphosis"),
}

ROSTER = {
    "Tanks": [("Kraggor", "warrior-protection", ""), ("Thokhan", "warrior-protection", ""), ("Barkhide", "druid-feral-bear", "")],
    "Healers": [("Mendara", "priest-holy", ""), ("Velaine", "priest-discipline", ""), ("Gravemend", "priest-holy", ""),
                ("Rainspeaker", "shaman-restoration", ""), ("Tidecaller", "shaman-restoration", "Can swap to Elemental"), ("Oakenmother", "druid-restoration", "")],
    "DPS": [("Draknar", "warrior-fury", ""), ("Ruzga", "warrior-fury", ""), ("Ashfang", "hunter-marksmanship", ""), ("Talonmaw", "hunter-survival", ""),
            ("Skarn", "rogue-combat", ""), ("Zekra", "rogue-combat", ""), ("Grimshade", "priest-shadow", ""), ("Lightningbolltz", "shaman-elemental", ""),
            ("Stormjaw", "shaman-enhancement", ""), ("Vexxa", "mage-frost", ""), ("Kelthis", "mage-fire", ""), ("Morbith", "warlock-destruction", ""),
            ("Hexlash", "warlock-affliction", ""), ("Nightbloom", "druid-balance", "")],
    "Late": [("Bonecrank", "warlock-demonology", "Work till 8:30")],
    "Bench": [("Pyrewick", "mage-arcane", ""), ("Fenrik", "hunter-beast-mastery", "")],
    "Out": [("Ozzak", "rogue-subtlety", "Away this week"), ("Murkfin", "shaman-elemental", ""), ("Ruskar", "warrior-arms", "")],
}
NOT_SIGNED = [("Fenna", "priest"), ("Ugnar", "warrior"), ("Selith", "mage"), ("Korvash", "warlock"), ("Jinxa", "rogue")]
CLASS_ORDER = ["warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid"]


def class_counts() -> dict:
    counts = {c: 0 for c in CLASS_ORDER}
    for group in ("Tanks", "Healers", "DPS", "Late"):
        for _, spec, _ in ROSTER[group]:
            counts[SPEC[spec][0]] += 1
    return counts
