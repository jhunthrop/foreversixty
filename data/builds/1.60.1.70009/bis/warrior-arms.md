# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 30.4. Weights run: 1.2s. Verify run: 1.0s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.053 ± 0.002 per rating point (10 rating = 1%, 0.529 per %), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.09 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.3 attack_power points (0.26 DPS) | yes | Rough Bronze Shoulders (3480, +0.00 DPS, sim-verified) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Grave Shroud (279865, -0.05 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.18 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 attack_power points (0.88 DPS) | yes | Veteran's Chain Shirt (250488, -0.17 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Cryptwalker Bracers (280095, -0.05 DPS, sim-verified) [quest]; Bristlebark Bindings (14569, -0.26 DPS) [world_drop]; Bravo's Armbands (270015, -0.26 DPS) [quest] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 attack_power points (0.70 DPS) | yes | Gold-flecked Gloves (5195, -0.05 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -0.18 DPS) [quest]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.09 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Warchief's Girdle (5750, -0.32 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.9 attack_power points (0.96 DPS) | yes | Veteran's Chain Leggings (250493, -0.09 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.18 DPS) [crafted]; Totemic Leather Pants (252446, -0.18 DPS) [crafted] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 attack_power points (0.35 DPS) | yes | Ring of the Moon (12052, -0.17 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.26 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 323.2 attack_power points (13.58 DPS) | yes | Smite's Mighty Hammer (7230, -1.09 DPS, sim-verified) [dungeon]; Living Root (6631, -1.10 DPS) [dungeon]; Duskbringer (2205, -1.21 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567)) | Blacksmithing [crafted] | 4.2 attack_power points (0.18 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS, sim-verified) [quest]; Fine Longbow (11304, -0.01 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Veteran's Boots; finger1: Demon Band; finger2: Protector's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings

### Band 30 (human, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 70.7. Weights run: 1.3s. Verify run: 1.1s. 488 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.012 ± 0.002 per rating point (14 rating = 1%, 0.171 per %), hit=0.074 ± 0.003 per rating point (10 rating = 1%, 0.743 per %), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.26 DPS) | yes | Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted]; Veteran's Chain Helm (250498, -0.22 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.68 DPS) | yes | River Pride Choker (13087, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.38 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.68 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.49 DPS) | yes | Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.21 DPS, sim-verified) [pvp] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.9 attack_power points (1.46 DPS) | yes | Shining Silver Breastplate (2870, -0.22 DPS, sim-verified) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted]; Hard Gold Cuirass (250533, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 attack_power points (0.78 DPS) | yes | Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Yorgen Bracers (13012, -0.26 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.29 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.07 DPS) | yes | Bonefist Gauntlets (4465, -0.20 DPS) [world]; The Frozen Clutch (23170, -0.22 DPS, sim-verified) [dungeon]; Heavy Earthen Gloves (7359, -0.29 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.17 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.27 DPS) | yes | Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Slayer's Pants (14757, -0.20 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -0.25 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 attack_power points (0.68 DPS) | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 attack_power points (0.78 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.0 attack_power points (0.59 DPS) | yes | Tiger Band (6749, +0.00 DPS, sim-verified) [quest]; Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (70.7 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -20.81 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.44 DPS) | yes | Long Battle Bow (15284, -0.15 DPS) [world_drop]; Fine Longbow (11304, -0.24 DPS) [vendor]; Double-barreled Shotgun (2098, -0.31 DPS, sim-verified) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 488, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 90.3. Weights run: 1.6s. Verify run: 1.2s. 671 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.304, strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.034 ± 0.004 per rating point (14 rating = 1%, 0.473 per %), hit=0.114 ± 0.006 per rating point (10 rating = 1%, 1.140 per %), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Icemetal Barbute (10763) (or Hard Gold Coif (250537)) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 34.6 attack_power points (1.43 DPS) | yes | Hard Gold Coif (250537, +0.00 DPS, sim-verified) [crafted]; Raging Berserker's Helm (7719, -0.08 DPS) [dungeon]; Tusken Helm (6686, -0.10 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.83 DPS) | yes | Kaleidoscope Chain (13084, -0.42 DPS) [world_drop]; Gazlowe's Charm (13088, -0.42 DPS) [dungeon]; Ghostshard Talisman (7731, -0.72 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 attack_power points (1.12 DPS) | yes | Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [dungeon]; Chromite Pauldrons (8144, -0.23 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.9 attack_power points (0.61 DPS) | yes | Dark Hooded Cape (5257, -0.20 DPS) [world]; Sergeant Major's Cape (16315, -0.20 DPS) [pvp]; Wolfmaster Cape (6314, -0.28 DPS, sim-verified) [dungeon] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 37.1 attack_power points (1.53 DPS) | yes | Avenger's Armor (1488, +0.00 DPS, sim-verified) [dungeon]; Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Shining Mithril Breastplate (250540, -0.10 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Yorgen Bracers (13012, -0.21 DPS) [world_drop]; Pugilist Bracers (4438, -0.59 DPS, sim-verified) [dungeon] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 attack_power points (1.63 DPS) | yes | Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.72 DPS, sim-verified) [world_drop] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 37.1 attack_power points (1.53 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS, sim-verified) [dungeon]; Highlander's Leather Girdle (20116, -0.29 DPS) [rep]; Highlander's Lamellar Girdle (20107, -0.31 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 attack_power points (2.15 DPS) | yes | Orcish War Leggings (7929, -0.41 DPS) [crafted]; Firemane Leggings (13129, -0.59 DPS, sim-verified) [world_drop]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 attack_power points (1.33 DPS) | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.59 DPS, sim-verified) [crafted] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.83 DPS) | yes | Protector's Band (19515, -0.01 DPS) [rep]; Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor]; Protector's Band (19515, -0.54 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (90.3 DPS) | yes | Bonebiter (6830, -0.60 DPS) [quest]; Darkspear Raider's Reaper (272081, -0.83 DPS) [vendor]; Fiery War Axe (870, -10.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (90.3 DPS) | yes | The Silencer (13138, -0.04 DPS) [world_drop]; Mithril Blacksmith Hammer (285280, -0.10 DPS) [crafted]; Bow of Searing Arrows (2825, -0.90 DPS, sim-verified) [world_drop] |

