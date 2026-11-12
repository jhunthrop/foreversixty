# ux-designer, round 1 (boards in this folder)

**Verdict: SHIP WITH FIXES.** The addon surfaces are on spec: the leads, the question, the tiles, the dialog, the Following gold and the heights all hold. The must-fix items are board and render defects, plus one footer alignment.

## Must fix

1. `builds-dialog-druid`, the site row "Bear for Deadmines": it carries the paladin crest on a druid's dialog (the renderer's `site_row` hard-codes the paladin crest). Re-render with the class crest; if the addon did the same it would be a blocker there too.
2. `panel-ask`, the "Rotation at 13" block: section 6.2 says the rotation row is absent while the spec is undecided, and `Panel.rotation` returns nil with no rows. The mock typed the block by hand. Remove it.
3. `guild-tab-undecided`, Mallory's Build cell runs out of the frame; the Level and Spec headers sit tight. Widen the frame or the column.
4. The last page's footer line "/fs help opens this again." starts at an offset that aligns with neither the buttons nor the body column (`again:SetPoint("LEFT", footer, "LEFT", 4, 0)`). Set it to 0 and use the same on every last-page board.

## Should fix

- `chip-path-card-ask`: three solid gold buttons under a gold heading compete as the one emphasis. The chip pill's icon on the board is a spell icon, not the crest; if the shipped chip draws the crest, re-render it.
- Dialog own rows: the gap between rows equals the gap inside a row, so rows read as a flat list. Consider 6 to 8 px between rows.
- Dialog footer hints sit about 9 px above the bottom border on most boards; give them the same inset.
- No board shows a tour tile in hover, focus or active state, and none shows the tour stepping aside (the dialog opaque over the Set it up page).
- The title "Forever Sixty" and the active tab are gold, so "gold means only Following" holds for the rows, not the frame; the acceptance line should say "on the rows".

## What reads right

The leads at the content's top at 588 wide; the heights (346, 346, 434, 348, 282, 435, 435, 435, 435, 409, 333; dialogs 338, 359, 433, 338, 418, 378); the question on `-ask`, `-zero` and `-druid`; the tiles at 190 and 141; the followed path row with its count, next point and Change; the dialog's crest heading, own rows, Following, the body-coloured cursor, the companion line only without site rows, "Enter to preview" and "Enter to follow"; the guild tab's empty cells; the panel's question block. The Escape and combat paths for the step aside and the pooled rows are the pure tests' to prove.
