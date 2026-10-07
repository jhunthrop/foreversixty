# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 05321000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 26.9. Weights run: 2.0s. Verify run: 0.9s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.004, strength=2.063 ± 0.008, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.051 ± 0.001 per rating point (10 rating = 1%, 0.508 per %), melee_haste=1.701 ± 0.053

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.6 attack_power points (0.80 DPS) | yes | Defender's Leather Hood (252447, -0.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.2 attack_power points (0.24 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.32 DPS) | yes | Grave Shroud (279865, -0.08 DPS) [quest]; Catacomb Cloak (279899, -0.09 DPS) [quest]; Dark Leather Cloak (2316, -0.16 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.6 attack_power points (0.80 DPS) | yes | Veteran's Chain Shirt (250488, -0.21 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.24 DPS) [crafted]; Totemic Leather Armor (252435, -0.24 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.3 attack_power points (0.40 DPS) | yes | Cryptwalker Bracers (280095, -0.08 DPS) [quest]; Runed Copper Bracers (2854, -0.24 DPS) [crafted]; Bristlebark Bindings (14569, -0.24 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.5 attack_power points (0.64 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Polar Gauntlets (7606, -0.16 DPS) [quest]; Blackened Defias Gloves (10401, -0.16 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.70 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.30 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.7 attack_power points (0.88 DPS) | yes | Veteran's Chain Leggings (250493, -0.14 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.16 DPS) [crafted]; Totemic Leather Pants (252446, -0.16 DPS) [crafted] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 10.3 attack_power points (0.40 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 8.3 attack_power points (0.32 DPS) | yes | The 1 Ring (8350, -0.24 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 attack_power points (0.32 DPS) | yes | Ring of the Moon (12052, -0.21 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.24 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.9 attack_power points (12.50 DPS) | yes | Smite's Mighty Hammer (7230, -1.00 DPS, sim-verified) [dungeon]; Living Root (6631, -1.00 DPS) [dungeon]; Duskbringer (2205, -1.10 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567)) | Blacksmithing [crafted] | 4.1 attack_power points (0.16 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, -0.00 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.08 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Protector's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 05325213000000000-000000000000000000-000000000000000000)

Set DPS (verified): 68.8. Weights run: 2.2s. Verify run: 1.1s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.009, strength=1.967 ± 0.013, agility=0.014 ± 0.001, crit=0.020 ± 0.002 per rating point (14 rating = 1%, 0.275 per %), hit=0.079 ± 0.003 per rating point (10 rating = 1%, 0.795 per %), melee_haste=2.700 ± 0.608

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.6 attack_power points (1.50 DPS) | yes | Veteran's Chain Helm (250498, -0.12 DPS) [crafted]; Defender's Leather Helm (252455, -0.12 DPS) [crafted]; Crusader's Chain Helm (250502, -0.23 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.82 DPS) | yes | River Pride Choker (13087, -0.36 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.39 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.82 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.8 attack_power points (0.81 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.11 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.12 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.59 DPS) | yes | Sergeant Major's Cape (16315, -0.12 DPS) [pvp]; Lambent Scale Cloak (4706, -0.13 DPS) [world_drop]; Slayer's Cape (14752, -0.13 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.5 attack_power points (1.73 DPS) | yes | Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.35 DPS) [crafted]; Hard Gold Cuirass (250533, -0.46 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.7 attack_power points (0.92 DPS) | yes | Yorgen Bracers (13012, -0.23 DPS) [world_drop]; Bands of Serra'kis (6902, -0.23 DPS) [dungeon]; Cultist's Armguards (270032, -0.34 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 21.9 attack_power points (1.29 DPS) | yes | The Frozen Clutch (23170, -0.11 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.25 DPS) [world]; Heavy Earthen Gloves (7359, -0.35 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.41 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.02 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.02 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.53 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.25 DPS) [crafted]; Golden Scale Leggings (3843, -0.26 DPS) [crafted]; Chausses of Westfall (6087, -0.26 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 13.9 attack_power points (0.81 DPS) | yes | Hard Gold Boots (250534, -0.01 DPS) [crafted]; Disjointed Shoes (277226, -0.11 DPS) [quest]; Glimmering Mail Greaves (4073, -0.12 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 15.8 attack_power points (0.93 DPS) | yes | Tiger Band (6749, -0.23 DPS) [quest]; Silverlaine's Family Seal (6321, -0.35 DPS) [dungeon]; Insurgent's Band (272067, -0.40 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 11.9 attack_power points (0.70 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.12 DPS) [dungeon]; Insurgent's Band (272067, -0.17 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (68.8 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.16 DPS) [dungeon]; Viscous Hammer (13045, -18.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Double-barreled Shotgun (2098, -0.18 DPS) [world_drop]; Long Battle Bow (15284, -0.18 DPS) [world_drop]; Fine Longbow (11304, -0.29 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 05325213032310001-000000000000000000-000000000000000000)

Set DPS (verified): 105.2. Weights run: 2.6s. Verify run: 1.1s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.021, strength=2.006 ± 0.022, agility=0.038 ± 0.002, crit=0.054 ± 0.003 per rating point (14 rating = 1%, 0.761 per %), hit=0.103 ± 0.003 per rating point (10 rating = 1%, 1.030 per %), melee_haste=2.995 ± 0.380

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 28.1 attack_power points (1.44 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.06 DPS) [dungeon]; Tusken Helm (6686, -0.10 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.03 DPS) | yes | Kaleidoscope Chain (13084, -0.61 DPS) [world_drop]; River Pride Choker (13087, -0.61 DPS) [world_drop]; Ghostshard Talisman (7731, -0.85 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.1 attack_power points (1.13 DPS) | yes | Chromite Pauldrons (8144, -0.10 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.21 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 12.3 attack_power points (0.63 DPS) | yes | Dark Hooded Cape (5257, -0.20 DPS) [world]; Slayer's Cape (14752, -0.22 DPS) [world_drop]; Wolfmaster Cape (6314, -0.62 DPS, sim-verified) [dungeon] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.1 attack_power points (1.54 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.09 DPS) [quest]; Shining Mithril Breastplate (250540, -0.10 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.03 DPS) | yes | Ravager's Armguards (14770, -0.30 DPS) [world_drop]; Yorgen Bracers (13012, -0.40 DPS) [world_drop]; Pugilist Bracers (4438, -0.63 DPS, sim-verified) [dungeon] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.1 attack_power points (1.65 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.31 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.39 DPS) [dungeon] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 30.1 attack_power points (1.54 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Scarlet Belt (10329, -0.31 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.1 attack_power points (2.16 DPS) | yes | Orcish War Leggings (7929, -0.41 DPS) [crafted]; Symbolic Legplates (14829, -0.61 DPS) [world_drop]; Firemane Leggings (13129, -0.63 DPS, sim-verified) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 26.3 attack_power points (1.35 DPS) | yes | Skirmisher's Mail Boots (252564, -0.22 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.63 DPS, sim-verified) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.03 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.20 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.03 DPS) | yes | Thunderbrow Ring (13097, -0.20 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor]; Protector's Band (19515, -0.57 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (105.2 DPS) | yes | Bonebiter (6830, -1.22 DPS) [quest]; Darkspear Raider's Reaper (272081, -1.56 DPS) [vendor]; The Jackhammer (9423, -13.30 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (105.2 DPS) | yes | Monolithic Bow (9426, -0.09 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.20 DPS) [crafted]; Bow of Searing Arrows (2825, -1.43 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: The Silencer

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 05325213032310001-050500000000000000-000000000000000000)

Set DPS (verified): 149.0. Weights run: 2.8s. Verify run: 1.3s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.034, strength=1.994 ± 0.035, agility=0.388 ± 0.026, crit=0.555 ± 0.037 per rating point (14 rating = 1%, 7.769 per %), hit=0.149 ± 0.005 per rating point (10 rating = 1%, 1.487 per %), melee_haste=3.581 ± 0.676

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.61 DPS) | yes | Fury Visor (20521, +0.00 DPS, sim-verified) [quest]; Bloomsprout Headpiece (17767, -0.65 DPS) [dungeon]; Knight-Lieutenant's Plate Helm (220804, -0.70 DPS) [vendor] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.09 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.33 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.57 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 29.7 attack_power points (1.62 DPS) | yes | Wyrmslayer Spaulders (13066, -0.15 DPS) [world_drop]; Earthslag Shoulders (11632, -0.21 DPS) [dungeon]; Officer's Pauldrons (250576, -0.76 DPS, sim-verified) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 17.9 attack_power points (0.98 DPS) | yes | Blackveil Cape (11626, -0.03 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.20 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 47.9 attack_power points (2.61 DPS) | yes | Mixologist's Tunic (12793, -0.42 DPS) [dungeon]; Valorous Chestguard (8274, -0.43 DPS) [world_drop]; Coldmetal Guard (274758, -0.54 DPS) [vendor] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.53 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Runed Golem Shackles (12550, -0.00 DPS) [dungeon]; Officer's Wristguards (250581, -0.20 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 47.8 attack_power points (2.60 DPS) | yes | Gauntlets of Divinity (7724, -0.86 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.86 DPS) [crafted]; Officer's Gloves (250551, -1.46 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (2.50 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.55 DPS) [dungeon]; Belt of the Gladiator (13134, -0.55 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.73 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 43.9 attack_power points (2.39 DPS) | yes | Scarlet Leggings (10330, -0.11 DPS) [dungeon]; Silvershell Leggings (10633, -0.22 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.22 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 33.0 attack_power points (1.80 DPS) | yes | Prowler's Leather Boots (252468, -0.15 DPS) [crafted]; Officer's Sabatons (250561, -0.19 DPS) [crafted]; Officer's Boots (250546, -0.24 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 23.4 attack_power points (1.28 DPS) | yes | Mark of Kern (2262, -0.19 DPS) [dungeon]; Assault Band (13095, -0.19 DPS) [world_drop]; Thunderbrow Ring (13097, -0.34 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.5 attack_power points (1.17 DPS) | yes | Assault Band (13095, -0.08 DPS) [world_drop]; Thunderbrow Ring (13097, -0.24 DPS) [world_drop]; Mark of Kern (2262, -2.13 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (149.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (149.0 DPS) | yes | Molten Heart of the Mountain (249470, -0.94 DPS, sim-verified) [crafted] |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (149.0 DPS) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -7.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (149.0 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.11 DPS) [dungeon]; Dark Iron Rifle (16004, -1.57 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 05325213032310001-050500000000000000-500500000000000000)

Set DPS (verified): 228.9. Weights run: 2.8s. Verify run: 1.3s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.044, strength=2.009 ± 0.046, agility=0.496 ± 0.034, crit=0.709 ± 0.049 per rating point (14 rating = 1%, 9.919 per %), hit=0.207 ± 0.007 per rating point (10 rating = 1%, 2.071 per %), melee_haste=5.891 ± 0.852

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 60.1 attack_power points (3.34 DPS) | yes | Field Marshal's Plate Helm (231538, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -0.33 DPS) [vendor]; Crown of Heroism (226860, -9.09 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (228.9 DPS) | yes | Imperial Jewel (11933, -0.11 DPS) [dungeon]; Will of the Martyr (17044, -0.22 DPS) [quest]; Rage of Mugamba (19577, -2.45 DPS, sim-verified) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 44.6 attack_power points (2.48 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, -0.03 DPS) [pvp] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 36.5 attack_power points (2.03 DPS) | yes | Shroud of Domination (22337, -0.13 DPS) [dungeon]; Howler's Furs (272414, -0.36 DPS) [vendor]; Cape of the Black Baron (13340, -0.50 DPS) [dungeon] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (228.9 DPS) | yes | Timbermaw Tunic (252484, -0.04 DPS) [crafted]; Cadaverous Armor (14637, -0.33 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -12.32 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-verified (228.9 DPS) | yes | Windtalker's Wristguards (19582, -0.23 DPS) [rep]; Marshal's Plate Bracers (16481, -0.25 DPS) [pvp]; Bracers of Undead Slaying (23090, -6.56 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (228.9 DPS) | yes | Radiant Gloves of the Dawn (227817, -0.47 DPS) [vendor]; Marshal's Plate Gauntlets (231541, -0.53 DPS) [pvp]; Razor Gauntlets (18326, -6.85 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 64.2 attack_power points (3.57 DPS) | yes | Ferocity of the Timbermaw (227805, -0.19 DPS) [vendor]; Marshal's Plate Girdle (16482, -0.44 DPS) [pvp]; Dense Timbermaw Belt (227807, -2.26 DPS, sim-verified) [vendor] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (228.9 DPS) | yes | Titanic Leggings (22385, -0.10 DPS) [crafted]; Warbear Woolies (15065, -0.61 DPS) [crafted]; Cloudkeeper Legplates (14554, -6.01 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 42.3 attack_power points (2.35 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Scalegut Treaders (275618, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (228.9 DPS) | yes | Don Julio's Band (19325, -0.09 DPS) [rep]; Myrmidon's Signet (2246, -0.33 DPS) [world_drop]; Naglering (11669, -5.06 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-verified (228.9 DPS) | yes | Don Julio's Band (19325, -0.01 DPS) [rep]; Myrmidon's Signet (2246, -0.25 DPS) [world_drop]; Naglering (11669, -4.32 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (228.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (228.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, -3.63 DPS, sim-verified) [dungeon] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (228.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -4.81 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (228.9 DPS) | yes | Stinging Bow (10624, -0.11 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.11 DPS) [world_drop]; Dark Iron Rifle (16004, -4.34 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Highlander's Plate Spaulders; back: Cloak of the Honor Guard; chest: Obsidian Mail Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Protector's Band; finger2: Band of the Ogre King; trinket2: Blackhand's Breadth; main_hand: Blackblade of Shahram; ranged: Riphook

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 05321000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 30.4. Weights run: 2.0s. Verify run: 0.8s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.004, strength=2.063 ± 0.008, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.051 ± 0.001 per rating point (10 rating = 1%, 0.508 per %), melee_haste=1.701 ± 0.053

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.6 attack_power points (0.80 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.2 attack_power points (0.24 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.32 DPS) | yes | Grave Shroud (279865, -0.08 DPS) [quest]; Subterranean Cape (14149, -0.08 DPS, sim-verified) [dungeon]; Catacomb Cloak (279899, -0.09 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.6 attack_power points (0.80 DPS) | yes | Defender's Leather Armor (252434, -0.24 DPS) [crafted]; Totemic Leather Armor (252435, -0.24 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.25 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.3 attack_power points (0.40 DPS) | yes | Raptorcrest Bracers (270010, -0.17 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.24 DPS) [crafted]; Bristlebark Bindings (14569, -0.24 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.5 attack_power points (0.64 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.16 DPS) [dungeon]; Foreman's Gloves (2167, -0.24 DPS) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.70 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.30 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.6 attack_power points (0.72 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 10.3 attack_power points (0.40 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.3 attack_power points (0.32 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; The 1 Ring (8350, -0.24 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 attack_power points (0.32 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.24 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.9 attack_power points (12.50 DPS) | yes | Forsaken Greataxe (251533, -0.75 DPS) [quest]; Smite's Mighty Hammer (7230, -0.94 DPS) [dungeon]; Hammerbone (270018, -3.74 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.1 attack_power points (0.16 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Heavy Shortbow (3036, -0.08 DPS) [world_drop]; Orcish Battle Bow (5346, -0.08 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Legionnaire's Band; main_hand: The Axe of Severing; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 05325213000000000-000000000000000000-000000000000000000)

Set DPS (verified): 69.9. Weights run: 2.2s. Verify run: 1.1s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.009, strength=1.967 ± 0.013, agility=0.014 ± 0.001, crit=0.020 ± 0.002 per rating point (14 rating = 1%, 0.275 per %), hit=0.079 ± 0.003 per rating point (10 rating = 1%, 0.795 per %), melee_haste=2.700 ± 0.608

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.6 attack_power points (1.50 DPS) | yes | Veteran's Chain Helm (250498, -0.12 DPS) [crafted]; Defender's Leather Helm (252455, -0.12 DPS) [crafted]; Crusader's Chain Helm (250502, -0.23 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.82 DPS) | yes | Kaleidoscope Chain (13084, -0.33 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.36 DPS) [world_drop]; Scout's Medallion (19537, -0.82 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.8 attack_power points (0.81 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.11 DPS) [crafted]; Elite Shoulders (4835, -0.12 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.59 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.13 DPS) [world_drop]; Slayer's Cape (14752, -0.13 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 29.5 attack_power points (1.73 DPS) | yes | Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.35 DPS) [crafted]; Hard Gold Cuirass (250533, -0.46 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.7 attack_power points (0.92 DPS) | yes | Yorgen Bracers (13012, -0.23 DPS) [world_drop]; Bands of Serra'kis (6902, -0.23 DPS) [dungeon]; Cultist's Armguards (270032, -0.34 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 21.9 attack_power points (1.29 DPS) | yes | The Frozen Clutch (23170, -0.11 DPS) [dungeon]; Warsong Gauntlets (16978, -0.13 DPS) [quest]; Bonefist Gauntlets (4465, -0.25 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.41 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.02 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.02 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.53 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.25 DPS) [crafted]; Golden Scale Leggings (3843, -0.26 DPS) [crafted]; Slayer's Pants (14757, -0.26 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 13.9 attack_power points (0.81 DPS) | yes | Hard Gold Boots (250534, -0.01 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.12 DPS) [world_drop]; Slayer's Slippers (14756, -0.12 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 15.8 attack_power points (0.93 DPS) | yes | Tiger Band (6749, -0.23 DPS) [quest]; Silverlaine's Family Seal (6321, -0.35 DPS) [dungeon]; Insurgent's Band (272067, -0.40 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 11.9 attack_power points (0.70 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.12 DPS) [dungeon]; Insurgent's Band (272067, -0.17 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (69.9 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.16 DPS) [dungeon]; Viscous Hammer (13045, -18.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Double-barreled Shotgun (2098, -0.18 DPS) [world_drop]; Long Battle Bow (15284, -0.18 DPS) [world_drop]; Fine Longbow (11304, -0.29 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 05325213032310001-000000000000000000-000000000000000000)

Set DPS (verified): 102.9. Weights run: 2.6s. Verify run: 1.1s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.021, strength=2.006 ± 0.022, agility=0.038 ± 0.002, crit=0.054 ± 0.003 per rating point (14 rating = 1%, 0.761 per %), hit=0.103 ± 0.003 per rating point (10 rating = 1%, 1.030 per %), melee_haste=2.995 ± 0.380

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 28.1 attack_power points (1.44 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.06 DPS) [dungeon]; Tusken Helm (6686, -0.10 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.03 DPS) | yes | Ghostshard Talisman (7731, -0.31 DPS) [dungeon]; Ethereal Talisman (4430, -0.50 DPS) [quest]; Kaleidoscope Chain (13084, -0.61 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.1 attack_power points (1.13 DPS) | yes | Chromite Pauldrons (8144, -0.10 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.21 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 12.3 attack_power points (0.63 DPS) | yes | Wolfmaster Cape (6314, -0.12 DPS) [dungeon]; Wildhunter Cloak (16658, -0.12 DPS) [quest]; Dark Hooded Cape (5257, -0.20 DPS) [world] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.1 attack_power points (1.54 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.09 DPS) [quest]; Shining Mithril Breastplate (250540, -0.10 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.03 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.1 attack_power points (1.65 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.31 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.39 DPS) [dungeon] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 30.1 attack_power points (1.54 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.21 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.1 attack_power points (2.16 DPS) | yes | Firemane Leggings (13129, -0.21 DPS) [world_drop]; Orcish War Leggings (7929, -0.41 DPS) [crafted]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 26.3 attack_power points (1.35 DPS) | yes | Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.22 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.03 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.20 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.03 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.20 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (102.9 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -1.01 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (102.9 DPS) | yes | Monolithic Bow (9426, -0.09 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.20 DPS) [crafted]; Bow of Searing Arrows (2825, -1.79 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Fiery War Axe; ranged: The Silencer

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 05325213032310001-050500000000000000-000000000000000000)

Set DPS (verified): 152.5. Weights run: 2.8s. Verify run: 1.4s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.034, strength=1.994 ± 0.035, agility=0.388 ± 0.026, crit=0.555 ± 0.037 per rating point (14 rating = 1%, 7.769 per %), hit=0.149 ± 0.005 per rating point (10 rating = 1%, 1.487 per %), melee_haste=3.581 ± 0.676

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloomsprout Headpiece (17767, -0.03 DPS) [dungeon]; Blood Guard's Plate Helm (220803, -0.08 DPS) [vendor]; Embrace of the Lycan (9479, -3.92 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.09 DPS) | yes | Woven Ivy Necklace (19159, -0.25 DPS) [quest]; Skibi's Pendant (13089, -0.27 DPS) [world_drop]; Ghostshard Talisman (7731, -0.33 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 29.0 attack_power points (1.58 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.11 DPS) [world_drop]; Earthslag Shoulders (11632, -0.17 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 17.9 attack_power points (0.98 DPS) | yes | Blackveil Cape (11626, +0.00 DPS, sim-verified) [dungeon]; First Sergeant's Cloak (16340, -0.20 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 47.9 attack_power points (2.61 DPS) | yes | Mixologist's Tunic (12793, -0.42 DPS) [dungeon]; Valorous Chestguard (8274, -0.43 DPS) [world_drop]; Coldmetal Guard (274758, -0.54 DPS) [vendor] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.53 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 47.8 attack_power points (2.60 DPS) | yes | Officer's Gloves (250551, -0.67 DPS) [crafted]; Gauntlets of Divinity (7724, -0.86 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.86 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (2.50 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.55 DPS) [dungeon]; Belt of the Gladiator (13134, -0.55 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.73 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 43.9 attack_power points (2.39 DPS) | yes | Scarlet Leggings (10330, -0.11 DPS) [dungeon]; Silvershell Leggings (10633, -0.22 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.22 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 33.0 attack_power points (1.80 DPS) | yes | Prowler's Leather Boots (252468, -0.15 DPS) [crafted]; Officer's Sabatons (250561, -0.19 DPS) [crafted]; Officer's Boots (250546, -0.24 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.31 DPS) | yes | Mark of Kern (2262, -0.22 DPS) [dungeon]; Assault Band (13095, -0.22 DPS) [world_drop]; Legionnaire's Band (19511, -1.86 DPS, sim-verified) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mark of Kern (2262, -0.08 DPS) [dungeon]; Assault Band (13095, -0.08 DPS) [world_drop]; Legionnaire's Band (19511, -1.69 DPS, sim-verified) [rep] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.27 DPS) [crafted]; Molten Heart of the Mountain (249470, -4.30 DPS, sim-verified) [crafted] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -3.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.11 DPS) [dungeon]; Dark Iron Rifle (16004, -1.89 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Rune of the Guard Captain; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 05325213032310001-050500000000000000-500500000000000000)

Set DPS (verified): 222.7. Weights run: 2.8s. Verify run: 1.2s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.044, strength=2.009 ± 0.046, agility=0.496 ± 0.034, crit=0.709 ± 0.049 per rating point (14 rating = 1%, 9.919 per %), hit=0.207 ± 0.007 per rating point (10 rating = 1%, 2.071 per %), melee_haste=5.891 ± 0.852

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 60.1 attack_power points (3.34 DPS) | yes | Warlord's Plate Headpiece (231535, +0.00 DPS) [pvp]; Champion's Plate Helm (227043, -0.33 DPS) [pvp]; Crown of Heroism (226860, -7.50 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (222.7 DPS) | yes | Imperial Jewel (11933, -0.11 DPS) [dungeon]; Will of the Martyr (17044, -0.22 DPS) [quest]; Rage of Mugamba (19577, -2.15 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 44.6 attack_power points (2.48 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Champion's Plate Shoulders (227042, -0.03 DPS) [pvp]; Black Dragonscale Shoulders (15051, -1.58 DPS, sim-verified) [crafted] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 36.5 attack_power points (2.03 DPS) | yes | Shroud of Domination (22337, -0.13 DPS) [dungeon]; Howler's Furs (272414, -0.36 DPS) [vendor]; Cape of the Black Baron (13340, -0.50 DPS) [dungeon] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (222.7 DPS) | yes | Timbermaw Tunic (252484, -0.04 DPS) [crafted]; Cadaverous Armor (14637, -0.33 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -11.18 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (222.7 DPS) | yes | Windtalker's Wristguards (19582, -0.23 DPS) [rep]; General's Plate Armguards (16546, -0.25 DPS) [pvp]; Bracers of Undead Slaying (23090, -6.30 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (222.7 DPS) | yes | Radiant Gloves of the Dawn (227817, -0.47 DPS) [vendor]; General's Plate Gauntlets (231532, -0.53 DPS) [pvp]; Razor Gauntlets (18326, -6.64 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 64.2 attack_power points (3.57 DPS) | yes | Ferocity of the Timbermaw (227805, -0.19 DPS) [vendor]; General's Plate Girdle (16547, -0.44 DPS) [pvp]; Dense Timbermaw Belt (227807, -1.72 DPS, sim-verified) [vendor] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (222.7 DPS) | yes | Titanic Leggings (22385, -0.10 DPS) [crafted]; Outrider's Plate Legguards (22651, -0.44 DPS) [rep]; Cloudkeeper Legplates (14554, -5.77 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 42.3 attack_power points (2.35 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Scalegut Treaders (275618, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (222.7 DPS) | yes | Don Julio's Band (19325, -0.09 DPS) [rep]; White Bone Band (11862, -0.31 DPS) [quest]; Naglering (11669, -4.88 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-verified (222.7 DPS) | yes | Don Julio's Band (19325, -0.01 DPS) [rep]; White Bone Band (11862, -0.23 DPS) [quest]; Naglering (11669, -4.25 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (222.7 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -3.33 DPS, sim-verified) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (222.7 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (222.7 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -7.88 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (222.7 DPS) | yes | Stinging Bow (10624, -0.11 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.11 DPS) [world_drop]; Dark Iron Rifle (16004, -2.74 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Defiler's Plate Spaulders; back: Deathguard's Cloak; chest: Obsidian Mail Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Boots of Heroism; finger1: Legionnaire's Band; finger2: Band of the Ogre King; trinket1: Hand of Justice; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Riphook

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

