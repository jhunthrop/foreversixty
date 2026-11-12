#!/usr/bin/env bash
# bare60.sh <spec> [iterations]: the bare band-60 entry at both factions and both
# seeds (1 and 50001) on the guide build the ref carries, printing DPS, error and
# seconds out of mana. Run from sim/ on the box (tools/remote-test.sh --dir sim
# --pin-engine <ref>); the walk (walk.sh) covers seed 1 at every checkpoint level.
set -uo pipefail
spec=$1; iters=${2:-10000}
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
for f in alliance horde; do
  for seed in 1 50001; do
    out=$(nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -iterations "$iters" -seed "$seed" -spec "$spec" -preset bare -band 60 -level 60 -faction $f -json 2>/dev/null) || { echo "$spec $f seed $seed FAILED"; continue; }
    echo "$out" | python3 -c "import json,sys;d=json.load(sys.stdin);print('$spec','$f','seed $seed',round(d['DPS'],2),round(d['Error'],2),round(d['SecondsOOM'],1))"
  done
done
