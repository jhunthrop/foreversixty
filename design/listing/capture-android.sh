#!/usr/bin/env bash
# design/listing/capture-android.sh ROOT OUT RAID_LINK SETTLES: the Android half of capture-apps.sh,
# run under the emulator lock. Demo mode gives the clean status bar; the font scale is set to the
# default for the shots and put back after.
set -euo pipefail

ROOT=$1 OUT=$2 RAID_LINK=$3 SETTLES=$4
ADB="$ANDROID_HOME/platform-tools/adb"
APK="$ROOT/apps/android/app/build/outputs/apk/debug/app-debug.apk"
[ -f "$APK" ] || { echo "capture-android: build the debug APK first (apps/android: ./gradlew :app:assembleDebug)" >&2; exit 1; }
PKG=$("$ANDROID_HOME/build-tools/$(ls "$ANDROID_HOME/build-tools" | sort -V | tail -1)/aapt2" dump packagename "$APK")

shot() { "$ADB" exec-out screencap -p > "$OUT/$1.png"; }

# Taps the centre of the first node whose text or content-desc is exactly $1.
tap_text() {
	"$ADB" shell uiautomator dump /sdcard/ui.xml >/dev/null
	local bounds
	bounds=$("$ADB" exec-out cat /sdcard/ui.xml | python3 -c '
import re, sys
want = sys.argv[1]
for node in re.findall(r"<node [^>]*>", sys.stdin.read()):
    attrs = dict(re.findall(r"(\w[\w-]*)=\"([^\"]*)\"", node))
    if want in (attrs.get("text"), attrs.get("content-desc")):
        print(attrs["bounds"]); break' "$1")
	[ -n "$bounds" ] || { echo "capture-android: no node '$1' on screen" >&2; return 1; }
	python3 -c '
import re, sys
x1, y1, x2, y2 = map(int, re.findall(r"\d+", sys.argv[1]))
print((x1 + x2) // 2, (y1 + y2) // 2)' "$bounds" | { read -r x y; "$ADB" shell input tap "$x" "$y"; }
}

# Scrolls until the first node whose text is $1 (case-insensitive) sits $2 px from the top, so the
# shot opens on a section heading and never on the tail of the one above it.
scroll_to() {
	local y
	"$ADB" shell uiautomator dump /sdcard/ui.xml >/dev/null
	y=$("$ADB" exec-out cat /sdcard/ui.xml | python3 -c '
import re, sys
want = sys.argv[1].lower()
for node in re.findall(r"<node [^>]*>", sys.stdin.read()):
    attrs = dict(re.findall(r"(\w[\w-]*)=\"([^\"]*)\"", node))
    if attrs.get("text", "").lower() == want:
        print(re.findall(r"\d+", attrs["bounds"])[1]); break' "$1")
	[ -n "$y" ] || { echo "capture-android: no '$1' to scroll to" >&2; return 1; }
	# A slow drag (2 s) does not fling, so the page moves by the drag's length.
	"$ADB" shell input swipe 540 "$y" 540 "$2" 2000
}

demo() { "$ADB" shell am broadcast -a com.android.systemui.demo -e command "$@" >/dev/null; }

font_scale=$("$ADB" shell settings get system font_scale | tr -d '\r')
restore() {
	demo exit
	"$ADB" shell settings put system font_scale "${font_scale:-1.0}"
}
trap restore EXIT

"$ADB" install -r "$APK" >/dev/null
"$ADB" shell settings put system font_scale 1.0
"$ADB" shell cmd uimode night yes
"$ADB" shell settings put global sysui_demo_allowed 1
demo enter
demo clock -e hhmm 0941
demo battery -e level 100 -e plugged false
demo network -e wifi show -e level 4 -e fully true
demo network -e mobile hide
demo notifications -e visible false

"$ADB" shell am force-stop "$PKG"
"$ADB" shell am start -W -n "$("$ADB" shell cmd package resolve-activity --brief "$PKG" | tail -1 | tr -d '\r')" >/dev/null
sleep "$SETTLES"
# Fixture mode signs in at once from Activity's Battle.net button, if the app is signed out.
if tap_text "Activity" 2>/dev/null; then sleep 2; tap_text "Sign in with Battle.net" 2>/dev/null && sleep "$SETTLES" || true; fi
tap_text "Home"
sleep "$SETTLES"
shot home
tap_text "Gear"
sleep "$SETTLES"
shot gear

"$ADB" shell am start -W -a android.intent.action.VIEW -d "$RAID_LINK" "$PKG" >/dev/null
sleep "$SETTLES"
shot raid
"$ADB" shell input swipe 540 1800 540 1300 2000
scroll_to "Raid composition" 330
sleep "$SETTLES"
shot raid-roster
