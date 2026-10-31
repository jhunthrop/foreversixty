# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 66.1. Weights run: 4.4s. Verify run: 5.0s. 196 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.098 ± 0.013, crit=0.679 ± 0.012 per rating point (14 rating = 1%, 9.508 per %), hit=0.913 ± 0.030 per rating point (10 rating = 1%, 9.134 per %), melee_haste=5.966 ± 0.364

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 17.6 attack_power points (0.96 DPS) | yes | Brawler's Leather Hood (252504, -0.04 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.2 attack_power points (0.23 DPS) | yes | Tarnished Locket (279870, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.5 attack_power points (0.57 DPS) | yes | Slime-encrusted Pads (6461, -0.64 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.6 attack_power points (0.68 DPS) | yes | Grave Shroud (279865, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted]; Cape of the Brotherhood (5193, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (66.1 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.07 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 12.8 attack_power points (0.69 DPS) | yes | Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Forest Leather Bracers (3202, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.24 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.1 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.19 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.98 DPS) | yes | Deviate Scale Belt (6468, -0.05 DPS) [crafted]; Murloc Scale Belt (5780, -0.40 DPS) [crafted]; Brawler's Leather Belt (252428, -3.72 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.1 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.32 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.4 attack_power points (1.27 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.59 DPS) [dungeon]; Footpads of the Fang (10411, -0.59 DPS) [dungeon] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.8 attack_power points (0.48 DPS) | yes | Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Ring of the Moon (12052, -0.36 DPS) [world_drop] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 4.3 attack_power points (0.23 DPS) | yes | Lavishly Jeweled Ring (1156, -0.01 DPS) [dungeon]; Ring of the Moon (12052, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 308.5 attack_power points (16.75 DPS) | yes | Staff of Westfall (2042, -1.17 DPS) [quest]; Twisted Chanter's Staff (890, -1.21 DPS) [world_drop]; Living Root (6631, -1.88 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: The 1 Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 121.0. Weights run: 4.6s. Verify run: 3.0s. 380 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.984 ± 0.015, crit=0.780 ± 0.016 per rating point (14 rating = 1%, 10.922 per %), hit=0.963 ± 0.032 per rating point (10 rating = 1%, 9.634 per %), melee_haste=5.524 ± 0.453

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.72 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.43 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.43 DPS) [crafted]; Brawler's Leather Helm (252512, -0.46 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.7 attack_power points (1.09 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.40 DPS, sim-verified) [world]; River Pride Choker (13087, -0.52 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.66 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.8 attack_power points (2.13 DPS) | yes | Mantle of Thieves (2264, -0.84 DPS) [dungeon]; Bristlebark Amice (14573, -0.93 DPS) [world_drop]; Barbaric Shoulders (5964, -1.16 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.5 attack_power points (1.33 DPS) | yes | Sergeant Major's Cape (16315, -0.24 DPS) [pvp]; Tigerstrike Mantle (13108, -0.30 DPS) [world_drop]; Fenrus' Hide (6340, -0.56 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 29.1 attack_power points (1.89 DPS) | yes | Dusky Leather Armor (7374, -0.08 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS) [crafted]; Brawler's Leather Armor (252490, -0.27 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.7 attack_power points (1.34 DPS) | yes | Barbaric Bracers (18948, -0.38 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.40 DPS) [quest]; Bands of Serra'kis (6902, -0.49 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.3 attack_power points (1.77 DPS) | yes | Toughened Leather Gloves (4253, -0.14 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.43 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 31.7 attack_power points (2.06 DPS) | yes | Skulker's Leather Belt (252520, -0.04 DPS) [crafted]; Warden's Leather Belt (252460, -0.43 DPS) [crafted]; Skirmisher's Leather Belt (252461, -0.77 DPS) [crafted] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 33.3 attack_power points (2.16 DPS) | yes | Brawler's Leather Pants (252500, -0.14 DPS) [crafted]; Trapper's Leather Pants (252501, -0.14 DPS) [crafted]; Barbaric Leggings (5963, -0.26 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.5 attack_power points (1.46 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Insignia Boots (4055, -0.43 DPS) [world_drop]; Disjointed Shoes (277226, -0.68 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.6 attack_power points (1.53 DPS) | yes | Ring of Precision (1491, -0.76 DPS) [dungeon]; Tiger Band (6749, -0.82 DPS) [quest]; Bonebinder's Signet (273032, -0.82 DPS) [world] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 12.6 attack_power points (0.82 DPS) | yes | Ring of Precision (1491, -0.04 DPS) [dungeon]; Tiger Band (6749, -0.10 DPS) [quest]; Bonebinder's Signet (273032, -0.10 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (121.0 DPS) | yes | Cobalt Crusher (7730, -3.01 DPS) [dungeon]; Gnarled Ash Staff (791, -6.01 DPS) [world_drop]; Viscous Hammer (13045, -23.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Seal of Wrynn; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 165.4. Weights run: 5.1s. Verify run: 3.2s. 516 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.938 ± 0.021, crit=0.933 ± 0.025 per rating point (14 rating = 1%, 13.058 per %), hit=0.951 ± 0.039 per rating point (10 rating = 1%, 9.512 per %), melee_haste=5.776 ± 0.555

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.5 attack_power points (3.38 DPS) | yes | Defender's Leather Helm (252455, -1.42 DPS) [crafted]; Warden's Wizard Hat (14604, -1.51 DPS) [world_drop]; Hawkeye's Helm (14591, -1.55 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.6 attack_power points (1.23 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.48 DPS, sim-verified) [world]; River Pride Choker (13087, -0.58 DPS) [world_drop]; Gazlowe's Charm (13088, -0.58 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.3 attack_power points (2.40 DPS) | yes | Flintrock Shoulders (7755, -0.14 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.72 DPS) [dungeon]; Nightscape Shoulders (8192, -0.82 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 24.8 attack_power points (1.84 DPS) | yes | Yeti Fur Cloak (2805, -0.49 DPS) [quest]; Parachute Cloak (10518, -0.69 DPS) [crafted]; Hawkeye's Cloak (14593, -1.22 DPS, sim-verified) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 33.0 attack_power points (2.45 DPS) | yes | Nightscape Tunic (8175, -0.29 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.32 DPS) [crafted]; Barbaric Harness (5739, -0.74 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (1.52 DPS) | yes | Branded Leather Bracers (19508, -0.03 DPS) [dungeon]; Barbaric Bracers (18948, -0.29 DPS) [crafted]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 39.4 attack_power points (2.93 DPS) | yes | Skulker's Leather Gloves (252525, -0.02 DPS) [crafted]; Imperial Leather Gloves (4063, -0.16 DPS) [dungeon]; Gloves of Holy Might (867, -0.47 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.2 attack_power points (2.77 DPS) | yes | Prowler's Leather Belt (252459, -0.43 DPS) [crafted]; Skulker's Leather Belt (252520, -0.49 DPS) [crafted]; Scorpashi Sash (14652, -0.67 DPS) [world_drop] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 41.5 attack_power points (3.08 DPS) | yes | Basilisk Hide Pants (1718, -0.06 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.64 DPS) [crafted]; Brawler's Leather Pants (252500, -0.81 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 37.8 attack_power points (2.81 DPS) | yes | Skulker's Leather Shoes (252531, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.37 DPS) [quest]; Imperial Leather Boots (6431, -0.40 DPS) [dungeon] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.0 attack_power points (1.79 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Mark of Kern (2262, -0.30 DPS) [dungeon]; Assault Band (13095, -0.30 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.8 attack_power points (1.77 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Mark of Kern (2262, -0.28 DPS) [dungeon]; Assault Band (13095, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (165.4 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Gut Ripper (2164, -31.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Falcon's Hook; finger2: Ring of the Underwood

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 213.5. Weights run: 5.1s. Verify run: 3.8s. 686 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.863 ± 0.025, crit=1.029 ± 0.032 per rating point (14 rating = 1%, 14.413 per %), hit=1.063 ± 0.049 per rating point (10 rating = 1%, 10.627 per %), melee_haste=6.044 ± 0.770

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.7 attack_power points (3.59 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.29 DPS) [vendor]; Helm of Fire (8348, -1.05 DPS) [crafted]; Scorpashi Skullcap (14658, -1.34 DPS, sim-verified) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.2 attack_power points (2.83 DPS) | yes | Kaleidoscope Chain (13084, -1.74 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -1.87 DPS) [world]; River Pride Choker (13087, -2.12 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 41.0 attack_power points (3.29 DPS) | yes | Failed Flying Experiment (9647, -0.03 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.05 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.64 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 27.9 attack_power points (2.25 DPS) | yes | Sergeant Major's Cape (16336, -0.29 DPS) [pvp]; Pridelord Cape (14673, -0.37 DPS) [world_drop]; Blackflame Cape (13109, -0.45 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 57.7 attack_power points (4.64 DPS) | yes | Blazewind Breastplate (11193, -0.96 DPS, sim-verified) [quest]; Knight's Leather Armor (220854, -1.18 DPS) [vendor]; Charred Leather Tunic (19127, -1.95 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 36.7 attack_power points (2.95 DPS) | yes | Prowler's Leather Bracers (252539, -0.31 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.37 DPS) [crafted]; Pridelord Bands (14672, -0.57 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 47.6 attack_power points (3.82 DPS) | yes | Serpentskin Gloves (8260, -0.23 DPS) [world_drop]; Heraldic Gloves (8121, -0.56 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.83 DPS, sim-verified) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 53.2 attack_power points (4.27 DPS) | yes | Skulker's Leather Waistguard (252474, -0.05 DPS) [crafted]; Ogron's Sash (13117, -1.33 DPS) [world_drop]; Pridelord Girdle (14674, -1.77 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 60.9 attack_power points (4.90 DPS) | yes | Serpentskin Leggings (8262, -0.20 DPS) [world_drop]; Knight's Leather Pants (220858, -1.44 DPS) [vendor]; Triprunner Dungarees (9624, -1.67 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.1 attack_power points (3.95 DPS) | yes | Skulker's Leather Boots (252469, -0.05 DPS) [crafted]; Albino Crocscale Boots (17728, -0.95 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.95 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.1 attack_power points (2.10 DPS) | yes | Thunderbrow Ring (13097, -0.23 DPS) [world_drop]; Ring of the Underwood (2951, -0.25 DPS) [world_drop]; Blackstone Ring (17713, -0.49 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.4 attack_power points (1.88 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Ring of the Underwood (2951, -0.03 DPS) [world_drop]; Blackstone Ring (17713, -0.27 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (213.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (213.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (213.5 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -35.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 686, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 252.3. Weights run: 5.0s. Verify run: 9.9s. 1564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.789 ± 0.028, crit=1.127 ± 0.040 per rating point (14 rating = 1%, 15.772 per %), hit=1.129 ± 0.063 per rating point (10 rating = 1%, 11.291 per %), melee_haste=5.837 ± 1.062

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadowcraft Cap (16707) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (252.3 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -5.57 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.3 attack_power points (2.93 DPS) | yes | Sunprism Pendant (20649, -1.23 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.90 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (252.3 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -5.54 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.8 attack_power points (4.00 DPS) | yes | Sergeant Major's Cape (16337, -1.27 DPS) [pvp]; Dark Phantom Cape (13122, -1.71 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 120.6 attack_power points (10.30 DPS) | yes | Dawn Armor (252483, -1.95 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -2.09 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -2.11 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (252.3 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Blue Suede Bracers (252559, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -5.77 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (252.3 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -4.36 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 71.2 attack_power points (6.08 DPS) | yes | Skulker's Leather Waistguard (252474, -1.69 DPS) [crafted]; Prowler's Leather Waistguard (252473, -1.91 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -2.25 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 93.8 attack_power points (8.01 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.61 DPS) [vendor]; Traveler's Leggings (8300, -2.52 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (252.3 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Prowler's Leather Boots (252468, -6.01 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 25.0 attack_power points (2.14 DPS) | yes | Falcon's Hook (7552, -0.20 DPS) [dungeon]; Ring of the Underwood (2951, -0.24 DPS) [world_drop]; Ring of Living Stone (18400, -0.26 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.0 attack_power points (1.96 DPS) | yes | Falcon's Hook (7552, -0.02 DPS) [dungeon]; Ring of the Underwood (2951, -0.06 DPS) [world_drop]; Ring of Living Stone (18400, -0.08 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (252.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (252.3 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (252.3 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -0.92 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (252.3 DPS) | yes | Idol of the Moon (23197, -8.90 DPS, sim-verified) [world_drop] |

**New at 60:** head: Shadowcraft Cap; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Thunderbrow Ring; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 0500200000000000-45211031021032212001-5053000000000000)

Set DPS (verified): 596.0. Weights run: 4.9s. Verify run: 9.7s. 1564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.289 ± 0.035, crit=1.554 ± 0.047 per rating point (14 rating = 1%, 21.756 per %), hit=1.821 ± 0.084 per rating point (10 rating = 1%, 18.214 per %), melee_haste=12.856 ± 1.295

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadowcraft Cap (16707) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (596.0 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -11.16 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.9 attack_power points (6.42 DPS) | yes | Sunprism Pendant (20649, -2.83 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -3.53 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -4.58 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (596.0 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -10.23 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.3 attack_power points (8.33 DPS) | yes | Sergeant Major's Cape (16337, -2.55 DPS) [pvp]; Dark Phantom Cape (13122, -3.07 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 142.8 attack_power points (21.90 DPS) | yes | Dawn Armor (252483, -3.22 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -3.43 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -4.37 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (596.0 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -10.63 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (596.0 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -8.53 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 82.9 attack_power points (12.71 DPS) | yes | Skulker's Leather Waistguard (252474, -3.34 DPS) [crafted]; Prowler's Leather Waistguard (252473, -3.78 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -4.45 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 109.0 attack_power points (16.71 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.91 DPS) [vendor]; Traveler's Leggings (8300, -4.90 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (596.0 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Prowler's Leather Boots (252468, -11.77 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.0 attack_power points (4.91 DPS) | yes | Ring of the Underwood (2951, -0.66 DPS) [world_drop]; Thunderbrow Ring (13097, -0.89 DPS) [world_drop]; Ring of Living Stone (18400, -1.20 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 27.9 attack_power points (4.27 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.25 DPS) [world_drop]; Ring of Living Stone (18400, -0.56 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (596.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (596.0 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (596.0 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -5.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (596.0 DPS) | yes | Idol of the Moon (23197, -6.43 DPS, sim-verified) [world_drop] |

**New at 60:** head: Shadowcraft Cap; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 65.0. Weights run: 4.4s. Verify run: 4.9s. 192 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.098 ± 0.013, crit=0.679 ± 0.012 per rating point (14 rating = 1%, 9.508 per %), hit=0.913 ± 0.030 per rating point (10 rating = 1%, 9.134 per %), melee_haste=5.966 ± 0.364

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 17.6 attack_power points (0.96 DPS) | yes | Brawler's Leather Hood (252504, -0.04 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.2 attack_power points (0.23 DPS) | yes | Roadwatcher's Confidence (281265, -0.27 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.5 attack_power points (0.57 DPS) | yes | Slime-encrusted Pads (6461, -0.64 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.6 attack_power points (0.68 DPS) | yes | Grave Shroud (279865, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted]; Cape of the Brotherhood (5193, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (65.0 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.07 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 10.7 attack_power points (0.58 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Spare Part Bindings (279875, -0.24 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (65.0 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.18 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.98 DPS) | yes | Deviate Scale Belt (6468, -0.05 DPS) [crafted]; Murloc Scale Belt (5780, -0.40 DPS) [crafted]; Brawler's Leather Belt (252428, -3.70 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (65.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.32 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.4 attack_power points (1.27 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.59 DPS) [dungeon]; Footpads of the Fang (10411, -0.59 DPS) [dungeon] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.8 attack_power points (0.48 DPS) | yes | Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; The 1 Ring (8350, -0.24 DPS) [world]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.6 attack_power points (0.36 DPS) | yes | Bounty Hunter's Ring (5351, -0.02 DPS) [quest]; The 1 Ring (8350, -0.12 DPS) [world]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 308.5 attack_power points (16.75 DPS) | yes | Crescent Staff (6505, -0.49 DPS) [quest]; Living Root (6631, -0.63 DPS) [dungeon]; Hammerbone (270018, -0.63 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 118.9. Weights run: 4.6s. Verify run: 3.0s. 380 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.984 ± 0.015, crit=0.780 ± 0.016 per rating point (14 rating = 1%, 10.922 per %), hit=0.963 ± 0.032 per rating point (10 rating = 1%, 9.634 per %), melee_haste=5.524 ± 0.453

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.72 DPS) | yes | Brawler's Leather Helm (252512, -0.41 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.43 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.43 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.7 attack_power points (1.09 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.41 DPS, sim-verified) [world]; River Pride Choker (13087, -0.52 DPS) [world_drop]; Erudite's Amulet (277204, -0.83 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.8 attack_power points (2.13 DPS) | yes | Mantle of Thieves (2264, -0.84 DPS) [dungeon]; Bristlebark Amice (14573, -0.93 DPS) [world_drop]; Barbaric Shoulders (5964, -1.17 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.5 attack_power points (1.33 DPS) | yes | Tigerstrike Mantle (13108, -0.37 DPS, sim-verified) [world_drop]; Fenrus' Hide (6340, -0.56 DPS) [dungeon]; Swiftrunner Cape (6745, -0.56 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 29.1 attack_power points (1.89 DPS) | yes | Dusky Leather Armor (7374, -0.08 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS) [crafted]; Brawler's Leather Armor (252490, -0.27 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.7 attack_power points (1.34 DPS) | yes | Barbaric Bracers (18948, -0.39 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.49 DPS) [dungeon]; Insignia Bracers (6410, -0.57 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.3 attack_power points (1.77 DPS) | yes | Toughened Leather Gloves (4253, -0.14 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.43 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 31.7 attack_power points (2.06 DPS) | yes | Skulker's Leather Belt (252520, -0.04 DPS) [crafted]; Warden's Leather Belt (252460, -0.43 DPS) [crafted]; Deftkin Belt (16659, -0.76 DPS) [quest] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 33.3 attack_power points (2.16 DPS) | yes | Brawler's Leather Pants (252500, -0.14 DPS) [crafted]; Trapper's Leather Pants (252501, -0.14 DPS) [crafted]; Barbaric Leggings (5963, -0.26 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.5 attack_power points (1.46 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Stomping Boots (3741, -0.37 DPS) [quest]; Insignia Boots (4055, -0.43 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.6 attack_power points (1.53 DPS) | yes | Ring of Precision (1491, -0.76 DPS) [dungeon]; Tiger Band (6749, -0.82 DPS) [quest]; Bonebinder's Signet (273032, -0.82 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 16.7 attack_power points (1.09 DPS) | yes | Ring of Precision (1491, -0.37 DPS, sim-verified) [dungeon]; Tiger Band (6749, -0.37 DPS) [quest]; Bonebinder's Signet (273032, -0.37 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (118.9 DPS) | yes | Cobalt Crusher (7730, -3.01 DPS) [dungeon]; Gnarled Ash Staff (791, -6.01 DPS) [world_drop]; Viscous Hammer (13045, -23.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 161.8. Weights run: 5.1s. Verify run: 3.1s. 514 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.938 ± 0.021, crit=0.933 ± 0.025 per rating point (14 rating = 1%, 13.058 per %), hit=0.951 ± 0.039 per rating point (10 rating = 1%, 9.512 per %), melee_haste=5.776 ± 0.555

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.5 attack_power points (3.38 DPS) | yes | Defender's Leather Helm (252455, -1.42 DPS) [crafted]; Warden's Wizard Hat (14604, -1.51 DPS) [world_drop]; Hawkeye's Helm (14591, -1.51 DPS, sim-verified) [world_drop] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 18.8 attack_power points (1.39 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.50 DPS) [world]; River Pride Choker (13087, -0.74 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.3 attack_power points (2.40 DPS) | yes | Flintrock Shoulders (7755, -0.14 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.72 DPS) [dungeon]; Nightscape Shoulders (8192, -0.82 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.2 attack_power points (1.50 DPS) | yes | Imperial Cloak (6432, -0.35 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.35 DPS) [world_drop]; Parachute Cloak (10518, -0.48 DPS, sim-verified) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 33.0 attack_power points (2.45 DPS) | yes | Nightscape Tunic (8175, -0.29 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.32 DPS) [crafted]; Barbaric Harness (5739, -0.81 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (1.52 DPS) | yes | Branded Leather Bracers (19508, -0.03 DPS) [dungeon]; Barbaric Bracers (18948, -0.29 DPS) [crafted]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 39.4 attack_power points (2.93 DPS) | yes | Skulker's Leather Gloves (252525, -0.02 DPS) [crafted]; Imperial Leather Gloves (4063, -0.16 DPS) [dungeon]; Gloves of Holy Might (867, -0.47 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.2 attack_power points (2.77 DPS) | yes | Prowler's Leather Belt (252459, -0.46 DPS, sim-verified) [crafted]; Skulker's Leather Belt (252520, -0.49 DPS) [crafted]; Tharg's Shoelace (9705, -0.64 DPS) [quest] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 41.5 attack_power points (3.08 DPS) | yes | Basilisk Hide Pants (1718, -0.06 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.64 DPS) [crafted]; Brawler's Leather Pants (252500, -0.81 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 37.8 attack_power points (2.81 DPS) | yes | Excelsior Boots (4109, -0.37 DPS) [quest]; Imperial Leather Boots (6431, -0.40 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.74 DPS, sim-verified) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.0 attack_power points (1.79 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Mark of Kern (2262, -0.30 DPS) [dungeon]; Assault Band (13095, -0.30 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.8 attack_power points (1.77 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Mark of Kern (2262, -0.28 DPS) [dungeon]; Assault Band (13095, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (161.8 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Gut Ripper (2164, -30.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Wolffear Harness; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Falcon's Hook; finger2: Ring of the Underwood

No-known-source sample (15 of 514, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 217.9. Weights run: 5.1s. Verify run: 3.8s. 684 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.863 ± 0.025, crit=1.029 ± 0.032 per rating point (14 rating = 1%, 14.413 per %), hit=1.063 ± 0.049 per rating point (10 rating = 1%, 10.627 per %), melee_haste=6.044 ± 0.770

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.7 attack_power points (3.59 DPS) | yes | Undercity Reservist's Cap (20643, -0.18 DPS) [quest]; Blood Guard's Leather Headband (220851, -0.29 DPS) [vendor]; Scorpashi Skullcap (14658, -0.97 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.2 attack_power points (2.83 DPS) | yes | Woven Ivy Necklace (19159, -0.42 DPS) [quest]; Ethereal Talisman (4430, -1.35 DPS) [quest]; Kaleidoscope Chain (13084, -1.53 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 41.0 attack_power points (3.29 DPS) | yes | Skulker's Leather Shoulder (252535, -0.05 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.64 DPS) [vendor]; Failed Flying Experiment (9647, -0.79 DPS, sim-verified) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 27.9 attack_power points (2.25 DPS) | yes | Blackflame Cape (13109, -0.45 DPS) [world_drop]; Serpentskin Cloak (8259, -0.60 DPS) [world_drop]; Pridelord Cape (14673, -0.64 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 57.7 attack_power points (4.64 DPS) | yes | Blazewind Breastplate (11193, -0.75 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -1.18 DPS) [vendor]; Charred Leather Tunic (19127, -1.95 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 36.7 attack_power points (2.95 DPS) | yes | Skulker's Leather Bracers (252540, -0.37 DPS) [crafted]; Pridelord Bands (14672, -0.57 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.62 DPS, sim-verified) [crafted] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 47.6 attack_power points (3.82 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.14 DPS) [crafted]; Serpentskin Gloves (8260, -0.23 DPS) [world_drop]; Heraldic Gloves (8121, -0.56 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 53.2 attack_power points (4.27 DPS) | yes | Skulker's Leather Waistguard (252474, -0.05 DPS) [crafted]; Ogron's Sash (13117, -1.33 DPS) [world_drop]; Pridelord Girdle (14674, -1.77 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 58.4 attack_power points (4.70 DPS) | yes | Triprunner Dungarees (9624, -1.47 DPS) [quest]; Basilisk Hide Pants (1718, -1.55 DPS) [world_drop]; Stone Guard's Leather Pants (220859, -1.57 DPS, sim-verified) [vendor] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.1 attack_power points (3.95 DPS) | yes | Skulker's Leather Boots (252469, -0.05 DPS) [crafted]; Albino Crocscale Boots (17728, -0.95 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.95 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.1 attack_power points (2.10 DPS) | yes | Falcon's Hook (7552, -0.22 DPS) [dungeon]; Thunderbrow Ring (13097, -0.23 DPS) [world_drop]; Ring of the Underwood (2951, -0.25 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.93 DPS) | yes | Falcon's Hook (7552, -0.05 DPS) [dungeon]; Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Ring of the Underwood (2951, -0.08 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (217.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (217.9 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (217.9 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -35.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 258.1. Weights run: 5.0s. Verify run: 9.8s. 1564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.789 ± 0.028, crit=1.127 ± 0.040 per rating point (14 rating = 1%, 15.772 per %), hit=1.129 ± 0.063 per rating point (10 rating = 1%, 11.291 per %), melee_haste=5.837 ± 1.062

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadowcraft Cap (16707) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (258.1 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -6.36 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.3 attack_power points (2.93 DPS) | yes | Woven Ivy Necklace (19159, -0.42 DPS) [quest]; Sunprism Pendant (20649, -1.09 DPS) [quest]; Ethereal Talisman (4430, -1.38 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (258.1 DPS) | yes | Warlord's Dragonhide Epaulets (16551, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -6.35 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.8 attack_power points (4.00 DPS) | yes | Dark Phantom Cape (13122, -1.68 DPS, sim-verified) [world_drop]; Pridelord Cape (14673, -2.06 DPS) [world_drop]; Blackflame Cape (13109, -2.17 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 120.6 attack_power points (10.30 DPS) | yes | Dawn Armor (252483, -1.70 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -2.09 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -2.11 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (258.1 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Blue Suede Bracers (252559, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -6.40 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (258.1 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -5.35 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 71.2 attack_power points (6.08 DPS) | yes | Skulker's Leather Waistguard (252474, -1.69 DPS) [crafted]; Prowler's Leather Waistguard (252473, -1.80 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -2.25 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 93.8 attack_power points (8.01 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -1.61 DPS) [pvp]; Traveler's Leggings (8300, -2.49 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (258.1 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Prowler's Leather Boots (252468, -6.94 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 25.0 attack_power points (2.14 DPS) | yes | Thunderbrow Ring (13097, -0.18 DPS) [world_drop]; Falcon's Hook (7552, -0.20 DPS) [dungeon]; Ring of the Underwood (2951, -0.24 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.05 DPS) | yes | Thunderbrow Ring (13097, -0.09 DPS) [world_drop]; Falcon's Hook (7552, -0.11 DPS) [dungeon]; Ring of the Underwood (2951, -0.15 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (258.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (258.1 DPS) | yes | Molten Heart of the Mountain (249470, -1.57 DPS, sim-verified) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (258.1 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -2.48 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (258.1 DPS) | yes | Idol of the Moon (23197, -9.12 DPS, sim-verified) [world_drop] |

**New at 60:** head: Shadowcraft Cap; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0500200000000000-45211031021032212001-5053000000000000)

Set DPS (verified): 604.0. Weights run: 4.9s. Verify run: 9.7s. 1564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.289 ± 0.035, crit=1.554 ± 0.047 per rating point (14 rating = 1%, 21.756 per %), hit=1.821 ± 0.084 per rating point (10 rating = 1%, 18.214 per %), melee_haste=12.856 ± 1.295

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadowcraft Cap (16707) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (604.0 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -11.42 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.9 attack_power points (6.42 DPS) | yes | Woven Ivy Necklace (19159, -1.03 DPS) [quest]; Sunprism Pendant (20649, -2.21 DPS) [quest]; Ethereal Talisman (4430, -3.16 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (604.0 DPS) | yes | Warlord's Dragonhide Epaulets (16551, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -9.91 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.3 attack_power points (8.33 DPS) | yes | Dark Phantom Cape (13122, -3.00 DPS, sim-verified) [world_drop]; Pridelord Cape (14673, -4.06 DPS) [world_drop]; Blackflame Cape (13109, -4.12 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 142.8 attack_power points (21.90 DPS) | yes | Dawn Armor (252483, -2.86 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -3.43 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -4.37 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (604.0 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -10.73 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (604.0 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -8.41 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 82.9 attack_power points (12.71 DPS) | yes | Skulker's Leather Waistguard (252474, -3.34 DPS) [crafted]; Prowler's Leather Waistguard (252473, -3.77 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -4.45 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 109.0 attack_power points (16.71 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -1.91 DPS) [pvp]; Traveler's Leggings (8300, -4.79 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (604.0 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Prowler's Leather Boots (252468, -11.61 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.0 attack_power points (4.91 DPS) | yes | Ring of the Underwood (2951, -0.66 DPS) [world_drop]; Thunderbrow Ring (13097, -0.89 DPS) [world_drop]; Ring of Living Stone (18400, -1.20 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 27.9 attack_power points (4.27 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.25 DPS) [world_drop]; Ring of Living Stone (18400, -0.56 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (604.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (604.0 DPS) | yes | Molten Heart of the Mountain (249470, -2.43 DPS, sim-verified) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (604.0 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -5.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (604.0 DPS) | yes | Idol of the Moon (23197, -7.17 DPS, sim-verified) [world_drop] |

**New at 60:** head: Shadowcraft Cap; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Falcon's Hook; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

