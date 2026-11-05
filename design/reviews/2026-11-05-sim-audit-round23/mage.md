# Sim-accuracy round 23: mage

Specs: mage-arcane, mage-fire, mage-frost. Site wt/r23-mage (from int/r23 112ab3197). Fork lane/r23-mage (from int/r23). Client 1.60.1.70790.

Method:
- Every sim ran through `$R/sim.sh`. I explored at 3,000 iterations and used 10,000 at seeds 1 and 50001, both factions, for every reported figure.
- Bare bands were run level by level with `-level N` and `-talents` set to the ladder build for that level. I printed each ladder from the site's own `leveling.LadderTalentString`, using a scratch tool that was not committed. Gear came from the frozen BiS (`-bis-dir $R/bis-frozen`).
- Raid 60 used `$R/rt` against `$R/before.tsv`.

**Confirmed and fixed: 2.**
- F1 (`old`, round 22 open item): the Fire level 60 build and its ladder.
- A1 (`old`, guides-handoff): Arcane Presence of Mind on the raid order.

Frost: aligned.

## Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **The Fire level 60 (leveling) build left damage on the table at 60, and its points did nothing in the 20s.**<ul><li>**What was wrong.** The level 60 build was `2520150000002/51520101230133031/003`. It spent its Combustion toll on Flame Throwing 2 (range) and Burning Soul 1 (pushback and threat), plus Arcane Resilience, Magic Absorption and Arcane Meditation. With the raid build's talents, bare 60 does +2.3% (round 22, Still open).</li><li>**Just using the raid build is worse while leveling.** The plain walk spends Accelerant's first-row ranks at 16 to 19. That pushes Heating Up to 33 and Combustion to 43: -13% at 30 and -6% at 40. It also drops Wand Specialization, which a leveling mage uses while it drinks (Shoot is 7 to 19% of bare-band damage): -1.7% at 55.</li><li>**New level 60 build:** `2301052001000/55500201230133031/003`. It is the raid build with Wand Specialization 2/2 in place of two Improved Channeling ranks. Improved Channeling is pushback only, for Arcane Missiles, Blast and Wrath (client text).</li><li>**New ladder order** (data/curated/ladder-first.json mage-fire):<ul><li>Pyroblast at 20, Heating Up at 30 and Combustion at 40 stay where they were.</li><li>Accelerant 2 to 4 moves to 22 to 24, where Flame Throwing and Burning Soul were. Fireball is the filler below 30, so this is a 3.1 s Fireball.</li><li>Then Accelerant 5 and Burning Soul 2 at 41 to 43, Wand Specialization at 44 and 45, Improved Channeling 3, Subtlety 1, Arcane Concentration from 50, Missile Barrage at 55 and 56, Arcane Blast at 57 and Precision at 58 to 60.</li></ul></li><li>**Adoption rule:**<ul><li>Fire holds the most points.</li><li>Nothing the rotation casts is dropped.</li><li>It gains well over 1% at 60.</li><li>Burning Soul, an unmodeled talent, is kept and grows to 2.</li><li>The points that go do nothing in the rotation (Flame Throwing, Arcane Resilience, Magic Absorption). This is the Shadow precedent, as with round 22's raid build.</li><li>Arcane Meditation is modeled, and the sim prices it.</li></ul></li><li>**Fillers checked on the new ladder:**<ul><li>Scorch beats a Fireball filler at 33 to 55 by 6 to 26%.</li><li>Fireball beats a Scorch filler at 57 and 60 by 4%.</li><li>The `spellIsKnown(Arcane Blast)` gates therefore stay. Only the notes name the levels.</li></ul></li><li>**Forever sources disagree on the order.** wowf.io takes Accelerant 5 at 10 to 14 and no Wand Specialization (https://wowf.io/en/leveling/mage/fire). Mobalytics takes Wand Specialization at 50 and 51 (https://mobalytics.gg/wow-forever/classes/fire-mage-guide-progress). Both were read on 5 Nov 2026. The sims decided.</li></ul> | runs `$R/mage/lad/{OLD,D}.*.10k`; tests `TestMageLadderTakesItsCastTalentsEarly` (fire cases) and vitest `Fire levels into the Fireball build` (builds.test.ts): red, then green; `make apl-check`, test_apl.py, `_prose.test.ts` pass | CONFIRMED-BUG (build defect under the adoption rule) | old (round 22 found it and left it open) | S ffe450d15, 5b191be1f; E 9398a2711, 174b5e310 |
| F2 | Guide-accuracy round 24, sims requested 2: the Fireball filler against a Frostbolt filler at 20, 22 and 24, on 60 s and 180 s fights. Fireball wins by 10 to 18% at every level, length, faction and seed. | `$R/mage/fire-frostbolt-filler.txt` | Aligned | - | - |
| F3 | Re-read of round 22's raid changes against the client text:<ul><li>Accelerant: -0.1 s a rank on Fireball, Frostfire Bolt and Arcane Wrath (`talents.go:799`).</li><li>Incineration: 2% a rank on Fire Blast, Ice Lance, Arcane Blast and Scorch (`talents.go:783`).</li><li>Master of Elements: Fire and Frost crits only (`talents.go:405`).</li><li>Heating Up and Critical Mass: unchanged since round 22.</li></ul>Raid 60 is unchanged this round: horde seed 1 is 712.4 against 712.4 in before.tsv. | engine files above | Aligned | - | - |

## Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **On the raid order, Presence of Mind went to Frostbolt from 32 to 53 while a longer Arcane Wrath was off cooldown** (guide-accuracy round 24, sims requested 3).<ul><li>The raid order takes Arcane Wrath at 31 and Presence of Mind at 32. Arcane Wrath is 3.5 s, 3.3 s once Accelerant is in, on a 6 s cooldown. Frostbolt is 3 s.</li><li>The new row puts Presence of Mind on Arcane Wrath when it is ready and Pyroblast is not known. Frostbolt still takes it when Arcane Wrath is on cooldown.</li><li>Gain on bare bands with the raid order, 10,000 iterations, seeds 1 and 50001, both factions: +0.40 to +0.88% at 32, 40, 48 and 53.</li><li>The leveling order takes no Arcane Wrath, so it is unchanged (45 and 60 identical). Raid 60 is unchanged (Pyroblast is known): 698.8 against 698.8.</li><li>The notes and arcane.md:56 now name Arcane Wrath.</li></ul> | `$R/mage/arc-leads.txt`; test `TestArcanePresenceOfMindTakesArcaneWrathBeforeFrostbolt`: red, then green | CONFIRMED-BUG (rotation) | old, guides-handoff (round 24) | S 5b191be1f; E 174b5e310 |
| A2 | Guide round 24 sim finding: the Arcane prepull keys on Heating Up, so a Pyroblast prepull from 54 instead of a Frostbolt was tested at 54 to 56. Alliance gains 0.3 to 0.7%. Horde gains 0.1% at 54 and 55 and loses 0.4% at 56 (seed 1). Not a gain at every level and faction, so it is held. | `$R/mage/arc-leads.txt` | Aligned (no change) | - | - |

## Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Ice Lance's coefficient at 60 still has no live log. | Logs | OPEN (log needed) | - | - |
| R2 | Nothing changed in the Frost APL or engine. Raid 60 is unchanged. | before.tsv | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Blizzard's Forever notes** (thread 2360696, read 5 Nov). The newest post is 4 Nov 19:35 (UI and accessibility) and has no mage line. Every 3 Nov mage line (Missile Barrage, Clearcasting, Arcane Blast +100%, Accelerant with Arcane Wrath, Critical Mass 3, Incineration 5, Ice Lance 8 s shared and 300%, Winter's Chill 20%, Precision 3) was checked in round 22 X1 and is unchanged in the engine. | https://us.forums.blizzard.com/en/wow/t/2360696 | Aligned | - | - |
| X2 | Arcane Subtlety's resistance cut (SIM-R7-1, carried) is still not modeled. The engine subtracts SpellPenetration from the target's own resistance (`core/spell_resistances.go:146`), not the level-based part, and the raid target has no school resistance, so the point does nothing on our fight. | engine | OPEN (carried) | - | - |

