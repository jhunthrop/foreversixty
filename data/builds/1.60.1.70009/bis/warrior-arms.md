# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.3. Weights run: 2.4s. Verify run: 1.2s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.001, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.049 ± 0.001 per rating point (10 rating = 1%, 0.492 per %), melee_haste=1.606 ± 0.044

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Grave Shroud (279865, -0.07 DPS) [quest]; Miner's Cape (5444, -0.14 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.21 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.55 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Polar Gauntlets (7606, -0.14 DPS) [quest]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Hulking Belt (14746, -0.28 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.76 DPS) | yes | Defender's Leather Pants (252445, -0.14 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.14 DPS, sim-verified) [crafted] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.16 DPS) | yes | Living Root (6631, -0.87 DPS) [dungeon]; Duskbringer (2205, -0.96 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.14 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Protector's Band; main_hand: The Axe of Severing; ranged: Fine Longbow

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.0. Weights run: 2.8s. Verify run: 1.4s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.021 ± 0.002, crit=0.030 ± 0.002 per rating point (14 rating = 1%, 0.416 per %), hit=0.130 ± 0.006 per rating point (10 rating = 1%, 1.300 per %), melee_haste=not significant (4.380 ± 1.332)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.95 DPS) | yes | Veteran's Chain Helm (250498, -0.07 DPS) [crafted]; Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Crusader's Chain Helm (250502, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.51 DPS) | yes | River Pride Choker (13087, -0.22 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.35 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.51 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.51 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.07 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.07 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.37 DPS) | yes | Sergeant Major's Cape (16315, -0.07 DPS) [pvp]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.10 DPS) | yes | Shining Silver Breastplate (2870, -0.07 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.22 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.59 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS) [world_drop]; Bands of Serra'kis (6902, -0.15 DPS) [dungeon]; Cultist's Armguards (270032, -0.22 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.81 DPS) | yes | The Frozen Clutch (23170, -0.07 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.15 DPS) [world]; Mail Combat Gauntlets (4075, -0.22 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (0.88 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.95 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.14 DPS) [crafted]; Golden Scale Leggings (3843, -0.15 DPS) [crafted]; Chausses of Westfall (6087, -0.15 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.1 attack_power points (0.52 DPS) | yes | Hard Gold Boots (250534, -0.01 DPS) [crafted]; Disjointed Shoes (277226, -0.08 DPS) [quest]; Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.1 attack_power points (0.59 DPS) | yes | Tiger Band (6749, -0.15 DPS) [quest]; Silverlaine's Family Seal (6321, -0.22 DPS) [dungeon]; Insurgent's Band (272067, -0.26 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.1 attack_power points (0.44 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.08 DPS) [dungeon]; Insurgent's Band (272067, -0.11 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.0 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.10 DPS) [dungeon]; Viscous Hammer (13045, -19.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.33 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.11 DPS) [world_drop]; Fine Longbow (11304, -0.18 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 97.2. Weights run: 3.0s. Verify run: 1.4s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.055 ± 0.003, crit=0.078 ± 0.004 per rating point (14 rating = 1%, 1.099 per %), hit=0.175 ± 0.006 per rating point (10 rating = 1%, 1.754 per %), melee_haste=not significant (3.233 ± 1.300)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 28.0 attack_power points (1.23 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.04 DPS) [dungeon]; Tusken Helm (6686, -0.09 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.88 DPS) | yes | Ghostshard Talisman (7731, -0.26 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.52 DPS) [world_drop]; River Pride Choker (13087, -0.53 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.96 DPS) | yes | Chromite Pauldrons (8144, -0.09 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.18 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 12.3 attack_power points (0.54 DPS) | yes | Wolfmaster Cape (6314, -0.10 DPS) [dungeon]; Dark Hooded Cape (5257, -0.17 DPS) [world]; Slayer's Cape (14752, -0.19 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.32 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.07 DPS) [quest]; Shining Mithril Breastplate (250540, -0.09 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.88 DPS) | yes | Pugilist Bracers (4438, -0.18 DPS) [dungeon]; Ravager's Armguards (14770, -0.25 DPS) [world_drop]; Yorgen Bracers (13012, -0.34 DPS) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.40 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.26 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.33 DPS) [dungeon] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 30.0 attack_power points (1.32 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Scarlet Belt (10329, -0.26 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.84 DPS) | yes | Firemane Leggings (13129, -0.18 DPS) [world_drop]; Orcish War Leggings (7929, -0.35 DPS) [crafted]; Symbolic Legplates (14829, -0.51 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 26.4 attack_power points (1.16 DPS) | yes | Prowler's Leather Shoes (252465, -0.18 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.19 DPS) [crafted]; Obsidian Greaves (13068, -0.27 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.88 DPS) | yes | Protector's Band (19515, -0.16 DPS) [rep]; Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Suspicious Spare Part (274754, -0.26 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.88 DPS) | yes | Protector's Band (19515, -0.16 DPS) [rep]; Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Suspicious Spare Part (274754, -0.26 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (97.2 DPS) | yes | Bonebiter (6830, -1.05 DPS) [quest]; Darkspear Raider's Reaper (272081, -1.34 DPS) [vendor]; The Jackhammer (9423, -2.21 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (97.2 DPS) | yes | Monolithic Bow (9426, -0.08 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.18 DPS) [crafted]; Bow of Searing Arrows (2825, -0.82 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: The Silencer

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 177.5. Weights run: 3.2s. Verify run: 1.7s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.631 ± 0.040, crit=0.902 ± 0.057 per rating point (14 rating = 1%, 12.626 per %), hit=0.224 ± 0.007 per rating point (10 rating = 1%, 2.241 per %), melee_haste=7.075 ± 0.975

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | sim-verified (177.5 DPS) | yes | Raging Berserker's Helm (7719, -0.11 DPS) [dungeon]; Fury Visor (20521, -0.18 DPS) [quest]; Embrace of the Lycan (9479, -2.01 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.97 DPS) | yes | Skibi's Pendant (13089, -0.09 DPS) [world_drop]; Ghostshard Talisman (7731, -0.29 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.46 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 34.6 attack_power points (1.69 DPS) | yes | Wyrmslayer Spaulders (13066, -0.27 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.34 DPS) [crafted]; Officer's Pauldrons (250576, -1.07 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.8 attack_power points (1.02 DPS) | yes | Sergeant Major's Cape (16336, -0.25 DPS) [pvp]; Dark Hooded Cape (5257, -0.32 DPS) [world]; Bloodlust Cape (14801, -1.03 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 48.0 attack_power points (2.34 DPS) | yes | Mixologist's Tunic (12793, -0.25 DPS) [dungeon]; Knight's Plate Hauberk (220794, -0.26 DPS) [vendor]; Valorous Chestguard (8274, -0.39 DPS) [world_drop] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.36 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Runed Golem Shackles (12550, -0.00 DPS) [dungeon]; Officer's Wristguards (250581, -0.11 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 50.3 attack_power points (2.45 DPS) | yes | Raider Gloves (272100, -0.80 DPS) [vendor]; Prowler's Leather Gauntlets (252547, -0.81 DPS) [crafted]; Officer's Gloves (250551, -1.17 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (2.24 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.49 DPS) [dungeon]; Belt of the Gladiator (13134, -0.49 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.51 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 44.0 attack_power points (2.14 DPS) | yes | Knight's Plate Leggings (220797, +0.00 DPS, sim-verified) [vendor]; Scarlet Leggings (10330, -0.10 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.19 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 36.2 attack_power points (1.76 DPS) | yes | Prowler's Leather Boots (252468, -0.16 DPS) [crafted]; Officer's Sabatons (250561, -0.22 DPS) [crafted]; Officer's Boots (250546, -0.28 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 25.7 attack_power points (1.25 DPS) | yes | Mark of Kern (2262, -0.28 DPS) [dungeon]; Assault Band (13095, -0.28 DPS) [world_drop]; Thunderbrow Ring (13097, -0.38 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 22.2 attack_power points (1.08 DPS) | yes | Assault Band (13095, -0.11 DPS) [world_drop]; Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Mark of Kern (2262, -2.39 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.95 DPS, sim-verified) [crafted] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -2.25 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.10 DPS) [dungeon]; Dark Iron Rifle (16004, -1.76 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 264.9. Weights run: 3.3s. Verify run: 1.5s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.961 ± 0.058, crit=1.373 ± 0.083 per rating point (14 rating = 1%, 19.218 per %), hit=0.317 ± 0.010 per rating point (10 rating = 1%, 3.172 per %), melee_haste=12.382 ± 1.379

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 80.8 attack_power points (3.92 DPS) | yes | Field Marshal's Plate Helm (231538, -0.27 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -0.80 DPS) [vendor]; Crown of Heroism (226860, -10.36 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (264.9 DPS) | yes | Amulet of the Darkmoon (19491, -0.24 DPS) [quest]; Imperial Jewel (11933, -0.54 DPS) [dungeon]; Rage of Mugamba (19577, -2.76 DPS, sim-verified) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 52.3 attack_power points (2.54 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 38.8 attack_power points (1.88 DPS) | yes | Cape of the Black Baron (13340, -0.21 DPS) [dungeon]; Shroud of Domination (22337, -0.23 DPS) [dungeon]; Howler's Furs (272414, -0.37 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (264.9 DPS) | yes | Obsidian Mail Tunic (22191, -0.35 DPS) [crafted]; Cadaverous Armor (14637, -0.91 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -11.50 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-verified (264.9 DPS) | yes | Marshal's Plate Bracers (16481, -0.24 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.26 DPS) [rep]; Bracers of Undead Slaying (23090, -5.14 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (264.9 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.26 DPS) [pvp]; Radiant Gloves of the Dawn (227817, -0.44 DPS) [vendor]; Razor Gauntlets (18326, -5.79 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 73.2 attack_power points (3.56 DPS) | yes | Ferocity of the Timbermaw (227805, -0.24 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.45 DPS) [vendor]; Marshal's Plate Girdle (16482, -0.57 DPS) [pvp] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (264.9 DPS) | yes | Titanic Leggings (22385, -0.49 DPS) [crafted]; Marshal's Plate Legguards (231540, -0.68 DPS) [pvp]; Cloudkeeper Legplates (14554, -6.15 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 45.7 attack_power points (2.22 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Boots of Heroism (21995, -0.12 DPS) [quest]; Battleboots of Heroism (226857, -0.12 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (264.9 DPS) | yes | Band of the Ogre King (18522, -0.50 DPS) [dungeon]; Myrmidon's Signet (2246, -0.57 DPS) [world_drop]; Naglering (11669, -8.12 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (264.9 DPS) | yes | Band of the Ogre King (18522, -0.32 DPS) [dungeon]; Myrmidon's Signet (2246, -0.38 DPS) [world_drop]; Naglering (11669, -4.36 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (264.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (264.9 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.83 DPS) [crafted]; Diamond Flask (20130, -2.85 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (264.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -11.62 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-verified (264.9 DPS) | yes | Riphook (12653, -0.04 DPS) [dungeon]; Malgen's Long Bow (22318, -0.13 DPS) [dungeon]; Dark Iron Rifle (16004, -2.84 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Highlander's Plate Spaulders; back: Cloak of the Honor Guard; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force; ranged: Bloodseeker

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.5. Weights run: 2.4s. Verify run: 1.1s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.001, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.049 ± 0.001 per rating point (10 rating = 1%, 0.492 per %), melee_haste=1.606 ± 0.044

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.22 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Raptorcrest Bracers (270010, -0.14 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.55 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon]; Foreman's Gloves (2167, -0.21 DPS) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Hulking Belt (14746, -0.28 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.0 attack_power points (0.62 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.07 DPS) [world_drop] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.16 DPS) | yes | Forsaken Greataxe (251533, -0.66 DPS) [quest]; Hammerbone (270018, -0.66 DPS, sim-verified) [quest]; Smite's Mighty Hammer (7230, -0.83 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.14 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Legionnaire's Band; main_hand: The Axe of Severing; ranged: Fine Longbow

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.8. Weights run: 2.8s. Verify run: 1.4s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.021 ± 0.002, crit=0.030 ± 0.002 per rating point (14 rating = 1%, 0.416 per %), hit=0.130 ± 0.006 per rating point (10 rating = 1%, 1.300 per %), melee_haste=not significant (4.380 ± 1.332)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.95 DPS) | yes | Veteran's Chain Helm (250498, -0.07 DPS) [crafted]; Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Crusader's Chain Helm (250502, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.51 DPS) | yes | River Pride Choker (13087, -0.22 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.36 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.51 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.51 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.07 DPS) [crafted]; Elite Shoulders (4835, -0.07 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.37 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.10 DPS) | yes | Shining Silver Breastplate (2870, -0.07 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.22 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.59 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS) [world_drop]; Bands of Serra'kis (6902, -0.15 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.22 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.81 DPS) | yes | The Frozen Clutch (23170, -0.07 DPS) [dungeon]; Warsong Gauntlets (16978, -0.07 DPS) [quest]; Bonefist Gauntlets (4465, -0.15 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (0.88 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.95 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.14 DPS) [crafted]; Golden Scale Leggings (3843, -0.15 DPS) [crafted]; Slayer's Pants (14757, -0.15 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.1 attack_power points (0.52 DPS) | yes | Hard Gold Boots (250534, -0.01 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.1 attack_power points (0.59 DPS) | yes | Tiger Band (6749, -0.15 DPS) [quest]; Silverlaine's Family Seal (6321, -0.22 DPS) [dungeon]; Insurgent's Band (272067, -0.26 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.1 attack_power points (0.44 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.08 DPS) [dungeon]; Insurgent's Band (272067, -0.11 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.8 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.10 DPS) [dungeon]; Viscous Hammer (13045, -20.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.33 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.11 DPS) [world_drop]; Fine Longbow (11304, -0.18 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 97.3. Weights run: 3.0s. Verify run: 1.4s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.055 ± 0.003, crit=0.078 ± 0.004 per rating point (14 rating = 1%, 1.099 per %), hit=0.175 ± 0.006 per rating point (10 rating = 1%, 1.754 per %), melee_haste=not significant (3.233 ± 1.300)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 28.0 attack_power points (1.23 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.04 DPS) [dungeon]; Tusken Helm (6686, -0.09 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.88 DPS) | yes | Ghostshard Talisman (7731, -0.26 DPS) [dungeon]; Ethereal Talisman (4430, -0.43 DPS) [quest]; Kaleidoscope Chain (13084, -0.52 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.96 DPS) | yes | Chromite Pauldrons (8144, -0.09 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.18 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 12.3 attack_power points (0.54 DPS) | yes | Wolfmaster Cape (6314, -0.10 DPS) [dungeon]; Wildhunter Cloak (16658, -0.10 DPS) [quest]; Dark Hooded Cape (5257, -0.17 DPS) [world] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.32 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.07 DPS) [quest]; Shining Mithril Breastplate (250540, -0.09 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.88 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.40 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.26 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.33 DPS) [dungeon] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 30.0 attack_power points (1.32 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.18 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.84 DPS) | yes | Firemane Leggings (13129, -0.18 DPS) [world_drop]; Orcish War Leggings (7929, -0.35 DPS) [crafted]; Symbolic Legplates (14829, -0.51 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 26.4 attack_power points (1.16 DPS) | yes | Prowler's Leather Shoes (252465, -0.18 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.19 DPS) [crafted]; Obsidian Greaves (13068, -0.27 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.88 DPS) | yes | Legionnaire's Band (19512, -0.16 DPS) [rep]; Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Suspicious Spare Part (274754, -0.26 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.88 DPS) | yes | Legionnaire's Band (19512, -0.16 DPS) [rep]; Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Suspicious Spare Part (274754, -0.26 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (97.3 DPS) | yes | Darkspear Raider's Reaper (272081, -1.34 DPS) [vendor]; Primitive Fishing Pole (276203, -1.72 DPS) [vendor]; The Jackhammer (9423, -2.14 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (97.3 DPS) | yes | Monolithic Bow (9426, -0.08 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.18 DPS) [crafted]; Bow of Searing Arrows (2825, -0.83 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: The Silencer

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 179.7. Weights run: 3.2s. Verify run: 1.7s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.631 ± 0.040, crit=0.902 ± 0.057 per rating point (14 rating = 1%, 12.626 per %), hit=0.224 ± 0.007 per rating point (10 rating = 1%, 2.241 per %), melee_haste=7.075 ± 0.975

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.34 DPS) | yes | Blood Guard's Plate Helm (220803, -0.35 DPS) [vendor]; Raging Berserker's Helm (7719, -0.46 DPS) [dungeon]; Fury Visor (20521, -0.53 DPS) [quest] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.97 DPS) | yes | Skibi's Pendant (13089, -0.09 DPS) [world_drop]; Woven Ivy Necklace (19159, -0.11 DPS) [quest]; Ghostshard Talisman (7731, -0.29 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 31.1 attack_power points (1.51 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.10 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.16 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.8 attack_power points (1.02 DPS) | yes | Bloodlust Cape (14801, -0.14 DPS) [world_drop]; First Sergeant's Cloak (16340, -0.25 DPS) [pvp]; Dark Hooded Cape (5257, -0.32 DPS) [world] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 48.0 attack_power points (2.34 DPS) | yes | Mixologist's Tunic (12793, -0.25 DPS) [dungeon]; Stone Guard's Plate Armor (220801, -0.26 DPS) [vendor]; Valorous Chestguard (8274, -0.39 DPS) [world_drop] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.36 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 50.3 attack_power points (2.45 DPS) | yes | Raider Gloves (272100, -0.80 DPS) [vendor]; Prowler's Leather Gauntlets (252547, -0.81 DPS) [crafted]; Officer's Gloves (250551, -1.06 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (2.24 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.49 DPS) [dungeon]; Belt of the Gladiator (13134, -0.49 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.51 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 44.0 attack_power points (2.14 DPS) | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS, sim-verified) [vendor]; Scarlet Leggings (10330, -0.10 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.19 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 36.2 attack_power points (1.76 DPS) | yes | Prowler's Leather Boots (252468, -0.16 DPS) [crafted]; Officer's Sabatons (250561, -0.22 DPS) [crafted]; Officer's Boots (250546, -0.28 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 25.7 attack_power points (1.25 DPS) | yes | Mark of Kern (2262, -0.28 DPS) [dungeon]; Assault Band (13095, -0.28 DPS) [world_drop]; White Bone Band (11862, -2.73 DPS, sim-verified) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (179.7 DPS) | yes | Mark of Kern (2262, -0.11 DPS) [dungeon]; Assault Band (13095, -0.11 DPS) [world_drop]; White Bone Band (11862, -3.11 DPS, sim-verified) [quest] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.03 DPS) [crafted]; Molten Heart of the Mountain (249470, -3.24 DPS, sim-verified) [crafted] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -3.38 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.10 DPS) [dungeon]; Dark Iron Rifle (16004, -1.75 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Rune of the Guard Captain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 265.2. Weights run: 3.3s. Verify run: 1.6s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.961 ± 0.058, crit=1.373 ± 0.083 per rating point (14 rating = 1%, 19.218 per %), hit=0.317 ± 0.010 per rating point (10 rating = 1%, 3.172 per %), melee_haste=12.382 ± 1.379

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 80.8 attack_power points (3.92 DPS) | yes | Warlord's Plate Headpiece (231535, -0.27 DPS) [pvp]; Champion's Plate Helm (227043, -0.80 DPS) [pvp]; Crown of Heroism (226860, -10.60 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (265.2 DPS) | yes | Amulet of the Darkmoon (19491, -0.24 DPS) [quest]; Imperial Jewel (11933, -0.54 DPS) [dungeon]; Rage of Mugamba (19577, -2.29 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 52.3 attack_power points (2.54 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Defiler's Leather Shoulders (20194, -0.24 DPS) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 38.8 attack_power points (1.88 DPS) | yes | Cape of the Black Baron (13340, -0.21 DPS) [dungeon]; Shroud of Domination (22337, -0.23 DPS) [dungeon]; Howler's Furs (272414, -0.37 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (265.2 DPS) | yes | Obsidian Mail Tunic (22191, -0.35 DPS) [crafted]; Cadaverous Armor (14637, -0.91 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -11.82 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (265.2 DPS) | yes | General's Plate Armguards (16546, -0.24 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.26 DPS) [rep]; Bracers of Undead Slaying (23090, -4.93 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (265.2 DPS) | yes | General's Plate Gauntlets (231532, -0.26 DPS) [pvp]; Radiant Gloves of the Dawn (227817, -0.44 DPS) [vendor]; Razor Gauntlets (18326, -5.19 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 73.2 attack_power points (3.56 DPS) | yes | Ferocity of the Timbermaw (227805, -0.24 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.45 DPS) [vendor]; General's Plate Girdle (16547, -0.57 DPS) [pvp] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (265.2 DPS) | yes | Titanic Leggings (22385, -0.49 DPS) [crafted]; General's Plate Leggings (231533, -0.68 DPS) [pvp]; Cloudkeeper Legplates (14554, -6.73 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 45.7 attack_power points (2.22 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Boots of Heroism (21995, -0.12 DPS) [quest]; Battleboots of Heroism (226857, -0.12 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (265.2 DPS) | yes | Band of the Ogre King (18522, -0.50 DPS) [dungeon]; Myrmidon's Signet (2246, -0.57 DPS) [world_drop]; Naglering (11669, -8.70 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (265.2 DPS) | yes | Band of the Ogre King (18522, -0.32 DPS) [dungeon]; Myrmidon's Signet (2246, -0.38 DPS) [world_drop]; Naglering (11669, -3.74 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (265.2 DPS) | yes | Counterattack Lodestone (18537, -1.08 DPS) [dungeon]; Hand of Justice (11815, -1.18 DPS) [dungeon]; Diamond Flask (20130, -3.69 DPS, sim-verified) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (265.2 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -1.70 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (265.2 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -11.70 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | The Legend of Korrak [quest] | sim-verified (265.2 DPS) | yes | Riphook (12653, -0.04 DPS) [dungeon]; Malgen's Long Bow (22318, -0.13 DPS) [dungeon]; Dark Iron Rifle (16004, -2.53 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Defiler's Plate Spaulders; back: Deathguard's Cloak; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Bloodseeker

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

