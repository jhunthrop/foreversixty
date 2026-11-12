# Home upgrades unit, persona verdicts (2026-11-11)

Capture: `home-41-upgrades-1440.png`, the signed-in home's Your upgrades panel at 1440 on the fixture build
(level 41 Protection warrior, `web/tests/e2e/fixtures/home/me-41.json`). Scope: the unit wording on the rows and
the foot after the fix (the model's `scoreUnit` through `scoreUnitWord`, the paperdoll's own word).

## ux-designer: SHIP

- Rows and foot use the same word; no DPS/score mix left in the panel.
- The wording matches the BiS page (`scoreUnitWord`, the word `doll-cells.ts` and `gearCopy.gainSummary` use).
- "score" reads clearly as a muted suffix after the bold figure; no hint needed, a Forever Score hover is optional polish.
- Make sure the foot is not mistaken for a five-row sum (it sits beside See all 17).

## wow-player: SHIP (after one fix)

First pass: "score" beside an item is honest now, but the foot "+1,058.4 score" under five rows read as their sum.
Fixed: with more upgrades than rows the foot says "score across all 17" (`upgradesFacts`).

Out of this lane (follow-ups): a one-line key or hover saying what a tank's score point is, carried by the
paperdoll too; Sword of Serenity doubted as a Protection main hand at 41 (a ranker question, not the unit);
"Aegis of the Scarlet Comma..." cut at this width.
