"""What the logged raid healers cast: per class, each heal's casts a minute (by rank), its share of
the class's casts and effective healing, who it lands on (self, the main tank, another tank, anyone
else), and the healer's active time (Warcraft Logs' casting time over the kill).

Reads the healer-mana sample (logs-kills.json) through wcl.py's cache only: every request below is
one collect.py already made, so no key is needed. Usage:
WCL_CACHE=<cache dir> python3 -I castmix.py <healer-mana dir> <spellranks.json>
"""
import collections
import json
import os
import sys

HERE = sys.argv[1]
sys.path.insert(0, os.path.join(HERE, "logs"))
import wcl  # noqa: E402

RANKS = json.load(open(sys.argv[2]))["classes"]
CLASS_KEY = {"Priest": "priest", "Paladin": "paladin", "Shaman": "shaman", "Druid": "druid"}
# The healing table names a heal by its effect's spell, which for these differs from the cast's.
HEAL_EFFECT = {19993: "Flash of Light", 19968: "Holy Light", 25903: "Holy Shock"}
MIN_SHARE = 0.005


def spell_name(cls, guid):
    for name, rows in RANKS[CLASS_KEY[cls]].items():
        for r in rows:
            if r["id"] == guid:
                return name, r["rank"]
    return None, None


def all_events(view, code, start, end):
    out, cursor = [], start
    while cursor is not None:
        body = wcl.get(f"report/events/{view}/{code}", start=cursor, end=end)
        out.extend(body.get("events", []))
        cursor = body.get("nextPageTimestamp")
    return out


def main():
    kills = json.load(open(os.path.join(HERE, "logs-kills.json")))
    casts = collections.defaultdict(collections.Counter)      # class -> (spell, rank) -> casts
    targets = collections.defaultdict(collections.Counter)    # class -> spell -> target kind
    healed = collections.defaultdict(collections.Counter)     # class -> spell -> effective healing
    minutes = collections.Counter()
    active = collections.defaultdict(list)
    for k in kills:
        code, fid = k["report"], k["fight"]
        fight = next(f for f in wcl.get(f"report/fights/{code}", translate="true")["fights"] if f["id"] == fid)
        s, e = fight["start_time"], fight["end_time"]
        roles = wcl.get(f"report/tables/summary/{code}", start=s, end=e)["playerDetails"]
        taken = {x["id"]: x["total"] for x in wcl.get(f"report/tables/damage-taken/{code}", start=s, end=e)["entries"]}
        heal = {x["id"]: x for x in wcl.get(f"report/tables/healing/{code}", start=s, end=e)["entries"]}
        tanks = [p["id"] for p in roles.get("tanks", [])]
        mt = max(tanks, key=lambda i: taken.get(i, 0)) if tanks else None
        healers = {p["id"]: p["type"] for p in roles.get("healers", []) if p["type"] in CLASS_KEY}
        events = all_events("casts", code, s, e)
        for hid, cls in healers.items():
            minutes[cls] += (e - s) / 60000
            h = heal.get(hid, {})
            if h.get("activeTime"):
                active[cls].append(h["activeTime"] / (e - s))
            for a in h.get("abilities", []):
                name = HEAL_EFFECT.get(a["guid"]) or spell_name(cls, a["guid"])[0]
                healed[cls][name or a["name"]] += a["total"]
            for ev in events:
                if ev.get("type") != "cast" or ev.get("sourceID") != hid:
                    continue
                name, rank = spell_name(cls, ev["ability"]["guid"])
                if name is None:
                    continue
                casts[cls][(name, rank)] += 1
                t = ev.get("targetID")
                kind = "self" if t == hid else "main tank" if t == mt else "tank" if t in tanks else "other"
                targets[cls][(name, kind)] += 1
    for cls in sorted(casts):
        mins = minutes[cls]
        tot = sum(casts[cls].values())
        htot = sum(healed[cls].values()) or 1
        act = sorted(active[cls])
        print(f"## {cls}: {mins:.0f} healer-minutes, {tot / mins:.1f} casts a minute, "
              f"active time median {act[len(act) // 2]:.2f}")
        by_spell = collections.Counter()
        for (n, r), c in casts[cls].items():
            by_spell[n] += c
        print("| spell | casts a min | share of casts | share of effective healing | ranks (share of the spell's casts) | on self / main tank / other tank / other |")
        print("|---|---|---|---|---|---|")
        for n, c in by_spell.most_common():
            if c / tot < MIN_SHARE:
                continue
            ranks = sorted(((r, cc) for (nn, r), cc in casts[cls].items() if nn == n), key=lambda x: -x[1])
            rk = ", ".join(f"r{r} {cc / c:.0%}" for r, cc in ranks if cc / c >= 0.03)
            tg = " / ".join(f"{targets[cls][(n, kind)] / c:.0%}" for kind in ("self", "main tank", "tank", "other"))
            print(f"| {n} | {c / mins:.2f} | {c / tot:.0%} | {healed[cls][n] / htot:.0%} | {rk} | {tg} |")
        print()
        print("| healing source (the healing table) | share of effective healing |")
        print("|---|---|")
        for n, v in healed[cls].most_common():
            if v / htot >= MIN_SHARE:
                print(f"| {n} | {v / htot:.1%} |")
        print()


if __name__ == "__main__":
    main()
