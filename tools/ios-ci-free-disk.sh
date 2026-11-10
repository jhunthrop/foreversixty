#!/usr/bin/env bash
# tools/ios-ci-free-disk.sh: make room on CI's macOS runner before the iOS job (ios.yml).
#
#   tools/ios-ci-free-disk.sh KEEP_RUNTIME_BUILD    e.g. 23F77: the simulator runtime ios.yml pins
#
# Record runs filled the xcode-27 runner's disk mid-test ("No space left on device", runs
# 42908207505, 42912202294 and 42916220914): the image's preinstalled simulator runtimes, other
# Xcodes and toolchains, the 8.5 GB iOS 26.5 download, DerivedData and every failing snapshot's
# attachments in the result bundle. This removes what the job never uses:
#   - every simulator runtime except KEEP_RUNTIME_BUILD (tvOS, watchOS, visionOS and other iOS),
#   - every Xcode except the selected one (setup-xcode runs first),
#   - the Android SDK, .NET, Haskell (ghcup), Rust, Go, PowerShell and the hosted tool cache,
#   - old DerivedData and simulator caches.
# It prints `df -h` before and after, and the space each group freed. Nothing here fails the job:
# a group that is missing is skipped, and a failed removal is a warning.
set -uo pipefail

[ $# -eq 1 ] || { echo "::error::usage: $0 KEEP_RUNTIME_BUILD (e.g. 23F77)"; exit 1; }
readonly KEEP_BUILD="$1"

warn() { echo "::warning title=Free disk::$*"; }
free_gb() { df -g / | awk 'NR==2 {print $4}'; }

# Runs "$@" as a named group and logs how much space it freed and how long it took.
step() {
	local name="$1" before start
	shift
	before=$(free_gb)
	start=$SECONDS
	"$@" || warn "$name: removal failed"
	echo "free-disk: $name freed $(($(free_gb) - before)) GB in $((SECONDS - start))s ($(free_gb) GB free)"
}

remove_paths() {
	local path
	for path in "$@"; do
		[ -e "$path" ] || continue
		sudo rm -rf "$path" || warn "could not remove $path"
	done
}

delete_runtimes() {
	local ids id
	ids=$(xcrun simctl runtime list -j | python3 -c '
import json, sys
keep = sys.argv[1]
for rid, r in json.load(sys.stdin).items():
    if r.get("build") != keep:
        print(rid)' "$KEEP_BUILD") || return 1
	for id in $ids; do
		xcrun simctl runtime delete "$id" >/dev/null 2>&1 || warn "could not delete simulator runtime $id"
	done
}

other_xcodes() {
	local selected app
	selected=$(cd "$(xcode-select -p)/../.." && pwd -P)
	for app in /Applications/Xcode*.app; do
		[ -d "$app" ] && [ ! -L "$app" ] || continue
		[ "$(cd "$app" && pwd -P)" = "$selected" ] || echo "$app"
	done
}

remove_other_xcodes() {
	local apps
	apps=$(other_xcodes)
	# shellcheck disable=SC2086 # one path per line, none with spaces (/Applications/Xcode_*.app)
	[ -z "$apps" ] || remove_paths $apps
}

main() {
	echo "free-disk: before"
	df -h /
	step "simulator runtimes other than $KEEP_BUILD" delete_runtimes
	step "Xcodes other than $(xcode-select -p)" remove_other_xcodes
	step "Android SDK" remove_paths "${ANDROID_HOME:-$HOME/Library/Android/sdk}" "$HOME/Library/Android/sdk"
	step ".NET, Haskell, Rust, Go and PowerShell" remove_paths "$HOME/.dotnet" /usr/local/share/dotnet \
		"$HOME/.ghcup" "$HOME/.cargo" "$HOME/.rustup" /usr/local/go /usr/local/microsoft/powershell
	step "the hosted tool cache" remove_paths "${RUNNER_TOOL_CACHE:-$HOME/hostedtoolcache}"
	step "DerivedData and simulator caches" remove_paths "$HOME/Library/Developer/Xcode/DerivedData" \
		"$HOME/Library/Developer/CoreSimulator/Caches"
	echo "free-disk: after"
	df -h /
}

main
