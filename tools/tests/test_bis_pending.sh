#!/usr/bin/env bash
# Tests for tools/bis-pending.sh, bis.yml's list of the specs still to rank on these inputs.
#
#   bash tools/tests/test_bis_pending.sh
set -uo pipefail
HERE=$(cd "$(dirname "$0")" && pwd)
readonly SCRIPT=$HERE/../bis-pending.sh
SCRATCH=$(mktemp -d)
trap 'rm -rf "$SCRATCH"' EXIT
PASSED=0
FAILED=0
check() {
	if [[ $2 == "$3" ]]; then PASSED=$((PASSED + 1)); else FAILED=$((FAILED + 1)); echo "FAIL: $1: got '$2', want '$3'" >&2; fi
}
readonly MATRIX='{"include":[{"class":"mage","specs":"mage-arcane,mage-fire"},{"class":"rogue","specs":"rogue-combat"}]}'
bis=$SCRATCH/data/builds/b1/bis
mkdir -p "$bis"
run() { (cd "$SCRATCH" && "$SCRIPT" b1 "$1" "$MATRIX"); }

check "no markers: every spec, in matrix order" "$(run h1)" "mage-arcane,mage-fire,rogue-combat"
echo h1 >"$bis/mage-arcane.inputs"
echo h0 >"$bis/rogue-combat.inputs"
check "a matching marker is done, another hash is not" "$(run h1)" "mage-fire,rogue-combat"
echo h1 >"$bis/mage-fire.inputs"
echo h1 >"$bis/rogue-combat.inputs"
check "every marker matches: nothing" "$(run h1)" ""
check "new inputs: everything again" "$(run h2)" "mage-arcane,mage-fire,rogue-combat"
status=0
"$SCRIPT" b1 h1 >/dev/null 2>&1 || status=$?
check "usage: two arguments exit 2" "$status" 2
mode=$(git -C "$HERE/../.." ls-files -s tools/bis-pending.sh tools/bis-inputs-hash.sh | cut -c1-6 | sort -u)
check "the scripts are committed executable" "$mode" 100755

if ((FAILED > 0)); then echo "$FAILED of $((PASSED + FAILED)) failed" >&2; exit 1; fi
echo "ok $PASSED tests"
