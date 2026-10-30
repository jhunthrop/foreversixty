#!/usr/bin/env bash
# Tests for tools/activation-gate.sh, the test gate tools/activate-build.sh runs before it
# commits, pushes or tags a client build move, and for activate-build.sh's failure
# report. Plain bash, no bats, the shape of tools/engine/tests/test_scripts.sh: the gate
# cases stub `npm` and `go` (NPM_BIN, GO_BIN) to write the reports the real tools write;
# the report cases run the real script in a scratch repository whose steps are stubs.
#
#   bash tools/tests/test_activation.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly GATE=$HERE/../activation-gate.sh
readonly ACTIVATE=$HERE/../activate-build.sh
readonly PINNED=$HERE/../engine/pinned-build.sh
readonly OLD_BUILD=1.60.1.70658
readonly NEW_BUILD=1.60.1.70707

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

# stub_npm <dir> <test exit> <vitest json> [<red script>...]: an npm whose `test` writes
# the JSON report to the --outputFile.json path it was given and exits <test exit>, and
# whose `run <script>` exits 1 for each named red script and 0 otherwise. Each call is
# recorded with the CF_PAGES, FOREVER_DATA and PUBLIC_SIM_ENGINE it saw.
stub_npm() {
	local dir=$1 status=$2 report=$3
	shift 3
	mkdir -p "$dir"
	printf '%s' "$report" >"$dir/vitest.json"
	printf '%s\n' "$@" >"$dir/red"
	cat >"$dir/npm" <<EOF
#!/usr/bin/env bash
echo "npm \$* [CF_PAGES=\${CF_PAGES:-} FOREVER_DATA=\${FOREVER_DATA:-} PUBLIC_SIM_ENGINE=\${PUBLIC_SIM_ENGINE:-}]" >>"$dir/calls"
if [ "\$1" = run ]; then
  if grep -qxF "\$2" "$dir/red"; then echo "\$2 failed: prerender of /guides/mage/frost threw"; exit 1; fi
  exit 0
fi
for arg in "\$@"; do
  case "\$arg" in --outputFile.json=*) cp "$dir/vitest.json" "\${arg#--outputFile.json=}" ;; esac
done
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
assert "green: runs astro check" contains "$(cat "$d/calls")" "npm run check"
assert "green: runs lint" contains "$(cat "$d/calls")" "npm run lint"
assert "green: runs the format check" contains "$(cat "$d/calls")" "npm run format:check"
assert "green: the real build runs as deploy runs it" contains "$(cat "$d/calls")" "npm run build [CF_PAGES=1 FOREVER_DATA= PUBLIC_SIM_ENGINE=]"
assert "green: check's real sync runs before the build, vitest's fixture sync after" \
	equal "$(sed -E 's/ \[.*//; s/ --.*//' "$d/calls" | tr '\n' ',')" "npm run check,npm run lint,npm run format:check,npm run build,npm test,"
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

# A real-build prerender failure is named, with the scripts after it still run.
d=$SCRATCH/web-build
stub_npm "$d" 0 "$WEB_GREEN" build
stub_go "$d" 0 "$API_GREEN"
STATUS=0
OUT=$(FOREVER_DATA=fixture PUBLIC_SIM_ENGINE=wasm NPM_BIN="$d/npm" GO_BIN="$d/go" "$GATE" "$d/failures" 2>&1) || STATUS=$?
LIST=$(cat "$d/failures")
assert "build red: exits 1" equal "$STATUS" 1
assert "build red: names the build" equal "$LIST" "web (npm run build) exited 1"
assert "build red: quotes its log" contains "$OUT" "prerender of /guides/mage/frost threw"
assert "build red: an inherited FOREVER_DATA and engine are cleared for the build" contains "$(cat "$d/calls")" "npm run build [CF_PAGES=1 FOREVER_DATA= PUBLIC_SIM_ENGINE=]"
assert "build red: vitest still runs" contains "$(cat "$d/calls")" "npm test --"

