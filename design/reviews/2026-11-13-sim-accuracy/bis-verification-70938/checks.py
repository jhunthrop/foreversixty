"""Offline checks over the active build's published bis files.

checks.py <repo-root> <build>: slot legality, weapon picks, the ranker's guards
(set_dps across bands, raid over bare, kept_alternative), tier order of the
alternatives in the json and md tables, MR94's metrics.health and MR100's
metrics.raid_fight. Every fact is re-derived from the build's own data.
Prints one line per failure and a summary per check.
"""
import json
import re
import sys
from collections import Counter, defaultdict
from pathlib import Path

ARMOR_GATES = {("hunter", 3): 40, ("shaman", 3): 40, ("warrior", 4): 40, ("paladin", 4): 40}
WEAPONS = {
    "warrior": {0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 13, 15, 16, 18},
    "paladin": {0, 1, 4, 5, 6, 7, 8},
    "hunter": {0, 1, 2, 3, 6, 7, 8, 10, 13, 15, 16, 18},
    "rogue": {0, 2, 3, 4, 7, 13, 15, 16, 18},
    "priest": {4, 10, 15, 19},
    "shaman": {0, 1, 4, 5, 10, 13, 15},
    "mage": {7, 10, 15, 19},
    "warlock": {7, 10, 15, 19},
    "druid": {4, 5, 10, 13, 15},
}
RANGED_SUB = {2, 3, 16, 18, 19}
DUAL_WIELD = {"warrior", "rogue", "hunter"}
HEALERS = {"priest-discipline", "priest-holy", "paladin-holy", "shaman-restoration", "druid-restoration"}
TANKS = {"warrior-protection", "paladin-protection", "druid-feral-bear"}
SLOT_OF = {"finger1": "finger", "finger2": "finger", "trinket1": "trinket", "trinket2": "trinket"}


def load(root, build):
    bdir = Path(root) / "data" / "builds" / build
    items = {i["id"]: i for i in json.loads((bdir / "items.json").read_text())}
    cls = {}
    for f in (bdir / "items").glob("*.json"):
        d = json.loads(f.read_text())
        cls[d["class_slug"]] = {i["id"]: i for i in d["items"]}
    specs = {f.stem: json.loads(f.read_text()) for f in sorted((bdir / "bis").glob("*.json"))}
    mds = {f.stem: f.read_text() for f in sorted((bdir / "bis").glob("*.md"))}
    return items, cls, specs, mds


def slot_fits(slot, entry):
    want = SLOT_OF.get(slot, slot)
    have = entry["slot"]
    if want == have:
        return True
    # a one-hand weapon in the off hand
    return slot == "off_hand" and have == "main_hand" and not entry.get("two_hand")


def legal(spec, klass, band, faction, slot, iid, items, cls):
    errs = []
    ce = cls[klass].get(iid)
    it = items.get(iid)
    if ce is None or it is None:
        return [f"{iid} not in {klass}'s items"]
    if not slot_fits(slot, ce):
        errs.append(f"{iid} slot {ce['slot']} in {slot}")
    if it["class_id"] == 4:
        gate = ARMOR_GATES.get((klass, it["subclass_id"]))
        if gate and band < gate:
            errs.append(f"{iid} armor subclass {it['subclass_id']} at band {band}")
    if it["class_id"] == 2 and it["subclass_id"] not in WEAPONS[klass]:
        errs.append(f"{iid} weapon subclass {it['subclass_id']}")
    fr = (it.get("faction_restriction") or "").lower().removesuffix("_only")
    if fr and fr != faction:
        errs.append(f"{iid} {fr}-only on {faction}")
    if (it.get("required_level") or 0) > band:
        errs.append(f"{iid} requires {it['required_level']} > {band}")
    return errs


