# Density review, round 8 (final): sim, character, planner

1440 read for every surface, sim-talents also at 390. Handed-off and needs-data items excluded.

## Sim
- sim-landing-in: DENSE WITH ADDITIONS. Gear, Talents, Your sims with deltas, Rotation all present; columns close (~y1605 vs ~1520).
- sim-result: DENSE. Damage share bars, spread vs BiS, timeline, gear, talents, rage, buffs, rotation with casts, run facts, Copy JSON. Optional: Copy CSV on Damage by ability.
- sim-saved: DENSE. Same body as result, saved date and Open in the sim / Copy link in the band.
- sim-gear (Top Gear): DENSE WITH ADDITIONS. "0 valid combinations" still shows above the tool band (cut while Run is disabled). Rail ends ~y1100 against left ~y1365: ~260 px empty under Find an item; let the finder show more rows. Candidates rows carry only ilvl; add the equipped item and its ilvl for comparison.
- sim-talents: THIN (MUST). 1) Guide build row: the per-talent delta list ("+5 Unbridled Wrath", "-3 Lingering Rage") is drawn on top of the Arms/Fury/Protection trees at 1440 and 390; labels collide with icons and tree headings and are unreadable. Lay the delta list as its own column beside the trees. 2) The three trees fill ~90 px of a 1250 px plate; the right ~1100 px is empty at 1440. Put current and guide builds side by side with the deltas between. 3) Drop "0 valid combinations".
- sim-drops: DENSE WITH ADDITIONS. Source groups with counts and ilvl. Cut "0 valid combinations". The Raids and Crafted date lines still restate and run long (Crafted line repeats the six rows above it); keep the one date.
- sim-weights: DENSE. Scale bars, scale, per point, error, rating conversions, link to the level 60 raid set.
- sim-specs: DENSE WITH ADDITIONS. Rank, bar, DPS, class-coloured icons. Add sortable columns and a guide link per row. Healers and tanks is a bare list of names with no figure (needs-data if no HPS/TPS ladder exists).

## Character
- character-visitor-60: DENSE WITH ADDITIONS. Rail ends ~y1180 against left ~y1415, ~235 px empty under Activity; Standing absent (show it if the public payload carries it, lib/character/standing.ts, else extend Activity). Best per encounter: only Lucifron has spark and delta; Magmadar, Gehennas, Garr, Baron Geddon lack both.
- character-owner-60-premium: DENSE WITH ADDITIONS. Columns level (~y1580/1557); Standing, Builds seen, Activity, Score, Upgrades, Coach reviews. Same per-boss spark gap; no per-part Forever Score delta.

## Planner
- planner: DENSE WITH ADDITIONS. Rail and Talents level. Point order uses ~420 of 1320 px, right two-thirds empty. Guide builds rows carry only name, level, spec.
- planner-you: DENSE WITH ADDITIONS. ~140 px gap under Talents beside the tall rail (small). Point order same empty right. Guide builds carry points off, split delta, first missing talent; no Load action.
- planner-guide: DENSE WITH ADDITIONS. Level rail, Loaded tag. Same Point order empty right; per-cell tooltip (talent, rank, level) is the cheap win.

## Verdicts
sim-landing-in DWA | sim-result DENSE | sim-saved DENSE | sim-gear DWA | sim-talents THIN | sim-drops DWA | sim-weights DENSE | sim-specs DWA | character-visitor-60 DWA | character-owner-60-premium DWA | planner DWA | planner-you DWA | planner-guide DWA
