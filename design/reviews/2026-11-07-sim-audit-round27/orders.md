# Sim-accuracy round 27: orders lane

Site `wt/r27-orders` from int/r27 af9721524. Fork `lane/r27-orders` from int/r27 227d95bd1, **unchanged** (no engine bug found). Client 1.60.1.70790. Ship mode: touched tests only, no regens. Every sim went through `$R/sim.sh`, one at a time.

**Result: 6 confirmed bugs, all fixed, plus one stale test.** Four leveling orders (origin: old) spent their first points on talents that do nothing in a fight while a damage talent of the same row waited (Subtlety, Arms, and two found in the fresh re-audit: Assassination and Combat). Seed 3: the Skyborne druid rank rows below 20 were inherited from night elf (High Order) and tauren (Windshaper) and are wrong in 10 of 40 level-race cells; both specs now have their own Skyborne rows. Combat's bare-60 Slice and Dice stays: no real in-game condition separates bare from raid.

Method for the orders: `spec-breakdown -preset bare -band 20 -level N -talents <walk string>` on frozen BiS (`$R/bis-frozen`), old walk against new walk at every level the change touches, 10,000 iterations, seeds 1 and 50001, both factions' published races. A walk row that moves a checkpoint (10, 20, ...) needs a measured gain (ladder-first.json); every row here has one. Outputs: `$R/orders/out/` (sub-final.txt, arms-final.txt, assa-final.txt, combat-final.txt, combat-bare60.txt, combat-buffs2.txt) and `$R/orders/t/` (the Skyborne grid).

## Rogue: Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Leveling order: Master of Deception at 10 to 12 before Opportunity.** Master of Deception only lowers the chance mobs see you in Stealth; it does nothing in a fight. Opportunity (+4% a point to Backstab, Garrote and Ambush, `talents/rogue.json`) waited until 13 to 17. Opportunity now comes at 10 to 14 and Master of Deception at 15 to 17; 18 on is unchanged. On the band 20 table's mace (Diamond Hammer; the sim casts Sinister Strike below 20) level 10 ties exactly (15.020 both, neither point does anything there), with a dagger (the guide's advice) level 10 gains +1.36% (14.29 → 14.48), level 14 +0.45 / +0.53% (the opening Garrote), level 15 +0.30 / +0.35%, both seeds, both factions. The talent adoption rule: a point that does nothing is a bug by itself when its replacement is legal, does something and costs nothing; this one costs nothing at any level. Red first: `TestSubtletyTakesOpportunityBeforeMasterOfDeception`. | `out/sub-final.txt`; `data/curated/ladder-first.json` | CONFIRMED-BUG | old (round 26 melee flagged it, not taken) | S 02c23cd10 |
| S2 | The band 20 table gives Subtlety a mace below 20 while the guide says "keep a dagger in the main hand while you level" (subtlety.md:123). The sim's dagger at 10 is 5% behind the mace on DPS. Not a sim bug; it is the BiS pick against the guide's advice (see For the guide lane). | `bis-frozen/rogue-subtlety.json` band 20 | Aligned (guide lane) | - | - |

## Warrior: Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Leveling order: Deflection at 11 to 13 before Improved Rend.** Deflection's parry does nothing while the target hits someone else (and the bare fight's target never swings at you); Improved Rend (+12% a point to Rend, cast from 4) waited until 14, 17 and 18. A sweep at 10 to 14 put Improved Rend first at every level, including level 10 over Improved Heroic Strike (+1.9%, 3,000 iterations), and Improved Heroic Strike over Deflection at 13 (+0.7%). The order is now Improved Rend 10 to 12, Improved Heroic Strike 13, the toll point of Deflection 14, Improved Overpower 15 and 16, Deflection 17 and 18, Improved Charge 19 and 20. Level 20 is unchanged. 10,000 iterations, seeds 1 / 50001: level 10 +1.8% Horde, +0.9 to 1.1% Alliance; 11 +2.6 to 3.1%; 12 +4.8%; 13 +4.2 to 4.3%; 14 +2.6%; 15 +2.3 to 2.6%; 16 +2.2%; 17 +1.0 to 1.1%. Forever leveling practice agrees: the Mobalytics Forever Arms guide (Ahlaundoh, updated 5 Nov 2026, https://mobalytics.gg/wow-forever/classes/arms-warrior-guide) opens Improved Rend at 10, 11 and 12; Lexd's Forever build takes Rend first, then Deflection. Red first: `TestArmsTakesImprovedRendBeforeDeflection`; round 26's `TestArmsTakesImprovedOverpowerBeforeTacticalMastery` level 10 case moved. Moves checkpoint 10 (the ladder golden), not the band 20 build. | `out/arms-explore.txt`, `out/arms-final.txt` | CONFIRMED-BUG | old (walk order since the 70755 trees) | S 02c23cd10 |

