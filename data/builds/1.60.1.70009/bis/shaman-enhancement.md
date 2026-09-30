# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 43.1. Weights run: 1.5s. Verify run: 1.5s. 267 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=0.309 ± 0.007 per rating point (14 rating = 1%, 4.333 per %), hit=0.439 ± 0.005 per rating point (10 rating = 1%, 4.391 per %), melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.61 DPS) [crafted]; Brawler's Leather Hood (252504, -0.63 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 1.4 attack_power points (0.05 DPS) | yes | Erudite's Amulet (277204, +0.00 DPS, sim-verified) [quest] |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, -0.13 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Grave Shroud (279865, -0.08 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.11 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.13 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.17 DPS) [quest]; Bristlebark Bindings (14569, -0.18 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.55 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -0.14 DPS) [quest]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.08 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Warchief's Girdle (5750, -0.28 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.76 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.11 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 attack_power points (0.39 DPS) | yes | Veteran's Boots (250503, +0.00 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.04 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.9 attack_power points (0.31 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop]; Signet of the Zhevra (285330, -0.26 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.16 DPS, sim-verified) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop]; Signet of the Zhevra (285330, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.0 attack_power points (8.18 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 attack_power points (8.03 DPS) | yes | Blackfang (2236, -2.09 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -7.62 DPS) [world_drop]; Bear Buckler (4821, -7.82 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer

