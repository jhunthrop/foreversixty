#!/usr/bin/env bash
# tools/mobile-device.sh: the ONE iOS simulator and the ONE Android emulator that every local
# lane shares, plus a lock so only one lane drives each at a time.
#
# Owner 2026-10-30: "can we have them use one emulator for all of the lane testing?" Parallel
# lanes each booting their own device collided, crashed and stacked macOS crash dialogs.
#
#   tools/mobile-device.sh ios-udid          print the shared simulator's UDID (creating it once)
#   tools/mobile-device.sh ios-run CMD...    run CMD holding the iOS lock; FS_SIM_UDID is exported
#   tools/mobile-device.sh android-run CMD...  boot the shared AVD if needed, run CMD holding the
#                                            Android lock (screenshots only: unit and Roborazzi
#                                            tests run on the JVM and need no emulator)
#
# A lane waiting on a lock simply blocks until the lane ahead of it finishes; nothing fails.
# CI never uses this script: its runners are single-tenant.
set -euo pipefail

IOS_NAME="Forever Sixty Test (iPhone 17)"
IOS_TYPE="com.apple.CoreSimulator.SimDeviceType.iPhone-17"
AVD_NAME="foreversixty-test"
AVD_DEVICE="pixel_9"
LOCK_DIR="${TMPDIR:-/tmp}"
IOS_LOCK="$LOCK_DIR/foreversixty-ios-simulator.lock"
ANDROID_LOCK="$LOCK_DIR/foreversixty-android-emulator.lock"
SDK="${ANDROID_HOME:-$HOME/Library/Android/sdk}"

die() { echo "mobile-device: $*" >&2; exit 1; }

# Every installed iOS runtime, newest first.
ios_runtimes() {
	xcrun simctl list runtimes -j | python3 -c '
import json, sys
rs = [r for r in json.load(sys.stdin)["runtimes"] if r["platform"] == "iOS" and r["isAvailable"]]
for r in sorted(rs, key=lambda r: [int(p) for p in r["version"].split(".")], reverse=True):
    print(r["identifier"])'
}

ios_udid() {
	local udid
	udid=$(xcrun simctl list devices -j | python3 -c '
import json, sys
name = sys.argv[1]
for devices in json.load(sys.stdin)["devices"].values():
    for d in devices:
        if d["name"] == name and d["isAvailable"]:
            print(d["udid"]); sys.exit(0)' "$IOS_NAME")
	if [ -z "$udid" ]; then
		# The newest runtime that accepts the device type (iOS 27.1 refuses iPhone 17 on 2026-10-30).
		local runtime
		for runtime in $(ios_runtimes); do
			udid=$(xcrun simctl create "$IOS_NAME" "$IOS_TYPE" "$runtime" 2>/dev/null) && break
		done
		[ -n "$udid" ] || die "no installed iOS runtime accepts $IOS_TYPE"
	fi
	echo "$udid"
}

ios_run() {
	[ $# -gt 0 ] || die "ios-run needs a command"
	local udid
	udid=$(ios_udid)
	export FS_SIM_UDID="$udid"
	# lockf -k keeps the lock file; the lock is released when the command exits, however it exits.
	exec lockf -k "$IOS_LOCK" "$@"
}

android_boot() {
	local adb="$SDK/platform-tools/adb" emulator="$SDK/emulator/emulator" avdmanager
	avdmanager=$(command -v avdmanager || echo "$SDK/cmdline-tools/latest/bin/avdmanager")
	if ! "$emulator" -list-avds | grep -qx "$AVD_NAME"; then
		local image api pkg
		image=$(find "$SDK"/system-images -maxdepth 3 -path '*/google_apis/arm64-v8a' -type d 2>/dev/null | sort -V | tail -1)
		[ -n "$image" ] || die "no arm64 google_apis system image under $SDK/system-images"
		api=$(basename "$(dirname "$(dirname "$image")")")
		pkg="system-images;$api;google_apis;arm64-v8a"
		echo no | "$avdmanager" create avd -n "$AVD_NAME" -k "$pkg" -d "$AVD_DEVICE" >/dev/null
	fi
	if ! "$adb" devices | grep -q "^emulator-"; then
		nohup "$emulator" -avd "$AVD_NAME" -no-window -no-audio -no-boot-anim >/dev/null 2>&1 &
		"$adb" wait-for-device
		until [ "$("$adb" shell getprop sys.boot_completed 2>/dev/null | tr -d '\r')" = 1 ]; do sleep 2; done
	fi
}

android_run() {
	[ $# -gt 0 ] || die "android-run needs a command"
	exec lockf -k "$ANDROID_LOCK" bash -c "$(declare -f android_boot die); SDK='$SDK' AVD_NAME='$AVD_NAME' AVD_DEVICE='$AVD_DEVICE' android_boot && exec \"\$@\"" _ "$@"
}

case "${1:-}" in
ios-udid) ios_udid ;;
ios-run) shift; ios_run "$@" ;;
android-run) shift; android_run "$@" ;;
*) die "usage: $0 ios-udid | ios-run CMD... | android-run CMD..." ;;
esac
