# Round 10 UX verdicts: sim pages (1440 + 390), regression check vs round 9

| Page | Verdict |
|---|---|
| sim-landing-out | SHIP |
| sim-gear | SHIP |
| sim-talents | SHIP |
| sim-drops | SHIP |
| sim-weights | SHIP |
| sim-specs | SHIP |

No regressions against round 9. Handed-off items and fixture artefacts (dim healer/tank icons, blank icons) do not fail a page.

## Held from round 9
- sim-talents 1440: trees left, +/- list in its own column; 390 list under trees. Still clean.
- sim-gear 390: Candidates rows keep "Lionheart Helm" on one line beside ilvl and Modify. Round 9 polish closed.
- sim-weights: Run enabled; stat/scale/per point/error aligned at 1440; at 390 each row stacks name, number, bar, "0.738 DPS · ±0.12 error".

## Polish, not failures
- sim-weights 1440: ~450 px gap between stat name and bar (open from round 9).
- sim-talents: flat diff list, no Gains / Drops grouping; ~450 px empty right of each build row at 1440.
- sim-gear 1440: "Minimum item level" and "Slot" labels still on different baselines.
- sim-landing-out 390: best-in-slot class rows leave an empty right half on 1-spec classes (Paladin, Priest); page ~2,140 px.
