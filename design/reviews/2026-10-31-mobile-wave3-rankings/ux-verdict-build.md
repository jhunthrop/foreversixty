# F11 build review, round 1: rankings, character pages, Paperdoll, Home Your gear

Captures: design/reviews/2026-10-31-mobile-wave3-rankings/build/. Boards: design/mocks/2026-10-31-mobile-wave3/renders/f11-rankings-ios.png and -and.png (only the iOS board was compared pixel for pixel). Spec: design/specs/2026-10-31-mobile-wave3.md section 5.3, sections 4, 6 to 9.

## Verdict: HOLD

The layout, copy, colours and states are close to the spec. Five defects stop it.

What matches: answer card, filter line plus 44 pt Filters button, board rows with ladder-coloured ranks, the pin under `Your best kill` with a gold edge, `You` tag, `Jump to your row, 14th`, empty, error, loading skeleton, offline strip, not found, no logs, no gear. Score card, five bars, Bests with percentile colours, Recent fights. Paperdoll: two mirrored columns, gold dots, legend line, `Gear gap` row, one column at largest text. Copy is verbatim everywhere I checked.

## MUST (blockers)

1. **iOS character page has no Watch star and no share in the bar.** character-other-ios, character-yours-ios, character-gear-ios, character-nologs-ios, character-large-ios show a bare back chevron. The board F11.4 and the spec put a 44 pt star toggle and share in the bar. Android shows both (character-other-and). The three-state strip character-watch-ios draws them inside a card, which does not prove the bar. Change: render the star and share in the iOS navigation bar and recapture the iOS pages with them visible, plus the watch on and off states in the real bar.
2. **Android filter sheet is covered by the bottom tab bar.** filters-and (and filters-encounters-and): the sheet runs under the nav bar, and `Today only` and `Reset filters` are hidden behind it. The sheet must sit above the tab bar, or end above it, and scroll to the last control. Recapture scrolled to the bottom.
3. **iOS filter sheet controls are broken.** rankings-filters-ios: every chip has stray tick marks on its border (left and right on pills, top and bottom on circles). The `Today only` switch is a blank grey pill with no thumb, and it overshoots the right margin (about 888 px against the 884 px content edge). Fix the chip border drawing and use a real switch with an on and an off state. Capture both.
4. **iOS encounter list sheet has a dead top half.** rankings-encounters-ios: the handle floats mid-screen and about 40 percent of the sheet is empty above the list. Anchor content to the top of the sheet at its detent.
5. **Raw string in the off-hand slot.** home-gear-ios and home-gear-and: the empty off-hand reads `none: Legionite Glaive is two-handed.` (lowercase, internal wording, wrapping to four lines on iOS). Replace with a player sentence (suggest `Two-handed weapon equipped.`) or leave the slot empty with `Off hand`. This is separate from the owner's ruling on unknown items. Also remove the gold dot from an unknown or empty slot unless a BiS item exists for it (Wrist and Trinket 2 carry dots).

## SHOULD

1. **Answer numeral colour differs by platform.** rankings-ios and rankings-offline-ios paint `14th` purple (parse ladder); rankings-onpage-and paints it white; the board shows white. Decide one (spec: stat size, no colour rule) and match both. Off the page, 204th is grey on iOS and white on Android.
2. **Android unranked card drops the `Your rank` micro label.** rankings-empty-and versus rankings-empty-ios. Add it.
3. **Off-page answer card differs.** rankings-offpage-ios shows `Sparkwick · best kill on...` with a gold name; rankings-offpage-and shows `Your best kill on Onyxia, Normal, Phase 1.` Spec `rank.meta` is the Android text. Make iOS match.
4. **Age wraps to an orphan.** Row meta breaks as `2 d / ago` on Android (rankings-onpage-and, rankings-offpage-and, rankings-signedout-and) and `· 6 d ago` starts a line on iOS. Bind `{n} d ago` with non-breaking spaces and keep the separator on the previous line.
5. **Pin row edge.** rankings-offpage-ios and -and: the gold edge is a straight bar poking outside the rounded corner. Clip it to the card radius.
6. **Score value font on Android.** character-other-and sets `1,830` in the condensed display face; spec and iOS say mono 36. Same for the `Builds seen` totals (spec: mono). iOS also adds an icon tile that Android lacks.
7. **Dead space in the score card.** About 40 pt under `How is this calculated?` on both platforms (character-other-ios, character-other-and). Trim to the card padding.
8. **Weapons row differs.** home-gear-ios left-aligns three columns and wraps names to four lines; home-gear-and centres a column per weapon with the icon above. Use the Android arrangement on iOS (the spec's phone frame says weapons in a row under the columns).
9. **Android character bar repeats the name.** The title `VOLTARIS` sits directly above the head name. Leave the bar title empty until the head scrolls away.
10. **Largest text.** rankings-large-ios and character-large-ios reflow correctly (value under meta, one column of sockets). The Filters button falls under the filter line; that is acceptable. Attach the Android largest-text rankings capture (only rankings-onpage-large-and exists; I did not see the filter line at that size).

## Not reviewed or missing

- iOS error and character-error captures, loading for character pages, `character-watch-failed`: not opened in this round; confirm the copy.
- No capture of the long-press menu (`Watch`, `Open the report`, `Share`) or of keyboard/switch focus. Attach for round 2.
- Android rankings largest text with the filter line and the sheet at largest text.

Round 2 needs: MUST 1 to 5 fixed, recaptures named above, and SHOULD 1 to 4 decided.
