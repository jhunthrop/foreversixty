# Sim audit round 14: healers

Lane: site wt/r14-healers and fork lane/r14-healers, both from int/r14 (site 72471fcb3, fork 5195a8889). Specs: Restoration druid, Holy paladin, Discipline and Holy priest, Restoration shaman, and the heal profile. Client 1.60.1.70753 with the 2026-11-03 push (the push moves hunter rows only; no healer row changed). Every run was sequential under `nice -n 15`, GOMAXPROCS=9.

The fixes were measured on the lane before it merged int/r14 (engine 5195a8889). Both branches then merged int/r14 (site f882bc7b1, fork 4bed3e869) to pick up regcast's Holy Heal row change and casters' Inner Focus fix. The "final" columns below are on the merged trees, so they include other lanes' changes.

The short version:
- **The heal scale goes from 1.85 to 2.3, by the profile's own retune rule.** At 1.85 the druid's raid set covered 95.5% and its bare set 59.8%. At 2.3 the druid's raid set covers 83% and every bare set covers under half. No multiplier meets every clause of the rule (see A1).
- **Holy's tank Flash Heal was rank 2 at raid 60** while the list finished fights with mana unspent. Rank 7 is +10.3 to +10.9%.
- **Renew is rank 1 for both priests at 60** (it already was at every band below 60). **Holy's Heal is rank 1 in a raid.** Together with the Flash Heal change, Holy raid 60 is +11.5 to +12.0% and Discipline raid 60 +1.6 to +1.8%.
- The level-60 split from Warcraft Logs is still blocked by gcloud auth (A2).

