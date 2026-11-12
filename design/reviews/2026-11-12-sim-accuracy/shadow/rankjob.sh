#!/usr/bin/env bash
# rankjob.sh: the Shadow lane's rank-row re-judge on the engine with the
# wand's first-shot windup (fea547248). Every Shadow rank row's band edges
# (rankcells.txt: authored spell id, level, rank, race; the row's rank, the
# next rank and the top rank the level can train), bare band-20 gear, seed
# 1, 5,000 iterations, one sim at a time, niced, GOMAXPROCS=5. Run from the
# site's sim directory (tools/remote-test.sh --dir sim --engine-sha <sha>
# HEAD -- bash ../design/reviews/2026-11-12-sim-accuracy/shadow/rankjob.sh).
set -uo pipefail
here=$(cd "$(dirname "$0")" && pwd)
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
while read -r sid lvl rank race; do
	case $race in undead | troll | orc) f=horde ;; *) f=alliance ;; esac
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec priest-shadow -preset bare -band 20 -level "$lvl" -race "$race" -faction $f -rank "$sid=$rank" -seed "${SEED:-1}" -iterations 5000 -json 2>/dev/null |
		python3 -c 'import json,sys; d=json.load(sys.stdin); print(round(d["DPS"],3), round(d["Error"],3), round(d.get("SecondsOOM",0),1), end="")' || printf FAILED
	echo "	$sid level $lvl rank $rank $race"
done <"${CELLS:-$here/rankcells.txt}"
