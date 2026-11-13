"""Summarise collect.py's kills: fight length, each healer class's overheal and mana gained by
source (potions and runes counted as the gain over the item's mean), and the damage the healers
face. The mana bar itself is not in these logs (classResources reads 0 of 0 on every kill). Usage: python3 -I analyze.py <kills.json>..."""
import collections
import json
import sys

SOURCES = {"17531": "Major Mana Potion", "17530": "Superior Mana Potion", "11903": "Greater Mana Potion",
           "16666": "Demonic Rune", "27869": "Dark Rune",
           "17360": "Mana Tide", "16191": "Mana Tide", "17355": "Mana Tide", "17359": "Mana Tide",
           "10494": "Mana Spring", "20272": "Illumination",
           "25894": "Blessing of Wisdom", "19854": "Blessing of Wisdom", "20351": "Seal of Wisdom",
           "20353": "Judgement of Wisdom", "18194": "Mana Regeneration (food)", "15604": "Second Wind"}
POTIONS = ("Major Mana Potion", "Superior Mana Potion", "Greater Mana Potion")
RUNES = ("Demonic Rune", "Dark Rune")
MAJOR_POTION_MEAN = 1800
RUNE_MEAN = 1200


def q(values, p):
    values = sorted(values)
    if not values:
        return float("nan")
    k = (len(values) - 1) * p
    lo = int(k)
    hi = min(lo + 1, len(values) - 1)
    return values[lo] + (values[hi] - values[lo]) * (k - lo)


def main():
    kills = [k for path in sys.argv[1:] for k in json.load(open(path))]
    secs = [k["seconds"] for k in kills]
    print(f"kills {len(kills)}; length s: median {q(secs, .5):.0f}, p75 {q(secs, .75):.0f}, p90 {q(secs, .9):.0f}, max {max(secs):.0f}")
    by_boss = collections.defaultdict(list)
    for k in kills:
        by_boss[k["boss"]].append(k["seconds"])
    print("by boss:", ", ".join(f"{b} {q(v, .5):.0f}" for b, v in by_boss.items()))
    mt = [k["mt_dtps"] for k in kills if "mt_dtps" in k]
    sp = [k["mt_spike5"] for k in kills if "mt_spike5" in k]
    mem = [k["member_dtps"] for k in kills]
    nh = [len(k["healers"]) for k in kills]
    print(f"main tank dtps median {q(mt, .5):.0f} (p25 {q(mt, .25):.0f}, p75 {q(mt, .75):.0f}); "
          f"largest 5 s over mean 5 s: median {q(sp, .5):.2f} (p75 {q(sp, .75):.2f})")
    print(f"other players dtps median {q(mem, .5):.0f} (p75 {q(mem, .75):.0f}); healers a kill median {q(nh, .5):.0f}")
    load = [((k.get("mt_dtps", 0)) + k["member_dtps"] * (k["players"] - k["tanks"])) / max(1, len(k["healers"])) for k in kills]
    print(f"damage a healer faces (main tank + every other player, over the healers) median {q(load, .5):.0f}")
    rows = collections.defaultdict(list)
    for k in kills:
        for h in k["healers"]:
            rows[f'{h["spec"]} {h["class"]}'].append((k, h))
    print("\n| healer | healer kills | minutes | overheal % median (p25 to p75) | Major Mana Potions a min | runes a min | logged mana a min by source |")
    print("|---|---|---|---|---|---|---|")
    for name, items in sorted(rows.items(), key=lambda kv: -len(kv[1])):
        oh = [100 * h["overheal"] / (h["healing"] + h["overheal"]) for _, h in items if h["healing"] + h["overheal"] > 0]
        minutes = sum(k["seconds"] for k, _ in items) / 60
        src = collections.Counter()
        for _, h in items:
            for guid, (nm, gain) in h["gains"].items():
                src[SOURCES.get(guid, nm)] += gain
        top = "; ".join(f"{n} {v / minutes:.0f}" for n, v in src.most_common(5))
        print(f"| {name} | {len(items)} | {minutes:.0f} | {q(oh, .5):.0f} ({q(oh, .25):.0f} to {q(oh, .75):.0f}) | "
              f"{src['Major Mana Potion'] / MAJOR_POTION_MEAN / minutes:.2f} | "
              f"{sum(src[r] for r in RUNES) / RUNE_MEAN / minutes:.2f} | {top} |")


if __name__ == "__main__":
    main()
