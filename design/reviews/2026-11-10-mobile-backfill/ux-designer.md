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

## Round 4 (2026-11-11, base 693e7a952)

# ux-designer, round 4: Forever Sixty mobile backfill

All three round-3 must-fixes are cleared. Two small new defects, found in the officer evidence you added, block the verdict. A third item is missing evidence.

I judged from the `ios/r4-*` and `android-r4/` captures. I did not open the iOS flow snapshots (run 42938631285 is in flight) or the Android `flows/` steps, so none of the sign-up findings below come from them.

## Round-3 must-fixes

| # | Item | Result |
|---|---|---|
| 1 | One stale tone | PASS on Home and the drawer, both apps. Home's `Synced 49 d ago` (iOS) and `Synced 12 d ago` (Android) draw in the warning orange under Next upgrades. Both drawers draw `Updated … d ago from Battle.net` in the same orange, with `Refreshing…` in the muted tone. Gear is not verified: `android-r4/gear-01-top.png` is the notification shade, not Gear. Recapture it (polish 1). |
| 2 | Covered, no empty frames | PASS. The line reads `Missing [Kings icon] Kings · Might · Wisdom and 2 more` on both apps. Might and Wisdom close up with no frame and no gap. Each name stays on one line. The separator is the same ` · ` on both. The remaining ragged edge, an icon on the first name only, is the known fixture limit. |
| 3 | Notification fixture | PASS. The iOS landing (`r4-push-raid-02-landing`) is the Onyxia's Lair raid page: `Saturday 19 December`, pull 7:00 pm, 24 of 40, Your sign-up with In selected and the Raid tab badge visible. The title, body and landing agree. The Home Molten Core row now reads `Thu 24 Dec`, after the raid. |

## Must fix

1. **The not-synced line has a different tone and affordance on each app.**
   - What shipped: iOS (`r4-officer-raid-02-scrolled`) draws `Not synced this week: Draknar, Skarn, Zekra, Lightningbolltz` in the primary text colour with a trailing chevron. Android (`raid-officer-02-covered-notsynced`) draws the same line in warning red with no chevron.
   - Why it matters: the Android line is the same tap target that opens the filtered roster, and nothing on the line says it is tappable.
   - Fix: match iOS on Android. Draw the chevron at the right edge, centred on the first line, with a 48 dp row. Then pick one tone for the line on both apps. I recommend the primary text colour. The warning tone is already used for stale stamps, and red reads as an error.
   - Files: `RaidTabParts.kt` (Android) and the iOS Raid tab's not-synced row.

2. **The filtered roster does not show that it is filtered.**
   - What shipped: on both apps (`r4-officer-raid-06-roster-settled`, `raid-officer-03-roster-not-synced`) the tap lands on four rows (Draknar, Lightningbolltz, Skarn, Zekra).
   - The stat plate still reads `38 on the roster`.
   - The `All` chip is selected.
   - Nothing says why only four people are shown or how to see the rest.
   - The roster screen is another lane's. Its filtered state is the direct result of a tap on this tab, though, so the brief's rule for main-tab interactions applies.
   - Fix: when the route carries the not-synced filter, draw a chip row item `Not synced this week · 4` in the selected style, with a clear control (`×`, 44 pt / 48 dp), and deselect `All`. Do not keep `All` selected.
   - Files: the shared roster filter state on both apps. The change is the filter chip only, not a roster redesign.

## Evidence still missing (not blockers)

