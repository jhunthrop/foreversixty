# Density review: simulator (round 3)

Verdict: DENSE WITH ADDITIONS. No MUST. Read at 1440: landing-in, result, gear, talents, drops, weights, specs; landing-in at 390. Handed-off and needs-data items excluded.
Job: the sim owns numbers.

## Round 2 additions, status
1. Candidates lists only Head and Main hand (sim-gear): still open. Show every equipped slot (icon, ilvl, Lock). Data: character gear. CandidateRows.svelte, TopGear.svelte.
2. Result Damage table Copy CSV and sticky header: not visible, open.
3. Your sims sparkline: open (SimYourSims.svelte). Deltas present.
4. Weights grey lines still there: "Per point of stat, normalized to Hit = 1.00..." and "3,120 iterations". Cut both.
5. "0 valid combinations" grey line on gear, drops, talents: open; Run is already disabled.
6. Specs healer/tank rows are icon + name only; honest, no sim exists.

## New additions (ranked, optional)
1. sim-drops: collapsed Dungeons, Reputation, PvP show only "0 of N ticked"; give boss and item counts like Crafted. Data: droptimizer source list.
2. sim-gear: Try each and Talent builds are bare checkbox lists; show the DPS delta from the last run once one exists.
3. sim-specs: mark the player's own spec row.

## Strong
Result page (damage table with share bars, spread, timeline, rage, buffs, this run, rotation with cast counts), specs table with rank and bars, weights bars with error, talents compare with diff list, landing-in has no dead space at 1440.
