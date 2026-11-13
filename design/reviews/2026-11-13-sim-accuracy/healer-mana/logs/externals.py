"""How often a logged raid healer receives another player's mana or throughput cooldown: Innervate
(29166) and Power Infusion (10060), on the kills collect.py sampled, by the healer's class.
Usage: WCL_KEY=... WCL_CACHE=<dir> python3 -I externals.py <this dir> <kills.json>"""
import collections
import json
import sys

sys.path.insert(0, sys.argv[1])
import wcl  # noqa: E402

AURAS = {"Innervate": 29166, "Power Infusion": 10060}


def main():
    kills = json.load(open(sys.argv[2]))
    got = collections.Counter()
    seen = collections.Counter()
    minutes = collections.Counter()
    for k in kills:
        fights = wcl.get(f"report/fights/{k['report']}", translate="true")
        fight = next(f for f in fights["fights"] if f["id"] == k["fight"])
        s, e = fight["start_time"], fight["end_time"]
        roles = wcl.get(f"report/tables/summary/{k['report']}", start=s, end=e)["playerDetails"]
        healers = {p["id"]: p["type"] for p in roles.get("healers", [])}
        for cls in healers.values():
            seen[cls] += 1
            minutes[cls] += k["seconds"] / 60
        for label, aura in AURAS.items():
            body = wcl.get(f"report/tables/buffs/{k['report']}", start=s, end=e, abilityid=aura, by="target")
            for a in body.get("auras", []):
                if a["id"] in healers:
                    got[(healers[a["id"]], label)] += a.get("totalUses", 0)
    print("| healer class | healer kills | minutes | Innervates received | Power Infusions received |")
    print("|---|---|---|---|---|")
    for cls, n in seen.most_common():
        print(f"| {cls} | {n} | {minutes[cls]:.0f} | {got[(cls, 'Innervate')]} | {got[(cls, 'Power Infusion')]} |")


if __name__ == "__main__":
    main()
