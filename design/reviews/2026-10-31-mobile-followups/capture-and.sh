#!/usr/bin/env bash
# Mobile follow-up captures on the shared emulator (run under tools/mobile-device.sh android-run).
# Usage: capture-and.sh <out dir> <fixture debug apk>
# Signed in on the fixtures: Home and BiS for the level-58 Fury warrior (both on 50 to 59); the
# selector on Raids and Home for Yellfury (Hunter · 31) at font scale 1.0 and 2.0; and the raid
# page's composition block with every buff tile named.
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
swipe_up() { "$ADB" shell input swipe 540 1900 540 900 600; sleep 2; }
pick() { tap_text "Character, $1"; sleep 3; tap_text "$2"; sleep 5; }

"$ADB" install -r "$APK" >/dev/null
"$ADB" shell pm clear "$PKG" >/dev/null
scale 1.0
"$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" >/dev/null
sleep 5
tap_text "You"; sleep 2
tap_text "SIGN IN WITH"; sleep 5
tap_text "Home"; sleep 5
shot home-band-58
swipe_up; shot home-band-58-scrolled
tap_text "BiS"; sleep 5; shot bis-band-58
tap_text "Raids"; sleep 5
pick "Obnoxious Yell" "Yellfury"
shot selector-raids-default
tap_text "Home"; sleep 4; shot selector-home-default
scale 2.0
shot selector-home-ax5
tap_text "Raids"; sleep 4; shot selector-raids-ax5
scale 1.0
pick "Yellfury" "Obnoxious Yell"
tap_text "Who is coming"; sleep 4
swipe_up; shot raid-buffs
"$ADB" shell input swipe 540 1500 540 1000 600; sleep 2; shot raid-buffs-2
"$ADB" shell am force-stop "$PKG"
