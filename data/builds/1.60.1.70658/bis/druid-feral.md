# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 70.7. Weights run: 4.4s. Verify run: 4.9s. 196 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.236 ± 0.013, crit=0.765 ± 0.012 per rating point (14 rating = 1%, 10.708 per %), hit=0.882 ± 0.037 per rating point (10 rating = 1%, 8.820 per %), melee_haste=5.316 ± 0.436

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.9 attack_power points (1.02 DPS) | yes | Defender's Leather Hood (252447, -0.02 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.5 attack_power points (0.26 DPS) | yes | Tarnished Locket (279870, -0.29 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 11.2 attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.73 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.4 attack_power points (0.77 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.13 DPS) [crafted]; Grave Shroud (279865, -0.13 DPS) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 26.7 attack_power points (1.52 DPS) | yes | Tunic of Westfall (2041, -0.12 DPS) [quest]; Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Prospector's Chestpiece (14562, -0.50 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 13.3 attack_power points (0.76 DPS) | yes | Forest Leather Bracers (3202, -0.12 DPS) [world_drop]; Bristlebark Bindings (14569, -0.13 DPS) [world_drop]; Wolf Bracers (4794, -0.25 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (70.7 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.50 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Belt (252428, -0.01 DPS) [crafted]; Dusty Belt (279897, -0.39 DPS) [quest]; Deviate Scale Belt (6468, -4.07 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (70.7 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.66 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (70.7 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.40 DPS, sim-verified) [world_drop] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.8 attack_power points (0.50 DPS) | yes | The 1 Ring (8350, -0.25 DPS) [world]; Ring of the Moon (12052, -0.38 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.5 attack_power points (0.26 DPS) | yes | Ring of the Moon (12052, -0.13 DPS) [world_drop]; The 1 Ring (8350, -0.22 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 309.1 attack_power points (17.63 DPS) | yes | Staff of Westfall (2042, -1.26 DPS) [quest]; Twisted Chanter's Staff (890, -1.31 DPS) [world_drop]; Living Root (6631, -2.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Lavishly Jeweled Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 125.1. Weights run: 4.4s. Verify run: 2.9s. 380 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.067 ± 0.014, crit=0.841 ± 0.015 per rating point (14 rating = 1%, 11.778 per %), hit=0.953 ± 0.045 per rating point (10 rating = 1%, 9.528 per %), melee_haste=5.814 ± 0.552

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.82 DPS) | yes | Brawler's Leather Helm (252512, -0.37 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.46 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.46 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.1 attack_power points (1.18 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.41 DPS, sim-verified) [world]; River Pride Choker (13087, -0.57 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.72 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.7 attack_power points (2.33 DPS) | yes | Mantle of Thieves (2264, -0.90 DPS) [dungeon]; Bristlebark Amice (14573, -1.02 DPS) [world_drop]; Barbaric Shoulders (5964, -1.23 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 21.1 attack_power points (1.45 DPS) | yes | Sergeant Major's Cape (16315, -0.28 DPS) [pvp]; Tigerstrike Mantle (13108, -0.31 DPS) [world_drop]; Fenrus' Hide (6340, -0.60 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 29.7 attack_power points (2.05 DPS) | yes | Dusky Leather Armor (7374, -0.06 DPS) [crafted]; Defender's Leather Tunic (252450, -0.29 DPS) [crafted]; Brawler's Leather Armor (252490, -0.29 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.2 attack_power points (1.46 DPS) | yes | Barbaric Bracers (18948, -0.40 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.44 DPS) [quest]; Bands of Serra'kis (6902, -0.55 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.8 attack_power points (1.92 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.30 DPS) [dungeon]; Gloves of the Fang (10413, -0.46 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 32.2 attack_power points (2.22 DPS) | yes | Skulker's Leather Belt (252520, -0.03 DPS) [crafted]; Warden's Leather Belt (252460, -0.46 DPS) [crafted]; Skirmisher's Leather Belt (252461, -0.86 DPS) [crafted] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 34.0 attack_power points (2.35 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Barbaric Leggings (5963, -0.29 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.1 attack_power points (1.60 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Insignia Boots (4055, -0.46 DPS) [world_drop]; Footpads of the Fang (10411, -0.74 DPS) [dungeon] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.8 attack_power points (1.64 DPS) | yes | Ring of Precision (1491, -0.79 DPS) [dungeon]; Tiger Band (6749, -0.88 DPS) [quest]; Bonebinder's Signet (273032, -0.88 DPS) [world] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 12.8 attack_power points (0.88 DPS) | yes | Ring of Precision (1491, -0.03 DPS) [dungeon]; Tiger Band (6749, -0.12 DPS) [quest]; Bonebinder's Signet (273032, -0.12 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (125.1 DPS) | yes | Cobalt Crusher (7730, -3.23 DPS) [dungeon]; Gnarled Ash Staff (791, -6.41 DPS) [world_drop]; Viscous Hammer (13045, -23.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Seal of Wrynn; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 175.0. Weights run: 4.8s. Verify run: 3.0s. 516 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.100 ± 0.028, crit=1.049 ± 0.033 per rating point (14 rating = 1%, 14.690 per %), hit=0.928 ± 0.071 per rating point (10 rating = 1%, 9.281 per %), melee_haste=5.250 ± 0.874

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 47.3 attack_power points (3.79 DPS) | yes | Warden's Wizard Hat (14604, -1.60 DPS) [world_drop]; Hawkeye's Helm (14591, -1.67 DPS, sim-verified) [world_drop]; Defender's Leather Helm (252455, -1.68 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.2 attack_power points (1.38 DPS) | yes | River Pride Choker (13087, -0.67 DPS) [world_drop]; Gazlowe's Charm (13088, -0.67 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.73 DPS, sim-verified) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.1 attack_power points (2.73 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.84 DPS) [dungeon]; Nightscape Shoulders (8192, -0.88 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 25.8 attack_power points (2.07 DPS) | yes | Yeti Fur Cloak (2805, -0.53 DPS) [quest]; Parachute Cloak (10518, -0.72 DPS) [crafted]; Hawkeye's Cloak (14593, -1.38 DPS, sim-verified) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 35.7 attack_power points (2.86 DPS) | yes | Nightscape Tunic (8175, -0.34 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.46 DPS) [crafted]; Barbaric Harness (5739, -0.62 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.4 attack_power points (1.72 DPS) | yes | Branded Leather Bracers (19508, -0.11 DPS) [dungeon]; Barbaric Bracers (18948, -0.34 DPS) [crafted]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 40.9 attack_power points (3.28 DPS) | yes | Skulker's Leather Gloves (252525, -0.01 DPS) [crafted]; Imperial Leather Gloves (4063, -0.18 DPS) [dungeon]; Gloves of Holy Might (867, -0.50 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 38.7 attack_power points (3.10 DPS) | yes | Prowler's Leather Belt (252459, -0.50 DPS) [crafted]; Skulker's Leather Belt (252520, -0.53 DPS) [crafted]; Scorpashi Sash (14652, -0.71 DPS) [world_drop] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 44.4 attack_power points (3.56 DPS) | yes | Basilisk Hide Pants (1718, -0.02 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.81 DPS) [crafted]; Brawler's Leather Pants (252500, -0.99 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 38.9 attack_power points (3.12 DPS) | yes | Skulker's Leather Shoes (252531, -0.03 DPS) [crafted]; Excelsior Boots (4109, -0.37 DPS) [quest]; Imperial Leather Boots (6431, -0.38 DPS) [dungeon] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.5 attack_power points (2.04 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Mark of Kern (2262, -0.44 DPS) [dungeon]; Assault Band (13095, -0.44 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.4 attack_power points (2.04 DPS) | yes | Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Mark of Kern (2262, -0.43 DPS) [dungeon]; Assault Band (13095, -0.43 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (175.0 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Staff of Jordan (873, -30.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Falcon's Hook; finger2: Ring of the Underwood

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 229.2. Weights run: 5.0s. Verify run: 3.7s. 686 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.943 ± 0.029, crit=1.130 ± 0.037 per rating point (14 rating = 1%, 15.826 per %), hit=1.242 ± 0.098 per rating point (10 rating = 1%, 12.415 per %), melee_haste=not significant (4.219 ± 1.249)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.6 attack_power points (3.95 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.12 DPS) [vendor]; Helm of Fire (8348, -1.09 DPS) [crafted]; Scorpashi Skullcap (14658, -1.38 DPS, sim-verified) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.3 attack_power points (3.15 DPS) | yes | Archmagister's Faceted Pendant (273049, -2.10 DPS) [world]; Kaleidoscope Chain (13084, -2.32 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -2.38 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 41.7 attack_power points (3.62 DPS) | yes | Skulker's Leather Shoulder (252535, -0.04 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.47 DPS) [vendor]; Failed Flying Experiment (9647, -1.33 DPS, sim-verified) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 29.1 attack_power points (2.53 DPS) | yes | Pridelord Cape (14673, -0.44 DPS) [world_drop]; Blackflame Cape (13109, -0.51 DPS) [world_drop]; Sergeant Major's Cape (16336, -0.65 DPS, sim-verified) [pvp] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 59.2 attack_power points (5.13 DPS) | yes | Blazewind Breastplate (11193, -0.92 DPS, sim-verified) [quest]; Knight's Leather Armor (220854, -1.12 DPS) [vendor]; Charred Leather Tunic (19127, -2.10 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.9 attack_power points (3.29 DPS) | yes | Prowler's Leather Bracers (252539, -0.39 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.44 DPS) [crafted]; Pridelord Bands (14672, -0.65 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 48.3 attack_power points (4.19 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.11 DPS) [crafted]; Serpentskin Gloves (8260, -0.24 DPS) [world_drop]; Heraldic Gloves (8121, -0.59 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 54.1 attack_power points (4.69 DPS) | yes | Skulker's Leather Waistguard (252474, -0.04 DPS) [crafted]; Ogron's Sash (13117, -1.46 DPS) [world_drop]; Pridelord Girdle (14674, -1.91 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 62.1 attack_power points (5.39 DPS) | yes | Serpentskin Leggings (8262, -0.21 DPS) [world_drop]; Knight's Leather Pants (220858, -1.38 DPS) [vendor]; Triprunner Dungarees (9624, -1.78 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 50.0 attack_power points (4.34 DPS) | yes | Skulker's Leather Boots (252469, -0.04 DPS) [crafted]; Albino Crocscale Boots (17728, -0.96 DPS) [dungeon]; Prowler's Leather Shoes (252465, -1.06 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 27.2 attack_power points (2.36 DPS) | yes | Ring of the Underwood (2951, -0.29 DPS) [world_drop]; Thunderbrow Ring (13097, -0.33 DPS) [world_drop]; Blackstone Ring (17713, -0.63 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.1 attack_power points (2.09 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Blackstone Ring (17713, -0.35 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (229.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (229.2 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (229.2 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -33.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 686, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 281.1. Weights run: 4.8s. Verify run: 7.0s. 1564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.025 ± 0.037, crit=1.464 ± 0.053 per rating point (14 rating = 1%, 20.500 per %), hit=1.310 ± 0.122 per rating point (10 rating = 1%, 13.096 per %), melee_haste=not significant (5.153 ± 1.616)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 73.3 attack_power points (6.75 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -1.33 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.3 attack_power points (3.44 DPS) | yes | Sunprism Pendant (20649, -1.51 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -1.88 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.33 DPS) [world] |
| shoulder | Dark Warder's Pauldrons (22241) | Blackrock Depths: Warder Stilgiss [dungeon] | 58.6 attack_power points (5.40 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Golden Mantle of the Dawn (19058, -1.30 DPS, sim-verified) [crafted] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.4 attack_power points (4.64 DPS) | yes | Sergeant Major's Cape (16337, -1.07 DPS, sim-verified) [pvp]; Dark Phantom Cape (13122, -1.84 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 129.5 attack_power points (11.94 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, -2.17 DPS) [vendor]; Dawn Armor (252483, -2.25 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Breastplate (16452, -2.37 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 39.2 attack_power points (3.61 DPS) | yes | Skulker's Leather Bracers (252540, -0.51 DPS) [crafted]; Blue Suede Bracers (252559, -0.62 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.92 DPS, sim-verified) [crafted] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 74.5 attack_power points (6.87 DPS) | yes | Marshal's Dragonhide Grips (231694, -0.07 DPS) [vendor]; Knight-Lieutenant's Dragonhide Grips (227183, -1.16 DPS) [pvp]; Gloves of the Pathfinder (21319, -2.60 DPS, sim-verified) [quest] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 74.5 attack_power points (6.87 DPS) | yes | Skulker's Leather Waistguard (252474, -1.82 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.37 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -2.43 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 98.0 attack_power points (9.03 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.27 DPS) [vendor]; Traveler's Leggings (8300, -2.71 DPS, sim-verified) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 50.9 attack_power points (4.69 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -0.03 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.3 attack_power points (2.61 DPS) | yes | Ring of the Underwood (2951, -0.34 DPS) [world_drop]; Thunderbrow Ring (13097, -0.43 DPS) [world_drop]; Ring of Living Stone (18400, -0.59 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.8 attack_power points (2.29 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Ring of Living Stone (18400, -0.26 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (281.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (281.1 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (281.1 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Hammer of the Northern Wind (810, -5.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (281.1 DPS) | yes | Idol of the Moon (23197, -11.80 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Dark Warder's Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; hands: Timbermaw Brawlers; waist: Might of the Timbermaw; legs: Warbear Woolies; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 0500200000000000-35201030021030212051-5053000000000000)

Set DPS (verified): 679.1. Weights run: 4.7s. Verify run: 9.5s. 1564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.434 ± 0.023, crit=1.731 ± 0.031 per rating point (14 rating = 1%, 24.234 per %), hit=2.094 ± 0.159 per rating point (10 rating = 1%, 20.942 per %), melee_haste=9.006 ± 2.016

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadowcraft Cap (16707) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (679.1 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -16.97 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 43.7 attack_power points (7.45 DPS) | yes | Sunprism Pendant (20649, -3.47 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -4.14 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -5.40 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (679.1 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -17.58 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 56.5 attack_power points (9.62 DPS) | yes | Sergeant Major's Cape (16337, -3.01 DPS) [pvp]; Dark Phantom Cape (13122, -3.41 DPS) [world_drop]; Blackflame Cape (13109, -4.65 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 147.9 attack_power points (25.18 DPS) | yes | Dawn Armor (252483, -2.16 DPS) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -3.45 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -4.95 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (679.1 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -18.71 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (679.1 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -16.62 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 84.9 attack_power points (14.45 DPS) | yes | Prowler's Leather Waistguard (252473, -3.71 DPS) [crafted]; Skulker's Leather Waistguard (252474, -4.21 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -4.94 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 111.6 attack_power points (19.00 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.39 DPS) [vendor]; Traveler's Leggings (8300, -5.59 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (679.1 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -19.62 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 34.1 attack_power points (5.80 DPS) | yes | Falcon's Hook (7552, -0.84 DPS) [dungeon]; Thunderbrow Ring (13097, -1.26 DPS) [world_drop]; Ring of Living Stone (18400, -1.68 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.2 attack_power points (4.97 DPS) | yes | Falcon's Hook (7552, -0.00 DPS) [dungeon]; Thunderbrow Ring (13097, -0.43 DPS) [world_drop]; Ring of Living Stone (18400, -0.85 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (679.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (679.1 DPS) | yes | Mark of the Chosen (17774, -3.28 DPS, sim-verified) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (679.1 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -9.11 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (679.1 DPS) | yes | Idol of the Moon (23197, -17.25 DPS, sim-verified) [world_drop] |

**New at 60:** head: Shadowcraft Cap; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 69.0. Weights run: 4.4s. Verify run: 4.9s. 192 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.236 ± 0.013, crit=0.765 ± 0.012 per rating point (14 rating = 1%, 10.708 per %), hit=0.882 ± 0.037 per rating point (10 rating = 1%, 8.820 per %), melee_haste=5.316 ± 0.436

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.9 attack_power points (1.02 DPS) | yes | Defender's Leather Hood (252447, -0.02 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.5 attack_power points (0.26 DPS) | yes | Roadwatcher's Confidence (281265, -0.29 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 11.2 attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.72 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.4 attack_power points (0.77 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.13 DPS) [crafted]; Grave Shroud (279865, -0.13 DPS) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 26.7 attack_power points (1.52 DPS) | yes | Defender's Leather Armor (252434, -0.32 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.50 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.51 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 11.2 attack_power points (0.64 DPS) | yes | Bristlebark Bindings (14569, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Spare Part Bindings (279875, -0.26 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (69.0 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.40 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Belt (252428, -0.01 DPS) [crafted]; Murloc Scale Belt (5780, -0.39 DPS) [crafted]; Deviate Scale Belt (6468, -3.96 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (69.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.59 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (69.0 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.34 DPS, sim-verified) [world_drop] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.8 attack_power points (0.50 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; The 1 Ring (8350, -0.25 DPS) [world] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 6.7 attack_power points (0.38 DPS) | yes | Loop of Sacrifice (281673, -0.01 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; The 1 Ring (8350, -0.13 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 309.1 attack_power points (17.63 DPS) | yes | Crescent Staff (6505, -0.49 DPS) [quest]; Living Root (6631, -0.69 DPS) [dungeon]; Hammerbone (270018, -0.74 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Bounty Hunter's Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 123.1. Weights run: 4.4s. Verify run: 2.8s. 380 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.067 ± 0.014, crit=0.841 ± 0.015 per rating point (14 rating = 1%, 11.778 per %), hit=0.953 ± 0.045 per rating point (10 rating = 1%, 9.528 per %), melee_haste=5.814 ± 0.552

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.82 DPS) | yes | Brawler's Leather Helm (252512, -0.36 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.46 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.46 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.1 attack_power points (1.18 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.46 DPS, sim-verified) [world]; River Pride Choker (13087, -0.57 DPS) [world_drop]; Erudite's Amulet (277204, -0.89 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.7 attack_power points (2.33 DPS) | yes | Mantle of Thieves (2264, -0.90 DPS) [dungeon]; Bristlebark Amice (14573, -1.02 DPS) [world_drop]; Barbaric Shoulders (5964, -1.27 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 21.1 attack_power points (1.45 DPS) | yes | Tigerstrike Mantle (13108, -0.40 DPS, sim-verified) [world_drop]; Fenrus' Hide (6340, -0.60 DPS) [dungeon]; Swiftrunner Cape (6745, -0.60 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 29.7 attack_power points (2.05 DPS) | yes | Dusky Leather Armor (7374, -0.06 DPS) [crafted]; Defender's Leather Tunic (252450, -0.29 DPS) [crafted]; Brawler's Leather Armor (252490, -0.29 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.2 attack_power points (1.46 DPS) | yes | Barbaric Bracers (18948, -0.41 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.55 DPS) [dungeon]; Insignia Bracers (6410, -0.61 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.8 attack_power points (1.92 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.30 DPS) [dungeon]; Gloves of the Fang (10413, -0.46 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 32.2 attack_power points (2.22 DPS) | yes | Skulker's Leather Belt (252520, -0.03 DPS) [crafted]; Warden's Leather Belt (252460, -0.46 DPS) [crafted]; Deftkin Belt (16659, -0.82 DPS) [quest] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 34.0 attack_power points (2.35 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Barbaric Leggings (5963, -0.29 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.1 attack_power points (1.60 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Stomping Boots (3741, -0.42 DPS) [quest]; Insignia Boots (4055, -0.46 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.8 attack_power points (1.64 DPS) | yes | Ring of Precision (1491, -0.79 DPS) [dungeon]; Tiger Band (6749, -0.88 DPS) [quest]; Bonebinder's Signet (273032, -0.88 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 17.1 attack_power points (1.18 DPS) | yes | Ring of Precision (1491, -0.41 DPS, sim-verified) [dungeon]; Tiger Band (6749, -0.42 DPS) [quest]; Bonebinder's Signet (273032, -0.42 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (123.1 DPS) | yes | Cobalt Crusher (7730, -3.23 DPS) [dungeon]; Gnarled Ash Staff (791, -6.41 DPS) [world_drop]; Viscous Hammer (13045, -23.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 170.3. Weights run: 4.8s. Verify run: 3.0s. 514 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.100 ± 0.028, crit=1.049 ± 0.033 per rating point (14 rating = 1%, 14.690 per %), hit=0.928 ± 0.071 per rating point (10 rating = 1%, 9.281 per %), melee_haste=5.250 ± 0.874

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 47.3 attack_power points (3.79 DPS) | yes | Warden's Wizard Hat (14604, -1.60 DPS) [world_drop]; Hawkeye's Helm (14591, -1.62 DPS, sim-verified) [world_drop]; Defender's Leather Helm (252455, -1.68 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 19.4 attack_power points (1.55 DPS) | yes | Kaleidoscope Chain (13084, -0.18 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.59 DPS) [world]; River Pride Choker (13087, -0.85 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.1 attack_power points (2.73 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.84 DPS) [dungeon]; Nightscape Shoulders (8192, -0.88 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 21.3 attack_power points (1.71 DPS) | yes | Imperial Cloak (6432, -0.36 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.36 DPS) [world_drop]; Parachute Cloak (10518, -0.53 DPS, sim-verified) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 35.7 attack_power points (2.86 DPS) | yes | Barbaric Harness (5739, -0.28 DPS) [crafted]; Nightscape Tunic (8175, -0.34 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.46 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.4 attack_power points (1.72 DPS) | yes | Branded Leather Bracers (19508, -0.11 DPS) [dungeon]; Barbaric Bracers (18948, -0.34 DPS) [crafted]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 40.9 attack_power points (3.28 DPS) | yes | Skulker's Leather Gloves (252525, -0.01 DPS) [crafted]; Imperial Leather Gloves (4063, -0.18 DPS) [dungeon]; Gloves of Holy Might (867, -0.50 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 38.7 attack_power points (3.10 DPS) | yes | Prowler's Leather Belt (252459, -0.50 DPS) [crafted]; Skulker's Leather Belt (252520, -0.53 DPS) [crafted]; Scorpashi Sash (14652, -0.71 DPS) [world_drop] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 44.4 attack_power points (3.56 DPS) | yes | Basilisk Hide Pants (1718, -0.02 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.81 DPS) [crafted]; Brawler's Leather Pants (252500, -0.99 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 38.9 attack_power points (3.12 DPS) | yes | Excelsior Boots (4109, -0.37 DPS) [quest]; Imperial Leather Boots (6431, -0.38 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.57 DPS, sim-verified) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.5 attack_power points (2.04 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Mark of Kern (2262, -0.44 DPS) [dungeon]; Assault Band (13095, -0.44 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.4 attack_power points (2.04 DPS) | yes | Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Mark of Kern (2262, -0.43 DPS) [dungeon]; Assault Band (13095, -0.43 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (170.3 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Staff of Jordan (873, -29.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Wolffear Harness; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Falcon's Hook; finger2: Ring of the Underwood

No-known-source sample (15 of 514, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 233.4. Weights run: 5.0s. Verify run: 3.7s. 684 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.943 ± 0.029, crit=1.130 ± 0.037 per rating point (14 rating = 1%, 15.826 per %), hit=1.242 ± 0.098 per rating point (10 rating = 1%, 12.415 per %), melee_haste=not significant (4.219 ± 1.249)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.6 attack_power points (3.95 DPS) | yes | Blood Guard's Leather Headband (220851, -0.12 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.19 DPS) [quest]; Scorpashi Skullcap (14658, -1.02 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.3 attack_power points (3.15 DPS) | yes | Woven Ivy Necklace (19159, -0.76 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -1.52 DPS) [quest]; Kaleidoscope Chain (13084, -1.71 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 41.7 attack_power points (3.62 DPS) | yes | Failed Flying Experiment (9647, -0.02 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.04 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.47 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 29.1 attack_power points (2.53 DPS) | yes | Pridelord Cape (14673, -0.44 DPS) [world_drop]; Blackflame Cape (13109, -0.51 DPS) [world_drop]; Serpentskin Cloak (8259, -0.67 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 59.2 attack_power points (5.13 DPS) | yes | Blazewind Breastplate (11193, -0.82 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -1.12 DPS) [vendor]; Charred Leather Tunic (19127, -2.10 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.9 attack_power points (3.29 DPS) | yes | Prowler's Leather Bracers (252539, -0.39 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.44 DPS) [crafted]; Pridelord Bands (14672, -0.65 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 48.3 attack_power points (4.19 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.11 DPS) [crafted]; Serpentskin Gloves (8260, -0.24 DPS) [world_drop]; Heraldic Gloves (8121, -0.59 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 54.1 attack_power points (4.69 DPS) | yes | Skulker's Leather Waistguard (252474, -0.04 DPS) [crafted]; Ogron's Sash (13117, -1.46 DPS) [world_drop]; Pridelord Girdle (14674, -1.91 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 59.7 attack_power points (5.18 DPS) | yes | Stone Guard's Leather Pants (220859, -1.17 DPS) [vendor]; Triprunner Dungarees (9624, -1.57 DPS) [quest]; Basilisk Hide Pants (1718, -1.64 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 50.0 attack_power points (4.34 DPS) | yes | Skulker's Leather Boots (252469, -0.04 DPS) [crafted]; Albino Crocscale Boots (17728, -0.96 DPS) [dungeon]; Prowler's Leather Shoes (252465, -1.06 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 27.2 attack_power points (2.36 DPS) | yes | White Bone Band (11862, -0.28 DPS) [quest]; Ring of the Underwood (2951, -0.29 DPS) [world_drop]; Thunderbrow Ring (13097, -0.33 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.1 attack_power points (2.09 DPS) | yes | White Bone Band (11862, -0.01 DPS) [quest]; Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.06 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (233.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (233.4 DPS) | yes | Molten Heart of the Mountain (249470, -1.25 DPS, sim-verified) [crafted] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (233.4 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -34.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 287.6. Weights run: 4.8s. Verify run: 9.6s. 1564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.025 ± 0.037, crit=1.464 ± 0.053 per rating point (14 rating = 1%, 20.500 per %), hit=1.310 ± 0.122 per rating point (10 rating = 1%, 13.096 per %), melee_haste=not significant (5.153 ± 1.616)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadowcraft Cap (16707) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (287.6 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -4.76 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.3 attack_power points (3.44 DPS) | yes | Woven Ivy Necklace (19159, -0.54 DPS) [quest]; Sunprism Pendant (20649, -1.20 DPS) [quest]; Ethereal Talisman (4430, -1.68 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (287.6 DPS) | yes | Warlord's Dragonhide Epaulets (16551, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -6.72 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.4 attack_power points (4.64 DPS) | yes | Dark Phantom Cape (13122, -1.71 DPS, sim-verified) [world_drop]; Pridelord Cape (14673, -2.35 DPS) [world_drop]; Blackflame Cape (13109, -2.40 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 129.5 attack_power points (11.94 DPS) | yes | Dawn Armor (252483, -1.96 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -2.17 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -2.37 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (287.6 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -6.86 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (287.6 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -5.74 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 74.5 attack_power points (6.87 DPS) | yes | Skulker's Leather Waistguard (252474, -1.82 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.02 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -2.43 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 98.0 attack_power points (9.03 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -1.27 DPS) [pvp]; Traveler's Leggings (8300, -2.73 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (287.6 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Prowler's Leather Boots (252468, -7.51 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.3 attack_power points (2.61 DPS) | yes | Ring of the Underwood (2951, -0.34 DPS) [world_drop]; White Bone Band (11862, -0.40 DPS) [quest]; Thunderbrow Ring (13097, -0.43 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.8 attack_power points (2.29 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; White Bone Band (11862, -0.08 DPS) [quest]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (287.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (287.6 DPS) | yes | Molten Heart of the Mountain (249470, -1.63 DPS, sim-verified) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (287.6 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -3.71 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (287.6 DPS) | yes | Idol of the Moon (23197, -10.06 DPS, sim-verified) [world_drop] |

**New at 60:** head: Shadowcraft Cap; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0500200000000000-35201030021030212051-5053000000000000)

Set DPS (verified): 685.4. Weights run: 4.7s. Verify run: 9.5s. 1564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.434 ± 0.023, crit=1.731 ± 0.031 per rating point (14 rating = 1%, 24.234 per %), hit=2.094 ± 0.159 per rating point (10 rating = 1%, 20.942 per %), melee_haste=9.006 ± 2.016

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadowcraft Cap (16707) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (685.4 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -10.78 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 43.7 attack_power points (7.45 DPS) | yes | Woven Ivy Necklace (19159, -2.36 DPS, sim-verified) [quest]; Sunprism Pendant (20649, -2.47 DPS) [quest]; Ethereal Talisman (4430, -3.73 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (685.4 DPS) | yes | Warlord's Dragonhide Epaulets (16551, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -14.67 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 56.5 attack_power points (9.62 DPS) | yes | Dark Phantom Cape (13122, -3.64 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.65 DPS) [world_drop]; Pridelord Cape (14673, -4.66 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 147.9 attack_power points (25.18 DPS) | yes | Dawn Armor (252483, -2.16 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -3.45 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -4.95 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (685.4 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -13.17 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (685.4 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -11.27 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 84.9 attack_power points (14.45 DPS) | yes | Prowler's Leather Waistguard (252473, -3.71 DPS) [crafted]; Skulker's Leather Waistguard (252474, -4.20 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -4.94 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 111.6 attack_power points (19.00 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -1.39 DPS) [pvp]; Traveler's Leggings (8300, -5.66 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (685.4 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -16.56 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 34.1 attack_power points (5.80 DPS) | yes | Falcon's Hook (7552, -0.84 DPS) [dungeon]; Thunderbrow Ring (13097, -1.26 DPS) [world_drop]; Ring of Living Stone (18400, -1.68 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.2 attack_power points (4.97 DPS) | yes | Falcon's Hook (7552, -0.00 DPS) [dungeon]; Thunderbrow Ring (13097, -0.43 DPS) [world_drop]; Ring of Living Stone (18400, -0.85 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (685.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (685.4 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (685.4 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -6.46 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (685.4 DPS) | yes | Idol of the Moon (23197, -16.19 DPS, sim-verified) [world_drop] |

**New at 60:** head: Shadowcraft Cap; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

