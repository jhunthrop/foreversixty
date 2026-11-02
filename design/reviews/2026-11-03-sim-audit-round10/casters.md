# Sim audit round 10: casters, 2026-11-02

Specs: mage-arcane, mage-fire, mage-frost, priest-shadow, druid-balance, shaman-elemental.

## Method

- **Engine.** `lane/r10-casters` from the pin 4d3129cae. **Site.** `wt/r10-casters` from origin/main bb619a188, with `make simdb`. "Before" is the shared baseline (`r10/sb-before` run from `r10/site-before/sim`).
- **Client.** 1.60.1.70707 with the hotfix push layered in (the round-9 layered copy, fetched 2026-11-01 21:42 UTC): SpellEffect, SpellClassOptions, SpellAuraOptions, SpellMisc, SpellCategories, SpellPower, SpellCooldowns, SpellDuration, SpellCastTimes, SkillLineAbility; talent rank texts from `data/builds/1.60.1.70707/talents/`. Class masks were decoded against every learned spell of the class set (`mask.py` in the lane scratch dir).
- **Crit check.** For each spec at raid 60 the final crit stats (every-table, spell-only, melee-only) were printed and each main spell's observed crit rate (crits / attempts) compared with the stat plus that spell's own talents, times the hit rate. Every spell matched to 0.1 point except the rows below that turned into findings.
- **Runs.** spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions, bare 20 to 60 and raid 60. Grids: `before.tsv`, `after1.tsv` in the lane scratch dir.
- **Tests.** Each confirmed bug got a failing test first; I saw it fail, then pass. `go test -tags=with_db ./sim/core/... ./sim/mage/...` passes in the fork.

