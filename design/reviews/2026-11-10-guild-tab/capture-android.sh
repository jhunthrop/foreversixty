#!/usr/bin/env bash
# design/reviews/2026-11-10-guild-tab/capture-android.sh: the Guild tab's interaction walk on the shared emulator
# (spec 2026-11-10-guild-tab.md), dark and light: in a guild under 60 (Gorvash, 58), the officer's tab (the
# officer raid and badge fixtures), no guild under 60 (Yellfury, 31) and a raid link landing inside the Guild tab.
# Run from the repo root after `./gradlew :app:assembleDebug` (apps/android), holding the lock:
#   tools/mobile-device.sh android-run design/reviews/2026-11-10-guild-tab/capture-android.sh
set -euo pipefail

ADB="$ANDROID_HOME/platform-tools/adb"
APK=apps/android/app/build/outputs/apk/debug/app-debug.apk
PKG=$("$ANDROID_HOME/build-tools/$(ls "$ANDROID_HOME/build-tools" | sort -V | tail -1)/aapt2" dump packagename "$APK")
ROOT=design/reviews/2026-11-10-guild-tab/android
SETTLES=5
OFFICER="guild.raid=officer,guild.raid-badge=officer"

shot() { "$ADB" exec-out screencap -p > "$OUT/$1.png"; }

# The centre of the first node whose text or content-desc is (or, with a trailing *, begins with) $1; empty when none.
centre() {
	"$ADB" shell uiautomator dump /sdcard/ui.xml >/dev/null
	"$ADB" exec-out cat /sdcard/ui.xml | python3 -c '
import re, sys
want = sys.argv[1]
prefix = want.endswith("*")
for node in re.findall(r"<node [^>]*>", sys.stdin.read()):
    attrs = dict(re.findall(r"(\w[\w-]*)=\"([^\"]*)\"", node))
    for value in (attrs.get("text", ""), attrs.get("content-desc", ""), attrs.get("resource-id", "")):
        if (prefix and value.startswith(want[:-1])) or value == want:
            x1, y1, x2, y2 = map(int, re.findall(r"\d+", attrs["bounds"]))
            print((x1 + x2) // 2, (y1 + y2) // 2)
            sys.exit()' "$1"
}

tap() {
	local at
	at=$(centre "$1")
	[ -n "$at" ] || { echo "capture-android: no '$1' on screen" >&2; return 1; }
	# shellcheck disable=SC2086
	"$ADB" shell input tap $at
}

# Taps $1 and captures at once, after 400 ms and settled.
steps() {
	tap "$1" || return 1
	shot "$2-t0"
	sleep 0.4
	shot "$2-t400"
	sleep 2
	shot "$2-settled"
}

# Scrolls down until $1 shows (up to eight swipes).
reveal() {
	for _ in 1 2 3 4 5 6 7 8; do
		[ -n "$(centre "$1")" ] && return 0
		"$ADB" shell input swipe 540 1800 540 1100 400
		sleep 1
	done
	return 1
}

back() { "$ADB" shell input keyevent KEYCODE_BACK; sleep 2; }

# Brings the app back to the front after a site link opened the browser.
front() {
	local activity
	activity=$("$ADB" shell cmd package resolve-activity --brief "$PKG" | tail -1 | tr -d "\r")
	"$ADB" shell am start -W -n "$activity" >/dev/null
	sleep 2
}

# Switches to the character whose drawer row begins with $1 (the switch persists across launches).
pick() {
	tap "Open the menu" 2>/dev/null || tap "Open menu"
	sleep 2
	local list row
	list=$(centre "Switch character")
	row=$(centre "$1*")
	# Under Switch character only: the drawer's head names the current character, and a tap there opens it.
	if [ -n "$list" ] && [ -n "$row" ] && [ "${row#* }" -gt "${list#* }" ]; then
		# shellcheck disable=SC2086
		"$ADB" shell input tap $row
	else
		back
	fi
	sleep "$SETTLES"
}

launch() {
	"$ADB" shell am force-stop "$PKG"
	local activity
	activity=$("$ADB" shell cmd package resolve-activity --brief "$PKG" | tail -1 | tr -d "\r")
	if [ -n "${1:-}" ]; then
		"$ADB" shell am start -W -n "$activity" --es fs.debug.fixtures "$1" >/dev/null
	else
		"$ADB" shell am start -W -n "$activity" >/dev/null
	fi
	sleep "$SETTLES"
	if tap "Activity" 2>/dev/null; then sleep 2; tap "Sign in with Battle.net" 2>/dev/null && sleep "$SETTLES" || true; fi
	tap "Guild"
	sleep "$SETTLES"
	pick "${2:-Gorvash}"
}

"$ADB" install -r "$APK" >/dev/null
for mode in dark light; do
	[ "$mode" = dark ] && night=yes || night=no
	"$ADB" shell cmd uimode night "$night"
	OUT="$ROOT/$mode"
	mkdir -p "$OUT"

	# In a guild, under 60.
	launch ""
	shot leveler-01-top
	steps "Late" leveler-02-answer-late || true
	steps "In" leveler-03-answer-in || true
	if reveal "Roster ›"; then shot leveler-04-roster-section; steps "Roster ›" leveler-05-roster-page; back; fi
	reveal "Progression" && shot leveler-06-progression || true
	if reveal "Calendar"; then steps "Calendar" leveler-07-calendar || true; front; fi
	"$ADB" shell input swipe 540 700 540 2000 300
	"$ADB" shell input swipe 540 700 540 2000 300
	sleep 2
	shot leveler-08-pull-refresh

	# Officer.
	launch "$OFFICER"
	shot officer-01-top
	if reveal "Applications"; then shot officer-02-officer-rows; steps "Applications" officer-03-applications; back; fi
	if reveal "Setup"; then steps "Setup" officer-04-setup-site || true; front; fi
	if reveal "Plan a raid"; then steps "Plan a raid" officer-05-plan-site || true; front; fi
	"$ADB" shell input swipe 540 1800 540 1100 400
	sleep 1
	shot officer-06-scrolled

	# No guild, under 60: switch to Yellfury in the drawer.
	launch "" Yellfury
	shot noguild-01-top
	if reveal "Search guilds"; then steps "Search guilds" noguild-02-search-site || true; front; fi
	if reveal "Claim your guild"; then steps "Claim your guild" noguild-03-claim-site || true; front; fi
	pick Gorvash

	# The tab bar: the four icons with Guild unselected (Home, Gear) and selected.
	tap "Home"; sleep "$SETTLES"; shot tabbar-01-home
	tap "Gear"; sleep "$SETTLES"; shot tabbar-02-gear
	tap "Guild"; sleep "$SETTLES"; shot tabbar-03-guild

	# The deep link: a raid link opens the raid page inside the Guild tab.
	"$ADB" shell am start -W -a android.intent.action.VIEW -d "https://foreversixty.gg/app/raid/2/31" "$PKG" >/dev/null
	sleep "$SETTLES"
	shot link-01-raid-page
	back
	shot link-02-back-to-guild-root
done
"$ADB" shell cmd uimode night yes