# Two red scripts and a red test: all listed.
d=$SCRATCH/web-many
stub_npm "$d" 1 '{"testResults":[{"name":"src/d.test.ts","status":"failed","assertionResults":[{"fullName":"d","status":"failed"}]}]}' check format:check
stub_go "$d" 0 "$API_GREEN"
gate "$d"
assert "many red: three lines" equal "$(wc -l <"$d/failures" | tr -d ' ')" 3
assert "many red: check, format, then the test" equal "$(tr '\n' '|' <"$d/failures")" "web (npm run check) exited 1|web (npm run format:check) exited 1|web src/d.test.ts > d|"

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


# make_move_repo <dir>: a scratch checkout with the active build, the new build's fetched
# tables and an addon TOC, and a re-encode stub that fails the way a talent-tree change does.
make_move_repo() {
	local dir=$1 f
	mkdir -p "$dir/tools" "$dir/web/src/data" "$dir/web/scripts" "$dir/data/builds/$NEW_BUILD" \
		"$dir/data/builds/$OLD_BUILD" "$dir/addon/ForeverSixty"
	cp "$ACTIVATE" "$dir/tools/activate-build.sh"
	printf '{\n  "build": "%s"\n}\n' "$OLD_BUILD" >"$dir/web/src/data/active-build.json"
	for f in items.json simdb.bin talents.json levels.json spellranks.json loot.json manifest.json; do
		echo '{}' >"$dir/data/builds/$NEW_BUILD/$f"
	done
	printf '## Interface: 16001\n## Version: 2.1.7\n' >"$dir/addon/ForeverSixty/ForeverSixty.toc"
	cat >"$dir/web/scripts/reencode-guide-builds.mjs" <<'EOF'
console.error('mage/frost build: Arcane Focus is not in 1.60.1.70707\'s Arcane tree');
process.exit(1);
EOF
	git init --quiet "$dir"
	git -C "$dir" add -A
	git -C "$dir" -c user.name=test -c user.email=test@example.com commit --quiet -m repo
}

# activate <dir> <args...>: runs the scratch copy with a step summary, sets STATUS, OUT, SUMMARY.
activate() {
	local dir=$1
	shift
	STATUS=0
	OUT=$(cd "$dir" && GITHUB_STEP_SUMMARY="$dir/summary.md" PRODUCT=wowf bash tools/activate-build.sh "$@" 2>&1) || STATUS=$?
	SUMMARY=$(cat "$dir/summary.md" 2>/dev/null)
}

# The re-encode stops on a tree change: step, reason and the hand command are named.
d=$SCRATCH/move-reencode
make_move_repo "$d"
activate "$d" "$NEW_BUILD" --no-commit --no-gate
assert "re-encode: exits 1" equal "$STATUS" 1
assert "re-encode: the summary names the step" contains "$SUMMARY" "**Step:** re-encoding the guide builds onto $NEW_BUILD's talent trees"
assert "re-encode: the reason quotes the script" contains "$SUMMARY" "Arcane Focus is not in 1.60.1.70707"
assert "re-encode: the hand command re-spends then reruns the move" contains "$SUMMARY" "then from main with those edits in the tree run: tools/activate-build.sh $NEW_BUILD"
assert "re-encode: or dispatches client with the build and product" contains "$SUMMARY" "gh workflow run client.yml -f build=$NEW_BUILD -f product=wowf"
assert "re-encode: a local run is told how to undo the half-move" contains "$SUMMARY" "git restore --staged --worktree --"
assert "re-encode: the same report goes to stderr" contains "$OUT" "Client move to $NEW_BUILD stopped"
assert "re-encode: nothing was committed" equal "$(git -C "$d" rev-list --count HEAD)" 1

# A long report keeps its head (what failed) and its tail.
d=$SCRATCH/move-long
make_move_repo "$d"
cat >"$d/web/scripts/reencode-guide-builds.mjs" <<'EOF'
console.error('9 guide build(s) do not carry cleanly');
for (let i = 1; i <= 150; i++) console.error(`line ${i}`);
process.exit(1);
EOF
activate "$d" "$NEW_BUILD" --no-commit --no-gate
assert "long: the head line stays" contains "$SUMMARY" "9 guide build(s) do not carry cleanly"
assert "long: the tail stays" contains "$SUMMARY" "line 150"
assert "long: the middle is elided" contains "$SUMMARY" "lines omitted"
assert "long: and left out" lacks "$SUMMARY" "line 60"$'\n'


