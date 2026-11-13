#!/usr/bin/env bash
# bis-land.sh: bis.yml's landing of a finished ranking onto the branch as it is NOW.
#
#   tools/bis-land.sh <build> <engine pin> <picks dir> <trailer>
#
# Run from the repository root of the commit job's checkout (any clone whose `origin` the
# job may push to). The ranking takes one to two hours and main moves under it (client
# activations, other lanes, data.yml's own Data.lua commits), so a picks commit made on the
# ranking's base and rebased afterwards conflicts on the derived files (run 43386119123
# threw away a whole night on Data.lua). This never rebases. Each attempt:
#
#   1. fetches the branch's head and checks it out clean in a scratch worktree;
#   2. stops, with a warning and exit 0, when that head's active build is no longer the
#      build that was ranked: picks ranked for an older build never land in a newer one,
#      and the ranking's artifacts stay on the run;
#   3. copies the picks (<spec>.json, <spec>.md) over data/builds/<build>/bis/, and the
#      re-measured dps-per-point.json when there is one, through
#      tools/remeasure-dps-per-point.sh so the marker and the manifest follow the fresh base;
#   4. renders the derived files (addon-data.json, the addon's Data.lua) on that base;
#   5. commits, naming the engine pin the picks were ranked on, and saying so in the body
#      when the branch pins another engine by now (the picks still land; a later ranking
#      replaces them);
#   6. pushes as a fast-forward (with PUBLISH=true; otherwise a `git push --dry-run`, which
#      still asks the remote whether the push would fast-forward). A rejected push starts
#      the next attempt from a fresh fetch.
#
# Environment:
#   BRANCH        the branch to land on (required)
#   PUBLISH       "true" pushes; anything else is the dry run (off main)
#   FIGURES       a re-measured dps-per-point.json to land with the picks (optional)
#   RANKED_HEAD   the sha the ranking started from, for the log (optional)
#   LAND_RENDER   the command that renders the derived files, run from the landing
#                 worktree's root with BUILD exported (tests stub it); default the pipeline
#   LAND_ATTEMPTS how many fetch-apply-render-push rounds (default 5)
#   LAND_SLEEP    seconds between rounds (default 20)
#
# With GITHUB_OUTPUT set, appends landed=true when a commit was pushed, landed=false otherwise.
#
# Exit 0 when the commit landed (or would, off the branch), when there was nothing to
# commit, or when the active build moved; 1 when every attempt was rejected; 2 on a usage
# or input error.
set -euo pipefail

readonly ACTIVE_BUILD_FILE=web/src/data/active-build.json
readonly PIN_FILE=sim/enginever/version.go
readonly DATA_LUA=addon/ForeverSixty/Data.lua

default_render() {
	(cd data && uv sync --quiet && uv run python -m pipeline addon-data --build "$BUILD")
}

active_build() { sed -n 's/.*"build"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' "$1/$ACTIVE_BUILD_FILE"; }
engine_pin() { sed -n 's/.*Version = "\(.*\)"/\1/p' "$1/$PIN_FILE"; }

