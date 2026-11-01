#!/usr/bin/env bash
# Tests for tools/engine/follow-inputs-hash.sh, tools/engine/pin-conditions.sh,
# tools/engine/push-review-branch.sh and tools/engine/pinned-build.sh, the scripts
# .github/workflows/engine.yml, bis.yml and tools/activate-build.sh call. Plain bash, no bats: each
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
readonly PUSH=$HERE/../push-review-branch.sh
readonly PINNED=$HERE/../pinned-build.sh
readonly WAIT=$HERE/../ranking-wait.sh
readonly VERDICT=$HERE/../ranking-verdict.sh
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
	mkdir -p "$dir/tools/talentgen/testdata" "$dir/sim/core/spellconst/gen" "$dir/assets/db_inputs"
	git init --quiet "$dir"
	echo 'package main' >"$dir/tools/main.go"
	echo 'package main' >"$dir/tools/talentgen/main.go"
	echo '{"build": "old"}' >"$dir/tools/talentgen/testdata/mage.json"
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

	echo '{"build": "new"}' >"$fork/tools/talentgen/testdata/mage.json"
	commit_all "$fork" "a follow rewrote a talent fixture"
	assert "a talent fixture the follow writes leaves the hash alone" \
		equal "$first" "$(bash "$HASH" "$site" "$fork" "$BUILD")"

	echo 'package main // moved' >"$fork/tools/talentgen/main.go"
	commit_all "$fork" "a generator moved"
	assert "a changed fork generator changes the hash" \
		differ "$first" "$(bash "$HASH" "$site" "$fork" "$BUILD")"
}

# The raw tables are not committed, but the fork's relicmods, periodiccrit,
# eureka and the set rows read them, and pipeline/hotfix_push.py layers the
# server's push into them. Its committed record (hotfix-push.json) moves with
# every push, so a push that changes only those generators still follows.
test_hash_follows_the_hotfix_push_record() {
	local site=$SCRATCH/push-site fork=$SCRATCH/push-fork first
	make_site "$site" aaaaaaaaa
	make_fork "$fork" "$BUILD"
	first=$(bash "$HASH" "$site" "$fork" "$BUILD")
	echo '{"search_totals": {"Spell": 6134}}' >"$site/data/builds/$BUILD/hotfix-push.json"
	commit_all "$site" "the server pushed"
	assert "a new hotfix push record changes the hash" \
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
	assert "absent CI: waits out the grace checks, then gives up" equal "$(cat "$SCRATCH/gh-$CASE/count")" 3
}

