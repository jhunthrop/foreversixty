#!/usr/bin/env bash
# Tests for tools/publish-feed.sh, which publishes a signed release to the site's release
# feed. Plain bash, no bats, the shape of tools/tests/test_wasm_compress.sh. wrangler is a
# stub that records each upload and keeps a copy of what it was given.
#
#   bash tools/tests/test_publish_feed.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly SCRIPT=$HERE/../publish-feed.sh

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

# The stub: "r2 object put <bucket/key> --file <f> --content-type <t> --remote".
stub=$SCRATCH/wrangler
cat >"$stub" <<'STUB'
#!/usr/bin/env bash
set -eu
key=$4 file=$6 type=$8
echo "$key|$type|${9:-}" >>"$UPLOADS"
mkdir -p "$STORE/$(dirname "$key")"
cp "$file" "$STORE/$key"
STUB
chmod +x "$stub"

run() { # run <args...>: sets OUT and STATUS
	export UPLOADS=$SCRATCH/uploads STORE=$SCRATCH/store
	rm -rf "$UPLOADS" "$STORE"
	STATUS=0
	OUT=$(WRANGLER=$stub FEED_PUBLISHED_AT=2026-11-06T10:00:00Z "$SCRIPT" "$@" 2>&1) || STATUS=$?
}

cd "$SCRATCH" || exit 1
printf 'zip' >ForeverSixty-2.1.0.zip
printf 'sig' >ForeverSixty-2.1.0.zip.minisig

run addon addon-v2.1.0 ForeverSixty-2.1.0.zip ForeverSixty-2.1.0.zip.minisig
assert "a signed release publishes ($OUT)" equal "$STATUS" 0
assert "three uploads: two files, then the manifest last" equal "$(cut -d'|' -f1 "$UPLOADS" | paste -sd, -)" \
	"foreversixty-logs/downloads/addon/addon-v2.1.0/ForeverSixty-2.1.0.zip,foreversixty-logs/downloads/addon/addon-v2.1.0/ForeverSixty-2.1.0.zip.minisig,foreversixty-logs/downloads/addon/latest.json"
assert "every upload is remote" equal "$(cut -d'|' -f3 "$UPLOADS" | sort -u)" "--remote"
assert "the zip is typed as a zip" contains "$(head -1 "$UPLOADS")" "|application/zip|"
manifest=$STORE/foreversixty-logs/downloads/addon/latest.json
assert "the manifest names the family and tag" equal "$(jq -r '"\(.schema) \(.family) \(.tag_name) \(.published_at)"' "$manifest")" \
	"1 addon addon-v2.1.0 2026-11-06T10:00:00Z"
assert "each asset carries its site URL" equal "$(jq -r '.assets[0].browser_download_url' "$manifest")" \
	"https://foreversixty.gg/dl/addon/addon-v2.1.0/ForeverSixty-2.1.0.zip"
assert "each asset carries its size and sha256" equal "$(jq -r '.assets[0] | "\(.size) \(.sha256)"' "$manifest")" \
	"3 $(printf zip | sha256sum | cut -d' ' -f1)"

printf 'dmg' >foreversixty-companion_darwin_universal.dmg
printf 'sig' >foreversixty-companion_darwin_universal.dmg.minisig
run companion companion-v0.3.0 foreversixty-companion_darwin_universal.dmg foreversixty-companion_darwin_universal.dmg.minisig
assert "a Mac disk image publishes ($OUT)" equal "$STATUS" 0
assert "and is typed as a disk image" contains "$(head -1 "$UPLOADS")" "|application/x-apple-diskimage|"

run addon addon-v2.1.0 ForeverSixty-2.1.0.zip
assert "an unsigned file is refused" equal "$STATUS" 64
assert "and says which" contains "$OUT" "ForeverSixty-2.1.0.zip has no ForeverSixty-2.1.0.zip.minisig"
assert "and uploads nothing" test ! -e "$UPLOADS"

run addon data-v2026.11.06 ForeverSixty-2.1.0.zip ForeverSixty-2.1.0.zip.minisig
assert "a tag of another family is refused" equal "$STATUS" 64
run engine engine-v1 ForeverSixty-2.1.0.zip ForeverSixty-2.1.0.zip.minisig
assert "an unknown family is refused" equal "$STATUS" 64
run addon addon-v2.1.0
assert "no files is usage" equal "$STATUS" 64

printf 'key' >MINISIGN_PUBLIC_KEY.txt
run companion companion-v0.2.31 ForeverSixty-2.1.0.zip ForeverSixty-2.1.0.zip.minisig MINISIGN_PUBLIC_KEY.txt
assert "a .txt needs no signature" equal "$STATUS" 0

if ((FAILED > 0)); then
	echo "$FAILED of $((PASSED + FAILED)) assertions failed" >&2
	exit 1
fi
echo "ok $PASSED tests"
