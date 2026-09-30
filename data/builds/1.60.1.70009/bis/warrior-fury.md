# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 32.8. Weights run: 1.1s. Verify run: 1.1s. 303 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=0.170 ± 0.005 per rating point (14 rating = 1%, 2.375 per %), hit=0.048 ± 0.002 per rating point (10 rating = 1%, 0.483 per %), melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, +0.00 DPS, sim-verified) [quest] |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Rough Bronze Shoulders (3480, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, +0.00 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.13 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 attack_power points (0.73 DPS) | yes | Veteran's Chain Shirt (250488, -0.19 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.36 DPS) | yes | Cryptwalker Bracers (280095, -0.06 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.1 attack_power points (0.58 DPS) | yes | Gold-flecked Gloves (5195, -0.06 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -0.15 DPS) [quest]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Warchief's Girdle (5750, -0.29 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.1 attack_power points (0.80 DPS) | yes | Veteran's Chain Leggings (250493, -0.10 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.38 DPS) | yes | Veteran's Boots (250503, +0.00 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.31 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop]; Signet of the Zhevra (285330, -0.28 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Ring of the Moon (12052, -0.22 DPS) [world_drop]; The 1 Ring (8350, -0.22 DPS, sim-verified) [world]; Signet of the Zhevra (285330, -0.27 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.02 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 attack_power points (8.58 DPS) | yes | Diamond Hammer (2194, -0.46 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.14 DPS) [world_drop]; Bear Buckler (4821, -8.36 DPS) [vendor] |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567)) | Blacksmithing [crafted] | 4.0 attack_power points (0.15 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS, sim-verified) [quest]; Fine Longbow (11304, -0.00 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 303, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings

