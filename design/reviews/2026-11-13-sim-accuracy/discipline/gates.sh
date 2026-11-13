#!/usr/bin/env bash
# gates.sh: the Discipline lane's gates on the box, from the repo root
# (tools/remote-test.sh --engine-sha <fork sha> --dir . HEAD -- bash design/reviews/2026-11-13-sim-accuracy/discipline/gates.sh):
# apl-check against the fork, the request suite with the two priest ladder goldens rewritten on the
# new card (printed back as a patch, since the box's worktree is thrown away), the priest card and
# rank tests, the spec-breakdown tests, the guide unit tests, lint and types.
set -uo pipefail
status=0
run() { echo "### $*"; "$@" || { echo "### FAILED: $*"; status=1; }; }
engine=$(cd sim && go list -m -f '{{.Dir}}' github.com/wowsims/classic)
echo "engine dir $engine"
run make simdb
run env ENGINE_DIR="$engine" CURATED_APL_DIR=data/curated/apl CURATED_SPECS_JSON=data/curated/specs.json python3 tools/apl_check.py
(cd sim && run go vet ./request/ ./cmd/spec-breakdown/)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./cmd/spec-breakdown/ -count=1 2>&1 | tail -4)
(cd sim && FOREVER_UPDATE_GOLDEN=1 nice -n 15 env GOMAXPROCS=5 go test ./request/ -run 'TestRotationLadder' -count=1 2>&1 | tail -4)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./request/ -count=1 2>&1 | tail -6)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./rankchoices/ ./leveling/ -count=1 2>&1 | tail -4)
(cd web && run npx vitest run src/content/guides src/lib/guides src/lib/sim 2>&1 | tail -6)
(cd web && run npm run lint --silent 2>&1 | tail -5)
(cd web && run npx astro check 2>&1 | tail -3)
run python3 -m pytest -q data/tests/test_apl.py data/tests/test_card_note_levels.py 2>&1 | tail -3
echo "=====PATCH====="
git diff -- sim/request/testdata/ladder
echo "=====END====="
echo "gates status $status"
exit $status
