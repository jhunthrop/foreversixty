# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 46.8. Weights run: 3.3s. Verify run: 5.4s. 230 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.088 ± 0.012, crit=0.674 ± 0.011 per rating point (14 rating = 1%, 9.443 per %), hit=0.708 ± 0.036 per rating point (10 rating = 1%, 7.075 per %), melee_haste=4.408 ± 0.424

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.7 attack_power points (0.97 DPS) | yes | Defender's Leather Hood (252447, -0.04 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.2 attack_power points (0.24 DPS) | yes | Scholarly Pendant (277203, -0.24 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.59 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.4 attack_power points (0.61 DPS) | yes | Slime-encrusted Pads (6461, -0.60 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.5 attack_power points (0.73 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.13 DPS) [crafted]; Bristlebark Cape (14571, -0.24 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 24.6 attack_power points (1.43 DPS) | yes | Tunic of Westfall (2041, -0.17 DPS, sim-verified) [quest]; Defender's Leather Armor (252434, -0.25 DPS) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 12.3 attack_power points (0.71 DPS) | yes | Forest Leather Bracers (3202, -0.11 DPS) [world_drop]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.23 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (46.8 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.71 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.04 DPS) | yes | Brawler's Leather Belt (252428, -0.10 DPS) [crafted]; Murloc Scale Belt (5780, -0.45 DPS) [crafted]; Deviate Scale Belt (6468, -4.32 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (46.8 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.80 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (46.8 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.55 DPS, sim-verified) [world_drop] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.46 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.35 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.2 attack_power points (0.24 DPS) | yes | The 1 Ring (8350, -0.01 DPS) [world]; Ring of the Moon (12052, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 306.4 attack_power points (17.78 DPS) | yes | Twisted Chanter's Staff (890, -1.09 DPS, sim-verified) [world_drop]; Rakzur Club (12983, -2.47 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -2.80 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Lavishly Jeweled Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde

### Band 30 (night-elf, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 98.5. Weights run: 3.4s. Verify run: 6.0s. 414 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.983 ± 0.013, crit=0.782 ± 0.013 per rating point (14 rating = 1%, 10.942 per %), hit=0.917 ± 0.051 per rating point (10 rating = 1%, 9.169 per %), melee_haste=5.337 ± 0.578

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.62 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.40 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.40 DPS) [crafted]; Brawler's Leather Helm (252512, -0.43 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.9 attack_power points (1.07 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.27 DPS) [world]; River Pride Choker (13087, -0.53 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.67 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 attack_power points (2.15 DPS) | yes | Mantle of Thieves (2264, -0.81 DPS) [dungeon]; Bristlebark Amice (14573, -0.94 DPS) [world_drop]; Barbaric Shoulders (5964, -1.07 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.9 attack_power points (1.34 DPS) | yes | Sergeant Major's Cape (16315, -0.27 DPS) [pvp]; Tigerstrike Mantle (13108, -0.27 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.54 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 27.9 attack_power points (1.88 DPS) | yes | Dusky Leather Armor (7374, -0.01 DPS) [crafted]; Defender's Leather Tunic (252450, -0.27 DPS) [crafted]; Brawler's Leather Armor (252490, -0.27 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.9 attack_power points (1.34 DPS) | yes | Barbaric Bracers (18948, -0.37 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.40 DPS) [quest]; Bands of Serra'kis (6902, -0.53 DPS) [dungeon] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (98.5 DPS) | yes | Wolfclaw Gloves (1978, +0.00 DPS) [dungeon]; Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, -4.04 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (98.5 DPS) | yes | Warden's Leather Belt (252460, +0.00 DPS) [crafted]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Prowler's Leather Belt (252459, -4.09 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (98.5 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Legguards (252516, -4.56 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (98.5 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -4.38 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.9 attack_power points (1.48 DPS) | yes | Ring of Precision (1491, -0.68 DPS) [dungeon]; Tiger Band (6749, -0.81 DPS) [quest]; Bonebinder's Signet (273032, -0.81 DPS) [world] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 11.9 attack_power points (0.81 DPS) | yes | Ring of Precision (1491, -0.00 DPS) [dungeon]; Tiger Band (6749, -0.13 DPS) [quest]; Bonebinder's Signet (273032, -0.13 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (98.5 DPS) | yes | Cobalt Crusher (7730, -2.91 DPS) [dungeon]; Gnarled Ash Staff (791, -6.01 DPS) [world_drop]; Viscous Hammer (13045, -16.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; finger1: Thunderbrow Ring; finger2: Seal of Wrynn; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (night-elf, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 134.2. Weights run: 4.0s. Verify run: 3.8s. 549 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.031 ± 0.026, crit=1.003 ± 0.030 per rating point (14 rating = 1%, 14.045 per %), hit=1.115 ± 0.068 per rating point (10 rating = 1%, 11.146 per %), melee_haste=4.770 ± 0.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.3 attack_power points (3.40 DPS) | yes | Hawkeye's Helm (14591, -1.07 DPS) [world_drop]; Warden's Wizard Hat (14604, -1.38 DPS) [world_drop]; Nightscape Headband (8176, -1.53 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.1 attack_power points (1.24 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.32 DPS) [world]; River Pride Choker (13087, -0.62 DPS) [world_drop]; Gazlowe's Charm (13088, -0.62 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.3 attack_power points (2.48 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.77 DPS) [crafted]; Fleshhide Shoulders (10774, -0.78 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 24.2 attack_power points (1.86 DPS) | yes | Hawkeye's Cloak (14593, -0.30 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.46 DPS) [quest]; Parachute Cloak (10518, -0.61 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 34.5 attack_power points (2.65 DPS) | yes | Nightscape Tunic (8175, -0.31 DPS) [crafted]; Barbaric Harness (5739, -0.33 DPS) [crafted]; Dusky Leather Armor (7374, -0.47 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (134.2 DPS) | yes | Hawkeye's Bracers (14590, +0.00 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.29 DPS) [dungeon]; Dusky Bracers (7378, -0.29 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | sim-verified (134.2 DPS) | yes | Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Imperial Leather Gloves (4063, -0.15 DPS) [dungeon]; Tiger Hunter Gloves (4107, -0.62 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.3 attack_power points (2.78 DPS) | yes | Skulker's Leather Belt (252520, -0.46 DPS) [crafted]; Prowler's Leather Belt (252459, -0.47 DPS) [crafted]; Scorpashi Sash (14652, -0.61 DPS) [world_drop] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | sim-verified (134.2 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.79 DPS) [crafted]; Brawler's Leather Pants (252500, -0.94 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | sim-verified (134.2 DPS) | yes | Skulker's Leather Shoes (252531, +0.00 DPS) [crafted]; Imperial Leather Boots (6431, -0.30 DPS) [dungeon]; Excelsior Boots (4109, -0.30 DPS) [quest] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | sim-verified (134.2 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Mark of Kern (2262, -0.33 DPS) [dungeon]; Assault Band (13095, -0.33 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | sim-verified (134.2 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Mark of Kern (2262, -0.33 DPS) [dungeon]; Assault Band (13095, -0.33 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Ironshod Bludgeon (9408, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Falcon's Hook; finger2: Ring of the Underwood

No-known-source sample (15 of 549, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 0000000000000000-50212032123012123001-5500000000000000)

Set DPS (verified): 176.5. Weights run: 3.9s. Verify run: 4.1s. 719 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.018 ± 0.029, crit=1.204 ± 0.037 per rating point (14 rating = 1%, 16.853 per %), hit=1.147 ± 0.085 per rating point (10 rating = 1%, 11.473 per %), melee_haste=not significant (3.897 ± 1.139)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.4 attack_power points (3.55 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.16 DPS) [vendor]; Scorpashi Skullcap (14658, -0.88 DPS) [world_drop]; Helm of Fire (8348, -0.93 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (2.85 DPS) | yes | Kaleidoscope Chain (13084, -1.56 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.93 DPS) [world]; River Pride Choker (13087, -2.18 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 42.4 attack_power points (3.24 DPS) | yes | Failed Flying Experiment (9647, -0.01 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.03 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.46 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 30.3 attack_power points (2.32 DPS) | yes | Sergeant Major's Cape (16336, -0.38 DPS) [pvp]; Pridelord Cape (14673, -0.42 DPS) [world_drop]; Blackflame Cape (13109, -0.46 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 60.5 attack_power points (4.64 DPS) | yes | Blazewind Breastplate (11193, -0.58 DPS) [quest]; Knight's Leather Armor (220854, -1.09 DPS) [vendor]; Charred Leather Tunic (19127, -1.85 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 39.1 attack_power points (2.99 DPS) | yes | Prowler's Leather Bracers (252539, -0.39 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.42 DPS) [crafted]; Pridelord Bands (14672, -0.60 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 49.0 attack_power points (3.75 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.07 DPS) [crafted]; Serpentskin Gloves (8260, -0.20 DPS) [world_drop]; Heraldic Gloves (8121, -0.52 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 55.0 attack_power points (4.21 DPS) | yes | Skulker's Leather Waistguard (252474, -0.03 DPS) [crafted]; Ogron's Sash (13117, -1.31 DPS) [world_drop]; Pridelord Girdle (14674, -1.68 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 63.3 attack_power points (4.85 DPS) | yes | Serpentskin Leggings (8262, -0.18 DPS) [world_drop]; Knight's Leather Pants (220858, -1.30 DPS) [vendor]; Triprunner Dungarees (9624, -1.56 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 50.8 attack_power points (3.89 DPS) | yes | Skulker's Leather Boots (252469, -0.03 DPS) [crafted]; Albino Crocscale Boots (17728, -0.80 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.96 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (176.5 DPS) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Ring of the Underwood (2951, -0.28 DPS) [world_drop]; Thunderbrow Ring (13097, -0.35 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | sim-verified (176.5 DPS) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Ring of the Underwood (2951, -0.01 DPS) [world_drop]; Thunderbrow Ring (13097, -0.08 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (176.5 DPS) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 719, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 0000000000000000-50212032123012123001-5550230000000000)

Set DPS (verified): 269.0. Weights run: 4.0s. Verify run: 9.8s. 1597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.457 ± 0.050, crit=2.081 ± 0.071 per rating point (14 rating = 1%, 29.137 per %), hit=2.289 ± 0.166 per rating point (10 rating = 1%, 22.895 per %), melee_haste=not significant (6.260 ± 2.136)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 95.3 attack_power points (7.94 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, +0.00 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.9 attack_power points (3.58 DPS) | yes | Sunprism Pendant (20649, -1.09 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -2.02 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.58 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (269.0 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -6.20 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 56.9 attack_power points (4.73 DPS) | yes | Sergeant Major's Cape (16337, -0.85 DPS, sim-verified) [pvp]; Dark Phantom Cape (13122, -1.67 DPS) [world_drop]; Blackflame Cape (13109, -2.28 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 146.0 attack_power points (12.15 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, -1.29 DPS) [vendor]; Dawn Armor (252483, -1.60 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Breastplate (16452, -2.28 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (269.0 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -6.61 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (269.0 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -5.41 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (269.0 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -5.41 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 105.8 attack_power points (8.81 DPS) | yes | Knight-Captain's Dragonhide Leggings (227178, +0.00 DPS) [vendor]; Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Traveler's Leggings (8300, -2.54 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 58.5 attack_power points (4.87 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -7.17 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.9 attack_power points (3.57 DPS) | yes | Band of the Penitent (13217, -1.15 DPS) [quest]; Ring of the Underwood (2951, -1.16 DPS) [world_drop]; Falcon's Hook (7552, -1.18 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 34.4 attack_power points (2.86 DPS) | yes | Band of the Penitent (13217, -0.44 DPS) [quest]; Ring of the Underwood (2951, -0.45 DPS) [world_drop]; Falcon's Hook (7552, -0.47 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (269.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (269.0 DPS) | yes | Mark of the Chosen (17774, -1.62 DPS, sim-verified) [quest] |
| main_hand | Hammer of the Titans (12796) | Blacksmithing [crafted] | sim-verified (269.0 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Staff of Hale Magefire (13000, -2.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (269.0 DPS) | yes | Idol of the Moon (23197, -9.20 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; main_hand: Hammer of the Titans; ranged: Idol of the Dream

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (night-elf, 0500200000000000-50230032023011125001-5050030000000000)

Set DPS (verified): 616.0. Weights run: 3.9s. Verify run: 9.5s. 1597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.849 ± 0.058, crit=2.291 ± 0.078 per rating point (14 rating = 1%, 32.080 per %), hit=2.500 ± 0.191 per rating point (10 rating = 1%, 25.003 per %), melee_haste=14.640 ± 2.348

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 107.9 attack_power points (16.40 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Headguard (227173, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 48.0 attack_power points (7.30 DPS) | yes | Sunprism Pendant (20649, -2.10 DPS) [quest]; Kaleidoscope Chain (13084, -4.23 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -5.47 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.7 attack_power points (9.53 DPS) | yes | Dark Phantom Cape (13122, -3.04 DPS) [world_drop]; Sergeant Major's Cape (16337, -3.39 DPS) [pvp]; Blackflame Cape (13109, -4.34 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 162.5 attack_power points (24.69 DPS) | yes | Timbermaw Tunic (252484, -0.99 DPS) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -3.27 DPS) [vendor]; Knight-Captain's Dragonhide Chestpiece (227176, -5.28 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Skulker's Leather Bracers (252540, -0.26 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.45 DPS) [crafted] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Might of the Timbermaw (19044, +0.00 DPS) [crafted]; Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 112.9 attack_power points (17.15 DPS) | yes | Knight-Captain's Dragonhide Leggings (227178, +0.00 DPS) [vendor]; Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 67.3 attack_power points (10.23 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -0.93 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (616.0 DPS) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Falcon's Hook (7552, -1.16 DPS) [dungeon]; Band of the Penitent (13217, -1.19 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | sim-verified (616.0 DPS) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Falcon's Hook (7552, -0.10 DPS) [dungeon]; Band of the Penitent (13217, -0.12 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (616.0 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (616.0 DPS) | yes | - |
| main_hand | Hammer of the Titans (12796) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Titans; ranged: Idol of the Dream

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (tauren, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 45.8. Weights run: 3.3s. Verify run: 5.4s. 226 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.088 ± 0.012, crit=0.674 ± 0.011 per rating point (14 rating = 1%, 9.443 per %), hit=0.708 ± 0.036 per rating point (10 rating = 1%, 7.075 per %), melee_haste=4.408 ± 0.424

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.7 attack_power points (0.97 DPS) | yes | Defender's Leather Hood (252447, -0.04 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.2 attack_power points (0.24 DPS) | yes | Roadwatcher's Confidence (281265, -0.23 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.4 attack_power points (0.61 DPS) | yes | Slime-encrusted Pads (6461, -0.57 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.5 attack_power points (0.73 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.13 DPS) [crafted]; Bristlebark Cape (14571, -0.24 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 24.6 attack_power points (1.43 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.48 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.4 attack_power points (0.61 DPS) | yes | Bristlebark Bindings (14569, -0.01 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Spare Part Bindings (279875, -0.24 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (45.8 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.60 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.04 DPS) | yes | Brawler's Leather Belt (252428, -0.10 DPS) [crafted]; Murloc Scale Belt (5780, -0.45 DPS) [crafted]; Deviate Scale Belt (6468, -4.21 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (45.8 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.70 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (45.8 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.45 DPS, sim-verified) [world_drop] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.46 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; The 1 Ring (8350, -0.23 DPS) [world] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 6.3 attack_power points (0.36 DPS) | yes | Loop of Sacrifice (281673, -0.02 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon]; The 1 Ring (8350, -0.13 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 306.4 attack_power points (17.78 DPS) | yes | Hammerbone (270018, -0.46 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -1.17 DPS) [world_drop]; Rakzur Club (12983, -2.47 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Bounty Hunter's Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 226, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (tauren, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 97.1. Weights run: 3.4s. Verify run: 6.0s. 414 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.983 ± 0.013, crit=0.782 ± 0.013 per rating point (14 rating = 1%, 10.942 per %), hit=0.917 ± 0.051 per rating point (10 rating = 1%, 9.169 per %), melee_haste=5.337 ± 0.578

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.62 DPS) | yes | Brawler's Leather Helm (252512, -0.33 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.40 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.40 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.9 attack_power points (1.07 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.30 DPS, sim-verified) [world]; River Pride Choker (13087, -0.53 DPS) [world_drop]; Erudite's Amulet (277204, -0.81 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 attack_power points (2.15 DPS) | yes | Mantle of Thieves (2264, -0.81 DPS) [dungeon]; Bristlebark Amice (14573, -0.94 DPS) [world_drop]; Barbaric Shoulders (5964, -1.06 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.9 attack_power points (1.34 DPS) | yes | Tigerstrike Mantle (13108, -0.35 DPS, sim-verified) [world_drop]; Glowing Lizardscale Cloak (6449, -0.54 DPS) [dungeon]; Swiftrunner Cape (6745, -0.54 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 27.9 attack_power points (1.88 DPS) | yes | Dusky Leather Armor (7374, -0.01 DPS) [crafted]; Defender's Leather Tunic (252450, -0.27 DPS) [crafted]; Brawler's Leather Armor (252490, -0.27 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.9 attack_power points (1.34 DPS) | yes | Barbaric Bracers (18948, -0.35 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.53 DPS) [dungeon]; Insignia Bracers (6410, -0.54 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (97.1 DPS) | yes | Wolfclaw Gloves (1978, +0.00 DPS) [dungeon]; Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, -3.95 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (97.1 DPS) | yes | Warden's Leather Belt (252460, +0.00 DPS) [crafted]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Prowler's Leather Belt (252459, -3.99 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (97.1 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Legguards (252516, -4.47 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (97.1 DPS) | yes | Stomping Boots (3741, +0.00 DPS) [quest]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -4.29 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.9 attack_power points (1.48 DPS) | yes | Ring of Precision (1491, -0.68 DPS) [dungeon]; Tiger Band (6749, -0.81 DPS) [quest]; Bonebinder's Signet (273032, -0.81 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 15.9 attack_power points (1.07 DPS) | yes | Ring of Precision (1491, -0.35 DPS, sim-verified) [dungeon]; Tiger Band (6749, -0.40 DPS) [quest]; Bonebinder's Signet (273032, -0.40 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (97.1 DPS) | yes | Cobalt Crusher (7730, -2.91 DPS) [dungeon]; Gnarled Ash Staff (791, -6.01 DPS) [world_drop]; Viscous Hammer (13045, -16.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 130.3. Weights run: 4.0s. Verify run: 3.8s. 547 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.031 ± 0.026, crit=1.003 ± 0.030 per rating point (14 rating = 1%, 14.045 per %), hit=1.115 ± 0.068 per rating point (10 rating = 1%, 11.146 per %), melee_haste=4.770 ± 0.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.3 attack_power points (3.40 DPS) | yes | Warden's Wizard Hat (14604, -1.38 DPS) [world_drop]; Hawkeye's Helm (14591, -1.45 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -1.53 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 18.1 attack_power points (1.39 DPS) | yes | Kaleidoscope Chain (13084, -0.15 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.47 DPS) [world]; River Pride Choker (13087, -0.78 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.3 attack_power points (2.48 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.77 DPS) [crafted]; Fleshhide Shoulders (10774, -0.78 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.2 attack_power points (1.55 DPS) | yes | Imperial Cloak (6432, -0.30 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.30 DPS) [world_drop]; Parachute Cloak (10518, -0.50 DPS, sim-verified) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 34.5 attack_power points (2.65 DPS) | yes | Barbaric Harness (5739, -0.33 DPS) [crafted]; Dusky Leather Armor (7374, -0.47 DPS) [crafted]; Nightscape Tunic (8175, -0.75 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.2 attack_power points (1.55 DPS) | yes | Branded Leather Bracers (19508, -0.01 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.30 DPS) [dungeon]; Dusky Bracers (7378, -0.30 DPS) [crafted] |
| hands | Skulker's Leather Gloves (252525) | Leatherworking [crafted] | 38.3 attack_power points (2.94 DPS) | yes | Prowler's Leather Gloves (252524, -0.00 DPS) [crafted]; Imperial Leather Gloves (4063, -0.16 DPS) [dungeon]; Tiger Hunter Gloves (4107, -0.62 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.3 attack_power points (2.78 DPS) | yes | Prowler's Leather Belt (252459, -0.47 DPS) [crafted]; Skulker's Leather Belt (252520, -0.60 DPS, sim-verified) [crafted]; Scorpashi Sash (14652, -0.61 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 42.7 attack_power points (3.27 DPS) | yes | Triprunner Dungarees (9624, -0.39 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -0.80 DPS) [crafted]; Brawler's Leather Pants (252500, -0.95 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 36.3 attack_power points (2.79 DPS) | yes | Prowler's Leather Shoes (252465, -0.01 DPS) [crafted]; Imperial Leather Boots (6431, -0.31 DPS) [dungeon]; Excelsior Boots (4109, -0.31 DPS) [quest] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 24.3 attack_power points (1.86 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Mark of Kern (2262, -0.33 DPS) [dungeon]; Assault Band (13095, -0.33 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.3 attack_power points (1.86 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Mark of Kern (2262, -0.33 DPS) [dungeon]; Assault Band (13095, -0.33 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (130.3 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Ironshod Bludgeon (9408, +0.00 DPS) [dungeon]; Mograine's Might (7723, -11.77 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Wolffear Harness; hands: Skulker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Ring of the Underwood; finger2: Falcon's Hook

No-known-source sample (15 of 547, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 50 (tauren, 0000000000000000-50212032123012123001-5500000000000000)

Set DPS (verified): 175.1. Weights run: 3.9s. Verify run: 4.0s. 717 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.018 ± 0.029, crit=1.204 ± 0.037 per rating point (14 rating = 1%, 16.853 per %), hit=1.147 ± 0.085 per rating point (10 rating = 1%, 11.473 per %), melee_haste=not significant (3.897 ± 1.139)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.4 attack_power points (3.55 DPS) | yes | Blood Guard's Leather Headband (220851, -0.16 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.17 DPS) [quest]; Scorpashi Skullcap (14658, -0.88 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (2.85 DPS) | yes | Woven Ivy Necklace (19159, -0.45 DPS) [quest]; Ethereal Talisman (4430, -1.39 DPS) [quest]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 42.4 attack_power points (3.24 DPS) | yes | Failed Flying Experiment (9647, -0.01 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.03 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.46 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 30.3 attack_power points (2.32 DPS) | yes | Pridelord Cape (14673, -0.42 DPS) [world_drop]; Blackflame Cape (13109, -0.46 DPS) [world_drop]; Serpentskin Cloak (8259, -0.62 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 60.5 attack_power points (4.64 DPS) | yes | Blazewind Breastplate (11193, -0.58 DPS) [quest]; Stone Guard's Leather Armor (220855, -1.09 DPS) [vendor]; Charred Leather Tunic (19127, -1.85 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 39.1 attack_power points (2.99 DPS) | yes | Prowler's Leather Bracers (252539, -0.39 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.42 DPS) [crafted]; Pridelord Bands (14672, -0.60 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 49.0 attack_power points (3.75 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.07 DPS) [crafted]; Serpentskin Gloves (8260, -0.20 DPS) [world_drop]; Heraldic Gloves (8121, -0.52 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 55.0 attack_power points (4.21 DPS) | yes | Skulker's Leather Waistguard (252474, -0.03 DPS) [crafted]; Ogron's Sash (13117, -1.31 DPS) [world_drop]; Pridelord Girdle (14674, -1.68 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 60.9 attack_power points (4.66 DPS) | yes | Stone Guard's Leather Pants (220859, -1.12 DPS) [vendor]; Triprunner Dungarees (9624, -1.38 DPS) [quest]; Basilisk Hide Pants (1718, -1.42 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 50.8 attack_power points (3.89 DPS) | yes | Skulker's Leather Boots (252469, -0.03 DPS) [crafted]; Albino Crocscale Boots (17728, -0.80 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.96 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (175.1 DPS) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Ring of the Underwood (2951, -0.28 DPS) [world_drop]; White Bone Band (11862, -0.33 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | sim-verified (175.1 DPS) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Ring of the Underwood (2951, -0.01 DPS) [world_drop]; White Bone Band (11862, -0.06 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (175.1 DPS) | yes | - |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (175.1 DPS) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 717, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60 (tauren, 0000000000000000-50212032123012123001-5550230000000000)

Set DPS (verified): 269.0. Weights run: 4.0s. Verify run: 9.7s. 1597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.457 ± 0.050, crit=2.081 ± 0.071 per rating point (14 rating = 1%, 29.137 per %), hit=2.289 ± 0.166 per rating point (10 rating = 1%, 22.895 per %), melee_haste=not significant (6.260 ± 2.136)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 95.3 attack_power points (7.94 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -1.76 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.9 attack_power points (3.58 DPS) | yes | Sunprism Pendant (20649, -1.12 DPS) [quest]; Woven Ivy Necklace (19159, -1.59 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -1.84 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (269.0 DPS) | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS) [pvp]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -7.19 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 56.9 attack_power points (4.73 DPS) | yes | Dark Phantom Cape (13122, -1.36 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.28 DPS) [world_drop]; Fel Cape (279269, -2.31 DPS) [crafted] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 146.0 attack_power points (12.15 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, -1.29 DPS) [vendor]; Dawn Armor (252483, -1.54 DPS, sim-verified) [crafted]; Warlord's Dragonhide Hauberk (16549, -2.28 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (269.0 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -7.65 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (269.0 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -5.47 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (269.0 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -5.47 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 105.8 attack_power points (8.81 DPS) | yes | Legionnaire's Dragonhide Leggings (227177, +0.00 DPS) [pvp]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Traveler's Leggings (8300, -2.65 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 58.5 attack_power points (4.87 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -7.92 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.9 attack_power points (3.57 DPS) | yes | Band of the Penitent (13217, -1.15 DPS) [quest]; Ring of the Underwood (2951, -1.16 DPS) [world_drop]; Falcon's Hook (7552, -1.18 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 34.4 attack_power points (2.86 DPS) | yes | Band of the Penitent (13217, -0.44 DPS) [quest]; Ring of the Underwood (2951, -0.45 DPS) [world_drop]; Falcon's Hook (7552, -0.47 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (269.0 DPS) | yes | Frozen Heart of the Mountain (249469, -3.12 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (269.0 DPS) | yes | Frozen Heart of the Mountain (249469, -1.29 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Titans (12796) | Blacksmithing [crafted] | sim-verified (269.0 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Staff of Hale Magefire (13000, -2.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (269.0 DPS) | yes | Idol of the Moon (23197, -9.58 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Titans; ranged: Idol of the Dream

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60, raid preset (tauren, 0500200000000000-50230032023011125001-5050030000000000)

Set DPS (verified): 618.6. Weights run: 3.9s. Verify run: 9.3s. 1597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.849 ± 0.058, crit=2.291 ± 0.078 per rating point (14 rating = 1%, 32.080 per %), hit=2.500 ± 0.191 per rating point (10 rating = 1%, 25.003 per %), melee_haste=14.640 ± 2.348

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 107.9 attack_power points (16.40 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -4.19 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 48.0 attack_power points (7.30 DPS) | yes | Woven Ivy Necklace (19159, -1.40 DPS) [quest]; Sunprism Pendant (20649, -2.10 DPS) [quest]; Ethereal Talisman (4430, -3.90 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (618.6 DPS) | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS) [pvp]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -16.91 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.7 attack_power points (9.53 DPS) | yes | Dark Phantom Cape (13122, -3.07 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.34 DPS) [world_drop]; Pridelord Cape (14673, -4.63 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 162.5 attack_power points (24.69 DPS) | yes | Timbermaw Tunic (252484, -0.99 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -3.27 DPS) [vendor]; Legionnaire's Dragonhide Chestpiece (227179, -5.28 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (618.6 DPS) | yes | Skulker's Leather Bracers (252540, -0.26 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.45 DPS) [crafted]; Deepfury Bracers (13120, -16.69 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (618.6 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -14.67 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (618.6 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -14.67 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 112.9 attack_power points (17.15 DPS) | yes | Legionnaire's Dragonhide Leggings (227177, +0.00 DPS) [pvp]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Traveler's Leggings (8300, -4.82 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 67.3 attack_power points (10.23 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -19.42 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.0 attack_power points (6.84 DPS) | yes | Ring of the Underwood (2951, -1.84 DPS) [world_drop]; Falcon's Hook (7552, -1.94 DPS) [dungeon]; Band of the Penitent (13217, -1.96 DPS) [quest] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.9 attack_power points (6.06 DPS) | yes | Ring of the Underwood (2951, -1.06 DPS) [world_drop]; Falcon's Hook (7552, -1.16 DPS) [dungeon]; Band of the Penitent (13217, -1.19 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (618.6 DPS) | yes | Frozen Heart of the Mountain (249469, -5.62 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (618.6 DPS) | yes | Frozen Heart of the Mountain (249469, -3.45 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Titans (12796) | Blacksmithing [crafted] | sim-verified (618.6 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Staff of Hale Magefire (13000, -7.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (618.6 DPS) | yes | Idol of the Moon (23197, -15.77 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Titans; ranged: Idol of the Dream

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

