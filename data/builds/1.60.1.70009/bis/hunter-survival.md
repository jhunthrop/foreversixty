# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 75.0. Weights run: 2.4s. Verify run: 3.4s. 220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.122 ± 0.014, strength=1.210 ± 0.002, crit=0.832 ± 0.015 per rating point (14 rating = 1%, 11.646 per %), hit=1.129 ± 0.042 per rating point (10 rating = 1%, 11.287 per %), melee_haste=not significant (1.723 ± 1.303)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.0 attack_power points (0.74 DPS) | yes | Defender's Leather Hood (252447, -0.32 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.7 attack_power points (0.55 DPS) | yes | Erudite's Amulet (277204, -0.18 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.6 attack_power points (0.46 DPS) | yes | Slime-encrusted Pads (6461, -0.44 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.7 attack_power points (0.55 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.17 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 23.3 attack_power points (1.02 DPS) | yes | Brawler's Leather Armor (252490, -0.11 DPS) [crafted]; Prospector's Chestpiece (14562, -0.36 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.37 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.9 attack_power points (0.47 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (75.0 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.02 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.78 DPS) | yes | Brawler's Leather Belt (252428, -0.20 DPS) [crafted]; Dusty Belt (279897, -0.32 DPS) [quest]; Deviate Scale Belt (6468, -2.71 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (75.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.16 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (75.0 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.07 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.3 attack_power points (0.58 DPS) | yes | Demon Band (12054, -0.37 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; The 1 Ring (8350, -0.43 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.7 attack_power points (0.55 DPS) | yes | Demon Band (12054, -0.30 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon]; The 1 Ring (8350, -0.41 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.83 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 attack_power points (10.30 DPS) | yes | Cruel Barb (5191, -0.91 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.5 attack_power points (0.37 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Owlsight Rifle (15205, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 100.5. Weights run: 2.7s. Verify run: 2.0s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.847 ± 0.012, strength=1.210 ± 0.002, crit=0.902 ± 0.016 per rating point (14 rating = 1%, 12.626 per %), hit=1.431 ± 0.052 per rating point (10 rating = 1%, 14.312 per %), melee_haste=6.357 ± 1.318

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.5 attack_power points (0.95 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.20 DPS) [crafted] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 14.8 attack_power points (0.76 DPS) | yes | Ghostshard Talisman (7731, -0.04 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Erudite's Amulet (277204, -0.38 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.4 attack_power points (1.36 DPS) | yes | Mantle of Thieves (2264, -0.28 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.57 DPS) [crafted]; Bristlebark Amice (14573, -0.60 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.6 attack_power points (0.85 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.22 DPS) [pvp]; Cloak of Night (4447, -0.28 DPS) [world] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 25.9 attack_power points (1.33 DPS) | yes | Brawler's Leather Tunic (252508, -0.20 DPS) [crafted]; Tunic of Westfall (2041, -0.29 DPS) [quest]; Brawler's Leather Armor (252490, -0.35 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.9 attack_power points (0.82 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.25 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.6 attack_power points (1.01 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.18 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.24 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.01 DPS) [crafted]; Prowler's Leather Belt (252459, -0.10 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.34 DPS) | yes | Troll's Bane Leggings (13114, -0.01 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.05 DPS) [crafted]; Petrolspill Leggings (9509, -1.38 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.4 attack_power points (0.95 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.5 attack_power points (1.11 DPS) | yes | Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Monkey Ring (6748, -0.44 DPS) [quest]; Ring of Precision (1491, -0.54 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.3 attack_power points (0.95 DPS) | yes | Thunderbrow Ring (13097, -0.16 DPS) [world_drop]; Monkey Ring (6748, -0.28 DPS) [quest]; Ring of Precision (1491, -0.37 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.62 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (100.5 DPS) | yes | Shoni's Disarming Tool (9608, -4.92 DPS) [quest]; Satyr's Rod (15962, -16.41 DPS) [world_drop]; Swinetusk Shank (6691, -21.62 DPS, sim-verified) [dungeon] |
| ranged | Silver Star (3463) | Stealing Supplies [quest] | 9.2 attack_power points (0.48 DPS) | yes | Double-barreled Shotgun (2098, -0.00 DPS) [world_drop]; Booty Bay Bruiser's Buckshot (274748, -0.01 DPS) [vendor]; BKP "Sparrow" Smallbore (3042, -0.10 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Silver Star

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 145.9. Weights run: 2.8s. Verify run: 2.4s. 597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.839 ± 0.010, strength=1.210 ± 0.001, crit=0.906 ± 0.015 per rating point (14 rating = 1%, 12.680 per %), hit=1.474 ± 0.083 per rating point (10 rating = 1%, 14.740 per %), melee_haste=8.739 ± 1.188

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.5 attack_power points (2.27 DPS) | yes | Raging Berserker's Helm (7719, -0.35 DPS) [dungeon]; Barbaric Iron Helm (7915, -0.41 DPS) [crafted]; Hawkeye's Helm (14591, -0.57 DPS) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 20.2 attack_power points (1.37 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.02 DPS) [quest]; Ghostshard Talisman (7731, -0.42 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.54 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.2 attack_power points (2.18 DPS) | yes | Forest Tracker Epaulets (2278, -0.43 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.53 DPS) [dungeon]; Nightscape Shoulders (8192, -0.81 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 23.2 attack_power points (1.57 DPS) | yes | Sergeant Major's Cape (16336, -0.33 DPS) [pvp]; Hawkeye's Cloak (14593, -0.46 DPS) [world_drop]; Parachute Cloak (10518, -0.58 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 44.6 attack_power points (3.02 DPS) | yes | Kolkar Marauder Chain (6773, -1.00 DPS) [quest]; Tough Scorpid Breastplate (8203, -1.15 DPS) [crafted]; Wolffear Harness (13110, -1.23 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.35 DPS) | yes | Hawkeye's Bracers (14590, -0.28 DPS) [world_drop]; Ravager's Armguards (14770, -0.28 DPS) [world_drop]; Dusky Bracers (7378, -0.36 DPS) [crafted] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 32.9 attack_power points (2.23 DPS) | yes | Gloves of Holy Might (867, -0.02 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.06 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.25 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.03 DPS) | yes | Ogron's Sash (13117, -0.17 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.41 DPS) [rep]; Skulker's Leather Belt (252520, -0.42 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 38.6 attack_power points (2.61 DPS) | yes | Triprunner Dungarees (9624, -0.49 DPS, sim-verified) [quest]; Veteran's Silvered Chain Leggings (250523, -0.84 DPS) [crafted]; Ferine Leggings (6690, -0.85 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.3 attack_power points (1.98 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 24.4 attack_power points (1.65 DPS) | yes | Ring of the Underwood (2951, -0.24 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon]; Mark of Kern (2262, -0.30 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.4 attack_power points (1.45 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Falcon's Hook (7552, -0.08 DPS) [dungeon]; Mark of Kern (2262, -0.09 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (145.9 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.26 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 447.5 attack_power points (30.27 DPS) | yes | Vanquisher's Sword (10823, -0.89 DPS, sim-verified) [quest]; Shoni's Disarming Tool (9608, -15.07 DPS) [quest]; Stonecloth Branch (15963, -30.02 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (145.9 DPS) | yes | Swiftwind (13038, -0.08 DPS) [world_drop]; Monolithic Bow (9426, -0.08 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.02 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 199.0. Weights run: 2.8s. Verify run: 4.9s. 754 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.777 ± 0.010, strength=1.210 ± 0.001, crit=1.006 ± 0.017 per rating point (14 rating = 1%, 14.084 per %), hit=1.497 ± 0.085 per rating point (10 rating = 1%, 14.972 per %), melee_haste=8.562 ± 1.253

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 42.5 attack_power points (3.17 DPS) | yes | Bloomsprout Headpiece (17767, -0.49 DPS) [dungeon]; White Bandit Mask (10008, -0.72 DPS) [crafted]; Embrace of the Lycan (9479, -4.88 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.2 attack_power points (2.17 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.68 DPS) [quest]; Sentinel's Medallion (19540, -0.72 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 37.2 attack_power points (2.77 DPS) | yes | Skulker's Leather Shoulder (252535, -0.50 DPS) [crafted]; Failed Flying Experiment (9647, -0.55 DPS) [quest]; Sunburn Spaulders (274751, -3.35 DPS, sim-verified) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 32.1 attack_power points (2.40 DPS) | yes | Blisterbane Wrap (12552, -0.41 DPS) [dungeon]; Dark Phantom Cape (13122, -0.41 DPS) [world_drop]; Dark Hooded Cape (5257, -0.71 DPS) [world] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blazewind Breastplate (11193, +0.00 DPS) [quest]; Fungus Shroud Armor (17742, +0.00 DPS) [dungeon]; Warbear Harness (15064, -3.98 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.5 attack_power points (2.35 DPS) | yes | Bracers of the Stone Princess (17714, -0.26 DPS) [dungeon]; Arena Bands (18711, -0.26 DPS) [world]; Skulker's Leather Bracers (252540, -0.52 DPS) [crafted] |
| hands | Sergeant Major's Chain Gauntlets (220829) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Raider Gloves (272100, +0.00 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 39.7 attack_power points (2.96 DPS) | yes | Substandard Belt Chain (274757, -0.00 DPS) [vendor]; Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.11 DPS) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Basilisk Hide Pants (1718, -0.12 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -3.95 DPS, sim-verified) [quest] |
| feet | Sergeant Major's Chain Sabatons (220837) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Albino Crocscale Boots (17728, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Sandstalker Ankleguards (12470, -1.76 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.0 attack_power points (2.61 DPS) | yes | Masons Fraternity Ring (9533, -0.75 DPS) [quest]; Ironspine's Eye (7686, -1.05 DPS) [dungeon]; Ring of the Underwood (2951, -1.10 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 28.1 attack_power points (2.09 DPS) | yes | Masons Fraternity Ring (9533, -0.24 DPS) [quest]; Ironspine's Eye (7686, -0.54 DPS) [dungeon]; Ring of the Underwood (2951, -0.59 DPS) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.43 DPS, sim-verified) [crafted] |
| trinket2 | Devilsaur Tooth (19992) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-verified (199.0 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Shadowblade (2163, -12.30 DPS, sim-verified) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Claw of Celebras (17738, -3.30 DPS) [dungeon]; Shoni's Disarming Tool (9608, -22.48 DPS) [quest]; Grizzle's Skinner (11702, -27.91 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -0.52 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.54 DPS) [world_drop]; Dark Iron Rifle (16004, -2.30 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Blackveil Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; hands: Sergeant Major's Chain Gauntlets; waist: Girdle of Beastial Fury; legs: Knight's Chain Legplates; feet: Sergeant Major's Chain Sabatons; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Devilsaur Eye; trinket2: Devilsaur Tooth; main_hand: Bloodrazor; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 301.3. Weights run: 2.6s. Verify run: 9.3s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.154 ± 0.025, strength=1.210 ± 0.002, crit=2.357 ± 0.049 per rating point (14 rating = 1%, 32.993 per %), hit=3.225 ± 0.253 per rating point (10 rating = 1%, 32.252 per %), melee_haste=not significant (11.626 ± 3.665)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -7.63 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.6 attack_power points (5.16 DPS) | yes | Medallion of the Dawn (22659, -0.60 DPS) [quest]; Beads of Ogre Might (22150, -0.66 DPS) [quest]; Mark of Fordring (15411, -2.52 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 91.1 attack_power points (7.28 DPS) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.21 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 attack_power points (4.81 DPS) | yes | Cloak of the Honor Guard (20073, -1.24 DPS) [rep]; Windshear Cape (20691, -1.46 DPS) [world]; Cape of the Black Baron (13340, -1.98 DPS, sim-verified) [dungeon] |
| chest | Beastmaster's Tunic (226886) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Chain Armor (227089, +0.00 DPS) [vendor]; Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Dawn Armor (252483, -5.28 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -6.10 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Beaststalker's Gloves (16676, -3.88 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -4.62 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 173.6 attack_power points (13.87 DPS) | yes | Marshal's Chain Legplates (231558, -3.69 DPS) [vendor]; Sentinel's Leather Pants (237818, -3.95 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -4.41 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -5.71 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (301.3 DPS) | yes | Tarnished Elven Ring (18500, -1.33 DPS) [dungeon]; Cutthroat's Signet (272408, -1.50 DPS) [vendor]; The Postmaster's Seal (13392, -3.04 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.52 DPS) [dungeon]; Cutthroat's Signet (272408, -0.69 DPS) [vendor]; Naglering (11669, -7.18 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Burst of Knowledge (11832, -9.98 DPS, sim-verified) [dungeon] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -14.55 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.4 attack_power points (66.97 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -6.50 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -20.69 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.26 DPS) [dungeon]; The Purifier (22656, -0.46 DPS) [quest]; Dark Iron Rifle (16004, -3.55 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Beastmaster's Tunic; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Second Wind; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 800.5. Weights run: 2.3s. Verify run: 8.2s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.339 ± 0.027, strength=1.210 ± 0.002, crit=2.842 ± 0.053 per rating point (14 rating = 1%, 39.782 per %), hit=4.680 ± 0.448 per rating point (10 rating = 1%, 46.797 per %), melee_haste=not significant (13.407 ± 5.621)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (800.5 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -8.05 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 81.9 attack_power points (14.00 DPS) | yes | Beads of Ogre Might (22150, -1.89 DPS) [quest]; Mark of Fordring (15411, -2.75 DPS) [quest]; Medallion of the Dawn (22659, -3.09 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 117.6 attack_power points (20.10 DPS) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Field Marshal's Chain Spaulders (16468, -1.71 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 74.8 attack_power points (12.79 DPS) | yes | Cape of the Black Baron (13340, -3.37 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.79 DPS) [vendor]; Stalwart Cloak (272415, -4.79 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (800.5 DPS) | yes | Field Marshal's Chain Armor (231563, -6.12 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -6.71 DPS) [vendor]; Tunic of Undead Slaying (23089, -26.46 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (800.5 DPS) | yes | Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Slashclaw Bracers (13211, -3.56 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 114.6 attack_power points (19.60 DPS) | yes | Marshal's Chain Grips (231560, -2.97 DPS) [pvp]; Stormshroud Gloves (21278, -3.84 DPS, sim-verified) [crafted]; Marshal's Chain Vices (231578, -4.40 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 98.8 attack_power points (16.89 DPS) | yes | Marksman's Girdle (22232, -0.50 DPS) [dungeon]; Highlander's Chain Girdle (20043, -4.28 DPS) [rep]; Highlander's Leather Girdle (20045, -4.28 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 208.2 attack_power points (35.59 DPS) | yes | Sentinel's Leather Pants (237818, -11.20 DPS) [vendor]; Marshal's Chain Legplates (231558, -12.12 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -12.71 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 93.6 attack_power points (16.00 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (800.5 DPS) | yes | Tarnished Elven Ring (18500, -3.54 DPS) [dungeon]; Cutthroat's Signet (272408, -3.94 DPS) [vendor]; Naglering (11669, -16.14 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (800.5 DPS) | yes | Tarnished Elven Ring (18500, -1.20 DPS) [dungeon]; Cutthroat's Signet (272408, -1.60 DPS) [vendor]; Naglering (11669, -13.97 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (800.5 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Devilsaur Eye (19991, -11.69 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (800.5 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -7.31 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (800.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -75.43 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 840.1 attack_power points (143.61 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -10.51 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -44.55 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (800.5 DPS) | yes | Blackcrow (12651, -0.80 DPS) [dungeon]; The Purifier (22656, -2.40 DPS) [quest]; Dark Iron Rifle (16004, -7.35 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 75.3. Weights run: 2.4s. Verify run: 3.3s. 209 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.122 ± 0.014, strength=1.210 ± 0.002, crit=0.832 ± 0.015 per rating point (14 rating = 1%, 11.646 per %), hit=1.129 ± 0.042 per rating point (10 rating = 1%, 11.287 per %), melee_haste=not significant (1.723 ± 1.303)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.0 attack_power points (0.74 DPS) | yes | Defender's Leather Hood (252447, -0.32 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.7 attack_power points (0.55 DPS) | yes | Erudite's Amulet (277204, -0.18 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.6 attack_power points (0.46 DPS) | yes | Slime-encrusted Pads (6461, -0.47 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.7 attack_power points (0.55 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.17 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (75.3 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.97 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.6 attack_power points (0.46 DPS) | yes | Bristlebark Bindings (14569, -0.08 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Ratchet Wristwraps (274742, -0.18 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 17.6 attack_power points (0.76 DPS) | yes | Bristlebark Gloves (14572, -0.18 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.18 DPS) [crafted]; Serpent Gloves (5970, -0.21 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.78 DPS) | yes | Brawler's Leather Belt (252428, -0.20 DPS) [crafted]; Murloc Scale Belt (5780, -0.40 DPS) [crafted]; Deviate Scale Belt (6468, -2.66 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (75.3 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.12 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (75.3 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.02 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.3 attack_power points (0.58 DPS) | yes | Bounty Hunter's Ring (5351, -0.30 DPS) [quest]; Demon Band (12054, -0.37 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.7 attack_power points (0.55 DPS) | yes | Bounty Hunter's Ring (5351, -0.28 DPS, sim-verified) [quest]; Demon Band (12054, -0.34 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.83 DPS) | yes | Crescent Staff (6505, +0.00 DPS) [quest]; Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 attack_power points (10.30 DPS) | yes | Wingblade (6504, -6.50 DPS, sim-verified) [quest]; Tork Wrench (11855, -10.20 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.5 attack_power points (0.37 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Privateer Musket (5309, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 101.4. Weights run: 2.7s. Verify run: 2.0s. 352 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.847 ± 0.012, strength=1.210 ± 0.002, crit=0.902 ± 0.016 per rating point (14 rating = 1%, 12.626 per %), hit=1.431 ± 0.052 per rating point (10 rating = 1%, 14.312 per %), melee_haste=6.357 ± 1.318

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.5 attack_power points (0.95 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.20 DPS) [crafted] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 14.8 attack_power points (0.76 DPS) | yes | Ghostshard Talisman (7731, -0.04 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Erudite's Amulet (277204, -0.38 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.4 attack_power points (1.36 DPS) | yes | Mantle of Thieves (2264, -0.28 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.57 DPS) [crafted]; Bristlebark Amice (14573, -0.60 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.6 attack_power points (0.85 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Cloak of Night (4447, -0.28 DPS) [world]; Swiftrunner Cape (6745, -0.28 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 25.9 attack_power points (1.33 DPS) | yes | Brawler's Leather Tunic (252508, -0.20 DPS) [crafted]; Panther Armor (6670, -0.35 DPS) [quest]; Brawler's Leather Armor (252490, -0.35 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.9 attack_power points (0.82 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.25 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.6 attack_power points (1.01 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.18 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.24 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.01 DPS) [crafted]; Prowler's Leather Belt (252459, -0.10 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.34 DPS) | yes | Troll's Bane Leggings (13114, -0.01 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.05 DPS) [crafted]; Petrolspill Leggings (9509, -1.33 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.4 attack_power points (0.95 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.5 attack_power points (1.11 DPS) | yes | Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Monkey Ring (6748, -0.44 DPS) [quest]; Band of the Fist (17694, -0.48 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.3 attack_power points (0.95 DPS) | yes | Thunderbrow Ring (13097, -0.16 DPS) [world_drop]; Monkey Ring (6748, -0.28 DPS) [quest]; Band of the Fist (17694, -0.32 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.62 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (101.4 DPS) | yes | Tork Wrench (11855, -16.39 DPS) [quest]; Satyr's Rod (15962, -16.41 DPS) [world_drop]; Swinetusk Shank (6691, -21.97 DPS, sim-verified) [dungeon] |
| ranged | Silver Star (3463) | Stealing Supplies [quest] | 9.2 attack_power points (0.48 DPS) | yes | Double-barreled Shotgun (2098, -0.00 DPS) [world_drop]; Booty Bay Bruiser's Buckshot (274748, -0.01 DPS) [vendor]; BKP "Sparrow" Smallbore (3042, -0.10 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Silver Star

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 145.8. Weights run: 2.8s. Verify run: 2.4s. 563 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.839 ± 0.010, strength=1.210 ± 0.001, crit=0.906 ± 0.015 per rating point (14 rating = 1%, 12.680 per %), hit=1.474 ± 0.083 per rating point (10 rating = 1%, 14.740 per %), melee_haste=8.739 ± 1.188

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.5 attack_power points (2.27 DPS) | yes | Raging Berserker's Helm (7719, -0.35 DPS) [dungeon]; Barbaric Iron Helm (7915, -0.41 DPS) [crafted]; Hawkeye's Helm (14591, -0.57 DPS) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 20.2 attack_power points (1.37 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.02 DPS) [quest]; Ghostshard Talisman (7731, -0.42 DPS) [dungeon]; Ethereal Talisman (4430, -0.46 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.2 attack_power points (2.18 DPS) | yes | Forest Tracker Epaulets (2278, -0.43 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.53 DPS) [dungeon]; Nightscape Shoulders (8192, -0.81 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 23.2 attack_power points (1.57 DPS) | yes | First Sergeant's Cloak (16340, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.46 DPS) [world_drop]; Parachute Cloak (10518, -0.58 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 44.6 attack_power points (3.02 DPS) | yes | Wolffear Harness (13110, -0.88 DPS, sim-verified) [world_drop]; Kolkar Marauder Chain (6773, -1.00 DPS) [quest]; Tough Scorpid Breastplate (8203, -1.15 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.35 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.28 DPS) [world_drop] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 32.9 attack_power points (2.23 DPS) | yes | Gloves of Holy Might (867, -0.02 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.06 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.25 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.03 DPS) | yes | Ogron's Sash (13117, -0.17 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.41 DPS) [rep]; Skulker's Leather Belt (252520, -0.42 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 38.6 attack_power points (2.61 DPS) | yes | Triprunner Dungarees (9624, -0.13 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.84 DPS) [crafted]; Ferine Leggings (6690, -0.85 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.3 attack_power points (1.98 DPS) | yes | Imperial Leather Boots (6431, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.43 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 24.4 attack_power points (1.65 DPS) | yes | Ring of the Underwood (2951, -0.24 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon]; Mark of Kern (2262, -0.30 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.4 attack_power points (1.45 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Falcon's Hook (7552, -0.08 DPS) [dungeon]; Mark of Kern (2262, -0.09 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (145.8 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.32 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 447.5 attack_power points (30.27 DPS) | yes | Vanquisher's Sword (10823, +0.00 DPS) [quest]; Stonecloth Branch (15963, -30.02 DPS) [world_drop]; Tork Wrench (11855, -30.11 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (145.8 DPS) | yes | Swiftwind (13038, -0.08 DPS) [world_drop]; Monolithic Bow (9426, -0.08 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.02 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Legionnaire's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 189.7. Weights run: 2.8s. Verify run: 2.6s. 713 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.777 ± 0.010, strength=1.210 ± 0.001, crit=1.006 ± 0.017 per rating point (14 rating = 1%, 14.084 per %), hit=1.497 ± 0.085 per rating point (10 rating = 1%, 14.972 per %), melee_haste=8.562 ± 1.253

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 41.7 attack_power points (3.11 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.42 DPS) [dungeon]; White Bandit Mask (10008, -0.66 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.2 attack_power points (2.17 DPS) | yes | Woven Ivy Necklace (19159, -0.44 DPS) [quest]; Scout's Medallion (19535, -0.58 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.68 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 31.6 attack_power points (2.35 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted]; Failed Flying Experiment (9647, -0.12 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 32.1 attack_power points (2.40 DPS) | yes | Blisterbane Wrap (12552, -0.41 DPS) [dungeon]; Dark Phantom Cape (13122, -0.41 DPS) [world_drop]; Dark Hooded Cape (5257, -0.71 DPS) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 45.3 attack_power points (3.38 DPS) | yes | Blazewind Breastplate (11193, -0.06 DPS) [quest]; Fungus Shroud Armor (17742, -0.06 DPS) [dungeon]; Quillward Harness (10583, -0.14 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.5 attack_power points (2.35 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 51.2 attack_power points (3.82 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.77 DPS, sim-verified) [crafted]; Gloves of Holy Might (867, -1.28 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -1.36 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 39.7 attack_power points (2.96 DPS) | yes | Substandard Belt Chain (274757, -0.00 DPS) [vendor]; Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.11 DPS) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Basilisk Hide Pants (1718, -0.12 DPS) [world_drop]; Triprunner Dungarees (9624, -0.25 DPS) [quest]; Serpentskin Leggings (8262, -1.94 DPS, sim-verified) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 37.5 attack_power points (2.79 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Albino Crocscale Boots (17728, -0.14 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.16 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.0 attack_power points (2.61 DPS) | yes | Masons Fraternity Ring (9533, -0.75 DPS) [quest]; White Bone Band (11862, -0.82 DPS) [quest]; Ironspine's Eye (7686, -1.05 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 28.1 attack_power points (2.09 DPS) | yes | Masons Fraternity Ring (9533, -0.24 DPS) [quest]; White Bone Band (11862, -0.31 DPS) [quest]; Ironspine's Eye (7686, -0.54 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.91 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Devilsaur Tooth (19992, -1.87 DPS, sim-verified) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodrazor (809, +0.00 DPS) [world_drop]; Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (+32.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Claw of Celebras (17738, -3.30 DPS) [dungeon]; White Bone Shredder (11863, -5.53 DPS) [quest]; Grizzle's Skinner (11702, -32.81 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -0.52 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.54 DPS) [world_drop]; Dark Iron Rifle (16004, -1.38 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Chain Legplates; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Shadowblade; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 303.3. Weights run: 2.6s. Verify run: 9.5s. 1650 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.154 ± 0.025, strength=1.210 ± 0.002, crit=2.357 ± 0.049 per rating point (14 rating = 1%, 32.993 per %), hit=3.225 ± 0.253 per rating point (10 rating = 1%, 32.252 per %), melee_haste=not significant (11.626 ± 3.665)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (303.3 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -5.21 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.6 attack_power points (5.16 DPS) | yes | Mark of Fordring (15411, -0.44 DPS) [quest]; Medallion of the Dawn (22659, -0.60 DPS) [quest]; Beads of Ogre Might (22150, -0.66 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 91.1 attack_power points (7.28 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.21 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 attack_power points (4.81 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -1.24 DPS) [rep]; Windshear Cape (20691, -1.46 DPS) [world] |
| chest | Beastmaster's Tunic (226886) | Saving the Best for Last [quest] | sim-verified (303.3 DPS) | yes | Legionnaire's Chain Armor (227083, +0.00 DPS) [vendor]; Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Dawn Armor (252483, -3.66 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (303.3 DPS) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -5.70 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-verified (303.3 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Beaststalker's Gloves (16676, -4.05 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (303.3 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -4.12 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 173.6 attack_power points (13.87 DPS) | yes | Outrider's Chain Leggings (22673, -2.35 DPS, sim-verified) [rep]; General's Chain Legplates (231567, -3.69 DPS) [vendor]; Sentinel's Leather Pants (237818, -3.95 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (303.3 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -5.49 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (303.3 DPS) | yes | Tarnished Elven Ring (18500, -1.33 DPS) [dungeon]; Cutthroat's Signet (272408, -1.50 DPS) [vendor]; Naglering (11669, -5.75 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (303.3 DPS) | yes | Tarnished Elven Ring (18500, -0.52 DPS) [dungeon]; Cutthroat's Signet (272408, -0.69 DPS) [vendor]; Naglering (11669, -4.27 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (303.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (303.3 DPS) | yes | Hand of Justice (11815, +0.00 DPS, sim-verified) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (303.3 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -14.08 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.4 attack_power points (66.97 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -6.73 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -20.69 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (303.3 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.26 DPS) [dungeon]; The Purifier (22656, -0.46 DPS) [quest] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Beastmaster's Tunic; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Second Wind; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 803.2. Weights run: 2.3s. Verify run: 8.4s. 1650 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.339 ± 0.027, strength=1.210 ± 0.002, crit=2.842 ± 0.053 per rating point (14 rating = 1%, 39.782 per %), hit=4.680 ± 0.448 per rating point (10 rating = 1%, 46.797 per %), melee_haste=not significant (13.407 ± 5.621)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 133.4 attack_power points (22.80 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 81.9 attack_power points (14.00 DPS) | yes | Mark of Fordring (15411, -2.75 DPS) [quest]; Medallion of the Dawn (22659, -3.09 DPS) [quest]; Beads of Ogre Might (22150, -3.45 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (803.2 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -7.70 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 74.8 attack_power points (12.79 DPS) | yes | Cape of the Black Baron (13340, -3.37 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.79 DPS) [vendor]; Stalwart Cloak (272415, -4.79 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (803.2 DPS) | yes | Warlord's Chain Armor (231566, -6.12 DPS) [vendor]; Legionnaire's Chain Armor (227083, -6.71 DPS) [vendor]; Tunic of Undead Slaying (23089, -28.51 DPS, sim-verified) [world] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-verified (803.2 DPS) | yes | Forest Stalker's Bracers (19587, -0.93 DPS) [rep]; Blackmist Armguards (12966, -1.76 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -10.78 DPS, sim-verified) [world] |
| hands | Bloodmail Gauntlets (14615) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (803.2 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -3.03 DPS, sim-verified) [quest] |
| waist | Bloodmail Belt (14614) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (803.2 DPS) | yes | Defiler's Chain Girdle (20150, +0.00 DPS) [rep]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Belt of Preserved Heads (20216, -6.10 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 208.2 attack_power points (35.59 DPS) | yes | Outrider's Chain Leggings (22673, -10.30 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -11.20 DPS) [vendor]; General's Chain Legplates (231567, -12.12 DPS) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (803.2 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Windreaver Greaves (13967, -7.65 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (803.2 DPS) | yes | Tarnished Elven Ring (18500, -3.54 DPS) [dungeon]; Cutthroat's Signet (272408, -3.94 DPS) [vendor]; Naglering (11669, -15.95 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (803.2 DPS) | yes | Tarnished Elven Ring (18500, -1.20 DPS) [dungeon]; Cutthroat's Signet (272408, -1.60 DPS) [vendor]; Naglering (11669, -12.99 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (803.2 DPS) | yes | Rune of the Guard Captain (19120, -0.82 DPS) [quest]; Frozen Heart of the Mountain (249469, -6.40 DPS) [crafted]; Counterattack Lodestone (18537, -9.84 DPS) [dungeon] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (803.2 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Devilsaur Eye (19991, -13.47 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (803.2 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -73.86 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 840.1 attack_power points (143.61 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -13.61 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -44.55 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (803.2 DPS) | yes | Blackcrow (12651, -0.80 DPS) [dungeon]; The Purifier (22656, -2.40 DPS) [quest]; Dark Iron Rifle (16004, -8.58 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Slashclaw Bracers; hands: Bloodmail Gauntlets; waist: Bloodmail Belt; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Blackhand's Breadth; trinket2: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

