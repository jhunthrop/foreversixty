#!/usr/bin/env bash
# jobs.sh <jobs file> [iterations]: builds spec-breakdown once and runs each job line in turn
# (one sim at a time, nice 15, GOMAXPROCS=5), printing one line per job: the label, effective
# HPS +/- error, how long the mana lasted and the casts of the Holy paladin's heals. It syncs
# the embedded simdb inputs first (make simdb). A job line is <label> <spec-breakdown flags...>;
# a hash starts a comment. Run from sim/ on the box through tools/remote-test.sh with
# --dir sim and the engine pinned (the pin, or an --engine-sha for the lane's branch):
#   bash ../design/reviews/2026-11-12-sim-accuracy/holy-paladin/jobs.sh ../design/reviews/2026-11-12-sim-accuracy/holy-paladin/<jobs>.txt
set -uo pipefail
jobs=$1; iters=${2:-10000}
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
heals='Flash of Light|Holy Light|Holy Shock|Divine Favor|Oath of Light|Light.s Vigil'
while read -r label rest; do
	[[ -z $label || $label == \#* ]] && continue
	# shellcheck disable=SC2086
	out=$(nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -heal -iterations "$iters" $rest 2>&1) || { echo "$label FAILED: $(echo "$out" | tail -3 | tr '\n' ' ')"; continue; }
	hps=$(echo "$out" | sed -n 's/^| effective HPS | \(.*\) |$/\1/p')
	mana=$(echo "$out" | sed -n 's/^| mana lasts (s) | \(.*\) |$/\1/p')
	casts=$(echo "$out" | grep -E "^\| ($heals)" | sed -n 's/^| \([^|]*\) | \([^|]*\) |.*/\1=\2/p' | tr -d ' ' | paste -sd, -)
	echo "$label hps=$hps mana=$mana $casts"
done < "$jobs"
