"""Everything the boards draw, read from the real run and the active build. No figure is typed in.

The run: design/mocks/data/simview/fury-result.json (Fury warrior, level 60, raid buffs, 3:00,
3000 iterations, engine d5c708834). Names: spells.json, then data/curated/engine-spell-names.json,
then the class talent file, the same order sync-data.mjs folds them for simnames. Icons:
spellicons.json, then the talent file's own icon; ICON_GAPS lists the few the build carries for
neither (each is the client SpellIcon of that id, to be added to extra-spell-icons.json by the
data lane, spec section 11).
"""
import json
import math
import pathlib

ROOT = pathlib.Path(__file__).resolve().parents[3]
BUILD = json.loads((ROOT / "web/src/data/active-build.json").read_text())["build"]
DATA = ROOT / "data/builds" / BUILD
RUN = json.loads((ROOT / "design/mocks/data/simview/fury-result.json").read_text())
SUMMARY = RUN["summary"]
REQ = RUN["request"]
CLASS = REQ["character"]["class"]

SPELLS = {s["id"]: s["name"] for s in json.loads((DATA / "spells.json").read_text())}
ENGINE_NAMES = {int(k): v["name"] for k, v in json.loads((ROOT / "data/curated/engine-spell-names.json").read_text())["spells"].items()} \
    if "spells" in json.loads((ROOT / "data/curated/engine-spell-names.json").read_text()) else {}
SPELL_ICONS = json.loads((DATA / "spellicons.json").read_text())
CONST = json.loads((DATA / f"spellconst/{CLASS}.json").read_text())["spells"]
TALENTS = json.loads((DATA / f"talents/{CLASS}.json").read_text())
TALENT_SPELL = {}
for tree in TALENTS["trees"]:
    for t in tree["talents"]:
        for r in t["ranks"]:
            TALENT_SPELL[r["spell_id"]] = (t["name"], t["icon"])
ITEMS = {i["id"]: i for i in json.loads((DATA / f"items/{CLASS}.json").read_text())["items"]}
ENCHANTS = {e["id"]: e for e in json.loads((DATA / "enchants.json").read_text())}
SIMBUFFS = json.loads((DATA / "simbuffs.json").read_text())["entries"]
APL = json.loads((ROOT / f"data/curated/apl/{REQ['spec']}.json").read_text())

# The client SpellIcon of ids neither spellicons.json nor the talent file covers (spec 11, data lane).
ICON_GAPS = {12319: "ability_ghoulfrenzy", 20572: "racial_orc_berserkerstrength", 20007: "spell_holy_blessingofstrength",
             12867: "ability_backstab", 462891: "ability_whirlwind", 11585: "ability_meleedamage",
             9910: "spell_nature_thorns", 20906: "ability_trueshot", 20218: "spell_holy_mindvision", 10301: "spell_holy_auraoflight"}
# Names the build tables do not carry for these engine ids (spec 11, data lane: engine-spell-names.json).
NAME_GAPS = {462891: "Whirlwind (off hand)"}
HAND_TAG = {"1": "main hand", "2": "off hand"}


def icon_path(name):
    return DATA / "icons" / f"{name}.webp"


def parse_key(key):
    """'spell:11567/1' -> ('spell', 11567, '1')."""
    kind, _, rest = key.partition(":")
    base, _, tag = rest.partition("/")
    base = base.split("+")[0]
    return kind, (int(base) if base.isdigit() else base), tag


def weapon(slot):
    g = next(x for x in REQ["character"]["gear"] if x["slot"] == slot)
    return ITEMS[g["item_id"]]


def name_of(key):
    kind, sid, tag = parse_key(key)
    if kind == "other" and sid == "attack":
        return "Main hand" if tag == "1" else "Off hand"
    if kind != "spell":
        return key
    n = NAME_GAPS.get(sid) or ENGINE_NAMES.get(sid) or SPELLS.get(sid) or (TALENT_SPELL.get(sid) or (None,))[0]
    if n and tag in HAND_TAG and sid == 20007:
        n = f"{n} ({HAND_TAG[tag]})"
    return n or "An unnamed spell"


def icon_of(key):
    kind, sid, tag = parse_key(key)
    if kind == "other" and sid == "attack":
        return weapon("main_hand" if tag == "1" else "off_hand")["icon"]
    if kind != "spell":
        return None
    return SPELL_ICONS.get(str(sid)) or (TALENT_SPELL.get(sid) or (None, None))[1] or ICON_GAPS.get(sid)


def rank_of(key):
    kind, sid, _ = parse_key(key)
    row = CONST.get(str(sid)) if kind == "spell" else None
    return row["rank"] if row and row.get("rank") else 0


