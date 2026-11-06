# Sim-accuracy round 26: ranks lane

Site branch `wt/r26-ranks` (from int/r26 79096663f, platform's gear fixes in). Fork branch `lane/r26-ranks` (from int/r26 271b42137). Client 1.60.1.70790. BUILD MODE: touched tests only, no regens. Every sim went through `$R/sim.sh`, one at a time.

**Confirmed and fixed: 11.** Origins: R24 6 (rows judged on the gear round 25's P1/P2 and round 26's P6/P7 later fixed), old 5. No engine change; the two fork commits are APL copies.
**Cross-class DPS movers:** none outside the specs named here.

Method: `spec-breakdown -preset bare -band 20 -level N -rank id=r`, frozen BiS (`$R/bis-frozen`), the leveling order's talents at each level, both factions' published races (Balance and Restoration druid night elf / tauren, Holy and Discipline Human / Undead, Protection Dwarf / Undead, Elemental Dwarf / Orc, Restoration shaman Dwarf / Tauren), plus every other race of the class where a row is race-keyed. Every rank combination learned at each level 10 to 19, on a joint grid where spells share the mana bar (Balance: Wrath x Moonfire; Restoration druid: tank Rejuvenation x party Rejuvenation x Healing Touch x Regrowth; Discipline: Renew x Lesser Heal x Power Word: Shield; Holy: Renew x Lesser Heal; Elemental: Lightning Bolt x Flame Shock; Protection: Holy Strike x Seal of Fury). Seed 1 at 10,000 iterations for the full grid (1,084 runs), seed 50001 for the current combination and every challenger beyond error (231 runs). A row loses when a challenger beats it by more than 2 x the combined error at both seeds. Healer HPS prints to 0.01, so the error floor is 0.01. Data: `$R/ranks/t/` (grid1.tsv, conf.tsv, v1.tsv, v3.tsv; an.py, best.py).

## Live logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | No new live Forever logs. `GET https://api.foreversixty.gg/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct, no kills). Nothing here needs a log. | last looked 21:49 UTC 6 Nov | OPEN (no data) | - | - |

## Seed 1: the 20 caster and healer rows at 10 to 19

Per row: old row, verdict, margin of the new choice over the old at seed 1 / seed 50001 (A = Alliance race, H = Horde race).

| spec | spell | old row | verdict | margins (s1 / s50001) |
|---|---|---|---|---|
| Balance | Wrath | 10-12 r1 | holds at 10, 11; **extends to 13**, and **15-16 r1** (with Moonfire r1) | 12 A +1.69/+1.61, H +1.01/+0.89; 13 A +3.13/+3.03, H +2.28/+2.12; 15 A +2.08/+2.09, H +2.26/+2.22; 16 A +4.79/+4.77, H +4.75/+4.71 |
| Balance | Wrath | 14-18 r2 | **14 and 17-18 only** | 14 holds (H r1 +0.24/+0.13, within error; A r1 -0.4); 17 holds; 18 r2 with Moonfire r2: A +4.74/+4.69, H +3.54/+3.52 |
| Balance | Moonfire | 16 r2 | **lost: r1 at 15-17** | 16 as Wrath above (+4.7 to 4.8%) |
| Balance | Moonfire | 17-19 r1 | **17 holds; 18-19 r2** | 18 as above; 19 A +1.10/+1.22, H +1.28/+1.41 |
| Restoration druid | tank Rejuvenation | 13-16 r1 (night elf) | **13-16 r1 every race** (13-22 merged); **10 r1** added | tauren 13 +0.69/+0.69, 14 +1.43/+1.31, 15 +5.87/+5.87, 16 +0.52/+0.44; 10 tauren +0.23/+0.31 |
| Restoration druid | tank Rejuvenation | 17-22 r1 | holds | no challenger at 17-19 |
| Restoration druid | party Rejuvenation | 13-16 r1 (night elf) | **lost at 13-14 (r2 every race); r1 at 15 every race; 16 night elf r1** | night elf 13 +3.57/+3.57, 14 +1.97/+2.04 (with HT r2); tauren 15 +5.87 (with tank r1); 16 night elf holds, tauren r3 holds |
| Restoration druid | party Rejuvenation | 17-24 r1 | holds | no challenger at 17-19; 10 night elf r1 added (+0.23/+0.23, with tank and HT r1) |
| Restoration druid | Healing Touch (line 9) | 14-16 r2 (tauren) | holds; **14 r2 every race** | night elf 14 +1.97/+2.04 |
| Restoration druid | Healing Touch (line 9) | 14-16 r1 (night elf) | **15-16 only** | 15-16 night elf holds (r2 -0.3 to -0.6); 10 night elf r1 added |
| Restoration druid | Healing Touch (line 9) | 17-19 r1 | holds | no challenger |
| Restoration druid | Regrowth | 18-19 r1 | holds | no challenger |
| Holy paladin | Holy Light | 18-19 r1 | holds | r1 over r2: 18 +4.3/+3.9%, 19 +8.2/+8.1% |
| Protection paladin | Holy Strike | 12-14 r1 | **lost: r2** | Dwarf 12 +2.71/+2.71, 13 +2.70/+2.66, 14 +2.16/+2.16; Undead 12 +2.49, 13 +2.75/+2.72, 14 +2.75/+2.75 (threat) |
| Protection paladin | Holy Strike | 15 r1 (Undead) | **lost: r2** | +2.42/+2.42 |
| Discipline | Renew | 14-29 r1 | holds at 14-19 | Renew r2 never a challenger |
| Holy priest | Renew | 14-27 r1 | **lost at 16 (r2)**; rows now 14-15 and 17-27 | 16 Human +0.61/+0.61, Undead +0.64/+0.73 |
| Elemental | Flame Shock (line 3) | 18 r1 (troll) | **lost: r2; dwarf r1 at 18** | troll r2 +1.0/+1.0; dwarf r1 +0.40/+0.44; orc level; tauren r2 +3.0, Windshaper Skyborne r2 +4.6 |
| Elemental | Flame Shock (line 3) | 19 r1 (orc) | **lost: r2; dwarf r1 at 19** | orc r2 +0.48/+0.32; dwarf r1 +0.68/+0.67; troll r2 +1.3, tauren +0.9, Windshaper +3.1 |
| Restoration shaman | tank Healing Wave (line 17) | 18-19 r3 | holds | r3 over r4: 18 +3.5/+3.2%, 19 +4.5/+3.5% |

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | **Balance Wrath and Moonfire below 20.** On gear with every armor slot filled, the cheap pair pays at more levels: Wrath rank 1 with Moonfire rank 1 at 12, 13, 15 and 16 (+1.0 to 4.8%), rank 2 of both at 14 (held) and 18 (+3.5 to 4.7%), Wrath 2 with Moonfire 1 at 17 (held), and Moonfire rank 2 with Wrath rank 3 at 19 (+1.1 to 1.4%). | grid1.tsv, conf.tsv; test `TestRound26BalanceRankRows`, `TestBalanceRankRowsBelowTwentyAndMoonfireAtThirtyFive` moved | CONFIRMED-BUG (rank rows) | R24 af0fb2481 (judged on the pre-P7 gear) | S c62caa3d5 |
| K2 | **Restoration druid heals below 20.** The tank's Rejuvenation rank 1 now wins from 13 for a tauren too (+0.4 to 5.9%), the party keeps rank 2 at 13 and 14 for a night elf (+2.0 to 3.6%) and takes rank 1 at 15 for a tauren, Healing Touch rank 2 at 14 for a night elf (+2.0%). At 10: the tank's rank 1 (+0.2% tauren; a night elf gains it only with rank 1 of all three). | as K1; tests `TestRound26RestorationDruidRankRows`, `TestDruidRankRowsSecondPassRound24` moved | CONFIRMED-BUG (rank rows) | R24 9ba42a787, 14c33aa48 | S c62caa3d5, ea819a7fc |
| K3 | **Protection Holy Strike rank 1 at 12 to 15 loses.** Rank 2 makes 2.2 to 2.8% more threat for a Dwarf and an Undead at every level; the cheaper rank no longer buys enough Seals and Judgements to pay for the weaker strike. Rows removed; rank 1 at 40 stays. | grid1.tsv, conf.tsv; tests `TestRound26ProtectionHolyStrikeTakesRankTwo`, `TestProtectionHolyStrikeRankRows` (renamed, 12-15 cases moved) | CONFIRMED-BUG (rank row) | R24 215ee9421 | S c62caa3d5 |
| K4 | **Holy priest Renew rank 2 at 16** (+0.6 to 0.7%, the level Heal arrives). The 14-27 rank 1 row splits around it. | test `TestRound26HolyPriestRenewAtSixteen` | CONFIRMED-BUG (rank row) | R24 92ba8e5db | S c62caa3d5 |
| K5 | **Elemental Flame Shock at 18 and 19.** Round 24's troll row at 18 and orc row at 19 now lose (rank 2: troll +1.0%, orc +0.3 to 0.5%); a dwarf keeps rank 1 at both (+0.4 to 0.7%). Every race of the class judged. | grid1.tsv, conf.tsv, v1.tsv; `TestRound24ShamanRankRows` moved | CONFIRMED-BUG (rank rows) | R24 08434b177, 3b78f3963 | S 066c1f711 |

### Rows that should exist (the sweep of every ranked spell the cards cast at 10 to 19)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K6 | **Discipline Power Word: Shield rank 1 below 20.** No row; the shield cast the newest rank. Rank 1 wins at 12 (+1.1 to 2.2%), 14 (+1.4 to 1.5%), 15 for every race but an Undead (+0.6 to 2.0%), 16 (+0.2 to 0.5%) and for a Human at 18 and 19 (+0.2% over rank 3; every other race keeps rank 3 by 0.3 to 0.6%). At 13 rank 2 holds (Human, Gnome, Night Elf, Troll keep it by 0.5 to 2.1%), and at 17. Every priest race judged at 13, 15, 18 and 19. | grid1.tsv, conf.tsv, v1.tsv, v3.tsv; test `TestRound26DisciplineShieldRankRows` | CONFIRMED-BUG (missing rank rows) | old (round 24 swept the shield at 36-40 and 54-55 only) | S 774e998df |
| K7 | **Discipline Lesser Heal rank 2 at 17** (+0.5% Human, +0.6% Undead, both seeds); rank 3 holds at 16, 18, 19. | grid1.tsv, conf.tsv, v1.tsv; test `TestRound26DisciplineLesserHealAtSeventeen` | CONFIRMED-BUG (missing rank row) | old | S 066c1f711 |
| K8 | Aligned at 10 to 19: Holy Light (rows hold, newest rank elsewhere), Protection Seal of Fury (rank 2 from 18), Restoration shaman Healing Wave, Elemental Lightning Bolt (newest rank; Bolt 2 with Flame Shock 1 at dwarf 19 ties), Holy priest Lesser Heal (not cast from 16), Discipline Renew. | grid1.tsv | Aligned | - | - |

## Seed 2: Combat and Assassination finishers at 14 to 17

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Combat never Eviscerated below 19**: Slice and Dice at one or two points ate every point. Grid of 46 finisher rules (Eviscerate at 2 to 5, Slice and Dice at a..b points or never) at 10 to 60: skipping Slice and Dice and spending every five points on Eviscerate wins at **every bare level from 10 to 60**, beyond error at both seeds and both factions: +4.8% at 10, +5.3 to 5.5% at 11-12, +2.5 to 3.0% at 13-15, +6.0 to 6.2% at 16, +2.8 to 3.5% at 17-19, +1.9% at 20, +3.3% at 24, +4.1 to 4.8% at 30, +4.4% at 40, +2.4 to 2.6% at 50, +1.1 to 1.3% at 58, +0.5% at bare 60. It holds with a sword and Sinister Strike too (Cruel Barb at 19 +2.7%, 24 +4.1%). The raid fight at 60 keeps Slice and Dice (794.5 vs 775.8, +2.4%). Why: below raid gear the white swings miss 27% and are dodged 6.5% of the time, so +20% attack speed for 9 to 12 seconds returns less than the 25 energy and the points would buy as Backstab and Eviscerate. The line now carries `currentLevel >= 60`; the bare-60 loss (0.5%) stays, since the APL cannot tell bare from raid. Eviscerate stays at five (pending with the owner). Forever leveling practice: Classic guides keep Slice and Dice up while leveling ([Icy Veins, sword leveling](https://www.icy-veins.com/wow-classic/sword-rogue-leveling-talent-build-from-1-to-60)), others skip it on short mobs ([Icy Veins, Subtlety leveling](https://www.icy-veins.com/wow-classic/subtlety-rogue-leveling-talent-build-from-1-to-60)); the sims are decisive here, so no tie to break. | rog1-3.tsv, `rogan.py`; test `TestCombatSliceAndDiceWaitsForSixty` | CONFIRMED-BUG (rotation line) | old (2026-11-02 guide rewrite; round 25 O2 flagged it) | S 6b082c691, E 2eab62e34 |
| R2 | **Assassination never Eviscerated at 10 to 22**: Slice and Dice at three points every time it dropped, Eviscerate at four never reached. Below 17, Slice and Dice at four once it drops and Eviscerate at five wins +0.2 to 0.5% at 10 to 16 at both seeds and factions; at 17 to 19 it is level, so the 60 rule takes over at 17. No Slice and Dice at all wins at 10-12 and 16 but loses 0.4 to 0.8% at 13-15 and 18-19, so it is not adopted. From 30 Assassination's Slice and Dice pays (no Slice and Dice -2.1% at 30, -4.3% at 42). | rog1, rog2, rog4.tsv; test `TestAssassinationLowLevelFinishersWaitForFive` | CONFIRMED-BUG (rotation lines) | old | S 1fcd842a9, E 282ade883 |

## Before/after (relative; absolute at seed 1)

Rank rows (bare band 20 at the level; the new combination against the old one, seeds 1 / 50001):

| spec | level | before (s1) | after (s1) | change |
|---|---|---|---|---|
| Balance, night elf | 13 | 10.92 | 11.27 | +3.1 / +3.0% |
| Balance, night elf | 16 | 20.11 | 21.07 | +4.8 / +4.8% |
| Balance, tauren | 18 | 27.24 | 28.20 | +3.5 / +3.5% |
| Restoration druid, tauren | 15 | 18.23 | 19.30 | +5.9 / +5.9% |
| Restoration druid, night elf | 13 | 14.27 | 14.78 | +3.6 / +3.6% |
| Protection, Undead (threat) | 14 | 28.38 | 29.16 | +2.8 / +2.8% |
| Discipline, Human | 12 | 20.62 | 21.07 | +2.2 / +2.2% |
| Holy priest, Undead | 16 | 23.35 | 23.50 | +0.6 / +0.7% |
| Elemental, dwarf | 19 | 24.22 | 24.38 | +0.7 / +0.7% |
| Elemental, orc | 19 | 23.03 | 23.14 | +0.5 / +0.3% |

Rogue finishers (10,000 iterations, seeds 1 / 50001, horde):

| spec | level | before (s1) | after (s1) | change |
|---|---|---|---|---|
| Combat | 12 | 17.92 | 18.88 | +5.3 / +5.4% |
| Combat | 16 | 26.89 | 28.50 | +6.0 / +6.1% |
| Combat | 30 | 72.37 | 75.83 | +4.8 / +4.8% |
| Combat | 50 | 199.37 | 204.16 | +2.4 / +2.4% |
| Combat | raid 60 | 794.51 | 794.51 | 0 |
| Assassination | 12 | 18.40 | 18.47 | +0.4 / +0.4% |
| Assassination | 16 | 28.16 | 28.25 | +0.3 / +0.4% |
| Assassination | 18 and up | | | 0 |

## Re-rank needed

- **rogue-combat, bare bands 20 to 50**: the Slice and Dice line no longer runs below 60, so the bare-band weights (attack speed against Backstab stats) and the published bare numbers move. Raid 60 and bare 60 are unchanged.
- Rank rows below 20 move no picks.

## Files to regenerate (integrator, at the pin)

- `TestRotationLadder` goldens (sim/request): rogue-combat rungs below 60 and rogue-assassination rungs below 17; druid-balance, druid-restoration, paladin-protection, priest-holy, priest-discipline, shaman-elemental rungs at 10 to 19 where a ladder level falls on a changed row.
- addon-data, then Data.lua (addonrotation): Combat's Slice and Dice line and note, Assassination's Slice and Dice and Eviscerate lines and notes, the changed rank rows and the Restoration druid and Protection notes.
- Coach starters only if the integrator's starter check reads the rogue lines (persona_starter.py --all).
- Fork APL copies are committed by hand (E 2eab62e34, 282ade883); the site copies of the notes-only APL changes (druid-restoration, paladin-protection) are refreshed in the site commits.

## Measure at refresh

- Combat bare bands 20 to 50 after the re-rank (published numbers rise by the margins above).

## For other lanes

- **melee (Subtlety)**: the integrator's optional Opportunity versus Master of Deception item at 10 to 12 was not taken (ship mode: no sweeps beyond this lane's rows). See Still open.
- **Anyone using `spec-breakdown -race`**: the tool accepts a race the class cannot be (Skyborne priests ran without error); only `rank_choices_test` refuses it. Worth a guard in the tool.

## For the guide lane

Shipped with the rows, `_prose.test.ts`, `_sections.test.ts` and `builds.test.ts` pass: rogue/combat.md (Levels 1 to 19, Levels 20 to 60: Slice and Dice waits for 60), rogue/assassination.md (Levels 1 to 19: spend at five until 17), druid/balance.md:127, druid/restoration.md:113, paladin/protection.md:131, priest/holy.md:114, priest/discipline.md:116, shaman/elemental.md (Levels 1 to 19, Flame Shock ranks). Card notes: Combat Slice and Dice, Assassination Slice and Dice and Eviscerate, Restoration druid tank and party Rejuvenation and Healing Touch, Protection Holy Strike.

## Still open

- L1 (logs).
- Subtlety Opportunity versus Master of Deception at 10 to 12 (melee lane's item, offered as optional; not taken in ship mode).
- Combat at bare 60 loses 0.5% with Slice and Dice; an APL cannot split bare from raid at the same level.
- The Restoration druid and Elemental rows keyed with the Skyborne races follow the grouping of the earlier rows; Skyborne druids were not re-judged below 20.
