# Round 2 verdicts (captures at a6c1c819b, in captures/)

Round 2 changes: the frame centred when nothing sits beside it; from a 760 px panel the run's five figures, the measured miss line and the set lines in the column beside it; Upgrades on Best in slot the section's action in every state; an older run's note under the frame; Fury enchants on the capture set; sim-result-* is the parity fixture's server run with the stats sheet.

## ux-designer
Round 1 shoulds fixed: the frame centred, the action in the head in every state, a capture with the stats sheet, the phone list intact. Shoulds: the stats column ends high and leaves the panel beside the lower frame empty at 1440 and 2000; the result capture is a different character (no enchants, level 60) from the landing and browser-run captures; the "after a rerun" line still floats; say in the notes why the levels differ.

VERDICT: SHIP

## wow-player
The frame, title bar, crest, borders, ring and tooltip match the BiS and character dolls; the tooltip wording is the client's; the action sits top right in every state; the Fury enchants show. Must: the result capture wears a different set with no enchants and reads Level 60, while the landing and browser-run captures wear the capture set at level 29. Shoulds: the empty column under the figures; "Weapon skill 300 · 300" does not name the hands; "Haste from gear 0%" is a dead stat here; the "after a rerun" line; the unmet Requires Level in red (routed).

VERDICT: FIX

## ai-slop-reviewer
Reads as a player-made surface; the touched copy is clean. Non-blocking: the floating "after a rerun" line; the residual empty column beside the frame; cut `statsHead` to the boss level if it ever renders.

VERDICT: PASS

## info-density-reviewer
390 dense. 1440 and 2000: the column beside the frame uses about a quarter of its height; at 2000 the frame is no longer centred, so the panel is lopsided. Fill the column (the set lines; the enchant and item lines the phone list carries; totals) or move the figures under the frame and centre it.

VERDICT: THIN

## Response (round 3, 9054a48c4)
- The worn list stands beside the frame in every state (the report's Gear at the pull pattern, now one shared component, character/WornList.svelte, drawn by both): each piece over its slot and item level, the enchant on its green line, the covered off hand saying two-handed. After a run the figures, the measured miss line and the set lines sit above it. The column now runs the frame's height at 1440 and 2000. Under 423 px the list form carries the same lines and the worn list hides.
- The result capture answers the fixture's server run for the same capture set, so every state wears the one set with its enchants. The level on the title bar comes from the loaded character (the capture's export yields 29; the fixture's own run is 60).
- Unchanged, pre-existing copy outside this lane: "Your stats in the fight show after a rerun." (lib/sim/input-cards-copy statsNeedNewerRun, shown only for a result without a stats sheet; it sits under the pair), "Weapon skill 300 · 300" and "Haste from gear" (the stats block's labels).
