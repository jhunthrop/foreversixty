# Round 4 verdicts (captures at 2b3f13726, in captures/)

Round 4 changed the capture set only: it wields the fixture's two one-handers, so the fixture's measured white misses are the set's own. The product is round 3's.

## wow-player
The frame, the worn list and the figures read like the game; quality colours right; both weapons carry Crusader. Musts: the level 29 landing beside the level 60 result on the same set; the 13.9% white miss against a 9.5% hit (expected about 17.5% dual wielding); Gloves - Greater Strength at +10 against Bracer - Greater Strength at +7; the enchant line's "Strength +8" order against the tooltip's "+18 Strength"; "+28 Crit" and "+20 Hit" without their rating unit. Shoulds: the figures first at 390; the lone Haste cell; the unmet Requires Level in red.

VERDICT: FIX

## Response (round 5, 0959b3592)
- The level: the capture's export now spends a level 60 build (the BiS rail's Fury talents), so the landing, the browser run and the server run all read Level 60 over the same set.
- Outside this lane, routed to the coordinator with the h-doll lane's earlier notes:
  - The measured miss line prints the run's own counts (lib/sim/gear-card measuredMiss over the fixture's summary); the figure is the engine's for src/fixtures/sim/result-full.json. A white miss the engine should not produce is a sim-accuracy item, not a gear panel one.
  - The enchant values and names are the build's enchants.json rows as the client prints them (Enchant Gloves - Greater Strength +10 in this client's table; Lesser Arcanum of Voracity "Strength +8"); the sim and the report read the one table.
  - "+28 Crit" and "+20 Hit" without a rating unit: the shared ItemTooltip's stat lines, already routed by the h-doll review (2026-11-10-h-doll/final-verdicts.md).
  - The figures' order at 390 and the lone Haste cell: the stats block the sim results spec places with the gear; the unmet Requires Level in red: the shared tooltip.
