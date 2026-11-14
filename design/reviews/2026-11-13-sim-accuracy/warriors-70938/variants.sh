#!/usr/bin/env bash
# variants.sh: the rotation search's two raid winners at 10,000 iterations, seeds 1 and 50001, both factions,
# against the curated cards. Run from sim/ with --engine-sha.
set -euo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown; go build -o "$bin" ./cmd/spec-breakdown
R=../design/reviews/2026-11-13-sim-accuracy/warriors-70938
run() { nice --adjustment=15 env GOMAXPROCS=5 "$bin" -repo-root .. -iterations 10000 -json "$@" 2>/dev/null | python3 -c 'import json,sys; r=json.load(sys.stdin); print("%.1f %.1f" % (r["DPS"], r["Error"]))'; }
for seed in 1 50001; do for f in alliance horde; do
	echo "fury raid 60 $f seed $seed card $(run -spec warrior-fury -preset raid -band 60 -faction $f -seed $seed) hamstring $(run -spec warrior-fury -preset raid -band 60 -faction $f -seed $seed -rotation $R/rot-fury-hamstring.json)"
	echo "arms raid 60 $f seed $seed card $(run -spec warrior-arms -preset raid -band 60 -faction $f -seed $seed) slam-first $(run -spec warrior-arms -preset raid -band 60 -faction $f -seed $seed -rotation $R/rot-arms-slam-first.json)"
done; done
for b in 60 50; do echo "fury bare $b alliance card $(run -spec warrior-fury -preset bare -band $b -seed 1) hamstring $(run -spec warrior-fury -preset bare -band $b -seed 1 -rotation $R/rot-fury-hamstring.json)"; done
