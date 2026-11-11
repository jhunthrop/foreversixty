# ux-designer: mobile backfill (changed surfaces, both platforms)

## Round 2 (2026-11-10, base 9bbf15e37)

iOS and Android are not at parity on the first job, and the evidence for several interaction steps is missing or contradicts the claims. I judged from the captures and a grep of `apps/ios/Sources`. I did not run the apps, and the iOS flow snapshots have not landed from CI.

## Must fix

1. **The role tally disagrees with itself on Android.**
   - Home's raid card reads `Damage 14`. The Raid tab, with the same fixture raid, reads `Damage 15 · 6 melee, 8 ranged`, which adds up to 14.
   - iOS reads `Damage 15 · 7 melee, 8 ranged`, which adds up to 15.
   - On raid day this is the first answer the player reads, so a wrong count is a blocker.
   - Fix the melee/ranged and total derivation in the one shared tally model, then add a parity test that asserts melee + ranged = total.
   - Files: Android `home/flows/HomeRaidCard.kt`, `raids/flows/RaidTabParts.kt`; iOS Raid tally view.

2. **iOS Home's raid card has no role tally, though round-2-decisions.md accepted it.**
   - iOS shows `Onyxia's Lair · Sat 19 Dec, 7:00 pm CST` and `You are In`, with no tally line.
   - Android has `Tanks 3 · Healers 6 · Damage 14`.
   - Fix: add the shared tally component to the iOS Home card.
   - File: iOS Home raid card (`apps/ios/Sources/Features/Home`).

3. **Android Home breaks the ruled time buckets.**
   - Home's `Molten Core posted` row reads `· 3 d ago`, and the `ago` wraps onto a second line.
   - Activity shows the same row as `Sat`, and iOS Home shows `Sat`.
   - The ruling is one row, one reading, on both apps.
   - Fix: use the shared bucket formatter on Home rows, and keep the stamp from wrapping (non-breaking space, or on its own line).
   - File: Android `home/flows/HomeFlowScreen.kt`.

4. **iOS has no flow evidence for the sign-up answer control.**
   - Every iOS step capture (`raid-signup-02-late-t0`, `-t400`, `-settled`, `home-signup-01-late-pressed`) is the same frame: In selected, `Saved.`.
   - Because t0 already shows the final state, I cannot see selection feedback, the Late minutes chips, `Saving…`, the red error with Try again, or the offline state.
   - The fixture snapping back to In explains the frames. It does not excuse missing proof for the new Segmented and the new states.
   - Fix: attach the CI flow snapshots, or capture against an echoing service. Required frames: pressed Late, chips, saving, failed with Try again, and offline with dimmed segments.
   - Also confirm the haptic fires on press and not on settle.
   - Android has the full set, and those pass (see Passing).

5. **Pull to refresh and the drawer's Refresh show no feedback in any capture.**
   - `home-03-pull-refresh`, `raid-09/06-pull-refresh` and `activity-05-pull-refresh` look identical to the rest state, with no indicator.
   - Drawer `menu-02-refresh-t0` and `-t400` (iOS) and `drawer-03-refresh-t300` (Android) are unchanged, and the stamp still reads `Updated 12 d ago`.
   - Feedback timing cannot be judged, and the spec requires visible feedback in the first 100 ms.
   - Fix: capture while the indicator is showing, and show an in-flight state on the drawer's Refresh (`Refreshing…`, or a disabled link) so the stamp visibly changes.
   - Files: `ios/menu-02-refresh-*`, `android-r2/drawer-03-refresh-t300.png`.

6. **iOS Activity has no "Show older".**
   - Android ends the list with `Show older`.
   - The iOS capture ends after four rows with empty space, and my source grep found no equivalent in `apps/ios/Sources`.
   - Anyone past the first page cannot reach older items on iOS.
   - Fix: add the control with the same copy and a 44 pt target.

7. **The Android push tap-through capture is a black screen.**
   - `android-r2/notif-02-expanded.png` is a black screen with a spinner, not an expanded notification.
   - The iOS `notif-04-tap-through` is the lock screen again, so it never shows the app opening.
   - Fix: recapture both as the notification expanded, then the landing screen.
   - The landing should be the Raid tab with the sign-up visible. Back should return to the tab (Android `raid-link-03` already does this).

