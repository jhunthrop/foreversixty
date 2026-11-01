#!/usr/bin/env bash
# follow-inputs-hash.sh: the skip-when-unchanged key for .github/workflows/engine.yml
# (docs/superpowers/specs/2026-10-23-engine-follow-design.md, "The workflow").
#
#   tools/engine/follow-inputs-hash.sh <site-dir> <fork-dir> <build>
#     prints a 16-hex hash of the build id, the site's committed
#     damage-relevant inputs for that build and the fork's generator sources
#   tools/engine/follow-inputs-hash.sh --last <fork-dir>
#     prints the Follow-Inputs: trailer of the newest fork commit carrying
#     one (searching the last 50), or "none"
#
# The same shape as bis.yml's "Inputs:" hash: git object ids, not file
# contents, so the hash reads only what is committed. The raw tables are not
# in the tree and are not hashed; a new build always follows because its id is,
# and a new hotfix push for the same build because hotfix-push.json moves.
set -euo pipefail

# The build's inputs a follow copies or generates from, relative to data/builds/<build>/.
# hotfix-push.json stands in for the raw tables the fork's relicmods, periodiccrit,
# eureka and set rows read: they are not committed, and the server's hotfix push
# (data/pipeline/hotfix_push.py) is all that moves them within one build.
readonly SITE_INPUTS=(spellconst talents trainables gametables levels.json sets.json races.json hotfix-push.json)
# The fork's generator sources: a change to them can change every generated file.
readonly FORK_GENERATORS=(tools sim/core/spellconst/gen)
# Under tools/ but written by the follow itself (follow.sh copies the build's
# talent trees there), so hashing it would make every follow look unfollowed.
readonly FORK_FOLLOW_OUTPUTS=tools/talentgen/testdata/
readonly TRAILER=Follow-Inputs
readonly TRAILER_SEARCH_DEPTH=50
readonly HASH_LENGTH=16

usage() {
	echo "usage: $0 <site-dir> <fork-dir> <build> | --last <fork-dir>" >&2
	exit 2
}

sha256() {
	if command -v sha256sum >/dev/null 2>&1; then sha256sum; else shasum -a 256; fi
}

last_trailer() {
	local fork=$1 found
	found=$(git -C "$fork" log --format=%B -n "$TRAILER_SEARCH_DEPTH" | sed -n "s/^$TRAILER: //p" | head -1)
	echo "${found:-none}"
}

inputs_hash() {
	local site=$1 fork=$2 build=$3 paths=() input tree generators
	for input in "${SITE_INPUTS[@]}"; do paths+=("data/builds/$build/$input"); done
	tree=$(git -C "$site" ls-tree -r HEAD -- "${paths[@]}")
	[[ -n $tree ]] || { echo "the site has no committed data/builds/$build inputs at $site" >&2; exit 1; }
	generators=$(git -C "$fork" ls-tree -r HEAD -- "${FORK_GENERATORS[@]}" |
		awk -F '\t' -v skip="$FORK_FOLLOW_OUTPUTS" 'index($2, skip) != 1')
	[[ -n $generators ]] || { echo "the fork at $fork has no generator sources (${FORK_GENERATORS[*]})" >&2; exit 1; }
	printf '%s\n%s\n%s\n' "$build" "$tree" "$generators" | sha256 | cut -c1-"$HASH_LENGTH"
}

main() {
	if [[ ${1:-} == --last ]]; then
		[[ $# -eq 2 ]] || usage
		last_trailer "$2"
	else
		[[ $# -eq 3 ]] || usage
		inputs_hash "$1" "$2" "$3"
	fi
}

main "$@"
