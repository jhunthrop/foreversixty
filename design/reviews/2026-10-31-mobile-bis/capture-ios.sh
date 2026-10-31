#!/usr/bin/env bash
# iOS captures of the BiS tab on the shared simulator (run under tools/mobile-device.sh ios-run,
# which exports FS_SIM_UDID). The debug build opens on the tab named by -FSTab.
set -euo pipefail
OUT="$1"; APP="$2"
BUNDLE=gg.foreversixty.app
xcrun simctl bootstatus "$FS_SIM_UDID" -b >/dev/null
xcrun simctl install "$FS_SIM_UDID" "$APP"
# shoot <name> [extra launch arguments]: the debug build opens on the tab -FSTab names, and
# `-FSBisPicker open` opens the real picker sheet over it.
shoot() {
  local name="$1"; shift
  xcrun simctl terminate "$FS_SIM_UDID" "$BUNDLE" >/dev/null 2>&1 || true
  xcrun simctl launch "$FS_SIM_UDID" "$BUNDLE" -FSTab bis "$@" >/dev/null
  sleep 7
  xcrun simctl io "$FS_SIM_UDID" screenshot "$OUT/$name-ios.png" >/dev/null
}
xcrun simctl ui "$FS_SIM_UDID" content_size large
shoot list-default
shoot picker-default -FSBisPicker open
xcrun simctl ui "$FS_SIM_UDID" content_size accessibility-extra-extra-extra-large
shoot list-large
shoot picker-large -FSBisPicker open
xcrun simctl ui "$FS_SIM_UDID" content_size large
xcrun simctl terminate "$FS_SIM_UDID" "$BUNDLE" >/dev/null 2>&1 || true
