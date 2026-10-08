# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.3. Weights run: 1.8s. Verify run: 1.1s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=1.241 ± 0.074 per rating point (10 rating = 1%, 12.414 per %), melee_haste=9.065 ± 0.670

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.14 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.21 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.57 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Polar Gauntlets (7606, -0.14 DPS) [quest]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Hulking Belt (14746, -0.28 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.78 DPS) | yes | Defender's Leather Pants (252445, -0.14 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.14 DPS, sim-verified) [crafted] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.43 DPS) | yes | Living Root (6631, -0.89 DPS) [dungeon]; Duskbringer (2205, -0.98 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567)) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, -0.00 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Protector's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.0. Weights run: 2.1s. Verify run: 1.4s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.085 ± 0.011, crit=0.121 ± 0.015 per rating point (14 rating = 1%, 1.697 per %), hit=1.845 ± 0.169 per rating point (10 rating = 1%, 18.449 per %), melee_haste=11.000 ± 1.318

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
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.0 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.12 DPS) [dungeon]; Viscous Hammer (13045, -19.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.36 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.12 DPS) [world_drop]; Precision Bow (8183, -0.14 DPS) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 97.3. Weights run: 2.3s. Verify run: 1.4s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=0.203 ± 0.015, crit=0.291 ± 0.021 per rating point (14 rating = 1%, 4.070 per %), hit=2.094 ± 0.152 per rating point (10 rating = 1%, 20.945 per %), melee_haste=12.534 ± 1.424

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.1 attack_power points (1.69 DPS) | yes | Icemetal Barbute (10763, -0.12 DPS) [dungeon]; Hard Gold Coif (250537, -0.12 DPS) [crafted]; Chromite Barbute (8142, -0.20 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.12 DPS) | yes | Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.63 DPS) [world_drop]; River Pride Choker (13087, -0.67 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.23 DPS) | yes | Chromite Pauldrons (8144, -0.11 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.11 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.22 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 13.2 attack_power points (0.74 DPS) | yes | Dark Hooded Cape (5257, -0.18 DPS) [world]; Wolfmaster Cape (6314, -0.18 DPS) [dungeon]; Slayer's Cape (14752, -0.29 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.68 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.03 DPS) [quest]; Shining Mithril Breastplate (250540, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Ravager's Armguards (14770, -0.29 DPS) [world_drop]; Yorgen Bracers (13012, -0.41 DPS) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.79 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.33 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.34 DPS) [world_drop] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 30.0 attack_power points (1.68 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Scarlet Belt (10329, -0.34 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.35 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted]; Symbolic Legplates (14829, -0.60 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.4 attack_power points (1.54 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.30 DPS) [crafted]; Obsidian Greaves (13068, -0.35 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Protector's Band (19515, -0.13 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.34 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.12 DPS) | yes | Protector's Band (19515, -0.13 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.34 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (97.3 DPS) | yes | Bonebiter (6830, -1.34 DPS) [quest]; Darkspear Raider's Reaper (272081, -1.72 DPS) [vendor]; The Jackhammer (9423, -2.22 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (97.3 DPS) | yes | Monolithic Bow (9426, -0.08 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.22 DPS) [crafted]; Bow of Searing Arrows (2825, -0.82 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: The Silencer

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 180.5. Weights run: 2.4s. Verify run: 1.6s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.167 ± 0.050, crit=1.667 ± 0.071 per rating point (14 rating = 1%, 23.335 per %), hit=2.325 ± 0.184 per rating point (10 rating = 1%, 23.250 per %), melee_haste=13.511 ± 1.447

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 82.6 attack_power points (6.69 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -0.81 DPS) [vendor]; Raging Berserker's Helm (7719, -2.69 DPS) [dungeon]; Embrace of the Lycan (9479, -2.80 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.2 attack_power points (2.04 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.42 DPS) [quest]; Sentinel's Medallion (19539, -0.90 DPS) [rep]; Ghostshard Talisman (7731, -0.90 DPS) [dungeon] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 54.9 attack_power points (4.45 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, -0.78 DPS) [vendor]; Officer's Pauldrons (250576, -1.59 DPS) [crafted]; Wyrmslayer Spaulders (13066, -1.75 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.3 attack_power points (2.30 DPS) | yes | Sergeant Major's Cape (16336, -0.76 DPS) [pvp]; Bloodlust Cape (14801, -0.84 DPS) [world_drop]; Dark Hooded Cape (5257, -1.02 DPS, sim-verified) [world] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 53.3 attack_power points (4.32 DPS) | yes | Mixologist's Tunic (12793, -0.36 DPS) [dungeon]; Warforged Chestplate (11195, -0.43 DPS) [quest]; Warbear Harness (15064, -0.84 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.0 attack_power points (2.35 DPS) | yes | Runed Golem Shackles (12550, -0.08 DPS) [dungeon]; Bracers of the Stone Princess (17714, -0.08 DPS) [dungeon]; Arena Bands (18711, -0.08 DPS) [world] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 55.7 attack_power points (4.51 DPS) | yes | Raider Gloves (272100, -0.81 DPS) [vendor]; Gloves of Holy Might (867, -1.00 DPS) [world_drop]; Officer's Gloves (250551, -1.07 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.73 DPS) | yes | Highlander's Lamellar Girdle (20106, -0.05 DPS) [rep]; Highlander's Plate Girdle (20124, -0.22 DPS) [rep]; Highlander's Chain Girdle (20088, -0.22 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 59.0 attack_power points (4.78 DPS) | yes | Centurion Legplates (10740, -0.93 DPS) [quest]; Stormshroud Pants (15057, -1.00 DPS) [crafted]; Gryphon Rider's Leggings (9652, -1.39 DPS, sim-verified) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 43.2 attack_power points (3.50 DPS) | yes | Prowler's Leather Boots (252468, -0.35 DPS) [crafted]; Skulker's Leather Boots (252469, -0.49 DPS) [crafted]; Officer's Sabatons (250561, -0.54 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.2 attack_power points (3.50 DPS) | yes | Mark of Kern (2262, -1.88 DPS) [dungeon]; Assault Band (13095, -1.88 DPS) [world_drop]; Thunderbrow Ring (13097, -1.92 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 30.5 attack_power points (2.47 DPS) | yes | Mark of Kern (2262, -0.85 DPS) [dungeon]; Assault Band (13095, -0.85 DPS) [world_drop]; Thunderbrow Ring (13097, -0.89 DPS) [world_drop] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (180.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (180.5 DPS) | yes | Molten Heart of the Mountain (249470, -1.61 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (180.5 DPS) | yes | Blight (7959, +0.00 DPS, sim-verified) [crafted]; Warmonger (13052, -1.00 DPS) [world_drop]; Thorium Greatmace (250613, -1.05 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (180.5 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.88 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 263.1. Weights run: 2.5s. Verify run: 1.7s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.581 ± 0.063, crit=2.258 ± 0.090 per rating point (14 rating = 1%, 31.616 per %), hit=3.450 ± 0.250 per rating point (10 rating = 1%, 34.500 per %), melee_haste=17.154 ± 2.042

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 168.2 attack_power points (14.27 DPS) | yes | Fury Visor (20521, -2.57 DPS, sim-verified) [quest]; Lieutenant Commander's Plate Helm (23314, -5.10 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Pendant of Celerity (22340, -0.02 DPS) [dungeon]; Mark of Fordring (15411, -0.08 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -3.18 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 62.5 attack_power points (5.30 DPS) | yes | Cape of the Black Baron (13340, -1.59 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.75 DPS) [rep]; Windshear Cape (20691, -1.93 DPS) [world] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -0.12 DPS) [crafted]; Obsidian Mail Tunic (22191, -1.57 DPS) [crafted]; Breastplate of Undead Slaying (23087, -11.54 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.19 DPS) [rep]; Berserker Bracers (19578, -1.31 DPS) [rep]; Bracers of Undead Slaying (23090, -2.57 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; Marshal's Plate Gauntlets (231541, -0.98 DPS) [pvp]; Raider Gauntlets (272095, -1.00 DPS) [vendor] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.57 DPS) [vendor]; Marshal's Plate Girdle (16482, -1.41 DPS) [pvp]; Belt of Preserved Heads (20216, -3.04 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Sentinel's Chain Leggings (237819, -0.64 DPS) [vendor]; Sentinel's Plate Legguards (237825, -0.75 DPS) [vendor] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 74.5 attack_power points (6.32 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -0.66 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.03 DPS) [dungeon]; Cutthroat's Signet (272408, -2.16 DPS) [vendor]; Naglering (11669, -4.26 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Cutthroat's Signet (272408, -0.54 DPS) [vendor]; Naglering (11669, -2.22 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+9.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Diamond Flask (20130, -3.06 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -11.10 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.29 DPS) [dungeon]; The Purifier (22656, -0.65 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force; ranged: Satyr's Bow

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60, raid preset (human, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 683.6. Weights run: 2.2s. Verify run: 1.5s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.618 ± 0.076, crit=2.311 ± 0.108 per rating point (14 rating = 1%, 32.356 per %), hit=4.182 ± 0.324 per rating point (10 rating = 1%, 41.820 per %), melee_haste=14.377 ± 2.512

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 184.4 attack_power points (32.64 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -12.07 DPS) [vendor]; Mask of the Unforgiven (13404, -13.67 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -0.05 DPS) [quest]; Mark of Fordring (15411, -1.37 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 107.6 attack_power points (19.06 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, -0.70 DPS) [pvp]; Razorsteel Shoulders (20517, -5.25 DPS) [quest]; Wyrmhide Spaulders (12082, -5.65 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 69.8 attack_power points (12.36 DPS) | yes | Cape of the Black Baron (13340, -4.52 DPS) [dungeon]; Cloak of the Honor Guard (20073, -4.91 DPS) [rep]; Stalwart Cloak (272415, -4.96 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -1.16 DPS) [crafted]; Obsidian Mail Tunic (22191, -4.55 DPS) [crafted]; Breastplate of Undead Slaying (23087, -15.21 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -3.72 DPS) [dungeon]; Forest Stalker's Bracers (19587, -3.80 DPS) [rep]; Bracers of Undead Slaying (23090, -5.25 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; Stormshroud Gloves (21278, -4.25 DPS) [crafted]; Marshal's Plate Gauntlets (231541, -4.58 DPS) [pvp] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (683.6 DPS) | yes | Ferocity of the Timbermaw (227805, -1.21 DPS) [vendor]; Marksman's Girdle (22232, -1.87 DPS) [dungeon]; Belt of Preserved Heads (20216, -7.39 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Sentinel's Chain Leggings (237819, -2.27 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.74 DPS) [vendor] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 81.8 attack_power points (14.49 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.32 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -4.27 DPS) [dungeon]; Cutthroat's Signet (272408, -4.55 DPS) [vendor]; Naglering (11669, -8.99 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.86 DPS) [dungeon]; Cutthroat's Signet (272408, -1.15 DPS) [vendor]; Naglering (11669, -5.16 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Diamond Flask (20130, -7.18 DPS, sim-verified) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -3.81 DPS, sim-verified) [quest] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -5.73 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.74 DPS) [dungeon]; The Purifier (22656, -2.54 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Darkmoon Card: Maelstrom; main_hand: Blackblade of Shahram; ranged: Satyr's Bow

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.5. Weights run: 1.8s. Verify run: 1.2s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=1.241 ± 0.074 per rating point (10 rating = 1%, 12.414 per %), melee_haste=9.065 ± 0.670

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.22 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Raptorcrest Bracers (270010, -0.14 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.57 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon]; Foreman's Gloves (2167, -0.21 DPS) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Hulking Belt (14746, -0.28 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.0 attack_power points (0.64 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.07 DPS) [world_drop] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.43 DPS) | yes | Hammerbone (270018, -0.66 DPS, sim-verified) [quest]; Forsaken Greataxe (251533, -0.67 DPS) [quest]; Smite's Mighty Hammer (7230, -0.85 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Legionnaire's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.8. Weights run: 2.1s. Verify run: 1.4s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.085 ± 0.011, crit=0.121 ± 0.015 per rating point (14 rating = 1%, 1.697 per %), hit=1.845 ± 0.169 per rating point (10 rating = 1%, 18.449 per %), melee_haste=11.000 ± 1.318

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.03 DPS) | yes | Veteran's Chain Helm (250498, -0.08 DPS) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.16 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.55 DPS) | yes | River Pride Choker (13087, -0.24 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.36 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.53 DPS) [rep] |
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
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.8 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.12 DPS) [dungeon]; Viscous Hammer (13045, -20.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.36 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.12 DPS) [world_drop]; Precision Bow (8183, -0.14 DPS) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 97.5. Weights run: 2.3s. Verify run: 1.4s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=0.203 ± 0.015, crit=0.291 ± 0.021 per rating point (14 rating = 1%, 4.070 per %), hit=2.094 ± 0.152 per rating point (10 rating = 1%, 20.945 per %), melee_haste=12.534 ± 1.424

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.1 attack_power points (1.69 DPS) | yes | Icemetal Barbute (10763, -0.12 DPS) [dungeon]; Hard Gold Coif (250537, -0.12 DPS) [crafted]; Chromite Barbute (8142, -0.20 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.12 DPS) | yes | Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Ethereal Talisman (4430, -0.51 DPS) [quest]; Kaleidoscope Chain (13084, -0.63 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.23 DPS) | yes | Chromite Pauldrons (8144, -0.11 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.11 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.22 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 13.2 attack_power points (0.74 DPS) | yes | Dark Hooded Cape (5257, -0.18 DPS) [world]; Wolfmaster Cape (6314, -0.18 DPS) [dungeon]; Wildhunter Cloak (16658, -0.18 DPS) [quest] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.68 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.03 DPS) [quest]; Shining Mithril Breastplate (250540, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.79 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.33 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.34 DPS) [world_drop] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 30.0 attack_power points (1.68 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.22 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.35 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted]; Symbolic Legplates (14829, -0.60 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.4 attack_power points (1.54 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.30 DPS) [crafted]; Obsidian Greaves (13068, -0.35 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Legionnaire's Band (19512, -0.13 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.34 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.12 DPS) | yes | Legionnaire's Band (19512, -0.13 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.34 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (97.5 DPS) | yes | Darkspear Raider's Reaper (272081, -1.72 DPS) [vendor]; Primitive Fishing Pole (276203, -2.11 DPS) [vendor]; The Jackhammer (9423, -2.14 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (97.5 DPS) | yes | Monolithic Bow (9426, -0.08 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.22 DPS) [crafted]; Bow of Searing Arrows (2825, -0.83 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: The Silencer

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 184.6. Weights run: 2.4s. Verify run: 1.6s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.167 ± 0.050, crit=1.667 ± 0.071 per rating point (14 rating = 1%, 23.335 per %), hit=2.325 ± 0.184 per rating point (10 rating = 1%, 23.250 per %), melee_haste=13.511 ± 1.447

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 82.6 attack_power points (6.69 DPS) | yes | Blood Guard's Plate Helm (220803, -0.81 DPS) [vendor]; Embrace of the Lycan (9479, -2.80 DPS) [dungeon]; Raging Berserker's Helm (7719, -3.66 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.2 attack_power points (2.04 DPS) | yes | Woven Ivy Necklace (19159, -0.22 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.42 DPS) [quest]; Ethereal Talisman (4430, -0.85 DPS) [quest] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 54.9 attack_power points (4.45 DPS) | yes | Blood Guard's Plate Pauldrons (220796, -0.78 DPS) [vendor]; Wyrmslayer Spaulders (13066, -1.75 DPS) [world_drop]; Officer's Pauldrons (250576, -2.46 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.3 attack_power points (2.30 DPS) | yes | Dark Hooded Cape (5257, -0.70 DPS) [world]; First Sergeant's Cloak (16340, -0.76 DPS) [pvp]; Bloodlust Cape (14801, -0.84 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 53.3 attack_power points (4.32 DPS) | yes | Mixologist's Tunic (12793, -0.36 DPS) [dungeon]; Warforged Chestplate (11195, -0.43 DPS) [quest]; Warbear Harness (15064, -0.84 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.0 attack_power points (2.35 DPS) | yes | Runed Golem Shackles (12550, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 55.7 attack_power points (4.51 DPS) | yes | Raider Gloves (272100, -0.81 DPS) [vendor]; Gloves of Holy Might (867, -1.00 DPS) [world_drop]; Officer's Gloves (250551, -1.07 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.73 DPS) | yes | Defiler's Plate Girdle (20205, -0.22 DPS) [rep]; Defiler's Chain Girdle (20151, -0.22 DPS) [rep]; Defiler's Leather Girdle (20193, -0.22 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 59.0 attack_power points (4.78 DPS) | yes | Stormshroud Pants (15057, -1.14 DPS, sim-verified) [crafted]; Serpentskin Leggings (8262, -1.16 DPS) [world_drop]; Golem Shard Leggings (13074, -1.22 DPS) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 43.2 attack_power points (3.50 DPS) | yes | Prowler's Leather Boots (252468, -0.35 DPS) [crafted]; Skulker's Leather Boots (252469, -0.49 DPS) [crafted]; Officer's Sabatons (250561, -0.54 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.2 attack_power points (3.50 DPS) | yes | White Bone Band (11862, -1.56 DPS) [quest]; Mark of Kern (2262, -1.88 DPS) [dungeon]; Assault Band (13095, -1.88 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 30.5 attack_power points (2.47 DPS) | yes | White Bone Band (11862, -0.53 DPS) [quest]; Mark of Kern (2262, -0.85 DPS) [dungeon]; Assault Band (13095, -0.85 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (184.6 DPS) | yes | Frozen Heart of the Mountain (249469, -3.03 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (184.6 DPS) | yes | Frozen Heart of the Mountain (249469, -2.41 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (184.6 DPS) | yes | Blight (7959, +0.00 DPS, sim-verified) [crafted]; Warmonger (13052, -1.00 DPS) [world_drop]; Thorium Greatmace (250613, -1.05 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (184.6 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.90 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Glowing Brightwood Staff; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 261.1. Weights run: 2.5s. Verify run: 1.6s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.581 ± 0.063, crit=2.258 ± 0.090 per rating point (14 rating = 1%, 31.616 per %), hit=3.450 ± 0.250 per rating point (10 rating = 1%, 34.500 per %), melee_haste=17.154 ± 2.042

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 168.2 attack_power points (14.27 DPS) | yes | Fury Visor (20521, -2.91 DPS, sim-verified) [quest]; Champion's Plate Helm (227043, -5.10 DPS) [pvp]; Mask of the Unforgiven (13404, -5.73 DPS) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Pendant of Celerity (22340, -0.02 DPS) [dungeon]; Mark of Fordring (15411, -0.08 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -3.13 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 62.5 attack_power points (5.30 DPS) | yes | Cape of the Black Baron (13340, -1.59 DPS) [dungeon]; Deathguard's Cloak (20068, -1.75 DPS) [rep]; Windshear Cape (20691, -1.93 DPS) [world] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -0.12 DPS) [crafted]; Obsidian Mail Tunic (22191, -1.57 DPS) [crafted]; Breastplate of Undead Slaying (23087, -12.32 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.19 DPS) [rep]; Berserker Bracers (19578, -1.31 DPS) [rep]; Bracers of Undead Slaying (23090, -2.91 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; General's Plate Gauntlets (231532, -0.98 DPS) [pvp]; Raider Gauntlets (272095, -1.00 DPS) [vendor] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.57 DPS) [vendor]; General's Plate Girdle (16547, -1.41 DPS) [pvp]; Belt of Preserved Heads (20216, -3.23 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Sentinel's Chain Leggings (237819, -0.64 DPS) [vendor]; Sentinel's Plate Legguards (237825, -0.75 DPS) [vendor] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 74.5 attack_power points (6.32 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -0.66 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.03 DPS) [dungeon]; Cutthroat's Signet (272408, -2.16 DPS) [vendor]; Naglering (11669, -4.58 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Cutthroat's Signet (272408, -0.54 DPS) [vendor]; Naglering (11669, -2.64 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS, sim-verified) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -9.53 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.29 DPS) [dungeon]; The Purifier (22656, -0.65 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Satyr's Bow

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (orc, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 686.5. Weights run: 2.2s. Verify run: 1.5s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.618 ± 0.076, crit=2.311 ± 0.108 per rating point (14 rating = 1%, 32.356 per %), hit=4.182 ± 0.324 per rating point (10 rating = 1%, 41.820 per %), melee_haste=14.377 ± 2.512

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 184.4 attack_power points (32.64 DPS) | yes | Champion's Plate Helm (227043, -12.07 DPS) [pvp]; Helm of the Executioner (22411, -12.88 DPS) [dungeon]; Mask of the Unforgiven (13404, -14.13 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -0.05 DPS) [quest]; Mark of Fordring (15411, -1.37 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 107.6 attack_power points (19.06 DPS) | yes | Warlord's Plate Shoulders (231534, -0.70 DPS) [pvp]; Razorsteel Shoulders (20517, -5.25 DPS) [quest]; Wyrmhide Spaulders (12082, -5.72 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 69.8 attack_power points (12.36 DPS) | yes | Cape of the Black Baron (13340, -4.52 DPS) [dungeon]; Deathguard's Cloak (20068, -4.91 DPS) [rep]; Stalwart Cloak (272415, -4.96 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -1.16 DPS) [crafted]; Obsidian Mail Tunic (22191, -4.55 DPS) [crafted]; Breastplate of Undead Slaying (23087, -13.88 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -3.72 DPS) [dungeon]; Forest Stalker's Bracers (19587, -3.80 DPS) [rep]; Bracers of Undead Slaying (23090, -5.62 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; Stormshroud Gloves (21278, -4.25 DPS) [crafted]; General's Plate Gauntlets (231532, -4.58 DPS) [pvp] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (686.5 DPS) | yes | Ferocity of the Timbermaw (227805, -1.21 DPS) [vendor]; Marksman's Girdle (22232, -1.87 DPS) [dungeon]; Belt of Preserved Heads (20216, -7.32 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Sentinel's Chain Leggings (237819, -2.27 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.74 DPS) [vendor] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 81.8 attack_power points (14.49 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.32 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -4.27 DPS) [dungeon]; Cutthroat's Signet (272408, -4.55 DPS) [vendor]; Naglering (11669, -9.40 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.86 DPS) [dungeon]; Cutthroat's Signet (272408, -1.15 DPS) [vendor]; Naglering (11669, -4.88 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+22.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Diamond Flask (20130, -7.65 DPS, sim-verified) [quest] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -13.06 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.74 DPS) [dungeon]; The Purifier (22656, -2.54 DPS) [quest] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Darkmoon Card: Maelstrom; main_hand: Blackblade of Shahram; ranged: Satyr's Bow

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

