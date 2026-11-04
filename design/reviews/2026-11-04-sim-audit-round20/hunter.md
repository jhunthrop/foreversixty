# Round 20, lane hunter: Beast Mastery, Marksmanship, Survival (melee)

Client 1.60.1.70755 with the late re-layer. Site wt/r20-hunter from int/r20 (eb0f16900, with the shaman lane); fork lane/r20-hunter from int/r20 da54b7b06.
Method: spec-breakdown through `$R/sim.sh`, guide builds (`$R/rt`), `-bis-dir $R/bis-frozen`. 3000 iterations to explore; 10000 at seeds 1 and 50001, both factions, for every reported number. Log tools: `$R/hunter/tools/elune.py` (Elune's Light and Heart of the Lion on pets in the beta corpus) and `$R/hunter/tools/stingfull.py` (Serpent Sting and Arcane Shot against the hunter's ranged attack power read off Multi-Shot against Auto Shot); both copied to `hunter-tools/` beside this file.

**Result: 4 CONFIRMED-BUGs, all fixed, all `old`.** Blizzard's 3 November notes say Serpent Sting and Arcane Shot carry ranged attack power (5% a tick, 15% a hit) and the 16 October notes say fire traps scale with attack power; the client rows carry 0, so no round saw it, and the beta logs had said it since October (round 19 H3 held the data and could not decompose it against the post-push bases). The engine now adds the shares (E 628f99a7c): Beast Mastery +14.5% and Marksmanship +13.9% at raid 60, Survival +1.9%. With Arcane Shot and Serpent Sting that much better for their mana, Multi-Shot is the worst shot per mana, so the rotations hold it under a quarter mana and free Arcane Shot (S, with guide text and a wow-player SHIP): another +0.4% and +2.4% at raid 60, +1.9% and +4.4% on a 420 s fight. The seeds: the hawk's level-60 swing and Heart of the Lion on the pet stay OPEN (no live log before 23:00 UTC); the pet's crit copy is dynamic in the client data; Lone Wolf against the pet is re-measured and unchanged.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Summon Hawk's level-60 swing: no client row carries it.** 1293241/1293525/1293526/1293527 effect 1 is Effect 32 (trigger missile) of 1312639, whose own rows are a dummy with no creature, amount or period (SpellEffect 1313790, 1314234, 1314237, 1314240; the cached 70755 csv). CreatureDisplayInfo and CreatureModelData carry models, not damage; the hawk's swing lives in a server creature template. The forever-hunter wiki's Attack Formulas page (read 2026-11-04) lists both hawk components as "Unverified". So the only way to size it is a level-60 log. foreverlogs has nothing after report 2814 (last checked 21:53 UTC, launch at 23:00), and `api.foreversixty.gg/v1/reports/recent` still lists the 2026-10-20 report. The floor model (0.32 a level) stands; round 18's bar (adopt Summon Hawk if swing / dive >= 24 / (108 + 0.05 x RAP)) stands. | `$SP/r15/hunter/csv/SpellEffect.csv`; https://github-wiki-see.page/m/classic-hunter/forever-hunter/wiki/Attack-Formulas | OPEN (a level-60 Beast Mastery log with Summon Hawk: ten hawk swings and the dive on one target) | - | - |
| H2 | **The pet copies the hunter's crit live, Elune's Light included: the client says so; the bug tracker has one unanswered report that it snapshots.** Hunter Pet Scaling 415429's SpellAuraOptions row (235687) is ProcChance 100, ProcCategoryRecovery 2000, ProcTypeMask 69904: the inheritance is a proc the pet re-applies on its own actions every 2 s, so a 15 s crit buff on the hunter reaches the pet within a swing (forever-bugs #48, neteyesking 2026-09-20, describes exactly this). The engine does the same through `DynamicStatsPets` (`core/pet.go:171-176`, `unit.go:297`): a hunter pet is not a guardian, so Elune's Light's `NewTemporaryStatsAura` reaches it at once. Against that, forever-bugs #48 (Prox-1988, 2026-10-11, build 70425) says the pet snapshots at summon whenever Focused Fire has a point, and updates live without it; the issue is "Pending Approval" with no Blizzard reply and no later comment. The 1742 cached beta timelines hold no Elune's Light caster at all (`elune.py`: 0 casts of 1259799), so the logs cannot settle it. Worth about 0.15% for a Night Elf Beast Mastery hunter (10% pet crit for 15 of 180 s). | wago SpellAuraOptions 70755 (`$SP/r15/hunter/csv`); https://github.com/ClassicWoWCommunity/forever-bugs/issues/48; `$R/hunter/tools/elune.py` | Aligned with the client (OPEN only on the Focused Fire snapshot report: a live log of a Night Elf Beast Mastery hunter with Focused Fire, pet crit inside and outside Elune's Light's 15 s) | - | - |
| H3 | **Heart of the Lion on the pet's buff bar: still no log.** No beta timeline carries 409583 or 409580 on a pet target (`elune.py` over `fl/` and `full/`: 0 rows), as round 17 found; no launch log exists yet. The round-17 basis (409583 unflagged, ImplicitTarget 31, Trueshot Aura on pets in the logs) stands. | `elune.py` | OPEN (a live hunter log with Heart of the Lion up: a `buff_applied` of 409583 on the pet) | - | - |
| H4 | **Ranged Weapon Specialization on Serpent Sting (round 19 H3): the ratios decompose now.** Round 19 compared the October ticks with the post-push bases (14 and 26) and found 1.43-1.57 "unexplained". Against the pre-push bases (12 and 22; the 3 November push raised every rank about 20%, 22 -> 26) and the attack power share, every hunter fits: Serpent Sting tick = base x (1 + Improved Stings) x (Lone Wolf) + 3% RAP, Arcane Shot = base x (Lone Wolf) + 10% RAP, with one ranged attack power per hunter read off Multi-Shot minus Auto Shot (RAP/14 x (2.8 - weapon speed)). Huntscar and Salami (2705, Lone Wolf, about 400-450 RAP): sting 38-40 = 22 x 1.2 + 12-14; Arcane Shot 119-124 = 65 x 1.2 + 41-46. Miss (2712, about 280 RAP, pet out): sting 20-21 = 12 + 8-9; Arcane Shot 67 = 39 + 28. Lokronn (2695, 310 RAP, pet out): sting 23 = 12 + 11; Arcane Shot 74-75 = 39 + 35-36. Draexan and Miss at rank 4 (2705, 2762, about 340-350 RAP): sting 33 = 22 + 11. These hunters are 25-35, below Ranged Weapon Specialization's tier, so the talent's share stays untested. | `$R/hunter/tools/stingfull.py` on `$SP/r16/hunter/full`; round 19's `stingratio.py` on `$SP/r19/hunter/dmg` | OPEN on the talent itself (a level-40+ log with the talent); the attack power share is B2 | - | - |
| H5 | **Lacerating Strikes and the queued Raptor Strike (round 19 H4): no new log.** Every foreverlogs report is beta, levels 13-35. | foreverlogs | OPEN (level-40+ Survival log) | - | - |
| H6 | **Lone Wolf against the pet at 60, re-measured on the landed engine and again after B2.** Before B2, identical to round 19: Survival raid Lone Wolf 860.2/858.8 (Alliance) and 860.5/861.7 (Horde) against the pet build's 839.2/838.3 and 838.5/838.1; bare 60 the pet leads 377.1/376.9 and 375.5/375.2 against 370.0/369.7 and 368.7/368.7. Marksmanship, the same shape: raid Lone Wolf 801.5/802.6 and 801.2/801.7 against the pet's 788.8/789.6 and 787.6/788.1 (-1.6%); bare 60 the pet leads 361.3/361.4 and 360.1/360.2 against 350.5/350.4 and 349.5/349.8 (+3.1%). After B2 the Marksmanship bare gap narrows to +1.8% (pet 382.6/382.7 and 380.0/380.1 against 376.1/375.9 and 373.3/373.6) and stays on the pet's side. The guides already say both: dismiss the pet for a raid, keep it out in the world. | `$R/hunter/lw.tsv`, `$R/hunter/recheck.tsv` | Aligned (no change) | - | - |

