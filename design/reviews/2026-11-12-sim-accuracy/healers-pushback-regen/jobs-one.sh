#!/usr/bin/env bash
# jobs-one.sh: one full spec-breakdown -heal print per healer at raid 60 (2,000 iterations), to read every
# mana income and spend row. Run from sim/ on the box as jobs.sh is.
set -uo pipefail
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
for spec in priest-holy priest-discipline druid-restoration shaman-restoration; do
	echo "##### $spec raid 60 alliance"
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -heal -iterations 2000 -spec "$spec" -preset raid -band 60 -faction alliance -seed 1 2>&1
done
