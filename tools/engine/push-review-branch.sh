#!/usr/bin/env bash
# push-review-branch.sh: push HEAD to a review branch without wiping hand work.
# Used by .github/workflows/engine.yml for the fork's follow/<build> and the
# site's engine-pin/<sha>, which are regenerated every night a follow or a pin
# goes to review.
#
#   tools/engine/push-review-branch.sh <remote-url> <branch> [bot name]
#
# Run inside the repository whose HEAD is pushed. No remote branch, or one whose
# head commit the bot (default forever-bis-bot) authored: push with
# --force-with-lease on the head just read, exit 0. A head someone else authored
# is hand work: push nothing, exit 10. Any failure: exit 1.
set -euo pipefail

readonly EXIT_PUSHED=0
readonly EXIT_HAND_WORK=10
readonly EXIT_FAILURE=1
readonly DEFAULT_BOT=forever-bis-bot

die() {
	echo "push-review-branch: $*" >&2
	exit "$EXIT_FAILURE"
}

main() {
	[[ $# -eq 2 || $# -eq 3 ]] || { echo "usage: $0 <remote-url> <branch> [bot name]" >&2; exit "$EXIT_FAILURE"; }
	local url=$1 ref=refs/heads/$2 bot=${3:-$DEFAULT_BOT} remote_head author listing
	listing=$(git ls-remote "$url" "$ref") || die "cannot read $ref from the remote"
	remote_head=$(cut -f1 <<<"$listing")
	if [[ -n $remote_head ]]; then
		git fetch --quiet "$url" "$ref" || die "cannot fetch $ref"
		author=$(git log -1 --format=%an FETCH_HEAD)
		if [[ $author != "$bot" ]]; then
			echo "push-review-branch: $2 has hand work ($author); not pushed"
			exit "$EXIT_HAND_WORK"
		fi
	fi
	git push --quiet --force-with-lease="$ref:$remote_head" "$url" "HEAD:$ref" || die "cannot push $ref"
	echo "push-review-branch: pushed $2"
	exit "$EXIT_PUSHED"
}

main "$@"
