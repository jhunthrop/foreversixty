#!/usr/bin/env bash
# Tests for tools/prod-smoke.sh, the launch-mode production smoke (docs/runbooks/landing.md,
# "Launch mode"). curl is stubbed (CURL): every page answers 200 with each check's key text,
# /go/ answers 302, and STUB_STATUS / STUB_BODY / STUB_FAIL break the URL that contains
# STUB_URL.
#
#   bash tools/tests/test_prod_smoke.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly SMOKE=$HERE/../prod-smoke.sh

SCRATCH=$(mktemp -d)
trap 'rm -rf "$SCRATCH"' EXIT
PASSED=0
FAILED=0

assert() {
	local what=$1
	shift
	if "$@"; then
		PASSED=$((PASSED + 1))
	else
		FAILED=$((FAILED + 1))
		echo "FAIL: $what" >&2
	fi
}

equal() { [[ $1 == "$2" ]]; }
contains() { [[ $1 == *"$2"* ]]; }

cat >"$SCRATCH/curl" <<'EOF'
#!/usr/bin/env bash
# curl -sS --max-time 30 -o <file> -w <format> <url>
while (($# > 1)); do
  [[ $1 == -o ]] && out=$2
  shift
done
url=$1
status=200 location='' body='<title data-og="title">Forever Sixty</title><h1 class="x">Guides</h1>
<h3 id="priority" data-c>Priority</h3> Best in slot <h1>Tier list</h1><h1>Simulator</h1>
{"ok":true,"data":{"status":"ok"}}'
[[ $url == */go/* ]] && status=302 location="https://site.test/?golink=discord" body=''
if [[ -n ${STUB_URL:-} && $url == *"$STUB_URL"* ]]; then
  [[ -n ${STUB_FAIL:-} ]] && { echo "curl: (6) Could not resolve host" >&2; exit 6; }
  status=${STUB_STATUS:-$status}
  body=${STUB_BODY:-$body}
fi
printf '%s' "$body" >"$out"
printf '%s %s' "$status" "$location"
EOF
chmod +x "$SCRATCH/curl"

# smoke [VAR=value...]: runs the script with the stub; output in $SCRATCH/out, status returned.
smoke() {
	env CURL="$SCRATCH/curl" SMOKE_SITE=https://site.test SMOKE_API=https://api.test \
		SMOKE_RETRY_SECONDS=0 GITHUB_STEP_SUMMARY="$SCRATCH/summary" "$@" \
		bash "$SMOKE" >"$SCRATCH/out" 2>&1
}

# --- green: every page answers with its key text ---
rm -f "$SCRATCH/summary"
smoke
assert "a green smoke exits 0" equal $? 0
out=$(cat "$SCRATCH/out")
assert "it counts sixteen checks" contains "$out" '0 of 16 failed'
assert "it checks one guide per class" contains "$out" 'https://site.test/guides/warlock/affliction'
assert "it checks the API's health" contains "$out" 'https://api.test/health'
assert "the summary reaches the step summary" contains "$(cat "$SCRATCH/summary")" '0 of 16 failed'

# --- a page that answers 500 ---
smoke STUB_URL=/guides/mage/frost STUB_STATUS=500
assert "a 500 fails the smoke" equal $? 1
out=$(cat "$SCRATCH/out")
assert "the failed page is named with its status" contains "$out" '| FAIL | https://site.test/guides/mage/frost | status 500, want 200 |'
assert "it is an annotation" contains "$out" '::error title=production smoke::https://site.test/guides/mage/frost'
assert "the other checks still ran" contains "$out" '1 of 16 failed'

# --- a guide that renders without its Priority list ---
smoke STUB_URL=/guides/rogue/combat STUB_BODY='<h1>Combat Rogue</h1>'
assert "a page without its key text fails" equal $? 1
assert "the missing text is named" contains "$(cat "$SCRATCH/out")" 'https://site.test/guides/rogue/combat | the page lacks'

# --- a /go link that stopped redirecting ---
smoke STUB_URL=/go/discord STUB_STATUS=200
assert "a /go link that answers 200 fails" equal $? 1
assert "it wants a redirect" contains "$(cat "$SCRATCH/out")" 'status 200, want 30[1278]'

# --- the host does not answer at all ---
smoke STUB_URL=api.test STUB_FAIL=1
assert "an unreachable host fails" equal $? 1
assert "it reads as status 000" contains "$(cat "$SCRATCH/out")" '| FAIL | https://api.test/health | status 000'

if ((FAILED > 0)); then
	echo "$FAILED of $((PASSED + FAILED)) tests failed" >&2
	exit 1
fi
echo "ok $PASSED tests"
