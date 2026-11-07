# Sim-accuracy round 28: hybrids lane

Site `wt/r28-hybrids` from int/r28 eb22d9382. Fork `lane/r28-hybrids` **not created** (no engine bug found; measured on the base engine 978e4c8aa). Client 1.60.1.70790. Ship mode: touched tests only, no regens. Every sim went through `$R/sim.sh`, one at a time.

**Result: 4 confirmed bugs, all fixed (all `old`), plus one rank row that the new order moved and one race row it changed.** The leveling orders of Feral, Survival and Protection spent points that do nothing in the fight while a live talent of an open row waited. None of the band 20 builds move. Survival's band 30 build moves (re-rank). Feral's and Protection's band 20, 30 and 40 builds do not.

Method (round 27 orders.md): `spec-breakdown -preset bare -band <20|30> -level N -talents <walk string> -bis-dir $R/bis-frozen`, old walk against new walk at every level the change touches, 10,000 iterations, seeds 1 and 50001, both factions' published races (Feral: Night Elf / Tauren; Survival: Dwarf / Troll; Protection: Dwarf / Undead). Protection is judged on threat on the tank fight (`-tank`). Feral's order was found by a greedy level-by-level sweep (seed 1, every legal next point), then confirmed at both seeds and factions. Outputs: `$R/hybrids/out/` (`feral-final.tsv`, `feral-optC.tsv`, `feral-optB.tsv`, `feral-greedy-*.txt`, `hunter-*.tsv`, `prot-orders.tsv`, `sv-raptor-*.txt`).

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Seed: Feral Swiftness (15, 16) before Feral Instinct (17). Below 20 the sim's druid is a caster (Moonfire and Wrath; no Cat Form until 20), so every Feral point at 10 to 19 does nothing in the fight, and at 20 all five row-two points are in either way. The order only matters for real play: the guide levels 10 to 19 in Bear Form, where Feral Swiftness's dodge counts and Feral Instinct's Swipe bonus matters only on packs. Order only; no checkpoint moves. | `talents/druid.json` 17002, 16947; APL Moonfire/Wrath lines gated `not spellIsKnown(Claw)` | Aligned (order only) | - | - |
| F2 | **Savage Fury at 23 and 24 before Sharpened Claws.** Sharpened Claws (3% crit a point in Cat Form) waited until 26 and 27. Now Sharpened Claws 23 and 24, Feral Charge 25, Savage Fury 26 and 27: +1.4% at 23, +2.7% at 24, +1.3% at 26, both seeds and factions. Taking Sharpened Claws even earlier, at 20, gains 3.3% at 20 but loses 1.9 to 3.0% at 21 and 22: Shredding Attacks' first point does nothing alone (it ties a Feral Charge point to the last digit at 20), but the second point at 21 makes Shred the builder, so it is the toll for a cheap Shred and stays at 20 to 22. Two Forever leveling guides put Sharpened Claws ahead of Savage Fury (Mobalytics, doublezug; Mobalytics, worthy-shadow). | `out/feral-final.tsv`, `out/feral-optB.tsv`, `out/feral-greedy-horde.txt` | CONFIRMED-BUG | old | S 6cdbfb90d |
| F3 | **Predatory Discipline at 31 before Leader of the Pack, Blood Frenzy, Natural Instinct and Rend and Tear.** Predatory Discipline only changes Ferocious Bite, which the leveling cat seldom presses (the bare fight never does: APL expected_idle 22829); its point ties a Feral Charge point exactly at 34 to 39. Now Leader of the Pack 31, Blood Frenzy 32, Natural Instinct 33 and 34, Blood Frenzy 35, Rend and Tear 36 to 38, Predatory Discipline 39 (it still opens Shifting Power at 40). +3.0 to 3.4% at 31 to 33, +0.7 to 0.9% at 34 and 35, +0.5% at 36 to 38, both seeds and factions. Moves checkpoints 35 and 38 (ladder goldens), not band 40. | `out/feral-optC.tsv`, `out/feral-greedy-C.txt` | CONFIRMED-BUG | old | S 6cdbfb90d |
| F4 | Feral Charge at 25 does nothing in the sim's fight (a stationary boss you already stand behind), and deferring it to 39 measured +1.6 to 4.7% at 25 to 33. **Kept at 25 by real play:** in Forever the talent is also a Cat Form leap behind the target (Wowhead Forever 1238122: "Feral Charge (Cat) ... Leap behind an enemy", 30 s), the only way a solo cat Shreds a mob that faces it (feral.md "Levels 20 to 60"), and Forever leveling guides take it early. So the point is not provably dead for the spec; keeping it also keeps band 30's build unchanged. | https://nether.wowhead.com/forever/tooltip/spell/1238122; `out/feral-final.tsv` (option A) | Aligned (real play; sim cannot value the leap) | - | - |