DUR_MS = SUMMARY["duration_ms"]
ACTOR = SUMMARY["damage_done"][0]
TOTAL = ACTOR["total"]
CASTS = {c["spell_name"]: c["succeeded"] for c in SUMMARY["casts"]}


def ability_rows():
    rows = []
    for a in ACTOR["abilities"]:
        if a["total"] <= 0:
            continue
        key = a["name"]
        misses = a.get("misses") or {}
        landed = a["hits"]
        attempts = landed + sum(v for k, v in misses.items() if k != "GLANCING")
        periodic = a["ticks"] > 0 and attempts == 0
        casts = CASTS.get(key, attempts)
        row = {
            "key": key, "name": name_of(key), "icon": icon_of(key), "rank": rank_of(key),
            "dps": a["total"] / (DUR_MS / 1000), "share": a["total"] / TOTAL, "total": a["total"],
            "casts": None if periodic else casts, "ticks": a["ticks"],
            "avg": a["total"] / (a["ticks"] if periodic else casts),
            "crit": a["crits"] / (a["ticks"] if periodic else attempts),
            "periodic": periodic, "white": key.startswith("other:attack"),
            "split": [(k.title() if k != "GLANCING" else "Glance", v / attempts) for k, v in
                      sorted(misses.items(), key=lambda kv: ["MISS", "DODGE", "PARRY", "GLANCING", "BLOCK"].index(kv[0]))] if attempts else [],
        }
        rows.append(row)
    rows.sort(key=lambda r: -r["dps"])
    return rows


SAMPLE = RUN["sample"]
SAMPLE_END = max(x["at_ms"] for x in SAMPLE)


def sample_times(key):
    return [x["at_ms"] for x in SAMPLE if x["action"] == key]


def rage_points():
    return [(x["at_ms"], x["resources"].get("rage", 0)) for x in SAMPLE if "resources" in x]


def duration_ms(key):
    _, sid, _ = parse_key(key)
    row = CONST.get(str(sid)) or {}
    return row.get("duration_ms") or 0


def merged(times, gap):
    segs = []
    for t in times:
        if segs and t - segs[-1][1] <= gap:
            segs[-1][1] = t
        else:
            segs.append([t, t])
    return segs


AURAS = [a for a in SUMMARY["auras"] if a["type"] == "BUFF"]


def aura_rows():
    perm, rows = [], []
    for a in AURAS:
        kind, sid, tag = parse_key(a["name"])
        if a["uptime_ms"] <= 0 or kind != "spell":
            continue
        if sid == 11567 and tag == "1":
            continue   # the Heroic Strike queue: a pressed button, drawn on its timeline lane, not a buff
        up = min(1.0, a["uptime_ms"] / DUR_MS)
        item = {"name": name_of(a["name"]), "icon": icon_of(a["name"]), "up": up, "n": a["applications"]}
        (perm if up >= 0.999 else rows).append(item)
    rows.sort(key=lambda r: -r["up"])
    return perm, rows


def talent_trees(string):
    out = []
    for tree, digits in zip(TALENTS["trees"], string.split("-")):
        cells = {(t["tier"], t["column"]): (t, int(d)) for t, d in zip(tree["talents"], digits)}
        out.append({"name": tree["name"], "bg": tree.get("background"), "points": sum(int(d) for d in digits),
                    "tiers": max(t["tier"] for t in tree["talents"]) + 1, "cells": cells})
    return out


def talents_valid(string):
    parts = string.split("-")
    if len(parts) != len(TALENTS["trees"]):
        return False
    for tree, digits in zip(TALENTS["trees"], parts):
        if len(digits) != len(tree["talents"]):
            return False
        if any(int(d) > t["max_rank"] for t, d in zip(tree["talents"], digits)):
            return False
    return True


def apl_lines():
    def cast_of(o):
        if isinstance(o, dict):
            if "castSpell" in o:
                return o["castSpell"]["spellId"]
            for v in o.values():
                r = cast_of(v)
                if r:
                    return r
        return None
    pre = []
    for p in APL["rotation"]["prepullActions"]:
        sid = cast_of(p)
        pre.append((sid, p["doAtValue"]["const"]["val"]))
    lines = []
    for p in APL["rotation"]["priorityList"]:
        sid = cast_of(p.get("action"))
        lines.append((sid, p.get("notes") or ""))
    return pre, lines


def histogram(bins=40):
    """Illustrative only: the engine's bins are new data (spec 11). Normal at the run's own mean and sd."""
    d = RUN["dps"]
    lo, hi = d["min"], d["max"]
    width = (hi - lo) / bins
    dens = []
    for i in range(bins):
        at = lo + width * (i + .5)
        z = (at - d["mean"]) / d["stddev"]
        dens.append(math.exp(-.5 * z * z))
    total = sum(dens)
    return [round(x / total * RUN["iterations_run"]) for x in dens], lo, hi, width
