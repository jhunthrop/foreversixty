#!/usr/bin/env bash
# Tests for tools/go-test-shard.sh, which splits api.yml's go test across its shard matrix.
# Plain bash, the shape of test_release_apps.sh: each case feeds a package list and a weights
# file and checks what each shard prints.
#
#   bash tools/tests/test_go_test_shard.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly SPLIT=$HERE/../go-test-shard.sh

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

readonly WEIGHTS=$SCRATCH/weights
printf '%s\n' '# seconds per package' 'm/heavy 300' 'm/mid 200' 'm/light 100' 'm/gone 50' >"$WEIGHTS"
readonly PACKAGES='m/light
m/new1
m/heavy
m/new2
m/mid'

shard() { "$SPLIT" "$1" "$2" "$WEIGHTS" <<<"$PACKAGES"; }

# --- every package lands in exactly one shard ---
all=$(for s in 1 2 3; do shard "$s" 3; done | sort)
assert "the shards together hold every package once" equal "$all" "$(sort <<<"$PACKAGES")"

# --- heaviest first to the lightest shard ---
assert "the heaviest package has a shard to itself" equal "$(shard 1 3)" 'm/heavy'
assert "the next heaviest has the next shard" equal "$(shard 2 3)" 'm/mid'
assert "the light and unlisted packages fill the lightest shard" \
	equal "$(shard 3 3)" $'m/light\nm/new1\nm/new2'

# --- a listed package that no longer exists is ignored ---
assert "a weight for a removed package shards nothing" equal "$(for s in 1 2 3; do shard "$s" 3; done | grep -c gone)" 0

# --- one shard is the whole list ---
assert "one shard tests everything" equal "$(shard 1 1 | sort)" "$(sort <<<"$PACKAGES")"

# --- bad arguments ---
"$SPLIT" 0 4 "$WEIGHTS" </dev/null 2>/dev/null
assert "shard 0 is a usage error" equal $? 64
"$SPLIT" 5 4 "$WEIGHTS" </dev/null 2>/dev/null
assert "a shard past the total is a usage error" equal $? 64
"$SPLIT" 1 4 "$SCRATCH/missing" </dev/null 2>/dev/null
assert "a missing weights file fails" equal $? 1

if ((FAILED > 0)); then
	echo "$FAILED of $((PASSED + FAILED)) tests failed" >&2
	exit 1
fi
echo "ok $PASSED tests"
