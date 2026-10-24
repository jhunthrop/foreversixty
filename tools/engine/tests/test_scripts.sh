#!/usr/bin/env bash
# Tests for tools/engine/follow-inputs-hash.sh and tools/engine/pin-conditions.sh,
# the two scripts .github/workflows/engine.yml calls. Plain bash, no bats: each
# case builds scratch git repositories under mktemp -d and stubs `gh`.
#
#   bash tools/engine/tests/test_scripts.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each
# failure and exits 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly HASH=$HERE/../follow-inputs-hash.sh
readonly PIN=$HERE/../pin-conditions.sh
readonly BUILD=1.60.1.70532
readonly OTHER_BUILD=1.60.1.70425

SCRATCH=$(mktemp -d)
trap 'rm -rf "$SCRATCH"' EXIT
PASSED=0
FAILED=0

# assert <description> <command...>: runs the command, counts the result.
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

is_hex16() { [[ $1 =~ ^[0-9a-f]{16}$ ]]; }
equal() { [[ $1 == "$2" ]]; }
differ() { [[ $1 != "$2" ]]; }
contains() { [[ $1 == *"$2"* ]]; }
lacks() { [[ $1 != *"$2"* ]]; }

commit_all() {
	git -C "$1" add -A
	git -C "$1" -c user.name=test -c user.email=test@example.com commit --quiet -m "$2"
}

# make_site <dir> <pinned sha>: a site checkout with one build's hashed inputs.
make_site() {
	local dir=$1 pin=$2 sub
	mkdir -p "$dir/sim/enginever"
	git init --quiet "$dir"
	for sub in spellconst talents trainables gametables; do
		mkdir -p "$dir/data/builds/$BUILD/$sub"
		echo "{\"$sub\": 1}" >"$dir/data/builds/$BUILD/$sub/mage.json"
	done
	echo '{}' >"$dir/data/builds/$BUILD/levels.json"
	echo '{}' >"$dir/data/builds/$BUILD/sets.json"
	echo '{}' >"$dir/data/builds/$BUILD/races.json"
	echo 'not hashed' >"$dir/data/builds/$BUILD/loot.json"
	printf 'package enginever\n\nconst Version = "%s"\n' "$pin" >"$dir/sim/enginever/version.go"
	commit_all "$dir" "site"
}

# make_fork <dir> <followed build>: a fork checkout with generator sources.
make_fork() {
	local dir=$1 build=$2
	mkdir -p "$dir/tools" "$dir/sim/core/spellconst/gen" "$dir/assets/db_inputs"
	git init --quiet "$dir"
	echo 'package main' >"$dir/tools/main.go"
	echo 'package gen' >"$dir/sim/core/spellconst/gen/gen.go"
	echo "$build" >"$dir/assets/db_inputs/FOREVER_BUILD"
	commit_all "$dir" "fork"
}

# stub_gh <dir> <json...>: a gh that prints its arguments' JSON in turn, one
# per call, repeating the last; GH_BIN points the scripts at it.
stub_gh() {
	local dir=$1 i=0
	shift
	mkdir -p "$dir"
	for json in "$@"; do
		echo "$json" >"$dir/reply.$i"
		i=$((i + 1))
	done
	cat >"$dir/gh" <<STUB
#!/usr/bin/env bash
echo "\$*" >>"$dir/calls"
n=\$(cat "$dir/count" 2>/dev/null || echo 0)
last=$((i - 1))
[ "\$n" -le "\$last" ] || n=\$last
cat "$dir/reply.\$n"
echo \$((\$(cat "$dir/count" 2>/dev/null || echo 0) + 1)) >"$dir/count"
STUB
	chmod +x "$dir/gh"
}

