# UX review: planner talent tooltip clip fix (2026-10-31)

Verdict: SHIP

Reference: the site's item tooltip (same placement math, same panel, same type).

## Region by region (after vs before)
- Tooltip whole: before, 1440 top cut at the tree panel edge ("Improved T..."). After, at 1024, 1440 and 2000 the title, rank line, NEXT RANK line and gold "Level 60 build" line are all visible, with the border and rounded corners intact. Pass.
- On top of the neighbour column: after, the tooltip draws over the Level 60 build rail and the Point order rail (1440, 2000) and over the lower cells of its own tree (1024). Nothing paints over it. Pass.
- Legible: title in the heading face, rank in muted grey, "NEXT RANK" small caps, level-60 line in gold. Same hierarchy as the item tooltip. Body text stays readable over the rail text behind it because the panel is opaque. Pass.
- Phone, 360, keyboard focus: the tooltip is 100% inside the viewport (x 185 to 703 of 720 px), the focus ring on the cell is visible, and the tooltip clears the cell it describes. Pass.
- Regressions: panel corners, panel borders, header rule under ARMS / FURY / PROTECTION and the background art are all unchanged and none is clipped or bleeding. Removing overflow-hidden has no visible side effect at any width. Pass.

## Findings
- Polish: the "bottom" captures hover the fourth row (Defiance), not the lowest right-column cell, so the flip-above branch is not shown. Attach one capture of the bottom cell (the 0/1 shield at the foot of Protection) at 1440 to prove it.
- Polish: at 2000 the tooltip covers the top of the "See what differs" button while hovering. Transient, no action.

## Checklist
- [x] Whole at 1024, 1440, 2000
- [x] Above the neighbouring column
- [x] Matches the site tooltip look
- [x] 360 keyboard focus inside the viewport
- [x] No regression in corners, header rule, art
- [ ] Flip-above capture (polish, not blocking)
