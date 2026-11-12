#!/usr/bin/env bash
# rows2.sh: Fire's Fireball rank 4 against rank 5 at 25 and 27, every mage
# race, on the conjure-level gem engine (rows.sh found the default Alliance
# race +0.2% on rank 4 at both levels and the Horde one on rank 5 at 25).
# Seeds 1 and 50001, 10,000 iterations, one sim at a time, niced,
# GOMAXPROCS=5. Run from the site's sim directory
# (tools/remote-test.sh --dir sim --pin-engine).
set -euo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
run() { # level race rank seed
	nice --adjustment=15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec mage-fire -preset bare -band 20 -level "$1" \
		-race "$2" -rank "10151=$3" -seed "$4" -iterations 10000 -json 2>/dev/null |
		python3 -c 'import json,sys; d=json.load(sys.stdin); print(round(d["DPS"],3), round(d["Error"],3), round(d.get("SecondsOOM",0),1), end="")' || echo "FAILED"
	echo "	mage-fire level $1 10151 rank $3 $2 seed $4"
}
for level in 25 27; do
	for race in human gnome high-order-skyborne undead orc troll; do
		for seed in 1 50001; do
			run "$level" "$race" 4 "$seed"
			run "$level" "$race" 5 "$seed"
		done
	done
done