## Coordinator items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **The heal scale was outside its own retune rule.**<br>- **The rule** (heal-profile.json notes; design/reviews/2026-10-08-heal-profile-retune.md section 2): take the smallest multiplier at which (1) the strongest raid set covers about 80% and the weakest about 60%, (2) every bare set covers less than half, and (3) mana is a live constraint.<br>- **At 1.85** (10,000 iterations, both factions, seeds 1 and 50001, published gear): the druid's raid set covers 95.5%, the shaman's 86%, Holy 73%, paladin 66 to 67%, Discipline 67%. The bare druid covers 59.5 to 59.9%. That breaks clause 1's ceiling and clause 2.<br>- **Sweep** in 0.05 steps, the step earlier retunes used (raid coverage druid / shaman / Holy / paladin / Disc; bare druid):<br>&nbsp;&nbsp;- 2.25: 84.4 / 63.7 / 59.8 / 56.7 / 55.2; bare druid 50.2% alliance, 49.9% horde<br>&nbsp;&nbsp;- 2.30: 82.8 / 61.3 / 58.5 / 55.7 / 54.0; bare druid 49.2 / 48.9%<br>&nbsp;&nbsp;- 2.35: 81.2 / 59.0 / 57.2 / 54.6 / 52.9<br>- **2.3 is the smallest multiplier that meets clause 2 on both factions.** Its strongest raid set is 83%. Mana stays live: every bare set runs dry (49 to 254 s), and the paladin, Discipline and shaman raid sets run dry too (96 to 171 s).<br>- **Clause 1's floor cannot be met at the same time.** The druid out-heals the weakest healer by 1.43 to 1.53 times, wider than the rule's 80-to-60 spread (1.33), so no multiplier meets every clause. At 2.3 the weakest raid sets cover 54 to 55% (paladin and Discipline). The strict clause, "less than half", decides; the profile notes now say so.<br>- **Other healers under the same rule:** no healer has its own scale. The profile is one fight for every spec, so it ranks gear, not fights. Every healer moves with the one scale (Before/after). After the priest fixes below, Holy covers 65 to 66% at raid 60 on the merged tree.<br>- **Bounds** move to [1.85, 2.3] (the old and new answers), as the 1.65-to-1.85 retune did. Every band's tank hit and pulse follow the scale (band 20 147/35 to 183/44; 30 274/63 to 341/78; 40 482/116 to 600/144; 50 1037/250 to 1289/311; 60 2128/555 to 2645/690). | Tests changed first and failing: `TestCuratedHealProfileStatesItsTuning`, `TestHealProfileAtScaleRescalesTheHitAndThePulse`, `TestHealProfileRefusesATuningThatDisagreesWithItsFigures`, `TestHealProfileAtTheReferenceBandIsUnchanged`, `TestHealProfileScalesEachBandByItsContent`, `TestHealProfileAttachUsesTheHealersBand`, `TestHealProfileRefusesABandFigureThatIsNotTheRules`. Sweep files were written by a scratch test that called `AtScale` and `figuresFor` (removed). | CONFIRMED-BUG (profile, by its own rule; same class as round 7 X8) | R13 383d137f1 (D2's rank 1 party Rejuvenation took the druid from 85 to 95%) | S 78d5b20f7 |
| A2 | **The level-60 tank and group split from Warcraft Logs is BLOCKED.** gcloud's login expired: `gcloud secrets versions access` fails with "Reauthentication failed" (the coordinator checked on 2026-11-03). I did not try to work around it.<br>Round 13's 12 cached responses are still in `.../scratchpad/r13/healers/wcl/cache/`. As round 13 A1 said, they hold one usable kill: Onyxia, report ZCgbD2ydYN8rtkVw fight 2, tank share 0.312. One kill is not a sample, so level 60 keeps 0.659.<br>**Steps to finish it:**<br>1. The owner runs `gcloud auth login`.<br>2. With the WCL_V1_API_KEY secret in the environment (never written to disk): `WCL_KEY=... python3 -I design/reviews/2026-11-03-sim-audit-round13/healers/raidsplit.py design/reviews/2026-11-03-sim-audit-round13/healers 10`.<br>3. The heal-profile change, tests first. Today `validateSplit` (sim/request/healprofile_bands.go:362-367) refuses any `split` on the reference band and pins its `tank_share` to the base hit and pulse ratio. The 60 row needs a `split` block like the sub-60 rows (fights, median, quartiles, source). `validateSplit` must accept a logged split at the reference band. `figuresFor` (healprofile_bands.go:278-295) must read `row.TankShare` at the reference too. `baseSplit` keeps the total incoming damage while the tank-to-pulse split changes. `tuning.base_tank_hit_damage` and `base_pulse_damage` then follow from the logged share, and the band tests and profile tests move with them.<br>4. Repeat A1's coverage check: a new split changes how much group healing counts. | `wcl.py`, `raidsplit.py` (round 13) | BLOCKED | - | - |
| A3 | **Discipline's 0% haste tie** belongs to lane core (owner ruling 2). Core's DONE file did not exist when this report was written, so I did not re-measure it. | - | Not mine | - | - |

