# Guild tab: ux-designer verdict, round 2 (2026-11-11)

Reviewed on screenshots: ios/dark and android/dark tabbar-01-home, -02-gear, -03-guild; officer-02 (Android); parity/ios-full-guild-tab-leveler.png; parity/leveler.ios-guild-tab-leveler-ax5.png; round1-response.md. Contrast figures are sampled from the pixels.

## Verdict: NOT YET

One must stands: the Late answer is still never seen to select. Everything else from round 1 is closed or downgraded. When one re-capture shows Late and In selecting, this ships with the shoulds below as follow-ups.

## Closed since round 1
- Round-1 finding 1, quiet line: sampled 5.15:1 on iOS and 5.18:1 on Android (227,223,217 on the tan). Closed.
- Round-1 finding 2, "Holy Paladin Paladin": Android officer-02 now reads "Test Applicant, Holy Paladin 60" in Needs you and in Applications. Closed.
- Round-1 finding 5, unselected Guild icon: captured on both platforms from Home and from Gear. Closed.
- Round-1 finding 6, icon weight: judged on the captures. Ruled acceptable, see Could 6.
- Round-1 finding 9, crest: the Alliance render now shows the ringed circular crest. Closed.
- Round-1 findings 14, 15, 16 and the Roster link shown twice: closed in the renders.

## Tab bar (Home and Gear with Guild unselected, Guild selected)
- iOS: unselected Guild is cream on the dark pill, on the same tint path as Home, Gear and Activity. Selected Guild is gold on the lighter pill, the same treatment as Home and Gear selected. The banner is a solid silhouette beside the outlined helm and horn. It carries a little more ink but sits at the same visual size and baseline, and the badge sits at the Activity offset. Reads as one family.
- Android: unselected Guild is muted blue-grey, the same as the other three. Selected Guild sits on the green-grey pill with gold glyph and label, matching Home and Gear. Weight difference is small at this size.
- The badge covers the banner's right finial on both platforms. It is legible as a banner and a count. Accepted.
- Pass.

## Findings

### Must
1. **Late and In never visibly select.**
   - Where: leveler-02 and leveler-03 on both platforms, as in round 1.
   - Ruled by the lane: the fixture answers one canned raid, so a tap snaps back to In, and the flows tests cover the real path.
   - Not accepted. The control's one job is the answer, and no frame shows the selection moving. A reviewer cannot judge a test.
   - Fix: make the fixture echo the tapped answer. Re-capture Late, then In, on both platforms. The tapped segment is filled at tap, and the counts under Who is coming change with it.

### Should
2. **Under 60, Who is coming still leads and carries the 60-only panel.**
   - Where: parity/ios-full-guild-tab-leveler.png. Tanks, Healers "5 of 8 - 3 short", Damage "11 of 30 - 19 short", Bench/Out line and Missing Fort/WF fill the screen below the raid head. Roster starts about 1.5 screens down.
   - Not addressed in round1-response.md; only the duplicate Roster link was removed.
   - Fix: under 60 show the counts line and the roles line only. Keep the full panel for 60.
3. **Site hand-offs: officer destinations still unverified, and the row does not say it leaves the app.**
   - Where: ios officer-04 (setup did not load), android officer-05 (signed out), the search and claim captures.
   - Ruled by the lane: the system browser carries no app session. Accepted as a platform fact, not as a must.
   - Fix: give Plan a raid, Setup, Search guilds and Open in planner an external-link glyph instead of a chevron, so the sign-in page is expected. Add one capture of each destination, signed in by hand, to the review folder. Raise a handover link (one-time sign-in) with the owner as a later item.
4. **Android Calendar leaves the app; iOS pushes a native page.**
   - Ruled out of scope by the lane.
   - Same row, two results, and the Android player loses the tab bar.
   - Fix: push the planned-raids list on Android, or record it as an open item with an owner.
5. **Android rows that navigate carry no chevron.**
   - Where: officer-02, the Officer card (Plan a raid, Applications, Setup). The response says Recent logs and Progression chevrons landed on iOS only.
   - Fix: the same chevron on every Android row that navigates, none on rows that do not (Needs you's Ping on Discord is a button).
6. **"Saved." inserts a line under the answers and shifts everything below.** Not addressed. Reserve the line height. Capture the pull indicator mid-pull; leveler-08 on both platforms still shows the settled result only.

### Could
7. **Icon weight.** The solid banner is the site's own Guild door, as the lane rules. Read on the captures it is acceptable. Flag to the owner as the lane proposes; no change required.
8. **ax5 head** (parity/leveler.ios-guild-tab-leveler-ax5.png). The spread now stacks as a column, which is correct. Still: the crest ring is about 25 pt against a two-line name, "US / PvP" splits across lines, and the subtitle takes three lines. Put "US PvP" on its own line and hold the crest at normal size.
9. **Ping on Discord** is a filled blue button on Android and sits alone in Needs you. Check it against the Home and web planner button once, so it reads as the same control.
10. No loading, empty or error state of the guild home is captured. Add three parity renders when next touched.

## Checklist from round 1
- Quiet line contrast: ticked.
- Applicant line formatter: ticked.
- Answer buttons show the answer: NOT ticked (1).
- Site hand-offs: downgraded to should (3).
- Tab bar unselected, selected, both platforms: ticked.
- Roster action once: ticked.
- Guild-60 panel under 60: not ticked (2).
- ax5 head: mostly ticked (8).
- Crest language: ticked.
- Tap affordance: iOS ticked, Android not (5).
- Pull to refresh and layout jump: not ticked (6).
