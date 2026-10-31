#!/usr/bin/env bash
# The client watcher's alert: a GitHub issue that stays open while a client build move
# is blocked, so a stopped move is never only a line in a run summary (29-31 Oct 2026 it
# failed daily unseen and the site sat two builds behind). client.yml runs it.
#
#   tools/client-alert.sh blocked <report.md>   a move stopped: open "Client build <build>
#                                               is blocked" (label client-build-blocked),
#                                               or comment on the open one
#   tools/client-alert.sh cleared <build>       a move landed: close every open
#                                               client-build-blocked issue, naming <build>
#
# <report.md> is the "Client move to <build> stopped" report (step, reason, hand command).
# The alert adds how far behind the site is and the run link. Environment:
#   BUILD         the build the move was for (blank when the watch failed before naming it)
#   ACTIVE        the build the site serves (blank: read web/src/data/active-build.json)
#   NEWEST        the newest Forever build wago.tools lists (blank when unknown)
#   BEHIND_SINCE  when wago first listed a Forever build newer than ACTIVE (UTC)
#   BEHIND_DAYS   whole days since then (blank when unknown)
#   RUN_URL       this run's page
#   GH_BIN        the gh to call (tests stub it); gh reads GH_TOKEN, which is never printed
set -euo pipefail

readonly LABEL=client-build-blocked
readonly LABEL_COLOR=B60205
readonly LABEL_DESCRIPTION="The client watcher (client.yml) could not move the site to a new client build"
readonly ACTIVE_BUILD_FILE=web/src/data/active-build.json
GH=${GH_BIN:-gh}

usage() { echo "usage: $0 blocked <report.md> | cleared <build>" >&2; exit 2; }

issue_title() { echo "Client build ${1:-unknown} is blocked"; }

active_build() {
	if [ -n "${ACTIVE:-}" ]; then echo "$ACTIVE"; return; fi
	sed -n 's/.*"build": *"\([^"]*\)".*/\1/p' "$ACTIVE_BUILD_FILE" 2>/dev/null || true
}

# behind_line: the site's build against the newest, and for how long it has lagged.
behind_line() {
	local active newest=${NEWEST:-} days=${BEHIND_DAYS:-}
	active=$(active_build)
	if [ -z "$newest" ]; then
		echo "**Behind:** the site serves ${active:-an unknown build}; the newest Forever build is unknown (the watch did not report one)."
	elif [ -z "$days" ]; then
		echo "**Behind:** the site serves ${active:-an unknown build}; the newest Forever build is $newest."
	else
		local unit=days
		[ "$days" != 1 ] || unit=day
		echo "**Behind:** the site serves ${active:-an unknown build}; the newest Forever build is $newest; behind for $days $unit (since ${BEHIND_SINCE:-?} UTC)."
	fi
}

# alert_body <report.md>: the report, the lag and the run link.
alert_body() {
	cat "$1"
	echo
	behind_line
	if [ -n "${RUN_URL:-}" ] && ! grep -qF "$RUN_URL" "$1"; then
		echo
		echo "Run: $RUN_URL"
	fi
}

# open_issue <title>: the number of the open issue for that title, found by label then
# by title (an issue whose label was removed); blank when none. A blank build's title
# takes any open labelled issue, since the watch did not say which build it stopped on.
open_issue() {
	local title=$1 any=$2 found
	found=$("$GH" issue list --state open --label "$LABEL" --limit 100 --json number,title \
		--jq '.[] | "\(.number)\t\(.title)"' | first_match "$title" "$any")
	if [ -z "$found" ]; then
		found=$("$GH" issue list --state open --search "\"$title\" in:title" --limit 20 --json number,title \
			--jq '.[] | "\(.number)\t\(.title)"' | first_match "$title" false)
	fi
	echo "$found"
}

# first_match <title> <any>: from "number<TAB>title" lines, the first number whose title
# is exactly <title> (or the first at all when <any> is true).
first_match() {
	awk -F '\t' -v want="$1" -v any="$2" 'any == "true" || $2 == want { print $1; exit }'
}

blocked() {
	local report=$1 title any=false number body
	if [ ! -s "$report" ]; then
		echo "client-alert: no report at $report; the alert says only that the move stopped" >&2
		report=$(mktemp)
		printf '## Client move to %s stopped\n\nNo report was written; the run log has the step and the reason.\n' "${BUILD:-a new build}" >"$report"
	fi
	title=$(issue_title "${BUILD:-}")
	[ -n "${BUILD:-}" ] || any=true
	body=$(mktemp)
	alert_body "$report" >"$body"
	number=$(open_issue "$title" "$any")
	if [ -n "$number" ]; then
		"$GH" issue comment "$number" --body-file "$body" >/dev/null
		echo "client-alert: commented on #$number ($title)"
	else
		"$GH" label create "$LABEL" --color "$LABEL_COLOR" --description "$LABEL_DESCRIPTION" --force >/dev/null
		"$GH" issue create --title "$title" --label "$LABEL" --body-file "$body"
		echo "client-alert: opened \"$title\""
	fi
	rm -f "$body"
}

cleared() {
	local build=$1 number note
	note="Client build $build is active: the client watcher moved the site to it${RUN_URL:+ ($RUN_URL)}. Closing."
	for number in $("$GH" issue list --state open --label "$LABEL" --limit 100 --json number --jq '.[].number'); do
		"$GH" issue close "$number" --comment "$note" >/dev/null
		echo "client-alert: closed #$number"
	done
}

[ $# -eq 2 ] || usage
case $1 in
blocked) blocked "$2" ;;
cleared) cleared "$2" ;;
*) usage ;;
esac
