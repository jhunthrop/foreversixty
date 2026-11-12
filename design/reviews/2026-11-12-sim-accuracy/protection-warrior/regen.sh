#!/usr/bin/env bash
# regen.sh: after the walk's first points moved, regenerate what follows the order on the box, from the repo root
# (tools/remote-test.sh --pin-engine --dir . HEAD -- bash design/reviews/2026-11-12-sim-accuracy/protection-warrior/regen.sh):
# the addon builds file, the Protection ladder golden, and the leveling, addon-builds and request tests; the
# changed files come back as a patch, since the box's worktree is thrown away.
set -uo pipefail
status=0
run() { echo "### $*"; "$@" || { echo "### FAILED: $*"; status=1; }; }
run make simdb
(cd sim && run go run ./cmd/addon-builds -repo-root ..)
(cd sim && FOREVER_UPDATE_GOLDEN=1 nice -n 15 env GOMAXPROCS=5 go test ./request/ -run TestRotationLadder -count=1 2>&1 | tail -3)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./leveling/ ./cmd/addon-builds/ ./cmd/leveling-bis/ -count=1 2>&1 | tail -6)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./request/ -run 'TestRotationLadder|TestRotations|TestAPL|TestRankChoices' -count=1 2>&1 | tail -3)
echo "=====PATCH====="
git diff -- data/builds sim/request/testdata/ladder
echo "=====END====="
echo "regen status $status"
exit $status
