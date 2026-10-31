# Mobile Raids integration C: UX verdict (2026-10-31)

Reviewed captures only: live Pixel emulator and iOS 27 simulator captures, plus the Roborazzi and CI (iOS 26.5) goldens named in the brief. Spec: design/specs/2026-10-30-mobile-apps.md, 3.4 (including "Deviation (integration C)"), 4.1 to 4.5, F4. Scope: what changed since the F4 verdict.

## Verdict: SHIP

No blockers. The four changes render as specified. The shoulds below do not gate the merge.

## Blockers

None.

## Judged and passing

1. **Bless fallback label.** The earlier contrast must-fix is resolved.
   - raid-composition-integrated-and and raid-page-roster-ios: the Bless tile has a dimmed ground and the label "Bless" is clearly legible in muted text. It reads as missing without being unreadable. Fort, MotW, Int and the other present tiles keep strong labels and full-colour art. "Missing: Blessings" sits under the grid.
   - The Android and iOS officer goldens draw every tile on its label, with Bless dimmed and still readable. Both platforms match.
2. **One selector.**
   - Android: the same pill (crest 34, class-coloured full name, `Fury · 58`, chevron) at the same y on Home, Raids and BiS at 1.0x. At 1.3x the description line drops and the name stays whole. At 2.0x the name is capped and the pill stays one line. Nothing is cut at any scale. BiS carries the search button on the same row. The full-width 48 dp row is the recorded deviation and is accepted.
   - Android: selector-home-relaunch-loaded-and is pixel-identical in position to selector-home-1.0-and.
   - iOS: the same pill at the same y on Home, Raids and BiS. The Raids root has no inline title next to it. This closes the earlier iOS tab-top should. The BiS search is a glass circle at the trailing edge.
3. **Shared design system.** In the Android and iOS goldens (reminder-ask, offline, tab-loaded, xl):
   - Cards use the warm card surface for the next raid and the raised surface for quiet cards.
   - The Planned, Locked and Not signed up tags are 22 pt capsules in uppercase.
   - The offline strip is the same blue-grey bar with the icon and "Offline. Showing what you had at 3:30 pm.", and the segments are disabled.
   - The segments, date blocks and bars match on both platforms.
   - The reminder ask uses gold-accent `Turn on reminders` as the secondary button, with `Not now` beside it. Copy matches `ask.turnOn`.
   - The XL goldens keep every earlier fix: no mid-word breaks, a stacked who-line and a 2x2 segment grid.
4. **Band ground.** account.signedout-default (iOS and Android) shows the `#14100b` ground with a softer ember glow at the top. It is visibly calmer than the earlier 0.28 glow. The two platforms match, and the headline, gold primary and secondary button read clearly. Home's hero band uses the same ground in the live captures.

## Shoulds

1. **`Not now` renders uppercase on both platforms** (raid-page-reminder-ask-and and -ios). The spec and `ask.notNow` say `Not now` as a text button in `--color-nav`. Set it in sentence case. This is the same sentence-case rule you applied to the secondary buttons.
2. **iOS goldens show a stale top bar.**
   - raids-tab-loaded-ios and raids-tab-offline-ios still draw a centred inline "Raids" title and no selector. The live capture has the selector and no title.
   - Make the golden host render the same toolbar, or the goldens are no evidence for the change under review.
3. **Faction mark differs by platform.**
   - The guild head fallback is a red star on Android and a red shield on iOS.
   - The spec says the faction mark, as the site's `GuildCrest` falls back. Use the one Horde mark on both.
4. **Dead space under the segments on the next-raid card** (carry-over should 1 from the F4 verdict, improved but not gone).
   - About 32 dp (Android) and 23 pt (iOS) between the segments and `Who is coming`, then about 30 dp under the link.
   - When the status line has no text, collapse its slot. Keep 8 above the link.
5. **Mono is still too wide in rows** (carry-over). "11 signed up" in Also planned is set whole in mono. Mono only the digits, and keep "signed up" in Barlow. This is the same fix the header already has for "24 of 40 signed up".
6. **Android XL note field.** The placeholder "Note for the officers (optional)" wraps to two lines and the field's bottom border crosses the parenthesis (raid-page-xl-and). Let the field grow with its text.
7. **iOS raid page repeats the raid name.**
   - The inline nav title and the large heading both say Onyxia's Lair (officer golden, reminder-ask golden).
   - The Android app bar carries the title once. Match iOS to that, or show the name only in the heading once scrolled.
8. **Evidence still missing for the selector.**
   - Android: a capture of the reserved 48 dp row while characters are still loading, next to the loaded one, to prove nothing shifts.
   - iOS: the selector at AX1 and at AX5 on Home.
9. **iOS 27 floating tab bar overlaps the last visible row** in the live captures (roster, BiS, Raids). This is platform behaviour, and the content scrolls clear. Confirm the bottom inset lets the final row scroll fully above the bar.

## Checklist

- [x] Bless label legible on a dimmed tile, both platforms, live and golden
- [x] One selector, same place and look on Home, Raids and BiS, both platforms
- [x] Name never cut at 1.3x, 2.0x (Android)
- [x] Description drops at large text (Android)
- [x] iOS Raids root has no inline title beside the selector (live)
- [x] Android 48 dp row deviation documented in spec 3.4
- [x] Cards, tags, offline strip and secondary buttons are the DS ones
- [x] Reminder ask uses a gold accent secondary
- [x] Band ground and ember glow match the spec on both platforms
- [ ] `Not now` in sentence case (should 1)
- [ ] iOS goldens match the live toolbar (should 2)
- [ ] One faction mark on both platforms (should 3)
- [ ] Loading-state selector capture (should 8)
