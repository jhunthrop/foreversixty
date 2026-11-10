# Batch H integration review (2026-11-10)

The merged build of h-report-tail, h-api-html, h-doll, h-guide-spec (with h-audit) and h-guild-tabs
(with h-edges), captured by web/tests/captures/batch-h at 1440 and 390 on the fixture build, every
icon drawn from the real build data: sim result, a report's Player view, the Frost Mage guide,
guild Raids and Applications, /ratings and the shared build /b/{id} (the API's fixture HTML served
as is, its header loading from /_chrome/).

## ai-slop-reviewer

- Round 1: VERDICT: PASS WITH FIXES (ratings qualifiers, the /b explainer, the guide sources
  explainer, empty icon sockets).
- Round 2: VERDICT: PASS WITH FIXES (lazy icons not yet loaded in the shot, the sources label,
  "The sample" said three times on /ratings).
- Round 3: **VERDICT: PASS**

## ux-designer

- Round 1: VERDICT: NOT YET (missing icons, the guide's At a glance plate down to the Build
  column with 800 px empty, the report Player gear on fixture ids).
- Round 2: VERDICT: NOT YET (the At a glance plate: back to the b2-guide-spec three columns).
- Round 3: **VERDICT: SHIP**

## Fixed in integration

- /b/{id}: "Read only." beside Edit a copy; /ratings "0 to 2,100"; the tier rail shows only what is
  left to climb; guide sources read "Talents Forever".
- The guide's At a glance plate is Build, Rotation and Stat priority again (round-5 mock).

## Follow-ups (not blocking)

- Guide: "Full priority ↓" sits at the column foot with a gap above it; the Build trees are 28 px
  cells in a 400 px column; capture 2000.
- Guild Applications: header meta lacks the raider count the other tabs show; 390 drops the
  Waiting column and wraps Decline under Accept; the guild tab strip clips at 390 with no fade.
- Guild Raids 390: the meta line wraps with a dangling separator; Absent is dropped on collapsed
  nights.
- Shared build: at 390 "In the game" sits under Point order; Point order fills 670 of 1250 px at
  1440; Edit a copy and Copy for the game are both primary.
- Sim result 390: Damage by ability loses its share bars; timeline lane labels sit above their row.
- Report Player 390: empty slot rows; the sub-tab strip has no fade; "What these mean" floats
  outside its card.
