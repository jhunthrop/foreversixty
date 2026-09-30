# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 29.3. Weights run: 1.2s. Verify run: 1.1s. 303 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.053 ± 0.002 per rating point (10 rating = 1%, 0.529 per %), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.18 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.3 attack_power points (0.26 DPS) | yes | Rough Bronze Shoulders (3480, +0.00 DPS, sim-verified) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Catacomb Cloak (279899, -0.10 DPS) [quest]; Grave Shroud (279865, -0.10 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.18 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.27 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Cryptwalker Bracers (280095, -0.10 DPS, sim-verified) [quest]; Bristlebark Bindings (14569, -0.26 DPS) [world_drop]; Bravo's Armbands (270015, -0.26 DPS) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 attack_power points (0.70 DPS) | yes | Gold-flecked Gloves (5195, -0.10 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -0.18 DPS) [quest]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.18 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Warchief's Girdle (5750, -0.32 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.9 attack_power points (0.96 DPS) | yes | Defender's Leather Pants (252445, -0.18 DPS) [crafted]; Totemic Leather Pants (252446, -0.18 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.18 DPS, sim-verified) [crafted] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 attack_power points (0.35 DPS) | yes | The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.27 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | Westfall: Mr. Smite [dungeon] | sim-verified (29.3 DPS) | yes | Living Root (6631, -0.08 DPS) [dungeon]; Duskbringer (2205, -0.19 DPS) [dungeon]; The Axe of Severing (23171, -13.03 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567)) | Blacksmithing [crafted] | 4.2 attack_power points (0.18 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS, sim-verified) [quest]; Fine Longbow (11304, -0.01 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Veteran's Boots; finger1: Demon Band; finger2: Protector's Band; main_hand: Smite's Mighty Hammer; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 303, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings

### Band 30 (human, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 70.5. Weights run: 1.3s. Verify run: 1.1s. 489 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.012 ± 0.002 per rating point (14 rating = 1%, 0.171 per %), hit=0.074 ± 0.003 per rating point (10 rating = 1%, 0.743 per %), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.26 DPS) | yes | Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Veteran's Chain Helm (250498, -0.17 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.68 DPS) | yes | River Pride Choker (13087, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.46 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.68 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.49 DPS) | yes | Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.16 DPS, sim-verified) [pvp] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (70.5 DPS) | yes | Barbaric Iron Breastplate (7914, -0.19 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Avenger's Armor (1488, -1.83 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 attack_power points (0.78 DPS) | yes | Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Cultist's Armguards (270032, -0.29 DPS) [quest]; Yorgen Bracers (13012, -0.44 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.07 DPS) | yes | Bonefist Gauntlets (4465, -0.20 DPS) [world]; Heavy Earthen Gloves (7359, -0.29 DPS) [crafted]; The Frozen Clutch (23170, -1.59 DPS, sim-verified) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.17 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.27 DPS) | yes | Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Slayer's Pants (14757, -0.20 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -0.43 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 attack_power points (0.68 DPS) | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 attack_power points (0.78 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.0 attack_power points (0.59 DPS) | yes | Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Tiger Band (6749, -0.93 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -20.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.44 DPS) | yes | Long Battle Bow (15284, -0.15 DPS) [world_drop]; Fine Longbow (11304, -0.24 DPS) [vendor]; Double-barreled Shotgun (2098, -0.25 DPS, sim-verified) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 489, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 77.9. Weights run: 1.6s. Verify run: 1.3s. 673 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.304, strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.034 ± 0.004 per rating point (14 rating = 1%, 0.473 per %), hit=0.114 ± 0.006 per rating point (10 rating = 1%, 1.140 per %), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Icemetal Barbute (10763) (or Hard Gold Coif (250537)) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 34.6 attack_power points (1.43 DPS) | yes | Hard Gold Coif (250537, +0.00 DPS, sim-verified) [crafted]; Raging Berserker's Helm (7719, -0.08 DPS) [dungeon]; Tusken Helm (6686, -0.10 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; Gazlowe's Charm (13088, -0.17 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.39 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 attack_power points (1.12 DPS) | yes | Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Chromite Pauldrons (8144, -0.12 DPS, sim-verified) [world_drop]; Imperial Leather Spaulders (4737, -0.20 DPS) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.9 attack_power points (0.61 DPS) | yes | Wolfmaster Cape (6314, -0.16 DPS, sim-verified) [dungeon]; Dark Hooded Cape (5257, -0.20 DPS) [world]; Sergeant Major's Cape (16315, -0.20 DPS) [pvp] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | World drop [world_drop] | 37.1 attack_power points (1.53 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Shining Mithril Breastplate (250540, -0.10 DPS) [crafted]; Avenger's Armor (1488, -1.87 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Yorgen Bracers (13012, -0.21 DPS) [world_drop]; Pugilist Bracers (4438, -0.25 DPS, sim-verified) [dungeon] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 attack_power points (1.63 DPS) | yes | Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.42 DPS, sim-verified) [world_drop] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 37.1 attack_power points (1.53 DPS) | yes | Highlander's Leather Girdle (20116, -0.29 DPS) [rep]; Highlander's Lamellar Girdle (20107, -0.31 DPS) [rep]; Boar Champion's Belt (10768, -1.87 DPS, sim-verified) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 attack_power points (2.15 DPS) | yes | Firemane Leggings (13129, -0.25 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.41 DPS) [crafted]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 attack_power points (1.33 DPS) | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.25 DPS, sim-verified) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.83 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor]; Mark of Kern (2262, -1.70 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrow Ring (13097, -0.00 DPS) [world_drop]; Suspicious Spare Part (274754, -0.10 DPS) [vendor]; Mark of Kern (2262, -1.62 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Nightblade (1982, -41.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Bow of Searing Arrows (2825, +0.00 DPS, sim-verified) [world_drop]; The Silencer (13138, -0.04 DPS) [world_drop]; Mithril Blacksmith Hammer (285280, -0.10 DPS) [crafted] |

