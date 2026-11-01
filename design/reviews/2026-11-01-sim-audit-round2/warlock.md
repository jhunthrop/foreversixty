# Sim-accuracy round 2: warlock (2026-11-01)

Specs: Affliction, Demonology, Destruction. Client 1.60.1.70707. Pinned engine 17d31ecfa; lane engine `lane/r2-warlock` (fork). Site branch `wt/r2-warlock`.

Method: the client tables come from wago.tools for build 1.60.1.70707: SpellEffect, SpellMisc, Spell, SpellClassOptions, SkillLineAbility, SpellCooldowns and SpellLevels. Each talent's class mask was resolved to the warlock spells it covers. Wowhead Forever tooltips were checked at `https://nether.wowhead.com/forever/tooltip/spell/<id>`. Engine code was read in `sim/warlock`. Runs used `sim/cmd/spec-breakdown` (horde raid preset, band 60, troll, 10000 iterations at seeds 1 and 50001) and `sim/cmd/talent-search` (raid preset, `-keep` from raid-keeps.json, lane engine). The talent-search report header prints the pinned sha, but GOWORK pointed it at the lane engine.

## Real Forever logs

- **Warcraft Logs.** No Forever partition was found. The logs are on foreverlogs.gg, a community site: 119 public reports, all beta dungeon runs at levels 13 to 30. They include 29 report-warlock pairs from 20 warlocks. The raw downloads are in the scratchpad (`r2/../fl/`). Nothing is at 60, and Wrack (level 40) never appears, so these logs only check mechanics.
- **Our API.** Round one found one public report with fights, a beta trash log. No level-60 warlock data exists before launch.

| Check | Logs (levels 13 to 30) | Sim | Verdict |
|---|---|---|---|
| Non-channelled DoT ticks crit | Corruption 6.5% of 4,076 ticks, Agony 6.2% of 6,174, Immolate 5.4% of 1,391 | All roll crit | Aligned |
| Siphon Life ticks crit | **0 crits in 772 ticks** | Rolls crit (owner rule) | OPEN for the owner (below) |
| Drain Life (a channel) crits | 2.1% of 1,546 | Client flag 0x200 set, so it crits | Aligned |
| Life Tap scales with Spirit, not spell power | Mana per cast moves with gear: 214 to 273 near level 30 | (amount + Spirit) × 1.2 | Aligned |
| Pet crit | Imp Firebolt 2.4%, pet melee 1.4%, Lash of Pain 7.6% | Succubus melee 10.2% at 60 | OPEN: the logs are low level against higher-level mobs |
| Pet share of damage | Usually 6% to 20% | Affliction 23%, Demonology and Destruction lower | No level-60 basis |

