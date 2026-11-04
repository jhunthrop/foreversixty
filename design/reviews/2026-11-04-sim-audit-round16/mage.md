# Sim-accuracy round 16: mage

Specs: mage-arcane, mage-fire, mage-frost. Site wt/r16-mage (from int/r16 85c7536d1, later merged with int/r16 for wave 2), fork lane/r16-mage (from int/r16 913221285, later merged). Client 1.60.1.70755 with the late re-layer.

Method:
- Seed items first, then the re-audit.
- Raid 60 numbers: spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions, `-bis-dir $R/bis-frozen`, guide raid builds (`ALLRT=1`). Exploration at 3,000. Every run went through `$R/sim.sh`.
- Wave 2: a level-by-level sweep with `-level`, done after merging int/r16.

**Confirmed and fixed: 5.** By origin: R15 3, old 2. Two are no-DPS consistency fixes. Raid 60 moves only for Arcane (+0.4 to +0.6%).

Guide text: round 15's mage prose had never landed (guide-accuracy round 10, eabba8f84 on land/guides-r10, was not merged). This lane carries it over unchanged (S dc14f242e) and adds the opener sentence for G1. Wow-player review: see "Guide review" below.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Winter's Chill's spell power on Cone of Cold skipped the low-rank penalty.** Each stack adds 7% of spell power to the next Cone of Cold. The code added that raw, so a low rank kept all of it while the rank's own coefficient lost its level-gap share. It now takes the same penalty as the rank (`core.SpellLevelPenalty`; after the wave-2 merge it keys on the client MaxLevel). No rotation casts a low rank, so nothing moves. | `cone_of_cold.go` `wintersChillConeOfColdCoefficient`. `TestWintersChillCoefficientTakesTheRankLevelPenalty` failed (no such hook), then passed. | CONFIRMED-BUG (no DPS change) | R15 8d9ca6aae: the penalty rule did not reach this hand-added coefficient | E 3b1966bb6, fddf01ae2 |
| S2 | **Read Leyline stays out.** Alliance Skyborne (high-order-skyborne) Fire at raid 60, 10,000 iterations: <ul><li>no line: 727.8 / 729.1</li><li>Read Leyline on cooldown: 717.3 / 718.7 (-1.4%)</li><li>below 50% mana: 723.1 / 723.2 (-0.6 / -0.8%)</li></ul>Arcane and Frost below 50% mana (3,000 iterations): -0.5% and -0.3%. It costs a global, and Mage Armor already keeps half the regeneration running while casting. | runs | Aligned | - | - |
| S3 | **Round 15 regression targets, all hold.** <ul><li>The Arcane Fire-second raid build and its four lines: decoded build 32/16/3, rows [4] [7] [8] [9] [11] present. Raid 60 equals before.tsv to the decimal (717.6 / 717.1 / 704.3 / 705.1).</li><li>Arcane Power and Presence of Mind wait 3 s: `TestArcaneHoldsArcanePowerForTheTanksFirstSunders` passes.</li><li>Frost Arcane Impact 3: in both Frost builds (`230005203/5/253510130101301221`).</li><li>Improved Frostbolt first, and Arcane Concentration then Blast: `TestMageLadderTakesItsCastTalentsEarly` and the curated ladder-first tests pass.</li><li>Scorch rank 2 at 40-49: the row is present. Wave 2 later splits it at 46 (W1).</li><li>Burning Soul: Fire's leveling build has it; builds.test.ts passes.</li></ul>One regression is G1. | tests above; `$R/mage/after0.tsv` | Aligned (+ G1) | - | - |

## Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 (G1) | **The raid build's pre-pull Frostbolt landed half a second into the fight.** Round 15 (A5) kept the -2.5 s pre-pull Frostbolt because "both builds take Improved Frostbolt". Its wave 4 then gave the raid build a Fire second tree and dropped Improved Frostbolt. That makes Frostbolt rank 10 a 3.0 s cast (spellconst 10181, cast 3000). The guide itself said "3 seconds with the raid build". Now, with Heating Up known, the opener is Pyroblast 6 s out, the way Fire opens: its crit is the first Ignite, and it is the longest cast. Without Heating Up (the leveling build, Improved Frostbolt 5/5) it stays a 2.5 s Frostbolt. Raid 60 pre-pull options, alliance seed 1, 3,000 iterations: <ul><li>none 713.9</li><li>Frostbolt at -2.5 s 717.7</li><li>Frostbolt at -3 s 719.3</li><li>Fireball at -3.3 s 719.2</li><li>Pyroblast at -6 s 721.0</li></ul>Web: Warcraft Tavern's Forever Arcane guide gives no opener (https://www.warcrafttavern.com/forever/guides/arcane-mage/, no date, read 2026-11-04). So the guide's own Fire opener and the measurement decide. | `TestArcaneOpensWithPyroblastOnTheRaidBuildAndFrostboltLeveling` failed, then passed. The guide card renders "Pyroblast 6 s before the pull, with Heating Up" and "Frostbolt 2.5 s before the pull, without Heating Up". The opener prose in arcane.md is updated. | CONFIRMED-BUG | R15 e01dc2dff | S d5eb38628, E ba5899b5b |
| A2 (W2) | **A phantom line (shaman lane's flag).** Line 4 casts Pyroblast while Presence of Mind is up, gated on Heating Up being known. The engine drops an `auraIsActive` on an aura it never registered. So a build with Heating Up and no Presence of Mind (any custom build in the site's sim) would cast Pyroblast from this line on every global. It now checks `spellIsKnown(12043)` too, as line 5 already did. No published build has Heating Up without Presence of Mind, so raid and bare bands are identical. Every other mage line was checked: each one that reads a talent aura already checks the talent, or reads a spell that only exists with the talent. Frost's Ice Barrier rank 4 aura is rewritten per level: 3 casts in 180 s at 40 and 45, 93% uptime. | `TestArcanePresenceOfMindLinesCheckTheTalent` failed, then passed | CONFIRMED-BUG (latent) | R15 e01dc2dff | S 3fb5f255e, E 7c40e95dd |
| A3 | The Arcane rotation re-read at raid 60. <ul><li>Arcane Wrath fires once per Blast set. That is because the GCD budget is full (cast time per iteration sums to 178 of 180 s), not a cooldown bug.</li><li>Crit rates line up: Fire Blast base +10 (Incineration), Arcane +6 (Arcane Impact), Blast +16.</li><li>Clearcasting as a free fourth Blast: 717.7 against 717.7, a tie, so no line.</li><li>Fire Blast on cooldown in the leveling build: bare 30 to 60 loses 1 to 6% (mana), so no line.</li></ul> | runs | Aligned | - | - |
| A4 | Every Arcane spell and talent at 60 against the 70755 rows: <ul><li>Arcane Blast and its 400573 buff (12% Intellect, +100% cost, 18 s);</li><li>Arcane Wrath (451 + 3.0/level, coefficient 1.0, 3.5 s, 6 s, 395 mana, 5% max mana);</li><li>Missiles 0.286 a missile;</li><li>Missile Barrage 40/20%;</li><li>Presence of Mind, Arcane Power 30/30;</li><li>Accelerant on Arcane Wrath;</li><li>Eureka! 1259817.</li></ul>The conformance golden has 0 mismatches. spellconst equals the live push (`push-check` moved=false at 05:00). | `cmpconst.py`; conformance mage.golden.md and TALENTS.md | Aligned | - | - |

## Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 (W1) | **Scorch rank 2 lost its spell power from 45 but stayed the row to 49.** Swept level by level, bare: <ul><li>ranks 2 and 3 at 44: 168.9 against 164.8 (rank 2 wins)</li><li>at 45: 167.9 against 166.4 (rank 2 wins)</li><li>at 46: 167.0 against 167.5 (a tie)</li><li>at 47: 165.9 against 168.9 (+1.8%)</li><li>at 49: 166.4 against 173.6 (+4.3%)</li></ul>Seed 50001 is the same; horde ties at 46. The row is now rank 2 at 40-45 and rank 3 at 46-49. | `TestRound16MageRankRowsFollowTheLevel` failed, then passed | CONFIRMED-BUG | R15 c094ddd7f (measured at the band floor only) | S bb66d5be6 |
| F2 | The other rank rows hold at every level: <ul><li>Fire Blast rank 3 at 30-39 wins every swept level (77.3 against 76.2 at 30, 98.1 against 96.1 at 39);</li><li>Fire Blast rank 4 at 40-49 wins to 48 and ties at 49 (166.0 / 166.2 against 166.2 / 166.6);</li><li>Scorch rank 1 at 30-39 wins at every level (98.1 against 97.1 at 39);</li><li>bands 20 and 50 defaults win.</li></ul> | sweep | Aligned | - | - |
| F3 | Scorch stays the filler: Fireball over Scorch is 644.9 against 735.1. Flame Throwing 2 is kept on purpose (raid-keeps.json, range for the spells Fire casts). The re-read of Ignite, Heating Up, Combustion, Fire Power and Pyroblast found nothing new. | runs, TALENTS.md | Aligned | - | - |

## Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 (W1) | **Frost had no Frostbolt rank row, and the highest rank lost mana-bound bands.** Bare, seeds 1 and 50001, both factions: <ul><li>rank 4 against rank 5 at 26: 56.0 against 52.9 (horde 52.7 against 50.2);</li><li>at 29: 58.4 against 56.7;</li><li>rank 5 against rank 6 at 32: 84.1 against 80.6;</li><li>at 33: 84.5 against 81.8;</li><li>at 34, with Mage Armor, rank 6 wins: 93.3 against 88.3;</li><li>at 30-31 the default rank 5 wins (83.3 against 82.2).</li></ul>New rows: rank 4 at 26-29 and rank 5 at 32-33 (+3 to +6%). | `TestRound16MageRankRowsFollowTheLevel` | CONFIRMED-BUG | old (no round swept Frost below 40 per level) | S bb66d5be6 |
| R2 | Ice Barrier against 70755: the shield's 5% and Frigid Power's 15% (1326891) match `ice_barrier.go`. Frost Warding's absorb is unmodeled and does nothing in a fight where the mage takes no hits (a row filler). Frostbite never freezes a boss (`chill.go`). Ice Lance's coefficient is still the carried open item. | spellconst; engine | Aligned | - | - |

## Arcane (wave 2 rank row)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W3 | **The Arcane Frostbolt rank 4 row ran to 39.** It wins at 30-33. From 34 (Mage Armor) the highest rank wins: 103.4 against 102.9 at 34, 104.3 against 103.5 at 36, 104.1 against 102.5 at 37, and +6% at 39 (seed 50001 the same). Rank 4 also loses spell power from 37 (MaxLevel 24 + 12). Horde ties at 34. The row now ends at 33. Arcane bands 20, 40 and 50 defaults win. | `TestRound16MageRankRowsFollowTheLevel` | CONFIRMED-BUG | old (round 13 row measured at 30 only) | S bb66d5be6 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The 70755 push for mage: <ul><li>`python -m pipeline push-check --build 1.60.1.70755` gives moved=false;</li><li>mage spellconst is byte-identical to the fork's conformance copy;</li><li>no mage trait curve differs between the shipped and layered tables (all 40 diffs are other classes).</li></ul> | `curvediff2.py`, `cmpconst.py` | Aligned | - | - |
| X2 | `talents.go` said wand casts are not modeled. `shoot.go` applies Wand Specialization. The comment is fixed. | - | comment only | - | E f637ec32c |

## Before/after (raid 60, 10,000 iterations)

Before: `$R/before.tsv`. After: E f637ec32c on S bb66d5be6.

