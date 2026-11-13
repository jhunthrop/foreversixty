#!/usr/bin/env bash
# Tests for tools/bis-land.sh, bis.yml's landing of a finished ranking onto the branch as
# it is now. Plain bash, no bats, the shape of tools/tests/test_remeasure.sh: each case
# builds a scratch origin, a job checkout at the ranking's base, moves the origin on, and
# runs the real script with the render stubbed (it concatenates the picks into the derived
# files, so a stale render is visible).
#
#   bash tools/tests/test_bis_land.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly SCRIPT=$HERE/../bis-land.sh
readonly BUILD=1.60.1.70938
readonly NEXT_BUILD=1.60.1.71000
readonly PIN=aaaaaaaaa
readonly NEXT_PIN=bbbbbbbbb
# shellcheck disable=SC2016 # expanded by the script's bash -c, with BUILD exported
readonly RENDER='cat data/builds/$BUILD/bis/*.json >data/builds/$BUILD/addon-data.json; cp data/builds/$BUILD/addon-data.json addon/ForeverSixty/Data.lua'

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

g() { git -c user.name=t -c user.email=t@t -c init.defaultBranch=main "$@"; }

# write_tree <dir> <build> <pin>: the files the script reads and writes.
write_tree() {
	local d=$1
	mkdir -p "$d/web/src/data" "$d/sim/enginever" "$d/addon/ForeverSixty" "$d/data/builds/$BUILD/bis" "$d/tools"
	printf '{\n  "build": "%s"\n}\n' "$2" >"$d/web/src/data/active-build.json"
	printf 'package enginever\n\nconst Version = "%s"\n' "$3" >"$d/sim/enginever/version.go"
	cp "$SCRIPT" "$HERE/../remeasure-dps-per-point.sh" "$d/tools/"
}

# scaffold <root>: an origin with the ranking's base, a job checkout of it, and tonight's
# picks; then origin moves on with a hand edit of Data.lua and the old picks file.
scaffold() {
	local root=$1 seed=$1/seed
	mkdir -p "$root"
	g init --quiet --bare "$root/origin.git"
	g init --quiet "$seed"
	write_tree "$seed" "$BUILD" "$PIN"
	echo '{"pick":"old"}' >"$seed/data/builds/$BUILD/bis/rogue-combat.json"
	echo 'old' >"$seed/data/builds/$BUILD/bis/rogue-combat.md"
	echo 'old' >"$seed/data/builds/$BUILD/addon-data.json"
	echo 'old' >"$seed/addon/ForeverSixty/Data.lua"
	echo 'carried' >"$seed/data/builds/$BUILD/dps-per-point.json"
	echo '1.60.1.70790' >"$seed/data/builds/$BUILD/dps-per-point.carried-from"
	printf '{"files": {"dps-per-point.json": "old", "dps-per-point.carried-from": "x"}}\n' >"$seed/data/builds/$BUILD/manifest.json"
	g -C "$seed" add -A && g -C "$seed" commit --quiet -m base
	g -C "$seed" remote add origin "$root/origin.git"
	g -C "$seed" push --quiet origin HEAD:refs/heads/main
	RANKED_HEAD=$(g -C "$seed" rev-parse HEAD)
	g clone --quiet "$root/origin.git" "$root/job"
	g -C "$root/job" checkout --quiet --detach "$RANKED_HEAD"
	mkdir -p "$root/picks"
	echo '{"pick":"tonight"}' >"$root/picks/rogue-combat.json"
	echo 'tonight' >"$root/picks/rogue-combat.md"
	echo 'measured' >"$root/figures.json"
	# main moves while the legs rank: the derived file the old rebase conflicted on.
	echo 'someone else' >"$seed/addon/ForeverSixty/Data.lua"
	echo 'other' >"$seed/other.txt"
	g -C "$seed" add -A && g -C "$seed" commit --quiet -m 'main moved'
	g -C "$seed" push --quiet origin HEAD:refs/heads/main
	MOVED_HEAD=$(g -C "$seed" rev-parse HEAD)
}

# move_origin <root> <file> <content> [<message>]: one more commit on origin's main.
move_origin() {
	local seed=$1/seed
	g -C "$seed" pull --quiet --rebase origin main
	mkdir -p "$(dirname "$seed/$2")"
	printf '%s\n' "$3" >"$seed/$2"
	g -C "$seed" add -A && g -C "$seed" commit --quiet -m "${4:-move}"
	g -C "$seed" push --quiet origin HEAD:refs/heads/main
}

