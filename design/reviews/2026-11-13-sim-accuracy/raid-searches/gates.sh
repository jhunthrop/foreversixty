#!/usr/bin/env bash
# gates.sh: the raid-searches lane gates on the box, from the repo root
# (tools/remote-test.sh --pin-engine --dir . HEAD -- bash design/reviews/2026-11-13-sim-accuracy/raid-searches/gates.sh):
# the sim version unit tests, the request suite with its ladder goldens read
# (not rewritten: Hammer of the Righteous' weapon rule moves no golden), and
# the three search tools and the tools that load the engine.
#   gates.sh [vitest|request|tools ...]   (default: all three)
set -uo pipefail
root=$PWD
status=0
gate() {
	local dir=$1 log
	shift
	log=$(mktemp)
	echo "### $dir: $*"
	if ! (cd "$root/$dir" && "$@") >"$log" 2>&1; then
		status=1
		echo "### FAILED"
		grep -E -- '--- FAIL|FAIL|✗|×|Error|expected|golden|differs' "$log" | head -60
	fi
	tail -8 "$log"
}
make simdb >/dev/null
for g in ${*:-vitest request tools}; do
	case $g in
	vitest) gate web npm test -- src/lib/sim ;;
	request) gate sim nice -n 15 env GOMAXPROCS=5 go test ./request/ -count=1 ;;
	tools)
		gate sim go vet ./cmd/heal-builds/ ./cmd/rotation-search/ ./cmd/talent-search/ ./enginever/
		gate sim nice -n 15 env GOMAXPROCS=5 go test ./enginever/ ./score/ ./cmd/heal-builds/ ./cmd/rotation-search/ ./cmd/talent-search/ -count=1 ;;
	esac
done
echo "gates status $status"
exit $status
