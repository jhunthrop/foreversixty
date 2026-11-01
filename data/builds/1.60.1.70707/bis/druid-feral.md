# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 45.6. Weights run: 4.4s. Verify run: 4.9s. 203 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.994 ± 0.011, crit=0.615 ± 0.010 per rating point (14 rating = 1%, 8.610 per %), hit=0.708 ± 0.030 per rating point (10 rating = 1%, 7.078 per %), melee_haste=4.033 ± 0.358

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.91 DPS) | yes | Brawler's Leather Hood (252504, -0.00 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.0 attack_power points (0.23 DPS) | yes | Tarnished Locket (279870, -0.23 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 attack_power points (0.57 DPS) | yes | Slime-encrusted Pads (6461, -0.56 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 attack_power points (0.68 DPS) | yes | Grave Shroud (279865, -0.11 DPS) [quest]; Dark Leather Cloak (2316, -0.11 DPS) [crafted]; Cape of the Brotherhood (5193, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (45.6 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.52 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 12.0 attack_power points (0.68 DPS) | yes | Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Forest Leather Bracers (3202, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.23 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (45.6 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.65 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.02 DPS) | yes | Deviate Scale Belt (6468, -0.12 DPS) [crafted]; Murloc Scale Belt (5780, -0.46 DPS) [crafted]; Brawler's Leather Belt (252428, -4.24 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (45.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.75 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.0 attack_power points (1.25 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Blackened Defias Boots (10402, -0.57 DPS) [dungeon]; Footpads of the Fang (10411, -0.57 DPS) [dungeon] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.45 DPS) | yes | Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Ring of the Moon (12052, -0.34 DPS) [world_drop] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 4.0 attack_power points (0.23 DPS) | yes | Lavishly Jeweled Ring (1156, -0.00 DPS) [dungeon]; Ring of the Moon (12052, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 305.9 attack_power points (17.40 DPS) | yes | Staff of Westfall (2042, -1.08 DPS) [quest]; Twisted Chanter's Staff (890, -1.12 DPS) [world_drop]; Living Root (6631, -1.90 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: The 1 Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 103.1. Weights run: 4.4s. Verify run: 3.1s. 387 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.887 ± 0.012, crit=0.708 ± 0.012 per rating point (14 rating = 1%, 9.906 per %), hit=0.728 ± 0.043 per rating point (10 rating = 1%, 7.282 per %), melee_haste=5.148 ± 0.533

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.76 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.44 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.44 DPS) [crafted]; Brawler's Leather Helm (252512, -0.48 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.5 attack_power points (1.14 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.26 DPS) [world]; River Pride Choker (13087, -0.55 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.70 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.8 attack_power points (2.26 DPS) | yes | Mantle of Thieves (2264, -0.87 DPS) [dungeon]; Bristlebark Amice (14573, -0.99 DPS) [world_drop]; Barbaric Shoulders (5964, -1.12 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.2 attack_power points (1.41 DPS) | yes | Sergeant Major's Cape (16315, -0.27 DPS) [pvp]; Tigerstrike Mantle (13108, -0.30 DPS) [world_drop]; Fenrus' Hide (6340, -0.58 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 27.1 attack_power points (1.99 DPS) | yes | Dusky Leather Armor (7374, -0.05 DPS) [crafted]; Defender's Leather Tunic (252450, -0.28 DPS) [crafted]; Brawler's Leather Armor (252490, -0.29 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.3 attack_power points (1.42 DPS) | yes | Barbaric Bracers (18948, -0.38 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.42 DPS) [quest]; Bands of Serra'kis (6902, -0.54 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.3 attack_power points (1.86 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.44 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.3 attack_power points (2.15 DPS) | yes | Skulker's Leather Belt (252520, -0.03 DPS) [crafted]; Warden's Leather Belt (252460, -0.44 DPS) [crafted]; Blackened Defias Belt (10403, -0.83 DPS) [dungeon] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 31.0 attack_power points (2.28 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Barbaric Leggings (5963, -0.28 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.1 attack_power points (1.55 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Insignia Boots (4055, -0.44 DPS) [world_drop]; Disjointed Shoes (277226, -0.67 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.7 attack_power points (1.59 DPS) | yes | Ring of Precision (1491, -0.76 DPS) [dungeon]; Tiger Band (6749, -0.86 DPS) [quest]; Bonebinder's Signet (273032, -0.86 DPS) [world] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 11.7 attack_power points (0.86 DPS) | yes | Ring of Precision (1491, -0.03 DPS) [dungeon]; Tiger Band (6749, -0.12 DPS) [quest]; Bonebinder's Signet (273032, -0.12 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (103.1 DPS) | yes | Cobalt Crusher (7730, -3.14 DPS) [dungeon]; Gnarled Ash Staff (791, -6.52 DPS) [world_drop]; Viscous Hammer (13045, -16.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Seal of Wrynn; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 147.7. Weights run: 5.2s. Verify run: 3.3s. 523 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.952 ± 0.028, crit=0.951 ± 0.034 per rating point (14 rating = 1%, 13.311 per %), hit=0.788 ± 0.072 per rating point (10 rating = 1%, 7.881 per %), melee_haste=4.266 ± 0.914

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 43.5 attack_power points (3.70 DPS) | yes | Warden's Wizard Hat (14604, -1.54 DPS) [world_drop]; Hawkeye's Helm (14591, -1.59 DPS, sim-verified) [world_drop]; Defender's Leather Helm (252455, -1.66 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.8 attack_power points (1.35 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.32 DPS) [world]; River Pride Choker (13087, -0.66 DPS) [world_drop]; Gazlowe's Charm (13088, -0.66 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.5 attack_power points (2.68 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.83 DPS) [dungeon]; Nightscape Shoulders (8192, -0.85 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 23.7 attack_power points (2.02 DPS) | yes | Yeti Fur Cloak (2805, -0.51 DPS) [quest]; Parachute Cloak (10518, -0.69 DPS) [crafted]; Hawkeye's Cloak (14593, -1.24 DPS, sim-verified) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 33.2 attack_power points (2.83 DPS) | yes | Nightscape Tunic (8175, -0.33 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.47 DPS) [crafted]; Barbaric Harness (5739, -0.56 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.70 DPS) | yes | Hawkeye's Bracers (14590, -0.02 DPS) [world_drop]; Barbaric Bracers (18948, -0.36 DPS) [crafted]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 37.6 attack_power points (3.20 DPS) | yes | Skulker's Leather Gloves (252525, -0.00 DPS) [crafted]; Imperial Leather Gloves (4063, -0.17 DPS) [dungeon]; Gloves of Holy Might (867, -0.36 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 35.6 attack_power points (3.03 DPS) | yes | Skulker's Leather Belt (252520, -0.51 DPS) [crafted]; Prowler's Leather Belt (252459, -0.58 DPS, sim-verified) [crafted]; Scorpashi Sash (14652, -0.69 DPS) [world_drop] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 41.1 attack_power points (3.50 DPS) | yes | Basilisk Hide Pants (1718, -0.01 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.81 DPS) [crafted]; Brawler's Leather Pants (252500, -0.99 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 35.7 attack_power points (3.04 DPS) | yes | Skulker's Leather Shoes (252531, -0.02 DPS) [crafted]; Excelsior Boots (4109, -0.35 DPS) [quest]; Imperial Leather Boots (6431, -0.36 DPS) [dungeon] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.6 attack_power points (2.01 DPS) | yes | Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Mark of Kern (2262, -0.30 DPS) [dungeon]; Assault Band (13095, -0.30 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.5 attack_power points (2.00 DPS) | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Mark of Kern (2262, -0.30 DPS) [dungeon]; Assault Band (13095, -0.30 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (147.7 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Staff of Jordan (873, -11.25 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Falcon's Hook; finger2: Ring of the Underwood

No-known-source sample (15 of 523, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0000000000000000-50212032123012123001-5500000000000000)

Set DPS (verified): 191.4. Weights run: 5.1s. Verify run: 3.8s. 693 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.941 ± 0.030, crit=1.123 ± 0.038 per rating point (14 rating = 1%, 15.723 per %), hit=0.998 ± 0.091 per rating point (10 rating = 1%, 9.976 per %), melee_haste=not significant (3.893 ± 1.141)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.5 attack_power points (3.88 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.33 DPS) [vendor]; Helm of Fire (8348, -1.07 DPS) [crafted]; Scorpashi Skullcap (14658, -1.47 DPS, sim-verified) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.2 attack_power points (3.08 DPS) | yes | Archmagister's Faceted Pendant (273049, -2.06 DPS) [world]; River Pride Choker (13087, -2.33 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.66 DPS, sim-verified) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 41.7 attack_power points (3.55 DPS) | yes | Failed Flying Experiment (9647, -0.02 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.04 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.68 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 29.1 attack_power points (2.48 DPS) | yes | Sergeant Major's Cape (16336, -0.36 DPS) [pvp]; Pridelord Cape (14673, -0.43 DPS) [world_drop]; Blackflame Cape (13109, -0.50 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 59.1 attack_power points (5.03 DPS) | yes | Blazewind Breastplate (11193, -0.92 DPS, sim-verified) [quest]; Knight's Leather Armor (220854, -1.31 DPS) [vendor]; Charred Leather Tunic (19127, -2.06 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.9 attack_power points (3.23 DPS) | yes | Prowler's Leather Bracers (252539, -0.39 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.43 DPS) [crafted]; Pridelord Bands (14672, -0.64 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 48.3 attack_power points (4.11 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.11 DPS) [crafted]; Serpentskin Gloves (8260, -0.23 DPS) [world_drop]; Heraldic Gloves (8121, -0.58 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 54.1 attack_power points (4.60 DPS) | yes | Skulker's Leather Waistguard (252474, -0.04 DPS) [crafted]; Ogron's Sash (13117, -1.43 DPS) [world_drop]; Pridelord Girdle (14674, -1.87 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 62.1 attack_power points (5.29 DPS) | yes | Serpentskin Leggings (8262, -0.21 DPS) [world_drop]; Knight's Leather Pants (220858, -1.57 DPS) [vendor]; Triprunner Dungarees (9624, -1.75 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.9 attack_power points (4.25 DPS) | yes | Skulker's Leather Boots (252469, -0.04 DPS) [crafted]; Albino Crocscale Boots (17728, -0.95 DPS) [dungeon]; Prowler's Leather Shoes (252465, -1.04 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 27.2 attack_power points (2.31 DPS) | yes | Ring of the Underwood (2951, -0.29 DPS) [world_drop]; Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Blackstone Ring (17713, -0.61 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.1 attack_power points (2.05 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Blackstone Ring (17713, -0.35 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (191.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (191.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (191.4 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -17.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 693, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0000000000000000-50212032123012123001-5550230000000000)

Set DPS (verified): 276.5. Weights run: 5.1s. Verify run: 10.1s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.351 ± 0.047, crit=1.930 ± 0.067 per rating point (14 rating = 1%, 27.024 per %), hit=1.980 ± 0.144 per rating point (10 rating = 1%, 19.801 per %), melee_haste=not significant (6.636 ± 1.940)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 89.9 attack_power points (8.54 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -1.85 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.6 attack_power points (3.95 DPS) | yes | Sunprism Pendant (20649, -1.90 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -2.22 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.81 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (276.5 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -5.71 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 55.3 attack_power points (5.25 DPS) | yes | Sergeant Major's Cape (16337, -1.60 DPS, sim-verified) [pvp]; Dark Phantom Cape (13122, -1.90 DPS) [world_drop]; Blackflame Cape (13109, -2.57 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 141.9 attack_power points (13.49 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, -1.73 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -2.57 DPS) [vendor]; Dawn Armor (252483, -2.87 DPS, sim-verified) [crafted] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (276.5 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -6.16 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (276.5 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -3.49 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (276.5 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -3.49 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 103.9 attack_power points (9.87 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -0.24 DPS) [vendor]; Traveler's Leggings (8300, -3.01 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 55.3 attack_power points (5.26 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -6.23 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.9 attack_power points (3.13 DPS) | yes | Falcon's Hook (7552, -0.49 DPS) [dungeon]; Thunderbrow Ring (13097, -0.79 DPS) [world_drop]; Ring of Living Stone (18400, -1.04 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.9 attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.31 DPS) [world_drop]; Ring of Living Stone (18400, -0.56 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (276.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (276.5 DPS) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (276.5 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -6.18 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (276.5 DPS) | yes | Idol of the Moon (23197, -8.97 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; trinket2: Mark of the Chosen; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 0500200000000000-50230032023011125001-5050030000000000)

Set DPS (verified): 656.5. Weights run: 4.9s. Verify run: 9.9s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.859 ± 0.061, crit=2.286 ± 0.083 per rating point (14 rating = 1%, 31.998 per %), hit=2.077 ± 0.194 per rating point (10 rating = 1%, 20.774 per %), melee_haste=11.317 ± 2.487

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 108.0 attack_power points (18.89 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -3.36 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 48.2 attack_power points (8.42 DPS) | yes | Sunprism Pendant (20649, -3.26 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -4.89 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -6.33 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (656.5 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -15.43 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.9 attack_power points (11.00 DPS) | yes | Dark Phantom Cape (13122, -3.17 DPS, sim-verified) [world_drop]; Sergeant Major's Cape (16337, -3.92 DPS) [pvp]; Blackflame Cape (13109, -5.00 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 158.6 attack_power points (27.73 DPS) | yes | Timbermaw Tunic (252484, -0.44 DPS) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -3.81 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -5.52 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (656.5 DPS) | yes | Skulker's Leather Bracers (252540, -0.31 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.54 DPS) [crafted]; Deepfury Bracers (13120, -17.27 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (656.5 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -13.27 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (656.5 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -13.27 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 113.1 attack_power points (19.77 DPS) | yes | Knight-Captain's Dragonhide Leggings (227178, +0.00 DPS) [vendor]; Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Traveler's Leggings (8300, -5.35 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 66.3 attack_power points (11.59 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -17.34 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 40.0 attack_power points (7.00 DPS) | yes | Falcon's Hook (7552, -1.35 DPS) [dungeon]; Thunderbrow Ring (13097, -2.42 DPS) [world_drop]; Blush Ember Ring (13093, -2.96 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 33.0 attack_power points (5.77 DPS) | yes | Falcon's Hook (7552, -0.12 DPS) [dungeon]; Thunderbrow Ring (13097, -1.19 DPS) [world_drop]; Blush Ember Ring (13093, -1.73 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (656.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (656.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (656.5 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -17.60 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (656.5 DPS) | yes | Idol of the Moon (23197, -9.82 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 44.8. Weights run: 4.4s. Verify run: 4.9s. 199 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.994 ± 0.011, crit=0.615 ± 0.010 per rating point (14 rating = 1%, 8.610 per %), hit=0.708 ± 0.030 per rating point (10 rating = 1%, 7.078 per %), melee_haste=4.033 ± 0.358

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.91 DPS) | yes | Brawler's Leather Hood (252504, -0.00 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.0 attack_power points (0.23 DPS) | yes | Roadwatcher's Confidence (281265, -0.23 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 attack_power points (0.57 DPS) | yes | Slime-encrusted Pads (6461, -0.57 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 attack_power points (0.68 DPS) | yes | Grave Shroud (279865, -0.11 DPS) [quest]; Dark Leather Cloak (2316, -0.11 DPS) [crafted]; Cape of the Brotherhood (5193, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (44.8 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.38 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 10.0 attack_power points (0.57 DPS) | yes | Forest Leather Bracers (3202, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor]; Spare Part Bindings (279875, -0.23 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (44.8 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.51 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.02 DPS) | yes | Deviate Scale Belt (6468, -0.12 DPS) [crafted]; Murloc Scale Belt (5780, -0.46 DPS) [crafted]; Brawler's Leather Belt (252428, -4.08 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (44.8 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.60 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.0 attack_power points (1.25 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Blackened Defias Boots (10402, -0.57 DPS) [dungeon]; Footpads of the Fang (10411, -0.57 DPS) [dungeon] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.45 DPS) | yes | Bounty Hunter's Ring (5351, -0.11 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.34 DPS) | yes | Bounty Hunter's Ring (5351, -0.00 DPS) [quest]; The 1 Ring (8350, -0.11 DPS) [world]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (17.41 DPS) | yes | Smite's Mighty Hammer (7230, -0.01 DPS) [dungeon]; Crescent Staff (6505, -0.42 DPS) [quest]; Living Root (6631, -0.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Hammerbone

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 101.7. Weights run: 4.4s. Verify run: 3.2s. 387 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.887 ± 0.012, crit=0.708 ± 0.012 per rating point (14 rating = 1%, 9.906 per %), hit=0.728 ± 0.043 per rating point (10 rating = 1%, 7.282 per %), melee_haste=5.148 ± 0.533

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.76 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.44 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.44 DPS) [crafted]; Brawler's Leather Helm (252512, -0.47 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.5 attack_power points (1.14 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.29 DPS, sim-verified) [world]; River Pride Choker (13087, -0.55 DPS) [world_drop]; Erudite's Amulet (277204, -0.87 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.8 attack_power points (2.26 DPS) | yes | Mantle of Thieves (2264, -0.87 DPS) [dungeon]; Bristlebark Amice (14573, -0.99 DPS) [world_drop]; Barbaric Shoulders (5964, -1.13 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.2 attack_power points (1.41 DPS) | yes | Tigerstrike Mantle (13108, -0.40 DPS, sim-verified) [world_drop]; Fenrus' Hide (6340, -0.58 DPS) [dungeon]; Swiftrunner Cape (6745, -0.58 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 27.1 attack_power points (1.99 DPS) | yes | Dusky Leather Armor (7374, -0.05 DPS) [crafted]; Defender's Leather Tunic (252450, -0.28 DPS) [crafted]; Brawler's Leather Armor (252490, -0.29 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.3 attack_power points (1.42 DPS) | yes | Barbaric Bracers (18948, -0.38 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.54 DPS) [dungeon]; Insignia Bracers (6410, -0.59 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.3 attack_power points (1.86 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.44 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.3 attack_power points (2.15 DPS) | yes | Skulker's Leather Belt (252520, -0.03 DPS) [crafted]; Warden's Leather Belt (252460, -0.44 DPS) [crafted]; Deftkin Belt (16659, -0.72 DPS) [quest] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 31.0 attack_power points (2.28 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Barbaric Leggings (5963, -0.28 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.1 attack_power points (1.55 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Stomping Boots (3741, -0.41 DPS) [quest]; Insignia Boots (4055, -0.44 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.7 attack_power points (1.59 DPS) | yes | Ring of Precision (1491, -0.76 DPS) [dungeon]; Tiger Band (6749, -0.86 DPS) [quest]; Bonebinder's Signet (273032, -0.86 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 15.5 attack_power points (1.14 DPS) | yes | Ring of Precision (1491, -0.40 DPS, sim-verified) [dungeon]; Tiger Band (6749, -0.41 DPS) [quest]; Bonebinder's Signet (273032, -0.41 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (101.7 DPS) | yes | Cobalt Crusher (7730, -3.14 DPS) [dungeon]; Gnarled Ash Staff (791, -6.52 DPS) [world_drop]; Viscous Hammer (13045, -16.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 143.3. Weights run: 5.2s. Verify run: 3.2s. 521 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.952 ± 0.028, crit=0.951 ± 0.034 per rating point (14 rating = 1%, 13.311 per %), hit=0.788 ± 0.072 per rating point (10 rating = 1%, 7.881 per %), melee_haste=4.266 ± 0.914

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 43.5 attack_power points (3.70 DPS) | yes | Warden's Wizard Hat (14604, -1.54 DPS) [world_drop]; Hawkeye's Helm (14591, -1.55 DPS, sim-verified) [world_drop]; Defender's Leather Helm (252455, -1.66 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 17.8 attack_power points (1.52 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.49 DPS) [world]; River Pride Choker (13087, -0.84 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.5 attack_power points (2.68 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.83 DPS) [dungeon]; Nightscape Shoulders (8192, -0.85 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.7 attack_power points (1.67 DPS) | yes | Imperial Cloak (6432, -0.34 DPS) [dungeon]; Parachute Cloak (10518, -0.34 DPS) [crafted]; Tigerstrike Mantle (13108, -0.34 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 33.2 attack_power points (2.83 DPS) | yes | Barbaric Harness (5739, -0.30 DPS) [crafted]; Nightscape Tunic (8175, -0.33 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.47 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.70 DPS) | yes | Hawkeye's Bracers (14590, -0.02 DPS) [world_drop]; Barbaric Bracers (18948, -0.36 DPS) [crafted]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 37.6 attack_power points (3.20 DPS) | yes | Skulker's Leather Gloves (252525, -0.00 DPS) [crafted]; Imperial Leather Gloves (4063, -0.17 DPS) [dungeon]; Gloves of Holy Might (867, -0.36 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 35.6 attack_power points (3.03 DPS) | yes | Prowler's Leather Belt (252459, -0.50 DPS) [crafted]; Skulker's Leather Belt (252520, -0.51 DPS) [crafted]; Scorpashi Sash (14652, -0.69 DPS) [world_drop] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 41.1 attack_power points (3.50 DPS) | yes | Basilisk Hide Pants (1718, -0.01 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.81 DPS) [crafted]; Brawler's Leather Pants (252500, -0.99 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 35.7 attack_power points (3.04 DPS) | yes | Excelsior Boots (4109, -0.35 DPS) [quest]; Imperial Leather Boots (6431, -0.36 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.56 DPS, sim-verified) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.6 attack_power points (2.01 DPS) | yes | Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Mark of Kern (2262, -0.30 DPS) [dungeon]; Assault Band (13095, -0.30 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.5 attack_power points (2.00 DPS) | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Mark of Kern (2262, -0.30 DPS) [dungeon]; Assault Band (13095, -0.30 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (143.3 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Staff of Jordan (873, -10.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Falcon's Hook; finger2: Ring of the Underwood

No-known-source sample (15 of 521, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0000000000000000-50212032123012123001-5500000000000000)

Set DPS (verified): 193.4. Weights run: 5.1s. Verify run: 3.8s. 691 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.941 ± 0.030, crit=1.123 ± 0.038 per rating point (14 rating = 1%, 15.723 per %), hit=0.998 ± 0.091 per rating point (10 rating = 1%, 9.976 per %), melee_haste=not significant (3.893 ± 1.141)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.5 attack_power points (3.88 DPS) | yes | Undercity Reservist's Cap (20643, -0.19 DPS) [quest]; Blood Guard's Leather Headband (220851, -0.33 DPS) [vendor]; Scorpashi Skullcap (14658, -1.00 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.2 attack_power points (3.08 DPS) | yes | Woven Ivy Necklace (19159, -0.47 DPS) [quest]; Ethereal Talisman (4430, -1.49 DPS) [quest]; Kaleidoscope Chain (13084, -1.67 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 41.7 attack_power points (3.55 DPS) | yes | Failed Flying Experiment (9647, -0.02 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.04 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.68 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 29.1 attack_power points (2.48 DPS) | yes | Blackflame Cape (13109, -0.50 DPS) [world_drop]; Serpentskin Cloak (8259, -0.66 DPS) [world_drop]; Pridelord Cape (14673, -0.70 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 59.1 attack_power points (5.03 DPS) | yes | Blazewind Breastplate (11193, -0.68 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -1.31 DPS) [vendor]; Charred Leather Tunic (19127, -2.06 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.9 attack_power points (3.23 DPS) | yes | Prowler's Leather Bracers (252539, -0.39 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.43 DPS) [crafted]; Pridelord Bands (14672, -0.64 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 48.3 attack_power points (4.11 DPS) | yes | Skulker's Leather Gauntlets (252548, +0.00 DPS, sim-verified) [crafted]; Serpentskin Gloves (8260, -0.23 DPS) [world_drop]; Heraldic Gloves (8121, -0.58 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 54.1 attack_power points (4.60 DPS) | yes | Skulker's Leather Waistguard (252474, -0.04 DPS) [crafted]; Ogron's Sash (13117, -1.43 DPS) [world_drop]; Pridelord Girdle (14674, -1.87 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 59.7 attack_power points (5.08 DPS) | yes | Stone Guard's Leather Pants (220859, -0.95 DPS, sim-verified) [vendor]; Triprunner Dungarees (9624, -1.54 DPS) [quest]; Basilisk Hide Pants (1718, -1.61 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.9 attack_power points (4.25 DPS) | yes | Skulker's Leather Boots (252469, -0.04 DPS) [crafted]; Albino Crocscale Boots (17728, -0.95 DPS) [dungeon]; Prowler's Leather Shoes (252465, -1.04 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 27.2 attack_power points (2.31 DPS) | yes | White Bone Band (11862, -0.27 DPS) [quest]; Ring of the Underwood (2951, -0.29 DPS) [world_drop]; Thunderbrow Ring (13097, -0.32 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.1 attack_power points (2.05 DPS) | yes | White Bone Band (11862, -0.01 DPS) [quest]; Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (193.4 DPS) | yes | Molten Heart of the Mountain (249470, -4.87 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (193.4 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (193.4 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -17.33 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0000000000000000-50212032123012123001-5550230000000000)

Set DPS (verified): 279.7. Weights run: 5.1s. Verify run: 9.9s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.351 ± 0.047, crit=1.930 ± 0.067 per rating point (14 rating = 1%, 27.024 per %), hit=1.980 ± 0.144 per rating point (10 rating = 1%, 19.801 per %), melee_haste=not significant (6.636 ± 1.940)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 89.9 attack_power points (8.54 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, +0.00 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.6 attack_power points (3.95 DPS) | yes | Woven Ivy Necklace (19159, -0.68 DPS) [quest]; Sunprism Pendant (20649, -1.27 DPS) [quest]; Ethereal Talisman (4430, -2.01 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (279.7 DPS) | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS) [pvp]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -5.94 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 55.3 attack_power points (5.25 DPS) | yes | Dark Phantom Cape (13122, -1.33 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.57 DPS) [world_drop]; Pridelord Cape (14673, -2.61 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 141.9 attack_power points (13.49 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, -1.73 DPS) [vendor]; Dawn Armor (252483, -2.05 DPS, sim-verified) [crafted]; Warlord's Dragonhide Hauberk (16549, -2.57 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (279.7 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -6.35 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (279.7 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -5.23 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (279.7 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -5.23 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 103.9 attack_power points (9.87 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -0.24 DPS) [pvp]; Traveler's Leggings (8300, -2.86 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 55.3 attack_power points (5.26 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -7.21 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.9 attack_power points (3.13 DPS) | yes | Falcon's Hook (7552, -0.49 DPS) [dungeon]; Thunderbrow Ring (13097, -0.79 DPS) [world_drop]; White Bone Band (11862, -0.85 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.9 attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.31 DPS) [world_drop]; White Bone Band (11862, -0.37 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (279.7 DPS) | yes | Molten Heart of the Mountain (249470, -5.76 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (279.7 DPS) | yes | Molten Heart of the Mountain (249470, -0.94 DPS, sim-verified) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (279.7 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -7.44 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (279.7 DPS) | yes | Idol of the Moon (23197, -8.66 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0500200000000000-50230032023011125001-5050030000000000)

Set DPS (verified): 660.1. Weights run: 4.9s. Verify run: 9.8s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.859 ± 0.061, crit=2.286 ± 0.083 per rating point (14 rating = 1%, 31.998 per %), hit=2.077 ± 0.194 per rating point (10 rating = 1%, 20.774 per %), melee_haste=11.317 ± 2.487

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 108.0 attack_power points (18.89 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -2.93 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 48.2 attack_power points (8.42 DPS) | yes | Woven Ivy Necklace (19159, -2.04 DPS, sim-verified) [quest]; Sunprism Pendant (20649, -2.42 DPS) [quest]; Ethereal Talisman (4430, -4.50 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (660.1 DPS) | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS) [pvp]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -17.05 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.9 attack_power points (11.00 DPS) | yes | Dark Phantom Cape (13122, -3.59 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -5.00 DPS) [world_drop]; Pridelord Cape (14673, -5.34 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 158.6 attack_power points (27.73 DPS) | yes | Timbermaw Tunic (252484, -0.44 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -3.81 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -5.52 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (660.1 DPS) | yes | Skulker's Leather Bracers (252540, -0.31 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.54 DPS) [crafted]; Deepfury Bracers (13120, -18.63 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (660.1 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -15.26 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (660.1 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -15.26 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 113.1 attack_power points (19.77 DPS) | yes | Legionnaire's Dragonhide Leggings (227177, +0.00 DPS) [pvp]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Traveler's Leggings (8300, -5.88 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 66.3 attack_power points (11.59 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -19.05 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 40.0 attack_power points (7.00 DPS) | yes | Falcon's Hook (7552, -1.35 DPS) [dungeon]; Thunderbrow Ring (13097, -2.42 DPS) [world_drop]; White Bone Band (11862, -2.80 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 33.0 attack_power points (5.77 DPS) | yes | Falcon's Hook (7552, -0.12 DPS) [dungeon]; Thunderbrow Ring (13097, -1.19 DPS) [world_drop]; White Bone Band (11862, -1.57 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (660.1 DPS) | yes | Frozen Heart of the Mountain (249469, -5.95 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (660.1 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (660.1 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -12.30 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (660.1 DPS) | yes | Idol of the Moon (23197, -14.33 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; trinket2: Molten Heart of the Mountain; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