## Logs

Last looked 12:37 UTC 5 Nov.
- api.foreversixty.gg `/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct).
- foreverlogs.gg `/api/reports/public` answers 403.
- The WCL key was not tried (401 per the integrator).

## Before/after

Bare bands, horde seed 1, 10,000 iterations, the old ladder against the new:

| Level | Before | After |
|---|---|---|
| 22 | 39.1 | 39.5 |
| 24 | 42.4 | 43.1 |
| 25 | 42.7 | 43.5 |
| 27 | 47.2 | 47.6 |
| 50 | 197.8 | 201.0 |
| 57 | 248.6 | 252.7 |
| 60 | 340.6 | 353.4 |

Relative change, all four runs (both factions, seeds 1 and 50001):
- 22: +0.7 to +1.0%. 24 and 25: +1.4 to +2.7%. 27: +0.8 to +1.0%. 50: +1.4 to +1.6%. 57: +1.5 to +1.8%. 60: +2.3 to +3.8%.
- 10, 20, 30, 35, 38, 40, 45 and 55: unchanged.
- **41 to 43: -0.2 to -0.6%.** Wand Specialization moves from 41 and 42 to 44 and 45, because the ladder walk spends the spec's own tree first. These are not checkpoint levels, and no checkpoint loses. Moving off-tree points ahead of own-tree points would need a sim/leveling change.

Arcane raid order (card only): +0.4 to +0.9% at 32 to 53. Raid 60, all three specs: unchanged. No spec outside mage moves.

## Files to regenerate

- addon-builds.json: the mage-fire leveling order.
- addon-data.json and Data.lua: the mage-fire notes, plus the new mage-arcane row and its notes. Both were regenerated locally only to preview the page, then restored.
- The ladder golden sim/request/testdata/ladder/mage-fire.golden.md. Every rung from 22 changes talents; 60 becomes the new build.
- The coach starter data/curated/coach/mage-fire.md and mage-arcane.md, if persona_starter.py --all reads them.
- The mage-fire bare BiS bands, which carry the old ladder talents (see Measure at refresh).

## Measure at refresh

- The mage-fire bare-band BiS and weights on the new ladder. Fireball, Arcane Blast and Missiles at 57 to 60 change the 60 band's stat picture.
- Carried: the Fire raid BiS on round 22's build, the Arcane re-rank, and the Frost re-rank after Ice Lance's 0.107.

## For other lanes

- **Platform (component), from the wow-player SHOULDs:**
  - The guide's "Level 60 build" label collides with the Level 60 BiS band. "Leveling build" would match the Arcane prose.
  - The generated line "The raid build moves 2 points from Wand Specialization (2) to Improved Channeling (2)" reads backwards. Suggested wording: "2 more points in Improved Channeling, taken from Wand Specialization".
- **sim/leveling:** off-tree reservations are always spent before the spec's own tree (`reserveOffTree`). A spec cannot take an off-tree point (Fire's Wand Specialization) between two own-tree points. This costs Fire 0.2 to 0.6% at 41 to 43.

## For the guide lane

- Done here:
  - fire.md: the Talents last paragraph and the Leveling 20-60 paragraphs (Accelerant at 22 to 24, why Scorch takes over at 30, the order after 40, and Missile Barrage waiting for 57).
  - The mage-fire Scorch-filler and Fireball notes.
  - arcane.md:56 and the three Presence of Mind notes.
- wow-player SHIP on the rendered /guides/mage/fire and the Arcane Presence of Mind line, on the second pass. First pass: NO-SHIP on capture hydration and three wording MUSTs, all taken.
- SHOULD left for the accuracy reviewer: the page ties Spell power and Fire power on the weight (0.61) but its prose ranks one over the other.

## Still open

- Ice Lance's coefficient at 60 (needs a level 56+ Frost log). Arcane Power's cost (needs a mana log).
- Arcane Subtlety's resistance cut (SIM-R7-1).
- The Fire dip at 41 to 43 (needs the sim/leveling change above).
- Carried: the level-60 book ranks (pending owner).
