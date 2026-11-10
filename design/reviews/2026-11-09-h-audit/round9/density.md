# Density review, round 9: sim, character, planner

1440 read for every surface, sim-talents also at 390. Handed-off and needs-data items excluded.

## Sim
- sim-landing-in: DENSE WITH ADDITIONS. Gear (with stat lines), Talents, Your sims with deltas, Rotation. Left column ends ~y1605 against rail ~y1520. Optional: link each Your sims row to its saved result.
- sim-result: DENSE. Damage by ability with share bars, spread vs BiS, timeline, gear stat strip, talents, rage, buffs, rotation with casts, run facts, Copy JSON. Optional: Copy CSV on Damage by ability.
- sim-saved: DENSE. Same body, Saved date, Open in the sim and Copy link in the band.
- sim-gear: DENSE WITH ADDITIONS. "0 valid combinations" is gone. Rail ends ~y1085 against left ~y1345 (~260 px empty under Find an item); the finder shows 8 of 100 with "8 more", show more rows. Candidates rows show only ilvl; add the equipped item and its ilvl beside each.
- sim-talents: DENSE WITH ADDITIONS (was THIN). The delta list is now its own column, no collisions, at 1440 and at 390 (stacks under the tree). Remaining: trees use ~430 px of the 1250 plate, right ~400 px beside the current build is empty; put the delta column between the builds or fill that space.
- sim-drops: DENSE WITH ADDITIONS. Counts and ilvl per source, one Raids date line. Dungeons, Reputation and PvP are collapsed to a bare count; show the ilvl range or top item on the collapsed row.
- sim-weights: DENSE. Scale bars, scale, per point, error, rating conversions, link to level 60 raid set.
- sim-specs: DENSE WITH ADDITIONS. Rank, bar, DPS, class colours. Add sortable columns and a guide link per row; link each Healers and tanks name to its guide.

## Character
- character-visitor-60: DENSE WITH ADDITIONS. Rail ends ~y1200 against left ~y1415 (~215 px empty under Activity). Standing absent on the visitor view; show it if the public payload carries it (lib/character/standing.ts), else extend Activity. Best per encounter: only Lucifron has spark and delta; Magmadar, Gehennas, Garr, Baron Geddon lack both.
- character-owner-60-premium: DENSE WITH ADDITIONS. Columns level (~y1580/1577); Standing, Builds seen, Activity, Score, Upgrades, Coach reviews present. Same per-boss spark/delta gap on four of five bosses.

## Planner
- planner: DENSE WITH ADDITIONS. Rail and Talents level, guide order present. Guide order uses ~420 of 1320 px, right two-thirds empty. Guide builds rows carry only name, level, spec; add points off and split delta as planner-you does, plus a Load action.
- planner-you: DENSE WITH ADDITIONS. Rich guide-build rows, Share, Import. Point order uses ~420 of 1320 px; no Load action on guide builds; ~90 px gap under the rail card.
- planner-guide: DENSE WITH ADDITIONS. Level rail, Loaded tag, delta rows. Same Point order empty right; per-cell tooltip (talent, rank, level) is the cheap win.

## Verdicts
sim-landing-in DWA | sim-result DENSE | sim-saved DENSE | sim-gear DWA | sim-talents DWA | sim-drops DWA | sim-weights DENSE | sim-specs DWA | character-visitor-60 DWA | character-owner-60-premium DWA | planner DWA | planner-you DWA | planner-guide DWA