### Band 30 (human, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.6. Weights run: 1.2s. Verify run: 1.2s. 489 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=0.297 ± 0.022 per rating point (14 rating = 1%, 4.156 per %), hit=0.085 ± 0.004 per rating point (10 rating = 1%, 0.852 per %), melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 attack_power points (1.16 DPS) | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Veteran's Chain Helm (250498, -0.17 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | River Pride Choker (13087, -0.14 DPS) [world_drop]; Sentinel's Medallion (19541, -0.43 DPS) [rep]; Kaleidoscope Chain (13084, -0.44 DPS, sim-verified) [world_drop] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 attack_power points (0.62 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 10.9 attack_power points (0.38 DPS) | yes | Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Wolfmaster Cape (6314, -0.03 DPS) [dungeon]; Slayer's Cape (14752, -0.19 DPS, sim-verified) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (53.6 DPS) | yes | Barbaric Iron Breastplate (7914, -0.18 DPS) [crafted]; Hard Gold Cuirass (250533, -0.27 DPS) [crafted]; Avenger's Armor (1488, -2.12 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 attack_power points (0.71 DPS) | yes | Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Technician's Bracers (270042, -0.27 DPS) [quest]; Yorgen Bracers (13012, -0.34 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 23.6 attack_power points (0.83 DPS) | yes | Mail Combat Gauntlets (4075, -0.09 DPS) [world_drop]; Brawler Gloves (720, -0.12 DPS) [world_drop]; Bonefist Gauntlets (4465, -0.43 DPS, sim-verified) [world] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 30.3 attack_power points (1.07 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20108, -0.09 DPS) [rep]; Officer's Belt (250556, -0.13 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 attack_power points (1.03 DPS) | yes | Chausses of Westfall (6087, -0.05 DPS) [quest]; Slayer's Pants (14757, -0.05 DPS) [world_drop]; Golden Scale Leggings (3843, -0.30 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 attack_power points (0.67 DPS) | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop]; Hard Gold Boots (250534, -0.30 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 attack_power points (0.73 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Ironspine's Eye (7686, -0.31 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.4 attack_power points (0.58 DPS) | yes | Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Tiger Band (6749, -1.16 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 attack_power points (12.23 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 attack_power points (11.90 DPS) | yes | Shoni's Disarming Tool (9608, -4.00 DPS) [quest]; Shield of Thorsen (13079, -11.28 DPS) [world_drop]; Royal Diplomatic Scepter (9457, -13.47 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.32 DPS) | yes | Long Battle Bow (15284, -0.05 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.14 DPS) [crafted]; Double-barreled Shotgun (2098, -0.15 DPS, sim-verified) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 489, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 68.1. Weights run: 1.5s. Verify run: 1.3s. 673 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.333, strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=0.564 ± 0.039 per rating point (14 rating = 1%, 7.890 per %), hit=0.103 ± 0.005 per rating point (10 rating = 1%, 1.030 per %), melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.0 attack_power points (1.60 DPS) | yes | White Bandit Mask (10008, -0.25 DPS) [crafted]; Icemetal Barbute (10763, -0.29 DPS) [dungeon]; Chromite Barbute (8142, -0.61 DPS, sim-verified) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; Gazlowe's Charm (13088, -0.30 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.74 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 attack_power points (1.03 DPS) | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor]; Chromite Pauldrons (8144, -0.23 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 15.2 attack_power points (0.74 DPS) | yes | Sergeant Major's Cape (16315, -0.25 DPS) [pvp]; Hawkeye's Cloak (14593, -0.25 DPS) [world_drop]; Dark Hooded Cape (5257, -1.31 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 attack_power points (1.52 DPS) | yes | Avenger's Armor (1488, -0.11 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.21 DPS) [crafted]; Jouster's Chestplate (8157, -0.44 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.97 DPS) | yes | Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Yorgen Bracers (13012, -0.32 DPS) [world_drop]; Ravager's Armguards (14770, -0.51 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.55 DPS) | yes | Truesilver Gauntlets (7938, +0.00 DPS, sim-verified) [crafted]; Scarlet Gauntlets (10331, -0.13 DPS) [dungeon]; Gloves of Holy Might (867, -0.20 DPS) [world_drop] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.46 DPS) | yes | Highlander's Plate Girdle (20125, +0.00 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -0.05 DPS) [dungeon]; Highlander's Chain Girdle (20090, -0.29 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 attack_power points (1.97 DPS) | yes | Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.38 DPS) [world_drop]; Firemane Leggings (13129, -0.48 DPS, sim-verified) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 attack_power points (1.42 DPS) | yes | Blackforge Greaves (6423, -0.29 DPS) [world_drop]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.48 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 20.3 attack_power points (0.99 DPS) | yes | Mark of Kern (2262, -0.02 DPS) [dungeon]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Protector's Band (19517, -0.25 DPS) [rep] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.97 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor]; Mark of Kern (2262, -1.02 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.02 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (+7.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ardent Custodian (868, -7.21 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -10.54 DPS) [quest]; Savage Boar's Guard (10767, -20.41 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Monolithic Bow (9426, -0.03 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.19 DPS) [quest]; Bow of Searing Arrows (2825, -0.88 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 673, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 97.4. Weights run: 1.6s. Verify run: 1.5s. 865 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=0.333 ± 0.028 per rating point (14 rating = 1%, 4.668 per %), hit=0.100 ± 0.004 per rating point (10 rating = 1%, 0.999 per %), melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.7 attack_power points (4.23 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Bloomsprout Headpiece (17767, -0.82 DPS) [dungeon]; Fury Visor (20521, -1.49 DPS) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Skibi's Pendant (13089, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.59 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.07 DPS, sim-verified) [quest] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 23.4 attack_power points (2.22 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.49 DPS, sim-verified) [vendor]; Wyrmslayer Spaulders (13066, -0.15 DPS) [world_drop]; Earthslag Shoulders (11632, -0.27 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Sergeant Major's Cape (16336, -0.25 DPS) [pvp]; Wolfmaster Cape (6314, -0.40 DPS) [dungeon]; Blackveil Cape (11626, -1.61 DPS, sim-verified) [dungeon] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 38.0 attack_power points (3.60 DPS) | yes | Mixologist's Tunic (12793, -0.53 DPS) [dungeon]; Valorous Chestguard (8274, -0.60 DPS) [world_drop]; Grizzled Pelt (22274, -4.27 DPS, sim-verified) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.65 DPS) | yes | Branded Leather Bracers (19508, -0.76 DPS) [dungeon]; Officer's Wristguards (250581, -0.80 DPS) [crafted]; Runed Golem Shackles (12550, -2.70 DPS, sim-verified) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 38.4 attack_power points (3.63 DPS) | yes | Officer's Gloves (250551, -0.93 DPS) [crafted]; Rockgrip Gauntlets (17736, -0.98 DPS) [dungeon]; Gauntlets of Divinity (7724, -2.16 DPS, sim-verified) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Atal'alarion's Tusk Ring (10798, -0.14 DPS) [dungeon]; Belt of the Gladiator (13134, -0.14 DPS) [world_drop]; Girdle of Beastial Fury (11686, -2.39 DPS, sim-verified) [dungeon] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.8 attack_power points (3.30 DPS) | yes | Scarlet Leggings (10330, +0.07 DPS, sim-verified) [dungeon]; Silvershell Leggings (10633, -0.30 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.30 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Officer's Sabatons (250561, -0.07 DPS) [crafted]; Officer's Boots (250546, -0.13 DPS) [crafted]; Battlechaser's Greaves (12555, -2.68 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.99 DPS) | yes | Mark of Kern (2262, -0.09 DPS) [dungeon]; Protector's Band (19516, -0.19 DPS) [rep]; Protector's Band (19515, -0.52 DPS) [rep] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (1.89 DPS) | yes | Protector's Band (19516, -0.09 DPS) [rep]; Protector's Band (19515, -0.43 DPS) [rep]; Mark of Kern (2262, -2.12 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hammer of the Northern Wind (810, -8.50 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (51.63 DPS) | yes | Doomforged Straightedge (12535, +0.00 DPS, sim-verified) [dungeon]; Claw of Celebras (17738, -5.96 DPS) [dungeon]; Shoni's Disarming Tool (9608, -30.35 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.19 DPS) [dungeon]; Dark Iron Rifle (16004, -2.12 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Bracers of the Stone Princess; hands: Raider Gauntlets; legs: Golem Shard Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; trinket1: Frozen Heart of the Mountain; main_hand: Hanzo Sword; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 865, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 301.6. Weights run: 1.6s. Verify run: 1.7s. 1885 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.427, strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=0.933 ± 0.062 per rating point (14 rating = 1%, 13.060 per %), hit=0.227 ± 0.010 per rating point (10 rating = 1%, 2.273 per %), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lightbreaker Greathelm (239517) | Leonid Barthalomew the Revered [vendor] | 141.6 attack_power points (8.18 DPS) | yes | Field Marshal's Plate Helm (16478, -3.63 DPS) [vendor]; Field Marshal's Plate Helm (231538, -3.63 DPS) [pvp]; Lightbreaker Helmet (239525, -12.90 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 attack_power points (0.00 DPS) | yes | Imperial Jewel (11933, -0.29 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.29 DPS) [quest]; Rage of Mugamba (19577, -5.52 DPS, sim-verified) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 101.8 attack_power points (5.88 DPS) | yes | Lieutenant Commander's Plate Shoulders (23315, -2.82 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -2.82 DPS) [pvp]; Lightbreaker Pauldrons (239524, -14.22 DPS, sim-verified) [vendor] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | sim-verified (301.6 DPS) | yes | Shadewood Cloak (18328, -0.33 DPS) [dungeon]; Howler's Furs (272414, -0.35 DPS) [vendor]; Shroud of Domination (22337, -6.06 DPS, sim-verified) [dungeon] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Timbermaw Tunic (252484, -2.76 DPS) [crafted]; Cadaverous Armor (14637, -3.70 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -24.82 DPS, sim-verified) [world] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Berserker Bracers (19578, -1.54 DPS) [rep]; Marshal's Plate Bracers (16481, -1.83 DPS) [pvp]; Bracers of Undead Slaying (23090, -15.89 DPS, sim-verified) [world] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Raider Gauntlets (272095, -1.76 DPS) [vendor]; Marshal's Plate Gauntlets (16484, -2.24 DPS) [vendor]; Razor Gauntlets (18326, -15.74 DPS, sim-verified) [dungeon] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 101.9 attack_power points (5.89 DPS) | yes | Ferocity of the Timbermaw (227805, -1.93 DPS) [vendor]; Marshal's Plate Girdle (16482, -2.19 DPS) [pvp]; Radiant Girdle of the Dawn (227814, -6.52 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Sentinel's Plate Legguards (237825, -3.24 DPS) [vendor]; Titanic Leggings (22385, -3.46 DPS) [crafted]; Cloudkeeper Legplates (14554, -25.26 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 101.9 attack_power points (5.88 DPS) | yes | Marshal's Plate Boots (231539, -3.00 DPS) [pvp]; Marshal's Plate Boots (16483, -3.00 DPS) [vendor]; Boots of Heroism (21995, -9.75 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.27 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234026, -0.41 DPS) [vendor]; Naglering (11669, -11.72 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | 0.0 attack_power points (0.00 DPS) | yes | Band of the Ogre King (18522, -0.02 DPS) [dungeon]; Don Julio's Band (19325, -0.10 DPS) [rep]; Naglering (11669, -4.75 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | 0.0 attack_power points (0.00 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -6.14 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; The Lobotomizer (19324, -10.98 DPS, sim-verified) [rep] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 0.0 attack_power points (0.00 DPS) | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Skullflame Shield (1168, -93.60 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS, sim-verified) [crafted]; Bloodseeker (19107, -0.01 DPS) [quest]; Skull Splitting Crossbow (13039, -0.06 DPS) [world_drop] |

