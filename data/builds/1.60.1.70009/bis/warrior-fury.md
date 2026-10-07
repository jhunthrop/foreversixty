# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-350300000000000000-000000000000000000)

Set DPS (verified): 33.2. Weights run: 2.5s. Verify run: 1.5s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.026, strength=1.998 ± 0.032, agility=0.104 ± 0.011, crit=0.169 ± 0.004 per rating point (14 rating = 1%, 2.366 per %), hit=0.050 ± 0.001 per rating point (10 rating = 1%, 0.499 per %), melee_haste=1.728 ± 0.034

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.30 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.70 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.14 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.73 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.44 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.37 DPS) | yes | Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop]; Cryptwalker Bracers (280095, -0.22 DPS, sim-verified) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.59 DPS) | yes | Polar Gauntlets (7606, -0.15 DPS) [quest]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Gold-flecked Gloves (5195, -0.22 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.66 DPS) | yes | Ruffian Belt (5975, -0.22 DPS) [world]; Cobrahn's Grasp (6460, -0.23 DPS, sim-verified) [dungeon]; Hulking Belt (14746, -0.29 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.81 DPS) | yes | Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.21 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.39 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.31 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop]; Signet of the Zhevra (285330, -0.29 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Ring of the Moon (12052, -0.22 DPS) [world_drop]; Signet of the Zhevra (285330, -0.27 DPS) [world]; The 1 Ring (8350, -0.47 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.14 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (8.69 DPS) | yes | Diamond Hammer (2194, -0.39 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.25 DPS) [world_drop]; Furen's Favor (6970, -8.47 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.15 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 00000000000000000-350511005010000000-000000000000000000)

