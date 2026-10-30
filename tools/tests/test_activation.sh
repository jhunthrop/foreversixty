#!/usr/bin/env bash
# Tests for tools/activation-gate.sh, the test gate tools/activate-build.sh runs before it
# commits, pushes or tags a client build move. Plain bash, no bats, the shape of
# tools/engine/tests/test_scripts.sh: each case stubs `npm` and `go` (NPM_BIN, GO_BIN) to
# write the reports the real tools write, and checks what the gate makes of them.
#
#   bash tools/tests/test_activation.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly GATE=$HERE/../activation-gate.sh

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
lacks() { [[ $1 != *"$2"* ]]; }

# stub_npm <dir> <exit> <vitest json>: an npm whose `test` writes the JSON report to the
# --outputFile.json path it was given and exits <exit>.
stub_npm() {
	local dir=$1 status=$2 report=$3
	mkdir -p "$dir"
	printf '%s' "$report" >"$dir/vitest.json"
	cat >"$dir/npm" <<EOF
#!/usr/bin/env bash
for arg in "\$@"; do
  case "\$arg" in --outputFile.json=*) cp "$dir/vitest.json" "\${arg#--outputFile.json=}" ;; esac
done
echo "npm \$*" >>"$dir/calls"
exit $status
EOF
	chmod +x "$dir/npm"
}

# stub_go <dir> <exit> <go test -json lines>: a go that prints the events and records the
# packages it was asked for.
stub_go() {
	local dir=$1 status=$2 events=$3
	mkdir -p "$dir"
	printf '%s\n' "$events" >"$dir/events.json"
	cat >"$dir/go" <<EOF
#!/usr/bin/env bash
echo "\$*" >>"$dir/go-calls"
cat "$dir/events.json"
exit $status
EOF
	chmod +x "$dir/go"
}

# gate <stub dir>: runs the gate with both stubs, sets OUT, STATUS and LIST.
gate() {
	local dir=$1
	STATUS=0
	OUT=$(NPM_BIN="$dir/npm" GO_BIN="$dir/go" "$GATE" "$dir/failures" 2>&1) || STATUS=$?
	LIST=$(cat "$dir/failures" 2>/dev/null)
}

readonly WEB_GREEN='{"testResults":[{"name":"/x/web/src/a.test.ts","status":"passed","assertionResults":[{"fullName":"a works","status":"passed"}]}]}'
readonly API_GREEN='{"Action":"pass","Package":"github.com/jhunthrop/foreversixty/api/internal/trees","Test":"TestActive"}'

# Both green.
d=$SCRATCH/green
stub_npm "$d" 0 "$WEB_GREEN"
stub_go "$d" 0 "$API_GREEN"
gate "$d"
assert "green: exits 0" equal "$STATUS" 0
assert "green: says so" contains "$OUT" "are green"
assert "green: lists nothing" equal "$LIST" ""
assert "green: runs the whole vitest suite through npm test" contains "$(cat "$d/calls")" "npm test --"
assert "green: go test runs one package at a time" contains "$(cat "$d/go-calls")" "test -p 1 -count=1 -json"
assert "green: the API packages that read the active build are found" contains "$(cat "$d/go-calls")" "./internal/trees"
assert "green: a package that does not read it is left out" lacks "$(cat "$d/go-calls")" "./internal/httpx"

# A web assertion fails: the gate names its file and full test name.
d=$SCRATCH/web-red
stub_npm "$d" 1 '{"testResults":[{"name":"'"$(cd "$HERE/../.." && pwd)"'/web/src/pages/guides/_spec.test.ts","status":"failed","assertionResults":[{"fullName":"/guides/mage/frost the hit line","status":"failed"},{"fullName":"/guides/mage/frost the header","status":"passed"}]}]}'
stub_go "$d" 0 "$API_GREEN"
gate "$d"
assert "web red: exits 1" equal "$STATUS" 1
assert "web red: names the failing test, file relative to web/" equal "$LIST" "web src/pages/guides/_spec.test.ts > /guides/mage/frost the hit line"
assert "web red: a passing test is not listed" lacks "$LIST" "the header"
assert "web red: prints the list" contains "$OUT" "the hit line"

# A web file that failed without an assertion (an import error) is listed as the suite.
d=$SCRATCH/web-suite
stub_npm "$d" 1 '{"testResults":[{"name":"src/b.test.ts","status":"failed","assertionResults":[]}]}'
stub_go "$d" 0 "$API_GREEN"
gate "$d"
assert "web suite: listed as (suite)" equal "$LIST" "web src/b.test.ts > (suite)"

# npm failed before vitest wrote a report (the pretest sync, say).
d=$SCRATCH/web-noreport
stub_npm "$d" 1 'not json'
stub_go "$d" 0 "$API_GREEN"
gate "$d"
assert "web no report: still fails" equal "$STATUS" 1
assert "web no report: says vitest wrote none" contains "$LIST" "web (suite)"

# An API subtest fails: only the deepest name is listed, with the package short.
d=$SCRATCH/api-red
stub_npm "$d" 0 "$WEB_GREEN"
stub_go "$d" 1 '{"Action":"fail","Package":"github.com/jhunthrop/foreversixty/api/internal/sims","Test":"TestLoot/mage"}
{"Action":"fail","Package":"github.com/jhunthrop/foreversixty/api/internal/sims","Test":"TestLoot"}
{"Action":"fail","Package":"github.com/jhunthrop/foreversixty/api/internal/sims"}'
gate "$d"
assert "api red: exits 1" equal "$STATUS" 1
assert "api red: names the subtest" equal "$LIST" "api api/internal/sims TestLoot/mage"

# An API package that does not build fails with no test named.
d=$SCRATCH/api-build
stub_npm "$d" 0 "$WEB_GREEN"
stub_go "$d" 1 '{"Action":"fail","Package":"github.com/jhunthrop/foreversixty/api/internal/trees"}'
gate "$d"
assert "api build: listed as (suite)" equal "$LIST" "api api/internal/trees (suite)"

# Both red: both listed, web first.
d=$SCRATCH/both
stub_npm "$d" 1 '{"testResults":[{"name":"src/c.test.ts","status":"failed","assertionResults":[{"fullName":"c","status":"failed"}]}]}'
stub_go "$d" 1 '{"Action":"fail","Package":"github.com/jhunthrop/foreversixty/api/internal/bis","Test":"TestPicks"}'
gate "$d"
assert "both red: two lines" equal "$(wc -l <"$d/failures" | tr -d ' ')" 2
assert "both red: web first" equal "$(head -1 "$d/failures")" "web src/c.test.ts > c"

if [[ $FAILED -gt 0 ]]; then
	echo "$FAILED failed, $PASSED passed" >&2
	exit 1
fi
echo "ok $PASSED tests"
