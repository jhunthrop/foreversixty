# F4 mobile guild and raids: UX verdict (2026-10-31)

Reviewed: iOS snapshots, Android Roborazzi and emulator captures in this folder, against spec §F4 revision 2 and the f4-* boards. Captures only; no diff reviewed.

## Verdict: HOLD

The sign-up beats a Discord button at default size on both platforms. Two things stop the ship: the moved-to-Bench block and the large-text layouts.

## Blockers

1. **F4.12 moved-to-Bench draws no state (both platforms).** raid-page-moved-and shows Yellfury with every segment unchecked and the line "Yellfury moved to Bench." The board (f4-states F4.12) shows the moved character's Bench segment checked with the 2 pt muted underline. A notice that says Bench over a block that shows nothing is the exact trust failure the line exists to prevent. Fix: the replaced raid view must draw Bench checked on that block; if the fixture lacks the state, fix the fixture. Re-check raid-page-moved-ios for the same fault.
2. **Large text breaks the layout, iOS worst.** raid-page-xl-ios and raids-tab-loaded-xl-ios:
   - Words break mid-word: "PLAN / NED", "OLYMP / US", "Obno / xious", "Chang / e spec", "Barro / w D...", "NOT SIGN ED UP", "Tuesda / y".
   - The count line wraps to four lines and the bar shrinks to a 40 pt stub.
   - The date block collapses into the title column on the tab card.
   - The waiting row wraps its icon off the text.
   Spec ask: nothing breaks inside a word; rows restack. Required at AX5 (iOS) and 2.0 (Android):
   - status chip stays on one line, the guild name drops to its own line;
   - the who-line stacks the crest above the name and Change spec goes below it, full width;
   - the count text sits above a full-width bar;
   - the chip in Also planned rows moves under the title, not beside it;
   - the date block stays 56 pt wide and the text column takes the rest.
   Android 2.0 (raid-page-xl-and) is close: it only wraps "signed / up", the name to two lines and the note placeholder. Same restack applies. Recapture all four xl states.

## Shoulds

