#!/usr/bin/env bash
# Tests for tools/bis-matrix.sh, bis.yml's rank matrix (one leg per class). Plain bash, the
# shape of test_go_test_shard.sh: each case builds a scratch repository root with a
# specs.json and apl files and checks the matrix the script prints.
#
#   bash tools/tests/test_bis_matrix.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly SCRIPT=$HERE/../bis-matrix.sh

SCRATCH=$(mktemp -d)
trap 'rm -rf "$SCRATCH"' EXIT
PASSED=0
FAILED=0

assert() {
	local what=$1
	shift
	if "$@"; then
		PASSED=$((PASSED + 1))
	else
		FAILED=$((FAILED + 1))
		echo "FAIL: $what" >&2
	fi
}

equal() { [[ $1 == "$2" ]]; }

# A root with four classes; priest-shadow and the rogue are not written, mage-frost has
# no apl file at all. specs.json lists the priests before the mages.
readonly ROOT=$SCRATCH/root
mkdir -p "$ROOT/data/curated/apl"
cat >"$ROOT/data/curated/specs.json" <<'JSON'
[
  {"spec": "priest-holy", "class_slug": "priest"},
  {"spec": "priest-shadow", "class_slug": "priest"},
  {"spec": "mage-fire", "class_slug": "mage"},
  {"spec": "mage-frost", "class_slug": "mage"},
  {"spec": "druid-feral-bear", "class_slug": "druid"},
  {"spec": "rogue-combat", "class_slug": "rogue"},
  {"spec": "priest-discipline", "class_slug": "priest"}
]
JSON
for spec in priest-holy mage-fire druid-feral-bear priest-discipline; do
	echo '{"state": "written"}' >"$ROOT/data/curated/apl/$spec.json"
done
echo '{"state": "draft"}' >"$ROOT/data/curated/apl/priest-shadow.json"
echo '{}' >"$ROOT/data/curated/apl/rogue-combat.json"

# Every written spec, one leg per class, classes in specs.json order, specs in written order.
got=$("$SCRIPT" "$ROOT")
assert "full matrix" equal "$got" \
	'{"include":[{"class":"priest","specs":"priest-holy,priest-discipline"},{"class":"mage","specs":"mage-fire"},{"class":"druid","specs":"druid-feral-bear"}]}'

# A comma list narrows the legs to the named specs (and drops a class with none); the
# legs follow the kept specs' specs.json order, so the druid leg now comes first.
got=$("$SCRIPT" "$ROOT" "druid-feral-bear,priest-discipline")
assert "narrowed matrix" equal "$got" \
	'{"include":[{"class":"druid","specs":"druid-feral-bear"},{"class":"priest","specs":"priest-discipline"}]}'

# Spaces around the names are tolerated (a dispatch input typed by hand).
got=$("$SCRIPT" "$ROOT" "mage-fire, priest-holy")
assert "spaces in the list" equal "$got" \
	'{"include":[{"class":"priest","specs":"priest-holy"},{"class":"mage","specs":"mage-fire"}]}'

# A name that is not written fails, naming it: a draft, a missing apl file, a typo.
for bad in priest-shadow mage-frost rogue-combat preist-holy; do
	err=$("$SCRIPT" "$ROOT" "priest-holy,$bad" 2>&1 >/dev/null)
	status=$?
	assert "$bad is refused" equal "$status" 2
	assert "$bad is named" grep -q "$bad" <<<"$err"
done

# A root with no written spec, and a missing specs.json, fail with exit 2.
mkdir -p "$SCRATCH/empty/data/curated/apl"
echo '[{"spec": "mage-fire", "class_slug": "mage"}]' >"$SCRATCH/empty/data/curated/specs.json"
"$SCRIPT" "$SCRATCH/empty" >/dev/null 2>&1
assert "no written spec exits 2" equal "$?" 2
"$SCRIPT" "$SCRATCH/nowhere" >/dev/null 2>&1
assert "no specs.json exits 2" equal "$?" 2
"$SCRIPT" >/dev/null 2>&1
assert "usage exits 2" equal "$?" 2

if [[ $FAILED -eq 0 ]]; then
	echo "ok $PASSED tests"
else
	echo "$FAILED of $((PASSED + FAILED)) assertions failed" >&2
	exit 1
fi
