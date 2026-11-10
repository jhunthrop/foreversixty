# Density review, round 7 (final): planner, character, sim

1440 read in full for every surface; planner-you 390 read. Handed-off and needs-data items excluded.

## Planner: DENSE WITH ADDITIONS (was THIN)
Job: pick and order talent points.
- planner-you 1440: the ~340 px hole is gone. Rail ends y~950, Talents y~845, a ~105 px step. Fixed.
- planner (base) 1440: rail y~800 vs Talents y~845, level.
- planner-guide 1440: rail y~860 vs Talents y~845, level; Guide builds rows carry points off, split delta, first missing talent, Loaded tag.
Additions (optional, ranked):
1. Point order plate: 10 cells fill ~420 of 1320 px, the right two-thirds is empty on all three captures. Add a per-row points total or a tooltip per cell with talent name, rank, level (planner/Planner.svelte).
2. Guide builds rows: Load action and hover; planner-you band card could list "you n, guide m".
3. planner-you 390 holds: Talents, Point order, band, Guide builds, Share, Import; no dead space.

## Planner-you: DENSE WITH ADDITIONS
## Planner-guide: DENSE WITH ADDITIONS

## Character visitor (character-visitor-60): DENSE WITH ADDITIONS (no regression)
Rail ends ~y1180, left column ~y1415: ~235 px empty under the rail, Standing absent (owner has it). Show Standing if the public rating carries it (lib/character/standing.ts), else give Activity more rows. Best per encounter: Magmadar, Gehennas, Garr, Baron Geddon still lack spark and delta.

## Character owner (character-owner-60-premium): DENSE WITH ADDITIONS (no regression)
Columns end level (~y1580/1557); Standing, Builds seen, Activity, Forever Score present. Same per-boss spark gap; no per-part score delta.

## Sim landing (sim-landing-in): DENSE WITH ADDITIONS (no regression)
Gear, Talents, Your sims with deltas, Rotation all present and columns level enough.

## Sim result (sim-result): DENSE WITH ADDITIONS (no regression)
Damage share bars, spread, timeline, gear, talents, rage, buffs, rotation, run facts. Optional: Copy CSV on Damage by ability.

## Sim drops (sim-drops): DENSE WITH ADDITIONS (no regression)
Source groups with counts and ilvl. Cut "0 valid combinations" while Run is disabled; the Raids and Crafted date lines restate and run long, trim to what is available.
