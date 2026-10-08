# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 34.9. Weights run: 2.7s. Verify run: 3.7s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.576 ± 0.017, crit=0.824 ± 0.024 per rating point (14 rating = 1%, 11.530 per %), hit=1.087 ± 0.048 per rating point (10 rating = 1%, 10.869 per %), melee_haste=7.278 ± 0.474

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.07 DPS) | yes | Defender's Leather Hood (252447, -0.24 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.76 DPS) [crafted]; Brawler's Leather Hood (252504, -0.83 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 3.5 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.06 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.32 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.43 DPS) | yes | Grave Shroud (279865, -0.05 DPS) [quest]; Catacomb Cloak (279899, -0.11 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (1.07 DPS) | yes | Veteran's Chain Shirt (250488, -0.22 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.23 DPS) [crafted]; Brawler's Leather Armor (252490, -0.32 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.54 DPS) | yes | Cryptwalker Bracers (280095, -0.11 DPS) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.23 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (34.9 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Polar Gauntlets (7606, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, -2.43 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.97 DPS) | yes | Ruffian Belt (5975, -0.32 DPS) [world]; Brawler's Leather Belt (252428, -0.41 DPS) [crafted]; Cobrahn's Grasp (6460, -2.78 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (34.9 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Chausses of Westfall (6087, -2.33 DPS, sim-verified) [quest] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (34.9 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -2.08 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.3 attack_power points (0.55 DPS) | yes | Signet of the Zhevra (285330, -0.37 DPS) [world]; The 1 Ring (8350, -0.41 DPS) [world]; Ring of the Moon (12052, -0.45 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.43 DPS) | yes | Signet of the Zhevra (285330, -0.26 DPS, sim-verified) [world]; The 1 Ring (8350, -0.29 DPS) [world]; Ring of the Moon (12052, -0.32 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.35 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.7 attack_power points (12.75 DPS) | yes | Diamond Hammer (2194, -0.31 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -12.10 DPS) [world_drop]; Furen's Favor (6970, -12.43 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.21 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Lil Timmy's Peashooter (13136, -0.09 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 62.1. Weights run: 2.9s. Verify run: 2.2s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.802 ± 0.024, crit=1.146 ± 0.034 per rating point (14 rating = 1%, 16.041 per %), hit=1.313 ± 0.068 per rating point (10 rating = 1%, 13.129 per %), melee_haste=10.156 ± 0.752

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.47 DPS) | yes | Barbaric Iron Helm (7915, -0.04 DPS) [crafted]; Veteran's Chain Helm (250498, -0.11 DPS) [crafted]; Defender's Leather Helm (252455, -0.11 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.79 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.34 DPS) [world_drop]; Sentinel's Medallion (19541, -0.43 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 18.8 attack_power points (1.07 DPS) | yes | Barbaric Iron Shoulders (7913, -0.11 DPS) [crafted]; Barbaric Shoulders (5964, -0.27 DPS) [crafted]; Golden Scale Shoulders (3841, -0.27 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 11.6 attack_power points (0.66 DPS) | yes | Sergeant Major's Cape (16315, -0.02 DPS) [pvp]; Wolfmaster Cape (6314, -0.09 DPS) [dungeon]; Slayer's Cape (14752, -0.20 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.70 DPS) | yes | Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.29 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.34 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.91 DPS) | yes | Yorgen Bracers (13012, -0.09 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop]; Bands of Serra'kis (6902, -0.23 DPS) [dungeon] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.25 DPS) | yes | The Frozen Clutch (23170, -0.11 DPS) [dungeon]; Mail Combat Gauntlets (4075, -0.16 DPS) [world_drop]; Insignia Gloves (6408, -0.18 DPS) [world_drop] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 24.8 attack_power points (1.41 DPS) | yes | Highlander's Chain Girdle (20090, -0.05 DPS) [rep]; Highlander's Leather Girdle (20117, -0.05 DPS) [rep]; Girdle of Golem Strength (9405, -0.05 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 27.6 attack_power points (1.56 DPS) | yes | Ferine Leggings (6690, -0.09 DPS) [dungeon]; Veteran's Chain Leggings (250493, -0.32 DPS) [crafted]; Golden Scale Leggings (3843, -0.32 DPS) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.6 attack_power points (1.11 DPS) | yes | Brawler's Leather Boots (252439, -0.32 DPS) [crafted]; Hard Gold Boots (250534, -0.32 DPS) [crafted]; Alacritous Treads (277234, -0.37 DPS, sim-verified) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 18.4 attack_power points (1.04 DPS) | yes | Ironspine's Eye (7686, -0.18 DPS) [dungeon]; Tiger Band (6749, -0.36 DPS) [quest]; Silverlaine's Family Seal (6321, -0.48 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.8 attack_power points (0.95 DPS) | yes | Ironspine's Eye (7686, -0.09 DPS) [dungeon]; Tiger Band (6749, -0.27 DPS) [quest]; Silverlaine's Family Seal (6321, -0.39 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (19.57 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (18.97 DPS) | yes | Royal Diplomatic Scepter (9457, -1.20 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -6.24 DPS) [quest]; Slayer's Shield (15892, -18.15 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.51 DPS) | yes | Double-barreled Shotgun (2098, -0.03 DPS) [world_drop]; Long Battle Bow (15284, -0.17 DPS) [world_drop]; Silver Star (3463, -0.28 DPS) [quest] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 118.9. Weights run: 3.1s. Verify run: 2.6s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.151 ± 0.049, crit=1.644 ± 0.070 per rating point (14 rating = 1%, 23.013 per %), hit=1.696 ± 0.117 per rating point (10 rating = 1%, 16.956 per %), melee_haste=9.360 ± 0.956

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 49.0 attack_power points (4.03 DPS) | yes | Chromite Barbute (8142, -0.78 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.18 DPS) [crafted]; Barbaric Iron Helm (7915, -1.70 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.64 DPS) | yes | Ghostshard Talisman (7731, -0.49 DPS) [dungeon]; Sentinel's Medallion (19540, -0.60 DPS) [rep]; Kaleidoscope Chain (13084, -0.61 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.7 attack_power points (2.03 DPS) | yes | Forest Tracker Epaulets (2278, -0.16 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.22 DPS) [crafted]; Flintrock Shoulders (7755, -0.26 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.5 attack_power points (1.60 DPS) | yes | Sergeant Major's Cape (16336, -0.05 DPS) [pvp]; Hawkeye's Cloak (14593, -0.45 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.54 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.9 attack_power points (3.11 DPS) | yes | Kolkar Marauder Chain (6773, -0.15 DPS) [quest]; Avenger's Armor (1488, -0.65 DPS) [dungeon]; Jouster's Chestplate (8157, -0.65 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.64 DPS) | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Pugilist Bracers (4438, -0.33 DPS) [dungeon]; Yorgen Bracers (13012, -0.37 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 43.0 attack_power points (3.54 DPS) | yes | Scarlet Gauntlets (10331, -0.62 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.91 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.91 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 31.0 attack_power points (2.55 DPS) | yes | Highlander's Leather Girdle (20116, -0.08 DPS) [rep]; Boar Champion's Belt (10768, -0.08 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.08 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.45 DPS) | yes | Firemane Leggings (13129, -0.33 DPS) [world_drop]; Symbolic Legplates (14829, -0.42 DPS) [world_drop]; Orcish War Leggings (7929, -0.66 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 34.1 attack_power points (2.80 DPS) | yes | Prowler's Leather Shoes (252465, -0.33 DPS) [crafted]; Blackforge Greaves (6423, -0.37 DPS) [dungeon]; Obsidian Greaves (13068, -0.59 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 25.2 attack_power points (2.07 DPS) | yes | Assault Band (13095, -0.43 DPS) [world_drop]; Thunderbrow Ring (13097, -0.47 DPS) [world_drop]; Ironspine's Eye (7686, -0.56 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.64 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest]; Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (118.9 DPS) | yes | Shoni's Disarming Tool (9608, -18.20 DPS) [quest]; Savage Boar's Guard (10767, -34.87 DPS) [dungeon]; Ardent Custodian (868, -38.62 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -0.12 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.23 DPS) [quest]; Bow of Searing Arrows (2825, -1.04 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 185.6. Weights run: 3.1s. Verify run: 3.4s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.309 ± 0.055, crit=1.870 ± 0.079 per rating point (14 rating = 1%, 26.180 per %), hit=2.144 ± 0.136 per rating point (10 rating = 1%, 21.436 per %), melee_haste=10.989 ± 1.155

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 83.6 attack_power points (7.96 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.12 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -2.99 DPS) [dungeon]; Embrace of the Lycan (9479, -3.39 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.0 attack_power points (2.57 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.67 DPS) [quest]; Sentinel's Medallion (19539, -1.08 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 54.5 attack_power points (5.19 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, -0.60 DPS) [vendor]; Officer's Pauldrons (250576, -1.72 DPS) [crafted]; Wyrmslayer Spaulders (13066, -1.91 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 30.3 attack_power points (2.89 DPS) | yes | Dark Hooded Cape (5257, -0.88 DPS) [world]; Sergeant Major's Cape (16336, -1.00 DPS) [pvp]; Dark Phantom Cape (13122, -1.02 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 56.2 attack_power points (5.35 DPS) | yes | Mixologist's Tunic (12793, -0.55 DPS) [dungeon]; Warforged Chestplate (11195, -0.78 DPS) [quest]; Warbear Harness (15064, -1.01 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.9 attack_power points (2.84 DPS) | yes | Runed Golem Shackles (12550, -0.18 DPS) [dungeon]; Bracers of the Stone Princess (17714, -0.18 DPS) [dungeon]; Arena Bands (18711, -0.18 DPS) [world] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 57.1 attack_power points (5.44 DPS) | yes | Raider Gloves (272100, -0.79 DPS) [vendor]; Gloves of Holy Might (867, -1.04 DPS) [world_drop]; Officer's Gloves (250551, -1.27 DPS) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 48.2 attack_power points (4.59 DPS) | yes | Highlander's Plate Girdle (20124, -0.19 DPS) [rep]; Highlander's Chain Girdle (20088, -0.19 DPS) [rep]; Highlander's Leather Girdle (20115, -0.19 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 63.3 attack_power points (6.03 DPS) | yes | Stormshroud Pants (15057, -1.22 DPS, sim-verified) [crafted]; Gryphon Rider's Leggings (9652, -1.30 DPS) [quest]; Centurion Legplates (10740, -1.30 DPS) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 45.0 attack_power points (4.29 DPS) | yes | Prowler's Leather Boots (252468, -0.44 DPS) [crafted]; Skulker's Leather Boots (252469, -0.57 DPS) [crafted]; Officer's Sabatons (250561, -0.69 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.4 attack_power points (3.95 DPS) | yes | Mark of Kern (2262, -2.04 DPS) [dungeon]; Assault Band (13095, -2.04 DPS) [world_drop]; Thunderbrow Ring (13097, -2.05 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 31.8 attack_power points (3.03 DPS) | yes | Mark of Kern (2262, -1.12 DPS) [dungeon]; Assault Band (13095, -1.12 DPS) [world_drop]; Thunderbrow Ring (13097, -1.13 DPS) [world_drop] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (185.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (185.6 DPS) | yes | Molten Heart of the Mountain (249470, -1.98 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (185.6 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hanzo Sword (8190, -1.46 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (51.92 DPS) | yes | Dawn's Edge (12774, +0.00 DPS) [crafted]; Claw of Celebras (17738, -5.99 DPS) [dungeon]; Shoni's Disarming Tool (9608, -30.51 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (185.6 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.16 DPS) [world_drop]; Dark Iron Rifle (16004, -2.25 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 300.6. Weights run: 3.1s. Verify run: 11.0s. 1943 eligible items had no known source.

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

Set DPS (verified): 876.2. Weights run: 3.5s. Verify run: 11.8s. 1943 eligible items had no known source.

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

Set DPS (verified): 35.4. Weights run: 2.7s. Verify run: 3.7s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.576 ± 0.017, crit=0.824 ± 0.024 per rating point (14 rating = 1%, 11.530 per %), hit=1.087 ± 0.048 per rating point (10 rating = 1%, 10.869 per %), melee_haste=7.278 ± 0.474

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.07 DPS) | yes | Defender's Leather Hood (252447, -0.24 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.76 DPS) [crafted]; Brawler's Leather Hood (252504, -0.83 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 3.5 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.06 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.32 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.43 DPS) | yes | Grave Shroud (279865, -0.05 DPS) [quest]; Catacomb Cloak (279899, -0.11 DPS) [quest]; Subterranean Cape (14149, -0.11 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (1.07 DPS) | yes | Defender's Leather Armor (252434, -0.23 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.24 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.32 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.54 DPS) | yes | Bristlebark Bindings (14569, -0.23 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.24 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.32 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (35.4 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.01 DPS) [quest]; Thorbia's Gauntlets (12994, -2.40 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.97 DPS) | yes | Ruffian Belt (5975, -0.32 DPS) [world]; Brawler's Leather Belt (252428, -0.41 DPS) [crafted]; Cobrahn's Grasp (6460, -2.80 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (35.4 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.32 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (35.4 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -2.08 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.3 attack_power points (0.55 DPS) | yes | Loop of Sacrifice (281673, -0.23 DPS) [quest]; Signet of the Zhevra (285330, -0.37 DPS) [world]; The 1 Ring (8350, -0.41 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.43 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Signet of the Zhevra (285330, -0.24 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.35 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.7 attack_power points (12.75 DPS) | yes | Diamond Hammer (2194, -0.36 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -12.10 DPS) [world_drop]; Ruga's Bulwark (7120, -12.43 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.21 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Lil Timmy's Peashooter (13136, -0.09 DPS) [world_drop]; Heavy Shortbow (3036, -0.11 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 62.8. Weights run: 2.9s. Verify run: 2.1s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.802 ± 0.024, crit=1.146 ± 0.034 per rating point (14 rating = 1%, 16.041 per %), hit=1.313 ± 0.068 per rating point (10 rating = 1%, 13.129 per %), melee_haste=10.156 ± 0.752

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.47 DPS) | yes | Barbaric Iron Helm (7915, -0.04 DPS) [crafted]; Veteran's Chain Helm (250498, -0.11 DPS) [crafted]; Defender's Leather Helm (252455, -0.11 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.79 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.34 DPS) [world_drop]; Scout's Medallion (19537, -0.43 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 18.8 attack_power points (1.07 DPS) | yes | Barbaric Iron Shoulders (7913, -0.11 DPS) [crafted]; Barbaric Shoulders (5964, -0.27 DPS) [crafted]; Golden Scale Shoulders (3841, -0.27 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 11.6 attack_power points (0.66 DPS) | yes | Wolfmaster Cape (6314, -0.09 DPS) [dungeon]; Wildhunter Cloak (16658, -0.09 DPS) [quest]; Slayer's Cape (14752, -0.20 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.70 DPS) | yes | Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.29 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.34 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.91 DPS) | yes | Yorgen Bracers (13012, -0.09 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop]; Bands of Serra'kis (6902, -0.23 DPS) [dungeon] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.25 DPS) | yes | The Frozen Clutch (23170, -0.11 DPS) [dungeon]; Warsong Gauntlets (16978, -0.11 DPS) [quest]; Mail Combat Gauntlets (4075, -0.16 DPS) [world_drop] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 24.8 attack_power points (1.41 DPS) | yes | Defiler's Chain Girdle (20152, -0.05 DPS) [rep]; Defiler's Leather Girdle (20191, -0.05 DPS) [rep]; Girdle of Golem Strength (9405, -0.05 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 27.6 attack_power points (1.56 DPS) | yes | Ferine Leggings (6690, -0.09 DPS) [dungeon]; Veteran's Chain Leggings (250493, -0.32 DPS) [crafted]; Golden Scale Leggings (3843, -0.32 DPS) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.6 attack_power points (1.11 DPS) | yes | Brawler's Leather Boots (252439, -0.31 DPS, sim-verified) [crafted]; Hard Gold Boots (250534, -0.32 DPS) [crafted]; Veteran's Boots (250503, -0.36 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 18.4 attack_power points (1.04 DPS) | yes | Ironspine's Eye (7686, -0.18 DPS) [dungeon]; Tiger Band (6749, -0.36 DPS) [quest]; Band of the Fist (17694, -0.41 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.8 attack_power points (0.95 DPS) | yes | Ironspine's Eye (7686, -0.09 DPS) [dungeon]; Tiger Band (6749, -0.27 DPS) [quest]; Band of the Fist (17694, -0.32 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (19.57 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (18.97 DPS) | yes | Royal Diplomatic Scepter (9457, -0.84 DPS, sim-verified) [dungeon]; Slayer's Shield (15892, -18.15 DPS) [world_drop]; Shield of Thorsen (13079, -18.17 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.51 DPS) | yes | Double-barreled Shotgun (2098, -0.03 DPS) [world_drop]; Long Battle Bow (15284, -0.17 DPS) [world_drop]; Silver Star (3463, -0.28 DPS) [quest] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 120.5. Weights run: 3.1s. Verify run: 2.6s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.151 ± 0.049, crit=1.644 ± 0.070 per rating point (14 rating = 1%, 23.013 per %), hit=1.696 ± 0.117 per rating point (10 rating = 1%, 16.956 per %), melee_haste=9.360 ± 0.956

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 49.0 attack_power points (4.03 DPS) | yes | Chromite Barbute (8142, -0.63 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.18 DPS) [crafted]; Barbaric Iron Helm (7915, -1.70 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.64 DPS) | yes | Ethereal Talisman (4430, -0.44 DPS) [quest]; Ghostshard Talisman (7731, -0.49 DPS) [dungeon]; Scout's Medallion (19536, -0.60 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.7 attack_power points (2.03 DPS) | yes | Forest Tracker Epaulets (2278, -0.16 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.22 DPS) [crafted]; Flintrock Shoulders (7755, -0.26 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.5 attack_power points (1.60 DPS) | yes | First Sergeant's Cloak (16340, -0.05 DPS) [pvp]; Hawkeye's Cloak (14593, -0.45 DPS) [world_drop]; Wildhunter Cloak (16658, -0.78 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.9 attack_power points (3.11 DPS) | yes | Kolkar Marauder Chain (6773, -0.15 DPS) [quest]; Avenger's Armor (1488, -0.65 DPS) [dungeon]; Jouster's Chestplate (8157, -0.65 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.64 DPS) | yes | Ravager's Armguards (14770, +0.00 DPS) [world_drop]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 43.0 attack_power points (3.54 DPS) | yes | Scarlet Gauntlets (10331, -0.55 DPS, sim-verified) [dungeon]; Gauntlets of Divinity (7724, -0.91 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.91 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 31.0 attack_power points (2.55 DPS) | yes | Defiler's Leather Girdle (20192, -0.08 DPS) [rep]; Boar Champion's Belt (10768, -0.08 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.08 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.45 DPS) | yes | Firemane Leggings (13129, -0.33 DPS) [world_drop]; Symbolic Legplates (14829, -0.42 DPS) [world_drop]; Orcish War Leggings (7929, -0.66 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 34.1 attack_power points (2.80 DPS) | yes | Prowler's Leather Shoes (252465, -0.33 DPS) [crafted]; Blackforge Greaves (6423, -0.37 DPS) [dungeon]; Obsidian Greaves (13068, -0.59 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 25.2 attack_power points (2.07 DPS) | yes | Assault Band (13095, -0.43 DPS) [world_drop]; Thunderbrow Ring (13097, -0.47 DPS) [world_drop]; Ironspine's Eye (7686, -0.56 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.64 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (120.5 DPS) | yes | Savage Boar's Guard (10767, -34.87 DPS) [dungeon]; Skullance Shield (13081, -34.99 DPS) [world_drop]; Ardent Custodian (868, -39.50 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -0.12 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.23 DPS) [quest]; Bow of Searing Arrows (2825, -1.03 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 192.7. Weights run: 3.1s. Verify run: 3.5s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.309 ± 0.055, crit=1.870 ± 0.079 per rating point (14 rating = 1%, 26.180 per %), hit=2.144 ± 0.136 per rating point (10 rating = 1%, 21.436 per %), melee_haste=10.989 ± 1.155

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 83.6 attack_power points (7.96 DPS) | yes | Blood Guard's Plate Helm (220803, -0.95 DPS) [vendor]; Embrace of the Lycan (9479, -3.39 DPS) [dungeon]; Raging Berserker's Helm (7719, -4.27 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.0 attack_power points (2.57 DPS) | yes | Woven Ivy Necklace (19159, -0.31 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.67 DPS) [quest]; Scout's Medallion (19535, -1.08 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 54.5 attack_power points (5.19 DPS) | yes | Blood Guard's Plate Pauldrons (220796, -0.60 DPS) [vendor]; Wyrmslayer Spaulders (13066, -1.91 DPS) [world_drop]; Officer's Pauldrons (250576, -3.16 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 30.3 attack_power points (2.89 DPS) | yes | First Sergeant's Cloak (16340, -1.00 DPS) [pvp]; Dark Phantom Cape (13122, -1.02 DPS) [world_drop]; Dark Hooded Cape (5257, -1.28 DPS, sim-verified) [world] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 56.2 attack_power points (5.35 DPS) | yes | Warforged Chestplate (11195, -0.78 DPS) [quest]; Warbear Harness (15064, -1.01 DPS) [crafted]; Mixologist's Tunic (12793, -1.45 DPS, sim-verified) [dungeon] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.9 attack_power points (2.84 DPS) | yes | Runed Golem Shackles (12550, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 57.1 attack_power points (5.44 DPS) | yes | Gloves of Holy Might (867, -1.04 DPS) [world_drop]; Raider Gloves (272100, -1.07 DPS, sim-verified) [vendor]; Officer's Gloves (250551, -1.27 DPS) [crafted] |
| waist | Defiler's Plate Girdle (20205) | The Defilers [rep] | 46.2 attack_power points (4.40 DPS) | yes | Defiler's Chain Girdle (20151, -0.00 DPS) [rep]; Defiler's Leather Girdle (20193, -0.00 DPS) [rep]; Girdle of Beastial Fury (11686, -0.02 DPS) [dungeon] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 63.3 attack_power points (6.03 DPS) | yes | Stormshroud Pants (15057, -1.18 DPS, sim-verified) [crafted]; Serpentskin Leggings (8262, -1.56 DPS) [world_drop]; Golem Shard Leggings (13074, -1.84 DPS) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 45.0 attack_power points (4.29 DPS) | yes | Prowler's Leather Boots (252468, -0.44 DPS) [crafted]; Skulker's Leather Boots (252469, -0.57 DPS) [crafted]; Officer's Sabatons (250561, -0.69 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.4 attack_power points (3.95 DPS) | yes | White Bone Band (11862, -1.66 DPS) [quest]; Mark of Kern (2262, -2.04 DPS) [dungeon]; Assault Band (13095, -2.04 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 31.8 attack_power points (3.03 DPS) | yes | White Bone Band (11862, -0.74 DPS) [quest]; Mark of Kern (2262, -1.12 DPS) [dungeon]; Assault Band (13095, -1.12 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (192.7 DPS) | yes | Frozen Heart of the Mountain (249469, -3.59 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (192.7 DPS) | yes | Frozen Heart of the Mountain (249469, -6.72 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (192.7 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hookfang Shanker (11635, -2.98 DPS, sim-verified) [dungeon] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (51.92 DPS) | yes | Dawn's Edge (12774, +0.00 DPS) [crafted]; Claw of Celebras (17738, -5.99 DPS) [dungeon]; White Bone Shredder (11863, -8.85 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (192.7 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.16 DPS) [world_drop]; Dark Iron Rifle (16004, -2.26 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Defiler's Plate Girdle; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 304.1. Weights run: 3.1s. Verify run: 11.1s. 1913 eligible items had no known source.

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

Set DPS (verified): 887.8. Weights run: 3.5s. Verify run: 12.0s. 1913 eligible items had no known source.

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

