# Guild tab: ux-designer verdict (2026-11-11)

Reviewed on screenshots only: ios/ and android/ dark captures (light compared to dark by pixel diff), parity goldens, ios-full renders, ax5 renders, mocks. Spec: design/specs/2026-11-10-guild-tab.md. Rulings applied: a player under 60 signs up with the quiet line (never a block); apps are dark-only.

## Verdict: NOT YET

The structure matches the spec and the mocks: Guild head, spread, live raid head, Roster, Recent logs, Progression, Pages, Officer rows, no-guild state, raid-link push and back all render and read well. Five items block: the quiet line is unreadable on iOS, Android prints "Holy Paladin Paladin", the Late/In answer never visibly changes in any frame, the site hand-offs land signed out or in an error, and the unselected Guild icon was never captured. Fix the musts, re-capture, and this ships.

## Findings

### Must

1. **iOS, quiet line "Raids open at 60 - you're N", contrast.** Where: ios/dark/leveler-01-top.png (y 1647), noguild-01-top.png (y 570). Shipped: grey (139,135,132) on the tan/blue gradient, measured 1.9:1. Android renders the same line at 5.0:1 (226,220,211). Fix: use the Android token (cream at about 80 percent) on both; the line stays quiet by size (13 pt) and position, not by contrast. On no-guild it also sits on the blue-to-tan seam; same token, 12 pt of space under the title.

2. **Android, Applications row subtitle doubles the class.** Where: android/dark/officer-02-officer-rows.png: "Test Applicant, Holy Paladin Paladin 60". iOS says "Test Applicant, Holy Paladin 60". Fix: use the shared "Spec Class level" formatter the iOS row and the applicant sheet use. Check the Needs you applicant line too.

3. **The answer buttons never show the answer.** Where: ios and android leveler-02 (Late) and -03 (In), t0, +400 ms and settled. Shipped: after tapping Late the segmented control still highlights In in all three frames on both platforms; only "Saved." appears. Spec step: the four answers work at any level. As captured, nothing proves Late selects. interactions.md blames the fixture echoing one canned answer. Fix: make the fixture echo the tapped answer, re-capture Late then In, show the tapped segment filled at t0 (optimistic), and the counts under Who is coming changing with it.

4. **Site hand-offs land signed out or in an error, so the officer path is unproven.** Where: android officer-05-plan-site-settled.png shows the guild's Raids tab with "Sign in with Battle.net", "No public raid nights yet." and no Plan a raid; ios officer-04-setup-site-settled.png shows "The setup did not load. Try again"; ios officer-04 t400 is a black screen; ios noguild-02 and -03 and android leveler-07 (Calendar) are loading skeletons or blank at +400 ms. Spec 2.3: each officer row is one tap from the tab to the thing. If the system browser does not carry the app's sign-in, an officer taps Plan a raid and meets a login wall. Fix: hand the session over (one-time sign-in link or in-app browser with the token), or capture a real signed-in run for Plan a raid, Setup and Open in planner. Until a capture shows the destination page itself, those officer steps are not verified.

5. **Tab bar: the unselected Guild icon is in no capture.** Every tab-root capture is taken on Guild, so only the selected Guild is seen beside unselected Home, Gear and Activity. The brief and interactions.md both say the four icons are checked selected and unselected. Fix: capture the tab bar from Home and from Gear on both platforms, dark. Judge finding 6 on that, where the weight difference matters most.

### Should

6. **Tab bar: the Guild door does not yet read as one family with Home, Gear and Activity.** Compared on ios noguild-01-top and android leveler-01-top:
   - Colour and opacity: selected gold on a pill, unselected cream (iOS) or muted blue-grey (Android), the same tint path as the other three. Pass, as far as one frame shows.
   - Size: 24 pt / 24 dp box, but the banner is narrow. Visible width is about 17 pt (iOS) and 17 dp (Android) against 20 to 22 for the stone and 18 to 20 for the helm and horn. It reads smaller.
   - Weight: Home, Gear and Activity are outlined drawings with interior line detail. Guild is a solid flat silhouette with one shield cut-out and carries clearly more ink. Unselected cream, it will read as the heaviest of the four. Fix: outline plus the same interior detail, or reduce the fill so ink coverage is within about 15 percent of Gear; widen the cloth by about 2 pt.
   - Finials: the top bar and its ends are the thinnest part of the glyph. At 24 they collapse to a stub on Android (top-left) and the badge covers the right end on both platforms. Thicken the rod to the Gear helm's stroke; keep the badge at the Activity offset.
   - Pill: iOS lighter tan, Android green-grey; platform tokens, not a Guild defect.