# make_full_move_repo <dir>: a scratch checkout every step of a move passes in, up to the
# gate: stub uv and go on PATH (stubs/), a re-encode that carries cleanly, the three TOCs,
# the changelog, the API fixture and one API package that reads the active build; plus a
# bare origin (<dir>.git) to show what a run pushed or tagged.
make_full_move_repo() {
	local dir=$1 toc
	make_move_repo "$dir"
	cp "$GATE" "$dir/tools/activation-gate.sh"
	mkdir -p "$dir/tools/engine" "$dir/stubs" "$dir/web/src/content/guides/mage" "$dir/api/internal/trees" \
		"$dir/api/internal/bnetbuild/testdata" "$dir/addon/ForeverSixtyData" "$dir/addon/ForeverSixtyInbox" \
		"$dir/data/curated/coach" "$dir/sim/enginever"
	cp "$PINNED" "$dir/tools/engine/pinned-build.sh"
	echo 'console.log("43 guide build(s) re-stamped");' >"$dir/web/scripts/reencode-guide-builds.mjs"
	echo "see https://foreversixty.gg/data/$OLD_BUILD/talents" >"$dir/web/src/content/guides/mage/frost.md"
	echo 'package trees // reads web/src/data/active-build.json' >"$dir/api/internal/trees/trees.go"
	echo "FS1:$OLD_BUILD:warrior:human:x" >"$dir/api/internal/bnetbuild/testdata/era-kiloz.fs1"
	for toc in addon/ForeverSixtyData/ForeverSixtyData.toc addon/ForeverSixtyInbox/ForeverSixtyInbox.toc; do
		printf '## Interface: 16001\n' >"$dir/$toc"
	done
	printf '# Changelog\n\n## 2.1.7\n\nOlder.\n' >"$dir/addon/ForeverSixty/CHANGELOG.md"
	printf 'package enginever\n\nconst Version = "abcdef123"\n' >"$dir/sim/enginever/version.go"
	printf '#!/usr/bin/env bash\nexit 0\n' >"$dir/stubs/uv"
	printf '#!/usr/bin/env bash\nexit 0\n' >"$dir/stubs/go"
	chmod +x "$dir/stubs/uv" "$dir/stubs/go"
	git -C "$dir" add -A
	git -C "$dir" -c user.name=test -c user.email=test@example.com commit --quiet -m "the rest of the move's inputs"
	git clone --quiet --bare "$dir" "$dir.git"
	git -C "$dir" remote add origin "$dir.git"
	git -C "$dir" branch --quiet -M main
	git -C "$dir" push --quiet origin main
}

# A real-build prerender failure stops the move before the commit, the push and the tag.
d=$SCRATCH/move-prerender
make_full_move_repo "$d"
stub_npm "$d/npm-stub" 0 "$WEB_GREEN" build
stub_go "$d/go-stub" 0 "$API_GREEN"
before=$(git -C "$d.git" rev-parse main)
STATUS=0
OUT=$(cd "$d" && PATH="$d/stubs:$PATH" NPM_BIN="$d/npm-stub/npm" GO_BIN="$d/go-stub/go" \
	GITHUB_STEP_SUMMARY="$d/summary.md" PRODUCT=wowf bash tools/activate-build.sh "$NEW_BUILD" 2>&1) || STATUS=$?
assert "prerender: the move exits 1" equal "$STATUS" 1
assert "prerender: it reached the gate (the earlier steps passed)" contains "$OUT" "== the test gate"
assert "prerender: the summary names the real build" contains "$(cat "$d/summary.md")" "web (npm run build) exited 1"
assert "prerender: nothing was pushed" equal "$(git -C "$d.git" rev-parse main)" "$before"
assert "prerender: no tag on origin" equal "$(git -C "$d.git" tag -l 'addon-v*')" ""
assert "prerender: no tag locally" equal "$(git -C "$d" tag -l 'addon-v*')" ""
assert "prerender: nothing was committed" equal "$(git -C "$d" rev-parse HEAD)" "$before"