- **iOS officer push results.** `ios/r4-push-short-03-result.png` is Home with the Onyxia's Lair card, not a result frame. It does not show that Mention happened. The Android equivalent (`push-raid-short-03-mention-result`) does: `Mentioned 3 healers in Discord.` / `Onyxia's Lair`. Show the iOS result as the banner or the in-app confirmation. If iOS has neither, say so in the brief and I will accept it from `AppTests`.
- **Drawer Refreshing.** Both in-flight frames are present and correct. `Refreshing…` replaces the link, in the muted tone, with no layout jump.
- **iOS flow snapshots.** Still pending on run 42938631285. Accepted on the condition from round 3: pressed Late, chips, saving, failed with Try again, and offline.

## Polish

1. **Android Gear capture.** `android-r4/gear-01-top.png` shows the notification shade. Recapture Gear so the stale tone on its `synced … ago` line and the `Not modelled` / `+0.0 DPS` cells can be checked.
2. **Android drawer, `drawer-02-refreshing`.** A speckled teal patch sits behind `Refreshing…`. It reads as a leftover ripple or a loading shimmer clipped to the link bounds. The iOS frame shows none. Clip it to the 48 dp link rect, or remove it while disabled.
3. **Drawer stamp clock.** The stamps read 12 d on iOS and 49 d on Android because the fixture clocks differ. This is not a defect, but use one fixture date when you attach both for the same review.
4. **Home meta rows.** The Molten Core row reads `just now` on Android and `4 h ago` on iOS. The same fixture clock gap causes it. Use one clock.
5. **Android roster, `raid-officer-03-roster-not-synced`.** The row name column wraps `Elemental Shaman · 60 · Member` and the nights column stays right-aligned. The first row of the pair is shorter than the others. This is the roster lane's. I mention it only because the filtered view makes it visible.

## Kept as ruled

- Panels on the list blocks.
- Home's Damage without the split.
- The three notification actions on Android.
- The Android tap-through landing on the raid page.

## Checklist

- [x] Stale tone is one rule on Home and the drawer, both apps.
- [ ] Stale tone on Gear (needs the recapture).
- [x] Covered has no empty frames.
- [x] Covered uses the ruled labels.
- [x] Notification title, body and landing agree.
- [x] iOS landing captured.
- [x] `Refreshing…` frames on both apps.
- [x] Android officer push steps and results.
- [ ] iOS officer push result.
- [ ] Not-synced line matches across apps (must-fix 1).
- [ ] Filtered roster shows its filter (must-fix 2).
- [ ] iOS CI flow snapshots attached.

The remaining blockers are the two small must-fixes above. Neither touches the data or the logic. Fix them, recapture Gear on Android and the iOS result frame, and attach the snapshots from the CI run. I will check those frames only.

VERDICT: NO SHIP

## Round 5

Round 5 clears UX1 and UX2 and both round-4 must-fixes, but one round-4 evidence request is still open: no Gear frame shows the stale tone, and Gear's stale tone is verified from code and unit tests only. I rate that a SHOULD, not a BLOCKER. The Android flows, the Android Home, Raid and Activity tab frames, and the iOS flow snapshots were not opened.

**Round-4 items**
- **UX1, not-synced row:** PASS on both apps.
  - iOS `r4-officer-raid-05-not-synced-line.png` and Android `raid-officer-02-covered-notsynced.png` draw the same row: primary text colour, two lines, trailing chevron at the right edge.
  - The Android row is no longer warning red.
  - The Android row is about 100 px tall at 2.7x, so the 48 dp target holds. The iOS row is taller than 44 pt.
- **UX2, filtered roster:** PASS on both apps.
  - iOS `r4-officer-raid-06-roster-settled.png` and Android `raid-officer-03-roster-not-synced.png` show a selected chip `Not synced this week · 4` with `×`, and `All` deselected.
  - The stat reads `4 of 38 on the roster`, and four rows show.
  - The `×` is about 44 px of glyph inside a chip taller than 44 pt / 48 dp, so the target is fine. I am accepting the clear frame (`r5-officer-raid-07-roster-cleared.png`, `raid-officer-04-roster-cleared.png`) from the filenames only. I did not open it.
- **iOS push results:** PASS.
  - `r4-push-short-03-result.png` shows a banner `Mentioned 3 healers in Discord.` with `Onyxia's Lair` under it.
  - `r4-push-applicant-accept-03-result.png` shows `Application accepted` with `New applicant: Mirelle, Holy Priest 60`.
  - This matches Android's `push-raid-short-03-mention-result` and answers round-4 evidence item 1. I did not open the Decline result frame.
- **Android Gear recapture:** the frame is clean and no longer the shade.
  - `android-r5/gear-01-top.png` reconciles. Thirteen cells carry gains and sum to +18.0, which matches `+18.0 DPS from 13 upgrades`.
  - Head, Neck and Main hand read `Equipped`, giving `3 of 17`.
  - Shoulder reads `+0.0 DPS`, as ruled.
  - The iOS Gear frame matches, apart from the Main hand cell, which wraps to three lines.
