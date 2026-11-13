#!/usr/bin/env bash
# bands.sh: DPS, out-of-mana time and the conjured items used for every published entry of SPECS
# (default the three mages) at raid 60 and bare 20 to 60, both factions; ITER iterations (default 5000),
# seed SEED (default 1), one sim at a time, niced to 15, GOMAXPROCS=5. Run on the box from sim/:
#   tools/remote-test.sh --pin-engine [--engine-sha <sha>] --dir sim HEAD -- env SPECS="..." bash ../design/reviews/2026-11-13-sim-accuracy/arcane-shadow-70938/bands.sh
set -euo pipefail
echo "engine $(grep 'replace github.com/wowsims/classic' go.mod | tail -1)"
make -C .. simdb >/dev/null
bin=$(mktemp -d)
go build -o "$bin/spec-breakdown" ./cmd/spec-breakdown
for spec in ${SPECS:-mage-arcane mage-fire mage-frost}; do
	for f in alliance horde; do
		for pb in ${BANDS:-raid:60 bare:60 bare:50 bare:40 bare:30 bare:20}; do
			p=${pb%:*}; b=${pb#*:}
			out=$(nice --adjustment=15 env GOMAXPROCS=5 "$bin/spec-breakdown" -iterations "${ITER:-5000}" -seed "${SEED:-1}" -spec $spec -preset $p -band $b -faction $f "$@" 2>/dev/null) || { echo "$spec $p $b $f FAILED"; continue; }
			dps=$(echo "$out" | grep -m1 -o 'DPS [0-9.]* +/- [0-9.]*' | awk '{print $2, $4}')
			oom=$(echo "$out" | grep -m1 -o 'out of mana: [0-9.]*' | awk '{print $4}')
			items=$(echo "$out" | awk -F'|' '/^\| item (5514|5513|8007|8008|12662) \|/ && $5+0>0 {gsub(/ /,"",$2); printf "%s=%s ", $2, $5+0}')
			echo "$spec $p $b $f dps $dps oom $oom $items"
		done
	done
done
