# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 05321000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.9. Weights run: 2.5s. Verify run: 1.2s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.001, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.054 ± 0.001 per rating point (10 rating = 1%, 0.539 per %), melee_haste=1.768 ± 0.046

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
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.16 DPS) | yes | Living Root (6631, -0.87 DPS) [dungeon]; Duskbringer (2205, -0.96 DPS) [dungeon]; Smite's Mighty Hammer (7230, -1.00 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.14 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Protector's Band; main_hand: The Axe of Severing; ranged: Fine Longbow

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 05325213000000000-00000000000000000-000000000000000000)

Set DPS (verified): 71.9. Weights run: 2.9s. Verify run: 1.5s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.022 ± 0.002, crit=0.031 ± 0.002 per rating point (14 rating = 1%, 0.441 per %), hit=0.136 ± 0.006 per rating point (10 rating = 1%, 1.361 per %), melee_haste=not significant (4.590 ± 1.331)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.95 DPS) | yes | Veteran's Chain Helm (250498, -0.07 DPS) [crafted]; Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Crusader's Chain Helm (250502, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.51 DPS) | yes | River Pride Choker (13087, -0.22 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.34 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.51 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.51 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.07 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.07 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.37 DPS) | yes | Sergeant Major's Cape (16315, -0.07 DPS) [pvp]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.10 DPS) | yes | Shining Silver Breastplate (2870, -0.07 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.22 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.59 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS) [world_drop]; Bands of Serra'kis (6902, -0.15 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.22 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.81 DPS) | yes | The Frozen Clutch (23170, -0.07 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.15 DPS) [world]; Mail Combat Gauntlets (4075, -0.22 DPS) [world_drop] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 attack_power points (0.88 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.95 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.14 DPS) [crafted]; Golden Scale Leggings (3843, -0.15 DPS) [crafted]; Chausses of Westfall (6087, -0.15 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.2 attack_power points (0.52 DPS) | yes | Hard Gold Boots (250534, -0.01 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.1 attack_power points (0.59 DPS) | yes | Tiger Band (6749, -0.15 DPS) [quest]; Silverlaine's Family Seal (6321, -0.22 DPS) [dungeon]; Insurgent's Band (272067, -0.26 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.1 attack_power points (0.45 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.08 DPS) [dungeon]; Insurgent's Band (272067, -0.11 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (71.9 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.10 DPS) [dungeon]; Viscous Hammer (13045, -19.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.33 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.11 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.18 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 05325213032310001-00000000000000000-000000000000000000)

Set DPS (verified): 119.5. Weights run: 3.3s. Verify run: 1.5s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.052 ± 0.003, crit=0.074 ± 0.004 per rating point (14 rating = 1%, 1.042 per %), hit=0.184 ± 0.005 per rating point (10 rating = 1%, 1.842 per %), melee_haste=5.844 ± 0.623

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 28.0 attack_power points (1.25 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.04 DPS) [dungeon]; Tusken Helm (6686, -0.09 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.90 DPS) | yes | Ghostshard Talisman (7731, -0.27 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.53 DPS) [world_drop]; River Pride Choker (13087, -0.54 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.99 DPS) | yes | Chromite Pauldrons (8144, -0.09 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.18 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 12.3 attack_power points (0.55 DPS) | yes | Wolfmaster Cape (6314, -0.10 DPS) [dungeon]; Dark Hooded Cape (5257, -0.17 DPS) [world]; Slayer's Cape (14752, -0.19 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.34 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.07 DPS) [quest]; Shining Mithril Breastplate (250540, -0.09 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.90 DPS) | yes | Pugilist Bracers (4438, -0.18 DPS) [dungeon]; Ravager's Armguards (14770, -0.26 DPS) [world_drop]; Yorgen Bracers (13012, -0.35 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.43 DPS) | yes | Truesilver Gauntlets (7938, -0.00 DPS) [crafted]; Reticulated Bone Gauntlets (9435, -0.27 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.33 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.34 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.00 DPS) [rep]; Highlander's Chain Girdle (20090, -0.27 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.88 DPS) | yes | Firemane Leggings (13129, -0.18 DPS) [world_drop]; Orcish War Leggings (7929, -0.36 DPS) [crafted]; Symbolic Legplates (14829, -0.52 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 26.4 attack_power points (1.18 DPS) | yes | Prowler's Leather Shoes (252465, -0.18 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.20 DPS) [crafted]; Obsidian Greaves (13068, -0.27 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.90 DPS) | yes | Protector's Band (19515, -0.16 DPS) [rep]; Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Suspicious Spare Part (274754, -0.27 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.90 DPS) | yes | Protector's Band (19515, -0.16 DPS) [rep]; Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Suspicious Spare Part (274754, -0.27 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (119.5 DPS) | yes | Bonebiter (6830, -1.07 DPS) [quest]; Darkspear Raider's Reaper (272081, -1.37 DPS) [vendor]; The Jackhammer (9423, -10.28 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (119.5 DPS) | yes | Monolithic Bow (9426, -0.08 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.18 DPS) [crafted]; Bow of Searing Arrows (2825, -1.03 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: The Silencer

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 05325213032310001-05050000000000000-000000000000000000)

Set DPS (verified): 170.4. Weights run: 3.6s. Verify run: 1.8s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.732 ± 0.044, crit=1.045 ± 0.063 per rating point (14 rating = 1%, 14.630 per %), hit=0.264 ± 0.008 per rating point (10 rating = 1%, 2.640 per %), melee_haste=9.189 ± 1.068

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.30 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -0.35 DPS) [dungeon]; Fury Visor (20521, -0.51 DPS) [quest] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.96 DPS) | yes | Skibi's Pendant (13089, -0.02 DPS) [world_drop]; Ghostshard Talisman (7731, -0.29 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.43 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 36.6 attack_power points (1.75 DPS) | yes | Wyrmslayer Spaulders (13066, -0.32 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.38 DPS) [crafted]; Officer's Pauldrons (250576, -0.84 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.2 attack_power points (1.06 DPS) | yes | Sergeant Major's Cape (16336, -0.28 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world]; Bloodlust Cape (14801, -0.96 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 48.0 attack_power points (2.30 DPS) | yes | Knight's Plate Hauberk (220794, -0.16 DPS) [vendor]; Mixologist's Tunic (12793, -0.19 DPS) [dungeon]; Valorous Chestguard (8274, -0.38 DPS) [world_drop] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.34 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Runed Golem Shackles (12550, -0.00 DPS) [dungeon]; Officer's Wristguards (250581, -0.08 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 51.3 attack_power points (2.45 DPS) | yes | Raider Gloves (272100, -0.73 DPS) [vendor]; Gloves of Holy Might (867, -0.80 DPS) [world_drop]; Officer's Gloves (250551, -1.19 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (2.20 DPS) | yes | Prowler's Leather Waistguard (252473, -0.44 DPS) [crafted]; Highlander's Lamellar Girdle (20106, -0.45 DPS) [rep]; Atal'alarion's Tusk Ring (10798, -0.48 DPS) [dungeon] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 45.9 attack_power points (2.20 DPS) | yes | Scarlet Leggings (10330, -0.19 DPS) [dungeon]; Gryphon Rider's Leggings (9652, -0.24 DPS) [quest]; Golem Shard Leggings (13074, -1.61 DPS, sim-verified) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 37.5 attack_power points (1.79 DPS) | yes | Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Officer's Sabatons (250561, -0.24 DPS) [crafted]; Skulker's Leather Boots (252469, -0.29 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 26.6 attack_power points (1.27 DPS) | yes | Mark of Kern (2262, -0.31 DPS) [dungeon]; Assault Band (13095, -0.31 DPS) [world_drop]; Thunderbrow Ring (13097, -0.40 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 22.6 attack_power points (1.08 DPS) | yes | Assault Band (13095, -0.13 DPS) [world_drop]; Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Mark of Kern (2262, -2.72 DPS, sim-verified) [dungeon] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (170.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (170.4 DPS) | yes | Molten Heart of the Mountain (249470, -1.68 DPS, sim-verified) [crafted] |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (170.4 DPS) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -6.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (170.4 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.10 DPS) [dungeon]; Dark Iron Rifle (16004, -1.73 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 05325213032310001-05050000000000000-055000000000000000)

Set DPS (verified): 258.5. Weights run: 3.5s. Verify run: 1.8s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.946 ± 0.058, crit=1.352 ± 0.083 per rating point (14 rating = 1%, 18.924 per %), hit=0.369 ± 0.011 per rating point (10 rating = 1%, 3.693 per %), melee_haste=10.426 ± 1.396

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 81.2 attack_power points (3.83 DPS) | yes | Field Marshal's Plate Helm (231538, -0.30 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -0.78 DPS) [vendor]; Crown of Heroism (226860, -9.37 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (258.5 DPS) | yes | Amulet of the Darkmoon (19491, -0.23 DPS) [quest]; Imperial Jewel (11933, -0.51 DPS) [dungeon]; Rage of Mugamba (19577, -2.59 DPS, sim-verified) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 52.1 attack_power points (2.45 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 38.7 attack_power points (1.82 DPS) | yes | Cape of the Black Baron (13340, -0.21 DPS) [dungeon]; Shroud of Domination (22337, -0.22 DPS) [dungeon]; Howler's Furs (272414, -0.33 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (258.5 DPS) | yes | Obsidian Mail Tunic (22191, -0.33 DPS) [crafted]; Cadaverous Armor (14637, -0.87 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -11.25 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-verified (258.5 DPS) | yes | Marshal's Plate Bracers (16481, -0.23 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.26 DPS) [rep]; Bracers of Undead Slaying (23090, -4.93 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (258.5 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.26 DPS) [pvp]; Radiant Gloves of the Dawn (227817, -0.42 DPS) [vendor]; Razor Gauntlets (18326, -5.38 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 72.9 attack_power points (3.44 DPS) | yes | Ferocity of the Timbermaw (227805, -0.23 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.42 DPS) [vendor]; Marshal's Plate Girdle (16482, -0.54 DPS) [pvp] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (258.5 DPS) | yes | Titanic Leggings (22385, -0.43 DPS) [crafted]; Marshal's Plate Legguards (231540, -0.66 DPS) [pvp]; Cloudkeeper Legplates (14554, -5.24 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 45.6 attack_power points (2.15 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Boots of Heroism (21995, -0.09 DPS) [quest]; Battleboots of Heroism (226857, -0.09 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (258.5 DPS) | yes | Band of the Ogre King (18522, -0.50 DPS) [dungeon]; Myrmidon's Signet (2246, -0.56 DPS) [world_drop]; Naglering (11669, -7.11 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (258.5 DPS) | yes | Band of the Ogre King (18522, -0.30 DPS) [dungeon]; Myrmidon's Signet (2246, -0.37 DPS) [world_drop]; Naglering (11669, -4.03 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (258.5 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -4.08 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (258.5 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.79 DPS) [crafted]; Blackhand's Breadth (13965, -4.92 DPS, sim-verified) [quest] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (258.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -3.91 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-verified (258.5 DPS) | yes | Riphook (12653, -0.03 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.12 DPS) [world_drop]; Dark Iron Rifle (16004, -2.76 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Highlander's Plate Spaulders; back: Cloak of the Honor Guard; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Blackblade of Shahram; ranged: Bloodseeker

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 05321000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 30.4. Weights run: 2.5s. Verify run: 1.2s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.001, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000) per rating point (14 rating = 1%, 0.000 per %), hit=0.054 ± 0.001 per rating point (10 rating = 1%, 0.539 per %), melee_haste=1.768 ± 0.046

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.09 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.25 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Raptorcrest Bracers (270010, -0.17 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.55 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon]; Foreman's Gloves (2167, -0.21 DPS) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.14 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Hulking Belt (14746, -0.28 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.0 attack_power points (0.62 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.07 DPS) [world_drop] |
| feet | Defender's Leather Boots (252441) (or Veteran's Boots (250503), Guard's Boots (250504), Brawler's Leather Boots (252439), Totemic Leather Boots (252442)) | Leatherworking [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Veteran's Boots (250503, +0.00 DPS) [crafted]; Guard's Boots (250504, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.16 DPS) | yes | Forsaken Greataxe (251533, -0.66 DPS) [quest]; Smite's Mighty Hammer (7230, -0.83 DPS) [dungeon]; Hammerbone (270018, -3.73 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.14 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Defender's Leather Boots; finger1: Demon Band; finger2: Legionnaire's Band; main_hand: The Axe of Severing; ranged: Fine Longbow

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (orc, 05325213000000000-00000000000000000-000000000000000000)

Set DPS (verified): 73.1. Weights run: 2.9s. Verify run: 1.5s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.022 ± 0.002, crit=0.031 ± 0.002 per rating point (14 rating = 1%, 0.441 per %), hit=0.136 ± 0.006 per rating point (10 rating = 1%, 1.361 per %), melee_haste=not significant (4.590 ± 1.331)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.95 DPS) | yes | Veteran's Chain Helm (250498, -0.07 DPS) [crafted]; Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Crusader's Chain Helm (250502, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.51 DPS) | yes | River Pride Choker (13087, -0.22 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.32 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.51 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.51 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.07 DPS) [crafted]; Elite Shoulders (4835, -0.07 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.37 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.10 DPS) | yes | Shining Silver Breastplate (2870, -0.07 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.22 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.59 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS) [world_drop]; Bands of Serra'kis (6902, -0.15 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.22 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.81 DPS) | yes | Warsong Gauntlets (16978, -0.07 DPS) [quest]; The Frozen Clutch (23170, -0.07 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.15 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 attack_power points (0.88 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.95 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.14 DPS) [crafted]; Golden Scale Leggings (3843, -0.15 DPS) [crafted]; Slayer's Pants (14757, -0.15 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.2 attack_power points (0.52 DPS) | yes | Hard Gold Boots (250534, -0.01 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.1 attack_power points (0.59 DPS) | yes | Tiger Band (6749, -0.15 DPS) [quest]; Silverlaine's Family Seal (6321, -0.22 DPS) [dungeon]; Insurgent's Band (272067, -0.26 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.1 attack_power points (0.45 DPS) | yes | Tiger Band (6749, -0.00 DPS) [quest]; Silverlaine's Family Seal (6321, -0.08 DPS) [dungeon]; Insurgent's Band (272067, -0.11 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (73.1 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.10 DPS) [dungeon]; Viscous Hammer (13045, -20.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.33 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Long Battle Bow (15284, -0.11 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.18 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 05325213032310001-00000000000000000-000000000000000000)

Set DPS (verified): 119.5. Weights run: 3.3s. Verify run: 1.5s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.052 ± 0.003, crit=0.074 ± 0.004 per rating point (14 rating = 1%, 1.042 per %), hit=0.184 ± 0.005 per rating point (10 rating = 1%, 1.842 per %), melee_haste=5.844 ± 0.623

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Hard Gold Coif (250537) (or Icemetal Barbute (10763)) | Blacksmithing [crafted] | 28.0 attack_power points (1.25 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.04 DPS) [dungeon]; Tusken Helm (6686, -0.09 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.90 DPS) | yes | Ghostshard Talisman (7731, -0.27 DPS) [dungeon]; Ethereal Talisman (4430, -0.44 DPS) [quest]; Kaleidoscope Chain (13084, -0.53 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.99 DPS) | yes | Chromite Pauldrons (8144, -0.09 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.18 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 12.3 attack_power points (0.55 DPS) | yes | Wolfmaster Cape (6314, -0.10 DPS) [dungeon]; Wildhunter Cloak (16658, -0.10 DPS) [quest]; Dark Hooded Cape (5257, -0.17 DPS) [world] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.34 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.07 DPS) [quest]; Shining Mithril Breastplate (250540, -0.09 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.90 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.43 DPS) | yes | Truesilver Gauntlets (7938, -0.00 DPS) [crafted]; Reticulated Bone Gauntlets (9435, -0.27 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.33 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.34 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.18 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.88 DPS) | yes | Firemane Leggings (13129, -0.18 DPS) [world_drop]; Orcish War Leggings (7929, -0.36 DPS) [crafted]; Symbolic Legplates (14829, -0.52 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 26.4 attack_power points (1.18 DPS) | yes | Prowler's Leather Shoes (252465, -0.18 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.20 DPS) [crafted]; Obsidian Greaves (13068, -0.27 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.90 DPS) | yes | Legionnaire's Band (19512, -0.16 DPS) [rep]; Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Suspicious Spare Part (274754, -0.27 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.90 DPS) | yes | Legionnaire's Band (19512, -0.16 DPS) [rep]; Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Suspicious Spare Part (274754, -0.27 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pendulum of Doom (9425) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (119.5 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Fiery War Axe (870, -1.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (119.5 DPS) | yes | Monolithic Bow (9426, -0.08 DPS) [dungeon]; Mithril Blacksmith Hammer (285280, -0.18 DPS) [crafted]; Bow of Searing Arrows (2825, -1.01 DPS, sim-verified) [world_drop] |

**New at 40:** head: Hard Gold Coif; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Pendulum of Doom; ranged: The Silencer

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 05325213032310001-05050000000000000-000000000000000000)

Set DPS (verified): 171.0. Weights run: 3.6s. Verify run: 1.9s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.732 ± 0.044, crit=1.045 ± 0.063 per rating point (14 rating = 1%, 14.630 per %), hit=0.264 ± 0.008 per rating point (10 rating = 1%, 2.640 per %), melee_haste=9.189 ± 1.068

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.30 DPS) | yes | Blood Guard's Plate Helm (220803, -0.23 DPS) [vendor]; Raging Berserker's Helm (7719, -0.35 DPS) [dungeon]; Fury Visor (20521, -0.51 DPS) [quest] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.96 DPS) | yes | Skibi's Pendant (13089, -0.02 DPS) [world_drop]; Woven Ivy Necklace (19159, -0.07 DPS) [quest]; Ghostshard Talisman (7731, -0.29 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 31.9 attack_power points (1.52 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.10 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.16 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.2 attack_power points (1.06 DPS) | yes | First Sergeant's Cloak (16340, -0.28 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world]; Bloodlust Cape (14801, -1.17 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 48.0 attack_power points (2.30 DPS) | yes | Stone Guard's Plate Armor (220801, -0.16 DPS) [vendor]; Mixologist's Tunic (12793, -0.19 DPS) [dungeon]; Valorous Chestguard (8274, -0.38 DPS) [world_drop] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.34 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 51.3 attack_power points (2.45 DPS) | yes | Raider Gloves (272100, -0.73 DPS) [vendor]; Gloves of Holy Might (867, -0.80 DPS) [world_drop]; Officer's Gloves (250551, -1.09 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (2.20 DPS) | yes | Prowler's Leather Waistguard (252473, -0.44 DPS) [crafted]; Atal'alarion's Tusk Ring (10798, -0.48 DPS) [dungeon]; Belt of the Gladiator (13134, -0.48 DPS) [world_drop] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 45.9 attack_power points (2.20 DPS) | yes | Scarlet Leggings (10330, -0.19 DPS) [dungeon]; Sunscale Legplates (14850, -0.27 DPS) [world_drop]; Golem Shard Leggings (13074, -1.77 DPS, sim-verified) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 37.5 attack_power points (1.79 DPS) | yes | Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Officer's Sabatons (250561, -0.24 DPS) [crafted]; Skulker's Leather Boots (252469, -0.29 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 26.6 attack_power points (1.27 DPS) | yes | Mark of Kern (2262, -0.31 DPS) [dungeon]; Assault Band (13095, -0.31 DPS) [world_drop]; White Bone Band (11862, -2.70 DPS, sim-verified) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (171.0 DPS) | yes | Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop]; White Bone Band (11862, -2.00 DPS, sim-verified) [quest] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.98 DPS) [crafted]; Molten Heart of the Mountain (249470, -4.88 DPS, sim-verified) [crafted] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Blight (7959, -3.10 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Houndmaster's Bow (11628, -0.10 DPS) [dungeon]; Dark Iron Rifle (16004, -1.69 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Diamond Flask; trinket2: Rune of the Guard Captain; main_hand: Fiery War Axe; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 05325213032310001-05050000000000000-055000000000000000)

Set DPS (verified): 251.8. Weights run: 3.5s. Verify run: 1.7s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.946 ± 0.058, crit=1.352 ± 0.083 per rating point (14 rating = 1%, 18.924 per %), hit=0.369 ± 0.011 per rating point (10 rating = 1%, 3.693 per %), melee_haste=10.426 ± 1.396

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 81.2 attack_power points (3.83 DPS) | yes | Warlord's Plate Headpiece (231535, -0.30 DPS) [pvp]; Champion's Plate Helm (227043, -0.78 DPS) [pvp]; Crown of Heroism (226860, -8.92 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (251.8 DPS) | yes | Amulet of the Darkmoon (19491, -0.23 DPS) [quest]; Imperial Jewel (11933, -0.51 DPS) [dungeon]; Rage of Mugamba (19577, -2.26 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 52.1 attack_power points (2.45 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Defiler's Leather Shoulders (20194, -0.24 DPS) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 38.7 attack_power points (1.82 DPS) | yes | Cape of the Black Baron (13340, -0.21 DPS) [dungeon]; Shroud of Domination (22337, -0.22 DPS) [dungeon]; Howler's Furs (272414, -0.33 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (251.8 DPS) | yes | Obsidian Mail Tunic (22191, -0.33 DPS) [crafted]; Cadaverous Armor (14637, -0.87 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -10.82 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (251.8 DPS) | yes | General's Plate Armguards (16546, -0.23 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.26 DPS) [rep]; Bracers of Undead Slaying (23090, -4.54 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (251.8 DPS) | yes | General's Plate Gauntlets (231532, -0.26 DPS) [pvp]; Radiant Gloves of the Dawn (227817, -0.42 DPS) [vendor]; Razor Gauntlets (18326, -5.08 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 72.9 attack_power points (3.44 DPS) | yes | Ferocity of the Timbermaw (227805, -0.23 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.42 DPS) [vendor]; General's Plate Girdle (16547, -0.54 DPS) [pvp] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (251.8 DPS) | yes | Titanic Leggings (22385, -0.43 DPS) [crafted]; General's Plate Leggings (231533, -0.66 DPS) [pvp]; Cloudkeeper Legplates (14554, -5.97 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 45.6 attack_power points (2.15 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Boots of Heroism (21995, -0.09 DPS) [quest]; Battleboots of Heroism (226857, -0.09 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (251.8 DPS) | yes | Band of the Ogre King (18522, -0.50 DPS) [dungeon]; Myrmidon's Signet (2246, -0.56 DPS) [world_drop]; Naglering (11669, -7.61 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (251.8 DPS) | yes | Band of the Ogre King (18522, -0.30 DPS) [dungeon]; Myrmidon's Signet (2246, -0.37 DPS) [world_drop]; Naglering (11669, -3.76 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (251.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (251.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Diamond Flask (20130, -3.00 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (251.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -8.27 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | The Legend of Korrak [quest] | sim-verified (251.8 DPS) | yes | Riphook (12653, -0.03 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.12 DPS) [world_drop]; Dark Iron Rifle (16004, -2.25 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Defiler's Plate Spaulders; back: Deathguard's Cloak; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force; ranged: Bloodseeker

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