- **Drawer `Refreshing…` ripple (polish 2):** fixed. `android-r5/drawer-03-refreshing.png` shows the muted `Refreshing…` with no speckled patch.

**BLOCKER**
- None.

**SHOULD**
1. **Gear stale tone has no frame.**
   - I asked for it in round 4 and the brief says it is "kept from round 4", but `android-r4/gear-01-top.png` does not exist in the worktree's design folder.
   - Both Gear captures are fresh (`synced 8 h ago`, muted).
   - The rule is in the code: `GearComposer.kt` `GearHeader.syncedStale` and `GearPresenter.swift:92`. `GearHeaderTest` and `GearParityTest` reference `syncedStale`.
   - I accept it from the code and tests. Attach one stale Gear frame on each app with the next capture.
2. **iOS flow snapshots.** `apps/ios/AppTests/Flows` exists, but the CI record run (42953910502) is still in flight. Every iOS sign-up state (pressed Late, chips, saving, failed, offline) is unverified from renders. Attach the snapshots when they land.
3. **Drawer stamp.** Android `drawer-03-refreshing.png` reads `Updated 12 d ago from Battle.net` (orange) while Gear on the same build reads `synced 8 h ago`. The two stamps disagree on one character, which the fresh-stamp fixture fix (s2) should have prevented. The drawer's stamp comes from the Battle.net import time, so this may be a different field. If it is, say so. If not, move the fixture.
4. **Roster stat plate under the filter.** `4 of 38 on the roster`, then `4 in game, no account` and `2 waiting`, still count the whole roster. After a filter the three figures describe different populations. The roster is another lane's, so I am only flagging it. The first figure is correct, so I am not blocking.

**NIT**
- Android drawer, Notifications row: the value `Raids, gear and sims, logs, guild ladder, tiers, officer on` wraps and strands `on` on its own line. Shorten it, or drop the trailing `on` when every group is on.
- iOS Gear: `Main hand · Equipped` wraps to a third line. Android's cell fits in two lines.
- iOS Raid tab badge reads 6 in the officer state while Android's Gear frame shows 1 for the non-officer state. This is consistent with the shared rule, so no action.
- The officer `r4-*` iOS frames pin the clock to 17 Dec and read `Synced 37 d ago`. This is the one stale capture the brief says it keeps, so not a defect.
- Android's `Preview mention` and the Missing icon are consistent with ruling 5. Nothing to report.

**Checklist**
- [x] Stale tone is one rule on Home and the drawer, both apps.
- [ ] Stale tone on Gear (code and tests only, no frame; SHOULD 1).
- [x] Covered has no empty frames.
- [x] Notification title, body and landing agree.
- [x] iOS landing captured.
- [x] `Refreshing…` frames, both apps, no ripple patch on Android.
- [x] Android officer push steps and results.
- [x] iOS officer push result.
- [x] Not-synced line matches across apps.
- [x] Filtered roster shows its filter, with a clear control.
- [x] Android Gear recaptured and reconciled.
- [ ] iOS CI flow snapshots attached (in flight).

The open items are evidence and fixture polish, not defects in what renders. I keep the round-4 KEEP rulings: panels, Damage without the split, three Android notification actions, and the Android tap-through landing.

VERDICT: SHIP

## Round 6

Round 6 clears all four round-5 items (S1 to S4). I found no blockers, and round 7 is not needed to ship.

I opened 14 captures and 2 iOS flow snapshots, listed in the Evidence limits section. I did not open the other Android goldens (about 40 flow frames) or the rest of the iOS flow snapshots.

**Round-5 items**
- **S1, stale Gear frame: cleared.**
  - Android `gear-01-stale.png` and iOS `r6-gear-01-stale.png` both read `+18.0 DPS from 13 upgrades · synced 9 d ago`.
  - The stale stamp is orange, set apart from the muted fresh tone, and the cell grid is unchanged. Both frames reconcile (13 gains, 3 of 17 slots match).
