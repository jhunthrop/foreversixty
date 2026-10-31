#!/usr/bin/env bash
# ranking-wait.sh: may tonight's ranking run, and if not, has it waited too long?
#
#   tools/engine/ranking-wait.sh <active-build> <engine-stamp> <activated-epoch> <now-epoch>
#
# <engine-stamp> is tools/engine/pinned-build.sh's output, or "" when it could not be
# read. <activated-epoch> is when active-build.json last changed (the move's commit).
# Prints one verdict word and a reason, tab-separated:
#
#   rank     the pinned engine is stamped for the active build;
#   wait     it is not, and the move is younger than STALL_HOURS (default 20): the
#            normal state between a move and the engine pin, one night;
#   stalled  it is not and the move is older than that (a second night of waiting: the
#            engine follow went to review and nobody has finished it), or the stamp
#            could not be read at all, which is never a reason to wait quietly.
#
# bis.yml keeps the carried picks on wait and stalled, and fails its run, after
# committing the crawl, on stalled (the 2026-10-30 review, M1). Exit 2 on bad arguments.
set -euo pipefail

readonly STALL_HOURS=${STALL_HOURS:-20}

main() {
	[[ $# -eq 4 ]] || { echo "usage: $0 <active-build> <engine-stamp> <activated-epoch> <now-epoch>" >&2; exit 2; }
	local active=$1 stamp=$2 activated=$3 now=$4 age_hours
	[[ $activated =~ ^[0-9]+$ && $now =~ ^[0-9]+$ ]] || { echo "ranking-wait: the epochs must be integers" >&2; exit 2; }
	if [[ -z $stamp ]]; then
		printf 'stalled\tthe pinned engine'\''s FOREVER_BUILD could not be read, so nothing can rank %s\n' "$active"
		return
	fi
	if [[ $stamp == "$active" ]]; then
		printf 'rank\tthe pinned engine is stamped for %s\n' "$active"
		return
	fi
	age_hours=$(((now - activated) / 3600))
	if ((age_hours < STALL_HOURS)); then
		printf 'wait\tranking waits for engine follow: the pinned engine is stamped for %s, not the active %s (moved %dh ago)\n' "$stamp" "$active" "$age_hours"
	else
		printf 'stalled\tranking has waited for engine follow for %dh: the pinned engine is stamped for %s, not the active %s; finish the engine follow or pin pull request (docs/engine-follow.md)\n' "$age_hours" "$stamp" "$active"
	fi
}

main "$@"
