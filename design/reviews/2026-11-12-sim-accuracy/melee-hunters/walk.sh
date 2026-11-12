#!/usr/bin/env bash
# walk.sh <spec> [iterations]: the guide's leveling walk at the checkpoint levels, both factions,
# bare, printing level, faction, DPS, error and seconds out of mana. Run from sim/ on the box
# at a ref whose guide frontmatter carries the build under test.
set -uo pipefail
spec=$1; iters=${2:-10000}
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
for f in horde alliance; do
  for l in 10 15 20 25 28 30 33 35 38 40 43 45 48 50 53 55 58 60; do
    if ((l < 30)); then b=20; else b=$((l / 10 * 10)); fi
    out=$(nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -iterations "$iters" -seed 1 -spec "$spec" -preset bare -band $b -level $l -faction $f -json 2>/dev/null) || { echo "$spec $f L$l FAILED $(echo "$out" | tail -2 | tr '\n' ' ')"; continue; }
    echo "$out" | python3 -c "import json,sys;d=json.load(sys.stdin);print('$spec','$f','L$l',round(d['DPS'],2),round(d['Error'],2),round(d['SecondsOOM'],1))"
  done
done
go run ./cmd/addon-builds -repo-root .. >/dev/null 2>&1 && python3 -c "import json;d=json.load(open('../data/builds/1.60.1.70790/addon-builds.json'));print('ORDER', '$spec', json.dumps(d['$spec']['leveling']['order']))"
