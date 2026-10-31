# UX verdict: mobile F1 sign-in, account, switcher (2026-10-31)

Reviewed: the renders in this folder (iOS snapshots 402x874, Android Roborazzi), against design/specs/2026-10-30-mobile-apps.md §F1, §3.4, F8.1 and the f1/f8 boards. iOS confirmation dialogs were not capturable; I read their copy from the string catalog.

**Verdict: HOLD**

## Blockers

1. **You is missing Card two and most of Card one (F8.1).** Both platforms show only Plan, Devices, Delete account, Sign out. Spec: Notifications, Characters, Discord, then Help on the Discord server, Privacy, Terms, About Forever Sixty (with the Blizzard disclaimer), and the version line. Privacy and Terms must be reachable in the app (App Store 5.1.1(i)) and sit beside an account-deletion flow. Ship Privacy, Terms, About and the version line in this lane; the three feature rows may stay out until their lanes land.
2. **Android tab bar disagrees with the spec and with iOS.** Android: Character, BiS, Raids, Sims, You. Spec §3.1 and iOS: Home, Raids, BiS, You (Sims is pushed from Home). At AX5 the Android label "Character" breaks mid-word ("Chara / cter", signedout-ax5-and), which XL.3 forbids. This lane edited ForeverTab.kt (the You label), so it owns the fix: rename and reorder to match the spec, and keep tab labels on one line at 2.0x.
3. **The deleted-account notice renders as an error.** signedout-deleted-ios/-and use the red alert toast (wipe border, "!" glyph) for a completed request. Copy is also stilted: `Your account is being deleted. It is gone by 29 Nov 2026.` Use a neutral toast (no wipe colour, check or info glyph) with `Your account is deleted. Everything is removed by 29 Nov 2026.`. Red is for failures only (spec §5.0).

## Shoulds

Delete account (new, judged on flow, copy, confirmation)
- Flow is right: a visible row on You, a screen that lists exactly what goes (it matches the privacy page), a destructive button, one confirmation alert, then signed out with the date. Keep it.
- The alert body `This cannot be undone.` repeats nothing the screen says. Make it `Your characters, logs and sims are removed within thirty days. Every device is signed out now.` and name the button `Delete account` (iOS `.destructive`, Android red), not `Delete`.
- Guild owners: deleting cancels the guild plan and the seats it gave. The screen is silent. Add a line when the account owns a guild plan: `Your guild's plan is cancelled and the seats you gave end.` Same rule as the plan line: show only when it applies.
- Sign-in with the same Battle.net during the 30 days is undefined. Decide and state it on the screen (`Signing in again before then does not restore it.`).
- The Delete row on You is a red full card above Sign out, so a mis-tap lands on a destructive screen. Move it below Sign out as a quiet text row (red text, no card), or keep it and separate by 24 pt. The screen already requires a second tap and the alert, so this is not a blocker.
- The failed capture is named `delete-plan-failed` but shows the plan line plus the error. Rename or add the capture with no plan line and a failed request.
- Attach iOS captures of the three alerts from a UI test or the simulator. I could not verify them.

Parity
- Premium Plan row: iOS wraps `Premium · renews 27 Oct 2027` to two right-aligned lines beside the Founder tag; Android keeps one line. Spec puts the plan line as meta under `Plan`, tag on the right. Use that on both.
- Account head lacks the gold `Premium` tag for Premium (you-premium both). Add it.
- Device icon: Android uses a handset (phone call) glyph; iOS an iPhone outline. Use a smartphone outline on Android (vector asset) or the handset reads as "call".
- You title: iOS sentence case `You`, Android condensed caps `YOU`; the same for Devices and Delete account. The mock uses condensed caps on both. Pick the mock.
- Android has no status bar inset in the renders, so toast placement (12 pt below the bar) is unverified. Add an inset-aware capture.

Switcher (passes the core)
- One selector, full name ("Obnoxious Yell"), ringed HD crests, gold open border with glow, check on the current row, spec footer copy. Both platforms match the board.
- At AX5 the iOS selector stops at 28 pt and drops its description: correct. In the sheet the iOS rows break "Obnoxious / Yell" and push the guild line two lines down; acceptable, but the sheet title at AX5 is 48 pt and eats a third of the screen. Cap it.
- The sheet drag handle is not visible on iOS; confirm it exists.
- Not yet placed on Home, Raids, BiS (report concern 3). That is the F2/F3/F4 lanes, but nobody can use the switcher until it is. Track it as a ship gate for the wave.

Signed out and errors
- F1.1, F1.4 and 401 match the spec to the pixel in copy, 52 pt buttons, one primary. Failed text uses the copy catalog, which is right.
- iOS AX5 capture is clipped at the viewport and shows the second button cut off. Attach a scrolled-to-bottom capture showing both buttons and the meta line. Android AX5 fits.
- No-characters: no seal and no tab context in the capture; it reads as a bare page. Add the seal as in F1.3 or confirm the real screen sits under the tab bar.
- Cancel of the system sheet returns to F1.1 silently: confirm by test (listed as tested).

Devices
- Matches F1.6 including the inline failure under the row and the footer. The revoke row's button needs a 44 pt hit target: it renders ~40 pt tall on both. Raise it.
