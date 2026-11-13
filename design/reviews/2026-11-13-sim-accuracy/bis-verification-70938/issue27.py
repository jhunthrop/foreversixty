"""issue27.py <repo-root> <old-build> <new-build>: issue #27's items (stats,
armor, damage or speed differ between the two builds' class item files) and
every published pick or alternative of the new build that wears one."""
import json
import sys
from pathlib import Path

FIELDS = ("stats", "armor", "damage_min", "damage_max", "speed", "percent_stats", "effect_text")


def class_items(bdir):
    out = {}
    for f in (bdir / "items").glob("*.json"):
        for i in json.loads(f.read_text())["items"]:
            out[i["id"]] = i
    return out


def main():
    root, old_b, new_b = Path(sys.argv[1]), sys.argv[2], sys.argv[3]
    old, new = class_items(root / "data/builds" / old_b), class_items(root / "data/builds" / new_b)
    moved = {}
    for iid, n in new.items():
        o = old.get(iid)
        if o is None:
            continue
        diff = [f for f in FIELDS if o.get(f) != n.get(f)]
        if diff:
            moved[iid] = (n["name"], {f: (o.get(f), n.get(f)) for f in diff})
    print(f"{len(moved)} items differ between {old_b} and {new_b}")
    for iid, (name, d) in sorted(moved.items()):
        print(f"  {iid} {name}: " + "; ".join(f"{k} {v[0]} -> {v[1]}" for k, v in d.items())[:300])
    print("== picks wearing one")
    for f in sorted((root / "data/builds" / new_b / "bis").glob("*.json")):
        for b in json.loads(f.read_text())["bands"]:
            for s in b["slots"]:
                if s.get("item_id") in moved:
                    print(f"  PICK {f.stem} {b['preset']} {b['band']} {b['faction']} {s['slot']}: {s['item_id']} {moved[s['item_id']][0]}")


if __name__ == "__main__":
    main()
