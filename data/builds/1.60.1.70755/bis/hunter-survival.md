# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 46.7. Weights run: 8.4s. Verify run: 24.2s. 258 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.527 ± 0.013, strength=1.210 ± 0.002, crit=0.985 ± 0.015 per rating point (14 rating = 1%, 13.787 per %), hit=1.335 ± 0.037 per rating point (10 rating = 1%, 13.352 per %), melee_haste=5.440 ± 0.834, intellect=0.131 ± 0.014

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 20.2 attack_power points (0.91 DPS) | yes | Defender's Leather Hood (252447, -0.44 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -0.86 DPS) [crafted]; Shadow Goggles (4373, -0.88 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 15.2 attack_power points (0.68 DPS) | yes | Erudite's Amulet (277204, -0.44 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.85 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 12.6 attack_power points (0.57 DPS) | yes | Magician's Mantle (12998, -0.52 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.54 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (46.7 DPS) | yes | Cape of the Brotherhood (5193, -0.11 DPS) [dungeon]; Sentry Cloak (2059, -0.23 DPS) [world_drop]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 27.8 attack_power points (1.25 DPS) | yes | Brawler's Leather Armor (252490, -0.19 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.44 DPS) [crafted]; Prospector's Chestpiece (14562, -0.46 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 12.6 attack_power points (0.57 DPS) | yes | Bravo's Armbands (270015, -0.06 DPS) [quest]; Wolf Bracers (4794, -0.11 DPS) [vendor]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (46.7 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.69 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.81 DPS) | yes | Brawler's Leather Belt (252428, -0.14 DPS) [crafted]; Trapper's Leather Belt (252429, -0.34 DPS) [crafted]; Deviate Scale Belt (6468, -2.42 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (46.7 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, -1.86 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (46.7 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.85 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 14.9 attack_power points (0.67 DPS) | yes | Demon Band (12054, -0.45 DPS) [world_drop]; The 1 Ring (8350, -0.50 DPS) [world]; Ring of the Moon (12052, -0.62 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.8 attack_power points (0.26 DPS) | yes | Demon Band (12054, -0.04 DPS) [world_drop]; The 1 Ring (8350, -0.09 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.6 attack_power points (10.69 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Duskbringer (2205, +0.00 DPS) [dungeon]; Impaling Harpoon (5200, +0.00 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 231.1 attack_power points (10.40 DPS) | yes | Cruel Barb (5191, -0.33 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 10.1 attack_power points (0.45 DPS) | yes | Deadly Blunderbuss (4369, -0.22 DPS, sim-verified) [crafted]; Light Bow (4576, -0.23 DPS) [world_drop]; Owlsight Rifle (15205, -0.23 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Lavishly Jeweled Ring; main_hand: Butcher's Cleaver; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230231040000100)

Set DPS (verified): 91.2. Weights run: 8.9s. Verify run: 18.1s. 447 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.265 ± 0.013, strength=1.210 ± 0.001, crit=1.278 ± 0.018 per rating point (14 rating = 1%, 17.896 per %), hit=1.717 ± 0.050 per rating point (10 rating = 1%, 17.167 per %), melee_haste=10.689 ± 0.999, intellect=0.704 ± 0.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.7 attack_power points (1.51 DPS) | yes | Brawler's Leather Hood (252504, -0.29 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.45 DPS) [world]; Defender's Leather Helm (252455, -0.54 DPS) [crafted] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 18.1 attack_power points (1.21 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.13 DPS) [world]; Kaleidoscope Chain (13084, -0.56 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.89 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.0 attack_power points (2.07 DPS) | yes | Barbaric Shoulders (5964, -0.91 DPS) [crafted]; Bristlebark Amice (14573, -0.92 DPS) [world_drop]; Mantle of Thieves (2264, -0.95 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (91.2 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.19 DPS) [pvp]; Glowing Lizardscale Cloak (6449, -0.39 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 31.7 attack_power points (2.12 DPS) | yes | Garb of Florid Feathers (273026, -0.35 DPS, sim-verified) [dungeon]; Brawler's Leather Tunic (252508, -0.42 DPS) [crafted]; Tunic of Westfall (2041, -0.45 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.4 attack_power points (1.23 DPS) | yes | Barbaric Bracers (18948, -0.29 DPS, sim-verified) [crafted]; Insignia Bracers (6410, -0.32 DPS) [world_drop]; Demonhide Bracers (270033, -0.38 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.1 attack_power points (1.48 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Cartilage Shapers (273054, -0.11 DPS) [world]; Wolfclaw Gloves (1978, -0.16 DPS) [dungeon] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 27.6 attack_power points (1.85 DPS) | yes | Stalker's Leather Belt (252521, -0.20 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.24 DPS) [rep]; Prowler's Leather Belt (252459, -0.29 DPS, sim-verified) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 34.5 attack_power points (2.30 DPS) | yes | Troll's Bane Leggings (13114, -0.18 DPS) [world_drop]; Trapper's Leather Pants (252501, -0.27 DPS) [crafted]; Petrolspill Leggings (9509, -1.50 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.8 attack_power points (1.45 DPS) | yes | Insignia Boots (4055, -0.24 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.29 DPS) [crafted]; Trapper's Leather Boots (252440, -0.51 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 20.9 attack_power points (1.39 DPS) | yes | Ring of Precision (1491, -0.49 DPS) [dungeon]; Seal of Wrynn (2933, -0.51 DPS) [quest]; Monkey Ring (6748, -0.64 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.5 attack_power points (1.10 DPS) | yes | Seal of Wrynn (2933, -0.22 DPS) [quest]; Monkey Ring (6748, -0.34 DPS) [quest]; Ring of Precision (1491, -0.56 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-verified (91.2 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Toxic Revenger (9453, -4.18 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (21.28 DPS) | yes | Zealot Blade (13033, -1.44 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -5.76 DPS) [quest] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 13.6 attack_power points (0.91 DPS) | yes | Double-barreled Shotgun (2098, -0.21 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.30 DPS) [world_drop]; Crystalpine Stinger (13037, -0.47 DPS, sim-verified) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Skulker's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 447, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 157.8. Weights run: 8.9s. Verify run: 26.9s. 681 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.330 ± 0.014, strength=1.210 ± 0.001, crit=1.815 ± 0.024 per rating point (14 rating = 1%, 25.412 per %), hit=2.410 ± 0.093 per rating point (10 rating = 1%, 24.103 per %), melee_haste=7.796 ± 1.237, intellect=0.038 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 41.1 attack_power points (2.90 DPS) | yes | White Bandit Mask (10008, -0.16 DPS) [crafted]; Barbaric Iron Helm (7915, -0.66 DPS) [crafted]; Hawkeye's Helm (14591, -0.75 DPS) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 25.6 attack_power points (1.81 DPS) | yes | Kaleidoscope Chain (13084, -0.80 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.95 DPS) [world]; River Pride Choker (13087, -1.47 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.7 attack_power points (2.24 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.43 DPS) [crafted]; Mantle of Thieves (2264, -0.59 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (157.8 DPS) | yes | Imperial Cloak (6432, -0.20 DPS) [dungeon]; Parachute Cloak (10518, -0.20 DPS) [crafted]; Hawkeye's Cloak (14593, -0.80 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 80.0 attack_power points (5.65 DPS) | yes | Tough Scorpid Breastplate (8203, -3.16 DPS) [crafted]; Nightscape Tunic (8175, -3.18 DPS) [crafted]; Wolffear Harness (13110, -3.72 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Hawkeye's Bracers (14590, -0.08 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.10 DPS) [dungeon]; Dusky Bracers (7378, -0.10 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (157.8 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.09 DPS) [crafted]; Dragonscale Gauntlets (8347, -3.38 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 33.4 attack_power points (2.36 DPS) | yes | Ogron's Sash (13117, -0.11 DPS) [world_drop]; Highlander's Leather Girdle (20116, -0.24 DPS) [rep]; Skulker's Leather Belt (252520, -0.37 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 48.9 attack_power points (3.45 DPS) | yes | Triprunner Dungarees (9624, -0.24 DPS) [quest]; Petrolspill Leggings (9509, -1.15 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.15 DPS) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 34.2 attack_power points (2.41 DPS) | yes | Skulker's Leather Shoes (252531, -0.01 DPS) [crafted]; Imperial Leather Boots (6431, -0.18 DPS) [dungeon]; Worn Running Boots (9398, -0.28 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 28.3 attack_power points (2.00 DPS) | yes | Falcon's Hook (7552, -0.26 DPS) [dungeon]; Mark of Kern (2262, -0.59 DPS) [dungeon]; Assault Band (13095, -0.59 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.7 attack_power points (1.82 DPS) | yes | Falcon's Hook (7552, -0.08 DPS) [dungeon]; Mark of Kern (2262, -0.40 DPS) [dungeon]; Assault Band (13095, -0.40 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (157.8 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Black Menace (6831, -6.88 DPS, sim-verified) [quest] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (31.20 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Shoni's Disarming Tool (9608, -14.82 DPS) [quest] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-verified (157.8 DPS) | yes | Monolithic Bow (9426, -0.15 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.15 DPS) [quest]; The Silencer (13138, -0.16 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 681, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band

### Band 50 (dwarf, 5000000000000000-0010000000000000-500230231051222151)

Set DPS (verified): 245.8. Weights run: 9.1s. Verify run: 35.0s. 854 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.394 ± 0.015, strength=1.210 ± 0.002, crit=1.968 ± 0.026 per rating point (14 rating = 1%, 27.551 per %), hit=2.658 ± 0.132 per rating point (10 rating = 1%, 26.579 per %), melee_haste=9.991 ± 1.699, intellect=0.040 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 43.3 attack_power points (4.01 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Knight-Lieutenant's Mail Helmet (223075, -0.20 DPS) [vendor]; Helm of Fire (8348, -0.24 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (3.44 DPS) | yes | Sentinel's Medallion (19540, -1.00 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 37.2 attack_power points (3.45 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.11 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackflame Cape (13109, -0.67 DPS) [world_drop]; Serpentskin Cloak (8259, -0.89 DPS) [world_drop]; Pridelord Cape (14673, -0.99 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 84.0 attack_power points (7.78 DPS) | yes | Knight's Chain Armor (220828, -1.91 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -2.35 DPS) [quest]; Warbear Harness (15064, -2.56 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 40.8 attack_power points (3.77 DPS) | yes | Skulker's Leather Bracers (252540, -0.99 DPS) [crafted]; Pridelord Bands (14672, -1.00 DPS) [world_drop]; Bracers of the Stone Princess (17714, -1.16 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -4.78 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 48.0 attack_power points (4.45 DPS) | yes | Highlander's Chain Girdle (20088, -0.05 DPS) [rep]; Highlander's Leather Girdle (20115, -0.05 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.22 DPS) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 61.5 attack_power points (5.69 DPS) | yes | Gryphon Rider's Leggings (9652, -0.69 DPS) [quest]; Serpentskin Leggings (8262, -0.69 DPS) [world_drop]; Stormshroud Pants (15057, -1.25 DPS, sim-verified) [crafted] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 47.9 attack_power points (4.43 DPS) | yes | Fleetfoot Greaves (11627, -0.20 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.32 DPS) [crafted]; Elven Chain Boots (13125, -0.44 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.6 attack_power points (4.31 DPS) | yes | Masons Fraternity Ring (9533, -1.21 DPS) [quest]; Ring of the Underwood (2951, -1.87 DPS) [world_drop]; Falcon's Hook (7552, -1.98 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 33.6 attack_power points (3.12 DPS) | yes | Masons Fraternity Ring (9533, -0.01 DPS) [quest]; Ring of the Underwood (2951, -0.67 DPS) [world_drop]; Falcon's Hook (7552, -0.78 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -0.91 DPS, sim-verified) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.51 DPS, sim-verified) [crafted] |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Stoneraven (13059, +0.00 DPS) [world_drop]; Dawn's Edge (12774, -3.46 DPS, sim-verified) [crafted] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | sim-verified (245.8 DPS) | yes | Claw of Celebras (17738, -2.71 DPS) [dungeon]; Thorium Cestus (250614, -6.29 DPS, sim-verified) [crafted]; Shoni's Disarming Tool (9608, -25.87 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -1.10 DPS, sim-verified) [world_drop]; Gryphonwing Long Bow (13022, -1.33 DPS) [dungeon]; Skull Splitting Crossbow (13039, -1.47 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Thrash Blade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 854, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60 (dwarf, 5000000000000000-0050050010000000-500230231051222151)

Set DPS (verified): 369.4. Weights run: 6.6s. Verify run: 62.6s. 1594 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.114 ± 0.016, strength=1.210 ± 0.002, crit=2.048 ± 0.031 per rating point (14 rating = 1%, 28.678 per %), hit=1.634 ± 0.134 per rating point (10 rating = 1%, 16.340 per %), melee_haste=16.916 ± 1.713, intellect=1.247 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 120.5 attack_power points (15.57 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Lieutenant Commander's Chain Greathelm (227086, -0.33 DPS) [vendor]; Outlaw's Collar (279253, -5.56 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.5 attack_power points (4.33 DPS) | yes | Sentinel's Medallion (19538, -0.24 DPS) [rep]; Sunprism Pendant (20649, -1.05 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 85.8 attack_power points (11.08 DPS) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Chain Spaulders (16468, -0.26 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.92 DPS) [rep]; Sergeant Major's Cape (16337, -1.96 DPS) [pvp]; Dark Phantom Cape (13122, -2.58 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 123.2 attack_power points (15.93 DPS) | yes | Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Field Marshal's Chain Breastplate (16466, -1.00 DPS) [vendor]; Timbermaw Tunic (252484, -3.41 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 53.5 attack_power points (6.91 DPS) | yes | Marshal's Chain Bracers (16461, -0.32 DPS) [pvp]; Windtalker's Wristguards (19582, -0.71 DPS) [rep]; Knight-Captain's Chain Armguards (16402, -1.48 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Theramore Gloves (271948, -1.63 DPS, sim-verified) [rep] |
| waist | Highlander's Chain Girdle (20043) (or Highlander's Leather Girdle (20045)) | The League of Arathor [rep] | 62.7 attack_power points (8.10 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Highlander's Leather Girdle (20045, +0.00 DPS) [rep]; Might of the Timbermaw (19044, -0.99 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 147.7 attack_power points (19.08 DPS) | yes | Marshal's Chain Legplates (231558, -1.33 DPS) [vendor]; Marshal's Chain Legguards (231577, -4.16 DPS) [pvp]; Sentinel's Leather Pants (237818, -4.30 DPS) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 54.9 attack_power points (7.10 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 59.0 attack_power points (7.63 DPS) | yes | Blackstone Ring (17713, -2.93 DPS) [dungeon]; Masons Fraternity Ring (9533, -3.80 DPS) [quest]; Band of the Penitent (13217, -3.92 DPS) [quest] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | 37.8 attack_power points (4.88 DPS) | yes | Blackstone Ring (17713, -0.19 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.06 DPS) [quest]; Band of the Penitent (13217, -1.18 DPS) [quest] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -5.51 DPS, sim-verified) [crafted] |
| off_hand | Electrified Dagger (19100) | Stormpike Guard [rep] | 616.1 attack_power points (79.62 DPS) | yes | Heartseeker (12783, +0.00 DPS) [crafted]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -17.31 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.66 DPS) [quest]; Blackcrow (12651, -0.89 DPS) [dungeon]; Eaglehorn Long Bow (13023, -1.09 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Highlander's Chain Girdle; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Electrified Dagger

No-known-source sample (15 of 1594, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60, raid preset (dwarf, 5000000000000000-0050050010000000-500230231051222151)

Set DPS (verified): 862.3. Weights run: 5.7s. Verify run: 53.6s. 1594 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.368 ± 0.020, strength=1.210 ± 0.002, crit=2.765 ± 0.040 per rating point (14 rating = 1%, 38.711 per %), hit=4.129 ± 0.293 per rating point (10 rating = 1%, 41.288 per %), melee_haste=20.858 ± 3.180, intellect=1.238 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 135.5 attack_power points (25.43 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -11.56 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.8 attack_power points (6.91 DPS) | yes | Sentinel's Medallion (19538, -0.25 DPS) [rep]; Sunprism Pendant (20649, -1.58 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 102.6 attack_power points (19.27 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -1.82 DPS) [rep]; Cloak of Flames (3475, -2.67 DPS) [world_drop]; Fel Cape (279269, -3.16 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 167.6 attack_power points (31.46 DPS) | yes | Field Marshal's Chain Armor (231563, -3.01 DPS) [vendor]; Field Marshal's Chain Breastplate (16466, -6.30 DPS) [vendor]; Savage Gladiator Chain (11726, -8.54 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 58.3 attack_power points (10.94 DPS) | yes | Marshal's Chain Bracers (16461, -0.43 DPS) [pvp]; Windtalker's Wristguards (19582, -1.95 DPS) [rep]; Knight-Captain's Chain Armguards (16402, -2.25 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Mourning Handguards (279257, -1.91 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20043) (or Highlander's Leather Girdle (20045)) | The League of Arathor [rep] | 72.7 attack_power points (13.65 DPS) | yes | Highlander's Leather Girdle (20045, +0.00 DPS) [rep]; Marshal's Chain Girdle (16464, -0.33 DPS) [pvp]; Highlander's Mail Girdle (20044, -2.43 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 201.6 attack_power points (37.84 DPS) | yes | Marshal's Chain Legplates (231558, -9.38 DPS) [vendor]; Sentinel's Leather Pants (237818, -11.31 DPS) [vendor]; Marshal's Chain Legguards (231577, -12.67 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.1 attack_power points (11.66 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 94.0 attack_power points (17.64 DPS) | yes | Protector's Band (19514, -10.03 DPS) [rep]; Band of the Penitent (13217, -10.38 DPS) [quest]; Don Rodrigo's Band (21563, -10.38 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 61.3 attack_power points (11.50 DPS) | yes | Protector's Band (19514, -3.89 DPS) [rep]; Band of the Penitent (13217, -4.24 DPS) [quest]; Don Rodrigo's Band (21563, -4.24 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -5.61 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -19.77 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 626.6 attack_power points (117.62 DPS) | yes | Electrified Dagger (19100, +0.00 DPS) [rep]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -27.10 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -1.29 DPS) [quest]; Blackcrow (12651, -1.71 DPS, sim-verified) [dungeon]; Eaglehorn Long Bow (13023, -1.78 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Highlander's Chain Girdle; legs: Sentinel's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket2: Molten Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Heartseeker

No-known-source sample (15 of 1594, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 48.0. Weights run: 8.4s. Verify run: 23.5s. 254 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.527 ± 0.013, strength=1.210 ± 0.002, crit=0.985 ± 0.015 per rating point (14 rating = 1%, 13.787 per %), hit=1.335 ± 0.037 per rating point (10 rating = 1%, 13.352 per %), melee_haste=5.440 ± 0.834, intellect=0.131 ± 0.014

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 20.2 attack_power points (0.91 DPS) | yes | Defender's Leather Hood (252447, -0.46 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -0.86 DPS) [crafted]; Shadow Goggles (4373, -0.88 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 15.2 attack_power points (0.68 DPS) | yes | Erudite's Amulet (277204, -0.46 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 12.6 attack_power points (0.57 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.54 DPS) [crafted]; Magician's Mantle (12998, -0.57 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (48.0 DPS) | yes | Cape of the Brotherhood (5193, -0.11 DPS) [dungeon]; Sentry Cloak (2059, -0.23 DPS) [world_drop]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 23.7 attack_power points (1.07 DPS) | yes | Prospector's Chestpiece (14562, -0.28 DPS) [world_drop]; Defender's Leather Armor (252434, -0.35 DPS) [crafted]; Trapper's Leather Armor (252491, -0.36 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 12.6 attack_power points (0.57 DPS) | yes | Wolf Bracers (4794, -0.11 DPS) [vendor]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Spare Part Bindings (279875, -0.23 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (48.0 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.60 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.81 DPS) | yes | Brawler's Leather Belt (252428, -0.14 DPS) [crafted]; Trapper's Leather Belt (252429, -0.34 DPS) [crafted]; Deviate Scale Belt (6468, -2.36 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (48.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, -1.77 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (48.0 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.77 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 14.9 attack_power points (0.67 DPS) | yes | Lavishly Jeweled Ring (1156, -0.41 DPS) [dungeon]; Demon Band (12054, -0.45 DPS) [world_drop]; Loop of Sacrifice (281673, -0.48 DPS) [quest] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 7.6 attack_power points (0.34 DPS) | yes | Lavishly Jeweled Ring (1156, -0.08 DPS) [dungeon]; Demon Band (12054, -0.12 DPS) [world_drop]; Loop of Sacrifice (281673, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wingblade (6504) | Leaders of the Fang [quest] | sim-verified (48.0 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Butcher's Cleaver (1292, -0.45 DPS, sim-verified) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 231.1 attack_power points (10.40 DPS) | yes | Cruel Barb (5191, -0.47 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 10.1 attack_power points (0.45 DPS) | yes | Light Bow (4576, -0.23 DPS) [world_drop]; Privateer Musket (5309, -0.23 DPS) [quest]; Deadly Blunderbuss (4369, -0.23 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Bounty Hunter's Ring; main_hand: Wingblade; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 254, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade

### Band 30 (troll, 0000000000000000-0000000000000000-500230231040000100)

Set DPS (verified): 91.1. Weights run: 8.9s. Verify run: 22.0s. 446 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.265 ± 0.013, strength=1.210 ± 0.001, crit=1.278 ± 0.018 per rating point (14 rating = 1%, 17.896 per %), hit=1.717 ± 0.050 per rating point (10 rating = 1%, 17.167 per %), melee_haste=10.689 ± 0.999, intellect=0.704 ± 0.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.7 attack_power points (1.51 DPS) | yes | Brawler's Leather Hood (252504, -0.31 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.45 DPS) [world]; Defender's Leather Helm (252455, -0.54 DPS) [crafted] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 18.1 attack_power points (1.21 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.13 DPS) [world]; Kaleidoscope Chain (13084, -0.15 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.89 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.0 attack_power points (2.07 DPS) | yes | Mantle of Thieves (2264, -0.58 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.91 DPS) [crafted]; Bristlebark Amice (14573, -0.92 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.39 DPS) [dungeon]; Swiftrunner Cape (6745, -0.39 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 31.7 attack_power points (2.12 DPS) | yes | Garb of Florid Feathers (273026, -0.05 DPS) [dungeon]; Brawler's Leather Tunic (252508, -0.42 DPS) [crafted]; Green Leather Armor (4255, -0.53 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.4 attack_power points (1.23 DPS) | yes | Barbaric Bracers (18948, -0.31 DPS, sim-verified) [crafted]; Insignia Bracers (6410, -0.32 DPS) [world_drop]; Loamflake Bracers (15462, -0.48 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.1 attack_power points (1.48 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Cartilage Shapers (273054, -0.11 DPS) [world]; Wolfclaw Gloves (1978, -0.16 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Stalker's Leather Belt (252521, -0.18 DPS) [crafted]; Skulker's Leather Belt (252520, -0.19 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.22 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 34.5 attack_power points (2.30 DPS) | yes | Troll's Bane Leggings (13114, -0.18 DPS) [world_drop]; Trapper's Leather Pants (252501, -0.27 DPS) [crafted]; Petrolspill Leggings (9509, -0.78 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.8 attack_power points (1.45 DPS) | yes | Insignia Boots (4055, -0.24 DPS) [world_drop]; Warsong Boots (16977, -0.24 DPS) [quest]; Brawler's Leather Boots (252439, -0.29 DPS) [crafted] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 20.9 attack_power points (1.39 DPS) | yes | Band of the Fist (17694, -0.46 DPS) [quest]; Ring of Precision (1491, -0.49 DPS) [dungeon]; Monkey Ring (6748, -0.64 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.5 attack_power points (1.10 DPS) | yes | Band of the Fist (17694, -0.17 DPS) [quest]; Ring of Precision (1491, -0.19 DPS) [dungeon]; Monkey Ring (6748, -0.34 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Claw of the Shadowmancer (2912, -3.55 DPS, sim-verified) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (21.28 DPS) | yes | Zealot Blade (13033, -0.77 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 13.6 attack_power points (0.91 DPS) | yes | Crystalpine Stinger (13037, -0.11 DPS) [world_drop]; Double-barreled Shotgun (2098, -0.21 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.30 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 446, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 155.3. Weights run: 8.9s. Verify run: 27.1s. 671 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.330 ± 0.014, strength=1.210 ± 0.001, crit=1.815 ± 0.024 per rating point (14 rating = 1%, 25.412 per %), hit=2.410 ± 0.093 per rating point (10 rating = 1%, 24.103 per %), melee_haste=7.796 ± 1.237, intellect=0.038 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 41.1 attack_power points (2.90 DPS) | yes | White Bandit Mask (10008, -0.16 DPS) [crafted]; Barbaric Iron Helm (7915, -0.66 DPS) [crafted]; Hawkeye's Helm (14591, -0.75 DPS) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 25.6 attack_power points (1.81 DPS) | yes | Ethereal Talisman (4430, -0.72 DPS) [quest]; Kaleidoscope Chain (13084, -0.80 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.95 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.7 attack_power points (2.24 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.43 DPS) [crafted]; Mantle of Thieves (2264, -0.59 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (155.3 DPS) | yes | Imperial Cloak (6432, -0.09 DPS) [dungeon]; Parachute Cloak (10518, -0.09 DPS) [crafted]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 80.0 attack_power points (5.65 DPS) | yes | Wolffear Harness (13110, -2.05 DPS, sim-verified) [world_drop]; Tough Scorpid Breastplate (8203, -3.16 DPS) [crafted]; Nightscape Tunic (8175, -3.18 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Hawkeye's Bracers (14590, -0.08 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.10 DPS) [dungeon]; Dusky Bracers (7378, -0.10 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (155.3 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.09 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.79 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 33.4 attack_power points (2.36 DPS) | yes | Ogron's Sash (13117, -0.11 DPS) [world_drop]; Defiler's Leather Girdle (20192, -0.24 DPS) [rep]; Skulker's Leather Belt (252520, -0.37 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 48.9 attack_power points (3.45 DPS) | yes | Triprunner Dungarees (9624, -0.24 DPS) [quest]; Petrolspill Leggings (9509, -1.15 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.15 DPS) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 34.2 attack_power points (2.41 DPS) | yes | Skulker's Leather Shoes (252531, -0.01 DPS) [crafted]; Imperial Leather Boots (6431, -0.18 DPS) [dungeon]; Worn Running Boots (9398, -0.28 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 28.3 attack_power points (2.00 DPS) | yes | Falcon's Hook (7552, -0.26 DPS) [dungeon]; Mark of Kern (2262, -0.59 DPS) [dungeon]; Assault Band (13095, -0.59 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.7 attack_power points (1.82 DPS) | yes | Falcon's Hook (7552, -0.08 DPS) [dungeon]; Mark of Kern (2262, -0.40 DPS) [dungeon]; Assault Band (13095, -0.40 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (155.3 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.69 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (31.20 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-verified (155.3 DPS) | yes | Monolithic Bow (9426, -0.15 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.15 DPS) [quest]; The Silencer (13138, -0.16 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Legionnaire's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 671, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 5000000000000000-0010000000000000-500230231051222151)

Set DPS (verified): 246.4. Weights run: 9.1s. Verify run: 34.8s. 845 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.394 ± 0.015, strength=1.210 ± 0.002, crit=1.968 ± 0.026 per rating point (14 rating = 1%, 27.551 per %), hit=2.658 ± 0.132 per rating point (10 rating = 1%, 26.579 per %), melee_haste=9.991 ± 1.699, intellect=0.040 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 43.3 attack_power points (4.01 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Blood Guard's Mail Helmet (220820, -0.20 DPS) [vendor]; Helm of Fire (8348, -0.24 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (3.44 DPS) | yes | Scout's Medallion (19535, -0.78 DPS) [rep]; Woven Ivy Necklace (19159, -0.90 DPS, sim-verified) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 37.2 attack_power points (3.45 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.11 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Cloak (8259, -0.89 DPS) [world_drop]; Pridelord Cape (14673, -0.99 DPS) [world_drop]; Blackflame Cape (13109, -1.00 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 84.0 attack_power points (7.78 DPS) | yes | Stone Guard's Chain Armor (220827, -2.02 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -2.35 DPS) [quest]; Warbear Harness (15064, -2.56 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 40.8 attack_power points (3.77 DPS) | yes | Pridelord Bands (14672, -1.00 DPS) [world_drop]; Skulker's Leather Bracers (252540, -1.12 DPS, sim-verified) [crafted]; Bracers of the Stone Princess (17714, -1.16 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -3.02 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 48.0 attack_power points (4.45 DPS) | yes | Defiler's Chain Girdle (20151, -0.05 DPS) [rep]; Defiler's Leather Girdle (20193, -0.05 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.22 DPS) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | 61.5 attack_power points (5.69 DPS) | yes | Serpentskin Leggings (8262, -0.69 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.04 DPS) [world_drop]; Stormshroud Pants (15057, -1.42 DPS, sim-verified) [crafted] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 47.9 attack_power points (4.43 DPS) | yes | Fleetfoot Greaves (11627, -0.20 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.32 DPS) [crafted]; Elven Chain Boots (13125, -0.44 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.6 attack_power points (4.31 DPS) | yes | Masons Fraternity Ring (9533, -1.21 DPS) [quest]; Ring of the Underwood (2951, -1.87 DPS) [world_drop]; Falcon's Hook (7552, -1.98 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 33.6 attack_power points (3.12 DPS) | yes | Masons Fraternity Ring (9533, -0.01 DPS) [quest]; Ring of the Underwood (2951, -0.67 DPS) [world_drop]; Falcon's Hook (7552, -0.78 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Stoneraven (13059, +0.00 DPS) [world_drop]; Dawn's Edge (12774, -1.74 DPS, sim-verified) [crafted] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Claw of Celebras (17738, -2.71 DPS) [dungeon]; Thorium Cestus (250614, -3.56 DPS, sim-verified) [crafted]; White Bone Shredder (11863, -5.09 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -1.12 DPS, sim-verified) [world_drop]; Gryphonwing Long Bow (13022, -1.33 DPS) [dungeon]; Skull Splitting Crossbow (13039, -1.47 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Stone Guard's Chain Legplates; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thrash Blade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 845, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 5000000000000000-0050050010000000-500230231051222151)

Set DPS (verified): 366.5. Weights run: 6.6s. Verify run: 61.3s. 1595 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.114 ± 0.016, strength=1.210 ± 0.002, crit=2.048 ± 0.031 per rating point (14 rating = 1%, 28.678 per %), hit=1.634 ± 0.134 per rating point (10 rating = 1%, 16.340 per %), melee_haste=16.916 ± 1.713, intellect=1.247 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 120.5 attack_power points (15.57 DPS) | yes | Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Champion's Chain Greathelm (227080, -0.33 DPS) [vendor]; Outlaw's Collar (279253, -6.43 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.5 attack_power points (4.33 DPS) | yes | Scout's Medallion (19534, -0.24 DPS) [rep]; Woven Ivy Necklace (19159, -0.94 DPS) [quest]; Sunprism Pendant (20649, -1.05 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 85.8 attack_power points (11.08 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -0.26 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.92 DPS) [rep]; Dark Phantom Cape (13122, -2.58 DPS) [world_drop]; Fel Cape (279269, -2.98 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 123.2 attack_power points (15.93 DPS) | yes | Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Warlord's Chain Chestpiece (16565, -1.00 DPS) [vendor]; Timbermaw Tunic (252484, -4.29 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 53.5 attack_power points (6.91 DPS) | yes | General's Chain Wristguards (16570, -0.32 DPS) [pvp]; Windtalker's Wristguards (19582, -0.71 DPS) [rep]; Legionnaire's Chain Bracers (16517, -1.48 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Raider Gloves (272099, -3.01 DPS, sim-verified) [vendor] |
| waist | Defiler's Chain Girdle (20150) (or Defiler's Leather Girdle (20190)) | The Defilers [rep] | 62.7 attack_power points (8.10 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Defiler's Leather Girdle (20190, +0.00 DPS) [rep]; Might of the Timbermaw (19044, -0.99 DPS) [crafted] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 119.0 attack_power points (15.38 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, -0.45 DPS) [pvp]; Outrider's Leather Pants (22740, -6.24 DPS, sim-verified) [rep] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 54.9 attack_power points (7.10 DPS) | yes | Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 59.0 attack_power points (7.63 DPS) | yes | Legionnaire's Band (19511, -3.60 DPS) [rep]; Masons Fraternity Ring (9533, -3.80 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Band (19511, -0.67 DPS) [rep]; Masons Fraternity Ring (9533, -0.87 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.39 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Ironwood Blade (279259, -4.27 DPS, sim-verified) [crafted] |
| off_hand | Glacial Blade (19099) | Frostwolf Clan [rep] | 616.1 attack_power points (79.62 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Heartseeker (12783, -2.28 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -17.31 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.66 DPS) [quest]; Blackcrow (12651, -0.89 DPS) [dungeon]; Eaglehorn Long Bow (13023, -1.09 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Defiler's Chain Girdle; legs: Outrider's Chain Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Glacial Blade

No-known-source sample (15 of 1595, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (troll, 5000000000000000-0050050010000000-500230231051222151)

Set DPS (verified): 861.9. Weights run: 5.7s. Verify run: 53.3s. 1595 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.368 ± 0.020, strength=1.210 ± 0.002, crit=2.765 ± 0.040 per rating point (14 rating = 1%, 38.711 per %), hit=4.129 ± 0.293 per rating point (10 rating = 1%, 41.288 per %), melee_haste=20.858 ± 3.180, intellect=1.238 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 135.5 attack_power points (25.43 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -13.16 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.8 attack_power points (6.91 DPS) | yes | Scout's Medallion (19534, -0.25 DPS) [rep]; Woven Ivy Necklace (19159, -1.55 DPS) [quest]; Sunprism Pendant (20649, -1.58 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 102.6 attack_power points (19.27 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -1.82 DPS) [rep]; Cloak of Flames (3475, -2.67 DPS) [world_drop]; Fel Cape (279269, -3.16 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 167.6 attack_power points (31.46 DPS) | yes | Warlord's Chain Armor (231566, -3.01 DPS) [vendor]; Warlord's Chain Chestpiece (16565, -6.30 DPS) [vendor]; Savage Gladiator Chain (11726, -14.01 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 58.3 attack_power points (10.94 DPS) | yes | General's Chain Wristguards (16570, -0.43 DPS) [pvp]; Windtalker's Wristguards (19582, -1.95 DPS) [rep]; Legionnaire's Chain Bracers (16517, -2.25 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Raider Gloves (272099, -4.88 DPS, sim-verified) [vendor] |
| waist | Defiler's Chain Girdle (20150) (or Defiler's Leather Girdle (20190)) | The Defilers [rep] | 72.7 attack_power points (13.65 DPS) | yes | Defiler's Leather Girdle (20190, +0.00 DPS) [rep]; General's Chain Girdle (16572, -0.33 DPS) [pvp]; Defiler's Lizardhide Girdle (20171, -2.43 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 162.9 attack_power points (30.57 DPS) | yes | General's Chain Legplates (231567, -2.12 DPS) [vendor]; General's Chain Legguards (231574, -5.41 DPS) [pvp]; Outrider's Leather Pants (22740, -14.04 DPS, sim-verified) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.1 attack_power points (11.66 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 94.0 attack_power points (17.64 DPS) | yes | Legionnaire's Band (19510, -10.03 DPS) [rep]; Band of the Penitent (13217, -10.38 DPS) [quest]; Don Rodrigo's Band (21563, -10.38 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 61.3 attack_power points (11.50 DPS) | yes | Band of the Penitent (13217, -4.24 DPS) [quest]; Don Rodrigo's Band (21563, -4.24 DPS) [rep]; Legionnaire's Band (19510, -5.04 DPS, sim-verified) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+11.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.08 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Mark of the Chosen (17774, -6.06 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Ironwood Blade (279259, -24.26 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 626.6 attack_power points (117.62 DPS) | yes | Glacial Blade (19099, +0.00 DPS) [rep]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -27.10 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -2.86 DPS) [world_drop]; Bloodseeker (19107, -4.16 DPS) [quest]; Eaglehorn Long Bow (13023, -4.64 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Defiler's Chain Girdle; legs: Outrider's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket2: Molten Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Heartseeker; ranged: Blackcrow

No-known-source sample (15 of 1595, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

