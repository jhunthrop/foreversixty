# Round 1 response (lane h-doll)

## Fixed for round 2
- **One row format on all three surfaces:** "Slot · Item level N · Enchant", with the enchant in the client's green.
  - The report names enchants from the build's enchant table (lib/sim/enchant-names.ts) and falls back to "Enchanted" only when an id has no name.
  - The sim Gear rows keep the item level on enchanted pieces.
  - The saved sim's Character strip shows enchants through the same `withEnchants` (wow 1, 4; density 2, 5, 6; slop 6).
- **Tooltip:**
  - The enchant sits on the tooltip's own green line, in the client's position.
  - The duplicate item level footer is gone; a gain line shows only when the set has an upgrade (wow 3; slop 4).
- **Inner box:** the headless list form has no second box inside the section panel (slop 1, ux 5).
- **Phone tooltip:** the report's list form no longer stacks the tooltip under 17 rows. Each item opens the page's tooltip by the row tapped (`listCentre={false}`) (wow 7; slop 2 on phone).
- **Item level:** the orphan "Item level 60.0" line is gone; the section meta carries the log's item level (ux 3; wow 5; density 1).
- **Off hand:** the report list shows a two-hander's off hand again, as "Off hand / Two-hand in main hand" with the greyed icon (ux 6; density 4).
- **Sim footer link:** now "Upgrades on Best in slot ›", the b2 mock's words (wow 6; slop 4).
- **Sim Gear stats and sets:**
  - A run without the engine's stats sheet says "The stats sheet needs a newer run."
  - The empty set-lines span is gone (density 7).

## Kept, with the reason
- **The frame's centre tooltip on desktop (slop 2):** the directions spec Paperdoll section puts the selected slot's client tooltip in the centre, and BiS and the character page do the same. It opens on the first item so the centre is never empty (web layout §4.5).
- **The frame plus the worn list at 1440 (slop 3):** this is the b1-report-player mock's layout. The frame shows the set, and the list gives the facts a log reader scans.
- **Two even columns and the client's slot order (slop 7, wow 8):** this is the spec's list form and the client's pane order.
  - The report list runs head, neck, shoulder, back, chest, wrist, hands … trinkets, then weapons, which is the frame's order.
  - One column under 360 px.
- **Tooltip stat wording ("+28 Crit") and weapon damage lines (wow 2, 10):** these come from the shared ItemTooltip and items model used site-wide (BiS, character, planner), and Forever's client data. That is outside this lane; it is routed to the coordinator.
- **Source pill on the strip and "Open in the planner" on the report (slop 5, 8):** both predate this lane and were not changed by it. Routed.
- **Out of scope here:**
  - missing-enchant flags (needs per-slot enchantable data);
  - gems and bonus ids;
  - stat deltas against the last run.
