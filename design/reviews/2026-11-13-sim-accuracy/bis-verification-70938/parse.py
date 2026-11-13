"""parse.py <spot.log> <bis-dir>: the spot sims against their baselines.

Each check run (prevband, emptied-<slot>) is compared with the same entry's
'published' run from the same log: DPS and healers on the figure and its
error; tanks on the tank score's ratio (health / damage taken a second x
(threat / 100)^0.25 x exp(-0.005 TMI); the boss's raw damage cancels).
Also prints each published run against the file's set_dps (calibration).
Verdicts: below (beyond the combined error), tie (inside it), ABOVE (beyond it).
"""
import json
import math
import re
import sys
from collections import Counter
from pathlib import Path

NUM = r"([0-9.]+)"


def parse_block(name, text):
    if "-tank-" in name:
        g = lambda pat: float(re.search(pat, text).group(1))
        try:
            h, dt, tmi = g(r"\| health \| " + NUM), g(r"\| damage taken /s \| " + NUM), g(r"\| TMI \| " + NUM)
            tps, tpe = float(re.search(r"\| threat /s \| " + NUM + r" \+/- " + NUM, text).group(1)), 0
            death = g(r"\| chance of death \| " + NUM)
        except AttributeError:
            return None
        score = h / max(dt, 1e-9) * (max(tps, 1) / 100) ** 0.25 * math.exp(-0.005 * tmi)
        return {"v": score, "e": 0.0, "health": h, "dtps": dt, "tmi": tmi, "tps": tps, "death": death}
    if "-heal-" in name:
        m = re.search(r"\| effective HPS \| " + NUM + r" \+/- " + NUM, text)
        return {"v": float(m.group(1)), "e": float(m.group(2))} if m else None
    m = re.search(r"^\{.*\}$", text, re.M)
    if not m:
        return None
    r = json.loads(m.group(0))
    return {"v": r["DPS"], "e": r["Error"]}


def main():
    log, bis = Path(sys.argv[1]).read_text(), Path(sys.argv[2])
    blocks = re.split(r"^##### (\S+)\n", log, flags=re.M)
    res = {blocks[i]: parse_block(blocks[i], blocks[i + 1]) for i in range(1, len(blocks) - 1, 2) if blocks[i] != "provenance"}
    files = {f.stem: json.loads(f.read_text()) for f in bis.glob("*.json")}
    calib = []
    for n, r in res.items():
        if n.endswith("-published") and r and "-tank-" not in n:
            m = re.match(r"(.+)-(dps|heal)-(bare|raid)(\d+)-(alliance|horde)-published", n)
            spec, _, preset, band, faction = m.groups()
            b = next(x for x in files[spec]["bands"] if x["preset"] == preset and x["band"] == int(band) and x["faction"] == faction)
            calib.append((n, r["v"], b["set_dps"]))
    off = [c for c in calib if abs(c[1] - c[2]) > 0.006 * max(1, c[2]) / 100 + 0.006]
    print(f"calibration: {len(calib)} published runs, {len(off)} differ from set_dps by more than rounding")
    for c in off[:20]:
        print(f"  {c[0]}: sim {c[1]:.2f} file {c[2]:.2f}")
    verdicts = Counter()
    for n, r in sorted(res.items()):
        if n.endswith("-published"):
            continue
        key = re.sub(r"-(prevband|emptied-\w+)$", "-published", n)
        kind = "prevband" if n.endswith("-prevband") else "emptied"
        b = res.get(key)
        if r is None or b is None:
            verdicts[(kind, "no figure")] += 1
            print(f"  NOFIG {n}")
            continue
        err = math.hypot(r["e"], b["e"])
        d = r["v"] - b["v"]
        v = "below" if d < -err else ("ABOVE" if d > err else "tie")
        if "-tank-" in n:
            v = "below" if d < -1e-9 * b["v"] else ("ABOVE" if d > 0.005 * b["v"] else "tie")
        verdicts[(kind, v)] += 1
        if v != "below":
            extra = ""
            if "-tank-" in n:
                extra = f" health {r['health']:.0f}/{b['health']:.0f} dtps {r['dtps']:.1f}/{b['dtps']:.1f} tmi {r['tmi']:.1f}/{b['tmi']:.1f} tps {r['tps']:.1f}/{b['tps']:.1f} death {r['death']:.1f}/{b['death']:.1f}"
            print(f"  {v} {n}: {r['v']:.3f} vs {b['v']:.3f} (+-{err:.3f}, {100 * d / b['v']:+.2f}%){extra}")
    for k, c in sorted(verdicts.items()):
        print(f"{k[0]} {k[1]}: {c}")


if __name__ == "__main__":
    main()
