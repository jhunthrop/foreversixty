#!/usr/bin/env bash
# Integration captures of the BiS tab on the shared emulator (run under tools/mobile-device.sh
# android-run). Usage: capture-integrated-and.sh <out dir> <fixture debug apk> <live debug apk>
# - the fixture build, signed in through You: the list under the selector and the four-tab bar,
#   at font scale 1.0 and 2.0;
# - the live build (the real CDN files), signed out: the tooltip from a healer and a tank list.
set -euo pipefail
ADB="$HOME/Library/Android/sdk/platform-tools/adb"
OUT="$1"; FIXTURE_APK="$2"; LIVE_APK="$3"
PKG=gg.foreversixty.app.debug

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
        print(*((a + 60, b + 60) if corner else ((a + c) // 2, (b + d) // 2)))
        break' "$1")
    [ -n "$b" ] && break
    sleep 2
  done
  [ -n "$b" ] || { echo "no node $1" >&2; return 1; }
  "$ADB" shell input tap $b
}
# Closes the open sheet: its scrim when one shows, else the back key.
close_sheet() { tap_text "Close sheet" 2>/dev/null || "$ADB" shell input keyevent 4; sleep 2; }
expand_sheet() { for _ in 1 2; do "$ADB" shell input swipe 540 1700 540 500 300; sleep 1; done; }
shot() { "$ADB" exec-out screencap -p > "$OUT/integrated-$1-and.png"; }
fresh() {
  "$ADB" install -r "$1" >/dev/null
  "$ADB" shell pm clear "$PKG" >/dev/null
  "$ADB" shell settings put system font_scale 1.0
  "$ADB" shell am start -n "$PKG/gg.foreversixty.app.MainActivity" >/dev/null
  sleep 5
}

# Signed in on the fixtures: the character's own list under the selector.
fresh "$FIXTURE_APK"
tap_text "You"; sleep 2
tap_text "SIGN IN WITH"; sleep 5
tap_text "BiS"; sleep 5; shot list-default
"$ADB" shell settings put system font_scale 2.0; sleep 4; shot list-large
"$ADB" shell settings put system font_scale 1.0; sleep 2

# Signed out on the live CDN: the first visit opens the picker.
pick_list() { # class, spec, then a slot row to open
  tap_text "$1"; sleep 2
  tap_text "$2"; sleep 2
  tap_text "60 · Phase 1 raid"; sleep 2
  close_sheet; sleep 3
}
fresh "$LIVE_APK"
tap_text "BiS"; sleep 5
pick_list "Priest" "Holy"
tap_text "Shoulder."; sleep 4; expand_sheet; shot tooltip-healer
close_sheet
tap_text "@Choose a list"; sleep 3
pick_list "Warrior" "Protection"
tap_text "Neck."; sleep 4; expand_sheet; shot tooltip-tank
close_sheet
"$ADB" shell am force-stop "$PKG"
