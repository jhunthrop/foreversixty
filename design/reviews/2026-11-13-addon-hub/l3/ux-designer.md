# L3 Gear tab: ux-designer

## Round 1: FIX

Reviewed what renders: `gear-hunter-60.png` and `gear-warrior-34.png` against `hub-gear.png`, same scale (740 px, device scale 2), plus the spec 3.5, 3.7, 4.3, 5, 11. Region by region the tab matches the board: header figures line, tab strip with the Gear dot, Level row, two-column doll with outer icons, Upgrades (icon, name, gain, slot and source, total note), weights table, Sources in two columns with counts and the state-word note, footer. Header counts check out against the cells (hunter 9 best in slot, warrior 5; empty slots not counted). The doll never needs the board's Conditions row, so it starts 56 px higher and stays whole.

### Known departures, judged
1. No Conditions row: accepted. One list per band, nothing to choose; a dead control would be worse.
2. No error column or haste error: accepted. The table's three columns (Stat, Weight, DPS) hold their widths and right edge (1506 px at 2x) on both renders.
3. Drop chance only where the data has one: accepted. Warrior row 1 (`World · 1% drop`) reads correctly; rows without one are not padded.
4. `Beats the pick` muted: accepted. It is the quiet state and keeps the row from reading as a gap.
5. Source line ellipsis: accepted in principle, see should 2 for the case where it fails.

### Must
1. Section headings. Spec 3.5 and the board set them semibold sans in the strong colour (Marcellus only for the window title). Hunter and warrior renders draw `Upgrades`, `Stat weights`, `Sources` in the serif title face, cream. The Builds tab (`l1/window-hunter-42-hub.png`) does the same, so L3 is consistent with L1 but both contradict the spec. One decision, one change in `BuildsWindow.buildSection`, for every tab: either the headings go semibold sans per spec, or the spec and board are amended to the serif and the Rotation, Guild and Officer boards follow. Not shippable with the spec and the build disagreeing; do not fix on this tab alone.
2. Source checkboxes are not evidence. `design/mocks/addon_hub_l3.py` draws them by hand (16 px solid gold square with a dark tick; unchecked a 14 px muted outline), so the renders show 14 loud gold blocks that the board does not have (board: outlined box, gold tick, one level of emphasis). Attach a capture, or a mock built from the real `Widgets.toggle` recipe, that shows the checked state as built. If the real widget does fill gold, it fails the board: the Sources block becomes the loudest region in the window, louder than the gains. Acceptance: checked and unchecked, hover and keyboard focus ring, at the real size.

