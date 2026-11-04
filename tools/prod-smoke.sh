#!/usr/bin/env bash
# tools/prod-smoke.sh: the production smoke check smoke.yml runs after a main deploy in
# launch mode (docs/runbooks/landing.md, "Launch mode"). It fetches the pages a player
# reaches first, one guide per class, a BiS page, the tier list, the simulator, the API's
# health and one /go link, and checks each answers with the expected status and carries its
# key text (a redirect: its Location). Every check runs; a summary lists each one and the
# script exits 1 if any failed.
#
#   tools/prod-smoke.sh
#
# Environment (defaults): SMOKE_SITE (https://foreversixty.gg), SMOKE_API
# (https://api.foreversixty.gg), SMOKE_ATTEMPTS (3) and SMOKE_RETRY_SECONDS (10): a check is
# retried so a deploy still settling at the edge is not a failure, CURL (curl),
# GITHUB_STEP_SUMMARY (the summary is appended there too when set).
set -uo pipefail

readonly SITE=${SMOKE_SITE:-https://foreversixty.gg}
readonly API=${SMOKE_API:-https://api.foreversixty.gg}
readonly ATTEMPTS=${SMOKE_ATTEMPTS:-3}
readonly RETRY=${SMOKE_RETRY_SECONDS:-10}
readonly CURL=${CURL:-curl}

# One check per line: url|expected status (an extended regex)|an extended regex the body
# must match, or for a redirect the Location.
checks() {
	local guide
	cat <<EOF
$SITE/|200|<title[^>]*>[^<]*Forever Sixty
$SITE/guides|200|<h1[^>]*>Guides
EOF
	# One spec guide per class, each with its rotation's Priority list.
	for guide in druid/feral hunter/marksmanship mage/frost paladin/retribution priest/shadow \
		rogue/combat shaman/enhancement warlock/affliction warrior/fury; do
		echo "$SITE/guides/$guide|200|id=\"priority\"[^>]*>Priority<"
	done
	cat <<EOF
$SITE/bis/warrior/fury|200|Best in slot
$SITE/tiers|200|<h1[^>]*>Tier list
$SITE/sim|200|<h1[^>]*>Simulator
$API/health|200|"status":"ok"
$SITE/go/discord|30[1278]|golink=discord
EOF
}

# probe <url>: "<status> <location>" on the first line, the body after it.
probe() {
	local body status
	body=$(mktemp)
	status=$("$CURL" -sS --max-time 30 -o "$body" -w '%{http_code} %{redirect_url}' "$1" 2>&1) || status="000 $status"
	echo "$status"
	cat "$body"
	rm -f "$body"
}

# check <url> <status regex> <text regex>: empty on a pass, the reason on a failure.
check() {
	local url=$1 want_status=$2 want_text=$3 out status location
	out=$(probe "$url")
	read -r status location <<<"$(head -n 1 <<<"$out")"
	if ! [[ $status =~ ^($want_status)$ ]]; then
		echo "status $status, want $want_status"
	elif [[ $status == 3* ]]; then
		[[ $location =~ $want_text ]] || echo "redirects to '$location', want /$want_text/"
	# A here-string, not a pipe: grep -q exits at the first match and, under pipefail, the
	# writer's SIGPIPE would read as a miss.
	elif ! grep -qE -- "$want_text" <<<"${out#*$'\n'}"; then
		echo "the page lacks /$want_text/"
	fi
}

main() {
	local url want_status want_text attempt reason failed=0 total=0 report=''
	while IFS='|' read -r url want_status want_text; do
		total=$((total + 1))
		for ((attempt = 1; ; attempt++)); do
			reason=$(check "$url" "$want_status" "$want_text")
			[[ -n $reason && $attempt -lt $ATTEMPTS ]] || break
			sleep "$RETRY"
		done
		if [[ -n $reason ]]; then
			failed=$((failed + 1))
			report+="| FAIL | $url | $reason |"$'\n'
			echo "::error title=production smoke::$url: $reason"
		else
			report+="| ok | $url | |"$'\n'
		fi
	done < <(checks)
	local summary
	summary=$(printf '### Production smoke: %s of %s failed\n\n| | URL | Why |\n|---|---|---|\n%s' \
		"$failed" "$total" "$report")
	echo "$summary"
	[[ -z ${GITHUB_STEP_SUMMARY:-} ]] || echo "$summary" >>"$GITHUB_STEP_SUMMARY"
	((failed == 0))
}

main "$@"