- **S2, iOS flow snapshots: cleared.** 18 goldens are in `apps/ios/AppTests/Flows/__Snapshots__/FlowSnapshotTests`. They cover home-signup, raid-signup (late, saving, failed, offline), raid-link and segmented-answers, including `-ax5`.
- **S3, drawer stamp: cleared.** Android and iOS both read `Updated 9 h ago from Battle.net`, with `Refresh` in the accent colour. This now agrees with Gear's `synced 9 d ago` in kind, since the two stamps come from different fields (drawer from `bnet_imported_at`, Gear from the character's `synced_at`).
- **S4, filtered roster plate: cleared.**
  - Android `raid-officer-02-roster-filtered.png` and iOS `r6-roster-01-filtered.png` show one plate, `4 of 38` over `on the roster`.
  - The `in game, no account` and `waiting` figures are gone. The selected chip `Not synced this week · 4 ×` is intact, and four rows match the 4.
- **Drawer Notifications value: cleared.**
  - Android reads `All on` and iOS reads `Gear and sims, guild ladder, tiers off`. Both fit one line, which closes the round-5 stranded-`on` nit.
  - The shared copy keys exist in `apps/shared/copy/en.json`.
- **Activity in the house format: cleared.** `204 DPS`, `+12 over your 191 (Fury Warrior, 60).` and `190 DPS` read correctly on both apps.
- **Push failure results: both apps captured, with the copy issues below.**
  - Android and iOS both show a banner titled `That did not save. Try again.`.
  - The Mention failure carries `Onyxia's Lair` under it. The Accept failure carries `Open the application.`.
  - iOS has all four steps (delivered, actions, result, settled).

**BLOCKER**
- None.

**SHOULD**
1. **The Mention failure says "save".**
   - Posting a mention to Discord is not a save. `That did not save. Try again.` is the generic `common.failed` string.
   - Use a mention-specific line such as `Could not mention the healers. Try again.`.
   - It is one copy key, shared across both apps.
2. **The Accept failure drops the applicant.**
   - The success frame names the applicant (`New applicant: Mirelle, Holy Priest 60`). The failure body is `Open the application.` (`push.applicant.failed`), so the notification no longer says who is waiting.
   - Keep the applicant line and put the instruction in a second line, or use `Open Forever Sixty to accept Mirelle.`.
   - `the application` reads as a placeholder to a WoW player, so say the app's name.
3. **The failed notification has no way to retry.**
   - Both apps replace the actions with the failure line, on iOS and Android. Android's frame has an expand chevron but no actions, and iOS's banner has none.
   - Tapping through to the raid or applicant list is the implied retry. State that in the body, or restore the action buttons on the failure.
   - I rate this SHOULD because the player is stuck on a transient failure with no visible next step.
4. **Drawer Notifications states are only partly shown.**
   - Android captures only `All on`. iOS captures only a some-off value.
   - `All off` is not captured, and the brief's `Logs, tiers off` form is not captured on either app.
   - The copy keys exist and I accept the rest from code, but one frame per state would close it.

**NIT**
- iOS `home-signup-05-bench-failed.1.png` shows Late selected after a failed Bench save, with `That did not save. Try again.` and `Try again`. That reads as a revert to the prior answer, which is correct. The filename says bench, so the reader has to infer the revert.
- iOS `raid-signup-08-offline.1.png` shows the segmented control dimmed under an `Offline. Showing what you had at 4:30 pm.` banner and a red `That did not save.` line. The raid page and its header band are another lane's, so I record this only.
- Android Activity has `Today` and `This week` groups and iOS has `This week` only. Molten Core reads `just now` on Android and `Thu` on iOS. Both come from fixture clocks, not app behaviour.
- iOS Gear `Main hand · Equipped` still wraps to three lines. This is a KEEP from round 6.
- The iOS push frames pin the clock to 17 Dec, so Home reads `Synced 37 d ago` in the same frame as a fresh drawer stamp. The brief calls this the one kept stale capture.

**Evidence limits**
- I opened these round-6 captures:
  - Android: `drawer-01-open`, `raid-officer-02-roster-filtered`, `push-raid-short-02-mention-failed`, `push-applicant-02-accept-failed`, `gear-01-stale`, `activity-01-top`.
  - iOS: `r6-menu-01-open`, `r6-roster-01-filtered`, `r6-push-short-fail-03-result`, `r6-push-applicant-accept-fail-03-result`, `r6-gear-01-stale`, `r6-activity-01-top`.
