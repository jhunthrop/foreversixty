# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 37.5. Weights run: 1.8s. Verify run: 2.7s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.392 ± 0.024, crit=0.766 ± 0.026 per rating point (14 rating = 1%, 10.720 per %), hit=1.097 ± 0.048 per rating point (10 rating = 1%, 10.973 per %), melee_haste=6.723 ± 0.500

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.12 DPS) | yes | Defender's Leather Hood (252447, -0.26 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.34 DPS) [crafted]; Brawler's Leather Hood (252504, -0.50 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.47 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.0 attack_power points (0.39 DPS) | yes | Rough Bronze Shoulders (3480, -0.05 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.05 DPS) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 8.8 attack_power points (0.49 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.02 DPS) [dungeon]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Lambent Scale Cloak (4706, -0.04 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.5 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.69 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.56 DPS) | yes | Bravo's Armbands (270015, -0.02 DPS) [quest]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Cryptwalker Bracers (280095, -0.11 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.5 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.93 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (37.5 DPS) | yes | Brawler's Leather Belt (252428, -0.25 DPS) [crafted]; Deviate Scale Belt (6468, -0.28 DPS) [crafted]; Cobrahn's Grasp (6460, -3.21 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.5 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.85 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.1 attack_power points (0.96 DPS) | yes | Brawler's Leather Boots (252439, -0.01 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.40 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.6 attack_power points (0.76 DPS) | yes | Signet of the Zhevra (285330, -0.31 DPS, sim-verified) [world]; The 1 Ring (8350, -0.57 DPS) [world]; Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon] |
| finger2 | Demon Band (12054) | World drop [world_drop] | sim-verified (37.5 DPS) | yes | Signet of the Zhevra (285330, +0.00 DPS) [world]; The 1 Ring (8350, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.93 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 239.3 attack_power points (13.39 DPS) | yes | Diamond Hammer (2194, -0.80 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -12.72 DPS) [world_drop]; Furen's Favor (6970, -13.06 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.6 attack_power points (0.31 DPS) | yes | Dwarven Fishing Pole (3567, -0.09 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.09 DPS) [crafted]; Fine Longbow (11304, -0.09 DPS) [vendor] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 66.6. Weights run: 1.9s. Verify run: 1.4s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.517 ± 0.032, crit=1.125 ± 0.037 per rating point (14 rating = 1%, 15.752 per %), hit=1.410 ± 0.073 per rating point (10 rating = 1%, 14.102 per %), melee_haste=10.298 ± 0.816

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 31.7 attack_power points (1.85 DPS) | yes | Tusken Helm (6686, -0.33 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.45 DPS) [crafted]; Defender's Leather Helm (252455, -0.45 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.1 attack_power points (0.82 DPS) | yes | Ghostshard Talisman (7731, -0.00 DPS) [dungeon]; Sentinel's Medallion (19541, -0.11 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.7 attack_power points (1.56 DPS) | yes | Barbaric Iron Shoulders (7913, -0.42 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -0.53 DPS) [crafted]; Mantle of Thieves (2264, -0.67 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.6 attack_power points (0.97 DPS) | yes | Sergeant Major's Cape (16315, -0.15 DPS) [pvp]; Tigerstrike Mantle (13108, -0.26 DPS) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.75 DPS) | yes | Veteran's Silvered Chain Shirt (250518, -0.05 DPS) [crafted]; Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Slayer's Surcoat (14751, -0.32 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.1 attack_power points (1.00 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Pugilist Bracers (4438, -0.06 DPS) [dungeon]; Barbaric Bracers (18948, -0.18 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.1 attack_power points (1.35 DPS) | yes | Mail Combat Gauntlets (4075, -0.06 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.06 DPS) [world]; Toughened Leather Gloves (4253, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 29.1 attack_power points (1.70 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skulker's Leather Belt (252520, -0.20 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.30 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 32.6 attack_power points (1.91 DPS) | yes | Brawler's Leather Legguards (252516, -0.29 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.35 DPS) [world_drop]; Ferine Leggings (6690, -0.39 DPS) [dungeon] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 24.6 attack_power points (1.44 DPS) | yes | Brawler's Leather Boots (252439, -0.41 DPS) [crafted]; Alacritous Treads (277234, -0.41 DPS) [quest]; Feet of the Lynx (1121, -0.46 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.7 attack_power points (1.27 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Tiger Band (6749, -0.56 DPS) [quest]; Monkey Ring (6748, -0.65 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.1 attack_power points (1.23 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Tiger Band (6749, -0.53 DPS) [quest]; Monkey Ring (6748, -0.61 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (20.21 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (19.59 DPS) | yes | Royal Diplomatic Scepter (9457, -0.92 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -6.44 DPS) [quest]; Slayer's Shield (15892, -18.62 DPS) [world_drop] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.6 attack_power points (0.62 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.09 DPS) [vendor]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Double-barreled Shotgun

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 127.3. Weights run: 2.0s. Verify run: 1.8s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.816 ± 0.056, crit=1.814 ± 0.072 per rating point (14 rating = 1%, 25.396 per %), hit=2.053 ± 0.118 per rating point (10 rating = 1%, 20.529 per %), melee_haste=8.443 ± 1.023

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 51.4 attack_power points (4.35 DPS) | yes | Chromite Barbute (8142, -0.47 DPS) [dungeon]; White Bandit Mask (10008, -0.80 DPS) [crafted]; Barbaric Iron Helm (7915, -1.44 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.69 DPS) | yes | Sentinel's Medallion (19540, -0.00 DPS) [rep]; Kaleidoscope Chain (13084, -0.40 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.0 attack_power points (2.71 DPS) | yes | Forest Tracker Epaulets (2278, -0.17 DPS) [world_drop]; Flintrock Shoulders (7755, -0.32 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.77 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 26.2 attack_power points (2.21 DPS) | yes | Sergeant Major's Cape (16336, -0.28 DPS) [pvp]; Hawkeye's Cloak (14593, -0.63 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.78 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 50.5 attack_power points (4.27 DPS) | yes | Kolkar Marauder Chain (6773, -0.56 DPS, sim-verified) [quest]; Carapace of Tuten'kash (10775, -1.35 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -1.66 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.3 attack_power points (1.80 DPS) | yes | Branded Leather Bracers (19508, -0.11 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.20 DPS) [world_drop]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 45.4 attack_power points (3.84 DPS) | yes | Scarlet Gauntlets (10331, -0.27 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.61 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.77 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.3 attack_power points (2.91 DPS) | yes | Highlander's Chain Girdle (20089, -0.08 DPS) [rep]; Officer's Belt (250556, -0.29 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.37 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.55 DPS) | yes | Symbolic Legplates (14829, -0.09 DPS) [world_drop]; Triprunner Dungarees (9624, -0.28 DPS) [quest]; Basilisk Hide Pants (1718, -0.33 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 38.7 attack_power points (3.28 DPS) | yes | Blackforge Greaves (6423, -0.22 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.34 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.40 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 30.5 attack_power points (2.58 DPS) | yes | Falcon's Hook (7552, -0.69 DPS) [dungeon]; Ring of the Underwood (2951, -0.71 DPS) [world_drop]; Thunderbrow Ring (13097, -0.77 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 24.3 attack_power points (2.06 DPS) | yes | Falcon's Hook (7552, -0.17 DPS) [dungeon]; Ring of the Underwood (2951, -0.18 DPS) [world_drop]; Thunderbrow Ring (13097, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (127.3 DPS) | yes | Shoni's Disarming Tool (9608, -19.24 DPS) [quest]; Skullance Shield (13081, -36.28 DPS) [world_drop]; Ardent Custodian (868, -42.49 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.18 DPS) [quest]; The Silencer (13138, -0.29 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.19 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 190.8. Weights run: 2.1s. Verify run: 2.4s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.574 ± 0.058, crit=1.884 ± 0.078 per rating point (14 rating = 1%, 26.382 per %), hit=1.955 ± 0.136 per rating point (10 rating = 1%, 19.550 per %), melee_haste=13.338 ± 1.221

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 81.9 attack_power points (7.92 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.17 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -2.86 DPS) [dungeon]; Embrace of the Lycan (9479, -3.28 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.5 attack_power points (2.94 DPS) | yes | Sentinel's Medallion (19539, -1.12 DPS) [rep]; Zealous Shadowshard Pendant (17772, -1.67 DPS, sim-verified) [quest] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 55.3 attack_power points (5.34 DPS) | yes | Officer's Pauldrons (250576, -1.61 DPS) [crafted]; Wyrmslayer Spaulders (13066, -1.81 DPS) [world_drop]; Knight-Lieutenant's Plate Pauldrons (220795, -2.27 DPS, sim-verified) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 34.0 attack_power points (3.29 DPS) | yes | Blisterbane Wrap (12552, -1.01 DPS) [dungeon]; Dark Phantom Cape (13122, -1.01 DPS) [world_drop]; Dark Hooded Cape (5257, -1.60 DPS, sim-verified) [world] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 56.4 attack_power points (5.45 DPS) | yes | Mixologist's Tunic (12793, -0.30 DPS) [dungeon]; Warbear Harness (15064, -0.58 DPS) [crafted]; Warforged Chestplate (11195, -0.81 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.6 attack_power points (3.05 DPS) | yes | Officer's Wristguards (250581, -0.02 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.25 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.33 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 59.7 attack_power points (5.77 DPS) | yes | Raider Gloves (272100, -0.49 DPS) [vendor]; Gloves of Holy Might (867, -1.29 DPS) [world_drop]; Officer's Gloves (250551, -1.31 DPS) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 48.4 attack_power points (4.67 DPS) | yes | Prowler's Leather Waistguard (252473, -0.14 DPS) [crafted]; Highlander's Plate Girdle (20124, -0.19 DPS) [rep]; Highlander's Chain Girdle (20088, -0.19 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 66.1 attack_power points (6.39 DPS) | yes | Centurion Legplates (10740, -1.21 DPS) [quest]; Stormshroud Pants (15057, -1.29 DPS) [crafted]; Gryphon Rider's Leggings (9652, -2.01 DPS, sim-verified) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 48.5 attack_power points (4.68 DPS) | yes | Prowler's Leather Boots (252468, -0.50 DPS) [crafted]; Skulker's Leather Boots (252469, -0.58 DPS) [crafted]; Officer's Sabatons (250561, -0.80 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 39.6 attack_power points (3.82 DPS) | yes | Ironspine's Eye (7686, -1.68 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.69 DPS) [quest]; Thunderbrow Ring (13097, -1.82 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 34.2 attack_power points (3.30 DPS) | yes | Ironspine's Eye (7686, -1.16 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.17 DPS) [quest]; Thunderbrow Ring (13097, -1.30 DPS) [world_drop] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (190.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (190.8 DPS) | yes | Molten Heart of the Mountain (249470, -3.02 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (190.8 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Hanzo Sword (8190, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (52.66 DPS) | yes | Dawn's Edge (12774, +0.00 DPS) [crafted]; Claw of Celebras (17738, -6.07 DPS) [dungeon]; Shoni's Disarming Tool (9608, -30.95 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (190.8 DPS) | yes | Stinging Bow (10624, -0.20 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.20 DPS) [world_drop]; Dark Iron Rifle (16004, -3.05 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Knight's Plate Hauberk; wrist: Deepfury Bracers; hands: Raider Gauntlets; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 300.6. Weights run: 2.1s. Verify run: 7.6s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.343 ± 0.121, crit=3.346 ± 0.173 per rating point (14 rating = 1%, 46.851 per %), hit=2.900 ± 0.318 per rating point (10 rating = 1%, 28.997 per %), melee_haste=18.914 ± 2.945

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 187.7 attack_power points (19.20 DPS) | yes | Outlaw's Collar (279253, -6.98 DPS) [crafted]; Lieutenant Commander's Plate Helm (23314, -7.15 DPS) [vendor]; Eye of Rend (12587, -10.23 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-verified (300.6 DPS) | yes | Medallion of the Dawn (22659, -0.20 DPS) [quest]; Amulet of the Darkmoon (19491, -0.85 DPS) [quest]; Rage of Mugamba (19577, -4.56 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 110.1 attack_power points (11.26 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Field Marshal's Plate Shoulderguards (231537, -0.78 DPS) [pvp]; Truestrike Shoulders (12927, -2.88 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 57.0 attack_power points (5.83 DPS) | yes | Windshear Cape (20691, -0.60 DPS) [world]; Fel Cape (279269, -1.04 DPS) [crafted]; Cape of the Black Baron (13340, -3.33 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (300.6 DPS) | yes | Timbermaw Tunic (252484, -0.77 DPS) [crafted]; Savage Gladiator Chain (11726, -2.81 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -15.85 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (300.6 DPS) | yes | Forest Stalker's Bracers (19587, -0.96 DPS) [rep]; Vambraces of the Sadist (13400, -1.53 DPS) [dungeon]; Bracers of Undead Slaying (23090, -8.60 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (300.6 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gauntlets of Heroism (226861, -0.25 DPS) [quest]; Razor Gauntlets (18326, -9.25 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 100.9 attack_power points (10.32 DPS) | yes | Ferocity of the Timbermaw (227805, -0.92 DPS) [vendor]; Valiant's Waistguard (272403, -1.70 DPS) [vendor]; Belt of Preserved Heads (20216, -2.66 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (300.6 DPS) | yes | Sentinel's Plate Legguards (237825, -2.86 DPS) [vendor]; Titanic Leggings (22385, -4.08 DPS) [crafted]; Cloudkeeper Legplates (14554, -8.77 DPS, sim-verified) [world_drop] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 75.8 attack_power points (7.76 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Boots of Heroism (21995, -0.70 DPS) [quest]; Battleboots of Heroism (226857, -0.70 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (300.6 DPS) | yes | Tarnished Elven Ring (18500, -2.84 DPS) [dungeon]; Painweaver Band (13098, -2.97 DPS) [dungeon]; Naglering (11669, -10.82 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (300.6 DPS) | yes | Tarnished Elven Ring (18500, -0.72 DPS) [dungeon]; Painweaver Band (13098, -0.85 DPS) [dungeon]; Naglering (11669, -8.39 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (300.6 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (300.6 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -3.10 DPS, sim-verified) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (300.6 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Felstriker (12590, -2.37 DPS, sim-verified) [dungeon] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (300.6 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -27.72 DPS) [dungeon]; Skullflame Shield (1168, -98.98 DPS, sim-verified) [world_drop] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (300.6 DPS) | yes | Satyr's Bow (18323, -1.11 DPS) [dungeon]; Blackcrow (12651, -1.40 DPS) [dungeon]; Dark Iron Rifle (16004, -6.39 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Raider Gloves; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Diamond Flask; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60, raid preset (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 876.2. Weights run: 2.5s. Verify run: 8.6s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.708 ± 0.169, crit=3.571 ± 0.234 per rating point (14 rating = 1%, 49.990 per %), hit=4.439 ± 0.425 per rating point (10 rating = 1%, 44.386 per %), melee_haste=32.332 ± 4.248

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 228.4 attack_power points (50.03 DPS) | yes | Mask of the Unforgiven (13404, -19.13 DPS, sim-verified) [dungeon]; Lieutenant Commander's Plate Helm (23314, -19.23 DPS) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (876.2 DPS) | yes | Mark of Fordring (15411, -1.98 DPS) [quest]; Medallion of the Dawn (22659, -2.41 DPS) [quest]; Rage of Mugamba (19577, -4.36 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 123.1 attack_power points (26.97 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -2.26 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.4 attack_power points (15.86 DPS) | yes | Windshear Cape (20691, -3.10 DPS) [world]; Cape of the Black Baron (13340, -4.58 DPS, sim-verified) [dungeon]; Blackveil Cape (11626, -4.66 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (876.2 DPS) | yes | Timbermaw Tunic (252484, -5.09 DPS) [crafted]; Savage Gladiator Chain (11726, -10.18 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -34.82 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (876.2 DPS) | yes | Forest Stalker's Bracers (19587, -4.10 DPS) [rep]; Bracers of Undead Slaying (23090, -14.97 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | sim-verified (876.2 DPS) | yes | Raider Gloves (272099, -3.47 DPS) [vendor]; Stormshroud Gloves (21278, -4.11 DPS) [crafted]; Razor Gauntlets (18326, -15.10 DPS, sim-verified) [dungeon] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 115.8 attack_power points (25.37 DPS) | yes | Radiant Girdle of the Dawn (227814, -1.41 DPS) [vendor]; Ferocity of the Timbermaw (227805, -2.75 DPS) [vendor]; Marksman's Girdle (22232, -3.19 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (876.2 DPS) | yes | Titanic Leggings (22385, -7.53 DPS) [crafted]; Sentinel's Plate Legguards (237825, -7.75 DPS) [vendor]; Cloudkeeper Legplates (14554, -16.10 DPS, sim-verified) [world_drop] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 98.6 attack_power points (21.59 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -2.19 DPS) [dungeon]; Boots of Heroism (21995, -2.23 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (876.2 DPS) | yes | Tarnished Elven Ring (18500, -5.56 DPS) [dungeon]; Cutthroat's Signet (272408, -6.15 DPS) [vendor]; Naglering (11669, -19.54 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (876.2 DPS) | yes | Tarnished Elven Ring (18500, -1.78 DPS) [dungeon]; Cutthroat's Signet (272408, -2.37 DPS) [vendor]; Naglering (11669, -15.11 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (876.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (876.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (876.2 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Felstriker (12590, -5.99 DPS, sim-verified) [dungeon] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (876.2 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -60.65 DPS) [dungeon]; Skullflame Shield (1168, -239.37 DPS, sim-verified) [world_drop] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (876.2 DPS) | yes | The Purifier (22656, -0.55 DPS) [quest]; Blackcrow (12651, -0.97 DPS) [dungeon]; Dark Iron Rifle (16004, -8.40 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (troll, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 37.6. Weights run: 1.8s. Verify run: 2.6s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.392 ± 0.024, crit=0.766 ± 0.026 per rating point (14 rating = 1%, 10.720 per %), hit=1.097 ± 0.048 per rating point (10 rating = 1%, 10.973 per %), melee_haste=6.723 ± 0.500

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.12 DPS) | yes | Defender's Leather Hood (252447, -0.26 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.34 DPS) [crafted]; Brawler's Leather Hood (252504, -0.50 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.4 attack_power points (0.47 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.0 attack_power points (0.39 DPS) | yes | Rough Bronze Shoulders (3480, -0.05 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.05 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 8.8 attack_power points (0.49 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.02 DPS) [dungeon]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Lambent Scale Cloak (4706, -0.04 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.6 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.06 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.56 DPS) | yes | Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Forest Leather Bracers (3202, -0.17 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.22 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.6 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.71 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (37.6 DPS) | yes | Brawler's Leather Belt (252428, -0.25 DPS) [crafted]; Deviate Scale Belt (6468, -0.28 DPS) [crafted]; Cobrahn's Grasp (6460, -2.82 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.6 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.36 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.1 attack_power points (0.96 DPS) | yes | Brawler's Leather Boots (252439, -0.01 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.40 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.6 attack_power points (0.76 DPS) | yes | Demon Band (12054, -0.31 DPS) [world_drop]; Loop of Sacrifice (281673, -0.42 DPS) [quest]; Bounty Hunter's Ring (5351, -0.53 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 8.4 attack_power points (0.47 DPS) | yes | Demon Band (12054, -0.02 DPS) [world_drop]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Bounty Hunter's Ring (5351, -0.23 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.93 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 239.3 attack_power points (13.39 DPS) | yes | Diamond Hammer (2194, -0.39 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -12.72 DPS) [world_drop]; Ruga's Bulwark (7120, -13.06 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.6 attack_power points (0.31 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.09 DPS) [crafted]; Fine Longbow (11304, -0.09 DPS) [vendor]; Deadly Blunderbuss (4369, -0.16 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 67.5. Weights run: 1.9s. Verify run: 1.5s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.517 ± 0.032, crit=1.125 ± 0.037 per rating point (14 rating = 1%, 15.752 per %), hit=1.410 ± 0.073 per rating point (10 rating = 1%, 14.102 per %), melee_haste=10.298 ± 0.816

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 31.7 attack_power points (1.85 DPS) | yes | Tusken Helm (6686, -0.33 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.45 DPS) [crafted]; Defender's Leather Helm (252455, -0.45 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.1 attack_power points (0.82 DPS) | yes | Ghostshard Talisman (7731, -0.00 DPS) [dungeon]; Scout's Medallion (19537, -0.11 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.7 attack_power points (1.56 DPS) | yes | Barbaric Iron Shoulders (7913, -0.33 DPS) [crafted]; Barbaric Shoulders (5964, -0.53 DPS) [crafted]; Mantle of Thieves (2264, -0.67 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.6 attack_power points (0.97 DPS) | yes | Tigerstrike Mantle (13108, -0.26 DPS) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon]; Wildhunter Cloak (16658, -0.39 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.75 DPS) | yes | Veteran's Silvered Chain Shirt (250518, -0.05 DPS) [crafted]; Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Slayer's Surcoat (14751, -0.32 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.1 attack_power points (1.00 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Pugilist Bracers (4438, -0.06 DPS) [dungeon]; Barbaric Bracers (18948, -0.18 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.1 attack_power points (1.35 DPS) | yes | Mail Combat Gauntlets (4075, -0.06 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.06 DPS) [world]; Toughened Leather Gloves (4253, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 29.1 attack_power points (1.70 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skulker's Leather Belt (252520, -0.20 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.30 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 32.6 attack_power points (1.91 DPS) | yes | Brawler's Leather Legguards (252516, -0.29 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.35 DPS) [world_drop]; Ferine Leggings (6690, -0.39 DPS) [dungeon] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 24.6 attack_power points (1.44 DPS) | yes | Feet of the Lynx (1121, -0.38 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.41 DPS) [crafted]; Veteran's Boots (250503, -0.50 DPS) [crafted] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.7 attack_power points (1.27 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Band of the Fist (17694, -0.44 DPS) [quest]; Tiger Band (6749, -0.56 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.1 attack_power points (1.23 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Band of the Fist (17694, -0.41 DPS) [quest]; Tiger Band (6749, -0.53 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (20.21 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (19.59 DPS) | yes | Royal Diplomatic Scepter (9457, -0.83 DPS, sim-verified) [dungeon]; Slayer's Shield (15892, -18.62 DPS) [world_drop]; Shield of Thorsen (13079, -18.77 DPS) [world_drop] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.6 attack_power points (0.62 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.09 DPS) [vendor]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Double-barreled Shotgun

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 128.4. Weights run: 2.0s. Verify run: 1.8s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.816 ± 0.056, crit=1.814 ± 0.072 per rating point (14 rating = 1%, 25.396 per %), hit=2.053 ± 0.118 per rating point (10 rating = 1%, 20.529 per %), melee_haste=8.443 ± 1.023

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 51.4 attack_power points (4.35 DPS) | yes | Chromite Barbute (8142, -0.47 DPS) [dungeon]; White Bandit Mask (10008, -0.80 DPS) [crafted]; Barbaric Iron Helm (7915, -1.44 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.69 DPS) | yes | Scout's Medallion (19536, -0.00 DPS) [rep]; Ethereal Talisman (4430, -0.23 DPS) [quest]; Kaleidoscope Chain (13084, -0.40 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.0 attack_power points (2.71 DPS) | yes | Forest Tracker Epaulets (2278, -0.17 DPS) [world_drop]; Flintrock Shoulders (7755, -0.32 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.77 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 26.2 attack_power points (2.21 DPS) | yes | First Sergeant's Cloak (16340, -0.28 DPS) [pvp]; Hawkeye's Cloak (14593, -0.63 DPS) [world_drop]; Parachute Cloak (10518, -0.98 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 50.5 attack_power points (4.27 DPS) | yes | Kolkar Marauder Chain (6773, -0.83 DPS) [quest]; Carapace of Tuten'kash (10775, -1.35 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -1.66 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.3 attack_power points (1.80 DPS) | yes | Branded Leather Bracers (19508, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 45.4 attack_power points (3.84 DPS) | yes | Scarlet Gauntlets (10331, -0.27 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.61 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.77 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.3 attack_power points (2.91 DPS) | yes | Defiler's Chain Girdle (20153, -0.08 DPS) [rep]; Officer's Belt (250556, -0.29 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.37 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.55 DPS) | yes | Symbolic Legplates (14829, -0.09 DPS) [world_drop]; Triprunner Dungarees (9624, -0.28 DPS) [quest]; Basilisk Hide Pants (1718, -0.33 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 38.7 attack_power points (3.28 DPS) | yes | Blackforge Greaves (6423, -0.22 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.34 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.40 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 30.5 attack_power points (2.58 DPS) | yes | Falcon's Hook (7552, -0.69 DPS) [dungeon]; Ring of the Underwood (2951, -0.71 DPS) [world_drop]; Thunderbrow Ring (13097, -0.77 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 24.3 attack_power points (2.06 DPS) | yes | Falcon's Hook (7552, -0.17 DPS) [dungeon]; Ring of the Underwood (2951, -0.18 DPS) [world_drop]; Thunderbrow Ring (13097, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (128.4 DPS) | yes | Skullance Shield (13081, -36.28 DPS) [world_drop]; Savage Boar's Guard (10767, -36.39 DPS) [dungeon]; Ardent Custodian (868, -42.82 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.18 DPS) [quest]; The Silencer (13138, -0.29 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.16 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 199.2. Weights run: 2.1s. Verify run: 2.4s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.574 ± 0.058, crit=1.884 ± 0.078 per rating point (14 rating = 1%, 26.382 per %), hit=1.955 ± 0.136 per rating point (10 rating = 1%, 19.550 per %), melee_haste=13.338 ± 1.221

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 81.9 attack_power points (7.92 DPS) | yes | Blood Guard's Plate Helm (220803, -0.97 DPS) [vendor]; Embrace of the Lycan (9479, -3.28 DPS) [dungeon]; Raging Berserker's Helm (7719, -5.21 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.5 attack_power points (2.94 DPS) | yes | Woven Ivy Necklace (19159, -0.42 DPS) [quest]; Zealous Shadowshard Pendant (17772, -1.01 DPS) [quest]; Scout's Medallion (19535, -1.12 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 55.3 attack_power points (5.34 DPS) | yes | Blood Guard's Plate Pauldrons (220796, -0.67 DPS) [vendor]; Wyrmslayer Spaulders (13066, -1.81 DPS) [world_drop]; Officer's Pauldrons (250576, -3.84 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 34.0 attack_power points (3.29 DPS) | yes | Blisterbane Wrap (12552, -1.01 DPS) [dungeon]; Dark Phantom Cape (13122, -1.01 DPS) [world_drop]; Dark Hooded Cape (5257, -1.45 DPS, sim-verified) [world] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 56.4 attack_power points (5.45 DPS) | yes | Mixologist's Tunic (12793, -0.30 DPS) [dungeon]; Warbear Harness (15064, -0.58 DPS) [crafted]; Warforged Chestplate (11195, -0.81 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.6 attack_power points (3.05 DPS) | yes | Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Officer's Wristguards (250581, +0.00 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 59.7 attack_power points (5.77 DPS) | yes | Raider Gloves (272100, -1.15 DPS, sim-verified) [vendor]; Gloves of Holy Might (867, -1.29 DPS) [world_drop]; Officer's Gloves (250551, -1.31 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 46.9 attack_power points (4.53 DPS) | yes | Defiler's Plate Girdle (20205, -0.05 DPS) [rep]; Defiler's Chain Girdle (20151, -0.05 DPS) [rep]; Defiler's Leather Girdle (20193, -0.05 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 66.1 attack_power points (6.39 DPS) | yes | Serpentskin Leggings (8262, -1.44 DPS) [world_drop]; Golem Shard Leggings (13074, -2.14 DPS) [world_drop]; Stormshroud Pants (15057, -2.19 DPS, sim-verified) [crafted] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 48.5 attack_power points (4.68 DPS) | yes | Prowler's Leather Boots (252468, -0.50 DPS) [crafted]; Skulker's Leather Boots (252469, -0.58 DPS) [crafted]; Officer's Sabatons (250561, -0.80 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 39.6 attack_power points (3.82 DPS) | yes | White Bone Band (11862, -1.50 DPS) [quest]; Ironspine's Eye (7686, -1.68 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.69 DPS) [quest] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 34.2 attack_power points (3.30 DPS) | yes | White Bone Band (11862, -0.98 DPS) [quest]; Ironspine's Eye (7686, -1.16 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.17 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (199.2 DPS) | yes | Frozen Heart of the Mountain (249469, -3.68 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (199.2 DPS) | yes | Frozen Heart of the Mountain (249469, -5.33 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (199.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hanzo Sword (8190, -2.28 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (52.66 DPS) | yes | Dawn's Edge (12774, +0.00 DPS) [crafted]; Claw of Celebras (17738, -6.07 DPS) [dungeon]; White Bone Shredder (11863, -8.79 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (199.2 DPS) | yes | Stinging Bow (10624, -0.20 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.20 DPS) [world_drop]; Dark Iron Rifle (16004, -2.95 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Stone Guard's Plate Armor; wrist: Deepfury Bracers; hands: Raider Gauntlets; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 304.1. Weights run: 2.1s. Verify run: 8.1s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.343 ± 0.121, crit=3.346 ± 0.173 per rating point (14 rating = 1%, 46.851 per %), hit=2.900 ± 0.318 per rating point (10 rating = 1%, 28.997 per %), melee_haste=18.914 ± 2.945

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 187.7 attack_power points (19.20 DPS) | yes | Outlaw's Collar (279253, -6.98 DPS) [crafted]; Champion's Plate Helm (227043, -7.15 DPS) [pvp]; Eye of Rend (12587, -10.85 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-verified (304.1 DPS) | yes | Medallion of the Dawn (22659, -0.20 DPS) [quest]; Amulet of the Darkmoon (19491, -0.85 DPS) [quest]; Rage of Mugamba (19577, -4.40 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 110.1 attack_power points (11.26 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Plate Shoulders (231534, -0.78 DPS) [pvp]; Truestrike Shoulders (12927, -2.88 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 57.0 attack_power points (5.83 DPS) | yes | Windshear Cape (20691, -0.60 DPS) [world]; Fel Cape (279269, -1.04 DPS) [crafted]; Cape of the Black Baron (13340, -1.67 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (304.1 DPS) | yes | Timbermaw Tunic (252484, -0.77 DPS) [crafted]; Savage Gladiator Chain (11726, -2.81 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -14.17 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (304.1 DPS) | yes | Forest Stalker's Bracers (19587, -0.96 DPS) [rep]; Vambraces of the Sadist (13400, -1.53 DPS) [dungeon]; Bracers of Undead Slaying (23090, -6.97 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (304.1 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gauntlets of Heroism (226861, -0.25 DPS) [quest]; Razor Gauntlets (18326, -5.86 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 100.9 attack_power points (10.32 DPS) | yes | Ferocity of the Timbermaw (227805, -0.92 DPS) [vendor]; Valiant's Waistguard (272403, -1.70 DPS) [vendor]; Belt of Preserved Heads (20216, -2.70 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (304.1 DPS) | yes | Sentinel's Plate Legguards (237825, -2.86 DPS) [vendor]; Titanic Leggings (22385, -4.08 DPS) [crafted]; Cloudkeeper Legplates (14554, -6.26 DPS, sim-verified) [world_drop] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 75.8 attack_power points (7.76 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Boots of Heroism (21995, -0.70 DPS) [quest]; Battleboots of Heroism (226857, -0.70 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (304.1 DPS) | yes | Tarnished Elven Ring (18500, -2.84 DPS) [dungeon]; Painweaver Band (13098, -2.97 DPS) [dungeon]; Naglering (11669, -9.22 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (304.1 DPS) | yes | Tarnished Elven Ring (18500, -0.72 DPS) [dungeon]; Painweaver Band (13098, -0.85 DPS) [dungeon]; Naglering (11669, -7.32 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (304.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (304.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (304.1 DPS) | yes | Felstriker (12590, +0.00 DPS) [dungeon]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (304.1 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -27.72 DPS) [dungeon]; Skullflame Shield (1168, -100.30 DPS, sim-verified) [world_drop] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (304.1 DPS) | yes | Satyr's Bow (18323, -1.11 DPS) [dungeon]; Blackcrow (12651, -1.40 DPS) [dungeon]; Dark Iron Rifle (16004, -4.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Raider Gloves; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 887.8. Weights run: 2.5s. Verify run: 8.6s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.708 ± 0.169, crit=3.571 ± 0.234 per rating point (14 rating = 1%, 49.990 per %), hit=4.439 ± 0.425 per rating point (10 rating = 1%, 44.386 per %), melee_haste=32.332 ± 4.248

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 228.4 attack_power points (50.03 DPS) | yes | Champion's Plate Helm (227043, -19.23 DPS) [pvp]; Fury Visor (20521, -20.68 DPS) [quest]; Mask of the Unforgiven (13404, -22.14 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (887.8 DPS) | yes | Mark of Fordring (15411, -1.98 DPS) [quest]; Medallion of the Dawn (22659, -2.41 DPS) [quest]; Rage of Mugamba (19577, -10.10 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 123.1 attack_power points (26.97 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -2.26 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.4 attack_power points (15.86 DPS) | yes | Windshear Cape (20691, -3.10 DPS) [world]; Blackveil Cape (11626, -4.66 DPS) [dungeon]; Cape of the Black Baron (13340, -10.86 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (887.8 DPS) | yes | Timbermaw Tunic (252484, -5.09 DPS) [crafted]; Savage Gladiator Chain (11726, -10.18 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -34.66 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (887.8 DPS) | yes | Forest Stalker's Bracers (19587, -4.10 DPS) [rep]; Bracers of Undead Slaying (23090, -19.75 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | sim-verified (887.8 DPS) | yes | Raider Gloves (272099, -3.47 DPS) [vendor]; Stormshroud Gloves (21278, -4.11 DPS) [crafted]; Razor Gauntlets (18326, -21.84 DPS, sim-verified) [dungeon] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 115.8 attack_power points (25.37 DPS) | yes | Radiant Girdle of the Dawn (227814, -1.41 DPS) [vendor]; Ferocity of the Timbermaw (227805, -2.75 DPS) [vendor]; Marksman's Girdle (22232, -3.19 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (887.8 DPS) | yes | Titanic Leggings (22385, -7.53 DPS) [crafted]; Sentinel's Plate Legguards (237825, -7.75 DPS) [vendor]; Cloudkeeper Legplates (14554, -15.80 DPS, sim-verified) [world_drop] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 98.6 attack_power points (21.59 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -2.19 DPS) [dungeon]; Boots of Heroism (21995, -2.23 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (887.8 DPS) | yes | Tarnished Elven Ring (18500, -5.56 DPS) [dungeon]; Cutthroat's Signet (272408, -6.15 DPS) [vendor]; Naglering (11669, -24.35 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (887.8 DPS) | yes | Tarnished Elven Ring (18500, -1.78 DPS) [dungeon]; Cutthroat's Signet (272408, -2.37 DPS) [vendor]; Naglering (11669, -19.92 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (887.8 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (887.8 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -9.65 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (887.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -9.82 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (887.8 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -60.65 DPS) [dungeon]; Skullflame Shield (1168, -242.60 DPS, sim-verified) [world_drop] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (887.8 DPS) | yes | The Purifier (22656, -0.55 DPS) [quest]; Blackcrow (12651, -0.97 DPS) [dungeon]; Dark Iron Rifle (16004, -13.08 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