## Rogue: Assassination (fresh re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Leveling order: Improved Gouge and Remorseless Attacks at 10 to 14 before Malice.** The rotation never casts Gouge, and Remorseless Attacks needs a kill, so on one target neither does anything, while Malice (+1% crit a point on attacks and poisons) waited until 15 to 19, and Ruthlessness, Improved Slice and Dice and Lethality after it. Now: Malice 10 to 14, Ruthlessness 15 to 17, Improved Slice and Dice 18 and 19, Lethality 20 and 21, Improved Gouge 22 to 24, Remorseless Attacks 25 and 26; the order rejoins the old one at 28. 10,000 iterations, seeds 1 / 50001, both factions: 10 +1.3%, 14 +6.6 to 6.7%, 17 +3.2 to 3.3%, 20 +3.0%, 22 +2.7 to 2.8%, 25 +0.3%. Two Forever leveling orders open Malice 10 to 14 then Ruthlessness 15 to 17 (https://wowfdb.com/classes/rogue/assassination/leveling/, https://wowforevercompass.com/class-guides/rogue/assassination/, read 7 Nov). Red first: `TestAssassinationTakesMaliceBeforeGouge`. **Moves the band 20 build** (`...-005302001...` in place of `325100...`). | `out/assa-final.txt`, `out/assa-old-walk.txt` | CONFIRMED-BUG | old | S 0960b9067 |

## Rogue: Combat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Leveling order: Lightning Reflexes at 14 to 18 before Precision and Puncturing Wounds.** Lightning Reflexes is dodge, which does nothing while the target hits someone else; one point of it is the 5-point row toll. Puncturing Wounds (Backstab crit) waited until 19 to 21 and Precision (+1% hit a point) until 22 to 24. Now: Coup de Grace 10 to 12, Efficient Strikes 13, Lightning Reflexes 14, Precision 15 to 17, Puncturing Wounds 18 to 20, Lightning Reflexes 21 to 24, where the order rejoins the old one; Adrenaline Rush is reserved so it stays at 40 (without it the reservation path moved the 40 build). Band 20 dagger, 10,000 iterations, seeds 1 / 50001, both factions: 15 level, 16 +1.0%, 17 +2.1 to 2.2%, 18 +5.4 to 5.5%, 20 +5.4 to 5.7%, 22 +2.1 to 2.3%. With a sword (Thief's Blade 5192) Precision first gains +2.1 to 2.2% at 17 and +2.9% at 20 while Puncturing Wounds gains nothing, so Precision leads (with a dagger Puncturing Wounds first would be 7% better at 17, but the guide levels Combat with a sword). Red first: `TestCombatTakesPrecisionBeforeLightningReflexes`; `rogue-combat` added to `TestCuratedLadderFirstIsSound`. **Moves the band 20 build** (`3113030...` in place of `3152...`). | `out/combat-final.txt`, `out/combat-old-walk.txt` | CONFIRMED-BUG | old | S aaaf1785b |
| C2 | Precision's first point does nothing at 15: 3,000-iteration runs with Precision 1 and with a second Lightning Reflexes point are identical to the last digit (24.578), and Precision 2 adds. It is the level-gap rule: against a target three levels up the first 1% of hit is lost (the Classic skill-gap hit suppression the engine models). Aligned, recorded so nobody reads it as an engine bug. | `out/` 3,000-iteration runs | Aligned | - | - |
| C3 | **Seed 4: Slice and Dice at bare 60.** The line runs only at 60 (round 26 R1); bare 60 loses with it, 10,000 iterations, seeds 1 / 50001: Horde 287.79 / 287.65 against 289.07 / 289.19 without (-0.44 / -0.53%), Alliance 290.09 / 289.86 against 291.74 / 291.57 (-0.57 / -0.59%). What flips it (3,000 iterations): one attack power buff ties (Battle Shout 303.96 / 303.88, Blessing of Might 303.54 / 303.49, Windfury Totem 303.70 / 303.95); two attack power buffs already favour Slice and Dice (319.60 / 318.53, +0.3%); armor debuffs alone favour skipping it. So the break-even sits at roughly one attack power buff, on a continuum of attack power and armor. No single in-game state (a named buff, the target, the group) separates bare from raid; an attack-power threshold would be a number fitted to this gear. The 0.5% stays, by the seed's rule. No change. | `out/combat-bare60.txt`, `out/combat-buffs2.txt` | Aligned (recorded) | - | - |

## Druid: Skyborne rank rows below 20 (seed 3)

