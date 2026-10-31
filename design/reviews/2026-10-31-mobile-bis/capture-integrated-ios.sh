#!/usr/bin/env bash
# Integration captures on the shared simulator (run under tools/mobile-device.sh ios-run, which
# exports FS_SIM_UDID): the item tooltip from a healer and a tank list, on the live debug build
# (`xcodebuild ... API_MODE=live`), which reads the real CDN files. simctl cannot tap, so the
# debug-only -FSBisSpec and -FSBisTooltip launch arguments open the list and the row's tooltip.
# Usage: capture-integrated-ios.sh <out dir> <live ForeverSixty.app>
set -euo pipefail
OUT="$1"; APP="$2"
BUNDLE=gg.foreversixty.app
xcrun simctl bootstatus "$FS_SIM_UDID" -b >/dev/null
xcrun simctl install "$FS_SIM_UDID" "$APP"
xcrun simctl ui "$FS_SIM_UDID" content_size large
shoot() { # name, spec, slot
  xcrun simctl terminate "$FS_SIM_UDID" "$BUNDLE" >/dev/null 2>&1 || true
  xcrun simctl launch "$FS_SIM_UDID" "$BUNDLE" -FSTab bis -FSBisSpec "$2" -FSBisTooltip "$3" >/dev/null
  sleep 12
  xcrun simctl io "$FS_SIM_UDID" screenshot "$OUT/integrated-$1-ios.png" >/dev/null
}
shoot tooltip-healer priest-holy shoulder
shoot tooltip-tank warrior-protection back
xcrun simctl terminate "$FS_SIM_UDID" "$BUNDLE" >/dev/null 2>&1 || true
