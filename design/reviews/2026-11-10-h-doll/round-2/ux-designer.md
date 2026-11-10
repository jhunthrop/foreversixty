# ux-designer, round 2 (captures at 485622479)

Round 2 is a SHIP.

Round 1 shoulds:
- **3, the orphan item level line:** gone. The section meta carries the log's item level once.
- **5, the inner box:** gone at 390. The icon column now starts at 30 px instead of 56 px.
- **6, the off hand row:** back, as "Off hand · Two-hand in main hand" with the greyed icon.

Also checked:
- **Row format:** all three surfaces read "Slot · Item level N · Enchant", with the enchant in client green. The footer now reads "Upgrades on Best in slot ›".
- **Report tooltip:** the named enchant sits on its own green line, in the client's position. I accept this over spec §4.2's "Item level 61 · Enchanted"; the spec line should be updated to match.

Shoulds, none blocking:
1. Long enchants wrap in the report list at 1440, so row heights are uneven. Widen the list or clamp the enchant with a title.
2. "The stats sheet needs a newer run." shows on sim-gear 390. It is only the fixture's missing-stats state, and the empty state is correct.
3. Capture the phone's row-tap tooltip.
4. Bring "Warrior" back to the phone medallion in a later pass.

VERDICT: SHIP
