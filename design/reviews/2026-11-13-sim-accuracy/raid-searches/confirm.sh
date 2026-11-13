#!/usr/bin/env bash
# confirm.sh: the raid-searches lane's follow-ups, one at a time, niced to
# 15, GOMAXPROCS=5. Run from the site's sim directory.
#   confirm.sh ret   the Retribution rotation search on an engine sha that
#                    holds Hammer of the Righteous to its one-handed weapon
#                    (remote-test.sh --engine-sha), plus the paladin tests
#   confirm.sh arms  the Arms rotation search confirmed at 5,000 iterations,
#                    seeds 7 and 50001, both factions
set -euo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)
go build -o "$bin/rotation-search" ./cmd/rotation-search
run() { nice --adjustment=15 env GOMAXPROCS=5 "$@"; }
out=$(mktemp -d)
emit() { for r in "$out"/*.md; do echo "=== FILE $1/$(basename "$r")"; cat "$r"; rm "$r"; done; }
case ${1:?ret or arms} in
ret)
	# the paladin goldens write beside their test, so they run from a writable copy of the module: see the report
	for f in alliance horde; do
		run "$bin/rotation-search" -repo-root .. -spec paladin-retribution -faction "$f" -preset raid -out "$out" 2>&1 | grep 'wrote\|rror' || true
		emit "hotr-fix/$f"
	done ;;
arms)
	for seed in 7 50001; do for f in alliance horde; do
		run "$bin/rotation-search" -repo-root .. -spec warrior-arms -faction "$f" -preset raid -confirm 5000 -seed "$seed" -out "$out" 2>&1 | grep 'wrote\|rror' || true
		emit "arms-5000/$f-seed$seed"
	done; done ;;
esac
