#!/usr/bin/env bash
# tools/ios-ci-simulator.sh: make, boot and select CI's iOS simulator (ios.yml).
#
#   tools/ios-ci-simulator.sh VERSION BUILD     e.g. 26.5 23F77
#
# The snapshot goldens are rendered on one runtime (ios.yml, IOS_RUNTIME_*), so the script
#   1. picks that runtime by version AND build from `simctl list runtimes -j`, downloading it
#      (`xcodebuild -downloadPlatform iOS -buildVersion`) when the image lacks it, else falls back
#      to the newest installed iOS <major>.x runtime with a warning (goldens may then differ);
#   2. waits for the runtime to report available (a fresh download mounts in the background, and
#      testing on it too early lost the runtime mid-run: "cannot determine the runtime bundle",
#      "iOS Simulator (null)", run 42503034639);
#   3. creates its own iPhone 17 on that runtime identifier and boots it, up to BOOT_ATTEMPTS
#      times with a shutdown and erase between attempts, waiting on `simctl bootstatus -b`;
#   4. exports SIMULATOR=platform=iOS Simulator,id=<udid> to GITHUB_ENV (stdout without it).
# It fails with an ::error:: annotation when no iOS <major>.x runtime can be had.
set -euo pipefail

readonly DEVICE_NAME="Forever Sixty CI (iPhone 17)"
readonly DEVICE_TYPE="com.apple.CoreSimulator.SimDeviceType.iPhone-17"
readonly BOOT_ATTEMPTS=3
readonly BOOT_TIMEOUT_SECONDS=300
readonly RUNTIME_READY_TIMEOUT_SECONDS=180
readonly DOWNLOAD_ATTEMPTS=2
readonly POLL_SECONDS=5

fail() { echo "::error title=iOS simulator::$*" >&2; exit 1; }
log() { echo "ios-ci-simulator: $*" >&2; }
# Annotations go to stderr too: stdout carries the values the functions return.
warn() { echo "::warning title=iOS simulator::$*" >&2; }

[ $# -eq 2 ] || fail "usage: $0 VERSION BUILD (e.g. 26.5 23F77)"
readonly VERSION="$1" BUILD="$2"
readonly MAJOR="${VERSION%%.*}"

# Prints the identifier of the available iOS runtime matching the query, or nothing.
#   exact        version VERSION and build BUILD
#   any          any build of VERSION, even unavailable (to detect a mount still in progress)
#   major        the newest available MAJOR.x runtime
find_runtime() {
	xcrun simctl list runtimes -j | python3 -c '
import json, sys
mode, version, build, major = sys.argv[1:5]
def key(r): return [int(p) for p in r["version"].split(".")]
rs = [r for r in json.load(sys.stdin)["runtimes"] if r.get("platform") == "iOS"]
if mode == "exact":
    rs = [r for r in rs if r["isAvailable"] and r["version"] == version and r["buildversion"] == build]
elif mode == "any":
    rs = [r for r in rs if r["version"] == version and r["buildversion"] == build]
else:
    rs = [r for r in rs if r["isAvailable"] and r["version"].split(".")[0] == major]
for r in sorted(rs, key=key, reverse=True)[:1]:
    print(r["identifier"])' "$1" "$VERSION" "$BUILD" "$MAJOR"
}

download_runtime() {
	local attempt
	for attempt in $(seq 1 "$DOWNLOAD_ATTEMPTS"); do
		log "downloading iOS $VERSION ($BUILD), attempt $attempt of $DOWNLOAD_ATTEMPTS"
		if xcodebuild -downloadPlatform iOS -buildVersion "$BUILD" >&2; then return 0; fi
		warn "xcodebuild -downloadPlatform iOS -buildVersion $BUILD failed (attempt $attempt)"
	done
	return 1
}

# Waits until the exact runtime is listed as available; a fresh install mounts in the background.
wait_for_runtime() {
	local waited=0 runtime
	while [ "$waited" -lt "$RUNTIME_READY_TIMEOUT_SECONDS" ]; do
		runtime=$(find_runtime exact)
		if [ -n "$runtime" ]; then echo "$runtime"; return 0; fi
		# Not even listed: waiting will not help.
		[ -n "$(find_runtime any)" ] || return 1
		sleep "$POLL_SECONDS"
		waited=$((waited + POLL_SECONDS))
	done
	return 1
}

pick_runtime() {
	local runtime
	runtime=$(find_runtime exact)
	if [ -z "$runtime" ]; then
		download_runtime || log "the iOS $VERSION ($BUILD) download failed"
		runtime=$(wait_for_runtime || true)
	fi
	if [ -n "$runtime" ]; then
		log "using iOS $VERSION ($BUILD): $runtime"
	else
		runtime=$(find_runtime major)
		[ -n "$runtime" ] || fail "no iOS $MAJOR.x simulator runtime is installed and iOS $VERSION ($BUILD) could not be downloaded; installed runtimes: $(xcrun simctl list runtimes | tr '\n' ' ')"
		warn "iOS $VERSION ($BUILD) is unavailable; falling back to $runtime, so snapshot goldens may not match"
	fi
	echo "$runtime"
}

# Runs "$@" for at most $1 seconds; macOS ships no timeout(1).
run_with_timeout() {
	local seconds="$1" pid waited=0
	shift
	"$@" &
	pid=$!
	while kill -0 "$pid" 2>/dev/null; do
		if [ "$waited" -ge "$seconds" ]; then
			kill "$pid" 2>/dev/null || true
			wait "$pid" 2>/dev/null || true
			return 124
		fi
		sleep 1
		waited=$((waited + 1))
	done
	wait "$pid"
}

boot_device() {
	local udid="$1" attempt
	for attempt in $(seq 1 "$BOOT_ATTEMPTS"); do
		log "booting $udid, attempt $attempt of $BOOT_ATTEMPTS"
		if run_with_timeout "$BOOT_TIMEOUT_SECONDS" xcrun simctl bootstatus "$udid" -b >&2; then
			log "booted $udid"
			return 0
		fi
		warn "simulator $udid failed to boot (attempt $attempt of $BOOT_ATTEMPTS); shutting down and erasing"
		xcrun simctl shutdown "$udid" 2>/dev/null || true
		xcrun simctl erase "$udid" || true
		sleep "$POLL_SECONDS"
	done
	return 1
}

main() {
	local runtime udid destination
	runtime=$(pick_runtime)
	# Installing a runtime also creates default devices, so a name could match twice: the job makes
	# its own device and tests on it by id.
	udid=$(xcrun simctl create "$DEVICE_NAME" "$DEVICE_TYPE" "$runtime") \
		|| fail "could not create $DEVICE_TYPE on $runtime"
	log "created $DEVICE_NAME: $udid"
	boot_device "$udid" || fail "simulator $udid ($runtime) did not boot after $BOOT_ATTEMPTS attempts"
	destination="platform=iOS Simulator,id=$udid"
	if [ -n "${GITHUB_ENV:-}" ]; then
		echo "SIMULATOR=$destination" >>"$GITHUB_ENV"
	fi
	echo "SIMULATOR=$destination"
}

main