test_hash_is_stable_and_follows_the_inputs() {
	local site=$SCRATCH/hash-site fork=$SCRATCH/hash-fork clone=$SCRATCH/hash-clone first second edited
	make_site "$site" aaaaaaaaa
	make_fork "$fork" "$BUILD"
	first=$(bash "$HASH" "$site" "$fork" "$BUILD")
	second=$(bash "$HASH" "$site" "$fork" "$BUILD")
	assert "the hash is 16 hex digits (got '$first')" is_hex16 "$first"
	assert "the hash is stable across two calls" equal "$first" "$second"

	git clone --quiet "$site" "$clone"
	echo '{"spellconst": 2}' >"$clone/data/builds/$BUILD/spellconst/mage.json"
	commit_all "$clone" "a spellconst byte moved"
	edited=$(bash "$HASH" "$clone" "$fork" "$BUILD")
	assert "a changed spellconst file changes the hash" differ "$first" "$edited"

	echo 'moved' >"$clone/data/builds/$BUILD/loot.json"
	commit_all "$clone" "an unhashed file moved"
	assert "a file outside the hashed inputs leaves the hash alone" \
		equal "$edited" "$(bash "$HASH" "$clone" "$fork" "$BUILD")"

	echo 'package main // moved' >"$fork/tools/main.go"
	commit_all "$fork" "a generator moved"
	assert "a changed fork generator changes the hash" \
		differ "$first" "$(bash "$HASH" "$site" "$fork" "$BUILD")"
}

test_hash_refuses_a_build_the_site_does_not_have() {
	local site=$SCRATCH/missing-site fork=$SCRATCH/missing-fork out status
	make_site "$site" aaaaaaaaa
	make_fork "$fork" "$BUILD"
	out=$(bash "$HASH" "$site" "$fork" "$OTHER_BUILD" 2>&1)
	status=$?
	assert "a missing build exits non-zero" differ "$status" 0
	assert "a missing build names the build directory" contains "$out" "data/builds/$OTHER_BUILD"
}

test_last_reads_the_newest_follow_inputs_trailer() {
	local fork=$SCRATCH/last-fork
	make_fork "$fork" "$BUILD"
	assert "--last prints none with no trailer" equal "$(bash "$HASH" --last "$fork")" none

	echo x >"$fork/tools/a"
	git -C "$fork" add -A
	git -C "$fork" -c user.name=bot -c user.email=bot@example.com commit --quiet \
		-m "chore(client): the engine follows client $OTHER_BUILD" -m "Follow-Inputs: 1111111111111111"
	echo y >"$fork/tools/b"
	git -C "$fork" add -A
	git -C "$fork" -c user.name=bot -c user.email=bot@example.com commit --quiet \
		-m "chore(client): the engine follows client $BUILD" -m "Follow-Inputs: 0123456789abcdef"
	echo z >"$fork/tools/c"
	commit_all "$fork" "a hand fix after the follow"
	assert "--last prints the newest trailer, past a later commit without one" \
		equal "$(bash "$HASH" --last "$fork")" 0123456789abcdef
}

# run_pin <site> <fork> <build>: sets PIN_OUT and PIN_STATUS.
run_pin() {
	PIN_OUT=$(GH_BIN=$SCRATCH/gh-$CASE/gh PIN_POLL_SECONDS=0 bash "$PIN" "$1" "$2" "$3" 2>&1)
	PIN_STATUS=$?
}

pin_case() {
	CASE=$1
	SITE=$SCRATCH/pin-$CASE-site
	FORK=$SCRATCH/pin-$CASE-fork
	make_fork "$FORK" "$2"
	make_site "$SITE" "${3:-$(git -C "$FORK" rev-parse --short=9 HEAD)}"
}

test_pin_holds_when_all_three_conditions_do() {
	pin_case green "$BUILD" 000000000
	stub_gh "$SCRATCH/gh-$CASE" '[{"status":"completed","conclusion":"success"}]'
	run_pin "$SITE" "$FORK" "$BUILD"
	assert "green: exits 0 (got $PIN_STATUS: $PIN_OUT)" equal "$PIN_STATUS" 0
	assert "green: asks gh about the fork's test workflow at the fork sha" \
		contains "$(cat "$SCRATCH/gh-$CASE/calls")" "$(git -C "$FORK" rev-parse HEAD)"
}

test_pin_refuses_a_failed_fork_ci() {
	pin_case red "$BUILD" 000000000
	stub_gh "$SCRATCH/gh-$CASE" '[{"conclusion":"failure"}]'
	run_pin "$SITE" "$FORK" "$BUILD"
	assert "red CI: exits 1" equal "$PIN_STATUS" 1
	assert "red CI: names the CI condition" contains "$PIN_OUT" "test workflow"
	assert "red CI: names the conclusion" contains "$PIN_OUT" failure
	assert "red CI: does not blame the build" lacks "$PIN_OUT" "FOREVER_BUILD"
}

