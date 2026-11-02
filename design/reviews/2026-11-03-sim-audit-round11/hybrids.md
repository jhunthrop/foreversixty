# Sim audit round 11: hybrids (2026-11-02)

This report covers Retribution, Protection paladin (threat), Enhancement, Feral cat and Feral bear (threat).

**Setup:**
- Client 1.60.1.70707, layered tables (hotfix push included; `scratchpad/r10/crit/raw`, fetched 2026-11-01 21:42Z). `python -m pipeline push-check` logic run today: wago lists no push the committed record lacks, so these tables are current.
- Extra client tables fetched today from wago.tools for 70707: SpellShapeshift, SpellProcsPerMinute, SpellAuraRestrictions, SpellTargetRestrictions (`scratchpad/r11/hybrids/wago`).
- Published sets and builds: the committed `data/builds/1.60.1.70707/bis/`. Nothing nightly-owned was regenerated or committed.
- Engine: pin cfcb76e80 plus `lane/r11-hybrids`. Site: `wt/r11-hybrids` on origin/main 8f047a62f.
- Method:
  - every band (bare 20 to 60, raid 60) of every spec was run and its cast, aura and resource tables read (`scratchpad/r11/hybrids/bd`);
  - one debug-log iteration of each spec at raid 60 and bare 40 was read line by line (a scratch, never-committed copy of spec-breakdown with `Debug`; `scratchpad/r11/hybrids/logs`);
  - every talent the published builds take was decoded and its client effects and class masks resolved to the spells they name (`tools/talmask.py`), then compared with the engine;
  - every spell cast at a bare band was compared with the highest rank the client teaches by then (SkillLineAbility, AcquireMethod not 3);
  - the use and equip spells of every published item, the published races' racials (ChrRaceRacialAbility, SkillLineAbility) and the in-form rules (SpellShapeshift) were read against the engine;
  - Wowhead Forever tooltips (nether.wowhead.com/forever/tooltip, read 2026-11-02) where a client row needs the server's text.

Fix-sha column: "E" is a fork commit on `lane/r11-hybrids`, "S" a site commit on `wt/r11-hybrids`.

**In short: 1 CONFIRMED-BUG, fixed, origin old.** Benediction did not discount Holy Shield. No published build takes both, so no published number moves. Everything else checked is aligned or stays open on logs.

