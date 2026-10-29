# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 73.2. Weights run: 3.7s. Verify run: 5.0s. 217 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.068 ± 0.016, strength=1.210 ± 0.002, crit=0.779 ± 0.017 per rating point (14 rating = 1%, 10.906 per %), hit=1.140 ± 0.041 per rating point (10 rating = 1%, 11.405 per %), melee_haste=7.736 ± 1.342

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.5 attack_power points (0.72 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.4 attack_power points (0.54 DPS) | yes | Erudite's Amulet (277204, -0.35 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.3 attack_power points (0.45 DPS) | yes | Slime-encrusted Pads (6461, -0.44 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.4 attack_power points (0.54 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.17 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 22.8 attack_power points (0.99 DPS) | yes | Brawler's Leather Armor (252490, -0.10 DPS) [crafted]; Prospector's Chestpiece (14562, -0.35 DPS) [world_drop]; Defender's Leather Armor (252434, -0.35 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.7 attack_power points (0.47 DPS) | yes | Forest Leather Bracers (3202, -0.02 DPS) [world_drop]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.2 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.01 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Brawler's Leather Belt (252428, -0.21 DPS) [crafted]; Dusty Belt (279897, -0.33 DPS) [quest]; Deviate Scale Belt (6468, -2.67 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.2 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.15 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (73.2 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.06 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.1 attack_power points (0.57 DPS) | yes | Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; The 1 Ring (8350, -0.43 DPS) [world]; Ring of the Moon (12052, -0.52 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 4.8 attack_power points (0.21 DPS) | yes | Lavishly Jeweled Ring (1156, -0.03 DPS) [dungeon]; The 1 Ring (8350, -0.07 DPS) [world]; Ring of the Moon (12052, -0.16 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (10.87 DPS) | yes | Night Reaver (1318, +0.00 DPS) [dungeon]; Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (10.34 DPS) | yes | Cruel Barb (5191, -0.90 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.3 attack_power points (0.36 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Owlsight Rifle (15205, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 98.2. Weights run: 4.0s. Verify run: 2.9s. 397 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.833 ± 0.014, strength=1.210 ± 0.002, crit=0.882 ± 0.020 per rating point (14 rating = 1%, 12.355 per %), hit=1.229 ± 0.054 per rating point (10 rating = 1%, 12.295 per %), melee_haste=not significant (4.788 ± 1.299)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.3 attack_power points (0.94 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.19 DPS) [crafted]; Humbert's Helm (4724, -0.28 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 14.7 attack_power points (0.75 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; River Pride Choker (13087, -0.50 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.2 attack_power points (1.34 DPS) | yes | Mantle of Thieves (2264, -0.32 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.56 DPS) [crafted]; Bristlebark Amice (14573, -0.59 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.5 attack_power points (0.84 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.22 DPS) [pvp]; Fenrus' Hide (6340, -0.28 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 25.7 attack_power points (1.31 DPS) | yes | Brawler's Leather Tunic (252508, -0.19 DPS) [crafted]; Tunic of Westfall (2041, -0.28 DPS) [quest]; Brawler's Leather Armor (252490, -0.35 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.8 attack_power points (0.81 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.25 DPS) [world_drop]; Demonhide Bracers (270033, -0.25 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.5 attack_power points (0.99 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Cartilage Shapers (273054, -0.15 DPS) [world] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.23 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.01 DPS) [crafted]; Prowler's Leather Belt (252459, -0.11 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.33 DPS) | yes | Troll's Bane Leggings (13114, -0.02 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.05 DPS) [crafted]; Petrolspill Leggings (9509, -1.26 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.3 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.3 attack_power points (0.93 DPS) | yes | Ring of Precision (1491, -0.37 DPS) [dungeon]; Defias Renegade Ring (1076, -0.46 DPS) [dungeon]; Monkey Ring (6748, -0.46 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 15.2 attack_power points (0.78 DPS) | yes | Ring of Precision (1491, -0.21 DPS) [dungeon]; Defias Renegade Ring (1076, -0.31 DPS) [dungeon]; Monkey Ring (6748, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.44 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (98.2 DPS) | yes | Shoni's Disarming Tool (9608, -4.78 DPS) [quest]; Swinetusk Shank (6691, -16.79 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 11.0 attack_power points (0.56 DPS) | yes | Double-barreled Shotgun (2098, -0.10 DPS) [world_drop]; Booty Bay Bruiser's Buckshot (274748, -0.10 DPS) [vendor]; Gravespike Repeater (273033, -0.15 DPS) [world] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 142.9. Weights run: 4.1s. Verify run: 3.4s. 626 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.849 ± 0.012, strength=1.210 ± 0.002, crit=0.905 ± 0.019 per rating point (14 rating = 1%, 12.664 per %), hit=1.370 ± 0.071 per rating point (10 rating = 1%, 13.697 per %), melee_haste=6.765 ± 1.061

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.6 attack_power points (2.39 DPS) | yes | Barbaric Iron Helm (7915, -0.58 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.60 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.68 DPS) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 20.3 attack_power points (1.45 DPS) | yes | Kaleidoscope Chain (13084, -0.58 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.59 DPS) [world]; River Pride Choker (13087, -1.10 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.3 attack_power points (2.30 DPS) | yes | Forest Tracker Epaulets (2278, -0.42 DPS) [world_drop]; Flintrock Shoulders (7755, -0.55 DPS) [dungeon]; Nightscape Shoulders (8192, -0.85 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 18.4 attack_power points (1.30 DPS) | yes | Imperial Cloak (6432, -0.25 DPS) [dungeon]; Parachute Cloak (10518, -0.25 DPS) [crafted]; Hawkeye's Cloak (14593, -0.54 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 44.4 attack_power points (3.15 DPS) | yes | Wolffear Harness (13110, -0.96 DPS, sim-verified) [world_drop]; Kolkar Marauder Chain (6773, -1.03 DPS) [quest]; Tough Scorpid Breastplate (8203, -1.18 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.42 DPS) | yes | Hawkeye's Bracers (14590, -0.29 DPS) [world_drop]; Ravager's Armguards (14770, -0.29 DPS) [world_drop]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 33.0 attack_power points (2.35 DPS) | yes | Gloves of Holy Might (867, -0.02 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.07 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.26 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.13 DPS) | yes | Ogron's Sash (13117, -0.18 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.43 DPS) [rep]; Skulker's Leather Belt (252520, -0.43 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 38.8 attack_power points (2.76 DPS) | yes | Triprunner Dungarees (9624, -0.77 DPS, sim-verified) [quest]; Veteran's Silvered Chain Leggings (250523, -0.89 DPS) [crafted]; Ferine Leggings (6690, -0.91 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.4 attack_power points (2.09 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.22 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 24.5 attack_power points (1.74 DPS) | yes | Falcon's Hook (7552, -0.30 DPS) [dungeon]; Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 20.9 attack_power points (1.49 DPS) | yes | Falcon's Hook (7552, -0.05 DPS) [dungeon]; Mark of Kern (2262, -0.06 DPS) [dungeon]; Assault Band (13095, -0.06 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (142.9 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -5.50 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 447.5 attack_power points (31.82 DPS) | yes | Vanquisher's Sword (10823, -0.67 DPS, sim-verified) [quest]; Shoni's Disarming Tool (9608, -15.84 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (142.9 DPS) | yes | Swiftwind (13038, -0.08 DPS) [world_drop]; Monolithic Bow (9426, -0.08 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.00 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Ring of the Underwood; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 626, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 50 (dwarf, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 189.6. Weights run: 4.1s. Verify run: 7.0s. 782 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.761 ± 0.011, strength=1.210 ± 0.002, crit=0.978 ± 0.019 per rating point (14 rating = 1%, 13.697 per %), hit=1.292 ± 0.082 per rating point (10 rating = 1%, 12.925 per %), melee_haste=8.958 ± 1.260

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 41.9 attack_power points (3.28 DPS) | yes | Helm of Fire (8348, -0.94 DPS) [crafted]; Scorpashi Skullcap (14658, -1.07 DPS) [world_drop]; White Bandit Mask (10008, -5.71 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.9 attack_power points (2.27 DPS) | yes | Sentinel's Medallion (19540, -0.75 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 36.6 attack_power points (2.87 DPS) | yes | Skulker's Leather Shoulder (252535, -0.50 DPS) [crafted]; Failed Flying Experiment (9647, -0.54 DPS) [quest]; Sunburn Spaulders (274751, -4.07 DPS, sim-verified) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 26.4 attack_power points (2.07 DPS) | yes | Pridelord Cape (14673, -0.54 DPS) [world_drop]; Serpentskin Cloak (8259, -0.55 DPS) [world_drop]; Blackflame Cape (13109, -0.79 DPS, sim-verified) [world_drop] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | sim-verified (189.6 DPS) | yes | Icemail Jerkin (1981, +0.00 DPS) [world_drop]; Blazewind Breastplate (11193, +0.00 DPS) [quest]; Warbear Harness (15064, -4.75 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.3 attack_power points (2.45 DPS) | yes | Bracers of the Stone Princess (17714, -0.26 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.54 DPS) [crafted]; Arena Bands (18711, -0.71 DPS, sim-verified) [world] |
| hands | Sergeant Major's Chain Gauntlets (220829) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (189.6 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Raider Gloves (272100, -0.74 DPS, sim-verified) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 39.4 attack_power points (3.09 DPS) | yes | Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.11 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.45 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | sim-verified (189.6 DPS) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Basilisk Hide Pants (1718, -0.11 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -4.61 DPS, sim-verified) [quest] |
| feet | Sergeant Major's Chain Sabatons (220837) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (189.6 DPS) | yes | Albino Crocscale Boots (17728, +0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, -1.86 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 27.9 attack_power points (2.19 DPS) | yes | Ring of the Underwood (2951, -0.62 DPS) [world_drop]; Mark of Kern (2262, -0.62 DPS) [dungeon]; Blackstone Ring (17713, -0.62 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.7 attack_power points (1.93 DPS) | yes | Ring of the Underwood (2951, -0.36 DPS) [world_drop]; Mark of Kern (2262, -0.36 DPS) [dungeon]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (189.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (189.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (189.6 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thrash Blade (17705, +0.00 DPS, sim-verified) [quest]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 523.7 attack_power points (41.07 DPS) | yes | Julie's Dagger (6660, -1.08 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.26 DPS) [dungeon]; Shoni's Disarming Tool (9608, -23.45 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (189.6 DPS) | yes | Guttbuster (13139, -0.54 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.55 DPS) [world_drop]; Dark Iron Rifle (16004, -2.60 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; hands: Sergeant Major's Chain Gauntlets; waist: Substandard Belt Chain; legs: Knight's Chain Legplates; feet: Sergeant Major's Chain Sabatons; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Cestus; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 782, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60 (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 288.6. Weights run: 4.0s. Verify run: 12.3s. 1415 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.966 ± 0.019, strength=1.210 ± 0.002, crit=1.809 ± 0.037 per rating point (14 rating = 1%, 25.328 per %), hit=2.634 ± 0.153 per rating point (10 rating = 1%, 26.341 per %), melee_haste=15.325 ± 2.330

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (288.6 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -5.22 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 50.3 attack_power points (3.50 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -0.07 DPS) [quest]; Skibi's Pendant (13089, -1.30 DPS) [world_drop] |
| shoulder | Beastmaster's Mantle (226884) | Anthion's Parting Words [quest] | sim-verified (288.6 DPS) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Beaststalker's Mantle (16679, -5.26 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 54.3 attack_power points (3.78 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -0.73 DPS) [rep]; Shifting Cloak (18511, -1.46 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.4 attack_power points (8.66 DPS) | yes | Field Marshal's Chain Armor (231563, -0.55 DPS) [vendor]; Obsidian Mail Tunic (22191, -1.61 DPS) [crafted]; Knight-Captain's Chain Armor (227089, -1.62 DPS) [vendor] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (288.6 DPS) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -5.91 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-verified (288.6 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Beaststalker's Gloves (16676, -3.79 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (288.6 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -3.81 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 145.8 attack_power points (10.15 DPS) | yes | Marshal's Chain Legplates (231558, -2.04 DPS) [vendor]; Sentinel's Leather Pants (237818, -2.93 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -3.10 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (288.6 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Knight-Lieutenant's Chain Sabatons (227088, -5.05 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 73.5 attack_power points (5.12 DPS) | yes | Cutthroat's Signet (272408, -1.37 DPS) [vendor]; Protector's Band (19514, -2.60 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 65.7 attack_power points (4.57 DPS) | yes | Cutthroat's Signet (272408, -0.82 DPS) [vendor]; Protector's Band (19514, -2.05 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (288.6 DPS) | yes | Frozen Heart of the Mountain (249469, -8.29 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (288.6 DPS) | yes | Talisman of Ascendance (22678, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (288.6 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -7.46 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 836.7 attack_power points (58.24 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -4.20 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -24.68 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (288.6 DPS) | yes | The Purifier (22656, -0.15 DPS) [quest]; Bloodseeker (19107, -0.28 DPS) [quest]; Dark Iron Rifle (16004, -1.99 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Beastmaster's Mantle; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy

No-known-source sample (15 of 1415, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 765.8. Weights run: 3.4s. Verify run: 8.6s. 1415 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.200 ± 0.021, strength=1.210 ± 0.002, crit=2.524 ± 0.044 per rating point (14 rating = 1%, 35.330 per %), hit=4.085 ± 0.325 per rating point (10 rating = 1%, 40.846 per %), melee_haste=17.800 ± 4.170

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 99.2 attack_power points (13.96 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Beastmaster's Cap (226887, -4.66 DPS, sim-verified) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 64.8 attack_power points (9.12 DPS) | yes | Medallion of the Dawn (22659, -0.78 DPS) [quest]; Amulet of the Darkmoon (19491, -1.54 DPS) [quest]; Skibi's Pendant (13089, -4.25 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.7 attack_power points (13.33 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 68.8 attack_power points (9.69 DPS) | yes | Cloak of the Honor Guard (20073, -3.35 DPS) [rep]; Stalwart Cloak (272415, -3.94 DPS) [vendor]; Cape of the Black Baron (13340, -6.91 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.6 attack_power points (22.17 DPS) | yes | Field Marshal's Chain Armor (231563, -3.81 DPS) [vendor]; Timbermaw Tunic (252484, -6.01 DPS) [crafted]; Knight-Captain's Chain Armor (227089, -10.44 DPS, sim-verified) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 55.1 attack_power points (7.75 DPS) | yes | Bracers of Subterfuge (22668, -1.41 DPS) [quest]; Marshal's Chain Bracers (16461, -1.56 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.64 DPS) [pvp] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 76.2 attack_power points (10.72 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, -0.49 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 90.8 attack_power points (12.77 DPS) | yes | Highlander's Leather Girdle (20045, -3.02 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.08 DPS) [vendor]; Highlander's Chain Girdle (20043, -6.60 DPS, sim-verified) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 188.5 attack_power points (26.52 DPS) | yes | Marshal's Chain Legplates (231558, -8.16 DPS) [vendor]; Sentinel's Leather Pants (237818, -8.22 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -9.14 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 65.1 attack_power points (9.15 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -4.62 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 93.7 attack_power points (13.18 DPS) | yes | Cutthroat's Signet (272408, -3.10 DPS) [vendor]; Channeler's Ring (272406, -7.43 DPS) [vendor]; Blessed Band of Light (272407, -7.43 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 90.2 attack_power points (12.69 DPS) | yes | Cutthroat's Signet (272408, -3.95 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -6.94 DPS) [vendor]; Blessed Band of Light (272407, -6.94 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (765.8 DPS) | yes | Frozen Heart of the Mountain (249469, -23.47 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (765.8 DPS) | yes | Talisman of Ascendance (22678, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (765.8 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -49.39 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.9 attack_power points (118.02 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -14.43 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -50.18 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (765.8 DPS) | yes | Precisely Calibrated Boomstick (2100, -0.64 DPS) [world_drop]; Bloodseeker (19107, -1.44 DPS) [quest]; Dark Iron Rifle (16004, -7.20 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1415, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 73.6. Weights run: 3.7s. Verify run: 4.8s. 206 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.068 ± 0.016, strength=1.210 ± 0.002, crit=0.779 ± 0.017 per rating point (14 rating = 1%, 10.906 per %), hit=1.140 ± 0.041 per rating point (10 rating = 1%, 11.405 per %), melee_haste=7.736 ± 1.342

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.5 attack_power points (0.72 DPS) | yes | Defender's Leather Hood (252447, -0.30 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.4 attack_power points (0.54 DPS) | yes | Erudite's Amulet (277204, -0.35 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.3 attack_power points (0.45 DPS) | yes | Slime-encrusted Pads (6461, -0.43 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.4 attack_power points (0.54 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.17 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (73.6 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.93 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.3 attack_power points (0.45 DPS) | yes | Bristlebark Bindings (14569, -0.07 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Ratchet Wristwraps (274742, -0.18 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 17.2 attack_power points (0.75 DPS) | yes | Bristlebark Gloves (14572, -0.18 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.18 DPS) [crafted]; Serpent Gloves (5970, -0.21 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Brawler's Leather Belt (252428, -0.21 DPS) [crafted]; Murloc Scale Belt (5780, -0.41 DPS) [crafted]; Deviate Scale Belt (6468, -2.58 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.07 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (73.6 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.99 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.1 attack_power points (0.57 DPS) | yes | Demon Band (12054, -0.36 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; Loop of Sacrifice (281673, -0.41 DPS) [quest] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 6.2 attack_power points (0.27 DPS) | yes | Demon Band (12054, -0.06 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Loop of Sacrifice (281673, -0.11 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (10.87 DPS) | yes | Crescent Staff (6505, +0.00 DPS) [quest]; Living Root (6631, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (10.34 DPS) | yes | Wingblade (6504, -6.26 DPS, sim-verified) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.3 attack_power points (0.36 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Privateer Musket (5309, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Bounty Hunter's Ring; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 206, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 209611 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (troll, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 99.1. Weights run: 4.0s. Verify run: 2.8s. 380 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.833 ± 0.014, strength=1.210 ± 0.002, crit=0.882 ± 0.020 per rating point (14 rating = 1%, 12.355 per %), hit=1.229 ± 0.054 per rating point (10 rating = 1%, 12.295 per %), melee_haste=not significant (4.788 ± 1.299)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.3 attack_power points (0.94 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.19 DPS) [crafted]; Humbert's Helm (4724, -0.28 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 14.7 attack_power points (0.75 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; River Pride Choker (13087, -0.50 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.2 attack_power points (1.34 DPS) | yes | Mantle of Thieves (2264, -0.23 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.56 DPS) [crafted]; Bristlebark Amice (14573, -0.59 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.5 attack_power points (0.84 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Fenrus' Hide (6340, -0.28 DPS) [dungeon]; Swiftrunner Cape (6745, -0.28 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 25.7 attack_power points (1.31 DPS) | yes | Brawler's Leather Tunic (252508, -0.19 DPS) [crafted]; Panther Armor (6670, -0.34 DPS) [quest]; Brawler's Leather Armor (252490, -0.35 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.8 attack_power points (0.81 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.25 DPS) [world_drop]; Cultist's Armguards (270032, -0.30 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.5 attack_power points (0.99 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Cartilage Shapers (273054, -0.15 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.23 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.01 DPS) [crafted]; Prowler's Leather Belt (252459, -0.11 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.33 DPS) | yes | Troll's Bane Leggings (13114, -0.02 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.05 DPS) [crafted]; Petrolspill Leggings (9509, -1.24 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.3 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.3 attack_power points (0.93 DPS) | yes | Band of the Fist (17694, -0.31 DPS) [quest]; Ring of Precision (1491, -0.37 DPS) [dungeon]; Monkey Ring (6748, -0.46 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 15.2 attack_power points (0.78 DPS) | yes | Band of the Fist (17694, -0.15 DPS) [quest]; Ring of Precision (1491, -0.21 DPS) [dungeon]; Monkey Ring (6748, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.44 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (99.1 DPS) | yes | Swinetusk Shank (6691, -17.21 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 11.0 attack_power points (0.56 DPS) | yes | Double-barreled Shotgun (2098, -0.10 DPS) [world_drop]; Booty Bay Bruiser's Buckshot (274748, -0.10 DPS) [vendor]; Gravespike Repeater (273033, -0.15 DPS) [world] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 142.8. Weights run: 4.1s. Verify run: 3.4s. 590 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.849 ± 0.012, strength=1.210 ± 0.002, crit=0.905 ± 0.019 per rating point (14 rating = 1%, 12.664 per %), hit=1.370 ± 0.071 per rating point (10 rating = 1%, 13.697 per %), melee_haste=6.765 ± 1.061

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.6 attack_power points (2.39 DPS) | yes | Barbaric Iron Helm (7915, -0.43 DPS) [crafted]; Hawkeye's Helm (14591, -0.60 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.68 DPS) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 20.3 attack_power points (1.45 DPS) | yes | Ethereal Talisman (4430, -0.49 DPS) [quest]; Kaleidoscope Chain (13084, -0.58 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.59 DPS) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.3 attack_power points (2.30 DPS) | yes | Forest Tracker Epaulets (2278, -0.43 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.55 DPS) [dungeon]; Nightscape Shoulders (8192, -0.85 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.6 attack_power points (1.18 DPS) | yes | Imperial Cloak (6432, -0.13 DPS) [dungeon]; Parachute Cloak (10518, -0.13 DPS) [crafted]; Tigerstrike Mantle (13108, -0.13 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 44.4 attack_power points (3.15 DPS) | yes | Wolffear Harness (13110, -0.94 DPS, sim-verified) [world_drop]; Kolkar Marauder Chain (6773, -1.03 DPS) [quest]; Tough Scorpid Breastplate (8203, -1.18 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.42 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.29 DPS) [world_drop] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 33.0 attack_power points (2.35 DPS) | yes | Gloves of Holy Might (867, -0.02 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.07 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.26 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.13 DPS) | yes | Ogron's Sash (13117, -0.18 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.43 DPS) [rep]; Skulker's Leather Belt (252520, -0.43 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 38.8 attack_power points (2.76 DPS) | yes | Triprunner Dungarees (9624, -0.14 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.89 DPS) [crafted]; Ferine Leggings (6690, -0.91 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.4 attack_power points (2.09 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.22 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 24.5 attack_power points (1.74 DPS) | yes | Falcon's Hook (7552, -0.30 DPS) [dungeon]; Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 20.9 attack_power points (1.49 DPS) | yes | Falcon's Hook (7552, -0.05 DPS) [dungeon]; Mark of Kern (2262, -0.06 DPS) [dungeon]; Assault Band (13095, -0.06 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (142.8 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.49 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 447.5 attack_power points (31.82 DPS) | yes | Vanquisher's Sword (10823, -0.43 DPS, sim-verified) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (142.8 DPS) | yes | Swiftwind (13038, -0.08 DPS) [world_drop]; Monolithic Bow (9426, -0.08 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.00 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Legionnaire's Band; finger2: Ring of the Underwood; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 191.1. Weights run: 4.1s. Verify run: 3.7s. 739 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.761 ± 0.011, strength=1.210 ± 0.002, crit=0.978 ± 0.019 per rating point (14 rating = 1%, 13.697 per %), hit=1.292 ± 0.082 per rating point (10 rating = 1%, 12.925 per %), melee_haste=8.958 ± 1.260

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 32.7 attack_power points (2.56 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.09 DPS) [quest]; Helm of Fire (8348, -0.22 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.9 attack_power points (2.27 DPS) | yes | Scout's Medallion (19535, -0.61 DPS) [rep]; Woven Ivy Necklace (19159, -0.71 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 31.4 attack_power points (2.46 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.09 DPS) [crafted]; Failed Flying Experiment (9647, -0.13 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 26.4 attack_power points (2.07 DPS) | yes | Pridelord Cape (14673, -0.54 DPS) [world_drop]; Serpentskin Cloak (8259, -0.55 DPS) [world_drop]; Blackflame Cape (13109, -0.99 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 45.0 attack_power points (3.53 DPS) | yes | Blazewind Breastplate (11193, -0.07 DPS) [quest]; Icemail Jerkin (1981, -0.22 DPS) [world_drop]; Stone Guard's Chain Armor (220827, -0.38 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.3 attack_power points (2.45 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Arena Bands (18711, -0.66 DPS, sim-verified) [world] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 50.8 attack_power points (3.99 DPS) | yes | Gloves of Holy Might (867, -1.34 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -1.42 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.99 DPS, sim-verified) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 39.4 attack_power points (3.09 DPS) | yes | Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.11 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.45 DPS) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 43.9 attack_power points (3.44 DPS) | yes | Stone Guard's Chain Legplates (220833, -0.44 DPS) [vendor]; Basilisk Hide Pants (1718, -0.54 DPS) [world_drop]; Triprunner Dungarees (9624, -0.67 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 36.2 attack_power points (2.84 DPS) | yes | Albino Crocscale Boots (17728, -0.08 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.09 DPS) [crafted]; Fleetfoot Greaves (11627, -0.22 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 27.9 attack_power points (2.19 DPS) | yes | White Bone Band (11862, -0.31 DPS) [quest]; Ring of the Underwood (2951, -0.62 DPS) [world_drop]; Blackstone Ring (17713, -0.62 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.7 attack_power points (1.93 DPS) | yes | White Bone Band (11862, -0.05 DPS) [quest]; Ring of the Underwood (2951, -0.36 DPS) [world_drop]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (191.1 DPS) | yes | Molten Heart of the Mountain (249470, -4.81 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (191.1 DPS) | yes | Molten Heart of the Mountain (249470, -1.71 DPS, sim-verified) [crafted] |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-verified (191.1 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Thorium Cestus (250614, -2.35 DPS, sim-verified) [crafted] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 523.7 attack_power points (41.07 DPS) | yes | Julie's Dagger (6660, +0.00 DPS) [world_drop]; Claw of Celebras (17738, -3.26 DPS) [dungeon]; White Bone Shredder (11863, -5.61 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (191.1 DPS) | yes | Guttbuster (13139, -0.54 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.55 DPS) [world_drop]; Dark Iron Rifle (16004, -2.51 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: Legionnaire's Band; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thrash Blade; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 739, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 293.9. Weights run: 4.0s. Verify run: 12.5s. 1395 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.966 ± 0.019, strength=1.210 ± 0.002, crit=1.809 ± 0.037 per rating point (14 rating = 1%, 25.328 per %), hit=2.634 ± 0.153 per rating point (10 rating = 1%, 26.341 per %), melee_haste=15.325 ± 2.330

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (293.9 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -6.84 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 50.3 attack_power points (3.50 DPS) | yes | Amulet of the Darkmoon (19491, -0.06 DPS) [quest]; Medallion of the Dawn (22659, -0.07 DPS) [quest]; Skibi's Pendant (13089, -1.30 DPS) [world_drop] |
| shoulder | Beastmaster's Mantle (226884) | Anthion's Parting Words [quest] | sim-verified (293.9 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Beaststalker's Mantle (16679, -5.49 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 54.3 attack_power points (3.78 DPS) | yes | Cape of the Black Baron (13340, -0.34 DPS) [dungeon]; Deathguard's Cloak (20068, -0.73 DPS) [rep]; Shifting Cloak (18511, -1.46 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.4 attack_power points (8.66 DPS) | yes | Warlord's Chain Armor (231566, -0.55 DPS) [vendor]; Obsidian Mail Tunic (22191, -1.30 DPS, sim-verified) [crafted]; Legionnaire's Chain Armor (227083, -1.62 DPS) [vendor] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (293.9 DPS) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -6.57 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-verified (293.9 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Beaststalker's Gloves (16676, -3.57 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (293.9 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -3.78 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 145.8 attack_power points (10.15 DPS) | yes | Outrider's Chain Leggings (22673, -1.78 DPS, sim-verified) [rep]; General's Chain Legplates (231567, -2.04 DPS) [vendor]; Sentinel's Leather Pants (237818, -2.93 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (293.9 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -5.60 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 73.5 attack_power points (5.12 DPS) | yes | Cutthroat's Signet (272408, -1.37 DPS) [vendor]; Legionnaire's Band (19510, -2.60 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 65.7 attack_power points (4.57 DPS) | yes | Cutthroat's Signet (272408, -0.82 DPS) [vendor]; Legionnaire's Band (19510, -2.05 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (293.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (293.9 DPS) | yes | Frozen Heart of the Mountain (249469, -2.56 DPS) [crafted]; Molten Heart of the Mountain (249470, -4.71 DPS, sim-verified) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (293.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -7.48 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 836.7 attack_power points (58.24 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -5.02 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -24.68 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (293.9 DPS) | yes | The Purifier (22656, -0.15 DPS) [quest]; Bloodseeker (19107, -0.28 DPS) [quest]; Dark Iron Rifle (16004, -2.50 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Beastmaster's Mantle; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy

No-known-source sample (15 of 1395, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 776.4. Weights run: 3.4s. Verify run: 8.7s. 1395 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.200 ± 0.021, strength=1.210 ± 0.002, crit=2.524 ± 0.044 per rating point (14 rating = 1%, 35.330 per %), hit=4.085 ± 0.325 per rating point (10 rating = 1%, 40.846 per %), melee_haste=17.800 ± 4.170

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 99.2 attack_power points (13.96 DPS) | yes | Beastmaster's Cap (226887, +0.00 DPS) [quest]; Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 64.8 attack_power points (9.12 DPS) | yes | Medallion of the Dawn (22659, -0.78 DPS) [quest]; Amulet of the Darkmoon (19491, -1.54 DPS) [quest]; Skibi's Pendant (13089, -4.25 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.7 attack_power points (13.33 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 68.8 attack_power points (9.69 DPS) | yes | Cape of the Black Baron (13340, -2.23 DPS) [dungeon]; Deathguard's Cloak (20068, -3.35 DPS) [rep]; Stalwart Cloak (272415, -3.94 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.6 attack_power points (22.17 DPS) | yes | Warlord's Chain Armor (231566, -3.81 DPS) [vendor]; Legionnaire's Chain Armor (227083, -4.79 DPS) [vendor]; Timbermaw Tunic (252484, -5.15 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 55.1 attack_power points (7.75 DPS) | yes | Bracers of Subterfuge (22668, -1.41 DPS) [quest]; General's Chain Wristguards (16570, -1.56 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.64 DPS) [pvp] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 76.2 attack_power points (10.72 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 90.8 attack_power points (12.77 DPS) | yes | Defiler's Chain Girdle (20150, -3.02 DPS) [rep]; Defiler's Leather Girdle (20190, -3.02 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.08 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 188.5 attack_power points (26.52 DPS) | yes | General's Chain Legplates (231567, -8.16 DPS) [vendor]; Sentinel's Leather Pants (237818, -8.22 DPS) [vendor]; Outrider's Chain Leggings (22673, -8.52 DPS, sim-verified) [rep] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 65.1 attack_power points (9.15 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -0.93 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 93.7 attack_power points (13.18 DPS) | yes | Cutthroat's Signet (272408, -3.10 DPS) [vendor]; Channeler's Ring (272406, -7.43 DPS) [vendor]; Blessed Band of Light (272407, -7.43 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 90.2 attack_power points (12.69 DPS) | yes | Cutthroat's Signet (272408, -5.98 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -6.94 DPS) [vendor]; Blessed Band of Light (272407, -6.94 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (776.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (776.4 DPS) | yes | Frozen Heart of the Mountain (249469, -4.76 DPS) [crafted]; Molten Heart of the Mountain (249470, -11.62 DPS, sim-verified) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (776.4 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -45.90 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.9 attack_power points (118.02 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -12.33 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -50.18 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (776.4 DPS) | yes | Precisely Calibrated Boomstick (2100, -0.64 DPS) [world_drop]; Bloodseeker (19107, -1.44 DPS) [quest]; Dark Iron Rifle (16004, -8.52 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1395, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

