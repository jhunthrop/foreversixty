"""Writes the lane's trial rotations: each is a curated healer rotation with its health thresholds
moved or top-off lines appended, as a -rotation file for spec-breakdown ({"rotation": ...}).

Usage: python3 -I variants.py <repo root> <out dir>
Units: the healing sim's raid is player 0 (the healer), 1 to 4 its party, 5 the tank.
"""
import copy
import json
import os
import sys

ROOT, OUT = sys.argv[1], sys.argv[2]
SELF, PARTY, TANK = 0, (1, 2, 3, 4), 5
GROUP = (0, 1, 2, 3, 4)


def const(v):
    return {"const": {"val": str(v)}}


def hp(unit):
    return {"currentHealthPercent": {"sourceUnit": {"type": "Player", "index": unit}}}


def lt(lhs, v):
    return {"cmp": {"op": "OpLt", "lhs": lhs, "rhs": const(v)}}


def on_pace(k):
    """Mana no further behind the fight clock than k of it (the lists' own guard)."""
    return {"cmp": {"op": "OpGe", "lhs": {"currentManaPercent": {}},
                    "rhs": {"math": {"op": "OpMul", "lhs": {"remainingTimePercent": {}}, "rhs": const(k)}}}}


def and_(*vals):
    return {"and": {"vals": list(vals)}}


def not_ticking(spell, rank, unit):
    return {"not": {"val": {"dotIsActive": {"spellId": {"spellId": spell, "rank": rank},
                                             "targetUnit": {"type": "Player", "index": unit}}}}}


def cast(spell, rank, unit, cond):
    return {"action": {"condition": cond,
                       "castSpell": {"spellId": {"spellId": spell, "rank": rank},
                                     "target": {"type": "Player", "index": unit}}}}


def load(spec):
    with open(os.path.join(ROOT, "data", "curated", "apl", spec + ".json")) as f:
        return json.load(f)["rotation"]


def set_hp(rot, lines, old, new):
    """Moves every `currentHealthPercent < old` test on the given lines to `< new`."""
    def walk(node):
        if isinstance(node, dict):
            c = node.get("cmp")
            if c and c.get("op") == "OpLt" and "currentHealthPercent" in c.get("lhs", {}) \
                    and c.get("rhs", {}).get("const", {}).get("val") == str(old):
                c["rhs"] = const(new)
            for v in node.values():
                walk(v)
        elif isinstance(node, list):
            for v in node:
                walk(v)
    for i in lines:
        walk(rot["priorityList"][i])
    return rot


def write(name, rot):
    with open(os.path.join(OUT, name + ".json"), "w") as f:
        json.dump({"rotation": rot}, f, indent=1)


def top_off(spell, rank, tank_at, group_at, pace):
    return [cast(spell, rank, TANK, and_(lt(hp(TANK), tank_at), on_pace(pace)))] + \
        [cast(spell, rank, u, and_(lt(hp(u), group_at), on_pace(pace))) for u in GROUP]


def paladin():
    for tag, tank, group in (("p1", 0.95, 0.9), ("p2", 0.97, 0.93)):
        rot = load("paladin-holy")
        set_hp(rot, [11], 0.9, tank)
        set_hp(rot, [12, 13, 14, 15, 16], 0.8, group)
        write("paladin-holy-" + tag, rot)


def shaman():
    rot = load("shaman-restoration")
    set_hp(rot, [16], 0.9, 0.95)
    set_hp(rot, [6], 0.9, 0.95)
    write("shaman-restoration-s1", rot)
    rot = copy.deepcopy(rot)
    rot["priorityList"] += top_off(10468, 6, 0.9, 0.9, 0.9)
    write("shaman-restoration-s2", rot)
    for tag, chain in (("s3", 0.93), ("s4", 0.92), ("s5", 0.92)):
        rot = load("shaman-restoration")
        set_hp(rot, [16], 0.9, chain)
        set_hp(rot, [6], 0.9, 0.95)
        if tag == "s4":
            rot["priorityList"] += top_off(10468, 6, 0.9, 0.9, 1.0)
        write("shaman-restoration-" + tag, rot)


def priests():
    for spec in ("priest-holy", "priest-discipline"):
        rot = load(spec)
        rot["priorityList"] += top_off(10917, 7, 0.9, 0.85, 0.9)
        write(spec + "-h1", rot)
        rot = copy.deepcopy(rot)
        rot["priorityList"] += [cast(10927, 7, u, and_(lt(hp(u), 0.9), not_ticking(10927, 7, u), on_pace(0.9)))
                                for u in GROUP]
        write(spec + "-h2", rot)
        # h3: the tank's Renew only once he is hurt, and the Flash Heal top-off on pace.
        rot = load(spec)
        renew = next(line for line in rot["priorityList"]
                     if line["action"].get("castSpell", {}).get("spellId", {}).get("spellId") == 10927)
        renew["action"]["condition"] = and_(renew["action"]["condition"], lt(hp(TANK), 0.95))
        rot["priorityList"] += top_off(10917, 7, 0.9, 0.85, 1.0)
        write(spec + "-h3", rot)


def druid():
    rot = load("druid-restoration")
    set_hp(rot, [10, 11, 12, 13, 14], 0.7, 0.9)
    write("druid-restoration-d1", rot)
    rot = copy.deepcopy(rot)
    rot["priorityList"][7]["action"]["condition"]["and"]["vals"].append(lt(hp(TANK), 0.9))
    write("druid-restoration-d2", rot)
    # d3: the old lines kept; the tank's Regrowth once he is hurt, Rejuvenation on the group on pace.
    rot = load("druid-restoration")
    rot["priorityList"][7]["action"]["condition"]["and"]["vals"].append(lt(hp(TANK), 0.9))
    rot["priorityList"] += [cast(8910, 7, u, and_(lt(hp(u), 0.9), not_ticking(8910, 7, u), on_pace(1.0)))
                            for u in GROUP]
    write("druid-restoration-d3", rot)
    # d4: d3, and the tank's Rejuvenation once he is hurt as well.
    rot = copy.deepcopy(rot)
    rot["priorityList"][8]["action"]["condition"]["and"]["vals"].append(lt(hp(TANK), 0.97))
    write("druid-restoration-d4", rot)


if __name__ == "__main__":
    os.makedirs(OUT, exist_ok=True)
    paladin()
    shaman()
    priests()
    druid()
