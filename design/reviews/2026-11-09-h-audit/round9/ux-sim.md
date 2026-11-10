# Round 9 UX verdicts: sim pages (1440 + 390)

Handed-off items (h-doll ilvl rows, gear stat strip, SimTimeline toggle label) and fixture artefacts (blank icons, Orc/Troll mismatch) do not fail a page.

| Page | Verdict |
|---|---|
| sim-landing-out | SHIP |
| sim-landing-in | SHIP |
| sim-result | SHIP |
| sim-saved | SHIP |
| sim-gear | SHIP |
| sim-talents | SHIP |
| sim-drops | SHIP |
| sim-weights | SHIP |
| sim-specs | SHIP |

## Round 8 blocker closed
- sim-talents, Guide build row: trees and the diff list no longer overlap. 1440: trees left, +/- list in its own column at x 850. 390: list under the trees. Spec labels (Arms 17, Fury 32, Protection 2) and rank badges read cleanly.

## Polish, not failures
- sim-talents: the diff is one flat list; the Gains / Drops grouping from round 8 is not done (green plus and red minus carry it).
- sim-talents 1440: about 450 px empty at the right of each build row.
- sim-gear 390: Candidates rows wrap the item name ("Lionheart / Helm") against the ilvl and "Copy and modify". Give the name its own line.
- sim-gear 1440: Find an item "Minimum item level" and "Slot" labels sit on different baselines.
- sim-result 390: Timeline icon rows are about 8 px tall; readable only by zoom.
- sim-weights 1440: 450 px gap between stat name and bar.