- Of the iOS flow snapshots I opened `home-signup-05-bench-failed` and `raid-signup-08-offline`.
- I did not open `drawer-02-scrolled`, the expanded push frames, the iOS fail-01, fail-02 and fail-04 steps, or the `ax5` and `light` sets.
- I judge the round-6 deltas only.

**Checklist**
- [x] Drawer value copy shipped (`All on`, some-off list)
- [ ] `All off` and `Logs, tiers off` frames (SHOULD 4)
- [x] Filtered roster one figure, both apps
- [x] Drawer stamp fresh and consistent
- [x] Stale Gear frame, both apps
- [x] Activity in the house format
- [x] Push failure frames exist, both apps
- [ ] Push failure copy specific and recoverable (SHOULD 1 to 3)
- [x] iOS flow snapshots attached
- [x] Android flow goldens present

VERDICT: SHIP

## Round 7

Round 7 clears all four round-6 items, with two SHOULD fixes and three NITs. The "Where" line works, but Android and iOS label it differently and the iOS row has no label at all.

**Round-6 items**
- **S1 and S2, push failure copy: cleared on both apps.**
  - Android `push-raid-short-mention-failed.png` reads `Could not mention the healers in Discord` over `Onyxia's Lair. Open the raid to try again.`.
  - Android `push-applicant-accept-failed.png` reads `Could not accept Mirelle` over `Mirelle, Holy Priest 60. Open Forever Sixty to try again.`.
  - iOS `r7-push-short-fail-03-result` and `r7-push-applicant-accept-fail-03-result` carry the same strings. The applicant is named, "save" is gone, and the app is named.
- **S3, retry: cleared.** The body says where to retry, and the tap opens the raid or the applications list. Neither failure shows action buttons, and the body line is the retry cue.
- **S4, drawer Notifications states: cleared on both apps.**
  - Android shows `Gear and sims, guild ladder, tiers off`, `All on` and `All off`. I opened some-off and all-off and took all-on from the file list.
  - iOS shows `All off` and `All on`, and I took its some-off frame from the file list.
  - All states fit one line, with no stranded word.
- **Gear header figure: cleared.** Both apps read `182 DPS` and `+18.0 DPS from 13 upgrades · synced 10 h ago`. The thirteen gains sum to 18.0, and 3 of 17 slots match.

**BLOCKER**
- None.

**SHOULD**
1. **The "Where" line is labelled differently on each app, and not at all on iOS.**
   - Android shows `Crafted: Blacksmithing` with `Where to get it` as a grey caption under it, so the label comes after the value.
   - iOS shows `World drop` in bold beside the item icon, with no caption. It looks like a list row for the item itself rather than a source.
   - The round-7 decision calls this one `Where` line, the way Home's Best run line reads.
   - Fix: use one row pattern on both apps. Either put a `Where` caption above the value, or write `Where · World drop` and `Where · Crafted: Blacksmithing`.
   - The icon is the item's own, so it also matches the `Runner-up` row below it. Drop it from this row or put the caption in front.
2. **The iOS banner truncates the Mention title.**
   - It reads `Could not mention the healers in Di...`, so the word that says what failed is the one cut off. Android shows the full title.
   - Fix: use a shorter `push.short.failed` title, for example `Could not mention the healers`. The body already says Onyxia's Lair, and Discord is implied by the Mention action.

**NIT**
- **`Shoulder · +0.0 DPS` on both apps.** It reads like an upgrade, but the shoulder is not among the 13. Show the equipped-style line, or a `No gain` line.
- **The iOS sheet steps do not show the sheet moving.**
  - `r7-gear-02-sheet-where-t0`, `-t400` and `-settled` all show the sheet at the same height. They differ by at most 2 px.
  - Only `-scrolled` shows the full-height sheet and the `More about this item` link. I judged the settled layout. I cannot judge the presentation timing from these frames.