**New at 60:** head: Lightbreaker Greathelm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Cloak of the Honor Guard; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Ebon Hand; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1885, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (troll, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 29.7. Weights run: 1.1s. Verify run: 1.1s. 286 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=0.170 ± 0.005 per rating point (14 rating = 1%, 2.375 per %), hit=0.048 ± 0.002 per rating point (10 rating = 1%, 0.483 per %), melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.39 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, +0.00 DPS, sim-verified) [quest] |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Rough Bronze Shoulders (3480, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, +0.00 DPS, sim-verified) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.26 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.36 DPS) | yes | Bristlebark Bindings (14569, -0.21 DPS) [world_drop]; Runed Copper Bracers (2854, -0.22 DPS) [crafted]; Raptorcrest Bracers (270010, -0.39 DPS, sim-verified) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.1 attack_power points (0.58 DPS) | yes | Gold-flecked Gloves (5195, -0.11 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.21 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Ruffian Belt (5975, -0.22 DPS) [world]; Warchief's Girdle (5750, -0.29 DPS) [world]; Cobrahn's Grasp (6460, -0.32 DPS, sim-verified) [dungeon] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.6 attack_power points (0.67 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.38 DPS) | yes | Veteran's Boots (250503, +0.00 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.02 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 attack_power points (8.58 DPS) | yes | Diamond Hammer (2194, -0.16 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.14 DPS) [world_drop]; Bear Buckler (4821, -8.36 DPS) [vendor] |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.0 attack_power points (0.15 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.5. Weights run: 1.2s. Verify run: 1.1s. 472 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=0.297 ± 0.022 per rating point (14 rating = 1%, 4.156 per %), hit=0.085 ± 0.004 per rating point (10 rating = 1%, 0.852 per %), melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 attack_power points (1.16 DPS) | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted]; Veteran's Chain Helm (250498, -0.19 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | River Pride Choker (13087, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.33 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.43 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 attack_power points (0.62 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Slayer's Cape (14752) (or Lambent Scale Cloak (4706)) | World drop [world_drop] | 10.1 attack_power points (0.36 DPS) | yes | Lambent Scale Cloak (4706, +0.00 DPS, sim-verified) [world_drop]; Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (53.5 DPS) | yes | Barbaric Iron Breastplate (7914, -0.18 DPS) [crafted]; Hard Gold Cuirass (250533, -0.27 DPS) [crafted]; Avenger's Armor (1488, -1.67 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 attack_power points (0.71 DPS) | yes | Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.26 DPS) [quest]; Yorgen Bracers (13012, -0.43 DPS, sim-verified) [world_drop] |
| hands | Warsong Gauntlets (16978) | Warsong Supplies [quest] | 25.3 attack_power points (0.89 DPS) | yes | Gauntlets of Ogre Strength (3341, +0.29 DPS, sim-verified) [world]; Bonefist Gauntlets (4465, -0.09 DPS) [world]; Mail Combat Gauntlets (4075, -0.15 DPS) [world_drop] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 30.3 attack_power points (1.07 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Officer's Belt (250556, -0.13 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.22 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 attack_power points (1.03 DPS) | yes | Slayer's Pants (14757, -0.05 DPS) [world_drop]; Ferine Leggings (6690, -0.11 DPS) [dungeon]; Golden Scale Leggings (3843, -0.34 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 attack_power points (0.67 DPS) | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop]; Hard Gold Boots (250534, -0.34 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 attack_power points (0.73 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Ironspine's Eye (7686, -0.31 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.4 attack_power points (0.58 DPS) | yes | Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Tiger Band (6749, -1.35 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 attack_power points (12.23 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 attack_power points (11.90 DPS) | yes | Shield of Thorsen (13079, -11.28 DPS) [world_drop]; Slayer's Shield (15892, -11.34 DPS) [world_drop]; Royal Diplomatic Scepter (9457, -13.71 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.32 DPS) | yes | Long Battle Bow (15284, -0.05 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.14 DPS) [crafted]; Double-barreled Shotgun (2098, -0.33 DPS, sim-verified) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Slayer's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Warsong Gauntlets; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 472, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 60.6. Weights run: 1.5s. Verify run: 1.3s. 654 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.333, strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=0.564 ± 0.039 per rating point (14 rating = 1%, 7.890 per %), hit=0.103 ± 0.005 per rating point (10 rating = 1%, 1.030 per %), melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.0 attack_power points (1.60 DPS) | yes | White Bandit Mask (10008, -0.25 DPS) [crafted]; Icemetal Barbute (10763, -0.29 DPS) [dungeon]; Chromite Barbute (8142, -0.67 DPS, sim-verified) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ethereal Talisman (4430, -0.09 DPS) [quest]; Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.90 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 attack_power points (1.03 DPS) | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor]; Chromite Pauldrons (8144, -0.13 DPS, sim-verified) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest]; Dark Hooded Cape (5257, -0.73 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 attack_power points (1.52 DPS) | yes | Avenger's Armor (1488, -0.11 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.21 DPS) [crafted]; Jouster's Chestplate (8157, -0.39 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.97 DPS) | yes | Ravager's Armguards (14770, +0.09 DPS, sim-verified) [world_drop]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Darkspear Armsplints (4132, -0.31 DPS) [quest] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.55 DPS) | yes | Truesilver Gauntlets (7938, +0.00 DPS, sim-verified) [crafted]; Scarlet Gauntlets (10331, -0.13 DPS) [dungeon]; Gloves of Holy Might (867, -0.20 DPS) [world_drop] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.46 DPS) | yes | Defiler's Plate Girdle (20206, +0.00 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -0.05 DPS) [dungeon]; Tharg's Shoelace (9705, -0.24 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 attack_power points (1.97 DPS) | yes | Firemane Leggings (13129, -0.17 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.38 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 attack_power points (1.42 DPS) | yes | Prowler's Leather Shoes (252465, -0.17 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.29 DPS) [world_drop]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 20.3 attack_power points (0.99 DPS) | yes | Mark of Kern (2262, -0.02 DPS) [dungeon]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Legionnaire's Band (19513, -0.25 DPS) [rep] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.97 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor]; Mark of Kern (2262, -1.03 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.02 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 attack_power points (22.32 DPS) | yes | Vanquisher's Sword (10823, -1.31 DPS, sim-verified) [quest]; Savage Boar's Guard (10767, -21.29 DPS) [dungeon]; Skullance Shield (13081, -21.45 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Monolithic Bow (9426, -0.03 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.19 DPS) [quest]; Bow of Searing Arrows (2825, -0.79 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Hawkeye's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Ardent Custodian; ranged: The Silencer

No-known-source sample (15 of 654, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 93.2. Weights run: 1.6s. Verify run: 1.6s. 846 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=0.333 ± 0.028 per rating point (14 rating = 1%, 4.668 per %), hit=0.100 ± 0.004 per rating point (10 rating = 1%, 0.999 per %), melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.7 attack_power points (4.23 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Bloomsprout Headpiece (17767, -0.82 DPS) [dungeon]; Fury Visor (20521, -1.49 DPS) [quest] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.89 DPS) | yes | Ghostshard Talisman (7731, +0.67 DPS, sim-verified) [dungeon]; Woven Ivy Necklace (19159, -0.69 DPS) [quest]; Skibi's Pendant (13089, -0.71 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 23.4 attack_power points (2.22 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.21 DPS, sim-verified) [vendor]; Wyrmslayer Spaulders (13066, -0.15 DPS) [world_drop]; Earthslag Shoulders (11632, -0.27 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Battlehard Cape (11858, -0.40 DPS) [quest]; Wildhunter Cloak (16658, -0.40 DPS) [quest]; Blackveil Cape (11626, -0.87 DPS, sim-verified) [dungeon] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 38.0 attack_power points (3.60 DPS) | yes | Mixologist's Tunic (12793, -0.53 DPS) [dungeon]; Valorous Chestguard (8274, -0.60 DPS) [world_drop]; Grizzled Pelt (22274, -3.77 DPS, sim-verified) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.65 DPS) | yes | Branded Leather Bracers (19508, -0.76 DPS) [dungeon]; Officer's Wristguards (250581, -0.80 DPS) [crafted]; Runed Golem Shackles (12550, -2.84 DPS, sim-verified) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 38.4 attack_power points (3.63 DPS) | yes | Officer's Gloves (250551, -0.93 DPS) [crafted]; Rockgrip Gauntlets (17736, -0.98 DPS) [dungeon]; Gauntlets of Divinity (7724, -2.29 DPS, sim-verified) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Atal'alarion's Tusk Ring (10798, -0.14 DPS) [dungeon]; Belt of the Gladiator (13134, -0.14 DPS) [world_drop]; Girdle of Beastial Fury (11686, -2.14 DPS, sim-verified) [dungeon] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.8 attack_power points (3.30 DPS) | yes | Scarlet Leggings (10330, -0.16 DPS, sim-verified) [dungeon]; Silvershell Leggings (10633, -0.30 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.30 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Officer's Sabatons (250561, -0.07 DPS) [crafted]; Officer's Boots (250546, -0.13 DPS) [crafted]; Battlechaser's Greaves (12555, -2.57 DPS, sim-verified) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.27 DPS) | yes | Mark of Kern (2262, -0.38 DPS) [dungeon]; Assault Band (13095, -0.38 DPS) [world_drop]; Legionnaire's Band (19511, -0.47 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.99 DPS) | yes | Mark of Kern (2262, -0.09 DPS) [dungeon]; Legionnaire's Band (19511, -0.19 DPS) [rep]; Assault Band (13095, -1.45 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Blessed Prayer Beads (19990, -1.03 DPS, sim-verified) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hanzo Sword (8190, -3.36 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (51.63 DPS) | yes | Doomforged Straightedge (12535, +0.00 DPS, sim-verified) [dungeon]; Claw of Celebras (17738, -5.96 DPS) [dungeon]; White Bone Shredder (11863, -9.59 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.19 DPS) [dungeon]; Dark Iron Rifle (16004, -2.40 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Zealous Shadowshard Pendant; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Bracers of the Stone Princess; hands: Raider Gauntlets; legs: Golem Shard Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 846, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 301.0. Weights run: 1.6s. Verify run: 1.7s. 1866 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.427, strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=0.933 ± 0.062 per rating point (14 rating = 1%, 13.060 per %), hit=0.227 ± 0.010 per rating point (10 rating = 1%, 2.273 per %), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lightbreaker Greathelm (239517) | Leonid Barthalomew the Revered [vendor] | 141.6 attack_power points (8.18 DPS) | yes | Warlord's Plate Headpiece (16542, -3.63 DPS) [vendor]; Warlord's Plate Headpiece (231535, -3.63 DPS) [pvp]; Lightbreaker Helmet (239525, -13.81 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 attack_power points (0.00 DPS) | yes | Imperial Jewel (11933, -0.29 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.29 DPS) [quest]; Rage of Mugamba (19577, -5.46 DPS, sim-verified) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 101.8 attack_power points (5.88 DPS) | yes | Champion's Plate Shoulders (23243, -2.82 DPS) [vendor]; Champion's Plate Shoulders (227042, -2.82 DPS) [pvp]; Lightbreaker Pauldrons (239524, -14.72 DPS, sim-verified) [vendor] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | sim-verified (301.0 DPS) | yes | Shadewood Cloak (18328, -0.33 DPS) [dungeon]; Howler's Furs (272414, -0.35 DPS) [vendor]; Shroud of Domination (22337, -6.74 DPS, sim-verified) [dungeon] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Timbermaw Tunic (252484, -2.76 DPS) [crafted]; Cadaverous Armor (14637, -3.70 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -28.46 DPS, sim-verified) [world] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Berserker Bracers (19578, -1.54 DPS) [rep]; General's Plate Armguards (16546, -1.83 DPS) [pvp]; Bracers of Undead Slaying (23090, -14.49 DPS, sim-verified) [world] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Raider Gauntlets (272095, -1.76 DPS) [vendor]; General's Plate Gauntlets (16548, -2.24 DPS) [vendor]; Razor Gauntlets (18326, -16.70 DPS, sim-verified) [dungeon] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 101.9 attack_power points (5.89 DPS) | yes | Ferocity of the Timbermaw (227805, -1.93 DPS) [vendor]; General's Plate Girdle (16547, -2.19 DPS) [pvp]; Radiant Girdle of the Dawn (227814, -5.56 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Sentinel's Plate Legguards (237825, -3.24 DPS) [vendor]; Titanic Leggings (22385, -3.46 DPS) [crafted]; Cloudkeeper Legplates (14554, -26.10 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 101.9 attack_power points (5.88 DPS) | yes | General's Plate Boots (16545, -3.00 DPS) [vendor]; General's Plate Boots (231531, -3.00 DPS) [pvp]; Boots of Heroism (21995, -11.72 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.27 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234026, -0.41 DPS) [vendor]; Naglering (11669, -13.26 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | 0.0 attack_power points (0.00 DPS) | yes | Band of the Ogre King (18522, -0.02 DPS) [dungeon]; Don Julio's Band (19325, -0.10 DPS) [rep]; Naglering (11669, -8.43 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+8.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 attack_power points (0.00 DPS) | yes | Counterattack Lodestone (18537, -1.25 DPS) [dungeon]; Hand of Justice (11815, -1.36 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -4.53 DPS, sim-verified) [crafted] |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; The Lobotomizer (19324, -10.66 DPS, sim-verified) [rep] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 0.0 attack_power points (0.00 DPS) | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Skullflame Shield (1168, -82.65 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS, sim-verified) [crafted]; Bloodseeker (19107, -0.01 DPS) [quest]; Skull Splitting Crossbow (13039, -0.06 DPS) [world_drop] |

**New at 60:** head: Lightbreaker Greathelm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Deathguard's Cloak; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Legionnaire's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Ebon Hand; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1866, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

