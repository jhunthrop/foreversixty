#!/usr/bin/env bash
# gates.sh: the raid-searches lane gates on the box, from the repo root
# (tools/remote-test.sh --pin-engine --dir . HEAD -- bash design/reviews/2026-11-13-sim-accuracy/raid-searches/gates.sh):
# the sim version unit tests, the request suite with its ladder goldens read
# (not rewritten: Hammer of the Righteous' weapon rule moves no golden), and
# the three search tools and the tools that load the engine.
set -uo pipefail
status=0
run() { echo "### $*"; "$@" || { echo "### FAILED: $*"; status=1; }; }
run make simdb
(cd web && run npx vitest run src/lib/sim 2>&1 | tail -6)
(cd sim && run go vet ./cmd/heal-builds/ ./cmd/rotation-search/ ./cmd/talent-search/ ./enginever/)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./request/ -count=1 2>&1 | tail -4)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./enginever/ ./score/ ./cmd/heal-builds/ ./cmd/rotation-search/ ./cmd/talent-search/ ./cmd/spec-breakdown/ -count=1 2>&1 | tail -10)
echo "gates status $status"
exit $status
