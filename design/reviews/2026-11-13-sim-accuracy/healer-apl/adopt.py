"""Applies the lane's adopted trial rotations (variants.py: paladin p2, shaman s5, priests h3,
druid d4) to data/curated/apl/<spec>.json, with each touched line's notes and panel text.
Run once from the repo root: python3 -I adopt.py <repo root>. `make apl-sync` then copies them.
"""
import json
import os
import sys

ROOT = sys.argv[1]
SELF, TANK = 0, 5
GROUP = (0, 1, 2, 3, 4)


def const(v):
    return {"const": {"val": str(v)}}


def hp(unit):
    return {"currentHealthPercent": {"sourceUnit": {"type": "Player", "index": unit}}}


def lt(lhs, v):
    return {"cmp": {"op": "OpLt", "lhs": lhs, "rhs": const(v)}}


def on_pace():
    return {"cmp": {"op": "OpGe", "lhs": {"currentManaPercent": {}},
                    "rhs": {"math": {"op": "OpMul", "lhs": {"remainingTimePercent": {}}, "rhs": const(1.0)}}}}


def not_ticking(spell, rank, unit):
    return {"not": {"val": {"dotIsActive": {"spellId": {"spellId": spell, "rank": rank},
                                             "targetUnit": {"type": "Player", "index": unit}}}}}


def cast(spell, rank, unit, cond, notes=None):
    line = {"action": {"condition": cond,
                       "castSpell": {"spellId": {"spellId": spell, "rank": rank},
                                     "target": {"type": "Player", "index": unit}}}}
    if notes:
        line["notes"] = notes
    return line


def set_hp(node, old, new):
    if isinstance(node, dict):
        c = node.get("cmp")
        if c and c.get("op") == "OpLt" and "currentHealthPercent" in c.get("lhs", {}) \
                and c.get("rhs", {}).get("const", {}).get("val") == str(old):
            c["rhs"] = const(new)
        for v in node.values():
            set_hp(v, old, new)
    elif isinstance(node, list):
        for v in node:
            set_hp(v, old, new)


def edit(spec, change):
    path = os.path.join(ROOT, "data", "curated", "apl", spec + ".json")
    with open(path) as f:
        doc = json.load(f)
    change(doc["rotation"]["priorityList"], doc["panel_when"])
    with open(path, "w") as f:
        f.write(json.dumps(doc, indent=2, ensure_ascii=False) + "\n")


def with_cond(line, extra):
    cond = line["action"]["condition"]
    if "and" in cond:
        cond["and"]["vals"].append(extra)
    else:
        line["action"]["condition"] = {"and": {"vals": [cond, extra]}}


def paladin(lines, when):
    set_hp(lines[11], 0.9, 0.97)
    lines[11]["notes"] = lines[11]["notes"].replace(
        "On the tank below 90%: your staple.",
        "On the tank as soon as one hit has landed on him: your staple, and in a raid the heal you cast most.")
    when["11"] = "on the tank below 97%"
    for i in (12, 13, 14, 15, 16):
        set_hp(lines[i], 0.8, 0.93)
        when[str(i)] = "on yourself below 93%" if i == 12 else "on a party member below 93%"
    lines[12]["notes"] = lines[12]["notes"].replace(
        "On everyone else below 80%, yourself included.",
        "On everyone else once a hit or two has landed, yourself included.")


def shaman(lines, when):
    set_hp(lines[6], 0.9, 0.95)
    when["6"] = "when the tank is below 95% with no Riptide"
    set_hp(lines[16], 0.9, 0.92)
    when["16"] = "when anyone is below 92% while mana % is over boss health %"
    lines[16]["notes"] = lines[16]["notes"].replace("under 90%", "under 92%")


def renew_when_hurt(lines, when, index):
    with_cond(lines[index], lt(hp(TANK), 0.95))
    when[str(index)] = "when Renew is not on the tank and he is below 95%"
    lines[index]["notes"] = ("Renew on the tank once he has taken a hit and it is not already on him; at 60 "
                             "that means rank 7 or higher. On a full tank every tick is wasted.")


def flash_top_off(lines, when):
    start = len(lines)
    lines.append(cast(10917, 7, TANK, {"and": {"vals": [lt(hp(TANK), 0.9), on_pace()]}},
                      "Flash Heal on the tank below 90% while your mana is level with the fight or ahead "
                      "of it: in a raid most of your casts are this one, keeping people topped between the big hits."))
    for u in GROUP:
        lines.append(cast(10917, 7, u, {"and": {"vals": [lt(hp(u), 0.85), on_pace()]}}))
    when[str(start)] = "on the tank below 90% while mana keeps pace with the fight"
    when[str(start + 1)] = "on yourself below 85% while mana keeps pace with the fight"
    for i in range(start + 2, start + 6):
        when[str(i)] = "on a party member below 85% while mana keeps pace with the fight"


def holy(lines, when):
    renew_when_hurt(lines, when, 3)
    flash_top_off(lines, when)


def discipline(lines, when):
    renew_when_hurt(lines, when, 8)
    flash_top_off(lines, when)


def druid(lines, when):
    with_cond(lines[7], lt(hp(TANK), 0.9))
    when["7"] = "when Regrowth is not ticking on the tank and he is below 90%"
    lines[7]["notes"] = lines[7]["notes"].replace(
        "On the tank whenever it is not ticking.", "On the tank once he is below 90% and it is not ticking.")
    with_cond(lines[8], lt(hp(TANK), 0.97))
    when["8"] = "when Rejuvenation is not ticking on the tank and he is below 97%"
    lines[8]["notes"] = lines[8]["notes"].replace(
        "On the tank whenever it is not ticking,", "On the tank whenever it is not ticking and he has taken a hit,")
    start = len(lines)
    lines.append(cast(8910, 7, SELF, {"and": {"vals": [lt(hp(SELF), 0.9), not_ticking(8910, 7, SELF), on_pace()]}},
                      "Rejuvenation on anyone in your party below 90%, you included, while your mana is level "
                      "with the fight or ahead of it: in a raid the group's small hits add up, and a ticking "
                      "Rejuvenation clears them before the next one lands. Same rank as the line above."))
    for u in (1, 2, 3, 4):
        lines.append(cast(8910, 7, u, {"and": {"vals": [lt(hp(u), 0.9), not_ticking(8910, 7, u), on_pace()]}}))
    when[str(start)] = "when you are under 90% with no Rejuvenation, while mana keeps pace with the fight"
    for i in range(start + 1, start + 5):
        when[str(i)] = "on a party member under 90% with no Rejuvenation, while mana keeps pace with the fight"


if __name__ == "__main__":
    edit("paladin-holy", paladin)
    edit("shaman-restoration", shaman)
    edit("priest-holy", holy)
    edit("priest-discipline", discipline)
    edit("druid-restoration", druid)
