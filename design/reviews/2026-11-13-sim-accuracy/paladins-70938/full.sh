#!/usr/bin/env bash
# full.sh <jobs file> [iterations]: builds spec-breakdown once and prints each job's whole breakdown
# (one sim at a time, nice 15, GOMAXPROCS=5). A job line is
# <label> <spec-breakdown flags...>; # starts a comment; @R is this folder. Run from sim/ on the
# box:
#   tools/remote-test.sh --pin-engine [--engine-sha <sha>] --dir sim HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/paladins-70938/full.sh ../design/reviews/2026-11-13-sim-accuracy/paladins-70938/<jobs>.txt
set -uo pipefail
jobs=$1; iters=${2:-10000}
R=../design/reviews/2026-11-13-sim-accuracy/paladins-70938
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
while read -r label rest; do
	[[ -z $label || $label == \#* ]] && continue
	rest=${rest//@R/$R}
	echo "===== $label: $rest"
	# shellcheck disable=SC2086
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -iterations "$iters" -seed 1 $rest 2>&1 | grep -v '^$'
done < "$jobs"
