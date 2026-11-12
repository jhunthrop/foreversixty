#!/usr/bin/env bash
# diag.sh: the three gate failures in full, from the repo root on the box.
set -uo pipefail
make simdb >/dev/null
echo "### astro"; (cd web && npx astro check 2>&1 | grep -v "^\s*$" | grep -i "error\|Result\|hint\|warn" | tail -12)
echo "### vitest"; (cd web && npx vitest run src/lib/sim src/content/guides src/lib/guides --reporter=dot 2>&1 | grep -E "FAIL|✗|×|Error|expected|Test Files|Tests " | head -40)
echo "### addon-builds"; (cd sim && go run ./cmd/addon-builds -repo-root .. 2>&1 | tail -3); git diff --stat -- data/builds/1.60.1.70790/addon-builds.json; git diff -- data/builds/1.60.1.70790/addon-builds.json | grep "^[-+][^-+]" | cut -c1-200 | head -30