Joint grid on the Skyborne races themselves: `spec-breakdown -preset bare -band 20 -level N -race <skyborne> -rank ...`, every rank combination of the spells sharing the mana bar (Balance: Wrath 9912, Moonfire 9835; Restoration: tank Rejuvenation 9841, party Rejuvenation 8910, Healing Touch 9889, Regrowth 9858), levels 10 to 19, 10,000 iterations, seed 1 (564 runs), then every flagged cell's current and best combination at seed 50001 (20 runs). Adopted where the best beats the current rows by more than 2 standard errors at both seeds. Round 26's night elf and tauren values for the same combinations (`r26/ranks/t/grid1.tsv`, `conf.tsv`) show the cells are race-specific, so every new band names its races; overlapping bands for one race are refused by the parser, so the shared bands were split rather than layered.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **Balance, both Skyborne races at 18: Wrath 3 with Moonfire 1, not 2 and 2.** High Order +1.30 / +1.46%, Windshaper +1.75 / +1.96% (seeds 1 / 50001). Tauren is the other way (Wrath 3 with Moonfire 1 loses 1.8%, round 26 grid), night elf is a tie (+0.07 / +0.23%, left as it was). **Windshaper at 14: rank 1 of both**, +1.17 / +1.13%; High Order at 14 ties (+0.06%), unchanged. The Wrath 17 to 18 band splits at 18 (night elf and tauren keep rank 2), the Moonfire 18 to 19 band splits into 18 (night elf and tauren rank 2, Skyborne rank 1) and 19 (rank 2 for all); new Windshaper bands at 14. Red first: `TestRound27SkyborneBalanceRankRows`. | `$R/orders/t/sky1.tsv`, `conf1.tsv`, `an.py` | CONFIRMED-BUG | new (round 26 rows) | S 35ad9e758 |
| D2 | **Restoration, Skyborne rows.** At 10 both Skyborne races heal most with rank 2 of everything (High Order +0.23 / +0.23%, Windshaper +0.61 / +0.61%): the tank's rank 1 band at 10 now names night elf and tauren, and the night elf bands for the party Rejuvenation and Healing Touch at 10 no longer name High Order. At 12 a High Order puts rank 1 on the tank (+0.47 / +0.40%). A Windshaper heals like a night elf, not a tauren: party Rejuvenation and Healing Touch rank 1 at 13 and 14 (+1.28% at 13, +1.33% at 14, both seeds), Healing Touch 1 at 15 (+0.50 / +0.45%), and both rank 1 at 16 in place of the tauren's party rank 3 (+2.52 / +2.52%). Red first: `TestRound27SkyborneRestorationRankRows`; the round 24 second-pass case for Windshaper at 15 moves to rank 1. | `$R/orders/t/sky1.tsv`, `conf1.tsv` | CONFIRMED-BUG | new (round 26 rows) | S 35ad9e758 |
| D3 | `TestRestorationRankRowsBelowTwenty` (`sim/request/druid_audit_rotation_test.go`) still asserted round 24's ranks (tauren tank rank 2 at 13, rank 3 at 16; night elf party rank 2 at 10) and was **red on int/r27 before this lane** (checked at HEAD with the change stashed). Round 26 changed the rows and updated the Balance test but not this one. Now a per-race table of the current tank and party ranks. | `go test -run TestRestorationRankRowsBelowTwenty` at HEAD | CONFIRMED (stale test) | old (round 26) | S 35ad9e758 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Other specs whose walks open on points that do nothing in the sim's fight, ahead of a damage or healing talent of the same row (not mine; for their lanes): warlock-destruction Destructive Reach 10 to 11 before Bane, warlock-demonology Improved Health Funnel and Demonic Embrace before Unholy Power, mage-frost Frost Warding 15 to 16 before Precision, druid-feral Feral Swiftness before Feral Instinct (both inert on one target; order only), hunter-survival Improved Tracking 10 to 14 (check what it does in Forever first), paladin-protection Divine Intellect at 10. Each needs the same old-against-new sweep. | ladder walk 10 to 20 for every spec (scratch print) | OPEN (other lanes) | - | - |
| X2 | Logs, last looked 7 Nov 00:00 UTC: `GET https://api.foreversixty.gg/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct). Nothing here needs a log. | curl | OPEN (no data) | - | - |

## Before/after

Bare band 20 at the level, 10,000 iterations, seed 1 absolute (Horde), change at seeds 1 / 50001:

| spec | level | before | after | change |
|---|---|---|---|---|
| rogue-subtlety | 10 (mace) | 15.020 | 15.020 | 0 / 0 |
| rogue-subtlety | 14 | 25.464 | 25.579 | +0.45 / +0.45% (Alliance +0.53%) |
| warrior-arms | 10 | 12.481 | 12.707 | +1.81 / +1.75% (Alliance +1.1 / +0.9%) |
| warrior-arms | 12 | 15.879 | 16.641 | +4.80 / +4.85% |
| rogue-assassination | 14 | 24.000 | 25.616 | +6.73 / +6.73% |
| rogue-assassination | 20 | 43.106 | 44.379 | +2.95 / +3.00% |
| rogue-combat | 18 | 34.576 | 36.452 | +5.43 / +5.41% |
| rogue-combat | 20 | 43.213 | 45.660 | +5.66 / +5.61% |
| druid-balance (High Order) | 18 | 29.25 | 29.63 | +1.30 / +1.46% |
| druid-balance (Windshaper) | 18 | 28.82 | 29.32 | +1.75 / +1.96% |
| druid-restoration (Windshaper, HPS) | 16 | 25.04 | 25.67 | +2.52 / +2.52% |

Raid 60 and every band from 25 up: unchanged for all four melee specs (the walks rejoin by 28; level 40 checked for Combat). No spec outside rogue, warrior and druid moves.

## Files to regenerate

- `sim/request/testdata/ladder/` goldens (`TestRotationLadder`): rogue-subtlety (rung 10), warrior-arms (rung 10), rogue-assassination (10, 20, 25), rogue-combat (20). Until then `web/src/lib/guides/leveling-builds.test.ts` fails on those rungs.
- `addon-builds.json` (the four leveling orders), then `Data.lua`.
- `dps-per-point` / stat weights for rogue-assassination and rogue-combat band 20.
- Coach starters (`persona_starter.py --all`) once `addon-builds.json` is regenerated, if they read talent levels.
- Engine fork copy of the Restoration APL (`ui/druid/apls/forever_restoration.apl.json`, Makefile:363): card notes changed in `data/curated/apl/druid-restoration.json` and `sim/request/apl/druid-restoration.apl.json`, sync the fork copy. No druid ladder rung moves (`TestRotationLadder` fails only on the four melee specs).
- Hand copy: `sim/rankchoices/rank-choices.json` = `data/curated/rank-choices.json`.
- Hand copies done, not regens: `sim/leveling/ladder-first.json` = `data/curated/ladder-first.json`.

## Re-rank needed

- **rogue-combat band 20 (LOUD: Combat is being re-ranked now on lane/bis-r26-rerank).** C1 changes its band 20 build (Precision 3 and Puncturing Wounds 3 in place of Lightning Reflexes 4 and Puncturing Wounds 2), so its band 20 picks and published number move again (hit is now in the build).
- **rogue-assassination band 20.** R1 changes its band 20 build (Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Lethality 1).
- rogue-subtlety and warrior-arms: no band build moves (only rung 10 of the ladder).

## Measure at refresh

- Assassination and Combat band 20 on their re-ranked gear.

## For other lanes

- X1 (the walks above).
- **BiS / guide:** Subtlety's band 20 table picks a mace (Diamond Hammer), the guide says level with a dagger (S2). Combat's band 20 table picks a dagger, the guide says level with a sword (combat.md, Levels 20 to 60) and uses a dagger in dungeons, which matches the table. Worth one line in Subtlety's guide or a dagger-only band 20 pick.
- Coach: `tools/coach/starter_overrides.py`:489 still names "Initiative full only from thirty-two" (round 26); the Subtlety walk's first points now are Opportunity.

## For the guide lane

Done here, `_prose.test.ts` and `_sections.test.ts` pass: rogue/subtlety.md:129, warrior/arms.md:125, rogue/assassination.md (Levels 1 to 19, the talent paragraph), rogue/combat.md:126. druid/balance.md:127 and druid/restoration.md:113 (the low-level rank lines name races, not factions), and three Restoration card notes (tank Rejuvenation, party Rejuvenation, Healing Touch). No card note changes.

## Guide text not done

- None.

## Still open

- X1, X2. C3 stays as recorded (bare 60 -0.5% with Slice and Dice). Night elf at Balance 18 ties Wrath 3 / Moonfire 1 (+0.07 / +0.23%, not beyond error at seed 1), left on 2 and 2. Owner pending items untouched.

## Sources

- Client 1.60.1.70790: `talents/{rogue,warrior,druid}.json`, `items.json`.
- Mobalytics, WoW Forever Arms Warrior Guide, Ahlaundoh, updated 5 Nov 2026: https://mobalytics.gg/wow-forever/classes/arms-warrior-guide (leveling path: Improved Rend 10 to 12).
- Mobalytics, Lexd's Arms warrior PvP build leveling to 60 (search summary only, page 403): https://mobalytics.gg/wow-forever/classes/lexd-s-arms-warrior-pvp-build-leveling-to-60
- WoW Forever DB, Assassination Rogue leveling guide: https://wowfdb.com/classes/rogue/assassination/leveling/ (Malice 5/5 10 to 14, Ruthlessness 15 to 17); Forever Atlas: https://wowforevercompass.com/class-guides/rogue/assassination/ (Malice at 10). Read 7 Nov.
