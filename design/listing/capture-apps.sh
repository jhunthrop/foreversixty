#!/usr/bin/env bash
# design/listing/capture-apps.sh: the store screenshots' raw captures from the fixture apps on the
# shared simulator and emulator (tools/mobile-device.sh holds each lock), in dark appearance, with a
# clean status bar (9:41, full bars, full battery) and the system font at its default size.
# design/listing/build-store.mjs then draws each capture onto its store plate.
#
#   design/listing/capture-apps.sh ios        # apps/ios UITests/StoreCaptureTests
#   design/listing/capture-apps.sh android    # after ./gradlew :app:assembleDebug in apps/android
#   design/listing/capture-apps.sh            # both, one after the other
#
# Writes design/listing/app-captures/{ios,android}/{home,gear,raid,raid-roster}.png.
# Run from the repo root.
set -euo pipefail

ROOT=$(cd "$(dirname "$0")/../.." && pwd)
OUT="$ROOT/design/listing/app-captures"
DEVICE="$ROOT/tools/mobile-device.sh"
RAID_LINK="https://foreversixty.gg/app/raid/2/31"
SETTLES=6

ios() {
	mkdir -p "$OUT/ios"
	(cd "$ROOT/apps/ios" && make project >/dev/null)
	"$DEVICE" ios-run bash -c '
		set -euo pipefail
		xcrun simctl boot "$FS_SIM_UDID" 2>/dev/null || true
		xcrun simctl ui "$FS_SIM_UDID" appearance dark
		xcrun simctl status_bar "$FS_SIM_UDID" override --time 9:41 --dataNetwork wifi --wifiMode active \
			--wifiBars 3 --cellularMode active --cellularBars 4 --batteryState charged --batteryLevel 100
		trap "xcrun simctl status_bar \"\$FS_SIM_UDID\" clear" EXIT
		cd "$1/apps/ios"
		xcodebuild -project ForeverSixty.xcodeproj -scheme ForeverSixtyCaptures \
			-destination "platform=iOS Simulator,id=$FS_SIM_UDID" -derivedDataPath build/DerivedData \
			-only-testing:ForeverSixtyUITests/StoreCaptureTests test
	' _ "$ROOT"
}

android() {
	mkdir -p "$OUT/android"
	"$DEVICE" android-run bash "$ROOT/design/listing/capture-android.sh" "$ROOT" "$OUT/android" "$RAID_LINK" "$SETTLES"
}

case "${1:-all}" in
ios) ios ;;
android) android ;;
all) ios; android ;;
*) echo "usage: $0 [ios|android]" >&2; exit 2 ;;
esac