## Holy priest

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **The tank's Flash Heal was rank 2 (9472) at raid 60.**<br>- **Why it was rank 2:** round 9 picked it when the list ran dry at 137 s with rank 7 (+2.3% then).<br>- **What changed:** later rounds raised Holy's mana (Prayer of Mending ungated, Litany of Light). The list now finished raid fights with mana to spare (mana lasts about 444 s at 1.85, 534 s at 2.3). Meanwhile the tank sat under 35%, so about 126 globals a fight went on a 600-point Flash Heal.<br>- **Rank sweep at raid 60** (3,000 iterations, seed 1; alliance / horde): rank 2 711 / 693; rank 4 733 / 715; rank 5 742 / 722; rank 6 770 / 753; rank 7 785 / 767. The same order holds at 1.85.<br>- **Measured** (10,000 iterations): rank 7 is +10.35 / +10.34% alliance and +10.84 / +10.86% horde (seeds 1 / 50001). The fight now runs dry at about 115 s.<br>- **Below 60:** the line's rank rows (40-49 rank 4, 50-59 rank 6, bare 60 rank 7) were exactly what the highest learned rank casts at each published band, so they are gone. Bare 20 to 60 are unchanged. Between bands (for example level 45) the line now casts the highest rank learned. | `TestHolyTankFlashHealIsTheTopRankInARaid` failed first, then passed. Round 9 ranks.md row "the tank's Flash Heal rank 2". | CONFIRMED-BUG | old (R9 dcd25dee3, right then, overtaken; round 13 P3 swept only Renew's rank) | S 2f9d5ad6d, E b4d7e0390 |
| P2 | **Renew rank 3 (6075) at 60, though both priests' rows already cast rank 1 at every band from 20 to 59.** Rank 1 (139) is learned at 8, and heals learned below 20 keep their whole coefficient in Forever (round 11 C1).<br>- **Holy** raid 60: 784.6 to 791.7 / 791.4 alliance, 767.9 to 775.1 / 774.7 horde (+0.9%). Bare 60: +0.2% alliance, +1.6% horde.<br>- **Discipline** raid 60: +0.9 / +0.9% alliance, +1.1 / +1.2% horde. Bare 60: +4.3% alliance, +4.4% horde.<br>- Both seeds agree within 0.1.<br>Renew is authored at rank 1 now. Its rows are gone, and every band below 60 casts what it did. | `TestPriestsCastRenewRankOne` failed first, then passed. A sweep of every ranked line, every rank, at raid 60 (horde, 3,000 iterations) found only this and P3 beyond +0.6%. | CONFIRMED-BUG | old (the rank-3 Renew predates round 9; round 9 added the rows below 60 and kept 60) | S 2f9d5ad6d, E b4d7e0390 |
| P3 | **Holy's Heal is rank 1 (2054, level 16) in a raid.** On top of P2: raid 60 +1.97 / +1.99% alliance, +1.97 / +2.01% horde.<br>Bare bands lose with it: bare 30 -3.6%, bare 40 -0.9 to -2.6%, bare 60 -5.2 to -5.4%. So rows keep what each bare level cast: rank 2 at 22-27, rank 3 at 28-33, rank 4 at 34-59 (regcast's 9dba07e78 rank 4 at 50-59 kept), and rank 4 at bare 60. | `TestHolyHealRanksByBand` failed first, then passed. Regcast's `TestHolyHealAtFiftyCastsTheListsRankFour` passes on the merge. | CONFIRMED-BUG | old | S 2f9d5ad6d, E b4d7e0390 |
| P4 | **Fresh read** of the spells this list casts, against the 70753 rows. Prayer of Mending 1240827: 413 + 0.429, 5 charges, 10 s cooldown, 390 mana. Lightwell 27871: 1.5 s cast, 3 min, 10 min cooldown, 365 mana. Lightwell Renew 27874: 320 every 2 s for 10 s, no bonus coefficient. Binding Heal 1240774: 825 + 4.2 a level, 0.429, 380 mana. Each matches the engine (prayer_of_mending.go, lightwell.go, binding_heal.go). The engine's conformance tests (`TestHealingLaddersMatchTheClient`) pin every heal rank's level, cost, cast time, coefficient and roll to the vendored 70753 spellconst, which is byte-identical to data/builds/1.60.1.70753/spellconst. They pass. | rows.py; priest tests pass | Aligned | - | - |
| P5 | **Inner Focus's crit mask** (casters lane's question). 14751 effect 1 (SpellEffect 692357, aura 107, +25) masks are 0xd9f42c90 / 0x00098434 / 0 / 0x10200000. Checked against each spell's SpellClassOptions:<br>- Prayer of Mending (word 1 0x20) is in the mask. The engine gives it the crit. Aligned.<br>- Renew (word 0 0x40) is not in the mask. The engine excludes it. Aligned.<br>- Penance's cast spell (1316995, word 1 0x800000) is not in the mask, but its heal pulses (1316991, 402289, 1240723: word 1 0x10000, direct heals) are. The engine excludes every channel, so Penance's pulses get no Inner Focus crit. Whether the charge is still up when the pulses land is a server question (ProcTypeMask 87376 consumes it on the cast). It is worth at most one Penance a 3-minute cooldown, well under 0.1%. | client rows | OPEN (one in-game Penance under Inner Focus settles it) | - | - |

