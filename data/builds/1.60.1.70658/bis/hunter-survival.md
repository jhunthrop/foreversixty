# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 72.6. Weights run: 3.7s. Verify run: 5.0s. 227 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.068 ± 0.016, strength=1.210 ± 0.002, crit=0.779 ± 0.017 per rating point (14 rating = 1%, 10.906 per %), hit=1.140 ± 0.041 per rating point (10 rating = 1%, 11.405 per %), melee_haste=7.736 ± 1.342

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.5 attack_power points (0.72 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.1 attack_power points (0.18 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.3 attack_power points (0.45 DPS) | yes | Slime-encrusted Pads (6461, -0.44 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.4 attack_power points (0.54 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.17 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 22.8 attack_power points (0.99 DPS) | yes | Brawler's Leather Armor (252490, -0.10 DPS) [crafted]; Prospector's Chestpiece (14562, -0.35 DPS) [world_drop]; Defender's Leather Armor (252434, -0.35 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.7 attack_power points (0.47 DPS) | yes | Forest Leather Bracers (3202, -0.02 DPS) [world_drop]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (72.6 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.01 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Brawler's Leather Belt (252428, -0.21 DPS) [crafted]; Dusty Belt (279897, -0.33 DPS) [quest]; Deviate Scale Belt (6468, -2.66 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (72.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.15 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (72.6 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.07 DPS, sim-verified) [world_drop] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 4.8 attack_power points (0.21 DPS) | yes | The 1 Ring (8350, -0.07 DPS) [world]; Ring of the Moon (12052, -0.16 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.1 attack_power points (0.18 DPS) | yes | The 1 Ring (8350, -0.04 DPS) [world]; Ring of the Moon (12052, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (10.87 DPS) | yes | Night Reaver (1318, +0.00 DPS) [dungeon]; Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (10.34 DPS) | yes | Cruel Barb (5191, -0.91 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.3 attack_power points (0.36 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Owlsight Rifle (15205, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Lavishly Jeweled Ring; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 227, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 98.0. Weights run: 4.0s. Verify run: 2.9s. 434 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.833 ± 0.014, strength=1.210 ± 0.002, crit=0.882 ± 0.020 per rating point (14 rating = 1%, 12.355 per %), hit=1.229 ± 0.054 per rating point (10 rating = 1%, 12.295 per %), melee_haste=not significant (4.788 ± 1.299)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.3 attack_power points (0.94 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.19 DPS) [crafted]; Humbert's Helm (4724, -0.28 DPS) [world] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.2 attack_power points (0.62 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.01 DPS) [world]; River Pride Choker (13087, -0.37 DPS) [world_drop]; Erudite's Amulet (277204, -0.43 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.2 attack_power points (1.34 DPS) | yes | Mantle of Thieves (2264, -0.28 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.56 DPS) [crafted]; Bristlebark Amice (14573, -0.59 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.5 attack_power points (0.84 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.22 DPS) [pvp]; Fenrus' Hide (6340, -0.28 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 25.7 attack_power points (1.31 DPS) | yes | Brawler's Leather Tunic (252508, -0.19 DPS) [crafted]; Tunic of Westfall (2041, -0.28 DPS) [quest]; Brawler's Leather Armor (252490, -0.35 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.8 attack_power points (0.81 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.25 DPS) [world_drop]; Demonhide Bracers (270033, -0.25 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.5 attack_power points (0.99 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Cartilage Shapers (273054, -0.15 DPS) [world] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 23.8 attack_power points (1.21 DPS) | yes | Prowler's Leather Belt (252459, +0.00 DPS, sim-verified) [crafted]; Warden's Leather Belt (252460, -0.28 DPS) [crafted]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.33 DPS) | yes | Troll's Bane Leggings (13114, -0.02 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.05 DPS) [crafted]; Petrolspill Leggings (9509, -1.34 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.3 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Disjointed Shoes (277226, -0.32 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 15.2 attack_power points (0.78 DPS) | yes | Defias Renegade Ring (1076, -0.31 DPS) [dungeon]; Monkey Ring (6748, -0.31 DPS) [quest]; Seal of Wrynn (2933, -0.31 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 11.0 attack_power points (0.56 DPS) | yes | Defias Renegade Ring (1076, -0.09 DPS) [dungeon]; Monkey Ring (6748, -0.09 DPS) [quest]; Seal of Wrynn (2933, -0.10 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.44 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (98.0 DPS) | yes | Shoni's Disarming Tool (9608, -4.78 DPS) [quest]; Swinetusk Shank (6691, -16.48 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 11.0 attack_power points (0.56 DPS) | yes | Double-barreled Shotgun (2098, -0.10 DPS) [world_drop]; Gravespike Repeater (273033, -0.15 DPS) [world]; BKP "Sparrow" Smallbore (3042, -0.19 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Skulker's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Ring of Precision; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 141.0. Weights run: 4.1s. Verify run: 3.4s. 690 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.849 ± 0.012, strength=1.210 ± 0.002, crit=0.905 ± 0.019 per rating point (14 rating = 1%, 12.664 per %), hit=1.370 ± 0.071 per rating point (10 rating = 1%, 13.697 per %), melee_haste=6.765 ± 1.061

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.6 attack_power points (2.39 DPS) | yes | Barbaric Iron Helm (7915, -0.47 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.60 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.68 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.2 attack_power points (0.87 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.02 DPS) [world]; River Pride Choker (13087, -0.53 DPS) [world_drop]; Gazlowe's Charm (13088, -0.53 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.4 attack_power points (1.88 DPS) | yes | Flintrock Shoulders (7755, -0.13 DPS) [dungeon]; Nightscape Shoulders (8192, -0.43 DPS) [crafted]; Mantle of Thieves (2264, -0.56 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 18.4 attack_power points (1.30 DPS) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Imperial Cloak (6432, -0.25 DPS) [dungeon]; Parachute Cloak (10518, -0.25 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 44.4 attack_power points (3.15 DPS) | yes | Wolffear Harness (13110, -0.67 DPS, sim-verified) [world_drop]; Kolkar Marauder Chain (6773, -1.03 DPS) [quest]; Tough Scorpid Breastplate (8203, -1.18 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.42 DPS) | yes | Hawkeye's Bracers (14590, -0.29 DPS) [world_drop]; Ravager's Armguards (14770, -0.29 DPS) [world_drop]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 33.0 attack_power points (2.35 DPS) | yes | Gloves of Holy Might (867, -0.02 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.07 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.26 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 27.5 attack_power points (1.96 DPS) | yes | Skulker's Leather Belt (252520, -0.26 DPS) [crafted]; Scorpashi Sash (14652, -0.30 DPS) [world_drop]; Officer's Belt (250556, -0.31 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 38.8 attack_power points (2.76 DPS) | yes | Triprunner Dungarees (9624, -0.14 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.89 DPS) [crafted]; Ferine Leggings (6690, -0.91 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.4 attack_power points (2.09 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.22 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 20.9 attack_power points (1.49 DPS) | yes | Mark of Kern (2262, -0.06 DPS) [dungeon]; Assault Band (13095, -0.06 DPS) [world_drop]; Thunderbrow Ring (13097, -0.40 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 20.3 attack_power points (1.44 DPS) | yes | Mark of Kern (2262, -0.02 DPS) [dungeon]; Assault Band (13095, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.36 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (141.0 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -5.72 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (31.42 DPS) | yes | Shoni's Disarming Tool (9608, -15.45 DPS) [quest]; Dazzling Longsword (869, -22.55 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (141.0 DPS) | yes | Swiftwind (13038, -0.08 DPS) [world_drop]; Monolithic Bow (9426, -0.08 DPS) [dungeon]; Bow of Searing Arrows (2825, -0.99 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band; 5008 Quicksilver Ring

### Band 50 (dwarf, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 183.2. Weights run: 4.1s. Verify run: 3.9s. 879 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.761 ± 0.011, strength=1.210 ± 0.002, crit=0.978 ± 0.019 per rating point (14 rating = 1%, 13.697 per %), hit=1.292 ± 0.082 per rating point (10 rating = 1%, 12.925 per %), melee_haste=8.958 ± 1.260

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 32.7 attack_power points (2.56 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.22 DPS) [crafted]; Scorpashi Skullcap (14658, -0.34 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.9 attack_power points (2.27 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.13 DPS, sim-verified) [world]; Kaleidoscope Chain (13084, -1.34 DPS) [world_drop]; River Pride Choker (13087, -1.89 DPS) [world_drop] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 30.3 attack_power points (2.37 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.04 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.09 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 26.4 attack_power points (2.07 DPS) | yes | Blackflame Cape (13109, -0.41 DPS) [world_drop]; Pridelord Cape (14673, -0.54 DPS) [world_drop]; Serpentskin Cloak (8259, -0.55 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 45.0 attack_power points (3.53 DPS) | yes | Blazewind Breastplate (11193, -0.07 DPS) [quest]; Icemail Jerkin (1981, -0.22 DPS) [world_drop]; Knight's Chain Armor (220828, -0.38 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.3 attack_power points (2.45 DPS) | yes | Bracers of the Stone Princess (17714, -0.26 DPS) [dungeon]; Arena Bands (18711, -0.26 DPS) [world]; Skulker's Leather Bracers (252540, -0.54 DPS) [crafted] |
| hands | Skulker's Leather Gauntlets (252548) | Leatherworking [crafted] | 35.5 attack_power points (2.79 DPS) | yes | Gloves of Holy Might (867, -0.14 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.22 DPS) [crafted]; Serpentskin Gloves (8260, -0.22 DPS) [world_drop] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 39.2 attack_power points (3.07 DPS) | yes | Prowler's Leather Waistguard (252473, -0.09 DPS) [crafted]; Ogron's Sash (13117, -0.98 DPS) [world_drop]; Sagebrush Girdle (17778, -1.00 DPS) [quest] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 44.6 attack_power points (3.49 DPS) | yes | Serpentskin Leggings (8262, -0.05 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.49 DPS) [vendor]; Basilisk Hide Pants (1718, -0.59 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 36.2 attack_power points (2.84 DPS) | yes | Albino Crocscale Boots (17728, -0.08 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.09 DPS) [crafted]; Fleetfoot Greaves (11627, -0.22 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.7 attack_power points (1.93 DPS) | yes | Mark of Kern (2262, -0.36 DPS) [dungeon]; Assault Band (13095, -0.36 DPS) [world_drop]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 20.0 attack_power points (1.57 DPS) | yes | Mark of Kern (2262, -0.00 DPS) [dungeon]; Assault Band (13095, -0.00 DPS) [world_drop]; Blackstone Ring (17713, -0.00 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (183.2 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (183.2 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-verified (183.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Cestus (250614, -1.93 DPS, sim-verified) [crafted] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 523.7 attack_power points (41.07 DPS) | yes | Julie's Dagger (6660, +0.00 DPS) [world_drop]; Claw of Celebras (17738, -3.26 DPS) [dungeon]; Shoni's Disarming Tool (9608, -23.45 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (183.2 DPS) | yes | Guttbuster (13139, -0.54 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.55 DPS) [world_drop]; Dark Iron Rifle (16004, -1.88 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Skulker's Leather Gauntlets; waist: Skulker's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thrash Blade; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 256.0. Weights run: 4.0s. Verify run: 11.0s. 1749 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.966 ± 0.019, strength=1.210 ± 0.002, crit=1.809 ± 0.037 per rating point (14 rating = 1%, 25.328 per %), hit=2.634 ± 0.153 per rating point (10 rating = 1%, 26.341 per %), melee_haste=15.325 ± 2.330

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (256.0 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -8.89 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (2.20 DPS) | yes | Sunprism Pendant (20649, -0.56 DPS) [quest]; Kaleidoscope Chain (13084, -1.32 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.36 DPS) [world] |
| shoulder | Beaststalker's Mantle (16679) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-verified (256.0 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Devout Mantle (16695, +0.00 DPS) [dungeon]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 49.5 attack_power points (3.44 DPS) | yes | Sergeant Major's Cape (16337, -1.68 DPS) [pvp]; Fel Cape (279269, -1.68 DPS) [crafted]; Dark Phantom Cape (13122, -1.88 DPS, sim-verified) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.4 attack_power points (8.66 DPS) | yes | Field Marshal's Chain Armor (231563, -0.55 DPS) [vendor]; Timbermaw Tunic (252484, -1.53 DPS, sim-verified) [crafted]; Knight-Captain's Chain Armor (227089, -1.62 DPS) [vendor] |
| wrist | Beaststalker's Bindings (16681) | Stratholme: Fleshflayer Ghoul [dungeon] | sim-verified (256.0 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Devout Bracers (16697, -2.40 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (256.0 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -2.85 DPS, sim-verified) [crafted] |
| waist | Beaststalker's Belt (16680) | Blackrock Spire: Smolderthorn Headhunter [dungeon] | sim-verified (256.0 DPS) | yes | Heavy Timbermaw Belt (19043, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -2.34 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (256.0 DPS) | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS) [vendor]; Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Devilsaur Leggings (15062, -3.58 DPS, sim-verified) [crafted] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (256.0 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Scalegut Treaders (275618, -1.64 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 27.5 attack_power points (1.92 DPS) | yes | Falcon's Hook (7552, -0.43 DPS) [dungeon]; Mark of Kern (2262, -0.52 DPS) [dungeon]; Blackstone Ring (17713, -0.52 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 22.1 attack_power points (1.54 DPS) | yes | Falcon's Hook (7552, -0.05 DPS) [dungeon]; Mark of Kern (2262, -0.14 DPS) [dungeon]; Blackstone Ring (17713, -0.14 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (256.0 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (256.0 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (256.0 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -2.30 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (256.0 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -9.16 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -42.64 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (256.0 DPS) | yes | Eaglehorn Long Bow (13023, -0.55 DPS) [world_drop]; Guttbuster (13139, -0.57 DPS) [world_drop]; Dark Iron Rifle (16004, -2.04 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Beaststalker's Mantle; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Beaststalker's Gloves; waist: Beaststalker's Belt; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1749, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 686.6. Weights run: 3.5s. Verify run: 7.1s. 1749 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.200 ± 0.021, strength=1.210 ± 0.002, crit=2.524 ± 0.044 per rating point (14 rating = 1%, 35.330 per %), hit=4.085 ± 0.325 per rating point (10 rating = 1%, 40.846 per %), melee_haste=17.800 ± 4.170

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (+8.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -8.42 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.7 attack_power points (4.88 DPS) | yes | Sunprism Pendant (20649, -1.16 DPS) [quest]; Kaleidoscope Chain (13084, -2.96 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -3.19 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 56.6 attack_power points (7.96 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 53.0 attack_power points (7.46 DPS) | yes | Dark Phantom Cape (13122, -2.81 DPS) [world_drop]; Fel Cape (279269, -3.40 DPS, sim-verified) [crafted]; Sergeant Major's Cape (16337, -3.62 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.6 attack_power points (22.17 DPS) | yes | Field Marshal's Chain Armor (231563, -3.81 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -4.79 DPS) [vendor]; Timbermaw Tunic (252484, -7.01 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.8 attack_power points (5.32 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Arena Wristguards (18709, -0.35 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 63.3 attack_power points (8.91 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -13.93 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 56.2 attack_power points (7.91 DPS) | yes | Marshal's Chain Girdle (16464, -1.49 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.53 DPS) [crafted]; Prowler's Leather Waistguard (252473, -1.81 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 81.3 attack_power points (11.44 DPS) | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS) [vendor]; Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Warbear Woolies (15065, -13.11 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 58.5 attack_power points (8.22 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.8 attack_power points (4.33 DPS) | yes | Falcon's Hook (7552, -1.04 DPS) [dungeon]; Mark of Kern (2262, -1.52 DPS) [dungeon]; Blackstone Ring (17713, -1.52 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.4 attack_power points (3.44 DPS) | yes | Falcon's Hook (7552, -0.14 DPS) [dungeon]; Mark of Kern (2262, -0.62 DPS) [dungeon]; Blackstone Ring (17713, -0.62 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Annihilator (12798, +0.00 DPS) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (+64.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -20.06 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -64.63 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -1.24 DPS) [world_drop]; Guttbuster (13139, -1.35 DPS) [world_drop]; Dark Iron Rifle (16004, -4.05 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1749, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 73.0. Weights run: 3.7s. Verify run: 4.8s. 223 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.068 ± 0.016, strength=1.210 ± 0.002, crit=0.779 ± 0.017 per rating point (14 rating = 1%, 10.906 per %), hit=1.140 ± 0.041 per rating point (10 rating = 1%, 11.405 per %), melee_haste=7.736 ± 1.342

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.5 attack_power points (0.72 DPS) | yes | Defender's Leather Hood (252447, -0.26 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 4.1 attack_power points (0.18 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.3 attack_power points (0.45 DPS) | yes | Slime-encrusted Pads (6461, -0.45 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.4 attack_power points (0.54 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.17 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (73.0 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.95 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.3 attack_power points (0.45 DPS) | yes | Bristlebark Bindings (14569, -0.07 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Spare Part Bindings (279875, -0.18 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 17.2 attack_power points (0.75 DPS) | yes | Bristlebark Gloves (14572, -0.18 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.18 DPS) [crafted]; Serpent Gloves (5970, -0.21 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Brawler's Leather Belt (252428, -0.21 DPS) [crafted]; Murloc Scale Belt (5780, -0.41 DPS) [crafted]; Deviate Scale Belt (6468, -2.61 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.10 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (73.0 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.01 DPS, sim-verified) [world_drop] |
| finger1 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 6.2 attack_power points (0.27 DPS) | yes | Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Loop of Sacrifice (281673, -0.11 DPS) [quest]; The 1 Ring (8350, -0.13 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 4.8 attack_power points (0.21 DPS) | yes | Lavishly Jeweled Ring (1156, -0.03 DPS) [dungeon]; Loop of Sacrifice (281673, -0.05 DPS) [quest]; The 1 Ring (8350, -0.07 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (10.87 DPS) | yes | Crescent Staff (6505, +0.00 DPS) [quest]; Living Root (6631, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (10.34 DPS) | yes | Wingblade (6504, -6.16 DPS, sim-verified) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.3 attack_power points (0.36 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Privateer Musket (5309, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Bounty Hunter's Ring; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 223, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (troll, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 99.1. Weights run: 4.0s. Verify run: 2.8s. 431 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.833 ± 0.014, strength=1.210 ± 0.002, crit=0.882 ± 0.020 per rating point (14 rating = 1%, 12.355 per %), hit=1.229 ± 0.054 per rating point (10 rating = 1%, 12.295 per %), melee_haste=not significant (4.788 ± 1.299)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.3 attack_power points (0.94 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.19 DPS) [crafted]; Humbert's Helm (4724, -0.28 DPS) [world] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.2 attack_power points (0.62 DPS) | yes | Archmagister's Faceted Pendant (273049, +0.00 DPS, sim-verified) [world]; River Pride Choker (13087, -0.37 DPS) [world_drop]; Erudite's Amulet (277204, -0.43 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.2 attack_power points (1.34 DPS) | yes | Mantle of Thieves (2264, -0.31 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.56 DPS) [crafted]; Bristlebark Amice (14573, -0.59 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.5 attack_power points (0.84 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Fenrus' Hide (6340, -0.28 DPS) [dungeon]; Swiftrunner Cape (6745, -0.28 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 25.7 attack_power points (1.31 DPS) | yes | Brawler's Leather Tunic (252508, -0.19 DPS) [crafted]; Panther Armor (6670, -0.34 DPS) [quest]; Brawler's Leather Armor (252490, -0.35 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.8 attack_power points (0.81 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.25 DPS) [world_drop]; Cultist's Armguards (270032, -0.30 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.5 attack_power points (0.99 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Cartilage Shapers (273054, -0.15 DPS) [world] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 23.8 attack_power points (1.21 DPS) | yes | Prowler's Leather Belt (252459, -0.10 DPS) [crafted]; Deftkin Belt (16659, -0.23 DPS) [quest]; Warden's Leather Belt (252460, -0.28 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.33 DPS) | yes | Troll's Bane Leggings (13114, -0.02 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.05 DPS) [crafted]; Petrolspill Leggings (9509, -1.39 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.3 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Warsong Boots (16977, -0.19 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 15.2 attack_power points (0.78 DPS) | yes | Ring of Precision (1491, -0.21 DPS) [dungeon]; Defias Renegade Ring (1076, -0.31 DPS) [dungeon]; Monkey Ring (6748, -0.31 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 12.2 attack_power points (0.62 DPS) | yes | Ring of Precision (1491, -0.06 DPS) [dungeon]; Defias Renegade Ring (1076, -0.15 DPS) [dungeon]; Monkey Ring (6748, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.44 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (99.1 DPS) | yes | Swinetusk Shank (6691, -16.88 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 11.0 attack_power points (0.56 DPS) | yes | Double-barreled Shotgun (2098, -0.10 DPS) [world_drop]; Gravespike Repeater (273033, -0.15 DPS) [world]; BKP "Sparrow" Smallbore (3042, -0.19 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Skulker's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 431, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 140.8. Weights run: 4.1s. Verify run: 3.3s. 678 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.849 ± 0.012, strength=1.210 ± 0.002, crit=0.905 ± 0.019 per rating point (14 rating = 1%, 12.664 per %), hit=1.370 ± 0.071 per rating point (10 rating = 1%, 13.697 per %), melee_haste=6.765 ± 1.061

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.6 attack_power points (2.39 DPS) | yes | Barbaric Iron Helm (7915, -0.43 DPS) [crafted]; Hawkeye's Helm (14591, -0.60 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.68 DPS) [world_drop] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 13.4 attack_power points (0.96 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; River Pride Choker (13087, -0.61 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.4 attack_power points (1.88 DPS) | yes | Flintrock Shoulders (7755, -0.13 DPS) [dungeon]; Nightscape Shoulders (8192, -0.43 DPS) [crafted]; Mantle of Thieves (2264, -0.56 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.6 attack_power points (1.18 DPS) | yes | Imperial Cloak (6432, -0.13 DPS) [dungeon]; Parachute Cloak (10518, -0.13 DPS) [crafted]; Tigerstrike Mantle (13108, -0.13 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 44.4 attack_power points (3.15 DPS) | yes | Wolffear Harness (13110, -0.82 DPS, sim-verified) [world_drop]; Kolkar Marauder Chain (6773, -1.03 DPS) [quest]; Tough Scorpid Breastplate (8203, -1.18 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.42 DPS) | yes | Hawkeye's Bracers (14590, -0.29 DPS) [world_drop]; Ravager's Armguards (14770, -0.29 DPS) [world_drop]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 33.0 attack_power points (2.35 DPS) | yes | Gloves of Holy Might (867, -0.02 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.07 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.26 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 27.5 attack_power points (1.96 DPS) | yes | Skulker's Leather Belt (252520, -0.26 DPS) [crafted]; Scorpashi Sash (14652, -0.30 DPS) [world_drop]; Officer's Belt (250556, -0.31 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 38.8 attack_power points (2.76 DPS) | yes | Triprunner Dungarees (9624, -0.14 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.89 DPS) [crafted]; Ferine Leggings (6690, -0.91 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.4 attack_power points (2.09 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.22 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 20.9 attack_power points (1.49 DPS) | yes | Mark of Kern (2262, -0.06 DPS) [dungeon]; Assault Band (13095, -0.06 DPS) [world_drop]; Thunderbrow Ring (13097, -0.40 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 20.3 attack_power points (1.44 DPS) | yes | Mark of Kern (2262, -0.02 DPS) [dungeon]; Assault Band (13095, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.36 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (140.8 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -5.30 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (31.42 DPS) | yes | Dazzling Longsword (869, -24.42 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (140.8 DPS) | yes | Swiftwind (13038, -0.08 DPS) [world_drop]; Monolithic Bow (9426, -0.08 DPS) [dungeon]; Bow of Searing Arrows (2825, -0.99 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 678, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 189.1. Weights run: 4.1s. Verify run: 3.7s. 868 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.761 ± 0.011, strength=1.210 ± 0.002, crit=0.978 ± 0.019 per rating point (14 rating = 1%, 13.697 per %), hit=1.292 ± 0.082 per rating point (10 rating = 1%, 12.925 per %), melee_haste=8.958 ± 1.260

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 32.7 attack_power points (2.56 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.09 DPS) [quest]; Helm of Fire (8348, -0.22 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.9 attack_power points (2.27 DPS) | yes | Woven Ivy Necklace (19159, -0.46 DPS) [quest]; Ethereal Talisman (4430, -1.24 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.33 DPS) [world] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 30.3 attack_power points (2.37 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.04 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.09 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 26.4 attack_power points (2.07 DPS) | yes | Pridelord Cape (14673, -0.54 DPS) [world_drop]; Serpentskin Cloak (8259, -0.55 DPS) [world_drop]; Blackflame Cape (13109, -0.71 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 45.0 attack_power points (3.53 DPS) | yes | Blazewind Breastplate (11193, -0.07 DPS) [quest]; Icemail Jerkin (1981, -0.22 DPS) [world_drop]; Stone Guard's Chain Armor (220827, -0.38 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.3 attack_power points (2.45 DPS) | yes | Bracers of the Stone Princess (17714, -0.26 DPS) [dungeon]; Arena Bands (18711, -0.26 DPS) [world]; Skulker's Leather Bracers (252540, -0.54 DPS) [crafted] |
| hands | Skulker's Leather Gauntlets (252548) | Leatherworking [crafted] | 35.5 attack_power points (2.79 DPS) | yes | Gloves of Holy Might (867, -0.14 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.22 DPS) [crafted]; Serpentskin Gloves (8260, -0.22 DPS) [world_drop] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 39.2 attack_power points (3.07 DPS) | yes | Prowler's Leather Waistguard (252473, -0.09 DPS) [crafted]; Ogron's Sash (13117, -0.98 DPS) [world_drop]; Sagebrush Girdle (17778, -1.00 DPS) [quest] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 43.9 attack_power points (3.44 DPS) | yes | Stone Guard's Chain Legplates (220833, -0.44 DPS) [vendor]; Basilisk Hide Pants (1718, -0.54 DPS) [world_drop]; Triprunner Dungarees (9624, -0.67 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 36.2 attack_power points (2.84 DPS) | yes | Albino Crocscale Boots (17728, -0.08 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.09 DPS) [crafted]; Fleetfoot Greaves (11627, -0.22 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.7 attack_power points (1.93 DPS) | yes | Ring of the Underwood (2951, -0.36 DPS) [world_drop]; Mark of Kern (2262, -0.36 DPS) [dungeon]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.88 DPS) | yes | Ring of the Underwood (2951, -0.31 DPS) [world_drop]; Mark of Kern (2262, -0.31 DPS) [dungeon]; Blackstone Ring (17713, -0.31 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (189.1 DPS) | yes | Molten Heart of the Mountain (249470, -4.08 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (189.1 DPS) | yes | Molten Heart of the Mountain (249470, -0.98 DPS, sim-verified) [crafted] |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-verified (189.1 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Cestus (250614, -2.64 DPS, sim-verified) [crafted] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 523.7 attack_power points (41.07 DPS) | yes | Julie's Dagger (6660, +0.00 DPS) [world_drop]; Claw of Celebras (17738, -3.26 DPS) [dungeon]; White Bone Shredder (11863, -5.61 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (189.1 DPS) | yes | Guttbuster (13139, -0.54 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.55 DPS) [world_drop]; Dark Iron Rifle (16004, -2.44 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Skulker's Leather Gauntlets; waist: Skulker's Leather Waistguard; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thrash Blade; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 868, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 259.9. Weights run: 4.0s. Verify run: 10.5s. 1746 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.966 ± 0.019, strength=1.210 ± 0.002, crit=1.809 ± 0.037 per rating point (14 rating = 1%, 25.328 per %), hit=2.634 ± 0.153 per rating point (10 rating = 1%, 26.341 per %), melee_haste=15.325 ± 2.330

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (259.9 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -9.98 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (2.20 DPS) | yes | Woven Ivy Necklace (19159, -0.46 DPS) [quest]; Sunprism Pendant (20649, -0.56 DPS) [quest]; Ethereal Talisman (4430, -1.23 DPS) [quest] |
| shoulder | Beaststalker's Mantle (16679) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-verified (259.9 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Devout Mantle (16695, -6.14 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 49.5 attack_power points (3.44 DPS) | yes | Dark Phantom Cape (13122, -1.56 DPS, sim-verified) [world_drop]; Fel Cape (279269, -1.68 DPS) [crafted]; Blackflame Cape (13109, -1.80 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.4 attack_power points (8.66 DPS) | yes | Warlord's Chain Armor (231566, -0.55 DPS) [vendor]; Legionnaire's Chain Armor (227083, -1.62 DPS) [vendor]; Timbermaw Tunic (252484, -2.30 DPS, sim-verified) [crafted] |
| wrist | Beaststalker's Bindings (16681) | Stratholme: Fleshflayer Ghoul [dungeon] | sim-verified (259.9 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Devout Bracers (16697, -8.03 DPS, sim-verified) [dungeon] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 53.3 attack_power points (3.71 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; General's Chain Grips (231569, +0.00 DPS) [vendor] |
| waist | Beaststalker's Belt (16680) | Blackrock Spire: Smolderthorn Headhunter [dungeon] | sim-verified (259.9 DPS) | yes | Might of the Timbermaw (19044, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Devout Belt (16696, -7.26 DPS, sim-verified) [dungeon] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (259.9 DPS) | yes | Legionnaire's Chain Legplates (227079, +0.00 DPS) [vendor]; General's Chain Legplates (231567, +0.00 DPS) [vendor]; Devilsaur Leggings (15062, -3.94 DPS, sim-verified) [crafted] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (259.9 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Scalegut Treaders (275618, -7.42 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 27.5 attack_power points (1.92 DPS) | yes | Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.43 DPS) [dungeon]; Blackstone Ring (17713, -0.52 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.67 DPS) | yes | Ring of the Underwood (2951, -0.13 DPS) [world_drop]; Falcon's Hook (7552, -0.19 DPS) [dungeon]; Blackstone Ring (17713, -0.28 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (259.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (259.9 DPS) | yes | Molten Heart of the Mountain (249470, -1.42 DPS, sim-verified) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (259.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -1.96 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (259.9 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -9.16 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -45.27 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (259.9 DPS) | yes | Eaglehorn Long Bow (13023, -0.55 DPS) [world_drop]; Guttbuster (13139, -0.57 DPS) [world_drop]; Dark Iron Rifle (16004, -2.78 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Beaststalker's Mantle; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Devilsaur Gauntlets; waist: Beaststalker's Belt; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1746, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 702.3. Weights run: 3.5s. Verify run: 7.1s. 1746 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.200 ± 0.021, strength=1.210 ± 0.002, crit=2.524 ± 0.044 per rating point (14 rating = 1%, 35.330 per %), hit=4.085 ± 0.325 per rating point (10 rating = 1%, 40.846 per %), melee_haste=17.800 ± 4.170

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (+9.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -9.13 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.7 attack_power points (4.88 DPS) | yes | Woven Ivy Necklace (19159, -1.07 DPS) [quest]; Sunprism Pendant (20649, -1.16 DPS) [quest]; Ethereal Talisman (4430, -2.79 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 56.6 attack_power points (7.96 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 53.0 attack_power points (7.46 DPS) | yes | Fel Cape (279269, -2.49 DPS) [crafted]; Dark Phantom Cape (13122, -2.81 DPS) [world_drop]; Blackflame Cape (13109, -3.74 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.6 attack_power points (22.17 DPS) | yes | Warlord's Chain Armor (231566, -3.81 DPS) [vendor]; Legionnaire's Chain Armor (227083, -4.79 DPS) [vendor]; Timbermaw Tunic (252484, -6.54 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.8 attack_power points (5.32 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Arena Wristguards (18709, -0.35 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 63.3 attack_power points (8.91 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -8.80 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 56.2 attack_power points (7.91 DPS) | yes | General's Chain Girdle (16572, -1.49 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.53 DPS) [crafted]; Prowler's Leather Waistguard (252473, -1.81 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 81.3 attack_power points (11.44 DPS) | yes | Legionnaire's Chain Legplates (227079, +0.00 DPS) [vendor]; General's Chain Legplates (231567, +0.00 DPS) [vendor]; Warbear Woolies (15065, -9.32 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 58.5 attack_power points (8.22 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.8 attack_power points (4.33 DPS) | yes | White Bone Band (11862, -0.96 DPS) [quest]; Falcon's Hook (7552, -1.04 DPS) [dungeon]; Blackstone Ring (17713, -1.52 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.4 attack_power points (3.44 DPS) | yes | White Bone Band (11862, -0.06 DPS) [quest]; Falcon's Hook (7552, -0.14 DPS) [dungeon]; Blackstone Ring (17713, -0.62 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+10.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -4.14 DPS, sim-verified) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Annihilator (12798, +0.00 DPS) [crafted]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (+67.7 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -20.06 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -67.71 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -1.24 DPS) [world_drop]; Guttbuster (13139, -1.35 DPS) [world_drop]; Dark Iron Rifle (16004, -3.88 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger2: Ring of the Underwood; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1746, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

