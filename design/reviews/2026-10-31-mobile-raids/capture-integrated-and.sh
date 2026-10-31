#!/usr/bin/env bash
# Integration C captures on the shared emulator (run under tools/mobile-device.sh android-run).
# Usage: capture-integrated-and.sh <out dir> <fixture debug apk>
# Signed in on the fixtures: the character selector row closed on Home at font scale 1.0, 1.3 and
# 2.0; the same row on Raids and BiS; Home again after a relaunch; and the Raids composition block
# with the fallback labels.
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
shot() { "$ADB" exec-out screencap -p > "$OUT/$1-and.png"; }
scale() { "$ADB" shell settings put system font_scale "$1"; sleep 4; }

"$ADB" install -r "$APK" >/dev/null
"$ADB" shell pm clear "$PKG" >/dev/null
scale 1.0
"$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" >/dev/null
sleep 5
tap_text "You"; sleep 2
tap_text "SIGN IN WITH"; sleep 5
tap_text "Home"; sleep 5
shot selector-home-1.0
scale 1.3; shot selector-home-1.3
scale 2.0; shot selector-home-2.0
scale 1.0
tap_text "Raids"; sleep 5; shot selector-raids-1.0
tap_text "BiS"; sleep 5; shot selector-bis-1.0
# The relaunch: the session restores the token and reads /v1/me; the row sits where it did.
"$ADB" shell am force-stop "$PKG"
"$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" >/dev/null
sleep 5; shot selector-home-relaunch-loaded
# The composition block: the raid page, scrolled to Sign-ups.
tap_text "Raids"; sleep 4
tap_text "Who is coming"; sleep 4
"$ADB" shell input swipe 540 1900 540 900 600; sleep 2
shot raid-composition-integrated
"$ADB" shell am force-stop "$PKG"
