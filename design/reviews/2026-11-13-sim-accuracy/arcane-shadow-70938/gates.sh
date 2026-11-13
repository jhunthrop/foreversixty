#!/usr/bin/env bash
# gates.sh: the arcane-shadow lane gates on the box, from the repo root
# (tools/remote-test.sh --pin-engine --dir . HEAD -- bash design/reviews/2026-11-13-sim-accuracy/arcane-shadow-70938/gates.sh):
# web lint, the data sync, astro check and the guide and sim unit tests, apl-check against the pinned fork,
# go vet, the request suite with the ladder goldens and the adapter goldens rewritten on the new engine
# (printed back as a patch, since the box's worktree is thrown away), and the leveling, score, rank-choice
# and tool tests.
set -uo pipefail
status=0
run() { echo "### $*"; "$@" || { echo "### FAILED: $*"; status=1; }; }
engine=$(cd sim && go list -m -f '{{.Dir}}' github.com/wowsims/classic)
echo "engine dir $engine"
run make simdb
(cd web && run npm run lint --silent)
(cd web && npm run sync --silent 2>&1 | tail -2)
(cd web && npx astro check 2>&1 | grep -E "^- [0-9]+ (errors|warnings|hints)|Result"; echo "### astro check status ${PIPESTATUS[0]}")
(cd web && npx vitest run src/lib/sim src/content/guides src/lib/guides 2>&1 | grep -E "Test Files|Tests |FAIL|✗|×" | head -20; echo "### vitest status ${PIPESTATUS[0]}")
run env ENGINE_DIR="$engine" CURATED_APL_DIR=data/curated/apl CURATED_SPECS_JSON=data/curated/specs.json python3 tools/apl_check.py
(cd sim && run go vet ./...)
(cd sim && FOREVER_UPDATE_GOLDEN=1 nice -n 15 env GOMAXPROCS=5 go test ./request/ -run TestRotationLadder -count=1 2>&1 | tail -4)
(cd sim && FOREVER_UPDATE_GOLDEN=1 nice -n 15 env GOMAXPROCS=5 go test ./adapter/ -count=1 2>&1 | tail -4)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./request/ -count=1 2>&1 | tail -4)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./leveling/ ./rankchoices/ ./score/ ./adapter/ ./cmd/spec-breakdown/ ./cmd/leveling-bis/ ./cmd/addon-builds/ ./cmd/talent-search/ ./cmd/rotation-search/ -count=1 2>&1 | tail -14)
echo "=====PATCH====="
git diff -- sim/request/testdata sim/adapter/testdata
echo "=====END====="
echo "gates status $status"
exit $status
