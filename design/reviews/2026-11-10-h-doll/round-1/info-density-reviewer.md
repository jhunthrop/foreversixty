# info-density-reviewer, round 1 (captures at 2b92f51cd)

The three surfaces keep their content. Lost: the report's off-hand "Two-hand in main hand" line, since the list holds only filled cells.

VERDICT: DENSE WITH ADDITIONS

1. Report: "item level 183" in the head and "Item level 60.0" above the frame disagree. Show one.
2. Report: enchant ids are reduced to "Enchanted". Name them via lib/sim/enchants, as the sim does.
3. Report: gems and bonus ids are not shown (only if the build has them).
4. Report: list the off hand ("Two-handed") again. Flag missing enchants (needs the enchantable-slot data).
5. Sim Gear and strip: an enchant hides the item level. Use "Slot · item level · enchant".
6. Strip: pass enchants through the same gearSlots path.
7. Sim Gear: with a result but no stats, say so. Drop the empty set-lines span.
8. The commit text promises gains on the sim, but there is no gain path there.
9. Stat deltas against the last run (needs data).
