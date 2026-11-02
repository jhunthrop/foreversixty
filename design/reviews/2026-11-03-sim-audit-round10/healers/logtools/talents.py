"""Decode a spec's published 60 raid talent string: talents.py <spec> [band] [preset] [faction]."""
import json
import sys

B = "/Users/jh/code/forever/.claude/worktrees/r10-healers/data/builds/1.60.1.70707/"
spec = sys.argv[1]
band = int(sys.argv[2]) if len(sys.argv) > 2 else 60
preset = sys.argv[3] if len(sys.argv) > 3 else "raid"
fac = sys.argv[4] if len(sys.argv) > 4 else "horde"
cls = spec.split("-")[0]
bis = json.load(open(B + f"bis/{spec}.json"))
row = [b for b in bis["bands"] if b["band"] == band and b["preset"] == preset and b["faction"] == fac][0]
print(row["talents"], row.get("race"))
tal = json.load(open(B + f"talents/{cls}.json"))
for tree, part in zip(sorted(tal["trees"], key=lambda t: t["position"]), row["talents"].split("-")):
    ts = sorted(tree["talents"], key=lambda t: (t["tier"], t["column"]))
    for t, ch in zip(ts, part):
        if ch != "0":
            r = int(ch)
            sid = t["ranks"][r - 1]["spell_id"]
            print(f"{tree['name']:12} {t['name']:28} {r}/{t['max_rank']} spell={sid} :: {t['ranks'][r-1]['description'][:160]}")
