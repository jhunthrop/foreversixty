# AI-slop review, sim pages, round 2 (2026-11-09)

Method: all nine pages opened at 1440 and 390. Result/Saved 390 read at the header, ability table and Rage panel (12k px pages), rest skimmed against the 1440 capture. Framed inset plates are sanctioned (design-system Plate frame). No glows, gradients, glass, em dashes or marketing verbs anywhere.

## First impression
Reads as made by a player. Round-1 FAILs are fixed: Top Gear fills with Try each and consumable icons, Talents shows real trees, Droptimizer carries item counts and ilvl, Specs has crests and DPS rank.

## Verdicts
- sim-landing-out: PASS WITH FIXES
- sim-landing-in: PASS WITH FIXES
- sim-result: PASS WITH FIXES
- sim-saved: PASS WITH FIXES (same rage unit and caption nits as result)
- sim-gear: PASS WITH FIXES (tell 13 cleared)
- sim-talents: PASS WITH FIXES (tell 13 cleared)
- sim-drops: PASS WITH FIXES (no icons, flat lists)
- sim-weights: PASS WITH FIXES
- sim-specs: PASS WITH FIXES

## Findings (not in handoffs.md)
All tool pages, 1440+390: band sub-line "Damage specs only." still under every title (tell 7/21, web/src/lib/sim/page-copy.ts). Cut; Spec support says it.
sim-landing-out
- Same sub-line. Rest clean; Get set up hint kept.
sim-landing-in
- 22 top-right "No sim yet" at 1440 and under the name at 390. Cut; Run sim is the cue. page-copy.ts.
- 22 gear sub-lines mix "Item level 65", "Strength +8", "Crusader", "All Stats +4" (SimGearSection.svelte). Pick one; enchant/stat line only where it exists, ilvl otherwise is fine, but not both rules per row.
- 390: 15 gear rows plus both full talent trees push Run results/Rotation below 2000px; collapse the trees to the 17/32/2 split on phone.
sim-result / sim-saved
- 22 Rage panel "2,127 rage / 2,116 rage / 57 rage · 2.7%" and "Gained from · 2,127" (SimResourceSection.svelte). Title carries the unit: 2,127 / 2,116 / 57 · 2.7%. At 390 "2.7%" touches the plate edge.
- 22 "Recklessness counts once a pull here. In a raid it is back every 30 minutes." (SimFidelityCaption.svelte). At 390 its icon sits alone on a line above the text. Drop caption, or inline the icon.
- 8 Rotation card still lists greyed "Not cast in this run" steps 5 and 11 (SimRotationSection.svelte). Drop uncast steps.
sim-gear
- 21 "Nothing ticked yet." under the bar and "Saved builds did not load." (copy.ts:906) dev-speak. Show "No saved builds".
- 22 context strip "Orc Fury Warrior · Level 60 · Item level 62.0 · Entered by hand · 20 of 51 talent points". Drop "Entered by hand" (lib/sim/sources.ts:145) and "of 51"; "Level 60" repeats the title.
- Find an item list still clips mid-row ("Lightbreaker Cuirass") at both widths; "Showing 100 of 9026. Narrow the search." fine.
- Candidates plate is two rows with "Equipped" under the name; "Lock to equipped" checkbox twice. Ok, but 390 drops the ilvl under the name onto its own line.
sim-talents
- Same "Nothing ticked yet." and "Saved builds did not load." plus context strip as Top Gear. Current build row shows "0/20/0" on a Fury 20 tree: fine.
- Trees sit tiny (28px cells, mostly empty greyed) with a 400px dead label column at 1440; give each row the diff vs the other build.
sim-drops
- 19 no icons or crests on any dungeon, profession or faction row; 21 dungeons each carry a "Bosses n" disclosure line that doubles row height (Droptimizer, lib/sim/loot.ts). Put the boss count inline.
- "Nothing ticked yet." appears twice (bar and above Sources).
- PvP: 15 flat "Rank n (Horde)" rows, "Rank 15" and "Rank 5" read "0 items", three "PvP reward (Horde)" rows with no distinction, Rank 6 "ilvl 1". Hide empty rows, name the rewards, one row per rank with both factions.
- Raids block is now one line (good). Reputation grouped by faction (good).
sim-weights
- 7/21 grey line "Per point of stat, normalized to Hit = 1.00 ... Haste: 9.25 per 1%, not in the table because no item at this level range has it". Make Haste a row.
- 7/21 "Weights hold near the gear you have on. For a real choice, sim the items. Open Top Gear" explainer above Stats to weigh. Cut, keep the link in the band.
- 22 "Level 60 raid set" link beside Published weights; "3,120 iterations" caption under the bar (copy.ts:1085) move to the Run title.
- Stats have no icons; disabled ticked "Attack power" box.
sim-specs
- 7 "What it does" disclosure repeated on 18 rows; at 390 names wrap ("Beast Mastery Hunter", 65px rows). Use one link per class or an icon-only info toggle.
- 22 "Raid BiS DPS at 60" next to the Damage specs title and "DPS" after every number: keep the unit in the column, drop one.
- Healers and tanks: 8 rows of "Guide ›" with no detail; thin. Fold into one line of links.

## Authentic, keep
904.4 DPS as the hierarchy top; ability table with miss/dodge/glance; timeline with Execute band; rage ledger with lost-at-cap; buff uptimes with counts; Your sims rows; SimC-convention weights table with bars; class-coloured crests and names on Spec support; Try each with real consumables; talent trees with point splits; Droptimizer item counts and ilvl per dungeon; rank + DPS on specs.

## Handed off (noted, not failed)
Timeline Off/On toggle unlabeled (f-content). Sim gear stat strip and BiS DPS target (h-doll).