8. **The Android 2x header squeeze, already known.**
   - At font scale 2, `Gear ›` renders as a vertical column of single letters.
   - Past that, the first upgrade starts at about 62% of the screen height.
   - Fix: wrap the section header into two rows (title, then meta and `Gear ›`) at large text.
   - iOS ax5 already does this.
   - File: `home/flows/HomeFlowScreen.kt`.

## Polish

- **Orphan stamp wraps.**
  - iOS Home: `· Sat` lands alone on its own line, and `Sign up In, Late or Out. · Sat` breaks the same way.
  - Fix: put the age on its own meta line without the leading dot.
- **Synced stamp placement differs.**
  - iOS puts `Synced 12 d ago` above the first upgrade row.
  - Android puts it below the last row.
  - The brief puts the stamp under the meta, so move it to the same place on both. The top suits the player's first read.
- **Stale warning tone.**
  - A 12-day-old sync reads as plain grey text.
  - Past a staleness threshold it should take the warning tone, with a hint to refresh from the drawer.
- **Gear's first-screen answer is the same on both platforms (the header line plus the slot grid).**
  - The iOS grid starts about 440 px lower because of the large title.
  - That follows platform convention. Keep it.
- **Bring flasks placement.**
  - iOS sets it inside the head card.
  - Android sets it outside the card as loose text under the segmented control.
  - Pick one, and I recommend inside the card.
- **Missing-buffs line.**
  - Android renders it in red and shows every missing buff, but it reads as a wall of `Missing …` repeated six times.
  - iOS keeps it neutral and clips with `and 2 more`.
  - Align both: neutral label, red only on the missing names, one `Missing` prefix, three names and `and N more`.
- **Android fs2 Raid.**
  - `Bench` touches the segment dividers at 2x.
  - Stack the segments at large text, as iOS does.
  - The `In` chip drops below the date. That is acceptable.
- **Android Late minutes.**
  - The Raid tab shows the 15/30/60 chips with no label.
  - The Home flow shows `How late?` above them.
  - Use the label in both.
- **Android Saving copy.** Use the ellipsis character (`Saving…`), not three dots.
- **Offline tap.**
  - `raid-signup-11-offline-tap` is identical to `-10`, so tapping a disabled segment gives no feedback.
  - The persistent explanation line covers it.
  - Add a light error haptic to match R14.
- **Android drawer at 2x.**
  - `Refresh` is an inline text link inside a wrapped sentence, so its touch target is under 48 dp.
  - Move it to its own row at large text, and do the same on iOS ax5.
- **Notification fixture.**
  - The title says `Onyxia's Lair`.
  - The body says `Sign-ups are open: Thu 5 Nov` (Molten Core).
  - The collapsed Android line truncates before the count, `24 of 40 signed`.
  - Fix the fixture, and put the count before the date, or cut the date to `Sat 19 Dec, 7 pm`.
- **Tab bar overlap on iOS.**
  - The floating bar overlaps content with no fade or inset.
  - The Raid captures show text running under the bar.
  - Add bottom content inset and a scrim.
- **Activity filter chips.**
  - They are clipped at the right edge (`Guil`), which does signal scrolling.
  - Keep it, and add a trailing fade.
- **Gear counter copy.**
  - `0 of 17 slots match the list` is correct.
  - It reads as a failure when the list is an old class-band list.
  - The `Change` link beside it covers it.
- **Drawer.** The head, the rows and the plan and device values match the decisions, and the head's portrait fallback never shows an empty ring.

## Passing

- Raid-day answer: the Raid tab shows the sign-up control with the current answer in the first 560 px on both apps. `Roster ›` and Who is coming sit directly below, and Needs you is not crowding it.
- Non-raid-day: Home leads with Next upgrades, and the upgrade meta fits on one line at 1x. The Home raid card stays one row, with the tally as its only added line.
- Android sign-up flow:
  - Feedback is tick, then optimistic selection, then `Saving…`, then `Saved.` in green.
  - A failure restores In and shows `That did not save. Try again.` with a Try again button.
  - Offline dims the segments and says `You are offline. Sign-ups need a connection.`
  - Back from the raid page keeps the tab's scroll position (`raid-05-back-keeps-scroll`).
