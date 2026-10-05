# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 30.4. Weights run: 0.9s. Verify run: 0.8s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.053 ± 0.002 per rating point (10 rating = 1%, 0.529 per %), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.18 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 attack_power points (0.26 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Grave Shroud (279865, -0.09 DPS) [quest]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.18 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 attack_power points (0.88 DPS) | yes | Veteran's Chain Shirt (250488, -0.17 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Cryptwalker Bracers (280095, -0.09 DPS) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted]; Bristlebark Bindings (14569, -0.26 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 attack_power points (0.70 DPS) | yes | Gold-flecked Gloves (5195, -0.09 DPS) [dungeon]; Polar Gauntlets (7606, -0.18 DPS) [quest]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Hulking Belt (14746, -0.32 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.9 attack_power points (0.96 DPS) | yes | Veteran's Chain Leggings (250493, -0.18 DPS) [crafted]; Defender's Leather Pants (252445, -0.18 DPS) [crafted]; Totemic Leather Pants (252446, -0.18 DPS) [crafted] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 attack_power points (0.35 DPS) | yes | Ring of the Moon (12052, -0.17 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.26 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 323.2 attack_power points (13.58 DPS) | yes | Smite's Mighty Hammer (7230, -1.09 DPS, sim-verified) [dungeon]; Living Root (6631, -1.10 DPS) [dungeon]; Duskbringer (2205, -1.21 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567)) | Blacksmithing [crafted] | 4.2 attack_power points (0.18 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, -0.01 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Protector's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 70.7. Weights run: 1.1s. Verify run: 0.9s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.012 ± 0.002 per rating point (14 rating = 1%, 0.171 per %), hit=0.074 ± 0.003 per rating point (10 rating = 1%, 0.743 per %), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.26 DPS) | yes | Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted]; Veteran's Chain Helm (250498, -0.22 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.68 DPS) | yes | River Pride Choker (13087, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.38 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.68 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.10 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.49 DPS) | yes | Sergeant Major's Cape (16315, -0.10 DPS) [pvp]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.9 attack_power points (1.46 DPS) | yes | Shining Silver Breastplate (2870, -0.22 DPS, sim-verified) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted]; Hard Gold Cuirass (250533, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 attack_power points (0.78 DPS) | yes | Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Yorgen Bracers (13012, -0.26 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.29 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.07 DPS) | yes | Bonefist Gauntlets (4465, -0.20 DPS) [world]; The Frozen Clutch (23170, -0.22 DPS, sim-verified) [dungeon]; Heavy Earthen Gloves (7359, -0.29 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.17 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.27 DPS) | yes | Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Chausses of Westfall (6087, -0.20 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.25 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 attack_power points (0.68 DPS) | yes | Hard Gold Boots (250534, -0.00 DPS) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 attack_power points (0.78 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.0 attack_power points (0.59 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (70.7 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -20.81 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.44 DPS) | yes | Long Battle Bow (15284, -0.15 DPS) [world_drop]; Fine Longbow (11304, -0.24 DPS) [vendor]; Double-barreled Shotgun (2098, -0.31 DPS, sim-verified) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 90.3. Weights run: 1.3s. Verify run: 1.0s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.304, strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.034 ± 0.004 per rating point (14 rating = 1%, 0.473 per %), hit=0.114 ± 0.006 per rating point (10 rating = 1%, 1.140 per %), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 34.6 attack_power points (1.43 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.08 DPS) [dungeon]; Tusken Helm (6686, -0.10 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.83 DPS) | yes | Kaleidoscope Chain (13084, -0.42 DPS) [world_drop]; River Pride Choker (13087, -0.42 DPS) [world_drop]; Ghostshard Talisman (7731, -0.72 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 attack_power points (1.12 DPS) | yes | Chromite Pauldrons (8144, -0.10 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.9 attack_power points (0.61 DPS) | yes | Wolfmaster Cape (6314, -0.20 DPS) [dungeon]; Dark Hooded Cape (5257, -0.20 DPS) [world]; Slayer's Cape (14752, -0.21 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 37.1 attack_power points (1.53 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Shining Mithril Breastplate (250540, -0.10 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Yorgen Bracers (13012, -0.21 DPS) [world_drop]; Pugilist Bracers (4438, -0.59 DPS, sim-verified) [dungeon] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 attack_power points (1.63 DPS) | yes | Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.72 DPS, sim-verified) [world_drop] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 37.1 attack_power points (1.53 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.29 DPS) [rep]; Scarlet Belt (10329, -0.31 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 attack_power points (2.15 DPS) | yes | Orcish War Leggings (7929, -0.41 DPS) [crafted]; Firemane Leggings (13129, -0.59 DPS, sim-verified) [world_drop]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 attack_power points (1.33 DPS) | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.59 DPS, sim-verified) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Protector's Band (19515, -0.01 DPS) [rep]; Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.83 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor]; Protector's Band (19515, -0.54 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (90.3 DPS) | yes | Bonebiter (6830, -0.60 DPS) [quest]; Darkspear Raider's Reaper (272081, -0.83 DPS) [vendor]; Fiery War Axe (870, -10.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (90.3 DPS) | yes | The Silencer (13138, -0.04 DPS) [world_drop]; Mithril Blacksmith Hammer (285280, -0.10 DPS) [crafted]; Bow of Searing Arrows (2825, -0.90 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 129.5. Weights run: 1.3s. Verify run: 1.1s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.275, strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=0.440 ± 0.037 per rating point (14 rating = 1%, 6.161 per %), hit=0.098 ± 0.005 per rating point (10 rating = 1%, 0.978 per %), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.6 attack_power points (3.46 DPS) | yes | Fury Visor (20521, -1.22 DPS) [quest]; Knight-Lieutenant's Plate Helm (220804, -1.32 DPS) [vendor]; Bloomsprout Headpiece (17767, -1.69 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.55 DPS) | yes | Skibi's Pendant (13089, -0.65 DPS) [world_drop]; Ghostshard Talisman (7731, -0.69 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -0.97 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 23.5 attack_power points (1.82 DPS) | yes | Wyrmslayer Spaulders (13066, -0.18 DPS) [world_drop]; Earthslag Shoulders (11632, -0.23 DPS) [dungeon]; Officer's Pauldrons (250576, -1.13 DPS, sim-verified) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 attack_power points (1.10 DPS) | yes | Blackveil Cape (11626, -0.06 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.23 DPS) [pvp]; Wolfmaster Cape (6314, -0.32 DPS) [dungeon] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 37.8 attack_power points (2.93 DPS) | yes | Mixologist's Tunic (12793, -0.49 DPS) [dungeon]; Coldmetal Guard (274758, -0.61 DPS) [vendor]; Valorous Chestguard (8274, -0.79 DPS, sim-verified) [world_drop] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.17 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Runed Golem Shackles (12550, -0.46 DPS) [dungeon]; Branded Leather Bracers (19508, -0.62 DPS) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 37.5 attack_power points (2.90 DPS) | yes | Rockgrip Gauntlets (17736, -0.73 DPS) [dungeon]; Officer's Gloves (250551, -0.75 DPS) [crafted]; Gauntlets of Divinity (7724, -2.84 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 42.6 attack_power points (3.30 DPS) | yes | Atal'alarion's Tusk Ring (10798, -1.11 DPS) [dungeon]; Belt of the Gladiator (13134, -1.11 DPS) [world_drop]; Highlander's Leather Girdle (20116, -1.53 DPS, sim-verified) [rep] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.6 attack_power points (2.68 DPS) | yes | Scarlet Leggings (10330, -0.12 DPS) [dungeon]; Silvershell Leggings (10633, -0.24 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.24 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 25.7 attack_power points (1.99 DPS) | yes | Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Officer's Sabatons (250561, -0.21 DPS) [crafted]; Officer's Boots (250546, -0.25 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.63 DPS) | yes | Assault Band (13095, -0.08 DPS) [world_drop]; Protector's Band (19516, -0.21 DPS) [rep]; Insurgent's Band (272065, -0.46 DPS) [vendor] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.55 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Protector's Band (19516, -0.13 DPS) [rep]; Insurgent's Band (272065, -0.39 DPS) [vendor] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (129.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (129.5 DPS) | yes | Molten Heart of the Mountain (249470, -1.20 DPS, sim-verified) [crafted] |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (129.5 DPS) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -7.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (129.5 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.87 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Mark of Kern; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 201.6. Weights run: 1.3s. Verify run: 1.1s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.801, strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=1.448 ± 0.116 per rating point (14 rating = 1%, 20.276 per %), hit=0.304 ± 0.015 per rating point (10 rating = 1%, 3.037 per %), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 116.6 attack_power points (3.73 DPS) | yes | Field Marshal's Plate Helm (231538, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -0.37 DPS) [vendor]; Crown of Heroism (226860, -6.11 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-verified (201.6 DPS) | yes | Medallion of the Dawn (22659, -0.41 DPS) [quest]; Strength of Mugamba (19576, -0.59 DPS) [quest]; Rage of Mugamba (19577, -0.59 DPS) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 86.3 attack_power points (2.76 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 66.1 attack_power points (2.12 DPS) | yes | Shadewood Cloak (18328, -0.50 DPS) [dungeon]; Cloak of Revanchion (23127, -0.56 DPS) [dungeon]; Phantasmal Cloak (18689, -0.62 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (201.6 DPS) | yes | Savage Gladiator Chain (11726, -1.91 DPS) [dungeon]; Knight-Captain's Plate Hauberk (227046, -2.05 DPS) [pvp]; Breastplate of Undead Slaying (23087, -10.57 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-verified (201.6 DPS) | yes | Marshal's Plate Bracers (16481, -0.28 DPS) [pvp]; Gordok Bracers of Power (18533, -0.49 DPS) [dungeon]; Bracers of Undead Slaying (23090, -4.29 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (201.6 DPS) | yes | Radiant Gloves of the Dawn (227817, -0.53 DPS) [vendor]; Marshal's Plate Gauntlets (231541, -0.56 DPS) [pvp]; Razor Gauntlets (18326, -4.76 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 125.2 attack_power points (4.01 DPS) | yes | Ferocity of the Timbermaw (227805, -0.25 DPS) [vendor]; Marshal's Plate Girdle (16482, -0.53 DPS) [pvp]; Heavy Obsidian Belt (22197, -0.90 DPS) [crafted] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (201.6 DPS) | yes | Titanic Leggings (22385, -0.18 DPS) [crafted]; Warbear Woolies (15065, -0.72 DPS) [crafted]; Cloudkeeper Legplates (14554, -4.36 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 80.8 attack_power points (2.59 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Clutchlord's Stompers (275627, -0.22 DPS) [crafted] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (201.6 DPS) | yes | Myrmidon's Signet (2246, -0.37 DPS) [world_drop]; Don Julio's Band (19325, -0.57 DPS) [rep]; Naglering (11669, -3.97 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-verified (201.6 DPS) | yes | Myrmidon's Signet (2246, -0.28 DPS) [world_drop]; Don Julio's Band (19325, -0.48 DPS) [rep]; Naglering (11669, -2.41 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (201.6 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (201.6 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -0.55 DPS) [quest]; Frozen Heart of the Mountain (249469, -4.59 DPS, sim-verified) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (201.6 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -2.38 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-verified (201.6 DPS) | yes | Unsophisticated Hand Cannon (18460, -0.22 DPS) [dungeon]; Monolithic Bow (9426, -0.37 DPS) [dungeon]; Dark Iron Rifle (16004, -2.04 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Amulet of the Darkmoon; shoulder: Highlander's Plate Spaulders; back: Shroud of Domination; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Protector's Band; finger2: Band of the Ogre King; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Blackblade of Shahram; ranged: Bloodseeker

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 33.5. Weights run: 0.9s. Verify run: 0.8s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.053 ± 0.002 per rating point (10 rating = 1%, 0.529 per %), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.22 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 attack_power points (0.26 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Grave Shroud (279865, -0.09 DPS) [quest]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Subterranean Cape (14149, -0.12 DPS, sim-verified) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.32 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Raptorcrest Bracers (270010, -0.22 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted]; Bristlebark Bindings (14569, -0.26 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 attack_power points (0.70 DPS) | yes | Gold-flecked Gloves (5195, -0.12 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon]; Foreman's Gloves (2167, -0.26 DPS) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.16 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Hulking Belt (14746, -0.32 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.8 attack_power points (0.79 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.09 DPS) [world_drop] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 323.2 attack_power points (13.58 DPS) | yes | Forsaken Greataxe (251533, -0.82 DPS) [quest]; Smite's Mighty Hammer (7230, -1.02 DPS) [dungeon]; Hammerbone (270018, -3.33 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.2 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Heavy Shortbow (3036, -0.09 DPS) [world_drop]; Orcish Battle Bow (5346, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Legionnaire's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 72.1. Weights run: 1.1s. Verify run: 0.9s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.012 ± 0.002 per rating point (14 rating = 1%, 0.171 per %), hit=0.074 ± 0.003 per rating point (10 rating = 1%, 0.743 per %), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.26 DPS) | yes | Veteran's Chain Helm (250498, -0.10 DPS) [crafted]; Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.68 DPS) | yes | River Pride Choker (13087, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.40 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.68 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.49 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.9 attack_power points (1.46 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted]; Hard Gold Cuirass (250533, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 attack_power points (0.78 DPS) | yes | Yorgen Bracers (13012, -0.19 DPS) [world_drop]; Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Cultist's Armguards (270032, -0.29 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.07 DPS) | yes | The Frozen Clutch (23170, -0.10 DPS) [dungeon]; Warsong Gauntlets (16978, -0.10 DPS) [quest]; Bonefist Gauntlets (4465, -0.20 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.17 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.27 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.20 DPS) [crafted]; Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Slayer's Pants (14757, -0.20 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 attack_power points (0.68 DPS) | yes | Hard Gold Boots (250534, -0.00 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop]; Slayer's Slippers (14756, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 attack_power points (0.78 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.0 attack_power points (0.59 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.1 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -21.33 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.44 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS) [world_drop]; Long Battle Bow (15284, -0.15 DPS) [world_drop]; Fine Longbow (11304, -0.24 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 89.1. Weights run: 1.3s. Verify run: 1.0s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.304, strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.034 ± 0.004 per rating point (14 rating = 1%, 0.473 per %), hit=0.114 ± 0.006 per rating point (10 rating = 1%, 1.140 per %), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 34.6 attack_power points (1.43 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.08 DPS) [dungeon]; Tusken Helm (6686, -0.10 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.83 DPS) | yes | Ethereal Talisman (4430, -0.31 DPS) [quest]; Kaleidoscope Chain (13084, -0.42 DPS) [world_drop]; Ghostshard Talisman (7731, -0.58 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 attack_power points (1.12 DPS) | yes | Chromite Pauldrons (8144, -0.10 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 14.9 attack_power points (0.61 DPS) | yes | Wolfmaster Cape (6314, -0.20 DPS) [dungeon]; Wildhunter Cloak (16658, -0.20 DPS) [quest]; Dark Hooded Cape (5257, -0.20 DPS) [world] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 37.1 attack_power points (1.53 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Shining Mithril Breastplate (250540, -0.10 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Pugilist Bracers (4438, -0.40 DPS, sim-verified) [dungeon] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 attack_power points (1.63 DPS) | yes | Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.51 DPS, sim-verified) [world_drop] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 37.1 attack_power points (1.53 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Tharg's Shoelace (9705, -0.20 DPS) [quest]; Defiler's Leather Girdle (20192, -0.29 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 attack_power points (2.15 DPS) | yes | Orcish War Leggings (7929, -0.41 DPS) [crafted]; Firemane Leggings (13129, -0.44 DPS, sim-verified) [world_drop]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 attack_power points (1.33 DPS) | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.44 DPS, sim-verified) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Legionnaire's Band (19512, -0.01 DPS) [rep]; Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.83 DPS) | yes | Legionnaire's Band (19512, -0.01 DPS) [rep]; Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (89.1 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -1.19 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (89.1 DPS) | yes | The Silencer (13138, -0.04 DPS) [world_drop]; Mithril Blacksmith Hammer (285280, -0.10 DPS) [crafted]; Bow of Searing Arrows (2825, -0.96 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 129.9. Weights run: 1.3s. Verify run: 1.1s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.275, strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=0.440 ± 0.037 per rating point (14 rating = 1%, 6.161 per %), hit=0.098 ± 0.005 per rating point (10 rating = 1%, 0.978 per %), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 44.6 attack_power points (3.46 DPS) | yes | Bloomsprout Headpiece (17767, -0.99 DPS, sim-verified) [dungeon]; Fury Visor (20521, -1.22 DPS) [quest]; Blood Guard's Plate Helm (220803, -1.32 DPS) [vendor] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.55 DPS) | yes | Ghostshard Talisman (7731, -0.47 DPS) [dungeon]; Woven Ivy Necklace (19159, -0.62 DPS) [quest]; Skibi's Pendant (13089, -0.65 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 22.7 attack_power points (1.76 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.12 DPS) [world_drop]; Earthslag Shoulders (11632, -0.18 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 attack_power points (1.10 DPS) | yes | Blackveil Cape (11626, +0.00 DPS, sim-verified) [dungeon]; First Sergeant's Cloak (16340, -0.23 DPS) [pvp]; Battlehard Cape (11858, -0.32 DPS) [quest] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 37.8 attack_power points (2.93 DPS) | yes | Valorous Chestguard (8274, -0.49 DPS) [world_drop]; Mixologist's Tunic (12793, -0.49 DPS) [dungeon]; Coldmetal Guard (274758, -0.61 DPS) [vendor] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.17 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 37.5 attack_power points (2.90 DPS) | yes | Rockgrip Gauntlets (17736, -0.73 DPS) [dungeon]; Officer's Gloves (250551, -0.75 DPS) [crafted]; Gauntlets of Divinity (7724, -2.29 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 42.6 attack_power points (3.30 DPS) | yes | Atal'alarion's Tusk Ring (10798, -1.11 DPS) [dungeon]; Belt of the Gladiator (13134, -1.11 DPS) [world_drop]; Defiler's Leather Girdle (20192, -1.43 DPS, sim-verified) [rep] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 34.6 attack_power points (2.68 DPS) | yes | Scarlet Leggings (10330, -0.12 DPS) [dungeon]; Silvershell Leggings (10633, -0.24 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.24 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 25.7 attack_power points (1.99 DPS) | yes | Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Officer's Sabatons (250561, -0.21 DPS) [crafted]; Officer's Boots (250546, -0.25 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.86 DPS) | yes | Mark of Kern (2262, -0.31 DPS) [dungeon]; Assault Band (13095, -0.31 DPS) [world_drop]; Legionnaire's Band (19511, -0.44 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.63 DPS) | yes | Assault Band (13095, -0.08 DPS) [world_drop]; Legionnaire's Band (19511, -0.21 DPS) [rep]; Mark of Kern (2262, -1.03 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (129.9 DPS) | yes | Frozen Heart of the Mountain (249469, -3.24 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (129.9 DPS) | yes | Frozen Heart of the Mountain (249469, -2.79 DPS, sim-verified) [crafted] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (129.9 DPS) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -4.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (129.9 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Dark Iron Rifle (16004, -1.60 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 195.1. Weights run: 1.3s. Verify run: 1.0s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.801, strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=1.448 ± 0.116 per rating point (14 rating = 1%, 20.276 per %), hit=0.304 ± 0.015 per rating point (10 rating = 1%, 3.037 per %), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 116.6 attack_power points (3.73 DPS) | yes | Warlord's Plate Headpiece (231535, +0.00 DPS) [pvp]; Champion's Plate Helm (227043, -0.37 DPS) [pvp]; Crown of Heroism (226860, -8.37 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-verified (195.1 DPS) | yes | Medallion of the Dawn (22659, -0.41 DPS) [quest]; Conqueror's Medallion (12059, -0.46 DPS) [quest]; Rage of Mugamba (19577, -1.50 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 86.3 attack_power points (2.76 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Spaulders (272108, -1.23 DPS, sim-verified) [vendor] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 66.1 attack_power points (2.12 DPS) | yes | Cloak of Revanchion (23127, -0.56 DPS) [dungeon]; Phantasmal Cloak (18689, -0.62 DPS) [dungeon]; Shadewood Cloak (18328, -1.27 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (195.1 DPS) | yes | Savage Gladiator Chain (11726, -1.91 DPS) [dungeon]; Legionnaire's Plate Hauberk (227049, -2.05 DPS) [pvp]; Breastplate of Undead Slaying (23087, -12.23 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (195.1 DPS) | yes | General's Plate Armguards (16546, -0.28 DPS) [pvp]; Gordok Bracers of Power (18533, -0.49 DPS) [dungeon]; Bracers of Undead Slaying (23090, -5.58 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (195.1 DPS) | yes | Radiant Gloves of the Dawn (227817, -0.53 DPS) [vendor]; General's Plate Gauntlets (231532, -0.56 DPS) [pvp]; Razor Gauntlets (18326, -6.55 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 125.2 attack_power points (4.01 DPS) | yes | Ferocity of the Timbermaw (227805, -0.25 DPS) [vendor]; General's Plate Girdle (16547, -0.53 DPS) [pvp]; Heavy Obsidian Belt (22197, -0.90 DPS) [crafted] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (195.1 DPS) | yes | Titanic Leggings (22385, -0.18 DPS) [crafted]; Outrider's Plate Legguards (22651, -0.52 DPS) [rep]; Cloudkeeper Legplates (14554, -5.20 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 80.8 attack_power points (2.59 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Clutchlord's Stompers (275627, -0.22 DPS) [crafted] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (195.1 DPS) | yes | Myrmidon's Signet (2246, -0.37 DPS) [world_drop]; Don Julio's Band (19325, -0.57 DPS) [rep]; Naglering (11669, -3.89 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-verified (195.1 DPS) | yes | Myrmidon's Signet (2246, -0.28 DPS) [world_drop]; Don Julio's Band (19325, -0.48 DPS) [rep]; Naglering (11669, -4.06 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (195.1 DPS) | yes | Counterattack Lodestone (18537, -0.71 DPS) [dungeon]; Hand of Justice (11815, -0.77 DPS) [dungeon]; Blackhand's Breadth (13965, -2.10 DPS, sim-verified) [quest] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (195.1 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, +0.00 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (195.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -8.13 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | The Legend of Korrak [quest] | sim-verified (195.1 DPS) | yes | Unsophisticated Hand Cannon (18460, -0.22 DPS) [dungeon]; Monolithic Bow (9426, -0.37 DPS) [dungeon]; Dark Iron Rifle (16004, -2.44 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Amulet of the Darkmoon; shoulder: Defiler's Plate Spaulders; back: Shroud of Domination; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Legionnaire's Band; finger2: Band of the Ogre King; main_hand: The Unstoppable Force; ranged: Bloodseeker

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

