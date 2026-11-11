# Guild tab: ux-designer verdict, round 3 (2026-11-11)

Reviewed on images: ios/dark and android/dark leveler-01, -02 (Late), -03 (In), leveler-08, officer-02 (Android), tabbar set; parity/ios-full-guild-tab-leveler.png, parity/ios-full-guild-tab-officer.png, parity/android-full-guild-tab-officer.png. Round-3 claims that rest on unopened captures (Your attendance, Gear gaps, the gold loot award) were not checked.

## Verdict: NOT YET

The round-2 must is closed on both devices. One new must: the full-page leveler render contradicts the round-3 claim that Who is coming is roles-only under 60.

## Round-2 musts and shoulds against the captures

### Closed
- **Late and In select (round-2 must 1).** ios and android leveler-02-answer-late-settled: Late filled teal, minute chips, "Saved."; leveler-03-answer-in-settled: In filled, chips gone. Closed.
- **Under 60, Before you pull is left out.** Absent in leveler-01, -03 and -08 on both. Closed.
- **Under 60 on the device captures, Who is coming is roles only** ("Tanks 3", "Healers 6", "Damage 15 · 7 melee, 8 ranged"). Closed on the device, not on the full-page render (Must 1).
- **Quiet line "Raids open at 60 · you're 58"** in the text ink under the raid card. Closed.
- **Applicant line.** Closed.
- **Tab bar.** Unchanged since round 2, passed.

### Still open from round 2
- **Should 6 ("Saved." shifts the layout).** "Bring flasks" moves about 83 px (iOS) and 74 px (Android) when "Saved." appears; Late adds a chip row too. Check that "Saved." clears after about 2 s and capture the cleared frame.
- **Should 5 (Android navigating rows carry no chevron):** Plan a raid, Applications, Setup in android/dark/officer-02-officer-rows.png; iOS has them.
- **Should 3 (external-link glyph for hand-offs).** Not addressed.
- **Should 4 (Android Calendar leaves the app).** Recorded as unchanged, no owner.
- **Leveler-08 pull-refresh** shows no pull indicator.

## Findings

### Must
1. **The full-page leveler render still shows the full 60-only Who is coming panel** (parity/ios-full-guild-tab-leveler.png): "Tanks 2 of 2" in green, "Healers 5 of 8 · 3 short", "Damage 11 of 30 · 19 short", "Brightwen (Late · 15 min)", under "Raids open at 60 · you're 36". The device capture complies, so one view renders two ways. Fix the path or re-render with roles only; attach the Android pair.

### Should
2. **Under Late, the counts fall from "Tanks 3, Healers 6, Damage 15" to "Tanks 1 (Kraggor), Damage 1 (Gorvash)"** (leveler-02-answer-late-settled on both). The fixture echo should move only your own row and keep the rest of the roster.
3. **"Saved." shift.** Reserve the line height.
4. **Android chevrons** on Plan a raid, Applications and Setup; an external-link glyph on hand-off rows on both.
5. **Calendar parity.** Record the Android calendar leaving the app with an owner, or push the list.

### Could
6. Late chips: Android has "How late?", iOS does not. Pick one.
7. Compare Ping on Discord once with Home and the web planner button.
8. No loading, empty or error capture of the guild home.
9. Capture the mid-pull frame.

If Must 1 is cleared by a re-render or a fix that matches the device walk, this is SHIP with the shoulds as follow-ups.