## Discipline priest

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Renew rank 1: see P2. | - | CONFIRMED-BUG | old | S 2f9d5ad6d |
| D2 | **Rank sweep of every other ranked line at raid 60** (horde, 3,000 iterations): Lesser Heal rank 1 +0.6% (only +0.3 to +0.4% on top of Renew rank 1). Penance, Shield, Binding Heal, Flash Heal, Prayer of Healing and Greater Heal at every other rank tie or lose. Lesser Heal stays at rank 3; a 10,000-iteration check at refresh can decide it. | `.../r14/healers/rs-priest-discipline.tsv` | Aligned (see Measure at refresh) | - | - |
| D3 | **Penance** 1316995: 12 s cooldown (category), 385 mana, instant first pulse plus two at 1 s. The heal helper 1316991 is 425, coefficient 0.19. Matches penance.go and the ladder test. Dark Sacrifice 1277328: 320 a tick every 3 s for 15 s, 10 min cooldown. Matches. | rows.py | Aligned | - | - |
| D4 | **Bare-60 rank rows** under the new scale: with all of them removed, Discipline loses 0.8% (Binding Heal rank 1 row), so the rows stay. | scratch build without the bare-60 healer rows | Aligned | - | - |

## Holy paladin

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Rank sweep at raid 60** (horde, 3,000 iterations): every Holy Light, Holy Shock and Flash of Light rank ties or loses against the list's (the list runs dry at 160 s). | `rs-paladin-holy.tsv` | Aligned | - | - |
| H2 | **Divine Favor** 20216: +100% crit on one heal (aura 107, misc 7), one charge, 2 min, 4% of base mana. Oath of Light 1328548 as round 7 H1. The conformance tests pin Holy Light, Flash of Light, Holy Shock (cooldown included) and Light's Vigil, and pass. | rows.py; healing_tables_test.go | Aligned | - | - |
| H3 | **Bare-60 rows** still win at 2.3: Holy Light rank 4 at bare 60 is +4.7% over the authored rank (both factions, both seeds). | scratch build | Aligned | - | - |

## Restoration shaman

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Rank sweep at raid 60:** Healing Wave rank 2 on the tank +0.7% at 3,000 iterations. That is within 3,000-iteration noise of the row file's bar. It is listed for a 10,000-iteration check at refresh. Everything else ties or loses. | `rs-shaman-restoration.tsv` | Aligned (see Measure at refresh) | - | - |
| S2 | **Fresh read against 70753:** Riptide 1239243: 841 + 4.9 a level at 0.214, then 161 a tick every 3 s for 15 s at 0.10, 6 s cooldown, 385 mana (riptide.go). Chain Heal 10623: 3 targets, 50% each jump (chain_heal.go chainHealJumpFactor 0.5). Mana Tide Totem 1325256: 13 s, 5 min, 60 mana (mana_tide_totem.go). Water Shield 408510: 3 charges, 3.5 s internal cooldown. Healing Wave 10396: 1344 + 5.5, 0.857, 3 s, 560 mana. All match. | rows.py; healing_client_test.go passes | Aligned | - | - |
| S3 | Bare-60 rows (Healing Wave rank 2, Riptide rank 2) still win at 2.3: +35% over the authored ranks. | scratch build | Aligned | - | - |

