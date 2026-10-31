#!/usr/bin/env bash
# F6 Talents captures on the shared emulator (run under tools/mobile-device.sh android-run).
# Usage: capture-and.sh <out dir> <fixture debug apk>
# Signed in on the fixtures: Home's Talents row, the pushed screen for Obnoxious Yell (level 58,
# the 50 to 59 guide), the tree scrolled with the pinned compare bar, a rank sheet, and font scale 2.0.
set -euo pipefail
ADB="$HOME/Library/Android/sdk/platform-tools/adb"
OUT="$1"; APK="$2"
PKG=gg.foreversixty.app.debug

tap_text() {
  local b="" try
  for try in 1 2 3 4 5 6 7 8 9 10; do
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
    "$ADB" shell input swipe 540 1700 540 1100 400
    sleep 2
  done
  [ -n "$b" ] || { echo "no node $1" >&2; return 1; }
  "$ADB" shell input tap $b
}
shot() { "$ADB" exec-out screencap -p > "$OUT/$1-and.png"; }
scale() { "$ADB" shell settings put system font_scale "$1"; sleep 4; }
open_talents() {
  "$ADB" shell am force-stop "$PKG"
  "$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" >/dev/null
  sleep 6
  tap_text "Talents"; sleep 5
}

"$ADB" install -r "$APK" >/dev/null
"$ADB" shell pm clear "$PKG" >/dev/null
scale 1.0
"$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" >/dev/null
sleep 5
tap_text "You"; sleep 2
tap_text "SIGN IN WITH"; sleep 5
tap_text "Home"; sleep 5
"$ADB" shell input swipe 540 1700 540 900 500; sleep 2
shot home-talents-row
tap_text "Talents"; sleep 5
shot app-talents-differs
"$ADB" shell input swipe 540 1800 540 900 500; sleep 2
shot app-talents-differs-scrolled
tap_text "Gore Drinker"; sleep 3
shot app-talents-rank
"$ADB" shell input keyevent KEYCODE_BACK; sleep 2
scale 2.0
open_talents
shot app-talents-differs-large
"$ADB" shell input swipe 540 1800 540 700 500; sleep 2
shot app-talents-differs-large-tree
scale 1.0
"$ADB" shell am force-stop "$PKG"
