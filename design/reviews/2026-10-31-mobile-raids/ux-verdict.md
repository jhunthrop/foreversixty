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
