# Sim-accuracy round 30: the open items from round 29

Lane `round30-open`: round 29's "Open, for round 30" list (design/reviews/2026-11-07-sim-audit-round29/integration.md) without the warlock items (another lane), the icon regen and the coach generator comments. Site branch from origin/main 450eddafa, rebased on 1378671e1 (the warlock and hotfix lanes' landings); fork branch `lane/r30-open` on 7d0e129ad (the hotfix lane's pin). Client 1.60.1.70790.

Every sim ran on the Windows test box through `tools/remote-test.sh` (a scratch commit per batch carrying the job list; `spec-breakdown` built once per batch). From 14:28 UTC the batches ran one at a time, one sim at a time, under `nice -n 15` with `GOMAXPROCS=6` (coordinator throttle); the earlier parallel batches were killed and rerun serially. Nothing ran on the Mac but Python checks.

Method, unless a section says otherwise: bare band gear (the committed BiS files), the guide's leveling walk at the level (`spec-breakdown -level N`), 10,000 iterations, seeds 1 and 50001, both factions' default races. A change is adopted when it gains beyond the combined error at both seeds in every cell it touches (rank rows: the round 24 method, race by race where a row is race-keyed); tanks are judged on threat with the round 29 tank rule (chance of death may rise by under 5 points, absolute, at every level touched).

## Result

| # | item | verdict | fix |
|---|---|---|---|
| T1 | Below 60 the stand-in tank boss was 1.4 to 2x the dungeon table | CONFIRMED-BUG (old), fixed | S d812515f5 |
| T2 | Tank order steps the round 29 ruling touched, re-checked on the new boss | Protection warrior: CONFIRMED (old ruling input changed), fixed; bear: holds | S a36a17d9f |
| T3 | Protection paladin rank rows on the new boss | holds | - |
| B1 | Bear Maul gate at 40 to 53 tuned on the old order | CONFIRMED-BUG (R29 order + T1), fixed for every level | S 112267b31, F 4a3601e1f |
| A1 | Arcane: Precision before 40 with fewer Arcane points | HELD under the adoption rule (gain is real; needs the owner) | - |
| R1 | Rank rows on the moved levels (damage specs) | 5 rows changed, the rest hold | S b30508858 |
| R2 | Rank rows on the moved levels (healers) | 6 rows changed, Healing Touch holds | S cf57f8a96, c46ace75c |

Engine: no engine code changed. The fork carries only the synced bear rotation (`lane/r30-open` 4a3601e1f on 7d0e129ad); the pin moves to 4a3601e1f with the ladder goldens regenerated.

## T1: the tank fight below 60 (CONFIRMED-BUG, fixed)

**Finding.** `TankProfileForLevel` scaled the level-60 boss's damage by (band/60)^2 and the healers by (band/60)^1.5. Against the heal profile's sourced dungeon table (data/curated/heal-profile.json bands: each band's tank hit, from its dungeons' last bosses and 75th-percentile trash pulls with the logged tank share) the damage was 1.61x at band 20, 1.94x at 30, 1.96x at 40 and 1.42x at 50 (round 29 walks-heal T2).

**Source.** The heal profile's band rows are the only sourced dungeon damage the site has (cmangos classic-db creature_template and spell lists, client ContentTuning modifier 1.2, foreverlogs.gg beta tank share). The tank fight now reads them: a band's boss damage is the level-60 boss times `tank_hit_damage(band) / tank_hit_damage(60)` (0.069, 0.129, 0.227, 0.487). Below 20 a level scales down from band 20 by the curated exponent, as before.

**Healing.** Scaling the healers by the same ratio first (cut 2.4 to 2.8x) killed every bare paladin and bear at 20 to 39 (100% death at 3,000 iterations): the healing fell further than the damage. Round 29 had measured the stand-in's healing-to-damage ratio as close to a dungeon healer's and only the spikes against health as too big, so the healing keeps that ratio: damage scale times (band/60)^(1.5 - 2). Only the size of the swings against the tank's health changes. Band 60 is unchanged, so no raid figure moves.

Red first: `TestTankProfileScalesDamageAndHealingByTheHealProfileBandTable` (ae238f167 red, d812515f5 green).

**Before/after**, chance of death % (Horde / Alliance, seed 1, 10,000 iterations) and threat change:

Protection paladin:

| level | before H / A | after H / A | threat H / A |
|---|---|---|---|
| 20 | 32.5 / 28.4 | 3.4 / 2.5 | +3.7% / +3.9% |
| 25 | 53.8 / 50.4 | 10.1 / 8.9 | +2.5% / +2.6% |
| 29 | 54.2 / 52.2 | 9.3 / 7.6 | +3.0% / +3.3% |
| 30 | 96.8 / 96.5 | 47.4 / 49.1 | +1.3% / +1.2% |
| 32 | 87.8 / 86.7 | 20.3 / 20.2 | +1.2% / +0.9% |
| 34 | 73.3 / 73.1 | 7.4 / 7.9 | +3.0% / +3.0% |
| 36 | 73.1 / 73.1 | 7.0 / 7.5 | +3.3% / +3.3% |
| 39 | 69.7 / 68.6 | 6.2 / 6.0 | +4.0% / +3.9% |
| 40 | 41.5 / 38.6 | 0.5 / 0.6 | +3.2% / +3.2% |
| 45 | 42.3 / 39.8 | 1.0 / 1.0 | +3.1% / +3.0% |
| 50 | 26.3 / 23.8 | 2.5 / 1.9 | +1.1% / +1.5% |
| 59 | 25.0 / 23.3 | 3.2 / 3.0 | +1.9% / +1.8% |

Feral bear:

| level | before H / A | after H / A | threat H / A |
|---|---|---|---|
| 20 | 9.6 / 7.1 | 0.2 / 0.1 | -9.2% / -9.2% |
| 25 | 67.4 / 59.8 | 12.6 / 9.9 | -7.9% / -7.7% |
| 29 | 82.3 / 76.3 | 25.6 / 21.4 | -7.7% / -7.3% |
| 30 | 99.9 / 99.9 | 93.6 / 93.2 | -15.6% / -15.2% |
| 35 | 100.0 / 100.0 | 96.8 / 96.6 | -14.1% / -13.8% |
| 36 | 99.7 / 99.7 | 65.0 / 66.2 | -14.0% / -13.5% |
| 39 | 99.3 / 99.2 | 56.6 / 58.8 | -13.6% / -13.1% |
| 40 | 87.2 / 85.7 | 14.0 / 14.2 | -14.3% / -13.8% |
| 45 | 26.6 / 26.6 | 0.1 / 0.1 | -11.4% / -11.0% |
| 50 | 0.9 / 1.5 | 0.0 / 0.0 | -7.2% / -6.8% |
| 59 | 0.5 / 0.7 | 0.0 / 0.0 | -7.3% / -7.0% |

Protection warrior: 0% death before and after at every level 17 to 59; threat -3 to -9% (less rage from damage taken). The bear's and warrior's threat falls for the same reason.

**Still open (T1b).** A bare bear at 30 to 35 still dies in 93 to 97% of fights: its damage taken (100 to 104 a second) is above the band 30 healing (80 a second) until Frenzied Regeneration trains at 36. The paladin at the same band takes 86 to 90. That is attrition, not spikes, and the healing is the stand-in's ratio; a dungeon healer's tank healing per band (from the healer sims on the heal profile) would settle it. Left for the tank lane.

## T2: the tank order steps the ruling touched

All on the new boss, 10,000 iterations, seeds 1 and 50001, both factions, every level 16 or 17 to 59.

**Protection warrior (CONFIRMED, fixed, S a36a17d9f).** Round 29 held Anticipation 25 to 27, Last Stand 28 and Anticipation 29 and 30 ahead of the threat talents because the stand-in band 20 boss started to kill a tank at 28. On the new boss a bare warrior's chance of death is 0.0% at every level with every order tried, so the threat talents go first: Improved Sunder Armor 25 to 27, Improved Thunder Clap 28 to 30, Focused Rage 31 and 32, Anticipation 33 to 35, Last Stand 36, Anticipation 37 and 38; Concussion Blow 39, Shield Slam 40 and Bastion keep their levels (Anticipation still pays Shield Slam's row by 40). Threat, all four cells: +1.3 to +1.4% at 25, +2.8 to +3.1% at 26, +4.4 to +4.7% at 27 to 29, +5.8 to +6.0% at 30, +7.7 to +8.0% at 31, +10.7 to +11.2% at 32, +8.9 to +9.2% at 33, +7.6 to +8.5% at 34 and 35, +7.1 to +7.5% at 36, +3.4 to +3.6% at 37; 38 on identical. Two other orders failed: Anticipation after Bastion and Last Stand after Bastion both let the walk push Shield Slam to 46 (-25 to -29% at 40 to 45). Red first: `TestProtectionWarriorTakesDefianceBeforeAnticipation` updated (095734775 red, a36a17d9f green). Guide levels and the coach's Last Stand level (36) follow. Band 30 moves: warrior-protection is in `knownStaleBISFiles` until the re-rank.