# apply <worktree>: the picks and the figures over the fresh base, then the derived files.
apply() {
	local tree=$1 dir=$1/data/builds/$BUILD
	cp "$PICKS"/*.json "$PICKS"/*.md "$dir/bis/"
	if [[ -n ${FIGURES:-} ]]; then
		# The re-measure script with the measure replaced by a copy of tonight's figures:
		# it deletes the carried-from marker and re-hashes the manifest of THIS base. A base
		# whose figures are already its own (no marker) keeps them.
		(cd "$tree" && GITHUB_OUTPUT='' DPP_MEASURE="cp $FIGURES data/builds/$BUILD/dps-per-point.json" \
			tools/remeasure-dps-per-point.sh "$BUILD")
	fi
	(cd "$tree" && export BUILD && if [[ -n ${LAND_RENDER:-} ]]; then bash -c "$LAND_RENDER"; else default_render; fi)
}

# message <worktree>: the commit message as -m arguments, one per paragraph, in MSG.
message() {
	local now_pin
	now_pin=$(engine_pin "$1")
	MSG=(-m "chore(bis): nightly leveling BiS for $BUILD @ $PIN")
	if [[ $now_pin != "$PIN" ]]; then
		MSG+=(-m "Ranked on engine $PIN; $BRANCH pins $now_pin since the ranking started, so the next ranking replaces these picks.")
		echo "::notice title=engine pin moved during the ranking::ranked on $PIN, $BRANCH pins $now_pin; the picks land and name $PIN"
	fi
	MSG+=(-m "$TRAILER")
}

# stage <worktree>: the picks, the derived files and, with figures, what the re-measure touched.
stage() {
	local dir=data/builds/$BUILD
	git -C "$1" add "$dir/bis" "$dir/addon-data.json" "$DATA_LUA"
	if [[ -n ${FIGURES:-} ]]; then
		git -C "$1" add -A -- "$dir/dps-per-point.json" "$dir/manifest.json"
		# The marker's deletion, when this base still tracked one.
		if git -C "$1" ls-files --error-unmatch "$dir/dps-per-point.carried-from" >/dev/null 2>&1; then
			git -C "$1" add -A -- "$dir/dps-per-point.carried-from"
		fi
	fi
}

push() {
	if [[ ${PUBLISH:-} == true ]]; then
		git -C "$1" push origin "HEAD:refs/heads/$BRANCH"
	else
		git -C "$1" push --dry-run origin "HEAD:refs/heads/$BRANCH"
	fi
}

# attempt <n> <worktree>: one fetch-apply-render-push round. Sets LANDED when the round
# finished (a sha, "unchanged" or "skipped") and leaves it empty for a retry. Called as a
# plain statement, never in a condition, so `set -e` still stops the script on a failed
# copy, render or commit (a broken render must not land either).
attempt() {
	local n=$1 tree=$2 base active
	LANDED=""
	if ! git fetch --quiet origin "$BRANCH"; then
		echo "attempt $n: could not fetch $BRANCH"
		return 0
	fi
	base=$(git rev-parse FETCH_HEAD)
	git worktree remove --force "$tree" 2>/dev/null || rm -rf "$tree"
	git worktree add --quiet --detach "$tree" "$base"
	echo "attempt $n: landing on $BRANCH at $base${RANKED_HEAD:+ (the ranking started from $RANKED_HEAD, $(git rev-list --count "$RANKED_HEAD..$base" 2>/dev/null || echo '?') commits ago)}"

	active=$(active_build "$tree")
	if [[ $active != "$BUILD" ]]; then
		echo "::warning title=picks not committed::$BRANCH's active build is $active now, the ranking ranked $BUILD; nothing is committed and the picks stay in this run's artifacts"
		LANDED=skipped
		return 0
	fi
	[[ -d $tree/data/builds/$BUILD/bis ]] || { echo "$tree/data/builds/$BUILD/bis does not exist" >&2; exit 2; }

	apply "$tree"
	stage "$tree"
	if git -C "$tree" diff --cached --quiet; then
		echo "nothing to commit: $BRANCH already holds these picks"
		LANDED=unchanged
		return 0
	fi
	message "$tree"
	git -C "$tree" commit --quiet "${MSG[@]}"
	git -C "$tree" show --stat --format='%H (parent %P)%n%B' HEAD
	if push "$tree"; then
		LANDED=$(git -C "$tree" rev-parse HEAD)
	else
		echo "attempt $n: the push was rejected; fetching $BRANCH again"
	fi
}

main() {
	[[ $# -eq 4 ]] || { echo "usage: $0 <build> <engine pin> <picks dir> <trailer>" >&2; exit 2; }
	BUILD=$1 PIN=$2 PICKS=$3 TRAILER=$4
	[[ -n ${BRANCH:-} ]] || { echo "BRANCH is not set" >&2; exit 2; }
	compgen -G "$PICKS/*.json" >/dev/null || { echo "$PICKS holds no picks" >&2; exit 2; }
	if [[ -n ${FIGURES:-} && ! -f $FIGURES ]]; then echo "FIGURES $FIGURES does not exist" >&2; exit 2; fi
	PICKS=$(cd "$PICKS" && pwd)
	if [[ -n ${FIGURES:-} ]]; then FIGURES=$(cd "$(dirname "$FIGURES")" && pwd)/$(basename "$FIGURES"); fi

	local tree n attempts=${LAND_ATTEMPTS:-5}
	tree=$(mktemp -d)/land
	trap 'git worktree remove --force "$tree" 2>/dev/null || true' EXIT
	for ((n = 1; n <= attempts; n++)); do
		attempt "$n" "$tree"
		if [[ -n $LANDED ]]; then
			if [[ ${PUBLISH:-} == true ]]; then echo "landed: $LANDED"; else echo "dry run, would land: $LANDED"; fi
			# landed=true only for a pushed commit: the caller deploys on it.
			if [[ -n ${GITHUB_OUTPUT:-} ]]; then
				if [[ ${PUBLISH:-} == true && $LANDED =~ ^[0-9a-f]{40}$ ]]; then echo "landed=true"; else echo "landed=false"; fi >>"$GITHUB_OUTPUT"
			fi
			exit 0
		fi
		if ((n < attempts)); then sleep "${LAND_SLEEP:-20}"; fi
	done
	echo "::error title=picks not pushed::$attempts attempts to land the picks on $BRANCH were rejected; the picks stay in this run's artifacts"
	exit 1
}

main "$@"
