#!/usr/bin/env bash
# tools/dev-stack/smoke.sh
#
# Signs in as the seeded guild master through the dev sign-in and checks that GET /v1/me
# shows the account, its characters and the still-unclaimed guild, and that the guild's
# progression credits the two seeded boss kills to the guild master. Exits non-zero on the
# first thing that is missing. Needs the stack up and seeded (up.sh, then seed.sh).
set -euo pipefail
. "$(dirname "${BASH_SOURCE[0]}")/lib.sh"

api="http://127.0.0.1:${API_PORT}"
jar="$(mktemp)"
trap 'rm -f "$jar"' EXIT

fail() { echo "smoke: FAIL: $*" >&2; exit 1; }

api_healthy || fail "the API is not answering on $api (run tools/dev-stack/up.sh and seed.sh)"

code="$(curl -sS -o /dev/null -w '%{http_code}' -c "$jar" -X POST "$api/v1/dev/login" \
  -H 'Content-Type: application/json' -d "{\"battletag\":\"$GM_BATTLETAG\"}")"
[ "$code" = 200 ] || fail "POST /v1/dev/login as $GM_BATTLETAG answered $code (is the stack seeded?)"

me="$(curl -fsS -b "$jar" "$api/v1/me")" || fail "GET /v1/me failed after signing in"
want() { printf '%s' "$me" | grep -qF "$1" || fail "GET /v1/me does not contain $1"; }
want "\"battletag\":\"$GM_BATTLETAG\""
want '"key":"us/pvp/bow-jackzon"'
want '"key":"us/pvp/pipsqueak"'
want '"name":"Lifecycle Test"'
want '"rank":"leader"'
echo "smoke: $GM_BATTLETAG is signed in with Bow Jackzon and Pipsqueak, and leads Lifecycle Test"

gid="$(printf '%s' "$me" | grep -o '"guild":{"id":[0-9]*' | head -1 | grep -o '[0-9]*$')"
[ -n "$gid" ] || fail "GET /v1/me names no guild id"

progression="$(curl -fsS "$api/v1/guilds/$gid/progression")" || fail "GET /v1/guilds/$gid/progression failed"
printf '%s' "$progression" | grep -qF '"kills":2' || fail "the guild's progression does not show 2 kills"
printf '%s' "$progression" | grep -qF '"name":"Bow Jackzon"' || fail "the best parse is not credited to Bow Jackzon"
echo "smoke: guild $gid shows 2 Onyxia kills credited to Bow Jackzon"

# The dev sign-in never creates an account.
code="$(curl -sS -o /dev/null -w '%{http_code}' -X POST "$api/v1/dev/login" \
  -H 'Content-Type: application/json' -d '{"battletag":"Nobody#0000"}')"
[ "$code" = 404 ] || fail "an unknown battletag answered $code, want 404"
echo "smoke: ok"
