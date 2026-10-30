#!/usr/bin/env bash
# ranking-verdict.sh: tonight's ranking verdict for bis.yml.
#
#   tools/engine/ranking-verdict.sh <site-dir> <fork-dir> <active-build> <moved-epoch|unknown> [<now-epoch>]
#
# Reads the pinned engine's stamp (pinned-build.sh, from <fork-dir>) and prints
# ranking-wait.sh's "<verdict>\t<line>" for it. <moved-epoch> is when
# web/src/data/active-build.json last changed; bis.yml asks the GitHub API, because its
# checkout is shallow.
#
# "unknown" (the lookup failed) never turns into "now", which would wait quietly forever
# (re-review M3). With an engine on the active build it does not matter: the verdict is
# rank. Otherwise the verdict is stalled, naming the failed lookup, so the run goes red
# that night rather than green every night.
#
# Exit 0 with the verdict; exit non-zero when ranking-wait.sh fails or the arguments are
# wrong, never a verdict-less success (re-review M2: bis.yml once lost that status
# inside a process substitution).
set -euo pipefail

HERE=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly HERE

main() {
	[[ $# -ge 4 && $# -le 5 ]] ||
		{ echo "usage: $0 <site-dir> <fork-dir> <active-build> <moved-epoch|unknown> [<now-epoch>]" >&2; exit 2; }
	local site=$1 fork=$2 build=$3 moved=$4 now=${5:-$(date -u +%s)} stamp out
	stamp=$("$HERE/pinned-build.sh" "$site" "$fork" 2>/dev/null) || stamp=""
	if [[ $moved == unknown ]]; then
		if [[ -n $stamp && $stamp == "$build" ]]; then
			moved=$now
		else
			printf 'stalled\tranking waits for engine follow (the engine is stamped for %s, not %s), and when the move happened could not be read, so the wait cannot be bounded; check the GitHub API token, then rerun bis\n' "${stamp:-an unreadable build}" "$build"
			return
		fi
	fi
	out=$("$HERE/ranking-wait.sh" "$build" "$stamp" "$moved" "$now")
	[[ $out == rank$'\t'* || $out == wait$'\t'* || $out == stalled$'\t'* ]] ||
		{ echo "ranking-verdict: ranking-wait.sh printed no verdict: '$out'" >&2; exit 1; }
	echo "$out"
}

main "$@"