## Affliction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **Wrack is a 6 s channel; the engine cast it as an instant DoT.** The warlock kept casting Shadow Bolts over Wrack's ticks, which made Wrack the top damage source at 25% of the spec's damage. | SpellMisc 1316697 Attributes_1 = 0x4004, the same bits as Drain Life 11700 and Mind Flay 15407. Wowhead Forever: "Channeled (6 sec cast)". Engine `wrack.go` had no `SpellFlagChanneled`. | CONFIRMED-BUG | E a38dedc7e |
| A2 | Wrack's vulnerability raised all Shadow damage on the target from every source, Wrack itself and Shadow Bolt included. | SpellEffect 1316697 effect 1: aura 271 (damage taken from the caster) on class mask 0x402, which is Corruption (0x2) and Bane of Agony (0x400) only. It is now on the warlock's attack table and ends with the channel. | CONFIRMED-BUG | E a38dedc7e |
| A3 | Nightfall rolled only on Corruption and Drain Life. | Talent text 18094: "Corruption, Drain Soul, Drain Life, and Wrack". The 2026-10-07 Blizzard deep dive lists Drain Soul and Wrack as new triggers. The sim gave 2.26 procs per fight; 4% of all eligible ticks is about 8. `talents.go:288`. | CONFIRMED-BUG | E a38dedc7e |
| A4 | Shadow Mastery missed Wrack, Bane of Doom and Shadowburn, and a test asserted that Bane of Doom was outside its masks. | 18271 effect 0 mask (0x80093,0,0) and effect 1 mask (0x440A,0x40003,0). Bane of Doom 603 is word 1 0x2, Wrack is word 1 0x40000, Shadowburn is word 0 0x80. `talents.go:337`. | CONFIRMED-BUG | E a38dedc7e |
| A5 | With Wrack a channel, the APL chained it and nothing else: 461 DPS. As a filler, Wrack measured 11% below Shadow Bolt (539 against 605). Its line is gone, and Wrack leaves the raid keep list. | Runs `aff-wrackfiller` and `aff-nowrack`. Wrack is 0.143 coefficient a second; Shadow Bolt is 0.857 over a 2.5 s cast. | CONFIRMED-BUG (APL) | S c18ae5887 |
| A6 | **Held talent change, round one.** The raid search on the lane engine (keeps: Amplify Curse, Siphon Life, Suppression, Demonic Aegis) wins with `2552010013500105/00052/0550004` (30/7/14): Malevolence 5/5 and Ruin 4/5 in place of Improved Drains, Nightfall, Soul Siphon and Wrack. It keeps Grim Reach. Checked at 10000 iterations: **+5.1%** (605.1 to 636.2 at seed 1, 604.9 to 635.6 at 50001). Every adoption-rule condition holds: the gain beats error and 1%, every unmodeled guide talent is kept, Affliction holds the most points, and nothing the APL casts is dropped. Adopted. | `r2/warlock/ts/warlock-affliction-raid.md` | CONFIRMED (build) | S b1b962400 |
| A7 | Life Tap returns about 832 mana a cast. | The Forever text is `($m1+$SPI*1)*(1+ILT%)`, with effect coefficient 0. That is 424 at level 60 (RealPointsPerLevel 1, SpellLevel 56, MaxLevel 66) plus 269 Spirit, times 1.2. Wowhead shows 430 because it renders at MaxLevel. It does **not** scale with spell power. | Aligned | - |
| A8 | Shadow Bolt misses 10.4% of the time. | Suppression (18174) is aura 55 at 5 with class mask 0, so it covers every spell, not only Affliction. 17% base miss against a level-63 boss, less 5% from Suppression, less about 1.6% from gear. | Aligned | - |
| A9 | The raid preset has Curse of the Elements and no Curse of Shadow or Shadow Weaving. | Warlocks do not learn Curse of Shadow (402791 has no SkillLineAbility row). Forever's Curse of the Elements 1311680 covers all magic (mask 126): -75 resistance, +10%. Shadow Weaving 15258 is aura 270 "from the caster", so a priest's stacks do not help a warlock. Improved Shadow Bolt 17794 is also caster-only, 12 s, no charges. The preset is right. | Aligned | - |
| A10 | Improved Shadow Bolt is modelled: 4% a rank, 12 s, refreshed by Shadow Bolt crits. The engine applies it to every source's Shadow damage. Only the pet's Lash of Pain gains wrongly, under 0.1%. | `core/debuffs.go:322` | Aligned (minor) | - |
| A11 | Pandemic covers Corruption, both Banes, Drain Life, Drain Soul, Siphon Life and Wrack, and doubles the crit bonus (2.0×). Ruin raises Immolate's ticks the same way. | 427712 mask matches `pandemicAffectedSpellCodes`. `critMultiplierAs` gives 1 + 0.5 × CritDamageBonus. | Aligned | - |
| A12 | Improved Drains: the hotfixed text (talents json and Wowhead) names Drain Mana, but the SpellEffect 403511 mask still carries Wrack. The raw Spell.csv text also still says Wrack. The 70707 lane followed the text. The raid build no longer takes the talent. | wago Spell.csv 403511; SpellEffect mask (0x4008, 0x40000) | OPEN (moot for the raid build) | - |
| A13 | Drain Soul has no execute bonus. | A player reports it "is gone" (single source, Blizzard forum 2026-10-22). The engine has none. | Aligned (single source) | - |
| A14 | **Wrack re-check before ship** (engine 61ef5d396, which matches fork master c8d29ea33 for warlocks; horde raid; 10000 iterations; seeds 1 / 50001). On the old Wrack-branch build: shipped rotation (no Wrack) 605.3 / 604.9. (a) Wrack whenever its vulnerability is down while Corruption and Agony run: 539.5 / 539.8. The vulnerability lasts only as long as the channel, so this chain-casts Wrack. (b) Shadow Bolt on Shadow Trance, then Wrack as filler: 550.4 / 550.3. (c) Wrack as pure filler: 539.5 / 539.8. (d) wowtbc's Forever rotation verbatim (Shadow Bolt on Nightfall > Agony > Corruption > Siphon Life > Immolate > Shadow Bolt to keep Improved Shadow Bolt up > Wrack): 590.2 / 590.7, -2.4%. The adopted Malevolence build is 635.6 / 636.2. No variant beats the shipped rotation, so the Wrack-branch build stays behind. Even if Improved Drains' mask (+20% to Wrack, A12) were real, (b) would reach about 576 and (d) about 598. Wrack's damage was re-verified: SpellEffect 1316697 effect 0 is aura 3, base 36, coefficient 0.143 a tick, a 1 s period over 6 s (6 ticks). Wowhead: "36 Shadow damage every 1 sec", "Channeled (6 sec cast)". The engine matches (`wrack.go`). Over the channel that is a 0.858 coefficient, 0.143 a second, against Shadow Bolt's 0.343 a second and Drain Life's 0.1. | `r2/warlock/wrack-variants.txt`, rotations `r2/warlock/rot/w-*.json`; wowtbc.gg Forever Affliction guide; Warcraft Tavern's Forever Affliction guide covers only the level-30 beta and never mentions Wrack | Aligned (no change) | - |