## Round 10 open items (assigned)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| O1 | **A stone's damage on a form's swing.** Nothing new: no cat or bear log with and without a stone. GET api.foreversixty.gg/v1/reports/recent returns three reports, the newest from 2026-10-21. foreverlogs.gg holds only the beta (levels 13 to 30). | API, 2026-11-02 | OPEN (needs a launch log) | - | - |
| O2 | **Oath of Light.** The live tooltip (1328548) still reads "increasing the healing of your Holy Light on them by up to 0". No published hybrid takes it. | Wowhead Forever, 2026-11-02 | OPEN (healers) | - | - |
| O3 | **Holy Shield crit.** No 40+ Forever log exists yet. The engine still rolls spell crit on the shield's damage. | as O1 | OPEN | - | - |
| O4 | **Thundering Strikes on shaman spells.** No new evidence. The engine follows the beta logs (every table), as round 10 recorded. | as O1 | OPEN (crit lane's ruling) | - | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Seals and the twist.** Read in the iteration log:<br>- Seal of Command 20920 (210 mana, 1 s proc cooldown) triggers 20424 at every rank, 70% weapon damage.<br>- Seal of Righteousness 20293: the proc's damage comes from the seal's dummy (1786 + 47 a level). Wowhead's tooltip gives the same per-swing range and "slower weapons cause more Holy damage per swing".<br>- Judgement does not consume the Seal in Forever (the engine's note quotes the Deep Dive), so 22 Judgements need only 30 Seal casts.<br>- The round-10 twist waits for the Righteousness swing. Unsealed Wrath's stored copies roll without Seal of Command's 1 s cooldown. That is unconfirmed: the client row (1310735) is a dummy, and the engine says so. | client rows; `logs/paladin-retribution-raid60.log` | Aligned | - | - |
| P2 | **Judgement.** 20271 costs 6% of base mana, 10 s, minus 2 s for Improved Judgement. Judgement of Righteousness (0.429 coefficient, melee table, no dodge or parry). Judgement of Command is 356 ± 9.6%, halved unless stunned. Wowhead's rank-1 text, "69 to 73 Holy damage, 138 to 146 if the target is stunned", matches 97 + 5.6 a level at its level 28 cap, halved. Sanctified Judgement 3/3 returns 60% of the judged Seal's base cost (2,749 mana a fight). | client rows; tooltip; breakdown | Aligned | - | - |
| P3 | **Holy Strike** (Forever's Crusader Strike slot): 93 + 3.2 a level (effect 121) times 50% (effect 31), 10 s, 20 mana, melee weapons only (SpellEquippedItems class 2, mask 173555). Sacred Arbiter +20%. Ranks by band are 1866, 680, 2495, 5569 and 10333, the highest the client teaches at each. Its crit equals the white swings' on the two-roll table (21.4% of attempts against white's 22.8% single-roll at raid 60). | client rows; breakdowns | Aligned | - | - |
| P4 | **Hammer of Wrath.** 498 ± 9.8%, ranged table, 6 s, 425 mana, minus 40% from Holy Conduit (255 paid). Benediction's mask leaves it out. Champion of the Light makes it instant. It is cast from 25% of the fight left because the ranker's encounter spends 25% below 20% health (`api.DefaultEncounter`, ExecuteRatio 0.25). That is a platform setting, not a spell bug. | client rows; log | Aligned | - | - |
| P5 | **Consecration.** Rank tick values 12 + 27 for the first four targets plus 0.042 of spell power, 8 one-second ticks, 8 s cooldown, 565 mana, minus Holy Conduit and Benediction (282.5 paid). The pace line is round 10's; owner ruling 3 accepts the bare-20 loss, so it was not reopened. | client rows (20924, 1280349) | Aligned | - | - |
| P6 | **Talent masks**, resolved from each talent's SpellClassMask to the spells it names:<br>- Improved Seals: Seals and Judgements of Command, Fury and Righteousness, and their procs.<br>- Holy Conduit: Consecration, Exorcism, Hammer of Wrath, Holy Wrath.<br>- Sacred Arbiter: Holy Strike.<br>- Champion of the Light: spell damage from Intellect, Hammer of Wrath's cast.<br>- Two-Handed Weapon Specialization: physical damage and -20% threat.<br>- Vengeance: 3% a stack, three stacks, from any non-periodic crit.<br>- Conviction: melee crit, and heals per round 10.<br>- Benediction: see X1.<br>All match the engine. | `tools/talmask.py` | Aligned | - | - |
| P7 | **Sanctity Aura** is not in Forever's Retribution tree (talents/paladin.json 70707), and no preset carries it. | talents | Aligned | - | - |
| P8 | **Undead racial.** Touch of the Grave: 20%, 15 s cooldown (pushed), 5% of max health, spell hit, cannot crit (Attributes_2 0x20000000), fixed damage (Attributes_4 0x100). The other Undead racials have no combat effect. | client rows; tooltip | Aligned | - | - |
| P9 | **Gear.** Rune of the Guard Captain (Authority: static attack power), Molten and Frozen Heart of the Mountain (5 min use, 15 s), Arcanite Champion's proc, Cloak of Flames (a damage shield), Manual Crowd Pummeler (3 charges, 3 min). | ItemEffect rows | Aligned | - | - |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **Righteous Fury** (25780): +90% Holy threat (effect 0) and +33% to all threat (effect 4, pushed 31 October). The two multiply. Improved Righteous Fury is damage taken, not threat. | client row | Aligned | - | - |
| T2 | **Holy Shield** 20925/20927/20928: 30% block, 4 charges, 10 s, damage 75/110/153 plus 0.08 of spell power, +20% threat. Shieldsworn's Holy Strike fires the damage without spending a charge, which is why the damage lands 61 times a fight against 43 blocks. The mana cost missed Benediction: X1. | client rows; log | Aligned (cost: X1) | - | - |
| T3 | **Redoubt** (20127: 10% on a damaging melee hit taken; 20128: +30% block, 5 charges, 10 s): 6.4 procs from 64 landed swings. **Reckoning** is not in the published builds. | client rows; breakdown | Aligned | - | - |
| T4 | Seal of Fury 35 flat, plus a 50% absorb with a shield (Wowhead: "an additional 35 Holy damage"). Judgement of Fury, Swift Judgement (60 s, does not consume), Hammer of the Righteous (6% base mana, 6 s, one-handers only), Iron Creed (+30% Holy Strike threat), Shield Specialization (+30% block value), Divine Shield (its -50% damage dealt is modelled), Templar's Bulwark. | client rows; tooltips | Aligned | - | - |
| T5 | Gear: Guardian Talisman's "Guarded" is a stealth-detection buff, and Arena Master works only in the Battle Ring and The Maul. Both carry stats only. | tooltips 1300364, 1317285 | Aligned | - | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | **Stormstrike** 17364: 8% base mana, 8 s, weapon damage. Its 20% (aura 271) covers Lightning Bolt, Chain Lightning and Earth Shock only (mask 0x100003). It has 1 charge, and Attributes_3 0x80 makes only the shaman's own spell spend it. | client row | Aligned | - | - |
| N2 | **Windfury Weapon.** Every enchant rank procs 439431 (20%, 1.5 s). The bonus attack power is the rank's own row: Wowhead gives 104/221/315/433 at each rank's cap, and the engine reads 333 + 12.5 a level for rank 4. Elemental Weapons +40% at 3/3. **Windfury Totem** rank 3 is 246 attack power (10610), read by the engine. | client rows; tooltips 8232/8235/10486/16362 | Aligned | - | - |
| N3 | **Shock rotation.** Earth Shock 10414: 450 mana, minus 45% (Shamanistic Focus) to 247.5, 6 s minus Reverberation. Lightning Bolt is cast free and instant at five Maelstrom stacks. Searing Totem lasts 55 s. Flametongue and Flame Shock are not cast at 60, as rounds 9 and 10 measured. | log; client rows | Aligned | - | - |
| N4 | **Talent masks:** Concussion, Convection, Elemental Alacrity, Reverberation, Call of Flame (fire totems' "Attack" included), Shamanistic Focus, Elemental Weapons (Rockbiter, Windfury, Flametongue, Frostbrand). Tidal Focus is +1% hit a rank on melee and spell, and the raid build takes 5/5. Mindfulness lets 50% of regen continue while casting. Mental Dexterity: 100% of Intellect as attack power (tooltip). Mental Quickness: 30% of Intellect as spell damage and healing. Ancestral Knowledge. Rage of the Farseer: 30% attack speed, 25 s, 3 min. | `tools/talmask.py`; tooltip 415140 | Aligned | - | - |
| N5 | **Flurry** (25% at 5/5, 3 charges) and **Elemental Devastation** (melee crit from a direct-damage spell crit) are as round 10 found. **Unleashed Rage** and Shamanistic Rage are not in Forever's tree. The shaman wields a two-hander. | talents 70707 | Aligned | - | - |
| N6 | **Orc.** Blood Fury: 10% attack power and spell power, 15 s, 2 min. Shatter Curse, Hardiness, and Command, which damages pets only. The UI table ChrRaceRacialAbility still lists "Savage, +1% crit", but the SkillLineAbility rows that teach orcs their racials have Command and no Savage. The engine follows SkillLineAbility. Forever orcs have no axe specialization. | SkillLineAbility skill 125 | Aligned | - | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Shred** 155% + 80, 60 energy, minus 18 for Shredding Attacks. **Rip**: 30 energy, 7 ticks with Idol of the Dream (+2 s; 446212 aura 107, op 1, Rip's mask). Rip ticks crit on the melee table and take the stone. **Ferocious Bite** rank 4 in Forever (22829): 45 to 95 + 128 a point + 2.5 an energy. That matches Wowhead's "1 point: 173-223 ... 5 points: 685-735". **Ravage** from Prowl. | client rows; tooltips 22829, 31018, 22568 | Aligned | - | - |
| F2 | **Rip refresh line.** The rotation refreshes Rip under 6 s left. Thresholds of 0 to 4 s measure 666.6 to 672.8 against 670.3 (3,000 iterations, seed 1). The best gains 0.4%, under the 1% bar, so the line stays. | `-rotation` runs (`cat-rip*.json`) | Aligned | - | - |
| F3 | **Tiger's Fury** is not learnable in Forever (AcquireMethod 3), as round 3 found. **Shifting Power**: 55% of base mana for 40 energy, 8 s, 1 s global (684 mana at 60). **Omen of Clarity**: 2 procs a minute on melee and 4% on spells, with a 10 s cooldown, from the beta logs. **Faerie Fire** is castable in Cat, Bear and Dire Bear Form (SpellShapeshift mask 0x40000091), free with a 6 s cooldown through the form passives (1178/9635 effects 4 and 5). | client rows; SpellShapeshift | Aligned | - | - |
| F4 | **Powershifting.** No Wolfshead Helm in the published sets, and Furor's cat half restores the energy you had. Not modelled, as the APL notes say. | talents; APL | Aligned | - | - |
| F5 | **Leader of the Pack** (24932) is aura 290 at 3%, the same aura as the raid's. Nature's Majesty, Sharpened Claws, Natural Instinct (Claw, Rake, Shred, Ravage, Pounce, Bite, Rip, Maul, Swipe, Lacerate and Primal Bite), Savage Fury (Rake's tick included), Genesis (periodic), Naturalist (+10% all damage), Rend and Tear, Blood Frenzy, Predatory Discipline. | `tools/talmask.py` | Aligned | - | - |
| F6 | **Gear.** Shadowcraft 5-piece (Rogue Armor Energize, 20 energy at an unstated rate, 1 a minute: the item_sets lane's assumption), Hammer of the Titans' stun, Thistle Tea, Idol of the Dream. **Tauren** Endurance: the push removed the shipped 2% crit damage and healing (1322984, 1349372). Wowhead's live text is "Total Health increased by 5%." | shipped vs layered SpellEffect; tooltip 20550 | Aligned | - | - |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | **Bear Form.** Dire Bear Form (Passive) 9635: armor +360% (auras 142 and 466), health 600 + 32 a level (1,240 at 60), attack power 120 + 3 a level. 21178 gives +50% threat. | client rows | Aligned | - | - |
| G2 | **Maul** 15 rage minus Ferocity's 5, x1.75 threat, on the next swing, and it takes a Windfury Totem swing. **Swipe** rank 5 (9908), the top druid rank (the 1264498 to 1264502 "Swipe" ranks are a hunter pet's). **Lacerate** ranks by band (1328616 at 40, 1235827 at 60), 12 rage. Primal Bite and Lacerate threat use the unconfirmed 2.0 "high threat" tier (round 9). | client rows; SkillLineAbility | Aligned (threat tier open) | - | - |
| G3 | **Rage.** Missed, dodged and parried abilities refund 80% (logged on Maul and Lacerate misses). Natural Reaction gives 5 rage a dodge and Blood Frenzy 5 a crit. Barkskin is castable in form: no SpellShapeshift row and no NOT_SHAPESHIFTED flag. | log; SpellShapeshift | Aligned | - | - |
| G4 | **Rage from damage taken uses the boss's level.** `sim/core/rage.go:175` reads the attacker's level (63, conversion 243.6). The comment is upstream wowsims' ("based on attacker's level, not player's", 47f540b04, 2024). cmangos-classic's `Player::RewardRage` uses the victim's own level (230.6 at 60). Using the bear's level measured +1.2% threat at raid 60 (992.4 to 1004.7 at seed 1, 991.6 to 1003.3 at 50001; scratch build, not committed). One emulator against one engine comment, so not changed. | cmangos-classic master, Player.cpp `RewardRage`; scratch measure | OPEN (needs a Forever log of rage from hits taken) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Benediction did not discount Holy Shield.**<br>- What the client says: Benediction (20101, aura 108, cost) names every instant its SpellClassMask covers. Holy Shield 20925/20927/20928 is in that mask (SpellClassMask_1 0x40), next to the Seals, Judgement, Consecration, Exorcism and Holy Strike.<br>- What was wrong: every other masked instant took `paladin.benediction()`, but Holy Shield's mana cost did not. That line dates from the upstream port (2024). Round 4's Benediction fix (a3e569f13) checked the mask and missed this one.<br>- What changed: Holy Shield costs 90% at Benediction 5/5.<br>- Published numbers: none move. No published build takes both Benediction and Holy Shield. A talent search that tries Benediction on Protection now prices Holy Shield right. | Test first: `TestBenedictionDiscountsEveryInstantItsMaskNames` (sim/paladin/retribution) now covers Holy Shield. It failed at 100% and passes at 90. | CONFIRMED-BUG | old (upstream; round 4 missed it) | E 93d07f72d |
| X2 | Every spell cast at a bare band is the highest rank the client teaches by then, except the curated Consecration rank choices, the Ahn'Qiraj books (Blessing of Might 25291, Grace of Air 25359, Strength of Earth 25361, Ferocious Bite 31018) and Prowl, which has no combat effect. | SkillLineAbility sweep | Aligned | - | - |
| X3 | Engine ids missing from the client (Seal of Command procs 20944 to 20947, Holy Shield 20955 to 20957, the Seal of Fury shield 1311648) are named in `data/curated/engine-spell-names.json`, so players see names. | curated names | Aligned | - | - |

## Before/after

Before: `$R/sb-before` (pin cfcb76e80, site 8f047a62f). After: `lane/r11-hybrids` 93d07f72d with `wt/r11-hybrids`. Raid 60, 10,000 iterations. DPS, threat per second for the tanks.

| spec | horde s1 | horde s50001 | alliance s1 | alliance s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| paladin-retribution | 0.0% | 0.0% | 0.0% | 0.0% | 685.7 -> 685.7 |
| paladin-protection (threat) | 0.0% | 0.0% | 0.0% | 0.0% | 1013.5 -> 1013.5 |
| shaman-enhancement | 0.0% | 0.0% | 0.0% | 0.0% | 580.8 -> 580.8 |
| druid-feral | 0.0% | 0.0% | 0.0% | 0.0% | 670.0 -> 670.0 |
| druid-feral-bear (threat) | 0.0% | 0.0% | 0.0% | 0.0% | 992.4 -> 992.4 |

The bare bands do not move either: no bare build takes Benediction and Holy Shield together.

## For other lanes

- **Physical lane (warrior tank) and integrator:** G4. The rage conversion for damage taken reads the attacking boss's level. cmangos-classic reads the victim's. For the bear it is +1.2% threat at raid 60; the warrior tank is likely similar. Only a log of rage gained from hits taken on a Forever server settles it.
- **Platform:** P4. The ranker's encounter spends 25% of the fight below 20% health (`api.DefaultEncounter` ExecuteRatio 0.25), so execute spells start at 75% of the fight. This is noted, not filed as a bug.
- **Integrator:** E 93d07f72d moves no golden (`go test -tags=with_db ./sim/paladin/...` passes). It moves no spec outside my group. No held-output test was deselected. `TestRotationLadder` was not run or regenerated.

## Still open

1. O1 to O4: a stone on a form's swing, Oath of Light, Holy Shield crit, Thundering Strikes on spells. No log at the right level yet; launch is 4 November.
2. G4: which level the rage conversion reads for damage taken.
3. From round 10: Faerie Fire's form threat; Lacerate's and Primal Bite's 2.0 threat; Seal of Command and Stormstrike firing weapon procs (no SpellEquippedItems rows; emulator-only evidence); Unsealed Wrath's stored copies against Seal of Command's 1 s cooldown.