No-known-source sample (15 of 267, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 69.4. Weights run: 1.5s. Verify run: 1.6s. 433 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=0.525 ± 0.024 per rating point (14 rating = 1%, 7.348 per %), hit=0.619 ± 0.009 per rating point (10 rating = 1%, 6.193 per %), melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Veteran's Chain Helm (250498, -0.11 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Sentinel's Medallion (19541, -0.37 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.7 attack_power points (0.52 DPS) | yes | Barbaric Iron Shoulders (7913, +0.00 DPS, sim-verified) [crafted]; Golden Scale Shoulders (3841, -0.03 DPS) [crafted]; Mail Combat Spaulders (6404, -0.03 DPS) [world_drop] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (69.4 DPS) | yes | Hawkeye's Cloak (14593, -0.02 DPS) [world_drop]; Slayer's Cape (14752, -0.06 DPS) [world_drop]; Wolfmaster Cape (6314, -0.77 DPS, sim-verified) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.06 DPS) | yes | Shining Silver Breastplate (2870, -0.11 DPS, sim-verified) [crafted]; Barbaric Iron Breastplate (7914, -0.21 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.26 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.57 DPS) | yes | Yorgen Bracers (13012, +0.00 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.14 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.78 DPS) | yes | The Frozen Clutch (23170, -0.11 DPS, sim-verified) [dungeon]; Bonefist Gauntlets (4465, -0.14 DPS) [world]; Mail Combat Gauntlets (4075, -0.15 DPS) [world_drop] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 attack_power points (0.85 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Golden Scale Leggings (3843, -0.14 DPS) [crafted]; Slayer's Pants (14757, -0.14 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -0.89 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.0 attack_power points (0.60 DPS) | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Alacritous Treads (277234, -0.17 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.3 attack_power points (0.61 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.6 attack_power points (0.52 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Tiger Band (6749, -0.16 DPS, sim-verified) [quest]; Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (12.21 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (11.83 DPS) | yes | Royal Diplomatic Scepter (9457, -1.12 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -3.89 DPS) [quest]; Shield of Thorsen (13079, -11.34 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist

No-known-source sample (15 of 433, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 88.3. Weights run: 1.8s. Verify run: 1.6s. 582 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=0.551 ± 0.015 per rating point (14 rating = 1%, 7.710 per %), hit=0.521 ± 0.008 per rating point (10 rating = 1%, 5.206 per %), melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.7 attack_power points (1.25 DPS) | yes | White Bandit Mask (10008, -0.28 DPS) [crafted]; Tusken Helm (6686, -0.29 DPS) [dungeon]; Hard Gold Coif (250537, -0.81 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.74 DPS) | yes | Ghostshard Talisman (7731, -0.38 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; Gazlowe's Charm (13088, -0.44 DPS) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.81 DPS) | yes | Imperial Leather Spaulders (4737, +0.00 DPS, sim-verified) [dungeon]; Wrangling Spaulders (15698, -0.19 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.3 attack_power points (0.53 DPS) | yes | Wolfmaster Cape (6314, -0.16 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.18 DPS) [pvp]; Dark Hooded Cape (5257, -1.56 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 attack_power points (1.14 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Golden Scale Cuirass (3845, -0.10 DPS) [crafted]; Avenger's Armor (1488, -0.24 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Ravager's Armguards (14770, -0.17 DPS) [world_drop]; Pugilist Bracers (4438, -0.25 DPS, sim-verified) [dungeon]; Yorgen Bracers (13012, -0.25 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.18 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS, sim-verified) [dungeon]; Gloves of Holy Might (867, -0.16 DPS) [world_drop]; Reticulated Bone Gauntlets (9435, -0.22 DPS) [world_drop] |
| waist | Boar Champion's Belt (10768) | Razorfen Downs: Ragglesnout [dungeon] | 30.0 attack_power points (1.11 DPS) | yes | Highlander's Leather Girdle (20116, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.22 DPS) [world_drop]; Scarlet Belt (10329, -0.22 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.55 DPS) | yes | Firemane Leggings (13129, -0.25 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.30 DPS) [crafted]; Legguards of the Vault (9396, -0.52 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 attack_power points (0.91 DPS) | yes | Skirmisher's Mail Boots (252564, +0.00 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.11 DPS) [dungeon]; Ironheel Boots (4653, -0.17 DPS) [quest] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.74 DPS) | yes | Protector's Band (19515, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Protector's Band (19515, +0.00 DPS, sim-verified) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (17.56 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 attack_power points (17.03 DPS) | yes | Shoni's Disarming Tool (9608, -8.71 DPS) [quest]; Curve-bladed Ripper (2815, -9.47 DPS, sim-verified) [world_drop]; Savage Boar's Guard (10767, -16.22 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Boar Champion's Belt; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Assault Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Ardent Custodian

No-known-source sample (15 of 582, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (dwarf, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 127.9. Weights run: 1.8s. Verify run: 2.1s. 752 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=0.710 ± 0.019 per rating point (14 rating = 1%, 9.934 per %), hit=0.802 ± 0.012 per rating point (10 rating = 1%, 8.019 per %), melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.74 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Bloomsprout Headpiece (17767, -0.43 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.44 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.72 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.22 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.36 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 27.9 attack_power points (1.01 DPS) | yes | Failed Flying Experiment (9647, -0.10 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted]; Prowler's Leather Shoulder (252534, -0.75 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.2 attack_power points (0.69 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Dark Hooded Cape (5257, -0.22 DPS) [world]; Bloodlust Cape (14801, -0.30 DPS, sim-verified) [world_drop] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 42.2 attack_power points (1.53 DPS) | yes | Mixologist's Tunic (12793, +0.00 DPS, sim-verified) [dungeon]; Knight's Mail Armor (223078, -0.30 DPS) [vendor]; Kolkar Marauder Chain (6773, -0.38 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) (or Arena Bands (18711)) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.01 DPS) | yes | Arena Bands (18711, +0.00 DPS, sim-verified) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted]; Branded Leather Bracers (19508, -0.29 DPS) [dungeon] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 32.6 attack_power points (1.18 DPS) | yes | Gauntlets of Divinity (7724, -0.02 DPS) [dungeon]; Raider Gloves (272100, -0.05 DPS) [vendor]; Maddening Gauntlets (11867, -0.32 DPS, sim-verified) [quest] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.66 DPS) | yes | Prowler's Leather Waistguard (252473, -0.43 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.54 DPS) [crafted]; Belt of the Gladiator (13134, -0.67 DPS, sim-verified) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.52 DPS) | yes | Gryphon Rider's Leggings (9652, -0.16 DPS) [quest]; Firemane Leggings (13129, -0.27 DPS, sim-verified) [world_drop]; Serpentskin Leggings (8262, -0.28 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 31.6 attack_power points (1.14 DPS) | yes | Skirmisher's Mail Sabatons (252578, -0.20 DPS) [crafted]; Shadefiend Boots (11675, -0.22 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.36 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.01 DPS) | yes | Mark of Kern (2262, -0.29 DPS) [dungeon]; Assault Band (13095, -0.29 DPS) [world_drop]; Thunderbrow Ring (13097, -0.38 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.6 attack_power points (0.89 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; Protector's Band (19515, -0.39 DPS, sim-verified) [rep] |
| trinket1 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-verified (127.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (127.9 DPS) | yes | Frozen Heart of the Mountain (249469, -2.27 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (127.9 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Barman Shanker (12791, -13.44 DPS, sim-verified) [dungeon] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (19.72 DPS) | yes | Claw of Celebras (17738, -2.27 DPS) [dungeon]; Shoni's Disarming Tool (9608, -11.59 DPS) [quest]; Grizzle's Skinner (11702, -16.53 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Mail Epaulets; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Bracers of the Stone Princess; hands: Prowler's Leather Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Fire Ruby; trinket2: Diamond Flask; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade

No-known-source sample (15 of 752, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (dwarf, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 222.2. Weights run: 1.9s. Verify run: 2.1s. 1692 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=0.737 ± 0.023 per rating point (14 rating = 1%, 10.313 per %), hit=0.977 ± 0.017 per rating point (10 rating = 1%, 9.771 per %), melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Helmet (240131) | Leonid Barthalomew the Revered [vendor] | 90.0 attack_power points (3.24 DPS) | yes | Soulcrusher Faceguard (240104, -1.67 DPS) [vendor]; Warbear Helm (252485, -1.68 DPS) [crafted]; Embrace of the Lycan (9479, -4.80 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.3 attack_power points (1.24 DPS) | yes | Beads of Ogre Might (22150, +0.00 DPS, sim-verified) [quest]; Imperial Jewel (11933, -0.08 DPS) [dungeon]; Will of the Martyr (17044, -0.16 DPS) [quest] |
| shoulder | Soulcrusher Epaulets (240135) | Leonid Barthalomew the Revered [vendor] | 75.8 attack_power points (2.73 DPS) | yes | Highlander's Leather Shoulders (20059, -1.31 DPS) [rep]; Highlander's Lizardhide Shoulders (20060, -1.43 DPS) [rep]; Black Dragonscale Shoulders (15051, -8.10 DPS, sim-verified) [crafted] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 37.8 attack_power points (1.36 DPS) | yes | Shroud of Domination (22337, -0.14 DPS) [dungeon]; Cape of the Black Baron (13340, -0.36 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.64 DPS, sim-verified) [rep] |
| chest | Soulcrusher Armor (240128) | Leonid Barthalomew the Revered [vendor] | sim-verified (222.2 DPS) | yes | Obsidian Mail Tunic (22191, -1.87 DPS) [crafted]; Timbermaw Tunic (252484, -1.89 DPS) [crafted]; Tunic of Undead Slaying (23089, -22.31 DPS, sim-verified) [world] |
| wrist | Soulcrusher Vambraces (240137) | Leonid Barthalomew the Revered [vendor] | sim-verified (222.2 DPS) | yes | Windtalker's Wristguards (19582, -0.86 DPS) [rep]; Wristwraps of Undead Slaying (23093, -7.72 DPS, sim-verified) [world] |
| hands | Soulcrusher Grips (240130) | Leonid Barthalomew the Revered [vendor] | 84.6 attack_power points (3.05 DPS) | yes | Studded Timbermaw Brawlers (227809, -1.42 DPS) [vendor]; Cadaverous Gloves (14640, -1.46 DPS) [dungeon]; Timbermaw Brawlers (19049, -6.61 DPS, sim-verified) [crafted] |
| waist | Soulcrusher Belt (240136) | Leonid Barthalomew the Revered [vendor] | 82.3 attack_power points (2.97 DPS) | yes | Ferocity of the Timbermaw (227805, -0.77 DPS) [vendor]; Might of the Timbermaw (19044, -1.19 DPS) [crafted]; Dense Timbermaw Belt (227807, -5.20 DPS, sim-verified) [vendor] |
| legs | Soulcrusher Leggings (240134) | Leonid Barthalomew the Revered [vendor] | 120.6 attack_power points (4.35 DPS) | yes | Devilsaur Leggings (15062, -2.32 DPS) [crafted]; Black Dragonscale Leggings (15052, -2.40 DPS) [crafted]; Warbear Woolies (15065, -8.38 DPS, sim-verified) [crafted] |
| feet | Soulcrusher Treads (240129) | Leonid Barthalomew the Revered [vendor] | 72.0 attack_power points (2.60 DPS) | yes | Pads of the Dread Wolf (13210, -1.15 DPS) [dungeon]; Drudge Boots (21532, -1.34 DPS) [quest]; Scalegut Treaders (275618, -1.63 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (222.2 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.14 DPS) [vendor]; Naglering (11669, -8.10 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (222.2 DPS) | yes | Blackstone Ring (17713, -0.23 DPS) [dungeon]; Protector's Band (19514, -0.23 DPS) [rep]; Naglering (11669, -7.39 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (222.2 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (222.2 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Eye of the Beast (13968, +0.00 DPS, sim-verified) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (222.2 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS, sim-verified) [world_drop]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | sim-verified (222.2 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Distracting Dagger (18392, -3.84 DPS) [dungeon]; Skullflame Shield (1168, -24.50 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Soulcrusher Helmet; neck: Medallion of the Dawn; shoulder: Soulcrusher Epaulets; back: Howler's Furs; chest: Soulcrusher Armor; wrist: Soulcrusher Vambraces; hands: Soulcrusher Grips; waist: Soulcrusher Belt; legs: Soulcrusher Leggings; feet: Soulcrusher Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Shadowsong's Sorrow; off_hand: Greenhammer

No-known-source sample (15 of 1692, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (orc, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 45.5. Weights run: 1.5s. Verify run: 1.4s. 248 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=0.309 ± 0.007 per rating point (14 rating = 1%, 4.333 per %), hit=0.439 ± 0.005 per rating point (10 rating = 1%, 4.391 per %), melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.18 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.61 DPS) [crafted]; Brawler's Leather Hood (252504, -0.63 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 1.4 attack_power points (0.05 DPS) | yes | Erudite's Amulet (277204, -0.06 DPS, sim-verified) [quest] |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, +0.00 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.16 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Bristlebark Bindings (14569, -0.18 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.18 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.55 DPS) | yes | Gold-flecked Gloves (5195, -0.09 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.18 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.06 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Warchief's Girdle (5750, -0.28 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 19.2 attack_power points (0.66 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.04 DPS) [crafted]; Hulking Leggings (14748, -0.09 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 attack_power points (0.39 DPS) | yes | Veteran's Boots (250503, +0.00 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.04 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.9 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.0 attack_power points (8.18 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 attack_power points (8.03 DPS) | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -7.62 DPS) [world_drop]; Bear Buckler (4821, -7.82 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders

### Band 30 (orc, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 74.2. Weights run: 1.5s. Verify run: 1.5s. 417 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=0.525 ± 0.024 per rating point (14 rating = 1%, 7.348 per %), hit=0.619 ± 0.009 per rating point (10 rating = 1%, 6.193 per %), melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Veteran's Chain Helm (250498, -0.13 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.14 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (74.2 DPS) | yes | River Pride Choker (13087, -0.06 DPS) [world_drop]; Scout's Medallion (19537, -0.22 DPS) [rep]; Ghostshard Talisman (7731, -1.05 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.7 attack_power points (0.52 DPS) | yes | Barbaric Iron Shoulders (7913, +0.00 DPS, sim-verified) [crafted]; Golden Scale Shoulders (3841, -0.03 DPS) [crafted]; Mail Combat Spaulders (6404, -0.03 DPS) [world_drop] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.35 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Hawkeye's Cloak (14593, -0.03 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.06 DPS) | yes | Shining Silver Breastplate (2870, -0.13 DPS, sim-verified) [crafted]; Barbaric Iron Breastplate (7914, -0.21 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.26 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.57 DPS) | yes | Yorgen Bracers (13012, +0.00 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.14 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.78 DPS) | yes | The Frozen Clutch (23170, -0.07 DPS) [dungeon]; Warsong Gauntlets (16978, -0.08 DPS, sim-verified) [quest]; Bonefist Gauntlets (4465, -0.14 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 attack_power points (0.85 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Golden Scale Leggings (3843, -0.14 DPS) [crafted]; Slayer's Pants (14757, -0.14 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -0.70 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.0 attack_power points (0.60 DPS) | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.18 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.3 attack_power points (0.61 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.6 attack_power points (0.52 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon]; Tiger Band (6749, -0.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (12.21 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (11.83 DPS) | yes | Royal Diplomatic Scepter (9457, -2.26 DPS, sim-verified) [dungeon]; Shield of Thorsen (13079, -11.34 DPS) [world_drop]; Slayer's Shield (15892, -11.36 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 40 (orc, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 85.2. Weights run: 1.8s. Verify run: 1.7s. 565 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=0.551 ± 0.015 per rating point (14 rating = 1%, 7.710 per %), hit=0.521 ± 0.008 per rating point (10 rating = 1%, 5.206 per %), melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.7 attack_power points (1.25 DPS) | yes | White Bandit Mask (10008, -0.28 DPS) [crafted]; Tusken Helm (6686, -0.29 DPS) [dungeon]; Hard Gold Coif (250537, -1.30 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.74 DPS) | yes | Ethereal Talisman (4430, -0.31 DPS) [quest]; Ghostshard Talisman (7731, -0.34 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -0.39 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.81 DPS) | yes | Wrangling Spaulders (15698, -0.19 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor]; Imperial Leather Spaulders (4737, -0.25 DPS, sim-verified) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 11.9 attack_power points (0.44 DPS) | yes | Wolfmaster Cape (6314, -0.07 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.12 DPS) [world_drop]; Wildhunter Cloak (16658, -0.94 DPS, sim-verified) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 attack_power points (1.14 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Golden Scale Cuirass (3845, -0.10 DPS) [crafted]; Avenger's Armor (1488, -0.72 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Ravager's Armguards (14770, -0.17 DPS) [world_drop]; Pugilist Bracers (4438, -0.19 DPS, sim-verified) [dungeon]; Darkspear Armsplints (4132, -0.22 DPS) [quest] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.18 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS, sim-verified) [dungeon]; Gloves of Holy Might (867, -0.16 DPS) [world_drop]; Reticulated Bone Gauntlets (9435, -0.22 DPS) [world_drop] |
| waist | Boar Champion's Belt (10768) | Razorfen Downs: Ragglesnout [dungeon] | 30.0 attack_power points (1.11 DPS) | yes | Defiler's Leather Girdle (20192, -0.07 DPS, sim-verified) [rep]; Tharg's Shoelace (9705, -0.15 DPS) [quest]; Scarlet Belt (10329, -0.22 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.55 DPS) | yes | Firemane Leggings (13129, -0.24 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.30 DPS) [crafted]; Legguards of the Vault (9396, -0.52 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 attack_power points (0.91 DPS) | yes | Blackforge Greaves (6423, -0.11 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.24 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.52 DPS, sim-verified) [crafted] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.74 DPS) | yes | Legionnaire's Band (19512, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Legionnaire's Band (19512, +0.00 DPS, sim-verified) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (17.56 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Curve-bladed Ripper (2815) | World drop [world_drop] | sim-verified (85.2 DPS) | yes | Ardent Custodian (868, -3.23 DPS, sim-verified) [world_drop]; Savage Boar's Guard (10767, -15.46 DPS) [dungeon]; Skullance Shield (13081, -15.63 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Dark Hooded Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Boar Champion's Belt; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Assault Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Curve-bladed Ripper

No-known-source sample (15 of 565, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 128.1. Weights run: 1.8s. Verify run: 2.1s. 722 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=0.710 ± 0.019 per rating point (14 rating = 1%, 9.934 per %), hit=0.802 ± 0.012 per rating point (10 rating = 1%, 8.019 per %), melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.74 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Bloomsprout Headpiece (17767, -0.43 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.44 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.72 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Woven Ivy Necklace (19159, -0.12 DPS) [quest]; Ghostshard Talisman (7731, -0.22 DPS) [dungeon] |
| shoulder | Blood Guard's Mail Epaulets (220823) | Lady Palanseer [vendor] | 27.9 attack_power points (1.01 DPS) | yes | Failed Flying Experiment (9647, -0.10 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted]; Prowler's Leather Shoulder (252534, -1.40 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.2 attack_power points (0.69 DPS) | yes | Dark Hooded Cape (5257, -0.22 DPS) [world]; Pridelord Cape (14673, -0.31 DPS) [world_drop]; Bloodlust Cape (14801, -1.35 DPS, sim-verified) [world_drop] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 42.2 attack_power points (1.53 DPS) | yes | Stone Guard's Mail Armor (220826, -0.30 DPS) [vendor]; Kolkar Marauder Chain (6773, -0.38 DPS) [quest]; Mixologist's Tunic (12793, -0.64 DPS, sim-verified) [dungeon] |
| wrist | Bracers of the Stone Princess (17714) (or Arena Bands (18711)) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.01 DPS) | yes | Prowler's Leather Bracers (252539, -0.23 DPS) [crafted]; Branded Leather Bracers (19508, -0.29 DPS) [dungeon]; Arena Bands (18711, -1.17 DPS, sim-verified) [world] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 32.6 attack_power points (1.18 DPS) | yes | Raider Gloves (272100, -0.05 DPS) [vendor]; Gloves of Holy Might (867, -0.10 DPS) [world_drop]; Gauntlets of Divinity (7724, -1.35 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.66 DPS) | yes | Prowler's Leather Waistguard (252473, -0.43 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.54 DPS) [crafted]; Belt of the Gladiator (13134, -0.63 DPS, sim-verified) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.52 DPS) | yes | Serpentskin Leggings (8262, -0.28 DPS) [world_drop]; Firemane Leggings (13129, -0.29 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.29 DPS) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 31.6 attack_power points (1.14 DPS) | yes | Skirmisher's Mail Sabatons (252578, -0.20 DPS) [crafted]; Shadefiend Boots (11675, -0.22 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.32 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.01 DPS) | yes | White Bone Band (11862, -0.15 DPS) [quest]; Mark of Kern (2262, -0.29 DPS) [dungeon]; Assault Band (13095, -0.29 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.6 attack_power points (0.89 DPS) | yes | Legionnaire's Band (19512, -0.16 DPS) [rep]; Assault Band (13095, -0.17 DPS) [world_drop]; White Bone Band (11862, -0.88 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (128.1 DPS) | yes | Frozen Heart of the Mountain (249469, -1.46 DPS) [crafted]; Fire Ruby (20036, -3.14 DPS, sim-verified) [quest] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (128.1 DPS) | yes | Frozen Heart of the Mountain (249469, -2.18 DPS, sim-verified) [crafted] |
| main_hand | Flurry Axe (871) | World drop [world_drop] | sim-verified (128.1 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hammer of the Northern Wind (810, -7.44 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (19.72 DPS) | yes | Claw of Celebras (17738, -2.27 DPS) [dungeon]; Grizzle's Skinner (11702, -3.11 DPS, sim-verified) [dungeon]; White Bone Shredder (11863, -3.56 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Blood Guard's Mail Epaulets; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Bracers of the Stone Princess; hands: Prowler's Leather Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Flurry Axe; off_hand: Shadowblade

No-known-source sample (15 of 722, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 218.6. Weights run: 1.9s. Verify run: 2.0s. 1616 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=0.737 ± 0.023 per rating point (14 rating = 1%, 10.313 per %), hit=0.977 ± 0.017 per rating point (10 rating = 1%, 9.771 per %), melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Helmet (240131) | Leonid Barthalomew the Revered [vendor] | 90.0 attack_power points (3.24 DPS) | yes | Champion's Mail Headguard (227155, -1.21 DPS) [pvp]; Soulcrusher Faceguard (240104, -1.67 DPS) [vendor]; Embrace of the Lycan (9479, -4.55 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.3 attack_power points (1.24 DPS) | yes | Beads of Ogre Might (22150, +0.00 DPS, sim-verified) [quest]; Imperial Jewel (11933, -0.08 DPS) [dungeon]; Will of the Martyr (17044, -0.16 DPS) [quest] |
| shoulder | Soulcrusher Epaulets (240135) | Leonid Barthalomew the Revered [vendor] | 75.8 attack_power points (2.73 DPS) | yes | Warlord's Mail Pauldrons (231654, -0.81 DPS, sim-verified) [vendor]; Champion's Mail Pauldrons (227154, -1.07 DPS) [pvp]; Black Dragonscale Shoulders (15051, -1.29 DPS) [crafted] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 37.8 attack_power points (1.36 DPS) | yes | Shroud of Domination (22337, -0.14 DPS) [dungeon]; Cape of the Black Baron (13340, -0.36 DPS) [dungeon]; Deathguard's Cloak (20068, -1.63 DPS, sim-verified) [rep] |
| chest | Soulcrusher Armor (240128) | Leonid Barthalomew the Revered [vendor] | sim-verified (218.6 DPS) | yes | Obsidian Mail Tunic (22191, -1.87 DPS) [crafted]; Timbermaw Tunic (252484, -1.89 DPS) [crafted]; Tunic of Undead Slaying (23089, -21.92 DPS, sim-verified) [world] |
| wrist | Soulcrusher Vambraces (240137) | Leonid Barthalomew the Revered [vendor] | sim-verified (218.6 DPS) | yes | Windtalker's Wristguards (19582, -0.86 DPS) [rep]; Wristwraps of Undead Slaying (23093, -8.94 DPS, sim-verified) [world] |
| hands | Soulcrusher Grips (240130) | Leonid Barthalomew the Revered [vendor] | 84.6 attack_power points (3.05 DPS) | yes | Timbermaw Brawlers (19049, -1.27 DPS) [crafted]; Studded Timbermaw Brawlers (227809, -1.42 DPS) [vendor]; General's Mail Vices (231655, -3.33 DPS, sim-verified) [vendor] |
| waist | Soulcrusher Belt (240136) | Leonid Barthalomew the Revered [vendor] | 82.3 attack_power points (2.97 DPS) | yes | Ferocity of the Timbermaw (227805, -0.77 DPS) [vendor]; Might of the Timbermaw (19044, -1.19 DPS) [crafted]; Dense Timbermaw Belt (227807, -2.32 DPS, sim-verified) [vendor] |
| legs | Soulcrusher Leggings (240134) | Leonid Barthalomew the Revered [vendor] | 120.6 attack_power points (4.35 DPS) | yes | Warbear Woolies (15065, -1.99 DPS) [crafted]; Legionnaire's Mail Legguards (227156, -2.04 DPS) [pvp]; General's Mail Legguards (231658, -3.64 DPS, sim-verified) [vendor] |
| feet | Soulcrusher Treads (240129) | Leonid Barthalomew the Revered [vendor] | 72.0 attack_power points (2.60 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS, sim-verified) [vendor]; Scalegut Treaders (275618, -1.08 DPS) [crafted]; Pads of the Dread Wolf (13210, -1.15 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (218.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.14 DPS) [vendor]; Naglering (11669, -7.09 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (218.6 DPS) | yes | Blackstone Ring (17713, -0.23 DPS) [dungeon]; Legionnaire's Band (19510, -0.23 DPS) [rep]; Naglering (11669, -6.29 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (218.6 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -1.95 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (218.6 DPS) | yes | Smolderweb's Eye (13213, -0.81 DPS, sim-verified) [dungeon]; Counterattack Lodestone (18537, -0.97 DPS) [dungeon]; Hand of Justice (11815, -1.04 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (218.6 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Felstriker (12590, -1.77 DPS, sim-verified) [dungeon] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | sim-verified (218.6 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Distracting Dagger (18392, -3.84 DPS) [dungeon]; Skullflame Shield (1168, -16.17 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Soulcrusher Helmet; neck: Medallion of the Dawn; shoulder: Soulcrusher Epaulets; back: Howler's Furs; chest: Soulcrusher Armor; wrist: Soulcrusher Vambraces; hands: Soulcrusher Grips; waist: Soulcrusher Belt; legs: Soulcrusher Leggings; feet: Soulcrusher Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Greenhammer

No-known-source sample (15 of 1616, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

