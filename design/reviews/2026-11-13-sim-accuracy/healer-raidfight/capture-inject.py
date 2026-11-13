"""Capture only: give the active build's band 60 healer entries a raid_fight block from the lane's box run."""
import json, pathlib, sys
root = pathlib.Path(sys.argv[1])
build = json.loads((root / "web/src/data/active-build.json").read_text())["build"]
measured = {  # alliance raid 60, seed 1, 2000 iterations (healer-raidfight/run1)
    "druid-restoration": (111.94, 0.22, 433.4, 8350, 8572),
    "paladin-holy": (109.62, 0.15, 112.3, 9776, 9820),
    "priest-discipline": (119.24, 0.31, 216.5, 9353, 9775),
    "priest-holy": (125.62, 0.23, 363.6, 9142, 9576),
    "shaman-restoration": (131.59, 0.11, 181.7, 9782, 9810),
}
for spec, (hps, err, raw, left, pool) in measured.items():
    path = root / f"data/builds/{build}/bis/{spec}.json"
    if not path.exists():
        print("missing", path); continue
    data = json.loads(path.read_text())
    for band in data["bands"]:
        if band["band"] == 60 and band.get("metrics"):
            band["metrics"]["raid_fight"] = {"profile": "raid60-logs", "duration_sec": 110, "hps": hps, "hps_error": err,
                "raw_hps": raw, "overheal_pct": 1 - hps / raw, "mana_left": left, "mana_left_pct": left / pool}
    path.write_text(json.dumps(data))
    print("injected", spec)