test_pin_waits_for_a_fork_ci_not_registered_yet() {
	pin_case late "$BUILD" 000000000
	stub_gh "$SCRATCH/gh-$CASE" '[]' '[{"status":"completed","conclusion":"success"}]'
	run_pin "$SITE" "$FORK" "$BUILD"
	assert "late CI: exits 0 once the run appears and passes (got $PIN_STATUS: $PIN_OUT)" equal "$PIN_STATUS" 0
	assert "late CI: polled twice" equal "$(cat "$SCRATCH/gh-$CASE/count")" 2
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

# make_review_repos <name>: a bare remote and a clone of it with one commit,
# setting REMOTE and WORK.
make_review_repos() {
	REMOTE=$SCRATCH/$1-remote.git
	WORK=$SCRATCH/$1-work
	git init --quiet --bare "$REMOTE"
	git init --quiet "$WORK"
	echo base >"$WORK/f"
	commit_all "$WORK" "base"
}

# commit_as <dir> <author> <message>: one commit by that author.
commit_as() {
	echo "$3" >>"$1/f"
	git -C "$1" add -A
	git -C "$1" -c user.name="$2" -c user.email=x@example.com commit --quiet -m "$3"
}

remote_head() { git --git-dir="$REMOTE" rev-parse --verify --quiet "refs/heads/$1"; }

run_push() {
	PUSH_OUT=$(cd "$WORK" && bash "$PUSH" "$REMOTE" "$1" 2>&1)
	PUSH_STATUS=$?
}

test_push_review_branch_creates_a_missing_branch() {
	make_review_repos new
	run_push follow/b
	assert "no branch: exits 0 (got $PUSH_STATUS: $PUSH_OUT)" equal "$PUSH_STATUS" 0
	assert "no branch: the remote has HEAD" equal "$(remote_head follow/b)" "$(git -C "$WORK" rev-parse HEAD)"
}

test_push_review_branch_replaces_the_bots_own_head() {
	make_review_repos bot
	commit_as "$WORK" forever-bis-bot "last night"
	git -C "$WORK" push --quiet "$REMOTE" HEAD:refs/heads/follow/b
	git -C "$WORK" reset --quiet --hard HEAD~1
	commit_as "$WORK" forever-bis-bot "tonight"
	run_push follow/b
	assert "bot head: exits 0 (got $PUSH_STATUS: $PUSH_OUT)" equal "$PUSH_STATUS" 0
	assert "bot head: the remote has tonight's HEAD" equal "$(remote_head follow/b)" "$(git -C "$WORK" rev-parse HEAD)"
}

test_push_review_branch_keeps_hand_work() {
	local before
	make_review_repos hand
	commit_as "$WORK" forever-bis-bot "last night"
	commit_as "$WORK" "A Person" "hand work"
	git -C "$WORK" push --quiet "$REMOTE" HEAD:refs/heads/follow/b
	before=$(remote_head follow/b)
	git -C "$WORK" reset --quiet --hard HEAD~2
	commit_as "$WORK" forever-bis-bot "tonight"
	run_push follow/b
	assert "hand work: exits 10 (got $PUSH_STATUS: $PUSH_OUT)" equal "$PUSH_STATUS" 10
	assert "hand work: names it" contains "$PUSH_OUT" "hand work"
	assert "hand work: the remote is unchanged" equal "$(remote_head follow/b)" "$before"
}

test_push_review_branch_replaces_merged_hand_work() {
	make_review_repos merged
	commit_as "$WORK" "A Person" "hand work, merged"
	git -C "$WORK" push --quiet "$REMOTE" HEAD:refs/heads/follow/b HEAD:refs/heads/master
	git --git-dir="$REMOTE" symbolic-ref HEAD refs/heads/master
	commit_as "$WORK" forever-bis-bot "tonight"
	run_push follow/b
	assert "merged hand work: exits 0 (got $PUSH_STATUS: $PUSH_OUT)" equal "$PUSH_STATUS" 0
	assert "merged hand work: the remote has tonight's HEAD" equal "$(remote_head follow/b)" "$(git -C "$WORK" rev-parse HEAD)"
}

test_push_review_branch_fails_on_a_bad_remote() {
	make_review_repos bad
	PUSH_OUT=$(cd "$WORK" && bash "$PUSH" "$SCRATCH/no-such-remote.git" follow/b 2>&1)
	PUSH_STATUS=$?
	assert "bad remote: exits 1 (got $PUSH_STATUS)" equal "$PUSH_STATUS" 1
}

test_hash_is_stable_and_follows_the_inputs
test_hash_follows_the_hotfix_push_record
test_hash_refuses_a_build_the_site_does_not_have
test_last_reads_the_newest_follow_inputs_trailer
test_pin_holds_when_all_three_conditions_do
test_pin_refuses_a_failed_fork_ci
test_pin_refuses_a_fork_ci_that_never_ran
test_pin_waits_for_a_fork_ci_not_registered_yet
test_pin_waits_for_a_running_fork_ci

test_pinned_build_reads_the_forks_stamp_at_the_pin() {
	local site=$SCRATCH/pb-site fork=$SCRATCH/pb-fork sha out
	make_fork "$fork" "$OTHER_BUILD"
	sha=$(git -C "$fork" rev-parse --short HEAD)
	echo "$BUILD" >"$fork/assets/db_inputs/FOREVER_BUILD"
	commit_all "$fork" "a later follow the site has not pinned"
	make_site "$site" "$sha"
	out=$(bash "$PINNED" "$site" "$fork")
	assert "the stamp at the pinned sha, not the fork's HEAD (got '$out')" equal "$out" "$OTHER_BUILD"
}

test_pinned_build_asks_github_without_a_fork() {
	local site=$SCRATCH/pb-gh-site stub=$SCRATCH/pb-gh out
	make_site "$site" abcdef123
	stub_gh "$stub" "$BUILD"
	out=$(GH_BIN="$stub/gh" bash "$PINNED" "$site")
	assert "the stamp from GitHub (got '$out')" equal "$out" "$BUILD"
	assert "asks for the file at the pinned sha" contains "$(cat "$stub/calls")" "FOREVER_BUILD?ref=abcdef123"
}

test_pinned_build_fails_rather_than_guess() {
	local site=$SCRATCH/pb-bad-site stub=$SCRATCH/pb-bad fork=$SCRATCH/pb-bad-fork out status
	make_site "$site" abcdef123
	stub_gh "$stub" '{"message": "Not Found"}'
	status=0
	out=$(GH_BIN="$stub/gh" bash "$PINNED" "$site" 2>&1) || status=$?
	assert "a reply that is not a build exits 1" equal "$status" 1
	assert "and says what it read" contains "$out" "not a client build"
	make_fork "$fork" "$BUILD"
	status=0
	out=$(bash "$PINNED" "$site" "$fork" 2>&1) || status=$?
	assert "a sha the fork does not have exits 1" equal "$status" 1
	assert "and names the sha" contains "$out" "abcdef123"
	status=0
	bash "$PINNED" >/dev/null 2>&1 || status=$?
	assert "no arguments is a usage error" equal "$status" 2
}


test_ranking_wait_verdicts() {
	local moved=1730000000 out status
	out=$(bash "$WAIT" "$BUILD" "$BUILD" "$moved" $((moved + 3600)))
	assert "an engine on the active build ranks (got '$out')" equal "${out%%$'\t'*}" rank
	out=$(bash "$WAIT" "$BUILD" "$OTHER_BUILD" "$moved" $((moved + 3 * 3600)))
	assert "the move night's ranking waits (got '$out')" equal "${out%%$'\t'*}" wait
	assert "the wait line says why" contains "$out" "ranking waits for engine follow"
	out=$(bash "$WAIT" "$BUILD" "$OTHER_BUILD" "$moved" $((moved + 26 * 3600)))
	assert "the second night of waiting is stalled (got '$out')" equal "${out%%$'\t'*}" stalled
	assert "the stall names the hours and the fix" contains "$out" "waited for engine follow for 26h"
	out=$(bash "$WAIT" "$BUILD" "" "$moved" $((moved + 60)))
	assert "an unreadable stamp is stalled at once, never a quiet wait" equal "${out%%$'\t'*}" stalled
	out=$(STALL_HOURS=48 bash "$WAIT" "$BUILD" "$OTHER_BUILD" "$moved" $((moved + 26 * 3600)))
	assert "STALL_HOURS moves the line" equal "${out%%$'\t'*}" wait
	status=0
	bash "$WAIT" "$BUILD" "$BUILD" soon later >/dev/null 2>&1 || status=$?
	assert "non-integer epochs are a usage error" equal "$status" 2
}


test_ranking_verdict() {
	local fork=$SCRATCH/rv-fork other=$SCRATCH/rv-other site=$SCRATCH/rv-site osite=$SCRATCH/rv-osite now=1730100000 out status
	make_fork "$fork" "$BUILD"
	make_fork "$other" "$OTHER_BUILD"
	make_site "$site" "$(git -C "$fork" rev-parse --short HEAD)"
	make_site "$osite" "$(git -C "$other" rev-parse --short HEAD)"

	out=$(bash "$VERDICT" "$site" "$fork" "$BUILD" $((now - 3600)) "$now")
	assert "verdict: an engine on the active build ranks (got '$out')" equal "${out%%$'\t'*}" rank
	out=$(bash "$VERDICT" "$site" "$fork" "$BUILD" unknown "$now")
	assert "verdict: and still ranks when the move time is unknown (got '$out')" equal "${out%%$'\t'*}" rank

	out=$(bash "$VERDICT" "$osite" "$other" "$BUILD" $((now - 3 * 3600)) "$now")
	assert "verdict: the move night waits (got '$out')" equal "${out%%$'\t'*}" wait
	out=$(bash "$VERDICT" "$osite" "$other" "$BUILD" $((now - 26 * 3600)) "$now")
	assert "verdict: the second night stalls (got '$out')" equal "${out%%$'\t'*}" stalled

	out=$(bash "$VERDICT" "$osite" "$other" "$BUILD" unknown "$now")
	assert "verdict: a failed move-time lookup stalls, never waits as now (got '$out')" equal "${out%%$'\t'*}" stalled
	assert "verdict: and says the lookup failed" contains "$out" "could not be read"

	out=$(bash "$VERDICT" "$osite" "$SCRATCH/no-such-fork" "$BUILD" $((now - 60)) "$now")
	assert "verdict: an unreadable stamp stalls (got '$out')" equal "${out%%$'\t'*}" stalled

	status=0
	out=$(bash "$VERDICT" "$osite" "$other" "$BUILD" "not-a-time" "$now" 2>&1) || status=$?
	assert "verdict: ranking-wait's own failure is the verdict's failure (got $status)" differ "$status" 0
	assert "verdict: and prints no verdict" lacks "$out" "wait"$'\t'
}

test_pin_gives_up_on_a_fork_ci_that_never_finishes
test_pin_refuses_a_fork_on_another_build
test_pin_refuses_the_sha_already_pinned
test_pin_usage
test_push_review_branch_creates_a_missing_branch
test_push_review_branch_replaces_the_bots_own_head
test_push_review_branch_keeps_hand_work
test_push_review_branch_replaces_merged_hand_work
test_push_review_branch_fails_on_a_bad_remote
test_pinned_build_reads_the_forks_stamp_at_the_pin
test_pinned_build_asks_github_without_a_fork
test_pinned_build_fails_rather_than_guess
test_ranking_wait_verdicts
test_ranking_verdict

if [[ $FAILED -gt 0 ]]; then
	echo "FAILED $FAILED of $((PASSED + FAILED)) assertions" >&2
	exit 1
fi
echo "ok $PASSED tests"