## Restoration druid

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **The tank's Rejuvenation at raid 60** is still best at rank 10 under the new scale (3,000 iterations, seed 1; alliance / horde): rank 10 1007 / 1005; 9 995 / 994; 7 979 / 978; 6 973 / 972; 4 969 / 969. Mana lasts 342 s, so the raid fight is still not mana-bound. | spec-breakdown -rotation | Aligned | - | - |
| R2 | **Swiftmend** 18562: 20% of base mana, 15 s cooldown. **Innervate** 29166: +400% spirit regen (aura 110) and +100% regen while casting (aura 134) for 20 s, 6 min, 5% of base mana. **Bloom** 1328622 is a hidden dummy the Rejuvenation expiry triggers. Each matches the engine. The druid's conformance test pins every Healing Touch, Regrowth, Rejuvenation, Tranquility and Wild Growth rank, and passes. | rows.py | Aligned | - | - |
| R3 | **The druid's lead now rests on rank 1 Rejuvenation keeping its whole coefficient.** The owner's in-game check of Rejuvenation rank 1's tooltip (ruling 5) is still pending. If it shows a sub-20 penalty, the druid's coverage drops and A1 has to be redone downward. P2 and P3 rest on the same rule (Renew rank 1 is level 8, Heal rank 1 level 16). | round 11 C1 | OPEN (pending the owner) | - | - |
| R4 | **Bare-60 row** (tank Rejuvenation rank 4) still wins at 2.3: +1.0%. | scratch build | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Mana model, re-read:** spirit regen per class (priest 12.5 + Spirit/4, shaman 17 + Spirit/5, druid and paladin 15 + Spirit/5, `core/mana.go:157`, priest.go, shaman.go). The five-second rule. Aura 134 stacking (round 6). Every one is what rounds 2, 6 and 8 established, and no 70753 row or push row changes it. | engine | Aligned | - | - |
| X2 | **Racials of the published races:** night elf and tauren druids, human and undead paladins and Discipline, gnome and troll Holy, dwarf and tauren shamans. They are unchanged since round 12 R1. The troll's Berserking (26297) fires twice a fight from the autocast line. | heal report | Aligned | - | - |
| X3 | **Band coverage** (regcast's note 3). The retune rule reads band 60 only. Below 60 the rule scales every band by the same multiplier. At 2.3 the druid covers 50 / 77 / 85 / 67% of bands 20 / 30 / 40 / 50 (was 61 / 92 / 93 / 79% at 1.85). Its healing-power weight at bands 30 and 40 will stay low while its band set covers that much. The profile states no band-level rule, so this goes to the owner. | final.tsv | OPEN (owner: does the coverage rule apply per band?) | - | - |
| X4 | Regcast's note 2 (the druid's band-30 tank rank row ended fights with mana unspent). At 2.3 bare 30 runs dry at 190 s, so the row no longer breaks the row file's rule. | final.tsv | Aligned at 2.3 | - | - |
| X5 | The fork test `TestStormstrikeUsesNormalizedWeaponDamage` (sim/shaman) panics, "No item with id: 18817", on base-engine 5195a8889 too. It is an Enhancement test, not a healer one. | `go test ./shaman/` | For the shaman lane | - | - |

## Follow-up: rank re-sweep at 2.3 and core's C4

