# UX review: character and planner, round 3 (round3-captures/, 1440 and 390)

Handoffs not failed: Paperdoll Upgrades card, item level/enchant, PointOrder dash marker, "Client build" footer, tier-gate gutter, spec median.

## Verdicts
- character-visitor-60: SHIP
- character-owner-60-premium: SHIP
- planner: SHIP
- planner-you: SHIP WITH FIXES
- planner-guide: SHIP WITH FIXES

## Round 2 must-fix
- Owner upgrade nudge: FIXED. Plate with Devout Skirt icon in a quality ring, yellow text, one line, truncates ("Stratholme: B...") at 390; plate ~46 px tall. At 1440 the plate is 720 px wide against a 1248 px grid (POLISH: span full width).

## Remaining
- character-owner-60-premium, 390: Forever Score and Standing still sit after History, 700 px below Gear. POLISH (CharacterOwnerSections.svelte).
- History plate: still two closed rows with nothing previewed. POLISH (CharacterHistory.svelte).
- Standing line now fits one line; orphan fixed.
- planner-you, 1440 and 390: Point order plate still has a dashed border no other plate has. MUST FIX (use the solid plate).
- planner-you, 1440: right column ends y~1005 against a 1830 left column. POLISH, sticky top 16 px.
- planner-guide, 1440: right column ends y~990 against a 1795 left column, ~800 px void beside Point order and Share. MUST FIX as sticky top 16 px (unfixed since round 2).
- planner, 1440: tan "Warrior talents" on the blue band; signed-out default "Human . Level 9" with "0 of 51 points". POLISH.
- planner 390: tab per tree with counts, tree fits, plates stack, no horizontal scroll.

## Per page
- character-visitor-60: compact band, doll with quality rings, Upgrades card, Forever Score bars, Builds seen, Activity, Best per encounter. Right column ends 215 px above the left (POLISH). 390 slot list, no overflow. SHIP.
- character-owner-60-premium: all plates, nudge plated, columns balance at 1440. SHIP (polish above).
- planner: SHIP.
- planner-you: SHIP WITH FIXES (dashed plate).
- planner-guide: SHIP WITH FIXES (right-column void).

## Checklist
- [x] Compact title band, crest
- [x] Paperdoll ornate, quality rings
- [x] Unit follows number
- [x] No em dashes
- [x] 390 no horizontal scroll, tree fits
- [x] Lists sit on plates (nudge fixed)
- [x] No empty icon squares
- [x] One job per page
- [ ] One plate style (dashed Point order)
- [ ] Columns balance at 1440 (planner-guide, planner-you)
