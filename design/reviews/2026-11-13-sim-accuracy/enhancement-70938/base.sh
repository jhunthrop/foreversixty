#!/usr/bin/env bash
# base.sh [full]: Enhancement's published entries under the engine this run is
# pinned to: raid 60 and bare 20 to 60, both factions, 10,000 iterations, seed 1,
# one sim at a time, nice 15, GOMAXPROCS=5. With "full", the raid 60 Alliance
# and bare 60 Horde breakdowns print whole (casts, procs, auras). Run from the
# site's sim directory on the box:
#   tools/remote-test.sh --dir sim [--engine-sha <sha>] HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/enhancement-70938/base.sh
set -uo pipefail
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
sb() { # preset band faction
	printf '%s %s %s\t' "$1" "$2" "$3"
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec shaman-enhancement -preset "$1" -band "$2" -faction "$3" -seed 1 -iterations 10000 -json 2>/dev/null |
		python3 -c 'import json,sys; d=json.load(sys.stdin); print("%.1f\t%.1f\t%.1f" % (d["DPS"], d["Error"], d.get("SecondsOOM", 0)))' || echo FAILED
}
for faction in alliance horde; do
	sb raid 60 "$faction"
	for band in 60 50 40 30 20; do sb bare "$band" "$faction"; done
done
if [[ ${1:-} == full ]]; then
	for cell in "raid 60 alliance" "bare 60 horde"; do
		set -- $cell
		echo "### $cell"
		nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec shaman-enhancement -preset "$1" -band "$2" -faction "$3" -seed 1 -iterations 10000 2>&1 | grep -vE '^\s*$'
	done
fi
