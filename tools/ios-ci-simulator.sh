#!/usr/bin/env bash
# tools/ios-ci-simulator.sh: make, boot and select CI's iOS simulator (ios.yml).
#
#   tools/ios-ci-simulator.sh create VERSION BUILD       pick the runtime (e.g. 26.5 23F77) and make
#                                                        the job's own iPhone 17 on it, not booted
#   tools/ios-ci-simulator.sh boot-start UDID            boot UDID in the background, so the job
#                                                        builds while the simulator boots
#   tools/ios-ci-simulator.sh boot-wait                  wait for that boot, print its log, fail if
#                                                        it failed
#   tools/ios-ci-simulator.sh ensure VERSION BUILD UDID  keep UDID if it is still available and
#                                                        booted, else log why, make and boot a new one
#
# create and ensure print the xcodebuild destination (platform=iOS Simulator,id=<udid>) on stdout
# and, on Actions, write SIMULATOR and SIMULATOR_UDID to GITHUB_ENV.
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
#   4. exports the destination.
# A first boot is mostly data migration (8 minutes on a downloaded runtime, run 42842739728), so
# ios.yml starts it before the build and waits for it only before testing. A booted device has
# still lost its runtime minutes later ("Unable to find a device matching the provided destination
# specifier", run 42515970747), so ios.yml ensures again right before testing.
# It fails with an ::error:: annotation when no iOS <major>.x runtime can be had.
set -euo pipefail

readonly DEVICE_NAME="Forever Sixty CI (iPhone 17)"
readonly DEVICE_TYPE="com.apple.CoreSimulator.SimDeviceType.iPhone-17"
readonly BOOT_ATTEMPTS=3
# A fresh runtime's first boot took 6m18s on the xcode-27 image (run 42512503765).
readonly BOOT_TIMEOUT_SECONDS=600
readonly RUNTIME_READY_TIMEOUT_SECONDS=180
readonly DOWNLOAD_ATTEMPTS=2
readonly POLL_SECONDS=5
# Where boot-start leaves the background boot's log and exit status for boot-wait.
readonly BOOT_STATE="${RUNNER_TEMP:-${TMPDIR:-/tmp}}/ios-ci-simulator-boot"

fail() { echo "::error title=iOS simulator::$*" >&2; exit 1; }
log() { echo "ios-ci-simulator: $*" >&2; }
# Annotations go to stderr too: stdout carries the values the functions return.
warn() { echo "::warning title=iOS simulator::$*" >&2; }

usage() { fail "usage: $0 create VERSION BUILD | boot-start UDID | boot-wait | ensure VERSION BUILD UDID (e.g. 26.5 23F77)"; }

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
	if [ -z "$runtime" ] && [ -n "$(find_runtime any)" ]; then
		log "iOS $VERSION ($BUILD) is listed but not available yet; waiting for it to mount"
		runtime=$(wait_for_runtime || true)
	fi
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
	local deadline=$((SECONDS + $1)) pid
	shift
	"$@" &
	pid=$!
	while kill -0 "$pid" 2>/dev/null; do
		if [ "$SECONDS" -ge "$deadline" ]; then
			kill "$pid" 2>/dev/null || true
			wait "$pid" 2>/dev/null || true
			return 124
		fi
		sleep 1
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
		xcrun simctl shutdown "$udid" >/dev/null 2>&1 || true
		xcrun simctl erase "$udid" >&2 || true
		sleep "$POLL_SECONDS"
	done
	return 1
}

# True when the device exists, is available and is booted.
device_usable() {
	xcrun simctl list devices -j | python3 -c '
import json, sys
udid = sys.argv[1]
for devices in json.load(sys.stdin)["devices"].values():
    for d in devices:
        if d["udid"] == udid and d["isAvailable"] and d["state"] == "Booted":
            sys.exit(0)
sys.exit(1)' "$1"
}

diagnose() {
	log "simulator state:"
	{ xcrun simctl list runtimes; xcrun simctl runtime list; xcrun simctl list devices available; } >&2 || true
}

create_device() {
	local runtime udid
	runtime=$(pick_runtime)
	# Installing a runtime also creates default devices, so a name could match twice: the job makes
	# its own device and tests on it by id.
	udid=$(xcrun simctl create "$DEVICE_NAME" "$DEVICE_TYPE" "$runtime") \
		|| fail "could not create $DEVICE_TYPE on $runtime"
	log "created $DEVICE_NAME: $udid"
	echo "$udid"
}

boot_or_fail() {
	boot_device "$1" || { diagnose; fail "simulator $1 did not boot after $BOOT_ATTEMPTS attempts"; }
}

# Prints the destination for UDID and, on Actions, exports it.
export_destination() {
	local destination="platform=iOS Simulator,id=$1"
	if [ -n "${GITHUB_ENV:-}" ]; then
		{ echo "SIMULATOR=$destination"; echo "SIMULATOR_UDID=$1"; } >>"$GITHUB_ENV"
	fi
	log "destination: $destination"
	echo "$destination"
}

# The boot outlives the step that starts it (Actions ends a job's leftover processes only when the
# job completes); boot-wait reads its status file, which appears when the boot is over.
boot_start() {
	rm -f "$BOOT_STATE.status"
	# The background boot must not write to this step's GITHUB_ENV after the step has ended. The
	# inner script expands its own positional parameters, hence the single quotes.
	# shellcheck disable=SC2016
	nohup env -u GITHUB_ENV bash -c 'bash "$0" boot "$1" >"$2.log" 2>&1; echo $? >"$2.status"' \
		"$0" "$1" "$BOOT_STATE" >/dev/null 2>&1 &
	log "booting $1 in the background (log: $BOOT_STATE.log)"
}

boot_wait() {
	local deadline=$((SECONDS + BOOT_ATTEMPTS * (BOOT_TIMEOUT_SECONDS + 60))) status
	while [ ! -s "$BOOT_STATE.status" ]; do
		if [ "$SECONDS" -ge "$deadline" ]; then
			cat "$BOOT_STATE.log" >&2 || true
			fail "the background boot did not finish"
		fi
		sleep "$POLL_SECONDS"
	done
	cat "$BOOT_STATE.log" >&2
	status=$(cat "$BOOT_STATE.status")
	[ "$status" -eq 0 ] || fail "the background boot failed (exit $status)"
}

ensure() {
	local udid="$1"
	if device_usable "$udid"; then
		log "$udid is still available and booted"
	else
		warn "simulator $udid is no longer available and booted; setting up a new one"
		diagnose
		udid=$(create_device)
		boot_or_fail "$udid"
	fi
	export_destination "$udid"
}

main() {
	case "${1:-}" in
		create)
			[ $# -eq 3 ] || usage
			VERSION="$2" BUILD="$3" MAJOR="${2%%.*}"
			export_destination "$(create_device)"
			;;
		boot) [ $# -eq 2 ] || usage; boot_or_fail "$2" ;;
		boot-start) [ $# -eq 2 ] || usage; boot_start "$2" ;;
		boot-wait) [ $# -eq 1 ] || usage; boot_wait ;;
		ensure)
			[ $# -eq 4 ] || usage
			VERSION="$2" BUILD="$3" MAJOR="${2%%.*}"
			ensure "$4"
			;;
		*) usage ;;
	esac
}

main "$@"
