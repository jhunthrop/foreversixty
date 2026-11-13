# ux-designer verdict, addon 2.0.43 bug fixes

SHIP

MUST: none.

SHOULD
- builds-follow-moves.png: the right panel header reads "Level 60" while the left reads "Level 60 · 51 of 51 points". Confirm the points text is a real state difference (reopen) and not dropped by the new hide-unused redraw. If dropped, restore it.
- tooltip-long-wrap.png: the source line breaks inside "Scarlet Commander Mograine" ("Scarlet / Commander Mograine"). Acceptable as the client wraps its own lines the same way, but a break at the " · " separators would read better if cheap.

Checked, fine
- builds-follow-moves.png: Following moves to the pressed row, gold bar and gold label, exactly one Following on the tab, the old row returns to the Follow button in the same position, no layout shift between panels.
- bags-baganator-arrows.png: arrow sits top right at 12 px, green on the dark icon corner, legible against the green quality borders, matches the paper doll glyph, drawn only on the two chip-listed upgrades.
- tooltip-long-wrap.png: the 2.0.43 tooltip keeps the belt tooltip width, nothing escapes the frame (2.0.42 spilled past it), the pick name and source show whole, colours and the +4.2 DPS stay on the first line group.
