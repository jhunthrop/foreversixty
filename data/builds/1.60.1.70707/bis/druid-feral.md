# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 69.7. Weights run: 2.8s. Verify run: 3.3s. 203 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.234 ± 0.013, crit=0.762 ± 0.012 per rating point (14 rating = 1%, 10.672 per %), hit=0.896 ± 0.037 per rating point (10 rating = 1%, 8.962 per %), melee_haste=5.312 ± 0.428

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.9 attack_power points (1.02 DPS) | yes | Defender's Leather Hood (252447, -0.21 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.5 attack_power points (0.25 DPS) | yes | Tarnished Locket (279870, -0.29 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 11.2 attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.73 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.4 attack_power points (0.76 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.15 DPS) [crafted]; Grave Shroud (279865, -0.17 DPS) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 25.6 attack_power points (1.46 DPS) | yes | Tunic of Westfall (2041, -0.06 DPS) [quest]; Defender's Leather Armor (252434, -0.28 DPS) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 12.9 attack_power points (0.74 DPS) | yes | Forest Leather Bracers (3202, -0.10 DPS) [world_drop]; Bristlebark Bindings (14569, -0.13 DPS) [world_drop]; Wolf Bracers (4794, -0.23 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (69.7 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.46 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.02 DPS) | yes | Brawler's Leather Belt (252428, -0.06 DPS) [crafted]; Dusty Belt (279897, -0.39 DPS) [quest]; Deviate Scale Belt (6468, -4.11 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (69.7 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.66 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (69.7 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.43 DPS, sim-verified) [world_drop] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.45 DPS) | yes | The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.34 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.5 attack_power points (0.25 DPS) | yes | Ring of the Moon (12052, -0.14 DPS) [world_drop]; The 1 Ring (8350, -0.22 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 306.9 attack_power points (17.45 DPS) | yes | Staff of Westfall (2042, -1.13 DPS) [quest]; Twisted Chanter's Staff (890, -1.18 DPS) [world_drop]; Living Root (6631, -2.03 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Lavishly Jeweled Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 134.3. Weights run: 2.7s. Verify run: 2.1s. 387 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.066 ± 0.014, crit=0.851 ± 0.015 per rating point (14 rating = 1%, 11.911 per %), hit=0.896 ± 0.051 per rating point (10 rating = 1%, 8.959 per %), melee_haste=6.448 ± 0.610

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.76 DPS) | yes | Brawler's Leather Helm (252512, -0.25 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.44 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.44 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.3 attack_power points (1.20 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.31 DPS) [world]; River Pride Choker (13087, -0.61 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.75 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.7 attack_power points (2.40 DPS) | yes | Barbaric Shoulders (5964, -0.91 DPS) [crafted]; Bristlebark Amice (14573, -1.05 DPS) [world_drop]; Mantle of Thieves (2264, -1.21 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.5 attack_power points (1.50 DPS) | yes | Tigerstrike Mantle (13108, -0.29 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.31 DPS) [pvp]; Fenrus' Hide (6340, -0.59 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 28.9 attack_power points (2.13 DPS) | yes | Brawler's Leather Tunic (252508, -0.03 DPS) [crafted]; Brawler's Leather Armor (252490, -0.33 DPS) [crafted]; Defender's Leather Tunic (252450, -0.33 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (1.50 DPS) | yes | Barbaric Bracers (18948, -0.43 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.45 DPS) [quest]; Insignia Bracers (6410, -0.59 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 26.4 attack_power points (1.94 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.44 DPS) [dungeon] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 30.6 attack_power points (2.25 DPS) | yes | Prowler's Leather Belt (252459, -0.01 DPS) [crafted]; Warden's Leather Belt (252460, -0.46 DPS) [crafted]; Stalker's Leather Belt (252521, -0.88 DPS) [crafted] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 32.6 attack_power points (2.40 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Petrolspill Leggings (9509, -0.27 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.5 attack_power points (1.66 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.44 DPS) [world_drop]; Footpads of the Fang (10411, -0.74 DPS) [dungeon] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.2 attack_power points (1.63 DPS) | yes | Seal of Wrynn (2933, -0.73 DPS) [quest]; Defias Renegade Ring (1076, -0.87 DPS) [dungeon]; Monkey Ring (6748, -0.87 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 12.4 attack_power points (0.91 DPS) | yes | Seal of Wrynn (2933, -0.01 DPS) [quest]; Defias Renegade Ring (1076, -0.15 DPS) [dungeon]; Monkey Ring (6748, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (134.3 DPS) | yes | Cobalt Crusher (7730, -3.20 DPS) [dungeon]; Gnarled Ash Staff (791, -6.59 DPS) [world_drop]; Viscous Hammer (13045, -25.14 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Skulker's Leather Belt; legs: Brawler's Leather Legguards; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Ring of Precision; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 181.3. Weights run: 3.2s. Verify run: 2.0s. 523 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.074 ± 0.030, crit=1.064 ± 0.037 per rating point (14 rating = 1%, 14.902 per %), hit=0.946 ± 0.082 per rating point (10 rating = 1%, 9.465 per %), melee_haste=5.387 ± 1.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.8 attack_power points (3.78 DPS) | yes | Warden's Wizard Hat (14604, -1.51 DPS) [world_drop]; Hawkeye's Helm (14591, -1.60 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -1.68 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.3 attack_power points (1.38 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.36 DPS) [world]; River Pride Choker (13087, -0.70 DPS) [world_drop]; Gazlowe's Charm (13088, -0.70 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.8 attack_power points (2.77 DPS) | yes | Flintrock Shoulders (7755, -0.18 DPS) [dungeon]; Nightscape Shoulders (8192, -0.84 DPS) [crafted]; Fleshhide Shoulders (10774, -0.88 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 24.4 attack_power points (2.06 DPS) | yes | Yeti Fur Cloak (2805, -0.51 DPS) [quest]; Parachute Cloak (10518, -0.66 DPS) [crafted]; Hawkeye's Cloak (14593, -1.89 DPS, sim-verified) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 35.3 attack_power points (2.98 DPS) | yes | Barbaric Harness (5739, -0.39 DPS) [crafted]; Dusky Leather Armor (7374, -0.53 DPS) [crafted]; Nightscape Tunic (8175, -1.00 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (1.73 DPS) | yes | Branded Leather Bracers (19508, -0.04 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.33 DPS) [dungeon]; Dusky Bracers (7378, -0.33 DPS) [crafted] |
| hands | Skulker's Leather Gloves (252525) | Leatherworking [crafted] | 38.7 attack_power points (3.27 DPS) | yes | Prowler's Leather Gloves (252524, -0.01 DPS) [crafted]; Imperial Leather Gloves (4063, -0.18 DPS) [dungeon]; Gloves of Holy Might (867, -0.32 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.7 attack_power points (3.10 DPS) | yes | Prowler's Leather Belt (252459, -0.53 DPS) [crafted]; Scorpashi Sash (14652, -0.67 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.67 DPS, sim-verified) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 43.6 attack_power points (3.68 DPS) | yes | Triprunner Dungarees (9624, -0.56 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -0.92 DPS) [crafted]; Brawler's Leather Pants (252500, -1.09 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 36.8 attack_power points (3.11 DPS) | yes | Prowler's Leather Shoes (252465, -0.03 DPS) [crafted]; Imperial Leather Boots (6431, -0.34 DPS) [dungeon]; Excelsior Boots (4109, -0.35 DPS) [quest] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 24.7 attack_power points (2.09 DPS) | yes | Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Mark of Kern (2262, -0.40 DPS) [dungeon]; Assault Band (13095, -0.40 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.7 attack_power points (2.08 DPS) | yes | Mark of Kern (2262, -0.39 DPS) [dungeon]; Assault Band (13095, -0.39 DPS) [world_drop]; Thunderbrow Ring (13097, -0.53 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (181.3 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Staff of Jordan (873, -30.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; hands: Skulker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Ring of the Underwood; finger2: Falcon's Hook

No-known-source sample (15 of 523, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0000000000000000-50212032123012123001-5500000000000000)

Set DPS (verified): 223.7. Weights run: 3.1s. Verify run: 2.5s. 693 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.053 ± 0.035, crit=1.261 ± 0.044 per rating point (14 rating = 1%, 17.651 per %), hit=1.143 ± 0.104 per rating point (10 rating = 1%, 11.432 per %), melee_haste=5.545 ± 1.283

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.8 attack_power points (3.95 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.14 DPS) [vendor]; Helm of Fire (8348, -1.00 DPS) [crafted]; Scorpashi Skullcap (14658, -1.10 DPS, sim-verified) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.7 attack_power points (3.18 DPS) | yes | Kaleidoscope Chain (13084, -1.18 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -2.17 DPS) [world]; River Pride Choker (13087, -2.44 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 42.7 attack_power points (3.60 DPS) | yes | Failed Flying Experiment (9647, -0.01 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.02 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.47 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 30.8 attack_power points (2.60 DPS) | yes | Sergeant Major's Cape (16336, -0.45 DPS) [pvp]; Pridelord Cape (14673, -0.48 DPS) [world_drop]; Blackflame Cape (13109, -0.52 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 61.2 attack_power points (5.16 DPS) | yes | Blazewind Breastplate (11193, -0.62 DPS) [quest]; Knight's Leather Armor (220854, -1.19 DPS) [vendor]; Charred Leather Tunic (19127, -2.04 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 39.6 attack_power points (3.34 DPS) | yes | Prowler's Leather Bracers (252539, -0.46 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.48 DPS) [crafted]; Pridelord Bands (14672, -0.68 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 49.3 attack_power points (4.16 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.06 DPS) [crafted]; Serpentskin Gloves (8260, -0.21 DPS) [world_drop]; Heraldic Gloves (8121, -0.57 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 55.4 attack_power points (4.68 DPS) | yes | Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Ogron's Sash (13117, -1.45 DPS) [world_drop]; Pridelord Girdle (14674, -1.86 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 63.8 attack_power points (5.39 DPS) | yes | Serpentskin Leggings (8262, -0.20 DPS) [world_drop]; Knight's Leather Pants (220858, -1.41 DPS) [vendor]; Triprunner Dungarees (9624, -1.71 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 51.2 attack_power points (4.32 DPS) | yes | Skulker's Leather Boots (252469, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.85 DPS) [dungeon]; Prowler's Leather Shoes (252465, -1.06 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.7 attack_power points (2.43 DPS) | yes | Ring of the Underwood (2951, -0.32 DPS) [world_drop]; Thunderbrow Ring (13097, -0.42 DPS) [world_drop]; Blackstone Ring (17713, -0.74 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.1 attack_power points (2.12 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Blackstone Ring (17713, -0.43 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (223.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (223.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (223.7 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -32.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 693, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0000000000000000-50212032123012123001-5550230000000000)

Set DPS (verified): 290.6. Weights run: 3.1s. Verify run: 6.4s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.175 ± 0.041, crit=1.678 ± 0.059 per rating point (14 rating = 1%, 23.487 per %), hit=1.452 ± 0.126 per rating point (10 rating = 1%, 14.521 per %), melee_haste=6.453 ± 1.611

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadowcraft Cap (16707) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (290.6 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -5.75 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.3 attack_power points (3.70 DPS) | yes | Sunprism Pendant (20649, -1.85 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -2.05 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.57 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (290.6 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -6.26 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 52.6 attack_power points (4.96 DPS) | yes | Sergeant Major's Cape (16337, -1.50 DPS, sim-verified) [pvp]; Dark Phantom Cape (13122, -1.88 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 135.2 attack_power points (12.74 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, -2.15 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -2.48 DPS) [vendor]; Dawn Armor (252483, -2.86 DPS, sim-verified) [crafted] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (290.6 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -6.49 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (290.6 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -5.21 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 76.6 attack_power points (7.22 DPS) | yes | Skulker's Leather Waistguard (252474, -1.86 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.02 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -2.49 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 100.7 attack_power points (9.49 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -0.96 DPS) [vendor]; Traveler's Leggings (8300, -2.78 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (290.6 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Prowler's Leather Boots (252468, -7.71 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.4 attack_power points (2.87 DPS) | yes | Ring of the Underwood (2951, -0.40 DPS) [world_drop]; Thunderbrow Ring (13097, -0.60 DPS) [world_drop]; Ring of Living Stone (18400, -0.80 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 26.2 attack_power points (2.46 DPS) | yes | Ring of the Underwood (2951, -0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Ring of Living Stone (18400, -0.39 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (290.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (290.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (290.6 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -4.16 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (290.6 DPS) | yes | Idol of the Moon (23197, -9.80 DPS, sim-verified) [world_drop] |

**New at 60:** head: Shadowcraft Cap; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 0500200000000000-50230032023011125001-5050030000000000)

Set DPS (verified): 684.3. Weights run: 2.9s. Verify run: 6.2s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.548 ± 0.053, crit=1.911 ± 0.074 per rating point (14 rating = 1%, 26.748 per %), hit=1.573 ± 0.150 per rating point (10 rating = 1%, 15.732 per %), melee_haste=10.520 ± 2.043

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 93.6 attack_power points (16.26 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -3.63 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 44.1 attack_power points (7.66 DPS) | yes | Sunprism Pendant (20649, -2.52 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -4.36 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -5.58 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (684.3 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -15.63 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 58.2 attack_power points (10.11 DPS) | yes | Dark Phantom Cape (13122, -2.66 DPS, sim-verified) [world_drop]; Sergeant Major's Cape (16337, -3.51 DPS) [pvp]; Blackflame Cape (13109, -4.80 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 145.2 attack_power points (25.22 DPS) | yes | Dawn Armor (252483, -2.46 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -4.00 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -4.83 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (684.3 DPS) | yes | Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -14.76 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (684.3 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -13.06 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (684.3 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -13.06 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 107.5 attack_power points (18.66 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.39 DPS) [vendor]; Traveler's Leggings (8300, -5.42 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 58.2 attack_power points (10.11 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -18.29 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 35.7 attack_power points (6.19 DPS) | yes | Falcon's Hook (7552, -1.07 DPS) [dungeon]; Thunderbrow Ring (13097, -1.81 DPS) [world_drop]; Ring of Living Stone (18400, -2.37 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.9 attack_power points (5.19 DPS) | yes | Falcon's Hook (7552, -0.06 DPS) [dungeon]; Thunderbrow Ring (13097, -0.80 DPS) [world_drop]; Ring of Living Stone (18400, -1.37 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (684.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (684.3 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (684.3 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -15.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (684.3 DPS) | yes | Idol of the Moon (23197, -15.42 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 68.0. Weights run: 2.8s. Verify run: 3.3s. 199 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.234 ± 0.013, crit=0.762 ± 0.012 per rating point (14 rating = 1%, 10.672 per %), hit=0.896 ± 0.037 per rating point (10 rating = 1%, 8.962 per %), melee_haste=5.312 ± 0.428

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.9 attack_power points (1.02 DPS) | yes | Defender's Leather Hood (252447, -0.11 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.5 attack_power points (0.25 DPS) | yes | Roadwatcher's Confidence (281265, -0.28 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 11.2 attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.71 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.4 attack_power points (0.76 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.15 DPS) [crafted]; Grave Shroud (279865, -0.17 DPS) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 25.6 attack_power points (1.46 DPS) | yes | Defender's Leather Armor (252434, -0.33 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.51 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 11.2 attack_power points (0.64 DPS) | yes | Bristlebark Bindings (14569, -0.03 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Spare Part Bindings (279875, -0.25 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (68.0 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.37 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.02 DPS) | yes | Brawler's Leather Belt (252428, -0.06 DPS) [crafted]; Murloc Scale Belt (5780, -0.42 DPS) [crafted]; Deviate Scale Belt (6468, -4.00 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (68.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.57 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (68.0 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.34 DPS, sim-verified) [world_drop] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.45 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; The 1 Ring (8350, -0.21 DPS) [world] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 6.7 attack_power points (0.38 DPS) | yes | Loop of Sacrifice (281673, -0.04 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; The 1 Ring (8350, -0.14 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 306.9 attack_power points (17.45 DPS) | yes | Crescent Staff (6505, -0.37 DPS) [quest]; Living Root (6631, -0.56 DPS) [dungeon]; Hammerbone (270018, -0.72 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Bounty Hunter's Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 132.3. Weights run: 2.7s. Verify run: 2.2s. 387 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.066 ± 0.014, crit=0.851 ± 0.015 per rating point (14 rating = 1%, 11.911 per %), hit=0.896 ± 0.051 per rating point (10 rating = 1%, 8.959 per %), melee_haste=6.448 ± 0.610

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.76 DPS) | yes | Brawler's Leather Helm (252512, -0.25 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.44 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.44 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.3 attack_power points (1.20 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.38 DPS, sim-verified) [world]; River Pride Choker (13087, -0.61 DPS) [world_drop]; Erudite's Amulet (277204, -0.89 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.7 attack_power points (2.40 DPS) | yes | Barbaric Shoulders (5964, -0.91 DPS) [crafted]; Bristlebark Amice (14573, -1.05 DPS) [world_drop]; Mantle of Thieves (2264, -1.18 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.5 attack_power points (1.50 DPS) | yes | Tigerstrike Mantle (13108, -0.29 DPS) [world_drop]; Fenrus' Hide (6340, -0.59 DPS) [dungeon]; Swiftrunner Cape (6745, -0.59 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 28.9 attack_power points (2.13 DPS) | yes | Brawler's Leather Tunic (252508, -0.03 DPS) [crafted]; Brawler's Leather Armor (252490, -0.33 DPS) [crafted]; Defender's Leather Tunic (252450, -0.33 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (1.50 DPS) | yes | Barbaric Bracers (18948, -0.44 DPS, sim-verified) [crafted]; Insignia Bracers (6410, -0.59 DPS) [world_drop]; Bands of Serra'kis (6902, -0.62 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 26.4 attack_power points (1.94 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.44 DPS) [dungeon] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 30.6 attack_power points (2.25 DPS) | yes | Prowler's Leather Belt (252459, -0.01 DPS) [crafted]; Warden's Leather Belt (252460, -0.46 DPS) [crafted]; Deftkin Belt (16659, -0.76 DPS) [quest] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 32.6 attack_power points (2.40 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Petrolspill Leggings (9509, -0.27 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.5 attack_power points (1.66 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.44 DPS) [world_drop]; Warsong Boots (16977, -0.44 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.2 attack_power points (1.63 DPS) | yes | Ring of Precision (1491, -0.72 DPS) [dungeon]; Defias Renegade Ring (1076, -0.87 DPS) [dungeon]; Monkey Ring (6748, -0.87 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 16.3 attack_power points (1.20 DPS) | yes | Ring of Precision (1491, -0.28 DPS) [dungeon]; Defias Renegade Ring (1076, -0.44 DPS) [dungeon]; Monkey Ring (6748, -0.44 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (132.3 DPS) | yes | Cobalt Crusher (7730, -3.20 DPS) [dungeon]; Gnarled Ash Staff (791, -6.59 DPS) [world_drop]; Viscous Hammer (13045, -24.65 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Skulker's Leather Belt; legs: Brawler's Leather Legguards; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 176.3. Weights run: 3.2s. Verify run: 2.0s. 521 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.074 ± 0.030, crit=1.064 ± 0.037 per rating point (14 rating = 1%, 14.902 per %), hit=0.946 ± 0.082 per rating point (10 rating = 1%, 9.465 per %), melee_haste=5.387 ± 1.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.8 attack_power points (3.78 DPS) | yes | Warden's Wizard Hat (14604, -1.51 DPS) [world_drop]; Hawkeye's Helm (14591, -1.55 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -1.68 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 18.3 attack_power points (1.54 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.53 DPS) [world]; River Pride Choker (13087, -0.87 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.8 attack_power points (2.77 DPS) | yes | Flintrock Shoulders (7755, -0.18 DPS) [dungeon]; Nightscape Shoulders (8192, -0.84 DPS) [crafted]; Fleshhide Shoulders (10774, -0.88 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.5 attack_power points (1.73 DPS) | yes | Imperial Cloak (6432, -0.33 DPS) [dungeon]; Parachute Cloak (10518, -0.33 DPS) [crafted]; Tigerstrike Mantle (13108, -0.33 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 35.3 attack_power points (2.98 DPS) | yes | Barbaric Harness (5739, -0.39 DPS) [crafted]; Dusky Leather Armor (7374, -0.53 DPS) [crafted]; Nightscape Tunic (8175, -1.07 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (1.73 DPS) | yes | Branded Leather Bracers (19508, -0.04 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.33 DPS) [dungeon]; Dusky Bracers (7378, -0.33 DPS) [crafted] |
| hands | Skulker's Leather Gloves (252525) | Leatherworking [crafted] | 38.7 attack_power points (3.27 DPS) | yes | Prowler's Leather Gloves (252524, -0.01 DPS) [crafted]; Imperial Leather Gloves (4063, -0.18 DPS) [dungeon]; Gloves of Holy Might (867, -0.32 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.7 attack_power points (3.10 DPS) | yes | Prowler's Leather Belt (252459, -0.53 DPS) [crafted]; Scorpashi Sash (14652, -0.67 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.75 DPS, sim-verified) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 43.6 attack_power points (3.68 DPS) | yes | Triprunner Dungarees (9624, -0.71 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -0.92 DPS) [crafted]; Brawler's Leather Pants (252500, -1.09 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 36.8 attack_power points (3.11 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.34 DPS) [dungeon]; Excelsior Boots (4109, -0.35 DPS) [quest] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 24.7 attack_power points (2.09 DPS) | yes | Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Mark of Kern (2262, -0.40 DPS) [dungeon]; Assault Band (13095, -0.40 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.7 attack_power points (2.08 DPS) | yes | Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Mark of Kern (2262, -0.39 DPS) [dungeon]; Assault Band (13095, -0.39 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (176.3 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Staff of Jordan (873, -29.53 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Wolffear Harness; hands: Skulker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Ring of the Underwood; finger2: Falcon's Hook

No-known-source sample (15 of 521, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0000000000000000-50212032123012123001-5500000000000000)

Set DPS (verified): 226.6. Weights run: 3.1s. Verify run: 2.5s. 691 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.053 ± 0.035, crit=1.261 ± 0.044 per rating point (14 rating = 1%, 17.651 per %), hit=1.143 ± 0.104 per rating point (10 rating = 1%, 11.432 per %), melee_haste=5.545 ± 1.283

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.8 attack_power points (3.95 DPS) | yes | Blood Guard's Leather Headband (220851, -0.14 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.19 DPS) [quest]; Scorpashi Skullcap (14658, -0.97 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.7 attack_power points (3.18 DPS) | yes | Woven Ivy Necklace (19159, -0.82 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -1.56 DPS) [quest]; Kaleidoscope Chain (13084, -1.75 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 42.7 attack_power points (3.60 DPS) | yes | Skulker's Leather Shoulder (252535, -0.02 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.47 DPS) [vendor]; Failed Flying Experiment (9647, -0.91 DPS, sim-verified) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 30.8 attack_power points (2.60 DPS) | yes | Blackflame Cape (13109, -0.52 DPS) [world_drop]; Serpentskin Cloak (8259, -0.69 DPS) [world_drop]; Pridelord Cape (14673, -0.97 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 61.2 attack_power points (5.16 DPS) | yes | Blazewind Breastplate (11193, -0.86 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -1.19 DPS) [vendor]; Charred Leather Tunic (19127, -2.04 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 39.6 attack_power points (3.34 DPS) | yes | Prowler's Leather Bracers (252539, -0.46 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.48 DPS) [crafted]; Pridelord Bands (14672, -0.68 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 49.3 attack_power points (4.16 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.06 DPS) [crafted]; Serpentskin Gloves (8260, -0.21 DPS) [world_drop]; Heraldic Gloves (8121, -0.57 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 55.4 attack_power points (4.68 DPS) | yes | Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Ogron's Sash (13117, -1.45 DPS) [world_drop]; Pridelord Girdle (14674, -1.86 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 61.4 attack_power points (5.19 DPS) | yes | Stone Guard's Leather Pants (220859, -0.86 DPS, sim-verified) [vendor]; Triprunner Dungarees (9624, -1.51 DPS) [quest]; Basilisk Hide Pants (1718, -1.55 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 51.2 attack_power points (4.32 DPS) | yes | Skulker's Leather Boots (252469, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.85 DPS) [dungeon]; Prowler's Leather Shoes (252465, -1.06 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.7 attack_power points (2.43 DPS) | yes | Ring of the Underwood (2951, -0.32 DPS) [world_drop]; White Bone Band (11862, -0.40 DPS) [quest]; Thunderbrow Ring (13097, -0.42 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.1 attack_power points (2.12 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; White Bone Band (11862, -0.09 DPS) [quest]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (226.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (226.6 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (226.6 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -34.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0000000000000000-50212032123012123001-5550230000000000)

Set DPS (verified): 293.0. Weights run: 3.1s. Verify run: 6.4s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.175 ± 0.041, crit=1.678 ± 0.059 per rating point (14 rating = 1%, 23.487 per %), hit=1.452 ± 0.126 per rating point (10 rating = 1%, 14.521 per %), melee_haste=6.453 ± 1.611

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadowcraft Cap (16707) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (293.0 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -5.22 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.3 attack_power points (3.70 DPS) | yes | Woven Ivy Necklace (19159, -1.20 DPS, sim-verified) [quest]; Sunprism Pendant (20649, -1.24 DPS) [quest]; Ethereal Talisman (4430, -1.84 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (293.0 DPS) | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS) [pvp]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -6.89 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 52.6 attack_power points (4.96 DPS) | yes | Pridelord Cape (14673, -2.49 DPS) [world_drop]; Blackflame Cape (13109, -2.50 DPS) [world_drop]; Dark Phantom Cape (13122, -2.53 DPS, sim-verified) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 135.2 attack_power points (12.74 DPS) | yes | Dawn Armor (252483, -1.94 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -2.15 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -2.48 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (293.0 DPS) | yes | Prowler's Leather Bracers (252539, +0.00 DPS) [crafted]; Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Deepfury Bracers (13120, -7.27 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (293.0 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -5.75 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 76.6 attack_power points (7.22 DPS) | yes | Skulker's Leather Waistguard (252474, -1.86 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.42 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -2.49 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 100.7 attack_power points (9.49 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -0.96 DPS) [pvp]; Traveler's Leggings (8300, -2.98 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (293.0 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Prowler's Leather Boots (252468, -8.16 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.4 attack_power points (2.87 DPS) | yes | Ring of the Underwood (2951, -0.40 DPS) [world_drop]; Thunderbrow Ring (13097, -0.60 DPS) [world_drop]; White Bone Band (11862, -0.61 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 26.2 attack_power points (2.46 DPS) | yes | Ring of the Underwood (2951, -0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; White Bone Band (11862, -0.20 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (293.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (293.0 DPS) | yes | Molten Heart of the Mountain (249470, -1.66 DPS, sim-verified) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (293.0 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Hammer of the Northern Wind (810, -4.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (293.0 DPS) | yes | Idol of the Moon (23197, -9.76 DPS, sim-verified) [world_drop] |

**New at 60:** head: Shadowcraft Cap; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0500200000000000-50230032023011125001-5050030000000000)

Set DPS (verified): 690.3. Weights run: 2.9s. Verify run: 6.1s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.548 ± 0.053, crit=1.911 ± 0.074 per rating point (14 rating = 1%, 26.748 per %), hit=1.573 ± 0.150 per rating point (10 rating = 1%, 15.732 per %), melee_haste=10.520 ± 2.043

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 93.6 attack_power points (16.26 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -5.93 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 44.1 attack_power points (7.66 DPS) | yes | Woven Ivy Necklace (19159, -1.39 DPS) [quest]; Sunprism Pendant (20649, -2.35 DPS) [quest]; Ethereal Talisman (4430, -3.98 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (690.3 DPS) | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS) [pvp]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -19.38 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 58.2 attack_power points (10.11 DPS) | yes | Dark Phantom Cape (13122, -4.55 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.80 DPS) [world_drop]; Pridelord Cape (14673, -4.98 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 145.2 attack_power points (25.22 DPS) | yes | Dawn Armor (252483, -1.47 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -4.00 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -4.83 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (690.3 DPS) | yes | Skulker's Leather Bracers (252540, +0.00 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -18.59 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (690.3 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -16.92 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (690.3 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -16.92 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 107.5 attack_power points (18.66 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -1.39 DPS) [pvp]; Traveler's Leggings (8300, -5.05 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 58.2 attack_power points (10.11 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -20.83 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 35.7 attack_power points (6.19 DPS) | yes | Falcon's Hook (7552, -1.07 DPS) [dungeon]; Thunderbrow Ring (13097, -1.81 DPS) [world_drop]; White Bone Band (11862, -2.03 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.9 attack_power points (5.19 DPS) | yes | Falcon's Hook (7552, -0.06 DPS) [dungeon]; Thunderbrow Ring (13097, -0.80 DPS) [world_drop]; White Bone Band (11862, -1.02 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (690.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (690.3 DPS) | yes | Molten Heart of the Mountain (249470, -3.15 DPS, sim-verified) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (690.3 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -8.23 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (690.3 DPS) | yes | Idol of the Moon (23197, -15.54 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; main_hand: Elemental Mage Staff; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

