#!/usr/bin/env bash
# rows.sh: the Arcane lane's rank-row re-measure on the engine with the mana
# gem at its conjure level (28). Every mage rank row that sits at 23 to 27,
# where the old engine gave the mage a Mana Agate it cannot conjure yet:
# Arcane Frostbolt rank 1 against 4 at 24 to 27, Frost Frostbolt rank 4
# against 5 at 26 and 27, Fire Fireball rank 4 against 5 at 24. Both
# factions' default races, seeds 1 and 50001, 10,000 iterations, one sim at
# a time, niced, GOMAXPROCS=5 (two other lanes share the box). Run from the
# site's sim directory (tools/remote-test.sh --dir sim --engine-sha <sha>).
set -euo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
run() { # spec level spell rank faction seed
	nice --adjustment=15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec "$1" -preset bare -band 20 -level "$2" \
		-rank "$3=$4" -faction "$5" -seed "$6" -iterations 10000 -json 2>/dev/null |
		python3 -c 'import json,sys; d=json.load(sys.stdin); print(round(d["DPS"],3), round(d["Error"],3), round(d.get("SecondsOOM",0),1), end="")' || echo "FAILED"
	echo "	$1 level $2 $3 rank $4 $5 seed $6"
}
for cell in "mage-arcane 24 10181 1 4" "mage-arcane 25 10181 1 4" "mage-arcane 26 10181 1 4" "mage-arcane 27 10181 1 4" \
	"mage-frost 26 10181 4 5" "mage-frost 27 10181 4 5" "mage-fire 24 10151 4 5"; do
	set -- $cell
	for faction in alliance horde; do
		for seed in 1 50001; do
			run "$1" "$2" "$3" "$4" "$faction" "$seed"
			run "$1" "$2" "$3" "$5" "$faction" "$seed"
		done
	done
done