## Demonology

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | **Decimation's execute gate ignored encounter dummies.** The engine treated any target without a health bar, which is every encounter, as below 35% health. Soul Fire's 10 s buff was therefore up 83% of the fight, and Shadow Bolt took the 6% bonus for the whole fight. | Talent text 440870: "on an enemy below 35% health". Such a target now qualifies only in the encounter's execute phase (`soul_fire.go`, `decimationTargetEligible`). | CONFIRMED-BUG | E f3ae3fad9 |
| D2 | The Decimation buff's cast time was cut by 40% of the base 6 s, giving 1.6 s. | 440873 effect 0 is aura 108, misc 10, a percent modifier, so it scales after Bane's flat -2 s: 2.4 s. | CONFIRMED-BUG | E f3ae3fad9 |
| D3 | The buff took 40% off Soul Fire's mana cost. | 440873 effect 1 is aura 256, no reagent use: that is the free Soul Shard. The mana cost stands. | CONFIRMED-BUG | E f3ae3fad9 |
| D4 | The APL cast Soul Fire on cooldown. Outside the buff, Soul Fire is a 4 s cast that costs a shard and deals less per second of casting than Shadow Bolt. Soul Fire now casts only under Decimation's buff: 702.7, against 702.1 with no Soul Fire at all and 692.7 for Soul Fire in execute only. | Runs `demo-sfdecim`, `demo-nosf`, `demo-sfexec` | CONFIRMED-BUG (APL) | S c18ae5887, 83c08beb8 |
| D5 | Demonic Pact flow: Imp sacrifice (+15% Shadow) holds for 100% of the fight with the Succubus out. Master Demonologist's +10% Shadow is active: removing it lowers Shadow Bolt's average hit by 9.9%. The aura's uptime metric reads -25% and 66.7%, because prepull time counts negative; this is a display artefact. | Runs with `-talents ...1301...` | Aligned | - |
| D6 | The raid search finds no winner beyond error. The best gain is +2.1 ± 2.2, which shows Decimation's two points are worth keeping. | `r2/warlock/ts/warlock-demonology-raid.md` | Aligned | - |

