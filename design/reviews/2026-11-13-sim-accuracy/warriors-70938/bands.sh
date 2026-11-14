#!/usr/bin/env bash
# bands.sh: every melee spec's published raid 60 entry (both factions) plus the warriors' bare 60, DPS for
# damage specs and threat /s for tanks, under the engine this checkout is pinned to. ITER (default 10000)
# iterations, seed 1, one sim at a time, niced to 15, GOMAXPROCS=5. Run on the box from sim/:
#   tools/remote-test.sh [--engine-sha <sha>] --dir sim HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/warriors-70938/bands.sh
set -euo pipefail
echo "engine $(grep 'replace github.com/wowsims/classic' go.mod | tail -1)"
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
run() { nice --adjustment=15 env GOMAXPROCS=5 "$bin" -repo-root .. -seed 1 -iterations "${ITER:-10000}" "$@" 2>/dev/null; }
for spec in ${SPECS:-warrior-arms warrior-fury rogue-assassination rogue-combat rogue-subtlety paladin-retribution shaman-enhancement druid-feral hunter-survival}; do
	for pb in ${BANDS:-raid:60 bare:60}; do
		for f in alliance horde; do
			p=${pb%:*}; b=${pb#*:}
			out=$(run -spec "$spec" -preset "$p" -band "$b" -faction "$f" -json) || { echo "$spec $p $b $f error"; continue; }
			python3 -c 'import json,sys; r=json.loads(sys.argv[2]); print("%s dps %.1f %.1f" % (sys.argv[1], r["DPS"], r["Error"]))' "$spec $p $b $f" "$out"
		done
	done
done
for spec in ${TANKS:-warrior-protection paladin-protection druid-feral-bear}; do
	for f in alliance horde; do
		out=$(run -spec "$spec" -preset raid -band 60 -faction "$f" -tank) || { echo "$spec raid 60 $f error"; continue; }
		echo "$spec raid 60 $f $(echo "$out" | grep -m1 '| threat /s |') $(echo "$out" | grep -m1 '| damage taken /s |') $(echo "$out" | grep -m1 '| chance of death |')"
	done
done
