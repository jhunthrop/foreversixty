# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 27.6. Weights run: 1.1s. Verify run: 2.0s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.727 ± 0.028, crit=0.316 ± 0.020 per rating point (14 rating = 1%, 4.418 per %), hit=1.247 ± 0.074 per rating point (10 rating = 1%, 12.474 per %), melee_haste=9.069 ± 0.670

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.45 DPS) [crafted]; Brawler's Leather Hood (252504, -0.50 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 4.4 attack_power points (0.15 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.08 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Veteran's Chain Shirt (250488, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.14 DPS) [crafted]; Brawler's Leather Armor (252490, -0.17 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Bravo's Armbands (270015, -0.11 DPS) [quest]; Bristlebark Bindings (14569, -0.14 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.6 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.25 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Ruffian Belt (5975, -0.21 DPS) [world]; Brawler's Leather Belt (252428, -0.25 DPS) [crafted]; Cobrahn's Grasp (6460, -1.42 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.6 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Chausses of Westfall (6087, -1.25 DPS, sim-verified) [quest] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (27.6 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -1.06 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.9 attack_power points (0.39 DPS) | yes | Signet of the Zhevra (285330, -0.23 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world]; Ring of the Moon (12052, -0.32 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Signet of the Zhevra (285330, -0.13 DPS) [world]; The 1 Ring (8350, -0.19 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.43 DPS) | yes | Living Root (6631, -0.89 DPS) [dungeon]; Duskbringer (2205, -0.98 DPS) [dungeon]; Smite's Mighty Hammer (7230, -1.08 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567)) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, -0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.04 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 77.2. Weights run: 1.3s. Verify run: 1.3s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.086 ± 0.045, crit=0.543 ± 0.035 per rating point (14 rating = 1%, 7.605 per %), hit=1.755 ± 0.170 per rating point (10 rating = 1%, 17.552 per %), melee_haste=11.589 ± 1.312

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 27.8 attack_power points (1.10 DPS) | yes | Veteran's Chain Helm (250498, -0.15 DPS) [crafted]; Defender's Leather Helm (252455, -0.15 DPS) [crafted]; Tusken Helm (6686, -0.41 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.55 DPS) | yes | Kaleidoscope Chain (13084, -0.07 DPS) [world_drop]; Sentinel's Medallion (19541, -0.21 DPS) [rep]; River Pride Choker (13087, -0.24 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.9 attack_power points (0.87 DPS) | yes | Barbaric Shoulders (5964, -0.26 DPS) [crafted]; Golden Scale Shoulders (3841, -0.31 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.34 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.6 attack_power points (0.54 DPS) | yes | Sergeant Major's Cape (16315, -0.05 DPS) [pvp]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.19 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.14 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.63 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | Insignia Gloves (6408, -0.06 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; The Frozen Clutch (23170, -0.08 DPS) [dungeon] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 26.5 attack_power points (1.05 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.10 DPS) [rep]; Highlander's Leather Girdle (20117, -0.10 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.6 attack_power points (1.17 DPS) | yes | Brawler's Leather Legguards (252516, -0.23 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.24 DPS) [world_drop]; Ferine Leggings (6690, -0.41 DPS, sim-verified) [dungeon] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 21.6 attack_power points (0.86 DPS) | yes | Brawler's Leather Boots (252439, -0.24 DPS) [crafted]; Feet of the Lynx (1121, -0.27 DPS) [world_drop]; Alacritous Treads (277234, -0.45 DPS, sim-verified) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.3 attack_power points (0.76 DPS) | yes | Ironspine's Eye (7686, -0.06 DPS) [dungeon]; Tiger Band (6749, -0.29 DPS) [quest]; Silverlaine's Family Seal (6321, -0.37 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.5 attack_power points (0.73 DPS) | yes | Ironspine's Eye (7686, -0.03 DPS) [dungeon]; Tiger Band (6749, -0.26 DPS) [quest]; Silverlaine's Family Seal (6321, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (77.2 DPS) | yes | Morbid Dawn (7689, -0.04 DPS) [dungeon]; Corpsemaker (6687, -0.32 DPS) [dungeon]; Viscous Hammer (13045, -20.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.3 attack_power points (0.37 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.01 DPS) [vendor]; Long Battle Bow (15284, -0.13 DPS) [world_drop]; Silver Star (3463, -0.15 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 100.6. Weights run: 1.5s. Verify run: 1.6s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=0.709 ± 0.031, crit=0.521 ± 0.031 per rating point (14 rating = 1%, 7.296 per %), hit=2.038 ± 0.153 per rating point (10 rating = 1%, 20.383 per %), melee_haste=12.556 ± 1.422

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.3 attack_power points (1.87 DPS) | yes | Chromite Barbute (8142, -0.04 DPS) [dungeon]; White Bandit Mask (10008, -0.20 DPS) [crafted]; Hard Gold Coif (250537, -0.30 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.13 DPS) | yes | Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.52 DPS) [world_drop]; River Pride Choker (13087, -0.68 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.24 DPS) | yes | Chromite Pauldrons (8144, -0.11 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.11 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 16.3 attack_power points (0.92 DPS) | yes | Dark Hooded Cape (5257, -0.07 DPS) [world]; Hawkeye's Cloak (14593, -0.30 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.34 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 33.0 attack_power points (1.86 DPS) | yes | Avenger's Armor (1488, -0.17 DPS) [dungeon]; Quillward Harness (10583, -0.20 DPS) [dungeon]; Jouster's Chestplate (8157, -0.78 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.13 DPS) | yes | Ravager's Armguards (14770, -0.18 DPS) [world_drop]; Pugilist Bracers (4438, -0.23 DPS) [dungeon]; Yorgen Bracers (13012, -0.33 DPS) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.80 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.05 DPS) [dungeon]; Gloves of Holy Might (867, -0.26 DPS) [world_drop] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 30.0 attack_power points (1.69 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Ogron's Sash (13117, -0.32 DPS) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.36 DPS) | yes | Firemane Leggings (13129, -0.23 DPS) [world_drop]; Symbolic Legplates (14829, -0.44 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 31.0 attack_power points (1.74 DPS) | yes | Prowler's Leather Shoes (252465, -0.23 DPS) [crafted]; Blackforge Greaves (6423, -0.33 DPS) [dungeon]; Obsidian Greaves (13068, -0.38 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.7 attack_power points (1.22 DPS) | yes | Assault Band (13095, -0.09 DPS) [world_drop]; Thunderbrow Ring (13097, -0.20 DPS) [world_drop]; Ironspine's Eye (7686, -0.41 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.13 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Ironspine's Eye (7686, -0.32 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (100.6 DPS) | yes | Bonebiter (6830, -1.35 DPS) [quest]; Darkspear Raider's Reaper (272081, -1.72 DPS) [vendor]; The Jackhammer (9423, -2.18 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (100.6 DPS) | yes | The Silencer (13138, -0.01 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.19 DPS) [quest]; Bow of Searing Arrows (2825, -1.20 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 184.7. Weights run: 1.5s. Verify run: 1.9s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.460 ± 0.054, crit=1.662 ± 0.069 per rating point (14 rating = 1%, 23.270 per %), hit=2.383 ± 0.192 per rating point (10 rating = 1%, 23.834 per %), melee_haste=12.745 ± 1.456

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 83.1 attack_power points (6.81 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -0.97 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -2.77 DPS) [dungeon]; Embrace of the Lycan (9479, -2.88 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.0 attack_power points (2.37 DPS) | yes | Sentinel's Medallion (19539, -0.94 DPS) [rep]; Zealous Shadowshard Pendant (17772, -1.14 DPS, sim-verified) [quest] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 58.4 attack_power points (4.79 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, -0.99 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -1.70 DPS) [crafted]; Wyrmslayer Spaulders (13066, -1.86 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 32.4 attack_power points (2.66 DPS) | yes | Dark Hooded Cape (5257, -0.81 DPS) [world]; Blisterbane Wrap (12552, -0.86 DPS) [dungeon]; Dark Phantom Cape (13122, -0.86 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 53.3 attack_power points (4.37 DPS) | yes | Mixologist's Tunic (12793, -0.10 DPS) [dungeon]; Warbear Harness (15064, -0.41 DPS) [crafted]; Warforged Chestplate (11195, -0.43 DPS) [quest] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 30.8 attack_power points (2.52 DPS) | yes | Deepfury Bracers (13120, -0.07 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.21 DPS) [crafted]; Runed Golem Shackles (12550, -0.23 DPS) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 58.6 attack_power points (4.80 DPS) | yes | Raider Gloves (272100, -0.53 DPS) [vendor]; Officer's Gloves (250551, -1.10 DPS) [crafted]; Gloves of Holy Might (867, -1.26 DPS) [world_drop] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.77 DPS) | yes | Prowler's Leather Waistguard (252473, -0.04 DPS) [crafted]; Highlander's Lamellar Girdle (20106, -0.06 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 61.9 attack_power points (5.07 DPS) | yes | Centurion Legplates (10740, -0.82 DPS) [quest]; Serpentskin Leggings (8262, -1.03 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.10 DPS, sim-verified) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 47.0 attack_power points (3.85 DPS) | yes | Prowler's Leather Boots (252468, -0.40 DPS) [crafted]; Skulker's Leather Boots (252469, -0.49 DPS) [crafted]; Officer's Sabatons (250561, -0.64 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.8 attack_power points (3.59 DPS) | yes | Ironspine's Eye (7686, -1.86 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.92 DPS) [quest]; Thunderbrow Ring (13097, -1.92 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 33.1 attack_power points (2.72 DPS) | yes | Ironspine's Eye (7686, -0.98 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.04 DPS) [quest]; Thunderbrow Ring (13097, -1.05 DPS) [world_drop] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (184.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (184.7 DPS) | yes | Molten Heart of the Mountain (249470, -1.69 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (184.7 DPS) | yes | Blight (7959, +0.00 DPS) [crafted]; Warmonger (13052, -0.87 DPS) [world_drop]; Thorium Greatmace (250613, -1.07 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (184.7 DPS) | yes | Stinging Bow (10624, -0.04 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.04 DPS) [world_drop]; Dark Iron Rifle (16004, -2.14 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 261.9. Weights run: 1.6s. Verify run: 7.5s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.559 ± 0.063, crit=2.228 ± 0.091 per rating point (14 rating = 1%, 31.186 per %), hit=3.715 ± 0.250 per rating point (10 rating = 1%, 37.149 per %), melee_haste=17.614 ± 2.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 172.7 attack_power points (14.65 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -5.29 DPS) [vendor]; Mask of the Unforgiven (13404, -6.26 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (261.9 DPS) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Pendant of Celerity (22340, -0.05 DPS) [dungeon]; Mark of Fordring (15411, -0.34 DPS) [quest] |
| shoulder | Spaulders of Heroism (226858) | Anthion's Parting Words [quest] | sim-verified (261.9 DPS) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -5.28 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 65.1 attack_power points (5.53 DPS) | yes | Cape of the Black Baron (13340, -1.85 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.98 DPS) [rep]; Windshear Cape (20691, -2.19 DPS) [world] |
| chest | Breastplate of Heroism (226862) | Saving the Best for Last [quest] | sim-verified (261.9 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -1.37 DPS, sim-verified) [dungeon] |
| wrist | Bracers of Heroism (226863) | An Earnest Proposition [quest] | sim-verified (261.9 DPS) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Battleborn Armbraces (12936, -4.16 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Heroism (226861) | Just Compensation [quest] | sim-verified (261.9 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Savage Gladiator Grips (11730, -8.19 DPS, sim-verified) [dungeon] |
| waist | Belt of Heroism (226864) | Just Compensation [quest] | sim-verified (261.9 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (261.9 DPS) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Sentinel's Chain Leggings (237819, -0.97 DPS) [vendor]; Sentinel's Plate Legguards (237825, -1.02 DPS) [vendor] |
| feet | Battleboots of Heroism (226857) | Mokvar [vendor] | sim-verified (261.9 DPS) | yes | Boots of Heroism (21995, +0.00 DPS) [quest]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Savage Gladiator Greaves (11731, -4.53 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (261.9 DPS) | yes | Tarnished Elven Ring (18500, -2.02 DPS) [dungeon]; Cutthroat's Signet (272408, -2.15 DPS) [vendor]; Naglering (11669, -4.22 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (261.9 DPS) | yes | Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Cutthroat's Signet (272408, -0.53 DPS) [vendor]; Naglering (11669, -2.15 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (261.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -1.22 DPS, sim-verified) [dungeon] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (261.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (261.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -11.50 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (261.9 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.32 DPS) [dungeon]; The Purifier (22656, -0.90 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Spaulders of Heroism; back: Howler's Furs; chest: Breastplate of Heroism; wrist: Bracers of Heroism; hands: Gauntlets of Heroism; waist: Belt of Heroism; legs: Titanic Leggings; feet: Battleboots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Diamond Flask; main_hand: The Unstoppable Force; ranged: Satyr's Bow

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60, raid preset (human, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 770.9. Weights run: 1.4s. Verify run: 7.0s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.863 ± 0.094, crit=2.475 ± 0.126 per rating point (14 rating = 1%, 34.645 per %), hit=4.251 ± 0.385 per rating point (10 rating = 1%, 42.508 per %), melee_haste=15.560 ± 2.924

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 193.9 attack_power points (33.42 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -12.16 DPS) [vendor]; Mask of the Unforgiven (13404, -14.67 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (770.9 DPS) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Beads of Ogre Might (22150, -0.68 DPS) [quest]; Mark of Fordring (15411, -1.69 DPS) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) | The League of Arathor [rep] | sim-verified (770.9 DPS) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -15.69 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 70.5 attack_power points (12.15 DPS) | yes | Cape of the Black Baron (13340, -3.89 DPS) [dungeon]; Windshear Cape (20691, -4.30 DPS) [world]; Cloak of the Honor Guard (20073, -4.69 DPS) [rep] |
| chest | Savage Gladiator Chain (11726) | Blackrock Depths: Gorosh the Dervish [dungeon] | sim-verified (770.9 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Breastplate of Heroism (226862, -0.06 DPS) [quest]; Dawn Armor (252483, -4.91 DPS, sim-verified) [crafted] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (770.9 DPS) | yes | Forest Stalker's Bracers (19587, -3.03 DPS) [rep]; Berserker Bracers (19578, -3.52 DPS) [rep]; Bracers of Undead Slaying (23090, -5.47 DPS, sim-verified) [world] |
| hands | Savage Gladiator Grips (11730) | Blackrock Depths: Eviscerator [dungeon] | sim-verified (770.9 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Raider Gauntlets (272095, +0.00 DPS) [vendor] |
| waist | Highlander's Plate Girdle (20041) | The League of Arathor [rep] | sim-verified (770.9 DPS) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -8.79 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (770.9 DPS) | yes | Sentinel's Chain Leggings (237819, -1.49 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.49 DPS) [vendor]; Cloudkeeper Legplates (14554, -11.33 DPS, sim-verified) [world_drop] |
| feet | Highlander's Plate Greaves (20048) | The League of Arathor [rep] | sim-verified (770.9 DPS) | yes | Savage Gladiator Greaves (11731, +0.00 DPS) [dungeon]; Boots of Heroism (21995, +0.00 DPS) [quest]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (770.9 DPS) | yes | Tarnished Elven Ring (18500, -3.91 DPS) [dungeon]; Cutthroat's Signet (272408, -4.23 DPS) [vendor]; Naglering (11669, -9.22 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (770.9 DPS) | yes | Tarnished Elven Ring (18500, -0.96 DPS) [dungeon]; Cutthroat's Signet (272408, -1.28 DPS) [vendor]; Naglering (11669, -5.44 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (770.9 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -12.29 DPS, sim-verified) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (770.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Diamond Flask (20130, -4.95 DPS, sim-verified) [quest] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (770.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -10.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (770.9 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.73 DPS) [dungeon]; The Purifier (22656, -2.32 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Highlander's Plate Spaulders; back: Howler's Furs; chest: Savage Gladiator Chain; wrist: Battleborn Armbraces; hands: Savage Gladiator Grips; waist: Highlander's Plate Girdle; legs: Titanic Leggings; feet: Highlander's Plate Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Darkmoon Card: Maelstrom; main_hand: Blackblade of Shahram; ranged: Satyr's Bow

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 27.9. Weights run: 1.1s. Verify run: 2.0s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.727 ± 0.028, crit=0.316 ± 0.020 per rating point (14 rating = 1%, 4.418 per %), hit=1.247 ± 0.074 per rating point (10 rating = 1%, 12.474 per %), melee_haste=9.069 ± 0.670

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.45 DPS) [crafted]; Brawler's Leather Hood (252504, -0.50 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 4.4 attack_power points (0.15 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.08 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Veteran's Chain Shirt (250488, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.14 DPS) [crafted]; Brawler's Leather Armor (252490, -0.17 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Bristlebark Bindings (14569, -0.14 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.14 DPS) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.9 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.26 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Ruffian Belt (5975, -0.21 DPS) [world]; Brawler's Leather Belt (252428, -0.25 DPS) [crafted]; Cobrahn's Grasp (6460, -1.45 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.9 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.21 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (27.9 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -1.06 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.9 attack_power points (0.39 DPS) | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Signet of the Zhevra (285330, -0.23 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; Signet of the Zhevra (285330, -0.13 DPS) [world]; The 1 Ring (8350, -0.19 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.43 DPS) | yes | Forsaken Greataxe (251533, -0.67 DPS) [quest]; Smite's Mighty Hammer (7230, -0.75 DPS) [dungeon]; Hammerbone (270018, -0.80 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.04 DPS) [world_drop]; Heavy Shortbow (3036, -0.07 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 77.7. Weights run: 1.3s. Verify run: 1.4s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.086 ± 0.045, crit=0.543 ± 0.035 per rating point (14 rating = 1%, 7.605 per %), hit=1.755 ± 0.170 per rating point (10 rating = 1%, 17.552 per %), melee_haste=11.589 ± 1.312

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 27.8 attack_power points (1.10 DPS) | yes | Veteran's Chain Helm (250498, -0.15 DPS) [crafted]; Defender's Leather Helm (252455, -0.15 DPS) [crafted]; Tusken Helm (6686, -0.47 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.55 DPS) | yes | Kaleidoscope Chain (13084, -0.07 DPS) [world_drop]; Scout's Medallion (19537, -0.21 DPS) [rep]; River Pride Choker (13087, -0.24 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.9 attack_power points (0.87 DPS) | yes | Barbaric Shoulders (5964, -0.26 DPS) [crafted]; Golden Scale Shoulders (3841, -0.31 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.45 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.6 attack_power points (0.54 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.19 DPS) [world_drop]; Wildhunter Cloak (16658, -0.52 DPS, sim-verified) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.14 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.63 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | Insignia Gloves (6408, -0.06 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; The Frozen Clutch (23170, -0.08 DPS) [dungeon] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 26.5 attack_power points (1.05 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.10 DPS) [rep]; Defiler's Leather Girdle (20191, -0.10 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.6 attack_power points (1.17 DPS) | yes | Brawler's Leather Legguards (252516, -0.23 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.24 DPS) [world_drop]; Ferine Leggings (6690, -0.52 DPS, sim-verified) [dungeon] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 21.6 attack_power points (0.86 DPS) | yes | Feet of the Lynx (1121, -0.27 DPS) [world_drop]; Veteran's Boots (250503, -0.29 DPS) [crafted]; Brawler's Leather Boots (252439, -0.52 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.3 attack_power points (0.76 DPS) | yes | Ironspine's Eye (7686, -0.06 DPS) [dungeon]; Band of the Fist (17694, -0.27 DPS) [quest]; Tiger Band (6749, -0.29 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.5 attack_power points (0.73 DPS) | yes | Ironspine's Eye (7686, -0.03 DPS) [dungeon]; Band of the Fist (17694, -0.24 DPS) [quest]; Tiger Band (6749, -0.26 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (77.7 DPS) | yes | Morbid Dawn (7689, -0.04 DPS) [dungeon]; Corpsemaker (6687, -0.32 DPS) [dungeon]; Viscous Hammer (13045, -21.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.3 attack_power points (0.37 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.01 DPS) [vendor]; Long Battle Bow (15284, -0.13 DPS) [world_drop]; Silver Star (3463, -0.15 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 100.1. Weights run: 1.5s. Verify run: 1.5s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=0.709 ± 0.031, crit=0.521 ± 0.031 per rating point (14 rating = 1%, 7.296 per %), hit=2.038 ± 0.153 per rating point (10 rating = 1%, 20.383 per %), melee_haste=12.556 ± 1.422

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.3 attack_power points (1.87 DPS) | yes | Chromite Barbute (8142, -0.04 DPS) [dungeon]; White Bandit Mask (10008, -0.20 DPS) [crafted]; Hard Gold Coif (250537, -0.30 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.13 DPS) | yes | Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Ethereal Talisman (4430, -0.40 DPS) [quest]; Kaleidoscope Chain (13084, -0.52 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.24 DPS) | yes | Chromite Pauldrons (8144, -0.11 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.11 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 16.3 attack_power points (0.92 DPS) | yes | Dark Hooded Cape (5257, -0.07 DPS) [world]; Hawkeye's Cloak (14593, -0.30 DPS) [world_drop]; Wildhunter Cloak (16658, -0.35 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 33.0 attack_power points (1.86 DPS) | yes | Avenger's Armor (1488, -0.17 DPS) [dungeon]; Jouster's Chestplate (8157, -0.17 DPS) [dungeon]; Quillward Harness (10583, -0.20 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.13 DPS) | yes | Ravager's Armguards (14770, +0.00 DPS) [world_drop]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.80 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.05 DPS) [dungeon]; Gloves of Holy Might (867, -0.26 DPS) [world_drop] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 30.0 attack_power points (1.69 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.23 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.36 DPS) | yes | Firemane Leggings (13129, -0.23 DPS) [world_drop]; Symbolic Legplates (14829, -0.44 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 31.0 attack_power points (1.74 DPS) | yes | Prowler's Leather Shoes (252465, -0.23 DPS) [crafted]; Blackforge Greaves (6423, -0.33 DPS) [dungeon]; Obsidian Greaves (13068, -0.38 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.7 attack_power points (1.22 DPS) | yes | Assault Band (13095, -0.09 DPS) [world_drop]; Thunderbrow Ring (13097, -0.20 DPS) [world_drop]; Ironspine's Eye (7686, -0.41 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.13 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Ironspine's Eye (7686, -0.32 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (100.1 DPS) | yes | The Jackhammer (9423, -1.62 DPS, sim-verified) [dungeon]; Darkspear Raider's Reaper (272081, -1.72 DPS) [vendor]; Primitive Fishing Pole (276203, -1.84 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (100.1 DPS) | yes | The Silencer (13138, -0.01 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.19 DPS) [quest]; Bow of Searing Arrows (2825, -1.03 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 190.5. Weights run: 1.5s. Verify run: 1.9s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.460 ± 0.054, crit=1.662 ± 0.069 per rating point (14 rating = 1%, 23.270 per %), hit=2.383 ± 0.192 per rating point (10 rating = 1%, 23.834 per %), melee_haste=12.745 ± 1.456

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 83.1 attack_power points (6.81 DPS) | yes | Blood Guard's Plate Helm (220803, -0.82 DPS) [vendor]; Embrace of the Lycan (9479, -2.88 DPS) [dungeon]; Raging Berserker's Helm (7719, -3.20 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.0 attack_power points (2.37 DPS) | yes | Woven Ivy Necklace (19159, -0.31 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.74 DPS) [quest]; Scout's Medallion (19535, -0.94 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 58.4 attack_power points (4.79 DPS) | yes | Blood Guard's Plate Pauldrons (220796, -1.08 DPS) [vendor]; Wyrmslayer Spaulders (13066, -1.86 DPS) [world_drop]; Officer's Pauldrons (250576, -1.92 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 32.4 attack_power points (2.66 DPS) | yes | Blisterbane Wrap (12552, -0.86 DPS) [dungeon]; Dark Phantom Cape (13122, -0.86 DPS) [world_drop]; Dark Hooded Cape (5257, -0.95 DPS, sim-verified) [world] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 53.3 attack_power points (4.37 DPS) | yes | Mixologist's Tunic (12793, -0.10 DPS) [dungeon]; Warbear Harness (15064, -0.41 DPS) [crafted]; Warforged Chestplate (11195, -0.43 DPS) [quest] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 30.8 attack_power points (2.52 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 58.6 attack_power points (4.80 DPS) | yes | Raider Gloves (272100, -1.02 DPS, sim-verified) [vendor]; Officer's Gloves (250551, -1.10 DPS) [crafted]; Gloves of Holy Might (867, -1.26 DPS) [world_drop] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.77 DPS) | yes | Prowler's Leather Waistguard (252473, -0.04 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted]; Defiler's Plate Girdle (20205, -0.22 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 61.9 attack_power points (5.07 DPS) | yes | Serpentskin Leggings (8262, -1.24 DPS, sim-verified) [world_drop]; Stormshroud Pants (15057, -1.26 DPS) [crafted]; Golem Shard Leggings (13074, -1.46 DPS) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 47.0 attack_power points (3.85 DPS) | yes | Prowler's Leather Boots (252468, -0.40 DPS) [crafted]; Skulker's Leather Boots (252469, -0.49 DPS) [crafted]; Officer's Sabatons (250561, -0.64 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.8 attack_power points (3.59 DPS) | yes | White Bone Band (11862, -1.63 DPS) [quest]; Ironspine's Eye (7686, -1.86 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.92 DPS) [quest] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 33.1 attack_power points (2.72 DPS) | yes | White Bone Band (11862, -0.75 DPS) [quest]; Ironspine's Eye (7686, -0.98 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.04 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -3.05 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.76 DPS, sim-verified) [crafted] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (190.5 DPS) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -2.05 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.04 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.04 DPS) [world_drop]; Dark Iron Rifle (16004, -1.96 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Blight; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 261.4. Weights run: 1.6s. Verify run: 7.6s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.559 ± 0.063, crit=2.228 ± 0.091 per rating point (14 rating = 1%, 31.186 per %), hit=3.715 ± 0.250 per rating point (10 rating = 1%, 37.149 per %), melee_haste=17.614 ± 2.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 172.7 attack_power points (14.65 DPS) | yes | Champion's Plate Helm (227043, -5.29 DPS) [pvp]; Fury Visor (20521, -5.80 DPS) [quest]; Mask of the Unforgiven (13404, -6.13 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (261.4 DPS) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Pendant of Celerity (22340, -0.05 DPS) [dungeon]; Mark of Fordring (15411, -0.34 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 98.3 attack_power points (8.34 DPS) | yes | Warlord's Plate Shoulders (231534, -0.02 DPS) [pvp]; Darkspear Pauldrons (272105, -2.12 DPS) [vendor]; Wyrmhide Spaulders (12082, -2.51 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 65.1 attack_power points (5.53 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -1.98 DPS) [rep]; Windshear Cape (20691, -2.19 DPS) [world] |
| chest | Breastplate of Heroism (226862) | Saving the Best for Last [quest] | sim-verified (261.4 DPS) | yes | Savage Gladiator Chain (11726, +0.00 DPS) [dungeon]; Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted] |
| wrist | Bracers of Heroism (226863) | An Earnest Proposition [quest] | sim-verified (261.4 DPS) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Battleborn Armbraces (12936, -4.73 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Heroism (226861) | Just Compensation [quest] | sim-verified (261.4 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Savage Gladiator Grips (11730, -8.85 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (261.4 DPS) | yes | Ferocity of the Timbermaw (227805, -0.57 DPS) [vendor]; Marksman's Girdle (22232, -1.30 DPS) [dungeon]; Belt of Preserved Heads (20216, -2.84 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (261.4 DPS) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Sentinel's Chain Leggings (237819, -0.97 DPS) [vendor]; Sentinel's Plate Legguards (237825, -1.02 DPS) [vendor] |
| feet | Battleboots of Heroism (226857) | Mokvar [vendor] | sim-verified (261.4 DPS) | yes | Boots of Heroism (21995, +0.00 DPS) [quest]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Savage Gladiator Greaves (11731, -4.35 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (261.4 DPS) | yes | Tarnished Elven Ring (18500, -2.02 DPS) [dungeon]; Cutthroat's Signet (272408, -2.15 DPS) [vendor]; Naglering (11669, -4.07 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (261.4 DPS) | yes | Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Cutthroat's Signet (272408, -0.53 DPS) [vendor]; Naglering (11669, -2.06 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (261.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Hand of Justice (11815, -1.24 DPS, sim-verified) [dungeon] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (261.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (261.4 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -11.31 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (261.4 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.32 DPS) [dungeon]; The Purifier (22656, -0.90 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Breastplate of Heroism; wrist: Bracers of Heroism; hands: Gauntlets of Heroism; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Battleboots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Satyr's Bow

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (orc, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 774.4. Weights run: 1.4s. Verify run: 7.2s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.863 ± 0.094, crit=2.475 ± 0.126 per rating point (14 rating = 1%, 34.645 per %), hit=4.251 ± 0.385 per rating point (10 rating = 1%, 42.508 per %), melee_haste=15.560 ± 2.924

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 193.9 attack_power points (33.42 DPS) | yes | Champion's Plate Helm (227043, -12.16 DPS) [pvp]; Fury Visor (20521, -13.30 DPS) [quest]; Mask of the Unforgiven (13404, -15.32 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (774.4 DPS) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Beads of Ogre Might (22150, -0.68 DPS) [quest]; Mark of Fordring (15411, -1.69 DPS) [quest] |
| shoulder | Spaulders of Heroism (226858) | Anthion's Parting Words [quest] | sim-verified (774.4 DPS) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -17.76 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 70.5 attack_power points (12.15 DPS) | yes | Cape of the Black Baron (13340, -3.89 DPS) [dungeon]; Windshear Cape (20691, -4.30 DPS) [world]; Deathguard's Cloak (20068, -4.69 DPS) [rep] |
| chest | Breastplate of Heroism (226862) | Saving the Best for Last [quest] | sim-verified (774.4 DPS) | yes | Savage Gladiator Chain (11726, +0.00 DPS) [dungeon]; Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted] |
| wrist | Bracers of Heroism (226863) | An Earnest Proposition [quest] | sim-verified (774.4 DPS) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Battleborn Armbraces (12936, -11.24 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Heroism (226861) | Just Compensation [quest] | sim-verified (774.4 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Raider Gauntlets (272095, +0.00 DPS) [vendor]; Savage Gladiator Grips (11730, -19.54 DPS, sim-verified) [dungeon] |
| waist | Belt of Heroism (226864) | Just Compensation [quest] | sim-verified (774.4 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (774.4 DPS) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Sentinel's Chain Leggings (237819, -1.49 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.49 DPS) [vendor] |
| feet | Battleboots of Heroism (226857) | Mokvar [vendor] | sim-verified (774.4 DPS) | yes | Boots of Heroism (21995, +0.00 DPS) [quest]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Savage Gladiator Greaves (11731, -10.58 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (774.4 DPS) | yes | Tarnished Elven Ring (18500, -3.91 DPS) [dungeon]; Cutthroat's Signet (272408, -4.23 DPS) [vendor]; Naglering (11669, -10.09 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (774.4 DPS) | yes | Tarnished Elven Ring (18500, -0.96 DPS) [dungeon]; Cutthroat's Signet (272408, -1.28 DPS) [vendor]; Naglering (11669, -6.34 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (774.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (774.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Diamond Flask (20130, -14.29 DPS, sim-verified) [quest] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (774.4 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -15.81 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (774.4 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.73 DPS) [dungeon]; The Purifier (22656, -2.32 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Spaulders of Heroism; back: Howler's Furs; chest: Breastplate of Heroism; wrist: Bracers of Heroism; hands: Gauntlets of Heroism; waist: Belt of Heroism; legs: Titanic Leggings; feet: Battleboots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Blackblade of Shahram; ranged: Satyr's Bow

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

