"""Everything the boards draw, read from the real run, the hand fixture and the active build.

The run: design/mocks/data/simview/fury-result.json, a real forever-sim run of the published
warrior-fury level 60 raid BiS entry (fury-request.json): troll, the guide talents, the raid preset
with Windfury Totem and consumables, 10,000 runs, 904.4 DPS ± 0.49.
The new fields the engine does not return yet (histogram, resource sources, sample aura segments,
sample outcomes, target debuff tracks, created_at) come from web/src/fixtures/sim/result-full.json,
the hand fixture the web lanes build against, so the mock and the lanes agree (spec 11.2).

Names: spells.json, then data/curated/engine-spell-names.json, then the class talent file, the order
sync-data.mjs folds them for simnames; items from items.json. Icons: spellicons.json, then the talent
file, then simbuffs.json for a buff or consumable. ICON_GAPS lists ids none of those carry an icon for;
each is the client SpellIcon or the simbuffs icon of the same buff. MISSING_ICONS are Forever-new ids
the repo has no icon for at all: they draw the skeleton box and are on the data lane's list (spec 11.4).
"""
import json
import pathlib

ROOT = pathlib.Path(__file__).resolve().parents[3]
BUILD = json.loads((ROOT / "web/src/data/active-build.json").read_text())["build"]
DATA = ROOT / "data/builds" / BUILD
RUN = json.loads((ROOT / "design/mocks/data/simview/fury-result.json").read_text())
FIX = json.loads((ROOT / "web/src/fixtures/sim/result-full.json").read_text())
SUMMARY = RUN["summary"]
REQ = RUN["request"]
CLASS = REQ["character"]["class"]

SPELLS = {s["id"]: s["name"] for s in json.loads((DATA / "spells.json").read_text())}
ENGINE_NAMES = {int(k): v["name"] for k, v in json.loads((ROOT / "data/curated/engine-spell-names.json").read_text())["spells"].items()}
SPELL_ICONS = json.loads((DATA / "spellicons.json").read_text())
CONST = json.loads((DATA / f"spellconst/{CLASS}.json").read_text())["spells"]
TALENTS = json.loads((DATA / f"talents/{CLASS}.json").read_text())
TALENT_SPELL = {}
for tree in TALENTS["trees"]:
    for t in tree["talents"]:
        for r in t["ranks"]:
            TALENT_SPELL[r["spell_id"]] = (t["name"], t["icon"])
ITEMS = {i["id"]: i for i in json.loads((DATA / f"items/{CLASS}.json").read_text())["items"]}
ALL_ITEMS = {i["id"]: i["name"] for i in json.loads((DATA / "items.json").read_text())}
ENCHANTS = {e["id"]: e for e in json.loads((DATA / "enchants.json").read_text())}
SIMBUFFS = json.loads((DATA / "simbuffs.json").read_text())["entries"]
SIMBUFF_ICON_BY_NAME = {v["name"]: v["icon"] for v in SIMBUFFS.values()}
APL = json.loads((ROOT / f"data/curated/apl/{REQ['spec']}.json").read_text())

ICON_GAPS = {12319: "ability_ghoulfrenzy", 20554: "racial_troll_berserk", 12867: "ability_backstab",
             462891: "ability_whirlwind", 11585: "ability_meleedamage", 20906: "ability_trueshot",
             10610: "spell_nature_windfury", 409583: "ability_hunter_pet_cat", 11722: "spell_shadow_chilltouch"}
MISSING_ICONS = {1249110, 21919}     # Chthonic Power, Thrash: no icon anywhere in the repo
NAME_GAPS = {462891: "Whirlwind (off hand)", 11722: "Curse of the Elements"}
HANDS = {"1": "Main hand", "2": "Off hand", "3": "Extra attacks"}


def icon_path(name):
    return DATA / "icons" / f"{name}.webp"


def parse_key(key):
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
        return HANDS[tag]
    if kind == "item":
        return ALL_ITEMS.get(sid, "An unnamed item")
    if kind != "spell":
        return key
    return NAME_GAPS.get(sid) or ENGINE_NAMES.get(sid) or SPELLS.get(sid) or (TALENT_SPELL.get(sid) or (None,))[0] or "An unnamed spell"


