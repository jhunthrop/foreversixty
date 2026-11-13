#!/usr/bin/env bash
# run.sh: the rogues-60 lane's sims on the current engine, one at a time,
# niced to 15, GOMAXPROCS=5 (other lanes share the box). Run from the site's
# sim directory (tools/remote-test.sh --dir sim). Prints every breakdown and
# the raid rotation searches (raid build, the Arcane lane's fix).
#   run.sh breakdown | search | all
set -euo pipefail
what=${1:-all}
make -C .. simdb >/dev/null
bin=$(mktemp -d)
go build -o "$bin/spec-breakdown" ./cmd/spec-breakdown
go build -o "$bin/rotation-search" ./cmd/rotation-search
run() { nice --adjustment=15 env GOMAXPROCS=5 "$@"; }
specs="rogue-assassination rogue-combat rogue-subtlety"
if [[ $what == breakdown || $what == all ]]; then
	for s in $specs; do
		echo "=== $s raid 60"; run "$bin/spec-breakdown" -spec "$s" -preset raid -iterations 5000
		echo "=== $s bare 60"; run "$bin/spec-breakdown" -spec "$s" -preset bare -band 60 -iterations 5000
		echo "=== $s bare 50"; run "$bin/spec-breakdown" -spec "$s" -preset bare -band 50 -iterations 5000
		echo "=== $s bare 50 at 55"; run "$bin/spec-breakdown" -spec "$s" -preset bare -band 50 -level 55 -iterations 5000
	done
fi
if [[ $what == search || $what == all ]]; then
	out=$(mktemp -d)
	for s in $specs; do
		run "$bin/rotation-search" -repo-root .. -spec "$s" -preset raid -out "$out" 2>&1 | grep -v '^$' | tail -3
	done
	for f in "$out"/*.md; do echo "=== $(basename "$f")"; cat "$f"; done
fi