## Beast Mastery

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Raid 60 breakdown before this lane (guide raid build, frozen BiS, Horde, 3000 it, 780.0): Auto Shot 43.8%, cat 23.4%, Multi-Shot 10.0, Arcane Shot 7.2, Serpent Sting 6.0, Aimed Shot 5.8; 0 s out of mana. Identical to round 19 B1. After B2 and B3 (893.0): Auto Shot 38%, cat 20%, Arcane Shot 13%, Serpent Sting 11%, Multi-Shot 9%, Aimed Shot 5%. Serpent Sting casts 13555 (rank 8, the top trained rank; 25295 is the pending level-60 book). | `$R/hunter/bd-beast-mastery-h.txt`, `bd2-beast-mastery-h.txt` | Aligned | - | - |
| B2 | **Serpent Sting and Arcane Shot carry ranged attack power; the engine had the client's flat only.** The 70755 SpellEffect rows (679708, 690132, 690655, 690760, 691479, 691570, 691588, 691589, 702403; 1313789 and the Arcane Shot ranks) carry BonusCoefficientFromAP 0, and the shipped tooltips are plain ("An instant shot that causes 280 Arcane damage", "causing 500 Nature damage over 15 sec"), so the engine read 0 (rounds 12, 16, 17, 19 each checked the column and the tooltip). The server adds the share anyway. **Blizzard, WoW Forever Beta Development Notes, updated 3 November 2026** (us.forums.blizzard.com; quoted whole in https://us.forums.blizzard.com/en/wow/t/major-class-final-tuning/2399448, and item by item at https://forever-xp.com/changes/hunter/): "Arcane Shot: Base damage on all ranks increased by ~30% and Attack Power coefficient raised from 0.1 to 0.15." and "Serpent Sting: Attack Power coefficient per tick increased to 5% from 3%. Base damage on all ranks increased ~20%." **The beta logs** (H4): eight hunters in five reports tick 3% and hit 10% of their own ranged attack power over the pre-push bases, the values the notes say were in force then. **The forever-hunter wiki** (Attack Formulas, "Verified" by its testing method) gives the same two pre-push values. New: `attack_power_scaling.go` (5% a tick, 15% a hit, read when the sting is applied or the shot is cast; Blizzard and the wiki say the ticks recalculate from current attack power, which in a sim fight differs only during a trinket or racial use). Tests first: `TestSerpentStingTicksScaleWithRangedAttackPower` (red: tick 100.0, want 142.5 on a naked 60) and `TestArcaneShotScalesWithRangedAttackPower` (red: 280 against 407.5); the three spellconst guards now expect the flat plus the share. Raid 60: Beast Mastery 779.6/779.6 -> 889.5/889.5 (Horde) and 782.8/781.9 -> 892.3/891.3 (Alliance), +14.0 to +14.1%; Marksmanship +13.9%; bare 60 +7.2 to +7.8%, bare 40 and 50 +6.7 to +10.3%. | sources above; `sim/hunter/attack_power_scaling.go`, `serpent_sting.go`, `arcane_shot.go` | CONFIRMED-BUG | old (no round found it; round 19 H3 had the logs; the pending Arcane Shot tooltip item is about spell power, which stays 0) | E 628f99a7c |
| B3 | **The low-mana order was backwards once the shots carry attack power.** At raid 60 after B2 the damage per mana is Serpent Sting 6.4, Arcane Shot 4.5, Aimed Shot 3.8, Multi-Shot 3.1 (Beast Mastery; Marksmanship 8.8, 6.5, 3.5, 3.4), so the shipped list (Arcane Shot gated above 25% mana, Multi-Shot never) dropped the second-best shot first. Variants on a 300 s and a 420 s Horde fight, 3000 it: Multi-Shot gated above 25% with Arcane Shot's gate removed (D) beats the shipped list by +1.4%/+1.9% (Beast Mastery) and +3.7%/+4.4% (Marksmanship); gating Multi-Shot alone (C) and a 15% Arcane Shot gate (E) win less; holding Aimed Shot too (F) loses. D at 180 s, 10000 it, both factions, seeds 1 and 50001: Beast Mastery 889.5/889.5 -> 893.0/893.2 (Horde) and 892.3/891.3 -> 895.0/894.1 (Alliance), +0.3 to +0.4%; Marksmanship 912.6/913.1 -> 934.4/934.4 and 912.4/913.6 -> 934.2/935.1, +2.4%; 420 s, 10000 it: Beast Mastery 871.9/872.0 (Horde) and 875.5/875.1 (Alliance) against the before engine's 760.0/759.9 and 763.1/762.4, Marksmanship 879.2/879.2 and 879.5/880.1 against 757.0/757.0 and 757.3/757.6. Beyond error at every seed. APL: Multi-Shot carries `currentManaPercent > 25%`; Arcane Shot keeps only its cooldown spacing; the three card notes are rewritten; the fork copies are synced by hand (E df3b8dbd6) and `sim/request/apl` too. Guide: beast-mastery.md:91, 99, 113 and marksmanship.md:96, 105, 119 (the rotation paragraph names the attack power on both shots, the low-mana paragraph drops Multi-Shot first, the stat paragraph says attack power now feeds two shots). `_prose.test.ts` passes. wow-player: see "Guide text". | `$R/hunter/lowmana.tsv`, `lowmana2.tsv`, `after2-raid.tsv`, `after2-420.tsv`; `$R/hunter/rot/*.json` | CONFIRMED-BUG (rotation; guides-handoff, round 16 sims requested 1) | old (the order was right while the shots were flat) | S (this lane's site commit) |
| B4 | **Dead points: none.** Every point of the raid and leveling builds re-read against the 70755 text: Endurance Training (pet health and armor), Bestial Swiftness and Intimidation (the road to Bestial Wrath, and a stun with 100% crit on trash), Bestial Discipline (focus and 50% mana regeneration while casting), Focused Fire, Ferocity, Unleashed Fury, Frenzy, Bestial Wrath, Lethal Attacks, Improved Stings, Efficiency, Careful Aim (client 1223984 carries both halves, aura 580 and 598), Trueshot Aura, Mortal Shots, Rapid Killing in the leveling build. Each does something a hunter can see. The shared guard `sim/leveling/dead_points_test.go` now carries the three hunter builds against the cross-tree talents that would be dead (Survival's melee-only talents for the two ranged builds; Marksmanship's ranged-only talents for Survival); it passes. | talents/hunter.json; `sim/leveling/dead_points_test.go` | Aligned | - | S |
| B5 | Gear on the frozen BiS (both factions): the engine registers every item with an effect. Frozen Heart of the Mountain's Chthonic Power (a Frost and Shadow spell damage use, picked for its hit) and Thunderbrew's Boot Flask are in `sim/common`; Rune of the Guard Captain and the rest are plain stats in simdb. Barrage 5/10/15% and Lightning Reflexes 3% a rank match the 3 November notes (the talent file's curve). Hunter's Mark rank 4 is 143 (the notes' rank-4 fix leaves rank 4 above rank 3). Ammo (Thorium Headed Arrow from 52) and quiver (15% from 55) by level. | `sim/common/item_effects.go:3724`, `client_procs.go:26`; `talents.go:46-48`; `huntersmark.go:13`; `sim/request/request.go:702-724` | Aligned | - | - |

## Marksmanship

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Raid 60 breakdown before (801.1, Horde, 3000 it): Auto Shot 53%, Multi-Shot 15.9, Sniper Shot 11.6, Serpent Sting 6.8, Aimed Shot 6.5, Arcane Shot 6.1; 4.3 s out of mana. After B2 (912.3): Auto Shot 46.5%, Multi-Shot 14.0, Serpent Sting 12.2, Arcane Shot 11.2, Sniper Shot 10.2, Aimed Shot 5.7. After B3 (934.4 at 10000 it): the shared cooldown goes to Aimed Shot under a quarter mana and Arcane Shot fills every gap. | `$R/hunter/bd-marksmanship-h.txt`, `bd2-marksmanship-h.txt` | Aligned | - | - |
| M2 | **Aimed Shot rank 1 at 50-51 stays a tie after B2.** Bare band 50, `-level 50` and `51`, Horde, 10000 it: the top rank (4) 272.3/272.4 and 275.7/275.7 against rank 1's 272.6/272.5 and 275.8/275.8. Inside the 0.1 error at one seed each, so no row; the guide's "either serves" at 50 and 51 holds. The 28-49, 52-59 and bare-60 rows gain if anything: the mana rank 1 saves now buys better shots. | `$R/hunter/recheck.tsv`, `lowmana.tsv` | Aligned | - | - |
| M3 | Serpent Sting and Arcane Shot downranks (round 19 H5) cannot flip: the attack power share is the same for every rank, so a lower rank loses the same flat damage and saves the same mana as before, on a larger total. Not re-swept. | arithmetic on `SerpentStingDamage`, `ArcaneShotDamage` | Aligned | - | - |

## Survival (melee)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Immolation Trap carries attack power; the engine had the client's flat only.** Blizzard's notes of 16 October 2026: "The Hunter's Fire traps now correctly scale with Attack Power." The client rows (13797, 14298-14301) carry 0. The coefficient, 2% of ranged attack power a tick (10% over the burn), is the forever-hunter wiki's tested value (Attack Formulas, "Verified"; its patch-notes page writes "2% RAP per tick"). **Single source for the number:** every beta report with trap ticks (2745, 2762: 21-22 and 44 a tick, the flat) is dated 7-10 October, before the fix. Test first: `TestImmolationTrapTicksScaleWithRangedAttackPower` (red: 138, want 168.1 on the test Survival hunter). Raid 60: 860.5/861.7 -> 876.6/877.9 (Horde) and 860.2/858.8 -> 877.0/875.5 (Alliance), +1.9 to +1.9%; bare 60 +3.1%, bare 40 +2.6%, bare 50 +2.3%. The trap's card note now says the burn carries attack power. | sources above; `immolation_trap.go` | CONFIRMED-BUG (the scaling); the 2% is single-source | old (rounds 17 and 19 checked the column) | E 628f99a7c |
| S2 | Raid 60 breakdown (Horde, 3000 it): before 861.0, after 877.2 with Immolation Trap 56.7 -> 72.8 DPS (341 a tick with Clever Traps and crits); the melee kit is unchanged (Raptor Strike 24 hits, Mongoose Bite 20.3, Strider Kick 19.1, Lacerate 9.5 with 53 Lacerating Strikes ticks, Wing Clip 61.8 under Windfury). 0 s out of mana. | `$R/hunter/bd-survival-h.txt`, `bd2-survival-h.txt` | Aligned | - | - |
| S3 | Every Survival point re-read (round 18 S1's list plus Careful Aim: both attack powers, so a melee hunter's point does something). No dead point; the shared guard holds the build off Ranged Weapon Specialization, Mortal Shots, Barrage, Improved Stings and Sniper Shot. | talents/hunter.json; `dead_points_test.go` | Aligned | - | S |
| S4 | Explosive Trap and Volley also scale with attack power on the server per the same notes and wiki (coefficients unverified). No build casts either; left at the flat and noted. | wiki Attack Formulas | Aligned (not cast) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Arcane Shot's pending spell-power item.** The 70755 tooltip has no formula, the client spell-power coefficient is 0, and B2 leaves it at 0. Blizzard's 3 November line names only attack power. Pending stays pending; the owner's in-game check now has a cleaner form: with spell damage gear on, does Arcane Shot's hit rise beyond 280 + 15% of the sheet's ranged attack power. | spelltips 14287 | Pending owner (recorded) | - | - |
| X2 | **Live logs.** Last looked 21:53 UTC on 2026-11-04: foreverlogs `reports/2815-2900` "Report not found", `/v1/reports/recent` the 2026-10-20 report only. The direct `/api/reports/<id>` path refuses plain clients; the browser-headed `fl.sh` wrapper works. Warcraft Logs still blocked (gcloud). | foreverlogs | OPEN (no data) | - | - |
| X3 | Round 19's three engine commits (every Serpent Sting and Raptor Strike rank registered; the 46-47 row) hold: all ranks register, the rows cast what they name, and the Survival 46-47 row still reads rank 5 in `rank-choices.json`. `go vet ./sim/hunter/` reports only the test-only lock copy at `talents_test.go:496` (rounds 18 and 19). | test run; go vet | Aligned | - | - |

## Before/after

Raid 60, guide builds, frozen BiS, 10000 iterations, seed 1 / seed 50001, against `$R/before.tsv`. The engine (E 628f99a7c) and the rotation (B3) together:

| spec | faction | seed 1 | seed 50001 |
|---|---|---|---|
| Beast Mastery | alliance | 782.8 -> 895.0 (+14.3%) | 781.9 -> 894.1 (+14.4%) |
| Beast Mastery | horde | 779.6 -> 893.0 (+14.5%) | 779.6 -> 893.2 (+14.6%) |
| Marksmanship | alliance | 801.5 -> 934.2 (+16.6%) | 802.6 -> 935.1 (+16.5%) |
| Marksmanship | horde | 801.2 -> 934.4 (+16.6%) | 801.7 -> 934.4 (+16.6%) |
| Survival | alliance | 860.2 -> 877.0 (+2.0%) | 858.8 -> 875.5 (+1.9%) |
| Survival | horde | 860.5 -> 876.6 (+1.9%) | 861.7 -> 877.9 (+1.9%) |

Of which the engine alone (B2, S1): Beast Mastery 889.5/889.5 Horde and 892.3/891.3 Alliance; Marksmanship 912.6/913.1 and 912.4/913.6; Survival as above. The rotation (B3) adds the rest.

Bare bands (published entry talents, 10000 it), before -> after (engine; B3 reaches them too but was measured at 60 only):

| spec | band | faction | seed 1 | seed 50001 |
|---|---|---|---|---|
| Beast Mastery | bare 60 | alliance / horde | 378.7 -> 409.2 (+8.1%) / 374.7 -> 404.0 (+7.8%) | 378.7 -> 409.3 / 374.4 -> 403.9 |
| Beast Mastery | bare 40 / 50 | horde | 188.4 -> 201.9 (+7.2%) / 255.4 -> 271.9 (+6.5%) | 188.5 -> 202.0 / 255.5 -> 272.0 |
| Marksmanship | bare 60 | alliance / horde | 371.7 -> 399.2 (+7.4%) / 368.6 -> 395.3 (+7.2%) | 371.8 -> 399.3 / 368.8 -> 395.5 |
| Marksmanship | bare 40 / 50 | horde | 184.4 -> 203.3 (+10.2%) / 252.4 -> 272.3 (+7.9%) | 184.4 -> 203.4 / 252.5 -> 272.4 |
| Survival | bare 60 | alliance / horde | 370.0 -> 381.3 (+3.1%) / 368.7 -> 380.0 (+3.1%) | 369.7 -> 380.9 / 368.7 -> 380.0 |
| Survival | bare 40 / 50 | horde | 166.9 -> 171.2 (+2.6%) / 260.4 -> 266.4 (+2.3%) | 166.8 -> 171.2 / 260.5 -> 266.5 |

Marksmanship now sits above Survival at raid 60 (934 against 877) and Beast Mastery close behind (893); the tier list and the class page's spec order move at the refresh.

No commit moves DPS for a spec outside the hunter class: both engine commits touch `sim/hunter` and `ui/hunter` only; the site changes are hunter rotations, hunter guides and a hunter-only addition to a shared test.

## Files to regenerate

- addon-data / Data.lua (the three hunter card notes), the coach starters for the three hunters, apl-sync copies (synced by hand here; apl-check should pass as is).
- Ladder goldens (`sim/request TestRotationLadder`) for every hunter checkpoint: every band moves.
- dps-per-point (ranged attack power is worth more; haste and crit shift) and the planner vectors for hunter.
- Fork: no golden (TestP1Hunter is skipped awaiting the talent rewrite).
- The BiS re-rank: hunter picks and weights move with the attack power share (Measure at refresh).

## Measure at refresh

- Beast Mastery and Marksmanship picks and weights on the new attack power value (MR20's list, every band); Survival's trap share.
- Round 15's guide list, carried: MM race order, Survival Crusader against Agility, the race re-runs, the best non-BoP two-hander.
- The tier order among the three hunters (Marksmanship first at raid 60 on this engine).

## For other lanes

- **Platform (low-rank proc curve, priest P1):** Blizzard's 3 November notes say "Various talents no longer have a penalty to their chance to trigger when casting spells of a low rank." That is the direct answer to whether a curve exists now.
- **Platform (addon):** the pet frame question stays open; the addon has no pet buff surface (`Recorder.lua` only logs the pet). If one is wanted, a pet buff row showing Heart of the Lion (409583) and the pet's crit would settle H2 and H3 in game.
- **Warlock:** Blizzard's same notes may carry attack power coefficients the client rows hide for other classes; worth a read for every lane whose client coefficient is 0.

## For the guide lane

- beast-mastery.md:91, 99, 113 and marksmanship.md:96, 105, 119 are rewritten here (B3). The card notes for Multi-Shot, Serpent Sting and Arcane Shot (both specs) and Immolation Trap (Survival) are rewritten in `data/curated/apl`; the rendered cards show them after the addon-data regen.
- The class page's "which spec" prose, if any is written after the refresh, should know Marksmanship leads at raid 60 on this engine.

## Guide text

Done in this lane: beast-mastery.md and marksmanship.md (B3), the four card notes. `_prose.test.ts` passes (18). wow-player reviewed the rendered pages (astro dev on this branch, 1440, `$R/hunter/shots/*-1440.png`): first pass SHIP WITH FIXES (two MUSTs: the Beast Mastery weakness bullet still blamed Arcane Shot for the mana, and "Multi-Shot on cooldown is fine" contradicted the low-mana paragraph; three SHOULDs on the Marksmanship Multi-Shot clause, the movement bullet and "hold Multi-Shot first"), all five applied, second pass **SHIP** on the served pages. The reviewer asks for one look at the three rotation cards after the addon-data regen, since the cards read the nightly file.

## Guide text not done

None.

## Still open

- H1 the hawk's level-60 swing; H3 Heart of the Lion on the pet; H2's Focused Fire snapshot report; H4 Ranged Weapon Specialization on the sting; H5 Lacerating Strikes and the queued Raptor Strike. All need live logs (level 40+ or 60).
- S1's 2% is single-source until a live level-16+ hunter log shows Immolation Trap rank 1 above 21 a tick.
- Pending owner, not reopened: Arcane Shot's tooltip with spell power (X1), level-60 book ranks, Heart of the Lion 280/276.
