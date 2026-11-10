# Density review, round 6 (final): sim, character, planner

Read at 1440 (390 not re-read in detail). Handed-off and needs-data items excluded.

## Sim: DENSE WITH ADDITIONS
Job: the sim owns numbers. Result and saved (damage table with share bars, spread, timeline, rage, buffs, rotation, run facts), landing (gear, talents, Your sims with deltas, rotation), specs (rank and bars), weights (bars, error) are all dense. All additions are round-4 carry-overs, optional:
1. sim-gear: Candidates plate still holds Head and Main hand only beside a 9026-item finder; list every equipped slot (CandidateRows.svelte, TopGear.svelte).
2. "0 valid combinations" still printed on gear, drops and talents with Run already disabled; cut.
3. Weights: "Haste: 9.25 per 1%, not in the table because no item at this level range has it" keeps its clause; trim to the number. Add Copy weights.
4. Result: Copy CSV on Damage by ability; sim-specs: own spec row (Fury Warrior) unmarked.
5. Try each and Talent builds checkboxes show no last-run DPS delta; sim-talents Guide build has a net list but no DPS.

## Character: DENSE WITH ADDITIONS
Job: one character's standing and gear. History folded into Best per encounter works: owner left and right columns end level (~y1580/1557), Standing, Builds seen, Activity and Forever Score fill the rail.
1. Visitor 1440: the rail ends ~y1180, the left column ~y1415, so ~235 px of empty ground right of Best per encounter, and Standing is absent while the owner has it. Show Standing for visitors if the public rating carries it (lib/character/standing.ts), else let Activity run longer.
2. Best per encounter: only Lucifron has a spark and "+82 vs 27 Oct"; Magmadar, Gehennas, Garr and Baron Geddon have none though two kills exist on 3 Nov and 27 Oct (history encounters[].points; CharacterFightTable.svelte). No boss row shows a report link.
3. Forever Score parts carry no delta per part (HistoryScorePoint first vs last); Builds seen has no fights per build.

## Planner: THIN
Job: pick and order talent points. Talents-beside-rail with Point order and Build vs guide beneath fixed the guide-mode dead ground (planner-guide gap is ~40 px; rows carry points off, split delta, first missing talent, Loaded tag). But:
MUST 1. planner-you 1440: the rail (band card, In the game, Share, Import) runs to y~1190 while Talents ends y~845, leaving a ~340 px empty block under the trees, then Point order starts full width. Move Point order into the left column under Talents (10 cells fit in 880 px) so it fills that block and Build vs guide stays full width; or fold Share and Import to one row. Planner.svelte layout.
Should: Guide builds rows still have no Load or hover (guideBuilds in planner.astro, lib/planner/guide-builds.ts); band card lists "Past it" but not "you n, guide m"; Point order cells lack a tooltip with talent name, rank and level.
