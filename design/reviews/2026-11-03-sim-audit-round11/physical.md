# Sim audit, round 11: physical group (2026-11-02)

This lane covers Arms, Fury, Protection, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee), on client 1.60.1.70707.

- **Baseline:** fork pin cfcb76e80 with site origin/main 8f047a62f (the shared `r11/sb-before`).
- **Branches:** fork fixes on `lane/r11-physical` (E), site changes and this report on `wt/r11-physical` (S).
- **Measurements:** `sim/cmd/spec-breakdown`, 10,000 iterations, seeds 1 and 50001, both factions, on the published sets.
- **Client rows:** the layered 70707 raw tables with the hotfix push (round 10's copy, fetched 2026-11-01 21:42Z). Wowhead Forever tooltips were read 2026-11-02; a tooltip and its client row count as one source.
- **New client build.** wago lists 1.60.1.70753 (2026-11-02 21:35). For the warrior, rogue and hunter families, its SpellEffect, SpellMisc, SpellCooldowns, SpellPower and SpellAuraOptions rows are identical to the layered 70707 tables. The only change is to SpellMisc, where 343 rows gain Attributes_14 0x40, the bit the hotfix push already layered in. So nothing moves for this group.
- **No new logs.** foreverlogs.gg reports stop at id 2811 ("Stocks as Beta Ends", 23 October), and api.foreversixty.gg/v1/reports/recent has nothing newer. No log can settle a level-40+ question this round.

Origin tags: `old` means the bug predates round 9. No bug here came from a round-9 or round-10 commit, or from new data.

## Assigned items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Marksmanship's held build stays held.** I re-ran the search at the current pin (cfcb76e80; `-keep` Trueshot Aura; horde raid; 3,000 finalist iterations). It finds round 9's winner again, 7/31/13 (`502-005145020150305-55000003`): +2 Improved Arcane Shot, +5 Deflection and +3 Surefooted, paid for by Endurance Training 3, Unleashed Fury 5, Ferocity 1 and Efficiency 1. It drops nothing the rotation casts. At 10,000 iterations it gains +0.90 / +0.87% alliance (618.5 to 624.1) and +0.93 / +0.87% horde (623.0 to 628.8). That is beyond error but under the 1% bar. | `scratchpad/r11/physical/ts-mm/hunter-marksmanship-raid.md` | HOLD (under 1%) | - | - |
| A2 | **Normalized rage: off-hand factor 0.5, crit x2.0.** There is no new evidence. The beta closed on 23 October, and no public log reaches a level where a warrior dual wields with Dual Wield Specialization. wowfdb's change log (read 2026-11-02) still lists the off-hand rage as an open question. | foreverlogs.gg report ids; https://wowfdb.com/classes/warrior/changes/ | OPEN (as round 10 H2) | - | - |
| A3 | **Lacerating Strikes' refresh** (a new Mongoose Bite replaces the bleed with no rollover). There is no new evidence. The talent sits at Survival tier 7 (40+), above every beta log. The 70707 rows are unchanged: 1310533 is a dummy (aura 4, 40%), and 1310536 is aura 3 every 3 s. | SpellEffect 1310533, 1310536 | OPEN (as round 10 H3) | - | - |
| A4 | **The one-a-minute proc rate.** No new evidence. | - | OPEN | - | - |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Execute converted rage at rank 5's 15 a point at every rank.** The client description of every rank is "converting each extra point of rage into $*10;F1 additional damage". Wowhead Forever renders it per rank as 3, 6, 9, 12 and 15 (5308, 20658, 20660, 20661, 20662), which matches Era. The flat bases (125, 200, 325, 450, 600) were already right. A warrior below 56 hit too hard in the execute phase. Moves: Arms bare 30 −6.7%, 40 −2.0%, 50 −0.7%; Fury bare 30 −5.9%, 40 −2.4%, 50 −1.4 to −1.6%. Level 60 does not move. Test `TestExecuteRageConversionFollowsTheRank` failed first. | `sim/warrior/execute.go:13` (`executeRageConversion = 15.0`, since f6054c4f1, 2026-09-18); https://nether.wowhead.com/forever/tooltip/spell/5308 (and 20658, 20660, 20661) | CONFIRMED-BUG (client text plus Era) | old (round 8's rank check covered the flat base only) | E 83cd1abf7 |
| W2 | Checked against the client rows and found aligned:<br>- Mortal Strike: normalized (effect 121), +85 to +160 by rank, 30 rage, 6 s.<br>- Overpower: +35 normalized, 77.7% crit (Improved Overpower 50).<br>- Slam: weapon damage +87, 1.5 s cast, 18 s cooldown, 15 s with Improved Slam.<br>- Rend: 21 a tick ×7 at rank 7, Improved Rend 35%.<br>- Hamstring: 45 flat.<br>- Deep Wounds: 60% over 12 s.<br>- Weaponmaster: axe and polearm 5% crit; mace and staff 3% a point armor ignore (15% at 5/5); sword 1% a point extra attack.<br>- Bloodthrill: 20%, main hand, Rend target.<br>- Also Recklessness, Berserker Stance, Blood Fury, Bloodrage, Anger Management, Two-Handed Weapon Specialization and Impale.<br>- Gear: Arcanite Champion's Strength of the Champion and Molten Heart's Molten Fury are modelled. Cloak of Flames (21142, aura 15) is a damage shield and does nothing for a DPS warrior who is not hit. | SpellEffect / SpellCooldowns / SpellPower rows; Forever tooltips 21553, 11585, 11605, 11574, 7373; `bd/warrior-arms-raid60.txt` | aligned | - | - |

## Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Whirlwind's off-hand strike was dealt as a main-hand special.** The client's off-hand strike is its own spell, 462891. It carries SpellMisc Attributes_3 0x01000000, the off-hand attack flag. It does normalized weapon damage (effect 121) to one target, has no cost and no stance requirement, and shares Whirlwind's family mask. The engine dealt it through 1680 with the main-hand special mask. So the off hand's hit rolled Windfury Totem (a main-hand-only proc) and any main-hand enchant, and never the off hand's own procs. It is now a separate off-hand special; Dual Wield Specialization and Furious Precision reach it through their own masks. Fury raid moves −0.37 to −0.50%. Bare 60 has no totem and moves −0.07%. Test `TestWhirlwindOffHandStrikeIsAnOffHandSpecial` failed first; TestP1DPSWarrior regenerated. | SpellMisc 462891 (Attributes_3 16777728), SpellEffect 462891, SpellClassOptions 462891; `sim/warrior/whirlwind.go`; `core/buffs.go:1641` (Windfury requires ProcMaskMeleeMH) | CONFIRMED-BUG (client row) | old (since the 1 October baseline change; round 7's a3b0a2611 touched the strike and kept the main-hand mask) | E 405bfeda0 |
| F2 | Checked and found aligned:<br>- Bloodthirst: 45% AP plus 30/…/48 by rank (23881, 23894).<br>- Whirlwind: both hands, 25 rage, 10 s.<br>- Heroic Strike queueing, with no off-hand bonus from the queue.<br>- Furious Precision: off-hand hit 4/7/10% (off-hand white miss 8.1% against main-hand 18.2%, both labelled correctly now).<br>- Dual Wield Specialization: +25% damage, +50% off-hand rage.<br>- Flurry 25%; Death Wish 20%, 3 min; Unbridled Wrath 1 rage, 12% a point; Cruelty.<br>- Troll Berserking: 10%, 15 s, 3 min. | rows above; `bd/warrior-fury-raid60.txt` | aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Rage from damage taken converted at the attacker's level.** The engine computes 2.5 × damage / c, and it took c at the attacker's level. A level-60 tank hit by a level-63 boss divided by c(63) = 247.3 instead of c(60) = 230.6, about 7% less rage. The formula uses the receiver's level in two independent sources. Kalgan's 1.x rage post, as reposted, gives "rage conversion at your level" with 230.6 at 60. cmangos-classic's `Player::RewardRage` computes c from `GetLevel()` of the player gaining the rage. Berserker Rage's spell-damage clause had the same mistake and is fixed too. Found by the **hybrids lane**. Test `TestRageFromDamageTakenUsesTheReceiversLevel` failed first. TestP1TankWarrior and TestFeralTank regenerated. | `sim/core/rage.go:175` (upstream a9175da8e, 2024-05-10, "depend on attacker's level"); https://raw.githubusercontent.com/cmangos/mangos-classic/master/src/game/Entities/Player.cpp `RewardRage` (read 2026-11-02); Kalgan reposts at https://www.tapatalk.com/groups/purgatory468/new-rage-generation-formulas-t2866.html and https://vanilla-wow-archive.fandom.com/wiki/Rage | CONFIRMED-BUG (two sources) | old (inherited from upstream; no round looked at it) | E 12979ebe9 |
| P2 | Checked and found aligned:<br>- Revenge: 109 to 133 plus 25% AP, Improved Revenge 60%.<br>- Shield Slam: 640 to 670 plus block value; Shield Block: 75%, 2 charges, 5 s.<br>- Thunder Clap: 103, DefenseType 1, so it rolls the spell table (16% miss).<br>- Sunder: 450 armor, 206 threat plus 5% AP.<br>- Defiance: 15%; Bastion: 10%; Focused Rage: −3.<br>- Shield Specialization and Master of Defense: 5 rage.<br>- Threat values without client rows stay open (carried). | Forever tooltips 11601, 23925, 2565, 11581; SpellEffect 11597 | aligned | - | - |