# run_land <root> [VAR=value ...]: runs the script from the job checkout; sets OUT, STATUS.
run_land() {
	local root=$1
	shift
	STATUS=0
	OUT=$(cd "$root/job" && env GIT_AUTHOR_NAME=t GIT_AUTHOR_EMAIL=t@t GIT_COMMITTER_NAME=t GIT_COMMITTER_EMAIL=t@t \
		BRANCH=main PUBLISH=true RANKED_HEAD="$RANKED_HEAD" LAND_RENDER="$RENDER" LAND_SLEEP=0 "$@" \
		tools/bis-land.sh "$BUILD" "$PIN" "$root/picks" "Inputs: 0123456789abcdef" 2>&1) || STATUS=$?
}

origin_head() { git -C "$1/origin.git" rev-parse main; }
origin_file() { git -C "$1/origin.git" show "main:$2"; }
origin_msg() { git -C "$1/origin.git" log -1 --format=%B main; }

# main moved under the ranking: the picks land on the new head, the derived files are
# rendered there (no conflict, no stale render), and the pin is the one ranked on.
d=$SCRATCH/moved
scaffold "$d"
run_land "$d" GITHUB_OUTPUT="$d/output"
assert "moved: tells the caller a commit landed" equal "$(cat "$d/output")" landed=true
assert "moved: exits 0" equal "$STATUS" 0
assert "moved: the commit's parent is main's new head" equal "$(git -C "$d/origin.git" rev-parse main^)" "$MOVED_HEAD"
assert "moved: the picks landed" equal "$(origin_file "$d" "data/builds/$BUILD/bis/rogue-combat.json")" '{"pick":"tonight"}'
assert "moved: Data.lua is rendered from the new picks" equal "$(origin_file "$d" addon/ForeverSixty/Data.lua)" '{"pick":"tonight"}'
assert "moved: main's other change is kept" equal "$(origin_file "$d" other.txt)" other
assert "moved: the subject names the build and pin" contains "$(origin_msg "$d")" "chore(bis): nightly leveling BiS for $BUILD @ $PIN"
assert "moved: the trailer is kept" contains "$(origin_msg "$d")" "Inputs: 0123456789abcdef"
assert "moved: an unmoved pin adds no note" lacks "$(origin_msg "$d")" "Ranked on engine"
assert "moved: no marker without one in the picks" lacks "$(git -C "$d/origin.git" ls-tree -r --name-only main)" ".inputs"
assert "moved: names how far main moved" contains "$OUT" "1 commits ago"

# Run again on the landed tree: nothing to commit, nothing pushed.
before=$(origin_head "$d")
run_land "$d"
assert "again: exits 0" equal "$STATUS" 0
assert "again: says nothing to commit" contains "$OUT" "nothing to commit"
assert "again: pushes nothing" equal "$(origin_head "$d")" "$before"

# A spec's inputs marker lands with its picks; a picks dir without markers lands as before.
d=$SCRATCH/inputs
scaffold "$d"
echo 0123456789abcdef >"$d/picks/rogue-combat.inputs"
run_land "$d"
assert "inputs: exits 0" equal "$STATUS" 0
assert "inputs: the marker landed" equal "$(origin_file "$d" "data/builds/$BUILD/bis/rogue-combat.inputs")" 0123456789abcdef

# The engine pin moved during the ranking: the picks still land, the body says on which pin.
d=$SCRATCH/pin
scaffold "$d"
move_origin "$d" sim/enginever/version.go "$(printf 'package enginever\n\nconst Version = "%s"' "$NEXT_PIN")" 'pin'
run_land "$d"
assert "pin: exits 0" equal "$STATUS" 0
assert "pin: the picks landed" equal "$(origin_file "$d" "data/builds/$BUILD/bis/rogue-combat.json")" '{"pick":"tonight"}'
assert "pin: the subject keeps the ranked pin" contains "$(origin_msg "$d")" "@ $PIN"
assert "pin: the body names both pins" contains "$(origin_msg "$d")" "Ranked on engine $PIN; main pins $NEXT_PIN"
assert "pin: a notice says so" contains "$OUT" "::notice title=engine pin moved"

# The active build moved during the ranking: nothing is committed.
d=$SCRATCH/build
scaffold "$d"
move_origin "$d" web/src/data/active-build.json "{\"build\": \"$NEXT_BUILD\"}" 'activate'
before=$(origin_head "$d")
run_land "$d"
assert "build: exits 0" equal "$STATUS" 0
assert "build: warns that nothing is committed" contains "$OUT" "::warning title=picks not committed::main's active build is $NEXT_BUILD now"
assert "build: pushes nothing" equal "$(origin_head "$d")" "$before"

# Off main: the commit is built on the fresh head and checked with a dry-run push only.
d=$SCRATCH/dry
scaffold "$d"
before=$(origin_head "$d")
run_land "$d" PUBLISH=false GITHUB_OUTPUT="$d/output"
assert "dry: tells the caller nothing landed" equal "$(cat "$d/output")" landed=false
assert "dry: exits 0" equal "$STATUS" 0
assert "dry: says it would land" contains "$OUT" "dry run, would land:"
assert "dry: shows the commit on the fresh head" contains "$OUT" "(parent $MOVED_HEAD)"
assert "dry: pushes nothing" equal "$(origin_head "$d")" "$before"

