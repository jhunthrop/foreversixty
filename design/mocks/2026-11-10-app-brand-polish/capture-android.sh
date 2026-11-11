#!/usr/bin/env bash
# design/mocks/2026-11-10-app-brand-polish/capture-android.sh: the Android review captures for lane
# app-brand-polish on the shared emulator: the four tab roots with the head fade, at rest and
# scrolled, under the system's dark and light themes, plus the launcher with the new icon.
# Run from the repo root after `./gradlew :app:assembleDebug` (apps/android), holding the lock:
#   tools/mobile-device.sh android-run design/mocks/2026-11-10-app-brand-polish/capture-android.sh
set -euo pipefail

ADB="$ANDROID_HOME/platform-tools/adb"
APK=apps/android/app/build/outputs/apk/debug/app-debug.apk
PKG=$("$ANDROID_HOME/build-tools/$(ls "$ANDROID_HOME/build-tools" | sort -V | tail -1)/aapt2" dump packagename "$APK")
OUT=design/reviews/2026-11-10-app-brand-polish
SETTLES=5

shot() { "$ADB" exec-out screencap -p > "$OUT/$1-android.png"; }

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

mkdir -p "$OUT"
"$ADB" install -r "$APK" >/dev/null
for mode in dark light; do
	[ "$mode" = dark ] && night=yes || night=no
	"$ADB" shell cmd uimode night "$night"
	"$ADB" shell am force-stop "$PKG"
	"$ADB" shell am start -W -n "$("$ADB" shell cmd package resolve-activity --brief "$PKG" | tail -1 | tr -d "\r")" >/dev/null
	sleep "$SETTLES"
	# Fixture mode signs in at once from Activity's Battle.net button, if the app is signed out.
	if tap_text "Activity" 2>/dev/null; then sleep 2; tap_text "Sign in with Battle.net" 2>/dev/null && sleep "$SETTLES" || true; fi
	tap_text "Home"
	sleep "$SETTLES"
	shot "fade-home-$mode"
	for tab in Raid Gear Activity; do
		tap_text "$tab"
		sleep "$SETTLES"
		shot "fade-$(echo "$tab" | tr '[:upper:]' '[:lower:]')-$mode"
	done
	"$ADB" shell input swipe 540 1700 540 900 300
	sleep 2
	shot "fade-activity-scrolled-$mode"
done
"$ADB" shell cmd uimode night yes
"$ADB" shell input keyevent KEYCODE_HOME
sleep 2
shot launcher-home
