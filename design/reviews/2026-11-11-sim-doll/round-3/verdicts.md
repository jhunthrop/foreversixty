# Round 3 verdicts (captures at 9054a48c4, in captures/)

Round 3 changes: the worn list beside the frame in every state (character/WornList, shared with the report's Gear at the pull), the figures and the measured miss line above it after a run; the result capture answering the fixture's server run for the one capture set.

## ux-designer
The column beside the frame is full; the result wears the same set as the other states; the action top right; the covered off hand reads "Two-handed"; one row pattern; the 390 list form carries the same lines. Shoulds: at 1440 the list outruns the frame by about 300 px; one figures grid at every width; highlight the hovered slot's row; the figures above the list at 390.

VERDICT: SHIP

## wow-player
The frame and the worn list match the report's; the tooltip wording is the client's; the round 2 blocker is fixed. Must: "Miss 0% specials · 13.9% white" beside a two-hander reads as a dual-wield miss table. Shoulds: the lone Haste cell; the level 60 result beside the level 29 captures; the floating "after a rerun" line; the enchant values to check against the client; the report and the sim reading the same enchant names.

VERDICT: FIX

## ai-slop-reviewer
Reads as made by someone who plays; the touched comments carry no tells. Minor: the lone Haste cell; the floating note; the figures after the list at 390.

VERDICT: PASS

## info-density-reviewer
The worn list runs the panel's height at 1440 and 2000 in every state; 390 dense. Suggestions: the figures under the frame at 1440, ahead of the list at 390; two-line phone rows.

VERDICT: DENSE

## Response (round 4, 2b3f13726)
- The miss line was the capture's: the fixture's result (a dual-wield run, src/fixtures/sim/result-full.json) answered for a two-hander set. The capture set now wields the fixture's two one-handers (Ironwood Blade, Persuader, both Crusader), so the measured white misses are the set's own. The product line is unchanged.
- The enchants are the build's own enchants.json rows (Lesser Arcanum of Voracity, Enchant Gloves - Greater Strength, Enchant Bracer - Greater Strength, Enchant Cloak - Lesser Agility, Enchant Chest - Greater Stats, Enchant Boots - Minor Speed, Enchant Weapon - Crusader), named as the client prints them; the report and the sim read the one table (lib/sim/gear-card enchantNames, character/WornList).
- Carried, outside this lane: the figures' order on a phone and under the frame at 1440 (the sim results spec places the stats block with the gear; a move is a spec change), the lone Haste cell, the "after a rerun" line, the hovered slot's row highlight, the unmet Requires Level in red.