7. **Under 60, Who is coming sits above Roster and carries 60-only content.** Where: parity/ios-full-guild-tab-leveler.png. The full panel (Tanks, Healers 5 of 8 - 3 short, Damage 11 of 30 - 19 short, Missing Fort - WF 1 of 2) fills the first screen below the raid head; Roster, the leveler's reason to open the tab, starts about 1.7 screens down. Spec 2.1 asked for the counts line and the roles line. Fix: under 60 show those two lines; keep the full panel for 60. The Roster action also appears twice (Who is coming and Roster) to the same page; keep the one on the Roster section.

8. **Large text (ax5), the guild head breaks.** Where: parity/leveler.ios-guild-tab-leveler-ax5.png. "38 members - Alliance - US PvP" wraps to three lines, "55 to 59" wraps to "55 to / 59", the crest is centred in a 480 pt block and shrinks to about 36 pt, the raid head subtitle wraps one word per line. Fix: at accessibility sizes stack the three spread stats in a column (number left, label right), top-align the crest at normal size, put "US PvP" on its own line.

9. **Guild crest breaks the one crest language.** Where: parity ios-full renders (Alliance): a bare 36 pt shield with no ring. The mock and the Horde device capture show the ringed circular crest at about 54 pt. Fix: guild crest and faction mark both use the ringed circle at 48 pt on both platforms, ring token as the character crests in Roster.

10. **Pull to refresh: no spinner captured; the layout jumps.** Where: leveler-08 both platforms show the settled result only. "Saved." also inserts a 28 px line under the answers and removes it later, shifting everything below (ios leveler-02 frames vs leveler-08). Fix: reserve the line height, and capture the indicator mid-pull.

11. **Android Calendar leaves the app; iOS pushes a native page.** Where: android leveler-07-calendar-settled.png (Chrome, blank skeleton, meta line "Progression, roster, raid nights and loot.") vs ios leveler-07 (native list). Same row, two results; the Android player loses the tab bar. Fix: push the planned-raids list on Android with the same rows (date tile, title, time, count bar, answer pill).

12. **Tap affordance is inconsistent.** Pages rows and Plan a raid have chevrons; Recent logs rows and the Progression card open pages but show none. Fix: chevron on every row that navigates, none on rows that do not.

### Could

13. Light captures: compared to dark by pixel diff, app frames are identical (the ruling); differences are the system browser and animation timing. No finding.
14. Claim your guild repeats itself: title "Claim your guild", line "Claim it: a guild master or an officer confirms with the addon or Battle.net." Drop "Claim it:". "By name or realm" under Search guilds explains the title; drop it.
15. Close to 60 repeats a member already in Recently on (Thornwick twice, adjacent, ios-full leveler). Exclude names already shown.
16. "Applied Mon" is grey on iOS and cream on Android; the mock shows grey. Use grey on both.
17. iOS Plan a raid subtitle reads "Onyxia's Lair Sat, Barrow Deeps Tue, Onyxia's Lair Sat" over two lines (officer-02). Show the next two raids, as Android does.
18. The recruiting card has an inner double border (parity noGuild) the Find your guild card below lacks; one card frame per screen.
19. The at-60 guild head is a card where spec 2.2 asked for a compact title row "38 members - 31 at 60". The 31 at 60 moved into the Roster title. Acceptable; match the spec when next touched.
20. No loading, empty or error state is captured (guild home failing, no raid planned, no logs). Add three parity renders.

## Interaction walk (spec step list)

- Open the Guild tab: ok on both; spread and raid head lead; finding 1 on the line.
- Answer Late, then In: fail on both (finding 3).
- Roster: ok on both; native page, back to root. The capture leveler-04 is mislabelled (it shows the top of the tab with "Saved.", not the Roster section); re-capture the section.
- Progression: no destination shown on either (finding 4).
- Calendar: iOS ok, Android differs (finding 11).
- Pull to refresh: spinner missing on both (finding 10).
- Officer, the tab: ok on both; Android finding 2. Officer rows sit below Needs you, one scroll down on iOS at 844 pt; acceptable, the badge leads there.
- Officer, Applications: ok on both; applicant sheet with Accept and Decline in a bar above the tab bar.
- Officer, Setup and Plan a raid: iOS error page, Android signed out (finding 4).
- Officer, Assignments, Open in planner: parity only; renders correctly with "Assigning is on the site for now."; destination unproven.
- Your job: parity only; correct ("Whelps, west cave with Kaelra and Thalrin" over "Onyxia - purpose"); left out with no slot.
- No guild: ok on both; the device fixture has no recruiting rows, Apply and "Applied Mon" are in parity only and match the mock.
- Search guilds, Claim: blank at +400 ms on both (finding 4).
- A raid link: ok on both; raid pushed inside Guild, back returns to the root, Guild stays selected.

## Re-review needs

New captures: tab bar from Home and Gear (both platforms); Late and In with the selection moving; signed-in Plan a raid, Setup and Open in planner; the pull spinner; the ax5 leveler head. SHIP when 1 to 5 are closed and 6 and 7 are done or ruled.
