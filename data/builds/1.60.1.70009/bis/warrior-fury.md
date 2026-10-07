# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-350300000000000000-000000000000000000)

Set DPS (verified): 33.2. Weights run: 2.1s. Verify run: 1.2s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=1.995 ± 0.003, agility=0.114 ± 0.003, crit=0.163 ± 0.004 per rating point (14 rating = 1%, 2.285 per %), hit=0.048 ± 0.001 per rating point (10 rating = 1%, 0.482 per %), melee_haste=1.669 ± 0.033

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.76 DPS) | yes | Defender's Leather Hood (252447, -0.30 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.72 DPS) [crafted]; Brawler's Leather Hood (252504, -0.72 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.7 attack_power points (0.03 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.23 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.21 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.30 DPS) | yes | Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.08 DPS) [quest]; Dark Leather Cloak (2316, -0.14 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.76 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.23 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.44 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.38 DPS) | yes | Bravo's Armbands (270015, -0.21 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop]; Cryptwalker Bracers (280095, -0.22 DPS, sim-verified) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.61 DPS) | yes | Polar Gauntlets (7606, -0.15 DPS) [quest]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Gold-flecked Gloves (5195, -0.22 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.68 DPS) | yes | Cobrahn's Grasp (6460, -0.23 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Hulking Belt (14746, -0.31 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 21.9 attack_power points (0.83 DPS) | yes | Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.21 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.40 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.32 DPS) | yes | The 1 Ring (8350, -0.24 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop]; Signet of the Zhevra (285330, -0.29 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.30 DPS) | yes | Ring of the Moon (12052, -0.23 DPS) [world_drop]; Signet of the Zhevra (285330, -0.28 DPS) [world]; The 1 Ring (8350, -0.47 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.46 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (9.00 DPS) | yes | Diamond Hammer (2194, -0.39 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.55 DPS) [world_drop]; Furen's Favor (6970, -8.77 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.15 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Daryl's Hunting Rifle (2904, -0.08 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 00000000000000000-350511005010000000-000000000000000000)

Set DPS (verified): 60.9. Weights run: 2.2s. Verify run: 1.2s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.016, strength=1.980 ± 0.015, agility=0.199 ± 0.010, crit=0.285 ± 0.014 per rating point (14 rating = 1%, 3.984 per %), hit=0.074 ± 0.002 per rating point (10 rating = 1%, 0.739 per %), melee_haste=2.827 ± 0.224

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.7 attack_power points (1.09 DPS) | yes | Veteran's Chain Helm (250498, -0.08 DPS) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.59 DPS) | yes | Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.26 DPS) [world_drop]; Sentinel's Medallion (19541, -0.53 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.59 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.03 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.42 DPS) | yes | Sergeant Major's Cape (16315, -0.05 DPS) [pvp]; Lambent Scale Cloak (4706, -0.09 DPS) [world_drop]; Slayer's Cape (14752, -0.09 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.7 attack_power points (1.26 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.25 DPS) [crafted]; Hard Gold Cuirass (250533, -0.34 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.8 attack_power points (0.67 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Cultist's Armguards (270032, -0.25 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 21.9 attack_power points (0.93 DPS) | yes | The Frozen Clutch (23170, -0.08 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.17 DPS) [world]; Mail Combat Gauntlets (4075, -0.22 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.02 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.10 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.12 DPS) [crafted]; Golden Scale Leggings (3843, -0.18 DPS) [crafted]; Chausses of Westfall (6087, -0.18 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 15.3 attack_power points (0.65 DPS) | yes | Disjointed Shoes (277226, -0.14 DPS) [quest]; Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Hard Gold Boots (250534, -0.46 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 attack_power points (0.70 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Silverlaine's Family Seal (6321, -0.28 DPS) [dungeon]; Ironspine's Eye (7686, -0.28 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 13.1 attack_power points (0.55 DPS) | yes | Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.14 DPS) [dungeon]; Tiger Band (6749, -0.36 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.3 attack_power points (14.62 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.7 attack_power points (14.17 DPS) | yes | Royal Diplomatic Scepter (9457, -0.76 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -4.66 DPS) [quest]; Shield of Thorsen (13079, -13.58 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.38 DPS) | yes | Double-barreled Shotgun (2098, -0.10 DPS) [world_drop]; Long Battle Bow (15284, -0.13 DPS) [world_drop]; Fine Longbow (11304, -0.21 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-350511005050010050-000000000000000000)

Set DPS (verified): 101.1. Weights run: 2.6s. Verify run: 1.3s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.025, strength=1.956 ± 0.030, agility=0.408 ± 0.021, crit=0.583 ± 0.030 per rating point (14 rating = 1%, 8.169 per %), hit=0.108 ± 0.004 per rating point (10 rating = 1%, 1.081 per %), melee_haste=4.362 ± 0.390

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.6 attack_power points (1.62 DPS) | yes | Chromite Barbute (8142, -0.25 DPS) [dungeon]; Icemetal Barbute (10763, -0.30 DPS) [dungeon]; Hard Gold Coif (250537, -0.30 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.96 DPS) | yes | Ghostshard Talisman (7731, -0.29 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop]; River Pride Choker (13087, -0.59 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.5 attack_power points (1.04 DPS) | yes | Chromite Pauldrons (8144, -0.09 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.19 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.2 attack_power points (0.68 DPS) | yes | Dark Hooded Cape (5257, -0.11 DPS) [world]; Wolfmaster Cape (6314, -0.20 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.26 DPS) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.2 attack_power points (1.45 DPS) | yes | Avenger's Armor (1488, -0.04 DPS) [dungeon]; Jouster's Chestplate (8157, -0.04 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.14 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.96 DPS) | yes | Pugilist Bracers (4438, -0.21 DPS) [dungeon]; Ravager's Armguards (14770, -0.22 DPS) [world_drop]; Yorgen Bracers (13012, -0.34 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.54 DPS) | yes | Truesilver Gauntlets (7938, -0.03 DPS) [crafted]; Gloves of Holy Might (867, -0.18 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.21 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.44 DPS) | yes | Boar Champion's Belt (10768, -0.03 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.03 DPS) [rep]; Highlander's Chain Girdle (20090, -0.29 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 41.1 attack_power points (1.98 DPS) | yes | Firemane Leggings (13129, -0.19 DPS) [world_drop]; Orcish War Leggings (7929, -0.38 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 28.3 attack_power points (1.36 DPS) | yes | Prowler's Leather Shoes (252465, -0.19 DPS) [crafted]; Obsidian Greaves (13068, -0.30 DPS) [world_drop]; Blackforge Greaves (6423, -0.32 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.96 DPS) | yes | Protector's Band (19515, -0.05 DPS) [rep]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.96 DPS) | yes | Protector's Band (19515, +0.00 DPS, sim-verified) [rep]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (22.82 DPS) | yes | Ardent Custodian (868, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (101.1 DPS) | yes | Shoni's Disarming Tool (9608, -10.45 DPS) [quest]; Savage Boar's Guard (10767, -20.22 DPS) [dungeon]; Ardent Custodian (868, -39.23 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.05 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.20 DPS) [crafted]; Bow of Searing Arrows (2825, -0.52 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35100000000000000-350511005050010051-000000000000000000)

Set DPS (verified): 181.9. Weights run: 2.8s. Verify run: 1.7s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.026, strength=2.007 ± 0.030, agility=0.427 ± 0.021, crit=0.610 ± 0.029 per rating point (14 rating = 1%, 8.536 per %), hit=0.152 ± 0.005 per rating point (10 rating = 1%, 1.517 per %), melee_haste=4.716 ± 0.406

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Plate Helm (220804, -0.05 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.06 DPS) [dungeon]; Embrace of the Lycan (9479, -3.67 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.32 DPS) | yes | Skibi's Pendant (13089, -0.29 DPS) [world_drop]; Ghostshard Talisman (7731, -0.40 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.68 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 30.6 attack_power points (2.03 DPS) | yes | Wyrmslayer Spaulders (13066, -0.21 DPS) [world_drop]; Earthslag Shoulders (11632, -0.30 DPS) [dungeon]; Officer's Pauldrons (250576, -2.10 DPS, sim-verified) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.1 attack_power points (1.20 DPS) | yes | Blackveil Cape (11626, -0.00 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.23 DPS) [pvp]; Dark Hooded Cape (5257, -0.38 DPS) [world] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 48.2 attack_power points (3.19 DPS) | yes | Valorous Chestguard (8274, -0.53 DPS) [world_drop]; Knight's Plate Hauberk (220794, -0.63 DPS) [vendor]; Mixologist's Tunic (12793, -1.20 DPS, sim-verified) [dungeon] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 28.1 attack_power points (1.86 DPS) | yes | Bracers of the Stone Princess (17714, -0.01 DPS) [dungeon]; Arena Bands (18711, -0.01 DPS) [world]; Officer's Wristguards (250581, -0.23 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 48.4 attack_power points (3.21 DPS) | yes | Maddening Gauntlets (11867, -1.07 DPS) [quest]; Truesilver Gauntlets (7938, -1.08 DPS) [crafted]; Officer's Gloves (250551, -2.85 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.1 attack_power points (3.05 DPS) | yes | Belt of the Gladiator (13134, -0.66 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.85 DPS) [crafted]; Atal'alarion's Tusk Ring (10798, -2.38 DPS, sim-verified) [dungeon] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 44.2 attack_power points (2.92 DPS) | yes | Silvershell Leggings (10633, -0.27 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.27 DPS) [dungeon]; Scarlet Leggings (10330, -0.94 DPS, sim-verified) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 33.6 attack_power points (2.23 DPS) | yes | Officer's Sabatons (250561, -0.25 DPS) [crafted]; Officer's Boots (250546, -0.30 DPS) [crafted]; Prowler's Leather Boots (252468, -1.33 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 23.9 attack_power points (1.58 DPS) | yes | Mark of Kern (2262, -0.26 DPS) [dungeon]; Assault Band (13095, -0.26 DPS) [world_drop]; Thunderbrow Ring (13097, -0.44 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.5 attack_power points (1.43 DPS) | yes | Assault Band (13095, -0.10 DPS) [world_drop]; Thunderbrow Ring (13097, -0.28 DPS) [world_drop]; Mark of Kern (2262, -3.46 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+9.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -2.19 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Hanzo Sword (8190, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Doomforged Straightedge (12535) | Blackrock Depths: Anvilrage Overseer [dungeon] | sim-verified (+8.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Claw of Celebras (17738, -3.38 DPS) [dungeon]; Shadowblade (2163, -8.64 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -20.43 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.13 DPS) [dungeon]; Dark Iron Rifle (16004, -3.36 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Bloodrazor; off_hand: Doomforged Straightedge; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35311103002000000-350511005050010051-000000000000000000)

Set DPS (verified): 306.3. Weights run: 2.9s. Verify run: 1.7s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.035, strength=2.062 ± 0.041, agility=0.592 ± 0.030, crit=0.846 ± 0.043 per rating point (14 rating = 1%, 11.839 per %), hit=0.205 ± 0.007 per rating point (10 rating = 1%, 2.052 per %), melee_haste=6.013 ± 0.553

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 64.9 attack_power points (4.33 DPS) | yes | Field Marshal's Plate Helm (231538, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -0.51 DPS) [vendor]; Crown of Heroism (226860, -13.48 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (306.3 DPS) | yes | Imperial Jewel (11933, -0.26 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.27 DPS) [quest]; Rage of Mugamba (19577, -5.55 DPS, sim-verified) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 47.2 attack_power points (3.15 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, -0.02 DPS) [pvp] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 37.0 attack_power points (2.47 DPS) | yes | Shroud of Domination (22337, -0.13 DPS) [dungeon]; Howler's Furs (272414, -0.46 DPS) [vendor]; Cape of the Black Baron (13340, -0.54 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (306.3 DPS) | yes | Obsidian Mail Tunic (22191, -0.18 DPS) [crafted]; Cadaverous Armor (14637, -0.62 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -21.28 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-verified (306.3 DPS) | yes | Marshal's Plate Bracers (16481, -0.31 DPS) [pvp]; Windtalker's Wristguards (19582, -0.39 DPS) [rep]; Bracers of Undead Slaying (23090, -9.05 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (306.3 DPS) | yes | Radiant Gloves of the Dawn (227817, -0.59 DPS) [vendor]; Marshal's Plate Gauntlets (231541, -0.61 DPS) [pvp]; Razor Gauntlets (18326, -9.87 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 67.5 attack_power points (4.51 DPS) | yes | Ferocity of the Timbermaw (227805, -0.26 DPS) [vendor]; Marshal's Plate Girdle (16482, -0.59 DPS) [pvp]; Dense Timbermaw Belt (227807, -3.60 DPS, sim-verified) [vendor] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (306.3 DPS) | yes | Titanic Leggings (22385, -0.24 DPS) [crafted]; Warbear Woolies (15065, -0.87 DPS) [crafted]; Cloudkeeper Legplates (14554, -10.63 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 43.3 attack_power points (2.89 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Scalegut Treaders (275618, -0.04 DPS) [crafted] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (306.3 DPS) | yes | Band of the Ogre King (18522, -0.16 DPS) [dungeon]; Myrmidon's Signet (2246, -0.43 DPS) [world_drop]; Naglering (11669, -6.99 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (306.3 DPS) | yes | Band of the Ogre King (18522, -0.07 DPS) [dungeon]; Myrmidon's Signet (2246, -0.34 DPS) [world_drop]; Naglering (11669, -13.72 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (306.3 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (306.3 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.21 DPS) [crafted] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (306.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Quel'Serrar (18348, -11.99 DPS, sim-verified) [quest] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (306.3 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -16.66 DPS) [dungeon]; Skullflame Shield (1168, -100.90 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (306.3 DPS) | yes | Bloodseeker (19107, -0.09 DPS) [quest]; Skull Splitting Crossbow (13039, -0.12 DPS) [world_drop]; Dark Iron Rifle (16004, -5.26 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Highlander's Plate Spaulders; back: Cloak of the Honor Guard; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Protector's Band; finger2: Don Julio's Band; trinket2: Hand of Justice; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Riphook

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (troll, 00000000000000000-350300000000000000-000000000000000000)

Set DPS (verified): 30.5. Weights run: 2.1s. Verify run: 1.2s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=1.995 ± 0.003, agility=0.114 ± 0.003, crit=0.163 ± 0.004 per rating point (14 rating = 1%, 2.285 per %), hit=0.048 ± 0.001 per rating point (10 rating = 1%, 0.482 per %), melee_haste=1.669 ± 0.033

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.76 DPS) | yes | Defender's Leather Hood (252447, -0.23 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.72 DPS) [crafted]; Brawler's Leather Hood (252504, -0.72 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.7 attack_power points (0.03 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.23 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.21 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.30 DPS) | yes | Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.08 DPS) [quest]; Subterranean Cape (14149, -0.08 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.76 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.23 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.27 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.38 DPS) | yes | Bristlebark Bindings (14569, -0.21 DPS) [world_drop]; Runed Copper Bracers (2854, -0.23 DPS) [crafted]; Raptorcrest Bracers (270010, -0.23 DPS, sim-verified) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.61 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.21 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.68 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Hulking Belt (14746, -0.31 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.5 attack_power points (0.70 DPS) | yes | Defender's Leather Pants (252445, -0.00 DPS) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.40 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 attack_power points (0.32 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.24 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.46 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (9.00 DPS) | yes | Diamond Hammer (2194, -0.41 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.55 DPS) [world_drop]; Ruga's Bulwark (7120, -8.77 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.15 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Heavy Shortbow (3036, -0.08 DPS) [world_drop]; Orcish Battle Bow (5346, -0.08 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-350511005010000000-000000000000000000)

Set DPS (verified): 61.4. Weights run: 2.2s. Verify run: 1.2s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.016, strength=1.980 ± 0.015, agility=0.199 ± 0.010, crit=0.285 ± 0.014 per rating point (14 rating = 1%, 3.984 per %), hit=0.074 ± 0.002 per rating point (10 rating = 1%, 0.739 per %), melee_haste=2.827 ± 0.224

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.7 attack_power points (1.09 DPS) | yes | Veteran's Chain Helm (250498, -0.08 DPS) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.59 DPS) | yes | Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.26 DPS) [world_drop]; Scout's Medallion (19537, -0.53 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.59 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.03 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.42 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.09 DPS) [world_drop]; Slayer's Cape (14752, -0.09 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.7 attack_power points (1.26 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.25 DPS) [crafted]; Hard Gold Cuirass (250533, -0.34 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.8 attack_power points (0.67 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.24 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 21.9 attack_power points (0.93 DPS) | yes | The Frozen Clutch (23170, -0.08 DPS) [dungeon]; Warsong Gauntlets (16978, -0.09 DPS) [quest]; Bonefist Gauntlets (4465, -0.17 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.02 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.10 DPS) | yes | Veteran's Silvered Chain Leggings (250523, +0.00 DPS, sim-verified) [crafted]; Golden Scale Leggings (3843, -0.18 DPS) [crafted]; Slayer's Pants (14757, -0.18 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 15.3 attack_power points (0.65 DPS) | yes | Hard Gold Boots (250534, -0.06 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 attack_power points (0.70 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Silverlaine's Family Seal (6321, -0.28 DPS) [dungeon]; Ironspine's Eye (7686, -0.28 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 13.1 attack_power points (0.55 DPS) | yes | Tiger Band (6749, -0.05 DPS) [quest]; Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.3 attack_power points (14.62 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.7 attack_power points (14.17 DPS) | yes | Royal Diplomatic Scepter (9457, -0.86 DPS, sim-verified) [dungeon]; Shield of Thorsen (13079, -13.58 DPS) [world_drop]; Slayer's Shield (15892, -13.64 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.38 DPS) | yes | Double-barreled Shotgun (2098, -0.10 DPS) [world_drop]; Long Battle Bow (15284, -0.13 DPS) [world_drop]; Fine Longbow (11304, -0.21 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-350511005050010050-000000000000000000)

Set DPS (verified): 95.6. Weights run: 2.6s. Verify run: 1.4s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.025, strength=1.956 ± 0.030, agility=0.408 ± 0.021, crit=0.583 ± 0.030 per rating point (14 rating = 1%, 8.169 per %), hit=0.108 ± 0.004 per rating point (10 rating = 1%, 1.081 per %), melee_haste=4.362 ± 0.390

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.6 attack_power points (1.62 DPS) | yes | Icemetal Barbute (10763, -0.30 DPS) [dungeon]; Hard Gold Coif (250537, -0.30 DPS) [crafted]; Chromite Barbute (8142, -0.52 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.96 DPS) | yes | Ghostshard Talisman (7731, -0.29 DPS) [dungeon]; Ethereal Talisman (4430, -0.41 DPS) [quest]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.5 attack_power points (1.04 DPS) | yes | Chromite Pauldrons (8144, -0.09 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.19 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 14.2 attack_power points (0.68 DPS) | yes | Dark Hooded Cape (5257, -0.11 DPS) [world]; Wolfmaster Cape (6314, -0.20 DPS) [dungeon]; Wildhunter Cloak (16658, -0.20 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.2 attack_power points (1.45 DPS) | yes | Avenger's Armor (1488, -0.04 DPS) [dungeon]; Jouster's Chestplate (8157, -0.04 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.14 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.96 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.54 DPS) | yes | Truesilver Gauntlets (7938, -0.03 DPS) [crafted]; Gloves of Holy Might (867, -0.18 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.21 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.44 DPS) | yes | Boar Champion's Belt (10768, -0.03 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.03 DPS) [rep]; Tharg's Shoelace (9705, -0.22 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 41.1 attack_power points (1.98 DPS) | yes | Firemane Leggings (13129, -0.19 DPS) [world_drop]; Orcish War Leggings (7929, -0.38 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 28.3 attack_power points (1.36 DPS) | yes | Prowler's Leather Shoes (252465, -0.19 DPS) [crafted]; Obsidian Greaves (13068, -0.30 DPS) [world_drop]; Blackforge Greaves (6423, -0.32 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.96 DPS) | yes | Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Suspicious Spare Part (274754, -0.30 DPS) [vendor]; Assault Band (13095, -0.87 DPS, sim-verified) [world_drop] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor]; Assault Band (13095, -0.66 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (22.82 DPS) | yes | Ardent Custodian (868, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (+33.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Savage Boar's Guard (10767, -20.22 DPS) [dungeon]; Skullance Shield (13081, -20.43 DPS) [world_drop]; Ardent Custodian (868, -33.41 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.05 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.20 DPS) [crafted]; Bow of Searing Arrows (2825, -0.58 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Legionnaire's Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 35100000000000000-350511005050010051-000000000000000000)

Set DPS (verified): 167.4. Weights run: 2.8s. Verify run: 1.8s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.026, strength=2.007 ± 0.030, agility=0.427 ± 0.021, crit=0.610 ± 0.029 per rating point (14 rating = 1%, 8.536 per %), hit=0.152 ± 0.005 per rating point (10 rating = 1%, 1.517 per %), melee_haste=4.716 ± 0.406

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Plate Helm (220803, -0.05 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.06 DPS) [dungeon]; Embrace of the Lycan (9479, -4.91 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.32 DPS) | yes | Woven Ivy Necklace (19159, -0.27 DPS) [quest]; Skibi's Pendant (13089, -0.29 DPS) [world_drop]; Ghostshard Talisman (7731, -0.40 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 29.5 attack_power points (1.95 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.13 DPS) [world_drop]; Earthslag Shoulders (11632, -0.23 DPS) [dungeon] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | First Sergeant's Cloak (16340, -0.23 DPS) [pvp]; Dark Hooded Cape (5257, -0.38 DPS) [world]; Bloodlust Cape (14801, -1.62 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 48.2 attack_power points (3.19 DPS) | yes | Mixologist's Tunic (12793, -0.49 DPS) [dungeon]; Valorous Chestguard (8274, -0.53 DPS) [world_drop]; Stone Guard's Plate Armor (220801, -0.63 DPS) [vendor] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 28.1 attack_power points (1.86 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 48.4 attack_power points (3.21 DPS) | yes | Truesilver Gauntlets (7938, -1.08 DPS) [crafted]; Gauntlets of Divinity (7724, -1.09 DPS) [dungeon]; Officer's Gloves (250551, -1.55 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.1 attack_power points (3.05 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.66 DPS) [dungeon]; Belt of the Gladiator (13134, -0.66 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.85 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 44.2 attack_power points (2.92 DPS) | yes | Scarlet Leggings (10330, -0.13 DPS) [dungeon]; Silvershell Leggings (10633, -0.27 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.27 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 33.6 attack_power points (2.23 DPS) | yes | Officer's Sabatons (250561, -0.25 DPS) [crafted]; Officer's Boots (250546, -0.30 DPS) [crafted]; Prowler's Leather Boots (252468, -1.05 DPS, sim-verified) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.59 DPS) | yes | Blackstone Ring (17713, -0.16 DPS) [dungeon]; Mark of Kern (2262, -0.26 DPS) [dungeon]; Assault Band (13095, -0.26 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 23.9 attack_power points (1.58 DPS) | yes | Blackstone Ring (17713, -0.16 DPS) [dungeon]; Mark of Kern (2262, -0.26 DPS) [dungeon]; Assault Band (13095, -0.26 DPS) [world_drop] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+8.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -5.45 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Bloodrazor (809, -1.58 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (36.11 DPS) | yes | Doomforged Straightedge (12535, -1.56 DPS, sim-verified) [dungeon]; Claw of Celebras (17738, -4.16 DPS) [dungeon]; White Bone Shredder (11863, -6.56 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.13 DPS) [dungeon]; Dark Iron Rifle (16004, -3.25 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: White Bone Band; finger2: Legionnaire's Band; trinket1: Diamond Flask; trinket2: Rune of the Guard Captain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 35311103002000000-350511005050010051-000000000000000000)

Set DPS (verified): 291.9. Weights run: 2.9s. Verify run: 1.7s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.035, strength=2.062 ± 0.041, agility=0.592 ± 0.030, crit=0.846 ± 0.043 per rating point (14 rating = 1%, 11.839 per %), hit=0.205 ± 0.007 per rating point (10 rating = 1%, 2.052 per %), melee_haste=6.013 ± 0.553

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 64.9 attack_power points (4.33 DPS) | yes | Warlord's Plate Headpiece (231535, +0.00 DPS) [pvp]; Champion's Plate Helm (227043, -0.51 DPS) [pvp]; Crown of Heroism (226860, -12.68 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (291.9 DPS) | yes | Imperial Jewel (11933, -0.26 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.27 DPS) [quest]; Rage of Mugamba (19577, -5.25 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 47.2 attack_power points (3.15 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Champion's Plate Shoulders (227042, -0.02 DPS) [pvp]; Defiler's Leather Shoulders (20194, -0.44 DPS) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 37.0 attack_power points (2.47 DPS) | yes | Howler's Furs (272414, -0.46 DPS) [vendor]; Cape of the Black Baron (13340, -0.54 DPS) [dungeon]; Shroud of Domination (22337, -1.67 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (291.9 DPS) | yes | Obsidian Mail Tunic (22191, -0.18 DPS) [crafted]; Cadaverous Armor (14637, -0.62 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -17.40 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (291.9 DPS) | yes | General's Plate Armguards (16546, -0.31 DPS) [pvp]; Windtalker's Wristguards (19582, -0.39 DPS) [rep]; Bracers of Undead Slaying (23090, -7.82 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (291.9 DPS) | yes | Radiant Gloves of the Dawn (227817, -0.59 DPS) [vendor]; General's Plate Gauntlets (231532, -0.61 DPS) [pvp]; Razor Gauntlets (18326, -8.89 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 67.5 attack_power points (4.51 DPS) | yes | Ferocity of the Timbermaw (227805, -0.26 DPS) [vendor]; General's Plate Girdle (16547, -0.59 DPS) [pvp]; Dense Timbermaw Belt (227807, -1.83 DPS, sim-verified) [vendor] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (291.9 DPS) | yes | Titanic Leggings (22385, -0.24 DPS) [crafted]; Outrider's Plate Legguards (22651, -0.65 DPS) [rep]; Cloudkeeper Legplates (14554, -9.41 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 43.3 attack_power points (2.89 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Scalegut Treaders (275618, -0.04 DPS) [crafted] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (291.9 DPS) | yes | Band of the Ogre King (18522, -0.16 DPS) [dungeon]; Myrmidon's Signet (2246, -0.43 DPS) [world_drop]; Naglering (11669, -5.38 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (291.9 DPS) | yes | Band of the Ogre King (18522, -0.07 DPS) [dungeon]; Myrmidon's Signet (2246, -0.34 DPS) [world_drop]; Naglering (11669, -11.23 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (291.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -1.43 DPS) [dungeon]; Hand of Justice (11815, -1.56 DPS) [dungeon] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (291.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -1.96 DPS, sim-verified) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (291.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Felstriker (12590, -4.89 DPS, sim-verified) [dungeon] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (291.9 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -16.66 DPS) [dungeon]; Skullflame Shield (1168, -105.62 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (291.9 DPS) | yes | Bloodseeker (19107, -0.09 DPS) [quest]; Skull Splitting Crossbow (13039, -0.12 DPS) [world_drop]; Dark Iron Rifle (16004, -2.26 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Defiler's Plate Spaulders; back: Deathguard's Cloak; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Legionnaire's Band; finger2: Don Julio's Band; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Riphook

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

