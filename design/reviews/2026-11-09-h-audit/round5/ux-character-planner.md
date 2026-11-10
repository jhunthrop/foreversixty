# UX review: character and planner, round 5 (round5-captures/, 1440 and 390)

Handoffs not failed: dashed Point order border, PointOrder dash marker, "Client build" footer, Paperdoll Upgrades card, item level/enchant, tier-gate gutter, spec median.

## Verdicts
- character-visitor-60: SHIP
- character-owner-60-premium: SHIP
- planner: SHIP
- planner-you: SHIP
- planner-guide: SHIP

## Regression check (round-4 layout changes): no blockers, no must-fix
- 390: no horizontal scroll on any of the five; tree fits and has a tab per tree; Score and Standing sit above Best per encounter on owner; no empty icon squares; no em dashes.
- 1440: planner-guide columns balanced (right ~1085 vs left ~1275); planner-you right column now shorter than left (was longer in r4), the r4 void is gone.

## Polish
- planner 1440: right column ends ~y800 against left ~1590, so ~800 px of empty tan beside Point order and Guide builds. Signed-out only; tan "Warrior talents" title on the blue band still reads low-contrast.
- character-visitor 1440: right column ends ~y1180 against left ~1415 (~235 px void).
- owner 1440: the upgrade nudge plate is 720 px of 1248, left-aligned, with an empty right half. Span full width.
- owner/visitor 390: the Lucifron row wraps "+82 vs 27 Oct" onto its own line.
- owner: History is still two closed rows, nothing previewed.
- planner-you/guide 390: the "Show DPS" toggle runs ~8 px past the plate content padding; the level build card with Load sits below Point order (~y1470), so the primary action is two screens down. Consider moving it above Point order.

## Checklist
- [x] Compact title band, crest
- [x] Paperdoll ornate, quality rings
- [x] Unit follows number
- [x] No em dashes
- [x] 390 no horizontal scroll, tree fits, tab per tree
- [x] Lists sit on plates
- [x] No empty icon squares
- [x] One job per page
- [x] Columns balance at 1440 (planner-guide, planner-you)
- [ ] One plate style (dashed Point order, handed off)
