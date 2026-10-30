#!/usr/bin/env bash
# tools/activation-gate.sh [<failures-file>]
#
# The tests a client build move must pass before tools/activate-build.sh commits, pushes
# or tags anything. Run against the activated tree (active-build.json already names the
# new build), it runs, at the same time:
#
#   web -- the whole vitest suite, exactly what web.yml's `build` job runs before it
#          deploys (`npm test`, whose pretest syncs the fixture data). A subset would let
#          through the one failure that stops the deploy after the addon has shipped,
#          which is what the 2026-10-30 rehearsal found (design/reviews/
#          2026-10-30-launch-rehearsal.md);
#   api -- `go test -p 1` over the API packages that read the active build: every
#          package with a Go file naming active-build, ActiveBuild or data/builds
#          (twenty of them on 2026-10-30). Found by that grep rather than listed, so a
#          new reader is gated the day it lands. The database-backed ones need
#          TEST_DATABASE_URL (client.yml's activate job runs a Postgres service); without
#          it they skip, and the gate says so.
#
# Exit 0 when both are green. Otherwise exit 1, with one failing test per line in
# <failures-file> (default: a temp file whose path is printed) and on stdout, as
# "web <file> > <test>" or "api <package> <Test>"; a suite that failed without naming a
# test (a compile error, a crashed worker) is listed as "(suite)".
#
# NPM_BIN and GO_BIN override the commands (the tests stub them).
set -euo pipefail

readonly NPM=${NPM_BIN:-npm}
readonly GO=${GO_BIN:-go}
readonly API_READS_ACTIVE_BUILD='active-build|ActiveBuild|data/builds'

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
readonly ROOT
FAILURES=${1:-$(mktemp)}
WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT

# The API packages that read the active build, as ./-relative package paths.
api_packages() {
	(cd "$ROOT/api" && { grep -rlE "$API_READS_ACTIVE_BUILD" --include='*.go' . || true; } | xargs -n1 dirname | sort -u)
}

run_web() {
	(cd "$ROOT/web" && "$NPM" test -- --reporter=dot --reporter=json --outputFile.json="$WORK/vitest.json") \
		>"$WORK/web.log" 2>&1
}

run_api() {
	local packages
	packages=$(api_packages)
	[ -n "$packages" ] || { echo "no API package reads the active build" >"$WORK/api.log"; return 1; }
	# shellcheck disable=SC2086 # one argument per package
	(cd "$ROOT/api" && "$GO" test -p 1 -count=1 -json $packages) >"$WORK/api.json" 2>"$WORK/api.log"
}

# web_failures: "web <file> > <test>" per failed assertion, "(suite)" for a file that
# failed without one, from vitest's JSON report.
web_failures() {
	python3 - "$WORK/vitest.json" "$ROOT/web/" <<'PY'
import json, sys
path, prefix = sys.argv[1:3]
try:
    report = json.load(open(path, encoding="utf-8"))
except (OSError, ValueError):
    print("web (suite) vitest wrote no report; see the log")
    sys.exit()
for result in report.get("testResults", []):
    name = result.get("name", "").removeprefix(prefix)
    failed = [a for a in result.get("assertionResults", []) if a.get("status") == "failed"]
    for a in failed:
        print(f"web {name} > {a.get('fullName') or a.get('title')}")
    if not failed and result.get("status") == "failed":
        print(f"web {name} > (suite)")
PY
}

# api_failures: "api <package> <Test>" per failed test, "(suite)" for a package that
# failed without one (it did not build, or a TestMain exited), from go test -json.
api_failures() {
	python3 - "$WORK/api.json" <<'PY'
import json, sys
tests, packages = [], []
for line in open(sys.argv[1], encoding="utf-8", errors="replace"):
    try:
        event = json.loads(line)
    except ValueError:
        continue
    if event.get("Action") != "fail":
        continue
    package = event.get("Package", "?").removeprefix("github.com/jhunthrop/foreversixty/")
    test = event.get("Test")
    if test:
        tests.append((package, test))
    else:
        packages.append(package)
# A subtest's failure fails its parent too; name the deepest test only.
named = {(p, t) for p, t in tests if not any(q == p and u.startswith(t + "/") for q, u in tests)}
for p, t in sorted(named):
    print(f"api {p} {t}")
for p in sorted(set(packages) - {p for p, _ in named}):
    print(f"api {p} (suite)")
PY
}

main() {
	[ -n "${TEST_DATABASE_URL:-}" ] ||
		echo "gate: TEST_DATABASE_URL is unset; the API's database-backed tests skip"
	echo "gate: web vitest and $(api_packages | wc -l | tr -d ' ') API packages, in parallel"
	local web=0 api=0 web_pid api_pid
	run_web &
	web_pid=$!
	run_api &
	api_pid=$!
	wait "$web_pid" || web=$?
	wait "$api_pid" || api=$?

	: >"$FAILURES"
	if [ "$web" -ne 0 ]; then
		web_failures >>"$FAILURES"
		[ -s "$FAILURES" ] || echo "web (suite) npm test exited $web" >>"$FAILURES"
		echo "--- web log (tail) ---"
		tail -n 40 "$WORK/web.log"
	fi
	if [ "$api" -ne 0 ]; then
		local before
		before=$(wc -l <"$FAILURES")
		api_failures >>"$FAILURES"
		[ "$(wc -l <"$FAILURES")" -gt "$before" ] || echo "api (suite) go test exited $api" >>"$FAILURES"
		echo "--- api log (tail) ---"
		tail -n 40 "$WORK/api.log"
	fi
	if [ -s "$FAILURES" ]; then
		echo "gate: $(wc -l <"$FAILURES" | tr -d ' ') failing (listed in $FAILURES):"
		cat "$FAILURES"
		return 1
	fi
	echo "gate: web vitest and the API's active-build tests are green"
}

main
