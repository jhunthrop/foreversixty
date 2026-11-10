# Density review: simulator (round 4)

Verdict: DENSE WITH ADDITIONS. No MUST. Read at 1440: result, saved, gear, talents, drops, weights, specs (landing-in and 390 not re-read in detail). Handed-off and needs-data items excluded. Job: the sim owns numbers.

## Round 3 status
- Weights grey lines ("normalized to Hit = 1.00", "3,120 iterations"): fixed. Only "Haste: 9.25 per 1%, not in the table because..." remains; drop the clause after the number.
- "0 valid combinations" line on gear, drops, talents: still shown (Run already disabled). Cut.
- Candidates only Head and Main hand (sim-gear): open. Show every equipped slot (icon, ilvl, Lock). CandidateRows.svelte, TopGear.svelte.
- Result damage table Copy CSV and sticky header: open. Your sims sparkline: open (deltas fine).
- Drops: Dungeons, Reputation, PvP now show "N sources" but not item counts like Crafted ("349 items ilvl 70"). Improved, optional.
- Specs healer/tank rows icon + name only: honest, no sim. Own-spec row still unmarked.

## Additions, ranked (all optional)
1. sim-gear: Try each and Talent builds are bare checkbox lists; show last-run DPS delta per row once a run exists. Same on talents for the Guide build (net vs current).
2. sim-gear: Candidates plate holds two rows beside a 9026-item finder; fill with the full slot list.
3. sim-specs: mark the player's own spec row.
4. sim-weights: add Copy weights.
5. Result: Copy CSV on Damage by ability.

## Strong
Result/saved (damage table with share bars, spread, timeline, rage, buffs, rotation with casts, run facts), specs table with rank and bars, weights bars with error, talents compare with net-point diff list.
