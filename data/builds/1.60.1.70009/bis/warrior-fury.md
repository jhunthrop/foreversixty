# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 32.8. Weights run: 1.2s. Verify run: 1.1s. 293 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=0.170 ± 0.005 per rating point (14 rating = 1%, 2.375 per %), hit=0.048 ± 0.002 per rating point (10 rating = 1%, 0.483 per %), melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, +0.00 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, +0.00 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.13 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 attack_power points (0.73 DPS) | yes | Veteran's Chain Shirt (250488, -0.19 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.36 DPS) | yes | Cryptwalker Bracers (280095, -0.06 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.1 attack_power points (0.58 DPS) | yes | Gold-flecked Gloves (5195, -0.06 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -0.15 DPS) [quest]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.29 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.1 attack_power points (0.80 DPS) | yes | Veteran's Chain Leggings (250493, -0.10 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.38 DPS) | yes | Veteran's Boots (250503, +0.00 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.31 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop]; Signet of the Zhevra (285330, -0.28 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Ring of the Moon (12052, -0.22 DPS) [world_drop]; The 1 Ring (8350, -0.22 DPS, sim-verified) [world]; Signet of the Zhevra (285330, -0.27 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.02 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 attack_power points (8.58 DPS) | yes | Diamond Hammer (2194, -0.46 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.14 DPS) [world_drop]; Furen's Favor (6970, -8.36 DPS) [quest] |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567)) | Blacksmithing [crafted] | 4.0 attack_power points (0.15 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS, sim-verified) [quest]; Fine Longbow (11304, -0.00 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 293, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings

