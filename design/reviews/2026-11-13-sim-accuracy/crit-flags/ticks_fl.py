"""Count critical over all damage and healing events per class and spell in
the foreverlogs.gg report caches (one <id>.json per report with an "events"
list, as the 2026-11-12 lanes fetched them).

Usage: python3 -I ticks_fl.py <cache dir> > ticks_fl.tsv
"""
import collections
import glob
import json
import os
import re
import sys


def main(cache):
    n = collections.Counter()
    c = collections.Counter()
    who = collections.defaultdict(set)
    names = {}
    seen = set()
    for f in sorted(glob.glob(f"{cache}/*.json")):
        if not re.fullmatch(r"\d+\.json", os.path.basename(f)):
            continue
        for e in json.load(open(f)).get("events", []):
            if e["id"] in seen or e.get("type") not in ("damage", "heal", "periodic_damage", "periodic_heal"):
                continue
            seen.add(e["id"])
            src = e.get("source") or {}
            if not src.get("isPlayer") or not e.get("spell"):
                continue
            ev = e["type"] + ("_tick" if e.get("isPeriodic") or e.get("periodic") else "")
            k = (src.get("class") or "?", str(e["spell"]["id"]), ev)
            n[k] += 1
            c[k] += bool(e.get("isCritical"))
            who[k].add(src.get("id"))
            names[k[1]] = e["spell"].get("name", "?")
    print("class\tspell\tname\tevent\tevents\tcrits\tplayers")
    for k in sorted(n):
        print(f"{k[0]}\t{k[1]}\t{names[k[1]]}\t{k[2]}\t{n[k]}\t{c[k]}\t{len(who[k])}")


if __name__ == "__main__":
    main(sys.argv[1])