# A push rejected because main moved again mid-landing: the next round fetches and lands.
d=$SCRATCH/race
scaffold "$d"
race="if [ ! -e $d/raced ]; then touch $d/raced; (cd $d/seed && git -c user.name=t -c user.email=t@t pull --quiet --rebase origin main && echo race >race.txt && git add race.txt && git -c user.name=t -c user.email=t@t commit --quiet -m race && git push --quiet origin HEAD:refs/heads/main); fi; $RENDER"
run_land "$d" LAND_RENDER="$race"
assert "race: exits 0" equal "$STATUS" 0
assert "race: the first push was rejected" contains "$OUT" "attempt 1: the push was rejected"
assert "race: the second round landed" contains "$OUT" "attempt 2: landing on main"
assert "race: the racing commit is kept" equal "$(origin_file "$d" race.txt)" race
assert "race: the picks landed" equal "$(origin_file "$d" "data/builds/$BUILD/bis/rogue-combat.json")" '{"pick":"tonight"}'

# Every round rejected: exits 1 with an error, nothing pushed.
d=$SCRATCH/reject
scaffold "$d"
before=$(origin_head "$d")
reject="(cd $d/seed && git -c user.name=t -c user.email=t@t pull --quiet --rebase origin main && date +%s%N >>race.txt && git add race.txt && git -c user.name=t -c user.email=t@t commit --quiet -m race && git push --quiet origin HEAD:refs/heads/main); $RENDER"
run_land "$d" LAND_RENDER="$reject" LAND_ATTEMPTS=2
assert "reject: exits 1" equal "$STATUS" 1
assert "reject: names the failure" contains "$OUT" "::error title=picks not pushed::2 attempts"
assert "reject: no picks landed" equal "$(origin_file "$d" "data/builds/$BUILD/bis/rogue-combat.json")" '{"pick":"old"}'

# Re-measured figures: they land with the marker deleted and the manifest re-hashed.
d=$SCRATCH/figures
scaffold "$d"
run_land "$d" FIGURES="$d/figures.json"
assert "figures: exits 0" equal "$STATUS" 0
assert "figures: the figures landed" equal "$(origin_file "$d" "data/builds/$BUILD/dps-per-point.json")" measured
assert "figures: the marker is deleted" equal "$(git -C "$d/origin.git" ls-tree main "data/builds/$BUILD/dps-per-point.carried-from")" ""
assert "figures: the manifest drops the marker" lacks "$(origin_file "$d" "data/builds/$BUILD/manifest.json")" carried-from

# A failed render stops the script: nothing is committed or pushed.
d=$SCRATCH/broken
scaffold "$d"
before=$(origin_head "$d")
run_land "$d" LAND_RENDER='exit 3'
assert "broken: fails" test "$STATUS" -ne 0
assert "broken: pushes nothing" equal "$(origin_head "$d")" "$before"

# LAND_SUBJECT replaces the subject (the re-measure's landing).
d=$SCRATCH/subject
scaffold "$d"
run_land "$d" LAND_SUBJECT="chore(bis): re-measured dps-per-point for $BUILD"
assert "subject: exits 0" equal "$STATUS" 0
assert "subject: the given subject" contains "$(origin_msg "$d")" "chore(bis): re-measured dps-per-point for $BUILD"
assert "subject: not the nightly's" lacks "$(origin_msg "$d")" "nightly leveling BiS"

# Usage errors.
d=$SCRATCH/usage
scaffold "$d"
STATUS=0
OUT=$(cd "$d/job" && tools/bis-land.sh "$BUILD" "$PIN" 2>&1) || STATUS=$?
assert "usage: too few arguments exit 2" equal "$STATUS" 2
STATUS=0
OUT=$(cd "$d/job" && tools/bis-land.sh "$BUILD" "$PIN" "$d/picks" t 2>&1) || STATUS=$?
assert "usage: no BRANCH exits 2" equal "$STATUS" 2
mkdir -p "$d/empty"
STATUS=0
OUT=$(cd "$d/job" && BRANCH=main tools/bis-land.sh "$BUILD" "$PIN" "$d/empty" t 2>&1) || STATUS=$?
assert "usage: an empty picks dir exits 2" equal "$STATUS" 2

# bis.yml runs the script directly and this repository has core.fileMode off, so a script
# committed without its executable bit fails only on the runner ("Permission denied", the
# first proof run of this script).
mode=$(git -C "$HERE/../.." ls-files -s tools/bis-land.sh | cut -c1-6)
assert "tools/bis-land.sh is committed executable (mode $mode)" equal "$mode" 100755

if ((FAILED > 0)); then
	echo "$FAILED of $((PASSED + FAILED)) failed" >&2
	exit 1
fi
echo "ok $PASSED tests"
