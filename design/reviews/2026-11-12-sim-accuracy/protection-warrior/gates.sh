#!/usr/bin/env bash
# gates.sh: the Protection warrior lane's gates on the box, from the repo root
# (tools/remote-test.sh --pin-engine --dir . HEAD -- bash design/reviews/2026-11-12-sim-accuracy/protection-warrior/gates.sh):
# the web lint and types, the guide and sim unit tests, apl-check against the
# pinned fork, the request suite with the three tank ladder goldens rewritten
# on the new engine (printed back as a patch, since the box's worktree is
# thrown away), and the leveling, score, rank-choice and tool tests.
set -uo pipefail
status=0
run() { echo "### $*"; "$@" || { echo "### FAILED: $*"; status=1; }; }
engine=$(cd sim && go list -m -f '{{.Dir}}' github.com/wowsims/classic)
echo "engine dir $engine"
run make simdb
(cd web && run npm run lint --silent)
(cd web && run npx astro check 2>&1 | tail -3)
(cd web && run npx vitest run src/lib/sim src/content/guides src/lib/guides 2>&1 | tail -8)
run env ENGINE_DIR="$engine" CURATED_APL_DIR=data/curated/apl CURATED_SPECS_JSON=data/curated/specs.json python3 tools/apl_check.py
(cd sim && run go vet ./...)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./request/ -run 'TestRotations|TestAPL|TestRankChoices|TestTank|TestFightSeconds|TestWarrior' -count=1 2>&1 | tail -4)
(cd sim && FOREVER_UPDATE_GOLDEN=1 nice -n 15 env GOMAXPROCS=5 go test ./request/ -run TestRotationLadder -count=1 2>&1 | tail -4)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./request/ -run TestRotationLadder -count=1 2>&1 | tail -3)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./leveling/ ./rankchoices/ ./score/ ./adapter/ ./cmd/spec-breakdown/ ./cmd/leveling-bis/ ./cmd/addon-builds/ -count=1 2>&1 | tail -10)
echo "=====PATCH====="
git diff -- sim/request/testdata/ladder
echo "=====END====="
echo "gates status $status"
exit $status
