# ai-slop review, sim pages, round 10 (1440 and 390)

Excluded: handed-off items (per-row "Item level N", timeline toggle label, gear stat strip), fixture artefacts, haste-per-1% caption.
Round 9 fixes confirmed: Change source is a text link at 390; Whole sets stacks at 390; Candidates inner frame gone; Healers and tanks is a one-per-row list at 390.

## sim-landing-out: PASS WITH FIXES
- 390: spec links still wrap one orphan per row (Hunter "Survival >", Mage "Frost >", Warlock "Destruction >"); the class list is about 1,100 px. Two columns of class rows, or spec links stacked in a fixed column.
- 1440 clean: three plates with real labels, crests, "/fs copy".

## sim-landing-in: PASS
- Real gear, talents, sims, rotation with per-spell notes. Nit: 390 dead band between the character line and the tabs.

## sim-result: PASS
- Dense, data-led; "Level 60 raid BiS 1,004.0 . 9.9% under" is the point. 390 folds Rage and Buffs sensibly.

## sim-saved: PASS
- Same body, "Open in the sim / Copy link", "Saved 4 Nov".

## sim-gear (Top Gear): PASS
- Candidates, Try each, Whole sets, Talent builds read as one job each; Find an item list uses quality colours and real ilvls. Nit: "No saved builds" over an indented "Add a build".

## sim-talents: PASS WITH FIXES
- 390: both trees are ~100 px wide with ~14 px icons (illegible), and the +/- diff is still one 19-row column under them. Scale the trees to the plate width and put the diff in two columns.
- 1440: trees still occupy ~600 of 1,240 px; the current-build row is empty to the right of the tree.

## sim-drops: PASS WITH FIXES
- Raids is the one open group with no count, only a grey "First raids, 9 December 2026: Onyxia's Lair . No date yet: ..." line; reads as an empty expander (tell 13, mild). List Onyxia's Lair with its item count.
- Crafted: real professions, item counts, ilvl, authentic.

## sim-weights: PASS
- Scale, per-point DPS and error, bars, "10 Hit rating = 1% Hit". No tells.

## sim-specs: PASS WITH FIXES
- Healers and tanks has no numbers: names plus grey "Healer"/"Tank" chips; wraps ragged at 1440 and is a thin plate under a data-heavy one. Give it a DPS-equivalent column or group by role.
- "#n" repeats the sort order beside the bar; nit.

## Summary
PASS: landing-in, result, saved, gear, weights. PASS WITH FIXES: landing-out, talents, drops, specs. FAIL: none. No visual tell 1, 2, 5, 8, 12 or 15 on any page.