## Assassination, Combat, Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | No new finding. What I checked:<br>- Each published raid breakdown, spell by spell: Mutilate (both hands, 1241586 and 1241590; 75% plus flat; 60 energy; ranks 30/40/50/60), Eviscerate at 4 and 5 points, Backstab, Hemorrhage (145% with a dagger), Ghostly Strike, Rupture, Ambush, Premeditation, Cold Blood, Slice and Dice, Blade Flurry (2 casts, 16.4% uptime), Adrenaline Rush, and the poisons' tables.<br>- The published builds' talents against their client text: Malice, Ruthlessness, Relentless Strikes, Lethality, Vile Poisons, Improved Poisons, Seal Fate, Improved Eviscerate, Puncturing Wounds, Precision, Flawless Execution, Dual Wield Specialization, Hack and Slash, Weapon Expertise, Aggression, Opportunity, Serrated Blades, Quietus, Thousand Cuts, Initiative and Improved Ambush.<br>- Sets: Shadowcraft 3-piece (Subtlety: 18679 resist and 9336 +30 AP, both generic stat bonuses), Stormshroud 2/3 and Devilsaur 2.<br>- Carried, unchanged: Eviscerate's and Rupture's AP, Venom, and Combat's Slice and Dice gate (round 2's C2 measured it within error). | `bd/rogue-*-raid60.txt`; talents/rogue.json; ItemSet 184, ItemSetSpell 2418, 7910 | aligned (carried opens stand) | - | - |