- Segmented on iOS matches Android and the web: one group, chosen segment teal, hairline dividers. It stacks at ax5 with rows of roughly 53 pt.
- The `Covered`, `Bench · Out · not answered`, and `7 melee, 8 ranged` lines match the accepted decisions.

VERDICT: NO SHIP

The blockers are items 1 to 3, which are wrong or contradictory numbers on the first job. Items 4 to 7 are missing proof for steps the brief asked me to judge, and items 6 and 8 are functional gaps. After the fixes, re-run the iOS flows and the pull and drawer captures and send them back.

## Round 3 (2026-11-10, base 7911bf87e)

# ux-designer round 3: mobile backfill (Home, Raid, Gear, Activity, drawer, notification; iOS and Android)

Base is the round-3 tip. I judged from the captures in `ios/` and `android-r3/`, the step logic in `apps/ios/AppTests/Flows`, and a grep of the drawer code. I did not run the apps. The iOS CI flow snapshots are not in the record yet.

## Round-2 must-fixes, checked

| # | Item | Result |
|---|---|---|
| 1 | One count model | PASS. Home reads `Damage 15` on both apps. The Raid tab reads `Damage 15 · 7 melee, 8 ranged` on both. 7 + 8 = 15, and 3 + 6 + 15 = 24 matches the raid page's 24 of 40. |
| 2 | iOS Home tally | PASS. `Tanks 3 · Healers 6 · Damage 15` now sits above the Segmented, as on Android. |
| 3 | Home time buckets | PASS. Both Home and Activity read `Sat` on its own meta line on both apps. No `· 3 d ago` and no orphan dot. |
| 4 | iOS sign-up evidence | CONDITIONAL, see below. |
| 5 | Drawer Refresh | Code PASS, capture still missing, see below. Pull to refresh uses the system controls, which is accepted. |
| 6 | iOS Show older | PASS. `activity-07-older-settled` shows a centred yellow `Show older` link below the list. |
| 7 | Push tap-through | PASS on Android: `notif-02-expanded` is the raid page, Back returns to the tab, and the count agrees. The iOS capture is still a lock screen, see must-fix 3. |
| 8 | Android 2x header | PASS. In `fs2-home`, `Gear ›` stays on the title row and the meta and stamp drop to their own rows. |

The i3 bottom inset also passes. `raid-06-scrolled` ends with clear space above the tab bar.

Sign-up flow (item 4) is judged from the tests, not from snapshots. `RaidSignupFlowTests` asserts these behaviours:
- Late selected and chips open on press.
- `busy` while the write is in flight, with the press drawn before the write returns.
- On failure, a return to In, `Copy.commonFailed`, a retry, and haptics `["selection","error"]`.
- On success, haptics `["selection","success"]`.
- Offline disables the answers and shows the sentence. A tap sends nothing.

That covers the logic and the haptic order. The steps match Android's `flows/*.png`. I accept item 4 on the condition that the CI snapshots are attached before merge. The required frames are pressed Late, chips, saving, failed with Try again, and offline. I will not re-open it unless a frame contradicts the tests.

## Must fix

1. **The stale-sync tone differs between the apps and inside iOS.**
   - Home: Android draws `Synced 12 d ago` in the warning orange. iOS draws it in plain grey (`home-01-top`, `ax5-home`).
   - The iOS drawer draws `Updated 12 d ago` in orange.
   - Round-3 change 12 asked for the warning tone past 7 days, with Refresh as the action.
   - Fix: apply the warning tone to the iOS Home and Gear stamps. Android's Home treatment is the reference.
   - Gear's `synced 12 d ago` is grey on both apps. Either tone it on both or on neither, to match Home.
   - Files: the iOS Home section-header stamp, and iOS and Android Gear's sync line.

