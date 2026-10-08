# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.8. Weights run: 1.7s. Verify run: 2.9s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=1.247 ± 0.074 per rating point (10 rating = 1%, 12.474 per %), melee_haste=9.069 ± 0.670

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.14 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.21 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.8 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Polar Gauntlets (7606, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, -1.22 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Ruffian Belt (5975, -0.21 DPS) [world]; Hulking Belt (14746, -0.28 DPS) [world_drop]; Cobrahn's Grasp (6460, -1.51 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.8 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Chausses of Westfall (6087, -1.01 DPS, sim-verified) [quest] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (26.8 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, -1.01 DPS, sim-verified) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.43 DPS) | yes | Living Root (6631, -0.89 DPS) [dungeon]; Duskbringer (2205, -0.98 DPS) [dungeon]; Smite's Mighty Hammer (7230, -1.10 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567)) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, -0.00 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Protector's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 71.9. Weights run: 2.0s. Verify run: 2.1s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.084 ± 0.010, crit=0.119 ± 0.015 per rating point (14 rating = 1%, 1.671 per %), hit=1.777 ± 0.170 per rating point (10 rating = 1%, 17.770 per %), melee_haste=11.743 ± 1.316

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.03 DPS) | yes | Veteran's Chain Helm (250498, -0.08 DPS) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.16 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.55 DPS) | yes | River Pride Choker (13087, -0.24 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.35 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.53 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.55 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.06 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.08 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.40 DPS) | yes | Sergeant Major's Cape (16315, -0.07 DPS) [pvp]; Lambent Scale Cloak (4706, -0.08 DPS) [world_drop]; Slayer's Cape (14752, -0.08 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted]; Hard Gold Cuirass (250533, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.63 DPS) | yes | Yorgen Bracers (13012, -0.15 DPS) [world_drop]; Bands of Serra'kis (6902, -0.16 DPS) [dungeon]; Cultist's Armguards (270032, -0.24 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | The Frozen Clutch (23170, -0.08 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.16 DPS) [world]; Mail Combat Gauntlets (4075, -0.22 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (0.95 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.03 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.13 DPS) [crafted]; Golden Scale Leggings (3843, -0.16 DPS) [crafted]; Chausses of Westfall (6087, -0.16 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.6 attack_power points (0.58 DPS) | yes | Hard Gold Boots (250534, -0.02 DPS) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.3 attack_power points (0.64 DPS) | yes | Tiger Band (6749, -0.17 DPS) [quest]; Silverlaine's Family Seal (6321, -0.25 DPS) [dungeon]; Insurgent's Band (272067, -0.29 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.5 attack_power points (0.49 DPS) | yes | Tiger Band (6749, -0.02 DPS) [quest]; Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.14 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (71.9 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.12 DPS) [dungeon]; Viscous Hammer (13045, -19.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.36 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.12 DPS) [world_drop]; Precision Bow (8183, -0.14 DPS) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 97.2. Weights run: 2.3s. Verify run: 2.1s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=0.199 ± 0.015, crit=0.285 ± 0.021 per rating point (14 rating = 1%, 3.986 per %), hit=2.011 ± 0.153 per rating point (10 rating = 1%, 20.107 per %), melee_haste=12.432 ± 1.429

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.0 attack_power points (1.68 DPS) | yes | Icemetal Barbute (10763, -0.11 DPS) [dungeon]; Hard Gold Coif (250537, -0.11 DPS) [crafted]; Chromite Barbute (8142, -0.20 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.12 DPS) | yes | Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.63 DPS) [world_drop]; River Pride Choker (13087, -0.67 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.23 DPS) | yes | Chromite Pauldrons (8144, -0.11 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.11 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.22 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 13.2 attack_power points (0.74 DPS) | yes | Wolfmaster Cape (6314, -0.18 DPS) [dungeon]; Dark Hooded Cape (5257, -0.18 DPS) [world]; Slayer's Cape (14752, -0.29 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.68 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.03 DPS) [quest]; Shining Mithril Breastplate (250540, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Ravager's Armguards (14770, -0.29 DPS) [world_drop]; Yorgen Bracers (13012, -0.41 DPS) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.79 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.34 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.34 DPS) [dungeon] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 30.0 attack_power points (1.68 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Scarlet Belt (10329, -0.34 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.35 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.4 attack_power points (1.54 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.30 DPS) [crafted]; Obsidian Greaves (13068, -0.35 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Protector's Band (19515, -0.13 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.34 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.12 DPS) | yes | Protector's Band (19515, -0.13 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.34 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (97.2 DPS) | yes | Bonebiter (6830, -1.34 DPS) [quest]; Darkspear Raider's Reaper (272081, -1.72 DPS) [vendor]; The Jackhammer (9423, -2.46 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (97.2 DPS) | yes | Monolithic Bow (9426, -0.08 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.22 DPS) [crafted]; Bow of Searing Arrows (2825, -0.82 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: The Silencer

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 180.5. Weights run: 2.3s. Verify run: 2.8s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.169 ± 0.049, crit=1.670 ± 0.071 per rating point (14 rating = 1%, 23.385 per %), hit=2.432 ± 0.185 per rating point (10 rating = 1%, 24.324 per %), melee_haste=12.497 ± 1.433

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 83.7 attack_power points (6.79 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -0.81 DPS) [vendor]; Raging Berserker's Helm (7719, -2.78 DPS) [dungeon]; Embrace of the Lycan (9479, -2.89 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.2 attack_power points (2.04 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.42 DPS) [quest]; Sentinel's Medallion (19539, -0.91 DPS) [rep]; Ghostshard Talisman (7731, -0.91 DPS) [dungeon] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 56.0 attack_power points (4.54 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, -0.86 DPS) [vendor]; Officer's Pauldrons (250576, -1.67 DPS) [crafted]; Wyrmslayer Spaulders (13066, -1.84 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.4 attack_power points (2.30 DPS) | yes | Sergeant Major's Cape (16336, -0.76 DPS) [pvp]; Bloodlust Cape (14801, -0.84 DPS) [world_drop]; Dark Hooded Cape (5257, -1.09 DPS, sim-verified) [world] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 53.4 attack_power points (4.33 DPS) | yes | Mixologist's Tunic (12793, -0.37 DPS) [dungeon]; Warforged Chestplate (11195, -0.44 DPS) [quest]; Warbear Harness (15064, -0.84 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.0 attack_power points (2.35 DPS) | yes | Runed Golem Shackles (12550, -0.08 DPS) [dungeon]; Bracers of the Stone Princess (17714, -0.08 DPS) [dungeon]; Arena Bands (18711, -0.08 DPS) [world] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 55.7 attack_power points (4.51 DPS) | yes | Raider Gloves (272100, -0.81 DPS) [vendor]; Gloves of Holy Might (867, -1.00 DPS) [world_drop]; Officer's Gloves (250551, -1.07 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.73 DPS) | yes | Highlander's Lamellar Girdle (20106, -0.05 DPS) [rep]; Highlander's Plate Girdle (20124, -0.21 DPS) [rep]; Highlander's Chain Girdle (20088, -0.21 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 59.1 attack_power points (4.79 DPS) | yes | Centurion Legplates (10740, -0.94 DPS) [quest]; Stormshroud Pants (15057, -1.00 DPS) [crafted]; Gryphon Rider's Leggings (9652, -1.53 DPS, sim-verified) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 43.2 attack_power points (3.50 DPS) | yes | Prowler's Leather Boots (252468, -0.35 DPS) [crafted]; Skulker's Leather Boots (252469, -0.49 DPS) [crafted]; Officer's Sabatons (250561, -0.54 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.3 attack_power points (3.59 DPS) | yes | Mark of Kern (2262, -1.97 DPS) [dungeon]; Assault Band (13095, -1.97 DPS) [world_drop]; Thunderbrow Ring (13097, -2.01 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 30.5 attack_power points (2.47 DPS) | yes | Mark of Kern (2262, -0.85 DPS) [dungeon]; Assault Band (13095, -0.85 DPS) [world_drop]; Thunderbrow Ring (13097, -0.89 DPS) [world_drop] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (180.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (180.5 DPS) | yes | Molten Heart of the Mountain (249470, -1.44 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (180.5 DPS) | yes | Blight (7959, +0.00 DPS, sim-verified) [crafted]; Warmonger (13052, -0.74 DPS) [world_drop]; Thorium Greatmace (250613, -1.06 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (180.5 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.88 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 261.9. Weights run: 2.4s. Verify run: 10.9s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.559 ± 0.063, crit=2.228 ± 0.091 per rating point (14 rating = 1%, 31.186 per %), hit=3.715 ± 0.250 per rating point (10 rating = 1%, 37.149 per %), melee_haste=17.614 ± 2.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 172.7 attack_power points (14.65 DPS) | yes | Lieutenant Commander's Plate Helm (227044, -5.29 DPS) [vendor]; Mask of the Unforgiven (13404, -6.26 DPS, sim-verified) [dungeon] |
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

Set DPS (verified): 792.1. Weights run: 2.1s. Verify run: 10.2s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.901 ± 0.088, crit=2.440 ± 0.120 per rating point (14 rating = 1%, 34.162 per %), hit=4.420 ± 0.364 per rating point (10 rating = 1%, 44.203 per %), melee_haste=16.652 ± 2.739

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 196.3 attack_power points (35.16 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -12.85 DPS) [vendor]; Mask of the Unforgiven (13404, -15.40 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (792.1 DPS) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Beads of Ogre Might (22150, -0.81 DPS) [quest]; Mark of Fordring (15411, -2.25 DPS) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) | The League of Arathor [rep] | sim-verified (792.1 DPS) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -14.51 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.2 attack_power points (12.93 DPS) | yes | Cape of the Black Baron (13340, -4.24 DPS) [dungeon]; Windshear Cape (20691, -4.67 DPS) [world]; Stalwart Cloak (272415, -5.01 DPS) [vendor] |
| chest | Savage Gladiator Chain (11726) | Blackrock Depths: Gorosh the Dervish [dungeon] | sim-verified (792.1 DPS) | yes | Breastplate of Heroism (226862, +0.00 DPS) [quest]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -4.56 DPS, sim-verified) [crafted] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (792.1 DPS) | yes | Forest Stalker's Bracers (19587, -3.23 DPS) [rep]; Slashclaw Bracers (13211, -3.73 DPS) [dungeon]; Bracers of Undead Slaying (23090, -6.19 DPS, sim-verified) [world] |
| hands | Savage Gladiator Grips (11730) | Blackrock Depths: Eviscerator [dungeon] | sim-verified (792.1 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Raider Gauntlets (272095, +0.00 DPS) [vendor] |
| waist | Highlander's Plate Girdle (20041) | The League of Arathor [rep] | sim-verified (792.1 DPS) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -9.59 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (792.1 DPS) | yes | Sentinel's Chain Leggings (237819, -1.70 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.98 DPS) [vendor]; Cloudkeeper Legplates (14554, -10.45 DPS, sim-verified) [world_drop] |
| feet | Highlander's Plate Greaves (20048) | The League of Arathor [rep] | sim-verified (792.1 DPS) | yes | Savage Gladiator Greaves (11731, +0.00 DPS) [dungeon]; Boots of Heroism (21995, +0.00 DPS) [quest]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (792.1 DPS) | yes | Tarnished Elven Ring (18500, -3.88 DPS) [dungeon]; Cutthroat's Signet (272408, -4.22 DPS) [vendor]; Naglering (11669, -10.04 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (792.1 DPS) | yes | Tarnished Elven Ring (18500, -1.02 DPS) [dungeon]; Cutthroat's Signet (272408, -1.36 DPS) [vendor]; Naglering (11669, -6.13 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (792.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Diamond Flask (20130, -11.51 DPS, sim-verified) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (792.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (792.1 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -11.37 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (792.1 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.79 DPS) [dungeon]; The Purifier (22656, -2.82 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Highlander's Plate Spaulders; back: Howler's Furs; chest: Savage Gladiator Chain; wrist: Battleborn Armbraces; hands: Savage Gladiator Grips; waist: Highlander's Plate Girdle; legs: Titanic Leggings; feet: Highlander's Plate Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Darkmoon Card: Maelstrom; main_hand: Blackblade of Shahram; ranged: Satyr's Bow

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 27.2. Weights run: 1.7s. Verify run: 2.9s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=1.247 ± 0.074 per rating point (10 rating = 1%, 12.474 per %), melee_haste=9.069 ± 0.670

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.22 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Raptorcrest Bracers (270010, -0.15 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.2 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Foreman's Gloves (2167, -0.07 DPS) [world]; Thorbia's Gauntlets (12994, -1.24 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Ruffian Belt (5975, -0.21 DPS) [world]; Hulking Belt (14746, -0.28 DPS) [world_drop]; Cobrahn's Grasp (6460, -1.53 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.2 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.17 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (27.2 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, -1.02 DPS, sim-verified) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.43 DPS) | yes | Forsaken Greataxe (251533, -0.67 DPS) [quest]; Hammerbone (270018, -0.78 DPS, sim-verified) [quest]; Smite's Mighty Hammer (7230, -0.85 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Legionnaire's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.5. Weights run: 2.0s. Verify run: 1.9s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.084 ± 0.010, crit=0.119 ± 0.015 per rating point (14 rating = 1%, 1.671 per %), hit=1.777 ± 0.170 per rating point (10 rating = 1%, 17.770 per %), melee_haste=11.743 ± 1.316

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.03 DPS) | yes | Veteran's Chain Helm (250498, -0.08 DPS) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.16 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.55 DPS) | yes | River Pride Choker (13087, -0.24 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.35 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.53 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.55 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.06 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.40 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.08 DPS) [world_drop]; Slayer's Cape (14752, -0.08 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted]; Hard Gold Cuirass (250533, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.63 DPS) | yes | Yorgen Bracers (13012, -0.15 DPS) [world_drop]; Bands of Serra'kis (6902, -0.16 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.23 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | The Frozen Clutch (23170, -0.08 DPS) [dungeon]; Warsong Gauntlets (16978, -0.08 DPS) [quest]; Bonefist Gauntlets (4465, -0.16 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (0.95 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.03 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.13 DPS) [crafted]; Golden Scale Leggings (3843, -0.16 DPS) [crafted]; Slayer's Pants (14757, -0.16 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.6 attack_power points (0.58 DPS) | yes | Hard Gold Boots (250534, -0.02 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop]; Slayer's Slippers (14756, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.3 attack_power points (0.64 DPS) | yes | Tiger Band (6749, -0.17 DPS) [quest]; Silverlaine's Family Seal (6321, -0.25 DPS) [dungeon]; Insurgent's Band (272067, -0.29 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.5 attack_power points (0.49 DPS) | yes | Tiger Band (6749, -0.02 DPS) [quest]; Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.14 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.5 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.12 DPS) [dungeon]; Viscous Hammer (13045, -19.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.36 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.12 DPS) [world_drop]; Precision Bow (8183, -0.14 DPS) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 97.1. Weights run: 2.3s. Verify run: 2.1s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=0.199 ± 0.015, crit=0.285 ± 0.021 per rating point (14 rating = 1%, 3.986 per %), hit=2.011 ± 0.153 per rating point (10 rating = 1%, 20.107 per %), melee_haste=12.432 ± 1.429

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.0 attack_power points (1.68 DPS) | yes | Icemetal Barbute (10763, -0.11 DPS) [dungeon]; Hard Gold Coif (250537, -0.11 DPS) [crafted]; Chromite Barbute (8142, -0.20 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.12 DPS) | yes | Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Ethereal Talisman (4430, -0.52 DPS) [quest]; Kaleidoscope Chain (13084, -0.63 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.23 DPS) | yes | Chromite Pauldrons (8144, -0.11 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.11 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.22 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 13.2 attack_power points (0.74 DPS) | yes | Wolfmaster Cape (6314, -0.18 DPS) [dungeon]; Wildhunter Cloak (16658, -0.18 DPS) [quest]; Dark Hooded Cape (5257, -0.18 DPS) [world] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.68 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.03 DPS) [quest]; Shining Mithril Breastplate (250540, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.79 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.34 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.34 DPS) [dungeon] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 30.0 attack_power points (1.68 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.22 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.35 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.4 attack_power points (1.54 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.30 DPS) [crafted]; Obsidian Greaves (13068, -0.35 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Legionnaire's Band (19512, -0.13 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.34 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.12 DPS) | yes | Legionnaire's Band (19512, -0.13 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.34 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (97.1 DPS) | yes | Darkspear Raider's Reaper (272081, -1.72 DPS) [vendor]; The Jackhammer (9423, -1.83 DPS, sim-verified) [dungeon]; Primitive Fishing Pole (276203, -2.12 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (97.1 DPS) | yes | Monolithic Bow (9426, -0.08 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.22 DPS) [crafted]; Bow of Searing Arrows (2825, -0.83 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: The Silencer

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 185.3. Weights run: 2.3s. Verify run: 2.7s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.169 ± 0.049, crit=1.670 ± 0.071 per rating point (14 rating = 1%, 23.385 per %), hit=2.432 ± 0.185 per rating point (10 rating = 1%, 24.324 per %), melee_haste=12.497 ± 1.433

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 83.7 attack_power points (6.79 DPS) | yes | Blood Guard's Plate Helm (220803, -0.81 DPS) [vendor]; Embrace of the Lycan (9479, -2.89 DPS) [dungeon]; Raging Berserker's Helm (7719, -3.59 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.2 attack_power points (2.04 DPS) | yes | Woven Ivy Necklace (19159, -0.22 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.42 DPS) [quest]; Ethereal Talisman (4430, -0.85 DPS) [quest] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 56.0 attack_power points (4.54 DPS) | yes | Blood Guard's Plate Pauldrons (220796, -0.86 DPS) [vendor]; Wyrmslayer Spaulders (13066, -1.84 DPS) [world_drop]; Officer's Pauldrons (250576, -2.38 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.4 attack_power points (2.30 DPS) | yes | Dark Hooded Cape (5257, -0.70 DPS) [world]; First Sergeant's Cloak (16340, -0.76 DPS) [pvp]; Bloodlust Cape (14801, -0.84 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 53.4 attack_power points (4.33 DPS) | yes | Mixologist's Tunic (12793, -0.37 DPS) [dungeon]; Warforged Chestplate (11195, -0.44 DPS) [quest]; Warbear Harness (15064, -0.84 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.0 attack_power points (2.35 DPS) | yes | Runed Golem Shackles (12550, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 55.7 attack_power points (4.51 DPS) | yes | Raider Gloves (272100, -0.81 DPS) [vendor]; Gloves of Holy Might (867, -1.00 DPS) [world_drop]; Officer's Gloves (250551, -1.07 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.73 DPS) | yes | Defiler's Plate Girdle (20205, -0.21 DPS) [rep]; Defiler's Chain Girdle (20151, -0.21 DPS) [rep]; Defiler's Leather Girdle (20193, -0.21 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 59.1 attack_power points (4.79 DPS) | yes | Serpentskin Leggings (8262, -1.16 DPS) [world_drop]; Stormshroud Pants (15057, -1.20 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -1.22 DPS) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 43.2 attack_power points (3.50 DPS) | yes | Prowler's Leather Boots (252468, -0.35 DPS) [crafted]; Skulker's Leather Boots (252469, -0.49 DPS) [crafted]; Officer's Sabatons (250561, -0.54 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.3 attack_power points (3.59 DPS) | yes | White Bone Band (11862, -1.65 DPS) [quest]; Mark of Kern (2262, -1.97 DPS) [dungeon]; Assault Band (13095, -1.97 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 30.5 attack_power points (2.47 DPS) | yes | White Bone Band (11862, -0.53 DPS) [quest]; Mark of Kern (2262, -0.85 DPS) [dungeon]; Assault Band (13095, -0.85 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (185.3 DPS) | yes | Frozen Heart of the Mountain (249469, -3.01 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (185.3 DPS) | yes | Frozen Heart of the Mountain (249469, -2.89 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (185.3 DPS) | yes | Blight (7959, +0.00 DPS, sim-verified) [crafted]; Warmonger (13052, -0.74 DPS) [world_drop]; Thorium Greatmace (250613, -1.06 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (185.3 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.91 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Glowing Brightwood Staff; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 261.4. Weights run: 2.4s. Verify run: 11.0s. 1913 eligible items had no known source.

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

Set DPS (verified): 790.0. Weights run: 2.1s. Verify run: 10.1s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.901 ± 0.088, crit=2.440 ± 0.120 per rating point (14 rating = 1%, 34.162 per %), hit=4.420 ± 0.364 per rating point (10 rating = 1%, 44.203 per %), melee_haste=16.652 ± 2.739

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 196.3 attack_power points (35.16 DPS) | yes | Champion's Plate Helm (227043, -12.85 DPS) [pvp]; Helm of the Executioner (22411, -13.81 DPS) [dungeon]; Mask of the Unforgiven (13404, -15.95 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (790.0 DPS) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Beads of Ogre Might (22150, -0.81 DPS) [quest]; Mark of Fordring (15411, -2.25 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 112.4 attack_power points (20.13 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -4.82 DPS) [vendor]; Wyrmhide Spaulders (12082, -5.81 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.2 attack_power points (12.93 DPS) | yes | Cape of the Black Baron (13340, -4.24 DPS) [dungeon]; Windshear Cape (20691, -4.67 DPS) [world]; Stalwart Cloak (272415, -5.01 DPS) [vendor] |
| chest | Savage Gladiator Chain (11726) | Blackrock Depths: Gorosh the Dervish [dungeon] | sim-verified (790.0 DPS) | yes | Breastplate of Heroism (226862, +0.00 DPS) [quest]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -9.19 DPS, sim-verified) [crafted] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (790.0 DPS) | yes | Forest Stalker's Bracers (19587, -3.23 DPS) [rep]; Slashclaw Bracers (13211, -3.73 DPS) [dungeon]; Bracers of Undead Slaying (23090, -6.13 DPS, sim-verified) [world] |
| hands | Savage Gladiator Grips (11730) | Blackrock Depths: Eviscerator [dungeon] | sim-verified (790.0 DPS) | yes | Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Raider Gauntlets (272095, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -7.36 DPS, sim-verified) [quest] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (790.0 DPS) | yes | Ferocity of the Timbermaw (227805, -0.72 DPS) [vendor]; Marksman's Girdle (22232, -1.69 DPS) [dungeon]; Belt of Preserved Heads (20216, -7.37 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (790.0 DPS) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Sentinel's Chain Leggings (237819, -1.70 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.98 DPS) [vendor] |
| feet | Savage Gladiator Greaves (11731) | Blackrock Depths: Anub'shiah [dungeon] | sim-verified (790.0 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Boots of Heroism (21995, -3.78 DPS, sim-verified) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (790.0 DPS) | yes | Tarnished Elven Ring (18500, -3.88 DPS) [dungeon]; Cutthroat's Signet (272408, -4.22 DPS) [vendor]; Naglering (11669, -9.98 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (790.0 DPS) | yes | Tarnished Elven Ring (18500, -1.02 DPS) [dungeon]; Cutthroat's Signet (272408, -1.36 DPS) [vendor]; Naglering (11669, -6.10 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (790.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (790.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Diamond Flask (20130, -11.92 DPS, sim-verified) [quest] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (790.0 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -14.70 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (790.0 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.79 DPS) [dungeon]; The Purifier (22656, -2.82 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Savage Gladiator Chain; wrist: Battleborn Armbraces; hands: Savage Gladiator Grips; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Savage Gladiator Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Blackblade of Shahram; ranged: Satyr's Bow

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