The coordinator asked for a second pass. Both branches merged int/r14 again (site 88e92b54e, fork with core's cbdc60fd5 and 2fe671caa). The sweep covered every ranked line of Restoration shaman, Holy paladin and Discipline, at raid 60 and bare 60 on the 2.3 profile, horde, 3,000 iterations. Each candidate was then confirmed at 10,000 iterations, both factions, seeds 1 and 50001. The seeds agree within 0.05% everywhere, so every gain below is beyond error. Every fight still runs dry. Tables: `.../r14/healers/rs2-*.tsv`.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Discipline's party Lesser Heal is rank 1 (2050) at 60**, not rank 3. Raid 60: 662.6 to 664.7 / 665.0 alliance, 667.0 to 669.6 / 669.8 horde (+0.3 to +0.4%). Bare 60: 329.1 to 344.0 alliance, 338.4 to 353.0 horde (+4.3 to +4.5%). Mana lasts the same (50 to 108 s). Rows keep rank 2 at 4-9 and rank 3 at 10-59, so no band below 60 moves. Rank 1 is learned at level 1, and rests on the same full coefficient for low ranks as P2 (R3). | `TestDisciplineFillsWithLesserHealRankOne` (was `...RankThree`) failed first, then passed | CONFIRMED-BUG | old (the 2.3 scale made the fight heavier; round 11 chose rank 3) | S cf5fe63bb, E 887e6b747 |
| F2 | **The shaman's tank Healing Wave is rank 2 (332) at raid 60**, as it already was at 40-59 and bare 60. Raid 60 +0.67 / +0.68% alliance and +0.77 / +0.73% horde (744.7 to 749.8). Mana lasts 162 to 146 s. Rows keep the highest rank learned below 40 (rank 3 at 12-17, 4 at 18-23, 5 at 24-31, 6 at 32-39). The old 40-59 and bare-60 rows are now what the list casts anyway, so they are gone. | `TestRestorationShamanHealingWaveRanksByBand` failed first, then passed | CONFIRMED-BUG | old (round 13 S1 measured +0.7% at 3,000 iterations and did not confirm it) | S cf5fe63bb, E 887e6b747 |
| F3 | **Holy paladin's Holy Shock is rank 4 at bare 60** (a new bare row; the raid keeps rank 1): +0.31 / +0.29% alliance, +0.34 / +0.37% horde. | `TestHolyPaladinHolyShockAtBareSixty` failed first, then passed | CONFIRMED-BUG | old | S cf5fe63bb |
| F4 | **Not adopted:** <br>- Holy Light rank 3 at bare 60 is +4.2% at 3,000 iterations, but mana lasts 636 s against 142 s. That turns a fight that runs dry into one that ends with mana unspent, which the row file refuses.<br>- Discipline's Power Word: Shield rank 10 is +0.4 to +0.6% raid and +0.3 to +0.4% bare. The guide keeps the low rank deliberately, because a full-rank shield on a warrior tank eats the rage he needs. The sim does not model that cost, so the low rank stays.<br>- Discipline's Flash Heal rank 2 at bare 60: +0.12 to +0.56%. Alliance is under the 0.2% healer bar, so it stays.<br>Every other rank of every line ties or loses. The raid paladin sweep found nothing. | rs2 tables | Aligned / held | - | - |
| F5 | **Core's C4: Discipline at exactly 0% haste** sits 0.24% above +0.01% at raid 60 on this tree (669.6 against 668.0, horde, seed 1, core's FS_PROBE_HASTE patch on a scratch engine).<br>- **The slot:** Penance falls from 21.8 to 21.5 casts a fight. Pulses a cast fall from 2.79 to 2.78. Binding Heal (+0.3) and Lesser Heal (+0.4) take the time.<br>- **Penance gap histogram** (18 logged iterations; delay past the 12 s cooldown): at 0%, 76 casts land on the instant the cooldown ends and 20 wait 1.5 s. Hasted, 37 land on the instant and 39 wait 1.5 s. At 0% a 1.5 s or 2.5 s cast ends exactly as the cooldown does. Hasted it ends about 1 ms early, so Lesser Heal (1.5 s) starts and Penance waits a full cast. Lesser Heal starts 107 to 113 of the delayed casts.<br>- **A wait does not close it.** A line "wait for Penance when it is under 0.1 s from ready" restores the hasted cast count (21.8) and pulses (60.3), but HPS falls to 667.4 because the wait costs time elsewhere. The remainder is the pulse-against-tick tie (round 13 X3, second tie).<br>- Both ties come from the engine's 0.5 s timing grid, not from a Discipline line. | probe build at `.../r14/healers/probe-engine` (removed); cast logs `.../r14/healers/dcast-*.txt` | OPEN (core: the cooldown-ends-on-cast-end tie and the pushback-on-tick tie at 0% haste) | - | - |

## Before/after

Raid 60, 10,000 iterations, effective HPS. Before is $R/before.tsv (5195a8889 / 72471fcb3). Lane (pre-merge) is lane/r14-healers on 5195a8889 with this branch. Final is the branch after merging int/r14 (it includes other lanes). Absolute values are horde, seed 1.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 | horde s1 absolute | coverage at 2.3 |
|---|---|---|---|---|---|---|
| Restoration druid | +7.72% | +7.74% | +7.67% | +7.69% | 933.7 to 1005.3 | 82.7 to 82.8% |
| Restoration shaman | -11.59% | -11.64% | -11.56% | -11.58% | 840.0 to 742.9 | 61.1 to 61.3% |
| Holy priest | +12.00% | +11.83% | +11.53% | +11.44% | 708.7 to 790.4 | 65.0 to 66.4% |
| Holy paladin | +3.00% | +3.06% | +2.70% | +2.67% | 644.6 to 662.0 | 54.5 to 55.7% |
| Discipline priest | +1.61% | +1.70% | +1.72% | +1.75% | 655.7 to 667.0 | 54.5 to 54.9% |

