#!/usr/bin/env bash
# webgates.sh: the lane's web and data gates with their full failure output, from the repo root on the box.
set -uo pipefail
cd web
npx vitest run src/content/guides src/lib/guides src/lib/sim 2>&1 | sed 's/\x1b\[[0-9;]*m//g' | grep -E "FAIL|✗|×|AssertionError|Expected|Received|Test Files|Tests  |toBe|Error" | head -60
echo "vitest exit ${PIPESTATUS[0]}"
npx astro check 2>&1 | sed 's/\x1b\[[0-9;]*m//g' | grep -E "error|Error|Result" | head -30
echo "astro exit ${PIPESTATUS[0]}"
cd .. && (cd data && uv run pytest -q --no-cov tests/test_apl.py tests/test_card_note_levels.py 2>&1 | tail -15); echo "pytest exit $?"
