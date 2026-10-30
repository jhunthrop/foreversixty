#!/usr/bin/env bash
# Android captures of the BiS tab on the shared emulator (run under tools/mobile-device.sh android-run).
set -euo pipefail
ADB="$HOME/Library/Android/sdk/platform-tools/adb"
OUT="$1"; APK="$2"
PKG=gg.foreversixty.app.debug
"$ADB" install -r "$APK" >/dev/null
"$ADB" shell pm clear "$PKG" >/dev/null
"$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" >/dev/null
sleep 4
tap_text() {
  "$ADB" shell uiautomator dump /sdcard/ui.xml >/dev/null
  local b
  b=$("$ADB" shell cat /sdcard/ui.xml | python3 -c '
import re, sys, xml.etree.ElementTree as ET
want = sys.argv[1]
for node in ET.fromstring(sys.stdin.read()).iter("node"):
    if node.get("text", "").startswith(want) or node.get("content-desc", "").startswith(want):
        a, b, c, d = map(int, re.findall(r"\d+", node.get("bounds")))
        print((a + c) // 2, (b + d) // 2)
        break' "$1")
  [ -n "$b" ] || { echo "no node $1" >&2; return 1; }
  "$ADB" shell input tap $b
}
shot() { "$ADB" exec-out screencap -p > "$OUT/$1-and.png"; }
tap_text "BiS"; sleep 4; shot list-default
tap_text "Head."; sleep 3; shot tooltip-default
"$ADB" shell input keyevent 4; sleep 2
tap_text "Choose a list"; sleep 3; shot picker-default
"$ADB" shell input keyevent 4; sleep 2
"$ADB" shell settings put system font_scale 2.0; sleep 3; shot list-large
"$ADB" shell settings put system font_scale 1.0
"$ADB" shell am force-stop "$PKG"
