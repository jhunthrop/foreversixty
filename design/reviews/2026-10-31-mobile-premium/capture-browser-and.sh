#!/usr/bin/env bash
# F8: See plans opens foreversixty.gg/premium in the browser (run under tools/mobile-device.sh
# android-run, after capture-and.sh has signed in). Steps past Chrome's first run if it shows.
# Usage: capture-browser-and.sh <out dir> <fixture debug apk>
set -euo pipefail
ADB="$HOME/Library/Android/sdk/platform-tools/adb"
OUT="$1"; APK="$2"
PKG=gg.foreversixty.app.debug
node_at() {
  "$ADB" shell uiautomator dump /sdcard/ui.xml >/dev/null
  "$ADB" shell cat /sdcard/ui.xml | python3 -c '
import re, sys, xml.etree.ElementTree as ET
want = sys.argv[1]
for node in ET.fromstring(sys.stdin.read()).iter("node"):
    if node.get("text", "").startswith(want) or node.get("content-desc", "").startswith(want):
        a, b, c, d = map(int, re.findall(r"\d+", node.get("bounds")))
        print((a + c) // 2, (b + d) // 2)
        break' "$1"
}
tap_if() { local b; b=$(node_at "$1"); [ -n "$b" ] && "$ADB" shell input tap $b; return 0; }
"$ADB" install -r "$APK" >/dev/null
"$ADB" shell settings put system font_scale 1.0
"$ADB" shell am force-stop "$PKG"
"$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" --es fs_storefront US >/dev/null
sleep 5
tap_if "You"; sleep 3
tap_if "SIGN IN WITH"; sleep 5
tap_if "Plan, Free, See plans"; sleep 5
for step in "Use without an account" "No thanks" "No, thanks" "Got it" "Accept & continue"; do tap_if "$step"; sleep 3; done
sleep 8
"$ADB" exec-out screencap -p > "$OUT/device-see-plans-browser-and.png"
"$ADB" shell input keyevent KEYCODE_BACK
