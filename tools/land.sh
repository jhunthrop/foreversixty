#!/usr/bin/env bash
# tools/land.sh <land/branch>: lands a batch the integrator pushed to a land/<n> branch
# (docs/runbooks/landing.md). It waits for every check workflow GitHub started for the push
# of the branch's head sha, stops at the first failed job and lists every failed job, and
# when all of them are green fast-forwards main to that sha (git push origin <sha>:main,
# which the remote refuses unless it is a fast-forward), then deletes the land branch. The
# push to main starts main's own runs and deploys like any other push.
#
#   tools/land.sh land/7
#
# Exit status: 0 landed; 1 a check failed, timed out or the push failed; 2 main moved since
# the branch was cut (merge origin/main into the batch and push a new land branch); 64 usage.
#
# Environment (defaults): GH (gh), GIT (git), REMOTE (origin), LAND_POLL_SECONDS (30),
# LAND_SETTLE_SECONDS (90: how long the set of runs must hold still, all completed, before
# it counts as the whole set; path-filtered workflows that never start leave no run to wait
# for), LAND_TIMEOUT_SECONDS (5400).
set -euo pipefail

readonly GH=${GH:-gh}
readonly GIT=${GIT:-git}
readonly REMOTE=${REMOTE:-origin}
readonly POLL=${LAND_POLL_SECONDS:-30}
readonly SETTLE=${LAND_SETTLE_SECONDS:-90}
readonly TIMEOUT=${LAND_TIMEOUT_SECONDS:-5400}
# A completed job or run with any other conclusion failed the landing.
readonly GREEN_CONCLUSIONS='["success","skipped","neutral"]'

usage() {
	echo "usage: tools/land.sh land/<n>" >&2
	exit 64
}

log() { printf 'land: %s\n' "$*" >&2; }

# head_sha <branch>: the remote branch's head commit, or a failure naming the branch.
head_sha() {
	local sha
	sha=$("$GIT" ls-remote "$REMOTE" "refs/heads/$1" | cut -f1)
	[[ -n $sha ]] || {
		log "$REMOTE has no branch $1"
		return 1
	}
	echo "$sha"
}

# retry <command...>: runs the command up to three times, POLL seconds apart, so one
# network blip does not abandon a half-hour wait.
retry() {
	local attempt
	for attempt in 1 2 3; do
		"$@" && return 0
		((attempt == 3)) || sleep "$POLL"
	done
	log "failed three times: $*"
	return 1
}

# runs_json <branch> <sha>: every push run on the branch for the sha, as gh prints them.
runs_json() {
	"$GH" run list --branch "$1" --commit "$2" --event push --limit 200 \
		--json databaseId,name,status,conclusion
}

# run_jobs_json <id>: the run's name, status and conclusion and its jobs'.
run_jobs_json() {
	"$GH" run view "$1" --json name,status,conclusion,jobs
}

# failed_lines <run json>: one "workflow / job (conclusion) url" line per failed job, or one
# for the run itself when it failed without a failed job (a workflow that never started).
failed_lines() {
	jq -r --argjson green "$GREEN_CONCLUSIONS" '
		. as $run
		| [.jobs[]? | select(.status == "completed" and ((.conclusion // "") as $c | $green | index($c) | not))]
		| if length > 0 then .[] | "\($run.name) / \(.name) (\(.conclusion)) \(.url // "")"
		  elif $run.status == "completed" and (($run.conclusion // "") as $c | $green | index($c) | not)
		  then "\($run.name) (\($run.conclusion))"
		  else empty end'
}

# wait_for_checks <branch> <sha>: polls until every run is completed and green (0) or any
# job failed (1, the failed jobs listed) or the timeout passed (1).
wait_for_checks() {
	local branch=$1 sha=$2 start now runs id failed pending count stable_since=-1 last_count=-1
	start=$SECONDS
	while :; do
		runs=$(retry runs_json "$branch" "$sha")
		count=$(jq length <<<"$runs")
		failed=$(for id in $(jq -r '.[].databaseId' <<<"$runs"); do
			retry run_jobs_json "$id" | failed_lines
		done)
		if [[ -n $failed ]]; then
			log "checks failed on $sha:"
			printf '  %s\n' "${failed//$'\n'/$'\n'  }" >&2
			return 1
		fi
		pending=$(jq '[.[] | select(.status != "completed")] | length' <<<"$runs")
		now=$SECONDS
		if ((count != last_count || pending > 0)); then
			stable_since=$now
			last_count=$count
		fi
		if ((count > 0 && pending == 0 && now - stable_since >= SETTLE)); then
			log "all $count check workflows green on $sha"
			return 0
		fi
		if ((now - start >= TIMEOUT)); then
			log "timed out after ${TIMEOUT}s: $count runs, $pending still running on $sha"
			return 1
		fi
		log "$count runs, $pending running; waiting"
		sleep "$POLL"
	done
}

# fast_forward_main <sha>: pushes the sha to main; 2 when main is no longer its ancestor.
fast_forward_main() {
	local sha=$1
	if "$GIT" push "$REMOTE" "$sha:refs/heads/main"; then
		log "main fast-forwarded to $sha"
		return 0
	fi
	"$GIT" fetch -q "$REMOTE" main
	if ! "$GIT" merge-base --is-ancestor FETCH_HEAD "$sha"; then
		log "main moved since the batch was cut: merge $REMOTE/main into it, push a new land branch and land that"
		return 2
	fi
	log "the push to main failed"
	return 1
}

main() {
	(($# == 1)) || usage
	local branch=$1 sha status=0
	[[ $branch == land/?* ]] || usage
	command -v jq >/dev/null || {
		log "jq is required"
		exit 1
	}
	sha=$(head_sha "$branch")
	log "landing $branch at $sha"
	wait_for_checks "$branch" "$sha" || exit 1
	[[ $(head_sha "$branch") == "$sha" ]] || {
		log "$branch moved while its checks ran; land again"
		exit 1
	}
	fast_forward_main "$sha" || status=$?
	((status == 0)) || exit "$status"
	"$GIT" push -q "$REMOTE" --delete "$branch"
	log "deleted $branch; main's runs and deploys start from the push"
}

main "$@"
