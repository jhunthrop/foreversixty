# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 33.3. Weights run: 2.6s. Verify run: 1.5s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.165 ± 0.004, crit=0.236 ± 0.006 per rating point (14 rating = 1%, 3.305 per %), hit=0.068 ± 0.002 per rating point (10 rating = 1%, 0.685 per %), melee_haste=2.484 ± 0.100

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.77 DPS) | yes | Defender's Leather Hood (252447, -0.23 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.71 DPS) [crafted]; Brawler's Leather Hood (252504, -0.72 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 1.0 attack_power points (0.04 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.23 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.31 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.08 DPS) [quest]; Dark Leather Cloak (2316, -0.14 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.77 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.22 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.23 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.39 DPS) | yes | Cryptwalker Bracers (280095, -0.08 DPS) [quest]; Bravo's Armbands (270015, -0.21 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.62 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Polar Gauntlets (7606, -0.15 DPS) [quest]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.70 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Hulking Belt (14746, -0.31 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.85 DPS) | yes | Veteran's Chain Leggings (250493, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.8 attack_power points (0.42 DPS) | yes | Veteran's Boots (250503, -0.01 DPS) [crafted]; Guard's Boots (250504, -0.03 DPS) [crafted]; Defender's Leather Boots (252441, -0.03 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.7 attack_power points (0.34 DPS) | yes | The 1 Ring (8350, -0.25 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop]; Signet of the Zhevra (285330, -0.30 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.31 DPS) | yes | Ring of the Moon (12052, -0.23 DPS) [world_drop]; Signet of the Zhevra (285330, -0.27 DPS) [world]; The 1 Ring (8350, -0.32 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.63 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.9 attack_power points (9.16 DPS) | yes | Diamond Hammer (2194, +0.00 DPS) [world_drop]; Redbeard Crest (12997, -8.70 DPS) [world_drop]; Furen's Favor (6970, -8.93 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.15 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Daryl's Hunting Rifle (2904, -0.08 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 61.2. Weights run: 2.9s. Verify run: 1.6s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.388 ± 0.015, crit=0.554 ± 0.021 per rating point (14 rating = 1%, 7.751 per %), hit=0.141 ± 0.004 per rating point (10 rating = 1%, 1.409 per %), melee_haste=5.108 ± 0.469

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.97 DPS) | yes | Veteran's Chain Helm (250498, -0.07 DPS) [crafted]; Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Crusader's Chain Helm (250502, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.52 DPS) | yes | Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; River Pride Choker (13087, -0.22 DPS) [world_drop]; Sentinel's Medallion (19541, -0.41 DPS) [rep] |
| shoulder | Barbaric Iron Shoulders (7913) | Blacksmithing [crafted] | 14.3 attack_power points (0.54 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS, sim-verified) [world_drop]; Golden Scale Shoulders (3841, -0.01 DPS) [crafted]; Mail Combat Spaulders (6404, -0.01 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.37 DPS) | yes | Sergeant Major's Cape (16315, -0.02 DPS) [pvp]; Hawkeye's Cloak (14593, -0.05 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.12 DPS) | yes | Shining Silver Breastplate (2870, -0.07 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.22 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.60 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Bands of Serra'kis (6902, -0.15 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.21 DPS) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.82 DPS) | yes | The Frozen Clutch (23170, -0.07 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.15 DPS) [world]; Mail Combat Gauntlets (4075, -0.17 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (0.90 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.97 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.05 DPS) [crafted]; Golden Scale Leggings (3843, -0.15 DPS) [crafted]; Chausses of Westfall (6087, -0.15 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 16.7 attack_power points (0.63 DPS) | yes | Hard Gold Boots (250534, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.18 DPS) [quest]; Glimmering Mail Greaves (4073, -0.18 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.2 attack_power points (0.64 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.27 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.3 attack_power points (0.54 DPS) | yes | Tiger Band (6749, -0.09 DPS) [quest]; Ironspine's Eye (7686, -0.11 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (12.94 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (12.55 DPS) | yes | Royal Diplomatic Scepter (9457, -0.59 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -4.13 DPS) [quest]; Shield of Thorsen (13079, -12.02 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.34 DPS) | yes | Double-barreled Shotgun (2098, -0.07 DPS) [world_drop]; Long Battle Bow (15284, -0.11 DPS) [world_drop]; Fine Longbow (11304, -0.19 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Barbaric Iron Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-05153105022011500-000000000000000000)

Set DPS (verified): 112.8. Weights run: 3.1s. Verify run: 1.6s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.086 ± 0.046, crit=1.552 ± 0.066 per rating point (14 rating = 1%, 21.726 per %), hit=0.276 ± 0.009 per rating point (10 rating = 1%, 2.756 per %), melee_haste=9.252 ± 0.940

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.7 attack_power points (1.94 DPS) | yes | White Bandit Mask (10008, -0.56 DPS) [crafted]; Chromite Barbute (8142, -0.79 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -0.80 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.81 DPS) | yes | Ghostshard Talisman (7731, -0.24 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.31 DPS) [world_drop]; Sentinel's Medallion (19540, -0.33 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.9 attack_power points (0.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.08 DPS) [world_drop]; Flintrock Shoulders (7755, -0.13 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.56 DPS, sim-verified) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 18.9 attack_power points (0.77 DPS) | yes | Sergeant Major's Cape (16336, -0.01 DPS) [pvp]; Hawkeye's Cloak (14593, -0.21 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.26 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 36.6 attack_power points (1.49 DPS) | yes | Avenger's Armor (1488, -0.27 DPS) [dungeon]; Jouster's Chestplate (8157, -0.27 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.59 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.81 DPS) | yes | Ravager's Armguards (14770, -0.07 DPS) [world_drop]; Pugilist Bracers (4438, -0.16 DPS) [dungeon]; Yorgen Bracers (13012, -0.19 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.7 attack_power points (1.69 DPS) | yes | Truesilver Gauntlets (7938, -0.39 DPS) [crafted]; Gauntlets of Divinity (7724, -0.39 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.64 DPS, sim-verified) [dungeon] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 30.0 attack_power points (1.22 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Highlander's Chain Girdle (20089, -0.01 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.70 DPS) | yes | Firemane Leggings (13129, -0.16 DPS) [world_drop]; Symbolic Legplates (14829, -0.22 DPS) [world_drop]; Orcish War Leggings (7929, -0.32 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 33.6 attack_power points (1.36 DPS) | yes | Prowler's Leather Shoes (252465, -0.16 DPS) [crafted]; Blackforge Greaves (6423, -0.19 DPS) [dungeon]; Obsidian Greaves (13068, -0.29 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 24.7 attack_power points (1.00 DPS) | yes | Assault Band (13095, -0.19 DPS) [world_drop]; Thunderbrow Ring (13097, -0.22 DPS) [world_drop]; Ironspine's Eye (7686, -0.28 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.81 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Ironspine's Eye (7686, -0.09 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest]; Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (112.8 DPS) | yes | Shoni's Disarming Tool (9608, -8.95 DPS) [quest]; Savage Boar's Guard (10767, -17.17 DPS) [dungeon]; Ardent Custodian (868, -38.82 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -0.05 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.12 DPS) [quest]; Bow of Searing Arrows (2825, -0.92 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 180.3. Weights run: 3.3s. Verify run: 2.2s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.805 ± 0.037, crit=1.151 ± 0.053 per rating point (14 rating = 1%, 16.107 per %), hit=0.250 ± 0.008 per rating point (10 rating = 1%, 2.499 per %), melee_haste=8.551 ± 0.697

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | sim-verified (180.3 DPS) | yes | Raging Berserker's Helm (7719, -0.17 DPS) [dungeon]; Fury Visor (20521, -0.50 DPS) [quest]; Embrace of the Lycan (9479, -3.13 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.5 attack_power points (1.42 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.03 DPS) [quest]; Ghostshard Talisman (7731, -0.45 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.64 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 38.1 attack_power points (2.65 DPS) | yes | Officer's Pauldrons (250576, -0.39 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.53 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.62 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 23.3 attack_power points (1.62 DPS) | yes | Sergeant Major's Cape (16336, -0.45 DPS) [pvp]; Dark Hooded Cape (5257, -0.50 DPS) [world]; Bloodlust Cape (14801, -1.21 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 48.0 attack_power points (3.34 DPS) | yes | Knight's Plate Hauberk (220794, -0.13 DPS) [vendor]; Mixologist's Tunic (12793, -0.22 DPS) [dungeon]; Valorous Chestguard (8274, -0.56 DPS) [world_drop] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 28.0 attack_power points (1.95 DPS) | yes | Bracers of the Stone Princess (17714, -0.00 DPS) [dungeon]; Arena Bands (18711, -0.00 DPS) [world]; Officer's Wristguards (250581, -0.08 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 52.1 attack_power points (3.62 DPS) | yes | Officer's Gloves (250551, -0.98 DPS, sim-verified) [crafted]; Raider Gloves (272100, -1.00 DPS) [vendor]; Gloves of Holy Might (867, -1.11 DPS) [world_drop] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.20 DPS) | yes | Highlander's Lamellar Girdle (20106, -0.55 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.58 DPS) [crafted]; Highlander's Plate Girdle (20124, -0.69 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 48.2 attack_power points (3.35 DPS) | yes | Gryphon Rider's Leggings (9652, -0.42 DPS) [quest]; Centurion Legplates (10740, -0.42 DPS) [quest]; Golem Shard Leggings (13074, -2.47 DPS, sim-verified) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 38.5 attack_power points (2.67 DPS) | yes | Prowler's Leather Boots (252468, -0.25 DPS) [crafted]; Officer's Sabatons (250561, -0.36 DPS) [crafted]; Skulker's Leather Boots (252469, -0.42 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 27.2 attack_power points (1.89 DPS) | yes | Mark of Kern (2262, -0.50 DPS) [dungeon]; Assault Band (13095, -0.50 DPS) [world_drop]; Thunderbrow Ring (13097, -0.61 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 22.5 attack_power points (1.56 DPS) | yes | Assault Band (13095, -0.17 DPS) [world_drop]; Thunderbrow Ring (13097, -0.28 DPS) [world_drop]; Mark of Kern (2262, -3.52 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -3.09 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hookfang Shanker (11635, -1.68 DPS, sim-verified) [dungeon] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (37.87 DPS) | yes | Doomforged Straightedge (12535, +0.00 DPS) [dungeon]; Claw of Celebras (17738, -4.37 DPS) [dungeon]; Shoni's Disarming Tool (9608, -22.26 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.14 DPS) [dungeon]; Dark Iron Rifle (16004, -2.09 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Plate Helm; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 288.0. Weights run: 3.4s. Verify run: 2.1s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.029 ± 0.051, crit=1.469 ± 0.072 per rating point (14 rating = 1%, 20.572 per %), hit=0.312 ± 0.010 per rating point (10 rating = 1%, 3.125 per %), melee_haste=9.776 ± 0.954

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 83.4 attack_power points (6.04 DPS) | yes | Field Marshal's Plate Helm (231538, -0.49 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -1.28 DPS) [vendor]; Crown of Heroism (226860, -9.86 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (288.0 DPS) | yes | Amulet of the Darkmoon (19491, -0.36 DPS) [quest]; Imperial Jewel (11933, -0.91 DPS) [dungeon]; Rage of Mugamba (19577, -3.17 DPS, sim-verified) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 53.5 attack_power points (3.88 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 39.1 attack_power points (2.84 DPS) | yes | Cape of the Black Baron (13340, -0.27 DPS) [dungeon]; Shroud of Domination (22337, -0.37 DPS) [dungeon]; Windshear Cape (20691, -0.56 DPS) [world] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (288.0 DPS) | yes | Obsidian Mail Tunic (22191, -0.62 DPS) [crafted]; Cadaverous Armor (14637, -1.51 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -15.23 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-verified (288.0 DPS) | yes | Forest Stalker's Bracers (19587, -0.34 DPS) [rep]; Marshal's Plate Bracers (16481, -0.36 DPS) [pvp]; Bracers of Undead Slaying (23090, -6.71 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (288.0 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.34 DPS) [pvp]; Gauntlets of Heroism (226861, -0.63 DPS) [quest]; Razor Gauntlets (18326, -6.98 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 74.6 attack_power points (5.41 DPS) | yes | Ferocity of the Timbermaw (227805, -0.37 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.77 DPS) [vendor]; Marshal's Plate Girdle (16482, -0.89 DPS) [pvp] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (288.0 DPS) | yes | Titanic Leggings (22385, -0.83 DPS) [crafted]; Marshal's Plate Legguards (231540, -1.01 DPS) [pvp]; Cloudkeeper Legplates (14554, -7.52 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 46.2 attack_power points (3.35 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Drudge Boots (21532, -0.20 DPS) [quest]; Boots of Heroism (21995, -0.22 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (288.0 DPS) | yes | Band of the Ogre King (18522, -0.85 DPS) [dungeon]; Myrmidon's Signet (2246, -0.91 DPS) [world_drop]; Naglering (11669, -9.47 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (288.0 DPS) | yes | Band of the Ogre King (18522, -0.53 DPS) [dungeon]; Myrmidon's Signet (2246, -0.59 DPS) [world_drop]; Naglering (11669, -5.20 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (288.0 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (288.0 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, +0.00 DPS, sim-verified) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (288.0 DPS) | yes | Felstriker (12590, +0.00 DPS) [dungeon]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (288.0 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -18.47 DPS) [dungeon]; Skullflame Shield (1168, -87.16 DPS, sim-verified) [world_drop] |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-verified (288.0 DPS) | yes | Riphook (12653, -0.09 DPS) [dungeon]; The Purifier (22656, -0.19 DPS) [quest]; Dark Iron Rifle (16004, -3.33 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Highlander's Plate Spaulders; back: Cloak of the Honor Guard; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Protector's Band; trinket2: Blackhand's Breadth; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Bloodseeker

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (troll, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 33.9. Weights run: 2.6s. Verify run: 1.5s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.165 ± 0.004, crit=0.236 ± 0.006 per rating point (14 rating = 1%, 3.305 per %), hit=0.068 ± 0.002 per rating point (10 rating = 1%, 0.685 per %), melee_haste=2.484 ± 0.100

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.77 DPS) | yes | Defender's Leather Hood (252447, -0.24 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.71 DPS) [crafted]; Brawler's Leather Hood (252504, -0.72 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 1.0 attack_power points (0.04 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.23 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.31 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.08 DPS) [quest]; Subterranean Cape (14149, -0.08 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.77 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.23 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.25 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.39 DPS) | yes | Bristlebark Bindings (14569, -0.21 DPS) [world_drop]; Runed Copper Bracers (2854, -0.23 DPS) [crafted]; Raptorcrest Bracers (270010, -0.24 DPS, sim-verified) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.62 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.21 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.70 DPS) | yes | Cobrahn's Grasp (6460, -0.18 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Hulking Belt (14746, -0.31 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.8 attack_power points (0.73 DPS) | yes | Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Totemic Leather Pants (252446, -0.03 DPS) [crafted]; Hulking Leggings (14748, -0.09 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.8 attack_power points (0.42 DPS) | yes | Veteran's Boots (250503, -0.01 DPS) [crafted]; Guard's Boots (250504, -0.03 DPS) [crafted]; Defender's Leather Boots (252441, -0.03 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.7 attack_power points (0.34 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; The 1 Ring (8350, -0.25 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.63 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.9 attack_power points (9.16 DPS) | yes | Diamond Hammer (2194, +0.00 DPS) [world_drop]; Redbeard Crest (12997, -8.70 DPS) [world_drop]; Ruga's Bulwark (7120, -8.93 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.15 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Heavy Shortbow (3036, -0.08 DPS) [world_drop]; Orcish Battle Bow (5346, -0.08 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 62.1. Weights run: 2.9s. Verify run: 1.6s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.388 ± 0.015, crit=0.554 ± 0.021 per rating point (14 rating = 1%, 7.751 per %), hit=0.141 ± 0.004 per rating point (10 rating = 1%, 1.409 per %), melee_haste=5.108 ± 0.469

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.97 DPS) | yes | Veteran's Chain Helm (250498, -0.07 DPS) [crafted]; Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Crusader's Chain Helm (250502, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.52 DPS) | yes | Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; River Pride Choker (13087, -0.22 DPS) [world_drop]; Scout's Medallion (19537, -0.41 DPS) [rep] |
| shoulder | Barbaric Iron Shoulders (7913) | Blacksmithing [crafted] | 14.3 attack_power points (0.54 DPS) | yes | Forest Tracker Epaulets (2278, -0.00 DPS) [world_drop]; Golden Scale Shoulders (3841, -0.01 DPS) [crafted]; Mail Combat Spaulders (6404, -0.01 DPS) [world_drop] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.37 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.05 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.12 DPS) | yes | Shining Silver Breastplate (2870, -0.07 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.22 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.60 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Bands of Serra'kis (6902, -0.15 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.21 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.82 DPS) | yes | The Frozen Clutch (23170, -0.07 DPS) [dungeon]; Warsong Gauntlets (16978, -0.07 DPS) [quest]; Bonefist Gauntlets (4465, -0.15 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (0.90 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.97 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.05 DPS) [crafted]; Golden Scale Leggings (3843, -0.15 DPS) [crafted]; Slayer's Pants (14757, -0.15 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 16.7 attack_power points (0.63 DPS) | yes | Glimmering Mail Greaves (4073, -0.18 DPS) [world_drop]; Slayer's Slippers (14756, -0.18 DPS) [world_drop]; Hard Gold Boots (250534, -0.44 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.2 attack_power points (0.64 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.27 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.3 attack_power points (0.54 DPS) | yes | Ironspine's Eye (7686, -0.11 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon]; Tiger Band (6749, -0.37 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (12.94 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (12.55 DPS) | yes | Royal Diplomatic Scepter (9457, -0.60 DPS, sim-verified) [dungeon]; Shield of Thorsen (13079, -12.02 DPS) [world_drop]; Slayer's Shield (15892, -12.05 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.34 DPS) | yes | Double-barreled Shotgun (2098, -0.07 DPS) [world_drop]; Long Battle Bow (15284, -0.11 DPS) [world_drop]; Fine Longbow (11304, -0.19 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Barbaric Iron Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-05153105022011500-000000000000000000)

Set DPS (verified): 114.0. Weights run: 3.1s. Verify run: 1.7s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.086 ± 0.046, crit=1.552 ± 0.066 per rating point (14 rating = 1%, 21.726 per %), hit=0.276 ± 0.009 per rating point (10 rating = 1%, 2.756 per %), melee_haste=9.252 ± 0.940

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.7 attack_power points (1.94 DPS) | yes | Chromite Barbute (8142, -0.50 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -0.56 DPS) [crafted]; Hard Gold Coif (250537, -0.80 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.81 DPS) | yes | Ethereal Talisman (4430, -0.23 DPS) [quest]; Ghostshard Talisman (7731, -0.24 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.31 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.9 attack_power points (0.97 DPS) | yes | Hard Gold Pauldrons (250539, -0.08 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.08 DPS) [world_drop]; Flintrock Shoulders (7755, -0.13 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 18.9 attack_power points (0.77 DPS) | yes | First Sergeant's Cloak (16340, -0.01 DPS) [pvp]; Hawkeye's Cloak (14593, -0.21 DPS) [world_drop]; Wildhunter Cloak (16658, -0.36 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 36.6 attack_power points (1.49 DPS) | yes | Kolkar Marauder Chain (6773, -0.04 DPS) [quest]; Avenger's Armor (1488, -0.27 DPS) [dungeon]; Jouster's Chestplate (8157, -0.27 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.81 DPS) | yes | Ravager's Armguards (14770, +0.00 DPS) [world_drop]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.7 attack_power points (1.69 DPS) | yes | Scarlet Gauntlets (10331, -0.28 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.39 DPS) [crafted]; Gauntlets of Divinity (7724, -0.39 DPS) [dungeon] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 30.0 attack_power points (1.22 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Defiler's Chain Girdle (20153, -0.01 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.70 DPS) | yes | Firemane Leggings (13129, -0.16 DPS) [world_drop]; Symbolic Legplates (14829, -0.22 DPS) [world_drop]; Orcish War Leggings (7929, -0.32 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 33.6 attack_power points (1.36 DPS) | yes | Prowler's Leather Shoes (252465, -0.16 DPS) [crafted]; Blackforge Greaves (6423, -0.19 DPS) [dungeon]; Obsidian Greaves (13068, -0.29 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 24.7 attack_power points (1.00 DPS) | yes | Assault Band (13095, -0.19 DPS) [world_drop]; Thunderbrow Ring (13097, -0.22 DPS) [world_drop]; Ironspine's Eye (7686, -0.28 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.81 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Ironspine's Eye (7686, -0.09 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (114.0 DPS) | yes | Savage Boar's Guard (10767, -17.17 DPS) [dungeon]; Skullance Shield (13081, -17.24 DPS) [world_drop]; Ardent Custodian (868, -39.11 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -0.05 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.12 DPS) [quest]; Bow of Searing Arrows (2825, -0.67 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 182.5. Weights run: 3.3s. Verify run: 2.3s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.805 ± 0.037, crit=1.151 ± 0.053 per rating point (14 rating = 1%, 16.107 per %), hit=0.250 ± 0.008 per rating point (10 rating = 1%, 2.499 per %), melee_haste=8.551 ± 0.697

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (3.34 DPS) | yes | Raging Berserker's Helm (7719, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Plate Helm (220803, -0.24 DPS) [vendor]; Fury Visor (20521, -0.74 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.5 attack_power points (1.42 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.03 DPS) [quest]; Woven Ivy Necklace (19159, -0.08 DPS) [quest]; Ghostshard Talisman (7731, -0.45 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 32.4 attack_power points (2.25 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.14 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.22 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 23.3 attack_power points (1.62 DPS) | yes | Bloodlust Cape (14801, -0.37 DPS) [world_drop]; First Sergeant's Cloak (16340, -0.45 DPS) [pvp]; Dark Hooded Cape (5257, -0.50 DPS) [world] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 48.0 attack_power points (3.34 DPS) | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Mixologist's Tunic (12793, -0.22 DPS) [dungeon]; Valorous Chestguard (8274, -0.56 DPS) [world_drop] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 28.0 attack_power points (1.95 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 52.1 attack_power points (3.62 DPS) | yes | Raider Gloves (272100, -1.00 DPS) [vendor]; Gloves of Holy Might (867, -1.11 DPS) [world_drop]; Officer's Gloves (250551, -1.35 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.20 DPS) | yes | Prowler's Leather Waistguard (252473, -0.58 DPS) [crafted]; Defiler's Plate Girdle (20205, -0.69 DPS) [rep]; Defiler's Chain Girdle (20151, -0.69 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 48.2 attack_power points (3.35 DPS) | yes | Scarlet Leggings (10330, -0.43 DPS) [dungeon]; Sunscale Legplates (14850, -0.54 DPS) [world_drop]; Golem Shard Leggings (13074, -2.69 DPS, sim-verified) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 38.5 attack_power points (2.67 DPS) | yes | Prowler's Leather Boots (252468, -0.25 DPS) [crafted]; Officer's Sabatons (250561, -0.36 DPS) [crafted]; Skulker's Leather Boots (252469, -0.42 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 27.2 attack_power points (1.89 DPS) | yes | Mark of Kern (2262, -0.50 DPS) [dungeon]; Assault Band (13095, -0.50 DPS) [world_drop]; White Bone Band (11862, -3.58 DPS, sim-verified) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (182.5 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; White Bone Band (11862, -2.68 DPS, sim-verified) [quest] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.89 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Hanzo Sword (8190, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (37.87 DPS) | yes | Doomforged Straightedge (12535, +0.00 DPS) [dungeon]; Claw of Celebras (17738, -4.37 DPS) [dungeon]; White Bone Shredder (11863, -6.70 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.14 DPS) [dungeon]; Dark Iron Rifle (16004, -2.05 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Rune of the Guard Captain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 292.1. Weights run: 3.4s. Verify run: 2.2s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.029 ± 0.051, crit=1.469 ± 0.072 per rating point (14 rating = 1%, 20.572 per %), hit=0.312 ± 0.010 per rating point (10 rating = 1%, 3.125 per %), melee_haste=9.776 ± 0.954

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 83.4 attack_power points (6.04 DPS) | yes | Warlord's Plate Headpiece (231535, -0.49 DPS) [pvp]; Champion's Plate Helm (227043, -1.28 DPS) [pvp]; Crown of Heroism (226860, -11.61 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (292.1 DPS) | yes | Amulet of the Darkmoon (19491, -0.36 DPS) [quest]; Imperial Jewel (11933, -0.91 DPS) [dungeon]; Rage of Mugamba (19577, -3.00 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 53.5 attack_power points (3.88 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Defiler's Leather Shoulders (20194, -0.36 DPS) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 39.1 attack_power points (2.84 DPS) | yes | Cape of the Black Baron (13340, -0.27 DPS) [dungeon]; Shroud of Domination (22337, -0.37 DPS) [dungeon]; Windshear Cape (20691, -0.56 DPS) [world] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (292.1 DPS) | yes | Obsidian Mail Tunic (22191, -0.62 DPS) [crafted]; Cadaverous Armor (14637, -1.51 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -17.62 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (292.1 DPS) | yes | Forest Stalker's Bracers (19587, -0.34 DPS) [rep]; General's Plate Armguards (16546, -0.36 DPS) [pvp]; Bracers of Undead Slaying (23090, -5.74 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (292.1 DPS) | yes | General's Plate Gauntlets (231532, -0.34 DPS) [pvp]; Gauntlets of Heroism (226861, -0.63 DPS) [quest]; Razor Gauntlets (18326, -6.52 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 74.6 attack_power points (5.41 DPS) | yes | Ferocity of the Timbermaw (227805, -0.37 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.77 DPS) [vendor]; General's Plate Girdle (16547, -0.89 DPS) [pvp] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (292.1 DPS) | yes | Titanic Leggings (22385, -0.83 DPS) [crafted]; General's Plate Leggings (231533, -1.01 DPS) [pvp]; Cloudkeeper Legplates (14554, -8.11 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 46.2 attack_power points (3.35 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Drudge Boots (21532, -0.20 DPS) [quest]; Boots of Heroism (21995, -0.22 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (292.1 DPS) | yes | Band of the Ogre King (18522, -0.85 DPS) [dungeon]; Myrmidon's Signet (2246, -0.91 DPS) [world_drop]; Naglering (11669, -11.61 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (292.1 DPS) | yes | Band of the Ogre King (18522, -0.53 DPS) [dungeon]; Myrmidon's Signet (2246, -0.59 DPS) [world_drop]; Naglering (11669, -4.70 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (292.1 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, +0.00 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (292.1 DPS) | yes | Diamond Flask (20130, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -1.61 DPS) [dungeon]; Hand of Justice (11815, -1.75 DPS) [dungeon] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (292.1 DPS) | yes | Felstriker (12590, +0.00 DPS, sim-verified) [dungeon]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (292.1 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -18.47 DPS) [dungeon]; Skullflame Shield (1168, -88.50 DPS, sim-verified) [world_drop] |
| ranged | Bloodseeker (19107) | The Legend of Korrak [quest] | sim-verified (292.1 DPS) | yes | Riphook (12653, -0.09 DPS) [dungeon]; The Purifier (22656, -0.19 DPS) [quest]; Dark Iron Rifle (16004, -2.90 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Defiler's Plate Spaulders; back: Deathguard's Cloak; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Blackhand's Breadth; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Bloodseeker

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

