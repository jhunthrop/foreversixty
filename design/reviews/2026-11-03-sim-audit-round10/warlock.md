# Sim-accuracy round 10: warlock (2026-11-02)

Specs: Affliction, Demonology, Destruction and their demons, bands 20 to 60, raid and bare presets. Client 1.60.1.70707 with the hotfix push layered in (layered tables in scratch `r9/casters/raw`, fetched 2026-11-01 21:42 UTC). Pin 4d3129cae, site origin/main bb619a188. Lane engine `lane/r10-warlock`, site branch `wt/r10-warlock`.

Method: spec-breakdown on the published band entries (`data/builds/1.60.1.70707/bis/warlock-*.json`), 10000 iterations, seeds 1 and 50001, horde troll and alliance gnome. "Before" is the shared baseline (`r10/sb-before` from `r10/site-before/sim`). Log evidence: the 31 cached foreverlogs.gg warlock character reports (scratch `fl/wl`, `r9/crit/fl`, `r9/crit/fl2`) with their roster snapshots (talent points per tree), aggregated by `r10/warlock/py/petcrit.py`; forever-bugs #331 read 2026-11-02.

Totals: 2 CONFIRMED-BUGs, both fixed in the fork. One moves every published warlock number by -0.3 to -0.6% (the demons' crit); the other (Spellstone) moves none.

Note on spec-breakdown's "Crit %": it is crits over attempts, misses included (`cmd/spec-breakdown/main.go:409`). The logs' figure is crits over landed hits. At a +3 target a demon spell misses 10 to 17%, so 2.7% per attempt is 3.3% per landed hit.

## Assigned item 1: crit split side effects on warlock spells

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | Every warlock damage spell rolls the spell table. Client DefenseType is 1 (magic) for Shadow Bolt, Incinerate, Immolate, Corruption, Agony, Doom, Siphon Life, Conflagrate (both id sets), Shadowburn, Soul Fire, Searing Pain, Rain of Fire, Wrack, Drain Mana, Death Coil, wand Shoot, Lash of Pain and Firebolt. Each engine spell is `DefenseTypeMagic`, so it reads Crit + SpellCrit (`core/spell_result.go:243`). | SpellCategories rows; `grep DefenseType sim/warlock` | Aligned | - |
| X2 | Crit talents land on their own spells only, as per-spell bonus crit, so they reach whichever table that spell rolls. Malevolence 1310949 (aura 107, misc 7, 1% a rank): its mask names the Shadow list plus Lash of Pain; the engine gives it to the warlock's Shadow spells (`talents.go:1102`) and Lash (`succubus.go:151`). Agonizing Flames 17927 effect 0 is aura 107 on Searing Pain only (3/7/10); effects 1 and 2 are +10% damage on the Destruction list and Immolate/Rain of Fire periodic (`talents.go:1296`). Fire and Brimstone 412751 effect 1 is aura 107 on Conflagrate only (8/17/25). Forever has no Devastation and no Improved Searing Pain (talents/warlock.json); the engine's Devastation is commented out. | SpellEffect masks resolved through SpellClassOptions set 5 (`r10/warlock/py/mask.py`) | Aligned | - |
| X3 | Crit-damage talents: Ruin 17959 (aura 108 misc 15, 20% a rank) masks Shadow Bolt, Incinerate, Immolate, Conflagrate, Shadowburn, Soul Fire, Searing Pain and Rain of Fire, which is the engine's `WarlockFlagDestruction` set (`talents.go:896`). Pandemic 427712 (33/67/100%) masks Corruption, Agony, Doom, Drain Life, Drain Soul, Siphon Life and Wrack, the engine's list (`talents.go:1018`). With a base bonus of 1, a 100% talent gives a 2.0x spell crit (`core/spell_outcome.go:1083`). | as X2 | Aligned | - |
| X4 | Crit triggers read the outcome, not a table: Improved Shadow Bolt on a Shadow Bolt crit (`talents.go:806`; Shadow Vulnerability is the warlock's own Shadow damage only). Decimation triggers on Shadow Bolt or Searing Pain below 35% health, not on a crit (440870 text and effect 2, 35; `soul_fire.go:157`). Nightfall is a 2/4% chance on Corruption, Drain Life, Drain Soul and Wrack ticks, with no crit (`talents.go:338`). Shadowburn, Conflagrate and Soul Fire use `OutcomeMagicHitAndCrit`. | 17793, 440870, 18094, 17941 rows | Aligned | - |
| X5 | The warlock's own crit after the split: base 1.7% spell (Era's column), Intellect at the warlock row of PlayerExpectedStat (0.000165 a point at 60), Brilliant Wizard Oil +1% spell only. Kindling Stave (+1%), Fel Cape (+1%) and Gloves of Spell Mastery (+2%) are item crit (critstrkpct), every table. Leader of the Pack and Moonkin Aura reach the warlock as every-table crit (aura 290, round 9). | crit.md; raw/wowhead/items.json | Aligned | - |

## Assigned item 2: the demons after the split

Round 9 measured Lash of Pain at 2.7% against 7.1% in the logs. The gap is mostly two artifacts and partly a real error:
- 2.7% is per attempt; per landed hit it is 3.3%, the Succubus's base.
- 7.1% pools every Lash of Pain in the logs, and the Affliction warlocks there (15 or more points, so Malevolence is reachable) crit 16 of 209 (7.7%). The rest crit 7 of 105 (6.7%).
- The engine's 3.3% is still low against the warlocks without Malevolence, and the engine's melee and Firebolt are far too high.

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Demons crit far too often in melee.** The engine gave the Succubus, Voidwalker and Felhunter Era's merged 3.34% on every table plus Agility at a warrior's rate: 7.6 to 7.8% against a target of their own level at every band (3.34 + 4.3 to 4.4). Against a +3 boss, after the 4.8% suppression, that was 2.8 to 3.0% of swings. Forever measures much less: <br>- Fierywind's test on #331 (level 30 warlock and Succubus on level 30 Cave Yetis, three logs): 12 crits in 1,718 swings, 0.70%;<br>- Blizzard's reply there (AidanZMoon, 2026-10-14): "Their melee crit is extremely low and their base spell crit is around normal levels";<br>- the 29 public logs: demon melee 33 in 2,400 (1.4%), Succubus alone 30 in 1,470 (2.0%) on mobs of mixed level, Voidwalker 0 in 452.<br>Demons now have 0.70% melee crit and Agility no longer raises it. Against a boss the suppression leaves them none. Test `TestDemonCritMatchesTheForeverLogs` failed first (7.69% at 20, 7.76% at 30, 7.73% at 40, 7.59% at 60) and passes. | forever-bugs #331 (read 2026-11-02); logs per character in `r10/warlock/py/petcrit.json`; `sim/warlock/pet.go:249` | CONFIRMED-BUG (one test, one Blizzard statement, 29 logs) | E de87d2393 |
| P2 | **Lash of Pain crit was too low.** It was 3.34%, Era's base. Forever measures 13 in 205 in Fierywind's test and 7 in 105 in the public logs of warlocks without Malevolence: 20 in 310, 6.4% (95% interval about 4 to 10%). Blizzard says the demons' base spell crit is "around normal levels". Lash of Pain is now 6.4%, plus Malevolence where taken. It is 1 to 3% of warlock damage. | as P1 | CONFIRMED-BUG (test and logs; the value is measured at levels 19 to 35) | E de87d2393 |
| P3 | **Firebolt crit was far too high.** It was the Imp's merged 3.45% plus Intellect at a mage's rate: 7.7 to 7.9%. The logs give 53 in 2,383 (2.2%) over 17 reports at levels 12 to 30, with no trend by level (2.0% at 12 to 18, 2.4% at 20 to 30). It is now 2.2%. No published rotation keeps the Imp out (Demonology's raid prepull sacrifices it), so no number moves. | logs as P1 | CONFIRMED-BUG (logs only, a large sample; flagged single-source) | E de87d2393 |
| P4 | Pets and the warlock's crit. Blizzard on #331: demons "don't benefit from their master's critical strike unless it's purely 'critical strike' as a stat". A raid warlock has 3 to 4% of item crit (Kindling Stave, Fel Cape, Gloves of Spell Mastery) and 6% of every-table party buffs, all in one engine stat. Whether the server passes the buff part as well is unpublished. Wiring the gear part would raise Lash by about 3 to 4 points and leave demon melee against a boss at 0. Not wired; single source. | #331; `pet.go` makeStatInheritance | OPEN: a level-60 log with the Succubus's Lash and melee crit next to the warlock's gear settles it | - |
| P5 | The Felhunter and Voidwalker cast nothing the engine models; they take the Succubus's spell value. No published rotation uses them. | `felhunter.go`, `voidwalker.go` | Aligned (moot) | - |

Pet crit plumbing in core is untouched. Everything above is in `sim/warlock`, so nothing goes to the crit lane except P4's question.

## Assigned item 3: round 9's held rows

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | Affliction: 1 Improved Drains to Improved Bane of Agony, re-measured on the lane engine (the pet fix moves Affliction): 571.0 to 575.1 (seed 1), 571.3 to 575.4 (50001), +0.72% / +0.72%. Still under 1%. Held. | spec-breakdown `-talents 25520010135001050-0005200000000000000-0550004000000000` | Aligned (held) | - |
| H2 | Destruction's Immolate below Conflagrate: no fix this round touches Immolate, Conflagrate or the warlock's crit, so it was not re-measured. Held at round 9's +0.1 to +0.2%. | round 9 A2 | Aligned (held) | - |

## Assigned item 4: demon sheets against Era's pet table

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | The sheets (Strength, Agility, Stamina, Intellect, Spirit, mana and weapon at 25, 40, 50 and 60, interpolated at 20 and 30) are wowsims Era's figures, one source. Forever's creature tables are server-side and no level-60 demon sheet or log exists. The crit columns no longer come from the sheet (P1 to P3). Agility now feeds nothing on a demon. Unchanged. | `succubus.go`, `imp.go`, `pet.go:petSheetAtLevel` | OPEN: a level-60 pet sheet or launch log | - |

## Spellstone (round 9 crit lane's note)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| W1 | **Spellstone was Era's +1% crit at 55 and nothing below.** In Forever the Spellstones are weapon enchants 8059, 8060 and 8061. Items 5522, 13602 and 13603 (made at 36, 48 and 60) apply them through 1237152, 1237156 and 1237159. Their spells 1237162, 1237164 and 1237165 are aura 13 Shadow damage 14/17/21 and aura 65 cast speed 1/2/2%. Now modelled as written. Test `TestSpellstoneIsShadowDamageAndCastSpeed` failed first (no Shadow damage or cast speed, +1% crit at 59 and 60) and passes. No site preset, bis run or /sim panel offers the stone, so nothing published moves. | SpellItemEnchantment, ItemXItemEffect, ItemEffect, SpellEffect rows; `talents.go:179` | CONFIRMED-BUG (client tables; no published number) | E 4a9938a16 |

## Fresh re-audit

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Cast spells against spellconst/warlock.json 70707: Shadow Bolt r9 251 + 1.8 a level, 0.857, 3 s, 370; Incinerate r3 217 + 1.4, 0.714, 2.5 s, 325, +25% on Immolate; Immolate r7 146 + 2.3 at 0.20 then 52 a tick for 5 at 0.13; Corruption r6 57 a tick for 6, 0.20; Bane of Agony r6 46 a tick for 12, 0.133; Bane of Doom 1742 at 4.0 after 60 s; Siphon Life r4 41 a tick for 10, 0.05; Conflagrate r6 282 + 1.3, 0.429, 10 s, 255 (ranks 1 and 2 are Forever's 1293817/1293818 at 25 and 32); Shadowburn r6 266 + 1.8, 0.429; Soul Fire r2 431 + 1.9, 1.0, 6 s; Life Tap r6 420 + 1 a level + Spirit; Lash of Pain r6 50, 0.429; Firebolt r7 44 + 0.6, 0.571. The engine's tables are generated from the same rows, and `spellconst_damage_test.go` and `spellconst_tables_test.go` hold them. | spellconst rows; fork tests pass | Aligned | - |
| F2 | Ranks: the rotations cast Shadow Bolt r9 and Immolate r7 at 60. r10 (25307) and r8 (25309) are book ranks, gated with their content. Corruption r7 (25311) is also a book rank. Per-band ranks come from the band rewrite (the ranks lane's). | spellranks.json | Aligned | - |
| F3 | Published builds decoded against the trees: Affliction raid takes Malevolence 5, Pandemic 3, Shadow Mastery 5, Siphon Life, Improved Shadow Bolt 5, Bane 5, Ruin 4, Unholy Power 5; no Nightfall and no Wrack. The bare 40 to 60 builds take Nightfall 2 and Wrack 1. The APL has no Wrack line (a 6 s channel at 36 a second; round 3 kept it out), and the Wrack point is a path point. Destruction raid: Bane 5, Ruin 5, Aftermath 4, Agonizing Flames 3, Conflagrate, Fire and Brimstone 3, Shadow and Flame 5, Incinerate; Bane of Havoc is taken and never cast single target, rightly. Demonology raid: Master Demonologist 5, Soul Link, Demonic Knowledge 3, Demonic Pact, Decimation 2, Ruin 5. | bis json talents; talents_auto_gen.go order | Aligned | - |
| F4 | Talents with no damage effect, checked: Demonic Aegis (aura 108 on Demon Skin/Armor, defensive), Fel Vitality (+15% mana for the warlock and demons, `talents.go:1167`), Suppression (aura 54 and 55 hit 5, threat -20), Improved Life Tap (x1.2), Destructive Reach and Grim Reach (range), Molten Skin (damage taken). Demon Armor (aura 674 armor 570, Shadow resistance 15, regen) is defensive; demons take 35% of the armor. | client rows | Aligned | - |
| F5 | Gnome Eureka! for warlocks (1259821): its damage mask names Conflagrate, Death Coil, Immolate, Rain of Fire, Searing Pain, Shadow Bolt, Shadowburn and Soul Fire, and its periodic mask the drains and Wrack. It does not name Incinerate, so Destruction's filler rightly gets nothing (`core/eureka_auto_gen.go` has no 1293813). Troll Berserking is unchanged since round 9. | 1259821 masks | Aligned | - |
| F6 | Items on the published sets: Frozen Heart of the Mountain (Chthonic Power, round 9), Kindling Stave (Ignition, round 8), Rune of the Guard Captain (Authority, equip), Robe of the Void (use: Heal Pet, no damage), Torch of Light (wand stats only). No set bonus is in a published set. | ItemXItemEffect, ItemEffect rows | Aligned | - |
| F7 | Rotations by band (bare 20 to 60, horde): Affliction casts Shadow Bolt, Corruption, Agony, Immolate and from 40 Siphon Life, with Amplify Curse from 30 and Shadow Trance from 40. Demonology and Destruction cast what their builds take. Life Tap covers mana: out-of-mana time is 0.0 s at 30 and 0.4 s at 60 raid. | spec-breakdown | Aligned | - |

## Before/after

10000 iterations. Before: the shared baseline (origin/main bb619a188 on 4d3129cae). After: wt/r10-warlock on lane/r10-warlock 4a9938a16. Seed 1 absolute, then % at seed 1 / 50001. All of it is P1 to P3 (Spellstone moves nothing published).

| Spec, band, preset | Horde s1 | Horde % (1 / 50001) | Alliance s1 | Alliance % (1 / 50001) |
|---|---|---|---|---|
| Affliction 20 bare | 53.7 to 53.4 | -0.56 / -0.56 | 53.9 to 53.6 | -0.56 / -0.56 |
| Affliction 30 bare | 87.1 to 86.6 | -0.57 / -0.57 | 87.2 to 86.7 | -0.57 / -0.57 |
| Affliction 40 bare | 132.9 to 132.2 | -0.53 / -0.53 | 133.1 to 132.5 | -0.45 / -0.45 |
| Affliction 50 bare | 193.3 to 192.6 | -0.36 / -0.41 | 194.8 to 194.1 | -0.36 / -0.41 |
| Affliction 60 bare | 286.9 to 285.6 | -0.45 / -0.49 | 286.3 to 285.0 | -0.45 / -0.45 |
| Affliction 60 raid | 573.2 to 571.0 | -0.38 / -0.38 | 571.5 to 569.3 | -0.38 / -0.38 |
| Demonology 20 bare | 47.7 to 47.4 | -0.63 / -0.63 | 48.9 to 48.6 | -0.61 / -0.61 |
| Demonology 30 bare | 79.5 to 79.0 | -0.63 / -0.63 | 79.4 to 78.9 | -0.63 / -0.63 |
| Demonology 40 bare | 140.7 to 140.1 | -0.43 / -0.50 | 140.7 to 140.0 | -0.50 / -0.50 |
| Demonology 50 bare | 206.7 to 206.0 | -0.34 / -0.34 | 206.1 to 205.4 | -0.34 / -0.34 |
| Demonology 60 bare | 323.4 to 322.2 | -0.37 / -0.40 | 322.8 to 321.5 | -0.40 / -0.37 |
| Demonology 60 raid | 642.6 to 640.4 | -0.34 / -0.34 | 637.8 to 635.6 | -0.34 / -0.34 |
| Destruction 20 bare | 52.3 to 52.0 | -0.57 / -0.57 | 52.5 to 52.2 | -0.57 / -0.57 |
| Destruction 30 bare | 88.2 to 87.7 | -0.57 / -0.57 | 88.6 to 88.1 | -0.56 / -0.56 |
| Destruction 40 bare | 138.7 to 138.0 | -0.50 / -0.43 | 139.4 to 138.8 | -0.43 / -0.43 |
| Destruction 50 bare | 234.6 to 233.9 | -0.30 / -0.30 | 235.6 to 234.8 | -0.34 / -0.30 |
| Destruction 60 bare | 351.9 to 350.7 | -0.34 / -0.34 | 355.1 to 353.9 | -0.34 / -0.34 |
| Destruction 60 raid | 638.9 to 636.9 | -0.31 / -0.31 | 636.6 to 634.6 | -0.31 / -0.31 |

Where it comes from (60 raid, horde, seed 1): Succubus melee crit goes from 2.8% of swings to 0, and Lash of Pain from 3.1% to 5.8% per attempt (Affliction 7.3% to 10.1% with Malevolence).

The fork suite `GOWORK=off go test -tags=with_db ./sim/warlock/... ./sim/core/...` passes; no fork golden moved. Site: `sim/cmd/spec-breakdown` passes. `sim/request` TestRotationLadder's three warlock goldens move (the demon's damage); not regenerated here (the integrator's job at the pin).

## For other lanes

- Integrator: E de87d2393 and E 4a9938a16 touch only `sim/warlock`. No spec outside the warlock group moves. Re-run the three warlock ladder goldens at the pin.
- Crit lane: P4 (does a demon take the warlock's every-table crit, gear only or buffs too?). Core crit plumbing is unchanged by this lane.
- Platform / guides: the Demonology guide's demon paragraph ("It does not yet take your spell crit") stays true. It could add that demons rarely crit in melee; optional.

## Still open

- P4, the warlock's crit reaching demons (single source); S1, the level-60 demon sheets; and the demon crit values at 60 (measured at levels 12 to 35). A launch log with a Succubus settles all three.
- From earlier rounds, unchanged: Ignition's one-a-minute rate, Improved Corruption with Shadow Mastery's periodic half, shadow wands and Shadow multipliers (#419), Siphon Life tick crit (held for the owner), Demonic Brand (no build takes it).
