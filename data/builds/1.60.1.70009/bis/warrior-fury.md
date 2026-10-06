# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 32.8. Weights run: 1.6s. Verify run: 1.0s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.026, strength=1.989 ± 0.033, agility=0.105 ± 0.011, crit=0.169 ± 0.004 per rating point (14 rating = 1%, 2.367 per %), hit=0.049 ± 0.001 per rating point (10 rating = 1%, 0.486 per %), melee_haste=1.719 ± 0.033

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 19.9 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.14 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.68 DPS) [crafted]; Brawler's Leather Hood (252504, -0.68 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.13 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 19.9 attack_power points (0.71 DPS) | yes | Veteran's Chain Shirt (250488, -0.19 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.20 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 9.9 attack_power points (0.36 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.20 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 15.9 attack_power points (0.57 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Polar Gauntlets (7606, -0.14 DPS) [quest]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.29 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 21.9 attack_power points (0.79 DPS) | yes | Veteran's Chain Leggings (250493, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.38 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.30 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop]; Signet of the Zhevra (285330, -0.28 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Ring of the Moon (12052, -0.21 DPS) [world_drop]; The 1 Ring (8350, -0.22 DPS, sim-verified) [world]; Signet of the Zhevra (285330, -0.26 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (8.93 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (8.49 DPS) | yes | Diamond Hammer (2194, -0.46 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.07 DPS) [world_drop]; Furen's Favor (6970, -8.28 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.14 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.7. Weights run: 1.8s. Verify run: 1.0s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.162, strength=2.313 ± 0.195, agility=not significant (0.231 ± 0.058), crit=0.314 ± 0.016 per rating point (14 rating = 1%, 4.401 per %), hit=0.087 ± 0.003 per rating point (10 rating = 1%, 0.873 per %), melee_haste=2.850 ± 0.251

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.1 attack_power points (1.05 DPS) | yes | Veteran's Chain Helm (250498, -0.08 DPS) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.16 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | River Pride Choker (13087, -0.17 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.35 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.42 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 16.2 attack_power points (0.57 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.03 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 10.2 attack_power points (0.36 DPS) | yes | Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Slayer's Cape (14752, -0.03 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 34.7 attack_power points (1.21 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted]; Hard Gold Cuirass (250533, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 18.5 attack_power points (0.65 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS) [world_drop]; Bands of Serra'kis (6902, -0.16 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.24 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.9 attack_power points (0.80 DPS) | yes | The Frozen Clutch (23170, -0.10 DPS) [dungeon]; Mail Combat Gauntlets (4075, -0.12 DPS) [world_drop]; Bonefist Gauntlets (4465, -0.36 DPS, sim-verified) [world] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 27.8 attack_power points (0.97 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS) [rep]; Highlander's Lamellar Girdle (20108, -0.08 DPS) [rep]; Officer's Belt (250556, -0.11 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 27.1 attack_power points (0.95 DPS) | yes | Ferine Leggings (6690, -0.04 DPS) [dungeon]; Golden Scale Leggings (3843, -0.06 DPS) [crafted]; Chausses of Westfall (6087, -0.06 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.8 attack_power points (0.62 DPS) | yes | Hard Gold Boots (250534, -0.06 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.2 attack_power points (0.67 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Silverlaine's Family Seal (6321, -0.27 DPS) [dungeon]; Ironspine's Eye (7686, -0.28 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.3 attack_power points (0.53 DPS) | yes | Tiger Band (6749, -0.05 DPS) [quest]; Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 347.0 attack_power points (12.14 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 337.1 attack_power points (11.79 DPS) | yes | Royal Diplomatic Scepter (9457, -0.85 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -3.93 DPS) [quest]; Shield of Thorsen (13079, -11.22 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.31 DPS) | yes | Double-barreled Shotgun (2098, -0.05 DPS) [world_drop]; Long Battle Bow (15284, -0.07 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.15 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 88.8. Weights run: 2.1s. Verify run: 1.0s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.248, strength=2.094 ± 0.316, agility=0.571 ± 0.116, crit=0.678 ± 0.033 per rating point (14 rating = 1%, 9.498 per %), hit=0.123 ± 0.004 per rating point (10 rating = 1%, 1.228 per %), melee_haste=4.539 ± 0.444

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 36.7 attack_power points (1.50 DPS) | yes | White Bandit Mask (10008, -0.30 DPS) [crafted]; Hard Gold Coif (250537, -0.30 DPS) [crafted]; Chromite Barbute (8142, -0.72 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.82 DPS) | yes | Kaleidoscope Chain (13084, -0.38 DPS) [world_drop]; River Pride Choker (13087, -0.47 DPS) [world_drop]; Ghostshard Talisman (7731, -0.52 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 23.0 attack_power points (0.94 DPS) | yes | Chromite Pauldrons (8144, -0.09 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 16.0 attack_power points (0.65 DPS) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Wolfmaster Cape (6314, -0.24 DPS) [dungeon]; Dark Hooded Cape (5257, -0.44 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 33.3 attack_power points (1.36 DPS) | yes | Avenger's Armor (1488, -0.08 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.16 DPS) [crafted]; Jouster's Chestplate (8157, -0.54 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Ravager's Armguards (14770, -0.12 DPS) [world_drop]; Pugilist Bracers (4438, -0.13 DPS) [dungeon]; Yorgen Bracers (13012, -0.23 DPS) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 33.5 attack_power points (1.37 DPS) | yes | Gauntlets of Divinity (7724, -0.06 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.11 DPS) [dungeon]; Gloves of Holy Might (867, -0.16 DPS) [world_drop] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 31.4 attack_power points (1.28 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.06 DPS) [rep]; Scarlet Belt (10329, -0.26 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 44.0 attack_power points (1.80 DPS) | yes | Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.37 DPS) [world_drop]; Firemane Leggings (13129, -0.57 DPS, sim-verified) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 31.2 attack_power points (1.27 DPS) | yes | Blackforge Greaves (6423, -0.27 DPS) [dungeon]; Obsidian Greaves (13068, -0.28 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.57 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.3 attack_power points (0.87 DPS) | yes | Assault Band (13095, -0.05 DPS) [world_drop]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Suspicious Spare Part (274754, -0.27 DPS) [vendor] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Suspicious Spare Part (274754, -0.22 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (19.37 DPS) | yes | Ardent Custodian (868, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (88.8 DPS) | yes | Shoni's Disarming Tool (9608, -8.87 DPS) [quest]; Savage Boar's Guard (10767, -17.11 DPS) [dungeon]; Ardent Custodian (868, -27.14 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -0.01 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.15 DPS) [quest]; Bow of Searing Arrows (2825, -1.33 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 153.9. Weights run: 2.2s. Verify run: 1.4s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.208, strength=1.830 ± 0.271, agility=not significant (0.437 ± 0.118), crit=0.444 ± 0.026 per rating point (14 rating = 1%, 6.213 per %), hit=0.127 ± 0.004 per rating point (10 rating = 1%, 1.275 per %), melee_haste=3.441 ± 0.365

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 46.6 attack_power points (3.33 DPS) | yes | Fury Visor (20521, -0.94 DPS) [quest]; Sunscale Helmet (14849, -1.05 DPS) [world_drop]; Bloomsprout Headpiece (17767, -1.80 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.43 DPS) | yes | Skibi's Pendant (13089, -0.37 DPS) [world_drop]; Ghostshard Talisman (7731, -0.43 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.78 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 27.3 attack_power points (1.95 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS, sim-verified) [vendor]; Wyrmslayer Spaulders (13066, -0.13 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.23 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.1 attack_power points (1.22 DPS) | yes | Sergeant Major's Cape (16336, -0.25 DPS) [pvp]; Dark Hooded Cape (5257, -0.39 DPS) [world]; Bloodlust Cape (14801, -1.93 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 43.9 attack_power points (3.13 DPS) | yes | Valorous Chestguard (8274, -0.52 DPS) [world_drop]; Coldmetal Guard (274758, -0.65 DPS) [vendor]; Mixologist's Tunic (12793, -1.33 DPS, sim-verified) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.00 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Runed Golem Shackles (12550, -0.17 DPS) [dungeon]; Officer's Wristguards (250581, -0.37 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 44.6 attack_power points (3.19 DPS) | yes | Gauntlets of Divinity (7724, -0.90 DPS) [dungeon]; Maddening Gauntlets (11867, -1.07 DPS) [quest]; Officer's Gloves (250551, -2.13 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 44.6 attack_power points (3.19 DPS) | yes | Belt of the Gladiator (13134, -0.84 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.98 DPS) [crafted]; Atal'alarion's Tusk Ring (10798, -1.77 DPS, sim-verified) [dungeon] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 40.3 attack_power points (2.87 DPS) | yes | Silvershell Leggings (10633, -0.26 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.26 DPS) [dungeon]; Scarlet Leggings (10330, -1.34 DPS, sim-verified) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 31.3 attack_power points (2.23 DPS) | yes | Officer's Sabatons (250561, -0.26 DPS) [crafted]; Officer's Boots (250546, -0.32 DPS) [crafted]; Prowler's Leather Boots (252468, -1.26 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.2 attack_power points (1.59 DPS) | yes | Mark of Kern (2262, -0.16 DPS) [dungeon]; Assault Band (13095, -0.16 DPS) [world_drop]; Thunderbrow Ring (13097, -0.45 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.3 attack_power points (1.52 DPS) | yes | Assault Band (13095, -0.09 DPS) [world_drop]; Thunderbrow Ring (13097, -0.38 DPS) [world_drop]; Mark of Kern (2262, -3.85 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -6.53 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smoking Heart of the Mountain (11811, -3.12 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Hanzo Sword (8190, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Doomforged Straightedge (12535) | Blackrock Depths: Anvilrage Overseer [dungeon] | sim-verified (153.9 DPS) | yes | Claw of Celebras (17738, -3.51 DPS) [dungeon]; Shadowblade (2163, -7.26 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -21.89 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.14 DPS) [dungeon]; Dark Iron Rifle (16004, -2.22 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Bloodrazor; off_hand: Doomforged Straightedge; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 259.3. Weights run: 2.4s. Verify run: 1.4s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.446, strength=3.029 ± 0.608, agility=not significant (0.744 ± 0.233), crit=1.407 ± 0.064 per rating point (14 rating = 1%, 19.691 per %), hit=0.323 ± 0.011 per rating point (10 rating = 1%, 3.229 per %), melee_haste=10.549 ± 0.889

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 100.4 attack_power points (3.99 DPS) | yes | Field Marshal's Plate Helm (231538, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -0.55 DPS) [vendor]; Crown of Heroism (226860, -9.52 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-verified (259.3 DPS) | yes | Medallion of the Dawn (22659, -0.03 DPS) [quest]; Imperial Jewel (11933, -0.49 DPS) [dungeon]; Rage of Mugamba (19577, -3.21 DPS, sim-verified) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 67.2 attack_power points (2.67 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 51.5 attack_power points (2.05 DPS) | yes | Cloak of Revanchion (23127, -0.55 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.55 DPS) [rep]; Shadewood Cloak (18328, -1.50 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (259.3 DPS) | yes | Obsidian Mail Tunic (22191, -1.48 DPS) [crafted]; Cadaverous Armor (14637, -1.70 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -16.07 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-verified (259.3 DPS) | yes | Marshal's Plate Bracers (16481, -0.27 DPS) [pvp]; Gordok Bracers of Power (18533, -0.48 DPS) [dungeon]; Bracers of Undead Slaying (23090, -6.87 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (259.3 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.39 DPS) [pvp]; Radiant Gloves of the Dawn (227817, -0.51 DPS) [vendor]; Razor Gauntlets (18326, -7.68 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 101.5 attack_power points (4.04 DPS) | yes | Marshal's Plate Girdle (16482, -0.67 DPS) [pvp]; Heavy Obsidian Belt (22197, -1.02 DPS) [crafted]; Ferocity of the Timbermaw (227805, -1.48 DPS, sim-verified) [vendor] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (259.3 DPS) | yes | Titanic Leggings (22385, -0.29 DPS) [crafted]; Marshal's Plate Legguards (231540, -0.84 DPS) [pvp]; Cloudkeeper Legplates (14554, -5.58 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 63.8 attack_power points (2.54 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Clutchlord's Stompers (275627, -0.25 DPS) [crafted] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (259.3 DPS) | yes | Don Julio's Band (19325, -0.22 DPS) [rep]; Myrmidon's Signet (2246, -0.36 DPS) [world_drop]; Naglering (11669, -4.49 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-verified (259.3 DPS) | yes | Don Julio's Band (19325, -0.14 DPS) [rep]; Myrmidon's Signet (2246, -0.27 DPS) [world_drop]; Naglering (11669, -3.27 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (259.3 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Diamond Flask (20130, +0.00 DPS, sim-verified) [quest] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (259.3 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Diamond Flask (20130, +0.00 DPS) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (259.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Quel'Serrar (18348, -9.86 DPS, sim-verified) [quest] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (259.3 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -10.35 DPS) [dungeon]; Skullflame Shield (1168, -64.33 DPS, sim-verified) [world_drop] |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-verified (259.3 DPS) | yes | Unsophisticated Hand Cannon (18460, -0.21 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.25 DPS) [world_drop]; Dark Iron Rifle (16004, -2.09 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Amulet of the Darkmoon; shoulder: Highlander's Plate Spaulders; back: Shroud of Domination; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Protector's Band; finger2: Band of the Ogre King; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Bloodseeker

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (troll, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 29.7. Weights run: 1.6s. Verify run: 1.0s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.026, strength=1.989 ± 0.033, agility=0.105 ± 0.011, crit=0.169 ± 0.004 per rating point (14 rating = 1%, 2.367 per %), hit=0.049 ± 0.001 per rating point (10 rating = 1%, 0.486 per %), melee_haste=1.719 ± 0.033

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 19.9 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.39 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.68 DPS) [crafted]; Brawler's Leather Hood (252504, -0.68 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 19.9 attack_power points (0.71 DPS) | yes | Defender's Leather Armor (252434, -0.20 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.26 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 9.9 attack_power points (0.36 DPS) | yes | Bristlebark Bindings (14569, -0.20 DPS) [world_drop]; Runed Copper Bracers (2854, -0.21 DPS) [crafted]; Raptorcrest Bracers (270010, -0.39 DPS, sim-verified) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 15.9 attack_power points (0.57 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.20 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.29 DPS) [world_drop]; Cobrahn's Grasp (6460, -0.32 DPS, sim-verified) [dungeon] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.4 attack_power points (0.66 DPS) | yes | Defender's Leather Pants (252445, -0.00 DPS) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.38 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (8.93 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (8.49 DPS) | yes | Diamond Hammer (2194, +0.00 DPS) [world_drop]; Redbeard Crest (12997, -8.07 DPS) [world_drop]; Ruga's Bulwark (7120, -8.28 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.14 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.8. Weights run: 1.8s. Verify run: 1.0s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.162, strength=2.313 ± 0.195, agility=not significant (0.231 ± 0.058), crit=0.314 ± 0.016 per rating point (14 rating = 1%, 4.401 per %), hit=0.087 ± 0.003 per rating point (10 rating = 1%, 0.873 per %), melee_haste=2.850 ± 0.251

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.1 attack_power points (1.05 DPS) | yes | Veteran's Chain Helm (250498, -0.08 DPS) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.16 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; River Pride Choker (13087, -0.17 DPS) [world_drop]; Scout's Medallion (19537, -0.42 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 16.2 attack_power points (0.57 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.03 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.35 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Slayer's Cape (14752, -0.03 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 34.7 attack_power points (1.21 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted]; Hard Gold Cuirass (250533, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 18.5 attack_power points (0.65 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS) [world_drop]; Bands of Serra'kis (6902, -0.16 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.23 DPS) [quest] |
| hands | Warsong Gauntlets (16978) | Warsong Supplies [quest] | 23.1 attack_power points (0.81 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.01 DPS) [world]; Bonefist Gauntlets (4465, -0.08 DPS) [world]; The Frozen Clutch (23170, -0.11 DPS) [dungeon] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 27.8 attack_power points (0.97 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS) [rep]; Officer's Belt (250556, -0.11 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.13 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 27.1 attack_power points (0.95 DPS) | yes | Ferine Leggings (6690, -0.04 DPS) [dungeon]; Golden Scale Leggings (3843, -0.06 DPS) [crafted]; Slayer's Pants (14757, -0.06 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.8 attack_power points (0.62 DPS) | yes | Hard Gold Boots (250534, -0.06 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.2 attack_power points (0.67 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Silverlaine's Family Seal (6321, -0.27 DPS) [dungeon]; Ironspine's Eye (7686, -0.28 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.3 attack_power points (0.53 DPS) | yes | Tiger Band (6749, -0.05 DPS) [quest]; Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 347.0 attack_power points (12.14 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 337.1 attack_power points (11.79 DPS) | yes | Royal Diplomatic Scepter (9457, -0.96 DPS, sim-verified) [dungeon]; Shield of Thorsen (13079, -11.22 DPS) [world_drop]; Slayer's Shield (15892, -11.28 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.31 DPS) | yes | Double-barreled Shotgun (2098, -0.05 DPS) [world_drop]; Long Battle Bow (15284, -0.07 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.15 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Warsong Gauntlets; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 84.3. Weights run: 2.1s. Verify run: 1.1s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.248, strength=2.094 ± 0.316, agility=0.571 ± 0.116, crit=0.678 ± 0.033 per rating point (14 rating = 1%, 9.498 per %), hit=0.123 ± 0.004 per rating point (10 rating = 1%, 1.228 per %), melee_haste=4.539 ± 0.444

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 36.7 attack_power points (1.50 DPS) | yes | White Bandit Mask (10008, -0.30 DPS) [crafted]; Hard Gold Coif (250537, -0.30 DPS) [crafted]; Chromite Barbute (8142, -0.48 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.82 DPS) | yes | Ghostshard Talisman (7731, -0.24 DPS) [dungeon]; Ethereal Talisman (4430, -0.30 DPS) [quest]; Kaleidoscope Chain (13084, -0.38 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 23.0 attack_power points (0.94 DPS) | yes | Chromite Pauldrons (8144, -0.09 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 16.0 attack_power points (0.65 DPS) | yes | Dark Hooded Cape (5257, -0.08 DPS) [world]; Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Wildhunter Cloak (16658, -0.24 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 33.3 attack_power points (1.36 DPS) | yes | Avenger's Armor (1488, -0.08 DPS) [dungeon]; Jouster's Chestplate (8157, -0.08 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.16 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Ravager's Armguards (14770, +0.00 DPS) [world_drop]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 33.5 attack_power points (1.37 DPS) | yes | Gauntlets of Divinity (7724, -0.06 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.11 DPS) [dungeon]; Gloves of Holy Might (867, -0.16 DPS) [world_drop] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 31.4 attack_power points (1.28 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.06 DPS) [rep]; Tharg's Shoelace (9705, -0.17 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 44.0 attack_power points (1.80 DPS) | yes | Firemane Leggings (13129, -0.17 DPS) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.37 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 31.2 attack_power points (1.27 DPS) | yes | Prowler's Leather Shoes (252465, -0.17 DPS) [crafted]; Blackforge Greaves (6423, -0.27 DPS) [dungeon]; Obsidian Greaves (13068, -0.28 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.3 attack_power points (0.87 DPS) | yes | Assault Band (13095, -0.05 DPS) [world_drop]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Suspicious Spare Part (274754, -0.27 DPS) [vendor] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Suspicious Spare Part (274754, -0.22 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (19.37 DPS) | yes | Ardent Custodian (868, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (84.3 DPS) | yes | Savage Boar's Guard (10767, -17.11 DPS) [dungeon]; Skullance Shield (13081, -17.27 DPS) [world_drop]; Ardent Custodian (868, -23.11 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -0.01 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.15 DPS) [quest]; Bow of Searing Arrows (2825, -0.69 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 137.7. Weights run: 2.2s. Verify run: 1.4s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.208, strength=1.830 ± 0.271, agility=not significant (0.437 ± 0.118), crit=0.444 ± 0.026 per rating point (14 rating = 1%, 6.213 per %), hit=0.127 ± 0.004 per rating point (10 rating = 1%, 1.275 per %), melee_haste=3.441 ± 0.365

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 46.6 attack_power points (3.33 DPS) | yes | Fury Visor (20521, -0.94 DPS) [quest]; Sunscale Helmet (14849, -1.05 DPS) [world_drop]; Bloomsprout Headpiece (17767, -1.16 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.43 DPS) | yes | Woven Ivy Necklace (19159, -0.36 DPS) [quest]; Skibi's Pendant (13089, -0.37 DPS) [world_drop]; Ghostshard Talisman (7731, -0.43 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 27.3 attack_power points (1.95 DPS) | yes | Blood Guard's Plate Pauldrons (220796, -0.07 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.13 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.23 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.1 attack_power points (1.22 DPS) | yes | First Sergeant's Cloak (16340, -0.25 DPS) [pvp]; Dark Hooded Cape (5257, -0.39 DPS) [world]; Bloodlust Cape (14801, -1.16 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 43.9 attack_power points (3.13 DPS) | yes | Mixologist's Tunic (12793, -0.44 DPS) [dungeon]; Valorous Chestguard (8274, -0.52 DPS) [world_drop]; Coldmetal Guard (274758, -0.65 DPS) [vendor] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.00 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 44.6 attack_power points (3.19 DPS) | yes | Gauntlets of Divinity (7724, -0.90 DPS) [dungeon]; Prowler's Leather Gauntlets (252547, -1.08 DPS) [crafted]; Officer's Gloves (250551, -1.15 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 44.6 attack_power points (3.19 DPS) | yes | Belt of the Gladiator (13134, -0.84 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.98 DPS) [crafted]; Atal'alarion's Tusk Ring (10798, -1.15 DPS, sim-verified) [dungeon] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 40.3 attack_power points (2.87 DPS) | yes | Scarlet Leggings (10330, -0.13 DPS) [dungeon]; Silvershell Leggings (10633, -0.26 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.26 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 31.3 attack_power points (2.23 DPS) | yes | Officer's Sabatons (250561, -0.26 DPS) [crafted]; Officer's Boots (250546, -0.32 DPS) [crafted]; Prowler's Leather Boots (252468, -1.06 DPS, sim-verified) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.71 DPS) | yes | Blackstone Ring (17713, -0.19 DPS) [dungeon]; Mark of Kern (2262, -0.29 DPS) [dungeon]; Assault Band (13095, -0.29 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 22.2 attack_power points (1.59 DPS) | yes | Blackstone Ring (17713, +0.00 DPS, sim-verified) [dungeon]; Mark of Kern (2262, -0.16 DPS) [dungeon]; Assault Band (13095, -0.16 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (137.7 DPS) | yes | Frozen Heart of the Mountain (249469, -2.98 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (137.7 DPS) | yes | Frozen Heart of the Mountain (249469, -4.50 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (137.7 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Hanzo Sword (8190, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (38.90 DPS) | yes | Doomforged Straightedge (12535, +0.00 DPS) [dungeon]; Claw of Celebras (17738, -4.49 DPS) [dungeon]; White Bone Shredder (11863, -7.11 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (137.7 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.14 DPS) [dungeon]; Dark Iron Rifle (16004, -1.70 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: White Bone Band; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 243.3. Weights run: 2.4s. Verify run: 1.5s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.446, strength=3.029 ± 0.608, agility=not significant (0.744 ± 0.233), crit=1.407 ± 0.064 per rating point (14 rating = 1%, 19.691 per %), hit=0.323 ± 0.011 per rating point (10 rating = 1%, 3.229 per %), melee_haste=10.549 ± 0.889

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 100.4 attack_power points (3.99 DPS) | yes | Warlord's Plate Headpiece (231535, +0.00 DPS) [pvp]; Champion's Plate Helm (227043, -0.55 DPS) [pvp]; Crown of Heroism (226860, -11.28 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-verified (243.3 DPS) | yes | Medallion of the Dawn (22659, -0.03 DPS) [quest]; Conqueror's Medallion (12059, -0.44 DPS) [quest]; Rage of Mugamba (19577, -1.85 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 67.2 attack_power points (2.67 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Spaulders (272108, -0.44 DPS) [vendor] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 51.5 attack_power points (2.05 DPS) | yes | Shadewood Cloak (18328, -0.48 DPS) [dungeon]; Cloak of Revanchion (23127, -0.55 DPS) [dungeon]; Deathguard's Cloak (20068, -0.55 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (243.3 DPS) | yes | Obsidian Mail Tunic (22191, -1.48 DPS) [crafted]; Cadaverous Armor (14637, -1.70 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -16.96 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (243.3 DPS) | yes | General's Plate Armguards (16546, -0.27 DPS) [pvp]; Gordok Bracers of Power (18533, -0.48 DPS) [dungeon]; Bracers of Undead Slaying (23090, -8.59 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (243.3 DPS) | yes | General's Plate Gauntlets (231532, -0.39 DPS) [pvp]; Radiant Gloves of the Dawn (227817, -0.51 DPS) [vendor]; Razor Gauntlets (18326, -8.83 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 101.5 attack_power points (4.04 DPS) | yes | Ferocity of the Timbermaw (227805, -0.40 DPS) [vendor]; General's Plate Girdle (16547, -0.67 DPS) [pvp]; Heavy Obsidian Belt (22197, -1.02 DPS) [crafted] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (243.3 DPS) | yes | Titanic Leggings (22385, -0.29 DPS) [crafted]; Outrider's Plate Legguards (22651, -0.66 DPS) [rep]; Cloudkeeper Legplates (14554, -8.49 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 63.8 attack_power points (2.54 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Clutchlord's Stompers (275627, -0.25 DPS) [crafted] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (243.3 DPS) | yes | Don Julio's Band (19325, -0.22 DPS) [rep]; Myrmidon's Signet (2246, -0.36 DPS) [world_drop]; Naglering (11669, -5.95 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-verified (243.3 DPS) | yes | Don Julio's Band (19325, -0.14 DPS) [rep]; Myrmidon's Signet (2246, -0.27 DPS) [world_drop]; Naglering (11669, -5.46 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (243.3 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (243.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (243.3 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Felstriker (12590, -4.54 DPS, sim-verified) [dungeon] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (243.3 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -10.35 DPS) [dungeon]; Skullflame Shield (1168, -65.51 DPS, sim-verified) [world_drop] |
| ranged | Bloodseeker (19107) | The Legend of Korrak [quest] | sim-verified (243.3 DPS) | yes | Unsophisticated Hand Cannon (18460, -0.21 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.25 DPS) [world_drop]; Dark Iron Rifle (16004, -2.99 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Amulet of the Darkmoon; shoulder: Defiler's Plate Spaulders; back: Shroud of Domination; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Legionnaire's Band; finger2: Band of the Ogre King; trinket1: Darkmoon Card: Maelstrom; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Bloodseeker

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

