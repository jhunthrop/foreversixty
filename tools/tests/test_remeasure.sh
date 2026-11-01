#!/usr/bin/env bash
# Tests for tools/remeasure-dps-per-point.sh, bis.yml's re-measure of a carried
# dps-per-point.json. Plain bash, no bats, the shape of tools/tests/test_activation.sh:
# each case runs the real script in a scratch repository root with DPP_MEASURE stubbed to
# write new figures.
#
#   bash tools/tests/test_remeasure.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly SCRIPT=$HERE/../remeasure-dps-per-point.sh
readonly BUILD=1.60.1.70707
readonly OLD_BUILD=1.60.1.70658
readonly NEW_FIGURES='{"measured":"tonight"}'

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

# scaffold <root> <with marker: yes|no>: a build directory with carried figures, an
# optional marker and a manifest listing every file, and a stub measure that rewrites the
# figures and records that it ran.
scaffold() {
	local root=$1 marker=$2 dir=$1/data/builds/$BUILD
	mkdir -p "$dir"
	echo '{"measured":"carried"}' >"$dir/dps-per-point.json"
	echo '{}' >"$dir/items.json"
	local entries='"dps-per-point.json": "old", "items.json": "keep"'
	if [ "$marker" = yes ]; then
		echo "$OLD_BUILD" >"$dir/dps-per-point.carried-from"
		entries="$entries, \"dps-per-point.carried-from\": \"marker\""
	fi
	printf '{"build": "%s", "files": {%s}}\n' "$BUILD" "$entries" | jq --indent 2 . >"$dir/manifest.json"
	cat >"$root/measure" <<EOF
#!/usr/bin/env bash
echo ran >>"$root/measure-calls"
printf '%s\n' '$NEW_FIGURES' >"$dir/dps-per-point.json"
EOF
	chmod +x "$root/measure"
}

# run_script <root> [<exit of the measure>]: runs the script from <root>; sets OUT, STATUS.
run_script() {
	local root=$1
	STATUS=0
	OUT=$(cd "$root" && DPP_MEASURE="$root/measure" GITHUB_OUTPUT="$root/output" "$SCRIPT" "$BUILD" 2>&1) || STATUS=$?
}

files_of() { jq -r '.files | keys | join(",")' "$1/data/builds/$BUILD/manifest.json"; }
digest_of() { jq -r --arg f "$2" '.files[$f]' "$1/data/builds/$BUILD/manifest.json"; }

# A carried build: measured, marker deleted, the manifest drops it and re-hashes the figures.
d=$SCRATCH/carried
scaffold "$d" yes
run_script "$d"
dir=$d/data/builds/$BUILD
assert "carried: exits 0" equal "$STATUS" 0
assert "carried: names the build it was carried from" contains "$OUT" "carried from $OLD_BUILD"
assert "carried: measures" equal "$(cat "$d/measure-calls")" ran
assert "carried: deletes the marker" test ! -e "$dir/dps-per-point.carried-from"
assert "carried: the manifest no longer lists the marker" equal "$(files_of "$d")" "dps-per-point.json,items.json"
assert "carried: the manifest holds the re-measured figures' digest" \
	equal "$(digest_of "$d" dps-per-point.json)" "$(sha256sum "$dir/dps-per-point.json" | cut -d' ' -f1)"
assert "carried: other entries are kept" equal "$(digest_of "$d" items.json)" keep
assert "carried: the manifest keeps its other fields" equal "$(jq -r .build "$dir/manifest.json")" "$BUILD"
assert "carried: reports measured=true" equal "$(cat "$d/output")" "measured=true"

# A build's own figures: no measure, nothing changes.
d=$SCRATCH/own
scaffold "$d" no
before=$(cat "$d/data/builds/$BUILD/manifest.json")
run_script "$d"
assert "own: exits 0" equal "$STATUS" 0
assert "own: says the figures are kept" contains "$OUT" "kept"
assert "own: does not measure" test ! -e "$d/measure-calls"
assert "own: the manifest is unchanged" equal "$(cat "$d/data/builds/$BUILD/manifest.json")" "$before"
assert "own: reports nothing" test ! -e "$d/output"

# A failed measure keeps the marker and the manifest, and fails the step.
d=$SCRATCH/red
scaffold "$d" yes
printf '#!/usr/bin/env bash\nexit 1\n' >"$d/measure"
before=$(cat "$d/data/builds/$BUILD/manifest.json")
run_script "$d"
assert "red: exits non-zero" test "$STATUS" -ne 0
assert "red: keeps the marker" test -e "$d/data/builds/$BUILD/dps-per-point.carried-from"
assert "red: the manifest is unchanged" equal "$(cat "$d/data/builds/$BUILD/manifest.json")" "$before"
assert "red: reports nothing" test ! -e "$d/output"

# A build that does not exist is a usage error.
d=$SCRATCH/none
mkdir -p "$d"
STATUS=0
OUT=$(cd "$d" && "$SCRIPT" "$BUILD" 2>&1) || STATUS=$?
assert "no build: exits 2" equal "$STATUS" 2

if [[ $FAILED -gt 0 ]]; then
	echo "$FAILED failed, $PASSED passed" >&2
	exit 1
fi
echo "ok $PASSED tests"
