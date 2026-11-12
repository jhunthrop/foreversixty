#!/usr/bin/env bash
# gates2.sh: the three gates gates.sh ran without the web data sync, rerun after `npm run sync`, plus the
# addon-builds regeneration (the shaman lane's level 60 orders had not reached the file) printed as a patch.
# From the repo root: tools/remote-test.sh --pin-engine --dir . HEAD -- bash design/reviews/2026-11-12-sim-accuracy/dps-warriors/gates2.sh
set -uo pipefail
status=0
run() { echo "### $*"; "$@" || { echo "### FAILED: $*"; status=1; }; }
run make simdb
(cd web && run npm run sync --silent 2>&1 | tail -2)
(cd web && npx astro check 2>&1 | grep -E "^- [0-9]+ (errors|warnings|hints)|Result" ; echo "### astro check status ${PIPESTATUS[0]}")
(cd web && npx vitest run src/lib/sim src/content/guides src/lib/guides 2>&1 | grep -E "Test Files|Tests |FAIL" | head -12)
(cd sim && run go run ./cmd/addon-builds -repo-root .. 2>&1 | tail -1)
(cd sim && run nice -n 15 env GOMAXPROCS=5 go test ./leveling/ ./cmd/addon-builds/ -count=1 2>&1 | tail -4)
echo "=====PATCH====="
git diff -- data/builds/1.60.1.70790/addon-builds.json
echo "=====END====="
echo "gates2 status $status"
