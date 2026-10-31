# Paperdoll build review: UX designer (2026-10-31)

Reviewed: the side-by-sides (raid 360/1024/1280/2000, beta 360/1280, longest 360/1024), sim-raid, character-beta (360, 1024, 2000), report (360, 1280), loading, failed, tip-right/flip/bind/phone/phoneup, discord-sheet. Judged against design/specs/2026-10-31-paperdoll.md and the earlier verdict.

## Verdict: HOLD

The build matches the mocks closely. Raid 1280/2000, beta 1280, longest 1024/360 and the sheet are at parity. Two defects remain, both on the cases the owner will open first.

## Blockers

1. **Trinket 2 line collides at 1024** (character-beta-1024.png; spec 3.3 "Detail and upgrade line", 3.4 unscored). "Trinket 2 · Item 288893" is overprinted by "Not scored", which wraps to two lines ("Not / scored"). The pick is cut to "Counterattack Lode...". The spec says the pick name takes the ellipsis (`min-width: 0`) and the trailing figure never shrinks. "Not scored" needs `white-space: nowrap; flex: none`, like the gain. Retest every container from 640 to 1100, since C changes at 880 and 1024 sits in the narrow band. The same cause truncates the phone upgrade row at 360 (should 3).
2. **Report draws a broken-image glyph on unknown items** (report-1280.png Chest and Ranged, report-360.png; spec 4.1 "Not in our item data yet", 8 "no missing-icon flash"). The unknown socket must show the game's `inv_misc_questionmark` at 40%, as it does on the character page (character-beta-1280, same fixture build). The report call site loads a different image source for unknown ids. The brief says icons do not resolve in this build, but the question mark does resolve on the character page, so this is a real path difference. Fix it, or attach a report capture from a build with icons that shows the question mark.

## Shoulds

3. **Mid-word ellipsis at 360** (side-beta-360.png, Trinket 2 upgrade row: "Counteratta..." beside "Not scored"). Spec 3.3 says a long word is hyphenated. The row's slot word is 64 px and "Not scored" is wide. Use the 56 px slot word under 340 from the spec and apply `hyphens: auto` to the pick name. The mock drew "Counterattack Lodestone" in full.
4. **Two-hander on the report** (report-1280.png, report-360.png; spec 4.1). Arcanite Reaper is a two-hander, yet the off hand reads "Off hand", not "Two-hand in main hand". Confirm the report payload carries `twoHand`. If it does not, the report is wrong where the character page is right.
5. **Report header contradicts the note** (report-1280.png; spec 3.5). The row above the card reads "Protection · item level 183 · 8 items" while the note says item level shows once every item is known. Rename it (it is not the mean) or drop it when any item is unknown.
6. **Missing acceptance captures** (spec 9, items 2, 6, 7). Not attached: `/sim/<id>`, a focus-visible state (2 px gold outline, offset 2), and a phone tooltip closed by Esc or an outside tap. Attach them.

## Polish

7. tip-bind-1280.png is captioned 1280 but its container is about 860 px, so the tooltip clips the crest and name ("ELGRIM"). §4.2.4 accepts this between 640 and 879. Retake at a true 1280 to show the clear-crest rule.
8. The tooltip source line reads "Crafted: Tailoring" then "Quests" on Lionheart Helm in the fixture. Confirm this is fixture data only.
9. The sheet's off-hand socket is a plain grey square, not the dashed empty treatment. Parity with the site is close enough. Leave it.

## Closed since the spec verdict

- Must-fix A, the 360 hyphen: closed. longest-360.png shows "Electromag-netic", "Shoul-derguard" and "Counterat-tack". No name is cut without a mark.
- The two-hander, the item-level rule (hidden with unknowns on beta, shown 62.0 and 48.1 on full sets, 63.5 on the sheet), "Finger 1/2", the phone upgrade rows, "Not scored" and "Best in slot": all as spec.
- Tooltip placement: the right column aligns end (tip-right), the bottom cell flips above (tip-flip), the phone panel spans the gutters (tip-phone, tip-phoneup). The new "Binds when equipped" and "Plate" lines render (tip-bind).
- Loading keeps row heights and reserves the item-level box. Failed shows plain sockets, text-colour names and the correct note (failed-390).