| Spec | Faction | s1 before → after | s1 % | s50001 % |
|---|---|---|---|---|
| mage-arcane | alliance | 717.6 → 721.0 | +0.47 | +0.39 |
| mage-arcane | horde | 704.3 → 708.3 | +0.57 | +0.47 |
| mage-fire | both | 735.6 / 721.0 unchanged | 0 | 0 |
| mage-frost | both | 696.0 / 686.9 unchanged | 0 | 0 |

Below 60, the wave-2 rows gain:
- Frost: +3 to +6% at 26-29 and 32-33.
- Arcane: +0.3 to +6% at 34-39.
- Fire: +1.8 to +4.3% at 47-49.

No spec outside mage moves. No shared code was changed.

## Files to regenerate

- Ladder goldens:
  - mage-arcane: rungs 35 and 38, where Frostbolt rank 4 becomes rank 6 or 7.
  - The 60 rung's opener is now Pyroblast.
  - `TestRotationLadder` was already red at the merge (126 violations, not this lane).
- Addon data: addon-data.json and Data.lua (rank rows, the Arcane opener, the Arcane notes from guide round 10).
- The guide's card clauses for Arcane lines 4, 7, 8, 9 and 11, which render without a clause until regenerated.
- Coach starters for mage-arcane.

## Measure at refresh

- Arcane raid 60 on the Pyroblast opener (BiS and weights; tiny).

## For other lanes

- None.

## For the guide lane

- Done here: round 10's mage text, and the Arcane opener sentence (arcane.md:72).

## Guide review

The wow-player agent reviewed the rendered arcane, fire and frost pages, served by astro dev at /guides/mage/*. The captures are in `$R/mage/shots/`.

**First pass: NO-SHIP.** The reviewer raised these points, all fixed in S 529a6bfca:
- The Arcane "Brings" line read as if Arcane Subtlety were only a resistance cut.
- The raid-build point arithmetic was unclear.
- "Before you have the Fire points" is wrong for a leveling build that never takes them. It now says "with the leveling build", in the prose and in two card notes.
- The Frost line on Arcane Impact needed checking. It now follows the client text: crit on Arcane spells, which for Frost means the free Missiles.

**Second pass: SHIP**, on one condition: the talent trees must draw. They load lazily, and a capture after the scroll shows both Arcane trees drawn (`arcane-trees3.png`).

The prose test (`_prose.test.ts`) and builds.test.ts pass.

## Still open

Carried with no new evidence:
- Ice Lance's coefficient (client 0, engine 0.429 by convention).
- Arcane Subtlety's 15 resistance cut (SIM-R7-1).
- Arcane Concentration from Shoot and the Kindling Stave's Ignition.
- The level-60 book ranks (pending owner).
- Warcraft Logs is blocked.

## Wave 2

Wave 2 ran after merging int/r16, for spec-breakdown `-level` and the MaxLevel + 12 penalty.

1. **Per-level rank rows.** Every mage rank row was swept at the band floor, +3, +6 and +9, and at each switch level. The Frostbolt default was also swept in Arcane and Frost bands 20 to 50, and Fire Blast and Scorch in Fire bands 20 to 50. Switches were confirmed at 10,000 iterations, seeds 1 and 50001, alliance and horde. Changes, all in S bb66d5be6:
   - Frost: Frostbolt rank 4 at 26-29 and rank 5 at 32-33 (R1, new rows).
   - Arcane: the Frostbolt rank 4 row now ends at 33 (W3).
   - Fire: Scorch rank 2 at 40-45, then rank 3 at 46-49 (F1).
   - Fire Blast rows hold (F2).
   - No race key: horde agrees or ties at every switch level.
2. **Phantom talent lines.** Arcane line 4 (Pyroblast while Presence of Mind is up) now also checks `spellIsKnown(12043)` (A2; S 3fb5f255e, E 7c40e95dd). Line 5 already checked it. No other mage line reads a talent aura without its talent gate or a talent-only spell.
3. **Text.** The rank-choice notes are updated. No guide prose names ranks. The rendered review is above.