## Hunter: Survival

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | Seed: Improved Tracking at 10 to 14. In Forever it is +1% damage a point to the tracked type (`talents/hunter.json` 24293), and the engine applies it to every trackable type and the type-neutral target (`sim/hunter/talents.go:648-690`). It is live: Deflection's parry in its place loses 0.5% at 10 and 2.7 to 2.8% at 14, both seeds and factions. | `out/hunter-tracking.tsv` | Aligned | - | - |
| H2 | Clever Traps (20, 21) before Surefooted (22 to 24): both live; Surefooted first loses 1.1% at 20 for a Troll, gains 0.15% for a Dwarf. No adoption. | `out/hunter-traps.tsv` | Aligned | - | - |
| H3 | **Deterrence at 25 before Predator's Edge.** Deterrence is a defensive the rotation never presses, and the 15-point row is open at 25. Predator's Edge (crit damage and off-hand damage) now takes 25 to 29, Strider Kick 30 (reserved, unchanged level). +1.3 to 1.5% at 25 to 29, +1.5% at 30, both seeds and factions. **Moves the band 30 build** (`500230230050000100` for `500230231040000100`). | `out/hunter-deterrence.tsv` | CONFIRMED-BUG | old | S 6cdbfb90d |
| H4 | **Counterattack at 32 (and Deterrence at 31) before Resourcefulness, Expose Prey, Survivalist's Discipline and Lightning Reflexes.** Counterattack needs a parry; a target hitting the tank or your pet never gives one (APL expected_idle 20910). Now Resourcefulness 31 and 32, Expose Prey 33 and 34, Survivalist's Discipline 35 and 36, Lightning Reflexes 37, Deterrence 38, Counterattack 39; both still pay the toll for Lacerating Strikes at 40, which is unchanged. +2.0 to 2.5% at 31, +4.9 to 5.7% at 32 to 35, +4.0% at 37, +0.7% at 38, both seeds and factions. Forever leveling guides (Mobalytics, Expcarry) take Deterrence and Counterattack in the mid-twenties out of Classic habit; the guide's own leveling advice keeps the pet on the mob, where neither adds damage. Decided on the measured gain. | `out/hunter-counterattack.tsv` | CONFIRMED-BUG | old | S 6cdbfb90d |
| H5 | **Raptor Strike rank rows follow the new order.** The rank 1 row sat at 33 "while only the first point of Resourcefulness is in"; that level is now 31. At 31 rank 1 beats rank 4 by 0.7 to 1.1% (rank 2 by 0.4 to 0.6%), all eight races, both seeds; at 33 rank 1 now loses 2.2 to 2.4% to rank 5. Row moved 33 to 31. At 26 to 29 the rows hold on the new order, except an Orc at 27, who now gains 0.2% with rank 2 over rank 3 at both seeds (a Windshaper at 28 ties, 0.03%, unchanged). | `out/sv-raptor-30-35.txt`, `sv-raptor-31-33-races.txt`, `sv-raptor-26-29.txt`; test `TestRound24SurvivalRaptorStrikeRankRows` | CONFIRMED (follows H4) | this lane | S 6cdbfb90d |
| H6 | Deterrence and Counterattack also sit in the level 60 (raid) build, where the boss hits the tank. Not touched: Deterrence is the page's emergency button when a raider pulls aggro, so it is not provably dead, and dropping it is a raid-build change outside this round's scope. | survival.md "Threat" | OPEN (raid build) | - | - |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Seed: Divine Intellect at 10. It is 2% Intellect, so not provably nothing (a little mana). Redoubt in its place ties on threat at 10 to 14 (-0.05 to +0.15%, not beyond error at both seeds), with chance of death about a point lower at 14. Threat decides tanks: no change. | `out/prot-orders.tsv` | Aligned (tie) | - | - |
| P2 | **Sacred Duty's second point at 21 before Shield Specialization.** Sacred Duty only shortens Bulwark and the bubbles; Shield Specialization's blocks return 4% mana. Now Shield Specialization 21 and 22, Sacred Duty 23. Threat +2.3 to 3.2% at 21 and 22, chance of death no higher (80.3 to 79.2% at 21, Horde seed 1). Level 20 keeps Bulwark (round 14 ruling); between checkpoints, so no band moves. | `out/prot-orders.tsv` | CONFIRMED-BUG | old | S 6cdbfb90d |
| P3 | **Toughness 2 to 5 at 25 to 28 before Swift Judgement.** The 15-point row opens at 25; Toughness only adds armor. Swift Judgement (a second Judgement and taunt) now at 25, Toughness 26 to 29. Threat +5.2 to 6.2% at 25 to 28, both seeds and factions. Band 30 unchanged; moves checkpoint 25 (ladder golden). | `out/prot-orders.tsv` | CONFIRMED-BUG | old | S 6cdbfb90d |
| P4 | The tank fight's chance of death is 99 to 100% at 25 to 28 on either order: the band 20 boss outclasses a bare level 25 tank. Not an order question. | `out/prot-orders.tsv` | OPEN (encounter tuning, for the tank lane) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | `data/tests/test_card_note_levels.py`: the Feral Rip note no longer needs its allowlist line (the comma that split "once you have Rend and Tear" is gone), deleted. The Survival Hunter's Mark line (`ef096649bb13`) is kept although `test_every_allowlist_line_still_matches_a_note` fails on it **until addon-data.json is regenerated**: the note's first sentence changed, so it matches no committed card line now, and after the regen the unchanged second sentence hits again. | test run | Expected (regen) | - | S 6cdbfb90d |
| X2 | Logs, last looked 01:23 UTC 7 Nov: `GET https://api.foreversixty.gg/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct). Nothing here needs a log. | curl | OPEN (no data) | - | - |

## Before/after

Bare band 20 (levels 20 to 29) or 30 (30 to 39) gear at the level, 10,000 iterations. Absolute is Horde seed 1 (Protection: threat per second); change is Horde 1 / 50001, Alliance 1 / 50001.

| spec | level | before | after | change |
|---|---|---|---|---|
| druid-feral | 23 | 57.33 | 58.15 | +1.43 / +1.40, +1.37 / +1.35% |
| druid-feral | 24 | 58.39 | 60.03 | +2.81 / +2.78, +2.68 / +2.66% |
| druid-feral | 26 | 60.57 | 61.39 | +1.36 / +1.38, +1.30 / +1.30% |
| druid-feral | 31 | 96.26 | 99.18 | +3.03 / +3.04, +2.97 / +2.98% |
| druid-feral | 33 | 99.07 | 102.38 | +3.34 / +3.37, +3.22 / +3.15% |
| druid-feral | 35 | 103.40 | 104.20 | +0.78 / +0.78, +0.87 / +0.87% |
| druid-feral | 38 | 114.32 | 114.94 | +0.54 / +0.54, +0.55 / +0.55% |
| druid-feral | 20, 25, 27 to 30, 39, 40 | | | unchanged (same build) |
| hunter-survival | 25 | 63.22 | 64.14 | +1.47 / +1.46, +1.45 / +1.45% |
| hunter-survival | 29 | 74.75 | 75.73 | +1.32 / +1.32, +1.29 / +1.29% |
| hunter-survival | 30 | 100.40 | 101.95 | +1.54 / +1.54, +1.54 / +1.53% |
| hunter-survival | 32 | 106.35 | 111.98 | +5.29 / +5.27, +5.28 / +5.28% |
| hunter-survival | 35 | 120.61 | 127.46 | +5.67 / +5.72, +5.66 / +5.69% |
| hunter-survival | 37 | 135.40 | 140.96 | +4.11 / +4.11, +4.02 / +3.97% |
| hunter-survival | 38 | 141.49 | 142.48 | +0.70 / +0.70, +0.68 / +0.68% |
| hunter-survival | 31 (Raptor Strike rank 1 row) | 105.58 | 106.67 | +1.03% Troll seed 1; +0.65 to +1.12% every race, both seeds |
| paladin-protection | 21 | 39.77 | 40.79 | +2.56 / +2.49, +3.06 / +3.15% |
| paladin-protection | 22 | 41.89 | 42.85 | +2.29 / +2.27, +2.29 / +2.43% |
| paladin-protection | 25 | 43.91 | 46.20 | +5.22 / +5.29, +5.78 / +5.75% |
| paladin-protection | 28 | 48.50 | 51.32 | +5.81 / +5.54, +6.21 / +6.00% |

Raid 60 and every level from 40 up: unchanged for all three specs (the orders rejoin by 40). No spec outside druid-feral, hunter-survival and paladin-protection moves.

## Files to regenerate

- `sim/request/testdata/ladder/` goldens (`TestRotationLadder`): druid-feral (rungs 35, 38), hunter-survival (25, 30, 35, 38), paladin-protection (25). Until then `web/src/lib/guides/leveling-builds.test.ts` may fail on those rungs.
- `addon-builds.json` (the three leveling orders), `addon-data.json` (card notes: Feral Rip; Survival Hunter's Mark, Raptor Strike, Mongoose Bite; the Raptor Strike rank rows), then `Data.lua`. `test_card_note_levels.py::test_every_allowlist_line_still_matches_a_note` fails until addon-data is regenerated (X1).
- `dps-per-point` / stat weights for hunter-survival band 30.
- Coach starters (`persona_starter.py --all`) after addon-builds.
- Engine fork APL copies: `ui/feral_druid/apls/forever_feral.apl.json` (Rip note) and `ui/hunter/apls/forever_survival.apl.json` (three notes) from `data/curated/apl/`.
- Hand copies done, not regens: `sim/leveling/ladder-first.json` = `data/curated/ladder-first.json`; `sim/rankchoices/rank-choices.json` = `data/curated/rank-choices.json`; `sim/request/apl/{druid-feral,hunter-survival}.apl.json` notes edited by hand to match.

## Re-rank needed

- **hunter-survival band 30.** H3 changes its band 30 build (Predator's Edge 5, no Deterrence).
- druid-feral and paladin-protection: none (band 20, 30 and 40 builds unchanged).

## Measure at refresh

- Survival band 30 on its re-ranked gear.

## For other lanes

- **Coach (tools/coach/starter_overrides.py, not touched here):** line 126 "From thirty-five, Expose Prey" is now thirty-three; line 673 "Counterattack from thirty-two" is now thirty-nine; line 688 "Deterrence from twenty-five" is now thirty-eight. Guide-accuracy HUNTER-R26-7's proposed "from" levels for Counterattack and Deterrence need the same update.
- **Tank lane:** P4, the band 20 tank boss kills a bare level 25 to 28 Protection paladin almost every time.

## For the guide lane

Done here (`_prose.test.ts` and `_sections.test.ts` pass): druid/feral.md "Levels 20 to 60" (the talent order), hunter/survival.md "Levels 20 to 60" (the talent order paragraph, the Orc at 27, the rank 1 level 31), paladin/protection.md "Levels 20 to 60" (Shield Specialization and Swift Judgement). Card notes in both APL copies: Feral Rip ("from 36 in either order"), Survival Hunter's Mark (33), Raptor Strike (31, a Horde hunter at 26 and 27), Mongoose Bite (Expose Prey at 33 and 34).

## Guide text not done

- None.

## Still open

- H6 (Deterrence and Counterattack in the raid build), P4, X2. Owner pending items untouched.

## Sources

- Client 1.60.1.70790: `talents/{druid,hunter,paladin}.json`, `trainables/druid.json`, `spellranks.json`.
- Wowhead Forever, Feral Charge: https://nether.wowhead.com/forever/tooltip/spell/1238122 (read 7 Nov).
- [Mobalytics, Feral druid leveling guide 1-60 (doublezug)](https://mobalytics.gg/wow-forever/profile/doublezug/classes/feral-druid-leveling-guide-level-1-60): Feral Charge early, Sharpened Claws before Savage Fury.
- [Mobalytics, worthy-shadow druid build](https://mobalytics.gg/wow-forever/profile/worthy-shadow-yj9332/classes/my-class): Sharpened Claws 20 to 21, Savage Fury 22 to 23, Feral Charge 24, the cat leap.
- [Mobalytics, Survival hunter guide](https://mobalytics.gg/wow-forever/classes/survival-hunter-guide) and [Expcarry, WoW Forever Survival hunter guide](https://expcarry.com/wow-forever-survival-hunter-guide): Deterrence and Counterattack in the mid-twenties (H4 disagrees, on the measured gain). Read 7 Nov.
