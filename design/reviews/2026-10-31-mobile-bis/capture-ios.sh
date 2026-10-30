#!/usr/bin/env bash
# iOS captures of the BiS tab on the shared simulator (run under tools/mobile-device.sh ios-run,
# which exports FS_SIM_UDID). The debug build opens on the tab named by -FSTab.
set -euo pipefail
OUT="$1"; APP="$2"
BUNDLE=gg.foreversixty.app
xcrun simctl bootstatus "$FS_SIM_UDID" -b >/dev/null
xcrun simctl install "$FS_SIM_UDID" "$APP"
shoot() {
  xcrun simctl terminate "$FS_SIM_UDID" "$BUNDLE" >/dev/null 2>&1 || true
  xcrun simctl launch "$FS_SIM_UDID" "$BUNDLE" -FSTab bis >/dev/null
  sleep 6
  xcrun simctl io "$FS_SIM_UDID" screenshot "$OUT/$1-ios.png" >/dev/null
}
xcrun simctl ui "$FS_SIM_UDID" content_size large
shoot list-default
xcrun simctl ui "$FS_SIM_UDID" content_size accessibility-extra-extra-extra-large
shoot list-large
xcrun simctl ui "$FS_SIM_UDID" content_size large
xcrun simctl terminate "$FS_SIM_UDID" "$BUNDLE" >/dev/null 2>&1 || true
