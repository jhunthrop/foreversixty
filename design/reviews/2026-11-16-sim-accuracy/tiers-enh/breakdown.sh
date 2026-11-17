#!/usr/bin/env bash
# breakdown.sh <spec>...: raid 60 Alliance full breakdown (casts, procs, auras, damage
# by ability) for each spec named, 3,000 iterations, seed 1, nice 15, GOMAXPROCS 5.
# Run from the site's sim directory on the box:
#   tools/remote-test.sh --dir sim HEAD -- bash ../design/reviews/2026-11-16-sim-accuracy/tiers-enh/breakdown.sh shaman-enhancement
set -uo pipefail
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
for spec in "$@"; do
	echo "### $spec"
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec "$spec" -preset raid -band 60 -faction alliance -seed 1 -iterations 3000 2>&1 | grep -vE '^\s*$'
done