- **The scale alone** (lane, before the priest fixes; seed 1 alliance / horde): druid +7.8 / +7.7%, shaman -11.6 / -11.6%, Holy -1.3 / -2.2%, paladin +3.1 / +2.7%, Discipline +0.7 / +0.6%. A heavier fight leaves the druid less overheal and runs the shaman dry sooner (mana lasts 268 to 162 s).
- **P1 to P3** on top of the scale: Holy +13.5% alliance, +14.1% horde; Discipline +0.9% alliance, +1.1% horde.

Bare bands, horde, seeds 1 / 50001 (final against $R/sb-before), effective HPS:
- Druid: 20 +2.0 / +2.0%, 30 +4.9 / +5.1%, 40 +13.4 / +13.4%, 50 +6.6 / +6.7%, 60 +2.1 / +2.2%.
- Shaman: 20 -1.9 / -1.9%, 30 -3.4 / -3.3%, 40 -2.3 / -2.3%, 50 -7.0 / -7.0%, 60 -0.2 / -0.2%.
- Holy: 20 -1.4 / -1.2%, 30 +0.3 / +0.3%, 40 +1.2 / +1.3%, 50 +11.4 / +11.6% (regcast's Heal rank 4 row), 60 0.0 / 0.0% (the scale's -1.6%, then P2's +1.6%).
- Paladin: 20 +2.6%, 30 +4.0%, 40 +9.3%, 50 -0.3%, 60 +0.1%.
- Discipline: 20 +5.2%, 30 +5.2%, 40 +1.9%, 50 +1.6%, 60 +5.6% (scale +1.3%, P2 +4.3%).

Every band below 60 moves only by the scale, apart from Holy 50 (regcast).

## Files to regenerate

- sim/request `TestRotationLadder` goldens for all five healers at every band (the heal profile is the ladder fight), and for priest-holy, priest-discipline, shaman-restoration and paladin-holy for P1 to P3 and F1 to F3.
- Every healer BiS file and its `heal_profile` block, and dps-per-point for the five healers (the nightly re-picks gear at 2.3).
- addon-data, addon-builds and Data.lua (priest rotations and rank rows).
- `make apl-sync` would rewrite the fork's forever_holy and forever_discipline copies with the site's notes. The ids are already synced by hand in E b4d7e0390, and sim/request/apl copies in S 2f9d5ad6d.

## Measure at refresh

- **The coverage check on the re-picked sets.** The ranker re-picks gear at 2.3, and A1 was measured on gear picked at 1.85. If the re-picked druid leaves the 80 to 85% band, or a bare set reaches half, repeat the sweep.
- **Every healer pick at 2.3:** talent builds (the adoption rule), stat weights and the sub-60 rank rows. The scale is 24% heavier at every band, so mana counts for more.
- Discipline Lesser Heal rank 1 and the shaman's tank Healing Wave rank 2 at raid 60, at 10,000 iterations (D2, S1).
- /tiers/healer order at both bounds (1.85 and 2.3).

## For other lanes

- **Guides** (the APL `notes` and guide prose that now contradict the rotation):
  - priest-holy.json, the tank Flash Heal line (now rank 7). Suggested note: "Flash Heal at the top rank on a tank who is really low: in a raid your mana outlasts the fight, so a big fast heal is worth its cost."
  - The Renew lines of both priests. Suggested note: "Renew on the tank all fight at rank 1: Forever keeps a low rank's whole healing bonus, so it heals the most per mana."
  - The Holy Heal lines. Suggested note: "Heal at rank 1 in a raid, the cheapest healing per mana you have; while levelling, the highest rank you know."
  - web/src/content/guides/priest/holy.md:70 ("Renew ... rank 3 costs about a quarter of the top rank's mana") should say rank 1. holy.md:74 ("On the tank the low ranks are enough") should say the top rank in a raid. Check discipline.md's Renew wording too.
- **Casters:** P5 answers the Inner Focus mask question.
- **Shaman / enhancement:** X5.
- **Core:** A3 is yours (no change from me).
- **wow-player review:** A1 (what a heavier fight means for published numbers) and P1 to P3.

## Still open

- A2 (blocked on gcloud auth; steps above).
- A3, once core's report exists.
- P5, R3 (owner's pending in-game checks), X3 (owner).
- Carried from round 13: D4, H4, S3, P6, A6 (walk X2).
