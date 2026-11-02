# Row bands: review verdicts (2026-11-01)

Spec: design/specs/2026-11-01-row-bands.md. Boards: `compare-<page>-<width>-<list>.png` (before on
the left), built by design/mocks/compose_row_bands.py from web/tests/e2e/row-band-captures.spec.ts.

## ux-designer: SHIP

No must-fix. Roster at 1440, 2000 and 390: the band is rounded, starts left of the crest, and the
right-hand text clears its edge. Readiness, sign-ups, raid report blocks, hover bands and the
addressed fight row all follow the spec. Full-bleed lists (rankings, tiers, BiS, account) rightly
untouched. Nothing closer than ~10 px to the screen edge at 390.

Shoulds:

1. Roster at 390: the row keeps ~44 px of empty space under its text (pre-existing: the
   Open in simulator / planner slot holds its height when the member has no addon export).
2. Character fight table at 390: the last row's rule meets the card's rounded bottom corners
   (pre-existing, not a band).
3. Attendance Present grid: the 4/4 sits ~8 px inside the band at its narrowest cell.
4. Readiness at 390: the lit band is 8 px wider than the rules (the reach rule, accepted).

## wow-player: SHIP

No must-fix. The own row is findable at a glance and reads as a deliberate row; the crest and the
right-hand text sit inside it; nothing clipped.

Shoulds: the planned-raid hover is quiet on the dark card at 390; the same roster phone dead space
as ux-designer's first should.
