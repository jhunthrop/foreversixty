#!/usr/bin/env bash
# jobs.sh <jobs file>: builds spec-breakdown once and runs each job line in turn
# (one sim at a time, niced to 15, GOMAXPROCS=6), printing one line per job:
# the label, effective HPS +/- error, how long the mana lasted and the casts of
# the heals that matter. It syncs the embedded simdb inputs first (make simdb). A job line is: <label> <spec-breakdown flags...>;
# # starts a comment. Run from the site's sim directory (tools/remote-test.sh --dir sim).
set -euo pipefail
jobs=$1
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
while read -r label rest; do
	[[ -z $label || $label == \#* ]] && continue
	# shellcheck disable=SC2086
	out=$(nice -n 15 env GOMAXPROCS=6 "$bin" -repo-root .. -heal -iterations 10000 $rest 2>&1) || { echo "$label FAILED: $(echo "$out" | tail -3)"; continue; }
	hps=$(echo "$out" | sed -n 's/^| effective HPS | \(.*\) |$/\1/p')
	mana=$(echo "$out" | sed -n 's/^| mana lasts (s) | \(.*\) |$/\1/p')
	casts=$(echo "$out" | grep -E '^\| (Wild Growth|Swiftmend|Rejuvenation|Regrowth|Healing Touch|Tranquility|Healing Wave|Chain Heal|Riptide|Lesser Healing Wave|Water Shield)' | sed -n 's/^| \([^|]*\) | \([^|]*\) |.*/\1=\2/p' | tr -d ' ' | paste -sd, -)
	echo "$label hps=$hps mana=$mana $casts"
done < "$jobs"
