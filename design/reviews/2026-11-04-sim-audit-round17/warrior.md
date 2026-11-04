# Round 17, lane warrior (Arms, Fury, Protection)

Site `wt/r17-warrior` on int/r17 d7fc51fe7; engine `lane/r17-warrior` on fork int/r17 f637ec32c (no engine commit this round). Client 1.60.1.70755 with the late hotfix re-layer. Client rows: the layered 70755 raw tables (`scratchpad/c755/fetch/data/builds/1.60.1.70755/raw/`), `talents/`, `trainables/`, `spellconst/` and `spelltips/warrior.json`, Wowhead Forever tooltips. Beta logs from foreverlogs.gg (report 2763 and the cached rosters).

Result: **0 confirmed sim bugs** for the class. One test-coverage change (the monotone-ladder guard now covers raid builds). No guide-visible change, so no guide text and no rendered review were needed.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| 1 | **Heroic Strike as a white swing for procs (round 16 F5).** The engine gives Heroic Strike and Cleave `ProcMaskMeleeMHSpecial \| ProcMaskMeleeMHAuto` (heroic_strike_cleave.go:84), so they proc Unbridled Wrath, spend Flurry charges and spend Reckless Swings charges. The rage bar still pays them no swing rage: it requires a pure auto mask (`core/rage.go:151`). **One client rule decides all three:** Unbridled Wrath 12322, the Flurry charge aura 12966 (3 charges) and Reckless Swings 1329555 (4 charges) all carry SpellAuraOptions ProcTypeMask 4 (done melee auto attack). So the question is only whether the server counts a next-swing ability as an auto attack. **Logs:** there are no launch logs yet (newest foreverlogs.gg report is 2814, "Test", 2026-11-04 12:24). Energize events still return `{"success":false,"error":"Internal server error"}`, so Unbridled Wrath's rage can't be read. The API serves no stack-dose events, so Flurry or Reckless Swings charge use can't be read either. The beta warriors that used Heroic Strike (Hades 98 casts and Deaken, report 2763, about level 34) never had Flurry, and no beta log holds a Flurry or Reckless Swings aura at all. Gore Drinker's charge aura (1323969) is mask 20 (auto plus ability), so it cannot tell the two apart. **Web:** the sources disagree. Aurana's Classic 2019 warrior sim (https://wow-aurana.github.io/bigdick/rtfm.html, fetched 2026-11-04) says "Heroic Strike (and Cleave) does not consume Flurry charges". Upstream wowsims (SoD, 2024-03-08, fork commit 543752637) made next-swing abilities count as autos "since they behave like autos in almost all circumstances", and said Unbridled Wrath "correctly ... triggers from HS and Cleave now", from SoD logs. A TBC Wowhead comment (https://www.wowhead.com/tbc/spell=13002, comment by tobarstep) says Unbridled Wrath procs "off things that use the swing timer such as autoattack, heroic strike and cleave". **Decision:** Forever runs on the 1.15 server line that SoD runs, and the SoD evidence is the newer of the two and log-based, so the engine stays as it is. A level-60 Forever log settles it (a Flurry removal on a Heroic Strike hit, or Unbridled Wrath energize events once the API serves them). If Heroic Strike turns out not to count, Fury loses about 18 Unbridled Wrath rage a fight (0.6 x 31.9 Heroic Strike hits, `bd-fury`), plus Reckless Swings and Flurry charge effects. | SpellAuraOptions 124699, 126501, 249144, 248613; foreverlogs.gg API 2026-11-04; sources above | OPEN (kept on best evidence; settles with a level-60 log) | - | - |
| 1b | **Queued Heroic Strike and the off-hand miss penalty.** Icy Veins Forever (Abide, "Forever Fury Warrior Stat Priority", updated 2026-11-03, https://www.icy-veins.com/wow-forever/fury-warrior-melee-dps-pve-stat-priority): "Queuing Heroic Strike no longer removes the extra offhand miss chance in Forever." The engine never lifts the dual-wield penalty for a warrior queue; `DisableDWMissPenalty` is set only by the hunter's Raptor Strike (`hunter/raptor_strike.go:117`). | Icy Veins Forever; `core/spell_outcome.go:743` | Aligned (Forever change, engine already right) | - | - |
| 2a | **Monotone ladders.** Every warrior ladder, leveling `build:` and `raidBuild:`, levels 10 to 60, keeps every point from the level below (306 rows dumped, zero drops). Protection 41 is `...001301` in both builds: Shield Slam is kept and the new point goes to Bastion. All 28 specs' raid ladders are monotone too (1428 rows). | ladder dump with `LadderTalentString` + `droppedTalents` | Aligned (round 16 fix holds) | - | - |
| 2b | **The guard covered only the leveling build.** `TestEveryLadderOnlyAddsPoints` walked `GuideBuildTalents` only, but round 16's bug hit both builds (Protection 41, Balance raid 42). It now walks `raidBuild:` too, for all 28 specs (56 subtests, green). This is test coverage, not a bug. | `sim/leveling/ladder_monotone_test.go` | coverage | - | S 49b36d196 |
| 3 | Fixed-rank casts, level by level: see "Level sweep" below. Every rank is the trainer's rank at that level for all three specs. | spec-breakdown `-level`, ladder talents | Aligned | - | - |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Raid breakdown (horde Orc, guide raid build, frozen BiS, seed 1): Mortal Strike 23.8, Overpower 19.4 (77% crit; 10.6 dodge windows plus 27.7 Bloodthrill windows), Slam 10.5, Execute 14.7, Spearing Strike 6.4, Reckless Swings 11.3, Rend 7.6, Hamstring 22.7 (Windfury only), Heroic Strike 7.9 at 90 rage or more, one Berserker swap for Recklessness, Blood Fury twice. | `bd-arms` | Aligned | - | - |
| A2 | Bloodthrill: 1289682 ProcTypeMask 20 (auto plus melee ability), so main-hand abilities can open Overpower too. Engine `ProcMaskMeleeMH` with Rend up, 10/20%, a 5 s window (1282733 DurationIndex 28). | SpellAuraOptions 245638 | Aligned | - | - |
| A3 | Slam (70755): 1.5 s cast, 15 s cooldown, 15 rage. Improved Slam 2 gives -0.5 s cast, -0.5 s global and -5 s cooldown, and keeps the swing. Engine `slam.go` matches. Mortal Strike rank 4 is +160 normalized; Execute rank 5 is 600 plus 15 a rage (EffectChainAmplitude 1.5 x 10, ranks 3/6/9/12/15). | trainables, SpellEffect 20662 | Aligned | - | - |
| A4 | Dead points: Sweeping Strikes is Mortal Strike's prerequisite (talents.json prereq). Enrage 5 is the Fury tier-3 gate filler for Reckless Swings; every other tier 0 to 2 Fury talent also does nothing single-target, and Enrage fires on raid AoE. | talents/warrior.json | Aligned | - | - |
| A5 | Hamstring under Windfury, raid horde, 10000 iterations: with the line 826.4 / 826.5, without it 813.4 / 813.3 (seeds 1 / 50001), so the line is worth +1.6%. It stays. | spec-breakdown `-rotation` | Aligned | - | - |
| A6 | Guide lead "Blood Fury held timing" (guide rounds 10 to 13): holding the second Blood Fury for the execute phase gives 824.2 / 824.3 against 826.4 / 826.5 on cooldown (-0.3%; shorter fights lose the second use). On cooldown stays, and the prose needs no change. | spec-breakdown | Aligned (lead closed) | guides-handoff | - |

## Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Raid breakdown (horde Troll, seed 1): Bloodthirst 22.1, Whirlwind 16.4 (both hands, 30% on the closest target per hand), Execute 24.1, Heroic Strike 34.1, Reckless Swings 10.6, Berserker Rage 5.8, Death Wish 1.09, Recklessness 1, Flurry up 84%. Unbridled Wrath rage 97.2 = 0.6 x (white hits + Heroic Strike hits). | `bd-fury` | Aligned | - | - |
| F2 | Bloodthirst rank 4: 48 + 45% AP (SpellEffect 23894 e0 48, e1 45), 6 s, 30 rage. Whirlwind: 25 rage, 10 s, Improved Whirlwind -2 s. | SpellEffect | Aligned | - | - |
| F3 | Normalized rage: Icy Veins Forever (above, 2026-11-03) now states main hand speed x 3.46 and **off hand speed x 1.73**, with crits double. The engine's off-hand factor 0.5 was an unmeasured assumption (`core/rage.go:73`); it now has a second source. | Icy Veins Forever | Aligned | - | - |
| F4 | Unbridled Wrath 12964: energize 10 tenths (1 rage), no two-hander clause, which agrees with Kaivax's post. A web listing of "2 rage with a two-hander" is stale. | SpellEffect 12964 | Aligned | - | - |
| F5 | Fury `build:` Arms points: Improved Charge 2 and Improved Tactical Mastery 3 do nothing in a raid (no Charge, no stance swap). They are tier-gate fillers: Anger Management and Deep Wounds need 10 Arms points, and the only tier 0 and 1 options left (Deflection, Improved Overpower) do nothing in a raid either. Improved Charge counts while leveling. | talents/warrior.json | Aligned | - | - |
| F6 | Persuader (22384) in the off hand: client InventoryType 13 (One-Hand) for Persuader and Ironwood Blade, so the pair is legal. `items/warrior.json` labels every one-hander `main_hand`. Neither weapon has an effect. | Item.csv, ItemSparse.csv | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Tank fight (alliance, seed 1): 1097 threat/s, Shield Slam 30.2, Revenge 30.0, Shield Block 78% up, Thunder Clap 7.2, Demoralizing Shout 5.1, Heroic Strike on 56 of 85 swings, chance of death 0%. | `bd-protection` | Aligned | - | - |
| P2 | **Focused Rage against the client mask.** 29787 EffectSpellClassMask (1852722926, 1057316, 1, 0), applied to each spell's family flags: it reduces Shield Slam, Revenge, Thunder Clap, Sunder Armor, Heroic Strike and Demoralizing Shout, but not Shield Block or Battle Shout. The engine's `SpellFlagOffensive` set gives the same answer for every tank spell. | SpellEffect, spellconst family_mask | Aligned | - | - |
| P3 | Impale's mask (16493) includes Rend and excludes Deep Wounds; the engine applies it the same way (`rendCritDamageBonus`, and no bonus on Deep Wounds). | same | Aligned | - | - |
| P4 | Stances: Defensive 7376 (-10% damage taken, -10% damage done, +30% threat), with Defiance multiplied on top; Berserker 7381 (+3% crit on every table, +10% damage taken, -20% threat); Battle 21156 (-20% threat). Engine `stances.go` matches. | SpellEffect | Aligned | - | - |
| P5 | Revenge rank 5: 109 to 133 + 0.25 AP, a 5 s category cooldown (SpellCategories, not SpellCooldowns). Shield Slam rank 4: 655 with a 4.57% spread, plus block value. Thunder Clap rank 6: 103 and a 20% slow (aura 319). Sunder Armor -450, Demoralizing Shout -196. Shield Block: 75% for 7 s, 2 blocks, 5 s. Engine matches. | spellconst, SpellEffect | Aligned | - | - |
| P6 | Arena Grand Master (19024) use, "Absorbs 810 to 1350 damage, 20 sec, 9 min": not modelled. It moves damage taken, not threat, and its equip clause only works in the Battle Ring and The Maul. | Wowhead Forever item 19024; ItemEffect 102030 | note (For other lanes) | - | - |
| P7 | WARRIOR-13 (Improved Shield Wall) stays with the owner. | | OPEN (owner) | guides-handoff | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Round 16's low-rank penalty (`core/spell.go` SpellLevelPenalty) leaves physical spells whole, so no warrior ability is touched. | fork diff 9f7f1f6e4..f637ec32c | Aligned | - | - |
| X2 | Generator dedup traps (Battle Shout 6, Mortal Strike 4, Hamstring 3, Charge 1 to 3, Thunder Clap 1 to 6 and Demoralizing Shout 5 kept SoD or duplicate ids): each spell file maps back to the trainer id, and the level sweep shows every trainer id cast. | `constants_auto_gen.go` notes; sweep | Aligned | - | - |
| X3 | Trainer-book ranks at 60 (Battle Shout 7 25289, Heroic Strike 9 25286, Revenge 6 25288) are in 70755 SkillLineAbility. They are pending with the owner ("level-60 trainer book ranks"), so not reopened. | trainables/warrior.json | pending owner | - | - |
| X4 | Manual Crowd Pummeler (Arms 30 to 39): a 3 min cooldown with 3 charges (ItemEffect 98990), so 1.53 uses in a 180 s fight is right. | ItemEffect 98990 | Aligned | - | - |
| X5 | Stale comment: `core/rage.go:18-21` says Dual Wield Specialization adds 20% off-hand rage a point; 70755 says 10% (50% at 5), which `dualWieldSpecializationOffHandRageMultiplier` uses. No behaviour change. | talents/warrior.json | note | - | - |

### Level sweep (seed item 3)

Bare preset, horde, levels 20, 23, 26, 29, 30, 31, 33, 36, 39, 40, 43, 46, 49, 50, 53, 56 and 59, with `-talents` set to the ladder string at each level. The frozen BiS bands still carry pre-round-15 talent strings, so a sweep without `-talents` misreads the build. Band 10 has no bare entry. Every cast rank is the trainer's rank at that level:

- Heroic Strike 285/1608/11564/11565/11566/11567
- Rend 6547/6548/11572/11573
- Overpower 7384/7887/11584
- Execute 5308/20658/20660/20661/20662
- Mortal Strike 12294/21551/21552
- Slam 8820/11604/11605
- Bloodthirst 23881/23892/23893
- Battle Shout 5242/6192/11549/11550/11551
- Charge 100/6178/11578
- Thunder Clap 8198/8204/8205/11580/11581
- Demoralizing Shout 1160/6190/11554/11555/11556
- Revenge 6572/6574/7379/11600/11601
- Sunder Armor 7386/7405/8380/11596/11597
- Shield Slam 23922/23923/23924

The talent spells appear where the ladder takes them: Fury Reckless Swings from 31, Whirlwind and Berserker Rage from 36, Bloodthirst from 40; Arms Slam from 43 (Improved Slam), Mortal Strike from 40, Recklessness from 50.

## Before/after (raid 60, 10000 iterations, guide builds via `rt`, frozen BiS)

No commit moves a number. Every warrior value equals $R/before.tsv:

| Spec | alliance s1 / s50001 | horde s1 / s50001 |
|---|---|---|
| warrior-arms | 820.4 / 820.4 (0.00%) | 826.4 / 826.5 (0.00%) |
| warrior-fury | 879.4 / 878.9 (0.00%) | 873.3 / 873.2 (0.00%) |
| warrior-protection (threat/s) | 1097.0 / 1096.3 (0.00%) | 1088.0 / 1088.0 (0.00%) |

## Files to regenerate

- None.

## Measure at refresh

- Carried from round 16: Arms and Fury race order and picks on the round 15 builds; Arms Alliance Thunderbrew's Boot Flask and the Frozen Heart picks.
- Protection trinkets if Arena Grand Master's absorb is ever modelled (P6).

## For other lanes

- Core (all melee): seed 1, Heroic Strike/Cleave counting as an auto attack for mask-4 procs, stays OPEN until a level-60 log. Note it applies to every on-next-swing ability: Maul and Raptor Strike carry the same mask in the engine.
- Items: Arena Grand Master (19024) use absorb not modelled (P6).
- Log tooling: foreverlogs.gg still answers `event_types=energize` with a server error and serves no stack-dose events, so no rage or charge question can be settled from its API.
- Engine comment nit: X5.

## For the guide lane

- Nothing new. The Blood Fury timing lead (A6) is closed: on cooldown is right, so no prose change.

## Still open

- Seed 1 (Heroic Strike as an auto attack for procs): needs a level-60 Forever log.
- WARRIOR-13 (owner).
- Warcraft Logs blocked (gcloud login), not worked around.
