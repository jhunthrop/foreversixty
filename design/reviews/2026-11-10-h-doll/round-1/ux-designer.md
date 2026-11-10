# ux-designer, round 1 (captures at 2b92f51cd)

No blockers. The doll swap holds up at 1440, 2000 and 390.

- **Report, desktop:** the frame sits left with the "Baelgrim · Warrior · 73.3 DPS" strip. The tooltip is centred in the frame and ends on "Item level 61 · Enchanted". Rows read "Head · Item level 61 · Enchanted". This matches b1-report-player-1440.
- **Report, 390:** the list form, with the tooltip under the list, as b1 places it.
- **Sim Gear:** the list form with no strip. The green enchant or the grey item level survives from before, and "Upgrades... Best in slot ›" is kept.
- **Saved sim Character strip:** the same list form. The slot, item level, "Two-handed" and "Open in planner" are all present.

Shoulds:
1. At 390, the report title row wraps.
2. At 390, the report strip drops the class.
3. The orphan "Item level 60.0" line sits beside "item level 183" (fixture).
4. At 390, the report list is about 1,600 px in a single column; b1 uses two columns.
5. At 390, the sim list's inner bordered box doubles the border.
6. On desktop, the report list drops the off hand when a two-hander is worn.

VERDICT: SHIP
