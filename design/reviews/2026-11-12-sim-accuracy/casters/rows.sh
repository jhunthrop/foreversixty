#!/usr/bin/env bash
# rows.sh: the Fire and Frost rank rows at 23 to 27 re-judged on the engine
# that gives a mage its first mana gem at 28 (the Arcane lane's G1, ledger
# MR67). The Arcane lane re-judged the rows that sit there (Fire Fireball
# rank 4 at 24, Frost Frostbolt rank 4 at 26 and 27); this sweep asks
# whether a lower rank now wins at the levels with no row: Fire Fireball
# rank 4 against 5 at 25 to 27, Fire Blast rank 2 against 3 at 23 to 27
# (both specs' cards cast it: Fire's for the hit, Frost's never), Frost
# Frostbolt rank 3 against 4 at 23 to 25. Both factions' default races,
# seeds 1 and 50001, 10,000 iterations, one sim at a time, niced,
# GOMAXPROCS=5. Run from the site's sim directory
# (tools/remote-test.sh --dir sim --pin-engine).
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
for cell in "mage-fire 25 10151 4 5" "mage-fire 26 10151 4 5" "mage-fire 27 10151 4 5" \
	"mage-fire 23 10199 2 3" "mage-fire 24 10199 2 3" "mage-fire 25 10199 2 3" "mage-fire 26 10199 2 3" "mage-fire 27 10199 2 3" \
	"mage-frost 23 10181 3 4" "mage-frost 24 10181 3 4" "mage-frost 25 10181 3 4"; do
	set -- $cell
	for faction in alliance horde; do
		for seed in 1 50001; do
			run "$1" "$2" "$3" "$4" "$faction" "$seed"
			run "$1" "$2" "$3" "$5" "$faction" "$seed"
		done
	done
done
