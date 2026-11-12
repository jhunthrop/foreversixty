#!/usr/bin/env bash
# run.sh <jobs file> <out dir> [iterations]: builds spec-breakdown once and runs each job line in
# turn (one sim at a time, nice 15, GOMAXPROCS=5), saving each job's whole breakdown to
# <out dir>/<label>.txt, printing its DPS line and then the whole file between markers. A job line is: <label> <spec-breakdown flags...>.
# Run from sim/ on the box through tools/remote-test.sh --dir sim --engine-sha <sha>.
set -uo pipefail
jobs=$1; out=$2; iters=${3:-10000}
mkdir -p "$out"
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
while read -r label rest; do
	[[ -z $label || $label == \#* ]] && continue
	# shellcheck disable=SC2086
	if ! nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -iterations "$iters" -seed "${SEED:-1}" $rest > "$out/$label.txt" 2>&1; then
		echo "$label FAILED: $(tail -3 "$out/$label.txt" | tr '\n' ' ')"; continue
	fi
	echo "$label seed ${SEED:-1} $(grep -m1 ' DPS ' "$out/$label.txt" | tr -s ' ')"
	echo "=== begin $label seed ${SEED:-1}"; cat "$out/$label.txt"; echo "=== end $label"
done < "$jobs"
