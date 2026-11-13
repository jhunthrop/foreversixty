#!/usr/bin/env bash
# gates2.sh: the web gates gates.sh ran without the data sync (astro check, the sim and guide vitest), rerun
# after `npm run sync`. From the repo root:
#   tools/remote-test.sh --pin-engine --dir . HEAD -- bash design/reviews/2026-11-13-sim-accuracy/crit-flags/gates2.sh
set -uo pipefail
make simdb >/dev/null
(cd web && npm run sync --silent 2>&1 | tail -2)
(cd web && npx astro check 2>&1 | grep -E "^- [0-9]+ (errors|warnings|hints)|Result"; echo "### astro check status ${PIPESTATUS[0]}")
(cd web && npx vitest run src/lib/sim src/content/guides src/lib/guides 2>&1 | grep -E "Test Files|Tests |FAIL" | head -12; echo "### vitest status ${PIPESTATUS[0]}")
