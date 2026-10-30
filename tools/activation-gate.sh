#!/usr/bin/env bash
# tools/activation-gate.sh [<failures-file>]
#
# The tests a client build move must pass before tools/activate-build.sh commits, pushes
# or tags anything. Run against the activated tree (active-build.json already names the
# new build), it runs two legs at the same time:
#
#   web -- what web.yml runs before the site goes live, in this order (one after another:
#          `check`'s presync publishes the real data that the build reads, and the test's
#          presync then swaps in the fixture):
#            npm run check            (astro check; its precheck syncs the real data)
#            npm run lint
#            npm run format:check     (the re-encode and the link rewrite edit markdown)
#            CF_PAGES=1 npm run build (the real-data build deploy runs: FOREVER_DATA and
#                                      PUBLIC_SIM_ENGINE unset, so no wasm is needed, and
#                                      every guide, BiS and planner page prerenders from
#                                      the new build; vitest reads the fixture, so only
#                                      this catches a page the new client breaks)
#            npm test                 (the whole vitest suite)
#          Every command runs even after one fails, so one run lists every red;
#          the 2026-10-30 rehearsal found a red vitest case that stopped the deploy
#          after the addon had shipped (design/reviews/2026-10-30-launch-rehearsal.md);
#          the review of this gate (docs/superpowers/reports/
#          2026-10-30-client-hardening-review.md, H1) found the build was missing;
#   api -- `go test -p 1` over the API packages that read the active build: every
#          package with a Go file naming active-build, ActiveBuild or data/builds
#          (twenty of them on 2026-10-30). Found by that grep rather than listed, so a
#          new reader is gated the day it lands. The database-backed ones need
#          TEST_DATABASE_URL (client.yml's activate job runs a Postgres service); without
#          it they skip, and the gate says so. Not api.yml's whole `-race ./...` and
#          `go vet`: an API red after the tag does not strand the addon the way a web red
#          does, and the API image is rebuilt from main on dispatch either way.
#
# Exit 0 when both are green. Otherwise exit 1, with one failure per line in
# <failures-file> (default: a temp file whose path is printed) and on stdout, as
# "web <file> > <test>", "web (npm run <script>) exited <n>" or "api <package> <Test>";
# a suite that failed without naming a test (a compile error, a crashed worker) is
# listed as "(suite)".
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

# run_web: the web leg. Each red script's name goes to $WORK/web-red (vitest's as
# "test"), its output to $WORK/web-<script>.log.
run_web() {
	local script status
	: >"$WORK/web-red"
	for script in check lint format:check build; do
		status=0
		if [ "$script" = build ]; then
			(cd "$ROOT/web" && env -u FOREVER_DATA -u PUBLIC_SIM_ENGINE CF_PAGES=1 "$NPM" run build) \
				>"$WORK/web-$script.log" 2>&1 || status=$?
		else
			(cd "$ROOT/web" && "$NPM" run "$script") >"$WORK/web-$script.log" 2>&1 || status=$?
		fi
		[ "$status" -eq 0 ] || echo "$script $status" >>"$WORK/web-red"
	done
	status=0
	(cd "$ROOT/web" && "$NPM" test -- --reporter=dot --reporter=json --outputFile.json="$WORK/vitest.json") \
		>"$WORK/web-test.log" 2>&1 || status=$?
	[ "$status" -eq 0 ] || echo "test $status" >>"$WORK/web-red"
	[ ! -s "$WORK/web-red" ]
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
	echo "gate: web (check, lint, format, the real build, vitest) and $(api_packages | wc -l | tr -d ' ') API packages, in parallel"
	local web=0 api=0 web_pid api_pid script status before
	run_web &
	web_pid=$!
	run_api &
	api_pid=$!
	wait "$web_pid" || web=$?
	wait "$api_pid" || api=$?

	: >"$FAILURES"
	if [ "$web" -ne 0 ]; then
		[ -s "$WORK/web-red" ] || echo "web (suite) the web leg exited $web" >>"$FAILURES"
		while read -r script status; do
			if [ "$script" = test ]; then
				before=$(wc -l <"$FAILURES")
				web_failures >>"$FAILURES"
				[ "$(wc -l <"$FAILURES")" -gt "$before" ] || echo "web (suite) npm test exited $status" >>"$FAILURES"
			else
				echo "web (npm run $script) exited $status" >>"$FAILURES"
			fi
			echo "--- web $script log (tail) ---"
			tail -n 40 "$WORK/web-$script.log"
		done <"$WORK/web-red"
	fi
	if [ "$api" -ne 0 ]; then
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
	echo "gate: web (check, lint, format, the real build, vitest) and the API's active-build tests are green"
}

main
