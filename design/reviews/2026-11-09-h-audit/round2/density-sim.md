# Density review: simulator (round 2)

Verdict: DENSE WITH ADDITIONS. No MUST remains in this lane. Read at 1440 (all eight pages) and sim-landing-in at 390; sim-result and sim-saved 390 not read in detail.
Page job: the sim owns numbers. Items in handoffs.md (SimGearSection stat strip, BiS DPS target on landing, SimTimeline toggle) are excluded.

## Round 1 MUSTs, status
1. landing-in dead space: FIXED. Left (Gear, Talents) ends ~y1330, rail (Rotation, Your sims) ends ~1375. Rotation rail is icon + name + rank rows.
2. Your sims delta: FIXED (+497.3, -478.9 under DPS). No sparkline, optional.
3. result/saved header gap to BiS: FIXED (+/- 1.0 DPS, ~99.6 vs BiS under 904.4 DPS).
4. Top Gear dead left column: FIXED. Try each, Whole sets, Talent builds sit under Candidates, rail holds Find an item.
5. Droptimizer: FIXED. Every source row has items and ilvl, dungeons/crafted/rep/PvP in 3 columns, boss counts, rep standings with counts.
6. Talents compare: FIXED. Current and guide builds as side-by-side tree strips with point totals; Open in planner link.

## Additions worth making (ranked, optional)
1. sim-gear Candidates plate lists only Head and Main hand. Show every equipped slot as a row (icon, ilvl, Lock to equipped) so it matches the paperdoll. Data: character gear. Files: web/src/components/sim/tools/CandidateRows.svelte, TopGear.svelte.
2. sim-result Damage by ability: Copy CSV and sticky header at 390.
3. sim-landing-in Your sims: sparkline across saved runs of the same preset. File: web/src/components/sim/SimYourSims.svelte.
4. sim-result header: delta to the last run of the same preset (needs prior-run link, see needs data).
5. sim-weights: remove the grey lines "Per point of stat, normalized to Hit = 1.00...", "Weights hold near the gear you have on..." and "3,120 iterations" (don't-explain rule); the "10 Hit rating = 1%" sub-line is enough.
6. sim-gear/talents/drops: "Nothing ticked yet." and "No source ticked yet." are grey lines for a state the disabled Run button already shows.
7. sim-specs: healer/tank rows are name + Guide only (no sim exists, honest).

## Strong
Result page (damage table, spread, timeline, rage, buffs, this run, rotation with cast counts), specs table with rank and Raid BiS DPS, droptimizer counts, weights bars with error.

## Needs data
- Per-ability DPS of the BiS benchmark.
- Stored link to the prior run of the same preset (result delta, sparkline).
- Healer and tank sim output.