# The same move with a green gate commits, pushes and tags (the test above is not green by
# accident of a broken fixture).
d=$SCRATCH/move-green
make_full_move_repo "$d"
stub_npm "$d/npm-stub" 0 "$WEB_GREEN"
stub_go "$d/go-stub" 0 "$API_GREEN"
printf '#!/usr/bin/env bash\nexit 1\n' >"$d/stubs/gh"
chmod +x "$d/stubs/gh"
STATUS=0
OUT=$(cd "$d" && PATH="$d/stubs:$PATH" NPM_BIN="$d/npm-stub/npm" GO_BIN="$d/go-stub/go" \
	GITHUB_STEP_SUMMARY="$d/summary.md" bash tools/activate-build.sh "$NEW_BUILD" 2>&1) || STATUS=$?
assert "green move: pushed the move" contains "$(git -C "$d.git" log -1 --format=%s main)" "activate build $NEW_BUILD"
assert "green move: tagged the addon" equal "$(git -C "$d.git" tag -l 'addon-v*')" "addon-v2.1.8"
assert "green move: the ranking waits (the stub engine stamp is unreadable)" contains "$OUT" "ranking waits for engine follow"
assert "green move: the failed dispatches fail the run" equal "$STATUS" 1



# Under Actions: an error annotation, and no local undo line.
d=$SCRATCH/move-actions
make_move_repo "$d"
STATUS=0
OUT=$(cd "$d" && GITHUB_ACTIONS=true GITHUB_STEP_SUMMARY="$d/summary.md" bash tools/activate-build.sh "$NEW_BUILD" --no-commit --no-gate 2>&1) || STATUS=$?
assert "actions: an error annotation names the step" contains "$OUT" "::error title=client move to $NEW_BUILD stopped at: re-encoding the guide builds"
assert "actions: no undo line on a throwaway runner" lacks "$(cat "$d/summary.md")" "undo the half-made move"
assert "actions: without PRODUCT the redo names the build" contains "$(cat "$d/summary.md")" "gh workflow run client.yml -f build=$NEW_BUILD"
assert "actions: and no empty product" lacks "$(cat "$d/summary.md")" "-f product="

# The fetch has not landed: the hand command is the fetch, then the move.
d=$SCRATCH/move-unfetched
make_move_repo "$d"
rm "$d/data/builds/$NEW_BUILD/simdb.bin"
activate "$d" "$NEW_BUILD" --no-commit --no-gate
assert "unfetched: exits 1" equal "$STATUS" 1
assert "unfetched: names the missing table" contains "$SUMMARY" "data/builds/$NEW_BUILD/simdb.bin is missing"
assert "unfetched: the hand command fetches first" contains "$SUMMARY" "gh workflow run data.yml -f product=wowf -f build=$NEW_BUILD, then gh workflow run client.yml"
assert "unfetched: the active build was not touched" contains "$(cat "$d/web/src/data/active-build.json")" "$OLD_BUILD"

# Argument errors and the already-active build are not failures to report.
d=$SCRATCH/move-args
make_move_repo "$d"
activate "$d" "$NEW_BUILD" --no-gate
assert "--no-gate without --no-commit is a usage error" equal "$STATUS" 2
assert "and writes no failure summary" equal "$SUMMARY" ""
activate "$d" "$OLD_BUILD"
assert "the active build is nothing to do" equal "$STATUS" 0

# The workflows and activate-build.sh run these directly, and this repository has
# core.fileMode off, so a script committed without its executable bit fails only on the
# runner ("Permission denied", which the engine check would read as an unreadable stamp).
for script in tools/activate-build.sh tools/activation-gate.sh tools/engine/pinned-build.sh \
	tools/engine/pin-conditions.sh tools/engine/follow-inputs-hash.sh tools/engine/push-review-branch.sh; do
	mode=$(git -C "$HERE/../.." ls-files -s "$script" | cut -c1-6)
	assert "$script is committed executable (mode $mode)" equal "$mode" 100755
done

if [[ $FAILED -gt 0 ]]; then
	echo "$FAILED failed, $PASSED passed" >&2
	exit 1
fi
echo "ok $PASSED tests"
