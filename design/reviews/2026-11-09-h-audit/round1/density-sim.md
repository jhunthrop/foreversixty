# Density review: simulator (round 1)

Verdict: THIN (landing-in, Top Gear, Droptimizer, Talents compare). sim-result and sim-saved are DENSE WITH ADDITIONS.
Page job: the sim owns numbers (DPS, breakdowns, weights, what-if comparisons). Captures read at 1440; 390 captures were not read in detail.

## MUST (page, element, source, file)
1. sim-landing-in, 1440: left column ends at y~1330 (Gear, Talents) while the Rotation rail runs to y~1890 plus Your sims to 2150. About 900px of empty ground under Talents. The rail's 12 rotation rows are full prose paragraphs (guide text, restates the guide page). Fix: rotation rail = icon + name + rank rows only (prose to the Fury guide link or a tooltip), and fill the left with the attack power / hit / crit / haste strip the result page already shows under Gear plus the Level 60 raid BiS DPS target. Data: guide rotation, bis/warrior-fury.json benchmark (lib/sim/bis-benchmark.ts), SimGearSection stat row. Files: web/src/components/sim/SimRotationSection.svelte, SimPreRun.svelte, SimGearSection.svelte.
2. sim-landing-in, 1440 and 390: Your sims rail rows give DPS only, no delta or trend. Add "+/- vs previous run of the same preset" and a small sparkline across the saved runs. Data: SimListRow DPS and dates already in the rows. File: web/src/components/sim/SimYourSims.svelte.
3. sim-result and sim-saved, header: 904.4 DPS carries only +/- 1.0. Add the gap to Level 60 raid BiS (1,004.0, already computed in the DPS spread card), and on sim-result the delta to your last run. Data: bis-benchmark match gap, saved rows. Files: web/src/components/sim/SimCharacterHead.svelte, SimAnswer.svelte, DpsDistribution.svelte.
4. sim-gear (Top Gear), 1440: Candidates plate holds 2 rows, then about 1000px of empty left column while the rail stacks Find an item, Try each, Whole sets, Talent builds down to y~2180. Move Try each, Whole sets and Talent builds into the left column under Candidates (two-up), keep Find an item in the rail; show every equipped slot as a plate row (icon, ilvl) so the plate is not empty. Data: character gear, items catalogue. Files: web/src/components/sim/tools/TopGear.svelte, ToolsView.svelte, CandidateRows.svelte.
5. sim-drops (Droptimizer), 1440: 7300px of unticked rows with no figures beyond "Every boss 320"; ~45 reputation and PvP rows in one column. Add per-source count of items this character can use and best ilvl, drop chance where loot.json has it, and set rep and PvP in 3 columns. Data: web/public/data/<build>/loot.json, items. Files: web/src/components/sim/tools/SourcePicker.svelte, Droptimizer.svelte.
6. sim-talents, 1440: one plate with two checkboxes and a failed-saved-builds line, 600px empty under it. Show the current build and guide build as compact tree strips (or differing-talent chips) with a link to the planner; do not restate the planner's job. Data: web/public/data/<build>/talents, guide build. Files: web/src/components/sim/tools/TalentCandidates.svelte, SimTalentsSection.svelte.

## Additions worth making (ranked)
- sim-result/saved Damage by ability: Copy CSV and a delta-vs-BiS column per ability (bis benchmark), sticky header on the 9-row table at 390.
- sim-weights: right of the table is empty under "Stats to weigh"; add per-stat weight vs the published spec value and ranking change. Remove the grey explainer line under Published weights ("Per point of stat, normalized to Hit = 1.00...") and the "3,120 engine iterations..." line (don't-explain rule); Hit is "10 rating = 1%" already in the row.
- sim-gear/talents/weights/drops all carry the same grey sentences ("Nothing to try yet...", "Three stages: ..." , "Simmed as a level 60 with these talents"): cut to figures.
- sim-specs: 20 rows of name + "What it does" + "Not yet" with no figures (fixture state); show best DPS or ladder rank per supported spec from the ladder data, and crest icons; Healers and tanks rows are bare names, add the spec crest.
- Tooltips: Rage and Buffs sections already have icons; add spell tooltips on hover for ability and buff names if missing.

## Strong already
sim-result: damage table with share bars, DPS spread, timeline, rage and buffs sections, this-run facts, BiS benchmark, and sim-saved mirrors it.

## Needs data
- Per-ability DPS of the BiS benchmark (only totals exist) for a per-row gap.
- Previous-run delta on a saved sim needs a stored link to the prior run of the same preset.

## Mobile
At 390 the result page is a single column; fold Rotation and This run behind a disclosure so Damage by ability and the timeline come first. Not verified against the 390 captures.