### Should
1. Weights rows. Hunter shows `Attack Power` and `Ranged Attack Power` both 0.33 / 0.14, then `Strength 0.07` and `Intellect 0.03`: seven rows, two of them duplicates or noise. The board lists five that matter. Collapse an identical pair into one row (`Ranged attack power`), or cut rows under 0.10 weight, so the table reads as a ranking. Also confirm whether the shipped hunter weights carry no haste; the board shows a haste line and the render shows none.
2. Source line truncation (warrior). `Boulderfist Enforc…`, `Scarlet Mon…`, `Razorfen Kr…`, and on row 3 the boss itself (`Charlga Razorfla…`) is cut at the column edge. The rule in departure 5 keeps the boss; here it does not. Cut the place to its floor first, then the kind, and keep the boss whole; if the boss alone overflows, put the full source line in the row's hover (spec 3.5: cut cells are shown whole on hover). Say which in the render notes.
3. Empty doll cell (warrior Hands, Finger, Trinket x2). A plain black square. The character pane the doll copies draws the slot silhouette there; add it at the muted quality-edge ground so an empty cell reads as a slot, not a missing icon.
4. Doll name column. `Brawler's Leather Leggua…` is cut with 40+ px of gutter free between the columns. Give each column the gutter, or the longer of the two names, before ellipsis.
5. Vocabulary in source lines. The tail mixes a kind (`World`, `Vendor`, `Quest`) with a boss or NPC name; row 1 ends `World · 1% drop`, row 2 `Interrogator Vishas`. Fine as built (the game's words), but state in the code comment which rows carry a kind so a later lane does not "fix" it.

### Checklist
- [x] Question answered above the fold: what you wear, what beats it, where from.
- [x] Header figures line `N of M best in slot · K enchants missing`, counts match the cells.
- [x] Gear dot on the tab; Guild dot absent where the guild has nothing urgent.
- [x] Level row: bands 20 to 60, chosen band teal, others blue.
- [x] Doll: two columns, icons on outer edges, off hand after main hand for a one-hander (warrior).
- [x] Cell states: `Best in slot` muted, gain green, `in your bags` client gold, `No enchant` warning, `Empty`.
- [x] Upgrades: top four, total in the note, `In your bags` client gold.
- [x] Weights: top stat 1.00, note `shipped, per point`.
- [x] Sources: eight kinds, two columns, counts, state-word note, `Off: ...` line, `Default` only while a choice is saved (warrior shows none, correct).
- [x] Footer right `Shift-click an item to link it`.
- [x] Alt-click on a marked slot and the character pane button (code present in `SlotMarks.lua`, `PanelTab.lua`; not visible in a render, verify in the in-game capture).
- [ ] Section headings per spec 3.5 (must 1).
- [ ] Checkbox as built is shown (must 2).
- [ ] Weights table reads as a ranking (should 1).
- [ ] Boss never cut in a source line (should 2).
- [ ] Empty slot silhouette (should 3).
- [ ] Hover, focus-visible, busy states of Level, Source toggle and doll cells: not in a render; attach captures (hover cell with the pick tooltip, Alt-click 2 s hover ground).

Re-review after must 1 and 2 land.

## Round 2: SHIP

Re-reviewed the re-made `gear-hunter-60.png` and `gear-warrior-34.png` against the board and Round 1.

### Round 1 items
- Must 1, headings: closed. Spec 3.5 amended, one rule for every tab, the render matches the amended spec and the Builds tab.
- Must 2, checkbox: closed on the render's evidence that it is `Widgets.toggle`'s own recipe (gold filled square with a dark tick, 1 px muted outline unchecked). It is heavier than the board's outline box, but it is the one Settings control and spec 3.5 allows no second recipe. Condition: the in-game hover and keyboard-focus captures from the Windows box are attached before the lane is called done.
- Should 1, weights: half closed. Intellect is gone (under 0.05 of the top). Hunter still lists `Attack Power` and `Ranged Attack Power` at an identical 0.33 / 0.14. Fold the pair into one row.
- Should 2, source line: closed. Warrior rows keep boss and chance (`Boulderfist Enforcer · 1% drop`, `Charlga Razorfla… ` no longer cut), the place gives way (`Scarlet M…`, `Razorfen K…`), the hover carries the whole line.
- Should 3, empty cell silhouette: not visible. The render still shows the black square; the lane says the code draws `Theme.SLOT_ICONS` and the mock stands in. Accepted on that word; the in-game capture of an empty Hands cell proves it.
- Should 4, name column: nearly closed. `Brawler's Leather Legguar…` still loses one letter in the right column.
- Should 5, vocabulary: closed by dropping kind words for vendor and mob names.

### New in this round, judged
- `Better than the list` replaces `Beats the pick` and counts in the header (hunter `10 of 16`: nine `Best in slot` cells plus Slashclaw Bracers). Clear, and the figure now equals what the doll shows.
- `No enchant: Chest, Feet` under the doll, warning colour. Matches the cells (hunter Dawn Armor, Bloodmail Boots; warrior Legs). Good: the player reads the list without scanning the cells.
- `More upgrades` (ranks 5 and up): accepted. Same row grammar at 18 px icons, clearly secondary. It makes the window taller for the warrior (about 40 px at 1x over the hunter); height still fits one value per open window.
- Sources note blank at `All sources` and the `Off:` line hidden at the leveling default: accepted as quieter, but the spec's copy still lists `All sources`; amend 4.3 so the next reader does not file it as a miss.

### Should, remaining (none blocks)
1. Fold identical `Attack Power` / `Ranged Attack Power` rows.
2. Give the right doll column the width its longest name needs.
3. When a slot is dropped from the upgrade source line (warrior `Bloodmail`-style rows 2 and 3 lose `Off Hand`, `Main Hand`), keep it in the hover; confirm it does.
4. Hunter haste line absent: confirm the shipped weights carry none.
5. Attach in-game captures: hover cell with the pick tooltip, Alt-click 2 s hover ground, empty-slot silhouette, toggle hover and focus.