### Band 30 (human, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.7. Weights run: 1.2s. Verify run: 1.1s. 479 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=0.297 ± 0.022 per rating point (14 rating = 1%, 4.156 per %), hit=0.085 ± 0.004 per rating point (10 rating = 1%, 0.852 per %), melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 attack_power points (1.16 DPS) | yes | Veteran's Chain Helm (250498, -0.06 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | River Pride Choker (13087, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.35 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.43 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 attack_power points (0.62 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 10.9 attack_power points (0.38 DPS) | yes | Slayer's Cape (14752, +0.00 DPS, sim-verified) [world_drop]; Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Wolfmaster Cape (6314, -0.03 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 37.9 attack_power points (1.33 DPS) | yes | Shining Silver Breastplate (2870, -0.06 DPS, sim-verified) [crafted]; Barbaric Iron Breastplate (7914, -0.27 DPS) [crafted]; Hard Gold Cuirass (250533, -0.36 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 attack_power points (0.71 DPS) | yes | Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.27 DPS) [crafted]; Yorgen Bracers (13012, -0.27 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 23.6 attack_power points (0.83 DPS) | yes | Mail Combat Gauntlets (4075, -0.09 DPS) [world_drop]; Brawler Gloves (720, -0.12 DPS) [world_drop]; Bonefist Gauntlets (4465, -0.36 DPS, sim-verified) [world] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 30.3 attack_power points (1.07 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20108, -0.09 DPS) [rep]; Officer's Belt (250556, -0.13 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 attack_power points (1.03 DPS) | yes | Chausses of Westfall (6087, -0.05 DPS) [quest]; Slayer's Pants (14757, -0.05 DPS) [world_drop]; Golden Scale Leggings (3843, -0.18 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 attack_power points (0.67 DPS) | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop]; Hard Gold Boots (250534, -0.18 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 attack_power points (0.73 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Ironspine's Eye (7686, -0.31 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.4 attack_power points (0.58 DPS) | yes | Tiger Band (6749, -0.11 DPS, sim-verified) [quest]; Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.16 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 attack_power points (12.23 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 attack_power points (11.90 DPS) | yes | Royal Diplomatic Scepter (9457, -0.85 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -4.00 DPS) [quest]; Shield of Thorsen (13079, -11.28 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.32 DPS) | yes | Long Battle Bow (15284, -0.05 DPS) [world_drop]; Double-barreled Shotgun (2098, -0.13 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -0.14 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 479, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 89.0. Weights run: 1.4s. Verify run: 1.3s. 671 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.333, strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=0.564 ± 0.039 per rating point (14 rating = 1%, 7.890 per %), hit=0.103 ± 0.005 per rating point (10 rating = 1%, 1.030 per %), melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.0 attack_power points (1.60 DPS) | yes | White Bandit Mask (10008, -0.25 DPS) [crafted]; Hard Gold Coif (250537, -0.29 DPS) [crafted]; Chromite Barbute (8142, -0.42 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.97 DPS) | yes | Ghostshard Talisman (7731, -0.46 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop]; River Pride Choker (13087, -0.60 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 attack_power points (1.03 DPS) | yes | Chromite Pauldrons (8144, +0.00 DPS, sim-verified) [dungeon]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 15.2 attack_power points (0.74 DPS) | yes | Dark Hooded Cape (5257, +0.00 DPS, sim-verified) [world]; Hawkeye's Cloak (14593, -0.25 DPS) [world_drop]; Wolfmaster Cape (6314, -0.25 DPS) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 attack_power points (1.52 DPS) | yes | Avenger's Armor (1488, -0.11 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.21 DPS) [crafted]; Jouster's Chestplate (8157, -0.42 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.97 DPS) | yes | Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Ravager's Armguards (14770, -0.26 DPS, sim-verified) [world_drop]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.55 DPS) | yes | Truesilver Gauntlets (7938, +0.00 DPS, sim-verified) [crafted]; Scarlet Gauntlets (10331, -0.13 DPS) [dungeon]; Gloves of Holy Might (867, -0.20 DPS) [world_drop] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.46 DPS) | yes | Highlander's Plate Girdle (20125, +0.00 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -0.05 DPS) [dungeon]; Highlander's Chain Girdle (20090, -0.29 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 attack_power points (1.97 DPS) | yes | Firemane Leggings (13129, -0.05 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.38 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 attack_power points (1.42 DPS) | yes | Prowler's Leather Shoes (252465, -0.05 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.29 DPS) [dungeon]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 20.3 attack_power points (0.99 DPS) | yes | Assault Band (13095, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Suspicious Spare Part (274754, -0.33 DPS) [vendor] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.97 DPS) | yes | Assault Band (13095, +0.00 DPS, sim-verified) [world_drop]; Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.02 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (89.0 DPS) | yes | Shoni's Disarming Tool (9608, -10.54 DPS) [quest]; Savage Boar's Guard (10767, -20.41 DPS) [dungeon]; Ardent Custodian (868, -27.61 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.03 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.19 DPS) [quest]; Bow of Searing Arrows (2825, -1.01 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 671, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 152.6. Weights run: 1.6s. Verify run: 1.8s. 856 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=0.333 ± 0.028 per rating point (14 rating = 1%, 4.668 per %), hit=0.100 ± 0.004 per rating point (10 rating = 1%, 0.999 per %), melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.7 attack_power points (4.23 DPS) | yes | Fury Visor (20521, -1.49 DPS) [quest]; Sunscale Helmet (14849, -1.63 DPS) [world_drop]; Bloomsprout Headpiece (17767, -2.01 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.89 DPS) | yes | Skibi's Pendant (13089, -0.71 DPS) [world_drop]; Ghostshard Talisman (7731, -1.07 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -1.16 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 23.4 attack_power points (2.22 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS, sim-verified) [vendor]; Wyrmslayer Spaulders (13066, -0.15 DPS) [world_drop]; Earthslag Shoulders (11632, -0.27 DPS) [dungeon] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.4 attack_power points (1.37 DPS) | yes | Sergeant Major's Cape (16336, -0.27 DPS) [pvp]; Wolfmaster Cape (6314, -0.42 DPS) [dungeon]; Bloodlust Cape (14801, -1.60 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 38.0 attack_power points (3.60 DPS) | yes | Mixologist's Tunic (12793, -0.31 DPS, sim-verified) [dungeon]; Valorous Chestguard (8274, -0.60 DPS) [world_drop]; Coldmetal Guard (274758, -0.75 DPS) [vendor] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.65 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Runed Golem Shackles (12550, -0.55 DPS) [dungeon]; Branded Leather Bracers (19508, -0.76 DPS) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 38.4 attack_power points (3.63 DPS) | yes | Officer's Gloves (250551, -0.93 DPS) [crafted]; Rockgrip Gauntlets (17736, -0.98 DPS) [dungeon]; Gauntlets of Divinity (7724, -3.61 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 42.7 attack_power points (4.04 DPS) | yes | Atal'alarion's Tusk Ring (10798, -1.34 DPS) [dungeon]; Belt of the Gladiator (13134, -1.34 DPS) [world_drop]; Highlander's Leather Girdle (20116, -2.87 DPS, sim-verified) [rep] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.8 attack_power points (3.30 DPS) | yes | Silvershell Leggings (10633, -0.30 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.30 DPS) [dungeon]; Scarlet Leggings (10330, -0.73 DPS, sim-verified) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 26.7 attack_power points (2.53 DPS) | yes | Officer's Sabatons (250561, -0.28 DPS) [crafted]; Officer's Boots (250546, -0.35 DPS) [crafted]; Prowler's Leather Boots (252468, -1.11 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.99 DPS) | yes | Assault Band (13095, -0.09 DPS) [world_drop]; Protector's Band (19516, -0.19 DPS) [rep] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.89 DPS) | yes | Assault Band (13095, +0.00 DPS, sim-verified) [world_drop]; Protector's Band (19516, -0.09 DPS) [rep] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -6.03 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -2.93 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hanzo Sword (8190, -1.13 DPS, sim-verified) [world_drop] |
| off_hand | Doomforged Straightedge (12535) | Blackrock Depths: Anvilrage Overseer [dungeon] | sim-verified (152.6 DPS) | yes | Claw of Celebras (17738, -4.43 DPS) [dungeon]; Shadowblade (2163, -7.10 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -28.82 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.19 DPS) [dungeon]; Dark Iron Rifle (16004, -2.24 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Bloodrazor; off_hand: Doomforged Straightedge; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 856, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 329.6. Weights run: 1.6s. Verify run: 1.7s. 1891 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.427, strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=0.933 ± 0.062 per rating point (14 rating = 1%, 13.060 per %), hit=0.227 ± 0.010 per rating point (10 rating = 1%, 2.273 per %), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lightbreaker Greathelm (239517) | Leonid Barthalomew the Revered [vendor] | 141.6 attack_power points (8.18 DPS) | yes | Field Marshal's Plate Helm (231538, -3.63 DPS) [pvp]; Lightbreaker Helmet (239525, -16.01 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (329.6 DPS) | yes | Imperial Jewel (11933, -0.29 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.29 DPS) [quest]; Rage of Mugamba (19577, -4.25 DPS, sim-verified) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 101.8 attack_power points (5.88 DPS) | yes | Lieutenant Commander's Plate Shoulders (227045, -2.82 DPS) [pvp]; Lightbreaker Pauldrons (239524, -19.46 DPS, sim-verified) [vendor] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 39.8 attack_power points (2.30 DPS) | yes | Cloak of the Honor Guard (20073, +0.00 DPS, sim-verified) [rep]; Shadewood Cloak (18328, -0.54 DPS) [dungeon]; Howler's Furs (272414, -0.55 DPS) [vendor] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | sim-verified (329.6 DPS) | yes | Timbermaw Tunic (252484, -2.76 DPS) [crafted]; Obsidian Mail Tunic (22191, -3.31 DPS) [crafted]; Breastplate of Undead Slaying (23087, -28.90 DPS, sim-verified) [world] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (329.6 DPS) | yes | Berserker Bracers (19578, -1.54 DPS) [rep]; Marshal's Plate Bracers (16481, -1.83 DPS) [pvp]; Bracers of Undead Slaying (23090, -17.04 DPS, sim-verified) [world] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | sim-verified (329.6 DPS) | yes | Raider Gauntlets (272095, -1.76 DPS) [vendor]; Marshal's Plate Gauntlets (231541, -2.24 DPS) [pvp] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 101.9 attack_power points (5.89 DPS) | yes | Ferocity of the Timbermaw (227805, -1.93 DPS) [vendor]; Marshal's Plate Girdle (16482, -2.19 DPS) [pvp]; Radiant Girdle of the Dawn (227814, -4.72 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | sim-verified (329.6 DPS) | yes | Sentinel's Plate Legguards (237825, -3.24 DPS) [vendor]; Titanic Leggings (22385, -3.46 DPS) [crafted]; Cloudkeeper Legplates (14554, -13.88 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 101.9 attack_power points (5.88 DPS) | yes | Marshal's Plate Boots (231539, -3.00 DPS) [pvp]; Boots of Heroism (21995, -3.04 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (329.6 DPS) | yes | Band of the Ogre King (18522, -1.48 DPS) [dungeon]; Don Julio's Band (19325, -1.57 DPS) [rep]; Naglering (11669, -14.31 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (329.6 DPS) | yes | Band of the Ogre King (18522, -0.02 DPS) [dungeon]; Don Julio's Band (19325, -0.10 DPS) [rep]; Naglering (11669, -7.34 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (329.6 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Diamond Flask (20130, -0.85 DPS, sim-verified) [quest] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (329.6 DPS) | yes | Hand of Justice (11815, +0.00 DPS, sim-verified) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (329.6 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -15.59 DPS, sim-verified) [world_drop] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (329.6 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -14.46 DPS) [dungeon]; Skullflame Shield (1168, -82.07 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (329.6 DPS) | yes | Bloodseeker (19107, -0.01 DPS) [quest]; Skull Splitting Crossbow (13039, -0.06 DPS) [world_drop]; Dark Iron Rifle (16004, -4.23 DPS, sim-verified) [crafted] |

**New at 60:** head: Lightbreaker Greathelm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Shroud of Domination; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Riphook

No-known-source sample (15 of 1891, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (troll, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 29.7. Weights run: 1.2s. Verify run: 1.1s. 276 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=0.170 ± 0.005 per rating point (14 rating = 1%, 2.375 per %), hit=0.048 ± 0.002 per rating point (10 rating = 1%, 0.483 per %), melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.39 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, +0.00 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, +0.00 DPS, sim-verified) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.26 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.36 DPS) | yes | Bristlebark Bindings (14569, -0.21 DPS) [world_drop]; Runed Copper Bracers (2854, -0.22 DPS) [crafted]; Raptorcrest Bracers (270010, -0.39 DPS, sim-verified) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.1 attack_power points (0.58 DPS) | yes | Gold-flecked Gloves (5195, -0.11 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.21 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.29 DPS) [world_drop]; Cobrahn's Grasp (6460, -0.32 DPS, sim-verified) [dungeon] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.6 attack_power points (0.67 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.38 DPS) | yes | Veteran's Boots (250503, +0.00 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.02 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 attack_power points (8.58 DPS) | yes | Diamond Hammer (2194, -0.16 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.14 DPS) [world_drop]; Ruga's Bulwark (7120, -8.36 DPS) [quest] |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.0 attack_power points (0.15 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 276, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.8. Weights run: 1.2s. Verify run: 1.2s. 462 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=0.297 ± 0.022 per rating point (14 rating = 1%, 4.156 per %), hit=0.085 ± 0.004 per rating point (10 rating = 1%, 0.852 per %), melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 attack_power points (1.16 DPS) | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted]; Veteran's Chain Helm (250498, -0.24 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | River Pride Choker (13087, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.20 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.43 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 attack_power points (0.62 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Slayer's Cape (14752) (or Lambent Scale Cloak (4706)) | World drop [world_drop] | 10.1 attack_power points (0.36 DPS) | yes | Lambent Scale Cloak (4706, +0.00 DPS, sim-verified) [world_drop]; Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 37.9 attack_power points (1.33 DPS) | yes | Shining Silver Breastplate (2870, -0.24 DPS, sim-verified) [crafted]; Barbaric Iron Breastplate (7914, -0.27 DPS) [crafted]; Hard Gold Cuirass (250533, -0.36 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 attack_power points (0.71 DPS) | yes | Yorgen Bracers (13012, -0.16 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.26 DPS) [quest] |
| hands | Warsong Gauntlets (16978) | Warsong Supplies [quest] | 25.3 attack_power points (0.89 DPS) | yes | Gauntlets of Ogre Strength (3341, +0.00 DPS, sim-verified) [world]; Bonefist Gauntlets (4465, -0.09 DPS) [world]; Mail Combat Gauntlets (4075, -0.15 DPS) [world_drop] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 30.3 attack_power points (1.07 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Officer's Belt (250556, -0.13 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.22 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 attack_power points (1.03 DPS) | yes | Slayer's Pants (14757, -0.05 DPS) [world_drop]; Ferine Leggings (6690, -0.11 DPS) [dungeon]; Golden Scale Leggings (3843, -0.21 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 attack_power points (0.67 DPS) | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop]; Hard Gold Boots (250534, -0.21 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 attack_power points (0.73 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Ironspine's Eye (7686, -0.31 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.4 attack_power points (0.58 DPS) | yes | Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Tiger Band (6749, -0.16 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 attack_power points (12.23 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 attack_power points (11.90 DPS) | yes | Royal Diplomatic Scepter (9457, -1.18 DPS, sim-verified) [dungeon]; Shield of Thorsen (13079, -11.28 DPS) [world_drop]; Slayer's Shield (15892, -11.34 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.32 DPS) | yes | Long Battle Bow (15284, -0.05 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.14 DPS) [crafted]; Double-barreled Shotgun (2098, -0.17 DPS, sim-verified) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Slayer's Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Warsong Gauntlets; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 84.4. Weights run: 1.4s. Verify run: 1.4s. 652 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.333, strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=0.564 ± 0.039 per rating point (14 rating = 1%, 7.890 per %), hit=0.103 ± 0.005 per rating point (10 rating = 1%, 1.030 per %), melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.0 attack_power points (1.60 DPS) | yes | White Bandit Mask (10008, -0.25 DPS) [crafted]; Hard Gold Coif (250537, -0.29 DPS) [crafted]; Chromite Barbute (8142, -0.87 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.97 DPS) | yes | Ghostshard Talisman (7731, -0.27 DPS, sim-verified) [dungeon]; Ethereal Talisman (4430, -0.38 DPS) [quest]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 attack_power points (1.03 DPS) | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor]; Chromite Pauldrons (8144, -0.19 DPS, sim-verified) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 13.8 attack_power points (0.67 DPS) | yes | Wolfmaster Cape (6314, -0.18 DPS) [dungeon]; Wildhunter Cloak (16658, -0.18 DPS) [quest]; Hawkeye's Cloak (14593, -0.47 DPS, sim-verified) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 attack_power points (1.52 DPS) | yes | Avenger's Armor (1488, -0.11 DPS) [dungeon]; Jouster's Chestplate (8157, -0.20 DPS, sim-verified) [dungeon]; Shining Mithril Breastplate (250540, -0.21 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.97 DPS) | yes | Ravager's Armguards (14770, +0.00 DPS, sim-verified) [world_drop]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Darkspear Armsplints (4132, -0.31 DPS) [quest] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.55 DPS) | yes | Truesilver Gauntlets (7938, +0.00 DPS, sim-verified) [crafted]; Scarlet Gauntlets (10331, -0.13 DPS) [dungeon]; Gloves of Holy Might (867, -0.20 DPS) [world_drop] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.46 DPS) | yes | Defiler's Plate Girdle (20206, +0.00 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -0.05 DPS) [dungeon]; Tharg's Shoelace (9705, -0.24 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 attack_power points (1.97 DPS) | yes | Firemane Leggings (13129, -0.26 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.38 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 attack_power points (1.42 DPS) | yes | Prowler's Leather Shoes (252465, -0.26 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.29 DPS) [dungeon]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 20.3 attack_power points (0.99 DPS) | yes | Assault Band (13095, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Suspicious Spare Part (274754, -0.33 DPS) [vendor] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.97 DPS) | yes | Assault Band (13095, +0.00 DPS, sim-verified) [world_drop]; Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.02 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (84.4 DPS) | yes | Savage Boar's Guard (10767, -20.41 DPS) [dungeon]; Skullance Shield (13081, -20.58 DPS) [world_drop]; Ardent Custodian (868, -23.10 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.03 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.19 DPS) [quest]; Bow of Searing Arrows (2825, -0.59 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Dark Hooded Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 652, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 140.4. Weights run: 1.6s. Verify run: 1.8s. 837 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=0.333 ± 0.028 per rating point (14 rating = 1%, 4.668 per %), hit=0.100 ± 0.004 per rating point (10 rating = 1%, 0.999 per %), melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.7 attack_power points (4.23 DPS) | yes | Fury Visor (20521, -1.49 DPS) [quest]; Bloomsprout Headpiece (17767, -1.62 DPS, sim-verified) [dungeon]; Sunscale Helmet (14849, -1.63 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.89 DPS) | yes | Ghostshard Talisman (7731, -0.63 DPS, sim-verified) [dungeon]; Woven Ivy Necklace (19159, -0.69 DPS) [quest]; Skibi's Pendant (13089, -0.71 DPS) [world_drop] |
| shoulder | Blood Guard's Plate Pauldrons (220796) | Lady Palanseer [vendor] | sim-verified (140.4 DPS) | yes | Wyrmslayer Spaulders (13066, -0.03 DPS) [world_drop]; Earthslag Shoulders (11632, -0.14 DPS) [dungeon]; Officer's Pauldrons (250576, -1.89 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.4 attack_power points (1.37 DPS) | yes | Battlehard Cape (11858, -0.42 DPS) [quest]; Wildhunter Cloak (16658, -0.42 DPS) [quest]; Bloodlust Cape (14801, -0.79 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 38.0 attack_power points (3.60 DPS) | yes | Mixologist's Tunic (12793, +0.00 DPS, sim-verified) [dungeon]; Valorous Chestguard (8274, -0.60 DPS) [world_drop]; Coldmetal Guard (274758, -0.75 DPS) [vendor] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.65 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Runed Golem Shackles (12550, -0.55 DPS) [dungeon]; Branded Leather Bracers (19508, -0.76 DPS) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 38.4 attack_power points (3.63 DPS) | yes | Officer's Gloves (250551, -0.93 DPS) [crafted]; Rockgrip Gauntlets (17736, -0.98 DPS) [dungeon]; Gauntlets of Divinity (7724, -3.14 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 42.7 attack_power points (4.04 DPS) | yes | Atal'alarion's Tusk Ring (10798, -1.34 DPS) [dungeon]; Belt of the Gladiator (13134, -1.34 DPS) [world_drop]; Defiler's Leather Girdle (20192, -1.71 DPS, sim-verified) [rep] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.8 attack_power points (3.30 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Silvershell Leggings (10633, -0.30 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.30 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 26.7 attack_power points (2.53 DPS) | yes | Prowler's Leather Boots (252468, -0.08 DPS, sim-verified) [crafted]; Officer's Sabatons (250561, -0.28 DPS) [crafted]; Officer's Boots (250546, -0.35 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.27 DPS) | yes | Mark of Kern (2262, -0.38 DPS) [dungeon]; Assault Band (13095, -0.38 DPS) [world_drop]; Legionnaire's Band (19511, -0.47 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.99 DPS) | yes | Assault Band (13095, -0.09 DPS) [world_drop]; Legionnaire's Band (19511, -0.19 DPS) [rep]; Mark of Kern (2262, -2.94 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -3.96 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -4.88 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hanzo Sword (8190, -0.98 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (51.63 DPS) | yes | Doomforged Straightedge (12535, +0.00 DPS, sim-verified) [dungeon]; Claw of Celebras (17738, -5.96 DPS) [dungeon]; White Bone Shredder (11863, -9.59 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.19 DPS) [dungeon]; Dark Iron Rifle (16004, -1.92 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Blood Guard's Plate Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 837, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 310.5. Weights run: 1.6s. Verify run: 1.7s. 1872 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.427, strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=0.933 ± 0.062 per rating point (14 rating = 1%, 13.060 per %), hit=0.227 ± 0.010 per rating point (10 rating = 1%, 2.273 per %), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lightbreaker Greathelm (239517) | Leonid Barthalomew the Revered [vendor] | 141.6 attack_power points (8.18 DPS) | yes | Warlord's Plate Headpiece (231535, -3.63 DPS) [pvp]; Lightbreaker Helmet (239525, -15.04 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (310.5 DPS) | yes | Imperial Jewel (11933, -0.29 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.29 DPS) [quest]; Rage of Mugamba (19577, -2.52 DPS, sim-verified) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 101.8 attack_power points (5.88 DPS) | yes | Champion's Plate Shoulders (227042, -2.82 DPS) [pvp]; Lightbreaker Pauldrons (239524, -14.64 DPS, sim-verified) [vendor] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 39.8 attack_power points (2.30 DPS) | yes | Deathguard's Cloak (20068, +0.00 DPS, sim-verified) [rep]; Shadewood Cloak (18328, -0.54 DPS) [dungeon]; Howler's Furs (272414, -0.55 DPS) [vendor] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | sim-verified (310.5 DPS) | yes | Timbermaw Tunic (252484, -2.76 DPS) [crafted]; Obsidian Mail Tunic (22191, -3.31 DPS) [crafted]; Breastplate of Undead Slaying (23087, -25.73 DPS, sim-verified) [world] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (310.5 DPS) | yes | Berserker Bracers (19578, -1.54 DPS) [rep]; General's Plate Armguards (16546, -1.83 DPS) [pvp]; Bracers of Undead Slaying (23090, -15.28 DPS, sim-verified) [world] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | sim-verified (310.5 DPS) | yes | Raider Gauntlets (272095, -1.76 DPS) [vendor]; General's Plate Gauntlets (231532, -2.24 DPS) [pvp] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 101.9 attack_power points (5.89 DPS) | yes | Ferocity of the Timbermaw (227805, -1.93 DPS) [vendor]; General's Plate Girdle (16547, -2.19 DPS) [pvp]; Radiant Girdle of the Dawn (227814, -5.83 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | sim-verified (310.5 DPS) | yes | Sentinel's Plate Legguards (237825, -3.24 DPS) [vendor]; Titanic Leggings (22385, -3.46 DPS) [crafted]; Cloudkeeper Legplates (14554, -13.89 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 101.9 attack_power points (5.88 DPS) | yes | General's Plate Boots (231531, -3.00 DPS) [pvp]; Boots of Heroism (21995, -3.04 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (310.5 DPS) | yes | Band of the Ogre King (18522, -1.48 DPS) [dungeon]; Don Julio's Band (19325, -1.57 DPS) [rep]; Naglering (11669, -10.55 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (310.5 DPS) | yes | Band of the Ogre King (18522, -0.02 DPS) [dungeon]; Don Julio's Band (19325, -0.10 DPS) [rep]; Naglering (11669, -7.01 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (310.5 DPS) | yes | Counterattack Lodestone (18537, -1.25 DPS) [dungeon]; Hand of Justice (11815, -1.36 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.40 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (310.5 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -2.24 DPS, sim-verified) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (310.5 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Felstriker (12590, -5.38 DPS, sim-verified) [dungeon] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (310.5 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -14.46 DPS) [dungeon]; Skullflame Shield (1168, -84.40 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (310.5 DPS) | yes | Bloodseeker (19107, -0.01 DPS) [quest]; Skull Splitting Crossbow (13039, -0.06 DPS) [world_drop]; Dark Iron Rifle (16004, -3.47 DPS, sim-verified) [crafted] |

**New at 60:** head: Lightbreaker Greathelm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Shroud of Domination; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Legionnaire's Band; trinket2: Darkmoon Card: Maelstrom; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Riphook

No-known-source sample (15 of 1872, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

