# Density review, round 10: sim, character, planner

1440 read for every surface (sim-saved judged from sim-result, same body). Handed-off and needs-data items excluded.

## Sim
- sim-landing-in: DENSE. Gear with stat lines, Talents, Your sims with deltas, Rotation; rail ends ~y1520 against left ~y1605, balanced. Optional: link each Your sims row to its saved result.
- sim-result: DENSE. Damage by ability with share bars, DPS spread vs BiS, timeline, gear, talents, rage, buffs, rotation with casts, This run. Optional: Copy CSV on Damage by ability.
- sim-saved: DENSE. Same body as sim-result.
- sim-gear (Top Gear): DENSE WITH ADDITIONS. Rail ends ~y1085 against left ~y1345 (~260 px empty under Find an item); finder shows 8 rows of 100 with "8 more". Show more finder rows. Candidates rows show only ilvl; add the equipped item and its ilvl beside Lionheart Helm and Arcanite Reaper.
- sim-talents (Talent compare): DENSE. Delta list fills the right of the guide build row; trees and deltas use the plate. Optional: delta column on saved builds when one exists.
- sim-drops (Droptimizer): DENSE WITH ADDITIONS. Counts and ilvl per crafted profession, one Raids date line. Dungeons, Reputation and PvP are still a bare "n sources" with no ilvl range or top item; Raids lists no sources. Show ilvl range or top item on each collapsed row.
- sim-weights: DENSE. Scale bars, scale, per point, error, rating conversions, link to level 60 raid set.
- sim-specs: DENSE WITH ADDITIONS. Rank, bar, DPS, class colours. Add sortable columns, a guide link per row, and link each Healers and tanks name to its guide.

## Character
- character-visitor-60: DENSE WITH ADDITIONS. Rail ends ~y1200 against left ~y1415 (~215 px empty under Activity). Standing absent on the visitor view; show it if the public payload carries it (lib/character/standing.ts). Best per encounter: only Lucifron has spark and delta; Magmadar, Gehennas, Garr, Baron Geddon lack both.
- character-owner-60-premium: DENSE WITH ADDITIONS. Columns level (~y1580/1577); Standing, Builds seen, Activity, Score, Upgrades, History, Coach present. Same per-boss spark/delta gap on four of five bosses.

## Planner
- planner: DENSE WITH ADDITIONS. Guide order uses ~420 of 1320 px, right two-thirds empty. Guide builds rows carry only spec, level and spread; add points off and the split delta as planner-you does, plus a Load action.
- planner-you: DENSE WITH ADDITIONS. Rich guide-build rows, Share, Import. Point order uses ~420 of 1320 px; guide build rows have no Load action; ~50 px gap under the rail card.
- planner-guide: DENSE WITH ADDITIONS. Loaded tag, delta rows, Share and Import in the rail. Point order empty right; per-cell tooltip (talent, rank, level) on the order icons is the cheap win.

## Verdicts
sim-landing-in DENSE | sim-result DENSE | sim-saved DENSE | sim-gear DWA | sim-talents DENSE | sim-drops DWA | sim-weights DENSE | sim-specs DWA | character-visitor-60 DWA | character-owner-60-premium DWA | planner DWA | planner-you DWA | planner-guide DWA
