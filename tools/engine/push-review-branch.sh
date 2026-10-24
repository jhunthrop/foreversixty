#!/usr/bin/env bash
# push-review-branch.sh: push HEAD to a review branch without wiping hand work.
# Used by .github/workflows/engine.yml for the fork's follow/<build> and the
# site's engine-pin/<sha>, which are regenerated every night a follow or a pin
# goes to review.
#
#   tools/engine/push-review-branch.sh <remote-url or remote name> <branch> [bot name]
#
# Run inside the repository whose HEAD is pushed. No remote branch, one whose
# head commit the bot (default forever-bis-bot) authored, or one whose head is
# already merged into the remote's default branch (master on the fork, main on
# the site; a merged review branch is stale, not hand work): push with
# --force-with-lease on the head just read, exit 0. An unmerged head someone
# else authored is hand work: push nothing, exit 10. Any failure: exit 1.
set -euo pipefail

readonly EXIT_PUSHED=0
readonly EXIT_HAND_WORK=10
readonly EXIT_FAILURE=1
readonly DEFAULT_BOT=forever-bis-bot

die() {
	echo "push-review-branch: $*" >&2
	exit "$EXIT_FAILURE"
}

# merged_into_default <remote-url> <sha>: the sha is reachable from the
# remote's default branch. A remote without one has merged nothing.
merged_into_default() {
	git fetch --quiet "$1" HEAD 2>/dev/null || return 1
	git merge-base --is-ancestor "$2" FETCH_HEAD
}

main() {
	[[ $# -eq 2 || $# -eq 3 ]] || { echo "usage: $0 <remote-url> <branch> [bot name]" >&2; exit "$EXIT_FAILURE"; }
	local url=$1 ref=refs/heads/$2 bot=${3:-$DEFAULT_BOT} remote_head author listing
	listing=$(git ls-remote "$url" "$ref") || die "cannot read $ref from the remote"
	remote_head=$(cut -f1 <<<"$listing")
	if [[ -n $remote_head ]]; then
		git fetch --quiet "$url" "$ref" || die "cannot fetch $ref"
		author=$(git log -1 --format=%an "$remote_head")
		if [[ $author != "$bot" ]] && ! merged_into_default "$url" "$remote_head"; then
			echo "push-review-branch: $2 has hand work ($author); not pushed"
			exit "$EXIT_HAND_WORK"
		fi
	fi
	git push --quiet --force-with-lease="$ref:$remote_head" "$url" "HEAD:$ref" || die "cannot push $ref"
	echo "push-review-branch: pushed $2"
	exit "$EXIT_PUSHED"
}

main "$@"
