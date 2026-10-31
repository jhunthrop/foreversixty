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
  local b="" try
  for try in 1 2 3 4 5; do
  "$ADB" shell uiautomator dump /sdcard/ui.xml >/dev/null
  b=$("$ADB" shell cat /sdcard/ui.xml | python3 -c '
import re, sys, xml.etree.ElementTree as ET
want = sys.argv[1]
corner = want.startswith("@")
want = want.lstrip("@")
for node in ET.fromstring(sys.stdin.read()).iter("node"):
    if node.get("text", "").startswith(want) or node.get("content-desc", "").startswith(want):
        a, b, c, d = map(int, re.findall(r"\d+", node.get("bounds")))
        # "@corner" taps just inside the top-left (the first chip of the merged chip row).
        print(*((a + 60, b + 60) if corner else ((a + c) // 2, (b + d) // 2)))
        break' "$1")
  [ -n "$b" ] && break
  sleep 2
  done
  [ -n "$b" ] || { echo "no node $1" >&2; return 1; }
  "$ADB" shell input tap $b
}
# Closes the open sheet: its scrim when one shows, else the back key.
close_sheet() {
  tap_text "Close sheet" 2>/dev/null || "$ADB" shell input keyevent 4
  sleep 2
}
shot() { "$ADB" exec-out screencap -p > "$OUT/$1-and.png"; }
tap_text "BiS"; sleep 4; shot list-default
tap_text "Head."; sleep 3; shot tooltip-default
close_sheet
tap_text "@Choose a list"; sleep 3; shot picker-default
close_sheet
"$ADB" shell settings put system font_scale 2.0; sleep 3; shot list-large
# The picker at 2.0, scrolled to its end: Faction and Back must clear the system bars.
tap_text "@Choose a list"; sleep 3; shot picker-large-top
for _ in 1 2 3 4; do "$ADB" shell input swipe 540 1900 540 700 300; sleep 1; done
sleep 1; shot picker-large-bottom
close_sheet
tap_text "Head."; sleep 3
for _ in 1 2 3; do "$ADB" shell input swipe 540 1900 540 700 300; sleep 1; done
sleep 1; shot tooltip-large-bottom
close_sheet
"$ADB" shell settings put system font_scale 1.0
"$ADB" shell am force-stop "$PKG"
