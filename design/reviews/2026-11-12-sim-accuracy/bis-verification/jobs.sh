#!/usr/bin/env bash
# jobs.sh <jobs file> [iterations] [seed]: builds spec-breakdown once and runs each job line in
# turn (one sim at a time, nice 15, GOMAXPROCS=5) at the ranker's own count and seed (3000, 7)
# unless told otherwise, printing the label and the figure the ranker publishes for that entry:
# the DPS line, the effective HPS line (-heal) or the tank lines (-tank). A job line is
# <label> <spec-breakdown flags...>; # starts a comment. FULL=1 prints each job's whole output
# between "### <label>" markers instead (the tank figures and score). Run from sim/ on the box:
#   tools/remote-test.sh --engine-sha <ranked engine sha> --dir sim HEAD -- \
#     bash ../design/reviews/2026-11-12-sim-accuracy/bis-verification/jobs.sh \
#     ../design/reviews/2026-11-12-sim-accuracy/bis-verification/<jobs>.txt
set -uo pipefail
jobs=$1; iters=${2:-3000}; seed=${3:-7}
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
while read -r label rest; do
	[[ -z $label || $label == \#* ]] && continue
	rest=${rest%%#*}
	# shellcheck disable=SC2086
	out=$(nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -iterations "$iters" -seed "$seed" $rest 2>&1) || { echo "$label FAILED: $(echo "$out" | tail -3 | tr '\n' ' ')"; continue; }
	if [[ ${FULL:-0} == 1 ]]; then
		echo "### $label"; echo "$out"
	elif [[ $rest == *-tank* ]]; then
		echo "$label $(echo "$out" | grep -E '^\| (damage taken /s|TMI|chance of death|threat /s) \|' | tr -s ' ' | tr '\n' ' ')"
	elif [[ $rest == *-heal* ]]; then
		echo "$label $(echo "$out" | grep -m1 'effective HPS' | tr -s ' ')"
	else
		echo "$label $(echo "$out" | grep -m1 ' DPS ' | tr -s ' ')"
	fi
done < "$jobs"
