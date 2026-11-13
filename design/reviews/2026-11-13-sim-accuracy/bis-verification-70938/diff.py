"""diff.py <old-dir> <new-dir>: per band entry, old against new: set_dps move,
talents changed (T), race changed (R), the re-picked slots, and for tanks and
healers the metrics that moved. One line per entry."""
import json
import sys
from pathlib import Path


def entries(d):
    return {(b["preset"], b["band"], b["faction"]): b for b in d["bands"]}


def main():
    old_dir, new_dir = Path(sys.argv[1]), Path(sys.argv[2])
    only = set(sys.argv[3:])
    for f in sorted(new_dir.glob("*.json")):
        if only and f.stem not in only:
            continue
        new = json.loads(f.read_text())
        old_path = old_dir / f.name
        old = json.loads(old_path.read_text()) if old_path.exists() else {"bands": []}
        oe = entries(old)
        print(f"## {new['spec']} engine {old.get('engine_version')} -> {new.get('engine_version')}")
        for key, b in entries(new).items():
            o = oe.get(key)
            tag = f"{key[0]} {key[1]} {key[2][0].upper()}"
            if not o:
                print(f"  {tag}: new entry {b['set_dps']:.2f}")
                continue
            pct = 100 * (b["set_dps"] / o["set_dps"] - 1) if o["set_dps"] else 0
            flags = ("T" if b["talents"] != o["talents"] else "") + ("R" if b["race"] != o["race"] else "")
            os_ = {s["slot"]: s.get("item_id") for s in o["slots"]}
            moved = [f"{s['slot']}:{os_.get(s['slot'])}>{s.get('item_id')}" for s in b["slots"] if os_.get(s["slot"]) != s.get("item_id")]
            line = f"  {tag}: {o['set_dps']:.2f} -> {b['set_dps']:.2f} ({pct:+.1f}%) {flags} {' '.join(moved)}"
            m = b.get("metrics") or {}
            if b.get("role") == "tank":
                om = o.get("metrics") or {}
                line += f" | tps {om.get('tps', 0):.0f}->{m.get('tps', 0):.0f} dtps {om.get('dtps', 0):.1f}->{m.get('dtps', 0):.1f} tmi {om.get('tmi', 0):.1f}->{m.get('tmi', 0):.1f} death {om.get('chance_of_death', 0):.3f}->{m.get('chance_of_death', 0):.3f} health {m.get('health')}"
            if m.get("raid_fight"):
                rf = m["raid_fight"]
                line += " | raid_fight " + json.dumps(rf, sort_keys=True)[:200]
            print(line)


if __name__ == "__main__":
    main()
