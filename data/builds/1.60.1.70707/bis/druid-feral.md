# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 46.7. Weights run: 4.2s. Verify run: 4.8s. 203 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.998 ± 0.011, crit=0.617 ± 0.010 per rating point (14 rating = 1%, 8.645 per %), hit=0.704 ± 0.030 per rating point (10 rating = 1%, 7.035 per %), melee_haste=3.941 ± 0.365

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.93 DPS) | yes | Brawler's Leather Hood (252504, -0.00 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.0 attack_power points (0.23 DPS) | yes | Tarnished Locket (279870, -0.24 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 attack_power points (0.58 DPS) | yes | Slime-encrusted Pads (6461, -0.57 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 attack_power points (0.70 DPS) | yes | Grave Shroud (279865, -0.12 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted]; Cape of the Brotherhood (5193, -0.12 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (46.7 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.51 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 12.0 attack_power points (0.70 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Forest Leather Bracers (3202, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.23 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (46.7 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.64 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.04 DPS) | yes | Deviate Scale Belt (6468, -0.12 DPS) [crafted]; Murloc Scale Belt (5780, -0.46 DPS) [crafted]; Brawler's Leather Belt (252428, -4.24 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (46.7 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.74 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.0 attack_power points (1.28 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Blackened Defias Boots (10402, -0.58 DPS) [dungeon]; Footpads of the Fang (10411, -0.58 DPS) [dungeon] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.46 DPS) | yes | Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Ring of the Moon (12052, -0.35 DPS) [world_drop] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 4.0 attack_power points (0.23 DPS) | yes | Lavishly Jeweled Ring (1156, -0.00 DPS) [dungeon]; Ring of the Moon (12052, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 306.0 attack_power points (17.75 DPS) | yes | Staff of the Blessed Seer (2271, -0.98 DPS, sim-verified) [dungeon]; Staff of Westfall (2042, -1.10 DPS) [quest]; Twisted Chanter's Staff (890, -1.15 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: The 1 Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 98.5. Weights run: 4.2s. Verify run: 5.3s. 387 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.911 ± 0.012, crit=0.728 ± 0.012 per rating point (14 rating = 1%, 10.185 per %), hit=0.760 ± 0.043 per rating point (10 rating = 1%, 7.604 per %), melee_haste=5.466 ± 0.514

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.61 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.40 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.40 DPS) [crafted]; Brawler's Leather Helm (252512, -0.43 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.6 attack_power points (1.05 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.25 DPS) [world]; River Pride Choker (13087, -0.51 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.65 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.0 attack_power points (2.09 DPS) | yes | Mantle of Thieves (2264, -0.80 DPS) [dungeon]; Bristlebark Amice (14573, -0.91 DPS) [world_drop]; Barbaric Shoulders (5964, -1.07 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.4 attack_power points (1.30 DPS) | yes | Sergeant Major's Cape (16315, -0.25 DPS) [pvp]; Tigerstrike Mantle (13108, -0.27 DPS) [world_drop]; Fenrus' Hide (6340, -0.53 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 27.3 attack_power points (1.83 DPS) | yes | Dusky Leather Armor (7374, -0.04 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS) [crafted]; Brawler's Leather Armor (252490, -0.26 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.5 attack_power points (1.31 DPS) | yes | Barbaric Bracers (18948, -0.37 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.39 DPS) [quest]; Bands of Serra'kis (6902, -0.50 DPS) [dungeon] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (98.5 DPS) | yes | Wolfclaw Gloves (1978, +0.00 DPS) [dungeon]; Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, -4.04 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (98.5 DPS) | yes | Warden's Leather Belt (252460, +0.00 DPS) [crafted]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Prowler's Leather Belt (252459, -4.09 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (98.5 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Legguards (252516, -4.56 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (98.5 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -4.38 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.7 attack_power points (1.46 DPS) | yes | Ring of Precision (1491, -0.69 DPS) [dungeon]; Tiger Band (6749, -0.79 DPS) [quest]; Bonebinder's Signet (273032, -0.79 DPS) [world] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 11.7 attack_power points (0.79 DPS) | yes | Ring of Precision (1491, -0.02 DPS) [dungeon]; Tiger Band (6749, -0.12 DPS) [quest]; Bonebinder's Signet (273032, -0.12 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (98.5 DPS) | yes | Cobalt Crusher (7730, -2.88 DPS) [dungeon]; Hydrocane (9452, -5.05 DPS) [dungeon]; Viscous Hammer (13045, -16.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; feet: Blackened Defias Boots; finger1: Thunderbrow Ring; finger2: Seal of Wrynn; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 134.3. Weights run: 4.7s. Verify run: 3.0s. 523 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.946 ± 0.025, crit=0.943 ± 0.031 per rating point (14 rating = 1%, 13.207 per %), hit=0.800 ± 0.065 per rating point (10 rating = 1%, 8.005 per %), melee_haste=4.599 ± 0.815

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 43.4 attack_power points (3.36 DPS) | yes | Warden's Wizard Hat (14604, -1.40 DPS) [world_drop]; Hawkeye's Helm (14591, -1.48 DPS, sim-verified) [world_drop]; Defender's Leather Helm (252455, -1.50 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.8 attack_power points (1.22 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.29 DPS) [world]; River Pride Choker (13087, -0.60 DPS) [world_drop]; Gazlowe's Charm (13088, -0.60 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.4 attack_power points (2.43 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.75 DPS) [dungeon]; Nightscape Shoulders (8192, -0.77 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 23.7 attack_power points (1.83 DPS) | yes | Yeti Fur Cloak (2805, -0.46 DPS) [quest]; Parachute Cloak (10518, -0.63 DPS) [crafted]; Hawkeye's Cloak (14593, -1.06 DPS, sim-verified) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 33.1 attack_power points (2.56 DPS) | yes | Nightscape Tunic (8175, -0.30 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.43 DPS) [crafted]; Barbaric Harness (5739, -0.48 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.55 DPS) | yes | Hawkeye's Bracers (14590, -0.03 DPS) [world_drop]; Barbaric Bracers (18948, -0.33 DPS) [crafted]; Dusky Bracers (7378, -0.34 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 37.5 attack_power points (2.90 DPS) | yes | Skulker's Leather Gloves (252525, -0.00 DPS) [crafted]; Imperial Leather Gloves (4063, -0.15 DPS) [dungeon]; Gloves of Holy Might (867, -0.33 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 35.5 attack_power points (2.75 DPS) | yes | Skulker's Leather Belt (252520, -0.46 DPS) [crafted]; Prowler's Leather Belt (252459, -0.51 DPS, sim-verified) [crafted]; Scorpashi Sash (14652, -0.62 DPS) [world_drop] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 41.0 attack_power points (3.18 DPS) | yes | Basilisk Hide Pants (1718, -0.01 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.74 DPS) [crafted]; Brawler's Leather Pants (252500, -0.89 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 35.6 attack_power points (2.76 DPS) | yes | Skulker's Leather Shoes (252531, -0.02 DPS) [crafted]; Excelsior Boots (4109, -0.32 DPS) [quest]; Imperial Leather Boots (6431, -0.33 DPS) [dungeon] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.5 attack_power points (1.82 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Mark of Kern (2262, -0.27 DPS) [dungeon]; Assault Band (13095, -0.27 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.5 attack_power points (1.82 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Mark of Kern (2262, -0.27 DPS) [dungeon]; Assault Band (13095, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (134.3 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Staff of Jordan (873, -12.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Falcon's Hook; finger2: Ring of the Underwood

No-known-source sample (15 of 523, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0000000000000000-50212032123012123001-5500000000000000)

Set DPS (verified): 173.9. Weights run: 4.8s. Verify run: 3.7s. 693 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.018 ± 0.029, crit=1.204 ± 0.037 per rating point (14 rating = 1%, 16.853 per %), hit=1.147 ± 0.085 per rating point (10 rating = 1%, 11.473 per %), melee_haste=not significant (3.897 ± 1.139)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.4 attack_power points (3.55 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.16 DPS) [vendor]; Helm of Fire (8348, -0.93 DPS) [crafted]; Scorpashi Skullcap (14658, -1.35 DPS, sim-verified) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (2.85 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.93 DPS) [world]; River Pride Choker (13087, -2.18 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.44 DPS, sim-verified) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 42.4 attack_power points (3.24 DPS) | yes | Failed Flying Experiment (9647, -0.01 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.03 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.46 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 30.3 attack_power points (2.32 DPS) | yes | Sergeant Major's Cape (16336, -0.38 DPS) [pvp]; Pridelord Cape (14673, -0.42 DPS) [world_drop]; Blackflame Cape (13109, -0.46 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 60.5 attack_power points (4.64 DPS) | yes | Blazewind Breastplate (11193, -0.85 DPS, sim-verified) [quest]; Knight's Leather Armor (220854, -1.09 DPS) [vendor]; Charred Leather Tunic (19127, -1.85 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 39.1 attack_power points (2.99 DPS) | yes | Prowler's Leather Bracers (252539, -0.39 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.42 DPS) [crafted]; Pridelord Bands (14672, -0.60 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 49.0 attack_power points (3.75 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.07 DPS) [crafted]; Serpentskin Gloves (8260, -0.20 DPS) [world_drop]; Heraldic Gloves (8121, -0.52 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 55.0 attack_power points (4.21 DPS) | yes | Skulker's Leather Waistguard (252474, -0.03 DPS) [crafted]; Ogron's Sash (13117, -1.31 DPS) [world_drop]; Pridelord Girdle (14674, -1.68 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 63.3 attack_power points (4.85 DPS) | yes | Serpentskin Leggings (8262, -0.18 DPS) [world_drop]; Knight's Leather Pants (220858, -1.30 DPS) [vendor]; Triprunner Dungarees (9624, -1.56 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 50.8 attack_power points (3.89 DPS) | yes | Skulker's Leather Boots (252469, -0.03 DPS) [crafted]; Albino Crocscale Boots (17728, -0.80 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.96 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.3 attack_power points (2.16 DPS) | yes | Ring of the Underwood (2951, -0.28 DPS) [world_drop]; Thunderbrow Ring (13097, -0.35 DPS) [world_drop]; Blackstone Ring (17713, -0.63 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.8 attack_power points (1.90 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (173.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (173.9 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (173.9 DPS) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, -6.67 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 693, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0000000000000000-50212032123012123001-5550230000000000)

Set DPS (verified): 264.4. Weights run: 4.7s. Verify run: 9.7s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.669 ± 0.049, crit=2.384 ± 0.071 per rating point (14 rating = 1%, 33.377 per %), hit=2.323 ± 0.148 per rating point (10 rating = 1%, 23.225 per %), melee_haste=10.121 ± 1.848

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 106.1 attack_power points (9.03 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -1.19 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 45.7 attack_power points (3.89 DPS) | yes | Sunprism Pendant (20649, -0.99 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -2.23 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.87 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (264.4 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -6.10 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 60.0 attack_power points (5.11 DPS) | yes | Dark Phantom Cape (13122, -1.70 DPS) [world_drop]; Sergeant Major's Cape (16337, -1.79 DPS) [pvp]; Fel Cape (279269, -2.27 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.3 attack_power points (13.22 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -1.48 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -2.52 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (264.4 DPS) | yes | Skulker's Leather Bracers (252540, -0.05 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.13 DPS) [crafted]; Deepfury Bracers (13120, -6.18 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (264.4 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -5.86 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (264.4 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -5.86 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 109.6 attack_power points (9.33 DPS) | yes | Knight-Captain's Dragonhide Leggings (227178, +0.00 DPS) [vendor]; Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Traveler's Leggings (8300, -2.93 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 63.0 attack_power points (5.36 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -6.95 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 37.4 attack_power points (3.18 DPS) | yes | Falcon's Hook (7552, -0.57 DPS) [dungeon]; Thunderbrow Ring (13097, -1.00 DPS) [world_drop]; Blush Ember Ring (13093, -1.29 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 31.1 attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.04 DPS) [dungeon]; Thunderbrow Ring (13097, -0.47 DPS) [world_drop]; Blush Ember Ring (13093, -0.76 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (264.4 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (264.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Titans (12796) | Blacksmithing [crafted] | sim-verified (264.4 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -20.32 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (264.4 DPS) | yes | Idol of the Moon (23197, -9.04 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Titans; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 0500200000000000-50230032023011125001-5050030000000000)

Set DPS (verified): 610.3. Weights run: 4.5s. Verify run: 9.6s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=3.135 ± 0.062, crit=2.625 ± 0.085 per rating point (14 rating = 1%, 36.751 per %), hit=2.636 ± 0.193 per rating point (10 rating = 1%, 26.361 per %), melee_haste=16.255 ± 2.324

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 121.0 attack_power points (18.62 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -2.15 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 51.8 attack_power points (7.97 DPS) | yes | Sunprism Pendant (20649, -3.00 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -4.68 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -6.12 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (610.3 DPS) | yes | Lieutenant Commander's Dragonhide Shoulders (227172, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -12.00 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 67.0 attack_power points (10.32 DPS) | yes | Sergeant Major's Cape (16337, -3.75 DPS) [pvp]; Dark Phantom Cape (13122, -4.01 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.53 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 179.1 attack_power points (27.57 DPS) | yes | Timbermaw Tunic (252484, -2.05 DPS) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -4.33 DPS) [vendor]; Knight-Captain's Dragonhide Chestpiece (227176, -6.36 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (610.3 DPS) | yes | Skulker's Leather Bracers (252540, -0.52 DPS) [crafted]; Pridelord Bands (14672, -0.72 DPS) [world_drop]; Deepfury Bracers (13120, -12.03 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (610.3 DPS) | yes | Knight-Lieutenant's Dragonhide Grips (227183, +0.00 DPS) [pvp]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -11.19 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (610.3 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -11.19 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 118.0 attack_power points (18.17 DPS) | yes | Knight-Captain's Dragonhide Leggings (227178, +0.00 DPS) [vendor]; Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Traveler's Leggings (8300, -4.79 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 73.7 attack_power points (11.35 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -14.26 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.9 attack_power points (6.76 DPS) | yes | Falcon's Hook (7552, -1.40 DPS) [dungeon]; Thunderbrow Ring (13097, -2.60 DPS) [world_drop]; Disengagement Ring (276202, -2.90 DPS) [vendor] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 35.7 attack_power points (5.50 DPS) | yes | Falcon's Hook (7552, -0.14 DPS) [dungeon]; Thunderbrow Ring (13097, -1.35 DPS) [world_drop]; Disengagement Ring (276202, -1.64 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (610.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (610.3 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Titans (12796) | Blacksmithing [crafted] | sim-verified (610.3 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -42.16 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (610.3 DPS) | yes | Idol of the Moon (23197, -11.02 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Titans; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 45.9. Weights run: 4.2s. Verify run: 4.8s. 199 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.998 ± 0.011, crit=0.617 ± 0.010 per rating point (14 rating = 1%, 8.645 per %), hit=0.704 ± 0.030 per rating point (10 rating = 1%, 7.035 per %), melee_haste=3.941 ± 0.365

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.93 DPS) | yes | Brawler's Leather Hood (252504, -0.00 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.0 attack_power points (0.23 DPS) | yes | Roadwatcher's Confidence (281265, -0.23 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 attack_power points (0.58 DPS) | yes | Slime-encrusted Pads (6461, -0.59 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 attack_power points (0.70 DPS) | yes | Grave Shroud (279865, -0.12 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted]; Cape of the Brotherhood (5193, -0.12 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (45.9 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.37 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 10.0 attack_power points (0.58 DPS) | yes | Forest Leather Bracers (3202, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Spare Part Bindings (279875, -0.23 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (45.9 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.50 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.04 DPS) | yes | Deviate Scale Belt (6468, -0.12 DPS) [crafted]; Murloc Scale Belt (5780, -0.46 DPS) [crafted]; Brawler's Leather Belt (252428, -4.08 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (45.9 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.60 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.0 attack_power points (1.28 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Blackened Defias Boots (10402, -0.58 DPS) [dungeon]; Footpads of the Fang (10411, -0.58 DPS) [dungeon] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.46 DPS) | yes | Bounty Hunter's Ring (5351, -0.12 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.35 DPS) | yes | Bounty Hunter's Ring (5351, -0.00 DPS) [quest]; The 1 Ring (8350, -0.12 DPS) [world]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (17.76 DPS) | yes | Smite's Mighty Hammer (7230, -0.01 DPS) [dungeon]; Crescent Staff (6505, -0.43 DPS) [quest]; Staff of the Blessed Seer (2271, -0.71 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Hammerbone

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 97.1. Weights run: 4.2s. Verify run: 5.3s. 387 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.911 ± 0.012, crit=0.728 ± 0.012 per rating point (14 rating = 1%, 10.185 per %), hit=0.760 ± 0.043 per rating point (10 rating = 1%, 7.604 per %), melee_haste=5.466 ± 0.514

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.61 DPS) | yes | Brawler's Leather Helm (252512, -0.33 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.40 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.40 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.6 attack_power points (1.05 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.30 DPS, sim-verified) [world]; River Pride Choker (13087, -0.51 DPS) [world_drop]; Erudite's Amulet (277204, -0.79 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.0 attack_power points (2.09 DPS) | yes | Mantle of Thieves (2264, -0.80 DPS) [dungeon]; Bristlebark Amice (14573, -0.91 DPS) [world_drop]; Barbaric Shoulders (5964, -1.06 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.4 attack_power points (1.30 DPS) | yes | Tigerstrike Mantle (13108, -0.35 DPS, sim-verified) [world_drop]; Fenrus' Hide (6340, -0.53 DPS) [dungeon]; Swiftrunner Cape (6745, -0.53 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 27.3 attack_power points (1.83 DPS) | yes | Dusky Leather Armor (7374, -0.04 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS) [crafted]; Brawler's Leather Armor (252490, -0.26 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.5 attack_power points (1.31 DPS) | yes | Barbaric Bracers (18948, -0.35 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.50 DPS) [dungeon]; Insignia Bracers (6410, -0.54 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (97.1 DPS) | yes | Wolfclaw Gloves (1978, +0.00 DPS) [dungeon]; Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, -3.95 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (97.1 DPS) | yes | Warden's Leather Belt (252460, +0.00 DPS) [crafted]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Prowler's Leather Belt (252459, -3.99 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (97.1 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Legguards (252516, -4.47 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (97.1 DPS) | yes | Stomping Boots (3741, +0.00 DPS) [quest]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -4.29 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.7 attack_power points (1.46 DPS) | yes | Ring of Precision (1491, -0.69 DPS) [dungeon]; Tiger Band (6749, -0.79 DPS) [quest]; Bonebinder's Signet (273032, -0.79 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 15.6 attack_power points (1.05 DPS) | yes | Ring of Precision (1491, -0.35 DPS, sim-verified) [dungeon]; Tiger Band (6749, -0.38 DPS) [quest]; Bonebinder's Signet (273032, -0.38 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (97.1 DPS) | yes | Cobalt Crusher (7730, -2.88 DPS) [dungeon]; Hydrocane (9452, -5.05 DPS) [dungeon]; Viscous Hammer (13045, -16.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; feet: Blackened Defias Boots; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 130.4. Weights run: 4.7s. Verify run: 3.0s. 521 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.946 ± 0.025, crit=0.943 ± 0.031 per rating point (14 rating = 1%, 13.207 per %), hit=0.800 ± 0.065 per rating point (10 rating = 1%, 8.005 per %), melee_haste=4.599 ± 0.815

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 43.4 attack_power points (3.36 DPS) | yes | Warden's Wizard Hat (14604, -1.40 DPS) [world_drop]; Hawkeye's Helm (14591, -1.43 DPS, sim-verified) [world_drop]; Defender's Leather Helm (252455, -1.50 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 17.8 attack_power points (1.38 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.45 DPS) [world]; River Pride Choker (13087, -0.76 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.4 attack_power points (2.43 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.75 DPS) [dungeon]; Nightscape Shoulders (8192, -0.77 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.6 attack_power points (1.52 DPS) | yes | Imperial Cloak (6432, -0.31 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.31 DPS) [world_drop]; Parachute Cloak (10518, -0.40 DPS, sim-verified) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 33.1 attack_power points (2.56 DPS) | yes | Barbaric Harness (5739, -0.27 DPS) [crafted]; Nightscape Tunic (8175, -0.30 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.43 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.55 DPS) | yes | Hawkeye's Bracers (14590, -0.03 DPS) [world_drop]; Barbaric Bracers (18948, -0.33 DPS) [crafted]; Dusky Bracers (7378, -0.34 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 37.5 attack_power points (2.90 DPS) | yes | Skulker's Leather Gloves (252525, -0.00 DPS) [crafted]; Imperial Leather Gloves (4063, -0.15 DPS) [dungeon]; Gloves of Holy Might (867, -0.33 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 35.5 attack_power points (2.75 DPS) | yes | Prowler's Leather Belt (252459, -0.45 DPS) [crafted]; Skulker's Leather Belt (252520, -0.46 DPS) [crafted]; Scorpashi Sash (14652, -0.62 DPS) [world_drop] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 41.0 attack_power points (3.18 DPS) | yes | Basilisk Hide Pants (1718, -0.01 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.74 DPS) [crafted]; Brawler's Leather Pants (252500, -0.89 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 35.6 attack_power points (2.76 DPS) | yes | Excelsior Boots (4109, -0.32 DPS) [quest]; Imperial Leather Boots (6431, -0.33 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.53 DPS, sim-verified) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.5 attack_power points (1.82 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Mark of Kern (2262, -0.27 DPS) [dungeon]; Assault Band (13095, -0.27 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.5 attack_power points (1.82 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Mark of Kern (2262, -0.27 DPS) [dungeon]; Assault Band (13095, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (130.4 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Staff of Jordan (873, -11.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Falcon's Hook; finger2: Ring of the Underwood

No-known-source sample (15 of 521, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0000000000000000-50212032123012123001-5500000000000000)

Set DPS (verified): 176.0. Weights run: 4.8s. Verify run: 3.6s. 691 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.018 ± 0.029, crit=1.204 ± 0.037 per rating point (14 rating = 1%, 16.853 per %), hit=1.147 ± 0.085 per rating point (10 rating = 1%, 11.473 per %), melee_haste=not significant (3.897 ± 1.139)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.4 attack_power points (3.55 DPS) | yes | Blood Guard's Leather Headband (220851, -0.16 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.17 DPS) [quest]; Scorpashi Skullcap (14658, -0.88 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (2.85 DPS) | yes | Woven Ivy Necklace (19159, -0.45 DPS) [quest]; Ethereal Talisman (4430, -1.39 DPS) [quest]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 42.4 attack_power points (3.24 DPS) | yes | Failed Flying Experiment (9647, -0.01 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.03 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.46 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 30.3 attack_power points (2.32 DPS) | yes | Blackflame Cape (13109, -0.46 DPS) [world_drop]; Serpentskin Cloak (8259, -0.62 DPS) [world_drop]; Pridelord Cape (14673, -0.66 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 60.5 attack_power points (4.64 DPS) | yes | Blazewind Breastplate (11193, -0.67 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -1.09 DPS) [vendor]; Charred Leather Tunic (19127, -1.85 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 39.1 attack_power points (2.99 DPS) | yes | Prowler's Leather Bracers (252539, -0.39 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.42 DPS) [crafted]; Pridelord Bands (14672, -0.60 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 49.0 attack_power points (3.75 DPS) | yes | Skulker's Leather Gauntlets (252548, +0.00 DPS, sim-verified) [crafted]; Serpentskin Gloves (8260, -0.20 DPS) [world_drop]; Heraldic Gloves (8121, -0.52 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 55.0 attack_power points (4.21 DPS) | yes | Skulker's Leather Waistguard (252474, -0.03 DPS) [crafted]; Ogron's Sash (13117, -1.31 DPS) [world_drop]; Pridelord Girdle (14674, -1.68 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 60.9 attack_power points (4.66 DPS) | yes | Stone Guard's Leather Pants (220859, -0.94 DPS, sim-verified) [vendor]; Triprunner Dungarees (9624, -1.38 DPS) [quest]; Basilisk Hide Pants (1718, -1.42 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 50.8 attack_power points (3.89 DPS) | yes | Skulker's Leather Boots (252469, -0.03 DPS) [crafted]; Albino Crocscale Boots (17728, -0.80 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.96 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.3 attack_power points (2.16 DPS) | yes | Ring of the Underwood (2951, -0.28 DPS) [world_drop]; White Bone Band (11862, -0.33 DPS) [quest]; Thunderbrow Ring (13097, -0.35 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.8 attack_power points (1.90 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; White Bone Band (11862, -0.06 DPS) [quest]; Thunderbrow Ring (13097, -0.08 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (176.0 DPS) | yes | Molten Heart of the Mountain (249470, -4.55 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (176.0 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (176.0 DPS) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, -8.21 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0000000000000000-50212032123012123001-5550230000000000)

Set DPS (verified): 267.5. Weights run: 4.7s. Verify run: 9.8s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.669 ± 0.049, crit=2.384 ± 0.071 per rating point (14 rating = 1%, 33.377 per %), hit=2.323 ± 0.148 per rating point (10 rating = 1%, 23.225 per %), melee_haste=10.121 ± 1.848

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 106.1 attack_power points (9.03 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -1.88 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 45.7 attack_power points (3.89 DPS) | yes | Woven Ivy Necklace (19159, -0.72 DPS) [quest]; Sunprism Pendant (20649, -1.16 DPS) [quest]; Ethereal Talisman (4430, -2.04 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (267.5 DPS) | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS) [pvp]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -6.57 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 60.0 attack_power points (5.11 DPS) | yes | Dark Phantom Cape (13122, -1.63 DPS, sim-verified) [world_drop]; Fel Cape (279269, -2.27 DPS) [crafted]; Blackflame Cape (13109, -2.38 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.3 attack_power points (13.22 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -1.48 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -2.52 DPS) [vendor] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (267.5 DPS) | yes | Skulker's Leather Bracers (252540, -0.05 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.13 DPS) [crafted]; Deepfury Bracers (13120, -6.48 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (267.5 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -4.52 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (267.5 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -4.52 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 109.6 attack_power points (9.33 DPS) | yes | Legionnaire's Dragonhide Leggings (227177, +0.00 DPS) [pvp]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Traveler's Leggings (8300, -2.71 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 63.0 attack_power points (5.36 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -7.43 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 37.4 attack_power points (3.18 DPS) | yes | Falcon's Hook (7552, -0.57 DPS) [dungeon]; Thunderbrow Ring (13097, -1.00 DPS) [world_drop]; White Bone Band (11862, -1.14 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 31.1 attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.04 DPS) [dungeon]; Thunderbrow Ring (13097, -0.47 DPS) [world_drop]; White Bone Band (11862, -0.60 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (267.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (267.5 DPS) | yes | Molten Heart of the Mountain (249470, -1.36 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Titans (12796) | Blacksmithing [crafted] | sim-verified (267.5 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -20.64 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (267.5 DPS) | yes | Idol of the Moon (23197, -9.76 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; main_hand: Hammer of the Titans; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0500200000000000-50230032023011125001-5050030000000000)

Set DPS (verified): 617.2. Weights run: 4.5s. Verify run: 9.4s. 1571 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=3.135 ± 0.062, crit=2.625 ± 0.085 per rating point (14 rating = 1%, 36.751 per %), hit=2.636 ± 0.193 per rating point (10 rating = 1%, 26.361 per %), melee_haste=16.255 ± 2.324

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 121.0 attack_power points (18.62 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -2.22 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 51.8 attack_power points (7.97 DPS) | yes | Woven Ivy Necklace (19159, -1.84 DPS, sim-verified) [quest]; Sunprism Pendant (20649, -2.18 DPS) [quest]; Ethereal Talisman (4430, -4.34 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (617.2 DPS) | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS) [pvp]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -14.66 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 67.0 attack_power points (10.32 DPS) | yes | Dark Phantom Cape (13122, -1.95 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.53 DPS) [world_drop]; Fel Cape (279269, -4.66 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 179.1 attack_power points (27.57 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -4.33 DPS) [vendor]; Legionnaire's Dragonhide Chestpiece (227179, -6.36 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (617.2 DPS) | yes | Skulker's Leather Bracers (252540, -0.52 DPS) [crafted]; Pridelord Bands (14672, -0.72 DPS) [world_drop]; Deepfury Bracers (13120, -14.15 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (617.2 DPS) | yes | Blood Guard's Dragonhide Grips (227180, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -10.55 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (617.2 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -10.55 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 118.0 attack_power points (18.17 DPS) | yes | Legionnaire's Dragonhide Leggings (227177, +0.00 DPS) [pvp]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Traveler's Leggings (8300, -4.76 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 73.7 attack_power points (11.35 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -16.43 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.9 attack_power points (6.76 DPS) | yes | Falcon's Hook (7552, -1.40 DPS) [dungeon]; Thunderbrow Ring (13097, -2.60 DPS) [world_drop]; Disengagement Ring (276202, -2.90 DPS) [vendor] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 35.7 attack_power points (5.50 DPS) | yes | Falcon's Hook (7552, -0.14 DPS) [dungeon]; Thunderbrow Ring (13097, -1.35 DPS) [world_drop]; Disengagement Ring (276202, -1.64 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (617.2 DPS) | yes | Molten Heart of the Mountain (249470, -8.56 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (617.2 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Hammer of the Titans (12796) | Blacksmithing [crafted] | sim-verified (617.2 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -40.99 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (617.2 DPS) | yes | Idol of the Moon (23197, -12.69 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Ring of the Underwood; main_hand: Hammer of the Titans; ranged: Idol of the Dream

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

