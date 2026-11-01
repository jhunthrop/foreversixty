# Sim audit round 3: casters (Priest and Mage), 2026-11-01

Specs: priest-shadow, priest-discipline, priest-holy, mage-arcane, mage-fire, mage-frost. Warlocks are another lane.

## Method

- **Engine.** Lane branch `lane/r3-casters` from the pin `c8d29ea33`. The "before" figures run a second worktree at the pin.
- **Client data.** Client 1.60.1.70707 raw tables (round 2's copies, plus SpellCooldowns and SpellCategories for 70374 to 70707 and Era 1.15.9.70576 from `wago.tools/db2/<table>/csv`). Server text from `nether.wowhead.com/forever/tooltip/spell/<id>`, read 2026-11-01.
- **Logs.** foreverlogs.gg encounter timelines for 24 public beta reports (2740 to 2811, 8 to 23 October 2026), 1,103 encounters, cached in the lane scratch directory. Report 2807 is a TBC log (Illidan, Vampiric Touch) and is left out.
- **Runs.** `sim/cmd/spec-breakdown`, horde, raid preset, band 60, 10,000 iterations at seeds 1 and 50001; healers with `-heal`. Talent searches: `sim/cmd/talent-search -preset raid -faction horde -keep <raid-keeps.json>`.
- **Tests.** Each confirmed bug got a failing test first, seen failing, then passing.
- **A caution on tooltips.** Wowhead Forever renders a spell's amounts at the spell's maximum level, not at 60. Power Word: Shield 10901 reads 950 = 928 + 4.3 × 5 (MaxLevel 65); Ice Lance 1240047 reads 140-164, which is its level-64 roll. Compare tooltip amounts with the client table at MaxLevel.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | Pre-pull landing hold (fork c880120ea) on casters. | It holds only a pre-pull cast at an enemy with direct or spell damage. The caster pre-pulls are Lightwell (heal), Shadowform and Ice Barrier (self buffs), so nothing moves. Fire and Arcane have no pre-pull cast. | Aligned | - |
| X2 | Blood Fury off the GCD for orc mages (fork f44ba933d). | Fire and Arcane horde wear orc. Client 20572 is 10% attack power, ranged attack power and spell power (aura 166, 167, 317), 15 s, 2 min, StartRecoveryTime 0. `core/racial_blood_fury.go` multiplies SpellPower, SpellDamage and every school power by 1.10. Cast twice a fight, 16.7% uptime. | Aligned | - |
| X3 | Judgement of the Crusader in the raid preset (hybrids S bc6a1dce1). | `core/debuffs.go` JudgementOfTheCrusaderAura adds 161 to Holy damage taken only. No caster list casts a Holy damage spell: Shadow casts 5019, 10894, 10947, 14751, 15473, 18807, 19280 and 1309636; the healer lists cast heals only. Smite and Holy Fire are unused, so nothing moves. | Aligned | - |
| X4 | Presence of Mind and Combustion share cooldown category 1151 (180 s), in Era and in 70707. | The engine gives them separate timers (`mage/talents.go:270`, `:416`). No build takes both: Arcane casts Presence of Mind and has no Combustion, Fire has Combustion and no Presence of Mind. | Aligned (no build reaches it) | - |
| X5 | Heal profile coverage band (`data/curated/heal-profile.json`, rule: strongest raid set about 80% of incoming, weakest about 60, bare sets under half, and mana a live constraint). | Incoming damage: tank 1725 / 3.3 s plus 3 × 450 / 5 s = 792.7 a second. Raid sets, lane engine, seed 1: Holy 628.6 (79.3%), Discipline 540.6 (68.2%), Resto druid 721.5 (91.0%), Resto shaman 610.7 (77.0%), Holy paladin 531.3 (67.0%). Bare sets (Holy, Discipline, druid, shaman, paladin): 37%, 36%, 48.5%, 28%, 20%, all under half. Raid mana runs dry for Discipline (136 s), Holy (207 s) and Holy paladin (249 s). After round 2, Holy covered 96%, out of band. Removing Shadowfiend (P1) brings it back to 79%. The druid at 91% is now the only set out of band. | OPEN (the druid is the hybrids lane's; re-measure after its Omen of Clarity change, then retune if it stays above 85%) | - |

## Priest: Shadow

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Shadowfiend (401977) is not learnable, but round 2 put it in the Shadow list below 30% mana, and both healing lists cast it below half mana (round 2 row C2). | **Logs:** in 1,103 beta encounters (foreverlogs.gg, 8 to 23 October, builds 70291 to 70532, which already carry SkillLineAbility row 49886), about 25 priests cast Binding Heal 401937, Dark Sacrifice 1277325, Inner Focus, Penance, Restore Mana and Berserking. None cast Shadowfiend or got a Mana Leech energize. The only Shadowfiend events come from TBC report 2807 (34433). **Client:** row 49886 has SpellLevels BaseLevel 0. Under the trainables rule (memory `forever-trainables-source`), a normal trainer spell has a learn level above 0. **Guide:** Warcraft Tavern's Forever Shadow guide says Shadow has no self-power cooldowns. Test: `data/tests/test_apl.py` test_no_rotation_casts_a_spell_a_player_cannot_learn (failed, then passed); the fork's TestRotationsCastTheManaAbilitiesBehindTheirGates now requires 0 casts. | CONFIRMED-BUG (the logs and the guide are independent sources) | S 09848d599, E 72aefbb89 |
| P2 | Shadowfiend pet damage. | Moot after P1. | - | - |
| P3 | Round-2 fixes re-checked. **Darkness:** one aura 79 school multiplier, the same shape as Shadowform's effect 1. The two multiply, and Touch of the Grave (Shadow) takes both. **Shadow Word: Death:** 18 s cooldown (SpellCooldowns 100966). 8.97 casts a fight, against a 10-cast ceiling. Coefficient 0.429 (client). **Backlash:** health only. **Shadow Weaving uptime:** 99.7%. | Client 70707 SpellEffect 15259, 15473 and 1309636; spec-breakdown. | Aligned | - |
| P4 | Coefficients and amounts for every spell the list casts. | Client EffectBonusCoefficient, matched by `priest/constants_auto_gen.go`: Mind Flay 18807 0.167 a tick (3 × 130), Mind Blast 0.429, Shadow Word: Pain 0.2 a tick, Devouring Plague 0.1 a tick. Mind Blast is in category 19 (8 s), less Improved Mind Blast 5/5. | Aligned | - |
| P5 | Raid talent search, with `-keep`. | Best variant is +0.2% (±0.2%): under 1%. | Hold | - |

## Priest: Discipline and Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | Penance and Power Word: Shield coefficients (round 2 row H6). | The client gives them. SpellEffect EffectBonusCoefficient is 0.19 on every Penance pulse spell: heal 402289, 1240723, 1240724 and 1316991; damage 402284, 1240727, 1240730 and 1316993. It is 0.10 on every Power Word: Shield rank, 17 through 10901. `priest/heal_ranks.go:140` and `:193` use exactly those numbers. The "unconfirmed (vanilla convention)" marks sit in `priest/constants_auto_gen.go:2399` and `:2468`, tables the healer code never reads. That file's Power Word: Shield rank 10 is NPC spell 27607 at 2.0, which is dead data. The Wowhead tooltips show no spell-power share. | Aligned (settled by the client) | - |
| H2 | Shadowfiend in both healing lists. | See P1. Effective HPS: Holy -17.3%, Discipline -7.8%. The lists ran on its 4,360 mana a fight. | CONFIRMED-BUG | S 09848d599, E 72aefbb89 |
| H3 | Prayer of Mending charge on damage taken (round 2). | 13.7 casts, 41,075 healed: about 4.4 of 5 charges spent a cast, each at 413 + 0.429 × 608. Dark Sacrifice: Undead-only, cast by the Undead Discipline healer. | Aligned | - |
| H4 | Power Infusion on yourself (round 2 row H8). | `priest/power_infusion.go` registers it only on a picked ally. Neither build takes it. Settling it needs a raid policy for who receives Power Infusion, then a `sim/cmd/heal-builds` run with a self-target. | OPEN | - |

## Mage: Arcane

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Arcane Blast, Missile Barrage, Arcane Power and Presence of Mind against 70707 text. | Tooltips 400573, 400589, 12042 and 12043 match `mage/arcane_blast.go` and `talents.go`. | Aligned | - |
| A2 | Raid talent search, with `-keep`. | Best variant is deep Arcane (2 Frost Channeling to Incineration): +0.6% (±0.3%). Under 1%. | Hold | - |

## Mage: Fire

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Ignite skipping target modifiers (round 2). | Ticks carry SpellFlagIgnoreTargetModifiers and take 40% of a crit that already took Curse of the Elements, Improved Scorch and Fire Power. That matches Blizzard's beta note. | Aligned | - |
| F2 | Fire Blast now shares a timer with Ice Lance (E1 below). | The Fire build has no Ice Lance, so nothing moves (710.9 / 712.5 before and after). | Aligned | - |
| F3 | Raid talent search, with `-keep`. | No variant beats the guide beyond error. | Hold | - |

## Mage: Frost

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Every Ice Lance rank has an 8 s cooldown, shared with Fire Blast (round 2 row R4). The engine gave the trained ranks none. | **Client:** 1.60.1.70685 put rank 1 (1312002) into SpellCategories Category 19, with SpellCooldowns CategoryRecoveryTime 8000. That is the same category and recovery as every Fire Blast rank (2136 to 10199, 1318432). In 70374, 70450, 70532 and 70658, rank 1 had no cooldown. The trained ranks 400640 and 1240044-1240047 have Category 0 and no recovery. **Server text:** Wowhead Forever shows "8 sec cooldown" on every rank. That is the same data that gives the 300% Frozen bonus over the shipped base points of 200, which the engine already follows. A talent rank on a cooldown that its own trainer upgrades remove would be incoherent. **Logs:** beta casts of 400640 and 1312002 every 1.4 to 1.5 s (reports 2785 and 2786) predate 70685, so they do not speak to the live build. This is one source (the client change plus the tooltip), and in-game confirmation is still wanted. Test: `mage/ice_lance_cooldown_test.go` TestEveryIceLanceRankHasFireBlastsEightSecondCategoryCooldown (failed on ranks 2-6 and the timer, then passed). The old TestIceLanceCooldownFollowsTheClientRank asserted the opposite and was removed. | CONFIRMED-BUG (single source, see evidence) | E cf8231b8b |
| R2 | The Frost list after R1. | The Ice Lance line still fires on a Fingers of Frost charge. The second charge now goes to Frostbolt, which spends it for Shatter crit. Ice Lance casts drop from 7.08 to 5.18; Frostbolt crit rises from 21.6% to 22.7%. | Aligned | - |
| R3 | Raid talent search after R1, with `-keep`. | No variant beats the guide beyond error (best +0.2%). | Hold | - |

## Before/after (relative; lane engine plus site branch, against the pin)

| Spec | Preset | Before (seed 1 / 50001) | After | Change | Cause |
|---|---|---|---|---|---|
| Shadow, horde | raid | 669.2 / 669.2 | 669.2 / 669.2 | 0.0% | P1 (Shadowfiend never fired; 0 s out of mana) |
| Shadow, horde | bare | 409.7 / 409.4 | 360.2 / 360.3 | -12.0% | P1 |
| Arcane, horde | raid | 665.7 / 666.4 | 665.7 / 666.4 | 0.0% | - |
| Fire, horde | raid | 710.9 / 712.5 | 710.9 / 712.5 | 0.0% | - |
| Frost, horde | raid | 562.0 / 561.7 | 557.3 / 556.7 | -0.8% | R1 |
| Holy, horde (troll), effective HPS | raid | 760.1 / 760.7 | 628.6 / 628.6 | -17.3% | H2 |
| Discipline, horde (undead), effective HPS | raid | 586.2 / 586.3 | 540.6 / 540.6 | -7.8% | H2 |

## Still open

- **Ice Lance's cooldown on the trained ranks.** Settled from the client change and Wowhead's hotfix data. One in-game check settles it fully: two Ice Lances inside 8 s at rank 2 or higher, or a Fire Blast followed by an Ice Lance.
- **Power Infusion for healers (H4).**
- **Heal profile band (X5).** The Resto druid covers 91%.
- **Ladder goldens.** The Shadow, Holy and Discipline rotations and Frost's Ice Lance moved. TestRotationLadder goldens are left for the integrator, per the brief.
- **Shadowfiend at launch.** If a launch-day trainer or spellbook shows it, revert S 09848d599 and E 72aefbb89, and model the pet's damage.

## Sources

- foreverlogs.gg public reports 2740 to 2811, encounter timelines: https://foreverlogs.gg/ (read 2026-11-01)
- Wowhead Forever tooltips (read 2026-11-01): https://nether.wowhead.com/forever/tooltip/spell/ with the ids 1316995, 1316991, 10901, 401977, 34433, 400640, 1240044-1240047, 1312002, 10060, 15473, 18807, 1309636, 20572, 400625, 400624, 11129, 400589, 400669, 400647, 12042, 12043, 400573, 11124, 29074 and 11160
- wago.tools SpellCooldowns and SpellCategories for 1.60.1.70374, 70450, 70532, 70658 and 70707, and 1.15.9.70576
- Warcraft Tavern, Shadow Priest Guide (Forever): https://www.warcrafttavern.com/forever/guides/shadow-priest/