- **The iOS menu captures are scrolled.**
  - `r7-menu-notifications-all-off` and `-all-on` start mid-drawer.
  - The status-bar clock sits on top of the `Tools` heading and the `Hunter · 31` line shows above it.
  - The drawer header and the Notifications row never share a frame on iOS. This is a capture position problem, since I did not see a layout fault.
- **Fixture differences, no action.** Raid badge `1` on Gear and `6` on Home, and Home reading `Synced 37 d ago` in the iOS push frames (the one kept stale capture).

**Evidence limits**
- I did not open the mocks in `design/mocks/app/` or the Android goldens under `apps/android/app/src/test/screenshots/`. I judged against the round-7 decisions and my round-6 verdict.
- I did not open Android `drawer-02-notifications-all-on`, the iOS fail-01, fail-02 and fail-04 frames for the applicant sequence, `r7-gear-03` beyond what is noted, or `r7-home-01-top`.
- I opened the iOS short-failure sequence only at steps 03 and 04.

**Checklist**
- [x] Gear in the house format on both apps
- [x] Mention failure copy specific, with a stated retry
- [x] Accept failure names the applicant and the app
- [x] Drawer Notifications value in some-off, All on and All off
- [x] Gear sheet shows a source row
- [ ] Source row labelled and identical across apps (SHOULD 1)
- [ ] Banner title survives one iOS line (SHOULD 2)

Files are under `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/`, in `android-r7/` and `ios/` (the `r7-*` files).

VERDICT: SHIP

## Round 8

Round 8 clears both round-7 SHOULD items and the NIT on the zero gain. I found no blockers.

