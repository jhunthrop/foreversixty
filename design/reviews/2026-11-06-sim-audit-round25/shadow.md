# Sim-accuracy round 25: shadow lane

Seed: Shadowform at 40 against Wand Specialization 2/2 at 10 and 11 in the Shadow priest leveling order, under the talent adoption rule. Build mode. Engine bf727699f (unchanged, no engine worktree), site wt/r25-shadow from int/r25 fd90dd921. Every run: `spec-breakdown -spec priest-shadow -preset bare -band B -level L -talents <order at L> -bis-dir $R/bis-frozen -iterations 10000`, seeds 1 and 50001, both factions (frozen band races: gnome / undead), through `$R/sim.sh`. Raw runs, order strings, driver and weighing script: `design/reviews/2026-11-06-sim-audit-round25/shadow/`.

## Priest: Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| SH1 | **Shadowform at 40 cannot be combined with Wand Specialization before 40 (candidate c).** | Client talents 70790 (`data/builds/1.60.1.70790/talents/priest.json`): Shadowform is tier 6 (0-based), so it needs 30 points in Shadow (`sim/leveling/talents.go:181` pointsPerTier 5, same gate as the planner, `web/src/lib/planner/rules.ts` validateOrder). A priest has level minus 9 points: 31 at 40. Shadowform at 40 therefore needs all 31 points in Shadow; each Wand Specialization point taken by 40 moves the form one level later (1 point: 41, candidate d; 2 points: 42, the current order). No single point sequence gets both. Only a talent reset does. | Proven | - | - |
| SH2 | **Candidate orders measured on the whole path 10 to 44.** (a) current: Wand Specialization 10-11, Shadowform 42. (b) Wand Specialization dropped from the ladder-first row: Shadowform 40, Wand Specialization 41-42. (d) one point of Wand Specialization at 10: Shadowform 41. (f) the best Shadowform-40 order I found: (b) with Blackout 2 ahead of Mind Flay (so Mind Flay lands at 28 with its rank 2) and Shadow Weaving 3 ahead of Vampiric Embrace (31-33). From 42 every order is the same string (table below shows 42-44 equal to the decimal). | Table below; weighing rules below | Aligned: current order holds | - | - |
| SH3 | **No candidate wins.** Over levels 10 to 41, (b), (d) and (f) all take longer than (a) at both seeds and both factions, under both weighing rules. Shadowform at 40-41 gains +58 to +62% for two levels, but the wand points are worth -6 to -13% at 10-19, -4 to -9% at 20-25 and -1.5 to -2.5% at 34-39 to the Shadowform-40 orders. The adoption rule needs a gain of at least 1% beyond error at both seeds; every candidate loses instead, so the guide order stays. | Weighing table below | Aligned | - | - |
| SH4 | **The guide now gives the reset at 40.** Forever's guides take the wand early and reset for Shadowform at 40 (Kargoz, Method), or skip the wand and take it back after the form (Kargoz's alternative). A reset at 40 gets the (a) path to 39 and the (b) path at 40-41, the best of both by the table. The leveling section adds one sentence: reset at the trainer at 40, all 31 points in Shadow with Shadowform, Wand Specialization back at 41 and 42. The ladder, the planner and the addon keep one sequence (a), which cannot express a reset. | shadow.md leveling section; sources below | Guide improvement (not a sim bug) | guides-handoff | S 9d02416ed |
| SH5 | **Bare band 20 runs a priest with no wand at 10 to 16.** The frozen band-20 wand (Cookie's Stirring Rod) requires 17 and `wearableGear` (`sim/cmd/spec-breakdown/main.go`, wearableGear) strips it with nothing in its place, so Wand Specialization reads 0.0% at 10-14 (`shadow/runs-bare-gear.tsv`). I measured 10-16 with the wands the guide names instead: `-gear ranged:11287` (Lesser Magic Wand, req 5) at 10-12, `-gear ranged:11288` (Greater Magic Wand, req 13) at 13-16. Same family as the platform lane's seed 1 (unarmed rogues, bowless hunters); casters lose their wand the same way. | runs-bare-gear.tsv rows 10-16; items.json required_level | CONFIRMED-BUG, routed to platform (not fixed here) | old (round 15 measured band 20 at level 20 only and missed it) | - |

### Weighing rule

Two rules, both stated before reading the result:
- **One weight per level:** the time at each level counts the same, so path time is the sum over levels 10-41 of 1/DPS (normalised to (a) per level).
- **Weighted by kills per level:** a level takes about (XP to next level) / (XP of a same-level kill) kills, each lasting mob HP / DPS; HP cancels between orders at the same level. Classic 1.12 XP-to-level table and the same-level kill XP 45 + 5 x level (Wowpedia, Experience point). This puts more weight on the 30s and on 40-41.

Path time against (a), levels 10-41 (positive = slower than the current order):

| Faction | Seed | rule | (b) SF 40 | (d) 1 wand, SF 41 | (f) SF 40 tuned |
|---|---|---|---|---|---|
| alliance | 1 | one weight per level | +5.58% | +2.49% | +4.42% |
| alliance | 1 | weighted by kills per level | +2.16% | +0.93% | +0.80% |
| alliance | 50001 | one weight per level | +5.60% | +2.47% | +4.42% |
| alliance | 50001 | weighted by kills per level | +2.17% | +0.90% | +0.79% |
| horde | 1 | one weight per level | +5.02% | +2.15% | +4.01% |
| horde | 1 | weighted by kills per level | +1.44% | +0.51% | +0.24% |
| horde | 50001 | one weight per level | +4.98% | +2.16% | +3.97% |
| horde | 50001 | weighted by kills per level | +1.41% | +0.51% | +0.22% |

(f), the strongest Shadowform-40 order, is still +0.2 to +0.8% slower by kills and about +4% by level. (a) holds.

### Every level measured (DPS; b/a etc. relative to the current order)

Levels 10-16 wear the guide's starter wands (SH5); 17-44 the frozen band gear stripped to the level. A = alliance (gnome), H = horde (undead). (f) equals (b) at 10-25 and from 34 on; it was run at 26-41. Errors are at most ±0.1 DPS on every run.

| Lvl | Band | Faction | Seed | (a) | (b) | (d) | (f) | b/a | d/a | f/a |
|---|---|---|---|---|---|---|---|---|---|---|
| 10 | 20 | A | 1 | 14.2 | 13.3 | 14.2 | 13.3 | -6.3% | +0.0% | -6.3% |
| 10 | 20 | A | 50001 | 14.2 | 13.3 | 14.2 | 13.3 | -6.3% | +0.0% | -6.3% |
| 10 | 20 | H | 1 | 14.0 | 13.1 | 14.0 | 13.1 | -6.4% | +0.0% | -6.4% |
| 10 | 20 | H | 50001 | 14.0 | 13.1 | 14.0 | 13.1 | -6.4% | +0.0% | -6.4% |
| 11 | 20 | A | 1 | 15.1 | 13.4 | 14.3 | 13.4 | -11.3% | -5.3% | -11.3% |
| 11 | 20 | A | 50001 | 15.1 | 13.4 | 14.3 | 13.4 | -11.3% | -5.3% | -11.3% |
| 11 | 20 | H | 1 | 15.0 | 13.2 | 14.1 | 13.2 | -12.0% | -6.0% | -12.0% |
| 11 | 20 | H | 50001 | 15.0 | 13.2 | 14.1 | 13.2 | -12.0% | -6.0% | -12.0% |
| 12 | 20 | A | 1 | 15.2 | 13.5 | 14.4 | 13.5 | -11.2% | -5.3% | -11.2% |
| 12 | 20 | A | 50001 | 15.2 | 13.5 | 14.4 | 13.5 | -11.2% | -5.3% | -11.2% |
| 12 | 20 | H | 1 | 15.2 | 13.4 | 14.4 | 13.4 | -11.8% | -5.3% | -11.8% |
| 12 | 20 | H | 50001 | 15.2 | 13.5 | 14.4 | 13.5 | -11.2% | -5.3% | -11.2% |
| 13 | 20 | A | 1 | 20.0 | 17.3 | 18.7 | 17.3 | -13.5% | -6.5% | -13.5% |
| 13 | 20 | A | 50001 | 20.0 | 17.3 | 18.7 | 17.3 | -13.5% | -6.5% | -13.5% |
| 13 | 20 | H | 1 | 20.1 | 17.4 | 18.8 | 17.4 | -13.4% | -6.5% | -13.4% |
| 13 | 20 | H | 50001 | 20.1 | 17.4 | 18.8 | 17.4 | -13.4% | -6.5% | -13.4% |
| 14 | 20 | A | 1 | 20.1 | 17.4 | 18.8 | 17.4 | -13.4% | -6.5% | -13.4% |
| 14 | 20 | A | 50001 | 20.1 | 17.4 | 18.8 | 17.4 | -13.4% | -6.5% | -13.4% |
| 14 | 20 | H | 1 | 20.2 | 17.5 | 18.9 | 17.5 | -13.4% | -6.4% | -13.4% |
| 14 | 20 | H | 50001 | 20.2 | 17.5 | 18.9 | 17.5 | -13.4% | -6.4% | -13.4% |
| 15 | 20 | A | 1 | 20.8 | 18.2 | 19.5 | 18.2 | -12.5% | -6.2% | -12.5% |
| 15 | 20 | A | 50001 | 20.8 | 18.2 | 19.5 | 18.2 | -12.5% | -6.2% | -12.5% |
| 15 | 20 | H | 1 | 21.0 | 18.4 | 19.8 | 18.4 | -12.4% | -5.7% | -12.4% |
| 15 | 20 | H | 50001 | 21.0 | 18.4 | 19.8 | 18.4 | -12.4% | -5.7% | -12.4% |
| 16 | 20 | A | 1 | 22.6 | 20.1 | 21.4 | 20.1 | -11.1% | -5.3% | -11.1% |
| 16 | 20 | A | 50001 | 22.6 | 20.1 | 21.4 | 20.1 | -11.1% | -5.3% | -11.1% |
| 16 | 20 | H | 1 | 22.7 | 20.1 | 21.4 | 20.1 | -11.5% | -5.7% | -11.5% |
| 16 | 20 | H | 50001 | 22.7 | 20.1 | 21.4 | 20.1 | -11.5% | -5.7% | -11.5% |
| 17 | 20 | A | 1 | 27.5 | 24.3 | 26.0 | 24.3 | -11.6% | -5.5% | -11.6% |
| 17 | 20 | A | 50001 | 27.5 | 24.3 | 26.0 | 24.3 | -11.6% | -5.5% | -11.6% |
| 17 | 20 | H | 1 | 28.0 | 24.7 | 26.4 | 24.7 | -11.8% | -5.7% | -11.8% |
| 17 | 20 | H | 50001 | 28.0 | 24.8 | 26.4 | 24.8 | -11.4% | -5.7% | -11.4% |
| 18 | 20 | A | 1 | 29.7 | 26.4 | 28.1 | 26.4 | -11.1% | -5.4% | -11.1% |
| 18 | 20 | A | 50001 | 29.7 | 26.4 | 28.1 | 26.4 | -11.1% | -5.4% | -11.1% |
| 18 | 20 | H | 1 | 30.2 | 26.9 | 28.7 | 26.9 | -10.9% | -5.0% | -10.9% |
| 18 | 20 | H | 50001 | 30.2 | 26.9 | 28.7 | 26.9 | -10.9% | -5.0% | -10.9% |
| 19 | 20 | A | 1 | 30.0 | 26.8 | 28.4 | 26.8 | -10.7% | -5.3% | -10.7% |
| 19 | 20 | A | 50001 | 30.0 | 26.8 | 28.4 | 26.8 | -10.7% | -5.3% | -10.7% |
| 19 | 20 | H | 1 | 30.5 | 27.2 | 28.9 | 27.2 | -10.8% | -5.2% | -10.8% |
| 19 | 20 | H | 50001 | 30.5 | 27.2 | 28.9 | 27.2 | -10.8% | -5.2% | -10.8% |
| 20 | 20 | A | 1 | 38.5 | 35.6 | 37.0 | 35.6 | -7.5% | -3.9% | -7.5% |
| 20 | 20 | A | 50001 | 38.5 | 35.5 | 37.0 | 35.5 | -7.8% | -3.9% | -7.8% |
| 20 | 20 | H | 1 | 41.1 | 38.1 | 39.7 | 38.1 | -7.3% | -3.4% | -7.3% |
| 20 | 20 | H | 50001 | 41.1 | 38.1 | 39.7 | 38.1 | -7.3% | -3.4% | -7.3% |
| 21 | 20 | A | 1 | 39.1 | 35.8 | 37.6 | 35.8 | -8.4% | -3.8% | -8.4% |
| 21 | 20 | A | 50001 | 39.1 | 35.8 | 37.7 | 35.8 | -8.4% | -3.6% | -8.4% |
| 21 | 20 | H | 1 | 41.4 | 38.5 | 40.2 | 38.5 | -7.0% | -2.9% | -7.0% |
| 21 | 20 | H | 50001 | 41.5 | 38.5 | 40.2 | 38.5 | -7.2% | -3.1% | -7.2% |
| 22 | 20 | A | 1 | 40.1 | 36.4 | 38.4 | 36.4 | -9.2% | -4.2% | -9.2% |
| 22 | 20 | A | 50001 | 40.1 | 36.5 | 38.4 | 36.5 | -9.0% | -4.2% | -9.0% |
| 22 | 20 | H | 1 | 42.9 | 39.2 | 41.2 | 39.2 | -8.6% | -4.0% | -8.6% |
| 22 | 20 | H | 50001 | 42.9 | 39.2 | 41.2 | 39.2 | -8.6% | -4.0% | -8.6% |
| 23 | 20 | A | 1 | 40.3 | 36.7 | 38.5 | 36.7 | -8.9% | -4.5% | -8.9% |
| 23 | 20 | A | 50001 | 40.4 | 36.7 | 38.5 | 36.7 | -9.2% | -4.7% | -9.2% |
| 23 | 20 | H | 1 | 43.1 | 39.5 | 41.2 | 39.5 | -8.4% | -4.4% | -8.4% |
| 23 | 20 | H | 50001 | 43.1 | 39.5 | 41.2 | 39.5 | -8.4% | -4.4% | -8.4% |
| 24 | 20 | A | 1 | 40.5 | 38.0 | 38.9 | 38.0 | -6.2% | -4.0% | -6.2% |
| 24 | 20 | A | 50001 | 40.5 | 38.0 | 38.9 | 38.0 | -6.2% | -4.0% | -6.2% |
| 24 | 20 | H | 1 | 43.6 | 40.9 | 42.0 | 40.9 | -6.2% | -3.7% | -6.2% |
| 24 | 20 | H | 50001 | 43.6 | 40.9 | 42.0 | 40.9 | -6.2% | -3.7% | -6.2% |
| 25 | 20 | A | 1 | 40.7 | 39.0 | 40.0 | 39.0 | -4.2% | -1.7% | -4.2% |
| 25 | 20 | A | 50001 | 40.7 | 39.0 | 40.1 | 39.0 | -4.2% | -1.5% | -4.2% |
| 25 | 20 | H | 1 | 43.9 | 41.9 | 42.9 | 41.9 | -4.6% | -2.3% | -4.6% |
| 25 | 20 | H | 50001 | 43.8 | 41.9 | 42.9 | 41.9 | -4.3% | -2.1% | -4.3% |
| 26 | 20 | A | 1 | 42.1 | 35.5 | 41.2 | 39.4 | -15.7% | -2.1% | -6.4% |
| 26 | 20 | A | 50001 | 42.1 | 35.5 | 41.2 | 39.4 | -15.7% | -2.1% | -6.4% |
| 26 | 20 | H | 1 | 44.9 | 38.9 | 44.2 | 42.5 | -13.4% | -1.6% | -5.3% |
| 26 | 20 | H | 50001 | 44.9 | 38.9 | 44.2 | 42.5 | -13.4% | -1.6% | -5.3% |
| 27 | 20 | A | 1 | 43.1 | 37.6 | 36.8 | 39.7 | -12.8% | -14.6% | -7.9% |
| 27 | 20 | A | 50001 | 43.1 | 37.6 | 36.8 | 39.7 | -12.8% | -14.6% | -7.9% |
| 27 | 20 | H | 1 | 46.2 | 40.8 | 40.1 | 42.8 | -11.7% | -13.2% | -7.4% |
| 27 | 20 | H | 50001 | 46.2 | 40.8 | 40.1 | 42.8 | -11.7% | -13.2% | -7.4% |
| 28 | 20 | A | 1 | 44.3 | 45.4 | 44.9 | 41.9 | +2.5% | +1.4% | -5.4% |
| 28 | 20 | A | 50001 | 44.3 | 45.4 | 44.9 | 41.9 | +2.5% | +1.4% | -5.4% |
| 28 | 20 | H | 1 | 46.7 | 48.0 | 47.4 | 44.4 | +2.8% | +1.5% | -4.9% |
| 28 | 20 | H | 50001 | 46.7 | 48.0 | 47.3 | 44.4 | +2.8% | +1.3% | -4.9% |
| 29 | 20 | A | 1 | 45.9 | 45.4 | 46.6 | 43.5 | -1.1% | +1.5% | -5.2% |
| 29 | 20 | A | 50001 | 45.9 | 45.4 | 46.6 | 43.6 | -1.1% | +1.5% | -5.0% |
| 29 | 20 | H | 1 | 49.5 | 49.1 | 50.3 | 47.2 | -0.8% | +1.6% | -4.6% |
| 29 | 20 | H | 50001 | 49.5 | 49.1 | 50.3 | 47.2 | -0.8% | +1.6% | -4.6% |
| 30 | 30 | A | 1 | 63.3 | 60.0 | 61.7 | 60.0 | -5.2% | -2.5% | -5.2% |
| 30 | 30 | A | 50001 | 63.3 | 60.0 | 61.7 | 60.0 | -5.2% | -2.5% | -5.2% |
| 30 | 30 | H | 1 | 68.8 | 65.8 | 67.4 | 65.8 | -4.4% | -2.0% | -4.4% |
| 30 | 30 | H | 50001 | 68.8 | 65.7 | 67.3 | 65.7 | -4.5% | -2.2% | -4.5% |
| 31 | 30 | A | 1 | 63.7 | 55.8 | 62.1 | 62.2 | -12.4% | -2.5% | -2.4% |
| 31 | 30 | A | 50001 | 63.7 | 55.8 | 62.1 | 62.2 | -12.4% | -2.5% | -2.4% |
| 31 | 30 | H | 1 | 69.2 | 61.7 | 67.8 | 68.2 | -10.8% | -2.0% | -1.4% |
| 31 | 30 | H | 50001 | 69.2 | 61.6 | 67.8 | 68.2 | -11.0% | -2.0% | -1.4% |
| 32 | 30 | A | 1 | 65.2 | 58.6 | 58.8 | 65.0 | -10.1% | -9.8% | -0.3% |
| 32 | 30 | A | 50001 | 65.3 | 58.5 | 58.9 | 65.0 | -10.4% | -9.8% | -0.5% |
| 32 | 30 | H | 1 | 71.1 | 65.1 | 65.0 | 71.7 | -8.4% | -8.6% | +0.8% |
| 32 | 30 | H | 50001 | 71.1 | 65.1 | 65.0 | 71.7 | -8.4% | -8.6% | +0.8% |
| 33 | 30 | A | 1 | 61.2 | 60.7 | 61.1 | 66.8 | -0.8% | -0.2% | +9.2% |
| 33 | 30 | A | 50001 | 61.2 | 60.7 | 61.1 | 66.9 | -0.8% | -0.2% | +9.3% |
| 33 | 30 | H | 1 | 67.1 | 67.3 | 67.5 | 73.5 | +0.3% | +0.6% | +9.5% |
| 33 | 30 | H | 50001 | 67.1 | 67.4 | 67.5 | 73.5 | +0.4% | +0.6% | +9.5% |
| 34 | 30 | A | 1 | 63.0 | 61.9 | 62.7 | 61.9 | -1.7% | -0.5% | -1.7% |
| 34 | 30 | A | 50001 | 63.0 | 61.9 | 62.7 | 61.9 | -1.7% | -0.5% | -1.7% |
| 34 | 30 | H | 1 | 69.5 | 68.8 | 69.5 | 68.8 | -1.0% | +0.0% | -1.0% |
| 34 | 30 | H | 50001 | 69.5 | 68.8 | 69.5 | 68.8 | -1.0% | +0.0% | -1.0% |
| 35 | 30 | A | 1 | 65.0 | 63.6 | 64.3 | 63.6 | -2.2% | -1.1% | -2.2% |
| 35 | 30 | A | 50001 | 65.0 | 63.6 | 64.3 | 63.6 | -2.2% | -1.1% | -2.2% |
| 35 | 30 | H | 1 | 71.3 | 70.3 | 70.8 | 70.3 | -1.4% | -0.7% | -1.4% |
| 35 | 30 | H | 50001 | 71.3 | 70.3 | 70.8 | 70.3 | -1.4% | -0.7% | -1.4% |
| 36 | 30 | A | 1 | 68.9 | 67.2 | 68.1 | 67.2 | -2.5% | -1.2% | -2.5% |
| 36 | 30 | A | 50001 | 69.0 | 67.2 | 68.2 | 67.2 | -2.6% | -1.2% | -2.6% |
| 36 | 30 | H | 1 | 75.9 | 74.7 | 75.4 | 74.7 | -1.6% | -0.7% | -1.6% |
| 36 | 30 | H | 50001 | 75.8 | 74.7 | 75.3 | 74.7 | -1.5% | -0.7% | -1.5% |
| 37 | 30 | A | 1 | 70.2 | 68.5 | 69.4 | 68.5 | -2.4% | -1.1% | -2.4% |
| 37 | 30 | A | 50001 | 70.2 | 68.6 | 69.5 | 68.6 | -2.3% | -1.0% | -2.3% |
| 37 | 30 | H | 1 | 76.9 | 75.6 | 76.3 | 75.6 | -1.7% | -0.8% | -1.7% |
| 37 | 30 | H | 50001 | 76.9 | 75.6 | 76.3 | 75.6 | -1.7% | -0.8% | -1.7% |
| 38 | 30 | A | 1 | 70.9 | 69.2 | 70.1 | 69.2 | -2.4% | -1.1% | -2.4% |
| 38 | 30 | A | 50001 | 70.9 | 69.2 | 70.2 | 69.2 | -2.4% | -1.0% | -2.4% |
| 38 | 30 | H | 1 | 78.1 | 76.9 | 77.5 | 76.9 | -1.5% | -0.8% | -1.5% |
| 38 | 30 | H | 50001 | 78.0 | 76.8 | 77.5 | 76.8 | -1.5% | -0.6% | -1.5% |
| 39 | 30 | A | 1 | 71.6 | 69.9 | 70.8 | 69.9 | -2.4% | -1.1% | -2.4% |
| 39 | 30 | A | 50001 | 71.6 | 70.0 | 70.9 | 70.0 | -2.2% | -1.0% | -2.2% |
| 39 | 30 | H | 1 | 79.0 | 77.9 | 78.5 | 77.9 | -1.4% | -0.6% | -1.4% |
| 39 | 30 | H | 50001 | 79.0 | 77.8 | 78.5 | 77.8 | -1.5% | -0.6% | -1.5% |
| 40 | 40 | A | 1 | 93.3 | 150.5 | 93.2 | 150.5 | +61.3% | -0.1% | +61.3% |
| 40 | 40 | A | 50001 | 93.4 | 150.7 | 93.3 | 150.7 | +61.3% | -0.1% | +61.3% |
| 40 | 40 | H | 1 | 104.6 | 169.1 | 104.6 | 169.1 | +61.7% | +0.0% | +61.7% |
| 40 | 40 | H | 50001 | 104.5 | 169.3 | 104.5 | 169.3 | +62.0% | +0.0% | +62.0% |
| 41 | 40 | A | 1 | 96.8 | 152.7 | 152.7 | 152.7 | +57.7% | +57.7% | +57.7% |
| 41 | 40 | A | 50001 | 96.9 | 152.9 | 152.9 | 152.9 | +57.8% | +57.8% | +57.8% |
| 41 | 40 | H | 1 | 106.3 | 169.7 | 169.7 | 169.7 | +59.6% | +59.6% | +59.6% |
| 41 | 40 | H | 50001 | 106.3 | 169.9 | 169.9 | 169.9 | +59.8% | +59.8% | +59.8% |
| 42 | 40 | A | 1 | 157.7 | 157.7 | 157.7 | 157.7 | +0.0% | +0.0% | +0.0% |
| 42 | 40 | A | 50001 | 157.8 | 157.8 | 157.8 | 157.8 | +0.0% | +0.0% | +0.0% |
| 42 | 40 | H | 1 | 176.1 | 176.1 | 176.1 | 176.1 | +0.0% | +0.0% | +0.0% |
| 42 | 40 | H | 50001 | 176.3 | 176.3 | 176.3 | 176.3 | +0.0% | +0.0% | +0.0% |
| 43 | 40 | A | 1 | 159.6 | 159.6 | 159.6 | 159.6 | +0.0% | +0.0% | +0.0% |
| 43 | 40 | A | 50001 | 159.9 | 159.9 | 159.9 | 159.9 | +0.0% | +0.0% | +0.0% |
| 43 | 40 | H | 1 | 177.8 | 177.8 | 177.8 | 177.8 | +0.0% | +0.0% | +0.0% |
| 43 | 40 | H | 50001 | 178.0 | 178.0 | 178.0 | 178.0 | +0.0% | +0.0% | +0.0% |
| 44 | 40 | A | 1 | 166.0 | 166.0 | 166.0 | 166.0 | +0.0% | +0.0% | +0.0% |
| 44 | 40 | A | 50001 | 166.2 | 166.2 | 166.2 | 166.2 | +0.0% | +0.0% | +0.0% |
| 44 | 40 | H | 1 | 188.2 | 188.2 | 188.2 | 188.2 | +0.0% | +0.0% | +0.0% |
| 44 | 40 | H | 50001 | 188.4 | 188.4 | 188.4 | 188.4 | +0.0% | +0.0% | +0.0% |

### Web research (Forever sources first)

- Kargoz, Forever 1-60 Shadow priest leveling (Mobalytics), seen through search 2026-11-06: Wand Specialization into Shadow talents to 39, "Respec at 40 to get Shadowform ASAP"; without a reset, skip the early wand and take it after Shadowform. https://mobalytics.gg/wow-forever/profile/kargoz/classes/1-60-shadow-priest-leveling-guide-wands-talents-rotation-tips-and-tricks
- Method, WoW Forever priest leveling, last updated 2026-11-06: "At Level 40 we can now grab Shadowform, so I recommend re-speccing to Shadow." https://www.method.gg/wow-forever/wow-forever-priest-leveling-guide-and-talents
- Icy Veins, Forever Shadow priest leveling (round 15 citation, 2026-11-02): Wand Specialization first; reset at 40 or 42 (search summary, page answers 403 here). https://www.icy-veins.com/wow-forever/shadow-priest-leveling
- Warcraft Tavern Forever priest leveling: Shadowform at 40; builds still capped at the level-30 beta. https://www.warcrafttavern.com/forever/guides/priest-leveling/

Consensus: early wand, then a reset at 40. Among single sequences, the early wand wins, which is what the measurement says too.

Live logs: last looked 2026-11-06 18:32 UTC (api.foreversixty.gg/v1/reports/recent: one 20 Oct report, no Shadow leveling data). No log can settle a leveling talent order; nothing left open on logs.

## Before/after

No engine, APL, ladder or build change: every published number is unchanged (0.0% at every band). The guide change is prose only.

## Files to regenerate

None. (Had (b) won, the ladder-first row change would have needed the ladder goldens, the addon-builds data and the planner vectors regenerated.)

## Measure at refresh

- The frozen BiS copy (`$R/bis-frozen/priest-shadow.json`) carries band talents that are not the guide ladder (band 20 `...-54200...`, band 40 `...-54342001201300051` with Shadowform and no wand) and races gnome / undead, while the live `data/builds/1.60.1.70790/bis/priest-shadow.json` carries the ladder strings. After the refresh, check that the published Shadow band talents equal the guide ladder at 20, 30 and 40 (wand at 20 and 30, no Shadowform at 40) and that the alliance race is the guide's Human.

## For other lanes

- **platform (SH5):** bare band 20 strips the band wand below 17 and leaves the ranged slot empty, so priest (and by the same code mage and warlock) runs at 10-16 have no wand. Fold casters into the "unarmed bare bands" fix: give the level the best legal wand the band's sources offer (Lesser Magic Wand 11287 from 5, Greater Magic Wand 11288 from 13 are crafted and always available).

## For the guide lane

- shadow.md Weaknesses still says "Shadowform at 42", which stays right for the point order. The new leveling sentence offers the reset at 40.

## Still open

Nothing for this seed.