## Beast Mastery, Marksmanship, Survival

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Survival's Mongoose Bite note said Expose Prey opens "the same 5 second window".** Expose Prey's activation has lasted 10 s since 70291, and the engine already uses 10 s; the dodge window is 5 s. Round 10 F9 found it and left it. Prose only. | `data/curated/apl/hunter-survival.json:89`; `sim/hunter/mongoose_bite.go:93` | CONFIRMED (copy) | old (round 10 F9 saw it) | S 545df3fd3, E d7cffdb40 |
| H2 | **A queued Raptor Strike still lifts the off hand's dual-wield miss penalty.** Blizzard removed that trick for Heroic Strike: "Queueing Heroic Strike will no longer increase off-hand hit chance" (level-30 development notes; wowfdb change log, 1 October). No note names Raptor Strike. A Survival hunter's off hand misses 15.5% against the main hand's 18.3% on the raid set because of it. If the fix covers every on-next-swing ability, Survival loses 0.8% raid (745.3 to 739.5) and 1.5% bare 60 (296.7 to 292.3). Measured with a scratch build, not committed. | `sim/hunter/raptor_strike.go:117`; https://eu.forums.blizzard.com/en/wow/t/one-of-warriors-most-satisfying-mechanics-is-disappearing-in-forever/634526 (HS only) | OPEN (a beta log of off-hand misses with Raptor Strike queued, or a Blizzard line on all next-swing abilities) | - | - |
| H3 | Checked and found aligned:<br>- Multi-Shot: 100% normalized, 13.9% base mana, shares category 2 with Aimed Shot.<br>- Aimed Shot: +166, 2 s.<br>- Arcane Shot: 217 flat, its own 6 s category 1173.<br>- Serpent Sting: 83 a tick, 415 over 15 s.<br>- Rapid Fire: 40%; Bestial Wrath: 50% for 18 s; Deadly Aspects: 10% for 30% over 12 s.<br>- Careful Aim: 100% of Intellect to both AP, auras 580 and 598.<br>- Raptor Strike: weapon damage +70, effect 58; Mongoose Bite: normalized +57; Strider Kick: 100% normalized, 5.81% base mana, 8 s; Wing Clip: 50; Counterattack: 50% normalized +110.<br>- Lacerate: 58 a tick ×7 at 60 (the tooltip's 451 is drawn at level 68).<br>- Predator's Edge: +30% crit damage, +50% off-hand damage.<br>- Lone Wolf, Improved Tracking and Surefooted.<br>- Survival at bare 20 (47 DPS against the published 70) is explained by round 9's pets-by-level fix (the published BiS dates from 8baffe5fc).<br>- Carried, unchanged: the Hawk proc mask, and Serpent Sting's and Arcane Shot's spell power. | SpellEffect, SpellCategories, SpellPower rows; Forever tooltips 2643, 14287, 13555, 20904, 1299332; `bd/hunter-*.txt` | aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Racials of the published races (human, orc, troll, dwarf, tauren, night elf) are Forever's reworked set: Ruthless 1%, Command, Blood Fury, Berserking 15 s, Artisan's Precision, Elune's Light and Endurance. No weapon-skill racials remain. The shipped ChrRaceRacialAbility text is older than the 26 October notes the engine follows. | ChrRaceRacialAbility (40 rows); `sim/core/racials.go` | aligned (as earlier rounds) | - | - |
| X2 | Glancing blows follow Era's formula: 10% plus 2% a point of defence over skill (40% against a level-63 boss), and a 0.55 to 0.75 damage band at skill 300. | `sim/core/target.go:473` | aligned | - | - |
| X3 | Every equip and use effect on the nine specs' published gear, at every band (82 rows, `items-effects.txt`), is either modelled or has no damage, threat or mana effect against a raid boss: heals, absorbs, movement, fear resist, slow fall, aura 87 taken-damage effects, detection, XP. Satchels of bombs proc only from a ranged auto attack, which a Protection warrior never makes. | ItemXItemEffect → ItemEffect → SpellEffect; `scratchpad/r11/physical/py/items.py` | aligned | - | - |

## Before/after

**Before:** `r11/sb-before` (cfcb76e80, site 8f047a62f). **After:** `lane/r11-physical` 12979ebe9 with `wt/r11-physical`. Each cell is seed 1 / seed 50001. Protection warrior and the bear are in threat per second.

| Spec, band | alliance % | horde % | seed 1 alliance / horde | cause |
|---|---|---|---|---|
| Fury 60 raid | −0.50 / −0.38 | −0.37 / −0.45 | 753.9 → 750.1 / 757.8 → 755.0 | F1 |
| Fury 30 bare | −5.96 / −5.96 | −5.92 / −5.93 | 63.8 → 60.0 / 64.2 → 60.4 | W1 |
| Fury 40 bare | −2.42 / −2.26 | −2.35 / −2.35 | 128.1 → 125.0 / 127.4 → 124.4 | W1, F1 |
| Fury 50 bare | −1.60 / −1.40 | −1.49 / −1.34 | 193.4 → 190.3 / 194.7 → 191.8 | W1, F1 |
| Fury 60 bare | −0.07 / −0.07 | −0.07 / −0.07 | 272.6 → 272.4 / 276.9 → 276.7 | F1 |
| Arms 30 bare | −6.75 / −6.76 | −6.58 / −6.60 | 56.3 → 52.5 / 56.2 → 52.5 | W1 |
| Arms 40 bare | −2.07 / −2.07 | −1.99 / −1.89 | 96.4 → 94.4 / 95.4 → 93.5 | W1 |
| Arms 50 bare | −0.70 / −0.70 | −0.70 / −0.70 | 184.6 → 183.3 / 186.0 → 184.7 | W1 |
| Protection 60 raid | +0.35 / +0.39 | +0.37 / +0.35 | 1079.0 → 1082.8 / 1067.5 → 1071.5 | P1 |
| Protection 20 / 30 / 40 / 50 / 60 bare (seed 1, alliance) | +2.4 / +2.2 / +1.5 / +1.1 / +0.7 | about the same (+2.1 to +2.4 at 20 and 30, +0.75 at 60) | 71.3 → 73.0 at 20; 634.8 → 639.5 at 60 | P1 |
| Bear 60 raid (hybrids group) | +1.32 / +1.26 | +1.24 / +1.18 | 1001.0 → 1014.2 / 992.4 → 1004.7 | P1 |
| Bear 20 / 30 / 40 / 50 / 60 bare | +2.4 / +2.4 / +2.1 / +1.7 / +1.3 | +2.5 / +2.4 / +2.1 / +1.8 / +1.4 | 55.2 → 56.5 at 20; 433.8 → 439.4 at 60 | P1 |

These are level to the decimal at both seeds and factions:
- Arms 60 raid, 20 and 60 bare; Fury 20 bare.
- Every rogue and hunter spec at 60 raid.
- Arms and Fury after P1 (DPS warriors are not hit).

## For other lanes

- **Integrator:** TestRotationLadder rows move for warrior-arms and warrior-fury (E 405bfeda0, E 83cd1abf7) and for warrior-protection and druid-feral-bear (E 12979ebe9). They were not regenerated here, per the brief.
- **Integrator:** E 12979ebe9 is in `sim/core/rage.go` and moves **druid-feral-bear**, which is outside this group (+1.2 to +1.3% raid, +1.3 to +2.5% bare). The hybrids lane found it and asked this lane to land it. The cat takes no damage, so it should not move; re-measure it at the pin.
- **Ranks lane:** Execute is now worth less below 56 (W1). If any per-band rank or rotation choice for Arms or Fury at 30 to 50 rests on Execute, it may need re-measuring.
- **Hybrids lane:** P1 is landed, with the measurements above.

## Still open

- Raptor Strike's queue and the off-hand penalty (H2).
- Normalized rage's off-hand factor and crit multiplier (A2), Lacerating Strikes' refresh (A3) and the one-a-minute proc rate (A4). The beta has closed; launch logs are the next chance.
- Marksmanship's +0.9% build (A1, held).
- Carried: armor-debuff stacking, Eviscerate's and Rupture's AP, Venom, Combat's Slice and Dice gate, the Hawk proc mask, Serpent Sting's and Arcane Shot's spell power, threat values, weapon-proc rates.

Sources (read 2026-11-02): the layered 70707 raw tables; wago.tools 1.60.1.70753 SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellAuraOptions; https://nether.wowhead.com/forever/tooltip/spell/{5308,20658,20660,20661,20662,1680,11601,23925,11574,21553,11585,23894,11567,11605,7373,2565,11581,2643,14287,13555,20904,1299332}; https://wowfdb.com/classes/warrior/changes/; https://wowfdb.com/classes/hunter/changes/; the cmangos-classic Player.cpp above; the Kalgan reposts above.
