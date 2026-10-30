#!/usr/bin/env bash
# pinned-build.sh: which client build is the site's pinned engine stamped for?
#
#   tools/engine/pinned-build.sh <site-dir> [<fork-dir>]
#
# Prints the fork's assets/db_inputs/FOREVER_BUILD at the sha the site pins
# (<site>/sim/enginever/version.go): read from <fork-dir>'s history when one is
# given (bis.yml clones the fork at the pin), else from GitHub (the fork is
# public). The ranking and the dps-per-point re-measure run only when this
# names the active build: an engine stamped for another build reads the new
# client's talent strings by the old trees' indexes (the 2026-10-30 rehearsal,
# design/reviews/2026-10-30-launch-rehearsal.md).
#
# Exit 0 with the build on stdout; exit 1 with the reason on stderr when the
# stamp cannot be read or is not a four-part build (callers treat that as "not
# the active build", never as a match); exit 2 on bad arguments. GH_BIN
# overrides the gh command (the tests stub it).
set -euo pipefail

readonly FORK_REPO=jhunthrop/wowsims-forever
readonly FOREVER_BUILD_FILE=assets/db_inputs/FOREVER_BUILD
readonly VERSION_FILE=sim/enginever/version.go
readonly GH=${GH_BIN:-gh}
readonly BUILD_RE='^[0-9]+(\.[0-9]+){3}$'

main() {
	[[ $# -ge 1 && $# -le 2 ]] || { echo "usage: $0 <site-dir> [<fork-dir>]" >&2; exit 2; }
	local site=$1 fork=${2:-} sha stamp
	sha=$(sed -n 's/.*Version = "\(.*\)"/\1/p' "$site/$VERSION_FILE" 2>/dev/null)
	[[ -n $sha ]] || { echo "pinned-build: $site/$VERSION_FILE has no Version" >&2; exit 1; }
	if [[ -n $fork ]]; then
		stamp=$(git -C "$fork" show "$sha:$FOREVER_BUILD_FILE" 2>/dev/null) ||
			{ echo "pinned-build: $fork has no $FOREVER_BUILD_FILE at $sha" >&2; exit 1; }
	else
		stamp=$("$GH" api -H 'Accept: application/vnd.github.raw' \
			"repos/$FORK_REPO/contents/$FOREVER_BUILD_FILE?ref=$sha" 2>/dev/null) ||
			{ echo "pinned-build: could not read $FOREVER_BUILD_FILE at $sha from $FORK_REPO" >&2; exit 1; }
	fi
	stamp=$(tr -d '[:space:]' <<<"$stamp")
	[[ $stamp =~ $BUILD_RE ]] ||
		{ echo "pinned-build: $FOREVER_BUILD_FILE at $sha is '$stamp', not a client build" >&2; exit 1; }
	echo "$stamp"
}

main "$@"