**New at 40:** head: Icemetal Barbute; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Protector's Band; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 673, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 110.3. Weights run: 1.7s. Verify run: 1.4s. 865 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.275, strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=0.440 ± 0.037 per rating point (14 rating = 1%, 6.161 per %), hit=0.098 ± 0.005 per rating point (10 rating = 1%, 0.978 per %), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.6 attack_power points (3.46 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Bloomsprout Headpiece (17767, -0.67 DPS) [dungeon]; Fury Visor (20521, -1.22 DPS) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Skibi's Pendant (13089, -0.19 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.77 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 23.5 attack_power points (1.82 DPS) | yes | Wyrmslayer Spaulders (13066, -0.18 DPS) [world_drop]; Earthslag Shoulders (11632, -0.23 DPS) [dungeon]; Officer's Pauldrons (250576, -0.49 DPS, sim-verified) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 attack_power points (1.10 DPS) | yes | Sergeant Major's Cape (16336, -0.23 DPS) [pvp]; Wolfmaster Cape (6314, -0.32 DPS) [dungeon]; Blackveil Cape (11626, -0.85 DPS, sim-verified) [dungeon] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 37.8 attack_power points (2.93 DPS) | yes | Valorous Chestguard (8274, -0.47 DPS, sim-verified) [world_drop]; Mixologist's Tunic (12793, -0.49 DPS) [dungeon]; Grizzled Pelt (22274, -0.50 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.17 DPS) | yes | Branded Leather Bracers (19508, -0.62 DPS) [dungeon]; Officer's Wristguards (250581, -0.70 DPS) [crafted]; Runed Golem Shackles (12550, -1.61 DPS, sim-verified) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 37.5 attack_power points (2.90 DPS) | yes | Rockgrip Gauntlets (17736, -0.73 DPS) [dungeon]; Officer's Gloves (250551, -0.75 DPS) [crafted]; Gauntlets of Divinity (7724, -1.23 DPS, sim-verified) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Atal'alarion's Tusk Ring (10798, -0.13 DPS) [dungeon]; Belt of the Gladiator (13134, -0.13 DPS) [world_drop]; Girdle of Beastial Fury (11686, -3.50 DPS, sim-verified) [dungeon] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.6 attack_power points (2.68 DPS) | yes | Scarlet Leggings (10330, -0.23 DPS, sim-verified) [dungeon]; Silvershell Leggings (10633, -0.24 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.24 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Officer's Sabatons (250561, -0.04 DPS) [crafted]; Officer's Boots (250546, -0.09 DPS) [crafted]; Battlechaser's Greaves (12555, -3.90 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.63 DPS) | yes | Mark of Kern (2262, -0.08 DPS) [dungeon]; Protector's Band (19516, -0.21 DPS) [rep]; Insurgent's Band (272065, -0.46 DPS) [vendor] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (1.55 DPS) | yes | Protector's Band (19516, -0.13 DPS) [rep]; Insurgent's Band (272065, -0.39 DPS) [vendor]; Mark of Kern (2262, -1.18 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Blight (7959) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Fiery War Axe (870, -5.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.18 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Bracers of the Stone Princess; hands: Raider Gauntlets; waist: Highlander's Leather Girdle; legs: Golem Shard Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Assault Band; trinket1: Frozen Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 865, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 224.6. Weights run: 1.6s. Verify run: 1.5s. 1885 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.801, strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=1.448 ± 0.116 per rating point (14 rating = 1%, 20.276 per %), hit=0.304 ± 0.015 per rating point (10 rating = 1%, 3.037 per %), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lightbreaker Greathelm (239517) | Leonid Barthalomew the Revered [vendor] | 235.9 attack_power points (7.55 DPS) | yes | Field Marshal's Plate Helm (16478, -3.42 DPS) [vendor]; Field Marshal's Plate Helm (231538, -3.42 DPS) [pvp]; Lightbreaker Helmet (239525, -6.03 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 0.0 attack_power points (0.00 DPS) | yes | Medallion of the Dawn (22659, -0.41 DPS) [quest]; Strength of Mugamba (19576, -0.59 DPS) [quest]; Rage of Mugamba (19577, -1.09 DPS, sim-verified) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 169.8 attack_power points (5.44 DPS) | yes | Field Marshal's Plate Shoulderguards (16480, -2.61 DPS) [vendor]; Field Marshal's Plate Shoulderguards (231537, -2.61 DPS) [pvp]; Lightbreaker Pauldrons (239524, -5.73 DPS, sim-verified) [vendor] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 66.1 attack_power points (2.12 DPS) | yes | Shadewood Cloak (18328, +0.00 DPS, sim-verified) [dungeon]; Cloak of Revanchion (23127, -0.56 DPS) [dungeon]; Phantasmal Cloak (18689, -0.62 DPS) [dungeon] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Timbermaw Tunic (252484, -2.46 DPS) [crafted]; Lightbreaker Breastplate (239527, -4.28 DPS) [vendor]; Breastplate of Undead Slaying (23087, -13.60 DPS, sim-verified) [world] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Berserker Bracers (19578, -1.40 DPS) [rep]; Marshal's Plate Bracers (16481, -1.68 DPS) [pvp]; Bracers of Undead Slaying (23090, -7.31 DPS, sim-verified) [world] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Raider Gauntlets (272095, -1.55 DPS) [vendor]; Radiant Gloves of the Dawn (227817, -2.08 DPS) [vendor]; Razor Gauntlets (18326, -8.31 DPS, sim-verified) [dungeon] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 169.8 attack_power points (5.44 DPS) | yes | Radiant Girdle of the Dawn (227814, -1.00 DPS, sim-verified) [vendor]; Ferocity of the Timbermaw (227805, -1.68 DPS) [vendor]; Marshal's Plate Girdle (16482, -1.96 DPS) [pvp] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Sentinel's Plate Legguards (237825, -3.07 DPS) [vendor]; Titanic Leggings (22385, -3.25 DPS) [crafted]; Cloudkeeper Legplates (14554, -13.44 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 168.9 attack_power points (5.41 DPS) | yes | Marshal's Plate Boots (16483, -2.70 DPS) [vendor]; Marshal's Plate Boots (231539, -2.70 DPS) [pvp]; Boots of Heroism (21995, -4.86 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.25 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234026, -0.37 DPS) [vendor]; Naglering (11669, -5.55 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | 0.0 attack_power points (0.00 DPS) | yes | Band of the Ogre King (18522, -0.09 DPS) [dungeon]; Protector's Band (19516, -0.31 DPS) [rep]; Naglering (11669, -1.44 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (224.6 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Heroism (19287, -2.94 DPS, sim-verified) [quest] |
| main_hand | Blackfury (19167) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Nightfall (19169, +1.68 DPS, sim-verified) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | 0.0 attack_power points (0.00 DPS) | yes | Unsophisticated Hand Cannon (18460, -0.22 DPS) [dungeon]; Monolithic Bow (9426, -0.37 DPS) [dungeon]; Dark Iron Rifle (16004, -1.14 DPS, sim-verified) [crafted] |

**New at 60:** head: Lightbreaker Greathelm; neck: Amulet of the Darkmoon; shoulder: Lightbreaker Shoulders; back: Shroud of Domination; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: Blackfury; ranged: Bloodseeker

No-known-source sample (15 of 1885, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 30.2. Weights run: 1.2s. Verify run: 1.1s. 286 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.053 ± 0.002 per rating point (10 rating = 1%, 0.529 per %), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.22 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.3 attack_power points (0.26 DPS) | yes | Rough Bronze Shoulders (3480, +0.00 DPS, sim-verified) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Subterranean Cape (14149, -0.09 DPS) [dungeon]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Grave Shroud (279865, -0.10 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.29 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Raptorcrest Bracers (270010, -0.22 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted]; Bristlebark Bindings (14569, -0.26 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 attack_power points (0.70 DPS) | yes | Gold-flecked Gloves (5195, -0.10 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.26 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.17 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Warchief's Girdle (5750, -0.32 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.8 attack_power points (0.79 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.09 DPS) [world_drop] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | sim-verified (30.2 DPS) | yes | Forsaken Greataxe (251533, -0.14 DPS) [quest]; Smite's Mighty Hammer (7230, -0.34 DPS) [dungeon]; The Axe of Severing (23171, -13.32 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.2 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Heavy Shortbow (3036, -0.09 DPS) [world_drop]; Orcish Battle Bow (5346, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Veteran's Boots; finger1: Demon Band; finger2: Legionnaire's Band; main_hand: Hammerbone; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 72.0. Weights run: 1.3s. Verify run: 1.1s. 472 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.012 ± 0.002 per rating point (14 rating = 1%, 0.171 per %), hit=0.074 ± 0.003 per rating point (10 rating = 1%, 0.743 per %), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.26 DPS) | yes | Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted]; Veteran's Chain Helm (250498, -0.22 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.68 DPS) | yes | Kaleidoscope Chain (13084, -0.28 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.29 DPS) [world_drop]; Scout's Medallion (19537, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.68 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.49 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (72.0 DPS) | yes | Barbaric Iron Breastplate (7914, -0.19 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Avenger's Armor (1488, -2.19 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 attack_power points (0.78 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Cultist's Armguards (270032, -0.29 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.07 DPS) | yes | Warsong Gauntlets (16978, -0.10 DPS) [quest]; Bonefist Gauntlets (4465, -0.20 DPS) [world]; The Frozen Clutch (23170, -1.64 DPS, sim-verified) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.17 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.27 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.13 DPS, sim-verified) [crafted]; Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Slayer's Pants (14757, -0.20 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 attack_power points (0.68 DPS) | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop]; Slayer's Slippers (14756, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 attack_power points (0.78 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.0 attack_power points (0.59 DPS) | yes | Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Tiger Band (6749, -0.92 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -20.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.44 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS, sim-verified) [world_drop]; Long Battle Bow (15284, -0.15 DPS) [world_drop]; Fine Longbow (11304, -0.24 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 472, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 86.3. Weights run: 1.6s. Verify run: 1.2s. 654 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.304, strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.034 ± 0.004 per rating point (14 rating = 1%, 0.473 per %), hit=0.114 ± 0.006 per rating point (10 rating = 1%, 1.140 per %), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Icemetal Barbute (10763) (or Hard Gold Coif (250537)) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 34.6 attack_power points (1.43 DPS) | yes | Hard Gold Coif (250537, +0.00 DPS, sim-verified) [crafted]; Raging Berserker's Helm (7719, -0.08 DPS) [dungeon]; Tusken Helm (6686, -0.10 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.83 DPS) | yes | Ghostshard Talisman (7731, +0.74 DPS, sim-verified) [dungeon]; Ethereal Talisman (4430, -0.31 DPS) [quest]; Kaleidoscope Chain (13084, -0.42 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 attack_power points (1.12 DPS) | yes | Chromite Pauldrons (8144, +0.00 DPS, sim-verified) [world_drop]; Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [world_drop] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.41 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Dark Hooded Cape (5257, -0.00 DPS) [world]; Khan's Cloak (14781, -0.00 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | World drop [world_drop] | 37.1 attack_power points (1.53 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Shining Mithril Breastplate (250540, -0.10 DPS) [crafted]; Avenger's Armor (1488, -2.06 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Darkspear Armsplints (4132, -0.11 DPS) [quest]; Pugilist Bracers (4438, -0.18 DPS, sim-verified) [dungeon] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 attack_power points (1.63 DPS) | yes | Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.44 DPS, sim-verified) [world_drop] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 37.1 attack_power points (1.53 DPS) | yes | Tharg's Shoelace (9705, -0.20 DPS) [quest]; Defiler's Leather Girdle (20192, -0.29 DPS) [rep]; Boar Champion's Belt (10768, -2.06 DPS, sim-verified) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 attack_power points (2.15 DPS) | yes | Firemane Leggings (13129, -0.14 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.41 DPS) [crafted]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 attack_power points (1.33 DPS) | yes | Prowler's Leather Shoes (252465, -0.14 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.83 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor]; Mark of Kern (2262, -0.99 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | sim-verified (86.3 DPS) | yes | Thunderbrow Ring (13097, -0.00 DPS) [world_drop]; Suspicious Spare Part (274754, -0.10 DPS) [vendor]; Mark of Kern (2262, -0.95 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Nightblade (1982, -50.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Bow of Searing Arrows (2825, +0.00 DPS, sim-verified) [world_drop]; The Silencer (13138, -0.04 DPS) [world_drop]; Mithril Blacksmith Hammer (285280, -0.10 DPS) [crafted] |

**New at 40:** head: Icemetal Barbute; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Legionnaire's Band; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 654, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 122.3. Weights run: 1.7s. Verify run: 1.5s. 846 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.275, strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=0.440 ± 0.037 per rating point (14 rating = 1%, 6.161 per %), hit=0.098 ± 0.005 per rating point (10 rating = 1%, 0.978 per %), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.6 attack_power points (3.46 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Bloomsprout Headpiece (17767, -0.67 DPS) [dungeon]; Fury Visor (20521, -1.22 DPS) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Woven Ivy Necklace (19159, -0.15 DPS) [quest]; Skibi's Pendant (13089, -0.19 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.60 DPS, sim-verified) [quest] |
| shoulder | Blood Guard's Plate Pauldrons (220796) | Lady Palanseer [vendor] | 23.5 attack_power points (1.82 DPS) | yes | Officer's Pauldrons (250576, -0.06 DPS, sim-verified) [crafted]; Wyrmslayer Spaulders (13066, -0.18 DPS) [world_drop]; Earthslag Shoulders (11632, -0.23 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 attack_power points (1.10 DPS) | yes | Battlehard Cape (11858, -0.32 DPS) [quest]; Wildhunter Cloak (16658, -0.32 DPS) [quest]; Blackveil Cape (11626, -1.44 DPS, sim-verified) [dungeon] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 37.8 attack_power points (2.93 DPS) | yes | Mixologist's Tunic (12793, -0.49 DPS) [dungeon]; Grizzled Pelt (22274, -0.50 DPS) [quest]; Valorous Chestguard (8274, -0.61 DPS, sim-verified) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.17 DPS) | yes | Branded Leather Bracers (19508, -0.62 DPS) [dungeon]; Officer's Wristguards (250581, -0.70 DPS) [crafted]; Runed Golem Shackles (12550, -2.35 DPS, sim-verified) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 37.5 attack_power points (2.90 DPS) | yes | Rockgrip Gauntlets (17736, -0.73 DPS) [dungeon]; Officer's Gloves (250551, -0.75 DPS) [crafted]; Gauntlets of Divinity (7724, -1.91 DPS, sim-verified) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Atal'alarion's Tusk Ring (10798, -0.13 DPS) [dungeon]; Belt of the Gladiator (13134, -0.13 DPS) [world_drop]; Girdle of Beastial Fury (11686, -2.51 DPS, sim-verified) [dungeon] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.6 attack_power points (2.68 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Silvershell Leggings (10633, -0.24 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.24 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Officer's Sabatons (250561, -0.04 DPS) [crafted]; Officer's Boots (250546, -0.09 DPS) [crafted]; Battlechaser's Greaves (12555, -3.41 DPS, sim-verified) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.86 DPS) | yes | Mark of Kern (2262, -0.31 DPS) [dungeon]; Assault Band (13095, -0.31 DPS) [world_drop]; Legionnaire's Band (19511, -0.44 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.63 DPS) | yes | Mark of Kern (2262, -0.08 DPS) [dungeon]; Legionnaire's Band (19511, -0.21 DPS) [rep]; Assault Band (13095, -1.38 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Blessed Prayer Beads (19990, -1.19 DPS, sim-verified) [quest] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Blight (7959, -4.35 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.52 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Ghostshard Talisman; shoulder: Blood Guard's Plate Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Bracers of the Stone Princess; hands: Raider Gauntlets; waist: Defiler's Leather Girdle; legs: Golem Shard Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 846, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 251.2. Weights run: 1.6s. Verify run: 1.3s. 1866 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.801, strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=1.448 ± 0.116 per rating point (14 rating = 1%, 20.276 per %), hit=0.304 ± 0.015 per rating point (10 rating = 1%, 3.037 per %), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lightbreaker Greathelm (239517) | Leonid Barthalomew the Revered [vendor] | 235.9 attack_power points (7.55 DPS) | yes | Warlord's Plate Headpiece (16542, -3.42 DPS) [vendor]; Warlord's Plate Headpiece (231535, -3.42 DPS) [pvp]; Lightbreaker Helmet (239525, -8.87 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-verified (251.2 DPS) | yes | Medallion of the Dawn (22659, -0.41 DPS) [quest]; Conqueror's Medallion (12059, -0.46 DPS) [quest]; Rage of Mugamba (19577, -2.25 DPS, sim-verified) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 169.8 attack_power points (5.44 DPS) | yes | Warlord's Plate Shoulders (16544, -2.61 DPS) [vendor]; Warlord's Plate Shoulders (231534, -2.61 DPS) [pvp]; Lightbreaker Pauldrons (239524, -7.73 DPS, sim-verified) [vendor] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 66.1 attack_power points (2.12 DPS) | yes | Shadewood Cloak (18328, +0.00 DPS, sim-verified) [dungeon]; Cloak of Revanchion (23127, -0.56 DPS) [dungeon]; Phantasmal Cloak (18689, -0.62 DPS) [dungeon] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | sim-verified (251.2 DPS) | yes | Timbermaw Tunic (252484, -2.46 DPS) [crafted]; Lightbreaker Breastplate (239527, -4.28 DPS) [vendor]; Breastplate of Undead Slaying (23087, -15.52 DPS, sim-verified) [world] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (251.2 DPS) | yes | Berserker Bracers (19578, -1.40 DPS) [rep]; General's Plate Armguards (16546, -1.68 DPS) [pvp]; Bracers of Undead Slaying (23090, -9.00 DPS, sim-verified) [world] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | sim-verified (251.2 DPS) | yes | Raider Gauntlets (272095, -1.55 DPS) [vendor]; Radiant Gloves of the Dawn (227817, -2.08 DPS) [vendor]; Razor Gauntlets (18326, -10.17 DPS, sim-verified) [dungeon] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 169.8 attack_power points (5.44 DPS) | yes | Ferocity of the Timbermaw (227805, -1.68 DPS) [vendor]; General's Plate Girdle (16547, -1.96 DPS) [pvp]; Radiant Girdle of the Dawn (227814, -2.45 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | sim-verified (251.2 DPS) | yes | Sentinel's Plate Legguards (237825, -3.07 DPS) [vendor]; Titanic Leggings (22385, -3.25 DPS) [crafted]; Cloudkeeper Legplates (14554, -16.39 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 168.9 attack_power points (5.41 DPS) | yes | General's Plate Boots (16545, -2.70 DPS) [vendor]; General's Plate Boots (231531, -2.70 DPS) [pvp]; Boots of Heroism (21995, -7.06 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (251.2 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.25 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234026, -0.37 DPS) [vendor]; Naglering (11669, -5.91 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (251.2 DPS) | yes | Band of the Ogre King (18522, -0.09 DPS) [dungeon]; Legionnaire's Band (19511, -0.31 DPS) [rep]; Naglering (11669, -3.16 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (251.2 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (251.2 DPS) | yes | Counterattack Lodestone (18537, -0.71 DPS) [dungeon]; Hand of Justice (11815, -0.77 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.10 DPS, sim-verified) [crafted] |
| main_hand | Nightfall (19169) | Blacksmithing [crafted] | sim-verified (251.2 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackfury (19167, -18.19 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | The Legend of Korrak [quest] | sim-verified (251.2 DPS) | yes | Unsophisticated Hand Cannon (18460, -0.22 DPS) [dungeon]; Monolithic Bow (9426, -0.37 DPS) [dungeon]; Dark Iron Rifle (16004, -3.12 DPS, sim-verified) [crafted] |

**New at 60:** head: Lightbreaker Greathelm; neck: Amulet of the Darkmoon; shoulder: Lightbreaker Shoulders; back: Shroud of Domination; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Legionnaire's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Nightfall; ranged: Bloodseeker

No-known-source sample (15 of 1866, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