2. **The Covered/Missing line draws empty icon boxes and does not use the ruled names.**
   - iOS draws `Blessing of Kings` with its icon. `Blessing of Might` and `Blessing of Wisdom` each get an empty dark rounded square (`raid-01-top`).
   - Android draws the Kings icon and leaves a blank gap before Might and Wisdom, so the baseline is ragged (`raid-01-top`).
   - A missing-icon flash is a defect under the performance limits. Showing an icon only for the first name is also a defect.
   - Round-3 change 4 asked for buff short names from the data's `label` (`Fort`, `MotW`, `Kings`) with spell icons beside them. The captures show full names instead.
   - Fix:
     - Resolve every spell icon, or omit the icon box when there is none. Never draw an empty frame.
     - Use `label` for the names.
     - Keep the order missing before short, at most three names, then `and N more`.
     - Do not wrap mid-name: Android wraps `Blessing of Wisdom` onto a second line under the gap.
   - Files: the iOS Raid tab's Covered view and Android `raids/flows/RaidTabParts.kt`.

3. **The notification fixture still contradicts itself, and iOS has no tap-through frame.**
   - iOS `notif-04-tap-through` shows the title `Onyxia's Lair` with the body `24 of 40 signed. Thu 5 Nov, 8:00 pm CST.` Round 2 flagged this mismatch. The raid is Sat 19 Dec, 7:00 pm.
   - The same date confusion shows in the `Molten Core` feed row, `Thu 12 Nov`.
   - The file is named tap-through but it is the lock screen. The landing is not captured.
   - Fix: set `docs/contracts/fixtures/app/push.raid-posted.json` so the body carries the same raid as the title: `24 of 40 signed. Sat 19 Dec, 7:00 pm CST.`
   - Recapture the iOS landing on the Raid tab with the sign-up visible.

## Evidence to attach before merge (not blockers by themselves)

- **Drawer Refresh in flight.** `menu-02-refresh-t0` and `drawer-03-refresh-t300` still show `Refresh`. Neither shows `Refreshing…`.
  - The code is present: `ProfileMenu.swift:211` and `menu_refreshing` in `ProfileMenu.kt:182`.
  - I accept it from the code. Capture the in-flight frame with the link disabled, then the new stamp, for both apps.
- **iOS flow snapshots.** Attach them from the CI record run (see item 4 above).

## Polish

- **`drawer-03-refresh-t300` (Android).** A teal rectangle sits behind the `Updated 12 d ago from Battle.net · Refresh` line. It looks like a pressed or focus fill on the whole row, not on the link.
  - Confine the pressed state to the `Refresh` link, with a 48 dp target.
  - Files: `ProfileMenu.kt`, and `ProfileMenu.swift` for the matching target.
- **Hero line at large text.**
  - On Android fs2, `Level 58 Fury Warrior · 2 levels to the 60 list` wraps with `· 2` stranded at the end of the first line.
  - On iOS ax5 it takes four lines and the Next upgrades header spreads over five rows. That pushes the first upgrade below the fold at about 2300 px.
  - Break the line at the dot, with `2 levels to the 60 list` as a second line. At ax5, collapse the band to the name only.
- **Gear counter on iOS.** The fixture app still reads `0 of 17 slots match the list` with no Equipped cells, while Android reads `3 of 17`. The brief lists this as known and being fixed. The iOS capture predates the alignment.
- **Gear on Android.** The counts reconcile. `13 upgrades · +63.0 DPS` equals the cell gains, which sum to 63.0, and 3 slots read Equipped. Shoulder carries neither a gain nor Equipped. That is acceptable as long as a no-gain, not-equipped slot is intended to read plain.
- **Activity filter chips on iOS.** `Guil` is still clipped at the right edge with no trailing fade. Round 2 asked for the fade.
- **Android Show older.** The control is already ruled. Confirm its 48 dp height.

## Passing

- **Segmented.** Teal fill and hairline dividers match on both apps. The raid notes (`Bring flasks`) sit inside the card on both. There is one answer state per surface, with no chip beside the title.
- **Raid-day first screen.** The answer, the tally and `Roster ›` all sit inside the first 560 px on both apps.
- **Activity rows.** The figure shows once, as `1,196 DPS`. Rows use the same bucket format as Home.
- **Android Gear.** Slot labels are `Finger 1` and `Trinket 2`, and Equipped cells show.

VERDICT: NO SHIP

The blockers are must-fixes 1 to 3. All three are small. Fix them, attach the CI snapshots and the in-flight drawer frames, and this ships.