test_pin_refuses_a_fork_ci_that_never_ran() {
	pin_case absent "$BUILD" 000000000
	stub_gh "$SCRATCH/gh-$CASE" '[]'
	run_pin "$SITE" "$FORK" "$BUILD"
	assert "absent CI: exits 1" equal "$PIN_STATUS" 1
	assert "absent CI: names the CI condition" contains "$PIN_OUT" "test workflow"
}

test_pin_waits_for_a_running_fork_ci() {
	pin_case wait "$BUILD" 000000000
	stub_gh "$SCRATCH/gh-$CASE" '[{"status":"queued","conclusion":""}]' \
		'[{"status":"in_progress","conclusion":""}]' '[{"status":"completed","conclusion":"success"}]'
	run_pin "$SITE" "$FORK" "$BUILD"
	assert "running CI: exits 0 once it passes (got $PIN_STATUS: $PIN_OUT)" equal "$PIN_STATUS" 0
	assert "running CI: polled three times" equal "$(cat "$SCRATCH/gh-$CASE/count")" 3
}

test_pin_gives_up_on_a_fork_ci_that_never_finishes() {
	pin_case stuck "$BUILD" 000000000
	stub_gh "$SCRATCH/gh-$CASE" '[{"status":"in_progress","conclusion":""}]'
	PIN_OUT=$(GH_BIN=$SCRATCH/gh-$CASE/gh PIN_POLL_SECONDS=0 PIN_POLL_ATTEMPTS=4 bash "$PIN" "$SITE" "$FORK" "$BUILD" 2>&1)
	PIN_STATUS=$?
	assert "stuck CI: exits 1" equal "$PIN_STATUS" 1
	assert "stuck CI: says it is still running" contains "$PIN_OUT" in_progress
	assert "stuck CI: polled the attempt budget" equal "$(cat "$SCRATCH/gh-$CASE/count")" 4
}

test_pin_refuses_a_fork_on_another_build() {
	pin_case build "$OTHER_BUILD" 000000000
	stub_gh "$SCRATCH/gh-$CASE" '[{"status":"completed","conclusion":"success"}]'
	run_pin "$SITE" "$FORK" "$BUILD"
	assert "build mismatch: exits 1" equal "$PIN_STATUS" 1
	assert "build mismatch: names the build condition" contains "$PIN_OUT" "FOREVER_BUILD"
	assert "build mismatch: names both builds" contains "$PIN_OUT" "$OTHER_BUILD"
}

test_pin_refuses_the_sha_already_pinned() {
	pin_case same "$BUILD"
	stub_gh "$SCRATCH/gh-$CASE" '[{"status":"completed","conclusion":"success"}]'
	run_pin "$SITE" "$FORK" "$BUILD"
	assert "already pinned: exits 1" equal "$PIN_STATUS" 1
	assert "already pinned: names the pin condition" contains "$PIN_OUT" "already"
}

test_pin_usage() {
	local out status
	out=$(bash "$PIN" 2>&1)
	status=$?
	assert "no arguments: exits 2" equal "$status" 2
	assert "no arguments: prints usage" contains "$out" usage
}

test_hash_is_stable_and_follows_the_inputs
test_hash_refuses_a_build_the_site_does_not_have
test_last_reads_the_newest_follow_inputs_trailer
test_pin_holds_when_all_three_conditions_do
test_pin_refuses_a_failed_fork_ci
test_pin_refuses_a_fork_ci_that_never_ran
test_pin_waits_for_a_running_fork_ci
test_pin_gives_up_on_a_fork_ci_that_never_finishes
test_pin_refuses_a_fork_on_another_build
test_pin_refuses_the_sha_already_pinned
test_pin_usage

if [[ $FAILED -gt 0 ]]; then
	echo "FAILED $FAILED of $((PASSED + FAILED)) assertions" >&2
	exit 1
fi
echo "ok $PASSED tests"
