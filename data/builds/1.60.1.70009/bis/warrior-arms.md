# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 05321000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 30.0. Weights run: 2.5s. Verify run: 1.2s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.027, strength=1.946 ± 0.035, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.050 ± 0.001 per rating point (10 rating = 1%, 0.496 per %), melee_haste=1.661 ± 0.053

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 19.5 attack_power points (0.84 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 5.8 attack_power points (0.25 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 7.8 attack_power points (0.34 DPS) | yes | Catacomb Cloak (279899, -0.08 DPS) [quest]; Grave Shroud (279865, -0.08 DPS) [quest]; Miner's Cape (5444, -0.16 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 19.5 attack_power points (0.84 DPS) | yes | Veteran's Chain Shirt (250488, -0.21 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.25 DPS) [crafted]; Totemic Leather Armor (252435, -0.25 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 9.7 attack_power points (0.42 DPS) | yes | Cryptwalker Bracers (280095, -0.08 DPS) [quest]; Runed Copper Bracers (2854, -0.25 DPS) [crafted]; Bristlebark Bindings (14569, -0.25 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 15.6 attack_power points (0.67 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Polar Gauntlets (7606, -0.17 DPS) [quest]; Blackened Defias Gloves (10401, -0.17 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.78 DPS) | yes | Cobrahn's Grasp (6460, -0.19 DPS) [dungeon]; Ruffian Belt (5975, -0.27 DPS) [world]; Hulking Belt (14746, -0.36 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 21.4 attack_power points (0.93 DPS) | yes | Veteran's Chain Leggings (250493, -0.17 DPS) [crafted]; Defender's Leather Pants (252445, -0.17 DPS) [crafted]; Totemic Leather Pants (252446, -0.17 DPS) [crafted] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 9.7 attack_power points (0.42 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 7.8 attack_power points (0.34 DPS) | yes | The 1 Ring (8350, -0.25 DPS) [world]; Ring of the Moon (12052, -0.25 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 7.8 attack_power points (0.34 DPS) | yes | Ring of the Moon (12052, -0.21 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.25 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 321.2 attack_power points (13.91 DPS) | yes | Living Root (6631, -1.05 DPS) [dungeon]; Smite's Mighty Hammer (7230, -1.07 DPS, sim-verified) [dungeon]; Duskbringer (2205, -1.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.17 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Daryl's Hunting Rifle (2904, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Protector's Band; main_hand: The Axe of Severing; ranged: Fine Longbow

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 05325213000000000-000000000000000000-000000000000000000)

Set DPS (verified): 71.5. Weights run: 2.7s. Verify run: 1.4s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.106, strength=2.129 ± 0.135, agility=not significant (0.017 ± 0.006), crit=0.025 ± 0.002 per rating point (14 rating = 1%, 0.345 per %), hit=0.083 ± 0.003 per rating point (10 rating = 1%, 0.831 per %), melee_haste=2.836 ± 0.293

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 27.7 attack_power points (1.19 DPS) | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted]; Veteran's Chain Helm (250498, -0.26 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.60 DPS) | yes | Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.24 DPS) [world_drop]; Sentinel's Medallion (19541, -0.60 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.9 attack_power points (0.64 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.09 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.43 DPS) | yes | Sergeant Major's Cape (16315, -0.06 DPS) [pvp]; Lambent Scale Cloak (4706, -0.06 DPS) [world_drop]; Slayer's Cape (14752, -0.06 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 31.9 attack_power points (1.37 DPS) | yes | Shining Silver Breastplate (2870, -0.26 DPS, sim-verified) [crafted]; Barbaric Iron Breastplate (7914, -0.27 DPS) [crafted]; Hard Gold Cuirass (250533, -0.37 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 17.0 attack_power points (0.73 DPS) | yes | Yorgen Bracers (13012, -0.18 DPS) [world_drop]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.27 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.4 attack_power points (0.96 DPS) | yes | Bonefist Gauntlets (4465, -0.14 DPS) [world]; Mail Combat Gauntlets (4075, -0.23 DPS) [world_drop]; The Frozen Clutch (23170, -0.26 DPS, sim-verified) [dungeon] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 25.5 attack_power points (1.10 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS) [rep]; Highlander's Chain Girdle (20090, -0.07 DPS) [rep]; Highlander's Leather Girdle (20117, -0.07 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.12 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.11 DPS) [crafted]; Golden Scale Leggings (3843, -0.11 DPS) [crafted]; Chausses of Westfall (6087, -0.11 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 15.0 attack_power points (0.64 DPS) | yes | Hard Gold Boots (250534, -0.01 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop]; Slayer's Slippers (14756, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.73 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Silverlaine's Family Seal (6321, -0.28 DPS) [dungeon]; Insurgent's Band (272067, -0.35 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.9 attack_power points (0.55 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.17 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (71.5 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -20.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.39 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.11 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.20 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 05325213032310001-000000000000000000-000000000000000000)

Set DPS (verified): 105.0. Weights run: 3.3s. Verify run: 1.4s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.187, strength=2.059 ± 0.233, agility=not significant (0.020 ± 0.007), crit=0.050 ± 0.003 per rating point (14 rating = 1%, 0.701 per %), hit=0.095 ± 0.003 per rating point (10 rating = 1%, 0.950 per %), melee_haste=2.762 ± 0.350

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 28.8 attack_power points (1.60 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.08 DPS) [dungeon]; Tusken Helm (6686, -0.11 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.11 DPS) | yes | Kaleidoscope Chain (13084, -0.65 DPS) [world_drop]; River Pride Choker (13087, -0.65 DPS) [world_drop]; Ghostshard Talisman (7731, -0.75 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.7 attack_power points (1.26 DPS) | yes | Chromite Pauldrons (8144, -0.11 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.11 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.23 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 12.5 attack_power points (0.69 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Dark Hooded Cape (5257, -0.22 DPS) [world]; Slayer's Cape (14752, -0.24 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.9 attack_power points (1.72 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.11 DPS) [quest]; Shining Mithril Breastplate (250540, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.11 DPS) | yes | Pugilist Bracers (4438, -0.20 DPS) [dungeon]; Ravager's Armguards (14770, -0.31 DPS) [world_drop]; Yorgen Bracers (13012, -0.42 DPS) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.9 attack_power points (1.83 DPS) | yes | Gauntlets of Divinity (7724, -0.05 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.34 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.45 DPS) [dungeon] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 30.9 attack_power points (1.72 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.05 DPS) [rep]; Scarlet Belt (10329, -0.34 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 43.2 attack_power points (2.40 DPS) | yes | Firemane Leggings (13129, -0.23 DPS) [world_drop]; Orcish War Leggings (7929, -0.46 DPS) [crafted]; Symbolic Legplates (14829, -0.68 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 26.9 attack_power points (1.50 DPS) | yes | Prowler's Leather Shoes (252465, -0.23 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.34 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.11 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.11 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (105.0 DPS) | yes | Bonebiter (6830, -1.27 DPS) [quest]; Darkspear Raider's Reaper (272081, -1.63 DPS) [vendor]; Fiery War Axe (870, -13.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (105.0 DPS) | yes | Monolithic Bow (9426, -0.09 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.21 DPS) [crafted]; Bow of Searing Arrows (2825, -1.45 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: The Silencer

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 05325213032310001-050500000000000000-000000000000000000)

Set DPS (verified): 150.1. Weights run: 3.5s. Verify run: 1.8s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.322, strength=2.310 ± 0.409, agility=not significant (0.237 ± 0.138), crit=0.494 ± 0.033 per rating point (14 rating = 1%, 6.915 per %), hit=0.134 ± 0.004 per rating point (10 rating = 1%, 1.339 per %), melee_haste=3.330 ± 0.612

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | sim-verified (150.1 DPS) | yes | Sunscale Helmet (14849, -0.23 DPS) [world_drop]; Knight-Lieutenant's Plate Helm (220804, -0.24 DPS) [vendor]; Embrace of the Lycan (9479, -1.96 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.21 DPS) | yes | Skibi's Pendant (13089, -0.32 DPS) [world_drop]; Ghostshard Talisman (7731, -0.36 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.59 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 32.3 attack_power points (1.95 DPS) | yes | Earthslag Shoulders (11632, -0.14 DPS) [dungeon]; Wyrmslayer Spaulders (13066, -0.16 DPS) [world_drop]; Officer's Pauldrons (250576, -0.95 DPS, sim-verified) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 20.8 attack_power points (1.26 DPS) | yes | Blackveil Cape (11626, -0.22 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.33 DPS) [pvp]; Dark Hooded Cape (5257, -0.55 DPS) [world] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 55.4 attack_power points (3.35 DPS) | yes | Mixologist's Tunic (12793, -0.68 DPS) [dungeon]; Coldmetal Guard (274758, -0.70 DPS) [vendor]; Valorous Chestguard (8274, -0.95 DPS, sim-verified) [world_drop] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 32.3 attack_power points (1.95 DPS) | yes | Bracers of the Stone Princess (17714, -0.26 DPS) [dungeon]; Arena Bands (18711, -0.26 DPS) [world]; Officer's Wristguards (250581, -0.33 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 53.2 attack_power points (3.21 DPS) | yes | Truesilver Gauntlets (7938, -0.98 DPS) [crafted]; Officer's Gloves (250551, -1.04 DPS, sim-verified) [crafted]; Maddening Gauntlets (11867, -1.05 DPS) [quest] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 48.5 attack_power points (2.93 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.42 DPS) [dungeon]; Belt of the Gladiator (13134, -0.42 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.80 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 50.8 attack_power points (3.07 DPS) | yes | Scarlet Leggings (10330, -0.14 DPS) [dungeon]; Silvershell Leggings (10633, -0.28 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.28 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 35.4 attack_power points (2.14 DPS) | yes | Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.23 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 25.2 attack_power points (1.52 DPS) | yes | Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop]; Thunderbrow Ring (13097, -0.37 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.3 attack_power points (1.29 DPS) | yes | Assault Band (13095, -0.08 DPS) [world_drop]; Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Mark of Kern (2262, -2.20 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smoking Heart of the Mountain (11811, -2.04 DPS, sim-verified) [crafted] |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -8.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.12 DPS) [dungeon]; Dark Iron Rifle (16004, -2.05 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 05325213032310001-050500000000000000-500500000000000000)

Set DPS (verified): 223.8. Weights run: 9.7s. Verify run: 1.7s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.428, strength=1.460 ± 0.305, agility=not significant (0.334 ± 0.199), crit=0.721 ± 0.050 per rating point (14 rating = 1%, 10.092 per %), hit=0.213 ± 0.007 per rating point (10 rating = 1%, 2.131 per %), melee_haste=6.262 ± 0.875

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 50.7 attack_power points (2.74 DPS) | yes | Field Marshal's Plate Helm (231538, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -0.42 DPS) [vendor]; Embrace of the Lycan (9479, -10.09 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (223.8 DPS) | yes | Imperial Jewel (11933, -0.11 DPS) [dungeon]; Will of the Martyr (17044, -0.22 DPS) [quest]; Rage of Mugamba (19577, -3.01 DPS, sim-verified) [quest] |
| shoulder | Black Dragonscale Shoulders (15051) | Leatherworking [crafted] | 40.0 attack_power points (2.16 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS, sim-verified) [rep]; Lieutenant Commander's Plate Shoulders (227045, -0.27 DPS) [pvp]; Highlander's Lizardhide Shoulders (20060, -0.32 DPS) [rep] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 35.7 attack_power points (1.93 DPS) | yes | Howler's Furs (272414, -0.30 DPS) [vendor]; Cape of the Black Baron (13340, -0.58 DPS) [dungeon]; Shroud of Domination (22337, -0.59 DPS) [dungeon] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (223.8 DPS) | yes | Cadaverous Armor (14637, -0.63 DPS) [dungeon]; Timbermaw Tunic (252484, -1.18 DPS) [crafted]; Breastplate of Undead Slaying (23087, -10.54 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | sim-verified (223.8 DPS) | yes | Berserker Bracers (19578, -0.41 DPS) [rep]; Arena Bands (18711, -0.54 DPS) [world]; Bracers of Undead Slaying (23090, -4.96 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (223.8 DPS) | yes | Raider Gauntlets (272095, -0.05 DPS) [vendor]; Skul's Fingerbone Claws (13395, -0.22 DPS) [dungeon]; Razor Gauntlets (18326, -3.09 DPS, sim-verified) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (3.46 DPS) | yes | Radiant Girdle of the Dawn (227814, +0.00 DPS, sim-verified) [vendor]; Highlander's Chain Girdle (20043, -1.07 DPS) [rep]; Highlander's Leather Girdle (20045, -1.07 DPS) [rep] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (223.8 DPS) | yes | Titanic Leggings (22385, -0.19 DPS) [crafted]; Devilsaur Leggings (15062, -0.30 DPS) [crafted]; Cloudkeeper Legplates (14554, -6.46 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 40.7 attack_power points (2.20 DPS) | yes | Pads of the Dread Wolf (13210, -0.04 DPS) [dungeon]; Goregasher Stompers (275624, -0.14 DPS) [crafted]; Marshal's Plate Boots (231539, -0.45 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (223.8 DPS) | yes | Protector's Band (19514, -0.38 DPS) [rep]; Band of the Ogre King (18522, -0.42 DPS) [dungeon]; Naglering (11669, -7.51 DPS, sim-verified) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (223.8 DPS) | yes | Protector's Band (19514, -0.05 DPS) [rep]; Band of the Ogre King (18522, -0.09 DPS) [dungeon]; Naglering (11669, -5.08 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (223.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Diamond Flask (20130, -1.26 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (223.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.98 DPS) [crafted]; Smolderweb's Eye (13213, -5.93 DPS, sim-verified) [dungeon] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (223.8 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -5.57 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (223.8 DPS) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.20 DPS) [world_drop]; Dark Iron Rifle (16004, -3.18 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Black Dragonscale Shoulders; back: Cloak of the Honor Guard; chest: Obsidian Mail Tunic; wrist: Windtalker's Wristguards; hands: Cadaverous Gloves; waist: Dense Timbermaw Belt; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Blackblade of Shahram; ranged: Riphook

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 05321000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 33.2. Weights run: 2.5s. Verify run: 1.2s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.027, strength=1.946 ± 0.035, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.050 ± 0.001 per rating point (10 rating = 1%, 0.496 per %), melee_haste=1.661 ± 0.053

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 19.5 attack_power points (0.84 DPS) | yes | Defender's Leather Hood (252447, -0.22 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 5.8 attack_power points (0.25 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 7.8 attack_power points (0.34 DPS) | yes | Subterranean Cape (14149, -0.08 DPS) [dungeon]; Grave Shroud (279865, -0.08 DPS) [quest]; Catacomb Cloak (279899, -0.09 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 19.5 attack_power points (0.84 DPS) | yes | Defender's Leather Armor (252434, -0.25 DPS) [crafted]; Totemic Leather Armor (252435, -0.25 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.28 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 9.7 attack_power points (0.42 DPS) | yes | Raptorcrest Bracers (270010, -0.22 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.25 DPS) [crafted]; Bristlebark Bindings (14569, -0.25 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 15.6 attack_power points (0.67 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.17 DPS) [dungeon]; Foreman's Gloves (2167, -0.25 DPS) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.78 DPS) | yes | Cobrahn's Grasp (6460, -0.16 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.27 DPS) [world]; Hulking Belt (14746, -0.36 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 17.5 attack_power points (0.76 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 9.7 attack_power points (0.42 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 7.8 attack_power points (0.34 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; The 1 Ring (8350, -0.25 DPS) [world]; Ring of the Moon (12052, -0.25 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 7.8 attack_power points (0.34 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; The 1 Ring (8350, -0.25 DPS) [world]; Ring of the Moon (12052, -0.25 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 321.2 attack_power points (13.91 DPS) | yes | Forsaken Greataxe (251533, -0.81 DPS) [quest]; Smite's Mighty Hammer (7230, -1.03 DPS) [dungeon]; Hammerbone (270018, -3.33 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.17 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Heavy Shortbow (3036, -0.09 DPS) [world_drop]; Orcish Battle Bow (5346, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Legionnaire's Band; main_hand: The Axe of Severing; ranged: Fine Longbow

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 05325213000000000-000000000000000000-000000000000000000)

Set DPS (verified): 72.8. Weights run: 2.7s. Verify run: 1.4s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.106, strength=2.129 ± 0.135, agility=not significant (0.017 ± 0.006), crit=0.025 ± 0.002 per rating point (14 rating = 1%, 0.345 per %), hit=0.083 ± 0.003 per rating point (10 rating = 1%, 0.831 per %), melee_haste=2.836 ± 0.293

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 27.7 attack_power points (1.19 DPS) | yes | Veteran's Chain Helm (250498, -0.09 DPS) [crafted]; Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.60 DPS) | yes | River Pride Choker (13087, -0.24 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.45 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.60 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.9 attack_power points (0.64 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.43 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.06 DPS) [world_drop]; Slayer's Cape (14752, -0.06 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 31.9 attack_power points (1.37 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.27 DPS) [crafted]; Hard Gold Cuirass (250533, -0.37 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 17.0 attack_power points (0.73 DPS) | yes | Yorgen Bracers (13012, -0.18 DPS) [world_drop]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.27 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.4 attack_power points (0.96 DPS) | yes | Warsong Gauntlets (16978, -0.05 DPS) [quest]; The Frozen Clutch (23170, -0.10 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.14 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 25.5 attack_power points (1.10 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS) [rep]; Defiler's Chain Girdle (20152, -0.07 DPS) [rep]; Defiler's Leather Girdle (20191, -0.07 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.12 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.11 DPS) [crafted]; Golden Scale Leggings (3843, -0.11 DPS) [crafted]; Slayer's Pants (14757, -0.11 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 15.0 attack_power points (0.64 DPS) | yes | Hard Gold Boots (250534, -0.01 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop]; Slayer's Slippers (14756, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.73 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Silverlaine's Family Seal (6321, -0.28 DPS) [dungeon]; Insurgent's Band (272067, -0.35 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.9 attack_power points (0.55 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.17 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.8 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -21.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.39 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.11 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.20 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 05325213032310001-000000000000000000-000000000000000000)

Set DPS (verified): 102.6. Weights run: 3.3s. Verify run: 1.4s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.187, strength=2.059 ± 0.233, agility=not significant (0.020 ± 0.007), crit=0.050 ± 0.003 per rating point (14 rating = 1%, 0.701 per %), hit=0.095 ± 0.003 per rating point (10 rating = 1%, 0.950 per %), melee_haste=2.762 ± 0.350

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 28.8 attack_power points (1.60 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.08 DPS) [dungeon]; Tusken Helm (6686, -0.11 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.11 DPS) | yes | Ghostshard Talisman (7731, -0.33 DPS) [dungeon]; Ethereal Talisman (4430, -0.54 DPS) [quest]; Kaleidoscope Chain (13084, -0.65 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.7 attack_power points (1.26 DPS) | yes | Chromite Pauldrons (8144, -0.11 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.11 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.23 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 12.5 attack_power points (0.69 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Wildhunter Cloak (16658, -0.14 DPS) [quest]; Dark Hooded Cape (5257, -0.22 DPS) [world] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.9 attack_power points (1.72 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.11 DPS) [quest]; Shining Mithril Breastplate (250540, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.11 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.9 attack_power points (1.83 DPS) | yes | Gauntlets of Divinity (7724, -0.05 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.34 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.45 DPS) [dungeon] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 30.9 attack_power points (1.72 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.05 DPS) [rep]; Tharg's Shoelace (9705, -0.23 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 43.2 attack_power points (2.40 DPS) | yes | Firemane Leggings (13129, -0.23 DPS) [world_drop]; Orcish War Leggings (7929, -0.46 DPS) [crafted]; Symbolic Legplates (14829, -0.68 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 26.9 attack_power points (1.50 DPS) | yes | Prowler's Leather Shoes (252465, -0.23 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.34 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.11 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.11 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (102.6 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -1.91 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (102.6 DPS) | yes | Monolithic Bow (9426, -0.09 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.21 DPS) [crafted]; Bow of Searing Arrows (2825, -1.54 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Fiery War Axe; ranged: The Silencer

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 05325213032310001-050500000000000000-000000000000000000)

Set DPS (verified): 152.0. Weights run: 3.5s. Verify run: 1.8s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.322, strength=2.310 ± 0.409, agility=not significant (0.237 ± 0.138), crit=0.494 ± 0.033 per rating point (14 rating = 1%, 6.915 per %), hit=0.134 ± 0.004 per rating point (10 rating = 1%, 1.339 per %), melee_haste=3.330 ± 0.612

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Sunscale Helmet (14849, -0.23 DPS) [world_drop]; Blood Guard's Plate Helm (220803, -0.24 DPS) [vendor]; Embrace of the Lycan (9479, -3.49 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.21 DPS) | yes | Woven Ivy Necklace (19159, -0.24 DPS) [quest]; Skibi's Pendant (13089, -0.32 DPS) [world_drop]; Ghostshard Talisman (7731, -0.36 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 31.9 attack_power points (1.93 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Earthslag Shoulders (11632, -0.11 DPS) [dungeon]; Wyrmslayer Spaulders (13066, -0.14 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 20.8 attack_power points (1.26 DPS) | yes | Blackveil Cape (11626, +0.00 DPS, sim-verified) [dungeon]; First Sergeant's Cloak (16340, -0.33 DPS) [pvp]; Dark Hooded Cape (5257, -0.55 DPS) [world] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 55.4 attack_power points (3.35 DPS) | yes | Valorous Chestguard (8274, -0.56 DPS) [world_drop]; Mixologist's Tunic (12793, -0.68 DPS) [dungeon]; Coldmetal Guard (274758, -0.70 DPS) [vendor] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 32.3 attack_power points (1.95 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 53.2 attack_power points (3.21 DPS) | yes | Officer's Gloves (250551, -0.85 DPS) [crafted]; Truesilver Gauntlets (7938, -0.98 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -1.13 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 48.5 attack_power points (2.93 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.42 DPS) [dungeon]; Belt of the Gladiator (13134, -0.42 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.80 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 50.8 attack_power points (3.07 DPS) | yes | Scarlet Leggings (10330, -0.14 DPS) [dungeon]; Silvershell Leggings (10633, -0.28 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.28 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 35.4 attack_power points (2.14 DPS) | yes | Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.23 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 25.2 attack_power points (1.52 DPS) | yes | Band of Allegiance (18585, -0.27 DPS) [quest]; Mark of Kern (2262, -0.32 DPS) [dungeon]; White Bone Band (11862, -2.69 DPS, sim-verified) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Band of Allegiance (18585, -0.03 DPS) [quest]; Mark of Kern (2262, -0.08 DPS) [dungeon]; White Bone Band (11862, -2.61 DPS, sim-verified) [quest] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.52 DPS) [crafted]; Molten Heart of the Mountain (249470, -4.02 DPS, sim-verified) [crafted] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -3.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.12 DPS) [dungeon]; Dark Iron Rifle (16004, -1.84 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Rune of the Guard Captain; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 05325213032310001-050500000000000000-500500000000000000)

Set DPS (verified): 210.8. Weights run: 9.7s. Verify run: 1.6s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.428, strength=1.460 ± 0.305, agility=not significant (0.334 ± 0.199), crit=0.721 ± 0.050 per rating point (14 rating = 1%, 10.092 per %), hit=0.213 ± 0.007 per rating point (10 rating = 1%, 2.131 per %), melee_haste=6.262 ± 0.875

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 50.7 attack_power points (2.74 DPS) | yes | Warlord's Plate Headpiece (231535, +0.00 DPS) [pvp]; Champion's Plate Helm (227043, -0.42 DPS) [pvp]; Embrace of the Lycan (9479, -10.56 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (210.8 DPS) | yes | Imperial Jewel (11933, -0.11 DPS) [dungeon]; Will of the Martyr (17044, -0.22 DPS) [quest]; Rage of Mugamba (19577, -1.85 DPS, sim-verified) [quest] |
| shoulder | Black Dragonscale Shoulders (15051) | Leatherworking [crafted] | 40.0 attack_power points (2.16 DPS) | yes | Defiler's Leather Shoulders (20194, -0.22 DPS) [rep]; Champion's Plate Shoulders (227042, -0.27 DPS) [pvp]; Defiler's Lizardhide Shoulders (20175, -0.32 DPS) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 35.7 attack_power points (1.93 DPS) | yes | Howler's Furs (272414, -0.30 DPS) [vendor]; Cape of the Black Baron (13340, -0.58 DPS) [dungeon]; Shroud of Domination (22337, -0.59 DPS) [dungeon] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (210.8 DPS) | yes | Cadaverous Armor (14637, -0.63 DPS) [dungeon]; Timbermaw Tunic (252484, -1.18 DPS) [crafted]; Breastplate of Undead Slaying (23087, -12.06 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | sim-verified (210.8 DPS) | yes | Berserker Bracers (19578, -0.41 DPS) [rep]; Arena Bands (18711, -0.54 DPS) [world]; Bracers of Undead Slaying (23090, -4.05 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (210.8 DPS) | yes | Raider Gauntlets (272095, -0.05 DPS) [vendor]; Skul's Fingerbone Claws (13395, -0.22 DPS) [dungeon]; Razor Gauntlets (18326, -3.85 DPS, sim-verified) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (3.46 DPS) | yes | Radiant Girdle of the Dawn (227814, +0.00 DPS, sim-verified) [vendor]; Defiler's Chain Girdle (20150, -1.07 DPS) [rep]; Defiler's Leather Girdle (20190, -1.07 DPS) [rep] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (210.8 DPS) | yes | Titanic Leggings (22385, -0.19 DPS) [crafted]; Devilsaur Leggings (15062, -0.30 DPS) [crafted]; Cloudkeeper Legplates (14554, -5.92 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 40.7 attack_power points (2.20 DPS) | yes | Pads of the Dread Wolf (13210, -0.04 DPS) [dungeon]; Goregasher Stompers (275624, -0.14 DPS) [crafted]; General's Plate Boots (231531, -0.45 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (210.8 DPS) | yes | Blackstone Ring (17713, -0.33 DPS) [dungeon]; Legionnaire's Band (19510, -0.38 DPS) [rep]; Naglering (11669, -5.81 DPS, sim-verified) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | sim-verified (210.8 DPS) | yes | Blackstone Ring (17713, -0.10 DPS) [dungeon]; Legionnaire's Band (19510, -0.15 DPS) [rep]; Naglering (11669, -3.80 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (210.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Diamond Flask (20130, -2.53 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (210.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.98 DPS) [crafted]; Rune of the Guard Captain (19120, -1.32 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (210.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -7.62 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (210.8 DPS) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.20 DPS) [world_drop]; Dark Iron Rifle (16004, -3.45 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Black Dragonscale Shoulders; back: Deathguard's Cloak; chest: Obsidian Mail Tunic; wrist: Windtalker's Wristguards; hands: Cadaverous Gloves; waist: Dense Timbermaw Belt; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: White Bone Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force; ranged: Riphook

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

