#!/usr/bin/env bash
# gates.sh: the warrior 70938 lane's site gates on the box, from the repo root
# (tools/remote-test.sh --pin-engine --dir . HEAD -- bash design/reviews/2026-11-13-sim-accuracy/warriors-70938/gates.sh):
# web lint, the data sync, types and the sim and guide unit tests, apl-check against the pinned fork, the curated
# data tests, go vet, the request suite with the ladder goldens rewritten on the new engine (printed back as a
# patch, the box's worktree is thrown away), and the leveling, score, adapter and tool tests.
set -uo pipefail
status=0
run() { echo "### $*"; "$@" || { echo "### FAILED: $*"; status=1; }; }
engine=$(cd sim && go list -m -f '{{.Dir}}' github.com/wowsims/classic)
echo "engine dir $engine"
run make simdb
(cd web && run npm run lint --silent)
(cd web && npm run sync --silent 2>&1 | tail -2)
(cd web && run npx astro check 2>&1 | tail -3)
(cd web && run npx vitest run src/lib/sim src/content/guides src/lib/guides src/fixtures 2>&1 | tail -8)
run env ENGINE_DIR="$engine" CURATED_APL_DIR=data/curated/apl CURATED_SPECS_JSON=data/curated/specs.json python3 tools/apl_check.py
(cd data && run python3 -m pytest -q tests/test_apl.py tests/test_card_note_levels.py tests/test_addonrotation.py 2>&1 | tail -3)
(cd sim && run go vet ./...)
(cd sim && FOREVER_UPDATE_GOLDEN=1 nice -n 15 env GOMAXPROCS=5 go test ./request/ -run TestRotationLadder -count=1 2>&1 | tail -4)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./request/ -count=1 2>&1 | tail -4)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./leveling/ ./rankchoices/ ./score/ ./adapter/ ./api/ ./combine/ ./cmd/spec-breakdown/ ./cmd/leveling-bis/ ./cmd/addon-builds/ ./cmd/talent-search/ -count=1 2>&1 | tail -14)
echo "=====PATCH====="
git diff -- sim/request/testdata/ladder
echo "=====END====="
echo "gates status $status"
exit $status