def icon_of(key):
    kind, sid, tag = parse_key(key)
    if kind == "other" and sid == "attack":
        return weapon("off_hand" if tag == "2" else "main_hand")["icon"]
    if kind == "item":
        return SIMBUFF_ICON_BY_NAME.get(ALL_ITEMS.get(sid))
    if kind != "spell" or sid in MISSING_ICONS:
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
DEBUFFS = [a for a in FIX["summary"]["auras"] if a["type"] == "DEBUFF"]          # ILLUSTRATIVE (fixture)


def ability_rows():
    """Avg hit = damage / landed hits (crits and glances are inside `hits`, adapter.go ability()).
    Crit = crits / landed hits, glances counted as landed (Warcraft Logs' Crit %; spec 3.6 says why not
    non-glancing). Splits are shares of attempts. A periodic row is per tick."""
    rows = []
    for a in ACTOR["abilities"]:
        if a["total"] <= 0:
            continue
        key = a["name"]
        misses = a.get("misses") or {}
        landed = a["hits"]
        glances = misses.get("GLANCING", 0)
        attempts = landed + sum(v for k, v in misses.items() if k != "GLANCING")
        periodic = a["ticks"] > 0 and attempts == 0
        _, sid, _ = parse_key(key)
        up = next((d["uptime_ms"] / DUR_MS for d in DEBUFFS if d["spell_id"] == sid), None)
        rows.append({
            "key": key, "name": name_of(key), "icon": icon_of(key), "rank": rank_of(key),
            "dps": a["total"] / (DUR_MS / 1000), "share": a["total"] / TOTAL, "total": a["total"],
            "casts": None if periodic else CASTS.get(key, attempts), "ticks": a["ticks"],
            "avg": a["total"] / (a["ticks"] if periodic else landed),
            "crit": a["crits"] / (a["ticks"] if periodic else landed),
            "periodic": periodic, "white": key.startswith("other:attack"), "uptime": min(1.0, up) if up else None,
            "split": [(k.title() if k != "GLANCING" else "Glance", v / attempts) for k, v in
                      sorted(misses.items(), key=lambda kv: ["MISS", "DODGE", "PARRY", "GLANCING", "BLOCK"].index(kv[0]))] if attempts else [],
        })
    rows.sort(key=lambda r: -r["dps"])
    return rows


SAMPLE = FIX["sample"]                 # the run's own events, plus the fixture's illustrative outcome and amount
SAMPLE_END = FIX["sample_duration_ms"]
SAMPLE_AURAS = FIX["sample_auras"]     # ILLUSTRATIVE (fixture)


def sample_times(key):
    return [x["at_ms"] for x in SAMPLE if x["action"] == key]


def sample_events(key):
    return [x for x in SAMPLE if x["action"] == key]


def rage_points():
    return [(x["at_ms"], x["resources"].get("rage", 0)) for x in SAMPLE if "resources" in x]


def segments(key, on):
    return next((a["segments"] for a in SAMPLE_AURAS if a["action"] == key and a["on"] == on), [])


ROTATION_CASTS_OVERPOWER = CASTS.get("spell:11585", 0) > 0


def aura_rows():
    perm, rows = [], []
    for a in SUMMARY["auras"]:
        kind, sid, tag = parse_key(a["name"])
        if a["uptime_ms"] <= 0 or kind not in ("spell", "item"):
            continue
        if sid == 11567 and tag == "1":
            continue   # the Heroic Strike queue: a pressed button, drawn on its timeline lane
        if sid == 11585 and not ROTATION_CASTS_OVERPOWER:
            continue   # the Overpower window: never a buff; a timeline lane only when the rotation casts it
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


def talent_split(string):
    return "/".join(str(sum(int(c) for c in part)) for part in string.split("-"))


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
    pre = [(cast_of(p), p["doAtValue"]["const"]["val"]) for p in APL["rotation"]["prepullActions"]]
    lines = [(cast_of(p.get("action")), p.get("notes") or "") for p in APL["rotation"]["priorityList"]]
    return pre, lines


def histogram(buckets=40):
    """The fixture's 1 DPS bins (ILLUSTRATIVE until lane E), grouped client-side as spec 3.6 says."""
    h = FIX["dps_histogram"]
    counts, start = h["counts"], h["start"]
    width = -(-len(counts) // buckets)
    out = [sum(counts[i:i + width]) for i in range(0, len(counts), width)]
    return out, start, start + len(counts) * h["bin_dps"], width
