#!/usr/bin/env bash
# jobs.sh <jobs file> [iterations]: builds spec-breakdown once and runs each job line in turn
# (one sim at a time, nice 15, GOMAXPROCS=5), printing the label and the DPS line, or for a
# -tank job the damage taken, TMI, chance of death and threat lines. A job line is
# <label> <spec-breakdown flags...>; # starts a comment; @R is this folder. Run from sim/ on the
# box:
#   tools/remote-test.sh --pin-engine [--engine-sha <sha>] --dir sim HEAD -- bash ../design/reviews/2026-11-12-sim-accuracy/paladins/jobs.sh ../design/reviews/2026-11-12-sim-accuracy/paladins/<jobs>.txt
set -uo pipefail
jobs=$1; iters=${2:-10000}
R=../design/reviews/2026-11-12-sim-accuracy/paladins
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
while read -r label rest; do
	[[ -z $label || $label == \#* ]] && continue
	rest=${rest//@R/$R}
	# shellcheck disable=SC2086
	out=$(nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -iterations "$iters" -seed 1 $rest 2>&1) || { echo "$label FAILED: $(echo "$out" | tail -3 | tr '\n' ' ')"; continue; }
	if [[ $rest == *-tank* ]]; then
		echo "$label $(echo "$out" | grep -E '^\| (damage taken /s|TMI|chance of death|threat /s) \|' | tr -s ' ' | tr '\n' ' ')"
	else
		echo "$label $(echo "$out" | grep -m1 ' DPS ' | tr -s ' ')"
	fi
done < "$jobs"