**Feral bear (holds).** Every round 29 step is still in and its chance of death only fell. The steps round 29 rejected on death were re-tried: Naturalist cannot come before Natural Reaction or Heart of the Wild (Heart of the Wild pays Naturalist's Restoration row); Heart of the Wild and Naturalist before Natural Reaction gain 3 to 7% at 50 to 56 but lose 0.4 to 4.1% at 43 to 49; Heart of the Wild before Natural Reaction then Naturalist loses 0.4 to 4.1% at 43 to 51; Feral Swiftness's second point after 40 only delays Primal Bite to 42. None gains at every level it touches, so the order stays.

## T3: Protection paladin rank rows on the new boss (hold)

Consecration (20924) at 30 to 59 and Holy Strike (10333) at 38 to 43, every learned rank, tank fight, both factions and seeds. One cell moves beyond the error: Consecration rank 1 over rank 2 at 37 for the Horde default (+0.19 to +0.20%), not for the Alliance and not at 36 or 38. An isolated one-faction level is not a row; the rows hold.

## B1: the bear's Maul gate (CONFIRMED-BUG, fixed)

**Finding.** The Maul line queued at 15 rage, and at 40 to 53 only at 30 (round 16, tuned with Natural Reaction at 46 to 50). Swept on round 29's order and the new boss: gates 15 to 70 at every level 10 to 59, and bare and raid 60, tank fight, 10,000 iterations, seeds 1 and 50001, both factions. A higher gate wins almost everywhere from 25: Maul only replaces a white swing, so the rage it holds goes into Lacerate and Primal Bite. Below 25 every higher gate loses (-0.1 to -2.4%). From 36 to 39 the best gates raise chance of death by 5 to 6.4 points (Frenzied Regeneration needs the rage), so the tank rule caps the gate there at 25 (+4.2 at most).

**Fix.** Maul at 15 rage below 25, 20 at 25 to 29, 45 at 30 to 35, 25 at 36 to 39, 45 at 40 to 59 and 25 at 60 (bare and raid best at 25 to 35; 25 is best for the raid). Red first: `TestBearMaulGateFollowsTheLevel` (7e1fc81bd red, 112267b31 green); the spec-breakdown Primal Bite test still holds at 43.

**Before/after**, threat (Horde s1 / s50001 / Alliance s1 / s50001) and Horde seed 1 chance of death:

| level | threat | death % | worst death rise |
|---|---|---|---|
| 20 | +0.00 / +0.00 / +0.00 / +0.00% | 0.2 -> 0.2 | +0.0 |
| 25 | +0.64 / +0.77 / +0.67 / +0.72% | 12.6 -> 13.5 | +1.3 |
| 27 | +0.64 / +0.64 / +0.51 / +0.64% | 20.4 -> 22.7 | +2.3 |
| 29 | +0.86 / +0.93 / +0.66 / +0.74% | 25.6 -> 26.7 | +1.9 |
| 30 | +5.55 / +5.51 / +5.42 / +5.53% | 93.6 -> 96.0 | +2.4 |
| 33 | +6.39 / +6.29 / +6.14 / +6.27% | 95.4 -> 97.1 | +2.0 |
| 35 | +4.76 / +4.88 / +4.79 / +4.83% | 96.8 -> 97.7 | +1.6 |
| 36 | +2.67 / +2.55 / +2.53 / +2.36% | 65.0 -> 67.9 | +4.2 |
| 39 | +2.56 / +2.55 / +2.60 / +2.52% | 56.6 -> 59.5 | +3.7 |
| 40 | +2.00 / +2.08 / +2.02 / +2.01% | 14.0 -> 14.5 | +1.8 |
| 43 | +1.69 / +1.74 / +1.94 / +1.55% | 3.1 -> 3.1 | +0.3 |
| 45 | +1.65 / +1.63 / +1.78 / +1.72% | 0.1 -> 0.1 | +0.0 |
| 50 | +0.95 / +0.93 / +0.79 / +0.88% | 0.0 -> 0.0 | +0.0 |
| 53 | +1.03 / +1.04 / +1.03 / +1.11% | 0.0 -> 0.0 | +0.0 |
| 55 | +4.32 / +4.45 / +4.49 / +4.34% | 0.0 -> 0.0 | +0.0 |
| 59 | +3.76 / +3.70 / +3.64 / +3.64% | 0.0 -> 0.0 | +0.0 |
| bare 60 | +1.43 / +1.30 / +0.83 / +0.71% | 3.6 -> 3.6 | +0.3 |
| raid 60 | +0.79 / +0.58 / +0.65 / +0.54% | 0.0 -> 0.0 | +0.0 |

The rotation note and the guide's Rage paragraph name the levels. Raid 60 moves, so the bear's raid weights and picks go on "Measure at refresh".

## A1: Arcane Precision before 40 (HELD, needs the owner)

There is no slack for it: Presence of Mind (tier 5) needs 20 Arcane points with the 21 a level 30 mage has, and Arcane Power (tier 7) 30 with the 31 at 40. Any point of Precision before 40 delays Arcane Power to 43, and any before 30 delays Presence of Mind as well. Measured (every level 10 to 45, 10,000 iterations, both seeds and factions, against the shipped walk):

| order | where Precision goes | change |
|---|---|---|
| V1 | 10 to 12 (Presence of Mind 33, Arcane Power 43) | -2.2% at 10, -0.3 to -2.2% at 18 and 19, -0.5% at 20 (Horde) and 28; +0.1 to +7.4% elsewhere 11 to 42 |
| V2 | 21 to 23, after Arcane Blast (Presence of Mind 33, Arcane Power 43) | +1.3 to +6.1% at 21 to 27 and 29 to 42; -0.3 to -0.5% at 28 |
| V3 | 31 to 33, after Presence of Mind (Arcane Power 43) | +1.1 to +4.4% at 31 to 41, +0.3 to +0.5% at 42; never below the shipped walk |

V3 gains at every level it touches, including 40 to 42 where it has no Arcane Power (6% spell hit against a level+3 target is worth more than one 15 second Arcane Power in the fight). It is not a model fault: Arcane Power is modeled (30% damage, 30% more cost, 15 s, 3 min). But the adoption rule refuses a change that drops a talent the rotation casts, and V3 drops Arcane Power at 40 to 42; round 29 read the rule the same way (walksdps MA3). Held. If the owner rules that a three-level delay of a cast talent is allowed when the walk gains at each of those levels, V3 is ready: `first` gets Precision 3 after Presence of Mind.

## R1: damage rank rows on the moved levels

Cells round 29 listed (walksdps "Measure at refresh"), Affliction excluded: every level of each range, every learned rank (the top 3 first, then every rank where one moved), both factions' default races, seeds 1 and 50001; race by race (every legal race) where a row is race-keyed or the two factions split.

| spec, spell | levels | verdict | change |
|---|---|---|---|
| Arcane Frostbolt (10181) | 20 to 35 | CONFIRMED | rank 1 beats rank 4 by +1.2 to +2.0% at 24 to 26 and +0.4 to +0.9% at 27; on the hotfix engine (re-check below) also at 28 (+0.4 to +1.1%), so rank 1 at 24 to 28 and the rank 4 row from 29 |
| Arcane Missiles (10212) | 21 to 30 | holds | no rank beats the default beyond the error |
| Fire Scorch (10207) | 41 to 57 | CONFIRMED | rank 3 at 50 and 51 (+0.5 to +1.9%); rank 4 at 52 to 55 (+0.4 to +1.0% at 52 to 54; +0.1 to +0.5% at 55 for every race, beyond the error for a Human, a High Order Skyborne and a Troll); 56 and 57 hold |
| Fire Pyroblast (18809) | 40 to 44 | CONFIRMED | a Human leaves the rank 3 row at 41 (rank 4 +0.3 to +0.5%, both seeds); an Orc, a Troll and an Undead gain under the error and keep rank 3 |
| Balance Wrath (9912) | 42 to 58 | CONFIRMED | rank 7 beats rank 6 at 48 (+0.24 to +0.34%, all four cells); the rank 6 row ends at 47 |
| Balance Starfire (9876) | 42 to 58 | holds | - |
| Retribution Consecration (20924) | 30 to 33 | holds | (24 to 29 have only rank 1) |
| Shadow Mind Flay (18807), Shadow Word: Pain (10894) | 31 to 35 | hold | - |
| Elemental Lightning Bolt (15208) | 15 to 21, 30 to 33 | holds | rank 3 at 20 and 21 within 0.04% |

**Re-check on the landed hotfix engine.** The warlock and hotfix lanes landed while this lane ran (Arcane Concentration now splits over a Missiles channel, which Arcane has from 15 and Fire from 49). Every lane measurement above ran on pin 0e04b1284; the Arcane Frostbolt and Fire Scorch rows were re-run on the rebased branch (pin 4a3601e1f, 10,000 iterations, both seeds and factions, every rank near the row): every new row holds, and Arcane rank 1 also wins at 28. The tank, bear and healer items touch nothing the hotfixes changed below 60 except the Restoration druid's Nature's Swiftness on Wild Growth at 60.

Red first: `TestRound30DamageRankRows` (bea9aa80f red, b30508858 green); round 15's Wrath-at-48 case and round 24's Arcane Frostbolt-at-26 case moved to 47 and 28. The Arcane, Fire and Balance guides' rank lists follow.

## R2: healer rank rows on the moved levels

Cells round 29 flagged (walksheal X1), on each band's heal profile (`spec-breakdown -heal`), every learned rank near the row, both seeds; race by race (every legal race) for the race-keyed rows and wherever the two factions split. Red first: `TestRound30HealerRankRows` (6125139bd red, cf57f8a96 and c46ace75c green).

| spec, spell | levels | verdict | change |
|---|---|---|---|
| Holy Renew (10927) | 14 to 39 | CONFIRMED | rank 1 at 16 for every race (+3.2 to +4.6%; an Undead +0.3%) and at 28 and 29 (an Undead +1.8 to +3.7%, a Dwarf +3.3% at 29); 30: rank 2 for a Human, Gnome, Dwarf or Night Elf (a Human +2.7%), rank 1 for a Troll (+0.8%) or Undead; 31: rank 2 for everyone (+2.4 to +5.5%); 32: rank 1 for the Alliance races (+0.2 to +2.0%), rank 2 for a Troll or Undead (an Undead +1.9%); 33: rank 2 for everyone (a Dwarf +1.9%, a Human +0.6%); 34: rank 3 (a Human +0.9%, a Troll +0.3%), a Dwarf rank 1 (+0.1%); 35 to 39 rank 3 (a Human +1.4% at 35) |
| Discipline Renew (10927) | 33 to 42 | CONFIRMED | rank 3 for every race at 35 (+0.3 to +0.7%); a Human casts rank 5 at 41 (+0.2 to +0.3%) and the newest (rank 6) at 42 (+0.3%); every other race rank 4 at 42 (a Troll +1.7%, an Undead +0.1%) |
| Discipline Prayer of Healing (10961) | 50 to 59 | CONFIRMED | every race but the Undead: rank 1 at 50 to 54 (+0.2 to +0.6%), rank 2 at 55 to 58 (+0.1 to +0.5%; a Night Elf keeps rank 3 at 57, where rank 2 loses 0.2%); an Undead keeps rank 3 to 58; rank 2 for every race at 59 (+0.1 to +0.4%) |
| Restoration shaman Healing Wave (10396) | 37 to 46 | CONFIRMED | the rank 4 row at 37 to 41 is gone: the newest rank (6) wins at 37 to 39 (+0.1 to +0.9%), rank 5 at 40 and 41 (+1.2 to +3.2% over rank 7) and holds at 42 and 43; rank 6 from 44 (+0.2 to +1.8% at 44 and 45) |
| Restoration druid party Rejuvenation (8910) | 40 to 50 | CONFIRMED | rank 4 at 44 (+2.5 to +2.6%), rank 5 at 45 to 48 (+1.0 to +2.7%); 40 to 43 and 49 on hold |
| Restoration druid tank Rejuvenation (9841) | 56 to 58 | CONFIRMED | rank 9 at 56 to 58 (+1.8 to +2.0% over rank 6 at 56, +0.3 to +0.4% over rank 8 at 57 and 58); 59 keeps rank 8 (not re-measured) |
| Restoration druid Healing Touch (9889) | 30 to 39, 49, 57 | holds | no rank beats the row beyond the error |

Restoration druid party Rejuvenation at 40 to 42 was measured for all four druid races (rank 2 holds); the other druid and shaman cells on the two factions' default races, as both agreed.

## wow-player review

Every touched class reviewed on its guide and card text, fixes applied, then SHIP:
- **Warrior:** SHIP on the first read (Protection leveling levels, coach Last Stand at 36).
- **Mage:** HOLD on Fire (the Scorch and Pyroblast ranks sat at the end in a separate paragraph, with a race clause); the ranks moved into the 28 to 39 list as one line each and the Arcane Frostbolt sentence was reordered, then SHIP.
- **Druid:** HOLD on the bear's Rage paragraph (open-ended ranges, no reason for the 60 drop); closed ranges and the reasons went in, and the card note spells each range out; then SHIP. Balance and Restoration SHIP.
- **Priest and shaman:** HOLD (Holy Renew 30 to 33 incomplete, Prayer of Healing wording, Healing Wave going up, down and up without a reason); fixed, then SHIP.

## Gates

On the rebased branch (dc18c6cd0, pin 4a3601e1f), on the box, one run at a time:
- sim: `go vet ./...`; `go test` for request, rankchoices, leveling, cmd/leveling-bis, cmd/addon-builds, score and the bear and tank spec-breakdown tests (all ok). The ladder goldens were regenerated at the pin (the 15 goldens origin/main carried stale are current again); addon-builds.json follows the warrior walk.
- web: `npm test` on the guide, guides, bis and sim suites (176 files, 3,142 tests) and `astro check` (0 errors, 0 warnings).
- data: `pipeline coach-rules --check`, `pytest tests/test_genproto.py tests/test_addondata.py` (43 passed).
- apl-check against the fork branch.
- Red first for every fix: the red commits fail on the box and the fix commits pass.
- **Already red on origin/main, not this lane's:** spec-breakdown's `TestWarlockAfflictionAmplifyCurseBeforeEachAgonyBelow60` and `TestAfflictionRaidRotationKeepsBaneOfDoomAndAmplifiesTheLastAgony` fail on 1378671e1 as on this branch (the warlock lane moved Affliction to Agony over Doom at 60).

## Re-rank and Measure at refresh

- warrior-protection band 30 (talents moved; in `knownStaleBISFiles`).
- The three tanks, bands 20 to 50: every tank figure (threat, chance of death, TMI, the tank weights) below 60 moves with T1.
- druid-feral-bear every band and raid 60: the Maul gate.
- mage-arcane bands 20 (level 24 to 27 rows), mage-fire band 50, druid-balance band 40: rank rows.
- druid-restoration tank Rejuvenation: rank 9 at 56 to 58 (rows 49 to 55 and 58).
- Healer rank rows: priest-holy band 30, priest-discipline bands 30, 40 and 50, shaman-restoration band 40, druid-restoration bands 40 and 50 (picks and weights move with the casts).

## Logs

No new public level-60 logs were read for these items; none of them is settled by a raid log (leveling orders and gates, and a below-60 tank fight).

## Shas

Site: ae238f167 to dc18c6cd0 (22 commits) and the report commit. Fork: `lane/r30-open` 4a3601e1f (bear rotation sync only). Engine pin after: **4a3601e1f** (S 43b56ca0c, goldens and addon builds S 91e33364b, dc18c6cd0).
