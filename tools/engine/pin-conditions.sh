#!/usr/bin/env bash
# pin-conditions.sh: may the site pin the fork's checked-out commit?
# (docs/superpowers/specs/2026-10-23-engine-follow-design.md, "The pin step")
#
#   tools/engine/pin-conditions.sh <site-dir> <fork-dir> <active-build>
#
# Three conditions, all required:
#   1. the fork's HEAD is not the sha in <site>/sim/enginever/version.go;
#   2. the fork's assets/db_inputs/FOREVER_BUILD names <active-build>
#      (a half-followed fork is never pinned);
#   3. the fork's `test` workflow concluded success for that sha, waiting
#      for a queued or running one (every PIN_POLL_SECONDS, default 60, up to
#      PIN_POLL_ATTEMPTS, default 30: half an hour; a run not registered yet
#      counts as pending for the first two checks).
# Exit 0 when all hold; exit 1 with one line per failed condition; exit 2 on
# bad arguments. GH_BIN overrides the gh command (the tests stub it).
set -euo pipefail

readonly FORK_REPO=jhunthrop/wowsims-forever
readonly FORK_WORKFLOW=test
readonly FOREVER_BUILD_FILE=assets/db_inputs/FOREVER_BUILD
readonly VERSION_FILE=sim/enginever/version.go
readonly GH=${GH_BIN:-gh}
readonly POLL_SECONDS=${PIN_POLL_SECONDS:-60}
readonly POLL_ATTEMPTS=${PIN_POLL_ATTEMPTS:-30}
# A run is registered a few seconds after the push that starts it, so while
# waiting, the first checks treat "no run yet" as one about to start.
readonly ABSENT_GRACE_ATTEMPTS=2

FAILURES=()

fail() { FAILURES+=("pin: $1"); }

pinned_sha() {
	sed -n 's/.*Version = "\(.*\)"/\1/p' "$1/$VERSION_FILE"
}

# ci_state <sha>: "pending", or the newest run's conclusion, or "absent".
ci_state() {
	local runs
	runs=$("$GH" run list -R "$FORK_REPO" -w "$FORK_WORKFLOW" --commit "$1" --json status,conclusion)
	jq -r 'if length == 0 then "absent"
		elif ((.[0].status // "completed") | IN("queued", "in_progress", "waiting", "pending", "requested")) then "pending"
		else (.[0].conclusion // "absent") end' <<<"$runs"
}

# ci_conclusion <sha> <attempts>: the run's conclusion once it finishes, or
# "still in_progress" when the attempts run out.
ci_conclusion() {
	local sha=$1 attempts=$2 state attempt
	for ((attempt = 1; attempt <= attempts; attempt++)); do
		state=$(ci_state "$sha")
		if [[ $state == absent && $attempts -gt 1 && $attempt -le $ABSENT_GRACE_ATTEMPTS ]]; then
			state=pending
		fi
		[[ $state == pending ]] || { echo "$state"; return; }
		[[ $attempt -eq $attempts ]] || sleep "$POLL_SECONDS"
	done
	echo "still in_progress after $attempts checks"
}

main() {
	[[ $# -eq 3 ]] || { echo "usage: $0 <site-dir> <fork-dir> <active-build>" >&2; exit 2; }
	local site=$1 fork=$2 build=$3 sha pinned followed attempts conclusion
	sha=$(git -C "$fork" rev-parse HEAD)
	pinned=$(pinned_sha "$site")
	[[ -n $pinned ]] || { echo "pin: $site/$VERSION_FILE has no Version" >&2; exit 2; }
	followed=$(cat "$fork/$FOREVER_BUILD_FILE" 2>/dev/null || echo "nothing")

	[[ $sha != "$pinned"* ]] || fail "the fork's ${sha:0:9} is already the site's pin; nothing to pin"
	[[ $followed == "$build" ]] || fail "the fork's $FOREVER_BUILD_FILE names $followed, the site is on $build; the fork is not followed yet"

	# Wait for a running CI only when it is the last thing standing.
	attempts=1
	[[ ${#FAILURES[@]} -gt 0 ]] || attempts=$POLL_ATTEMPTS
	conclusion=$(ci_conclusion "$sha" "$attempts")
	[[ $conclusion == success ]] || fail "the fork's $FORK_WORKFLOW workflow has not passed ${sha:0:9} ($conclusion)"

	if [[ ${#FAILURES[@]} -gt 0 ]]; then
		printf '%s\n' "${FAILURES[@]}"
		exit 1
	fi
	echo "pin: ${sha:0:9} differs from the pin $pinned, follows $build, and passed the fork's $FORK_WORKFLOW workflow"
}

main "$@"
