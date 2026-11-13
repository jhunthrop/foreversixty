#!/usr/bin/env bash
# variants.sh: Elemental's card against three single-target Chain Lightning
# lines (ele-cl-always: on its cooldown; ele-cl-lm: only inside Lightning
# Mastery; ele-cl-mana: while mana is ahead of the fight), raid 60 both
# factions, bare 60 and bare 50 Horde, 10,000 iterations, seed 1, one sim at
# a time, niced to 15, GOMAXPROCS=5. Run from sim/ (tools/remote-test.sh
# --dir sim --pin-engine).
set -euo pipefail
here=../design/reviews/2026-11-13-sim-accuracy/casters-60
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
run() { # label preset band faction [extra...]
	local label=$1 preset=$2 band=$3 faction=$4; shift 4
	printf '%s %s %s %s ' "$label" "$preset" "$band" "$faction"
	nice --adjustment=15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec shaman-elemental -preset "$preset" -band "$band" -faction "$faction" -seed 1 -iterations 10000 -json "$@" 2>/dev/null | tr -d '\n'
	echo
}
for cell in "raid 60 alliance" "raid 60 horde" "bare 60 horde" "bare 50 horde"; do
	run card $cell
	for v in ele-cl-always ele-cl-lm ele-cl-mana; do run $v $cell -rotation "$here/$v.json"; done
done
