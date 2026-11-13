#!/usr/bin/env bash
# jobs4.sh: the owner's Flame Shock question, three cards on the same gear and builds:
# Earth Shock only (the card before this lane, rot-es-only.json), Flame Shock kept up with
# Earth Shock on the other shock cooldowns (rot-fs-es.json, this lane's card), and Flame
# Shock only (rot-fs-only.json). Raid 60 on the old and the new raid build, bare 60, 50,
# 40 and 30; seeds 1 and 50001 at 60, both factions; 10,000 iterations, nice 15,
# GOMAXPROCS=5, one sim at a time. Run from the site's sim directory on the box:
#   tools/remote-test.sh --dir sim --pin-engine HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/enhancement-70938/jobs4.sh
set -uo pipefail
R=../design/reviews/2026-11-13-sim-accuracy/enhancement-70938
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
sb() {
	local label=$1; shift
	printf '%s\t' "$label"
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec shaman-enhancement -iterations 10000 -json "$@" 2>/dev/null |
		python3 -c 'import json,sys; d=json.load(sys.stdin); print("%.1f\t%.1f\t%.1f" % (d["DPS"], d["Error"], d.get("SecondsOOM", 0)))' || echo FAILED
}
old=0533001000000000-053030310051202051-0530000000000000
new=0535001000000000-054030310051202051-0500000000000000
for card in es-only fs-es fs-only; do
	for faction in alliance horde; do
		for seed in 1 50001; do
			sb "$card raid60-oldbuild $faction $seed" -preset raid -band 60 -faction $faction -seed $seed -talents $old -rotation "$R/rot-$card.json"
			sb "$card raid60-newbuild $faction $seed" -preset raid -band 60 -faction $faction -seed $seed -talents $new -rotation "$R/rot-$card.json"
			sb "$card bare60 $faction $seed" -preset bare -band 60 -faction $faction -seed $seed -rotation "$R/rot-$card.json"
		done
		for band in 50 40 30; do sb "$card bare$band $faction 1" -preset bare -band $band -faction $faction -seed 1 -rotation "$R/rot-$card.json"; done
	done
done