1. **The gap under the segments (the lane's flag): fix it.** On the next-raid card (raids-tab-loaded-ios/-and) there is about 60 pt of dead space between the segments and "Who is coming", and another 20 pt under the link. Reserve the status line only on the raid page, where the note field sits under it. On the card, draw the line only when it has text and let the card grow; keep the link at a 44 pt target with 8 pt above it.
2. **Mono is too wide.** The system mono is applied to "pm", to the whole "24 of 40 signed up" sentence and to "in 2 days 2 h". Bundle JetBrains Mono, and set only digits and the unit in mono; keep words ("of", "signed up", "pm") in Barlow. This is also what makes the XL wrapping so bad.
3. **iOS tab top.** The large "Raids" title leaves about 100 pt of empty ground before the guild head. Use the inline title; the next raid card should start within the first 250 pt on a 393 x 852 device. Spec calls for the character selector here (lane F1 gap, noted in the report).
4. **Locked line for officers.** The officer's segments stay enabled but the status line reads "Members cannot change theirs; officers can." The spec's F4.8 sentence is for members; this officer variant is sensible, but put it in the spec copy deck so web and apps share it.
5. **Toast covers the page.** The F4.21 toast sits over the Notes and Sign-ups heading with no inset; lift it 16 pt above the tab bar and give it a surface that hides what is under it.
6. **Android move menu and iOS move menu.** Android capture exists; iOS cannot be snapshotted and has no evidence. Someone must drive it by hand before ship.
7. **Missing evidence the spec asked for:** keyboard focus on a segment, VoiceOver/TalkBack reads (`Saving`, `dimmed`), the calendar sheet, the More menu. Attach before the next review.
8. **Raids badge dot and the 5-tab shell** are lane F1; track them.

## Judged and passing

- **One-tap sign-up:** segments at 48+, optimistic In with the 2 pt kill line; failure restores In and shows "That did not save. Try again." in wipe red (save-failed-ios); "Saved." in kill (moved-and); locked, cancelled race, rate-limited and offline copy match the spec.
- **Spec ask:** pending dashed outline and sheet captures exist (spec-sheet-ask-*, raid-page-spec-pending-ios); Out never asks per the report.
- **Roster with spec icons:** emulator-raid-page-roster-and shows real spec icons, class-coloured names, notes unclipped, own row washed, role groups with counts, Late/Bench/Out/Not signed up groups. Matches F4.2. Snapshot tiles are empty only because the tests have no network.
- **Officer tools:** the Officer row, Lock/Unlock, Copy the list, place buttons with chevrons and Add on Not signed up appear only on the officer captures; member captures show none.
- **Your raid night:** stat figures, loot with quality border and colour, parse-coloured percentile, numbered coach notes; the free capture has no prompt. Passes (loot icon is the no-network tile).
- **States:** offline strip with the timestamp plus disabled segments, cancelled, locked and done captures exist.
- **Parity:** same order, copy and colours on both platforms at default size; the Android bar carries the title and iOS an inline title; both are acceptable.

## Re-review needs

Moved capture on both platforms with Bench drawn; xl captures (tab, raid page, officer view, raid night) on both platforms with no mid-word breaks; the tightened card; the hand-driven iOS move menu.

## Re-check (fix round)

Verdict: HOLD on one new blocker. Both original blockers are resolved.

### Original blockers
1. **Moved to Bench: resolved.** raid-page-moved-ios and -and draw Bench checked with the 2 pt underline on Obnoxious Yell. "Obnoxious Yell moved to Bench." sits under that block. Yellfury shows In with "Saved." in kill green. The Vorgash waiting row is correct.
2. **Large text: resolved.** Checked raid-page-xl, raids-tab-loaded-xl and the spec/officer/raid-night xl captures, iOS AX5 and Android 2.0.
   - No mid-word breaks. Wraps fall only between words ("Pull 7:00 / pm", "signed / up").
   - The status chip stays on one line, with the guild name on its own line.
   - The who-line stacks crest, name, Change spec. The segments become a 2x2 grid.
   - The count sits above a full-width bar.
   - The tab-card date block stacks above the title. The Also planned date block stays 56 wide, with the chip under the title.

### Composition block
- Class names and counts are right. Zeros are muted. All nine classes show in class colours with crests. The moved capture sums to 24 (4+0+3+2+4+4+2+2+3), and the XL capture's 5 Warrior and 2 Hunter is the same roster with Yell In and Yellfury not In.
- Role counts: Tanks 3, Healers 6, DPS 14 match the header line (23 In plus 1 Late). No targets show because the fixture's targets are null, as intended.
- "Missing: Blessings" matches the one dimmed tile, the 4th in the order.

### Blocker (new)
**Buff and utility icons render as seven blank bordered tiles on both platforms.** This includes Android, whose emulator loads the class crests from the CDN in the same capture. The report names only two missing icons (Mark of the Wild, Blessings). So five more are missing or unresolved.
- The row cannot be read. Present buffs cannot be told apart. Only the one muted tile and the "Missing" text carry meaning.
- That is a missing-icon defect on a first-class element.
- Fix: resolve all seven icon names against the export and recapture. Or, until the export lands, draw a short text label ("Fort", "MotW", "AI", "Bless", "SP", "SS", "Reb") inside the tile when the icon fails to load. Recapture on both platforms.

### Shoulds (not blocking)
- Mono is still wide at XL ("11  signed  up", "8/40"). Bundle JetBrains Mono and set only the digits in mono.
- On the iOS tab xl capture, the second Also planned row's date block sits about 12 pt right of the first row's.
- The toast inset is unchanged. Keyboard focus and VoiceOver/TalkBack captures are still owed.
- Everything else from the first pass stands: the More menu, calendar sheet, officer move menu and notification answers now have captures.