Set DPS (verified): 54.5. Weights run: 2.7s. Verify run: 1.5s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.116, strength=1.995 ± 0.139, agility=0.181 ± 0.041, crit=0.260 ± 0.012 per rating point (14 rating = 1%, 3.640 per %), hit=0.067 ± 0.002 per rating point (10 rating = 1%, 0.675 per %), melee_haste=2.583 ± 0.205

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.20 DPS) | yes | Veteran's Chain Helm (250498, -0.09 DPS) [crafted]; Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.65 DPS) | yes | Kaleidoscope Chain (13084, -0.25 DPS) [world_drop]; River Pride Choker (13087, -0.28 DPS) [world_drop]; Sentinel's Medallion (19541, -0.58 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.65 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.04 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.09 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.46 DPS) | yes | Sergeant Major's Cape (16315, -0.06 DPS) [pvp]; Lambent Scale Cloak (4706, -0.09 DPS) [world_drop]; Slayer's Cape (14752, -0.09 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.9 attack_power points (1.39 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.28 DPS) [crafted]; Hard Gold Cuirass (250533, -0.37 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.74 DPS) | yes | Yorgen Bracers (13012, -0.16 DPS) [world_drop]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Cultist's Armguards (270032, -0.28 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.02 DPS) | yes | The Frozen Clutch (23170, -0.09 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.19 DPS) [world]; Mail Combat Gauntlets (4075, -0.25 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.11 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.20 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.13 DPS) [crafted]; Golden Scale Leggings (3843, -0.19 DPS) [crafted]; Chausses of Westfall (6087, -0.19 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 15.2 attack_power points (0.71 DPS) | yes | Hard Gold Boots (250534, -0.06 DPS) [crafted]; Disjointed Shoes (277226, -0.15 DPS) [quest]; Glimmering Mail Greaves (4073, -0.15 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.5 attack_power points (0.76 DPS) | yes | Tiger Band (6749, -0.21 DPS) [quest]; Silverlaine's Family Seal (6321, -0.30 DPS) [dungeon]; Ironspine's Eye (7686, -0.32 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 13.1 attack_power points (0.60 DPS) | yes | Tiger Band (6749, -0.05 DPS) [quest]; Silverlaine's Family Seal (6321, -0.14 DPS) [dungeon]; Ironspine's Eye (7686, -0.16 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (16.00 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.8 attack_power points (15.52 DPS) | yes | Royal Diplomatic Scepter (9457, -0.96 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -5.10 DPS) [quest]; Shield of Thorsen (13079, -14.87 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.14 DPS) [world_drop]; Fine Longbow (11304, -0.23 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-350511005050010050-000000000000000000)

Set DPS (verified): 90.5. Weights run: 3.1s. Verify run: 1.6s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.186, strength=1.773 ± 0.242, agility=not significant (0.300 ± 0.097), crit=0.508 ± 0.026 per rating point (14 rating = 1%, 7.111 per %), hit=0.094 ± 0.003 per rating point (10 rating = 1%, 0.941 per %), melee_haste=3.798 ± 0.339

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.2 attack_power points (1.67 DPS) | yes | Chromite Barbute (8142, -0.29 DPS) [dungeon]; Icemetal Barbute (10763, -0.29 DPS) [dungeon]; Hard Gold Coif (250537, -0.29 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.10 DPS) | yes | Ghostshard Talisman (7731, -0.33 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.65 DPS) [world_drop]; River Pride Choker (13087, -0.71 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 19.5 attack_power points (1.08 DPS) | yes | Chromite Pauldrons (8144, -0.10 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 12.4 attack_power points (0.69 DPS) | yes | Dark Hooded Cape (5257, -0.13 DPS) [world]; Wolfmaster Cape (6314, -0.13 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.28 DPS) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 26.9 attack_power points (1.49 DPS) | yes | Avenger's Armor (1488, -0.02 DPS) [dungeon]; Jouster's Chestplate (8157, -0.02 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.12 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.10 DPS) | yes | Pugilist Bracers (4438, -0.32 DPS) [dungeon]; Ravager's Armguards (14770, -0.35 DPS) [world_drop]; Yorgen Bracers (13012, -0.47 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.77 DPS) | yes | Truesilver Gauntlets (7938, -0.20 DPS) [crafted]; Gloves of Holy Might (867, -0.27 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.43 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.66 DPS) | yes | Boar Champion's Belt (10768, -0.19 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.19 DPS) [rep]; Highlander's Chain Girdle (20090, -0.33 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 37.2 attack_power points (2.06 DPS) | yes | Firemane Leggings (13129, -0.20 DPS) [world_drop]; Orcish War Leggings (7929, -0.39 DPS) [crafted]; Symbolic Legplates (14829, -0.49 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 25.1 attack_power points (1.39 DPS) | yes | Prowler's Leather Shoes (252465, -0.20 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Skirmisher's Mail Boots (252564, -0.31 DPS) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.10 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.27 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.10 DPS) | yes | Protector's Band (19515, +0.00 DPS, sim-verified) [rep]; Thunderbrow Ring (13097, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (26.21 DPS) | yes | Ardent Custodian (868, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (90.5 DPS) | yes | Shoni's Disarming Tool (9608, -12.01 DPS) [quest]; Savage Boar's Guard (10767, -23.34 DPS) [dungeon]; Ardent Custodian (868, -28.70 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.14 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.28 DPS) [vendor]; Bow of Searing Arrows (2825, -0.52 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35100000000000000-350511005050010051-000000000000000000)

Set DPS (verified): 161.9. Weights run: 3.3s. Verify run: 2.3s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.248, strength=2.192 ± 0.328, agility=0.573 ± 0.112, crit=0.711 ± 0.034 per rating point (14 rating = 1%, 9.950 per %), hit=0.177 ± 0.006 per rating point (10 rating = 1%, 1.768 per %), melee_haste=5.497 ± 0.473

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Plate Helm (220804, -0.01 DPS) [vendor]; Sunscale Helmet (14849, -0.10 DPS) [world_drop]; Embrace of the Lycan (9479, -4.00 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.14 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 34.1 attack_power points (1.94 DPS) | yes | Wyrmslayer Spaulders (13066, -0.18 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.27 DPS) [crafted]; Officer's Pauldrons (250576, -1.39 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.2 attack_power points (1.20 DPS) | yes | Sergeant Major's Cape (16336, -0.26 DPS) [pvp]; Dark Hooded Cape (5257, -0.38 DPS) [world]; Bloodlust Cape (14801, -1.16 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 52.6 attack_power points (2.99 DPS) | yes | Mixologist's Tunic (12793, +0.00 DPS, sim-verified) [dungeon]; Valorous Chestguard (8274, -0.50 DPS) [world_drop]; Knight's Plate Hauberk (220794, -0.56 DPS) [vendor] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 30.7 attack_power points (1.74 DPS) | yes | Bracers of the Stone Princess (17714, -0.15 DPS) [dungeon]; Arena Bands (18711, -0.15 DPS) [world]; Officer's Wristguards (250581, -0.18 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 54.0 attack_power points (3.07 DPS) | yes | Prowler's Leather Gauntlets (252547, -1.03 DPS) [crafted]; Maddening Gauntlets (11867, -1.04 DPS) [quest]; Officer's Gloves (250551, -1.05 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 47.5 attack_power points (2.70 DPS) | yes | Belt of the Gladiator (13134, -0.46 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.57 DPS) [crafted]; Atal'alarion's Tusk Ring (10798, -0.69 DPS, sim-verified) [dungeon] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 48.2 attack_power points (2.74 DPS) | yes | Scarlet Leggings (10330, -0.12 DPS) [dungeon]; Silvershell Leggings (10633, -0.25 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.25 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 38.1 attack_power points (2.17 DPS) | yes | Prowler's Leather Boots (252468, -0.19 DPS) [crafted]; Officer's Sabatons (250561, -0.25 DPS) [crafted]; Officer's Boots (250546, -0.32 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 27.1 attack_power points (1.54 DPS) | yes | Mark of Kern (2262, -0.40 DPS) [dungeon]; Assault Band (13095, -0.40 DPS) [world_drop]; Thunderbrow Ring (13097, -0.44 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.8 attack_power points (1.24 DPS) | yes | Assault Band (13095, -0.10 DPS) [world_drop]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Mark of Kern (2262, -2.68 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.46 DPS, sim-verified) [crafted] |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodrazor (809, +0.00 DPS) [world_drop]; Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Doomforged Straightedge (12535) | Blackrock Depths: Anvilrage Overseer [dungeon] | sim-verified (+7.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Claw of Celebras (17738, -3.00 DPS) [dungeon]; Shadowblade (2163, -7.71 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -17.63 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.11 DPS) [dungeon]; Dark Iron Rifle (16004, -2.10 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Hanzo Sword; off_hand: Doomforged Straightedge; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35311103002000000-350511005050010051-000000000000000000)

Set DPS (verified): 265.9. Weights run: 3.4s. Verify run: 2.1s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.274, strength=2.173 ± 0.373, agility=not significant (0.388 ± 0.117), crit=0.808 ± 0.041 per rating point (14 rating = 1%, 11.309 per %), hit=0.196 ± 0.006 per rating point (10 rating = 1%, 1.960 per %), melee_haste=5.744 ± 0.528

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 65.6 attack_power points (4.59 DPS) | yes | Field Marshal's Plate Helm (231538, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -0.47 DPS) [vendor]; Crown of Heroism (226860, -10.75 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (265.9 DPS) | yes | Imperial Jewel (11933, -0.23 DPS) [dungeon]; Will of the Martyr (17044, -0.37 DPS) [quest]; Rage of Mugamba (19577, -5.26 DPS, sim-verified) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 45.7 attack_power points (3.19 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 36.9 attack_power points (2.58 DPS) | yes | Cloak of the Honor Guard (20073, -0.07 DPS) [rep]; Howler's Furs (272414, -0.49 DPS) [vendor]; Shadewood Cloak (18328, -0.61 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (265.9 DPS) | yes | Obsidian Mail Tunic (22191, -0.19 DPS) [crafted]; Cadaverous Armor (14637, -0.66 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -17.30 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-verified (265.9 DPS) | yes | Marshal's Plate Bracers (16481, -0.33 DPS) [pvp]; Windtalker's Wristguards (19582, -0.45 DPS) [rep]; Bracers of Undead Slaying (23090, -7.46 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (265.9 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.57 DPS) [pvp]; Radiant Gloves of the Dawn (227817, -0.63 DPS) [vendor]; Razor Gauntlets (18326, -8.88 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 70.0 attack_power points (4.89 DPS) | yes | Ferocity of the Timbermaw (227805, -0.48 DPS) [vendor]; Marshal's Plate Girdle (16482, -0.77 DPS) [pvp]; Dense Timbermaw Belt (227807, -2.91 DPS, sim-verified) [vendor] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (265.9 DPS) | yes | Titanic Leggings (22385, -0.20 DPS) [crafted]; Marshal's Plate Legguards (231540, -1.06 DPS) [pvp]; Cloudkeeper Legplates (14554, -5.90 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 45.4 attack_power points (3.17 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Clutchlord's Stompers (275627, -0.29 DPS) [crafted] |
| finger1 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-verified (265.9 DPS) | yes | Don Julio's Band (19325, -0.08 DPS) [rep]; Myrmidon's Signet (2246, -0.42 DPS) [world_drop]; Naglering (11669, -4.96 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (265.9 DPS) | yes | Don Julio's Band (19325, -0.07 DPS) [rep]; Myrmidon's Signet (2246, -0.41 DPS) [world_drop]; Naglering (11669, -5.54 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (265.9 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Diamond Flask (20130, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (265.9 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.27 DPS) [crafted]; Blackhand's Breadth (13965, -1.45 DPS, sim-verified) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (265.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Quel'Serrar (18348, -12.32 DPS, sim-verified) [quest] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (265.9 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -17.33 DPS) [dungeon]; Skullflame Shield (1168, -66.44 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (265.9 DPS) | yes | Stinging Bow (10624, -0.10 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.10 DPS) [world_drop]; Dark Iron Rifle (16004, -3.30 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Highlander's Plate Spaulders; back: Shroud of Domination; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Band of the Ogre King; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Riphook

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (troll, 00000000000000000-350300000000000000-000000000000000000)

Set DPS (verified): 30.5. Weights run: 2.5s. Verify run: 1.5s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.026, strength=1.998 ± 0.032, agility=0.104 ± 0.011, crit=0.169 ± 0.004 per rating point (14 rating = 1%, 2.366 per %), hit=0.050 ± 0.001 per rating point (10 rating = 1%, 0.499 per %), melee_haste=1.728 ± 0.034

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.23 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.70 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.73 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.27 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.37 DPS) | yes | Bristlebark Bindings (14569, -0.21 DPS) [world_drop]; Runed Copper Bracers (2854, -0.22 DPS) [crafted]; Raptorcrest Bracers (270010, -0.23 DPS, sim-verified) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.59 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.21 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.66 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.29 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.5 attack_power points (0.68 DPS) | yes | Defender's Leather Pants (252445, -0.00 DPS) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.39 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.14 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (8.69 DPS) | yes | Diamond Hammer (2194, -0.41 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.25 DPS) [world_drop]; Ruga's Bulwark (7120, -8.47 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.15 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-350511005010000000-000000000000000000)

Set DPS (verified): 55.1. Weights run: 2.7s. Verify run: 1.5s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.116, strength=1.995 ± 0.139, agility=0.181 ± 0.041, crit=0.260 ± 0.012 per rating point (14 rating = 1%, 3.640 per %), hit=0.067 ± 0.002 per rating point (10 rating = 1%, 0.675 per %), melee_haste=2.583 ± 0.205

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.20 DPS) | yes | Veteran's Chain Helm (250498, -0.09 DPS) [crafted]; Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.65 DPS) | yes | Kaleidoscope Chain (13084, -0.25 DPS) [world_drop]; River Pride Choker (13087, -0.28 DPS) [world_drop]; Scout's Medallion (19537, -0.58 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.65 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.04 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.46 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.09 DPS) [world_drop]; Slayer's Cape (14752, -0.09 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.9 attack_power points (1.39 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.28 DPS) [crafted]; Hard Gold Cuirass (250533, -0.37 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.74 DPS) | yes | Yorgen Bracers (13012, -0.16 DPS) [world_drop]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.27 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.02 DPS) | yes | The Frozen Clutch (23170, -0.09 DPS) [dungeon]; Warsong Gauntlets (16978, -0.09 DPS) [quest]; Bonefist Gauntlets (4465, -0.19 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.11 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.20 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.13 DPS) [crafted]; Golden Scale Leggings (3843, -0.19 DPS) [crafted]; Slayer's Pants (14757, -0.19 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 15.2 attack_power points (0.71 DPS) | yes | Glimmering Mail Greaves (4073, -0.15 DPS) [world_drop]; Slayer's Slippers (14756, -0.15 DPS) [world_drop]; Hard Gold Boots (250534, -0.40 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.5 attack_power points (0.76 DPS) | yes | Tiger Band (6749, -0.21 DPS) [quest]; Silverlaine's Family Seal (6321, -0.30 DPS) [dungeon]; Ironspine's Eye (7686, -0.32 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 13.1 attack_power points (0.60 DPS) | yes | Silverlaine's Family Seal (6321, -0.14 DPS) [dungeon]; Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Tiger Band (6749, -0.37 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (16.00 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.8 attack_power points (15.52 DPS) | yes | Royal Diplomatic Scepter (9457, -0.73 DPS, sim-verified) [dungeon]; Shield of Thorsen (13079, -14.87 DPS) [world_drop]; Slayer's Shield (15892, -14.94 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.14 DPS) [world_drop]; Fine Longbow (11304, -0.23 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-350511005050010050-000000000000000000)

Set DPS (verified): 86.5. Weights run: 3.1s. Verify run: 1.7s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.186, strength=1.773 ± 0.242, agility=not significant (0.300 ± 0.097), crit=0.508 ± 0.026 per rating point (14 rating = 1%, 7.111 per %), hit=0.094 ± 0.003 per rating point (10 rating = 1%, 0.941 per %), melee_haste=3.798 ± 0.339

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.2 attack_power points (1.67 DPS) | yes | Icemetal Barbute (10763, -0.29 DPS) [dungeon]; Hard Gold Coif (250537, -0.29 DPS) [crafted]; Chromite Barbute (8142, -0.52 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.10 DPS) | yes | Ghostshard Talisman (7731, -0.33 DPS) [dungeon]; Ethereal Talisman (4430, -0.55 DPS) [quest]; Kaleidoscope Chain (13084, -0.65 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 19.5 attack_power points (1.08 DPS) | yes | Chromite Pauldrons (8144, -0.10 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 12.4 attack_power points (0.69 DPS) | yes | Dark Hooded Cape (5257, -0.13 DPS) [world]; Wolfmaster Cape (6314, -0.13 DPS) [dungeon]; Wildhunter Cloak (16658, -0.13 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 26.9 attack_power points (1.49 DPS) | yes | Avenger's Armor (1488, -0.02 DPS) [dungeon]; Jouster's Chestplate (8157, -0.02 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.12 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.10 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.77 DPS) | yes | Truesilver Gauntlets (7938, -0.20 DPS) [crafted]; Gloves of Holy Might (867, -0.27 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.43 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.66 DPS) | yes | Boar Champion's Belt (10768, -0.19 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.19 DPS) [rep]; Defiler's Chain Girdle (20152, -0.33 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 37.2 attack_power points (2.06 DPS) | yes | Firemane Leggings (13129, -0.20 DPS) [world_drop]; Orcish War Leggings (7929, -0.39 DPS) [crafted]; Symbolic Legplates (14829, -0.49 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 25.1 attack_power points (1.39 DPS) | yes | Prowler's Leather Shoes (252465, -0.20 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Skirmisher's Mail Boots (252564, -0.31 DPS) [crafted] |
| finger1 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.10 DPS) | yes | Thunderbrow Ring (13097, -0.27 DPS) [world_drop]; Suspicious Spare Part (274754, -0.42 DPS) [vendor]; Assault Band (13095, -0.87 DPS, sim-verified) [world_drop] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Suspicious Spare Part (274754, -0.23 DPS) [vendor]; Assault Band (13095, -0.66 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (26.21 DPS) | yes | Ardent Custodian (868, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (+23.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Savage Boar's Guard (10767, -23.34 DPS) [dungeon]; Skullance Shield (13081, -23.57 DPS) [world_drop]; Ardent Custodian (868, -23.90 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.14 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.28 DPS) [vendor]; Bow of Searing Arrows (2825, -0.58 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Legionnaire's Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 35100000000000000-350511005050010051-000000000000000000)

Set DPS (verified): 145.7. Weights run: 3.3s. Verify run: 2.3s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.248, strength=2.192 ± 0.328, agility=0.573 ± 0.112, crit=0.711 ± 0.034 per rating point (14 rating = 1%, 9.950 per %), hit=0.177 ± 0.006 per rating point (10 rating = 1%, 1.768 per %), melee_haste=5.497 ± 0.473

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | sim-verified (145.7 DPS) | yes | Blood Guard's Plate Helm (220803, -0.01 DPS) [vendor]; Sunscale Helmet (14849, -0.10 DPS) [world_drop]; Embrace of the Lycan (9479, -2.79 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.14 DPS) | yes | Skibi's Pendant (13089, -0.09 DPS) [world_drop]; Woven Ivy Necklace (19159, -0.10 DPS) [quest]; Ghostshard Talisman (7731, -0.34 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 33.1 attack_power points (1.88 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.12 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.22 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.2 attack_power points (1.20 DPS) | yes | First Sergeant's Cloak (16340, -0.26 DPS) [pvp]; Dark Hooded Cape (5257, -0.38 DPS) [world]; Bloodlust Cape (14801, -1.98 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 52.6 attack_power points (2.99 DPS) | yes | Valorous Chestguard (8274, -0.50 DPS) [world_drop]; Stone Guard's Plate Armor (220801, -0.56 DPS) [vendor]; Mixologist's Tunic (12793, -0.91 DPS, sim-verified) [dungeon] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 30.7 attack_power points (1.74 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Berserker Bracers (19580, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 54.0 attack_power points (3.07 DPS) | yes | Prowler's Leather Gauntlets (252547, -1.03 DPS) [crafted]; Truesilver Gauntlets (7938, -1.07 DPS) [crafted]; Officer's Gloves (250551, -2.41 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 47.5 attack_power points (2.70 DPS) | yes | Belt of the Gladiator (13134, -0.46 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.57 DPS) [crafted]; Atal'alarion's Tusk Ring (10798, -1.98 DPS, sim-verified) [dungeon] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 48.2 attack_power points (2.74 DPS) | yes | Scarlet Leggings (10330, -0.12 DPS) [dungeon]; Silvershell Leggings (10633, -0.25 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.25 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 38.1 attack_power points (2.17 DPS) | yes | Prowler's Leather Boots (252468, -0.19 DPS) [crafted]; Officer's Sabatons (250561, -0.25 DPS) [crafted]; Officer's Boots (250546, -0.32 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 27.1 attack_power points (1.54 DPS) | yes | Blackstone Ring (17713, -0.30 DPS) [dungeon]; Mark of Kern (2262, -0.40 DPS) [dungeon]; Assault Band (13095, -0.40 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.36 DPS) | yes | Blackstone Ring (17713, -0.13 DPS) [dungeon]; Mark of Kern (2262, -0.23 DPS) [dungeon]; Assault Band (13095, -0.23 DPS) [world_drop] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+8.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -5.93 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hanzo Sword (8190, -2.42 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (30.98 DPS) | yes | Doomforged Straightedge (12535, -1.91 DPS, sim-verified) [dungeon]; Claw of Celebras (17738, -3.57 DPS) [dungeon]; White Bone Shredder (11863, -5.53 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.11 DPS) [dungeon]; Dark Iron Rifle (16004, -3.10 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Diamond Flask; trinket2: Rune of the Guard Captain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 35311103002000000-350511005050010051-000000000000000000)

Set DPS (verified): 247.7. Weights run: 3.4s. Verify run: 2.2s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.274, strength=2.173 ± 0.373, agility=not significant (0.388 ± 0.117), crit=0.808 ± 0.041 per rating point (14 rating = 1%, 11.309 per %), hit=0.196 ± 0.006 per rating point (10 rating = 1%, 1.960 per %), melee_haste=5.744 ± 0.528

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 65.6 attack_power points (4.59 DPS) | yes | Warlord's Plate Headpiece (231535, +0.00 DPS) [pvp]; Champion's Plate Helm (227043, -0.47 DPS) [pvp]; Crown of Heroism (226860, -10.77 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (247.7 DPS) | yes | Imperial Jewel (11933, -0.23 DPS) [dungeon]; Will of the Martyr (17044, -0.37 DPS) [quest]; Rage of Mugamba (19577, -3.68 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 45.7 attack_power points (3.19 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Black Dragonscale Shoulders (15051, -2.75 DPS, sim-verified) [crafted] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 36.9 attack_power points (2.58 DPS) | yes | Deathguard's Cloak (20068, -0.07 DPS) [rep]; Howler's Furs (272414, -0.49 DPS) [vendor]; Shadewood Cloak (18328, -0.61 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (247.7 DPS) | yes | Obsidian Mail Tunic (22191, -0.19 DPS) [crafted]; Cadaverous Armor (14637, -0.66 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -16.96 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (247.7 DPS) | yes | General's Plate Armguards (16546, -0.33 DPS) [pvp]; Windtalker's Wristguards (19582, -0.45 DPS) [rep]; Bracers of Undead Slaying (23090, -6.38 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (247.7 DPS) | yes | General's Plate Gauntlets (231532, -0.57 DPS) [pvp]; Radiant Gloves of the Dawn (227817, -0.63 DPS) [vendor]; Razor Gauntlets (18326, -7.26 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 70.0 attack_power points (4.89 DPS) | yes | Ferocity of the Timbermaw (227805, -0.48 DPS) [vendor]; General's Plate Girdle (16547, -0.77 DPS) [pvp]; Dense Timbermaw Belt (227807, -3.85 DPS, sim-verified) [vendor] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (247.7 DPS) | yes | Titanic Leggings (22385, -0.20 DPS) [crafted]; Outrider's Plate Legguards (22651, -0.64 DPS) [rep]; Cloudkeeper Legplates (14554, -7.43 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 45.4 attack_power points (3.17 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Clutchlord's Stompers (275627, -0.29 DPS) [crafted] |
| finger1 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-verified (247.7 DPS) | yes | Don Julio's Band (19325, -0.08 DPS) [rep]; Myrmidon's Signet (2246, -0.42 DPS) [world_drop]; Naglering (11669, -4.58 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (247.7 DPS) | yes | Don Julio's Band (19325, -0.07 DPS) [rep]; Myrmidon's Signet (2246, -0.41 DPS) [world_drop]; Naglering (11669, -5.28 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (247.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (247.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (247.7 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Felstriker (12590, -5.81 DPS, sim-verified) [dungeon] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (247.7 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -17.33 DPS) [dungeon]; Skullflame Shield (1168, -66.59 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (247.7 DPS) | yes | Stinging Bow (10624, -0.10 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.10 DPS) [world_drop]; Dark Iron Rifle (16004, -4.15 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Defiler's Plate Spaulders; back: Shroud of Domination; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Band of the Ogre King; finger2: Legionnaire's Band; trinket2: Blackhand's Breadth; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Riphook

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

