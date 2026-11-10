# ai-slop review, sim pages, round 8 (1440 and 390)

Handed-off items (h-doll gear strip, per-row "Item level N", timeline toggle label), fixture artefacts and the haste-per-1% caption are excluded.

## sim-landing-out: PASS WITH FIXES
- Only the "Sim your character" plate carries an inner second frame; "Other ways in" and "Sim a best-in-slot set" do not. Border by role or not at all (tell 8, mild).
- 390: spec links wrap one orphan per row ("Survival >", "Destruction >"); the class list runs about 900 px. Use two columns of class rows at 390.
- Real class crests, real spec names, "/fs copy" instruction: authentic.

## sim-landing-in: PASS
- Gear, Talents, Your sims, Rotation are real data with real names. Rotation notes ("Never at 60: Overpower needs Battle Stance") read as player prose.
- Nit: at 390 the character head leaves a dead band between the name and the tabs.

## sim-result: PASS
- Dense and data-led: Damage by ability, DPS spread, buffs, rage. The stacked plates are justified by distinct data. "Level 60 raid BiS 1,004.0 · 9.9% under" is the point of the page.
- Nit: "Crit on the boss 29.3%" and "Miss 0% specials · 13.9% white" are the only qualifier-heavy labels.

## sim-saved: PASS
- Same body as result, with "Open in the sim / Copy link" and "Saved 4 Nov". The header meta line is long but all concrete.

## sim-gear (Top Gear): PASS WITH FIXES
- "0 valid combinations" floats under the toolbar before anything is chosen: an empty counter posing as content (tell 13, mild). Hide it until a candidate exists.
- "20 talent points" beside a 17/32/2 build contradicts the other pages and is a qualifier nobody asked for (tell 22).
- 390: the character strip wraps to three rows and "Change source" is a full blue button floating alone.
- The Candidates plate for two items is thin next to the full-height Find an item plate.

## sim-talents: FAIL
- 1440 and 390: in the "Guide build" row the talent diff list ("+5 Unbridled Wrath ... -1 Piercing Howl") renders on top of the three tree grids. Text sits over icons and the Arms/Fury/Protection headings. A broken layout, unshippable.
- Trees are about 90 px wide, left-aligned in a 1,240 px plate, leaving a dead field at 1440. Put the diff list beside the trees and scale the trees up.
- "0 valid combinations" again (see sim-gear).

## sim-drops: PASS WITH FIXES
- The grey sentences "First raids, 9 December 2026: ... No date yet: ..." appear under Raids and again under Crafted. Under Crafted it lists professions with no date, which is meaningless (tells 7, 19). Keep one release line on Raids, drop it from Crafted.
- "0 valid combinations" again.
- Real professions, item counts and ilvl: authentic.

## sim-weights: PASS WITH FIXES
- 390: the error prints unlabeled after the DPS ("0.738 DPS · ±0.12", "0.669 DPS · < 0.01"). Label it or drop it on phones. The 1440 table labels it correctly.
- Normalized 1.00 top stat with bars and "10 Hit rating = 1% Hit": authentic. Haste caption stays by owner rule.

## sim-specs: PASS WITH FIXES
- The DPS column has no header and no unit ("725.5", "1,083.6"). A number takes its unit. Add a DPS column head.
- The page does not say what the ranks measure, and the rank (#19) sits beside a bar that does not match the list order (grouped by class). Sort by DPS or label the order.
- Healers and tanks is a bare list of names, six rows at 390, under a data-heavy plate. Fixture icons are blank here, so no finding on those, but the plate needs real data to stop reading as filler.

## Summary
PASS: landing-in, result, saved. PASS WITH FIXES: landing-out, gear, drops, weights, specs. FAIL: talents.