**Round-7 items**
- **S1, the Where row: cleared on both apps.** The row is one bold white line with no caption and no icon, and it matches the mocks' `Onyxia, Onyxia’s Lair, 8%` row. Android reads `Crafted: Blacksmithing (245)`, `Crafted: Blacksmithing (275)` and `Vendor: Lady Palanseer`. iOS reads `Crafted: Blacksmithing (245)`, `Crafted: Blacksmithing (275)` and `World drop`. The skill figure shows where the data carries it. The row can no longer be mistaken for an item row, because the item icon moved to the `Runner-up` row where it belongs.
- **S2, the iOS Mention title: cleared.** `r8-push-short-fail-03-result` shows `Could not mention the healers` on one banner line, over `Onyxia's Lair. Open the raid to try again.`. Android carries the same strings.
- **Zero gain: cleared.** `Shoulder · No gain` shows on both Gear grids.
- **Android runner-up gain: cleared.** The rows read `−5.0 DPS` (Thorium Cestus), `−1.3 DPS` (Serpentskin Leggings) and `Same DPS as the pick` (Blood Guard's Plate Helm). iOS shows the same.
- **iOS drawer from the top: cleared.** `r8-menu-notifications-some-off` shows the header, Switch character and Tools, with `Gear and sims, guild ladder, tiers off` on one line.
- **Decline failure: matches on both apps.** The title reads `Could not decline Mirelle` and the body `Mirelle, Holy Priest 60. Open Forever Sixty to try again.`, with no action buttons, as the KEEP ruling says. iOS fits it in a three-line banner and Android in a two-line shade card.

**BLOCKER**
- None.

**SHOULD**
1. **The crafted runner-up words differ between apps.** In the crafted-upgrade sheet for Dawn's Edge, Android reads `Runner-up · Crafted: Blacksmithing · −5.0 DPS`. iOS reads `Runner-up · Blacksmithing · −5.0 DPS`. The heading above already says `Crafted:`, so I prefer iOS's shorter line. Take it from one shared string (`bis.source.*`) and apply it to both. Round 7's S1 was this same kind of drift.

**NIT**
- **The boss-drop row has no capture.** Neither app shows the `Onyxia, Onyxia's Lair, 8%` heading with its chance, or a boss-drop runner-up. I can only match it to the mock. The brief says the fixture cannot produce it, so there is no action beyond a golden when one exists.
- **The iOS sheet frames still show no motion.** `r8-gear-sheet-*-t0`, `-t400` and `-settled` differ by about 2 px. I cannot judge presentation timing from them, and this is unchanged from round 7.
- **The Android decline body leaves `again.` alone on its second line.** The wording is right. Short of the title and body changing, there is nothing to fix.
- **The sheets cover the two fixtures differently.** Android's vendor-upgrade sheet has no iOS pair, and iOS's world-drop cape sheet has no Android pair. The anatomy matches in the frames I have.

**Evidence limits**
- I did not open `r8-menu-notifications-all-on` or `-all-off`, the iOS Decline steps 01, 02 and 04, the Mention steps 01, 02 and 04, or the Android item-sheet golden under `apps/android/app/src/test/screenshots/parity/`.
- For the mocks I read only the `Onyxia, Onyxia’s Lair, 8%` source row in `android-item-sheet.html` and did not open `ios-item-sheet.html`.
- The brief says the iOS goldens are still being re-recorded, so I judged iOS from the step captures.

**Checklist**
- [x] Source row identical in anatomy on both apps, with no caption and no icon
- [x] Skill figure shown where the data has it
- [x] Android runner-up carries its gain
- [x] `No gain` on Gear cells
- [x] Mention title fits one iOS line
- [x] Decline failure names the applicant and the app, with no action buttons
- [x] iOS drawer frames start from the top
- [ ] Crafted runner-up copy identical across apps (SHOULD 1)
- [ ] Boss-drop row captured (NIT)

Files are under `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/`, in `android-r8/` and `ios/` (the `r8-*` files).

VERDICT: SHIP

## Round 9

Round 9 clears S1: both apps now build the runner-up row from the same shared words. I found no blockers.

I read the Android sheets from `android-r9/gear-02` to `gear-04`. The iOS sheets are `ios/r9-gear-sheet-*-settled`. The Android item-sheet golden is `apps/android/app/src/test/screenshots/parity/android-item-sheet.png`.

**Round-8 S1 (crafted words differing between apps): cleared.**
- Crafted upgrade, Dawn's Edge: both apps read `Crafted: Blacksmithing (275)`, then `Runner-up · −5.0 DPS`. The source is not repeated, and the line is identical on both.
- Crafted helm: both apps read `Runner-up · PvP rank 11 · Blood Guard · Same DPS`. The faction word is dropped as decided, `Same DPS` is a figure-like tail, and the source is named because it differs from the pick's.
- World-drop pick on iOS: `Runner-up · −0.3 DPS` under `World drop`. The source is correctly omitted because it matches.
- Vendor pick on Android: `Runner-up · World drop · −1.3 DPS` under `Vendor: Lady Palanseer`. The source is correctly named because it differs.
- Structure is unchanged and matches the mocks:
  - The runner-up icon is on the left, with the name in its quality colour.
  - The caption is grey, in the mocks' `Runner-up · −1.8 DPS` shape.
  - The source row above it is bold white.
- Spec checks: separators are intact, there is no clipping, there is no layout shift between the t0, t400 and settled captures, and the two apps match in type and spacing.

**BLOCKER:** none.

**SHOULD**
1. **The iOS and Android captures use different pick fixtures.** The third sheet is a world-drop pick on iOS (Dark Phantom Cape) and a vendor pick on Android (Stone Guard's Plate Leggings). Each app shows one of the two naming rules, not both. Because of that, I could not compare the named-source row across the apps in one fixture. Round 10 should capture the same three fixtures on both apps. The code and strings are shared, so I expect no drift.
2. **The Android golden and the mocks show no source difference.** The golden shows a boss pick, `Onyxia, Onyxia's Lair, 8%`, with the runner-up `Vis'kag the Bloodletter · −1.8 DPS`. The mocks match it. That is consistent only if Vis'kag shares that source. If it does not, the rule requires the runner-up to name its own source. Confirm the fixture data. The golden was re-recorded on Nov 10 and shows no named boss runner-up, so the boss shape `Boss · 8%` has no capture in this round.

**NIT**
- iOS sheets are cut at the bottom (the `Ironhide` row is clipped in the helm capture). The Android captures show the whole list with the "3 guildmates have not shared their gear yet" footer. This is presumably the simulator frame, not a defect, so I am not holding the verdict for it.
- The iOS directory in the BRIEF is `ios/r9-*`, and `android-r9/` uses different filenames (`gear-0N-sheet-*`). Name them the same way in round 10.

VERDICT: SHIP