**New at 40:** head: Icemetal Barbute; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Mark of Kern; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 671, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 129.5. Weights run: 1.7s. Verify run: 1.4s. 856 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.275, strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=0.440 ± 0.037 per rating point (14 rating = 1%, 6.161 per %), hit=0.098 ± 0.005 per rating point (10 rating = 1%, 0.978 per %), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.6 attack_power points (3.46 DPS) | yes | Fury Visor (20521, -1.22 DPS) [quest]; Knight-Lieutenant's Plate Helm (220804, -1.32 DPS) [vendor]; Bloomsprout Headpiece (17767, -1.69 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.55 DPS) | yes | Skibi's Pendant (13089, -0.65 DPS) [world_drop]; Ghostshard Talisman (7731, -0.69 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -0.97 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 23.5 attack_power points (1.82 DPS) | yes | Wyrmslayer Spaulders (13066, -0.18 DPS) [world_drop]; Earthslag Shoulders (11632, -0.23 DPS) [dungeon]; Officer's Pauldrons (250576, -1.13 DPS, sim-verified) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 attack_power points (1.10 DPS) | yes | Blackveil Cape (11626, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Cape (16336, -0.23 DPS) [pvp]; Wolfmaster Cape (6314, -0.32 DPS) [dungeon] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 37.8 attack_power points (2.93 DPS) | yes | Mixologist's Tunic (12793, -0.49 DPS) [dungeon]; Coldmetal Guard (274758, -0.61 DPS) [vendor]; Valorous Chestguard (8274, -0.79 DPS, sim-verified) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) (or Arena Bands (18711)) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.17 DPS) | yes | Arena Bands (18711, +0.00 DPS, sim-verified) [world]; Runed Golem Shackles (12550, -0.46 DPS) [dungeon]; Branded Leather Bracers (19508, -0.62 DPS) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 37.5 attack_power points (2.90 DPS) | yes | Rockgrip Gauntlets (17736, -0.73 DPS) [dungeon]; Officer's Gloves (250551, -0.75 DPS) [crafted]; Gauntlets of Divinity (7724, -2.84 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 42.6 attack_power points (3.30 DPS) | yes | Atal'alarion's Tusk Ring (10798, -1.11 DPS) [dungeon]; Belt of the Gladiator (13134, -1.11 DPS) [world_drop]; Highlander's Leather Girdle (20116, -1.53 DPS, sim-verified) [rep] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.6 attack_power points (2.68 DPS) | yes | Scarlet Leggings (10330, -0.16 DPS, sim-verified) [dungeon]; Silvershell Leggings (10633, -0.24 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.24 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 25.7 attack_power points (1.99 DPS) | yes | Officer's Sabatons (250561, -0.21 DPS) [crafted]; Officer's Boots (250546, -0.25 DPS) [crafted]; Prowler's Leather Boots (252468, -0.52 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.63 DPS) | yes | Mark of Kern (2262, -0.08 DPS) [dungeon]; Protector's Band (19516, -0.21 DPS) [rep]; Insurgent's Band (272065, -0.46 DPS) [vendor] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (1.55 DPS) | yes | Mark of Kern (2262, +0.00 DPS, sim-verified) [dungeon]; Protector's Band (19516, -0.13 DPS) [rep]; Insurgent's Band (272065, -0.39 DPS) [vendor] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (129.5 DPS) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (129.5 DPS) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -7.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (129.5 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.87 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Bracers of the Stone Princess; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Assault Band; trinket1: Diamond Flask; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 856, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 245.3. Weights run: 1.7s. Verify run: 1.4s. 1891 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.801, strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=1.448 ± 0.116 per rating point (14 rating = 1%, 20.276 per %), hit=0.304 ± 0.015 per rating point (10 rating = 1%, 3.037 per %), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lightbreaker Greathelm (239517) | Leonid Barthalomew the Revered [vendor] | 235.9 attack_power points (7.55 DPS) | yes | Field Marshal's Plate Helm (231538, -3.42 DPS) [pvp]; Lightbreaker Helmet (239525, -9.53 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-verified (245.3 DPS) | yes | Medallion of the Dawn (22659, -0.41 DPS) [quest]; Strength of Mugamba (19576, -0.59 DPS) [quest]; Rage of Mugamba (19577, -1.96 DPS, sim-verified) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 169.8 attack_power points (5.44 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, -2.61 DPS) [pvp]; Lightbreaker Pauldrons (239524, -8.55 DPS, sim-verified) [vendor] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 66.1 attack_power points (2.12 DPS) | yes | Shadewood Cloak (18328, +0.00 DPS, sim-verified) [dungeon]; Cloak of Revanchion (23127, -0.56 DPS) [dungeon]; Phantasmal Cloak (18689, -0.62 DPS) [dungeon] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | sim-verified (245.3 DPS) | yes | Timbermaw Tunic (252484, -2.46 DPS) [crafted]; Lightbreaker Breastplate (239527, -4.28 DPS) [vendor]; Breastplate of Undead Slaying (23087, -16.39 DPS, sim-verified) [world] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (245.3 DPS) | yes | Berserker Bracers (19578, -1.40 DPS) [rep]; Marshal's Plate Bracers (16481, -1.68 DPS) [pvp]; Bracers of Undead Slaying (23090, -10.03 DPS, sim-verified) [world] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | sim-verified (245.3 DPS) | yes | Raider Gauntlets (272095, -1.55 DPS) [vendor]; Radiant Gloves of the Dawn (227817, -2.08 DPS) [vendor]; Marshal's Plate Gauntlets (16484, -3.77 DPS, sim-verified) [vendor] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 169.8 attack_power points (5.44 DPS) | yes | Ferocity of the Timbermaw (227805, -1.68 DPS) [vendor]; Marshal's Plate Girdle (16482, -1.96 DPS) [pvp]; Radiant Girdle of the Dawn (227814, -3.20 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | sim-verified (245.3 DPS) | yes | Sentinel's Plate Legguards (237825, -3.07 DPS) [vendor]; Titanic Leggings (22385, -3.25 DPS) [crafted]; Cloudkeeper Legplates (14554, -7.36 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 168.9 attack_power points (5.41 DPS) | yes | Marshal's Plate Boots (231539, -2.70 DPS) [pvp]; Boots of Heroism (21995, -2.82 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (245.3 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.25 DPS) [vendor]; Naglering (11669, -6.04 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (245.3 DPS) | yes | Band of the Ogre King (18522, -0.09 DPS) [dungeon]; Protector's Band (19516, -0.31 DPS) [rep]; Naglering (11669, -3.71 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (245.3 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.55 DPS) [crafted]; Blackhand's Breadth (13965, -2.45 DPS, sim-verified) [quest] |
| trinket2 | - | - |  |  |  |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (245.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -2.53 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-verified (245.3 DPS) | yes | Unsophisticated Hand Cannon (18460, -0.22 DPS) [dungeon]; Monolithic Bow (9426, -0.37 DPS) [dungeon]; Dark Iron Rifle (16004, -3.90 DPS, sim-verified) [crafted] |

**New at 60:** head: Lightbreaker Greathelm; neck: Amulet of the Darkmoon; shoulder: Lightbreaker Shoulders; back: Shroud of Domination; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Protector's Band; trinket1: Hand of Justice; main_hand: Blackblade of Shahram; ranged: Bloodseeker

No-known-source sample (15 of 1891, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 33.5. Weights run: 1.2s. Verify run: 0.9s. 285 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.053 ± 0.002 per rating point (10 rating = 1%, 0.529 per %), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.22 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.3 attack_power points (0.26 DPS) | yes | Rough Bronze Shoulders (3480, +0.00 DPS, sim-verified) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Subterranean Cape (14149, -0.09 DPS) [dungeon]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Grave Shroud (279865, -0.12 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.32 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Raptorcrest Bracers (270010, -0.22 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted]; Bristlebark Bindings (14569, -0.26 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 attack_power points (0.70 DPS) | yes | Gold-flecked Gloves (5195, -0.12 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.26 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.16 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Warchief's Girdle (5750, -0.32 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.8 attack_power points (0.79 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.09 DPS) [world_drop] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 323.2 attack_power points (13.58 DPS) | yes | Forsaken Greataxe (251533, -0.82 DPS) [quest]; Smite's Mighty Hammer (7230, -1.02 DPS) [dungeon]; Hammerbone (270018, -3.33 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.2 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Heavy Shortbow (3036, -0.09 DPS) [world_drop]; Orcish Battle Bow (5346, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Veteran's Boots; finger1: Demon Band; finger2: Legionnaire's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 285, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 72.1. Weights run: 1.3s. Verify run: 1.1s. 471 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.012 ± 0.002 per rating point (14 rating = 1%, 0.171 per %), hit=0.074 ± 0.003 per rating point (10 rating = 1%, 0.743 per %), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.26 DPS) | yes | Veteran's Chain Helm (250498, -0.09 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.68 DPS) | yes | River Pride Choker (13087, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.40 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.68 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.49 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.9 attack_power points (1.46 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS, sim-verified) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted]; Hard Gold Cuirass (250533, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 attack_power points (0.78 DPS) | yes | Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Yorgen Bracers (13012, -0.22 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.29 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.07 DPS) | yes | The Frozen Clutch (23170, -0.08 DPS, sim-verified) [dungeon]; Warsong Gauntlets (16978, -0.10 DPS) [quest]; Bonefist Gauntlets (4465, -0.20 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.17 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.27 DPS) | yes | Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Slayer's Pants (14757, -0.20 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -0.21 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 attack_power points (0.68 DPS) | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop]; Slayer's Slippers (14756, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 attack_power points (0.78 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.0 attack_power points (0.59 DPS) | yes | Tiger Band (6749, +0.00 DPS, sim-verified) [quest]; Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.1 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -21.33 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.44 DPS) | yes | Long Battle Bow (15284, -0.15 DPS) [world_drop]; Double-barreled Shotgun (2098, -0.15 DPS, sim-verified) [world_drop]; Fine Longbow (11304, -0.24 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 471, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 89.0. Weights run: 1.6s. Verify run: 1.2s. 652 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.304, strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.034 ± 0.004 per rating point (14 rating = 1%, 0.473 per %), hit=0.114 ± 0.006 per rating point (10 rating = 1%, 1.140 per %), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Icemetal Barbute (10763) (or Hard Gold Coif (250537)) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 34.6 attack_power points (1.43 DPS) | yes | Hard Gold Coif (250537, +0.00 DPS, sim-verified) [crafted]; Raging Berserker's Helm (7719, -0.08 DPS) [dungeon]; Tusken Helm (6686, -0.10 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.83 DPS) | yes | Ethereal Talisman (4430, -0.31 DPS) [quest]; Kaleidoscope Chain (13084, -0.42 DPS) [world_drop]; Ghostshard Talisman (7731, -0.49 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 attack_power points (1.12 DPS) | yes | Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [dungeon]; Chromite Pauldrons (8144, -0.38 DPS, sim-verified) [dungeon] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.41 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Dark Hooded Cape (5257, -0.00 DPS) [world]; Khan's Cloak (14781, -0.00 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 37.1 attack_power points (1.53 DPS) | yes | Avenger's Armor (1488, +0.00 DPS, sim-verified) [dungeon]; Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Shining Mithril Breastplate (250540, -0.10 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Darkspear Armsplints (4132, -0.11 DPS) [quest]; Pugilist Bracers (4438, -0.33 DPS, sim-verified) [dungeon] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 attack_power points (1.63 DPS) | yes | Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.58 DPS, sim-verified) [world_drop] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 37.1 attack_power points (1.53 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS, sim-verified) [dungeon]; Tharg's Shoelace (9705, -0.20 DPS) [quest]; Defiler's Leather Girdle (20192, -0.29 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 attack_power points (2.15 DPS) | yes | Firemane Leggings (13129, -0.38 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.41 DPS) [crafted]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 attack_power points (1.33 DPS) | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.38 DPS, sim-verified) [crafted] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.83 DPS) | yes | Legionnaire's Band (19512, -0.01 DPS) [rep]; Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor]; Legionnaire's Band (19512, -0.32 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (89.0 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -1.33 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (89.0 DPS) | yes | The Silencer (13138, -0.04 DPS) [world_drop]; Mithril Blacksmith Hammer (285280, -0.10 DPS) [crafted]; Bow of Searing Arrows (2825, -1.03 DPS, sim-verified) [world_drop] |

**New at 40:** head: Icemetal Barbute; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Mark of Kern; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 652, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 130.6. Weights run: 1.7s. Verify run: 1.4s. 837 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.275, strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=0.440 ± 0.037 per rating point (14 rating = 1%, 6.161 per %), hit=0.098 ± 0.005 per rating point (10 rating = 1%, 0.978 per %), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.6 attack_power points (3.46 DPS) | yes | Bloomsprout Headpiece (17767, -0.82 DPS, sim-verified) [dungeon]; Fury Visor (20521, -1.22 DPS) [quest]; Blood Guard's Plate Helm (220803, -1.32 DPS) [vendor] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.55 DPS) | yes | Ghostshard Talisman (7731, -0.15 DPS, sim-verified) [dungeon]; Woven Ivy Necklace (19159, -0.62 DPS) [quest]; Skibi's Pendant (13089, -0.65 DPS) [world_drop] |
| shoulder | Blood Guard's Plate Pauldrons (220796) | Lady Palanseer [vendor] | 23.5 attack_power points (1.82 DPS) | yes | Wyrmslayer Spaulders (13066, -0.18 DPS) [world_drop]; Earthslag Shoulders (11632, -0.23 DPS) [dungeon]; Officer's Pauldrons (250576, -0.67 DPS, sim-verified) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 attack_power points (1.10 DPS) | yes | Blackveil Cape (11626, +0.00 DPS, sim-verified) [dungeon]; Battlehard Cape (11858, -0.32 DPS) [quest]; Wildhunter Cloak (16658, -0.32 DPS) [quest] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 37.8 attack_power points (2.93 DPS) | yes | Valorous Chestguard (8274, -0.48 DPS, sim-verified) [world_drop]; Mixologist's Tunic (12793, -0.49 DPS) [dungeon]; Coldmetal Guard (274758, -0.61 DPS) [vendor] |
| wrist | Bracers of the Stone Princess (17714) (or Arena Bands (18711)) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.17 DPS) | yes | Arena Bands (18711, +0.00 DPS, sim-verified) [world]; Runed Golem Shackles (12550, -0.46 DPS) [dungeon]; Branded Leather Bracers (19508, -0.62 DPS) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 37.5 attack_power points (2.90 DPS) | yes | Rockgrip Gauntlets (17736, -0.73 DPS) [dungeon]; Officer's Gloves (250551, -0.75 DPS) [crafted]; Gauntlets of Divinity (7724, -2.39 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 42.6 attack_power points (3.30 DPS) | yes | Defiler's Leather Girdle (20192, -0.93 DPS, sim-verified) [rep]; Atal'alarion's Tusk Ring (10798, -1.11 DPS) [dungeon]; Belt of the Gladiator (13134, -1.11 DPS) [world_drop] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.6 attack_power points (2.68 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Silvershell Leggings (10633, -0.24 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.24 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 25.7 attack_power points (1.99 DPS) | yes | Officer's Sabatons (250561, -0.21 DPS) [crafted]; Prowler's Leather Boots (252468, -0.23 DPS, sim-verified) [crafted]; Officer's Boots (250546, -0.25 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.86 DPS) | yes | Mark of Kern (2262, -0.31 DPS) [dungeon]; Assault Band (13095, -0.31 DPS) [world_drop]; Legionnaire's Band (19511, -0.44 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.63 DPS) | yes | Mark of Kern (2262, -0.08 DPS) [dungeon]; Legionnaire's Band (19511, -0.21 DPS) [rep]; Assault Band (13095, -1.12 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (130.6 DPS) | yes | Frozen Heart of the Mountain (249469, -3.24 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (130.6 DPS) | yes | Frozen Heart of the Mountain (249469, -2.80 DPS, sim-verified) [crafted] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (130.6 DPS) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -4.01 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (130.6 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.15 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Blood Guard's Plate Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Bracers of the Stone Princess; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 837, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 237.7. Weights run: 1.7s. Verify run: 1.4s. 1872 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.801, strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=1.448 ± 0.116 per rating point (14 rating = 1%, 20.276 per %), hit=0.304 ± 0.015 per rating point (10 rating = 1%, 3.037 per %), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lightbreaker Greathelm (239517) | Leonid Barthalomew the Revered [vendor] | 235.9 attack_power points (7.55 DPS) | yes | Warlord's Plate Headpiece (231535, -3.42 DPS) [pvp]; Lightbreaker Helmet (239525, -8.77 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-verified (237.7 DPS) | yes | Medallion of the Dawn (22659, -0.41 DPS) [quest]; Conqueror's Medallion (12059, -0.46 DPS) [quest]; Rage of Mugamba (19577, -2.40 DPS, sim-verified) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 169.8 attack_power points (5.44 DPS) | yes | Warlord's Plate Shoulders (231534, -2.61 DPS) [pvp]; Lightbreaker Pauldrons (239524, -8.62 DPS, sim-verified) [vendor] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 66.1 attack_power points (2.12 DPS) | yes | Shadewood Cloak (18328, -0.08 DPS, sim-verified) [dungeon]; Cloak of Revanchion (23127, -0.56 DPS) [dungeon]; Phantasmal Cloak (18689, -0.62 DPS) [dungeon] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | sim-verified (237.7 DPS) | yes | Timbermaw Tunic (252484, -2.46 DPS) [crafted]; Lightbreaker Breastplate (239527, -4.28 DPS) [vendor]; Breastplate of Undead Slaying (23087, -19.04 DPS, sim-verified) [world] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (237.7 DPS) | yes | Berserker Bracers (19578, -1.40 DPS) [rep]; General's Plate Armguards (16546, -1.68 DPS) [pvp]; Bracers of Undead Slaying (23090, -9.70 DPS, sim-verified) [world] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | sim-verified (237.7 DPS) | yes | Raider Gauntlets (272095, -1.55 DPS) [vendor]; Radiant Gloves of the Dawn (227817, -2.08 DPS) [vendor]; General's Plate Gauntlets (16548, -4.26 DPS, sim-verified) [vendor] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 169.8 attack_power points (5.44 DPS) | yes | Ferocity of the Timbermaw (227805, -1.68 DPS) [vendor]; General's Plate Girdle (16547, -1.96 DPS) [pvp]; Radiant Girdle of the Dawn (227814, -3.93 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | sim-verified (237.7 DPS) | yes | Sentinel's Plate Legguards (237825, -3.07 DPS) [vendor]; Titanic Leggings (22385, -3.25 DPS) [crafted]; Cloudkeeper Legplates (14554, -9.71 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 168.9 attack_power points (5.41 DPS) | yes | General's Plate Boots (231531, -2.70 DPS) [pvp]; Boots of Heroism (21995, -2.82 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (237.7 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.25 DPS) [vendor]; Naglering (11669, -5.96 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (237.7 DPS) | yes | Band of the Ogre King (18522, -0.09 DPS) [dungeon]; Legionnaire's Band (19511, -0.31 DPS) [rep]; Naglering (11669, -4.54 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (237.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -1.57 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (237.7 DPS) | yes | Hand of Justice (11815, -0.25 DPS, sim-verified) [dungeon]; Counterattack Lodestone (18537, -0.71 DPS) [dungeon]; Blackhand's Breadth (13965, -1.32 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (237.7 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -8.40 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | The Legend of Korrak [quest] | sim-verified (237.7 DPS) | yes | Unsophisticated Hand Cannon (18460, -0.22 DPS) [dungeon]; Monolithic Bow (9426, -0.37 DPS) [dungeon]; Dark Iron Rifle (16004, -3.82 DPS, sim-verified) [crafted] |

**New at 60:** head: Lightbreaker Greathelm; neck: Amulet of the Darkmoon; shoulder: Lightbreaker Shoulders; back: Shroud of Domination; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Legionnaire's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Bloodseeker

No-known-source sample (15 of 1872, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