## Destruction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | Incinerate gets +25% with Immolate up, a 2.5 s base cast and 2.0 s with Bane, coefficient 0.714. Conflagrate keeps Immolate (Shadow and Flame 100%). Shadowburn's shard refund is 100%, so it is cast on cooldown. Shadow and Flame keeps +10% Shadow and +10% Fire up 94% and 86% of the fight. | Client 1293813, 426316, 18932. `incinerate.go:50` | Aligned | - |
| S2 | Agonizing Flames (+10% direct damage on Destruction spells, +10% periodic on Immolate) and Ruin (crit bonus on the Destruction mask, Immolate ticks included) match their masks. | 17927, 17959 masks | Aligned | - |
| S3 | Raid search: `2550000002/0/2053345103101351` (Pandemic 2 and Cataclysm 3 for Unholy Power 5). With buffed demons it gains +0.9%, below the bar. With demons refusing player buffs (X3) it gains **+1.9%** at 10000 iterations (652.4 to 665.1 at seed 1, 653.0 to 665.6 at 50001) and passes the adoption rule. Adopted, but it depends on fork dfbb63ae3: drop this commit if that one is dropped. | `r2/warlock/ts3/warlock-destruction-raid.md` | OPEN (conditional on X3, held with it) | held (lane S 5d013a9c0, not merged) |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Blood Fury triggered a 1.5 s global cooldown.** Its 2 min cooldown, 15 s duration and +10% spell power (aura 317), attack power and ranged attack power were already right. The published horde warlocks are troll, so warlock DPS does not move. Orc goldens move: Enhancement, Fury/Arms, Protection warrior. | SpellCooldowns 20572: StartRecoveryTime 0, as for Berserking 20554. `core/racial_blood_fury.go` | CONFIRMED-BUG | E f44ba933d, c20b3e345 |
| X2 | Pets inherit nothing from the warlock (`pet.go`, `makeStatInheritance` returns empty stats), although the deep dive says pets scale from their summoner. The client has Warlock Pet Scaling 416189, with channels for spell power, attack power, hit (54 and 55), crit (52 and 57), haste (319 and 65), stamina and resistances, but every amount is computed on the server and no figure is published. Hunter pets use guide figures (10% AP, 100% crit). Lash of Pain misses 17% because the pet has no hit. | Client 416189; Blizzard deep dive 2026-10-07 | OPEN: needs level-60 logs (pet hit and crit against warlock gear) or a published coefficient | - |
| X3 | Demons took every player buff, including Windfury Totem extra attacks (2.8% of Affliction's DPS), Battle Shout, Blessings and Heart of the Lion. Hunter pets already refuse them (`hunter/pet.go:110`). The evidence for demons is indirect. Blizzard's hunter deep dive (2026-09-30) says pets that scale from their owner "can no longer receive player buffs that increase their stats". The warlock deep dive (2026-10-07) says only "Pets now scale from their summoner's stats". The client's Warlock Pet Scaling 416189 carries the same effect list as Hunter Pet Scaling 415429. Kargoz's Forever Demonology guide (kami-labs.fr) says demons scale "exactly like hunter pets". Demons now refuse player buffs; test `pet_buffs_test.go`. **Caveat:** the inheritance in X2 is still zero, so the demon now has neither the buffs nor its scaling. Pet share falls to about 9 to 11%, against "around 35%" in Kargoz's guide and 6 to 20% in the low-level logs. The integrator may prefer to land this together with an inheritance model. | Blizzard hunter and warlock deep dives; client 416189 and 415429; kami-labs.fr Kargoz Demonology leveling | OPEN (indirect evidence; held by the integrator) | held (lane E dfbb63ae3, not merged): refusing buffs while demons still inherit nothing leaves pet share at 9-11% against about 35% in guides and 6-20% in low-level logs, further from the truth than today. Lands together with a pet-inheritance model. |
| X4 | Siphon Life ticks never crit in 772 logged ticks (foreverlogs, levels 13 to 30), against 6% for the other DoTs. Siphon Life is aura 53 (periodic leech). The owner rule says every non-channelled DoT crits, so the engine is unchanged. | foreverlogs.gg ability breakdowns | OPEN for the owner | - |
| X5 | spec-breakdown left pet rows out, so a warlock table summed to about 80%. Pet actions are now listed under the pet's name. | `sim/cmd/spec-breakdown/main.go` | Tool fix | S 24d977725 |
| R | Round-one regressions on the 70707 engine. W1: Siphon Life upkeep is still cast (6.4 casts a fight). W2: Agony over Doom still holds. W3: DoT ticks crit at about 16%. X1: wand Shoot is still gated by the global cooldown (no Shoot casts with mana above 10%). D2: Pact flow holds. S1: Shadowburn refund holds. All still hold. | Baseline runs `r2/warlock/runs/base-*` | Aligned | - |

Haste weights are not re-measured here, because the nightly owns the weights. The new Affliction build and the Wrack channel change both alter them.

## Before/after (relative, horde raid, band 60, 10000 iterations)

| Spec | Step | Seed 1 | Seed 50001 |
|---|---|---|---|
| Affliction | Wrack channel, vulnerability scope, Nightfall, Shadow Mastery, APL drops Wrack | -15.9% | -15.9% |
| Affliction | plus the adopted raid build | **-11.6%** | **-11.7%** |
| Demonology | Decimation gate, cast time, cost, APL gates Soul Fire | **-4.8%** | **-5.0%** |
| Destruction | none of the fixes above reach its build | 0.0% | 0.0% |
| Affliction | demons refuse player buffs (X3) | -10.9% | -10.8% |
| Demonology | demons refuse player buffs (X3) | -10.0% | -10.1% |
| Destruction | demons refuse player buffs (X3) | -9.6% | -9.5% |
| Destruction | plus the adopted raid build | +1.9% | +1.9% |
| **All steps** | Affliction / Demonology / Destruction | **-21.2% / -14.3% / -7.8%** | **-21.2% / -14.5% / -7.7%** |

Absolute DPS at seed 1, for the integrator only: Affliction 719.5 to 567.0, Demonology 737.6 to 631.8, Destruction 721.4 to 665.1. With X3 in place, the Affliction raid re-search's best variant that keeps every unmodeled guide talent gains only +0.7%, so the A6 build stands, and Demonology still has no winner. The new order puts Destruction first, then Demonology, then Affliction. Two competing community sims (SixtySim, MythicSim) rank Demonology first.

## Still open

- Pet scaling from gear (X2): the largest open warlock question now that demons take no buffs. It decides pet damage, Unholy Power and Destruction's build.
- Siphon Life tick crits (X4), for the owner.
- Improved Drains, text against mask (A12).
- sim/request TestRotationLadder goldens are not regenerated, as the brief requires. The integrator regenerates them at the pin.
- The leveling `build:` of the Affliction guide still takes Wrack. That path was not searched.
