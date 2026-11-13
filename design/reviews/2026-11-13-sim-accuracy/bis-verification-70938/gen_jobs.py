"""gen_jobs.py <repo-root> <build> <outdir>: the spot-sim job files.

prevband.txt: every bare band 30 to 60, both factions, worn with the band below's
published set (slots it lacks emptied); must score under the band's own figure.
emptied.txt: the published set with one trinket, ranged, relic or off-hand pick
emptied, where the ranker decided the pick by sim or it had no competing
alternative; must score under its own set.
published.txt: every entry those touch, at the same settings, as the baseline.
Name: <spec>-<role>-<preset><band>-<faction>-<kind>[-<slot>], role tank/heal/dps.
"""
import json
import sys
from pathlib import Path

TANKS = {"warrior-protection", "paladin-protection", "druid-feral-bear"}
HEALERS = {"priest-discipline", "priest-holy", "paladin-holy", "shaman-restoration", "druid-restoration"}
SAMPLED = ("trinket1", "trinket2", "ranged", "off_hand")
COMMON = "-iterations 3000 -seed 7"


def role(spec):
    return "tank" if spec in TANKS else "heal" if spec in HEALERS else "dps"


def base(spec, b):
    return f"-spec {spec} -preset {b['preset']} -band {b['band']} -faction {b['faction']} {COMMON}"


def name(spec, b, kind):
    return f"{spec}-{role(spec)}-{b['preset']}{b['band']}-{b['faction']}-{kind}"


def main():
    root, build, out = sys.argv[1:4]
    bdir = Path(root) / "data" / "builds" / build / "bis"
    prev, emptied, published = [], [], []
    for f in sorted(bdir.glob("*.json")):
        d = json.loads(f.read_text())
        spec = d["spec"]
        idx = {(b["preset"], b["band"], b["faction"]): b for b in d["bands"]}
        for b in d["bands"]:
            published.append(f"{name(spec, b, 'published')} {base(spec, b)}")
            if b["preset"] == "bare" and b["band"] >= 30:
                lower = idx.get(("bare", b["band"] - 10, b["faction"]))
                if lower:
                    lp = {s["slot"]: s.get("item_id") or 0 for s in lower["slots"]}
                    gear = ",".join(f"{s['slot']}:{lp.get(s['slot'], 0)}" for s in b["slots"])
                    prev.append(f"{name(spec, b, 'prevband')} {base(spec, b)} -gear {gear}")
            for s in b["slots"]:
                if s["slot"] not in SAMPLED or not s.get("item_id"):
                    continue
                if s.get("score") is None or not s.get("alternatives"):
                    emptied.append(f"{name(spec, b, 'emptied')}-{s['slot']} {base(spec, b)} -gear {s['slot']}:0")
    Path(out).mkdir(parents=True, exist_ok=True)
    for fname, lines in (("prevband.txt", prev), ("emptied.txt", emptied), ("published.txt", published)):
        Path(out, fname).write_text("\n".join(lines) + "\n")
        print(fname, len(lines))


if __name__ == "__main__":
    main()
