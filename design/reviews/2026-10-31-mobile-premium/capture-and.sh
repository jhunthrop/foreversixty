#!/usr/bin/env bash
# F8 captures on the shared emulator (run under tools/mobile-device.sh android-run).
# Usage: capture-and.sh <out dir> <fixture debug apk>
# Signed in on the fixtures (me.json, a free account): You with the US storefront (See plans),
# with GB (no link) and with no readable storefront, each at font scale 1.0 and 2.0.
# The debug-only `fs_storefront` launch extra stands in for Play Billing's country.
set -euo pipefail
ADB="$HOME/Library/Android/sdk/platform-tools/adb"
OUT="$1"; APK="$2"
PKG=gg.foreversixty.app.debug

tap_text() {
  local b="" try
  for try in 1 2 3 4 5; do
    "$ADB" shell uiautomator dump /sdcard/ui.xml >/dev/null
    b=$("$ADB" shell cat /sdcard/ui.xml | python3 -c '
import re, sys, xml.etree.ElementTree as ET
want = sys.argv[1]
for node in ET.fromstring(sys.stdin.read()).iter("node"):
    if node.get("text", "").startswith(want) or node.get("content-desc", "").startswith(want):
        a, b, c, d = map(int, re.findall(r"\d+", node.get("bounds")))
        print((a + c) // 2, (b + d) // 2)
        break' "$1")
    [ -n "$b" ] && break
    sleep 2
  done
  [ -n "$b" ] || { echo "no node $1" >&2; return 1; }
  "$ADB" shell input tap $b
}
shot() { "$ADB" exec-out screencap -p > "$OUT/device-$(echo "$1" | tr A-Z a-z)-and.png"; }
scale() { "$ADB" shell settings put system font_scale "$1"; sleep 4; }
launch() {
  "$ADB" shell am force-stop "$PKG"
  "$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" --es fs_storefront "$1" >/dev/null
  sleep 5
  tap_text "You"; sleep 3
}

"$ADB" install -r "$APK" >/dev/null
"$ADB" shell pm clear "$PKG" >/dev/null
scale 1.0
"$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" >/dev/null
sleep 5
tap_text "You"; sleep 2
tap_text "SIGN IN WITH"; sleep 5
for region in US GB none; do
  launch "$region"
  scale 1.0; shot "you-free-$region"
  scale 2.0; shot "you-free-$region-large"
  scale 1.0
done
# The See plans row opens the browser (no in-app web view).
launch US
tap_text "Plan, Free, See plans"; sleep 4; shot "see-plans-browser"
"$ADB" shell input keyevent KEYCODE_BACK
scale 1.0