## Assigned items: crit side effects

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **Combustion's crit reached Ignite's burn.** Combustion added its crit to every Fire-school mage spell, so during Combustion the Ignite ticks crit more (raid Fire's Ignite ticks crit 19.5% against 17.6% with Combustion off; the base spell crit is 17.7). The buff's mask names Fireball, Scorch, Pyroblast, Fire Blast, Flamestrike, Blast Wave and Frostfire Bolt, the same mask as Critical Mass; Forever's Ignite burn is outside it. It now reaches those seven spells only. | Combustion buff 28682: aura 107 misc 7, mask 0x00C00017 / 0x1000. Ignite burn 412538: class mask 0x08000000. Test `TestCombustionCritDoesNotReachIgnite` failed (30 crit rating on the burn), then passed. | CONFIRMED-BUG (client masks) | E 2af230c0e |
| A2 | Ignite: 8% a rank (40% at 5/5) of a Fire spell crit's finished damage, two 2 s ticks, rolling, no double dip on % modifiers. It reads Fire-school spell damage crits (item procs included, as the client's proc mask 0x10000 with no class mask allows), never its own ticks. | 11119 (proc 0x10000), 412538 (CumulativeAura 1, 2 s period); `mage/ignite.go` | Aligned | - |
| A3 | Critical Mass (2% a rank, Fire mask) and Incineration (2% a rank, Fire Blast, Ice Lance, Arcane Blast, Scorch): the engine's masks equal the client's. Raid Fire Scorch crits 30.8% = (17.69 + 6 + 10) x 0.915 hit, with Combustion off. Critical Mass correctly misses Ignite. | 11115 mask 0x00C00017/0x1000; 18459 mask 0x2001_8012 (Arcane Blast, Fire Blast, Ice Lance, Scorch) | Aligned | - |
| A4 | **Improved Flamestrike counted twice.** Flamestrike's config added 5% crit a rank and the declarative mod added the client's 5% a rank again: 30% at 3/3 instead of 15%. No published rotation casts Flamestrike; the custom /sim panel and AoE requests would. | 11108: one aura 107 misc 7 effect, 15 at 3/3. Test `TestImprovedFlamestrikeIsFifteenPercentAtThreeRanks` failed (30), then passed. | CONFIRMED-BUG (no published number) | E 0673b3844 |
| A5 | Arcane Instability's crit is aura 290 with no mask (every table, wands and procs too); Arcane Impact 2% a rank on the Arcane mask (missiles included, 0x200000); Arcane Mind and Ice Shards +20% crit damage a rank on their masks. Raid Arcane: Frostbolt 20.5%, Arcane Blast 31.6%, Arcane Missiles 26.1%, each equal to the prediction. | 15058, 11242, 11232, 11207 | Aligned | - |
| A6 | Shatter (17/33/50%) applies to any mage cast against a Frozen target (Frostbite, Frost Nova or a Fingers of Frost charge). Raid Frost Ice Lance crits 69.6% = (16.85 + 10 + 50) x 0.906. But see A7. | 11170 text by rank; `mage/fingers_of_frost.go` | Aligned | - |
| A7 | **Arcane Missiles missiles spent Fingers of Frost.** The check for "a direct-damage cast of the mage's own" asked for the mage flag OR the APL flag, so every Arcane Missiles missile (a passive tick with the mage flag) spent a charge and took Shatter's crit. In raid Frost, missiles crit 17.9% instead of 15.3% and ate about a charge a fight that the Ice Lance line should have spent: Ice Lance went from 5.2 to 6.5 casts a fight after the fix. The engine's documented model (design/reviews/2026-10-07-frost-fingers-of-frost.md): channels, ticks and procs neither spend nor benefit. | `consumesFingersOfFrost` used `Flags.Matches(A|B)`, which is "either". Client 400669 is aura 262 on Deep Freeze's mask only, so Forever's "as if Frozen" behaviour is server-side and the engine's design note is the model. Test `TestArcaneMissilesMissilesDoNotSpendFingersOfFrost` failed, then passed. | CONFIRMED-BUG (engine implementation against its own model) | E 3b57afcab |
| A8 | **Precision missed the spells items cast.** Precision was a mod on spells with a mage class mask, so Kindling Stave's Ignition and wand Shoot missed 6% more than the mage's own spells (raid Arcane Ignition missed 14.6% against Frostbolt's 8.5%). Client 29438 is aura 55 (spell hit) with no mask: "all spells". The druid's Nature's Reach (16819, aura 55) already reached every spell. Precision is now a school hit bonus on every magic school; the mage's own spells keep exactly 6%. | 29438; `druid/talents.go` applyNaturesReach. Test `TestPrecisionReachesSpellsWithoutAMageClassMask` failed (0.00), then passed; `TestPrecisionAddsHitToEverySchool` now checks hit chance, not the mod. | CONFIRMED-BUG (client aura) | E 1499d3f21 |
| A9 | Winter's Chill is no longer crit in Forever: it empowers Cone of Cold (20% a stack, 5 stacks), and the engine does that. Era's "+2% Frost crit a stack" is the web's. | 11180 text; `mage/cone_of_cold.go` | FOREVER-CHANGE | - |
| A10 | Master of Elements: Fire and Frost crits refund 10% a rank of base cost, mage spells only, not on a free cast. Whether it refunds on a Clearcasting crit (it did in TBC) is open: the client row is a server-side dummy (aura 4). | 29074 | Aligned / OPEN (see Still open) | - |
| A11 | Clearcasting (Arcane Concentration 12536, Elemental Focus 16246): one charge, 15 s, -100%, spent at the proc's moment (round 9). Elemental Focus in Forever procs on any Fire, Frost or Nature damage cast (10%), not on a crit. | SpellAuraOptions ProcCharges 1; SpellDuration 8 = 15 s | Aligned | - |
| A12 | Balance: Moonfury is aura 79 on Arcane and Nature (all sources); Vengeance +20% crit damage a rank on Wrath, Starfire, Moonfire, Insect Swarm (client mask also names Hurricane and Roots, which do not crit or are not cast); Improved Moonfire 5% crit and damage a rank on Moonfire only; Nature's Majesty aura 290; Nature's Grace on non-periodic spell crits. Raid Balance: Starfire and Wrath 16.9 / 16.8% = 19.07 x 0.885; Moonfire 28.5% (+10 on direct and ticks); Insect Swarm 19.2%. Moonkin Aura in the raid preset is aura 290, all tables, and the druid's own form does not add it twice (`raidHasMoonkinAura`). "Natural Weapons" is not in Forever's tree. | 16896, 16909 (mask 0x600207), 16821, 1223082, 16880, 24907 | Aligned | - |
| A13 | Elemental: Elemental Fury +20% crit damage a rank on the shaman's Fire, Frost and Nature damage and Searing and Magma Totems, never heals; Totem of Thunder +1% on Lightning Bolt (overloads included, same spell); Elemental Devastation is melee crit only. Raid Elemental: Lightning Bolt 20.1% = 21.37 x 0.906 + 1, Lava Burst 19.4%, Flame Shock 21.1%, Searing Totem 19.3%. Call of Thunder does not exist in Forever and Tidal Mastery is heals only. | 16089, 461295 (relic 228176), 30160, 16194 | Aligned / FOREVER-CHANGE | - |
| A14 | Shadow: Shadowform's +100% crit damage on the priest's Shadow class spells; Inner Focus +25% on the next non-periodic priest spell; Early Demise +30% on Shadow Word: Death in the execute window. Raid Shadow: Shadow Word: Pain and Devouring Plague 17.3% (= 17.31), Mind Blast 16.8% (with Inner Focus 1.38 casts in 21.85), Shadow Word: Death 22.1%. Shadow Power (Era) is not in Forever. | 15473 eff 4, 14751, 1310076 | Aligned / FOREVER-CHANGE | - |

## Assigned items: Evocation and ranks

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Evocation: +1500% spirit regen (x16, the 5 s rule lifted), 8 s channel, 8 min cooldown, learned at 20; every 8 s channel takes exactly four 2 s regen ticks. Cast below 5% mana with more than 30 s left: once a fight at every bare band for Fire and Frost, 0.79 to 0.99 for Arcane, 0.67 raid Fire, 0.03 raid Arcane, none raid Frost (never short). At Arcane bare 40 (spirit 189) one Evocation returns about 3,850 mana, more than the 3,560 pool, so it refills the bar. | 12051 (aura 110 1500, aura 134 100, dur 31 = 8 s, cd 480000); breakdowns | Aligned | - |
| R2 | Elemental raid 60 casts Lightning Bolt rank 9 (15207, 190 mana) all fight. Rank 10 (15208, 220 mana) measured -0.6% (466.9 to 464.2, seed 1; 467.2 to 464.6, seed 50001), out of mana 0.3 s to 5.5 s. Rank 9 holds beyond error. Every caster rank-choice row in data/curated/rank-choices.json casts its own client amount and cost (spellranks; Lightning Bolt checked rank by rank against SpellPower and SpellEffect). | `ele-lb9.json`, `ele-lb10.json` | Aligned | - |
| R3 | Fire's Improved Scorch line (round 9: +0.1 to +0.6%) stays held. A1 moves raid Fire by -0.1%, which does not change the line. | before/after grid | Held | - |

## Fresh re-audit

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | **Frost's Ice Barrier note said 3%.** The player note ends "its extra 3% holds all minute ... only the first 3% stays", while its first sentence and the engine carry the 70707 numbers, 15% for a minute (Frigid Power) and 5% more while the shield holds. Note text only. | 1326891 eff 0 = 15; 13033 eff 1 = 5; `mage/ice_barrier.go:31` | CONFIRMED (note) | S eafd5f520, E f69639907 |
| F2 | Mage talents read rank by rank against the talent texts: Arcane Meditation 17/33/50, Accelerant 0.1 a rank, Improved Fire Blast 1 s a rank, Heating Up 25% a stack, three stacks, 20 s (SpellDuration 18), Pyroblast 6 s (SpellCastTimes 171), Fire Power 3% a Fire effect, Frost Channeling 5% a rank, Frostbite 5% a rank, Missile Barrage 40/20%. Ice Barrier 15% + 5%. | talents/mage.json; client rows | Aligned | - |
| F3 | Balance: Improved Wrath -0.5 s and -50%, Moonglow 25%, Genesis 2% a rank (hotfixed text), Nature's Splendor, Improved Starfire -0.5 s, Eclipse 0.5 s on two Starfires, Nature's Reach +4% hit on every spell and attack. Idol of the Moon (28854, "up to 33") is spell power on Moonfire, as the engine has it. | talents/druid.json; 28854 | Aligned | - |
| F4 | Elemental: Convection 10%, Concussion 5%, Call of Flame 15% and +9 s, Elemental Alacrity 0.5 s, Lightning Overload 3/7/10% at half damage, Elemental Mastery 15 s / 3 min. Lightning Bolt ranks 1 to 10: amount, per-level growth, coefficient and cost equal SpellEffect and SpellPower. | talents/shaman.json; client rows | Aligned | - |
| F5 | Shadow: Darkness 2% a rank, Shadowform 10%, Shadow Weaving 2% x 5 for 15 s (aura 270 from the caster; the engine's dealt multiplier is the same on one target), Improved Mind Flay 10% a rank compounding, Improved Mind Blast 2.5 s, Shadow Focus 1% a rank on Shadow spells (raid Mind Blast misses 10.4% = 17 - 1.6 - 5), Twin Disciplines on Mind Flay, Devouring Plague, Shadow Word: Pain and Shadow Word: Death (client masks equal). | talents/priest.json; 1225132 masks | Aligned | - |
| F6 | Racials of the published races: Orc Blood Fury (20572: 10% attack and spell power, 15 s, 2 min, 2 casts a fight), Troll Berserking, Undead Touch of the Grave (5% max health as a flat leech that does not crit), Tauren. | 20572, 20554, 1260198; `core/racial_blood_fury.go` | Aligned | - |
| F7 | Raid talent searches on the lane engine with `-keep` from raid-keeps.json (seed 7, 800 final iterations): Fire and Balance, the guide is best; Arcane +0.5%, Shadow +0.5%, Elemental +0.6%, Frost +1.0% at seed 7. Frost's winner (Arcane Impact for Arcane Concentration and Accelerant) re-measured at 10,000 iterations: +0.5 / +0.4% horde and +0.5 / +0.3% alliance (seeds 1 / 50001), under the 1% bar. Every guide build holds. | `ts-out/*-raid.md`; spec-breakdown | Held (adoption rule) | - |
| F8 | Bare-band mana: Fire bare 20/30 and Frost bare 20 are out of mana 50 to 68 s of 180 even with Evocation; Shadow bare 20/30 about 90 s; Elemental bare 20 92 s. These are the bands with no cheaper rank that wins (round 9, band 20 pays the sub-20 penalty) and no mana cooldown before 40. Not a model error. | breakdowns | Aligned | - |

## Before/after

Before: pin 4d3129cae on origin/main bb619a188. After: lane engine f69639907 on site eafd5f520. 10,000 iterations. Rows that move less than 0.05% at every seed and faction are left out: every Shadow, Balance and Elemental band, Arcane bare 20 to 50 and raid 60, Fire bare 20/30, Frost bare 40/50/60.

| Spec, band | before (horde, seed 1) | after | horde (seed 1 / 50001) | alliance (seed 1 / 50001) | cause |
|---|---|---|---|---|---|
| mage-frost raid60 | 594.1 | 614.0 | +3.3 / +3.5% | +3.5 / +3.5% | A7 |
| mage-frost bare20 | 44.0 | 44.4 | +0.9 / +0.9% | +0.7 / +0.7% | A8 (wand Shoot gains Precision; the bare-20 mage wands when dry) |
| mage-frost bare30 | 75.7 | 75.9 | +0.3 / +0.3% | +0.1 / +0.1% | A8 |
| mage-arcane bare60 | 329.5 | 329.9 | +0.1 / +0.1% | +0.1 / +0.0% | A8 (Ignition, wand) |
| mage-fire raid60 | 660.8 | 660.1 | -0.1 / -0.1% | -0.1 / -0.1% | A1 |
| mage-fire bare40 | 153.6 | 153.4 | -0.1 / -0.1% | -0.1 / -0.1% | A1 |
| mage-fire bare50 | 190.2 | 189.9 | -0.2 / -0.2% | -0.1 / -0.2% | A1 |
| mage-fire bare60 | 329.6 | 329.8 | +0.1 / +0.1% | +0.0 / +0.1% | A8 |

## For other lanes

- **Integrator.**
  - Fork goldens: TestP1Mage regenerated (E 438f79a38); its character is the Forever Frost build, so every row moves with A7 and A8. sim/request TestRotationLadder goldens move for mage-frost (raid and bare 20/30), mage-fire and mage-arcane bare 60. I did not regenerate them.
  - A8 can move the published caster `hit_to_cap` for mages: the engine's hit profile takes the lowest-hit damaging spell, and wand Shoot now carries Precision. Held output; it changes at the refresh.
  - Frost's stat weights and BiS at raid 60 should be re-ranked after A7 (Ice Lance and Shatter weigh more).
  - No commit moves a spec outside this group: every engine change is in sim/mage.
- **Crit lane.** No crit-table finding outside the mage code. Every caster crit talent sits on the right table per its client aura (57/107 spell, 290 every table, 52 melee).

## Still open

- Master of Elements on a Clearcasting crit: the engine refunds nothing when the cast was free. TBC refunded the base cost anyway; the Forever client row is a server-side dummy. A level-60 log with a Clearcasting Fire crit and a mana line would settle it.
- Lightning Overload's 3/7/10% comes from the talent texts; the spell row carries 30 (the rank curve sets it). Single source plus the tooltip.
- Carried from round 9, blocked on level-60 logs: Kindling Stave's rate, Ignite's 0.73, Ice Lance's coefficient, Missile Barrage on a miss, Shadowfiend, Searing Totem's 2.2 s attack, Starfire rank 7 at trainers, Wrath's cost stacking.

## Sources

- 1.60.1.70707 client tables with the hotfix push (round-9 layered copy, fetched 2026-11-01 21:42 UTC). Spells read: 11069, 11078, 11095, 11100, 11103, 11108, 11115, 11119, 11124, 11129, 11170, 11180, 11207, 11213, 11232, 11242, 11426, 12051, 12536, 13033, 15058, 15257, 15258, 15473, 16035, 16038, 16089, 16164, 16166, 16246, 16578, 16821, 16850, 16880, 16896, 16909, 18459, 18809, 20554, 20572, 22959, 24907, 28682, 28854, 29074, 29438, 400588, 400589, 400624, 400625, 400647, 400669, 408248, 408438, 412538, 461295, 1223081, 1223082, 1223083, 1225132, 1260198, 1310076, 1326891, and the Lightning Bolt ranks.
- Talent rank texts: data/builds/1.60.1.70707/talents/{mage,priest,druid,shaman}.json.