def check_band(spec, b, items, cls, out):
    klass = spec.split("-")[0]
    band, faction = b["band"], b["faction"]
    tag = f"{spec} {b['preset']} {band} {faction}"
    picks = {s["slot"]: s.get("item_id") or 0 for s in b["slots"]}
    n = 0
    for s in b["slots"]:
        ids = ([s["item_id"]] if s.get("item_id") else []) + [a["item_id"] for a in s.get("alternatives") or []]
        for iid in ids:
            n += 1
            for e in legal(spec, klass, band, faction, s["slot"], iid, items, cls):
                out["legality"].append(f"{tag} {s['slot']}: {e}")
        alts = s.get("alternatives") or []
        deltas = [a.get("dps_delta") for a in alts if a.get("dps_delta") is not None]
        if any(d2 > d1 + 1e-9 for d1, d2 in zip(deltas, deltas[1:])):
            out["tier"].append(f"{tag} {s['slot']}: alternatives not descending {deltas}")
        if any(d > 1e-6 for d in deltas):
            out["tier"].append(f"{tag} {s['slot']}: an alternative above its pick {max(deltas)}")
    for a, c in (("finger1", "finger2"), ("trinket1", "trinket2")):
        if picks.get(a) and picks.get(a) == picks.get(c):
            out["legality"].append(f"{tag}: {a} and {c} both {picks[a]}")
    mh, oh, rg = picks.get("main_hand"), picks.get("off_hand"), picks.get("ranged")
    if not mh:
        out["weapons"].append(f"{tag}: no main hand")
    elif cls[klass].get(mh, {}).get("two_hand") and oh:
        out["weapons"].append(f"{tag}: two-hander {mh} beside off hand {oh}")
    if oh:
        oi = items.get(oh, {})
        if oi.get("class_id") == 2:
            if klass not in DUAL_WIELD:
                out["weapons"].append(f"{tag}: weapon {oh} in the off hand")
    if mh and items.get(mh, {}).get("subclass_id") in RANGED_SUB and items[mh]["class_id"] == 2:
        out["weapons"].append(f"{tag}: ranged weapon {mh} in the main hand")
    if rg:
        ri = items.get(rg, {})
        if ri.get("class_id") == 2 and ri.get("subclass_id") not in RANGED_SUB:
            out["weapons"].append(f"{tag}: ranged slot holds {rg} subclass {ri.get('subclass_id')}")
    if spec == "hunter-survival":
        names = [w["stat"] for w in b.get("weights") or []]
        if "ranged_attack_power" in names and "attack_power" not in names:
            out["weapons"].append(f"{tag}: Survival weights are the ranged table")
    if spec in TANKS:
        m = b.get("metrics") or {}
        if not m.get("health"):
            out["mr94"].append(f"{tag}: no metrics.health")
        elif m.get("effective_health") and m["health"] > m["effective_health"] + 1e-6:
            out["mr94"].append(f"{tag}: health {m['health']:.0f} above effective health {m['effective_health']:.0f}")
    if spec in HEALERS and band == 60:
        rf = (b.get("metrics") or {}).get("raid_fight")
        if not rf:
            out["mr100"].append(f"{tag}: no metrics.raid_fight")
    if b.get("kept_alternative"):
        out["kept"].append((tag, b["kept_alternative"]))
    return n


def check_bands(spec, d, out):
    series = defaultdict(dict)
    for b in d["bands"]:
        series[(b["faction"], b["preset"])][b["band"]] = b["set_dps"]
    for (faction, preset), s in series.items():
        if preset != "bare":
            continue
        bands = sorted(s)
        for lo, hi in zip(bands, bands[1:]):
            if s[hi] <= s[lo]:
                out["bands"].append(f"{spec} {faction}: bare {hi} {s[hi]:.2f} <= bare {lo} {s[lo]:.2f}")
        raid = series.get((faction, "raid"), {}).get(60)
        if raid is not None and 60 in s and raid <= s[60]:
            out["bands"].append(f"{spec} {faction}: raid 60 {raid:.2f} <= bare 60 {s[60]:.2f}")


ROW = re.compile(r"^\| (\w+) \| .*\| ([^|]*) \|$")
DELTA = re.compile(r"\(\d+, ([+-]\d+\.\d+) [A-Z]+")


def check_md(spec, text, out):
    rows = 0
    for line in text.splitlines():
        m = ROW.match(line)
        if not m or m.group(1) in ("Slot",):
            continue
        rows += 1
        deltas = [float(x) for x in DELTA.findall(m.group(2))]
        if any(b > a + 1e-9 for a, b in zip(deltas, deltas[1:])):
            out["tier"].append(f"{spec}.md {m.group(1)}: {deltas}")
    return rows


def main():
    root, build = sys.argv[1], sys.argv[2]
    items, cls, specs, mds = load(root, build)
    out = defaultdict(list)
    picks = bands = mdrows = 0
    for spec, d in specs.items():
        for b in d["bands"]:
            bands += 1
            picks += check_band(spec, b, items, cls, out)
        check_bands(spec, d, out)
        mdrows += check_md(spec, mds.get(spec, ""), out)
    healers60 = Counter()
    for spec in HEALERS:
        for b in specs[spec]["bands"]:
            if b["band"] == 60 and (b.get("metrics") or {}).get("raid_fight"):
                healers60[(b["faction"], b["preset"])] += 1
    print(f"specs {len(specs)} bands {bands} picks+alternatives {picks} md rows {mdrows}")
    for k in ("legality", "weapons", "bands", "tier", "mr94", "mr100"):
        print(f"== {k}: {len(out[k])} failures")
        for line in out[k][:60]:
            print("  " + line)
    print(f"== raid_fight blocks at band 60 by (faction, preset): {dict(healers60)}")
    print(f"== kept_alternative: {len(out['kept'])}")
    for tag, k in out["kept"]:
        print(f"  {tag}: kept {k.get('kept')} {k.get('kept_score', 0):.2f}+-{k.get('kept_error', 0):.2f} over new "
              f"{k.get('new_score', 0):.2f}+-{k.get('new_error', 0):.2f} slots {','.join(k.get('differing_slots') or [])}")


if __name__ == "__main__":
    main()