## Checklist (spec 9)

- [x] One Paperdoll across call sites 1 to 3.
- [x] Slot positions match the sheet.
- [x] Clamp and ellipsis, worst case at 360/390/1024.
- [x] Phone upgrade rows: full width, 44 px, slot word first.
- [x] No gain on a worn pick.
- [x] Item level only when every counted slot is known (sheet agrees).
- [x] Two-hander off hand on the character page and sim. NOT on the report (should 4).
- [x] Tooltip placement (crest clear at 880 and above).
- [ ] No icon flash: blocker 2.
- [ ] No overprint at any width: blocker 1.
- [ ] All acceptance captures: should 6.

Ship when blockers 1 and 2 are fixed and re-captured (character-beta at 1024, report at 1280 and 360). Shoulds 3 to 6 may follow in the same pass.

## Re-check (2026-10-31, after fix round)

Reviewed: character-beta-1024, side-beta-360, report-1280, focus-1280, tip-phoneclosed-360, sim-saved-1280.

### Verdict: HOLD (one blocker, a regression at 360)

### Blocker

1. **Trinket 2 upgrade row clips at 360** (side-beta-360.png, tip-phoneclosed-360.png; spec 3.3 upgrade row, 2-line clamp, "no two slots collapse", must-fix A).
   - The pick reads "Counterat-/tack..." and "Lodestone" is lost.
   - A sliver of a third line (the tops of "Lodestone") shows at the bottom edge of the row. A clamped name must show no stray glyphs.
   - The cause is the same as before: "Not scored" and the 64 px slot word take the name's width.
   - Fix: the 56 px slot word under 340, so the name fits the mock's "Counterattack / Lodestone". Check the line clamp clips cleanly with no third-line fragment.
   - Retake side-beta-360, and also check beta at 390.

### Closed

- Old blocker 1, 1024 overprint: closed. character-beta-1024.png shows "Counterattack Lod..." then "Not scored" on one line. Nothing wraps or overprints.
- Old blocker 2, report glyph: closed. The report sockets are plain, with no broken-image glyph (report-1280.png).
- Should 4, report two-hander: closed. The report off hand reads "Two-hand in main hand".
- Should 5, report item level: closed. The head shows "item level 66" in agreement with the injected gear.
- Should 6, captures: the added focus-1280 shows a 2 px gold outline, offset 2, with the tooltip open. sim-saved-1280 and tip-phoneclosed-360 are present and correct.

### Should (carried)

- Unknown report sockets are plain dark squares, not the 40% question mark. Attach one report capture from a real-data build that shows the question mark, or make the unknown socket use the static `inv_misc_questionmark` asset on the report as on the character page.

### Polish

- sim-saved-1280.png ends with "Run Top Gear to find your next upgrade. Run Top Gear" flush left, below the card edge, in body type. Confirm this is a crop, not the layout.

Ship when blocker 1 is fixed and side-beta-360 is re-shot. No further review round is needed if the retake shows "Counterattack" whole or cleanly hyphenated, and no clipped fragment.

## Re-check 2 (2026-10-31)

Reviewed: side-beta-360, character-beta-390, report-360.

### Verdict: SHIP

- Blocker 1, the 360 Trinket 2 row: closed. side-beta-360.png reads "Counterat- / tack / Lodestone" on three clean lines beside "Not scored". Nothing is clipped and no third-line fragment shows. The name is whole, with a hyphen mark, and still distinct from every other row. At 390 (character-beta-390.png) it reads "Counterattack / Lodestone" on two lines.
- Carried should, the report question mark: closed. report-360.png shows the 40% desaturated question mark in the Chest and Ranged unknown sockets. Their names read "Chest · Item 288892" and "Ranged · Item 288893".
- Polish, the trailing line in sim-saved: accepted as page content in the crop.

### Remaining polish (non-blocking)

- The hyphenation at 360 ("Counterat-/tack") is a legal break but ugly. The mock fit "Counterattack" on one line. Tightening the row (a 52 px slot word, or "Not scored" at 12 px) would restore it. Do this when next in the file. It does not hold the merge.
- Retake tip-bind at a true 1280 for the file record (polish 7).
